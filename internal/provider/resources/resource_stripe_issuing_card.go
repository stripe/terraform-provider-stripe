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

var _ resource.Resource = &IssuingCardResource{}

var _ resource.ResourceWithConfigure = &IssuingCardResource{}

var _ resource.ResourceWithImportState = &IssuingCardResource{}

func NewIssuingCardResource() resource.Resource {
	return &IssuingCardResource{}
}

type IssuingCardResource struct {
	client *stripe.Client
}

type IssuingCardResourceModel struct {
	Object                types.String `tfsdk:"object"`
	Brand                 types.String `tfsdk:"brand"`
	Cardholder            types.String `tfsdk:"cardholder"`
	Created               types.Int64  `tfsdk:"created"`
	Currency              types.String `tfsdk:"currency"`
	CVC                   types.String `tfsdk:"cvc"`
	ExpMonth              types.Int64  `tfsdk:"exp_month"`
	ExpYear               types.Int64  `tfsdk:"exp_year"`
	FinancialAccount      types.String `tfsdk:"financial_account"`
	ID                    types.String `tfsdk:"id"`
	Last4                 types.String `tfsdk:"last4"`
	LatestFraudWarning    types.Object `tfsdk:"latest_fraud_warning"`
	LifecycleControls     types.Object `tfsdk:"lifecycle_controls"`
	Livemode              types.Bool   `tfsdk:"livemode"`
	Metadata              types.Map    `tfsdk:"metadata"`
	Number                types.String `tfsdk:"number"`
	PersonalizationDesign types.String `tfsdk:"personalization_design"`
	ReplacedBy            types.String `tfsdk:"replaced_by"`
	ReplacementFor        types.String `tfsdk:"replacement_for"`
	ReplacementReason     types.String `tfsdk:"replacement_reason"`
	Shipping              types.Object `tfsdk:"shipping"`
	SpendingControls      types.Object `tfsdk:"spending_controls"`
	Status                types.String `tfsdk:"status"`
	Type                  types.String `tfsdk:"type"`
	Wallets               types.Object `tfsdk:"wallets"`
	PIN                   types.Object `tfsdk:"pin"`
}

func (r *IssuingCardResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IssuingCardResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_issuing_card"
}

