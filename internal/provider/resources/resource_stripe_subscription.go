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

var _ resource.Resource = &SubscriptionResource{}

var _ resource.ResourceWithConfigure = &SubscriptionResource{}

var _ resource.ResourceWithImportState = &SubscriptionResource{}

func NewSubscriptionResource() resource.Resource {
	return &SubscriptionResource{}
}

type SubscriptionResource struct {
	client *stripe.Client
}

type SubscriptionResourceModel struct {
	Object                        types.String  `tfsdk:"object"`
	Application                   types.String  `tfsdk:"application"`
	ApplicationFeePercent         types.Float64 `tfsdk:"application_fee_percent"`
	AutomaticTax                  types.Object  `tfsdk:"automatic_tax"`
	BillingCycleAnchor            types.Int64   `tfsdk:"billing_cycle_anchor"`
	BillingCycleAnchorConfig      types.Object  `tfsdk:"billing_cycle_anchor_config"`
	BillingMode                   types.Object  `tfsdk:"billing_mode"`
	BillingSchedules              types.List    `tfsdk:"billing_schedules"`
	BillingThresholds             types.Object  `tfsdk:"billing_thresholds"`
	CancelAt                      types.Int64   `tfsdk:"cancel_at"`
	CancelAtPeriodEnd             types.Bool    `tfsdk:"cancel_at_period_end"`
	CanceledAt                    types.Int64   `tfsdk:"canceled_at"`
	CancellationDetails           types.Object  `tfsdk:"cancellation_details"`
	CollectionMethod              types.String  `tfsdk:"collection_method"`
	Created                       types.Int64   `tfsdk:"created"`
	Currency                      types.String  `tfsdk:"currency"`
	Customer                      types.String  `tfsdk:"customer"`
	CustomerAccount               types.String  `tfsdk:"customer_account"`
	DaysUntilDue                  types.Int64   `tfsdk:"days_until_due"`
	DefaultPaymentMethod          types.String  `tfsdk:"default_payment_method"`
	DefaultSource                 types.String  `tfsdk:"default_source"`
	DefaultTaxRates               types.List    `tfsdk:"default_tax_rates"`
	Description                   types.String  `tfsdk:"description"`
	Discounts                     types.List    `tfsdk:"discounts"`
	EndedAt                       types.Int64   `tfsdk:"ended_at"`
	ID                            types.String  `tfsdk:"id"`
	InvoiceSettings               types.Object  `tfsdk:"invoice_settings"`
	Items                         types.List    `tfsdk:"items"`
	Livemode                      types.Bool    `tfsdk:"livemode"`
	ManagedPayments               types.Object  `tfsdk:"managed_payments"`
	Metadata                      types.Map     `tfsdk:"metadata"`
	NextPendingInvoiceItemInvoice types.Int64   `tfsdk:"next_pending_invoice_item_invoice"`
	OnBehalfOf                    types.String  `tfsdk:"on_behalf_of"`
	PauseCollection               types.Object  `tfsdk:"pause_collection"`
	PaymentSettings               types.Object  `tfsdk:"payment_settings"`
	PendingInvoiceItemInterval    types.Object  `tfsdk:"pending_invoice_item_interval"`
	PendingSetupIntent            types.String  `tfsdk:"pending_setup_intent"`
	PendingUpdate                 types.Object  `tfsdk:"pending_update"`
	PresentmentDetails            types.Object  `tfsdk:"presentment_details"`
	Schedule                      types.String  `tfsdk:"schedule"`
	StartDate                     types.Int64   `tfsdk:"start_date"`
	Status                        types.String  `tfsdk:"status"`
	TestClock                     types.String  `tfsdk:"test_clock"`
	TransferData                  types.Object  `tfsdk:"transfer_data"`
	TrialEnd                      types.Int64   `tfsdk:"trial_end"`
	TrialSettings                 types.Object  `tfsdk:"trial_settings"`
	TrialStart                    types.Int64   `tfsdk:"trial_start"`
	AddInvoiceItems               types.List    `tfsdk:"add_invoice_items"`
	BackdateStartDate             types.Int64   `tfsdk:"backdate_start_date"`
	OffSession                    types.Bool    `tfsdk:"off_session"`
	PaymentBehavior               types.String  `tfsdk:"payment_behavior"`
	ProrationBehavior             types.String  `tfsdk:"proration_behavior"`
	TrialFromPlan                 types.Bool    `tfsdk:"trial_from_plan"`
	TrialPeriodDays               types.Int64   `tfsdk:"trial_period_days"`
}

func (r *SubscriptionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SubscriptionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscription"
}

