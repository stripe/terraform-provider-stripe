//
// File generated from our OpenAPI spec
//

package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	stripe "github.com/stripe/stripe-go/v86"
)

func snakeToPascalName(name string) string {
	parts := strings.Split(name, "_")
	for i, part := range parts {
		if part == "" {
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
	}
	return strings.Join(parts, "")
}

func pascalToSnakeName(name string) string {
	if name == "" {
		return ""
	}
	runes := []rune(name)
	result := make([]rune, 0, len(runes)+4)
	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				prev := runes[i-1]
				prevLowerOrDigit := (prev >= 'a' && prev <= 'z') || (prev >= '0' && prev <= '9')
				nextLower := i+1 < len(runes) && runes[i+1] >= 'a' && runes[i+1] <= 'z'
				if prevLowerOrDigit || ((prev >= 'A' && prev <= 'Z') && nextLower) {
					result = append(result, '_')
				}
			}
			result = append(result, r+'a'-'A')
			continue
		}
		result = append(result, r)
	}
	return string(result)
}

func derefValue(v reflect.Value) (reflect.Value, bool) {
	if !v.IsValid() {
		return reflect.Value{}, false
	}
	for v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return reflect.Value{}, false
		}
		v = v.Elem()
	}
	return v, true
}

func reflectValueToPlain(v reflect.Value) interface{} {
	deref, ok := derefValue(v)
	if !ok {
		return nil
	}
	switch deref.Kind() {
	case reflect.String:
		return deref.String()
	case reflect.Bool:
		return deref.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return deref.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(deref.Uint())
	case reflect.Float32, reflect.Float64:
		return deref.Float()
	case reflect.Slice, reflect.Array:
		items := make([]interface{}, 0, deref.Len())
		for i := 0; i < deref.Len(); i++ {
			items = append(items, reflectValueToPlain(deref.Index(i)))
		}
		return items
	case reflect.Map:
		result := map[string]interface{}{}
		for _, key := range deref.MapKeys() {
			if key.Kind() != reflect.String {
				continue
			}
			result[key.String()] = reflectValueToPlain(deref.MapIndex(key))
		}
		return result
	case reflect.Struct:
		result := map[string]interface{}{}
		for i := 0; i < deref.NumField(); i++ {
			field := deref.Type().Field(i)
			if field.PkgPath != "" {
				continue
			}
			result[pascalToSnakeName(field.Name)] = reflectValueToPlain(deref.Field(i))
		}
		return result
	default:
		return fmt.Sprintf("%v", deref.Interface())
	}
}

func attrValueToPlain(value attr.Value) interface{} {
	if value == nil || value.IsNull() || value.IsUnknown() {
		return nil
	}
	switch v := value.(type) {
	case basetypes.StringValue:
		return v.ValueString()
	case basetypes.Int64Value:
		return v.ValueInt64()
	case basetypes.Float64Value:
		return v.ValueFloat64()
	case basetypes.BoolValue:
		return v.ValueBool()
	case basetypes.ListValue:
		result := make([]interface{}, 0, len(v.Elements()))
		for _, element := range v.Elements() {
			result = append(result, attrValueToPlain(element))
		}
		return result
	case basetypes.SetValue:
		result := make([]interface{}, 0, len(v.Elements()))
		for _, element := range v.Elements() {
			result = append(result, attrValueToPlain(element))
		}
		return result
	case basetypes.MapValue:
		result := map[string]interface{}{}
		for key, element := range v.Elements() {
			result[key] = attrValueToPlain(element)
		}
		return result
	case basetypes.ObjectValue:
		result := map[string]interface{}{}
		for key, element := range v.Attributes() {
			result[key] = attrValueToPlain(element)
		}
		return result
	default:
		return nil
	}
}

func stringValueAtAttrPath(value attr.Value, path ...string) (string, bool) {
	plain := attrValueToPlain(value)
	if plain == nil {
		return "", false
	}
	if len(path) == 0 {
		stringValue, ok := plain.(string)
		return stringValue, ok
	}
	nestedPlain, ok := plainValueAtPath(plain, path...)
	if !ok {
		return "", false
	}
	stringValue, ok := nestedPlain.(string)
	return stringValue, ok
}

func setPlainValueAtPath(target map[string]interface{}, path []string, value interface{}) bool {
	if len(path) == 0 {
		return false
	}
	current := target
	for i, segment := range path {
		if i == len(path)-1 {
			current[segment] = value
			return true
		}
		next, exists := current[segment]
		if !exists || next == nil {
			nested := map[string]interface{}{}
			current[segment] = nested
			current = nested
			continue
		}
		nextMap, ok := next.(map[string]interface{})
		if !ok {
			return false
		}
		current = nextMap
	}
	return false
}

func plainKeyedListToMap(plain interface{}, keyField string) (interface{}, bool) {
	if plain == nil {
		return nil, true
	}
	items, ok := plain.([]interface{})
	if !ok {
		return nil, false
	}
	result := map[string]interface{}{}
	for _, item := range items {
		entry, ok := item.(map[string]interface{})
		if !ok {
			return nil, false
		}
		rawKey, exists := entry[keyField]
		if !exists {
			return nil, false
		}
		key, ok := rawKey.(string)
		if !ok || key == "" {
			return nil, false
		}
		copied := map[string]interface{}{}
		for entryKey, entryValue := range entry {
			if entryKey == keyField {
				continue
			}
			copied[entryKey] = entryValue
		}
		result[key] = copied
	}
	return result, true
}

func plainMapToKeyedList(plain interface{}, keyField string) interface{} {
	sourceMap, ok := plain.(map[string]interface{})
	if !ok {
		return plain
	}
	keys := make([]string, 0, len(sourceMap))
	for key := range sourceMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	result := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		entryMap, ok := sourceMap[key].(map[string]interface{})
		if !ok {
			result = append(result, map[string]interface{}{keyField: key, "value": sourceMap[key]})
			continue
		}
		copied := map[string]interface{}{keyField: key}
		for entryKey, entryValue := range entryMap {
			copied[entryKey] = entryValue
		}
		result = append(result, copied)
	}
	return result
}

func keyedListCurrentEntries(items []interface{}, keyField string) (map[string]interface{}, []string, bool) {
	entries := map[string]interface{}{}
	order := make([]string, 0, len(items))
	for _, item := range items {
		entryMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, nil, false
		}
		rawKey, exists := entryMap[keyField]
		if !exists {
			return nil, nil, false
		}
		key, ok := rawKey.(string)
		if !ok || key == "" {
			return nil, nil, false
		}
		entries[key] = item
		order = append(order, key)
	}
	return entries, order, true
}

func plainMapToKeyedListWithCurrent(plain map[string]interface{}, current []interface{}, keyField string) []interface{} {
	currentEntries, currentOrder, ok := keyedListCurrentEntries(current, keyField)
	if !ok {
		converted, _ := plainMapToKeyedList(plain, keyField).([]interface{})
		return converted
	}
	keys := make([]string, 0, len(plain))
	seen := map[string]bool{}
	for _, key := range currentOrder {
		if _, exists := plain[key]; exists {
			keys = append(keys, key)
			seen[key] = true
		}
	}
	remaining := make([]string, 0, len(plain))
	for key := range plain {
		if !seen[key] {
			remaining = append(remaining, key)
		}
	}
	sort.Strings(remaining)
	keys = append(keys, remaining...)

	result := make([]interface{}, 0, len(keys))
	for _, key := range keys {
		currentEntry := currentEntries[key]
		convertedEntry := applyConfiguredKeyedListShapes(plain[key], currentEntry)
		entryMap, ok := convertedEntry.(map[string]interface{})
		if !ok {
			result = append(result, map[string]interface{}{keyField: key, "value": convertedEntry})
			continue
		}
		copied := map[string]interface{}{keyField: key}
		for entryKey, entryValue := range entryMap {
			copied[entryKey] = entryValue
		}
		result = append(result, copied)
	}
	return result
}

