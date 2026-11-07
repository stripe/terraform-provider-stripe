package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithModifyPlan = &PaymentMethodConfigurationResource{}

// ModifyPlan runs after the schema's UseStateForUnknown modifiers. The effective
// preference and availability depend on the requested preference, so their old
// values are only valid predictions when that preference is unchanged.
func (r *PaymentMethodConfigurationResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	var planned types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planned)...)
	if resp.Diagnostics.HasError() {
		return
	}
	for name, value := range planned.Attributes() {
		method, ok := value.(types.Object)
		if !ok || method.IsNull() || method.IsUnknown() {
			continue
		}
		display, ok := method.Attributes()["display_preference"].(types.Object)
		if !ok || display.IsNull() {
			continue
		}
		// Select the shared payment-method response shape, excluding write-only
		// methods which do not expose these computed fields.
		if _, ok := method.Attributes()["available"]; !ok {
			continue
		}
		preferencePath := path.Root(name).AtName("display_preference").AtName("preference")
		var prior types.String
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, preferencePath, &prior)...)
		if resp.Diagnostics.HasError() {
			return
		}
		preference := display.Attributes()["preference"]
		if !display.IsUnknown() && !preference.IsUnknown() && preference.Equal(prior) {
			continue
		}
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root(name).AtName("available"), types.BoolUnknown())...)
		if !display.IsUnknown() {
			resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root(name).AtName("display_preference").AtName("value"), types.StringUnknown())...)
		}
	}
}
