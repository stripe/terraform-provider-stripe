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

var _ resource.Resource = &SetupIntentResource{}

var _ resource.ResourceWithConfigure = &SetupIntentResource{}

var _ resource.ResourceWithImportState = &SetupIntentResource{}

func NewSetupIntentResource() resource.Resource {
	return &SetupIntentResource{}
}

type SetupIntentResource struct {
	client *stripe.Client
}

type SetupIntentResourceModel struct {
	Object                            types.String `tfsdk:"object"`
	Application                       types.String `tfsdk:"application"`
	AttachToSelf                      types.Bool   `tfsdk:"attach_to_self"`
	AutomaticPaymentMethods           types.Object `tfsdk:"automatic_payment_methods"`
	CancellationReason                types.String `tfsdk:"cancellation_reason"`
	ClientSecret                      types.String `tfsdk:"client_secret"`
	Created                           types.Int64  `tfsdk:"created"`
	Customer                          types.String `tfsdk:"customer"`
	CustomerAccount                   types.String `tfsdk:"customer_account"`
	Description                       types.String `tfsdk:"description"`
	ExcludedPaymentMethodTypes        types.List   `tfsdk:"excluded_payment_method_types"`
	FlowDirections                    types.List   `tfsdk:"flow_directions"`
	ID                                types.String `tfsdk:"id"`
	LastSetupError                    types.Object `tfsdk:"last_setup_error"`
	LatestAttempt                     types.String `tfsdk:"latest_attempt"`
	Livemode                          types.Bool   `tfsdk:"livemode"`
	ManagedPayments                   types.Object `tfsdk:"managed_payments"`
	Mandate                           types.String `tfsdk:"mandate"`
	Metadata                          types.Map    `tfsdk:"metadata"`
	NextAction                        types.Object `tfsdk:"next_action"`
	OnBehalfOf                        types.String `tfsdk:"on_behalf_of"`
	PaymentMethod                     types.String `tfsdk:"payment_method"`
	PaymentMethodConfigurationDetails types.Object `tfsdk:"payment_method_configuration_details"`
	PaymentMethodOptions              types.Object `tfsdk:"payment_method_options"`
	PaymentMethodTypes                types.List   `tfsdk:"payment_method_types"`
	SingleUseMandate                  types.String `tfsdk:"single_use_mandate"`
	Status                            types.String `tfsdk:"status"`
	Usage                             types.String `tfsdk:"usage"`
	Confirm                           types.Bool   `tfsdk:"confirm"`
	ConfirmationToken                 types.String `tfsdk:"confirmation_token"`
	MandateData                       types.Object `tfsdk:"mandate_data"`
	PaymentMethodConfiguration        types.String `tfsdk:"payment_method_configuration"`
	PaymentMethodData                 types.Object `tfsdk:"payment_method_data"`
	ReturnURL                         types.String `tfsdk:"return_url"`
	SingleUse                         types.Object `tfsdk:"single_use"`
	UseStripeSDK                      types.Bool   `tfsdk:"use_stripe_sdk"`
}

func (r *SetupIntentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SetupIntentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_setup_intent"
}

