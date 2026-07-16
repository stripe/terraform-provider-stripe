//
// File generated from our OpenAPI spec
//

package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"reflect"
	"strconv"
	"strings"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &PriceResource{}

var _ resource.ResourceWithConfigure = &PriceResource{}

var _ resource.ResourceWithImportState = &PriceResource{}

func NewPriceResource() resource.Resource {
	return &PriceResource{}
}

type PriceResource struct {
	client *stripe.Client
}

type PriceResourceModel struct {
	Object            types.String `tfsdk:"object"`
	Active            types.Bool   `tfsdk:"active"`
	BillingScheme     types.String `tfsdk:"billing_scheme"`
	Created           types.Int64  `tfsdk:"created"`
	Currency          types.String `tfsdk:"currency"`
	CurrencyOptions   types.List   `tfsdk:"currency_options"`
	CustomUnitAmount  types.List   `tfsdk:"custom_unit_amount"`
	ID                types.String `tfsdk:"id"`
	Livemode          types.Bool   `tfsdk:"livemode"`
	LookupKey         types.String `tfsdk:"lookup_key"`
	Metadata          types.Map    `tfsdk:"metadata"`
	Nickname          types.String `tfsdk:"nickname"`
	Product           types.String `tfsdk:"product"`
	Recurring         types.List   `tfsdk:"recurring"`
	TaxBehavior       types.String `tfsdk:"tax_behavior"`
	Tiers             types.List   `tfsdk:"tiers"`
	TiersMode         types.String `tfsdk:"tiers_mode"`
	TransformQuantity types.Object `tfsdk:"transform_quantity"`
	Type              types.String `tfsdk:"type"`
	UnitAmount        types.Int64  `tfsdk:"unit_amount"`
	UnitAmountDecimal types.String `tfsdk:"unit_amount_decimal"`
	ProductData       types.List   `tfsdk:"product_data"`
	TransferLookupKey types.Bool   `tfsdk:"transfer_lookup_key"`
}

func (r *PriceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PriceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_price"
}

var _ resource.ResourceWithUpgradeState = &PriceResource{}

func (r *PriceResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: priceResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior PriceResourceModel
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradePriceStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
		1: {
			PriorSchema: priceResourceV1Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior PriceResourceV1Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradePriceStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func priceResourceV1Schema() *schema.Schema {
	return &schema.Schema{
		Description: "Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.\n[Products](https://api.stripe.com#products) help you track inventory or provisioning, and prices help you track payment terms. Different physical goods or levels of service should be represented by products, and pricing options should be represented by prices. This approach lets you change prices without having to change your provisioning scheme.\n\nFor example, you might have a single \"gold\" product that has prices for $10/month, $100/year, and €9 once.\n\nRelated guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription), [create an invoice](https://docs.stripe.com/billing/invoices/create), and more about [products and prices](https://docs.stripe.com/products-prices/overview).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("price")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the price can be used for new purchases.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"billing_scheme": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `unit_amount` or `unit_amount_decimal`) will be charged per unit in `quantity` (for prices with `usage_type=licensed`), or per unit of total usage (for prices with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("per_unit", "tiered")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Required:      true,
				Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"currency_options": schema.ListNestedAttribute{
				Optional:    true,
				Description: "Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required:    true,
							Description: "Key for this entry.",
						},
						"custom_unit_amount": schema.SingleNestedAttribute{
							Optional:    true,
							Description: "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
							Attributes: map[string]schema.Attribute{
								"maximum": schema.Int64Attribute{
									Optional:    true,
									Description: "The maximum unit amount the customer can specify for this item.",
								},
								"minimum": schema.Int64Attribute{
									Optional:    true,
									Description: "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
								},
								"preset": schema.Int64Attribute{
									Optional:    true,
									Description: "The starting unit amount which can be updated by the customer.",
								},
								"enabled": schema.BoolAttribute{
									Required:    true,
									Description: "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
								},
							},
						},
						"tax_behavior": schema.StringAttribute{
							Optional:    true,
							Description: "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
							Validators:  []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
						},
						"tiers": schema.ListNestedAttribute{
							Optional:    true,
							Description: "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"flat_amount": schema.Int64Attribute{
										Optional:    true,
										Description: "Price for the entire tier.",
									},
									"flat_amount_decimal": schema.StringAttribute{
										Optional:      true,
										Description:   "Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.",
										PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier()},
									},
									"unit_amount": schema.Int64Attribute{
										Optional:    true,
										Description: "Per unit price for units relevant to the tier.",
									},
									"unit_amount_decimal": schema.StringAttribute{
										Optional:      true,
										Description:   "Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.",
										PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier()},
									},
									"up_to": schema.StringAttribute{
										Required:    true,
										Description: "Up to and including to this quantity will be contained in the tier.",
									},
								},
							},
						},
						"unit_amount": schema.Int64Attribute{
							Optional:    true,
							Description: "The unit amount in cents (or local equivalent) to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.",
						},
						"unit_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier()},
						},
					},
				},
			},
			"custom_unit_amount": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"maximum": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The maximum unit amount the customer can specify for this item.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"minimum": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"preset": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The starting unit amount which can be updated by the customer.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"enabled": schema.BoolAttribute{
						Required:      true,
						Description:   "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"lookup_key": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A lookup key used to retrieve prices dynamically from a static string. This may be up to 200 characters.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"nickname": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A brief description of the price, hidden from customers.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"product": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ID of the product this price is associated with.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"recurring": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "The recurring components of a price such as `interval` and `usage_type`.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"interval": schema.StringAttribute{
						Required:      true,
						Description:   "The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("day", "month", "week", "year")},
					},
					"interval_count": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"meter": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The meter tracking the usage of a metered price",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"trial_period_days": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "Default number of trial days when subscribing a customer to this price using [`trial_from_plan=true`](https://docs.stripe.com/api#create_subscription-trial_from_plan).",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"usage_type": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("licensed", "metered")},
					},
				},
			},
			"tax_behavior": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
			},
			"tiers": schema.ListNestedAttribute{
				Optional:      true,
				Description:   "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"flat_amount": schema.Int64Attribute{
							Optional:      true,
							Description:   "Price for the entire tier.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
						"flat_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
						},
						"unit_amount": schema.Int64Attribute{
							Optional:      true,
							Description:   "Per unit price for units relevant to the tier.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
						"unit_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
						},
						"up_to": schema.StringAttribute{
							Required:      true,
							Description:   "Up to and including to this quantity will be contained in the tier.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
			"tiers_mode": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price. In `graduated` tiering, pricing can change as the quantity grows.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("graduated", "volume")},
			},
			"transform_quantity": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Apply a transformation to the reported usage or set quantity before computing the amount billed. Cannot be combined with `tiers`.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"divide_by": schema.Int64Attribute{
						Required:      true,
						Description:   "Divide usage by this number.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
					"round": schema.StringAttribute{
						Required:      true,
						Description:   "After division, either round the result `up` or `down`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("down", "up")},
					},
				},
			},
			"type": schema.StringAttribute{
				Computed:      true,
				Description:   "One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("one_time", "recurring")},
			},
			"unit_amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"unit_amount_decimal": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
			},
			"product_data": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "These fields can be used to create a new product that this price will belong to.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"active": schema.BoolAttribute{
						Optional:      true,
						Description:   "Whether the product is currently available for purchase. Defaults to `true`.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
					},
					"id": schema.StringAttribute{
						Optional:      true,
						Description:   "The identifier for the product. Must be unique. If not provided, an identifier will be randomly generated.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"metadata": schema.MapAttribute{
						Optional:      true,
						Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
						ElementType:   types.StringType,
					},
					"name": schema.StringAttribute{
						Required:      true,
						Description:   "The product's name, meant to be displayable to the customer.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"statement_descriptor": schema.StringAttribute{
						Optional:      true,
						Description:   "An arbitrary string to be displayed on your customer's credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all.\n\nThis may be up to 22 characters. The statement description may not include `<`, `>`, `\\`, `\"`, `'` characters, and will appear on your customer's statement in capital letters. Non-ASCII characters are automatically stripped.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"tax_code": schema.StringAttribute{
						Optional:      true,
						Description:   "A [tax code](https://docs.stripe.com/tax/tax-categories) ID.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"unit_label": schema.StringAttribute{
						Optional:      true,
						Description:   "A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"transfer_lookup_key": schema.BoolAttribute{
				Optional:    true,
				Description: "If set to true, will atomically remove the lookup key from the existing price, and assign it to this price.",
				WriteOnly:   true,
			},
		},
	}
}

func priceResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.\n[Products](https://api.stripe.com#products) help you track inventory or provisioning, and prices help you track payment terms. Different physical goods or levels of service should be represented by products, and pricing options should be represented by prices. This approach lets you change prices without having to change your provisioning scheme.\n\nFor example, you might have a single \"gold\" product that has prices for $10/month, $100/year, and €9 once.\n\nRelated guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription), [create an invoice](https://docs.stripe.com/billing/invoices/create), and more about [products and prices](https://docs.stripe.com/products-prices/overview).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("price")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the price can be used for new purchases.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"billing_scheme": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `unit_amount` or `unit_amount_decimal`) will be charged per unit in `quantity` (for prices with `usage_type=licensed`), or per unit of total usage (for prices with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("per_unit", "tiered")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Required:      true,
				Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"lookup_key": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A lookup key used to retrieve prices dynamically from a static string. This may be up to 200 characters.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"nickname": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A brief description of the price, hidden from customers.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"product": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ID of the product this price is associated with.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"tax_behavior": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
			},
			"tiers_mode": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price. In `graduated` tiering, pricing can change as the quantity grows.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("graduated", "volume")},
			},
			"transform_quantity": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Apply a transformation to the reported usage or set quantity before computing the amount billed. Cannot be combined with `tiers`.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"divide_by": schema.Int64Attribute{
						Required:      true,
						Description:   "Divide usage by this number.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
					"round": schema.StringAttribute{
						Required:      true,
						Description:   "After division, either round the result `up` or `down`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("down", "up")},
					},
				},
			},
			"type": schema.StringAttribute{
				Computed:      true,
				Description:   "One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("one_time", "recurring")},
			},
			"unit_amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"unit_amount_decimal": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
			},
			"transfer_lookup_key": schema.BoolAttribute{
				Optional:    true,
				Description: "If set to true, will atomically remove the lookup key from the existing price, and assign it to this price.",
				WriteOnly:   true,
			},
		},
		Blocks: map[string]schema.Block{
			"currency_options": schema.ListNestedBlock{
				Description: "Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required:    true,
							Description: "Key for this entry.",
						},
						"tax_behavior": schema.StringAttribute{
							Optional:    true,
							Description: "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
							Validators:  []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
						},
						"tiers": schema.ListNestedAttribute{
							Optional:    true,
							Description: "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"flat_amount": schema.Int64Attribute{
										Optional:    true,
										Description: "Price for the entire tier.",
									},
									"flat_amount_decimal": schema.StringAttribute{
										Optional:      true,
										Description:   "Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.",
										PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier()},
									},
									"unit_amount": schema.Int64Attribute{
										Optional:    true,
										Description: "Per unit price for units relevant to the tier.",
									},
									"unit_amount_decimal": schema.StringAttribute{
										Optional:      true,
										Description:   "Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.",
										PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier()},
									},
									"up_to": schema.StringAttribute{
										Required:    true,
										Description: "Up to and including to this quantity will be contained in the tier.",
									},
								},
							},
						},
						"unit_amount": schema.Int64Attribute{
							Optional:    true,
							Description: "The unit amount in cents (or local equivalent) to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.",
						},
						"unit_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier()},
						},
					},
					Blocks: map[string]schema.Block{
						"custom_unit_amount": schema.ListNestedBlock{
							Description: "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"maximum": schema.Int64Attribute{
										Optional:    true,
										Description: "The maximum unit amount the customer can specify for this item.",
									},
									"minimum": schema.Int64Attribute{
										Optional:    true,
										Description: "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
									},
									"preset": schema.Int64Attribute{
										Optional:    true,
										Description: "The starting unit amount which can be updated by the customer.",
									},
									"enabled": schema.BoolAttribute{
										Required:    true,
										Description: "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
									},
								},
							},
						},
					},
				},
			},
			"custom_unit_amount": schema.ListNestedBlock{
				Description:   "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"maximum": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The maximum unit amount the customer can specify for this item.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"minimum": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"preset": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The starting unit amount which can be updated by the customer.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"enabled": schema.BoolAttribute{
							Required:      true,
							Description:   "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
							PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
						},
					},
				},
			},
			"recurring": schema.ListNestedBlock{
				Description:   "The recurring components of a price such as `interval` and `usage_type`.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"interval": schema.StringAttribute{
							Required:      true,
							Description:   "The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							Validators:    []validator.String{stringvalidator.OneOf("day", "month", "week", "year")},
						},
						"interval_count": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"meter": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The meter tracking the usage of a metered price",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						},
						"trial_period_days": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "Default number of trial days when subscribing a customer to this price using [`trial_from_plan=true`](https://docs.stripe.com/api#create_subscription-trial_from_plan).",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"usage_type": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
							Validators:    []validator.String{stringvalidator.OneOf("licensed", "metered")},
						},
					},
				},
			},
			"tiers": schema.ListNestedBlock{
				Description:   "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"flat_amount": schema.Int64Attribute{
							Optional:      true,
							Description:   "Price for the entire tier.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
						"flat_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
						},
						"unit_amount": schema.Int64Attribute{
							Optional:      true,
							Description:   "Per unit price for units relevant to the tier.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
						"unit_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
						},
						"up_to": schema.StringAttribute{
							Required:      true,
							Description:   "Up to and including to this quantity will be contained in the tier.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
			"product_data": schema.ListNestedBlock{
				Description:   "These fields can be used to create a new product that this price will belong to.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"active": schema.BoolAttribute{
							Optional:      true,
							Description:   "Whether the product is currently available for purchase. Defaults to `true`.",
							PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
						},
						"id": schema.StringAttribute{
							Optional:      true,
							Description:   "The identifier for the product. Must be unique. If not provided, an identifier will be randomly generated.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"metadata": schema.MapAttribute{
							Optional:      true,
							Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
							PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
							ElementType:   types.StringType,
						},
						"name": schema.StringAttribute{
							Required:      true,
							Description:   "The product's name, meant to be displayable to the customer.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"statement_descriptor": schema.StringAttribute{
							Optional:      true,
							Description:   "An arbitrary string to be displayed on your customer's credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all.\n\nThis may be up to 22 characters. The statement description may not include `<`, `>`, `\\`, `\"`, `'` characters, and will appear on your customer's statement in capital letters. Non-ASCII characters are automatically stripped.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"tax_code": schema.StringAttribute{
							Optional:      true,
							Description:   "A [tax code](https://docs.stripe.com/tax/tax-categories) ID.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"unit_label": schema.StringAttribute{
							Optional:      true,
							Description:   "A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
		},
	}
}

