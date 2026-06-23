//
// File generated from our OpenAPI spec
//

package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"reflect"
	"strconv"
	"strings"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &CouponResource{}

var _ resource.ResourceWithConfigure = &CouponResource{}

var _ resource.ResourceWithImportState = &CouponResource{}

func NewCouponResource() resource.Resource {
	return &CouponResource{}
}

type CouponResource struct {
	client *stripe.Client
}

type CouponResourceModel struct {
	Object           types.String  `tfsdk:"object"`
	AmountOff        types.Int64   `tfsdk:"amount_off"`
	AppliesTo        types.Object  `tfsdk:"applies_to"`
	Created          types.Int64   `tfsdk:"created"`
	Currency         types.String  `tfsdk:"currency"`
	CurrencyOptions  types.List    `tfsdk:"currency_options"`
	Duration         types.String  `tfsdk:"duration"`
	DurationInMonths types.Int64   `tfsdk:"duration_in_months"`
	ID               types.String  `tfsdk:"id"`
	Livemode         types.Bool    `tfsdk:"livemode"`
	MaxRedemptions   types.Int64   `tfsdk:"max_redemptions"`
	Metadata         types.Map     `tfsdk:"metadata"`
	Name             types.String  `tfsdk:"name"`
	PercentOff       types.Float64 `tfsdk:"percent_off"`
	RedeemBy         types.Int64   `tfsdk:"redeem_by"`
	TimesRedeemed    types.Int64   `tfsdk:"times_redeemed"`
	Valid            types.Bool    `tfsdk:"valid"`
}

func (r *CouponResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stripe.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *stripe.Client, got: %T", req.ProviderData))
		return
	}

	r.client = client
}

func (r *CouponResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_coupon"
}

var _ resource.ResourceWithUpgradeState = &CouponResource{}

func (r *CouponResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: couponResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior CouponResourceV0Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeCouponStateV0(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func couponResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "A coupon contains information about a percent-off or amount-off discount you\nmight want to apply to a customer. Coupons may be applied to [subscriptions](https://api.stripe.com#subscriptions), [invoices](https://api.stripe.com#invoices),\n[checkout sessions](https://docs.stripe.com/api/checkout/sessions), [quotes](https://api.stripe.com#quotes), and more. Coupons do not work with conventional one-off [charges](/api/charges/create) or [payment intents](https://docs.stripe.com/api/payment_intents).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("coupon")},
			},
			"amount_off": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "If `amount_off` has been set, the three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the amount to take off.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"duration": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "One of `forever`, `once`, or `repeating`. Describes how long a customer who applies this coupon will get the discount.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("forever", "once", "repeating")},
			},
			"duration_in_months": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "If `duration` is `repeating`, the number of months the coupon applies. Null if coupon `duration` is `forever` or `once`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"max_redemptions": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Name of the coupon displayed to customers on for instance invoices or receipts.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"percent_off": schema.Float64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percent_off of 50 will make a $ (or local equivalent)100 invoice $ (or local equivalent)50 instead.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown(), float64planmodifier.RequiresReplace()},
			},
			"redeem_by": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Date after which the coupon can no longer be redeemed.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"times_redeemed": schema.Int64Attribute{
				Computed:      true,
				Description:   "Number of times this coupon has been applied to a customer.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"valid": schema.BoolAttribute{
				Computed:      true,
				Description:   "Taking account of the above properties, whether this coupon can still be applied to a customer.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
		},
		Blocks: map[string]schema.Block{
			"applies_to": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"products": schema.ListAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "A list of product IDs this coupon applies to",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
							ElementType:   types.StringType,
						},
					},
				},
			},
			"currency_options": schema.ListNestedBlock{
				Description: "Coupons defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required:    true,
							Description: "Key for this entry.",
						},
						"amount_off": schema.Int64Attribute{
							Required:    true,
							Description: "Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.",
						},
					},
				},
			},
		},
	}
}

