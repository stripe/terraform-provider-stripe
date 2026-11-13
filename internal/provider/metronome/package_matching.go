//
// File generated from our OpenAPI spec
//

// Package normalization preserves the reviewed immutable-package matching policy.
package metronome

import (
	"encoding/json"
	"fmt"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
	"math/big"
	"reflect"
	"strings"
)

func decodeJSON(raw []byte) (any, error) { return bindingruntime.DecodeJSON(raw) }
func isUnknownAPI(value any) bool        { return bindingruntime.IsUnknown(value) }
func validateKnownPackageValue(value any) error {
	if isUnknownAPI(value) {
		return nil
	}
	root, ok := value.(map[string]any)
	if !ok {
		return fmt.Errorf("package request must be an object")
	}
	name, _ := root["name"].(string)
	if !isUnknownAPI(root["name"]) && strings.TrimSpace(name) == "" {
		return fmt.Errorf("package request needs a nonempty name")
	}
	key, _ := root["uniqueness_key"].(string)
	if root["uniqueness_key"] != nil && !isUnknownAPI(root["uniqueness_key"]) && strings.TrimSpace(key) == "" {
		return fmt.Errorf("uniqueness_key must be nonempty when supplied")
	}
	subscriptions, _ := root["subscriptions"].([]any)
	for _, item := range subscriptions {
		if isUnknownAPI(item) {
			continue
		}
		subscription, ok := item.(map[string]any)
		if !ok {
			return fmt.Errorf("subscription must be an object")
		}
		collection, _ := subscription["collection_schedule"].(string)
		if strings.EqualFold(collection, "ADVANCE") {
			if isUnknownAPI(subscription["proration"]) {
				continue
			}
			proration, _ := subscription["proration"].(map[string]any)
			behavior, _ := proration["invoice_behavior"].(string)
			if !isUnknownAPI(proration["invoice_behavior"]) && behavior == "" {
				return fmt.Errorf("ADVANCE subscriptions require proration.invoice_behavior")
			}
		}
	}
	return nil
}

// Response objects embed product/credit-type records where creation accepts IDs.
// Preserve response fields but add the corresponding creation field for matching.
func normalizePackageResponse(value any) any {
	switch v := value.(type) {
	case map[string]any:
		result := map[string]any{}
		for key, child := range v {
			result[key] = normalizePackageResponse(child)
		}
		for _, key := range []string{"product", "credit_type"} {
			if nested, ok := v[key].(map[string]any); ok {
				if id, ok := nested["id"]; ok {
					result[key+"_id"] = id
				}
			}
		}
		return result
	case []any:
		result := make([]any, len(v))
		for i, child := range v {
			result[i] = normalizePackageResponse(child)
		}
		return result
	default:
		return value
	}
}

func packageRequestMatches(expected, actual any, field string) error {
	if e, ok := expected.(map[string]any); ok {
		a, ok := actual.(map[string]any)
		if !ok {
			return fmt.Errorf("%s is not an object", field)
		}
		for key, child := range e {
			value, exists := a[key]
			if !exists && (child == nil || emptyJSONArray(child)) {
				continue
			}
			if err := packageRequestMatches(child, value, field+"."+key); err != nil {
				return err
			}
		}
		for key, value := range a {
			if _, exists := e[key]; exists || emptyAPIValue(value) {
				continue
			}
			if packageResponseDefault(key, value, e, field) {
				continue
			}
			return fmt.Errorf("%s.%s exists remotely but is absent from the creation request", field, key)
		}
		return nil
	}
	if e, ok := expected.([]any); ok {
		a, ok := actual.([]any)
		if !ok || len(e) != len(a) {
			return fmt.Errorf("%s has different items", field)
		}
		for i, child := range e {
			if err := packageRequestMatches(child, a[i], fmt.Sprintf("%s[%d]", field, i)); err != nil {
				return err
			}
		}
		return nil
	}
	if e, ok := expected.(json.Number); ok {
		var text string
		switch a := actual.(type) {
		case json.Number:
			text = string(a)
		case string:
			text = a
		}
		x, xok := new(big.Rat).SetString(string(e))
		y, yok := new(big.Rat).SetString(text)
		if xok && yok && x.Cmp(y) == 0 {
			return nil
		}
	}
	if !reflect.DeepEqual(expected, actual) {
		if e, ok := expected.(string); ok {
			if a, ok := actual.(string); ok {
				key := field[strings.LastIndex(field, ".")+1:]
				if key == "uniqueness_key" && samePackageKey(e, a) {
					return nil
				}
				for _, enum := range []string{"unit", "frequency", "billing_frequency", "recurrence_frequency", "rate_type", "access_type", "collection_schedule", "quantity_management_mode", "invoice_behavior", "day", "multiplier_override_prioritization"} {
					if key == enum && strings.EqualFold(e, a) {
						return nil
					}
				}
			}
		}
		return fmt.Errorf("%s differs from the remote package", field)
	}
	return nil
}