func (r *SubscriptionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Subscriptions allow you to charge a customer on a recurring basis.\n\nRelated guide: [Creating subscriptions](https://docs.stripe.com/billing/subscriptions/creating)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("subscription")},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the Connect Application that created the subscription.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"application_fee_percent": schema.Float64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
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
						Validators:    []validator.String{stringvalidator.OneOf("requires_location_inputs")},
					},
					"enabled": schema.BoolAttribute{
						Required:    true,
						Description: "Whether Stripe automatically computes tax on this subscription.",
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
				},
			},
			"billing_cycle_anchor": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The reference point that aligns future [billing cycle](https://docs.stripe.com/subscriptions/billing-cycle) dates. It sets the day of week for `week` intervals, the day of month for `month` and `year` intervals, and the month of year for `year` intervals. The timestamp is in UTC format.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"billing_cycle_anchor_config": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The fixed values used to calculate the `billing_cycle_anchor`.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"day_of_month": schema.Int64Attribute{
						Required:      true,
						Description:   "The day of the month of the billing_cycle_anchor.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
					"hour": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The hour of the day of the billing_cycle_anchor.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"minute": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The minute of the hour of the billing_cycle_anchor.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"month": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The month to start full cycle billing periods.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"second": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The second of the minute of the billing_cycle_anchor.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
				},
			},
			"billing_mode": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The billing mode of the subscription.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"flexible": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Configure behavior for flexible billing mode",
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
					"updated_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "Details on when the current billing_mode was adopted.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"billing_schedules": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Billing schedules for this subscription.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"applies_to": schema.ListNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Specifies which subscription items the billing schedule applies to.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"price": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The billing schedule will apply to the subscription item with the given price ID.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"type": schema.StringAttribute{
										Required:    true,
										Description: "Controls which subscription items the billing schedule applies to.",
										Validators:  []validator.String{stringvalidator.OneOf("price")},
									},
								},
							},
						},
						"bill_until": schema.SingleNestedAttribute{
							Required:    true,
							Description: "Specifies the end of billing period.",
							Attributes: map[string]schema.Attribute{
								"computed_timestamp": schema.Int64Attribute{
									Computed:      true,
									Description:   "The timestamp the billing schedule will apply until.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"duration": schema.SingleNestedAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "Specifies the billing period.",
									PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
									Attributes: map[string]schema.Attribute{
										"interval": schema.StringAttribute{
											Required:    true,
											Description: "Specifies billing duration. Either `day`, `week`, `month` or `year`.",
											Validators:  []validator.String{stringvalidator.OneOf("day", "month", "week", "year")},
										},
										"interval_count": schema.Int64Attribute{
											Optional:      true,
											Computed:      true,
											Description:   "The multiplier applied to the interval.",
											PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
										},
									},
								},
								"timestamp": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "If specified, the billing schedule will apply until the specified timestamp.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"type": schema.StringAttribute{
									Required:    true,
									Description: "Describes how the billing schedule will determine the end date. Either `duration` or `timestamp`.",
									Validators:  []validator.String{stringvalidator.OneOf("duration", "timestamp")},
								},
							},
						},
						"key": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Unique identifier for the billing schedule.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
					},
				},
			},
			"billing_thresholds": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"amount_gte": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "Monetary threshold that triggers the subscription to create an invoice",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"reset_billing_cycle_anchor": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"cancel_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "A date in the future at which the subscription will automatically get canceled",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"cancel_at_period_end": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether this subscription will (if `status=active`) or did (if `status=canceled`) cancel at the end of the current billing period.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"canceled_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "If the subscription has been canceled, the date of that cancellation. If the subscription was canceled with `cancel_at_period_end`, `canceled_at` will reflect the time of the most recent update request, not the end of the subscription period when the subscription is automatically moved to a canceled state.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"cancellation_details": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Details about why this subscription was cancelled",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"comment": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Additional comments about why the user canceled the subscription, if the subscription was canceled explicitly by the user.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"feedback": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The customer submitted reason for why they canceled, if the subscription was canceled explicitly by the user.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("customer_service", "low_quality", "missing_features", "other", "switched_service", "too_complex", "too_expensive", "unused")},
					},
					"reason": schema.StringAttribute{
						Computed:      true,
						Description:   "Why this subscription was canceled.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("canceled_by_retention_policy", "cancellation_requested", "payment_disputed", "payment_failed")},
					},
				},
			},
			"collection_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this subscription at the end of the cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("charge_automatically", "send_invoice")},
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
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the customer who owns the subscription.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"customer_account": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the account representing the customer who owns the subscription.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"days_until_due": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Number of days a customer has to pay invoices generated by this subscription. This value will be `null` for subscriptions where `collection_method=charge_automatically`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"default_payment_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the default payment method for the subscription. It must belong to the customer associated with the subscription. This takes precedence over `default_source`. If neither are set, invoices will use the customer's [invoice_settings.default_payment_method](https://docs.stripe.com/api/customers/object#customer_object-invoice_settings-default_payment_method) or [default_source](https://docs.stripe.com/api/customers/object#customer_object-default_source).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_source": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the default payment source for the subscription. It must belong to the customer associated with the subscription and be in a chargeable state. If `default_payment_method` is also set, `default_payment_method` will take precedence. If neither are set, invoices will use the customer's [invoice_settings.default_payment_method](https://docs.stripe.com/api/customers/object#customer_object-invoice_settings-default_payment_method) or [default_source](https://docs.stripe.com/api/customers/object#customer_object-default_source).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_tax_rates": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The tax rates that will apply to any subscription item that does not have `tax_rates` set. Invoices created will have their `default_tax_rates` populated from the subscription.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"discounts": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The discounts applied to the subscription. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.",
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
			"ended_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "If the subscription has ended, the date the subscription ended.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice_settings": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"account_tax_ids": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The account tax IDs associated with the subscription. Will be set on invoices generated by the subscription.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
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
			"items": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "List of subscription items, each with an attached price.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:      true,
							Description:   "Unique identifier for the object.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"billing_thresholds": schema.SingleNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period. Pass an empty string to remove previously-defined thresholds.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"usage_gte": schema.Int64Attribute{
									Required:    true,
									Description: "Number of units that meets the billing threshold to advance the subscription to a new billing period (e.g., it takes 10 $5 units to meet a $50 [monetary threshold](https://docs.stripe.com/api/subscriptions/update#update_subscription-billing_thresholds-amount_gte))",
								},
							},
						},
						"discounts": schema.ListNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The coupons to redeem into discounts for the subscription item.",
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
						"metadata": schema.MapAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
							PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
							ElementType:   types.StringType,
						},
						"price": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The ID of the price object.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"price_data": schema.SingleNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline.",
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
									Required:    true,
									Description: "The recurring components of a price such as `interval` and `interval_count`.",
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
							Description:   "Quantity for this item.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"tax_rates": schema.ListAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "A list of [Tax Rate](https://docs.stripe.com/api/tax_rates) ids. These Tax Rates will override the [`default_tax_rates`](https://docs.stripe.com/api/subscriptions/create#create_subscription-default_tax_rates) on the Subscription. When updating, pass an empty string to remove previously-defined tax rates.",
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
			"managed_payments": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Settings for Managed Payments for this Subscription and resulting [Invoices](/api/invoices/object) and [PaymentIntents](/api/payment_intents/object).",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Computed:      true,
						Description:   "Set to `true` to enable [Managed Payments](https://docs.stripe.com/payments/managed-payments), Stripe's merchant of record solution, for this session.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
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
			"next_pending_invoice_item_invoice": schema.Int64Attribute{
				Computed:      true,
				Description:   "Specifies the approximate timestamp on which any pending invoice items will be billed according to the schedule provided at `pending_invoice_item_interval`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"on_behalf_of": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account (if any) the charge was made on behalf of for charges associated with this subscription. See the [Connect documentation](https://docs.stripe.com/connect/subscriptions#on-behalf-of) for details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"pause_collection": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "If specified, payment collection for this subscription will be paused. Note that the subscription status will be unchanged and will not be updated to `paused`. Learn more about [pausing collection](https://docs.stripe.com/billing/subscriptions/pause-payment).",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"behavior": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The payment collection behavior for this subscription while paused.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("keep_as_draft", "mark_uncollectible", "void")},
					},
					"resumes_at": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The time after which the subscription will resume collecting payments.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"payment_settings": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Payment settings passed on to invoices created by the subscription.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"payment_method_options": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Payment-method-specific configuration to provide to invoices created by the subscription.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"acss_debit": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "This sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to invoices created by the subscription.",
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
								Description:   "This sub-hash contains details about the Bancontact payment method options to pass to invoices created by the subscription.",
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
								Description:   "This sub-hash contains details about the Card payment method options to pass to invoices created by the subscription.",
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
												Description:   "Amount to be charged for future payments, specified in the presentment currency.",
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
										},
									},
									"network": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Selected network to process this Subscription on. Depends on the available networks of the card attached to the Subscription. Can be only set confirm-time.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										Validators:    []validator.String{stringvalidator.OneOf("amex", "cartes_bancaires", "diners", "discover", "eftpos_au", "girocard", "interac", "jcb", "link", "mastercard", "unionpay", "unknown", "visa")},
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
								Description:   "This sub-hash contains details about the Bank transfer payment method options to pass to invoices created by the subscription.",
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
								Description:   "This sub-hash contains details about the PayTo payment method options to pass to invoices created by the subscription.",
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
								Description:   "This sub-hash contains details about the Pix payment method options to pass to invoices created by the subscription.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"expires_after_seconds": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "The number of seconds (between 10 and 1209600) after which Pix payment will expire. Defaults to 86400 seconds.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
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
											"amount_includes_iof": schema.StringAttribute{
												Optional:      true,
												Computed:      true,
												Description:   "Determines if the amount includes the IOF tax.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												Validators:    []validator.String{stringvalidator.OneOf("always", "never")},
											},
											"end_date": schema.StringAttribute{
												Optional:      true,
												Computed:      true,
												Description:   "Date when the mandate expires and no further payments will be charged, in `YYYY-MM-DD`.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
											},
											"payment_schedule": schema.StringAttribute{
												Optional:      true,
												Computed:      true,
												Description:   "Schedule at which the future payments will be charged.",
												PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
												Validators:    []validator.String{stringvalidator.OneOf("halfyearly", "monthly", "quarterly", "weekly", "yearly")},
											},
										},
									},
								},
							},
							"upi": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "This sub-hash contains details about the UPI payment method options to pass to invoices created by the subscription.",
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
								Description:   "This sub-hash contains details about the ACH direct debit payment method options to pass to invoices created by the subscription.",
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
						Description:   "The list of payment method types to provide to every invoice created by the subscription. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice’s default payment method, the subscription’s default payment method, the customer’s default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"save_default_payment_method": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Configure whether Stripe updates `subscription.default_payment_method` when payment succeeds. Defaults to `off`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("off", "on_subscription")},
					},
				},
			},
			"pending_invoice_item_interval": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Specifies an interval for how often to bill for any pending invoice items. It is analogous to calling [Create an invoice](/api/invoices/create) for the given subscription at the specified interval.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"interval": schema.StringAttribute{
						Required:    true,
						Description: "Specifies invoicing frequency. Either `day`, `week`, `month` or `year`.",
						Validators:  []validator.String{stringvalidator.OneOf("day", "month", "week", "year")},
					},
					"interval_count": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The number of intervals between invoices. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"pending_setup_intent": schema.StringAttribute{
				Computed:      true,
				Description:   "You can use this [SetupIntent](https://docs.stripe.com/api/setup_intents) to collect user authentication when creating a subscription without immediate payment or updating a subscription's payment method, allowing you to optimize for off-session payments. Learn more in the [SCA Migration Guide](https://docs.stripe.com/billing/migration/strong-customer-authentication#scenario-2).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"pending_update": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "If specified, [pending updates](https://docs.stripe.com/billing/subscriptions/pending-updates) that will be applied to the subscription once the `latest_invoice` has been paid.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"billing_cycle_anchor": schema.Int64Attribute{
						Computed:      true,
						Description:   "If the update is applied, determines the date of the first full invoice, and, for plans with `month` or `year` intervals, the day of the month for subsequent invoices. The timestamp is in UTC format.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"discount": schema.StringAttribute{
						Computed:      true,
						Description:   "The pending subscription-level discount that will be applied when the pending update is applied.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"expires_at": schema.Int64Attribute{
						Computed:      true,
						Description:   "The point after which the changes reflected by this update will be discarded and no longer applied.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"metadata": schema.MapAttribute{
						Computed:      true,
						Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"trial_end": schema.Int64Attribute{
						Computed:      true,
						Description:   "Unix timestamp representing the end of the trial period the customer will get before being charged for the first time, if the update is applied.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"trial_from_plan": schema.BoolAttribute{
						Computed:      true,
						Description:   "Indicates if a plan's `trial_period_days` should be applied to the subscription. Setting `trial_end` per subscription is preferred, and this defaults to `false`. Setting this flag to `true` together with `trial_end` is not allowed. See [Using trial periods on subscriptions](https://docs.stripe.com/billing/subscriptions/trials) to learn more.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"presentment_details": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"presentment_currency": schema.StringAttribute{
						Computed:      true,
						Description:   "Currency used for customer payments.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"schedule": schema.StringAttribute{
				Computed:      true,
				Description:   "The schedule attached to the subscription",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"start_date": schema.Int64Attribute{
				Computed:      true,
				Description:   "Date when the subscription was first created. The date might differ from the `created` date due to backdating.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Possible values are `incomplete`, `incomplete_expired`, `trialing`, `active`, `past_due`, `canceled`, `unpaid`, or `paused`. \n\nFor `collection_method=charge_automatically` a subscription moves into `incomplete` if the initial payment attempt fails. A subscription in this status can only have metadata and default_source updated. Once the first invoice is paid, the subscription moves into an `active` status. If the first invoice is not paid within 23 hours, the subscription transitions to `incomplete_expired`. This is a terminal status, the open invoice will be voided and no further invoices will be generated. \n\nA subscription that is currently in a trial period is `trialing` and moves to `active` when the trial period is over. \n\nA subscription can only enter a `paused` status [when a trial ends without a payment method](https://docs.stripe.com/billing/subscriptions/trials#create-free-trials-without-payment). A `paused` subscription doesn't generate invoices and can be resumed after your customer adds their payment method. The `paused` status is different from [pausing collection](https://docs.stripe.com/billing/subscriptions/pause-payment), which still generates invoices and leaves the subscription's status unchanged. \n\nIf subscription `collection_method=charge_automatically`, it becomes `past_due` when payment is required but cannot be paid (due to failed payment or awaiting additional user actions). Once Stripe has exhausted all payment retry attempts, the subscription will become `canceled` or `unpaid` (depending on your subscriptions settings). \n\nIf subscription `collection_method=send_invoice` it becomes `past_due` when its invoice is not paid by the due date, and `canceled` or `unpaid` if it is still not paid by an additional deadline after that. Note that when a subscription has a status of `unpaid`, no subsequent invoices will be attempted (invoices will be created, but then immediately automatically closed). After receiving updated payment information from a customer, you may choose to reopen and pay their closed invoices.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "canceled", "incomplete", "incomplete_expired", "past_due", "paused", "trialing", "unpaid")},
			},
			"test_clock": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the test clock this subscription belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"transfer_data": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account (if any) the subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"amount_percent": schema.Float64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the destination account. By default, the entire amount is transferred to the destination.",
						PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
					},
					"destination": schema.StringAttribute{
						Required:    true,
						Description: "The account where funds from the payment will be transferred to upon payment success.",
					},
				},
			},
			"trial_end": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "If the subscription has a trial, the end of that trial.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"trial_settings": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Settings related to subscription trials.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"end_behavior": schema.SingleNestedAttribute{
						Required:    true,
						Description: "Defines how a subscription behaves when a trial ends.",
						Attributes: map[string]schema.Attribute{
							"missing_payment_method": schema.StringAttribute{
								Required:    true,
								Description: "Indicates how the subscription should change when the trial ends if the user did not provide a payment method.",
								Validators:  []validator.String{stringvalidator.OneOf("cancel", "create_invoice", "pause")},
							},
						},
					},
				},
			},
			"trial_start": schema.Int64Attribute{
				Computed:      true,
				Description:   "If the subscription has a trial, the beginning of that trial.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"add_invoice_items": schema.ListNestedAttribute{
				Optional:    true,
				Description: "A list of prices and quantities that will generate invoice items appended to the next invoice for this subscription. You may pass up to 20 items.",
				WriteOnly:   true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"discountable": schema.BoolAttribute{
							Optional:    true,
							Description: "Controls whether discounts apply to this invoice item. Defaults to true if no value is provided.",
							WriteOnly:   true,
						},
						"discounts": schema.ListNestedAttribute{
							Optional:    true,
							Description: "The coupons to redeem into discounts for the item.",
							WriteOnly:   true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"coupon": schema.StringAttribute{
										Optional:    true,
										Description: "ID of the coupon to create a new discount for.",
										WriteOnly:   true,
									},
									"discount": schema.StringAttribute{
										Optional:    true,
										Description: "ID of an existing discount on the object (or one of its ancestors) to reuse.",
										WriteOnly:   true,
									},
									"promotion_code": schema.StringAttribute{
										Optional:    true,
										Description: "ID of the promotion code to create a new discount for.",
										WriteOnly:   true,
									},
								},
							},
						},
						"metadata": schema.MapAttribute{
							Optional:    true,
							Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
							WriteOnly:   true,
							ElementType: types.StringType,
						},
						"period": schema.SingleNestedAttribute{
							Optional:    true,
							Description: "The period associated with this invoice item. If not set, `period.start.type` defaults to `max_item_period_start` and `period.end.type` defaults to `min_item_period_end`.",
							WriteOnly:   true,
							Attributes: map[string]schema.Attribute{
								"end": schema.SingleNestedAttribute{
									Required:    true,
									Description: "End of the invoice item period.",
									WriteOnly:   true,
									Attributes: map[string]schema.Attribute{
										"timestamp": schema.Int64Attribute{
											Optional:    true,
											Description: "A precise Unix timestamp for the end of the invoice item period. Must be greater than or equal to `period.start`.",
											WriteOnly:   true,
										},
										"type": schema.StringAttribute{
											Required:    true,
											Description: "Select how to calculate the end of the invoice item period.",
											WriteOnly:   true,
										},
									},
								},
								"start": schema.SingleNestedAttribute{
									Required:    true,
									Description: "Start of the invoice item period.",
									WriteOnly:   true,
									Attributes: map[string]schema.Attribute{
										"timestamp": schema.Int64Attribute{
											Optional:    true,
											Description: "A precise Unix timestamp for the start of the invoice item period. Must be less than or equal to `period.end`.",
											WriteOnly:   true,
										},
										"type": schema.StringAttribute{
											Required:    true,
											Description: "Select how to calculate the start of the invoice item period.",
											WriteOnly:   true,
										},
									},
								},
							},
						},
						"price": schema.StringAttribute{
							Optional:    true,
							Description: "The ID of the price object. One of `price` or `price_data` is required.",
							WriteOnly:   true,
						},
						"price_data": schema.SingleNestedAttribute{
							Optional:    true,
							Description: "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.",
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
									Description: "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge or a negative integer representing the amount to credit to the customer.",
									WriteOnly:   true,
								},
								"unit_amount_decimal": schema.Float64Attribute{
									Optional:    true,
									Description: "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
									WriteOnly:   true,
								},
							},
						},
						"quantity": schema.Int64Attribute{
							Optional:    true,
							Description: "Quantity for this item. Defaults to 1.",
							WriteOnly:   true,
						},
						"tax_rates": schema.ListAttribute{
							Optional:    true,
							Description: "The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.",
							WriteOnly:   true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"backdate_start_date": schema.Int64Attribute{
				Optional:      true,
				Description:   "A past timestamp to backdate the subscription's start date to. If set, the first invoice will contain line items for the timespan between the start date and the current time. Can be combined with trials and the billing cycle anchor.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"off_session": schema.BoolAttribute{
				Optional:    true,
				Description: "Indicates if a customer is on or off-session while an invoice payment is attempted. Defaults to `false` (on-session).",
				WriteOnly:   true,
			},
			"payment_behavior": schema.StringAttribute{
				Optional:    true,
				Description: "Controls how Stripe handles the first invoice when payment is required and `collection_method=charge_automatically`. Subscriptions with `collection_method=send_invoice` are automatically activated regardless of the first Invoice status.",
				WriteOnly:   true,
				Validators:  []validator.String{stringvalidator.OneOf("allow_incomplete", "default_incomplete", "error_if_incomplete", "pending_if_incomplete")},
			},
			"proration_behavior": schema.StringAttribute{
				Optional:    true,
				Description: "Determines how to handle [prorations](https://docs.stripe.com/billing/subscriptions/prorations) resulting from the `billing_cycle_anchor`. If no value is passed, the default is `create_prorations`.",
				WriteOnly:   true,
				Validators:  []validator.String{stringvalidator.OneOf("always_invoice", "create_prorations", "none")},
			},
			"trial_from_plan": schema.BoolAttribute{
				Optional:    true,
				Description: "Indicates if a plan's `trial_period_days` should be applied to the subscription. Setting `trial_end` per subscription is preferred, and this defaults to `false`. Setting this flag to `true` together with `trial_end` is not allowed. See [Using trial periods on subscriptions](https://docs.stripe.com/billing/subscriptions/trials) to learn more.",
				WriteOnly:   true,
			},
			"trial_period_days": schema.Int64Attribute{
				Optional:      true,
				Description:   "Integer representing the number of trial period days before the customer is charged for the first time. This will always overwrite any trials that might apply via a subscribed plan. See [Using trial periods on subscriptions](https://docs.stripe.com/billing/subscriptions/trials) to learn more.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
		},
	}
}

