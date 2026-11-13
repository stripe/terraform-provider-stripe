//
// File generated from our OpenAPI spec
//

package main

import (
	"context"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestMetronomeGeneratedReadAndExplicitActionInvocation(t *testing.T) {
	terraform, err := exec.LookPath("terraform")
	if err != nil {
		if os.Getenv("METRONOME_REQUIRE_ACTION_TEST") == "1" {
			t.Fatal("required action gate cannot find Terraform")
		}
		t.Fatal("required offline gate cannot find Terraform")
	}
	versionOutput, err := exec.Command(terraform, "version", "-json").Output()
	if err != nil {
		t.Fatal(err)
	}
	var version struct {
		Version string `json:"terraform_version"`
	}
	if err := json.Unmarshal(versionOutput, &version); err != nil {
		t.Fatal(err)
	}
	parts := strings.Split(version.Version, ".")
	if len(parts) < 2 {
		t.Fatalf("unexpected Terraform version %q", version.Version)
	}
	minor, minorErr := strconv.Atoi(parts[1])
	major, majorErr := strconv.Atoi(parts[0])
	if minorErr != nil || majorErr != nil {
		t.Fatalf("unexpected Terraform version %q", version.Version)
	}
	if major == 1 && minor < 14 {
		if os.Getenv("METRONOME_REQUIRE_ACTION_TEST") == "1" {
			t.Fatal("release action gate requires Terraform 1.14+")
		}
		t.Fatal("required offline action gate needs Terraform 1.14+")
	}
	var reads, writes atomic.Int32
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer sk_test_offline" {
			t.Error("missing auth")
		}
		switch r.URL.EscapedPath() {
		case "/v1/contract-pricing/products/get":
			reads.Add(1)
			w.Write([]byte(`{"data":{"id":"product","name":"lookup"}}`))
		case "/v1/customers/acct%20123/setName":
			writes.Add(1)
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
			if body["name"] != "New name" {
				t.Errorf("bad action body %+v", body)
			}
			w.Write([]byte(`{"data":{"id":"acct 123"}}`))
		default:
			t.Errorf("unexpected request %s", r.URL.EscapedPath())
			w.WriteHeader(404)
		}
	}))
	defer server.Close()
	temp := t.TempDir()
	providerDir := filepath.Join(temp, "provider")
	if err := os.Mkdir(providerDir, 0700); err != nil {
		t.Fatal(err)
	}
	build := exec.Command("go", "build", "-o", filepath.Join(providerDir, "terraform-provider-stripe"), ".")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, out)
	}
	certPath := filepath.Join(temp, "cert.pem")
	if err := os.WriteFile(certPath, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600); err != nil {
		t.Fatal(err)
	}
	cliPath := filepath.Join(temp, "terraformrc")
	cli := fmt.Sprintf(`provider_installation { dev_overrides { "registry.terraform.io/stripe/stripe" = %q } direct {} }`, providerDir)
	if err := os.WriteFile(cliPath, []byte(cli), 0600); err != nil {
		t.Fatal(err)
	}
	hcl := fmt.Sprintf(`
terraform {
 required_providers {
  stripe = { source = "stripe/stripe" }
 }
}
provider "stripe" {
 metronome = {
 base_url = %q
 ca_cert_file = %q
 }
}
data "stripe_metronome_get_product_v1" "lookup" { request = { id = "product" } }
output "result" {
 value = data.stripe_metronome_get_product_v1.lookup.response_json
 sensitive = true
}
action "stripe_metronome_set_customer_name_v1" "rename" {
 config {
  path_params = { customer_id = "acct 123" }
  request = { name = "New name" }
 }
}
`, server.URL, certPath)
	if err := os.WriteFile(filepath.Join(temp, "main.tf"), []byte(hcl), 0600); err != nil {
		t.Fatal(err)
	}
	run := func(args ...string) string {
		ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, terraform, args...)
		cmd.Dir = temp
		cmd.Env = terraformTestEnvironment("TF_CLI_CONFIG_FILE="+cliPath, "STRIPE_API_KEY=sk_test_offline", "TF_IN_AUTOMATION=1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("terraform %v: %v\n%s", args, err, out)
		}
		return string(out)
	}
	run("apply", "-auto-approve", "-no-color")
	if reads.Load() == 0 || writes.Load() != 0 {
		t.Fatalf("action ran implicitly: %d reads, %d writes", reads.Load(), writes.Load())
	}
	run("apply", "-invoke=action.stripe_metronome_set_customer_name_v1.rename", "-auto-approve", "-no-color")
	if writes.Load() != 1 {
		t.Fatalf("explicit action invocation count %d", writes.Load())
	}
}
