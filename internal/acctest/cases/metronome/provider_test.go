//
// File generated from our OpenAPI spec
//

package metronome_test

import (
	"context"
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

	runner "github.com/stripe/terraform-provider-stripe/internal/acctest/runner/metronome"
)

func TestMetronomeModuleWithAutomaticProvisioning(t *testing.T) {
	runner.RequireTerraform(t)
	root := t.TempDir()
	binaryDir := t.TempDir()
	build := exec.Command("go", "build", "-o", filepath.Join(binaryDir, "terraform-provider-stripe"), ".")
	build.Dir = runner.ProviderRoot(t)
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("provider build: %v\n%s", err, out)
	}
	var mu sync.Mutex
	provisioned, created := false, false
	reads, mutations := 0, 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if r.Header.Get("Authorization") != "Bearer sk_test_shared_account" || r.Header.Get("Stripe-Account") != "" {
			t.Error("module did not use the shared Stripe key")
			http.Error(w, "unauthorized", 401)
			return
		}
		// The API provisions the account on its first authenticated request.
		provisioned = true
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v1/credit-types/list":
			if r.Method != http.MethodGet {
				t.Errorf("unexpected list request: %s %s", r.Method, r.URL)
			}
			reads++
			if created {
				_, _ = io.WriteString(w, `{"data":[{"id":"unit_module","name":"module credits","is_currency":false}]}`)
			} else {
				_, _ = io.WriteString(w, `{"data":[]}`)
			}
		case "/v1/credit-types/create":
			if r.Method != http.MethodPost {
				t.Error("create method changed")
			}
			created = true
			mutations++
			_, _ = io.WriteString(w, `{"data":{"id":"unit_module"}}`)
		case "/v1/credit-types/archive":
			if r.Method != http.MethodPost {
				t.Error("archive method changed")
			}
			created = false
			mutations++
			_, _ = io.WriteString(w, `{"data":{"id":"unit_module"}}`)
		default:
			t.Errorf("unexpected API request: %s %s", r.Method, r.URL)
			http.Error(w, "unexpected", 400)
		}
	}))
	defer server.Close()
	ca := filepath.Join(root, "ca.pem")
	if err := os.WriteFile(ca, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600); err != nil {
		t.Fatal(err)
	}
	cli := filepath.Join(root, "terraform.rc")
	if err := os.WriteFile(cli, []byte(fmt.Sprintf("provider_installation {\n dev_overrides {\n \"stripe/stripe\" = %q\n }\n direct {}\n}\n", binaryDir)), 0600); err != nil {
		t.Fatal(err)
	}
	env := runner.Environment("TF_CLI_CONFIG_FILE="+cli, "TF_IN_AUTOMATION=1", "STRIPE_API_KEY=sk_test_shared_account")
	config := runner.RenderConfig(t, "shared_account/main.tf", map[string]string{"API_URL": server.URL, "CA_CERT_FILE": ca})
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "billing"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "billing", "main.tf"), []byte(runner.RenderConfig(t, "shared_account/billing/main.tf", nil)), 0600); err != nil {
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
	planned := provisioned && reads > 0 && mutations == 0
	mu.Unlock()
	if !planned {
		t.Fatal("initial plan did not read Metronome directly without resource mutations")
	}
	run("apply", "-no-color", "initial.plan")
	mu.Lock()
	applied := created && mutations == 1
	mu.Unlock()
	if !applied {
		t.Fatal("module did not create its resource directly")
	}
	run("plan", "-no-color", "-detailed-exitcode")
	run("destroy", "-no-color", "-auto-approve")
	mu.Lock()
	defer mu.Unlock()
	if created || mutations != 2 {
		t.Fatalf("module lifecycle had unexpected writes: created=%t mutations=%d", created, mutations)
	}
}

func TestMetronomeInstalledProviderWithoutCredentials(t *testing.T) {
	runner.RequireTerraform(t)
	root := t.TempDir()
	mirror := filepath.Join(root, "mirror")
	version := "99.0.0"
	install := filepath.Join(mirror, "registry.terraform.io", "stripe", "stripe", version, runtime.GOOS+"_"+runtime.GOARCH)
	if err := os.MkdirAll(install, 0700); err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(install, "terraform-provider-stripe_v"+version)
	build := exec.Command("go", "build", "-ldflags", "-X main.version="+version, "-o", binary, ".")
	build.Dir = runner.ProviderRoot(t)
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("provider build: %v\n%s", err, out)
	}
	config := runner.RenderConfig(t, "installed_provider.tf", map[string]string{"VERSION": version})
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
		cmd.Env = runner.Environment("TF_CLI_CONFIG_FILE="+cli, "TF_IN_AUTOMATION=1")
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
	examples, err := os.ReadDir(filepath.Join(runner.ProviderRoot(t), "examples", "metronome"))
	if err != nil {
		t.Fatal(err)
	}
	for _, example := range examples {
		if !example.IsDir() {
			continue
		}
		t.Run(example.Name(), func(t *testing.T) {
			directory := filepath.Join(root, "examples", example.Name())
			err := filepath.WalkDir(filepath.Join(runner.ProviderRoot(t), "examples", "metronome", example.Name()), func(source string, entry os.DirEntry, walkErr error) error {
				if walkErr != nil {
					return walkErr
				}
				relative, e := filepath.Rel(filepath.Join(runner.ProviderRoot(t), "examples", "metronome", example.Name()), source)
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
				cmd.Env = runner.Environment("TF_CLI_CONFIG_FILE="+cli, "TF_IN_AUTOMATION=1")
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
	}{{"resources", p.Resources, 14}, {"data sources", p.DataSources, 57}, {"actions", p.Actions, 69}} {
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
