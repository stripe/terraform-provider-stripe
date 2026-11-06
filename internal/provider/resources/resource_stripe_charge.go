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

var _ resource.Resource = &ChargeResource{}

var _ resource.ResourceWithConfigure = &ChargeResource{}

var _ resource.ResourceWithImportState = &ChargeResource{}

func NewChargeResource() resource.Resource {
	return &ChargeResource{}
}

type ChargeResource struct {
	client *stripe.Client
}

type ChargeResourceModel struct {
	Object                        types.String `tfsdk:"object"`
	Amount                        types.Int64  `tfsdk:"amount"`
	AmountCaptured                types.Int64  `tfsdk:"amount_captured"`
	AmountRefunded                types.Int64  `tfsdk:"amount_refunded"`
	Application                   types.String `tfsdk:"application"`
	ApplicationFee                types.Int64  `tfsdk:"application_fee"`
	ApplicationFeeAmount          types.Int64  `tfsdk:"application_fee_amount"`
	AuthorizationCode             types.String `tfsdk:"authorization_code"`
	BalanceTransaction            types.String `tfsdk:"balance_transaction"`
	BillingDetails                types.Object `tfsdk:"billing_details"`
	CalculatedStatementDescriptor types.String `tfsdk:"calculated_statement_descriptor"`
	Captured                      types.Bool   `tfsdk:"captured"`
	Created                       types.Int64  `tfsdk:"created"`
	Currency                      types.String `tfsdk:"currency"`
	Customer                      types.String `tfsdk:"customer"`
	Description                   types.String `tfsdk:"description"`
	Disputed                      types.Bool   `tfsdk:"disputed"`
	FailureBalanceTransaction     types.String `tfsdk:"failure_balance_transaction"`
	FailureCode                   types.String `tfsdk:"failure_code"`
	FailureMessage                types.String `tfsdk:"failure_message"`
	FraudDetails                  types.Object `tfsdk:"fraud_details"`
	ID                            types.String `tfsdk:"id"`
	Level3                        types.Object `tfsdk:"level3"`
	Livemode                      types.Bool   `tfsdk:"livemode"`
	Metadata                      types.Map    `tfsdk:"metadata"`
	OnBehalfOf                    types.String `tfsdk:"on_behalf_of"`
	Outcome                       types.Object `tfsdk:"outcome"`
	Paid                          types.Bool   `tfsdk:"paid"`
	PaymentIntent                 types.String `tfsdk:"payment_intent"`
	PaymentMethod                 types.String `tfsdk:"payment_method"`
	PaymentMethodDetails          types.Object `tfsdk:"payment_method_details"`
	PresentmentDetails            types.Object `tfsdk:"presentment_details"`
	RadarOptions                  types.Object `tfsdk:"radar_options"`
	ReceiptEmail                  types.String `tfsdk:"receipt_email"`
	ReceiptNumber                 types.String `tfsdk:"receipt_number"`
	Refunded                      types.Bool   `tfsdk:"refunded"`
	Review                        types.String `tfsdk:"review"`
	Shipping                      types.Object `tfsdk:"shipping"`
	Source                        types.String `tfsdk:"source"`
	SourceTransfer                types.String `tfsdk:"source_transfer"`
	StatementDescriptor           types.String `tfsdk:"statement_descriptor"`
	StatementDescriptorSuffix     types.String `tfsdk:"statement_descriptor_suffix"`
	Status                        types.String `tfsdk:"status"`
	Transfer                      types.String `tfsdk:"transfer"`
	TransferData                  types.Object `tfsdk:"transfer_data"`
	TransferGroup                 types.String `tfsdk:"transfer_group"`
	Capture                       types.Bool   `tfsdk:"capture"`
	Destination                   types.Object `tfsdk:"destination"`
}

func (r *ChargeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ChargeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_charge"
}

