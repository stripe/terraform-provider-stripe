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

var _ resource.Resource = &PaymentIntentResource{}

var _ resource.ResourceWithConfigure = &PaymentIntentResource{}

var _ resource.ResourceWithImportState = &PaymentIntentResource{}

func NewPaymentIntentResource() resource.Resource {
	return &PaymentIntentResource{}
}

type PaymentIntentResource struct {
	client *stripe.Client
}

type PaymentIntentResourceModel struct {
	Object                            types.String `tfsdk:"object"`
	Amount                            types.Int64  `tfsdk:"amount"`
	AmountCapturable                  types.Int64  `tfsdk:"amount_capturable"`
	AmountReceived                    types.Int64  `tfsdk:"amount_received"`
	Application                       types.String `tfsdk:"application"`
	ApplicationFeeAmount              types.Int64  `tfsdk:"application_fee_amount"`
	AutomaticPaymentMethods           types.Object `tfsdk:"automatic_payment_methods"`
	CanceledAt                        types.Int64  `tfsdk:"canceled_at"`
	CancellationReason                types.String `tfsdk:"cancellation_reason"`
	CaptureMethod                     types.String `tfsdk:"capture_method"`
	ClientSecret                      types.String `tfsdk:"client_secret"`
	ConfirmationMethod                types.String `tfsdk:"confirmation_method"`
	Created                           types.Int64  `tfsdk:"created"`
	Currency                          types.String `tfsdk:"currency"`
	Customer                          types.String `tfsdk:"customer"`
	Description                       types.String `tfsdk:"description"`
	ExcludedPaymentMethodTypes        types.List   `tfsdk:"excluded_payment_method_types"`
	Hooks                             types.Object `tfsdk:"hooks"`
	ID                                types.String `tfsdk:"id"`
	LastPaymentError                  types.Object `tfsdk:"last_payment_error"`
	LatestCharge                      types.String `tfsdk:"latest_charge"`
	Livemode                          types.Bool   `tfsdk:"livemode"`
	ManagedPayments                   types.Object `tfsdk:"managed_payments"`
	Metadata                          types.Map    `tfsdk:"metadata"`
	NextAction                        types.Object `tfsdk:"next_action"`
	OnBehalfOf                        types.String `tfsdk:"on_behalf_of"`
	PaymentDetails                    types.Object `tfsdk:"payment_details"`
	PaymentMethod                     types.String `tfsdk:"payment_method"`
	PaymentMethodConfigurationDetails types.Object `tfsdk:"payment_method_configuration_details"`
	PaymentMethodOptions              types.Object `tfsdk:"payment_method_options"`
	PresentmentDetails                types.Object `tfsdk:"presentment_details"`
	Processing                        types.Object `tfsdk:"processing"`
	ReceiptEmail                      types.String `tfsdk:"receipt_email"`
	Review                            types.String `tfsdk:"review"`
	SetupFutureUsage                  types.String `tfsdk:"setup_future_usage"`
	Shipping                          types.Object `tfsdk:"shipping"`
	Source                            types.String `tfsdk:"source"`
	StatementDescriptor               types.String `tfsdk:"statement_descriptor"`
	StatementDescriptorSuffix         types.String `tfsdk:"statement_descriptor_suffix"`
	Status                            types.String `tfsdk:"status"`
	TransferData                      types.Object `tfsdk:"transfer_data"`
	TransferGroup                     types.String `tfsdk:"transfer_group"`
	Confirm                           types.Bool   `tfsdk:"confirm"`
	ConfirmationToken                 types.String `tfsdk:"confirmation_token"`
	ErrorOnRequiresAction             types.Bool   `tfsdk:"error_on_requires_action"`
	Mandate                           types.String `tfsdk:"mandate"`
	MandateData                       types.Object `tfsdk:"mandate_data"`
	PaymentMethodConfiguration        types.String `tfsdk:"payment_method_configuration"`
	PaymentMethodData                 types.Object `tfsdk:"payment_method_data"`
	RadarOptions                      types.Object `tfsdk:"radar_options"`
	ReturnURL                         types.String `tfsdk:"return_url"`
	UseStripeSDK                      types.Bool   `tfsdk:"use_stripe_sdk"`
}

func (r *PaymentIntentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PaymentIntentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_payment_intent"
}

