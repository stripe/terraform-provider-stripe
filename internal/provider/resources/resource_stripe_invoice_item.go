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

var _ resource.Resource = &InvoiceItemResource{}

var _ resource.ResourceWithConfigure = &InvoiceItemResource{}

var _ resource.ResourceWithImportState = &InvoiceItemResource{}

func NewInvoiceItemResource() resource.Resource {
	return &InvoiceItemResource{}
}

type InvoiceItemResource struct {
	client *stripe.Client
}

type InvoiceItemResourceModel struct {
	Object            types.String  `tfsdk:"object"`
	Amount            types.Int64   `tfsdk:"amount"`
	Currency          types.String  `tfsdk:"currency"`
	Customer          types.String  `tfsdk:"customer"`
	Date              types.Int64   `tfsdk:"date"`
	Description       types.String  `tfsdk:"description"`
	Discountable      types.Bool    `tfsdk:"discountable"`
	Discounts         types.List    `tfsdk:"discounts"`
	ID                types.String  `tfsdk:"id"`
	Invoice           types.String  `tfsdk:"invoice"`
	Livemode          types.Bool    `tfsdk:"livemode"`
	Metadata          types.Map     `tfsdk:"metadata"`
	Parent            types.Object  `tfsdk:"parent"`
	Period            types.Object  `tfsdk:"period"`
	Pricing           types.Object  `tfsdk:"pricing"`
	Proration         types.Bool    `tfsdk:"proration"`
	ProrationDetails  types.Object  `tfsdk:"proration_details"`
	Quantity          types.Int64   `tfsdk:"quantity"`
	QuantityDecimal   types.Float64 `tfsdk:"quantity_decimal"`
	TaxRates          types.List    `tfsdk:"tax_rates"`
	TestClock         types.String  `tfsdk:"test_clock"`
	PriceData         types.Object  `tfsdk:"price_data"`
	Subscription      types.String  `tfsdk:"subscription"`
	TaxBehavior       types.String  `tfsdk:"tax_behavior"`
	TaxCode           types.String  `tfsdk:"tax_code"`
	UnitAmountDecimal types.Float64 `tfsdk:"unit_amount_decimal"`
}

func (r *InvoiceItemResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *InvoiceItemResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_invoice_item"
}