type PriceResourceV1Model struct {
	Object            types.String `tfsdk:"object"`
	Active            types.Bool   `tfsdk:"active"`
	BillingScheme     types.String `tfsdk:"billing_scheme"`
	Created           types.Int64  `tfsdk:"created"`
	Currency          types.String `tfsdk:"currency"`
	CurrencyOptions   types.List   `tfsdk:"currency_options"`
	CustomUnitAmount  types.Object `tfsdk:"custom_unit_amount"`
	ID                types.String `tfsdk:"id"`
	Livemode          types.Bool   `tfsdk:"livemode"`
	LookupKey         types.String `tfsdk:"lookup_key"`
	Metadata          types.Map    `tfsdk:"metadata"`
	Nickname          types.String `tfsdk:"nickname"`
	Product           types.String `tfsdk:"product"`
	Recurring         types.Object `tfsdk:"recurring"`
	TaxBehavior       types.String `tfsdk:"tax_behavior"`
	Tiers             types.List   `tfsdk:"tiers"`
	TiersMode         types.String `tfsdk:"tiers_mode"`
	TransformQuantity types.Object `tfsdk:"transform_quantity"`
	Type              types.String `tfsdk:"type"`
	UnitAmount        types.Int64  `tfsdk:"unit_amount"`
	UnitAmountDecimal types.String `tfsdk:"unit_amount_decimal"`
	ProductData       types.Object `tfsdk:"product_data"`
	TransferLookupKey types.Bool   `tfsdk:"transfer_lookup_key"`
}

type priceStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]priceStateUpgradeAttrMeta
}

var priceStateUpgradeRootMeta = map[string]priceStateUpgradeAttrMeta{"object": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "active": priceStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "billing_scheme": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "created": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "currency": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "currency_options": priceStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "custom_unit_amount": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type, "enabled": types.BoolType}}}, "tax_behavior": types.StringType, "tiers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.StringType}}}, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", PreserveConfiguredValue: true, Nested: map[string]priceStateUpgradeAttrMeta{"key": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "custom_unit_amount": priceStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type, "enabled": types.BoolType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]priceStateUpgradeAttrMeta{"maximum": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "minimum": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "preset": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "enabled": priceStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "required", LegacyBehavior: "required"}}}, "tax_behavior": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "tiers": priceStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]priceStateUpgradeAttrMeta{"flat_amount": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "flat_amount_decimal": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "unit_amount": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "unit_amount_decimal": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "up_to": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "unit_amount": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "unit_amount_decimal": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}}}, "custom_unit_amount": priceStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type, "enabled": types.BoolType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]priceStateUpgradeAttrMeta{"maximum": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "minimum": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "preset": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "enabled": priceStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "required", LegacyBehavior: "required"}}}, "id": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "livemode": priceStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "lookup_key": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "metadata": priceStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "nickname": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "product": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "recurring": priceStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "meter": types.StringType, "trial_period_days": types.Int64Type, "usage_type": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]priceStateUpgradeAttrMeta{"interval": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "interval_count": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "meter": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "trial_period_days": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "usage_type": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}}, "tax_behavior": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "tiers": priceStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", PreserveConfiguredValue: true, Nested: map[string]priceStateUpgradeAttrMeta{"flat_amount": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "flat_amount_decimal": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "unit_amount": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "unit_amount_decimal": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "up_to": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "tiers_mode": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "transform_quantity": priceStateUpgradeAttrMeta{AttrType: types.ObjectType{AttrTypes: map[string]attr.Type{"divide_by": types.Int64Type, "round": types.StringType}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]priceStateUpgradeAttrMeta{"divide_by": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "required", LegacyBehavior: "required"}, "round": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "type": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "unit_amount": priceStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "unit_amount_decimal": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "product_data": priceStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"active": types.BoolType, "id": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "name": types.StringType, "statement_descriptor": types.StringType, "tax_code": types.StringType, "unit_label": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", PreserveConfiguredValue: true, Nested: map[string]priceStateUpgradeAttrMeta{"active": priceStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional", LegacyBehavior: "optional"}, "id": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "metadata": priceStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional", LegacyBehavior: "optional"}, "name": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "statement_descriptor": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "tax_code": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "unit_label": priceStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}}}, "transfer_lookup_key": priceStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional", LegacyBehavior: "optional"}}

var priceStateUpgradeSingletonPaths = map[string]struct{}{}

var priceStateUpgradeLegacyObjectPaths = map[string]struct{}{"currency_options.custom_unit_amount": struct{}{}, "custom_unit_amount": struct{}{}, "product_data": struct{}{}, "recurring": struct{}{}}

func priceAttrMapFromModel(model interface{}) map[string]attr.Value {
	value := reflect.ValueOf(model)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if !value.IsValid() || value.Kind() != reflect.Struct {
		return map[string]attr.Value{}
	}

	result := make(map[string]attr.Value, value.NumField())
	typeInfo := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := typeInfo.Field(i)
		tag := field.Tag.Get("tfsdk")
		if tag == "" || tag == "-" {
			continue
		}

		attrValue, ok := value.Field(i).Interface().(attr.Value)
		if !ok {
			continue
		}
		result[tag] = attrValue
	}
	return result
}

func priceSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
	value := reflect.ValueOf(target)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return
	}

	elem := value.Elem()
	if !elem.IsValid() || elem.Kind() != reflect.Struct {
		return
	}

	typeInfo := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		field := typeInfo.Field(i)
		tag := field.Tag.Get("tfsdk")
		if tag == "" || tag == "-" {
			continue
		}

		attrValue, ok := values[tag]
		if !ok {
			continue
		}

		fieldValue := elem.Field(i)
		renderedValue := reflect.ValueOf(attrValue)
		if renderedValue.Type().AssignableTo(fieldValue.Type()) {
			fieldValue.Set(renderedValue)
			continue
		}
		if renderedValue.Type().ConvertibleTo(fieldValue.Type()) {
			fieldValue.Set(renderedValue.Convert(fieldValue.Type()))
		}
	}
}

func priceIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func priceShouldPreserveChild(parent priceStateUpgradeAttrMeta, child priceStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !priceIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !priceIsComputedBehavior(child.Behavior)
}

func priceNullValueForType(attributeType attr.Type) attr.Value {
	switch t := attributeType.(type) {
	case basetypes.StringType:
		return types.StringNull()
	case basetypes.Int64Type:
		return types.Int64Null()
	case basetypes.Float64Type:
		return types.Float64Null()
	case basetypes.BoolType:
		return types.BoolNull()
	case basetypes.MapType:
		return types.MapNull(t.ElemType)
	case basetypes.ListType:
		return types.ListNull(t.ElemType)
	case basetypes.SetType:
		return types.SetNull(t.ElemType)
	case basetypes.ObjectType:
		return types.ObjectNull(t.AttrTypes)
	default:
		return types.StringNull()
	}
}

func priceLegacyUpgradeIsEmptyValue(value attr.Value) bool {
	if value == nil || value.IsNull() || value.IsUnknown() {
		return true
	}
	switch typed := value.(type) {
	case types.String:
		return typed.ValueString() == ""
	case types.Int64:
		return typed.ValueInt64() == 0
	case types.Float64:
		return typed.ValueFloat64() == 0
	case types.Bool:
		return !typed.ValueBool()
	case types.Map:
		return len(typed.Elements()) == 0
	case types.List:
		return len(typed.Elements()) == 0
	case types.Set:
		return len(typed.Elements()) == 0
	case types.Object:
		return len(typed.Attributes()) == 0
	default:
		return false
	}
}

func priceLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
	switch typed := value.(type) {
	case types.Int64:
		if typed.IsNull() || typed.IsUnknown() {
			return 0, false
		}
		return typed.ValueInt64(), true
	case types.Float64:
		if typed.IsNull() || typed.IsUnknown() {
			return 0, false
		}
		return int64(typed.ValueFloat64()), true
	default:
		return 0, false
	}
}

func priceLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := priceLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func priceLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.Int64)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || typedValue.ValueInt64() != 0 {
		return false
	}
	decimalValue, ok := siblings[name+"_decimal"]
	if !ok {
		return false
	}
	typedDecimal, ok := decimalValue.(types.String)
	if !ok || typedDecimal.IsNull() || typedDecimal.IsUnknown() {
		return false
	}
	return typedDecimal.ValueString() == "0"
}

func priceLegacyUpgradeNormalizeChild(parent priceStateUpgradeAttrMeta, name string, child priceStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if priceLegacyUpgradeIsEmptyValue(value) {
		return priceNullValueForType(child.AttrType)
	}
	if priceLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return priceNullValueForType(child.AttrType)
	}
	if priceLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return priceNullValueForType(child.AttrType)
	}
	return value
}

func priceLegacyUpgradeChildAttr(path []string, parent priceStateUpgradeAttrMeta, name string, child priceStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !priceShouldPreserveChild(parent, child) {
		return priceNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return priceNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := priceUpgradeValue(nextPath, child, childValue)
	return priceLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func priceUpgradeAttrs(path []string, meta map[string]priceStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = priceNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := priceUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = priceLegacyUpgradeNormalizeChild(
				priceStateUpgradeAttrMeta{PreserveConfiguredValue: true},
				name,
				fieldMeta,
				upgradedValue,
				prior,
			)
		}
		upgraded[name] = upgradedValue
	}
	return upgraded
}

func priceUpgradeObjectValue(path []string, meta priceStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = priceLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func priceUpgradeSingletonListToObject(path []string, meta priceStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return priceNullValueForType(meta.AttrType)
	}
	if listValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if listValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}

	elements := listValue.Elements()
	if len(elements) == 0 {
		return types.ObjectNull(objectType.AttrTypes)
	}

	firstObject, ok := elements[0].(types.Object)
	if !ok {
		if baseObject, baseOk := elements[0].(basetypes.ObjectValue); baseOk {
			firstObject = types.Object(baseObject)
		} else {
			return types.ObjectUnknown(objectType.AttrTypes)
		}
	}

	sourceAttrs := firstObject.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = priceLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func priceUpgradeObjectValueToSingletonList(path []string, meta priceStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	if listValue, ok := priorValue.(types.List); ok {
		return priceUpgradeListValue(path, meta, listType, listValue)
	}
	if baseList, ok := priorValue.(basetypes.ListValue); ok {
		return priceUpgradeListValue(path, meta, listType, types.List(baseList))
	}

	objectValue, ok := priorValue.(types.Object)
	if !ok {
		if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
			objectValue = types.Object(baseObject)
		} else {
			return priceNullValueForType(meta.AttrType)
		}
	}
	if objectValue.IsNull() {
		return types.ListNull(listType.ElemType)
	}
	if objectValue.IsUnknown() {
		return types.ListUnknown(listType.ElemType)
	}

	elementObjectType, ok := listType.ElemType.(basetypes.ObjectType)
	if !ok {
		return priceNullValueForType(meta.AttrType)
	}

	upgradedObject := priceUpgradeObjectValue(path, meta, elementObjectType, objectValue)
	return types.ListValueMust(listType.ElemType, []attr.Value{upgradedObject})
}

