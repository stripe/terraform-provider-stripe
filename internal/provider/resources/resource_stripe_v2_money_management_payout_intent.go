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

var _ resource.Resource = &V2MoneyManagementPayoutIntentResource{}

var _ resource.ResourceWithConfigure = &V2MoneyManagementPayoutIntentResource{}

var _ resource.ResourceWithImportState = &V2MoneyManagementPayoutIntentResource{}

func NewV2MoneyManagementPayoutIntentResource() resource.Resource {
	return &V2MoneyManagementPayoutIntentResource{}
}

type V2MoneyManagementPayoutIntentResource struct {
	client *stripe.Client
}

type V2MoneyManagementPayoutIntentResourceModel struct {
	Object                types.String `tfsdk:"object"`
	Amount                types.Object `tfsdk:"amount"`
	ConfirmationMethod    types.String `tfsdk:"confirmation_method"`
	Created               types.String `tfsdk:"created"`
	Description           types.String `tfsdk:"description"`
	EstimatedFees         types.List   `tfsdk:"estimated_fees"`
	From                  types.Object `tfsdk:"from"`
	FxQuote               types.Object `tfsdk:"fx_quote"`
	ID                    types.String `tfsdk:"id"`
	LatestPayout          types.Object `tfsdk:"latest_payout"`
	Livemode              types.Bool   `tfsdk:"livemode"`
	Metadata              types.Map    `tfsdk:"metadata"`
	NextAction            types.Object `tfsdk:"next_action"`
	RecipientNotification types.Object `tfsdk:"recipient_notification"`
	RecipientVerification types.Object `tfsdk:"recipient_verification"`
	ScheduleOptions       types.Object `tfsdk:"schedule_options"`
	StatementDescriptor   types.String `tfsdk:"statement_descriptor"`
	Status                types.String `tfsdk:"status"`
	StatusTransitions     types.Object `tfsdk:"status_transitions"`
	To                    types.Object `tfsdk:"to"`
}

func (r *V2MoneyManagementPayoutIntentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *V2MoneyManagementPayoutIntentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v2_money_management_payout_intent"
}

func (r *V2MoneyManagementPayoutIntentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "PayoutIntent represents an intent to send funds from a Financial Account to a payout method.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value of the object field.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("v2.money_management.payout_intent")},
			},
			"amount": schema.SingleNestedAttribute{
				Required:      true,
				Description:   "The monetary amount to be sent.",
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
			"confirmation_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Controls whether the PayoutIntent requires explicit confirmation before transitioning to `pending`.\nEven when `automatic`, regulatory requirements may require confirmation. See [Send money with Global Payouts](https://docs.stripe.com/global-payouts/send-money) for details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("automatic", "manual")},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				Description:   "Time at which the PayoutIntent was created.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the PayoutIntent. Often useful for displaying to users.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"estimated_fees": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "Estimated fees and taxes.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"amount": schema.SingleNestedAttribute{
							Computed:      true,
							Description:   "The fee amount.",
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
						"network_fee_details": schema.SingleNestedAttribute{
							Computed:      true,
							Description:   "Details about the network and options associated with this fee. Present when `type` is `network_fee`.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"network": schema.StringAttribute{
									Computed:      true,
									Description:   "The network associated with the fee.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									Validators:    []validator.String{stringvalidator.OneOf("ach", "fedwire", "fps", "local", "rtp", "sepa", "swift")},
								},
								"network_options": schema.SingleNestedAttribute{
									Computed:      true,
									Description:   "Per-network options that affect the fee.",
									PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
									Attributes: map[string]schema.Attribute{
										"ach": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "ACH-specific network fee options.",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"submission": schema.StringAttribute{
													Computed:      true,
													Description:   "Open Enum. ACH submission timing.",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
													Validators:    []validator.String{stringvalidator.OneOf("next_day", "same_day")},
												},
											},
										},
									},
								},
							},
						},
						"tax_amount": schema.SingleNestedAttribute{
							Computed:      true,
							Description:   "Tax charged for this fee, if applicable. Value expressed as a decimal string in major units.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"currency": schema.StringAttribute{
									Computed:      true,
									Description:   "Currency code.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"value_decimal": schema.StringAttribute{
									Computed:      true,
									Description:   "Tax amount value represented as a decimal string in major units.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), equivalentDecimalStringPlanModifier()},
								},
							},
						},
						"type": schema.StringAttribute{
							Computed:      true,
							Description:   "Open Enum. The type of fee.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("cross_border_fee", "foreign_exchange_fee", "instant_card_payout_fee", "network_fee", "next_day_payout_fee", "real_time_payout_fee", "standard_payout_fee", "wire_payout_fee")},
						},
					},
				},
			},
			"from": schema.SingleNestedAttribute{
				Required:      true,
				Description:   "The FinancialAccount that funds are pulled from.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"currency": schema.StringAttribute{
						Required:      true,
						Description:   "The currency of the financial account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"debited": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Estimated amount to be debited from the financial account.",
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
						Description:   "The FinancialAccount that funds are pulled from.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"fx_quote": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "FX rate information for fee transparency.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"lock_duration": schema.StringAttribute{
						Computed:      true,
						Description:   "Open Enum. Duration of the FX rate lock.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("five_minutes", "none")},
					},
					"lock_expires_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp when the rate lock expires. Null when rate locking is not supported.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"lock_status": schema.StringAttribute{
						Computed:      true,
						Description:   "Open Enum. Lock status of the FX rate.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("active", "expired", "none")},
					},
					"rates": schema.MapNestedAttribute{
						Computed:      true,
						Description:   "Maps each source currency to its exchange rate into `to_currency`.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"exchange_rate": schema.StringAttribute{
									Computed:      true,
									Description:   "The exchange rate from the source currency to `to_currency`, represented as a decimal string\n(for example, `1.1520`) to preserve the full precision of the rate.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
							},
						},
					},
					"to_currency": schema.StringAttribute{
						Computed:      true,
						Description:   "The destination currency.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the PayoutIntent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"latest_payout": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Details about the latest payout associated with this PayoutIntent.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"outbound_payment": schema.StringAttribute{
						Computed:      true,
						Description:   "The ID of the OutboundPayment, if applicable.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"outbound_transfer": schema.StringAttribute{
						Computed:      true,
						Description:   "The ID of the OutboundTransfer, if applicable.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "The type of payout.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("outbound_payment", "outbound_transfer")},
					},
				},
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
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"next_action": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Next action required for a PayoutIntent in the `requires_action` state. Populated when `status` is `requires_action`.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"confirm": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details about a confirmation required. Populated when `type` is `confirm`.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fields_to_review": schema.ListAttribute{
								Computed:      true,
								Description:   "The PayoutIntent fields that must be reviewed before confirmation.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"reason": schema.StringAttribute{
								Computed:      true,
								Description:   "Open Enum. The reason the PayoutIntent requires confirmation.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("automatically_required", "manually_requested")},
							},
						},
					},
					"handle_failure": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details about a failure that requires user action. Populated when `type` is `handle_failure`.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"failure_reason": schema.StringAttribute{
								Computed:      true,
								Description:   "Open Enum. The reason for the failure.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("account_not_configured_as_recipient", "fx_rate_drift_exceeded_after_review", "insufficient_funds", "payout_method_account_type_incorrect", "payout_method_amount_limit_exceeded", "payout_method_canceled_by_customer", "payout_method_closed", "payout_method_currency_unsupported", "payout_method_declined", "payout_method_does_not_exist", "payout_method_expired", "payout_method_holder_address_incorrect", "payout_method_holder_details_incorrect", "payout_method_holder_name_incorrect", "payout_method_invalid", "payout_method_invalid_account_number", "payout_method_restricted", "payout_method_unsupported", "payout_method_usage_frequency_limit_exceeded", "recalled", "recipient_capability_not_active", "review_rejected", "unknown_failure")},
							},
						},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "Open Enum. The type of next action required.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("confirm", "handle_failure")},
					},
				},
			},
			"recipient_notification": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Recipient notification settings for the PayoutIntent. Only applicable when the PayoutIntent creates an OutboundPayment.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"setting": schema.StringAttribute{
						Required:    true,
						Description: "Closed Enum. Configuration option to enable or disable notifications to recipients.\nDo not send notifications when `setting` is `none`. Use the account setting when `setting` is `configured` or omitted.",
						Validators:  []validator.String{stringvalidator.OneOf("configured", "none")},
					},
				},
			},
			"recipient_verification": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Recipient verification details for this PayoutIntent.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"match_result": schema.StringAttribute{
						Computed:      true,
						Description:   "Closed Enum. The recipient-verification provider's match result.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("close_match", "match", "no_match", "unavailable")},
					},
					"match_result_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details associated with the match result.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"matched_business_type": schema.StringAttribute{
								Computed:      true,
								Description:   "The business type returned by the recipient-verification provider.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("business", "personal")},
							},
							"matched_name": schema.StringAttribute{
								Computed:      true,
								Description:   "The account name returned by the recipient-verification provider.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"message": schema.StringAttribute{
								Computed:      true,
								Description:   "A message describing the match result.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"provided_business_type": schema.StringAttribute{
								Computed:      true,
								Description:   "The business type provided for recipient verification.",
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
				},
			},
			"schedule_options": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Scheduling options for the payout. If omitted, the payout is scheduled for immediate execution.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"execute_on": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The date when the payout should be executed, in YYYY-MM-DD format.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"statement_descriptor": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The description that appears on the receiving end for the payout (for example, on a bank statement).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Current status of the PayoutIntent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("canceled", "pending", "posted", "processing", "requires_action")},
			},
			"status_transitions": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Hash containing timestamps of when transitioned to a particular status.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"canceled_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp describing when a PayoutIntent changed status to `canceled`.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"posted_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp describing when a PayoutIntent changed status to `posted`.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"processing_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp describing when a PayoutIntent changed status to `processing`.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"requires_action_at": schema.StringAttribute{
						Computed:      true,
						Description:   "Timestamp describing when a PayoutIntent changed status to `requires_action`.\nRepresented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"to": schema.SingleNestedAttribute{
				Required:    true,
				Description: "To which payout method the payout is sent.",
				Attributes: map[string]schema.Attribute{
					"credited": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Estimated amount to be credited to the recipient in the destination currency.",
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
					"currency": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The currency to send to the recipient.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"payout_method": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The payout method ID for the PayoutIntent. Optional when the PayoutIntent creates an OutboundPayment and the recipient has a default payout method. Required when the PayoutIntent creates an OutboundTransfer.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"payout_method_options": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Payout method options for the PayoutIntent.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank_account": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Options for bank account payout methods.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"preferred_network_options": schema.SingleNestedAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Per-network configuration options.",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"ach": schema.SingleNestedAttribute{
												Optional:      true,
												Computed:      true,
												Description:   "ACH-specific network options.",
												PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
												Attributes: map[string]schema.Attribute{
													"submission": schema.StringAttribute{
														Optional:      true,
														Computed:      true,
														Description:   "Open Enum. ACH submission timing.",
														PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														Validators:    []validator.String{stringvalidator.OneOf("next_day", "same_day")},
													},
													"transaction_purpose": schema.StringAttribute{
														Optional:      true,
														Computed:      true,
														Description:   "The transaction purpose for this ACH payment.",
														PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														Validators:    []validator.String{stringvalidator.OneOf("payroll")},
													},
												},
											},
										},
									},
									"preferred_networks": schema.ListAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The preferred networks to use for this PayoutIntent.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
								},
							},
						},
					},
					"recipient": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The recipient ID for the PayoutIntent.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
		},
	}
}

func (r *V2MoneyManagementPayoutIntentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan V2MoneyManagementPayoutIntentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandV2MoneyManagementPayoutIntentCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2MoneyManagementPayoutIntent create params", err.Error())
		return
	}

	obj, err := r.client.V2MoneyManagementPayoutIntents.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating V2MoneyManagementPayoutIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementPayoutIntents.B, r.client.V2MoneyManagementPayoutIntents.Key, stripe.FormatURLPath("/v2/money_management/payout_intents/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementPayoutIntent create raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementPayoutIntent(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementPayoutIntent create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementPayoutIntentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState V2MoneyManagementPayoutIntentResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2MoneyManagementPayoutIntentResourceModel
	state = priorState

	obj, err := r.client.V2MoneyManagementPayoutIntents.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading V2MoneyManagementPayoutIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementPayoutIntents.B, r.client.V2MoneyManagementPayoutIntents.Key, stripe.FormatURLPath("/v2/money_management/payout_intents/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementPayoutIntent raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementPayoutIntent(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementPayoutIntent read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementPayoutIntentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan V2MoneyManagementPayoutIntentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2MoneyManagementPayoutIntentResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandV2MoneyManagementPayoutIntentUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2MoneyManagementPayoutIntent update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building V2MoneyManagementPayoutIntent update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V2MoneyManagementPayoutIntents.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating V2MoneyManagementPayoutIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementPayoutIntents.B, r.client.V2MoneyManagementPayoutIntents.Key, stripe.FormatURLPath("/v2/money_management/payout_intents/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementPayoutIntent update raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementPayoutIntent(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementPayoutIntent update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementPayoutIntentResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *V2MoneyManagementPayoutIntentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandV2MoneyManagementPayoutIntentCreate(plan V2MoneyManagementPayoutIntentResourceModel) (*stripe.V2MoneyManagementPayoutIntentCreateParams, error) {
	params := &stripe.V2MoneyManagementPayoutIntentCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Amount", plan.Amount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "amount", params)
		}
	}
	if !plan.ConfirmationMethod.IsNull() && !plan.ConfirmationMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ConfirmationMethod", "ConfirmationMethod", plan.ConfirmationMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "confirmation_method", params)
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
	if !plan.ScheduleOptions.IsNull() && !plan.ScheduleOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ScheduleOptions", plan.ScheduleOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "schedule_options", params)
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

func expandV2MoneyManagementPayoutIntentUpdate(plan V2MoneyManagementPayoutIntentResourceModel, state V2MoneyManagementPayoutIntentResourceModel) (*stripe.V2MoneyManagementPayoutIntentUpdateParams, error) {
	params := &stripe.V2MoneyManagementPayoutIntentUpdateParams{}

	if !plan.Description.Equal(state.Description) && !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			if !plan.Description.Equal(state.Description) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "description", params)
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
	if !plan.RecipientNotification.Equal(state.RecipientNotification) && !plan.RecipientNotification.IsNull() && !plan.RecipientNotification.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RecipientNotification", plan.RecipientNotification) {
			if !plan.RecipientNotification.Equal(state.RecipientNotification) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "recipient_notification", params)
			}
		}
	}
	if !plan.ScheduleOptions.Equal(state.ScheduleOptions) && !plan.ScheduleOptions.IsNull() && !plan.ScheduleOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ScheduleOptions", plan.ScheduleOptions) {
			if !plan.ScheduleOptions.Equal(state.ScheduleOptions) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "schedule_options", params)
			}
		}
	}
	if !plan.StatementDescriptor.Equal(state.StatementDescriptor) && !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			if !plan.StatementDescriptor.Equal(state.StatementDescriptor) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "statement_descriptor", params)
			}
		}
	}
	if !plan.To.Equal(state.To) && !plan.To.IsNull() && !plan.To.IsUnknown() {
		if !assignAttrValueToNamedField(params, "To", plan.To) {
			if !plan.To.Equal(state.To) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "to", params)
			}
		}
	}

	return params, nil
}

