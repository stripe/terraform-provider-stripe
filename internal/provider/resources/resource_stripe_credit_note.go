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

var _ resource.Resource = &CreditNoteResource{}

var _ resource.ResourceWithConfigure = &CreditNoteResource{}

var _ resource.ResourceWithImportState = &CreditNoteResource{}

func NewCreditNoteResource() resource.Resource {
	return &CreditNoteResource{}
}

type CreditNoteResource struct {
	client *stripe.Client
}

type CreditNoteResourceModel struct {
	Object                     types.String `tfsdk:"object"`
	Amount                     types.Int64  `tfsdk:"amount"`
	AmountShipping             types.Int64  `tfsdk:"amount_shipping"`
	Created                    types.Int64  `tfsdk:"created"`
	Currency                   types.String `tfsdk:"currency"`
	Customer                   types.String `tfsdk:"customer"`
	CustomerBalanceTransaction types.String `tfsdk:"customer_balance_transaction"`
	DiscountAmount             types.Int64  `tfsdk:"discount_amount"`
	DiscountAmounts            types.List   `tfsdk:"discount_amounts"`
	EffectiveAt                types.Int64  `tfsdk:"effective_at"`
	ID                         types.String `tfsdk:"id"`
	Invoice                    types.String `tfsdk:"invoice"`
	Livemode                   types.Bool   `tfsdk:"livemode"`
	Memo                       types.String `tfsdk:"memo"`
	Metadata                   types.Map    `tfsdk:"metadata"`
	Number                     types.String `tfsdk:"number"`
	OutOfBandAmount            types.Int64  `tfsdk:"out_of_band_amount"`
	PostPaymentAmount          types.Int64  `tfsdk:"post_payment_amount"`
	PrePaymentAmount           types.Int64  `tfsdk:"pre_payment_amount"`
	PretaxCreditAmounts        types.List   `tfsdk:"pretax_credit_amounts"`
	Reason                     types.String `tfsdk:"reason"`
	Refunds                    types.List   `tfsdk:"refunds"`
	ShippingCost               types.Object `tfsdk:"shipping_cost"`
	Status                     types.String `tfsdk:"status"`
	Subtotal                   types.Int64  `tfsdk:"subtotal"`
	SubtotalExcludingTax       types.Int64  `tfsdk:"subtotal_excluding_tax"`
	Total                      types.Int64  `tfsdk:"total"`
	TotalExcludingTax          types.Int64  `tfsdk:"total_excluding_tax"`
	TotalTaxes                 types.List   `tfsdk:"total_taxes"`
	Type                       types.String `tfsdk:"type"`
	VoidedAt                   types.Int64  `tfsdk:"voided_at"`
	CreditAmount               types.Int64  `tfsdk:"credit_amount"`
	EmailType                  types.String `tfsdk:"email_type"`
	RefundAmount               types.Int64  `tfsdk:"refund_amount"`
}

func (r *CreditNoteResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CreditNoteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_credit_note"
}

