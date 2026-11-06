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

var _ resource.Resource = &InvoiceResource{}

var _ resource.ResourceWithConfigure = &InvoiceResource{}

var _ resource.ResourceWithImportState = &InvoiceResource{}

func NewInvoiceResource() resource.Resource {
	return &InvoiceResource{}
}

type InvoiceResource struct {
	client *stripe.Client
}

type InvoiceResourceModel struct {
	Object                       types.String `tfsdk:"object"`
	AccountCountry               types.String `tfsdk:"account_country"`
	AccountName                  types.String `tfsdk:"account_name"`
	AccountTaxIDs                types.List   `tfsdk:"account_tax_ids"`
	AmountDue                    types.Int64  `tfsdk:"amount_due"`
	AmountOverpaid               types.Int64  `tfsdk:"amount_overpaid"`
	AmountPaid                   types.Int64  `tfsdk:"amount_paid"`
	AmountPaidOffStripe          types.Int64  `tfsdk:"amount_paid_off_stripe"`
	AmountRemaining              types.Int64  `tfsdk:"amount_remaining"`
	AmountShipping               types.Int64  `tfsdk:"amount_shipping"`
	Application                  types.String `tfsdk:"application"`
	AttemptCount                 types.Int64  `tfsdk:"attempt_count"`
	Attempted                    types.Bool   `tfsdk:"attempted"`
	AutoAdvance                  types.Bool   `tfsdk:"auto_advance"`
	AutomaticTax                 types.Object `tfsdk:"automatic_tax"`
	AutomaticallyFinalizesAt     types.Int64  `tfsdk:"automatically_finalizes_at"`
	BillingReason                types.String `tfsdk:"billing_reason"`
	CollectionMethod             types.String `tfsdk:"collection_method"`
	ConfirmationSecret           types.Object `tfsdk:"confirmation_secret"`
	Created                      types.Int64  `tfsdk:"created"`
	Currency                     types.String `tfsdk:"currency"`
	CustomFields                 types.List   `tfsdk:"custom_fields"`
	Customer                     types.String `tfsdk:"customer"`
	CustomerAccount              types.String `tfsdk:"customer_account"`
	CustomerAddress              types.Object `tfsdk:"customer_address"`
	CustomerEmail                types.String `tfsdk:"customer_email"`
	CustomerName                 types.String `tfsdk:"customer_name"`
	CustomerPhone                types.String `tfsdk:"customer_phone"`
	CustomerShipping             types.Object `tfsdk:"customer_shipping"`
	CustomerTaxIDs               types.List   `tfsdk:"customer_tax_ids"`
	DefaultPaymentMethod         types.String `tfsdk:"default_payment_method"`
	DefaultSource                types.String `tfsdk:"default_source"`
	DefaultTaxRates              types.List   `tfsdk:"default_tax_rates"`
	Description                  types.String `tfsdk:"description"`
	Discounts                    types.List   `tfsdk:"discounts"`
	DueDate                      types.Int64  `tfsdk:"due_date"`
	EffectiveAt                  types.Int64  `tfsdk:"effective_at"`
	EndingBalance                types.Int64  `tfsdk:"ending_balance"`
	Footer                       types.String `tfsdk:"footer"`
	FromInvoice                  types.Object `tfsdk:"from_invoice"`
	HostedInvoiceURL             types.String `tfsdk:"hosted_invoice_url"`
	ID                           types.String `tfsdk:"id"`
	InvoicePDF                   types.String `tfsdk:"invoice_pdf"`
	Issuer                       types.Object `tfsdk:"issuer"`
	LastFinalizationError        types.Object `tfsdk:"last_finalization_error"`
	LatestRevision               types.String `tfsdk:"latest_revision"`
	Livemode                     types.Bool   `tfsdk:"livemode"`
	Metadata                     types.Map    `tfsdk:"metadata"`
	NextPaymentAttempt           types.Int64  `tfsdk:"next_payment_attempt"`
	Number                       types.String `tfsdk:"number"`
	OnBehalfOf                   types.String `tfsdk:"on_behalf_of"`
	Parent                       types.Object `tfsdk:"parent"`
	PaymentSettings              types.Object `tfsdk:"payment_settings"`
	PeriodEnd                    types.Int64  `tfsdk:"period_end"`
	PeriodStart                  types.Int64  `tfsdk:"period_start"`
	PostPaymentCreditNotesAmount types.Int64  `tfsdk:"post_payment_credit_notes_amount"`
	PrePaymentCreditNotesAmount  types.Int64  `tfsdk:"pre_payment_credit_notes_amount"`
	ReceiptNumber                types.String `tfsdk:"receipt_number"`
	Rendering                    types.Object `tfsdk:"rendering"`
	ShippingCost                 types.Object `tfsdk:"shipping_cost"`
	ShippingDetails              types.Object `tfsdk:"shipping_details"`
	StartingBalance              types.Int64  `tfsdk:"starting_balance"`
	StatementDescriptor          types.String `tfsdk:"statement_descriptor"`
	Status                       types.String `tfsdk:"status"`
	StatusTransitions            types.Object `tfsdk:"status_transitions"`
	Subscription                 types.String `tfsdk:"subscription"`
	Subtotal                     types.Int64  `tfsdk:"subtotal"`
	SubtotalExcludingTax         types.Int64  `tfsdk:"subtotal_excluding_tax"`
	TestClock                    types.String `tfsdk:"test_clock"`
	ThresholdReason              types.Object `tfsdk:"threshold_reason"`
	Total                        types.Int64  `tfsdk:"total"`
	TotalDiscountAmounts         types.List   `tfsdk:"total_discount_amounts"`
	TotalExcludingTax            types.Int64  `tfsdk:"total_excluding_tax"`
	TotalPretaxCreditAmounts     types.List   `tfsdk:"total_pretax_credit_amounts"`
	TotalTaxes                   types.List   `tfsdk:"total_taxes"`
	WebhooksDeliveredAt          types.Int64  `tfsdk:"webhooks_delivered_at"`
	ApplicationFeeAmount         types.Int64  `tfsdk:"application_fee_amount"`
	DaysUntilDue                 types.Int64  `tfsdk:"days_until_due"`
	PendingInvoiceItemsBehavior  types.String `tfsdk:"pending_invoice_items_behavior"`
	TransferData                 types.Object `tfsdk:"transfer_data"`
}

func (r *InvoiceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *InvoiceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_invoice"
}

