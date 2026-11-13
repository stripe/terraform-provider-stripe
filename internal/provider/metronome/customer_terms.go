//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"errors"
	"fmt"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
	"strings"
	"time"
	"unicode/utf8"
)

func validateContract(_ context.Context, v *values) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	if err = nonempty(r, "customer_id"); err != nil {
		return err
	}
	if !isUnknownAPI(r["package_id"]) && !isUnknownAPI(r["rate_card_id"]) && (r["package_id"] == nil) == (r["rate_card_id"] == nil) {
		return errors.New("set exactly one of package_id or rate_card_id")
	}
	for _, field := range []string{"package_id", "rate_card_id", "name"} {
		if r[field] != nil {
			if err = nonempty(r, field); err != nil {
				return err
			}
		}
	}
	if r["package_id"] != nil && !isUnknownAPI(r["package_id"]) {
		for _, field := range []string{"name", "ending_before"} {
			if r[field] != nil && !isUnknownAPI(r[field]) {
				return fmt.Errorf("%s cannot be set when creating a package-backed contract", field)
			}
		}
	}
	if !isUnknownAPI(r["uniqueness_key"]) && (utf8.RuneCountInString(str(r["uniqueness_key"])) < 1 || utf8.RuneCountInString(str(r["uniqueness_key"])) > 128) {
		return errors.New("uniqueness_key must be 1–128 characters")
	}
	if !isUnknownAPI(r["starting_at"]) {
		at, err := time.Parse(time.RFC3339Nano, str(r["starting_at"]))
		if err != nil || !at.Equal(at.Truncate(time.Hour)) {
			return errors.New("starting_at must be RFC 3339 aligned to the hour")
		}
	}
	if r["ending_before"] != nil && !isUnknownAPI(r["ending_before"]) {
		end, err := time.Parse(time.RFC3339Nano, str(r["ending_before"]))
		if err != nil {
			return errors.New("ending_before must be RFC 3339")
		}
		if !isUnknownAPI(r["starting_at"]) {
			start, _ := time.Parse(time.RFC3339Nano, str(r["starting_at"]))
			if !start.Before(end) {
				return errors.New("ending_before must follow starting_at")
			}
		}
	}
	return nil
}
func planContract(_ context.Context, v, prior *values) ([]string, error) {
	if !v.Known("uniqueness_key") {
		return nil, errors.New("contract uniqueness_key must be known before replacement")
	}
	fields := []string{"customer_id", "package_id", "rate_card_id", "name", "starting_at"}
	if v.String("uniqueness_key") == prior.String("uniqueness_key") {
		for _, field := range fields {
			if value := v.Get(field); value.IsUnknown() {
				return nil, errors.New("contract replacement cannot be ruled out while retaining the uniqueness key")
			}
		}
	}
	replace := knownChange(v, prior, fields[:4]...)
	startChanged := knownChange(v, prior, "starting_at") && !sameTime(v.String("starting_at"), prior.String("starting_at"))
	replace = replace || startChanged
	if replace && v.String("uniqueness_key") == prior.String("uniqueness_key") {
		return nil, errors.New("contract replacement needs a new uniqueness_key; archived keys remain reserved")
	}
	if startChanged {
		return []string{"starting_at"}, nil
	}
	return nil, nil
}
func contractEndRequest(_ context.Context, v, prior *values, body map[string]any) (any, error) {
	if prior != nil && sameTime(v.String("ending_before"), prior.String("ending_before")) {
		return bindingruntime.SkipRequest{}, nil
	}
	if !v.Known("ending_before") {
		return nil, errors.New("set a known ending_before or use the editContract action to clear it")
	}
	return body, nil
}
func readContract(ctx context.Context, v *values, remote any) error {
	if err := checkIdentity(ctx, v, remote); err != nil {
		return err
	}
	r := obj(remote)
	if r["customer_id"] != v.String("customer_id") {
		return errors.New("contract customer identity mismatch")
	}
	if err := nonempty(r, "uniqueness_key", "starting_at"); err != nil {
		return err
	}
	if emptyAPIValue(r["package_id"]) && emptyAPIValue(r["rate_card_id"]) {
		return errors.New("contract response omitted package and rate-card identity")
	}
	updates := map[string]any{"package_id": nullable(r["package_id"]), "rate_card_id": nullable(r["rate_card_id"]), "name": nullable(r["name"]), "uniqueness_key": r["uniqueness_key"], "starting_at": retainTime(v.String("starting_at"), r["starting_at"]), "ending_before": nullable(retainTime(v.String("ending_before"), r["ending_before"]))}
	if !emptyAPIValue(r["package_id"]) {
		updates["rate_card_id"] = nil
		updates["name"] = nil
	}
	return put(ctx, v, updates)
}
func validateAlert(_ context.Context, v *values) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	if err = nonempty(r, "customer_id", "name"); err != nil {
		return err
	}
	if err = enumValue(r, "alert_type", customerAlertTypeValues, false); err != nil {
		return err
	}
	if r["alert_type"] == "low_remaining_seat_balance_reached" {
		return errors.New("seat balance alerts require an unmodeled seat_filter; use the API action")
	}
	return nil
}
func planAlert(_ context.Context, v, prior *values) ([]string, error) {
	return nil, reservedKey(v, prior, "customer_id", "name", "alert_type", "threshold", "credit_type_id", "billable_metric_id", "evaluate_on_create")
}
func updateAlert(_ context.Context, v, prior *values) error {
	return ensureUnchanged(v, prior, "customer_id", "name", "alert_type", "threshold", "credit_type_id", "billable_metric_id", "uniqueness_key", "evaluate_on_create")
}
func readAlert(ctx context.Context, v *values, remote any) error {
	r := obj(get(remote, "alert"))
	if r == nil {
		return errors.New("alert response omitted alert object")
	}
	if r["status"] == "archived" {
		return bindingruntime.ErrMissing
	}
	if err := checkIdentity(ctx, v, r); err != nil {
		return err
	}
	if err := nonempty(r, "name", "type"); err != nil {
		return err
	}
	if _, ok := rat(r["threshold"]); !ok {
		return errors.New("alert response omitted numeric threshold")
	}
	if r["access_type"] != nil && r["access_type"] != "" && r["access_type"] != "SPEND" {
		return errors.New("alert has unsupported access type")
	}
	for _, key := range []string{"alert_specifiers", "credit_grant_type_filters", "custom_field_filters", "group_key_filter", "group_values", "invoice_types_filter", "seat_filter"} {
		if !emptyAPIValue(r[key]) {
			return errors.New("alert has filters this resource cannot round-trip")
		}
	}
	updates := map[string]any{"name": r["name"], "alert_type": r["type"], "threshold": r["threshold"], "status": r["status"], "credit_type_id": nullable(get(r, "credit_type", "id"))}
	if !v.Known("uniqueness_key") {
		updates["uniqueness_key"] = nullable(r["uniqueness_key"])
	}
	return put(ctx, v, updates)
}
func validateBalance(v *values, kind string) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	if err = nonempty(r, "customer_id", "product_id", "name"); err != nil {
		return err
	}
	if r["uniqueness_key"] != nil {
		if err = nonempty(r, "uniqueness_key"); err != nil {
			return err
		}
	}
	if !isUnknownAPI(r["priority"]) {
		if _, ok := rat(r["priority"]); !ok {
			return errors.New("priority must be numeric")
		}
	}
	if kind == "commit" && !isUnknownAPI(r["commit_type"]) && r["commit_type"] != "PREPAID" {
		return errors.New("customer commits support PREPAID without an invoice schedule")
	}
	if isUnknownAPI(r["access_schedule"]) {
		return nil
	}
	schedule := obj(r["access_schedule"])
	if schedule == nil {
		return errors.New("access_schedule must be an object")
	}
	if !isUnknownAPI(schedule["access_type"]) && schedule["access_type"] != "SPEND" && schedule["access_type"] != "QUANTITY" {
		return errors.New("access_type must be SPEND or QUANTITY")
	}
	if err = nonempty(schedule, "credit_type_id"); err != nil {
		return err
	}
	if isUnknownAPI(schedule["schedule_items"]) {
		return nil
	}
	items := arr(schedule["schedule_items"])
	if len(items) == 0 {
		return errors.New("access schedule needs at least one schedule item")
	}
	for _, item := range items {
		if isUnknownAPI(item) {
			continue
		}
		s := obj(item)
		if !isUnknownAPI(s["amount"]) {
			n, ok := rat(s["amount"])
			if !ok || n.Sign() < 0 {
				return errors.New("schedule amount must be nonnegative")
			}
		}
		if !isUnknownAPI(s["starting_at"]) && !isUnknownAPI(s["ending_before"]) {
			start, e1 := time.Parse(time.RFC3339Nano, str(s["starting_at"]))
			end, e2 := time.Parse(time.RFC3339Nano, str(s["ending_before"]))
			if e1 != nil || e2 != nil || !start.Before(end) {
				return errors.New("schedule item requires RFC 3339 start before end")
			}
		}
	}
	return nil
}
func validateCustomerCredit(_ context.Context, v *values) error { return validateBalance(v, "credit") }
func validateCustomerCommit(_ context.Context, v *values) error { return validateBalance(v, "commit") }
func balanceRequest(_ context.Context, v, prior *values, body map[string]any) (any, error) {
	if strings.TrimSpace(v.String("uniqueness_key")) == "" {
		return nil, errors.New("new customer balances require a known uniqueness_key")
	}
	return body, nil
}
func planCustomerBalance(_ context.Context, v, prior *values) ([]string, error) {
	change := bindingruntime.Changed(v, prior, "customer_id", "product_id", "name", "priority", "access_schedule", "applicable_product_ids", "uniqueness_key") || knownChange(v, prior, "commit_type")
	if change {
		old, err := prior.JSON(true)
		if err != nil {
			return nil, err
		}
		if old["deletion_protection"] != false {
			return nil, errors.New("apply deletion_protection = false before replacing this balance")
		}
		if !v.Known("uniqueness_key") || v.String("uniqueness_key") == prior.String("uniqueness_key") {
			return nil, errors.New("balance replacement needs a new uniqueness_key")
		}
	}
	return nil, nil
}
func readCustomerCredit(ctx context.Context, v *values, remote any) error {
	return readBalance(ctx, v, remote, "credit")
}
func readCustomerCommit(ctx context.Context, v *values, remote any) error {
	return readBalance(ctx, v, remote, "commit")
}
func readBalance(ctx context.Context, v *values, remote any, kind string) error {
	if err := checkIdentity(ctx, v, remote); err != nil {
		return err
	}
	r := obj(remote)
	if err := nonempty(r, "name"); err != nil {
		return err
	}
	if str(get(r, "product", "id")) == "" || str(get(r, "access_schedule", "credit_type", "id")) == "" || len(arr(get(r, "access_schedule", "schedule_items"))) == 0 {
		return errors.New("balance response omitted required product or schedule")
	}
	if _, ok := rat(r["priority"]); !ok {
		return errors.New("balance response omitted numeric priority")
	}
	for _, field := range []string{"description", "applicable_product_tags", "applicable_contract_ids", "specifiers", "custom_fields", "recurring_commit_id", "recurring_credit_id", "rolled_over_from", "hierarchy_configuration", "spend_tracker_attributes", "subscription_config", "netsuite_sales_order_id", "salesforce_opportunity_id"} {
		if !emptyAPIValue(r[field]) {
			return errors.New("balance has unsupported contract, invoice, recurrence or applicability terms")
		}
	}
	if !emptyAPIValue(get(r, "contract", "id")) || !emptyAPIValue(get(r, "invoice_contract", "id")) || !emptyAPIValue(get(r, "invoice_schedule", "schedule_items")) || get(r, "invoice_schedule", "do_not_invoice") == true {
		return errors.New("balance has unsupported contract or invoice schedule")
	}
	for _, field := range []string{"cost_basis", "rollover_fraction"} {
		if !emptyAPIValue(r[field]) {
			n, ok := rat(r[field])
			if !ok || n.Sign() != 0 {
				return errors.New("balance has unsupported cost basis or rollover")
			}
		}
	}
	if (!emptyAPIValue(r["rate_type"]) && r["rate_type"] != "LIST_RATE") || (kind == "commit" && r["type"] != "PREPAID") {
		return errors.New("balance has unsupported rate or commit type")
	}
	old, _ := v.JSON(true)
	oldItems := arr(get(old, "access_schedule", "schedule_items"))
	items := []any{}
	for _, item := range arr(get(r, "access_schedule", "schedule_items")) {
		s := obj(item)
		if _, ok := rat(s["amount"]); !ok {
			return errors.New("balance schedule amount is not numeric")
		}
		next := map[string]any{"amount": s["amount"], "starting_at": s["starting_at"], "ending_before": s["ending_before"]}
		for _, previous := range oldItems {
			if sameTime(str(get(previous, "starting_at")), str(s["starting_at"])) && sameTime(str(get(previous, "ending_before")), str(s["ending_before"])) && numberEqual(get(previous, "amount"), s["amount"]) {
				next["starting_at"] = get(previous, "starting_at")
				next["ending_before"] = get(previous, "ending_before")
				break
			}
		}
		items = append(items, next)
	}
	updates := map[string]any{"product_id": get(r, "product", "id"), "name": r["name"], "priority": r["priority"], "uniqueness_key": nullable(r["uniqueness_key"]), "commit_type": r["type"], "access_schedule": map[string]any{"access_type": get(r, "access_schedule", "access_type"), "credit_type_id": get(r, "access_schedule", "credit_type", "id"), "schedule_items": items}}
	if old["applicable_product_ids"] != nil || len(arr(r["applicable_product_ids"])) > 0 {
		ids := r["applicable_product_ids"]
		if ids == nil {
			ids = []any{}
		}
		updates["applicable_product_ids"] = ids
	}
	return put(ctx, v, updates)
}
