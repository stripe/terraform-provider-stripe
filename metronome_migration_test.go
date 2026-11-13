//
// File generated from our OpenAPI spec
//

package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

type migrationCase struct {
	Name, Kind, HCL string
	ImportID        string `json:"import_id"`
	Module          bool
	Responses       []struct{ Method, Path, Response string }
}

func writeMigrationConfig(t *testing.T, root string, test migrationCase, legacy bool, origin, ca string) string {
	t.Helper()
	providerName, source, prefix := "stripe", "stripe/stripe", "stripe_metronome_"
	if legacy {
		providerName, source, prefix = "metronome", "tonyxiao/metronome", "metronome_"
	}
	required := fmt.Sprintf("terraform {\n required_providers {\n %s = { source = %q }\n }\n}\n", providerName, source)
	transport := fmt.Sprintf("base_url = %q\n ca_cert_file = %q", origin, ca)
	if !legacy {
		transport = "metronome = {\n " + transport + "\n }"
	}
	config := required + fmt.Sprintf("provider %q {\n alias = \"migration\"\n %s\n}\n", providerName, transport)
	address := prefix + test.Kind + `.fixture["instance"]`
	resourceConfig := fmt.Sprintf("resource %q \"fixture\" {\n for_each = toset([\"instance\"])\n %s\n", prefix+test.Kind, test.HCL)
	if test.Module {
		child := filepath.Join(root, "billing")
		if err := os.MkdirAll(child, 0700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(child, "main.tf"), []byte(required+resourceConfig+"}\n"), 0600); err != nil {
			t.Fatal(err)
		}
		config += fmt.Sprintf("module \"billing\" {\n source = \"./billing\"\n providers = { %s = %s.migration }\n}\n", providerName, providerName)
		address = "module.billing." + address
	} else {
		config += resourceConfig + fmt.Sprintf(" provider = %s.migration\n}\n", providerName)
	}
	if !legacy {
		oldAddress := strings.Replace(address, "stripe_metronome_", "metronome_", 1)
		config += fmt.Sprintf("moved {\n from = %s\n to = %s\n}\n", oldAddress, address)
	}
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	return address
}

