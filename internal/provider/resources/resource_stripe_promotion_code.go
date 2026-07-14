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

var _ resource.Resource = &PromotionCodeResource{}

var _ resource.ResourceWithConfigure = &PromotionCodeResource{}

var _ resource.ResourceWithImportState = &PromotionCodeResource{}

func NewPromotionCodeResource() resource.Resource {
	return &PromotionCodeResource{}
}

type PromotionCodeResource struct {
	client *stripe.Client
}

type PromotionCodeResourceModel struct {
	Object          types.String `tfsdk:"object"`
	Active          types.Bool   `tfsdk:"active"`
	Code            types.String `tfsdk:"code"`
	Created         types.Int64  `tfsdk:"created"`
	Customer        types.String `tfsdk:"customer"`
	CustomerAccount types.String `tfsdk:"customer_account"`
	ExpiresAt       types.Int64  `tfsdk:"expires_at"`
	ID              types.String `tfsdk:"id"`
	Livemode        types.Bool   `tfsdk:"livemode"`
	MaxRedemptions  types.Int64  `tfsdk:"max_redemptions"`
	Metadata        types.Map    `tfsdk:"metadata"`
	Promotion       types.List   `tfsdk:"promotion"`
	Restrictions    types.List   `tfsdk:"restrictions"`
	TimesRedeemed   types.Int64  `tfsdk:"times_redeemed"`
}

func (r *PromotionCodeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PromotionCodeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_promotion_code"
}

var _ resource.ResourceWithUpgradeState = &PromotionCodeResource{}

func (r *PromotionCodeResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		1: {
			PriorSchema: promotionCodeResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior PromotionCodeResourceV0Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradePromotionCodeStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func promotionCodeResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "A Promotion Code represents a customer-redeemable code for an underlying promotion.\nYou can create multiple codes for a single promotion.\n\nIf you enable promotion codes in your [customer portal configuration](https://docs.stripe.com/customer-management/configure-portal), then customers can redeem a code themselves when updating a subscription in the portal.\nCustomers can also view the currently active promotion codes and coupons on each of their subscriptions in the portal.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("promotion_code")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the promotion code is currently active. A promotion code is only active if the coupon is also valid.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"code": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for each customer. Valid characters are lower case letters (a-z), upper case letters (A-Z), digits (0-9), and dashes (-).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer who can use this promotion code.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"customer_account": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account representing the customer who can use this promotion code.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"expires_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Date at which the promotion code can no longer be redeemed.",
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
				Description:   "Maximum number of times this promotion code can be redeemed.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"promotion": schema.SingleNestedAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"coupon": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "If promotion `type` is `coupon`, the coupon for this promotion.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"type": schema.StringAttribute{
						Required:      true,
						Description:   "The type of promotion.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("coupon")},
					},
				},
			},
			"restrictions": schema.SingleNestedAttribute{
				Optional: true,

				Attributes: map[string]schema.Attribute{
					"currency_options": schema.ListNestedAttribute{
						Optional:    true,
						Description: "Promotion code restrictions defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Required:    true,
									Description: "Key for this entry.",
								},
								"minimum_amount": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
							},
						},
					},
					"first_time_transaction": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A Boolean indicating if the Promotion Code should only be redeemed for Customers without any successful payments or invoices",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown(), boolplanmodifier.RequiresReplace()},
					},
					"minimum_amount": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"minimum_amount_currency": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Three-letter [ISO code](https://stripe.com/docs/currencies) for minimum_amount",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"times_redeemed": schema.Int64Attribute{
				Computed:      true,
				Description:   "Number of times this promotion code has been used.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
		},
	}
}