func applyConfiguredKeyedListShapes(observed interface{}, current interface{}) interface{} {
	switch currentValue := current.(type) {
	case []interface{}:
		switch observedValue := observed.(type) {
		case []interface{}:
			result := make([]interface{}, 0, len(observedValue))
			for index, item := range observedValue {
				var currentItem interface{}
				if index < len(currentValue) {
					currentItem = currentValue[index]
				}
				result = append(result, applyConfiguredKeyedListShapes(item, currentItem))
			}
			return result
		case map[string]interface{}:
			if converted := plainMapToKeyedListWithCurrent(observedValue, currentValue, "key"); converted != nil {
				return converted
			}
		}
		return observed
	case map[string]interface{}:
		observedMap, ok := observed.(map[string]interface{})
		if !ok {
			return observed
		}
		result := map[string]interface{}{}
		for key, value := range observedMap {
			result[key] = applyConfiguredKeyedListShapes(value, currentValue[key])
		}
		return result
	default:
		return observed
	}
}

type plainStringLeafAlias struct {
	Target []string
	Source []string
}

func applyPlainStringLeafAliases(plain interface{}, aliases []plainStringLeafAlias) interface{} {
	sourceMap, ok := plain.(map[string]interface{})
	if !ok {
		return plain
	}
	for _, alias := range aliases {
		if len(alias.Target) == 0 || len(alias.Source) == 0 {
			continue
		}
		targetValue, targetExists := plainValueAtPath(sourceMap, alias.Target...)
		if targetExists && targetValue != nil {
			continue
		}
		sourceValue, sourceExists := plainValueAtPath(sourceMap, alias.Source...)
		if !sourceExists || sourceValue == nil {
			continue
		}
		setPlainValueAtPath(sourceMap, alias.Target, sourceValue)
	}
	return sourceMap
}

type plainObjectDefaultLeafValue struct {
	Target []string
	Value  interface{}
}

type plainNestedObjectDefaultLeafValues struct {
	ObjectPath []string
	Defaults   []plainObjectDefaultLeafValue
}

func applyPlainObjectDefaultLeafValues(plain interface{}, defaults []plainObjectDefaultLeafValue) interface{} {
	sourceMap, ok := plain.(map[string]interface{})
	if !ok {
		return plain
	}
	for _, defaultLeaf := range defaults {
		if len(defaultLeaf.Target) == 0 {
			continue
		}
		existingValue, exists := plainValueAtPath(sourceMap, defaultLeaf.Target...)
		if exists && existingValue != nil {
			continue
		}
		setPlainValueAtPath(sourceMap, defaultLeaf.Target, defaultLeaf.Value)
	}
	return sourceMap
}

func applyPlainObjectDefaultLeafValuesAtPath(plain interface{}, path []string, defaults []plainObjectDefaultLeafValue) (interface{}, bool) {
	if len(path) == 0 {
		return applyPlainObjectDefaultLeafValues(plain, defaults), true
	}
	segment := path[0]
	switch sourceValue := plain.(type) {
	case map[string]interface{}:
		if segment == "*" {
			changed := false
			for key, nextSource := range sourceValue {
				updated, itemChanged := applyPlainObjectDefaultLeafValuesAtPath(nextSource, path[1:], defaults)
				if !itemChanged {
					continue
				}
				sourceValue[key] = updated
				changed = true
			}
			return sourceValue, changed
		}
		nextSource, ok := sourceValue[segment]
		if !ok {
			return plain, false
		}
		updated, changed := applyPlainObjectDefaultLeafValuesAtPath(nextSource, path[1:], defaults)
		if !changed {
			return plain, false
		}
		sourceValue[segment] = updated
		return sourceValue, true
	case []interface{}:
		if segment == "*" {
			changed := false
			for index, item := range sourceValue {
				updated, itemChanged := applyPlainObjectDefaultLeafValuesAtPath(item, path[1:], defaults)
				if !itemChanged {
					continue
				}
				sourceValue[index] = updated
				changed = true
			}
			return sourceValue, changed
		}
		index, err := strconv.Atoi(segment)
		if err != nil || index < 0 || index >= len(sourceValue) {
			return plain, false
		}
		updated, changed := applyPlainObjectDefaultLeafValuesAtPath(sourceValue[index], path[1:], defaults)
		if !changed {
			return plain, false
		}
		sourceValue[index] = updated
		return sourceValue, true
	default:
		return plain, false
	}
}

func applyPlainNestedObjectDefaultLeafValues(plain interface{}, defaults []plainNestedObjectDefaultLeafValues) interface{} {
	current := plain
	for _, defaultSet := range defaults {
		updated, changed := applyPlainObjectDefaultLeafValuesAtPath(current, defaultSet.ObjectPath, defaultSet.Defaults)
		if changed {
			current = updated
		}
	}
	return current
}

func mergeMissingPlainLeaves(observed interface{}, current interface{}) interface{} {
	switch observedValue := observed.(type) {
	case nil:
		return current
	case map[string]interface{}:
		currentValues, ok := current.(map[string]interface{})
		if !ok {
			return observed
		}
		result := map[string]interface{}{}
		for key, value := range observedValue {
			result[key] = value
		}
		for key, currentValue := range currentValues {
			observedLeaf, exists := result[key]
			if !exists || observedLeaf == nil {
				result[key] = currentValue
				continue
			}
			result[key] = mergeMissingPlainLeaves(observedLeaf, currentValue)
		}
		return result
	case []interface{}:
		currentValues, ok := current.([]interface{})
		if !ok {
			return observed
		}
		result := make([]interface{}, len(observedValue))
		for index, item := range observedValue {
			if index < len(currentValues) {
				result[index] = mergeMissingPlainLeaves(item, currentValues[index])
			} else {
				result[index] = item
			}
		}
		return result
	default:
		return observed
	}
}

func suppressUnconfiguredOptionalReadbackLeaves(observed interface{}, current interface{}, paths [][]string) interface{} {
	result := observed
	for _, path := range paths {
		updated, changed := suppressUnconfiguredOptionalReadbackLeafAtPath(result, current, path)
		if changed {
			result = updated
		}
	}
	return result
}

func suppressUnconfiguredOptionalReadbackLeafAtPath(observed interface{}, current interface{}, path []string) (interface{}, bool) {
	if len(path) == 0 {
		if current == nil {
			return nil, true
		}
		return observed, false
	}
	switch observedValue := observed.(type) {
	case []interface{}:
		currentValues, _ := current.([]interface{})
		result := make([]interface{}, len(observedValue))
		changed := false
		for index, item := range observedValue {
			var currentItem interface{}
			if index < len(currentValues) {
				currentItem = currentValues[index]
			}
			updatedItem, itemChanged := suppressUnconfiguredOptionalReadbackLeafAtPath(item, currentItem, path)
			result[index] = updatedItem
			changed = changed || itemChanged
		}
		return result, changed
	case map[string]interface{}:
		currentValues, _ := current.(map[string]interface{})
		segment := path[0]
		child, exists := observedValue[segment]
		if !exists {
			return observed, false
		}
		var currentChild interface{}
		if currentValues != nil {
			currentChild = currentValues[segment]
		}
		if len(path) == 1 {
			if currentChild != nil {
				return observed, false
			}
			result := map[string]interface{}{}
			for key, value := range observedValue {
				if key == segment {
					continue
				}
				result[key] = value
			}
			return result, true
		}
		updatedChild, changed := suppressUnconfiguredOptionalReadbackLeafAtPath(child, currentChild, path[1:])
		if !changed {
			return observed, false
		}
		result := map[string]interface{}{}
		for key, value := range observedValue {
			result[key] = value
		}
		if updatedChild == nil {
			delete(result, segment)
		} else {
			result[segment] = updatedChild
		}
		return result, true
	default:
		return observed, false
	}
}

func extractListObjectData(plain interface{}) interface{} {
	sourceMap, ok := plain.(map[string]interface{})
	if !ok {
		return plain
	}
	data, exists := sourceMap["data"]
	if !exists {
		return plain
	}
	return data
}

