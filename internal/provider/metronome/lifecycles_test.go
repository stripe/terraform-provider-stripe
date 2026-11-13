//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	metronomeclient "github.com/stripe/terraform-provider-stripe/internal/metronomeclient"
	"github.com/stripe/terraform-provider-stripe/internal/providerdata"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type apiCall struct {
	method, path, query, body, response string
	status                              int
}

func resourceFixture(t *testing.T, factory func() resource.Resource, calls []apiCall) resource.Resource {
	t.Helper()
	index := 0
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if index >= len(calls) {
			t.Errorf("unexpected extra request %s %s", r.Method, r.URL)
			http.Error(w, "unexpected", 400)
			return
		}
		call := calls[index]
		index++
		if r.Header.Get("Authorization") != "Bearer sk_test_fixture" {
			t.Error("shared Stripe bearer key missing")
		}
		if r.Method != call.method || r.URL.Path != call.path || r.URL.RawQuery != call.query {
			t.Errorf("request %d: got %s %s; want %s %s?%s", index, r.Method, r.URL, call.method, call.path, call.query)
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		if call.body != "" {
			actual, err := bindingruntime.DecodeJSON(body)
			if err != nil {
				t.Errorf("request %d JSON: %v", index, err)
			}
			if !reflect.DeepEqual(actual, bindingruntime.Literal(call.body)) {
				t.Errorf("request %d body: got %s; want %s", index, body, call.body)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		if call.status != 0 {
			w.WriteHeader(call.status)
		}
		if call.response != "" {
			_, _ = w.Write([]byte(call.response))
		} else {
			_, _ = w.Write([]byte(`{"data":{}}`))
		}
	}))
	t.Cleanup(server.Close)
	t.Cleanup(func() {
		if index != len(calls) {
			t.Errorf("only %d/%d API calls executed", index, len(calls))
		}
	})
	cert := filepath.Join(t.TempDir(), "ca.pem")
	if err := os.WriteFile(cert, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600); err != nil {
		t.Fatal(err)
	}
	client, err := metronomeclient.New("sk_test_fixture", server.URL, cert, "test")
	if err != nil {
		t.Fatal(err)
	}
	r := factory()
	var configured resource.ConfigureResponse
	r.(resource.ResourceWithConfigure).Configure(context.Background(), resource.ConfigureRequest{ProviderData: &providerdata.Data{Metronome: client, MetronomeStatus: providerdata.Status{State: "ready"}}}, &configured)
	if configured.Diagnostics.HasError() {
		t.Fatal(configured.Diagnostics)
	}
	return r
}
func fixtureState(t *testing.T, r resource.Resource, source string, overrides map[string]attr.Value) tfsdk.State {
	t.Helper()
	ctx := context.Background()
	var schema resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &schema)
	value, err := bindingruntime.FromJSON(ctx, schema.Schema.Type(), bindingruntime.Literal(source), nil)
	if err != nil {
		t.Fatal(err)
	}
	values := bindingruntime.NewValues(value.(types.Object))
	for key, value := range overrides {
		if err := values.Set(key, value); err != nil {
			t.Fatal(err)
		}
	}
	state := tfsdk.State{Schema: schema.Schema}
	if diags := state.Set(ctx, values.Object); diags.HasError() {
		t.Fatal(diags)
	}
	return state
}
func stateValues(t *testing.T, state tfsdk.State) *values {
	t.Helper()
	var result types.Object
	if diags := state.Get(context.Background(), &result); diags.HasError() {
		t.Fatal(diags)
	}
	return bindingruntime.NewValues(result)
}
func createFixture(t *testing.T, r resource.Resource, source string) tfsdk.State {
	t.Helper()
	state := fixtureState(t, r, source, map[string]attr.Value{"id": types.StringUnknown()})
	plan := tfsdk.Plan{Schema: state.Schema, Raw: state.Raw}
	var response resource.CreateResponse
	response.State.Schema = state.Schema
	r.Create(context.Background(), resource.CreateRequest{Plan: plan}, &response)
	if response.Diagnostics.HasError() {
		t.Fatal(response.Diagnostics)
	}
	if !stateValues(t, response.State).Known("id") {
		t.Fatal("create lost identity")
	}
	return response.State
}
func readFixture(t *testing.T, r resource.Resource, state tfsdk.State) tfsdk.State {
	t.Helper()
	response := resource.ReadResponse{State: state}
	r.Read(context.Background(), resource.ReadRequest{State: state}, &response)
	if response.Diagnostics.HasError() {
		t.Fatal(response.Diagnostics)
	}
	return response.State
}
func deleteFixture(t *testing.T, r resource.Resource, state tfsdk.State) {
	t.Helper()
	var response resource.DeleteResponse
	r.Delete(context.Background(), resource.DeleteRequest{State: state}, &response)
	if response.Diagnostics.HasError() {
		t.Fatal(response.Diagnostics)
	}
}