type CouponResourceV0Model struct {
	Object           types.String  `tfsdk:"object"`
	AmountOff        types.Int64   `tfsdk:"amount_off"`
	AppliesTo        types.List    `tfsdk:"applies_to"`
	Created          types.Int64   `tfsdk:"created"`
	Currency         types.String  `tfsdk:"currency"`
	CurrencyOptions  types.List    `tfsdk:"currency_options"`
	Duration         types.String  `tfsdk:"duration"`
	DurationInMonths types.Int64   `tfsdk:"duration_in_months"`
	ID               types.String  `tfsdk:"id"`
	Livemode         types.Bool    `tfsdk:"livemode"`
	MaxRedemptions   types.Int64   `tfsdk:"max_redemptions"`
	Metadata         types.Map     `tfsdk:"metadata"`
	Name             types.String  `tfsdk:"name"`
	PercentOff       types.Float64 `tfsdk:"percent_off"`
	RedeemBy         types.Int64   `tfsdk:"redeem_by"`
	TimesRedeemed    types.Int64   `tfsdk:"times_redeemed"`
	Valid            types.Bool    `tfsdk:"valid"`
}

type couponStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]couponStateUpgradeAttrMeta
}

var couponStateUpgradeRootMeta = map[string]couponStateUpgradeAttrMeta{"object": couponStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "amount_off": couponStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "applies_to": couponStateUpgradeAttrMeta{AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{"products": types.ListType{ElemType: types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", PreserveConfiguredValue: true, Nested: map[string]couponStateUpgradeAttrMeta{"products": couponStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}}, "created": couponStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "currency": couponStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "currency_options": couponStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "amount_off": types.Int64Type}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]couponStateUpgradeAttrMeta{"key": couponStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "amount_off": couponStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "required", LegacyBehavior: "required"}}}, "duration": couponStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "duration_in_months": couponStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "id": couponStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "livemode": couponStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "max_redemptions": couponStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "metadata": couponStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "name": couponStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "percent_off": couponStateUpgradeAttrMeta{AttrType: types.Float64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "redeem_by": couponStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "times_redeemed": couponStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "valid": couponStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}}

var couponStateUpgradeSingletonPaths = map[string]struct{}{"applies_to": struct{}{}}

func couponAttrMapFromModel(model interface{}) map[string]attr.Value {
	value := reflect.ValueOf(model)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if !value.IsValid() || value.Kind() != reflect.Struct {
		return map[string]attr.Value{}
	}

	result := make(map[string]attr.Value, value.NumField())
	typeInfo := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := typeInfo.Field(i)
		tag := field.Tag.Get("tfsdk")
		if tag == "" || tag == "-" {
			continue
		}

		attrValue, ok := value.Field(i).Interface().(attr.Value)
		if !ok {
			continue
		}
		result[tag] = attrValue
	}
	return result
}

func couponSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
	value := reflect.ValueOf(target)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return
	}

	elem := value.Elem()
	if !elem.IsValid() || elem.Kind() != reflect.Struct {
		return
	}

	typeInfo := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		field := typeInfo.Field(i)
		tag := field.Tag.Get("tfsdk")
		if tag == "" || tag == "-" {
			continue
		}

		attrValue, ok := values[tag]
		if !ok {
			continue
		}

		fieldValue := elem.Field(i)
		renderedValue := reflect.ValueOf(attrValue)
		if renderedValue.Type().AssignableTo(fieldValue.Type()) {
			fieldValue.Set(renderedValue)
			continue
		}
		if renderedValue.Type().ConvertibleTo(fieldValue.Type()) {
			fieldValue.Set(renderedValue.Convert(fieldValue.Type()))
		}
	}
}

func couponIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func couponShouldPreserveChild(parent couponStateUpgradeAttrMeta, child couponStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !couponIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !couponIsComputedBehavior(child.Behavior)
}

func couponNullValueForType(attributeType attr.Type) attr.Value {
	switch t := attributeType.(type) {
	case basetypes.StringType:
		return types.StringNull()
	case basetypes.Int64Type:
		return types.Int64Null()
	case basetypes.Float64Type:
		return types.Float64Null()
	case basetypes.BoolType:
		return types.BoolNull()
	case basetypes.MapType:
		return types.MapNull(t.ElemType)
	case basetypes.ListType:
		return types.ListNull(t.ElemType)
	case basetypes.SetType:
		return types.SetNull(t.ElemType)
	case basetypes.ObjectType:
		return types.ObjectNull(t.AttrTypes)
	default:
		return types.StringNull()
	}
}

func couponLegacyUpgradeIsEmptyValue(value attr.Value) bool {
	if value == nil || value.IsNull() || value.IsUnknown() {
		return true
	}
	switch typed := value.(type) {
	case types.String:
		return typed.ValueString() == ""
	case types.Int64:
		return typed.ValueInt64() == 0
	case types.Float64:
		return typed.ValueFloat64() == 0
	case types.Bool:
		return !typed.ValueBool()
	case types.Map:
		return len(typed.Elements()) == 0
	case types.List:
		return len(typed.Elements()) == 0
	case types.Set:
		return len(typed.Elements()) == 0
	case types.Object:
		return len(typed.Attributes()) == 0
	default:
		return false
	}
}

func couponLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
	switch typed := value.(type) {
	case types.Int64:
		if typed.IsNull() || typed.IsUnknown() {
			return 0, false
		}
		return typed.ValueInt64(), true
	case types.Float64:
		if typed.IsNull() || typed.IsUnknown() {
			return 0, false
		}
		return int64(typed.ValueFloat64()), true
	default:
		return 0, false
	}
}

func couponLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := couponLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func couponLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.Int64)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || typedValue.ValueInt64() != 0 {
		return false
	}
	decimalValue, ok := siblings[name+"_decimal"]
	if !ok {
		return false
	}
	typedDecimal, ok := decimalValue.(types.String)
	if !ok || typedDecimal.IsNull() || typedDecimal.IsUnknown() {
		return false
	}
	return typedDecimal.ValueString() == "0"
}

func couponLegacyUpgradeNormalizeChild(parent couponStateUpgradeAttrMeta, name string, child couponStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if couponLegacyUpgradeIsEmptyValue(value) {
		return couponNullValueForType(child.AttrType)
	}
	if couponLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return couponNullValueForType(child.AttrType)
	}
	if couponLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return couponNullValueForType(child.AttrType)
	}
	return value
}

func couponLegacyUpgradeChildAttr(path []string, parent couponStateUpgradeAttrMeta, name string, child couponStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !couponShouldPreserveChild(parent, child) {
		return couponNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return couponNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := couponUpgradeValue(nextPath, child, childValue)
	return couponLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func couponUpgradeAttrs(path []string, meta map[string]couponStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = couponNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := couponUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = couponLegacyUpgradeNormalizeChild(
				couponStateUpgradeAttrMeta{PreserveConfiguredValue: true},
				name,
				fieldMeta,
				upgradedValue,
				prior,
			)
		}
		upgraded[name] = upgradedValue
	}
	return upgraded
}

func couponUpgradeObjectValue(path []string, meta couponStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = couponLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func couponUpgradeSingletonListToObject(path []string, meta couponStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return couponNullValueForType(meta.AttrType)
	}
	if listValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if listValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}

	elements := listValue.Elements()
	if len(elements) == 0 {
		return types.ObjectNull(objectType.AttrTypes)
	}

	firstObject, ok := elements[0].(types.Object)
	if !ok {
		if baseObject, baseOk := elements[0].(basetypes.ObjectValue); baseOk {
			firstObject = types.Object(baseObject)
		} else {
			return types.ObjectUnknown(objectType.AttrTypes)
		}
	}

	sourceAttrs := firstObject.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = couponLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func couponUpgradeListValue(path []string, meta couponStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return couponNullValueForType(meta.AttrType)
	}
	if listValue.IsNull() {
		return types.ListNull(listType.ElemType)
	}
	if listValue.IsUnknown() {
		return types.ListUnknown(listType.ElemType)
	}
	if len(meta.Nested) == 0 {
		return listValue
	}

	elementObjectType, ok := listType.ElemType.(basetypes.ObjectType)
	if !ok {
		return listValue
	}

	elements := listValue.Elements()
	upgradedElements := make([]attr.Value, 0, len(elements))
	for _, element := range elements {
		objectValue, ok := element.(types.Object)
		if !ok {
			if baseObject, baseOk := element.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				upgradedElements = append(upgradedElements, couponNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			couponUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func couponUpgradeValue(path []string, meta couponStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := couponStateUpgradeSingletonPaths[pathKey]; ok {
			return couponUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return couponNullValueForType(meta.AttrType)
			}
		}
		return couponUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		return couponUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradeCouponStateV0(ctx context.Context, prior CouponResourceV0Model) (CouponResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := couponUpgradeAttrs(nil, couponStateUpgradeRootMeta, couponAttrMapFromModel(prior))
	var upgraded CouponResourceModel
	couponSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *CouponResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     1,
		Description: "A coupon contains information about a percent-off or amount-off discount you\nmight want to apply to a customer. Coupons may be applied to [subscriptions](https://api.stripe.com#subscriptions), [invoices](https://api.stripe.com#invoices),\n[checkout sessions](https://docs.stripe.com/api/checkout/sessions), [quotes](https://api.stripe.com#quotes), and more. Coupons do not work with conventional one-off [charges](/api/charges/create) or [payment intents](https://docs.stripe.com/api/payment_intents).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("coupon")},
			},
			"amount_off": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "If `amount_off` has been set, the three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the amount to take off.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"duration": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "One of `forever`, `once`, or `repeating`. Describes how long a customer who applies this coupon will get the discount.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("forever", "once", "repeating")},
			},
			"duration_in_months": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "If `duration` is `repeating`, the number of months the coupon applies. Null if coupon `duration` is `forever` or `once`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"max_redemptions": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Maximum number of times this coupon can be redeemed, in total, across all customers, before it is no longer valid.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Name of the coupon displayed to customers on for instance invoices or receipts.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"percent_off": schema.Float64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Percent that will be taken off the subtotal of any invoices for this customer for the duration of the coupon. For example, a coupon with percent_off of 50 will make a $ (or local equivalent)100 invoice $ (or local equivalent)50 instead.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown(), float64planmodifier.RequiresReplace()},
			},
			"redeem_by": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Date after which the coupon can no longer be redeemed.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"times_redeemed": schema.Int64Attribute{
				Computed:      true,
				Description:   "Number of times this coupon has been applied to a customer.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"valid": schema.BoolAttribute{
				Computed:      true,
				Description:   "Taking account of the above properties, whether this coupon can still be applied to a customer.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
		},
		Blocks: map[string]schema.Block{
			"applies_to": schema.SingleNestedBlock{
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"products": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A list of product IDs this coupon applies to",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
						ElementType:   types.StringType,
					},
				},
			},
			"currency_options": schema.ListNestedBlock{
				Description: "Coupons defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required:    true,
							Description: "Key for this entry.",
						},
						"amount_off": schema.Int64Attribute{
							Required:    true,
							Description: "Amount (in the `currency` specified) that will be taken off the subtotal of any invoices for this customer.",
						},
					},
				},
			},
		},
	}
}

