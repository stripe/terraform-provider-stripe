//
// File generated from our OpenAPI spec
//

package combinedprovider

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	original "github.com/stripe/terraform-provider-stripe/internal/provider"
	"github.com/stripe/terraform-provider-stripe/internal/providerdata"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

func configureFixture(t *testing.T, source string, overrides map[string]attr.Value) *providerdata.Data {
	t.Helper()
	ctx := context.Background()
	p := New("qualification")()
	var definition provider.SchemaResponse
	p.Schema(ctx, provider.SchemaRequest{}, &definition)
	object, err := bindingruntime.FromJSON(ctx, definition.Schema.Type(), bindingruntime.Literal(source), nil)
	if err != nil {
		t.Fatal(err)
	}
	values := bindingruntime.NewValues(object.(types.Object))
	for key, value := range overrides {
		if err := values.Set(key, value); err != nil {
			t.Fatal(err)
		}
	}
	raw, err := values.Object.ToTerraformValue(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var response provider.ConfigureResponse
	p.Configure(ctx, provider.ConfigureRequest{Config: tfsdk.Config{Schema: definition.Schema, Raw: raw}}, &response)
	if response.Diagnostics.HasError() {
		t.Fatal("service-specific initialization error leaked from Configure", response.Diagnostics)
	}
	data, ok := response.ResourceData.(*providerdata.Data)
	if !ok {
		t.Fatal("missing service data")
	}
	if response.DataSourceData != data || response.ActionData != data || response.EphemeralResourceData != data {
		t.Fatal("provider instance uses different service data per surface")
	}
	return data
}

func TestCredentialResolutionAndServiceIsolation(t *testing.T) {
	t.Setenv("STRIPE_API_KEY", "sk_test_environment")
	t.Setenv("STRIPE_ACCOUNT", "")
	t.Setenv("METRONOME_API_TOKEN", "ignored-native-token")
	t.Setenv("METRONOME_BEARER_TOKEN", "ignored-native-token")
	t.Setenv("METRONOME_BASE_URL", "")
	t.Setenv("METRONOME_CA_CERT_FILE", "")
	for _, source := range []string{`{}`, `{"api_key":null}`, `{"api_key":""}`} {
		data := configureFixture(t, source, nil)
		if data.StripeKey != "sk_test_environment" || data.StripeStatus.State != "ready" || data.MetronomeStatus.State != "ready" {
			t.Fatalf("shared key fallback failed: Stripe=%s Metronome=%s", data.StripeStatus.State, data.MetronomeStatus.State)
		}
	}
	one := configureFixture(t, `{"api_key":"sk_test_one"}`, nil)
	two := configureFixture(t, `{"api_key":"sk_test_two"}`, nil)
	if one.StripeKey != "sk_test_one" || two.StripeKey != "sk_test_two" || one.Stripe == two.Stripe || one.Metronome == two.Metronome {
		t.Fatal("provider aliases share clients or ignore explicit keys")
	}
	invalid := configureFixture(t, `{"metronome":{"ca_cert_file":"/nonexistent/metronome-fixture-ca"}}`, nil)
	if invalid.StripeStatus.State != "ready" || invalid.MetronomeStatus.State != "invalid" || invalid.Stripe == nil {
		t.Fatal("invalid unused Metronome settings disabled Stripe")
	}
	unknown := configureFixture(t, `{}`, map[string]attr.Value{"api_key": types.StringUnknown()})
	if unknown.StripeStatus.State != "unknown" || unknown.MetronomeStatus.State != "unknown" || unknown.StripeKey != "" || unknown.Metronome != nil {
		t.Fatal("unknown key selected ambient credentials")
	}
	metType := map[string]attr.Type{"base_url": types.StringType, "ca_cert_file": types.StringType}
	unknownTransport := configureFixture(t, `{}`, map[string]attr.Value{"metronome": types.ObjectUnknown(metType)})
	if unknownTransport.StripeStatus.State != "ready" || unknownTransport.MetronomeStatus.State != "unknown" {
		t.Fatal("unknown Metronome transport deferred Stripe")
	}
	unknownField := configureFixture(t, `{}`, map[string]attr.Value{"metronome": types.ObjectValueMust(metType, map[string]attr.Value{"base_url": types.StringUnknown(), "ca_cert_file": types.StringNull()})})
	if unknownField.MetronomeStatus.State != "unknown" {
		t.Fatal("unknown Metronome field selected ambient configuration")
	}
	t.Setenv("STRIPE_API_KEY", "")
	missing := configureFixture(t, `{}`, nil)
	if missing.StripeStatus.State != "missing" || missing.MetronomeStatus.State != "missing" || missing.Metronome != nil {
		t.Fatal("native token substituted for a missing Stripe key")
	}
}

func TestCombinedProviderPreservesEveryStripeProtocolSchema(t *testing.T) {
	ctx := context.Background()
	old := original.New("test")()
	combined := New("test")()
	before, err := providerserver.NewProtocol6(old)().GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatal(err)
	}
	after, err := providerserver.NewProtocol6(combined)().GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range append(before.Diagnostics, after.Diagnostics...) {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Fatal(d)
		}
	}
	for name, schema := range before.ResourceSchemas {
		if !reflect.DeepEqual(schema, after.ResourceSchemas[name]) {
			t.Errorf("Stripe resource schema changed: %s", name)
		}
	}
	for name, schema := range before.ActionSchemas {
		if !reflect.DeepEqual(schema, after.ActionSchemas[name]) {
			t.Errorf("Stripe action schema changed: %s", name)
		}
	}
	for name, schema := range before.EphemeralResourceSchemas {
		if !reflect.DeepEqual(schema, after.EphemeralResourceSchemas[name]) {
			t.Errorf("Stripe ephemeral schema changed: %s", name)
		}
	}
	if len(after.ResourceSchemas)-len(before.ResourceSchemas) != 15 || len(after.DataSourceSchemas)-len(before.DataSourceSchemas) != 59 || len(after.ActionSchemas)-len(before.ActionSchemas) != 73 {
		t.Fatalf("combined surface inventory is wrong: resources=%d reads=%d actions=%d", len(after.ResourceSchemas), len(after.DataSourceSchemas), len(after.ActionSchemas))
	}
	if len(after.Provider.Block.Attributes) != len(before.Provider.Block.Attributes)+1 {
		t.Fatal("provider schema added unexpected shared fields")
	}
	for _, attribute := range before.Provider.Block.Attributes {
		found := false
		for _, candidate := range after.Provider.Block.Attributes {
			if attribute.Name == candidate.Name {
				found = true
				if !reflect.DeepEqual(attribute, candidate) {
					t.Errorf("Stripe provider attribute changed: %s", attribute.Name)
				}
			}
		}
		if !found {
			t.Errorf("missing provider field %s", attribute.Name)
		}
	}
}