func flattenV2MoneyManagementPayoutIntent(obj *stripe.V2MoneyManagementPayoutIntent, state *V2MoneyManagementPayoutIntentResourceModel) error {
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
		if rawValueConfirmationMethod, rawOk := plainValueAtPath(raw, "confirmation_method"); rawOk {
			if valueConfirmationMethod, err := flattenPlainValue(rawValueConfirmationMethod, types.StringType, "confirmation_method", "raw response"); err != nil {
				return err
			} else {
				if typedConfirmationMethod, ok := valueConfirmationMethod.(types.String); ok {
					state.ConfirmationMethod = typedConfirmationMethod
				}
			}
		} else if !hasRaw {
			if responseValueConfirmationMethod, ok := plainFromResponseField(obj, "ConfirmationMethod"); ok {
				if valueConfirmationMethod, err := flattenPlainValue(responseValueConfirmationMethod, types.StringType, "confirmation_method", "response struct"); err != nil {
					return err
				} else {
					if typedConfirmationMethod, ok := valueConfirmationMethod.(types.String); ok {
						state.ConfirmationMethod = typedConfirmationMethod
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
		if rawValueEstimatedFees, rawOk := plainValueAtPath(raw, "estimated_fees"); rawOk {
			if valueEstimatedFees, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawValueEstimatedFees, attrValueToPlain(state.EstimatedFees)), attrValueToPlain(state.EstimatedFees)), attrValueToPlain(state.EstimatedFees)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "network_fee_details": types.ObjectType{AttrTypes: map[string]attr.Type{"network": types.StringType, "network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"submission": types.StringType}}}}}}, "tax_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value_decimal": types.StringType}}, "type": types.StringType}}}, "estimated_fees", "raw response"); err != nil {
				return err
			} else {
				if typedEstimatedFees, ok := valueEstimatedFees.(types.List); ok {
					state.EstimatedFees = typedEstimatedFees
				}
			}
		} else if !hasRaw {
			if responseValueEstimatedFees, ok := plainFromResponseField(obj, "EstimatedFees"); ok {
				if valueEstimatedFees, err := flattenPlainValue(
					preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(responseValueEstimatedFees, attrValueToPlain(state.EstimatedFees)), attrValueToPlain(state.EstimatedFees)), attrValueToPlain(state.EstimatedFees)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "network_fee_details": types.ObjectType{AttrTypes: map[string]attr.Type{"network": types.StringType, "network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"submission": types.StringType}}}}}}, "tax_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value_decimal": types.StringType}}, "type": types.StringType}}},
					"estimated_fees",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedEstimatedFees, ok := valueEstimatedFees.(types.List); ok {
						state.EstimatedFees = typedEstimatedFees
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
				if valueFrom, err := flattenPlainValue(sourceFrom, types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "debited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "financial_account": types.StringType}}, "from", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "debited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "financial_account": types.StringType}},
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
			if nullFrom, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "debited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "financial_account": types.StringType}}); ok {
				if typedFrom, ok := nullFrom.(types.Object); ok {
					state.From = typedFrom
				}
			}
		}
	}
	{
		assignedFxQuote := false
		hadRawFxQuote := false
		if rawValueFxQuote, rawOk := plainValueAtPath(raw, "fx_quote"); rawOk {
			hadRawFxQuote = true
			if rawValueFxQuote != nil {
				sourceFxQuote := applyConfiguredKeyedListShapes(rawValueFxQuote, attrValueToPlain(state.FxQuote))
				if valueFxQuote, err := flattenPlainValue(sourceFxQuote, types.ObjectType{AttrTypes: map[string]attr.Type{"lock_duration": types.StringType, "lock_expires_at": types.StringType, "lock_status": types.StringType, "rates": types.MapType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"exchange_rate": types.StringType}}}, "to_currency": types.StringType}}, "fx_quote", "raw response"); err != nil {
					return err
				} else {
					if typedFxQuote, ok := valueFxQuote.(types.Object); ok {
						state.FxQuote = typedFxQuote
						assignedFxQuote = true
					}
				}
			}
		}
		if !assignedFxQuote {
			if !hasRaw {
				if responseValueFxQuote, ok := plainFromResponseField(obj, "FxQuote"); ok {
					sourceFxQuote := applyConfiguredKeyedListShapes(responseValueFxQuote, attrValueToPlain(state.FxQuote))
					if valueFxQuote, err := flattenPlainValue(
						sourceFxQuote,
						types.ObjectType{AttrTypes: map[string]attr.Type{"lock_duration": types.StringType, "lock_expires_at": types.StringType, "lock_status": types.StringType, "rates": types.MapType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"exchange_rate": types.StringType}}}, "to_currency": types.StringType}},
						"fx_quote",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFxQuote, ok := valueFxQuote.(types.Object); ok {
							state.FxQuote = typedFxQuote
							assignedFxQuote = true
						}
					}
				}
			}
		}
		if !assignedFxQuote && hadRawFxQuote {
			if nullFxQuote, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"lock_duration": types.StringType, "lock_expires_at": types.StringType, "lock_status": types.StringType, "rates": types.MapType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"exchange_rate": types.StringType}}}, "to_currency": types.StringType}}); ok {
				if typedFxQuote, ok := nullFxQuote.(types.Object); ok {
					state.FxQuote = typedFxQuote
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
		assignedLatestPayout := false
		hadRawLatestPayout := false
		if rawValueLatestPayout, rawOk := plainValueAtPath(raw, "latest_payout"); rawOk {
			hadRawLatestPayout = true
			if rawValueLatestPayout != nil {
				sourceLatestPayout := applyConfiguredKeyedListShapes(rawValueLatestPayout, attrValueToPlain(state.LatestPayout))
				if valueLatestPayout, err := flattenPlainValue(sourceLatestPayout, types.ObjectType{AttrTypes: map[string]attr.Type{"outbound_payment": types.StringType, "outbound_transfer": types.StringType, "type": types.StringType}}, "latest_payout", "raw response"); err != nil {
					return err
				} else {
					if typedLatestPayout, ok := valueLatestPayout.(types.Object); ok {
						state.LatestPayout = typedLatestPayout
						assignedLatestPayout = true
					}
				}
			}
		}
		if !assignedLatestPayout {
			if !hasRaw {
				if responseValueLatestPayout, ok := plainFromResponseField(obj, "LatestPayout"); ok {
					sourceLatestPayout := applyConfiguredKeyedListShapes(responseValueLatestPayout, attrValueToPlain(state.LatestPayout))
					if valueLatestPayout, err := flattenPlainValue(
						sourceLatestPayout,
						types.ObjectType{AttrTypes: map[string]attr.Type{"outbound_payment": types.StringType, "outbound_transfer": types.StringType, "type": types.StringType}},
						"latest_payout",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLatestPayout, ok := valueLatestPayout.(types.Object); ok {
							state.LatestPayout = typedLatestPayout
							assignedLatestPayout = true
						}
					}
				}
			}
		}
		if !assignedLatestPayout && hadRawLatestPayout {
			if nullLatestPayout, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"outbound_payment": types.StringType, "outbound_transfer": types.StringType, "type": types.StringType}}); ok {
				if typedLatestPayout, ok := nullLatestPayout.(types.Object); ok {
					state.LatestPayout = typedLatestPayout
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
		assignedNextAction := false
		hadRawNextAction := false
		if rawValueNextAction, rawOk := plainValueAtPath(raw, "next_action"); rawOk {
			hadRawNextAction = true
			if rawValueNextAction != nil {
				sourceNextAction := applyConfiguredKeyedListShapes(rawValueNextAction, attrValueToPlain(state.NextAction))
				if valueNextAction, err := flattenPlainValue(sourceNextAction, types.ObjectType{AttrTypes: map[string]attr.Type{"confirm": types.ObjectType{AttrTypes: map[string]attr.Type{"fields_to_review": types.ListType{ElemType: types.StringType}, "reason": types.StringType}}, "handle_failure": types.ObjectType{AttrTypes: map[string]attr.Type{"failure_reason": types.StringType}}, "type": types.StringType}}, "next_action", "raw response"); err != nil {
					return err
				} else {
					if typedNextAction, ok := valueNextAction.(types.Object); ok {
						state.NextAction = typedNextAction
						assignedNextAction = true
					}
				}
			}
		}
		if !assignedNextAction {
			if !hasRaw {
				if responseValueNextAction, ok := plainFromResponseField(obj, "NextAction"); ok {
					sourceNextAction := applyConfiguredKeyedListShapes(responseValueNextAction, attrValueToPlain(state.NextAction))
					if valueNextAction, err := flattenPlainValue(
						sourceNextAction,
						types.ObjectType{AttrTypes: map[string]attr.Type{"confirm": types.ObjectType{AttrTypes: map[string]attr.Type{"fields_to_review": types.ListType{ElemType: types.StringType}, "reason": types.StringType}}, "handle_failure": types.ObjectType{AttrTypes: map[string]attr.Type{"failure_reason": types.StringType}}, "type": types.StringType}},
						"next_action",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedNextAction, ok := valueNextAction.(types.Object); ok {
							state.NextAction = typedNextAction
							assignedNextAction = true
						}
					}
				}
			}
		}
		if !assignedNextAction && hadRawNextAction {
			if nullNextAction, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"confirm": types.ObjectType{AttrTypes: map[string]attr.Type{"fields_to_review": types.ListType{ElemType: types.StringType}, "reason": types.StringType}}, "handle_failure": types.ObjectType{AttrTypes: map[string]attr.Type{"failure_reason": types.StringType}}, "type": types.StringType}}); ok {
				if typedNextAction, ok := nullNextAction.(types.Object); ok {
					state.NextAction = typedNextAction
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
		assignedRecipientVerification := false
		hadRawRecipientVerification := false
		if rawValueRecipientVerification, rawOk := plainValueAtPath(raw, "recipient_verification"); rawOk {
			hadRawRecipientVerification = true
			if rawValueRecipientVerification != nil {
				sourceRecipientVerification := applyConfiguredKeyedListShapes(rawValueRecipientVerification, attrValueToPlain(state.RecipientVerification))
				if valueRecipientVerification, err := flattenPlainValue(sourceRecipientVerification, types.ObjectType{AttrTypes: map[string]attr.Type{"match_result": types.StringType, "match_result_details": types.ObjectType{AttrTypes: map[string]attr.Type{"matched_business_type": types.StringType, "matched_name": types.StringType, "message": types.StringType, "provided_business_type": types.StringType, "provided_name": types.StringType}}}}, "recipient_verification", "raw response"); err != nil {
					return err
				} else {
					if typedRecipientVerification, ok := valueRecipientVerification.(types.Object); ok {
						state.RecipientVerification = typedRecipientVerification
						assignedRecipientVerification = true
					}
				}
			}
		}
		if !assignedRecipientVerification {
			if !hasRaw {
				if responseValueRecipientVerification, ok := plainFromResponseField(obj, "RecipientVerification"); ok {
					sourceRecipientVerification := applyConfiguredKeyedListShapes(responseValueRecipientVerification, attrValueToPlain(state.RecipientVerification))
					if valueRecipientVerification, err := flattenPlainValue(
						sourceRecipientVerification,
						types.ObjectType{AttrTypes: map[string]attr.Type{"match_result": types.StringType, "match_result_details": types.ObjectType{AttrTypes: map[string]attr.Type{"matched_business_type": types.StringType, "matched_name": types.StringType, "message": types.StringType, "provided_business_type": types.StringType, "provided_name": types.StringType}}}},
						"recipient_verification",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRecipientVerification, ok := valueRecipientVerification.(types.Object); ok {
							state.RecipientVerification = typedRecipientVerification
							assignedRecipientVerification = true
						}
					}
				}
			}
		}
		if !assignedRecipientVerification && hadRawRecipientVerification {
			if nullRecipientVerification, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"match_result": types.StringType, "match_result_details": types.ObjectType{AttrTypes: map[string]attr.Type{"matched_business_type": types.StringType, "matched_name": types.StringType, "message": types.StringType, "provided_business_type": types.StringType, "provided_name": types.StringType}}}}); ok {
				if typedRecipientVerification, ok := nullRecipientVerification.(types.Object); ok {
					state.RecipientVerification = typedRecipientVerification
				}
			}
		}
	}
	{
		assignedScheduleOptions := false
		hadRawScheduleOptions := false
		if rawValueScheduleOptions, rawOk := plainValueAtPath(raw, "schedule_options"); rawOk {
			hadRawScheduleOptions = true
			if rawValueScheduleOptions != nil {
				sourceScheduleOptions := applyConfiguredKeyedListShapes(rawValueScheduleOptions, attrValueToPlain(state.ScheduleOptions))
				if valueScheduleOptions, err := flattenPlainValue(sourceScheduleOptions, types.ObjectType{AttrTypes: map[string]attr.Type{"execute_on": types.StringType}}, "schedule_options", "raw response"); err != nil {
					return err
				} else {
					if typedScheduleOptions, ok := valueScheduleOptions.(types.Object); ok {
						state.ScheduleOptions = typedScheduleOptions
						assignedScheduleOptions = true
					}
				}
			}
		}
		if !assignedScheduleOptions {
			if !hasRaw {
				if responseValueScheduleOptions, ok := plainFromResponseField(obj, "ScheduleOptions"); ok {
					sourceScheduleOptions := applyConfiguredKeyedListShapes(responseValueScheduleOptions, attrValueToPlain(state.ScheduleOptions))
					if valueScheduleOptions, err := flattenPlainValue(
						sourceScheduleOptions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"execute_on": types.StringType}},
						"schedule_options",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedScheduleOptions, ok := valueScheduleOptions.(types.Object); ok {
							state.ScheduleOptions = typedScheduleOptions
							assignedScheduleOptions = true
						}
					}
				}
			}
		}
		if !assignedScheduleOptions && hadRawScheduleOptions {
			if nullScheduleOptions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"execute_on": types.StringType}}); ok {
				if typedScheduleOptions, ok := nullScheduleOptions.(types.Object); ok {
					state.ScheduleOptions = typedScheduleOptions
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
		assignedStatusTransitions := false
		hadRawStatusTransitions := false
		if rawValueStatusTransitions, rawOk := plainValueAtPath(raw, "status_transitions"); rawOk {
			hadRawStatusTransitions = true
			if rawValueStatusTransitions != nil {
				sourceStatusTransitions := applyConfiguredKeyedListShapes(rawValueStatusTransitions, attrValueToPlain(state.StatusTransitions))
				if valueStatusTransitions, err := flattenPlainValue(sourceStatusTransitions, types.ObjectType{AttrTypes: map[string]attr.Type{"canceled_at": types.StringType, "posted_at": types.StringType, "processing_at": types.StringType, "requires_action_at": types.StringType}}, "status_transitions", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"canceled_at": types.StringType, "posted_at": types.StringType, "processing_at": types.StringType, "requires_action_at": types.StringType}},
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
			if nullStatusTransitions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"canceled_at": types.StringType, "posted_at": types.StringType, "processing_at": types.StringType, "requires_action_at": types.StringType}}); ok {
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
				if valueTo, err := flattenPlainValue(sourceTo, types.ObjectType{AttrTypes: map[string]attr.Type{"credited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "currency": types.StringType, "payout_method": types.StringType, "payout_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"submission": types.StringType, "transaction_purpose": types.StringType}}}}, "preferred_networks": types.ListType{ElemType: types.StringType}}}}}, "recipient": types.StringType}}, "to", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"credited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "currency": types.StringType, "payout_method": types.StringType, "payout_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"submission": types.StringType, "transaction_purpose": types.StringType}}}}, "preferred_networks": types.ListType{ElemType: types.StringType}}}}}, "recipient": types.StringType}},
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
			if nullTo, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"credited": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "currency": types.StringType, "payout_method": types.StringType, "payout_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"submission": types.StringType, "transaction_purpose": types.StringType}}}}, "preferred_networks": types.ListType{ElemType: types.StringType}}}}}, "recipient": types.StringType}}); ok {
				if typedTo, ok := nullTo.(types.Object); ok {
					state.To = typedTo
				}
			}
		}
	}
	return nil
}
