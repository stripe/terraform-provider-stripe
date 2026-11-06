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

var _ resource.Resource = &BillingMeterResource{}

var _ resource.ResourceWithConfigure = &BillingMeterResource{}

var _ resource.ResourceWithImportState = &BillingMeterResource{}

func NewBillingMeterResource() resource.Resource {
	return &BillingMeterResource{}
}

type BillingMeterResource struct {
	client *stripe.Client
}

type BillingMeterResourceModel struct {
	Object             types.String `tfsdk:"object"`
	Created            types.Int64  `tfsdk:"created"`
	CustomerMapping    types.List   `tfsdk:"customer_mapping"`
	DefaultAggregation types.List   `tfsdk:"default_aggregation"`
	DisplayName        types.String `tfsdk:"display_name"`
	EventName          types.String `tfsdk:"event_name"`
	EventTimeWindow    types.String `tfsdk:"event_time_window"`
	ID                 types.String `tfsdk:"id"`
	Livemode           types.Bool   `tfsdk:"livemode"`
	Status             types.String `tfsdk:"status"`
	StatusTransitions  types.Object `tfsdk:"status_transitions"`
	Updated            types.Int64  `tfsdk:"updated"`
	ValueSettings      types.List   `tfsdk:"value_settings"`
}

func (r *BillingMeterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *BillingMeterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_billing_meter"
}

var _ resource.ResourceWithUpgradeState = &BillingMeterResource{}

func (r *BillingMeterResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: billingMeterResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior BillingMeterResourceModel
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeBillingMeterStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
		1: {
			PriorSchema: billingMeterResourceV1Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior BillingMeterResourceV1Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeBillingMeterStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func billingMeterResourceV1Schema() *schema.Schema {
	return &schema.Schema{
		Description: "Meters specify how to aggregate meter events over a billing period. Meter events represent the actions that customers take in your system. Meters attach to prices and form the basis of the bill.\n\nRelated guide: [Usage based billing](https://docs.stripe.com/billing/subscriptions/usage-based)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("billing.meter")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"customer_mapping": schema.SingleNestedAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"event_payload_key": schema.StringAttribute{
						Required:      true,
						Description:   "The key in the meter event payload to use for mapping the event to a customer.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"type": schema.StringAttribute{
						Required:      true,
						Description:   "The method for mapping a meter event to a customer.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("by_id")},
					},
				},
			},
			"default_aggregation": schema.SingleNestedAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"formula": schema.StringAttribute{
						Required:      true,
						Description:   "Specifies how events are aggregated.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("count", "last", "sum")},
					},
				},
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "The meter's name.",
			},
			"event_name": schema.StringAttribute{
				Required:      true,
				Description:   "The name of the meter event to record usage for. Corresponds with the `event_name` field on meter events.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"event_time_window": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The time window which meter events have been pre-aggregated for, if any.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("day", "hour")},
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
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The meter's status.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "inactive")},
			},
			"status_transitions": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"deactivated_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time the meter was deactivated, if any. Measured in seconds since Unix epoch.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"updated": schema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was last updated. Measured in seconds since the Unix epoch.",
			},
			"value_settings": schema.SingleNestedAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"event_payload_key": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The key in the meter event payload to use as the value for this meter.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
				},
			},
		},
	}
}

func billingMeterResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "Meters specify how to aggregate meter events over a billing period. Meter events represent the actions that customers take in your system. Meters attach to prices and form the basis of the bill.\n\nRelated guide: [Usage based billing](https://docs.stripe.com/billing/subscriptions/usage-based)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("billing.meter")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "The meter's name.",
			},
			"event_name": schema.StringAttribute{
				Required:      true,
				Description:   "The name of the meter event to record usage for. Corresponds with the `event_name` field on meter events.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"event_time_window": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The time window which meter events have been pre-aggregated for, if any.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("day", "hour")},
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
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The meter's status.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "inactive")},
			},
			"status_transitions": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"deactivated_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time the meter was deactivated, if any. Measured in seconds since Unix epoch.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"updated": schema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was last updated. Measured in seconds since the Unix epoch.",
			},
		},
		Blocks: map[string]schema.Block{
			"customer_mapping": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"event_payload_key": schema.StringAttribute{
							Required:      true,
							Description:   "The key in the meter event payload to use for mapping the event to a customer.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"type": schema.StringAttribute{
							Required:      true,
							Description:   "The method for mapping a meter event to a customer.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							Validators:    []validator.String{stringvalidator.OneOf("by_id")},
						},
					},
				},
			},
			"default_aggregation": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"formula": schema.StringAttribute{
							Required:      true,
							Description:   "Specifies how events are aggregated.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							Validators:    []validator.String{stringvalidator.OneOf("count", "last", "sum")},
						},
					},
				},
			},
			"value_settings": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"event_payload_key": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The key in the meter event payload to use as the value for this meter.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
		},
	}
}