func priceUpgradeListValue(path []string, meta priceStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return priceNullValueForType(meta.AttrType)
	}
	if listValue.IsNull() {
		return types.ListNull(listType.ElemType)
	}
	if listValue.IsUnknown() {
		return types.ListUnknown(listType.ElemType)
	}
	if len(meta.Nested) == 0 {
		return listValue
	}

	elementObjectType, ok := listType.ElemType.(basetypes.ObjectType)
	if !ok {
		return listValue
	}

	elements := listValue.Elements()
	upgradedElements := make([]attr.Value, 0, len(elements))
	for _, element := range elements {
		objectValue, ok := element.(types.Object)
		if !ok {
			if baseObject, baseOk := element.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				upgradedElements = append(upgradedElements, priceNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			priceUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func priceUpgradeValue(path []string, meta priceStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := priceStateUpgradeSingletonPaths[pathKey]; ok {
			return priceUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return priceNullValueForType(meta.AttrType)
			}
		}
		return priceUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		if _, ok := priceStateUpgradeLegacyObjectPaths[pathKey]; ok {
			return priceUpgradeObjectValueToSingletonList(path, meta, attrType, priorValue)
		}
		return priceUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradePriceStateV1(ctx context.Context, prior interface{}) (PriceResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := priceUpgradeAttrs(nil, priceStateUpgradeRootMeta, priceAttrMapFromModel(prior))
	var upgraded PriceResourceModel
	priceSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *PriceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     2,
		Description: "Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.\n[Products](https://api.stripe.com#products) help you track inventory or provisioning, and prices help you track payment terms. Different physical goods or levels of service should be represented by products, and pricing options should be represented by prices. This approach lets you change prices without having to change your provisioning scheme.\n\nFor example, you might have a single \"gold\" product that has prices for $10/month, $100/year, and €9 once.\n\nRelated guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription), [create an invoice](https://docs.stripe.com/billing/invoices/create), and more about [products and prices](https://docs.stripe.com/products-prices/overview).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("price")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the price can be used for new purchases.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"billing_scheme": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `unit_amount` or `unit_amount_decimal`) will be charged per unit in `quantity` (for prices with `usage_type=licensed`), or per unit of total usage (for prices with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("per_unit", "tiered")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Required:      true,
				Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"lookup_key": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A lookup key used to retrieve prices dynamically from a static string. This may be up to 200 characters.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"nickname": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A brief description of the price, hidden from customers.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"product": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ID of the product this price is associated with.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"tax_behavior": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
			},
			"tiers_mode": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price. In `graduated` tiering, pricing can change as the quantity grows.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("graduated", "volume")},
			},
			"transform_quantity": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Apply a transformation to the reported usage or set quantity before computing the amount billed. Cannot be combined with `tiers`.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"divide_by": schema.Int64Attribute{
						Required:      true,
						Description:   "Divide usage by this number.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
					"round": schema.StringAttribute{
						Required:      true,
						Description:   "After division, either round the result `up` or `down`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("down", "up")},
					},
				},
			},
			"type": schema.StringAttribute{
				Computed:      true,
				Description:   "One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("one_time", "recurring")},
			},
			"unit_amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"unit_amount_decimal": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
			},
			"transfer_lookup_key": schema.BoolAttribute{
				Optional:    true,
				Description: "If set to true, will atomically remove the lookup key from the existing price, and assign it to this price.",
				WriteOnly:   true,
			},
		},
		Blocks: map[string]schema.Block{
			"currency_options": schema.ListNestedBlock{
				Description: "Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required:    true,
							Description: "Key for this entry.",
						},
						"tax_behavior": schema.StringAttribute{
							Optional:    true,
							Description: "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
							Validators:  []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
						},
						"tiers": schema.ListNestedAttribute{
							Optional:    true,
							Description: "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"flat_amount": schema.Int64Attribute{
										Optional:    true,
										Description: "Price for the entire tier.",
									},
									"flat_amount_decimal": schema.StringAttribute{
										Optional:      true,
										Description:   "Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.",
										PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier()},
									},
									"unit_amount": schema.Int64Attribute{
										Optional:    true,
										Description: "Per unit price for units relevant to the tier.",
									},
									"unit_amount_decimal": schema.StringAttribute{
										Optional:      true,
										Description:   "Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.",
										PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier()},
									},
									"up_to": schema.StringAttribute{
										Required:    true,
										Description: "Up to and including to this quantity will be contained in the tier.",
									},
								},
							},
						},
						"unit_amount": schema.Int64Attribute{
							Optional:    true,
							Description: "The unit amount in cents (or local equivalent) to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.",
						},
						"unit_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier()},
						},
					},
					Blocks: map[string]schema.Block{
						"custom_unit_amount": schema.ListNestedBlock{
							Description: "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"maximum": schema.Int64Attribute{
										Optional:    true,
										Description: "The maximum unit amount the customer can specify for this item.",
									},
									"minimum": schema.Int64Attribute{
										Optional:    true,
										Description: "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
									},
									"preset": schema.Int64Attribute{
										Optional:    true,
										Description: "The starting unit amount which can be updated by the customer.",
									},
									"enabled": schema.BoolAttribute{
										Required:    true,
										Description: "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
									},
								},
							},
						},
					},
				},
			},
			"custom_unit_amount": schema.ListNestedBlock{
				Description:   "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"maximum": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The maximum unit amount the customer can specify for this item.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"minimum": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"preset": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The starting unit amount which can be updated by the customer.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"enabled": schema.BoolAttribute{
							Required:      true,
							Description:   "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
							PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
						},
					},
				},
			},
			"recurring": schema.ListNestedBlock{
				Description:   "The recurring components of a price such as `interval` and `usage_type`.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"interval": schema.StringAttribute{
							Required:      true,
							Description:   "The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							Validators:    []validator.String{stringvalidator.OneOf("day", "month", "week", "year")},
						},
						"interval_count": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"meter": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The meter tracking the usage of a metered price",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						},
						"trial_period_days": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "Default number of trial days when subscribing a customer to this price using [`trial_from_plan=true`](https://docs.stripe.com/api#create_subscription-trial_from_plan).",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"usage_type": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
							Validators:    []validator.String{stringvalidator.OneOf("licensed", "metered")},
						},
					},
				},
			},
			"tiers": schema.ListNestedBlock{
				Description:   "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"flat_amount": schema.Int64Attribute{
							Optional:      true,
							Description:   "Price for the entire tier.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
						"flat_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
						},
						"unit_amount": schema.Int64Attribute{
							Optional:      true,
							Description:   "Per unit price for units relevant to the tier.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
						"unit_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
						},
						"up_to": schema.StringAttribute{
							Required:      true,
							Description:   "Up to and including to this quantity will be contained in the tier.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
			"product_data": schema.ListNestedBlock{
				Description:   "These fields can be used to create a new product that this price will belong to.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"active": schema.BoolAttribute{
							Optional:      true,
							Description:   "Whether the product is currently available for purchase. Defaults to `true`.",
							PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
						},
						"id": schema.StringAttribute{
							Optional:      true,
							Description:   "The identifier for the product. Must be unique. If not provided, an identifier will be randomly generated.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"metadata": schema.MapAttribute{
							Optional:      true,
							Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
							PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
							ElementType:   types.StringType,
						},
						"name": schema.StringAttribute{
							Required:      true,
							Description:   "The product's name, meant to be displayable to the customer.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"statement_descriptor": schema.StringAttribute{
							Optional:      true,
							Description:   "An arbitrary string to be displayed on your customer's credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all.\n\nThis may be up to 22 characters. The statement description may not include `<`, `>`, `\\`, `\"`, `'` characters, and will appear on your customer's statement in capital letters. Non-ASCII characters are automatically stripped.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"tax_code": schema.StringAttribute{
							Optional:      true,
							Description:   "A [tax code](https://docs.stripe.com/tax/tax-categories) ID.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"unit_label": schema.StringAttribute{
							Optional:      true,
							Description:   "A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
		},
	}
}

func (r *PriceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PriceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PriceResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"TransferLookupKey"}})

	params, err := expandPriceCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Price create params", err.Error())
		return
	}

	obj, err := r.client.V1Prices.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Price", err.Error())
		return
	}

	rawReadParams := &stripe.PriceRetrieveParams{}
	rawReadParams.AddExpand("tiers")

	if err := ensureRawResponse(obj, r.client.V1Prices.B, r.client.V1Prices.Key, stripe.FormatURLPath("/v1/prices/%s", obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating Price create raw response", err.Error())
		return
	}

	if err := flattenPrice(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Price create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"TransferLookupKey"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PriceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState PriceResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PriceResourceModel
	state = priorState

	params := &stripe.PriceRetrieveParams{}
	params.AddExpand("tiers")

	obj, err := r.client.V1Prices.Retrieve(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Price", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Prices.B, r.client.V1Prices.Key, stripe.FormatURLPath("/v1/prices/%s", state.ID.ValueString()), params); err != nil {
		resp.Diagnostics.AddError("Error hydrating Price raw response", err.Error())
		return
	}

	if err := flattenPrice(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Price read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"TransferLookupKey"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *PriceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PriceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PriceResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"TransferLookupKey"}})

	var state PriceResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandPriceUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Price update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Price update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Prices.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Price", err.Error())
		return
	}

	rawReadParams := &stripe.PriceRetrieveParams{}
	rawReadParams.AddExpand("tiers")

	if err := ensureRawResponse(obj, r.client.V1Prices.B, r.client.V1Prices.Key, stripe.FormatURLPath("/v1/prices/%s", obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating Price update raw response", err.Error())
		return
	}

	if err := flattenPrice(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Price update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"TransferLookupKey"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PriceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PriceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.Active.IsNull() && !state.Active.IsUnknown() && !state.Active.ValueBool() {
		return
	}

	params := &stripe.PriceUpdateParams{}
	activeField := reflect.ValueOf(params).Elem().FieldByName("Active")
	if activeField.IsValid() && activeField.CanSet() {
		if activeField.Kind() == reflect.Pointer && activeField.Type().Elem().Kind() == reflect.Bool {
			activeField.Set(reflect.ValueOf(stripe.Bool(false)))
		} else if activeField.Kind() == reflect.Bool {
			activeField.SetBool(false)
		}
	}

	_, err := r.client.V1Prices.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating Price", err.Error())
		return
	}
}