func (r *SetupIntentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A SetupIntent guides you through the process of setting up and saving a customer's payment credentials for future payments.\nFor example, you can use a SetupIntent to set up and save your customer's card without immediately collecting a payment.\nLater, you can use [PaymentIntents](https://api.stripe.com#payment_intents) to drive the payment flow.\n\nCreate a SetupIntent when you're ready to collect your customer's payment credentials.\nDon't maintain long-lived, unconfirmed SetupIntents because they might not be valid.\nThe SetupIntent transitions through multiple [statuses](https://docs.stripe.com/payments/intents#intent-statuses) as it guides\nyou through the setup process.\n\nSuccessful SetupIntents result in payment credentials that are optimized for future payments.\nFor example, cardholders in [certain regions](https://stripe.com/guides/strong-customer-authentication) might need to be run through\n[Strong Customer Authentication](https://docs.stripe.com/strong-customer-authentication) during payment method collection\nto streamline later [off-session payments](https://docs.stripe.com/payments/setup-intents).\nIf you use the SetupIntent with a [Customer](https://api.stripe.com#setup_intent_object-customer),\nit automatically attaches the resulting payment method to that Customer after successful setup.\nWe recommend using SetupIntents or [setup_future_usage](https://api.stripe.com#payment_intent_object-setup_future_usage) on\nPaymentIntents to save payment methods to prevent saving invalid or unoptimized payment methods.\n\nBy using SetupIntents, you can reduce friction for your customers, even as regulations change over time.\n\nRelated guide: [Setup Intents API](https://docs.stripe.com/payments/setup-intents)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("setup_intent")},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the Connect application that created the SetupIntent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"attach_to_self": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.\n\nIt can only be used for this Stripe Account’s own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"automatic_payment_methods": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Settings for dynamic payment methods compatible with this Setup Intent",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"allow_redirects": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Controls whether this SetupIntent will accept redirect-based payment methods.\n\nRedirect-based payment methods may require your customer to be redirected to a payment method's app or site for authentication or additional steps. To [confirm](https://docs.stripe.com/api/setup_intents/confirm) this SetupIntent, you may be required to provide a `return_url` to redirect customers back to your site after they authenticate or complete the setup.",
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
			"cancellation_reason": schema.StringAttribute{
				Computed:      true,
				Description:   "Reason for cancellation of this SetupIntent, one of `abandoned`, `requested_by_customer`, or `duplicate`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("abandoned", "duplicate", "requested_by_customer")},
			},
			"client_secret": schema.StringAttribute{
				Computed:      true,
				Description:   "The client secret of this SetupIntent. Used for client-side retrieval using a publishable key.\n\nThe client secret can be used to complete payment setup from your frontend. It should not be stored, logged, or exposed to anyone other than the customer. Make sure that you have TLS enabled on any page that includes the client secret.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the Customer this SetupIntent belongs to, if one exists.\n\nIf present, the SetupIntent's payment method will be attached to the Customer on successful setup. Payment methods attached to other Customers cannot be used with this SetupIntent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_account": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the Account this SetupIntent belongs to, if one exists.\n\nIf present, the SetupIntent's payment method will be attached to the Account on successful setup. Payment methods attached to other Accounts cannot be used with this SetupIntent.",
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
				Description:   "Payment method types that are excluded from this SetupIntent.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"flow_directions": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Indicates the directions of money movement for which this payment method is intended to be used.\n\nInclude `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"last_setup_error": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The error encountered in the previous SetupIntent confirmation.",
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
			"latest_attempt": schema.StringAttribute{
				Computed:      true,
				Description:   "The most recent SetupAttempt for this SetupIntent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"managed_payments": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Computed:      true,
						Description:   "Set to `true` to enable [Managed Payments](https://docs.stripe.com/payments/managed-payments), Stripe's merchant of record solution, for this session.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"mandate": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the multi use Mandate generated by the SetupIntent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"next_action": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "If present, this property tells you what actions you need to take in order for your customer to continue payment setup.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
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
								Description:   "The URL you must redirect your customer to in order to authenticate.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
						Description:   "When confirming a SetupIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js.",
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
				},
			},
			"on_behalf_of": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account (if any) for which the setup is intended.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"payment_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the payment method used with this SetupIntent. If the payment method is `card_present` and isn't a digital wallet, then the [generated_card](https://docs.stripe.com/api/setup_attempts/object#setup_attempt_object-payment_method_details-card_present-generated_card) associated with the `latest_attempt` is attached to the Customer instead.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"payment_method_configuration_details": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Information about the [payment method configuration](https://docs.stripe.com/api/payment_method_configurations) used for this Setup Intent.",
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
				Description:   "Payment method-specific configuration for this SetupIntent.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"acss_debit": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"currency": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Currency supported by the bank account",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("cad", "usd")},
							},
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
									"default_for": schema.ListAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "List of Stripe products where this mandate can be selected automatically.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
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
							"verification_method": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Bank account verification method. The default value is `automatic`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("automatic", "instant", "microdeposits")},
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
						},
					},
					"card": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
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
									"currency": schema.StringAttribute{
										Required:    true,
										Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
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
								Description:   "Selected network to process this SetupIntent on. Depends on the available networks of the card attached to the setup intent. Can be only set confirm-time.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("amex", "cartes_bancaires", "diners", "discover", "eftpos_au", "girocard", "interac", "jcb", "link", "mastercard", "unionpay", "unknown", "visa")},
							},
							"request_three_d_secure": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("any", "automatic", "challenge")},
							},
							"moto": schema.BoolAttribute{
								Optional:    true,
								Description: "When specified, this parameter signals that a card has been collected\nas MOTO (Mail Order Telephone Order) and thus out of scope for SCA. This\nparameter can only be provided during confirmation.",
							},
							"three_d_secure": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "If 3D Secure authentication was performed with a third-party provider,\nthe authentication details to use for this setup.",
								Attributes: map[string]schema.Attribute{
									"ares_trans_status": schema.StringAttribute{
										Optional:    true,
										Description: "The `transStatus` returned from the card Issuer’s ACS in the ARes.",
									},
									"cryptogram": schema.StringAttribute{
										Optional:    true,
										Description: "The cryptogram, also known as the \"authentication value\" (AAV, CAVV or\nAEVV). This value is 20 bytes, base64-encoded into a 28-character string.\n(Most 3D Secure providers will return the base64-encoded version, which\nis what you should specify here.)",
									},
									"electronic_commerce_indicator": schema.StringAttribute{
										Optional:    true,
										Description: "The Electronic Commerce Indicator (ECI) is returned by your 3D Secure\nprovider and indicates what degree of authentication was performed.",
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
										Optional:    true,
										Description: "For 3D Secure 1, the XID. For 3D Secure 2, the Directory Server\nTransaction ID (dsTransID).",
									},
									"version": schema.StringAttribute{
										Optional:    true,
										Description: "The version of 3D Secure that was performed.",
									},
								},
							},
						},
					},
					"klarna": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"currency": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The currency of the setup intent. Three letter ISO currency code.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"preferred_locale": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Preferred locale of the Klarna checkout page that the customer is redirected to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"on_demand": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "On-demand details if setting up a payment method for on-demand payments.",
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
								Description: "Subscription details if setting up or charging a subscription",
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
											Required:    true,
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
					"link": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"persistent_token": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "[Deprecated] This is a legacy parameter that no longer has any function.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"paypal": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"billing_agreement_id": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The PayPal Billing Agreement ID (BAID). This is an ID generated by PayPal which represents the mandate between the merchant and the customer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
									"start_date": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Date, in YYYY-MM-DD format, from which payments will be collected. Defaults to confirmation time.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
					"pix": schema.SingleNestedAttribute{
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
						},
					},
					"upi": schema.SingleNestedAttribute{
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
										Description:   "Amount to be charged for future payments.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"amount_type": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("fixed", "maximum")},
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
										Description:   "End date of the mandate or subscription.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
								},
							},
							"setup_future_usage": schema.StringAttribute{
								Optional: true,

								Validators: []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
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
				},
			},
			"payment_method_types": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The list of payment method types (e.g. card) that this SetupIntent is allowed to set up. A list of valid payment method types can be found [here](https://docs.stripe.com/api/payment_methods/object#payment_method_object-type).",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"single_use_mandate": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the single_use Mandate generated by the SetupIntent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "[Status](https://docs.stripe.com/payments/intents#intent-statuses) of this SetupIntent, one of `requires_payment_method`, `requires_confirmation`, `requires_action`, `processing`, `canceled`, or `succeeded`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("canceled", "processing", "requires_action", "requires_confirmation", "requires_payment_method", "succeeded")},
			},
			"usage": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Indicates how the payment method is intended to be used in the future.\n\nUse `on_session` if you intend to only reuse the payment method when the customer is in your checkout flow. Use `off_session` if your customer may or may not be in your checkout flow. If not provided, this value defaults to `off_session`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"confirm": schema.BoolAttribute{
				Optional:      true,
				Description:   "Set to `true` to attempt to confirm this SetupIntent immediately. This parameter defaults to `false`. If a card is the attached payment method, you can provide a `return_url` in case further authentication is necessary.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"confirmation_token": schema.StringAttribute{
				Optional:      true,
				Description:   "ID of the ConfirmationToken used to confirm this SetupIntent.\n\nIf the provided ConfirmationToken contains properties that are also being provided in this request, such as `payment_method`, then the values in this request will take precedence.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"mandate_data": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "This hash contains details about the mandate to create. This parameter can only be used with [`confirm=true`](https://docs.stripe.com/api/setup_intents/create#create_setup_intent-confirm).",
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
				Description: "The ID of the [payment method configuration](https://docs.stripe.com/api/payment_method_configurations) to use with this SetupIntent.",
				WriteOnly:   true,
			},
			"payment_method_data": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "When included, this hash creates a PaymentMethod that is set as the [`payment_method`](https://docs.stripe.com/api/setup_intents/object#setup_intent_object-payment_method)\nvalue in the SetupIntent.",
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
			"return_url": schema.StringAttribute{
				Optional:      true,
				Description:   "The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method's app or site. To redirect to a mobile application, you can alternatively supply an application URI scheme. This parameter can only be used with [`confirm=true`](https://docs.stripe.com/api/setup_intents/create#create_setup_intent-confirm).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"single_use": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "If you populate this hash, this SetupIntent generates a `single_use` mandate after successful completion.\n\nSingle-use mandates are only valid for the following payment methods: `acss_debit`, `alipay`, `au_becs_debit`, `bacs_debit`, `bancontact`, `boleto`, `ideal`, `link`, `sepa_debit`, and `us_bank_account`.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				WriteOnly:     true,
				Attributes: map[string]schema.Attribute{
					"amount": schema.Int64Attribute{
						Required:      true,
						Description:   "Amount the customer is granting permission to collect later. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://docs.stripe.com/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
					"currency": schema.StringAttribute{
						Required:      true,
						Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
				},
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

func (r *SetupIntentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SetupIntentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config SetupIntentResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"PaymentMethodConfiguration"}, []string{"PaymentMethodData"}, []string{"PaymentMethodData", "acss_debit"}, []string{"PaymentMethodData", "acss_debit", "account_number"}, []string{"PaymentMethodData", "acss_debit", "institution_number"}, []string{"PaymentMethodData", "acss_debit", "transit_number"}, []string{"PaymentMethodData", "allow_redisplay"}, []string{"PaymentMethodData", "au_becs_debit"}, []string{"PaymentMethodData", "au_becs_debit", "account_number"}, []string{"PaymentMethodData", "au_becs_debit", "bsb_number"}, []string{"PaymentMethodData", "bacs_debit"}, []string{"PaymentMethodData", "bacs_debit", "account_number"}, []string{"PaymentMethodData", "bacs_debit", "sort_code"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "billing_details", "tax_id"}, []string{"PaymentMethodData", "boleto"}, []string{"PaymentMethodData", "boleto", "tax_id"}, []string{"PaymentMethodData", "eps"}, []string{"PaymentMethodData", "eps", "bank"}, []string{"PaymentMethodData", "fpx"}, []string{"PaymentMethodData", "fpx", "account_holder_type"}, []string{"PaymentMethodData", "fpx", "bank"}, []string{"PaymentMethodData", "ideal"}, []string{"PaymentMethodData", "ideal", "bank"}, []string{"PaymentMethodData", "klarna"}, []string{"PaymentMethodData", "klarna", "dob"}, []string{"PaymentMethodData", "klarna", "dob", "day"}, []string{"PaymentMethodData", "klarna", "dob", "month"}, []string{"PaymentMethodData", "klarna", "dob", "year"}, []string{"PaymentMethodData", "metadata"}, []string{"PaymentMethodData", "naver_pay"}, []string{"PaymentMethodData", "naver_pay", "funding"}, []string{"PaymentMethodData", "nz_bank_account"}, []string{"PaymentMethodData", "nz_bank_account", "account_holder_name"}, []string{"PaymentMethodData", "nz_bank_account", "account_number"}, []string{"PaymentMethodData", "nz_bank_account", "bank_code"}, []string{"PaymentMethodData", "nz_bank_account", "branch_code"}, []string{"PaymentMethodData", "nz_bank_account", "reference"}, []string{"PaymentMethodData", "nz_bank_account", "suffix"}, []string{"PaymentMethodData", "p24"}, []string{"PaymentMethodData", "p24", "bank"}, []string{"PaymentMethodData", "payto"}, []string{"PaymentMethodData", "payto", "account_number"}, []string{"PaymentMethodData", "payto", "bsb_number"}, []string{"PaymentMethodData", "payto", "pay_id"}, []string{"PaymentMethodData", "radar_options"}, []string{"PaymentMethodData", "radar_options", "session"}, []string{"PaymentMethodData", "sepa_debit"}, []string{"PaymentMethodData", "sepa_debit", "iban"}, []string{"PaymentMethodData", "sofort"}, []string{"PaymentMethodData", "sofort", "country"}, []string{"PaymentMethodData", "type"}, []string{"PaymentMethodData", "upi"}, []string{"PaymentMethodData", "upi", "mandate_options"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodData", "upi", "mandate_options", "description"}, []string{"PaymentMethodData", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodData", "us_bank_account"}, []string{"PaymentMethodData", "us_bank_account", "account_holder_type"}, []string{"PaymentMethodData", "us_bank_account", "account_number"}, []string{"PaymentMethodData", "us_bank_account", "account_type"}, []string{"PaymentMethodData", "us_bank_account", "financial_connections_account"}, []string{"PaymentMethodData", "us_bank_account", "routing_number"}, []string{"ReturnURL"}, []string{"SingleUse"}, []string{"SingleUse", "amount"}, []string{"SingleUse", "currency"}, []string{"UseStripeSDK"}})

	params, err := expandSetupIntentCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building SetupIntent create params", err.Error())
		return
	}

	obj, err := r.client.V1SetupIntents.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating SetupIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1SetupIntents.B, r.client.V1SetupIntents.Key, stripe.FormatURLPath("/v1/setup_intents/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating SetupIntent create raw response", err.Error())
		return
	}

	if err := flattenSetupIntent(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening SetupIntent create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PaymentMethodOptions", "card", "moto"}, []string{"PaymentMethodOptions", "card", "three_d_secure"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "ares_trans_status"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "cryptogram"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "electronic_commerce_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_avalgo"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_exemption"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_score"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "requestor_challenge_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "transaction_id"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "version"}, []string{"PaymentMethodOptions", "klarna", "on_demand"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "average_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "maximum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "minimum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "name"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "amount"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "date"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "reference"}, []string{"PaymentMethodOptions", "upi", "setup_future_usage"}, []string{"PaymentMethodOptions", "us_bank_account", "networks"}, []string{"PaymentMethodOptions", "us_bank_account", "networks", "requested"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"PaymentMethodConfiguration"}, []string{"PaymentMethodData"}, []string{"PaymentMethodData", "acss_debit"}, []string{"PaymentMethodData", "acss_debit", "account_number"}, []string{"PaymentMethodData", "acss_debit", "institution_number"}, []string{"PaymentMethodData", "acss_debit", "transit_number"}, []string{"PaymentMethodData", "allow_redisplay"}, []string{"PaymentMethodData", "au_becs_debit"}, []string{"PaymentMethodData", "au_becs_debit", "account_number"}, []string{"PaymentMethodData", "au_becs_debit", "bsb_number"}, []string{"PaymentMethodData", "bacs_debit"}, []string{"PaymentMethodData", "bacs_debit", "account_number"}, []string{"PaymentMethodData", "bacs_debit", "sort_code"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "billing_details", "tax_id"}, []string{"PaymentMethodData", "boleto"}, []string{"PaymentMethodData", "boleto", "tax_id"}, []string{"PaymentMethodData", "eps"}, []string{"PaymentMethodData", "eps", "bank"}, []string{"PaymentMethodData", "fpx"}, []string{"PaymentMethodData", "fpx", "account_holder_type"}, []string{"PaymentMethodData", "fpx", "bank"}, []string{"PaymentMethodData", "ideal"}, []string{"PaymentMethodData", "ideal", "bank"}, []string{"PaymentMethodData", "klarna"}, []string{"PaymentMethodData", "klarna", "dob"}, []string{"PaymentMethodData", "klarna", "dob", "day"}, []string{"PaymentMethodData", "klarna", "dob", "month"}, []string{"PaymentMethodData", "klarna", "dob", "year"}, []string{"PaymentMethodData", "metadata"}, []string{"PaymentMethodData", "naver_pay"}, []string{"PaymentMethodData", "naver_pay", "funding"}, []string{"PaymentMethodData", "nz_bank_account"}, []string{"PaymentMethodData", "nz_bank_account", "account_holder_name"}, []string{"PaymentMethodData", "nz_bank_account", "account_number"}, []string{"PaymentMethodData", "nz_bank_account", "bank_code"}, []string{"PaymentMethodData", "nz_bank_account", "branch_code"}, []string{"PaymentMethodData", "nz_bank_account", "reference"}, []string{"PaymentMethodData", "nz_bank_account", "suffix"}, []string{"PaymentMethodData", "p24"}, []string{"PaymentMethodData", "p24", "bank"}, []string{"PaymentMethodData", "payto"}, []string{"PaymentMethodData", "payto", "account_number"}, []string{"PaymentMethodData", "payto", "bsb_number"}, []string{"PaymentMethodData", "payto", "pay_id"}, []string{"PaymentMethodData", "radar_options"}, []string{"PaymentMethodData", "radar_options", "session"}, []string{"PaymentMethodData", "sepa_debit"}, []string{"PaymentMethodData", "sepa_debit", "iban"}, []string{"PaymentMethodData", "sofort"}, []string{"PaymentMethodData", "sofort", "country"}, []string{"PaymentMethodData", "type"}, []string{"PaymentMethodData", "upi"}, []string{"PaymentMethodData", "upi", "mandate_options"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodData", "upi", "mandate_options", "description"}, []string{"PaymentMethodData", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodData", "us_bank_account"}, []string{"PaymentMethodData", "us_bank_account", "account_holder_type"}, []string{"PaymentMethodData", "us_bank_account", "account_number"}, []string{"PaymentMethodData", "us_bank_account", "account_type"}, []string{"PaymentMethodData", "us_bank_account", "financial_connections_account"}, []string{"PaymentMethodData", "us_bank_account", "routing_number"}, []string{"ReturnURL"}, []string{"SingleUse"}, []string{"SingleUse", "amount"}, []string{"SingleUse", "currency"}, []string{"UseStripeSDK"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SetupIntentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState SetupIntentResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state SetupIntentResourceModel
	state = priorState

	obj, err := r.client.V1SetupIntents.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading SetupIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1SetupIntents.B, r.client.V1SetupIntents.Key, stripe.FormatURLPath("/v1/setup_intents/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating SetupIntent raw response", err.Error())
		return
	}

	if err := flattenSetupIntent(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening SetupIntent read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"PaymentMethodOptions", "card", "moto"}, []string{"PaymentMethodOptions", "card", "three_d_secure"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "ares_trans_status"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "cryptogram"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "electronic_commerce_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_avalgo"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_exemption"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_score"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "requestor_challenge_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "transaction_id"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "version"}, []string{"PaymentMethodOptions", "klarna", "on_demand"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "average_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "maximum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "minimum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "name"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "amount"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "date"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "reference"}, []string{"PaymentMethodOptions", "upi", "setup_future_usage"}, []string{"PaymentMethodOptions", "us_bank_account", "networks"}, []string{"PaymentMethodOptions", "us_bank_account", "networks", "requested"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"PaymentMethodConfiguration"}, []string{"PaymentMethodData"}, []string{"PaymentMethodData", "acss_debit"}, []string{"PaymentMethodData", "acss_debit", "account_number"}, []string{"PaymentMethodData", "acss_debit", "institution_number"}, []string{"PaymentMethodData", "acss_debit", "transit_number"}, []string{"PaymentMethodData", "allow_redisplay"}, []string{"PaymentMethodData", "au_becs_debit"}, []string{"PaymentMethodData", "au_becs_debit", "account_number"}, []string{"PaymentMethodData", "au_becs_debit", "bsb_number"}, []string{"PaymentMethodData", "bacs_debit"}, []string{"PaymentMethodData", "bacs_debit", "account_number"}, []string{"PaymentMethodData", "bacs_debit", "sort_code"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "billing_details", "tax_id"}, []string{"PaymentMethodData", "boleto"}, []string{"PaymentMethodData", "boleto", "tax_id"}, []string{"PaymentMethodData", "eps"}, []string{"PaymentMethodData", "eps", "bank"}, []string{"PaymentMethodData", "fpx"}, []string{"PaymentMethodData", "fpx", "account_holder_type"}, []string{"PaymentMethodData", "fpx", "bank"}, []string{"PaymentMethodData", "ideal"}, []string{"PaymentMethodData", "ideal", "bank"}, []string{"PaymentMethodData", "klarna"}, []string{"PaymentMethodData", "klarna", "dob"}, []string{"PaymentMethodData", "klarna", "dob", "day"}, []string{"PaymentMethodData", "klarna", "dob", "month"}, []string{"PaymentMethodData", "klarna", "dob", "year"}, []string{"PaymentMethodData", "metadata"}, []string{"PaymentMethodData", "naver_pay"}, []string{"PaymentMethodData", "naver_pay", "funding"}, []string{"PaymentMethodData", "nz_bank_account"}, []string{"PaymentMethodData", "nz_bank_account", "account_holder_name"}, []string{"PaymentMethodData", "nz_bank_account", "account_number"}, []string{"PaymentMethodData", "nz_bank_account", "bank_code"}, []string{"PaymentMethodData", "nz_bank_account", "branch_code"}, []string{"PaymentMethodData", "nz_bank_account", "reference"}, []string{"PaymentMethodData", "nz_bank_account", "suffix"}, []string{"PaymentMethodData", "p24"}, []string{"PaymentMethodData", "p24", "bank"}, []string{"PaymentMethodData", "payto"}, []string{"PaymentMethodData", "payto", "account_number"}, []string{"PaymentMethodData", "payto", "bsb_number"}, []string{"PaymentMethodData", "payto", "pay_id"}, []string{"PaymentMethodData", "radar_options"}, []string{"PaymentMethodData", "radar_options", "session"}, []string{"PaymentMethodData", "sepa_debit"}, []string{"PaymentMethodData", "sepa_debit", "iban"}, []string{"PaymentMethodData", "sofort"}, []string{"PaymentMethodData", "sofort", "country"}, []string{"PaymentMethodData", "type"}, []string{"PaymentMethodData", "upi"}, []string{"PaymentMethodData", "upi", "mandate_options"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodData", "upi", "mandate_options", "description"}, []string{"PaymentMethodData", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodData", "us_bank_account"}, []string{"PaymentMethodData", "us_bank_account", "account_holder_type"}, []string{"PaymentMethodData", "us_bank_account", "account_number"}, []string{"PaymentMethodData", "us_bank_account", "account_type"}, []string{"PaymentMethodData", "us_bank_account", "financial_connections_account"}, []string{"PaymentMethodData", "us_bank_account", "routing_number"}, []string{"ReturnURL"}, []string{"SingleUse"}, []string{"SingleUse", "amount"}, []string{"SingleUse", "currency"}, []string{"UseStripeSDK"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *SetupIntentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SetupIntentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config SetupIntentResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"PaymentMethodConfiguration"}, []string{"PaymentMethodData"}, []string{"PaymentMethodData", "acss_debit"}, []string{"PaymentMethodData", "acss_debit", "account_number"}, []string{"PaymentMethodData", "acss_debit", "institution_number"}, []string{"PaymentMethodData", "acss_debit", "transit_number"}, []string{"PaymentMethodData", "allow_redisplay"}, []string{"PaymentMethodData", "au_becs_debit"}, []string{"PaymentMethodData", "au_becs_debit", "account_number"}, []string{"PaymentMethodData", "au_becs_debit", "bsb_number"}, []string{"PaymentMethodData", "bacs_debit"}, []string{"PaymentMethodData", "bacs_debit", "account_number"}, []string{"PaymentMethodData", "bacs_debit", "sort_code"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "billing_details", "tax_id"}, []string{"PaymentMethodData", "boleto"}, []string{"PaymentMethodData", "boleto", "tax_id"}, []string{"PaymentMethodData", "eps"}, []string{"PaymentMethodData", "eps", "bank"}, []string{"PaymentMethodData", "fpx"}, []string{"PaymentMethodData", "fpx", "account_holder_type"}, []string{"PaymentMethodData", "fpx", "bank"}, []string{"PaymentMethodData", "ideal"}, []string{"PaymentMethodData", "ideal", "bank"}, []string{"PaymentMethodData", "klarna"}, []string{"PaymentMethodData", "klarna", "dob"}, []string{"PaymentMethodData", "klarna", "dob", "day"}, []string{"PaymentMethodData", "klarna", "dob", "month"}, []string{"PaymentMethodData", "klarna", "dob", "year"}, []string{"PaymentMethodData", "metadata"}, []string{"PaymentMethodData", "naver_pay"}, []string{"PaymentMethodData", "naver_pay", "funding"}, []string{"PaymentMethodData", "nz_bank_account"}, []string{"PaymentMethodData", "nz_bank_account", "account_holder_name"}, []string{"PaymentMethodData", "nz_bank_account", "account_number"}, []string{"PaymentMethodData", "nz_bank_account", "bank_code"}, []string{"PaymentMethodData", "nz_bank_account", "branch_code"}, []string{"PaymentMethodData", "nz_bank_account", "reference"}, []string{"PaymentMethodData", "nz_bank_account", "suffix"}, []string{"PaymentMethodData", "p24"}, []string{"PaymentMethodData", "p24", "bank"}, []string{"PaymentMethodData", "payto"}, []string{"PaymentMethodData", "payto", "account_number"}, []string{"PaymentMethodData", "payto", "bsb_number"}, []string{"PaymentMethodData", "payto", "pay_id"}, []string{"PaymentMethodData", "radar_options"}, []string{"PaymentMethodData", "radar_options", "session"}, []string{"PaymentMethodData", "sepa_debit"}, []string{"PaymentMethodData", "sepa_debit", "iban"}, []string{"PaymentMethodData", "sofort"}, []string{"PaymentMethodData", "sofort", "country"}, []string{"PaymentMethodData", "type"}, []string{"PaymentMethodData", "upi"}, []string{"PaymentMethodData", "upi", "mandate_options"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodData", "upi", "mandate_options", "description"}, []string{"PaymentMethodData", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodData", "us_bank_account"}, []string{"PaymentMethodData", "us_bank_account", "account_holder_type"}, []string{"PaymentMethodData", "us_bank_account", "account_number"}, []string{"PaymentMethodData", "us_bank_account", "account_type"}, []string{"PaymentMethodData", "us_bank_account", "financial_connections_account"}, []string{"PaymentMethodData", "us_bank_account", "routing_number"}, []string{"ReturnURL"}, []string{"SingleUse"}, []string{"SingleUse", "amount"}, []string{"SingleUse", "currency"}, []string{"UseStripeSDK"}})

	var state SetupIntentResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"ReturnURL"}, []string{"SingleUse"}, []string{"SingleUse", "amount"}, []string{"SingleUse", "currency"}, []string{"UseStripeSDK"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"ReturnURL"}, []string{"SingleUse"}, []string{"SingleUse", "amount"}, []string{"SingleUse", "currency"}, []string{"UseStripeSDK"}})

	params, err := expandSetupIntentUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building SetupIntent update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building SetupIntent update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1SetupIntents.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating SetupIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1SetupIntents.B, r.client.V1SetupIntents.Key, stripe.FormatURLPath("/v1/setup_intents/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating SetupIntent update raw response", err.Error())
		return
	}

	if err := flattenSetupIntent(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening SetupIntent update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PaymentMethodOptions", "card", "moto"}, []string{"PaymentMethodOptions", "card", "three_d_secure"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "ares_trans_status"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "cryptogram"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "electronic_commerce_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_avalgo"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_exemption"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "network_options", "cartes_bancaires", "cb_score"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "requestor_challenge_indicator"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "transaction_id"}, []string{"PaymentMethodOptions", "card", "three_d_secure", "version"}, []string{"PaymentMethodOptions", "klarna", "on_demand"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "average_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "maximum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "minimum_amount"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval"}, []string{"PaymentMethodOptions", "klarna", "on_demand", "purchase_interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "interval_count"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "name"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "amount"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "next_billing", "date"}, []string{"PaymentMethodOptions", "klarna", "subscriptions", "*", "reference"}, []string{"PaymentMethodOptions", "upi", "setup_future_usage"}, []string{"PaymentMethodOptions", "us_bank_account", "networks"}, []string{"PaymentMethodOptions", "us_bank_account", "networks", "requested"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Confirm"}, []string{"ConfirmationToken"}, []string{"MandateData"}, []string{"MandateData", "customer_acceptance"}, []string{"MandateData", "customer_acceptance", "accepted_at"}, []string{"MandateData", "customer_acceptance", "online"}, []string{"MandateData", "customer_acceptance", "online", "ip_address"}, []string{"MandateData", "customer_acceptance", "online", "user_agent"}, []string{"MandateData", "customer_acceptance", "type"}, []string{"PaymentMethodConfiguration"}, []string{"PaymentMethodData"}, []string{"PaymentMethodData", "acss_debit"}, []string{"PaymentMethodData", "acss_debit", "account_number"}, []string{"PaymentMethodData", "acss_debit", "institution_number"}, []string{"PaymentMethodData", "acss_debit", "transit_number"}, []string{"PaymentMethodData", "allow_redisplay"}, []string{"PaymentMethodData", "au_becs_debit"}, []string{"PaymentMethodData", "au_becs_debit", "account_number"}, []string{"PaymentMethodData", "au_becs_debit", "bsb_number"}, []string{"PaymentMethodData", "bacs_debit"}, []string{"PaymentMethodData", "bacs_debit", "account_number"}, []string{"PaymentMethodData", "bacs_debit", "sort_code"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "billing_details", "tax_id"}, []string{"PaymentMethodData", "boleto"}, []string{"PaymentMethodData", "boleto", "tax_id"}, []string{"PaymentMethodData", "eps"}, []string{"PaymentMethodData", "eps", "bank"}, []string{"PaymentMethodData", "fpx"}, []string{"PaymentMethodData", "fpx", "account_holder_type"}, []string{"PaymentMethodData", "fpx", "bank"}, []string{"PaymentMethodData", "ideal"}, []string{"PaymentMethodData", "ideal", "bank"}, []string{"PaymentMethodData", "klarna"}, []string{"PaymentMethodData", "klarna", "dob"}, []string{"PaymentMethodData", "klarna", "dob", "day"}, []string{"PaymentMethodData", "klarna", "dob", "month"}, []string{"PaymentMethodData", "klarna", "dob", "year"}, []string{"PaymentMethodData", "metadata"}, []string{"PaymentMethodData", "naver_pay"}, []string{"PaymentMethodData", "naver_pay", "funding"}, []string{"PaymentMethodData", "nz_bank_account"}, []string{"PaymentMethodData", "nz_bank_account", "account_holder_name"}, []string{"PaymentMethodData", "nz_bank_account", "account_number"}, []string{"PaymentMethodData", "nz_bank_account", "bank_code"}, []string{"PaymentMethodData", "nz_bank_account", "branch_code"}, []string{"PaymentMethodData", "nz_bank_account", "reference"}, []string{"PaymentMethodData", "nz_bank_account", "suffix"}, []string{"PaymentMethodData", "p24"}, []string{"PaymentMethodData", "p24", "bank"}, []string{"PaymentMethodData", "payto"}, []string{"PaymentMethodData", "payto", "account_number"}, []string{"PaymentMethodData", "payto", "bsb_number"}, []string{"PaymentMethodData", "payto", "pay_id"}, []string{"PaymentMethodData", "radar_options"}, []string{"PaymentMethodData", "radar_options", "session"}, []string{"PaymentMethodData", "sepa_debit"}, []string{"PaymentMethodData", "sepa_debit", "iban"}, []string{"PaymentMethodData", "sofort"}, []string{"PaymentMethodData", "sofort", "country"}, []string{"PaymentMethodData", "type"}, []string{"PaymentMethodData", "upi"}, []string{"PaymentMethodData", "upi", "mandate_options"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount"}, []string{"PaymentMethodData", "upi", "mandate_options", "amount_type"}, []string{"PaymentMethodData", "upi", "mandate_options", "description"}, []string{"PaymentMethodData", "upi", "mandate_options", "end_date"}, []string{"PaymentMethodData", "us_bank_account"}, []string{"PaymentMethodData", "us_bank_account", "account_holder_type"}, []string{"PaymentMethodData", "us_bank_account", "account_number"}, []string{"PaymentMethodData", "us_bank_account", "account_type"}, []string{"PaymentMethodData", "us_bank_account", "financial_connections_account"}, []string{"PaymentMethodData", "us_bank_account", "routing_number"}, []string{"ReturnURL"}, []string{"SingleUse"}, []string{"SingleUse", "amount"}, []string{"SingleUse", "currency"}, []string{"UseStripeSDK"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SetupIntentResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *SetupIntentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandSetupIntentCreate(plan SetupIntentResourceModel) (*stripe.SetupIntentCreateParams, error) {
	params := &stripe.SetupIntentCreateParams{}

	if !plan.AttachToSelf.IsNull() && !plan.AttachToSelf.IsUnknown() {
		params.AttachToSelf = stripe.Bool(plan.AttachToSelf.ValueBool())
	}
	if !plan.AutomaticPaymentMethods.IsNull() && !plan.AutomaticPaymentMethods.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AutomaticPaymentMethods", plan.AutomaticPaymentMethods) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "automatic_payment_methods", params)
		}
	}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerID", "Customer", plan.Customer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer", params)
		}
	}
	if !plan.CustomerAccount.IsNull() && !plan.CustomerAccount.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerAccount", "CustomerAccount", plan.CustomerAccount.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer_account", params)
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
	if !plan.FlowDirections.IsNull() && !plan.FlowDirections.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FlowDirections", plan.FlowDirections) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "flow_directions", params)
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
	if !plan.PaymentMethodTypes.IsNull() && !plan.PaymentMethodTypes.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodTypes", plan.PaymentMethodTypes) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method_types", params)
		}
	}
	if !plan.Usage.IsNull() && !plan.Usage.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Usage", "Usage", plan.Usage.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "usage", params)
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
	if !plan.ReturnURL.IsNull() && !plan.ReturnURL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReturnURL", "ReturnURL", plan.ReturnURL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "return_url", params)
		}
	}
	if !plan.SingleUse.IsNull() && !plan.SingleUse.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SingleUse", plan.SingleUse) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "single_use", params)
		}
	}
	if !plan.UseStripeSDK.IsNull() && !plan.UseStripeSDK.IsUnknown() {
		params.UseStripeSDK = stripe.Bool(plan.UseStripeSDK.ValueBool())
	}

	return params, nil
}