type BillingMeterResourceV1Model struct {
	Object             types.String `tfsdk:"object"`
	Created            types.Int64  `tfsdk:"created"`
	CustomerMapping    types.Object `tfsdk:"customer_mapping"`
	DefaultAggregation types.Object `tfsdk:"default_aggregation"`
	DisplayName        types.String `tfsdk:"display_name"`
	EventName          types.String `tfsdk:"event_name"`
	EventTimeWindow    types.String `tfsdk:"event_time_window"`
	ID                 types.String `tfsdk:"id"`
	Livemode           types.Bool   `tfsdk:"livemode"`
	Status             types.String `tfsdk:"status"`
	StatusTransitions  types.Object `tfsdk:"status_transitions"`
	Updated            types.Int64  `tfsdk:"updated"`
	ValueSettings      types.Object `tfsdk:"value_settings"`
}

type billingmeterStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]billingmeterStateUpgradeAttrMeta
}

var billingmeterStateUpgradeRootMeta = map[string]billingmeterStateUpgradeAttrMeta{"object": billingmeterStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "created": billingmeterStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "customer_mapping": billingmeterStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"event_payload_key": types.StringType, "type": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]billingmeterStateUpgradeAttrMeta{"event_payload_key": billingmeterStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "type": billingmeterStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "default_aggregation": billingmeterStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"formula": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]billingmeterStateUpgradeAttrMeta{"formula": billingmeterStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "display_name": billingmeterStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "event_name": billingmeterStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "event_time_window": billingmeterStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "id": billingmeterStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "livemode": billingmeterStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "status": billingmeterStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "status_transitions": billingmeterStateUpgradeAttrMeta{AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{"deactivated_at": types.Int64Type}}, Behavior: "computed", LegacyBehavior: "computed", Nested: map[string]billingmeterStateUpgradeAttrMeta{"deactivated_at": billingmeterStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}}}, "updated": billingmeterStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "value_settings": billingmeterStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"event_payload_key": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]billingmeterStateUpgradeAttrMeta{"event_payload_key": billingmeterStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}}}

var billingmeterStateUpgradeSingletonPaths = map[string]struct{}{}

var billingmeterStateUpgradeLegacyObjectPaths = map[string]struct{}{"customer_mapping": struct{}{}, "default_aggregation": struct{}{}, "value_settings": struct{}{}}

