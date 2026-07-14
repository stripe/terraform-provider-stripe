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

var _ resource.Resource = &V2CoreEventDestinationResource{}

var _ resource.ResourceWithConfigure = &V2CoreEventDestinationResource{}

var _ resource.ResourceWithImportState = &V2CoreEventDestinationResource{}

func NewV2CoreEventDestinationResource() resource.Resource {
	return &V2CoreEventDestinationResource{}
}

type V2CoreEventDestinationResource struct {
	client *stripe.Client
}

type V2CoreEventDestinationResourceModel struct {
	Object             types.String `tfsdk:"object"`
	AmazonEventbridge  types.List   `tfsdk:"amazon_eventbridge"`
	AzureEventGrid     types.Object `tfsdk:"azure_event_grid"`
	Created            types.String `tfsdk:"created"`
	Description        types.String `tfsdk:"description"`
	EnabledEvents      types.List   `tfsdk:"enabled_events"`
	EventPayload       types.String `tfsdk:"event_payload"`
	EventsFrom         types.List   `tfsdk:"events_from"`
	ID                 types.String `tfsdk:"id"`
	Livemode           types.Bool   `tfsdk:"livemode"`
	Metadata           types.Map    `tfsdk:"metadata"`
	Name               types.String `tfsdk:"name"`
	SnapshotAPIVersion types.String `tfsdk:"snapshot_api_version"`
	Status             types.String `tfsdk:"status"`
	StatusDetails      types.Object `tfsdk:"status_details"`
	Type               types.String `tfsdk:"type"`
	Updated            types.String `tfsdk:"updated"`
	WebhookEndpoint    types.List   `tfsdk:"webhook_endpoint"`
	Include            types.List   `tfsdk:"include"`
}

func (r *V2CoreEventDestinationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *V2CoreEventDestinationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v2_core_event_destination"
}

var _ resource.ResourceWithUpgradeState = &V2CoreEventDestinationResource{}

func (r *V2CoreEventDestinationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		1: {
			PriorSchema: v2CoreEventDestinationResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior V2CoreEventDestinationResourceV0Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeV2CoreEventDestinationStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func v2CoreEventDestinationResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "Set up an event destination to receive events from Stripe across multiple destination types, including [webhook endpoints](https://docs.stripe.com/webhooks) and [Amazon EventBridge](https://docs.stripe.com/event-destinations/eventbridge). Event destinations support receiving [thin events](https://docs.stripe.com/api/v2/events) and [snapshot events](https://docs.stripe.com/api/events).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value of the object field.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("v2.core.event_destination")},
			},
			"amazon_eventbridge": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Amazon EventBridge configuration.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"aws_account_id": schema.StringAttribute{
						Required:      true,
						Description:   "The AWS account ID.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"aws_event_source_arn": schema.StringAttribute{
						Computed:      true,
						Description:   "The ARN of the AWS event source.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"aws_event_source_status": schema.StringAttribute{
						Computed:      true,
						Description:   "The state of the AWS event source.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("active", "deleted", "pending", "unknown")},
					},
					"aws_region": schema.StringAttribute{
						Required:      true,
						Description:   "The region of the AWS event source.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"azure_event_grid": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Azure Event Grid configuration.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"azure_partner_topic_name": schema.StringAttribute{
						Computed:      true,
						Description:   "The name of the Azure partner topic.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"azure_partner_topic_status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the Azure partner topic.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("activated", "deleted", "never_activated", "unknown")},
					},
					"azure_region": schema.StringAttribute{
						Required:      true,
						Description:   "The Azure region.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"azure_resource_group_name": schema.StringAttribute{
						Required:      true,
						Description:   "The name of the Azure resource group.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"azure_subscription_id": schema.StringAttribute{
						Required:      true,
						Description:   "The Azure subscription ID.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				Description:   "Time at which the object was created.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An optional description of what the event destination is used for.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled_events": schema.ListAttribute{
				Required:    true,
				Description: "The list of events to enable for this endpoint.",
				ElementType: types.StringType,
			},
			"event_payload": schema.StringAttribute{
				Required:      true,
				Description:   "Payload type of events being subscribed to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("snapshot", "thin")},
			},
			"events_from": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Specifies which accounts' events route to this destination.\n`@self`: Receive events from the account that owns the event destination.\n`@accounts`: Receive events emitted from other accounts you manage which includes your v1 and v2 accounts.\n`@organization_members`: Receive events from accounts directly linked to the organization.\n`@organization_members/@accounts`: Receive events from all accounts connected to any platform accounts in the organization.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
				ElementType:   types.StringType,
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Metadata.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Event destination name.",
			},
			"snapshot_api_version": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "If using the snapshot event payload, the API version events are rendered as.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Status. It can be set to either enabled or disabled.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("disabled", "enabled")},
			},
			"status_details": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Additional information about event destination status.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"disabled": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details about why the event destination has been disabled.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"reason": schema.StringAttribute{
								Computed:      true,
								Description:   "Reason event destination has been disabled.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("no_aws_event_source_exists", "no_azure_partner_topic_exists", "user")},
							},
						},
					},
				},
			},
			"type": schema.StringAttribute{
				Required:      true,
				Description:   "Event destination type.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("amazon_eventbridge", "azure_event_grid", "webhook_endpoint")},
			},
			"updated": schema.StringAttribute{
				Computed:    true,
				Description: "Time at which the object was last updated.",
			},
			"webhook_endpoint": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Webhook endpoint configuration.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"signing_secret": schema.StringAttribute{
						Computed:      true,
						Description:   "The signing secret of the webhook endpoint, only includable on creation.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"url": schema.StringAttribute{
						Required:    true,
						Description: "The URL of the webhook endpoint, includable.",
					},
				},
			},
			"include": schema.ListAttribute{
				Optional:    true,
				Description: "Additional fields to include in the response.",
				WriteOnly:   true,
				ElementType: types.StringType,
			},
		},
	}
}