func expandSetupIntentUpdate(plan SetupIntentResourceModel, state SetupIntentResourceModel) (*stripe.SetupIntentUpdateParams, error) {
	params := &stripe.SetupIntentUpdateParams{}

	if !plan.AttachToSelf.Equal(state.AttachToSelf) && !plan.AttachToSelf.IsNull() && !plan.AttachToSelf.IsUnknown() {
		params.AttachToSelf = stripe.Bool(plan.AttachToSelf.ValueBool())
	}
	if !plan.Customer.Equal(state.Customer) && !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerID", "Customer", plan.Customer.ValueString()) {
			if !plan.Customer.Equal(state.Customer) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "customer", params)
			}
		}
	}
	if !plan.CustomerAccount.Equal(state.CustomerAccount) && !plan.CustomerAccount.IsNull() && !plan.CustomerAccount.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerAccount", "CustomerAccount", plan.CustomerAccount.ValueString()) {
			if !plan.CustomerAccount.Equal(state.CustomerAccount) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "customer_account", params)
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
	if !plan.FlowDirections.Equal(state.FlowDirections) && !plan.FlowDirections.IsNull() && !plan.FlowDirections.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FlowDirections", plan.FlowDirections) {
			if !plan.FlowDirections.Equal(state.FlowDirections) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "flow_directions", params)
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
	if !plan.PaymentMethodTypes.Equal(state.PaymentMethodTypes) && !plan.PaymentMethodTypes.IsNull() && !plan.PaymentMethodTypes.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodTypes", plan.PaymentMethodTypes) {
			if !plan.PaymentMethodTypes.Equal(state.PaymentMethodTypes) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_method_types", params)
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

