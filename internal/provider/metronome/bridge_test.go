//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/terraform-provider-stripe/internal/providerdata"
)

type activationBackend struct {
	t                    *testing.T
	account, key, client string
	activationStatus     int
	calls                []string
}

func (b *activationBackend) RawRequest(method, path, key, body string, params *stripe.RawParams) (*stripe.APIResponse, error) {
	b.t.Helper()
	b.calls = append(b.calls, method+" "+path)
	if key != b.key || body != "" || params.Context == nil {
		b.t.Fatal("activation credentials, body, or context changed")
	}
	switch {
	case method == "GET" && path == "/v1/account":
		return &stripe.APIResponse{RawJSON: []byte(fmt.Sprintf(`{"id":%q}`, b.account))}, nil
	case method == "POST" && path == "/v1/experimental-met-activate":
		if b.activationStatus != 0 {
			return nil, &stripe.Error{HTTPStatusCode: b.activationStatus, Msg: "sensitive-API-response"}
		}
		return &stripe.APIResponse{RawJSON: []byte(`{"status":"complete","client_id":` + b.client + `}`)}, nil
	default:
		b.t.Fatalf("unexpected Stripe call %s %s", method, path)
		return nil, nil
	}
}
func bridgeFixture(t *testing.T, calls []apiCall) (*MetronomeAccountBridgeResource, *activationBackend) {
	t.Helper()
	r := resourceFixture(t, NewMetronomeAccountBridgeResource, calls).(*MetronomeAccountBridgeResource)
	data := r.transport.(*providerdata.Data)
	backend := &activationBackend{t: t, account: "acct_fixture", key: "sk_test_fixture", client: "null"}
	data.StripeKey = backend.key
	data.StripeStatus = providerdata.Status{State: "ready"}
	data.ActivationBackend = backend
	return r, backend
}

const bridgeState = `{"id":"acct_fixture/test","stripe_account_id":"acct_fixture","mode":"test","client_id":null,"readiness_timeout":"7m"}`

func TestBridgeWaitsThenOnlyReadsAndForgets(t *testing.T) {
	r, backend := bridgeFixture(t, []apiCall{
		{method: "GET", path: "/v1/credit-types/list", query: "limit=1", status: 403},
		{method: "GET", path: "/v1/credit-types/list", query: "limit=1", response: `{"data":[],"next_page":null}`},
		{method: "GET", path: "/v1/credit-types/list", query: "limit=1", response: `{"data":[{"id":"unit","name":"do-not-persist","is_currency":true}],"next_page":"do-not-follow"}`},
	})
	state := createFixture(t, r, `{"readiness_timeout":"7m"}`)
	v := stateValues(t, state)
	if v.String("id") != "acct_fixture/test" || !v.Get("client_id").IsNull() {
		t.Fatal("bridge lost stable identity or rejected null client_id")
	}
	writes := 0
	for _, call := range backend.calls {
		if strings.HasPrefix(call, "POST") {
			writes++
		}
	}
	if writes != 1 {
		t.Fatalf("activation ran %d times", writes)
	}
	before := len(backend.calls)
	state = readFixture(t, r, state)
	deleteFixture(t, r, state)
	if len(backend.calls) != before+1 || backend.calls[len(backend.calls)-1] != "GET /v1/account" {
		t.Fatal("refresh/delete activated or revoked a bridge")
	}
	encoded, _ := json.Marshal(stateValues(t, state).Object.String())
	if strings.Contains(string(encoded), "do-not-persist") || strings.Contains(string(encoded), "sk_test_fixture") {
		t.Fatal("probe data or credentials persisted")
	}
}

func TestBridgeProbeClassificationAndMalformedSuccess(t *testing.T) {
	for _, test := range []struct {
		status   int
		body     string
		missing  bool
		attempts int
	}{
		{401, "", true, 1}, {403, "", true, 1}, {404, "", true, 1}, {400, "", false, 1}, {409, "", false, 1}, {429, "", false, 3}, {500, "", false, 3}, {503, "", false, 3},
		{200, `{"data":{}}`, false, 1}, {200, `{"data":[{"id":"a","name":"secret"}]}`, false, 1}, {200, `not JSON`, false, 1},
	} {
		t.Run(fmt.Sprintf("%d-%s", test.status, test.body), func(t *testing.T) {
			calls := make([]apiCall, test.attempts)
			for i := range calls {
				calls[i] = apiCall{method: "GET", path: "/v1/credit-types/list", query: "limit=1", status: test.status, response: test.body}
			}
			r, backend := bridgeFixture(t, calls)
			state := fixtureState(t, r, bridgeState, nil)
			response := resource.ReadResponse{State: state}
			r.Read(context.Background(), resource.ReadRequest{State: state}, &response)
			if test.missing {
				if response.Diagnostics.HasError() || !response.State.Raw.IsNull() {
					t.Fatalf("missing bridge retained: %v", response.Diagnostics)
				}
			} else if !response.Diagnostics.HasError() || !response.State.Raw.Equal(state.Raw) {
				t.Fatalf("failed probe destroyed state: %v", response.Diagnostics)
			}
			if len(backend.calls) != 1 || backend.calls[0] != "GET /v1/account" {
				t.Fatal("read did not verify the account without mutation")
			}
		})
	}
}

