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

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &TerminalReaderResource{}

var _ resource.ResourceWithConfigure = &TerminalReaderResource{}

var _ resource.ResourceWithImportState = &TerminalReaderResource{}

func NewTerminalReaderResource() resource.Resource {
	return &TerminalReaderResource{}
}

type TerminalReaderResource struct {
	client *stripe.Client
}

type TerminalReaderResourceModel struct {
	Object           types.String `tfsdk:"object"`
	Action           types.Object `tfsdk:"action"`
	DeviceSwVersion  types.String `tfsdk:"device_sw_version"`
	DeviceType       types.String `tfsdk:"device_type"`
	ID               types.String `tfsdk:"id"`
	IPAddress        types.String `tfsdk:"ip_address"`
	Label            types.String `tfsdk:"label"`
	Livemode         types.Bool   `tfsdk:"livemode"`
	Location         types.String `tfsdk:"location"`
	Metadata         types.Map    `tfsdk:"metadata"`
	SerialNumber     types.String `tfsdk:"serial_number"`
	Status           types.String `tfsdk:"status"`
	RegistrationCode types.String `tfsdk:"registration_code"`
}

func (r *TerminalReaderResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TerminalReaderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_terminal_reader"
}

func (r *TerminalReaderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Reader represents a physical device for accepting payment details.\n\nRelated guide: [Connecting to a reader](https://docs.stripe.com/terminal/payments/connect-reader)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("terminal.reader")},
			},
			"action": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The most recent action performed by the reader.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"api_error": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "The reader action failed due to an [API error](https://docs.stripe.com/api/errors). Only present when `status` is `failed` and the underlying failure was an API error. Avoid parsing the `message` field for programmatic logic; use `type` or `code` instead. The `message` field is for display to humans only and may be updated at anytime. Requires [reader version](https://docs.stripe.com/terminal/readers/stripe-reader-s700-s710#reader-software-version) 2.42 or later. Readers on older versions always return null.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"advice_code": schema.StringAttribute{
								Computed:      true,
								Description:   "For card errors resulting from a card issuer decline, a short string indicating [how to proceed with an error](https://docs.stripe.com/declines#retrying-issuer-declines) if they provide one.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"charge": schema.StringAttribute{
								Computed:      true,
								Description:   "For card errors, the ID of the failed charge.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"code": schema.StringAttribute{
								Computed:      true,
								Description:   "For some errors that could be handled programmatically, a short string indicating the [error code](https://docs.stripe.com/error-codes) reported.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("account_closed", "account_country_invalid_address", "account_error_country_change_requires_additional_steps", "account_information_mismatch", "account_invalid", "account_number_invalid", "account_token_required_for_v2_account", "acss_debit_session_incomplete", "action_blocked", "alipay_upgrade_required", "amount_too_large", "amount_too_small", "api_key_expired", "application_fees_not_allowed", "approval_required", "authentication_required", "balance_insufficient", "balance_invalid_parameter", "bank_account_bad_routing_numbers", "bank_account_declined", "bank_account_exists", "bank_account_restricted", "bank_account_unusable", "bank_account_unverified", "bank_account_verification_failed", "billing_invalid_mandate", "bitcoin_upgrade_required", "capture_charge_authorization_expired", "capture_unauthorized_payment", "card_decline_rate_limit_exceeded", "card_declined", "cardholder_phone_number_required", "charge_already_captured", "charge_already_refunded", "charge_disputed", "charge_exceeds_source_limit", "charge_exceeds_transaction_limit", "charge_expired_for_capture", "charge_invalid_parameter", "charge_not_refundable", "clearing_code_unsupported", "country_code_invalid", "country_unsupported", "coupon_expired", "customer_max_payment_methods", "customer_max_subscriptions", "customer_session_expired", "customer_tax_location_invalid", "debit_not_authorized", "email_invalid", "expired_card", "financial_connections_account_inactive", "financial_connections_account_pending_account_numbers", "financial_connections_account_unavailable_account_numbers", "financial_connections_no_successful_transaction_refresh", "forwarding_api_inactive", "forwarding_api_invalid_parameter", "forwarding_api_retryable_upstream_error", "forwarding_api_upstream_connection_error", "forwarding_api_upstream_connection_timeout", "forwarding_api_upstream_error", "idempotency_key_in_use", "incorrect_address", "incorrect_cvc", "incorrect_number", "incorrect_zip", "india_recurring_payment_mandate_canceled", "instant_payouts_config_disabled", "instant_payouts_currency_disabled", "instant_payouts_limit_exceeded", "instant_payouts_unsupported", "insufficient_funds", "intent_invalid_state", "intent_verification_method_missing", "invalid_card_type", "invalid_characters", "invalid_charge_amount", "invalid_cvc", "invalid_expiry_month", "invalid_expiry_year", "invalid_mandate_reference_prefix_format", "invalid_number", "invalid_source_usage", "invalid_tax_location", "invoice_no_customer_line_items", "invoice_no_payment_method_types", "invoice_no_subscription_line_items", "invoice_not_editable", "invoice_on_behalf_of_not_editable", "invoice_payment_intent_requires_action", "invoice_upcoming_none", "livemode_mismatch", "lock_timeout", "missing", "no_account", "not_allowed_on_standard_account", "out_of_inventory", "ownership_declaration_not_allowed", "parameter_invalid_empty", "parameter_invalid_integer", "parameter_invalid_string_blank", "parameter_invalid_string_empty", "parameter_missing", "parameter_unknown", "parameters_exclusive", "payment_intent_action_required", "payment_intent_authentication_failure", "payment_intent_incompatible_payment_method", "payment_intent_invalid_parameter", "payment_intent_konbini_rejected_confirmation_number", "payment_intent_mandate_invalid", "payment_intent_payment_attempt_expired", "payment_intent_payment_attempt_failed", "payment_intent_rate_limit_exceeded", "payment_intent_unexpected_state", "payment_method_bank_account_already_verified", "payment_method_bank_account_blocked", "payment_method_billing_details_address_missing", "payment_method_configuration_failures", "payment_method_currency_mismatch", "payment_method_customer_decline", "payment_method_invalid_parameter", "payment_method_invalid_parameter_testmode", "payment_method_microdeposit_failed", "payment_method_microdeposit_processing_error", "payment_method_microdeposit_verification_amounts_invalid", "payment_method_microdeposit_verification_amounts_mismatch", "payment_method_microdeposit_verification_attempts_exceeded", "payment_method_microdeposit_verification_descriptor_code_mismatch", "payment_method_microdeposit_verification_timeout", "payment_method_not_available", "payment_method_provider_decline", "payment_method_provider_timeout", "payment_method_unactivated", "payment_method_unexpected_state", "payment_method_unsupported_type", "payout_reconciliation_not_ready", "payouts_limit_exceeded", "payouts_not_allowed", "platform_account_required", "platform_api_key_expired", "postal_code_invalid", "processing_error", "product_inactive", "progressive_onboarding_limit_exceeded", "rate_limit", "refer_to_customer", "refund_disputed_payment", "request_blocked", "resource_already_exists", "resource_missing", "return_intent_already_processed", "routing_number_invalid", "secret_key_required", "sepa_unsupported_account", "service_period_coupon_with_metered_tiered_item_unsupported", "setup_attempt_failed", "setup_intent_authentication_failure", "setup_intent_invalid_parameter", "setup_intent_mandate_invalid", "setup_intent_mobile_wallet_unsupported", "setup_intent_setup_attempt_expired", "setup_intent_unexpected_state", "shipping_address_invalid", "shipping_calculation_failed", "siret_invalid", "sku_inactive", "state_unsupported", "status_transition_invalid", "storer_capability_missing", "storer_capability_not_active", "stripe_tax_inactive", "tax_id_invalid", "tax_id_prohibited", "taxes_calculation_failed", "terminal_location_country_unsupported", "terminal_reader_busy", "terminal_reader_hardware_fault", "terminal_reader_invalid_location_for_activation", "terminal_reader_invalid_location_for_payment", "terminal_reader_offline", "terminal_reader_timeout", "testmode_charges_only", "tls_version_unsupported", "token_already_used", "token_card_network_invalid", "token_in_use", "transfer_source_balance_parameters_mismatch", "transfers_not_allowed", "url_invalid")},
							},
							"decline_code": schema.StringAttribute{
								Computed:      true,
								Description:   "For card errors resulting from a card issuer decline, a short string indicating the [card issuer's reason for the decline](https://docs.stripe.com/declines#issuer-declines) if they provide one.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"doc_url": schema.StringAttribute{
								Computed:      true,
								Description:   "A URL to more information about the [error code](https://docs.stripe.com/error-codes) reported.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"message": schema.StringAttribute{
								Computed:      true,
								Description:   "A human-readable message providing more details about the error. For card errors, these messages can be shown to your users.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"network_advice_code": schema.StringAttribute{
								Computed:      true,
								Description:   "For card errors resulting from a card issuer decline, a 2 digit code which indicates the advice given to merchant by the card network on how to proceed with an error.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"network_decline_code": schema.StringAttribute{
								Computed:      true,
								Description:   "For payments declined by the network, an alphanumeric code which indicates the reason the payment failed.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"param": schema.StringAttribute{
								Computed:      true,
								Description:   "If the error is parameter-specific, the parameter related to the error. For example, you can use this to display a message near the correct form field.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payment_intent": schema.StringAttribute{
								Computed:      true,
								Description:   "A PaymentIntent guides you through the process of collecting a payment from your customer.\nWe recommend that you create exactly one PaymentIntent for each order or\ncustomer session in your system. You can reference the PaymentIntent later to\nsee the history of payment attempts for a particular session.\n\nA PaymentIntent transitions through\n[multiple statuses](/payments/paymentintents/lifecycle)\nthroughout its lifetime as it interfaces with Stripe.js to perform\nauthentication flows and ultimately creates at most one successful charge.\n\nRelated guide: [Payment Intents API](https://docs.stripe.com/payments/payment-intents)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payment_method": schema.StringAttribute{
								Computed:      true,
								Description:   "PaymentMethod objects represent your customer's payment instruments.\nYou can use them with [PaymentIntents](https://docs.stripe.com/payments/payment-intents) to collect payments or save them to\nCustomer objects to store instrument details for future payments.\n\nRelated guides: [Payment Methods](https://docs.stripe.com/payments/payment-methods) and [More Payment Scenarios](https://docs.stripe.com/payments/more-payment-scenarios).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payment_method_type": schema.StringAttribute{
								Computed:      true,
								Description:   "If the error is specific to the type of payment method, the payment method type that had a problem. This field is only populated for invoice-related errors.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"request_log_url": schema.StringAttribute{
								Computed:      true,
								Description:   "A URL to the request log entry in your dashboard.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"setup_intent": schema.StringAttribute{
								Computed:      true,
								Description:   "A SetupIntent guides you through the process of setting up and saving a customer's payment credentials for future payments.\nFor example, you can use a SetupIntent to set up and save your customer's card without immediately collecting a payment.\nLater, you can use [PaymentIntents](https://api.stripe.com#payment_intents) to drive the payment flow.\n\nCreate a SetupIntent when you're ready to collect your customer's payment credentials.\nDon't maintain long-lived, unconfirmed SetupIntents because they might not be valid.\nThe SetupIntent transitions through multiple [statuses](https://docs.stripe.com/payments/intents#intent-statuses) as it guides\nyou through the setup process.\n\nSuccessful SetupIntents result in payment credentials that are optimized for future payments.\nFor example, cardholders in [certain regions](https://stripe.com/guides/strong-customer-authentication) might need to be run through\n[Strong Customer Authentication](https://docs.stripe.com/strong-customer-authentication) during payment method collection\nto streamline later [off-session payments](https://docs.stripe.com/payments/setup-intents).\nIf you use the SetupIntent with a [Customer](https://api.stripe.com#setup_intent_object-customer),\nit automatically attaches the resulting payment method to that Customer after successful setup.\nWe recommend using SetupIntents or [setup_future_usage](https://api.stripe.com#payment_intent_object-setup_future_usage) on\nPaymentIntents to save payment methods to prevent saving invalid or unoptimized payment methods.\n\nBy using SetupIntents, you can reduce friction for your customers, even as regulations change over time.\n\nRelated guide: [Setup Intents API](https://docs.stripe.com/payments/setup-intents)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"source": schema.StringAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"type": schema.StringAttribute{
								Computed:      true,
								Description:   "The type of error returned. One of `api_error`, `card_error`, `idempotency_error`, or `invalid_request_error`",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("api_error", "card_error", "idempotency_error", "invalid_request_error")},
							},
						},
					},
					"collect_inputs": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Represents a reader action to collect customer inputs",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"inputs": schema.ListNestedAttribute{
								Computed:      true,
								Description:   "List of inputs to be collected.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"custom_text": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "Default text of input being collected.",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"description": schema.StringAttribute{
													Computed:      true,
													Description:   "Customize the default description for this input",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"skip_button": schema.StringAttribute{
													Computed:      true,
													Description:   "Customize the default label for this input's skip button",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"submit_button": schema.StringAttribute{
													Computed:      true,
													Description:   "Customize the default label for this input's submit button",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"title": schema.StringAttribute{
													Computed:      true,
													Description:   "Customize the default title for this input",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"email": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "Information about a email being collected using a reader",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Computed:      true,
													Description:   "The collected email address",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"numeric": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "Information about a number being collected using a reader",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Computed:      true,
													Description:   "The collected number",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"phone": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "Information about a phone number being collected using a reader",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Computed:      true,
													Description:   "The collected phone number",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"required": schema.BoolAttribute{
											Computed:      true,
											Description:   "Indicate that this input is required, disabling the skip button.",
											PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
										},
										"selection": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "Information about a selection being collected using a reader",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"choices": schema.ListNestedAttribute{
													Computed:      true,
													Description:   "List of possible choices to be selected",
													PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"id": schema.StringAttribute{
																Computed:      true,
																Description:   "The identifier for the selected choice. Maximum 50 characters.",
																PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
															},
															"style": schema.StringAttribute{
																Computed:      true,
																Description:   "The button style for the choice. Can be `primary` or `secondary`.",
																PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
																Validators:    []validator.String{stringvalidator.OneOf("primary", "secondary")},
															},
															"text": schema.StringAttribute{
																Computed:      true,
																Description:   "The text to be selected. Maximum 30 characters.",
																PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
															},
														},
													},
												},
												"id": schema.StringAttribute{
													Computed:      true,
													Description:   "The id of the selected choice",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"text": schema.StringAttribute{
													Computed:      true,
													Description:   "The text of the selected choice",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"signature": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "Information about a signature being collected using a reader",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Computed:      true,
													Description:   "The File ID of a collected signature image",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"skipped": schema.BoolAttribute{
											Computed:      true,
											Description:   "Indicate that this input was skipped by the user.",
											PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
										},
										"text": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "Information about text being collected using a reader",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Computed:      true,
													Description:   "The collected text value",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"toggles": schema.ListNestedAttribute{
											Computed:      true,
											Description:   "List of toggles being collected. Values are present if collection is complete.",
											PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"default_value": schema.StringAttribute{
														Computed:      true,
														Description:   "The toggle's default value. Can be `enabled` or `disabled`.",
														PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														Validators:    []validator.String{stringvalidator.OneOf("disabled", "enabled")},
													},
													"description": schema.StringAttribute{
														Computed:      true,
														Description:   "The toggle's description text. Maximum 50 characters.",
														PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
													},
													"title": schema.StringAttribute{
														Computed:      true,
														Description:   "The toggle's title text. Maximum 50 characters.",
														PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
													},
													"value": schema.StringAttribute{
														Computed:      true,
														Description:   "The toggle's collected value. Can be `enabled` or `disabled`.",
														PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														Validators:    []validator.String{stringvalidator.OneOf("disabled", "enabled")},
													},
												},
											},
										},
										"type": schema.StringAttribute{
											Computed:      true,
											Description:   "Type of input being collected.",
											PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											Validators:    []validator.String{stringvalidator.OneOf("email", "numeric", "phone", "selection", "signature", "text")},
										},
									},
								},
							},
							"metadata": schema.MapAttribute{
								Computed:      true,
								Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
								PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
						},
					},
					"collect_payment_method": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Represents a reader action to collect a payment method",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"collect_config": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Represents a per-transaction override of a reader configuration",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"enable_customer_cancellation": schema.BoolAttribute{
										Computed:      true,
										Description:   "Enable customer-initiated cancellation when processing this payment.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
									"skip_tipping": schema.BoolAttribute{
										Computed:      true,
										Description:   "Override showing a tipping selection screen on this transaction.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
									"tipping": schema.SingleNestedAttribute{
										Computed:      true,
										Description:   "Represents a per-transaction tipping configuration",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"amount_eligible": schema.Int64Attribute{
												Computed:      true,
												Description:   "Amount used to calculate tip suggestions on tipping selection screen for this transaction. Must be a positive integer in the smallest currency unit (e.g., 100 cents to represent $1.00 or 100 to represent ¥100, a zero-decimal currency).",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
											},
										},
									},
								},
							},
							"payment_intent": schema.StringAttribute{
								Computed:      true,
								Description:   "Most recent PaymentIntent processed by the reader.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payment_method": schema.StringAttribute{
								Computed:      true,
								Description:   "PaymentMethod objects represent your customer's payment instruments.\nYou can use them with [PaymentIntents](https://docs.stripe.com/payments/payment-intents) to collect payments or save them to\nCustomer objects to store instrument details for future payments.\n\nRelated guides: [Payment Methods](https://docs.stripe.com/payments/payment-methods) and [More Payment Scenarios](https://docs.stripe.com/payments/more-payment-scenarios).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"confirm_payment_intent": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Represents a reader action to confirm a payment",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"confirm_config": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Represents a per-transaction override of a reader configuration",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"return_url": schema.StringAttribute{
										Computed:      true,
										Description:   "If the customer doesn't abandon authenticating the payment, they're redirected to this URL after completion.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"payment_intent": schema.StringAttribute{
								Computed:      true,
								Description:   "Most recent PaymentIntent processed by the reader.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"failure_code": schema.StringAttribute{
						Computed:      true,
						Description:   "Failure code, only set if status is `failed`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"failure_message": schema.StringAttribute{
						Computed:      true,
						Description:   "Detailed failure message, only set if status is `failed`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"print_content": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Represents a reader action to print content",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"image": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Metadata of an uploaded file",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"created_at": schema.Int64Attribute{
										Computed:      true,
										Description:   "Creation time of the object (in seconds since the Unix epoch).",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"filename": schema.StringAttribute{
										Computed:      true,
										Description:   "The original name of the uploaded file (e.g. `receipt.png`).",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"size": schema.Int64Attribute{
										Computed:      true,
										Description:   "The size (in bytes) of the uploaded file.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"type": schema.StringAttribute{
										Computed:      true,
										Description:   "The format of the uploaded file.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"type": schema.StringAttribute{
								Computed:      true,
								Description:   "The type of content to print. Currently supports `image`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("image")},
							},
						},
					},
					"process_payment_intent": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Represents a reader action to process a payment intent",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"payment_intent": schema.StringAttribute{
								Computed:      true,
								Description:   "Most recent PaymentIntent processed by the reader.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"process_config": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Represents a per-transaction override of a reader configuration",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"enable_customer_cancellation": schema.BoolAttribute{
										Computed:      true,
										Description:   "Enable customer-initiated cancellation when processing this payment.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
									"return_url": schema.StringAttribute{
										Computed:      true,
										Description:   "If the customer doesn't abandon authenticating the payment, they're redirected to this URL after completion.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"skip_tipping": schema.BoolAttribute{
										Computed:      true,
										Description:   "Override showing a tipping selection screen on this transaction.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
									"tipping": schema.SingleNestedAttribute{
										Computed:      true,
										Description:   "Represents a per-transaction tipping configuration",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"amount_eligible": schema.Int64Attribute{
												Computed:      true,
												Description:   "Amount used to calculate tip suggestions on tipping selection screen for this transaction. Must be a positive integer in the smallest currency unit (e.g., 100 cents to represent $1.00 or 100 to represent ¥100, a zero-decimal currency).",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
											},
										},
									},
								},
							},
						},
					},
					"process_setup_intent": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Represents a reader action to process a setup intent",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"generated_card": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"process_config": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Represents a per-setup override of a reader configuration",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"enable_customer_cancellation": schema.BoolAttribute{
										Computed:      true,
										Description:   "Enable customer-initiated cancellation when processing this SetupIntent.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"setup_intent": schema.StringAttribute{
								Computed:      true,
								Description:   "Most recent SetupIntent processed by the reader.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"refund_payment": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Represents a reader action to refund a payment",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"amount": schema.Int64Attribute{
								Computed:      true,
								Description:   "The amount being refunded.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"charge": schema.StringAttribute{
								Computed:      true,
								Description:   "Charge that is being refunded.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"metadata": schema.MapAttribute{
								Computed:      true,
								Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
								PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"payment_intent": schema.StringAttribute{
								Computed:      true,
								Description:   "Payment intent that is being refunded.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reason": schema.StringAttribute{
								Computed:      true,
								Description:   "The reason for the refund.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("duplicate", "fraudulent", "requested_by_customer")},
							},
							"refund": schema.StringAttribute{
								Computed:      true,
								Description:   "Unique identifier for the refund object.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"refund_application_fee": schema.BoolAttribute{
								Computed:      true,
								Description:   "Boolean indicating whether the application fee should be refunded when refunding this charge. If a full charge refund is given, the full application fee will be refunded. Otherwise, the application fee will be refunded in an amount proportional to the amount of the charge refunded. An application fee can be refunded only by the application that created the charge.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"refund_payment_config": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Represents a per-transaction override of a reader configuration",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"enable_customer_cancellation": schema.BoolAttribute{
										Computed:      true,
										Description:   "Enable customer-initiated cancellation when refunding this payment.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"reverse_transfer": schema.BoolAttribute{
								Computed:      true,
								Description:   "Boolean indicating whether the transfer should be reversed when refunding this charge. The transfer will be reversed proportionally to the amount being refunded (either the entire or partial amount). A transfer can be reversed only by the application that created the charge.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"set_reader_display": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Represents a reader action to set the reader display",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"cart": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Cart object to be displayed by the reader, including line items, amounts, and currency.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"currency": schema.StringAttribute{
										Computed:      true,
										Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"line_items": schema.ListNestedAttribute{
										Computed:      true,
										Description:   "List of line items in the cart.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"amount": schema.Int64Attribute{
													Computed:      true,
													Description:   "The amount of the line item. A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
												},
												"description": schema.StringAttribute{
													Computed:      true,
													Description:   "Description of the line item.",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"quantity": schema.Int64Attribute{
													Computed:      true,
													Description:   "The quantity of the line item.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
												},
											},
										},
									},
									"tax": schema.Int64Attribute{
										Computed:      true,
										Description:   "Tax amount for the entire cart. A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"total": schema.Int64Attribute{
										Computed:      true,
										Description:   "Total amount for the entire cart, including tax. A positive integer in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
								},
							},
							"type": schema.StringAttribute{
								Computed:      true,
								Description:   "Type of information to be displayed by the reader. Only `cart` is currently supported.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("cart")},
							},
						},
					},
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "Status of the action performed by the reader.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("failed", "in_progress", "succeeded")},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "Type of action performed by the reader.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("collect_inputs", "collect_payment_method", "confirm_payment_intent", "print_content", "process_payment_intent", "process_setup_intent", "refund_payment", "set_reader_display")},
					},
				},
			},
			"device_sw_version": schema.StringAttribute{
				Computed:      true,
				Description:   "The current software version of the reader.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"device_type": schema.StringAttribute{
				Computed:      true,
				Description:   "Device type of the reader.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("bbpos_chipper2x", "bbpos_wisepad3", "bbpos_wisepos_e", "mobile_phone_reader", "simulated_stripe_s700", "simulated_stripe_s710", "simulated_verifone_m425", "simulated_verifone_p630", "simulated_verifone_ux700", "simulated_verifone_v660p", "simulated_wisepos_e", "stripe_m2", "stripe_s700", "stripe_s710", "verifone_P400", "verifone_m425", "verifone_p630", "verifone_ux700", "verifone_v660p")},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ip_address": schema.StringAttribute{
				Computed:      true,
				Description:   "The local IP address of the reader.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"label": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Custom label given to the reader for easier identification.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"location": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The location identifier of the reader.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"serial_number": schema.StringAttribute{
				Computed:      true,
				Description:   "Serial number of the reader.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The networking status of the reader. We do not recommend using this field in flows that may block taking payments.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("offline", "online")},
			},
			"registration_code": schema.StringAttribute{
				Required:      true,
				Description:   "A code generated by the reader used for registering to an account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
		},
	}
}