func (r *InvoiceItemResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Invoice Items represent the component lines of an [invoice](https://docs.stripe.com/api/invoices). When you create an invoice item with an `invoice` field, it is attached to the specified invoice and included as [an invoice line item](https://docs.stripe.com/api/invoices/line_item) within [invoice.lines](https://docs.stripe.com/api/invoices/object#invoice_object-lines).\n\nInvoice Items can be created before you are ready to actually send the invoice. This can be particularly useful when combined\nwith a [subscription](https://docs.stripe.com/api/subscriptions). Sometimes you want to add a charge or credit to a customer, but actually charge\nor credit the customer's card only at the end of a regular billing cycle. This is useful for combining several charges\n(to minimize per-transaction fees), or for having Stripe tabulate your usage-based billing totals.\n\nRelated guides: [Integrate with the Invoicing API](https://docs.stripe.com/invoicing/integration), [Subscription Invoices](https://docs.stripe.com/billing/invoices/subscription#adding-upcoming-invoice-items).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("invoiceitem")},
			},
			"amount": schema.Int64Attribute{
				Optional:    true,
				Description: "Amount (in the `currency` specified) of the invoice item. This should always be equal to `unit_amount * quantity`.",
				WriteOnly:   true,
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
				Description:   "The ID of the customer to bill for this invoice item.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"date": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the object. Often useful for displaying to users.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"discountable": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "If true, discounts will apply to this invoice item. Always false for prorations.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"discounts": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The discounts which apply to the invoice item. Item discounts are applied before invoice discounts. Use `expand[]=discounts` to expand each discount.",
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
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ID of the invoice this invoice item belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
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
			"parent": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The parent that generated this invoice item.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"subscription_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details about the subscription that generated this invoice item",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"subscription": schema.StringAttribute{
								Computed:      true,
								Description:   "The subscription that generated this invoice item",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"subscription_item": schema.StringAttribute{
								Computed:      true,
								Description:   "The subscription item that generated this invoice item",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "The type of parent that generated this invoice item",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("subscription_details")},
					},
				},
			},
			"period": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"end": schema.Int64Attribute{
						Required:    true,
						Description: "The end of the period, which must be greater than or equal to the start. This value is inclusive.",
					},
					"start": schema.Int64Attribute{
						Required:    true,
						Description: "The start of the period. This value is inclusive.",
					},
				},
			},
			"pricing": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "The pricing information of the invoice item.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"price_details": schema.SingleNestedAttribute{
						Optional: true,

						WriteOnly: true,
						Attributes: map[string]schema.Attribute{
							"price": schema.StringAttribute{
								Optional:    true,
								Description: "The ID of the price this item is associated with.",
								WriteOnly:   true,
							},
							"product": schema.StringAttribute{
								Optional:    true,
								Description: "The ID of the product this item is associated with.",
								WriteOnly:   true,
							},
						},
					},
					"type": schema.StringAttribute{
						Optional:    true,
						Description: "The type of the pricing details.",
						WriteOnly:   true,
						Validators:  []validator.String{stringvalidator.OneOf("price_details")},
					},
					"unit_amount_decimal": schema.Float64Attribute{
						Optional:    true,
						Description: "The unit amount (in the `currency` specified) of the item which contains a decimal value with at most 12 decimal places.",
						WriteOnly:   true,
					},
					"price": schema.StringAttribute{
						Optional:    true,
						Description: "The ID of the price object.",
						WriteOnly:   true,
					},
				},
			},
			"proration": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether the invoice item was created automatically as a proration adjustment when the customer switched plans.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"proration_details": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"credited_items": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "For a credit proration, links to the debit invoice line items or invoice item that the credit applies to.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"invoice_item": schema.StringAttribute{
								Computed:      true,
								Description:   "When `type` is `invoice_item`, the invoice item id for the debited invoice item corresponding to this credit proration.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"invoice_line_item_details": schema.SingleNestedAttribute{
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"invoice": schema.StringAttribute{
										Computed:      true,
										Description:   "The invoice id for the debited line item(s).",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"invoice_line_items": schema.ListAttribute{
										Computed:      true,
										Description:   "IDs of the debited invoice line item(s) on the invoice that correspond to the credit proration.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
								},
							},
							"type": schema.StringAttribute{
								Computed:      true,
								Description:   "Whether the credit references a pending invoice item or one or more invoice line items on an invoice.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("invoice_item", "invoice_line_items")},
							},
						},
					},
					"discount_amounts": schema.ListNestedAttribute{
						Computed:      true,
						Description:   "Discount amounts applied when the proration was created.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"amount": schema.Int64Attribute{
									Computed:      true,
									Description:   "The amount, in cents (or local equivalent), of the discount.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"discount": schema.StringAttribute{
									Computed:      true,
									Description:   "The discount that was applied to get this discount amount.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
							},
						},
					},
				},
			},
			"quantity": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Quantity of units for the invoice item in integer format, with any decimal precision truncated. For the item's full-precision decimal quantity, use `quantity_decimal`. This field will be deprecated in favor of `quantity_decimal` in a future version. If the invoice item is a proration, the quantity of the subscription that the proration was computed for.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"quantity_decimal": schema.Float64Attribute{
				Optional:    true,
				Description: "Non-negative decimal with at most 12 decimal places. The quantity of units for the invoice item.",
				WriteOnly:   true,
			},
			"tax_rates": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The tax rates which apply to the invoice item. When set, the `default_tax_rates` on the invoice do not apply to this invoice item.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"test_clock": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the test clock this invoice item belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"price_data": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"currency": schema.StringAttribute{
						Required:    true,
						Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
						WriteOnly:   true,
					},
					"product": schema.StringAttribute{
						Required:    true,
						Description: "The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to.",
						WriteOnly:   true,
					},
					"tax_behavior": schema.StringAttribute{
						Optional:    true,
						Description: "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
						WriteOnly:   true,
					},
					"unit_amount": schema.Int64Attribute{
						Optional:    true,
						Description: "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.",
						WriteOnly:   true,
					},
					"unit_amount_decimal": schema.Float64Attribute{
						Optional:    true,
						Description: "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
						WriteOnly:   true,
					},
				},
			},
			"subscription": schema.StringAttribute{
				Optional:      true,
				Description:   "The ID of a subscription to add this invoice item to. When left blank, the invoice item is added to the next upcoming scheduled invoice. When set, scheduled invoices for subscriptions other than the specified subscription will ignore the invoice item. Use this when you want to express that an invoice item has been accrued within the context of a particular subscription.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"tax_behavior": schema.StringAttribute{
				Optional:    true,
				Description: "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
				WriteOnly:   true,
				Validators:  []validator.String{stringvalidator.OneOf("exclusive", "inclusive", "unspecified")},
			},
			"tax_code": schema.StringAttribute{
				Optional:    true,
				Description: "A [tax code](https://docs.stripe.com/tax/tax-categories) ID.",
				WriteOnly:   true,
			},
			"unit_amount_decimal": schema.Float64Attribute{
				Optional:    true,
				Description: "The decimal unit amount in cents (or local equivalent) of the charge to be applied to the upcoming invoice. This `unit_amount_decimal` will be multiplied by the quantity to get the full amount. Passing in a negative `unit_amount_decimal` will reduce the `amount_due` on the invoice. Accepts at most 12 decimal places.",
				WriteOnly:   true,
			},
		},
	}
}

