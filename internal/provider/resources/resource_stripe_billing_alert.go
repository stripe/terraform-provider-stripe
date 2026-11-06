//
// File generated from our OpenAPI spec
//

package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &BillingAlertResource{}

var _ resource.ResourceWithConfigure = &BillingAlertResource{}

var _ resource.ResourceWithImportState = &BillingAlertResource{}

func NewBillingAlertResource() resource.Resource {
	return &BillingAlertResource{}
}

type BillingAlertResource struct {
	client *stripe.Client
}

type BillingAlertResourceModel struct {
	Object         types.String `tfsdk:"object"`
	AlertType      types.String `tfsdk:"alert_type"`
	ID             types.String `tfsdk:"id"`
	Livemode       types.Bool   `tfsdk:"livemode"`
	Status         types.String `tfsdk:"status"`
	Title          types.String `tfsdk:"title"`
	UsageThreshold types.Object `tfsdk:"usage_threshold"`
}

func (r *BillingAlertResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *BillingAlertResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_billing_alert"
}

func (r *BillingAlertResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A billing alert is a resource that notifies you when a certain usage threshold on a meter is crossed. For example, you might create a billing alert to notify you when a certain user made 100 API requests.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("billing.alert")},
			},
			"alert_type": schema.StringAttribute{
				Required:      true,
				Description:   "Defines the type of the alert.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("usage_threshold")},
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
				Description:   "Status of the alert. This can be active, inactive or archived.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "archived", "inactive")},
			},
			"title": schema.StringAttribute{
				Required:      true,
				Description:   "Title of the alert.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"usage_threshold": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Encapsulates configuration of the alert to monitor usage on a specific [Billing Meter](https://docs.stripe.com/api/billing/meter).",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"filters": schema.ListNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The filters allow limiting the scope of this usage alert. You can only specify up to one filter at this time.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"customer": schema.StringAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "Limit the scope of the alert to this customer ID",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
								},
								"type": schema.StringAttribute{
									Required: true,

									PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									Validators:    []validator.String{stringvalidator.OneOf("customer")},
								},
							},
						},
					},
					"gte": schema.Int64Attribute{
						Required:      true,
						Description:   "The value at which this alert will trigger.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
					"meter": schema.StringAttribute{
						Required:      true,
						Description:   "The [Billing Meter](/api/billing/meter) ID whose usage is monitored.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"recurrence": schema.StringAttribute{
						Required:      true,
						Description:   "Defines how the alert will behave.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("one_time")},
					},
				},
			},
		},
	}
}

func (r *BillingAlertResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan BillingAlertResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandBillingAlertCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building BillingAlert create params", err.Error())
		return
	}

	obj, err := r.client.V1BillingAlerts.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating BillingAlert", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingAlerts.B, r.client.V1BillingAlerts.Key, stripe.FormatURLPath("/v1/billing/alerts/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingAlert create raw response", err.Error())
		return
	}

	if err := flattenBillingAlert(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingAlert create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingAlertResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState BillingAlertResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state BillingAlertResourceModel
	state = priorState

	obj, err := r.client.V1BillingAlerts.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading BillingAlert", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingAlerts.B, r.client.V1BillingAlerts.Key, stripe.FormatURLPath("/v1/billing/alerts/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingAlert raw response", err.Error())
		return
	}

	if err := flattenBillingAlert(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingAlert read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingAlertResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "BillingAlert does not support updates")
}

func (r *BillingAlertResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *BillingAlertResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandBillingAlertCreate(plan BillingAlertResourceModel) (*stripe.BillingAlertCreateParams, error) {
	params := &stripe.BillingAlertCreateParams{}

	if !plan.AlertType.IsNull() && !plan.AlertType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "AlertType", "AlertType", plan.AlertType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "alert_type", params)
		}
	}
	if !plan.Title.IsNull() && !plan.Title.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Title", "Title", plan.Title.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "title", params)
		}
	}
	if !plan.UsageThreshold.IsNull() && !plan.UsageThreshold.IsUnknown() {
		if !assignAttrValueToNamedField(params, "UsageThreshold", plan.UsageThreshold) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "usage_threshold", params)
		}
	}

	return params, nil
}