func billingmeterAttrMapFromModel(model interface{}) map[string]attr.Value {
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

func billingmeterSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
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

func billingmeterIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func billingmeterShouldPreserveChild(parent billingmeterStateUpgradeAttrMeta, child billingmeterStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !billingmeterIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !billingmeterIsComputedBehavior(child.Behavior)
}

func billingmeterNullValueForType(attributeType attr.Type) attr.Value {
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

func billingmeterLegacyUpgradeIsEmptyValue(value attr.Value) bool {
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

func billingmeterLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
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

func billingmeterLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := billingmeterLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func billingmeterLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
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

func billingmeterLegacyUpgradeNormalizeChild(parent billingmeterStateUpgradeAttrMeta, name string, child billingmeterStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if billingmeterLegacyUpgradeIsEmptyValue(value) {
		return billingmeterNullValueForType(child.AttrType)
	}
	if billingmeterLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return billingmeterNullValueForType(child.AttrType)
	}
	if billingmeterLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return billingmeterNullValueForType(child.AttrType)
	}
	return value
}

func billingmeterLegacyUpgradeChildAttr(path []string, parent billingmeterStateUpgradeAttrMeta, name string, child billingmeterStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !billingmeterShouldPreserveChild(parent, child) {
		return billingmeterNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return billingmeterNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := billingmeterUpgradeValue(nextPath, child, childValue)
	return billingmeterLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func billingmeterUpgradeAttrs(path []string, meta map[string]billingmeterStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = billingmeterNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := billingmeterUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = billingmeterLegacyUpgradeNormalizeChild(
				billingmeterStateUpgradeAttrMeta{PreserveConfiguredValue: true},
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

func billingmeterUpgradeObjectValue(path []string, meta billingmeterStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = billingmeterLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func billingmeterUpgradeSingletonListToObject(path []string, meta billingmeterStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return billingmeterNullValueForType(meta.AttrType)
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
		upgradedAttrs[name] = billingmeterLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func billingmeterUpgradeObjectValueToSingletonList(path []string, meta billingmeterStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	if listValue, ok := priorValue.(types.List); ok {
		return billingmeterUpgradeListValue(path, meta, listType, listValue)
	}
	if baseList, ok := priorValue.(basetypes.ListValue); ok {
		return billingmeterUpgradeListValue(path, meta, listType, types.List(baseList))
	}

	objectValue, ok := priorValue.(types.Object)
	if !ok {
		if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
			objectValue = types.Object(baseObject)
		} else {
			return billingmeterNullValueForType(meta.AttrType)
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
		return billingmeterNullValueForType(meta.AttrType)
	}

	upgradedObject := billingmeterUpgradeObjectValue(path, meta, elementObjectType, objectValue)
	return types.ListValueMust(listType.ElemType, []attr.Value{upgradedObject})
}

func billingmeterUpgradeListValue(path []string, meta billingmeterStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return billingmeterNullValueForType(meta.AttrType)
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
				upgradedElements = append(upgradedElements, billingmeterNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			billingmeterUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func billingmeterUpgradeValue(path []string, meta billingmeterStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := billingmeterStateUpgradeSingletonPaths[pathKey]; ok {
			return billingmeterUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return billingmeterNullValueForType(meta.AttrType)
			}
		}
		return billingmeterUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		if _, ok := billingmeterStateUpgradeLegacyObjectPaths[pathKey]; ok {
			return billingmeterUpgradeObjectValueToSingletonList(path, meta, attrType, priorValue)
		}
		return billingmeterUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradeBillingMeterStateV1(ctx context.Context, prior interface{}) (BillingMeterResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := billingmeterUpgradeAttrs(nil, billingmeterStateUpgradeRootMeta, billingmeterAttrMapFromModel(prior))
	var upgraded BillingMeterResourceModel
	billingmeterSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *BillingMeterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     2,
		Description: "Meters specify how to aggregate meter events over a billing period. Meter events represent the actions that customers take in your system. Meters attach to prices and form the basis of the bill.\n\nRelated guide: [Usage based billing](https://docs.stripe.com/billing/subscriptions/usage-based)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("billing.meter")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "The meter's name.",
			},
			"event_name": schema.StringAttribute{
				Required:      true,
				Description:   "The name of the meter event to record usage for. Corresponds with the `event_name` field on meter events.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"event_time_window": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The time window which meter events have been pre-aggregated for, if any.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("day", "hour")},
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
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The meter's status.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "inactive")},
			},
			"status_transitions": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"deactivated_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time the meter was deactivated, if any. Measured in seconds since Unix epoch.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"updated": schema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was last updated. Measured in seconds since the Unix epoch.",
			},
		},
		Blocks: map[string]schema.Block{
			"customer_mapping": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"event_payload_key": schema.StringAttribute{
							Required:      true,
							Description:   "The key in the meter event payload to use for mapping the event to a customer.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"type": schema.StringAttribute{
							Required:      true,
							Description:   "The method for mapping a meter event to a customer.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							Validators:    []validator.String{stringvalidator.OneOf("by_id")},
						},
					},
				},
			},
			"default_aggregation": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"formula": schema.StringAttribute{
							Required:      true,
							Description:   "Specifies how events are aggregated.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							Validators:    []validator.String{stringvalidator.OneOf("count", "last", "sum")},
						},
					},
				},
			},
			"value_settings": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"event_payload_key": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The key in the meter event payload to use as the value for this meter.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
		},
	}
}

