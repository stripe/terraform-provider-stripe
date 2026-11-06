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

var _ resource.Resource = &ProductResource{}

var _ resource.ResourceWithConfigure = &ProductResource{}

var _ resource.ResourceWithImportState = &ProductResource{}

func NewProductResource() resource.Resource {
	return &ProductResource{}
}

type ProductResource struct {
	client *stripe.Client
}

type ProductResourceModel struct {
	Object              types.String `tfsdk:"object"`
	Active              types.Bool   `tfsdk:"active"`
	Created             types.Int64  `tfsdk:"created"`
	DefaultPrice        types.String `tfsdk:"default_price"`
	Description         types.String `tfsdk:"description"`
	ID                  types.String `tfsdk:"id"`
	Images              types.List   `tfsdk:"images"`
	Livemode            types.Bool   `tfsdk:"livemode"`
	MarketingFeatures   types.List   `tfsdk:"marketing_features"`
	Metadata            types.Map    `tfsdk:"metadata"`
	Name                types.String `tfsdk:"name"`
	PackageDimensions   types.List   `tfsdk:"package_dimensions"`
	Shippable           types.Bool   `tfsdk:"shippable"`
	StatementDescriptor types.String `tfsdk:"statement_descriptor"`
	TaxCode             types.String `tfsdk:"tax_code"`
	Type                types.String `tfsdk:"type"`
	UnitLabel           types.String `tfsdk:"unit_label"`
	Updated             types.Int64  `tfsdk:"updated"`
	URL                 types.String `tfsdk:"url"`
	DefaultPriceData    types.List   `tfsdk:"default_price_data"`
}

func (r *ProductResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ProductResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_product"
}

var _ resource.ResourceWithUpgradeState = &ProductResource{}

func (r *ProductResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: productResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior ProductResourceModel
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeProductStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
		1: {
			PriorSchema: productResourceV1Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior ProductResourceV1Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeProductStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func productResourceV1Schema() *schema.Schema {
	return &schema.Schema{
		Description: "Products describe the specific goods or services you offer to your customers.\nFor example, you might offer a Standard and Premium version of your goods or service; each version would be a separate Product.\nThey can be used in conjunction with [Prices](https://api.stripe.com#prices) to configure pricing in Payment Links, Checkout, and Subscriptions.\n\nRelated guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription),\n[share a Payment Link](https://docs.stripe.com/payment-links),\n[accept payments with Checkout](https://docs.stripe.com/payments/accept-a-payment#create-product-prices-upfront),\nand more about [Products and Prices](https://docs.stripe.com/products-prices/overview)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("product")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the product is currently available for purchase.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"default_price": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the [Price](https://docs.stripe.com/api/prices) object that is the default price for this product.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"images": schema.ListAttribute{
				Optional:    true,
				Description: "A list of up to 8 URLs of images for this product, meant to be displayable to the customer.",
				ElementType: types.StringType,
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"marketing_features": schema.ListNestedAttribute{
				Optional:    true,
				Description: "A list of up to 15 marketing features for this product. These are displayed in [pricing tables](https://docs.stripe.com/payments/checkout/pricing-table).",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:    true,
							Description: "The marketing feature name. Up to 80 characters long.",
						},
					},
				},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The product's name, meant to be displayable to the customer.",
			},
			"package_dimensions": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "The dimensions of this product for shipping purposes.",
				Attributes: map[string]schema.Attribute{
					"height": schema.Float64Attribute{
						Required:    true,
						Description: "Height, in inches.",
					},
					"length": schema.Float64Attribute{
						Required:    true,
						Description: "Length, in inches.",
					},
					"weight": schema.Float64Attribute{
						Required:    true,
						Description: "Weight, in ounces.",
					},
					"width": schema.Float64Attribute{
						Required:    true,
						Description: "Width, in inches.",
					},
				},
			},
			"shippable": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether this product is shipped (i.e., physical goods).",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"statement_descriptor": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Extra information about a product which will appear on your customer's credit card statement. In the case that multiple products are billed at once, the first statement descriptor will be used. Only used for subscription payments.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"tax_code": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A [tax code](https://docs.stripe.com/tax/tax-categories) ID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The type of the product. The product is either of type `good`, which is eligible for use with Orders and SKUs, or `service`, which is eligible for use with Subscriptions and Plans.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("good", "service")},
			},
			"unit_label": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"updated": schema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was last updated. Measured in seconds since the Unix epoch.",
			},
			"url": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A URL of a publicly-accessible webpage for this product.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_price_data": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object. This Price will be set as the default price for this product.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"currency": schema.StringAttribute{
						Required:      true,
						Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"currency_options": schema.ListNestedAttribute{
						Optional:      true,
						Description:   "Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
						PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Required:    true,
									Description: "Key for this entry.",
								},
								"custom_unit_amount": schema.SingleNestedAttribute{
									Optional:      true,
									Description:   "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
									PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Required:      true,
											Description:   "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
											PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
										},
										"maximum": schema.Int64Attribute{
											Optional:      true,
											Description:   "The maximum unit amount the customer can specify for this item.",
											PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
										},
										"minimum": schema.Int64Attribute{
											Optional:      true,
											Description:   "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
											PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
										},
										"preset": schema.Int64Attribute{
											Optional:      true,
											Description:   "The starting unit amount which can be updated by the customer.",
											PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
										},
									},
								},
								"tax_behavior": schema.StringAttribute{
									Optional:      true,
									Description:   "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								},
								"tiers": schema.ListNestedAttribute{
									Optional:      true,
									Description:   "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
									PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"flat_amount": schema.Int64Attribute{
												Optional:      true,
												Description:   "The flat billing amount for an entire tier, regardless of the number of units in the tier.",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
											},
											"flat_amount_decimal": schema.StringAttribute{
												Optional:      true,
												Description:   "Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.",
												PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
											},
											"unit_amount": schema.Int64Attribute{
												Optional:      true,
												Description:   "The per unit billing amount for each individual unit for which this tier applies.",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
											},
											"unit_amount_decimal": schema.StringAttribute{
												Optional:      true,
												Description:   "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
												PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
											},
											"up_to": schema.Int64Attribute{
												Required:      true,
												Description:   "Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
											},
										},
									},
								},
								"unit_amount": schema.Int64Attribute{
									Optional:      true,
									Description:   "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
								},
								"unit_amount_decimal": schema.StringAttribute{
									Optional:      true,
									Description:   "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
									PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
								},
							},
						},
					},
					"custom_unit_amount": schema.SingleNestedAttribute{
						Optional:      true,
						Description:   "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Required:      true,
								Description:   "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
							},
							"maximum": schema.Int64Attribute{
								Optional:      true,
								Description:   "The maximum unit amount the customer can specify for this item.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
							"minimum": schema.Int64Attribute{
								Optional:      true,
								Description:   "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
							"preset": schema.Int64Attribute{
								Optional:      true,
								Description:   "The starting unit amount which can be updated by the customer.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
						},
					},
					"metadata": schema.MapAttribute{
						Optional:      true,
						Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
						ElementType:   types.StringType,
					},
					"recurring": schema.SingleNestedAttribute{
						Optional:      true,
						Description:   "The recurring components of a price such as `interval` and `interval_count`.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"interval": schema.StringAttribute{
								Required:      true,
								Description:   "Specifies billing frequency. Either `day`, `week`, `month` or `year`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
							"interval_count": schema.Int64Attribute{
								Optional:      true,
								Description:   "The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
						},
					},
					"tax_behavior": schema.StringAttribute{
						Optional:      true,
						Description:   "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"unit_amount": schema.Int64Attribute{
						Optional:      true,
						Description:   "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge. One of `unit_amount`, `unit_amount_decimal`, or `custom_unit_amount` is required.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
					"unit_amount_decimal": schema.StringAttribute{
						Optional:      true,
						Description:   "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
						PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
					},
				},
			},
		},
	}
}

func productResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "Products describe the specific goods or services you offer to your customers.\nFor example, you might offer a Standard and Premium version of your goods or service; each version would be a separate Product.\nThey can be used in conjunction with [Prices](https://api.stripe.com#prices) to configure pricing in Payment Links, Checkout, and Subscriptions.\n\nRelated guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription),\n[share a Payment Link](https://docs.stripe.com/payment-links),\n[accept payments with Checkout](https://docs.stripe.com/payments/accept-a-payment#create-product-prices-upfront),\nand more about [Products and Prices](https://docs.stripe.com/products-prices/overview)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("product")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the product is currently available for purchase.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"default_price": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the [Price](https://docs.stripe.com/api/prices) object that is the default price for this product.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"images": schema.ListAttribute{
				Optional:    true,
				Description: "A list of up to 8 URLs of images for this product, meant to be displayable to the customer.",
				ElementType: types.StringType,
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
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The product's name, meant to be displayable to the customer.",
			},
			"shippable": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether this product is shipped (i.e., physical goods).",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"statement_descriptor": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Extra information about a product which will appear on your customer's credit card statement. In the case that multiple products are billed at once, the first statement descriptor will be used. Only used for subscription payments.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"tax_code": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A [tax code](https://docs.stripe.com/tax/tax-categories) ID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The type of the product. The product is either of type `good`, which is eligible for use with Orders and SKUs, or `service`, which is eligible for use with Subscriptions and Plans.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("good", "service")},
			},
			"unit_label": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"updated": schema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was last updated. Measured in seconds since the Unix epoch.",
			},
			"url": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A URL of a publicly-accessible webpage for this product.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
		Blocks: map[string]schema.Block{
			"marketing_features": schema.ListNestedBlock{
				Description: "A list of up to 15 marketing features for this product. These are displayed in [pricing tables](https://docs.stripe.com/payments/checkout/pricing-table).",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:    true,
							Description: "The marketing feature name. Up to 80 characters long.",
						},
					},
				},
			},
			"package_dimensions": schema.ListNestedBlock{
				Description: "The dimensions of this product for shipping purposes.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"height": schema.Float64Attribute{
							Required:    true,
							Description: "Height, in inches.",
						},
						"length": schema.Float64Attribute{
							Required:    true,
							Description: "Length, in inches.",
						},
						"weight": schema.Float64Attribute{
							Required:    true,
							Description: "Weight, in ounces.",
						},
						"width": schema.Float64Attribute{
							Required:    true,
							Description: "Width, in inches.",
						},
					},
				},
			},
			"default_price_data": schema.ListNestedBlock{
				Description:   "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object. This Price will be set as the default price for this product.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"currency": schema.StringAttribute{
							Required:      true,
							Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"metadata": schema.MapAttribute{
							Optional:      true,
							Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
							PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
							ElementType:   types.StringType,
						},
						"tax_behavior": schema.StringAttribute{
							Optional:      true,
							Description:   "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"unit_amount": schema.Int64Attribute{
							Optional:      true,
							Description:   "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge. One of `unit_amount`, `unit_amount_decimal`, or `custom_unit_amount` is required.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
						"unit_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
						},
					},
					Blocks: map[string]schema.Block{
						"currency_options": schema.ListNestedBlock{
							Description:   "Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
							PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Required:    true,
										Description: "Key for this entry.",
									},
									"tax_behavior": schema.StringAttribute{
										Optional:      true,
										Description:   "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
									"tiers": schema.ListNestedAttribute{
										Optional:      true,
										Description:   "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
										PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"flat_amount": schema.Int64Attribute{
													Optional:      true,
													Description:   "The flat billing amount for an entire tier, regardless of the number of units in the tier.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
												},
												"flat_amount_decimal": schema.StringAttribute{
													Optional:      true,
													Description:   "Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.",
													PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
												},
												"unit_amount": schema.Int64Attribute{
													Optional:      true,
													Description:   "The per unit billing amount for each individual unit for which this tier applies.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
												},
												"unit_amount_decimal": schema.StringAttribute{
													Optional:      true,
													Description:   "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
													PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
												},
												"up_to": schema.Int64Attribute{
													Required:      true,
													Description:   "Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
												},
											},
										},
									},
									"unit_amount": schema.Int64Attribute{
										Optional:      true,
										Description:   "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
									"unit_amount_decimal": schema.StringAttribute{
										Optional:      true,
										Description:   "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
										PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
									},
								},
								Blocks: map[string]schema.Block{
									"custom_unit_amount": schema.ListNestedBlock{
										Description:   "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
										PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"enabled": schema.BoolAttribute{
													Required:      true,
													Description:   "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
													PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
												},
												"maximum": schema.Int64Attribute{
													Optional:      true,
													Description:   "The maximum unit amount the customer can specify for this item.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
												},
												"minimum": schema.Int64Attribute{
													Optional:      true,
													Description:   "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
												},
												"preset": schema.Int64Attribute{
													Optional:      true,
													Description:   "The starting unit amount which can be updated by the customer.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
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
									"enabled": schema.BoolAttribute{
										Required:      true,
										Description:   "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
									},
									"maximum": schema.Int64Attribute{
										Optional:      true,
										Description:   "The maximum unit amount the customer can specify for this item.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
									"minimum": schema.Int64Attribute{
										Optional:      true,
										Description:   "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
									"preset": schema.Int64Attribute{
										Optional:      true,
										Description:   "The starting unit amount which can be updated by the customer.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
								},
							},
						},
						"recurring": schema.ListNestedBlock{
							Description:   "The recurring components of a price such as `interval` and `interval_count`.",
							PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
										Required:      true,
										Description:   "Specifies billing frequency. Either `day`, `week`, `month` or `year`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
									"interval_count": schema.Int64Attribute{
										Optional:      true,
										Description:   "The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
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

type ProductResourceV1Model struct {
	Object              types.String `tfsdk:"object"`
	Active              types.Bool   `tfsdk:"active"`
	Created             types.Int64  `tfsdk:"created"`
	DefaultPrice        types.String `tfsdk:"default_price"`
	Description         types.String `tfsdk:"description"`
	ID                  types.String `tfsdk:"id"`
	Images              types.List   `tfsdk:"images"`
	Livemode            types.Bool   `tfsdk:"livemode"`
	MarketingFeatures   types.List   `tfsdk:"marketing_features"`
	Metadata            types.Map    `tfsdk:"metadata"`
	Name                types.String `tfsdk:"name"`
	PackageDimensions   types.Object `tfsdk:"package_dimensions"`
	Shippable           types.Bool   `tfsdk:"shippable"`
	StatementDescriptor types.String `tfsdk:"statement_descriptor"`
	TaxCode             types.String `tfsdk:"tax_code"`
	Type                types.String `tfsdk:"type"`
	UnitLabel           types.String `tfsdk:"unit_label"`
	Updated             types.Int64  `tfsdk:"updated"`
	URL                 types.String `tfsdk:"url"`
	DefaultPriceData    types.Object `tfsdk:"default_price_data"`
}

type productStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]productStateUpgradeAttrMeta
}

var productStateUpgradeRootMeta = map[string]productStateUpgradeAttrMeta{"object": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "active": productStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "created": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "default_price": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "description": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "id": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "images": productStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.StringType}, Behavior: "optional", LegacyBehavior: "optional"}, "livemode": productStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "marketing_features": productStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]productStateUpgradeAttrMeta{"name": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "metadata": productStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "name": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "package_dimensions": productStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"height": types.Float64Type, "length": types.Float64Type, "weight": types.Float64Type, "width": types.Float64Type}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]productStateUpgradeAttrMeta{"height": productStateUpgradeAttrMeta{AttrType: types.Float64Type, Behavior: "required", LegacyBehavior: "required"}, "length": productStateUpgradeAttrMeta{AttrType: types.Float64Type, Behavior: "required", LegacyBehavior: "required"}, "weight": productStateUpgradeAttrMeta{AttrType: types.Float64Type, Behavior: "required", LegacyBehavior: "required"}, "width": productStateUpgradeAttrMeta{AttrType: types.Float64Type, Behavior: "required", LegacyBehavior: "required"}}}, "shippable": productStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "statement_descriptor": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "tax_code": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "type": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "unit_label": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "updated": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "url": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "default_price_data": productStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "custom_unit_amount": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type}}}, "tax_behavior": types.StringType, "tiers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.Int64Type}}}, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType}}}, "custom_unit_amount": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type}}}, "metadata": types.MapType{ElemType: types.StringType}, "recurring": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", PreserveConfiguredValue: true, Nested: map[string]productStateUpgradeAttrMeta{"currency": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "currency_options": productStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "custom_unit_amount": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type}}}, "tax_behavior": types.StringType, "tiers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.Int64Type}}}, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]productStateUpgradeAttrMeta{"key": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "custom_unit_amount": productStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]productStateUpgradeAttrMeta{"enabled": productStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "required", LegacyBehavior: "required"}, "maximum": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "minimum": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "preset": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}}}, "tax_behavior": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "tiers": productStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.Int64Type}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]productStateUpgradeAttrMeta{"flat_amount": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "flat_amount_decimal": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "unit_amount": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "unit_amount_decimal": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "up_to": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "required", LegacyBehavior: "required"}}}, "unit_amount": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "unit_amount_decimal": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}}}, "custom_unit_amount": productStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]productStateUpgradeAttrMeta{"enabled": productStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "required", LegacyBehavior: "required"}, "maximum": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "minimum": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "preset": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}}}, "metadata": productStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional", LegacyBehavior: "optional"}, "recurring": productStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}}, Behavior: "optional", LegacyBehavior: "optional", Nested: map[string]productStateUpgradeAttrMeta{"interval": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "interval_count": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}}}, "tax_behavior": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "unit_amount": productStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional", LegacyBehavior: "optional"}, "unit_amount_decimal": productStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}}}}

