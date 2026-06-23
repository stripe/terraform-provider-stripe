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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
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

var _ resource.Resource = &QuoteResource{}

var _ resource.ResourceWithConfigure = &QuoteResource{}

var _ resource.ResourceWithImportState = &QuoteResource{}

func NewQuoteResource() resource.Resource {
	return &QuoteResource{}
}

type QuoteResource struct {
	client *stripe.Client
}

type QuoteResourceModel struct {
	Object                types.String  `tfsdk:"object"`
	Application           types.String  `tfsdk:"application"`
	ApplicationFeeAmount  types.Int64   `tfsdk:"application_fee_amount"`
	ApplicationFeePercent types.Float64 `tfsdk:"application_fee_percent"`
	AutomaticTax          types.Object  `tfsdk:"automatic_tax"`
	CollectionMethod      types.String  `tfsdk:"collection_method"`
	Created               types.Int64   `tfsdk:"created"`
	Currency              types.String  `tfsdk:"currency"`
	Customer              types.String  `tfsdk:"customer"`
	CustomerAccount       types.String  `tfsdk:"customer_account"`
	DefaultTaxRates       types.List    `tfsdk:"default_tax_rates"`
	Description           types.String  `tfsdk:"description"`
	Discounts             types.List    `tfsdk:"discounts"`
	ExpiresAt             types.Int64   `tfsdk:"expires_at"`
	Footer                types.String  `tfsdk:"footer"`
	FromQuote             types.Object  `tfsdk:"from_quote"`
	Header                types.String  `tfsdk:"header"`
	ID                    types.String  `tfsdk:"id"`
	Invoice               types.String  `tfsdk:"invoice"`
	InvoiceSettings       types.Object  `tfsdk:"invoice_settings"`
	LineItems             types.List    `tfsdk:"line_items"`
	Livemode              types.Bool    `tfsdk:"livemode"`
	Metadata              types.Map     `tfsdk:"metadata"`
	Number                types.String  `tfsdk:"number"`
	OnBehalfOf            types.String  `tfsdk:"on_behalf_of"`
	Status                types.String  `tfsdk:"status"`
	StatusTransitions     types.Object  `tfsdk:"status_transitions"`
	Subscription          types.String  `tfsdk:"subscription"`
	SubscriptionData      types.Object  `tfsdk:"subscription_data"`
	SubscriptionSchedule  types.String  `tfsdk:"subscription_schedule"`
	TestClock             types.String  `tfsdk:"test_clock"`
	TotalDetails          types.Object  `tfsdk:"total_details"`
	TransferData          types.Object  `tfsdk:"transfer_data"`
}

func (r *QuoteResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *QuoteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quote"
}

