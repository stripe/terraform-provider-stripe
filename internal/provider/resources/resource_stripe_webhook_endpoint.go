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

var _ resource.Resource = &WebhookEndpointResource{}

var _ resource.ResourceWithConfigure = &WebhookEndpointResource{}

var _ resource.ResourceWithImportState = &WebhookEndpointResource{}

func NewWebhookEndpointResource() resource.Resource {
	return &WebhookEndpointResource{}
}

type WebhookEndpointResource struct {
	client *stripe.Client
}

type WebhookEndpointResourceModel struct {
	Object        types.String `tfsdk:"object"`
	APIVersion    types.String `tfsdk:"api_version"`
	Application   types.String `tfsdk:"application"`
	Created       types.Int64  `tfsdk:"created"`
	Description   types.String `tfsdk:"description"`
	EnabledEvents types.List   `tfsdk:"enabled_events"`
	ID            types.String `tfsdk:"id"`
	Livemode      types.Bool   `tfsdk:"livemode"`
	Metadata      types.Map    `tfsdk:"metadata"`
	Secret        types.String `tfsdk:"secret"`
	Status        types.String `tfsdk:"status"`
	URL           types.String `tfsdk:"url"`
	Connect       types.Bool   `tfsdk:"connect"`
}

func (r *WebhookEndpointResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WebhookEndpointResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook_endpoint"
}

var _ resource.ResourceWithUpgradeState = &WebhookEndpointResource{}

func (r *WebhookEndpointResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: webhookEndpointResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior WebhookEndpointResourceV0Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeWebhookEndpointStateV0(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func webhookEndpointResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "You can configure [webhook endpoints](https://docs.stripe.com/webhooks/) via the API to be\nnotified about events that happen in your Stripe account or connected\naccounts.\n\nMost users configure webhooks from [the dashboard](https://dashboard.stripe.com/webhooks), which provides a user interface for registering and testing your webhook endpoints.\n\nRelated guide: [Setting up webhooks](https://docs.stripe.com/webhooks/configure)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("webhook_endpoint")},
			},
			"api_version": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The API version events are rendered as for this webhook endpoint.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the associated Connect application.",
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
				Description:   "An optional description of what the webhook is used for.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled_events": schema.ListAttribute{
				Required:    true,
				Description: "The list of events to enable for this endpoint. `['*']` indicates that all events are enabled, except those that require explicit selection.",
				ElementType: types.StringType,
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
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"secret": schema.StringAttribute{
				Computed:      true,
				Description:   "The endpoint's secret, used to generate [webhook signatures](https://docs.stripe.com/webhooks/signatures). Only returned at creation.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Sensitive:     true,
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The status of the webhook. It can be `enabled` or `disabled`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"url": schema.StringAttribute{
				Required:    true,
				Description: "The URL of the webhook endpoint.",
			},
			"connect": schema.BoolAttribute{
				Optional:      true,
				Description:   "Whether this endpoint should receive events from connected accounts (`true`), or from your account (`false`). Defaults to `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
		},
	}
}

type WebhookEndpointResourceV0Model struct {
	Object        types.String `tfsdk:"object"`
	APIVersion    types.String `tfsdk:"api_version"`
	Application   types.String `tfsdk:"application"`
	Created       types.Int64  `tfsdk:"created"`
	Description   types.String `tfsdk:"description"`
	EnabledEvents types.List   `tfsdk:"enabled_events"`
	ID            types.String `tfsdk:"id"`
	Livemode      types.Bool   `tfsdk:"livemode"`
	Metadata      types.Map    `tfsdk:"metadata"`
	Secret        types.String `tfsdk:"secret"`
	Status        types.String `tfsdk:"status"`
	URL           types.String `tfsdk:"url"`
	Connect       types.Bool   `tfsdk:"connect"`
}

type webhookendpointStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]webhookendpointStateUpgradeAttrMeta
}

var webhookendpointStateUpgradeRootMeta = map[string]webhookendpointStateUpgradeAttrMeta{"object": webhookendpointStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "api_version": webhookendpointStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "application": webhookendpointStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "created": webhookendpointStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "description": webhookendpointStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "enabled_events": webhookendpointStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.StringType}, Behavior: "required", LegacyBehavior: "required"}, "id": webhookendpointStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "livemode": webhookendpointStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "metadata": webhookendpointStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "secret": webhookendpointStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "status": webhookendpointStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "url": webhookendpointStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "connect": webhookendpointStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional", LegacyBehavior: "optional", PreserveConfiguredValue: true}}

var webhookendpointStateUpgradeSingletonPaths = map[string]struct{}{}

var webhookendpointStateUpgradeLegacyObjectPaths = map[string]struct{}{}