func (r *IssuingCardResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "You can [create physical or virtual cards](https://docs.stripe.com/issuing) that are issued to cardholders.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("issuing.card")},
			},
			"brand": schema.StringAttribute{
				Computed:      true,
				Description:   "The brand of the card.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cardholder": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An Issuing `Cardholder` object represents an individual or business entity who is [issued](https://docs.stripe.com/issuing) cards.\n\nRelated guide: [How to create a cardholder](https://docs.stripe.com/issuing/cards/virtual/issue-cards#create-cardholder)",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Required:      true,
				Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Supported currencies are `usd` in the US, `eur` in the EU, and `gbp` in the UK.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cvc": schema.StringAttribute{
				Computed:      true,
				Description:   "The card's CVC. For security reasons, this is only available for virtual cards, and will be omitted unless you explicitly request it with [the `expand` parameter](https://docs.stripe.com/api/expanding_objects). Additionally, it's only available via the [\"Retrieve a card\" endpoint](https://docs.stripe.com/api/issuing/cards/retrieve), not via \"List all cards\" or any other endpoint.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"exp_month": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The expiration month of the card.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"exp_year": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The expiration year of the card.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"financial_account": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The financial account this card is attached to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"last4": schema.StringAttribute{
				Computed:      true,
				Description:   "The last 4 digits of the card number.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"latest_fraud_warning": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Stripe’s assessment of whether this card’s details have been compromised. If this property isn't null, cancel and reissue the card to prevent fraudulent activity risk.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"started_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "Timestamp of the most recent fraud warning.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "The type of fraud warning that most recently took place on this card. This field updates with every new fraud warning, so the value changes over time. If populated, cancel and reissue the card.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("card_testing_exposure", "fraud_dispute_filed", "third_party_reported", "user_indicated_fraud")},
					},
				},
			},
			"lifecycle_controls": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Rules that control the lifecycle of this card, such as automatic cancellation. Refer to our [documentation](/issuing/controls/lifecycle-controls) for more details.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"cancel_after": schema.SingleNestedAttribute{
						Required: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"payment_count": schema.Int64Attribute{
								Required:      true,
								Description:   "The card is automatically cancelled when it makes this number of non-zero payment authorizations and transactions. The count includes penny authorizations, but doesn't include non-payment actions, such as authorization advice.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
						},
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
			"number": schema.StringAttribute{
				Computed:      true,
				Description:   "The full unredacted card number. For security reasons, this is only available for virtual cards, and will be omitted unless you explicitly request it with [the `expand` parameter](https://docs.stripe.com/api/expanding_objects). Additionally, it's only available via the [\"Retrieve a card\" endpoint](https://docs.stripe.com/api/issuing/cards/retrieve), not via \"List all cards\" or any other endpoint.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"personalization_design": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The personalization design object belonging to this card.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"replaced_by": schema.StringAttribute{
				Computed:      true,
				Description:   "The latest card that replaces this card, if any.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"replacement_for": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The card this card replaces, if any.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"replacement_reason": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The reason why the previous card needed to be replaced.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("damaged", "expired", "fulfillment_error", "lost", "stolen")},
			},
			"shipping": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Where and how the card will be shipped.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"address": schema.SingleNestedAttribute{
						Required: true,

						Attributes: map[string]schema.Attribute{
							"city": schema.StringAttribute{
								Required:    true,
								Description: "City, district, suburb, town, or village.",
							},
							"country": schema.StringAttribute{
								Required:    true,
								Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
							},
							"line1": schema.StringAttribute{
								Required:    true,
								Description: "Address line 1, such as the street, PO Box, or company name.",
							},
							"line2": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Address line 2, such as the apartment, suite, unit, or building.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"postal_code": schema.StringAttribute{
								Required:    true,
								Description: "ZIP or postal code.",
							},
							"state": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"address_validation": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Address validation details for the shipment.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Required:    true,
								Description: "The address validation capabilities to use.",
								Validators:  []validator.String{stringvalidator.OneOf("disabled", "normalization_only", "validation_and_normalization")},
							},
							"normalized_address": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "The normalized shipping address.",
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
							"result": schema.StringAttribute{
								Computed:      true,
								Description:   "The validation result for the shipping address.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("indeterminate", "likely_deliverable", "likely_undeliverable")},
							},
						},
					},
					"carrier": schema.StringAttribute{
						Computed:      true,
						Description:   "The delivery company that shipped a card.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("dhl", "fedex", "royal_mail", "usps")},
					},
					"customs": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Additional information that may be required for clearing customs.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"eori_number": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "A registration number used for customs in Europe. See [https://www.gov.uk/eori](https://www.gov.uk/eori) for the UK and [https://ec.europa.eu/taxation_customs/business/customs-procedures-import-and-export/customs-procedures/economic-operators-registration-and-identification-number-eori_en](https://ec.europa.eu/taxation_customs/business/customs-procedures-import-and-export/customs-procedures/economic-operators-registration-and-identification-number-eori_en) for the EU.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"eta": schema.Int64Attribute{
						Computed:      true,
						Description:   "A unix timestamp representing a best estimate of when the card will be delivered.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"name": schema.StringAttribute{
						Required:    true,
						Description: "Recipient name.",
					},
					"phone_number": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The phone number of the receiver of the shipment. Our courier partners will use this number to contact you in the event of card delivery issues. For individual shipments to the EU/UK, if this field is empty, we will provide them with the phone number provided when the cardholder was initially created.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"require_signature": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Whether a signature is required for card delivery. This feature is only supported for US users. Standard shipping service does not support signature on delivery. The default value for standard shipping service is false and for express and priority services is true.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"service": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Shipment service, such as `standard` or `express`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("express", "priority", "standard")},
					},
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The delivery status of the card.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("canceled", "delivered", "failure", "pending", "returned", "shipped", "submitted")},
					},
					"tracking_number": schema.StringAttribute{
						Computed:      true,
						Description:   "A tracking number for a card shipment.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"tracking_url": schema.StringAttribute{
						Computed:      true,
						Description:   "A link to the shipping carrier's site where you can view detailed information about a card shipment.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"type": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Packaging options.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("bulk", "individual")},
					},
				},
			},
			"spending_controls": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"allowed_card_presences": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of card presence statuses from which authorizations will be allowed. Possible options are `present`, `not_present`. All other statuses will be blocked. Cannot be set with `blocked_card_presences`. Provide an empty value to unset this control.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"allowed_categories": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of strings containing [categories](https://docs.stripe.com/api#issuing_authorization_object-merchant_data-category) of authorizations to allow. All other categories will be blocked. Cannot be set with `blocked_categories`.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"allowed_merchant_countries": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of strings containing representing countries from which authorizations will be allowed. Authorizations from merchants in all other countries will be declined. Country codes should be ISO 3166 alpha-2 country codes (e.g. `US`). Cannot be set with `blocked_merchant_countries`. Provide an empty value to unset this control.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"blocked_card_presences": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of card presence statuses from which authorizations will be declined. Possible options are `present`, `not_present`. Cannot be set with `allowed_card_presences`. Provide an empty value to unset this control.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"blocked_categories": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of strings containing [categories](https://docs.stripe.com/api#issuing_authorization_object-merchant_data-category) of authorizations to decline. All other categories will be allowed. Cannot be set with `allowed_categories`.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"blocked_merchant_countries": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of strings containing representing countries from which authorizations will be declined. Country codes should be ISO 3166 alpha-2 country codes (e.g. `US`). Cannot be set with `allowed_merchant_countries`. Provide an empty value to unset this control.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"spending_limits": schema.ListNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Limit spending with amount-based rules that apply across any cards this card replaced (i.e., its `replacement_for` card and _that_ card's `replacement_for` card, up the chain).",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"amount": schema.Int64Attribute{
									Required:    true,
									Description: "Maximum amount allowed to spend per interval. This amount is in the card's currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).",
								},
								"categories": schema.ListAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "Array of strings containing [categories](https://docs.stripe.com/api#issuing_authorization_object-merchant_data-category) this limit applies to. Omitting this field will apply the limit to all categories.",
									PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
									ElementType:   types.StringType,
								},
								"interval": schema.StringAttribute{
									Required:    true,
									Description: "Interval (or event) to which the amount applies.",
									Validators:  []validator.String{stringvalidator.OneOf("all_time", "daily", "monthly", "per_authorization", "weekly", "yearly")},
								},
							},
						},
					},
					"spending_limits_currency": schema.StringAttribute{
						Computed:      true,
						Description:   "Currency of the amounts within `spending_limits`. Always the same as the currency of the card.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"status": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether authorizations can be approved on this card. May be blocked from activating cards depending on past-due Cardholder requirements. Defaults to `inactive`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "canceled", "inactive")},
			},
			"type": schema.StringAttribute{
				Required:      true,
				Description:   "The type of the card.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("physical", "virtual")},
			},
			"wallets": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Information relating to digital wallets (like Apple Pay and Google Pay).",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"apple_pay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"eligible": schema.BoolAttribute{
								Computed:      true,
								Description:   "Apple Pay Eligibility",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"ineligible_reason": schema.StringAttribute{
								Computed:      true,
								Description:   "Reason the card is ineligible for Apple Pay",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("missing_agreement", "missing_cardholder_contact", "unsupported_region")},
							},
						},
					},
					"google_pay": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"eligible": schema.BoolAttribute{
								Computed:      true,
								Description:   "Google Pay Eligibility",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"ineligible_reason": schema.StringAttribute{
								Computed:      true,
								Description:   "Reason the card is ineligible for Google Pay",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("missing_agreement", "missing_cardholder_contact", "unsupported_region")},
							},
						},
					},
					"primary_account_identifier": schema.StringAttribute{
						Computed:      true,
						Description:   "Unique identifier for a card used with digital wallets",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"pin": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "The desired PIN for this card.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"encrypted_number": schema.StringAttribute{
						Optional:    true,
						Description: "The card's desired new PIN, encrypted under Stripe's public key.",
						WriteOnly:   true,
					},
				},
			},
		},
	}
}

