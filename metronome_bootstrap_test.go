//
// File generated from our OpenAPI spec
//

package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/terraform-provider-stripe/internal/combinedprovider"
	"github.com/stripe/terraform-provider-stripe/internal/providerdata"
)

type offlineActivationProvider struct {
	*combinedprovider.Provider
	backend stripe.RawRequestBackend
}

func (p *offlineActivationProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	p.Provider.Configure(ctx, req, resp)
	if data, ok := resp.ResourceData.(*providerdata.Data); ok {
		data.ActivationBackend = p.backend
	}
}

// Only the test executable exposes this local transport. Released providers
// always use the fixed Stripe activation origin.
func TestMetronomeProviderProcess(t *testing.T) {
	if os.Getenv("METRONOME_TEST_PROVIDER_PROCESS") != "1" {
		return
	}
	pemData, err := os.ReadFile(os.Getenv("METRONOME_TEST_ACTIVATION_CA"))
	if err != nil {
		t.Fatal(err)
	}
	roots := x509.NewCertPool()
	if !roots.AppendCertsFromPEM(pemData) {
		t.Fatal("test CA is invalid")
	}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{RootCAs: roots, MinVersion: tls.VersionTLS12}
	backend := stripe.GetBackendWithConfig(stripe.APIBackend, &stripe.BackendConfig{URL: stripe.String(os.Getenv("METRONOME_TEST_ACTIVATION_ORIGIN")), HTTPClient: &http.Client{Transport: transport, Timeout: 10 * time.Second}, MaxNetworkRetries: stripe.Int64(0), LeveledLogger: &stripe.LeveledLogger{Level: stripe.LevelNull}}).(stripe.RawRequestBackend)
	err = providerserver.Serve(context.Background(), func() provider.Provider {
		return &offlineActivationProvider{Provider: combinedprovider.New("offline")().(*combinedprovider.Provider), backend: backend}
	}, providerserver.ServeOpts{Address: "registry.terraform.io/stripe/stripe"})
	if err != nil {
		t.Fatal(err)
	}
}

func requiredTerraform(t *testing.T) {
	t.Helper()
	out, err := exec.Command("terraform", "version", "-json").Output()
	if err != nil {
		t.Fatalf("required offline Terraform is unavailable: %v", err)
	}
	var version struct {
		Version string `json:"terraform_version"`
	}
	if err = json.Unmarshal(out, &version); err != nil {
		t.Fatal(err)
	}
	if version.Version != "1.16.5" {
		t.Fatalf("required offline Terraform is 1.16.5, got %s", version.Version)
	}
}

