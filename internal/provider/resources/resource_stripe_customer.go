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

var _ resource.Resource = &CustomerResource{}

var _ resource.ResourceWithConfigure = &CustomerResource{}

var _ resource.ResourceWithImportState = &CustomerResource{}

func NewCustomerResource() resource.Resource {
	return &CustomerResource{}
}

type CustomerResource struct {
	client *stripe.Client
}

type CustomerResourceModel struct {
	Object               types.String `tfsdk:"object"`
	Address              types.List   `tfsdk:"address"`
	Balance              types.Int64  `tfsdk:"balance"`
	BusinessName         types.String `tfsdk:"business_name"`
	CashBalance          types.List   `tfsdk:"cash_balance"`
	Created              types.Int64  `tfsdk:"created"`
	Currency             types.String `tfsdk:"currency"`
	CustomerAccount      types.String `tfsdk:"customer_account"`
	DefaultSource        types.String `tfsdk:"default_source"`
	Delinquent           types.Bool   `tfsdk:"delinquent"`
	Description          types.String `tfsdk:"description"`
	Discount             types.String `tfsdk:"discount"`
	Email                types.String `tfsdk:"email"`
	ID                   types.String `tfsdk:"id"`
	IndividualName       types.String `tfsdk:"individual_name"`
	InvoiceCreditBalance types.Map    `tfsdk:"invoice_credit_balance"`
	InvoicePrefix        types.String `tfsdk:"invoice_prefix"`
	InvoiceSettings      types.List   `tfsdk:"invoice_settings"`
	Livemode             types.Bool   `tfsdk:"livemode"`
	Metadata             types.Map    `tfsdk:"metadata"`
	Name                 types.String `tfsdk:"name"`
	NextInvoiceSequence  types.Int64  `tfsdk:"next_invoice_sequence"`
	Phone                types.String `tfsdk:"phone"`
	PreferredLocales     types.List   `tfsdk:"preferred_locales"`
	Shipping             types.List   `tfsdk:"shipping"`
	Tax                  types.List   `tfsdk:"tax"`
	TaxExempt            types.String `tfsdk:"tax_exempt"`
	TestClock            types.String `tfsdk:"test_clock"`
	PaymentMethod        types.String `tfsdk:"payment_method"`
	Source               types.String `tfsdk:"source"`
	TaxIDData            types.List   `tfsdk:"tax_id_data"`
	Validate             types.Bool   `tfsdk:"validate"`
}

func (r *CustomerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CustomerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_customer"
}

var _ resource.ResourceWithUpgradeState = &CustomerResource{}

func (r *CustomerResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: customerResourceV0Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior CustomerResourceModel
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeCustomerStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
		1: {
			PriorSchema: customerResourceV1Schema(),
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var prior CustomerResourceV1Model
				resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgraded, diags := upgradeCustomerStateV1(ctx, prior)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, &upgraded)...)
			},
		},
	}
}

func customerResourceV1Schema() *schema.Schema {
	return &schema.Schema{
		Description: "This object represents a customer of your business. Use it to [create recurring charges](https://docs.stripe.com/invoicing/customer), [save payment](https://docs.stripe.com/payments/save-during-payment) and contact information,\nand track payments that belong to the same customer.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("customer")},
			},
			"address": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's address.",
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
			"balance": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The current balance, if any, that's stored on the customer in their default currency. If negative, the customer has credit to apply to their next invoice. If positive, the customer has an amount owed that's added to their next invoice. The balance only considers amounts that Stripe hasn't successfully applied to any invoice. It doesn't reflect unpaid invoices. This balance is only taken into account after invoices finalize. For multi-currency balances, see [invoice_credit_balance](https://docs.stripe.com/api/customers/object#customer_object-invoice_credit_balance).",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"business_name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's business name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cash_balance": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The current funds being held by Stripe on behalf of the customer. You can apply these funds towards payment intents when the source is \"cash_balance\". The `settings[reconciliation_mode]` field describes if these funds apply to these payment intents manually or automatically.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"object": schema.StringAttribute{
						Computed:      true,
						Description:   "String representing the object's type. Objects of the same type share the same value.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("cash_balance")},
					},
					"available": schema.MapAttribute{
						Computed:      true,
						Description:   "A hash of all cash balances available to this customer. You cannot delete a customer with any cash balances, even if the balance is 0. Amounts are represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						ElementType:   types.Int64Type,
					},
					"customer": schema.StringAttribute{
						Computed:      true,
						Description:   "The ID of the customer whose cash balance this object represents.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"customer_account": schema.StringAttribute{
						Computed:      true,
						Description:   "The ID of an Account representing a customer whose cash balance this object represents.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"livemode": schema.BoolAttribute{
						Computed:      true,
						Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"settings": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"reconciliation_mode": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The configuration for how funds that land in the customer cash balance are reconciled.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("automatic", "manual")},
							},
							"using_merchant_default": schema.BoolAttribute{
								Computed:      true,
								Description:   "A flag to indicate if reconciliation mode returned is the user's default or is specific to this customer cash balance",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Computed:      true,
				Description:   "Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) the customer can be charged in for recurring billing purposes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_account": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of an Account representing a customer. You can use this ID with any v1 API that accepts a customer_account parameter.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_source": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the default payment source for the customer.\n\nIf you use payment methods created through the PaymentMethods API, see the [invoice_settings.default_payment_method](https://docs.stripe.com/api/customers/object#customer_object-invoice_settings-default_payment_method) field instead.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"delinquent": schema.BoolAttribute{
				Computed:      true,
				Description:   "Tracks the most recent state change on any invoice belonging to the customer. Paying an invoice or marking it uncollectible via the API will set this field to false. An automatic payment failure or passing the `invoice.due_date` will set this field to `true`.\n\nIf an invoice becomes uncollectible by [dunning](https://docs.stripe.com/billing/automatic-collection), `delinquent` doesn't reset to `false`.\n\nIf you care whether the customer has paid their most recent subscription invoice, use `subscription.status` instead. Paying or marking uncollectible any customer invoice regardless of whether it is the latest invoice for a subscription will always set this field to `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the object. Often useful for displaying to users.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"discount": schema.StringAttribute{
				Computed:      true,
				Description:   "Describes the current discount active on the customer, if there is one.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"email": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's email address.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"individual_name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's individual name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice_credit_balance": schema.MapAttribute{
				Computed:      true,
				Description:   "The current multi-currency balances, if any, that's stored on the customer. If positive in a currency, the customer has a credit to apply to their next invoice denominated in that currency. If negative, the customer has an amount owed that's added to their next invoice denominated in that currency. These balances don't apply to unpaid invoices. They solely track amounts that Stripe hasn't successfully applied to any invoice. Stripe only applies a balance in a specific currency to an invoice after that invoice (which is in the same currency) finalizes.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.Int64Type,
			},
			"invoice_prefix": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The prefix for the customer used to generate unique invoice numbers.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice_settings": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"custom_fields": schema.ListNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Default custom fields to be displayed on invoices for this customer.",
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
					"default_payment_method": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "ID of a payment method that's attached to the customer, to be used as the customer's default payment method for subscriptions and invoices.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"footer": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Default footer to be displayed on invoices for this customer.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"rendering_options": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Default options for invoice PDF rendering for this customer.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"amount_tax_display": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"template": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "ID of the invoice rendering template to be used for this customer's invoices. If set, the template will be used on all invoices for this customer unless a template is set directly on the invoice.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's full name or business name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"next_invoice_sequence": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The suffix of the customer's next invoice number (for example, 0001). When the account uses account level sequencing, this parameter is ignored in API requests and the field omitted in API responses.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"phone": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's phone number.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"preferred_locales": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's preferred locales (languages), ordered by preference.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"shipping": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Mailing and shipping address for the customer. Appears on invoices emailed to this customer.",
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
			"tax": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"automatic_tax": schema.StringAttribute{
						Computed:      true,
						Description:   "Surfaces if automatic tax computation is possible given the current customer location information.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("failed", "not_collecting", "supported", "unrecognized_location")},
					},
					"ip_address": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A recent IP address of the customer used for tax reporting and tax location inference.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"location": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "The identified tax location of the customer.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"country": schema.StringAttribute{
								Computed:      true,
								Description:   "The identified tax country of the customer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"source": schema.StringAttribute{
								Computed:      true,
								Description:   "The data source used to infer the customer's location.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("billing_address", "ip_address", "payment_method", "shipping_destination")},
							},
							"state": schema.StringAttribute{
								Computed:      true,
								Description:   "The identified tax state, county, province, or region of the customer.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"provider": schema.StringAttribute{
						Computed:      true,
						Description:   "The tax calculation provider used for location resolution. Defaults to `stripe` when not using a [third-party provider](/tax/third-party-apps).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("anrok", "avalara", "sphere", "stripe")},
					},
					"validate_location": schema.StringAttribute{
						Optional:    true,
						Description: "A flag that indicates when Stripe should validate the customer tax location. Defaults to `deferred`.",
						Validators:  []validator.String{stringvalidator.OneOf("deferred", "immediately")},
					},
				},
			},
			"tax_exempt": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Describes the customer's tax exemption status, which is `none`, `exempt`, or `reverse`. When set to `reverse`, invoice and receipt PDFs include the following text: **\"Reverse charge\"**.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("exempt", "none", "reverse")},
			},
			"test_clock": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the test clock that this customer belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"payment_method": schema.StringAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"source": schema.StringAttribute{
				Optional: true,

				WriteOnly: true,
			},
			"tax_id_data": schema.ListNestedAttribute{
				Optional:      true,
				Description:   "The customer's tax IDs.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Required:      true,
							Description:   "Type of the tax ID, one of `ad_nrt`, `ae_trn`, `al_tin`, `am_tin`, `ao_tin`, `ar_cuit`, `au_abn`, `au_arn`, `aw_tin`, `az_tin`, `ba_tin`, `bb_tin`, `bd_bin`, `bf_ifu`, `bg_uic`, `bh_vat`, `bj_ifu`, `bo_tin`, `br_cnpj`, `br_cpf`, `bs_tin`, `by_tin`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `cd_nif`, `ch_uid`, `ch_vat`, `cl_tin`, `cm_niu`, `cn_tin`, `co_nit`, `cr_tin`, `cv_nif`, `de_stn`, `do_rcn`, `ec_ruc`, `eg_tin`, `es_cif`, `et_tin`, `eu_oss_vat`, `eu_vat`, `fo_vat`, `gb_vat`, `ge_vat`, `gi_tin`, `gn_nif`, `hk_br`, `hr_oib`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `it_cf`, `jp_cn`, `jp_rn`, `jp_trn`, `ke_pin`, `kg_tin`, `kh_tin`, `kr_brn`, `kz_bin`, `la_tin`, `li_uid`, `li_vat`, `lk_vat`, `ma_vat`, `md_vat`, `me_pib`, `mk_vat`, `mr_nif`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `ng_tin`, `no_vat`, `no_voec`, `np_pan`, `nz_gst`, `om_vat`, `pe_ruc`, `ph_tin`, `pl_nip`, `py_ruc`, `ro_tin`, `rs_pib`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `sn_ninea`, `sr_fin`, `sv_nit`, `th_vat`, `tj_tin`, `tr_tin`, `tw_vat`, `tz_vat`, `ua_vat`, `ug_tin`, `us_ein`, `uy_ruc`, `uz_tin`, `uz_vat`, `ve_rif`, `vn_tin`, `za_vat`, `zm_tin`, or `zw_tin`",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"value": schema.StringAttribute{
							Required:      true,
							Description:   "Value of the tax ID.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
			"validate": schema.BoolAttribute{
				Optional: true,

				WriteOnly: true,
			},
		},
	}
}