func plainValueIsEmpty(plain interface{}) bool {
	switch v := plain.(type) {
	case nil:
		return true
	case bool:
		return !v
	case []interface{}:
		return len(v) == 0
	case map[string]interface{}:
		if len(v) == 0 {
			return true
		}
		for _, value := range v {
			if !plainValueIsEmpty(value) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func suppressUnconfiguredDecimalMirrorLeaves(observed interface{}, current interface{}) interface{} {
	switch observedValue := observed.(type) {
	case []interface{}:
		currentValues, _ := current.([]interface{})
		result := make([]interface{}, 0, len(observedValue))
		for index, item := range observedValue {
			var currentItem interface{}
			if index < len(currentValues) {
				currentItem = currentValues[index]
			}
			result = append(result, suppressUnconfiguredDecimalMirrorLeaves(item, currentItem))
		}
		return result
	case map[string]interface{}:
		currentValues, _ := current.(map[string]interface{})
		result := map[string]interface{}{}
		for key, value := range observedValue {
			result[key] = suppressUnconfiguredDecimalMirrorLeaves(value, currentValues[key])
		}
		looksLikeTierValue := false
		if _, exists := result["up_to"]; exists {
			looksLikeTierValue = true
		} else if _, exists := currentValues["up_to"]; exists {
			looksLikeTierValue = true
		}
		if observedUpTo, exists := result["up_to"]; !exists || observedUpTo == nil {
			if currentUpTo, ok := currentValues["up_to"].(string); ok && currentUpTo == "inf" {
				result["up_to"] = "inf"
			} else if looksLikeTierValue {
				result["up_to"] = "inf"
			}
		}
		for key := range result {
			if !strings.HasSuffix(key, "_decimal") {
				continue
			}
			baseKey := strings.TrimSuffix(key, "_decimal")
			baseValue, baseConfigured := currentValues[baseKey]
			decimalValue, decimalConfigured := currentValues[key]
			if decimalConfigured && decimalValue != nil && (!baseConfigured || baseValue == nil) {
				if observedDecimal, ok := result[key].(string); ok {
					if observedBase, exists := result[baseKey]; exists {
						if observedBaseInt, ok := numberToInt64(observedBase); ok &&
							observedDecimal == strconv.FormatInt(observedBaseInt, 10) {
							delete(result, baseKey)
						}
					}
				}
			}
			if observedDecimal, ok := result[key].(string); ok {
				if observedBase, exists := result[baseKey]; exists {
					if observedBaseInt, ok := numberToInt64(observedBase); ok &&
						observedDecimal == strconv.FormatInt(observedBaseInt, 10) &&
						(!decimalConfigured || decimalValue == nil) {
						delete(result, key)
						continue
					}
				}
			}
			if baseConfigured && baseValue != nil && (!decimalConfigured || decimalValue == nil) {
				delete(result, key)
			}
		}
		if looksLikeTierValue {
			for _, amountKey := range []string{"flat_amount", "unit_amount"} {
				currentAmount, amountConfigured := currentValues[amountKey]
				if amountConfigured && currentAmount != nil {
					continue
				}
				observedAmount, exists := result[amountKey]
				if !exists {
					continue
				}
				observedAmountInt, ok := numberToInt64(observedAmount)
				if !ok || observedAmountInt != 0 {
					continue
				}
				oppositeAmountKey := "flat_amount"
				oppositeDecimalKey := "flat_amount_decimal"
				if amountKey == "flat_amount" {
					oppositeAmountKey = "unit_amount"
					oppositeDecimalKey = "unit_amount_decimal"
				}
				if oppositeAmount, exists := result[oppositeAmountKey]; exists && oppositeAmount != nil {
					delete(result, amountKey)
					continue
				}
				if oppositeDecimal, exists := result[oppositeDecimalKey]; exists && oppositeDecimal != nil {
					delete(result, amountKey)
				}
			}
		}
		return result
	default:
		return observed
	}
}

func plainDecimalStringValue(value interface{}) (string, bool) {
	switch typed := value.(type) {
	case string:
		return typed, true
	case int64:
		return strconv.FormatInt(typed, 10), true
	case int:
		return strconv.FormatInt(int64(typed), 10), true
	case float64:
		return strconv.FormatFloat(typed, 'f', -1, 64), true
	case float32:
		return strconv.FormatFloat(float64(typed), 'f', -1, 64), true
	default:
		return "", false
	}
}

func decimalStringsEquivalent(left string, right string) bool {
	leftValue, ok := new(big.Rat).SetString(left)
	if !ok {
		return false
	}
	rightValue, ok := new(big.Rat).SetString(right)
	if !ok {
		return false
	}
	return leftValue.Cmp(rightValue) == 0
}

type equivalentDecimalStringModifier struct{}

func (m equivalentDecimalStringModifier) Description(context.Context) string {
	return "Treats numerically equivalent decimal string values as unchanged."
}

func (m equivalentDecimalStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m equivalentDecimalStringModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() {
		return
	}
	if decimalStringsEquivalent(req.ConfigValue.ValueString(), req.StateValue.ValueString()) {
		resp.PlanValue = req.StateValue
	}
}

func equivalentDecimalStringPlanModifier() planmodifier.String {
	return equivalentDecimalStringModifier{}
}

func preserveEquivalentDecimalStringLeaves(observed interface{}, current interface{}) interface{} {
	if currentValue, ok := current.(string); ok {
		if observedValue, ok := plainDecimalStringValue(observed); ok &&
			decimalStringsEquivalent(observedValue, currentValue) {
			return currentValue
		}
	}
	switch observedValue := observed.(type) {
	case []interface{}:
		currentValues, _ := current.([]interface{})
		result := make([]interface{}, 0, len(observedValue))
		for index, item := range observedValue {
			var currentItem interface{}
			if index < len(currentValues) {
				currentItem = currentValues[index]
			}
			result = append(result, preserveEquivalentDecimalStringLeaves(item, currentItem))
		}
		return result
	case map[string]interface{}:
		currentValues, _ := current.(map[string]interface{})
		result := map[string]interface{}{}
		for key, value := range observedValue {
			result[key] = preserveEquivalentDecimalStringLeaves(value, currentValues[key])
		}
		return result
	default:
		return observed
	}
}

func structFieldSourceKey(field reflect.StructField) string {
	for _, rawTag := range []string{field.Tag.Get("json"), field.Tag.Get("form")} {
		if rawTag == "" {
			continue
		}
		name := strings.Split(rawTag, ",")[0]
		if name == "" || name == "-" || name == "*" {
			continue
		}
		return name
	}
	return pascalToSnakeName(field.Name)
}

func structFieldHasFormOption(field reflect.StructField, option string) bool {
	rawTag := field.Tag.Get("form")
	if rawTag == "" {
		return false
	}
	parts := strings.Split(rawTag, ",")
	for _, part := range parts[1:] {
		if part == option {
			return true
		}
	}
	return false
}

func isHighPrecisionFloatField(field reflect.StructField) bool {
	if !structFieldHasFormOption(field, "high_precision") {
		return false
	}
	fieldType := field.Type
	for fieldType.Kind() == reflect.Pointer {
		fieldType = fieldType.Elem()
	}
	return fieldType.Kind() == reflect.Float32 || fieldType.Kind() == reflect.Float64
}

func addExtraToStruct(target reflect.Value, key string, value string) bool {
	structValue := target
	for structValue.IsValid() && structValue.Kind() == reflect.Pointer {
		if structValue.IsNil() {
			return false
		}
		structValue = structValue.Elem()
	}
	if !structValue.IsValid() || structValue.Kind() != reflect.Struct || !structValue.CanAddr() {
		return false
	}
	method := structValue.Addr().MethodByName("AddExtra")
	if !method.IsValid() {
		return false
	}
	method.Call([]reflect.Value{reflect.ValueOf(key), reflect.ValueOf(value)})
	return true
}

func formEncodedExtraKey(path []string) string {
	if len(path) == 0 {
		return ""
	}
	key := path[0]
	for _, segment := range path[1:] {
		key += "[" + segment + "]"
	}
	return key
}

func structFieldDecodeKey(field reflect.StructField) string {
	rawTag := field.Tag.Get("json")
	if rawTag != "" {
		name := strings.Split(rawTag, ",")[0]
		if name != "" && name != "-" {
			return name
		}
	}
	return field.Name
}

func transformPlainForTargetType(value interface{}, targetType reflect.Type) interface{} {
	for targetType.Kind() == reflect.Pointer {
		targetType = targetType.Elem()
	}
	switch targetType.Kind() {
	case reflect.Struct:
		sourceMap, ok := value.(map[string]interface{})
		if !ok {
			return value
		}
		result := map[string]interface{}{}
		for i := 0; i < targetType.NumField(); i++ {
			field := targetType.Field(i)
			if field.PkgPath != "" {
				continue
			}
			sourceKey := structFieldSourceKey(field)
			sourceValue, exists := sourceMap[sourceKey]
			if !exists {
				continue
			}
			if typedString, ok := sourceValue.(string); ok && typedString == "inf" {
				infField, found := targetType.FieldByName(field.Name + "Inf")
				if found && infField.PkgPath == "" && infField.Type.Kind() == reflect.Pointer && infField.Type.Elem().Kind() == reflect.Bool {
					result[structFieldDecodeKey(infField)] = true
					continue
				}
			}
			if _, ok := sourceValue.(string); ok && isHighPrecisionFloatField(field) {
				result[structFieldDecodeKey(field)] = nil
				continue
			}
			result[structFieldDecodeKey(field)] = transformPlainForTargetType(sourceValue, field.Type)
		}
		return result
	case reflect.Slice, reflect.Array:
		items, ok := value.([]interface{})
		if !ok {
			return value
		}
		result := make([]interface{}, 0, len(items))
		for _, item := range items {
			result = append(result, transformPlainForTargetType(item, targetType.Elem()))
		}
		return result
	case reflect.Map:
		sourceMap, ok := value.(map[string]interface{})
		if !ok {
			keyedSource, keyedOk := plainKeyedListToMap(value, "key")
			if !keyedOk {
				return value
			}
			castMap, castOk := keyedSource.(map[string]interface{})
			if !castOk {
				return value
			}
			sourceMap = castMap
		}
		result := map[string]interface{}{}
		for key, item := range sourceMap {
			result[key] = transformPlainForTargetType(item, targetType.Elem())
		}
		return result
	case reflect.String:
		switch typed := value.(type) {
		case string:
			return typed
		case int64:
			return strconv.FormatInt(typed, 10)
		case int:
			return strconv.Itoa(typed)
		case float64:
			return strconv.FormatFloat(typed, 'f', -1, 64)
		case bool:
			return strconv.FormatBool(typed)
		default:
			return value
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch typed := value.(type) {
		case string:
			parsed, err := strconv.ParseInt(typed, 10, 64)
			if err != nil {
				return value
			}
			return parsed
		default:
			return value
		}
	default:
		return value
	}
}

func assignAttrValueToTarget(target reflect.Value, value attr.Value) {
	plain := attrValueToPlain(value)
	assignPlainValueToTarget(target, plain)
}

func collectHighPrecisionExtraValues(plain interface{}, targetType reflect.Type, path []string, extras map[string]string) {
	if plain == nil || targetType == nil {
		return
	}
	for targetType.Kind() == reflect.Pointer {
		targetType = targetType.Elem()
	}
	switch targetType.Kind() {
	case reflect.Struct:
		sourceMap, ok := plain.(map[string]interface{})
		if !ok {
			return
		}
		for i := 0; i < targetType.NumField(); i++ {
			fieldMeta := targetType.Field(i)
			if fieldMeta.PkgPath != "" {
				continue
			}
			sourceKey := structFieldSourceKey(fieldMeta)
			sourceValue, exists := sourceMap[sourceKey]
			if !exists {
				continue
			}
			nextPath := append(append([]string{}, path...), sourceKey)
			if typedString, ok := sourceValue.(string); ok && isHighPrecisionFloatField(fieldMeta) {
				extras[formEncodedExtraKey(nextPath)] = typedString
				continue
			}
			collectHighPrecisionExtraValues(sourceValue, fieldMeta.Type, nextPath, extras)
		}
	case reflect.Slice, reflect.Array:
		sourceItems, ok := plain.([]interface{})
		if !ok {
			return
		}
		for i, sourceItem := range sourceItems {
			nextPath := append(append([]string{}, path...), strconv.Itoa(i))
			collectHighPrecisionExtraValues(sourceItem, targetType.Elem(), nextPath, extras)
		}
	case reflect.Map:
		sourceMap, ok := plain.(map[string]interface{})
		if !ok {
			return
		}
		for key, sourceValue := range sourceMap {
			nextPath := append(append([]string{}, path...), key)
			collectHighPrecisionExtraValues(sourceValue, targetType.Elem(), nextPath, extras)
		}
	}
}

func addHighPrecisionExtrasToTarget(target interface{}, fieldMeta reflect.StructField, plain interface{}) bool {
	extras := map[string]string{}
	collectHighPrecisionExtraValues(plain, fieldMeta.Type, []string{structFieldSourceKey(fieldMeta)}, extras)
	if len(extras) == 0 {
		return true
	}
	targetValue := reflect.ValueOf(target)
	keys := make([]string, 0, len(extras))
	for key := range extras {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if !addExtraToStruct(targetValue, key, extras[key]) {
			return false
		}
	}
	return true
}

func assignPlainValueToTarget(target reflect.Value, plain interface{}) {
	if plain == nil {
		return
	}
	transformed := transformPlainForTargetType(plain, target.Type())
	raw, err := json.Marshal(transformed)
	if err != nil {
		return
	}
	if target.Kind() == reflect.Pointer {
		instance := reflect.New(target.Type().Elem())
		if err := json.Unmarshal(raw, instance.Interface()); err == nil {
			target.Set(instance)
		}
		return
	}
	instance := reflect.New(target.Type())
	if err := json.Unmarshal(raw, instance.Interface()); err == nil {
		target.Set(instance.Elem())
	}
}

func assignAttrValueToNamedField(target interface{}, fieldName string, value attr.Value) bool {
	structValue := reflect.ValueOf(target)
	if !structValue.IsValid() || structValue.Kind() != reflect.Pointer || structValue.IsNil() {
		return false
	}
	structValue = structValue.Elem()
	if structValue.Kind() != reflect.Struct {
		return false
	}
	fieldMeta, ok := structValue.Type().FieldByName(fieldName)
	if !ok {
		return false
	}
	field := structValue.FieldByName(fieldName)
	if !field.IsValid() || !field.CanSet() {
		return false
	}
	plain := attrValueToPlain(value)
	assignPlainValueToTarget(field, plain)
	return addHighPrecisionExtrasToTarget(target, fieldMeta, plain)
}

func assignMetadataDiffToNamedField(target interface{}, fieldName string, plan types.Map, state types.Map) bool {
	structValue := reflect.ValueOf(target)
	if !structValue.IsValid() || structValue.Kind() != reflect.Pointer || structValue.IsNil() {
		return false
	}
	structValue = structValue.Elem()
	if structValue.Kind() != reflect.Struct {
		return false
	}
	field := structValue.FieldByName(fieldName)
	if !field.IsValid() || !field.CanSet() {
		return false
	}

	nextMetadata := map[string]interface{}{}
	for key, value := range plan.Elements() {
		stringValue, ok := value.(basetypes.StringValue)
		if !ok || stringValue.IsNull() || stringValue.IsUnknown() {
			continue
		}
		nextMetadata[key] = stringValue.ValueString()
	}
	for key := range state.Elements() {
		if _, ok := nextMetadata[key]; ok {
			continue
		}
		nextMetadata[key] = ""
	}

	assignPlainValueToTarget(field, nextMetadata)
	return true
}

func assignKeyedListValueToNamedField(target interface{}, fieldName string, value attr.Value, keyField string) bool {
	plain, ok := plainKeyedListToMap(attrValueToPlain(value), keyField)
	if !ok {
		return false
	}
	structValue := reflect.ValueOf(target)
	if !structValue.IsValid() || structValue.Kind() != reflect.Pointer || structValue.IsNil() {
		return false
	}
	structValue = structValue.Elem()
	if structValue.Kind() != reflect.Struct {
		return false
	}
	fieldMeta, ok := structValue.Type().FieldByName(fieldName)
	if !ok {
		return false
	}
	field := structValue.FieldByName(fieldName)
	if !field.IsValid() || !field.CanSet() {
		return false
	}
	assignPlainValueToTarget(field, plain)
	return addHighPrecisionExtrasToTarget(target, fieldMeta, plain)
}

func assignFilePathToNamedField(target interface{}, fieldName string, filePath string) error {
	structValue := reflect.ValueOf(target)
	if !structValue.IsValid() || structValue.Kind() != reflect.Pointer || structValue.IsNil() {
		return fmt.Errorf("target must be a non-nil pointer")
	}
	structValue = structValue.Elem()
	if structValue.Kind() != reflect.Struct {
		return fmt.Errorf("target must point to a struct")
	}
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	readerField := structValue.FieldByName(fieldName + "Reader")
	if !readerField.IsValid() || !readerField.CanSet() {
		return fmt.Errorf("missing assignable upload reader field %q", fieldName+"Reader")
	}
	readerField.Set(reflect.ValueOf(bytes.NewReader(contents)))
	filenameField := structValue.FieldByName("Filename")
	if !filenameField.IsValid() || !filenameField.CanSet() {
		return fmt.Errorf("missing assignable upload filename field")
	}
	filename := filepath.Base(filePath)
	switch filenameField.Kind() {
	case reflect.String:
		filenameField.SetString(filename)
		return nil
	case reflect.Pointer:
		if filenameField.Type().Elem().Kind() == reflect.String {
			filenameField.Set(reflect.ValueOf(stripe.String(filename)))
			return nil
		}
	}
	return fmt.Errorf("unsupported upload filename field type %s", filenameField.Type())
}

func assignStringToStructIDField(field reflect.Value, value string) bool {
	target := field
	switch field.Kind() {
	case reflect.Pointer:
		if field.Type().Elem().Kind() != reflect.Struct {
			return false
		}
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		target = field.Elem()
	case reflect.Struct:
		target = field
	default:
		return false
	}
	idField := target.FieldByName("ID")
	if !idField.IsValid() || !idField.CanSet() {
		return false
	}
	switch idField.Kind() {
	case reflect.String:
		idField.SetString(value)
		return true
	case reflect.Pointer:
		if idField.Type().Elem().Kind() == reflect.String {
			idField.Set(reflect.ValueOf(stripe.String(value)))
			return true
		}
	}
	return false
}

func assignStringToNamedField(target interface{}, primary string, fallback string, value string) bool {
	structValue := reflect.ValueOf(target)
	if !structValue.IsValid() || structValue.Kind() != reflect.Pointer || structValue.IsNil() {
		return false
	}
	structValue = structValue.Elem()
	if structValue.Kind() != reflect.Struct {
		return false
	}
	for _, fieldName := range []string{primary, fallback} {
		field := structValue.FieldByName(fieldName)
		if !field.IsValid() || !field.CanSet() {
			continue
		}
		fieldMeta, ok := structValue.Type().FieldByName(fieldName)
		if ok && isHighPrecisionFloatField(fieldMeta) && addExtraToStruct(structValue, structFieldSourceKey(fieldMeta), value) {
			return true
		}
		switch field.Kind() {
		case reflect.String:
			field.SetString(value)
			return true
		case reflect.Pointer:
			if field.Type().Elem().Kind() == reflect.String {
				field.Set(reflect.ValueOf(stripe.String(value)))
				return true
			}
		}
		if assignStringToStructIDField(field, value) {
			return true
		}
	}
	return false
}

func assignStringToNamedFieldOrMethod(target interface{}, primary string, fallback string, value string) bool {
	if assignStringToNamedField(target, primary, fallback, value) {
		return true
	}
	receiver := reflect.ValueOf(target)
	if !receiver.IsValid() || receiver.Kind() != reflect.Pointer || receiver.IsNil() {
		return false
	}
	for _, fieldName := range []string{fallback, primary} {
		if fieldName == "" {
			continue
		}
		setterName := "Set" + strings.TrimSuffix(fieldName, "ID")
		setter := receiver.MethodByName(setterName)
		if !setter.IsValid() {
			continue
		}
		results := setter.Call([]reflect.Value{reflect.ValueOf(value)})
		if len(results) == 0 {
			return true
		}
		if len(results) == 1 && results[0].Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			if results[0].IsNil() {
				return true
			}
			continue
		}
		return true
	}
	return false
}

func numberToInt64(value interface{}) (int64, bool) {
	switch n := value.(type) {
	case int64:
		return n, true
	case int:
		return int64(n), true
	case float64:
		return int64(n), true
	case string:
		parsed, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			return 0, false
		}
		return parsed, true
	default:
		return 0, false
	}
}

func numberToFloat64(value interface{}) (float64, bool) {
	switch n := value.(type) {
	case float64:
		return n, true
	case float32:
		return float64(n), true
	case int64:
		return float64(n), true
	case int:
		return float64(n), true
	case string:
		parsed, err := strconv.ParseFloat(n, 64)
		if err != nil {
			return 0, false
		}
		return parsed, true
	default:
		return 0, false
	}
}

func nullTerraformValue(tfType attr.Type) (attr.Value, bool) {
	switch t := tfType.(type) {
	case basetypes.StringType:
		return types.StringNull(), true
	case basetypes.BoolType:
		return types.BoolNull(), true
	case basetypes.Int64Type:
		return types.Int64Null(), true
	case basetypes.Float64Type:
		return types.Float64Null(), true
	case types.ListType:
		return types.ListNull(t.ElemType), true
	case types.SetType:
		return types.SetNull(t.ElemType), true
	case types.MapType:
		return types.MapNull(t.ElemType), true
	case types.ObjectType:
		return types.ObjectNull(t.AttrTypes), true
	default:
		return nil, false
	}
}

func rawResponseToPlain(resource interface{}) (interface{}, bool) {
	value, ok := derefValue(reflect.ValueOf(resource))
	if !ok || value.Kind() != reflect.Struct {
		return nil, false
	}
	lastResponse := value.FieldByName("LastResponse")
	if !lastResponse.IsValid() || lastResponse.Kind() != reflect.Pointer || lastResponse.IsNil() {
		return nil, false
	}
	rawJSON := lastResponse.Elem().FieldByName("RawJSON")
	if !rawJSON.IsValid() || rawJSON.Kind() != reflect.Slice || rawJSON.Len() == 0 {
		return nil, false
	}
	if rawJSON.Type().Elem().Kind() != reflect.Uint8 {
		return nil, false
	}
	var plain interface{}
	if err := json.Unmarshal(rawJSON.Bytes(), &plain); err != nil {
		return nil, false
	}
	return plain, true
}

func stripeParamsFromValue(params interface{}) *stripe.Params {
	if params == nil {
		return &stripe.Params{}
	}
	value := reflect.ValueOf(params)
	if !value.IsValid() || value.Kind() != reflect.Pointer || value.IsNil() {
		return &stripe.Params{}
	}
	elem := value.Elem()
	if elem.Kind() != reflect.Struct {
		return &stripe.Params{}
	}
	field := elem.FieldByName("Params")
	if !field.IsValid() || !field.CanAddr() {
		return &stripe.Params{}
	}
	if stripeParams, ok := field.Addr().Interface().(*stripe.Params); ok {
		return stripeParams
	}
	return &stripe.Params{}
}

func ensureRawResponse(resource stripe.LastResponseSetter, backend stripe.Backend, key string, path string, params interface{}) error {
	if _, ok := rawResponseToPlain(resource); ok {
		return nil
	}
	if err := backend.CallRaw("GET", path, key, nil, stripeParamsFromValue(params), resource); err != nil {
		return err
	}
	if _, ok := rawResponseToPlain(resource); !ok {
		return fmt.Errorf("raw response unavailable after hydration for %T", resource)
	}
	return nil
}

func plainFromResponseField(resource interface{}, fieldName string) (interface{}, bool) {
	value, ok := derefValue(reflect.ValueOf(resource))
	if !ok || value.Kind() != reflect.Struct {
		return nil, false
	}
	field := value.FieldByName(fieldName)
	if !field.IsValid() {
		return nil, false
	}
	return reflectValueToPlain(field), true
}

func plainValueAtPath(plain interface{}, path ...string) (interface{}, bool) {
	current := plain
	for _, segment := range path {
		switch value := current.(type) {
		case map[string]interface{}:
			next, exists := value[segment]
			if !exists {
				return nil, false
			}
			current = next
		case []interface{}:
			index, err := strconv.Atoi(segment)
			if err != nil || index < 0 || index >= len(value) {
				return nil, false
			}
			current = value[index]
		default:
			return nil, false
		}
	}
	return current, true
}

func plainValueAtAnyPath(plain interface{}, paths [][]string) (interface{}, bool) {
	for _, path := range paths {
		value, ok := plainValueAtPath(plain, path...)
		if !ok || value == nil {
			continue
		}
		return value, true
	}
	return nil, false
}

func plainToStringIDValue(plain interface{}) (types.String, bool) {
	if plain == nil {
		return types.StringNull(), true
	}
	switch v := plain.(type) {
	case string:
		return types.StringValue(v), true
	case int64:
		return types.StringValue(strconv.FormatInt(v, 10)), true
	case int:
		return types.StringValue(strconv.Itoa(v)), true
	case float64:
		return types.StringValue(strconv.FormatFloat(v, 'f', -1, 64)), true
	case float32:
		return types.StringValue(strconv.FormatFloat(float64(v), 'f', -1, 32)), true
	case bool:
		return types.StringValue(strconv.FormatBool(v)), true
	case map[string]interface{}:
		rawID, exists := v["id"]
		if !exists {
			return types.String{}, false
		}
		id, ok := rawID.(string)
		if !ok {
			return types.String{}, false
		}
		return types.StringValue(id), true
	default:
		return types.String{}, false
	}
}

func nullValueForUnknown(value attr.Value) (attr.Value, bool) {
	switch v := value.(type) {
	case basetypes.StringValue:
		if !v.IsUnknown() {
			return nil, false
		}
		return types.StringNull(), true
	case basetypes.Int64Value:
		if !v.IsUnknown() {
			return nil, false
		}
		return types.Int64Null(), true
	case basetypes.Float64Value:
		if !v.IsUnknown() {
			return nil, false
		}
		return types.Float64Null(), true
	case basetypes.BoolValue:
		if !v.IsUnknown() {
			return nil, false
		}
		return types.BoolNull(), true
	case basetypes.ListValue:
		if !v.IsUnknown() {
			return nil, false
		}
		return types.ListNull(v.ElementType(context.Background())), true
	case basetypes.SetValue:
		if !v.IsUnknown() {
			return nil, false
		}
		return types.SetNull(v.ElementType(context.Background())), true
	case basetypes.MapValue:
		if !v.IsUnknown() {
			return nil, false
		}
		return types.MapNull(v.ElementType(context.Background())), true
	case basetypes.ObjectValue:
		if !v.IsUnknown() {
			return nil, false
		}
		return types.ObjectNull(v.AttributeTypes(context.Background())), true
	default:
		return nil, false
	}
}

func normalizeAttrValue(value attr.Value) (attr.Value, bool) {
	if replacement, ok := nullValueForUnknown(value); ok {
		return replacement, true
	}
	switch v := value.(type) {
	case basetypes.ListValue:
		elements := v.Elements()
		normalized := make([]attr.Value, 0, len(elements))
		changed := false
		for _, element := range elements {
			replacement, ok := normalizeAttrValue(element)
			if ok {
				normalized = append(normalized, replacement)
				changed = true
				continue
			}
			normalized = append(normalized, element)
		}
		if !changed {
			return nil, false
		}
		normalizedList, diags := types.ListValue(v.ElementType(context.Background()), normalized)
		if diags.HasError() {
			return nil, false
		}
		return normalizedList, true
	case basetypes.SetValue:
		elements := v.Elements()
		normalized := make([]attr.Value, 0, len(elements))
		changed := false
		for _, element := range elements {
			replacement, ok := normalizeAttrValue(element)
			if ok {
				normalized = append(normalized, replacement)
				changed = true
				continue
			}
			normalized = append(normalized, element)
		}
		if !changed {
			return nil, false
		}
		normalizedSet, diags := types.SetValue(v.ElementType(context.Background()), normalized)
		if diags.HasError() {
			return nil, false
		}
		return normalizedSet, true
	case basetypes.MapValue:
		normalized := map[string]attr.Value{}
		changed := false
		for key, element := range v.Elements() {
			replacement, ok := normalizeAttrValue(element)
			if ok {
				normalized[key] = replacement
				changed = true
				continue
			}
			normalized[key] = element
		}
		if !changed {
			return nil, false
		}
		normalizedMap, diags := types.MapValue(v.ElementType(context.Background()), normalized)
		if diags.HasError() {
			return nil, false
		}
		return normalizedMap, true
	case basetypes.ObjectValue:
		normalized := map[string]attr.Value{}
		changed := false
		for key, element := range v.Attributes() {
			replacement, ok := normalizeAttrValue(element)
			if ok {
				normalized[key] = replacement
				changed = true
				continue
			}
			normalized[key] = element
		}
		if !changed {
			return nil, false
		}
		normalizedObject, diags := types.ObjectValue(v.AttributeTypes(context.Background()), normalized)
		if diags.HasError() {
			return nil, false
		}
		return normalizedObject, true
	default:
		return nil, false
	}
}

func clearAttrValueAtPaths(value attr.Value, paths [][]string) (attr.Value, bool) {
	if value == nil {
		return nil, false
	}
	for _, path := range paths {
		if len(path) != 0 {
			continue
		}
		replacement, ok := nullTerraformValue(value.Type(context.Background()))
		return replacement, ok
	}
	switch v := value.(type) {
	case basetypes.ObjectValue:
		nextPaths := map[string][][]string{}
		for _, path := range paths {
			if len(path) == 0 {
				continue
			}
			head := path[0]
			nextPaths[head] = append(nextPaths[head], path[1:])
		}
		normalized := map[string]attr.Value{}
		changed := false
		for key, element := range v.Attributes() {
			childPaths, ok := nextPaths[key]
			if !ok {
				normalized[key] = element
				continue
			}
			replacement, ok := clearAttrValueAtPaths(element, childPaths)
			if ok {
				normalized[key] = replacement
				changed = true
				continue
			}
			normalized[key] = element
		}
		if !changed {
			return nil, false
		}
		normalizedObject, diags := types.ObjectValue(v.AttributeTypes(context.Background()), normalized)
		if diags.HasError() {
			return nil, false
		}
		return normalizedObject, true
	case basetypes.ListValue:
		wildcardPaths := [][]string{}
		indexedPaths := map[int][][]string{}
		for _, path := range paths {
			if len(path) == 0 {
				continue
			}
			if path[0] == "*" {
				wildcardPaths = append(wildcardPaths, path[1:])
				continue
			}
			index, err := strconv.Atoi(path[0])
			if err != nil {
				continue
			}
			indexedPaths[index] = append(indexedPaths[index], path[1:])
		}
		elements := v.Elements()
		normalized := make([]attr.Value, len(elements))
		changed := false
		for index, element := range elements {
			childPaths := append([][]string{}, wildcardPaths...)
			childPaths = append(childPaths, indexedPaths[index]...)
			if len(childPaths) == 0 {
				normalized[index] = element
				continue
			}
			replacement, ok := clearAttrValueAtPaths(element, childPaths)
			if ok {
				normalized[index] = replacement
				changed = true
				continue
			}
			normalized[index] = element
		}
		if !changed {
			return nil, false
		}
		normalizedList, diags := types.ListValue(v.ElementType(context.Background()), normalized)
		if diags.HasError() {
			return nil, false
		}
		return normalizedList, true
	case basetypes.MapValue:
		wildcardPaths := [][]string{}
		keyedPaths := map[string][][]string{}
		for _, path := range paths {
			if len(path) == 0 {
				continue
			}
			if path[0] == "*" {
				wildcardPaths = append(wildcardPaths, path[1:])
				continue
			}
			keyedPaths[path[0]] = append(keyedPaths[path[0]], path[1:])
		}
		normalized := map[string]attr.Value{}
		changed := false
		for key, element := range v.Elements() {
			childPaths := append([][]string{}, wildcardPaths...)
			childPaths = append(childPaths, keyedPaths[key]...)
			if len(childPaths) == 0 {
				normalized[key] = element
				continue
			}
			replacement, ok := clearAttrValueAtPaths(element, childPaths)
			if ok {
				normalized[key] = replacement
				changed = true
				continue
			}
			normalized[key] = element
		}
		if !changed {
			return nil, false
		}
		normalizedMap, diags := types.MapValue(v.ElementType(context.Background()), normalized)
		if diags.HasError() {
			return nil, false
		}
		return normalizedMap, true
	case basetypes.SetValue:
		return nil, false
	default:
		return nil, false
	}
}

func hydrateAttrValueAtPaths(target attr.Value, source attr.Value, paths [][]string) (attr.Value, bool) {
	if source == nil || source.IsNull() || source.IsUnknown() {
		return nil, false
	}
	for _, path := range paths {
		if len(path) == 0 {
			return source, true
		}
	}
	if target == nil || target.IsNull() || target.IsUnknown() {
		return source, true
	}
	switch targetValue := target.(type) {
	case basetypes.ObjectValue:
		sourceValue, ok := source.(basetypes.ObjectValue)
		if !ok {
			return source, true
		}
		nextPaths := map[string][][]string{}
		for _, path := range paths {
			if len(path) == 0 {
				continue
			}
			head := path[0]
			nextPaths[head] = append(nextPaths[head], path[1:])
		}
		normalized := map[string]attr.Value{}
		changed := false
		targetAttrs := targetValue.Attributes()
		sourceAttrs := sourceValue.Attributes()
		for key, element := range targetAttrs {
			childPaths, ok := nextPaths[key]
			if !ok {
				normalized[key] = element
				continue
			}
			sourceElement, ok := sourceAttrs[key]
			if !ok {
				normalized[key] = element
				continue
			}
			replacement, ok := hydrateAttrValueAtPaths(element, sourceElement, childPaths)
			if ok {
				normalized[key] = replacement
				changed = true
				continue
			}
			normalized[key] = element
		}
		if !changed {
			return nil, false
		}
		normalizedObject, diags := types.ObjectValue(targetValue.AttributeTypes(context.Background()), normalized)
		if diags.HasError() {
			return nil, false
		}
		return normalizedObject, true
	case basetypes.ListValue:
		sourceValue, ok := source.(basetypes.ListValue)
		if !ok {
			return source, true
		}
		targetElements := targetValue.Elements()
		sourceElements := sourceValue.Elements()
		if len(targetElements) != len(sourceElements) {
			return source, true
		}
		wildcardPaths := [][]string{}
		indexedPaths := map[int][][]string{}
		for _, path := range paths {
			if len(path) == 0 {
				continue
			}
			if path[0] == "*" {
				wildcardPaths = append(wildcardPaths, path[1:])
				continue
			}
			index, err := strconv.Atoi(path[0])
			if err != nil {
				continue
			}
			indexedPaths[index] = append(indexedPaths[index], path[1:])
		}
		normalized := make([]attr.Value, len(targetElements))
		changed := false
		for index, element := range targetElements {
			childPaths := append([][]string{}, wildcardPaths...)
			childPaths = append(childPaths, indexedPaths[index]...)
			if len(childPaths) == 0 {
				normalized[index] = element
				continue
			}
			replacement, ok := hydrateAttrValueAtPaths(element, sourceElements[index], childPaths)
			if ok {
				normalized[index] = replacement
				changed = true
				continue
			}
			normalized[index] = element
		}
		if !changed {
			return nil, false
		}
		normalizedList, diags := types.ListValue(targetValue.ElementType(context.Background()), normalized)
		if diags.HasError() {
			return nil, false
		}
		return normalizedList, true
	case basetypes.MapValue:
		sourceValue, ok := source.(basetypes.MapValue)
		if !ok {
			return source, true
		}
		targetElements := targetValue.Elements()
		sourceElements := sourceValue.Elements()
		wildcardPaths := [][]string{}
		keyedPaths := map[string][][]string{}
		for _, path := range paths {
			if len(path) == 0 {
				continue
			}
			if path[0] == "*" {
				wildcardPaths = append(wildcardPaths, path[1:])
				continue
			}
			keyedPaths[path[0]] = append(keyedPaths[path[0]], path[1:])
		}
		normalized := map[string]attr.Value{}
		changed := false
		for key, element := range targetElements {
			sourceElement, ok := sourceElements[key]
			if !ok {
				normalized[key] = element
				continue
			}
			childPaths := append([][]string{}, wildcardPaths...)
			childPaths = append(childPaths, keyedPaths[key]...)
			if len(childPaths) == 0 {
				normalized[key] = element
				continue
			}
			replacement, ok := hydrateAttrValueAtPaths(element, sourceElement, childPaths)
			if ok {
				normalized[key] = replacement
				changed = true
				continue
			}
			normalized[key] = element
		}
		if !changed {
			return nil, false
		}
		normalizedMap, diags := types.MapValue(targetValue.ElementType(context.Background()), normalized)
		if diags.HasError() {
			return nil, false
		}
		return normalizedMap, true
	case basetypes.SetValue:
		return source, true
	default:
		return nil, false
	}
}

func clearWriteOnlyPaths(model interface{}, paths [][]string) {
	value := reflect.ValueOf(model)
	if value.Kind() != reflect.Pointer || value.IsNil() {
		return
	}
	structValue := value.Elem()
	if structValue.Kind() != reflect.Struct {
		return
	}
	pathsByField := map[string][][]string{}
	for _, path := range paths {
		if len(path) == 0 {
			continue
		}
		fieldName := path[0]
		pathsByField[fieldName] = append(pathsByField[fieldName], path[1:])
	}
	for fieldName, childPaths := range pathsByField {
		field := structValue.FieldByName(fieldName)
		if !field.IsValid() || !field.CanSet() {
			continue
		}
		currentValue, ok := field.Interface().(attr.Value)
		if !ok {
			continue
		}
		replacement, ok := clearAttrValueAtPaths(currentValue, childPaths)
		if !ok {
			continue
		}
		replacementValue := reflect.ValueOf(replacement)
		if replacementValue.Type().AssignableTo(field.Type()) {
			field.Set(replacementValue)
			continue
		}
		if replacementValue.Type().ConvertibleTo(field.Type()) {
			field.Set(replacementValue.Convert(field.Type()))
		}
	}
}

func hydrateWriteOnlyPaths(target interface{}, source interface{}, paths [][]string) {
	targetValue := reflect.ValueOf(target)
	sourceValue := reflect.ValueOf(source)
	if targetValue.Kind() != reflect.Pointer || targetValue.IsNil() || sourceValue.Kind() != reflect.Pointer || sourceValue.IsNil() {
		return
	}
	targetStruct := targetValue.Elem()
	sourceStruct := sourceValue.Elem()
	if targetStruct.Kind() != reflect.Struct || sourceStruct.Kind() != reflect.Struct {
		return
	}
	pathsByField := map[string][][]string{}
	for _, path := range paths {
		if len(path) == 0 {
			continue
		}
		fieldName := path[0]
		pathsByField[fieldName] = append(pathsByField[fieldName], path[1:])
	}
	for fieldName, childPaths := range pathsByField {
		targetField := targetStruct.FieldByName(fieldName)
		sourceField := sourceStruct.FieldByName(fieldName)
		if !targetField.IsValid() || !targetField.CanSet() || !sourceField.IsValid() {
			continue
		}
		targetAttr, ok := targetField.Interface().(attr.Value)
		if !ok {
			continue
		}
		sourceAttr, ok := sourceField.Interface().(attr.Value)
		if !ok {
			continue
		}
		replacement, ok := hydrateAttrValueAtPaths(targetAttr, sourceAttr, childPaths)
		if !ok {
			continue
		}
		replacementValue := reflect.ValueOf(replacement)
		if replacementValue.Type().AssignableTo(targetField.Type()) {
			targetField.Set(replacementValue)
			continue
		}
		if replacementValue.Type().ConvertibleTo(targetField.Type()) {
			targetField.Set(replacementValue.Convert(targetField.Type()))
		}
	}
}

func normalizeUnknownValues(model interface{}) {
	value := reflect.ValueOf(model)
	if value.Kind() != reflect.Pointer || value.IsNil() {
		return
	}
	structValue := value.Elem()
	if structValue.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		if !field.CanSet() {
			continue
		}
		currentValue, ok := field.Interface().(attr.Value)
		if !ok {
			continue
		}
		replacement, ok := normalizeAttrValue(currentValue)
		if !ok {
			continue
		}
		replacementValue := reflect.ValueOf(replacement)
		if replacementValue.Type().AssignableTo(field.Type()) {
			field.Set(replacementValue)
			continue
		}
		if replacementValue.Type().ConvertibleTo(field.Type()) {
			field.Set(replacementValue.Convert(field.Type()))
		}
	}
}

func paramsHaveValues(value interface{}) bool {
	reflected := reflect.ValueOf(value)
	if !reflected.IsValid() {
		return false
	}
	if reflected.Kind() == reflect.Pointer {
		if reflected.IsNil() {
			return false
		}
		reflected = reflected.Elem()
	}
	return reflectValueHasValues(reflected)
}

func reflectValueHasValues(value reflect.Value) bool {
	if !value.IsValid() {
		return false
	}
	switch value.Kind() {
	case reflect.Interface:
		if value.IsNil() {
			return false
		}
		return reflectValueHasValues(value.Elem())
	case reflect.Pointer:
		return !value.IsNil()
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String:
		return !value.IsZero()
	case reflect.Slice, reflect.Array, reflect.Map:
		return value.Len() > 0
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			fieldInfo := value.Type().Field(i)
			if fieldInfo.PkgPath != "" {
				continue
			}
			if reflectValueHasValues(value.Field(i)) {
				return true
			}
		}
		return false
	default:
		return !value.IsZero()
	}
}

