//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
	"math/big"
	"reflect"
	"strings"
	"time"
	"unicode/utf8"
)

type values = bindingruntime.Values

func obj(value any) map[string]any { result, _ := value.(map[string]any); return result }
func arr(value any) []any          { result, _ := value.([]any); return result }
func str(value any) string         { result, _ := value.(string); return result }
func get(value any, path ...string) any {
	result, _ := bindingruntime.At(value, path...)
	return result
}
func rat(value any) (*big.Rat, bool) {
	switch n := value.(type) {
	case json.Number:
		return new(big.Rat).SetString(string(n))
	case string:
		return new(big.Rat).SetString(n)
	default:
		return nil, false
	}
}
func numberEqual(a, b any) bool {
	x, xok := rat(a)
	y, yok := rat(b)
	return xok && yok && x.Cmp(y) == 0
}
func nonempty(root map[string]any, fields ...string) error {
	for _, key := range fields {
		if isUnknownAPI(root[key]) {
			continue
		}
		if strings.TrimSpace(str(root[key])) == "" {
			return fmt.Errorf("%s must be nonempty", key)
		}
	}
	return nil
}
func put(ctx context.Context, v *values, fields map[string]any) error {
	for key, value := range fields {
		if err := v.SetJSON(ctx, key, value); err != nil {
			return err
		}
	}
	return nil
}
func nullable(value any) any {
	if value == "" {
		return nil
	}
	return value
}
func checkIdentity(_ context.Context, v *values, remote any) error {
	if str(get(remote, "id")) == "" || str(get(remote, "id")) != v.String("id") {
		return errors.New("API returned a missing or mismatched object identity")
	}
	return nil
}
func sameTime(a, b string) bool {
	x, xe := time.Parse(time.RFC3339Nano, a)
	y, ye := time.Parse(time.RFC3339Nano, b)
	return xe == nil && ye == nil && x.Equal(y)
}
func retainTime(previous, value any) any {
	if sameTime(str(previous), str(value)) {
		return previous
	}
	return value
}
func knownChange(v, prior *values, keys ...string) bool {
	for _, key := range keys {
		value := v.Get(key)
		if value != nil && !value.IsUnknown() && bindingruntime.Changed(v, prior, key) {
			return true
		}
	}
	return false
}
func ensureUnchanged(v, prior *values, keys ...string) error {
	if bindingruntime.Changed(v, prior, keys...) {
		return errors.New("this change requires resource replacement")
	}
	return nil
}
func reservedKey(v, prior *values, fields ...string) error {
	if bindingruntime.Changed(v, prior, fields...) && prior.Known("uniqueness_key") && (!v.Known("uniqueness_key") || v.String("uniqueness_key") == prior.String("uniqueness_key")) {
		return errors.New("replacement needs a new known uniqueness_key; Metronome reserves archived keys")
	}
	return nil
}
func enumValue(root map[string]any, field string, values map[string]struct{}, optional bool) error {
	value := root[field]
	if isUnknownAPI(value) || (optional && value == nil) {
		return nil
	}
	if _, ok := values[str(value)]; !ok {
		return fmt.Errorf("%s is not in the pinned API enum", field)
	}
	return nil
}
func subset(expected, actual any) bool {
	switch wanted := expected.(type) {
	case map[string]any:
		remote, ok := actual.(map[string]any)
		if !ok {
			return false
		}
		for k, x := range wanted {
			if !subset(x, remote[k]) {
				return false
			}
		}
		return true
	case []any:
		remote, ok := actual.([]any)
		if !ok || len(wanted) != len(remote) {
			return false
		}
		for i, x := range wanted {
			if !subset(x, remote[i]) {
				return false
			}
		}
		return true
	case json.Number:
		return numberEqual(expected, actual)
	default:
		return reflect.DeepEqual(expected, actual)
	}
}
func validateProduct(_ context.Context, v *values) error {
	root, err := v.JSON(true)
	if err != nil {
		return err
	}
	if err = nonempty(root, "name"); err != nil {
		return err
	}
	if isUnknownAPI(root["type"]) {
		return nil
	}
	kind := str(root["type"])
	if kind != "USAGE" && kind != "FIXED" && kind != "SUBSCRIPTION" && kind != "COMPOSITE" {
		return errors.New("product type must be USAGE, FIXED, SUBSCRIPTION, or COMPOSITE")
	}
	if kind == "USAGE" {
		if err = nonempty(root, "billable_metric_id"); err != nil {
			return err
		}
	} else if root["billable_metric_id"] != nil && !isUnknownAPI(root["billable_metric_id"]) {
		return errors.New("only USAGE accepts billable_metric_id")
	}
	for _, key := range []string{"composite_product_ids", "composite_tags"} {
		x := root[key]
		if isUnknownAPI(x) {
			continue
		}
		if kind == "COMPOSITE" && len(arr(x)) == 0 {
			return fmt.Errorf("COMPOSITE requires nonempty %s", key)
		}
		if kind != "COMPOSITE" && x != nil {
			return errors.New("composite fields apply only to COMPOSITE products")
		}
	}
	return nil
}
func readProduct(ctx context.Context, v *values, remote any) error {
	if err := checkIdentity(ctx, v, remote); err != nil {
		return err
	}
	root := obj(remote)
	current := obj(root["current"])
	if current == nil {
		return errors.New("product response omitted current fields")
	}
	if err := rejectUnrepresentedFields(current, []string{"name", "billable_metric_id", "composite_product_ids", "composite_tags", "starting_at", "created_at", "created_by"}, map[string]any{"exclude_free_usage": false, "include_composite_spend": false, "is_refundable": true, "composite_scope": "contract", "sql_breakdown_granularity": "service_period"}, "product"); err != nil {
		return err
	}
	updates := map[string]any{"name": current["name"], "type": root["type"], "billable_metric_id": nullable(current["billable_metric_id"])}
	if root["type"] == "COMPOSITE" {
		updates["composite_product_ids"] = current["composite_product_ids"]
		updates["composite_tags"] = current["composite_tags"]
	}
	if err := put(ctx, v, updates); err != nil {
		return err
	}
	return validateProduct(ctx, v)
}
func validateMetric(_ context.Context, v *values) error {
	root, err := v.JSON(true)
	if err != nil {
		return err
	}
	if err = nonempty(root, "name", "event_type"); err != nil {
		return err
	}
	if isUnknownAPI(root["aggregation_type"]) {
		return nil
	}
	kind := str(root["aggregation_type"])
	if kind == "" {
		kind = "COUNT"
	}
	if kind != "COUNT" && kind != "SUM" {
		return errors.New("aggregation_type must be COUNT or SUM")
	}
	key := root["aggregation_key"]
	if isUnknownAPI(key) {
		return nil
	}
	if kind == "COUNT" && key != nil {
		return errors.New("COUNT must omit aggregation_key")
	}
	if kind == "SUM" && (str(key) == "" || strings.TrimSpace(str(key)) != str(key)) {
		return errors.New("SUM requires an aggregation_key without surrounding whitespace")
	}
	return nil
}
func metricRequest(_ context.Context, v, prior *values, body map[string]any) (any, error) {
	if err := validateMetric(context.Background(), v); err != nil {
		return nil, err
	}
	kind := v.String("aggregation_type")
	if kind == "" {
		kind = "COUNT"
	}
	body["name"] = v.String("name")
	body["aggregation_type"] = kind
	body["event_type_filter"] = map[string]any{"in_values": []any{v.String("event_type")}}
	if kind == "SUM" {
		body["aggregation_key"] = v.String("aggregation_key")
		body["property_filters"] = []any{map[string]any{"name": v.String("aggregation_key"), "exists": true}}
	}
	return body, nil
}
func readMetric(ctx context.Context, v *values, remote any) error {
	if err := checkIdentity(ctx, v, remote); err != nil {
		return err
	}
	r := obj(remote)
	kind := strings.ToUpper(str(r["aggregation_type"]))
	events := arr(get(r, "event_type_filter", "in_values"))
	if (kind != "COUNT" && kind != "SUM") || len(events) != 1 || !emptyAPIValue(r["group_keys"]) || !emptyAPIValue(get(r, "event_type_filter", "not_in_values")) {
		return errors.New("metric has unsupported aggregation or filters")
	}
	filters := arr(r["property_filters"])
	if kind == "COUNT" && len(filters) > 0 {
		return errors.New("COUNT cannot round-trip property filters")
	}
	if kind == "SUM" {
		if len(filters) != 1 || get(filters[0], "name") != r["aggregation_key"] || get(filters[0], "exists") != true {
			return errors.New("SUM requires only its aggregation property's exists filter")
		}
		for key, x := range obj(filters[0]) {
			if key != "name" && key != "exists" && !emptyAPIValue(x) {
				return errors.New("SUM contains unsupported property constraints")
			}
		}
	}
	key := r["aggregation_key"]
	if kind == "COUNT" {
		key = nil
	}
	return put(ctx, v, map[string]any{"name": r["name"], "aggregation_type": kind, "event_type": events[0], "aggregation_key": key})
}
func validateCreditType(_ context.Context, v *values) error {
	if value := v.Get("name"); value != nil && value.IsUnknown() {
		return nil
	}
	name := v.String("name")
	if name == "" || strings.TrimSpace(name) != name {
		return errors.New("name must be nonempty with no surrounding whitespace")
	}
	return nil
}
func planCreditType(_ context.Context, v, prior *values) ([]string, error) {
	if knownChange(v, prior, "name") {
		return nil, errors.New("credit type rename is unsafe; create a new pricing unit and migrate consumers explicitly")
	}
	return nil, nil
}
func readCreditType(ctx context.Context, v *values, remote any) error {
	r := obj(remote)
	if r["is_currency"] == true {
		return errors.New("built-in currency credit types cannot be managed or archived")
	}
	if _, ok := r["is_currency"].(bool); !ok {
		return errors.New("credit type response omitted currency flag")
	}
	if err := checkIdentity(ctx, v, r); err != nil {
		return err
	}
	return put(ctx, v, map[string]any{"name": r["name"], "is_currency": r["is_currency"]})
}
func validateCustomFieldKey(_ context.Context, v *values) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	if err = enumValue(r, "entity", managedEntityValues, false); err != nil {
		return err
	}
	return nonempty(r, "key")
}
func planCustomFieldKey(_ context.Context, v, prior *values) ([]string, error) {
	if !bindingruntime.Changed(v, prior, "entity", "key") && bindingruntime.Changed(v, prior, "enforce_uniqueness") {
		return nil, errors.New("enforce_uniqueness cannot change for an existing entity/key")
	}
	return nil, nil
}
func checkCustomKeyAbsent(_ context.Context, v *values, remote any) error {
	items, ok := remote.([]any)
	if !ok {
		return errors.New("custom field list is missing")
	}
	for _, item := range items {
		if get(item, "entity") == v.String("entity") && get(item, "key") == v.String("key") {
			return errors.New("custom field key already exists; import it instead")
		}
	}
	return nil
}
func checkCreatedCustomKey(_ context.Context, v *values, remote any) error {
	wanted, err := bindingruntime.JSONValue(v.Get("enforce_uniqueness"), false)
	if err != nil {
		return err
	}
	if get(remote, "enforce_uniqueness") != wanted {
		return errors.New("created custom field key differs from enforce_uniqueness; identity retained for recovery")
	}
	return nil
}
func validateNotification(_ context.Context, v *values) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	if err = nonempty(r, "name"); err != nil {
		return err
	}
	if isUnknownAPI(r["policy"]) {
		return nil
	}
	return nonempty(obj(r["policy"]), "type", "offset")
}
func planNotification(_ context.Context, v, prior *values) ([]string, error) {
	return nil, reservedKey(v, prior, "name", "policy")
}
func updateNotification(_ context.Context, v, prior *values) error {
	if prior.Known("uniqueness_key") {
		return errors.New("notification cannot be updated")
	}
	return ensureUnchanged(v, prior, "name", "policy")
}
func readNotification(ctx context.Context, v *values, remote any) error {
	if err := checkIdentity(ctx, v, remote); err != nil {
		return err
	}
	r := obj(remote)
	if err := nonempty(r, "name"); err != nil {
		return err
	}
	if err := nonempty(obj(r["policy"]), "type", "offset"); err != nil {
		return err
	}
	return put(ctx, v, map[string]any{"name": r["name"], "policy": r["policy"]})
}
func validateCustomer(_ context.Context, v *values) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	if !isUnknownAPI(r["name"]) && (strings.TrimSpace(str(r["name"])) == "" || utf8.RuneCountInString(str(r["name"])) > 160) {
		return errors.New("customer name must be nonempty and at most 160 characters")
	}
	aliases := arr(r["ingest_aliases"])
	if len(aliases) > 2000 {
		return errors.New("ingest_aliases exceeds 2000")
	}
	for _, alias := range aliases {
		if isUnknownAPI(alias) {
			continue
		}
		if utf8.RuneCountInString(str(alias)) < 1 || utf8.RuneCountInString(str(alias)) > 128 {
			return errors.New("each ingest alias must be 1–128 characters")
		}
	}
	return nil
}
func customerAliases(_ context.Context, v *values) ([]any, error) {
	value, err := bindingruntime.JSONValue(v.Get("ingest_aliases"), false)
	if err != nil {
		return nil, err
	}
	return arr(value), nil
}
func checkAliasOwnership(_ context.Context, v *values, remote any) error {
	items, ok := remote.([]any)
	if !ok {
		return errors.New("customer list missing")
	}
	for _, item := range items {
		if str(get(item, "id")) == "" || str(get(item, "id")) != v.String("id") {
			return errors.New("ingest alias belongs to another customer")
		}
	}
	return nil
}
func readCustomer(ctx context.Context, v *values, remote any) error {
	if err := checkIdentity(ctx, v, remote); err != nil {
		return err
	}
	r := obj(remote)
	if err := nonempty(r, "name"); err != nil {
		return err
	}
	aliases := r["ingest_aliases"]
	if aliases == nil {
		aliases = []any{}
	}
	return put(ctx, v, map[string]any{"name": r["name"], "ingest_aliases": aliases})
}
func checkCustomerContracts(_ context.Context, _ *values, remote any) error {
	items, ok := remote.([]any)
	if !ok {
		return errors.New("contract list is missing")
	}
	for _, item := range items {
		if emptyAPIValue(get(item, "archived_at")) {
			return errors.New("customer has active contracts; archive all contracts before archiving the customer")
		}
	}
	return nil
}