func (r *PriceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandPriceCreate(plan PriceResourceModel) (*stripe.PriceCreateParams, error) {
	params := &stripe.PriceCreateParams{}

	if !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.BillingScheme.IsNull() && !plan.BillingScheme.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "BillingScheme", "BillingScheme", plan.BillingScheme.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_scheme", params)
		}
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.CurrencyOptions.IsNull() && !plan.CurrencyOptions.IsUnknown() {
		if !assignKeyedListValueToNamedField(params, "CurrencyOptions", plan.CurrencyOptions, "key") {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency_options", params)
		}
	}
	if !plan.CustomUnitAmount.IsNull() && !plan.CustomUnitAmount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomUnitAmount", plan.CustomUnitAmount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "custom_unit_amount", params)
		}
	}
	if !plan.LookupKey.IsNull() && !plan.LookupKey.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LookupKey", "LookupKey", plan.LookupKey.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "lookup_key", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Nickname.IsNull() && !plan.Nickname.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Nickname", "Nickname", plan.Nickname.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "nickname", params)
		}
	}
	if !plan.Product.IsNull() && !plan.Product.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ProductID", "Product", plan.Product.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "product", params)
		}
	}
	if !plan.Recurring.IsNull() && !plan.Recurring.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Recurring", plan.Recurring) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "recurring", params)
		}
	}
	if !plan.TaxBehavior.IsNull() && !plan.TaxBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxBehavior", "TaxBehavior", plan.TaxBehavior.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_behavior", params)
		}
	}
	if !plan.Tiers.IsNull() && !plan.Tiers.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Tiers", plan.Tiers) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tiers", params)
		}
	}
	if !plan.TiersMode.IsNull() && !plan.TiersMode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TiersMode", "TiersMode", plan.TiersMode.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tiers_mode", params)
		}
	}
	if !plan.TransformQuantity.IsNull() && !plan.TransformQuantity.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TransformQuantity", plan.TransformQuantity) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "transform_quantity", params)
		}
	}
	if !plan.UnitAmount.IsNull() && !plan.UnitAmount.IsUnknown() {
		params.UnitAmount = stripe.Int64(plan.UnitAmount.ValueInt64())
	}
	if !plan.UnitAmountDecimal.IsNull() && !plan.UnitAmountDecimal.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "UnitAmountDecimal", "UnitAmountDecimal", plan.UnitAmountDecimal.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "unit_amount_decimal", params)
		}
	}
	if !plan.ProductData.IsNull() && !plan.ProductData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ProductData", plan.ProductData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "product_data", params)
		}
	}
	if !plan.TransferLookupKey.IsNull() && !plan.TransferLookupKey.IsUnknown() {
		params.TransferLookupKey = stripe.Bool(plan.TransferLookupKey.ValueBool())
	}

	return params, nil
}

func expandPriceUpdate(plan PriceResourceModel, state PriceResourceModel) (*stripe.PriceUpdateParams, error) {
	params := &stripe.PriceUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.CurrencyOptions.Equal(state.CurrencyOptions) && !plan.CurrencyOptions.IsNull() && !plan.CurrencyOptions.IsUnknown() {
		if !assignKeyedListValueToNamedField(params, "CurrencyOptions", plan.CurrencyOptions, "key") {
			if !plan.CurrencyOptions.Equal(state.CurrencyOptions) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "currency_options", params)
			}
		}
	}
	if !plan.LookupKey.Equal(state.LookupKey) && !plan.LookupKey.IsNull() && !plan.LookupKey.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LookupKey", "LookupKey", plan.LookupKey.ValueString()) {
			if !plan.LookupKey.Equal(state.LookupKey) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "lookup_key", params)
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
	if !plan.Nickname.Equal(state.Nickname) && !plan.Nickname.IsNull() && !plan.Nickname.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Nickname", "Nickname", plan.Nickname.ValueString()) {
			if !plan.Nickname.Equal(state.Nickname) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "nickname", params)
			}
		}
	}
	if !plan.TaxBehavior.Equal(state.TaxBehavior) && !plan.TaxBehavior.IsNull() && !plan.TaxBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxBehavior", "TaxBehavior", plan.TaxBehavior.ValueString()) {
			if !plan.TaxBehavior.Equal(state.TaxBehavior) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tax_behavior", params)
			}
		}
	}
	if !plan.TransferLookupKey.Equal(state.TransferLookupKey) && !plan.TransferLookupKey.IsNull() && !plan.TransferLookupKey.IsUnknown() {
		params.TransferLookupKey = stripe.Bool(plan.TransferLookupKey.ValueBool())
	}

	return params, nil
}