var productStateUpgradeSingletonPaths = map[string]struct{}{}

var productStateUpgradeLegacyObjectPaths = map[string]struct{}{"default_price_data": struct{}{}, "default_price_data.currency_options.custom_unit_amount": struct{}{}, "default_price_data.custom_unit_amount": struct{}{}, "default_price_data.recurring": struct{}{}, "package_dimensions": struct{}{}}

func productAttrMapFromModel(model interface{}) map[string]attr.Value {
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

func productSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
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

func productIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func productShouldPreserveChild(parent productStateUpgradeAttrMeta, child productStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !productIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !productIsComputedBehavior(child.Behavior)
}

func productNullValueForType(attributeType attr.Type) attr.Value {
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

func productLegacyUpgradeIsEmptyValue(value attr.Value) bool {
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

func productLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
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

func productLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := productLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func productLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
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

func productLegacyUpgradeNormalizeChild(parent productStateUpgradeAttrMeta, name string, child productStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if productLegacyUpgradeIsEmptyValue(value) {
		return productNullValueForType(child.AttrType)
	}
	if productLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return productNullValueForType(child.AttrType)
	}
	if productLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return productNullValueForType(child.AttrType)
	}
	return value
}

func productLegacyUpgradeChildAttr(path []string, parent productStateUpgradeAttrMeta, name string, child productStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !productShouldPreserveChild(parent, child) {
		return productNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return productNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := productUpgradeValue(nextPath, child, childValue)
	return productLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func productUpgradeAttrs(path []string, meta map[string]productStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = productNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := productUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = productLegacyUpgradeNormalizeChild(
				productStateUpgradeAttrMeta{PreserveConfiguredValue: true},
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

func productUpgradeObjectValue(path []string, meta productStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = productLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func productUpgradeSingletonListToObject(path []string, meta productStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return productNullValueForType(meta.AttrType)
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
		upgradedAttrs[name] = productLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func productUpgradeObjectValueToSingletonList(path []string, meta productStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	if listValue, ok := priorValue.(types.List); ok {
		return productUpgradeListValue(path, meta, listType, listValue)
	}
	if baseList, ok := priorValue.(basetypes.ListValue); ok {
		return productUpgradeListValue(path, meta, listType, types.List(baseList))
	}

	objectValue, ok := priorValue.(types.Object)
	if !ok {
		if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
			objectValue = types.Object(baseObject)
		} else {
			return productNullValueForType(meta.AttrType)
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
		return productNullValueForType(meta.AttrType)
	}

	upgradedObject := productUpgradeObjectValue(path, meta, elementObjectType, objectValue)
	return types.ListValueMust(listType.ElemType, []attr.Value{upgradedObject})
}

func productUpgradeListValue(path []string, meta productStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return productNullValueForType(meta.AttrType)
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
				upgradedElements = append(upgradedElements, productNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			productUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func productUpgradeValue(path []string, meta productStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := productStateUpgradeSingletonPaths[pathKey]; ok {
			return productUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return productNullValueForType(meta.AttrType)
			}
		}
		return productUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		if _, ok := productStateUpgradeLegacyObjectPaths[pathKey]; ok {
			return productUpgradeObjectValueToSingletonList(path, meta, attrType, priorValue)
		}
		return productUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradeProductStateV1(ctx context.Context, prior interface{}) (ProductResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := productUpgradeAttrs(nil, productStateUpgradeRootMeta, productAttrMapFromModel(prior))
	var upgraded ProductResourceModel
	productSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *ProductResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     2,
		Description: "Products describe the specific goods or services you offer to your customers.\nFor example, you might offer a Standard and Premium version of your goods or service; each version would be a separate Product.\nThey can be used in conjunction with [Prices](https://api.stripe.com#prices) to configure pricing in Payment Links, Checkout, and Subscriptions.\n\nRelated guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription),\n[share a Payment Link](https://docs.stripe.com/payment-links),\n[accept payments with Checkout](https://docs.stripe.com/payments/accept-a-payment#create-product-prices-upfront),\nand more about [Products and Prices](https://docs.stripe.com/products-prices/overview)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("product")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the product is currently available for purchase.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"default_price": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the [Price](https://docs.stripe.com/api/prices) object that is the default price for this product.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"images": schema.ListAttribute{
				Optional:    true,
				Description: "A list of up to 8 URLs of images for this product, meant to be displayable to the customer.",
				ElementType: types.StringType,
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
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The product's name, meant to be displayable to the customer.",
			},
			"shippable": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether this product is shipped (i.e., physical goods).",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"statement_descriptor": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Extra information about a product which will appear on your customer's credit card statement. In the case that multiple products are billed at once, the first statement descriptor will be used. Only used for subscription payments.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"tax_code": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A [tax code](https://docs.stripe.com/tax/tax-categories) ID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The type of the product. The product is either of type `good`, which is eligible for use with Orders and SKUs, or `service`, which is eligible for use with Subscriptions and Plans.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("good", "service")},
			},
			"unit_label": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"updated": schema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was last updated. Measured in seconds since the Unix epoch.",
			},
			"url": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A URL of a publicly-accessible webpage for this product.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
		Blocks: map[string]schema.Block{
			"marketing_features": schema.ListNestedBlock{
				Description: "A list of up to 15 marketing features for this product. These are displayed in [pricing tables](https://docs.stripe.com/payments/checkout/pricing-table).",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:    true,
							Description: "The marketing feature name. Up to 80 characters long.",
						},
					},
				},
			},
			"package_dimensions": schema.ListNestedBlock{
				Description: "The dimensions of this product for shipping purposes.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"height": schema.Float64Attribute{
							Required:    true,
							Description: "Height, in inches.",
						},
						"length": schema.Float64Attribute{
							Required:    true,
							Description: "Length, in inches.",
						},
						"weight": schema.Float64Attribute{
							Required:    true,
							Description: "Weight, in ounces.",
						},
						"width": schema.Float64Attribute{
							Required:    true,
							Description: "Width, in inches.",
						},
					},
				},
			},
			"default_price_data": schema.ListNestedBlock{
				Description:   "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object. This Price will be set as the default price for this product.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"currency": schema.StringAttribute{
							Required:      true,
							Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"metadata": schema.MapAttribute{
							Optional:      true,
							Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
							PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
							ElementType:   types.StringType,
						},
						"tax_behavior": schema.StringAttribute{
							Optional:      true,
							Description:   "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"unit_amount": schema.Int64Attribute{
							Optional:      true,
							Description:   "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge. One of `unit_amount`, `unit_amount_decimal`, or `custom_unit_amount` is required.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
						"unit_amount_decimal": schema.StringAttribute{
							Optional:      true,
							Description:   "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
							PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
						},
					},
					Blocks: map[string]schema.Block{
						"currency_options": schema.ListNestedBlock{
							Description:   "Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
							PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Required:    true,
										Description: "Key for this entry.",
									},
									"tax_behavior": schema.StringAttribute{
										Optional:      true,
										Description:   "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
									"tiers": schema.ListNestedAttribute{
										Optional:      true,
										Description:   "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
										PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"flat_amount": schema.Int64Attribute{
													Optional:      true,
													Description:   "The flat billing amount for an entire tier, regardless of the number of units in the tier.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
												},
												"flat_amount_decimal": schema.StringAttribute{
													Optional:      true,
													Description:   "Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.",
													PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
												},
												"unit_amount": schema.Int64Attribute{
													Optional:      true,
													Description:   "The per unit billing amount for each individual unit for which this tier applies.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
												},
												"unit_amount_decimal": schema.StringAttribute{
													Optional:      true,
													Description:   "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
													PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
												},
												"up_to": schema.Int64Attribute{
													Required:      true,
													Description:   "Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
												},
											},
										},
									},
									"unit_amount": schema.Int64Attribute{
										Optional:      true,
										Description:   "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
									"unit_amount_decimal": schema.StringAttribute{
										Optional:      true,
										Description:   "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
										PlanModifiers: []planmodifier.String{equivalentDecimalStringPlanModifier(), stringplanmodifier.RequiresReplace()},
									},
								},
								Blocks: map[string]schema.Block{
									"custom_unit_amount": schema.ListNestedBlock{
										Description:   "When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.",
										PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"enabled": schema.BoolAttribute{
													Required:      true,
													Description:   "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
													PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
												},
												"maximum": schema.Int64Attribute{
													Optional:      true,
													Description:   "The maximum unit amount the customer can specify for this item.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
												},
												"minimum": schema.Int64Attribute{
													Optional:      true,
													Description:   "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
												},
												"preset": schema.Int64Attribute{
													Optional:      true,
													Description:   "The starting unit amount which can be updated by the customer.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
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
									"enabled": schema.BoolAttribute{
										Required:      true,
										Description:   "Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
									},
									"maximum": schema.Int64Attribute{
										Optional:      true,
										Description:   "The maximum unit amount the customer can specify for this item.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
									"minimum": schema.Int64Attribute{
										Optional:      true,
										Description:   "The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
									"preset": schema.Int64Attribute{
										Optional:      true,
										Description:   "The starting unit amount which can be updated by the customer.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
									},
								},
							},
						},
						"recurring": schema.ListNestedBlock{
							Description:   "The recurring components of a price such as `interval` and `interval_count`.",
							PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
										Required:      true,
										Description:   "Specifies billing frequency. Either `day`, `week`, `month` or `year`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
									"interval_count": schema.Int64Attribute{
										Optional:      true,
										Description:   "The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
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

func (r *ProductResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProductResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandProductCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Product create params", err.Error())
		return
	}

	obj, err := r.client.V1Products.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Product", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Products.B, r.client.V1Products.Key, stripe.FormatURLPath("/v1/products/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Product create raw response", err.Error())
		return
	}

	var createdState ProductResourceModel
	if err := flattenProduct(obj, &createdState); err != nil {
		resp.Diagnostics.AddError("Error flattening Product create response", err.Error())
		return
	}
	normalizeUnknownValues(&createdState)

	diffPlan := plan
	diffCreatedState := createdState

	postCreateParams, err := expandProductPostCreateUpdate(diffPlan, diffCreatedState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Product post-create update params", err.Error())
		return
	}

	if paramsHaveValues(postCreateParams) {
		if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
			!createdState.Metadata.IsNull() && !createdState.Metadata.IsUnknown() {
			if !assignMetadataDiffToNamedField(postCreateParams, "Metadata", plan.Metadata, createdState.Metadata) {
				resp.Diagnostics.AddError("Error building Product update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", postCreateParams))
				return
			}
		}
		obj, err = r.client.V1Products.Update(ctx, createdState.ID.ValueString(), postCreateParams)
		if err != nil {
			resp.Diagnostics.AddError("Error finalizing Product after create", err.Error())
			return
		}
		if err := ensureRawResponse(obj, r.client.V1Products.B, r.client.V1Products.Key, stripe.FormatURLPath("/v1/products/%s", obj.ID), nil); err != nil {
			resp.Diagnostics.AddError("Error hydrating Product post-create update raw response", err.Error())
			return
		}
	}

	if err := flattenProduct(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Product create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ProductResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState ProductResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ProductResourceModel
	state = priorState

	obj, err := r.client.V1Products.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Product", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Products.B, r.client.V1Products.Key, stripe.FormatURLPath("/v1/products/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Product raw response", err.Error())
		return
	}

	if err := flattenProduct(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Product read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *ProductResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ProductResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ProductResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandProductUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Product update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Product update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Products.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Product", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Products.B, r.client.V1Products.Key, stripe.FormatURLPath("/v1/products/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Product update raw response", err.Error())
		return
	}

	if err := flattenProduct(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Product update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ProductResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ProductResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.Active.IsNull() && !state.Active.IsUnknown() && !state.Active.ValueBool() {
		return
	}

	params := &stripe.ProductUpdateParams{}
	activeField := reflect.ValueOf(params).Elem().FieldByName("Active")
	if activeField.IsValid() && activeField.CanSet() {
		if activeField.Kind() == reflect.Pointer && activeField.Type().Elem().Kind() == reflect.Bool {
			activeField.Set(reflect.ValueOf(stripe.Bool(false)))
		} else if activeField.Kind() == reflect.Bool {
			activeField.SetBool(false)
		}
	}

	_, err := r.client.V1Products.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating Product", err.Error())
		return
	}
}

func (r *ProductResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandProductCreate(plan ProductResourceModel) (*stripe.ProductCreateParams, error) {
	params := &stripe.ProductCreateParams{}

	if !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.Images.IsNull() && !plan.Images.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Images", plan.Images) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "images", params)
		}
	}
	if !plan.MarketingFeatures.IsNull() && !plan.MarketingFeatures.IsUnknown() {
		if !assignAttrValueToNamedField(params, "MarketingFeatures", plan.MarketingFeatures) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "marketing_features", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Name", "Name", plan.Name.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "name", params)
		}
	}
	if !plan.PackageDimensions.IsNull() && !plan.PackageDimensions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PackageDimensions", plan.PackageDimensions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "package_dimensions", params)
		}
	}
	if !plan.Shippable.IsNull() && !plan.Shippable.IsUnknown() {
		params.Shippable = stripe.Bool(plan.Shippable.ValueBool())
	}
	if !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "statement_descriptor", params)
		}
	}
	if !plan.TaxCode.IsNull() && !plan.TaxCode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxCodeID", "TaxCode", plan.TaxCode.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_code", params)
		}
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}
	if !plan.UnitLabel.IsNull() && !plan.UnitLabel.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "UnitLabel", "UnitLabel", plan.UnitLabel.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "unit_label", params)
		}
	}
	if !plan.URL.IsNull() && !plan.URL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "URL", "URL", plan.URL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "url", params)
		}
	}
	if !plan.DefaultPriceData.IsNull() && !plan.DefaultPriceData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DefaultPriceData", plan.DefaultPriceData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "default_price_data", params)
		}
	}

	return params, nil
}

