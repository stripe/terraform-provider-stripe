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

// Exercises the explicit examples against a local HTTPS Metronome stand-in.
// It never touches the shared sandbox account.
func TestMetronomeExplicitResourcesLifecycle(t *testing.T) {
	for _, tc := range []struct {
		name           string
		rateCard       string
		proKey         string
		requestJSON    bool
		typedCreate    bool
		typedMigration bool
	}{
		{name: "subscription-only", rateCard: "stripe_metronome_rate_card.subscription"},
		{name: "composite-percentage", rateCard: "stripe_metronome_rate_card.usd"},
		{name: "subscription-with-overage", rateCard: "stripe_metronome_rate_card.usd", proKey: "example:overage:pro:v1"},
		{name: "openrouter-cost-markup", rateCard: "stripe_metronome_rate_card.openrouter_markup", proKey: "openrouter:pro:markup-20:v1"},
		{name: "api-mirror-mixed", rateCard: "stripe_metronome_rate_card.mixed", proKey: "openrouter:mixed:markup-30:v1"},
		{name: "api-mirror-mixed", rateCard: "stripe_metronome_rate_card.mixed", proKey: "openrouter:mixed:markup-30:v1", requestJSON: true},
		{name: "api-mirror-mixed", rateCard: "stripe_metronome_rate_card.mixed", proKey: "openrouter:mixed:markup-30:v1", typedCreate: true},
		{name: "api-mirror-mixed", rateCard: "stripe_metronome_rate_card.mixed", proKey: "openrouter:mixed:markup-30:v1", typedMigration: true},
	} {
		name := tc.name
		if tc.requestJSON {
			name += "-json"
		}
		if tc.typedCreate {
			name += "-typed"
		}
		if tc.typedMigration {
			name += "-migrate-typed"
		}
		t.Run(name, func(t *testing.T) {
			testExplicitResourcesLifecycle(t, tc.name, tc.rateCard, tc.proKey, tc.requestJSON, tc.typedCreate, tc.typedMigration)
		})
	}
}

