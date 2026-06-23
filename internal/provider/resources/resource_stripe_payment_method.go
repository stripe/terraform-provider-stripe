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

var _ resource.Resource = &PaymentMethodResource{}

var _ resource.ResourceWithConfigure = &PaymentMethodResource{}

var _ resource.ResourceWithImportState = &PaymentMethodResource{}

func NewPaymentMethodResource() resource.Resource {
	return &PaymentMethodResource{}
}

type PaymentMethodResource struct {
	client *stripe.Client
}

type PaymentMethodResourceModel struct {
	Object         types.String `tfsdk:"object"`
	ACSSDebit      types.Object `tfsdk:"acss_debit"`
	AllowRedisplay types.String `tfsdk:"allow_redisplay"`
	AUBECSDebit    types.Object `tfsdk:"au_becs_debit"`
	BACSDebit      types.Object `tfsdk:"bacs_debit"`
	BillingDetails types.Object `tfsdk:"billing_details"`
	Boleto         types.Object `tfsdk:"boleto"`
	Card           types.Object `tfsdk:"card"`
	CardPresent    types.Object `tfsdk:"card_present"`
	CashApp        types.Object `tfsdk:"cashapp"`
	Created        types.Int64  `tfsdk:"created"`
	Customer       types.String `tfsdk:"customer"`
	EPS            types.Object `tfsdk:"eps"`
	FPX            types.Object `tfsdk:"fpx"`
	ID             types.String `tfsdk:"id"`
	IDEAL          types.Object `tfsdk:"ideal"`
	InteracPresent types.Object `tfsdk:"interac_present"`
	Klarna         types.Object `tfsdk:"klarna"`
	KrCard         types.Object `tfsdk:"kr_card"`
	Link           types.Object `tfsdk:"link"`
	Livemode       types.Bool   `tfsdk:"livemode"`
	Metadata       types.Map    `tfsdk:"metadata"`
	NaverPay       types.Object `tfsdk:"naver_pay"`
	NzBankAccount  types.Object `tfsdk:"nz_bank_account"`
	P24            types.Object `tfsdk:"p24"`
	Paypal         types.Object `tfsdk:"paypal"`
	Payto          types.Object `tfsdk:"payto"`
	RadarOptions   types.Object `tfsdk:"radar_options"`
	SEPADebit      types.Object `tfsdk:"sepa_debit"`
	Sofort         types.Object `tfsdk:"sofort"`
	Type           types.String `tfsdk:"type"`
	Upi            types.Object `tfsdk:"upi"`
	USBankAccount  types.Object `tfsdk:"us_bank_account"`
	PaymentMethod  types.String `tfsdk:"payment_method"`
}

func (r *PaymentMethodResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PaymentMethodResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_payment_method"
}