type V2CoreEventDestinationResourceV0Model struct {
	Object             types.String `tfsdk:"object"`
	AmazonEventbridge  types.Object `tfsdk:"amazon_eventbridge"`
	AzureEventGrid     types.Object `tfsdk:"azure_event_grid"`
	Created            types.String `tfsdk:"created"`
	Description        types.String `tfsdk:"description"`
	EnabledEvents      types.List   `tfsdk:"enabled_events"`
	EventPayload       types.String `tfsdk:"event_payload"`
	EventsFrom         types.List   `tfsdk:"events_from"`
	ID                 types.String `tfsdk:"id"`
	Livemode           types.Bool   `tfsdk:"livemode"`
	Metadata           types.Map    `tfsdk:"metadata"`
	Name               types.String `tfsdk:"name"`
	SnapshotAPIVersion types.String `tfsdk:"snapshot_api_version"`
	Status             types.String `tfsdk:"status"`
	StatusDetails      types.Object `tfsdk:"status_details"`
	Type               types.String `tfsdk:"type"`
	Updated            types.String `tfsdk:"updated"`
	WebhookEndpoint    types.Object `tfsdk:"webhook_endpoint"`
	Include            types.List   `tfsdk:"include"`
}

type v2coreeventdestinationStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]v2coreeventdestinationStateUpgradeAttrMeta
}

var v2coreeventdestinationStateUpgradeRootMeta = map[string]v2coreeventdestinationStateUpgradeAttrMeta{"object": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "amazon_eventbridge": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"aws_account_id": types.StringType, "aws_event_source_arn": types.StringType, "aws_event_source_status": types.StringType, "aws_region": types.StringType}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", PreserveConfiguredValue: true, Nested: map[string]v2coreeventdestinationStateUpgradeAttrMeta{"aws_account_id": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "aws_event_source_arn": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "aws_event_source_status": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "aws_region": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "azure_event_grid": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{"azure_partner_topic_name": types.StringType, "azure_partner_topic_status": types.StringType, "azure_region": types.StringType, "azure_resource_group_name": types.StringType, "azure_subscription_id": types.StringType}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]v2coreeventdestinationStateUpgradeAttrMeta{"azure_partner_topic_name": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "azure_partner_topic_status": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "azure_region": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "azure_resource_group_name": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "azure_subscription_id": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "created": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "description": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "enabled_events": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.StringType}, Behavior: "required", LegacyBehavior: "required"}, "event_payload": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "events_from": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "id": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "livemode": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "metadata": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "name": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "snapshot_api_version": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "status": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "status_details": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{"disabled": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}}}, Behavior: "computed", LegacyBehavior: "computed", Nested: map[string]v2coreeventdestinationStateUpgradeAttrMeta{"disabled": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}, Behavior: "computed", LegacyBehavior: "computed", Nested: map[string]v2coreeventdestinationStateUpgradeAttrMeta{"reason": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}}}}}, "type": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "updated": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "webhook_endpoint": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"signing_secret": types.StringType, "url": types.StringType}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", PreserveConfiguredValue: true, Nested: map[string]v2coreeventdestinationStateUpgradeAttrMeta{"signing_secret": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "url": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "include": v2coreeventdestinationStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.StringType}, Behavior: "optional", LegacyBehavior: "optional"}}

var v2coreeventdestinationStateUpgradeSingletonPaths = map[string]struct{}{}