func expandProductUpdate(plan ProductResourceModel, state ProductResourceModel) (*stripe.ProductUpdateParams, error) {
	params := &stripe.ProductUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.Description.Equal(state.Description) && !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			if !plan.Description.Equal(state.Description) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "description", params)
			}
		}
	}
	if !plan.Images.Equal(state.Images) && !plan.Images.IsNull() && !plan.Images.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Images", plan.Images) {
			if !plan.Images.Equal(state.Images) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "images", params)
			}
		}
	}
	if !plan.MarketingFeatures.Equal(state.MarketingFeatures) && !plan.MarketingFeatures.IsNull() && !plan.MarketingFeatures.IsUnknown() {
		if !assignAttrValueToNamedField(params, "MarketingFeatures", plan.MarketingFeatures) {
			if !plan.MarketingFeatures.Equal(state.MarketingFeatures) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "marketing_features", params)
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
	if !plan.Name.Equal(state.Name) && !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Name", "Name", plan.Name.ValueString()) {
			if !plan.Name.Equal(state.Name) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "name", params)
			}
		}
	}
	if !plan.PackageDimensions.Equal(state.PackageDimensions) && !plan.PackageDimensions.IsNull() && !plan.PackageDimensions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PackageDimensions", plan.PackageDimensions) {
			if !plan.PackageDimensions.Equal(state.PackageDimensions) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "package_dimensions", params)
			}
		}
	}
	if !plan.Shippable.Equal(state.Shippable) && !plan.Shippable.IsNull() && !plan.Shippable.IsUnknown() {
		params.Shippable = stripe.Bool(plan.Shippable.ValueBool())
	}
	if !plan.StatementDescriptor.Equal(state.StatementDescriptor) && !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			if !plan.StatementDescriptor.Equal(state.StatementDescriptor) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "statement_descriptor", params)
			}
		}
	}
	if !plan.TaxCode.Equal(state.TaxCode) && !plan.TaxCode.IsNull() && !plan.TaxCode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxCodeID", "TaxCode", plan.TaxCode.ValueString()) {
			if !plan.TaxCode.Equal(state.TaxCode) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tax_code", params)
			}
		}
	}
	if !plan.UnitLabel.Equal(state.UnitLabel) && !plan.UnitLabel.IsNull() && !plan.UnitLabel.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "UnitLabel", "UnitLabel", plan.UnitLabel.ValueString()) {
			if !plan.UnitLabel.Equal(state.UnitLabel) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "unit_label", params)
			}
		}
	}
	if !plan.URL.Equal(state.URL) && !plan.URL.IsNull() && !plan.URL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "URL", "URL", plan.URL.ValueString()) {
			if !plan.URL.Equal(state.URL) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "url", params)
			}
		}
	}

	return params, nil
}