func TestManagedCatalogSDKLifecycles(t *testing.T) {
	tests := []struct {
		name    string
		factory func() resource.Resource
		config  string
		calls   []apiCall
	}{
		{"product", NewMetronomeProductResource, `{"name":"usage","type":"USAGE","billable_metric_id":"metric"}`, []apiCall{
			{method: "POST", path: "/v1/contract-pricing/products/create", body: `{"name":"usage","type":"USAGE","billable_metric_id":"metric"}`, response: `{"data":{"id":"p"}}`},
			{method: "POST", path: "/v1/contract-pricing/products/get", body: `{"id":"p"}`, response: `{"data":{"id":"p","type":"USAGE","current":{"name":"usage","billable_metric_id":"metric"}}}`},
			{method: "POST", path: "/v1/contract-pricing/products/archive", body: `{"product_id":"p"}`},
		}},
		{"metric", NewMetronomeBillableMetricResource, `{"name":"requests","event_type":"request","aggregation_type":"COUNT"}`, []apiCall{
			{method: "POST", path: "/v1/billable-metrics/create", body: `{"name":"requests","aggregation_type":"COUNT","event_type_filter":{"in_values":["request"]}}`, response: `{"data":{"id":"m"}}`},
			{method: "GET", path: "/v1/billable-metrics/m", response: `{"data":{"id":"m","name":"requests","aggregation_type":"count","event_type_filter":{"in_values":["request"]}}}`},
			{method: "POST", path: "/v1/billable-metrics/archive", body: `{"id":"m"}`},
		}},
		{"credit_type", NewMetronomeCreditTypeResource, `{"name":"credits"}`, []apiCall{
			{method: "POST", path: "/v1/credit-types/create", body: `{"name":"credits"}`, response: `{"data":{"id":"unit"}}`},
			{method: "GET", path: "/v1/credit-types/list", query: "limit=100", response: `{"data":[{"id":"unit","name":"credits","is_currency":false}]}`},
			{method: "GET", path: "/v1/credit-types/list", query: "limit=100", response: `{"data":[{"id":"unit","name":"credits","is_currency":false}]}`},
			{method: "POST", path: "/v1/credit-types/archive", body: `{"id":"unit"}`},
		}},
		{"custom_field_key", NewMetronomeCustomFieldKeyResource, `{"entity":"customer","key":"region","enforce_uniqueness":false}`, []apiCall{
			{method: "POST", path: "/v1/customFields/listKeys", body: `{"entities":["customer"]}`, response: `{"data":[]}`},
			{method: "POST", path: "/v1/customFields/addKey", body: `{"entity":"customer","key":"region","enforce_uniqueness":false}`},
			{method: "POST", path: "/v1/customFields/listKeys", body: `{"entities":["customer"]}`, response: `{"data":[{"entity":"customer","key":"region","enforce_uniqueness":false}]}`},
			{method: "POST", path: "/v1/customFields/listKeys", body: `{"entities":["customer"]}`, response: `{"data":[{"entity":"customer","key":"region","enforce_uniqueness":false}]}`},
			{method: "POST", path: "/v1/customFields/listKeys", body: `{"entities":["customer"]}`, response: `{"data":[{"entity":"customer","key":"region","enforce_uniqueness":false}]}`},
			{method: "POST", path: "/v1/customFields/removeKey", body: `{"entity":"customer","key":"region"}`},
		}},
		{"notification", NewMetronomeOffsetNotificationResource, `{"name":"reminder","uniqueness_key":"notice:1","policy":{"type":"contract.start","offset":"P1D"}}`, []apiCall{
			{method: "POST", path: "/v2/notifications/create", body: `{"name":"reminder","uniqueness_key":"notice:1","policy":{"type":"contract.start","offset":"P1D"}}`, response: `{"data":{"id":"n"}}`},
			{method: "POST", path: "/v2/notifications/get", body: `{"id":"n"}`, response: `{"data":{"id":"n","name":"reminder","policy":{"type":"contract.start","offset":"P1D"}}}`},
			{method: "POST", path: "/v2/notifications/archive", body: `{"id":"n"}`},
		}},
		{"contract", NewMetronomeContractResource, `{"customer_id":"customer","rate_card_id":"card","starting_at":"2026-11-01T00:00:00Z","uniqueness_key":"contract:1","void_invoices_on_destroy":false}`, []apiCall{
			{method: "POST", path: "/v1/contracts/create", body: `{"customer_id":"customer","rate_card_id":"card","starting_at":"2026-11-01T00:00:00Z","uniqueness_key":"contract:1"}`, response: `{"data":{"id":"contract"}}`},
			{method: "POST", path: "/v2/contracts/get", body: `{"customer_id":"customer","contract_id":"contract"}`, response: `{"data":{"id":"contract","customer_id":"customer","rate_card_id":"card","starting_at":"2026-11-01T00:00:00.000Z","uniqueness_key":"contract:1"}}`},
			{method: "POST", path: "/v1/contracts/archive", body: `{"customer_id":"customer","contract_id":"contract","void_invoices":false}`},
		}},
		{"alert", NewMetronomeCustomerAlertResource, `{"customer_id":"customer","name":"low","alert_type":"low_remaining_contract_credit_balance_reached","threshold":9007199254740993.125,"evaluate_on_create":false,"deletion_protection":false}`, []apiCall{
			{method: "POST", path: "/v1/alerts/create", body: `{"customer_id":"customer","name":"low","alert_type":"low_remaining_contract_credit_balance_reached","threshold":9007199254740993.125,"evaluate_on_create":false}`, response: `{"data":{"id":"alert"}}`},
			{method: "POST", path: "/v1/customer-alerts/get", body: `{"customer_id":"customer","alert_id":"alert"}`, response: `{"data":{"alert":{"id":"alert","name":"low","type":"low_remaining_contract_credit_balance_reached","threshold":9007199254740993.125,"status":"enabled"}}}`},
			{method: "POST", path: "/v1/alerts/archive", body: `{"id":"alert"}`},
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := resourceFixture(t, test.factory, test.calls)
			state := createFixture(t, r, test.config)
			state = readFixture(t, r, state)
			deleteFixture(t, r, state)
		})
	}
}