func customerResourceV0Schema() *schema.Schema {
	return &schema.Schema{
		Description: "This object represents a customer of your business. Use it to [create recurring charges](https://docs.stripe.com/invoicing/customer), [save payment](https://docs.stripe.com/payments/save-during-payment) and contact information,\nand track payments that belong to the same customer.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("customer")},
			},
			"balance": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The current balance, if any, that's stored on the customer in their default currency. If negative, the customer has credit to apply to their next invoice. If positive, the customer has an amount owed that's added to their next invoice. The balance only considers amounts that Stripe hasn't successfully applied to any invoice. It doesn't reflect unpaid invoices. This balance is only taken into account after invoices finalize. For multi-currency balances, see [invoice_credit_balance](https://docs.stripe.com/api/customers/object#customer_object-invoice_credit_balance).",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"business_name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's business name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Computed:      true,
				Description:   "Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) the customer can be charged in for recurring billing purposes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_account": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of an Account representing a customer. You can use this ID with any v1 API that accepts a customer_account parameter.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_source": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the default payment source for the customer.\n\nIf you use payment methods created through the PaymentMethods API, see the [invoice_settings.default_payment_method](https://docs.stripe.com/api/customers/object#customer_object-invoice_settings-default_payment_method) field instead.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"delinquent": schema.BoolAttribute{
				Computed:      true,
				Description:   "Tracks the most recent state change on any invoice belonging to the customer. Paying an invoice or marking it uncollectible via the API will set this field to false. An automatic payment failure or passing the `invoice.due_date` will set this field to `true`.\n\nIf an invoice becomes uncollectible by [dunning](https://docs.stripe.com/billing/automatic-collection), `delinquent` doesn't reset to `false`.\n\nIf you care whether the customer has paid their most recent subscription invoice, use `subscription.status` instead. Paying or marking uncollectible any customer invoice regardless of whether it is the latest invoice for a subscription will always set this field to `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the object. Often useful for displaying to users.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"discount": schema.StringAttribute{
				Computed:      true,
				Description:   "Describes the current discount active on the customer, if there is one.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"email": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's email address.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"individual_name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's individual name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice_credit_balance": schema.MapAttribute{
				Computed:      true,
				Description:   "The current multi-currency balances, if any, that's stored on the customer. If positive in a currency, the customer has a credit to apply to their next invoice denominated in that currency. If negative, the customer has an amount owed that's added to their next invoice denominated in that currency. These balances don't apply to unpaid invoices. They solely track amounts that Stripe hasn't successfully applied to any invoice. Stripe only applies a balance in a specific currency to an invoice after that invoice (which is in the same currency) finalizes.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.Int64Type,
			},
			"invoice_prefix": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The prefix for the customer used to generate unique invoice numbers.",
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
			"name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's full name or business name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"next_invoice_sequence": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The suffix of the customer's next invoice number (for example, 0001). When the account uses account level sequencing, this parameter is ignored in API requests and the field omitted in API responses.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"phone": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's phone number.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"preferred_locales": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's preferred locales (languages), ordered by preference.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"tax_exempt": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Describes the customer's tax exemption status, which is `none`, `exempt`, or `reverse`. When set to `reverse`, invoice and receipt PDFs include the following text: **\"Reverse charge\"**.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("exempt", "none", "reverse")},
			},
			"test_clock": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the test clock that this customer belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"payment_method": schema.StringAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"source": schema.StringAttribute{
				Optional: true,

				WriteOnly: true,
			},
			"validate": schema.BoolAttribute{
				Optional: true,

				WriteOnly: true,
			},
		},
		Blocks: map[string]schema.Block{
			"address": schema.ListNestedBlock{
				Description:   "The customer's address.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
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
			},
			"cash_balance": schema.ListNestedBlock{
				Description:   "The current funds being held by Stripe on behalf of the customer. You can apply these funds towards payment intents when the source is \"cash_balance\". The `settings[reconciliation_mode]` field describes if these funds apply to these payment intents manually or automatically.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"object": schema.StringAttribute{
							Computed:      true,
							Description:   "String representing the object's type. Objects of the same type share the same value.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("cash_balance")},
						},
						"available": schema.MapAttribute{
							Computed:      true,
							Description:   "A hash of all cash balances available to this customer. You cannot delete a customer with any cash balances, even if the balance is 0. Amounts are represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).",
							PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
							ElementType:   types.Int64Type,
						},
						"customer": schema.StringAttribute{
							Computed:      true,
							Description:   "The ID of the customer whose cash balance this object represents.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"customer_account": schema.StringAttribute{
							Computed:      true,
							Description:   "The ID of an Account representing a customer whose cash balance this object represents.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"livemode": schema.BoolAttribute{
							Computed:      true,
							Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
							PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
						},
					},
					Blocks: map[string]schema.Block{
						"settings": schema.ListNestedBlock{
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"reconciliation_mode": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The configuration for how funds that land in the customer cash balance are reconciled.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("automatic", "manual")},
									},
									"using_merchant_default": schema.BoolAttribute{
										Computed:      true,
										Description:   "A flag to indicate if reconciliation mode returned is the user's default or is specific to this customer cash balance",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
				},
			},
			"invoice_settings": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"custom_fields": schema.ListNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Default custom fields to be displayed on invoices for this customer.",
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
						"default_payment_method": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "ID of a payment method that's attached to the customer, to be used as the customer's default payment method for subscriptions and invoices.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"footer": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Default footer to be displayed on invoices for this customer.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
					},
					Blocks: map[string]schema.Block{
						"rendering_options": schema.ListNestedBlock{
							Description:   "Default options for invoice PDF rendering for this customer.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"amount_tax_display": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"template": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "ID of the invoice rendering template to be used for this customer's invoices. If set, the template will be used on all invoices for this customer unless a template is set directly on the invoice.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
				},
			},
			"shipping": schema.ListNestedBlock{
				Description:   "Mailing and shipping address for the customer. Appears on invoices emailed to this customer.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
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
					Blocks: map[string]schema.Block{
						"address": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
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
						},
					},
				},
			},
			"tax": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"automatic_tax": schema.StringAttribute{
							Computed:      true,
							Description:   "Surfaces if automatic tax computation is possible given the current customer location information.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("failed", "not_collecting", "supported", "unrecognized_location")},
						},
						"ip_address": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "A recent IP address of the customer used for tax reporting and tax location inference.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"provider": schema.StringAttribute{
							Computed:      true,
							Description:   "The tax calculation provider used for location resolution. Defaults to `stripe` when not using a [third-party provider](/tax/third-party-apps).",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("anrok", "avalara", "sphere", "stripe")},
						},
						"validate_location": schema.StringAttribute{
							Optional:    true,
							Description: "A flag that indicates when Stripe should validate the customer tax location. Defaults to `deferred`.",
							Validators:  []validator.String{stringvalidator.OneOf("deferred", "immediately")},
						},
					},
					Blocks: map[string]schema.Block{
						"location": schema.ListNestedBlock{
							Description:   "The identified tax location of the customer.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"country": schema.StringAttribute{
										Computed:      true,
										Description:   "The identified tax country of the customer.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"source": schema.StringAttribute{
										Computed:      true,
										Description:   "The data source used to infer the customer's location.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("billing_address", "ip_address", "payment_method", "shipping_destination")},
									},
									"state": schema.StringAttribute{
										Computed:      true,
										Description:   "The identified tax state, county, province, or region of the customer.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
				},
			},
			"tax_id_data": schema.ListNestedBlock{
				Description:   "The customer's tax IDs.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Required:      true,
							Description:   "Type of the tax ID, one of `ad_nrt`, `ae_trn`, `al_tin`, `am_tin`, `ao_tin`, `ar_cuit`, `au_abn`, `au_arn`, `aw_tin`, `az_tin`, `ba_tin`, `bb_tin`, `bd_bin`, `bf_ifu`, `bg_uic`, `bh_vat`, `bj_ifu`, `bo_tin`, `br_cnpj`, `br_cpf`, `bs_tin`, `by_tin`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `cd_nif`, `ch_uid`, `ch_vat`, `cl_tin`, `cm_niu`, `cn_tin`, `co_nit`, `cr_tin`, `cv_nif`, `de_stn`, `do_rcn`, `ec_ruc`, `eg_tin`, `es_cif`, `et_tin`, `eu_oss_vat`, `eu_vat`, `fo_vat`, `gb_vat`, `ge_vat`, `gi_tin`, `gn_nif`, `hk_br`, `hr_oib`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `it_cf`, `jp_cn`, `jp_rn`, `jp_trn`, `ke_pin`, `kg_tin`, `kh_tin`, `kr_brn`, `kz_bin`, `la_tin`, `li_uid`, `li_vat`, `lk_vat`, `ma_vat`, `md_vat`, `me_pib`, `mk_vat`, `mr_nif`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `ng_tin`, `no_vat`, `no_voec`, `np_pan`, `nz_gst`, `om_vat`, `pe_ruc`, `ph_tin`, `pl_nip`, `py_ruc`, `ro_tin`, `rs_pib`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `sn_ninea`, `sr_fin`, `sv_nit`, `th_vat`, `tj_tin`, `tr_tin`, `tw_vat`, `tz_vat`, `ua_vat`, `ug_tin`, `us_ein`, `uy_ruc`, `uz_tin`, `uz_vat`, `ve_rif`, `vn_tin`, `za_vat`, `zm_tin`, or `zw_tin`",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"value": schema.StringAttribute{
							Required:      true,
							Description:   "Value of the tax ID.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
		},
	}
}