func (r *TerminalReaderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TerminalReaderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config TerminalReaderResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"RegistrationCode"}})

	params, err := expandTerminalReaderCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building TerminalReader create params", err.Error())
		return
	}

	obj, err := r.client.V1TerminalReaders.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating TerminalReader", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TerminalReaders.B, r.client.V1TerminalReaders.Key, stripe.FormatURLPath("/v1/terminal/readers/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TerminalReader create raw response", err.Error())
		return
	}

	if err := flattenTerminalReader(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TerminalReader create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"RegistrationCode"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TerminalReaderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState TerminalReaderResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TerminalReaderResourceModel
	state = priorState

	obj, err := r.client.V1TerminalReaders.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading TerminalReader", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TerminalReaders.B, r.client.V1TerminalReaders.Key, stripe.FormatURLPath("/v1/terminal/readers/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TerminalReader raw response", err.Error())
		return
	}

	if err := flattenTerminalReader(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening TerminalReader read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"RegistrationCode"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *TerminalReaderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TerminalReaderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config TerminalReaderResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"RegistrationCode"}})

	var state TerminalReaderResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"RegistrationCode"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"RegistrationCode"}})

	params, err := expandTerminalReaderUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building TerminalReader update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building TerminalReader update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1TerminalReaders.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating TerminalReader", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TerminalReaders.B, r.client.V1TerminalReaders.Key, stripe.FormatURLPath("/v1/terminal/readers/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TerminalReader update raw response", err.Error())
		return
	}

	if err := flattenTerminalReader(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TerminalReader update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"RegistrationCode"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TerminalReaderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TerminalReaderResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1TerminalReaders.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting TerminalReader", err.Error())
		return
	}
}