func expandProductPostCreateUpdate(plan ProductResourceModel, state ProductResourceModel) (*stripe.ProductUpdateParams, error) {
	params := &stripe.ProductUpdateParams{}

	return params, nil
}

func flattenProduct(obj *stripe.Product, state *ProductResourceModel) error {
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
			if rawValueDefaultPrice, rawOk := plainValueAtPath(raw, "default_price"); rawOk {
				if typedDefaultPrice, ok := plainToStringIDValue(rawValueDefaultPrice); ok {
					state.DefaultPrice = typedDefaultPrice
				}
			} else if !hasRaw {
				if responseValueDefaultPrice, ok := plainFromResponseField(obj, "DefaultPrice"); ok {
					if typedDefaultPrice, ok := plainToStringIDValue(responseValueDefaultPrice); ok {
						state.DefaultPrice = typedDefaultPrice
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
		if rawValueImages, rawOk := plainValueAtPath(raw, "images"); rawOk {
			if !plainValueIsEmpty(applyConfiguredKeyedListShapes(rawValueImages, attrValueToPlain(state.Images))) || state.Images.IsUnknown() || !state.Images.IsNull() {
				if valueImages, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueImages, attrValueToPlain(state.Images)), types.ListType{ElemType: types.StringType}, "images", "raw response"); err != nil {
					return err
				} else {
					if typedImages, ok := valueImages.(types.List); ok {
						state.Images = typedImages
					}
				}
			}
		} else if !hasRaw {
			if responseValueImages, ok := plainFromResponseField(obj, "Images"); ok {
				if !plainValueIsEmpty(applyConfiguredKeyedListShapes(responseValueImages, attrValueToPlain(state.Images))) || state.Images.IsUnknown() || !state.Images.IsNull() {
					if valueImages, err := flattenPlainValue(
						applyConfiguredKeyedListShapes(responseValueImages, attrValueToPlain(state.Images)),
						types.ListType{ElemType: types.StringType},
						"images",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedImages, ok := valueImages.(types.List); ok {
							state.Images = typedImages
						}
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
		if rawValueMarketingFeatures, rawOk := plainValueAtPath(raw, "marketing_features"); rawOk {
			if valueMarketingFeatures, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueMarketingFeatures, attrValueToPlain(state.MarketingFeatures)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType}}}, "marketing_features", "raw response"); err != nil {
				return err
			} else {
				if typedMarketingFeatures, ok := valueMarketingFeatures.(types.List); ok {
					state.MarketingFeatures = typedMarketingFeatures
				}
			}
		} else if !hasRaw {
			if responseValueMarketingFeatures, ok := plainFromResponseField(obj, "MarketingFeatures"); ok {
				if valueMarketingFeatures, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueMarketingFeatures, attrValueToPlain(state.MarketingFeatures)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType}}},
					"marketing_features",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedMarketingFeatures, ok := valueMarketingFeatures.(types.List); ok {
						state.MarketingFeatures = typedMarketingFeatures
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
		if rawValueName, rawOk := plainValueAtPath(raw, "name"); rawOk {
			if valueName, err := flattenPlainValue(rawValueName, types.StringType, "name", "raw response"); err != nil {
				return err
			} else {
				if typedName, ok := valueName.(types.String); ok {
					state.Name = typedName
				}
			}
		} else if !hasRaw {
			if responseValueName, ok := plainFromResponseField(obj, "Name"); ok {
				if valueName, err := flattenPlainValue(responseValueName, types.StringType, "name", "response struct"); err != nil {
					return err
				} else {
					if typedName, ok := valueName.(types.String); ok {
						state.Name = typedName
					}
				}
			}
		}
	}
	{
		assignedPackageDimensions := false
		hadRawPackageDimensions := false
		if rawValuePackageDimensions, rawOk := plainValueAtPath(raw, "package_dimensions"); rawOk {
			hadRawPackageDimensions = true
			if rawValuePackageDimensions != nil {
				sourcePackageDimensions := applyConfiguredKeyedListShapes(rawValuePackageDimensions, unwrapPlainSingletonList(attrValueToPlain(state.PackageDimensions)))
				if valuePackageDimensions, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourcePackageDimensions, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"height": types.Float64Type, "length": types.Float64Type, "weight": types.Float64Type, "width": types.Float64Type}}}, "package_dimensions", "raw response"); err != nil {
					return err
				} else {
					if typedPackageDimensions, ok := valuePackageDimensions.(types.List); ok {
						state.PackageDimensions = typedPackageDimensions
						assignedPackageDimensions = true
					}
				}
			}
		}
		if !assignedPackageDimensions {
			if !hasRaw {
				if responseValuePackageDimensions, ok := plainFromResponseField(obj, "PackageDimensions"); ok {
					sourcePackageDimensions := applyConfiguredKeyedListShapes(responseValuePackageDimensions, unwrapPlainSingletonList(attrValueToPlain(state.PackageDimensions)))
					if valuePackageDimensions, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourcePackageDimensions, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"height": types.Float64Type, "length": types.Float64Type, "weight": types.Float64Type, "width": types.Float64Type}}},
						"package_dimensions",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPackageDimensions, ok := valuePackageDimensions.(types.List); ok {
							state.PackageDimensions = typedPackageDimensions
							assignedPackageDimensions = true
						}
					}
				}
			}
		}
		if !assignedPackageDimensions && hadRawPackageDimensions {
			if nullPackageDimensions, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"height": types.Float64Type, "length": types.Float64Type, "weight": types.Float64Type, "width": types.Float64Type}}}); ok {
				if typedPackageDimensions, ok := nullPackageDimensions.(types.List); ok {
					state.PackageDimensions = typedPackageDimensions
				}
			}
		}
	}
	{
		if rawValueShippable, rawOk := plainValueAtPath(raw, "shippable"); rawOk {
			if valueShippable, err := flattenPlainValue(rawValueShippable, types.BoolType, "shippable", "raw response"); err != nil {
				return err
			} else {
				if typedShippable, ok := valueShippable.(types.Bool); ok {
					state.Shippable = typedShippable
				}
			}
		} else if !hasRaw {
			if responseValueShippable, ok := plainFromResponseField(obj, "Shippable"); ok {
				if valueShippable, err := flattenPlainValue(responseValueShippable, types.BoolType, "shippable", "response struct"); err != nil {
					return err
				} else {
					if typedShippable, ok := valueShippable.(types.Bool); ok {
						state.Shippable = typedShippable
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
		if true {
			if rawValueTaxCode, rawOk := plainValueAtPath(raw, "tax_code"); rawOk {
				if typedTaxCode, ok := plainToStringIDValue(rawValueTaxCode); ok {
					state.TaxCode = typedTaxCode
				}
			} else if !hasRaw {
				if responseValueTaxCode, ok := plainFromResponseField(obj, "TaxCode"); ok {
					if typedTaxCode, ok := plainToStringIDValue(responseValueTaxCode); ok {
						state.TaxCode = typedTaxCode
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
		if rawValueUnitLabel, rawOk := plainValueAtPath(raw, "unit_label"); rawOk {
			if valueUnitLabel, err := flattenPlainValue(rawValueUnitLabel, types.StringType, "unit_label", "raw response"); err != nil {
				return err
			} else {
				if typedUnitLabel, ok := valueUnitLabel.(types.String); ok {
					state.UnitLabel = typedUnitLabel
				}
			}
		} else if !hasRaw {
			if responseValueUnitLabel, ok := plainFromResponseField(obj, "UnitLabel"); ok {
				if valueUnitLabel, err := flattenPlainValue(responseValueUnitLabel, types.StringType, "unit_label", "response struct"); err != nil {
					return err
				} else {
					if typedUnitLabel, ok := valueUnitLabel.(types.String); ok {
						state.UnitLabel = typedUnitLabel
					}
				}
			}
		}
	}
	{
		if rawValueUpdated, rawOk := plainValueAtPath(raw, "updated"); rawOk {
			if valueUpdated, err := flattenPlainValue(rawValueUpdated, types.Int64Type, "updated", "raw response"); err != nil {
				return err
			} else {
				if typedUpdated, ok := valueUpdated.(types.Int64); ok {
					state.Updated = typedUpdated
				}
			}
		} else if !hasRaw {
			if responseValueUpdated, ok := plainFromResponseField(obj, "Updated"); ok {
				if valueUpdated, err := flattenPlainValue(responseValueUpdated, types.Int64Type, "updated", "response struct"); err != nil {
					return err
				} else {
					if typedUpdated, ok := valueUpdated.(types.Int64); ok {
						state.Updated = typedUpdated
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
		assignedDefaultPriceData := false
		if rawValueDefaultPriceData, rawOk := plainValueAtPath(raw, "default_price_data"); rawOk {
			if rawValueDefaultPriceData != nil {
				sourceDefaultPriceData := mergeMissingPlainLeaves(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(suppressUnconfiguredOptionalReadbackLeaves(applyConfiguredKeyedListShapes(applyPlainNestedObjectDefaultLeafValues(rawValueDefaultPriceData, []plainNestedObjectDefaultLeafValues{{ObjectPath: []string{"currency_options", "*", "custom_unit_amount"}, Defaults: []plainObjectDefaultLeafValue{{Target: []string{"enabled"}, Value: true}}}, {ObjectPath: []string{"custom_unit_amount"}, Defaults: []plainObjectDefaultLeafValue{{Target: []string{"enabled"}, Value: true}}}}), unwrapPlainSingletonList(attrValueToPlain(state.DefaultPriceData))), unwrapPlainSingletonList(attrValueToPlain(state.DefaultPriceData)), [][]string{[]string{"currency_options"}, []string{"currency_options", "custom_unit_amount"}, []string{"currency_options", "custom_unit_amount", "maximum"}, []string{"currency_options", "custom_unit_amount", "minimum"}, []string{"currency_options", "custom_unit_amount", "preset"}, []string{"currency_options", "tax_behavior"}, []string{"currency_options", "tiers"}, []string{"currency_options", "tiers", "flat_amount"}, []string{"currency_options", "tiers", "flat_amount_decimal"}, []string{"currency_options", "tiers", "unit_amount"}, []string{"currency_options", "tiers", "unit_amount_decimal"}, []string{"currency_options", "unit_amount"}, []string{"currency_options", "unit_amount_decimal"}, []string{"custom_unit_amount"}, []string{"custom_unit_amount", "maximum"}, []string{"custom_unit_amount", "minimum"}, []string{"custom_unit_amount", "preset"}, []string{"metadata"}, []string{"recurring"}, []string{"recurring", "interval_count"}, []string{"tax_behavior"}, []string{"unit_amount"}, []string{"unit_amount_decimal"}}), unwrapPlainSingletonList(attrValueToPlain(state.DefaultPriceData))), unwrapPlainSingletonList(attrValueToPlain(state.DefaultPriceData))), unwrapPlainSingletonList(attrValueToPlain(state.DefaultPriceData)))
				if valueDefaultPriceData, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceDefaultPriceData, [][]string{[]string{}, []string{"currency_options", "custom_unit_amount"}, []string{"custom_unit_amount"}, []string{"recurring"}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "custom_unit_amount": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type}}}, "tax_behavior": types.StringType, "tiers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.Int64Type}}}, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType}}}, "custom_unit_amount": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type}}}, "metadata": types.MapType{ElemType: types.StringType}, "recurring": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType}}}, "default_price_data", "raw response"); err != nil {
					return err
				} else {
					if typedDefaultPriceData, ok := valueDefaultPriceData.(types.List); ok {
						state.DefaultPriceData = typedDefaultPriceData
						assignedDefaultPriceData = true
					}
				}
			}
		}
		if !assignedDefaultPriceData {
			if !hasRaw {
				if responseValueDefaultPriceData, ok := plainFromResponseField(obj, "DefaultPriceData"); ok {
					sourceDefaultPriceData := mergeMissingPlainLeaves(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(suppressUnconfiguredOptionalReadbackLeaves(applyConfiguredKeyedListShapes(applyPlainNestedObjectDefaultLeafValues(responseValueDefaultPriceData, []plainNestedObjectDefaultLeafValues{{ObjectPath: []string{"currency_options", "*", "custom_unit_amount"}, Defaults: []plainObjectDefaultLeafValue{{Target: []string{"enabled"}, Value: true}}}, {ObjectPath: []string{"custom_unit_amount"}, Defaults: []plainObjectDefaultLeafValue{{Target: []string{"enabled"}, Value: true}}}}), unwrapPlainSingletonList(attrValueToPlain(state.DefaultPriceData))), unwrapPlainSingletonList(attrValueToPlain(state.DefaultPriceData)), [][]string{[]string{"currency_options"}, []string{"currency_options", "custom_unit_amount"}, []string{"currency_options", "custom_unit_amount", "maximum"}, []string{"currency_options", "custom_unit_amount", "minimum"}, []string{"currency_options", "custom_unit_amount", "preset"}, []string{"currency_options", "tax_behavior"}, []string{"currency_options", "tiers"}, []string{"currency_options", "tiers", "flat_amount"}, []string{"currency_options", "tiers", "flat_amount_decimal"}, []string{"currency_options", "tiers", "unit_amount"}, []string{"currency_options", "tiers", "unit_amount_decimal"}, []string{"currency_options", "unit_amount"}, []string{"currency_options", "unit_amount_decimal"}, []string{"custom_unit_amount"}, []string{"custom_unit_amount", "maximum"}, []string{"custom_unit_amount", "minimum"}, []string{"custom_unit_amount", "preset"}, []string{"metadata"}, []string{"recurring"}, []string{"recurring", "interval_count"}, []string{"tax_behavior"}, []string{"unit_amount"}, []string{"unit_amount_decimal"}}), unwrapPlainSingletonList(attrValueToPlain(state.DefaultPriceData))), unwrapPlainSingletonList(attrValueToPlain(state.DefaultPriceData))), unwrapPlainSingletonList(attrValueToPlain(state.DefaultPriceData)))
					if valueDefaultPriceData, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceDefaultPriceData, [][]string{[]string{}, []string{"currency_options", "custom_unit_amount"}, []string{"custom_unit_amount"}, []string{"recurring"}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "currency_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"key": types.StringType, "custom_unit_amount": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type}}}, "tax_behavior": types.StringType, "tiers": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType, "up_to": types.Int64Type}}}, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType}}}, "custom_unit_amount": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type, "preset": types.Int64Type}}}, "metadata": types.MapType{ElemType: types.StringType}, "recurring": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.StringType}}},
						"default_price_data",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedDefaultPriceData, ok := valueDefaultPriceData.(types.List); ok {
							state.DefaultPriceData = typedDefaultPriceData
							assignedDefaultPriceData = true
						}
					}
				}
			}
		}
	}
	return nil
}