func (r *SubscriptionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SubscriptionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config SubscriptionResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"AddInvoiceItems"}, []string{"AddInvoiceItems", "*", "discountable"}, []string{"AddInvoiceItems", "*", "discounts"}, []string{"AddInvoiceItems", "*", "discounts", "*", "coupon"}, []string{"AddInvoiceItems", "*", "discounts", "*", "discount"}, []string{"AddInvoiceItems", "*", "discounts", "*", "promotion_code"}, []string{"AddInvoiceItems", "*", "metadata"}, []string{"AddInvoiceItems", "*", "period"}, []string{"AddInvoiceItems", "*", "period", "end"}, []string{"AddInvoiceItems", "*", "period", "end", "timestamp"}, []string{"AddInvoiceItems", "*", "period", "end", "type"}, []string{"AddInvoiceItems", "*", "period", "start"}, []string{"AddInvoiceItems", "*", "period", "start", "timestamp"}, []string{"AddInvoiceItems", "*", "period", "start", "type"}, []string{"AddInvoiceItems", "*", "price"}, []string{"AddInvoiceItems", "*", "price_data"}, []string{"AddInvoiceItems", "*", "price_data", "currency"}, []string{"AddInvoiceItems", "*", "price_data", "product"}, []string{"AddInvoiceItems", "*", "price_data", "tax_behavior"}, []string{"AddInvoiceItems", "*", "price_data", "unit_amount"}, []string{"AddInvoiceItems", "*", "price_data", "unit_amount_decimal"}, []string{"AddInvoiceItems", "*", "quantity"}, []string{"AddInvoiceItems", "*", "tax_rates"}, []string{"BackdateStartDate"}, []string{"OffSession"}, []string{"PaymentBehavior"}, []string{"ProrationBehavior"}, []string{"TrialFromPlan"}, []string{"TrialPeriodDays"}})

	params, err := expandSubscriptionCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Subscription create params", err.Error())
		return
	}

	obj, err := r.client.V1Subscriptions.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Subscription", err.Error())
		return
	}

	rawReadParams := &stripe.SubscriptionRetrieveParams{}
	rawReadParams.AddExpand("items")

	if err := ensureRawResponse(obj, r.client.V1Subscriptions.B, r.client.V1Subscriptions.Key, stripe.FormatURLPath("/v1/subscriptions/%s", obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating Subscription create raw response", err.Error())
		return
	}

	var createdState SubscriptionResourceModel
	if err := flattenSubscription(obj, &createdState); err != nil {
		resp.Diagnostics.AddError("Error flattening Subscription create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&createdState, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}, []string{"Items", "*", "billing_thresholds"}, []string{"Items", "*", "billing_thresholds", "usage_gte"}, []string{"Items", "*", "discounts"}, []string{"Items", "*", "discounts", "*", "coupon"}, []string{"Items", "*", "discounts", "*", "discount"}, []string{"Items", "*", "discounts", "*", "promotion_code"}, []string{"Items", "*", "metadata"}, []string{"Items", "*", "price"}, []string{"Items", "*", "price_data"}, []string{"Items", "*", "price_data", "currency"}, []string{"Items", "*", "price_data", "product"}, []string{"Items", "*", "price_data", "recurring"}, []string{"Items", "*", "price_data", "recurring", "interval"}, []string{"Items", "*", "price_data", "recurring", "interval_count"}, []string{"Items", "*", "price_data", "tax_behavior"}, []string{"Items", "*", "price_data", "unit_amount"}, []string{"Items", "*", "price_data", "unit_amount_decimal"}, []string{"Items", "*", "quantity"}, []string{"Items", "*", "tax_rates"}})
	normalizeUnknownValues(&createdState)

	diffPlan := plan
	diffCreatedState := createdState
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"BackdateStartDate"}, []string{"TrialPeriodDays"}})
	clearWriteOnlyPaths(&diffCreatedState, [][]string{[]string{"BackdateStartDate"}, []string{"TrialPeriodDays"}})

	postCreateParams, err := expandSubscriptionPostCreateUpdate(diffPlan, diffCreatedState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Subscription post-create update params", err.Error())
		return
	}

	if paramsHaveValues(postCreateParams) {
		if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
			!createdState.Metadata.IsNull() && !createdState.Metadata.IsUnknown() {
			if !assignMetadataDiffToNamedField(postCreateParams, "Metadata", plan.Metadata, createdState.Metadata) {
				resp.Diagnostics.AddError("Error building Subscription update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", postCreateParams))
				return
			}
		}
		obj, err = r.client.V1Subscriptions.Update(ctx, createdState.ID.ValueString(), postCreateParams)
		if err != nil {
			resp.Diagnostics.AddError("Error finalizing Subscription after create", err.Error())
			return
		}
		rawReadParams := &stripe.SubscriptionRetrieveParams{}
		rawReadParams.AddExpand("items")

		if err := ensureRawResponse(obj, r.client.V1Subscriptions.B, r.client.V1Subscriptions.Key, stripe.FormatURLPath("/v1/subscriptions/%s", obj.ID), rawReadParams); err != nil {
			resp.Diagnostics.AddError("Error hydrating Subscription post-create update raw response", err.Error())
			return
		}
	}

	if err := flattenSubscription(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Subscription create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}, []string{"Items", "*", "billing_thresholds"}, []string{"Items", "*", "billing_thresholds", "usage_gte"}, []string{"Items", "*", "discounts"}, []string{"Items", "*", "discounts", "*", "coupon"}, []string{"Items", "*", "discounts", "*", "discount"}, []string{"Items", "*", "discounts", "*", "promotion_code"}, []string{"Items", "*", "metadata"}, []string{"Items", "*", "price"}, []string{"Items", "*", "price_data"}, []string{"Items", "*", "price_data", "currency"}, []string{"Items", "*", "price_data", "product"}, []string{"Items", "*", "price_data", "recurring"}, []string{"Items", "*", "price_data", "recurring", "interval"}, []string{"Items", "*", "price_data", "recurring", "interval_count"}, []string{"Items", "*", "price_data", "tax_behavior"}, []string{"Items", "*", "price_data", "unit_amount"}, []string{"Items", "*", "price_data", "unit_amount_decimal"}, []string{"Items", "*", "quantity"}, []string{"Items", "*", "tax_rates"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"AddInvoiceItems"}, []string{"AddInvoiceItems", "*", "discountable"}, []string{"AddInvoiceItems", "*", "discounts"}, []string{"AddInvoiceItems", "*", "discounts", "*", "coupon"}, []string{"AddInvoiceItems", "*", "discounts", "*", "discount"}, []string{"AddInvoiceItems", "*", "discounts", "*", "promotion_code"}, []string{"AddInvoiceItems", "*", "metadata"}, []string{"AddInvoiceItems", "*", "period"}, []string{"AddInvoiceItems", "*", "period", "end"}, []string{"AddInvoiceItems", "*", "period", "end", "timestamp"}, []string{"AddInvoiceItems", "*", "period", "end", "type"}, []string{"AddInvoiceItems", "*", "period", "start"}, []string{"AddInvoiceItems", "*", "period", "start", "timestamp"}, []string{"AddInvoiceItems", "*", "period", "start", "type"}, []string{"AddInvoiceItems", "*", "price"}, []string{"AddInvoiceItems", "*", "price_data"}, []string{"AddInvoiceItems", "*", "price_data", "currency"}, []string{"AddInvoiceItems", "*", "price_data", "product"}, []string{"AddInvoiceItems", "*", "price_data", "tax_behavior"}, []string{"AddInvoiceItems", "*", "price_data", "unit_amount"}, []string{"AddInvoiceItems", "*", "price_data", "unit_amount_decimal"}, []string{"AddInvoiceItems", "*", "quantity"}, []string{"AddInvoiceItems", "*", "tax_rates"}, []string{"BackdateStartDate"}, []string{"OffSession"}, []string{"PaymentBehavior"}, []string{"ProrationBehavior"}, []string{"TrialFromPlan"}, []string{"TrialPeriodDays"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SubscriptionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState SubscriptionResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state SubscriptionResourceModel
	state = priorState

	params := &stripe.SubscriptionRetrieveParams{}
	params.AddExpand("items")

	obj, err := r.client.V1Subscriptions.Retrieve(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Subscription", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Subscriptions.B, r.client.V1Subscriptions.Key, stripe.FormatURLPath("/v1/subscriptions/%s", state.ID.ValueString()), params); err != nil {
		resp.Diagnostics.AddError("Error hydrating Subscription raw response", err.Error())
		return
	}

	if err := flattenSubscription(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Subscription read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}, []string{"Items", "*", "billing_thresholds"}, []string{"Items", "*", "billing_thresholds", "usage_gte"}, []string{"Items", "*", "discounts"}, []string{"Items", "*", "discounts", "*", "coupon"}, []string{"Items", "*", "discounts", "*", "discount"}, []string{"Items", "*", "discounts", "*", "promotion_code"}, []string{"Items", "*", "metadata"}, []string{"Items", "*", "price"}, []string{"Items", "*", "price_data"}, []string{"Items", "*", "price_data", "currency"}, []string{"Items", "*", "price_data", "product"}, []string{"Items", "*", "price_data", "recurring"}, []string{"Items", "*", "price_data", "recurring", "interval"}, []string{"Items", "*", "price_data", "recurring", "interval_count"}, []string{"Items", "*", "price_data", "tax_behavior"}, []string{"Items", "*", "price_data", "unit_amount"}, []string{"Items", "*", "price_data", "unit_amount_decimal"}, []string{"Items", "*", "quantity"}, []string{"Items", "*", "tax_rates"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"AddInvoiceItems"}, []string{"AddInvoiceItems", "*", "discountable"}, []string{"AddInvoiceItems", "*", "discounts"}, []string{"AddInvoiceItems", "*", "discounts", "*", "coupon"}, []string{"AddInvoiceItems", "*", "discounts", "*", "discount"}, []string{"AddInvoiceItems", "*", "discounts", "*", "promotion_code"}, []string{"AddInvoiceItems", "*", "metadata"}, []string{"AddInvoiceItems", "*", "period"}, []string{"AddInvoiceItems", "*", "period", "end"}, []string{"AddInvoiceItems", "*", "period", "end", "timestamp"}, []string{"AddInvoiceItems", "*", "period", "end", "type"}, []string{"AddInvoiceItems", "*", "period", "start"}, []string{"AddInvoiceItems", "*", "period", "start", "timestamp"}, []string{"AddInvoiceItems", "*", "period", "start", "type"}, []string{"AddInvoiceItems", "*", "price"}, []string{"AddInvoiceItems", "*", "price_data"}, []string{"AddInvoiceItems", "*", "price_data", "currency"}, []string{"AddInvoiceItems", "*", "price_data", "product"}, []string{"AddInvoiceItems", "*", "price_data", "tax_behavior"}, []string{"AddInvoiceItems", "*", "price_data", "unit_amount"}, []string{"AddInvoiceItems", "*", "price_data", "unit_amount_decimal"}, []string{"AddInvoiceItems", "*", "quantity"}, []string{"AddInvoiceItems", "*", "tax_rates"}, []string{"BackdateStartDate"}, []string{"OffSession"}, []string{"PaymentBehavior"}, []string{"ProrationBehavior"}, []string{"TrialFromPlan"}, []string{"TrialPeriodDays"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *SubscriptionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SubscriptionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config SubscriptionResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"AddInvoiceItems"}, []string{"AddInvoiceItems", "*", "discountable"}, []string{"AddInvoiceItems", "*", "discounts"}, []string{"AddInvoiceItems", "*", "discounts", "*", "coupon"}, []string{"AddInvoiceItems", "*", "discounts", "*", "discount"}, []string{"AddInvoiceItems", "*", "discounts", "*", "promotion_code"}, []string{"AddInvoiceItems", "*", "metadata"}, []string{"AddInvoiceItems", "*", "period"}, []string{"AddInvoiceItems", "*", "period", "end"}, []string{"AddInvoiceItems", "*", "period", "end", "timestamp"}, []string{"AddInvoiceItems", "*", "period", "end", "type"}, []string{"AddInvoiceItems", "*", "period", "start"}, []string{"AddInvoiceItems", "*", "period", "start", "timestamp"}, []string{"AddInvoiceItems", "*", "period", "start", "type"}, []string{"AddInvoiceItems", "*", "price"}, []string{"AddInvoiceItems", "*", "price_data"}, []string{"AddInvoiceItems", "*", "price_data", "currency"}, []string{"AddInvoiceItems", "*", "price_data", "product"}, []string{"AddInvoiceItems", "*", "price_data", "tax_behavior"}, []string{"AddInvoiceItems", "*", "price_data", "unit_amount"}, []string{"AddInvoiceItems", "*", "price_data", "unit_amount_decimal"}, []string{"AddInvoiceItems", "*", "quantity"}, []string{"AddInvoiceItems", "*", "tax_rates"}, []string{"BackdateStartDate"}, []string{"OffSession"}, []string{"PaymentBehavior"}, []string{"ProrationBehavior"}, []string{"TrialFromPlan"}, []string{"TrialPeriodDays"}})

	var state SubscriptionResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"BackdateStartDate"}, []string{"TrialPeriodDays"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"BackdateStartDate"}, []string{"TrialPeriodDays"}})

	params, err := expandSubscriptionUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Subscription update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Subscription update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Subscriptions.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Subscription", err.Error())
		return
	}

	rawReadParams := &stripe.SubscriptionRetrieveParams{}
	rawReadParams.AddExpand("items")

	if err := ensureRawResponse(obj, r.client.V1Subscriptions.B, r.client.V1Subscriptions.Key, stripe.FormatURLPath("/v1/subscriptions/%s", obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating Subscription update raw response", err.Error())
		return
	}

	if err := flattenSubscription(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Subscription update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}, []string{"Items", "*", "billing_thresholds"}, []string{"Items", "*", "billing_thresholds", "usage_gte"}, []string{"Items", "*", "discounts"}, []string{"Items", "*", "discounts", "*", "coupon"}, []string{"Items", "*", "discounts", "*", "discount"}, []string{"Items", "*", "discounts", "*", "promotion_code"}, []string{"Items", "*", "metadata"}, []string{"Items", "*", "price"}, []string{"Items", "*", "price_data"}, []string{"Items", "*", "price_data", "currency"}, []string{"Items", "*", "price_data", "product"}, []string{"Items", "*", "price_data", "recurring"}, []string{"Items", "*", "price_data", "recurring", "interval"}, []string{"Items", "*", "price_data", "recurring", "interval_count"}, []string{"Items", "*", "price_data", "tax_behavior"}, []string{"Items", "*", "price_data", "unit_amount"}, []string{"Items", "*", "price_data", "unit_amount_decimal"}, []string{"Items", "*", "quantity"}, []string{"Items", "*", "tax_rates"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"AddInvoiceItems"}, []string{"AddInvoiceItems", "*", "discountable"}, []string{"AddInvoiceItems", "*", "discounts"}, []string{"AddInvoiceItems", "*", "discounts", "*", "coupon"}, []string{"AddInvoiceItems", "*", "discounts", "*", "discount"}, []string{"AddInvoiceItems", "*", "discounts", "*", "promotion_code"}, []string{"AddInvoiceItems", "*", "metadata"}, []string{"AddInvoiceItems", "*", "period"}, []string{"AddInvoiceItems", "*", "period", "end"}, []string{"AddInvoiceItems", "*", "period", "end", "timestamp"}, []string{"AddInvoiceItems", "*", "period", "end", "type"}, []string{"AddInvoiceItems", "*", "period", "start"}, []string{"AddInvoiceItems", "*", "period", "start", "timestamp"}, []string{"AddInvoiceItems", "*", "period", "start", "type"}, []string{"AddInvoiceItems", "*", "price"}, []string{"AddInvoiceItems", "*", "price_data"}, []string{"AddInvoiceItems", "*", "price_data", "currency"}, []string{"AddInvoiceItems", "*", "price_data", "product"}, []string{"AddInvoiceItems", "*", "price_data", "tax_behavior"}, []string{"AddInvoiceItems", "*", "price_data", "unit_amount"}, []string{"AddInvoiceItems", "*", "price_data", "unit_amount_decimal"}, []string{"AddInvoiceItems", "*", "quantity"}, []string{"AddInvoiceItems", "*", "tax_rates"}, []string{"BackdateStartDate"}, []string{"OffSession"}, []string{"PaymentBehavior"}, []string{"ProrationBehavior"}, []string{"TrialFromPlan"}, []string{"TrialPeriodDays"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SubscriptionResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *SubscriptionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandSubscriptionCreate(plan SubscriptionResourceModel) (*stripe.SubscriptionCreateParams, error) {
	params := &stripe.SubscriptionCreateParams{}

	if !plan.ApplicationFeePercent.IsNull() && !plan.ApplicationFeePercent.IsUnknown() {
		params.ApplicationFeePercent = stripe.Float64(plan.ApplicationFeePercent.ValueFloat64())
	}
	if !plan.AutomaticTax.IsNull() && !plan.AutomaticTax.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AutomaticTax", plan.AutomaticTax) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "automatic_tax", params)
		}
	}
	if !plan.BillingCycleAnchor.IsNull() && !plan.BillingCycleAnchor.IsUnknown() {
		params.BillingCycleAnchor = stripe.Int64(plan.BillingCycleAnchor.ValueInt64())
	}
	if !plan.BillingCycleAnchorConfig.IsNull() && !plan.BillingCycleAnchorConfig.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingCycleAnchorConfig", plan.BillingCycleAnchorConfig) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_cycle_anchor_config", params)
		}
	}
	if !plan.BillingMode.IsNull() && !plan.BillingMode.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingMode", plan.BillingMode) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_mode", params)
		}
	}
	if !plan.BillingSchedules.IsNull() && !plan.BillingSchedules.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingSchedules", plan.BillingSchedules) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_schedules", params)
		}
	}
	if !plan.BillingThresholds.IsNull() && !plan.BillingThresholds.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingThresholds", plan.BillingThresholds) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_thresholds", params)
		}
	}
	if !plan.CancelAt.IsNull() && !plan.CancelAt.IsUnknown() {
		params.CancelAt = stripe.Int64(plan.CancelAt.ValueInt64())
	}
	if !plan.CancelAtPeriodEnd.IsNull() && !plan.CancelAtPeriodEnd.IsUnknown() {
		params.CancelAtPeriodEnd = stripe.Bool(plan.CancelAtPeriodEnd.ValueBool())
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
	if !plan.DaysUntilDue.IsNull() && !plan.DaysUntilDue.IsUnknown() {
		params.DaysUntilDue = stripe.Int64(plan.DaysUntilDue.ValueInt64())
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
	if !plan.InvoiceSettings.IsNull() && !plan.InvoiceSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "InvoiceSettings", plan.InvoiceSettings) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "invoice_settings", params)
		}
	}
	if !plan.Items.IsNull() && !plan.Items.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Items", plan.Items) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "items", params)
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
	if !plan.PaymentSettings.IsNull() && !plan.PaymentSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentSettings", plan.PaymentSettings) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_settings", params)
		}
	}
	if !plan.PendingInvoiceItemInterval.IsNull() && !plan.PendingInvoiceItemInterval.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PendingInvoiceItemInterval", plan.PendingInvoiceItemInterval) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "pending_invoice_item_interval", params)
		}
	}
	if !plan.TransferData.IsNull() && !plan.TransferData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TransferData", plan.TransferData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "transfer_data", params)
		}
	}
	if !plan.TrialEnd.IsNull() && !plan.TrialEnd.IsUnknown() {
		params.TrialEnd = stripe.Int64(plan.TrialEnd.ValueInt64())
	}
	if !plan.TrialSettings.IsNull() && !plan.TrialSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TrialSettings", plan.TrialSettings) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "trial_settings", params)
		}
	}
	if !plan.AddInvoiceItems.IsNull() && !plan.AddInvoiceItems.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AddInvoiceItems", plan.AddInvoiceItems) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "add_invoice_items", params)
		}
	}
	if !plan.BackdateStartDate.IsNull() && !plan.BackdateStartDate.IsUnknown() {
		params.BackdateStartDate = stripe.Int64(plan.BackdateStartDate.ValueInt64())
	}
	if !plan.OffSession.IsNull() && !plan.OffSession.IsUnknown() {
		params.OffSession = stripe.Bool(plan.OffSession.ValueBool())
	}
	if !plan.PaymentBehavior.IsNull() && !plan.PaymentBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentBehavior", "PaymentBehavior", plan.PaymentBehavior.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_behavior", params)
		}
	}
	if !plan.ProrationBehavior.IsNull() && !plan.ProrationBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ProrationBehavior", "ProrationBehavior", plan.ProrationBehavior.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "proration_behavior", params)
		}
	}
	if !plan.TrialFromPlan.IsNull() && !plan.TrialFromPlan.IsUnknown() {
		params.TrialFromPlan = stripe.Bool(plan.TrialFromPlan.ValueBool())
	}
	if !plan.TrialPeriodDays.IsNull() && !plan.TrialPeriodDays.IsUnknown() {
		params.TrialPeriodDays = stripe.Int64(plan.TrialPeriodDays.ValueInt64())
	}

	return params, nil
}