func (r *InvoiceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Invoices are statements of amounts owed by a customer, and are either\ngenerated one-off, or generated periodically from a subscription.\n\nThey contain [invoice items](https://api.stripe.com#invoiceitems), and proration adjustments\nthat may be caused by subscription upgrades/downgrades (if necessary).\n\nIf your invoice is configured to be billed through automatic charges,\nStripe automatically finalizes your invoice and attempts payment. Note\nthat finalizing the invoice,\n[when automatic](https://docs.stripe.com/invoicing/integration/automatic-advancement-collection), does\nnot happen immediately as the invoice is created. Stripe waits\nuntil one hour after the last webhook was successfully sent (or the last\nwebhook timed out after failing). If you (and the platforms you may have\nconnected to) have no webhooks configured, Stripe waits one hour after\ncreation to finalize the invoice.\n\nIf your invoice is configured to be billed by sending an email, then based on your\n[email settings](https://dashboard.stripe.com/account/billing/automatic),\nStripe will email the invoice to your customer and await payment. These\nemails can contain a link to a hosted page to pay the invoice.\n\nStripe applies any customer credit on the account before determining the\namount due for the invoice (i.e., the amount that will be actually\ncharged). If the amount due for the invoice is less than Stripe's [minimum allowed charge\nper currency](/docs/currencies#minimum-and-maximum-charge-amounts), the\ninvoice is automatically marked paid, and we add the amount due to the\ncustomer's credit balance which is applied to the next invoice.\n\nMore details on the customer's credit balance are\n[here](https://docs.stripe.com/billing/customer/balance).\n\nRelated guide: [Send invoices to customers](https://docs.stripe.com/billing/invoices/sending)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("invoice")},
			},
			"account_country": schema.StringAttribute{
				Computed:      true,
				Description:   "The country of the business associated with this invoice, most often the business creating the invoice.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_name": schema.StringAttribute{
				Computed:      true,
				Description:   "The public name of the business associated with this invoice, most often the business creating the invoice.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_tax_ids": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account tax IDs associated with the invoice. Only editable when the invoice is a draft.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"amount_due": schema.Int64Attribute{
				Computed:      true,
				Description:   "Final amount due at this time for this invoice. If the invoice's total is smaller than the minimum charge amount, for example, or if there is account credit that can be applied to the invoice, the `amount_due` may be 0. If there is a positive `starting_balance` for the invoice (the customer owes money), the `amount_due` will also take that into account. The charge that gets generated for the invoice will be for the amount specified in `amount_due`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"amount_overpaid": schema.Int64Attribute{
				Computed:      true,
				Description:   "Amount that was overpaid on the invoice. The amount overpaid is credited to the customer's credit balance.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"amount_paid": schema.Int64Attribute{
				Computed:      true,
				Description:   "The amount, in cents (or local equivalent), that was paid.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"amount_paid_off_stripe": schema.Int64Attribute{
				Computed:      true,
				Description:   "Amount, in cents (or local equivalent), that was paid on the invoice outside of Stripe.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"amount_remaining": schema.Int64Attribute{
				Computed:      true,
				Description:   "The difference between amount_due and amount_paid, in cents (or local equivalent).",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"amount_shipping": schema.Int64Attribute{
				Computed:      true,
				Description:   "This is the sum of all the shipping amounts.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the Connect Application that created the invoice.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"attempt_count": schema.Int64Attribute{
				Computed:      true,
				Description:   "Number of payment attempts made for this invoice, from the perspective of the payment retry schedule. Any payment attempt counts as the first attempt, and subsequently only automatic retries increment the attempt count. In other words, manual payment attempts after the first attempt do not affect the retry schedule. If a failure is returned with a non-retryable return code, the invoice can no longer be retried unless a new payment method is obtained. Retries will continue to be scheduled, and attempt_count will continue to increment, but retries will only be executed if a new payment method is obtained.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"attempted": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether an attempt has been made to pay the invoice. An invoice is not attempted until 1 hour after the `invoice.created` webhook, for example, so you might not want to display that invoice as unpaid to your users.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"auto_advance": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Controls whether Stripe performs [automatic collection](https://docs.stripe.com/invoicing/integration/automatic-advancement-collection) of the invoice. If `false`, the invoice's state doesn't automatically advance without an explicit action.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"automatic_tax": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"disabled_reason": schema.StringAttribute{
						Computed:      true,
						Description:   "If Stripe disabled automatic tax, this enum describes why.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("finalization_requires_location_inputs", "finalization_system_error")},
					},
					"enabled": schema.BoolAttribute{
						Required:    true,
						Description: "Whether Stripe automatically computes tax on this invoice. Note that incompatible invoice items (invoice items with manually specified [tax rates](https://docs.stripe.com/api/tax_rates), negative amounts, or `tax_behavior=unspecified`) cannot be added to automatic tax invoices.",
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
						Description:   "The status of the most recent automated tax calculation for this invoice.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("complete", "failed", "requires_location_inputs")},
					},
				},
			},
			"automatically_finalizes_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The time when this invoice is currently scheduled to be automatically finalized. The field will be `null` if the invoice is not scheduled to finalize in the future. If the invoice is not in the draft state, this field will always be `null` - see `finalized_at` for the time when an already-finalized invoice was finalized.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"billing_reason": schema.StringAttribute{
				Computed:      true,
				Description:   "Indicates the reason why the invoice was created.\n\n* `manual`: Unrelated to a subscription, for example, created via the invoice editor.\n* `subscription`: No longer in use. Applies to subscriptions from before May 2018 where no distinction was made between updates, cycles, and thresholds.\n* `subscription_create`: A new subscription was created.\n* `subscription_cycle`: A subscription advanced into a new period.\n* `subscription_threshold`: A subscription reached a billing threshold.\n* `subscription_update`: A subscription was updated.\n* `upcoming`: Reserved for upcoming invoices created through the Create Preview Invoice API or when an `invoice.upcoming` event is generated for an upcoming invoice on a subscription.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("automatic_pending_invoice_item_invoice", "manual", "quote_accept", "subscription", "subscription_create", "subscription_cycle", "subscription_threshold", "subscription_update", "upcoming")},
			},
			"collection_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("charge_automatically", "send_invoice")},
			},
			"confirmation_secret": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The confirmation secret associated with this invoice. Currently, this contains the client_secret of the PaymentIntent that Stripe creates during invoice finalization.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"client_secret": schema.StringAttribute{
						Computed:      true,
						Description:   "The client_secret of the payment that Stripe creates for the invoice after finalization.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Sensitive:     true,
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "The type of client_secret. Currently this is always payment_intent, referencing the default payment_intent that Stripe creates during invoice finalization",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
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
			"custom_fields": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Custom fields displayed on the invoice.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:    true,
							Description: "The name of the custom field.",
						},
						"value": schema.StringAttribute{
							Required:    true,
							Description: "The value of the custom field.",
						},
					},
				},
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ID of the customer to bill.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"customer_account": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ID of the account representing the customer to bill.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"customer_address": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The customer's address. Until the invoice is finalized, this field will equal `customer.address`. Once the invoice is finalized, this field will no longer be updated.",
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
			"customer_email": schema.StringAttribute{
				Computed:      true,
				Description:   "The customer's email. Until the invoice is finalized, this field will equal `customer.email`. Once the invoice is finalized, this field will no longer be updated.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_name": schema.StringAttribute{
				Computed:      true,
				Description:   "The customer's name. Until the invoice is finalized, this field will equal `customer.name`. Once the invoice is finalized, this field will no longer be updated.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_phone": schema.StringAttribute{
				Computed:      true,
				Description:   "The customer's phone number. Until the invoice is finalized, this field will equal `customer.phone`. Once the invoice is finalized, this field will no longer be updated.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_shipping": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The customer's shipping information. Until the invoice is finalized, this field will equal `customer.shipping`. Once the invoice is finalized, this field will no longer be updated.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"address": schema.SingleNestedAttribute{
						Computed: true,

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
					"carrier": schema.StringAttribute{
						Computed:      true,
						Description:   "The delivery service that shipped a physical product, such as Fedex, UPS, USPS, etc.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"name": schema.StringAttribute{
						Computed:      true,
						Description:   "Recipient name.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"phone": schema.StringAttribute{
						Computed:      true,
						Description:   "Recipient phone (including extension).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"tracking_number": schema.StringAttribute{
						Computed:      true,
						Description:   "The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"customer_tax_ids": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "The customer's tax IDs. Until the invoice is finalized, this field will contain the same tax IDs as `customer.tax_ids`. Once the invoice is finalized, this field will no longer be updated.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Computed:      true,
							Description:   "The type of the tax ID, one of `ad_nrt`, `ar_cuit`, `eu_vat`, `bo_tin`, `br_cnpj`, `br_cpf`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `eu_oss_vat`, `hr_oib`, `pe_ruc`, `ro_tin`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, `vn_tin`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `no_voec`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `pl_nip`, `it_cf`, `fo_vat`, `gi_tin`, `py_ruc`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `li_vat`, `lk_vat`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, `al_tin`, `bh_vat`, `kz_bin`, `ng_tin`, `om_vat`, `de_stn`, `ch_uid`, `tz_vat`, `uz_vat`, `uz_tin`, `md_vat`, `ma_vat`, `by_tin`, `ao_tin`, `bs_tin`, `bb_tin`, `cd_nif`, `mr_nif`, `me_pib`, `zw_tin`, `ba_tin`, `gn_nif`, `mk_vat`, `sr_fin`, `sn_ninea`, `am_tin`, `np_pan`, `tj_tin`, `ug_tin`, `zm_tin`, `kh_tin`, `aw_tin`, `az_tin`, `bd_bin`, `bj_ifu`, `et_tin`, `kg_tin`, `la_tin`, `cm_niu`, `cv_nif`, `bf_ifu`, or `unknown`",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("ad_nrt", "ae_trn", "al_tin", "am_tin", "ao_tin", "ar_cuit", "au_abn", "au_arn", "aw_tin", "az_tin", "ba_tin", "bb_tin", "bd_bin", "bf_ifu", "bg_uic", "bh_vat", "bj_ifu", "bo_tin", "br_cnpj", "br_cpf", "bs_tin", "by_tin", "ca_bn", "ca_gst_hst", "ca_pst_bc", "ca_pst_mb", "ca_pst_sk", "ca_qst", "cd_nif", "ch_uid", "ch_vat", "cl_tin", "cm_niu", "cn_tin", "co_nit", "cr_tin", "cv_nif", "de_stn", "do_rcn", "ec_ruc", "eg_tin", "es_cif", "et_tin", "eu_oss_vat", "eu_vat", "fo_vat", "gb_vat", "ge_vat", "gi_tin", "gn_nif", "hk_br", "hr_oib", "hu_tin", "id_npwp", "il_vat", "in_gst", "is_vat", "it_cf", "jp_cn", "jp_rn", "jp_trn", "ke_pin", "kg_tin", "kh_tin", "kr_brn", "kz_bin", "la_tin", "li_uid", "li_vat", "lk_vat", "ma_vat", "md_vat", "me_pib", "mk_vat", "mr_nif", "mx_rfc", "my_frp", "my_itn", "my_sst", "ng_tin", "no_vat", "no_voec", "np_pan", "nz_gst", "om_vat", "pe_ruc", "ph_tin", "pl_nip", "py_ruc", "ro_tin", "rs_pib", "ru_inn", "ru_kpp", "sa_vat", "sg_gst", "sg_uen", "si_tin", "sn_ninea", "sr_fin", "sv_nit", "th_vat", "tj_tin", "tr_tin", "tw_vat", "tz_vat", "ua_vat", "ug_tin", "unknown", "us_ein", "uy_ruc", "uz_tin", "uz_vat", "ve_rif", "vn_tin", "za_vat", "zm_tin", "zw_tin")},
						},
						"value": schema.StringAttribute{
							Computed:      true,
							Description:   "The value of the tax ID.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
					},
				},
			},
			"default_payment_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription's default payment method, if any, or to the default payment method in the customer's invoice settings.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_source": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the default payment source for the invoice. It must belong to the customer associated with the invoice and be in a chargeable state. If not set, defaults to the subscription's default source, if any, or to the customer's default source.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_tax_rates": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The tax rates applied to this invoice, if any.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the object. Often useful for displaying to users. Referenced as 'memo' in the Dashboard.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"discounts": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The discounts applied to the invoice. Line item discounts are applied before invoice discounts. Use `expand[]=discounts` to expand each discount.",
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
			"due_date": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The date on which payment for this invoice is due. This value will be `null` for invoices where `collection_method=charge_automatically`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"effective_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The date when this invoice is in effect. Same as `finalized_at` unless overwritten. When defined, this value replaces the system-generated 'Date of issue' printed on the invoice PDF and receipt.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"ending_balance": schema.Int64Attribute{
				Computed:      true,
				Description:   "Ending customer balance after the invoice is finalized. Invoices are finalized approximately an hour after successful webhook delivery or when payment collection is attempted for the invoice. If the invoice has not been finalized yet, this will be null.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"footer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Footer displayed on the invoice.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"from_invoice": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Details of the invoice that was cloned. See the [revision documentation](https://docs.stripe.com/invoicing/invoice-revisions) for more details.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Required:      true,
						Description:   "The relation between this invoice and the cloned invoice",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"invoice": schema.StringAttribute{
						Required:      true,
						Description:   "The invoice that was cloned.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"hosted_invoice_url": schema.StringAttribute{
				Computed:      true,
				Description:   "The URL for the hosted invoice page, which allows customers to view and pay an invoice. If the invoice has not been finalized yet, this will be null.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object. For preview invoices created using the [create preview](https://stripe.com/docs/api/invoices/create_preview) endpoint, this id will be prefixed with `upcoming_in`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice_pdf": schema.StringAttribute{
				Computed:      true,
				Description:   "The link to download the PDF for the invoice. If the invoice has not been finalized yet, this will be null.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"last_finalization_error": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The error encountered during the previous attempt to finalize the invoice. This field is cleared when the invoice is successfully finalized.",
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
			"latest_revision": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the most recent non-draft revision of this invoice",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"next_payment_attempt": schema.Int64Attribute{
				Computed:      true,
				Description:   "The time at which payment will next be attempted. This value will be `null` for invoices where `collection_method=send_invoice`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"number": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A unique, identifying string that appears on emails sent to the customer for this invoice. This starts with the customer's unique invoice_prefix if it is specified.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"on_behalf_of": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account (if any) for which the funds of the invoice payment are intended. If set, the invoice will be presented with the branding and support information of the specified account. See the [Invoices with Connect](https://docs.stripe.com/billing/invoices/connect) documentation for details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"parent": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The parent that generated this invoice",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"quote_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details about the quote that generated this invoice",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"quote": schema.StringAttribute{
								Computed:      true,
								Description:   "The quote that generated this invoice",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"subscription_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details about the subscription that generated this invoice",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"metadata": schema.MapAttribute{
								Computed:      true,
								Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) defined as subscription metadata when an invoice is created. Becomes an immutable snapshot of the subscription metadata at the time of invoice finalization.\n *Note: This attribute is populated only for invoices created on or after June 29, 2023.*",
								PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"subscription": schema.StringAttribute{
								Computed:      true,
								Description:   "The subscription that generated this invoice",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"subscription_proration_date": schema.Int64Attribute{
								Computed:      true,
								Description:   "Only set for upcoming invoices that preview prorations. The time used to calculate prorations.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "The type of parent that generated this invoice",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("quote_details", "subscription_details")},
					},
				},
			},
			"payment_settings": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"default_mandate": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "ID of the mandate to be used for this invoice. It must correspond to the payment method used to pay the invoice, including the invoice's default_payment_method or default_source, if set.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"payment_method_options": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Payment-method-specific configuration to provide to the invoice’s PaymentIntent.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"acss_debit": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice’s PaymentIntent.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"mandate_options": schema.SingleNestedAttribute{
										Optional: true,
										Computed: true,

										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
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
							"bancontact": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "If paying by `bancontact`, this sub-hash contains details about the Bancontact payment method options to pass to the invoice’s PaymentIntent.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"preferred_language": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Preferred language of the Bancontact authorization page that the customer is redirected to.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("de", "en", "fr", "nl")},
									},
								},
							},
							"card": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "If paying by `card`, this sub-hash contains details about the Card payment method options to pass to the invoice’s PaymentIntent.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"installments": schema.SingleNestedAttribute{
										Optional: true,
										Computed: true,

										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Optional:      true,
												Computed:      true,
												Description:   "Whether Installments are enabled for this Invoice.",
												PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
											},
											"plan": schema.SingleNestedAttribute{
												Optional:    true,
												Description: "The selected installment plan to use for this invoice.",
												Attributes: map[string]schema.Attribute{
													"count": schema.Int64Attribute{
														Optional:    true,
														Description: "For `fixed_count` installment plans, this is required. It represents the number of installment payments your customer will make to their credit card.",
													},
													"interval": schema.StringAttribute{
														Optional:    true,
														Description: "For `fixed_count` installment plans, this is required. It represents the interval between installment payments your customer will make to their credit card.\nOne of `month`.",
													},
													"type": schema.StringAttribute{
														Required:    true,
														Description: "Type of installment plan, one of `fixed_count`, `bonus`, or `revolving`.",
													},
												},
											},
										},
									},
									"request_three_d_secure": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("any", "automatic", "challenge")},
									},
								},
							},
							"customer_balance": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "If paying by `customer_balance`, this sub-hash contains details about the Bank transfer payment method options to pass to the invoice’s PaymentIntent.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"bank_transfer": schema.SingleNestedAttribute{
										Optional: true,
										Computed: true,

										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"eu_bank_transfer": schema.SingleNestedAttribute{
												Optional: true,
												Computed: true,

												PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
												Attributes: map[string]schema.Attribute{
													"country": schema.StringAttribute{
														Required:    true,
														Description: "The desired country code of the bank account information. Permitted values include: `DE`, `FR`, `IE`, or `NL`.",
														Validators:  []validator.String{stringvalidator.OneOf("BE", "DE", "ES", "FR", "IE", "NL")},
													},
												},
											},
											"type": schema.StringAttribute{
												Optional:      true,
												Computed:      true,
												Description:   "The bank transfer type that can be used for funding. Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, `mx_bank_transfer`, or `us_bank_transfer`.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
										},
									},
									"funding_type": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The funding method type to be used when there are not enough funds in the customer balance. Permitted values include: `bank_transfer`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("bank_transfer")},
									},
								},
							},
							"payto": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "If paying by `payto`, this sub-hash contains details about the PayTo payment method options to pass to the invoice’s PaymentIntent.",
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
												Description:   "The maximum amount that can be collected in a single invoice. If you don't specify a maximum, then there is no limit.",
												PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
											},
											"amount_type": schema.StringAttribute{
												Computed:      true,
												Description:   "Only `maximum` is supported.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												Validators:    []validator.String{stringvalidator.OneOf("fixed", "maximum")},
											},
											"purpose": schema.StringAttribute{
												Optional:      true,
												Computed:      true,
												Description:   "The purpose for which payments are made. Has a default value based on your merchant category code.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												Validators:    []validator.String{stringvalidator.OneOf("dependant_support", "government", "loan", "mortgage", "other", "pension", "personal", "retail", "salary", "tax", "utility")},
											},
										},
									},
								},
							},
							"pix": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "If paying by `pix`, this sub-hash contains details about the Pix payment method options to pass to the invoice’s PaymentIntent.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"amount_includes_iof": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Determines if the amount includes the IOF tax.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("always", "never")},
									},
									"expires_after_seconds": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "The number of seconds (between 10 and 1209600) after which Pix payment will expire. Defaults to 86400 seconds.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
								},
							},
							"upi": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "If paying by `upi`, this sub-hash contains details about the UPI payment method options to pass to the invoice’s PaymentIntent.",
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
								},
							},
							"us_bank_account": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "If paying by `us_bank_account`, this sub-hash contains details about the ACH direct debit payment method options to pass to the invoice’s PaymentIntent.",
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
						},
					},
					"payment_method_types": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The list of payment method types (e.g. card) to provide to the invoice’s PaymentIntent. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice’s default payment method, the subscription’s default payment method, the customer’s default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
				},
			},
			"period_end": schema.Int64Attribute{
				Computed:      true,
				Description:   "The latest timestamp at which invoice items can be associated with this invoice. Use the [line item period](/api/invoices/line_item#invoice_line_item_object-period) to get the service period for each price.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"period_start": schema.Int64Attribute{
				Computed:      true,
				Description:   "The earliest timestamp at which invoice items can be associated with this invoice. Use the [line item period](/api/invoices/line_item#invoice_line_item_object-period) to get the service period for each price.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"post_payment_credit_notes_amount": schema.Int64Attribute{
				Computed:      true,
				Description:   "Total amount of all post-payment credit notes issued for this invoice.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"pre_payment_credit_notes_amount": schema.Int64Attribute{
				Computed:      true,
				Description:   "Total amount of all pre-payment credit notes issued for this invoice.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"receipt_number": schema.StringAttribute{
				Computed:      true,
				Description:   "This is the transaction number that appears on email receipts sent for this invoice.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"rendering": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The rendering-related settings that control how the invoice is displayed on customer-facing surfaces such as PDF and Hosted Invoice Page.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"amount_tax_display": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pdf": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Invoice pdf rendering options",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"page_size": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Page size of invoice pdf. Options include a4, letter, and auto. If set to auto, page size will be switched to a4 or letter based on customer locale.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("a4", "auto", "letter")},
							},
						},
					},
					"template": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "ID of the rendering template that the invoice is formatted by.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"template_version": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "Version of the rendering template that the invoice is using.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"shipping_cost": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The details of the cost of shipping, including the ShippingRate applied on the invoice.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
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
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
					"shipping_rate_data": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "Parameters to create a new ad-hoc shipping rate for this order.",
						Attributes: map[string]schema.Attribute{
							"delivery_estimate": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "The estimated range for how long shipping will take, meant to be displayable to the customer. This will appear on CheckoutSessions.",
								Attributes: map[string]schema.Attribute{
									"maximum": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "The upper bound of the estimated range. If empty, represents no upper bound i.e., infinite.",
										Attributes: map[string]schema.Attribute{
											"unit": schema.StringAttribute{
												Required:    true,
												Description: "A unit of time.",
											},
											"value": schema.Int64Attribute{
												Required:    true,
												Description: "Must be greater than 0.",
											},
										},
									},
									"minimum": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "The lower bound of the estimated range. If empty, represents no lower bound.",
										Attributes: map[string]schema.Attribute{
											"unit": schema.StringAttribute{
												Required:    true,
												Description: "A unit of time.",
											},
											"value": schema.Int64Attribute{
												Required:    true,
												Description: "Must be greater than 0.",
											},
										},
									},
								},
							},
							"display_name": schema.StringAttribute{
								Required:    true,
								Description: "The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.",
							},
							"fixed_amount": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Describes a fixed amount to charge for shipping. Must be present if type is `fixed_amount`.",
								Attributes: map[string]schema.Attribute{
									"amount": schema.Int64Attribute{
										Required:    true,
										Description: "A non-negative integer in cents representing how much to charge.",
									},
									"currency": schema.StringAttribute{
										Required:    true,
										Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
									},
									"currency_options": schema.MapNestedAttribute{
										Optional:    true,
										Description: "Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"amount": schema.Int64Attribute{
													Required:    true,
													Description: "A non-negative integer in cents representing how much to charge.",
												},
												"tax_behavior": schema.StringAttribute{
													Optional:    true,
													Description: "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.",
												},
											},
										},
									},
								},
							},
							"metadata": schema.MapAttribute{
								Optional:    true,
								Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
								ElementType: types.StringType,
							},
							"tax_behavior": schema.StringAttribute{
								Optional:    true,
								Description: "Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.",
							},
							"tax_code": schema.StringAttribute{
								Optional:    true,
								Description: "A [tax code](https://docs.stripe.com/tax/tax-categories) ID. The Shipping tax code is `txcd_92010001`.",
							},
							"type": schema.StringAttribute{
								Optional:    true,
								Description: "The type of calculation to use on the shipping rate.",
							},
						},
					},
				},
			},
			"shipping_details": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Shipping details for the invoice. The Invoice PDF will use the `shipping_details` value if it is set, otherwise the PDF will render the shipping address from the customer.",
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
						Computed:      true,
						Description:   "The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"starting_balance": schema.Int64Attribute{
				Computed:      true,
				Description:   "Starting customer balance before the invoice is finalized. If the invoice has not been finalized yet, this will be the current customer balance. For revision invoices, this also includes any customer balance that was applied to the original invoice.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"statement_descriptor": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Extra information about an invoice for the customer's credit card statement.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The status of the invoice, one of `draft`, `open`, `paid`, `uncollectible`, or `void`. [Learn more](https://docs.stripe.com/billing/invoices/workflow#workflow-overview)",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("draft", "open", "paid", "uncollectible", "void")},
			},
			"status_transitions": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"finalized_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time that the invoice draft was finalized.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"marked_uncollectible_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time that the invoice was marked uncollectible.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"paid_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time that the invoice was paid.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"voided_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The time that the invoice was voided.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"subscription": schema.StringAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"subtotal": schema.Int64Attribute{
				Computed:      true,
				Description:   "Total of all subscriptions, invoice items, and prorations on the invoice before any invoice level discount or exclusive tax is applied. Item discounts are already incorporated",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"subtotal_excluding_tax": schema.Int64Attribute{
				Computed:      true,
				Description:   "The integer amount in cents (or local equivalent) representing the subtotal of the invoice before any invoice level discount or tax is applied. Item discounts are already incorporated",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"test_clock": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the test clock this invoice belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"threshold_reason": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"amount_gte": schema.Int64Attribute{
						Computed:      true,
						Description:   "The total invoice amount threshold boundary if it triggered the threshold invoice.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"item_reasons": schema.ListNestedAttribute{
						Computed:      true,
						Description:   "Indicates which line items triggered a threshold invoice.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"line_item_ids": schema.ListAttribute{
									Computed:      true,
									Description:   "The IDs of the line items that triggered the threshold invoice.",
									PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
									ElementType:   types.StringType,
								},
								"usage_gte": schema.Int64Attribute{
									Computed:      true,
									Description:   "The quantity threshold boundary that applied to the given line item.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
							},
						},
					},
				},
			},
			"total": schema.Int64Attribute{
				Computed:      true,
				Description:   "Total after discounts and taxes.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"total_discount_amounts": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "The aggregate amounts calculated per discount across all line items.",
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
			"total_excluding_tax": schema.Int64Attribute{
				Computed:      true,
				Description:   "The integer amount in cents (or local equivalent) representing the total amount of the invoice including all discounts but excluding all tax.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"total_pretax_credit_amounts": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "Contains pretax credit amounts (ex: discount, credit grants, etc) that apply to this invoice. This is a combined list of total_pretax_credit_amounts across all invoice line items.",
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
			"total_taxes": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "The aggregate tax information of all line items.",
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
			"webhooks_delivered_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "Invoices are automatically paid or sent 1 hour after webhooks are delivered, or until all webhook delivery attempts have [been exhausted](https://docs.stripe.com/billing/webhooks#understand). This field tracks the time when webhooks for this invoice were successfully delivered. If the invoice had no webhooks to deliver, this will be set while the invoice is being created.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"application_fee_amount": schema.Int64Attribute{
				Optional:    true,
				Description: "A fee in cents (or local equivalent) that will be applied to the invoice and transferred to the application owner's Stripe account. The request must be made with an OAuth key or the Stripe-Account header in order to take an application fee. For more information, see the application fees [documentation](https://docs.stripe.com/billing/invoices/connect#collecting-fees).",
				WriteOnly:   true,
			},
			"days_until_due": schema.Int64Attribute{
				Optional:    true,
				Description: "The number of days from when the invoice is created until it is due. Valid only for invoices where `collection_method=send_invoice`.",
				WriteOnly:   true,
			},
			"pending_invoice_items_behavior": schema.StringAttribute{
				Optional:      true,
				Description:   "How to handle pending invoice items on invoice creation. Defaults to `exclude` if the parameter is omitted.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
				Validators:    []validator.String{stringvalidator.OneOf("exclude", "include")},
			},
			"transfer_data": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "If specified, the funds from the invoice will be transferred to the destination and the ID of the resulting transfer will be found on the invoice's charge.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"amount": schema.Int64Attribute{
						Optional:    true,
						Description: "The amount that will be transferred automatically when the invoice is paid. If no amount is set, the full amount is transferred.",
						WriteOnly:   true,
					},
					"destination": schema.StringAttribute{
						Required:    true,
						Description: "ID of an existing, connected Stripe account.",
						WriteOnly:   true,
					},
				},
			},
		},
	}
}