func TestCustomerBalanceSDKLifecycles(t *testing.T) {
	for _, kind := range []string{"credit", "commit"} {
		t.Run(kind, func(t *testing.T) {
			title, typ, factory := "Credit", "CREDIT", NewMetronomeCustomerCreditResource
			if kind == "commit" {
				title, typ, factory = "Commit", "PREPAID", NewMetronomeCustomerCommitResource
			}
			schedule := `{"access_type":"SPEND","credit_type_id":"unit","schedule_items":[{"amount":9007199254740993.125,"starting_at":"2026-11-01T00:00:00Z","ending_before":"2026-12-01T00:00:00Z"}]}`
			config := fmt.Sprintf(`{"customer_id":"customer","product_id":"product","name":"trial","priority":100.25,"access_schedule":%s,"uniqueness_key":"balance:1","commit_type":%q,"deletion_protection":false}`, schedule, typ)
			body := fmt.Sprintf(`{"customer_id":"customer","product_id":"product","name":"trial","priority":100.25,"access_schedule":%s,"uniqueness_key":"balance:1"%s}`, schedule, map[bool]string{true: `,"type":"PREPAID"`, false: ""}[kind == "commit"])
			remoteSchedule := strings.Replace(schedule, `"credit_type_id":"unit"`, `"credit_type":{"id":"unit"}`, 1)
			r := resourceFixture(t, factory, []apiCall{
				{method: "POST", path: "/v1/contracts/customer" + title + "s/create", body: body, response: `{"data":{"id":"balance"}}`},
				{method: "POST", path: "/v1/contracts/customer" + title + "s/list", body: fmt.Sprintf(`{"customer_id":"customer",%q:"balance","include_archived":false,"limit":25}`, kind+"_id"), response: fmt.Sprintf(`{"data":[{"id":"balance","type":%q,"product":{"id":"product"},"name":"trial","priority":100.25,"access_schedule":%s,"uniqueness_key":"balance:1","rate_type":"LIST_RATE"}]}`, typ, remoteSchedule)},
				{method: "GET", path: "/v1/customers/customer", response: `{"data":{"id":"customer","name":"customer"}}`},
				{method: "POST", path: "/v2/contracts/" + kind + "s/archive", body: fmt.Sprintf(`{"customer_id":"customer",%q:"balance"}`, kind+"_id")},
			})
			state := createFixture(t, r, config)
			state = readFixture(t, r, state)
			deleteFixture(t, r, state)
		})
	}
}

