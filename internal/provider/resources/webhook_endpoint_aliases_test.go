// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	stripe "github.com/stripe/stripe-go/v87"
)

func TestWebhookEndpointAccountDefaultVersion(t *testing.T) {
	current := "payment_method.automatically_updated"
	legacy := "payment_method.card_automatically_updated"
	for _, tc := range []struct {
		name, version string
		status        int
		wantError     bool
		create        bool
	}{
		{"old account default", "2019-12-03", http.StatusOK, false, false},
		{"modern account default", "2026-05-27.dahlia", http.StatusOK, false, false},
		{"missing version", "", http.StatusOK, true, false},
		{"API error", "", http.StatusForbidden, true, false},
		{"old account before create", "2019-12-03", http.StatusOK, false, true},
		{"modern account before create", "2026-05-27.dahlia", http.StatusOK, false, true},
		{"missing version before create", "", http.StatusOK, true, true},
		{"API error before create", "", http.StatusForbidden, true, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var requests atomic.Int64
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				requests.Add(1)
				if req.Method == "POST" {
					if !tc.create || tc.wantError {
						t.Error("creation continued after failed version lookup")
					}
					if req.Header.Get("Stripe-Version") != stripe.APIVersion {
						t.Error("lookup changed the SDK request version")
					}
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprintf(w, `{"id":"we_version","object":"webhook_endpoint","api_version":null,"enabled_events":[%q,%q]}`, current, legacy)
					return
				}
				expectedPath := "/v1/webhook_endpoints/we_version"
				if tc.create {
					expectedPath = "/v1/webhook_endpoints"
					if req.URL.RawQuery != "limit=1" {
						t.Error("pre-create lookup must request at most one endpoint")
					}
				}
				if req.Method != "GET" || req.URL.Path != expectedPath {
					t.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
				}
				if _, exists := req.Header["Stripe-Version"]; exists {
					t.Error("account-default lookup must omit Stripe-Version entirely")
				}
				if req.Header.Get("Authorization") != "Bearer test-key" || req.Header.Get("Stripe-Context") != "acct_context" {
					t.Error("lookup lost authentication or account context")
				}
				if tc.version != "" {
					w.Header().Set("Stripe-Version", tc.version)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.status)
				if tc.status != http.StatusOK {
					fmt.Fprint(w, `{"error":{"type":"invalid_request_error","message":"test failure"}}`)
					return
				}
				fmt.Fprint(w, `{"id":"we_version","object":"webhook_endpoint","api_version":null,"enabled_events":[]}`)
			}))
			defer server.Close()
			backend := stripe.GetBackendWithConfig(stripe.APIBackend, &stripe.BackendConfig{URL: stripe.String(server.URL), HTTPClient: server.Client(), MaxNetworkRetries: stripe.Int64(0), StripeContext: stripe.String("acct_context")})
			client := stripe.NewClient("test-key", stripe.WithBackends(&stripe.Backends{API: &stripe.UsageBackend{B: backend}}))
			r := WebhookEndpointResource{client: client}
			obj := &stripe.WebhookEndpoint{ID: "we_version", EnabledEvents: []string{current, legacy}}
			data, err := json.Marshal(map[string]interface{}{"enabled_events": obj.EnabledEvents, "api_version": nil})
			if err != nil {
				t.Fatal(err)
			}
			obj.SetLastResponse(&stripe.APIResponse{RawJSON: data})
			state := WebhookEndpointResourceModel{EnabledEvents: webhookEventList(current), APIVersion: types.StringNull(), URL: types.StringValue("https://example.com/webhook"), Metadata: types.MapNull(types.StringType)}
			if tc.create {
				ctx := context.Background()
				var schemaResp resource.SchemaResponse
				r.Schema(ctx, resource.SchemaRequest{}, &schemaResp)
				plan := tfsdk.Plan{Schema: schemaResp.Schema}
				if diags := plan.Set(ctx, state); diags.HasError() {
					t.Fatal(diags)
				}
				resp := resource.CreateResponse{State: tfsdk.State{Schema: schemaResp.Schema, Raw: plan.Raw}}
				r.Create(ctx, resource.CreateRequest{Plan: plan}, &resp)
				if resp.Diagnostics.HasError() {
					err = fmt.Errorf("%v", resp.Diagnostics)
				}
				if diags := resp.State.Get(ctx, &state); diags.HasError() {
					t.Fatal(diags)
				}
			} else {
				err = r.flattenWebhookEndpoint(context.Background(), obj, &state)
			}
			if (err != nil) != tc.wantError {
				t.Fatalf("unexpected error: %v", err)
			}
			wantRequests := int64(1)
			if tc.create && !tc.wantError {
				wantRequests++
			}
			if requests.Load() != wantRequests {
				t.Fatalf("got %d requests; want %d", requests.Load(), wantRequests)
			}
			if !state.APIVersion.IsNull() {
				t.Fatal("account-default lookup changed endpoint api_version")
			}
			want := webhookEventList(current)
			if tc.version == "2019-12-03" {
				want = webhookEventList(current, legacy)
			}
			if !state.EnabledEvents.Equal(want) {
				t.Fatalf("got %s; want %s", state.EnabledEvents, want)
			}
		})
	}
}