func (r *BillingMeterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan BillingMeterResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandBillingMeterCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building BillingMeter create params", err.Error())
		return
	}

	obj, err := r.client.V1BillingMeters.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating BillingMeter", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingMeters.B, r.client.V1BillingMeters.Key, stripe.FormatURLPath("/v1/billing/meters/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingMeter create raw response", err.Error())
		return
	}

	if err := flattenBillingMeter(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingMeter create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingMeterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState BillingMeterResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state BillingMeterResourceModel
	state = priorState

	obj, err := r.client.V1BillingMeters.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading BillingMeter", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingMeters.B, r.client.V1BillingMeters.Key, stripe.FormatURLPath("/v1/billing/meters/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingMeter raw response", err.Error())
		return
	}

	if err := flattenBillingMeter(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingMeter read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingMeterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan BillingMeterResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state BillingMeterResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandBillingMeterUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building BillingMeter update params", err.Error())
		return
	}

	obj, err := r.client.V1BillingMeters.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating BillingMeter", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingMeters.B, r.client.V1BillingMeters.Key, stripe.FormatURLPath("/v1/billing/meters/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingMeter update raw response", err.Error())
		return
	}

	if err := flattenBillingMeter(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingMeter update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingMeterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state BillingMeterResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1BillingMeters.Deactivate(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating BillingMeter", err.Error())
		return
	}
}

func (r *BillingMeterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandBillingMeterCreate(plan BillingMeterResourceModel) (*stripe.BillingMeterCreateParams, error) {
	params := &stripe.BillingMeterCreateParams{}

	if !plan.CustomerMapping.IsNull() && !plan.CustomerMapping.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomerMapping", plan.CustomerMapping) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer_mapping", params)
		}
	}
	if !plan.DefaultAggregation.IsNull() && !plan.DefaultAggregation.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DefaultAggregation", plan.DefaultAggregation) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "default_aggregation", params)
		}
	}
	if !plan.DisplayName.IsNull() && !plan.DisplayName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayName", "DisplayName", plan.DisplayName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "display_name", params)
		}
	}
	if !plan.EventName.IsNull() && !plan.EventName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "EventName", "EventName", plan.EventName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "event_name", params)
		}
	}
	if !plan.EventTimeWindow.IsNull() && !plan.EventTimeWindow.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "EventTimeWindow", "EventTimeWindow", plan.EventTimeWindow.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "event_time_window", params)
		}
	}
	if !plan.ValueSettings.IsNull() && !plan.ValueSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ValueSettings", plan.ValueSettings) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "value_settings", params)
		}
	}

	return params, nil
}

func expandBillingMeterUpdate(plan BillingMeterResourceModel, state BillingMeterResourceModel) (*stripe.BillingMeterUpdateParams, error) {
	params := &stripe.BillingMeterUpdateParams{}

	if !plan.DisplayName.Equal(state.DisplayName) && !plan.DisplayName.IsNull() && !plan.DisplayName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayName", "DisplayName", plan.DisplayName.ValueString()) {
			if !plan.DisplayName.Equal(state.DisplayName) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "display_name", params)
			}
		}
	}

	return params, nil
}