func expandSubscriptionUpdate(plan SubscriptionResourceModel, state SubscriptionResourceModel) (*stripe.SubscriptionUpdateParams, error) {
	params := &stripe.SubscriptionUpdateParams{}

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
	if !plan.BillingCycleAnchor.Equal(state.BillingCycleAnchor) && !plan.BillingCycleAnchor.IsNull() && !plan.BillingCycleAnchor.IsUnknown() {
		params.BillingCycleAnchor = stripe.Int64(plan.BillingCycleAnchor.ValueInt64())
	}
	if !plan.BillingSchedules.Equal(state.BillingSchedules) && !plan.BillingSchedules.IsNull() && !plan.BillingSchedules.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingSchedules", plan.BillingSchedules) {
			if !plan.BillingSchedules.Equal(state.BillingSchedules) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "billing_schedules", params)
			}
		}
	}
	if !plan.BillingThresholds.Equal(state.BillingThresholds) && !plan.BillingThresholds.IsNull() && !plan.BillingThresholds.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingThresholds", plan.BillingThresholds) {
			if !plan.BillingThresholds.Equal(state.BillingThresholds) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "billing_thresholds", params)
			}
		}
	}
	if !plan.CancelAt.Equal(state.CancelAt) && !plan.CancelAt.IsNull() && !plan.CancelAt.IsUnknown() {
		params.CancelAt = stripe.Int64(plan.CancelAt.ValueInt64())
	}
	if !plan.CancelAtPeriodEnd.Equal(state.CancelAtPeriodEnd) && !plan.CancelAtPeriodEnd.IsNull() && !plan.CancelAtPeriodEnd.IsUnknown() {
		params.CancelAtPeriodEnd = stripe.Bool(plan.CancelAtPeriodEnd.ValueBool())
	}
	if !plan.CancellationDetails.Equal(state.CancellationDetails) && !plan.CancellationDetails.IsNull() && !plan.CancellationDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CancellationDetails", plan.CancellationDetails) {
			if !plan.CancellationDetails.Equal(state.CancellationDetails) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "cancellation_details", params)
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
	if !plan.DaysUntilDue.Equal(state.DaysUntilDue) && !plan.DaysUntilDue.IsNull() && !plan.DaysUntilDue.IsUnknown() {
		params.DaysUntilDue = stripe.Int64(plan.DaysUntilDue.ValueInt64())
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
	if !plan.InvoiceSettings.Equal(state.InvoiceSettings) && !plan.InvoiceSettings.IsNull() && !plan.InvoiceSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "InvoiceSettings", plan.InvoiceSettings) {
			if !plan.InvoiceSettings.Equal(state.InvoiceSettings) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "invoice_settings", params)
			}
		}
	}
	if !plan.Items.Equal(state.Items) && !plan.Items.IsNull() && !plan.Items.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Items", plan.Items) {
			if !plan.Items.Equal(state.Items) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "items", params)
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
	if !plan.PauseCollection.Equal(state.PauseCollection) && !plan.PauseCollection.IsNull() && !plan.PauseCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PauseCollection", plan.PauseCollection) {
			if !plan.PauseCollection.Equal(state.PauseCollection) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "pause_collection", params)
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
	if !plan.PendingInvoiceItemInterval.Equal(state.PendingInvoiceItemInterval) && !plan.PendingInvoiceItemInterval.IsNull() && !plan.PendingInvoiceItemInterval.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PendingInvoiceItemInterval", plan.PendingInvoiceItemInterval) {
			if !plan.PendingInvoiceItemInterval.Equal(state.PendingInvoiceItemInterval) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "pending_invoice_item_interval", params)
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
	if !plan.TrialEnd.Equal(state.TrialEnd) && !plan.TrialEnd.IsNull() && !plan.TrialEnd.IsUnknown() {
		params.TrialEnd = stripe.Int64(plan.TrialEnd.ValueInt64())
	}
	if !plan.TrialSettings.Equal(state.TrialSettings) && !plan.TrialSettings.IsNull() && !plan.TrialSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TrialSettings", plan.TrialSettings) {
			if !plan.TrialSettings.Equal(state.TrialSettings) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "trial_settings", params)
			}
		}
	}
	if !plan.AddInvoiceItems.Equal(state.AddInvoiceItems) && !plan.AddInvoiceItems.IsNull() && !plan.AddInvoiceItems.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AddInvoiceItems", plan.AddInvoiceItems) {
			if !plan.AddInvoiceItems.Equal(state.AddInvoiceItems) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "add_invoice_items", params)
			}
		}
	}
	if !plan.OffSession.Equal(state.OffSession) && !plan.OffSession.IsNull() && !plan.OffSession.IsUnknown() {
		params.OffSession = stripe.Bool(plan.OffSession.ValueBool())
	}
	if !plan.PaymentBehavior.Equal(state.PaymentBehavior) && !plan.PaymentBehavior.IsNull() && !plan.PaymentBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentBehavior", "PaymentBehavior", plan.PaymentBehavior.ValueString()) {
			if !plan.PaymentBehavior.Equal(state.PaymentBehavior) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_behavior", params)
			}
		}
	}
	if !plan.ProrationBehavior.Equal(state.ProrationBehavior) && !plan.ProrationBehavior.IsNull() && !plan.ProrationBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ProrationBehavior", "ProrationBehavior", plan.ProrationBehavior.ValueString()) {
			if !plan.ProrationBehavior.Equal(state.ProrationBehavior) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "proration_behavior", params)
			}
		}
	}
	if !plan.TrialFromPlan.Equal(state.TrialFromPlan) && !plan.TrialFromPlan.IsNull() && !plan.TrialFromPlan.IsUnknown() {
		params.TrialFromPlan = stripe.Bool(plan.TrialFromPlan.ValueBool())
	}

	return params, nil
}