func (r *PaymentIntentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A PaymentIntent guides you through the process of collecting a payment from your customer.\nWe recommend that you create exactly one PaymentIntent for each order or\ncustomer session in your system. You can reference the PaymentIntent later to\nsee the history of payment attempts for a particular session.\n\nA PaymentIntent transitions through\n[multiple statuses](/payments/paymentintents/lifecycle)\nthroughout its lifetime as it interfaces with Stripe.js to perform\nauthentication flows and ultimately creates at most one successful charge.\n\nRelated guide: [Payment Intents API](https://docs.stripe.com/payments/payment-intents)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("payment_intent")},
			},
			"amount": schema.Int64Attribute{
				Required:    true,
				Description: "Amount intended to be collected by this PaymentIntent. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://docs.stripe.com/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).",
			},
			"amount_capturable": schema.Int64Attribute{
				Computed:      true,
				Description:   "Amount that can be captured from this PaymentIntent.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"amount_received": schema.Int64Attribute{
				Computed:      true,
				Description:   "Amount that this PaymentIntent collects.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the Connect application that created the PaymentIntent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"application_fee_amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. The amount of the application fee collected will be capped at the total amount captured. For more information, see the PaymentIntents [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts).",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"automatic_payment_methods": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Settings to configure compatible payment methods from the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods)",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"allow_redirects": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Controls whether this PaymentIntent will accept redirect-based payment methods.\n\nRedirect-based payment methods may require your customer to be redirected to a payment method's app or site for authentication or additional steps. To [confirm](https://docs.stripe.com/api/payment_intents/confirm) this PaymentIntent, you may be required to provide a `return_url` to redirect customers back to your site after they authenticate or complete the payment.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("always", "never")},
					},
					"enabled": schema.BoolAttribute{
						Required:      true,
						Description:   "Automatically calculates compatible payment methods",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
					},
				},
			},
			"canceled_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "Populated when `status` is `canceled`, this is the time at which the PaymentIntent was canceled. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"cancellation_reason": schema.StringAttribute{
				Computed:      true,
				Description:   "Reason for cancellation of this PaymentIntent, either user-provided (`duplicate`, `fraudulent`, `requested_by_customer`, or `abandoned`) or generated by Stripe internally (`failed_invoice`, `void_invoice`, `automatic`, or `expired`).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("abandoned", "automatic", "duplicate", "expired", "failed_invoice", "fraudulent", "requested_by_customer", "void_invoice")},
			},
			"capture_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Controls when the funds will be captured from the customer's account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("automatic", "automatic_async", "manual")},
			},
			"client_secret": schema.StringAttribute{
				Computed:      true,
				Description:   "The client secret of this PaymentIntent. Used for client-side retrieval using a publishable key. \n\nThe client secret can be used to complete a payment from your frontend. It should not be stored, logged, or exposed to anyone other than the customer. Make sure that you have TLS enabled on any page that includes the client secret.\n\nRefer to our docs to [accept a payment](https://docs.stripe.com/payments/accept-a-payment?ui=elements) and learn about how `client_secret` should be handled.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Sensitive:     true,
			},
			"confirmation_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Describes whether we can confirm this PaymentIntent automatically, or if it requires customer action to confirm the payment.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("automatic", "manual")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Required:    true,
				Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the Customer this PaymentIntent belongs to, if one exists.\n\nPayment methods attached to other Customers cannot be used with this PaymentIntent.\n\nIf [setup_future_usage](https://api.stripe.com#payment_intent_object-setup_future_usage) is set and this PaymentIntent's payment method is not `card_present`, then the payment method attaches to the Customer after the PaymentIntent has been confirmed and any required actions from the user are complete. If the payment method is `card_present` and isn't a digital wallet, then a [generated_card](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card is created and attached to the Customer instead.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the object. Often useful for displaying to users.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"excluded_payment_method_types": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The list of payment method types to exclude from use with this payment.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"hooks": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"inputs": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"tax": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"calculation": schema.StringAttribute{
										Required:    true,
										Description: "The [TaxCalculation](https://docs.stripe.com/api/tax/calculations) id",
									},
								},
							},
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"last_payment_error": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The payment error encountered in the previous PaymentIntent confirmation. It will be cleared if the PaymentIntent is later updated for any reason.",
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
			"latest_charge": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the latest [Charge object](https://docs.stripe.com/api/charges) created by this PaymentIntent. This property is `null` until PaymentIntent confirmation is attempted.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"managed_payments": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Settings for Managed Payments.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Computed:      true,
						Description:   "Set to `true` to enable [Managed Payments](https://docs.stripe.com/payments/managed-payments), Stripe's merchant of record solution, for this session.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Learn more about [storing information in metadata](https://docs.stripe.com/payments/payment-intents/creating-payment-intents#storing-information-in-metadata).",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"next_action": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "If present, this property tells you what actions you need to take in order for your customer to fulfill a payment using the provided source.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"alipay_handle_redirect": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"native_data": schema.StringAttribute{
								Computed:      true,
								Description:   "The native data to be used with Alipay SDK you must redirect your customer to in order to authenticate the payment in an Android App.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"native_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The native URL you must redirect your customer to in order to authenticate the payment in an iOS App.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"return_url": schema.StringAttribute{
								Computed:      true,
								Description:   "If the customer does not exit their browser while authenticating, they will be redirected to this specified URL after completion.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL you must redirect your customer to in order to authenticate the payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"boleto_display_details": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"expires_at": schema.Int64Attribute{
								Computed:      true,
								Description:   "The timestamp after which the boleto expires.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"hosted_voucher_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL to the hosted boleto voucher page, which allows customers to view the boleto voucher.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"number": schema.StringAttribute{
								Computed:      true,
								Description:   "The boleto number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"pdf": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL to the downloadable boleto voucher PDF.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"card_await_notification": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"charge_attempt_at": schema.Int64Attribute{
								Computed:      true,
								Description:   "The time that payment will be attempted. If customer approval is required, they need to provide approval before this time.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"customer_approval_required": schema.BoolAttribute{
								Computed:      true,
								Description:   "For payments greater than INR 15000, the customer must provide explicit approval of the payment with their bank. For payments of lower amount, no customer action is required.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"cashapp_handle_redirect_or_display_qr_code": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"hosted_instructions_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL to the hosted Cash App Pay instructions page, which allows customers to view the QR code, and supports QR code refreshing on expiration.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"mobile_auth_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The url for mobile redirect based auth",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"qr_code": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"expires_at": schema.Int64Attribute{
										Computed:      true,
										Description:   "The date (unix timestamp) when the QR code expires.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"image_url_png": schema.StringAttribute{
										Computed:      true,
										Description:   "The image_url_png string used to render QR code",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"image_url_svg": schema.StringAttribute{
										Computed:      true,
										Description:   "The image_url_svg string used to render QR code",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
					"display_bank_transfer_instructions": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"amount_remaining": schema.Int64Attribute{
								Computed:      true,
								Description:   "The remaining amount that needs to be transferred to complete the payment.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"currency": schema.StringAttribute{
								Computed:      true,
								Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"financial_addresses": schema.ListNestedAttribute{
								Computed:      true,
								Description:   "A list of financial addresses that can be used to fund the customer balance",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"aba": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "ABA Records contain U.S. bank account details per the ABA format.",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"account_holder_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"account_holder_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The account holder name",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"account_number": schema.StringAttribute{
													Computed:      true,
													Description:   "The ABA account number",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"account_type": schema.StringAttribute{
													Computed:      true,
													Description:   "The account type",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"bank_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"bank_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The bank name",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"routing_number": schema.StringAttribute{
													Computed:      true,
													Description:   "The ABA routing number",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"iban": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "Iban Records contain E.U. bank account details per the SEPA format.",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"account_holder_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"account_holder_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The name of the person or business that owns the bank account",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"bank_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"bic": schema.StringAttribute{
													Computed:      true,
													Description:   "The BIC/SWIFT code of the account.",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"country": schema.StringAttribute{
													Computed:      true,
													Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"iban": schema.StringAttribute{
													Computed:      true,
													Description:   "The IBAN of the account.",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"sort_code": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "Sort Code Records contain U.K. bank account details per the sort code format.",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"account_holder_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"account_holder_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The name of the person or business that owns the bank account",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"account_number": schema.StringAttribute{
													Computed:      true,
													Description:   "The account number",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"bank_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"sort_code": schema.StringAttribute{
													Computed:      true,
													Description:   "The six-digit sort code",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"spei": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "SPEI Records contain Mexico bank account details per the SPEI format.",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"account_holder_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"account_holder_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The account holder name",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"bank_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"bank_code": schema.StringAttribute{
													Computed:      true,
													Description:   "The three-digit bank code",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"bank_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The short banking institution name",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"clabe": schema.StringAttribute{
													Computed:      true,
													Description:   "The CLABE number",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"supported_networks": schema.ListAttribute{
											Computed:      true,
											Description:   "The payment networks supported by this FinancialAddress",
											PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
											ElementType:   types.StringType,
										},
										"swift": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "SWIFT Records contain U.S. bank account details per the SWIFT format.",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"account_holder_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"account_holder_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The account holder name",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"account_number": schema.StringAttribute{
													Computed:      true,
													Description:   "The account number",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"account_type": schema.StringAttribute{
													Computed:      true,
													Description:   "The account type",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"bank_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"bank_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The bank name",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"swift_code": schema.StringAttribute{
													Computed:      true,
													Description:   "The SWIFT code",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
										"type": schema.StringAttribute{
											Computed:      true,
											Description:   "The type of financial address",
											PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											Validators:    []validator.String{stringvalidator.OneOf("aba", "iban", "sort_code", "spei", "swift", "zengin")},
										},
										"zengin": schema.SingleNestedAttribute{
											Computed:      true,
											Description:   "Zengin Records contain Japan bank account details per the Zengin format.",
											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"account_holder_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"account_holder_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The account holder name",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"account_number": schema.StringAttribute{
													Computed:      true,
													Description:   "The account number",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"account_type": schema.StringAttribute{
													Computed:      true,
													Description:   "The bank account type. In Japan, this can only be `futsu` or `toza`.",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"bank_address": schema.SingleNestedAttribute{
													Computed: true,

													PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
													Attributes: map[string]schema.Attribute{
														"city": schema.StringAttribute{
															Computed:      true,
															Description:   "City, district, suburb, town, or village.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"country": schema.StringAttribute{
															Computed:      true,
															Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line1": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 1, such as the street, PO Box, or company name.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"line2": schema.StringAttribute{
															Computed:      true,
															Description:   "Address line 2, such as the apartment, suite, unit, or building.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"postal_code": schema.StringAttribute{
															Computed:      true,
															Description:   "ZIP or postal code.",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
														"state": schema.StringAttribute{
															Computed:      true,
															Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
															PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
														},
													},
												},
												"bank_code": schema.StringAttribute{
													Computed:      true,
													Description:   "The bank code of the account",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"bank_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The bank name of the account",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"branch_code": schema.StringAttribute{
													Computed:      true,
													Description:   "The branch code of the account",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
												"branch_name": schema.StringAttribute{
													Computed:      true,
													Description:   "The branch name of the account",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												},
											},
										},
									},
								},
							},
							"hosted_instructions_url": schema.StringAttribute{
								Computed:      true,
								Description:   "A link to a hosted page that guides your customer through completing the transfer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reference": schema.StringAttribute{
								Computed:      true,
								Description:   "A string identifying this payment. Instruct your customer to include this code in the reference or memo field of their bank transfer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"type": schema.StringAttribute{
								Computed:      true,
								Description:   "Type of bank transfer",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("eu_bank_transfer", "gb_bank_transfer", "jp_bank_transfer", "mx_bank_transfer", "us_bank_transfer")},
							},
						},
					},
					"klarna_display_qr_code": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Computed:      true,
								Description:   "The data being used to generate QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expires_at": schema.Int64Attribute{
								Computed:      true,
								Description:   "The timestamp at which the QR code expires.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"image_url_png": schema.StringAttribute{
								Computed:      true,
								Description:   "The image_url_png string used to render QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"image_url_svg": schema.StringAttribute{
								Computed:      true,
								Description:   "The image_url_svg string used to render QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"konbini_display_details": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"expires_at": schema.Int64Attribute{
								Computed:      true,
								Description:   "The timestamp at which the pending Konbini payment expires.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"hosted_voucher_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL for the Konbini payment instructions page, which allows customers to view and print a Konbini voucher.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"stores": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"familymart": schema.SingleNestedAttribute{
										Computed:      true,
										Description:   "FamilyMart instruction details.",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"confirmation_number": schema.StringAttribute{
												Computed:      true,
												Description:   "The confirmation number.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"payment_code": schema.StringAttribute{
												Computed:      true,
												Description:   "The payment code.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
										},
									},
									"lawson": schema.SingleNestedAttribute{
										Computed:      true,
										Description:   "Lawson instruction details.",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"confirmation_number": schema.StringAttribute{
												Computed:      true,
												Description:   "The confirmation number.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"payment_code": schema.StringAttribute{
												Computed:      true,
												Description:   "The payment code.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
										},
									},
									"ministop": schema.SingleNestedAttribute{
										Computed:      true,
										Description:   "Ministop instruction details.",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"confirmation_number": schema.StringAttribute{
												Computed:      true,
												Description:   "The confirmation number.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"payment_code": schema.StringAttribute{
												Computed:      true,
												Description:   "The payment code.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
										},
									},
									"seicomart": schema.SingleNestedAttribute{
										Computed:      true,
										Description:   "Seicomart instruction details.",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"confirmation_number": schema.StringAttribute{
												Computed:      true,
												Description:   "The confirmation number.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"payment_code": schema.StringAttribute{
												Computed:      true,
												Description:   "The payment code.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
										},
									},
								},
							},
						},
					},
					"multibanco_display_details": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"entity": schema.StringAttribute{
								Computed:      true,
								Description:   "Entity number associated with this Multibanco payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expires_at": schema.Int64Attribute{
								Computed:      true,
								Description:   "The timestamp at which the Multibanco voucher expires.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"hosted_voucher_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL for the hosted Multibanco voucher page, which allows customers to view a Multibanco voucher.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reference": schema.StringAttribute{
								Computed:      true,
								Description:   "Reference number associated with this Multibanco payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"oxxo_display_details": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"expires_after": schema.Int64Attribute{
								Computed:      true,
								Description:   "The timestamp after which the OXXO voucher expires.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"hosted_voucher_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL for the hosted OXXO voucher page, which allows customers to view and print an OXXO voucher.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"number": schema.StringAttribute{
								Computed:      true,
								Description:   "OXXO reference number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"paynow_display_qr_code": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Computed:      true,
								Description:   "The raw data string used to generate QR code, it should be used together with QR code library.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"hosted_instructions_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL to the hosted PayNow instructions page, which allows customers to view the PayNow QR code.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"image_url_png": schema.StringAttribute{
								Computed:      true,
								Description:   "The image_url_png string used to render QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"image_url_svg": schema.StringAttribute{
								Computed:      true,
								Description:   "The image_url_svg string used to render QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"pix_display_qr_code": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Computed:      true,
								Description:   "The raw data string used to generate QR code, it should be used together with QR code library.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expires_at": schema.Int64Attribute{
								Computed:      true,
								Description:   "The date (unix timestamp) when the PIX expires.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"hosted_instructions_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL to the hosted pix instructions page, which allows customers to view the pix QR code.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"image_url_png": schema.StringAttribute{
								Computed:      true,
								Description:   "The image_url_png string used to render png QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"image_url_svg": schema.StringAttribute{
								Computed:      true,
								Description:   "The image_url_svg string used to render svg QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"promptpay_display_qr_code": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Computed:      true,
								Description:   "The raw data string used to generate QR code, it should be used together with QR code library.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"hosted_instructions_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL to the hosted PromptPay instructions page, which allows customers to view the PromptPay QR code.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"image_url_png": schema.StringAttribute{
								Computed:      true,
								Description:   "The PNG path used to render the QR code, can be used as the source in an HTML img tag",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"image_url_svg": schema.StringAttribute{
								Computed:      true,
								Description:   "The SVG path used to render the QR code, can be used as the source in an HTML img tag",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"redirect_to_url": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"return_url": schema.StringAttribute{
								Computed:      true,
								Description:   "If the customer does not exit their browser while authenticating, they will be redirected to this specified URL after completion.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL you must redirect your customer to in order to authenticate the payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"swish_handle_redirect_or_display_qr_code": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"hosted_instructions_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL to the hosted Swish instructions page, which allows customers to view the QR code.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"mobile_auth_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The url for mobile redirect based auth (for internal use only and not typically available in standard API requests).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"qr_code": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"data": schema.StringAttribute{
										Computed:      true,
										Description:   "The raw data string used to generate QR code, it should be used together with QR code library.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"image_url_png": schema.StringAttribute{
										Computed:      true,
										Description:   "The image_url_png string used to render QR code",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"image_url_svg": schema.StringAttribute{
										Computed:      true,
										Description:   "The image_url_svg string used to render QR code",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "Type of the next action to perform. Refer to the other child attributes under `next_action` for available values. Examples include: `redirect_to_url`, `use_stripe_sdk`, `alipay_handle_redirect`, `oxxo_display_details`, or `verify_with_microdeposits`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"upi_handle_redirect_or_display_qr_code": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"hosted_instructions_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL to the hosted UPI instructions page, which allows customers to view the QR code.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"qr_code": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"expires_at": schema.Int64Attribute{
										Computed:      true,
										Description:   "The date (unix timestamp) when the QR code expires.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"image_url_png": schema.StringAttribute{
										Computed:      true,
										Description:   "The image_url_png string used to render QR code",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"image_url_svg": schema.StringAttribute{
										Computed:      true,
										Description:   "The image_url_svg string used to render QR code",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
					"use_stripe_sdk": schema.MapAttribute{
						Computed:      true,
						Description:   "When confirming a PaymentIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"verify_with_microdeposits": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"arrival_date": schema.Int64Attribute{
								Computed:      true,
								Description:   "The timestamp when the microdeposits are expected to land.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"hosted_verification_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL for the hosted verification page, which allows customers to verify their bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"microdeposit_type": schema.StringAttribute{
								Computed:      true,
								Description:   "The type of the microdeposit sent to the customer. Used to distinguish between different verification methods.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("amounts", "descriptor_code")},
							},
						},
					},
					"wechat_pay_display_qr_code": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Computed:      true,
								Description:   "The data being used to generate QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"hosted_instructions_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The URL to the hosted WeChat Pay instructions page, which allows customers to view the WeChat Pay QR code.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"image_data_url": schema.StringAttribute{
								Computed:      true,
								Description:   "The base64 image data for a pre-generated QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"image_url_png": schema.StringAttribute{
								Computed:      true,
								Description:   "The image_url_png string used to render QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"image_url_svg": schema.StringAttribute{
								Computed:      true,
								Description:   "The image_url_svg string used to render QR code",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"wechat_pay_redirect_to_android_app": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"app_id": schema.StringAttribute{
								Computed:      true,
								Description:   "app_id is the APP ID registered on WeChat open platform",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"nonce_str": schema.StringAttribute{
								Computed:      true,
								Description:   "nonce_str is a random string",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"package": schema.StringAttribute{
								Computed:      true,
								Description:   "package is static value",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"partner_id": schema.StringAttribute{
								Computed:      true,
								Description:   "an unique merchant ID assigned by WeChat Pay",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"prepay_id": schema.StringAttribute{
								Computed:      true,
								Description:   "an unique trading ID assigned by WeChat Pay",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"sign": schema.StringAttribute{
								Computed:      true,
								Description:   "A signature",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"timestamp": schema.StringAttribute{
								Computed:      true,
								Description:   "Specifies the current time in epoch format",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"wechat_pay_redirect_to_ios_app": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"native_url": schema.StringAttribute{
								Computed:      true,
								Description:   "An universal link that redirect to WeChat Pay app",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"on_behalf_of": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "You can specify the settlement merchant as the\nconnected account using the `on_behalf_of` attribute on the charge. See the PaymentIntents [use case for connected accounts](/payments/connected-accounts) for details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"payment_details": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"customer_reference": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A unique value to identify the customer. This field is available only for card payments.\n\nThis field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"order_reference": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A unique value assigned by the business to identify the transaction. Required for L2 and L3 rates.\n\nFor Cards, this field is truncated to 25 alphanumeric characters, excluding spaces, before being sent to card networks. For Klarna, this field is truncated to 255 characters and is visible to customers when they view the order in the Klarna app.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"payment_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the payment method used in this PaymentIntent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"payment_method_configuration_details": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Information about the [payment method configuration](https://docs.stripe.com/api/payment_method_configurations) used for this PaymentIntent.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed:      true,
						Description:   "ID of the payment method configuration used.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"parent": schema.StringAttribute{
						Computed:      true,
						Description:   "ID of the parent payment method configuration used.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"payment_method_options": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Payment-method-specific configuration for this PaymentIntent.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"acss_debit": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"mandate_options": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"custom_mandate_url": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "A URL for custom mandate text",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"interval_description": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"payment_schedule": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Payment schedule for the mandate.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("combined", "interval", "sporadic")},
									},
									"transaction_type": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Transaction type of the mandate.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("business", "personal")},
									},
								},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"target_date": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"verification_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Bank account verification method. The default value is `automatic`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("automatic", "instant", "microdeposits")},
							},
						},
					},
					"affirm": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"preferred_locale": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Preferred language of the Affirm authorization page that the customer is redirected to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"afterpay_clearpay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"reference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "An internal identifier or reference that this payment corresponds to. You must limit the identifier to 128 characters, and it can only contain letters, numbers, underscores, backslashes, and dashes.\nThis field differs from the statement descriptor and item name.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"alipay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"alma": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"amazon_pay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"au_becs_debit": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"target_date": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"bacs_debit": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"mandate_options": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"reference_prefix": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'DDIC' or 'STRIPE'.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"target_date": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"bancontact": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"preferred_language": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Preferred language of the Bancontact authorization page that the customer is redirected to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("de", "en", "fr", "nl")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"billie": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"blik": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
							"code": schema.StringAttribute{
								Optional:    true,
								Description: "The 6-digit BLIK code that a customer has generated using their banking application. Can only be set on confirmation.",
							},
						},
					},
					"boleto": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"expires_after_days": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The number of calendar days before a Boleto voucher expires. For example, if you create a Boleto voucher on Monday and you set expires_after_days to 2, the Boleto voucher will expire on Wednesday at 23:59 America/Sao_Paulo time.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
						},
					},
					"card": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"installments": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Installment details for this payment.\n\nFor more information, see the [installments integration guide](https://docs.stripe.com/payments/installments).",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"available_plans": schema.ListNestedAttribute{
										Computed:      true,
										Description:   "Installment plans that may be selected for this PaymentIntent.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"count": schema.Int64Attribute{
													Computed:      true,
													Description:   "For `fixed_count` installment plans, this is the number of installment payments your customer will make to their credit card.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
												},
												"interval": schema.StringAttribute{
													Computed:      true,
													Description:   "For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card.\nOne of `month`.",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
													Validators:    []validator.String{stringvalidator.OneOf("month")},
												},
												"type": schema.StringAttribute{
													Computed:      true,
													Description:   "Type of installment plan, one of `fixed_count`, `bonus`, or `revolving`.",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
													Validators:    []validator.String{stringvalidator.OneOf("bonus", "fixed_count", "revolving")},
												},
											},
										},
									},
									"enabled": schema.BoolAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Whether Installments are enabled for this PaymentIntent.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
									"plan": schema.SingleNestedAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Installment plan selected for this PaymentIntent.",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"count": schema.Int64Attribute{
												Optional:      true,
												Computed:      true,
												Description:   "For `fixed_count` installment plans, this is the number of installment payments your customer will make to their credit card.",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
											},
											"interval": schema.StringAttribute{
												Optional:      true,
												Computed:      true,
												Description:   "For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card.\nOne of `month`.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												Validators:    []validator.String{stringvalidator.OneOf("month")},
											},
											"type": schema.StringAttribute{
												Required:    true,
												Description: "Type of installment plan, one of `fixed_count`, `bonus`, or `revolving`.",
												Validators:  []validator.String{stringvalidator.OneOf("bonus", "fixed_count", "revolving")},
											},
										},
									},
								},
							},
							"mandate_options": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Configuration options for setting up an eMandate for cards issued in India.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"amount": schema.Int64Attribute{
										Required:    true,
										Description: "Amount to be charged for future payments, specified in the presentment currency.",
									},
									"amount_type": schema.StringAttribute{
										Required:    true,
										Description: "One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.",
										Validators:  []validator.String{stringvalidator.OneOf("fixed", "maximum")},
									},
									"description": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "A description of the mandate or subscription that is meant to be displayed to the customer.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"end_date": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "End date of the mandate or subscription. If not provided, the mandate will be active until canceled. If provided, end date should be after start date.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"interval": schema.StringAttribute{
										Required:    true,
										Description: "Specifies payment frequency. One of `day`, `week`, `month`, `year`, or `sporadic`.",
										Validators:  []validator.String{stringvalidator.OneOf("day", "month", "sporadic", "week", "year")},
									},
									"interval_count": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "The number of intervals between payments. For example, `interval=month` and `interval_count=3` indicates one payment every three months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks). This parameter is optional when `interval=sporadic`.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"reference": schema.StringAttribute{
										Required:    true,
										Description: "Unique identifier for the mandate or subscription.",
									},
									"start_date": schema.Int64Attribute{
										Required:    true,
										Description: "Start date of the mandate or subscription. Start date should not be lesser than yesterday.",
									},
									"supported_types": schema.ListAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Specifies the type of mandates supported. Possible values are `india`.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
								},
							},
							"network": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Selected network to process this payment intent on. Depends on the available networks of the card attached to the payment intent. Can be only set confirm-time.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("amex", "cartes_bancaires", "diners", "discover", "eftpos_au", "girocard", "interac", "jcb", "link", "mastercard", "unionpay", "unknown", "visa")},
							},
							"request_extended_authorization": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Request ability to [capture beyond the standard authorization validity window](https://docs.stripe.com/payments/extended-authorization) for this PaymentIntent.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("if_available", "never")},
							},
							"request_incremental_authorization": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Request ability to [increment the authorization](https://docs.stripe.com/payments/incremental-authorization) for this PaymentIntent.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("if_available", "never")},
							},
							"request_multicapture": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Request ability to make [multiple captures](https://docs.stripe.com/payments/multicapture) for this PaymentIntent.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("if_available", "never")},
							},
							"request_overcapture": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Request ability to [overcapture](https://docs.stripe.com/payments/overcapture) for this PaymentIntent.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("if_available", "never")},
							},
							"request_three_d_secure": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("any", "automatic", "challenge")},
							},
							"require_cvc_recollection": schema.BoolAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "When enabled, using a card that is attached to a customer will require the CVC to be provided again (i.e. using the cvc_token parameter).",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"statement_descriptor_suffix_kana": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Provides information about a card payment that customers see on their statements. Concatenated with the Kana prefix (shortened Kana descriptor) or Kana statement descriptor that’s set on the account to form the complete statement descriptor. Maximum 22 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 22 characters.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"statement_descriptor_suffix_kanji": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Provides information about a card payment that customers see on their statements. Concatenated with the Kanji prefix (shortened Kanji descriptor) or Kanji statement descriptor that’s set on the account to form the complete statement descriptor. Maximum 17 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 17 characters.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"cvc_token": schema.StringAttribute{
								Optional:    true,
								Description: "A single-use `cvc_update` Token that represents a card CVC value. When provided, the CVC value will be verified during the card payment attempt. This parameter can only be provided during confirmation.",
							},
							"moto": schema.BoolAttribute{
								Optional:    true,
								Description: "When specified, this parameter indicates that a transaction will be marked\nas MOTO (Mail Order Telephone Order) and thus out of scope for SCA. This\nparameter can only be provided during confirmation.",
							},
							"three_d_secure": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "If 3D Secure authentication was performed with a third-party provider,\nthe authentication details to use for this payment.",
								Attributes: map[string]schema.Attribute{
									"ares_trans_status": schema.StringAttribute{
										Optional:    true,
										Description: "The `transStatus` returned from the card Issuer’s ACS in the ARes.",
									},
									"cryptogram": schema.StringAttribute{
										Required:    true,
										Description: "The cryptogram, also known as the \"authentication value\" (AAV, CAVV or\nAEVV). This value is 20 bytes, base64-encoded into a 28-character string.\n(Most 3D Secure providers will return the base64-encoded version, which\nis what you should specify here.)",
									},
									"electronic_commerce_indicator": schema.StringAttribute{
										Optional:    true,
										Description: "The Electronic Commerce Indicator (ECI) is returned by your 3D Secure\nprovider and indicates what degree of authentication was performed.",
									},
									"exemption_indicator": schema.StringAttribute{
										Optional:    true,
										Description: "The exemption requested via 3DS and accepted by the issuer at authentication time.",
									},
									"network_options": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "Network specific 3DS fields. Network specific arguments require an\nexplicit card brand choice. The parameter `payment_method_options.card.network``\nmust be populated accordingly",
										Attributes: map[string]schema.Attribute{
											"cartes_bancaires": schema.SingleNestedAttribute{
												Optional:    true,
												Description: "Cartes Bancaires-specific 3DS fields.",
												Attributes: map[string]schema.Attribute{
													"cb_avalgo": schema.StringAttribute{
														Required:    true,
														Description: "The cryptogram calculation algorithm used by the card Issuer's ACS\nto calculate the Authentication cryptogram. Also known as `cavvAlgorithm`.\nmessageExtension: CB-AVALGO",
													},
													"cb_exemption": schema.StringAttribute{
														Optional:    true,
														Description: "The exemption indicator returned from Cartes Bancaires in the ARes.\nmessage extension: CB-EXEMPTION; string (4 characters)\nThis is a 3 byte bitmap (low significant byte first and most significant\nbit first) that has been Base64 encoded",
													},
													"cb_score": schema.Int64Attribute{
														Optional:    true,
														Description: "The risk score returned from Cartes Bancaires in the ARes.\nmessage extension: CB-SCORE; numeric value 0-99",
													},
												},
											},
										},
									},
									"requestor_challenge_indicator": schema.StringAttribute{
										Optional:    true,
										Description: "The challenge indicator (`threeDSRequestorChallengeInd`) which was requested in the\nAReq sent to the card Issuer's ACS. A string containing 2 digits from 01-99.",
									},
									"transaction_id": schema.StringAttribute{
										Required:    true,
										Description: "For 3D Secure 1, the XID. For 3D Secure 2, the Directory Server\nTransaction ID (dsTransID).",
									},
									"version": schema.StringAttribute{
										Required:    true,
										Description: "The version of 3D Secure that was performed.",
									},
								},
							},
						},
					},
					"card_present": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual", "manual_preferred")},
							},
							"request_extended_authorization": schema.BoolAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Request ability to capture this payment beyond the standard [authorization validity window](https://docs.stripe.com/terminal/features/extended-authorizations#authorization-validity)",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"request_incremental_authorization_support": schema.BoolAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Request ability to [increment](https://docs.stripe.com/terminal/features/incremental-authorizations) this PaymentIntent if the combination of MCC and card brand is eligible. Check [incremental_authorization_supported](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-card_present-incremental_authorization_supported) in the [Confirm](https://docs.stripe.com/api/payment_intents/confirm) response to verify support.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"routing": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"requested_priority": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Requested routing priority",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("domestic", "international")},
									},
								},
							},
						},
					},
					"cashapp": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
						},
					},
					"crypto": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"customer_balance": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank_transfer": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"eu_bank_transfer": schema.SingleNestedAttribute{
										Optional: true,
										Computed: true,

										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"country": schema.StringAttribute{
												Required:    true,
												Description: "The desired country code of the bank account information. Permitted values include: `DE`, `FR`, `IE`, or `NL`.",
												Validators:  []validator.String{stringvalidator.OneOf("BE", "DE", "ES", "FR", "IE", "NL")},
											},
										},
									},
									"requested_address_types": schema.ListAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.\n\nPermitted values include: `sort_code`, `zengin`, `iban`, or `spei`.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
									"type": schema.StringAttribute{
										Required:    true,
										Description: "The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.",
										Validators:  []validator.String{stringvalidator.OneOf("eu_bank_transfer", "gb_bank_transfer", "jp_bank_transfer", "mx_bank_transfer", "us_bank_transfer")},
									},
								},
							},
							"funding_type": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("bank_transfer")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"eps": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"fpx": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"giropay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"grabpay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"ideal": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"kakao_pay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"klarna": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"preferred_locale": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Preferred locale of the Klarna checkout page that the customer is redirected to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"on_demand": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "On-demand details if setting up or charging an on-demand payment.",
								Attributes: map[string]schema.Attribute{
									"average_amount": schema.Int64Attribute{
										Optional:    true,
										Description: "Your average amount value. You can use a value across your customer base, or segment based on customer type, country, etc.",
									},
									"maximum_amount": schema.Int64Attribute{
										Optional:    true,
										Description: "The maximum value you may charge a customer per purchase. You can use a value across your customer base, or segment based on customer type, country, etc.",
									},
									"minimum_amount": schema.Int64Attribute{
										Optional:    true,
										Description: "The lowest or minimum value you may charge a customer per purchase. You can use a value across your customer base, or segment based on customer type, country, etc.",
									},
									"purchase_interval": schema.StringAttribute{
										Optional:    true,
										Description: "Interval at which the customer is making purchases",
									},
									"purchase_interval_count": schema.Int64Attribute{
										Optional:    true,
										Description: "The number of `purchase_interval` between charges",
									},
								},
							},
							"subscriptions": schema.ListNestedAttribute{
								Optional:    true,
								Description: "Subscription details if setting up or charging a subscription.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"interval": schema.StringAttribute{
											Required:    true,
											Description: "Unit of time between subscription charges.",
										},
										"interval_count": schema.Int64Attribute{
											Optional:    true,
											Description: "The number of intervals (specified in the `interval` attribute) between subscription charges. For example, `interval=month` and `interval_count=3` charges every 3 months.",
										},
										"name": schema.StringAttribute{
											Optional:    true,
											Description: "Name for subscription.",
										},
										"next_billing": schema.SingleNestedAttribute{
											Optional:    true,
											Description: "Describes the upcoming charge for this subscription.",
											Attributes: map[string]schema.Attribute{
												"amount": schema.Int64Attribute{
													Required:    true,
													Description: "The amount of the next charge for the subscription.",
												},
												"date": schema.StringAttribute{
													Required:    true,
													Description: "The date of the next charge for the subscription in YYYY-MM-DD format.",
												},
											},
										},
										"reference": schema.StringAttribute{
											Required:    true,
											Description: "A non-customer-facing reference to correlate subscription charges in the Klarna app. Use a value that persists across subscription charges.",
										},
									},
								},
							},
						},
					},
					"konbini": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"confirmation_number": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "An optional 10 to 11 digit numeric-only string determining the confirmation code at applicable convenience stores.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expires_after_days": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The number of calendar days (between 1 and 60) after which Konbini payment instructions will expire. For example, if a PaymentIntent is confirmed with Konbini and `expires_after_days` set to 2 on Monday JST, the instructions will expire on Wednesday 23:59:59 JST.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"expires_at": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The timestamp at which the Konbini payment instructions will expire. Only one of `expires_after_days` or `expires_at` may be set.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"product_description": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "A product descriptor of up to 22 characters, which will appear to customers at the convenience store.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"kr_card": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"mb_way": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"mobilepay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"multibanco": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"naver_pay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"nz_bank_account": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"target_date": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"oxxo": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"expires_after_days": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The number of calendar days before an OXXO invoice expires. For example, if you create an OXXO invoice on Monday and you set expires_after_days to 2, the OXXO invoice will expire on Wednesday at 23:59 America/Mexico_City time.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"p24": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
							"tos_shown_and_accepted": schema.BoolAttribute{
								Optional:    true,
								Description: "Confirm that the payer has accepted the P24 terms and conditions.",
							},
						},
					},
					"payco": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"paynow": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"paypal": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"preferred_locale": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Preferred locale of the PayPal checkout page that the customer is redirected to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "A reference of the PayPal transaction visible to customer which is mapped to PayPal's invoice ID. This must be a globally unique ID if you have configured in your PayPal settings to block multiple payments per invoice ID.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
							"risk_correlation_id": schema.StringAttribute{
								Optional:    true,
								Description: "The risk correlation ID for an on-session payment using a saved PayPal payment method.",
							},
						},
					},
					"payto": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"mandate_options": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"amount": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "Amount that will be collected. It is required when `amount_type` is `fixed`.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"amount_type": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The type of amount that will be collected. The amount charged must be exact or up to the value of `amount` param for `fixed` or `maximum` type respectively. Defaults to `maximum`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("fixed", "maximum")},
									},
									"end_date": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Date, in YYYY-MM-DD format, after which payments will not be collected. Defaults to no end date.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"payment_schedule": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The periodicity at which payments will be collected. Defaults to `adhoc`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("adhoc", "annual", "daily", "fortnightly", "monthly", "quarterly", "semi_annual", "weekly")},
									},
									"payments_per_period": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "The number of payments that will be made during a payment period. Defaults to 1 except for when `payment_schedule` is `adhoc`. In that case, it defaults to no limit.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"purpose": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The purpose for which payments are made. Has a default value based on your merchant category code.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("dependant_support", "government", "loan", "mortgage", "other", "pension", "personal", "retail", "salary", "tax", "utility")},
									},
								},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"pix": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"amount_includes_iof": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Determines if the amount includes the IOF tax.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("always", "never")},
							},
							"expires_after_seconds": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The number of seconds (between 10 and 1209600) after which Pix payment will expire.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"expires_at": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The timestamp at which the Pix expires.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"mandate_options": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"amount": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "Amount to be charged for future payments.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"amount_includes_iof": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Determines if the amount includes the IOF tax.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("always", "never")},
									},
									"amount_type": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Type of amount.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("fixed", "maximum")},
									},
									"currency": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"end_date": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Date when the mandate expires and no further payments will be charged, in `YYYY-MM-DD`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"payment_schedule": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Schedule at which the future payments will be charged.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("halfyearly", "monthly", "quarterly", "weekly", "yearly")},
									},
									"reference": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Subscription name displayed to buyers in their bank app.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"start_date": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Start date of the mandate, in `YYYY-MM-DD`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"promptpay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"revolut_pay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"samsung_pay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"satispay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"scalapay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"capture_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when the funds will be captured from the customer's account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"sepa_debit": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"mandate_options": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"reference_prefix": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'STRIPE'.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"target_date": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"sofort": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"preferred_language": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Preferred language of the SOFORT authorization page that the customer is redirected to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("de", "en", "es", "fr", "it", "nl", "pl")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"swish": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"reference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "A reference for this payment to be displayed in the Swish app.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"twint": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"upi": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off_session", "on_session")},
							},
							"mandate_options": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Configuration options for setting up an eMandate",
								Attributes: map[string]schema.Attribute{
									"amount": schema.Int64Attribute{
										Optional:    true,
										Description: "Amount to be charged for future payments.",
									},
									"amount_type": schema.StringAttribute{
										Optional:    true,
										Description: "One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.",
									},
									"description": schema.StringAttribute{
										Optional:    true,
										Description: "A description of the mandate or subscription that is meant to be displayed to the customer.",
									},
									"end_date": schema.Int64Attribute{
										Optional:    true,
										Description: "End date of the mandate or subscription.",
									},
								},
							},
						},
					},
					"us_bank_account": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"financial_connections": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"filters": schema.SingleNestedAttribute{
										Optional: true,
										Computed: true,

										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"account_subcategories": schema.ListAttribute{
												Optional:      true,
												Computed:      true,
												Description:   "The account subcategories to use to filter for possible accounts to link. Valid subcategories are `checking` and `savings`.",
												PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
												ElementType:   types.StringType,
											},
										},
									},
									"permissions": schema.ListAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The list of permissions to request. The `payment_method` permission must be included.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
									"prefetch": schema.ListAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Data features requested to be retrieved upon account creation.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
									"return_url": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"mandate_options": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"collection_method": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Mandate collection method",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("paper")},
									},
								},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"target_date": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_purpose": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The purpose of the transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("goods", "other", "services", "unspecified")},
							},
							"verification_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Bank account verification method. The default value is `automatic`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("automatic", "instant", "microdeposits")},
							},
							"networks": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Additional fields for network related functions",
								Attributes: map[string]schema.Attribute{
									"requested": schema.ListAttribute{
										Optional:    true,
										Description: "Triggers validations to run across the selected networks",
										ElementType: types.StringType,
									},
								},
							},
						},
					},
					"wechat_pay": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"app_id": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The app ID registered with WeChat Pay. Only required when client is ios or android.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"client": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The client type that the end customer will pay from",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("android", "ios", "web")},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"zip": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"setup_future_usage": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
				},
			},
			"presentment_details": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"presentment_amount": schema.Int64Attribute{
						Computed:      true,
						Description:   "Amount intended to be collected by this payment, denominated in `presentment_currency`.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"presentment_currency": schema.StringAttribute{
						Computed:      true,
						Description:   "Currency presented to the customer during payment.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"processing": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "If present, this property tells you about the processing state of the payment.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"card": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"customer_notification": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"approval_requested": schema.BoolAttribute{
										Computed:      true,
										Description:   "Whether customer approval has been requested for this payment. For payments greater than INR 15000 or mandate amount, the customer must provide explicit approval of the payment with their bank.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
									"completes_at": schema.Int64Attribute{
										Computed:      true,
										Description:   "If customer approval is required, they need to provide approval before this time.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "Type of the payment method for which payment is in `processing` state, one of `card`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("card")},
					},
				},
			},
			"receipt_email": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Email address that the receipt for the resulting payment will be sent to. If `receipt_email` is specified for a payment in live mode, a receipt will be sent regardless of your [email settings](https://dashboard.stripe.com/account/emails).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"review": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the review associated with this PaymentIntent, if any.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"setup_future_usage": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("off_session", "on_session")},
			},
			"shipping": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Shipping information for this PaymentIntent.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"address": schema.SingleNestedAttribute{
						Required: true,

						Attributes: map[string]schema.Attribute{
							"city": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "City, district, suburb, town, or village.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"country": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"line1": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Address line 1, such as the street, PO Box, or company name.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"line2": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Address line 2, such as the apartment, suite, unit, or building.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"postal_code": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "ZIP or postal code.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"state": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"carrier": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The delivery service that shipped a physical product, such as Fedex, UPS, USPS, etc.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"name": schema.StringAttribute{
						Required:    true,
						Description: "Recipient name.",
					},
					"phone": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Recipient phone (including extension).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"tracking_number": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"source": schema.StringAttribute{
				Computed:      true,
				Description:   "This is a legacy field that will be removed in the future. It is the ID of the Source object that is associated with this PaymentIntent, if one was supplied.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"statement_descriptor": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Text that appears on the customer's statement as the statement descriptor for a non-card charge. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).\n\nSetting this value for a card charge returns an error. For card charges, set the [statement_descriptor_suffix](https://docs.stripe.com/get-started/account/statement-descriptors#dynamic) instead.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"statement_descriptor_suffix": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Status of this PaymentIntent, one of `requires_payment_method`, `requires_confirmation`, `requires_action`, `processing`, `requires_capture`, `canceled`, or `succeeded`. Read more about each PaymentIntent [status](https://docs.stripe.com/payments/intents#intent-statuses).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("canceled", "processing", "requires_action", "requires_capture", "requires_confirmation", "requires_payment_method", "succeeded")},
			},
			"transfer_data": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The data that automatically creates a Transfer after the payment finalizes. Learn more about the [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts).",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"amount": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The amount transferred to the destination account. This transfer will occur automatically after the payment succeeds. If no amount is specified, by default the entire payment amount is transferred to the destination account.\n The amount must be less than or equal to the [amount](https://docs.stripe.com/api/payment_intents/object#payment_intent_object-amount), and must be a positive integer\n representing how much to transfer in the smallest currency unit (e.g., 100 cents to charge $1.00).",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"description": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "An arbitrary string attached to the transfer. Often useful for displaying to users.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"destination": schema.StringAttribute{
						Required:      true,
						Description:   "The account (if any) that the payment is attributed to for tax reporting, and where funds from the payment are transferred to after payment success.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"metadata": schema.MapAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"payment_data": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"description": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "An arbitrary string attached to the destination payment. Often useful for displaying to users.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"metadata": schema.MapAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
								PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
						},
					},
				},
			},
			"transfer_group": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A string that identifies the resulting payment as part of a group. Learn more about the [use case for connected accounts](https://docs.stripe.com/connect/separate-charges-and-transfers).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"confirm": schema.BoolAttribute{
				Optional:      true,
				Description:   "Set to `true` to attempt to [confirm this PaymentIntent](https://docs.stripe.com/api/payment_intents/confirm) immediately. This parameter defaults to `false`. When creating and confirming a PaymentIntent at the same time, you can also provide the parameters available in the [Confirm API](https://docs.stripe.com/api/payment_intents/confirm).",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"confirmation_token": schema.StringAttribute{
				Optional:      true,
				Description:   "ID of the ConfirmationToken used to confirm this PaymentIntent.\n\nIf the provided ConfirmationToken contains properties that are also being provided in this request, such as `payment_method`, then the values in this request will take precedence.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"error_on_requires_action": schema.BoolAttribute{
				Optional:      true,
				Description:   "Set to `true` to fail the payment attempt if the PaymentIntent transitions into `requires_action`. Use this parameter for simpler integrations that don't handle customer actions, such as [saving cards without authentication](https://docs.stripe.com/payments/save-card-without-authentication). This parameter can only be used with [`confirm=true`](https://docs.stripe.com/api/payment_intents/create#create_payment_intent-confirm).",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"mandate": schema.StringAttribute{
				Optional:      true,
				Description:   "ID of the mandate that's used for this payment. This parameter can only be used with [`confirm=true`](https://docs.stripe.com/api/payment_intents/create#create_payment_intent-confirm).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"mandate_data": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "This hash contains details about the Mandate to create. This parameter can only be used with [`confirm=true`](https://docs.stripe.com/api/payment_intents/create#create_payment_intent-confirm).",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				WriteOnly:     true,
				Attributes: map[string]schema.Attribute{
					"customer_acceptance": schema.SingleNestedAttribute{
						Required:      true,
						Description:   "This hash contains details about the customer acceptance of the Mandate.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
						WriteOnly:     true,
						Attributes: map[string]schema.Attribute{
							"accepted_at": schema.Int64Attribute{
								Optional:      true,
								Description:   "The time at which the customer accepted the Mandate.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
								WriteOnly:     true,
							},
							"online": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "If this is a Mandate accepted online, this hash contains details about the online acceptance.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								WriteOnly:     true,
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										Required:      true,
										Description:   "The IP address from which the Mandate was accepted by the customer.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										WriteOnly:     true,
									},
									"user_agent": schema.StringAttribute{
										Required:      true,
										Description:   "The user agent of the browser from which the Mandate was accepted by the customer.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										WriteOnly:     true,
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "The type of customer acceptance information included with the Mandate. One of `online` or `offline`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								WriteOnly:     true,
							},
						},
					},
				},
			},
			"payment_method_configuration": schema.StringAttribute{
				Optional:    true,
				Description: "The ID of the [payment method configuration](https://docs.stripe.com/api/payment_method_configurations) to use with this PaymentIntent.",
				WriteOnly:   true,
			},
			"payment_method_data": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "If provided, this hash will be used to create a PaymentMethod. The new PaymentMethod will appear\nin the [payment_method](https://docs.stripe.com/api/payment_intents/object#payment_intent_object-payment_method)\nproperty on the PaymentIntent.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"acss_debit": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"account_number": schema.StringAttribute{
								Required:    true,
								Description: "Customer's bank account number.",
								WriteOnly:   true,
							},
							"institution_number": schema.StringAttribute{
								Required:    true,
								Description: "Institution number of the customer's bank.",
								WriteOnly:   true,
							},
							"transit_number": schema.StringAttribute{
								Required:    true,
								Description: "Transit number of the customer's bank.",
								WriteOnly:   true,
							},
						},
					},
					"allow_redisplay": schema.StringAttribute{
						Optional:    true,
						Description: "This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow. The field defaults to `unspecified`.",
						WriteOnly:   true,
					},
					"au_becs_debit": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"account_number": schema.StringAttribute{
								Required:    true,
								Description: "The account number for the bank account.",
								WriteOnly:   true,
							},
							"bsb_number": schema.StringAttribute{
								Required:    true,
								Description: "Bank-State-Branch number of the bank account.",
								WriteOnly:   true,
							},
						},
					},
					"bacs_debit": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"account_number": schema.StringAttribute{
								Optional:    true,
								Description: "Account number of the bank account that the funds will be debited from.",
								WriteOnly:   true,
							},
							"sort_code": schema.StringAttribute{
								Optional:    true,
								Description: "Sort code of the bank account. (e.g., `10-20-30`)",
								WriteOnly:   true,
							},
						},
					},
					"billing_details": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"address": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Billing address.",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"city": schema.StringAttribute{
										Optional:    true,
										Description: "City, district, suburb, town, or village.",
										WriteOnly:   true,
									},
									"country": schema.StringAttribute{
										Optional:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
										WriteOnly:   true,
									},
									"line1": schema.StringAttribute{
										Optional:    true,
										Description: "Address line 1, such as the street, PO Box, or company name.",
										WriteOnly:   true,
									},
									"line2": schema.StringAttribute{
										Optional:    true,
										Description: "Address line 2, such as the apartment, suite, unit, or building.",
										WriteOnly:   true,
									},
									"postal_code": schema.StringAttribute{
										Optional:    true,
										Description: "ZIP or postal code.",
										WriteOnly:   true,
									},
									"state": schema.StringAttribute{
										Optional:    true,
										Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
										WriteOnly:   true,
									},
								},
							},
							"email": schema.StringAttribute{
								Optional:    true,
								Description: "Email address.",
								WriteOnly:   true,
							},
							"name": schema.StringAttribute{
								Optional:    true,
								Description: "Full name.",
								WriteOnly:   true,
							},
							"phone": schema.StringAttribute{
								Optional:    true,
								Description: "Billing phone number (including extension).",
								WriteOnly:   true,
							},
							"tax_id": schema.StringAttribute{
								Optional:    true,
								Description: "Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.",
								WriteOnly:   true,
							},
						},
					},
					"boleto": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"tax_id": schema.StringAttribute{
								Required:    true,
								Description: "The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)",
								WriteOnly:   true,
							},
						},
					},
					"eps": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"bank": schema.StringAttribute{
								Optional:    true,
								Description: "The customer's bank.",
								WriteOnly:   true,
							},
						},
					},
					"fpx": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"account_holder_type": schema.StringAttribute{
								Optional:    true,
								Description: "Account holder type for FPX transaction",
								WriteOnly:   true,
							},
							"bank": schema.StringAttribute{
								Required:    true,
								Description: "The customer's bank.",
								WriteOnly:   true,
							},
						},
					},
					"ideal": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"bank": schema.StringAttribute{
								Optional:    true,
								Description: "The customer's bank. Only use this parameter for existing customers. Don't use it for new customers.",
								WriteOnly:   true,
							},
						},
					},
					"klarna": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"dob": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Customer's date of birth",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"day": schema.Int64Attribute{
										Required:    true,
										Description: "The day of birth, between 1 and 31.",
										WriteOnly:   true,
									},
									"month": schema.Int64Attribute{
										Required:    true,
										Description: "The month of birth, between 1 and 12.",
										WriteOnly:   true,
									},
									"year": schema.Int64Attribute{
										Required:    true,
										Description: "The four-digit year of birth.",
										WriteOnly:   true,
									},
								},
							},
						},
					},
					"metadata": schema.MapAttribute{
						Optional:    true,
						Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
						WriteOnly:   true,
						ElementType: types.StringType,
					},
					"naver_pay": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"funding": schema.StringAttribute{
								Optional:    true,
								Description: "Whether to use Naver Pay points or a card to fund this transaction. If not provided, this defaults to `card`.",
								WriteOnly:   true,
							},
						},
					},
					"nz_bank_account": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"account_holder_name": schema.StringAttribute{
								Optional:    true,
								Description: "The name on the bank account. Only required if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod’s billing details.",
								WriteOnly:   true,
							},
							"account_number": schema.StringAttribute{
								Required:    true,
								Description: "The account number for the bank account.",
								WriteOnly:   true,
							},
							"bank_code": schema.StringAttribute{
								Required:    true,
								Description: "The numeric code for the bank account's bank.",
								WriteOnly:   true,
							},
							"branch_code": schema.StringAttribute{
								Required:    true,
								Description: "The numeric code for the bank account's bank branch.",
								WriteOnly:   true,
							},
							"reference": schema.StringAttribute{
								Optional: true,

								WriteOnly: true,
							},
							"suffix": schema.StringAttribute{
								Required:    true,
								Description: "The suffix of the bank account number.",
								WriteOnly:   true,
							},
						},
					},
					"p24": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"bank": schema.StringAttribute{
								Optional:    true,
								Description: "The customer's bank.",
								WriteOnly:   true,
							},
						},
					},
					"payto": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is a `payto` PaymentMethod, this hash contains details about the PayTo payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"account_number": schema.StringAttribute{
								Optional:    true,
								Description: "The account number for the bank account.",
								WriteOnly:   true,
							},
							"bsb_number": schema.StringAttribute{
								Optional:    true,
								Description: "Bank-State-Branch number of the bank account.",
								WriteOnly:   true,
							},
							"pay_id": schema.StringAttribute{
								Optional:    true,
								Description: "The PayID alias for the bank account.",
								WriteOnly:   true,
							},
						},
					},
					"radar_options": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"session": schema.StringAttribute{
								Optional:    true,
								Description: "A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.",
								WriteOnly:   true,
							},
						},
					},
					"sepa_debit": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"iban": schema.StringAttribute{
								Required:    true,
								Description: "IBAN of the bank account.",
								WriteOnly:   true,
							},
						},
					},
					"sofort": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"country": schema.StringAttribute{
								Required:    true,
								Description: "Two-letter ISO code representing the country the bank account is located in.",
								WriteOnly:   true,
							},
						},
					},
					"type": schema.StringAttribute{
						Required:    true,
						Description: "The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.",
						WriteOnly:   true,
					},
					"upi": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is a `upi` PaymentMethod, this hash contains details about the UPI payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"mandate_options": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Configuration options for setting up an eMandate",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"amount": schema.Int64Attribute{
										Optional:    true,
										Description: "Amount to be charged for future payments.",
										WriteOnly:   true,
									},
									"amount_type": schema.StringAttribute{
										Optional:    true,
										Description: "One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.",
										WriteOnly:   true,
									},
									"description": schema.StringAttribute{
										Optional:    true,
										Description: "A description of the mandate or subscription that is meant to be displayed to the customer.",
										WriteOnly:   true,
									},
									"end_date": schema.Int64Attribute{
										Optional:    true,
										Description: "End date of the mandate or subscription.",
										WriteOnly:   true,
									},
								},
							},
						},
					},
					"us_bank_account": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"account_holder_type": schema.StringAttribute{
								Optional:    true,
								Description: "Account holder type: individual or company.",
								WriteOnly:   true,
							},
							"account_number": schema.StringAttribute{
								Optional:    true,
								Description: "Account number of the bank account.",
								WriteOnly:   true,
							},
							"account_type": schema.StringAttribute{
								Optional:    true,
								Description: "Account type: checkings or savings. Defaults to checking if omitted.",
								WriteOnly:   true,
							},
							"financial_connections_account": schema.StringAttribute{
								Optional:    true,
								Description: "The ID of a Financial Connections Account to use as a payment method.",
								WriteOnly:   true,
							},
							"routing_number": schema.StringAttribute{
								Optional:    true,
								Description: "Routing number of the bank account.",
								WriteOnly:   true,
							},
						},
					},
				},
			},
			"radar_options": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "Options to configure Radar. Learn more about [Radar Sessions](https://docs.stripe.com/radar/radar-session).",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				WriteOnly:     true,
				Attributes: map[string]schema.Attribute{
					"session": schema.StringAttribute{
						Optional:      true,
						Description:   "A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
				},
			},
			"return_url": schema.StringAttribute{
				Optional:      true,
				Description:   "The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method's app or site. If you'd prefer to redirect to a mobile application, you can alternatively supply an application URI scheme. This parameter can only be used with [`confirm=true`](https://docs.stripe.com/api/payment_intents/create#create_payment_intent-confirm).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"use_stripe_sdk": schema.BoolAttribute{
				Optional:      true,
				Description:   "Set to `true` when confirming server-side and using Stripe.js, iOS, or Android client-side SDKs to handle the next actions.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
		},
	}
}