func flattenBillingMeter(obj *stripe.BillingMeter, state *BillingMeterResourceModel) error {
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
		assignedCustomerMapping := false
		hadRawCustomerMapping := false
		if rawValueCustomerMapping, rawOk := plainValueAtPath(raw, "customer_mapping"); rawOk {
			hadRawCustomerMapping = true
			if rawValueCustomerMapping != nil {
				sourceCustomerMapping := applyConfiguredKeyedListShapes(rawValueCustomerMapping, unwrapPlainSingletonList(attrValueToPlain(state.CustomerMapping)))
				if valueCustomerMapping, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceCustomerMapping, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"event_payload_key": types.StringType, "type": types.StringType}}}, "customer_mapping", "raw response"); err != nil {
					return err
				} else {
					if typedCustomerMapping, ok := valueCustomerMapping.(types.List); ok {
						state.CustomerMapping = typedCustomerMapping
						assignedCustomerMapping = true
					}
				}
			}
		}
		if !assignedCustomerMapping {
			if !hasRaw {
				if responseValueCustomerMapping, ok := plainFromResponseField(obj, "CustomerMapping"); ok {
					sourceCustomerMapping := applyConfiguredKeyedListShapes(responseValueCustomerMapping, unwrapPlainSingletonList(attrValueToPlain(state.CustomerMapping)))
					if valueCustomerMapping, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceCustomerMapping, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"event_payload_key": types.StringType, "type": types.StringType}}},
						"customer_mapping",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCustomerMapping, ok := valueCustomerMapping.(types.List); ok {
							state.CustomerMapping = typedCustomerMapping
							assignedCustomerMapping = true
						}
					}
				}
			}
		}
		if !assignedCustomerMapping && hadRawCustomerMapping {
			if nullCustomerMapping, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"event_payload_key": types.StringType, "type": types.StringType}}}); ok {
				if typedCustomerMapping, ok := nullCustomerMapping.(types.List); ok {
					state.CustomerMapping = typedCustomerMapping
				}
			}
		}
	}
	{
		assignedDefaultAggregation := false
		hadRawDefaultAggregation := false
		if rawValueDefaultAggregation, rawOk := plainValueAtPath(raw, "default_aggregation"); rawOk {
			hadRawDefaultAggregation = true
			if rawValueDefaultAggregation != nil {
				sourceDefaultAggregation := applyConfiguredKeyedListShapes(rawValueDefaultAggregation, unwrapPlainSingletonList(attrValueToPlain(state.DefaultAggregation)))
				if valueDefaultAggregation, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceDefaultAggregation, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"formula": types.StringType}}}, "default_aggregation", "raw response"); err != nil {
					return err
				} else {
					if typedDefaultAggregation, ok := valueDefaultAggregation.(types.List); ok {
						state.DefaultAggregation = typedDefaultAggregation
						assignedDefaultAggregation = true
					}
				}
			}
		}
		if !assignedDefaultAggregation {
			if !hasRaw {
				if responseValueDefaultAggregation, ok := plainFromResponseField(obj, "DefaultAggregation"); ok {
					sourceDefaultAggregation := applyConfiguredKeyedListShapes(responseValueDefaultAggregation, unwrapPlainSingletonList(attrValueToPlain(state.DefaultAggregation)))
					if valueDefaultAggregation, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceDefaultAggregation, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"formula": types.StringType}}},
						"default_aggregation",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedDefaultAggregation, ok := valueDefaultAggregation.(types.List); ok {
							state.DefaultAggregation = typedDefaultAggregation
							assignedDefaultAggregation = true
						}
					}
				}
			}
		}
		if !assignedDefaultAggregation && hadRawDefaultAggregation {
			if nullDefaultAggregation, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"formula": types.StringType}}}); ok {
				if typedDefaultAggregation, ok := nullDefaultAggregation.(types.List); ok {
					state.DefaultAggregation = typedDefaultAggregation
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
		if rawValueEventName, rawOk := plainValueAtPath(raw, "event_name"); rawOk {
			if valueEventName, err := flattenPlainValue(rawValueEventName, types.StringType, "event_name", "raw response"); err != nil {
				return err
			} else {
				if typedEventName, ok := valueEventName.(types.String); ok {
					state.EventName = typedEventName
				}
			}
		} else if !hasRaw {
			if responseValueEventName, ok := plainFromResponseField(obj, "EventName"); ok {
				if valueEventName, err := flattenPlainValue(responseValueEventName, types.StringType, "event_name", "response struct"); err != nil {
					return err
				} else {
					if typedEventName, ok := valueEventName.(types.String); ok {
						state.EventName = typedEventName
					}
				}
			}
		}
	}
	{
		if rawValueEventTimeWindow, rawOk := plainValueAtPath(raw, "event_time_window"); rawOk {
			if valueEventTimeWindow, err := flattenPlainValue(rawValueEventTimeWindow, types.StringType, "event_time_window", "raw response"); err != nil {
				return err
			} else {
				if typedEventTimeWindow, ok := valueEventTimeWindow.(types.String); ok {
					state.EventTimeWindow = typedEventTimeWindow
				}
			}
		} else if !hasRaw {
			if responseValueEventTimeWindow, ok := plainFromResponseField(obj, "EventTimeWindow"); ok {
				if valueEventTimeWindow, err := flattenPlainValue(responseValueEventTimeWindow, types.StringType, "event_time_window", "response struct"); err != nil {
					return err
				} else {
					if typedEventTimeWindow, ok := valueEventTimeWindow.(types.String); ok {
						state.EventTimeWindow = typedEventTimeWindow
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
		assignedStatusTransitions := false
		hadRawStatusTransitions := false
		if rawValueStatusTransitions, rawOk := plainValueAtPath(raw, "status_transitions"); rawOk {
			hadRawStatusTransitions = true
			if rawValueStatusTransitions != nil {
				sourceStatusTransitions := applyConfiguredKeyedListShapes(rawValueStatusTransitions, attrValueToPlain(state.StatusTransitions))
				if valueStatusTransitions, err := flattenPlainValue(sourceStatusTransitions, types.ObjectType{AttrTypes: map[string]attr.Type{"deactivated_at": types.Int64Type}}, "status_transitions", "raw response"); err != nil {
					return err
				} else {
					if typedStatusTransitions, ok := valueStatusTransitions.(types.Object); ok {
						state.StatusTransitions = typedStatusTransitions
						assignedStatusTransitions = true
					}
				}
			}
		}
		if !assignedStatusTransitions {
			if !hasRaw {
				if responseValueStatusTransitions, ok := plainFromResponseField(obj, "StatusTransitions"); ok {
					sourceStatusTransitions := applyConfiguredKeyedListShapes(responseValueStatusTransitions, attrValueToPlain(state.StatusTransitions))
					if valueStatusTransitions, err := flattenPlainValue(
						sourceStatusTransitions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"deactivated_at": types.Int64Type}},
						"status_transitions",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedStatusTransitions, ok := valueStatusTransitions.(types.Object); ok {
							state.StatusTransitions = typedStatusTransitions
							assignedStatusTransitions = true
						}
					}
				}
			}
		}
		if !assignedStatusTransitions && hadRawStatusTransitions {
			if nullStatusTransitions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"deactivated_at": types.Int64Type}}); ok {
				if typedStatusTransitions, ok := nullStatusTransitions.(types.Object); ok {
					state.StatusTransitions = typedStatusTransitions
				}
			}
		}
	}
	{
		if rawValueUpdated, rawOk := plainValueAtPath(raw, "updated"); rawOk {
			if valueUpdated, err := flattenPlainValue(rawValueUpdated, types.Int64Type, "updated", "raw response"); err != nil {
				return err
			} else {
				if typedUpdated, ok := valueUpdated.(types.Int64); ok {
					state.Updated = typedUpdated
				}
			}
		} else if !hasRaw {
			if responseValueUpdated, ok := plainFromResponseField(obj, "Updated"); ok {
				if valueUpdated, err := flattenPlainValue(responseValueUpdated, types.Int64Type, "updated", "response struct"); err != nil {
					return err
				} else {
					if typedUpdated, ok := valueUpdated.(types.Int64); ok {
						state.Updated = typedUpdated
					}
				}
			}
		}
	}
	{
		assignedValueSettings := false
		hadRawValueSettings := false
		if rawValueValueSettings, rawOk := plainValueAtPath(raw, "value_settings"); rawOk {
			hadRawValueSettings = true
			if rawValueValueSettings != nil {
				sourceValueSettings := applyConfiguredKeyedListShapes(rawValueValueSettings, unwrapPlainSingletonList(attrValueToPlain(state.ValueSettings)))
				if valueValueSettings, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceValueSettings, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"event_payload_key": types.StringType}}}, "value_settings", "raw response"); err != nil {
					return err
				} else {
					if typedValueSettings, ok := valueValueSettings.(types.List); ok {
						state.ValueSettings = typedValueSettings
						assignedValueSettings = true
					}
				}
			}
		}
		if !assignedValueSettings {
			if !hasRaw {
				if responseValueValueSettings, ok := plainFromResponseField(obj, "ValueSettings"); ok {
					sourceValueSettings := applyConfiguredKeyedListShapes(responseValueValueSettings, unwrapPlainSingletonList(attrValueToPlain(state.ValueSettings)))
					if valueValueSettings, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceValueSettings, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"event_payload_key": types.StringType}}},
						"value_settings",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedValueSettings, ok := valueValueSettings.(types.List); ok {
							state.ValueSettings = typedValueSettings
							assignedValueSettings = true
						}
					}
				}
			}
		}
		if !assignedValueSettings && hadRawValueSettings {
			if nullValueSettings, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"event_payload_key": types.StringType}}}); ok {
				if typedValueSettings, ok := nullValueSettings.(types.List); ok {
					state.ValueSettings = typedValueSettings
				}
			}
		}
	}
	return nil
}
