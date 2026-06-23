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

var _ resource.Resource = &TaxRateResource{}

var _ resource.ResourceWithConfigure = &TaxRateResource{}

var _ resource.ResourceWithImportState = &TaxRateResource{}

func NewTaxRateResource() resource.Resource {
	return &TaxRateResource{}
}

type TaxRateResource struct {
	client *stripe.Client
}

type TaxRateResourceModel struct {
	Object              types.String  `tfsdk:"object"`
	Active              types.Bool    `tfsdk:"active"`
	Country             types.String  `tfsdk:"country"`
	Created             types.Int64   `tfsdk:"created"`
	Description         types.String  `tfsdk:"description"`
	DisplayName         types.String  `tfsdk:"display_name"`
	EffectivePercentage types.Float64 `tfsdk:"effective_percentage"`
	FlatAmount          types.Object  `tfsdk:"flat_amount"`
	ID                  types.String  `tfsdk:"id"`
	Inclusive           types.Bool    `tfsdk:"inclusive"`
	Jurisdiction        types.String  `tfsdk:"jurisdiction"`
	JurisdictionLevel   types.String  `tfsdk:"jurisdiction_level"`
	Livemode            types.Bool    `tfsdk:"livemode"`
	Metadata            types.Map     `tfsdk:"metadata"`
	Percentage          types.Float64 `tfsdk:"percentage"`
	RateType            types.String  `tfsdk:"rate_type"`
	State               types.String  `tfsdk:"state"`
	TaxType             types.String  `tfsdk:"tax_type"`
}

func (r *TaxRateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TaxRateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tax_rate"
}

var _ resource.ResourceWithUpgradeState = &TaxRateResource{}

func (r *TaxRateResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: taxRateResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior TaxRateResourceV0Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeTaxRateStateV0(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func taxRateResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "Tax rates can be applied to [invoices](/invoicing/taxes/tax-rates), [subscriptions](/billing/taxes/tax-rates) and [Checkout Sessions](/payments/checkout/use-manual-tax-rates) to collect tax.\n\nRelated guide: [Tax rates](/billing/taxes/tax-rates)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("tax_rate")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Defaults to `true`. When set to `false`, this tax rate cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"country": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "The display name of the tax rates as it will appear to your customer on their receipt email, PDF, and the hosted invoice page.",
			},
			"effective_percentage": schema.Float64Attribute{
				Computed:      true,
				Description:   "Actual/effective tax rate percentage out of 100. For tax calculations with automatic_tax[enabled]=true,\nthis percentage reflects the rate actually used to calculate tax based on the product's taxability\nand whether the user is registered to collect taxes in the corresponding jurisdiction.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
			},
			"flat_amount": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The amount of the tax rate when the `rate_type` is `flat_amount`. Tax rates with `rate_type` `percentage` can vary based on the transaction, resulting in this field being `null`. This field exposes the amount and currency of the flat tax rate.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"amount": schema.Int64Attribute{
						Computed:      true,
						Description:   "Amount of the tax when the `rate_type` is `flat_amount`. This positive integer represents how much to charge in the smallest currency unit (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"currency": schema.StringAttribute{
						Computed:      true,
						Description:   "Three-letter ISO currency code, in lowercase.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"inclusive": schema.BoolAttribute{
				Required:      true,
				Description:   "This specifies if the tax rate is inclusive or exclusive.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"jurisdiction": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"jurisdiction_level": schema.StringAttribute{
				Computed:      true,
				Description:   "The level of the jurisdiction that imposes this tax rate. Will be `null` for manually defined tax rates.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("city", "country", "county", "district", "multiple", "state")},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"percentage": schema.Float64Attribute{
				Required:      true,
				Description:   "Tax rate percentage out of 100. For tax calculations with automatic_tax[enabled]=true, this percentage includes the statutory tax rate of non-taxable jurisdictions.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"rate_type": schema.StringAttribute{
				Computed:      true,
				Description:   "Indicates the type of tax rate applied to the taxable amount. This value can be `null` when no tax applies to the location. This field is only present for TaxRates created by Stripe Tax.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("flat_amount", "percentage")},
			},
			"state": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "[ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2), without country prefix. For example, \"NY\" for New York, United States.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"tax_type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The high-level tax type, such as `vat` or `sales_tax`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("amusement_tax", "communications_tax", "gst", "hst", "igst", "jct", "lease_tax", "pst", "qst", "retail_delivery_fee", "rst", "sales_tax", "service_tax", "vat")},
			},
		},
	}
}