func flattenBillingAlert(obj *stripe.BillingAlert, state *BillingAlertResourceModel) error {
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
		if rawValueAlertType, rawOk := plainValueAtPath(raw, "alert_type"); rawOk {
			if valueAlertType, err := flattenPlainValue(rawValueAlertType, types.StringType, "alert_type", "raw response"); err != nil {
				return err
			} else {
				if typedAlertType, ok := valueAlertType.(types.String); ok {
					state.AlertType = typedAlertType
				}
			}
		} else if !hasRaw {
			if responseValueAlertType, ok := plainFromResponseField(obj, "AlertType"); ok {
				if valueAlertType, err := flattenPlainValue(responseValueAlertType, types.StringType, "alert_type", "response struct"); err != nil {
					return err
				} else {
					if typedAlertType, ok := valueAlertType.(types.String); ok {
						state.AlertType = typedAlertType
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
		if rawValueTitle, rawOk := plainValueAtPath(raw, "title"); rawOk {
			if valueTitle, err := flattenPlainValue(rawValueTitle, types.StringType, "title", "raw response"); err != nil {
				return err
			} else {
				if typedTitle, ok := valueTitle.(types.String); ok {
					state.Title = typedTitle
				}
			}
		} else if !hasRaw {
			if responseValueTitle, ok := plainFromResponseField(obj, "Title"); ok {
				if valueTitle, err := flattenPlainValue(responseValueTitle, types.StringType, "title", "response struct"); err != nil {
					return err
				} else {
					if typedTitle, ok := valueTitle.(types.String); ok {
						state.Title = typedTitle
					}
				}
			}
		}
	}
	{
		assignedUsageThreshold := false
		hadRawUsageThreshold := false
		if rawValueUsageThreshold, rawOk := plainValueAtPath(raw, "usage_threshold"); rawOk {
			hadRawUsageThreshold = true
			if rawValueUsageThreshold != nil {
				sourceUsageThreshold := applyConfiguredKeyedListShapes(rawValueUsageThreshold, attrValueToPlain(state.UsageThreshold))
				if valueUsageThreshold, err := flattenPlainValue(sourceUsageThreshold, types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"customer": types.StringType, "type": types.StringType}}}, "gte": types.Int64Type, "meter": types.StringType, "recurrence": types.StringType}}, "usage_threshold", "raw response"); err != nil {
					return err
				} else {
					if typedUsageThreshold, ok := valueUsageThreshold.(types.Object); ok {
						state.UsageThreshold = typedUsageThreshold
						assignedUsageThreshold = true
					}
				}
			}
		}
		if !assignedUsageThreshold {
			if !hasRaw {
				if responseValueUsageThreshold, ok := plainFromResponseField(obj, "UsageThreshold"); ok {
					sourceUsageThreshold := applyConfiguredKeyedListShapes(responseValueUsageThreshold, attrValueToPlain(state.UsageThreshold))
					if valueUsageThreshold, err := flattenPlainValue(
						sourceUsageThreshold,
						types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"customer": types.StringType, "type": types.StringType}}}, "gte": types.Int64Type, "meter": types.StringType, "recurrence": types.StringType}},
						"usage_threshold",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedUsageThreshold, ok := valueUsageThreshold.(types.Object); ok {
							state.UsageThreshold = typedUsageThreshold
							assignedUsageThreshold = true
						}
					}
				}
			}
		}
		if !assignedUsageThreshold && hadRawUsageThreshold {
			if nullUsageThreshold, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"customer": types.StringType, "type": types.StringType}}}, "gte": types.Int64Type, "meter": types.StringType, "recurrence": types.StringType}}); ok {
				if typedUsageThreshold, ok := nullUsageThreshold.(types.Object); ok {
					state.UsageThreshold = typedUsageThreshold
				}
			}
		}
	}
	return nil
}