func TestBridgeActivationFailureIsNeverMissingOrRetried(t *testing.T) {
	for _, status := range []int{400, 401, 403, 404, 500} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			r, backend := bridgeFixture(t, nil)
			backend.activationStatus = status
			state := fixtureState(t, r, `{"readiness_timeout":"7m"}`, nil)
			response := resource.CreateResponse{State: tfsdk.State{Schema: state.Schema}}
			r.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan{Schema: state.Schema, Raw: state.Raw}}, &response)
			if !response.Diagnostics.HasError() {
				t.Fatal("activation failure reported success")
			}
			writes := 0
			for _, call := range backend.calls {
				if strings.HasPrefix(call, "POST") {
					writes++
				}
			}
			if writes != 1 {
				t.Fatalf("activation retried: %d writes", writes)
			}
			if strings.Contains(fmt.Sprint(response.Diagnostics), "sensitive-API-response") {
				t.Fatal("activation response leaked")
			}
		})
	}
}

func TestBridgePostActivationTimeoutRetainsState(t *testing.T) {
	r, backend := bridgeFixture(t, []apiCall{{method: "GET", path: "/v1/credit-types/list", query: "limit=1", status: 403}})
	state := fixtureState(t, r, `{"readiness_timeout":"7m"}`, nil)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	response := resource.CreateResponse{State: tfsdk.State{Schema: state.Schema}}
	r.Create(ctx, resource.CreateRequest{Plan: tfsdk.Plan{Schema: state.Schema, Raw: state.Raw}}, &response)
	if !response.Diagnostics.HasError() || stateValues(t, response.State).String("id") != "acct_fixture/test" {
		t.Fatal("upsert identity lost after canceled readiness")
	}
	writes := 0
	for _, call := range backend.calls {
		if strings.HasPrefix(call, "POST") {
			writes++
		}
	}
	if writes != 1 {
		t.Fatal("cancellation repeated upsert")
	}
}

func TestBridgeImportRotationAndScopeMismatch(t *testing.T) {
	r, backend := bridgeFixture(t, []apiCall{
		{method: "GET", path: "/v1/credit-types/list", query: "limit=1", response: `{"data":[]}`},
	})
	initial := fixtureState(t, r, `{}`, nil)
	response := resource.ImportStateResponse{State: tfsdk.State{Schema: initial.Schema}}
	r.ImportState(context.Background(), resource.ImportStateRequest{ID: "acct_fixture/test"}, &response)
	if response.Diagnostics.HasError() || len(backend.calls) != 0 {
		t.Fatal("import performed activation or rejected identity")
	}
	backend.key = "sk_test_rotated"
	r.transport.(*providerdata.Data).StripeKey = backend.key
	state := readFixture(t, r, response.State)
	backend.account = "acct_different"
	read := resource.ReadResponse{State: state}
	r.Read(context.Background(), resource.ReadRequest{State: state}, &read)
	if !read.Diagnostics.HasError() || !read.State.Raw.Equal(state.Raw) {
		t.Fatal("key scope mismatch replaced existing bridge identity")
	}
	for _, bad := range []string{"unit", "acct_fixture/preview", "acct_/test", "acct_fixture/test/extra"} {
		invalid := resource.ImportStateResponse{State: tfsdk.State{Schema: initial.Schema}}
		r.ImportState(context.Background(), resource.ImportStateRequest{ID: bad}, &invalid)
		if !invalid.Diagnostics.HasError() {
			t.Fatalf("invalid bridge import accepted: %q", bad)
		}
	}
}

func TestBridgeValidationCoversAuthorizationCache(t *testing.T) {
	r := NewMetronomeAccountBridgeResource()
	for _, timeout := range []string{"5m", "5m29s", "garbage", "61m"} {
		state := fixtureState(t, r, fmt.Sprintf(`{"readiness_timeout":%q}`, timeout), nil)
		var response resource.ValidateConfigResponse
		r.(resource.ResourceWithValidateConfig).ValidateConfig(context.Background(), resource.ValidateConfigRequest{Config: tfsdk.Config{Schema: state.Schema, Raw: state.Raw}}, &response)
		if !response.Diagnostics.HasError() {
			t.Fatalf("unsafe readiness timeout accepted: %s", timeout)
		}
	}
	for _, timeout := range []string{"5m30s", "7m", "1h"} {
		state := fixtureState(t, r, fmt.Sprintf(`{"readiness_timeout":%q}`, timeout), nil)
		var response resource.ValidateConfigResponse
		r.(resource.ResourceWithValidateConfig).ValidateConfig(context.Background(), resource.ValidateConfigRequest{Config: tfsdk.Config{Schema: state.Schema, Raw: state.Raw}}, &response)
		if response.Diagnostics.HasError() {
			t.Fatal(response.Diagnostics)
		}
	}
}