type PromotionCodeResourceV0Model struct {
	Object          types.String `tfsdk:"object"`
	Active          types.Bool   `tfsdk:"active"`
	Code            types.String `tfsdk:"code"`
	Created         types.Int64  `tfsdk:"created"`
	Customer        types.String `tfsdk:"customer"`
	CustomerAccount types.String `tfsdk:"customer_account"`
	ExpiresAt       types.Int64  `tfsdk:"expires_at"`
	ID              types.String `tfsdk:"id"`
	Livemode        types.Bool   `tfsdk:"livemode"`
	MaxRedemptions  types.Int64  `tfsdk:"max_redemptions"`
	Metadata        types.Map    `tfsdk:"metadata"`
	Promotion       types.Object `tfsdk:"promotion"`
	Restrictions    types.Object `tfsdk:"restrictions"`
	TimesRedeemed   types.Int64  `tfsdk:"times_redeemed"`
}

type promotioncodeStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]promotioncodeStateUpgradeAttrMeta
}

var promotioncodeStateUpgradeRootMeta = map[string]promotioncodeStateUpgradeAttrMeta{"object": promotioncodeStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "active": promotioncodeStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "code": promotioncodeStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "created": promotioncodeStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "customer": promotioncodeStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "customer_account": promotioncodeStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "expires_at": promotioncodeStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "id": promotioncodeStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "livemode": promotioncodeStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "max_redemptions": promotioncodeStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "metadata": promotioncodeStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "promotion": promotioncodeStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "type": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]promotioncodeStateUpgradeAttrMeta{"coupon": promotioncodeStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "type": promotioncodeStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "restrictions": promotioncodeStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "minimum_amount": types.Int64Type}}}, "first_time_transaction": types.BoolType, "minimum_amount": types.Int64Type, "minimum_amount_currency": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]promotioncodeStateUpgradeAttrMeta{"currency_options": promotioncodeStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "minimum_amount": types.Int64Type}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]promotioncodeStateUpgradeAttrMeta{"key": promotioncodeStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "minimum_amount": promotioncodeStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}}, "first_time_transaction": promotioncodeStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "minimum_amount": promotioncodeStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "minimum_amount_currency": promotioncodeStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}}, "times_redeemed": promotioncodeStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}}

var promotioncodeStateUpgradeSingletonPaths = map[string]struct{}{}

var promotioncodeStateUpgradeLegacyObjectPaths = map[string]struct{}{"promotion": struct{}{}, "restrictions": struct{}{}}