type TaxRateResourceV0Model struct {
	Object              types.String  `tfsdk:"object"`
	Active              types.Bool    `tfsdk:"active"`
	Country             types.String  `tfsdk:"country"`
	Created             types.Int64   `tfsdk:"created"`
	Description         types.String  `tfsdk:"description"`
	DisplayName         types.String  `tfsdk:"display_name"`
	EffectivePercentage types.Float64 `tfsdk:"effective_percentage"`
	FlatAmount          types.Object  `tfsdk:"flat_amount"`
	ID                  types.String  `tfsdk:"id"`
	Inclusive           types.Bool    `tfsdk:"inclusive"`
	Jurisdiction        types.String  `tfsdk:"jurisdiction"`
	JurisdictionLevel   types.String  `tfsdk:"jurisdiction_level"`
	Livemode            types.Bool    `tfsdk:"livemode"`
	Metadata            types.Map     `tfsdk:"metadata"`
	Percentage          types.Float64 `tfsdk:"percentage"`
	RateType            types.String  `tfsdk:"rate_type"`
	State               types.String  `tfsdk:"state"`
	TaxType             types.String  `tfsdk:"tax_type"`
}

type taxrateStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]taxrateStateUpgradeAttrMeta
}

var taxrateStateUpgradeRootMeta = map[string]taxrateStateUpgradeAttrMeta{"object": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "active": taxrateStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "country": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "created": taxrateStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "description": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "display_name": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "effective_percentage": taxrateStateUpgradeAttrMeta{AttrType: types.Float64Type, Behavior: "computed", LegacyBehavior: "computed"}, "flat_amount": taxrateStateUpgradeAttrMeta{AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType}}, Behavior: "computed", LegacyBehavior: "computed", Nested: map[string]taxrateStateUpgradeAttrMeta{"amount": taxrateStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "currency": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}}}, "id": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "inclusive": taxrateStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "required", LegacyBehavior: "required"}, "jurisdiction": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "jurisdiction_level": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "livemode": taxrateStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "metadata": taxrateStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "percentage": taxrateStateUpgradeAttrMeta{AttrType: types.Float64Type, Behavior: "required", LegacyBehavior: "required"}, "rate_type": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "state": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "tax_type": taxrateStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}

var taxrateStateUpgradeSingletonPaths = map[string]struct{}{}