type CustomerResourceV1Model struct {
	Object               types.String `tfsdk:"object"`
	Address              types.Object `tfsdk:"address"`
	Balance              types.Int64  `tfsdk:"balance"`
	BusinessName         types.String `tfsdk:"business_name"`
	CashBalance          types.Object `tfsdk:"cash_balance"`
	Created              types.Int64  `tfsdk:"created"`
	Currency             types.String `tfsdk:"currency"`
	CustomerAccount      types.String `tfsdk:"customer_account"`
	DefaultSource        types.String `tfsdk:"default_source"`
	Delinquent           types.Bool   `tfsdk:"delinquent"`
	Description          types.String `tfsdk:"description"`
	Discount             types.String `tfsdk:"discount"`
	Email                types.String `tfsdk:"email"`
	ID                   types.String `tfsdk:"id"`
	IndividualName       types.String `tfsdk:"individual_name"`
	InvoiceCreditBalance types.Map    `tfsdk:"invoice_credit_balance"`
	InvoicePrefix        types.String `tfsdk:"invoice_prefix"`
	InvoiceSettings      types.Object `tfsdk:"invoice_settings"`
	Livemode             types.Bool   `tfsdk:"livemode"`
	Metadata             types.Map    `tfsdk:"metadata"`
	Name                 types.String `tfsdk:"name"`
	NextInvoiceSequence  types.Int64  `tfsdk:"next_invoice_sequence"`
	Phone                types.String `tfsdk:"phone"`
	PreferredLocales     types.List   `tfsdk:"preferred_locales"`
	Shipping             types.Object `tfsdk:"shipping"`
	Tax                  types.Object `tfsdk:"tax"`
	TaxExempt            types.String `tfsdk:"tax_exempt"`
	TestClock            types.String `tfsdk:"test_clock"`
	PaymentMethod        types.String `tfsdk:"payment_method"`
	Source               types.String `tfsdk:"source"`
	TaxIDData            types.List   `tfsdk:"tax_id_data"`
	Validate             types.Bool   `tfsdk:"validate"`
}

type customerStateUpgradeAttrMeta struct {
	AttrType                attr.Type
	Behavior                string
	LegacyBehavior          string
	PreserveConfiguredValue bool
	Nested                  map[string]customerStateUpgradeAttrMeta
}

var customerStateUpgradeRootMeta = map[string]customerStateUpgradeAttrMeta{"object": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "address": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]customerStateUpgradeAttrMeta{"city": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "country": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "line1": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "line2": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "postal_code": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "state": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}}, "balance": customerStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "business_name": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "cash_balance": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"object": types.StringType, "available": types.MapType{ElemType: types.Int64Type}, "customer": types.StringType, "customer_account": types.StringType, "livemode": types.BoolType, "settings": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"reconciliation_mode": types.StringType, "using_merchant_default": types.BoolType}}}}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]customerStateUpgradeAttrMeta{"object": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "available": customerStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.Int64Type}, Behavior: "computed", LegacyBehavior: "computed"}, "customer": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "customer_account": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "livemode": customerStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "settings": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"reconciliation_mode": types.StringType, "using_merchant_default": types.BoolType}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]customerStateUpgradeAttrMeta{"reconciliation_mode": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "using_merchant_default": customerStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}}}}}, "created": customerStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "computed", LegacyBehavior: "computed"}, "currency": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "customer_account": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "default_source": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "delinquent": customerStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "description": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "discount": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "email": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "id": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "individual_name": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "invoice_credit_balance": customerStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.Int64Type}, Behavior: "computed", LegacyBehavior: "computed"}, "invoice_prefix": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "invoice_settings": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"custom_fields": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "default_payment_method": types.StringType, "footer": types.StringType, "rendering_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]customerStateUpgradeAttrMeta{"custom_fields": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]customerStateUpgradeAttrMeta{"name": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "value": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "default_payment_method": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "footer": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "rendering_options": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]customerStateUpgradeAttrMeta{"amount_tax_display": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "template": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}}}}, "livemode": customerStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "computed", LegacyBehavior: "computed"}, "metadata": customerStateUpgradeAttrMeta{AttrType: types.MapType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "name": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "next_invoice_sequence": customerStateUpgradeAttrMeta{AttrType: types.Int64Type, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "phone": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "preferred_locales": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.StringType}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "shipping": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]customerStateUpgradeAttrMeta{"address": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}, Behavior: "required", LegacyBehavior: "required", Nested: map[string]customerStateUpgradeAttrMeta{"city": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "country": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "line1": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "line2": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "postal_code": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "state": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}}}, "carrier": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "name": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "phone": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "tracking_number": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}}}, "tax": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"automatic_tax": types.StringType, "ip_address": types.StringType, "location": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "source": types.StringType, "state": types.StringType}}}, "provider": types.StringType, "validate_location": types.StringType}}}, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed", Nested: map[string]customerStateUpgradeAttrMeta{"automatic_tax": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "ip_address": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "location": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "source": types.StringType, "state": types.StringType}}}, Behavior: "computed", LegacyBehavior: "computed", Nested: map[string]customerStateUpgradeAttrMeta{"country": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "source": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "state": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}}}, "provider": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "computed", LegacyBehavior: "computed"}, "validate_location": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}}}, "tax_exempt": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "test_clock": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional_and_computed", LegacyBehavior: "optional_and_computed"}, "payment_method": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional", PreserveConfiguredValue: true}, "source": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "optional", LegacyBehavior: "optional"}, "tax_id_data": customerStateUpgradeAttrMeta{AttrType: types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}}, Behavior: "optional", LegacyBehavior: "optional", PreserveConfiguredValue: true, Nested: map[string]customerStateUpgradeAttrMeta{"type": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}, "value": customerStateUpgradeAttrMeta{AttrType: types.StringType, Behavior: "required", LegacyBehavior: "required"}}}, "validate": customerStateUpgradeAttrMeta{AttrType: types.BoolType, Behavior: "optional", LegacyBehavior: "optional"}}

var customerStateUpgradeSingletonPaths = map[string]struct{}{}

var customerStateUpgradeLegacyObjectPaths = map[string]struct{}{"address": struct{}{}, "cash_balance": struct{}{}, "cash_balance.settings": struct{}{}, "invoice_settings": struct{}{}, "invoice_settings.rendering_options": struct{}{}, "shipping": struct{}{}, "shipping.address": struct{}{}, "tax": struct{}{}, "tax.location": struct{}{}}