func flattenSetupIntent(obj *stripe.SetupIntent, state *SetupIntentResourceModel) error {
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
		if rawValueAttachToSelf, rawOk := plainValueAtPath(raw, "attach_to_self"); rawOk {
			if valueAttachToSelf, err := flattenPlainValue(rawValueAttachToSelf, types.BoolType, "attach_to_self", "raw response"); err != nil {
				return err
			} else {
				if typedAttachToSelf, ok := valueAttachToSelf.(types.Bool); ok {
					state.AttachToSelf = typedAttachToSelf
				}
			}
		} else if !hasRaw {
			if responseValueAttachToSelf, ok := plainFromResponseField(obj, "AttachToSelf"); ok {
				if valueAttachToSelf, err := flattenPlainValue(responseValueAttachToSelf, types.BoolType, "attach_to_self", "response struct"); err != nil {
					return err
				} else {
					if typedAttachToSelf, ok := valueAttachToSelf.(types.Bool); ok {
						state.AttachToSelf = typedAttachToSelf
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
		if rawValueCustomerAccount, rawOk := plainValueAtPath(raw, "customer_account"); rawOk {
			if valueCustomerAccount, err := flattenPlainValue(rawValueCustomerAccount, types.StringType, "customer_account", "raw response"); err != nil {
				return err
			} else {
				if typedCustomerAccount, ok := valueCustomerAccount.(types.String); ok {
					state.CustomerAccount = typedCustomerAccount
				}
			}
		} else if !hasRaw {
			if responseValueCustomerAccount, ok := plainFromResponseField(obj, "CustomerAccount"); ok {
				if valueCustomerAccount, err := flattenPlainValue(responseValueCustomerAccount, types.StringType, "customer_account", "response struct"); err != nil {
					return err
				} else {
					if typedCustomerAccount, ok := valueCustomerAccount.(types.String); ok {
						state.CustomerAccount = typedCustomerAccount
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
		if rawValueFlowDirections, rawOk := plainValueAtPath(raw, "flow_directions"); rawOk {
			if valueFlowDirections, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueFlowDirections, attrValueToPlain(state.FlowDirections)), types.ListType{ElemType: types.StringType}, "flow_directions", "raw response"); err != nil {
				return err
			} else {
				if typedFlowDirections, ok := valueFlowDirections.(types.List); ok {
					state.FlowDirections = typedFlowDirections
				}
			}
		} else if !hasRaw {
			if responseValueFlowDirections, ok := plainFromResponseField(obj, "FlowDirections"); ok {
				if valueFlowDirections, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueFlowDirections, attrValueToPlain(state.FlowDirections)),
					types.ListType{ElemType: types.StringType},
					"flow_directions",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedFlowDirections, ok := valueFlowDirections.(types.List); ok {
						state.FlowDirections = typedFlowDirections
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
		assignedLastSetupError := false
		hadRawLastSetupError := false
		if rawValueLastSetupError, rawOk := plainValueAtPath(raw, "last_setup_error"); rawOk {
			hadRawLastSetupError = true
			if rawValueLastSetupError != nil {
				sourceLastSetupError := applyConfiguredKeyedListShapes(rawValueLastSetupError, attrValueToPlain(state.LastSetupError))
				if valueLastSetupError, err := flattenPlainValue(sourceLastSetupError, types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}}, "last_setup_error", "raw response"); err != nil {
					return err
				} else {
					if typedLastSetupError, ok := valueLastSetupError.(types.Object); ok {
						state.LastSetupError = typedLastSetupError
						assignedLastSetupError = true
					}
				}
			}
		}
		if !assignedLastSetupError {
			if !hasRaw {
				if responseValueLastSetupError, ok := plainFromResponseField(obj, "LastSetupError"); ok {
					sourceLastSetupError := applyConfiguredKeyedListShapes(responseValueLastSetupError, attrValueToPlain(state.LastSetupError))
					if valueLastSetupError, err := flattenPlainValue(
						sourceLastSetupError,
						types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}},
						"last_setup_error",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLastSetupError, ok := valueLastSetupError.(types.Object); ok {
							state.LastSetupError = typedLastSetupError
							assignedLastSetupError = true
						}
					}
				}
			}
		}
		if !assignedLastSetupError && hadRawLastSetupError {
			if nullLastSetupError, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}}); ok {
				if typedLastSetupError, ok := nullLastSetupError.(types.Object); ok {
					state.LastSetupError = typedLastSetupError
				}
			}
		}
	}
	{
		if true {
			if rawValueLatestAttempt, rawOk := plainValueAtPath(raw, "latest_attempt"); rawOk {
				if typedLatestAttempt, ok := plainToStringIDValue(rawValueLatestAttempt); ok {
					state.LatestAttempt = typedLatestAttempt
				}
			} else if !hasRaw {
				if responseValueLatestAttempt, ok := plainFromResponseField(obj, "LatestAttempt"); ok {
					if typedLatestAttempt, ok := plainToStringIDValue(responseValueLatestAttempt); ok {
						state.LatestAttempt = typedLatestAttempt
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
		if true {
			if rawValueMandate, rawOk := plainValueAtPath(raw, "mandate"); rawOk {
				if typedMandate, ok := plainToStringIDValue(rawValueMandate); ok {
					state.Mandate = typedMandate
				}
			} else if !hasRaw {
				if responseValueMandate, ok := plainFromResponseField(obj, "Mandate"); ok {
					if typedMandate, ok := plainToStringIDValue(responseValueMandate); ok {
						state.Mandate = typedMandate
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
				if valueNextAction, err := flattenPlainValue(sourceNextAction, types.ObjectType{AttrTypes: map[string]attr.Type{"cashapp_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "mobile_auth_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "pix_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "expires_at": types.Int64Type, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "redirect_to_url": types.ObjectType{AttrTypes: map[string]attr.Type{"return_url": types.StringType, "url": types.StringType}}, "type": types.StringType, "upi_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "use_stripe_sdk": types.MapType{ElemType: types.StringType}, "verify_with_microdeposits": types.ObjectType{AttrTypes: map[string]attr.Type{"arrival_date": types.Int64Type, "hosted_verification_url": types.StringType, "microdeposit_type": types.StringType}}}}, "next_action", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"cashapp_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "mobile_auth_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "pix_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "expires_at": types.Int64Type, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "redirect_to_url": types.ObjectType{AttrTypes: map[string]attr.Type{"return_url": types.StringType, "url": types.StringType}}, "type": types.StringType, "upi_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "use_stripe_sdk": types.MapType{ElemType: types.StringType}, "verify_with_microdeposits": types.ObjectType{AttrTypes: map[string]attr.Type{"arrival_date": types.Int64Type, "hosted_verification_url": types.StringType, "microdeposit_type": types.StringType}}}},
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
			if nullNextAction, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"cashapp_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "mobile_auth_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "pix_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"data": types.StringType, "expires_at": types.Int64Type, "hosted_instructions_url": types.StringType, "image_url_png": types.StringType, "image_url_svg": types.StringType}}, "redirect_to_url": types.ObjectType{AttrTypes: map[string]attr.Type{"return_url": types.StringType, "url": types.StringType}}, "type": types.StringType, "upi_handle_redirect_or_display_qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_instructions_url": types.StringType, "qr_code": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_at": types.Int64Type, "image_url_png": types.StringType, "image_url_svg": types.StringType}}}}, "use_stripe_sdk": types.MapType{ElemType: types.StringType}, "verify_with_microdeposits": types.ObjectType{AttrTypes: map[string]attr.Type{"arrival_date": types.Int64Type, "hosted_verification_url": types.StringType, "microdeposit_type": types.StringType}}}}); ok {
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
				if valuePaymentMethodOptions, err := flattenPlainValue(sourcePaymentMethodOptions, types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_mandate_url": types.StringType, "default_for": types.ListType{ElemType: types.StringType}, "interval_description": types.StringType, "payment_schedule": types.StringType, "transaction_type": types.StringType}}, "verification_method": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "currency": types.StringType, "description": types.StringType, "end_date": types.Int64Type, "interval": types.StringType, "interval_count": types.Int64Type, "reference": types.StringType, "start_date": types.Int64Type, "supported_types": types.ListType{ElemType: types.StringType}}}, "network": types.StringType, "request_three_d_secure": types.StringType, "moto": types.BoolType, "three_d_secure": types.ObjectType{AttrTypes: map[string]attr.Type{"ares_trans_status": types.StringType, "cryptogram": types.StringType, "electronic_commerce_indicator": types.StringType, "network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"cartes_bancaires": types.ObjectType{AttrTypes: map[string]attr.Type{"cb_avalgo": types.StringType, "cb_exemption": types.StringType, "cb_score": types.Int64Type}}}}, "requestor_challenge_indicator": types.StringType, "transaction_id": types.StringType, "version": types.StringType}}}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "preferred_locale": types.StringType, "on_demand": types.ObjectType{AttrTypes: map[string]attr.Type{"average_amount": types.Int64Type, "maximum_amount": types.Int64Type, "minimum_amount": types.Int64Type, "purchase_interval": types.StringType, "purchase_interval_count": types.Int64Type}}, "subscriptions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "name": types.StringType, "next_billing": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "date": types.StringType}}, "reference": types.StringType}}}}}, "link": types.ObjectType{AttrTypes: map[string]attr.Type{"persistent_token": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_agreement_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "payments_per_period": types.Int64Type, "purpose": types.StringType, "start_date": types.StringType}}}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "amount_type": types.StringType, "currency": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "reference": types.StringType, "start_date": types.StringType}}}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}, "setup_future_usage": types.StringType}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}, "return_url": types.StringType}}, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"collection_method": types.StringType}}, "verification_method": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.ListType{ElemType: types.StringType}}}}}}}, "payment_method_options", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_mandate_url": types.StringType, "default_for": types.ListType{ElemType: types.StringType}, "interval_description": types.StringType, "payment_schedule": types.StringType, "transaction_type": types.StringType}}, "verification_method": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "currency": types.StringType, "description": types.StringType, "end_date": types.Int64Type, "interval": types.StringType, "interval_count": types.Int64Type, "reference": types.StringType, "start_date": types.Int64Type, "supported_types": types.ListType{ElemType: types.StringType}}}, "network": types.StringType, "request_three_d_secure": types.StringType, "moto": types.BoolType, "three_d_secure": types.ObjectType{AttrTypes: map[string]attr.Type{"ares_trans_status": types.StringType, "cryptogram": types.StringType, "electronic_commerce_indicator": types.StringType, "network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"cartes_bancaires": types.ObjectType{AttrTypes: map[string]attr.Type{"cb_avalgo": types.StringType, "cb_exemption": types.StringType, "cb_score": types.Int64Type}}}}, "requestor_challenge_indicator": types.StringType, "transaction_id": types.StringType, "version": types.StringType}}}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "preferred_locale": types.StringType, "on_demand": types.ObjectType{AttrTypes: map[string]attr.Type{"average_amount": types.Int64Type, "maximum_amount": types.Int64Type, "minimum_amount": types.Int64Type, "purchase_interval": types.StringType, "purchase_interval_count": types.Int64Type}}, "subscriptions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "name": types.StringType, "next_billing": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "date": types.StringType}}, "reference": types.StringType}}}}}, "link": types.ObjectType{AttrTypes: map[string]attr.Type{"persistent_token": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_agreement_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "payments_per_period": types.Int64Type, "purpose": types.StringType, "start_date": types.StringType}}}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "amount_type": types.StringType, "currency": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "reference": types.StringType, "start_date": types.StringType}}}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}, "setup_future_usage": types.StringType}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}, "return_url": types.StringType}}, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"collection_method": types.StringType}}, "verification_method": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.ListType{ElemType: types.StringType}}}}}}},
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
			if nullPaymentMethodOptions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_mandate_url": types.StringType, "default_for": types.ListType{ElemType: types.StringType}, "interval_description": types.StringType, "payment_schedule": types.StringType, "transaction_type": types.StringType}}, "verification_method": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "currency": types.StringType, "description": types.StringType, "end_date": types.Int64Type, "interval": types.StringType, "interval_count": types.Int64Type, "reference": types.StringType, "start_date": types.Int64Type, "supported_types": types.ListType{ElemType: types.StringType}}}, "network": types.StringType, "request_three_d_secure": types.StringType, "moto": types.BoolType, "three_d_secure": types.ObjectType{AttrTypes: map[string]attr.Type{"ares_trans_status": types.StringType, "cryptogram": types.StringType, "electronic_commerce_indicator": types.StringType, "network_options": types.ObjectType{AttrTypes: map[string]attr.Type{"cartes_bancaires": types.ObjectType{AttrTypes: map[string]attr.Type{"cb_avalgo": types.StringType, "cb_exemption": types.StringType, "cb_score": types.Int64Type}}}}, "requestor_challenge_indicator": types.StringType, "transaction_id": types.StringType, "version": types.StringType}}}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "preferred_locale": types.StringType, "on_demand": types.ObjectType{AttrTypes: map[string]attr.Type{"average_amount": types.Int64Type, "maximum_amount": types.Int64Type, "minimum_amount": types.Int64Type, "purchase_interval": types.StringType, "purchase_interval_count": types.Int64Type}}, "subscriptions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "name": types.StringType, "next_billing": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "date": types.StringType}}, "reference": types.StringType}}}}}, "link": types.ObjectType{AttrTypes: map[string]attr.Type{"persistent_token": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_agreement_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "payments_per_period": types.Int64Type, "purpose": types.StringType, "start_date": types.StringType}}}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "amount_type": types.StringType, "currency": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "reference": types.StringType, "start_date": types.StringType}}}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}, "setup_future_usage": types.StringType}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}, "return_url": types.StringType}}, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"collection_method": types.StringType}}, "verification_method": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.ListType{ElemType: types.StringType}}}}}}}); ok {
				if typedPaymentMethodOptions, ok := nullPaymentMethodOptions.(types.Object); ok {
					state.PaymentMethodOptions = typedPaymentMethodOptions
				}
			}
		}
	}
	{
		if rawValuePaymentMethodTypes, rawOk := plainValueAtPath(raw, "payment_method_types"); rawOk {
			if valuePaymentMethodTypes, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValuePaymentMethodTypes, attrValueToPlain(state.PaymentMethodTypes)), types.ListType{ElemType: types.StringType}, "payment_method_types", "raw response"); err != nil {
				return err
			} else {
				if typedPaymentMethodTypes, ok := valuePaymentMethodTypes.(types.List); ok {
					state.PaymentMethodTypes = typedPaymentMethodTypes
				}
			}
		} else if !hasRaw {
			if responseValuePaymentMethodTypes, ok := plainFromResponseField(obj, "PaymentMethodTypes"); ok {
				if valuePaymentMethodTypes, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValuePaymentMethodTypes, attrValueToPlain(state.PaymentMethodTypes)),
					types.ListType{ElemType: types.StringType},
					"payment_method_types",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedPaymentMethodTypes, ok := valuePaymentMethodTypes.(types.List); ok {
						state.PaymentMethodTypes = typedPaymentMethodTypes
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueSingleUseMandate, rawOk := plainValueAtPath(raw, "single_use_mandate"); rawOk {
				if typedSingleUseMandate, ok := plainToStringIDValue(rawValueSingleUseMandate); ok {
					state.SingleUseMandate = typedSingleUseMandate
				}
			} else if !hasRaw {
				if responseValueSingleUseMandate, ok := plainFromResponseField(obj, "SingleUseMandate"); ok {
					if typedSingleUseMandate, ok := plainToStringIDValue(responseValueSingleUseMandate); ok {
						state.SingleUseMandate = typedSingleUseMandate
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
		if rawValueUsage, rawOk := plainValueAtPath(raw, "usage"); rawOk {
			if valueUsage, err := flattenPlainValue(rawValueUsage, types.StringType, "usage", "raw response"); err != nil {
				return err
			} else {
				if typedUsage, ok := valueUsage.(types.String); ok {
					state.Usage = typedUsage
				}
			}
		} else if !hasRaw {
			if responseValueUsage, ok := plainFromResponseField(obj, "Usage"); ok {
				if valueUsage, err := flattenPlainValue(responseValueUsage, types.StringType, "usage", "response struct"); err != nil {
					return err
				} else {
					if typedUsage, ok := valueUsage.(types.String); ok {
						state.Usage = typedUsage
					}
				}
			}
		}
	}
	return nil
}
