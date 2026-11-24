//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"encoding/json"
	"encoding/pem"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func testClient(t *testing.T, handler http.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()
	server := httptest.NewTLSServer(handler)
	t.Cleanup(server.Close)
	ca := filepath.Join(t.TempDir(), "ca.pem")
	if err := os.WriteFile(ca, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600); err != nil {
		t.Fatal(err)
	}
	client, err := New("sk_test_synthetic", server.URL, ca, "qualification")
	if err != nil {
		t.Fatal(err)
	}
	return client, server
}

func TestEveryReviewedSDKBinding(t *testing.T) {
	t.Setenv("METRONOME_BEARER_TOKEN", "must-not-be-used")
	t.Setenv("METRONOME_API_TOKEN", "must-not-be-used")
	t.Setenv("METRONOME_CUSTOM_HEADERS", "Authorization: Bearer must-not-be-used\nX-Unwanted: true")
	t.Setenv("METRONOME_BASE_URL", "https://unwanted.invalid")
	const payload = `{"number":123456789012345678901234567890.1234567890123456789,"null":null,"empty":[],"opaque":{"flag":false}}`
	const result = `{"data":{"id":"fixture","number":123456789012345678901234567890.1234567890123456789,"extra":{"nested":null}},"next_page":"must-not-follow"}`
	var active Operation
	var expectedPath string
	var calls int
	client, _ := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Method != active.Method || r.URL.Path != expectedPath {
			t.Errorf("%s: got %s %s, expected %s %s", active.ID, r.Method, r.URL.Path, active.Method, expectedPath)
		}
		if r.Header.Get("Authorization") != "Bearer sk_test_synthetic" || r.Header.Get("X-Unwanted") != "" {
			t.Error("SDK used ambient credentials or headers")
		}
		if r.URL.Query().Get("limit") != "1" || len(r.URL.Query()["test"]) != 2 {
			t.Error("SDK changed query parameters")
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		if string(body) != payload {
			t.Errorf("%s: request body changed: %s", active.ID, body)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, result)
	})
	for _, operation := range sdkOperations {
		op := operation.Operation
		t.Run(op.ID, func(t *testing.T) {
			active = op
			calls = 0
			params := map[string]string{}
			placeholder := regexp.MustCompile(`\{([^}]+)\}`)
			expectedPath = placeholder.ReplaceAllStringFunc(op.Path, func(p string) string { params[p[1:len(p)-1]] = "id with %"; return "id with %" })
			got, err := client.Execute(context.Background(), op, params, url.Values{"limit": {"1"}, "test": {"first", "second"}}, []byte(payload))
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != result {
				t.Error("SDK changed response bytes, precision, or opaque fields")
			}
			if calls != 1 {
				t.Errorf("expected exactly one request/page; got %d", calls)
			}
		})
	}
	if len(sdkOperations) != 131 {
		t.Fatalf("operation coverage changed: %d", len(sdkOperations))
	}
}

func TestReadAndWriteRetryContracts(t *testing.T) {
	for _, tc := range []struct {
		name         string
		op           Operation
		status, want int
	}{
		{"write", Operation{ID: "createProduct-v1", Method: "POST", Path: "/v1/contract-pricing/products/create"}, 503, 1},
		{"post-read", Operation{ID: "getProduct-v1", Method: "POST", Path: "/v1/contract-pricing/products/get", Read: true}, 503, 3},
		{"rate-limit", Operation{ID: "listCreditTypes-v1", Method: "GET", Path: "/v1/credit-types/list", Read: true}, 429, 3},
		{"permission", Operation{ID: "listCreditTypes-v1", Method: "GET", Path: "/v1/credit-types/list", Read: true}, 403, 1},
		{"missing", Operation{ID: "listCreditTypes-v1", Method: "GET", Path: "/v1/credit-types/list", Read: true}, 404, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var calls atomic.Int32
			client, _ := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				calls.Add(1)
				w.WriteHeader(tc.status)
				_, _ = io.WriteString(w, `{"message":"sk_test_synthetic echoed-secret"}`)
			})
			_, err := client.Execute(context.Background(), tc.op, nil, nil, []byte(`{}`))
			if err == nil || HTTPStatus(err) != tc.status {
				t.Fatalf("unexpected error: %v", err)
			}
			if strings.Contains(err.Error(), "synthetic") || strings.Contains(err.Error(), "echoed-secret") {
				t.Fatal("response body leaked into diagnostic")
			}
			if int(calls.Load()) != tc.want {
				t.Fatalf("attempt count: %d, expected %d", calls.Load(), tc.want)
			}
		})
	}
}

func TestCancellationAndRedirectIsolation(t *testing.T) {
	var targetCalls atomic.Int32
	target := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { targetCalls.Add(1) }))
	defer target.Close()
	client, _ := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", target.URL)
		w.WriteHeader(307)
	})
	op := Operation{ID: "listCreditTypes-v1", Method: "GET", Path: "/v1/credit-types/list", Read: true}
	_, err := client.Execute(context.Background(), op, nil, nil, nil)
	if HTTPStatus(err) != 307 || targetCalls.Load() != 0 {
		t.Fatalf("redirect escaped origin: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = client.Execute(ctx, op, nil, nil, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("cancellation lost: %v", err)
	}
	if _, err = New("", DefaultBaseURL, "", "test"); err == nil {
		t.Fatal("empty key accepted")
	}
	for _, origin := range []string{"http://example.com", "https://user:pass@example.com", "https://example.com/path", "https://example.com?query", "https://example.com#fragment"} {
		if _, err = New("key", origin, "", "test"); err == nil {
			t.Fatalf("invalid origin accepted: %s", origin)
		}
	}
}

