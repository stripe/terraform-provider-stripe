//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"errors"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

func validateBillingConfiguration(_ context.Context, v *values) error {
	return validateConfiguration(v, "billing")
}
func validateRevenueConfiguration(_ context.Context, v *values) error {
	return validateConfiguration(v, "revenue")
}
func validateConfiguration(v *values, kind string) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	if err = nonempty(r, "customer_id"); err != nil {
		return err
	}
	allowed := billingProviderValues
	if kind == "revenue" {
		allowed = revenueProviderValues
	}
	if err = enumValue(r, "provider_name", allowed, false); err != nil {
		return err
	}
	if err = enumValue(r, "delivery_method", deliveryMethodValues, true); err != nil {
		return err
	}
	if err = enumValue(r, "tax_provider", taxProviderValues, true); err != nil {
		return err
	}
	if r["configuration"] != nil && !isUnknownAPI(r["configuration"]) {
		if obj(r["configuration"]) == nil {
			return errors.New("configuration must be an object")
		}
	}
	return nil
}
func wrapConfigurationRequest(_ context.Context, v, prior *values, body map[string]any) (any, error) {
	if obj(body["configuration"]) == nil {
		return nil, errors.New("configuration must be a known object")
	}
	if body["billing_provider"] != nil && body["delivery_method"] == nil && body["delivery_method_id"] == nil {
		return nil, errors.New("billing configuration requires delivery_method or delivery_method_id")
	}
	return map[string]any{"data": []any{body}}, nil
}
func checkBillingConfigurationAbsent(_ context.Context, v *values, remote any) error {
	return checkConfigurationAbsent(v, remote, "billing_provider")
}
func checkRevenueConfigurationAbsent(_ context.Context, v *values, remote any) error {
	return checkConfigurationAbsent(v, remote, "provider")
}
func checkConfigurationAbsent(v *values, remote any, provider string) error {
	items, ok := remote.([]any)
	if !ok {
		return errors.New("configuration list is missing")
	}
	for _, item := range items {
		methodMatches := (v.Known("delivery_method_id") && get(item, "delivery_method_id") == v.String("delivery_method_id")) || (v.Known("delivery_method") && get(item, "delivery_method") == v.String("delivery_method"))
		if emptyAPIValue(get(item, "archived_at")) && get(item, provider) == v.String("provider_name") && methodMatches {
			return errors.New("provider configuration already exists; import it instead of taking ownership")
		}
	}
	return nil
}
func createdConfiguration(ctx context.Context, v *values, remote any) error {
	items, ok := remote.([]any)
	if !ok || len(items) != 1 || str(get(items[0], "id")) == "" {
		return errors.New("API did not return exactly one configuration ID")
	}
	if err := v.SetJSON(ctx, "id", get(items[0], "id")); err != nil {
		return err
	}
	if !v.Known("delivery_method_id") && !emptyAPIValue(get(items[0], "delivery_method_id")) {
		return v.SetJSON(ctx, "delivery_method_id", get(items[0], "delivery_method_id"))
	}
	return nil
}
func readBillingConfiguration(ctx context.Context, v *values, remote any) error {
	return readConfiguration(ctx, v, remote, "billing_provider")
}
func readRevenueConfiguration(ctx context.Context, v *values, remote any) error {
	return readConfiguration(ctx, v, remote, "provider")
}
func readConfiguration(ctx context.Context, v *values, remote any, provider string) error {
	if err := checkIdentity(ctx, v, remote); err != nil {
		return err
	}
	r := obj(remote)
	if r["customer_id"] != v.String("customer_id") || str(r[provider]) == "" {
		return errors.New("API omitted configuration customer/provider identity")
	}
	if provider == "billing_provider" && !emptyAPIValue(r["unbillable_invoices_configuration"]) {
		return errors.New("configuration contains unsupported unbillable invoice terms")
	}
	if obj(r["configuration"]) == nil {
		return errors.New("API omitted configuration object")
	}
	updates := map[string]any{"provider_name": r[provider], "delivery_method_id": nullable(r["delivery_method_id"]), "delivery_method": nullable(r["delivery_method"])}
	if !emptyAPIValue(r["tax_provider"]) {
		updates["tax_provider"] = r["tax_provider"]
	}
	old, err := bindingruntime.JSONValue(v.Get("configuration"), true)
	if err != nil {
		return err
	}
	if !subset(bindingruntime.StripNulls(old), r["configuration"]) {
		updates["configuration"] = r["configuration"]
	}
	return put(ctx, v, updates)
}
func planConfiguration(_ context.Context, v, prior *values) ([]string, error) {
	wanted, err := bindingruntime.JSONValue(v.Get("configuration"), true)
	if err != nil {
		return nil, err
	}
	actual, err := bindingruntime.JSONValue(prior.Get("configuration"), true)
	if err != nil {
		return nil, err
	}
	if !isUnknownAPI(wanted) && !subset(bindingruntime.StripNulls(wanted), actual) {
		return []string{"configuration"}, nil
	}
	return nil, nil
}
func updateConfiguration(_ context.Context, v, prior *values) error {
	if err := ensureUnchanged(v, prior, "customer_id", "provider_name", "delivery_method_id", "delivery_method", "tax_provider"); err != nil {
		return err
	}
	wanted, err := bindingruntime.JSONValue(v.Get("configuration"), false)
	if err != nil {
		return err
	}
	actual, err := bindingruntime.JSONValue(prior.Get("configuration"), false)
	if err != nil {
		return err
	}
	if !subset(bindingruntime.StripNulls(wanted), actual) {
		return errors.New("configuration cannot be edited in place")
	}
	return nil
}
