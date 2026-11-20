//
// File generated from our OpenAPI spec
//

package providerdata

import (
	"context"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
	metronomeclient "github.com/stripe/terraform-provider-stripe/metronome-go"
)

func TestMetronomeUsesSharedKeyWithoutStripeRequests(t *testing.T) {
	calls := 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Method != http.MethodGet || r.URL.Path != "/v1/credit-types/list" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL)
		}
		if r.Header.Get("Authorization") != "Bearer sk_test_fixture" || r.Header.Get("Stripe-Account") != "" {
			t.Error("Metronome request did not use only the shared Stripe key")
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":[]}`))
	}))
	defer server.Close()
	ca := filepath.Join(t.TempDir(), "ca.pem")
	if err := os.WriteFile(ca, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600); err != nil {
		t.Fatal(err)
	}
	client, err := metronomeclient.New("sk_test_fixture", server.URL, ca, "test")
	if err != nil {
		t.Fatal(err)
	}
	data := &Data{Metronome: client, MetronomeStatus: Status{State: "ready"}, StripeStatus: Status{State: "unconfigured"}}
	op := bindingruntime.Operation{ID: "listCreditTypes-v1", Service: "metronome", Method: "GET", Path: "/v1/credit-types/list", Read: true}
	result, err := data.Execute(context.Background(), op, nil, nil, nil)
	if err != nil || string(result) != `{"data":[]}` || calls != 1 {
		t.Fatalf("direct Metronome request failed: calls=%d result=%s err=%v", calls, result, err)
	}
	data.StripeAccount = "acct_other"
	if _, err = data.Execute(context.Background(), op, nil, nil, nil); err == nil || !strings.Contains(err.Error(), "stripe_account") || calls != 1 {
		t.Fatalf("unsupported Connect context reached the API: calls=%d err=%v", calls, err)
	}
}

func TestUnavailableMetronomeClientReturnsDiagnostic(t *testing.T) {
	for _, status := range []Status{{State: "missing"}, {State: "unknown"}, {State: "invalid"}, {State: "ready"}} {
		t.Run(status.State, func(t *testing.T) {
			data := &Data{MetronomeStatus: status}
			_, err := data.Execute(context.Background(), bindingruntime.Operation{Service: "metronome"}, nil, nil, nil)
			if err == nil {
				t.Fatal("unavailable Metronome client was accepted")
			}
		})
	}
}