func customerAttrMapFromModel(model interface{}) map[string]attr.Value {
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

func customerSetModelFromAttrMap(target interface{}, values map[string]attr.Value) {
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

func customerIsComputedBehavior(behavior string) bool {
	return behavior == "computed" || behavior == "optional_and_computed"
}

func customerShouldPreserveChild(parent customerStateUpgradeAttrMeta, child customerStateUpgradeAttrMeta) bool {
	if parent.PreserveConfiguredValue {
		return true
	}
	if !customerIsComputedBehavior(child.LegacyBehavior) {
		return true
	}
	return !customerIsComputedBehavior(child.Behavior)
}

func customerNullValueForType(attributeType attr.Type) attr.Value {
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

func customerLegacyUpgradeIsEmptyValue(value attr.Value) bool {
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

func customerLegacyUpgradeInt64Value(value attr.Value) (int64, bool) {
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

func customerLegacyUpgradeDecimalMirrorsBase(name string, value attr.Value, siblings map[string]attr.Value) bool {
	typedValue, ok := value.(types.String)
	if !ok || typedValue.IsNull() || typedValue.IsUnknown() || !strings.HasSuffix(name, "_decimal") {
		return false
	}
	baseValue, ok := siblings[strings.TrimSuffix(name, "_decimal")]
	if !ok {
		return false
	}
	baseInt, ok := customerLegacyUpgradeInt64Value(baseValue)
	if !ok {
		return false
	}
	return typedValue.ValueString() == strconv.FormatInt(baseInt, 10)
}

func customerLegacyUpgradeZeroAmountMirrorsDecimal(name string, value attr.Value, siblings map[string]attr.Value) bool {
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

func customerLegacyUpgradeNormalizeChild(parent customerStateUpgradeAttrMeta, name string, child customerStateUpgradeAttrMeta, value attr.Value, siblings map[string]attr.Value) attr.Value {
	if !parent.PreserveConfiguredValue || child.Behavior != "optional" {
		return value
	}
	if customerLegacyUpgradeIsEmptyValue(value) {
		return customerNullValueForType(child.AttrType)
	}
	if customerLegacyUpgradeDecimalMirrorsBase(name, value, siblings) {
		return customerNullValueForType(child.AttrType)
	}
	if customerLegacyUpgradeZeroAmountMirrorsDecimal(name, value, siblings) {
		return customerNullValueForType(child.AttrType)
	}
	return value
}

func customerLegacyUpgradeChildAttr(path []string, parent customerStateUpgradeAttrMeta, name string, child customerStateUpgradeAttrMeta, source map[string]attr.Value) attr.Value {
	if !customerShouldPreserveChild(parent, child) {
		return customerNullValueForType(child.AttrType)
	}

	childValue, ok := source[name]
	if !ok {
		return customerNullValueForType(child.AttrType)
	}

	nextPath := append(append([]string{}, path...), name)
	upgradedChild := customerUpgradeValue(nextPath, child, childValue)
	return customerLegacyUpgradeNormalizeChild(parent, name, child, upgradedChild, source)
}

func customerUpgradeAttrs(path []string, meta map[string]customerStateUpgradeAttrMeta, prior map[string]attr.Value) map[string]attr.Value {
	upgraded := make(map[string]attr.Value, len(meta))
	for name, fieldMeta := range meta {
		nextPath := append(append([]string{}, path...), name)
		priorValue, ok := prior[name]
		if !ok {
			upgraded[name] = customerNullValueForType(fieldMeta.AttrType)
			continue
		}
		upgradedValue := customerUpgradeValue(nextPath, fieldMeta, priorValue)
		if fieldMeta.PreserveConfiguredValue && fieldMeta.Behavior == "optional" {
			upgradedValue = customerLegacyUpgradeNormalizeChild(
				customerStateUpgradeAttrMeta{PreserveConfiguredValue: true},
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

func customerUpgradeObjectValue(path []string, meta customerStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue types.Object) attr.Value {
	if priorValue.IsNull() {
		return types.ObjectNull(objectType.AttrTypes)
	}
	if priorValue.IsUnknown() {
		return types.ObjectUnknown(objectType.AttrTypes)
	}
	sourceAttrs := priorValue.Attributes()
	upgradedAttrs := make(map[string]attr.Value, len(meta.Nested))
	for name, childMeta := range meta.Nested {
		upgradedAttrs[name] = customerLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func customerUpgradeSingletonListToObject(path []string, meta customerStateUpgradeAttrMeta, objectType basetypes.ObjectType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return customerNullValueForType(meta.AttrType)
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
		upgradedAttrs[name] = customerLegacyUpgradeChildAttr(path, meta, name, childMeta, sourceAttrs)
	}
	return types.ObjectValueMust(objectType.AttrTypes, upgradedAttrs)
}

func customerUpgradeObjectValueToSingletonList(path []string, meta customerStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	if listValue, ok := priorValue.(types.List); ok {
		return customerUpgradeListValue(path, meta, listType, listValue)
	}
	if baseList, ok := priorValue.(basetypes.ListValue); ok {
		return customerUpgradeListValue(path, meta, listType, types.List(baseList))
	}

	objectValue, ok := priorValue.(types.Object)
	if !ok {
		if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
			objectValue = types.Object(baseObject)
		} else {
			return customerNullValueForType(meta.AttrType)
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
		return customerNullValueForType(meta.AttrType)
	}

	upgradedObject := customerUpgradeObjectValue(path, meta, elementObjectType, objectValue)
	return types.ListValueMust(listType.ElemType, []attr.Value{upgradedObject})
}

func customerUpgradeListValue(path []string, meta customerStateUpgradeAttrMeta, listType basetypes.ListType, priorValue attr.Value) attr.Value {
	listValue, ok := priorValue.(types.List)
	if !ok {
		return customerNullValueForType(meta.AttrType)
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
				upgradedElements = append(upgradedElements, customerNullValueForType(elementObjectType))
				continue
			}
		}
		upgradedElements = append(
			upgradedElements,
			customerUpgradeObjectValue(path, meta, elementObjectType, objectValue),
		)
	}

	return types.ListValueMust(listType.ElemType, upgradedElements)
}

func customerUpgradeValue(path []string, meta customerStateUpgradeAttrMeta, priorValue attr.Value) attr.Value {
	pathKey := strings.Join(path, ".")

	switch attrType := meta.AttrType.(type) {
	case basetypes.ObjectType:
		if _, ok := customerStateUpgradeSingletonPaths[pathKey]; ok {
			return customerUpgradeSingletonListToObject(path, meta, attrType, priorValue)
		}

		objectValue, ok := priorValue.(types.Object)
		if !ok {
			if baseObject, baseOk := priorValue.(basetypes.ObjectValue); baseOk {
				objectValue = types.Object(baseObject)
			} else {
				return customerNullValueForType(meta.AttrType)
			}
		}
		return customerUpgradeObjectValue(path, meta, attrType, objectValue)
	case basetypes.ListType:
		if _, ok := customerStateUpgradeLegacyObjectPaths[pathKey]; ok {
			return customerUpgradeObjectValueToSingletonList(path, meta, attrType, priorValue)
		}
		return customerUpgradeListValue(path, meta, attrType, priorValue)
	default:
		return priorValue
	}
}

func upgradeCustomerStateV1(ctx context.Context, prior interface{}) (CustomerResourceModel, diag.Diagnostics) {
	_ = ctx
	upgradedAttrs := customerUpgradeAttrs(nil, customerStateUpgradeRootMeta, customerAttrMapFromModel(prior))
	var upgraded CustomerResourceModel
	customerSetModelFromAttrMap(&upgraded, upgradedAttrs)
	return upgraded, nil
}

func (r *CustomerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     2,
		Description: "This object represents a customer of your business. Use it to [create recurring charges](https://docs.stripe.com/invoicing/customer), [save payment](https://docs.stripe.com/payments/save-during-payment) and contact information,\nand track payments that belong to the same customer.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("customer")},
			},
			"balance": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The current balance, if any, that's stored on the customer in their default currency. If negative, the customer has credit to apply to their next invoice. If positive, the customer has an amount owed that's added to their next invoice. The balance only considers amounts that Stripe hasn't successfully applied to any invoice. It doesn't reflect unpaid invoices. This balance is only taken into account after invoices finalize. For multi-currency balances, see [invoice_credit_balance](https://docs.stripe.com/api/customers/object#customer_object-invoice_credit_balance).",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"business_name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's business name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Computed:      true,
				Description:   "Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) the customer can be charged in for recurring billing purposes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_account": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of an Account representing a customer. You can use this ID with any v1 API that accepts a customer_account parameter.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_source": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the default payment source for the customer.\n\nIf you use payment methods created through the PaymentMethods API, see the [invoice_settings.default_payment_method](https://docs.stripe.com/api/customers/object#customer_object-invoice_settings-default_payment_method) field instead.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"delinquent": schema.BoolAttribute{
				Computed:      true,
				Description:   "Tracks the most recent state change on any invoice belonging to the customer. Paying an invoice or marking it uncollectible via the API will set this field to false. An automatic payment failure or passing the `invoice.due_date` will set this field to `true`.\n\nIf an invoice becomes uncollectible by [dunning](https://docs.stripe.com/billing/automatic-collection), `delinquent` doesn't reset to `false`.\n\nIf you care whether the customer has paid their most recent subscription invoice, use `subscription.status` instead. Paying or marking uncollectible any customer invoice regardless of whether it is the latest invoice for a subscription will always set this field to `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "An arbitrary string attached to the object. Often useful for displaying to users.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"discount": schema.StringAttribute{
				Computed:      true,
				Description:   "Describes the current discount active on the customer, if there is one.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"email": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's email address.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"individual_name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's individual name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice_credit_balance": schema.MapAttribute{
				Computed:      true,
				Description:   "The current multi-currency balances, if any, that's stored on the customer. If positive in a currency, the customer has a credit to apply to their next invoice denominated in that currency. If negative, the customer has an amount owed that's added to their next invoice denominated in that currency. These balances don't apply to unpaid invoices. They solely track amounts that Stripe hasn't successfully applied to any invoice. Stripe only applies a balance in a specific currency to an invoice after that invoice (which is in the same currency) finalizes.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.Int64Type,
			},
			"invoice_prefix": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The prefix for the customer used to generate unique invoice numbers.",
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
			"name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's full name or business name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"next_invoice_sequence": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The suffix of the customer's next invoice number (for example, 0001). When the account uses account level sequencing, this parameter is ignored in API requests and the field omitted in API responses.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"phone": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's phone number.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"preferred_locales": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The customer's preferred locales (languages), ordered by preference.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"tax_exempt": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Describes the customer's tax exemption status, which is `none`, `exempt`, or `reverse`. When set to `reverse`, invoice and receipt PDFs include the following text: **\"Reverse charge\"**.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("exempt", "none", "reverse")},
			},
			"test_clock": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the test clock that this customer belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"payment_method": schema.StringAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"source": schema.StringAttribute{
				Optional: true,

				WriteOnly: true,
			},
			"validate": schema.BoolAttribute{
				Optional: true,

				WriteOnly: true,
			},
		},
		Blocks: map[string]schema.Block{
			"address": schema.ListNestedBlock{
				Description:   "The customer's address.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
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
			},
			"cash_balance": schema.ListNestedBlock{
				Description:   "The current funds being held by Stripe on behalf of the customer. You can apply these funds towards payment intents when the source is \"cash_balance\". The `settings[reconciliation_mode]` field describes if these funds apply to these payment intents manually or automatically.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"object": schema.StringAttribute{
							Computed:      true,
							Description:   "String representing the object's type. Objects of the same type share the same value.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("cash_balance")},
						},
						"available": schema.MapAttribute{
							Computed:      true,
							Description:   "A hash of all cash balances available to this customer. You cannot delete a customer with any cash balances, even if the balance is 0. Amounts are represented in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).",
							PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
							ElementType:   types.Int64Type,
						},
						"customer": schema.StringAttribute{
							Computed:      true,
							Description:   "The ID of the customer whose cash balance this object represents.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"customer_account": schema.StringAttribute{
							Computed:      true,
							Description:   "The ID of an Account representing a customer whose cash balance this object represents.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"livemode": schema.BoolAttribute{
							Computed:      true,
							Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
							PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
						},
					},
					Blocks: map[string]schema.Block{
						"settings": schema.ListNestedBlock{
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"reconciliation_mode": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The configuration for how funds that land in the customer cash balance are reconciled.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("automatic", "manual")},
									},
									"using_merchant_default": schema.BoolAttribute{
										Computed:      true,
										Description:   "A flag to indicate if reconciliation mode returned is the user's default or is specific to this customer cash balance",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
				},
			},
			"invoice_settings": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"custom_fields": schema.ListNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Default custom fields to be displayed on invoices for this customer.",
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
						"default_payment_method": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "ID of a payment method that's attached to the customer, to be used as the customer's default payment method for subscriptions and invoices.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"footer": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Default footer to be displayed on invoices for this customer.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
					},
					Blocks: map[string]schema.Block{
						"rendering_options": schema.ListNestedBlock{
							Description:   "Default options for invoice PDF rendering for this customer.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"amount_tax_display": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"template": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "ID of the invoice rendering template to be used for this customer's invoices. If set, the template will be used on all invoices for this customer unless a template is set directly on the invoice.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
				},
			},
			"shipping": schema.ListNestedBlock{
				Description:   "Mailing and shipping address for the customer. Appears on invoices emailed to this customer.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
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
					Blocks: map[string]schema.Block{
						"address": schema.ListNestedBlock{
							NestedObject: schema.NestedBlockObject{
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
						},
					},
				},
			},
			"tax": schema.ListNestedBlock{
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"automatic_tax": schema.StringAttribute{
							Computed:      true,
							Description:   "Surfaces if automatic tax computation is possible given the current customer location information.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("failed", "not_collecting", "supported", "unrecognized_location")},
						},
						"ip_address": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "A recent IP address of the customer used for tax reporting and tax location inference.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"provider": schema.StringAttribute{
							Computed:      true,
							Description:   "The tax calculation provider used for location resolution. Defaults to `stripe` when not using a [third-party provider](/tax/third-party-apps).",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("anrok", "avalara", "sphere", "stripe")},
						},
						"validate_location": schema.StringAttribute{
							Optional:    true,
							Description: "A flag that indicates when Stripe should validate the customer tax location. Defaults to `deferred`.",
							Validators:  []validator.String{stringvalidator.OneOf("deferred", "immediately")},
						},
					},
					Blocks: map[string]schema.Block{
						"location": schema.ListNestedBlock{
							Description:   "The identified tax location of the customer.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"country": schema.StringAttribute{
										Computed:      true,
										Description:   "The identified tax country of the customer.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"source": schema.StringAttribute{
										Computed:      true,
										Description:   "The data source used to infer the customer's location.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("billing_address", "ip_address", "payment_method", "shipping_destination")},
									},
									"state": schema.StringAttribute{
										Computed:      true,
										Description:   "The identified tax state, county, province, or region of the customer.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
				},
			},
			"tax_id_data": schema.ListNestedBlock{
				Description:   "The customer's tax IDs.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Required:      true,
							Description:   "Type of the tax ID, one of `ad_nrt`, `ae_trn`, `al_tin`, `am_tin`, `ao_tin`, `ar_cuit`, `au_abn`, `au_arn`, `aw_tin`, `az_tin`, `ba_tin`, `bb_tin`, `bd_bin`, `bf_ifu`, `bg_uic`, `bh_vat`, `bj_ifu`, `bo_tin`, `br_cnpj`, `br_cpf`, `bs_tin`, `by_tin`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `cd_nif`, `ch_uid`, `ch_vat`, `cl_tin`, `cm_niu`, `cn_tin`, `co_nit`, `cr_tin`, `cv_nif`, `de_stn`, `do_rcn`, `ec_ruc`, `eg_tin`, `es_cif`, `et_tin`, `eu_oss_vat`, `eu_vat`, `fo_vat`, `gb_vat`, `ge_vat`, `gi_tin`, `gn_nif`, `hk_br`, `hr_oib`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `it_cf`, `jp_cn`, `jp_rn`, `jp_trn`, `ke_pin`, `kg_tin`, `kh_tin`, `kr_brn`, `kz_bin`, `la_tin`, `li_uid`, `li_vat`, `lk_vat`, `ma_vat`, `md_vat`, `me_pib`, `mk_vat`, `mr_nif`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `ng_tin`, `no_vat`, `no_voec`, `np_pan`, `nz_gst`, `om_vat`, `pe_ruc`, `ph_tin`, `pl_nip`, `py_ruc`, `ro_tin`, `rs_pib`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `sn_ninea`, `sr_fin`, `sv_nit`, `th_vat`, `tj_tin`, `tr_tin`, `tw_vat`, `tz_vat`, `ua_vat`, `ug_tin`, `us_ein`, `uy_ruc`, `uz_tin`, `uz_vat`, `ve_rif`, `vn_tin`, `za_vat`, `zm_tin`, or `zw_tin`",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
						"value": schema.StringAttribute{
							Required:      true,
							Description:   "Value of the tax ID.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
		},
	}
}

func (r *CustomerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CustomerResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config CustomerResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Source"}, []string{"Validate"}})

	params, err := expandCustomerCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Customer create params", err.Error())
		return
	}

	obj, err := r.client.V1Customers.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Customer", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Customers.B, r.client.V1Customers.Key, stripe.FormatURLPath("/v1/customers/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Customer create raw response", err.Error())
		return
	}

	var createdState CustomerResourceModel
	if err := flattenCustomer(obj, &createdState); err != nil {
		resp.Diagnostics.AddError("Error flattening Customer create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&createdState, &config, [][]string{[]string{"Tax", "validate_location"}})
	normalizeUnknownValues(&createdState)

	diffPlan := plan
	diffCreatedState := createdState

	postCreateParams, err := expandCustomerPostCreateUpdate(diffPlan, diffCreatedState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Customer post-create update params", err.Error())
		return
	}

	if paramsHaveValues(postCreateParams) {
		if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
			!createdState.Metadata.IsNull() && !createdState.Metadata.IsUnknown() {
			if !assignMetadataDiffToNamedField(postCreateParams, "Metadata", plan.Metadata, createdState.Metadata) {
				resp.Diagnostics.AddError("Error building Customer update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", postCreateParams))
				return
			}
		}
		obj, err = r.client.V1Customers.Update(ctx, createdState.ID.ValueString(), postCreateParams)
		if err != nil {
			resp.Diagnostics.AddError("Error finalizing Customer after create", err.Error())
			return
		}
		if err := ensureRawResponse(obj, r.client.V1Customers.B, r.client.V1Customers.Key, stripe.FormatURLPath("/v1/customers/%s", obj.ID), nil); err != nil {
			resp.Diagnostics.AddError("Error hydrating Customer post-create update raw response", err.Error())
			return
		}
	}

	if err := flattenCustomer(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Customer create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Tax", "validate_location"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Source"}, []string{"Validate"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState CustomerResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state CustomerResourceModel
	state = priorState

	obj, err := r.client.V1Customers.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Customer", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Customers.B, r.client.V1Customers.Key, stripe.FormatURLPath("/v1/customers/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Customer raw response", err.Error())
		return
	}

	if err := flattenCustomer(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Customer read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"Tax", "validate_location"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"Source"}, []string{"Validate"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CustomerResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config CustomerResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Source"}, []string{"Validate"}})

	var state CustomerResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandCustomerUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Customer update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Customer update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Customers.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Customer", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Customers.B, r.client.V1Customers.Key, stripe.FormatURLPath("/v1/customers/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Customer update raw response", err.Error())
		return
	}

	if err := flattenCustomer(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Customer update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Tax", "validate_location"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Source"}, []string{"Validate"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *CustomerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CustomerResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1Customers.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting Customer", err.Error())
		return
	}
}

func (r *CustomerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandCustomerCreate(plan CustomerResourceModel) (*stripe.CustomerCreateParams, error) {
	params := &stripe.CustomerCreateParams{}

	if !plan.Address.IsNull() && !plan.Address.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Address", plan.Address) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "address", params)
		}
	}
	if !plan.Balance.IsNull() && !plan.Balance.IsUnknown() {
		params.Balance = stripe.Int64(plan.Balance.ValueInt64())
	}
	if !plan.BusinessName.IsNull() && !plan.BusinessName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "BusinessName", "BusinessName", plan.BusinessName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "business_name", params)
		}
	}
	if !plan.CashBalance.IsNull() && !plan.CashBalance.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CashBalance", plan.CashBalance) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "cash_balance", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.Email.IsNull() && !plan.Email.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Email", "Email", plan.Email.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "email", params)
		}
	}
	if !plan.IndividualName.IsNull() && !plan.IndividualName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "IndividualName", "IndividualName", plan.IndividualName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "individual_name", params)
		}
	}
	if !plan.InvoicePrefix.IsNull() && !plan.InvoicePrefix.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "InvoicePrefix", "InvoicePrefix", plan.InvoicePrefix.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "invoice_prefix", params)
		}
	}
	if !plan.InvoiceSettings.IsNull() && !plan.InvoiceSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "InvoiceSettings", plan.InvoiceSettings) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "invoice_settings", params)
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
	if !plan.NextInvoiceSequence.IsNull() && !plan.NextInvoiceSequence.IsUnknown() {
		params.NextInvoiceSequence = stripe.Int64(plan.NextInvoiceSequence.ValueInt64())
	}
	if !plan.Phone.IsNull() && !plan.Phone.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Phone", "Phone", plan.Phone.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "phone", params)
		}
	}
	if !plan.PreferredLocales.IsNull() && !plan.PreferredLocales.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PreferredLocales", plan.PreferredLocales) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "preferred_locales", params)
		}
	}
	if !plan.Shipping.IsNull() && !plan.Shipping.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Shipping", plan.Shipping) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping", params)
		}
	}
	if !plan.Tax.IsNull() && !plan.Tax.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Tax", plan.Tax) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax", params)
		}
	}
	if !plan.TaxExempt.IsNull() && !plan.TaxExempt.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxExempt", "TaxExempt", plan.TaxExempt.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_exempt", params)
		}
	}
	if !plan.TestClock.IsNull() && !plan.TestClock.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TestClockID", "TestClock", plan.TestClock.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "test_clock", params)
		}
	}
	if !plan.PaymentMethod.IsNull() && !plan.PaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethod", "PaymentMethod", plan.PaymentMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method", params)
		}
	}
	if !plan.Source.IsNull() && !plan.Source.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Source", "Source", plan.Source.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "source", params)
		}
	}
	if !plan.TaxIDData.IsNull() && !plan.TaxIDData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TaxIDData", plan.TaxIDData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_id_data", params)
		}
	}
	if !plan.Validate.IsNull() && !plan.Validate.IsUnknown() {
		params.Validate = stripe.Bool(plan.Validate.ValueBool())
	}

	return params, nil
}