func (r *InvoiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InvoiceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config InvoiceResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"ApplicationFeeAmount"}, []string{"DaysUntilDue"}, []string{"PendingInvoiceItemsBehavior"}, []string{"TransferData"}, []string{"TransferData", "amount"}, []string{"TransferData", "destination"}})

	params, err := expandInvoiceCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Invoice create params", err.Error())
		return
	}

	obj, err := r.client.V1Invoices.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Invoice", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Invoices.B, r.client.V1Invoices.Key, stripe.FormatURLPath("/v1/invoices/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Invoice create raw response", err.Error())
		return
	}

	if err := flattenInvoice(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Invoice create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan", "count"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan", "interval"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan", "type"}, []string{"ShippingCost", "shipping_rate_data"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "maximum"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "maximum", "unit"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "maximum", "value"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "minimum"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "minimum", "unit"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "minimum", "value"}, []string{"ShippingCost", "shipping_rate_data", "display_name"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "amount"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency_options"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency_options", "*", "amount"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency_options", "*", "tax_behavior"}, []string{"ShippingCost", "shipping_rate_data", "metadata"}, []string{"ShippingCost", "shipping_rate_data", "tax_behavior"}, []string{"ShippingCost", "shipping_rate_data", "tax_code"}, []string{"ShippingCost", "shipping_rate_data", "type"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"ApplicationFeeAmount"}, []string{"DaysUntilDue"}, []string{"PendingInvoiceItemsBehavior"}, []string{"TransferData"}, []string{"TransferData", "amount"}, []string{"TransferData", "destination"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *InvoiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState InvoiceResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state InvoiceResourceModel
	state = priorState

	obj, err := r.client.V1Invoices.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Invoice", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Invoices.B, r.client.V1Invoices.Key, stripe.FormatURLPath("/v1/invoices/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Invoice raw response", err.Error())
		return
	}

	if err := flattenInvoice(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Invoice read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan", "count"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan", "interval"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan", "type"}, []string{"ShippingCost", "shipping_rate_data"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "maximum"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "maximum", "unit"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "maximum", "value"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "minimum"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "minimum", "unit"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "minimum", "value"}, []string{"ShippingCost", "shipping_rate_data", "display_name"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "amount"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency_options"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency_options", "*", "amount"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency_options", "*", "tax_behavior"}, []string{"ShippingCost", "shipping_rate_data", "metadata"}, []string{"ShippingCost", "shipping_rate_data", "tax_behavior"}, []string{"ShippingCost", "shipping_rate_data", "tax_code"}, []string{"ShippingCost", "shipping_rate_data", "type"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"ApplicationFeeAmount"}, []string{"DaysUntilDue"}, []string{"PendingInvoiceItemsBehavior"}, []string{"TransferData"}, []string{"TransferData", "amount"}, []string{"TransferData", "destination"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *InvoiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InvoiceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config InvoiceResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"ApplicationFeeAmount"}, []string{"DaysUntilDue"}, []string{"PendingInvoiceItemsBehavior"}, []string{"TransferData"}, []string{"TransferData", "amount"}, []string{"TransferData", "destination"}})

	var state InvoiceResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"PendingInvoiceItemsBehavior"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"PendingInvoiceItemsBehavior"}})

	params, err := expandInvoiceUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Invoice update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Invoice update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Invoices.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Invoice", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Invoices.B, r.client.V1Invoices.Key, stripe.FormatURLPath("/v1/invoices/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Invoice update raw response", err.Error())
		return
	}

	if err := flattenInvoice(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Invoice update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan", "count"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan", "interval"}, []string{"PaymentSettings", "payment_method_options", "card", "installments", "plan", "type"}, []string{"ShippingCost", "shipping_rate_data"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "maximum"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "maximum", "unit"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "maximum", "value"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "minimum"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "minimum", "unit"}, []string{"ShippingCost", "shipping_rate_data", "delivery_estimate", "minimum", "value"}, []string{"ShippingCost", "shipping_rate_data", "display_name"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "amount"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency_options"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency_options", "*", "amount"}, []string{"ShippingCost", "shipping_rate_data", "fixed_amount", "currency_options", "*", "tax_behavior"}, []string{"ShippingCost", "shipping_rate_data", "metadata"}, []string{"ShippingCost", "shipping_rate_data", "tax_behavior"}, []string{"ShippingCost", "shipping_rate_data", "tax_code"}, []string{"ShippingCost", "shipping_rate_data", "type"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"ApplicationFeeAmount"}, []string{"DaysUntilDue"}, []string{"PendingInvoiceItemsBehavior"}, []string{"TransferData"}, []string{"TransferData", "amount"}, []string{"TransferData", "destination"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *InvoiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InvoiceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1Invoices.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting Invoice", err.Error())
		return
	}
}