func taxrateAttrMapFromModel(model interface{}) map[string]attr.Value {
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

func taxrateSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
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

func taxrateIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func taxrateShouldPreserveChild(parent taxrateStateUpgradeAttrMeta, child taxrateStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !taxrateIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !taxrateIsComputedBehavior(child.Behavior)
}

func taxrateNullValueForType(attributeType attr.Type) attr.Value {
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

func taxrateLegacyUpgradeIsEmptyValue(value attr.Value) bool {
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

func taxrateLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
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

func taxrateLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := taxrateLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func taxrateLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
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

func taxrateLegacyUpgradeNormalizeChild(parent taxrateStateUpgradeAttrMeta, name string, child taxrateStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if taxrateLegacyUpgradeIsEmptyValue(value) {
		return taxrateNullValueForType(child.AttrType)
	}
	if taxrateLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return taxrateNullValueForType(child.AttrType)
	}
	if taxrateLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return taxrateNullValueForType(child.AttrType)
	}
	return value
}

func taxrateLegacyUpgradeChildAttr(path []string, parent taxrateStateUpgradeAttrMeta, name string, child taxrateStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !taxrateShouldPreserveChild(parent, child) {
		return taxrateNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return taxrateNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := taxrateUpgradeValue(nextPath, child, childValue)
	return taxrateLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func taxrateUpgradeAttrs(path []string, meta map[string]taxrateStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = taxrateNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := taxrateUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = taxrateLegacyUpgradeNormalizeChild(
				taxrateStateUpgradeAttrMeta{PreserveConfiguredValue: true},
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

func taxrateUpgradeObjectValue(path []string, meta taxrateStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = taxrateLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func taxrateUpgradeSingletonListToObject(path []string, meta taxrateStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return taxrateNullValueForType(meta.AttrType)
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
		upgradedAttrs[name] = taxrateLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func taxrateUpgradeListValue(path []string, meta taxrateStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return taxrateNullValueForType(meta.AttrType)
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
				upgradedElements = append(upgradedElements, taxrateNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			taxrateUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func taxrateUpgradeValue(path []string, meta taxrateStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := taxrateStateUpgradeSingletonPaths[pathKey]; ok {
			return taxrateUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return taxrateNullValueForType(meta.AttrType)
			}
		}
		return taxrateUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		return taxrateUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradeTaxRateStateV0(ctx context.Context, prior TaxRateResourceV0Model) (TaxRateResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := taxrateUpgradeAttrs(nil, taxrateStateUpgradeRootMeta, taxrateAttrMapFromModel(prior))
	var upgraded TaxRateResourceModel
	taxrateSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *TaxRateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     1,
		Description: "Tax rates can be applied to [invoices](/invoicing/taxes/tax-rates), [subscriptions](/billing/taxes/tax-rates) and [Checkout Sessions](/payments/checkout/use-manual-tax-rates) to collect tax.\n\nRelated guide: [Tax rates](/billing/taxes/tax-rates)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("tax_rate")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Defaults to `true`. When set to `false`, this tax rate cannot be used with new applications or Checkout Sessions, but will still work for subscriptions and invoices that already have it set.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"country": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the tax rate for your internal use only. It will not be visible to your customers.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "The display name of the tax rates as it will appear to your customer on their receipt email, PDF, and the hosted invoice page.",
			},
			"effective_percentage": schema.Float64Attribute{
				Computed:      true,
				Description:   "Actual/effective tax rate percentage out of 100. For tax calculations with automatic_tax[enabled]=true,\nthis percentage reflects the rate actually used to calculate tax based on the product's taxability\nand whether the user is registered to collect taxes in the corresponding jurisdiction.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
			},
			"flat_amount": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The amount of the tax rate when the `rate_type` is `flat_amount`. Tax rates with `rate_type` `percentage` can vary based on the transaction, resulting in this field being `null`. This field exposes the amount and currency of the flat tax rate.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"amount": schema.Int64Attribute{
						Computed:      true,
						Description:   "Amount of the tax when the `rate_type` is `flat_amount`. This positive integer represents how much to charge in the smallest currency unit (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"currency": schema.StringAttribute{
						Computed:      true,
						Description:   "Three-letter ISO currency code, in lowercase.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"inclusive": schema.BoolAttribute{
				Required:      true,
				Description:   "This specifies if the tax rate is inclusive or exclusive.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"jurisdiction": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The jurisdiction for the tax rate. You can use this label field for tax reporting purposes. It also appears on your customer’s invoice.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"jurisdiction_level": schema.StringAttribute{
				Computed:      true,
				Description:   "The level of the jurisdiction that imposes this tax rate. Will be `null` for manually defined tax rates.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("city", "country", "county", "district", "multiple", "state")},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"percentage": schema.Float64Attribute{
				Required:      true,
				Description:   "Tax rate percentage out of 100. For tax calculations with automatic_tax[enabled]=true, this percentage includes the statutory tax rate of non-taxable jurisdictions.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"rate_type": schema.StringAttribute{
				Computed:      true,
				Description:   "Indicates the type of tax rate applied to the taxable amount. This value can be `null` when no tax applies to the location. This field is only present for TaxRates created by Stripe Tax.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("flat_amount", "percentage")},
			},
			"state": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "[ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2), without country prefix. For example, \"NY\" for New York, United States.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"tax_type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The high-level tax type, such as `vat` or `sales_tax`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("amusement_tax", "communications_tax", "gst", "hst", "igst", "jct", "lease_tax", "pst", "qst", "retail_delivery_fee", "rst", "sales_tax", "service_tax", "vat")},
			},
		},
	}
}

func (r *TaxRateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TaxRateResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTaxRateCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building TaxRate create params", err.Error())
		return
	}

	obj, err := r.client.V1TaxRates.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating TaxRate", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TaxRates.B, r.client.V1TaxRates.Key, stripe.FormatURLPath("/v1/tax_rates/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TaxRate create raw response", err.Error())
		return
	}

	if err := flattenTaxRate(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxRate create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TaxRateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState TaxRateResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TaxRateResourceModel
	state = priorState

	obj, err := r.client.V1TaxRates.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading TaxRate", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TaxRates.B, r.client.V1TaxRates.Key, stripe.FormatURLPath("/v1/tax_rates/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TaxRate raw response", err.Error())
		return
	}

	if err := flattenTaxRate(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxRate read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *TaxRateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TaxRateResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TaxRateResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandTaxRateUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building TaxRate update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building TaxRate update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1TaxRates.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating TaxRate", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TaxRates.B, r.client.V1TaxRates.Key, stripe.FormatURLPath("/v1/tax_rates/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TaxRate update raw response", err.Error())
		return
	}

	if err := flattenTaxRate(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxRate update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TaxRateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TaxRateResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.Active.IsNull() && !state.Active.IsUnknown() && !state.Active.ValueBool() {
		return
	}

	params := &stripe.TaxRateUpdateParams{}
	activeField := reflect.ValueOf(params).Elem().FieldByName("Active")
	if activeField.IsValid() && activeField.CanSet() {
		if activeField.Kind() == reflect.Pointer && activeField.Type().Elem().Kind() == reflect.Bool {
			activeField.Set(reflect.ValueOf(stripe.Bool(false)))
		} else if activeField.Kind() == reflect.Bool {
			activeField.SetBool(false)
		}
	}

	_, err := r.client.V1TaxRates.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating TaxRate", err.Error())
		return
	}
}

func (r *TaxRateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandTaxRateCreate(plan TaxRateResourceModel) (*stripe.TaxRateCreateParams, error) {
	params := &stripe.TaxRateCreateParams{}

	if !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.Country.IsNull() && !plan.Country.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Country", "Country", plan.Country.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "country", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.DisplayName.IsNull() && !plan.DisplayName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayName", "DisplayName", plan.DisplayName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "display_name", params)
		}
	}
	if !plan.Inclusive.IsNull() && !plan.Inclusive.IsUnknown() {
		params.Inclusive = stripe.Bool(plan.Inclusive.ValueBool())
	}
	if !plan.Jurisdiction.IsNull() && !plan.Jurisdiction.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Jurisdiction", "Jurisdiction", plan.Jurisdiction.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "jurisdiction", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Percentage.IsNull() && !plan.Percentage.IsUnknown() {
		params.Percentage = stripe.Float64(plan.Percentage.ValueFloat64())
	}
	if !plan.State.IsNull() && !plan.State.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "State", "State", plan.State.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "state", params)
		}
	}
	if !plan.TaxType.IsNull() && !plan.TaxType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxType", "TaxType", plan.TaxType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_type", params)
		}
	}

	return params, nil
}