func TestMetronomeBridgeModuleOrderingAndTargetedBootstrap(t *testing.T) {
	requiredTerraform(t)
	root := t.TempDir()
	binaryDir := t.TempDir()
	testBinary, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	wrapper := filepath.Join(binaryDir, "terraform-provider-stripe")
	quote := func(value string) string { return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'" }
	if err = os.WriteFile(wrapper, []byte("#!/bin/sh\nexec "+quote(testBinary)+" '-test.run=^TestMetronomeProviderProcess$'\n"), 0700); err != nil {
		t.Fatal(err)
	}
	var mu sync.Mutex
	activated, ready, created := false, false, false
	activations, probes, earlyCalls, mutations := 0, 0, 0, 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if r.Header.Get("Authorization") != "Bearer sk_test_bootstrap" {
			t.Error("bootstrap did not use the shared Stripe key")
			http.Error(w, "unauthorized", 401)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v1/account":
			if r.Method != "GET" {
				t.Error("identity mutated")
			}
			_, _ = io.WriteString(w, `{"id":"acct_bootstrap"}`)
		case "/v1/experimental-met-activate":
			body, _ := io.ReadAll(r.Body)
			if r.Method != "POST" || len(body) != 0 {
				t.Error("activation method or payload changed")
			}
			activations++
			activated = true
			probes = 0
			_, _ = io.WriteString(w, `{"status":"complete","client_id":null}`)
		case "/v1/credit-types/list":
			if r.Method != "GET" {
				t.Error("probe method changed")
			}
			if r.URL.Query().Get("limit") == "1" {
				probes++
				if !activated || probes == 1 {
					http.Error(w, "cached authorization", 403)
					return
				}
				ready = true
				_, _ = io.WriteString(w, `{"data":[],"next_page":"must-not-follow"}`)
				return
			}
			if !ready {
				earlyCalls++
				http.Error(w, "bootstrap required", 403)
				return
			}
			if created {
				_, _ = io.WriteString(w, `{"data":[{"id":"unit_bootstrap","name":"module credits","is_currency":false}]}`)
			} else {
				_, _ = io.WriteString(w, `{"data":[]}`)
			}
		case "/v1/credit-types/create":
			if !ready {
				earlyCalls++
				http.Error(w, "bootstrap required", 403)
				return
			}
			created = true
			mutations++
			_, _ = io.WriteString(w, `{"data":{"id":"unit_bootstrap"}}`)
		case "/v1/credit-types/archive":
			created = false
			mutations++
			_, _ = io.WriteString(w, `{"data":{"id":"unit_bootstrap"}}`)
		default:
			t.Errorf("unexpected bootstrap request: %s %s", r.Method, r.URL)
			http.Error(w, "unexpected", 400)
		}
	}))
	defer server.Close()
	ca := filepath.Join(root, "ca.pem")
	if err = os.WriteFile(ca, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600); err != nil {
		t.Fatal(err)
	}
	cli := filepath.Join(root, "terraform.rc")
	if err = os.WriteFile(cli, []byte(fmt.Sprintf("provider_installation {\n dev_overrides {\n \"stripe/stripe\" = %q\n }\n direct {}\n}\n", binaryDir)), 0600); err != nil {
		t.Fatal(err)
	}
	env := terraformTestEnvironment("TF_CLI_CONFIG_FILE="+cli, "TF_IN_AUTOMATION=1", "STRIPE_API_KEY=sk_test_bootstrap", "METRONOME_TEST_PROVIDER_PROCESS=1", "METRONOME_TEST_ACTIVATION_ORIGIN="+server.URL, "METRONOME_TEST_ACTIVATION_CA="+ca)
	required := "terraform {\n required_version = \">= 1.16.5\"\n required_providers {\n stripe = { source = \"stripe/stripe\" }\n }\n}\n"
	config := required + fmt.Sprintf("provider \"stripe\" {\n alias = \"billing\"\n metronome = { base_url = %q, ca_cert_file = %q }\n}\n", server.URL, ca) + "resource \"stripe_metronome_account_bridge\" \"billing\" {\n provider = stripe.billing\n}\nmodule \"billing\" {\n source = \"./billing\"\n providers = { stripe = stripe.billing }\n depends_on = [stripe_metronome_account_bridge.billing]\n}\n"
	if err = os.WriteFile(filepath.Join(root, "main.tf"), []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	if err = os.Mkdir(filepath.Join(root, "billing"), 0700); err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(filepath.Join(root, "billing", "main.tf"), []byte(required+"resource \"stripe_metronome_credit_type\" \"credits\" {\n name = \"module credits\"\n}\n"), 0600); err != nil {
		t.Fatal(err)
	}
	run := func(args ...string) string {
		t.Helper()
		ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, "terraform", args...)
		cmd.Dir = root
		cmd.Env = env
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("terraform %s: %v\n%s", strings.Join(args, " "), err, out)
		}
		return string(out)
	}
	run("get", "-no-color")
	run("validate", "-no-color")
	run("plan", "-no-color", "-out=initial.plan")
	mu.Lock()
	before := activations
	mu.Unlock()
	if before != 0 {
		t.Fatal("planning activated Metronome")
	}
	run("apply", "-no-color", "initial.plan")
	mu.Lock()
	count, early := activations, earlyCalls
	mu.Unlock()
	if count != 1 || early != 0 {
		t.Fatalf("activation/module order: %d activations, %d early calls", count, early)
	}
	run("plan", "-no-color", "-detailed-exitcode")
	// Targeted repair must complete without refreshing existing dependent state.
	mu.Lock()
	activated = false
	ready = false
	mu.Unlock()
	run("apply", "-no-color", "-auto-approve", "-target=stripe_metronome_account_bridge.billing")
	mu.Lock()
	count, early = activations, earlyCalls
	mu.Unlock()
	if count != 2 || early != 0 {
		t.Fatalf("targeted bootstrap failed: %d activations, %d premature refreshes", count, early)
	}
	run("plan", "-no-color", "-detailed-exitcode")
	run("destroy", "-no-color", "-auto-approve")
	mu.Lock()
	defer mu.Unlock()
	if activations != 2 || mutations != 2 {
		t.Fatalf("repeat apply or destroy mutated bridge: activations=%d resource writes=%d", activations, mutations)
	}
}