func TestCustomerProviderConfigurationSDKLifecycles(t *testing.T) {
	for _, kind := range []string{"billing", "revenue"} {
		t.Run(kind, func(t *testing.T) {
			title, provider, field, factory := "BillingProvider", "stripe", "billing_provider", NewMetronomeCustomerBillingConfigurationResource
			if kind == "revenue" {
				title, provider, field, factory = "RevenueSystem", "netsuite", "provider", NewMetronomeCustomerRevenueConfigurationResource
			}
			config := fmt.Sprintf(`{"customer_id":"customer","provider_name":%q,"configuration":{"external_customer_id":"external"},"delivery_method_id":"delivery","deletion_protection":false}`, provider)
			remote := fmt.Sprintf(`{"data":[{"id":"configuration","customer_id":"customer",%q:%q,"configuration":{"external_customer_id":"external","server_default":true},"delivery_method_id":"delivery"}]}`, field, provider)
			r := resourceFixture(t, factory, []apiCall{
				{method: "POST", path: "/v1/getCustomer" + title + "Configurations", body: `{"customer_id":"customer","include_archived":true}`, response: `{"data":[]}`},
				{method: "POST", path: "/v1/setCustomer" + title + "Configurations", body: fmt.Sprintf(`{"data":[{"customer_id":"customer",%q:%q,"configuration":{"external_customer_id":"external"},"delivery_method_id":"delivery"}]}`, field, provider), response: `{"data":[{"id":"configuration"}]}`},
				{method: "POST", path: "/v1/getCustomer" + title + "Configurations", body: `{"customer_id":"customer","include_archived":true}`, response: remote},
				{method: "POST", path: "/v1/archiveCustomer" + title + "Configurations", body: fmt.Sprintf(`{"customer_id":"customer",%q:["configuration"]}`, "customer_"+map[string]string{"billing": "billing_provider", "revenue": "revenue_system"}[kind]+"_configuration_ids")},
			})
			state := createFixture(t, r, config)
			original := stateValues(t, state).Get("configuration")
			state = readFixture(t, r, state)
			if !stateValues(t, state).Get("configuration").Equal(original) {
				t.Fatal("server defaults changed configured dynamic values")
			}
			deleteFixture(t, r, state)
		})
	}
}

func TestCustomerPartialUpdateAndArchiveRecovery(t *testing.T) {
	ctx := context.Background()
	r := resourceFixture(t, NewMetronomeCustomerResource, []apiCall{
		{method: "GET", path: "/v1/customers", query: "ingest_alias=new", response: `{"data":[]}`},
		{method: "POST", path: "/v1/customers/customer/setName", body: `{"name":"new"}`},
		{method: "POST", path: "/v1/customers/customer/setIngestAliases", body: `{"ingest_aliases":["new"]}`, status: 500},
		{method: "GET", path: "/v1/customers/customer", response: `{"data":{"id":"customer","name":"new","ingest_aliases":["old"]}}`},
		{method: "POST", path: "/v2/contracts/list", body: `{"customer_id":"customer","include_archived":false,"limit":20}`, response: `{"data":[],"cursor":null}`},
		{method: "POST", path: "/v1/customers/customer/setIngestAliases", body: `{"ingest_aliases":[]}`},
		{method: "POST", path: "/v1/customers/archive", body: `{"id":"customer"}`, status: 500},
		{method: "GET", path: "/v1/customers/customer", response: `{"data":{"id":"customer","name":"new","ingest_aliases":[]}}`},
		{method: "POST", path: "/v1/customers/customer/setIngestAliases", body: `{"ingest_aliases":["old"]}`},
	})
	prior := fixtureState(t, r, `{"id":"customer","name":"old","ingest_aliases":["old"],"deletion_protection":false}`, nil)
	planned := fixtureState(t, r, `{"id":"customer","name":"new","ingest_aliases":["new"],"deletion_protection":false}`, nil)
	result := resource.UpdateResponse{State: prior}
	r.Update(ctx, resource.UpdateRequest{State: prior, Plan: tfsdk.Plan{Schema: planned.Schema, Raw: planned.Raw}}, &result)
	if !result.Diagnostics.HasError() {
		t.Fatal("alias failure was ignored")
	}
	actual := stateValues(t, result.State)
	if actual.String("name") != "new" || !actual.Get("ingest_aliases").Equal(stateValues(t, prior).Get("ingest_aliases")) {
		t.Fatal("partial state did not reflect exactly the applied changes")
	}
	var deleted resource.DeleteResponse
	r.Delete(ctx, resource.DeleteRequest{State: result.State}, &deleted)
	if !deleted.Diagnostics.HasError() {
		t.Fatal("archive failure must remain an error after alias restoration")
	}
}