var v2coreeventdestinationStateUpgradeLegacyObjectPaths = map[string]struct{}{"amazon_eventbridge": struct{}{}, "webhook_endpoint": struct{}{}}

func v2coreeventdestinationAttrMapFromModel(model interface{}) map[string]attr.Value {
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

func v2coreeventdestinationSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
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

func v2coreeventdestinationIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func v2coreeventdestinationShouldPreserveChild(parent v2coreeventdestinationStateUpgradeAttrMeta, child v2coreeventdestinationStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !v2coreeventdestinationIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !v2coreeventdestinationIsComputedBehavior(child.Behavior)
}

func v2coreeventdestinationNullValueForType(attributeType attr.Type) attr.Value {
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

func v2coreeventdestinationLegacyUpgradeIsEmptyValue(value attr.Value) bool {
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

func v2coreeventdestinationLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
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

func v2coreeventdestinationLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := v2coreeventdestinationLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func v2coreeventdestinationLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
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

func v2coreeventdestinationLegacyUpgradeNormalizeChild(parent v2coreeventdestinationStateUpgradeAttrMeta, name string, child v2coreeventdestinationStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if v2coreeventdestinationLegacyUpgradeIsEmptyValue(value) {
		return v2coreeventdestinationNullValueForType(child.AttrType)
	}
	if v2coreeventdestinationLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return v2coreeventdestinationNullValueForType(child.AttrType)
	}
	if v2coreeventdestinationLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return v2coreeventdestinationNullValueForType(child.AttrType)
	}
	return value
}

func v2coreeventdestinationLegacyUpgradeChildAttr(path []string, parent v2coreeventdestinationStateUpgradeAttrMeta, name string, child v2coreeventdestinationStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !v2coreeventdestinationShouldPreserveChild(parent, child) {
		return v2coreeventdestinationNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return v2coreeventdestinationNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := v2coreeventdestinationUpgradeValue(nextPath, child, childValue)
	return v2coreeventdestinationLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func v2coreeventdestinationUpgradeAttrs(path []string, meta map[string]v2coreeventdestinationStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = v2coreeventdestinationNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := v2coreeventdestinationUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = v2coreeventdestinationLegacyUpgradeNormalizeChild(
				v2coreeventdestinationStateUpgradeAttrMeta{PreserveConfiguredValue: true},
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

func v2coreeventdestinationUpgradeObjectValue(path []string, meta v2coreeventdestinationStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = v2coreeventdestinationLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func v2coreeventdestinationUpgradeSingletonListToObject(path []string, meta v2coreeventdestinationStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return v2coreeventdestinationNullValueForType(meta.AttrType)
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
		upgradedAttrs[name] = v2coreeventdestinationLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func v2coreeventdestinationUpgradeObjectValueToSingletonList(path []string, meta v2coreeventdestinationStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	objectValue, ok := priorValue.(types.Object)
	if !ok {
		if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
			objectValue = types.Object(baseObject)
		} else {
			return v2coreeventdestinationNullValueForType(meta.AttrType)
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
		return v2coreeventdestinationNullValueForType(meta.AttrType)
	}

	upgradedObject := v2coreeventdestinationUpgradeObjectValue(path, meta, elementObjectType, objectValue)
	return types.ListValueMust(listType.ElemType, []attr.Value{upgradedObject})
}

func v2coreeventdestinationUpgradeListValue(path []string, meta v2coreeventdestinationStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return v2coreeventdestinationNullValueForType(meta.AttrType)
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
				upgradedElements = append(upgradedElements, v2coreeventdestinationNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			v2coreeventdestinationUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func v2coreeventdestinationUpgradeValue(path []string, meta v2coreeventdestinationStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := v2coreeventdestinationStateUpgradeSingletonPaths[pathKey]; ok {
			return v2coreeventdestinationUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return v2coreeventdestinationNullValueForType(meta.AttrType)
			}
		}
		return v2coreeventdestinationUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		if _, ok := v2coreeventdestinationStateUpgradeLegacyObjectPaths[pathKey]; ok {
			return v2coreeventdestinationUpgradeObjectValueToSingletonList(path, meta, attrType, priorValue)
		}
		return v2coreeventdestinationUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradeV2CoreEventDestinationStateV1(ctx context.Context, prior V2CoreEventDestinationResourceV0Model) (V2CoreEventDestinationResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := v2coreeventdestinationUpgradeAttrs(nil, v2coreeventdestinationStateUpgradeRootMeta, v2coreeventdestinationAttrMapFromModel(prior))
	var upgraded V2CoreEventDestinationResourceModel
	v2coreeventdestinationSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *V2CoreEventDestinationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     2,
		Description: "Set up an event destination to receive events from Stripe across multiple destination types, including [webhook endpoints](https://docs.stripe.com/webhooks) and [Amazon EventBridge](https://docs.stripe.com/event-destinations/eventbridge). Event destinations support receiving [thin events](https://docs.stripe.com/api/v2/events) and [snapshot events](https://docs.stripe.com/api/events).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value of the object field.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("v2.core.event_destination")},
			},
			"azure_event_grid": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Azure Event Grid configuration.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"azure_partner_topic_name": schema.StringAttribute{
						Computed:      true,
						Description:   "The name of the Azure partner topic.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"azure_partner_topic_status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the Azure partner topic.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("activated", "deleted", "never_activated", "unknown")},
					},
					"azure_region": schema.StringAttribute{
						Required:      true,
						Description:   "The Azure region.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"azure_resource_group_name": schema.StringAttribute{
						Required:      true,
						Description:   "The name of the Azure resource group.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"azure_subscription_id": schema.StringAttribute{
						Required:      true,
						Description:   "The Azure subscription ID.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				Description:   "Time at which the object was created.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An optional description of what the event destination is used for.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled_events": schema.ListAttribute{
				Required:    true,
				Description: "The list of events to enable for this endpoint.",
				ElementType: types.StringType,
			},
			"event_payload": schema.StringAttribute{
				Required:      true,
				Description:   "Payload type of events being subscribed to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("snapshot", "thin")},
			},
			"events_from": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Specifies which accounts' events route to this destination.\n`@self`: Receive events from the account that owns the event destination.\n`@accounts`: Receive events emitted from other accounts you manage which includes your v1 and v2 accounts.\n`@organization_members`: Receive events from accounts directly linked to the organization.\n`@organization_members/@accounts`: Receive events from all accounts connected to any platform accounts in the organization.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
				ElementType:   types.StringType,
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Metadata.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Event destination name.",
			},
			"snapshot_api_version": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "If using the snapshot event payload, the API version events are rendered as.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Status. It can be set to either enabled or disabled.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("disabled", "enabled")},
			},
			"status_details": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Additional information about event destination status.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"disabled": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details about why the event destination has been disabled.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"reason": schema.StringAttribute{
								Computed:      true,
								Description:   "Reason event destination has been disabled.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("no_aws_event_source_exists", "no_azure_partner_topic_exists", "user")},
							},
						},
					},
				},
			},
			"type": schema.StringAttribute{
				Required:      true,
				Description:   "Event destination type.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("amazon_eventbridge", "azure_event_grid", "webhook_endpoint")},
			},
			"updated": schema.StringAttribute{
				Computed:    true,
				Description: "Time at which the object was last updated.",
			},
			"include": schema.ListAttribute{
				Optional:    true,
				Description: "Additional fields to include in the response.",
				WriteOnly:   true,
				ElementType: types.StringType,
			},
		},
		Blocks: map[string]schema.Block{
			"amazon_eventbridge": schema.ListNestedBlock{
				Description:   "Amazon EventBridge configuration.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"aws_account_id": schema.StringAttribute{
							Required:      true,
							Description:   "The AWS account ID.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"aws_event_source_arn": schema.StringAttribute{
							Computed:      true,
							Description:   "The ARN of the AWS event source.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"aws_event_source_status": schema.StringAttribute{
							Computed:      true,
							Description:   "The state of the AWS event source.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("active", "deleted", "pending", "unknown")},
						},
						"aws_region": schema.StringAttribute{
							Required:      true,
							Description:   "The region of the AWS event source.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
			"webhook_endpoint": schema.ListNestedBlock{
				Description:   "Webhook endpoint configuration.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"signing_secret": schema.StringAttribute{
							Computed:      true,
							Description:   "The signing secret of the webhook endpoint, only includable on creation.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"url": schema.StringAttribute{
							Required:    true,
							Description: "The URL of the webhook endpoint, includable.",
						},
					},
				},
			},
		},
	}
}

