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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &BillingFeedbackOptionResource{}

var _ resource.ResourceWithConfigure = &BillingFeedbackOptionResource{}

var _ resource.ResourceWithImportState = &BillingFeedbackOptionResource{}

func NewBillingFeedbackOptionResource() resource.Resource {
	return &BillingFeedbackOptionResource{}
}

type BillingFeedbackOptionResource struct {
	client *stripe.Client
}

type BillingFeedbackOptionResourceModel struct {
	Object            types.String `tfsdk:"object"`
	Description       types.String `tfsdk:"description"`
	ID                types.String `tfsdk:"id"`
	Livemode          types.Bool   `tfsdk:"livemode"`
	Status            types.String `tfsdk:"status"`
	StatusTransitions types.Object `tfsdk:"status_transitions"`
}

func (r *BillingFeedbackOptionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *BillingFeedbackOptionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_billing_feedback_option"
}

func (r *BillingFeedbackOptionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A feedback option is a reason you can present to customers when they cancel a\nsubscription through the customer portal. Configure the set of options a customer\ncan choose from on a [portal configuration](/api/customer_portal/configuration).\n\nRelated guide: [Customer management](/customer-management)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("billing.feedback_option")},
			},
			"description": schema.StringAttribute{
				Required:    true,
				Description: "An arbitrary string attached to the object. Often useful for displaying to users.",
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
				Description:   "The feedback option's status.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "inactive")},
			},
			"status_transitions": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"deactivated_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time the feedback option was deactivated, if any. Measured in seconds since Unix epoch.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
		},
	}
}

func (r *BillingFeedbackOptionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan BillingFeedbackOptionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandBillingFeedbackOptionCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building BillingFeedbackOption create params", err.Error())
		return
	}

	obj, err := r.client.V1BillingFeedbackOptions.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating BillingFeedbackOption", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingFeedbackOptions.B, r.client.V1BillingFeedbackOptions.Key, stripe.FormatURLPath("/v1/billing/feedback_options/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingFeedbackOption create raw response", err.Error())
		return
	}

	if err := flattenBillingFeedbackOption(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingFeedbackOption create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingFeedbackOptionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState BillingFeedbackOptionResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state BillingFeedbackOptionResourceModel
	state = priorState

	obj, err := r.client.V1BillingFeedbackOptions.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading BillingFeedbackOption", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingFeedbackOptions.B, r.client.V1BillingFeedbackOptions.Key, stripe.FormatURLPath("/v1/billing/feedback_options/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingFeedbackOption raw response", err.Error())
		return
	}

	if err := flattenBillingFeedbackOption(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingFeedbackOption read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingFeedbackOptionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan BillingFeedbackOptionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state BillingFeedbackOptionResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandBillingFeedbackOptionUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building BillingFeedbackOption update params", err.Error())
		return
	}

	obj, err := r.client.V1BillingFeedbackOptions.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating BillingFeedbackOption", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingFeedbackOptions.B, r.client.V1BillingFeedbackOptions.Key, stripe.FormatURLPath("/v1/billing/feedback_options/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingFeedbackOption update raw response", err.Error())
		return
	}

	if err := flattenBillingFeedbackOption(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingFeedbackOption update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingFeedbackOptionResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *BillingFeedbackOptionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandBillingFeedbackOptionCreate(plan BillingFeedbackOptionResourceModel) (*stripe.BillingFeedbackOptionCreateParams, error) {
	params := &stripe.BillingFeedbackOptionCreateParams{}

	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}

	return params, nil
}

func expandBillingFeedbackOptionUpdate(plan BillingFeedbackOptionResourceModel, state BillingFeedbackOptionResourceModel) (*stripe.BillingFeedbackOptionUpdateParams, error) {
	params := &stripe.BillingFeedbackOptionUpdateParams{}

	if !plan.Description.Equal(state.Description) && !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			if !plan.Description.Equal(state.Description) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "description", params)
			}
		}
	}

	return params, nil
}

func flattenBillingFeedbackOption(obj *stripe.BillingFeedbackOption, state *BillingFeedbackOptionResourceModel) error {
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
	return nil
}
