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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
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

var _ resource.Resource = &EntitlementsFeatureResource{}

var _ resource.ResourceWithConfigure = &EntitlementsFeatureResource{}

var _ resource.ResourceWithImportState = &EntitlementsFeatureResource{}

func NewEntitlementsFeatureResource() resource.Resource {
	return &EntitlementsFeatureResource{}
}

type EntitlementsFeatureResource struct {
	client *stripe.Client
}

type EntitlementsFeatureResourceModel struct {
	Object    types.String `tfsdk:"object"`
	Active    types.Bool   `tfsdk:"active"`
	ID        types.String `tfsdk:"id"`
	Livemode  types.Bool   `tfsdk:"livemode"`
	LookupKey types.String `tfsdk:"lookup_key"`
	Metadata  types.Map    `tfsdk:"metadata"`
	Name      types.String `tfsdk:"name"`
}

func (r *EntitlementsFeatureResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *EntitlementsFeatureResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_entitlements_feature"
}

var _ resource.ResourceWithUpgradeState = &EntitlementsFeatureResource{}

func (r *EntitlementsFeatureResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: entitlementsFeatureResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior EntitlementsFeatureResourceV0Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeEntitlementsFeatureStateV0(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func entitlementsFeatureResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "A feature represents a monetizable ability or functionality in your system.\nFeatures can be assigned to products, and when those products are purchased, Stripe will create an entitlement to the feature for the purchasing customer.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("entitlements.feature")},
			},
			"active": schema.BoolAttribute{
				Computed:      true,
				Description:   "Inactive features cannot be attached to new products and will not be returned from the features list endpoint.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
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
			"lookup_key": schema.StringAttribute{
				Required:      true,
				Description:   "A unique key you provide as your own system identifier. This may be up to 80 characters.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The feature's name, for your own purpose, not meant to be displayable to the customer.",
			},
		},
	}
}

type EntitlementsFeatureResourceV0Model struct {
	Object    types.String `tfsdk:"object"`
	Active    types.Bool   `tfsdk:"active"`
	ID        types.String `tfsdk:"id"`
	Livemode  types.Bool   `tfsdk:"livemode"`
	LookupKey types.String `tfsdk:"lookup_key"`
	Metadata  types.Map    `tfsdk:"metadata"`
	Name      types.String `tfsdk:"name"`
}

type entitlementsfeatureStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]entitlementsfeatureStateUpgradeAttrMeta
}

var entitlementsfeatureStateUpgradeRootMeta = map[string]entitlementsfeatureStateUpgradeAttrMeta{"object": entitlementsfeatureStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "active": entitlementsfeatureStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "id": entitlementsfeatureStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "livemode": entitlementsfeatureStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "lookup_key": entitlementsfeatureStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "metadata": entitlementsfeatureStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "name": entitlementsfeatureStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}

var entitlementsfeatureStateUpgradeSingletonPaths = map[string]struct{}{}