func testExplicitResourcesLifecycle(t *testing.T, exampleName, rateCardAddress, expectedProKey string, requestJSON, typedCreate, typedMigration bool) {
	if _, err := exec.LookPath("terraform"); err != nil {
		t.Fatal("required offline gate cannot find Terraform")
	}
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	temp := t.TempDir()
	var mu sync.Mutex
	var metric map[string]any
	products := map[string]map[string]any{}
	var card map[string]any
	var rate map[string]any
	var allRates []map[string]any
	packages := map[string]map[string]any{}
	archived := 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if r.Header.Get("Authorization") != "Bearer sk_test_offline" {
			http.Error(w, "unauthorized", 401)
			return
		}
		var body map[string]any
		if r.Method == http.MethodPost {
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("decode %s: %v", r.URL.Path, err)
				http.Error(w, "bad request", 400)
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		var response any
		switch r.URL.Path {
		case "/v1/listConfiguredBillingProviders":
			response = map[string]any{"data": []any{}}
		case "/v1/billable-metrics/create":
			metric = body
			metric["id"] = "metric-id"
			metric["aggregation_type"] = strings.ToLower(body["aggregation_type"].(string))
			response = map[string]any{"data": map[string]any{"id": "metric-id"}}
		case "/v1/billable-metrics/metric-id":
			response = map[string]any{"data": metric}
		case "/v1/billable-metrics/archive":
			archived++
			response = map[string]any{"data": map[string]any{"id": "metric-id"}}
		case "/v1/contract-pricing/products/create":
			id := "fixed-id"
			if body["type"] == "USAGE" {
				id = "usage-id"
			} else if body["type"] == "SUBSCRIPTION" {
				id = "subscription-id"
			} else if body["type"] == "COMPOSITE" {
				id = "composite-id"
			}
			products[id] = map[string]any{"id": id, "type": body["type"], "current": map[string]any{"name": body["name"], "billable_metric_id": body["billable_metric_id"], "composite_product_ids": body["composite_product_ids"], "composite_tags": body["composite_tags"]}}
			response = map[string]any{"data": map[string]any{"id": id}}
		case "/v1/contract-pricing/products/update":
			product := products[body["product_id"].(string)]
			product["current"].(map[string]any)["name"] = body["name"]
			response = map[string]any{"data": map[string]any{"id": body["product_id"]}}
		case "/v1/contract-pricing/products/get":
			response = map[string]any{"data": products[body["id"].(string)]}
		case "/v1/contract-pricing/products/archive":
			archived++
			response = map[string]any{"data": map[string]any{"id": body["product_id"]}}
		case "/v1/contract-pricing/rate-cards/create":
			creditTypeID := body["fiat_credit_type_id"]
			if creditTypeID == nil {
				creditTypeID = "usd-cents"
			}
			card = map[string]any{"id": "card-id", "name": body["name"], "description": body["description"], "fiat_credit_type": map[string]any{"id": creditTypeID}}
			response = map[string]any{"data": map[string]any{"id": "card-id"}}
		case "/v1/contract-pricing/rate-cards/addRate":
			creditTypeID := body["credit_type_id"]
			if creditTypeID == nil {
				creditTypeID = "usd-cents"
			}
			rate = map[string]any{"product_id": body["product_id"], "starting_at": body["starting_at"], "entitled": body["entitled"], "billing_frequency": body["billing_frequency"], "rate": map[string]any{"rate_type": body["rate_type"], "price": body["price"], "credit_type": map[string]any{"id": creditTypeID}}}
			allRates = append(allRates, rate)
			response = map[string]any{"data": map[string]any{"rate_type": "FLAT"}}
		case "/v1/contract-pricing/rate-cards/update":
			if body["rate_card_id"] != "card-id" {
				t.Error("update lost card ID")
				http.Error(w, "bad ID", 400)
				return
			}
			card["name"] = body["name"]
			if description, ok := body["description"]; ok {
				card["description"] = description
			}
			response = map[string]any{"data": map[string]string{"id": "card-id"}}
		case "/v1/contract-pricing/rate-cards/get":
			response = map[string]any{"data": card}
		case "/v1/contract-pricing/rate-cards/getRateSchedule":
			earlierRate := map[string]any{"product_id": rate["product_id"], "starting_at": "2025-01-01T00:00:00.000Z", "entitled": true, "rate": map[string]any{"rate_type": "FLAT", "price": float64(100), "credit_type": map[string]any{"id": "usd-cents"}}}
			data := []any{earlierRate}
			for _, item := range allRates {
				data = append(data, item)
			}
			response = map[string]any{"data": data}
		case "/v1/contract-pricing/rate-cards/archive":
			archived++
			response = map[string]any{"data": map[string]any{"id": "card-id"}}
		case "/v1/packages/create":
			id := "base-id"
			if body["recurring_commits"] != nil {
				id = "pro-id"
			}
			body["id"] = id
			key := body["uniqueness_key"].(string)
			if !strings.HasPrefix(key, "package:") {
				body["uniqueness_key"] = "package:" + key
			}
			packages[id] = body
			response = map[string]any{"data": map[string]any{"id": id}}
		case "/v1/packages/get":
			id := body["package_id"].(string)
			p := packages[id]
			if p["recurring_commits"] == nil {
				p["recurring_commits"] = []any{}
			}
			encoded, _ := json.Marshal(p)
			var current map[string]any
			_ = json.Unmarshal(encoded, &current)
			current["multiplier_override_prioritization"] = "LOWEST_MULTIPLIER"
			for _, section := range []string{"recurring_commits", "recurring_credits"} {
				items, _ := current[section].([]any)
				for _, item := range items {
					term := item.(map[string]any)
					if id, ok := term["product_id"]; ok {
						term["product"] = map[string]any{"id": id, "name": "Balance"}
						delete(term, "product_id")
					}
					if amount, ok := term["access_amount"].(map[string]any); ok {
						amount["access_type"] = "SPEND"
					}
					if offset, ok := term["starting_at_offset"].(map[string]any); ok {
						offset["value"] = fmt.Sprint(offset["value"])
					}
				}
			}
			subscriptions, _ := current["subscriptions"].([]any)
			for _, item := range subscriptions {
				rate := item.(map[string]any)["subscription_rate"].(map[string]any)
				if id, ok := rate["product_id"]; ok {
					rate["product"] = map[string]any{"id": id, "name": "Seat"}
					delete(rate, "product_id")
				}
			}
			response = map[string]any{"data": current}
		case "/v1/packages/archive":
			archived++
			response = map[string]any{"data": map[string]any{"id": body["package_id"]}}
		default:
			t.Errorf("unexpected API request: %s %s", r.Method, r.URL.Path)
			http.Error(w, "unknown", 404)
			return
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	certPath := filepath.Join(temp, "cert.pem")
	cert := server.Certificate()
	if cert == nil {
		t.Fatal("missing TLS certificate")
	}
	certFile := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
	if _, err := x509.ParseCertificate(cert.Raw); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(certPath, certFile, 0600); err != nil {
		t.Fatal(err)
	}
	providerDir := filepath.Join(temp, "provider")
	if err := os.Mkdir(providerDir, 0700); err != nil {
		t.Fatal(err)
	}
	providerBinary := filepath.Join(providerDir, "terraform-provider-stripe")
	build := exec.Command("go", "build", "-o", providerBinary, workingDir)
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build provider: %v\n%s", err, output)
	}
	config := fmt.Sprintf(`provider_installation { dev_overrides { "registry.terraform.io/stripe/stripe" = %q } direct {} }`, providerDir)
	cliPath := filepath.Join(temp, "terraformrc")
	if err := os.WriteFile(cliPath, []byte(config), 0600); err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(temp, "root")
	if err := os.Mkdir(root, 0700); err != nil {
		t.Fatal(err)
	}
	example, err := os.ReadFile(filepath.Join(workingDir, "examples", "metronome", exampleName, "main.tf"))
	if err != nil {
		t.Fatal(err)
	}
	providerConfig := fmt.Sprintf("provider \"stripe\" {\n  metronome = {\n  base_url = %q\n  ca_cert_file = %q\n  }\n}", server.URL, certPath)
	hcl := strings.Replace(string(example), "provider \"stripe\" {}", providerConfig, 1)
	if hcl == string(example) {
		t.Fatal("example provider block was not found")
	}
	if requestJSON {
		marker := "resource \"stripe_metronome_package\" \"mixed\" {\n  request = {"
		at := strings.Index(hcl, marker)
		if at < 0 {
			t.Fatal("mixed request marker missing")
		}
		hcl = strings.Replace(hcl, marker, "resource \"stripe_metronome_package\" \"mixed\" {\n  request_json = jsonencode({", 1)
		start := strings.Index(hcl[at:], "jsonencode({") + at + len("jsonencode(")
		depth := 0
		for i := start; i < len(hcl); i++ {
			if hcl[i] == '{' {
				depth++
			}
			if hcl[i] == '}' {
				depth--
				if depth == 0 {
					hcl = hcl[:i+1] + ")" + hcl[i+1:]
					break
				}
			}
		}
	}
	const mixedRequest = "resource \"stripe_metronome_package\" \"mixed\" {\n  request = {"
	const mixedTyped = "resource \"stripe_metronome_package\" \"mixed\" {\n  package = {"
	if typedCreate {
		hcl = strings.Replace(hcl, mixedRequest, mixedTyped, 1)
		if !strings.Contains(hcl, mixedTyped) {
			t.Fatal("typed package marker missing")
		}
	}
	cardMarker := fmt.Sprintf("resource %q %q {\n", "stripe_metronome_rate_card", strings.TrimPrefix(rateCardAddress, "stripe_metronome_rate_card."))
	hcl = strings.Replace(hcl, cardMarker, cardMarker+"  description = \"Acceptance description\"\n", 1)
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(hcl), 0600); err != nil {
		t.Fatal(err)
	}
	env := terraformTestEnvironment("TF_CLI_CONFIG_FILE="+cliPath, "STRIPE_API_KEY=sk_test_offline", "SSL_CERT_FILE="+certPath, "TF_IN_AUTOMATION=1", "TF_VAR_credit_type_id=usd-cents", "TF_VAR_rate_start_at=2026-01-01T00:00:00.000Z", "TF_VAR_rate_start=2026-01-01T00:00:00Z")
	run := func(args ...string) (string, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, "terraform", args...)
		cmd.Dir = root
		cmd.Env = env
		out, err := cmd.CombinedOutput()
		return string(out), err
	}
	if out, err := run("get", "-no-color"); err != nil {
		t.Fatalf("terraform get: %v\n%s", err, out)
	}
	creationWindowStart := time.Now().UTC().Truncate(time.Hour)
	if out, err := run("apply", "-auto-approve", "-no-color"); err != nil {
		t.Fatalf("terraform apply: %v\n%s", err, out)
	}
	creationWindowEnd := time.Now().UTC().Truncate(time.Hour)
	if out, err := run("plan", "-detailed-exitcode", "-no-color"); err != nil {
		t.Fatalf("terraform refresh plan: %v\n%s", err, out)
	}
	if typedMigration {
		hcl = strings.Replace(hcl, mixedRequest, mixedTyped, 1)
		if !strings.Contains(hcl, mixedTyped) {
			t.Fatal("typed migration marker missing")
		}
		if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(hcl), 0600); err != nil {
			t.Fatal(err)
		}
		if out, err := run("plan", "-out=typed-migration.tfplan", "-no-color"); err != nil {
			t.Fatalf("typed migration plan: %v\n%s", err, out)
		}
		planJSON, err := run("show", "-json", "typed-migration.tfplan")
		if err != nil {
			t.Fatalf("typed migration plan JSON: %v\n%s", err, planJSON)
		}
		var planned struct {
			ResourceChanges []struct {
				Address string `json:"address"`
				Change  struct {
					Actions      []string       `json:"actions"`
					After        map[string]any `json:"after"`
					AfterUnknown map[string]any `json:"after_unknown"`
				} `json:"change"`
			} `json:"resource_changes"`
		}
		if err := json.Unmarshal([]byte(planJSON), &planned); err != nil {
			t.Fatal(err)
		}
		seen := false
		for _, change := range planned.ResourceChanges {
			if change.Address != "stripe_metronome_package.mixed" {
				continue
			}
			seen = true
			if len(change.Change.Actions) != 1 || change.Change.Actions[0] != "update" || change.Change.After["id"] != "pro-id" || change.Change.AfterUnknown["id"] == true {
				t.Fatalf("typed migration changed package identity: %+v", change.Change)
			}
		}
		if !seen {
			t.Fatal("typed migration did not plan a state update")
		}
		if out, err := run("apply", "-auto-approve", "-no-color", "typed-migration.tfplan"); err != nil {
			t.Fatalf("typed migration apply: %v\n%s", err, out)
		}
		if out, err := run("plan", "-detailed-exitcode", "-no-color"); err != nil {
			t.Fatalf("typed migration refresh plan: %v\n%s", err, out)
		}
	}

	// Reconstruct every resource from a fresh import, then require an empty plan.
	// API-shaped imports carry the declared shape instead of guessing from defaults.
	snapshot, err := run("show", "-json")
	if err != nil {
		t.Fatalf("read state for imports: %v\n%s", err, snapshot)
	}
	var state struct {
		Values struct {
			RootModule struct {
				Resources []struct {
					Address string                     `json:"address"`
					Type    string                     `json:"type"`
					Values  map[string]json.RawMessage `json:"values"`
				} `json:"resources"`
			} `json:"root_module"`
		} `json:"values"`
	}
	if err := json.Unmarshal([]byte(snapshot), &state); err != nil {
		t.Fatalf("decode state: %v\n%s", err, snapshot)
	}
	managedCount := 0
	for _, managed := range state.Values.RootModule.Resources {
		if strings.HasPrefix(managed.Address, "data.") {
			continue
		}
		managedCount++
		var id string
		if err := json.Unmarshal(managed.Values["id"], &id); err != nil {
			t.Fatal(err)
		}
		importID := id
		if managed.Type == "stripe_metronome_package" && len(managed.Values["request"]) > 0 && string(managed.Values["request"]) != "null" {
			identifier, err := json.Marshal(map[string]any{"id": id, "request": managed.Values["request"]})
			if err != nil {
				t.Fatal(err)
			}
			importID = string(identifier)
		}
		if managed.Type == "stripe_metronome_package" && len(managed.Values["package"]) > 0 && string(managed.Values["package"]) != "null" {
			identifier, err := json.Marshal(map[string]any{"id": id, "package": managed.Values["package"]})
			if err != nil {
				t.Fatal(err)
			}
			importID = string(identifier)
		}
		if managed.Type == "stripe_metronome_package" && len(managed.Values["request_json"]) > 0 && string(managed.Values["request_json"]) != "null" {
			identifier, err := json.Marshal(map[string]any{"id": id, "request_json": managed.Values["request_json"]})
			if err != nil {
				t.Fatal(err)
			}
			importID = string(identifier)
		}
		if managed.Type == "stripe_metronome_rate_card" && len(managed.Values["rates"]) > 0 && string(managed.Values["rates"]) != "null" {
			identifier, err := json.Marshal(map[string]any{"id": id, "rates": managed.Values["rates"]})
			if err != nil {
				t.Fatal(err)
			}
			importID = string(identifier)
		}
		if out, err := run("state", "rm", managed.Address); err != nil {
			t.Fatalf("remove state before import: %v\n%s", err, out)
		}
		if out, err := run("import", "-no-color", managed.Address, importID); err != nil {
			t.Fatalf("import %s: %v\n%s", managed.Address, err, out)
		}
	}
	if out, err := run("plan", "-detailed-exitcode", "-no-color"); err != nil {
		t.Fatalf("plan after all imports: %v\n%s", err, out)
	}
	// A metadata-only config can import a card whose rates are owned by actions.
	metadataRoot := filepath.Join(temp, "metadata-root")
	if err := os.Mkdir(metadataRoot, 0700); err != nil {
		t.Fatal(err)
	}
	mu.Lock()
	metadataName := card["name"].(string)
	mu.Unlock()
	metadataHCL := fmt.Sprintf("terraform {\n required_providers {\n stripe = { source = \"stripe/stripe\" }\n }\n}\n%s\nresource \"stripe_metronome_rate_card\" \"metadata\" {\n name = %q\n description = \"Acceptance description\"\n}\n", providerConfig, metadataName)
	if err := os.WriteFile(filepath.Join(metadataRoot, "main.tf"), []byte(metadataHCL), 0600); err != nil {
		t.Fatal(err)
	}
	previousRoot := root
	root = metadataRoot
	if out, err := run("import", "-no-color", "stripe_metronome_rate_card.metadata", `{"id":"card-id","rates":null}`); err != nil {
		t.Fatalf("metadata-only import: %v\n%s", err, out)
	}
	if out, err := run("plan", "-detailed-exitcode", "-no-color"); err != nil {
		t.Fatalf("metadata-only import plan: %v\n%s", err, out)
	}
	root = previousRoot
	// Metadata edits must keep identities and all computed fields, not replace cards.
	override := fmt.Sprintf("resource %q %q { name = %q }\n", "stripe_metronome_rate_card", strings.TrimPrefix(rateCardAddress, "stripe_metronome_rate_card."), "Renamed rate card")
	for _, managed := range state.Values.RootModule.Resources {
		if managed.Type == "stripe_metronome_product" && !strings.HasPrefix(managed.Address, "data.") {
			override += fmt.Sprintf("resource %q %q { name = %q }\n", "stripe_metronome_product", strings.TrimPrefix(managed.Address, "stripe_metronome_product."), "Renamed product")
			break
		}
	}
	if err := os.WriteFile(filepath.Join(root, "metadata_override.tf"), []byte(override), 0600); err != nil {
		t.Fatal(err)
	}
	if out, err := run("apply", "-auto-approve", "-no-color"); err != nil {
		t.Fatalf("update metadata: %v\n%s", err, out)
	}
	if out, err := run("plan", "-detailed-exitcode", "-no-color"); err != nil {
		t.Fatalf("plan after metadata update: %v\n%s", err, out)
	}
	if err := os.Remove(filepath.Join(root, "metadata_override.tf")); err != nil {
		t.Fatal(err)
	}
	if out, err := run("apply", "-auto-approve", "-no-color"); err != nil {
		t.Fatalf("restore metadata: %v\n%s", err, out)
	}

	verifiedHCL := hcl + "\ndata \"stripe_metronome_stripe_billing_provider\" \"direct\" {}\n"
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(verifiedHCL), 0600); err != nil {
		t.Fatal(err)
	}
	if out, err := run("plan", "-no-color"); err == nil || !strings.Contains(strings.Join(strings.Fields(out), " "), "Stripe billing integration") {
		t.Fatalf("expected Stripe integration data source to reject an unconnected account: %v\n%s", err, out)
	}
	if err := os.WriteFile(filepath.Join(root, "main.tf"), []byte(hcl), 0600); err != nil {
		t.Fatal(err)
	}
	if out, err := run("destroy", "-auto-approve", "-no-color"); err != nil {
		t.Fatalf("terraform destroy: %v\n%s", err, out)
	}
	mu.Lock()
	defer mu.Unlock()
	if archived != managedCount {
		t.Errorf("expected %d objects to archive, got %d", managedCount, archived)
	}
	if p := packages["pro-id"]; expectedProKey != "" && (p == nil || strings.TrimPrefix(p["uniqueness_key"].(string), "package:") != strings.TrimPrefix(expectedProKey, "package:")) {
		t.Errorf("missing explicit Pro package key: %+v", p)
	}
	if exampleName == "openrouter-cost-markup" {
		if metric["aggregation_key"] != "cost_usd" || metric["event_type_filter"].(map[string]any)["in_values"].([]any)[0] != "openrouter_usage" {
			t.Errorf("sample event is not selected and summed: %+v", metric)
		}
		if card["fiat_credit_type"].(map[string]any)["id"] != "usd-cents" || rate["rate"].(map[string]any)["price"] != float64(120) {
			t.Errorf("default USD rate or 20%% markup is wrong: card=%+v rate=%+v", card, rate)
		}
		start, err := time.Parse("2006-01-02T15:04:05.000Z", rate["starting_at"].(string))
		if err != nil || start.Before(creationWindowStart) || start.After(creationWindowEnd) {
			t.Errorf("automatic rate start is outside the creation hour: %v, %v", start, err)
		}
	}
	if exampleName == "api-mirror-mixed" {
		if len(allRates) != 2 || allRates[0]["rate"].(map[string]any)["price"] != float64(130) || allRates[1]["billing_frequency"] != "MONTHLY" {
			t.Errorf("mixed card did not create usage and subscription rates: %+v", allRates)
		}
		p := packages["pro-id"]
		if p == nil || len(p["subscriptions"].([]any)) != 1 || len(p["recurring_credits"].([]any)) != 1 {
			t.Errorf("mixed package omitted subscriptions or recurring credits: %+v", p)
		}
	}
}
