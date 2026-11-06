//
// File generated from our OpenAPI spec
//

package ephemeralresources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	ephemeralSchema "github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ ephemeral.EphemeralResource = &CheckoutSessionEphemeralResource{}
var _ ephemeral.EphemeralResourceWithConfigure = &CheckoutSessionEphemeralResource{}

func NewCheckoutSessionEphemeralResource() ephemeral.EphemeralResource {
	return &CheckoutSessionEphemeralResource{}
}

type CheckoutSessionEphemeralResource struct {
	client *stripe.Client
}

type CheckoutSessionResourceModel struct {
	Object                            types.String `tfsdk:"object"`
	AdaptivePricing                   types.Object `tfsdk:"adaptive_pricing"`
	AfterExpiration                   types.Object `tfsdk:"after_expiration"`
	AllowPromotionCodes               types.Bool   `tfsdk:"allow_promotion_codes"`
	AmountSubtotal                    types.Int64  `tfsdk:"amount_subtotal"`
	AmountTotal                       types.Int64  `tfsdk:"amount_total"`
	AutomaticTax                      types.Object `tfsdk:"automatic_tax"`
	BillingAddressCollection          types.String `tfsdk:"billing_address_collection"`
	BrandingSettings                  types.Object `tfsdk:"branding_settings"`
	CancelURL                         types.String `tfsdk:"cancel_url"`
	ClientReferenceID                 types.String `tfsdk:"client_reference_id"`
	ClientSecret                      types.String `tfsdk:"client_secret"`
	CollectedInformation              types.Object `tfsdk:"collected_information"`
	Consent                           types.Object `tfsdk:"consent"`
	ConsentCollection                 types.Object `tfsdk:"consent_collection"`
	Created                           types.Int64  `tfsdk:"created"`
	Currency                          types.String `tfsdk:"currency"`
	CurrencyConversion                types.Object `tfsdk:"currency_conversion"`
	CustomFields                      types.List   `tfsdk:"custom_fields"`
	CustomText                        types.Object `tfsdk:"custom_text"`
	Customer                          types.String `tfsdk:"customer"`
	CustomerAccount                   types.String `tfsdk:"customer_account"`
	CustomerCreation                  types.String `tfsdk:"customer_creation"`
	CustomerDetails                   types.Object `tfsdk:"customer_details"`
	CustomerEmail                     types.String `tfsdk:"customer_email"`
	Discounts                         types.List   `tfsdk:"discounts"`
	ExcludedPaymentMethodTypes        types.List   `tfsdk:"excluded_payment_method_types"`
	ExpiresAt                         types.Int64  `tfsdk:"expires_at"`
	ID                                types.String `tfsdk:"id"`
	IntegrationIdentifier             types.String `tfsdk:"integration_identifier"`
	Invoice                           types.String `tfsdk:"invoice"`
	InvoiceCreation                   types.Object `tfsdk:"invoice_creation"`
	LineItems                         types.List   `tfsdk:"line_items"`
	Livemode                          types.Bool   `tfsdk:"livemode"`
	Locale                            types.String `tfsdk:"locale"`
	ManagedPayments                   types.Object `tfsdk:"managed_payments"`
	Metadata                          types.Map    `tfsdk:"metadata"`
	Mode                              types.String `tfsdk:"mode"`
	NameCollection                    types.Object `tfsdk:"name_collection"`
	OptionalItems                     types.List   `tfsdk:"optional_items"`
	OriginContext                     types.String `tfsdk:"origin_context"`
	PaymentIntent                     types.String `tfsdk:"payment_intent"`
	PaymentLink                       types.String `tfsdk:"payment_link"`
	PaymentMethodCollection           types.String `tfsdk:"payment_method_collection"`
	PaymentMethodConfigurationDetails types.Object `tfsdk:"payment_method_configuration_details"`
	PaymentMethodOptions              types.Object `tfsdk:"payment_method_options"`
	PaymentMethodTypes                types.List   `tfsdk:"payment_method_types"`
	PaymentStatus                     types.String `tfsdk:"payment_status"`
	Permissions                       types.Object `tfsdk:"permissions"`
	PhoneNumberCollection             types.Object `tfsdk:"phone_number_collection"`
	PresentmentDetails                types.Object `tfsdk:"presentment_details"`
	RecoveredFrom                     types.String `tfsdk:"recovered_from"`
	RedirectOnCompletion              types.String `tfsdk:"redirect_on_completion"`
	ReturnURL                         types.String `tfsdk:"return_url"`
	SavedPaymentMethodOptions         types.Object `tfsdk:"saved_payment_method_options"`
	SetupIntent                       types.String `tfsdk:"setup_intent"`
	ShippingAddressCollection         types.Object `tfsdk:"shipping_address_collection"`
	ShippingCost                      types.Object `tfsdk:"shipping_cost"`
	ShippingOptions                   types.List   `tfsdk:"shipping_options"`
	Status                            types.String `tfsdk:"status"`
	SubmitType                        types.String `tfsdk:"submit_type"`
	Subscription                      types.String `tfsdk:"subscription"`
	SuccessURL                        types.String `tfsdk:"success_url"`
	TaxIDCollection                   types.Object `tfsdk:"tax_id_collection"`
	TotalDetails                      types.Object `tfsdk:"total_details"`
	UIMode                            types.String `tfsdk:"ui_mode"`
	URL                               types.String `tfsdk:"url"`
	WalletOptions                     types.Object `tfsdk:"wallet_options"`
	CustomerUpdate                    types.Object `tfsdk:"customer_update"`
	PaymentIntentData                 types.Object `tfsdk:"payment_intent_data"`
	PaymentMethodConfiguration        types.String `tfsdk:"payment_method_configuration"`
	PaymentMethodData                 types.Object `tfsdk:"payment_method_data"`
	SetupIntentData                   types.Object `tfsdk:"setup_intent_data"`
	SubscriptionData                  types.Object `tfsdk:"subscription_data"`
}

func (r *CheckoutSessionEphemeralResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_checkout_session"
}