// The capture mode produces fixtures with the pinned standalone executable and
// Terraform 1.5.7. Ordinary tests consume only the checked-in, token-free states.
func TestMetronomeCrossProviderMigration(t *testing.T) {
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	fixtureRoot := filepath.Join(workingDir, "testdata", "metronome", "migration")
	content, err := os.ReadFile(filepath.Join(fixtureRoot, "cases.json"))
	if err != nil {
		t.Fatal(err)
	}
	var cases []migrationCase
	if err = json.Unmarshal(content, &cases); err != nil {
		t.Fatal(err)
	}
	kinds := map[string]bool{}
	for _, test := range cases {
		kinds[test.Kind] = true
	}
	if len(kinds) != 14 || len(cases) != 18 {
		t.Fatal("migration fixture inventory is incomplete")
	}
	captureDir := os.Getenv("METRONOME_FIXTURE_CAPTURE_DIR")
	capture := captureDir != ""
	terraform := "terraform"
	binarySource := os.Getenv("METRONOME_LEGACY_BINARY")
	if capture {
		terraform = os.Getenv("METRONOME_LEGACY_TERRAFORM")
		if terraform == "" || binarySource == "" {
			t.Fatal("fixture capture requires the pinned legacy Terraform and provider binaries")
		}
	}
	versionJSON, err := exec.Command(terraform, "version", "-json").Output()
	if err != nil {
		t.Fatal("required migration Terraform is unavailable", err)
	}
	var version struct {
		Version string `json:"terraform_version"`
	}
	if err = json.Unmarshal(versionJSON, &version); err != nil {
		t.Fatal(err)
	}
	if capture && version.Version != "1.5.7" {
		t.Fatal("legacy fixtures must be produced with Terraform 1.5.7")
	}
	if !capture && version.Version != "1.16.5" {
		t.Fatalf("required migration gate uses Terraform 1.16.5, got %s", version.Version)
	}
	binaryDir := t.TempDir()
	binaryName := "terraform-provider-stripe"
	if capture {
		binaryName = "terraform-provider-metronome"
	}
	binary := filepath.Join(binaryDir, binaryName)
	if capture {
		raw, err := os.ReadFile(binarySource)
		if err != nil {
			t.Fatal(err)
		}
		if err = os.WriteFile(binary, raw, 0700); err != nil {
			t.Fatal(err)
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer cancel()
		build := exec.CommandContext(ctx, "go", "build", "-o", binary, ".")
		if out, err := build.CombinedOutput(); err != nil {
			t.Fatalf("provider build: %v\n%s", err, out)
		}
	}
	legacyDir := binaryDir
	if !capture {
		legacyDir = t.TempDir()
		legacyBinary := filepath.Join(legacyDir, "terraform-provider-metronome")
		if binarySource != "" {
			raw, err := os.ReadFile(binarySource)
			if err != nil {
				t.Fatal(err)
			}
			if err = os.WriteFile(legacyBinary, raw, 0700); err != nil {
				t.Fatal(err)
			}
		} else {
			build := exec.Command("go", "build", "-o", legacyBinary, "./testdata/metronome/legacy_provider")
			if out, err := build.CombinedOutput(); err != nil {
				t.Fatalf("legacy schema replay build: %v\n%s", err, out)
			}
		}
	}
	manifest := map[string]string{}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			root := t.TempDir()
			var mutations atomic.Int32
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				expected := "Bearer sk_test_migration"
				if capture {
					expected = "Bearer legacy-fixture-token"
				}
				if r.Header.Get("Authorization") != expected {
					t.Error("migration used the wrong provider credentials")
					http.Error(w, "forbidden", 403)
					return
				}
				if r.URL.Path == "/v1/alerts/create" {
					mutations.Add(1)
					if !capture {
						t.Error("migration recreated a non-importable alert")
						http.Error(w, "unexpected mutation", 400)
						return
					}
				}
				for _, response := range test.Responses {
					if r.Method == response.Method && r.URL.Path == response.Path {
						w.Header().Set("Content-Type", "application/json")
						_, _ = w.Write([]byte(response.Response))
						return
					}
				}
				t.Errorf("unexpected migration API request: %s %s", r.Method, r.URL.Path)
				http.Error(w, "unexpected", 400)
			}))
			defer server.Close()
			ca := filepath.Join(root, "ca.pem")
			if err = os.WriteFile(ca, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600); err != nil {
				t.Fatal(err)
			}
			source := "stripe/stripe"
			if capture {
				source = "tonyxiao/metronome"
			}
			overrides := fmt.Sprintf("%q = %q\n", source, binaryDir)
			if !capture {
				overrides += fmt.Sprintf("%q = %q\n", "tonyxiao/metronome", legacyDir)
			}
			cli := filepath.Join(root, "terraform.rc")
			if err = os.WriteFile(cli, []byte("provider_installation {\n dev_overrides {\n"+overrides+"}\n direct {}\n}\n"), 0600); err != nil {
				t.Fatal(err)
			}
			env := terraformTestEnvironment("TF_CLI_CONFIG_FILE="+cli, "TF_IN_AUTOMATION=1", "STRIPE_API_KEY=sk_test_migration", "METRONOME_API_TOKEN=legacy-fixture-token")
			run := func(args ...string) (string, error) {
				ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
				defer cancel()
				cmd := exec.CommandContext(ctx, terraform, args...)
				cmd.Dir = root
				cmd.Env = env
				out, err := cmd.CombinedOutput()
				return string(out), err
			}
			address := writeMigrationConfig(t, root, test, capture, server.URL, ca)
			if out, err := run("get", "-no-color"); err != nil {
				t.Fatalf("terraform get: %v\n%s", err, out)
			}
			if capture {
				if test.ImportID == "" {
					if out, err := run("apply", "-auto-approve", "-no-color"); err != nil {
						t.Fatalf("legacy apply: %v\n%s", err, out)
					}
				} else if out, err := run("import", "-no-color", address, test.ImportID); err != nil {
					t.Fatalf("legacy import: %v\n%s", err, out)
				}
				if out, err := run("plan", "-detailed-exitcode", "-no-color"); err != nil {
					t.Fatalf("legacy fixture has drift: %v\n%s", err, out)
				}
				state, err := os.ReadFile(filepath.Join(root, "terraform.tfstate"))
				if err != nil {
					t.Fatal(err)
				}
				if bytes.Contains(state, []byte("legacy-fixture-token")) {
					t.Fatal("fixture contains a credential")
				}
				if err = os.MkdirAll(captureDir, 0700); err != nil {
					t.Fatal(err)
				}
				if err = os.WriteFile(filepath.Join(captureDir, test.Name+".json"), state, 0600); err != nil {
					t.Fatal(err)
				}
				sum := sha256.Sum256(state)
				manifest[test.Name+".json"] = hex.EncodeToString(sum[:])
				return
			}
			legacyState, err := os.ReadFile(filepath.Join(fixtureRoot, test.Name+".json"))
			if err != nil {
				t.Fatal(err)
			}
			if err = os.WriteFile(filepath.Join(root, "terraform.tfstate"), legacyState, 0600); err != nil {
				t.Fatal(err)
			}
			if out, err := run("plan", "-out=migrate.tfplan", "-no-color"); err != nil {
				t.Fatalf("cross-provider move plan: %v\n%s", err, out)
			}
			planJSON, err := run("show", "-json", "migrate.tfplan")
			if err != nil {
				t.Fatalf("read move plan: %v\n%s", err, planJSON)
			}
			var plan struct {
				Changes []struct {
					Address         string
					PreviousAddress string `json:"previous_address"`
					Change          struct {
						Actions       []string
						Before, After map[string]json.RawMessage
					}
				} `json:"resource_changes"`
			}
			if err = json.Unmarshal([]byte(planJSON), &plan); err != nil {
				t.Fatal(err)
			}
			if len(plan.Changes) != 1 {
				t.Fatalf("expected one moved resource: %s", planJSON)
			}
			change := plan.Changes[0]
			if change.Address != address || change.PreviousAddress != strings.Replace(address, "stripe_metronome_", "metronome_", 1) || !reflect.DeepEqual(change.Change.Actions, []string{"no-op"}) || !bytes.Equal(change.Change.Before["id"], change.Change.After["id"]) {
				t.Fatalf("move changed identity or requested an API update: %+v", change)
			}
			if out, err := run("apply", "-auto-approve", "-no-color", "migrate.tfplan"); err != nil {
				t.Fatalf("move apply: %v\n%s", err, out)
			}
			if out, err := run("plan", "-detailed-exitcode", "-no-color"); err != nil {
				t.Fatalf("migrated state has drift: %v\n%s", err, out)
			}
			migrated, err := os.ReadFile(filepath.Join(root, "terraform.tfstate"))
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Contains(migrated, []byte(`registry.terraform.io/stripe/stripe`)) || bytes.Contains(migrated, []byte(`registry.terraform.io/tonyxiao/metronome`)) {
				t.Fatal("provider address was not migrated")
			}
			if mutations.Load() != 0 {
				t.Fatal("state conversion mutated Metronome")
			}
		})
	}
	if capture {
		if len(manifest) != len(cases) {
			t.Fatal("legacy fixture capture was incomplete")
		}
		proof := map[string]any{"sourceRevision": "4abb7576cbf86763322a5aa1c33e89b4ee6c67e6", "terraformVersion": "1.5.7", "method": "Terraform import or alert apply, followed by a zero-diff plan against local HTTPS fixtures", "files": manifest}
		bytes, err := json.MarshalIndent(proof, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
		if err = os.WriteFile(filepath.Join(captureDir, "provenance.json"), append(bytes, '\n'), 0600); err != nil {
			t.Fatal(err)
		}
	}
}