func entitlementsfeatureAttrMapFromModel(model interface{}) map[string]attr.Value {
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

func entitlementsfeatureSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
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

func entitlementsfeatureIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func entitlementsfeatureShouldPreserveChild(parent entitlementsfeatureStateUpgradeAttrMeta, child entitlementsfeatureStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !entitlementsfeatureIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !entitlementsfeatureIsComputedBehavior(child.Behavior)
}

func entitlementsfeatureNullValueForType(attributeType attr.Type) attr.Value {
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

func entitlementsfeatureLegacyUpgradeIsEmptyValue(value attr.Value) bool {
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

func entitlementsfeatureLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
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

func entitlementsfeatureLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := entitlementsfeatureLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func entitlementsfeatureLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
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

func entitlementsfeatureLegacyUpgradeNormalizeChild(parent entitlementsfeatureStateUpgradeAttrMeta, name string, child entitlementsfeatureStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if entitlementsfeatureLegacyUpgradeIsEmptyValue(value) {
		return entitlementsfeatureNullValueForType(child.AttrType)
	}
	if entitlementsfeatureLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return entitlementsfeatureNullValueForType(child.AttrType)
	}
	if entitlementsfeatureLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return entitlementsfeatureNullValueForType(child.AttrType)
	}
	return value
}

func entitlementsfeatureLegacyUpgradeChildAttr(path []string, parent entitlementsfeatureStateUpgradeAttrMeta, name string, child entitlementsfeatureStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !entitlementsfeatureShouldPreserveChild(parent, child) {
		return entitlementsfeatureNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return entitlementsfeatureNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := entitlementsfeatureUpgradeValue(nextPath, child, childValue)
	return entitlementsfeatureLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func entitlementsfeatureUpgradeAttrs(path []string, meta map[string]entitlementsfeatureStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = entitlementsfeatureNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := entitlementsfeatureUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = entitlementsfeatureLegacyUpgradeNormalizeChild(
				entitlementsfeatureStateUpgradeAttrMeta{PreserveConfiguredValue: true},
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

func entitlementsfeatureUpgradeObjectValue(path []string, meta entitlementsfeatureStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = entitlementsfeatureLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func entitlementsfeatureUpgradeSingletonListToObject(path []string, meta entitlementsfeatureStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return entitlementsfeatureNullValueForType(meta.AttrType)
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
		upgradedAttrs[name] = entitlementsfeatureLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func entitlementsfeatureUpgradeListValue(path []string, meta entitlementsfeatureStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return entitlementsfeatureNullValueForType(meta.AttrType)
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
				upgradedElements = append(upgradedElements, entitlementsfeatureNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			entitlementsfeatureUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func entitlementsfeatureUpgradeValue(path []string, meta entitlementsfeatureStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := entitlementsfeatureStateUpgradeSingletonPaths[pathKey]; ok {
			return entitlementsfeatureUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return entitlementsfeatureNullValueForType(meta.AttrType)
			}
		}
		return entitlementsfeatureUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		return entitlementsfeatureUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradeEntitlementsFeatureStateV0(ctx context.Context, prior EntitlementsFeatureResourceV0Model) (EntitlementsFeatureResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := entitlementsfeatureUpgradeAttrs(nil, entitlementsfeatureStateUpgradeRootMeta, entitlementsfeatureAttrMapFromModel(prior))
	var upgraded EntitlementsFeatureResourceModel
	entitlementsfeatureSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *EntitlementsFeatureResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     1,
		Description: "A feature represents a monetizable ability or functionality in your system.\nFeatures can be assigned to products, and when those products are purchased, Stripe will create an entitlement to the feature for the purchasing customer.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("entitlements.feature")},
			},
			"active": schema.BoolAttribute{
				Computed:      true,
				Description:   "Inactive features cannot be attached to new products and will not be returned from the features list endpoint.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
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
			"lookup_key": schema.StringAttribute{
				Required:      true,
				Description:   "A unique key you provide as your own system identifier. This may be up to 80 characters.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The feature's name, for your own purpose, not meant to be displayable to the customer.",
			},
		},
	}
}

func (r *EntitlementsFeatureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan EntitlementsFeatureResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandEntitlementsFeatureCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building EntitlementsFeature create params", err.Error())
		return
	}

	obj, err := r.client.V1EntitlementsFeatures.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating EntitlementsFeature", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1EntitlementsFeatures.B, r.client.V1EntitlementsFeatures.Key, stripe.FormatURLPath("/v1/entitlements/features/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating EntitlementsFeature create raw response", err.Error())
		return
	}

	var createdState EntitlementsFeatureResourceModel
	if err := flattenEntitlementsFeature(obj, &createdState); err != nil {
		resp.Diagnostics.AddError("Error flattening EntitlementsFeature create response", err.Error())
		return
	}
	normalizeUnknownValues(&createdState)

	diffPlan := plan
	diffCreatedState := createdState

	postCreateParams, err := expandEntitlementsFeaturePostCreateUpdate(diffPlan, diffCreatedState)
	if err != nil {
		resp.Diagnostics.AddError("Error building EntitlementsFeature post-create update params", err.Error())
		return
	}

	if paramsHaveValues(postCreateParams) {
		if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
			!createdState.Metadata.IsNull() && !createdState.Metadata.IsUnknown() {
			if !assignMetadataDiffToNamedField(postCreateParams, "Metadata", plan.Metadata, createdState.Metadata) {
				resp.Diagnostics.AddError("Error building EntitlementsFeature update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", postCreateParams))
				return
			}
		}
		obj, err = r.client.V1EntitlementsFeatures.Update(ctx, createdState.ID.ValueString(), postCreateParams)
		if err != nil {
			resp.Diagnostics.AddError("Error finalizing EntitlementsFeature after create", err.Error())
			return
		}
		if err := ensureRawResponse(obj, r.client.V1EntitlementsFeatures.B, r.client.V1EntitlementsFeatures.Key, stripe.FormatURLPath("/v1/entitlements/features/%s", obj.ID), nil); err != nil {
			resp.Diagnostics.AddError("Error hydrating EntitlementsFeature post-create update raw response", err.Error())
			return
		}
	}

	if err := flattenEntitlementsFeature(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening EntitlementsFeature create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *EntitlementsFeatureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState EntitlementsFeatureResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state EntitlementsFeatureResourceModel
	state = priorState

	obj, err := r.client.V1EntitlementsFeatures.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading EntitlementsFeature", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1EntitlementsFeatures.B, r.client.V1EntitlementsFeatures.Key, stripe.FormatURLPath("/v1/entitlements/features/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating EntitlementsFeature raw response", err.Error())
		return
	}

	if err := flattenEntitlementsFeature(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening EntitlementsFeature read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *EntitlementsFeatureResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan EntitlementsFeatureResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state EntitlementsFeatureResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandEntitlementsFeatureUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building EntitlementsFeature update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building EntitlementsFeature update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1EntitlementsFeatures.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating EntitlementsFeature", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1EntitlementsFeatures.B, r.client.V1EntitlementsFeatures.Key, stripe.FormatURLPath("/v1/entitlements/features/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating EntitlementsFeature update raw response", err.Error())
		return
	}

	if err := flattenEntitlementsFeature(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening EntitlementsFeature update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *EntitlementsFeatureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state EntitlementsFeatureResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.Active.IsNull() && !state.Active.IsUnknown() && !state.Active.ValueBool() {
		return
	}

	params := &stripe.EntitlementsFeatureUpdateParams{}
	activeField := reflect.ValueOf(params).Elem().FieldByName("Active")
	if activeField.IsValid() && activeField.CanSet() {
		if activeField.Kind() == reflect.Pointer && activeField.Type().Elem().Kind() == reflect.Bool {
			activeField.Set(reflect.ValueOf(stripe.Bool(false)))
		} else if activeField.Kind() == reflect.Bool {
			activeField.SetBool(false)
		}
	}

	_, err := r.client.V1EntitlementsFeatures.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating EntitlementsFeature", err.Error())
		return
	}
}

func (r *EntitlementsFeatureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandEntitlementsFeatureCreate(plan EntitlementsFeatureResourceModel) (*stripe.EntitlementsFeatureCreateParams, error) {
	params := &stripe.EntitlementsFeatureCreateParams{}

	if !plan.LookupKey.IsNull() && !plan.LookupKey.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LookupKey", "LookupKey", plan.LookupKey.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "lookup_key", params)
		}
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

	return params, nil
}

func expandEntitlementsFeatureUpdate(plan EntitlementsFeatureResourceModel, state EntitlementsFeatureResourceModel) (*stripe.EntitlementsFeatureUpdateParams, error) {
	params := &stripe.EntitlementsFeatureUpdateParams{}

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

func expandEntitlementsFeaturePostCreateUpdate(plan EntitlementsFeatureResourceModel, state EntitlementsFeatureResourceModel) (*stripe.EntitlementsFeatureUpdateParams, error) {
	params := &stripe.EntitlementsFeatureUpdateParams{}

	return params, nil
}

func flattenEntitlementsFeature(obj *stripe.EntitlementsFeature, state *EntitlementsFeatureResourceModel) error {
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
		if rawValueLookupKey, rawOk := plainValueAtPath(raw, "lookup_key"); rawOk {
			if valueLookupKey, err := flattenPlainValue(rawValueLookupKey, types.StringType, "lookup_key", "raw response"); err != nil {
				return err
			} else {
				if typedLookupKey, ok := valueLookupKey.(types.String); ok {
					state.LookupKey = typedLookupKey
				}
			}
		} else if !hasRaw {
			if responseValueLookupKey, ok := plainFromResponseField(obj, "LookupKey"); ok {
				if valueLookupKey, err := flattenPlainValue(responseValueLookupKey, types.StringType, "lookup_key", "response struct"); err != nil {
					return err
				} else {
					if typedLookupKey, ok := valueLookupKey.(types.String); ok {
						state.LookupKey = typedLookupKey
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
	return nil
}