func TestRateCardFailureRetainsIdentityUnlessCleanupSucceeds(t *testing.T) {
	for _, cleanupStatus := range []int{200, 400} {
		t.Run(fmt.Sprint(cleanupStatus), func(t *testing.T) {
			r := resourceFixture(t, NewMetronomeRateCardResource, []apiCall{
				{method: "POST", path: "/v1/contract-pricing/rate-cards/create", body: `{"name":"card","fiat_credit_type_id":"usd"}`, response: `{"data":{"id":"card"}}`},
				{method: "POST", path: "/v1/contract-pricing/rate-cards/addRate", status: 400},
				{method: "POST", path: "/v1/contract-pricing/rate-cards/archive", body: `{"id":"card"}`, status: cleanupStatus},
			})
			state := fixtureState(t, r, `{"name":"card","credit_type_id":"usd","usage_product_id":"product","price_cents":1,"starting_at":"2026-11-01T00:00:00.000Z"}`, map[string]attr.Value{"id": types.StringUnknown()})
			response := resource.CreateResponse{State: tfsdk.State{Schema: state.Schema}}
			r.Create(context.Background(), resource.CreateRequest{Plan: tfsdk.Plan{Schema: state.Schema, Raw: state.Raw}}, &response)
			if !response.Diagnostics.HasError() {
				t.Fatal("failed rate creation reported success")
			}
			if cleanupStatus == 200 && !response.State.Raw.IsNull() {
				t.Fatal("successfully archived card remained in state")
			}
			if cleanupStatus != 200 && stateValues(t, response.State).String("id") != "card" {
				t.Fatal("failed cleanup lost the recovery identity")
			}
		})
	}
}

func TestReadFailuresDoNotEraseOrdinaryResourceState(t *testing.T) {
	for _, status := range []int{401, 403, 404} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			r := resourceFixture(t, NewMetronomeProductResource, []apiCall{{method: "POST", path: "/v1/contract-pricing/products/get", body: `{"id":"product"}`, status: status}})
			state := fixtureState(t, r, `{"id":"product","name":"fixed","type":"FIXED"}`, nil)
			response := resource.ReadResponse{State: state}
			r.Read(context.Background(), resource.ReadRequest{State: state}, &response)
			if status == 404 {
				if response.Diagnostics.HasError() || !response.State.Raw.IsNull() {
					t.Fatal("missing product was retained")
				}
			} else if !response.Diagnostics.HasError() || !response.State.Raw.Equal(state.Raw) {
				t.Fatal("permission failure erased an ordinary resource")
			}
		})
	}
}

func TestPackageRepresentationChangeRetainsIdentity(t *testing.T) {
	ctx := context.Background()
	r := NewMetronomePackageResource()
	payload := `{"name":"package","contract_name":"contract","rate_card_id":"card","usage_statement_schedule":{"frequency":"MONTHLY","day":"CONTRACT_START"}}`
	prior := fixtureState(t, r, `{"id":"package","request":`+payload+`}`, nil)
	planned := fixtureState(t, r, `{"package":`+payload+`}`, map[string]attr.Value{"id": types.StringUnknown()})
	plan := tfsdk.Plan{Schema: planned.Schema, Raw: planned.Raw}
	response := resource.ModifyPlanResponse{Plan: plan}
	r.(resource.ResourceWithModifyPlan).ModifyPlan(ctx, resource.ModifyPlanRequest{State: prior, Plan: plan}, &response)
	if response.Diagnostics.HasError() || len(response.RequiresReplace) != 0 {
		t.Fatal("equivalent typed representation requested replacement", response.Diagnostics)
	}
	var value types.String
	if diags := response.Plan.GetAttribute(ctx, path.Root("id"), &value); diags.HasError() {
		t.Fatal(diags)
	}
	if value.ValueString() != "package" {
		t.Fatal("equivalent representation lost package identity")
	}
}
