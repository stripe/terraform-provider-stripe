//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"errors"
	"strings"
	"time"
)

func importBridge(ctx context.Context, v *values, id string) error {
	parts := strings.Split(id, "/")
	if len(parts) != 2 || !strings.HasPrefix(parts[0], "acct_") || len(parts[0]) <= 5 || (parts[1] != "test" && parts[1] != "live") {
		return errors.New("bridge import ID must be acct_<account>/test or acct_<account>/live")
	}
	return put(ctx, v, map[string]any{"stripe_account_id": parts[0], "mode": parts[1]})
}
func validateBridge(_ context.Context, v *values) error {
	if !v.Known("readiness_timeout") {
		return nil
	}
	duration, err := time.ParseDuration(v.String("readiness_timeout"))
	if err != nil || duration < 5*time.Minute+30*time.Second || duration > time.Hour {
		return errors.New("readiness_timeout must be between 5m30s and 1h to allow the five-minute authorization cache to expire")
	}
	return nil
}
func readBridgeIdentity(ctx context.Context, v *values, remote any) error {
	r := obj(remote)
	account, mode := str(r["id"]), str(r["metronome_mode"])
	if account == "" || (mode != "test" && mode != "live") {
		return errors.New("Stripe identity response omitted the account or supported environment")
	}
	id := account + "/" + mode
	if v.Known("id") && v.String("id") != id {
		return errors.New("Stripe key belongs to a different account or mode; use the correct provider alias or a separate bridge")
	}
	return put(ctx, v, map[string]any{"id": id, "stripe_account_id": account, "mode": mode})
}
func activatedBridge(ctx context.Context, v *values, remote any) error {
	if get(remote, "status") != "complete" {
		return errors.New("Stripe activation did not complete")
	}
	client := get(remote, "client_id")
	if client != nil && str(client) == "" {
		return errors.New("Stripe activation returned an invalid client_id")
	}
	return v.SetJSON(ctx, "client_id", client)
}
func updateBridge(ctx context.Context, v, prior *values) error { return validateBridge(ctx, v) }
func checkBridgeProbe(_ context.Context, _ *values, remote any) error {
	r, ok := remote.(map[string]any)
	if !ok {
		return errors.New("Metronome readiness response must be an object")
	}
	items, ok := r["data"].([]any)
	if !ok {
		return errors.New("Metronome readiness response omitted its data array")
	}
	if cursor, exists := r["next_page"]; exists && cursor != nil {
		if _, ok := cursor.(string); !ok {
			return errors.New("Metronome readiness pagination has an invalid shape")
		}
	}
	for _, item := range items {
		row, ok := item.(map[string]any)
		if !ok {
			return errors.New("Metronome readiness item must be an object")
		}
		for _, key := range []string{"id", "name"} {
			if _, ok := row[key].(string); !ok {
				return errors.New("Metronome readiness item omitted a string field")
			}
		}
		if str(row["id"]) == "" {
			return errors.New("Metronome readiness item omitted its identity")
		}
		if _, ok := row["is_currency"].(bool); !ok {
			return errors.New("Metronome readiness item omitted its currency flag")
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