func (r *V2CoreEventDestinationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan V2CoreEventDestinationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config V2CoreEventDestinationResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandV2CoreEventDestinationCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2CoreEventDestination create params", err.Error())
		return
	}

	obj, err := r.client.V2CoreEventDestinations.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating V2CoreEventDestination", err.Error())
		return
	}

	rawReadParams := &stripe.V2CoreEventDestinationRetrieveParams{}
	if !plan.Include.IsNull() && !plan.Include.IsUnknown() {
		if !assignAttrValueToNamedField(rawReadParams, "Include", plan.Include) {
			resp.Diagnostics.AddError("Error building V2CoreEventDestination read params", fmt.Sprintf("failed to assign attribute %q on %T", "include", rawReadParams))
			return
		}
	}

	if err := ensureRawResponse(obj, r.client.V2CoreEventDestinations.B, r.client.V2CoreEventDestinations.Key, stripe.FormatURLPath("/v2/core/event_destinations/%s", obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2CoreEventDestination create raw response", err.Error())
		return
	}

	if err := flattenV2CoreEventDestination(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2CoreEventDestination create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"AmazonEventbridge", "aws_region"}, []string{"Include"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2CoreEventDestinationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState V2CoreEventDestinationResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2CoreEventDestinationResourceModel
	state = priorState

	params := &stripe.V2CoreEventDestinationRetrieveParams{}
	if !state.Include.IsNull() && !state.Include.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Include", state.Include) {
			resp.Diagnostics.AddError("Error building V2CoreEventDestination read params", fmt.Sprintf("failed to assign attribute %q on %T", "include", params))
			return
		}
	}

	obj, err := r.client.V2CoreEventDestinations.Retrieve(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error reading V2CoreEventDestination", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2CoreEventDestinations.B, r.client.V2CoreEventDestinations.Key, stripe.FormatURLPath("/v2/core/event_destinations/%s", state.ID.ValueString()), params); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2CoreEventDestination raw response", err.Error())
		return
	}

	if err := flattenV2CoreEventDestination(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening V2CoreEventDestination read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"AmazonEventbridge", "aws_region"}, []string{"Include"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *V2CoreEventDestinationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan V2CoreEventDestinationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config V2CoreEventDestinationResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2CoreEventDestinationResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandV2CoreEventDestinationUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2CoreEventDestination update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building V2CoreEventDestination update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V2CoreEventDestinations.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating V2CoreEventDestination", err.Error())
		return
	}

	rawReadParams := &stripe.V2CoreEventDestinationRetrieveParams{}
	if !state.Include.IsNull() && !state.Include.IsUnknown() {
		if !assignAttrValueToNamedField(rawReadParams, "Include", state.Include) {
			resp.Diagnostics.AddError("Error building V2CoreEventDestination read params", fmt.Sprintf("failed to assign attribute %q on %T", "include", rawReadParams))
			return
		}
	}

	if err := ensureRawResponse(obj, r.client.V2CoreEventDestinations.B, r.client.V2CoreEventDestinations.Key, stripe.FormatURLPath("/v2/core/event_destinations/%s", obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2CoreEventDestination update raw response", err.Error())
		return
	}

	if err := flattenV2CoreEventDestination(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2CoreEventDestination update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"AmazonEventbridge", "aws_region"}, []string{"Include"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2CoreEventDestinationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state V2CoreEventDestinationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V2CoreEventDestinations.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting V2CoreEventDestination", err.Error())
		return
	}
}