func promotioncodeAttrMapFromModel(model interface{}) map[string]attr.Value {
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

func promotioncodeSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
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

func promotioncodeIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func promotioncodeShouldPreserveChild(parent promotioncodeStateUpgradeAttrMeta, child promotioncodeStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !promotioncodeIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !promotioncodeIsComputedBehavior(child.Behavior)
}

func promotioncodeNullValueForType(attributeType attr.Type) attr.Value {
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

func promotioncodeLegacyUpgradeIsEmptyValue(value attr.Value) bool {
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

func promotioncodeLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
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

func promotioncodeLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := promotioncodeLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func promotioncodeLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
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

func promotioncodeLegacyUpgradeNormalizeChild(parent promotioncodeStateUpgradeAttrMeta, name string, child promotioncodeStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if promotioncodeLegacyUpgradeIsEmptyValue(value) {
		return promotioncodeNullValueForType(child.AttrType)
	}
	if promotioncodeLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return promotioncodeNullValueForType(child.AttrType)
	}
	if promotioncodeLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return promotioncodeNullValueForType(child.AttrType)
	}
	return value
}

func promotioncodeLegacyUpgradeChildAttr(path []string, parent promotioncodeStateUpgradeAttrMeta, name string, child promotioncodeStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !promotioncodeShouldPreserveChild(parent, child) {
		return promotioncodeNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return promotioncodeNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := promotioncodeUpgradeValue(nextPath, child, childValue)
	return promotioncodeLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func promotioncodeUpgradeAttrs(path []string, meta map[string]promotioncodeStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = promotioncodeNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := promotioncodeUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = promotioncodeLegacyUpgradeNormalizeChild(
				promotioncodeStateUpgradeAttrMeta{PreserveConfiguredValue: true},
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

func promotioncodeUpgradeObjectValue(path []string, meta promotioncodeStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = promotioncodeLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func promotioncodeUpgradeSingletonListToObject(path []string, meta promotioncodeStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return promotioncodeNullValueForType(meta.AttrType)
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
		upgradedAttrs[name] = promotioncodeLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func promotioncodeUpgradeObjectValueToSingletonList(path []string, meta promotioncodeStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	objectValue, ok := priorValue.(types.Object)
	if !ok {
		if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
			objectValue = types.Object(baseObject)
		} else {
			return promotioncodeNullValueForType(meta.AttrType)
		}
	}
	if objectValue.IsNull() {
		return types.ListNull(listType.ElemType)
	}
	if objectValue.IsUnknown() {
		return types.ListUnknown(listType.ElemType)
	}

	elementObjectType, ok := listType.ElemType.(basetypes.ObjectType)
	if !ok {
		return promotioncodeNullValueForType(meta.AttrType)
	}

	upgradedObject := promotioncodeUpgradeObjectValue(path, meta, elementObjectType, objectValue)
	return types.ListValueMust(listType.ElemType, []attr.Value{upgradedObject})
}

func promotioncodeUpgradeListValue(path []string, meta promotioncodeStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return promotioncodeNullValueForType(meta.AttrType)
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
				upgradedElements = append(upgradedElements, promotioncodeNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			promotioncodeUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func promotioncodeUpgradeValue(path []string, meta promotioncodeStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := promotioncodeStateUpgradeSingletonPaths[pathKey]; ok {
			return promotioncodeUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return promotioncodeNullValueForType(meta.AttrType)
			}
		}
		return promotioncodeUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		if _, ok := promotioncodeStateUpgradeLegacyObjectPaths[pathKey]; ok {
			return promotioncodeUpgradeObjectValueToSingletonList(path, meta, attrType, priorValue)
		}
		return promotioncodeUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradePromotionCodeStateV1(ctx context.Context, prior PromotionCodeResourceV0Model) (PromotionCodeResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := promotioncodeUpgradeAttrs(nil, promotioncodeStateUpgradeRootMeta, promotioncodeAttrMapFromModel(prior))
	var upgraded PromotionCodeResourceModel
	promotioncodeSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *PromotionCodeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     2,
		Description: "A Promotion Code represents a customer-redeemable code for an underlying promotion.\nYou can create multiple codes for a single promotion.\n\nIf you enable promotion codes in your [customer portal configuration](https://docs.stripe.com/customer-management/configure-portal), then customers can redeem a code themselves when updating a subscription in the portal.\nCustomers can also view the currently active promotion codes and coupons on each of their subscriptions in the portal.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("promotion_code")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the promotion code is currently active. A promotion code is only active if the coupon is also valid.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"code": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer-facing code. Regardless of case, this code must be unique across all active promotion codes for each customer. Valid characters are lower case letters (a-z), upper case letters (A-Z), digits (0-9), and dashes (-).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer who can use this promotion code.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"customer_account": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account representing the customer who can use this promotion code.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"expires_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Date at which the promotion code can no longer be redeemed.",
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
				Description:   "Maximum number of times this promotion code can be redeemed.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"times_redeemed": schema.Int64Attribute{
				Computed:      true,
				Description:   "Number of times this promotion code has been used.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
		},
		Blocks: map[string]schema.Block{
			"promotion": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"coupon": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "If promotion `type` is `coupon`, the coupon for this promotion.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						},
						"type": schema.StringAttribute{
							Required:      true,
							Description:   "The type of promotion.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							Validators:    []validator.String{stringvalidator.OneOf("coupon")},
						},
					},
				},
			},
			"restrictions": schema.ListNestedBlock{
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"first_time_transaction": schema.BoolAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "A Boolean indicating if the Promotion Code should only be redeemed for Customers without any successful payments or invoices",
							PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown(), boolplanmodifier.RequiresReplace()},
						},
						"minimum_amount": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"minimum_amount_currency": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Three-letter [ISO code](https://stripe.com/docs/currencies) for minimum_amount",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						},
					},
					Blocks: map[string]schema.Block{
						"currency_options": schema.ListNestedBlock{
							Description: "Promotion code restrictions defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Required:    true,
										Description: "Key for this entry.",
									},
									"minimum_amount": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "Minimum amount required to redeem this Promotion Code into a Coupon (e.g., a purchase must be $100 or more to work).",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *PromotionCodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PromotionCodeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandPromotionCodeCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building PromotionCode create params", err.Error())
		return
	}

	obj, err := r.client.V1PromotionCodes.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating PromotionCode", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PromotionCodes.B, r.client.V1PromotionCodes.Key, stripe.FormatURLPath("/v1/promotion_codes/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PromotionCode create raw response", err.Error())
		return
	}

	if err := flattenPromotionCode(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PromotionCode create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PromotionCodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState PromotionCodeResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PromotionCodeResourceModel
	state = priorState

	obj, err := r.client.V1PromotionCodes.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading PromotionCode", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PromotionCodes.B, r.client.V1PromotionCodes.Key, stripe.FormatURLPath("/v1/promotion_codes/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PromotionCode raw response", err.Error())
		return
	}

	if err := flattenPromotionCode(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening PromotionCode read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *PromotionCodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PromotionCodeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PromotionCodeResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandPromotionCodeUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building PromotionCode update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building PromotionCode update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1PromotionCodes.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating PromotionCode", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PromotionCodes.B, r.client.V1PromotionCodes.Key, stripe.FormatURLPath("/v1/promotion_codes/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PromotionCode update raw response", err.Error())
		return
	}

	if err := flattenPromotionCode(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PromotionCode update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PromotionCodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PromotionCodeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.Active.IsNull() && !state.Active.IsUnknown() && !state.Active.ValueBool() {
		return
	}

	params := &stripe.PromotionCodeUpdateParams{}
	activeField := reflect.ValueOf(params).Elem().FieldByName("Active")
	if activeField.IsValid() && activeField.CanSet() {
		if activeField.Kind() == reflect.Pointer && activeField.Type().Elem().Kind() == reflect.Bool {
			activeField.Set(reflect.ValueOf(stripe.Bool(false)))
		} else if activeField.Kind() == reflect.Bool {
			activeField.SetBool(false)
		}
	}

	_, err := r.client.V1PromotionCodes.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating PromotionCode", err.Error())
		return
	}
}

func (r *PromotionCodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandPromotionCodeCreate(plan PromotionCodeResourceModel) (*stripe.PromotionCodeCreateParams, error) {
	params := &stripe.PromotionCodeCreateParams{}

	if !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.Code.IsNull() && !plan.Code.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Code", "Code", plan.Code.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "code", params)
		}
	}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerID", "Customer", plan.Customer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer", params)
		}
	}
	if !plan.CustomerAccount.IsNull() && !plan.CustomerAccount.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerAccount", "CustomerAccount", plan.CustomerAccount.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer_account", params)
		}
	}
	if !plan.ExpiresAt.IsNull() && !plan.ExpiresAt.IsUnknown() {
		params.ExpiresAt = stripe.Int64(plan.ExpiresAt.ValueInt64())
	}
	if !plan.MaxRedemptions.IsNull() && !plan.MaxRedemptions.IsUnknown() {
		params.MaxRedemptions = stripe.Int64(plan.MaxRedemptions.ValueInt64())
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Promotion.IsNull() && !plan.Promotion.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Promotion", plan.Promotion) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "promotion", params)
		}
	}
	if !plan.Restrictions.IsNull() && !plan.Restrictions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Restrictions", plan.Restrictions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "restrictions", params)
		}
	}

	return params, nil
}