func (r *InvoiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandInvoiceCreate(plan InvoiceResourceModel) (*stripe.InvoiceCreateParams, error) {
	params := &stripe.InvoiceCreateParams{}

	if !plan.AccountTaxIDs.IsNull() && !plan.AccountTaxIDs.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AccountTaxIDs", plan.AccountTaxIDs) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "account_tax_ids", params)
		}
	}
	if !plan.AutoAdvance.IsNull() && !plan.AutoAdvance.IsUnknown() {
		params.AutoAdvance = stripe.Bool(plan.AutoAdvance.ValueBool())
	}
	if !plan.AutomaticTax.IsNull() && !plan.AutomaticTax.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AutomaticTax", plan.AutomaticTax) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "automatic_tax", params)
		}
	}
	if !plan.AutomaticallyFinalizesAt.IsNull() && !plan.AutomaticallyFinalizesAt.IsUnknown() {
		params.AutomaticallyFinalizesAt = stripe.Int64(plan.AutomaticallyFinalizesAt.ValueInt64())
	}
	if !plan.CollectionMethod.IsNull() && !plan.CollectionMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CollectionMethod", "CollectionMethod", plan.CollectionMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "collection_method", params)
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
	if !plan.DefaultPaymentMethod.IsNull() && !plan.DefaultPaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DefaultPaymentMethodID", "DefaultPaymentMethod", plan.DefaultPaymentMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "default_payment_method", params)
		}
	}
	if !plan.DefaultSource.IsNull() && !plan.DefaultSource.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DefaultSourceID", "DefaultSource", plan.DefaultSource.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "default_source", params)
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
	if !plan.DueDate.IsNull() && !plan.DueDate.IsUnknown() {
		params.DueDate = stripe.Int64(plan.DueDate.ValueInt64())
	}
	if !plan.EffectiveAt.IsNull() && !plan.EffectiveAt.IsUnknown() {
		params.EffectiveAt = stripe.Int64(plan.EffectiveAt.ValueInt64())
	}
	if !plan.Footer.IsNull() && !plan.Footer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Footer", "Footer", plan.Footer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "footer", params)
		}
	}
	if !plan.FromInvoice.IsNull() && !plan.FromInvoice.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FromInvoice", plan.FromInvoice) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "from_invoice", params)
		}
	}
	if !plan.Issuer.IsNull() && !plan.Issuer.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Issuer", plan.Issuer) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "issuer", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Number.IsNull() && !plan.Number.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Number", "Number", plan.Number.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "number", params)
		}
	}
	if !plan.OnBehalfOf.IsNull() && !plan.OnBehalfOf.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "OnBehalfOfID", "OnBehalfOf", plan.OnBehalfOf.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "on_behalf_of", params)
		}
	}
	if !plan.PaymentSettings.IsNull() && !plan.PaymentSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentSettings", plan.PaymentSettings) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_settings", params)
		}
	}
	if !plan.Rendering.IsNull() && !plan.Rendering.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Rendering", plan.Rendering) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "rendering", params)
		}
	}
	if !plan.ShippingCost.IsNull() && !plan.ShippingCost.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingCost", plan.ShippingCost) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping_cost", params)
		}
	}
	if !plan.ShippingDetails.IsNull() && !plan.ShippingDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingDetails", plan.ShippingDetails) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping_details", params)
		}
	}
	if !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "statement_descriptor", params)
		}
	}
	if !plan.Subscription.IsNull() && !plan.Subscription.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SubscriptionID", "Subscription", plan.Subscription.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "subscription", params)
		}
	}
	if !plan.ApplicationFeeAmount.IsNull() && !plan.ApplicationFeeAmount.IsUnknown() {
		params.ApplicationFeeAmount = stripe.Int64(plan.ApplicationFeeAmount.ValueInt64())
	}
	if !plan.DaysUntilDue.IsNull() && !plan.DaysUntilDue.IsUnknown() {
		params.DaysUntilDue = stripe.Int64(plan.DaysUntilDue.ValueInt64())
	}
	if !plan.PendingInvoiceItemsBehavior.IsNull() && !plan.PendingInvoiceItemsBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PendingInvoiceItemsBehavior", "PendingInvoiceItemsBehavior", plan.PendingInvoiceItemsBehavior.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "pending_invoice_items_behavior", params)
		}
	}
	if !plan.TransferData.IsNull() && !plan.TransferData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TransferData", plan.TransferData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "transfer_data", params)
		}
	}

	return params, nil
}

