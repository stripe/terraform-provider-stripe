//
// File generated from our OpenAPI spec
//

package bindingruntime

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestExactJSONRoundTrip(t *testing.T) {
	for _, source := range []string{`9007199254740993`, `0.000000000000000000000000000123456789`, `123456789012345678901234567890.125`} {
		decoded, err := DecodeJSON([]byte(source))
		if err != nil {
			t.Fatal(err)
		}
		value, err := FromJSON(context.Background(), types.NumberType, decoded, nil)
		if err != nil {
			t.Fatal(err)
		}
		wire, err := JSONValue(value, false)
		if err != nil {
			t.Fatal(err)
		}
		got, err := json.Marshal(wire)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != source {
			t.Fatalf("number rounded: %s -> %s", source, got)
		}
	}
}

func TestUnknownNullEmptyAndOpaqueValuesStayDistinct(t *testing.T) {
	ctx := context.Background()
	object := types.ObjectValueMust(map[string]attr.Type{"future": types.StringType, "empty": types.ListType{ElemType: types.StringType}, "absent": types.StringType}, map[string]attr.Value{"future": types.StringUnknown(), "empty": types.ListValueMust(types.StringType, []attr.Value{}), "absent": types.StringNull()})
	decoded, err := JSONValue(object, true)
	if err != nil {
		t.Fatal(err)
	}
	fields := decoded.(map[string]any)
	if !IsUnknown(fields["future"]) || fields["absent"] != nil || len(fields["empty"].([]any)) != 0 {
		t.Fatalf("value distinctions lost: %#v", fields)
	}
	if _, err := JSONValue(object, false); err == nil {
		t.Fatal("unknown was accepted for an API write")
	}
	if _, err := json.Marshal(decoded); err == nil {
		t.Fatal("unknown silently serialized")
	}
	opaque := Literal(`{"new_field":{"amount":9007199254740993,"null":null,"items":[]}}`)
	dynamic, err := FromJSON(ctx, types.DynamicType, opaque, nil)
	if err != nil {
		t.Fatal(err)
	}
	actual, err := JSONValue(dynamic, false)
	if err != nil {
		t.Fatal(err)
	}
	wire, _ := json.Marshal(actual)
	if string(wire) != `{"new_field":{"amount":9007199254740993,"items":[],"null":null}}` {
		t.Fatalf("opaque fields lost: %s", wire)
	}
	if _, err := FromJSON(ctx, types.NumberType, float64(1), nil); err == nil {
		t.Fatal("accepted a lossy numeric conversion")
	}
}

func TestImportIdentityAndTypedKeys(t *testing.T) {
	ctx := context.Background()
	values := EmptyValues(types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "customer_id": types.StringType}})
	for _, bad := range []string{"", "a", "/b", "a/"} {
		if err := ImportIdentity(ctx, values, bad, []string{"customer_id", "id"}, "/"); err == nil {
			t.Fatalf("accepted invalid composite identity %q", bad)
		}
	}
	if err := ImportIdentity(ctx, values, "customer/contract", []string{"customer_id", "id"}, "/"); err != nil {
		t.Fatal(err)
	}
	if values.String("id") != "contract" || values.String("customer_id") != "customer" {
		t.Fatal("composite identity changed")
	}
	kind := types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType}}
	if err := CheckObjectKeys(kind, Literal(`{"id":"x","opaque":true}`)); err == nil {
		t.Fatal("unknown typed import field was silently dropped")
	}
}