func expandCustomerUpdate(plan CustomerResourceModel, state CustomerResourceModel) (*stripe.CustomerUpdateParams, error) {
	params := &stripe.CustomerUpdateParams{}

	if !plan.Address.Equal(state.Address) && !plan.Address.IsNull() && !plan.Address.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Address", plan.Address) {
			if !plan.Address.Equal(state.Address) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "address", params)
			}
		}
	}
	if !plan.Balance.Equal(state.Balance) && !plan.Balance.IsNull() && !plan.Balance.IsUnknown() {
		params.Balance = stripe.Int64(plan.Balance.ValueInt64())
	}
	if !plan.BusinessName.Equal(state.BusinessName) && !plan.BusinessName.IsNull() && !plan.BusinessName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "BusinessName", "BusinessName", plan.BusinessName.ValueString()) {
			if !plan.BusinessName.Equal(state.BusinessName) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "business_name", params)
			}
		}
	}
	if !plan.CashBalance.Equal(state.CashBalance) && !plan.CashBalance.IsNull() && !plan.CashBalance.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CashBalance", plan.CashBalance) {
			if !plan.CashBalance.Equal(state.CashBalance) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "cash_balance", params)
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
	if !plan.Description.Equal(state.Description) && !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			if !plan.Description.Equal(state.Description) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "description", params)
			}
		}
	}
	if !plan.Email.Equal(state.Email) && !plan.Email.IsNull() && !plan.Email.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Email", "Email", plan.Email.ValueString()) {
			if !plan.Email.Equal(state.Email) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "email", params)
			}
		}
	}
	if !plan.IndividualName.Equal(state.IndividualName) && !plan.IndividualName.IsNull() && !plan.IndividualName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "IndividualName", "IndividualName", plan.IndividualName.ValueString()) {
			if !plan.IndividualName.Equal(state.IndividualName) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "individual_name", params)
			}
		}
	}
	if !plan.InvoicePrefix.Equal(state.InvoicePrefix) && !plan.InvoicePrefix.IsNull() && !plan.InvoicePrefix.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "InvoicePrefix", "InvoicePrefix", plan.InvoicePrefix.ValueString()) {
			if !plan.InvoicePrefix.Equal(state.InvoicePrefix) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "invoice_prefix", params)
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
	if !plan.NextInvoiceSequence.Equal(state.NextInvoiceSequence) && !plan.NextInvoiceSequence.IsNull() && !plan.NextInvoiceSequence.IsUnknown() {
		params.NextInvoiceSequence = stripe.Int64(plan.NextInvoiceSequence.ValueInt64())
	}
	if !plan.Phone.Equal(state.Phone) && !plan.Phone.IsNull() && !plan.Phone.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Phone", "Phone", plan.Phone.ValueString()) {
			if !plan.Phone.Equal(state.Phone) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "phone", params)
			}
		}
	}
	if !plan.PreferredLocales.Equal(state.PreferredLocales) && !plan.PreferredLocales.IsNull() && !plan.PreferredLocales.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PreferredLocales", plan.PreferredLocales) {
			if !plan.PreferredLocales.Equal(state.PreferredLocales) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "preferred_locales", params)
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
	if !plan.Tax.Equal(state.Tax) && !plan.Tax.IsNull() && !plan.Tax.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Tax", plan.Tax) {
			if !plan.Tax.Equal(state.Tax) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tax", params)
			}
		}
	}
	if !plan.TaxExempt.Equal(state.TaxExempt) && !plan.TaxExempt.IsNull() && !plan.TaxExempt.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TaxExempt", "TaxExempt", plan.TaxExempt.ValueString()) {
			if !plan.TaxExempt.Equal(state.TaxExempt) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tax_exempt", params)
			}
		}
	}
	if !plan.Source.Equal(state.Source) && !plan.Source.IsNull() && !plan.Source.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Source", "Source", plan.Source.ValueString()) {
			if !plan.Source.Equal(state.Source) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "source", params)
			}
		}
	}
	if !plan.Validate.Equal(state.Validate) && !plan.Validate.IsNull() && !plan.Validate.IsUnknown() {
		params.Validate = stripe.Bool(plan.Validate.ValueBool())
	}

	return params, nil
}