func (r *ChargeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The `Charge` object represents a single attempt to move money into your Stripe account.\nPaymentIntent confirmation is the most common way to create Charges, but [Account Debits](https://docs.stripe.com/connect/account-debits) may also create Charges.\nSome legacy payment flows create Charges directly, which is not recommended for new integrations.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("charge")},
			},
			"amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Amount intended to be collected by this payment. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://docs.stripe.com/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"amount_captured": schema.Int64Attribute{
				Computed:      true,
				Description:   "Amount in cents (or local equivalent) captured (can be less than the amount attribute on the charge if a partial capture was made).",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"amount_refunded": schema.Int64Attribute{
				Computed:      true,
				Description:   "Amount in cents (or local equivalent) refunded (can be less than the amount attribute on the charge if a partial refund was issued).",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the Connect application that created the charge.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"application_fee": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The application fee (if any) for the charge. [See the Connect documentation](https://docs.stripe.com/connect/direct-charges#collect-fees) for details.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"application_fee_amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The amount of the application fee (if any) requested for the charge. [See the Connect documentation](https://docs.stripe.com/connect/direct-charges#collect-fees) for details.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"authorization_code": schema.StringAttribute{
				Computed:      true,
				Description:   "Authorization code on the charge.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"balance_transaction": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the balance transaction that describes the impact of this charge on your account balance (not including refunds or disputes).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"billing_details": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"address": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Billing address.",
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
					"email": schema.StringAttribute{
						Computed:      true,
						Description:   "Email address.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"name": schema.StringAttribute{
						Computed:      true,
						Description:   "Full name.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"phone": schema.StringAttribute{
						Computed:      true,
						Description:   "Billing phone number (including extension).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"tax_id": schema.StringAttribute{
						Computed:      true,
						Description:   "Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"calculated_statement_descriptor": schema.StringAttribute{
				Computed:      true,
				Description:   "The full statement descriptor that is passed to card networks, and that is displayed on your customers' credit card and bank statements. Allows you to see what the statement descriptor looks like after the static and dynamic portions are combined. This value only exists for card payments.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"captured": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the charge was created without capturing, this Boolean represents whether it is still uncaptured or has since been captured.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the customer this charge is for if one exists.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the object. Often useful for displaying to users.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"disputed": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether the charge has been disputed.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"failure_balance_transaction": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the balance transaction that describes the reversal of the balance on your account due to payment failure.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"failure_code": schema.StringAttribute{
				Computed:      true,
				Description:   "Error code explaining reason for charge failure if available (see [the errors section](https://docs.stripe.com/error-codes) for a list of codes).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"failure_message": schema.StringAttribute{
				Computed:      true,
				Description:   "Message to user further explaining reason for charge failure if available.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"fraud_details": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Information on fraud assessments for the charge.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"stripe_report": schema.StringAttribute{
						Computed:      true,
						Description:   "Assessments from Stripe. If set, the value is `fraudulent`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"user_report": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Assessments reported by you. If set, possible values of are `safe` and `fraudulent`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"level3": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"customer_reference": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"line_items": schema.ListNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"discount_amount": schema.Int64Attribute{
									Computed: true,

									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"product_code": schema.StringAttribute{
									Computed: true,

									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"product_description": schema.StringAttribute{
									Computed: true,

									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"quantity": schema.Int64Attribute{
									Computed: true,

									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"tax_amount": schema.Int64Attribute{
									Computed: true,

									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"unit_cost": schema.Int64Attribute{
									Computed: true,

									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
							},
						},
					},
					"merchant_reference": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"shipping_address_zip": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"shipping_amount": schema.Int64Attribute{
						Computed: true,

						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"shipping_from_zip": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"on_behalf_of": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account (if any) the charge was made on behalf of without triggering an automatic transfer. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"outcome": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Details about whether the payment was accepted, and why. See [understanding declines](https://docs.stripe.com/declines) for details.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"advice_code": schema.StringAttribute{
						Computed:      true,
						Description:   "An enumerated value providing a more detailed explanation on [how to proceed with an error](https://docs.stripe.com/declines#retrying-issuer-declines).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("confirm_card_data", "do_not_try_again", "try_again_later")},
					},
					"network_advice_code": schema.StringAttribute{
						Computed:      true,
						Description:   "For charges declined by the network, a 2 digit code which indicates the advice returned by the network on how to proceed with an error.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"network_decline_code": schema.StringAttribute{
						Computed:      true,
						Description:   "For charges declined by the network, an alphanumeric code which indicates the reason the charge failed.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"network_status": schema.StringAttribute{
						Computed:      true,
						Description:   "Possible values are `approved_by_network`, `declined_by_network`, `not_sent_to_network`, and `reversed_after_approval`. The value `reversed_after_approval` indicates the payment was [blocked by Stripe](https://docs.stripe.com/declines#blocked-payments) after bank authorization, and may temporarily appear as \"pending\" on a cardholder's statement.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"reason": schema.StringAttribute{
						Computed:      true,
						Description:   "An enumerated value providing a more detailed explanation of the outcome's `type`. Charges blocked by Radar's default block rule have the value `highest_risk_level`. Charges placed in review by Radar's default review rule have the value `elevated_risk_level`. Charges blocked because the payment is unlikely to be authorized have the value `low_probability_of_authorization`. Charges authorized, blocked, or placed in review by custom rules have the value `rule`. See [understanding declines](https://docs.stripe.com/declines) for more details.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"risk_level": schema.StringAttribute{
						Computed:      true,
						Description:   "Stripe Radar's evaluation of the riskiness of the payment. Possible values for evaluated payments are `normal`, `elevated`, `highest`. For non-card payments, and card-based payments predating the public assignment of risk levels, this field will have the value `not_assessed`. In the event of an error in the evaluation, this field will have the value `unknown`. This field is only available with Radar.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"risk_score": schema.Int64Attribute{
						Computed:      true,
						Description:   "Stripe Radar's evaluation of the riskiness of the payment. Possible values for evaluated payments are between 0 and 100. For non-card payments, card-based payments predating the public assignment of risk scores, or in the event of an error during evaluation, this field will not be present. This field is only available with Radar for Fraud Teams.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"rule": schema.StringAttribute{
						Computed:      true,
						Description:   "The ID of the Radar rule that matched the payment, if applicable.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"seller_message": schema.StringAttribute{
						Computed:      true,
						Description:   "A human-readable description of the outcome type and reason, designed for you (the recipient of the payment), not your customer.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "Possible values are `authorized`, `manual_review`, `issuer_declined`, `blocked`, and `invalid`. See [understanding declines](https://docs.stripe.com/declines) and [Radar reviews](https://docs.stripe.com/radar/reviews) for details.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"paid": schema.BoolAttribute{
				Computed:      true,
				Description:   "`true` if the charge succeeded, or was successfully authorized for later capture.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"payment_intent": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the PaymentIntent associated with this charge, if one exists.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"payment_method": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the payment method used in this charge.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"payment_method_details": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Details about the payment method at the time of the transaction.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"ach_credit_transfer": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account_number": schema.StringAttribute{
								Computed:      true,
								Description:   "Account number to transfer funds to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Name of the bank associated with the routing number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"routing_number": schema.StringAttribute{
								Computed:      true,
								Description:   "Routing transit number for the bank account to transfer funds to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"swift_code": schema.StringAttribute{
								Computed:      true,
								Description:   "SWIFT code of the bank associated with the routing number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"ach_debit": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account_holder_type": schema.StringAttribute{
								Computed:      true,
								Description:   "Type of entity that holds the account. This can be either `individual` or `company`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("company", "individual")},
							},
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Name of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"country": schema.StringAttribute{
								Computed:      true,
								Description:   "Two-letter ISO code representing the country the bank account is located in.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four digits of the bank account number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"routing_number": schema.StringAttribute{
								Computed:      true,
								Description:   "Routing transit number of the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"acss_debit": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Name of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expected_debit_date": schema.StringAttribute{
								Computed:      true,
								Description:   "Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"institution_number": schema.StringAttribute{
								Computed:      true,
								Description:   "Institution number of the bank account",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four digits of the bank account number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the mandate used to make this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transit_number": schema.StringAttribute{
								Computed:      true,
								Description:   "Transit number of the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"affirm": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"location": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the location that this reader is assigned to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reader": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the reader this transaction was made on.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Affirm transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"afterpay_clearpay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"order_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Afterpay order ID associated with this payment intent.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reference": schema.StringAttribute{
								Computed:      true,
								Description:   "Order identifier shown to the merchant in Afterpay's online portal.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"alipay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"buyer_id": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "Transaction ID of this particular Alipay transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"alma": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"installments": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"count": schema.Int64Attribute{
										Computed:      true,
										Description:   "The number of installments.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
								},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Alma transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"amazon_pay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"funding": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"card": schema.SingleNestedAttribute{
										Computed: true,

										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"brand": schema.StringAttribute{
												Computed:      true,
												Description:   "Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"country": schema.StringAttribute{
												Computed:      true,
												Description:   "Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"exp_month": schema.Int64Attribute{
												Computed:      true,
												Description:   "Two-digit number representing the card's expiration month.",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
											},
											"exp_year": schema.Int64Attribute{
												Computed:      true,
												Description:   "Four-digit number representing the card's expiration year.",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
											},
											"funding": schema.StringAttribute{
												Computed:      true,
												Description:   "Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"last4": schema.StringAttribute{
												Computed:      true,
												Description:   "The last four digits of the card.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
										},
									},
									"type": schema.StringAttribute{
										Computed:      true,
										Description:   "funding type of the underlying payment method.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("card")},
									},
								},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Amazon Pay transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"au_becs_debit": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bsb_number": schema.StringAttribute{
								Computed:      true,
								Description:   "Bank-State-Branch number of the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expected_debit_date": schema.StringAttribute{
								Computed:      true,
								Description:   "Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four digits of the bank account number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the mandate used to make this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"bacs_debit": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"expected_debit_date": schema.StringAttribute{
								Computed:      true,
								Description:   "Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four digits of the bank account number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the mandate used to make this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"sort_code": schema.StringAttribute{
								Computed:      true,
								Description:   "Sort code of the bank account. (e.g., `10-20-30`)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"bancontact": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank_code": schema.StringAttribute{
								Computed:      true,
								Description:   "Bank code of bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Name of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bic": schema.StringAttribute{
								Computed:      true,
								Description:   "Bank Identifier Code of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"generated_sepa_debit": schema.StringAttribute{
								Computed:      true,
								Description:   "The ID of the SEPA Direct Debit PaymentMethod which was generated by this Charge.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"generated_sepa_debit_mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "The mandate for the SEPA Direct Debit PaymentMethod which was generated by this Charge.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"iban_last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four characters of the IBAN.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"preferred_language": schema.StringAttribute{
								Computed:      true,
								Description:   "Preferred language of the Bancontact authorization page that the customer is redirected to.\nCan be one of `en`, `de`, `fr`, or `nl`",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("de", "en", "fr", "nl")},
							},
							"verified_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Owner's verified full name. Values are verified or provided by Bancontact directly\n(if supported) at the time of authorization or settlement. They cannot be set or mutated.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"billie": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Billie transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"bizum": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Bizum transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"blik": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"buyer_id": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique and immutable identifier assigned by BLIK to every buyer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"boleto": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"tax_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The tax ID of the customer (CPF for individuals consumers or CNPJ for businesses consumers)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"card": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"amount_authorized": schema.Int64Attribute{
								Computed:      true,
								Description:   "The authorized amount.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"authorization_code": schema.StringAttribute{
								Computed:      true,
								Description:   "Authorization code on the charge.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"brand": schema.StringAttribute{
								Computed:      true,
								Description:   "Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"capture_before": schema.Int64Attribute{
								Computed:      true,
								Description:   "When using manual capture, a future timestamp at which the charge will be automatically refunded if uncaptured.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"checks": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Check results by Card networks on Card address and CVC at time of payment.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"address_line1_check": schema.StringAttribute{
										Computed:      true,
										Description:   "If a address line1 was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"address_postal_code_check": schema.StringAttribute{
										Computed:      true,
										Description:   "If a address postal code was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"cvc_check": schema.StringAttribute{
										Computed:      true,
										Description:   "If a CVC was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"country": schema.StringAttribute{
								Computed:      true,
								Description:   "Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"description": schema.StringAttribute{
								Computed:      true,
								Description:   "A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"exp_month": schema.Int64Attribute{
								Computed:      true,
								Description:   "Two-digit number representing the card's expiration month.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"exp_year": schema.Int64Attribute{
								Computed:      true,
								Description:   "Four-digit number representing the card's expiration year.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"extended_authorization": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"status": schema.StringAttribute{
										Computed:      true,
										Description:   "Indicates whether or not the capture window is extended beyond the standard authorization.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("disabled", "enabled")},
									},
								},
							},
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular card number. You can use this attribute to check whether two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.\n\n*As of May 1, 2021, card fingerprint in India for Connect changed to allow two fingerprints for the same card---one for India and one for the rest of the world.*",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"funding": schema.StringAttribute{
								Computed:      true,
								Description:   "Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"iin": schema.StringAttribute{
								Computed:      true,
								Description:   "Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"incremental_authorization": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"status": schema.StringAttribute{
										Computed:      true,
										Description:   "Indicates whether or not the incremental authorization feature is supported.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("available", "unavailable")},
									},
								},
							},
							"installments": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Installment details for this payment.\n\nFor more information, see the [installments integration guide](https://docs.stripe.com/payments/installments).",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"plan": schema.SingleNestedAttribute{
										Computed:      true,
										Description:   "Installment plan selected for the payment.",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
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
							},
							"issuer": schema.StringAttribute{
								Computed:      true,
								Description:   "The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "The last four digits of the card.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the mandate used to make this payment or created by it.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"moto": schema.BoolAttribute{
								Computed:      true,
								Description:   "True if this payment was marked as MOTO and out of scope for SCA.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"multicapture": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"status": schema.StringAttribute{
										Computed:      true,
										Description:   "Indicates whether or not multiple captures are supported.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("available", "unavailable")},
									},
								},
							},
							"network": schema.StringAttribute{
								Computed:      true,
								Description:   "Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"network_token": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "If this card has network token credentials, this contains the details of the network token credentials.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"used": schema.BoolAttribute{
										Computed:      true,
										Description:   "Indicates if Stripe used a network token, either user provided or Stripe managed when processing the transaction.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"network_transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"overcapture": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"maximum_amount_capturable": schema.Int64Attribute{
										Computed:      true,
										Description:   "The maximum amount that can be captured.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"status": schema.StringAttribute{
										Computed:      true,
										Description:   "Indicates whether or not the authorized amount can be over-captured.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("available", "unavailable")},
									},
								},
							},
							"regulated_status": schema.StringAttribute{
								Computed:      true,
								Description:   "Status of a card based on the card issuer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("regulated", "unregulated")},
							},
							"three_d_secure": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Populated if this transaction used 3D Secure authentication.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"authentication_flow": schema.StringAttribute{
										Computed:      true,
										Description:   "For authenticated transactions: how the customer was authenticated by\nthe issuing bank.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("challenge", "frictionless")},
									},
									"electronic_commerce_indicator": schema.StringAttribute{
										Computed:      true,
										Description:   "The Electronic Commerce Indicator (ECI). A protocol-level field\nindicating what degree of authentication was performed.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("01", "02", "05", "06", "07")},
									},
									"exemption_indicator": schema.StringAttribute{
										Computed:      true,
										Description:   "The exemption requested via 3DS and accepted by the issuer at authentication time.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("low_risk", "none")},
									},
									"exemption_indicator_applied": schema.BoolAttribute{
										Computed:      true,
										Description:   "Whether Stripe requested the value of `exemption_indicator` in the transaction. This will depend on\nthe outcome of Stripe's internal risk assessment.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
									"result": schema.StringAttribute{
										Computed:      true,
										Description:   "Indicates the outcome of 3D Secure authentication.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("attempt_acknowledged", "authenticated", "exempted", "failed", "not_supported", "processing_error")},
									},
									"result_reason": schema.StringAttribute{
										Computed:      true,
										Description:   "Additional information about why 3D Secure succeeded or failed based\non the `result`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("abandoned", "bypassed", "canceled", "card_not_enrolled", "network_not_supported", "protocol_error", "rejected")},
									},
									"transaction_id": schema.StringAttribute{
										Computed:      true,
										Description:   "The 3D Secure 1 XID or 3D Secure 2 Directory Server Transaction ID\n(dsTransId) for this payment.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"version": schema.StringAttribute{
										Computed:      true,
										Description:   "The version of 3D Secure that was used.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("1.0.2", "2.1.0", "2.2.0", "2.3.0", "2.3.1")},
									},
								},
							},
							"wallet": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "If this Card is part of a card wallet, this contains the details of the card wallet.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"dynamic_last4": schema.StringAttribute{
										Computed:      true,
										Description:   "(For tokenized numbers only.) The last four digits of the device account number.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"masterpass": schema.SingleNestedAttribute{
										Computed: true,

										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"billing_address": schema.SingleNestedAttribute{
												Computed:      true,
												Description:   "Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.",
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
											"email": schema.StringAttribute{
												Computed:      true,
												Description:   "Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"name": schema.StringAttribute{
												Computed:      true,
												Description:   "Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"shipping_address": schema.SingleNestedAttribute{
												Computed:      true,
												Description:   "Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.",
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
										},
									},
									"type": schema.StringAttribute{
										Computed:      true,
										Description:   "The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, `visa_checkout`, or `link`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("amex_express_checkout", "apple_pay", "google_pay", "link", "masterpass", "samsung_pay", "visa_checkout")},
									},
									"visa_checkout": schema.SingleNestedAttribute{
										Computed: true,

										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"billing_address": schema.SingleNestedAttribute{
												Computed:      true,
												Description:   "Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.",
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
											"email": schema.StringAttribute{
												Computed:      true,
												Description:   "Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"name": schema.StringAttribute{
												Computed:      true,
												Description:   "Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"shipping_address": schema.SingleNestedAttribute{
												Computed:      true,
												Description:   "Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.",
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
										},
									},
								},
							},
						},
					},
					"card_present": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"amount_authorized": schema.Int64Attribute{
								Computed:      true,
								Description:   "The authorized amount",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"brand": schema.StringAttribute{
								Computed:      true,
								Description:   "Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"brand_product": schema.StringAttribute{
								Computed:      true,
								Description:   "The [product code](https://stripe.com/docs/card-product-codes) that identifies the specific program or product associated with a card.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"capture_before": schema.Int64Attribute{
								Computed:      true,
								Description:   "When using manual capture, a future timestamp after which the charge will be automatically refunded if uncaptured.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"cardholder_name": schema.StringAttribute{
								Computed:      true,
								Description:   "The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"country": schema.StringAttribute{
								Computed:      true,
								Description:   "Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"description": schema.StringAttribute{
								Computed:      true,
								Description:   "A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"emv_auth_data": schema.StringAttribute{
								Computed:      true,
								Description:   "Authorization response cryptogram.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"exp_month": schema.Int64Attribute{
								Computed:      true,
								Description:   "Two-digit number representing the card's expiration month.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"exp_year": schema.Int64Attribute{
								Computed:      true,
								Description:   "Four-digit number representing the card's expiration year.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular card number. You can use this attribute to check whether two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.\n\n*As of May 1, 2021, card fingerprint in India for Connect changed to allow two fingerprints for the same card---one for India and one for the rest of the world.*",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"funding": schema.StringAttribute{
								Computed:      true,
								Description:   "Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"generated_card": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"iin": schema.StringAttribute{
								Computed:      true,
								Description:   "Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"incremental_authorization_supported": schema.BoolAttribute{
								Computed:      true,
								Description:   "Whether this [PaymentIntent](https://docs.stripe.com/api/payment_intents) is eligible for incremental authorizations. Request support using [request_incremental_authorization_support](https://docs.stripe.com/api/payment_intents/create#create_payment_intent-payment_method_options-card_present-request_incremental_authorization_support).",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"issuer": schema.StringAttribute{
								Computed:      true,
								Description:   "The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "The last four digits of the card.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"location": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"network": schema.StringAttribute{
								Computed:      true,
								Description:   "Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"network_transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"offline": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Details about payments collected offline.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"stored_at": schema.Int64Attribute{
										Computed:      true,
										Description:   "Time at which the payment was collected while offline",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"type": schema.StringAttribute{
										Computed:      true,
										Description:   "The method used to process this payment method offline. Only deferred is allowed.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("deferred")},
									},
								},
							},
							"overcapture_supported": schema.BoolAttribute{
								Computed:      true,
								Description:   "Defines whether the authorized amount can be over-captured or not",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preferred_locales": schema.ListAttribute{
								Computed:      true,
								Description:   "The languages that the issuing bank recommends using for localizing any customer-facing text, as read from the card. Referenced from EMV tag 5F2D, data encoded on the card's chip.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"read_method": schema.StringAttribute{
								Computed:      true,
								Description:   "How card details were read in this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("contact_emv", "contactless_emv", "contactless_magstripe_mode", "magnetic_stripe_fallback", "magnetic_stripe_track2")},
							},
							"reader": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"receipt": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "A collection of fields required to be displayed on receipts. Only required for EMV transactions.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"account_type": schema.StringAttribute{
										Computed:      true,
										Description:   "The type of account being debited or credited",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("checking", "credit", "prepaid", "unknown")},
									},
									"application_cryptogram": schema.StringAttribute{
										Computed:      true,
										Description:   "The Application Cryptogram, a unique value generated by the card to authenticate the transaction with issuers.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"application_preferred_name": schema.StringAttribute{
										Computed:      true,
										Description:   "The Application Identifier (AID) on the card used to determine which networks are eligible to process the transaction. Referenced from EMV tag 9F12, data encoded on the card's chip.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"authorization_code": schema.StringAttribute{
										Computed:      true,
										Description:   "Identifier for this transaction.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"authorization_response_code": schema.StringAttribute{
										Computed:      true,
										Description:   "EMV tag 8A. A code returned by the card issuer.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"cardholder_verification_method": schema.StringAttribute{
										Computed:      true,
										Description:   "Describes the method used by the cardholder to verify ownership of the card. One of the following: `approval`, `failure`, `none`, `offline_pin`, `offline_pin_and_signature`, `online_pin`, or `signature`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"dedicated_file_name": schema.StringAttribute{
										Computed:      true,
										Description:   "Similar to the application_preferred_name, identifying the applications (AIDs) available on the card. Referenced from EMV tag 84.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"terminal_verification_results": schema.StringAttribute{
										Computed:      true,
										Description:   "A 5-byte string that records the checks and validations that occur between the card and the terminal. These checks determine how the terminal processes the transaction and what risk tolerance is acceptable. Referenced from EMV Tag 95.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"transaction_status_information": schema.StringAttribute{
										Computed:      true,
										Description:   "An indication of which steps were completed during the card read process. Referenced from EMV Tag 9B.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"wallet": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Computed:      true,
										Description:   "The type of mobile wallet, one of `apple_pay`, `google_pay`, `samsung_pay`, or `unknown`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("apple_pay", "google_pay", "samsung_pay", "unknown")},
									},
								},
							},
						},
					},
					"cashapp": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"buyer_id": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique and immutable identifier assigned by Cash App to every buyer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"cashtag": schema.StringAttribute{
								Computed:      true,
								Description:   "A public identifier for buyers using Cash App.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique and immutable identifier of payments assigned by Cash App",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"crypto": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"buyer_address": schema.StringAttribute{
								Computed:      true,
								Description:   "The wallet address of the customer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"network": schema.StringAttribute{
								Computed:      true,
								Description:   "The blockchain network that the transaction was sent on.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("base", "ethereum", "polygon", "solana", "tempo")},
							},
							"token_currency": schema.StringAttribute{
								Computed:      true,
								Description:   "The token currency that the transaction was sent with.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("phantom_cash", "usdc", "usdg", "usdp", "usdt")},
							},
							"transaction_hash": schema.StringAttribute{
								Computed:      true,
								Description:   "The blockchain transaction hash of the crypto payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"eps": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank": schema.StringAttribute{
								Computed:      true,
								Description:   "The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `deutsche_bank_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("arzte_und_apotheker_bank", "austrian_anadi_bank_ag", "bank_austria", "bankhaus_carl_spangler", "bankhaus_schelhammer_und_schattera_ag", "bawag_psk_ag", "bks_bank_ag", "brull_kallmus_bank_ag", "btv_vier_lander_bank", "capital_bank_grawe_gruppe_ag", "deutsche_bank_ag", "dolomitenbank", "easybank_ag", "erste_bank_und_sparkassen", "hypo_alpeadriabank_international_ag", "hypo_bank_burgenland_aktiengesellschaft", "hypo_noe_lb_fur_niederosterreich_u_wien", "hypo_oberosterreich_salzburg_steiermark", "hypo_tirol_bank_ag", "hypo_vorarlberg_bank_ag", "marchfelder_bank", "oberbank_ag", "raiffeisen_bankengruppe_osterreich", "schoellerbank_ag", "sparda_bank_wien", "volksbank_gruppe", "volkskreditbank_ag", "vr_bank_braunau")},
							},
							"verified_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Owner's verified full name. Values are verified or provided by EPS directly\n(if supported) at the time of authorization or settlement. They cannot be set or mutated.\nEPS rarely provides this information so the attribute is usually empty.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"fpx": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account_holder_type": schema.StringAttribute{
								Computed:      true,
								Description:   "Account holder type, if provided. Can be one of `individual` or `company`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("company", "individual")},
							},
							"bank": schema.StringAttribute{
								Computed:      true,
								Description:   "The customer's bank. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, `pb_enterprise`, or `bank_of_china`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("affin_bank", "agrobank", "alliance_bank", "ambank", "bank_islam", "bank_muamalat", "bank_of_china", "bank_rakyat", "bsn", "cimb", "deutsche_bank", "hong_leong_bank", "hsbc", "kfh", "maybank2e", "maybank2u", "ocbc", "pb_enterprise", "public_bank", "rhb", "standard_chartered", "uob")},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "Unique transaction id generated by FPX for every request from the merchant",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"giropay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank_code": schema.StringAttribute{
								Computed:      true,
								Description:   "Bank code of bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Name of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bic": schema.StringAttribute{
								Computed:      true,
								Description:   "Bank Identifier Code of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"verified_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Owner's verified full name. Values are verified or provided by Giropay directly\n(if supported) at the time of authorization or settlement. They cannot be set or mutated.\nGiropay rarely provides this information so the attribute is usually empty.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"grabpay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "Unique transaction id generated by GrabPay",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"ideal": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank": schema.StringAttribute{
								Computed:      true,
								Description:   "The customer's bank. Can be one of `abn_amro`, `adyen`, `asn_bank`, `bunq`, `buut`, `finom`, `handelsbanken`, `ing`, `knab`, `mollie`, `moneyou`, `n26`, `nn`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, `van_lanschot`, or `yoursafe`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("abn_amro", "adyen", "asn_bank", "bunq", "buut", "finom", "handelsbanken", "ing", "knab", "mollie", "moneyou", "n26", "nn", "rabobank", "regiobank", "revolut", "sns_bank", "triodos_bank", "van_lanschot", "yoursafe")},
							},
							"bic": schema.StringAttribute{
								Computed:      true,
								Description:   "The Bank Identifier Code of the customer's bank.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("ABNANL2A", "ADYBNL2A", "ASNBNL21", "BITSNL2A", "BUNQNL2A", "BUUTNL2A", "FNOMNL22", "FVLBNL22", "HANDNL2A", "INGBNL2A", "KNABNL2H", "MLLENL2A", "MOYONL21", "NNBANL2G", "NTSBDEB1", "RABONL2U", "RBRBNL21", "REVOIE23", "REVOLT21", "SNSBNL2A", "TRIONL2U")},
							},
							"generated_sepa_debit": schema.StringAttribute{
								Computed:      true,
								Description:   "The ID of the SEPA Direct Debit PaymentMethod which was generated by this Charge.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"generated_sepa_debit_mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "The mandate for the SEPA Direct Debit PaymentMethod which was generated by this Charge.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"iban_last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four characters of the IBAN.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "Unique transaction ID generated by iDEAL.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"verified_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Owner's verified full name. Values are verified or provided by iDEAL directly\n(if supported) at the time of authorization or settlement. They cannot be set or mutated.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"interac_present": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"brand": schema.StringAttribute{
								Computed:      true,
								Description:   "Card brand. Can be `interac`, `mastercard` or `visa`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"cardholder_name": schema.StringAttribute{
								Computed:      true,
								Description:   "The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"country": schema.StringAttribute{
								Computed:      true,
								Description:   "Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"description": schema.StringAttribute{
								Computed:      true,
								Description:   "A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"emv_auth_data": schema.StringAttribute{
								Computed:      true,
								Description:   "Authorization response cryptogram.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"exp_month": schema.Int64Attribute{
								Computed:      true,
								Description:   "Two-digit number representing the card's expiration month.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"exp_year": schema.Int64Attribute{
								Computed:      true,
								Description:   "Four-digit number representing the card's expiration year.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular card number. You can use this attribute to check whether two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.\n\n*As of May 1, 2021, card fingerprint in India for Connect changed to allow two fingerprints for the same card---one for India and one for the rest of the world.*",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"funding": schema.StringAttribute{
								Computed:      true,
								Description:   "Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"generated_card": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"iin": schema.StringAttribute{
								Computed:      true,
								Description:   "Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"issuer": schema.StringAttribute{
								Computed:      true,
								Description:   "The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "The last four digits of the card.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"location": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"network": schema.StringAttribute{
								Computed:      true,
								Description:   "Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"network_transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"preferred_locales": schema.ListAttribute{
								Computed:      true,
								Description:   "The languages that the issuing bank recommends using for localizing any customer-facing text, as read from the card. Referenced from EMV tag 5F2D, data encoded on the card's chip.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"read_method": schema.StringAttribute{
								Computed:      true,
								Description:   "How card details were read in this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("contact_emv", "contactless_emv", "contactless_magstripe_mode", "magnetic_stripe_fallback", "magnetic_stripe_track2")},
							},
							"reader": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"receipt": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "A collection of fields required to be displayed on receipts. Only required for EMV transactions.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"account_type": schema.StringAttribute{
										Computed:      true,
										Description:   "The type of account being debited or credited",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("checking", "savings", "unknown")},
									},
									"application_cryptogram": schema.StringAttribute{
										Computed:      true,
										Description:   "The Application Cryptogram, a unique value generated by the card to authenticate the transaction with issuers.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"application_preferred_name": schema.StringAttribute{
										Computed:      true,
										Description:   "The Application Identifier (AID) on the card used to determine which networks are eligible to process the transaction. Referenced from EMV tag 9F12, data encoded on the card's chip.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"authorization_code": schema.StringAttribute{
										Computed:      true,
										Description:   "Identifier for this transaction.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"authorization_response_code": schema.StringAttribute{
										Computed:      true,
										Description:   "EMV tag 8A. A code returned by the card issuer.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"cardholder_verification_method": schema.StringAttribute{
										Computed:      true,
										Description:   "Describes the method used by the cardholder to verify ownership of the card. One of the following: `approval`, `failure`, `none`, `offline_pin`, `offline_pin_and_signature`, `online_pin`, or `signature`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"dedicated_file_name": schema.StringAttribute{
										Computed:      true,
										Description:   "Similar to the application_preferred_name, identifying the applications (AIDs) available on the card. Referenced from EMV tag 84.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"terminal_verification_results": schema.StringAttribute{
										Computed:      true,
										Description:   "A 5-byte string that records the checks and validations that occur between the card and the terminal. These checks determine how the terminal processes the transaction and what risk tolerance is acceptable. Referenced from EMV Tag 95.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"transaction_status_information": schema.StringAttribute{
										Computed:      true,
										Description:   "An indication of which steps were completed during the card read process. Referenced from EMV Tag 9B.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
					"kakao_pay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"buyer_id": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique identifier for the buyer as determined by the local payment processor.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Kakao Pay transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"klarna": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"location": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payer_details": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "The payer details for this transaction.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"address": schema.SingleNestedAttribute{
										Computed:      true,
										Description:   "The payer's address",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"country": schema.StringAttribute{
												Computed:      true,
												Description:   "The payer address country",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
										},
									},
								},
							},
							"payment_method_category": schema.StringAttribute{
								Computed:      true,
								Description:   "The Klarna payment method used for this transaction.\nCan be one of `pay_later`, `pay_now`, `pay_with_financing`, or `pay_in_installments`",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"preferred_locale": schema.StringAttribute{
								Computed:      true,
								Description:   "Preferred language of the Klarna authorization page that the customer is redirected to.\nCan be one of `de-AT`, `en-AT`, `nl-BE`, `fr-BE`, `en-BE`, `de-DE`, `en-DE`, `da-DK`, `en-DK`, `es-ES`, `en-ES`, `fi-FI`, `sv-FI`, `en-FI`, `en-GB`, `en-IE`, `it-IT`, `en-IT`, `nl-NL`, `en-NL`, `nb-NO`, `en-NO`, `sv-SE`, `en-SE`, `en-US`, `es-US`, `fr-FR`, `en-FR`, `cs-CZ`, `en-CZ`, `ro-RO`, `en-RO`, `el-GR`, `en-GR`, `en-AU`, `en-NZ`, `en-CA`, `fr-CA`, `pl-PL`, `en-PL`, `pt-PT`, `en-PT`, `de-CH`, `fr-CH`, `it-CH`, or `en-CH`",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reader": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"konbini": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"store": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "If the payment succeeded, this contains the details of the convenience store where the payment was completed.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"chain": schema.StringAttribute{
										Computed:      true,
										Description:   "The name of the convenience store chain where the payment was completed.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("familymart", "lawson", "ministop", "seicomart")},
									},
								},
							},
						},
					},
					"kr_card": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"brand": schema.StringAttribute{
								Computed:      true,
								Description:   "The local credit or debit card brand.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("bc", "citi", "hana", "hyundai", "jeju", "jeonbuk", "kakaobank", "kbank", "kdbbank", "kookmin", "kwangju", "lotte", "mg", "nh", "post", "samsung", "savingsbank", "shinhan", "shinhyup", "suhyup", "tossbank", "woori")},
							},
							"buyer_id": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique identifier for the buyer as determined by the local payment processor.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "The last four digits of the card. This may not be present for American Express cards.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Korean Card transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"link": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"country": schema.StringAttribute{
								Computed:      true,
								Description:   "Two-letter ISO code representing the funding source country beneath the Link payment.\nYou could use this attribute to get a sense of international fees.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"mobilepay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"card": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Internal card details",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"brand": schema.StringAttribute{
										Computed:      true,
										Description:   "Brand of the card used in the transaction",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"country": schema.StringAttribute{
										Computed:      true,
										Description:   "Two-letter ISO code representing the country of the card",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"exp_month": schema.Int64Attribute{
										Computed:      true,
										Description:   "Two digit number representing the card's expiration month",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"exp_year": schema.Int64Attribute{
										Computed:      true,
										Description:   "Two digit number representing the card's expiration year",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"last4": schema.StringAttribute{
										Computed:      true,
										Description:   "The last 4 digits of the card",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
					"multibanco": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"entity": schema.StringAttribute{
								Computed:      true,
								Description:   "Entity number associated with this Multibanco payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reference": schema.StringAttribute{
								Computed:      true,
								Description:   "Reference number associated with this Multibanco payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"naver_pay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"buyer_id": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique identifier for the buyer as determined by the local payment processor.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Naver Pay transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"nz_bank_account": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account_holder_name": schema.StringAttribute{
								Computed:      true,
								Description:   "The name on the bank account. Only present if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod’s billing details.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bank_code": schema.StringAttribute{
								Computed:      true,
								Description:   "The numeric code for the bank account's bank.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "The name of the bank.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"branch_code": schema.StringAttribute{
								Computed:      true,
								Description:   "The numeric code for the bank account's bank branch.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expected_debit_date": schema.StringAttribute{
								Computed:      true,
								Description:   "Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four digits of the bank account number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"suffix": schema.StringAttribute{
								Computed:      true,
								Description:   "The suffix of the bank account number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"oxxo": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"number": schema.StringAttribute{
								Computed:      true,
								Description:   "OXXO reference number",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"p24": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank": schema.StringAttribute{
								Computed:      true,
								Description:   "The customer's bank. Can be one of `ing`, `citi_handlowy`, `tmobile_usbugi_bankowe`, `plus_bank`, `etransfer_pocztowy24`, `banki_spbdzielcze`, `bank_nowy_bfg_sa`, `getin_bank`, `velobank`, `blik`, `noble_pay`, `ideabank`, `envelobank`, `santander_przelew24`, `nest_przelew`, `mbank_mtransfer`, `inteligo`, `pbac_z_ipko`, `bnp_paribas`, `credit_agricole`, `toyota_bank`, `bank_pekao_sa`, `volkswagen_bank`, `bank_millennium`, `alior_bank`, or `boz`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("alior_bank", "bank_millennium", "bank_nowy_bfg_sa", "bank_pekao_sa", "banki_spbdzielcze", "blik", "bnp_paribas", "boz", "citi_handlowy", "credit_agricole", "envelobank", "etransfer_pocztowy24", "getin_bank", "ideabank", "ing", "inteligo", "mbank_mtransfer", "nest_przelew", "noble_pay", "pbac_z_ipko", "plus_bank", "santander_przelew24", "tmobile_usbugi_bankowe", "toyota_bank", "velobank", "volkswagen_bank")},
							},
							"reference": schema.StringAttribute{
								Computed:      true,
								Description:   "Unique reference for this Przelewy24 payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"verified_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Owner's verified full name. Values are verified or provided by Przelewy24 directly\n(if supported) at the time of authorization or settlement. They cannot be set or mutated.\nPrzelewy24 rarely provides this information so the attribute is usually empty.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"payco": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"buyer_id": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique identifier for the buyer as determined by the local payment processor.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Payco transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"paynow": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"location": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reader": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reference": schema.StringAttribute{
								Computed:      true,
								Description:   "Reference number associated with this PayNow payment",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"paypal": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"country": schema.StringAttribute{
								Computed:      true,
								Description:   "Two-letter ISO code representing the buyer's country. Values are provided by PayPal directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payer_email": schema.StringAttribute{
								Computed:      true,
								Description:   "Owner's email. Values are provided by PayPal directly\n(if supported) at the time of authorization or settlement. They cannot be set or mutated.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payer_id": schema.StringAttribute{
								Computed:      true,
								Description:   "PayPal account PayerID. This identifier uniquely identifies the PayPal customer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payer_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Owner's full name. Values provided by PayPal directly\n(if supported) at the time of authorization or settlement. They cannot be set or mutated.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"seller_protection": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "The level of protection offered as defined by PayPal Seller Protection for Merchants, for this transaction.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"dispute_categories": schema.ListAttribute{
										Computed:      true,
										Description:   "An array of conditions that are covered for the transaction, if applicable.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
									"status": schema.StringAttribute{
										Computed:      true,
										Description:   "Indicates whether the transaction is eligible for PayPal's seller protection.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("eligible", "not_eligible", "partially_eligible")},
									},
								},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique ID generated by PayPal for this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"payto": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bsb_number": schema.StringAttribute{
								Computed:      true,
								Description:   "Bank-State-Branch number of the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four digits of the bank account number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the mandate used to make this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"pay_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The PayID alias for the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"pix": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank_transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "Unique transaction id generated by BCB",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the multi use Mandate generated by the PaymentIntent",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"promptpay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"reference": schema.StringAttribute{
								Computed:      true,
								Description:   "Bill reference generated by PromptPay",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"revolut_pay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"funding": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"card": schema.SingleNestedAttribute{
										Computed: true,

										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"brand": schema.StringAttribute{
												Computed:      true,
												Description:   "Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"country": schema.StringAttribute{
												Computed:      true,
												Description:   "Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"exp_month": schema.Int64Attribute{
												Computed:      true,
												Description:   "Two-digit number representing the card's expiration month.",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
											},
											"exp_year": schema.Int64Attribute{
												Computed:      true,
												Description:   "Four-digit number representing the card's expiration year.",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
											},
											"funding": schema.StringAttribute{
												Computed:      true,
												Description:   "Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"last4": schema.StringAttribute{
												Computed:      true,
												Description:   "The last four digits of the card.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
										},
									},
									"type": schema.StringAttribute{
										Computed:      true,
										Description:   "funding type of the underlying payment method.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("card")},
									},
								},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Revolut Pay transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"samsung_pay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"buyer_id": schema.StringAttribute{
								Computed:      true,
								Description:   "A unique identifier for the buyer as determined by the local payment processor.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Samsung Pay transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"satispay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Satispay transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"scalapay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Scalapay transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"sepa_credit_transfer": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Name of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bic": schema.StringAttribute{
								Computed:      true,
								Description:   "Bank Identifier Code of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"iban": schema.StringAttribute{
								Computed:      true,
								Description:   "IBAN of the bank account to transfer funds to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"sepa_debit": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank_code": schema.StringAttribute{
								Computed:      true,
								Description:   "Bank code of bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"branch_code": schema.StringAttribute{
								Computed:      true,
								Description:   "Branch code of bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"country": schema.StringAttribute{
								Computed:      true,
								Description:   "Two-letter ISO code representing the country the bank account is located in.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expected_debit_date": schema.StringAttribute{
								Computed:      true,
								Description:   "Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four characters of the IBAN.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "Find the ID of the mandate used for this payment under the [payment_method_details.sepa_debit.mandate](https://docs.stripe.com/api/charges/object#charge_object-payment_method_details-sepa_debit-mandate) property on the Charge. Use this mandate ID to [retrieve the Mandate](https://docs.stripe.com/api/mandates/retrieve).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"sofort": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"bank_code": schema.StringAttribute{
								Computed:      true,
								Description:   "Bank code of bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Name of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bic": schema.StringAttribute{
								Computed:      true,
								Description:   "Bank Identifier Code of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"country": schema.StringAttribute{
								Computed:      true,
								Description:   "Two-letter ISO code representing the country the bank account is located in.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"generated_sepa_debit": schema.StringAttribute{
								Computed:      true,
								Description:   "The ID of the SEPA Direct Debit PaymentMethod which was generated by this Charge.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"generated_sepa_debit_mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "The mandate for the SEPA Direct Debit PaymentMethod which was generated by this Charge.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"iban_last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four characters of the IBAN.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"preferred_language": schema.StringAttribute{
								Computed:      true,
								Description:   "Preferred language of the SOFORT authorization page that the customer is redirected to.\nCan be one of `de`, `en`, `es`, `fr`, `it`, `nl`, or `pl`",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("de", "en", "es", "fr", "it", "nl", "pl")},
							},
							"verified_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Owner's verified full name. Values are verified or provided by SOFORT directly\n(if supported) at the time of authorization or settlement. They cannot be set or mutated.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"sunbit": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "The Sunbit transaction ID associated with this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"swish": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies the payer's Swish account. You can use this attribute to check whether two Swish transactions were paid for by the same payer",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payment_reference": schema.StringAttribute{
								Computed:      true,
								Description:   "Payer bank reference number for the payment",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"verified_phone_last4": schema.StringAttribute{
								Computed:      true,
								Description:   "The last four digits of the Swish account phone number",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"twint": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the multi use Mandate generated by the PaymentIntent",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "The type of transaction-specific details of the payment method used in the payment. See [PaymentMethod.type](https://docs.stripe.com/api/payment_methods/object#payment_method_object-type) for the full list of possible types.\nAn additional hash is included on `payment_method_details` with a name matching this value.\nIt contains information specific to the payment method.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"upi": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"vpa": schema.StringAttribute{
								Computed:      true,
								Description:   "Customer's unique Virtual Payment Address.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"us_bank_account": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account_holder_type": schema.StringAttribute{
								Computed:      true,
								Description:   "Account holder type: individual or company.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("company", "individual")},
							},
							"account_type": schema.StringAttribute{
								Computed:      true,
								Description:   "Account type: checkings or savings. Defaults to checking if omitted.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("checking", "savings")},
							},
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "Name of the bank associated with the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expected_debit_date": schema.StringAttribute{
								Computed:      true,
								Description:   "Estimated date to debit the customer's bank account. A date string in YYYY-MM-DD format.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"last4": schema.StringAttribute{
								Computed:      true,
								Description:   "Last four digits of the bank account number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"mandate": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the mandate used to make this payment.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payment_reference": schema.StringAttribute{
								Computed:      true,
								Description:   "Reference number to locate ACH payments with customer's bank.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"routing_number": schema.StringAttribute{
								Computed:      true,
								Description:   "Routing number of the bank account.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"wechat_pay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fingerprint": schema.StringAttribute{
								Computed:      true,
								Description:   "Uniquely identifies this particular WeChat Pay account. You can use this attribute to check whether two WeChat accounts are the same.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"location": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the [location](https://docs.stripe.com/api/terminal/locations) that this transaction's reader is assigned to.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"reader": schema.StringAttribute{
								Computed:      true,
								Description:   "ID of the [reader](https://docs.stripe.com/api/terminal/readers) this transaction was made on.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"transaction_id": schema.StringAttribute{
								Computed:      true,
								Description:   "Transaction ID of this particular WeChat Pay transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"radar_options": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"session": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"receipt_email": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "This is the email address that the receipt for this charge was sent to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"receipt_number": schema.StringAttribute{
				Computed:      true,
				Description:   "This is the transaction number that appears on email receipts sent for this charge. This attribute will be `null` until a receipt has been sent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"refunded": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether the charge has been fully refunded. If the charge is only partially refunded, this attribute will still be false.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"review": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the review associated with this charge if one exists.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"shipping": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Shipping information for the charge.",
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
				Optional:      true,
				Computed:      true,
				Description:   "This is a legacy field that will be removed in the future. It contains the Source, Card, or BankAccount object used for the charge. For details about the payment method used for this charge, refer to `payment_method` or `payment_method_details` instead.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"source_transfer": schema.StringAttribute{
				Computed:      true,
				Description:   "The transfer ID which created this charge. Only present if the charge came from another Stripe account. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"statement_descriptor": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "For a non-card charge, text that appears on the customer's statement as the statement descriptor. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).\n\nFor a card charge, this value is ignored unless you don't specify a `statement_descriptor_suffix`, in which case this value is used as the suffix.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"statement_descriptor_suffix": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement. If the account has no prefix value, the suffix is concatenated to the account's statement descriptor.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The status of the payment is either `succeeded`, `pending`, or `failed`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("failed", "pending", "succeeded")},
			},
			"transfer": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the transfer to the `destination` account (only applicable if the charge was created using the `destination` parameter).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"transfer_data": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://docs.stripe.com/connect/destination-charges) for details.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"amount": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The amount transferred to the destination account, if specified. By default, the entire charge amount is transferred to the destination account.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"destination": schema.StringAttribute{
						Required:      true,
						Description:   "ID of an existing, connected Stripe account to transfer funds to if `transfer_data` was specified in the charge request.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"description": schema.StringAttribute{
						Optional:      true,
						Description:   "An arbitrary string attached to the transfer. Often useful for displaying to users.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"transfer_group": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A string that identifies this transaction as part of a group. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options) for details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"capture": schema.BoolAttribute{
				Optional:      true,
				Description:   "Whether to immediately capture the charge. Defaults to `true`. When `false`, the charge issues an authorization (or pre-authorization), and will need to be [captured](https://api.stripe.com#capture_charge) later. Uncaptured charges expire after a set number of days (7 by default). For more information, see the [authorizing charges and settling later](https://docs.stripe.com/charges/placing-a-hold) documentation.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"destination": schema.SingleNestedAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				WriteOnly:     true,
				Attributes: map[string]schema.Attribute{
					"account": schema.StringAttribute{
						Required:      true,
						Description:   "ID of an existing, connected Stripe account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
					"amount": schema.Int64Attribute{
						Optional:      true,
						Description:   "The amount to transfer to the destination account without creating an `Application Fee` object. Cannot be combined with the `application_fee` parameter. Must be less than or equal to the charge amount.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
				},
			},
		},
	}
}

func (r *ChargeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ChargeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config ChargeResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Capture"}, []string{"Destination"}, []string{"Destination", "account"}, []string{"Destination", "amount"}})

	params, err := expandChargeCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Charge create params", err.Error())
		return
	}

	obj, err := r.client.V1Charges.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Charge", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Charges.B, r.client.V1Charges.Key, stripe.FormatURLPath("/v1/charges/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Charge create raw response", err.Error())
		return
	}

	var createdState ChargeResourceModel
	if err := flattenCharge(obj, &createdState); err != nil {
		resp.Diagnostics.AddError("Error flattening Charge create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&createdState, &config, [][]string{[]string{"TransferData", "description"}})
	normalizeUnknownValues(&createdState)

	diffPlan := plan
	diffCreatedState := createdState
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"Capture"}, []string{"Destination"}, []string{"Destination", "account"}, []string{"Destination", "amount"}})
	clearWriteOnlyPaths(&diffCreatedState, [][]string{[]string{"Capture"}, []string{"Destination"}, []string{"Destination", "account"}, []string{"Destination", "amount"}})

	postCreateParams, err := expandChargePostCreateUpdate(diffPlan, diffCreatedState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Charge post-create update params", err.Error())
		return
	}

	if paramsHaveValues(postCreateParams) {
		if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
			!createdState.Metadata.IsNull() && !createdState.Metadata.IsUnknown() {
			if !assignMetadataDiffToNamedField(postCreateParams, "Metadata", plan.Metadata, createdState.Metadata) {
				resp.Diagnostics.AddError("Error building Charge update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", postCreateParams))
				return
			}
		}
		obj, err = r.client.V1Charges.Update(ctx, createdState.ID.ValueString(), postCreateParams)
		if err != nil {
			resp.Diagnostics.AddError("Error finalizing Charge after create", err.Error())
			return
		}
		if err := ensureRawResponse(obj, r.client.V1Charges.B, r.client.V1Charges.Key, stripe.FormatURLPath("/v1/charges/%s", obj.ID), nil); err != nil {
			resp.Diagnostics.AddError("Error hydrating Charge post-create update raw response", err.Error())
			return
		}
	}

	if err := flattenCharge(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Charge create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"TransferData", "description"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Capture"}, []string{"Destination"}, []string{"Destination", "account"}, []string{"Destination", "amount"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ChargeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState ChargeResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ChargeResourceModel
	state = priorState

	obj, err := r.client.V1Charges.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Charge", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Charges.B, r.client.V1Charges.Key, stripe.FormatURLPath("/v1/charges/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Charge raw response", err.Error())
		return
	}

	if err := flattenCharge(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Charge read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"TransferData", "description"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"Capture"}, []string{"Destination"}, []string{"Destination", "account"}, []string{"Destination", "amount"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *ChargeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ChargeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config ChargeResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Capture"}, []string{"Destination"}, []string{"Destination", "account"}, []string{"Destination", "amount"}})

	var state ChargeResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"Capture"}, []string{"Destination"}, []string{"Destination", "account"}, []string{"Destination", "amount"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"Capture"}, []string{"Destination"}, []string{"Destination", "account"}, []string{"Destination", "amount"}})

	params, err := expandChargeUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Charge update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Charge update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Charges.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Charge", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Charges.B, r.client.V1Charges.Key, stripe.FormatURLPath("/v1/charges/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Charge update raw response", err.Error())
		return
	}

	if err := flattenCharge(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Charge update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"TransferData", "description"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Capture"}, []string{"Destination"}, []string{"Destination", "account"}, []string{"Destination", "amount"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ChargeResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *ChargeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandChargeCreate(plan ChargeResourceModel) (*stripe.ChargeCreateParams, error) {
	params := &stripe.ChargeCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.ApplicationFee.IsNull() && !plan.ApplicationFee.IsUnknown() {
		params.ApplicationFee = stripe.Int64(plan.ApplicationFee.ValueInt64())
	}
	if !plan.ApplicationFeeAmount.IsNull() && !plan.ApplicationFeeAmount.IsUnknown() {
		params.ApplicationFeeAmount = stripe.Int64(plan.ApplicationFeeAmount.ValueInt64())
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
	if !plan.RadarOptions.IsNull() && !plan.RadarOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RadarOptions", plan.RadarOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "radar_options", params)
		}
	}
	if !plan.ReceiptEmail.IsNull() && !plan.ReceiptEmail.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReceiptEmail", "ReceiptEmail", plan.ReceiptEmail.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "receipt_email", params)
		}
	}
	if !plan.Shipping.IsNull() && !plan.Shipping.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Shipping", plan.Shipping) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping", params)
		}
	}
	if !plan.Source.IsNull() && !plan.Source.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SourceID", "Source", plan.Source.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "source", params)
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
	if !plan.Capture.IsNull() && !plan.Capture.IsUnknown() {
		params.Capture = stripe.Bool(plan.Capture.ValueBool())
	}
	if !plan.Destination.IsNull() && !plan.Destination.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Destination", plan.Destination) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "destination", params)
		}
	}

	return params, nil
}

func expandChargeUpdate(plan ChargeResourceModel, state ChargeResourceModel) (*stripe.ChargeUpdateParams, error) {
	params := &stripe.ChargeUpdateParams{}

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
	if !plan.FraudDetails.Equal(state.FraudDetails) && !plan.FraudDetails.IsNull() && !plan.FraudDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FraudDetails", plan.FraudDetails) {
			if !plan.FraudDetails.Equal(state.FraudDetails) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "fraud_details", params)
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
	if !plan.ReceiptEmail.Equal(state.ReceiptEmail) && !plan.ReceiptEmail.IsNull() && !plan.ReceiptEmail.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReceiptEmail", "ReceiptEmail", plan.ReceiptEmail.ValueString()) {
			if !plan.ReceiptEmail.Equal(state.ReceiptEmail) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "receipt_email", params)
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
	if !plan.TransferGroup.Equal(state.TransferGroup) && !plan.TransferGroup.IsNull() && !plan.TransferGroup.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TransferGroup", "TransferGroup", plan.TransferGroup.ValueString()) {
			if !plan.TransferGroup.Equal(state.TransferGroup) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "transfer_group", params)
			}
		}
	}

	return params, nil
}

