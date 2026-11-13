//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"encoding/json"
	"errors"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
	"math/big"
	"strings"
)

func packagePayload(v *values, unknown bool) (any, bool, error) {
	r, err := v.JSON(true)
	if err != nil {
		return nil, false, err
	}
	for _, key := range []string{"package", "request", "request_json"} {
		if r[key] == nil {
			continue
		}
		if isUnknownAPI(r[key]) {
			if unknown {
				return r[key], true, nil
			}
			return nil, true, errors.New("package request must be known at apply")
		}
		switch key {
		case "package":
			return bindingruntime.StripNulls(r[key]), true, nil
		case "request":
			return r[key], true, nil
		case "request_json":
			value, err := decodeJSON([]byte(str(r[key])))
			return value, true, err
		}
	}
	body := map[string]any{"name": r["name"], "contract_name": r["contract_name"], "rate_card_id": r["rate_card_id"], "usage_statement_schedule": map[string]any{"frequency": "MONTHLY", "day": "CONTRACT_START"}}
	if r["uniqueness_key"] != nil {
		if isUnknownAPI(r["uniqueness_key"]) {
			body["uniqueness_key"] = r["uniqueness_key"]
		} else {
			body["uniqueness_key"] = packageKeyWithoutNamespace(str(r["uniqueness_key"]))
		}
	}
	if r["stripe_billing"] == true {
		body["billing_provider"] = "stripe"
		body["delivery_method"] = "direct_to_billing_provider"
	}
	if r["recurring_commit"] != nil {
		if isUnknownAPI(r["recurring_commit"]) {
			body["recurring_commits"] = r["recurring_commit"]
		} else {
			c := obj(r["recurring_commit"])
			body["recurring_commits"] = []any{map[string]any{"name": c["name"], "product_id": c["fixed_product_id"], "priority": c["priority"], "applicable_product_ids": c["applicable_product_ids"], "access_amount": map[string]any{"credit_type_id": c["credit_type_id"], "unit_price": c["monthly_grant_cents"], "quantity": json.Number("1")}, "invoice_amount": map[string]any{"credit_type_id": c["credit_type_id"], "unit_price": c["monthly_price_cents"], "quantity": json.Number("1")}, "starting_at_offset": map[string]any{"unit": "DAYS", "value": json.Number("0")}, "commit_duration": map[string]any{"unit": "PERIODS", "value": c["grant_duration_periods"]}, "recurrence_frequency": "MONTHLY", "proration": "NONE", "rollover_fraction": c["rollover_fraction"]}}
		}
	}
	return body, false, nil
}
func validatePackage(_ context.Context, v *values) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	modes := 0
	for _, key := range []string{"package", "request", "request_json"} {
		if r[key] != nil {
			modes++
		}
	}
	if modes > 1 {
		return errors.New("use exactly one of package, request, or request_json")
	}
	if modes == 1 {
		for _, key := range []string{"name", "contract_name", "rate_card_id", "uniqueness_key", "stripe_billing", "recurring_commit"} {
			if r[key] != nil && !isUnknownAPI(r[key]) {
				return errors.New("API-shaped package requests cannot be combined with legacy fields")
			}
		}
		payload, _, err := packagePayload(v, true)
		if err != nil {
			return err
		}
		return validateKnownPackageValue(payload)
	}
	if err = nonempty(r, "name", "contract_name", "rate_card_id"); err != nil {
		return err
	}
	if r["stripe_billing"] == nil {
		return errors.New("legacy package requires stripe_billing")
	}
	if key, ok := r["uniqueness_key"].(string); ok && (!strings.HasPrefix(key, "package:") || strings.TrimSpace(packageKeyWithoutNamespace(key)) == "") {
		return errors.New("legacy uniqueness_key requires package: and a nonempty key")
	}
	c := obj(r["recurring_commit"])
	if c == nil {
		return nil
	}
	if r["stripe_billing"] == false {
		return errors.New("legacy recurring commits require stripe_billing=true")
	}
	for _, key := range []string{"monthly_price_cents", "monthly_grant_cents", "grant_duration_periods", "priority", "rollover_fraction"} {
		if isUnknownAPI(c[key]) {
			continue
		}
		n, ok := rat(c[key])
		if !ok || n.Sign() < 0 {
			return errors.New("recurring commit values must be numeric and nonnegative")
		}
		if (key == "monthly_grant_cents" || key == "grant_duration_periods") && n.Sign() == 0 {
			return errors.New("grant and duration must be positive")
		}
		if key == "rollover_fraction" && n.Cmp(big.NewRat(1, 1)) > 0 {
			return errors.New("rollover_fraction must be at most 1")
		}
	}
	if !isUnknownAPI(c["applicable_product_ids"]) && len(arr(c["applicable_product_ids"])) == 0 {
		return errors.New("applicable_product_ids must not be empty")
	}
	return nil
}
func packageRequest(ctx context.Context, v, prior *values, _ map[string]any) (any, error) {
	if err := validatePackage(ctx, v); err != nil {
		return nil, err
	}
	payload, _, err := packagePayload(v, false)
	return payload, err
}
func samePayload(a, b any) bool {
	left, le := json.Marshal(a)
	right, re := json.Marshal(b)
	if le != nil || re != nil {
		return false
	}
	x, xe := decodeJSON(left)
	y, ye := decodeJSON(right)
	return xe == nil && ye == nil && packageRequestMatches(x, y, "request") == nil && packageRequestMatches(y, x, "request") == nil
}
func planPackage(_ context.Context, v, prior *values) ([]string, error) {
	changed := false
	for key := range v.Object.Attributes() {
		if key != "id" && bindingruntime.Changed(v, prior, key) {
			changed = true
		}
	}
	if !changed {
		return nil, nil
	}
	before, oldAPI, oldErr := packagePayload(prior, true)
	after, newAPI, newErr := packagePayload(v, true)
	if oldErr == nil && newErr == nil && oldAPI && newAPI && samePayload(before, after) {
		return nil, v.Set("id", prior.Get("id"))
	}
	next, old := get(after, "uniqueness_key"), get(before, "uniqueness_key")
	if newErr == nil && next == nil {
		return []string{"id"}, nil
	}
	if oldErr != nil || newErr != nil || isUnknownAPI(before) || isUnknownAPI(after) || isUnknownAPI(next) || isUnknownAPI(old) {
		return nil, errors.New("package replacement uniqueness_key must be known")
	}
	if samePackageKey(str(old), str(next)) {
		return nil, errors.New("package replacement needs a new uniqueness_key; archived keys remain reserved")
	}
	return []string{"id"}, nil
}
func updatePackage(_ context.Context, v, prior *values) error {
	a, aa, ae := packagePayload(v, false)
	b, bb, be := packagePayload(prior, false)
	if ae == nil && be == nil && aa && bb && samePayload(a, b) {
		return nil
	}
	return errors.New("packages are immutable; update requires replacement")
}
func readPackage(ctx context.Context, v *values, remote any) error {
	if err := checkIdentity(ctx, v, remote); err != nil {
		return err
	}
	expected, apiMode, err := packagePayload(v, false)
	if err != nil {
		return err
	}
	raw, err := json.Marshal(remote)
	if err != nil {
		return err
	}
	if apiMode {
		request, err := json.Marshal(expected)
		if err != nil {
			return err
		}
		return verifyPackageRequest(request, raw)
	}
	if err = verifyLegacyPackage(raw); err != nil {
		return err
	}
	r := obj(remote)
	if get(r, "usage_statement_schedule", "frequency") != "MONTHLY" || get(r, "usage_statement_schedule", "day") != "CONTRACT_START" || len(arr(r["recurring_commits"])) > 1 {
		return errors.New("legacy package requires monthly contract-start statements and at most one recurring commit")
	}
	updates := map[string]any{"name": r["name"], "contract_name": r["contract_name"], "rate_card_id": r["rate_card_id"], "stripe_billing": r["billing_provider"] == "stripe" && r["delivery_method"] == "direct_to_billing_provider", "recurring_commit": nil}
	if (!v.Known("name") || v.Known("uniqueness_key")) && !samePackageKey(v.String("uniqueness_key"), str(r["uniqueness_key"])) {
		updates["uniqueness_key"] = r["uniqueness_key"]
	}
	if commits := arr(r["recurring_commits"]); len(commits) == 1 {
		c := obj(commits[0])
		access, invoice := obj(c["access_amount"]), obj(c["invoice_amount"])
		if c["recurrence_frequency"] != "MONTHLY" || c["proration"] != "NONE" || get(c, "starting_at_offset", "unit") != "DAYS" || !numberEqual(get(c, "starting_at_offset", "value"), json.Number("0")) || access["credit_type_id"] != invoice["credit_type_id"] || !numberEqual(access["quantity"], json.Number("1")) || !numberEqual(invoice["quantity"], json.Number("1")) {
			return errors.New("remote recurring commit differs from the legacy monthly prepaid model")
		}
		product := c["product_id"]
		if emptyAPIValue(product) {
			product = get(c, "product", "id")
		}
		duration := get(c, "commit_duration", "value")
		if text, ok := duration.(string); ok {
			duration = json.Number(text)
		}
		rollover := c["rollover_fraction"]
		if rollover == nil {
			rollover = json.Number("0")
		}
		updates["recurring_commit"] = map[string]any{"name": c["name"], "fixed_product_id": product, "credit_type_id": access["credit_type_id"], "monthly_price_cents": invoice["unit_price"], "monthly_grant_cents": access["unit_price"], "applicable_product_ids": c["applicable_product_ids"], "grant_duration_periods": duration, "rollover_fraction": rollover, "priority": c["priority"]}
	}
	return put(ctx, v, updates)
}
func importPackage(ctx context.Context, v *values, id string) error {
	if !strings.HasPrefix(strings.TrimSpace(id), "{") {
		return bindingruntime.ImportIdentity(ctx, v, id, []string{"id"}, "/")
	}
	raw, err := decodeJSON([]byte(id))
	if err != nil {
		return err
	}
	r := obj(raw)
	if err = nonempty(r, "id"); err != nil {
		return err
	}
	modes := 0
	for key, x := range r {
		switch key {
		case "id":
		case "package", "request", "request_json":
			if x != nil {
				modes++
			}
		default:
			return errors.New("unknown package import field")
		}
	}
	if modes != 1 {
		return errors.New("package import requires exactly one of package, request or request_json")
	}
	if r["package"] != nil {
		if err = bindingruntime.CheckObjectKeys(v.Get("package").Type(ctx), r["package"]); err != nil {
			return err
		}
	}
	if err = put(ctx, v, r); err != nil {
		return err
	}
	return validatePackage(ctx, v)
}