func (r *QuoteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Quote is a way to model prices that you'd like to provide to a customer.\nOnce accepted, it will automatically create an invoice, subscription or subscription schedule.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("quote")},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the Connect Application that created the quote.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"application_fee_amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account. Only applicable if there are no line items with recurring prices on the quote.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"application_fee_percent": schema.Float64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account. Only applicable if there are line items with recurring prices on the quote.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
			},
			"automatic_tax": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Required:    true,
						Description: "Automatically calculate taxes",
					},
					"liability": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The account that's liable for tax. If set, the business address and tax registrations required to perform the tax calculation are loaded from this account. The tax transaction is returned in the report of the connected account.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The connected account being referenced when `type` is `account`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"type": schema.StringAttribute{
								Required:    true,
								Description: "Type of the account referenced.",
								Validators:  []validator.String{stringvalidator.OneOf("account", "self")},
							},
						},
					},
					"provider": schema.StringAttribute{
						Computed:      true,
						Description:   "The tax provider powering automatic tax.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the most recent automated tax calculation for this quote.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("complete", "failed", "requires_location_inputs")},
					},
				},
			},
			"collection_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay invoices at the end of the subscription cycle or on finalization using the default payment method attached to the subscription or customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`. Defaults to `charge_automatically`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("charge_automatically", "send_invoice")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Computed:      true,
				Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer who received this quote. A customer is required to finalize the quote. Once specified, you can't change it.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_account": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account representing the customer who received this quote. A customer or account is required to finalize the quote. Once specified, you can't change it.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_tax_rates": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The tax rates applied to this quote.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A description that will be displayed on the quote PDF.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"discounts": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The discounts applied to this quote.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"coupon": schema.StringAttribute{
							Optional:    true,
							Description: "ID of the coupon to create a new discount for.",
						},
						"discount": schema.StringAttribute{
							Optional:    true,
							Description: "ID of an existing discount on the object (or one of its ancestors) to reuse.",
						},
						"promotion_code": schema.StringAttribute{
							Optional:    true,
							Description: "ID of the promotion code to create a new discount for.",
						},
					},
				},
			},
			"expires_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The date on which the quote will be canceled if in `open` or `draft` status. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"footer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A footer that will be displayed on the quote PDF.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"from_quote": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Details of the quote that was cloned. See the [cloning documentation](https://docs.stripe.com/quotes/clone) for more details.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"is_revision": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Whether this quote is a revision of a different quote.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown(), boolplanmodifier.RequiresReplace()},
					},
					"quote": schema.StringAttribute{
						Required:      true,
						Description:   "The quote that was cloned.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"header": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A header that will be displayed on the quote PDF.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice": schema.StringAttribute{
				Computed:      true,
				Description:   "The invoice that was created from this quote.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice_settings": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"days_until_due": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "Number of days within which a customer must pay invoices generated by this quote. This value will be `null` for quotes where `collection_method=charge_automatically`.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"issuer": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The connected account being referenced when `type` is `account`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"type": schema.StringAttribute{
								Required:    true,
								Description: "Type of the account referenced.",
								Validators:  []validator.String{stringvalidator.OneOf("account", "self")},
							},
						},
					},
				},
			},
			"line_items": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A list of items the customer is being quoted for.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"discounts": schema.ListNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The discounts applied to this line item.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"coupon": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "ID of the coupon to create a new discount for.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"discount": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "ID of an existing discount on the object (or one of its ancestors) to reuse.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"promotion_code": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "ID of the promotion code to create a new discount for.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
						"price": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The ID of the price object. One of `price` or `price_data` is required.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"price_data": schema.SingleNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"currency": schema.StringAttribute{
									Required:    true,
									Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
								},
								"product": schema.StringAttribute{
									Required:    true,
									Description: "The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to.",
								},
								"recurring": schema.SingleNestedAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "The recurring components of a price such as `interval` and `interval_count`.",
									PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
									Attributes: map[string]schema.Attribute{
										"interval": schema.StringAttribute{
											Required:    true,
											Description: "Specifies billing frequency. Either `day`, `week`, `month` or `year`.",
										},
										"interval_count": schema.Int64Attribute{
											Optional:      true,
											Computed:      true,
											Description:   "The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).",
											PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
										},
									},
								},
								"tax_behavior": schema.StringAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"unit_amount": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"unit_amount_decimal": schema.Float64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
									PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
								},
							},
						},
						"quantity": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The quantity of the line item.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"tax_rates": schema.ListAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The tax rates which apply to the line item. When set, the `default_tax_rates` on the quote do not apply to this line item.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							ElementType:   types.StringType,
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
				Description:   "A unique number that identifies this particular quote. This number is assigned once the quote is [finalized](https://docs.stripe.com/quotes/overview#finalize).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"on_behalf_of": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account on behalf of which to charge. See the [Connect documentation](https://support.stripe.com/questions/sending-invoices-on-behalf-of-connected-accounts) for details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The status of the quote.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("accepted", "canceled", "draft", "open")},
			},
			"status_transitions": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"accepted_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time that the quote was accepted. Measured in seconds since Unix epoch.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"canceled_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time that the quote was canceled. Measured in seconds since Unix epoch.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"finalized_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time that the quote was finalized. Measured in seconds since Unix epoch.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"subscription": schema.StringAttribute{
				Computed:      true,
				Description:   "The subscription that was created or updated from this quote.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"subscription_data": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"billing_mode": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The billing mode of the quote.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"flexible": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"proration_discounts": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Controls how invoices and invoice items display proration amounts and discount amounts.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("included", "itemized")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Controls how prorations and invoices for subscriptions are calculated and orchestrated.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("classic", "flexible")},
							},
						},
					},
					"description": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"effective_date": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "When creating a new subscription, the date of which the subscription schedule will start after the quote is accepted. This date is ignored if it is in the past when the quote is accepted. Measured in seconds since the Unix epoch.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"metadata": schema.MapAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will set metadata on the subscription or subscription schedule when the quote is accepted. If a recurring price is included in `line_items`, this field will be passed to the resulting subscription's `metadata` field. If `subscription_data.effective_date` is used, this field will be passed to the resulting subscription schedule's `phases.metadata` field. Unlike object-level metadata, this field is declarative. Updates will clear prior values.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"trial_period_days": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "Integer representing the number of trial period days before the customer is charged for the first time.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"subscription_schedule": schema.StringAttribute{
				Computed:      true,
				Description:   "The subscription schedule that was created or updated from this quote.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"test_clock": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the test clock this quote belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"total_details": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"amount_discount": schema.Int64Attribute{
						Computed:      true,
						Description:   "This is the sum of all the discounts.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"amount_shipping": schema.Int64Attribute{
						Computed:      true,
						Description:   "This is the sum of all the shipping amounts.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"amount_tax": schema.Int64Attribute{
						Computed:      true,
						Description:   "This is the sum of all the tax amounts.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"breakdown": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"discounts": schema.ListNestedAttribute{
								Computed:      true,
								Description:   "The aggregated discounts.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"amount": schema.Int64Attribute{
											Computed:      true,
											Description:   "The amount discounted.",
											PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
										},
										"discount": schema.StringAttribute{
											Computed:      true,
											Description:   "A discount represents the actual application of a [coupon](https://api.stripe.com#coupons) or [promotion code](https://api.stripe.com#promotion_codes).\nIt contains information about when the discount began, when it will end, and what it is applied to.\n\nRelated guide: [Applying discounts to subscriptions](https://docs.stripe.com/billing/subscriptions/discounts)",
											PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										},
									},
								},
							},
							"taxes": schema.ListNestedAttribute{
								Computed:      true,
								Description:   "The aggregated tax amounts by rate.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"amount": schema.Int64Attribute{
											Computed:      true,
											Description:   "Amount of tax applied for this rate.",
											PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
										},
										"rate": schema.StringAttribute{
											Computed:      true,
											Description:   "Tax rates can be applied to [invoices](/invoicing/taxes/tax-rates), [subscriptions](/billing/taxes/tax-rates) and [Checkout Sessions](/payments/checkout/use-manual-tax-rates) to collect tax.\n\nRelated guide: [Tax rates](/billing/taxes/tax-rates)",
											PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										},
										"taxability_reason": schema.StringAttribute{
											Computed:      true,
											Description:   "The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.",
											PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											Validators:    []validator.String{stringvalidator.OneOf("customer_exempt", "not_collecting", "not_subject_to_tax", "not_supported", "portion_product_exempt", "portion_reduced_rated", "portion_standard_rated", "product_exempt", "product_exempt_holiday", "proportionally_rated", "reduced_rated", "reverse_charge", "standard_rated", "taxable_basis_reduced", "zero_rated")},
										},
										"taxable_amount": schema.Int64Attribute{
											Computed:      true,
											Description:   "The amount on which tax is calculated, in cents (or local equivalent).",
											PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
										},
									},
								},
							},
						},
					},
				},
			},
			"transfer_data": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the invoices.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"amount": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The amount in cents (or local equivalent) that will be transferred to the destination account when the invoice is paid. By default, the entire amount is transferred to the destination.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"amount_percent": schema.Float64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount will be transferred to the destination.",
						PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
					},
					"destination": schema.StringAttribute{
						Required:    true,
						Description: "The account where funds from the payment will be transferred to upon payment success.",
					},
				},
			},
		},
	}
}