func (r *CouponResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CouponResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandCouponCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Coupon create params", err.Error())
		return
	}

	obj, err := r.client.V1Coupons.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Coupon", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Coupons.B, r.client.V1Coupons.Key, stripe.FormatURLPath("/v1/coupons/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Coupon create raw response", err.Error())
		return
	}

	if err := flattenCoupon(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Coupon create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CouponResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState CouponResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CouponResourceModel
	state = priorState

	obj, err := r.client.V1Coupons.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Coupon", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Coupons.B, r.client.V1Coupons.Key, stripe.FormatURLPath("/v1/coupons/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Coupon raw response", err.Error())
		return
	}

	if err := flattenCoupon(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Coupon read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *CouponResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CouponResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CouponResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandCouponUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Coupon update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Coupon update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Coupons.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Coupon", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Coupons.B, r.client.V1Coupons.Key, stripe.FormatURLPath("/v1/coupons/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Coupon update raw response", err.Error())
		return
	}

	if err := flattenCoupon(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Coupon update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CouponResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CouponResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1Coupons.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting Coupon", err.Error())
		return
	}
}

func (r *CouponResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandCouponCreate(plan CouponResourceModel) (*stripe.CouponCreateParams, error) {
	params := &stripe.CouponCreateParams{}

	if !plan.AmountOff.IsNull() && !plan.AmountOff.IsUnknown() {
		params.AmountOff = stripe.Int64(plan.AmountOff.ValueInt64())
	}
	if !plan.AppliesTo.IsNull() && !plan.AppliesTo.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AppliesTo", plan.AppliesTo) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "applies_to", params)
		}
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.CurrencyOptions.IsNull() && !plan.CurrencyOptions.IsUnknown() {
		if !assignKeyedListValueToNamedField(params, "CurrencyOptions", plan.CurrencyOptions, "key") {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency_options", params)
		}
	}
	if !plan.Duration.IsNull() && !plan.Duration.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Duration", "Duration", plan.Duration.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "duration", params)
		}
	}
	if !plan.DurationInMonths.IsNull() && !plan.DurationInMonths.IsUnknown() {
		params.DurationInMonths = stripe.Int64(plan.DurationInMonths.ValueInt64())
	}
	if !plan.MaxRedemptions.IsNull() && !plan.MaxRedemptions.IsUnknown() {
		params.MaxRedemptions = stripe.Int64(plan.MaxRedemptions.ValueInt64())
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Name", "Name", plan.Name.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "name", params)
		}
	}
	if !plan.PercentOff.IsNull() && !plan.PercentOff.IsUnknown() {
		params.PercentOff = stripe.Float64(plan.PercentOff.ValueFloat64())
	}
	if !plan.RedeemBy.IsNull() && !plan.RedeemBy.IsUnknown() {
		params.RedeemBy = stripe.Int64(plan.RedeemBy.ValueInt64())
	}

	return params, nil
}