func expandTaxRateUpdate(plan TaxRateResourceModel, state TaxRateResourceModel) (*stripe.TaxRateUpdateParams, error) {
	params := &stripe.TaxRateUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.Country.Equal(state.Country) && !plan.Country.IsNull() && !plan.Country.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Country", "Country", plan.Country.ValueString()) {
			if !plan.Country.Equal(state.Country) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "country", params)
			}
		}
	}
	if !plan.Description.Equal(state.Description) && !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			if !plan.Description.Equal(state.Description) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "description", params)
			}
		}
	}
	if !plan.DisplayName.Equal(state.DisplayName) && !plan.DisplayName.IsNull() && !plan.DisplayName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayName", "DisplayName", plan.DisplayName.ValueString()) {
			if !plan.DisplayName.Equal(state.DisplayName) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "display_name", params)
			}
		}
	}
	if !plan.Jurisdiction.Equal(state.Jurisdiction) && !plan.Jurisdiction.IsNull() && !plan.Jurisdiction.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Jurisdiction", "Jurisdiction", plan.Jurisdiction.ValueString()) {
			if !plan.Jurisdiction.Equal(state.Jurisdiction) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "jurisdiction", params)
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
	if !plan.State.Equal(state.State) && !plan.State.IsNull() && !plan.State.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "State", "State", plan.State.ValueString()) {
			if !plan.State.Equal(state.State) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "state", params)
			}
		}
	}
	if !plan.TaxType.Equal(state.TaxType) && !plan.TaxType.IsNull() && !plan.TaxType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxType", "TaxType", plan.TaxType.ValueString()) {
			if !plan.TaxType.Equal(state.TaxType) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tax_type", params)
			}
		}
	}

	return params, nil
}