func (r *QuoteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan QuoteResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config QuoteResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandQuoteCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Quote create params", err.Error())
		return
	}

	obj, err := r.client.V1Quotes.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Quote", err.Error())
		return
	}

	rawReadParams := &stripe.QuoteRetrieveParams{}
	rawReadParams.AddExpand("line_items")

	if err := ensureRawResponse(obj, r.client.V1Quotes.B, r.client.V1Quotes.Key, stripe.FormatURLPath("/v1/quotes/%s", obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating Quote create raw response", err.Error())
		return
	}

	if err := flattenQuote(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Quote create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}, []string{"LineItems", "*", "discounts"}, []string{"LineItems", "*", "discounts", "*", "coupon"}, []string{"LineItems", "*", "discounts", "*", "discount"}, []string{"LineItems", "*", "discounts", "*", "promotion_code"}, []string{"LineItems", "*", "price"}, []string{"LineItems", "*", "price_data"}, []string{"LineItems", "*", "price_data", "currency"}, []string{"LineItems", "*", "price_data", "product"}, []string{"LineItems", "*", "price_data", "recurring"}, []string{"LineItems", "*", "price_data", "recurring", "interval"}, []string{"LineItems", "*", "price_data", "recurring", "interval_count"}, []string{"LineItems", "*", "price_data", "tax_behavior"}, []string{"LineItems", "*", "price_data", "unit_amount"}, []string{"LineItems", "*", "price_data", "unit_amount_decimal"}, []string{"LineItems", "*", "quantity"}, []string{"LineItems", "*", "tax_rates"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *QuoteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState QuoteResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state QuoteResourceModel
	state = priorState

	params := &stripe.QuoteRetrieveParams{}
	params.AddExpand("line_items")

	obj, err := r.client.V1Quotes.Retrieve(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Quote", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Quotes.B, r.client.V1Quotes.Key, stripe.FormatURLPath("/v1/quotes/%s", state.ID.ValueString()), params); err != nil {
		resp.Diagnostics.AddError("Error hydrating Quote raw response", err.Error())
		return
	}

	if err := flattenQuote(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Quote read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}, []string{"LineItems", "*", "discounts"}, []string{"LineItems", "*", "discounts", "*", "coupon"}, []string{"LineItems", "*", "discounts", "*", "discount"}, []string{"LineItems", "*", "discounts", "*", "promotion_code"}, []string{"LineItems", "*", "price"}, []string{"LineItems", "*", "price_data"}, []string{"LineItems", "*", "price_data", "currency"}, []string{"LineItems", "*", "price_data", "product"}, []string{"LineItems", "*", "price_data", "recurring"}, []string{"LineItems", "*", "price_data", "recurring", "interval"}, []string{"LineItems", "*", "price_data", "recurring", "interval_count"}, []string{"LineItems", "*", "price_data", "tax_behavior"}, []string{"LineItems", "*", "price_data", "unit_amount"}, []string{"LineItems", "*", "price_data", "unit_amount_decimal"}, []string{"LineItems", "*", "quantity"}, []string{"LineItems", "*", "tax_rates"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *QuoteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan QuoteResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config QuoteResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state QuoteResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandQuoteUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Quote update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Quote update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Quotes.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Quote", err.Error())
		return
	}

	rawReadParams := &stripe.QuoteRetrieveParams{}
	rawReadParams.AddExpand("line_items")

	if err := ensureRawResponse(obj, r.client.V1Quotes.B, r.client.V1Quotes.Key, stripe.FormatURLPath("/v1/quotes/%s", obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating Quote update raw response", err.Error())
		return
	}

	if err := flattenQuote(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Quote update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}, []string{"LineItems", "*", "discounts"}, []string{"LineItems", "*", "discounts", "*", "coupon"}, []string{"LineItems", "*", "discounts", "*", "discount"}, []string{"LineItems", "*", "discounts", "*", "promotion_code"}, []string{"LineItems", "*", "price"}, []string{"LineItems", "*", "price_data"}, []string{"LineItems", "*", "price_data", "currency"}, []string{"LineItems", "*", "price_data", "product"}, []string{"LineItems", "*", "price_data", "recurring"}, []string{"LineItems", "*", "price_data", "recurring", "interval"}, []string{"LineItems", "*", "price_data", "recurring", "interval_count"}, []string{"LineItems", "*", "price_data", "tax_behavior"}, []string{"LineItems", "*", "price_data", "unit_amount"}, []string{"LineItems", "*", "price_data", "unit_amount_decimal"}, []string{"LineItems", "*", "quantity"}, []string{"LineItems", "*", "tax_rates"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *QuoteResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *QuoteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandQuoteCreate(plan QuoteResourceModel) (*stripe.QuoteCreateParams, error) {
	params := &stripe.QuoteCreateParams{}

	if !plan.ApplicationFeeAmount.IsNull() && !plan.ApplicationFeeAmount.IsUnknown() {
		params.ApplicationFeeAmount = stripe.Int64(plan.ApplicationFeeAmount.ValueInt64())
	}
	if !plan.ApplicationFeePercent.IsNull() && !plan.ApplicationFeePercent.IsUnknown() {
		params.ApplicationFeePercent = stripe.Float64(plan.ApplicationFeePercent.ValueFloat64())
	}
	if !plan.AutomaticTax.IsNull() && !plan.AutomaticTax.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AutomaticTax", plan.AutomaticTax) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "automatic_tax", params)
		}
	}
	if !plan.CollectionMethod.IsNull() && !plan.CollectionMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CollectionMethod", "CollectionMethod", plan.CollectionMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "collection_method", params)
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
	if !plan.DefaultTaxRates.IsNull() && !plan.DefaultTaxRates.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DefaultTaxRates", plan.DefaultTaxRates) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "default_tax_rates", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.Discounts.IsNull() && !plan.Discounts.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Discounts", plan.Discounts) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "discounts", params)
		}
	}
	if !plan.ExpiresAt.IsNull() && !plan.ExpiresAt.IsUnknown() {
		params.ExpiresAt = stripe.Int64(plan.ExpiresAt.ValueInt64())
	}
	if !plan.Footer.IsNull() && !plan.Footer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Footer", "Footer", plan.Footer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "footer", params)
		}
	}
	if !plan.FromQuote.IsNull() && !plan.FromQuote.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FromQuote", plan.FromQuote) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "from_quote", params)
		}
	}
	if !plan.Header.IsNull() && !plan.Header.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Header", "Header", plan.Header.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "header", params)
		}
	}
	if !plan.InvoiceSettings.IsNull() && !plan.InvoiceSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "InvoiceSettings", plan.InvoiceSettings) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "invoice_settings", params)
		}
	}
	if !plan.LineItems.IsNull() && !plan.LineItems.IsUnknown() {
		if !assignAttrValueToNamedField(params, "LineItems", plan.LineItems) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "line_items", params)
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
	if !plan.SubscriptionData.IsNull() && !plan.SubscriptionData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SubscriptionData", plan.SubscriptionData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "subscription_data", params)
		}
	}
	if !plan.TestClock.IsNull() && !plan.TestClock.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TestClockID", "TestClock", plan.TestClock.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "test_clock", params)
		}
	}
	if !plan.TransferData.IsNull() && !plan.TransferData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TransferData", plan.TransferData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "transfer_data", params)
		}
	}

	return params, nil
}

