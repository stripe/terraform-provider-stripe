//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"encoding/json"
	"errors"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
	"strings"
	"time"
)

func validateRateCard(_ context.Context, v *values) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	if err = nonempty(r, "name"); err != nil {
		return err
	}
	seen := map[string]bool{}
	for _, item := range arr(r["credit_type_conversions"]) {
		if isUnknownAPI(item) {
			continue
		}
		c := obj(item)
		if err = nonempty(c, "custom_credit_type_id"); err != nil {
			return err
		}
		if !isUnknownAPI(c["fiat_per_custom_credit"]) {
			n, ok := rat(c["fiat_per_custom_credit"])
			if !ok || n.Sign() <= 0 {
				return errors.New("fiat_per_custom_credit must be positive")
			}
		}
		if id, ok := c["custom_credit_type_id"].(string); ok {
			if seen[id] {
				return errors.New("duplicate credit type conversion")
			}
			seen[id] = true
		}
	}
	product, price := r["usage_product_id"], r["price_cents"]
	if !isUnknownAPI(product) && !isUnknownAPI(price) && (product == nil) != (price == nil) {
		return errors.New("usage_product_id and price_cents must be set together")
	}
	if product != nil && !isUnknownAPI(product) && len(arr(r["rates"])) > 0 {
		return errors.New("use rates or legacy single-rate fields")
	}
	if price != nil && !isUnknownAPI(price) {
		n, ok := rat(price)
		if !ok || n.Sign() < 0 {
			return errors.New("price_cents must be nonnegative")
		}
	}
	if start, ok := r["starting_at"].(string); ok {
		if product == nil {
			return errors.New("starting_at requires the legacy usage rate")
		}
		at, err := time.Parse("2006-01-02T15:04:05.000Z", start)
		if err != nil || !at.Equal(at.Truncate(time.Hour)) {
			return errors.New("starting_at must be a UTC hour boundary")
		}
	}
	seen = map[string]bool{}
	for _, item := range arr(r["rates"]) {
		if isUnknownAPI(item) {
			continue
		}
		rate := obj(item)
		if err = nonempty(rate, "product_id"); err != nil {
			return err
		}
		if !isUnknownAPI(rate["rate_type"]) && rate["rate_type"] != "FLAT" && rate["rate_type"] != "PERCENTAGE" {
			return errors.New("rates supports FLAT and PERCENTAGE")
		}
		if !isUnknownAPI(rate["price"]) {
			n, ok := rat(rate["price"])
			if !ok || n.Sign() < 0 {
				return errors.New("rate price must be nonnegative")
			}
		}
		if isUnknownAPI(rate["starting_at"]) {
			continue
		}
		start, err := time.Parse(time.RFC3339Nano, str(rate["starting_at"]))
		if err != nil {
			return errors.New("rate starting_at must have an ISO 8601 timezone")
		}
		if isUnknownAPI(rate["product_id"]) || isUnknownAPI(rate["billing_frequency"]) {
			continue
		}
		key := str(rate["product_id"]) + "\x00" + start.UTC().Format(time.RFC3339Nano) + "\x00" + str(rate["billing_frequency"])
		if seen[key] {
			return errors.New("duplicate rate product/start/frequency")
		}
		seen[key] = true
	}
	return nil
}
func prepareRateCard(ctx context.Context, v *values) error {
	if err := validateRateCard(ctx, v); err != nil {
		return err
	}
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	mode := "none"
	if v.Known("usage_product_id") {
		mode = "legacy"
		if !v.Known("starting_at") {
			if err = v.SetJSON(ctx, "starting_at", time.Now().UTC().Truncate(time.Hour).Format("2006-01-02T15:04:05.000Z")); err != nil {
				return err
			}
		}
	} else {
		if len(arr(r["rates"])) > 0 {
			mode = "list"
		}
		if err = v.SetJSON(ctx, "starting_at", nil); err != nil {
			return err
		}
	}
	return v.SetJSON(ctx, "rate_mode", mode)
}
func rateRequests(_ context.Context, v *values) ([]any, error) {
	r, err := v.JSON(true)
	if err != nil {
		return nil, err
	}
	if !v.Known("credit_type_id") {
		return nil, errors.New("rate card default credit type is missing")
	}
	var items []any
	if v.String("rate_mode") == "legacy" {
		items = append(items, map[string]any{"rate_card_id": v.String("id"), "product_id": v.String("usage_product_id"), "credit_type_id": v.String("credit_type_id"), "starting_at": v.String("starting_at"), "rate_type": "FLAT", "entitled": true, "price": r["price_cents"]})
	}
	for _, item := range arr(r["rates"]) {
		rate := obj(bindingruntime.StripNulls(item))
		if rate == nil {
			return nil, errors.New("rate must be an object")
		}
		rate["rate_card_id"] = v.String("id")
		if emptyAPIValue(rate["credit_type_id"]) {
			rate["credit_type_id"] = v.String("credit_type_id")
		}
		items = append(items, rate)
	}
	return items, nil
}
func readRateCard(ctx context.Context, v *values, remote any) error {
	if err := checkIdentity(ctx, v, remote); err != nil {
		return err
	}
	r := obj(remote)
	credit := str(get(r, "fiat_credit_type", "id"))
	if credit == "" {
		return errors.New("rate card response omitted fiat credit type")
	}
	updates := map[string]any{"name": r["name"], "credit_type_id": credit}
	if v.Known("description") || !emptyAPIValue(r["description"]) {
		updates["description"] = nullable(r["description"])
	}
	before, err := v.JSON(true)
	if err != nil {
		return err
	}
	if before["credit_type_conversions"] != nil && !isUnknownAPI(before["credit_type_conversions"]) {
		refreshed := []any{}
		for _, declared := range arr(before["credit_type_conversions"]) {
			for _, actual := range arr(r["credit_type_conversions"]) {
				if get(actual, "custom_credit_type", "id") == get(declared, "custom_credit_type_id") {
					number := get(actual, "fiat_per_custom_credit")
					n, ok := rat(number)
					if !ok || n.Sign() <= 0 {
						return errors.New("API returned invalid credit conversion")
					}
					refreshed = append(refreshed, map[string]any{"custom_credit_type_id": get(declared, "custom_credit_type_id"), "fiat_per_custom_credit": number})
					break
				}
			}
		}
		updates["credit_type_conversions"] = refreshed
	}
	return put(ctx, v, updates)
}
func rateCardUpdateRequest(_ context.Context, v, prior *values, body map[string]any) (any, error) {
	if !v.Known("description") && prior != nil && prior.Known("description") {
		body["description"] = ""
	}
	return body, nil
}
func rateScheduleRequest(_ context.Context, v, prior *values, body map[string]any) (any, error) {
	if v.String("rate_mode") == "none" {
		return bindingruntime.SkipRequest{}, nil
	}
	body["starting_at"] = "1970-01-01T00:00:00.000Z"
	if v.String("rate_mode") == "legacy" && v.Known("starting_at") {
		body["starting_at"] = v.String("starting_at")
	}
	return body, nil
}
func matchRate(declared, actual any) bool {
	return get(declared, "product_id") == get(actual, "product_id") && sameTime(str(get(declared, "starting_at")), str(get(actual, "starting_at"))) && str(get(declared, "billing_frequency")) == str(get(actual, "billing_frequency"))
}
func readRateSchedule(ctx context.Context, v *values, remote any) error {
	rates, ok := remote.([]any)
	if !ok {
		return errors.New("rate schedule missing")
	}
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	mode := v.String("rate_mode")
	if mode == "" {
		if len(arr(r["rates"])) > 0 {
			mode = "list"
		} else if v.Known("usage_product_id") {
			mode = "legacy"
		} else if len(rates) == 0 {
			mode = "none"
		} else {
			product := get(rates[0], "product_id")
			for _, rate := range rates {
				price, ok := rat(get(rate, "rate", "price"))
				_, timeErr := time.Parse(time.RFC3339Nano, str(get(rate, "starting_at")))
				if get(rate, "product_id") != product || get(rate, "entitled") != true || get(rate, "rate", "rate_type") != "FLAT" || get(rate, "rate", "credit_type", "id") != v.String("credit_type_id") || !ok || !price.IsInt() || !emptyAPIValue(get(rate, "billing_frequency")) || timeErr != nil {
					return errors.New("explicit rate import required; bare IDs support one product with entitled whole-unit FLAT rates")
				}
			}
			mode = "legacy"
		}
	}
	if err = v.SetJSON(ctx, "rate_mode", mode); err != nil {
		return err
	}
	if mode == "none" || mode == "list" {
		updates := map[string]any{"price_cents": nil, "starting_at": nil}
		var refreshed []any
		if r["rates"] != nil {
			refreshed = []any{}
		}
		for _, declared := range arr(r["rates"]) {
			var match any
			for _, actual := range rates {
				if matchRate(declared, actual) {
					if match != nil {
						return errors.New("ambiguous managed rate")
					}
					match = actual
				}
			}
			if match == nil {
				continue
			}
			kind := get(match, "rate", "rate_type")
			if kind != "FLAT" && kind != "PERCENTAGE" {
				return errors.New("managed rate changed to an unsupported type")
			}
			if _, ok := rat(get(match, "rate", "price")); !ok {
				return errors.New("rate price missing")
			}
			next := obj(declared)
			next["rate_type"] = kind
			next["entitled"] = get(match, "entitled")
			next["price"] = get(match, "rate", "price")
			credit := get(match, "rate", "credit_type", "id")
			if next["credit_type_id"] != nil || credit != v.String("credit_type_id") {
				next["credit_type_id"] = credit
			}
			refreshed = append(refreshed, next)
		}
		if refreshed != nil {
			updates["rates"] = refreshed
		}
		return put(ctx, v, updates)
	}
	var match any
	for _, rate := range rates {
		if v.Known("usage_product_id") && get(rate, "product_id") != v.String("usage_product_id") {
			continue
		}
		if v.Known("starting_at") && !sameTime(str(get(rate, "starting_at")), v.String("starting_at")) {
			continue
		}
		if match != nil {
			if v.Known("starting_at") {
				return errors.New("ambiguous legacy rate")
			}
			if get(rate, "product_id") != get(match, "product_id") {
				return errors.New("bare import requires rates for one usage product")
			}
			a, _ := time.Parse(time.RFC3339Nano, str(get(rate, "starting_at")))
			b, _ := time.Parse(time.RFC3339Nano, str(get(match, "starting_at")))
			if a.After(b) {
				match = rate
			}
		} else {
			match = rate
		}
	}
	if match == nil {
		return v.SetJSON(ctx, "price_cents", nil)
	}
	price, ok := rat(get(match, "rate", "price"))
	if get(match, "entitled") != true || get(match, "rate", "rate_type") != "FLAT" || get(match, "rate", "credit_type", "id") != v.String("credit_type_id") || !ok || !price.IsInt() || !emptyAPIValue(get(match, "billing_frequency")) {
		return v.SetJSON(ctx, "price_cents", nil)
	}
	updates := map[string]any{"usage_product_id": get(match, "product_id"), "price_cents": json.Number(price.Num().String())}
	if !v.Known("starting_at") {
		at, err := time.Parse(time.RFC3339Nano, str(get(match, "starting_at")))
		if err != nil {
			return errors.New("invalid rate start")
		}
		updates["starting_at"] = at.UTC().Format("2006-01-02T15:04:05.000Z")
	}
	return put(ctx, v, updates)
}
func importRateCard(ctx context.Context, v *values, id string) error {
	if !strings.HasPrefix(strings.TrimSpace(id), "{") {
		return bindingruntime.ImportIdentity(ctx, v, id, []string{"id"}, "/")
	}
	decoded, err := decodeJSON([]byte(id))
	if err != nil {
		return err
	}
	r := obj(decoded)
	if err = nonempty(r, "id"); err != nil {
		return err
	}
	if _, exists := r["rates"]; !exists {
		return errors.New("JSON rate-card import requires an explicit rates array or null")
	}
	for key := range r {
		if key != "id" && key != "rates" && key != "credit_type_conversions" {
			return errors.New("unknown rate-card import field")
		}
	}
	for _, key := range []string{"rates", "credit_type_conversions"} {
		if err = bindingruntime.CheckObjectKeys(v.Get(key).Type(ctx), r[key]); err != nil {
			return err
		}
	}
	if err = put(ctx, v, r); err != nil {
		return err
	}
	mode := "list"
	if len(arr(r["rates"])) == 0 {
		mode = "none"
	}
	if err = v.SetJSON(ctx, "rate_mode", mode); err != nil {
		return err
	}
	for _, rate := range arr(r["rates"]) {
		if err = nonempty(obj(rate), "product_id", "rate_type", "starting_at"); err != nil {
			return err
		}
		if _, ok := rat(get(rate, "price")); !ok {
			return errors.New("import rate requires price")
		}
		if _, ok := get(rate, "entitled").(bool); !ok {
			return errors.New("import rate requires entitled")
		}
	}
	return nil
}
func verifyRateCardImport(ctx context.Context, v *values, remote any) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	for _, declared := range arr(r["credit_type_conversions"]) {
		found := false
		for _, actual := range arr(get(remote, "credit_type_conversions")) {
			if get(declared, "custom_credit_type_id") == get(actual, "custom_credit_type", "id") && numberEqual(get(declared, "fiat_per_custom_credit"), get(actual, "fiat_per_custom_credit")) {
				found = true
				break
			}
		}
		if !found {
			return errors.New("import conversion does not match remote card")
		}
	}
	return readRateCard(ctx, v, remote)
}
func verifyRateImport(_ context.Context, v *values, remote any) error {
	r, err := v.JSON(true)
	if err != nil {
		return err
	}
	for _, declared := range arr(r["rates"]) {
		found := false
		for _, actual := range arr(remote) {
			if matchRate(declared, actual) {
				found = true
				break
			}
		}
		if !found {
			return errors.New("import rate does not exist at its product/start/frequency")
		}
	}
	return nil
}