func (r *CheckoutSessionEphemeralResource) Schema(_ context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = ephemeralSchema.Schema{
		Description: "A Checkout Session represents your customer's session as they pay for\none-time purchases or subscriptions through [Checkout](https://docs.stripe.com/payments/checkout)\nor [Payment Links](https://docs.stripe.com/payments/payment-links). We recommend creating a\nnew Session each time your customer attempts to pay.\n\nOnce payment is successful, the Checkout Session will contain a reference\nto the [Customer](https://docs.stripe.com/api/customers), and either the successful\n[PaymentIntent](https://docs.stripe.com/api/payment_intents) or an active\n[Subscription](https://docs.stripe.com/api/subscriptions).\n\nYou can create a Checkout Session on your server and redirect to its URL\nto begin Checkout.\n\nRelated guide: [Checkout quickstart](https://docs.stripe.com/checkout/quickstart)",
		Attributes: map[string]ephemeralSchema.Attribute{
			"object": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "String representing the object's type. Objects of the same type share the same value.",
				Validators:  []validator.String{stringvalidator.OneOf("checkout.session")},
			},
			"adaptive_pricing": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Settings for price localization with [Adaptive Pricing](https://docs.stripe.com/payments/checkout/adaptive-pricing).",
				Attributes: map[string]ephemeralSchema.Attribute{
					"enabled": ephemeralSchema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Description: "If enabled, Adaptive Pricing is available on [eligible sessions](https://docs.stripe.com/payments/currencies/localize-prices/adaptive-pricing?payment-ui=stripe-hosted#restrictions).",
					},
				},
			},
			"after_expiration": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set, provides configuration for actions to take if this Checkout Session expires.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"recovery": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "When set, configuration used to recover the Checkout Session on expiry.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"allow_promotion_codes": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Enables user redeemable promotion codes on the recovered Checkout Sessions. Defaults to `false`",
							},
							"enabled": ephemeralSchema.BoolAttribute{
								Required:    true,
								Description: "If `true`, a recovery url will be generated to recover this Checkout Session if it\nexpires before a transaction is completed. It will be attached to the\nCheckout Session object upon expiration.",
							},
							"expires_at": ephemeralSchema.Int64Attribute{
								Computed:    true,
								Description: "The timestamp at which the recovery URL will expire.",
							},
							"url": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "URL that creates a new Checkout Session when clicked that is a copy of this expired Checkout Session",
							},
						},
					},
				},
			},
			"allow_promotion_codes": ephemeralSchema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables user redeemable promotion codes.",
			},
			"amount_subtotal": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "Total of all items before discounts or taxes are applied.",
			},
			"amount_total": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "Total of all items after discounts and taxes are applied.",
			},
			"automatic_tax": ephemeralSchema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"enabled": ephemeralSchema.BoolAttribute{
						Required:    true,
						Description: "Indicates whether automatic tax is enabled for the session",
					},
					"liability": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"account": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "The connected account being referenced when `type` is `account`.",
							},
							"type": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Type of the account referenced.",
								Validators:  []validator.String{stringvalidator.OneOf("account", "self")},
							},
						},
					},
					"provider": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The tax provider powering automatic tax.",
					},
					"status": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The status of the most recent automated tax calculation for this session.",
						Validators:  []validator.String{stringvalidator.OneOf("complete", "failed", "requires_location_inputs")},
					},
				},
			},
			"billing_address_collection": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Describes whether Checkout should collect the customer's billing address. Defaults to `auto`.",
				Validators:  []validator.String{stringvalidator.OneOf("auto", "required")},
			},
			"branding_settings": ephemeralSchema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"background_color": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "A hex color value starting with `#` representing the background color for the Checkout Session.",
					},
					"border_style": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The border style for the Checkout Session. Must be one of `rounded`, `rectangular`, or `pill`.",
						Validators:  []validator.String{stringvalidator.OneOf("pill", "rectangular", "rounded")},
					},
					"button_color": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "A hex color value starting with `#` representing the button color for the Checkout Session.",
					},
					"display_name": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The display name shown on the Checkout Session.",
					},
					"font_family": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The font family for the Checkout Session. Must be one of the [supported font families](https://docs.stripe.com/payments/checkout/customization/appearance?payment-ui=stripe-hosted#font-compatibility).",
					},
					"icon": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The icon for the Checkout Session. You cannot set both `logo` and `icon`.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"file": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "The ID of a [File upload](https://stripe.com/docs/api/files) representing the icon. Purpose must be `business_icon`. Required if `type` is `file` and disallowed otherwise.",
							},
							"type": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "The type of image for the icon. Must be one of `file` or `url`.",
								Validators:  []validator.String{stringvalidator.OneOf("file", "url")},
							},
							"url": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "The URL of the image. Present when `type` is `url`.",
							},
						},
					},
					"logo": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The logo for the Checkout Session. You cannot set both `logo` and `icon`.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"file": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "The ID of a [File upload](https://stripe.com/docs/api/files) representing the logo. Purpose must be `business_logo`. Required if `type` is `file` and disallowed otherwise.",
							},
							"type": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "The type of image for the logo. Must be one of `file` or `url`.",
								Validators:  []validator.String{stringvalidator.OneOf("file", "url")},
							},
							"url": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "The URL of the image. Present when `type` is `url`.",
							},
						},
					},
				},
			},
			"cancel_url": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If set, Checkout displays a back button and customers will be directed to this URL if they decide to cancel payment and return to your website.",
			},
			"client_reference_id": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A unique string to reference the Checkout Session. This can be a\ncustomer ID, a cart ID, or similar, and can be used to reconcile the\nSession with your internal systems.",
			},
			"client_secret": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The client secret of your Checkout Session. Applies to Checkout Sessions with `ui_mode: embedded_page` or `ui_mode: elements`. For `ui_mode: embedded_page`, the client secret is to be used when initializing Stripe.js embedded checkout.\n For `ui_mode: elements`, use the client secret with [initCheckout](https://docs.stripe.com/js/custom_checkout/init) on your front end.",
				Sensitive:   true,
			},
			"collected_information": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Information about the customer collected within the Checkout Session.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"business_name": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "Customer’s business name for this Checkout Session",
					},
					"individual_name": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "Customer’s individual name for this Checkout Session",
					},
					"shipping_details": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Shipping information for this Checkout Session.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"address": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"city": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "City, district, suburb, town, or village.",
									},
									"country": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
									},
									"line1": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Address line 1, such as the street, PO Box, or company name.",
									},
									"line2": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Address line 2, such as the apartment, suite, unit, or building.",
									},
									"postal_code": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "ZIP or postal code.",
									},
									"state": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
									},
								},
							},
							"name": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Customer name.",
							},
						},
					},
				},
			},
			"consent": ephemeralSchema.SingleNestedAttribute{
				Computed:    true,
				Description: "Results of `consent_collection` for this session.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"promotions": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "If `opt_in`, the customer consents to receiving promotional communications\nfrom the merchant about this Checkout Session.",
						Validators:  []validator.String{stringvalidator.OneOf("opt_in", "opt_out")},
					},
					"terms_of_service": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "If `accepted`, the customer in this Checkout Session has agreed to the merchant's terms of service.",
						Validators:  []validator.String{stringvalidator.OneOf("accepted")},
					},
				},
			},
			"consent_collection": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set, provides configuration for the Checkout Session to gather active consent from customers.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"payment_method_reuse_agreement": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "If set to `hidden`, it will hide legal text related to the reuse of a payment method.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"position": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Determines the position and visibility of the payment method reuse agreement in the UI. When set to `auto`, Stripe's defaults will be used.\n\nWhen set to `hidden`, the payment method reuse agreement text will always be hidden in the UI.",
								Validators:  []validator.String{stringvalidator.OneOf("auto", "hidden")},
							},
						},
					},
					"promotions": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout\nSession will determine whether to display an option to opt into promotional communication\nfrom the merchant depending on the customer's locale. Only available to US merchants and US customers.",
						Validators:  []validator.String{stringvalidator.OneOf("auto", "none")},
					},
					"terms_of_service": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "If set to `required`, it requires customers to accept the terms of service before being able to pay.",
						Validators:  []validator.String{stringvalidator.OneOf("none", "required")},
					},
				},
			},
			"created": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was created. Measured in seconds since the Unix epoch.",
			},
			"currency": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
			},
			"currency_conversion": ephemeralSchema.SingleNestedAttribute{
				Computed:    true,
				Description: "Currency conversion details for [Adaptive Pricing](https://docs.stripe.com/payments/checkout/adaptive-pricing) sessions created before 2025-03-31.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"amount_subtotal": ephemeralSchema.Int64Attribute{
						Computed:    true,
						Description: "Total of all items in source currency before discounts or taxes are applied.",
					},
					"amount_total": ephemeralSchema.Int64Attribute{
						Computed:    true,
						Description: "Total of all items in source currency after discounts and taxes are applied.",
					},
					"fx_rate": ephemeralSchema.Float64Attribute{
						Computed:    true,
						Description: "Exchange rate used to convert source currency amounts to customer currency amounts",
					},
					"source_currency": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "Creation currency of the CheckoutSession before localization",
					},
				},
			},
			"custom_fields": ephemeralSchema.ListNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.",
				NestedObject: ephemeralSchema.NestedAttributeObject{
					Attributes: map[string]ephemeralSchema.Attribute{
						"dropdown": ephemeralSchema.SingleNestedAttribute{
							Optional: true,
							Computed: true,

							Attributes: map[string]ephemeralSchema.Attribute{
								"default_value": ephemeralSchema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "The value that pre-fills on the payment page.",
								},
								"options": ephemeralSchema.ListNestedAttribute{
									Required:    true,
									Description: "The options available for the customer to select. Up to 200 options allowed.",
									NestedObject: ephemeralSchema.NestedAttributeObject{
										Attributes: map[string]ephemeralSchema.Attribute{
											"label": ephemeralSchema.StringAttribute{
												Required:    true,
												Description: "The label for the option, displayed to the customer. Up to 100 characters.",
											},
											"value": ephemeralSchema.StringAttribute{
												Required:    true,
												Description: "The value for this option, not displayed to the customer, used by your integration to reconcile the option selected by the customer. Must be unique to this option, alphanumeric, and up to 100 characters.",
											},
										},
									},
								},
								"value": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "The option selected by the customer. This will be the `value` for the option.",
								},
							},
						},
						"key": ephemeralSchema.StringAttribute{
							Required:    true,
							Description: "String of your choice that your integration can use to reconcile this field. Must be unique to this field, alphanumeric, and up to 200 characters.",
						},
						"label": ephemeralSchema.SingleNestedAttribute{
							Required: true,

							Attributes: map[string]ephemeralSchema.Attribute{
								"custom": ephemeralSchema.StringAttribute{
									Required:    true,
									Description: "Custom text for the label, displayed to the customer. Up to 50 characters.",
								},
								"type": ephemeralSchema.StringAttribute{
									Required:    true,
									Description: "The type of the label.",
									Validators:  []validator.String{stringvalidator.OneOf("custom")},
								},
							},
						},
						"numeric": ephemeralSchema.SingleNestedAttribute{
							Optional: true,
							Computed: true,

							Attributes: map[string]ephemeralSchema.Attribute{
								"default_value": ephemeralSchema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "The value that pre-fills the field on the payment page.",
								},
								"maximum_length": ephemeralSchema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "The maximum character length constraint for the customer's input.",
								},
								"minimum_length": ephemeralSchema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "The minimum character length requirement for the customer's input.",
								},
								"value": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "The value entered by the customer, containing only digits.",
								},
							},
						},
						"optional": ephemeralSchema.BoolAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.",
						},
						"text": ephemeralSchema.SingleNestedAttribute{
							Optional: true,
							Computed: true,

							Attributes: map[string]ephemeralSchema.Attribute{
								"default_value": ephemeralSchema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "The value that pre-fills the field on the payment page.",
								},
								"maximum_length": ephemeralSchema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "The maximum character length constraint for the customer's input.",
								},
								"minimum_length": ephemeralSchema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "The minimum character length requirement for the customer's input.",
								},
								"value": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "The value entered by the customer.",
								},
							},
						},
						"type": ephemeralSchema.StringAttribute{
							Required:    true,
							Description: "The type of the field.",
							Validators:  []validator.String{stringvalidator.OneOf("dropdown", "numeric", "text")},
						},
					},
				},
			},
			"custom_text": ephemeralSchema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"after_submit": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Custom text that should be displayed after the payment confirmation button.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"message": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Text can be up to 1200 characters in length.",
							},
						},
					},
					"shipping_address": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Custom text that should be displayed alongside shipping address collection.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"message": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Text can be up to 1200 characters in length.",
							},
						},
					},
					"submit": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Custom text that should be displayed alongside the payment confirmation button.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"message": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Text can be up to 1200 characters in length.",
							},
						},
					},
					"terms_of_service_acceptance": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Custom text that should be displayed in place of the default terms of service agreement text.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"message": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Text can be up to 1200 characters in length.",
							},
						},
					},
				},
			},
			"customer": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The ID of the customer for this Session.\nFor Checkout Sessions in `subscription` mode or Checkout Sessions with `customer_creation` set as `always` in `payment` mode, Checkout\nwill create a new customer object based on information provided\nduring the payment flow unless an existing customer was provided when\nthe Session was created.",
			},
			"customer_account": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The ID of the account for this Session.",
			},
			"customer_creation": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure whether a Checkout Session creates a Customer when the Checkout Session completes.",
				Validators:  []validator.String{stringvalidator.OneOf("always", "if_required")},
			},
			"customer_details": ephemeralSchema.SingleNestedAttribute{
				Computed:    true,
				Description: "The customer details including the customer's tax exempt status and the customer's tax IDs. Customer's address details are not present on Sessions in `setup` mode.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"address": ephemeralSchema.SingleNestedAttribute{
						Computed:    true,
						Description: "The customer's address after a completed Checkout Session. Note: This property is populated only for sessions on or after March 30, 2022.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"city": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "City, district, suburb, town, or village.",
							},
							"country": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
							},
							"line1": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "Address line 1, such as the street, PO Box, or company name.",
							},
							"line2": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "Address line 2, such as the apartment, suite, unit, or building.",
							},
							"postal_code": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "ZIP or postal code.",
							},
							"state": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
							},
						},
					},
					"business_name": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The customer's business name after a completed Checkout Session.",
					},
					"email": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The email associated with the Customer, if one exists, on the Checkout Session after a completed Checkout Session or at time of session expiry.\nOtherwise, if the customer has consented to promotional content, this value is the most recent valid email provided by the customer on the Checkout form.",
					},
					"individual_name": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The customer's individual name after a completed Checkout Session.",
					},
					"name": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The customer's name after a completed Checkout Session. Note: This property is populated only for sessions on or after March 30, 2022.",
					},
					"phone": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The customer's phone number after a completed Checkout Session.",
					},
					"tax_exempt": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The customer’s tax exempt status after a completed Checkout Session.",
						Validators:  []validator.String{stringvalidator.OneOf("exempt", "none", "reverse")},
					},
					"tax_ids": ephemeralSchema.ListNestedAttribute{
						Computed:    true,
						Description: "The customer’s tax IDs after a completed Checkout Session.",
						NestedObject: ephemeralSchema.NestedAttributeObject{
							Attributes: map[string]ephemeralSchema.Attribute{
								"type": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "The type of the tax ID, one of `ad_nrt`, `ar_cuit`, `eu_vat`, `bo_tin`, `br_cnpj`, `br_cpf`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `eu_oss_vat`, `hr_oib`, `pe_ruc`, `ro_tin`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, `vn_tin`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `no_voec`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `pl_nip`, `it_cf`, `fo_vat`, `gi_tin`, `py_ruc`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `li_vat`, `lk_vat`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, `al_tin`, `bh_vat`, `kz_bin`, `ng_tin`, `om_vat`, `de_stn`, `ch_uid`, `tz_vat`, `uz_vat`, `uz_tin`, `md_vat`, `ma_vat`, `by_tin`, `ao_tin`, `bs_tin`, `bb_tin`, `cd_nif`, `mr_nif`, `me_pib`, `zw_tin`, `ba_tin`, `gn_nif`, `mk_vat`, `sr_fin`, `sn_ninea`, `am_tin`, `np_pan`, `tj_tin`, `ug_tin`, `zm_tin`, `kh_tin`, `aw_tin`, `az_tin`, `bd_bin`, `bj_ifu`, `et_tin`, `kg_tin`, `la_tin`, `cm_niu`, `cv_nif`, `bf_ifu`, or `unknown`",
									Validators:  []validator.String{stringvalidator.OneOf("ad_nrt", "ae_trn", "al_tin", "am_tin", "ao_tin", "ar_cuit", "au_abn", "au_arn", "aw_tin", "az_tin", "ba_tin", "bb_tin", "bd_bin", "bf_ifu", "bg_uic", "bh_vat", "bj_ifu", "bo_tin", "br_cnpj", "br_cpf", "bs_tin", "by_tin", "ca_bn", "ca_gst_hst", "ca_pst_bc", "ca_pst_mb", "ca_pst_sk", "ca_qst", "cd_nif", "ch_uid", "ch_vat", "cl_tin", "cm_niu", "cn_tin", "co_nit", "cr_tin", "cv_nif", "de_stn", "do_rcn", "ec_ruc", "eg_tin", "es_cif", "et_tin", "eu_oss_vat", "eu_vat", "fo_vat", "gb_vat", "ge_vat", "gi_tin", "gn_nif", "hk_br", "hr_oib", "hu_tin", "id_npwp", "il_vat", "in_gst", "is_vat", "it_cf", "jp_cn", "jp_rn", "jp_trn", "ke_pin", "kg_tin", "kh_tin", "kr_brn", "kz_bin", "la_tin", "li_uid", "li_vat", "lk_vat", "ma_vat", "md_vat", "me_pib", "mk_vat", "mr_nif", "mx_rfc", "my_frp", "my_itn", "my_sst", "ng_tin", "no_vat", "no_voec", "np_pan", "nz_gst", "om_vat", "pe_ruc", "ph_tin", "pl_nip", "py_ruc", "ro_tin", "rs_pib", "ru_inn", "ru_kpp", "sa_vat", "sg_gst", "sg_uen", "si_tin", "sn_ninea", "sr_fin", "sv_nit", "th_vat", "tj_tin", "tr_tin", "tw_vat", "tz_vat", "ua_vat", "ug_tin", "unknown", "us_ein", "uy_ruc", "uz_tin", "uz_vat", "ve_rif", "vn_tin", "za_vat", "zm_tin", "zw_tin")},
								},
								"value": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "The value of the tax ID.",
								},
							},
						},
					},
				},
			},
			"customer_email": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If provided, this value will be used when the Customer object is created.\nIf not provided, customers will be asked to enter their email address.\nUse this parameter to prefill customer data if you already have an email\non file. To access information about the customer once the payment flow is\ncomplete, use the `customer` attribute.",
			},
			"discounts": ephemeralSchema.ListNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of coupons and promotion codes attached to the Checkout Session.",
				NestedObject: ephemeralSchema.NestedAttributeObject{
					Attributes: map[string]ephemeralSchema.Attribute{
						"coupon": ephemeralSchema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Coupon attached to the Checkout Session.",
						},
						"promotion_code": ephemeralSchema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Promotion code attached to the Checkout Session.",
						},
					},
				},
			},
			"excluded_payment_method_types": ephemeralSchema.ListAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A list of the types of payment methods (e.g., `card`) that should be excluded from this Checkout Session. This should only be used when payment methods for this Checkout Session are managed through the [Stripe Dashboard](https://dashboard.stripe.com/settings/payment_methods).",
				ElementType: types.StringType,
			},
			"expires_at": ephemeralSchema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The timestamp at which the Checkout Session will expire.",
			},
			"id": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier for the object.",
			},
			"integration_identifier": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The integration identifier for this Checkout Session. Multiple Checkout Sessions can have the same integration identifier.",
			},
			"invoice": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "ID of the invoice created by the Checkout Session, if it exists.",
			},
			"invoice_creation": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Details on the state of invoice creation for the Checkout Session.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"enabled": ephemeralSchema.BoolAttribute{
						Required:    true,
						Description: "Indicates whether invoice creation is enabled for the Checkout Session.",
					},
					"invoice_data": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"account_tax_ids": ephemeralSchema.ListAttribute{
								Optional:    true,
								Computed:    true,
								Description: "The account tax IDs associated with the invoice.",
								ElementType: types.StringType,
							},
							"custom_fields": ephemeralSchema.ListNestedAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Custom fields displayed on the invoice.",
								NestedObject: ephemeralSchema.NestedAttributeObject{
									Attributes: map[string]ephemeralSchema.Attribute{
										"name": ephemeralSchema.StringAttribute{
											Required:    true,
											Description: "The name of the custom field.",
										},
										"value": ephemeralSchema.StringAttribute{
											Required:    true,
											Description: "The value of the custom field.",
										},
									},
								},
							},
							"description": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "An arbitrary string attached to the object. Often useful for displaying to users.",
							},
							"footer": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Footer displayed on the invoice.",
							},
							"issuer": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Computed:    true,
								Description: "The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"account": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The connected account being referenced when `type` is `account`.",
									},
									"type": ephemeralSchema.StringAttribute{
										Required:    true,
										Description: "Type of the account referenced.",
										Validators:  []validator.String{stringvalidator.OneOf("account", "self")},
									},
								},
							},
							"metadata": ephemeralSchema.MapAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
								ElementType: types.StringType,
							},
							"rendering_options": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Options for invoice PDF rendering.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"amount_tax_display": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.",
									},
									"template": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "ID of the invoice rendering template to be used for the generated invoice.",
									},
								},
							},
						},
					},
				},
			},
			"line_items": ephemeralSchema.ListNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The line items purchased by the customer.",
				NestedObject: ephemeralSchema.NestedAttributeObject{
					Attributes: map[string]ephemeralSchema.Attribute{
						"adjustable_quantity": ephemeralSchema.SingleNestedAttribute{
							Optional:    true,
							Computed:    true,
							Description: "When set, provides configuration for this item’s quantity to be adjusted by the customer during Checkout.",
							Attributes: map[string]ephemeralSchema.Attribute{
								"enabled": ephemeralSchema.BoolAttribute{
									Required:    true,
									Description: "Set to true if the quantity can be adjusted to any non-negative integer.",
								},
								"maximum": ephemeralSchema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "The maximum quantity the customer can purchase for the Checkout Session. By default this value is 99. You can specify a value up to 999999.",
								},
								"minimum": ephemeralSchema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "The minimum quantity the customer must purchase for the Checkout Session. By default this value is 0.",
								},
							},
						},
						"dynamic_tax_rates": ephemeralSchema.ListAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The [tax rates](https://docs.stripe.com/api/tax_rates) that will be applied to this line item depending on the customer's billing/shipping address. We currently support the following countries: US, GB, AU, and all countries in the EU. You can't set this parameter if `ui_mode` is `custom`.",
							ElementType: types.StringType,
						},
						"metadata": ephemeralSchema.MapAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
							ElementType: types.StringType,
						},
						"price": ephemeralSchema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object. One of `price` or `price_data` is required.",
						},
						"price_data": ephemeralSchema.SingleNestedAttribute{
							Optional:    true,
							Computed:    true,
							Description: "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.",
							Attributes: map[string]ephemeralSchema.Attribute{
								"currency": ephemeralSchema.StringAttribute{
									Required:    true,
									Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
								},
								"product": ephemeralSchema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to. One of `product` or `product_data` is required.",
								},
								"product_data": ephemeralSchema.SingleNestedAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.",
									Attributes: map[string]ephemeralSchema.Attribute{
										"description": ephemeralSchema.StringAttribute{
											Optional:    true,
											Computed:    true,
											Description: "The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.",
										},
										"images": ephemeralSchema.ListAttribute{
											Optional:    true,
											Computed:    true,
											Description: "A list of up to 8 URLs of images for this product, meant to be displayable to the customer.",
											ElementType: types.StringType,
										},
										"metadata": ephemeralSchema.MapAttribute{
											Optional:    true,
											Computed:    true,
											Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
											ElementType: types.StringType,
										},
										"name": ephemeralSchema.StringAttribute{
											Required:    true,
											Description: "The product's name, meant to be displayable to the customer.",
										},
										"tax_code": ephemeralSchema.StringAttribute{
											Optional:    true,
											Computed:    true,
											Description: "A [tax code](https://docs.stripe.com/tax/tax-categories) ID.",
										},
										"unit_label": ephemeralSchema.StringAttribute{
											Optional:    true,
											Computed:    true,
											Description: "A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.",
										},
									},
								},
								"recurring": ephemeralSchema.SingleNestedAttribute{
									Optional:    true,
									Computed:    true,
									Description: "The recurring components of a price such as `interval` and `interval_count`.",
									Attributes: map[string]ephemeralSchema.Attribute{
										"interval": ephemeralSchema.StringAttribute{
											Required:    true,
											Description: "Specifies billing frequency. Either `day`, `week`, `month` or `year`.",
										},
										"interval_count": ephemeralSchema.Int64Attribute{
											Optional:    true,
											Computed:    true,
											Description: "The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).",
										},
									},
								},
								"tax_behavior": ephemeralSchema.StringAttribute{
									Optional:    true,
									Computed:    true,
									Description: "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
								},
								"unit_amount": ephemeralSchema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "A non-negative integer in cents (or local equivalent) representing how much to charge. One of `unit_amount` or `unit_amount_decimal` is required.",
								},
								"unit_amount_decimal": ephemeralSchema.Float64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
								},
							},
						},
						"quantity": ephemeralSchema.Int64Attribute{
							Optional:    true,
							Computed:    true,
							Description: "The quantity of the line item being purchased. Quantity should not be defined when `recurring.usage_type=metered`.",
						},
						"tax_rates": ephemeralSchema.ListAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The [tax rates](https://docs.stripe.com/api/tax_rates) which apply to this line item.",
							ElementType: types.StringType,
						},
					},
				},
			},
			"livemode": ephemeralSchema.BoolAttribute{
				Computed:    true,
				Description: "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
			},
			"locale": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IETF language tag of the locale Checkout is displayed in. If blank or `auto`, the browser's locale is used.",
				Validators:  []validator.String{stringvalidator.OneOf("auto", "bg", "cs", "da", "de", "el", "en", "en-GB", "es", "es-419", "et", "fi", "fil", "fr", "fr-CA", "hr", "hu", "id", "it", "ja", "ko", "lt", "lv", "ms", "mt", "nb", "nl", "pl", "pt", "pt-BR", "ro", "ru", "sk", "sl", "sv", "th", "tr", "vi", "zh", "zh-HK", "zh-TW")},
			},
			"managed_payments": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Settings for Managed Payments for this Checkout Session and resulting [PaymentIntents](/api/payment_intents/object), [Invoices](/api/invoices/object), and [Subscriptions](/api/subscriptions/object).",
				Attributes: map[string]ephemeralSchema.Attribute{
					"enabled": ephemeralSchema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Set to `true` to enable [Managed Payments](https://docs.stripe.com/payments/managed-payments), Stripe's merchant of record solution, for this session.",
					},
				},
			},
			"metadata": ephemeralSchema.MapAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"mode": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The mode of the Checkout Session.",
				Validators:  []validator.String{stringvalidator.OneOf("payment", "setup", "subscription")},
			},
			"name_collection": ephemeralSchema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"business": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"enabled": ephemeralSchema.BoolAttribute{
								Required:    true,
								Description: "Indicates whether business name collection is enabled for the session",
							},
							"optional": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.",
							},
						},
					},
					"individual": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"enabled": ephemeralSchema.BoolAttribute{
								Required:    true,
								Description: "Indicates whether individual name collection is enabled for the session",
							},
							"optional": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.",
							},
						},
					},
				},
			},
			"optional_items": ephemeralSchema.ListNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The optional items presented to the customer at checkout.",
				NestedObject: ephemeralSchema.NestedAttributeObject{
					Attributes: map[string]ephemeralSchema.Attribute{
						"adjustable_quantity": ephemeralSchema.SingleNestedAttribute{
							Optional: true,
							Computed: true,

							Attributes: map[string]ephemeralSchema.Attribute{
								"enabled": ephemeralSchema.BoolAttribute{
									Required:    true,
									Description: "Set to true if the quantity can be adjusted to any non-negative integer.",
								},
								"maximum": ephemeralSchema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "The maximum quantity of this item the customer can purchase. By default this value is 99. You can specify a value up to 999999.",
								},
								"minimum": ephemeralSchema.Int64Attribute{
									Optional:    true,
									Computed:    true,
									Description: "The minimum quantity of this item the customer must purchase, if they choose to purchase it. Because this item is optional, the customer will always be able to remove it from their order, even if the `minimum` configured here is greater than 0. By default this value is 0.",
								},
							},
						},
						"price": ephemeralSchema.StringAttribute{
							Required: true,
						},
						"quantity": ephemeralSchema.Int64Attribute{
							Required: true,
						},
					},
				},
			},
			"origin_context": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Where the user is coming from. This informs the optimizations that are applied to the session.",
				Validators:  []validator.String{stringvalidator.OneOf("mobile_app", "web")},
			},
			"payment_intent": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The ID of the PaymentIntent for Checkout Sessions in `payment` mode. You can't confirm or cancel the PaymentIntent for a Checkout Session. To cancel, [expire the Checkout Session](https://docs.stripe.com/api/checkout/sessions/expire) instead.",
			},
			"payment_link": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The ID of the Payment Link that created this Session.",
			},
			"payment_method_collection": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure whether a Checkout Session should collect a payment method. Defaults to `always`.",
				Validators:  []validator.String{stringvalidator.OneOf("always", "if_required")},
			},
			"payment_method_configuration_details": ephemeralSchema.SingleNestedAttribute{
				Computed:    true,
				Description: "Information about the payment method configuration used for this Checkout session if using dynamic payment methods.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"id": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "ID of the payment method configuration used.",
					},
					"parent": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "ID of the parent payment method configuration used.",
					},
				},
			},
			"payment_method_options": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Payment-method-specific configuration for the PaymentIntent or SetupIntent of this CheckoutSession.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"acss_debit": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"currency": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Currency supported by the bank account. Returned when the Session is in `setup` mode.",
								Validators:  []validator.String{stringvalidator.OneOf("cad", "usd")},
							},
							"mandate_options": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"custom_mandate_url": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "A URL for custom mandate text",
									},
									"default_for": ephemeralSchema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Description: "List of Stripe products where this mandate can be selected automatically. Returned when the Session is in `setup` mode.",
										ElementType: types.StringType,
									},
									"interval_description": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.",
									},
									"payment_schedule": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Payment schedule for the mandate.",
										Validators:  []validator.String{stringvalidator.OneOf("combined", "interval", "sporadic")},
									},
									"transaction_type": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Transaction type of the mandate.",
										Validators:  []validator.String{stringvalidator.OneOf("business", "personal")},
									},
								},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"target_date": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
							},
							"verification_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Bank account verification method. The default value is `automatic`.",
								Validators:  []validator.String{stringvalidator.OneOf("automatic", "instant", "microdeposits")},
							},
						},
					},
					"affirm": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"afterpay_clearpay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"alipay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"alma": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"amazon_pay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"au_becs_debit": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
							"target_date": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
							},
						},
					},
					"bacs_debit": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"mandate_options": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"reference_prefix": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'DDIC' or 'STRIPE'.",
									},
								},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"target_date": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
							},
						},
					},
					"bancontact": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"billie": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"boleto": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"expires_after_days": ephemeralSchema.Int64Attribute{
								Optional:    true,
								Computed:    true,
								Description: "The number of calendar days before a Boleto voucher expires. For example, if you create a Boleto voucher on Monday and you set expires_after_days to 2, the Boleto voucher will expire on Wednesday at 23:59 America/Sao_Paulo time.",
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
						},
					},
					"card": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"installments": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"enabled": ephemeralSchema.BoolAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Indicates if installments are enabled",
									},
								},
							},
							"request_extended_authorization": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Request ability to [capture beyond the standard authorization validity window](/payments/extended-authorization) for this CheckoutSession.",
								Validators:  []validator.String{stringvalidator.OneOf("if_available", "never")},
							},
							"request_incremental_authorization": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Request ability to [increment the authorization](/payments/incremental-authorization) for this CheckoutSession.",
								Validators:  []validator.String{stringvalidator.OneOf("if_available", "never")},
							},
							"request_multicapture": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Request ability to make [multiple captures](/payments/multicapture) for this CheckoutSession.",
								Validators:  []validator.String{stringvalidator.OneOf("if_available", "never")},
							},
							"request_overcapture": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Request ability to [overcapture](/payments/overcapture) for this CheckoutSession.",
								Validators:  []validator.String{stringvalidator.OneOf("if_available", "never")},
							},
							"request_three_d_secure": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. If not provided, this value defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.",
								Validators:  []validator.String{stringvalidator.OneOf("any", "automatic", "challenge")},
							},
							"restrictions": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"brands_blocked": ephemeralSchema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The card brands to block. If a customer enters or selects a card belonging to a blocked brand, they can't complete the payment.",
										ElementType: types.StringType,
									},
								},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"statement_descriptor_suffix_kana": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Provides information about a card payment that customers see on their statements. Concatenated with the Kana prefix (shortened Kana descriptor) or Kana statement descriptor that’s set on the account to form the complete statement descriptor. Maximum 22 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 22 characters.",
							},
							"statement_descriptor_suffix_kanji": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Provides information about a card payment that customers see on their statements. Concatenated with the Kanji prefix (shortened Kanji descriptor) or Kanji statement descriptor that’s set on the account to form the complete statement descriptor. Maximum 17 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 17 characters.",
							},
						},
					},
					"cashapp": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"customer_balance": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"bank_transfer": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"eu_bank_transfer": ephemeralSchema.SingleNestedAttribute{
										Optional: true,
										Computed: true,

										Attributes: map[string]ephemeralSchema.Attribute{
											"country": ephemeralSchema.StringAttribute{
												Required:    true,
												Description: "The desired country code of the bank account information. Permitted values include: `DE`, `FR`, `IE`, or `NL`.",
												Validators:  []validator.String{stringvalidator.OneOf("BE", "DE", "ES", "FR", "IE", "NL")},
											},
										},
									},
									"requested_address_types": ephemeralSchema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Description: "List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.\n\nPermitted values include: `sort_code`, `zengin`, `iban`, or `spei`.",
										ElementType: types.StringType,
									},
									"type": ephemeralSchema.StringAttribute{
										Required:    true,
										Description: "The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.",
										Validators:  []validator.String{stringvalidator.OneOf("eu_bank_transfer", "gb_bank_transfer", "jp_bank_transfer", "mx_bank_transfer", "us_bank_transfer")},
									},
								},
							},
							"funding_type": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.",
								Validators:  []validator.String{stringvalidator.OneOf("bank_transfer")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"eps": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"fpx": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"giropay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"grabpay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"ideal": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"kakao_pay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"klarna": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"subscriptions": ephemeralSchema.ListNestedAttribute{
								Optional:    true,
								Description: "Subscription details if the Checkout Session sets up a future subscription.",
								NestedObject: ephemeralSchema.NestedAttributeObject{
									Attributes: map[string]ephemeralSchema.Attribute{
										"interval": ephemeralSchema.StringAttribute{
											Required:    true,
											Description: "Unit of time between subscription charges.",
										},
										"interval_count": ephemeralSchema.Int64Attribute{
											Optional:    true,
											Description: "The number of intervals (specified in the `interval` attribute) between subscription charges. For example, `interval=month` and `interval_count=3` charges every 3 months.",
										},
										"name": ephemeralSchema.StringAttribute{
											Optional:    true,
											Description: "Name for subscription.",
										},
										"next_billing": ephemeralSchema.SingleNestedAttribute{
											Required:    true,
											Description: "Describes the upcoming charge for this subscription.",
											Attributes: map[string]ephemeralSchema.Attribute{
												"amount": ephemeralSchema.Int64Attribute{
													Required:    true,
													Description: "The amount of the next charge for the subscription.",
												},
												"date": ephemeralSchema.StringAttribute{
													Required:    true,
													Description: "The date of the next charge for the subscription in YYYY-MM-DD format.",
												},
											},
										},
										"reference": ephemeralSchema.StringAttribute{
											Required:    true,
											Description: "A non-customer-facing reference to correlate subscription charges in the Klarna app. Use a value that persists across subscription charges.",
										},
									},
								},
							},
						},
					},
					"konbini": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"expires_after_days": ephemeralSchema.Int64Attribute{
								Optional:    true,
								Computed:    true,
								Description: "The number of calendar days (between 1 and 60) after which Konbini payment instructions will expire. For example, if a PaymentIntent is confirmed with Konbini and `expires_after_days` set to 2 on Monday JST, the instructions will expire on Wednesday 23:59:59 JST.",
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"kr_card": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"link": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"mobilepay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"multibanco": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"naver_pay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"oxxo": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"expires_after_days": ephemeralSchema.Int64Attribute{
								Optional:    true,
								Computed:    true,
								Description: "The number of calendar days before an OXXO invoice expires. For example, if you create an OXXO invoice on Monday and you set expires_after_days to 2, the OXXO invoice will expire on Wednesday at 23:59 America/Mexico_City time.",
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"p24": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
							"tos_shown_and_accepted": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Confirm that the payer has accepted the P24 terms and conditions.",
							},
						},
					},
					"payco": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"paynow": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"paypal": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"preferred_locale": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Preferred locale of the PayPal checkout page that the customer is redirected to.",
							},
							"reference": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "A reference of the PayPal transaction visible to customer which is mapped to PayPal's invoice ID. This must be a globally unique ID if you have configured in your PayPal settings to block multiple payments per invoice ID.",
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
							"risk_correlation_id": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The risk correlation ID for an on-session payment using a saved PayPal payment method.",
							},
						},
					},
					"payto": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"mandate_options": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"amount": ephemeralSchema.Int64Attribute{
										Optional:    true,
										Computed:    true,
										Description: "Amount that will be collected. It is required when `amount_type` is `fixed`.",
									},
									"amount_type": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The type of amount that will be collected. The amount charged must be exact or up to the value of `amount` param for `fixed` or `maximum` type respectively. Defaults to `maximum`.",
										Validators:  []validator.String{stringvalidator.OneOf("fixed", "maximum")},
									},
									"end_date": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Date, in YYYY-MM-DD format, after which payments will not be collected. Defaults to no end date.",
									},
									"payment_schedule": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The periodicity at which payments will be collected. Defaults to `adhoc`.",
										Validators:  []validator.String{stringvalidator.OneOf("adhoc", "annual", "daily", "fortnightly", "monthly", "quarterly", "semi_annual", "weekly")},
									},
									"payments_per_period": ephemeralSchema.Int64Attribute{
										Optional:    true,
										Computed:    true,
										Description: "The number of payments that will be made during a payment period. Defaults to 1 except for when `payment_schedule` is `adhoc`. In that case, it defaults to no limit.",
									},
									"purpose": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The purpose for which payments are made. Has a default value based on your merchant category code.",
										Validators:  []validator.String{stringvalidator.OneOf("dependant_support", "government", "loan", "mortgage", "other", "pension", "personal", "retail", "salary", "tax", "utility")},
									},
									"start_date": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Date, in YYYY-MM-DD format, from which payments will be collected. Defaults to confirmation time.",
									},
								},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"pix": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"amount_includes_iof": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Determines if the amount includes the IOF tax.",
								Validators:  []validator.String{stringvalidator.OneOf("always", "never")},
							},
							"expires_after_seconds": ephemeralSchema.Int64Attribute{
								Optional:    true,
								Computed:    true,
								Description: "The number of seconds after which Pix payment will expire.",
							},
							"mandate_options": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"amount": ephemeralSchema.Int64Attribute{
										Optional:    true,
										Computed:    true,
										Description: "Amount to be charged for future payments.",
									},
									"amount_includes_iof": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Determines if the amount includes the IOF tax.",
										Validators:  []validator.String{stringvalidator.OneOf("always", "never")},
									},
									"amount_type": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Type of amount.",
										Validators:  []validator.String{stringvalidator.OneOf("fixed", "maximum")},
									},
									"currency": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase.",
									},
									"end_date": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Date when the mandate expires and no further payments will be charged, in `YYYY-MM-DD`.",
									},
									"payment_schedule": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Schedule at which the future payments will be charged.",
										Validators:  []validator.String{stringvalidator.OneOf("halfyearly", "monthly", "quarterly", "weekly", "yearly")},
									},
									"reference": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Subscription name displayed to buyers in their bank app.",
									},
									"start_date": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Start date of the mandate, in `YYYY-MM-DD`.",
									},
								},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"revolut_pay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"samsung_pay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"satispay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"scalapay": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"capture_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when the funds will be captured from the customer's account.",
								Validators:  []validator.String{stringvalidator.OneOf("manual")},
							},
						},
					},
					"sepa_debit": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"mandate_options": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"reference_prefix": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Prefix used to generate the Mandate reference. Must be at most 12 characters long. Must consist of only uppercase letters, numbers, spaces, or the following special characters: '/', '_', '-', '&', '.'. Cannot begin with 'STRIPE'.",
									},
								},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"target_date": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
							},
						},
					},
					"sofort": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none")},
							},
						},
					},
					"swish": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"reference": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "The order reference that will be displayed to customers in the Swish application. Defaults to the `id` of the Payment Intent.",
							},
						},
					},
					"twint": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session")},
							},
						},
					},
					"upi": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"mandate_options": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"amount": ephemeralSchema.Int64Attribute{
										Optional:    true,
										Computed:    true,
										Description: "Amount to be charged for future payments.",
									},
									"amount_type": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.",
										Validators:  []validator.String{stringvalidator.OneOf("fixed", "maximum")},
									},
									"description": ephemeralSchema.StringAttribute{
										Optional:    true,
										Computed:    true,
										Description: "A description of the mandate or subscription that is meant to be displayed to the customer.",
									},
									"end_date": ephemeralSchema.Int64Attribute{
										Optional:    true,
										Computed:    true,
										Description: "End date of the mandate or subscription.",
									},
								},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
						},
					},
					"us_bank_account": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"financial_connections": ephemeralSchema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								Attributes: map[string]ephemeralSchema.Attribute{
									"filters": ephemeralSchema.SingleNestedAttribute{
										Computed: true,

										Attributes: map[string]ephemeralSchema.Attribute{
											"account_subcategories": ephemeralSchema.ListAttribute{
												Computed:    true,
												Description: "The account subcategories to use to filter for possible accounts to link. Valid subcategories are `checking` and `savings`.",
												ElementType: types.StringType,
											},
										},
									},
									"permissions": ephemeralSchema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Description: "The list of permissions to request. The `payment_method` permission must be included.",
										ElementType: types.StringType,
									},
									"prefetch": ephemeralSchema.ListAttribute{
										Optional:    true,
										Computed:    true,
										Description: "Data features requested to be retrieved upon account creation.",
										ElementType: types.StringType,
									},
									"return_url": ephemeralSchema.StringAttribute{
										Computed:    true,
										Description: "For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.",
									},
								},
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
								Validators:  []validator.String{stringvalidator.OneOf("none", "off_session", "on_session")},
							},
							"target_date": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Controls when Stripe will attempt to debit the funds from the customer's account. The date must be a string in YYYY-MM-DD format. The date must be in the future and between 3 and 15 calendar days from now.",
							},
							"verification_method": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Bank account verification method. The default value is `automatic`.",
								Validators:  []validator.String{stringvalidator.OneOf("automatic", "instant")},
							},
						},
					},
					"crypto": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "contains details about the Crypto payment method options.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
							},
						},
					},
					"demo_pay": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "contains details about the DemoPay payment method options.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
							},
						},
					},
					"wechat_pay": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "contains details about the WeChat Pay payment method options.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"app_id": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The app ID registered with WeChat Pay. Only required when client is ios or android.",
							},
							"client": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "The client type that the end customer will pay from",
							},
							"setup_future_usage": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Indicates that you intend to make future payments with this PaymentIntent's payment method.\n\nIf you provide a Customer with the PaymentIntent, you can use this parameter to [attach the payment method](/payments/save-during-payment) to the Customer after the PaymentIntent is confirmed and the customer completes any required actions. If you don't provide a Customer, you can still [attach](/api/payment_methods/attach) the payment method to a Customer after the transaction completes.\n\nIf the payment method is `card_present` and isn't a digital wallet, Stripe creates and attaches a [generated_card](/api/charges/object#charge_object-payment_method_details-card_present-generated_card) payment method representing the card to the Customer instead.\n\nWhen processing card payments, Stripe uses `setup_future_usage` to help you comply with regional legislation and network rules, such as [SCA](/strong-customer-authentication).",
							},
						},
					},
				},
			},
			"payment_method_types": ephemeralSchema.ListAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A list of the types of payment methods (e.g. card) this Checkout\nSession is allowed to accept.",
				ElementType: types.StringType,
			},
			"payment_status": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The payment status of the Checkout Session, one of `paid`, `unpaid`, or `no_payment_required`.\nYou can use this value to decide when to fulfill your customer's order.",
				Validators:  []validator.String{stringvalidator.OneOf("no_payment_required", "paid", "unpaid")},
			},
			"permissions": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This property is used to set up permissions for various actions (e.g., update) on the CheckoutSession object.\n\nFor specific permissions, please refer to their dedicated subsections, such as `permissions.update_shipping_details`.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"update_shipping_details": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Determines which entity is allowed to update the shipping details.\n\nDefault is `client_only`. Stripe Checkout client will automatically update the shipping details. If set to `server_only`, only your server is allowed to update the shipping details.\n\nWhen set to `server_only`, you must add the onShippingDetailsChange event handler when initializing the Stripe Checkout client and manually update the shipping details from your server using the Stripe API.",
						Validators:  []validator.String{stringvalidator.OneOf("client_only", "server_only")},
					},
				},
			},
			"phone_number_collection": ephemeralSchema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"enabled": ephemeralSchema.BoolAttribute{
						Required:    true,
						Description: "Indicates whether phone number collection is enabled for the session",
					},
				},
			},
			"presentment_details": ephemeralSchema.SingleNestedAttribute{
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"presentment_amount": ephemeralSchema.Int64Attribute{
						Computed:    true,
						Description: "Amount intended to be collected by this payment, denominated in `presentment_currency`.",
					},
					"presentment_currency": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "Currency presented to the customer during payment.",
					},
				},
			},
			"recovered_from": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The ID of the original expired Checkout Session that triggered the recovery flow.",
			},
			"redirect_on_completion": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This parameter applies to `ui_mode: embedded_page`. Learn more about the [redirect behavior](https://docs.stripe.com/payments/checkout/custom-success-page?payment-ui=embedded-form) of embedded sessions. Defaults to `always`.",
				Validators:  []validator.String{stringvalidator.OneOf("always", "if_required", "never")},
			},
			"return_url": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applies to Checkout Sessions with `ui_mode: embedded_page` or `ui_mode: elements`. The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method's app or site.",
			},
			"saved_payment_method_options": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Controls saved payment method settings for the session. Only available in `payment` and `subscription` mode.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"allow_redisplay_filters": ephemeralSchema.ListAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Uses the `allow_redisplay` value of each saved payment method to filter the set presented to a returning customer. By default, only saved payment methods with ’allow_redisplay: ‘always’ are shown in Checkout.",
						ElementType: types.StringType,
					},
					"payment_method_remove": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Enable customers to choose if they wish to remove their saved payment methods. Disabled by default.",
						Validators:  []validator.String{stringvalidator.OneOf("disabled", "enabled")},
					},
					"payment_method_save": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Enable customers to choose if they wish to save their payment method for future use. Disabled by default.",
						Validators:  []validator.String{stringvalidator.OneOf("disabled", "enabled")},
					},
				},
			},
			"setup_intent": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The ID of the SetupIntent for Checkout Sessions in `setup` mode. You can't confirm or cancel the SetupIntent for a Checkout Session. To cancel, [expire the Checkout Session](https://docs.stripe.com/api/checkout/sessions/expire) instead.",
			},
			"shipping_address_collection": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set, provides configuration for Checkout to collect a shipping address from a customer.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"allowed_countries": ephemeralSchema.ListAttribute{
						Required:    true,
						Description: "An array of two-letter ISO country codes representing which countries Checkout should provide as options for\nshipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SY, UM, VI`.",
						ElementType: types.StringType,
					},
				},
			},
			"shipping_cost": ephemeralSchema.SingleNestedAttribute{
				Computed:    true,
				Description: "The details of the customer cost of shipping, including the customer chosen ShippingRate.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"amount_subtotal": ephemeralSchema.Int64Attribute{
						Computed:    true,
						Description: "Total shipping cost before any discounts or taxes are applied.",
					},
					"amount_tax": ephemeralSchema.Int64Attribute{
						Computed:    true,
						Description: "Total tax amount applied due to shipping costs. If no tax was applied, defaults to 0.",
					},
					"amount_total": ephemeralSchema.Int64Attribute{
						Computed:    true,
						Description: "Total shipping cost after discounts and taxes are applied.",
					},
					"shipping_rate": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The ID of the ShippingRate for this order.",
					},
					"taxes": ephemeralSchema.ListNestedAttribute{
						Computed:    true,
						Description: "The taxes applied to the shipping rate.",
						NestedObject: ephemeralSchema.NestedAttributeObject{
							Attributes: map[string]ephemeralSchema.Attribute{
								"amount": ephemeralSchema.Int64Attribute{
									Computed:    true,
									Description: "Amount of tax applied for this rate.",
								},
								"rate": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "Tax rates can be applied to [invoices](/invoicing/taxes/tax-rates), [subscriptions](/billing/taxes/tax-rates) and [Checkout Sessions](/payments/checkout/use-manual-tax-rates) to collect tax.\n\nRelated guide: [Tax rates](/billing/taxes/tax-rates)",
								},
								"taxability_reason": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.",
									Validators:  []validator.String{stringvalidator.OneOf("customer_exempt", "not_collecting", "not_subject_to_tax", "not_supported", "portion_product_exempt", "portion_reduced_rated", "portion_standard_rated", "product_exempt", "product_exempt_holiday", "proportionally_rated", "reduced_rated", "reverse_charge", "standard_rated", "taxable_basis_reduced", "zero_rated")},
								},
								"taxable_amount": ephemeralSchema.Int64Attribute{
									Computed:    true,
									Description: "The amount on which tax is calculated, in cents (or local equivalent).",
								},
							},
						},
					},
				},
			},
			"shipping_options": ephemeralSchema.ListNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The shipping rate options applied to this Session.",
				NestedObject: ephemeralSchema.NestedAttributeObject{
					Attributes: map[string]ephemeralSchema.Attribute{
						"shipping_amount": ephemeralSchema.Int64Attribute{
							Computed:    true,
							Description: "A non-negative integer in cents representing how much to charge.",
						},
						"shipping_rate": ephemeralSchema.StringAttribute{
							Optional:    true,
							Computed:    true,
							Description: "The shipping rate.",
						},
						"shipping_rate_data": ephemeralSchema.SingleNestedAttribute{
							Optional:    true,
							Description: "Parameters to be passed to Shipping Rate creation for this shipping option.",
							Attributes: map[string]ephemeralSchema.Attribute{
								"delivery_estimate": ephemeralSchema.SingleNestedAttribute{
									Optional:    true,
									Description: "The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.",
									Attributes: map[string]ephemeralSchema.Attribute{
										"maximum": ephemeralSchema.SingleNestedAttribute{
											Optional:    true,
											Description: "The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.",
											Attributes: map[string]ephemeralSchema.Attribute{
												"unit": ephemeralSchema.StringAttribute{
													Required:    true,
													Description: "A unit of time.",
												},
												"value": ephemeralSchema.Int64Attribute{
													Required:    true,
													Description: "Must be greater than 0.",
												},
											},
										},
										"minimum": ephemeralSchema.SingleNestedAttribute{
											Optional:    true,
											Description: "The lower bound of the estimated range. If empty, represents no lower bound.",
											Attributes: map[string]ephemeralSchema.Attribute{
												"unit": ephemeralSchema.StringAttribute{
													Required:    true,
													Description: "A unit of time.",
												},
												"value": ephemeralSchema.Int64Attribute{
													Required:    true,
													Description: "Must be greater than 0.",
												},
											},
										},
									},
								},
								"display_name": ephemeralSchema.StringAttribute{
									Required:    true,
									Description: "The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.",
								},
								"fixed_amount": ephemeralSchema.SingleNestedAttribute{
									Optional:    true,
									Description: "Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.",
									Attributes: map[string]ephemeralSchema.Attribute{
										"amount": ephemeralSchema.Int64Attribute{
											Required:    true,
											Description: "A non-negative integer in cents representing how much to charge.",
										},
										"currency": ephemeralSchema.StringAttribute{
											Required:    true,
											Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
										},
										"currency_options": ephemeralSchema.MapNestedAttribute{
											Optional:    true,
											Description: "Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
											NestedObject: ephemeralSchema.NestedAttributeObject{
												Attributes: map[string]ephemeralSchema.Attribute{
													"amount": ephemeralSchema.Int64Attribute{
														Required:    true,
														Description: "A non-negative integer in cents representing how much to charge.",
													},
													"tax_behavior": ephemeralSchema.StringAttribute{
														Optional:    true,
														Description: "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.",
													},
												},
											},
										},
									},
								},
								"metadata": ephemeralSchema.MapAttribute{
									Optional:    true,
									Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
									ElementType: types.StringType,
								},
								"tax_behavior": ephemeralSchema.StringAttribute{
									Optional:    true,
									Description: "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.",
								},
								"tax_code": ephemeralSchema.StringAttribute{
									Optional:    true,
									Description: "A [tax code](https://docs.stripe.com/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.",
								},
								"type": ephemeralSchema.StringAttribute{
									Optional:    true,
									Description: "The type of calculation to use on the shipping rate.",
								},
							},
						},
					},
				},
			},
			"status": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The status of the Checkout Session, one of `open`, `complete`, or `expired`.",
				Validators:  []validator.String{stringvalidator.OneOf("complete", "expired", "open")},
			},
			"submit_type": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Describes the type of transaction being performed by Checkout in order to customize\nrelevant text on the page, such as the submit button. `submit_type` can only be\nspecified on Checkout Sessions in `payment` mode. If blank or `auto`, `pay` is used.",
				Validators:  []validator.String{stringvalidator.OneOf("auto", "book", "donate", "pay", "subscribe")},
			},
			"subscription": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The ID of the [Subscription](https://docs.stripe.com/api/subscriptions) for Checkout Sessions in `subscription` mode.",
			},
			"success_url": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The URL the customer will be directed to after the payment or\nsubscription creation is successful.",
			},
			"tax_id_collection": ephemeralSchema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"enabled": ephemeralSchema.BoolAttribute{
						Required:    true,
						Description: "Indicates whether tax ID collection is enabled for the session",
					},
					"required": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Indicates whether a tax ID is required on the payment page",
						Validators:  []validator.String{stringvalidator.OneOf("if_supported", "never")},
					},
				},
			},
			"total_details": ephemeralSchema.SingleNestedAttribute{
				Computed:    true,
				Description: "Tax and discount details for the computed total amount.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"amount_discount": ephemeralSchema.Int64Attribute{
						Computed:    true,
						Description: "This is the sum of all the discounts.",
					},
					"amount_shipping": ephemeralSchema.Int64Attribute{
						Computed:    true,
						Description: "This is the sum of all the shipping amounts.",
					},
					"amount_tax": ephemeralSchema.Int64Attribute{
						Computed:    true,
						Description: "This is the sum of all the tax amounts.",
					},
					"breakdown": ephemeralSchema.SingleNestedAttribute{
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"discounts": ephemeralSchema.ListNestedAttribute{
								Computed:    true,
								Description: "The aggregated discounts.",
								NestedObject: ephemeralSchema.NestedAttributeObject{
									Attributes: map[string]ephemeralSchema.Attribute{
										"amount": ephemeralSchema.Int64Attribute{
											Computed:    true,
											Description: "The amount discounted.",
										},
										"discount": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "A discount represents the actual application of a [coupon](https://api.stripe.com#coupons) or [promotion code](https://api.stripe.com#promotion_codes).\nIt contains information about when the discount began, when it will end, and what it is applied to.\n\nRelated guide: [Applying discounts to subscriptions](https://docs.stripe.com/billing/subscriptions/discounts)",
										},
									},
								},
							},
							"taxes": ephemeralSchema.ListNestedAttribute{
								Computed:    true,
								Description: "The aggregated tax amounts by rate.",
								NestedObject: ephemeralSchema.NestedAttributeObject{
									Attributes: map[string]ephemeralSchema.Attribute{
										"amount": ephemeralSchema.Int64Attribute{
											Computed:    true,
											Description: "Amount of tax applied for this rate.",
										},
										"rate": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "Tax rates can be applied to [invoices](/invoicing/taxes/tax-rates), [subscriptions](/billing/taxes/tax-rates) and [Checkout Sessions](/payments/checkout/use-manual-tax-rates) to collect tax.\n\nRelated guide: [Tax rates](/billing/taxes/tax-rates)",
										},
										"taxability_reason": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.",
											Validators:  []validator.String{stringvalidator.OneOf("customer_exempt", "not_collecting", "not_subject_to_tax", "not_supported", "portion_product_exempt", "portion_reduced_rated", "portion_standard_rated", "product_exempt", "product_exempt_holiday", "proportionally_rated", "reduced_rated", "reverse_charge", "standard_rated", "taxable_basis_reduced", "zero_rated")},
										},
										"taxable_amount": ephemeralSchema.Int64Attribute{
											Computed:    true,
											Description: "The amount on which tax is calculated, in cents (or local equivalent).",
										},
									},
								},
							},
						},
					},
				},
			},
			"ui_mode": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The UI mode of the Session. Defaults to `hosted_page`.",
				Validators:  []validator.String{stringvalidator.OneOf("elements", "embedded_page", "form", "hosted_page")},
			},
			"url": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The URL to the Checkout Session. Applies to Checkout Sessions with `ui_mode: hosted_page`. Redirect customers to this URL to take them to Checkout. If you’re using [Custom Domains](https://docs.stripe.com/payments/checkout/custom-domains), the URL will use your subdomain. Otherwise, it’ll use `checkout.stripe.com.`\nThis value is only present when the session is active.",
			},
			"wallet_options": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Wallet-specific configuration for this Checkout Session.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"link": ephemeralSchema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"display": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Describes whether Checkout should display Link. Defaults to `auto`.",
								Validators:  []validator.String{stringvalidator.OneOf("auto", "never")},
							},
						},
					},
				},
			},
			"customer_update": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "Controls what fields on Customer can be updated by the Checkout Session. Can only be provided when `customer` is provided.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"address": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Describes whether Checkout saves the billing address onto `customer.address`.\nTo always collect a full billing address, use `billing_address_collection`. Defaults to `never`.",
					},
					"name": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Describes whether Checkout saves the name onto `customer.name`. Defaults to `never`.",
					},
					"shipping": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Describes whether Checkout saves shipping information onto `customer.shipping`.\nTo collect shipping information, use `shipping_address_collection`. Defaults to `never`.",
					},
				},
			},
			"payment_intent_data": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "A subset of parameters to be passed to PaymentIntent creation for Checkout Sessions in `payment` mode.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"application_fee_amount": ephemeralSchema.Int64Attribute{
						Optional:    true,
						Description: "The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. The amount of the application fee collected will be capped at the total amount captured. For more information, see the PaymentIntents [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts).",
					},
					"capture_method": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Controls when the funds will be captured from the customer's account.",
					},
					"description": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "An arbitrary string attached to the object. Often useful for displaying to users.",
					},
					"metadata": ephemeralSchema.MapAttribute{
						Optional:    true,
						Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
						ElementType: types.StringType,
					},
					"on_behalf_of": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The Stripe account ID for which these funds are intended. For details,\nsee the PaymentIntents [use case for connected\naccounts](/docs/payments/connected-accounts).",
					},
					"receipt_email": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Email address that the receipt for the resulting payment will be sent to. If `receipt_email` is specified for a payment in live mode, a receipt will be sent regardless of your [email settings](https://dashboard.stripe.com/account/emails).",
					},
					"setup_future_usage": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Indicates that you intend to [make future payments](https://docs.stripe.com/payments/payment-intents#future-usage) with the payment\nmethod collected by this Checkout Session.\n\nWhen setting this to `on_session`, Checkout will show a notice to the\ncustomer that their payment details will be saved.\n\nWhen setting this to `off_session`, Checkout will show a notice to the\ncustomer that their payment details will be saved and used for future\npayments.\n\nIf a Customer has been provided or Checkout creates a new Customer,\nCheckout will attach the payment method to the Customer.\n\nIf Checkout does not create a Customer, the payment method is not attached\nto a Customer. To reuse the payment method, you can retrieve it from the\nCheckout Session's PaymentIntent.\n\nWhen processing card payments, Checkout also uses `setup_future_usage`\nto dynamically optimize your payment flow and comply with regional\nlegislation and network rules, such as SCA.",
					},
					"shipping": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Shipping information for this payment.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"address": ephemeralSchema.SingleNestedAttribute{
								Required:    true,
								Description: "Shipping address.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"city": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "City, district, suburb, town, or village.",
									},
									"country": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
									},
									"line1": ephemeralSchema.StringAttribute{
										Required:    true,
										Description: "Address line 1, such as the street, PO Box, or company name.",
									},
									"line2": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Address line 2, such as the apartment, suite, unit, or building.",
									},
									"postal_code": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "ZIP or postal code.",
									},
									"state": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
									},
								},
							},
							"carrier": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The delivery service that shipped a physical product, such as Fedex, UPS, USPS, etc.",
							},
							"name": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Recipient name.",
							},
							"phone": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Recipient phone (including extension).",
							},
							"tracking_number": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.",
							},
						},
					},
					"statement_descriptor": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Text that appears on the customer's statement as the statement descriptor for a non-card charge. This value overrides the account's default statement descriptor. For information about requirements, including the 22-character limit, see [the Statement Descriptor docs](https://docs.stripe.com/get-started/account/statement-descriptors).\n\nSetting this value for a card charge returns an error. For card charges, set the [statement_descriptor_suffix](https://docs.stripe.com/get-started/account/statement-descriptors#dynamic) instead.",
					},
					"statement_descriptor_suffix": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Provides information about a card charge. Concatenated to the account's [statement descriptor prefix](https://docs.stripe.com/get-started/account/statement-descriptors#static) to form the complete statement descriptor that appears on the customer's statement.",
					},
					"transfer_data": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "The parameters used to automatically create a Transfer when the payment succeeds.\nFor more information, see the PaymentIntents [use case for connected accounts](https://docs.stripe.com/payments/connected-accounts).",
						Attributes: map[string]ephemeralSchema.Attribute{
							"amount": ephemeralSchema.Int64Attribute{
								Optional:    true,
								Description: "The amount that will be transferred automatically when a charge succeeds.",
							},
							"destination": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "If specified, successful charges will be attributed to the destination\naccount for tax reporting, and the funds from charges will be transferred\nto the destination account. The ID of the resulting transfer will be\nreturned on the successful charge's `transfer` field.",
							},
						},
					},
					"transfer_group": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "A string that identifies the resulting payment as part of a group. See the PaymentIntents [use case for connected accounts](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.",
					},
				},
			},
			"payment_method_configuration": ephemeralSchema.StringAttribute{
				Optional:    true,
				Description: "The ID of the payment method configuration to use with this Checkout session.",
			},
			"payment_method_data": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "This parameter allows you to set some attributes on the payment method created during a Checkout session.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"allow_redisplay": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Allow redisplay will be set on the payment method on confirmation and indicates whether this payment method can be shown again to the customer in a checkout flow. Only set this field if you wish to override the allow_redisplay value determined by Checkout.",
					},
				},
			},
			"setup_intent_data": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "A subset of parameters to be passed to SetupIntent creation for Checkout Sessions in `setup` mode.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"description": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "An arbitrary string attached to the object. Often useful for displaying to users.",
					},
					"metadata": ephemeralSchema.MapAttribute{
						Optional:    true,
						Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
						ElementType: types.StringType,
					},
					"on_behalf_of": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The Stripe account for which the setup is intended.",
					},
				},
			},
			"subscription_data": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "A subset of parameters to be passed to subscription creation for Checkout Sessions in `subscription` mode.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"application_fee_percent": ephemeralSchema.Float64Attribute{
						Optional:    true,
						Description: "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. To use an application fee percent, the request must be made on behalf of another account, using the `Stripe-Account` header or an OAuth key. For more information, see the application fees [documentation](https://stripe.com/docs/connect/subscriptions#collecting-fees-on-subscriptions).",
					},
					"billing_cycle_anchor": ephemeralSchema.Int64Attribute{
						Optional:    true,
						Description: "A future timestamp to anchor the subscription's billing cycle for new subscriptions. You can't set this parameter if `ui_mode` is `elements`.",
					},
					"billing_mode": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Controls how prorations and invoices for subscriptions are calculated and orchestrated.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"flexible": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "Configure behavior for flexible billing mode.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"proration_discounts": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Controls how invoices and invoice items display proration amounts and discount amounts.",
									},
								},
							},
							"type": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Controls the calculation and orchestration of prorations and invoices for subscriptions. If no value is passed, the default is `flexible`.",
							},
						},
					},
					"default_tax_rates": ephemeralSchema.ListAttribute{
						Optional:    true,
						Description: "The tax rates that will apply to any subscription item that does not have\n`tax_rates` set. Invoices created will have their `default_tax_rates` populated\nfrom the subscription.",
						ElementType: types.StringType,
					},
					"description": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The subscription's description, meant to be displayable to the customer.\nUse this field to optionally store an explanation of the subscription\nfor rendering in the [customer portal](https://docs.stripe.com/customer-management).",
					},
					"invoice_settings": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "All invoices will be billed using the specified settings.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"issuer": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"account": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The connected account being referenced when `type` is `account`.",
									},
									"type": ephemeralSchema.StringAttribute{
										Required:    true,
										Description: "Type of the account referenced in the request.",
									},
								},
							},
						},
					},
					"metadata": ephemeralSchema.MapAttribute{
						Optional:    true,
						Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
						ElementType: types.StringType,
					},
					"on_behalf_of": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The account on behalf of which to charge, for each of the subscription's invoices.",
					},
					"pending_invoice_item_interval": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Specifies an interval for how often to bill for any pending invoice items. It is analogous to calling [Create an invoice](https://docs.stripe.com/api#create_invoice) for the given subscription at the specified interval.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"interval": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Specifies invoicing frequency. Either `day`, `week`, `month` or `year`.",
							},
							"interval_count": ephemeralSchema.Int64Attribute{
								Optional:    true,
								Description: "The number of intervals between invoices. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).",
							},
						},
					},
					"proration_behavior": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Determines how to handle prorations resulting from the `billing_cycle_anchor`. If no value is passed, the default is `create_prorations`.",
					},
					"transfer_data": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "If specified, the funds from the subscription's invoices will be transferred to the destination and the ID of the resulting transfers will be found on the resulting charges.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"amount_percent": ephemeralSchema.Float64Attribute{
								Optional:    true,
								Description: "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount is transferred to the destination.",
							},
							"destination": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "ID of an existing, connected Stripe account.",
							},
						},
					},
					"trial_end": ephemeralSchema.Int64Attribute{
						Optional:    true,
						Description: "Unix timestamp representing the end of the trial period the customer will get before being charged for the first time. Has to be at least 48 hours in the future.",
					},
					"trial_period_days": ephemeralSchema.Int64Attribute{
						Optional:    true,
						Description: "Integer representing the number of trial period days before the customer is charged for the first time. Has to be at least 1.",
					},
					"trial_settings": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Settings related to subscription trials.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"end_behavior": ephemeralSchema.SingleNestedAttribute{
								Required:    true,
								Description: "Defines how the subscription should behave when the user's free trial ends.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"missing_payment_method": ephemeralSchema.StringAttribute{
										Required:    true,
										Description: "Indicates how the subscription should change when the trial ends if the user did not provide a payment method.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *CheckoutSessionEphemeralResource) Configure(_ context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stripe.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Ephemeral Resource Configure Type", fmt.Sprintf("Expected *stripe.Client, got: %T", req.ProviderData))
		return
	}

	r.client = client
}