func expandQuoteUpdate(plan QuoteResourceModel, state QuoteResourceModel) (*stripe.QuoteUpdateParams, error) {
	params := &stripe.QuoteUpdateParams{}

	if !plan.ApplicationFeeAmount.Equal(state.ApplicationFeeAmount) && !plan.ApplicationFeeAmount.IsNull() && !plan.ApplicationFeeAmount.IsUnknown() {
		params.ApplicationFeeAmount = stripe.Int64(plan.ApplicationFeeAmount.ValueInt64())
	}
	if !plan.ApplicationFeePercent.Equal(state.ApplicationFeePercent) && !plan.ApplicationFeePercent.IsNull() && !plan.ApplicationFeePercent.IsUnknown() {
		params.ApplicationFeePercent = stripe.Float64(plan.ApplicationFeePercent.ValueFloat64())
	}
	if !plan.AutomaticTax.Equal(state.AutomaticTax) && !plan.AutomaticTax.IsNull() && !plan.AutomaticTax.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AutomaticTax", plan.AutomaticTax) {
			if !plan.AutomaticTax.Equal(state.AutomaticTax) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "automatic_tax", params)
			}
		}
	}
	if !plan.CollectionMethod.Equal(state.CollectionMethod) && !plan.CollectionMethod.IsNull() && !plan.CollectionMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CollectionMethod", "CollectionMethod", plan.CollectionMethod.ValueString()) {
			if !plan.CollectionMethod.Equal(state.CollectionMethod) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "collection_method", params)
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
	if !plan.CustomerAccount.Equal(state.CustomerAccount) && !plan.CustomerAccount.IsNull() && !plan.CustomerAccount.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerAccount", "CustomerAccount", plan.CustomerAccount.ValueString()) {
			if !plan.CustomerAccount.Equal(state.CustomerAccount) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "customer_account", params)
			}
		}
	}
	if !plan.DefaultTaxRates.Equal(state.DefaultTaxRates) && !plan.DefaultTaxRates.IsNull() && !plan.DefaultTaxRates.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DefaultTaxRates", plan.DefaultTaxRates) {
			if !plan.DefaultTaxRates.Equal(state.DefaultTaxRates) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "default_tax_rates", params)
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
	if !plan.Discounts.Equal(state.Discounts) && !plan.Discounts.IsNull() && !plan.Discounts.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Discounts", plan.Discounts) {
			if !plan.Discounts.Equal(state.Discounts) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "discounts", params)
			}
		}
	}
	if !plan.ExpiresAt.Equal(state.ExpiresAt) && !plan.ExpiresAt.IsNull() && !plan.ExpiresAt.IsUnknown() {
		params.ExpiresAt = stripe.Int64(plan.ExpiresAt.ValueInt64())
	}
	if !plan.Footer.Equal(state.Footer) && !plan.Footer.IsNull() && !plan.Footer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Footer", "Footer", plan.Footer.ValueString()) {
			if !plan.Footer.Equal(state.Footer) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "footer", params)
			}
		}
	}
	if !plan.Header.Equal(state.Header) && !plan.Header.IsNull() && !plan.Header.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Header", "Header", plan.Header.ValueString()) {
			if !plan.Header.Equal(state.Header) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "header", params)
			}
		}
	}
	if !plan.InvoiceSettings.Equal(state.InvoiceSettings) && !plan.InvoiceSettings.IsNull() && !plan.InvoiceSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "InvoiceSettings", plan.InvoiceSettings) {
			if !plan.InvoiceSettings.Equal(state.InvoiceSettings) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "invoice_settings", params)
			}
		}
	}
	if !plan.LineItems.Equal(state.LineItems) && !plan.LineItems.IsNull() && !plan.LineItems.IsUnknown() {
		if !assignAttrValueToNamedField(params, "LineItems", plan.LineItems) {
			if !plan.LineItems.Equal(state.LineItems) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "line_items", params)
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
	if !plan.OnBehalfOf.Equal(state.OnBehalfOf) && !plan.OnBehalfOf.IsNull() && !plan.OnBehalfOf.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "OnBehalfOfID", "OnBehalfOf", plan.OnBehalfOf.ValueString()) {
			if !plan.OnBehalfOf.Equal(state.OnBehalfOf) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "on_behalf_of", params)
			}
		}
	}
	if !plan.SubscriptionData.Equal(state.SubscriptionData) && !plan.SubscriptionData.IsNull() && !plan.SubscriptionData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SubscriptionData", plan.SubscriptionData) {
			if !plan.SubscriptionData.Equal(state.SubscriptionData) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "subscription_data", params)
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

	return params, nil
}