func TestMetronomeInstalledProviderWithoutCredentials(t *testing.T) {
	requiredTerraform(t)
	root := t.TempDir()
	mirror := filepath.Join(root, "mirror")
	version := "99.0.0"
	install := filepath.Join(mirror, "registry.terraform.io", "stripe", "stripe", version, runtime.GOOS+"_"+runtime.GOARCH)
	if err := os.MkdirAll(install, 0700); err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(install, "terraform-provider-stripe_v"+version)
	build := exec.Command("go", "build", "-ldflags", "-X main.version="+version, "-o", binary, ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("provider build: %v\n%s", err, out)
	}
	config := fmt.Sprintf("terraform {\n required_version = \">= 1.16.5\"\n required_providers {\n stripe = { source = \"stripe/stripe\", version = %q }\n }\n}\nprovider \"stripe\" {}\nresource \"stripe_metronome_credit_type\" \"credits\" {\n name = \"validate without credentials\"\n}\n", version)
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	cli := filepath.Join(root, "terraform.rc")
	if err := os.WriteFile(cli, []byte(fmt.Sprintf("provider_installation {\n filesystem_mirror {\n path = %q\n include = [\"stripe/stripe\"]\n }\n}\n", mirror)), 0600); err != nil {
		t.Fatal(err)
	}
	run := func(args ...string) []byte {
		t.Helper()
		cmd := exec.Command("terraform", args...)
		cmd.Dir = root
		cmd.Env = terraformTestEnvironment("TF_CLI_CONFIG_FILE="+cli, "TF_IN_AUTOMATION=1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("installed terraform %s: %v\n%s", strings.Join(args, " "), err, out)
		}
		return out
	}
	run("init", "-backend=false", "-no-color")
	run("validate", "-no-color")
	var wire struct {
		Providers map[string]struct {
			Resources   map[string]json.RawMessage `json:"resource_schemas"`
			DataSources map[string]json.RawMessage `json:"data_source_schemas"`
			Actions     map[string]json.RawMessage `json:"action_schemas"`
		} `json:"provider_schemas"`
	}
	if err := json.Unmarshal(run("providers", "schema", "-json"), &wire); err != nil {
		t.Fatal(err)
	}
	p := wire.Providers["registry.terraform.io/stripe/stripe"]
	// Validate the shipped examples with a real installation and no API keys.
	examples, err := os.ReadDir("examples/metronome")
	if err != nil {
		t.Fatal(err)
	}
	for _, example := range examples {
		if !example.IsDir() {
			continue
		}
		t.Run(example.Name(), func(t *testing.T) {
			directory := filepath.Join(root, "examples", example.Name())
			err := filepath.WalkDir(filepath.Join("examples", "metronome", example.Name()), func(source string, entry os.DirEntry, walkErr error) error {
				if walkErr != nil {
					return walkErr
				}
				relative, e := filepath.Rel(filepath.Join("examples", "metronome", example.Name()), source)
				if e != nil {
					return e
				}
				target := filepath.Join(directory, relative)
				if entry.IsDir() {
					return os.MkdirAll(target, 0700)
				}
				if !strings.HasSuffix(source, ".tf") {
					return nil
				}
				content, e := os.ReadFile(source)
				if e != nil {
					return e
				}
				return os.WriteFile(target, content, 0600)
			})
			if err != nil {
				t.Fatal(err)
			}
			for _, args := range [][]string{{"init", "-backend=false", "-no-color"}, {"validate", "-no-color"}} {
				cmd := exec.Command("terraform", args...)
				cmd.Dir = directory
				cmd.Env = terraformTestEnvironment("TF_CLI_CONFIG_FILE="+cli, "TF_IN_AUTOMATION=1")
				if out, err := cmd.CombinedOutput(); err != nil {
					t.Fatalf("example %s %v: %v\n%s", example.Name(), args, err, out)
				}
			}
		})
	}
	for _, test := range []struct {
		name     string
		surfaces map[string]json.RawMessage
		expected int
	}{{"resources", p.Resources, 15}, {"data sources", p.DataSources, 59}, {"actions", p.Actions, 73}} {
		count := 0
		for name := range test.surfaces {
			if strings.HasPrefix(name, "stripe_metronome_") {
				count++
			}
		}
		if count != test.expected {
			t.Fatalf("installed %s inventory: %d, expected %d", test.name, count, test.expected)
		}
	}
}
