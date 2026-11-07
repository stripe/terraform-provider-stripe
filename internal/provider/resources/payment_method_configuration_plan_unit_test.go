package resources

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestPaymentMethodConfigurationModifyPlan(t *testing.T) {
	ctx := context.Background()
	r := &PaymentMethodConfigurationResource{}
	var schemaResponse resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &schemaResponse)
	schema := schemaResponse.Schema
	objectType := schema.Type().(types.ObjectType)
	null := tftypes.NewValue(objectType.TerraformType(ctx), nil)

	// Use the schema to cover every method with dependent response fields, so new
	// generated payment methods get the same regression coverage automatically.
	for name, attributeType := range objectType.AttrTypes {
		methodType, ok := attributeType.(types.ObjectType)
		if !ok || methodType.AttrTypes["available"] == nil {
			continue
		}
		t.Run(name, func(t *testing.T) {
			for _, tc := range []struct {
				name           string
				before         string
				after          types.String
				unknownDisplay bool
				unknownMethod  bool
				create         bool
				destroy        bool
				wantUnknown    bool
			}{
				{name: "enable", before: "off", after: types.StringValue("on"), wantUnknown: true},
				{name: "disable", before: "on", after: types.StringValue("off"), wantUnknown: true},
				{name: "unchanged_on", before: "on", after: types.StringValue("on")},
				{name: "unchanged_off", before: "off", after: types.StringValue("off")},
				{name: "inherit", before: "on", after: types.StringValue("none"), wantUnknown: true},
				{name: "unknown_preference", before: "off", after: types.StringUnknown(), wantUnknown: true},
				{name: "unknown_display", before: "off", unknownDisplay: true, wantUnknown: true},
				{name: "unknown_method", before: "off", unknownMethod: true, wantUnknown: true},
				{name: "create", before: "off", after: types.StringValue("on"), create: true, wantUnknown: true},
				{name: "destroy", before: "off", destroy: true},
			} {
				t.Run(tc.name, func(t *testing.T) {
					state := tfsdk.State{Schema: schema, Raw: null}
					for _, setting := range []struct {
						p path.Path
						v any
					}{
						{path.Root(name).AtName("available"), tc.before == "on"},
						{path.Root(name).AtName("display_preference").AtName("preference"), tc.before},
						{path.Root(name).AtName("display_preference").AtName("value"), tc.before},
					} {
						if d := state.SetAttribute(ctx, setting.p, setting.v); d.HasError() {
							t.Fatal(d)
						}
					}
					plan := tfsdk.Plan{Schema: schema, Raw: state.Raw.Copy()}
					if d := plan.SetAttribute(ctx, path.Root(name).AtName("display_preference").AtName("preference"), tc.after); d.HasError() {
						t.Fatal(d)
					}
					if tc.unknownDisplay {
						displayType := methodType.AttrTypes["display_preference"].(types.ObjectType)
						if d := plan.SetAttribute(ctx, path.Root(name).AtName("display_preference"), types.ObjectUnknown(displayType.AttrTypes)); d.HasError() {
							t.Fatal(d)
						}
					}
					if tc.unknownMethod {
						if d := plan.SetAttribute(ctx, path.Root(name), types.ObjectUnknown(methodType.AttrTypes)); d.HasError() {
							t.Fatal(d)
						}
					}
					if tc.create {
						state.Raw = null
						if d := plan.SetAttribute(ctx, path.Root(name).AtName("available"), types.BoolUnknown()); d.HasError() {
							t.Fatal(d)
						}
						if d := plan.SetAttribute(ctx, path.Root(name).AtName("display_preference").AtName("value"), types.StringUnknown()); d.HasError() {
							t.Fatal(d)
						}
					}
					if tc.destroy {
						plan.Raw = null
					}
					resp := resource.ModifyPlanResponse{Plan: plan}
					r.ModifyPlan(ctx, resource.ModifyPlanRequest{State: state, Plan: plan}, &resp)
					if resp.Diagnostics.HasError() {
						t.Fatal(resp.Diagnostics)
					}
					if tc.destroy {
						if !resp.Plan.Raw.IsNull() {
							t.Fatal("destroy plan changed")
						}
						return
					}
					if tc.unknownMethod {
						var method types.Object
						if d := resp.Plan.GetAttribute(ctx, path.Root(name), &method); d.HasError() {
							t.Fatal(d)
						}
						if !method.IsUnknown() {
							t.Fatal("unknown method changed")
						}
						return
					}
					var available types.Bool
					if tc.unknownDisplay {
						var display types.Object
						if d := resp.Plan.GetAttribute(ctx, path.Root(name).AtName("display_preference"), &display); d.HasError() {
							t.Fatal(d)
						}
						if !display.IsUnknown() {
							t.Fatal("unknown display changed")
						}
						if d := resp.Plan.GetAttribute(ctx, path.Root(name).AtName("available"), &available); d.HasError() {
							t.Fatal(d)
						}
						if !available.IsUnknown() {
							t.Fatal("availability must be unknown")
						}
						return
					}
					var effective types.String
					if d := resp.Plan.GetAttribute(ctx, path.Root(name).AtName("available"), &available); d.HasError() {
						t.Fatal(d)
					}
					if d := resp.Plan.GetAttribute(ctx, path.Root(name).AtName("display_preference").AtName("value"), &effective); d.HasError() {
						t.Fatal(d)
					}
					for field, value := range map[string]attr.Value{"available": available, "value": effective} {
						if value.IsUnknown() != tc.wantUnknown {
							t.Errorf("%s = %s, want unknown %t", field, value, tc.wantUnknown)
						}
					}
					if !tc.wantUnknown && !resp.Plan.Raw.Equal(plan.Raw) {
						t.Fatal("unchanged preference modified plan")
					}
				})
			}
		})
	}
}