func expandPromotionCodeUpdate(plan PromotionCodeResourceModel, state PromotionCodeResourceModel) (*stripe.PromotionCodeUpdateParams, error) {
	params := &stripe.PromotionCodeUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.Metadata.Equal(state.Metadata) && !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			if !plan.Metadata.Equal(state.Metadata) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "metadata", params)
			}
		}
	}
	if !plan.Restrictions.Equal(state.Restrictions) && !plan.Restrictions.IsNull() && !plan.Restrictions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Restrictions", plan.Restrictions) {
			if !plan.Restrictions.Equal(state.Restrictions) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "restrictions", params)
			}
		}
	}

	return params, nil
}

func flattenPromotionCode(obj *stripe.PromotionCode, state *PromotionCodeResourceModel) error {
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
		if rawValueActive, rawOk := plainValueAtPath(raw, "active"); rawOk {
			if valueActive, err := flattenPlainValue(rawValueActive, types.BoolType, "active", "raw response"); err != nil {
				return err
			} else {
				if typedActive, ok := valueActive.(types.Bool); ok {
					state.Active = typedActive
				}
			}
		} else if !hasRaw {
			if responseValueActive, ok := plainFromResponseField(obj, "Active"); ok {
				if valueActive, err := flattenPlainValue(responseValueActive, types.BoolType, "active", "response struct"); err != nil {
					return err
				} else {
					if typedActive, ok := valueActive.(types.Bool); ok {
						state.Active = typedActive
					}
				}
			}
		}
	}
	{
		if rawValueCode, rawOk := plainValueAtPath(raw, "code"); rawOk {
			if valueCode, err := flattenPlainValue(rawValueCode, types.StringType, "code", "raw response"); err != nil {
				return err
			} else {
				if typedCode, ok := valueCode.(types.String); ok {
					state.Code = typedCode
				}
			}
		} else if !hasRaw {
			if responseValueCode, ok := plainFromResponseField(obj, "Code"); ok {
				if valueCode, err := flattenPlainValue(responseValueCode, types.StringType, "code", "response struct"); err != nil {
					return err
				} else {
					if typedCode, ok := valueCode.(types.String); ok {
						state.Code = typedCode
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
		if state.Customer.IsNull() || state.Customer.IsUnknown() {
			if rawValueCustomer, rawOk := plainValueAtPath(raw, "customer"); rawOk {
				if typedCustomer, ok := plainToStringIDValue(rawValueCustomer); ok {
					state.Customer = typedCustomer
				}
			} else if !hasRaw {
				if responseValueCustomer, ok := plainFromResponseField(obj, "Customer"); ok {
					if typedCustomer, ok := plainToStringIDValue(responseValueCustomer); ok {
						state.Customer = typedCustomer
					}
				}
			}
		}
	}
	{
		if rawValueCustomerAccount, rawOk := plainValueAtPath(raw, "customer_account"); rawOk {
			if valueCustomerAccount, err := flattenPlainValue(rawValueCustomerAccount, types.StringType, "customer_account", "raw response"); err != nil {
				return err
			} else {
				if typedCustomerAccount, ok := valueCustomerAccount.(types.String); ok {
					state.CustomerAccount = typedCustomerAccount
				}
			}
		} else if !hasRaw {
			if responseValueCustomerAccount, ok := plainFromResponseField(obj, "CustomerAccount"); ok {
				if valueCustomerAccount, err := flattenPlainValue(responseValueCustomerAccount, types.StringType, "customer_account", "response struct"); err != nil {
					return err
				} else {
					if typedCustomerAccount, ok := valueCustomerAccount.(types.String); ok {
						state.CustomerAccount = typedCustomerAccount
					}
				}
			}
		}
	}
	{
		if rawValueExpiresAt, rawOk := plainValueAtPath(raw, "expires_at"); rawOk {
			if valueExpiresAt, err := flattenPlainValue(rawValueExpiresAt, types.Int64Type, "expires_at", "raw response"); err != nil {
				return err
			} else {
				if typedExpiresAt, ok := valueExpiresAt.(types.Int64); ok {
					state.ExpiresAt = typedExpiresAt
				}
			}
		} else if !hasRaw {
			if responseValueExpiresAt, ok := plainFromResponseField(obj, "ExpiresAt"); ok {
				if valueExpiresAt, err := flattenPlainValue(responseValueExpiresAt, types.Int64Type, "expires_at", "response struct"); err != nil {
					return err
				} else {
					if typedExpiresAt, ok := valueExpiresAt.(types.Int64); ok {
						state.ExpiresAt = typedExpiresAt
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
		assignedPromotion := false
		hadRawPromotion := false
		if rawValuePromotion, rawOk := plainValueAtPath(raw, "promotion"); rawOk {
			hadRawPromotion = true
			if rawValuePromotion != nil {
				sourcePromotion := applyConfiguredKeyedListShapes(rawValuePromotion, unwrapPlainSingletonList(attrValueToPlain(state.Promotion)))
				if valuePromotion, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourcePromotion, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "type": types.StringType}}}, "promotion", "raw response"); err != nil {
					return err
				} else {
					if typedPromotion, ok := valuePromotion.(types.List); ok {
						state.Promotion = typedPromotion
						assignedPromotion = true
					}
				}
			}
		}
		if !assignedPromotion {
			if !hasRaw {
				if responseValuePromotion, ok := plainFromResponseField(obj, "Promotion"); ok {
					sourcePromotion := applyConfiguredKeyedListShapes(responseValuePromotion, unwrapPlainSingletonList(attrValueToPlain(state.Promotion)))
					if valuePromotion, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourcePromotion, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "type": types.StringType}}},
						"promotion",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPromotion, ok := valuePromotion.(types.List); ok {
							state.Promotion = typedPromotion
							assignedPromotion = true
						}
					}
				}
			}
		}
		if !assignedPromotion && hadRawPromotion {
			if nullPromotion, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "type": types.StringType}}}); ok {
				if typedPromotion, ok := nullPromotion.(types.List); ok {
					state.Promotion = typedPromotion
				}
			}
		}
	}
	{
		assignedRestrictions := false
		hadRawRestrictions := false
		if rawValueRestrictions, rawOk := plainValueAtPath(raw, "restrictions"); rawOk {
			hadRawRestrictions = true
			if rawValueRestrictions != nil {
				sourceRestrictions := applyConfiguredKeyedListShapes(rawValueRestrictions, unwrapPlainSingletonList(attrValueToPlain(state.Restrictions)))
				if !plainValueIsEmpty(sourceRestrictions) || state.Restrictions.IsUnknown() || !state.Restrictions.IsNull() {
					if valueRestrictions, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceRestrictions, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "minimum_amount": types.Int64Type}}}, "first_time_transaction": types.BoolType, "minimum_amount": types.Int64Type, "minimum_amount_currency": types.StringType}}}, "restrictions", "raw response"); err != nil {
						return err
					} else {
						if typedRestrictions, ok := valueRestrictions.(types.List); ok {
							state.Restrictions = typedRestrictions
							assignedRestrictions = true
						}
					}
				}
			}
		}
		if !assignedRestrictions {
			if !hasRaw {
				if responseValueRestrictions, ok := plainFromResponseField(obj, "Restrictions"); ok {
					sourceRestrictions := applyConfiguredKeyedListShapes(responseValueRestrictions, unwrapPlainSingletonList(attrValueToPlain(state.Restrictions)))
					if !plainValueIsEmpty(sourceRestrictions) || state.Restrictions.IsUnknown() || !state.Restrictions.IsNull() {
						if valueRestrictions, err := flattenPlainValue(
							applyPlainSingletonListShapePaths(sourceRestrictions, [][]string{[]string{}}),
							types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "minimum_amount": types.Int64Type}}}, "first_time_transaction": types.BoolType, "minimum_amount": types.Int64Type, "minimum_amount_currency": types.StringType}}},
							"restrictions",
							"response struct",
						); err != nil {
							return err
						} else {
							if typedRestrictions, ok := valueRestrictions.(types.List); ok {
								state.Restrictions = typedRestrictions
								assignedRestrictions = true
							}
						}
					}
				}
			}
		}
		if !assignedRestrictions && hadRawRestrictions {
			if nullRestrictions, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "minimum_amount": types.Int64Type}}}, "first_time_transaction": types.BoolType, "minimum_amount": types.Int64Type, "minimum_amount_currency": types.StringType}}}); ok {
				if typedRestrictions, ok := nullRestrictions.(types.List); ok {
					state.Restrictions = typedRestrictions
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
	return nil
}
