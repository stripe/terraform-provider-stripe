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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &V2MoneyManagementOutboundPaymentResource{}

var _ resource.ResourceWithConfigure = &V2MoneyManagementOutboundPaymentResource{}

var _ resource.ResourceWithImportState = &V2MoneyManagementOutboundPaymentResource{}

func NewV2MoneyManagementOutboundPaymentResource() resource.Resource {
	return &V2MoneyManagementOutboundPaymentResource{}
}

type V2MoneyManagementOutboundPaymentResource struct {
	client *stripe.Client
}

type V2MoneyManagementOutboundPaymentResourceModel struct {
	Object                types.String `tfsdk:"object"`
	Amount                types.Object `tfsdk:"amount"`
	Cancelable            types.Bool   `tfsdk:"cancelable"`
	Created               types.String `tfsdk:"created"`
	Description           types.String `tfsdk:"description"`
	ExpectedArrivalDate   types.String `tfsdk:"expected_arrival_date"`
	From                  types.Object `tfsdk:"from"`
	ID                    types.String `tfsdk:"id"`
	Livemode              types.Bool   `tfsdk:"livemode"`
	Metadata              types.Map    `tfsdk:"metadata"`
	PayoutIntent          types.String `tfsdk:"payout_intent"`
	RecipientNotification types.Object `tfsdk:"recipient_notification"`
	RecipientVerification types.String `tfsdk:"recipient_verification"`
	StatementDescriptor   types.String `tfsdk:"statement_descriptor"`
	Status                types.String `tfsdk:"status"`
	StatusDetails         types.Object `tfsdk:"status_details"`
	StatusTransitions     types.Object `tfsdk:"status_transitions"`
	To                    types.Object `tfsdk:"to"`
	TrackingDetails       types.Object `tfsdk:"tracking_details"`
}

func (r *V2MoneyManagementOutboundPaymentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *V2MoneyManagementOutboundPaymentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v2_money_management_outbound_payment"
}

func (r *V2MoneyManagementOutboundPaymentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "OutboundPayment represents a single money movement from one FinancialAccount you own to a payout method someone else owns.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value of the object field.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("v2.money_management.outbound_payment")},
			},
			"amount": schema.SingleNestedAttribute{
				Required:      true,
				Description:   "The amount and currency requested for this OutboundPayment.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"currency": schema.StringAttribute{
						Required:      true,
						Description:   "The currency of the monetary amount.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"value": schema.Int64Attribute{
						Required:      true,
						Description:   "The monetary amount in minor units of the currency.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
				},
			},
			"cancelable": schema.BoolAttribute{
				Computed:      true,
				Description:   "Returns true if the OutboundPayment can be canceled, and false otherwise.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				Description:   "Time at which the OutboundPayment was created.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the OutboundPayment. Often useful for displaying to users.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"expected_arrival_date": schema.StringAttribute{
				Computed:      true,
				Description:   "The date when funds are expected to arrive in the payout method. This field is not set if the payout method is in a `failed`, `canceled`, or `returned` state.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"from": schema.SingleNestedAttribute{
				Required:      true,
				Description:   "The FinancialAccount that funds were pulled from.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"debited": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "The monetary amount debited from the sender, only set on responses.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"currency": schema.StringAttribute{
								Computed:      true,
								Description:   "The currency of the monetary amount.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"value": schema.Int64Attribute{
								Computed:      true,
								Description:   "The monetary amount in minor units of the currency.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"financial_account": schema.StringAttribute{
						Required:      true,
						Description:   "The FinancialAccount that funds were pulled from.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"currency": schema.StringAttribute{
						Required:      true,
						Description:   "Describes the FinancialAmount's currency drawn from.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the OutboundPayment.",
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
				Description:   "Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown(), mapplanmodifier.RequiresReplace()},
				ElementType:   types.StringType,
			},
			"payout_intent": schema.StringAttribute{
				Computed:      true,
				Description:   "The PayoutIntent ID that triggered this OutboundPayment.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"recipient_notification": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Details about the OutboundPayment notification settings for recipient.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"setting": schema.StringAttribute{
						Required:      true,
						Description:   "Closed Enum. Configuration option to enable or disable notifications to recipients.\nDo not send notifications when `setting` is `none`. Use the account setting when `setting` is `configured` or omitted.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("configured", "none")},
					},
				},
			},
			"recipient_verification": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ID of the RecipientVerification for this OutboundPayment. Required where recipient-name verification is mandated before creating an OutboundPayment.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"statement_descriptor": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The description that appears on the receiving end of an OutboundPayment (for example, on a bank statement). Defaults to `STRIPE` if not configured in the account settings.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Current status of the OutboundPayment.\nAn OutboundPayment is `processing` after creation. It becomes `posted` when funds leave the account,\n`failed` if it cannot be sent, `canceled` if canceled before posting, or `returned` if the funds are returned.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("canceled", "failed", "posted", "processing", "returned")},
			},
			"status_details": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Status details for an OutboundPayment in a `processing`, `failed`, or `returned` state.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"failed": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "The `failed` status reason.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"reason": schema.StringAttribute{
								Computed:      true,
								Description:   "Open Enum. The `failed` status reason.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("fx_rate_drift_exceeded_after_review", "payout_method_amount_limit_exceeded", "payout_method_declined", "payout_method_does_not_exist", "payout_method_expired", "payout_method_unsupported", "payout_method_usage_frequency_limit_exceeded", "review_rejected", "unknown_failure")},
							},
						},
					},
					"processing": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "The `processing` status details.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"reason": schema.StringAttribute{
								Computed:      true,
								Description:   "Open Enum. The `processing` status reason.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("under_review")},
							},
						},
					},
					"returned": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "The `returned` status reason.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"reason": schema.StringAttribute{
								Computed:      true,
								Description:   "Open Enum. The `returned` status reason.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("payout_method_canceled_by_customer", "payout_method_closed", "payout_method_currency_unsupported", "payout_method_does_not_exist", "payout_method_holder_address_incorrect", "payout_method_holder_details_incorrect", "payout_method_holder_name_incorrect", "payout_method_invalid_account_number", "payout_method_restricted", "recalled", "unknown_failure")},
							},
						},
					},
				},
			},
			"status_transitions": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Hash containing timestamps of when the object transitioned to a particular status.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"canceled_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp describing when an OutboundPayment changed status to `canceled`.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"failed_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp describing when an OutboundPayment changed status to `failed`.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"posted_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp describing when an OutboundPayment changed status to `posted`.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"returned_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp describing when an OutboundPayment changed status to `returned`.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"to": schema.SingleNestedAttribute{
				Required:      true,
				Description:   "To which payout method the OutboundPayment was sent.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"credited": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "The monetary amount being credited to the destination.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"currency": schema.StringAttribute{
								Computed:      true,
								Description:   "The currency of the monetary amount.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"value": schema.Int64Attribute{
								Computed:      true,
								Description:   "The monetary amount in minor units of the currency.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"payout_method": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The payout method which the OutboundPayment uses to send payout.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"payout_method_options": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Payout method options for the OutboundPayment.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"bank_account": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Options for bank account payout methods.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"preferred_network_options": schema.SingleNestedAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Per-network configuration options.",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
										Attributes: map[string]schema.Attribute{
											"ach": schema.SingleNestedAttribute{
												Optional:      true,
												Computed:      true,
												Description:   "ACH-specific network options.",
												PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
												Attributes: map[string]schema.Attribute{
													"submission": schema.StringAttribute{
														Optional:      true,
														Computed:      true,
														Description:   "Open Enum. ACH submission timing.",
														PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
														Validators:    []validator.String{stringvalidator.OneOf("next_day", "same_day")},
													},
													"transaction_purpose": schema.StringAttribute{
														Optional:      true,
														Computed:      true,
														Description:   "The transaction purpose for this ACH payment.",
														PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
														Validators:    []validator.String{stringvalidator.OneOf("payroll")},
													},
												},
											},
										},
									},
									"preferred_networks": schema.ListAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The preferred networks to use for this OutboundPayment.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
										ElementType:   types.StringType,
									},
								},
							},
						},
					},
					"recipient": schema.StringAttribute{
						Required:      true,
						Description:   "To which account the OutboundPayment is sent.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"currency": schema.StringAttribute{
						Optional:      true,
						Description:   "Describes the currency to send to the recipient.\nIf included, this currency must match a currency supported by the destination.\nCan be omitted in the following cases:\n- destination only supports one currency\n- destination supports multiple currencies and one matches `from.currency`\n- destination supports multiple currencies and one matches `amount.currency`\nWhen both `from.currency` and `amount.currency` are supported, Stripe selects `from.currency` to minimize currency conversion.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
				},
			},
			"tracking_details": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Information to track this OutboundPayment with the recipient bank.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bank_account": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Bank account tracking details.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"trace_id": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique identifier that can be used to track this OutboundPayment with the recipient bank. Banks might call this a “reference number” or something similar.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"card": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Card tracking details.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"retrieval_reference": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique identifier that can be used to track this OutboundPayment with the recipient card.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the tracking details for this OutboundPayment.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("pending", "supported", "unsupported")},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "The type of tracking details for this OutboundPayment.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("bank_account", "card")},
					},
				},
			},
		},
	}
}

func (r *V2MoneyManagementOutboundPaymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan V2MoneyManagementOutboundPaymentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config V2MoneyManagementOutboundPaymentResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"From", "currency"}, []string{"To", "currency"}})

	params, err := expandV2MoneyManagementOutboundPaymentCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2MoneyManagementOutboundPayment create params", err.Error())
		return
	}

	obj, err := r.client.V2MoneyManagementOutboundPayments.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating V2MoneyManagementOutboundPayment", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementOutboundPayments.B, r.client.V2MoneyManagementOutboundPayments.Key, stripe.FormatURLPath("/v2/money_management/outbound_payments/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementOutboundPayment create raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementOutboundPayment(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementOutboundPayment create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"From", "currency"}, []string{"To", "currency"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementOutboundPaymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState V2MoneyManagementOutboundPaymentResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2MoneyManagementOutboundPaymentResourceModel
	state = priorState

	obj, err := r.client.V2MoneyManagementOutboundPayments.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading V2MoneyManagementOutboundPayment", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementOutboundPayments.B, r.client.V2MoneyManagementOutboundPayments.Key, stripe.FormatURLPath("/v2/money_management/outbound_payments/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementOutboundPayment raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementOutboundPayment(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementOutboundPayment read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"From", "currency"}, []string{"To", "currency"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementOutboundPaymentResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "V2MoneyManagementOutboundPayment does not support updates")
}

func (r *V2MoneyManagementOutboundPaymentResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *V2MoneyManagementOutboundPaymentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandV2MoneyManagementOutboundPaymentCreate(plan V2MoneyManagementOutboundPaymentResourceModel) (*stripe.V2MoneyManagementOutboundPaymentCreateParams, error) {
	params := &stripe.V2MoneyManagementOutboundPaymentCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Amount", plan.Amount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "amount", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.From.IsNull() && !plan.From.IsUnknown() {
		if !assignAttrValueToNamedField(params, "From", plan.From) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "from", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.RecipientNotification.IsNull() && !plan.RecipientNotification.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RecipientNotification", plan.RecipientNotification) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "recipient_notification", params)
		}
	}
	if !plan.RecipientVerification.IsNull() && !plan.RecipientVerification.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "RecipientVerification", "RecipientVerification", plan.RecipientVerification.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "recipient_verification", params)
		}
	}
	if !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "statement_descriptor", params)
		}
	}
	if !plan.To.IsNull() && !plan.To.IsUnknown() {
		if !assignAttrValueToNamedField(params, "To", plan.To) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "to", params)
		}
	}

	return params, nil
}