func expandSubscriptionPostCreateUpdate(plan SubscriptionResourceModel, state SubscriptionResourceModel) (*stripe.SubscriptionUpdateParams, error) {
	params := &stripe.SubscriptionUpdateParams{}

	if !plan.CancellationDetails.Equal(state.CancellationDetails) && !plan.CancellationDetails.IsNull() && !plan.CancellationDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CancellationDetails", plan.CancellationDetails) {
			if !plan.CancellationDetails.Equal(state.CancellationDetails) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "cancellation_details", params)
			}
		}
	}
	if !plan.PauseCollection.Equal(state.PauseCollection) && !plan.PauseCollection.IsNull() && !plan.PauseCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PauseCollection", plan.PauseCollection) {
			if !plan.PauseCollection.Equal(state.PauseCollection) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "pause_collection", params)
			}
		}
	}

	return params, nil
}

func flattenSubscription(obj *stripe.Subscription, state *SubscriptionResourceModel) error {
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
				if valueAutomaticTax, err := flattenPlainValue(sourceAutomaticTax, types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "automatic_tax", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}},
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
			if nullAutomaticTax, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}); ok {
				if typedAutomaticTax, ok := nullAutomaticTax.(types.Object); ok {
					state.AutomaticTax = typedAutomaticTax
				}
			}
		}
	}
	{
		if rawValueBillingCycleAnchor, rawOk := plainValueAtPath(raw, "billing_cycle_anchor"); rawOk {
			if valueBillingCycleAnchor, err := flattenPlainValue(rawValueBillingCycleAnchor, types.Int64Type, "billing_cycle_anchor", "raw response"); err != nil {
				return err
			} else {
				if typedBillingCycleAnchor, ok := valueBillingCycleAnchor.(types.Int64); ok {
					state.BillingCycleAnchor = typedBillingCycleAnchor
				}
			}
		} else if !hasRaw {
			if responseValueBillingCycleAnchor, ok := plainFromResponseField(obj, "BillingCycleAnchor"); ok {
				if valueBillingCycleAnchor, err := flattenPlainValue(responseValueBillingCycleAnchor, types.Int64Type, "billing_cycle_anchor", "response struct"); err != nil {
					return err
				} else {
					if typedBillingCycleAnchor, ok := valueBillingCycleAnchor.(types.Int64); ok {
						state.BillingCycleAnchor = typedBillingCycleAnchor
					}
				}
			}
		}
	}
	{
		assignedBillingCycleAnchorConfig := false
		hadRawBillingCycleAnchorConfig := false
		if rawValueBillingCycleAnchorConfig, rawOk := plainValueAtPath(raw, "billing_cycle_anchor_config"); rawOk {
			hadRawBillingCycleAnchorConfig = true
			if rawValueBillingCycleAnchorConfig != nil {
				sourceBillingCycleAnchorConfig := applyConfiguredKeyedListShapes(rawValueBillingCycleAnchorConfig, attrValueToPlain(state.BillingCycleAnchorConfig))
				if valueBillingCycleAnchorConfig, err := flattenPlainValue(sourceBillingCycleAnchorConfig, types.ObjectType{AttrTypes: map[string]attr.Type{"day_of_month": types.Int64Type, "hour": types.Int64Type, "minute": types.Int64Type, "month": types.Int64Type, "second": types.Int64Type}}, "billing_cycle_anchor_config", "raw response"); err != nil {
					return err
				} else {
					if typedBillingCycleAnchorConfig, ok := valueBillingCycleAnchorConfig.(types.Object); ok {
						state.BillingCycleAnchorConfig = typedBillingCycleAnchorConfig
						assignedBillingCycleAnchorConfig = true
					}
				}
			}
		}
		if !assignedBillingCycleAnchorConfig {
			if !hasRaw {
				if responseValueBillingCycleAnchorConfig, ok := plainFromResponseField(obj, "BillingCycleAnchorConfig"); ok {
					sourceBillingCycleAnchorConfig := applyConfiguredKeyedListShapes(responseValueBillingCycleAnchorConfig, attrValueToPlain(state.BillingCycleAnchorConfig))
					if valueBillingCycleAnchorConfig, err := flattenPlainValue(
						sourceBillingCycleAnchorConfig,
						types.ObjectType{AttrTypes: map[string]attr.Type{"day_of_month": types.Int64Type, "hour": types.Int64Type, "minute": types.Int64Type, "month": types.Int64Type, "second": types.Int64Type}},
						"billing_cycle_anchor_config",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBillingCycleAnchorConfig, ok := valueBillingCycleAnchorConfig.(types.Object); ok {
							state.BillingCycleAnchorConfig = typedBillingCycleAnchorConfig
							assignedBillingCycleAnchorConfig = true
						}
					}
				}
			}
		}
		if !assignedBillingCycleAnchorConfig && hadRawBillingCycleAnchorConfig {
			if nullBillingCycleAnchorConfig, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"day_of_month": types.Int64Type, "hour": types.Int64Type, "minute": types.Int64Type, "month": types.Int64Type, "second": types.Int64Type}}); ok {
				if typedBillingCycleAnchorConfig, ok := nullBillingCycleAnchorConfig.(types.Object); ok {
					state.BillingCycleAnchorConfig = typedBillingCycleAnchorConfig
				}
			}
		}
	}
	{
		assignedBillingMode := false
		hadRawBillingMode := false
		if rawValueBillingMode, rawOk := plainValueAtPath(raw, "billing_mode"); rawOk {
			hadRawBillingMode = true
			if rawValueBillingMode != nil {
				sourceBillingMode := applyConfiguredKeyedListShapes(rawValueBillingMode, attrValueToPlain(state.BillingMode))
				if valueBillingMode, err := flattenPlainValue(sourceBillingMode, types.ObjectType{AttrTypes: map[string]attr.Type{"flexible": types.ObjectType{AttrTypes: map[string]attr.Type{"proration_discounts": types.StringType}}, "type": types.StringType, "updated_at": types.Int64Type}}, "billing_mode", "raw response"); err != nil {
					return err
				} else {
					if typedBillingMode, ok := valueBillingMode.(types.Object); ok {
						state.BillingMode = typedBillingMode
						assignedBillingMode = true
					}
				}
			}
		}
		if !assignedBillingMode {
			if !hasRaw {
				if responseValueBillingMode, ok := plainFromResponseField(obj, "BillingMode"); ok {
					sourceBillingMode := applyConfiguredKeyedListShapes(responseValueBillingMode, attrValueToPlain(state.BillingMode))
					if valueBillingMode, err := flattenPlainValue(
						sourceBillingMode,
						types.ObjectType{AttrTypes: map[string]attr.Type{"flexible": types.ObjectType{AttrTypes: map[string]attr.Type{"proration_discounts": types.StringType}}, "type": types.StringType, "updated_at": types.Int64Type}},
						"billing_mode",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBillingMode, ok := valueBillingMode.(types.Object); ok {
							state.BillingMode = typedBillingMode
							assignedBillingMode = true
						}
					}
				}
			}
		}
		if !assignedBillingMode && hadRawBillingMode {
			if nullBillingMode, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"flexible": types.ObjectType{AttrTypes: map[string]attr.Type{"proration_discounts": types.StringType}}, "type": types.StringType, "updated_at": types.Int64Type}}); ok {
				if typedBillingMode, ok := nullBillingMode.(types.Object); ok {
					state.BillingMode = typedBillingMode
				}
			}
		}
	}
	{
		if rawValueBillingSchedules, rawOk := plainValueAtPath(raw, "billing_schedules"); rawOk {
			if valueBillingSchedules, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueBillingSchedules, attrValueToPlain(state.BillingSchedules)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"applies_to": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"price": types.StringType, "type": types.StringType}}}, "bill_until": types.ObjectType{AttrTypes: map[string]attr.Type{"computed_timestamp": types.Int64Type, "duration": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "timestamp": types.Int64Type, "type": types.StringType}}, "key": types.StringType}}}, "billing_schedules", "raw response"); err != nil {
				return err
			} else {
				if typedBillingSchedules, ok := valueBillingSchedules.(types.List); ok {
					state.BillingSchedules = typedBillingSchedules
				}
			}
		} else if !hasRaw {
			if responseValueBillingSchedules, ok := plainFromResponseField(obj, "BillingSchedules"); ok {
				if valueBillingSchedules, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueBillingSchedules, attrValueToPlain(state.BillingSchedules)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"applies_to": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"price": types.StringType, "type": types.StringType}}}, "bill_until": types.ObjectType{AttrTypes: map[string]attr.Type{"computed_timestamp": types.Int64Type, "duration": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "timestamp": types.Int64Type, "type": types.StringType}}, "key": types.StringType}}},
					"billing_schedules",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedBillingSchedules, ok := valueBillingSchedules.(types.List); ok {
						state.BillingSchedules = typedBillingSchedules
					}
				}
			}
		}
	}
	{
		assignedBillingThresholds := false
		hadRawBillingThresholds := false
		if rawValueBillingThresholds, rawOk := plainValueAtPath(raw, "billing_thresholds"); rawOk {
			hadRawBillingThresholds = true
			if rawValueBillingThresholds != nil {
				sourceBillingThresholds := applyConfiguredKeyedListShapes(rawValueBillingThresholds, attrValueToPlain(state.BillingThresholds))
				if valueBillingThresholds, err := flattenPlainValue(sourceBillingThresholds, types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "reset_billing_cycle_anchor": types.BoolType}}, "billing_thresholds", "raw response"); err != nil {
					return err
				} else {
					if typedBillingThresholds, ok := valueBillingThresholds.(types.Object); ok {
						state.BillingThresholds = typedBillingThresholds
						assignedBillingThresholds = true
					}
				}
			}
		}
		if !assignedBillingThresholds {
			if !hasRaw {
				if responseValueBillingThresholds, ok := plainFromResponseField(obj, "BillingThresholds"); ok {
					sourceBillingThresholds := applyConfiguredKeyedListShapes(responseValueBillingThresholds, attrValueToPlain(state.BillingThresholds))
					if valueBillingThresholds, err := flattenPlainValue(
						sourceBillingThresholds,
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "reset_billing_cycle_anchor": types.BoolType}},
						"billing_thresholds",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBillingThresholds, ok := valueBillingThresholds.(types.Object); ok {
							state.BillingThresholds = typedBillingThresholds
							assignedBillingThresholds = true
						}
					}
				}
			}
		}
		if !assignedBillingThresholds && hadRawBillingThresholds {
			if nullBillingThresholds, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "reset_billing_cycle_anchor": types.BoolType}}); ok {
				if typedBillingThresholds, ok := nullBillingThresholds.(types.Object); ok {
					state.BillingThresholds = typedBillingThresholds
				}
			}
		}
	}
	{
		if rawValueCancelAt, rawOk := plainValueAtPath(raw, "cancel_at"); rawOk {
			if valueCancelAt, err := flattenPlainValue(rawValueCancelAt, types.Int64Type, "cancel_at", "raw response"); err != nil {
				return err
			} else {
				if typedCancelAt, ok := valueCancelAt.(types.Int64); ok {
					state.CancelAt = typedCancelAt
				}
			}
		} else if !hasRaw {
			if responseValueCancelAt, ok := plainFromResponseField(obj, "CancelAt"); ok {
				if valueCancelAt, err := flattenPlainValue(responseValueCancelAt, types.Int64Type, "cancel_at", "response struct"); err != nil {
					return err
				} else {
					if typedCancelAt, ok := valueCancelAt.(types.Int64); ok {
						state.CancelAt = typedCancelAt
					}
				}
			}
		}
	}
	{
		if rawValueCancelAtPeriodEnd, rawOk := plainValueAtPath(raw, "cancel_at_period_end"); rawOk {
			if valueCancelAtPeriodEnd, err := flattenPlainValue(rawValueCancelAtPeriodEnd, types.BoolType, "cancel_at_period_end", "raw response"); err != nil {
				return err
			} else {
				if typedCancelAtPeriodEnd, ok := valueCancelAtPeriodEnd.(types.Bool); ok {
					state.CancelAtPeriodEnd = typedCancelAtPeriodEnd
				}
			}
		} else if !hasRaw {
			if responseValueCancelAtPeriodEnd, ok := plainFromResponseField(obj, "CancelAtPeriodEnd"); ok {
				if valueCancelAtPeriodEnd, err := flattenPlainValue(responseValueCancelAtPeriodEnd, types.BoolType, "cancel_at_period_end", "response struct"); err != nil {
					return err
				} else {
					if typedCancelAtPeriodEnd, ok := valueCancelAtPeriodEnd.(types.Bool); ok {
						state.CancelAtPeriodEnd = typedCancelAtPeriodEnd
					}
				}
			}
		}
	}
	{
		if rawValueCanceledAt, rawOk := plainValueAtPath(raw, "canceled_at"); rawOk {
			if valueCanceledAt, err := flattenPlainValue(rawValueCanceledAt, types.Int64Type, "canceled_at", "raw response"); err != nil {
				return err
			} else {
				if typedCanceledAt, ok := valueCanceledAt.(types.Int64); ok {
					state.CanceledAt = typedCanceledAt
				}
			}
		} else if !hasRaw {
			if responseValueCanceledAt, ok := plainFromResponseField(obj, "CanceledAt"); ok {
				if valueCanceledAt, err := flattenPlainValue(responseValueCanceledAt, types.Int64Type, "canceled_at", "response struct"); err != nil {
					return err
				} else {
					if typedCanceledAt, ok := valueCanceledAt.(types.Int64); ok {
						state.CanceledAt = typedCanceledAt
					}
				}
			}
		}
	}
	{
		assignedCancellationDetails := false
		hadRawCancellationDetails := false
		if rawValueCancellationDetails, rawOk := plainValueAtPath(raw, "cancellation_details"); rawOk {
			hadRawCancellationDetails = true
			if rawValueCancellationDetails != nil {
				sourceCancellationDetails := applyConfiguredKeyedListShapes(rawValueCancellationDetails, attrValueToPlain(state.CancellationDetails))
				if valueCancellationDetails, err := flattenPlainValue(sourceCancellationDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"comment": types.StringType, "feedback": types.StringType, "reason": types.StringType}}, "cancellation_details", "raw response"); err != nil {
					return err
				} else {
					if typedCancellationDetails, ok := valueCancellationDetails.(types.Object); ok {
						state.CancellationDetails = typedCancellationDetails
						assignedCancellationDetails = true
					}
				}
			}
		}
		if !assignedCancellationDetails {
			if !hasRaw {
				if responseValueCancellationDetails, ok := plainFromResponseField(obj, "CancellationDetails"); ok {
					sourceCancellationDetails := applyConfiguredKeyedListShapes(responseValueCancellationDetails, attrValueToPlain(state.CancellationDetails))
					if valueCancellationDetails, err := flattenPlainValue(
						sourceCancellationDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"comment": types.StringType, "feedback": types.StringType, "reason": types.StringType}},
						"cancellation_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCancellationDetails, ok := valueCancellationDetails.(types.Object); ok {
							state.CancellationDetails = typedCancellationDetails
							assignedCancellationDetails = true
						}
					}
				}
			}
		}
		if !assignedCancellationDetails && hadRawCancellationDetails {
			if nullCancellationDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"comment": types.StringType, "feedback": types.StringType, "reason": types.StringType}}); ok {
				if typedCancellationDetails, ok := nullCancellationDetails.(types.Object); ok {
					state.CancellationDetails = typedCancellationDetails
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
		if rawValueDaysUntilDue, rawOk := plainValueAtPath(raw, "days_until_due"); rawOk {
			if valueDaysUntilDue, err := flattenPlainValue(rawValueDaysUntilDue, types.Int64Type, "days_until_due", "raw response"); err != nil {
				return err
			} else {
				if typedDaysUntilDue, ok := valueDaysUntilDue.(types.Int64); ok {
					state.DaysUntilDue = typedDaysUntilDue
				}
			}
		} else if !hasRaw {
			if responseValueDaysUntilDue, ok := plainFromResponseField(obj, "DaysUntilDue"); ok {
				if valueDaysUntilDue, err := flattenPlainValue(responseValueDaysUntilDue, types.Int64Type, "days_until_due", "response struct"); err != nil {
					return err
				} else {
					if typedDaysUntilDue, ok := valueDaysUntilDue.(types.Int64); ok {
						state.DaysUntilDue = typedDaysUntilDue
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
		if rawValueEndedAt, rawOk := plainValueAtPath(raw, "ended_at"); rawOk {
			if valueEndedAt, err := flattenPlainValue(rawValueEndedAt, types.Int64Type, "ended_at", "raw response"); err != nil {
				return err
			} else {
				if typedEndedAt, ok := valueEndedAt.(types.Int64); ok {
					state.EndedAt = typedEndedAt
				}
			}
		} else if !hasRaw {
			if responseValueEndedAt, ok := plainFromResponseField(obj, "EndedAt"); ok {
				if valueEndedAt, err := flattenPlainValue(responseValueEndedAt, types.Int64Type, "ended_at", "response struct"); err != nil {
					return err
				} else {
					if typedEndedAt, ok := valueEndedAt.(types.Int64); ok {
						state.EndedAt = typedEndedAt
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
		assignedInvoiceSettings := false
		hadRawInvoiceSettings := false
		if rawValueInvoiceSettings, rawOk := plainValueAtPath(raw, "invoice_settings"); rawOk {
			hadRawInvoiceSettings = true
			if rawValueInvoiceSettings != nil {
				sourceInvoiceSettings := applyConfiguredKeyedListShapes(rawValueInvoiceSettings, attrValueToPlain(state.InvoiceSettings))
				if valueInvoiceSettings, err := flattenPlainValue(sourceInvoiceSettings, types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "invoice_settings", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}},
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
			if nullInvoiceSettings, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}); ok {
				if typedInvoiceSettings, ok := nullInvoiceSettings.(types.Object); ok {
					state.InvoiceSettings = typedInvoiceSettings
				}
			}
		}
	}
	{
		if rawValueItems, rawOk := plainValueAtPath(raw, "items"); rawOk {
			rawPlainItems := extractListObjectData(rawValueItems)
			if valueItems, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawPlainItems, attrValueToPlain(state.Items)), attrValueToPlain(state.Items)), attrValueToPlain(state.Items)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "billing_thresholds": types.ObjectType{AttrTypes: map[string]attr.Type{"usage_gte": types.Int64Type}}, "discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "metadata": types.MapType{ElemType: types.StringType}, "price": types.StringType, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}, "quantity": types.Int64Type, "tax_rates": types.ListType{ElemType: types.StringType}}}}, "items", "raw response"); err != nil {
				return err
			} else {
				if typedItems, ok := valueItems.(types.List); ok {
					state.Items = typedItems
				}
			}
		} else if !hasRaw {
			if responseValueItems, ok := plainFromResponseField(obj, "Items"); ok {
				fallbackPlainItems := extractListObjectData(responseValueItems)
				if valueItems, err := flattenPlainValue(
					preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(fallbackPlainItems, attrValueToPlain(state.Items)), attrValueToPlain(state.Items)), attrValueToPlain(state.Items)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "billing_thresholds": types.ObjectType{AttrTypes: map[string]attr.Type{"usage_gte": types.Int64Type}}, "discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "metadata": types.MapType{ElemType: types.StringType}, "price": types.StringType, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}, "quantity": types.Int64Type, "tax_rates": types.ListType{ElemType: types.StringType}}}},
					"items",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedItems, ok := valueItems.(types.List); ok {
						state.Items = typedItems
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
		if rawValueNextPendingInvoiceItemInvoice, rawOk := plainValueAtPath(raw, "next_pending_invoice_item_invoice"); rawOk {
			if valueNextPendingInvoiceItemInvoice, err := flattenPlainValue(rawValueNextPendingInvoiceItemInvoice, types.Int64Type, "next_pending_invoice_item_invoice", "raw response"); err != nil {
				return err
			} else {
				if typedNextPendingInvoiceItemInvoice, ok := valueNextPendingInvoiceItemInvoice.(types.Int64); ok {
					state.NextPendingInvoiceItemInvoice = typedNextPendingInvoiceItemInvoice
				}
			}
		} else if !hasRaw {
			if responseValueNextPendingInvoiceItemInvoice, ok := plainFromResponseField(obj, "NextPendingInvoiceItemInvoice"); ok {
				if valueNextPendingInvoiceItemInvoice, err := flattenPlainValue(responseValueNextPendingInvoiceItemInvoice, types.Int64Type, "next_pending_invoice_item_invoice", "response struct"); err != nil {
					return err
				} else {
					if typedNextPendingInvoiceItemInvoice, ok := valueNextPendingInvoiceItemInvoice.(types.Int64); ok {
						state.NextPendingInvoiceItemInvoice = typedNextPendingInvoiceItemInvoice
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
		assignedPauseCollection := false
		hadRawPauseCollection := false
		if rawValuePauseCollection, rawOk := plainValueAtPath(raw, "pause_collection"); rawOk {
			hadRawPauseCollection = true
			if rawValuePauseCollection != nil {
				sourcePauseCollection := applyConfiguredKeyedListShapes(rawValuePauseCollection, attrValueToPlain(state.PauseCollection))
				if valuePauseCollection, err := flattenPlainValue(sourcePauseCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"behavior": types.StringType, "resumes_at": types.Int64Type}}, "pause_collection", "raw response"); err != nil {
					return err
				} else {
					if typedPauseCollection, ok := valuePauseCollection.(types.Object); ok {
						state.PauseCollection = typedPauseCollection
						assignedPauseCollection = true
					}
				}
			}
		}
		if !assignedPauseCollection {
			if !hasRaw {
				if responseValuePauseCollection, ok := plainFromResponseField(obj, "PauseCollection"); ok {
					sourcePauseCollection := applyConfiguredKeyedListShapes(responseValuePauseCollection, attrValueToPlain(state.PauseCollection))
					if valuePauseCollection, err := flattenPlainValue(
						sourcePauseCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"behavior": types.StringType, "resumes_at": types.Int64Type}},
						"pause_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPauseCollection, ok := valuePauseCollection.(types.Object); ok {
							state.PauseCollection = typedPauseCollection
							assignedPauseCollection = true
						}
					}
				}
			}
		}
		if !assignedPauseCollection && hadRawPauseCollection {
			if nullPauseCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"behavior": types.StringType, "resumes_at": types.Int64Type}}); ok {
				if typedPauseCollection, ok := nullPauseCollection.(types.Object); ok {
					state.PauseCollection = typedPauseCollection
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
				if valuePaymentSettings, err := flattenPlainValue(sourcePaymentSettings, types.ObjectType{AttrTypes: map[string]attr.Type{"payment_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_type": types.StringType}}, "verification_method": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType}}, "network": types.StringType, "request_three_d_secure": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "type": types.StringType}}, "funding_type": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "purpose": types.StringType}}}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_seconds": types.Int64Type, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType}}}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}}}, "verification_method": types.StringType}}}}, "payment_method_types": types.ListType{ElemType: types.StringType}, "save_default_payment_method": types.StringType}}, "payment_settings", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"payment_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_type": types.StringType}}, "verification_method": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType}}, "network": types.StringType, "request_three_d_secure": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "type": types.StringType}}, "funding_type": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "purpose": types.StringType}}}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_seconds": types.Int64Type, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType}}}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}}}, "verification_method": types.StringType}}}}, "payment_method_types": types.ListType{ElemType: types.StringType}, "save_default_payment_method": types.StringType}},
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
			if nullPaymentSettings, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"payment_method_options": types.ObjectType{AttrTypes: map[string]attr.Type{"acss_debit": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"transaction_type": types.StringType}}, "verification_method": types.StringType}}, "bancontact": types.ObjectType{AttrTypes: map[string]attr.Type{"preferred_language": types.StringType}}, "card": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType}}, "network": types.StringType, "request_three_d_secure": types.StringType}}, "customer_balance": types.ObjectType{AttrTypes: map[string]attr.Type{"bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"eu_bank_transfer": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType}}, "type": types.StringType}}, "funding_type": types.StringType}}, "payto": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "purpose": types.StringType}}}}, "pix": types.ObjectType{AttrTypes: map[string]attr.Type{"expires_after_seconds": types.Int64Type, "mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_includes_iof": types.StringType, "end_date": types.StringType, "payment_schedule": types.StringType}}}}, "upi": types.ObjectType{AttrTypes: map[string]attr.Type{"mandate_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_type": types.StringType, "description": types.StringType, "end_date": types.Int64Type}}}}, "us_bank_account": types.ObjectType{AttrTypes: map[string]attr.Type{"financial_connections": types.ObjectType{AttrTypes: map[string]attr.Type{"filters": types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}}}, "permissions": types.ListType{ElemType: types.StringType}, "prefetch": types.ListType{ElemType: types.StringType}}}, "verification_method": types.StringType}}}}, "payment_method_types": types.ListType{ElemType: types.StringType}, "save_default_payment_method": types.StringType}}); ok {
				if typedPaymentSettings, ok := nullPaymentSettings.(types.Object); ok {
					state.PaymentSettings = typedPaymentSettings
				}
			}
		}
	}
	{
		assignedPendingInvoiceItemInterval := false
		hadRawPendingInvoiceItemInterval := false
		if rawValuePendingInvoiceItemInterval, rawOk := plainValueAtPath(raw, "pending_invoice_item_interval"); rawOk {
			hadRawPendingInvoiceItemInterval = true
			if rawValuePendingInvoiceItemInterval != nil {
				sourcePendingInvoiceItemInterval := applyConfiguredKeyedListShapes(rawValuePendingInvoiceItemInterval, attrValueToPlain(state.PendingInvoiceItemInterval))
				if valuePendingInvoiceItemInterval, err := flattenPlainValue(sourcePendingInvoiceItemInterval, types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "pending_invoice_item_interval", "raw response"); err != nil {
					return err
				} else {
					if typedPendingInvoiceItemInterval, ok := valuePendingInvoiceItemInterval.(types.Object); ok {
						state.PendingInvoiceItemInterval = typedPendingInvoiceItemInterval
						assignedPendingInvoiceItemInterval = true
					}
				}
			}
		}
		if !assignedPendingInvoiceItemInterval {
			if !hasRaw {
				if responseValuePendingInvoiceItemInterval, ok := plainFromResponseField(obj, "PendingInvoiceItemInterval"); ok {
					sourcePendingInvoiceItemInterval := applyConfiguredKeyedListShapes(responseValuePendingInvoiceItemInterval, attrValueToPlain(state.PendingInvoiceItemInterval))
					if valuePendingInvoiceItemInterval, err := flattenPlainValue(
						sourcePendingInvoiceItemInterval,
						types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}},
						"pending_invoice_item_interval",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPendingInvoiceItemInterval, ok := valuePendingInvoiceItemInterval.(types.Object); ok {
							state.PendingInvoiceItemInterval = typedPendingInvoiceItemInterval
							assignedPendingInvoiceItemInterval = true
						}
					}
				}
			}
		}
		if !assignedPendingInvoiceItemInterval && hadRawPendingInvoiceItemInterval {
			if nullPendingInvoiceItemInterval, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}); ok {
				if typedPendingInvoiceItemInterval, ok := nullPendingInvoiceItemInterval.(types.Object); ok {
					state.PendingInvoiceItemInterval = typedPendingInvoiceItemInterval
				}
			}
		}
	}
	{
		if true {
			if rawValuePendingSetupIntent, rawOk := plainValueAtPath(raw, "pending_setup_intent"); rawOk {
				if typedPendingSetupIntent, ok := plainToStringIDValue(rawValuePendingSetupIntent); ok {
					state.PendingSetupIntent = typedPendingSetupIntent
				}
			} else if !hasRaw {
				if responseValuePendingSetupIntent, ok := plainFromResponseField(obj, "PendingSetupIntent"); ok {
					if typedPendingSetupIntent, ok := plainToStringIDValue(responseValuePendingSetupIntent); ok {
						state.PendingSetupIntent = typedPendingSetupIntent
					}
				}
			}
		}
	}
	{
		assignedPendingUpdate := false
		hadRawPendingUpdate := false
		if rawValuePendingUpdate, rawOk := plainValueAtPath(raw, "pending_update"); rawOk {
			hadRawPendingUpdate = true
			if rawValuePendingUpdate != nil {
				sourcePendingUpdate := applyConfiguredKeyedListShapes(rawValuePendingUpdate, attrValueToPlain(state.PendingUpdate))
				if valuePendingUpdate, err := flattenPlainValue(sourcePendingUpdate, types.ObjectType{AttrTypes: map[string]attr.Type{"billing_cycle_anchor": types.Int64Type, "discount": types.StringType, "expires_at": types.Int64Type, "metadata": types.MapType{ElemType: types.StringType}, "trial_end": types.Int64Type, "trial_from_plan": types.BoolType}}, "pending_update", "raw response"); err != nil {
					return err
				} else {
					if typedPendingUpdate, ok := valuePendingUpdate.(types.Object); ok {
						state.PendingUpdate = typedPendingUpdate
						assignedPendingUpdate = true
					}
				}
			}
		}
		if !assignedPendingUpdate {
			if !hasRaw {
				if responseValuePendingUpdate, ok := plainFromResponseField(obj, "PendingUpdate"); ok {
					sourcePendingUpdate := applyConfiguredKeyedListShapes(responseValuePendingUpdate, attrValueToPlain(state.PendingUpdate))
					if valuePendingUpdate, err := flattenPlainValue(
						sourcePendingUpdate,
						types.ObjectType{AttrTypes: map[string]attr.Type{"billing_cycle_anchor": types.Int64Type, "discount": types.StringType, "expires_at": types.Int64Type, "metadata": types.MapType{ElemType: types.StringType}, "trial_end": types.Int64Type, "trial_from_plan": types.BoolType}},
						"pending_update",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPendingUpdate, ok := valuePendingUpdate.(types.Object); ok {
							state.PendingUpdate = typedPendingUpdate
							assignedPendingUpdate = true
						}
					}
				}
			}
		}
		if !assignedPendingUpdate && hadRawPendingUpdate {
			if nullPendingUpdate, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"billing_cycle_anchor": types.Int64Type, "discount": types.StringType, "expires_at": types.Int64Type, "metadata": types.MapType{ElemType: types.StringType}, "trial_end": types.Int64Type, "trial_from_plan": types.BoolType}}); ok {
				if typedPendingUpdate, ok := nullPendingUpdate.(types.Object); ok {
					state.PendingUpdate = typedPendingUpdate
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
				if valuePresentmentDetails, err := flattenPlainValue(sourcePresentmentDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"presentment_currency": types.StringType}}, "presentment_details", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"presentment_currency": types.StringType}},
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
			if nullPresentmentDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"presentment_currency": types.StringType}}); ok {
				if typedPresentmentDetails, ok := nullPresentmentDetails.(types.Object); ok {
					state.PresentmentDetails = typedPresentmentDetails
				}
			}
		}
	}
	{
		if true {
			if rawValueSchedule, rawOk := plainValueAtPath(raw, "schedule"); rawOk {
				if typedSchedule, ok := plainToStringIDValue(rawValueSchedule); ok {
					state.Schedule = typedSchedule
				}
			} else if !hasRaw {
				if responseValueSchedule, ok := plainFromResponseField(obj, "Schedule"); ok {
					if typedSchedule, ok := plainToStringIDValue(responseValueSchedule); ok {
						state.Schedule = typedSchedule
					}
				}
			}
		}
	}
	{
		if rawValueStartDate, rawOk := plainValueAtPath(raw, "start_date"); rawOk {
			if valueStartDate, err := flattenPlainValue(rawValueStartDate, types.Int64Type, "start_date", "raw response"); err != nil {
				return err
			} else {
				if typedStartDate, ok := valueStartDate.(types.Int64); ok {
					state.StartDate = typedStartDate
				}
			}
		} else if !hasRaw {
			if responseValueStartDate, ok := plainFromResponseField(obj, "StartDate"); ok {
				if valueStartDate, err := flattenPlainValue(responseValueStartDate, types.Int64Type, "start_date", "response struct"); err != nil {
					return err
				} else {
					if typedStartDate, ok := valueStartDate.(types.Int64); ok {
						state.StartDate = typedStartDate
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
		assignedTransferData := false
		hadRawTransferData := false
		if rawValueTransferData, rawOk := plainValueAtPath(raw, "transfer_data"); rawOk {
			hadRawTransferData = true
			if rawValueTransferData != nil {
				sourceTransferData := applyConfiguredKeyedListShapes(rawValueTransferData, attrValueToPlain(state.TransferData))
				if valueTransferData, err := flattenPlainValue(sourceTransferData, types.ObjectType{AttrTypes: map[string]attr.Type{"amount_percent": types.Float64Type, "destination": types.StringType}}, "transfer_data", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount_percent": types.Float64Type, "destination": types.StringType}},
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
			if nullTransferData, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount_percent": types.Float64Type, "destination": types.StringType}}); ok {
				if typedTransferData, ok := nullTransferData.(types.Object); ok {
					state.TransferData = typedTransferData
				}
			}
		}
	}
	{
		if rawValueTrialEnd, rawOk := plainValueAtPath(raw, "trial_end"); rawOk {
			if valueTrialEnd, err := flattenPlainValue(rawValueTrialEnd, types.Int64Type, "trial_end", "raw response"); err != nil {
				return err
			} else {
				if typedTrialEnd, ok := valueTrialEnd.(types.Int64); ok {
					state.TrialEnd = typedTrialEnd
				}
			}
		} else if !hasRaw {
			if responseValueTrialEnd, ok := plainFromResponseField(obj, "TrialEnd"); ok {
				if valueTrialEnd, err := flattenPlainValue(responseValueTrialEnd, types.Int64Type, "trial_end", "response struct"); err != nil {
					return err
				} else {
					if typedTrialEnd, ok := valueTrialEnd.(types.Int64); ok {
						state.TrialEnd = typedTrialEnd
					}
				}
			}
		}
	}
	{
		assignedTrialSettings := false
		hadRawTrialSettings := false
		if rawValueTrialSettings, rawOk := plainValueAtPath(raw, "trial_settings"); rawOk {
			hadRawTrialSettings = true
			if rawValueTrialSettings != nil {
				sourceTrialSettings := applyConfiguredKeyedListShapes(rawValueTrialSettings, attrValueToPlain(state.TrialSettings))
				if valueTrialSettings, err := flattenPlainValue(sourceTrialSettings, types.ObjectType{AttrTypes: map[string]attr.Type{"end_behavior": types.ObjectType{AttrTypes: map[string]attr.Type{"missing_payment_method": types.StringType}}}}, "trial_settings", "raw response"); err != nil {
					return err
				} else {
					if typedTrialSettings, ok := valueTrialSettings.(types.Object); ok {
						state.TrialSettings = typedTrialSettings
						assignedTrialSettings = true
					}
				}
			}
		}
		if !assignedTrialSettings {
			if !hasRaw {
				if responseValueTrialSettings, ok := plainFromResponseField(obj, "TrialSettings"); ok {
					sourceTrialSettings := applyConfiguredKeyedListShapes(responseValueTrialSettings, attrValueToPlain(state.TrialSettings))
					if valueTrialSettings, err := flattenPlainValue(
						sourceTrialSettings,
						types.ObjectType{AttrTypes: map[string]attr.Type{"end_behavior": types.ObjectType{AttrTypes: map[string]attr.Type{"missing_payment_method": types.StringType}}}},
						"trial_settings",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTrialSettings, ok := valueTrialSettings.(types.Object); ok {
							state.TrialSettings = typedTrialSettings
							assignedTrialSettings = true
						}
					}
				}
			}
		}
		if !assignedTrialSettings && hadRawTrialSettings {
			if nullTrialSettings, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"end_behavior": types.ObjectType{AttrTypes: map[string]attr.Type{"missing_payment_method": types.StringType}}}}); ok {
				if typedTrialSettings, ok := nullTrialSettings.(types.Object); ok {
					state.TrialSettings = typedTrialSettings
				}
			}
		}
	}
	{
		if rawValueTrialStart, rawOk := plainValueAtPath(raw, "trial_start"); rawOk {
			if valueTrialStart, err := flattenPlainValue(rawValueTrialStart, types.Int64Type, "trial_start", "raw response"); err != nil {
				return err
			} else {
				if typedTrialStart, ok := valueTrialStart.(types.Int64); ok {
					state.TrialStart = typedTrialStart
				}
			}
		} else if !hasRaw {
			if responseValueTrialStart, ok := plainFromResponseField(obj, "TrialStart"); ok {
				if valueTrialStart, err := flattenPlainValue(responseValueTrialStart, types.Int64Type, "trial_start", "response struct"); err != nil {
					return err
				} else {
					if typedTrialStart, ok := valueTrialStart.(types.Int64); ok {
						state.TrialStart = typedTrialStart
					}
				}
			}
		}
	}
	return nil
}