func (r *TerminalReaderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandTerminalReaderCreate(plan TerminalReaderResourceModel) (*stripe.TerminalReaderCreateParams, error) {
	params := &stripe.TerminalReaderCreateParams{}

	if !plan.Label.IsNull() && !plan.Label.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Label", "Label", plan.Label.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "label", params)
		}
	}
	if !plan.Location.IsNull() && !plan.Location.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LocationID", "Location", plan.Location.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "location", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.RegistrationCode.IsNull() && !plan.RegistrationCode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "RegistrationCode", "RegistrationCode", plan.RegistrationCode.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "registration_code", params)
		}
	}

	return params, nil
}

func expandTerminalReaderUpdate(plan TerminalReaderResourceModel, state TerminalReaderResourceModel) (*stripe.TerminalReaderUpdateParams, error) {
	params := &stripe.TerminalReaderUpdateParams{}

	if !plan.Label.Equal(state.Label) && !plan.Label.IsNull() && !plan.Label.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Label", "Label", plan.Label.ValueString()) {
			if !plan.Label.Equal(state.Label) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "label", params)
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

	return params, nil
}

func flattenTerminalReader(obj *stripe.TerminalReader, state *TerminalReaderResourceModel) error {
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
		assignedAction := false
		hadRawAction := false
		if rawValueAction, rawOk := plainValueAtPath(raw, "action"); rawOk {
			hadRawAction = true
			if rawValueAction != nil {
				sourceAction := applyConfiguredKeyedListShapes(rawValueAction, attrValueToPlain(state.Action))
				if valueAction, err := flattenPlainValue(sourceAction, types.ObjectType{AttrTypes: map[string]attr.Type{"api_error": types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}}, "collect_inputs": types.ObjectType{AttrTypes: map[string]attr.Type{"inputs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"custom_text": types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "skip_button": types.StringType, "submit_button": types.StringType, "title": types.StringType}}, "email": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "numeric": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "phone": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "required": types.BoolType, "selection": types.ObjectType{AttrTypes: map[string]attr.Type{"choices": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "style": types.StringType, "text": types.StringType}}}, "id": types.StringType, "text": types.StringType}}, "signature": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "skipped": types.BoolType, "text": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "toggles": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "description": types.StringType, "title": types.StringType, "value": types.StringType}}}, "type": types.StringType}}}, "metadata": types.MapType{ElemType: types.StringType}}}, "collect_payment_method": types.ObjectType{AttrTypes: map[string]attr.Type{"collect_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType, "skip_tipping": types.BoolType, "tipping": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_eligible": types.Int64Type}}}}, "payment_intent": types.StringType, "payment_method": types.StringType}}, "confirm_payment_intent": types.ObjectType{AttrTypes: map[string]attr.Type{"confirm_config": types.ObjectType{AttrTypes: map[string]attr.Type{"return_url": types.StringType}}, "payment_intent": types.StringType}}, "failure_code": types.StringType, "failure_message": types.StringType, "print_content": types.ObjectType{AttrTypes: map[string]attr.Type{"image": types.ObjectType{AttrTypes: map[string]attr.Type{"created_at": types.Int64Type, "filename": types.StringType, "size": types.Int64Type, "type": types.StringType}}, "type": types.StringType}}, "process_payment_intent": types.ObjectType{AttrTypes: map[string]attr.Type{"payment_intent": types.StringType, "process_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType, "return_url": types.StringType, "skip_tipping": types.BoolType, "tipping": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_eligible": types.Int64Type}}}}}}, "process_setup_intent": types.ObjectType{AttrTypes: map[string]attr.Type{"generated_card": types.StringType, "process_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType}}, "setup_intent": types.StringType}}, "refund_payment": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "charge": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "payment_intent": types.StringType, "reason": types.StringType, "refund": types.StringType, "refund_application_fee": types.BoolType, "refund_payment_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType}}, "reverse_transfer": types.BoolType}}, "set_reader_display": types.ObjectType{AttrTypes: map[string]attr.Type{"cart": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "line_items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "description": types.StringType, "quantity": types.Int64Type}}}, "tax": types.Int64Type, "total": types.Int64Type}}, "type": types.StringType}}, "status": types.StringType, "type": types.StringType}}, "action", "raw response"); err != nil {
					return err
				} else {
					if typedAction, ok := valueAction.(types.Object); ok {
						state.Action = typedAction
						assignedAction = true
					}
				}
			}
		}
		if !assignedAction {
			if !hasRaw {
				if responseValueAction, ok := plainFromResponseField(obj, "Action"); ok {
					sourceAction := applyConfiguredKeyedListShapes(responseValueAction, attrValueToPlain(state.Action))
					if valueAction, err := flattenPlainValue(
						sourceAction,
						types.ObjectType{AttrTypes: map[string]attr.Type{"api_error": types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}}, "collect_inputs": types.ObjectType{AttrTypes: map[string]attr.Type{"inputs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"custom_text": types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "skip_button": types.StringType, "submit_button": types.StringType, "title": types.StringType}}, "email": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "numeric": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "phone": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "required": types.BoolType, "selection": types.ObjectType{AttrTypes: map[string]attr.Type{"choices": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "style": types.StringType, "text": types.StringType}}}, "id": types.StringType, "text": types.StringType}}, "signature": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "skipped": types.BoolType, "text": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "toggles": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "description": types.StringType, "title": types.StringType, "value": types.StringType}}}, "type": types.StringType}}}, "metadata": types.MapType{ElemType: types.StringType}}}, "collect_payment_method": types.ObjectType{AttrTypes: map[string]attr.Type{"collect_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType, "skip_tipping": types.BoolType, "tipping": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_eligible": types.Int64Type}}}}, "payment_intent": types.StringType, "payment_method": types.StringType}}, "confirm_payment_intent": types.ObjectType{AttrTypes: map[string]attr.Type{"confirm_config": types.ObjectType{AttrTypes: map[string]attr.Type{"return_url": types.StringType}}, "payment_intent": types.StringType}}, "failure_code": types.StringType, "failure_message": types.StringType, "print_content": types.ObjectType{AttrTypes: map[string]attr.Type{"image": types.ObjectType{AttrTypes: map[string]attr.Type{"created_at": types.Int64Type, "filename": types.StringType, "size": types.Int64Type, "type": types.StringType}}, "type": types.StringType}}, "process_payment_intent": types.ObjectType{AttrTypes: map[string]attr.Type{"payment_intent": types.StringType, "process_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType, "return_url": types.StringType, "skip_tipping": types.BoolType, "tipping": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_eligible": types.Int64Type}}}}}}, "process_setup_intent": types.ObjectType{AttrTypes: map[string]attr.Type{"generated_card": types.StringType, "process_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType}}, "setup_intent": types.StringType}}, "refund_payment": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "charge": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "payment_intent": types.StringType, "reason": types.StringType, "refund": types.StringType, "refund_application_fee": types.BoolType, "refund_payment_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType}}, "reverse_transfer": types.BoolType}}, "set_reader_display": types.ObjectType{AttrTypes: map[string]attr.Type{"cart": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "line_items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "description": types.StringType, "quantity": types.Int64Type}}}, "tax": types.Int64Type, "total": types.Int64Type}}, "type": types.StringType}}, "status": types.StringType, "type": types.StringType}},
						"action",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAction, ok := valueAction.(types.Object); ok {
							state.Action = typedAction
							assignedAction = true
						}
					}
				}
			}
		}
		if !assignedAction && hadRawAction {
			if nullAction, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"api_error": types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}}, "collect_inputs": types.ObjectType{AttrTypes: map[string]attr.Type{"inputs": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"custom_text": types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "skip_button": types.StringType, "submit_button": types.StringType, "title": types.StringType}}, "email": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "numeric": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "phone": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "required": types.BoolType, "selection": types.ObjectType{AttrTypes: map[string]attr.Type{"choices": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "style": types.StringType, "text": types.StringType}}}, "id": types.StringType, "text": types.StringType}}, "signature": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "skipped": types.BoolType, "text": types.ObjectType{AttrTypes: map[string]attr.Type{"value": types.StringType}}, "toggles": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "description": types.StringType, "title": types.StringType, "value": types.StringType}}}, "type": types.StringType}}}, "metadata": types.MapType{ElemType: types.StringType}}}, "collect_payment_method": types.ObjectType{AttrTypes: map[string]attr.Type{"collect_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType, "skip_tipping": types.BoolType, "tipping": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_eligible": types.Int64Type}}}}, "payment_intent": types.StringType, "payment_method": types.StringType}}, "confirm_payment_intent": types.ObjectType{AttrTypes: map[string]attr.Type{"confirm_config": types.ObjectType{AttrTypes: map[string]attr.Type{"return_url": types.StringType}}, "payment_intent": types.StringType}}, "failure_code": types.StringType, "failure_message": types.StringType, "print_content": types.ObjectType{AttrTypes: map[string]attr.Type{"image": types.ObjectType{AttrTypes: map[string]attr.Type{"created_at": types.Int64Type, "filename": types.StringType, "size": types.Int64Type, "type": types.StringType}}, "type": types.StringType}}, "process_payment_intent": types.ObjectType{AttrTypes: map[string]attr.Type{"payment_intent": types.StringType, "process_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType, "return_url": types.StringType, "skip_tipping": types.BoolType, "tipping": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_eligible": types.Int64Type}}}}}}, "process_setup_intent": types.ObjectType{AttrTypes: map[string]attr.Type{"generated_card": types.StringType, "process_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType}}, "setup_intent": types.StringType}}, "refund_payment": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "charge": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "payment_intent": types.StringType, "reason": types.StringType, "refund": types.StringType, "refund_application_fee": types.BoolType, "refund_payment_config": types.ObjectType{AttrTypes: map[string]attr.Type{"enable_customer_cancellation": types.BoolType}}, "reverse_transfer": types.BoolType}}, "set_reader_display": types.ObjectType{AttrTypes: map[string]attr.Type{"cart": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "line_items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "description": types.StringType, "quantity": types.Int64Type}}}, "tax": types.Int64Type, "total": types.Int64Type}}, "type": types.StringType}}, "status": types.StringType, "type": types.StringType}}); ok {
				if typedAction, ok := nullAction.(types.Object); ok {
					state.Action = typedAction
				}
			}
		}
	}
	{
		if rawValueDeviceSwVersion, rawOk := plainValueAtPath(raw, "device_sw_version"); rawOk {
			if valueDeviceSwVersion, err := flattenPlainValue(rawValueDeviceSwVersion, types.StringType, "device_sw_version", "raw response"); err != nil {
				return err
			} else {
				if typedDeviceSwVersion, ok := valueDeviceSwVersion.(types.String); ok {
					state.DeviceSwVersion = typedDeviceSwVersion
				}
			}
		} else if !hasRaw {
			if responseValueDeviceSwVersion, ok := plainFromResponseField(obj, "DeviceSwVersion"); ok {
				if valueDeviceSwVersion, err := flattenPlainValue(responseValueDeviceSwVersion, types.StringType, "device_sw_version", "response struct"); err != nil {
					return err
				} else {
					if typedDeviceSwVersion, ok := valueDeviceSwVersion.(types.String); ok {
						state.DeviceSwVersion = typedDeviceSwVersion
					}
				}
			}
		}
	}
	{
		if rawValueDeviceType, rawOk := plainValueAtPath(raw, "device_type"); rawOk {
			if valueDeviceType, err := flattenPlainValue(rawValueDeviceType, types.StringType, "device_type", "raw response"); err != nil {
				return err
			} else {
				if typedDeviceType, ok := valueDeviceType.(types.String); ok {
					state.DeviceType = typedDeviceType
				}
			}
		} else if !hasRaw {
			if responseValueDeviceType, ok := plainFromResponseField(obj, "DeviceType"); ok {
				if valueDeviceType, err := flattenPlainValue(responseValueDeviceType, types.StringType, "device_type", "response struct"); err != nil {
					return err
				} else {
					if typedDeviceType, ok := valueDeviceType.(types.String); ok {
						state.DeviceType = typedDeviceType
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
		if rawValueIPAddress, rawOk := plainValueAtPath(raw, "ip_address"); rawOk {
			if valueIPAddress, err := flattenPlainValue(rawValueIPAddress, types.StringType, "ip_address", "raw response"); err != nil {
				return err
			} else {
				if typedIPAddress, ok := valueIPAddress.(types.String); ok {
					state.IPAddress = typedIPAddress
				}
			}
		} else if !hasRaw {
			if responseValueIPAddress, ok := plainFromResponseField(obj, "IPAddress"); ok {
				if valueIPAddress, err := flattenPlainValue(responseValueIPAddress, types.StringType, "ip_address", "response struct"); err != nil {
					return err
				} else {
					if typedIPAddress, ok := valueIPAddress.(types.String); ok {
						state.IPAddress = typedIPAddress
					}
				}
			}
		}
	}
	{
		if rawValueLabel, rawOk := plainValueAtPath(raw, "label"); rawOk {
			if valueLabel, err := flattenPlainValue(rawValueLabel, types.StringType, "label", "raw response"); err != nil {
				return err
			} else {
				if typedLabel, ok := valueLabel.(types.String); ok {
					state.Label = typedLabel
				}
			}
		} else if !hasRaw {
			if responseValueLabel, ok := plainFromResponseField(obj, "Label"); ok {
				if valueLabel, err := flattenPlainValue(responseValueLabel, types.StringType, "label", "response struct"); err != nil {
					return err
				} else {
					if typedLabel, ok := valueLabel.(types.String); ok {
						state.Label = typedLabel
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
		if state.Location.IsNull() || state.Location.IsUnknown() {
			if rawValueLocation, rawOk := plainValueAtPath(raw, "location"); rawOk {
				if typedLocation, ok := plainToStringIDValue(rawValueLocation); ok {
					state.Location = typedLocation
				}
			} else if !hasRaw {
				if responseValueLocation, ok := plainFromResponseField(obj, "Location"); ok {
					if typedLocation, ok := plainToStringIDValue(responseValueLocation); ok {
						state.Location = typedLocation
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
		if rawValueSerialNumber, rawOk := plainValueAtPath(raw, "serial_number"); rawOk {
			if valueSerialNumber, err := flattenPlainValue(rawValueSerialNumber, types.StringType, "serial_number", "raw response"); err != nil {
				return err
			} else {
				if typedSerialNumber, ok := valueSerialNumber.(types.String); ok {
					state.SerialNumber = typedSerialNumber
				}
			}
		} else if !hasRaw {
			if responseValueSerialNumber, ok := plainFromResponseField(obj, "SerialNumber"); ok {
				if valueSerialNumber, err := flattenPlainValue(responseValueSerialNumber, types.StringType, "serial_number", "response struct"); err != nil {
					return err
				} else {
					if typedSerialNumber, ok := valueSerialNumber.(types.String); ok {
						state.SerialNumber = typedSerialNumber
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
	return nil
}