func flattenTaxRate(obj *stripe.TaxRate, state *TaxRateResourceModel) error {
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
		if rawValueCountry, rawOk := plainValueAtPath(raw, "country"); rawOk {
			if valueCountry, err := flattenPlainValue(rawValueCountry, types.StringType, "country", "raw response"); err != nil {
				return err
			} else {
				if typedCountry, ok := valueCountry.(types.String); ok {
					state.Country = typedCountry
				}
			}
		} else if !hasRaw {
			if responseValueCountry, ok := plainFromResponseField(obj, "Country"); ok {
				if valueCountry, err := flattenPlainValue(responseValueCountry, types.StringType, "country", "response struct"); err != nil {
					return err
				} else {
					if typedCountry, ok := valueCountry.(types.String); ok {
						state.Country = typedCountry
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
		if rawValueDescription, rawOk := plainValueAtPath(raw, "description"); rawOk {
			if valueDescription, err := flattenPlainValue(rawValueDescription, types.StringType, "description", "raw response"); err != nil {
				return err
			} else {
				if typedDescription, ok := valueDescription.(types.String); ok {
					state.Description = typedDescription
				}
			}
		} else if !hasRaw {
			if responseValueDescription, ok := plainFromResponseField(obj, "Description"); ok {
				if valueDescription, err := flattenPlainValue(responseValueDescription, types.StringType, "description", "response struct"); err != nil {
					return err
				} else {
					if typedDescription, ok := valueDescription.(types.String); ok {
						state.Description = typedDescription
					}
				}
			}
		}
	}
	{
		if rawValueDisplayName, rawOk := plainValueAtPath(raw, "display_name"); rawOk {
			if valueDisplayName, err := flattenPlainValue(rawValueDisplayName, types.StringType, "display_name", "raw response"); err != nil {
				return err
			} else {
				if typedDisplayName, ok := valueDisplayName.(types.String); ok {
					state.DisplayName = typedDisplayName
				}
			}
		} else if !hasRaw {
			if responseValueDisplayName, ok := plainFromResponseField(obj, "DisplayName"); ok {
				if valueDisplayName, err := flattenPlainValue(responseValueDisplayName, types.StringType, "display_name", "response struct"); err != nil {
					return err
				} else {
					if typedDisplayName, ok := valueDisplayName.(types.String); ok {
						state.DisplayName = typedDisplayName
					}
				}
			}
		}
	}
	{
		if rawValueEffectivePercentage, rawOk := plainValueAtPath(raw, "effective_percentage"); rawOk {
			if valueEffectivePercentage, err := flattenPlainValue(rawValueEffectivePercentage, types.Float64Type, "effective_percentage", "raw response"); err != nil {
				return err
			} else {
				if typedEffectivePercentage, ok := valueEffectivePercentage.(types.Float64); ok {
					state.EffectivePercentage = typedEffectivePercentage
				}
			}
		} else if !hasRaw {
			if responseValueEffectivePercentage, ok := plainFromResponseField(obj, "EffectivePercentage"); ok {
				if valueEffectivePercentage, err := flattenPlainValue(responseValueEffectivePercentage, types.Float64Type, "effective_percentage", "response struct"); err != nil {
					return err
				} else {
					if typedEffectivePercentage, ok := valueEffectivePercentage.(types.Float64); ok {
						state.EffectivePercentage = typedEffectivePercentage
					}
				}
			}
		}
	}
	{
		assignedFlatAmount := false
		hadRawFlatAmount := false
		if rawValueFlatAmount, rawOk := plainValueAtPath(raw, "flat_amount"); rawOk {
			hadRawFlatAmount = true
			if rawValueFlatAmount != nil {
				sourceFlatAmount := applyConfiguredKeyedListShapes(rawValueFlatAmount, attrValueToPlain(state.FlatAmount))
				if valueFlatAmount, err := flattenPlainValue(sourceFlatAmount, types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType}}, "flat_amount", "raw response"); err != nil {
					return err
				} else {
					if typedFlatAmount, ok := valueFlatAmount.(types.Object); ok {
						state.FlatAmount = typedFlatAmount
						assignedFlatAmount = true
					}
				}
			}
		}
		if !assignedFlatAmount {
			if !hasRaw {
				if responseValueFlatAmount, ok := plainFromResponseField(obj, "FlatAmount"); ok {
					sourceFlatAmount := applyConfiguredKeyedListShapes(responseValueFlatAmount, attrValueToPlain(state.FlatAmount))
					if valueFlatAmount, err := flattenPlainValue(
						sourceFlatAmount,
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType}},
						"flat_amount",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFlatAmount, ok := valueFlatAmount.(types.Object); ok {
							state.FlatAmount = typedFlatAmount
							assignedFlatAmount = true
						}
					}
				}
			}
		}
		if !assignedFlatAmount && hadRawFlatAmount {
			if nullFlatAmount, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType}}); ok {
				if typedFlatAmount, ok := nullFlatAmount.(types.Object); ok {
					state.FlatAmount = typedFlatAmount
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
		if rawValueInclusive, rawOk := plainValueAtPath(raw, "inclusive"); rawOk {
			if valueInclusive, err := flattenPlainValue(rawValueInclusive, types.BoolType, "inclusive", "raw response"); err != nil {
				return err
			} else {
				if typedInclusive, ok := valueInclusive.(types.Bool); ok {
					state.Inclusive = typedInclusive
				}
			}
		} else if !hasRaw {
			if responseValueInclusive, ok := plainFromResponseField(obj, "Inclusive"); ok {
				if valueInclusive, err := flattenPlainValue(responseValueInclusive, types.BoolType, "inclusive", "response struct"); err != nil {
					return err
				} else {
					if typedInclusive, ok := valueInclusive.(types.Bool); ok {
						state.Inclusive = typedInclusive
					}
				}
			}
		}
	}
	{
		if rawValueJurisdiction, rawOk := plainValueAtPath(raw, "jurisdiction"); rawOk {
			if valueJurisdiction, err := flattenPlainValue(rawValueJurisdiction, types.StringType, "jurisdiction", "raw response"); err != nil {
				return err
			} else {
				if typedJurisdiction, ok := valueJurisdiction.(types.String); ok {
					state.Jurisdiction = typedJurisdiction
				}
			}
		} else if !hasRaw {
			if responseValueJurisdiction, ok := plainFromResponseField(obj, "Jurisdiction"); ok {
				if valueJurisdiction, err := flattenPlainValue(responseValueJurisdiction, types.StringType, "jurisdiction", "response struct"); err != nil {
					return err
				} else {
					if typedJurisdiction, ok := valueJurisdiction.(types.String); ok {
						state.Jurisdiction = typedJurisdiction
					}
				}
			}
		}
	}
	{
		if rawValueJurisdictionLevel, rawOk := plainValueAtPath(raw, "jurisdiction_level"); rawOk {
			if valueJurisdictionLevel, err := flattenPlainValue(rawValueJurisdictionLevel, types.StringType, "jurisdiction_level", "raw response"); err != nil {
				return err
			} else {
				if typedJurisdictionLevel, ok := valueJurisdictionLevel.(types.String); ok {
					state.JurisdictionLevel = typedJurisdictionLevel
				}
			}
		} else if !hasRaw {
			if responseValueJurisdictionLevel, ok := plainFromResponseField(obj, "JurisdictionLevel"); ok {
				if valueJurisdictionLevel, err := flattenPlainValue(responseValueJurisdictionLevel, types.StringType, "jurisdiction_level", "response struct"); err != nil {
					return err
				} else {
					if typedJurisdictionLevel, ok := valueJurisdictionLevel.(types.String); ok {
						state.JurisdictionLevel = typedJurisdictionLevel
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
		if rawValuePercentage, rawOk := plainValueAtPath(raw, "percentage"); rawOk {
			if valuePercentage, err := flattenPlainValue(rawValuePercentage, types.Float64Type, "percentage", "raw response"); err != nil {
				return err
			} else {
				if typedPercentage, ok := valuePercentage.(types.Float64); ok {
					state.Percentage = typedPercentage
				}
			}
		} else if !hasRaw {
			if responseValuePercentage, ok := plainFromResponseField(obj, "Percentage"); ok {
				if valuePercentage, err := flattenPlainValue(responseValuePercentage, types.Float64Type, "percentage", "response struct"); err != nil {
					return err
				} else {
					if typedPercentage, ok := valuePercentage.(types.Float64); ok {
						state.Percentage = typedPercentage
					}
				}
			}
		}
	}
	{
		if rawValueRateType, rawOk := plainValueAtPath(raw, "rate_type"); rawOk {
			if valueRateType, err := flattenPlainValue(rawValueRateType, types.StringType, "rate_type", "raw response"); err != nil {
				return err
			} else {
				if typedRateType, ok := valueRateType.(types.String); ok {
					state.RateType = typedRateType
				}
			}
		} else if !hasRaw {
			if responseValueRateType, ok := plainFromResponseField(obj, "RateType"); ok {
				if valueRateType, err := flattenPlainValue(responseValueRateType, types.StringType, "rate_type", "response struct"); err != nil {
					return err
				} else {
					if typedRateType, ok := valueRateType.(types.String); ok {
						state.RateType = typedRateType
					}
				}
			}
		}
	}
	{
		if rawValueState, rawOk := plainValueAtPath(raw, "state"); rawOk {
			if valueState, err := flattenPlainValue(rawValueState, types.StringType, "state", "raw response"); err != nil {
				return err
			} else {
				if typedState, ok := valueState.(types.String); ok {
					state.State = typedState
				}
			}
		} else if !hasRaw {
			if responseValueState, ok := plainFromResponseField(obj, "State"); ok {
				if valueState, err := flattenPlainValue(responseValueState, types.StringType, "state", "response struct"); err != nil {
					return err
				} else {
					if typedState, ok := valueState.(types.String); ok {
						state.State = typedState
					}
				}
			}
		}
	}
	{
		if rawValueTaxType, rawOk := plainValueAtPath(raw, "tax_type"); rawOk {
			if valueTaxType, err := flattenPlainValue(rawValueTaxType, types.StringType, "tax_type", "raw response"); err != nil {
				return err
			} else {
				if typedTaxType, ok := valueTaxType.(types.String); ok {
					state.TaxType = typedTaxType
				}
			}
		} else if !hasRaw {
			if responseValueTaxType, ok := plainFromResponseField(obj, "TaxType"); ok {
				if valueTaxType, err := flattenPlainValue(responseValueTaxType, types.StringType, "tax_type", "response struct"); err != nil {
					return err
				} else {
					if typedTaxType, ok := valueTaxType.(types.String); ok {
						state.TaxType = typedTaxType
					}
				}
			}
		}
	}
	return nil
}