func expandCouponUpdate(plan CouponResourceModel, state CouponResourceModel) (*stripe.CouponUpdateParams, error) {
	params := &stripe.CouponUpdateParams{}

	if !plan.CurrencyOptions.Equal(state.CurrencyOptions) && !plan.CurrencyOptions.IsNull() && !plan.CurrencyOptions.IsUnknown() {
		if !assignKeyedListValueToNamedField(params, "CurrencyOptions", plan.CurrencyOptions, "key") {
			if !plan.CurrencyOptions.Equal(state.CurrencyOptions) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "currency_options", params)
			}
		}
	}
	if !plan.Metadata.Equal(state.Metadata) && !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			if !plan.Metadata.Equal(state.Metadata) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "metadata", params)
			}
		}
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Name", "Name", plan.Name.ValueString()) {
			if !plan.Name.Equal(state.Name) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "name", params)
			}
		}
	}

	return params, nil
}

func flattenCoupon(obj *stripe.Coupon, state *CouponResourceModel) error {
	raw, hasRaw := rawResponseToPlain(obj)
	{
		if rawValueObject, rawOk := plainValueAtPath(raw, "object"); rawOk {
			if valueObject, err := flattenPlainValue(rawValueObject, types.StringType, "object", "raw response"); err != nil {
				return err
			} else {
				if typedObject, ok := valueObject.(types.String); ok {
					state.Object = typedObject
				}
			}
		} else if !hasRaw {
			if responseValueObject, ok := plainFromResponseField(obj, "Object"); ok {
				if valueObject, err := flattenPlainValue(responseValueObject, types.StringType, "object", "response struct"); err != nil {
					return err
				} else {
					if typedObject, ok := valueObject.(types.String); ok {
						state.Object = typedObject
					}
				}
			}
		}
	}
	{
		if rawValueAmountOff, rawOk := plainValueAtPath(raw, "amount_off"); rawOk {
			if valueAmountOff, err := flattenPlainValue(rawValueAmountOff, types.Int64Type, "amount_off", "raw response"); err != nil {
				return err
			} else {
				if typedAmountOff, ok := valueAmountOff.(types.Int64); ok {
					state.AmountOff = typedAmountOff
				}
			}
		} else if !hasRaw {
			if responseValueAmountOff, ok := plainFromResponseField(obj, "AmountOff"); ok {
				if valueAmountOff, err := flattenPlainValue(responseValueAmountOff, types.Int64Type, "amount_off", "response struct"); err != nil {
					return err
				} else {
					if typedAmountOff, ok := valueAmountOff.(types.Int64); ok {
						state.AmountOff = typedAmountOff
					}
				}
			}
		}
	}
	{
		assignedAppliesTo := false
		if rawValueAppliesTo, rawOk := plainValueAtPath(raw, "applies_to"); rawOk {
			if rawValueAppliesTo != nil {
				sourceAppliesTo := mergeMissingPlainLeaves(applyConfiguredKeyedListShapes(rawValueAppliesTo, attrValueToPlain(state.AppliesTo)), attrValueToPlain(state.AppliesTo))
				if !state.AppliesTo.IsNull() && !state.AppliesTo.IsUnknown() {
					if valueAppliesTo, err := flattenPlainValue(sourceAppliesTo, types.ObjectType{AttrTypes: map[string]attr.Type{"products": types.ListType{ElemType: types.StringType}}}, "applies_to", "raw response"); err != nil {
						return err
					} else {
						if typedAppliesTo, ok := valueAppliesTo.(types.Object); ok {
							state.AppliesTo = typedAppliesTo
							assignedAppliesTo = true
						}
					}
				}
			}
		}
		if !assignedAppliesTo {
			if !hasRaw {
				if responseValueAppliesTo, ok := plainFromResponseField(obj, "AppliesTo"); ok {
					sourceAppliesTo := mergeMissingPlainLeaves(applyConfiguredKeyedListShapes(responseValueAppliesTo, attrValueToPlain(state.AppliesTo)), attrValueToPlain(state.AppliesTo))
					if !state.AppliesTo.IsNull() && !state.AppliesTo.IsUnknown() {
						if valueAppliesTo, err := flattenPlainValue(
							sourceAppliesTo,
							types.ObjectType{AttrTypes: map[string]attr.Type{"products": types.ListType{ElemType: types.StringType}}},
							"applies_to",
							"response struct",
						); err != nil {
							return err
						} else {
							if typedAppliesTo, ok := valueAppliesTo.(types.Object); ok {
								state.AppliesTo = typedAppliesTo
								assignedAppliesTo = true
							}
						}
					}
				}
			}
		}
	}
	{
		if rawValueCreated, rawOk := plainValueAtPath(raw, "created"); rawOk {
			if valueCreated, err := flattenPlainValue(rawValueCreated, types.Int64Type, "created", "raw response"); err != nil {
				return err
			} else {
				if typedCreated, ok := valueCreated.(types.Int64); ok {
					state.Created = typedCreated
				}
			}
		} else if !hasRaw {
			if responseValueCreated, ok := plainFromResponseField(obj, "Created"); ok {
				if valueCreated, err := flattenPlainValue(responseValueCreated, types.Int64Type, "created", "response struct"); err != nil {
					return err
				} else {
					if typedCreated, ok := valueCreated.(types.Int64); ok {
						state.Created = typedCreated
					}
				}
			}
		}
	}
	{
		if rawValueCurrency, rawOk := plainValueAtPath(raw, "currency"); rawOk {
			if valueCurrency, err := flattenPlainValue(rawValueCurrency, types.StringType, "currency", "raw response"); err != nil {
				return err
			} else {
				if typedCurrency, ok := valueCurrency.(types.String); ok {
					state.Currency = typedCurrency
				}
			}
		} else if !hasRaw {
			if responseValueCurrency, ok := plainFromResponseField(obj, "Currency"); ok {
				if valueCurrency, err := flattenPlainValue(responseValueCurrency, types.StringType, "currency", "response struct"); err != nil {
					return err
				} else {
					if typedCurrency, ok := valueCurrency.(types.String); ok {
						state.Currency = typedCurrency
					}
				}
			}
		}
	}
	{
		if rawValueCurrencyOptions, rawOk := plainValueAtPath(raw, "currency_options"); rawOk {
			if !plainValueIsEmpty(applyConfiguredKeyedListShapes(plainMapToKeyedList(rawValueCurrencyOptions, "key"), attrValueToPlain(state.CurrencyOptions))) || state.CurrencyOptions.IsUnknown() || !state.CurrencyOptions.IsNull() {
				if valueCurrencyOptions, err := flattenPlainValue(applyConfiguredKeyedListShapes(plainMapToKeyedList(rawValueCurrencyOptions, "key"), attrValueToPlain(state.CurrencyOptions)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "amount_off": types.Int64Type}}}, "currency_options", "raw response"); err != nil {
					return err
				} else {
					if typedCurrencyOptions, ok := valueCurrencyOptions.(types.List); ok {
						state.CurrencyOptions = typedCurrencyOptions
					}
				}
			}
		} else if !hasRaw {
			if responseValueCurrencyOptions, ok := plainFromResponseField(obj, "CurrencyOptions"); ok {
				if !plainValueIsEmpty(applyConfiguredKeyedListShapes(plainMapToKeyedList(responseValueCurrencyOptions, "key"), attrValueToPlain(state.CurrencyOptions))) || state.CurrencyOptions.IsUnknown() || !state.CurrencyOptions.IsNull() {
					if valueCurrencyOptions, err := flattenPlainValue(
						applyConfiguredKeyedListShapes(plainMapToKeyedList(responseValueCurrencyOptions, "key"), attrValueToPlain(state.CurrencyOptions)),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "amount_off": types.Int64Type}}},
						"currency_options",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCurrencyOptions, ok := valueCurrencyOptions.(types.List); ok {
							state.CurrencyOptions = typedCurrencyOptions
						}
					}
				}
			}
		}
	}
	{
		if rawValueDuration, rawOk := plainValueAtPath(raw, "duration"); rawOk {
			if valueDuration, err := flattenPlainValue(rawValueDuration, types.StringType, "duration", "raw response"); err != nil {
				return err
			} else {
				if typedDuration, ok := valueDuration.(types.String); ok {
					state.Duration = typedDuration
				}
			}
		} else if !hasRaw {
			if responseValueDuration, ok := plainFromResponseField(obj, "Duration"); ok {
				if valueDuration, err := flattenPlainValue(responseValueDuration, types.StringType, "duration", "response struct"); err != nil {
					return err
				} else {
					if typedDuration, ok := valueDuration.(types.String); ok {
						state.Duration = typedDuration
					}
				}
			}
		}
	}
	{
		if rawValueDurationInMonths, rawOk := plainValueAtPath(raw, "duration_in_months"); rawOk {
			if valueDurationInMonths, err := flattenPlainValue(rawValueDurationInMonths, types.Int64Type, "duration_in_months", "raw response"); err != nil {
				return err
			} else {
				if typedDurationInMonths, ok := valueDurationInMonths.(types.Int64); ok {
					state.DurationInMonths = typedDurationInMonths
				}
			}
		} else if !hasRaw {
			if responseValueDurationInMonths, ok := plainFromResponseField(obj, "DurationInMonths"); ok {
				if valueDurationInMonths, err := flattenPlainValue(responseValueDurationInMonths, types.Int64Type, "duration_in_months", "response struct"); err != nil {
					return err
				} else {
					if typedDurationInMonths, ok := valueDurationInMonths.(types.Int64); ok {
						state.DurationInMonths = typedDurationInMonths
					}
				}
			}
		}
	}
	{
		if rawValueID, rawOk := plainValueAtPath(raw, "id"); rawOk {
			if valueID, err := flattenPlainValue(rawValueID, types.StringType, "id", "raw response"); err != nil {
				return err
			} else {
				if typedID, ok := valueID.(types.String); ok {
					state.ID = typedID
				}
			}
		} else if !hasRaw {
			if responseValueID, ok := plainFromResponseField(obj, "ID"); ok {
				if valueID, err := flattenPlainValue(responseValueID, types.StringType, "id", "response struct"); err != nil {
					return err
				} else {
					if typedID, ok := valueID.(types.String); ok {
						state.ID = typedID
					}
				}
			}
		}
	}
	{
		if rawValueLivemode, rawOk := plainValueAtPath(raw, "livemode"); rawOk {
			if valueLivemode, err := flattenPlainValue(rawValueLivemode, types.BoolType, "livemode", "raw response"); err != nil {
				return err
			} else {
				if typedLivemode, ok := valueLivemode.(types.Bool); ok {
					state.Livemode = typedLivemode
				}
			}
		} else if !hasRaw {
			if responseValueLivemode, ok := plainFromResponseField(obj, "Livemode"); ok {
				if valueLivemode, err := flattenPlainValue(responseValueLivemode, types.BoolType, "livemode", "response struct"); err != nil {
					return err
				} else {
					if typedLivemode, ok := valueLivemode.(types.Bool); ok {
						state.Livemode = typedLivemode
					}
				}
			}
		}
	}
	{
		if rawValueMaxRedemptions, rawOk := plainValueAtPath(raw, "max_redemptions"); rawOk {
			if valueMaxRedemptions, err := flattenPlainValue(rawValueMaxRedemptions, types.Int64Type, "max_redemptions", "raw response"); err != nil {
				return err
			} else {
				if typedMaxRedemptions, ok := valueMaxRedemptions.(types.Int64); ok {
					state.MaxRedemptions = typedMaxRedemptions
				}
			}
		} else if !hasRaw {
			if responseValueMaxRedemptions, ok := plainFromResponseField(obj, "MaxRedemptions"); ok {
				if valueMaxRedemptions, err := flattenPlainValue(responseValueMaxRedemptions, types.Int64Type, "max_redemptions", "response struct"); err != nil {
					return err
				} else {
					if typedMaxRedemptions, ok := valueMaxRedemptions.(types.Int64); ok {
						state.MaxRedemptions = typedMaxRedemptions
					}
				}
			}
		}
	}
	{
		if rawValueMetadata, rawOk := plainValueAtPath(raw, "metadata"); rawOk {
			if valueMetadata, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueMetadata, attrValueToPlain(state.Metadata)), types.MapType{ElemType: types.StringType}, "metadata", "raw response"); err != nil {
				return err
			} else {
				if typedMetadata, ok := valueMetadata.(types.Map); ok {
					state.Metadata = typedMetadata
				}
			}
		} else if !hasRaw {
			if responseValueMetadata, ok := plainFromResponseField(obj, "Metadata"); ok {
				if valueMetadata, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueMetadata, attrValueToPlain(state.Metadata)),
					types.MapType{ElemType: types.StringType},
					"metadata",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedMetadata, ok := valueMetadata.(types.Map); ok {
						state.Metadata = typedMetadata
					}
				}
			}
		}
	}
	{
		if rawValueName, rawOk := plainValueAtPath(raw, "name"); rawOk {
			if valueName, err := flattenPlainValue(rawValueName, types.StringType, "name", "raw response"); err != nil {
				return err
			} else {
				if typedName, ok := valueName.(types.String); ok {
					state.Name = typedName
				}
			}
		} else if !hasRaw {
			if responseValueName, ok := plainFromResponseField(obj, "Name"); ok {
				if valueName, err := flattenPlainValue(responseValueName, types.StringType, "name", "response struct"); err != nil {
					return err
				} else {
					if typedName, ok := valueName.(types.String); ok {
						state.Name = typedName
					}
				}
			}
		}
	}
	{
		if rawValuePercentOff, rawOk := plainValueAtPath(raw, "percent_off"); rawOk {
			if valuePercentOff, err := flattenPlainValue(rawValuePercentOff, types.Float64Type, "percent_off", "raw response"); err != nil {
				return err
			} else {
				if typedPercentOff, ok := valuePercentOff.(types.Float64); ok {
					state.PercentOff = typedPercentOff
				}
			}
		} else if !hasRaw {
			if responseValuePercentOff, ok := plainFromResponseField(obj, "PercentOff"); ok {
				if valuePercentOff, err := flattenPlainValue(responseValuePercentOff, types.Float64Type, "percent_off", "response struct"); err != nil {
					return err
				} else {
					if typedPercentOff, ok := valuePercentOff.(types.Float64); ok {
						state.PercentOff = typedPercentOff
					}
				}
			}
		}
	}
	{
		if rawValueRedeemBy, rawOk := plainValueAtPath(raw, "redeem_by"); rawOk {
			if valueRedeemBy, err := flattenPlainValue(rawValueRedeemBy, types.Int64Type, "redeem_by", "raw response"); err != nil {
				return err
			} else {
				if typedRedeemBy, ok := valueRedeemBy.(types.Int64); ok {
					state.RedeemBy = typedRedeemBy
				}
			}
		} else if !hasRaw {
			if responseValueRedeemBy, ok := plainFromResponseField(obj, "RedeemBy"); ok {
				if valueRedeemBy, err := flattenPlainValue(responseValueRedeemBy, types.Int64Type, "redeem_by", "response struct"); err != nil {
					return err
				} else {
					if typedRedeemBy, ok := valueRedeemBy.(types.Int64); ok {
						state.RedeemBy = typedRedeemBy
					}
				}
			}
		}
	}
	{
		if rawValueTimesRedeemed, rawOk := plainValueAtPath(raw, "times_redeemed"); rawOk {
			if valueTimesRedeemed, err := flattenPlainValue(rawValueTimesRedeemed, types.Int64Type, "times_redeemed", "raw response"); err != nil {
				return err
			} else {
				if typedTimesRedeemed, ok := valueTimesRedeemed.(types.Int64); ok {
					state.TimesRedeemed = typedTimesRedeemed
				}
			}
		} else if !hasRaw {
			if responseValueTimesRedeemed, ok := plainFromResponseField(obj, "TimesRedeemed"); ok {
				if valueTimesRedeemed, err := flattenPlainValue(responseValueTimesRedeemed, types.Int64Type, "times_redeemed", "response struct"); err != nil {
					return err
				} else {
					if typedTimesRedeemed, ok := valueTimesRedeemed.(types.Int64); ok {
						state.TimesRedeemed = typedTimesRedeemed
					}
				}
			}
		}
	}
	{
		if rawValueValid, rawOk := plainValueAtPath(raw, "valid"); rawOk {
			if valueValid, err := flattenPlainValue(rawValueValid, types.BoolType, "valid", "raw response"); err != nil {
				return err
			} else {
				if typedValid, ok := valueValid.(types.Bool); ok {
					state.Valid = typedValid
				}
			}
		} else if !hasRaw {
			if responseValueValid, ok := plainFromResponseField(obj, "Valid"); ok {
				if valueValid, err := flattenPlainValue(responseValueValid, types.BoolType, "valid", "response struct"); err != nil {
					return err
				} else {
					if typedValid, ok := valueValid.(types.Bool); ok {
						state.Valid = typedValid
					}
				}
			}
		}
	}
	return nil
}