func webhookendpointAttrMapFromModel(model interface{}) map[string]attr.Value {
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

func webhookendpointSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
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

func webhookendpointIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func webhookendpointShouldPreserveChild(parent webhookendpointStateUpgradeAttrMeta, child webhookendpointStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !webhookendpointIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !webhookendpointIsComputedBehavior(child.Behavior)
}

func webhookendpointNullValueForType(attributeType attr.Type) attr.Value {
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

func webhookendpointLegacyUpgradeIsEmptyValue(value attr.Value) bool {
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

func webhookendpointLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
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

func webhookendpointLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := webhookendpointLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func webhookendpointLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
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

func webhookendpointLegacyUpgradeNormalizeChild(parent webhookendpointStateUpgradeAttrMeta, name string, child webhookendpointStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if webhookendpointLegacyUpgradeIsEmptyValue(value) {
		return webhookendpointNullValueForType(child.AttrType)
	}
	if webhookendpointLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return webhookendpointNullValueForType(child.AttrType)
	}
	if webhookendpointLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return webhookendpointNullValueForType(child.AttrType)
	}
	return value
}

func webhookendpointLegacyUpgradeChildAttr(path []string, parent webhookendpointStateUpgradeAttrMeta, name string, child webhookendpointStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !webhookendpointShouldPreserveChild(parent, child) {
		return webhookendpointNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return webhookendpointNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := webhookendpointUpgradeValue(nextPath, child, childValue)
	return webhookendpointLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func webhookendpointUpgradeAttrs(path []string, meta map[string]webhookendpointStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = webhookendpointNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := webhookendpointUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = webhookendpointLegacyUpgradeNormalizeChild(
				webhookendpointStateUpgradeAttrMeta{PreserveConfiguredValue: true},
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

func webhookendpointUpgradeObjectValue(path []string, meta webhookendpointStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = webhookendpointLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func webhookendpointUpgradeSingletonListToObject(path []string, meta webhookendpointStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return webhookendpointNullValueForType(meta.AttrType)
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
		upgradedAttrs[name] = webhookendpointLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func webhookendpointUpgradeObjectValueToSingletonList(path []string, meta webhookendpointStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	objectValue, ok := priorValue.(types.Object)
	if !ok {
		if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
			objectValue = types.Object(baseObject)
		} else {
			return webhookendpointNullValueForType(meta.AttrType)
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
		return webhookendpointNullValueForType(meta.AttrType)
	}

	upgradedObject := webhookendpointUpgradeObjectValue(path, meta, elementObjectType, objectValue)
	return types.ListValueMust(listType.ElemType, []attr.Value{upgradedObject})
}

func webhookendpointUpgradeListValue(path []string, meta webhookendpointStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return webhookendpointNullValueForType(meta.AttrType)
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
				upgradedElements = append(upgradedElements, webhookendpointNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			webhookendpointUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func webhookendpointUpgradeValue(path []string, meta webhookendpointStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := webhookendpointStateUpgradeSingletonPaths[pathKey]; ok {
			return webhookendpointUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return webhookendpointNullValueForType(meta.AttrType)
			}
		}
		return webhookendpointUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		if _, ok := webhookendpointStateUpgradeLegacyObjectPaths[pathKey]; ok {
			return webhookendpointUpgradeObjectValueToSingletonList(path, meta, attrType, priorValue)
		}
		return webhookendpointUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradeWebhookEndpointStateV0(ctx context.Context, prior WebhookEndpointResourceV0Model) (WebhookEndpointResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := webhookendpointUpgradeAttrs(nil, webhookendpointStateUpgradeRootMeta, webhookendpointAttrMapFromModel(prior))
	var upgraded WebhookEndpointResourceModel
	webhookendpointSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *WebhookEndpointResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     1,
		Description: "You can configure [webhook endpoints](https://docs.stripe.com/webhooks/) via the API to be\nnotified about events that happen in your Stripe account or connected\naccounts.\n\nMost users configure webhooks from [the dashboard](https://dashboard.stripe.com/webhooks), which provides a user interface for registering and testing your webhook endpoints.\n\nRelated guide: [Setting up webhooks](https://docs.stripe.com/webhooks/configure)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("webhook_endpoint")},
			},
			"api_version": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The API version events are rendered as for this webhook endpoint.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the associated Connect application.",
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
				Description:   "An optional description of what the webhook is used for.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled_events": schema.ListAttribute{
				Required:    true,
				Description: "The list of events to enable for this endpoint. `['*']` indicates that all events are enabled, except those that require explicit selection.",
				ElementType: types.StringType,
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
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"secret": schema.StringAttribute{
				Computed:      true,
				Description:   "The endpoint's secret, used to generate [webhook signatures](https://docs.stripe.com/webhooks/signatures). Only returned at creation.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Sensitive:     true,
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The status of the webhook. It can be `enabled` or `disabled`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"url": schema.StringAttribute{
				Required:    true,
				Description: "The URL of the webhook endpoint.",
			},
			"connect": schema.BoolAttribute{
				Optional:      true,
				Description:   "Whether this endpoint should receive events from connected accounts (`true`), or from your account (`false`). Defaults to `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *WebhookEndpointResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WebhookEndpointResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandWebhookEndpointCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building WebhookEndpoint create params", err.Error())
		return
	}

	obj, err := r.client.V1WebhookEndpoints.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating WebhookEndpoint", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1WebhookEndpoints.B, r.client.V1WebhookEndpoints.Key, stripe.FormatURLPath("/v1/webhook_endpoints/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating WebhookEndpoint create raw response", err.Error())
		return
	}

	if err := flattenWebhookEndpoint(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening WebhookEndpoint create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *WebhookEndpointResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState WebhookEndpointResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state WebhookEndpointResourceModel
	state = priorState

	obj, err := r.client.V1WebhookEndpoints.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading WebhookEndpoint", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1WebhookEndpoints.B, r.client.V1WebhookEndpoints.Key, stripe.FormatURLPath("/v1/webhook_endpoints/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating WebhookEndpoint raw response", err.Error())
		return
	}

	if err := flattenWebhookEndpoint(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening WebhookEndpoint read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *WebhookEndpointResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan WebhookEndpointResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state WebhookEndpointResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandWebhookEndpointUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building WebhookEndpoint update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building WebhookEndpoint update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1WebhookEndpoints.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating WebhookEndpoint", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1WebhookEndpoints.B, r.client.V1WebhookEndpoints.Key, stripe.FormatURLPath("/v1/webhook_endpoints/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating WebhookEndpoint update raw response", err.Error())
		return
	}

	if err := flattenWebhookEndpoint(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening WebhookEndpoint update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *WebhookEndpointResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WebhookEndpointResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1WebhookEndpoints.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting WebhookEndpoint", err.Error())
		return
	}
}

func (r *WebhookEndpointResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandWebhookEndpointCreate(plan WebhookEndpointResourceModel) (*stripe.WebhookEndpointCreateParams, error) {
	params := &stripe.WebhookEndpointCreateParams{}

	if !plan.APIVersion.IsNull() && !plan.APIVersion.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "APIVersion", "APIVersion", plan.APIVersion.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "api_version", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.EnabledEvents.IsNull() && !plan.EnabledEvents.IsUnknown() {
		if !assignAttrValueToNamedField(params, "EnabledEvents", plan.EnabledEvents) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "enabled_events", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.URL.IsNull() && !plan.URL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "URL", "URL", plan.URL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "url", params)
		}
	}
	if !plan.Connect.IsNull() && !plan.Connect.IsUnknown() {
		params.Connect = stripe.Bool(plan.Connect.ValueBool())
	}

	return params, nil
}

func expandWebhookEndpointUpdate(plan WebhookEndpointResourceModel, state WebhookEndpointResourceModel) (*stripe.WebhookEndpointUpdateParams, error) {
	params := &stripe.WebhookEndpointUpdateParams{}

	if !plan.Description.Equal(state.Description) && !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			if !plan.Description.Equal(state.Description) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "description", params)
			}
		}
	}
	if !plan.EnabledEvents.Equal(state.EnabledEvents) && !plan.EnabledEvents.IsNull() && !plan.EnabledEvents.IsUnknown() {
		if !assignAttrValueToNamedField(params, "EnabledEvents", plan.EnabledEvents) {
			if !plan.EnabledEvents.Equal(state.EnabledEvents) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "enabled_events", params)
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
	if !plan.URL.Equal(state.URL) && !plan.URL.IsNull() && !plan.URL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "URL", "URL", plan.URL.ValueString()) {
			if !plan.URL.Equal(state.URL) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "url", params)
			}
		}
	}

	return params, nil
}