func (r *IssuingCardResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IssuingCardResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config IssuingCardResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PIN"}, []string{"PIN", "encrypted_number"}})

	params, err := expandIssuingCardCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building IssuingCard create params", err.Error())
		return
	}

	obj, err := r.client.V1IssuingCards.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating IssuingCard", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingCards.B, r.client.V1IssuingCards.Key, stripe.FormatURLPath("/v1/issuing/cards/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingCard create raw response", err.Error())
		return
	}

	if err := flattenIssuingCard(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingCard create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"PIN"}, []string{"PIN", "encrypted_number"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingCardResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState IssuingCardResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state IssuingCardResourceModel
	state = priorState

	obj, err := r.client.V1IssuingCards.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading IssuingCard", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingCards.B, r.client.V1IssuingCards.Key, stripe.FormatURLPath("/v1/issuing/cards/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingCard raw response", err.Error())
		return
	}

	if err := flattenIssuingCard(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingCard read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"PIN"}, []string{"PIN", "encrypted_number"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingCardResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IssuingCardResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config IssuingCardResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PIN"}, []string{"PIN", "encrypted_number"}})

	var state IssuingCardResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandIssuingCardUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building IssuingCard update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building IssuingCard update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1IssuingCards.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating IssuingCard", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingCards.B, r.client.V1IssuingCards.Key, stripe.FormatURLPath("/v1/issuing/cards/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingCard update raw response", err.Error())
		return
	}

	if err := flattenIssuingCard(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingCard update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"PIN"}, []string{"PIN", "encrypted_number"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingCardResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *IssuingCardResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandIssuingCardCreate(plan IssuingCardResourceModel) (*stripe.IssuingCardCreateParams, error) {
	params := &stripe.IssuingCardCreateParams{}

	if !plan.Cardholder.IsNull() && !plan.Cardholder.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CardholderID", "Cardholder", plan.Cardholder.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "cardholder", params)
		}
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.ExpMonth.IsNull() && !plan.ExpMonth.IsUnknown() {
		params.ExpMonth = stripe.Int64(plan.ExpMonth.ValueInt64())
	}
	if !plan.ExpYear.IsNull() && !plan.ExpYear.IsUnknown() {
		params.ExpYear = stripe.Int64(plan.ExpYear.ValueInt64())
	}
	if !plan.FinancialAccount.IsNull() && !plan.FinancialAccount.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FinancialAccount", "FinancialAccount", plan.FinancialAccount.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "financial_account", params)
		}
	}
	if !plan.LifecycleControls.IsNull() && !plan.LifecycleControls.IsUnknown() {
		if !assignAttrValueToNamedField(params, "LifecycleControls", plan.LifecycleControls) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "lifecycle_controls", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.PersonalizationDesign.IsNull() && !plan.PersonalizationDesign.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PersonalizationDesignID", "PersonalizationDesign", plan.PersonalizationDesign.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "personalization_design", params)
		}
	}
	if !plan.ReplacementFor.IsNull() && !plan.ReplacementFor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReplacementForID", "ReplacementFor", plan.ReplacementFor.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "replacement_for", params)
		}
	}
	if !plan.ReplacementReason.IsNull() && !plan.ReplacementReason.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReplacementReason", "ReplacementReason", plan.ReplacementReason.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "replacement_reason", params)
		}
	}
	if !plan.Shipping.IsNull() && !plan.Shipping.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Shipping", plan.Shipping) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping", params)
		}
	}
	if !plan.SpendingControls.IsNull() && !plan.SpendingControls.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SpendingControls", plan.SpendingControls) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "spending_controls", params)
		}
	}
	if !plan.Status.IsNull() && !plan.Status.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Status", "Status", plan.Status.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "status", params)
		}
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}
	if !plan.PIN.IsNull() && !plan.PIN.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PIN", plan.PIN) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "pin", params)
		}
	}

	return params, nil
}