func samePackageKey(a, b string) bool {
	return packageKeyWithoutNamespace(a) == packageKeyWithoutNamespace(b)
}

// Metronome adds its package namespace to uniqueness keys in responses. Older
// requests could include the namespace too, producing a repeated prefix.
func packageKeyWithoutNamespace(key string) string {
	for strings.HasPrefix(key, "package:") {
		key = strings.TrimPrefix(key, "package:")
	}
	return key
}

func directSubscriptionPath(field string) bool {
	prefix := "request.subscriptions["
	if !strings.HasPrefix(field, prefix) || !strings.HasSuffix(field, "]") {
		return false
	}
	return !strings.Contains(strings.TrimPrefix(field, prefix), ".")
}

func directCreditPath(field string) bool {
	prefix := "request.credits["
	if !strings.HasPrefix(field, prefix) || !strings.HasSuffix(field, "]") {
		return false
	}
	return !strings.Contains(strings.TrimPrefix(field, prefix), ".")
}

func packageResponseDefault(key string, value any, expected map[string]any, field string) bool {
	switch key {
	case "uniqueness_key":
		return field == "request" // Optional duplicate-prevention token, not a billing term.
	case "id", "created_at", "created_by", "updated_at", "updated_by", "archived_at", "product", "credit_type":
		return true
	case "type":
		return directCreditPath(field) && value == "PREPAID" // Generated for a one-time package credit.
	case "usage_statement_schedule":
		if field != "request" {
			return false
		}
		schedule, ok := value.(map[string]any)
		return ok && len(schedule) == 2 && schedule["day"] == "FIRST_OF_MONTH" && schedule["frequency"] == "MONTHLY"
	case "invoice_schedule":
		if !directCreditPath(field) {
			return false
		}
		schedule, ok := value.(map[string]any)
		if !ok || len(schedule) != 4 || schedule["do_not_invoice"] != false || !emptyJSONArray(schedule["schedule_items"]) {
			return false
		}
		creditType, ok := schedule["credit_type"].(map[string]any)
		if !ok || len(creditType) < 1 || len(creditType) > 2 {
			return false
		}
		if _, named := creditType["name"]; len(creditType) == 2 && !named {
			return false
		}
		// The server emits an empty invoice schedule in the rate card's fiat
		// currency, which can differ from a custom-unit access schedule. With no
		// invoice items there is no billing amount or user-configurable term.
		currencyID, ok := creditType["id"].(string)
		return ok && currencyID != "" && reflect.DeepEqual(schedule["credit_type_id"], currencyID)
	case "multiplier_override_prioritization":
		return value == "LOWEST_MULTIPLIER"
	case "rate_type":
		return value == "LIST_RATE"
	case "access_type":
		return value == "SPEND"
	case "quantity_management_mode":
		return value == "QUANTITY_ONLY"
	case "unit":
		return strings.HasSuffix(field, "commit_duration") && value == "PERIODS"
	case "fiat_credit_type_id":
		return directSubscriptionPath(field) // Response-only in SubscriptionTemplate, absent from SubscriptionTemplateInput.
	case "billing_cycle_config":
		if !directSubscriptionPath(field) {
			return false
		}
		v, ok := value.(map[string]any)
		return ok && len(v) == 1 && v["invoice_placement"] == "ON_USAGE_INVOICE"
	case "contract_name":
		return expected["name"] == value
	case "starting_at_offset":
		if !directSubscriptionPath(field) {
			return false
		}
		v, ok := value.(map[string]any)
		if !ok || len(v) != 2 || v["unit"] != "DAYS" {
			return false
		}
		return v["value"] == json.Number("0") || v["value"] == "0"
	}
	return false
}

