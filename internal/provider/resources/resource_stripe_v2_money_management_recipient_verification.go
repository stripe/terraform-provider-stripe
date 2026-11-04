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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &V2MoneyManagementRecipientVerificationResource{}

var _ resource.ResourceWithConfigure = &V2MoneyManagementRecipientVerificationResource{}

var _ resource.ResourceWithImportState = &V2MoneyManagementRecipientVerificationResource{}

func NewV2MoneyManagementRecipientVerificationResource() resource.Resource {
	return &V2MoneyManagementRecipientVerificationResource{}
}

type V2MoneyManagementRecipientVerificationResource struct {
	client *stripe.Client
}

type V2MoneyManagementRecipientVerificationResourceModel struct {
	Object             types.String `tfsdk:"object"`
	ConsumedBy         types.Object `tfsdk:"consumed_by"`
	Created            types.String `tfsdk:"created"`
	ExpiresAt          types.String `tfsdk:"expires_at"`
	ID                 types.String `tfsdk:"id"`
	Livemode           types.Bool   `tfsdk:"livemode"`
	MatchResult        types.String `tfsdk:"match_result"`
	MatchResultDetails types.Object `tfsdk:"match_result_details"`
	PayoutMethod       types.String `tfsdk:"payout_method"`
	Recipient          types.String `tfsdk:"recipient"`
	Status             types.String `tfsdk:"status"`
	StatusTransitions  types.Object `tfsdk:"status_transitions"`
	VerificationType   types.String `tfsdk:"verification_type"`
}

func (r *V2MoneyManagementRecipientVerificationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *V2MoneyManagementRecipientVerificationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v2_money_management_recipient_verification"
}

func (r *V2MoneyManagementRecipientVerificationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "RecipientVerification represents a verification of recipient you intend to send funds to.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value of the object field.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("v2.money_management.recipient_verification")},
			},
			"consumed_by": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The resource that consumed this verification.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"outbound_payment": schema.StringAttribute{
						Computed:      true,
						Description:   "The ID of the OutboundPayment that consumed the RecipientVerification.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "The type of resource that consumed the RecipientVerification.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("outbound_payment")},
					},
				},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				Description:   "Time at which the RecipientVerification was created.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"expires_at": schema.StringAttribute{
				Computed:      true,
				Description:   "Time at which the RecipientVerification expires, 5 minutes after the creation. Only applicable for Verification of Payee (VoP) verifications.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the RecipientVerification.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"match_result": schema.StringAttribute{
				Computed:      true,
				Description:   "Closed Enum. Match level of the RecipientVerification: `match`, `close_match`, `no_match`, `unavailable`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("close_match", "match", "no_match", "unavailable")},
			},
			"match_result_details": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Details for the match result.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"matched_business_type": schema.StringAttribute{
						Computed:      true,
						Description:   "The business type matched by the provider. Only present for Confirmation of Payee (CoP) verifications.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("business", "personal")},
					},
					"matched_name": schema.StringAttribute{
						Computed:      true,
						Description:   "The account name associated with the bank account as provided by the VoP provider, only present if there is a match or close match.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"message": schema.StringAttribute{
						Computed:      true,
						Description:   "A message describing the match result.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"provided_business_type": schema.StringAttribute{
						Computed:      true,
						Description:   "The business type provided by the user. Only present for Confirmation of Payee (CoP) verifications.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("business", "personal")},
					},
					"provided_name": schema.StringAttribute{
						Computed:      true,
						Description:   "The name associated with the provided recipient.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"payout_method": schema.StringAttribute{
				Required:      true,
				Description:   "ID of the payout method that was verified.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"recipient": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the recipient account. Omitted for same-entity payouts.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Current status of the recipient verification.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("acknowledged", "awaiting_acknowledgement", "consumed", "expired", "verified")},
			},
			"status_transitions": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Hash containing timestamps of when the object transitioned to a particular status.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"acknowledged_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp describing when a RecipientVerification changed status to `acknowledged`.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"consumed_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp describing when a RecipientVerification changed status to `consumed`.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"verification_type": schema.StringAttribute{
				Computed:      true,
				Description:   "Open Enum. The type of verification performed: `verification_of_payee`, `confirmation_of_payee`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("confirmation_of_payee", "verification_of_payee")},
			},
		},
	}
}

func (r *V2MoneyManagementRecipientVerificationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan V2MoneyManagementRecipientVerificationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandV2MoneyManagementRecipientVerificationCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2MoneyManagementRecipientVerification create params", err.Error())
		return
	}

	obj, err := r.client.V2MoneyManagementRecipientVerifications.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating V2MoneyManagementRecipientVerification", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementRecipientVerifications.B, r.client.V2MoneyManagementRecipientVerifications.Key, stripe.FormatURLPath("/v2/money_management/recipient_verifications/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementRecipientVerification create raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementRecipientVerification(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementRecipientVerification create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementRecipientVerificationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState V2MoneyManagementRecipientVerificationResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2MoneyManagementRecipientVerificationResourceModel
	state = priorState

	obj, err := r.client.V2MoneyManagementRecipientVerifications.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading V2MoneyManagementRecipientVerification", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementRecipientVerifications.B, r.client.V2MoneyManagementRecipientVerifications.Key, stripe.FormatURLPath("/v2/money_management/recipient_verifications/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementRecipientVerification raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementRecipientVerification(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementRecipientVerification read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementRecipientVerificationResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "V2MoneyManagementRecipientVerification does not support updates")
}

func (r *V2MoneyManagementRecipientVerificationResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *V2MoneyManagementRecipientVerificationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandV2MoneyManagementRecipientVerificationCreate(plan V2MoneyManagementRecipientVerificationResourceModel) (*stripe.V2MoneyManagementRecipientVerificationCreateParams, error) {
	params := &stripe.V2MoneyManagementRecipientVerificationCreateParams{}

	if !plan.PayoutMethod.IsNull() && !plan.PayoutMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PayoutMethod", "PayoutMethod", plan.PayoutMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payout_method", params)
		}
	}
	if !plan.Recipient.IsNull() && !plan.Recipient.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Recipient", "Recipient", plan.Recipient.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "recipient", params)
		}
	}

	return params, nil
}

func flattenV2MoneyManagementRecipientVerification(obj *stripe.V2MoneyManagementRecipientVerification, state *V2MoneyManagementRecipientVerificationResourceModel) error {
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
		assignedConsumedBy := false
		hadRawConsumedBy := false
		if rawValueConsumedBy, rawOk := plainValueAtPath(raw, "consumed_by"); rawOk {
			hadRawConsumedBy = true
			if rawValueConsumedBy != nil {
				sourceConsumedBy := applyConfiguredKeyedListShapes(rawValueConsumedBy, attrValueToPlain(state.ConsumedBy))
				if valueConsumedBy, err := flattenPlainValue(sourceConsumedBy, types.ObjectType{AttrTypes: map[string]attr.Type{"outbound_payment": types.StringType, "type": types.StringType}}, "consumed_by", "raw response"); err != nil {
					return err
				} else {
					if typedConsumedBy, ok := valueConsumedBy.(types.Object); ok {
						state.ConsumedBy = typedConsumedBy
						assignedConsumedBy = true
					}
				}
			}
		}
		if !assignedConsumedBy {
			if !hasRaw {
				if responseValueConsumedBy, ok := plainFromResponseField(obj, "ConsumedBy"); ok {
					sourceConsumedBy := applyConfiguredKeyedListShapes(responseValueConsumedBy, attrValueToPlain(state.ConsumedBy))
					if valueConsumedBy, err := flattenPlainValue(
						sourceConsumedBy,
						types.ObjectType{AttrTypes: map[string]attr.Type{"outbound_payment": types.StringType, "type": types.StringType}},
						"consumed_by",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedConsumedBy, ok := valueConsumedBy.(types.Object); ok {
							state.ConsumedBy = typedConsumedBy
							assignedConsumedBy = true
						}
					}
				}
			}
		}
		if !assignedConsumedBy && hadRawConsumedBy {
			if nullConsumedBy, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"outbound_payment": types.StringType, "type": types.StringType}}); ok {
				if typedConsumedBy, ok := nullConsumedBy.(types.Object); ok {
					state.ConsumedBy = typedConsumedBy
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
		if rawValueExpiresAt, rawOk := plainValueAtPath(raw, "expires_at"); rawOk {
			if valueExpiresAt, err := flattenPlainValue(rawValueExpiresAt, types.StringType, "expires_at", "raw response"); err != nil {
				return err
			} else {
				if typedExpiresAt, ok := valueExpiresAt.(types.String); ok {
					state.ExpiresAt = typedExpiresAt
				}
			}
		} else if !hasRaw {
			if responseValueExpiresAt, ok := plainFromResponseField(obj, "ExpiresAt"); ok {
				if valueExpiresAt, err := flattenPlainValue(responseValueExpiresAt, types.StringType, "expires_at", "response struct"); err != nil {
					return err
				} else {
					if typedExpiresAt, ok := valueExpiresAt.(types.String); ok {
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
		if rawValueMatchResult, rawOk := plainValueAtPath(raw, "match_result"); rawOk {
			if valueMatchResult, err := flattenPlainValue(rawValueMatchResult, types.StringType, "match_result", "raw response"); err != nil {
				return err
			} else {
				if typedMatchResult, ok := valueMatchResult.(types.String); ok {
					state.MatchResult = typedMatchResult
				}
			}
		} else if !hasRaw {
			if responseValueMatchResult, ok := plainFromResponseField(obj, "MatchResult"); ok {
				if valueMatchResult, err := flattenPlainValue(responseValueMatchResult, types.StringType, "match_result", "response struct"); err != nil {
					return err
				} else {
					if typedMatchResult, ok := valueMatchResult.(types.String); ok {
						state.MatchResult = typedMatchResult
					}
				}
			}
		}
	}
	{
		assignedMatchResultDetails := false
		hadRawMatchResultDetails := false
		if rawValueMatchResultDetails, rawOk := plainValueAtPath(raw, "match_result_details"); rawOk {
			hadRawMatchResultDetails = true
			if rawValueMatchResultDetails != nil {
				sourceMatchResultDetails := applyConfiguredKeyedListShapes(rawValueMatchResultDetails, attrValueToPlain(state.MatchResultDetails))
				if valueMatchResultDetails, err := flattenPlainValue(sourceMatchResultDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"matched_business_type": types.StringType, "matched_name": types.StringType, "message": types.StringType, "provided_business_type": types.StringType, "provided_name": types.StringType}}, "match_result_details", "raw response"); err != nil {
					return err
				} else {
					if typedMatchResultDetails, ok := valueMatchResultDetails.(types.Object); ok {
						state.MatchResultDetails = typedMatchResultDetails
						assignedMatchResultDetails = true
					}
				}
			}
		}
		if !assignedMatchResultDetails {
			if !hasRaw {
				if responseValueMatchResultDetails, ok := plainFromResponseField(obj, "MatchResultDetails"); ok {
					sourceMatchResultDetails := applyConfiguredKeyedListShapes(responseValueMatchResultDetails, attrValueToPlain(state.MatchResultDetails))
					if valueMatchResultDetails, err := flattenPlainValue(
						sourceMatchResultDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"matched_business_type": types.StringType, "matched_name": types.StringType, "message": types.StringType, "provided_business_type": types.StringType, "provided_name": types.StringType}},
						"match_result_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedMatchResultDetails, ok := valueMatchResultDetails.(types.Object); ok {
							state.MatchResultDetails = typedMatchResultDetails
							assignedMatchResultDetails = true
						}
					}
				}
			}
		}
		if !assignedMatchResultDetails && hadRawMatchResultDetails {
			if nullMatchResultDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"matched_business_type": types.StringType, "matched_name": types.StringType, "message": types.StringType, "provided_business_type": types.StringType, "provided_name": types.StringType}}); ok {
				if typedMatchResultDetails, ok := nullMatchResultDetails.(types.Object); ok {
					state.MatchResultDetails = typedMatchResultDetails
				}
			}
		}
	}
	{
		if rawValuePayoutMethod, rawOk := plainValueAtPath(raw, "payout_method"); rawOk {
			if valuePayoutMethod, err := flattenPlainValue(rawValuePayoutMethod, types.StringType, "payout_method", "raw response"); err != nil {
				return err
			} else {
				if typedPayoutMethod, ok := valuePayoutMethod.(types.String); ok {
					state.PayoutMethod = typedPayoutMethod
				}
			}
		} else if !hasRaw {
			if responseValuePayoutMethod, ok := plainFromResponseField(obj, "PayoutMethod"); ok {
				if valuePayoutMethod, err := flattenPlainValue(responseValuePayoutMethod, types.StringType, "payout_method", "response struct"); err != nil {
					return err
				} else {
					if typedPayoutMethod, ok := valuePayoutMethod.(types.String); ok {
						state.PayoutMethod = typedPayoutMethod
					}
				}
			}
		}
	}
	{
		if rawValueRecipient, rawOk := plainValueAtPath(raw, "recipient"); rawOk {
			if valueRecipient, err := flattenPlainValue(rawValueRecipient, types.StringType, "recipient", "raw response"); err != nil {
				return err
			} else {
				if typedRecipient, ok := valueRecipient.(types.String); ok {
					state.Recipient = typedRecipient
				}
			}
		} else if !hasRaw {
			if responseValueRecipient, ok := plainFromResponseField(obj, "Recipient"); ok {
				if valueRecipient, err := flattenPlainValue(responseValueRecipient, types.StringType, "recipient", "response struct"); err != nil {
					return err
				} else {
					if typedRecipient, ok := valueRecipient.(types.String); ok {
						state.Recipient = typedRecipient
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
				if valueStatusTransitions, err := flattenPlainValue(sourceStatusTransitions, types.ObjectType{AttrTypes: map[string]attr.Type{"acknowledged_at": types.StringType, "consumed_at": types.StringType}}, "status_transitions", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"acknowledged_at": types.StringType, "consumed_at": types.StringType}},
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
			if nullStatusTransitions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"acknowledged_at": types.StringType, "consumed_at": types.StringType}}); ok {
				if typedStatusTransitions, ok := nullStatusTransitions.(types.Object); ok {
					state.StatusTransitions = typedStatusTransitions
				}
			}
		}
	}
	{
		if rawValueVerificationType, rawOk := plainValueAtPath(raw, "verification_type"); rawOk {
			if valueVerificationType, err := flattenPlainValue(rawValueVerificationType, types.StringType, "verification_type", "raw response"); err != nil {
				return err
			} else {
				if typedVerificationType, ok := valueVerificationType.(types.String); ok {
					state.VerificationType = typedVerificationType
				}
			}
		} else if !hasRaw {
			if responseValueVerificationType, ok := plainFromResponseField(obj, "VerificationType"); ok {
				if valueVerificationType, err := flattenPlainValue(responseValueVerificationType, types.StringType, "verification_type", "response struct"); err != nil {
					return err
				} else {
					if typedVerificationType, ok := valueVerificationType.(types.String); ok {
						state.VerificationType = typedVerificationType
					}
				}
			}
		}
	}
	return nil
}