func expandIssuingCardUpdate(plan IssuingCardResourceModel, state IssuingCardResourceModel) (*stripe.IssuingCardUpdateParams, error) {
	params := &stripe.IssuingCardUpdateParams{}

	if !plan.Metadata.Equal(state.Metadata) && !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			if !plan.Metadata.Equal(state.Metadata) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "metadata", params)
			}
		}
	}
	if !plan.PersonalizationDesign.Equal(state.PersonalizationDesign) && !plan.PersonalizationDesign.IsNull() && !plan.PersonalizationDesign.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PersonalizationDesignID", "PersonalizationDesign", plan.PersonalizationDesign.ValueString()) {
			if !plan.PersonalizationDesign.Equal(state.PersonalizationDesign) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "personalization_design", params)
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
	if !plan.SpendingControls.Equal(state.SpendingControls) && !plan.SpendingControls.IsNull() && !plan.SpendingControls.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SpendingControls", plan.SpendingControls) {
			if !plan.SpendingControls.Equal(state.SpendingControls) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "spending_controls", params)
			}
		}
	}
	if !plan.Status.Equal(state.Status) && !plan.Status.IsNull() && !plan.Status.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Status", "Status", plan.Status.ValueString()) {
			if !plan.Status.Equal(state.Status) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "status", params)
			}
		}
	}
	if !plan.PIN.Equal(state.PIN) && !plan.PIN.IsNull() && !plan.PIN.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PIN", plan.PIN) {
			if !plan.PIN.Equal(state.PIN) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "pin", params)
			}
		}
	}

	return params, nil
}