func (r *InvoiceItemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InvoiceItemResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config InvoiceItemResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Amount"}, []string{"Pricing"}, []string{"Pricing", "price"}, []string{"QuantityDecimal"}, []string{"PriceData"}, []string{"PriceData", "currency"}, []string{"PriceData", "product"}, []string{"PriceData", "tax_behavior"}, []string{"PriceData", "unit_amount"}, []string{"PriceData", "unit_amount_decimal"}, []string{"Subscription"}, []string{"TaxBehavior"}, []string{"TaxCode"}, []string{"UnitAmountDecimal"}})

	params, err := expandInvoiceItemCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building InvoiceItem create params", err.Error())
		return
	}

	obj, err := r.client.V1InvoiceItems.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating InvoiceItem", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1InvoiceItems.B, r.client.V1InvoiceItems.Key, stripe.FormatURLPath("/v1/invoiceitems/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating InvoiceItem create raw response", err.Error())
		return
	}

	if err := flattenInvoiceItem(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening InvoiceItem create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Amount"}, []string{"Pricing"}, []string{"Pricing", "price"}, []string{"QuantityDecimal"}, []string{"PriceData"}, []string{"PriceData", "currency"}, []string{"PriceData", "product"}, []string{"PriceData", "tax_behavior"}, []string{"PriceData", "unit_amount"}, []string{"PriceData", "unit_amount_decimal"}, []string{"Subscription"}, []string{"TaxBehavior"}, []string{"TaxCode"}, []string{"UnitAmountDecimal"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *InvoiceItemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState InvoiceItemResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state InvoiceItemResourceModel
	state = priorState

	obj, err := r.client.V1InvoiceItems.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading InvoiceItem", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1InvoiceItems.B, r.client.V1InvoiceItems.Key, stripe.FormatURLPath("/v1/invoiceitems/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating InvoiceItem raw response", err.Error())
		return
	}

	if err := flattenInvoiceItem(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening InvoiceItem read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"Amount"}, []string{"Pricing"}, []string{"Pricing", "price"}, []string{"QuantityDecimal"}, []string{"PriceData"}, []string{"PriceData", "currency"}, []string{"PriceData", "product"}, []string{"PriceData", "tax_behavior"}, []string{"PriceData", "unit_amount"}, []string{"PriceData", "unit_amount_decimal"}, []string{"Subscription"}, []string{"TaxBehavior"}, []string{"TaxCode"}, []string{"UnitAmountDecimal"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *InvoiceItemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InvoiceItemResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config InvoiceItemResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Amount"}, []string{"Pricing"}, []string{"Pricing", "price"}, []string{"QuantityDecimal"}, []string{"PriceData"}, []string{"PriceData", "currency"}, []string{"PriceData", "product"}, []string{"PriceData", "tax_behavior"}, []string{"PriceData", "unit_amount"}, []string{"PriceData", "unit_amount_decimal"}, []string{"Subscription"}, []string{"TaxBehavior"}, []string{"TaxCode"}, []string{"UnitAmountDecimal"}})

	var state InvoiceItemResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"Subscription"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"Subscription"}})

	params, err := expandInvoiceItemUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building InvoiceItem update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building InvoiceItem update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1InvoiceItems.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating InvoiceItem", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1InvoiceItems.B, r.client.V1InvoiceItems.Key, stripe.FormatURLPath("/v1/invoiceitems/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating InvoiceItem update raw response", err.Error())
		return
	}

	if err := flattenInvoiceItem(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening InvoiceItem update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Amount"}, []string{"Pricing"}, []string{"Pricing", "price"}, []string{"QuantityDecimal"}, []string{"PriceData"}, []string{"PriceData", "currency"}, []string{"PriceData", "product"}, []string{"PriceData", "tax_behavior"}, []string{"PriceData", "unit_amount"}, []string{"PriceData", "unit_amount_decimal"}, []string{"Subscription"}, []string{"TaxBehavior"}, []string{"TaxCode"}, []string{"UnitAmountDecimal"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *InvoiceItemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InvoiceItemResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1InvoiceItems.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting InvoiceItem", err.Error())
		return
	}
}