func expandCheckoutSessionCreate(plan CheckoutSessionResourceModel) (*stripe.CheckoutSessionCreateParams, error) {
	params := &stripe.CheckoutSessionCreateParams{}

	if !plan.AdaptivePricing.IsNull() && !plan.AdaptivePricing.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AdaptivePricing", plan.AdaptivePricing) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "adaptive_pricing", params)
		}
	}
	if !plan.AfterExpiration.IsNull() && !plan.AfterExpiration.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AfterExpiration", plan.AfterExpiration) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "after_expiration", params)
		}
	}
	if !plan.AllowPromotionCodes.IsNull() && !plan.AllowPromotionCodes.IsUnknown() {
		params.AllowPromotionCodes = stripe.Bool(plan.AllowPromotionCodes.ValueBool())
	}
	if !plan.AutomaticTax.IsNull() && !plan.AutomaticTax.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AutomaticTax", plan.AutomaticTax) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "automatic_tax", params)
		}
	}
	if !plan.BillingAddressCollection.IsNull() && !plan.BillingAddressCollection.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "BillingAddressCollection", "BillingAddressCollection", plan.BillingAddressCollection.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_address_collection", params)
		}
	}
	if !plan.BrandingSettings.IsNull() && !plan.BrandingSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BrandingSettings", plan.BrandingSettings) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "branding_settings", params)
		}
	}
	if !plan.CancelURL.IsNull() && !plan.CancelURL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CancelURL", "CancelURL", plan.CancelURL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "cancel_url", params)
		}
	}
	if !plan.ClientReferenceID.IsNull() && !plan.ClientReferenceID.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ClientReferenceID", "ClientReferenceID", plan.ClientReferenceID.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "client_reference_id", params)
		}
	}
	if !plan.ConsentCollection.IsNull() && !plan.ConsentCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ConsentCollection", plan.ConsentCollection) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "consent_collection", params)
		}
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.CustomFields.IsNull() && !plan.CustomFields.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomFields", plan.CustomFields) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "custom_fields", params)
		}
	}
	if !plan.CustomText.IsNull() && !plan.CustomText.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomText", plan.CustomText) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "custom_text", params)
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
	if !plan.CustomerCreation.IsNull() && !plan.CustomerCreation.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerCreation", "CustomerCreation", plan.CustomerCreation.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer_creation", params)
		}
	}
	if !plan.CustomerEmail.IsNull() && !plan.CustomerEmail.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerEmail", "CustomerEmail", plan.CustomerEmail.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer_email", params)
		}
	}
	if !plan.Discounts.IsNull() && !plan.Discounts.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Discounts", plan.Discounts) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "discounts", params)
		}
	}
	if !plan.ExcludedPaymentMethodTypes.IsNull() && !plan.ExcludedPaymentMethodTypes.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ExcludedPaymentMethodTypes", plan.ExcludedPaymentMethodTypes) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "excluded_payment_method_types", params)
		}
	}
	if !plan.ExpiresAt.IsNull() && !plan.ExpiresAt.IsUnknown() {
		params.ExpiresAt = stripe.Int64(plan.ExpiresAt.ValueInt64())
	}
	if !plan.IntegrationIdentifier.IsNull() && !plan.IntegrationIdentifier.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "IntegrationIdentifier", "IntegrationIdentifier", plan.IntegrationIdentifier.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "integration_identifier", params)
		}
	}
	if !plan.InvoiceCreation.IsNull() && !plan.InvoiceCreation.IsUnknown() {
		if !assignAttrValueToNamedField(params, "InvoiceCreation", plan.InvoiceCreation) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "invoice_creation", params)
		}
	}
	if !plan.LineItems.IsNull() && !plan.LineItems.IsUnknown() {
		if !assignAttrValueToNamedField(params, "LineItems", plan.LineItems) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "line_items", params)
		}
	}
	if !plan.Locale.IsNull() && !plan.Locale.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Locale", "Locale", plan.Locale.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "locale", params)
		}
	}
	if !plan.ManagedPayments.IsNull() && !plan.ManagedPayments.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ManagedPayments", plan.ManagedPayments) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "managed_payments", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Mode.IsNull() && !plan.Mode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Mode", "Mode", plan.Mode.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "mode", params)
		}
	}
	if !plan.NameCollection.IsNull() && !plan.NameCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "NameCollection", plan.NameCollection) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "name_collection", params)
		}
	}
	if !plan.OptionalItems.IsNull() && !plan.OptionalItems.IsUnknown() {
		if !assignAttrValueToNamedField(params, "OptionalItems", plan.OptionalItems) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "optional_items", params)
		}
	}
	if !plan.OriginContext.IsNull() && !plan.OriginContext.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "OriginContext", "OriginContext", plan.OriginContext.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "origin_context", params)
		}
	}
	if !plan.PaymentMethodCollection.IsNull() && !plan.PaymentMethodCollection.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethodCollection", "PaymentMethodCollection", plan.PaymentMethodCollection.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method_collection", params)
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
	if !plan.Permissions.IsNull() && !plan.Permissions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Permissions", plan.Permissions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "permissions", params)
		}
	}
	if !plan.PhoneNumberCollection.IsNull() && !plan.PhoneNumberCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PhoneNumberCollection", plan.PhoneNumberCollection) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "phone_number_collection", params)
		}
	}
	if !plan.RedirectOnCompletion.IsNull() && !plan.RedirectOnCompletion.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "RedirectOnCompletion", "RedirectOnCompletion", plan.RedirectOnCompletion.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "redirect_on_completion", params)
		}
	}
	if !plan.ReturnURL.IsNull() && !plan.ReturnURL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReturnURL", "ReturnURL", plan.ReturnURL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "return_url", params)
		}
	}
	if !plan.SavedPaymentMethodOptions.IsNull() && !plan.SavedPaymentMethodOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SavedPaymentMethodOptions", plan.SavedPaymentMethodOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "saved_payment_method_options", params)
		}
	}
	if !plan.ShippingAddressCollection.IsNull() && !plan.ShippingAddressCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingAddressCollection", plan.ShippingAddressCollection) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping_address_collection", params)
		}
	}
	if !plan.ShippingOptions.IsNull() && !plan.ShippingOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingOptions", plan.ShippingOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping_options", params)
		}
	}
	if !plan.SubmitType.IsNull() && !plan.SubmitType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SubmitType", "SubmitType", plan.SubmitType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "submit_type", params)
		}
	}
	if !plan.SuccessURL.IsNull() && !plan.SuccessURL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SuccessURL", "SuccessURL", plan.SuccessURL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "success_url", params)
		}
	}
	if !plan.TaxIDCollection.IsNull() && !plan.TaxIDCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TaxIDCollection", plan.TaxIDCollection) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_id_collection", params)
		}
	}
	if !plan.UIMode.IsNull() && !plan.UIMode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "UIMode", "UIMode", plan.UIMode.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "ui_mode", params)
		}
	}
	if !plan.WalletOptions.IsNull() && !plan.WalletOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "WalletOptions", plan.WalletOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "wallet_options", params)
		}
	}
	if !plan.CustomerUpdate.IsNull() && !plan.CustomerUpdate.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomerUpdate", plan.CustomerUpdate) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer_update", params)
		}
	}
	if !plan.PaymentIntentData.IsNull() && !plan.PaymentIntentData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentIntentData", plan.PaymentIntentData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_intent_data", params)
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
	if !plan.SetupIntentData.IsNull() && !plan.SetupIntentData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SetupIntentData", plan.SetupIntentData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "setup_intent_data", params)
		}
	}
	if !plan.SubscriptionData.IsNull() && !plan.SubscriptionData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SubscriptionData", plan.SubscriptionData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "subscription_data", params)
		}
	}

	return params, nil
}