func expandCustomerPostCreateUpdate(plan CustomerResourceModel, state CustomerResourceModel) (*stripe.CustomerUpdateParams, error) {
	params := &stripe.CustomerUpdateParams{}

	if !plan.DefaultSource.Equal(state.DefaultSource) && !plan.DefaultSource.IsNull() && !plan.DefaultSource.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DefaultSourceID", "DefaultSource", plan.DefaultSource.ValueString()) {
			if !plan.DefaultSource.Equal(state.DefaultSource) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "default_source", params)
			}
		}
	}

	return params, nil
}

func flattenCustomer(obj *stripe.Customer, state *CustomerResourceModel) error {
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
		assignedAddress := false
		hadRawAddress := false
		if rawValueAddress, rawOk := plainValueAtPath(raw, "address"); rawOk {
			hadRawAddress = true
			if rawValueAddress != nil {
				sourceAddress := applyConfiguredKeyedListShapes(rawValueAddress, unwrapPlainSingletonList(attrValueToPlain(state.Address)))
				if valueAddress, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceAddress, [][]string{[]string{}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}, "address", "raw response"); err != nil {
					return err
				} else {
					if typedAddress, ok := valueAddress.(types.List); ok {
						state.Address = typedAddress
						assignedAddress = true
					}
				}
			}
		}
		if !assignedAddress {
			if !hasRaw {
				if responseValueAddress, ok := plainFromResponseField(obj, "Address"); ok {
					sourceAddress := applyConfiguredKeyedListShapes(responseValueAddress, unwrapPlainSingletonList(attrValueToPlain(state.Address)))
					if valueAddress, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceAddress, [][]string{[]string{}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}},
						"address",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAddress, ok := valueAddress.(types.List); ok {
							state.Address = typedAddress
							assignedAddress = true
						}
					}
				}
			}
		}
		if !assignedAddress && hadRawAddress {
			if nullAddress, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}); ok {
				if typedAddress, ok := nullAddress.(types.List); ok {
					state.Address = typedAddress
				}
			}
		}
	}
	{
		if rawValueBalance, rawOk := plainValueAtPath(raw, "balance"); rawOk {
			if valueBalance, err := flattenPlainValue(rawValueBalance, types.Int64Type, "balance", "raw response"); err != nil {
				return err
			} else {
				if typedBalance, ok := valueBalance.(types.Int64); ok {
					state.Balance = typedBalance
				}
			}
		} else if !hasRaw {
			if responseValueBalance, ok := plainFromResponseField(obj, "Balance"); ok {
				if valueBalance, err := flattenPlainValue(responseValueBalance, types.Int64Type, "balance", "response struct"); err != nil {
					return err
				} else {
					if typedBalance, ok := valueBalance.(types.Int64); ok {
						state.Balance = typedBalance
					}
				}
			}
		}
	}
	{
		if rawValueBusinessName, rawOk := plainValueAtPath(raw, "business_name"); rawOk {
			if valueBusinessName, err := flattenPlainValue(rawValueBusinessName, types.StringType, "business_name", "raw response"); err != nil {
				return err
			} else {
				if typedBusinessName, ok := valueBusinessName.(types.String); ok {
					state.BusinessName = typedBusinessName
				}
			}
		} else if !hasRaw {
			if responseValueBusinessName, ok := plainFromResponseField(obj, "BusinessName"); ok {
				if valueBusinessName, err := flattenPlainValue(responseValueBusinessName, types.StringType, "business_name", "response struct"); err != nil {
					return err
				} else {
					if typedBusinessName, ok := valueBusinessName.(types.String); ok {
						state.BusinessName = typedBusinessName
					}
				}
			}
		}
	}
	{
		assignedCashBalance := false
		if rawValueCashBalance, rawOk := plainValueAtPath(raw, "cash_balance"); rawOk {
			if rawValueCashBalance != nil {
				sourceCashBalance := applyConfiguredKeyedListShapes(rawValueCashBalance, unwrapPlainSingletonList(attrValueToPlain(state.CashBalance)))
				if !state.CashBalance.IsNull() && !state.CashBalance.IsUnknown() {
					if valueCashBalance, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceCashBalance, [][]string{[]string{}, []string{"settings"}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"object": types.StringType, "available": types.MapType{ElemType: types.Int64Type}, "customer": types.StringType, "customer_account": types.StringType, "livemode": types.BoolType, "settings": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"reconciliation_mode": types.StringType, "using_merchant_default": types.BoolType}}}}}}, "cash_balance", "raw response"); err != nil {
						return err
					} else {
						if typedCashBalance, ok := valueCashBalance.(types.List); ok {
							state.CashBalance = typedCashBalance
							assignedCashBalance = true
						}
					}
				}
			}
		}
		if !assignedCashBalance {
			if !hasRaw {
				if responseValueCashBalance, ok := plainFromResponseField(obj, "CashBalance"); ok {
					sourceCashBalance := applyConfiguredKeyedListShapes(responseValueCashBalance, unwrapPlainSingletonList(attrValueToPlain(state.CashBalance)))
					if !state.CashBalance.IsNull() && !state.CashBalance.IsUnknown() {
						if valueCashBalance, err := flattenPlainValue(
							applyPlainSingletonListShapePaths(sourceCashBalance, [][]string{[]string{}, []string{"settings"}}),
							types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"object": types.StringType, "available": types.MapType{ElemType: types.Int64Type}, "customer": types.StringType, "customer_account": types.StringType, "livemode": types.BoolType, "settings": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"reconciliation_mode": types.StringType, "using_merchant_default": types.BoolType}}}}}},
							"cash_balance",
							"response struct",
						); err != nil {
							return err
						} else {
							if typedCashBalance, ok := valueCashBalance.(types.List); ok {
								state.CashBalance = typedCashBalance
								assignedCashBalance = true
							}
						}
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
		if rawValueDelinquent, rawOk := plainValueAtPath(raw, "delinquent"); rawOk {
			if valueDelinquent, err := flattenPlainValue(rawValueDelinquent, types.BoolType, "delinquent", "raw response"); err != nil {
				return err
			} else {
				if typedDelinquent, ok := valueDelinquent.(types.Bool); ok {
					state.Delinquent = typedDelinquent
				}
			}
		} else if !hasRaw {
			if responseValueDelinquent, ok := plainFromResponseField(obj, "Delinquent"); ok {
				if valueDelinquent, err := flattenPlainValue(responseValueDelinquent, types.BoolType, "delinquent", "response struct"); err != nil {
					return err
				} else {
					if typedDelinquent, ok := valueDelinquent.(types.Bool); ok {
						state.Delinquent = typedDelinquent
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
		if true {
			if rawValueDiscount, rawOk := plainValueAtPath(raw, "discount"); rawOk {
				if typedDiscount, ok := plainToStringIDValue(rawValueDiscount); ok {
					state.Discount = typedDiscount
				}
			} else if !hasRaw {
				if responseValueDiscount, ok := plainFromResponseField(obj, "Discount"); ok {
					if typedDiscount, ok := plainToStringIDValue(responseValueDiscount); ok {
						state.Discount = typedDiscount
					}
				}
			}
		}
	}
	{
		if rawValueEmail, rawOk := plainValueAtPath(raw, "email"); rawOk {
			if valueEmail, err := flattenPlainValue(rawValueEmail, types.StringType, "email", "raw response"); err != nil {
				return err
			} else {
				if typedEmail, ok := valueEmail.(types.String); ok {
					state.Email = typedEmail
				}
			}
		} else if !hasRaw {
			if responseValueEmail, ok := plainFromResponseField(obj, "Email"); ok {
				if valueEmail, err := flattenPlainValue(responseValueEmail, types.StringType, "email", "response struct"); err != nil {
					return err
				} else {
					if typedEmail, ok := valueEmail.(types.String); ok {
						state.Email = typedEmail
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
		if rawValueIndividualName, rawOk := plainValueAtPath(raw, "individual_name"); rawOk {
			if valueIndividualName, err := flattenPlainValue(rawValueIndividualName, types.StringType, "individual_name", "raw response"); err != nil {
				return err
			} else {
				if typedIndividualName, ok := valueIndividualName.(types.String); ok {
					state.IndividualName = typedIndividualName
				}
			}
		} else if !hasRaw {
			if responseValueIndividualName, ok := plainFromResponseField(obj, "IndividualName"); ok {
				if valueIndividualName, err := flattenPlainValue(responseValueIndividualName, types.StringType, "individual_name", "response struct"); err != nil {
					return err
				} else {
					if typedIndividualName, ok := valueIndividualName.(types.String); ok {
						state.IndividualName = typedIndividualName
					}
				}
			}
		}
	}
	{
		if rawValueInvoiceCreditBalance, rawOk := plainValueAtPath(raw, "invoice_credit_balance"); rawOk {
			if valueInvoiceCreditBalance, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueInvoiceCreditBalance, attrValueToPlain(state.InvoiceCreditBalance)), types.MapType{ElemType: types.Int64Type}, "invoice_credit_balance", "raw response"); err != nil {
				return err
			} else {
				if typedInvoiceCreditBalance, ok := valueInvoiceCreditBalance.(types.Map); ok {
					state.InvoiceCreditBalance = typedInvoiceCreditBalance
				}
			}
		} else if !hasRaw {
			if responseValueInvoiceCreditBalance, ok := plainFromResponseField(obj, "InvoiceCreditBalance"); ok {
				if valueInvoiceCreditBalance, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueInvoiceCreditBalance, attrValueToPlain(state.InvoiceCreditBalance)),
					types.MapType{ElemType: types.Int64Type},
					"invoice_credit_balance",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedInvoiceCreditBalance, ok := valueInvoiceCreditBalance.(types.Map); ok {
						state.InvoiceCreditBalance = typedInvoiceCreditBalance
					}
				}
			}
		}
	}
	{
		if rawValueInvoicePrefix, rawOk := plainValueAtPath(raw, "invoice_prefix"); rawOk {
			if valueInvoicePrefix, err := flattenPlainValue(rawValueInvoicePrefix, types.StringType, "invoice_prefix", "raw response"); err != nil {
				return err
			} else {
				if typedInvoicePrefix, ok := valueInvoicePrefix.(types.String); ok {
					state.InvoicePrefix = typedInvoicePrefix
				}
			}
		} else if !hasRaw {
			if responseValueInvoicePrefix, ok := plainFromResponseField(obj, "InvoicePrefix"); ok {
				if valueInvoicePrefix, err := flattenPlainValue(responseValueInvoicePrefix, types.StringType, "invoice_prefix", "response struct"); err != nil {
					return err
				} else {
					if typedInvoicePrefix, ok := valueInvoicePrefix.(types.String); ok {
						state.InvoicePrefix = typedInvoicePrefix
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
				sourceInvoiceSettings := applyConfiguredKeyedListShapes(rawValueInvoiceSettings, unwrapPlainSingletonList(attrValueToPlain(state.InvoiceSettings)))
				if !plainValueIsEmpty(sourceInvoiceSettings) || state.InvoiceSettings.IsUnknown() || !state.InvoiceSettings.IsNull() {
					if valueInvoiceSettings, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceInvoiceSettings, [][]string{[]string{}, []string{"rendering_options"}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"custom_fields": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "default_payment_method": types.StringType, "footer": types.StringType, "rendering_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}}}}, "invoice_settings", "raw response"); err != nil {
						return err
					} else {
						if typedInvoiceSettings, ok := valueInvoiceSettings.(types.List); ok {
							state.InvoiceSettings = typedInvoiceSettings
							assignedInvoiceSettings = true
						}
					}
				}
			}
		}
		if !assignedInvoiceSettings {
			if !hasRaw {
				if responseValueInvoiceSettings, ok := plainFromResponseField(obj, "InvoiceSettings"); ok {
					sourceInvoiceSettings := applyConfiguredKeyedListShapes(responseValueInvoiceSettings, unwrapPlainSingletonList(attrValueToPlain(state.InvoiceSettings)))
					if !plainValueIsEmpty(sourceInvoiceSettings) || state.InvoiceSettings.IsUnknown() || !state.InvoiceSettings.IsNull() {
						if valueInvoiceSettings, err := flattenPlainValue(
							applyPlainSingletonListShapePaths(sourceInvoiceSettings, [][]string{[]string{}, []string{"rendering_options"}}),
							types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"custom_fields": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "default_payment_method": types.StringType, "footer": types.StringType, "rendering_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}}}},
							"invoice_settings",
							"response struct",
						); err != nil {
							return err
						} else {
							if typedInvoiceSettings, ok := valueInvoiceSettings.(types.List); ok {
								state.InvoiceSettings = typedInvoiceSettings
								assignedInvoiceSettings = true
							}
						}
					}
				}
			}
		}
		if !assignedInvoiceSettings && hadRawInvoiceSettings {
			if nullInvoiceSettings, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"custom_fields": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "default_payment_method": types.StringType, "footer": types.StringType, "rendering_options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}}}}); ok {
				if typedInvoiceSettings, ok := nullInvoiceSettings.(types.List); ok {
					state.InvoiceSettings = typedInvoiceSettings
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
		if rawValueNextInvoiceSequence, rawOk := plainValueAtPath(raw, "next_invoice_sequence"); rawOk {
			if valueNextInvoiceSequence, err := flattenPlainValue(rawValueNextInvoiceSequence, types.Int64Type, "next_invoice_sequence", "raw response"); err != nil {
				return err
			} else {
				if typedNextInvoiceSequence, ok := valueNextInvoiceSequence.(types.Int64); ok {
					state.NextInvoiceSequence = typedNextInvoiceSequence
				}
			}
		} else if !hasRaw {
			if responseValueNextInvoiceSequence, ok := plainFromResponseField(obj, "NextInvoiceSequence"); ok {
				if valueNextInvoiceSequence, err := flattenPlainValue(responseValueNextInvoiceSequence, types.Int64Type, "next_invoice_sequence", "response struct"); err != nil {
					return err
				} else {
					if typedNextInvoiceSequence, ok := valueNextInvoiceSequence.(types.Int64); ok {
						state.NextInvoiceSequence = typedNextInvoiceSequence
					}
				}
			}
		}
	}
	{
		if rawValuePhone, rawOk := plainValueAtPath(raw, "phone"); rawOk {
			if valuePhone, err := flattenPlainValue(rawValuePhone, types.StringType, "phone", "raw response"); err != nil {
				return err
			} else {
				if typedPhone, ok := valuePhone.(types.String); ok {
					state.Phone = typedPhone
				}
			}
		} else if !hasRaw {
			if responseValuePhone, ok := plainFromResponseField(obj, "Phone"); ok {
				if valuePhone, err := flattenPlainValue(responseValuePhone, types.StringType, "phone", "response struct"); err != nil {
					return err
				} else {
					if typedPhone, ok := valuePhone.(types.String); ok {
						state.Phone = typedPhone
					}
				}
			}
		}
	}
	{
		if rawValuePreferredLocales, rawOk := plainValueAtPath(raw, "preferred_locales"); rawOk {
			if valuePreferredLocales, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValuePreferredLocales, attrValueToPlain(state.PreferredLocales)), types.ListType{ElemType: types.StringType}, "preferred_locales", "raw response"); err != nil {
				return err
			} else {
				if typedPreferredLocales, ok := valuePreferredLocales.(types.List); ok {
					state.PreferredLocales = typedPreferredLocales
				}
			}
		} else if !hasRaw {
			if responseValuePreferredLocales, ok := plainFromResponseField(obj, "PreferredLocales"); ok {
				if valuePreferredLocales, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValuePreferredLocales, attrValueToPlain(state.PreferredLocales)),
					types.ListType{ElemType: types.StringType},
					"preferred_locales",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedPreferredLocales, ok := valuePreferredLocales.(types.List); ok {
						state.PreferredLocales = typedPreferredLocales
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
				sourceShipping := applyConfiguredKeyedListShapes(rawValueShipping, unwrapPlainSingletonList(attrValueToPlain(state.Shipping)))
				if valueShipping, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceShipping, [][]string{[]string{}, []string{"address"}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}}, "shipping", "raw response"); err != nil {
					return err
				} else {
					if typedShipping, ok := valueShipping.(types.List); ok {
						state.Shipping = typedShipping
						assignedShipping = true
					}
				}
			}
		}
		if !assignedShipping {
			if !hasRaw {
				if responseValueShipping, ok := plainFromResponseField(obj, "Shipping"); ok {
					sourceShipping := applyConfiguredKeyedListShapes(responseValueShipping, unwrapPlainSingletonList(attrValueToPlain(state.Shipping)))
					if valueShipping, err := flattenPlainValue(
						applyPlainSingletonListShapePaths(sourceShipping, [][]string{[]string{}, []string{"address"}}),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}},
						"shipping",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedShipping, ok := valueShipping.(types.List); ok {
							state.Shipping = typedShipping
							assignedShipping = true
						}
					}
				}
			}
		}
		if !assignedShipping && hadRawShipping {
			if nullShipping, ok := nullTerraformValue(types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}}); ok {
				if typedShipping, ok := nullShipping.(types.List); ok {
					state.Shipping = typedShipping
				}
			}
		}
	}
	{
		assignedTax := false
		if rawValueTax, rawOk := plainValueAtPath(raw, "tax"); rawOk {
			if rawValueTax != nil {
				sourceTax := applyConfiguredKeyedListShapes(rawValueTax, unwrapPlainSingletonList(attrValueToPlain(state.Tax)))
				if !state.Tax.IsNull() && !state.Tax.IsUnknown() {
					if valueTax, err := flattenPlainValue(applyPlainSingletonListShapePaths(sourceTax, [][]string{[]string{}, []string{"location"}}), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"automatic_tax": types.StringType, "ip_address": types.StringType, "location": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "source": types.StringType, "state": types.StringType}}}, "provider": types.StringType, "validate_location": types.StringType}}}, "tax", "raw response"); err != nil {
						return err
					} else {
						if typedTax, ok := valueTax.(types.List); ok {
							state.Tax = typedTax
							assignedTax = true
						}
					}
				}
			}
		}
		if !assignedTax {
			if !hasRaw {
				if responseValueTax, ok := plainFromResponseField(obj, "Tax"); ok {
					sourceTax := applyConfiguredKeyedListShapes(responseValueTax, unwrapPlainSingletonList(attrValueToPlain(state.Tax)))
					if !state.Tax.IsNull() && !state.Tax.IsUnknown() {
						if valueTax, err := flattenPlainValue(
							applyPlainSingletonListShapePaths(sourceTax, [][]string{[]string{}, []string{"location"}}),
							types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"automatic_tax": types.StringType, "ip_address": types.StringType, "location": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "source": types.StringType, "state": types.StringType}}}, "provider": types.StringType, "validate_location": types.StringType}}},
							"tax",
							"response struct",
						); err != nil {
							return err
						} else {
							if typedTax, ok := valueTax.(types.List); ok {
								state.Tax = typedTax
								assignedTax = true
							}
						}
					}
				}
			}
		}
	}
	{
		if rawValueTaxExempt, rawOk := plainValueAtPath(raw, "tax_exempt"); rawOk {
			if valueTaxExempt, err := flattenPlainValue(rawValueTaxExempt, types.StringType, "tax_exempt", "raw response"); err != nil {
				return err
			} else {
				if typedTaxExempt, ok := valueTaxExempt.(types.String); ok {
					state.TaxExempt = typedTaxExempt
				}
			}
		} else if !hasRaw {
			if responseValueTaxExempt, ok := plainFromResponseField(obj, "TaxExempt"); ok {
				if valueTaxExempt, err := flattenPlainValue(responseValueTaxExempt, types.StringType, "tax_exempt", "response struct"); err != nil {
					return err
				} else {
					if typedTaxExempt, ok := valueTaxExempt.(types.String); ok {
						state.TaxExempt = typedTaxExempt
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
		if rawValuePaymentMethod, rawOk := plainValueAtPath(raw, "payment_method"); rawOk {
			if !state.PaymentMethod.IsNull() && !state.PaymentMethod.IsUnknown() {
				if valuePaymentMethod, err := flattenPlainValue(rawValuePaymentMethod, types.StringType, "payment_method", "raw response"); err != nil {
					return err
				} else {
					if typedPaymentMethod, ok := valuePaymentMethod.(types.String); ok {
						state.PaymentMethod = typedPaymentMethod
					}
				}
			}
		} else if !hasRaw {
			if responseValuePaymentMethod, ok := plainFromResponseField(obj, "PaymentMethod"); ok {
				if !state.PaymentMethod.IsNull() && !state.PaymentMethod.IsUnknown() {
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
	}
	{
		if rawValueTaxIDData, rawOk := plainValueAtPath(raw, "tax_id_data"); rawOk {
			if !state.TaxIDData.IsNull() && !state.TaxIDData.IsUnknown() {
				if valueTaxIDData, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueTaxIDData, attrValueToPlain(state.TaxIDData)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}}, "tax_id_data", "raw response"); err != nil {
					return err
				} else {
					if typedTaxIDData, ok := valueTaxIDData.(types.List); ok {
						state.TaxIDData = typedTaxIDData
					}
				}
			}
		} else if !hasRaw {
			if responseValueTaxIDData, ok := plainFromResponseField(obj, "TaxIDData"); ok {
				if !state.TaxIDData.IsNull() && !state.TaxIDData.IsUnknown() {
					if valueTaxIDData, err := flattenPlainValue(
						applyConfiguredKeyedListShapes(responseValueTaxIDData, attrValueToPlain(state.TaxIDData)),
						types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}},
						"tax_id_data",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTaxIDData, ok := valueTaxIDData.(types.List); ok {
							state.TaxIDData = typedTaxIDData
						}
					}
				}
			}
		}
	}
	return nil
}