func (r *InvoiceItemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandInvoiceItemCreate(plan InvoiceItemResourceModel) (*stripe.InvoiceItemCreateParams, error) {
	params := &stripe.InvoiceItemCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
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
	if !plan.Discountable.IsNull() && !plan.Discountable.IsUnknown() {
		params.Discountable = stripe.Bool(plan.Discountable.ValueBool())
	}
	if !plan.Discounts.IsNull() && !plan.Discounts.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Discounts", plan.Discounts) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "discounts", params)
		}
	}
	if !plan.Invoice.IsNull() && !plan.Invoice.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "InvoiceID", "Invoice", plan.Invoice.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "invoice", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Period.IsNull() && !plan.Period.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Period", plan.Period) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "period", params)
		}
	}
	if !plan.Pricing.IsNull() && !plan.Pricing.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Pricing", plan.Pricing) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "pricing", params)
		}
	}
	if !plan.Quantity.IsNull() && !plan.Quantity.IsUnknown() {
		params.Quantity = stripe.Int64(plan.Quantity.ValueInt64())
	}
	if !plan.QuantityDecimal.IsNull() && !plan.QuantityDecimal.IsUnknown() {
		params.QuantityDecimal = stripe.Float64(plan.QuantityDecimal.ValueFloat64())
	}
	if !plan.TaxRates.IsNull() && !plan.TaxRates.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TaxRates", plan.TaxRates) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_rates", params)
		}
	}
	if !plan.PriceData.IsNull() && !plan.PriceData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PriceData", plan.PriceData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "price_data", params)
		}
	}
	if !plan.Subscription.IsNull() && !plan.Subscription.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Subscription", "Subscription", plan.Subscription.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "subscription", params)
		}
	}
	if !plan.TaxBehavior.IsNull() && !plan.TaxBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxBehavior", "TaxBehavior", plan.TaxBehavior.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_behavior", params)
		}
	}
	if !plan.TaxCode.IsNull() && !plan.TaxCode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxCode", "TaxCode", plan.TaxCode.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_code", params)
		}
	}
	if !plan.UnitAmountDecimal.IsNull() && !plan.UnitAmountDecimal.IsUnknown() {
		params.UnitAmountDecimal = stripe.Float64(plan.UnitAmountDecimal.ValueFloat64())
	}

	return params, nil
}