func flattenIssuingCard(obj *stripe.IssuingCard, state *IssuingCardResourceModel) error {
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
		if rawValueBrand, rawOk := plainValueAtPath(raw, "brand"); rawOk {
			if valueBrand, err := flattenPlainValue(rawValueBrand, types.StringType, "brand", "raw response"); err != nil {
				return err
			} else {
				if typedBrand, ok := valueBrand.(types.String); ok {
					state.Brand = typedBrand
				}
			}
		} else if !hasRaw {
			if responseValueBrand, ok := plainFromResponseField(obj, "Brand"); ok {
				if valueBrand, err := flattenPlainValue(responseValueBrand, types.StringType, "brand", "response struct"); err != nil {
					return err
				} else {
					if typedBrand, ok := valueBrand.(types.String); ok {
						state.Brand = typedBrand
					}
				}
			}
		}
	}
	{
		if state.Cardholder.IsNull() || state.Cardholder.IsUnknown() {
			if rawValueCardholder, rawOk := plainValueAtPath(raw, "cardholder"); rawOk {
				if typedCardholder, ok := plainToStringIDValue(rawValueCardholder); ok {
					state.Cardholder = typedCardholder
				}
			} else if !hasRaw {
				if responseValueCardholder, ok := plainFromResponseField(obj, "Cardholder"); ok {
					if typedCardholder, ok := plainToStringIDValue(responseValueCardholder); ok {
						state.Cardholder = typedCardholder
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
		if rawValueCVC, rawOk := plainValueAtPath(raw, "cvc"); rawOk {
			if valueCVC, err := flattenPlainValue(rawValueCVC, types.StringType, "cvc", "raw response"); err != nil {
				return err
			} else {
				if typedCVC, ok := valueCVC.(types.String); ok {
					state.CVC = typedCVC
				}
			}
		} else if !hasRaw {
			if responseValueCVC, ok := plainFromResponseField(obj, "CVC"); ok {
				if valueCVC, err := flattenPlainValue(responseValueCVC, types.StringType, "cvc", "response struct"); err != nil {
					return err
				} else {
					if typedCVC, ok := valueCVC.(types.String); ok {
						state.CVC = typedCVC
					}
				}
			}
		}
	}
	{
		if rawValueExpMonth, rawOk := plainValueAtPath(raw, "exp_month"); rawOk {
			if valueExpMonth, err := flattenPlainValue(rawValueExpMonth, types.Int64Type, "exp_month", "raw response"); err != nil {
				return err
			} else {
				if typedExpMonth, ok := valueExpMonth.(types.Int64); ok {
					state.ExpMonth = typedExpMonth
				}
			}
		} else if !hasRaw {
			if responseValueExpMonth, ok := plainFromResponseField(obj, "ExpMonth"); ok {
				if valueExpMonth, err := flattenPlainValue(responseValueExpMonth, types.Int64Type, "exp_month", "response struct"); err != nil {
					return err
				} else {
					if typedExpMonth, ok := valueExpMonth.(types.Int64); ok {
						state.ExpMonth = typedExpMonth
					}
				}
			}
		}
	}
	{
		if rawValueExpYear, rawOk := plainValueAtPath(raw, "exp_year"); rawOk {
			if valueExpYear, err := flattenPlainValue(rawValueExpYear, types.Int64Type, "exp_year", "raw response"); err != nil {
				return err
			} else {
				if typedExpYear, ok := valueExpYear.(types.Int64); ok {
					state.ExpYear = typedExpYear
				}
			}
		} else if !hasRaw {
			if responseValueExpYear, ok := plainFromResponseField(obj, "ExpYear"); ok {
				if valueExpYear, err := flattenPlainValue(responseValueExpYear, types.Int64Type, "exp_year", "response struct"); err != nil {
					return err
				} else {
					if typedExpYear, ok := valueExpYear.(types.Int64); ok {
						state.ExpYear = typedExpYear
					}
				}
			}
		}
	}
	{
		if rawValueFinancialAccount, rawOk := plainValueAtPath(raw, "financial_account"); rawOk {
			if valueFinancialAccount, err := flattenPlainValue(rawValueFinancialAccount, types.StringType, "financial_account", "raw response"); err != nil {
				return err
			} else {
				if typedFinancialAccount, ok := valueFinancialAccount.(types.String); ok {
					state.FinancialAccount = typedFinancialAccount
				}
			}
		} else if !hasRaw {
			if responseValueFinancialAccount, ok := plainFromResponseField(obj, "FinancialAccount"); ok {
				if valueFinancialAccount, err := flattenPlainValue(responseValueFinancialAccount, types.StringType, "financial_account", "response struct"); err != nil {
					return err
				} else {
					if typedFinancialAccount, ok := valueFinancialAccount.(types.String); ok {
						state.FinancialAccount = typedFinancialAccount
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
		if rawValueLast4, rawOk := plainValueAtPath(raw, "last4"); rawOk {
			if valueLast4, err := flattenPlainValue(rawValueLast4, types.StringType, "last4", "raw response"); err != nil {
				return err
			} else {
				if typedLast4, ok := valueLast4.(types.String); ok {
					state.Last4 = typedLast4
				}
			}
		} else if !hasRaw {
			if responseValueLast4, ok := plainFromResponseField(obj, "Last4"); ok {
				if valueLast4, err := flattenPlainValue(responseValueLast4, types.StringType, "last4", "response struct"); err != nil {
					return err
				} else {
					if typedLast4, ok := valueLast4.(types.String); ok {
						state.Last4 = typedLast4
					}
				}
			}
		}
	}
	{
		assignedLatestFraudWarning := false
		hadRawLatestFraudWarning := false
		if rawValueLatestFraudWarning, rawOk := plainValueAtPath(raw, "latest_fraud_warning"); rawOk {
			hadRawLatestFraudWarning = true
			if rawValueLatestFraudWarning != nil {
				sourceLatestFraudWarning := applyConfiguredKeyedListShapes(rawValueLatestFraudWarning, attrValueToPlain(state.LatestFraudWarning))
				if valueLatestFraudWarning, err := flattenPlainValue(sourceLatestFraudWarning, types.ObjectType{AttrTypes: map[string]attr.Type{"started_at": types.Int64Type, "type": types.StringType}}, "latest_fraud_warning", "raw response"); err != nil {
					return err
				} else {
					if typedLatestFraudWarning, ok := valueLatestFraudWarning.(types.Object); ok {
						state.LatestFraudWarning = typedLatestFraudWarning
						assignedLatestFraudWarning = true
					}
				}
			}
		}
		if !assignedLatestFraudWarning {
			if !hasRaw {
				if responseValueLatestFraudWarning, ok := plainFromResponseField(obj, "LatestFraudWarning"); ok {
					sourceLatestFraudWarning := applyConfiguredKeyedListShapes(responseValueLatestFraudWarning, attrValueToPlain(state.LatestFraudWarning))
					if valueLatestFraudWarning, err := flattenPlainValue(
						sourceLatestFraudWarning,
						types.ObjectType{AttrTypes: map[string]attr.Type{"started_at": types.Int64Type, "type": types.StringType}},
						"latest_fraud_warning",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLatestFraudWarning, ok := valueLatestFraudWarning.(types.Object); ok {
							state.LatestFraudWarning = typedLatestFraudWarning
							assignedLatestFraudWarning = true
						}
					}
				}
			}
		}
		if !assignedLatestFraudWarning && hadRawLatestFraudWarning {
			if nullLatestFraudWarning, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"started_at": types.Int64Type, "type": types.StringType}}); ok {
				if typedLatestFraudWarning, ok := nullLatestFraudWarning.(types.Object); ok {
					state.LatestFraudWarning = typedLatestFraudWarning
				}
			}
		}
	}
	{
		assignedLifecycleControls := false
		hadRawLifecycleControls := false
		if rawValueLifecycleControls, rawOk := plainValueAtPath(raw, "lifecycle_controls"); rawOk {
			hadRawLifecycleControls = true
			if rawValueLifecycleControls != nil {
				sourceLifecycleControls := applyConfiguredKeyedListShapes(rawValueLifecycleControls, attrValueToPlain(state.LifecycleControls))
				if valueLifecycleControls, err := flattenPlainValue(sourceLifecycleControls, types.ObjectType{AttrTypes: map[string]attr.Type{"cancel_after": types.ObjectType{AttrTypes: map[string]attr.Type{"payment_count": types.Int64Type}}}}, "lifecycle_controls", "raw response"); err != nil {
					return err
				} else {
					if typedLifecycleControls, ok := valueLifecycleControls.(types.Object); ok {
						state.LifecycleControls = typedLifecycleControls
						assignedLifecycleControls = true
					}
				}
			}
		}
		if !assignedLifecycleControls {
			if !hasRaw {
				if responseValueLifecycleControls, ok := plainFromResponseField(obj, "LifecycleControls"); ok {
					sourceLifecycleControls := applyConfiguredKeyedListShapes(responseValueLifecycleControls, attrValueToPlain(state.LifecycleControls))
					if valueLifecycleControls, err := flattenPlainValue(
						sourceLifecycleControls,
						types.ObjectType{AttrTypes: map[string]attr.Type{"cancel_after": types.ObjectType{AttrTypes: map[string]attr.Type{"payment_count": types.Int64Type}}}},
						"lifecycle_controls",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLifecycleControls, ok := valueLifecycleControls.(types.Object); ok {
							state.LifecycleControls = typedLifecycleControls
							assignedLifecycleControls = true
						}
					}
				}
			}
		}
		if !assignedLifecycleControls && hadRawLifecycleControls {
			if nullLifecycleControls, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"cancel_after": types.ObjectType{AttrTypes: map[string]attr.Type{"payment_count": types.Int64Type}}}}); ok {
				if typedLifecycleControls, ok := nullLifecycleControls.(types.Object); ok {
					state.LifecycleControls = typedLifecycleControls
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
		if rawValueNumber, rawOk := plainValueAtPath(raw, "number"); rawOk {
			if valueNumber, err := flattenPlainValue(rawValueNumber, types.StringType, "number", "raw response"); err != nil {
				return err
			} else {
				if typedNumber, ok := valueNumber.(types.String); ok {
					state.Number = typedNumber
				}
			}
		} else if !hasRaw {
			if responseValueNumber, ok := plainFromResponseField(obj, "Number"); ok {
				if valueNumber, err := flattenPlainValue(responseValueNumber, types.StringType, "number", "response struct"); err != nil {
					return err
				} else {
					if typedNumber, ok := valueNumber.(types.String); ok {
						state.Number = typedNumber
					}
				}
			}
		}
	}
	{
		if true {
			if rawValuePersonalizationDesign, rawOk := plainValueAtPath(raw, "personalization_design"); rawOk {
				if typedPersonalizationDesign, ok := plainToStringIDValue(rawValuePersonalizationDesign); ok {
					state.PersonalizationDesign = typedPersonalizationDesign
				}
			} else if !hasRaw {
				if responseValuePersonalizationDesign, ok := plainFromResponseField(obj, "PersonalizationDesign"); ok {
					if typedPersonalizationDesign, ok := plainToStringIDValue(responseValuePersonalizationDesign); ok {
						state.PersonalizationDesign = typedPersonalizationDesign
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueReplacedBy, rawOk := plainValueAtPath(raw, "replaced_by"); rawOk {
				if typedReplacedBy, ok := plainToStringIDValue(rawValueReplacedBy); ok {
					state.ReplacedBy = typedReplacedBy
				}
			} else if !hasRaw {
				if responseValueReplacedBy, ok := plainFromResponseField(obj, "ReplacedBy"); ok {
					if typedReplacedBy, ok := plainToStringIDValue(responseValueReplacedBy); ok {
						state.ReplacedBy = typedReplacedBy
					}
				}
			}
		}
	}
	{
		if state.ReplacementFor.IsNull() || state.ReplacementFor.IsUnknown() {
			if rawValueReplacementFor, rawOk := plainValueAtPath(raw, "replacement_for"); rawOk {
				if typedReplacementFor, ok := plainToStringIDValue(rawValueReplacementFor); ok {
					state.ReplacementFor = typedReplacementFor
				}
			} else if !hasRaw {
				if responseValueReplacementFor, ok := plainFromResponseField(obj, "ReplacementFor"); ok {
					if typedReplacementFor, ok := plainToStringIDValue(responseValueReplacementFor); ok {
						state.ReplacementFor = typedReplacementFor
					}
				}
			}
		}
	}
	{
		if rawValueReplacementReason, rawOk := plainValueAtPath(raw, "replacement_reason"); rawOk {
			if valueReplacementReason, err := flattenPlainValue(rawValueReplacementReason, types.StringType, "replacement_reason", "raw response"); err != nil {
				return err
			} else {
				if typedReplacementReason, ok := valueReplacementReason.(types.String); ok {
					state.ReplacementReason = typedReplacementReason
				}
			}
		} else if !hasRaw {
			if responseValueReplacementReason, ok := plainFromResponseField(obj, "ReplacementReason"); ok {
				if valueReplacementReason, err := flattenPlainValue(responseValueReplacementReason, types.StringType, "replacement_reason", "response struct"); err != nil {
					return err
				} else {
					if typedReplacementReason, ok := valueReplacementReason.(types.String); ok {
						state.ReplacementReason = typedReplacementReason
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
				if valueShipping, err := flattenPlainValue(sourceShipping, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "address_validation": types.ObjectType{AttrTypes: map[string]attr.Type{"mode": types.StringType, "normalized_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "result": types.StringType}}, "carrier": types.StringType, "customs": types.ObjectType{AttrTypes: map[string]attr.Type{"eori_number": types.StringType}}, "eta": types.Int64Type, "name": types.StringType, "phone_number": types.StringType, "require_signature": types.BoolType, "service": types.StringType, "status": types.StringType, "tracking_number": types.StringType, "tracking_url": types.StringType, "type": types.StringType}}, "shipping", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "address_validation": types.ObjectType{AttrTypes: map[string]attr.Type{"mode": types.StringType, "normalized_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "result": types.StringType}}, "carrier": types.StringType, "customs": types.ObjectType{AttrTypes: map[string]attr.Type{"eori_number": types.StringType}}, "eta": types.Int64Type, "name": types.StringType, "phone_number": types.StringType, "require_signature": types.BoolType, "service": types.StringType, "status": types.StringType, "tracking_number": types.StringType, "tracking_url": types.StringType, "type": types.StringType}},
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
			if nullShipping, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "address_validation": types.ObjectType{AttrTypes: map[string]attr.Type{"mode": types.StringType, "normalized_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "result": types.StringType}}, "carrier": types.StringType, "customs": types.ObjectType{AttrTypes: map[string]attr.Type{"eori_number": types.StringType}}, "eta": types.Int64Type, "name": types.StringType, "phone_number": types.StringType, "require_signature": types.BoolType, "service": types.StringType, "status": types.StringType, "tracking_number": types.StringType, "tracking_url": types.StringType, "type": types.StringType}}); ok {
				if typedShipping, ok := nullShipping.(types.Object); ok {
					state.Shipping = typedShipping
				}
			}
		}
	}
	{
		assignedSpendingControls := false
		hadRawSpendingControls := false
		if rawValueSpendingControls, rawOk := plainValueAtPath(raw, "spending_controls"); rawOk {
			hadRawSpendingControls = true
			if rawValueSpendingControls != nil {
				sourceSpendingControls := applyConfiguredKeyedListShapes(rawValueSpendingControls, attrValueToPlain(state.SpendingControls))
				if valueSpendingControls, err := flattenPlainValue(sourceSpendingControls, types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_card_presences": types.ListType{ElemType: types.StringType}, "allowed_categories": types.ListType{ElemType: types.StringType}, "allowed_merchant_countries": types.ListType{ElemType: types.StringType}, "blocked_card_presences": types.ListType{ElemType: types.StringType}, "blocked_categories": types.ListType{ElemType: types.StringType}, "blocked_merchant_countries": types.ListType{ElemType: types.StringType}, "spending_limits": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "categories": types.ListType{ElemType: types.StringType}, "interval": types.StringType}}}, "spending_limits_currency": types.StringType}}, "spending_controls", "raw response"); err != nil {
					return err
				} else {
					if typedSpendingControls, ok := valueSpendingControls.(types.Object); ok {
						state.SpendingControls = typedSpendingControls
						assignedSpendingControls = true
					}
				}
			}
		}
		if !assignedSpendingControls {
			if !hasRaw {
				if responseValueSpendingControls, ok := plainFromResponseField(obj, "SpendingControls"); ok {
					sourceSpendingControls := applyConfiguredKeyedListShapes(responseValueSpendingControls, attrValueToPlain(state.SpendingControls))
					if valueSpendingControls, err := flattenPlainValue(
						sourceSpendingControls,
						types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_card_presences": types.ListType{ElemType: types.StringType}, "allowed_categories": types.ListType{ElemType: types.StringType}, "allowed_merchant_countries": types.ListType{ElemType: types.StringType}, "blocked_card_presences": types.ListType{ElemType: types.StringType}, "blocked_categories": types.ListType{ElemType: types.StringType}, "blocked_merchant_countries": types.ListType{ElemType: types.StringType}, "spending_limits": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "categories": types.ListType{ElemType: types.StringType}, "interval": types.StringType}}}, "spending_limits_currency": types.StringType}},
						"spending_controls",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSpendingControls, ok := valueSpendingControls.(types.Object); ok {
							state.SpendingControls = typedSpendingControls
							assignedSpendingControls = true
						}
					}
				}
			}
		}
		if !assignedSpendingControls && hadRawSpendingControls {
			if nullSpendingControls, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_card_presences": types.ListType{ElemType: types.StringType}, "allowed_categories": types.ListType{ElemType: types.StringType}, "allowed_merchant_countries": types.ListType{ElemType: types.StringType}, "blocked_card_presences": types.ListType{ElemType: types.StringType}, "blocked_categories": types.ListType{ElemType: types.StringType}, "blocked_merchant_countries": types.ListType{ElemType: types.StringType}, "spending_limits": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "categories": types.ListType{ElemType: types.StringType}, "interval": types.StringType}}}, "spending_limits_currency": types.StringType}}); ok {
				if typedSpendingControls, ok := nullSpendingControls.(types.Object); ok {
					state.SpendingControls = typedSpendingControls
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
		assignedWallets := false
		hadRawWallets := false
		if rawValueWallets, rawOk := plainValueAtPath(raw, "wallets"); rawOk {
			hadRawWallets = true
			if rawValueWallets != nil {
				sourceWallets := applyConfiguredKeyedListShapes(rawValueWallets, attrValueToPlain(state.Wallets))
				if valueWallets, err := flattenPlainValue(sourceWallets, types.ObjectType{AttrTypes: map[string]attr.Type{"apple_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"eligible": types.BoolType, "ineligible_reason": types.StringType}}, "google_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"eligible": types.BoolType, "ineligible_reason": types.StringType}}, "primary_account_identifier": types.StringType}}, "wallets", "raw response"); err != nil {
					return err
				} else {
					if typedWallets, ok := valueWallets.(types.Object); ok {
						state.Wallets = typedWallets
						assignedWallets = true
					}
				}
			}
		}
		if !assignedWallets {
			if !hasRaw {
				if responseValueWallets, ok := plainFromResponseField(obj, "Wallets"); ok {
					sourceWallets := applyConfiguredKeyedListShapes(responseValueWallets, attrValueToPlain(state.Wallets))
					if valueWallets, err := flattenPlainValue(
						sourceWallets,
						types.ObjectType{AttrTypes: map[string]attr.Type{"apple_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"eligible": types.BoolType, "ineligible_reason": types.StringType}}, "google_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"eligible": types.BoolType, "ineligible_reason": types.StringType}}, "primary_account_identifier": types.StringType}},
						"wallets",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedWallets, ok := valueWallets.(types.Object); ok {
							state.Wallets = typedWallets
							assignedWallets = true
						}
					}
				}
			}
		}
		if !assignedWallets && hadRawWallets {
			if nullWallets, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"apple_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"eligible": types.BoolType, "ineligible_reason": types.StringType}}, "google_pay": types.ObjectType{AttrTypes: map[string]attr.Type{"eligible": types.BoolType, "ineligible_reason": types.StringType}}, "primary_account_identifier": types.StringType}}); ok {
				if typedWallets, ok := nullWallets.(types.Object); ok {
					state.Wallets = typedWallets
				}
			}
		}
	}
	return nil
}