func (r *CreditNoteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Issue a credit note to adjust an invoice's amount after the invoice is finalized.\n\nRelated guide: [Credit notes](https://docs.stripe.com/billing/invoices/credit-notes)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("credit_note")},
			},
			"amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The integer amount in cents (or local equivalent) representing the total amount of the credit note, including tax.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"amount_shipping": schema.Int64Attribute{
				Computed:      true,
				Description:   "This is the sum of all the shipping amounts.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
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
				Computed:      true,
				Description:   "ID of the customer.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_balance_transaction": schema.StringAttribute{
				Computed:      true,
				Description:   "Customer balance transaction related to this credit note.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"discount_amount": schema.Int64Attribute{
				Computed:      true,
				Description:   "The integer amount in cents (or local equivalent) representing the total amount of discount that was credited.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"discount_amounts": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "The aggregate amounts calculated per discount for all line items.",
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
			"effective_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The date when this credit note is in effect. Same as `created` unless overwritten. When defined, this value replaces the system-generated 'Date of issue' printed on the credit note PDF.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice": schema.StringAttribute{
				Required:      true,
				Description:   "ID of the invoice.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"memo": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Customer-facing text that appears on the credit note PDF.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
				Description:   "A unique number that identifies this particular credit note and appears on the PDF of the credit note and its associated invoice.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"out_of_band_amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Amount that was credited outside of Stripe.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"post_payment_amount": schema.Int64Attribute{
				Computed:      true,
				Description:   "The amount of the credit note that was refunded to the customer, credited to the customer's balance, credited outside of Stripe, or any combination thereof.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"pre_payment_amount": schema.Int64Attribute{
				Computed:      true,
				Description:   "The amount of the credit note by which the invoice's `amount_remaining` and `amount_due` were reduced.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"pretax_credit_amounts": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "The pretax credit amounts (ex: discount, credit grants, etc) for all line items.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"amount": schema.Int64Attribute{
							Computed:      true,
							Description:   "The amount, in cents (or local equivalent), of the pretax credit amount.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"credit_balance_transaction": schema.StringAttribute{
							Computed:      true,
							Description:   "The credit balance transaction that was applied to get this pretax credit amount.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"discount": schema.StringAttribute{
							Computed:      true,
							Description:   "The discount that was applied to get this pretax credit amount.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"type": schema.StringAttribute{
							Computed:      true,
							Description:   "Type of the pretax credit amount referenced.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("credit_balance_transaction", "discount")},
						},
					},
				},
			},
			"reason": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("duplicate", "fraudulent", "order_change", "product_unsatisfactory")},
			},
			"refunds": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Refunds related to this credit note.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"amount_refunded": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "Amount of the refund that applies to this credit note, in cents (or local equivalent).",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"payment_record_refund": schema.SingleNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The PaymentRecord refund details associated with this credit note refund.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
							Attributes: map[string]schema.Attribute{
								"payment_record": schema.StringAttribute{
									Required:      true,
									Description:   "ID of the payment record.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								},
								"refund_group": schema.StringAttribute{
									Required:      true,
									Description:   "ID of the refund group.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								},
							},
						},
						"refund": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "ID of the refund.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						},
						"type": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Type of the refund, one of `refund` or `payment_record_refund`.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
							Validators:    []validator.String{stringvalidator.OneOf("payment_record_refund", "refund")},
						},
					},
				},
			},
			"shipping_cost": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The details of the cost of shipping, including the ShippingRate applied to the invoice.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"amount_subtotal": schema.Int64Attribute{
						Computed:      true,
						Description:   "Total shipping cost before any taxes are applied.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"amount_tax": schema.Int64Attribute{
						Computed:      true,
						Description:   "Total tax amount applied due to shipping costs. If no tax was applied, defaults to 0.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"amount_total": schema.Int64Attribute{
						Computed:      true,
						Description:   "Total shipping cost after taxes are applied.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"shipping_rate": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The ID of the ShippingRate for this invoice.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"taxes": schema.ListNestedAttribute{
						Computed:      true,
						Description:   "The taxes applied to the shipping rate.",
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
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Status of this credit note, one of `issued` or `void`. Learn more about [voiding credit notes](https://docs.stripe.com/billing/invoices/credit-notes#voiding).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("issued", "void")},
			},
			"subtotal": schema.Int64Attribute{
				Computed:      true,
				Description:   "The integer amount in cents (or local equivalent) representing the amount of the credit note, excluding exclusive tax and invoice level discounts.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"subtotal_excluding_tax": schema.Int64Attribute{
				Computed:      true,
				Description:   "The integer amount in cents (or local equivalent) representing the amount of the credit note, excluding all tax and invoice level discounts.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"total": schema.Int64Attribute{
				Computed:      true,
				Description:   "The integer amount in cents (or local equivalent) representing the total amount of the credit note, including tax and all discount.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"total_excluding_tax": schema.Int64Attribute{
				Computed:      true,
				Description:   "The integer amount in cents (or local equivalent) representing the total amount of the credit note, excluding tax, but including discounts.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"total_taxes": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "The aggregate tax information for all line items.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"amount": schema.Int64Attribute{
							Computed:      true,
							Description:   "The amount of the tax, in cents (or local equivalent).",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"tax_behavior": schema.StringAttribute{
							Computed:      true,
							Description:   "Whether this tax is inclusive or exclusive.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("exclusive", "inclusive")},
						},
						"tax_rate_details": schema.SingleNestedAttribute{
							Computed:      true,
							Description:   "Additional details about the tax rate. Only present when `type` is `tax_rate_details`.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"tax_rate": schema.StringAttribute{
									Computed:      true,
									Description:   "ID of the tax rate",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
							},
						},
						"taxability_reason": schema.StringAttribute{
							Computed:      true,
							Description:   "The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("customer_exempt", "not_available", "not_collecting", "not_subject_to_tax", "not_supported", "portion_product_exempt", "portion_reduced_rated", "portion_standard_rated", "product_exempt", "product_exempt_holiday", "proportionally_rated", "reduced_rated", "reverse_charge", "standard_rated", "taxable_basis_reduced", "zero_rated")},
						},
						"taxable_amount": schema.Int64Attribute{
							Computed:      true,
							Description:   "The amount on which tax is calculated, in cents (or local equivalent).",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"type": schema.StringAttribute{
							Computed:      true,
							Description:   "The type of tax information.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("tax_rate_details")},
						},
					},
				},
			},
			"type": schema.StringAttribute{
				Computed:      true,
				Description:   "Type of this credit note, one of `pre_payment` or `post_payment`. A `pre_payment` credit note means it was issued when the invoice was open. A `post_payment` credit note means it was issued when the invoice was paid.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("mixed", "post_payment", "pre_payment")},
			},
			"voided_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "The time that the credit note was voided.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"credit_amount": schema.Int64Attribute{
				Optional:      true,
				Description:   "The integer amount in cents (or local equivalent) representing the amount to credit the customer's balance, which will be automatically applied to their next invoice.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"email_type": schema.StringAttribute{
				Optional:      true,
				Description:   "Type of email to send to the customer, one of `credit_note` or `none` and the default is `credit_note`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
				Validators:    []validator.String{stringvalidator.OneOf("credit_note", "none")},
			},
			"refund_amount": schema.Int64Attribute{
				Optional:      true,
				Description:   "The integer amount in cents (or local equivalent) representing the amount to refund. If set, a refund will be created for the charge associated with the invoice.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
		},
	}
}