func expandInvoiceItemUpdate(plan InvoiceItemResourceModel, state InvoiceItemResourceModel) (*stripe.InvoiceItemUpdateParams, error) {
	params := &stripe.InvoiceItemUpdateParams{}

	if !plan.Amount.Equal(state.Amount) && !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Description.Equal(state.Description) && !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			if !plan.Description.Equal(state.Description) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "description", params)
			}
		}
	}
	if !plan.Discountable.Equal(state.Discountable) && !plan.Discountable.IsNull() && !plan.Discountable.IsUnknown() {
		params.Discountable = stripe.Bool(plan.Discountable.ValueBool())
	}
	if !plan.Discounts.Equal(state.Discounts) && !plan.Discounts.IsNull() && !plan.Discounts.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Discounts", plan.Discounts) {
			if !plan.Discounts.Equal(state.Discounts) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "discounts", params)
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
	if !plan.Period.Equal(state.Period) && !plan.Period.IsNull() && !plan.Period.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Period", plan.Period) {
			if !plan.Period.Equal(state.Period) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "period", params)
			}
		}
	}
	if !plan.Pricing.Equal(state.Pricing) && !plan.Pricing.IsNull() && !plan.Pricing.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Pricing", plan.Pricing) {
			if !plan.Pricing.Equal(state.Pricing) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "pricing", params)
			}
		}
	}
	if !plan.Quantity.Equal(state.Quantity) && !plan.Quantity.IsNull() && !plan.Quantity.IsUnknown() {
		params.Quantity = stripe.Int64(plan.Quantity.ValueInt64())
	}
	if !plan.QuantityDecimal.Equal(state.QuantityDecimal) && !plan.QuantityDecimal.IsNull() && !plan.QuantityDecimal.IsUnknown() {
		params.QuantityDecimal = stripe.Float64(plan.QuantityDecimal.ValueFloat64())
	}
	if !plan.TaxRates.Equal(state.TaxRates) && !plan.TaxRates.IsNull() && !plan.TaxRates.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TaxRates", plan.TaxRates) {
			if !plan.TaxRates.Equal(state.TaxRates) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tax_rates", params)
			}
		}
	}
	if !plan.PriceData.Equal(state.PriceData) && !plan.PriceData.IsNull() && !plan.PriceData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PriceData", plan.PriceData) {
			if !plan.PriceData.Equal(state.PriceData) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "price_data", params)
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
	if !plan.TaxCode.Equal(state.TaxCode) && !plan.TaxCode.IsNull() && !plan.TaxCode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxCode", "TaxCode", plan.TaxCode.ValueString()) {
			if !plan.TaxCode.Equal(state.TaxCode) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tax_code", params)
			}
		}
	}
	if !plan.UnitAmountDecimal.Equal(state.UnitAmountDecimal) && !plan.UnitAmountDecimal.IsNull() && !plan.UnitAmountDecimal.IsUnknown() {
		params.UnitAmountDecimal = stripe.Float64(plan.UnitAmountDecimal.ValueFloat64())
	}

	return params, nil
}