func (r *PaymentIntentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PaymentIntentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PaymentIntentResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"ErrorOnRequiresAction"}, []string{"Mandate"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"PaymentMethodConfiguration"}, []string{"PaymentMethodData"}, []string{"PaymentMethodData", "acss_debit"}, []string{"PaymentMethodData", "acss_debit", "account_number"}, []string{"PaymentMethodData", "acss_debit", "institution_number"}, []string{"PaymentMethodData", "acss_debit", "transit_number"}, []string{"PaymentMethodData", "allow_redisplay"}, []string{"PaymentMethodData", "au_becs_debit"}, []string{"PaymentMethodData", "au_becs_debit", "account_number"}, []string{"PaymentMethodData", "au_becs_debit", "bsb_number"}, []string{"PaymentMethodData", "bacs_debit"}, []string{"PaymentMethodData", "bacs_debit", "account_number"}, []string{"PaymentMethodData", "bacs_debit", "sort_code"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "billing_details", "tax_id"}, []string{"PaymentMethodData", "boleto"}, []string{"PaymentMethodData", "boleto", "tax_id"}, []string{"PaymentMethodData", "eps"}, []string{"PaymentMethodData", "eps", "bank"}, []string{"PaymentMethodData", "fpx"}, []string{"PaymentMethodData", "fpx", "account_holder_type"}, []string{"PaymentMethodData", "fpx", "bank"}, []string{"PaymentMethodData", "ideal"}, []string{"PaymentMethodData", "ideal", "bank"}, []string{"PaymentMethodData", "klarna"}, []string{"PaymentMethodData", "klarna", "dob"}, []string{"PaymentMethodData", "klarna", "dob", "day"}, []string{"PaymentMethodData", "klarna", "dob", "month"}, []string{"PaymentMethodData", "klarna", "dob", "year"}, []string{"PaymentMethodData", "metadata"}, []string{"PaymentMethodData", "naver_pay"}, []string{"PaymentMethodData", "naver_pay", "funding"}, []string{"PaymentMethodData", "nz_bank_account"}, []string{"PaymentMethodData", "nz_bank_account", "account_holder_name"}, []string{"PaymentMethodData", "nz_bank_account", "account_number"}, []string{"PaymentMethodData", "nz_bank_account", "bank_code"}, []string{"PaymentMethodData", "nz_bank_account", "branch_code"}, []string{"PaymentMethodData", "nz_bank_account", "reference"}, []string{"PaymentMethodData", "nz_bank_account", "suffix"}, []string{"PaymentMethodData", "p24"}, []string{"PaymentMethodData", "p24", "bank"}, []string{"PaymentMethodData", "payto"}, []string{"PaymentMethodData", "payto", "account_number"}, []string{"PaymentMethodData", "payto", "bsb_number"}, []string{"PaymentMethodData", "payto", "pay_id"}, []string{"PaymentMethodData", "radar_options"}, []string{"PaymentMethodData", "radar_options", "session"}, []string{"PaymentMethodData", "sepa_debit"}, []string{"PaymentMethodData", "sepa_debit", "iban"}, []string{"PaymentMethodData", "sofort"}, []string{"PaymentMethodData", "sofort", "country"}, []string{"PaymentMethodData", "type"}, []string{"PaymentMethodData", "upi"}, []string{"PaymentMethodData", "upi", "mandate_options"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodData", "upi", "mandate_options", "description"}, []string{"PaymentMethodData", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodData", "us_bank_account"}, []string{"PaymentMethodData", "us_bank_account", "account_holder_type"}, []string{"PaymentMethodData", "us_bank_account", "account_number"}, []string{"PaymentMethodData", "us_bank_account", "account_type"}, []string{"PaymentMethodData", "us_bank_account", "financial_connections_account"}, []string{"PaymentMethodData", "us_bank_account", "routing_number"}, []string{"RadarOptions"}, []string{"RadarOptions", "session"}, []string{"ReturnURL"}, []string{"UseStripeSDK"}})

	params, err := expandPaymentIntentCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentIntent create params", err.Error())
		return
	}

	obj, err := r.client.V1PaymentIntents.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating PaymentIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentIntents.B, r.client.V1PaymentIntents.Key, stripe.FormatURLPath("/v1/payment_intents/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentIntent create raw response", err.Error())
		return
	}

	if err := flattenPaymentIntent(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentIntent create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PaymentMethodOptions", "blik", "code"}, []string{"PaymentMethodOptions", "card", "cvc_token"}, []string{"PaymentMethodOptions", "card", "moto"}, []string{"PaymentMethodOptions", "card", "three_d_secure"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "ares_trans_status"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "cryptogram"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "electronic_commerce_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "exemption_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_avalgo"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_exemption"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_score"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "requestor_challenge_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "transaction_id"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "version"}, []string{"PaymentMethodOptions", "klarna", "on_demand"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "average_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "maximum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "minimum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "name"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "amount"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "date"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "reference"}, []string{"PaymentMethodOptions", "p24", "tos_shown_and_accepted"}, []string{"PaymentMethodOptions", "paypal", "risk_correlation_id"}, []string{"PaymentMethodOptions", "upi", "mandate_options"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "amount"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "description"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodOptions", "us_bank_account", "networks"}, []string{"PaymentMethodOptions", "us_bank_account", "networks", "requested"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"ErrorOnRequiresAction"}, []string{"Mandate"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"PaymentMethodConfiguration"}, []string{"PaymentMethodData"}, []string{"PaymentMethodData", "acss_debit"}, []string{"PaymentMethodData", "acss_debit", "account_number"}, []string{"PaymentMethodData", "acss_debit", "institution_number"}, []string{"PaymentMethodData", "acss_debit", "transit_number"}, []string{"PaymentMethodData", "allow_redisplay"}, []string{"PaymentMethodData", "au_becs_debit"}, []string{"PaymentMethodData", "au_becs_debit", "account_number"}, []string{"PaymentMethodData", "au_becs_debit", "bsb_number"}, []string{"PaymentMethodData", "bacs_debit"}, []string{"PaymentMethodData", "bacs_debit", "account_number"}, []string{"PaymentMethodData", "bacs_debit", "sort_code"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "billing_details", "tax_id"}, []string{"PaymentMethodData", "boleto"}, []string{"PaymentMethodData", "boleto", "tax_id"}, []string{"PaymentMethodData", "eps"}, []string{"PaymentMethodData", "eps", "bank"}, []string{"PaymentMethodData", "fpx"}, []string{"PaymentMethodData", "fpx", "account_holder_type"}, []string{"PaymentMethodData", "fpx", "bank"}, []string{"PaymentMethodData", "ideal"}, []string{"PaymentMethodData", "ideal", "bank"}, []string{"PaymentMethodData", "klarna"}, []string{"PaymentMethodData", "klarna", "dob"}, []string{"PaymentMethodData", "klarna", "dob", "day"}, []string{"PaymentMethodData", "klarna", "dob", "month"}, []string{"PaymentMethodData", "klarna", "dob", "year"}, []string{"PaymentMethodData", "metadata"}, []string{"PaymentMethodData", "naver_pay"}, []string{"PaymentMethodData", "naver_pay", "funding"}, []string{"PaymentMethodData", "nz_bank_account"}, []string{"PaymentMethodData", "nz_bank_account", "account_holder_name"}, []string{"PaymentMethodData", "nz_bank_account", "account_number"}, []string{"PaymentMethodData", "nz_bank_account", "bank_code"}, []string{"PaymentMethodData", "nz_bank_account", "branch_code"}, []string{"PaymentMethodData", "nz_bank_account", "reference"}, []string{"PaymentMethodData", "nz_bank_account", "suffix"}, []string{"PaymentMethodData", "p24"}, []string{"PaymentMethodData", "p24", "bank"}, []string{"PaymentMethodData", "payto"}, []string{"PaymentMethodData", "payto", "account_number"}, []string{"PaymentMethodData", "payto", "bsb_number"}, []string{"PaymentMethodData", "payto", "pay_id"}, []string{"PaymentMethodData", "radar_options"}, []string{"PaymentMethodData", "radar_options", "session"}, []string{"PaymentMethodData", "sepa_debit"}, []string{"PaymentMethodData", "sepa_debit", "iban"}, []string{"PaymentMethodData", "sofort"}, []string{"PaymentMethodData", "sofort", "country"}, []string{"PaymentMethodData", "type"}, []string{"PaymentMethodData", "upi"}, []string{"PaymentMethodData", "upi", "mandate_options"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodData", "upi", "mandate_options", "description"}, []string{"PaymentMethodData", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodData", "us_bank_account"}, []string{"PaymentMethodData", "us_bank_account", "account_holder_type"}, []string{"PaymentMethodData", "us_bank_account", "account_number"}, []string{"PaymentMethodData", "us_bank_account", "account_type"}, []string{"PaymentMethodData", "us_bank_account", "financial_connections_account"}, []string{"PaymentMethodData", "us_bank_account", "routing_number"}, []string{"RadarOptions"}, []string{"RadarOptions", "session"}, []string{"ReturnURL"}, []string{"UseStripeSDK"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentIntentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState PaymentIntentResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PaymentIntentResourceModel
	state = priorState

	obj, err := r.client.V1PaymentIntents.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading PaymentIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentIntents.B, r.client.V1PaymentIntents.Key, stripe.FormatURLPath("/v1/payment_intents/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentIntent raw response", err.Error())
		return
	}

	if err := flattenPaymentIntent(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentIntent read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"PaymentMethodOptions", "blik", "code"}, []string{"PaymentMethodOptions", "card", "cvc_token"}, []string{"PaymentMethodOptions", "card", "moto"}, []string{"PaymentMethodOptions", "card", "three_d_secure"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "ares_trans_status"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "cryptogram"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "electronic_commerce_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "exemption_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_avalgo"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_exemption"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_score"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "requestor_challenge_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "transaction_id"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "version"}, []string{"PaymentMethodOptions", "klarna", "on_demand"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "average_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "maximum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "minimum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "name"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "amount"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "date"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "reference"}, []string{"PaymentMethodOptions", "p24", "tos_shown_and_accepted"}, []string{"PaymentMethodOptions", "paypal", "risk_correlation_id"}, []string{"PaymentMethodOptions", "upi", "mandate_options"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "amount"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "description"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodOptions", "us_bank_account", "networks"}, []string{"PaymentMethodOptions", "us_bank_account", "networks", "requested"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"ErrorOnRequiresAction"}, []string{"Mandate"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"PaymentMethodConfiguration"}, []string{"PaymentMethodData"}, []string{"PaymentMethodData", "acss_debit"}, []string{"PaymentMethodData", "acss_debit", "account_number"}, []string{"PaymentMethodData", "acss_debit", "institution_number"}, []string{"PaymentMethodData", "acss_debit", "transit_number"}, []string{"PaymentMethodData", "allow_redisplay"}, []string{"PaymentMethodData", "au_becs_debit"}, []string{"PaymentMethodData", "au_becs_debit", "account_number"}, []string{"PaymentMethodData", "au_becs_debit", "bsb_number"}, []string{"PaymentMethodData", "bacs_debit"}, []string{"PaymentMethodData", "bacs_debit", "account_number"}, []string{"PaymentMethodData", "bacs_debit", "sort_code"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "billing_details", "tax_id"}, []string{"PaymentMethodData", "boleto"}, []string{"PaymentMethodData", "boleto", "tax_id"}, []string{"PaymentMethodData", "eps"}, []string{"PaymentMethodData", "eps", "bank"}, []string{"PaymentMethodData", "fpx"}, []string{"PaymentMethodData", "fpx", "account_holder_type"}, []string{"PaymentMethodData", "fpx", "bank"}, []string{"PaymentMethodData", "ideal"}, []string{"PaymentMethodData", "ideal", "bank"}, []string{"PaymentMethodData", "klarna"}, []string{"PaymentMethodData", "klarna", "dob"}, []string{"PaymentMethodData", "klarna", "dob", "day"}, []string{"PaymentMethodData", "klarna", "dob", "month"}, []string{"PaymentMethodData", "klarna", "dob", "year"}, []string{"PaymentMethodData", "metadata"}, []string{"PaymentMethodData", "naver_pay"}, []string{"PaymentMethodData", "naver_pay", "funding"}, []string{"PaymentMethodData", "nz_bank_account"}, []string{"PaymentMethodData", "nz_bank_account", "account_holder_name"}, []string{"PaymentMethodData", "nz_bank_account", "account_number"}, []string{"PaymentMethodData", "nz_bank_account", "bank_code"}, []string{"PaymentMethodData", "nz_bank_account", "branch_code"}, []string{"PaymentMethodData", "nz_bank_account", "reference"}, []string{"PaymentMethodData", "nz_bank_account", "suffix"}, []string{"PaymentMethodData", "p24"}, []string{"PaymentMethodData", "p24", "bank"}, []string{"PaymentMethodData", "payto"}, []string{"PaymentMethodData", "payto", "account_number"}, []string{"PaymentMethodData", "payto", "bsb_number"}, []string{"PaymentMethodData", "payto", "pay_id"}, []string{"PaymentMethodData", "radar_options"}, []string{"PaymentMethodData", "radar_options", "session"}, []string{"PaymentMethodData", "sepa_debit"}, []string{"PaymentMethodData", "sepa_debit", "iban"}, []string{"PaymentMethodData", "sofort"}, []string{"PaymentMethodData", "sofort", "country"}, []string{"PaymentMethodData", "type"}, []string{"PaymentMethodData", "upi"}, []string{"PaymentMethodData", "upi", "mandate_options"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodData", "upi", "mandate_options", "description"}, []string{"PaymentMethodData", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodData", "us_bank_account"}, []string{"PaymentMethodData", "us_bank_account", "account_holder_type"}, []string{"PaymentMethodData", "us_bank_account", "account_number"}, []string{"PaymentMethodData", "us_bank_account", "account_type"}, []string{"PaymentMethodData", "us_bank_account", "financial_connections_account"}, []string{"PaymentMethodData", "us_bank_account", "routing_number"}, []string{"RadarOptions"}, []string{"RadarOptions", "session"}, []string{"ReturnURL"}, []string{"UseStripeSDK"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentIntentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PaymentIntentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PaymentIntentResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"ErrorOnRequiresAction"}, []string{"Mandate"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"PaymentMethodConfiguration"}, []string{"PaymentMethodData"}, []string{"PaymentMethodData", "acss_debit"}, []string{"PaymentMethodData", "acss_debit", "account_number"}, []string{"PaymentMethodData", "acss_debit", "institution_number"}, []string{"PaymentMethodData", "acss_debit", "transit_number"}, []string{"PaymentMethodData", "allow_redisplay"}, []string{"PaymentMethodData", "au_becs_debit"}, []string{"PaymentMethodData", "au_becs_debit", "account_number"}, []string{"PaymentMethodData", "au_becs_debit", "bsb_number"}, []string{"PaymentMethodData", "bacs_debit"}, []string{"PaymentMethodData", "bacs_debit", "account_number"}, []string{"PaymentMethodData", "bacs_debit", "sort_code"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "billing_details", "tax_id"}, []string{"PaymentMethodData", "boleto"}, []string{"PaymentMethodData", "boleto", "tax_id"}, []string{"PaymentMethodData", "eps"}, []string{"PaymentMethodData", "eps", "bank"}, []string{"PaymentMethodData", "fpx"}, []string{"PaymentMethodData", "fpx", "account_holder_type"}, []string{"PaymentMethodData", "fpx", "bank"}, []string{"PaymentMethodData", "ideal"}, []string{"PaymentMethodData", "ideal", "bank"}, []string{"PaymentMethodData", "klarna"}, []string{"PaymentMethodData", "klarna", "dob"}, []string{"PaymentMethodData", "klarna", "dob", "day"}, []string{"PaymentMethodData", "klarna", "dob", "month"}, []string{"PaymentMethodData", "klarna", "dob", "year"}, []string{"PaymentMethodData", "metadata"}, []string{"PaymentMethodData", "naver_pay"}, []string{"PaymentMethodData", "naver_pay", "funding"}, []string{"PaymentMethodData", "nz_bank_account"}, []string{"PaymentMethodData", "nz_bank_account", "account_holder_name"}, []string{"PaymentMethodData", "nz_bank_account", "account_number"}, []string{"PaymentMethodData", "nz_bank_account", "bank_code"}, []string{"PaymentMethodData", "nz_bank_account", "branch_code"}, []string{"PaymentMethodData", "nz_bank_account", "reference"}, []string{"PaymentMethodData", "nz_bank_account", "suffix"}, []string{"PaymentMethodData", "p24"}, []string{"PaymentMethodData", "p24", "bank"}, []string{"PaymentMethodData", "payto"}, []string{"PaymentMethodData", "payto", "account_number"}, []string{"PaymentMethodData", "payto", "bsb_number"}, []string{"PaymentMethodData", "payto", "pay_id"}, []string{"PaymentMethodData", "radar_options"}, []string{"PaymentMethodData", "radar_options", "session"}, []string{"PaymentMethodData", "sepa_debit"}, []string{"PaymentMethodData", "sepa_debit", "iban"}, []string{"PaymentMethodData", "sofort"}, []string{"PaymentMethodData", "sofort", "country"}, []string{"PaymentMethodData", "type"}, []string{"PaymentMethodData", "upi"}, []string{"PaymentMethodData", "upi", "mandate_options"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodData", "upi", "mandate_options", "description"}, []string{"PaymentMethodData", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodData", "us_bank_account"}, []string{"PaymentMethodData", "us_bank_account", "account_holder_type"}, []string{"PaymentMethodData", "us_bank_account", "account_number"}, []string{"PaymentMethodData", "us_bank_account", "account_type"}, []string{"PaymentMethodData", "us_bank_account", "financial_connections_account"}, []string{"PaymentMethodData", "us_bank_account", "routing_number"}, []string{"RadarOptions"}, []string{"RadarOptions", "session"}, []string{"ReturnURL"}, []string{"UseStripeSDK"}})

	var state PaymentIntentResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"ErrorOnRequiresAction"}, []string{"Mandate"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"RadarOptions"}, []string{"RadarOptions", "session"}, []string{"ReturnURL"}, []string{"UseStripeSDK"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"ErrorOnRequiresAction"}, []string{"Mandate"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"RadarOptions"}, []string{"RadarOptions", "session"}, []string{"ReturnURL"}, []string{"UseStripeSDK"}})

	params, err := expandPaymentIntentUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentIntent update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building PaymentIntent update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1PaymentIntents.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating PaymentIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentIntents.B, r.client.V1PaymentIntents.Key, stripe.FormatURLPath("/v1/payment_intents/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentIntent update raw response", err.Error())
		return
	}

	if err := flattenPaymentIntent(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentIntent update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PaymentMethodOptions", "blik", "code"}, []string{"PaymentMethodOptions", "card", "cvc_token"}, []string{"PaymentMethodOptions", "card", "moto"}, []string{"PaymentMethodOptions", "card", "three_d_secure"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "ares_trans_status"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "cryptogram"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "electronic_commerce_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "exemption_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_avalgo"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_exemption"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_score"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "requestor_challenge_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "transaction_id"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "version"}, []string{"PaymentMethodOptions", "klarna", "on_demand"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "average_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "maximum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "minimum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "name"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "amount"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "date"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "reference"}, []string{"PaymentMethodOptions", "p24", "tos_shown_and_accepted"}, []string{"PaymentMethodOptions", "paypal", "risk_correlation_id"}, []string{"PaymentMethodOptions", "upi", "mandate_options"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "amount"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "description"}, []string{"PaymentMethodOptions", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodOptions", "us_bank_account", "networks"}, []string{"PaymentMethodOptions", "us_bank_account", "networks", "requested"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"ErrorOnRequiresAction"}, []string{"Mandate"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"PaymentMethodConfiguration"}, []string{"PaymentMethodData"}, []string{"PaymentMethodData", "acss_debit"}, []string{"PaymentMethodData", "acss_debit", "account_number"}, []string{"PaymentMethodData", "acss_debit", "institution_number"}, []string{"PaymentMethodData", "acss_debit", "transit_number"}, []string{"PaymentMethodData", "allow_redisplay"}, []string{"PaymentMethodData", "au_becs_debit"}, []string{"PaymentMethodData", "au_becs_debit", "account_number"}, []string{"PaymentMethodData", "au_becs_debit", "bsb_number"}, []string{"PaymentMethodData", "bacs_debit"}, []string{"PaymentMethodData", "bacs_debit", "account_number"}, []string{"PaymentMethodData", "bacs_debit", "sort_code"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "billing_details", "tax_id"}, []string{"PaymentMethodData", "boleto"}, []string{"PaymentMethodData", "boleto", "tax_id"}, []string{"PaymentMethodData", "eps"}, []string{"PaymentMethodData", "eps", "bank"}, []string{"PaymentMethodData", "fpx"}, []string{"PaymentMethodData", "fpx", "account_holder_type"}, []string{"PaymentMethodData", "fpx", "bank"}, []string{"PaymentMethodData", "ideal"}, []string{"PaymentMethodData", "ideal", "bank"}, []string{"PaymentMethodData", "klarna"}, []string{"PaymentMethodData", "klarna", "dob"}, []string{"PaymentMethodData", "klarna", "dob", "day"}, []string{"PaymentMethodData", "klarna", "dob", "month"}, []string{"PaymentMethodData", "klarna", "dob", "year"}, []string{"PaymentMethodData", "metadata"}, []string{"PaymentMethodData", "naver_pay"}, []string{"PaymentMethodData", "naver_pay", "funding"}, []string{"PaymentMethodData", "nz_bank_account"}, []string{"PaymentMethodData", "nz_bank_account", "account_holder_name"}, []string{"PaymentMethodData", "nz_bank_account", "account_number"}, []string{"PaymentMethodData", "nz_bank_account", "bank_code"}, []string{"PaymentMethodData", "nz_bank_account", "branch_code"}, []string{"PaymentMethodData", "nz_bank_account", "reference"}, []string{"PaymentMethodData", "nz_bank_account", "suffix"}, []string{"PaymentMethodData", "p24"}, []string{"PaymentMethodData", "p24", "bank"}, []string{"PaymentMethodData", "payto"}, []string{"PaymentMethodData", "payto", "account_number"}, []string{"PaymentMethodData", "payto", "bsb_number"}, []string{"PaymentMethodData", "payto", "pay_id"}, []string{"PaymentMethodData", "radar_options"}, []string{"PaymentMethodData", "radar_options", "session"}, []string{"PaymentMethodData", "sepa_debit"}, []string{"PaymentMethodData", "sepa_debit", "iban"}, []string{"PaymentMethodData", "sofort"}, []string{"PaymentMethodData", "sofort", "country"}, []string{"PaymentMethodData", "type"}, []string{"PaymentMethodData", "upi"}, []string{"PaymentMethodData", "upi", "mandate_options"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodData", "upi", "mandate_options", "description"}, []string{"PaymentMethodData", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodData", "us_bank_account"}, []string{"PaymentMethodData", "us_bank_account", "account_holder_type"}, []string{"PaymentMethodData", "us_bank_account", "account_number"}, []string{"PaymentMethodData", "us_bank_account", "account_type"}, []string{"PaymentMethodData", "us_bank_account", "financial_connections_account"}, []string{"PaymentMethodData", "us_bank_account", "routing_number"}, []string{"RadarOptions"}, []string{"RadarOptions", "session"}, []string{"ReturnURL"}, []string{"UseStripeSDK"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentIntentResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *PaymentIntentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandPaymentIntentCreate(plan PaymentIntentResourceModel) (*stripe.PaymentIntentCreateParams, error) {
	params := &stripe.PaymentIntentCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.ApplicationFeeAmount.IsNull() && !plan.ApplicationFeeAmount.IsUnknown() {
		params.ApplicationFeeAmount = stripe.Int64(plan.ApplicationFeeAmount.ValueInt64())
	}
	if !plan.AutomaticPaymentMethods.IsNull() && !plan.AutomaticPaymentMethods.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AutomaticPaymentMethods", plan.AutomaticPaymentMethods) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "automatic_payment_methods", params)
		}
	}
	if !plan.CaptureMethod.IsNull() && !plan.CaptureMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CaptureMethod", "CaptureMethod", plan.CaptureMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "capture_method", params)
		}
	}
	if !plan.ConfirmationMethod.IsNull() && !plan.ConfirmationMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ConfirmationMethod", "ConfirmationMethod", plan.ConfirmationMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "confirmation_method", params)
		}
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerID", "Customer", plan.Customer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.ExcludedPaymentMethodTypes.IsNull() && !plan.ExcludedPaymentMethodTypes.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ExcludedPaymentMethodTypes", plan.ExcludedPaymentMethodTypes) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "excluded_payment_method_types", params)
		}
	}
	if !plan.Hooks.IsNull() && !plan.Hooks.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Hooks", plan.Hooks) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "hooks", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.OnBehalfOf.IsNull() && !plan.OnBehalfOf.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "OnBehalfOfID", "OnBehalfOf", plan.OnBehalfOf.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "on_behalf_of", params)
		}
	}
	if !plan.PaymentDetails.IsNull() && !plan.PaymentDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentDetails", plan.PaymentDetails) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_details", params)
		}
	}
	if !plan.PaymentMethod.IsNull() && !plan.PaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethodID", "PaymentMethod", plan.PaymentMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method", params)
		}
	}
	if !plan.PaymentMethodOptions.IsNull() && !plan.PaymentMethodOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodOptions", plan.PaymentMethodOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method_options", params)
		}
	}
	if !plan.ReceiptEmail.IsNull() && !plan.ReceiptEmail.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReceiptEmail", "ReceiptEmail", plan.ReceiptEmail.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "receipt_email", params)
		}
	}
	if !plan.SetupFutureUsage.IsNull() && !plan.SetupFutureUsage.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SetupFutureUsage", "SetupFutureUsage", plan.SetupFutureUsage.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "setup_future_usage", params)
		}
	}
	if !plan.Shipping.IsNull() && !plan.Shipping.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Shipping", plan.Shipping) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping", params)
		}
	}
	if !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "statement_descriptor", params)
		}
	}
	if !plan.StatementDescriptorSuffix.IsNull() && !plan.StatementDescriptorSuffix.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptorSuffix", "StatementDescriptorSuffix", plan.StatementDescriptorSuffix.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "statement_descriptor_suffix", params)
		}
	}
	if !plan.TransferData.IsNull() && !plan.TransferData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TransferData", plan.TransferData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "transfer_data", params)
		}
	}
	if !plan.TransferGroup.IsNull() && !plan.TransferGroup.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TransferGroup", "TransferGroup", plan.TransferGroup.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "transfer_group", params)
		}
	}
	if !plan.Confirm.IsNull() && !plan.Confirm.IsUnknown() {
		params.Confirm = stripe.Bool(plan.Confirm.ValueBool())
	}
	if !plan.ConfirmationToken.IsNull() && !plan.ConfirmationToken.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ConfirmationToken", "ConfirmationToken", plan.ConfirmationToken.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "confirmation_token", params)
		}
	}
	if !plan.ErrorOnRequiresAction.IsNull() && !plan.ErrorOnRequiresAction.IsUnknown() {
		params.ErrorOnRequiresAction = stripe.Bool(plan.ErrorOnRequiresAction.ValueBool())
	}
	if !plan.Mandate.IsNull() && !plan.Mandate.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Mandate", "Mandate", plan.Mandate.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "mandate", params)
		}
	}
	if !plan.MandateData.IsNull() && !plan.MandateData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "MandateData", plan.MandateData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "mandate_data", params)
		}
	}
	if !plan.PaymentMethodConfiguration.IsNull() && !plan.PaymentMethodConfiguration.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethodConfiguration", "PaymentMethodConfiguration", plan.PaymentMethodConfiguration.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method_configuration", params)
		}
	}
	if !plan.PaymentMethodData.IsNull() && !plan.PaymentMethodData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodData", plan.PaymentMethodData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method_data", params)
		}
	}
	if !plan.RadarOptions.IsNull() && !plan.RadarOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RadarOptions", plan.RadarOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "radar_options", params)
		}
	}
	if !plan.ReturnURL.IsNull() && !plan.ReturnURL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReturnURL", "ReturnURL", plan.ReturnURL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "return_url", params)
		}
	}
	if !plan.UseStripeSDK.IsNull() && !plan.UseStripeSDK.IsUnknown() {
		params.UseStripeSDK = stripe.Bool(plan.UseStripeSDK.ValueBool())
	}

	return params, nil
}