func (r *CreditNoteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CreditNoteResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config CreditNoteResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"CreditAmount"}, []string{"EmailType"}, []string{"RefundAmount"}})

	params, err := expandCreditNoteCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building CreditNote create params", err.Error())
		return
	}

	obj, err := r.client.V1CreditNotes.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating CreditNote", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1CreditNotes.B, r.client.V1CreditNotes.Key, stripe.FormatURLPath("/v1/credit_notes/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating CreditNote create raw response", err.Error())
		return
	}

	if err := flattenCreditNote(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening CreditNote create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"CreditAmount"}, []string{"EmailType"}, []string{"RefundAmount"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CreditNoteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState CreditNoteResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CreditNoteResourceModel
	state = priorState

	obj, err := r.client.V1CreditNotes.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading CreditNote", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1CreditNotes.B, r.client.V1CreditNotes.Key, stripe.FormatURLPath("/v1/credit_notes/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating CreditNote raw response", err.Error())
		return
	}

	if err := flattenCreditNote(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening CreditNote read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"CreditAmount"}, []string{"EmailType"}, []string{"RefundAmount"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *CreditNoteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CreditNoteResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config CreditNoteResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"CreditAmount"}, []string{"EmailType"}, []string{"RefundAmount"}})

	var state CreditNoteResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"CreditAmount"}, []string{"EmailType"}, []string{"RefundAmount"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"CreditAmount"}, []string{"EmailType"}, []string{"RefundAmount"}})

	params, err := expandCreditNoteUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building CreditNote update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building CreditNote update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1CreditNotes.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating CreditNote", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1CreditNotes.B, r.client.V1CreditNotes.Key, stripe.FormatURLPath("/v1/credit_notes/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating CreditNote update raw response", err.Error())
		return
	}

	if err := flattenCreditNote(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening CreditNote update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"CreditAmount"}, []string{"EmailType"}, []string{"RefundAmount"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CreditNoteResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *CreditNoteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandCreditNoteCreate(plan CreditNoteResourceModel) (*stripe.CreditNoteCreateParams, error) {
	params := &stripe.CreditNoteCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.EffectiveAt.IsNull() && !plan.EffectiveAt.IsUnknown() {
		params.EffectiveAt = stripe.Int64(plan.EffectiveAt.ValueInt64())
	}
	if !plan.Invoice.IsNull() && !plan.Invoice.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "InvoiceID", "Invoice", plan.Invoice.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "invoice", params)
		}
	}
	if !plan.Memo.IsNull() && !plan.Memo.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Memo", "Memo", plan.Memo.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "memo", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.OutOfBandAmount.IsNull() && !plan.OutOfBandAmount.IsUnknown() {
		params.OutOfBandAmount = stripe.Int64(plan.OutOfBandAmount.ValueInt64())
	}
	if !plan.Reason.IsNull() && !plan.Reason.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Reason", "Reason", plan.Reason.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "reason", params)
		}
	}
	if !plan.Refunds.IsNull() && !plan.Refunds.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Refunds", plan.Refunds) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "refunds", params)
		}
	}
	if !plan.ShippingCost.IsNull() && !plan.ShippingCost.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingCost", plan.ShippingCost) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping_cost", params)
		}
	}
	if !plan.CreditAmount.IsNull() && !plan.CreditAmount.IsUnknown() {
		params.CreditAmount = stripe.Int64(plan.CreditAmount.ValueInt64())
	}
	if !plan.EmailType.IsNull() && !plan.EmailType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "EmailType", "EmailType", plan.EmailType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "email_type", params)
		}
	}
	if !plan.RefundAmount.IsNull() && !plan.RefundAmount.IsUnknown() {
		params.RefundAmount = stripe.Int64(plan.RefundAmount.ValueInt64())
	}

	return params, nil
}