func flattenCheckoutSession(obj *stripe.CheckoutSession, state *CheckoutSessionResourceModel) error {
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
		assignedAdaptivePricing := false
		hadRawAdaptivePricing := false
		if rawValueAdaptivePricing, rawOk := plainValueAtPath(raw, "adaptive_pricing"); rawOk {
			hadRawAdaptivePricing = true
			if rawValueAdaptivePricing != nil {
				sourceAdaptivePricing := applyConfiguredKeyedListShapes(rawValueAdaptivePricing, attrValueToPlain(state.AdaptivePricing))
				if valueAdaptivePricing, err := flattenPlainValue(sourceAdaptivePricing, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "adaptive_pricing", "raw response"); err != nil {
					return err
				} else {
					if typedAdaptivePricing, ok := valueAdaptivePricing.(types.Object); ok {
						state.AdaptivePricing = typedAdaptivePricing
						assignedAdaptivePricing = true
					}
				}
			}
		}
		if !assignedAdaptivePricing {
			if !hasRaw {
				if responseValueAdaptivePricing, ok := plainFromResponseField(obj, "AdaptivePricing"); ok {
					sourceAdaptivePricing := applyConfiguredKeyedListShapes(responseValueAdaptivePricing, attrValueToPlain(state.AdaptivePricing))
					if valueAdaptivePricing, err := flattenPlainValue(
						sourceAdaptivePricing,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}},
						"adaptive_pricing",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAdaptivePricing, ok := valueAdaptivePricing.(types.Object); ok {
							state.AdaptivePricing = typedAdaptivePricing
							assignedAdaptivePricing = true
						}
					}
				}
			}
		}
		if !assignedAdaptivePricing && hadRawAdaptivePricing {
			if nullAdaptivePricing, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}); ok {
				if typedAdaptivePricing, ok := nullAdaptivePricing.(types.Object); ok {
					state.AdaptivePricing = typedAdaptivePricing
				}
			}
		}
	}
	{
		assignedAfterExpiration := false
		hadRawAfterExpiration := false
		if rawValueAfterExpiration, rawOk := plainValueAtPath(raw, "after_expiration"); rawOk {
			hadRawAfterExpiration = true
			if rawValueAfterExpiration != nil {
				sourceAfterExpiration := applyConfiguredKeyedListShapes(rawValueAfterExpiration, attrValueToPlain(state.AfterExpiration))
				if valueAfterExpiration, err := flattenPlainValue(sourceAfterExpiration, types.ObjectType{AttrTypes: map[string]attr.Type{"recovery": types.ObjectType{AttrTypes: map[string]attr.Type{"allow_promotion_codes": types.BoolType, "enabled": types.BoolType, "expires_at": types.Int64Type, "url": types.StringType}}}}, "after_expiration", "raw response"); err != nil {
					return err
				} else {
					if typedAfterExpiration, ok := valueAfterExpiration.(types.Object); ok {
						state.AfterExpiration = typedAfterExpiration
						assignedAfterExpiration = true
					}
				}
			}
		}
		if !assignedAfterExpiration {
			if !hasRaw {
				if responseValueAfterExpiration, ok := plainFromResponseField(obj, "AfterExpiration"); ok {
					sourceAfterExpiration := applyConfiguredKeyedListShapes(responseValueAfterExpiration, attrValueToPlain(state.AfterExpiration))
					if valueAfterExpiration, err := flattenPlainValue(
						sourceAfterExpiration,
						types.ObjectType{AttrTypes: map[string]attr.Type{"recovery": types.ObjectType{AttrTypes: map[string]attr.Type{"allow_promotion_codes": types.BoolType, "enabled": types.BoolType, "expires_at": types.Int64Type, "url": types.StringType}}}},
						"after_expiration",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAfterExpiration, ok := valueAfterExpiration.(types.Object); ok {
							state.AfterExpiration = typedAfterExpiration
							assignedAfterExpiration = true
						}
					}
				}
			}
		}
		if !assignedAfterExpiration && hadRawAfterExpiration {
			if nullAfterExpiration, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"recovery": types.ObjectType{AttrTypes: map[string]attr.Type{"allow_promotion_codes": types.BoolType, "enabled": types.BoolType, "expires_at": types.Int64Type, "url": types.StringType}}}}); ok {
				if typedAfterExpiration, ok := nullAfterExpiration.(types.Object); ok {
					state.AfterExpiration = typedAfterExpiration
				}
			}
		}
	}
	{
		if rawValueAllowPromotionCodes, rawOk := plainValueAtPath(raw, "allow_promotion_codes"); rawOk {
			if valueAllowPromotionCodes, err := flattenPlainValue(rawValueAllowPromotionCodes, types.BoolType, "allow_promotion_codes", "raw response"); err != nil {
				return err
			} else {
				if typedAllowPromotionCodes, ok := valueAllowPromotionCodes.(types.Bool); ok {
					state.AllowPromotionCodes = typedAllowPromotionCodes
				}
			}
		} else if !hasRaw {
			if responseValueAllowPromotionCodes, ok := plainFromResponseField(obj, "AllowPromotionCodes"); ok {
				if valueAllowPromotionCodes, err := flattenPlainValue(responseValueAllowPromotionCodes, types.BoolType, "allow_promotion_codes", "response struct"); err != nil {
					return err
				} else {
					if typedAllowPromotionCodes, ok := valueAllowPromotionCodes.(types.Bool); ok {
						state.AllowPromotionCodes = typedAllowPromotionCodes
					}
				}
			}
		}
	}
	{
		if rawValueAmountSubtotal, rawOk := plainValueAtPath(raw, "amount_subtotal"); rawOk {
			if valueAmountSubtotal, err := flattenPlainValue(rawValueAmountSubtotal, types.Int64Type, "amount_subtotal", "raw response"); err != nil {
				return err
			} else {
				if typedAmountSubtotal, ok := valueAmountSubtotal.(types.Int64); ok {
					state.AmountSubtotal = typedAmountSubtotal
				}
			}
		} else if !hasRaw {
			if responseValueAmountSubtotal, ok := plainFromResponseField(obj, "AmountSubtotal"); ok {
				if valueAmountSubtotal, err := flattenPlainValue(responseValueAmountSubtotal, types.Int64Type, "amount_subtotal", "response struct"); err != nil {
					return err
				} else {
					if typedAmountSubtotal, ok := valueAmountSubtotal.(types.Int64); ok {
						state.AmountSubtotal = typedAmountSubtotal
					}
				}
			}
		}
	}
	{
		if rawValueAmountTotal, rawOk := plainValueAtPath(raw, "amount_total"); rawOk {
			if valueAmountTotal, err := flattenPlainValue(rawValueAmountTotal, types.Int64Type, "amount_total", "raw response"); err != nil {
				return err
			} else {
				if typedAmountTotal, ok := valueAmountTotal.(types.Int64); ok {
					state.AmountTotal = typedAmountTotal
				}
			}
		} else if !hasRaw {
			if responseValueAmountTotal, ok := plainFromResponseField(obj, "AmountTotal"); ok {
				if valueAmountTotal, err := flattenPlainValue(responseValueAmountTotal, types.Int64Type, "amount_total", "response struct"); err != nil {
					return err
				} else {
					if typedAmountTotal, ok := valueAmountTotal.(types.Int64); ok {
						state.AmountTotal = typedAmountTotal
					}
				}
			}
		}
	}
	{
		assignedAutomaticTax := false
		hadRawAutomaticTax := false
		if rawValueAutomaticTax, rawOk := plainValueAtPath(raw, "automatic_tax"); rawOk {
			hadRawAutomaticTax = true
			if rawValueAutomaticTax != nil {
				sourceAutomaticTax := applyConfiguredKeyedListShapes(rawValueAutomaticTax, attrValueToPlain(state.AutomaticTax))
				if valueAutomaticTax, err := flattenPlainValue(sourceAutomaticTax, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "provider": types.StringType, "status": types.StringType}}, "automatic_tax", "raw response"); err != nil {
					return err
				} else {
					if typedAutomaticTax, ok := valueAutomaticTax.(types.Object); ok {
						state.AutomaticTax = typedAutomaticTax
						assignedAutomaticTax = true
					}
				}
			}
		}
		if !assignedAutomaticTax {
			if !hasRaw {
				if responseValueAutomaticTax, ok := plainFromResponseField(obj, "AutomaticTax"); ok {
					sourceAutomaticTax := applyConfiguredKeyedListShapes(responseValueAutomaticTax, attrValueToPlain(state.AutomaticTax))
					if valueAutomaticTax, err := flattenPlainValue(
						sourceAutomaticTax,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "provider": types.StringType, "status": types.StringType}},
						"automatic_tax",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAutomaticTax, ok := valueAutomaticTax.(types.Object); ok {
							state.AutomaticTax = typedAutomaticTax
							assignedAutomaticTax = true
						}
					}
				}
			}
		}
		if !assignedAutomaticTax && hadRawAutomaticTax {
			if nullAutomaticTax, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "provider": types.StringType, "status": types.StringType}}); ok {
				if typedAutomaticTax, ok := nullAutomaticTax.(types.Object); ok {
					state.AutomaticTax = typedAutomaticTax
				}
			}
		}
	}
	{
		if rawValueBillingAddressCollection, rawOk := plainValueAtPath(raw, "billing_address_collection"); rawOk {
			if valueBillingAddressCollection, err := flattenPlainValue(rawValueBillingAddressCollection, types.StringType, "billing_address_collection", "raw response"); err != nil {
				return err
			} else {
				if typedBillingAddressCollection, ok := valueBillingAddressCollection.(types.String); ok {
					state.BillingAddressCollection = typedBillingAddressCollection
				}
			}
		} else if !hasRaw {
			if responseValueBillingAddressCollection, ok := plainFromResponseField(obj, "BillingAddressCollection"); ok {
				if valueBillingAddressCollection, err := flattenPlainValue(responseValueBillingAddressCollection, types.StringType, "billing_address_collection", "response struct"); err != nil {
					return err
				} else {
					if typedBillingAddressCollection, ok := valueBillingAddressCollection.(types.String); ok {
						state.BillingAddressCollection = typedBillingAddressCollection
					}
				}
			}
		}
	}
	{
		assignedBrandingSettings := false
		hadRawBrandingSettings := false
		if rawValueBrandingSettings, rawOk := plainValueAtPath(raw, "branding_settings"); rawOk {
			hadRawBrandingSettings = true
			if rawValueBrandingSettings != nil {
				sourceBrandingSettings := applyConfiguredKeyedListShapes(rawValueBrandingSettings, attrValueToPlain(state.BrandingSettings))
				if valueBrandingSettings, err := flattenPlainValue(sourceBrandingSettings, types.ObjectType{AttrTypes: map[string]attr.Type{"background_color": types.StringType, "border_style": types.StringType, "button_color": types.StringType, "display_name": types.StringType, "font_family": types.StringType, "icon": types.ObjectType{AttrTypes: map[string]attr.Type{"file": types.StringType, "type": types.StringType, "url": types.StringType}}, "logo": types.ObjectType{AttrTypes: map[string]attr.Type{"file": types.StringType, "type": types.StringType, "url": types.StringType}}}}, "branding_settings", "raw response"); err != nil {
					return err
				} else {
					if typedBrandingSettings, ok := valueBrandingSettings.(types.Object); ok {
						state.BrandingSettings = typedBrandingSettings
						assignedBrandingSettings = true
					}
				}
			}
		}
		if !assignedBrandingSettings {
			if !hasRaw {
				if responseValueBrandingSettings, ok := plainFromResponseField(obj, "BrandingSettings"); ok {
					sourceBrandingSettings := applyConfiguredKeyedListShapes(responseValueBrandingSettings, attrValueToPlain(state.BrandingSettings))
					if valueBrandingSettings, err := flattenPlainValue(
						sourceBrandingSettings,
						types.ObjectType{AttrTypes: map[string]attr.Type{"background_color": types.StringType, "border_style": types.StringType, "button_color": types.StringType, "display_name": types.StringType, "font_family": types.StringType, "icon": types.ObjectType{AttrTypes: map[string]attr.Type{"file": types.StringType, "type": types.StringType, "url": types.StringType}}, "logo": types.ObjectType{AttrTypes: map[string]attr.Type{"file": types.StringType, "type": types.StringType, "url": types.StringType}}}},
						"branding_settings",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBrandingSettings, ok := valueBrandingSettings.(types.Object); ok {
							state.BrandingSettings = typedBrandingSettings
							assignedBrandingSettings = true
						}
					}
				}
			}
		}
		if !assignedBrandingSettings && hadRawBrandingSettings {
			if nullBrandingSettings, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"background_color": types.StringType, "border_style": types.StringType, "button_color": types.StringType, "display_name": types.StringType, "font_family": types.StringType, "icon": types.ObjectType{AttrTypes: map[string]attr.Type{"file": types.StringType, "type": types.StringType, "url": types.StringType}}, "logo": types.ObjectType{AttrTypes: map[string]attr.Type{"file": types.StringType, "type": types.StringType, "url": types.StringType}}}}); ok {
				if typedBrandingSettings, ok := nullBrandingSettings.(types.Object); ok {
					state.BrandingSettings = typedBrandingSettings
				}
			}
		}
	}
	{
		if rawValueCancelURL, rawOk := plainValueAtPath(raw, "cancel_url"); rawOk {
			if valueCancelURL, err := flattenPlainValue(rawValueCancelURL, types.StringType, "cancel_url", "raw response"); err != nil {
				return err
			} else {
				if typedCancelURL, ok := valueCancelURL.(types.String); ok {
					state.CancelURL = typedCancelURL
				}
			}
		} else if !hasRaw {
			if responseValueCancelURL, ok := plainFromResponseField(obj, "CancelURL"); ok {
				if valueCancelURL, err := flattenPlainValue(responseValueCancelURL, types.StringType, "cancel_url", "response struct"); err != nil {
					return err
				} else {
					if typedCancelURL, ok := valueCancelURL.(types.String); ok {
						state.CancelURL = typedCancelURL
					}
				}
			}
		}
	}
	{
		if rawValueClientReferenceID, rawOk := plainValueAtPath(raw, "client_reference_id"); rawOk {
			if valueClientReferenceID, err := flattenPlainValue(rawValueClientReferenceID, types.StringType, "client_reference_id", "raw response"); err != nil {
				return err
			} else {
				if typedClientReferenceID, ok := valueClientReferenceID.(types.String); ok {
					state.ClientReferenceID = typedClientReferenceID
				}
			}
		} else if !hasRaw {
			if responseValueClientReferenceID, ok := plainFromResponseField(obj, "ClientReferenceID"); ok {
				if valueClientReferenceID, err := flattenPlainValue(responseValueClientReferenceID, types.StringType, "client_reference_id", "response struct"); err != nil {
					return err
				} else {
					if typedClientReferenceID, ok := valueClientReferenceID.(types.String); ok {
						state.ClientReferenceID = typedClientReferenceID
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
		assignedCollectedInformation := false
		hadRawCollectedInformation := false
		if rawValueCollectedInformation, rawOk := plainValueAtPath(raw, "collected_information"); rawOk {
			hadRawCollectedInformation = true
			if rawValueCollectedInformation != nil {
				sourceCollectedInformation := applyConfiguredKeyedListShapes(rawValueCollectedInformation, attrValueToPlain(state.CollectedInformation))
				if valueCollectedInformation, err := flattenPlainValue(sourceCollectedInformation, types.ObjectType{AttrTypes: map[string]attr.Type{"business_name": types.StringType, "individual_name": types.StringType, "shipping_details": types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "name": types.StringType}}}}, "collected_information", "raw response"); err != nil {
					return err
				} else {
					if typedCollectedInformation, ok := valueCollectedInformation.(types.Object); ok {
						state.CollectedInformation = typedCollectedInformation
						assignedCollectedInformation = true
					}
				}
			}
		}
		if !assignedCollectedInformation {
			if !hasRaw {
				if responseValueCollectedInformation, ok := plainFromResponseField(obj, "CollectedInformation"); ok {
					sourceCollectedInformation := applyConfiguredKeyedListShapes(responseValueCollectedInformation, attrValueToPlain(state.CollectedInformation))
					if valueCollectedInformation, err := flattenPlainValue(
						sourceCollectedInformation,
						types.ObjectType{AttrTypes: map[string]attr.Type{"business_name": types.StringType, "individual_name": types.StringType, "shipping_details": types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "name": types.StringType}}}},
						"collected_information",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCollectedInformation, ok := valueCollectedInformation.(types.Object); ok {
							state.CollectedInformation = typedCollectedInformation
							assignedCollectedInformation = true
						}
					}
				}
			}
		}
		if !assignedCollectedInformation && hadRawCollectedInformation {
			if nullCollectedInformation, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"business_name": types.StringType, "individual_name": types.StringType, "shipping_details": types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "name": types.StringType}}}}); ok {
				if typedCollectedInformation, ok := nullCollectedInformation.(types.Object); ok {
					state.CollectedInformation = typedCollectedInformation
				}
			}
		}
	}
	{
		assignedConsent := false
		hadRawConsent := false
		if rawValueConsent, rawOk := plainValueAtPath(raw, "consent"); rawOk {
			hadRawConsent = true
			if rawValueConsent != nil {
				sourceConsent := applyConfiguredKeyedListShapes(rawValueConsent, attrValueToPlain(state.Consent))
				if valueConsent, err := flattenPlainValue(sourceConsent, types.ObjectType{AttrTypes: map[string]attr.Type{"promotions": types.StringType, "terms_of_service": types.StringType}}, "consent", "raw response"); err != nil {
					return err
				} else {
					if typedConsent, ok := valueConsent.(types.Object); ok {
						state.Consent = typedConsent
						assignedConsent = true
					}
				}
			}
		}
		if !assignedConsent {
			if !hasRaw {
				if responseValueConsent, ok := plainFromResponseField(obj, "Consent"); ok {
					sourceConsent := applyConfiguredKeyedListShapes(responseValueConsent, attrValueToPlain(state.Consent))
					if valueConsent, err := flattenPlainValue(
						sourceConsent,
						types.ObjectType{AttrTypes: map[string]attr.Type{"promotions": types.StringType, "terms_of_service": types.StringType}},
						"consent",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedConsent, ok := valueConsent.(types.Object); ok {
							state.Consent = typedConsent
							assignedConsent = true
						}
					}
				}
			}
		}
		if !assignedConsent && hadRawConsent {
			if nullConsent, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"promotions": types.StringType, "terms_of_service": types.StringType}}); ok {
				if typedConsent, ok := nullConsent.(types.Object); ok {
					state.Consent = typedConsent
				}
			}
		}
	}
	{
		assignedConsentCollection := false
		hadRawConsentCollection := false
		if rawValueConsentCollection, rawOk := plainValueAtPath(raw, "consent_collection"); rawOk {
			hadRawConsentCollection = true
			if rawValueConsentCollection != nil {
				sourceConsentCollection := applyConfiguredKeyedListShapes(rawValueConsentCollection, attrValueToPlain(state.ConsentCollection))
				if valueConsentCollection, err := flattenPlainValue(sourceConsentCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"payment_method_reuse_agreement": types.ObjectType{AttrTypes: map[string]attr.Type{"position": types.StringType}}, "promotions": types.StringType, "terms_of_service": types.StringType}}, "consent_collection", "raw response"); err != nil {
					return err
				} else {
					if typedConsentCollection, ok := valueConsentCollection.(types.Object); ok {
						state.ConsentCollection = typedConsentCollection
						assignedConsentCollection = true
					}
				}
			}
		}
		if !assignedConsentCollection {
			if !hasRaw {
				if responseValueConsentCollection, ok := plainFromResponseField(obj, "ConsentCollection"); ok {
					sourceConsentCollection := applyConfiguredKeyedListShapes(responseValueConsentCollection, attrValueToPlain(state.ConsentCollection))
					if valueConsentCollection, err := flattenPlainValue(
						sourceConsentCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"payment_method_reuse_agreement": types.ObjectType{AttrTypes: map[string]attr.Type{"position": types.StringType}}, "promotions": types.StringType, "terms_of_service": types.StringType}},
						"consent_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedConsentCollection, ok := valueConsentCollection.(types.Object); ok {
							state.ConsentCollection = typedConsentCollection
							assignedConsentCollection = true
						}
					}
				}
			}
		}
		if !assignedConsentCollection && hadRawConsentCollection {
			if nullConsentCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"payment_method_reuse_agreement": types.ObjectType{AttrTypes: map[string]attr.Type{"position": types.StringType}}, "promotions": types.StringType, "terms_of_service": types.StringType}}); ok {
				if typedConsentCollection, ok := nullConsentCollection.(types.Object); ok {
					state.ConsentCollection = typedConsentCollection
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
		assignedCurrencyConversion := false
		hadRawCurrencyConversion := false
		if rawValueCurrencyConversion, rawOk := plainValueAtPath(raw, "currency_conversion"); rawOk {
			hadRawCurrencyConversion = true
			if rawValueCurrencyConversion != nil {
				sourceCurrencyConversion := applyConfiguredKeyedListShapes(rawValueCurrencyConversion, attrValueToPlain(state.CurrencyConversion))
				if valueCurrencyConversion, err := flattenPlainValue(sourceCurrencyConversion, types.ObjectType{AttrTypes: map[string]attr.Type{"amount_subtotal": types.Int64Type, "amount_total": types.Int64Type, "fx_rate": types.Float64Type, "source_currency": types.StringType}}, "currency_conversion", "raw response"); err != nil {
					return err
				} else {
					if typedCurrencyConversion, ok := valueCurrencyConversion.(types.Object); ok {
						state.CurrencyConversion = typedCurrencyConversion
						assignedCurrencyConversion = true
					}
				}
			}
		}
		if !assignedCurrencyConversion {
			if !hasRaw {
				if responseValueCurrencyConversion, ok := plainFromResponseField(obj, "CurrencyConversion"); ok {
					sourceCurrencyConversion := applyConfiguredKeyedListShapes(responseValueCurrencyConversion, attrValueToPlain(state.CurrencyConversion))
					if valueCurrencyConversion, err := flattenPlainValue(
						sourceCurrencyConversion,
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount_subtotal": types.Int64Type, "amount_total": types.Int64Type, "fx_rate": types.Float64Type, "source_currency": types.StringType}},
						"currency_conversion",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCurrencyConversion, ok := valueCurrencyConversion.(types.Object); ok {
							state.CurrencyConversion = typedCurrencyConversion
							assignedCurrencyConversion = true
						}
					}
				}
			}
		}
		if !assignedCurrencyConversion && hadRawCurrencyConversion {
			if nullCurrencyConversion, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount_subtotal": types.Int64Type, "amount_total": types.Int64Type, "fx_rate": types.Float64Type, "source_currency": types.StringType}}); ok {
				if typedCurrencyConversion, ok := nullCurrencyConversion.(types.Object); ok {
					state.CurrencyConversion = typedCurrencyConversion
				}
			}
		}
	}
	{
		if rawValueCustomFields, rawOk := plainValueAtPath(raw, "custom_fields"); rawOk {
			if valueCustomFields, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueCustomFields, attrValueToPlain(state.CustomFields)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"dropdown": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"label": types.StringType, "value": types.StringType}}}, "value": types.StringType}}, "key": types.StringType, "label": types.ObjectType{AttrTypes: map[string]attr.Type{"custom": types.StringType, "type": types.StringType}}, "numeric": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "maximum_length": types.Int64Type, "minimum_length": types.Int64Type, "value": types.StringType}}, "optional": types.BoolType, "text": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "maximum_length": types.Int64Type, "minimum_length": types.Int64Type, "value": types.StringType}}, "type": types.StringType}}}, "custom_fields", "raw response"); err != nil {
				return err
			} else {
				if typedCustomFields, ok := valueCustomFields.(types.List); ok {
					state.CustomFields = typedCustomFields
				}
			}
		} else if !hasRaw {
			if responseValueCustomFields, ok := plainFromResponseField(obj, "CustomFields"); ok {
				if valueCustomFields, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueCustomFields, attrValueToPlain(state.CustomFields)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"dropdown": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"label": types.StringType, "value": types.StringType}}}, "value": types.StringType}}, "key": types.StringType, "label": types.ObjectType{AttrTypes: map[string]attr.Type{"custom": types.StringType, "type": types.StringType}}, "numeric": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "maximum_length": types.Int64Type, "minimum_length": types.Int64Type, "value": types.StringType}}, "optional": types.BoolType, "text": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "maximum_length": types.Int64Type, "minimum_length": types.Int64Type, "value": types.StringType}}, "type": types.StringType}}},
					"custom_fields",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedCustomFields, ok := valueCustomFields.(types.List); ok {
						state.CustomFields = typedCustomFields
					}
				}
			}
		}
	}
	{
		assignedCustomText := false
		hadRawCustomText := false
		if rawValueCustomText, rawOk := plainValueAtPath(raw, "custom_text"); rawOk {
			hadRawCustomText = true
			if rawValueCustomText != nil {
				sourceCustomText := applyConfiguredKeyedListShapes(rawValueCustomText, attrValueToPlain(state.CustomText))
				if valueCustomText, err := flattenPlainValue(sourceCustomText, types.ObjectType{AttrTypes: map[string]attr.Type{"after_submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "terms_of_service_acceptance": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}}}, "custom_text", "raw response"); err != nil {
					return err
				} else {
					if typedCustomText, ok := valueCustomText.(types.Object); ok {
						state.CustomText = typedCustomText
						assignedCustomText = true
					}
				}
			}
		}
		if !assignedCustomText {
			if !hasRaw {
				if responseValueCustomText, ok := plainFromResponseField(obj, "CustomText"); ok {
					sourceCustomText := applyConfiguredKeyedListShapes(responseValueCustomText, attrValueToPlain(state.CustomText))
					if valueCustomText, err := flattenPlainValue(
						sourceCustomText,
						types.ObjectType{AttrTypes: map[string]attr.Type{"after_submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "terms_of_service_acceptance": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}}},
						"custom_text",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCustomText, ok := valueCustomText.(types.Object); ok {
							state.CustomText = typedCustomText
							assignedCustomText = true
						}
					}
				}
			}
		}
		if !assignedCustomText && hadRawCustomText {
			if nullCustomText, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"after_submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "terms_of_service_acceptance": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}}}); ok {
				if typedCustomText, ok := nullCustomText.(types.Object); ok {
					state.CustomText = typedCustomText
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
		if rawValueCustomerCreation, rawOk := plainValueAtPath(raw, "customer_creation"); rawOk {
			if valueCustomerCreation, err := flattenPlainValue(rawValueCustomerCreation, types.StringType, "customer_creation", "raw response"); err != nil {
				return err
			} else {
				if typedCustomerCreation, ok := valueCustomerCreation.(types.String); ok {
					state.CustomerCreation = typedCustomerCreation
				}
			}
		} else if !hasRaw {
			if responseValueCustomerCreation, ok := plainFromResponseField(obj, "CustomerCreation"); ok {
				if valueCustomerCreation, err := flattenPlainValue(responseValueCustomerCreation, types.StringType, "customer_creation", "response struct"); err != nil {
					return err
				} else {
					if typedCustomerCreation, ok := valueCustomerCreation.(types.String); ok {
						state.CustomerCreation = typedCustomerCreation
					}
				}
			}
		}
	}
	{
		assignedCustomerDetails := false
		hadRawCustomerDetails := false
		if rawValueCustomerDetails, rawOk := plainValueAtPath(raw, "customer_details"); rawOk {
			hadRawCustomerDetails = true
			if rawValueCustomerDetails != nil {
				sourceCustomerDetails := applyConfiguredKeyedListShapes(rawValueCustomerDetails, attrValueToPlain(state.CustomerDetails))
				if valueCustomerDetails, err := flattenPlainValue(sourceCustomerDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "business_name": types.StringType, "email": types.StringType, "individual_name": types.StringType, "name": types.StringType, "phone": types.StringType, "tax_exempt": types.StringType, "tax_ids": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}}}}, "customer_details", "raw response"); err != nil {
					return err
				} else {
					if typedCustomerDetails, ok := valueCustomerDetails.(types.Object); ok {
						state.CustomerDetails = typedCustomerDetails
						assignedCustomerDetails = true
					}
				}
			}
		}
		if !assignedCustomerDetails {
			if !hasRaw {
				if responseValueCustomerDetails, ok := plainFromResponseField(obj, "CustomerDetails"); ok {
					sourceCustomerDetails := applyConfiguredKeyedListShapes(responseValueCustomerDetails, attrValueToPlain(state.CustomerDetails))
					if valueCustomerDetails, err := flattenPlainValue(
						sourceCustomerDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "business_name": types.StringType, "email": types.StringType, "individual_name": types.StringType, "name": types.StringType, "phone": types.StringType, "tax_exempt": types.StringType, "tax_ids": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}}}},
						"customer_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCustomerDetails, ok := valueCustomerDetails.(types.Object); ok {
							state.CustomerDetails = typedCustomerDetails
							assignedCustomerDetails = true
						}
					}
				}
			}
		}
		if !assignedCustomerDetails && hadRawCustomerDetails {
			if nullCustomerDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "business_name": types.StringType, "email": types.StringType, "individual_name": types.StringType, "name": types.StringType, "phone": types.StringType, "tax_exempt": types.StringType, "tax_ids": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}}}}); ok {
				if typedCustomerDetails, ok := nullCustomerDetails.(types.Object); ok {
					state.CustomerDetails = typedCustomerDetails
				}
			}
		}
	}
	{
		if rawValueCustomerEmail, rawOk := plainValueAtPath(raw, "customer_email"); rawOk {
			if valueCustomerEmail, err := flattenPlainValue(rawValueCustomerEmail, types.StringType, "customer_email", "raw response"); err != nil {
				return err
			} else {
				if typedCustomerEmail, ok := valueCustomerEmail.(types.String); ok {
					state.CustomerEmail = typedCustomerEmail
				}
			}
		} else if !hasRaw {
			if responseValueCustomerEmail, ok := plainFromResponseField(obj, "CustomerEmail"); ok {
				if valueCustomerEmail, err := flattenPlainValue(responseValueCustomerEmail, types.StringType, "customer_email", "response struct"); err != nil {
					return err
				} else {
					if typedCustomerEmail, ok := valueCustomerEmail.(types.String); ok {
						state.CustomerEmail = typedCustomerEmail
					}
				}
			}
		}
	}
	{
		if rawValueDiscounts, rawOk := plainValueAtPath(raw, "discounts"); rawOk {
			if valueDiscounts, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueDiscounts, attrValueToPlain(state.Discounts)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "promotion_code": types.StringType}}}, "discounts", "raw response"); err != nil {
				return err
			} else {
				if typedDiscounts, ok := valueDiscounts.(types.List); ok {
					state.Discounts = typedDiscounts
				}
			}
		} else if !hasRaw {
			if responseValueDiscounts, ok := plainFromResponseField(obj, "Discounts"); ok {
				if valueDiscounts, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueDiscounts, attrValueToPlain(state.Discounts)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "promotion_code": types.StringType}}},
					"discounts",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedDiscounts, ok := valueDiscounts.(types.List); ok {
						state.Discounts = typedDiscounts
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
		if rawValueExpiresAt, rawOk := plainValueAtPath(raw, "expires_at"); rawOk {
			if valueExpiresAt, err := flattenPlainValue(rawValueExpiresAt, types.Int64Type, "expires_at", "raw response"); err != nil {
				return err
			} else {
				if typedExpiresAt, ok := valueExpiresAt.(types.Int64); ok {
					state.ExpiresAt = typedExpiresAt
				}
			}
		} else if !hasRaw {
			if responseValueExpiresAt, ok := plainFromResponseField(obj, "ExpiresAt"); ok {
				if valueExpiresAt, err := flattenPlainValue(responseValueExpiresAt, types.Int64Type, "expires_at", "response struct"); err != nil {
					return err
				} else {
					if typedExpiresAt, ok := valueExpiresAt.(types.Int64); ok {
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
		if rawValueIntegrationIdentifier, rawOk := plainValueAtPath(raw, "integration_identifier"); rawOk {
			if valueIntegrationIdentifier, err := flattenPlainValue(rawValueIntegrationIdentifier, types.StringType, "integration_identifier", "raw response"); err != nil {
				return err
			} else {
				if typedIntegrationIdentifier, ok := valueIntegrationIdentifier.(types.String); ok {
					state.IntegrationIdentifier = typedIntegrationIdentifier
				}
			}
		} else if !hasRaw {
			if responseValueIntegrationIdentifier, ok := plainFromResponseField(obj, "IntegrationIdentifier"); ok {
				if valueIntegrationIdentifier, err := flattenPlainValue(responseValueIntegrationIdentifier, types.StringType, "integration_identifier", "response struct"); err != nil {
					return err
				} else {
					if typedIntegrationIdentifier, ok := valueIntegrationIdentifier.(types.String); ok {
						state.IntegrationIdentifier = typedIntegrationIdentifier
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueInvoice, rawOk := plainValueAtPath(raw, "invoice"); rawOk {
				if typedInvoice, ok := plainToStringIDValue(rawValueInvoice); ok {
					state.Invoice = typedInvoice
				}
			} else if !hasRaw {
				if responseValueInvoice, ok := plainFromResponseField(obj, "Invoice"); ok {
					if typedInvoice, ok := plainToStringIDValue(responseValueInvoice); ok {
						state.Invoice = typedInvoice
					}
				}
			}
		}
	}
	{
		assignedInvoiceCreation := false
		hadRawInvoiceCreation := false
		if rawValueInvoiceCreation, rawOk := plainValueAtPath(raw, "invoice_creation"); rawOk {
			hadRawInvoiceCreation = true
			if rawValueInvoiceCreation != nil {
				sourceInvoiceCreation := applyConfiguredKeyedListShapes(rawValueInvoiceCreation, attrValueToPlain(state.InvoiceCreation))
				if valueInvoiceCreation, err := flattenPlainValue(sourceInvoiceCreation, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "invoice_data": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "custom_fields": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "description": types.StringType, "footer": types.StringType, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "metadata": types.MapType{ElemType: types.StringType}, "rendering_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}}}}, "invoice_creation", "raw response"); err != nil {
					return err
				} else {
					if typedInvoiceCreation, ok := valueInvoiceCreation.(types.Object); ok {
						state.InvoiceCreation = typedInvoiceCreation
						assignedInvoiceCreation = true
					}
				}
			}
		}
		if !assignedInvoiceCreation {
			if !hasRaw {
				if responseValueInvoiceCreation, ok := plainFromResponseField(obj, "InvoiceCreation"); ok {
					sourceInvoiceCreation := applyConfiguredKeyedListShapes(responseValueInvoiceCreation, attrValueToPlain(state.InvoiceCreation))
					if valueInvoiceCreation, err := flattenPlainValue(
						sourceInvoiceCreation,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "invoice_data": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "custom_fields": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "description": types.StringType, "footer": types.StringType, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "metadata": types.MapType{ElemType: types.StringType}, "rendering_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}}}},
						"invoice_creation",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedInvoiceCreation, ok := valueInvoiceCreation.(types.Object); ok {
							state.InvoiceCreation = typedInvoiceCreation
							assignedInvoiceCreation = true
						}
					}
				}
			}
		}
		if !assignedInvoiceCreation && hadRawInvoiceCreation {
			if nullInvoiceCreation, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "invoice_data": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "custom_fields": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "description": types.StringType, "footer": types.StringType, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "metadata": types.MapType{ElemType: types.StringType}, "rendering_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}}}}); ok {
				if typedInvoiceCreation, ok := nullInvoiceCreation.(types.Object); ok {
					state.InvoiceCreation = typedInvoiceCreation
				}
			}
		}
	}
	{
		if rawValueLineItems, rawOk := plainValueAtPath(raw, "line_items"); rawOk {
			rawPlainLineItems := extractListObjectData(rawValueLineItems)
			if valueLineItems, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawPlainLineItems, attrValueToPlain(state.LineItems)), attrValueToPlain(state.LineItems)), attrValueToPlain(state.LineItems)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "dynamic_tax_rates": types.ListType{ElemType: types.StringType}, "metadata": types.MapType{ElemType: types.StringType}, "price": types.StringType, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "product_data": types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "images": types.ListType{ElemType: types.StringType}, "metadata": types.MapType{ElemType: types.StringType}, "name": types.StringType, "tax_code": types.StringType, "unit_label": types.StringType}}, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}, "quantity": types.Int64Type, "tax_rates": types.ListType{ElemType: types.StringType}}}}, "line_items", "raw response"); err != nil {
				return err
			} else {
				if typedLineItems, ok := valueLineItems.(types.List); ok {
					state.LineItems = typedLineItems
				}
			}
		} else if !hasRaw {
			if responseValueLineItems, ok := plainFromResponseField(obj, "LineItems"); ok {
				fallbackPlainLineItems := extractListObjectData(responseValueLineItems)
				if valueLineItems, err := flattenPlainValue(
					preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(fallbackPlainLineItems, attrValueToPlain(state.LineItems)), attrValueToPlain(state.LineItems)), attrValueToPlain(state.LineItems)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "dynamic_tax_rates": types.ListType{ElemType: types.StringType}, "metadata": types.MapType{ElemType: types.StringType}, "price": types.StringType, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "product_data": types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "images": types.ListType{ElemType: types.StringType}, "metadata": types.MapType{ElemType: types.StringType}, "name": types.StringType, "tax_code": types.StringType, "unit_label": types.StringType}}, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}, "quantity": types.Int64Type, "tax_rates": types.ListType{ElemType: types.StringType}}}},
					"line_items",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedLineItems, ok := valueLineItems.(types.List); ok {
						state.LineItems = typedLineItems
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
		if rawValueLocale, rawOk := plainValueAtPath(raw, "locale"); rawOk {
			if valueLocale, err := flattenPlainValue(rawValueLocale, types.StringType, "locale", "raw response"); err != nil {
				return err
			} else {
				if typedLocale, ok := valueLocale.(types.String); ok {
					state.Locale = typedLocale
				}
			}
		} else if !hasRaw {
			if responseValueLocale, ok := plainFromResponseField(obj, "Locale"); ok {
				if valueLocale, err := flattenPlainValue(responseValueLocale, types.StringType, "locale", "response struct"); err != nil {
					return err
				} else {
					if typedLocale, ok := valueLocale.(types.String); ok {
						state.Locale = typedLocale
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
		if rawValueMode, rawOk := plainValueAtPath(raw, "mode"); rawOk {
			if valueMode, err := flattenPlainValue(rawValueMode, types.StringType, "mode", "raw response"); err != nil {
				return err
			} else {
				if typedMode, ok := valueMode.(types.String); ok {
					state.Mode = typedMode
				}
			}
		} else if !hasRaw {
			if responseValueMode, ok := plainFromResponseField(obj, "Mode"); ok {
				if valueMode, err := flattenPlainValue(responseValueMode, types.StringType, "mode", "response struct"); err != nil {
					return err
				} else {
					if typedMode, ok := valueMode.(types.String); ok {
						state.Mode = typedMode
					}
				}
			}
		}
	}
	{
		assignedNameCollection := false
		hadRawNameCollection := false
		if rawValueNameCollection, rawOk := plainValueAtPath(raw, "name_collection"); rawOk {
			hadRawNameCollection = true
			if rawValueNameCollection != nil {
				sourceNameCollection := applyConfiguredKeyedListShapes(rawValueNameCollection, attrValueToPlain(state.NameCollection))
				if valueNameCollection, err := flattenPlainValue(sourceNameCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"business": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}, "individual": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}}}, "name_collection", "raw response"); err != nil {
					return err
				} else {
					if typedNameCollection, ok := valueNameCollection.(types.Object); ok {
						state.NameCollection = typedNameCollection
						assignedNameCollection = true
					}
				}
			}
		}
		if !assignedNameCollection {
			if !hasRaw {
				if responseValueNameCollection, ok := plainFromResponseField(obj, "NameCollection"); ok {
					sourceNameCollection := applyConfiguredKeyedListShapes(responseValueNameCollection, attrValueToPlain(state.NameCollection))
					if valueNameCollection, err := flattenPlainValue(
						sourceNameCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"business": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}, "individual": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}}},
						"name_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedNameCollection, ok := valueNameCollection.(types.Object); ok {
							state.NameCollection = typedNameCollection
							assignedNameCollection = true
						}
					}
				}
			}
		}
		if !assignedNameCollection && hadRawNameCollection {
			if nullNameCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"business": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}, "individual": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}}}); ok {
				if typedNameCollection, ok := nullNameCollection.(types.Object); ok {
					state.NameCollection = typedNameCollection
				}
			}
		}
	}
	{
		if rawValueOptionalItems, rawOk := plainValueAtPath(raw, "optional_items"); rawOk {
			if valueOptionalItems, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueOptionalItems, attrValueToPlain(state.OptionalItems)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "price": types.StringType, "quantity": types.Int64Type}}}, "optional_items", "raw response"); err != nil {
				return err
			} else {
				if typedOptionalItems, ok := valueOptionalItems.(types.List); ok {
					state.OptionalItems = typedOptionalItems
				}
			}
		} else if !hasRaw {
			if responseValueOptionalItems, ok := plainFromResponseField(obj, "OptionalItems"); ok {
				if valueOptionalItems, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueOptionalItems, attrValueToPlain(state.OptionalItems)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "price": types.StringType, "quantity": types.Int64Type}}},
					"optional_items",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedOptionalItems, ok := valueOptionalItems.(types.List); ok {
						state.OptionalItems = typedOptionalItems
					}
				}
			}
		}
	}
	{
		if rawValueOriginContext, rawOk := plainValueAtPath(raw, "origin_context"); rawOk {
			if valueOriginContext, err := flattenPlainValue(rawValueOriginContext, types.StringType, "origin_context", "raw response"); err != nil {
				return err
			} else {
				if typedOriginContext, ok := valueOriginContext.(types.String); ok {
					state.OriginContext = typedOriginContext
				}
			}
		} else if !hasRaw {
			if responseValueOriginContext, ok := plainFromResponseField(obj, "OriginContext"); ok {
				if valueOriginContext, err := flattenPlainValue(responseValueOriginContext, types.StringType, "origin_context", "response struct"); err != nil {
					return err
				} else {
					if typedOriginContext, ok := valueOriginContext.(types.String); ok {
						state.OriginContext = typedOriginContext
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
		if true {
			if rawValuePaymentLink, rawOk := plainValueAtPath(raw, "payment_link"); rawOk {
				if typedPaymentLink, ok := plainToStringIDValue(rawValuePaymentLink); ok {
					state.PaymentLink = typedPaymentLink
				}
			} else if !hasRaw {
				if responseValuePaymentLink, ok := plainFromResponseField(obj, "PaymentLink"); ok {
					if typedPaymentLink, ok := plainToStringIDValue(responseValuePaymentLink); ok {
						state.PaymentLink = typedPaymentLink
					}
				}
			}
		}
	}
	{
		if rawValuePaymentMethodCollection, rawOk := plainValueAtPath(raw, "payment_method_collection"); rawOk {
			if valuePaymentMethodCollection, err := flattenPlainValue(rawValuePaymentMethodCollection, types.StringType, "payment_method_collection", "raw response"); err != nil {
				return err
			} else {
				if typedPaymentMethodCollection, ok := valuePaymentMethodCollection.(types.String); ok {
					state.PaymentMethodCollection = typedPaymentMethodCollection
				}
			}
		} else if !hasRaw {
			if responseValuePaymentMethodCollection, ok := plainFromResponseField(obj, "PaymentMethodCollection"); ok {
				if valuePaymentMethodCollection, err := flattenPlainValue(responseValuePaymentMethodCollection, types.StringType, "payment_method_collection", "response struct"); err != nil {
					return err
				} else {
					if typedPaymentMethodCollection, ok := valuePaymentMethodCollection.(types.String); ok {
						state.PaymentMethodCollection = typedPaymentMethodCollection
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
				if valuePaymentMethodOptions, err := flattenPlainValue(sourcePaymentMethodOptions, types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_mandate_url": types.StringType, "default_for": types.ListType{ElemType: types.StringType}, "interval_description": types.StringType, "payment_schedule": types.StringType, "transaction_type": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "verification_method": types.StringType}}, "affirm": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "afterpay_clearpay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "alipay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "alma": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "amazon_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "au_becs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "target_date": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "billie": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "boleto": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "installments": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "request_extended_authorization": types.StringType, "request_incremental_authorization": types.StringType, "request_multicapture": types.StringType, "request_overcapture": types.StringType, "request_three_d_secure": types.StringType, "restrictions": types.ObjectType{AttrTypes: map[string]attr.Type{"brands_blocked": types.ListType{ElemType: types.StringType}}}, "setup_future_usage": types.StringType, "statement_descriptor_suffix_kana": types.StringType, "statement_descriptor_suffix_kanji": types.StringType}}, "cashapp": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "requested_address_types": types.ListType{ElemType: types.StringType}, "type": types.StringType}}, "funding_type": types.StringType, "setup_future_usage": types.StringType}}, "eps": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "fpx": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "giropay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "grabpay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "ideal": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "kakao_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType, "subscriptions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "name": types.StringType, "next_billing": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "date": types.StringType}}, "reference": types.StringType}}}}}, "konbini": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "kr_card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "link": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "mobilepay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "multibanco": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "naver_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "oxxo": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "p24": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "tos_shown_and_accepted": types.BoolType}}, "payco": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "paynow": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "reference": types.StringType, "setup_future_usage": types.StringType, "risk_correlation_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "payments_per_period": types.Int64Type, "purpose": types.StringType, "start_date": types.StringType}}, "setup_future_usage": types.StringType}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_includes_iof": types.StringType, "expires_after_seconds": types.Int64Type, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "amount_type": types.StringType, "currency": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "reference": types.StringType, "start_date": types.StringType}}, "setup_future_usage": types.StringType}}, "revolut_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "samsung_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "satispay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "scalapay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "sofort": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "swish": types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType}}, "twint": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}, "setup_future_usage": types.StringType}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}, "return_url": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "verification_method": types.StringType}}, "crypto": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "demo_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "wechat_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"app_id": types.StringType, "client": types.StringType, "setup_future_usage": types.StringType}}}}, "payment_method_options", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_mandate_url": types.StringType, "default_for": types.ListType{ElemType: types.StringType}, "interval_description": types.StringType, "payment_schedule": types.StringType, "transaction_type": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "verification_method": types.StringType}}, "affirm": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "afterpay_clearpay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "alipay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "alma": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "amazon_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "au_becs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "target_date": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "billie": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "boleto": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "installments": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "request_extended_authorization": types.StringType, "request_incremental_authorization": types.StringType, "request_multicapture": types.StringType, "request_overcapture": types.StringType, "request_three_d_secure": types.StringType, "restrictions": types.ObjectType{AttrTypes: map[string]attr.Type{"brands_blocked": types.ListType{ElemType: types.StringType}}}, "setup_future_usage": types.StringType, "statement_descriptor_suffix_kana": types.StringType, "statement_descriptor_suffix_kanji": types.StringType}}, "cashapp": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "requested_address_types": types.ListType{ElemType: types.StringType}, "type": types.StringType}}, "funding_type": types.StringType, "setup_future_usage": types.StringType}}, "eps": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "fpx": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "giropay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "grabpay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "ideal": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "kakao_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType, "subscriptions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "name": types.StringType, "next_billing": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "date": types.StringType}}, "reference": types.StringType}}}}}, "konbini": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "kr_card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "link": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "mobilepay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "multibanco": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "naver_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "oxxo": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "p24": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "tos_shown_and_accepted": types.BoolType}}, "payco": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "paynow": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "reference": types.StringType, "setup_future_usage": types.StringType, "risk_correlation_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "payments_per_period": types.Int64Type, "purpose": types.StringType, "start_date": types.StringType}}, "setup_future_usage": types.StringType}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_includes_iof": types.StringType, "expires_after_seconds": types.Int64Type, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "amount_type": types.StringType, "currency": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "reference": types.StringType, "start_date": types.StringType}}, "setup_future_usage": types.StringType}}, "revolut_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "samsung_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "satispay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "scalapay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "sofort": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "swish": types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType}}, "twint": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}, "setup_future_usage": types.StringType}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}, "return_url": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "verification_method": types.StringType}}, "crypto": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "demo_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "wechat_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"app_id": types.StringType, "client": types.StringType, "setup_future_usage": types.StringType}}}},
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
			if nullPaymentMethodOptions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_mandate_url": types.StringType, "default_for": types.ListType{ElemType: types.StringType}, "interval_description": types.StringType, "payment_schedule": types.StringType, "transaction_type": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "verification_method": types.StringType}}, "affirm": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "afterpay_clearpay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "alipay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "alma": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "amazon_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "au_becs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "target_date": types.StringType}}, "bacs_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "billie": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "boleto": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "installments": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "request_extended_authorization": types.StringType, "request_incremental_authorization": types.StringType, "request_multicapture": types.StringType, "request_overcapture": types.StringType, "request_three_d_secure": types.StringType, "restrictions": types.ObjectType{AttrTypes: map[string]attr.Type{"brands_blocked": types.ListType{ElemType: types.StringType}}}, "setup_future_usage": types.StringType, "statement_descriptor_suffix_kana": types.StringType, "statement_descriptor_suffix_kanji": types.StringType}}, "cashapp": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "requested_address_types": types.ListType{ElemType: types.StringType}, "type": types.StringType}}, "funding_type": types.StringType, "setup_future_usage": types.StringType}}, "eps": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "fpx": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "giropay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "grabpay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "ideal": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "kakao_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "klarna": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType, "subscriptions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "name": types.StringType, "next_billing": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "date": types.StringType}}, "reference": types.StringType}}}}}, "konbini": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "kr_card": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "link": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "mobilepay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "multibanco": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "naver_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "oxxo": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_days": types.Int64Type, "setup_future_usage": types.StringType}}, "p24": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType, "tos_shown_and_accepted": types.BoolType}}, "payco": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "paynow": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "paypal": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "preferred_locale": types.StringType, "reference": types.StringType, "setup_future_usage": types.StringType, "risk_correlation_id": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "payments_per_period": types.Int64Type, "purpose": types.StringType, "start_date": types.StringType}}, "setup_future_usage": types.StringType}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_includes_iof": types.StringType, "expires_after_seconds": types.Int64Type, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "amount_type": types.StringType, "currency": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType, "reference": types.StringType, "start_date": types.StringType}}, "setup_future_usage": types.StringType}}, "revolut_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "setup_future_usage": types.StringType}}, "samsung_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "satispay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "scalapay": types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType}}, "sepa_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"reference_prefix": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType}}, "sofort": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "swish": types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType}}, "twint": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}, "setup_future_usage": types.StringType}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}, "return_url": types.StringType}}, "setup_future_usage": types.StringType, "target_date": types.StringType, "verification_method": types.StringType}}, "crypto": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "demo_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"setup_future_usage": types.StringType}}, "wechat_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"app_id": types.StringType, "client": types.StringType, "setup_future_usage": types.StringType}}}}); ok {
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
		if rawValuePaymentStatus, rawOk := plainValueAtPath(raw, "payment_status"); rawOk {
			if valuePaymentStatus, err := flattenPlainValue(rawValuePaymentStatus, types.StringType, "payment_status", "raw response"); err != nil {
				return err
			} else {
				if typedPaymentStatus, ok := valuePaymentStatus.(types.String); ok {
					state.PaymentStatus = typedPaymentStatus
				}
			}
		} else if !hasRaw {
			if responseValuePaymentStatus, ok := plainFromResponseField(obj, "PaymentStatus"); ok {
				if valuePaymentStatus, err := flattenPlainValue(responseValuePaymentStatus, types.StringType, "payment_status", "response struct"); err != nil {
					return err
				} else {
					if typedPaymentStatus, ok := valuePaymentStatus.(types.String); ok {
						state.PaymentStatus = typedPaymentStatus
					}
				}
			}
		}
	}
	{
		assignedPermissions := false
		hadRawPermissions := false
		if rawValuePermissions, rawOk := plainValueAtPath(raw, "permissions"); rawOk {
			hadRawPermissions = true
			if rawValuePermissions != nil {
				sourcePermissions := applyConfiguredKeyedListShapes(rawValuePermissions, attrValueToPlain(state.Permissions))
				if valuePermissions, err := flattenPlainValue(sourcePermissions, types.ObjectType{AttrTypes: map[string]attr.Type{"update_shipping_details": types.StringType}}, "permissions", "raw response"); err != nil {
					return err
				} else {
					if typedPermissions, ok := valuePermissions.(types.Object); ok {
						state.Permissions = typedPermissions
						assignedPermissions = true
					}
				}
			}
		}
		if !assignedPermissions {
			if !hasRaw {
				if responseValuePermissions, ok := plainFromResponseField(obj, "Permissions"); ok {
					sourcePermissions := applyConfiguredKeyedListShapes(responseValuePermissions, attrValueToPlain(state.Permissions))
					if valuePermissions, err := flattenPlainValue(
						sourcePermissions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"update_shipping_details": types.StringType}},
						"permissions",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPermissions, ok := valuePermissions.(types.Object); ok {
							state.Permissions = typedPermissions
							assignedPermissions = true
						}
					}
				}
			}
		}
		if !assignedPermissions && hadRawPermissions {
			if nullPermissions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"update_shipping_details": types.StringType}}); ok {
				if typedPermissions, ok := nullPermissions.(types.Object); ok {
					state.Permissions = typedPermissions
				}
			}
		}
	}
	{
		assignedPhoneNumberCollection := false
		hadRawPhoneNumberCollection := false
		if rawValuePhoneNumberCollection, rawOk := plainValueAtPath(raw, "phone_number_collection"); rawOk {
			hadRawPhoneNumberCollection = true
			if rawValuePhoneNumberCollection != nil {
				sourcePhoneNumberCollection := applyConfiguredKeyedListShapes(rawValuePhoneNumberCollection, attrValueToPlain(state.PhoneNumberCollection))
				if valuePhoneNumberCollection, err := flattenPlainValue(sourcePhoneNumberCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "phone_number_collection", "raw response"); err != nil {
					return err
				} else {
					if typedPhoneNumberCollection, ok := valuePhoneNumberCollection.(types.Object); ok {
						state.PhoneNumberCollection = typedPhoneNumberCollection
						assignedPhoneNumberCollection = true
					}
				}
			}
		}
		if !assignedPhoneNumberCollection {
			if !hasRaw {
				if responseValuePhoneNumberCollection, ok := plainFromResponseField(obj, "PhoneNumberCollection"); ok {
					sourcePhoneNumberCollection := applyConfiguredKeyedListShapes(responseValuePhoneNumberCollection, attrValueToPlain(state.PhoneNumberCollection))
					if valuePhoneNumberCollection, err := flattenPlainValue(
						sourcePhoneNumberCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}},
						"phone_number_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPhoneNumberCollection, ok := valuePhoneNumberCollection.(types.Object); ok {
							state.PhoneNumberCollection = typedPhoneNumberCollection
							assignedPhoneNumberCollection = true
						}
					}
				}
			}
		}
		if !assignedPhoneNumberCollection && hadRawPhoneNumberCollection {
			if nullPhoneNumberCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}); ok {
				if typedPhoneNumberCollection, ok := nullPhoneNumberCollection.(types.Object); ok {
					state.PhoneNumberCollection = typedPhoneNumberCollection
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
		if rawValueRecoveredFrom, rawOk := plainValueAtPath(raw, "recovered_from"); rawOk {
			if valueRecoveredFrom, err := flattenPlainValue(rawValueRecoveredFrom, types.StringType, "recovered_from", "raw response"); err != nil {
				return err
			} else {
				if typedRecoveredFrom, ok := valueRecoveredFrom.(types.String); ok {
					state.RecoveredFrom = typedRecoveredFrom
				}
			}
		} else if !hasRaw {
			if responseValueRecoveredFrom, ok := plainFromResponseField(obj, "RecoveredFrom"); ok {
				if valueRecoveredFrom, err := flattenPlainValue(responseValueRecoveredFrom, types.StringType, "recovered_from", "response struct"); err != nil {
					return err
				} else {
					if typedRecoveredFrom, ok := valueRecoveredFrom.(types.String); ok {
						state.RecoveredFrom = typedRecoveredFrom
					}
				}
			}
		}
	}
	{
		if rawValueRedirectOnCompletion, rawOk := plainValueAtPath(raw, "redirect_on_completion"); rawOk {
			if valueRedirectOnCompletion, err := flattenPlainValue(rawValueRedirectOnCompletion, types.StringType, "redirect_on_completion", "raw response"); err != nil {
				return err
			} else {
				if typedRedirectOnCompletion, ok := valueRedirectOnCompletion.(types.String); ok {
					state.RedirectOnCompletion = typedRedirectOnCompletion
				}
			}
		} else if !hasRaw {
			if responseValueRedirectOnCompletion, ok := plainFromResponseField(obj, "RedirectOnCompletion"); ok {
				if valueRedirectOnCompletion, err := flattenPlainValue(responseValueRedirectOnCompletion, types.StringType, "redirect_on_completion", "response struct"); err != nil {
					return err
				} else {
					if typedRedirectOnCompletion, ok := valueRedirectOnCompletion.(types.String); ok {
						state.RedirectOnCompletion = typedRedirectOnCompletion
					}
				}
			}
		}
	}
	{
		if rawValueReturnURL, rawOk := plainValueAtPath(raw, "return_url"); rawOk {
			if valueReturnURL, err := flattenPlainValue(rawValueReturnURL, types.StringType, "return_url", "raw response"); err != nil {
				return err
			} else {
				if typedReturnURL, ok := valueReturnURL.(types.String); ok {
					state.ReturnURL = typedReturnURL
				}
			}
		} else if !hasRaw {
			if responseValueReturnURL, ok := plainFromResponseField(obj, "ReturnURL"); ok {
				if valueReturnURL, err := flattenPlainValue(responseValueReturnURL, types.StringType, "return_url", "response struct"); err != nil {
					return err
				} else {
					if typedReturnURL, ok := valueReturnURL.(types.String); ok {
						state.ReturnURL = typedReturnURL
					}
				}
			}
		}
	}
	{
		assignedSavedPaymentMethodOptions := false
		hadRawSavedPaymentMethodOptions := false
		if rawValueSavedPaymentMethodOptions, rawOk := plainValueAtPath(raw, "saved_payment_method_options"); rawOk {
			hadRawSavedPaymentMethodOptions = true
			if rawValueSavedPaymentMethodOptions != nil {
				sourceSavedPaymentMethodOptions := applyConfiguredKeyedListShapes(rawValueSavedPaymentMethodOptions, attrValueToPlain(state.SavedPaymentMethodOptions))
				if valueSavedPaymentMethodOptions, err := flattenPlainValue(sourceSavedPaymentMethodOptions, types.ObjectType{AttrTypes: map[string]attr.Type{"allow_redisplay_filters": types.ListType{ElemType: types.StringType}, "payment_method_remove": types.StringType, "payment_method_save": types.StringType}}, "saved_payment_method_options", "raw response"); err != nil {
					return err
				} else {
					if typedSavedPaymentMethodOptions, ok := valueSavedPaymentMethodOptions.(types.Object); ok {
						state.SavedPaymentMethodOptions = typedSavedPaymentMethodOptions
						assignedSavedPaymentMethodOptions = true
					}
				}
			}
		}
		if !assignedSavedPaymentMethodOptions {
			if !hasRaw {
				if responseValueSavedPaymentMethodOptions, ok := plainFromResponseField(obj, "SavedPaymentMethodOptions"); ok {
					sourceSavedPaymentMethodOptions := applyConfiguredKeyedListShapes(responseValueSavedPaymentMethodOptions, attrValueToPlain(state.SavedPaymentMethodOptions))
					if valueSavedPaymentMethodOptions, err := flattenPlainValue(
						sourceSavedPaymentMethodOptions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"allow_redisplay_filters": types.ListType{ElemType: types.StringType}, "payment_method_remove": types.StringType, "payment_method_save": types.StringType}},
						"saved_payment_method_options",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSavedPaymentMethodOptions, ok := valueSavedPaymentMethodOptions.(types.Object); ok {
							state.SavedPaymentMethodOptions = typedSavedPaymentMethodOptions
							assignedSavedPaymentMethodOptions = true
						}
					}
				}
			}
		}
		if !assignedSavedPaymentMethodOptions && hadRawSavedPaymentMethodOptions {
			if nullSavedPaymentMethodOptions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"allow_redisplay_filters": types.ListType{ElemType: types.StringType}, "payment_method_remove": types.StringType, "payment_method_save": types.StringType}}); ok {
				if typedSavedPaymentMethodOptions, ok := nullSavedPaymentMethodOptions.(types.Object); ok {
					state.SavedPaymentMethodOptions = typedSavedPaymentMethodOptions
				}
			}
		}
	}
	{
		if true {
			if rawValueSetupIntent, rawOk := plainValueAtPath(raw, "setup_intent"); rawOk {
				if typedSetupIntent, ok := plainToStringIDValue(rawValueSetupIntent); ok {
					state.SetupIntent = typedSetupIntent
				}
			} else if !hasRaw {
				if responseValueSetupIntent, ok := plainFromResponseField(obj, "SetupIntent"); ok {
					if typedSetupIntent, ok := plainToStringIDValue(responseValueSetupIntent); ok {
						state.SetupIntent = typedSetupIntent
					}
				}
			}
		}
	}
	{
		assignedShippingAddressCollection := false
		hadRawShippingAddressCollection := false
		if rawValueShippingAddressCollection, rawOk := plainValueAtPath(raw, "shipping_address_collection"); rawOk {
			hadRawShippingAddressCollection = true
			if rawValueShippingAddressCollection != nil {
				sourceShippingAddressCollection := applyConfiguredKeyedListShapes(rawValueShippingAddressCollection, attrValueToPlain(state.ShippingAddressCollection))
				if valueShippingAddressCollection, err := flattenPlainValue(sourceShippingAddressCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_countries": types.ListType{ElemType: types.StringType}}}, "shipping_address_collection", "raw response"); err != nil {
					return err
				} else {
					if typedShippingAddressCollection, ok := valueShippingAddressCollection.(types.Object); ok {
						state.ShippingAddressCollection = typedShippingAddressCollection
						assignedShippingAddressCollection = true
					}
				}
			}
		}
		if !assignedShippingAddressCollection {
			if !hasRaw {
				if responseValueShippingAddressCollection, ok := plainFromResponseField(obj, "ShippingAddressCollection"); ok {
					sourceShippingAddressCollection := applyConfiguredKeyedListShapes(responseValueShippingAddressCollection, attrValueToPlain(state.ShippingAddressCollection))
					if valueShippingAddressCollection, err := flattenPlainValue(
						sourceShippingAddressCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_countries": types.ListType{ElemType: types.StringType}}},
						"shipping_address_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedShippingAddressCollection, ok := valueShippingAddressCollection.(types.Object); ok {
							state.ShippingAddressCollection = typedShippingAddressCollection
							assignedShippingAddressCollection = true
						}
					}
				}
			}
		}
		if !assignedShippingAddressCollection && hadRawShippingAddressCollection {
			if nullShippingAddressCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_countries": types.ListType{ElemType: types.StringType}}}); ok {
				if typedShippingAddressCollection, ok := nullShippingAddressCollection.(types.Object); ok {
					state.ShippingAddressCollection = typedShippingAddressCollection
				}
			}
		}
	}
	{
		assignedShippingCost := false
		hadRawShippingCost := false
		if rawValueShippingCost, rawOk := plainValueAtPath(raw, "shipping_cost"); rawOk {
			hadRawShippingCost = true
			if rawValueShippingCost != nil {
				sourceShippingCost := applyConfiguredKeyedListShapes(rawValueShippingCost, attrValueToPlain(state.ShippingCost))
				if valueShippingCost, err := flattenPlainValue(sourceShippingCost, types.ObjectType{AttrTypes: map[string]attr.Type{"amount_subtotal": types.Int64Type, "amount_tax": types.Int64Type, "amount_total": types.Int64Type, "shipping_rate": types.StringType, "taxes": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "rate": types.StringType, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}}}, "shipping_cost", "raw response"); err != nil {
					return err
				} else {
					if typedShippingCost, ok := valueShippingCost.(types.Object); ok {
						state.ShippingCost = typedShippingCost
						assignedShippingCost = true
					}
				}
			}
		}
		if !assignedShippingCost {
			if !hasRaw {
				if responseValueShippingCost, ok := plainFromResponseField(obj, "ShippingCost"); ok {
					sourceShippingCost := applyConfiguredKeyedListShapes(responseValueShippingCost, attrValueToPlain(state.ShippingCost))
					if valueShippingCost, err := flattenPlainValue(
						sourceShippingCost,
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount_subtotal": types.Int64Type, "amount_tax": types.Int64Type, "amount_total": types.Int64Type, "shipping_rate": types.StringType, "taxes": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "rate": types.StringType, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}}},
						"shipping_cost",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedShippingCost, ok := valueShippingCost.(types.Object); ok {
							state.ShippingCost = typedShippingCost
							assignedShippingCost = true
						}
					}
				}
			}
		}
		if !assignedShippingCost && hadRawShippingCost {
			if nullShippingCost, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount_subtotal": types.Int64Type, "amount_tax": types.Int64Type, "amount_total": types.Int64Type, "shipping_rate": types.StringType, "taxes": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "rate": types.StringType, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}}}); ok {
				if typedShippingCost, ok := nullShippingCost.(types.Object); ok {
					state.ShippingCost = typedShippingCost
				}
			}
		}
	}
	{
		if rawValueShippingOptions, rawOk := plainValueAtPath(raw, "shipping_options"); rawOk {
			if valueShippingOptions, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueShippingOptions, attrValueToPlain(state.ShippingOptions)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"shipping_amount": types.Int64Type, "shipping_rate": types.StringType, "shipping_rate_data": types.ObjectType{AttrTypes: map[string]attr.Type{"delivery_estimate": types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}, "minimum": types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}}, "display_name": types.StringType, "fixed_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "currency_options": types.MapType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "tax_behavior": types.StringType}}}}}, "metadata": types.MapType{ElemType: types.StringType}, "tax_behavior": types.StringType, "tax_code": types.StringType, "type": types.StringType}}}}}, "shipping_options", "raw response"); err != nil {
				return err
			} else {
				if typedShippingOptions, ok := valueShippingOptions.(types.List); ok {
					state.ShippingOptions = typedShippingOptions
				}
			}
		} else if !hasRaw {
			if responseValueShippingOptions, ok := plainFromResponseField(obj, "ShippingOptions"); ok {
				if valueShippingOptions, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueShippingOptions, attrValueToPlain(state.ShippingOptions)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"shipping_amount": types.Int64Type, "shipping_rate": types.StringType, "shipping_rate_data": types.ObjectType{AttrTypes: map[string]attr.Type{"delivery_estimate": types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}, "minimum": types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}}, "display_name": types.StringType, "fixed_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "currency_options": types.MapType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "tax_behavior": types.StringType}}}}}, "metadata": types.MapType{ElemType: types.StringType}, "tax_behavior": types.StringType, "tax_code": types.StringType, "type": types.StringType}}}}},
					"shipping_options",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedShippingOptions, ok := valueShippingOptions.(types.List); ok {
						state.ShippingOptions = typedShippingOptions
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
		if rawValueSubmitType, rawOk := plainValueAtPath(raw, "submit_type"); rawOk {
			if valueSubmitType, err := flattenPlainValue(rawValueSubmitType, types.StringType, "submit_type", "raw response"); err != nil {
				return err
			} else {
				if typedSubmitType, ok := valueSubmitType.(types.String); ok {
					state.SubmitType = typedSubmitType
				}
			}
		} else if !hasRaw {
			if responseValueSubmitType, ok := plainFromResponseField(obj, "SubmitType"); ok {
				if valueSubmitType, err := flattenPlainValue(responseValueSubmitType, types.StringType, "submit_type", "response struct"); err != nil {
					return err
				} else {
					if typedSubmitType, ok := valueSubmitType.(types.String); ok {
						state.SubmitType = typedSubmitType
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueSubscription, rawOk := plainValueAtPath(raw, "subscription"); rawOk {
				if typedSubscription, ok := plainToStringIDValue(rawValueSubscription); ok {
					state.Subscription = typedSubscription
				}
			} else if !hasRaw {
				if responseValueSubscription, ok := plainFromResponseField(obj, "Subscription"); ok {
					if typedSubscription, ok := plainToStringIDValue(responseValueSubscription); ok {
						state.Subscription = typedSubscription
					}
				}
			}
		}
	}
	{
		if rawValueSuccessURL, rawOk := plainValueAtPath(raw, "success_url"); rawOk {
			if valueSuccessURL, err := flattenPlainValue(rawValueSuccessURL, types.StringType, "success_url", "raw response"); err != nil {
				return err
			} else {
				if typedSuccessURL, ok := valueSuccessURL.(types.String); ok {
					state.SuccessURL = typedSuccessURL
				}
			}
		} else if !hasRaw {
			if responseValueSuccessURL, ok := plainFromResponseField(obj, "SuccessURL"); ok {
				if valueSuccessURL, err := flattenPlainValue(responseValueSuccessURL, types.StringType, "success_url", "response struct"); err != nil {
					return err
				} else {
					if typedSuccessURL, ok := valueSuccessURL.(types.String); ok {
						state.SuccessURL = typedSuccessURL
					}
				}
			}
		}
	}
	{
		assignedTaxIDCollection := false
		hadRawTaxIDCollection := false
		if rawValueTaxIDCollection, rawOk := plainValueAtPath(raw, "tax_id_collection"); rawOk {
			hadRawTaxIDCollection = true
			if rawValueTaxIDCollection != nil {
				sourceTaxIDCollection := applyConfiguredKeyedListShapes(rawValueTaxIDCollection, attrValueToPlain(state.TaxIDCollection))
				if valueTaxIDCollection, err := flattenPlainValue(sourceTaxIDCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "required": types.StringType}}, "tax_id_collection", "raw response"); err != nil {
					return err
				} else {
					if typedTaxIDCollection, ok := valueTaxIDCollection.(types.Object); ok {
						state.TaxIDCollection = typedTaxIDCollection
						assignedTaxIDCollection = true
					}
				}
			}
		}
		if !assignedTaxIDCollection {
			if !hasRaw {
				if responseValueTaxIDCollection, ok := plainFromResponseField(obj, "TaxIDCollection"); ok {
					sourceTaxIDCollection := applyConfiguredKeyedListShapes(responseValueTaxIDCollection, attrValueToPlain(state.TaxIDCollection))
					if valueTaxIDCollection, err := flattenPlainValue(
						sourceTaxIDCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "required": types.StringType}},
						"tax_id_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTaxIDCollection, ok := valueTaxIDCollection.(types.Object); ok {
							state.TaxIDCollection = typedTaxIDCollection
							assignedTaxIDCollection = true
						}
					}
				}
			}
		}
		if !assignedTaxIDCollection && hadRawTaxIDCollection {
			if nullTaxIDCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "required": types.StringType}}); ok {
				if typedTaxIDCollection, ok := nullTaxIDCollection.(types.Object); ok {
					state.TaxIDCollection = typedTaxIDCollection
				}
			}
		}
	}
	{
		assignedTotalDetails := false
		hadRawTotalDetails := false
		if rawValueTotalDetails, rawOk := plainValueAtPath(raw, "total_details"); rawOk {
			hadRawTotalDetails = true
			if rawValueTotalDetails != nil {
				sourceTotalDetails := applyConfiguredKeyedListShapes(rawValueTotalDetails, attrValueToPlain(state.TotalDetails))
				if valueTotalDetails, err := flattenPlainValue(sourceTotalDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"amount_discount": types.Int64Type, "amount_shipping": types.Int64Type, "amount_tax": types.Int64Type, "breakdown": types.ObjectType{AttrTypes: map[string]attr.Type{"discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "discount": types.StringType}}}, "taxes": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "rate": types.StringType, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}}}}}, "total_details", "raw response"); err != nil {
					return err
				} else {
					if typedTotalDetails, ok := valueTotalDetails.(types.Object); ok {
						state.TotalDetails = typedTotalDetails
						assignedTotalDetails = true
					}
				}
			}
		}
		if !assignedTotalDetails {
			if !hasRaw {
				if responseValueTotalDetails, ok := plainFromResponseField(obj, "TotalDetails"); ok {
					sourceTotalDetails := applyConfiguredKeyedListShapes(responseValueTotalDetails, attrValueToPlain(state.TotalDetails))
					if valueTotalDetails, err := flattenPlainValue(
						sourceTotalDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount_discount": types.Int64Type, "amount_shipping": types.Int64Type, "amount_tax": types.Int64Type, "breakdown": types.ObjectType{AttrTypes: map[string]attr.Type{"discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "discount": types.StringType}}}, "taxes": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "rate": types.StringType, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}}}}},
						"total_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTotalDetails, ok := valueTotalDetails.(types.Object); ok {
							state.TotalDetails = typedTotalDetails
							assignedTotalDetails = true
						}
					}
				}
			}
		}
		if !assignedTotalDetails && hadRawTotalDetails {
			if nullTotalDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount_discount": types.Int64Type, "amount_shipping": types.Int64Type, "amount_tax": types.Int64Type, "breakdown": types.ObjectType{AttrTypes: map[string]attr.Type{"discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "discount": types.StringType}}}, "taxes": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "rate": types.StringType, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}}}}}); ok {
				if typedTotalDetails, ok := nullTotalDetails.(types.Object); ok {
					state.TotalDetails = typedTotalDetails
				}
			}
		}
	}
	{
		if rawValueUIMode, rawOk := plainValueAtPath(raw, "ui_mode"); rawOk {
			if valueUIMode, err := flattenPlainValue(rawValueUIMode, types.StringType, "ui_mode", "raw response"); err != nil {
				return err
			} else {
				if typedUIMode, ok := valueUIMode.(types.String); ok {
					state.UIMode = typedUIMode
				}
			}
		} else if !hasRaw {
			if responseValueUIMode, ok := plainFromResponseField(obj, "UIMode"); ok {
				if valueUIMode, err := flattenPlainValue(responseValueUIMode, types.StringType, "ui_mode", "response struct"); err != nil {
					return err
				} else {
					if typedUIMode, ok := valueUIMode.(types.String); ok {
						state.UIMode = typedUIMode
					}
				}
			}
		}
	}
	{
		if rawValueURL, rawOk := plainValueAtPath(raw, "url"); rawOk {
			if valueURL, err := flattenPlainValue(rawValueURL, types.StringType, "url", "raw response"); err != nil {
				return err
			} else {
				if typedURL, ok := valueURL.(types.String); ok {
					state.URL = typedURL
				}
			}
		} else if !hasRaw {
			if responseValueURL, ok := plainFromResponseField(obj, "URL"); ok {
				if valueURL, err := flattenPlainValue(responseValueURL, types.StringType, "url", "response struct"); err != nil {
					return err
				} else {
					if typedURL, ok := valueURL.(types.String); ok {
						state.URL = typedURL
					}
				}
			}
		}
	}
	{
		assignedWalletOptions := false
		hadRawWalletOptions := false
		if rawValueWalletOptions, rawOk := plainValueAtPath(raw, "wallet_options"); rawOk {
			hadRawWalletOptions = true
			if rawValueWalletOptions != nil {
				sourceWalletOptions := applyConfiguredKeyedListShapes(rawValueWalletOptions, attrValueToPlain(state.WalletOptions))
				if valueWalletOptions, err := flattenPlainValue(sourceWalletOptions, types.ObjectType{AttrTypes: map[string]attr.Type{"link": types.ObjectType{AttrTypes: map[string]attr.Type{"display": types.StringType}}}}, "wallet_options", "raw response"); err != nil {
					return err
				} else {
					if typedWalletOptions, ok := valueWalletOptions.(types.Object); ok {
						state.WalletOptions = typedWalletOptions
						assignedWalletOptions = true
					}
				}
			}
		}
		if !assignedWalletOptions {
			if !hasRaw {
				if responseValueWalletOptions, ok := plainFromResponseField(obj, "WalletOptions"); ok {
					sourceWalletOptions := applyConfiguredKeyedListShapes(responseValueWalletOptions, attrValueToPlain(state.WalletOptions))
					if valueWalletOptions, err := flattenPlainValue(
						sourceWalletOptions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"link": types.ObjectType{AttrTypes: map[string]attr.Type{"display": types.StringType}}}},
						"wallet_options",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedWalletOptions, ok := valueWalletOptions.(types.Object); ok {
							state.WalletOptions = typedWalletOptions
							assignedWalletOptions = true
						}
					}
				}
			}
		}
		if !assignedWalletOptions && hadRawWalletOptions {
			if nullWalletOptions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"link": types.ObjectType{AttrTypes: map[string]attr.Type{"display": types.StringType}}}}); ok {
				if typedWalletOptions, ok := nullWalletOptions.(types.Object); ok {
					state.WalletOptions = typedWalletOptions
				}
			}
		}
	}
	return nil
}

func (r *CheckoutSessionEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var config CheckoutSessionResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandCheckoutSessionCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building CheckoutSession ephemeral params", err.Error())
		return
	}

	obj, err := r.client.V1CheckoutSessions.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error opening CheckoutSession ephemeral resource", err.Error())
		return
	}

	result := config
	if err := flattenCheckoutSession(obj, &result); err != nil {
		resp.Diagnostics.AddError("Error flattening CheckoutSession ephemeral response", err.Error())
		return
	}
	normalizeUnknownValues(&result)
	diags = resp.Result.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
}