func (r *PaymentMethodResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "PaymentMethod objects represent your customer's payment instruments.\nYou can use them with [PaymentIntents](https://docs.stripe.com/payments/payment-intents) to collect payments or save them to\nCustomer objects to store instrument details for future payments.\n\nRelated guides: [Payment Methods](https://docs.stripe.com/payments/payment-methods) and [More Payment Scenarios](https://docs.stripe.com/payments/more-payment-scenarios).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("payment_method")},
			},
			"acss_debit": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"bank_name": schema.StringAttribute{
						Computed:      true,
						Description:   "Name of the bank associated with the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"fingerprint": schema.StringAttribute{
						Computed:      true,
						Description:   "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"institution_number": schema.StringAttribute{
						Required:      true,
						Description:   "Institution number of the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"last4": schema.StringAttribute{
						Computed:      true,
						Description:   "Last four digits of the bank account number.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"transit_number": schema.StringAttribute{
						Required:      true,
						Description:   "Transit number of the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"account_number": schema.StringAttribute{
						Required:      true,
						Description:   "Customer's bank account number.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"allow_redisplay": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow. The field defaults to “unspecified”.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("always", "limited", "unspecified")},
			},
			"au_becs_debit": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"bsb_number": schema.StringAttribute{
						Required:      true,
						Description:   "Six-digit number identifying bank and branch associated with this bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
					"account_number": schema.StringAttribute{
						Required:      true,
						Description:   "The account number for the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"bacs_debit": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
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
					"sort_code": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Sort code of the bank account. (e.g., `10-20-30`)",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"account_number": schema.StringAttribute{
						Optional:      true,
						Description:   "Account number of the bank account that the funds will be debited from.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"billing_details": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"address": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Billing address.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
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
					"email": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Email address.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"name": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Full name.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"phone": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Billing phone number (including extension).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"tax_id": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"boleto": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"tax_id": schema.StringAttribute{
						Required:      true,
						Description:   "Uniquely identifies the customer tax id (CNPJ or CPF)",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"card": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"brand": schema.StringAttribute{
						Computed:      true,
						Description:   "Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"checks": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Checks on Card address and CVC if provided.",
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
					"display_brand": schema.StringAttribute{
						Computed:      true,
						Description:   "The brand to use when displaying the card, this accounts for customer's brand choice on dual-branded cards. Can be `american_express`, `cartes_bancaires`, `diners_club`, `discover`, `eftpos_australia`, `interac`, `jcb`, `mastercard`, `union_pay`, `visa`, or `other` and may contain more values in the future.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"exp_month": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "Two-digit number representing the card's expiration month.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"exp_year": schema.Int64Attribute{
						Optional:      true,
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
					"generated_from": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details of the original PaymentMethod that created this object.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"charge": schema.StringAttribute{
								Computed:      true,
								Description:   "The charge that created this object.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"payment_method_details": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Transaction-specific details of the payment method used in the payment.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
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
									"type": schema.StringAttribute{
										Computed:      true,
										Description:   "The type of payment method transaction-specific details from the transaction that generated this `card` payment method. Always `card_present`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"setup_attempt": schema.StringAttribute{
								Computed:      true,
								Description:   "The ID of the SetupAttempt that generated this PaymentMethod, if any.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
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
					"networks": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Contains information about card networks that can be used to process the payment.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"available": schema.ListAttribute{
								Computed:      true,
								Description:   "All networks available for selection via [payment_method_options.card.network](/api/payment_intents/confirm#confirm_payment_intent-payment_method_options-card-network).",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"preferred": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The preferred network for co-branded cards. Can be `cartes_bancaires`, `mastercard`, `visa` or `invalid_preference` if requested network is not valid for the card.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"regulated_status": schema.StringAttribute{
						Computed:      true,
						Description:   "Status of a card based on the card issuer.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("regulated", "unregulated")},
					},
					"three_d_secure_usage": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains details on how this Card may be used for 3D Secure authentication.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"supported": schema.BoolAttribute{
								Computed:      true,
								Description:   "Whether 3D Secure is supported on this card.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
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
					"cvc": schema.StringAttribute{
						Optional:      true,
						Description:   "The card's CVC. It is highly recommended to always include this value.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"number": schema.StringAttribute{
						Optional:      true,
						Description:   "The card number, as a string without any separators.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"token": schema.StringAttribute{
						Optional:      true,
						Description:   "For backwards compatibility, you can alternatively provide a Stripe token (e.g., for Apple Pay, Amex Express Checkout, or legacy Checkout) into the card hash with format card: {token: \"tok_visa\"}.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"card_present": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
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
					"networks": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains information about card networks that can be used to process the payment.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"available": schema.ListAttribute{
								Computed:      true,
								Description:   "All networks available for selection via [payment_method_options.card.network](/api/payment_intents/confirm#confirm_payment_intent-payment_method_options-card-network).",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"preferred": schema.StringAttribute{
								Computed:      true,
								Description:   "The preferred network for the card.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"offline": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details about payment methods collected offline.",
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
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
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
				},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ID of the Customer to which this PaymentMethod is saved. This will not be set when the PaymentMethod has not been saved to a Customer.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"eps": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"bank": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `deutsche_bank_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("arzte_und_apotheker_bank", "austrian_anadi_bank_ag", "bank_austria", "bankhaus_carl_spangler", "bankhaus_schelhammer_und_schattera_ag", "bawag_psk_ag", "bks_bank_ag", "brull_kallmus_bank_ag", "btv_vier_lander_bank", "capital_bank_grawe_gruppe_ag", "deutsche_bank_ag", "dolomitenbank", "easybank_ag", "erste_bank_und_sparkassen", "hypo_alpeadriabank_international_ag", "hypo_bank_burgenland_aktiengesellschaft", "hypo_noe_lb_fur_niederosterreich_u_wien", "hypo_oberosterreich_salzburg_steiermark", "hypo_tirol_bank_ag", "hypo_vorarlberg_bank_ag", "marchfelder_bank", "oberbank_ag", "raiffeisen_bankengruppe_osterreich", "schoellerbank_ag", "sparda_bank_wien", "volksbank_gruppe", "volkskreditbank_ag", "vr_bank_braunau")},
					},
				},
			},
			"fpx": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"account_holder_type": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Account holder type, if provided. Can be one of `individual` or `company`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("company", "individual")},
					},
					"bank": schema.StringAttribute{
						Required:      true,
						Description:   "The customer's bank, if provided. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, `pb_enterprise`, or `bank_of_china`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("affin_bank", "agrobank", "alliance_bank", "ambank", "bank_islam", "bank_muamalat", "bank_of_china", "bank_rakyat", "bsn", "cimb", "deutsche_bank", "hong_leong_bank", "hsbc", "kfh", "maybank2e", "maybank2u", "ocbc", "pb_enterprise", "public_bank", "rhb", "standard_chartered", "uob")},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ideal": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"bank": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The customer's bank, if provided. Can be one of `abn_amro`, `adyen`, `asn_bank`, `bunq`, `buut`, `finom`, `handelsbanken`, `ing`, `knab`, `mollie`, `moneyou`, `n26`, `nn`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, `van_lanschot`, or `yoursafe`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("abn_amro", "adyen", "asn_bank", "bunq", "buut", "finom", "handelsbanken", "ing", "knab", "mollie", "moneyou", "n26", "nn", "rabobank", "regiobank", "revolut", "sns_bank", "triodos_bank", "van_lanschot", "yoursafe")},
					},
					"bic": schema.StringAttribute{
						Computed:      true,
						Description:   "The Bank Identifier Code of the customer's bank, if the bank was provided.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("ABNANL2A", "ADYBNL2A", "ASNBNL21", "BITSNL2A", "BUNQNL2A", "BUUTNL2A", "FNOMNL22", "FVLBNL22", "HANDNL2A", "INGBNL2A", "KNABNL2H", "MLLENL2A", "MOYONL21", "NNBANL2G", "NTSBDEB1", "RABONL2U", "RBRBNL21", "REVOIE23", "REVOLT21", "SNSBNL2A", "TRIONL2U")},
					},
				},
			},
			"interac_present": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
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
					"networks": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains information about card networks that can be used to process the payment.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"available": schema.ListAttribute{
								Computed:      true,
								Description:   "All networks available for selection via [payment_method_options.card.network](/api/payment_intents/confirm#confirm_payment_intent-payment_method_options-card-network).",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"preferred": schema.StringAttribute{
								Computed:      true,
								Description:   "The preferred network for the card.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
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
				},
			},
			"klarna": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"dob": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The customer's date of birth, if provided.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"day": schema.Int64Attribute{
								Required:      true,
								Description:   "The day of birth, between 1 and 31.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
							"month": schema.Int64Attribute{
								Required:      true,
								Description:   "The month of birth, between 1 and 12.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
							"year": schema.Int64Attribute{
								Required:      true,
								Description:   "The four-digit year of birth.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
						},
					},
				},
			},
			"kr_card": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"brand": schema.StringAttribute{
						Computed:      true,
						Description:   "The local credit or debit card brand.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("bc", "citi", "hana", "hyundai", "jeju", "jeonbuk", "kakaobank", "kbank", "kdbbank", "kookmin", "kwangju", "lotte", "mg", "nh", "post", "samsung", "savingsbank", "shinhan", "shinhyup", "suhyup", "tossbank", "woori")},
					},
					"last4": schema.StringAttribute{
						Computed:      true,
						Description:   "The last four digits of the card. This may not be present for American Express cards.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"link": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"email": schema.StringAttribute{
						Computed:      true,
						Description:   "Account owner's email address.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"persistent_token": schema.StringAttribute{
						Computed:      true,
						Description:   "[Deprecated] This is a legacy parameter that no longer has any function.",
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
			"naver_pay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"buyer_id": schema.StringAttribute{
						Computed:      true,
						Description:   "Uniquely identifies this particular Naver Pay account. You can use this attribute to check whether two Naver Pay accounts are the same.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"funding": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Whether to fund this transaction with Naver Pay points or a card.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("card", "points")},
					},
				},
			},
			"nz_bank_account": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"account_holder_name": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The name on the bank account. Only present if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod’s billing details.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"bank_code": schema.StringAttribute{
						Required:      true,
						Description:   "The numeric code for the bank account's bank.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"bank_name": schema.StringAttribute{
						Computed:      true,
						Description:   "The name of the bank.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"branch_code": schema.StringAttribute{
						Required:      true,
						Description:   "The numeric code for the bank account's bank branch.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"last4": schema.StringAttribute{
						Computed:      true,
						Description:   "Last four digits of the bank account number.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"suffix": schema.StringAttribute{
						Required:      true,
						Description:   "The suffix of the bank account number.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"account_number": schema.StringAttribute{
						Required:      true,
						Description:   "The account number for the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"reference": schema.StringAttribute{
						Optional: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"p24": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"bank": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The customer's bank, if provided.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("alior_bank", "bank_millennium", "bank_nowy_bfg_sa", "bank_pekao_sa", "banki_spbdzielcze", "blik", "bnp_paribas", "boz", "citi_handlowy", "credit_agricole", "envelobank", "etransfer_pocztowy24", "getin_bank", "ideabank", "ing", "inteligo", "mbank_mtransfer", "nest_przelew", "noble_pay", "pbac_z_ipko", "plus_bank", "santander_przelew24", "tmobile_usbugi_bankowe", "toyota_bank", "velobank", "volkswagen_bank")},
					},
				},
			},
			"paypal": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
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
				},
			},
			"payto": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bsb_number": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Bank-State-Branch number of the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"last4": schema.StringAttribute{
						Computed:      true,
						Description:   "Last four digits of the bank account number.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_id": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The PayID alias for the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"account_number": schema.StringAttribute{
						Optional:    true,
						Description: "The account number for the bank account.",
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
			"sepa_debit": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
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
					"fingerprint": schema.StringAttribute{
						Computed:      true,
						Description:   "Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"generated_from": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Information about the object that generated this PaymentMethod.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"charge": schema.StringAttribute{
								Computed:      true,
								Description:   "The ID of the Charge that generated this PaymentMethod, if any.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"setup_attempt": schema.StringAttribute{
								Computed:      true,
								Description:   "The ID of the SetupAttempt that generated this PaymentMethod, if any.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"last4": schema.StringAttribute{
						Computed:      true,
						Description:   "Last four characters of the IBAN.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"iban": schema.StringAttribute{
						Required:      true,
						Description:   "IBAN of the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"sofort": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"country": schema.StringAttribute{
						Required:      true,
						Description:   "Two-letter ISO code representing the country the bank account is located in.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("acss_debit", "affirm", "afterpay_clearpay", "alipay", "alma", "amazon_pay", "au_becs_debit", "bacs_debit", "bancontact", "billie", "bizum", "blik", "boleto", "card", "card_present", "cashapp", "crypto", "custom", "customer_balance", "eps", "fpx", "giropay", "grabpay", "ideal", "interac_present", "kakao_pay", "klarna", "konbini", "kr_card", "link", "mb_way", "mobilepay", "multibanco", "naver_pay", "nz_bank_account", "oxxo", "p24", "pay_by_bank", "payco", "paynow", "paypal", "payto", "pix", "promptpay", "revolut_pay", "samsung_pay", "satispay", "scalapay", "sepa_debit", "sofort", "sunbit", "swish", "twint", "upi", "us_bank_account", "wechat_pay", "zip")},
			},
			"upi": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"vpa": schema.StringAttribute{
						Computed:      true,
						Description:   "Customer's unique Virtual Payment Address",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"mandate_options": schema.SingleNestedAttribute{
						Optional:      true,
						Description:   "Configuration options for setting up an eMandate",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"amount": schema.Int64Attribute{
								Optional:      true,
								Description:   "Amount to be charged for future payments.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
							"amount_type": schema.StringAttribute{
								Optional:      true,
								Description:   "One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
							"description": schema.StringAttribute{
								Optional:      true,
								Description:   "A description of the mandate or subscription that is meant to be displayed to the customer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
							"end_date": schema.Int64Attribute{
								Optional:      true,
								Description:   "End date of the mandate or subscription.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
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
					"account_holder_type": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Account holder type: individual or company.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("company", "individual")},
					},
					"account_type": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Account type: checkings or savings. Defaults to checking if omitted.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("checking", "savings")},
					},
					"bank_name": schema.StringAttribute{
						Computed:      true,
						Description:   "The name of the bank.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"financial_connections_account": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The ID of the Financial Connections Account used to create the payment method.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
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
					"networks": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains information about US bank account networks that can be used.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"preferred": schema.StringAttribute{
								Computed:      true,
								Description:   "The preferred network.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"supported": schema.ListAttribute{
								Computed:      true,
								Description:   "All supported networks.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
						},
					},
					"routing_number": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Routing number of the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"status_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains information about the future reusability of this PaymentMethod.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"blocked": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"network_code": schema.StringAttribute{
										Computed:      true,
										Description:   "The ACH network code that resulted in this block.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("R02", "R03", "R04", "R05", "R07", "R08", "R10", "R11", "R16", "R20", "R29", "R31")},
									},
									"reason": schema.StringAttribute{
										Computed:      true,
										Description:   "The reason why this PaymentMethod's fingerprint has been blocked",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("bank_account_closed", "bank_account_frozen", "bank_account_invalid_details", "bank_account_restricted", "bank_account_unusable", "debit_not_authorized", "tokenized_account_number_deactivated")},
									},
								},
							},
						},
					},
					"account_number": schema.StringAttribute{
						Optional:      true,
						Description:   "Account number of the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"payment_method": schema.StringAttribute{
				Optional:      true,
				Description:   "The PaymentMethod to share.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
		},
	}
}

func (r *PaymentMethodResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PaymentMethodResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PaymentMethodResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PaymentMethod"}})

	params, err := expandPaymentMethodCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentMethod create params", err.Error())
		return
	}

	obj, err := r.client.V1PaymentMethods.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating PaymentMethod", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentMethods.B, r.client.V1PaymentMethods.Key, stripe.FormatURLPath("/v1/payment_methods/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentMethod create raw response", err.Error())
		return
	}

	if err := flattenPaymentMethod(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentMethod create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"ACSSDebit", "account_number"}, []string{"AUBECSDebit", "account_number"}, []string{"BACSDebit", "account_number"}, []string{"Card", "cvc"}, []string{"Card", "number"}, []string{"Card", "token"}, []string{"NzBankAccount", "account_number"}, []string{"NzBankAccount", "reference"}, []string{"Payto", "account_number"}, []string{"SEPADebit", "iban"}, []string{"Upi", "mandate_options"}, []string{"Upi", "mandate_options", "amount"}, []string{"Upi", "mandate_options", "amount_type"}, []string{"Upi", "mandate_options", "description"}, []string{"Upi", "mandate_options", "end_date"}, []string{"USBankAccount", "account_number"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"PaymentMethod"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentMethodResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState PaymentMethodResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PaymentMethodResourceModel
	state = priorState

	obj, err := r.client.V1PaymentMethods.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading PaymentMethod", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentMethods.B, r.client.V1PaymentMethods.Key, stripe.FormatURLPath("/v1/payment_methods/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentMethod raw response", err.Error())
		return
	}

	if err := flattenPaymentMethod(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentMethod read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"ACSSDebit", "account_number"}, []string{"AUBECSDebit", "account_number"}, []string{"BACSDebit", "account_number"}, []string{"Card", "cvc"}, []string{"Card", "number"}, []string{"Card", "token"}, []string{"NzBankAccount", "account_number"}, []string{"NzBankAccount", "reference"}, []string{"Payto", "account_number"}, []string{"SEPADebit", "iban"}, []string{"Upi", "mandate_options"}, []string{"Upi", "mandate_options", "amount"}, []string{"Upi", "mandate_options", "amount_type"}, []string{"Upi", "mandate_options", "description"}, []string{"Upi", "mandate_options", "end_date"}, []string{"USBankAccount", "account_number"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"PaymentMethod"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentMethodResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PaymentMethodResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PaymentMethodResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PaymentMethod"}})

	var state PaymentMethodResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"PaymentMethod"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"PaymentMethod"}})

	params, err := expandPaymentMethodUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentMethod update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building PaymentMethod update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1PaymentMethods.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating PaymentMethod", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentMethods.B, r.client.V1PaymentMethods.Key, stripe.FormatURLPath("/v1/payment_methods/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentMethod update raw response", err.Error())
		return
	}

	if err := flattenPaymentMethod(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentMethod update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"ACSSDebit", "account_number"}, []string{"AUBECSDebit", "account_number"}, []string{"BACSDebit", "account_number"}, []string{"Card", "cvc"}, []string{"Card", "number"}, []string{"Card", "token"}, []string{"NzBankAccount", "account_number"}, []string{"NzBankAccount", "reference"}, []string{"Payto", "account_number"}, []string{"SEPADebit", "iban"}, []string{"Upi", "mandate_options"}, []string{"Upi", "mandate_options", "amount"}, []string{"Upi", "mandate_options", "amount_type"}, []string{"Upi", "mandate_options", "description"}, []string{"Upi", "mandate_options", "end_date"}, []string{"USBankAccount", "account_number"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"PaymentMethod"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentMethodResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *PaymentMethodResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandPaymentMethodCreate(plan PaymentMethodResourceModel) (*stripe.PaymentMethodCreateParams, error) {
	params := &stripe.PaymentMethodCreateParams{}

	if !plan.ACSSDebit.IsNull() && !plan.ACSSDebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ACSSDebit", plan.ACSSDebit) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "acss_debit", params)
		}
	}
	if !plan.AllowRedisplay.IsNull() && !plan.AllowRedisplay.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "AllowRedisplay", "AllowRedisplay", plan.AllowRedisplay.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "allow_redisplay", params)
		}
	}
	if !plan.AUBECSDebit.IsNull() && !plan.AUBECSDebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AUBECSDebit", plan.AUBECSDebit) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "au_becs_debit", params)
		}
	}
	if !plan.BACSDebit.IsNull() && !plan.BACSDebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BACSDebit", plan.BACSDebit) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "bacs_debit", params)
		}
	}
	if !plan.BillingDetails.IsNull() && !plan.BillingDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingDetails", plan.BillingDetails) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_details", params)
		}
	}
	if !plan.Boleto.IsNull() && !plan.Boleto.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Boleto", plan.Boleto) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "boleto", params)
		}
	}
	if !plan.Card.IsNull() && !plan.Card.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Card", plan.Card) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "card", params)
		}
	}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerID", "Customer", plan.Customer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer", params)
		}
	}
	if !plan.EPS.IsNull() && !plan.EPS.IsUnknown() {
		if !assignAttrValueToNamedField(params, "EPS", plan.EPS) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "eps", params)
		}
	}
	if !plan.FPX.IsNull() && !plan.FPX.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FPX", plan.FPX) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "fpx", params)
		}
	}
	if !plan.IDEAL.IsNull() && !plan.IDEAL.IsUnknown() {
		if !assignAttrValueToNamedField(params, "IDEAL", plan.IDEAL) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "ideal", params)
		}
	}
	if !plan.Klarna.IsNull() && !plan.Klarna.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Klarna", plan.Klarna) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "klarna", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.NaverPay.IsNull() && !plan.NaverPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "NaverPay", plan.NaverPay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "naver_pay", params)
		}
	}
	if !plan.NzBankAccount.IsNull() && !plan.NzBankAccount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "NzBankAccount", plan.NzBankAccount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "nz_bank_account", params)
		}
	}
	if !plan.P24.IsNull() && !plan.P24.IsUnknown() {
		if !assignAttrValueToNamedField(params, "P24", plan.P24) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "p24", params)
		}
	}
	if !plan.Payto.IsNull() && !plan.Payto.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Payto", plan.Payto) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payto", params)
		}
	}
	if !plan.RadarOptions.IsNull() && !plan.RadarOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RadarOptions", plan.RadarOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "radar_options", params)
		}
	}
	if !plan.SEPADebit.IsNull() && !plan.SEPADebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SEPADebit", plan.SEPADebit) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "sepa_debit", params)
		}
	}
	if !plan.Sofort.IsNull() && !plan.Sofort.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Sofort", plan.Sofort) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "sofort", params)
		}
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}
	if !plan.Upi.IsNull() && !plan.Upi.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Upi", plan.Upi) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "upi", params)
		}
	}
	if !plan.USBankAccount.IsNull() && !plan.USBankAccount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "USBankAccount", plan.USBankAccount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "us_bank_account", params)
		}
	}
	if !plan.PaymentMethod.IsNull() && !plan.PaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethod", "PaymentMethod", plan.PaymentMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method", params)
		}
	}

	return params, nil
}

func expandPaymentMethodUpdate(plan PaymentMethodResourceModel, state PaymentMethodResourceModel) (*stripe.PaymentMethodUpdateParams, error) {
	params := &stripe.PaymentMethodUpdateParams{}

	if !plan.AllowRedisplay.Equal(state.AllowRedisplay) && !plan.AllowRedisplay.IsNull() && !plan.AllowRedisplay.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "AllowRedisplay", "AllowRedisplay", plan.AllowRedisplay.ValueString()) {
			if !plan.AllowRedisplay.Equal(state.AllowRedisplay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "allow_redisplay", params)
			}
		}
	}
	if !plan.BillingDetails.Equal(state.BillingDetails) && !plan.BillingDetails.IsNull() && !plan.BillingDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingDetails", plan.BillingDetails) {
			if !plan.BillingDetails.Equal(state.BillingDetails) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "billing_details", params)
			}
		}
	}
	if !plan.Card.Equal(state.Card) && !plan.Card.IsNull() && !plan.Card.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Card", plan.Card) {
			if !plan.Card.Equal(state.Card) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "card", params)
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
	if !plan.Payto.Equal(state.Payto) && !plan.Payto.IsNull() && !plan.Payto.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Payto", plan.Payto) {
			if !plan.Payto.Equal(state.Payto) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payto", params)
			}
		}
	}
	if !plan.USBankAccount.Equal(state.USBankAccount) && !plan.USBankAccount.IsNull() && !plan.USBankAccount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "USBankAccount", plan.USBankAccount) {
			if !plan.USBankAccount.Equal(state.USBankAccount) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "us_bank_account", params)
			}
		}
	}

	return params, nil
}

func flattenPaymentMethod(obj *stripe.PaymentMethod, state *PaymentMethodResourceModel) error {
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
		assignedACSSDebit := false
		hadRawACSSDebit := false
		if rawValueACSSDebit, rawOk := plainValueAtPath(raw, "acss_debit"); rawOk {
			hadRawACSSDebit = true
			if rawValueACSSDebit != nil {
				sourceACSSDebit := applyConfiguredKeyedListShapes(rawValueACSSDebit, attrValueToPlain(state.ACSSDebit))
				if valueACSSDebit, err := flattenPlainValue(sourceACSSDebit, types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "fingerprint": types.StringType, "institution_number": types.StringType, "last4": types.StringType, "transit_number": types.StringType, "account_number": types.StringType}}, "acss_debit", "raw response"); err != nil {
					return err
				} else {
					if typedACSSDebit, ok := valueACSSDebit.(types.Object); ok {
						state.ACSSDebit = typedACSSDebit
						assignedACSSDebit = true
					}
				}
			}
		}
		if !assignedACSSDebit {
			if !hasRaw {
				if responseValueACSSDebit, ok := plainFromResponseField(obj, "ACSSDebit"); ok {
					sourceACSSDebit := applyConfiguredKeyedListShapes(responseValueACSSDebit, attrValueToPlain(state.ACSSDebit))
					if valueACSSDebit, err := flattenPlainValue(
						sourceACSSDebit,
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "fingerprint": types.StringType, "institution_number": types.StringType, "last4": types.StringType, "transit_number": types.StringType, "account_number": types.StringType}},
						"acss_debit",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedACSSDebit, ok := valueACSSDebit.(types.Object); ok {
							state.ACSSDebit = typedACSSDebit
							assignedACSSDebit = true
						}
					}
				}
			}
		}
		if !assignedACSSDebit && hadRawACSSDebit {
			if nullACSSDebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "fingerprint": types.StringType, "institution_number": types.StringType, "last4": types.StringType, "transit_number": types.StringType, "account_number": types.StringType}}); ok {
				if typedACSSDebit, ok := nullACSSDebit.(types.Object); ok {
					state.ACSSDebit = typedACSSDebit
				}
			}
		}
	}
	{
		if rawValueAllowRedisplay, rawOk := plainValueAtPath(raw, "allow_redisplay"); rawOk {
			if valueAllowRedisplay, err := flattenPlainValue(rawValueAllowRedisplay, types.StringType, "allow_redisplay", "raw response"); err != nil {
				return err
			} else {
				if typedAllowRedisplay, ok := valueAllowRedisplay.(types.String); ok {
					state.AllowRedisplay = typedAllowRedisplay
				}
			}
		} else if !hasRaw {
			if responseValueAllowRedisplay, ok := plainFromResponseField(obj, "AllowRedisplay"); ok {
				if valueAllowRedisplay, err := flattenPlainValue(responseValueAllowRedisplay, types.StringType, "allow_redisplay", "response struct"); err != nil {
					return err
				} else {
					if typedAllowRedisplay, ok := valueAllowRedisplay.(types.String); ok {
						state.AllowRedisplay = typedAllowRedisplay
					}
				}
			}
		}
	}
	{
		assignedAUBECSDebit := false
		hadRawAUBECSDebit := false
		if rawValueAUBECSDebit, rawOk := plainValueAtPath(raw, "au_becs_debit"); rawOk {
			hadRawAUBECSDebit = true
			if rawValueAUBECSDebit != nil {
				sourceAUBECSDebit := applyConfiguredKeyedListShapes(rawValueAUBECSDebit, attrValueToPlain(state.AUBECSDebit))
				if valueAUBECSDebit, err := flattenPlainValue(sourceAUBECSDebit, types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "account_number": types.StringType}}, "au_becs_debit", "raw response"); err != nil {
					return err
				} else {
					if typedAUBECSDebit, ok := valueAUBECSDebit.(types.Object); ok {
						state.AUBECSDebit = typedAUBECSDebit
						assignedAUBECSDebit = true
					}
				}
			}
		}
		if !assignedAUBECSDebit {
			if !hasRaw {
				if responseValueAUBECSDebit, ok := plainFromResponseField(obj, "AUBECSDebit"); ok {
					sourceAUBECSDebit := applyConfiguredKeyedListShapes(responseValueAUBECSDebit, attrValueToPlain(state.AUBECSDebit))
					if valueAUBECSDebit, err := flattenPlainValue(
						sourceAUBECSDebit,
						types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "account_number": types.StringType}},
						"au_becs_debit",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAUBECSDebit, ok := valueAUBECSDebit.(types.Object); ok {
							state.AUBECSDebit = typedAUBECSDebit
							assignedAUBECSDebit = true
						}
					}
				}
			}
		}
		if !assignedAUBECSDebit && hadRawAUBECSDebit {
			if nullAUBECSDebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "account_number": types.StringType}}); ok {
				if typedAUBECSDebit, ok := nullAUBECSDebit.(types.Object); ok {
					state.AUBECSDebit = typedAUBECSDebit
				}
			}
		}
	}
	{
		assignedBACSDebit := false
		hadRawBACSDebit := false
		if rawValueBACSDebit, rawOk := plainValueAtPath(raw, "bacs_debit"); rawOk {
			hadRawBACSDebit = true
			if rawValueBACSDebit != nil {
				sourceBACSDebit := applyConfiguredKeyedListShapes(rawValueBACSDebit, attrValueToPlain(state.BACSDebit))
				if valueBACSDebit, err := flattenPlainValue(sourceBACSDebit, types.ObjectType{AttrTypes: map[string]attr.Type{"fingerprint": types.StringType, "last4": types.StringType, "sort_code": types.StringType, "account_number": types.StringType}}, "bacs_debit", "raw response"); err != nil {
					return err
				} else {
					if typedBACSDebit, ok := valueBACSDebit.(types.Object); ok {
						state.BACSDebit = typedBACSDebit
						assignedBACSDebit = true
					}
				}
			}
		}
		if !assignedBACSDebit {
			if !hasRaw {
				if responseValueBACSDebit, ok := plainFromResponseField(obj, "BACSDebit"); ok {
					sourceBACSDebit := applyConfiguredKeyedListShapes(responseValueBACSDebit, attrValueToPlain(state.BACSDebit))
					if valueBACSDebit, err := flattenPlainValue(
						sourceBACSDebit,
						types.ObjectType{AttrTypes: map[string]attr.Type{"fingerprint": types.StringType, "last4": types.StringType, "sort_code": types.StringType, "account_number": types.StringType}},
						"bacs_debit",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBACSDebit, ok := valueBACSDebit.(types.Object); ok {
							state.BACSDebit = typedBACSDebit
							assignedBACSDebit = true
						}
					}
				}
			}
		}
		if !assignedBACSDebit && hadRawBACSDebit {
			if nullBACSDebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"fingerprint": types.StringType, "last4": types.StringType, "sort_code": types.StringType, "account_number": types.StringType}}); ok {
				if typedBACSDebit, ok := nullBACSDebit.(types.Object); ok {
					state.BACSDebit = typedBACSDebit
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
		assignedBoleto := false
		hadRawBoleto := false
		if rawValueBoleto, rawOk := plainValueAtPath(raw, "boleto"); rawOk {
			hadRawBoleto = true
			if rawValueBoleto != nil {
				sourceBoleto := applyConfiguredKeyedListShapes(rawValueBoleto, attrValueToPlain(state.Boleto))
				if valueBoleto, err := flattenPlainValue(sourceBoleto, types.ObjectType{AttrTypes: map[string]attr.Type{"tax_id": types.StringType}}, "boleto", "raw response"); err != nil {
					return err
				} else {
					if typedBoleto, ok := valueBoleto.(types.Object); ok {
						state.Boleto = typedBoleto
						assignedBoleto = true
					}
				}
			}
		}
		if !assignedBoleto {
			if !hasRaw {
				if responseValueBoleto, ok := plainFromResponseField(obj, "Boleto"); ok {
					sourceBoleto := applyConfiguredKeyedListShapes(responseValueBoleto, attrValueToPlain(state.Boleto))
					if valueBoleto, err := flattenPlainValue(
						sourceBoleto,
						types.ObjectType{AttrTypes: map[string]attr.Type{"tax_id": types.StringType}},
						"boleto",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBoleto, ok := valueBoleto.(types.Object); ok {
							state.Boleto = typedBoleto
							assignedBoleto = true
						}
					}
				}
			}
		}
		if !assignedBoleto && hadRawBoleto {
			if nullBoleto, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"tax_id": types.StringType}}); ok {
				if typedBoleto, ok := nullBoleto.(types.Object); ok {
					state.Boleto = typedBoleto
				}
			}
		}
	}
	{
		assignedCard := false
		hadRawCard := false
		if rawValueCard, rawOk := plainValueAtPath(raw, "card"); rawOk {
			hadRawCard = true
			if rawValueCard != nil {
				sourceCard := applyConfiguredKeyedListShapes(rawValueCard, attrValueToPlain(state.Card))
				if valueCard, err := flattenPlainValue(sourceCard, types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "checks": types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_postal_code_check": types.StringType, "cvc_check": types.StringType}}, "country": types.StringType, "description": types.StringType, "display_brand": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_from": types.ObjectType{AttrTypes: map[string]attr.Type{"charge": types.StringType, "payment_method_details": types.ObjectType{AttrTypes: map[string]attr.Type{"card_present": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_authorized": types.Int64Type, "brand": types.StringType, "brand_product": types.StringType, "capture_before": types.Int64Type, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_card": types.StringType, "iin": types.StringType, "incremental_authorization_supported": types.BoolType, "issuer": types.StringType, "last4": types.StringType, "location": types.StringType, "network": types.StringType, "network_transaction_id": types.StringType, "offline": types.ObjectType{AttrTypes: map[string]attr.Type{"stored_at": types.Int64Type, "type": types.StringType}}, "overcapture_supported": types.BoolType, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "reader": types.StringType, "receipt": types.ObjectType{AttrTypes: map[string]attr.Type{"account_type": types.StringType, "application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "cardholder_verification_method": types.StringType, "dedicated_file_name": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}, "type": types.StringType}}, "setup_attempt": types.StringType}}, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.ListType{ElemType: types.StringType}, "preferred": types.StringType}}, "regulated_status": types.StringType, "three_d_secure_usage": types.ObjectType{AttrTypes: map[string]attr.Type{"supported": types.BoolType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"dynamic_last4": types.StringType, "masterpass": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}, "type": types.StringType, "visa_checkout": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}}}, "cvc": types.StringType, "number": types.StringType, "token": types.StringType}}, "card", "raw response"); err != nil {
					return err
				} else {
					if typedCard, ok := valueCard.(types.Object); ok {
						state.Card = typedCard
						assignedCard = true
					}
				}
			}
		}
		if !assignedCard {
			if !hasRaw {
				if responseValueCard, ok := plainFromResponseField(obj, "Card"); ok {
					sourceCard := applyConfiguredKeyedListShapes(responseValueCard, attrValueToPlain(state.Card))
					if valueCard, err := flattenPlainValue(
						sourceCard,
						types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "checks": types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_postal_code_check": types.StringType, "cvc_check": types.StringType}}, "country": types.StringType, "description": types.StringType, "display_brand": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_from": types.ObjectType{AttrTypes: map[string]attr.Type{"charge": types.StringType, "payment_method_details": types.ObjectType{AttrTypes: map[string]attr.Type{"card_present": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_authorized": types.Int64Type, "brand": types.StringType, "brand_product": types.StringType, "capture_before": types.Int64Type, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_card": types.StringType, "iin": types.StringType, "incremental_authorization_supported": types.BoolType, "issuer": types.StringType, "last4": types.StringType, "location": types.StringType, "network": types.StringType, "network_transaction_id": types.StringType, "offline": types.ObjectType{AttrTypes: map[string]attr.Type{"stored_at": types.Int64Type, "type": types.StringType}}, "overcapture_supported": types.BoolType, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "reader": types.StringType, "receipt": types.ObjectType{AttrTypes: map[string]attr.Type{"account_type": types.StringType, "application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "cardholder_verification_method": types.StringType, "dedicated_file_name": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}, "type": types.StringType}}, "setup_attempt": types.StringType}}, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.ListType{ElemType: types.StringType}, "preferred": types.StringType}}, "regulated_status": types.StringType, "three_d_secure_usage": types.ObjectType{AttrTypes: map[string]attr.Type{"supported": types.BoolType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"dynamic_last4": types.StringType, "masterpass": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}, "type": types.StringType, "visa_checkout": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}}}, "cvc": types.StringType, "number": types.StringType, "token": types.StringType}},
						"card",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCard, ok := valueCard.(types.Object); ok {
							state.Card = typedCard
							assignedCard = true
						}
					}
				}
			}
		}
		if !assignedCard && hadRawCard {
			if nullCard, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "checks": types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_postal_code_check": types.StringType, "cvc_check": types.StringType}}, "country": types.StringType, "description": types.StringType, "display_brand": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_from": types.ObjectType{AttrTypes: map[string]attr.Type{"charge": types.StringType, "payment_method_details": types.ObjectType{AttrTypes: map[string]attr.Type{"card_present": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_authorized": types.Int64Type, "brand": types.StringType, "brand_product": types.StringType, "capture_before": types.Int64Type, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "generated_card": types.StringType, "iin": types.StringType, "incremental_authorization_supported": types.BoolType, "issuer": types.StringType, "last4": types.StringType, "location": types.StringType, "network": types.StringType, "network_transaction_id": types.StringType, "offline": types.ObjectType{AttrTypes: map[string]attr.Type{"stored_at": types.Int64Type, "type": types.StringType}}, "overcapture_supported": types.BoolType, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "reader": types.StringType, "receipt": types.ObjectType{AttrTypes: map[string]attr.Type{"account_type": types.StringType, "application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "cardholder_verification_method": types.StringType, "dedicated_file_name": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}, "type": types.StringType}}, "setup_attempt": types.StringType}}, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.ListType{ElemType: types.StringType}, "preferred": types.StringType}}, "regulated_status": types.StringType, "three_d_secure_usage": types.ObjectType{AttrTypes: map[string]attr.Type{"supported": types.BoolType}}, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"dynamic_last4": types.StringType, "masterpass": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}, "type": types.StringType, "visa_checkout": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}}}, "cvc": types.StringType, "number": types.StringType, "token": types.StringType}}); ok {
				if typedCard, ok := nullCard.(types.Object); ok {
					state.Card = typedCard
				}
			}
		}
	}
	{
		assignedCardPresent := false
		hadRawCardPresent := false
		if rawValueCardPresent, rawOk := plainValueAtPath(raw, "card_present"); rawOk {
			hadRawCardPresent = true
			if rawValueCardPresent != nil {
				sourceCardPresent := applyConfiguredKeyedListShapes(rawValueCardPresent, attrValueToPlain(state.CardPresent))
				if valueCardPresent, err := flattenPlainValue(sourceCardPresent, types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "brand_product": types.StringType, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.ListType{ElemType: types.StringType}, "preferred": types.StringType}}, "offline": types.ObjectType{AttrTypes: map[string]attr.Type{"stored_at": types.Int64Type, "type": types.StringType}}, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}, "card_present", "raw response"); err != nil {
					return err
				} else {
					if typedCardPresent, ok := valueCardPresent.(types.Object); ok {
						state.CardPresent = typedCardPresent
						assignedCardPresent = true
					}
				}
			}
		}
		if !assignedCardPresent {
			if !hasRaw {
				if responseValueCardPresent, ok := plainFromResponseField(obj, "CardPresent"); ok {
					sourceCardPresent := applyConfiguredKeyedListShapes(responseValueCardPresent, attrValueToPlain(state.CardPresent))
					if valueCardPresent, err := flattenPlainValue(
						sourceCardPresent,
						types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "brand_product": types.StringType, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.ListType{ElemType: types.StringType}, "preferred": types.StringType}}, "offline": types.ObjectType{AttrTypes: map[string]attr.Type{"stored_at": types.Int64Type, "type": types.StringType}}, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}},
						"card_present",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCardPresent, ok := valueCardPresent.(types.Object); ok {
							state.CardPresent = typedCardPresent
							assignedCardPresent = true
						}
					}
				}
			}
		}
		if !assignedCardPresent && hadRawCardPresent {
			if nullCardPresent, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "brand_product": types.StringType, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.ListType{ElemType: types.StringType}, "preferred": types.StringType}}, "offline": types.ObjectType{AttrTypes: map[string]attr.Type{"stored_at": types.Int64Type, "type": types.StringType}}, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType, "wallet": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}); ok {
				if typedCardPresent, ok := nullCardPresent.(types.Object); ok {
					state.CardPresent = typedCardPresent
				}
			}
		}
	}
	{
		assignedCashApp := false
		hadRawCashApp := false
		if rawValueCashApp, rawOk := plainValueAtPath(raw, "cashapp"); rawOk {
			hadRawCashApp = true
			if rawValueCashApp != nil {
				sourceCashApp := applyConfiguredKeyedListShapes(rawValueCashApp, attrValueToPlain(state.CashApp))
				if valueCashApp, err := flattenPlainValue(sourceCashApp, types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "cashtag": types.StringType}}, "cashapp", "raw response"); err != nil {
					return err
				} else {
					if typedCashApp, ok := valueCashApp.(types.Object); ok {
						state.CashApp = typedCashApp
						assignedCashApp = true
					}
				}
			}
		}
		if !assignedCashApp {
			if !hasRaw {
				if responseValueCashApp, ok := plainFromResponseField(obj, "CashApp"); ok {
					sourceCashApp := applyConfiguredKeyedListShapes(responseValueCashApp, attrValueToPlain(state.CashApp))
					if valueCashApp, err := flattenPlainValue(
						sourceCashApp,
						types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "cashtag": types.StringType}},
						"cashapp",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCashApp, ok := valueCashApp.(types.Object); ok {
							state.CashApp = typedCashApp
							assignedCashApp = true
						}
					}
				}
			}
		}
		if !assignedCashApp && hadRawCashApp {
			if nullCashApp, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "cashtag": types.StringType}}); ok {
				if typedCashApp, ok := nullCashApp.(types.Object); ok {
					state.CashApp = typedCashApp
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
		if state.Customer.IsNull() || state.Customer.IsUnknown() {
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
		assignedEPS := false
		hadRawEPS := false
		if rawValueEPS, rawOk := plainValueAtPath(raw, "eps"); rawOk {
			hadRawEPS = true
			if rawValueEPS != nil {
				sourceEPS := applyConfiguredKeyedListShapes(rawValueEPS, attrValueToPlain(state.EPS))
				if valueEPS, err := flattenPlainValue(sourceEPS, types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType}}, "eps", "raw response"); err != nil {
					return err
				} else {
					if typedEPS, ok := valueEPS.(types.Object); ok {
						state.EPS = typedEPS
						assignedEPS = true
					}
				}
			}
		}
		if !assignedEPS {
			if !hasRaw {
				if responseValueEPS, ok := plainFromResponseField(obj, "EPS"); ok {
					sourceEPS := applyConfiguredKeyedListShapes(responseValueEPS, attrValueToPlain(state.EPS))
					if valueEPS, err := flattenPlainValue(
						sourceEPS,
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType}},
						"eps",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedEPS, ok := valueEPS.(types.Object); ok {
							state.EPS = typedEPS
							assignedEPS = true
						}
					}
				}
			}
		}
		if !assignedEPS && hadRawEPS {
			if nullEPS, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType}}); ok {
				if typedEPS, ok := nullEPS.(types.Object); ok {
					state.EPS = typedEPS
				}
			}
		}
	}
	{
		assignedFPX := false
		hadRawFPX := false
		if rawValueFPX, rawOk := plainValueAtPath(raw, "fpx"); rawOk {
			hadRawFPX = true
			if rawValueFPX != nil {
				sourceFPX := applyConfiguredKeyedListShapes(rawValueFPX, attrValueToPlain(state.FPX))
				if valueFPX, err := flattenPlainValue(sourceFPX, types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "bank": types.StringType}}, "fpx", "raw response"); err != nil {
					return err
				} else {
					if typedFPX, ok := valueFPX.(types.Object); ok {
						state.FPX = typedFPX
						assignedFPX = true
					}
				}
			}
		}
		if !assignedFPX {
			if !hasRaw {
				if responseValueFPX, ok := plainFromResponseField(obj, "FPX"); ok {
					sourceFPX := applyConfiguredKeyedListShapes(responseValueFPX, attrValueToPlain(state.FPX))
					if valueFPX, err := flattenPlainValue(
						sourceFPX,
						types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "bank": types.StringType}},
						"fpx",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFPX, ok := valueFPX.(types.Object); ok {
							state.FPX = typedFPX
							assignedFPX = true
						}
					}
				}
			}
		}
		if !assignedFPX && hadRawFPX {
			if nullFPX, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "bank": types.StringType}}); ok {
				if typedFPX, ok := nullFPX.(types.Object); ok {
					state.FPX = typedFPX
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
		assignedIDEAL := false
		hadRawIDEAL := false
		if rawValueIDEAL, rawOk := plainValueAtPath(raw, "ideal"); rawOk {
			hadRawIDEAL = true
			if rawValueIDEAL != nil {
				sourceIDEAL := applyConfiguredKeyedListShapes(rawValueIDEAL, attrValueToPlain(state.IDEAL))
				if valueIDEAL, err := flattenPlainValue(sourceIDEAL, types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "bic": types.StringType}}, "ideal", "raw response"); err != nil {
					return err
				} else {
					if typedIDEAL, ok := valueIDEAL.(types.Object); ok {
						state.IDEAL = typedIDEAL
						assignedIDEAL = true
					}
				}
			}
		}
		if !assignedIDEAL {
			if !hasRaw {
				if responseValueIDEAL, ok := plainFromResponseField(obj, "IDEAL"); ok {
					sourceIDEAL := applyConfiguredKeyedListShapes(responseValueIDEAL, attrValueToPlain(state.IDEAL))
					if valueIDEAL, err := flattenPlainValue(
						sourceIDEAL,
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "bic": types.StringType}},
						"ideal",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedIDEAL, ok := valueIDEAL.(types.Object); ok {
							state.IDEAL = typedIDEAL
							assignedIDEAL = true
						}
					}
				}
			}
		}
		if !assignedIDEAL && hadRawIDEAL {
			if nullIDEAL, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "bic": types.StringType}}); ok {
				if typedIDEAL, ok := nullIDEAL.(types.Object); ok {
					state.IDEAL = typedIDEAL
				}
			}
		}
	}
	{
		assignedInteracPresent := false
		hadRawInteracPresent := false
		if rawValueInteracPresent, rawOk := plainValueAtPath(raw, "interac_present"); rawOk {
			hadRawInteracPresent = true
			if rawValueInteracPresent != nil {
				sourceInteracPresent := applyConfiguredKeyedListShapes(rawValueInteracPresent, attrValueToPlain(state.InteracPresent))
				if valueInteracPresent, err := flattenPlainValue(sourceInteracPresent, types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.ListType{ElemType: types.StringType}, "preferred": types.StringType}}, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType}}, "interac_present", "raw response"); err != nil {
					return err
				} else {
					if typedInteracPresent, ok := valueInteracPresent.(types.Object); ok {
						state.InteracPresent = typedInteracPresent
						assignedInteracPresent = true
					}
				}
			}
		}
		if !assignedInteracPresent {
			if !hasRaw {
				if responseValueInteracPresent, ok := plainFromResponseField(obj, "InteracPresent"); ok {
					sourceInteracPresent := applyConfiguredKeyedListShapes(responseValueInteracPresent, attrValueToPlain(state.InteracPresent))
					if valueInteracPresent, err := flattenPlainValue(
						sourceInteracPresent,
						types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.ListType{ElemType: types.StringType}, "preferred": types.StringType}}, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType}},
						"interac_present",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedInteracPresent, ok := valueInteracPresent.(types.Object); ok {
							state.InteracPresent = typedInteracPresent
							assignedInteracPresent = true
						}
					}
				}
			}
		}
		if !assignedInteracPresent && hadRawInteracPresent {
			if nullInteracPresent, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "cardholder_name": types.StringType, "country": types.StringType, "description": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.ListType{ElemType: types.StringType}, "preferred": types.StringType}}, "preferred_locales": types.ListType{ElemType: types.StringType}, "read_method": types.StringType}}); ok {
				if typedInteracPresent, ok := nullInteracPresent.(types.Object); ok {
					state.InteracPresent = typedInteracPresent
				}
			}
		}
	}
	{
		assignedKlarna := false
		hadRawKlarna := false
		if rawValueKlarna, rawOk := plainValueAtPath(raw, "klarna"); rawOk {
			hadRawKlarna = true
			if rawValueKlarna != nil {
				sourceKlarna := applyConfiguredKeyedListShapes(rawValueKlarna, attrValueToPlain(state.Klarna))
				if valueKlarna, err := flattenPlainValue(sourceKlarna, types.ObjectType{AttrTypes: map[string]attr.Type{"dob": types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}}}, "klarna", "raw response"); err != nil {
					return err
				} else {
					if typedKlarna, ok := valueKlarna.(types.Object); ok {
						state.Klarna = typedKlarna
						assignedKlarna = true
					}
				}
			}
		}
		if !assignedKlarna {
			if !hasRaw {
				if responseValueKlarna, ok := plainFromResponseField(obj, "Klarna"); ok {
					sourceKlarna := applyConfiguredKeyedListShapes(responseValueKlarna, attrValueToPlain(state.Klarna))
					if valueKlarna, err := flattenPlainValue(
						sourceKlarna,
						types.ObjectType{AttrTypes: map[string]attr.Type{"dob": types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}}},
						"klarna",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedKlarna, ok := valueKlarna.(types.Object); ok {
							state.Klarna = typedKlarna
							assignedKlarna = true
						}
					}
				}
			}
		}
		if !assignedKlarna && hadRawKlarna {
			if nullKlarna, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"dob": types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}}}); ok {
				if typedKlarna, ok := nullKlarna.(types.Object); ok {
					state.Klarna = typedKlarna
				}
			}
		}
	}
	{
		assignedKrCard := false
		hadRawKrCard := false
		if rawValueKrCard, rawOk := plainValueAtPath(raw, "kr_card"); rawOk {
			hadRawKrCard = true
			if rawValueKrCard != nil {
				sourceKrCard := applyConfiguredKeyedListShapes(rawValueKrCard, attrValueToPlain(state.KrCard))
				if valueKrCard, err := flattenPlainValue(sourceKrCard, types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "last4": types.StringType}}, "kr_card", "raw response"); err != nil {
					return err
				} else {
					if typedKrCard, ok := valueKrCard.(types.Object); ok {
						state.KrCard = typedKrCard
						assignedKrCard = true
					}
				}
			}
		}
		if !assignedKrCard {
			if !hasRaw {
				if responseValueKrCard, ok := plainFromResponseField(obj, "KrCard"); ok {
					sourceKrCard := applyConfiguredKeyedListShapes(responseValueKrCard, attrValueToPlain(state.KrCard))
					if valueKrCard, err := flattenPlainValue(
						sourceKrCard,
						types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "last4": types.StringType}},
						"kr_card",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedKrCard, ok := valueKrCard.(types.Object); ok {
							state.KrCard = typedKrCard
							assignedKrCard = true
						}
					}
				}
			}
		}
		if !assignedKrCard && hadRawKrCard {
			if nullKrCard, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"brand": types.StringType, "last4": types.StringType}}); ok {
				if typedKrCard, ok := nullKrCard.(types.Object); ok {
					state.KrCard = typedKrCard
				}
			}
		}
	}
	{
		assignedLink := false
		hadRawLink := false
		if rawValueLink, rawOk := plainValueAtPath(raw, "link"); rawOk {
			hadRawLink = true
			if rawValueLink != nil {
				sourceLink := applyConfiguredKeyedListShapes(rawValueLink, attrValueToPlain(state.Link))
				if valueLink, err := flattenPlainValue(sourceLink, types.ObjectType{AttrTypes: map[string]attr.Type{"email": types.StringType, "persistent_token": types.StringType}}, "link", "raw response"); err != nil {
					return err
				} else {
					if typedLink, ok := valueLink.(types.Object); ok {
						state.Link = typedLink
						assignedLink = true
					}
				}
			}
		}
		if !assignedLink {
			if !hasRaw {
				if responseValueLink, ok := plainFromResponseField(obj, "Link"); ok {
					sourceLink := applyConfiguredKeyedListShapes(responseValueLink, attrValueToPlain(state.Link))
					if valueLink, err := flattenPlainValue(
						sourceLink,
						types.ObjectType{AttrTypes: map[string]attr.Type{"email": types.StringType, "persistent_token": types.StringType}},
						"link",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLink, ok := valueLink.(types.Object); ok {
							state.Link = typedLink
							assignedLink = true
						}
					}
				}
			}
		}
		if !assignedLink && hadRawLink {
			if nullLink, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"email": types.StringType, "persistent_token": types.StringType}}); ok {
				if typedLink, ok := nullLink.(types.Object); ok {
					state.Link = typedLink
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
		assignedNaverPay := false
		hadRawNaverPay := false
		if rawValueNaverPay, rawOk := plainValueAtPath(raw, "naver_pay"); rawOk {
			hadRawNaverPay = true
			if rawValueNaverPay != nil {
				sourceNaverPay := applyConfiguredKeyedListShapes(rawValueNaverPay, attrValueToPlain(state.NaverPay))
				if valueNaverPay, err := flattenPlainValue(sourceNaverPay, types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "funding": types.StringType}}, "naver_pay", "raw response"); err != nil {
					return err
				} else {
					if typedNaverPay, ok := valueNaverPay.(types.Object); ok {
						state.NaverPay = typedNaverPay
						assignedNaverPay = true
					}
				}
			}
		}
		if !assignedNaverPay {
			if !hasRaw {
				if responseValueNaverPay, ok := plainFromResponseField(obj, "NaverPay"); ok {
					sourceNaverPay := applyConfiguredKeyedListShapes(responseValueNaverPay, attrValueToPlain(state.NaverPay))
					if valueNaverPay, err := flattenPlainValue(
						sourceNaverPay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "funding": types.StringType}},
						"naver_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedNaverPay, ok := valueNaverPay.(types.Object); ok {
							state.NaverPay = typedNaverPay
							assignedNaverPay = true
						}
					}
				}
			}
		}
		if !assignedNaverPay && hadRawNaverPay {
			if nullNaverPay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"buyer_id": types.StringType, "funding": types.StringType}}); ok {
				if typedNaverPay, ok := nullNaverPay.(types.Object); ok {
					state.NaverPay = typedNaverPay
				}
			}
		}
	}
	{
		assignedNzBankAccount := false
		hadRawNzBankAccount := false
		if rawValueNzBankAccount, rawOk := plainValueAtPath(raw, "nz_bank_account"); rawOk {
			hadRawNzBankAccount = true
			if rawValueNzBankAccount != nil {
				sourceNzBankAccount := applyConfiguredKeyedListShapes(rawValueNzBankAccount, attrValueToPlain(state.NzBankAccount))
				if valueNzBankAccount, err := flattenPlainValue(sourceNzBankAccount, types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "bank_code": types.StringType, "bank_name": types.StringType, "branch_code": types.StringType, "last4": types.StringType, "suffix": types.StringType, "account_number": types.StringType, "reference": types.StringType}}, "nz_bank_account", "raw response"); err != nil {
					return err
				} else {
					if typedNzBankAccount, ok := valueNzBankAccount.(types.Object); ok {
						state.NzBankAccount = typedNzBankAccount
						assignedNzBankAccount = true
					}
				}
			}
		}
		if !assignedNzBankAccount {
			if !hasRaw {
				if responseValueNzBankAccount, ok := plainFromResponseField(obj, "NzBankAccount"); ok {
					sourceNzBankAccount := applyConfiguredKeyedListShapes(responseValueNzBankAccount, attrValueToPlain(state.NzBankAccount))
					if valueNzBankAccount, err := flattenPlainValue(
						sourceNzBankAccount,
						types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "bank_code": types.StringType, "bank_name": types.StringType, "branch_code": types.StringType, "last4": types.StringType, "suffix": types.StringType, "account_number": types.StringType, "reference": types.StringType}},
						"nz_bank_account",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedNzBankAccount, ok := valueNzBankAccount.(types.Object); ok {
							state.NzBankAccount = typedNzBankAccount
							assignedNzBankAccount = true
						}
					}
				}
			}
		}
		if !assignedNzBankAccount && hadRawNzBankAccount {
			if nullNzBankAccount, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "bank_code": types.StringType, "bank_name": types.StringType, "branch_code": types.StringType, "last4": types.StringType, "suffix": types.StringType, "account_number": types.StringType, "reference": types.StringType}}); ok {
				if typedNzBankAccount, ok := nullNzBankAccount.(types.Object); ok {
					state.NzBankAccount = typedNzBankAccount
				}
			}
		}
	}
	{
		assignedP24 := false
		hadRawP24 := false
		if rawValueP24, rawOk := plainValueAtPath(raw, "p24"); rawOk {
			hadRawP24 = true
			if rawValueP24 != nil {
				sourceP24 := applyConfiguredKeyedListShapes(rawValueP24, attrValueToPlain(state.P24))
				if valueP24, err := flattenPlainValue(sourceP24, types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType}}, "p24", "raw response"); err != nil {
					return err
				} else {
					if typedP24, ok := valueP24.(types.Object); ok {
						state.P24 = typedP24
						assignedP24 = true
					}
				}
			}
		}
		if !assignedP24 {
			if !hasRaw {
				if responseValueP24, ok := plainFromResponseField(obj, "P24"); ok {
					sourceP24 := applyConfiguredKeyedListShapes(responseValueP24, attrValueToPlain(state.P24))
					if valueP24, err := flattenPlainValue(
						sourceP24,
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType}},
						"p24",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedP24, ok := valueP24.(types.Object); ok {
							state.P24 = typedP24
							assignedP24 = true
						}
					}
				}
			}
		}
		if !assignedP24 && hadRawP24 {
			if nullP24, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType}}); ok {
				if typedP24, ok := nullP24.(types.Object); ok {
					state.P24 = typedP24
				}
			}
		}
	}
	{
		assignedPaypal := false
		hadRawPaypal := false
		if rawValuePaypal, rawOk := plainValueAtPath(raw, "paypal"); rawOk {
			hadRawPaypal = true
			if rawValuePaypal != nil {
				sourcePaypal := applyConfiguredKeyedListShapes(rawValuePaypal, attrValueToPlain(state.Paypal))
				if valuePaypal, err := flattenPlainValue(sourcePaypal, types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "payer_email": types.StringType, "payer_id": types.StringType}}, "paypal", "raw response"); err != nil {
					return err
				} else {
					if typedPaypal, ok := valuePaypal.(types.Object); ok {
						state.Paypal = typedPaypal
						assignedPaypal = true
					}
				}
			}
		}
		if !assignedPaypal {
			if !hasRaw {
				if responseValuePaypal, ok := plainFromResponseField(obj, "Paypal"); ok {
					sourcePaypal := applyConfiguredKeyedListShapes(responseValuePaypal, attrValueToPlain(state.Paypal))
					if valuePaypal, err := flattenPlainValue(
						sourcePaypal,
						types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "payer_email": types.StringType, "payer_id": types.StringType}},
						"paypal",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPaypal, ok := valuePaypal.(types.Object); ok {
							state.Paypal = typedPaypal
							assignedPaypal = true
						}
					}
				}
			}
		}
		if !assignedPaypal && hadRawPaypal {
			if nullPaypal, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "payer_email": types.StringType, "payer_id": types.StringType}}); ok {
				if typedPaypal, ok := nullPaypal.(types.Object); ok {
					state.Paypal = typedPaypal
				}
			}
		}
	}
	{
		assignedPayto := false
		hadRawPayto := false
		if rawValuePayto, rawOk := plainValueAtPath(raw, "payto"); rawOk {
			hadRawPayto = true
			if rawValuePayto != nil {
				sourcePayto := applyConfiguredKeyedListShapes(rawValuePayto, attrValueToPlain(state.Payto))
				if valuePayto, err := flattenPlainValue(sourcePayto, types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "last4": types.StringType, "pay_id": types.StringType, "account_number": types.StringType}}, "payto", "raw response"); err != nil {
					return err
				} else {
					if typedPayto, ok := valuePayto.(types.Object); ok {
						state.Payto = typedPayto
						assignedPayto = true
					}
				}
			}
		}
		if !assignedPayto {
			if !hasRaw {
				if responseValuePayto, ok := plainFromResponseField(obj, "Payto"); ok {
					sourcePayto := applyConfiguredKeyedListShapes(responseValuePayto, attrValueToPlain(state.Payto))
					if valuePayto, err := flattenPlainValue(
						sourcePayto,
						types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "last4": types.StringType, "pay_id": types.StringType, "account_number": types.StringType}},
						"payto",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPayto, ok := valuePayto.(types.Object); ok {
							state.Payto = typedPayto
							assignedPayto = true
						}
					}
				}
			}
		}
		if !assignedPayto && hadRawPayto {
			if nullPayto, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "last4": types.StringType, "pay_id": types.StringType, "account_number": types.StringType}}); ok {
				if typedPayto, ok := nullPayto.(types.Object); ok {
					state.Payto = typedPayto
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
		assignedSEPADebit := false
		hadRawSEPADebit := false
		if rawValueSEPADebit, rawOk := plainValueAtPath(raw, "sepa_debit"); rawOk {
			hadRawSEPADebit = true
			if rawValueSEPADebit != nil {
				sourceSEPADebit := applyConfiguredKeyedListShapes(rawValueSEPADebit, attrValueToPlain(state.SEPADebit))
				if valueSEPADebit, err := flattenPlainValue(sourceSEPADebit, types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "branch_code": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "generated_from": types.ObjectType{AttrTypes: map[string]attr.Type{"charge": types.StringType, "setup_attempt": types.StringType}}, "last4": types.StringType, "iban": types.StringType}}, "sepa_debit", "raw response"); err != nil {
					return err
				} else {
					if typedSEPADebit, ok := valueSEPADebit.(types.Object); ok {
						state.SEPADebit = typedSEPADebit
						assignedSEPADebit = true
					}
				}
			}
		}
		if !assignedSEPADebit {
			if !hasRaw {
				if responseValueSEPADebit, ok := plainFromResponseField(obj, "SEPADebit"); ok {
					sourceSEPADebit := applyConfiguredKeyedListShapes(responseValueSEPADebit, attrValueToPlain(state.SEPADebit))
					if valueSEPADebit, err := flattenPlainValue(
						sourceSEPADebit,
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "branch_code": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "generated_from": types.ObjectType{AttrTypes: map[string]attr.Type{"charge": types.StringType, "setup_attempt": types.StringType}}, "last4": types.StringType, "iban": types.StringType}},
						"sepa_debit",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSEPADebit, ok := valueSEPADebit.(types.Object); ok {
							state.SEPADebit = typedSEPADebit
							assignedSEPADebit = true
						}
					}
				}
			}
		}
		if !assignedSEPADebit && hadRawSEPADebit {
			if nullSEPADebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "branch_code": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "generated_from": types.ObjectType{AttrTypes: map[string]attr.Type{"charge": types.StringType, "setup_attempt": types.StringType}}, "last4": types.StringType, "iban": types.StringType}}); ok {
				if typedSEPADebit, ok := nullSEPADebit.(types.Object); ok {
					state.SEPADebit = typedSEPADebit
				}
			}
		}
	}
	{
		assignedSofort := false
		hadRawSofort := false
		if rawValueSofort, rawOk := plainValueAtPath(raw, "sofort"); rawOk {
			hadRawSofort = true
			if rawValueSofort != nil {
				sourceSofort := applyConfiguredKeyedListShapes(rawValueSofort, attrValueToPlain(state.Sofort))
				if valueSofort, err := flattenPlainValue(sourceSofort, types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "sofort", "raw response"); err != nil {
					return err
				} else {
					if typedSofort, ok := valueSofort.(types.Object); ok {
						state.Sofort = typedSofort
						assignedSofort = true
					}
				}
			}
		}
		if !assignedSofort {
			if !hasRaw {
				if responseValueSofort, ok := plainFromResponseField(obj, "Sofort"); ok {
					sourceSofort := applyConfiguredKeyedListShapes(responseValueSofort, attrValueToPlain(state.Sofort))
					if valueSofort, err := flattenPlainValue(
						sourceSofort,
						types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}},
						"sofort",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSofort, ok := valueSofort.(types.Object); ok {
							state.Sofort = typedSofort
							assignedSofort = true
						}
					}
				}
			}
		}
		if !assignedSofort && hadRawSofort {
			if nullSofort, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}); ok {
				if typedSofort, ok := nullSofort.(types.Object); ok {
					state.Sofort = typedSofort
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
		assignedUpi := false
		hadRawUpi := false
		if rawValueUpi, rawOk := plainValueAtPath(raw, "upi"); rawOk {
			hadRawUpi = true
			if rawValueUpi != nil {
				sourceUpi := applyConfiguredKeyedListShapes(rawValueUpi, attrValueToPlain(state.Upi))
				if valueUpi, err := flattenPlainValue(sourceUpi, types.ObjectType{AttrTypes: map[string]attr.Type{"vpa": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}, "upi", "raw response"); err != nil {
					return err
				} else {
					if typedUpi, ok := valueUpi.(types.Object); ok {
						state.Upi = typedUpi
						assignedUpi = true
					}
				}
			}
		}
		if !assignedUpi {
			if !hasRaw {
				if responseValueUpi, ok := plainFromResponseField(obj, "Upi"); ok {
					sourceUpi := applyConfiguredKeyedListShapes(responseValueUpi, attrValueToPlain(state.Upi))
					if valueUpi, err := flattenPlainValue(
						sourceUpi,
						types.ObjectType{AttrTypes: map[string]attr.Type{"vpa": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}},
						"upi",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedUpi, ok := valueUpi.(types.Object); ok {
							state.Upi = typedUpi
							assignedUpi = true
						}
					}
				}
			}
		}
		if !assignedUpi && hadRawUpi {
			if nullUpi, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"vpa": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}); ok {
				if typedUpi, ok := nullUpi.(types.Object); ok {
					state.Upi = typedUpi
				}
			}
		}
	}
	{
		assignedUSBankAccount := false
		hadRawUSBankAccount := false
		if rawValueUSBankAccount, rawOk := plainValueAtPath(raw, "us_bank_account"); rawOk {
			hadRawUSBankAccount = true
			if rawValueUSBankAccount != nil {
				sourceUSBankAccount := applyConfiguredKeyedListShapes(rawValueUSBankAccount, attrValueToPlain(state.USBankAccount))
				if valueUSBankAccount, err := flattenPlainValue(sourceUSBankAccount, types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "account_type": types.StringType, "bank_name": types.StringType, "financial_connections_account": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred": types.StringType, "supported": types.ListType{ElemType: types.StringType}}}, "routing_number": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"blocked": types.ObjectType{AttrTypes: map[string]attr.Type{"network_code": types.StringType, "reason": types.StringType}}}}, "account_number": types.StringType}}, "us_bank_account", "raw response"); err != nil {
					return err
				} else {
					if typedUSBankAccount, ok := valueUSBankAccount.(types.Object); ok {
						state.USBankAccount = typedUSBankAccount
						assignedUSBankAccount = true
					}
				}
			}
		}
		if !assignedUSBankAccount {
			if !hasRaw {
				if responseValueUSBankAccount, ok := plainFromResponseField(obj, "USBankAccount"); ok {
					sourceUSBankAccount := applyConfiguredKeyedListShapes(responseValueUSBankAccount, attrValueToPlain(state.USBankAccount))
					if valueUSBankAccount, err := flattenPlainValue(
						sourceUSBankAccount,
						types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "account_type": types.StringType, "bank_name": types.StringType, "financial_connections_account": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred": types.StringType, "supported": types.ListType{ElemType: types.StringType}}}, "routing_number": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"blocked": types.ObjectType{AttrTypes: map[string]attr.Type{"network_code": types.StringType, "reason": types.StringType}}}}, "account_number": types.StringType}},
						"us_bank_account",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedUSBankAccount, ok := valueUSBankAccount.(types.Object); ok {
							state.USBankAccount = typedUSBankAccount
							assignedUSBankAccount = true
						}
					}
				}
			}
		}
		if !assignedUSBankAccount && hadRawUSBankAccount {
			if nullUSBankAccount, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_type": types.StringType, "account_type": types.StringType, "bank_name": types.StringType, "financial_connections_account": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "networks": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred": types.StringType, "supported": types.ListType{ElemType: types.StringType}}}, "routing_number": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"blocked": types.ObjectType{AttrTypes: map[string]attr.Type{"network_code": types.StringType, "reason": types.StringType}}}}, "account_number": types.StringType}}); ok {
				if typedUSBankAccount, ok := nullUSBankAccount.(types.Object); ok {
					state.USBankAccount = typedUSBankAccount
				}
			}
		}
	}
	return nil
}