func expandCreditNoteUpdate(plan CreditNoteResourceModel, state CreditNoteResourceModel) (*stripe.CreditNoteUpdateParams, error) {
	params := &stripe.CreditNoteUpdateParams{}

	if !plan.Memo.Equal(state.Memo) && !plan.Memo.IsNull() && !plan.Memo.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Memo", "Memo", plan.Memo.ValueString()) {
			if !plan.Memo.Equal(state.Memo) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "memo", params)
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

	return params, nil
}

func flattenCreditNote(obj *stripe.CreditNote, state *CreditNoteResourceModel) error {
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
		if rawValueAmountShipping, rawOk := plainValueAtPath(raw, "amount_shipping"); rawOk {
			if valueAmountShipping, err := flattenPlainValue(rawValueAmountShipping, types.Int64Type, "amount_shipping", "raw response"); err != nil {
				return err
			} else {
				if typedAmountShipping, ok := valueAmountShipping.(types.Int64); ok {
					state.AmountShipping = typedAmountShipping
				}
			}
		} else if !hasRaw {
			if responseValueAmountShipping, ok := plainFromResponseField(obj, "AmountShipping"); ok {
				if valueAmountShipping, err := flattenPlainValue(responseValueAmountShipping, types.Int64Type, "amount_shipping", "response struct"); err != nil {
					return err
				} else {
					if typedAmountShipping, ok := valueAmountShipping.(types.Int64); ok {
						state.AmountShipping = typedAmountShipping
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
		if true {
			if rawValueCustomerBalanceTransaction, rawOk := plainValueAtPath(raw, "customer_balance_transaction"); rawOk {
				if typedCustomerBalanceTransaction, ok := plainToStringIDValue(rawValueCustomerBalanceTransaction); ok {
					state.CustomerBalanceTransaction = typedCustomerBalanceTransaction
				}
			} else if !hasRaw {
				if responseValueCustomerBalanceTransaction, ok := plainFromResponseField(obj, "CustomerBalanceTransaction"); ok {
					if typedCustomerBalanceTransaction, ok := plainToStringIDValue(responseValueCustomerBalanceTransaction); ok {
						state.CustomerBalanceTransaction = typedCustomerBalanceTransaction
					}
				}
			}
		}
	}
	{
		if rawValueDiscountAmount, rawOk := plainValueAtPath(raw, "discount_amount"); rawOk {
			if valueDiscountAmount, err := flattenPlainValue(rawValueDiscountAmount, types.Int64Type, "discount_amount", "raw response"); err != nil {
				return err
			} else {
				if typedDiscountAmount, ok := valueDiscountAmount.(types.Int64); ok {
					state.DiscountAmount = typedDiscountAmount
				}
			}
		} else if !hasRaw {
			if responseValueDiscountAmount, ok := plainFromResponseField(obj, "DiscountAmount"); ok {
				if valueDiscountAmount, err := flattenPlainValue(responseValueDiscountAmount, types.Int64Type, "discount_amount", "response struct"); err != nil {
					return err
				} else {
					if typedDiscountAmount, ok := valueDiscountAmount.(types.Int64); ok {
						state.DiscountAmount = typedDiscountAmount
					}
				}
			}
		}
	}
	{
		if rawValueDiscountAmounts, rawOk := plainValueAtPath(raw, "discount_amounts"); rawOk {
			if valueDiscountAmounts, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueDiscountAmounts, attrValueToPlain(state.DiscountAmounts)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "discount": types.StringType}}}, "discount_amounts", "raw response"); err != nil {
				return err
			} else {
				if typedDiscountAmounts, ok := valueDiscountAmounts.(types.List); ok {
					state.DiscountAmounts = typedDiscountAmounts
				}
			}
		} else if !hasRaw {
			if responseValueDiscountAmounts, ok := plainFromResponseField(obj, "DiscountAmounts"); ok {
				if valueDiscountAmounts, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueDiscountAmounts, attrValueToPlain(state.DiscountAmounts)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "discount": types.StringType}}},
					"discount_amounts",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedDiscountAmounts, ok := valueDiscountAmounts.(types.List); ok {
						state.DiscountAmounts = typedDiscountAmounts
					}
				}
			}
		}
	}
	{
		if rawValueEffectiveAt, rawOk := plainValueAtPath(raw, "effective_at"); rawOk {
			if valueEffectiveAt, err := flattenPlainValue(rawValueEffectiveAt, types.Int64Type, "effective_at", "raw response"); err != nil {
				return err
			} else {
				if typedEffectiveAt, ok := valueEffectiveAt.(types.Int64); ok {
					state.EffectiveAt = typedEffectiveAt
				}
			}
		} else if !hasRaw {
			if responseValueEffectiveAt, ok := plainFromResponseField(obj, "EffectiveAt"); ok {
				if valueEffectiveAt, err := flattenPlainValue(responseValueEffectiveAt, types.Int64Type, "effective_at", "response struct"); err != nil {
					return err
				} else {
					if typedEffectiveAt, ok := valueEffectiveAt.(types.Int64); ok {
						state.EffectiveAt = typedEffectiveAt
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
		if rawValueMemo, rawOk := plainValueAtPath(raw, "memo"); rawOk {
			if valueMemo, err := flattenPlainValue(rawValueMemo, types.StringType, "memo", "raw response"); err != nil {
				return err
			} else {
				if typedMemo, ok := valueMemo.(types.String); ok {
					state.Memo = typedMemo
				}
			}
		} else if !hasRaw {
			if responseValueMemo, ok := plainFromResponseField(obj, "Memo"); ok {
				if valueMemo, err := flattenPlainValue(responseValueMemo, types.StringType, "memo", "response struct"); err != nil {
					return err
				} else {
					if typedMemo, ok := valueMemo.(types.String); ok {
						state.Memo = typedMemo
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
		if rawValueOutOfBandAmount, rawOk := plainValueAtPath(raw, "out_of_band_amount"); rawOk {
			if valueOutOfBandAmount, err := flattenPlainValue(rawValueOutOfBandAmount, types.Int64Type, "out_of_band_amount", "raw response"); err != nil {
				return err
			} else {
				if typedOutOfBandAmount, ok := valueOutOfBandAmount.(types.Int64); ok {
					state.OutOfBandAmount = typedOutOfBandAmount
				}
			}
		} else if !hasRaw {
			if responseValueOutOfBandAmount, ok := plainFromResponseField(obj, "OutOfBandAmount"); ok {
				if valueOutOfBandAmount, err := flattenPlainValue(responseValueOutOfBandAmount, types.Int64Type, "out_of_band_amount", "response struct"); err != nil {
					return err
				} else {
					if typedOutOfBandAmount, ok := valueOutOfBandAmount.(types.Int64); ok {
						state.OutOfBandAmount = typedOutOfBandAmount
					}
				}
			}
		}
	}
	{
		if rawValuePostPaymentAmount, rawOk := plainValueAtPath(raw, "post_payment_amount"); rawOk {
			if valuePostPaymentAmount, err := flattenPlainValue(rawValuePostPaymentAmount, types.Int64Type, "post_payment_amount", "raw response"); err != nil {
				return err
			} else {
				if typedPostPaymentAmount, ok := valuePostPaymentAmount.(types.Int64); ok {
					state.PostPaymentAmount = typedPostPaymentAmount
				}
			}
		} else if !hasRaw {
			if responseValuePostPaymentAmount, ok := plainFromResponseField(obj, "PostPaymentAmount"); ok {
				if valuePostPaymentAmount, err := flattenPlainValue(responseValuePostPaymentAmount, types.Int64Type, "post_payment_amount", "response struct"); err != nil {
					return err
				} else {
					if typedPostPaymentAmount, ok := valuePostPaymentAmount.(types.Int64); ok {
						state.PostPaymentAmount = typedPostPaymentAmount
					}
				}
			}
		}
	}
	{
		if rawValuePrePaymentAmount, rawOk := plainValueAtPath(raw, "pre_payment_amount"); rawOk {
			if valuePrePaymentAmount, err := flattenPlainValue(rawValuePrePaymentAmount, types.Int64Type, "pre_payment_amount", "raw response"); err != nil {
				return err
			} else {
				if typedPrePaymentAmount, ok := valuePrePaymentAmount.(types.Int64); ok {
					state.PrePaymentAmount = typedPrePaymentAmount
				}
			}
		} else if !hasRaw {
			if responseValuePrePaymentAmount, ok := plainFromResponseField(obj, "PrePaymentAmount"); ok {
				if valuePrePaymentAmount, err := flattenPlainValue(responseValuePrePaymentAmount, types.Int64Type, "pre_payment_amount", "response struct"); err != nil {
					return err
				} else {
					if typedPrePaymentAmount, ok := valuePrePaymentAmount.(types.Int64); ok {
						state.PrePaymentAmount = typedPrePaymentAmount
					}
				}
			}
		}
	}
	{
		if rawValuePretaxCreditAmounts, rawOk := plainValueAtPath(raw, "pretax_credit_amounts"); rawOk {
			if valuePretaxCreditAmounts, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValuePretaxCreditAmounts, attrValueToPlain(state.PretaxCreditAmounts)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "credit_balance_transaction": types.StringType, "discount": types.StringType, "type": types.StringType}}}, "pretax_credit_amounts", "raw response"); err != nil {
				return err
			} else {
				if typedPretaxCreditAmounts, ok := valuePretaxCreditAmounts.(types.List); ok {
					state.PretaxCreditAmounts = typedPretaxCreditAmounts
				}
			}
		} else if !hasRaw {
			if responseValuePretaxCreditAmounts, ok := plainFromResponseField(obj, "PretaxCreditAmounts"); ok {
				if valuePretaxCreditAmounts, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValuePretaxCreditAmounts, attrValueToPlain(state.PretaxCreditAmounts)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "credit_balance_transaction": types.StringType, "discount": types.StringType, "type": types.StringType}}},
					"pretax_credit_amounts",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedPretaxCreditAmounts, ok := valuePretaxCreditAmounts.(types.List); ok {
						state.PretaxCreditAmounts = typedPretaxCreditAmounts
					}
				}
			}
		}
	}
	{
		if rawValueReason, rawOk := plainValueAtPath(raw, "reason"); rawOk {
			if valueReason, err := flattenPlainValue(rawValueReason, types.StringType, "reason", "raw response"); err != nil {
				return err
			} else {
				if typedReason, ok := valueReason.(types.String); ok {
					state.Reason = typedReason
				}
			}
		} else if !hasRaw {
			if responseValueReason, ok := plainFromResponseField(obj, "Reason"); ok {
				if valueReason, err := flattenPlainValue(responseValueReason, types.StringType, "reason", "response struct"); err != nil {
					return err
				} else {
					if typedReason, ok := valueReason.(types.String); ok {
						state.Reason = typedReason
					}
				}
			}
		}
	}
	{
		if rawValueRefunds, rawOk := plainValueAtPath(raw, "refunds"); rawOk {
			if valueRefunds, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueRefunds, attrValueToPlain(state.Refunds)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount_refunded": types.Int64Type, "payment_record_refund": types.ObjectType{AttrTypes: map[string]attr.Type{"payment_record": types.StringType, "refund_group": types.StringType}}, "refund": types.StringType, "type": types.StringType}}}, "refunds", "raw response"); err != nil {
				return err
			} else {
				if typedRefunds, ok := valueRefunds.(types.List); ok {
					state.Refunds = typedRefunds
				}
			}
		} else if !hasRaw {
			if responseValueRefunds, ok := plainFromResponseField(obj, "Refunds"); ok {
				if valueRefunds, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueRefunds, attrValueToPlain(state.Refunds)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount_refunded": types.Int64Type, "payment_record_refund": types.ObjectType{AttrTypes: map[string]attr.Type{"payment_record": types.StringType, "refund_group": types.StringType}}, "refund": types.StringType, "type": types.StringType}}},
					"refunds",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedRefunds, ok := valueRefunds.(types.List); ok {
						state.Refunds = typedRefunds
					}
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
		if rawValueSubtotal, rawOk := plainValueAtPath(raw, "subtotal"); rawOk {
			if valueSubtotal, err := flattenPlainValue(rawValueSubtotal, types.Int64Type, "subtotal", "raw response"); err != nil {
				return err
			} else {
				if typedSubtotal, ok := valueSubtotal.(types.Int64); ok {
					state.Subtotal = typedSubtotal
				}
			}
		} else if !hasRaw {
			if responseValueSubtotal, ok := plainFromResponseField(obj, "Subtotal"); ok {
				if valueSubtotal, err := flattenPlainValue(responseValueSubtotal, types.Int64Type, "subtotal", "response struct"); err != nil {
					return err
				} else {
					if typedSubtotal, ok := valueSubtotal.(types.Int64); ok {
						state.Subtotal = typedSubtotal
					}
				}
			}
		}
	}
	{
		if rawValueSubtotalExcludingTax, rawOk := plainValueAtPath(raw, "subtotal_excluding_tax"); rawOk {
			if valueSubtotalExcludingTax, err := flattenPlainValue(rawValueSubtotalExcludingTax, types.Int64Type, "subtotal_excluding_tax", "raw response"); err != nil {
				return err
			} else {
				if typedSubtotalExcludingTax, ok := valueSubtotalExcludingTax.(types.Int64); ok {
					state.SubtotalExcludingTax = typedSubtotalExcludingTax
				}
			}
		} else if !hasRaw {
			if responseValueSubtotalExcludingTax, ok := plainFromResponseField(obj, "SubtotalExcludingTax"); ok {
				if valueSubtotalExcludingTax, err := flattenPlainValue(responseValueSubtotalExcludingTax, types.Int64Type, "subtotal_excluding_tax", "response struct"); err != nil {
					return err
				} else {
					if typedSubtotalExcludingTax, ok := valueSubtotalExcludingTax.(types.Int64); ok {
						state.SubtotalExcludingTax = typedSubtotalExcludingTax
					}
				}
			}
		}
	}
	{
		if rawValueTotal, rawOk := plainValueAtPath(raw, "total"); rawOk {
			if valueTotal, err := flattenPlainValue(rawValueTotal, types.Int64Type, "total", "raw response"); err != nil {
				return err
			} else {
				if typedTotal, ok := valueTotal.(types.Int64); ok {
					state.Total = typedTotal
				}
			}
		} else if !hasRaw {
			if responseValueTotal, ok := plainFromResponseField(obj, "Total"); ok {
				if valueTotal, err := flattenPlainValue(responseValueTotal, types.Int64Type, "total", "response struct"); err != nil {
					return err
				} else {
					if typedTotal, ok := valueTotal.(types.Int64); ok {
						state.Total = typedTotal
					}
				}
			}
		}
	}
	{
		if rawValueTotalExcludingTax, rawOk := plainValueAtPath(raw, "total_excluding_tax"); rawOk {
			if valueTotalExcludingTax, err := flattenPlainValue(rawValueTotalExcludingTax, types.Int64Type, "total_excluding_tax", "raw response"); err != nil {
				return err
			} else {
				if typedTotalExcludingTax, ok := valueTotalExcludingTax.(types.Int64); ok {
					state.TotalExcludingTax = typedTotalExcludingTax
				}
			}
		} else if !hasRaw {
			if responseValueTotalExcludingTax, ok := plainFromResponseField(obj, "TotalExcludingTax"); ok {
				if valueTotalExcludingTax, err := flattenPlainValue(responseValueTotalExcludingTax, types.Int64Type, "total_excluding_tax", "response struct"); err != nil {
					return err
				} else {
					if typedTotalExcludingTax, ok := valueTotalExcludingTax.(types.Int64); ok {
						state.TotalExcludingTax = typedTotalExcludingTax
					}
				}
			}
		}
	}
	{
		if rawValueTotalTaxes, rawOk := plainValueAtPath(raw, "total_taxes"); rawOk {
			if valueTotalTaxes, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueTotalTaxes, attrValueToPlain(state.TotalTaxes)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "tax_behavior": types.StringType, "tax_rate_details": types.ObjectType{AttrTypes: map[string]attr.Type{"tax_rate": types.StringType}}, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type, "type": types.StringType}}}, "total_taxes", "raw response"); err != nil {
				return err
			} else {
				if typedTotalTaxes, ok := valueTotalTaxes.(types.List); ok {
					state.TotalTaxes = typedTotalTaxes
				}
			}
		} else if !hasRaw {
			if responseValueTotalTaxes, ok := plainFromResponseField(obj, "TotalTaxes"); ok {
				if valueTotalTaxes, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueTotalTaxes, attrValueToPlain(state.TotalTaxes)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "tax_behavior": types.StringType, "tax_rate_details": types.ObjectType{AttrTypes: map[string]attr.Type{"tax_rate": types.StringType}}, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type, "type": types.StringType}}},
					"total_taxes",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedTotalTaxes, ok := valueTotalTaxes.(types.List); ok {
						state.TotalTaxes = typedTotalTaxes
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
		if rawValueVoidedAt, rawOk := plainValueAtPath(raw, "voided_at"); rawOk {
			if valueVoidedAt, err := flattenPlainValue(rawValueVoidedAt, types.Int64Type, "voided_at", "raw response"); err != nil {
				return err
			} else {
				if typedVoidedAt, ok := valueVoidedAt.(types.Int64); ok {
					state.VoidedAt = typedVoidedAt
				}
			}
		} else if !hasRaw {
			if responseValueVoidedAt, ok := plainFromResponseField(obj, "VoidedAt"); ok {
				if valueVoidedAt, err := flattenPlainValue(responseValueVoidedAt, types.Int64Type, "voided_at", "response struct"); err != nil {
					return err
				} else {
					if typedVoidedAt, ok := valueVoidedAt.(types.Int64); ok {
						state.VoidedAt = typedVoidedAt
					}
				}
			}
		}
	}
	return nil
}