func flattenV2MoneyManagementOutboundPayment(obj *stripe.V2MoneyManagementOutboundPayment, state *V2MoneyManagementOutboundPaymentResourceModel) error {
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
		assignedAmount := false
		hadRawAmount := false
		if rawValueAmount, rawOk := plainValueAtPath(raw, "amount"); rawOk {
			hadRawAmount = true
			if rawValueAmount != nil {
				sourceAmount := applyConfiguredKeyedListShapes(rawValueAmount, attrValueToPlain(state.Amount))
				if valueAmount, err := flattenPlainValue(sourceAmount, types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "amount", "raw response"); err != nil {
					return err
				} else {
					if typedAmount, ok := valueAmount.(types.Object); ok {
						state.Amount = typedAmount
						assignedAmount = true
					}
				}
			}
		}
		if !assignedAmount {
			if !hasRaw {
				if responseValueAmount, ok := plainFromResponseField(obj, "Amount"); ok {
					sourceAmount := applyConfiguredKeyedListShapes(responseValueAmount, attrValueToPlain(state.Amount))
					if valueAmount, err := flattenPlainValue(
						sourceAmount,
						types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}},
						"amount",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAmount, ok := valueAmount.(types.Object); ok {
							state.Amount = typedAmount
							assignedAmount = true
						}
					}
				}
			}
		}
		if !assignedAmount && hadRawAmount {
			if nullAmount, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}); ok {
				if typedAmount, ok := nullAmount.(types.Object); ok {
					state.Amount = typedAmount
				}
			}
		}
	}
	{
		if rawValueCancelable, rawOk := plainValueAtPath(raw, "cancelable"); rawOk {
			if valueCancelable, err := flattenPlainValue(rawValueCancelable, types.BoolType, "cancelable", "raw response"); err != nil {
				return err
			} else {
				if typedCancelable, ok := valueCancelable.(types.Bool); ok {
					state.Cancelable = typedCancelable
				}
			}
		} else if !hasRaw {
			if responseValueCancelable, ok := plainFromResponseField(obj, "Cancelable"); ok {
				if valueCancelable, err := flattenPlainValue(responseValueCancelable, types.BoolType, "cancelable", "response struct"); err != nil {
					return err
				} else {
					if typedCancelable, ok := valueCancelable.(types.Bool); ok {
						state.Cancelable = typedCancelable
					}
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
		if rawValueExpectedArrivalDate, rawOk := plainValueAtPath(raw, "expected_arrival_date"); rawOk {
			if valueExpectedArrivalDate, err := flattenPlainValue(rawValueExpectedArrivalDate, types.StringType, "expected_arrival_date", "raw response"); err != nil {
				return err
			} else {
				if typedExpectedArrivalDate, ok := valueExpectedArrivalDate.(types.String); ok {
					state.ExpectedArrivalDate = typedExpectedArrivalDate
				}
			}
		} else if !hasRaw {
			if responseValueExpectedArrivalDate, ok := plainFromResponseField(obj, "ExpectedArrivalDate"); ok {
				if valueExpectedArrivalDate, err := flattenPlainValue(responseValueExpectedArrivalDate, types.StringType, "expected_arrival_date", "response struct"); err != nil {
					return err
				} else {
					if typedExpectedArrivalDate, ok := valueExpectedArrivalDate.(types.String); ok {
						state.ExpectedArrivalDate = typedExpectedArrivalDate
					}
				}
			}
		}
	}
	{
		assignedFrom := false
		hadRawFrom := false
		if rawValueFrom, rawOk := plainValueAtPath(raw, "from"); rawOk {
			hadRawFrom = true
			if rawValueFrom != nil {
				sourceFrom := applyConfiguredKeyedListShapes(rawValueFrom, attrValueToPlain(state.From))
				if valueFrom, err := flattenPlainValue(sourceFrom, types.ObjectType{AttrTypes: map[string]attr.Type{"debited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "financial_account": types.StringType, "currency": types.StringType}}, "from", "raw response"); err != nil {
					return err
				} else {
					if typedFrom, ok := valueFrom.(types.Object); ok {
						state.From = typedFrom
						assignedFrom = true
					}
				}
			}
		}
		if !assignedFrom {
			if !hasRaw {
				if responseValueFrom, ok := plainFromResponseField(obj, "From"); ok {
					sourceFrom := applyConfiguredKeyedListShapes(responseValueFrom, attrValueToPlain(state.From))
					if valueFrom, err := flattenPlainValue(
						sourceFrom,
						types.ObjectType{AttrTypes: map[string]attr.Type{"debited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "financial_account": types.StringType, "currency": types.StringType}},
						"from",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFrom, ok := valueFrom.(types.Object); ok {
							state.From = typedFrom
							assignedFrom = true
						}
					}
				}
			}
		}
		if !assignedFrom && hadRawFrom {
			if nullFrom, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"debited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "financial_account": types.StringType, "currency": types.StringType}}); ok {
				if typedFrom, ok := nullFrom.(types.Object); ok {
					state.From = typedFrom
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
		if rawValuePayoutIntent, rawOk := plainValueAtPath(raw, "payout_intent"); rawOk {
			if valuePayoutIntent, err := flattenPlainValue(rawValuePayoutIntent, types.StringType, "payout_intent", "raw response"); err != nil {
				return err
			} else {
				if typedPayoutIntent, ok := valuePayoutIntent.(types.String); ok {
					state.PayoutIntent = typedPayoutIntent
				}
			}
		} else if !hasRaw {
			if responseValuePayoutIntent, ok := plainFromResponseField(obj, "PayoutIntent"); ok {
				if valuePayoutIntent, err := flattenPlainValue(responseValuePayoutIntent, types.StringType, "payout_intent", "response struct"); err != nil {
					return err
				} else {
					if typedPayoutIntent, ok := valuePayoutIntent.(types.String); ok {
						state.PayoutIntent = typedPayoutIntent
					}
				}
			}
		}
	}
	{
		assignedRecipientNotification := false
		hadRawRecipientNotification := false
		if rawValueRecipientNotification, rawOk := plainValueAtPath(raw, "recipient_notification"); rawOk {
			hadRawRecipientNotification = true
			if rawValueRecipientNotification != nil {
				sourceRecipientNotification := applyConfiguredKeyedListShapes(rawValueRecipientNotification, attrValueToPlain(state.RecipientNotification))
				if valueRecipientNotification, err := flattenPlainValue(sourceRecipientNotification, types.ObjectType{AttrTypes: map[string]attr.Type{"setting": types.StringType}}, "recipient_notification", "raw response"); err != nil {
					return err
				} else {
					if typedRecipientNotification, ok := valueRecipientNotification.(types.Object); ok {
						state.RecipientNotification = typedRecipientNotification
						assignedRecipientNotification = true
					}
				}
			}
		}
		if !assignedRecipientNotification {
			if !hasRaw {
				if responseValueRecipientNotification, ok := plainFromResponseField(obj, "RecipientNotification"); ok {
					sourceRecipientNotification := applyConfiguredKeyedListShapes(responseValueRecipientNotification, attrValueToPlain(state.RecipientNotification))
					if valueRecipientNotification, err := flattenPlainValue(
						sourceRecipientNotification,
						types.ObjectType{AttrTypes: map[string]attr.Type{"setting": types.StringType}},
						"recipient_notification",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRecipientNotification, ok := valueRecipientNotification.(types.Object); ok {
							state.RecipientNotification = typedRecipientNotification
							assignedRecipientNotification = true
						}
					}
				}
			}
		}
		if !assignedRecipientNotification && hadRawRecipientNotification {
			if nullRecipientNotification, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"setting": types.StringType}}); ok {
				if typedRecipientNotification, ok := nullRecipientNotification.(types.Object); ok {
					state.RecipientNotification = typedRecipientNotification
				}
			}
		}
	}
	{
		if rawValueRecipientVerification, rawOk := plainValueAtPath(raw, "recipient_verification"); rawOk {
			if valueRecipientVerification, err := flattenPlainValue(rawValueRecipientVerification, types.StringType, "recipient_verification", "raw response"); err != nil {
				return err
			} else {
				if typedRecipientVerification, ok := valueRecipientVerification.(types.String); ok {
					state.RecipientVerification = typedRecipientVerification
				}
			}
		} else if !hasRaw {
			if responseValueRecipientVerification, ok := plainFromResponseField(obj, "RecipientVerification"); ok {
				if valueRecipientVerification, err := flattenPlainValue(responseValueRecipientVerification, types.StringType, "recipient_verification", "response struct"); err != nil {
					return err
				} else {
					if typedRecipientVerification, ok := valueRecipientVerification.(types.String); ok {
						state.RecipientVerification = typedRecipientVerification
					}
				}
			}
		}
	}
	{
		if rawValueStatementDescriptor, rawOk := plainValueAtPath(raw, "statement_descriptor"); rawOk {
			if valueStatementDescriptor, err := flattenPlainValue(rawValueStatementDescriptor, types.StringType, "statement_descriptor", "raw response"); err != nil {
				return err
			} else {
				if typedStatementDescriptor, ok := valueStatementDescriptor.(types.String); ok {
					state.StatementDescriptor = typedStatementDescriptor
				}
			}
		} else if !hasRaw {
			if responseValueStatementDescriptor, ok := plainFromResponseField(obj, "StatementDescriptor"); ok {
				if valueStatementDescriptor, err := flattenPlainValue(responseValueStatementDescriptor, types.StringType, "statement_descriptor", "response struct"); err != nil {
					return err
				} else {
					if typedStatementDescriptor, ok := valueStatementDescriptor.(types.String); ok {
						state.StatementDescriptor = typedStatementDescriptor
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
				if valueStatusDetails, err := flattenPlainValue(sourceStatusDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"failed": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}, "processing": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}, "returned": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}}}, "status_details", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"failed": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}, "processing": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}, "returned": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}}},
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
			if nullStatusDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"failed": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}, "processing": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}, "returned": types.ObjectType{AttrTypes: map[string]attr.Type{"reason": types.StringType}}}}); ok {
				if typedStatusDetails, ok := nullStatusDetails.(types.Object); ok {
					state.StatusDetails = typedStatusDetails
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
				if valueStatusTransitions, err := flattenPlainValue(sourceStatusTransitions, types.ObjectType{AttrTypes: map[string]attr.Type{"canceled_at": types.StringType, "failed_at": types.StringType, "posted_at": types.StringType, "returned_at": types.StringType}}, "status_transitions", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"canceled_at": types.StringType, "failed_at": types.StringType, "posted_at": types.StringType, "returned_at": types.StringType}},
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
			if nullStatusTransitions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"canceled_at": types.StringType, "failed_at": types.StringType, "posted_at": types.StringType, "returned_at": types.StringType}}); ok {
				if typedStatusTransitions, ok := nullStatusTransitions.(types.Object); ok {
					state.StatusTransitions = typedStatusTransitions
				}
			}
		}
	}
	{
		assignedTo := false
		hadRawTo := false
		if rawValueTo, rawOk := plainValueAtPath(raw, "to"); rawOk {
			hadRawTo = true
			if rawValueTo != nil {
				sourceTo := applyConfiguredKeyedListShapes(rawValueTo, attrValueToPlain(state.To))
				if valueTo, err := flattenPlainValue(sourceTo, types.ObjectType{AttrTypes: map[string]attr.Type{"credited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "payout_method": types.StringType, "payout_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"submission": types.StringType, "transaction_purpose": types.StringType}}}}, "preferred_networks": types.ListType{ElemType: types.StringType}}}}}, "recipient": types.StringType, "currency": types.StringType}}, "to", "raw response"); err != nil {
					return err
				} else {
					if typedTo, ok := valueTo.(types.Object); ok {
						state.To = typedTo
						assignedTo = true
					}
				}
			}
		}
		if !assignedTo {
			if !hasRaw {
				if responseValueTo, ok := plainFromResponseField(obj, "To"); ok {
					sourceTo := applyConfiguredKeyedListShapes(responseValueTo, attrValueToPlain(state.To))
					if valueTo, err := flattenPlainValue(
						sourceTo,
						types.ObjectType{AttrTypes: map[string]attr.Type{"credited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "payout_method": types.StringType, "payout_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"submission": types.StringType, "transaction_purpose": types.StringType}}}}, "preferred_networks": types.ListType{ElemType: types.StringType}}}}}, "recipient": types.StringType, "currency": types.StringType}},
						"to",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTo, ok := valueTo.(types.Object); ok {
							state.To = typedTo
							assignedTo = true
						}
					}
				}
			}
		}
		if !assignedTo && hadRawTo {
			if nullTo, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"credited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "payout_method": types.StringType, "payout_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"submission": types.StringType, "transaction_purpose": types.StringType}}}}, "preferred_networks": types.ListType{ElemType: types.StringType}}}}}, "recipient": types.StringType, "currency": types.StringType}}); ok {
				if typedTo, ok := nullTo.(types.Object); ok {
					state.To = typedTo
				}
			}
		}
	}
	{
		assignedTrackingDetails := false
		hadRawTrackingDetails := false
		if rawValueTrackingDetails, rawOk := plainValueAtPath(raw, "tracking_details"); rawOk {
			hadRawTrackingDetails = true
			if rawValueTrackingDetails != nil {
				sourceTrackingDetails := applyConfiguredKeyedListShapes(rawValueTrackingDetails, attrValueToPlain(state.TrackingDetails))
				if valueTrackingDetails, err := flattenPlainValue(sourceTrackingDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"trace_id": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"retrieval_reference": types.StringType}}, "status": types.StringType, "type": types.StringType}}, "tracking_details", "raw response"); err != nil {
					return err
				} else {
					if typedTrackingDetails, ok := valueTrackingDetails.(types.Object); ok {
						state.TrackingDetails = typedTrackingDetails
						assignedTrackingDetails = true
					}
				}
			}
		}
		if !assignedTrackingDetails {
			if !hasRaw {
				if responseValueTrackingDetails, ok := plainFromResponseField(obj, "TrackingDetails"); ok {
					sourceTrackingDetails := applyConfiguredKeyedListShapes(responseValueTrackingDetails, attrValueToPlain(state.TrackingDetails))
					if valueTrackingDetails, err := flattenPlainValue(
						sourceTrackingDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"trace_id": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"retrieval_reference": types.StringType}}, "status": types.StringType, "type": types.StringType}},
						"tracking_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTrackingDetails, ok := valueTrackingDetails.(types.Object); ok {
							state.TrackingDetails = typedTrackingDetails
							assignedTrackingDetails = true
						}
					}
				}
			}
		}
		if !assignedTrackingDetails && hadRawTrackingDetails {
			if nullTrackingDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"trace_id": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"retrieval_reference": types.StringType}}, "status": types.StringType, "type": types.StringType}}); ok {
				if typedTrackingDetails, ok := nullTrackingDetails.(types.Object); ok {
					state.TrackingDetails = typedTrackingDetails
				}
			}
		}
	}
	return nil
}