func flattenQuote(obj *stripe.Quote, state *QuoteResourceModel) error {
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
		if rawValueApplicationFeePercent, rawOk := plainValueAtPath(raw, "application_fee_percent"); rawOk {
			if valueApplicationFeePercent, err := flattenPlainValue(rawValueApplicationFeePercent, types.Float64Type, "application_fee_percent", "raw response"); err != nil {
				return err
			} else {
				if typedApplicationFeePercent, ok := valueApplicationFeePercent.(types.Float64); ok {
					state.ApplicationFeePercent = typedApplicationFeePercent
				}
			}
		} else if !hasRaw {
			if responseValueApplicationFeePercent, ok := plainFromResponseField(obj, "ApplicationFeePercent"); ok {
				if valueApplicationFeePercent, err := flattenPlainValue(responseValueApplicationFeePercent, types.Float64Type, "application_fee_percent", "response struct"); err != nil {
					return err
				} else {
					if typedApplicationFeePercent, ok := valueApplicationFeePercent.(types.Float64); ok {
						state.ApplicationFeePercent = typedApplicationFeePercent
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
		if rawValueCollectionMethod, rawOk := plainValueAtPath(raw, "collection_method"); rawOk {
			if valueCollectionMethod, err := flattenPlainValue(rawValueCollectionMethod, types.StringType, "collection_method", "raw response"); err != nil {
				return err
			} else {
				if typedCollectionMethod, ok := valueCollectionMethod.(types.String); ok {
					state.CollectionMethod = typedCollectionMethod
				}
			}
		} else if !hasRaw {
			if responseValueCollectionMethod, ok := plainFromResponseField(obj, "CollectionMethod"); ok {
				if valueCollectionMethod, err := flattenPlainValue(responseValueCollectionMethod, types.StringType, "collection_method", "response struct"); err != nil {
					return err
				} else {
					if typedCollectionMethod, ok := valueCollectionMethod.(types.String); ok {
						state.CollectionMethod = typedCollectionMethod
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
		if rawValueDefaultTaxRates, rawOk := plainValueAtPath(raw, "default_tax_rates"); rawOk {
			if valueDefaultTaxRates, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueDefaultTaxRates, attrValueToPlain(state.DefaultTaxRates)), types.ListType{ElemType: types.StringType}, "default_tax_rates", "raw response"); err != nil {
				return err
			} else {
				if typedDefaultTaxRates, ok := valueDefaultTaxRates.(types.List); ok {
					state.DefaultTaxRates = typedDefaultTaxRates
				}
			}
		} else if !hasRaw {
			if responseValueDefaultTaxRates, ok := plainFromResponseField(obj, "DefaultTaxRates"); ok {
				if valueDefaultTaxRates, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueDefaultTaxRates, attrValueToPlain(state.DefaultTaxRates)),
					types.ListType{ElemType: types.StringType},
					"default_tax_rates",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedDefaultTaxRates, ok := valueDefaultTaxRates.(types.List); ok {
						state.DefaultTaxRates = typedDefaultTaxRates
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
		if rawValueDiscounts, rawOk := plainValueAtPath(raw, "discounts"); rawOk {
			if valueDiscounts, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueDiscounts, attrValueToPlain(state.Discounts)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "discounts", "raw response"); err != nil {
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
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}},
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
		if rawValueFooter, rawOk := plainValueAtPath(raw, "footer"); rawOk {
			if valueFooter, err := flattenPlainValue(rawValueFooter, types.StringType, "footer", "raw response"); err != nil {
				return err
			} else {
				if typedFooter, ok := valueFooter.(types.String); ok {
					state.Footer = typedFooter
				}
			}
		} else if !hasRaw {
			if responseValueFooter, ok := plainFromResponseField(obj, "Footer"); ok {
				if valueFooter, err := flattenPlainValue(responseValueFooter, types.StringType, "footer", "response struct"); err != nil {
					return err
				} else {
					if typedFooter, ok := valueFooter.(types.String); ok {
						state.Footer = typedFooter
					}
				}
			}
		}
	}
	{
		assignedFromQuote := false
		hadRawFromQuote := false
		if rawValueFromQuote, rawOk := plainValueAtPath(raw, "from_quote"); rawOk {
			hadRawFromQuote = true
			if rawValueFromQuote != nil {
				sourceFromQuote := applyConfiguredKeyedListShapes(rawValueFromQuote, attrValueToPlain(state.FromQuote))
				if valueFromQuote, err := flattenPlainValue(sourceFromQuote, types.ObjectType{AttrTypes: map[string]attr.Type{"is_revision": types.BoolType, "quote": types.StringType}}, "from_quote", "raw response"); err != nil {
					return err
				} else {
					if typedFromQuote, ok := valueFromQuote.(types.Object); ok {
						state.FromQuote = typedFromQuote
						assignedFromQuote = true
					}
				}
			}
		}
		if !assignedFromQuote {
			if !hasRaw {
				if responseValueFromQuote, ok := plainFromResponseField(obj, "FromQuote"); ok {
					sourceFromQuote := applyConfiguredKeyedListShapes(responseValueFromQuote, attrValueToPlain(state.FromQuote))
					if valueFromQuote, err := flattenPlainValue(
						sourceFromQuote,
						types.ObjectType{AttrTypes: map[string]attr.Type{"is_revision": types.BoolType, "quote": types.StringType}},
						"from_quote",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFromQuote, ok := valueFromQuote.(types.Object); ok {
							state.FromQuote = typedFromQuote
							assignedFromQuote = true
						}
					}
				}
			}
		}
		if !assignedFromQuote && hadRawFromQuote {
			if nullFromQuote, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"is_revision": types.BoolType, "quote": types.StringType}}); ok {
				if typedFromQuote, ok := nullFromQuote.(types.Object); ok {
					state.FromQuote = typedFromQuote
				}
			}
		}
	}
	{
		if rawValueHeader, rawOk := plainValueAtPath(raw, "header"); rawOk {
			if valueHeader, err := flattenPlainValue(rawValueHeader, types.StringType, "header", "raw response"); err != nil {
				return err
			} else {
				if typedHeader, ok := valueHeader.(types.String); ok {
					state.Header = typedHeader
				}
			}
		} else if !hasRaw {
			if responseValueHeader, ok := plainFromResponseField(obj, "Header"); ok {
				if valueHeader, err := flattenPlainValue(responseValueHeader, types.StringType, "header", "response struct"); err != nil {
					return err
				} else {
					if typedHeader, ok := valueHeader.(types.String); ok {
						state.Header = typedHeader
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
		assignedInvoiceSettings := false
		hadRawInvoiceSettings := false
		if rawValueInvoiceSettings, rawOk := plainValueAtPath(raw, "invoice_settings"); rawOk {
			hadRawInvoiceSettings = true
			if rawValueInvoiceSettings != nil {
				sourceInvoiceSettings := applyConfiguredKeyedListShapes(rawValueInvoiceSettings, attrValueToPlain(state.InvoiceSettings))
				if valueInvoiceSettings, err := flattenPlainValue(sourceInvoiceSettings, types.ObjectType{AttrTypes: map[string]attr.Type{"days_until_due": types.Int64Type, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "invoice_settings", "raw response"); err != nil {
					return err
				} else {
					if typedInvoiceSettings, ok := valueInvoiceSettings.(types.Object); ok {
						state.InvoiceSettings = typedInvoiceSettings
						assignedInvoiceSettings = true
					}
				}
			}
		}
		if !assignedInvoiceSettings {
			if !hasRaw {
				if responseValueInvoiceSettings, ok := plainFromResponseField(obj, "InvoiceSettings"); ok {
					sourceInvoiceSettings := applyConfiguredKeyedListShapes(responseValueInvoiceSettings, attrValueToPlain(state.InvoiceSettings))
					if valueInvoiceSettings, err := flattenPlainValue(
						sourceInvoiceSettings,
						types.ObjectType{AttrTypes: map[string]attr.Type{"days_until_due": types.Int64Type, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}},
						"invoice_settings",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedInvoiceSettings, ok := valueInvoiceSettings.(types.Object); ok {
							state.InvoiceSettings = typedInvoiceSettings
							assignedInvoiceSettings = true
						}
					}
				}
			}
		}
		if !assignedInvoiceSettings && hadRawInvoiceSettings {
			if nullInvoiceSettings, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"days_until_due": types.Int64Type, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}); ok {
				if typedInvoiceSettings, ok := nullInvoiceSettings.(types.Object); ok {
					state.InvoiceSettings = typedInvoiceSettings
				}
			}
		}
	}
	{
		if rawValueLineItems, rawOk := plainValueAtPath(raw, "line_items"); rawOk {
			rawPlainLineItems := extractListObjectData(rawValueLineItems)
			if valueLineItems, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawPlainLineItems, attrValueToPlain(state.LineItems)), attrValueToPlain(state.LineItems)), attrValueToPlain(state.LineItems)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "price": types.StringType, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}, "quantity": types.Int64Type, "tax_rates": types.ListType{ElemType: types.StringType}}}}, "line_items", "raw response"); err != nil {
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
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "price": types.StringType, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}, "quantity": types.Int64Type, "tax_rates": types.ListType{ElemType: types.StringType}}}},
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
				if valueStatusTransitions, err := flattenPlainValue(sourceStatusTransitions, types.ObjectType{AttrTypes: map[string]attr.Type{"accepted_at": types.Int64Type, "canceled_at": types.Int64Type, "finalized_at": types.Int64Type}}, "status_transitions", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"accepted_at": types.Int64Type, "canceled_at": types.Int64Type, "finalized_at": types.Int64Type}},
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
			if nullStatusTransitions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"accepted_at": types.Int64Type, "canceled_at": types.Int64Type, "finalized_at": types.Int64Type}}); ok {
				if typedStatusTransitions, ok := nullStatusTransitions.(types.Object); ok {
					state.StatusTransitions = typedStatusTransitions
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
		assignedSubscriptionData := false
		hadRawSubscriptionData := false
		if rawValueSubscriptionData, rawOk := plainValueAtPath(raw, "subscription_data"); rawOk {
			hadRawSubscriptionData = true
			if rawValueSubscriptionData != nil {
				sourceSubscriptionData := applyConfiguredKeyedListShapes(rawValueSubscriptionData, attrValueToPlain(state.SubscriptionData))
				if valueSubscriptionData, err := flattenPlainValue(sourceSubscriptionData, types.ObjectType{AttrTypes: map[string]attr.Type{"billing_mode": types.ObjectType{AttrTypes: map[string]attr.Type{"flexible": types.ObjectType{AttrTypes: map[string]attr.Type{"proration_discounts": types.StringType}}, "type": types.StringType}}, "description": types.StringType, "effective_date": types.Int64Type, "metadata": types.MapType{ElemType: types.StringType}, "trial_period_days": types.Int64Type}}, "subscription_data", "raw response"); err != nil {
					return err
				} else {
					if typedSubscriptionData, ok := valueSubscriptionData.(types.Object); ok {
						state.SubscriptionData = typedSubscriptionData
						assignedSubscriptionData = true
					}
				}
			}
		}
		if !assignedSubscriptionData {
			if !hasRaw {
				if responseValueSubscriptionData, ok := plainFromResponseField(obj, "SubscriptionData"); ok {
					sourceSubscriptionData := applyConfiguredKeyedListShapes(responseValueSubscriptionData, attrValueToPlain(state.SubscriptionData))
					if valueSubscriptionData, err := flattenPlainValue(
						sourceSubscriptionData,
						types.ObjectType{AttrTypes: map[string]attr.Type{"billing_mode": types.ObjectType{AttrTypes: map[string]attr.Type{"flexible": types.ObjectType{AttrTypes: map[string]attr.Type{"proration_discounts": types.StringType}}, "type": types.StringType}}, "description": types.StringType, "effective_date": types.Int64Type, "metadata": types.MapType{ElemType: types.StringType}, "trial_period_days": types.Int64Type}},
						"subscription_data",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSubscriptionData, ok := valueSubscriptionData.(types.Object); ok {
							state.SubscriptionData = typedSubscriptionData
							assignedSubscriptionData = true
						}
					}
				}
			}
		}
		if !assignedSubscriptionData && hadRawSubscriptionData {
			if nullSubscriptionData, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"billing_mode": types.ObjectType{AttrTypes: map[string]attr.Type{"flexible": types.ObjectType{AttrTypes: map[string]attr.Type{"proration_discounts": types.StringType}}, "type": types.StringType}}, "description": types.StringType, "effective_date": types.Int64Type, "metadata": types.MapType{ElemType: types.StringType}, "trial_period_days": types.Int64Type}}); ok {
				if typedSubscriptionData, ok := nullSubscriptionData.(types.Object); ok {
					state.SubscriptionData = typedSubscriptionData
				}
			}
		}
	}
	{
		if true {
			if rawValueSubscriptionSchedule, rawOk := plainValueAtPath(raw, "subscription_schedule"); rawOk {
				if typedSubscriptionSchedule, ok := plainToStringIDValue(rawValueSubscriptionSchedule); ok {
					state.SubscriptionSchedule = typedSubscriptionSchedule
				}
			} else if !hasRaw {
				if responseValueSubscriptionSchedule, ok := plainFromResponseField(obj, "SubscriptionSchedule"); ok {
					if typedSubscriptionSchedule, ok := plainToStringIDValue(responseValueSubscriptionSchedule); ok {
						state.SubscriptionSchedule = typedSubscriptionSchedule
					}
				}
			}
		}
	}
	{
		if state.TestClock.IsNull() || state.TestClock.IsUnknown() {
			if rawValueTestClock, rawOk := plainValueAtPath(raw, "test_clock"); rawOk {
				if typedTestClock, ok := plainToStringIDValue(rawValueTestClock); ok {
					state.TestClock = typedTestClock
				}
			} else if !hasRaw {
				if responseValueTestClock, ok := plainFromResponseField(obj, "TestClock"); ok {
					if typedTestClock, ok := plainToStringIDValue(responseValueTestClock); ok {
						state.TestClock = typedTestClock
					}
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
		assignedTransferData := false
		hadRawTransferData := false
		if rawValueTransferData, rawOk := plainValueAtPath(raw, "transfer_data"); rawOk {
			hadRawTransferData = true
			if rawValueTransferData != nil {
				sourceTransferData := applyConfiguredKeyedListShapes(rawValueTransferData, attrValueToPlain(state.TransferData))
				if valueTransferData, err := flattenPlainValue(sourceTransferData, types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_percent": types.Float64Type, "destination": types.StringType}}, "transfer_data", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_percent": types.Float64Type, "destination": types.StringType}},
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
			if nullTransferData, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_percent": types.Float64Type, "destination": types.StringType}}); ok {
				if typedTransferData, ok := nullTransferData.(types.Object); ok {
					state.TransferData = typedTransferData
				}
			}
		}
	}
	return nil
}