func TestWebhookEndpointVersionLookupBypass(t *testing.T) {
	for _, obj := range []*stripe.WebhookEndpoint{
		{APIVersion: "2019-12-03", EnabledEvents: []string{"payment_method.automatically_updated", "payment_method.card_automatically_updated"}},
		{EnabledEvents: []string{"customer.created"}},
		{EnabledEvents: []string{"*"}},
		{EnabledEvents: []string{"transfer.reversed"}},
	} {
		version, err := resolveWebhookEndpointEventAPIVersion(context.Background(), obj, nil, "unused")
		if err != nil || version != obj.APIVersion {
			t.Fatalf("unexpected lookup: %q %v", version, err)
		}
	}
	_, err := resolveWebhookEndpointEventAPIVersion(context.Background(), &stripe.WebhookEndpoint{EnabledEvents: []string{"transfer.reversed", "transfer.canceled"}}, nil, "unused")
	if err == nil {
		t.Fatal("ambiguous aliases must not guess a version for an unsupported backend")
	}
}

func webhookEventList(events ...string) types.List {
	values := make([]attr.Value, len(events))
	for i, event := range events {
		values[i] = types.StringValue(event)
	}
	return types.ListValueMust(types.StringType, values)
}

func TestFlattenWebhookEndpointEventAliases(t *testing.T) {
	modern := []string{"transfer.reversed", "payment_method.automatically_updated"}
	expanded := []string{"transfer.reversed", "transfer.canceled", "payment_method.automatically_updated", "payment_method.card_automatically_updated"}
	tests := []struct {
		name       string
		observed   []string
		configured types.List
		want       types.List
	}{
		{"current names", expanded, webhookEventList(modern...), webhookEventList(modern...)},
		{"legacy names", expanded, webhookEventList("transfer.canceled", "payment_method.card_automatically_updated"), webhookEventList("transfer.canceled", "payment_method.card_automatically_updated")},
		{"both names", expanded, webhookEventList(expanded...), webhookEventList(expanded...)},
		{"transfer alone", expanded[:2], webhookEventList(modern[0]), webhookEventList(modern[0])},
		{"payment method alone", expanded[2:], webhookEventList(modern[1]), webhookEventList(modern[1])},
		{"duplicate current name", expanded[:2], webhookEventList(modern[0], modern[0]), webhookEventList(modern[0], modern[0])},
		{"configured order", expanded, webhookEventList(modern[1], modern[0]), webhookEventList(modern[1], modern[0])},
		{"already canonical", modern, webhookEventList(modern...), webhookEventList(modern...)},
		{"remote addition", append(append([]string{}, expanded...), "customer.created"), webhookEventList(modern...), webhookEventList(modern[0], modern[1], "customer.created")},
		{"remote removal", expanded[:2], webhookEventList(modern...), webhookEventList(modern[0])},
		{"different subscription", []string{"customer.created"}, webhookEventList(modern...), webhookEventList("customer.created")},
		{"empty subscription", []string{}, webhookEventList(modern...), webhookEventList()},
		{"wildcard", []string{"*"}, webhookEventList("*"), webhookEventList("*")},
		{"wildcard drift", []string{"*"}, webhookEventList(modern...), webhookEventList("*")},
		{"ignored legacy request stays visible", []string{}, webhookEventList("transfer.canceled", "payment_method.card_automatically_updated"), webhookEventList()},
		{"legacy-only response", []string{"payment_method.card_automatically_updated"}, webhookEventList("payment_method.card_automatically_updated"), webhookEventList("payment_method.card_automatically_updated")},
		{"legacy-only import", []string{"payment_method.card_automatically_updated"}, types.ListNull(types.StringType), webhookEventList("payment_method.card_automatically_updated")},
		{"older transfer semantics stay visible", []string{"transfer.canceled", "payout.canceled", "recipient_transfer.canceled"}, webhookEventList("transfer.canceled"), webhookEventList("transfer.canceled", "payout.canceled", "recipient_transfer.canceled")},
		{"legacy-only is not assumed modern", []string{"transfer.canceled"}, webhookEventList("transfer.reversed"), webhookEventList("transfer.canceled")},
		{"import", expanded, types.ListNull(types.StringType), webhookEventList(modern...)},
		{"unknown prior", expanded, types.ListUnknown(types.StringType), webhookEventList(modern...)},
		{"unknown element", expanded, types.ListValueMust(types.StringType, []attr.Value{types.StringUnknown()}), webhookEventList(modern...)},
		{"unrelated order", []string{"customer.created", "customer.updated"}, webhookEventList("customer.updated", "customer.created"), webhookEventList("customer.updated", "customer.created")},
	}
	for _, tc := range tests {
		for _, raw := range []bool{true, false} {
			name := "typed/"
			if raw {
				name = "raw/"
			}
			t.Run(name+tc.name, func(t *testing.T) {
				obj := &stripe.WebhookEndpoint{EnabledEvents: tc.observed}
				if raw {
					data, err := json.Marshal(map[string]interface{}{"enabled_events": tc.observed})
					if err != nil {
						t.Fatal(err)
					}
					obj.SetLastResponse(&stripe.APIResponse{RawJSON: data})
				}
				state := WebhookEndpointResourceModel{EnabledEvents: tc.configured}
				if err := flattenWebhookEndpoint(obj, &state, "2026-05-27.dahlia"); err != nil {
					t.Fatal(err)
				}
				if !state.EnabledEvents.Equal(tc.want) {
					t.Fatalf("got %s; want %s", state.EnabledEvents, tc.want)
				}
			})
		}
	}
}