func (r *V2CoreEventDestinationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandV2CoreEventDestinationCreate(plan V2CoreEventDestinationResourceModel) (*stripe.V2CoreEventDestinationCreateParams, error) {
	params := &stripe.V2CoreEventDestinationCreateParams{}

	if !plan.AmazonEventbridge.IsNull() && !plan.AmazonEventbridge.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AmazonEventbridge", plan.AmazonEventbridge) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "amazon_eventbridge", params)
		}
	}
	if !plan.AzureEventGrid.IsNull() && !plan.AzureEventGrid.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AzureEventGrid", plan.AzureEventGrid) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "azure_event_grid", params)
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
	if !plan.EventPayload.IsNull() && !plan.EventPayload.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "EventPayload", "EventPayload", plan.EventPayload.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "event_payload", params)
		}
	}
	if !plan.EventsFrom.IsNull() && !plan.EventsFrom.IsUnknown() {
		if !assignAttrValueToNamedField(params, "EventsFrom", plan.EventsFrom) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "events_from", params)
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
	if !plan.SnapshotAPIVersion.IsNull() && !plan.SnapshotAPIVersion.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SnapshotAPIVersion", "SnapshotAPIVersion", plan.SnapshotAPIVersion.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "snapshot_api_version", params)
		}
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}
	if !plan.WebhookEndpoint.IsNull() && !plan.WebhookEndpoint.IsUnknown() {
		if !assignAttrValueToNamedField(params, "WebhookEndpoint", plan.WebhookEndpoint) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "webhook_endpoint", params)
		}
	}
	if !plan.Include.IsNull() && !plan.Include.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Include", plan.Include) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "include", params)
		}
	}

	return params, nil
}