func compareMethods(t *testing.T, old, wrapped any) {
	t.Helper()
	a, b := reflect.TypeOf(old), reflect.TypeOf(wrapped)
	if a.NumMethod() != b.NumMethod() {
		t.Fatalf("optional interface set changed: %s has %d methods, %s has %d", a, a.NumMethod(), b, b.NumMethod())
	}
	for i := 0; i < a.NumMethod(); i++ {
		method := a.Method(i)
		candidate, ok := b.MethodByName(method.Name)
		if !ok {
			t.Fatalf("missing method %s", method.Name)
		}
		if method.Type.NumIn() != candidate.Type.NumIn() || method.Type.NumOut() != candidate.Type.NumOut() {
			t.Fatalf("changed method signature %s", method.Name)
		}
		for n := 1; n < method.Type.NumIn(); n++ {
			if method.Type.In(n) != candidate.Type.In(n) {
				t.Fatalf("changed argument %s", method.Name)
			}
		}
		for n := 0; n < method.Type.NumOut(); n++ {
			if method.Type.Out(n) != candidate.Type.Out(n) {
				t.Fatalf("changed return type %s", method.Name)
			}
		}
	}
}
func TestStripeAdaptersPreserveOptionalInterfacesAndClientRouting(t *testing.T) {
	ctx := context.Background()
	old := original.New("test")()
	combined := New("test")()
	data := configureFixture(t, `{"api_key":"sk_test_routing"}`, nil)
	resources := map[string]resource.Resource{}
	for _, factory := range combined.Resources(ctx) {
		r := factory()
		var meta resource.MetadataResponse
		r.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "stripe"}, &meta)
		resources[meta.TypeName] = r
	}
	for _, factory := range old.Resources(ctx) {
		r := factory()
		var meta resource.MetadataResponse
		r.Metadata(ctx, resource.MetadataRequest{ProviderTypeName: "stripe"}, &meta)
		wrapped := resources[meta.TypeName]
		compareMethods(t, r, wrapped)
		var configured resource.ConfigureResponse
		wrapped.(resource.ResourceWithConfigure).Configure(ctx, resource.ConfigureRequest{ProviderData: data}, &configured)
		if configured.Diagnostics.HasError() {
			t.Fatal(configured.Diagnostics)
		}
		embedded := reflect.ValueOf(wrapped).Elem().Field(0).Elem()
		client := embedded.FieldByName("client")
		if !client.IsValid() || client.Pointer() != reflect.ValueOf(data.Stripe).Pointer() {
			t.Fatalf("Stripe client was not routed to %s", meta.TypeName)
		}
	}
	actions := map[string]action.Action{}
	for _, factory := range combined.(provider.ProviderWithActions).Actions(ctx) {
		a := factory()
		var meta action.MetadataResponse
		a.Metadata(ctx, action.MetadataRequest{ProviderTypeName: "stripe"}, &meta)
		actions[meta.TypeName] = a
	}
	for _, factory := range old.(provider.ProviderWithActions).Actions(ctx) {
		a := factory()
		var meta action.MetadataResponse
		a.Metadata(ctx, action.MetadataRequest{ProviderTypeName: "stripe"}, &meta)
		compareMethods(t, a, actions[meta.TypeName])
	}
	ephemerals := map[string]ephemeral.EphemeralResource{}
	for _, factory := range combined.(provider.ProviderWithEphemeralResources).EphemeralResources(ctx) {
		e := factory()
		var meta ephemeral.MetadataResponse
		e.Metadata(ctx, ephemeral.MetadataRequest{ProviderTypeName: "stripe"}, &meta)
		ephemerals[meta.TypeName] = e
	}
	for _, factory := range old.(provider.ProviderWithEphemeralResources).EphemeralResources(ctx) {
		e := factory()
		var meta ephemeral.MetadataResponse
		e.Metadata(ctx, ephemeral.MetadataRequest{ProviderTypeName: "stripe"}, &meta)
		compareMethods(t, e, ephemerals[meta.TypeName])
	}
}
