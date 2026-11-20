//
// File generated from our OpenAPI spec
//

package datasources

import (
	"context"
	"errors"

	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type values = bindingruntime.Values

func get(value any, path ...string) any {
	result, _ := bindingruntime.At(value, path...)
	return result
}
func str(value any) string { result, _ := value.(string); return result }
func nullable(value any) any {
	if value == "" {
		return nil
	}
	return value
}
func put(ctx context.Context, v *values, fields map[string]any) error {
	for key, value := range fields {
		if err := v.SetJSON(ctx, key, value); err != nil {
			return err
		}
	}
	return nil
}

func readBillableMetrics(ctx context.Context, v *values, remote any) error {
	items, ok := remote.([]any)
	if !ok {
		return errors.New("metric list missing")
	}
	metrics := make([]any, 0, len(items))
	for _, item := range items {
		metrics = append(metrics, map[string]any{"id": get(item, "id"), "name": get(item, "name"), "aggregation_type": nullable(get(item, "aggregation_type")), "archived_at": nullable(get(item, "archived_at"))})
	}
	return put(ctx, v, map[string]any{"id": "all", "metrics": metrics})
}
func readStripeBillingProvider(ctx context.Context, v *values, remote any) error {
	items, ok := remote.([]any)
	if !ok {
		return errors.New("billing provider list missing")
	}
	var matches []any
	for _, item := range items {
		if get(item, "billing_provider") == "stripe" && get(item, "delivery_method") == "direct_to_billing_provider" {
			matches = append(matches, item)
		}
	}
	if len(matches) != 1 {
		return errors.New("expected exactly one direct Stripe billing integration; configure the intended connection in Metronome")
	}
	if v.Known("expected_account_id") {
		actual := str(get(matches[0], "delivery_method_configuration", "stripe_account_id"))
		if actual == "" {
			return errors.New("Metronome did not expose the Stripe account ID for verification")
		}
		if actual != v.String("expected_account_id") {
			return errors.New("Metronome billing integration uses a different Stripe account")
		}
	}
	if str(get(matches[0], "delivery_method_id")) == "" {
		return errors.New("billing provider response omitted delivery method identity")
	}
	return v.SetJSON(ctx, "id", get(matches[0], "delivery_method_id"))
}
