//
// File generated from our OpenAPI spec
//

package main

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestMetronomeTerraformCustomerContractLifecycle(t *testing.T) {
	if _, err := exec.LookPath("terraform"); err != nil {
		t.Fatal("required offline gate cannot find Terraform")
	}
	const customerID = "11111111-1111-4111-8111-111111111111"
	const contractID = "22222222-2222-4222-8222-222222222222"
	const packageID = "33333333-3333-4333-8333-333333333333"
	var mu sync.Mutex
	var customer, contract bool
	createdContracts := 0
	var aliases []string
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		var body map[string]any
		if r.Method == http.MethodPost {
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Error(err)
			}
		}
		switch r.URL.Path {
		case "/v1/customers":
			if r.Method == http.MethodGet {
				if r.URL.Query().Get("ingest_alias") != "example-enterprise-customer" {
					t.Errorf("wrong alias lookup: %s", r.URL.RawQuery)
				}
				w.Write([]byte(`{"data":[],"next_page":null}`))
				return
			}
			customer = true
			aliases = []string{"example-enterprise-customer"}
			w.Write([]byte(`{"data":{"id":"` + customerID + `"}}`))
		case "/v1/customers/" + customerID:
			if !customer {
				http.Error(w, "not found", 404)
				return
			}
			json.NewEncoder(w).Encode(map[string]any{"data": map[string]any{"id": customerID, "name": "Example enterprise", "ingest_aliases": aliases}})
		case "/v1/customers/" + customerID + "/setIngestAliases":
			if values, ok := body["ingest_aliases"].([]any); !ok || len(values) != 0 {
				t.Errorf("wrong alias clear payload: %v", body)
			}
			aliases = nil
			w.Write([]byte(`{"data":{}}`))
		case "/v1/customers/archive":
			if contract {
				t.Error("customer archived before contract")
			}
			customer = false
			w.Write([]byte(`{"data":{"id":"` + customerID + `"}}`))
		case "/v1/contracts/create":
			createdContracts++
			if body["customer_id"] != customerID || body["package_id"] != packageID {
				t.Errorf("wrong contract request: %v", body)
			}
			contract = true
			w.Write([]byte(`{"data":{"id":"` + contractID + `"}}`))
		case "/v2/contracts/get":
			if body["customer_id"] != customerID || body["contract_id"] != contractID {
				t.Errorf("wrong contract get request: %v", body)
			}
			if !contract {
				http.Error(w, "not found", 404)
				return
			}
			w.Write([]byte(`{"data":{"id":"` + contractID + `","customer_id":"` + customerID + `","package_id":"` + packageID + `","uniqueness_key":"example-enterprise-contract:v1","starting_at":"2026-10-03T00:00:00.000Z"}}`))
		case "/v2/contracts/list":
			if contract {
				w.Write([]byte(`{"data":[{"id":"` + contractID + `","customer_id":"` + customerID + `"}],"cursor":null}`))
			} else {
				w.Write([]byte(`{"data":[],"cursor":null}`))
			}
		case "/v1/contracts/archive":
			if body["void_invoices"] != false {
				t.Errorf("unexpected void_invoices: %v", body)
			}
			contract = false
			w.Write([]byte(`{"data":{"id":"` + contractID + `"}}`))
		default:
			t.Errorf("unexpected path %s", r.URL.Path)
			http.Error(w, "not found", 404)
		}
	}))
	defer server.Close()
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	temp := t.TempDir()
	cert := filepath.Join(temp, "ca.pem")
	if _, err := x509.ParseCertificate(server.Certificate().Raw); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(cert, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600); err != nil {
		t.Fatal(err)
	}
	providerDir := filepath.Join(temp, "provider")
	if err := os.Mkdir(providerDir, 0700); err != nil {
		t.Fatal(err)
	}
	build := exec.Command("go", "build", "-o", filepath.Join(providerDir, "terraform-provider-stripe"), workingDir)
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build provider: %v\n%s", err, out)
	}
	cliPath := filepath.Join(temp, "terraformrc")
	if err := os.WriteFile(cliPath, []byte(fmt.Sprintf(`provider_installation { dev_overrides { "registry.terraform.io/stripe/stripe" = %q } direct {} }`, providerDir)), 0600); err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(temp, "root")
	if err := os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	example, err := os.ReadFile(filepath.Join(workingDir, "examples", "metronome", "customer-contract", "main.tf"))
	if err != nil {
		t.Fatal(err)
	}
	config := strings.Replace(string(example), `provider "stripe" {}`, fmt.Sprintf("provider \"stripe\" {\n  metronome = {\n  base_url = %q\n  ca_cert_file = %q\n  }\n}", server.URL, cert), 1)
	config = strings.Replace(config, "deletion_protection = true", "deletion_protection = false", 1)
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	env := terraformTestEnvironment("TF_CLI_CONFIG_FILE="+cliPath, "TF_DATA_DIR="+filepath.Join(temp, "terraform-data"), "STRIPE_API_KEY=sk_test_offline", "TF_IN_AUTOMATION=1", "TF_VAR_package_id="+packageID, "TF_VAR_start_at=2026-10-03T00:00:00Z")
	run := func(args ...string) (string, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, "terraform", args...)
		cmd.Dir = root
		cmd.Env = env
		out, err := cmd.CombinedOutput()
		return string(out), err
	}
	for _, args := range [][]string{{"get", "-no-color"}, {"apply", "-auto-approve", "-no-color"}, {"plan", "-detailed-exitcode", "-no-color"}} {
		if out, err := run(args...); err != nil {
			t.Fatalf("terraform %s: %v\n%s", args[0], err, out)
		}
	}
	stateOnly := strings.Replace(config, "void_invoices_on_destroy = false", "void_invoices_on_destroy = true", 1)
	if stateOnly == config {
		t.Fatal("failed to construct state-only contract update")
	}
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(stateOnly), 0600); err != nil {
		t.Fatal(err)
	}
	if out, err := run("plan", "-no-color"); err != nil || !strings.Contains(out, "Plan: 0 to add, 1 to change, 0 to destroy") {
		t.Fatalf("state-only update would replace contract: %v\n%s", err, out)
	}
	if out, err := run("apply", "-auto-approve", "-no-color"); err != nil {
		t.Fatalf("state-only update failed: %v\n%s", err, out)
	}
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	if out, err := run("apply", "-auto-approve", "-no-color"); err != nil {
		t.Fatalf("restoring archival preference failed: %v\n%s", err, out)
	}
	invalid := strings.Replace(config, "starting_at              = var.start_at", "starting_at              = \"2026-10-03T00:30:00Z\"", 1)
	invalid = strings.Replace(invalid, "example-enterprise-contract:v1", "example-enterprise-contract:v2", 1)
	if invalid == config {
		t.Fatal("failed to construct invalid contract plan")
	}
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(invalid), 0600); err != nil {
		t.Fatal(err)
	}
	if out, err := run("plan", "-no-color"); err == nil || !strings.Contains(out, "aligned to the hour") {
		t.Fatalf("invalid replacement was not rejected at plan time: %v\n%s", err, out)
	}
	mu.Lock()
	stillActive := contract
	mu.Unlock()
	if !stillActive {
		t.Fatal("contract was archived while planning invalid replacement")
	}
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{{"state", "rm", "stripe_metronome_contract.enterprise"}, {"state", "rm", "stripe_metronome_customer.enterprise"}, {"import", "-no-color", "stripe_metronome_customer.enterprise", customerID}, {"import", "-no-color", "stripe_metronome_contract.enterprise", customerID + "/" + contractID}, {"apply", "-auto-approve", "-no-color"}, {"plan", "-detailed-exitcode", "-no-color"}, {"destroy", "-auto-approve", "-no-color"}} {
		if out, err := run(args...); err != nil {
			t.Fatalf("terraform %s: %v\n%s", args[0], err, out)
		}
	}
	mu.Lock()
	defer mu.Unlock()
	if customer || contract || createdContracts != 1 {
		t.Fatalf("unexpected lifecycle: customer=%t contract=%t contract creates=%d", customer, contract, createdContracts)
	}
}