func expandChargePostCreateUpdate(plan ChargeResourceModel, state ChargeResourceModel) (*stripe.ChargeUpdateParams, error) {
	params := &stripe.ChargeUpdateParams{}

	if !plan.FraudDetails.Equal(state.FraudDetails) && !plan.FraudDetails.IsNull() && !plan.FraudDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FraudDetails", plan.FraudDetails) {
			if !plan.FraudDetails.Equal(state.FraudDetails) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "fraud_details", params)
			}
		}
	}

	return params, nil
}

func flattenCharge(obj *stripe.Charge, state *ChargeResourceModel) error {
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
		if rawValueAmountCaptured, rawOk := plainValueAtPath(raw, "amount_captured"); rawOk {
			if valueAmountCaptured, err := flattenPlainValue(rawValueAmountCaptured, types.Int64Type, "amount_captured", "raw response"); err != nil {
				return err
			} else {
				if typedAmountCaptured, ok := valueAmountCaptured.(types.Int64); ok {
					state.AmountCaptured = typedAmountCaptured
				}
			}
		} else if !hasRaw {
			if responseValueAmountCaptured, ok := plainFromResponseField(obj, "AmountCaptured"); ok {
				if valueAmountCaptured, err := flattenPlainValue(responseValueAmountCaptured, types.Int64Type, "amount_captured", "response struct"); err != nil {
					return err
				} else {
					if typedAmountCaptured, ok := valueAmountCaptured.(types.Int64); ok {
						state.AmountCaptured = typedAmountCaptured
					}
				}
			}
		}
	}
	{
		if rawValueAmountRefunded, rawOk := plainValueAtPath(raw, "amount_refunded"); rawOk {
			if valueAmountRefunded, err := flattenPlainValue(rawValueAmountRefunded, types.Int64Type, "amount_refunded", "raw response"); err != nil {
				return err
			} else {
				if typedAmountRefunded, ok := valueAmountRefunded.(types.Int64); ok {
					state.AmountRefunded = typedAmountRefunded
				}
			}
		} else if !hasRaw {
			if responseValueAmountRefunded, ok := plainFromResponseField(obj, "AmountRefunded"); ok {
				if valueAmountRefunded, err := flattenPlainValue(responseValueAmountRefunded, types.Int64Type, "amount_refunded", "response struct"); err != nil {
					return err
				} else {
					if typedAmountRefunded, ok := valueAmountRefunded.(types.Int64); ok {
						state.AmountRefunded = typedAmountRefunded
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
		if rawValueApplicationFee, rawOk := plainValueAtPath(raw, "application_fee"); rawOk {
			if valueApplicationFee, err := flattenPlainValue(rawValueApplicationFee, types.Int64Type, "application_fee", "raw response"); err != nil {
				return err
			} else {
				if typedApplicationFee, ok := valueApplicationFee.(types.Int64); ok {
					state.ApplicationFee = typedApplicationFee
				}
			}
		} else if !hasRaw {
			if responseValueApplicationFee, ok := plainFromResponseField(obj, "ApplicationFee"); ok {
				if valueApplicationFee, err := flattenPlainValue(responseValueApplicationFee, types.Int64Type, "application_fee", "response struct"); err != nil {
					return err
				} else {
					if typedApplicationFee, ok := valueApplicationFee.(types.Int64); ok {
						state.ApplicationFee = typedApplicationFee
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
		if rawValueAuthorizationCode, rawOk := plainValueAtPath(raw, "authorization_code"); rawOk {
			if valueAuthorizationCode, err := flattenPlainValue(rawValueAuthorizationCode, types.StringType, "authorization_code", "raw response"); err != nil {
				return err
			} else {
				if typedAuthorizationCode, ok := valueAuthorizationCode.(types.String); ok {
					state.AuthorizationCode = typedAuthorizationCode
				}
			}
		} else if !hasRaw {
			if responseValueAuthorizationCode, ok := plainFromResponseField(obj, "AuthorizationCode"); ok {
				if valueAuthorizationCode, err := flattenPlainValue(responseValueAuthorizationCode, types.StringType, "authorization_code", "response struct"); err != nil {
					return err
				} else {
					if typedAuthorizationCode, ok := valueAuthorizationCode.(types.String); ok {
						state.AuthorizationCode = typedAuthorizationCode
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueBalanceTransaction, rawOk := plainValueAtPath(raw, "balance_transaction"); rawOk {
				if typedBalanceTransaction, ok := plainToStringIDValue(rawValueBalanceTransaction); ok {
					state.BalanceTransaction = typedBalanceTransaction
				}
			} else if !hasRaw {
				if responseValueBalanceTransaction, ok := plainFromResponseField(obj, "BalanceTransaction"); ok {
					if typedBalanceTransaction, ok := plainToStringIDValue(responseValueBalanceTransaction); ok {
						state.BalanceTransaction = typedBalanceTransaction
					}
				}
			}
		}
	}
	{
		assignedBillingDetails := false
		hadRawBillingDetails := false
		if rawValueBillingDetails, rawOk := plainValueAtPath(raw, "billing_details"); rawOk {
			hadRawBillingDetails = true
			if rawValueBillingDetails != nil {
				sourceBillingDetails := applyConfiguredKeyedListShapes(rawValueBillingDetails, attrValueToPlain(state.BillingDetails))
				if valueBillingDetails, err := flattenPlainValue(sourceBillingDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "phone": types.StringType, "tax_id": types.StringType}}, "billing_details", "raw response"); err != nil {
					return err
				} else {
					if typedBillingDetails, ok := valueBillingDetails.(types.Object); ok {
						state.BillingDetails = typedBillingDetails
						assignedBillingDetails = true
					}
				}
			}
		}
		if !assignedBillingDetails {
			if !hasRaw {
				if responseValueBillingDetails, ok := plainFromResponseField(obj, "BillingDetails"); ok {
					sourceBillingDetails := applyConfiguredKeyedListShapes(responseValueBillingDetails, attrValueToPlain(state.BillingDetails))
					if valueBillingDetails, err := flattenPlainValue(
						sourceBillingDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "phone": types.StringType, "tax_id": types.StringType}},
						"billing_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBillingDetails, ok := valueBillingDetails.(types.Object); ok {
							state.BillingDetails = typedBillingDetails
							assignedBillingDetails = true
						}
					}
				}
			}
		}
		if !assignedBillingDetails && hadRawBillingDetails {
			if nullBillingDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "phone": types.StringType, "tax_id": types.StringType}}); ok {
				if typedBillingDetails, ok := nullBillingDetails.(types.Object); ok {
					state.BillingDetails = typedBillingDetails
				}
			}
		}
	}
	{
		if rawValueCalculatedStatementDescriptor, rawOk := plainValueAtPath(raw, "calculated_statement_descriptor"); rawOk {
			if valueCalculatedStatementDescriptor, err := flattenPlainValue(rawValueCalculatedStatementDescriptor, types.StringType, "calculated_statement_descriptor", "raw response"); err != nil {
				return err
			} else {
				if typedCalculatedStatementDescriptor, ok := valueCalculatedStatementDescriptor.(types.String); ok {
					state.CalculatedStatementDescriptor = typedCalculatedStatementDescriptor
				}
			}
		} else if !hasRaw {
			if responseValueCalculatedStatementDescriptor, ok := plainFromResponseField(obj, "CalculatedStatementDescriptor"); ok {
				if valueCalculatedStatementDescriptor, err := flattenPlainValue(responseValueCalculatedStatementDescriptor, types.StringType, "calculated_statement_descriptor", "response struct"); err != nil {
					return err
				} else {
					if typedCalculatedStatementDescriptor, ok := valueCalculatedStatementDescriptor.(types.String); ok {
						state.CalculatedStatementDescriptor = typedCalculatedStatementDescriptor
					}
				}
			}
		}
	}
	{
		if rawValueCaptured, rawOk := plainValueAtPath(raw, "captured"); rawOk {
			if valueCaptured, err := flattenPlainValue(rawValueCaptured, types.BoolType, "captured", "raw response"); err != nil {
				return err
			} else {
				if typedCaptured, ok := valueCaptured.(types.Bool); ok {
					state.Captured = typedCaptured
				}
			}
		} else if !hasRaw {
			if responseValueCaptured, ok := plainFromResponseField(obj, "Captured"); ok {
				if valueCaptured, err := flattenPlainValue(responseValueCaptured, types.BoolType, "captured", "response struct"); err != nil {
					return err
				} else {
					if typedCaptured, ok := valueCaptured.(types.Bool); ok {
						state.Captured = typedCaptured
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
		if rawValueDisputed, rawOk := plainValueAtPath(raw, "disputed"); rawOk {
			if valueDisputed, err := flattenPlainValue(rawValueDisputed, types.BoolType, "disputed", "raw response"); err != nil {
				return err
			} else {
				if typedDisputed, ok := valueDisputed.(types.Bool); ok {
					state.Disputed = typedDisputed
				}
			}
		} else if !hasRaw {
			if responseValueDisputed, ok := plainFromResponseField(obj, "Disputed"); ok {
				if valueDisputed, err := flattenPlainValue(responseValueDisputed, types.BoolType, "disputed", "response struct"); err != nil {
					return err
				} else {
					if typedDisputed, ok := valueDisputed.(types.Bool); ok {
						state.Disputed = typedDisputed
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueFailureBalanceTransaction, rawOk := plainValueAtPath(raw, "failure_balance_transaction"); rawOk {
				if typedFailureBalanceTransaction, ok := plainToStringIDValue(rawValueFailureBalanceTransaction); ok {
					state.FailureBalanceTransaction = typedFailureBalanceTransaction
				}
			} else if !hasRaw {
				if responseValueFailureBalanceTransaction, ok := plainFromResponseField(obj, "FailureBalanceTransaction"); ok {
					if typedFailureBalanceTransaction, ok := plainToStringIDValue(responseValueFailureBalanceTransaction); ok {
						state.FailureBalanceTransaction = typedFailureBalanceTransaction
					}
				}
			}
		}
	}
	{
		if rawValueFailureCode, rawOk := plainValueAtPath(raw, "failure_code"); rawOk {
			if valueFailureCode, err := flattenPlainValue(rawValueFailureCode, types.StringType, "failure_code", "raw response"); err != nil {
				return err
			} else {
				if typedFailureCode, ok := valueFailureCode.(types.String); ok {
					state.FailureCode = typedFailureCode
				}
			}
		} else if !hasRaw {
			if responseValueFailureCode, ok := plainFromResponseField(obj, "FailureCode"); ok {
				if valueFailureCode, err := flattenPlainValue(responseValueFailureCode, types.StringType, "failure_code", "response struct"); err != nil {
					return err
				} else {
					if typedFailureCode, ok := valueFailureCode.(types.String); ok {
						state.FailureCode = typedFailureCode
					}
				}
			}
		}
	}
	{
		if rawValueFailureMessage, rawOk := plainValueAtPath(raw, "failure_message"); rawOk {
			if valueFailureMessage, err := flattenPlainValue(rawValueFailureMessage, types.StringType, "failure_message", "raw response"); err != nil {
				return err
			} else {
				if typedFailureMessage, ok := valueFailureMessage.(types.String); ok {
					state.FailureMessage = typedFailureMessage
				}
			}
		} else if !hasRaw {
			if responseValueFailureMessage, ok := plainFromResponseField(obj, "FailureMessage"); ok {
				if valueFailureMessage, err := flattenPlainValue(responseValueFailureMessage, types.StringType, "failure_message", "response struct"); err != nil {
					return err
				} else {
					if typedFailureMessage, ok := valueFailureMessage.(types.String); ok {
						state.FailureMessage = typedFailureMessage
					}
				}
			}
		}
	}
	{
		assignedFraudDetails := false
		hadRawFraudDetails := false
		if rawValueFraudDetails, rawOk := plainValueAtPath(raw, "fraud_details"); rawOk {
			hadRawFraudDetails = true
			if rawValueFraudDetails != nil {
				sourceFraudDetails := applyConfiguredKeyedListShapes(rawValueFraudDetails, attrValueToPlain(state.FraudDetails))
				if valueFraudDetails, err := flattenPlainValue(sourceFraudDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"stripe_report": types.StringType, "user_report": types.StringType}}, "fraud_details", "raw response"); err != nil {
					return err
				} else {
					if typedFraudDetails, ok := valueFraudDetails.(types.Object); ok {
						state.FraudDetails = typedFraudDetails
						assignedFraudDetails = true
					}
				}
			}
		}
		if !assignedFraudDetails {
			if !hasRaw {
				if responseValueFraudDetails, ok := plainFromResponseField(obj, "FraudDetails"); ok {
					sourceFraudDetails := applyConfiguredKeyedListShapes(responseValueFraudDetails, attrValueToPlain(state.FraudDetails))
					if valueFraudDetails, err := flattenPlainValue(
						sourceFraudDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"stripe_report": types.StringType, "user_report": types.StringType}},
						"fraud_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFraudDetails, ok := valueFraudDetails.(types.Object); ok {
							state.FraudDetails = typedFraudDetails
							assignedFraudDetails = true
						}
					}
				}
			}
		}
		if !assignedFraudDetails && hadRawFraudDetails {
			if nullFraudDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"stripe_report": types.StringType, "user_report": types.StringType}}); ok {
				if typedFraudDetails, ok := nullFraudDetails.(types.Object); ok {
					state.FraudDetails = typedFraudDetails
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
		assignedLevel3 := false
		hadRawLevel3 := false
		if rawValueLevel3, rawOk := plainValueAtPath(raw, "level3"); rawOk {
			hadRawLevel3 = true
			if rawValueLevel3 != nil {
				sourceLevel3 := applyConfiguredKeyedListShapes(rawValueLevel3, attrValueToPlain(state.Level3))
				if valueLevel3, err := flattenPlainValue(sourceLevel3, types.ObjectType{AttrTypes: map[string]attr.Type{"customer_reference": types.StringType, "line_items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"discount_amount": types.Int64Type, "product_code": types.StringType, "product_description": types.StringType, "quantity": types.Int64Type, "tax_amount": types.Int64Type, "unit_cost": types.Int64Type}}}, "merchant_reference": types.StringType, "shipping_address_zip": types.StringType, "shipping_amount": types.Int64Type, "shipping_from_zip": types.StringType}}, "level3", "raw response"); err != nil {
					return err
				} else {
					if typedLevel3, ok := valueLevel3.(types.Object); ok {
						state.Level3 = typedLevel3
						assignedLevel3 = true
					}
				}
			}
		}
		if !assignedLevel3 {
			if !hasRaw {
				if responseValueLevel3, ok := plainFromResponseField(obj, "Level3"); ok {
					sourceLevel3 := applyConfiguredKeyedListShapes(responseValueLevel3, attrValueToPlain(state.Level3))
					if valueLevel3, err := flattenPlainValue(
						sourceLevel3,
						types.ObjectType{AttrTypes: map[string]attr.Type{"customer_reference": types.StringType, "line_items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"discount_amount": types.Int64Type, "product_code": types.StringType, "product_description": types.StringType, "quantity": types.Int64Type, "tax_amount": types.Int64Type, "unit_cost": types.Int64Type}}}, "merchant_reference": types.StringType, "shipping_address_zip": types.StringType, "shipping_amount": types.Int64Type, "shipping_from_zip": types.StringType}},
						"level3",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLevel3, ok := valueLevel3.(types.Object); ok {
							state.Level3 = typedLevel3
							assignedLevel3 = true
						}
					}
				}
			}
		}
		if !assignedLevel3 && hadRawLevel3 {
			if nullLevel3, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"customer_reference": types.StringType, "line_items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"discount_amount": types.Int64Type, "product_code": types.StringType, "product_description": types.StringType, "quantity": types.Int64Type, "tax_amount": types.Int64Type, "unit_cost": types.Int64Type}}}, "merchant_reference": types.StringType, "shipping_address_zip": types.StringType, "shipping_amount": types.Int64Type, "shipping_from_zip": types.StringType}}); ok {
				if typedLevel3, ok := nullLevel3.(types.Object); ok {
					state.Level3 = typedLevel3
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
		assignedOutcome := false
		hadRawOutcome := false
		if rawValueOutcome, rawOk := plainValueAtPath(raw, "outcome"); rawOk {
			hadRawOutcome = true
			if rawValueOutcome != nil {
				sourceOutcome := applyConfiguredKeyedListShapes(rawValueOutcome, attrValueToPlain(state.Outcome))
				if valueOutcome, err := flattenPlainValue(sourceOutcome, types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "network_status": types.StringType, "reason": types.StringType, "risk_level": types.StringType, "risk_score": types.Int64Type, "rule": types.StringType, "seller_message": types.StringType, "type": types.StringType}}, "outcome", "raw response"); err != nil {
					return err
				} else {
					if typedOutcome, ok := valueOutcome.(types.Object); ok {
						state.Outcome = typedOutcome
						assignedOutcome = true
					}
				}
			}
		}
		if !assignedOutcome {
			if !hasRaw {
				if responseValueOutcome, ok := plainFromResponseField(obj, "Outcome"); ok {
					sourceOutcome := applyConfiguredKeyedListShapes(responseValueOutcome, attrValueToPlain(state.Outcome))
					if valueOutcome, err := flattenPlainValue(
						sourceOutcome,
						types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "network_status": types.StringType, "reason": types.StringType, "risk_level": types.StringType, "risk_score": types.Int64Type, "rule": types.StringType, "seller_message": types.StringType, "type": types.StringType}},
						"outcome",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedOutcome, ok := valueOutcome.(types.Object); ok {
							state.Outcome = typedOutcome
							assignedOutcome = true
						}
					}
				}
			}
		}
		if !assignedOutcome && hadRawOutcome {
			if nullOutcome, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "network_status": types.StringType, "reason": types.StringType, "risk_level": types.StringType, "risk_score": types.Int64Type, "rule": types.StringType, "seller_message": types.StringType, "type": types.StringType}}); ok {
				if typedOutcome, ok := nullOutcome.(types.Object); ok {
					state.Outcome = typedOutcome
				}
			}
		}
	}
	{
		if rawValuePaid, rawOk := plainValueAtPath(raw, "paid"); rawOk {
			if valuePaid, err := flattenPlainValue(rawValuePaid, types.BoolType, "paid", "raw response"); err != nil {
				return err
			} else {
				if typedPaid, ok := valuePaid.(types.Bool); ok {
					state.Paid = typedPaid
				}
			}
		} else if !hasRaw {
			if responseValuePaid, ok := plainFromResponseField(obj, "Paid"); ok {
				if valuePaid, err := flattenPlainValue(responseValuePaid, types.BoolType, "paid", "response struct"); err != nil {
					return err
				} else {
					if typedPaid, ok := valuePaid.(types.Bool); ok {
						state.Paid = typedPaid
					}
				}
			}
		}
	}
	{
		if true {
			if rawValuePaymentIntent, rawOk := plainValueAtPath(raw, "payment_intent"); rawOk {
				if typedPaymentIntent, ok := plainToStringIDValue(rawValuePaymentIntent); ok {
					state.PaymentIntent = typedPaymentIntent
				}
			} else if !hasRaw {
				if responseValuePaymentIntent, ok := plainFromResponseField(obj, "PaymentIntent"); ok {
					if typedPaymentIntent, ok := plainToStringIDValue(responseValuePaymentIntent); ok {
						state.PaymentIntent = typedPaymentIntent
					}
				}
			}
		}
	}
	{
		if rawValuePaymentMethod, rawOk := plainValueAtPath(raw, "payment_method"); rawOk {
			if valuePaymentMethod, err := flattenPlainValue(rawValuePaymentMethod, types.StringType, "payment_method", "raw response"); err != nil {
				return err
			} else {
				if typedPaymentMethod, ok := valuePaymentMethod.(types.String); ok {
					state.PaymentMethod = typedPaymentMethod
				}
			}
		} else if !hasRaw {
			if responseValuePaymentMethod, ok := plainFromResponseField(obj, "PaymentMethod"); ok {
				if valuePaymentMethod, err := flattenPlainValue(responseValuePaymentMethod, types.StringType, "payment_method", "response struct"); err != nil {
					return err
				} else {
					if typedPaymentMethod, ok := valuePaymentMethod.(types.String); ok {
						state.PaymentMethod = typedPaymentMethod
					}
				}
			}
		}
	}
	{
		assignedPaymentMethodDetails := false
		hadRawPaymentMethodDetails := false
		if rawValuePaymentMethodDetails, rawOk := plainValueAtPath(raw, "payment_method_details"); rawOk {
			hadRawPaymentMethodDetails = true
			if rawValuePaymentMethodDetails != nil {
				sourcePaymentMethodDetails := applyConfiguredKeyedListShapes(rawValuePaymentMethodDetails, attrValueToPlain(state.PaymentMethodDetails))
				if valuePaymentMethodDetails, err := flattenPlainValue(sourcePaymentMethodDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"ach_credit_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"account_number": types.StringType, "bank_name": types.StringType, "routing_number": types.StringType, "swift_code": types.StringType}}, "ach_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "bank_name": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "routing_number": types.StringType}}, "acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "institution_number": types.StringType, "last4": types.StringType, "mandate": types.StringType, "transit_number": types.StringType}}, "affirm": types.ObjectType{AttrTypes: map[string]attr.Type{"location": types.StringType, "reader": types.StringType, "transaction_id": types.StringType}}, "afterpay_clearpay": types.ObjectType{AttrTypes: map[string]attr.Type{"order_id": types.StringType, "reference": types.StringType}}, "alipay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "fingerprint": types.StringType, "transaction_id": types.StringType}}, "alma": types.ObjectType{AttrTypes: map[string]attr.Type{"installments": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type}}, "transaction_id": types.StringType}}, "amazon_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"funding": types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "country": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "funding": types.StringType, "last4": types.StringType}}, "type": types.StringType}}, "transaction_id": types.StringType}}, "au_becs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType, "sort_code": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "generated_sepa_debit": types.StringType, "generated_sepa_debit_mandate": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "verified_name": types.StringType}}, "billie": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "bizum": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "blik": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType}}, "boleto": types.ObjectType{AttrTypes: map[string]attr.Type{"tax_id": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_authorized": types.Int64Type, "authorization_code": types.StringType, "brand": types.StringType, "capture_before": types.Int64Type, "checks": types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_postal_code_check": types.StringType, "cvc_check": types.StringType}}, "country": types.StringType, "description": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "extended_authorization": types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "incremental_authorization": types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}, "installments": types.ObjectType{AttrTypes: map[string]attr.Type{"plan": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}}, "issuer": types.StringType, "last4": types.StringType, "mandate": types.StringType, "moto": types.BoolType, "multicapture": types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}, "network": types.StringType, "network_token": types.ObjectType{AttrTypes: map[string]attr.Type{"used": types.BoolType}}, "network_transaction_id": types.StringType, "overcapture": types.ObjectType{AttrTypes: map[string]attr.Type{"maximum_amount_capturable": types.Int64Type, "status": types.StringType}}, "regulated_status": types.StringType, "three_d_secure": types.ObjectType{AttrTypes: map[string]attr.Type{"authentication_flow": types.StringType, "electronic_commerce_indicator": types.StringType, "exemption_indicator": types.StringType, "exemption_indicator_applied": types.BoolType, "result": types.StringType, "result_reason": types.StringType, "transaction_id": types.StringType, "version": types.StringType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"dynamic_last4": types.StringType, "masterpass": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}, "type": types.StringType, "visa_checkout": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}}}}}, "card_present": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_authorized": types.Int64Type, "brand": types.StringType, "brand_product": types.StringType, "capture_before": types.Int64Type, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_card": types.StringType, "iin": types.StringType, "incremental_authorization_supported": types.BoolType, "issuer": types.StringType, "last4": types.StringType, "location": types.StringType, "network": types.StringType, "network_transaction_id": types.StringType, "offline": types.ObjectType{AttrTypes: map[string]attr.Type{"stored_at": types.Int64Type, "type": types.StringType}}, "overcapture_supported": types.BoolType, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "reader": types.StringType, "receipt": types.ObjectType{AttrTypes: map[string]attr.Type{"account_type": types.StringType, "application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "cardholder_verification_method": types.StringType, "dedicated_file_name": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}, "cashapp": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "cashtag": types.StringType, "transaction_id": types.StringType}}, "crypto": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_address": types.StringType, "network": types.StringType, "token_currency": types.StringType, "transaction_hash": types.StringType}}, "eps": types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "verified_name": types.StringType}}, "fpx": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "bank": types.StringType, "transaction_id": types.StringType}}, "giropay": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "verified_name": types.StringType}}, "grabpay": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "ideal": types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "bic": types.StringType, "generated_sepa_debit": types.StringType, "generated_sepa_debit_mandate": types.StringType, "iban_last4": types.StringType, "transaction_id": types.StringType, "verified_name": types.StringType}}, "interac_present": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_card": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "location": types.StringType, "network": types.StringType, "network_transaction_id": types.StringType, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "reader": types.StringType, "receipt": types.ObjectType{AttrTypes: map[string]attr.Type{"account_type": types.StringType, "application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "cardholder_verification_method": types.StringType, "dedicated_file_name": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}}}, "kakao_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"location": types.StringType, "payer_details": types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}}}, "payment_method_category": types.StringType, "preferred_locale": types.StringType, "reader": types.StringType}}, "konbini": types.ObjectType{AttrTypes: map[string]attr.Type{"store": types.ObjectType{AttrTypes: map[string]attr.Type{"chain": types.StringType}}}}, "kr_card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "buyer_id": types.StringType, "last4": types.StringType, "transaction_id": types.StringType}}, "link": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "mobilepay": types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "country": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "last4": types.StringType}}}}, "multibanco": types.ObjectType{AttrTypes: map[string]attr.Type{"entity": types.StringType, "reference": types.StringType}}, "naver_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "nz_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "bank_code": types.StringType, "bank_name": types.StringType, "branch_code": types.StringType, "expected_debit_date": types.StringType, "last4": types.StringType, "suffix": types.StringType}}, "oxxo": types.ObjectType{AttrTypes: map[string]attr.Type{"number": types.StringType}}, "p24": types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "reference": types.StringType, "verified_name": types.StringType}}, "payco": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "paynow": types.ObjectType{AttrTypes: map[string]attr.Type{"location": types.StringType, "reader": types.StringType, "reference": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "payer_email": types.StringType, "payer_id": types.StringType, "payer_name": types.StringType, "seller_protection": types.ObjectType{AttrTypes: map[string]attr.Type{"dispute_categories": types.ListType{ElemType: types.StringType}, "status": types.StringType}}, "transaction_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "last4": types.StringType, "mandate": types.StringType, "pay_id": types.StringType}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transaction_id": types.StringType, "mandate": types.StringType}}, "promptpay": types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType}}, "revolut_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"funding": types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "country": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "funding": types.StringType, "last4": types.StringType}}, "type": types.StringType}}, "transaction_id": types.StringType}}, "samsung_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "satispay": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "scalapay": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "sepa_credit_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "bic": types.StringType, "iban": types.StringType}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "branch_code": types.StringType, "country": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType}}, "sofort": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "country": types.StringType, "generated_sepa_debit": types.StringType, "generated_sepa_debit_mandate": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "verified_name": types.StringType}}, "sunbit": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "swish": types.ObjectType{AttrTypes: map[string]attr.Type{"fingerprint": types.StringType, "payment_reference": types.StringType, "verified_phone_last4": types.StringType}}, "twint": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate": types.StringType}}, "type": types.StringType, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"vpa": types.StringType}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "account_type": types.StringType, "bank_name": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType, "payment_reference": types.StringType, "routing_number": types.StringType}}, "wechat_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"fingerprint": types.StringType, "location": types.StringType, "reader": types.StringType, "transaction_id": types.StringType}}}}, "payment_method_details", "raw response"); err != nil {
					return err
				} else {
					if typedPaymentMethodDetails, ok := valuePaymentMethodDetails.(types.Object); ok {
						state.PaymentMethodDetails = typedPaymentMethodDetails
						assignedPaymentMethodDetails = true
					}
				}
			}
		}
		if !assignedPaymentMethodDetails {
			if !hasRaw {
				if responseValuePaymentMethodDetails, ok := plainFromResponseField(obj, "PaymentMethodDetails"); ok {
					sourcePaymentMethodDetails := applyConfiguredKeyedListShapes(responseValuePaymentMethodDetails, attrValueToPlain(state.PaymentMethodDetails))
					if valuePaymentMethodDetails, err := flattenPlainValue(
						sourcePaymentMethodDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"ach_credit_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"account_number": types.StringType, "bank_name": types.StringType, "routing_number": types.StringType, "swift_code": types.StringType}}, "ach_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "bank_name": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "routing_number": types.StringType}}, "acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "institution_number": types.StringType, "last4": types.StringType, "mandate": types.StringType, "transit_number": types.StringType}}, "affirm": types.ObjectType{AttrTypes: map[string]attr.Type{"location": types.StringType, "reader": types.StringType, "transaction_id": types.StringType}}, "afterpay_clearpay": types.ObjectType{AttrTypes: map[string]attr.Type{"order_id": types.StringType, "reference": types.StringType}}, "alipay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "fingerprint": types.StringType, "transaction_id": types.StringType}}, "alma": types.ObjectType{AttrTypes: map[string]attr.Type{"installments": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type}}, "transaction_id": types.StringType}}, "amazon_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"funding": types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "country": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "funding": types.StringType, "last4": types.StringType}}, "type": types.StringType}}, "transaction_id": types.StringType}}, "au_becs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType, "sort_code": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "generated_sepa_debit": types.StringType, "generated_sepa_debit_mandate": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "verified_name": types.StringType}}, "billie": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "bizum": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "blik": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType}}, "boleto": types.ObjectType{AttrTypes: map[string]attr.Type{"tax_id": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_authorized": types.Int64Type, "authorization_code": types.StringType, "brand": types.StringType, "capture_before": types.Int64Type, "checks": types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_postal_code_check": types.StringType, "cvc_check": types.StringType}}, "country": types.StringType, "description": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "extended_authorization": types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "incremental_authorization": types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}, "installments": types.ObjectType{AttrTypes: map[string]attr.Type{"plan": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}}, "issuer": types.StringType, "last4": types.StringType, "mandate": types.StringType, "moto": types.BoolType, "multicapture": types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}, "network": types.StringType, "network_token": types.ObjectType{AttrTypes: map[string]attr.Type{"used": types.BoolType}}, "network_transaction_id": types.StringType, "overcapture": types.ObjectType{AttrTypes: map[string]attr.Type{"maximum_amount_capturable": types.Int64Type, "status": types.StringType}}, "regulated_status": types.StringType, "three_d_secure": types.ObjectType{AttrTypes: map[string]attr.Type{"authentication_flow": types.StringType, "electronic_commerce_indicator": types.StringType, "exemption_indicator": types.StringType, "exemption_indicator_applied": types.BoolType, "result": types.StringType, "result_reason": types.StringType, "transaction_id": types.StringType, "version": types.StringType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"dynamic_last4": types.StringType, "masterpass": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}, "type": types.StringType, "visa_checkout": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}}}}}, "card_present": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_authorized": types.Int64Type, "brand": types.StringType, "brand_product": types.StringType, "capture_before": types.Int64Type, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_card": types.StringType, "iin": types.StringType, "incremental_authorization_supported": types.BoolType, "issuer": types.StringType, "last4": types.StringType, "location": types.StringType, "network": types.StringType, "network_transaction_id": types.StringType, "offline": types.ObjectType{AttrTypes: map[string]attr.Type{"stored_at": types.Int64Type, "type": types.StringType}}, "overcapture_supported": types.BoolType, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "reader": types.StringType, "receipt": types.ObjectType{AttrTypes: map[string]attr.Type{"account_type": types.StringType, "application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "cardholder_verification_method": types.StringType, "dedicated_file_name": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}, "cashapp": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "cashtag": types.StringType, "transaction_id": types.StringType}}, "crypto": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_address": types.StringType, "network": types.StringType, "token_currency": types.StringType, "transaction_hash": types.StringType}}, "eps": types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "verified_name": types.StringType}}, "fpx": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "bank": types.StringType, "transaction_id": types.StringType}}, "giropay": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "verified_name": types.StringType}}, "grabpay": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "ideal": types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "bic": types.StringType, "generated_sepa_debit": types.StringType, "generated_sepa_debit_mandate": types.StringType, "iban_last4": types.StringType, "transaction_id": types.StringType, "verified_name": types.StringType}}, "interac_present": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_card": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "location": types.StringType, "network": types.StringType, "network_transaction_id": types.StringType, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "reader": types.StringType, "receipt": types.ObjectType{AttrTypes: map[string]attr.Type{"account_type": types.StringType, "application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "cardholder_verification_method": types.StringType, "dedicated_file_name": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}}}, "kakao_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"location": types.StringType, "payer_details": types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}}}, "payment_method_category": types.StringType, "preferred_locale": types.StringType, "reader": types.StringType}}, "konbini": types.ObjectType{AttrTypes: map[string]attr.Type{"store": types.ObjectType{AttrTypes: map[string]attr.Type{"chain": types.StringType}}}}, "kr_card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "buyer_id": types.StringType, "last4": types.StringType, "transaction_id": types.StringType}}, "link": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "mobilepay": types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "country": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "last4": types.StringType}}}}, "multibanco": types.ObjectType{AttrTypes: map[string]attr.Type{"entity": types.StringType, "reference": types.StringType}}, "naver_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "nz_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "bank_code": types.StringType, "bank_name": types.StringType, "branch_code": types.StringType, "expected_debit_date": types.StringType, "last4": types.StringType, "suffix": types.StringType}}, "oxxo": types.ObjectType{AttrTypes: map[string]attr.Type{"number": types.StringType}}, "p24": types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "reference": types.StringType, "verified_name": types.StringType}}, "payco": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "paynow": types.ObjectType{AttrTypes: map[string]attr.Type{"location": types.StringType, "reader": types.StringType, "reference": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "payer_email": types.StringType, "payer_id": types.StringType, "payer_name": types.StringType, "seller_protection": types.ObjectType{AttrTypes: map[string]attr.Type{"dispute_categories": types.ListType{ElemType: types.StringType}, "status": types.StringType}}, "transaction_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "last4": types.StringType, "mandate": types.StringType, "pay_id": types.StringType}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transaction_id": types.StringType, "mandate": types.StringType}}, "promptpay": types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType}}, "revolut_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"funding": types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "country": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "funding": types.StringType, "last4": types.StringType}}, "type": types.StringType}}, "transaction_id": types.StringType}}, "samsung_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "satispay": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "scalapay": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "sepa_credit_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "bic": types.StringType, "iban": types.StringType}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "branch_code": types.StringType, "country": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType}}, "sofort": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "country": types.StringType, "generated_sepa_debit": types.StringType, "generated_sepa_debit_mandate": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "verified_name": types.StringType}}, "sunbit": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "swish": types.ObjectType{AttrTypes: map[string]attr.Type{"fingerprint": types.StringType, "payment_reference": types.StringType, "verified_phone_last4": types.StringType}}, "twint": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate": types.StringType}}, "type": types.StringType, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"vpa": types.StringType}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "account_type": types.StringType, "bank_name": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType, "payment_reference": types.StringType, "routing_number": types.StringType}}, "wechat_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"fingerprint": types.StringType, "location": types.StringType, "reader": types.StringType, "transaction_id": types.StringType}}}},
						"payment_method_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPaymentMethodDetails, ok := valuePaymentMethodDetails.(types.Object); ok {
							state.PaymentMethodDetails = typedPaymentMethodDetails
							assignedPaymentMethodDetails = true
						}
					}
				}
			}
		}
		if !assignedPaymentMethodDetails && hadRawPaymentMethodDetails {
			if nullPaymentMethodDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"ach_credit_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"account_number": types.StringType, "bank_name": types.StringType, "routing_number": types.StringType, "swift_code": types.StringType}}, "ach_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "bank_name": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "routing_number": types.StringType}}, "acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "institution_number": types.StringType, "last4": types.StringType, "mandate": types.StringType, "transit_number": types.StringType}}, "affirm": types.ObjectType{AttrTypes: map[string]attr.Type{"location": types.StringType, "reader": types.StringType, "transaction_id": types.StringType}}, "afterpay_clearpay": types.ObjectType{AttrTypes: map[string]attr.Type{"order_id": types.StringType, "reference": types.StringType}}, "alipay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "fingerprint": types.StringType, "transaction_id": types.StringType}}, "alma": types.ObjectType{AttrTypes: map[string]attr.Type{"installments": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type}}, "transaction_id": types.StringType}}, "amazon_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"funding": types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "country": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "funding": types.StringType, "last4": types.StringType}}, "type": types.StringType}}, "transaction_id": types.StringType}}, "au_becs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType, "sort_code": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "generated_sepa_debit": types.StringType, "generated_sepa_debit_mandate": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "verified_name": types.StringType}}, "billie": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "bizum": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "blik": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType}}, "boleto": types.ObjectType{AttrTypes: map[string]attr.Type{"tax_id": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_authorized": types.Int64Type, "authorization_code": types.StringType, "brand": types.StringType, "capture_before": types.Int64Type, "checks": types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_postal_code_check": types.StringType, "cvc_check": types.StringType}}, "country": types.StringType, "description": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "extended_authorization": types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "incremental_authorization": types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}, "installments": types.ObjectType{AttrTypes: map[string]attr.Type{"plan": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}}, "issuer": types.StringType, "last4": types.StringType, "mandate": types.StringType, "moto": types.BoolType, "multicapture": types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}, "network": types.StringType, "network_token": types.ObjectType{AttrTypes: map[string]attr.Type{"used": types.BoolType}}, "network_transaction_id": types.StringType, "overcapture": types.ObjectType{AttrTypes: map[string]attr.Type{"maximum_amount_capturable": types.Int64Type, "status": types.StringType}}, "regulated_status": types.StringType, "three_d_secure": types.ObjectType{AttrTypes: map[string]attr.Type{"authentication_flow": types.StringType, "electronic_commerce_indicator": types.StringType, "exemption_indicator": types.StringType, "exemption_indicator_applied": types.BoolType, "result": types.StringType, "result_reason": types.StringType, "transaction_id": types.StringType, "version": types.StringType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"dynamic_last4": types.StringType, "masterpass": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}, "type": types.StringType, "visa_checkout": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}}}}}, "card_present": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_authorized": types.Int64Type, "brand": types.StringType, "brand_product": types.StringType, "capture_before": types.Int64Type, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_card": types.StringType, "iin": types.StringType, "incremental_authorization_supported": types.BoolType, "issuer": types.StringType, "last4": types.StringType, "location": types.StringType, "network": types.StringType, "network_transaction_id": types.StringType, "offline": types.ObjectType{AttrTypes: map[string]attr.Type{"stored_at": types.Int64Type, "type": types.StringType}}, "overcapture_supported": types.BoolType, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "reader": types.StringType, "receipt": types.ObjectType{AttrTypes: map[string]attr.Type{"account_type": types.StringType, "application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "cardholder_verification_method": types.StringType, "dedicated_file_name": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}, "cashapp": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "cashtag": types.StringType, "transaction_id": types.StringType}}, "crypto": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_address": types.StringType, "network": types.StringType, "token_currency": types.StringType, "transaction_hash": types.StringType}}, "eps": types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "verified_name": types.StringType}}, "fpx": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "bank": types.StringType, "transaction_id": types.StringType}}, "giropay": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "verified_name": types.StringType}}, "grabpay": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "ideal": types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "bic": types.StringType, "generated_sepa_debit": types.StringType, "generated_sepa_debit_mandate": types.StringType, "iban_last4": types.StringType, "transaction_id": types.StringType, "verified_name": types.StringType}}, "interac_present": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_card": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "location": types.StringType, "network": types.StringType, "network_transaction_id": types.StringType, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "reader": types.StringType, "receipt": types.ObjectType{AttrTypes: map[string]attr.Type{"account_type": types.StringType, "application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "cardholder_verification_method": types.StringType, "dedicated_file_name": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}}}, "kakao_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"location": types.StringType, "payer_details": types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}}}, "payment_method_category": types.StringType, "preferred_locale": types.StringType, "reader": types.StringType}}, "konbini": types.ObjectType{AttrTypes: map[string]attr.Type{"store": types.ObjectType{AttrTypes: map[string]attr.Type{"chain": types.StringType}}}}, "kr_card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "buyer_id": types.StringType, "last4": types.StringType, "transaction_id": types.StringType}}, "link": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "mobilepay": types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "country": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "last4": types.StringType}}}}, "multibanco": types.ObjectType{AttrTypes: map[string]attr.Type{"entity": types.StringType, "reference": types.StringType}}, "naver_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "nz_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "bank_code": types.StringType, "bank_name": types.StringType, "branch_code": types.StringType, "expected_debit_date": types.StringType, "last4": types.StringType, "suffix": types.StringType}}, "oxxo": types.ObjectType{AttrTypes: map[string]attr.Type{"number": types.StringType}}, "p24": types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "reference": types.StringType, "verified_name": types.StringType}}, "payco": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "paynow": types.ObjectType{AttrTypes: map[string]attr.Type{"location": types.StringType, "reader": types.StringType, "reference": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "payer_email": types.StringType, "payer_id": types.StringType, "payer_name": types.StringType, "seller_protection": types.ObjectType{AttrTypes: map[string]attr.Type{"dispute_categories": types.ListType{ElemType: types.StringType}, "status": types.StringType}}, "transaction_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "last4": types.StringType, "mandate": types.StringType, "pay_id": types.StringType}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transaction_id": types.StringType, "mandate": types.StringType}}, "promptpay": types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType}}, "revolut_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"funding": types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "country": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "funding": types.StringType, "last4": types.StringType}}, "type": types.StringType}}, "transaction_id": types.StringType}}, "samsung_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "transaction_id": types.StringType}}, "satispay": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "scalapay": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "sepa_credit_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "bic": types.StringType, "iban": types.StringType}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "branch_code": types.StringType, "country": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType}}, "sofort": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "country": types.StringType, "generated_sepa_debit": types.StringType, "generated_sepa_debit_mandate": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "verified_name": types.StringType}}, "sunbit": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_id": types.StringType}}, "swish": types.ObjectType{AttrTypes: map[string]attr.Type{"fingerprint": types.StringType, "payment_reference": types.StringType, "verified_phone_last4": types.StringType}}, "twint": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate": types.StringType}}, "type": types.StringType, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"vpa": types.StringType}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "account_type": types.StringType, "bank_name": types.StringType, "expected_debit_date": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate": types.StringType, "payment_reference": types.StringType, "routing_number": types.StringType}}, "wechat_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"fingerprint": types.StringType, "location": types.StringType, "reader": types.StringType, "transaction_id": types.StringType}}}}); ok {
				if typedPaymentMethodDetails, ok := nullPaymentMethodDetails.(types.Object); ok {
					state.PaymentMethodDetails = typedPaymentMethodDetails
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
		assignedRadarOptions := false
		hadRawRadarOptions := false
		if rawValueRadarOptions, rawOk := plainValueAtPath(raw, "radar_options"); rawOk {
			hadRawRadarOptions = true
			if rawValueRadarOptions != nil {
				sourceRadarOptions := applyConfiguredKeyedListShapes(rawValueRadarOptions, attrValueToPlain(state.RadarOptions))
				if valueRadarOptions, err := flattenPlainValue(sourceRadarOptions, types.ObjectType{AttrTypes: map[string]attr.Type{"session": types.StringType}}, "radar_options", "raw response"); err != nil {
					return err
				} else {
					if typedRadarOptions, ok := valueRadarOptions.(types.Object); ok {
						state.RadarOptions = typedRadarOptions
						assignedRadarOptions = true
					}
				}
			}
		}
		if !assignedRadarOptions {
			if !hasRaw {
				if responseValueRadarOptions, ok := plainFromResponseField(obj, "RadarOptions"); ok {
					sourceRadarOptions := applyConfiguredKeyedListShapes(responseValueRadarOptions, attrValueToPlain(state.RadarOptions))
					if valueRadarOptions, err := flattenPlainValue(
						sourceRadarOptions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"session": types.StringType}},
						"radar_options",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRadarOptions, ok := valueRadarOptions.(types.Object); ok {
							state.RadarOptions = typedRadarOptions
							assignedRadarOptions = true
						}
					}
				}
			}
		}
		if !assignedRadarOptions && hadRawRadarOptions {
			if nullRadarOptions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"session": types.StringType}}); ok {
				if typedRadarOptions, ok := nullRadarOptions.(types.Object); ok {
					state.RadarOptions = typedRadarOptions
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
		if rawValueReceiptNumber, rawOk := plainValueAtPath(raw, "receipt_number"); rawOk {
			if valueReceiptNumber, err := flattenPlainValue(rawValueReceiptNumber, types.StringType, "receipt_number", "raw response"); err != nil {
				return err
			} else {
				if typedReceiptNumber, ok := valueReceiptNumber.(types.String); ok {
					state.ReceiptNumber = typedReceiptNumber
				}
			}
		} else if !hasRaw {
			if responseValueReceiptNumber, ok := plainFromResponseField(obj, "ReceiptNumber"); ok {
				if valueReceiptNumber, err := flattenPlainValue(responseValueReceiptNumber, types.StringType, "receipt_number", "response struct"); err != nil {
					return err
				} else {
					if typedReceiptNumber, ok := valueReceiptNumber.(types.String); ok {
						state.ReceiptNumber = typedReceiptNumber
					}
				}
			}
		}
	}
	{
		if rawValueRefunded, rawOk := plainValueAtPath(raw, "refunded"); rawOk {
			if valueRefunded, err := flattenPlainValue(rawValueRefunded, types.BoolType, "refunded", "raw response"); err != nil {
				return err
			} else {
				if typedRefunded, ok := valueRefunded.(types.Bool); ok {
					state.Refunded = typedRefunded
				}
			}
		} else if !hasRaw {
			if responseValueRefunded, ok := plainFromResponseField(obj, "Refunded"); ok {
				if valueRefunded, err := flattenPlainValue(responseValueRefunded, types.BoolType, "refunded", "response struct"); err != nil {
					return err
				} else {
					if typedRefunded, ok := valueRefunded.(types.Bool); ok {
						state.Refunded = typedRefunded
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
		if state.Source.IsNull() || state.Source.IsUnknown() {
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
		if true {
			if rawValueSourceTransfer, rawOk := plainValueAtPath(raw, "source_transfer"); rawOk {
				if typedSourceTransfer, ok := plainToStringIDValue(rawValueSourceTransfer); ok {
					state.SourceTransfer = typedSourceTransfer
				}
			} else if !hasRaw {
				if responseValueSourceTransfer, ok := plainFromResponseField(obj, "SourceTransfer"); ok {
					if typedSourceTransfer, ok := plainToStringIDValue(responseValueSourceTransfer); ok {
						state.SourceTransfer = typedSourceTransfer
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
		if true {
			if rawValueTransfer, rawOk := plainValueAtPath(raw, "transfer"); rawOk {
				if typedTransfer, ok := plainToStringIDValue(rawValueTransfer); ok {
					state.Transfer = typedTransfer
				}
			} else if !hasRaw {
				if responseValueTransfer, ok := plainFromResponseField(obj, "Transfer"); ok {
					if typedTransfer, ok := plainToStringIDValue(responseValueTransfer); ok {
						state.Transfer = typedTransfer
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
				if valueTransferData, err := flattenPlainValue(sourceTransferData, types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "destination": types.StringType, "description": types.StringType}}, "transfer_data", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "destination": types.StringType, "description": types.StringType}},
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
			if nullTransferData, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "destination": types.StringType, "description": types.StringType}}); ok {
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
