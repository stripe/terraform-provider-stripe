//
// File generated from our OpenAPI spec
//

package providerdata

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	stripe "github.com/stripe/stripe-go/v86"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

type rawBackendFunc func(string, string, string, string, *stripe.RawParams) (*stripe.APIResponse, error)

func (f rawBackendFunc) RawRequest(m, p, k, b string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	return f(m, p, k, b, params)
}

func TestActivationOriginBodyAuthAndNoRetries(t *testing.T) {
	saved := http.DefaultTransport
	defer func() { http.DefaultTransport = saved }()
	calls := 0
	http.DefaultTransport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		calls++
		if r.URL.String() != "https://api.stripe.com/v1/experimental-met-activate" || r.Method != "POST" {
			t.Error("activation reached another origin or method")
		}
		if r.Header.Get("Stripe-Account") != "" {
			t.Error("Connect context leaked into activation")
		}
		body, _ := io.ReadAll(r.Body)
		if len(body) != 0 {
			t.Error("activation body must be empty")
		}
		if !strings.Contains(r.Header.Get("Authorization"), "sk_test_fixture") {
			t.Error("shared Stripe key was not supplied")
		}
		return &http.Response{StatusCode: 503, Header: http.Header{"Content-Type": []string{"application/json"}}, Body: io.NopCloser(strings.NewReader(`{"error":{"message":"private echoed data"}}`)), Request: r}, nil
	})
	backend := NewActivationBackend()
	_, err := backend.RawRequest("POST", "/v1/experimental-met-activate", "sk_test_fixture", "", &stripe.RawParams{Params: stripe.Params{Context: context.Background()}})
	if err == nil || calls != 1 {
		t.Fatalf("activation retried or ignored failure: calls=%d err=%v", calls, err)
	}
	request, _ := http.NewRequest("POST", "https://metronome.invalid/v1/experimental-met-activate", nil)
	transport := stripeOriginTransport{base: http.DefaultTransport}
	if _, err := transport.RoundTrip(request); err == nil || calls != 1 {
		t.Fatal("activation escaped pinned origin")
	}
}

func TestAccountModeAndConnectContextAreExplicit(t *testing.T) {
	for _, test := range []struct {
		key, context, mode string
		valid              bool
	}{
		{"sk_test_fixture", "", "test", true}, {"rkcs_test_fixture", "", "test", true}, {"sk_live_fixture", "", "live", true},
		{"rk_test_fixture", "", "", false}, {"native-metronome-token", "", "", false}, {"sk_test_fixture", "acct_other", "", false}, {"sk_test_fixture", "acct_fixture", "test", true},
	} {
		t.Run(test.key+test.context, func(t *testing.T) {
			calls := 0
			data := &Data{StripeKey: test.key, StripeAccount: test.context, StripeStatus: Status{State: "ready"}, ActivationBackend: rawBackendFunc(func(method, path, key, body string, params *stripe.RawParams) (*stripe.APIResponse, error) {
				calls++
				if method != "GET" || path != "/v1/account" || key != test.key {
					t.Fatal("identity lookup changed")
				}
				return &stripe.APIResponse{RawJSON: []byte(`{"id":"acct_fixture"}`)}, nil
			})}
			raw, err := data.Execute(context.Background(), bindingruntime.Operation{ID: "stripeAccountIdentity", Service: "stripe", Method: "GET", Path: "/v1/account", Read: true}, nil, nil, nil)
			if test.valid {
				if err != nil {
					t.Fatal(err)
				}
				var result map[string]string
				if err = json.Unmarshal(raw, &result); err != nil {
					t.Fatal(err)
				}
				if result["metronome_mode"] != test.mode || result["id"] != "acct_fixture" {
					t.Fatal("identity/mode changed")
				}
			} else if err == nil {
				t.Fatal("unsupported account/key context was accepted")
			}
			if !strings.HasPrefix(test.key, "sk_") && !strings.HasPrefix(test.key, "rkcs_test_") && calls != 0 {
				t.Fatal("unsupported native or restricted key sent to identity endpoint")
			}
		})
	}
}

func TestActivationDiagnosticsDoNotExposeAPIResponse(t *testing.T) {
	data := &Data{StripeKey: "sk_test_fixture", ActivationBackend: rawBackendFunc(func(string, string, string, string, *stripe.RawParams) (*stripe.APIResponse, error) {
		return nil, &stripe.Error{HTTPStatusCode: 403, Msg: "sk_test_fixture sensitive body"}
	})}
	_, err := data.stripeRequest(context.Background(), "POST", "/v1/experimental-met-activate")
	if err == nil || strings.Contains(err.Error(), "sensitive") || strings.Contains(err.Error(), "sk_test_fixture") || data.HTTPStatus(err) != 403 {
		t.Fatalf("unsafe diagnostic: %v", err)
	}
}