func TestNoUnreviewedOperationOrRetryEscalation(t *testing.T) {
	var calls atomic.Int32
	client, _ := testClient(t, func(w http.ResponseWriter, r *http.Request) { calls.Add(1); _, _ = io.WriteString(w, `{}`) })
	for _, op := range []Operation{
		{ID: "invented", Method: "GET", Path: "/v1/services", Read: true},
		{ID: "createProduct-v1", Method: "GET", Path: "/v1/contract-pricing/products/create"},
		{ID: "createProduct-v1", Method: "POST", Path: "/v1/contract-pricing/products/create", Read: true},
		{ID: "listCreditTypes-v1", Method: "GET", Path: "//unwanted.invalid", Read: true},
	} {
		if _, err := client.Execute(context.Background(), op, nil, nil, nil); err == nil {
			t.Fatalf("unreviewed binding accepted: %v", op)
		}
	}
	if calls.Load() != 0 {
		t.Fatal("invalid binding reached API")
	}
	if retryAfter("99999999") != 5*time.Second || retryAfter("-1") != 0 {
		t.Fatal("retry-after is not bounded")
	}
}

func TestTypedSDKCallsPreserveNumbersAndSeparateQueryParameters(t *testing.T) {
	const exact = "12345678901234567890.1234567890123456789"
	var calls int
	client, _ := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		body, _ := io.ReadAll(r.Body)
		switch r.URL.Path {
		case "/v1/contract-pricing/rate-cards/addRate":
			var params map[string]json.RawMessage
			if err := json.Unmarshal(body, &params); err != nil {
				t.Fatal(err)
			}
			if string(params["price"]) != exact || r.Header.Get("Content-Type") != "application/json" {
				t.Error("typed SDK request rounded a number or used form encoding")
			}
			if _, present := params["ending_before"]; present {
				t.Error("omitted optional field was serialized")
			}
			_, _ = io.WriteString(w, `{"data":{"id":"rate"}}`)
		case "/v1/credit-types/list":
			if r.Method != "GET" || r.URL.Query().Get("limit") != "1" || len(body) != 0 {
				t.Error("typed SDK query became a request body")
			}
			_, _ = io.WriteString(w, `{"data":[{"id":"unit","name":"credits","is_currency":false}],"next_page":null}`)
		default:
			t.Errorf("unexpected typed SDK path: %s", r.URL.Path)
		}
	})
	price, entitled := Number(exact), true
	_, err := client.V1RateCards.AddRate(context.Background(), &V1RateCardsAddRateParams{
		RateCardID: String("card"), ProductID: String("product"), RateType: String("FLAT"), Price: &price, Entitled: &entitled, StartingAt: String("2026-01-01T00:00:00Z"),
	})
	if err != nil {
		t.Fatal(err)
	}
	limit := int64(1)
	response, err := client.V1Settings.ListCreditTypes(context.Background(), &V1SettingsListCreditTypesParams{Limit: &limit})
	if err != nil {
		t.Fatal(err)
	}
	if len(response.Data) != 1 || response.Data[0].ID != "unit" || len(response.RawJSON) == 0 || calls != 2 {
		t.Fatal("typed SDK response or call count differs")
	}
}

func TestTypedSDKArrayBodyAndRepeatedQuery(t *testing.T) {
	client, _ := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		body, _ := io.ReadAll(r.Body)
		switch r.URL.Path {
		case "/v1/ingest":
			var events []struct {
				TransactionID string `json:"transaction_id"`
			}
			if err := json.Unmarshal(body, &events); err != nil || len(events) != 1 || events[0].TransactionID != "transaction" {
				t.Error("ingest request was not a top-level event array")
			}
			_, _ = io.WriteString(w, `{}`)
		case "/v1/customers":
			values := r.URL.Query()["customer_ids"]
			if len(values) != 2 || values[0] != "first" || values[1] != "second" || len(body) != 0 {
				t.Error("query array was not encoded as repeated parameters")
			}
			_, _ = io.WriteString(w, `{"data":[],"next_page":null}`)
		default:
			t.Errorf("unexpected typed SDK path: %s", r.URL.Path)
		}
	})
	_, err := client.V1Usage.Ingest(context.Background(), &V1UsageIngestParams{Body: []*V1UsageIngestBodyParams{{
		CustomerID: String("customer"), EventType: String("event"), Timestamp: String("2026-01-01T00:00:00Z"), TransactionID: String("transaction"),
	}}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.V1Customers.ListCustomers(context.Background(), &V1CustomersListCustomersParams{CustomerIDs: []*string{String("first"), String("second")}})
	if err != nil {
		t.Fatal(err)
	}
}