func expandPaymentIntentUpdate(plan PaymentIntentResourceModel, state PaymentIntentResourceModel) (*stripe.PaymentIntentUpdateParams, error) {
	params := &stripe.PaymentIntentUpdateParams{}

	if !plan.Amount.Equal(state.Amount) && !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.ApplicationFeeAmount.Equal(state.ApplicationFeeAmount) && !plan.ApplicationFeeAmount.IsNull() && !plan.ApplicationFeeAmount.IsUnknown() {
		params.ApplicationFeeAmount = stripe.Int64(plan.ApplicationFeeAmount.ValueInt64())
	}
	if !plan.CaptureMethod.Equal(state.CaptureMethod) && !plan.CaptureMethod.IsNull() && !plan.CaptureMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CaptureMethod", "CaptureMethod", plan.CaptureMethod.ValueString()) {
			if !plan.CaptureMethod.Equal(state.CaptureMethod) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "capture_method", params)
			}
		}
	}
	if !plan.Currency.Equal(state.Currency) && !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			if !plan.Currency.Equal(state.Currency) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "currency", params)
			}
		}
	}
	if !plan.Customer.Equal(state.Customer) && !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerID", "Customer", plan.Customer.ValueString()) {
			if !plan.Customer.Equal(state.Customer) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "customer", params)
			}
		}
	}
	if !plan.Description.Equal(state.Description) && !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			if !plan.Description.Equal(state.Description) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "description", params)
			}
		}
	}
	if !plan.ExcludedPaymentMethodTypes.Equal(state.ExcludedPaymentMethodTypes) && !plan.ExcludedPaymentMethodTypes.IsNull() && !plan.ExcludedPaymentMethodTypes.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ExcludedPaymentMethodTypes", plan.ExcludedPaymentMethodTypes) {
			if !plan.ExcludedPaymentMethodTypes.Equal(state.ExcludedPaymentMethodTypes) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "excluded_payment_method_types", params)
			}
		}
	}
	if !plan.Hooks.Equal(state.Hooks) && !plan.Hooks.IsNull() && !plan.Hooks.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Hooks", plan.Hooks) {
			if !plan.Hooks.Equal(state.Hooks) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "hooks", params)
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
	if !plan.PaymentDetails.Equal(state.PaymentDetails) && !plan.PaymentDetails.IsNull() && !plan.PaymentDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentDetails", plan.PaymentDetails) {
			if !plan.PaymentDetails.Equal(state.PaymentDetails) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_details", params)
			}
		}
	}
	if !plan.PaymentMethod.Equal(state.PaymentMethod) && !plan.PaymentMethod.IsNull() && !plan.PaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethodID", "PaymentMethod", plan.PaymentMethod.ValueString()) {
			if !plan.PaymentMethod.Equal(state.PaymentMethod) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_method", params)
			}
		}
	}
	if !plan.PaymentMethodOptions.Equal(state.PaymentMethodOptions) && !plan.PaymentMethodOptions.IsNull() && !plan.PaymentMethodOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodOptions", plan.PaymentMethodOptions) {
			if !plan.PaymentMethodOptions.Equal(state.PaymentMethodOptions) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_method_options", params)
			}
		}
	}
	if !plan.ReceiptEmail.Equal(state.ReceiptEmail) && !plan.ReceiptEmail.IsNull() && !plan.ReceiptEmail.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReceiptEmail", "ReceiptEmail", plan.ReceiptEmail.ValueString()) {
			if !plan.ReceiptEmail.Equal(state.ReceiptEmail) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "receipt_email", params)
			}
		}
	}
	if !plan.SetupFutureUsage.Equal(state.SetupFutureUsage) && !plan.SetupFutureUsage.IsNull() && !plan.SetupFutureUsage.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SetupFutureUsage", "SetupFutureUsage", plan.SetupFutureUsage.ValueString()) {
			if !plan.SetupFutureUsage.Equal(state.SetupFutureUsage) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "setup_future_usage", params)
			}
		}
	}
	if !plan.Shipping.Equal(state.Shipping) && !plan.Shipping.IsNull() && !plan.Shipping.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Shipping", plan.Shipping) {
			if !plan.Shipping.Equal(state.Shipping) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "shipping", params)
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
	if !plan.StatementDescriptorSuffix.Equal(state.StatementDescriptorSuffix) && !plan.StatementDescriptorSuffix.IsNull() && !plan.StatementDescriptorSuffix.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptorSuffix", "StatementDescriptorSuffix", plan.StatementDescriptorSuffix.ValueString()) {
			if !plan.StatementDescriptorSuffix.Equal(state.StatementDescriptorSuffix) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "statement_descriptor_suffix", params)
			}
		}
	}
	if !plan.TransferData.Equal(state.TransferData) && !plan.TransferData.IsNull() && !plan.TransferData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TransferData", plan.TransferData) {
			if !plan.TransferData.Equal(state.TransferData) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "transfer_data", params)
			}
		}
	}
	if !plan.TransferGroup.Equal(state.TransferGroup) && !plan.TransferGroup.IsNull() && !plan.TransferGroup.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TransferGroup", "TransferGroup", plan.TransferGroup.ValueString()) {
			if !plan.TransferGroup.Equal(state.TransferGroup) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "transfer_group", params)
			}
		}
	}
	if !plan.PaymentMethodConfiguration.Equal(state.PaymentMethodConfiguration) && !plan.PaymentMethodConfiguration.IsNull() && !plan.PaymentMethodConfiguration.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethodConfiguration", "PaymentMethodConfiguration", plan.PaymentMethodConfiguration.ValueString()) {
			if !plan.PaymentMethodConfiguration.Equal(state.PaymentMethodConfiguration) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_method_configuration", params)
			}
		}
	}
	if !plan.PaymentMethodData.Equal(state.PaymentMethodData) && !plan.PaymentMethodData.IsNull() && !plan.PaymentMethodData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodData", plan.PaymentMethodData) {
			if !plan.PaymentMethodData.Equal(state.PaymentMethodData) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_method_data", params)
			}
		}
	}

	return params, nil
}