func expandInvoiceUpdate(plan InvoiceResourceModel, state InvoiceResourceModel) (*stripe.InvoiceUpdateParams, error) {
	params := &stripe.InvoiceUpdateParams{}

	if !plan.AccountTaxIDs.Equal(state.AccountTaxIDs) && !plan.AccountTaxIDs.IsNull() && !plan.AccountTaxIDs.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AccountTaxIDs", plan.AccountTaxIDs) {
			if !plan.AccountTaxIDs.Equal(state.AccountTaxIDs) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "account_tax_ids", params)
			}
		}
	}
	if !plan.AutoAdvance.Equal(state.AutoAdvance) && !plan.AutoAdvance.IsNull() && !plan.AutoAdvance.IsUnknown() {
		params.AutoAdvance = stripe.Bool(plan.AutoAdvance.ValueBool())
	}
	if !plan.AutomaticTax.Equal(state.AutomaticTax) && !plan.AutomaticTax.IsNull() && !plan.AutomaticTax.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AutomaticTax", plan.AutomaticTax) {
			if !plan.AutomaticTax.Equal(state.AutomaticTax) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "automatic_tax", params)
			}
		}
	}
	if !plan.AutomaticallyFinalizesAt.Equal(state.AutomaticallyFinalizesAt) && !plan.AutomaticallyFinalizesAt.IsNull() && !plan.AutomaticallyFinalizesAt.IsUnknown() {
		params.AutomaticallyFinalizesAt = stripe.Int64(plan.AutomaticallyFinalizesAt.ValueInt64())
	}
	if !plan.CollectionMethod.Equal(state.CollectionMethod) && !plan.CollectionMethod.IsNull() && !plan.CollectionMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CollectionMethod", "CollectionMethod", plan.CollectionMethod.ValueString()) {
			if !plan.CollectionMethod.Equal(state.CollectionMethod) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "collection_method", params)
			}
		}
	}
	if !plan.CustomFields.Equal(state.CustomFields) && !plan.CustomFields.IsNull() && !plan.CustomFields.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomFields", plan.CustomFields) {
			if !plan.CustomFields.Equal(state.CustomFields) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "custom_fields", params)
			}
		}
	}
	if !plan.DefaultPaymentMethod.Equal(state.DefaultPaymentMethod) && !plan.DefaultPaymentMethod.IsNull() && !plan.DefaultPaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DefaultPaymentMethodID", "DefaultPaymentMethod", plan.DefaultPaymentMethod.ValueString()) {
			if !plan.DefaultPaymentMethod.Equal(state.DefaultPaymentMethod) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "default_payment_method", params)
			}
		}
	}
	if !plan.DefaultSource.Equal(state.DefaultSource) && !plan.DefaultSource.IsNull() && !plan.DefaultSource.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DefaultSourceID", "DefaultSource", plan.DefaultSource.ValueString()) {
			if !plan.DefaultSource.Equal(state.DefaultSource) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "default_source", params)
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
	if !plan.DueDate.Equal(state.DueDate) && !plan.DueDate.IsNull() && !plan.DueDate.IsUnknown() {
		params.DueDate = stripe.Int64(plan.DueDate.ValueInt64())
	}
	if !plan.EffectiveAt.Equal(state.EffectiveAt) && !plan.EffectiveAt.IsNull() && !plan.EffectiveAt.IsUnknown() {
		params.EffectiveAt = stripe.Int64(plan.EffectiveAt.ValueInt64())
	}
	if !plan.Footer.Equal(state.Footer) && !plan.Footer.IsNull() && !plan.Footer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Footer", "Footer", plan.Footer.ValueString()) {
			if !plan.Footer.Equal(state.Footer) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "footer", params)
			}
		}
	}
	if !plan.Issuer.Equal(state.Issuer) && !plan.Issuer.IsNull() && !plan.Issuer.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Issuer", plan.Issuer) {
			if !plan.Issuer.Equal(state.Issuer) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "issuer", params)
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
	if !plan.Number.Equal(state.Number) && !plan.Number.IsNull() && !plan.Number.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Number", "Number", plan.Number.ValueString()) {
			if !plan.Number.Equal(state.Number) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "number", params)
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
	if !plan.PaymentSettings.Equal(state.PaymentSettings) && !plan.PaymentSettings.IsNull() && !plan.PaymentSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentSettings", plan.PaymentSettings) {
			if !plan.PaymentSettings.Equal(state.PaymentSettings) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_settings", params)
			}
		}
	}
	if !plan.Rendering.Equal(state.Rendering) && !plan.Rendering.IsNull() && !plan.Rendering.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Rendering", plan.Rendering) {
			if !plan.Rendering.Equal(state.Rendering) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "rendering", params)
			}
		}
	}
	if !plan.ShippingCost.Equal(state.ShippingCost) && !plan.ShippingCost.IsNull() && !plan.ShippingCost.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingCost", plan.ShippingCost) {
			if !plan.ShippingCost.Equal(state.ShippingCost) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "shipping_cost", params)
			}
		}
	}
	if !plan.ShippingDetails.Equal(state.ShippingDetails) && !plan.ShippingDetails.IsNull() && !plan.ShippingDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingDetails", plan.ShippingDetails) {
			if !plan.ShippingDetails.Equal(state.ShippingDetails) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "shipping_details", params)
			}
		}
	}
	if !plan.StatementDescriptor.Equal(state.StatementDescriptor) && !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			if !plan.StatementDescriptor.Equal(state.StatementDescriptor) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "statement_descriptor", params)
			}
		}
	}
	if !plan.ApplicationFeeAmount.Equal(state.ApplicationFeeAmount) && !plan.ApplicationFeeAmount.IsNull() && !plan.ApplicationFeeAmount.IsUnknown() {
		params.ApplicationFeeAmount = stripe.Int64(plan.ApplicationFeeAmount.ValueInt64())
	}
	if !plan.DaysUntilDue.Equal(state.DaysUntilDue) && !plan.DaysUntilDue.IsNull() && !plan.DaysUntilDue.IsUnknown() {
		params.DaysUntilDue = stripe.Int64(plan.DaysUntilDue.ValueInt64())
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

func flattenInvoice(obj *stripe.Invoice, state *InvoiceResourceModel) error {
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
		if rawValueAccountCountry, rawOk := plainValueAtPath(raw, "account_country"); rawOk {
			if valueAccountCountry, err := flattenPlainValue(rawValueAccountCountry, types.StringType, "account_country", "raw response"); err != nil {
				return err
			} else {
				if typedAccountCountry, ok := valueAccountCountry.(types.String); ok {
					state.AccountCountry = typedAccountCountry
				}
			}
		} else if !hasRaw {
			if responseValueAccountCountry, ok := plainFromResponseField(obj, "AccountCountry"); ok {
				if valueAccountCountry, err := flattenPlainValue(responseValueAccountCountry, types.StringType, "account_country", "response struct"); err != nil {
					return err
				} else {
					if typedAccountCountry, ok := valueAccountCountry.(types.String); ok {
						state.AccountCountry = typedAccountCountry
					}
				}
			}
		}
	}
	{
		if rawValueAccountName, rawOk := plainValueAtPath(raw, "account_name"); rawOk {
			if valueAccountName, err := flattenPlainValue(rawValueAccountName, types.StringType, "account_name", "raw response"); err != nil {
				return err
			} else {
				if typedAccountName, ok := valueAccountName.(types.String); ok {
					state.AccountName = typedAccountName
				}
			}
		} else if !hasRaw {
			if responseValueAccountName, ok := plainFromResponseField(obj, "AccountName"); ok {
				if valueAccountName, err := flattenPlainValue(responseValueAccountName, types.StringType, "account_name", "response struct"); err != nil {
					return err
				} else {
					if typedAccountName, ok := valueAccountName.(types.String); ok {
						state.AccountName = typedAccountName
					}
				}
			}
		}
	}
	{
		if rawValueAccountTaxIDs, rawOk := plainValueAtPath(raw, "account_tax_ids"); rawOk {
			if valueAccountTaxIDs, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueAccountTaxIDs, attrValueToPlain(state.AccountTaxIDs)), types.ListType{ElemType: types.StringType}, "account_tax_ids", "raw response"); err != nil {
				return err
			} else {
				if typedAccountTaxIDs, ok := valueAccountTaxIDs.(types.List); ok {
					state.AccountTaxIDs = typedAccountTaxIDs
				}
			}
		} else if !hasRaw {
			if responseValueAccountTaxIDs, ok := plainFromResponseField(obj, "AccountTaxIDs"); ok {
				if valueAccountTaxIDs, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueAccountTaxIDs, attrValueToPlain(state.AccountTaxIDs)),
					types.ListType{ElemType: types.StringType},
					"account_tax_ids",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedAccountTaxIDs, ok := valueAccountTaxIDs.(types.List); ok {
						state.AccountTaxIDs = typedAccountTaxIDs
					}
				}
			}
		}
	}
	{
		if rawValueAmountDue, rawOk := plainValueAtPath(raw, "amount_due"); rawOk {
			if valueAmountDue, err := flattenPlainValue(rawValueAmountDue, types.Int64Type, "amount_due", "raw response"); err != nil {
				return err
			} else {
				if typedAmountDue, ok := valueAmountDue.(types.Int64); ok {
					state.AmountDue = typedAmountDue
				}
			}
		} else if !hasRaw {
			if responseValueAmountDue, ok := plainFromResponseField(obj, "AmountDue"); ok {
				if valueAmountDue, err := flattenPlainValue(responseValueAmountDue, types.Int64Type, "amount_due", "response struct"); err != nil {
					return err
				} else {
					if typedAmountDue, ok := valueAmountDue.(types.Int64); ok {
						state.AmountDue = typedAmountDue
					}
				}
			}
		}
	}
	{
		if rawValueAmountOverpaid, rawOk := plainValueAtPath(raw, "amount_overpaid"); rawOk {
			if valueAmountOverpaid, err := flattenPlainValue(rawValueAmountOverpaid, types.Int64Type, "amount_overpaid", "raw response"); err != nil {
				return err
			} else {
				if typedAmountOverpaid, ok := valueAmountOverpaid.(types.Int64); ok {
					state.AmountOverpaid = typedAmountOverpaid
				}
			}
		} else if !hasRaw {
			if responseValueAmountOverpaid, ok := plainFromResponseField(obj, "AmountOverpaid"); ok {
				if valueAmountOverpaid, err := flattenPlainValue(responseValueAmountOverpaid, types.Int64Type, "amount_overpaid", "response struct"); err != nil {
					return err
				} else {
					if typedAmountOverpaid, ok := valueAmountOverpaid.(types.Int64); ok {
						state.AmountOverpaid = typedAmountOverpaid
					}
				}
			}
		}
	}
	{
		if rawValueAmountPaid, rawOk := plainValueAtPath(raw, "amount_paid"); rawOk {
			if valueAmountPaid, err := flattenPlainValue(rawValueAmountPaid, types.Int64Type, "amount_paid", "raw response"); err != nil {
				return err
			} else {
				if typedAmountPaid, ok := valueAmountPaid.(types.Int64); ok {
					state.AmountPaid = typedAmountPaid
				}
			}
		} else if !hasRaw {
			if responseValueAmountPaid, ok := plainFromResponseField(obj, "AmountPaid"); ok {
				if valueAmountPaid, err := flattenPlainValue(responseValueAmountPaid, types.Int64Type, "amount_paid", "response struct"); err != nil {
					return err
				} else {
					if typedAmountPaid, ok := valueAmountPaid.(types.Int64); ok {
						state.AmountPaid = typedAmountPaid
					}
				}
			}
		}
	}
	{
		if rawValueAmountPaidOffStripe, rawOk := plainValueAtPath(raw, "amount_paid_off_stripe"); rawOk {
			if valueAmountPaidOffStripe, err := flattenPlainValue(rawValueAmountPaidOffStripe, types.Int64Type, "amount_paid_off_stripe", "raw response"); err != nil {
				return err
			} else {
				if typedAmountPaidOffStripe, ok := valueAmountPaidOffStripe.(types.Int64); ok {
					state.AmountPaidOffStripe = typedAmountPaidOffStripe
				}
			}
		} else if !hasRaw {
			if responseValueAmountPaidOffStripe, ok := plainFromResponseField(obj, "AmountPaidOffStripe"); ok {
				if valueAmountPaidOffStripe, err := flattenPlainValue(responseValueAmountPaidOffStripe, types.Int64Type, "amount_paid_off_stripe", "response struct"); err != nil {
					return err
				} else {
					if typedAmountPaidOffStripe, ok := valueAmountPaidOffStripe.(types.Int64); ok {
						state.AmountPaidOffStripe = typedAmountPaidOffStripe
					}
				}
			}
		}
	}
	{
		if rawValueAmountRemaining, rawOk := plainValueAtPath(raw, "amount_remaining"); rawOk {
			if valueAmountRemaining, err := flattenPlainValue(rawValueAmountRemaining, types.Int64Type, "amount_remaining", "raw response"); err != nil {
				return err
			} else {
				if typedAmountRemaining, ok := valueAmountRemaining.(types.Int64); ok {
					state.AmountRemaining = typedAmountRemaining
				}
			}
		} else if !hasRaw {
			if responseValueAmountRemaining, ok := plainFromResponseField(obj, "AmountRemaining"); ok {
				if valueAmountRemaining, err := flattenPlainValue(responseValueAmountRemaining, types.Int64Type, "amount_remaining", "response struct"); err != nil {
					return err
				} else {
					if typedAmountRemaining, ok := valueAmountRemaining.(types.Int64); ok {
						state.AmountRemaining = typedAmountRemaining
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
		if rawValueAttemptCount, rawOk := plainValueAtPath(raw, "attempt_count"); rawOk {
			if valueAttemptCount, err := flattenPlainValue(rawValueAttemptCount, types.Int64Type, "attempt_count", "raw response"); err != nil {
				return err
			} else {
				if typedAttemptCount, ok := valueAttemptCount.(types.Int64); ok {
					state.AttemptCount = typedAttemptCount
				}
			}
		} else if !hasRaw {
			if responseValueAttemptCount, ok := plainFromResponseField(obj, "AttemptCount"); ok {
				if valueAttemptCount, err := flattenPlainValue(responseValueAttemptCount, types.Int64Type, "attempt_count", "response struct"); err != nil {
					return err
				} else {
					if typedAttemptCount, ok := valueAttemptCount.(types.Int64); ok {
						state.AttemptCount = typedAttemptCount
					}
				}
			}
		}
	}
	{
		if rawValueAttempted, rawOk := plainValueAtPath(raw, "attempted"); rawOk {
			if valueAttempted, err := flattenPlainValue(rawValueAttempted, types.BoolType, "attempted", "raw response"); err != nil {
				return err
			} else {
				if typedAttempted, ok := valueAttempted.(types.Bool); ok {
					state.Attempted = typedAttempted
				}
			}
		} else if !hasRaw {
			if responseValueAttempted, ok := plainFromResponseField(obj, "Attempted"); ok {
				if valueAttempted, err := flattenPlainValue(responseValueAttempted, types.BoolType, "attempted", "response struct"); err != nil {
					return err
				} else {
					if typedAttempted, ok := valueAttempted.(types.Bool); ok {
						state.Attempted = typedAttempted
					}
				}
			}
		}
	}
	{
		if rawValueAutoAdvance, rawOk := plainValueAtPath(raw, "auto_advance"); rawOk {
			if valueAutoAdvance, err := flattenPlainValue(rawValueAutoAdvance, types.BoolType, "auto_advance", "raw response"); err != nil {
				return err
			} else {
				if typedAutoAdvance, ok := valueAutoAdvance.(types.Bool); ok {
					state.AutoAdvance = typedAutoAdvance
				}
			}
		} else if !hasRaw {
			if responseValueAutoAdvance, ok := plainFromResponseField(obj, "AutoAdvance"); ok {
				if valueAutoAdvance, err := flattenPlainValue(responseValueAutoAdvance, types.BoolType, "auto_advance", "response struct"); err != nil {
					return err
				} else {
					if typedAutoAdvance, ok := valueAutoAdvance.(types.Bool); ok {
						state.AutoAdvance = typedAutoAdvance
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
				if valueAutomaticTax, err := flattenPlainValue(sourceAutomaticTax, types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "provider": types.StringType, "status": types.StringType}}, "automatic_tax", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "provider": types.StringType, "status": types.StringType}},
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
			if nullAutomaticTax, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "provider": types.StringType, "status": types.StringType}}); ok {
				if typedAutomaticTax, ok := nullAutomaticTax.(types.Object); ok {
					state.AutomaticTax = typedAutomaticTax
				}
			}
		}
	}
	{
		if rawValueAutomaticallyFinalizesAt, rawOk := plainValueAtPath(raw, "automatically_finalizes_at"); rawOk {
			if valueAutomaticallyFinalizesAt, err := flattenPlainValue(rawValueAutomaticallyFinalizesAt, types.Int64Type, "automatically_finalizes_at", "raw response"); err != nil {
				return err
			} else {
				if typedAutomaticallyFinalizesAt, ok := valueAutomaticallyFinalizesAt.(types.Int64); ok {
					state.AutomaticallyFinalizesAt = typedAutomaticallyFinalizesAt
				}
			}
		} else if !hasRaw {
			if responseValueAutomaticallyFinalizesAt, ok := plainFromResponseField(obj, "AutomaticallyFinalizesAt"); ok {
				if valueAutomaticallyFinalizesAt, err := flattenPlainValue(responseValueAutomaticallyFinalizesAt, types.Int64Type, "automatically_finalizes_at", "response struct"); err != nil {
					return err
				} else {
					if typedAutomaticallyFinalizesAt, ok := valueAutomaticallyFinalizesAt.(types.Int64); ok {
						state.AutomaticallyFinalizesAt = typedAutomaticallyFinalizesAt
					}
				}
			}
		}
	}
	{
		if rawValueBillingReason, rawOk := plainValueAtPath(raw, "billing_reason"); rawOk {
			if valueBillingReason, err := flattenPlainValue(rawValueBillingReason, types.StringType, "billing_reason", "raw response"); err != nil {
				return err
			} else {
				if typedBillingReason, ok := valueBillingReason.(types.String); ok {
					state.BillingReason = typedBillingReason
				}
			}
		} else if !hasRaw {
			if responseValueBillingReason, ok := plainFromResponseField(obj, "BillingReason"); ok {
				if valueBillingReason, err := flattenPlainValue(responseValueBillingReason, types.StringType, "billing_reason", "response struct"); err != nil {
					return err
				} else {
					if typedBillingReason, ok := valueBillingReason.(types.String); ok {
						state.BillingReason = typedBillingReason
					}
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
		assignedConfirmationSecret := false
		hadRawConfirmationSecret := false
		if rawValueConfirmationSecret, rawOk := plainValueAtPath(raw, "confirmation_secret"); rawOk {
			hadRawConfirmationSecret = true
			if rawValueConfirmationSecret != nil {
				sourceConfirmationSecret := applyConfiguredKeyedListShapes(rawValueConfirmationSecret, attrValueToPlain(state.ConfirmationSecret))
				if valueConfirmationSecret, err := flattenPlainValue(sourceConfirmationSecret, types.ObjectType{AttrTypes: map[string]attr.Type{"client_secret": types.StringType, "type": types.StringType}}, "confirmation_secret", "raw response"); err != nil {
					return err
				} else {
					if typedConfirmationSecret, ok := valueConfirmationSecret.(types.Object); ok {
						state.ConfirmationSecret = typedConfirmationSecret
						assignedConfirmationSecret = true
					}
				}
			}
		}
		if !assignedConfirmationSecret {
			if !hasRaw {
				if responseValueConfirmationSecret, ok := plainFromResponseField(obj, "ConfirmationSecret"); ok {
					sourceConfirmationSecret := applyConfiguredKeyedListShapes(responseValueConfirmationSecret, attrValueToPlain(state.ConfirmationSecret))
					if valueConfirmationSecret, err := flattenPlainValue(
						sourceConfirmationSecret,
						types.ObjectType{AttrTypes: map[string]attr.Type{"client_secret": types.StringType, "type": types.StringType}},
						"confirmation_secret",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedConfirmationSecret, ok := valueConfirmationSecret.(types.Object); ok {
							state.ConfirmationSecret = typedConfirmationSecret
							assignedConfirmationSecret = true
						}
					}
				}
			}
		}
		if !assignedConfirmationSecret && hadRawConfirmationSecret {
			if nullConfirmationSecret, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"client_secret": types.StringType, "type": types.StringType}}); ok {
				if typedConfirmationSecret, ok := nullConfirmationSecret.(types.Object); ok {
					state.ConfirmationSecret = typedConfirmationSecret
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
		if rawValueCustomFields, rawOk := plainValueAtPath(raw, "custom_fields"); rawOk {
			if valueCustomFields, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueCustomFields, attrValueToPlain(state.CustomFields)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "custom_fields", "raw response"); err != nil {
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
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}},
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
		assignedCustomerAddress := false
		hadRawCustomerAddress := false
		if rawValueCustomerAddress, rawOk := plainValueAtPath(raw, "customer_address"); rawOk {
			hadRawCustomerAddress = true
			if rawValueCustomerAddress != nil {
				sourceCustomerAddress := applyConfiguredKeyedListShapes(rawValueCustomerAddress, attrValueToPlain(state.CustomerAddress))
				if valueCustomerAddress, err := flattenPlainValue(sourceCustomerAddress, types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "customer_address", "raw response"); err != nil {
					return err
				} else {
					if typedCustomerAddress, ok := valueCustomerAddress.(types.Object); ok {
						state.CustomerAddress = typedCustomerAddress
						assignedCustomerAddress = true
					}
				}
			}
		}
		if !assignedCustomerAddress {
			if !hasRaw {
				if responseValueCustomerAddress, ok := plainFromResponseField(obj, "CustomerAddress"); ok {
					sourceCustomerAddress := applyConfiguredKeyedListShapes(responseValueCustomerAddress, attrValueToPlain(state.CustomerAddress))
					if valueCustomerAddress, err := flattenPlainValue(
						sourceCustomerAddress,
						types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}},
						"customer_address",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCustomerAddress, ok := valueCustomerAddress.(types.Object); ok {
							state.CustomerAddress = typedCustomerAddress
							assignedCustomerAddress = true
						}
					}
				}
			}
		}
		if !assignedCustomerAddress && hadRawCustomerAddress {
			if nullCustomerAddress, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}); ok {
				if typedCustomerAddress, ok := nullCustomerAddress.(types.Object); ok {
					state.CustomerAddress = typedCustomerAddress
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
		if rawValueCustomerName, rawOk := plainValueAtPath(raw, "customer_name"); rawOk {
			if valueCustomerName, err := flattenPlainValue(rawValueCustomerName, types.StringType, "customer_name", "raw response"); err != nil {
				return err
			} else {
				if typedCustomerName, ok := valueCustomerName.(types.String); ok {
					state.CustomerName = typedCustomerName
				}
			}
		} else if !hasRaw {
			if responseValueCustomerName, ok := plainFromResponseField(obj, "CustomerName"); ok {
				if valueCustomerName, err := flattenPlainValue(responseValueCustomerName, types.StringType, "customer_name", "response struct"); err != nil {
					return err
				} else {
					if typedCustomerName, ok := valueCustomerName.(types.String); ok {
						state.CustomerName = typedCustomerName
					}
				}
			}
		}
	}
	{
		if rawValueCustomerPhone, rawOk := plainValueAtPath(raw, "customer_phone"); rawOk {
			if valueCustomerPhone, err := flattenPlainValue(rawValueCustomerPhone, types.StringType, "customer_phone", "raw response"); err != nil {
				return err
			} else {
				if typedCustomerPhone, ok := valueCustomerPhone.(types.String); ok {
					state.CustomerPhone = typedCustomerPhone
				}
			}
		} else if !hasRaw {
			if responseValueCustomerPhone, ok := plainFromResponseField(obj, "CustomerPhone"); ok {
				if valueCustomerPhone, err := flattenPlainValue(responseValueCustomerPhone, types.StringType, "customer_phone", "response struct"); err != nil {
					return err
				} else {
					if typedCustomerPhone, ok := valueCustomerPhone.(types.String); ok {
						state.CustomerPhone = typedCustomerPhone
					}
				}
			}
		}
	}
	{
		assignedCustomerShipping := false
		hadRawCustomerShipping := false
		if rawValueCustomerShipping, rawOk := plainValueAtPath(raw, "customer_shipping"); rawOk {
			hadRawCustomerShipping = true
			if rawValueCustomerShipping != nil {
				sourceCustomerShipping := applyConfiguredKeyedListShapes(rawValueCustomerShipping, attrValueToPlain(state.CustomerShipping))
				if valueCustomerShipping, err := flattenPlainValue(sourceCustomerShipping, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}, "customer_shipping", "raw response"); err != nil {
					return err
				} else {
					if typedCustomerShipping, ok := valueCustomerShipping.(types.Object); ok {
						state.CustomerShipping = typedCustomerShipping
						assignedCustomerShipping = true
					}
				}
			}
		}
		if !assignedCustomerShipping {
			if !hasRaw {
				if responseValueCustomerShipping, ok := plainFromResponseField(obj, "CustomerShipping"); ok {
					sourceCustomerShipping := applyConfiguredKeyedListShapes(responseValueCustomerShipping, attrValueToPlain(state.CustomerShipping))
					if valueCustomerShipping, err := flattenPlainValue(
						sourceCustomerShipping,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}},
						"customer_shipping",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCustomerShipping, ok := valueCustomerShipping.(types.Object); ok {
							state.CustomerShipping = typedCustomerShipping
							assignedCustomerShipping = true
						}
					}
				}
			}
		}
		if !assignedCustomerShipping && hadRawCustomerShipping {
			if nullCustomerShipping, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}); ok {
				if typedCustomerShipping, ok := nullCustomerShipping.(types.Object); ok {
					state.CustomerShipping = typedCustomerShipping
				}
			}
		}
	}
	{
		if rawValueCustomerTaxIDs, rawOk := plainValueAtPath(raw, "customer_tax_ids"); rawOk {
			if valueCustomerTaxIDs, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueCustomerTaxIDs, attrValueToPlain(state.CustomerTaxIDs)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}}, "customer_tax_ids", "raw response"); err != nil {
				return err
			} else {
				if typedCustomerTaxIDs, ok := valueCustomerTaxIDs.(types.List); ok {
					state.CustomerTaxIDs = typedCustomerTaxIDs
				}
			}
		} else if !hasRaw {
			if responseValueCustomerTaxIDs, ok := plainFromResponseField(obj, "CustomerTaxIDs"); ok {
				if valueCustomerTaxIDs, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueCustomerTaxIDs, attrValueToPlain(state.CustomerTaxIDs)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}},
					"customer_tax_ids",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedCustomerTaxIDs, ok := valueCustomerTaxIDs.(types.List); ok {
						state.CustomerTaxIDs = typedCustomerTaxIDs
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueDefaultPaymentMethod, rawOk := plainValueAtPath(raw, "default_payment_method"); rawOk {
				if typedDefaultPaymentMethod, ok := plainToStringIDValue(rawValueDefaultPaymentMethod); ok {
					state.DefaultPaymentMethod = typedDefaultPaymentMethod
				}
			} else if !hasRaw {
				if responseValueDefaultPaymentMethod, ok := plainFromResponseField(obj, "DefaultPaymentMethod"); ok {
					if typedDefaultPaymentMethod, ok := plainToStringIDValue(responseValueDefaultPaymentMethod); ok {
						state.DefaultPaymentMethod = typedDefaultPaymentMethod
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueDefaultSource, rawOk := plainValueAtPath(raw, "default_source"); rawOk {
				if typedDefaultSource, ok := plainToStringIDValue(rawValueDefaultSource); ok {
					state.DefaultSource = typedDefaultSource
				}
			} else if !hasRaw {
				if responseValueDefaultSource, ok := plainFromResponseField(obj, "DefaultSource"); ok {
					if typedDefaultSource, ok := plainToStringIDValue(responseValueDefaultSource); ok {
						state.DefaultSource = typedDefaultSource
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
		if rawValueDueDate, rawOk := plainValueAtPath(raw, "due_date"); rawOk {
			if valueDueDate, err := flattenPlainValue(rawValueDueDate, types.Int64Type, "due_date", "raw response"); err != nil {
				return err
			} else {
				if typedDueDate, ok := valueDueDate.(types.Int64); ok {
					state.DueDate = typedDueDate
				}
			}
		} else if !hasRaw {
			if responseValueDueDate, ok := plainFromResponseField(obj, "DueDate"); ok {
				if valueDueDate, err := flattenPlainValue(responseValueDueDate, types.Int64Type, "due_date", "response struct"); err != nil {
					return err
				} else {
					if typedDueDate, ok := valueDueDate.(types.Int64); ok {
						state.DueDate = typedDueDate
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
		if rawValueEndingBalance, rawOk := plainValueAtPath(raw, "ending_balance"); rawOk {
			if valueEndingBalance, err := flattenPlainValue(rawValueEndingBalance, types.Int64Type, "ending_balance", "raw response"); err != nil {
				return err
			} else {
				if typedEndingBalance, ok := valueEndingBalance.(types.Int64); ok {
					state.EndingBalance = typedEndingBalance
				}
			}
		} else if !hasRaw {
			if responseValueEndingBalance, ok := plainFromResponseField(obj, "EndingBalance"); ok {
				if valueEndingBalance, err := flattenPlainValue(responseValueEndingBalance, types.Int64Type, "ending_balance", "response struct"); err != nil {
					return err
				} else {
					if typedEndingBalance, ok := valueEndingBalance.(types.Int64); ok {
						state.EndingBalance = typedEndingBalance
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
		assignedFromInvoice := false
		hadRawFromInvoice := false
		if rawValueFromInvoice, rawOk := plainValueAtPath(raw, "from_invoice"); rawOk {
			hadRawFromInvoice = true
			if rawValueFromInvoice != nil {
				sourceFromInvoice := applyConfiguredKeyedListShapes(rawValueFromInvoice, attrValueToPlain(state.FromInvoice))
				if valueFromInvoice, err := flattenPlainValue(sourceFromInvoice, types.ObjectType{AttrTypes: map[string]attr.Type{"action": types.StringType, "invoice": types.StringType}}, "from_invoice", "raw response"); err != nil {
					return err
				} else {
					if typedFromInvoice, ok := valueFromInvoice.(types.Object); ok {
						state.FromInvoice = typedFromInvoice
						assignedFromInvoice = true
					}
				}
			}
		}
		if !assignedFromInvoice {
			if !hasRaw {
				if responseValueFromInvoice, ok := plainFromResponseField(obj, "FromInvoice"); ok {
					sourceFromInvoice := applyConfiguredKeyedListShapes(responseValueFromInvoice, attrValueToPlain(state.FromInvoice))
					if valueFromInvoice, err := flattenPlainValue(
						sourceFromInvoice,
						types.ObjectType{AttrTypes: map[string]attr.Type{"action": types.StringType, "invoice": types.StringType}},
						"from_invoice",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFromInvoice, ok := valueFromInvoice.(types.Object); ok {
							state.FromInvoice = typedFromInvoice
							assignedFromInvoice = true
						}
					}
				}
			}
		}
		if !assignedFromInvoice && hadRawFromInvoice {
			if nullFromInvoice, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"action": types.StringType, "invoice": types.StringType}}); ok {
				if typedFromInvoice, ok := nullFromInvoice.(types.Object); ok {
					state.FromInvoice = typedFromInvoice
				}
			}
		}
	}
	{
		if rawValueHostedInvoiceURL, rawOk := plainValueAtPath(raw, "hosted_invoice_url"); rawOk {
			if valueHostedInvoiceURL, err := flattenPlainValue(rawValueHostedInvoiceURL, types.StringType, "hosted_invoice_url", "raw response"); err != nil {
				return err
			} else {
				if typedHostedInvoiceURL, ok := valueHostedInvoiceURL.(types.String); ok {
					state.HostedInvoiceURL = typedHostedInvoiceURL
				}
			}
		} else if !hasRaw {
			if responseValueHostedInvoiceURL, ok := plainFromResponseField(obj, "HostedInvoiceURL"); ok {
				if valueHostedInvoiceURL, err := flattenPlainValue(responseValueHostedInvoiceURL, types.StringType, "hosted_invoice_url", "response struct"); err != nil {
					return err
				} else {
					if typedHostedInvoiceURL, ok := valueHostedInvoiceURL.(types.String); ok {
						state.HostedInvoiceURL = typedHostedInvoiceURL
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
		if rawValueInvoicePDF, rawOk := plainValueAtPath(raw, "invoice_pdf"); rawOk {
			if valueInvoicePDF, err := flattenPlainValue(rawValueInvoicePDF, types.StringType, "invoice_pdf", "raw response"); err != nil {
				return err
			} else {
				if typedInvoicePDF, ok := valueInvoicePDF.(types.String); ok {
					state.InvoicePDF = typedInvoicePDF
				}
			}
		} else if !hasRaw {
			if responseValueInvoicePDF, ok := plainFromResponseField(obj, "InvoicePDF"); ok {
				if valueInvoicePDF, err := flattenPlainValue(responseValueInvoicePDF, types.StringType, "invoice_pdf", "response struct"); err != nil {
					return err
				} else {
					if typedInvoicePDF, ok := valueInvoicePDF.(types.String); ok {
						state.InvoicePDF = typedInvoicePDF
					}
				}
			}
		}
	}
	{
		assignedIssuer := false
		hadRawIssuer := false
		if rawValueIssuer, rawOk := plainValueAtPath(raw, "issuer"); rawOk {
			hadRawIssuer = true
			if rawValueIssuer != nil {
				sourceIssuer := applyConfiguredKeyedListShapes(rawValueIssuer, attrValueToPlain(state.Issuer))
				if valueIssuer, err := flattenPlainValue(sourceIssuer, types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "issuer", "raw response"); err != nil {
					return err
				} else {
					if typedIssuer, ok := valueIssuer.(types.Object); ok {
						state.Issuer = typedIssuer
						assignedIssuer = true
					}
				}
			}
		}
		if !assignedIssuer {
			if !hasRaw {
				if responseValueIssuer, ok := plainFromResponseField(obj, "Issuer"); ok {
					sourceIssuer := applyConfiguredKeyedListShapes(responseValueIssuer, attrValueToPlain(state.Issuer))
					if valueIssuer, err := flattenPlainValue(
						sourceIssuer,
						types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}},
						"issuer",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedIssuer, ok := valueIssuer.(types.Object); ok {
							state.Issuer = typedIssuer
							assignedIssuer = true
						}
					}
				}
			}
		}
		if !assignedIssuer && hadRawIssuer {
			if nullIssuer, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}); ok {
				if typedIssuer, ok := nullIssuer.(types.Object); ok {
					state.Issuer = typedIssuer
				}
			}
		}
	}
	{
		assignedLastFinalizationError := false
		hadRawLastFinalizationError := false
		if rawValueLastFinalizationError, rawOk := plainValueAtPath(raw, "last_finalization_error"); rawOk {
			hadRawLastFinalizationError = true
			if rawValueLastFinalizationError != nil {
				sourceLastFinalizationError := applyConfiguredKeyedListShapes(rawValueLastFinalizationError, attrValueToPlain(state.LastFinalizationError))
				if valueLastFinalizationError, err := flattenPlainValue(sourceLastFinalizationError, types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}}, "last_finalization_error", "raw response"); err != nil {
					return err
				} else {
					if typedLastFinalizationError, ok := valueLastFinalizationError.(types.Object); ok {
						state.LastFinalizationError = typedLastFinalizationError
						assignedLastFinalizationError = true
					}
				}
			}
		}
		if !assignedLastFinalizationError {
			if !hasRaw {
				if responseValueLastFinalizationError, ok := plainFromResponseField(obj, "LastFinalizationError"); ok {
					sourceLastFinalizationError := applyConfiguredKeyedListShapes(responseValueLastFinalizationError, attrValueToPlain(state.LastFinalizationError))
					if valueLastFinalizationError, err := flattenPlainValue(
						sourceLastFinalizationError,
						types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}},
						"last_finalization_error",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLastFinalizationError, ok := valueLastFinalizationError.(types.Object); ok {
							state.LastFinalizationError = typedLastFinalizationError
							assignedLastFinalizationError = true
						}
					}
				}
			}
		}
		if !assignedLastFinalizationError && hadRawLastFinalizationError {
			if nullLastFinalizationError, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"advice_code": types.StringType, "charge": types.StringType, "code": types.StringType, "decline_code": types.StringType, "doc_url": types.StringType, "message": types.StringType, "network_advice_code": types.StringType, "network_decline_code": types.StringType, "param": types.StringType, "payment_intent": types.StringType, "payment_method": types.StringType, "payment_method_type": types.StringType, "request_log_url": types.StringType, "setup_intent": types.StringType, "source": types.StringType, "type": types.StringType}}); ok {
				if typedLastFinalizationError, ok := nullLastFinalizationError.(types.Object); ok {
					state.LastFinalizationError = typedLastFinalizationError
				}
			}
		}
	}
	{
		if true {
			if rawValueLatestRevision, rawOk := plainValueAtPath(raw, "latest_revision"); rawOk {
				if typedLatestRevision, ok := plainToStringIDValue(rawValueLatestRevision); ok {
					state.LatestRevision = typedLatestRevision
				}
			} else if !hasRaw {
				if responseValueLatestRevision, ok := plainFromResponseField(obj, "LatestRevision"); ok {
					if typedLatestRevision, ok := plainToStringIDValue(responseValueLatestRevision); ok {
						state.LatestRevision = typedLatestRevision
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
		if rawValueNextPaymentAttempt, rawOk := plainValueAtPath(raw, "next_payment_attempt"); rawOk {
			if valueNextPaymentAttempt, err := flattenPlainValue(rawValueNextPaymentAttempt, types.Int64Type, "next_payment_attempt", "raw response"); err != nil {
				return err
			} else {
				if typedNextPaymentAttempt, ok := valueNextPaymentAttempt.(types.Int64); ok {
					state.NextPaymentAttempt = typedNextPaymentAttempt
				}
			}
		} else if !hasRaw {
			if responseValueNextPaymentAttempt, ok := plainFromResponseField(obj, "NextPaymentAttempt"); ok {
				if valueNextPaymentAttempt, err := flattenPlainValue(responseValueNextPaymentAttempt, types.Int64Type, "next_payment_attempt", "response struct"); err != nil {
					return err
				} else {
					if typedNextPaymentAttempt, ok := valueNextPaymentAttempt.(types.Int64); ok {
						state.NextPaymentAttempt = typedNextPaymentAttempt
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
		assignedParent := false
		hadRawParent := false
		if rawValueParent, rawOk := plainValueAtPath(raw, "parent"); rawOk {
			hadRawParent = true
			if rawValueParent != nil {
				sourceParent := applyConfiguredKeyedListShapes(rawValueParent, attrValueToPlain(state.Parent))
				if valueParent, err := flattenPlainValue(sourceParent, types.ObjectType{AttrTypes: map[string]attr.Type{"quote_details": types.ObjectType{AttrTypes: map[string]attr.Type{"quote": types.StringType}}, "subscription_details": types.ObjectType{AttrTypes: map[string]attr.Type{"metadata": types.MapType{ElemType: types.StringType}, "subscription": types.StringType, "subscription_proration_date": types.Int64Type}}, "type": types.StringType}}, "parent", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"quote_details": types.ObjectType{AttrTypes: map[string]attr.Type{"quote": types.StringType}}, "subscription_details": types.ObjectType{AttrTypes: map[string]attr.Type{"metadata": types.MapType{ElemType: types.StringType}, "subscription": types.StringType, "subscription_proration_date": types.Int64Type}}, "type": types.StringType}},
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
			if nullParent, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"quote_details": types.ObjectType{AttrTypes: map[string]attr.Type{"quote": types.StringType}}, "subscription_details": types.ObjectType{AttrTypes: map[string]attr.Type{"metadata": types.MapType{ElemType: types.StringType}, "subscription": types.StringType, "subscription_proration_date": types.Int64Type}}, "type": types.StringType}}); ok {
				if typedParent, ok := nullParent.(types.Object); ok {
					state.Parent = typedParent
				}
			}
		}
	}
	{
		assignedPaymentSettings := false
		hadRawPaymentSettings := false
		if rawValuePaymentSettings, rawOk := plainValueAtPath(raw, "payment_settings"); rawOk {
			hadRawPaymentSettings = true
			if rawValuePaymentSettings != nil {
				sourcePaymentSettings := applyConfiguredKeyedListShapes(rawValuePaymentSettings, attrValueToPlain(state.PaymentSettings))
				if valuePaymentSettings, err := flattenPlainValue(sourcePaymentSettings, types.ObjectType{AttrTypes: map[string]attr.Type{"default_mandate": types.StringType, "payment_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_type": types.StringType}}, "verification_method": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"installments": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "plan": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}}, "request_three_d_secure": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "type": types.StringType}}, "funding_type": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "purpose": types.StringType}}}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_includes_iof": types.StringType, "expires_after_seconds": types.Int64Type}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}}}, "verification_method": types.StringType}}}}, "payment_method_types": types.ListType{ElemType: types.StringType}}}, "payment_settings", "raw response"); err != nil {
					return err
				} else {
					if typedPaymentSettings, ok := valuePaymentSettings.(types.Object); ok {
						state.PaymentSettings = typedPaymentSettings
						assignedPaymentSettings = true
					}
				}
			}
		}
		if !assignedPaymentSettings {
			if !hasRaw {
				if responseValuePaymentSettings, ok := plainFromResponseField(obj, "PaymentSettings"); ok {
					sourcePaymentSettings := applyConfiguredKeyedListShapes(responseValuePaymentSettings, attrValueToPlain(state.PaymentSettings))
					if valuePaymentSettings, err := flattenPlainValue(
						sourcePaymentSettings,
						types.ObjectType{AttrTypes: map[string]attr.Type{"default_mandate": types.StringType, "payment_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_type": types.StringType}}, "verification_method": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"installments": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "plan": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}}, "request_three_d_secure": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "type": types.StringType}}, "funding_type": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "purpose": types.StringType}}}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_includes_iof": types.StringType, "expires_after_seconds": types.Int64Type}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}}}, "verification_method": types.StringType}}}}, "payment_method_types": types.ListType{ElemType: types.StringType}}},
						"payment_settings",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPaymentSettings, ok := valuePaymentSettings.(types.Object); ok {
							state.PaymentSettings = typedPaymentSettings
							assignedPaymentSettings = true
						}
					}
				}
			}
		}
		if !assignedPaymentSettings && hadRawPaymentSettings {
			if nullPaymentSettings, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"default_mandate": types.StringType, "payment_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_type": types.StringType}}, "verification_method": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"installments": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "plan": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "interval": types.StringType, "type": types.StringType}}}}, "request_three_d_secure": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "type": types.StringType}}, "funding_type": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "purpose": types.StringType}}}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_includes_iof": types.StringType, "expires_after_seconds": types.Int64Type}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}}}, "verification_method": types.StringType}}}}, "payment_method_types": types.ListType{ElemType: types.StringType}}}); ok {
				if typedPaymentSettings, ok := nullPaymentSettings.(types.Object); ok {
					state.PaymentSettings = typedPaymentSettings
				}
			}
		}
	}
	{
		if rawValuePeriodEnd, rawOk := plainValueAtPath(raw, "period_end"); rawOk {
			if valuePeriodEnd, err := flattenPlainValue(rawValuePeriodEnd, types.Int64Type, "period_end", "raw response"); err != nil {
				return err
			} else {
				if typedPeriodEnd, ok := valuePeriodEnd.(types.Int64); ok {
					state.PeriodEnd = typedPeriodEnd
				}
			}
		} else if !hasRaw {
			if responseValuePeriodEnd, ok := plainFromResponseField(obj, "PeriodEnd"); ok {
				if valuePeriodEnd, err := flattenPlainValue(responseValuePeriodEnd, types.Int64Type, "period_end", "response struct"); err != nil {
					return err
				} else {
					if typedPeriodEnd, ok := valuePeriodEnd.(types.Int64); ok {
						state.PeriodEnd = typedPeriodEnd
					}
				}
			}
		}
	}
	{
		if rawValuePeriodStart, rawOk := plainValueAtPath(raw, "period_start"); rawOk {
			if valuePeriodStart, err := flattenPlainValue(rawValuePeriodStart, types.Int64Type, "period_start", "raw response"); err != nil {
				return err
			} else {
				if typedPeriodStart, ok := valuePeriodStart.(types.Int64); ok {
					state.PeriodStart = typedPeriodStart
				}
			}
		} else if !hasRaw {
			if responseValuePeriodStart, ok := plainFromResponseField(obj, "PeriodStart"); ok {
				if valuePeriodStart, err := flattenPlainValue(responseValuePeriodStart, types.Int64Type, "period_start", "response struct"); err != nil {
					return err
				} else {
					if typedPeriodStart, ok := valuePeriodStart.(types.Int64); ok {
						state.PeriodStart = typedPeriodStart
					}
				}
			}
		}
	}
	{
		if rawValuePostPaymentCreditNotesAmount, rawOk := plainValueAtPath(raw, "post_payment_credit_notes_amount"); rawOk {
			if valuePostPaymentCreditNotesAmount, err := flattenPlainValue(rawValuePostPaymentCreditNotesAmount, types.Int64Type, "post_payment_credit_notes_amount", "raw response"); err != nil {
				return err
			} else {
				if typedPostPaymentCreditNotesAmount, ok := valuePostPaymentCreditNotesAmount.(types.Int64); ok {
					state.PostPaymentCreditNotesAmount = typedPostPaymentCreditNotesAmount
				}
			}
		} else if !hasRaw {
			if responseValuePostPaymentCreditNotesAmount, ok := plainFromResponseField(obj, "PostPaymentCreditNotesAmount"); ok {
				if valuePostPaymentCreditNotesAmount, err := flattenPlainValue(responseValuePostPaymentCreditNotesAmount, types.Int64Type, "post_payment_credit_notes_amount", "response struct"); err != nil {
					return err
				} else {
					if typedPostPaymentCreditNotesAmount, ok := valuePostPaymentCreditNotesAmount.(types.Int64); ok {
						state.PostPaymentCreditNotesAmount = typedPostPaymentCreditNotesAmount
					}
				}
			}
		}
	}
	{
		if rawValuePrePaymentCreditNotesAmount, rawOk := plainValueAtPath(raw, "pre_payment_credit_notes_amount"); rawOk {
			if valuePrePaymentCreditNotesAmount, err := flattenPlainValue(rawValuePrePaymentCreditNotesAmount, types.Int64Type, "pre_payment_credit_notes_amount", "raw response"); err != nil {
				return err
			} else {
				if typedPrePaymentCreditNotesAmount, ok := valuePrePaymentCreditNotesAmount.(types.Int64); ok {
					state.PrePaymentCreditNotesAmount = typedPrePaymentCreditNotesAmount
				}
			}
		} else if !hasRaw {
			if responseValuePrePaymentCreditNotesAmount, ok := plainFromResponseField(obj, "PrePaymentCreditNotesAmount"); ok {
				if valuePrePaymentCreditNotesAmount, err := flattenPlainValue(responseValuePrePaymentCreditNotesAmount, types.Int64Type, "pre_payment_credit_notes_amount", "response struct"); err != nil {
					return err
				} else {
					if typedPrePaymentCreditNotesAmount, ok := valuePrePaymentCreditNotesAmount.(types.Int64); ok {
						state.PrePaymentCreditNotesAmount = typedPrePaymentCreditNotesAmount
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
		assignedRendering := false
		hadRawRendering := false
		if rawValueRendering, rawOk := plainValueAtPath(raw, "rendering"); rawOk {
			hadRawRendering = true
			if rawValueRendering != nil {
				sourceRendering := applyConfiguredKeyedListShapes(rawValueRendering, attrValueToPlain(state.Rendering))
				if valueRendering, err := flattenPlainValue(sourceRendering, types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "pdf": types.ObjectType{AttrTypes: map[string]attr.Type{"page_size": types.StringType}}, "template": types.StringType, "template_version": types.Int64Type}}, "rendering", "raw response"); err != nil {
					return err
				} else {
					if typedRendering, ok := valueRendering.(types.Object); ok {
						state.Rendering = typedRendering
						assignedRendering = true
					}
				}
			}
		}
		if !assignedRendering {
			if !hasRaw {
				if responseValueRendering, ok := plainFromResponseField(obj, "Rendering"); ok {
					sourceRendering := applyConfiguredKeyedListShapes(responseValueRendering, attrValueToPlain(state.Rendering))
					if valueRendering, err := flattenPlainValue(
						sourceRendering,
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "pdf": types.ObjectType{AttrTypes: map[string]attr.Type{"page_size": types.StringType}}, "template": types.StringType, "template_version": types.Int64Type}},
						"rendering",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRendering, ok := valueRendering.(types.Object); ok {
							state.Rendering = typedRendering
							assignedRendering = true
						}
					}
				}
			}
		}
		if !assignedRendering && hadRawRendering {
			if nullRendering, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "pdf": types.ObjectType{AttrTypes: map[string]attr.Type{"page_size": types.StringType}}, "template": types.StringType, "template_version": types.Int64Type}}); ok {
				if typedRendering, ok := nullRendering.(types.Object); ok {
					state.Rendering = typedRendering
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
				if valueShippingCost, err := flattenPlainValue(sourceShippingCost, types.ObjectType{AttrTypes: map[string]attr.Type{"amount_subtotal": types.Int64Type, "amount_tax": types.Int64Type, "amount_total": types.Int64Type, "shipping_rate": types.StringType, "taxes": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "rate": types.StringType, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}, "shipping_rate_data": types.ObjectType{AttrTypes: map[string]attr.Type{"delivery_estimate": types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}, "minimum": types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}}, "display_name": types.StringType, "fixed_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "currency_options": types.MapType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "tax_behavior": types.StringType}}}}}, "metadata": types.MapType{ElemType: types.StringType}, "tax_behavior": types.StringType, "tax_code": types.StringType, "type": types.StringType}}}}, "shipping_cost", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount_subtotal": types.Int64Type, "amount_tax": types.Int64Type, "amount_total": types.Int64Type, "shipping_rate": types.StringType, "taxes": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "rate": types.StringType, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}, "shipping_rate_data": types.ObjectType{AttrTypes: map[string]attr.Type{"delivery_estimate": types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}, "minimum": types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}}, "display_name": types.StringType, "fixed_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "currency_options": types.MapType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "tax_behavior": types.StringType}}}}}, "metadata": types.MapType{ElemType: types.StringType}, "tax_behavior": types.StringType, "tax_code": types.StringType, "type": types.StringType}}}},
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
			if nullShippingCost, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount_subtotal": types.Int64Type, "amount_tax": types.Int64Type, "amount_total": types.Int64Type, "shipping_rate": types.StringType, "taxes": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "rate": types.StringType, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}, "shipping_rate_data": types.ObjectType{AttrTypes: map[string]attr.Type{"delivery_estimate": types.ObjectType{AttrTypes: map[string]attr.Type{"maximum": types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}, "minimum": types.ObjectType{AttrTypes: map[string]attr.Type{"unit": types.StringType, "value": types.Int64Type}}}}, "display_name": types.StringType, "fixed_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "currency_options": types.MapType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "tax_behavior": types.StringType}}}}}, "metadata": types.MapType{ElemType: types.StringType}, "tax_behavior": types.StringType, "tax_code": types.StringType, "type": types.StringType}}}}); ok {
				if typedShippingCost, ok := nullShippingCost.(types.Object); ok {
					state.ShippingCost = typedShippingCost
				}
			}
		}
	}
	{
		assignedShippingDetails := false
		hadRawShippingDetails := false
		if rawValueShippingDetails, rawOk := plainValueAtPath(raw, "shipping_details"); rawOk {
			hadRawShippingDetails = true
			if rawValueShippingDetails != nil {
				sourceShippingDetails := applyConfiguredKeyedListShapes(rawValueShippingDetails, attrValueToPlain(state.ShippingDetails))
				if valueShippingDetails, err := flattenPlainValue(sourceShippingDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}, "shipping_details", "raw response"); err != nil {
					return err
				} else {
					if typedShippingDetails, ok := valueShippingDetails.(types.Object); ok {
						state.ShippingDetails = typedShippingDetails
						assignedShippingDetails = true
					}
				}
			}
		}
		if !assignedShippingDetails {
			if !hasRaw {
				if responseValueShippingDetails, ok := plainFromResponseField(obj, "ShippingDetails"); ok {
					sourceShippingDetails := applyConfiguredKeyedListShapes(responseValueShippingDetails, attrValueToPlain(state.ShippingDetails))
					if valueShippingDetails, err := flattenPlainValue(
						sourceShippingDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}},
						"shipping_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedShippingDetails, ok := valueShippingDetails.(types.Object); ok {
							state.ShippingDetails = typedShippingDetails
							assignedShippingDetails = true
						}
					}
				}
			}
		}
		if !assignedShippingDetails && hadRawShippingDetails {
			if nullShippingDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}); ok {
				if typedShippingDetails, ok := nullShippingDetails.(types.Object); ok {
					state.ShippingDetails = typedShippingDetails
				}
			}
		}
	}
	{
		if rawValueStartingBalance, rawOk := plainValueAtPath(raw, "starting_balance"); rawOk {
			if valueStartingBalance, err := flattenPlainValue(rawValueStartingBalance, types.Int64Type, "starting_balance", "raw response"); err != nil {
				return err
			} else {
				if typedStartingBalance, ok := valueStartingBalance.(types.Int64); ok {
					state.StartingBalance = typedStartingBalance
				}
			}
		} else if !hasRaw {
			if responseValueStartingBalance, ok := plainFromResponseField(obj, "StartingBalance"); ok {
				if valueStartingBalance, err := flattenPlainValue(responseValueStartingBalance, types.Int64Type, "starting_balance", "response struct"); err != nil {
					return err
				} else {
					if typedStartingBalance, ok := valueStartingBalance.(types.Int64); ok {
						state.StartingBalance = typedStartingBalance
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
				if valueStatusTransitions, err := flattenPlainValue(sourceStatusTransitions, types.ObjectType{AttrTypes: map[string]attr.Type{"finalized_at": types.Int64Type, "marked_uncollectible_at": types.Int64Type, "paid_at": types.Int64Type, "voided_at": types.Int64Type}}, "status_transitions", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"finalized_at": types.Int64Type, "marked_uncollectible_at": types.Int64Type, "paid_at": types.Int64Type, "voided_at": types.Int64Type}},
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
			if nullStatusTransitions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"finalized_at": types.Int64Type, "marked_uncollectible_at": types.Int64Type, "paid_at": types.Int64Type, "voided_at": types.Int64Type}}); ok {
				if typedStatusTransitions, ok := nullStatusTransitions.(types.Object); ok {
					state.StatusTransitions = typedStatusTransitions
				}
			}
		}
	}
	{
		if state.Subscription.IsNull() || state.Subscription.IsUnknown() {
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
	{
		assignedThresholdReason := false
		hadRawThresholdReason := false
		if rawValueThresholdReason, rawOk := plainValueAtPath(raw, "threshold_reason"); rawOk {
			hadRawThresholdReason = true
			if rawValueThresholdReason != nil {
				sourceThresholdReason := applyConfiguredKeyedListShapes(rawValueThresholdReason, attrValueToPlain(state.ThresholdReason))
				if valueThresholdReason, err := flattenPlainValue(sourceThresholdReason, types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "item_reasons": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"line_item_ids": types.ListType{ElemType: types.StringType}, "usage_gte": types.Int64Type}}}}}, "threshold_reason", "raw response"); err != nil {
					return err
				} else {
					if typedThresholdReason, ok := valueThresholdReason.(types.Object); ok {
						state.ThresholdReason = typedThresholdReason
						assignedThresholdReason = true
					}
				}
			}
		}
		if !assignedThresholdReason {
			if !hasRaw {
				if responseValueThresholdReason, ok := plainFromResponseField(obj, "ThresholdReason"); ok {
					sourceThresholdReason := applyConfiguredKeyedListShapes(responseValueThresholdReason, attrValueToPlain(state.ThresholdReason))
					if valueThresholdReason, err := flattenPlainValue(
						sourceThresholdReason,
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "item_reasons": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"line_item_ids": types.ListType{ElemType: types.StringType}, "usage_gte": types.Int64Type}}}}},
						"threshold_reason",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedThresholdReason, ok := valueThresholdReason.(types.Object); ok {
							state.ThresholdReason = typedThresholdReason
							assignedThresholdReason = true
						}
					}
				}
			}
		}
		if !assignedThresholdReason && hadRawThresholdReason {
			if nullThresholdReason, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "item_reasons": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"line_item_ids": types.ListType{ElemType: types.StringType}, "usage_gte": types.Int64Type}}}}}); ok {
				if typedThresholdReason, ok := nullThresholdReason.(types.Object); ok {
					state.ThresholdReason = typedThresholdReason
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
		if rawValueTotalDiscountAmounts, rawOk := plainValueAtPath(raw, "total_discount_amounts"); rawOk {
			if valueTotalDiscountAmounts, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueTotalDiscountAmounts, attrValueToPlain(state.TotalDiscountAmounts)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "discount": types.StringType}}}, "total_discount_amounts", "raw response"); err != nil {
				return err
			} else {
				if typedTotalDiscountAmounts, ok := valueTotalDiscountAmounts.(types.List); ok {
					state.TotalDiscountAmounts = typedTotalDiscountAmounts
				}
			}
		} else if !hasRaw {
			if responseValueTotalDiscountAmounts, ok := plainFromResponseField(obj, "TotalDiscountAmounts"); ok {
				if valueTotalDiscountAmounts, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueTotalDiscountAmounts, attrValueToPlain(state.TotalDiscountAmounts)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "discount": types.StringType}}},
					"total_discount_amounts",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedTotalDiscountAmounts, ok := valueTotalDiscountAmounts.(types.List); ok {
						state.TotalDiscountAmounts = typedTotalDiscountAmounts
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
		if rawValueTotalPretaxCreditAmounts, rawOk := plainValueAtPath(raw, "total_pretax_credit_amounts"); rawOk {
			if valueTotalPretaxCreditAmounts, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueTotalPretaxCreditAmounts, attrValueToPlain(state.TotalPretaxCreditAmounts)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "credit_balance_transaction": types.StringType, "discount": types.StringType, "type": types.StringType}}}, "total_pretax_credit_amounts", "raw response"); err != nil {
				return err
			} else {
				if typedTotalPretaxCreditAmounts, ok := valueTotalPretaxCreditAmounts.(types.List); ok {
					state.TotalPretaxCreditAmounts = typedTotalPretaxCreditAmounts
				}
			}
		} else if !hasRaw {
			if responseValueTotalPretaxCreditAmounts, ok := plainFromResponseField(obj, "TotalPretaxCreditAmounts"); ok {
				if valueTotalPretaxCreditAmounts, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueTotalPretaxCreditAmounts, attrValueToPlain(state.TotalPretaxCreditAmounts)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "credit_balance_transaction": types.StringType, "discount": types.StringType, "type": types.StringType}}},
					"total_pretax_credit_amounts",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedTotalPretaxCreditAmounts, ok := valueTotalPretaxCreditAmounts.(types.List); ok {
						state.TotalPretaxCreditAmounts = typedTotalPretaxCreditAmounts
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
		if rawValueWebhooksDeliveredAt, rawOk := plainValueAtPath(raw, "webhooks_delivered_at"); rawOk {
			if valueWebhooksDeliveredAt, err := flattenPlainValue(rawValueWebhooksDeliveredAt, types.Int64Type, "webhooks_delivered_at", "raw response"); err != nil {
				return err
			} else {
				if typedWebhooksDeliveredAt, ok := valueWebhooksDeliveredAt.(types.Int64); ok {
					state.WebhooksDeliveredAt = typedWebhooksDeliveredAt
				}
			}
		} else if !hasRaw {
			if responseValueWebhooksDeliveredAt, ok := plainFromResponseField(obj, "WebhooksDeliveredAt"); ok {
				if valueWebhooksDeliveredAt, err := flattenPlainValue(responseValueWebhooksDeliveredAt, types.Int64Type, "webhooks_delivered_at", "response struct"); err != nil {
					return err
				} else {
					if typedWebhooksDeliveredAt, ok := valueWebhooksDeliveredAt.(types.Int64); ok {
						state.WebhooksDeliveredAt = typedWebhooksDeliveredAt
					}
				}
			}
		}
	}
	return nil
}