func flattenWebhookEndpoint(obj *stripe.WebhookEndpoint, state *WebhookEndpointResourceModel) error {
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
		if rawValueAPIVersion, rawOk := plainValueAtPath(raw, "api_version"); rawOk {
			if valueAPIVersion, err := flattenPlainValue(rawValueAPIVersion, types.StringType, "api_version", "raw response"); err != nil {
				return err
			} else {
				if typedAPIVersion, ok := valueAPIVersion.(types.String); ok {
					state.APIVersion = typedAPIVersion
				}
			}
		} else if !hasRaw {
			if responseValueAPIVersion, ok := plainFromResponseField(obj, "APIVersion"); ok {
				if valueAPIVersion, err := flattenPlainValue(responseValueAPIVersion, types.StringType, "api_version", "response struct"); err != nil {
					return err
				} else {
					if typedAPIVersion, ok := valueAPIVersion.(types.String); ok {
						state.APIVersion = typedAPIVersion
					}
				}
			}
		}
	}
	{
		if rawValueApplication, rawOk := plainValueAtPath(raw, "application"); rawOk {
			if valueApplication, err := flattenPlainValue(rawValueApplication, types.StringType, "application", "raw response"); err != nil {
				return err
			} else {
				if typedApplication, ok := valueApplication.(types.String); ok {
					state.Application = typedApplication
				}
			}
		} else if !hasRaw {
			if responseValueApplication, ok := plainFromResponseField(obj, "Application"); ok {
				if valueApplication, err := flattenPlainValue(responseValueApplication, types.StringType, "application", "response struct"); err != nil {
					return err
				} else {
					if typedApplication, ok := valueApplication.(types.String); ok {
						state.Application = typedApplication
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
		if rawValueEnabledEvents, rawOk := plainValueAtPath(raw, "enabled_events"); rawOk {
			if valueEnabledEvents, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueEnabledEvents, attrValueToPlain(state.EnabledEvents)), types.ListType{ElemType: types.StringType}, "enabled_events", "raw response"); err != nil {
				return err
			} else {
				if typedEnabledEvents, ok := valueEnabledEvents.(types.List); ok {
					state.EnabledEvents = typedEnabledEvents
				}
			}
		} else if !hasRaw {
			if responseValueEnabledEvents, ok := plainFromResponseField(obj, "EnabledEvents"); ok {
				if valueEnabledEvents, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueEnabledEvents, attrValueToPlain(state.EnabledEvents)),
					types.ListType{ElemType: types.StringType},
					"enabled_events",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedEnabledEvents, ok := valueEnabledEvents.(types.List); ok {
						state.EnabledEvents = typedEnabledEvents
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
		if rawValueSecret, rawOk := plainValueAtPath(raw, "secret"); rawOk {
			if valueSecret, err := flattenPlainValue(rawValueSecret, types.StringType, "secret", "raw response"); err != nil {
				return err
			} else {
				if typedSecret, ok := valueSecret.(types.String); ok {
					state.Secret = typedSecret
				}
			}
		} else if !hasRaw {
			if responseValueSecret, ok := plainFromResponseField(obj, "Secret"); ok {
				if valueSecret, err := flattenPlainValue(responseValueSecret, types.StringType, "secret", "response struct"); err != nil {
					return err
				} else {
					if typedSecret, ok := valueSecret.(types.String); ok {
						state.Secret = typedSecret
					}
				}
			}
		}
	}
	{
		if rawValueStatus, rawOk := plainValueAtPath(raw, "status"); rawOk {
			if valueStatus, err := flattenPlainValue(rawValueStatus, types.StringType, "status", "raw response"); err != nil {
				return err
			} else {
				if typedStatus, ok := valueStatus.(types.String); ok {
					state.Status = typedStatus
				}
			}
		} else if !hasRaw {
			if responseValueStatus, ok := plainFromResponseField(obj, "Status"); ok {
				if valueStatus, err := flattenPlainValue(responseValueStatus, types.StringType, "status", "response struct"); err != nil {
					return err
				} else {
					if typedStatus, ok := valueStatus.(types.String); ok {
						state.Status = typedStatus
					}
				}
			}
		}
	}
	{
		if rawValueURL, rawOk := plainValueAtPath(raw, "url"); rawOk {
			if valueURL, err := flattenPlainValue(rawValueURL, types.StringType, "url", "raw response"); err != nil {
				return err
			} else {
				if typedURL, ok := valueURL.(types.String); ok {
					state.URL = typedURL
				}
			}
		} else if !hasRaw {
			if responseValueURL, ok := plainFromResponseField(obj, "URL"); ok {
				if valueURL, err := flattenPlainValue(responseValueURL, types.StringType, "url", "response struct"); err != nil {
					return err
				} else {
					if typedURL, ok := valueURL.(types.String); ok {
						state.URL = typedURL
					}
				}
			}
		}
	}
	{
		if rawValueConnect, rawOk := plainValueAtPath(raw, "connect"); rawOk {
			if !state.Connect.IsNull() && !state.Connect.IsUnknown() {
				if valueConnect, err := flattenPlainValue(rawValueConnect, types.BoolType, "connect", "raw response"); err != nil {
					return err
				} else {
					if typedConnect, ok := valueConnect.(types.Bool); ok {
						state.Connect = typedConnect
					}
				}
			}
		} else if !hasRaw {
			if responseValueConnect, ok := plainFromResponseField(obj, "Connect"); ok {
				if !state.Connect.IsNull() && !state.Connect.IsUnknown() {
					if valueConnect, err := flattenPlainValue(responseValueConnect, types.BoolType, "connect", "response struct"); err != nil {
						return err
					} else {
						if typedConnect, ok := valueConnect.(types.Bool); ok {
							state.Connect = typedConnect
						}
					}
				}
			}
		}
	}
	return nil
}