func flattenPaymentIntent(obj *stripe.PaymentIntent, state *PaymentIntentResourceModel) error {
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
		if rawValueAmount, rawOk := plainValueAtPath(raw, "amount"); rawOk {
			if valueAmount, err := flattenPlainValue(rawValueAmount, types.Int64Type, "amount", "raw response"); err != nil {
				return err
			} else {
				if typedAmount, ok := valueAmount.(types.Int64); ok {
					state.Amount = typedAmount
				}
			}
		} else if !hasRaw {
			if responseValueAmount, ok := plainFromResponseField(obj, "Amount"); ok {
				if valueAmount, err := flattenPlainValue(responseValueAmount, types.Int64Type, "amount", "response struct"); err != nil {
					return err
				} else {
					if typedAmount, ok := valueAmount.(types.Int64); ok {
						state.Amount = typedAmount
					}
				}
			}
		}
	}
	{
		if rawValueAmountCapturable, rawOk := plainValueAtPath(raw, "amount_capturable"); rawOk {
			if valueAmountCapturable, err := flattenPlainValue(rawValueAmountCapturable, types.Int64Type, "amount_capturable", "raw response"); err != nil {
				return err
			} else {
				if typedAmountCapturable, ok := valueAmountCapturable.(types.Int64); ok {
					state.AmountCapturable = typedAmountCapturable
				}
			}
		} else if !hasRaw {
			if responseValueAmountCapturable, ok := plainFromResponseField(obj, "AmountCapturable"); ok {
				if valueAmountCapturable, err := flattenPlainValue(responseValueAmountCapturable, types.Int64Type, "amount_capturable", "response struct"); err != nil {
					return err
				} else {
					if typedAmountCapturable, ok := valueAmountCapturable.(types.Int64); ok {
						state.AmountCapturable = typedAmountCapturable
					}
				}
			}
		}
	}
	{
		if rawValueAmountReceived, rawOk := plainValueAtPath(raw, "amount_received"); rawOk {
			if valueAmountReceived, err := flattenPlainValue(rawValueAmountReceived, types.Int64Type, "amount_received", "raw response"); err != nil {
				return err
			} else {
				if typedAmountReceived, ok := valueAmountReceived.(types.Int64); ok {
					state.AmountReceived = typedAmountReceived
				}
			}
		} else if !hasRaw {
			if responseValueAmountReceived, ok := plainFromResponseField(obj, "AmountReceived"); ok {
				if valueAmountReceived, err := flattenPlainValue(responseValueAmountReceived, types.Int64Type, "amount_received", "response struct"); err != nil {
					return err
				} else {
					if typedAmountReceived, ok := valueAmountReceived.(types.Int64); ok {
						state.AmountReceived = typedAmountReceived
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueApplication, rawOk := plainValueAtPath(raw, "application"); rawOk {
				if typedApplication, ok := plainToStringIDValue(rawValueApplication); ok {
					state.Application = typedApplication
				}
			} else if !hasRaw {
				if responseValueApplication, ok := plainFromResponseField(obj, "Application"); ok {
					if typedApplication, ok := plainToStringIDValue(responseValueApplication); ok {
						state.Application = typedApplication
					}
				}
			}
		}
	}
	{
		if rawValueApplicationFeeAmount, rawOk := plainValueAtPath(raw, "application_fee_amount"); rawOk {
			if valueApplicationFeeAmount, err := flattenPlainValue(rawValueApplicationFeeAmount, types.Int64Type, "application_fee_amount", "raw response"); err != nil {
				return err
			} else {
				if typedApplicationFeeAmount, ok := valueApplicationFeeAmount.(types.Int64); ok {
					state.ApplicationFeeAmount = typedApplicationFeeAmount
				}
			}
		} else if !hasRaw {
			if responseValueApplicationFeeAmount, ok := plainFromResponseField(obj, "ApplicationFeeAmount"); ok {
				if valueApplicationFeeAmount, err := flattenPlainValue(responseValueApplicationFeeAmount, types.Int64Type, "application_fee_amount", "response struct"); err != nil {
					return err
				} else {
					if typedApplicationFeeAmount, ok := valueApplicationFeeAmount.(types.Int64); ok {
						state.ApplicationFeeAmount = typedApplicationFeeAmount
					}
				}
			}
		}
	}
	{
		assignedAutomaticPaymentMethods := false
		hadRawAutomaticPaymentMethods := false
		if rawValueAutomaticPaymentMethods, rawOk := plainValueAtPath(raw, "automatic_payment_methods"); rawOk {
			hadRawAutomaticPaymentMethods = true
			if rawValueAutomaticPaymentMethods != nil {
				sourceAutomaticPaymentMethods := applyConfiguredKeyedListShapes(rawValueAutomaticPaymentMethods, attrValueToPlain(state.AutomaticPaymentMethods))
				if valueAutomaticPaymentMethods, err := flattenPlainValue(sourceAutomaticPaymentMethods, types.ObjectType{AttrTypes: map[string]attr.Type{"allow_redirects": types.StringType, "enabled": types.BoolType}}, "automatic_payment_methods", "raw response"); err != nil {
					return err
				} else {
					if typedAutomaticPaymentMethods, ok := valueAutomaticPaymentMethods.(types.Object); ok {
						state.AutomaticPaymentMethods = typedAutomaticPaymentMethods
						assignedAutomaticPaymentMethods = true
					}
				}
			}
		}
		if !assignedAutomaticPaymentMethods {
			if !hasRaw {
				if responseValueAutomaticPaymentMethods, ok := plainFromResponseField(obj, "AutomaticPaymentMethods"); ok {
					sourceAutomaticPaymentMethods := applyConfiguredKeyedListShapes(responseValueAutomaticPaymentMethods, attrValueToPlain(state.AutomaticPaymentMethods))
					if valueAutomaticPaymentMethods, err := flattenPlainValue(
						sourceAutomaticPaymentMethods,
						types.ObjectType{AttrTypes: map[string]attr.Type{"allow_redirects": types.StringType, "enabled": types.BoolType}},
						"automatic_payment_methods",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAutomaticPaymentMethods, ok := valueAutomaticPaymentMethods.(types.Object); ok {
							state.AutomaticPaymentMethods = typedAutomaticPaymentMethods
							assignedAutomaticPaymentMethods = true
						}
					}
				}
			}
		}
		if !assignedAutomaticPaymentMethods && hadRawAutomaticPaymentMethods {
			if nullAutomaticPaymentMethods, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"allow_redirects": types.StringType, "enabled": types.BoolType}}); ok {
				if typedAutomaticPaymentMethods, ok := nullAutomaticPaymentMethods.(types.Object); ok {
					state.AutomaticPaymentMethods = typedAutomaticPaymentMethods
				}
			}
		}
	}
	{
		if rawValueCanceledAt, rawOk := plainValueAtPath(raw, "canceled_at"); rawOk {
			if valueCanceledAt, err := flattenPlainValue(rawValueCanceledAt, types.Int64Type, "canceled_at", "raw response"); err != nil {
				return err
			} else {
				if typedCanceledAt, ok := valueCanceledAt.(types.Int64); ok {
					state.CanceledAt = typedCanceledAt
				}
			}
		} else if !hasRaw {
			if responseValueCanceledAt, ok := plainFromResponseField(obj, "CanceledAt"); ok {
				if valueCanceledAt, err := flattenPlainValue(responseValueCanceledAt, types.Int64Type, "canceled_at", "response struct"); err != nil {
					return err
				} else {
					if typedCanceledAt, ok := valueCanceledAt.(types.Int64); ok {
						state.CanceledAt = typedCanceledAt
					}
				}
			}
		}
	}
	{
		if rawValueCancellationReason, rawOk := plainValueAtPath(raw, "cancellation_reason"); rawOk {
			if valueCancellationReason, err := flattenPlainValue(rawValueCancellationReason, types.StringType, "cancellation_reason", "raw response"); err != nil {
				return err
			} else {
				if typedCancellationReason, ok := valueCancellationReason.(types.String); ok {
					state.CancellationReason = typedCancellationReason
				}
			}
		} else if !hasRaw {
			if responseValueCancellationReason, ok := plainFromResponseField(obj, "CancellationReason"); ok {
				if valueCancellationReason, err := flattenPlainValue(responseValueCancellationReason, types.StringType, "cancellation_reason", "response struct"); err != nil {
					return err
				} else {
					if typedCancellationReason, ok := valueCancellationReason.(types.String); ok {
						state.CancellationReason = typedCancellationReason
					}
				}
			}
		}
	}
	{
		if rawValueCaptureMethod, rawOk := plainValueAtPath(raw, "capture_method"); rawOk {
			if valueCaptureMethod, err := flattenPlainValue(rawValueCaptureMethod, types.StringType, "capture_method", "raw response"); err != nil {
				return err
			} else {
				if typedCaptureMethod, ok := valueCaptureMethod.(types.String); ok {
					state.CaptureMethod = typedCaptureMethod
				}
			}
		} else if !hasRaw {
			if responseValueCaptureMethod, ok := plainFromResponseField(obj, "CaptureMethod"); ok {
				if valueCaptureMethod, err := flattenPlainValue(responseValueCaptureMethod, types.StringType, "capture_method", "response struct"); err != nil {
					return err
				} else {
					if typedCaptureMethod, ok := valueCaptureMethod.(types.String); ok {
						state.CaptureMethod = typedCaptureMethod
					}
				}
			}
		}
	}
	{
		if rawValueClientSecret, rawOk := plainValueAtPath(raw, "client_secret"); rawOk {
			if valueClientSecret, err := flattenPlainValue(rawValueClientSecret, types.StringType, "client_secret", "raw response"); err != nil {
				return err
			} else {
				if typedClientSecret, ok := valueClientSecret.(types.String); ok {
					state.ClientSecret = typedClientSecret
				}
			}
		} else if !hasRaw {
			if responseValueClientSecret, ok := plainFromResponseField(obj, "ClientSecret"); ok {
				if valueClientSecret, err := flattenPlainValue(responseValueClientSecret, types.StringType, "client_secret", "response struct"); err != nil {
					return err
				} else {
					if typedClientSecret, ok := valueClientSecret.(types.String); ok {
						state.ClientSecret = typedClientSecret
					}
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
		if rawValueCurrency, rawOk := plainValueAtPath(raw, "currency"); rawOk {
			if valueCurrency, err := flattenPlainValue(rawValueCurrency, types.StringType, "currency", "raw response"); err != nil {
				return err
			} else {
				if typedCurrency, ok := valueCurrency.(types.String); ok {
					state.Currency = typedCurrency
				}
			}
		} else if !hasRaw {
			if responseValueCurrency, ok := plainFromResponseField(obj, "Currency"); ok {
				if valueCurrency, err := flattenPlainValue(responseValueCurrency, types.StringType, "currency", "response struct"); err != nil {
					return err
				} else {
					if typedCurrency, ok := valueCurrency.(types.String); ok {
						state.Currency = typedCurrency
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueCustomer, rawOk := plainValueAtPath(raw, "customer"); rawOk {
				if typedCustomer, ok := plainToStringIDValue(rawValueCustomer); ok {
					state.Customer = typedCustomer
				}
			} else if !hasRaw {
				if responseValueCustomer, ok := plainFromResponseField(obj, "Customer"); ok {
					if typedCustomer, ok := plainToStringIDValue(responseValueCustomer); ok {
						state.Customer = typedCustomer
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
		if rawValueExcludedPaymentMethodTypes, rawOk := plainValueAtPath(raw, "excluded_payment_method_types"); rawOk {
			if valueExcludedPaymentMethodTypes, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueExcludedPaymentMethodTypes, attrValueToPlain(state.ExcludedPaymentMethodTypes)), types.ListType{ElemType: types.StringType}, "excluded_payment_method_types", "raw response"); err != nil {
				return err
			} else {
				if typedExcludedPaymentMethodTypes, ok := valueExcludedPaymentMethodTypes.(types.List); ok {
					state.ExcludedPaymentMethodTypes = typedExcludedPaymentMethodTypes
				}
			}
		} else if !hasRaw {
			if responseValueExcludedPaymentMethodTypes, ok := plainFromResponseField(obj, "ExcludedPaymentMethodTypes"); ok {
				if valueExcludedPaymentMethodTypes, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueExcludedPaymentMethodTypes, attrValueToPlain(state.ExcludedPaymentMethodTypes)),
					types.ListType{ElemType: types.StringType},
					"excluded_payment_method_types",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedExcludedPaymentMethodTypes, ok := valueExcludedPaymentMethodTypes.(types.List); ok {
						state.ExcludedPaymentMethodTypes = typedExcludedPaymentMethodTypes
					}
				}
			}
		}
	}
	{
		assignedHooks := false
		hadRawHooks := false
		if rawValueHooks, rawOk := plainValueAtPath(raw, "hooks"); rawOk {
			hadRawHooks = true
			if rawValueHooks != nil {
				sourceHooks := applyConfiguredKeyedListShapes(rawValueHooks, attrValueToPlain(state.Hooks))
				if valueHooks, err := flattenPlainValue(sourceHooks, types.ObjectType{AttrTypes: map[string]attr.Type{"inputs": types.ObjectType{AttrTypes: map[string]attr.Type{"tax": types.ObjectType{AttrTypes: map[string]attr.Type{"calculation": types.StringType}}}}}}, "hooks", "raw response"); err != nil {
					return err
				} else {
					if typedHooks, ok := valueHooks.(types.Object); ok {
						state.Hooks = typedHooks
						assignedHooks = true
					}
				}
			}
		}
		if !assignedHooks {
			if !hasRaw {
				if responseValueHooks, ok := plainFromResponseField(obj, "Hooks"); ok {
					sourceHooks := applyConfiguredKeyedListShapes(responseValueHooks, attrValueToPlain(state.Hooks))
					if valueHooks, err := flattenPlainValue(
						sourceHooks,
						types.ObjectType{AttrTypes: map[string]attr.Type{"inputs": types.ObjectType{AttrTypes: map[string]attr.Type{"tax": types.ObjectType{AttrTypes: map[string]attr.Type{"calculation": types.StringType}}}}}},
						"hooks",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedHooks, ok := valueHooks.(types.Object); ok {
							state.Hooks = typedHooks
							assignedHooks = true
						}
					}
				}
			}
		}
		if !assignedHooks && hadRawHooks {
			if nullHooks, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"inputs": types.ObjectType{AttrTypes: map[string]attr.Type{"tax": types.ObjectType{AttrTypes: map[string]attr.Type{"calculation": types.StringType}}}}}}); ok {
				if typedHooks, ok := nullHooks.(types.Object); ok {
					state.Hooks = typedHooks
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
		assignedLastPaymentError := false
		hadRawLastPaymentError := false
		if rawValueLastPaymentError, rawOk := plainValueAtPath(raw, "last_payment_error"); rawOk {
			hadRawLastPaymentError = true
			if rawValueLastPaymentError != nil {
				sourceLastPaymentError := applyConfiguredKeyedListShapes(rawValueLastPaymentError, attrValueToPlain(state.LastPaymentError))
				if valueLastPaymentError, err := flattenPlainValue(sourceLastPaymentError, types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}}, "last_payment_error", "raw response"); err != nil {
					return err
				} else {
					if typedLastPaymentError, ok := valueLastPaymentError.(types.Object); ok {
						state.LastPaymentError = typedLastPaymentError
						assignedLastPaymentError = true
					}
				}
			}
		}
		if !assignedLastPaymentError {
			if !hasRaw {
				if responseValueLastPaymentError, ok := plainFromResponseField(obj, "LastPaymentError"); ok {
					sourceLastPaymentError := applyConfiguredKeyedListShapes(responseValueLastPaymentError, attrValueToPlain(state.LastPaymentError))
					if valueLastPaymentError, err := flattenPlainValue(
						sourceLastPaymentError,
						types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}},
						"last_payment_error",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLastPaymentError, ok := valueLastPaymentError.(types.Object); ok {
							state.LastPaymentError = typedLastPaymentError
							assignedLastPaymentError = true
						}
					}
				}
			}
		}
		if !assignedLastPaymentError && hadRawLastPaymentError {
			if nullLastPaymentError, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}}); ok {
				if typedLastPaymentError, ok := nullLastPaymentError.(types.Object); ok {
					state.LastPaymentError = typedLastPaymentError
				}
			}
		}
	}
	{
		if true {
			if rawValueLatestCharge, rawOk := plainValueAtPath(raw, "latest_charge"); rawOk {
				if typedLatestCharge, ok := plainToStringIDValue(rawValueLatestCharge); ok {
					state.LatestCharge = typedLatestCharge
				}
			} else if !hasRaw {
				if responseValueLatestCharge, ok := plainFromResponseField(obj, "LatestCharge"); ok {
					if typedLatestCharge, ok := plainToStringIDValue(responseValueLatestCharge); ok {
						state.LatestCharge = typedLatestCharge
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
		assignedManagedPayments := false
		hadRawManagedPayments := false
		if rawValueManagedPayments, rawOk := plainValueAtPath(raw, "managed_payments"); rawOk {
			hadRawManagedPayments = true
			if rawValueManagedPayments != nil {
				sourceManagedPayments := applyConfiguredKeyedListShapes(rawValueManagedPayments, attrValueToPlain(state.ManagedPayments))
				if valueManagedPayments, err := flattenPlainValue(sourceManagedPayments, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "managed_payments", "raw response"); err != nil {
					return err
				} else {
					if typedManagedPayments, ok := valueManagedPayments.(types.Object); ok {
						state.ManagedPayments = typedManagedPayments
						assignedManagedPayments = true
					}
				}
			}
		}
		if !assignedManagedPayments {
			if !hasRaw {
				if responseValueManagedPayments, ok := plainFromResponseField(obj, "ManagedPayments"); ok {
					sourceManagedPayments := applyConfiguredKeyedListShapes(responseValueManagedPayments, attrValueToPlain(state.ManagedPayments))
					if valueManagedPayments, err := flattenPlainValue(
						sourceManagedPayments,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}},
						"managed_payments",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedManagedPayments, ok := valueManagedPayments.(types.Object); ok {
							state.ManagedPayments = typedManagedPayments
							assignedManagedPayments = true
						}
					}
				}
			}
		}
		if !assignedManagedPayments && hadRawManagedPayments {
			if nullManagedPayments, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}); ok {
				if typedManagedPayments, ok := nullManagedPayments.(types.Object); ok {
					state.ManagedPayments = typedManagedPayments
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
				if valueNextAction, err := flattenPlainValue(sourceNextAction, types.ObjectType{AttrTypes: map[string]attr.Type{"alipay_handle_redirect": types.ObjectType{AttrTypes: map[string]attr.Type{"native_data": types.StringType, "native_url": types.StringType, "return_url": types.StringType, "url": types.StringType}}, "boleto_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "hosted_voucher_url": types.StringType, "number": types.StringType, "pdf": types.StringType}}, "card_await_notification": types.ObjectType{AttrTypes: map[string]attr.Type{"charge_attempt_at": types.Int64Type, "customer_approval_required": types.BoolType}}, "cashapp_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "mobile_auth_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "display_bank_transfer_instructions": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_remaining": types.Int64Type, "currency": types.StringType, "financial_addresses": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "account_type": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_name": types.StringType, "routing_number": types.StringType}}, "iban": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bic": types.StringType, "country": types.StringType, "iban": types.StringType}}, "sort_code": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "sort_code": types.StringType}}, "spei": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_code": types.StringType, "bank_name": types.StringType, "clabe": types.StringType}}, "supported_networks": types.ListType{ElemType: types.StringType}, "swift": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "account_type": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_name": types.StringType, "swift_code": types.StringType}}, "type": types.StringType, "zengin": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "account_type": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_code": types.StringType, "bank_name": types.StringType, "branch_code": types.StringType, "branch_name": types.StringType}}}}}, "hosted_instructions_url": types.StringType, "reference": types.StringType, "type": types.StringType}}, "klarna_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "konbini_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "hosted_voucher_url": types.StringType, "stores": types.ObjectType{AttrTypes: map[string]attr.Type{"familymart": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}, "lawson": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}, "ministop": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}, "seicomart": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}}}}}, "multibanco_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"entity": types.StringType, "expires_at": types.Int64Type, "hosted_voucher_url": types.StringType, "reference": types.StringType}}, "oxxo_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after": types.Int64Type, "hosted_voucher_url": types.StringType, "number": types.StringType}}, "paynow_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "pix_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "expires_at": types.Int64Type, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "promptpay_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "redirect_to_url": types.ObjectType{AttrTypes: map[string]attr.Type{"return_url": types.StringType, "url": types.StringType}}, "swish_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "mobile_auth_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "type": types.StringType, "upi_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "use_stripe_sdk": types.MapType{ElemType: types.StringType}, "verify_with_microdeposits": types.ObjectType{AttrTypes: map[string]attr.Type{"arrival_date": types.Int64Type, "hosted_verification_url": types.StringType, "microdeposit_type": types.StringType}}, "wechat_pay_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "hosted_instructions_url": types.StringType, "image_data_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "wechat_pay_redirect_to_android_app": types.ObjectType{AttrTypes: map[string]attr.Type{"app_id": types.StringType, "nonce_str": types.StringType, "package": types.StringType, "partner_id": types.StringType, "prepay_id": types.StringType, "sign": types.StringType, "timestamp": types.StringType}}, "wechat_pay_redirect_to_ios_app": types.ObjectType{AttrTypes: map[string]attr.Type{"native_url": types.StringType}}}}, "next_action", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"alipay_handle_redirect": types.ObjectType{AttrTypes: map[string]attr.Type{"native_data": types.StringType, "native_url": types.StringType, "return_url": types.StringType, "url": types.StringType}}, "boleto_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "hosted_voucher_url": types.StringType, "number": types.StringType, "pdf": types.StringType}}, "card_await_notification": types.ObjectType{AttrTypes: map[string]attr.Type{"charge_attempt_at": types.Int64Type, "customer_approval_required": types.BoolType}}, "cashapp_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "mobile_auth_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "display_bank_transfer_instructions": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_remaining": types.Int64Type, "currency": types.StringType, "financial_addresses": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "account_type": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_name": types.StringType, "routing_number": types.StringType}}, "iban": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bic": types.StringType, "country": types.StringType, "iban": types.StringType}}, "sort_code": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "sort_code": types.StringType}}, "spei": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_code": types.StringType, "bank_name": types.StringType, "clabe": types.StringType}}, "supported_networks": types.ListType{ElemType: types.StringType}, "swift": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "account_type": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_name": types.StringType, "swift_code": types.StringType}}, "type": types.StringType, "zengin": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "account_type": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_code": types.StringType, "bank_name": types.StringType, "branch_code": types.StringType, "branch_name": types.StringType}}}}}, "hosted_instructions_url": types.StringType, "reference": types.StringType, "type": types.StringType}}, "klarna_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "konbini_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "hosted_voucher_url": types.StringType, "stores": types.ObjectType{AttrTypes: map[string]attr.Type{"familymart": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}, "lawson": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}, "ministop": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}, "seicomart": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}}}}}, "multibanco_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"entity": types.StringType, "expires_at": types.Int64Type, "hosted_voucher_url": types.StringType, "reference": types.StringType}}, "oxxo_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after": types.Int64Type, "hosted_voucher_url": types.StringType, "number": types.StringType}}, "paynow_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "pix_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "expires_at": types.Int64Type, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "promptpay_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "redirect_to_url": types.ObjectType{AttrTypes: map[string]attr.Type{"return_url": types.StringType, "url": types.StringType}}, "swish_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "mobile_auth_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "type": types.StringType, "upi_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "use_stripe_sdk": types.MapType{ElemType: types.StringType}, "verify_with_microdeposits": types.ObjectType{AttrTypes: map[string]attr.Type{"arrival_date": types.Int64Type, "hosted_verification_url": types.StringType, "microdeposit_type": types.StringType}}, "wechat_pay_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "hosted_instructions_url": types.StringType, "image_data_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "wechat_pay_redirect_to_android_app": types.ObjectType{AttrTypes: map[string]attr.Type{"app_id": types.StringType, "nonce_str": types.StringType, "package": types.StringType, "partner_id": types.StringType, "prepay_id": types.StringType, "sign": types.StringType, "timestamp": types.StringType}}, "wechat_pay_redirect_to_ios_app": types.ObjectType{AttrTypes: map[string]attr.Type{"native_url": types.StringType}}}},
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
			if nullNextAction, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"alipay_handle_redirect": types.ObjectType{AttrTypes: map[string]attr.Type{"native_data": types.StringType, "native_url": types.StringType, "return_url": types.StringType, "url": types.StringType}}, "boleto_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "hosted_voucher_url": types.StringType, "number": types.StringType, "pdf": types.StringType}}, "card_await_notification": types.ObjectType{AttrTypes: map[string]attr.Type{"charge_attempt_at": types.Int64Type, "customer_approval_required": types.BoolType}}, "cashapp_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "mobile_auth_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "display_bank_transfer_instructions": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_remaining": types.Int64Type, "currency": types.StringType, "financial_addresses": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "account_type": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_name": types.StringType, "routing_number": types.StringType}}, "iban": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bic": types.StringType, "country": types.StringType, "iban": types.StringType}}, "sort_code": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "sort_code": types.StringType}}, "spei": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_code": types.StringType, "bank_name": types.StringType, "clabe": types.StringType}}, "supported_networks": types.ListType{ElemType: types.StringType}, "swift": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "account_type": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_name": types.StringType, "swift_code": types.StringType}}, "type": types.StringType, "zengin": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "account_holder_name": types.StringType, "account_number": types.StringType, "account_type": types.StringType, "bank_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "bank_code": types.StringType, "bank_name": types.StringType, "branch_code": types.StringType, "branch_name": types.StringType}}}}}, "hosted_instructions_url": types.StringType, "reference": types.StringType, "type": types.StringType}}, "klarna_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "konbini_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "hosted_voucher_url": types.StringType, "stores": types.ObjectType{AttrTypes: map[string]attr.Type{"familymart": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}, "lawson": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}, "ministop": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}, "seicomart": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "payment_code": types.StringType}}}}}}, "multibanco_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"entity": types.StringType, "expires_at": types.Int64Type, "hosted_voucher_url": types.StringType, "reference": types.StringType}}, "oxxo_display_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after": types.Int64Type, "hosted_voucher_url": types.StringType, "number": types.StringType}}, "paynow_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "pix_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "expires_at": types.Int64Type, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "promptpay_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "redirect_to_url": types.ObjectType{AttrTypes: map[string]attr.Type{"return_url": types.StringType, "url": types.StringType}}, "swish_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "mobile_auth_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "type": types.StringType, "upi_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "use_stripe_sdk": types.MapType{ElemType: types.StringType}, "verify_with_microdeposits": types.ObjectType{AttrTypes: map[string]attr.Type{"arrival_date": types.Int64Type, "hosted_verification_url": types.StringType, "microdeposit_type": types.StringType}}, "wechat_pay_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "hosted_instructions_url": types.StringType, "image_data_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "wechat_pay_redirect_to_android_app": types.ObjectType{AttrTypes: map[string]attr.Type{"app_id": types.StringType, "nonce_str": types.StringType, "package": types.StringType, "partner_id": types.StringType, "prepay_id": types.StringType, "sign": types.StringType, "timestamp": types.StringType}}, "wechat_pay_redirect_to_ios_app": types.ObjectType{AttrTypes: map[string]attr.Type{"native_url": types.StringType}}}}); ok {
				if typedNextAction, ok := nullNextAction.(types.Object); ok {
					state.NextAction = typedNextAction
				}
			}
		}
	}
	{
		if state.OnBehalfOf.IsNull() || state.OnBehalfOf.IsUnknown() {
			if rawValueOnBehalfOf, rawOk := plainValueAtPath(raw, "on_behalf_of"); rawOk {
				if typedOnBehalfOf, ok := plainToStringIDValue(rawValueOnBehalfOf); ok {
					state.OnBehalfOf = typedOnBehalfOf
				}
			} else if !hasRaw {
				if responseValueOnBehalfOf, ok := plainFromResponseField(obj, "OnBehalfOf"); ok {
					if typedOnBehalfOf, ok := plainToStringIDValue(responseValueOnBehalfOf); ok {
						state.OnBehalfOf = typedOnBehalfOf
					}
				}
			}
		}
	}
	{
		assignedPaymentDetails := false
		hadRawPaymentDetails := false
		if rawValuePaymentDetails, rawOk := plainValueAtPath(raw, "payment_details"); rawOk {
			hadRawPaymentDetails = true
			if rawValuePaymentDetails != nil {
				sourcePaymentDetails := applyConfiguredKeyedListShapes(rawValuePaymentDetails, attrValueToPlain(state.PaymentDetails))
				if valuePaymentDetails, err := flattenPlainValue(sourcePaymentDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"customer_reference": types.StringType, "order_reference": types.StringType}}, "payment_details", "raw response"); err != nil {
					return err
				} else {
					if typedPaymentDetails, ok := valuePaymentDetails.(types.Object); ok {
						state.PaymentDetails = typedPaymentDetails
						assignedPaymentDetails = true
					}
				}
			}
		}
		if !assignedPaymentDetails {
			if !hasRaw {
				if responseValuePaymentDetails, ok := plainFromResponseField(obj, "PaymentDetails"); ok {
					sourcePaymentDetails := applyConfiguredKeyedListShapes(responseValuePaymentDetails, attrValueToPlain(state.PaymentDetails))
					if valuePaymentDetails, err := flattenPlainValue(
						sourcePaymentDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"customer_reference": types.StringType, "order_reference": types.StringType}},
						"payment_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPaymentDetails, ok := valuePaymentDetails.(types.Object); ok {
							state.PaymentDetails = typedPaymentDetails
							assignedPaymentDetails = true
						}
					}
				}
			}
		}
		if !assignedPaymentDetails && hadRawPaymentDetails {
			if nullPaymentDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"customer_reference": types.StringType, "order_reference": types.StringType}}); ok {
				if typedPaymentDetails, ok := nullPaymentDetails.(types.Object); ok {
					state.PaymentDetails = typedPaymentDetails
				}
			}
		}
	}
	{
		if true {
			if rawValuePaymentMethod, rawOk := plainValueAtPath(raw, "payment_method"); rawOk {
				if typedPaymentMethod, ok := plainToStringIDValue(rawValuePaymentMethod); ok {
					state.PaymentMethod = typedPaymentMethod
				}
			} else if !hasRaw {
				if responseValuePaymentMethod, ok := plainFromResponseField(obj, "PaymentMethod"); ok {
					if typedPaymentMethod, ok := plainToStringIDValue(responseValuePaymentMethod); ok {
						state.PaymentMethod = typedPaymentMethod
					}
				}
			}
		}
	}
	{
		assignedPaymentMethodConfigurationDetails := false
		hadRawPaymentMethodConfigurationDetails := false
		if rawValuePaymentMethodConfigurationDetails, rawOk := plainValueAtPath(raw, "payment_method_configuration_details"); rawOk {
			hadRawPaymentMethodConfigurationDetails = true
			if rawValuePaymentMethodConfigurationDetails != nil {
				sourcePaymentMethodConfigurationDetails := applyConfiguredKeyedListShapes(rawValuePaymentMethodConfigurationDetails, attrValueToPlain(state.PaymentMethodConfigurationDetails))
				if valuePaymentMethodConfigurationDetails, err := flattenPlainValue(sourcePaymentMethodConfigurationDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "parent": types.StringType}}, "payment_method_configuration_details", "raw response"); err != nil {
					return err
				} else {
					if typedPaymentMethodConfigurationDetails, ok := valuePaymentMethodConfigurationDetails.(types.Object); ok {
						state.PaymentMethodConfigurationDetails = typedPaymentMethodConfigurationDetails
						assignedPaymentMethodConfigurationDetails = true
					}
				}
			}
		}
		if !assignedPaymentMethodConfigurationDetails {
			if !hasRaw {
				if responseValuePaymentMethodConfigurationDetails, ok := plainFromResponseField(obj, "PaymentMethodConfigurationDetails"); ok {
					sourcePaymentMethodConfigurationDetails := applyConfiguredKeyedListShapes(responseValuePaymentMethodConfigurationDetails, attrValueToPlain(state.PaymentMethodConfigurationDetails))
					if valuePaymentMethodConfigurationDetails, err := flattenPlainValue(
						sourcePaymentMethodConfigurationDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "parent": types.StringType}},
						"payment_method_configuration_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPaymentMethodConfigurationDetails, ok := valuePaymentMethodConfigurationDetails.(types.Object); ok {
							state.PaymentMethodConfigurationDetails = typedPaymentMethodConfigurationDetails
							assignedPaymentMethodConfigurationDetails = true
						}
					}
				}
			}
		}
		if !assignedPaymentMethodConfigurationDetails && hadRawPaymentMethodConfigurationDetails {
			if nullPaymentMethodConfigurationDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "parent": types.StringType}}); ok {
				if typedPaymentMethodConfigurationDetails, ok := nullPaymentMethodConfigurationDetails.(types.Object); ok {
					state.PaymentMethodConfigurationDetails = typedPaymentMethodConfigurationDetails
				}
			}
		}
	}
	{
		assignedPaymentMethodOptions := false
		hadRawPaymentMethodOptions := false
		if rawValuePaymentMethodOptions, rawOk := plainValueAtPath(raw, "payment_method_options"); rawOk {
			hadRawPaymentMethodOptions = true
			if rawValuePaymentMethodOptions != nil {
				sourcePaymentMethodOptions := applyConfiguredKeyedListShapes(rawValuePaymentMethodOptions, attrValueToPlain(state.PaymentMethodOptions))
				if valuePaymentMethodOptions, err := flattenPlainValue(sourcePaymentMethodOptions, types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_mandate_url": types.StringType, "interval_description": types.StringType, "payment_schedule": types.StringType, "transaction_type": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "verification_method": types.StringType}}, "affirm": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "setup_future_usage": types.StringType}}, "afterpay_clearpay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "reference": types.StringType, "setup_future_usage": types.StringType}}, "alipay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "alma": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "amazon_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "au_becs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "target_date": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType, "setup_future_usage": types.StringType}}, "billie": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "blik": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "code": types.StringType}}, "boleto": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "installments": types.ObjectType{AttrTypes: map[string]attr.Type{"available_plans": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}, "enabled": types.BoolType, "plan": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}}, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type, "interval": types.StringType, "interval_count": types.Int64Type, "reference": types.StringType, "start_date": types.Int64Type, "supported_types": types.ListType{ElemType: types.StringType}}}, "network": types.StringType, "request_extended_authorization": types.StringType, "request_incremental_authorization": types.StringType, "request_multicapture": types.StringType, "request_overcapture": types.StringType, "request_three_d_secure": types.StringType, "require_cvc_recollection": types.BoolType, "setup_future_usage": types.StringType, "statement_descriptor_suffix_kana": types.StringType, "statement_descriptor_suffix_kanji": types.StringType, "cvc_token": types.StringType, "moto": types.BoolType, "three_d_secure": types.ObjectType{AttrTypes: map[string]attr.Type{"ares_trans_status": types.StringType, "cryptogram": types.StringType, "electronic_commerce_indicator": types.StringType, "exemption_indicator": types.StringType, "network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"cartes_bancaires": types.ObjectType{AttrTypes: map[string]attr.Type{"cb_avalgo": types.StringType, "cb_exemption": types.StringType, "cb_score": types.Int64Type}}}}, "requestor_challenge_indicator": types.StringType, "transaction_id": types.StringType, "version": types.StringType}}}}, "card_present": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "request_extended_authorization": types.BoolType, "request_incremental_authorization_support": types.BoolType, "routing": types.ObjectType{AttrTypes: map[string]attr.Type{"requested_priority": types.StringType}}}}, "cashapp": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "crypto": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "requested_address_types": types.ListType{ElemType: types.StringType}, "type": types.StringType}}, "funding_type": types.StringType, "setup_future_usage": types.StringType}}, "eps": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "fpx": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "giropay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "grabpay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "ideal": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "kakao_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "setup_future_usage": types.StringType, "on_demand": types.ObjectType{AttrTypes: map[string]attr.Type{"average_amount": types.Int64Type, "maximum_amount": types.Int64Type, "minimum_amount": types.Int64Type, "purchase_interval": types.StringType, "purchase_interval_count": types.Int64Type}}, "subscriptions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "name": types.StringType, "next_billing": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "date": types.StringType}}, "reference": types.StringType}}}}}, "konbini": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "expires_after_days": types.Int64Type, "expires_at": types.Int64Type, "product_description": types.StringType, "setup_future_usage": types.StringType}}, "kr_card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "mb_way": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "mobilepay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "multibanco": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "naver_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "nz_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "target_date": types.StringType}}, "oxxo": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "p24": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "tos_shown_and_accepted": types.BoolType}}, "payco": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "paynow": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "reference": types.StringType, "setup_future_usage": types.StringType, "risk_correlation_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "payments_per_period": types.Int64Type, "purpose": types.StringType}}, "setup_future_usage": types.StringType}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_includes_iof": types.StringType, "expires_after_seconds": types.Int64Type, "expires_at": types.Int64Type, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "amount_type": types.StringType, "currency": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "reference": types.StringType, "start_date": types.StringType}}, "setup_future_usage": types.StringType}}, "promptpay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "revolut_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "samsung_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "satispay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "scalapay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "sofort": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType, "setup_future_usage": types.StringType}}, "swish": types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType, "setup_future_usage": types.StringType}}, "twint": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}, "return_url": types.StringType}}, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"collection_method": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "transaction_purpose": types.StringType, "verification_method": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.ListType{ElemType: types.StringType}}}}}, "wechat_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"app_id": types.StringType, "client": types.StringType, "setup_future_usage": types.StringType}}, "zip": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}}}, "payment_method_options", "raw response"); err != nil {
					return err
				} else {
					if typedPaymentMethodOptions, ok := valuePaymentMethodOptions.(types.Object); ok {
						state.PaymentMethodOptions = typedPaymentMethodOptions
						assignedPaymentMethodOptions = true
					}
				}
			}
		}
		if !assignedPaymentMethodOptions {
			if !hasRaw {
				if responseValuePaymentMethodOptions, ok := plainFromResponseField(obj, "PaymentMethodOptions"); ok {
					sourcePaymentMethodOptions := applyConfiguredKeyedListShapes(responseValuePaymentMethodOptions, attrValueToPlain(state.PaymentMethodOptions))
					if valuePaymentMethodOptions, err := flattenPlainValue(
						sourcePaymentMethodOptions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_mandate_url": types.StringType, "interval_description": types.StringType, "payment_schedule": types.StringType, "transaction_type": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "verification_method": types.StringType}}, "affirm": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "setup_future_usage": types.StringType}}, "afterpay_clearpay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "reference": types.StringType, "setup_future_usage": types.StringType}}, "alipay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "alma": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "amazon_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "au_becs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "target_date": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType, "setup_future_usage": types.StringType}}, "billie": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "blik": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "code": types.StringType}}, "boleto": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "installments": types.ObjectType{AttrTypes: map[string]attr.Type{"available_plans": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}, "enabled": types.BoolType, "plan": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}}, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type, "interval": types.StringType, "interval_count": types.Int64Type, "reference": types.StringType, "start_date": types.Int64Type, "supported_types": types.ListType{ElemType: types.StringType}}}, "network": types.StringType, "request_extended_authorization": types.StringType, "request_incremental_authorization": types.StringType, "request_multicapture": types.StringType, "request_overcapture": types.StringType, "request_three_d_secure": types.StringType, "require_cvc_recollection": types.BoolType, "setup_future_usage": types.StringType, "statement_descriptor_suffix_kana": types.StringType, "statement_descriptor_suffix_kanji": types.StringType, "cvc_token": types.StringType, "moto": types.BoolType, "three_d_secure": types.ObjectType{AttrTypes: map[string]attr.Type{"ares_trans_status": types.StringType, "cryptogram": types.StringType, "electronic_commerce_indicator": types.StringType, "exemption_indicator": types.StringType, "network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"cartes_bancaires": types.ObjectType{AttrTypes: map[string]attr.Type{"cb_avalgo": types.StringType, "cb_exemption": types.StringType, "cb_score": types.Int64Type}}}}, "requestor_challenge_indicator": types.StringType, "transaction_id": types.StringType, "version": types.StringType}}}}, "card_present": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "request_extended_authorization": types.BoolType, "request_incremental_authorization_support": types.BoolType, "routing": types.ObjectType{AttrTypes: map[string]attr.Type{"requested_priority": types.StringType}}}}, "cashapp": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "crypto": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "requested_address_types": types.ListType{ElemType: types.StringType}, "type": types.StringType}}, "funding_type": types.StringType, "setup_future_usage": types.StringType}}, "eps": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "fpx": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "giropay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "grabpay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "ideal": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "kakao_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "setup_future_usage": types.StringType, "on_demand": types.ObjectType{AttrTypes: map[string]attr.Type{"average_amount": types.Int64Type, "maximum_amount": types.Int64Type, "minimum_amount": types.Int64Type, "purchase_interval": types.StringType, "purchase_interval_count": types.Int64Type}}, "subscriptions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "name": types.StringType, "next_billing": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "date": types.StringType}}, "reference": types.StringType}}}}}, "konbini": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "expires_after_days": types.Int64Type, "expires_at": types.Int64Type, "product_description": types.StringType, "setup_future_usage": types.StringType}}, "kr_card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "mb_way": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "mobilepay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "multibanco": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "naver_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "nz_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "target_date": types.StringType}}, "oxxo": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "p24": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "tos_shown_and_accepted": types.BoolType}}, "payco": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "paynow": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "reference": types.StringType, "setup_future_usage": types.StringType, "risk_correlation_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "payments_per_period": types.Int64Type, "purpose": types.StringType}}, "setup_future_usage": types.StringType}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_includes_iof": types.StringType, "expires_after_seconds": types.Int64Type, "expires_at": types.Int64Type, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "amount_type": types.StringType, "currency": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "reference": types.StringType, "start_date": types.StringType}}, "setup_future_usage": types.StringType}}, "promptpay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "revolut_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "samsung_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "satispay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "scalapay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "sofort": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType, "setup_future_usage": types.StringType}}, "swish": types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType, "setup_future_usage": types.StringType}}, "twint": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}, "return_url": types.StringType}}, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"collection_method": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "transaction_purpose": types.StringType, "verification_method": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.ListType{ElemType: types.StringType}}}}}, "wechat_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"app_id": types.StringType, "client": types.StringType, "setup_future_usage": types.StringType}}, "zip": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}}},
						"payment_method_options",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPaymentMethodOptions, ok := valuePaymentMethodOptions.(types.Object); ok {
							state.PaymentMethodOptions = typedPaymentMethodOptions
							assignedPaymentMethodOptions = true
						}
					}
				}
			}
		}
		if !assignedPaymentMethodOptions && hadRawPaymentMethodOptions {
			if nullPaymentMethodOptions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_mandate_url": types.StringType, "interval_description": types.StringType, "payment_schedule": types.StringType, "transaction_type": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "verification_method": types.StringType}}, "affirm": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "setup_future_usage": types.StringType}}, "afterpay_clearpay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "reference": types.StringType, "setup_future_usage": types.StringType}}, "alipay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "alma": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "amazon_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "au_becs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "target_date": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType, "setup_future_usage": types.StringType}}, "billie": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "blik": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "code": types.StringType}}, "boleto": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "installments": types.ObjectType{AttrTypes: map[string]attr.Type{"available_plans": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}, "enabled": types.BoolType, "plan": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}}, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type, "interval": types.StringType, "interval_count": types.Int64Type, "reference": types.StringType, "start_date": types.Int64Type, "supported_types": types.ListType{ElemType: types.StringType}}}, "network": types.StringType, "request_extended_authorization": types.StringType, "request_incremental_authorization": types.StringType, "request_multicapture": types.StringType, "request_overcapture": types.StringType, "request_three_d_secure": types.StringType, "require_cvc_recollection": types.BoolType, "setup_future_usage": types.StringType, "statement_descriptor_suffix_kana": types.StringType, "statement_descriptor_suffix_kanji": types.StringType, "cvc_token": types.StringType, "moto": types.BoolType, "three_d_secure": types.ObjectType{AttrTypes: map[string]attr.Type{"ares_trans_status": types.StringType, "cryptogram": types.StringType, "electronic_commerce_indicator": types.StringType, "exemption_indicator": types.StringType, "network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"cartes_bancaires": types.ObjectType{AttrTypes: map[string]attr.Type{"cb_avalgo": types.StringType, "cb_exemption": types.StringType, "cb_score": types.Int64Type}}}}, "requestor_challenge_indicator": types.StringType, "transaction_id": types.StringType, "version": types.StringType}}}}, "card_present": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "request_extended_authorization": types.BoolType, "request_incremental_authorization_support": types.BoolType, "routing": types.ObjectType{AttrTypes: map[string]attr.Type{"requested_priority": types.StringType}}}}, "cashapp": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "crypto": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "requested_address_types": types.ListType{ElemType: types.StringType}, "type": types.StringType}}, "funding_type": types.StringType, "setup_future_usage": types.StringType}}, "eps": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "fpx": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "giropay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "grabpay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "ideal": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "kakao_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "setup_future_usage": types.StringType, "on_demand": types.ObjectType{AttrTypes: map[string]attr.Type{"average_amount": types.Int64Type, "maximum_amount": types.Int64Type, "minimum_amount": types.Int64Type, "purchase_interval": types.StringType, "purchase_interval_count": types.Int64Type}}, "subscriptions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "name": types.StringType, "next_billing": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "date": types.StringType}}, "reference": types.StringType}}}}}, "konbini": types.ObjectType{AttrTypes: map[string]attr.Type{"confirmation_number": types.StringType, "expires_after_days": types.Int64Type, "expires_at": types.Int64Type, "product_description": types.StringType, "setup_future_usage": types.StringType}}, "kr_card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "mb_way": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "mobilepay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "multibanco": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "naver_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "nz_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "target_date": types.StringType}}, "oxxo": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "p24": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "tos_shown_and_accepted": types.BoolType}}, "payco": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "paynow": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "reference": types.StringType, "setup_future_usage": types.StringType, "risk_correlation_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "payments_per_period": types.Int64Type, "purpose": types.StringType}}, "setup_future_usage": types.StringType}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_includes_iof": types.StringType, "expires_after_seconds": types.Int64Type, "expires_at": types.Int64Type, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "amount_type": types.StringType, "currency": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "reference": types.StringType, "start_date": types.StringType}}, "setup_future_usage": types.StringType}}, "promptpay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "revolut_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "samsung_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "satispay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "scalapay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "sofort": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType, "setup_future_usage": types.StringType}}, "swish": types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType, "setup_future_usage": types.StringType}}, "twint": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}, "return_url": types.StringType}}, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"collection_method": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "transaction_purpose": types.StringType, "verification_method": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.ListType{ElemType: types.StringType}}}}}, "wechat_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"app_id": types.StringType, "client": types.StringType, "setup_future_usage": types.StringType}}, "zip": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}}}); ok {
				if typedPaymentMethodOptions, ok := nullPaymentMethodOptions.(types.Object); ok {
					state.PaymentMethodOptions = typedPaymentMethodOptions
				}
			}
		}
	}
	{
		assignedPresentmentDetails := false
		hadRawPresentmentDetails := false
		if rawValuePresentmentDetails, rawOk := plainValueAtPath(raw, "presentment_details"); rawOk {
			hadRawPresentmentDetails = true
			if rawValuePresentmentDetails != nil {
				sourcePresentmentDetails := applyConfiguredKeyedListShapes(rawValuePresentmentDetails, attrValueToPlain(state.PresentmentDetails))
				if valuePresentmentDetails, err := flattenPlainValue(sourcePresentmentDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"presentment_amount": types.Int64Type, "presentment_currency": types.StringType}}, "presentment_details", "raw response"); err != nil {
					return err
				} else {
					if typedPresentmentDetails, ok := valuePresentmentDetails.(types.Object); ok {
						state.PresentmentDetails = typedPresentmentDetails
						assignedPresentmentDetails = true
					}
				}
			}
		}
		if !assignedPresentmentDetails {
			if !hasRaw {
				if responseValuePresentmentDetails, ok := plainFromResponseField(obj, "PresentmentDetails"); ok {
					sourcePresentmentDetails := applyConfiguredKeyedListShapes(responseValuePresentmentDetails, attrValueToPlain(state.PresentmentDetails))
					if valuePresentmentDetails, err := flattenPlainValue(
						sourcePresentmentDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"presentment_amount": types.Int64Type, "presentment_currency": types.StringType}},
						"presentment_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPresentmentDetails, ok := valuePresentmentDetails.(types.Object); ok {
							state.PresentmentDetails = typedPresentmentDetails
							assignedPresentmentDetails = true
						}
					}
				}
			}
		}
		if !assignedPresentmentDetails && hadRawPresentmentDetails {
			if nullPresentmentDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"presentment_amount": types.Int64Type, "presentment_currency": types.StringType}}); ok {
				if typedPresentmentDetails, ok := nullPresentmentDetails.(types.Object); ok {
					state.PresentmentDetails = typedPresentmentDetails
				}
			}
		}
	}
	{
		assignedProcessing := false
		hadRawProcessing := false
		if rawValueProcessing, rawOk := plainValueAtPath(raw, "processing"); rawOk {
			hadRawProcessing = true
			if rawValueProcessing != nil {
				sourceProcessing := applyConfiguredKeyedListShapes(rawValueProcessing, attrValueToPlain(state.Processing))
				if valueProcessing, err := flattenPlainValue(sourceProcessing, types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"customer_notification": types.ObjectType{AttrTypes: map[string]attr.Type{"approval_requested": types.BoolType, "completes_at": types.Int64Type}}}}, "type": types.StringType}}, "processing", "raw response"); err != nil {
					return err
				} else {
					if typedProcessing, ok := valueProcessing.(types.Object); ok {
						state.Processing = typedProcessing
						assignedProcessing = true
					}
				}
			}
		}
		if !assignedProcessing {
			if !hasRaw {
				if responseValueProcessing, ok := plainFromResponseField(obj, "Processing"); ok {
					sourceProcessing := applyConfiguredKeyedListShapes(responseValueProcessing, attrValueToPlain(state.Processing))
					if valueProcessing, err := flattenPlainValue(
						sourceProcessing,
						types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"customer_notification": types.ObjectType{AttrTypes: map[string]attr.Type{"approval_requested": types.BoolType, "completes_at": types.Int64Type}}}}, "type": types.StringType}},
						"processing",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedProcessing, ok := valueProcessing.(types.Object); ok {
							state.Processing = typedProcessing
							assignedProcessing = true
						}
					}
				}
			}
		}
		if !assignedProcessing && hadRawProcessing {
			if nullProcessing, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"customer_notification": types.ObjectType{AttrTypes: map[string]attr.Type{"approval_requested": types.BoolType, "completes_at": types.Int64Type}}}}, "type": types.StringType}}); ok {
				if typedProcessing, ok := nullProcessing.(types.Object); ok {
					state.Processing = typedProcessing
				}
			}
		}
	}
	{
		if rawValueReceiptEmail, rawOk := plainValueAtPath(raw, "receipt_email"); rawOk {
			if valueReceiptEmail, err := flattenPlainValue(rawValueReceiptEmail, types.StringType, "receipt_email", "raw response"); err != nil {
				return err
			} else {
				if typedReceiptEmail, ok := valueReceiptEmail.(types.String); ok {
					state.ReceiptEmail = typedReceiptEmail
				}
			}
		} else if !hasRaw {
			if responseValueReceiptEmail, ok := plainFromResponseField(obj, "ReceiptEmail"); ok {
				if valueReceiptEmail, err := flattenPlainValue(responseValueReceiptEmail, types.StringType, "receipt_email", "response struct"); err != nil {
					return err
				} else {
					if typedReceiptEmail, ok := valueReceiptEmail.(types.String); ok {
						state.ReceiptEmail = typedReceiptEmail
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueReview, rawOk := plainValueAtPath(raw, "review"); rawOk {
				if typedReview, ok := plainToStringIDValue(rawValueReview); ok {
					state.Review = typedReview
				}
			} else if !hasRaw {
				if responseValueReview, ok := plainFromResponseField(obj, "Review"); ok {
					if typedReview, ok := plainToStringIDValue(responseValueReview); ok {
						state.Review = typedReview
					}
				}
			}
		}
	}
	{
		if rawValueSetupFutureUsage, rawOk := plainValueAtPath(raw, "setup_future_usage"); rawOk {
			if valueSetupFutureUsage, err := flattenPlainValue(rawValueSetupFutureUsage, types.StringType, "setup_future_usage", "raw response"); err != nil {
				return err
			} else {
				if typedSetupFutureUsage, ok := valueSetupFutureUsage.(types.String); ok {
					state.SetupFutureUsage = typedSetupFutureUsage
				}
			}
		} else if !hasRaw {
			if responseValueSetupFutureUsage, ok := plainFromResponseField(obj, "SetupFutureUsage"); ok {
				if valueSetupFutureUsage, err := flattenPlainValue(responseValueSetupFutureUsage, types.StringType, "setup_future_usage", "response struct"); err != nil {
					return err
				} else {
					if typedSetupFutureUsage, ok := valueSetupFutureUsage.(types.String); ok {
						state.SetupFutureUsage = typedSetupFutureUsage
					}
				}
			}
		}
	}
	{
		assignedShipping := false
		hadRawShipping := false
		if rawValueShipping, rawOk := plainValueAtPath(raw, "shipping"); rawOk {
			hadRawShipping = true
			if rawValueShipping != nil {
				sourceShipping := applyConfiguredKeyedListShapes(rawValueShipping, attrValueToPlain(state.Shipping))
				if valueShipping, err := flattenPlainValue(sourceShipping, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}, "shipping", "raw response"); err != nil {
					return err
				} else {
					if typedShipping, ok := valueShipping.(types.Object); ok {
						state.Shipping = typedShipping
						assignedShipping = true
					}
				}
			}
		}
		if !assignedShipping {
			if !hasRaw {
				if responseValueShipping, ok := plainFromResponseField(obj, "Shipping"); ok {
					sourceShipping := applyConfiguredKeyedListShapes(responseValueShipping, attrValueToPlain(state.Shipping))
					if valueShipping, err := flattenPlainValue(
						sourceShipping,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}},
						"shipping",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedShipping, ok := valueShipping.(types.Object); ok {
							state.Shipping = typedShipping
							assignedShipping = true
						}
					}
				}
			}
		}
		if !assignedShipping && hadRawShipping {
			if nullShipping, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}); ok {
				if typedShipping, ok := nullShipping.(types.Object); ok {
					state.Shipping = typedShipping
				}
			}
		}
	}
	{
		if true {
			if rawValueSource, rawOk := plainValueAtPath(raw, "source"); rawOk {
				if typedSource, ok := plainToStringIDValue(rawValueSource); ok {
					state.Source = typedSource
				}
			} else if !hasRaw {
				if responseValueSource, ok := plainFromResponseField(obj, "Source"); ok {
					if typedSource, ok := plainToStringIDValue(responseValueSource); ok {
						state.Source = typedSource
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
		if rawValueStatementDescriptorSuffix, rawOk := plainValueAtPath(raw, "statement_descriptor_suffix"); rawOk {
			if valueStatementDescriptorSuffix, err := flattenPlainValue(rawValueStatementDescriptorSuffix, types.StringType, "statement_descriptor_suffix", "raw response"); err != nil {
				return err
			} else {
				if typedStatementDescriptorSuffix, ok := valueStatementDescriptorSuffix.(types.String); ok {
					state.StatementDescriptorSuffix = typedStatementDescriptorSuffix
				}
			}
		} else if !hasRaw {
			if responseValueStatementDescriptorSuffix, ok := plainFromResponseField(obj, "StatementDescriptorSuffix"); ok {
				if valueStatementDescriptorSuffix, err := flattenPlainValue(responseValueStatementDescriptorSuffix, types.StringType, "statement_descriptor_suffix", "response struct"); err != nil {
					return err
				} else {
					if typedStatementDescriptorSuffix, ok := valueStatementDescriptorSuffix.(types.String); ok {
						state.StatementDescriptorSuffix = typedStatementDescriptorSuffix
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
		assignedTransferData := false
		hadRawTransferData := false
		if rawValueTransferData, rawOk := plainValueAtPath(raw, "transfer_data"); rawOk {
			hadRawTransferData = true
			if rawValueTransferData != nil {
				sourceTransferData := applyConfiguredKeyedListShapes(rawValueTransferData, attrValueToPlain(state.TransferData))
				if valueTransferData, err := flattenPlainValue(sourceTransferData, types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "description": types.StringType, "destination": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "payment_data": types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "metadata": types.MapType{ElemType: types.StringType}}}}}, "transfer_data", "raw response"); err != nil {
					return err
				} else {
					if typedTransferData, ok := valueTransferData.(types.Object); ok {
						state.TransferData = typedTransferData
						assignedTransferData = true
					}
				}
			}
		}
		if !assignedTransferData {
			if !hasRaw {
				if responseValueTransferData, ok := plainFromResponseField(obj, "TransferData"); ok {
					sourceTransferData := applyConfiguredKeyedListShapes(responseValueTransferData, attrValueToPlain(state.TransferData))
					if valueTransferData, err := flattenPlainValue(
						sourceTransferData,
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "description": types.StringType, "destination": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "payment_data": types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "metadata": types.MapType{ElemType: types.StringType}}}}},
						"transfer_data",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTransferData, ok := valueTransferData.(types.Object); ok {
							state.TransferData = typedTransferData
							assignedTransferData = true
						}
					}
				}
			}
		}
		if !assignedTransferData && hadRawTransferData {
			if nullTransferData, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "description": types.StringType, "destination": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "payment_data": types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "metadata": types.MapType{ElemType: types.StringType}}}}}); ok {
				if typedTransferData, ok := nullTransferData.(types.Object); ok {
					state.TransferData = typedTransferData
				}
			}
		}
	}
	{
		if rawValueTransferGroup, rawOk := plainValueAtPath(raw, "transfer_group"); rawOk {
			if valueTransferGroup, err := flattenPlainValue(rawValueTransferGroup, types.StringType, "transfer_group", "raw response"); err != nil {
				return err
			} else {
				if typedTransferGroup, ok := valueTransferGroup.(types.String); ok {
					state.TransferGroup = typedTransferGroup
				}
			}
		} else if !hasRaw {
			if responseValueTransferGroup, ok := plainFromResponseField(obj, "TransferGroup"); ok {
				if valueTransferGroup, err := flattenPlainValue(responseValueTransferGroup, types.StringType, "transfer_group", "response struct"); err != nil {
					return err
				} else {
					if typedTransferGroup, ok := valueTransferGroup.(types.String); ok {
						state.TransferGroup = typedTransferGroup
					}
				}
			}
		}
	}
	return nil
}