func emptyJSONArray(value any) bool { a, ok := value.([]any); return ok && len(a) == 0 }

func emptyAPIValue(value any) bool {
	if value == nil {
		return true
	}
	switch v := value.(type) {
	case string:
		return v == ""
	case []any:
		return len(v) == 0
	case map[string]any:
		return len(v) == 0
	}
	return false
}

func rejectUnrepresentedFields(object map[string]any, allowed []string, defaults map[string]any, field string) error {
	known := map[string]bool{}
	for _, key := range allowed {
		known[key] = true
	}
	for key, value := range object {
		if known[key] || emptyAPIValue(value) {
			continue
		}
		if def, ok := defaults[key]; ok && reflect.DeepEqual(value, def) {
			continue
		}
		return fmt.Errorf("%s.%s is not represented by this resource; use an API-shaped configuration instead", field, key)
	}
	return nil
}

func verifyLegacyPackage(remote []byte) error {
	value, err := decodeJSON(remote)
	if err != nil {
		return err
	}
	root, ok := value.(map[string]any)
	if !ok {
		return fmt.Errorf("package response must be an object")
	}
	if err := rejectUnrepresentedFields(root, []string{"id", "name", "contract_name", "rate_card_id", "uniqueness_key", "created_at", "created_by", "archived_at", "usage_statement_schedule", "recurring_commits", "billing_provider", "delivery_method"}, map[string]any{"multiplier_override_prioritization": "LOWEST_MULTIPLIER"}, "package"); err != nil {
		return err
	}
	billing, delivery := root["billing_provider"], root["delivery_method"]
	if !emptyAPIValue(billing) && (billing != "stripe" || delivery != "direct_to_billing_provider") {
		return fmt.Errorf("legacy package supports only Stripe direct billing or no billing provider")
	}
	if emptyAPIValue(billing) && !emptyAPIValue(delivery) {
		return fmt.Errorf("delivery method without billing provider is unsupported")
	}
	commits, _ := root["recurring_commits"].([]any)
	for _, item := range commits {
		commit, ok := item.(map[string]any)
		if !ok {
			return fmt.Errorf("invalid recurring commit")
		}
		if err := rejectUnrepresentedFields(commit, []string{"id", "name", "product", "product_id", "priority", "applicable_product_ids", "access_amount", "invoice_amount", "starting_at_offset", "commit_duration", "recurrence_frequency", "proration", "rollover_fraction"}, map[string]any{"rate_type": "LIST_RATE"}, "recurring_commit"); err != nil {
			return err
		}
		duration, _ := commit["commit_duration"].(map[string]any)
		if duration["unit"] != "PERIODS" {
			return fmt.Errorf("legacy commit duration must use PERIODS")
		}
		for _, key := range []string{"access_amount", "invoice_amount"} {
			amount, _ := commit[key].(map[string]any)
			if err := rejectUnrepresentedFields(amount, []string{"credit_type_id", "unit_price", "quantity"}, map[string]any{"access_type": "SPEND"}, key); err != nil {
				return err
			}
		}
	}
	return nil
}

func verifyPackageRequest(request, remote []byte) error {
	expected, err := decodeJSON(request)
	if err != nil {
		return err
	}
	actual, err := decodeJSON(remote)
	if err != nil {
		return err
	}
	e, ok := expected.(map[string]any)
	if !ok {
		return fmt.Errorf("package request must be an object")
	}
	a, ok := actual.(map[string]any)
	if !ok {
		return fmt.Errorf("package response must be an object")
	}
	for _, key := range []string{"commits", "credits", "overrides", "scheduled_charges", "scheduled_charges_on_usage_invoices", "recurring_commits", "recurring_credits", "subscriptions", "spend_trackers", "spend_threshold_configuration", "prepaid_balance_threshold_configuration", "billing_provider", "delivery_method"} {
		if _, exists := e[key]; !exists && a[key] != nil && !emptyJSONArray(a[key]) && a[key] != "" {
			return fmt.Errorf("remote package contains %s absent from the request", key)
		}
	}
	return packageRequestMatches(expected, normalizePackageResponse(actual), "request")
}