func TestPreserveWebhookEndpointUnknownResponse(t *testing.T) {
	for _, value := range []types.List{types.ListNull(types.StringType), types.ListUnknown(types.StringType), types.ListValueMust(types.StringType, []attr.Value{types.StringUnknown()}), types.ListValueMust(types.StringType, []attr.Value{types.StringNull()})} {
		if got := preserveWebhookEndpointEventAliases(value, webhookEventList("transfer.reversed"), ""); !got.Equal(value) {
			t.Fatalf("got %s; want %s", got, value)
		}
	}
}

func TestFlattenWebhookEndpointVersionedAliases(t *testing.T) {
	transfer := []string{"transfer.reversed", "transfer.canceled"}
	payment := []string{"payment_method.automatically_updated", "payment_method.card_automatically_updated"}
	tests := []struct {
		name, version    string
		observed         []string
		configured, want types.List
	}{
		{"old payment alias addition", "2019-12-03", payment, webhookEventList(payment[0]), webhookEventList(payment...)},
		{"old payment alias removal", "2019-12-03", payment[:1], webhookEventList(payment...), webhookEventList(payment[0])},
		{"old payment import", "2019-12-03", payment, types.ListNull(types.StringType), webhookEventList(payment...)},
		{"payment version boundary", "2020-08-27", payment, webhookEventList(payment[0]), webhookEventList(payment[0])},
		{"old transfer import", "2015-01-26", append(append([]string{}, transfer...), "payout.canceled", "recipient_transfer.canceled"), types.ListNull(types.StringType), webhookEventList(transfer[0], transfer[1], "payout.canceled", "recipient_transfer.canceled")},
		{"transfer version boundary", "2017-04-06", transfer, webhookEventList(transfer[0]), webhookEventList(transfer[0])},
		{"mixed version semantics", "2019-12-03", append(append([]string{}, transfer...), payment...), types.ListNull(types.StringType), webhookEventList(transfer[0], payment[0], payment[1])},
		{"modern version suffix", "2026-05-27.dahlia", payment, webhookEventList(payment[0]), webhookEventList(payment[0])},
		{"modern missing alias is not invented", "2026-05-27.dahlia", payment[:1], webhookEventList(payment...), webhookEventList(payment[0])},
		{"unknown version retains distinct names", "", payment, types.ListNull(types.StringType), webhookEventList(payment...)},
		{"missing alias is not invented", "", payment[:1], webhookEventList(payment...), webhookEventList(payment[0])},
	}
	for _, tc := range tests {
		for _, raw := range []bool{true, false} {
			t.Run(fmt.Sprintf("%s/raw=%t", tc.name, raw), func(t *testing.T) {
				obj := &stripe.WebhookEndpoint{EnabledEvents: tc.observed, APIVersion: tc.version}
				if raw {
					data, err := json.Marshal(map[string]interface{}{"enabled_events": tc.observed, "api_version": tc.version})
					if err != nil {
						t.Fatal(err)
					}
					obj.SetLastResponse(&stripe.APIResponse{RawJSON: data})
				}
				state := WebhookEndpointResourceModel{EnabledEvents: tc.configured}
				if err := flattenWebhookEndpoint(obj, &state, tc.version); err != nil {
					t.Fatal(err)
				}
				if !state.EnabledEvents.Equal(tc.want) {
					t.Fatalf("got %s; want %s", state.EnabledEvents, tc.want)
				}
			})
		}
	}
}