func flattenInvoiceItem(obj *stripe.InvoiceItem, state *InvoiceItemResourceModel) error {
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
		if rawValueDate, rawOk := plainValueAtPath(raw, "date"); rawOk {
			if valueDate, err := flattenPlainValue(rawValueDate, types.Int64Type, "date", "raw response"); err != nil {
				return err
			} else {
				if typedDate, ok := valueDate.(types.Int64); ok {
					state.Date = typedDate
				}
			}
		} else if !hasRaw {
			if responseValueDate, ok := plainFromResponseField(obj, "Date"); ok {
				if valueDate, err := flattenPlainValue(responseValueDate, types.Int64Type, "date", "response struct"); err != nil {
					return err
				} else {
					if typedDate, ok := valueDate.(types.Int64); ok {
						state.Date = typedDate
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
		if rawValueDiscountable, rawOk := plainValueAtPath(raw, "discountable"); rawOk {
			if valueDiscountable, err := flattenPlainValue(rawValueDiscountable, types.BoolType, "discountable", "raw response"); err != nil {
				return err
			} else {
				if typedDiscountable, ok := valueDiscountable.(types.Bool); ok {
					state.Discountable = typedDiscountable
				}
			}
		} else if !hasRaw {
			if responseValueDiscountable, ok := plainFromResponseField(obj, "Discountable"); ok {
				if valueDiscountable, err := flattenPlainValue(responseValueDiscountable, types.BoolType, "discountable", "response struct"); err != nil {
					return err
				} else {
					if typedDiscountable, ok := valueDiscountable.(types.Bool); ok {
						state.Discountable = typedDiscountable
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
		if state.Invoice.IsNull() || state.Invoice.IsUnknown() {
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
		assignedParent := false
		hadRawParent := false
		if rawValueParent, rawOk := plainValueAtPath(raw, "parent"); rawOk {
			hadRawParent = true
			if rawValueParent != nil {
				sourceParent := applyConfiguredKeyedListShapes(rawValueParent, attrValueToPlain(state.Parent))
				if valueParent, err := flattenPlainValue(sourceParent, types.ObjectType{AttrTypes: map[string]attr.Type{"subscription_details": types.ObjectType{AttrTypes: map[string]attr.Type{"subscription": types.StringType, "subscription_item": types.StringType}}, "type": types.StringType}}, "parent", "raw response"); err != nil {
					return err
				} else {
					if typedParent, ok := valueParent.(types.Object); ok {
						state.Parent = typedParent
						assignedParent = true
					}
				}
			}
		}
		if !assignedParent {
			if !hasRaw {
				if responseValueParent, ok := plainFromResponseField(obj, "Parent"); ok {
					sourceParent := applyConfiguredKeyedListShapes(responseValueParent, attrValueToPlain(state.Parent))
					if valueParent, err := flattenPlainValue(
						sourceParent,
						types.ObjectType{AttrTypes: map[string]attr.Type{"subscription_details": types.ObjectType{AttrTypes: map[string]attr.Type{"subscription": types.StringType, "subscription_item": types.StringType}}, "type": types.StringType}},
						"parent",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedParent, ok := valueParent.(types.Object); ok {
							state.Parent = typedParent
							assignedParent = true
						}
					}
				}
			}
		}
		if !assignedParent && hadRawParent {
			if nullParent, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"subscription_details": types.ObjectType{AttrTypes: map[string]attr.Type{"subscription": types.StringType, "subscription_item": types.StringType}}, "type": types.StringType}}); ok {
				if typedParent, ok := nullParent.(types.Object); ok {
					state.Parent = typedParent
				}
			}
		}
	}
	{
		assignedPeriod := false
		hadRawPeriod := false
		if rawValuePeriod, rawOk := plainValueAtPath(raw, "period"); rawOk {
			hadRawPeriod = true
			if rawValuePeriod != nil {
				sourcePeriod := applyConfiguredKeyedListShapes(rawValuePeriod, attrValueToPlain(state.Period))
				if valuePeriod, err := flattenPlainValue(sourcePeriod, types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.Int64Type, "start": types.Int64Type}}, "period", "raw response"); err != nil {
					return err
				} else {
					if typedPeriod, ok := valuePeriod.(types.Object); ok {
						state.Period = typedPeriod
						assignedPeriod = true
					}
				}
			}
		}
		if !assignedPeriod {
			if !hasRaw {
				if responseValuePeriod, ok := plainFromResponseField(obj, "Period"); ok {
					sourcePeriod := applyConfiguredKeyedListShapes(responseValuePeriod, attrValueToPlain(state.Period))
					if valuePeriod, err := flattenPlainValue(
						sourcePeriod,
						types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.Int64Type, "start": types.Int64Type}},
						"period",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPeriod, ok := valuePeriod.(types.Object); ok {
							state.Period = typedPeriod
							assignedPeriod = true
						}
					}
				}
			}
		}
		if !assignedPeriod && hadRawPeriod {
			if nullPeriod, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.Int64Type, "start": types.Int64Type}}); ok {
				if typedPeriod, ok := nullPeriod.(types.Object); ok {
					state.Period = typedPeriod
				}
			}
		}
	}
	{
		assignedPricing := false
		hadRawPricing := false
		if rawValuePricing, rawOk := plainValueAtPath(raw, "pricing"); rawOk {
			hadRawPricing = true
			if rawValuePricing != nil {
				sourcePricing := applyPlainStringLeafAliases(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawValuePricing, attrValueToPlain(state.Pricing)), attrValueToPlain(state.Pricing)), attrValueToPlain(state.Pricing)), []plainStringLeafAlias{{Target: []string{"price"}, Source: []string{"price_details", "price"}}})
				if valuePricing, err := flattenPlainValue(sourcePricing, types.ObjectType{AttrTypes: map[string]attr.Type{"price_details": types.ObjectType{AttrTypes: map[string]attr.Type{"price": types.StringType, "product": types.StringType}}, "type": types.StringType, "unit_amount_decimal": types.Float64Type, "price": types.StringType}}, "pricing", "raw response"); err != nil {
					return err
				} else {
					if typedPricing, ok := valuePricing.(types.Object); ok {
						state.Pricing = typedPricing
						assignedPricing = true
					}
				}
			}
		}
		if !assignedPricing {
			if !hasRaw {
				if responseValuePricing, ok := plainFromResponseField(obj, "Pricing"); ok {
					sourcePricing := applyPlainStringLeafAliases(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(responseValuePricing, attrValueToPlain(state.Pricing)), attrValueToPlain(state.Pricing)), attrValueToPlain(state.Pricing)), []plainStringLeafAlias{{Target: []string{"price"}, Source: []string{"price_details", "price"}}})
					if valuePricing, err := flattenPlainValue(
						sourcePricing,
						types.ObjectType{AttrTypes: map[string]attr.Type{"price_details": types.ObjectType{AttrTypes: map[string]attr.Type{"price": types.StringType, "product": types.StringType}}, "type": types.StringType, "unit_amount_decimal": types.Float64Type, "price": types.StringType}},
						"pricing",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPricing, ok := valuePricing.(types.Object); ok {
							state.Pricing = typedPricing
							assignedPricing = true
						}
					}
				}
			}
		}
		if !assignedPricing && hadRawPricing {
			if nullPricing, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"price_details": types.ObjectType{AttrTypes: map[string]attr.Type{"price": types.StringType, "product": types.StringType}}, "type": types.StringType, "unit_amount_decimal": types.Float64Type, "price": types.StringType}}); ok {
				if typedPricing, ok := nullPricing.(types.Object); ok {
					state.Pricing = typedPricing
				}
			}
		}
	}
	{
		if rawValueProration, rawOk := plainValueAtPath(raw, "proration"); rawOk {
			if valueProration, err := flattenPlainValue(rawValueProration, types.BoolType, "proration", "raw response"); err != nil {
				return err
			} else {
				if typedProration, ok := valueProration.(types.Bool); ok {
					state.Proration = typedProration
				}
			}
		} else if !hasRaw {
			if responseValueProration, ok := plainFromResponseField(obj, "Proration"); ok {
				if valueProration, err := flattenPlainValue(responseValueProration, types.BoolType, "proration", "response struct"); err != nil {
					return err
				} else {
					if typedProration, ok := valueProration.(types.Bool); ok {
						state.Proration = typedProration
					}
				}
			}
		}
	}
	{
		assignedProrationDetails := false
		hadRawProrationDetails := false
		if rawValueProrationDetails, rawOk := plainValueAtPath(raw, "proration_details"); rawOk {
			hadRawProrationDetails = true
			if rawValueProrationDetails != nil {
				sourceProrationDetails := applyConfiguredKeyedListShapes(rawValueProrationDetails, attrValueToPlain(state.ProrationDetails))
				if valueProrationDetails, err := flattenPlainValue(sourceProrationDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"credited_items": types.ObjectType{AttrTypes: map[string]attr.Type{"invoice_item": types.StringType, "invoice_line_item_details": types.ObjectType{AttrTypes: map[string]attr.Type{"invoice": types.StringType, "invoice_line_items": types.ListType{ElemType: types.StringType}}}, "type": types.StringType}}, "discount_amounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "discount": types.StringType}}}}}, "proration_details", "raw response"); err != nil {
					return err
				} else {
					if typedProrationDetails, ok := valueProrationDetails.(types.Object); ok {
						state.ProrationDetails = typedProrationDetails
						assignedProrationDetails = true
					}
				}
			}
		}
		if !assignedProrationDetails {
			if !hasRaw {
				if responseValueProrationDetails, ok := plainFromResponseField(obj, "ProrationDetails"); ok {
					sourceProrationDetails := applyConfiguredKeyedListShapes(responseValueProrationDetails, attrValueToPlain(state.ProrationDetails))
					if valueProrationDetails, err := flattenPlainValue(
						sourceProrationDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"credited_items": types.ObjectType{AttrTypes: map[string]attr.Type{"invoice_item": types.StringType, "invoice_line_item_details": types.ObjectType{AttrTypes: map[string]attr.Type{"invoice": types.StringType, "invoice_line_items": types.ListType{ElemType: types.StringType}}}, "type": types.StringType}}, "discount_amounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "discount": types.StringType}}}}},
						"proration_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedProrationDetails, ok := valueProrationDetails.(types.Object); ok {
							state.ProrationDetails = typedProrationDetails
							assignedProrationDetails = true
						}
					}
				}
			}
		}
		if !assignedProrationDetails && hadRawProrationDetails {
			if nullProrationDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"credited_items": types.ObjectType{AttrTypes: map[string]attr.Type{"invoice_item": types.StringType, "invoice_line_item_details": types.ObjectType{AttrTypes: map[string]attr.Type{"invoice": types.StringType, "invoice_line_items": types.ListType{ElemType: types.StringType}}}, "type": types.StringType}}, "discount_amounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "discount": types.StringType}}}}}); ok {
				if typedProrationDetails, ok := nullProrationDetails.(types.Object); ok {
					state.ProrationDetails = typedProrationDetails
				}
			}
		}
	}
	{
		if rawValueQuantity, rawOk := plainValueAtPath(raw, "quantity"); rawOk {
			if valueQuantity, err := flattenPlainValue(rawValueQuantity, types.Int64Type, "quantity", "raw response"); err != nil {
				return err
			} else {
				if typedQuantity, ok := valueQuantity.(types.Int64); ok {
					state.Quantity = typedQuantity
				}
			}
		} else if !hasRaw {
			if responseValueQuantity, ok := plainFromResponseField(obj, "Quantity"); ok {
				if valueQuantity, err := flattenPlainValue(responseValueQuantity, types.Int64Type, "quantity", "response struct"); err != nil {
					return err
				} else {
					if typedQuantity, ok := valueQuantity.(types.Int64); ok {
						state.Quantity = typedQuantity
					}
				}
			}
		}
	}
	{
		if rawValueQuantityDecimal, rawOk := plainValueAtPath(raw, "quantity_decimal"); rawOk {
			if valueQuantityDecimal, err := flattenPlainValue(rawValueQuantityDecimal, types.Float64Type, "quantity_decimal", "raw response"); err != nil {
				return err
			} else {
				if typedQuantityDecimal, ok := valueQuantityDecimal.(types.Float64); ok {
					state.QuantityDecimal = typedQuantityDecimal
				}
			}
		} else if !hasRaw {
			if responseValueQuantityDecimal, ok := plainFromResponseField(obj, "QuantityDecimal"); ok {
				if valueQuantityDecimal, err := flattenPlainValue(responseValueQuantityDecimal, types.Float64Type, "quantity_decimal", "response struct"); err != nil {
					return err
				} else {
					if typedQuantityDecimal, ok := valueQuantityDecimal.(types.Float64); ok {
						state.QuantityDecimal = typedQuantityDecimal
					}
				}
			}
		}
	}
	{
		if rawValueTaxRates, rawOk := plainValueAtPath(raw, "tax_rates"); rawOk {
			if valueTaxRates, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueTaxRates, attrValueToPlain(state.TaxRates)), types.ListType{ElemType: types.StringType}, "tax_rates", "raw response"); err != nil {
				return err
			} else {
				if typedTaxRates, ok := valueTaxRates.(types.List); ok {
					state.TaxRates = typedTaxRates
				}
			}
		} else if !hasRaw {
			if responseValueTaxRates, ok := plainFromResponseField(obj, "TaxRates"); ok {
				if valueTaxRates, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueTaxRates, attrValueToPlain(state.TaxRates)),
					types.ListType{ElemType: types.StringType},
					"tax_rates",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedTaxRates, ok := valueTaxRates.(types.List); ok {
						state.TaxRates = typedTaxRates
					}
				}
			}
		}
	}
	{
		if true {
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
	return nil
}
