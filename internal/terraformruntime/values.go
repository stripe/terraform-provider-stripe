//
// File generated from our OpenAPI spec
//

// Package bindingruntime implements service-independent Terraform value and lifecycle policies.
package bindingruntime

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"math/big"
	"strconv"
)

type Unknown struct{}

func (Unknown) MarshalJSON() ([]byte, error) {
	return nil, errors.New("API request contains an unknown value")
}
func IsUnknown(v any) bool { _, ok := v.(Unknown); return ok }

// JSONValue preserves unknowns during validation and exact decimal tokens on the wire.
func JSONValue(v attr.Value, allowUnknown bool) (any, error) {
	if v == nil || v.IsNull() {
		return nil, nil
	}
	if v.IsUnknown() {
		if allowUnknown {
			return Unknown{}, nil
		}
		return nil, errors.New("API request contains an unknown value")
	}
	switch v := v.(type) {
	case types.Dynamic:
		return JSONValue(v.UnderlyingValue(), allowUnknown)
	case types.String:
		return v.ValueString(), nil
	case types.Bool:
		return v.ValueBool(), nil
	case types.Number:
		return json.Number(v.ValueBigFloat().Text('f', -1)), nil
	case types.Int64:
		return json.Number(strconv.FormatInt(v.ValueInt64(), 10)), nil
	case types.Object:
		return jsonValuesMap(v.Attributes(), allowUnknown)
	case types.Map:
		return jsonValuesMap(v.Elements(), allowUnknown)
	case types.List:
		return jsonValuesList(v.Elements(), allowUnknown)
	case types.Set:
		return jsonValuesList(v.Elements(), allowUnknown)
	case types.Tuple:
		return jsonValuesList(v.Elements(), allowUnknown)
	default:
		return nil, fmt.Errorf("unreviewed Terraform JSON type %T", v)
	}
}
func jsonValuesMap(values map[string]attr.Value, unknown bool) (map[string]any, error) {
	result := map[string]any{}
	for key, v := range values {
		x, err := JSONValue(v, unknown)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", key, err)
		}
		result[key] = x
	}
	return result, nil
}
func jsonValuesList(values []attr.Value, unknown bool) ([]any, error) {
	result := make([]any, len(values))
	for i, v := range values {
		x, err := JSONValue(v, unknown)
		if err != nil {
			return nil, fmt.Errorf("item %d: %w", i, err)
		}
		result[i] = x
	}
	return result, nil
}
func DecodeJSON(raw []byte) (any, error) {
	if !json.Valid(raw) {
		return nil, errors.New("invalid API JSON")
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	var v any
	if err := decoder.Decode(&v); err != nil {
		return nil, errors.New("invalid API JSON")
	}
	return v, nil
}
func Null(t attr.Type) attr.Value {
	v, err := t.ValueFromTerraform(context.Background(), tftNull(t))
	if err != nil {
		panic(err)
	}
	return v
}

// Values holds schema-typed values; absent JSON fields are distinct from explicit nulls.
type Values struct{ Object types.Object }

func NewValues(v types.Object) *Values      { return &Values{Object: v} }
func (v *Values) Get(key string) attr.Value { return v.Object.Attributes()[key] }
func (v *Values) String(key string) string  { x, _ := v.Get(key).(types.String); return x.ValueString() }
func (v *Values) Known(key string) bool {
	x := v.Get(key)
	return x != nil && !x.IsNull() && !x.IsUnknown()
}
func (v *Values) JSON(unknown bool) (map[string]any, error) {
	return jsonValuesMap(v.Object.Attributes(), unknown)
}
func (v *Values) Set(key string, value attr.Value) error {
	attrs := v.Object.Attributes()
	t, ok := v.Object.AttributeTypes(context.Background())[key]
	if !ok {
		return fmt.Errorf("unknown state attribute %s", key)
	}
	if !value.Type(context.Background()).Equal(t) {
		return fmt.Errorf("state type mismatch for %s", key)
	}
	attrs[key] = value
	object, diags := types.ObjectValue(v.Object.AttributeTypes(context.Background()), attrs)
	if diags.HasError() {
		return errors.New("invalid Terraform state value")
	}
	v.Object = object
	return nil
}
func (v *Values) SetJSON(ctx context.Context, key string, value any) error {
	t, ok := v.Object.AttributeTypes(ctx)[key]
	if !ok {
		return fmt.Errorf("unknown state attribute %s", key)
	}
	converted, err := FromJSON(ctx, t, value, v.Get(key))
	if err != nil {
		return fmt.Errorf("%s: %w", key, err)
	}
	return v.Set(key, converted)
}
func (v *Values) ResolveUnknowns() error {
	for key, value := range v.Object.Attributes() {
		if value.IsUnknown() {
			if err := v.Set(key, Null(value.Type(context.Background()))); err != nil {
				return err
			}
		}
	}
	return nil
}
func Changed(current, prior *Values, keys ...string) bool {
	if prior == nil {
		return true
	}
	for _, key := range keys {
		a, b := current.Get(key), prior.Get(key)
		if a == nil || b == nil || !a.Equal(b) {
			return true
		}
	}
	return false
}
func at(value any, path []string) (any, bool) {
	for _, key := range path {
		switch object := value.(type) {
		case map[string]any:
			var ok bool
			value, ok = object[key]
			if !ok {
				return nil, false
			}
		case []any:
			i, err := strconv.Atoi(key)
			if err != nil || i < 0 || i >= len(object) {
				return nil, false
			}
			value = object[i]
		default:
			return nil, false
		}
	}
	return value, true
}
func At(value any, path ...string) (any, bool) { return at(value, path) }
func Put(root map[string]any, path []string, value any) error {
	if len(path) == 0 {
		return errors.New("empty mapping path")
	}
	for _, key := range path[:len(path)-1] {
		child, ok := root[key]
		if !ok {
			child = map[string]any{}
			root[key] = child
		}
		var valid bool
		root, valid = child.(map[string]any)
		if !valid {
			return errors.New("overlapping request mappings")
		}
	}
	root[path[len(path)-1]] = value
	return nil
}
func StripNulls(value any) any {
	switch v := value.(type) {
	case map[string]any:
		r := map[string]any{}
		for k, x := range v {
			if x != nil {
				r[k] = StripNulls(x)
			}
		}
		return r
	case []any:
		r := make([]any, len(v))
		for i, x := range v {
			r[i] = StripNulls(x)
		}
		return r
	default:
		return value
	}
}

// FromJSON uses the declared shape for structured attributes and the configured
// shape for Dynamic. Extra response fields never silently rewrite user config.
func FromJSON(ctx context.Context, t attr.Type, value any, prior attr.Value) (attr.Value, error) {
	if value == nil {
		return Null(t), nil
	}
	switch t := t.(type) {
	case basetypes.StringType:
		s, ok := value.(string)
		if !ok {
			return nil, errors.New("expected string")
		}
		return types.StringValue(s), nil
	case basetypes.BoolType:
		b, ok := value.(bool)
		if !ok {
			return nil, errors.New("expected boolean")
		}
		return types.BoolValue(b), nil
	case basetypes.NumberType:
		n, ok := value.(json.Number)
		if !ok {
			return nil, errors.New("expected exact number")
		}
		f, _, err := big.ParseFloat(string(n), 10, 512, big.ToNearestEven)
		if err != nil {
			return nil, errors.New("invalid number")
		}
		return types.NumberValue(f), nil
	case basetypes.Int64Type:
		n, ok := value.(json.Number)
		if !ok {
			return nil, errors.New("expected integer")
		}
		i, err := n.Int64()
		if err != nil {
			return nil, errors.New("integer outside signed 64-bit range")
		}
		return types.Int64Value(i), nil
	case basetypes.DynamicType:
		if p, ok := prior.(types.Dynamic); ok && !p.IsNull() && !p.IsUnknown() && p.UnderlyingValue() != nil && !p.UnderlyingValue().IsNull() && !p.UnderlyingValue().IsUnknown() {
			child, err := FromJSON(ctx, p.UnderlyingValue().Type(ctx), value, p.UnderlyingValue())
			if err != nil {
				return nil, err
			}
			return types.DynamicValue(child), nil
		}
		child, err := InferJSON(ctx, value)
		if err != nil {
			return nil, err
		}
		return types.DynamicValue(child), nil
	case types.ObjectType:
		obj, ok := value.(map[string]any)
		if !ok {
			return nil, errors.New("expected object")
		}
		attrs := map[string]attr.Value{}
		old := map[string]attr.Value{}
		if p, ok := prior.(types.Object); ok {
			old = p.Attributes()
		}
		for key, childType := range t.AttrTypes {
			child, err := FromJSON(ctx, childType, obj[key], old[key])
			if err != nil {
				return nil, fmt.Errorf("%s: %w", key, err)
			}
			attrs[key] = child
		}
		v, d := types.ObjectValue(t.AttrTypes, attrs)
		return checkedValue(v, d)
	case types.MapType:
		obj, ok := value.(map[string]any)
		if !ok {
			return nil, errors.New("expected map")
		}
		attrs := map[string]attr.Value{}
		for key, x := range obj {
			child, err := FromJSON(ctx, t.ElemType, x, nil)
			if err != nil {
				return nil, err
			}
			attrs[key] = child
		}
		v, d := types.MapValue(t.ElemType, attrs)
		return checkedValue(v, d)
	case types.ListType:
		children, err := fromJSONList(ctx, t.ElemType, value)
		if err != nil {
			return nil, err
		}
		v, d := types.ListValue(t.ElemType, children)
		return checkedValue(v, d)
	case types.SetType:
		children, err := fromJSONList(ctx, t.ElemType, value)
		if err != nil {
			return nil, err
		}
		v, d := types.SetValue(t.ElemType, children)
		return checkedValue(v, d)
	case types.TupleType:
		list, ok := value.([]any)
		if !ok || len(list) != len(t.ElemTypes) {
			return nil, errors.New("response tuple differs from configured shape")
		}
		children := make([]attr.Value, len(list))
		for i, x := range list {
			v, err := FromJSON(ctx, t.ElemTypes[i], x, nil)
			if err != nil {
				return nil, err
			}
			children[i] = v
		}
		v, d := types.TupleValue(t.ElemTypes, children)
		return checkedValue(v, d)
	default:
		return nil, fmt.Errorf("unreviewed response type %T", t)
	}
}
func checkedValue(v attr.Value, d diag.Diagnostics) (attr.Value, error) {
	if d.HasError() {
		return nil, errors.New("invalid typed API response")
	}
	return v, nil
}
func fromJSONList(ctx context.Context, t attr.Type, value any) ([]attr.Value, error) {
	list, ok := value.([]any)
	if !ok {
		return nil, errors.New("expected array")
	}
	children := make([]attr.Value, len(list))
	for i, x := range list {
		v, err := FromJSON(ctx, t, x, nil)
		if err != nil {
			return nil, err
		}
		children[i] = v
	}
	return children, nil
}
func InferJSON(ctx context.Context, value any) (attr.Value, error) {
	switch v := value.(type) {
	case nil:
		return types.DynamicNull(), nil
	case string:
		return types.StringValue(v), nil
	case bool:
		return types.BoolValue(v), nil
	case json.Number:
		return FromJSON(ctx, types.NumberType, v, nil)
	case map[string]any:
		attrs := map[string]attr.Value{}
		ts := map[string]attr.Type{}
		for key, x := range v {
			child, err := InferJSON(ctx, x)
			if err != nil {
				return nil, err
			}
			attrs[key] = child
			ts[key] = child.Type(ctx)
		}
		o, d := types.ObjectValue(ts, attrs)
		return checkedValue(o, d)
	case []any:
		attrs := make([]attr.Value, len(v))
		ts := make([]attr.Type, len(v))
		for i, x := range v {
			child, err := InferJSON(ctx, x)
			if err != nil {
				return nil, err
			}
			attrs[i] = child
			ts[i] = child.Type(ctx)
		}
		o, d := types.TupleValue(ts, attrs)
		return checkedValue(o, d)
	default:
		return nil, errors.New("response contains non-JSON value")
	}
}

// CheckObjectKeys prevents typed imports from silently dropping opaque or new API fields.
func CheckObjectKeys(t attr.Type, value any) error {
	if value == nil {
		return nil
	}
	switch t := t.(type) {
	case types.ObjectType:
		root, ok := value.(map[string]any)
		if !ok {
			return errors.New("typed input must be an object")
		}
		for key, x := range root {
			kind, exists := t.AttrTypes[key]
			if !exists {
				return fmt.Errorf("unknown typed input field %s", key)
			}
			if err := CheckObjectKeys(kind, x); err != nil {
				return err
			}
		}
	case types.ListType:
		for _, x := range toJSONArray(value) {
			if err := CheckObjectKeys(t.ElemType, x); err != nil {
				return err
			}
		}
	case types.SetType:
		for _, x := range toJSONArray(value) {
			if err := CheckObjectKeys(t.ElemType, x); err != nil {
				return err
			}
		}
	}
	return nil
}
func toJSONArray(value any) []any { list, _ := value.([]any); return list }