func plainPathString(path []string) string {
	if len(path) == 0 {
		return "<root>"
	}
	return strings.Join(path, ".")
}

func appendPlainPath(path []string, segment string) []string {
	next := make([]string, 0, len(path)+1)
	next = append(next, path...)
	next = append(next, segment)
	return next
}

func plainToTerraformValue(plain interface{}, tfType attr.Type) (attr.Value, error) {
	return plainToTerraformValueAtPath(plain, tfType, nil)
}

func plainToTerraformValueAtPath(plain interface{}, tfType attr.Type, path []string) (attr.Value, error) {
	if plain == nil {
		value, ok := nullTerraformValue(tfType)
		if !ok {
			return nil, fmt.Errorf("unsupported terraform type %T at %q", tfType, plainPathString(path))
		}
		return value, nil
	}
	switch t := tfType.(type) {
	case basetypes.StringType:
		if len(path) > 0 && path[len(path)-1] == "up_to" {
			if numericValue, ok := numberToInt64(plain); ok && numericValue == 0 {
				return types.StringValue("inf"), nil
			}
		}
		if value, ok := plainToStringIDValue(plain); ok {
			return value, nil
		}
		return types.StringValue(fmt.Sprintf("%v", plain)), nil
	case basetypes.BoolType:
		value, ok := plain.(bool)
		if !ok {
			return nil, fmt.Errorf("expected bool at %q, got %T", plainPathString(path), plain)
		}
		return types.BoolValue(value), nil
	case basetypes.Int64Type:
		value, ok := numberToInt64(plain)
		if !ok {
			return nil, fmt.Errorf("expected int64-compatible value at %q, got %T", plainPathString(path), plain)
		}
		return types.Int64Value(value), nil
	case basetypes.Float64Type:
		value, ok := numberToFloat64(plain)
		if !ok {
			return nil, fmt.Errorf("expected float64-compatible value at %q, got %T", plainPathString(path), plain)
		}
		return types.Float64Value(value), nil
	case types.ListType:
		collectionPlain := extractListObjectData(plain)
		items, ok := collectionPlain.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected list at %q, got %T", plainPathString(path), collectionPlain)
		}
		result := make([]attr.Value, 0, len(items))
		for index, item := range items {
			elem, err := plainToTerraformValueAtPath(item, t.ElemType, appendPlainPath(path, strconv.Itoa(index)))
			if err != nil {
				return nil, err
			}
			result = append(result, elem)
		}
		value, diags := types.ListValue(t.ElemType, result)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to construct list at %q: %s", plainPathString(path), diags.Errors()[0].Detail())
		}
		return value, nil
	case types.SetType:
		collectionPlain := extractListObjectData(plain)
		items, ok := collectionPlain.([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected set-like list at %q, got %T", plainPathString(path), collectionPlain)
		}
		result := make([]attr.Value, 0, len(items))
		for index, item := range items {
			elem, err := plainToTerraformValueAtPath(item, t.ElemType, appendPlainPath(path, strconv.Itoa(index)))
			if err != nil {
				return nil, err
			}
			result = append(result, elem)
		}
		value, diags := types.SetValue(t.ElemType, result)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to construct set at %q: %s", plainPathString(path), diags.Errors()[0].Detail())
		}
		return value, nil
	case types.MapType:
		sourceMap, ok := plain.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected map at %q, got %T", plainPathString(path), plain)
		}
		result := map[string]attr.Value{}
		keys := make([]string, 0, len(sourceMap))
		for key := range sourceMap {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			elem, err := plainToTerraformValueAtPath(sourceMap[key], t.ElemType, appendPlainPath(path, key))
			if err != nil {
				return nil, err
			}
			result[key] = elem
		}
		value, diags := types.MapValue(t.ElemType, result)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to construct map at %q: %s", plainPathString(path), diags.Errors()[0].Detail())
		}
		return value, nil
	case types.ObjectType:
		sourceMap, ok := plain.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("expected object at %q, got %T", plainPathString(path), plain)
		}
		result := map[string]attr.Value{}
		keys := make([]string, 0, len(t.AttrTypes))
		for key := range t.AttrTypes {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			nestedType := t.AttrTypes[key]
			nestedPlain, exists := sourceMap[key]
			if !exists {
				nestedValue, ok := nullTerraformValue(nestedType)
				if !ok {
					return nil, fmt.Errorf("unsupported terraform type %T at %q", nestedType, plainPathString(appendPlainPath(path, key)))
				}
				result[key] = nestedValue
				continue
			}
			nestedValue, err := plainToTerraformValueAtPath(nestedPlain, nestedType, appendPlainPath(path, key))
			if err != nil {
				return nil, err
			}
			result[key] = nestedValue
		}
		value, diags := types.ObjectValue(t.AttrTypes, result)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to construct object at %q: %s", plainPathString(path), diags.Errors()[0].Detail())
		}
		return value, nil
	default:
		return nil, fmt.Errorf("unsupported terraform type %T at %q", tfType, plainPathString(path))
	}
}

func flattenPlainValue(plain interface{}, tfType attr.Type, attrName string, source string) (attr.Value, error) {
	value, err := plainToTerraformValue(plain, tfType)
	if err != nil {
		return nil, fmt.Errorf("failed to flatten attribute %q from %s: %w", attrName, source, err)
	}
	return value, nil
}
