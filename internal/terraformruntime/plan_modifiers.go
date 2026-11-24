//
// File generated from our OpenAPI spec
//

package bindingruntime

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/dynamicplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"reflect"
	"time"
)

func RequiresReplaceInstant() planmodifier.String {
	return stringplanmodifier.RequiresReplaceIf(func(_ context.Context, req planmodifier.StringRequest, resp *stringplanmodifier.RequiresReplaceIfFuncResponse) {
		if req.PlanValue.IsUnknown() {
			return
		}
		a, ae := time.Parse(time.RFC3339Nano, req.StateValue.ValueString())
		b, be := time.Parse(time.RFC3339Nano, req.PlanValue.ValueString())
		resp.RequiresReplace = ae != nil || be != nil || !a.Equal(b)
	}, "Replace when the starting instant changes", "Replace when the starting instant changes")
}
func RequiresReplaceSubset() planmodifier.Dynamic {
	return dynamicplanmodifier.RequiresReplaceIf(func(_ context.Context, req planmodifier.DynamicRequest, resp *dynamicplanmodifier.RequiresReplaceIfFuncResponse) {
		wanted, we := JSONValue(req.PlanValue, true)
		actual, ae := JSONValue(req.StateValue, true)
		resp.RequiresReplace = we != nil || ae != nil || !jsonSubset(StripNulls(wanted), actual)
	}, "Replace when configured provider fields change", "Replace when configured provider fields change")
}
func jsonSubset(expected, actual any) bool {
	if IsUnknown(expected) {
		return false
	}
	switch expected := expected.(type) {
	case map[string]any:
		remote, ok := actual.(map[string]any)
		if !ok {
			return false
		}
		for key, x := range expected {
			if !jsonSubset(x, remote[key]) {
				return false
			}
		}
		return true
	case []any:
		remote, ok := actual.([]any)
		if !ok || len(expected) != len(remote) {
			return false
		}
		for i, x := range expected {
			if !jsonSubset(x, remote[i]) {
				return false
			}
		}
		return true
	default:
		return reflect.DeepEqual(expected, actual)
	}
}