func flattenPrice(obj *stripe.Price, state *PriceResourceModel) error {
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
		if rawValueActive, rawOk := plainValueAtPath(raw, "active"); rawOk {
			if valueActive, err := flattenPlainValue(rawValueActive, types.BoolType, "active", "raw response"); err != nil {
				return err
			} else {
				if typedActive, ok := valueActive.(types.Bool); ok {
					state.Active = typedActive
				}
			}
		} else if !hasRaw {
			if responseValueActive, ok := plainFromResponseField(obj, "Active"); ok {
				if valueActive, err := flattenPlainValue(responseValueActive, types.BoolType, "active", "response struct"); err != nil {
					return err
				} else {
					if typedActive, ok := valueActive.(types.Bool); ok {
						state.Active = typedActive
					}
				}
			}
		}
	}
	{
		if rawValueBillingScheme, rawOk := plainValueAtPath(raw, "billing_scheme"); rawOk {
			if valueBillingScheme, err := flattenPlainValue(rawValueBillingScheme, types.StringType, "billing_scheme", "raw response"); err != nil {
				return err
			} else {
				if typedBillingScheme, ok := valueBillingScheme.(types.String); ok {
					state.BillingScheme = typedBillingScheme
				}
			}
		} else if !hasRaw {
			if responseValueBillingScheme, ok := plainFromResponseField(obj, "BillingScheme"); ok {
				if valueBillingScheme, err := flattenPlainValue(responseValueBillingScheme, types.StringType, "billing_scheme", "response struct"); err != nil {
					return err
				} else {
					if typedBillingScheme, ok := valueBillingScheme.(types.String); ok {
						state.BillingScheme = typedBillingScheme
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
		if rawValueCurrencyOptions, rawOk := plainValueAtPath(raw, "currency_options"); rawOk {
			if valueCurrencyOptions, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(applyPlainNestedObjectDefaultLeafValues(plainMapToKeyedList(rawValueCurrencyOptions, "key"), []plainNestedObjectDefaultLeafValues{{ObjectPath: []string{"*", "custom_unit_amount"}, Defaults: []plainObjectDefaultLeafValue{{Target: []string{"enabled"}, Value: true}}}}), attrValueToPlain(state.CurrencyOptions)), attrValueToPlain(state.CurrencyOptions)), attrValueToPlain(state.CurrencyOptions)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "custom_unit_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type, "enabled": types.BoolType}}, "tax_behavior": types.StringType, "tiers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.StringType}}}, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType}}}, "currency_options", "raw response"); err != nil {
				return err
			} else {
				if typedCurrencyOptions, ok := valueCurrencyOptions.(types.List); ok {
					state.CurrencyOptions = typedCurrencyOptions
				}
			}
		} else if !hasRaw {
			if responseValueCurrencyOptions, ok := plainFromResponseField(obj, "CurrencyOptions"); ok {
				if valueCurrencyOptions, err := flattenPlainValue(
					preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(applyPlainNestedObjectDefaultLeafValues(plainMapToKeyedList(responseValueCurrencyOptions, "key"), []plainNestedObjectDefaultLeafValues{{ObjectPath: []string{"*", "custom_unit_amount"}, Defaults: []plainObjectDefaultLeafValue{{Target: []string{"enabled"}, Value: true}}}}), attrValueToPlain(state.CurrencyOptions)), attrValueToPlain(state.CurrencyOptions)), attrValueToPlain(state.CurrencyOptions)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "custom_unit_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type, "enabled": types.BoolType}}, "tax_behavior": types.StringType, "tiers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.StringType}}}, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType}}},
					"currency_options",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedCurrencyOptions, ok := valueCurrencyOptions.(types.List); ok {
						state.CurrencyOptions = typedCurrencyOptions
					}
				}
			}
		}
	}
	{
		assignedCustomUnitAmount := false
		hadRawCustomUnitAmount := false
		if rawValueCustomUnitAmount, rawOk := plainValueAtPath(raw, "custom_unit_amount"); rawOk {
			hadRawCustomUnitAmount = true
			if rawValueCustomUnitAmount != nil {
				sourceCustomUnitAmount := applyConfiguredKeyedListShapes(applyPlainNestedObjectDefaultLeafValues(rawValueCustomUnitAmount, []plainNestedObjectDefaultLeafValues{{ObjectPath: []string{}, Defaults: []plainObjectDefaultLeafValue{{Target: []string{"enabled"}, Value: true}}}}), unwrapPlainSingletonList(attrValueToPlain(state.CustomUnitAmount)))
				if valueCustomUnitAmount, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceCustomUnitAmount, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type, "enabled": types.BoolType}}}, "custom_unit_amount", "raw response"); err != nil {
					return err
				} else {
					if typedCustomUnitAmount, ok := valueCustomUnitAmount.(types.List); ok {
						state.CustomUnitAmount = typedCustomUnitAmount
						assignedCustomUnitAmount = true
					}
				}
			}
		}
		if !assignedCustomUnitAmount {
			if !hasRaw {
				if responseValueCustomUnitAmount, ok := plainFromResponseField(obj, "CustomUnitAmount"); ok {
					sourceCustomUnitAmount := applyConfiguredKeyedListShapes(applyPlainNestedObjectDefaultLeafValues(responseValueCustomUnitAmount, []plainNestedObjectDefaultLeafValues{{ObjectPath: []string{}, Defaults: []plainObjectDefaultLeafValue{{Target: []string{"enabled"}, Value: true}}}}), unwrapPlainSingletonList(attrValueToPlain(state.CustomUnitAmount)))
					if valueCustomUnitAmount, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceCustomUnitAmount, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type, "enabled": types.BoolType}}},
						"custom_unit_amount",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCustomUnitAmount, ok := valueCustomUnitAmount.(types.List); ok {
							state.CustomUnitAmount = typedCustomUnitAmount
							assignedCustomUnitAmount = true
						}
					}
				}
			}
		}
		if !assignedCustomUnitAmount && hadRawCustomUnitAmount {
			if nullCustomUnitAmount, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type, "enabled": types.BoolType}}}); ok {
				if typedCustomUnitAmount, ok := nullCustomUnitAmount.(types.List); ok {
					state.CustomUnitAmount = typedCustomUnitAmount
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
		if rawValueLookupKey, rawOk := plainValueAtPath(raw, "lookup_key"); rawOk {
			if valueLookupKey, err := flattenPlainValue(rawValueLookupKey, types.StringType, "lookup_key", "raw response"); err != nil {
				return err
			} else {
				if typedLookupKey, ok := valueLookupKey.(types.String); ok {
					state.LookupKey = typedLookupKey
				}
			}
		} else if !hasRaw {
			if responseValueLookupKey, ok := plainFromResponseField(obj, "LookupKey"); ok {
				if valueLookupKey, err := flattenPlainValue(responseValueLookupKey, types.StringType, "lookup_key", "response struct"); err != nil {
					return err
				} else {
					if typedLookupKey, ok := valueLookupKey.(types.String); ok {
						state.LookupKey = typedLookupKey
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
		if rawValueNickname, rawOk := plainValueAtPath(raw, "nickname"); rawOk {
			if valueNickname, err := flattenPlainValue(rawValueNickname, types.StringType, "nickname", "raw response"); err != nil {
				return err
			} else {
				if typedNickname, ok := valueNickname.(types.String); ok {
					state.Nickname = typedNickname
				}
			}
		} else if !hasRaw {
			if responseValueNickname, ok := plainFromResponseField(obj, "Nickname"); ok {
				if valueNickname, err := flattenPlainValue(responseValueNickname, types.StringType, "nickname", "response struct"); err != nil {
					return err
				} else {
					if typedNickname, ok := valueNickname.(types.String); ok {
						state.Nickname = typedNickname
					}
				}
			}
		}
	}
	{
		if state.Product.IsNull() || state.Product.IsUnknown() {
			if rawValueProduct, rawOk := plainValueAtPath(raw, "product"); rawOk {
				if typedProduct, ok := plainToStringIDValue(rawValueProduct); ok {
					state.Product = typedProduct
				}
			} else if !hasRaw {
				if responseValueProduct, ok := plainFromResponseField(obj, "Product"); ok {
					if typedProduct, ok := plainToStringIDValue(responseValueProduct); ok {
						state.Product = typedProduct
					}
				}
			}
		}
	}
	{
		assignedRecurring := false
		hadRawRecurring := false
		if rawValueRecurring, rawOk := plainValueAtPath(raw, "recurring"); rawOk {
			hadRawRecurring = true
			if rawValueRecurring != nil {
				sourceRecurring := applyConfiguredKeyedListShapes(rawValueRecurring, unwrapPlainSingletonList(attrValueToPlain(state.Recurring)))
				if valueRecurring, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceRecurring, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "meter": types.StringType, "trial_period_days": types.Int64Type, "usage_type": types.StringType}}}, "recurring", "raw response"); err != nil {
					return err
				} else {
					if typedRecurring, ok := valueRecurring.(types.List); ok {
						state.Recurring = typedRecurring
						assignedRecurring = true
					}
				}
			}
		}
		if !assignedRecurring {
			if !hasRaw {
				if responseValueRecurring, ok := plainFromResponseField(obj, "Recurring"); ok {
					sourceRecurring := applyConfiguredKeyedListShapes(responseValueRecurring, unwrapPlainSingletonList(attrValueToPlain(state.Recurring)))
					if valueRecurring, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceRecurring, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "meter": types.StringType, "trial_period_days": types.Int64Type, "usage_type": types.StringType}}},
						"recurring",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRecurring, ok := valueRecurring.(types.List); ok {
							state.Recurring = typedRecurring
							assignedRecurring = true
						}
					}
				}
			}
		}
		if !assignedRecurring && hadRawRecurring {
			if nullRecurring, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type, "meter": types.StringType, "trial_period_days": types.Int64Type, "usage_type": types.StringType}}}); ok {
				if typedRecurring, ok := nullRecurring.(types.List); ok {
					state.Recurring = typedRecurring
				}
			}
		}
	}
	{
		if rawValueTaxBehavior, rawOk := plainValueAtPath(raw, "tax_behavior"); rawOk {
			if valueTaxBehavior, err := flattenPlainValue(rawValueTaxBehavior, types.StringType, "tax_behavior", "raw response"); err != nil {
				return err
			} else {
				if typedTaxBehavior, ok := valueTaxBehavior.(types.String); ok {
					state.TaxBehavior = typedTaxBehavior
				}
			}
		} else if !hasRaw {
			if responseValueTaxBehavior, ok := plainFromResponseField(obj, "TaxBehavior"); ok {
				if valueTaxBehavior, err := flattenPlainValue(responseValueTaxBehavior, types.StringType, "tax_behavior", "response struct"); err != nil {
					return err
				} else {
					if typedTaxBehavior, ok := valueTaxBehavior.(types.String); ok {
						state.TaxBehavior = typedTaxBehavior
					}
				}
			}
		}
	}
	{
		if rawValueTiers, rawOk := plainValueAtPath(raw, "tiers"); rawOk {
			if valueTiers, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawValueTiers, attrValueToPlain(state.Tiers)), attrValueToPlain(state.Tiers)), attrValueToPlain(state.Tiers)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.StringType}}}, "tiers", "raw response"); err != nil {
				return err
			} else {
				if typedTiers, ok := valueTiers.(types.List); ok {
					state.Tiers = typedTiers
				}
			}
		} else if !hasRaw {
			if responseValueTiers, ok := plainFromResponseField(obj, "Tiers"); ok {
				if valueTiers, err := flattenPlainValue(
					preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(responseValueTiers, attrValueToPlain(state.Tiers)), attrValueToPlain(state.Tiers)), attrValueToPlain(state.Tiers)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.StringType}}},
					"tiers",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedTiers, ok := valueTiers.(types.List); ok {
						state.Tiers = typedTiers
					}
				}
			}
		}
	}
	{
		if rawValueTiersMode, rawOk := plainValueAtPath(raw, "tiers_mode"); rawOk {
			if valueTiersMode, err := flattenPlainValue(rawValueTiersMode, types.StringType, "tiers_mode", "raw response"); err != nil {
				return err
			} else {
				if typedTiersMode, ok := valueTiersMode.(types.String); ok {
					state.TiersMode = typedTiersMode
				}
			}
		} else if !hasRaw {
			if responseValueTiersMode, ok := plainFromResponseField(obj, "TiersMode"); ok {
				if valueTiersMode, err := flattenPlainValue(responseValueTiersMode, types.StringType, "tiers_mode", "response struct"); err != nil {
					return err
				} else {
					if typedTiersMode, ok := valueTiersMode.(types.String); ok {
						state.TiersMode = typedTiersMode
					}
				}
			}
		}
	}
	{
		assignedTransformQuantity := false
		hadRawTransformQuantity := false
		if rawValueTransformQuantity, rawOk := plainValueAtPath(raw, "transform_quantity"); rawOk {
			hadRawTransformQuantity = true
			if rawValueTransformQuantity != nil {
				sourceTransformQuantity := applyConfiguredKeyedListShapes(rawValueTransformQuantity, attrValueToPlain(state.TransformQuantity))
				if valueTransformQuantity, err := flattenPlainValue(sourceTransformQuantity, types.ObjectType{AttrTypes: map[string]attr.Type{"divide_by": types.Int64Type, "round": types.StringType}}, "transform_quantity", "raw response"); err != nil {
					return err
				} else {
					if typedTransformQuantity, ok := valueTransformQuantity.(types.Object); ok {
						state.TransformQuantity = typedTransformQuantity
						assignedTransformQuantity = true
					}
				}
			}
		}
		if !assignedTransformQuantity {
			if !hasRaw {
				if responseValueTransformQuantity, ok := plainFromResponseField(obj, "TransformQuantity"); ok {
					sourceTransformQuantity := applyConfiguredKeyedListShapes(responseValueTransformQuantity, attrValueToPlain(state.TransformQuantity))
					if valueTransformQuantity, err := flattenPlainValue(
						sourceTransformQuantity,
						types.ObjectType{AttrTypes: map[string]attr.Type{"divide_by": types.Int64Type, "round": types.StringType}},
						"transform_quantity",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTransformQuantity, ok := valueTransformQuantity.(types.Object); ok {
							state.TransformQuantity = typedTransformQuantity
							assignedTransformQuantity = true
						}
					}
				}
			}
		}
		if !assignedTransformQuantity && hadRawTransformQuantity {
			if nullTransformQuantity, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"divide_by": types.Int64Type, "round": types.StringType}}); ok {
				if typedTransformQuantity, ok := nullTransformQuantity.(types.Object); ok {
					state.TransformQuantity = typedTransformQuantity
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
		if rawValueUnitAmount, rawOk := plainValueAtPath(raw, "unit_amount"); rawOk {
			if valueUnitAmount, err := flattenPlainValue(rawValueUnitAmount, types.Int64Type, "unit_amount", "raw response"); err != nil {
				return err
			} else {
				if typedUnitAmount, ok := valueUnitAmount.(types.Int64); ok {
					state.UnitAmount = typedUnitAmount
				}
			}
		} else if !hasRaw {
			if responseValueUnitAmount, ok := plainFromResponseField(obj, "UnitAmount"); ok {
				if valueUnitAmount, err := flattenPlainValue(responseValueUnitAmount, types.Int64Type, "unit_amount", "response struct"); err != nil {
					return err
				} else {
					if typedUnitAmount, ok := valueUnitAmount.(types.Int64); ok {
						state.UnitAmount = typedUnitAmount
					}
				}
			}
		}
	}
	{
		if rawValueUnitAmountDecimal, rawOk := plainValueAtPath(raw, "unit_amount_decimal"); rawOk {
			if valueUnitAmountDecimal, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(rawValueUnitAmountDecimal, attrValueToPlain(state.UnitAmountDecimal)), types.StringType, "unit_amount_decimal", "raw response"); err != nil {
				return err
			} else {
				if typedUnitAmountDecimal, ok := valueUnitAmountDecimal.(types.String); ok {
					state.UnitAmountDecimal = typedUnitAmountDecimal
				}
			}
		} else if !hasRaw {
			if responseValueUnitAmountDecimal, ok := plainFromResponseField(obj, "UnitAmountDecimal"); ok {
				if valueUnitAmountDecimal, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(responseValueUnitAmountDecimal, attrValueToPlain(state.UnitAmountDecimal)), types.StringType, "unit_amount_decimal", "response struct"); err != nil {
					return err
				} else {
					if typedUnitAmountDecimal, ok := valueUnitAmountDecimal.(types.String); ok {
						state.UnitAmountDecimal = typedUnitAmountDecimal
					}
				}
			}
		}
	}
	{
		assignedProductData := false
		if rawValueProductData, rawOk := plainValueAtPath(raw, "product_data"); rawOk {
			if rawValueProductData != nil {
				sourceProductData := mergeMissingPlainLeaves(suppressUnconfiguredOptionalReadbackLeaves(applyConfiguredKeyedListShapes(rawValueProductData, unwrapPlainSingletonList(attrValueToPlain(state.ProductData))), unwrapPlainSingletonList(attrValueToPlain(state.ProductData)), [][]string{[]string{"active"}, []string{"id"}, []string{"metadata"}, []string{"statement_descriptor"}, []string{"tax_code"}, []string{"unit_label"}}), unwrapPlainSingletonList(attrValueToPlain(state.ProductData)))
				if valueProductData, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceProductData, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"active": types.BoolType, "id": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "name": types.StringType, "statement_descriptor": types.StringType, "tax_code": types.StringType, "unit_label": types.StringType}}}, "product_data", "raw response"); err != nil {
					return err
				} else {
					if typedProductData, ok := valueProductData.(types.List); ok {
						state.ProductData = typedProductData
						assignedProductData = true
					}
				}
			}
		}
		if !assignedProductData {
			if !hasRaw {
				if responseValueProductData, ok := plainFromResponseField(obj, "ProductData"); ok {
					sourceProductData := mergeMissingPlainLeaves(suppressUnconfiguredOptionalReadbackLeaves(applyConfiguredKeyedListShapes(responseValueProductData, unwrapPlainSingletonList(attrValueToPlain(state.ProductData))), unwrapPlainSingletonList(attrValueToPlain(state.ProductData)), [][]string{[]string{"active"}, []string{"id"}, []string{"metadata"}, []string{"statement_descriptor"}, []string{"tax_code"}, []string{"unit_label"}}), unwrapPlainSingletonList(attrValueToPlain(state.ProductData)))
					if valueProductData, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceProductData, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"active": types.BoolType, "id": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "name": types.StringType, "statement_descriptor": types.StringType, "tax_code": types.StringType, "unit_label": types.StringType}}},
						"product_data",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedProductData, ok := valueProductData.(types.List); ok {
							state.ProductData = typedProductData
							assignedProductData = true
						}
					}
				}
			}
		}
	}
	return nil
}