func expandV2CoreEventDestinationUpdate(plan V2CoreEventDestinationResourceModel, state V2CoreEventDestinationResourceModel) (*stripe.V2CoreEventDestinationUpdateParams, error) {
	params := &stripe.V2CoreEventDestinationUpdateParams{}

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
	if !plan.Name.Equal(state.Name) && !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Name", "Name", plan.Name.ValueString()) {
			if !plan.Name.Equal(state.Name) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "name", params)
			}
		}
	}
	if !plan.WebhookEndpoint.Equal(state.WebhookEndpoint) && !plan.WebhookEndpoint.IsNull() && !plan.WebhookEndpoint.IsUnknown() {
		if !assignAttrValueToNamedField(params, "WebhookEndpoint", plan.WebhookEndpoint) {
			if !plan.WebhookEndpoint.Equal(state.WebhookEndpoint) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "webhook_endpoint", params)
			}
		}
	}
	if !plan.Include.Equal(state.Include) && !plan.Include.IsNull() && !plan.Include.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Include", plan.Include) {
			if !plan.Include.Equal(state.Include) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "include", params)
			}
		}
	}

	return params, nil
}

func flattenV2CoreEventDestination(obj *stripe.V2CoreEventDestination, state *V2CoreEventDestinationResourceModel) error {
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
		assignedAmazonEventbridge := false
		if rawValueAmazonEventbridge, rawOk := plainValueAtPath(raw, "amazon_eventbridge"); rawOk {
			if rawValueAmazonEventbridge != nil {
				sourceAmazonEventbridge := mergeMissingPlainLeaves(applyConfiguredKeyedListShapes(rawValueAmazonEventbridge, unwrapPlainSingletonList(attrValueToPlain(state.AmazonEventbridge))), unwrapPlainSingletonList(attrValueToPlain(state.AmazonEventbridge)))
				if valueAmazonEventbridge, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceAmazonEventbridge, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"aws_account_id": types.StringType, "aws_event_source_arn": types.StringType, "aws_event_source_status": types.StringType, "aws_region": types.StringType}}}, "amazon_eventbridge", "raw response"); err != nil {
					return err
				} else {
					if typedAmazonEventbridge, ok := valueAmazonEventbridge.(types.List); ok {
						state.AmazonEventbridge = typedAmazonEventbridge
						assignedAmazonEventbridge = true
					}
				}
			}
		}
		if !assignedAmazonEventbridge {
			if !hasRaw {
				if responseValueAmazonEventbridge, ok := plainFromResponseField(obj, "AmazonEventbridge"); ok {
					sourceAmazonEventbridge := mergeMissingPlainLeaves(applyConfiguredKeyedListShapes(responseValueAmazonEventbridge, unwrapPlainSingletonList(attrValueToPlain(state.AmazonEventbridge))), unwrapPlainSingletonList(attrValueToPlain(state.AmazonEventbridge)))
					if valueAmazonEventbridge, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceAmazonEventbridge, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"aws_account_id": types.StringType, "aws_event_source_arn": types.StringType, "aws_event_source_status": types.StringType, "aws_region": types.StringType}}},
						"amazon_eventbridge",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAmazonEventbridge, ok := valueAmazonEventbridge.(types.List); ok {
							state.AmazonEventbridge = typedAmazonEventbridge
							assignedAmazonEventbridge = true
						}
					}
				}
			}
		}
	}
	{
		assignedAzureEventGrid := false
		hadRawAzureEventGrid := false
		if rawValueAzureEventGrid, rawOk := plainValueAtPath(raw, "azure_event_grid"); rawOk {
			hadRawAzureEventGrid = true
			if rawValueAzureEventGrid != nil {
				sourceAzureEventGrid := applyConfiguredKeyedListShapes(rawValueAzureEventGrid, attrValueToPlain(state.AzureEventGrid))
				if valueAzureEventGrid, err := flattenPlainValue(sourceAzureEventGrid, types.ObjectType{AttrTypes: map[string]attr.Type{"azure_partner_topic_name": types.StringType, "azure_partner_topic_status": types.StringType, "azure_region": types.StringType, "azure_resource_group_name": types.StringType, "azure_subscription_id": types.StringType}}, "azure_event_grid", "raw response"); err != nil {
					return err
				} else {
					if typedAzureEventGrid, ok := valueAzureEventGrid.(types.Object); ok {
						state.AzureEventGrid = typedAzureEventGrid
						assignedAzureEventGrid = true
					}
				}
			}
		}
		if !assignedAzureEventGrid {
			if !hasRaw {
				if responseValueAzureEventGrid, ok := plainFromResponseField(obj, "AzureEventGrid"); ok {
					sourceAzureEventGrid := applyConfiguredKeyedListShapes(responseValueAzureEventGrid, attrValueToPlain(state.AzureEventGrid))
					if valueAzureEventGrid, err := flattenPlainValue(
						sourceAzureEventGrid,
						types.ObjectType{AttrTypes: map[string]attr.Type{"azure_partner_topic_name": types.StringType, "azure_partner_topic_status": types.StringType, "azure_region": types.StringType, "azure_resource_group_name": types.StringType, "azure_subscription_id": types.StringType}},
						"azure_event_grid",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAzureEventGrid, ok := valueAzureEventGrid.(types.Object); ok {
							state.AzureEventGrid = typedAzureEventGrid
							assignedAzureEventGrid = true
						}
					}
				}
			}
		}
		if !assignedAzureEventGrid && hadRawAzureEventGrid {
			if nullAzureEventGrid, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"azure_partner_topic_name": types.StringType, "azure_partner_topic_status": types.StringType, "azure_region": types.StringType, "azure_resource_group_name": types.StringType, "azure_subscription_id": types.StringType}}); ok {
				if typedAzureEventGrid, ok := nullAzureEventGrid.(types.Object); ok {
					state.AzureEventGrid = typedAzureEventGrid
				}
			}
		}
	}
	{
		if rawValueCreated, rawOk := plainValueAtPath(raw, "created"); rawOk {
			if valueCreated, err := flattenPlainValue(rawValueCreated, types.StringType, "created", "raw response"); err != nil {
				return err
			} else {
				if typedCreated, ok := valueCreated.(types.String); ok {
					state.Created = typedCreated
				}
			}
		} else if !hasRaw {
			if responseValueCreated, ok := plainFromResponseField(obj, "Created"); ok {
				if valueCreated, err := flattenPlainValue(responseValueCreated, types.StringType, "created", "response struct"); err != nil {
					return err
				} else {
					if typedCreated, ok := valueCreated.(types.String); ok {
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
		if rawValueEventPayload, rawOk := plainValueAtPath(raw, "event_payload"); rawOk {
			if valueEventPayload, err := flattenPlainValue(rawValueEventPayload, types.StringType, "event_payload", "raw response"); err != nil {
				return err
			} else {
				if typedEventPayload, ok := valueEventPayload.(types.String); ok {
					state.EventPayload = typedEventPayload
				}
			}
		} else if !hasRaw {
			if responseValueEventPayload, ok := plainFromResponseField(obj, "EventPayload"); ok {
				if valueEventPayload, err := flattenPlainValue(responseValueEventPayload, types.StringType, "event_payload", "response struct"); err != nil {
					return err
				} else {
					if typedEventPayload, ok := valueEventPayload.(types.String); ok {
						state.EventPayload = typedEventPayload
					}
				}
			}
		}
	}
	{
		if rawValueEventsFrom, rawOk := plainValueAtPath(raw, "events_from"); rawOk {
			if valueEventsFrom, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueEventsFrom, attrValueToPlain(state.EventsFrom)), types.ListType{ElemType: types.StringType}, "events_from", "raw response"); err != nil {
				return err
			} else {
				if typedEventsFrom, ok := valueEventsFrom.(types.List); ok {
					state.EventsFrom = typedEventsFrom
				}
			}
		} else if !hasRaw {
			if responseValueEventsFrom, ok := plainFromResponseField(obj, "EventsFrom"); ok {
				if valueEventsFrom, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueEventsFrom, attrValueToPlain(state.EventsFrom)),
					types.ListType{ElemType: types.StringType},
					"events_from",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedEventsFrom, ok := valueEventsFrom.(types.List); ok {
						state.EventsFrom = typedEventsFrom
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
		if rawValueSnapshotAPIVersion, rawOk := plainValueAtPath(raw, "snapshot_api_version"); rawOk {
			if valueSnapshotAPIVersion, err := flattenPlainValue(rawValueSnapshotAPIVersion, types.StringType, "snapshot_api_version", "raw response"); err != nil {
				return err
			} else {
				if typedSnapshotAPIVersion, ok := valueSnapshotAPIVersion.(types.String); ok {
					state.SnapshotAPIVersion = typedSnapshotAPIVersion
				}
			}
		} else if !hasRaw {
			if responseValueSnapshotAPIVersion, ok := plainFromResponseField(obj, "SnapshotAPIVersion"); ok {
				if valueSnapshotAPIVersion, err := flattenPlainValue(responseValueSnapshotAPIVersion, types.StringType, "snapshot_api_version", "response struct"); err != nil {
					return err
				} else {
					if typedSnapshotAPIVersion, ok := valueSnapshotAPIVersion.(types.String); ok {
						state.SnapshotAPIVersion = typedSnapshotAPIVersion
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
		assignedStatusDetails := false
		hadRawStatusDetails := false
		if rawValueStatusDetails, rawOk := plainValueAtPath(raw, "status_details"); rawOk {
			hadRawStatusDetails = true
			if rawValueStatusDetails != nil {
				sourceStatusDetails := applyConfiguredKeyedListShapes(rawValueStatusDetails, attrValueToPlain(state.StatusDetails))
				if valueStatusDetails, err := flattenPlainValue(sourceStatusDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"disabled": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}}}, "status_details", "raw response"); err != nil {
					return err
				} else {
					if typedStatusDetails, ok := valueStatusDetails.(types.Object); ok {
						state.StatusDetails = typedStatusDetails
						assignedStatusDetails = true
					}
				}
			}
		}
		if !assignedStatusDetails {
			if !hasRaw {
				if responseValueStatusDetails, ok := plainFromResponseField(obj, "StatusDetails"); ok {
					sourceStatusDetails := applyConfiguredKeyedListShapes(responseValueStatusDetails, attrValueToPlain(state.StatusDetails))
					if valueStatusDetails, err := flattenPlainValue(
						sourceStatusDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"disabled": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}}},
						"status_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedStatusDetails, ok := valueStatusDetails.(types.Object); ok {
							state.StatusDetails = typedStatusDetails
							assignedStatusDetails = true
						}
					}
				}
			}
		}
		if !assignedStatusDetails && hadRawStatusDetails {
			if nullStatusDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"disabled": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}}}); ok {
				if typedStatusDetails, ok := nullStatusDetails.(types.Object); ok {
					state.StatusDetails = typedStatusDetails
				}
			}
		}
	}
	{
		if rawValueType, rawOk := plainValueAtPath(raw, "type"); rawOk {
			if valueType, err := flattenPlainValue(rawValueType, types.StringType, "type", "raw response"); err != nil {
				return err
			} else {
				if typedType, ok := valueType.(types.String); ok {
					state.Type = typedType
				}
			}
		} else if !hasRaw {
			if responseValueType, ok := plainFromResponseField(obj, "Type"); ok {
				if valueType, err := flattenPlainValue(responseValueType, types.StringType, "type", "response struct"); err != nil {
					return err
				} else {
					if typedType, ok := valueType.(types.String); ok {
						state.Type = typedType
					}
				}
			}
		}
	}
	{
		if rawValueUpdated, rawOk := plainValueAtPath(raw, "updated"); rawOk {
			if valueUpdated, err := flattenPlainValue(rawValueUpdated, types.StringType, "updated", "raw response"); err != nil {
				return err
			} else {
				if typedUpdated, ok := valueUpdated.(types.String); ok {
					state.Updated = typedUpdated
				}
			}
		} else if !hasRaw {
			if responseValueUpdated, ok := plainFromResponseField(obj, "Updated"); ok {
				if valueUpdated, err := flattenPlainValue(responseValueUpdated, types.StringType, "updated", "response struct"); err != nil {
					return err
				} else {
					if typedUpdated, ok := valueUpdated.(types.String); ok {
						state.Updated = typedUpdated
					}
				}
			}
		}
	}
	{
		assignedWebhookEndpoint := false
		if rawValueWebhookEndpoint, rawOk := plainValueAtPath(raw, "webhook_endpoint"); rawOk {
			if rawValueWebhookEndpoint != nil {
				sourceWebhookEndpoint := mergeMissingPlainLeaves(applyConfiguredKeyedListShapes(rawValueWebhookEndpoint, unwrapPlainSingletonList(attrValueToPlain(state.WebhookEndpoint))), unwrapPlainSingletonList(attrValueToPlain(state.WebhookEndpoint)))
				if valueWebhookEndpoint, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceWebhookEndpoint, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"signing_secret": types.StringType, "url": types.StringType}}}, "webhook_endpoint", "raw response"); err != nil {
					return err
				} else {
					if typedWebhookEndpoint, ok := valueWebhookEndpoint.(types.List); ok {
						state.WebhookEndpoint = typedWebhookEndpoint
						assignedWebhookEndpoint = true
					}
				}
			}
		}
		if !assignedWebhookEndpoint {
			if !hasRaw {
				if responseValueWebhookEndpoint, ok := plainFromResponseField(obj, "WebhookEndpoint"); ok {
					sourceWebhookEndpoint := mergeMissingPlainLeaves(applyConfiguredKeyedListShapes(responseValueWebhookEndpoint, unwrapPlainSingletonList(attrValueToPlain(state.WebhookEndpoint))), unwrapPlainSingletonList(attrValueToPlain(state.WebhookEndpoint)))
					if valueWebhookEndpoint, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceWebhookEndpoint, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"signing_secret": types.StringType, "url": types.StringType}}},
						"webhook_endpoint",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedWebhookEndpoint, ok := valueWebhookEndpoint.(types.List); ok {
							state.WebhookEndpoint = typedWebhookEndpoint
							assignedWebhookEndpoint = true
						}
					}
				}
			}
		}
	}
	return nil
}
