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

var _ resource.Resource = &SubscriptionScheduleResource{}

var _ resource.ResourceWithConfigure = &SubscriptionScheduleResource{}

var _ resource.ResourceWithImportState = &SubscriptionScheduleResource{}

func NewSubscriptionScheduleResource() resource.Resource {
	return &SubscriptionScheduleResource{}
}

type SubscriptionScheduleResource struct {
	client *stripe.Client
}

type SubscriptionScheduleResourceModel struct {
	Object               types.String `tfsdk:"object"`
	Application          types.String `tfsdk:"application"`
	BillingMode          types.Object `tfsdk:"billing_mode"`
	CanceledAt           types.Int64  `tfsdk:"canceled_at"`
	CompletedAt          types.Int64  `tfsdk:"completed_at"`
	Created              types.Int64  `tfsdk:"created"`
	CurrentPhase         types.Object `tfsdk:"current_phase"`
	Customer             types.String `tfsdk:"customer"`
	CustomerAccount      types.String `tfsdk:"customer_account"`
	DefaultSettings      types.Object `tfsdk:"default_settings"`
	EndBehavior          types.String `tfsdk:"end_behavior"`
	ID                   types.String `tfsdk:"id"`
	Livemode             types.Bool   `tfsdk:"livemode"`
	Metadata             types.Map    `tfsdk:"metadata"`
	Phases               types.List   `tfsdk:"phases"`
	ReleasedAt           types.Int64  `tfsdk:"released_at"`
	ReleasedSubscription types.String `tfsdk:"released_subscription"`
	Status               types.String `tfsdk:"status"`
	Subscription         types.String `tfsdk:"subscription"`
	TestClock            types.String `tfsdk:"test_clock"`
	FromSubscription     types.String `tfsdk:"from_subscription"`
	StartDate            types.Int64  `tfsdk:"start_date"`
}

func (r *SubscriptionScheduleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SubscriptionScheduleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscription_schedule"
}

func (r *SubscriptionScheduleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A subscription schedule allows you to create and manage the lifecycle of a subscription by predefining expected changes.\n\nRelated guide: [Subscription schedules](https://docs.stripe.com/billing/subscriptions/subscription-schedules)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("subscription_schedule")},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the Connect Application that created the schedule.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"canceled_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the subscription schedule was canceled. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"completed_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the subscription schedule was completed. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"current_phase": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Object representing the start and end dates for the current phase of the subscription schedule, if it is `active`.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"end_date": schema.Int64Attribute{
						Computed:      true,
						Description:   "The end of this phase of the subscription schedule.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"start_date": schema.Int64Attribute{
						Computed:      true,
						Description:   "The start of this phase of the subscription schedule.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
				},
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the customer who owns the subscription schedule.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"customer_account": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the account who owns the subscription schedule.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"default_settings": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"application_fee_percent": schema.Float64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.",
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
								Description: "Whether Stripe automatically computes tax on invoices created during this phase.",
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
					"billing_cycle_anchor": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("automatic", "phase_start")},
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
					"collection_method": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("charge_automatically", "send_invoice")},
					},
					"default_payment_method": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "ID of the default payment method for the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"description": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.",
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
								Description:   "The account tax IDs associated with the subscription schedule. Will be set on invoices generated by the subscription schedule.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"days_until_due": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.",
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
					"on_behalf_of": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"transfer_data": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.",
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
				},
			},
			"end_behavior": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running. `cancel` will end the subscription schedule and cancel the underlying subscription.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("cancel", "none", "release", "renew")},
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
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"phases": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Configuration for the subscription schedule's phases.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"add_invoice_items": schema.ListNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"discountable": schema.BoolAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Controls whether discounts apply to this invoice item. Defaults to true if no value is provided.",
										PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
									},
									"discounts": schema.ListNestedAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The stackable discounts that will be applied to the item.",
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
										Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
										PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
									"period": schema.SingleNestedAttribute{
										Optional: true,
										Computed: true,

										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"end": schema.SingleNestedAttribute{
												Required: true,

												Attributes: map[string]schema.Attribute{
													"timestamp": schema.Int64Attribute{
														Optional:      true,
														Computed:      true,
														Description:   "A precise Unix timestamp for the end of the invoice item period. Must be greater than or equal to `period.start`.",
														PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
													},
													"type": schema.StringAttribute{
														Required:    true,
														Description: "Select how to calculate the end of the invoice item period.",
														Validators:  []validator.String{stringvalidator.OneOf("min_item_period_end", "phase_end", "timestamp")},
													},
												},
											},
											"start": schema.SingleNestedAttribute{
												Required: true,

												Attributes: map[string]schema.Attribute{
													"timestamp": schema.Int64Attribute{
														Optional:      true,
														Computed:      true,
														Description:   "A precise Unix timestamp for the start of the invoice item period. Must be less than or equal to `period.end`.",
														PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
													},
													"type": schema.StringAttribute{
														Required:    true,
														Description: "Select how to calculate the start of the invoice item period.",
														Validators:  []validator.String{stringvalidator.OneOf("max_item_period_start", "phase_start", "timestamp")},
													},
												},
											},
										},
									},
									"price": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "ID of the price used to generate the invoice item.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"quantity": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "The quantity of the invoice item.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"tax_rates": schema.ListAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
									"price_data": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline. One of `price` or `price_data` is required.",
										Attributes: map[string]schema.Attribute{
											"currency": schema.StringAttribute{
												Required:    true,
												Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
											},
											"product": schema.StringAttribute{
												Required:    true,
												Description: "The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to.",
											},
											"tax_behavior": schema.StringAttribute{
												Optional:    true,
												Description: "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
											},
											"unit_amount": schema.Int64Attribute{
												Optional:    true,
												Description: "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge or a negative integer representing the amount to credit to the customer.",
											},
											"unit_amount_decimal": schema.Float64Attribute{
												Optional:    true,
												Description: "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
											},
										},
									},
								},
							},
						},
						"application_fee_percent": schema.Float64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account during this phase of the schedule.",
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
									Description: "Whether Stripe automatically computes tax on invoices created during this phase.",
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
						"billing_cycle_anchor": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("automatic", "phase_start")},
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
						"collection_method": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("charge_automatically", "send_invoice")},
						},
						"currency": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"default_payment_method": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "ID of the default payment method for the subscription schedule. It must belong to the customer associated with the subscription schedule. If not set, invoices will use the default payment method in the customer's invoice settings.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"default_tax_rates": schema.ListAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The default tax rates to apply to the subscription during this phase of the subscription schedule.",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							ElementType:   types.StringType,
						},
						"description": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"discounts": schema.ListNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The stackable discounts that will be applied to the subscription on this phase. Subscription item discounts are applied before subscription discounts.",
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
						"end_date": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The end of this phase of the subscription schedule.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"invoice_settings": schema.SingleNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The invoice settings applicable during this phase.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"account_tax_ids": schema.ListAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "The account tax IDs associated with this phase of the subscription schedule. Will be set on invoices generated by this phase of the subscription schedule.",
									PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
									ElementType:   types.StringType,
								},
								"days_until_due": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"issuer": schema.SingleNestedAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "The connected account that issues the invoice. The invoice is presented with the branding and support information of the specified account.",
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
							Required:    true,
							Description: "Subscription items to configure the subscription to during this phase of the subscription schedule.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"billing_thresholds": schema.SingleNestedAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period",
										PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
										Attributes: map[string]schema.Attribute{
											"usage_gte": schema.Int64Attribute{
												Required:    true,
												Description: "Usage threshold that triggers the subscription to create an invoice",
											},
										},
									},
									"discounts": schema.ListNestedAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.",
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
										Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an item. Metadata on this item will update the underlying subscription item's `metadata` when the phase is entered.",
										PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
									"plan": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "ID of the plan to which the customer should be subscribed.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"price": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "ID of the price to which the customer should be subscribed.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"quantity": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "Quantity of the plan to which the customer should be subscribed.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"tax_rates": schema.ListAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The tax rates which apply to this `phase_item`. When set, the `default_tax_rates` on the phase do not apply to this `phase_item`.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
									"price_data": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline.",
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
														Optional:    true,
														Description: "The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).",
													},
												},
											},
											"tax_behavior": schema.StringAttribute{
												Optional:    true,
												Description: "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
											},
											"unit_amount": schema.Int64Attribute{
												Optional:    true,
												Description: "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.",
											},
											"unit_amount_decimal": schema.Float64Attribute{
												Optional:    true,
												Description: "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
											},
										},
									},
								},
							},
						},
						"metadata": schema.MapAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to a phase. Metadata on a schedule's phase will update the underlying subscription's `metadata` when the phase is entered. Updating the underlying subscription's `metadata` directly will not affect the current phase's `metadata`.",
							PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
							ElementType:   types.StringType,
						},
						"on_behalf_of": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The account (if any) the charge was made on behalf of for charges associated with the schedule's subscription. See the Connect documentation for details.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"proration_behavior": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "When transitioning phases, controls how prorations are handled (if any). Possible values are `create_prorations`, `none`, and `always_invoice`.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("always_invoice", "create_prorations", "none")},
						},
						"start_date": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "The start of this phase of the subscription schedule.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"transfer_data": schema.SingleNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The account (if any) the associated subscription's payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription's invoices.",
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
							Description:   "When the trial ends within the phase.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"duration": schema.SingleNestedAttribute{
							Optional:      true,
							Description:   "The number of intervals the phase should last. If set, `end_date` must not be set.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
							Attributes: map[string]schema.Attribute{
								"interval": schema.StringAttribute{
									Required:    true,
									Description: "Specifies phase duration. Either `day`, `week`, `month` or `year`.",
								},
								"interval_count": schema.Int64Attribute{
									Optional:    true,
									Description: "The multiplier applied to the interval.",
								},
							},
						},
						"trial": schema.BoolAttribute{
							Optional:    true,
							Description: "If set to true the entire phase is counted as a trial and the customer will not be charged for any fees.",
						},
					},
				},
			},
			"released_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the subscription schedule was released. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"released_subscription": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the subscription once managed by the subscription schedule (if it is released).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://docs.stripe.com/billing/subscriptions/subscription-schedules).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "canceled", "completed", "not_started", "released")},
			},
			"subscription": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the subscription managed by the subscription schedule.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"test_clock": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the test clock this subscription schedule belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"from_subscription": schema.StringAttribute{
				Optional:      true,
				Description:   "Migrate an existing subscription to be managed by a subscription schedule. If this parameter is set, a subscription schedule will be created using the subscription's item(s), set to auto-renew using the subscription's interval. When using this parameter, other parameters (such as phase values) cannot be set. To create a subscription schedule with other modifications, we recommend making two separate API calls.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"start_date": schema.Int64Attribute{
				Optional:      true,
				Description:   "When the subscription schedule starts. We recommend using `now` so that it starts the subscription immediately. You can also use a Unix timestamp to backdate the subscription so that it starts on a past date, or set a future date for the subscription to start on.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *SubscriptionScheduleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SubscriptionScheduleResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config SubscriptionScheduleResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"FromSubscription"}})

	params, err := expandSubscriptionScheduleCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building SubscriptionSchedule create params", err.Error())
		return
	}

	obj, err := r.client.V1SubscriptionSchedules.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating SubscriptionSchedule", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1SubscriptionSchedules.B, r.client.V1SubscriptionSchedules.Key, stripe.FormatURLPath("/v1/subscription_schedules/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating SubscriptionSchedule create raw response", err.Error())
		return
	}

	if err := flattenSubscriptionSchedule(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening SubscriptionSchedule create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Phases", "*", "add_invoice_items", "*", "price_data"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "currency"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "product"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "tax_behavior"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "unit_amount"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "unit_amount_decimal"}, []string{"Phases", "*", "items", "*", "price_data"}, []string{"Phases", "*", "items", "*", "price_data", "currency"}, []string{"Phases", "*", "items", "*", "price_data", "product"}, []string{"Phases", "*", "items", "*", "price_data", "recurring"}, []string{"Phases", "*", "items", "*", "price_data", "recurring", "interval"}, []string{"Phases", "*", "items", "*", "price_data", "recurring", "interval_count"}, []string{"Phases", "*", "items", "*", "price_data", "tax_behavior"}, []string{"Phases", "*", "items", "*", "price_data", "unit_amount"}, []string{"Phases", "*", "items", "*", "price_data", "unit_amount_decimal"}, []string{"Phases", "*", "duration"}, []string{"Phases", "*", "duration", "interval"}, []string{"Phases", "*", "duration", "interval_count"}, []string{"Phases", "*", "trial"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"FromSubscription"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SubscriptionScheduleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState SubscriptionScheduleResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state SubscriptionScheduleResourceModel
	state = priorState

	obj, err := r.client.V1SubscriptionSchedules.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading SubscriptionSchedule", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1SubscriptionSchedules.B, r.client.V1SubscriptionSchedules.Key, stripe.FormatURLPath("/v1/subscription_schedules/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating SubscriptionSchedule raw response", err.Error())
		return
	}

	if err := flattenSubscriptionSchedule(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening SubscriptionSchedule read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"Phases", "*", "add_invoice_items", "*", "price_data"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "currency"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "product"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "tax_behavior"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "unit_amount"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "unit_amount_decimal"}, []string{"Phases", "*", "items", "*", "price_data"}, []string{"Phases", "*", "items", "*", "price_data", "currency"}, []string{"Phases", "*", "items", "*", "price_data", "product"}, []string{"Phases", "*", "items", "*", "price_data", "recurring"}, []string{"Phases", "*", "items", "*", "price_data", "recurring", "interval"}, []string{"Phases", "*", "items", "*", "price_data", "recurring", "interval_count"}, []string{"Phases", "*", "items", "*", "price_data", "tax_behavior"}, []string{"Phases", "*", "items", "*", "price_data", "unit_amount"}, []string{"Phases", "*", "items", "*", "price_data", "unit_amount_decimal"}, []string{"Phases", "*", "duration"}, []string{"Phases", "*", "duration", "interval"}, []string{"Phases", "*", "duration", "interval_count"}, []string{"Phases", "*", "trial"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"FromSubscription"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *SubscriptionScheduleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SubscriptionScheduleResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config SubscriptionScheduleResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"FromSubscription"}})

	var state SubscriptionScheduleResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"FromSubscription"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"FromSubscription"}})

	params, err := expandSubscriptionScheduleUpdate(diffPlan, diffState, config)
	if err != nil {
		resp.Diagnostics.AddError("Error building SubscriptionSchedule update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building SubscriptionSchedule update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1SubscriptionSchedules.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating SubscriptionSchedule", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1SubscriptionSchedules.B, r.client.V1SubscriptionSchedules.Key, stripe.FormatURLPath("/v1/subscription_schedules/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating SubscriptionSchedule update raw response", err.Error())
		return
	}

	if err := flattenSubscriptionSchedule(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening SubscriptionSchedule update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Phases", "*", "add_invoice_items", "*", "price_data"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "currency"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "product"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "tax_behavior"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "unit_amount"}, []string{"Phases", "*", "add_invoice_items", "*", "price_data", "unit_amount_decimal"}, []string{"Phases", "*", "items", "*", "price_data"}, []string{"Phases", "*", "items", "*", "price_data", "currency"}, []string{"Phases", "*", "items", "*", "price_data", "product"}, []string{"Phases", "*", "items", "*", "price_data", "recurring"}, []string{"Phases", "*", "items", "*", "price_data", "recurring", "interval"}, []string{"Phases", "*", "items", "*", "price_data", "recurring", "interval_count"}, []string{"Phases", "*", "items", "*", "price_data", "tax_behavior"}, []string{"Phases", "*", "items", "*", "price_data", "unit_amount"}, []string{"Phases", "*", "items", "*", "price_data", "unit_amount_decimal"}, []string{"Phases", "*", "duration"}, []string{"Phases", "*", "duration", "interval"}, []string{"Phases", "*", "duration", "interval_count"}, []string{"Phases", "*", "trial"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"FromSubscription"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SubscriptionScheduleResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *SubscriptionScheduleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandSubscriptionScheduleCreate(plan SubscriptionScheduleResourceModel) (*stripe.SubscriptionScheduleCreateParams, error) {
	params := &stripe.SubscriptionScheduleCreateParams{}

	if !plan.BillingMode.IsNull() && !plan.BillingMode.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingMode", plan.BillingMode) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_mode", params)
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
	if !plan.DefaultSettings.IsNull() && !plan.DefaultSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DefaultSettings", plan.DefaultSettings) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "default_settings", params)
		}
	}
	if !plan.EndBehavior.IsNull() && !plan.EndBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "EndBehavior", "EndBehavior", plan.EndBehavior.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "end_behavior", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Phases.IsNull() && !plan.Phases.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Phases", plan.Phases) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "phases", params)
		}
	}
	if !plan.FromSubscription.IsNull() && !plan.FromSubscription.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FromSubscription", "FromSubscription", plan.FromSubscription.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "from_subscription", params)
		}
	}
	if !plan.StartDate.IsNull() && !plan.StartDate.IsUnknown() {
		params.StartDate = stripe.Int64(plan.StartDate.ValueInt64())
	}

	return params, nil
}

func expandSubscriptionScheduleUpdate(plan SubscriptionScheduleResourceModel, state SubscriptionScheduleResourceModel, config SubscriptionScheduleResourceModel) (*stripe.SubscriptionScheduleUpdateParams, error) {
	params := &stripe.SubscriptionScheduleUpdateParams{}

	if !plan.DefaultSettings.Equal(state.DefaultSettings) && !plan.DefaultSettings.IsNull() && !plan.DefaultSettings.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DefaultSettings", plan.DefaultSettings) {
			if !plan.DefaultSettings.Equal(state.DefaultSettings) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "default_settings", params)
			}
		}
	}
	if !plan.EndBehavior.Equal(state.EndBehavior) && !plan.EndBehavior.IsNull() && !plan.EndBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "EndBehavior", "EndBehavior", plan.EndBehavior.ValueString()) {
			if !plan.EndBehavior.Equal(state.EndBehavior) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "end_behavior", params)
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
	if !plan.Phases.Equal(state.Phases) && !config.Phases.IsNull() && !config.Phases.IsUnknown() {
		updatePhases := config.Phases
		if hydratedPhases, ok := hydrateAttrValueAtPaths(
			updatePhases,
			plan.Phases,
			[][]string{[]string{"*", "start_date"}},
		); ok {
			if typedPhases, ok := hydratedPhases.(types.List); ok {
				updatePhases = typedPhases
			}
		}
		if !assignAttrValueToNamedField(params, "Phases", updatePhases) {
			if !updatePhases.Equal(state.Phases) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "phases", params)
			}
		}
	}

	return params, nil
}

func flattenSubscriptionSchedule(obj *stripe.SubscriptionSchedule, state *SubscriptionScheduleResourceModel) error {
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
		if rawValueCompletedAt, rawOk := plainValueAtPath(raw, "completed_at"); rawOk {
			if valueCompletedAt, err := flattenPlainValue(rawValueCompletedAt, types.Int64Type, "completed_at", "raw response"); err != nil {
				return err
			} else {
				if typedCompletedAt, ok := valueCompletedAt.(types.Int64); ok {
					state.CompletedAt = typedCompletedAt
				}
			}
		} else if !hasRaw {
			if responseValueCompletedAt, ok := plainFromResponseField(obj, "CompletedAt"); ok {
				if valueCompletedAt, err := flattenPlainValue(responseValueCompletedAt, types.Int64Type, "completed_at", "response struct"); err != nil {
					return err
				} else {
					if typedCompletedAt, ok := valueCompletedAt.(types.Int64); ok {
						state.CompletedAt = typedCompletedAt
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
		assignedCurrentPhase := false
		hadRawCurrentPhase := false
		if rawValueCurrentPhase, rawOk := plainValueAtPath(raw, "current_phase"); rawOk {
			hadRawCurrentPhase = true
			if rawValueCurrentPhase != nil {
				sourceCurrentPhase := applyConfiguredKeyedListShapes(rawValueCurrentPhase, attrValueToPlain(state.CurrentPhase))
				if valueCurrentPhase, err := flattenPlainValue(sourceCurrentPhase, types.ObjectType{AttrTypes: map[string]attr.Type{"end_date": types.Int64Type, "start_date": types.Int64Type}}, "current_phase", "raw response"); err != nil {
					return err
				} else {
					if typedCurrentPhase, ok := valueCurrentPhase.(types.Object); ok {
						state.CurrentPhase = typedCurrentPhase
						assignedCurrentPhase = true
					}
				}
			}
		}
		if !assignedCurrentPhase {
			if !hasRaw {
				if responseValueCurrentPhase, ok := plainFromResponseField(obj, "CurrentPhase"); ok {
					sourceCurrentPhase := applyConfiguredKeyedListShapes(responseValueCurrentPhase, attrValueToPlain(state.CurrentPhase))
					if valueCurrentPhase, err := flattenPlainValue(
						sourceCurrentPhase,
						types.ObjectType{AttrTypes: map[string]attr.Type{"end_date": types.Int64Type, "start_date": types.Int64Type}},
						"current_phase",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCurrentPhase, ok := valueCurrentPhase.(types.Object); ok {
							state.CurrentPhase = typedCurrentPhase
							assignedCurrentPhase = true
						}
					}
				}
			}
		}
		if !assignedCurrentPhase && hadRawCurrentPhase {
			if nullCurrentPhase, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"end_date": types.Int64Type, "start_date": types.Int64Type}}); ok {
				if typedCurrentPhase, ok := nullCurrentPhase.(types.Object); ok {
					state.CurrentPhase = typedCurrentPhase
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
		assignedDefaultSettings := false
		hadRawDefaultSettings := false
		if rawValueDefaultSettings, rawOk := plainValueAtPath(raw, "default_settings"); rawOk {
			hadRawDefaultSettings = true
			if rawValueDefaultSettings != nil {
				sourceDefaultSettings := applyConfiguredKeyedListShapes(rawValueDefaultSettings, attrValueToPlain(state.DefaultSettings))
				if valueDefaultSettings, err := flattenPlainValue(sourceDefaultSettings, types.ObjectType{AttrTypes: map[string]attr.Type{"application_fee_percent": types.Float64Type, "automatic_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "billing_cycle_anchor": types.StringType, "billing_thresholds": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "reset_billing_cycle_anchor": types.BoolType}}, "collection_method": types.StringType, "default_payment_method": types.StringType, "description": types.StringType, "invoice_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "days_until_due": types.Int64Type, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "on_behalf_of": types.StringType, "transfer_data": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_percent": types.Float64Type, "destination": types.StringType}}}}, "default_settings", "raw response"); err != nil {
					return err
				} else {
					if typedDefaultSettings, ok := valueDefaultSettings.(types.Object); ok {
						state.DefaultSettings = typedDefaultSettings
						assignedDefaultSettings = true
					}
				}
			}
		}
		if !assignedDefaultSettings {
			if !hasRaw {
				if responseValueDefaultSettings, ok := plainFromResponseField(obj, "DefaultSettings"); ok {
					sourceDefaultSettings := applyConfiguredKeyedListShapes(responseValueDefaultSettings, attrValueToPlain(state.DefaultSettings))
					if valueDefaultSettings, err := flattenPlainValue(
						sourceDefaultSettings,
						types.ObjectType{AttrTypes: map[string]attr.Type{"application_fee_percent": types.Float64Type, "automatic_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "billing_cycle_anchor": types.StringType, "billing_thresholds": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "reset_billing_cycle_anchor": types.BoolType}}, "collection_method": types.StringType, "default_payment_method": types.StringType, "description": types.StringType, "invoice_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "days_until_due": types.Int64Type, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "on_behalf_of": types.StringType, "transfer_data": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_percent": types.Float64Type, "destination": types.StringType}}}},
						"default_settings",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedDefaultSettings, ok := valueDefaultSettings.(types.Object); ok {
							state.DefaultSettings = typedDefaultSettings
							assignedDefaultSettings = true
						}
					}
				}
			}
		}
		if !assignedDefaultSettings && hadRawDefaultSettings {
			if nullDefaultSettings, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"application_fee_percent": types.Float64Type, "automatic_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "billing_cycle_anchor": types.StringType, "billing_thresholds": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "reset_billing_cycle_anchor": types.BoolType}}, "collection_method": types.StringType, "default_payment_method": types.StringType, "description": types.StringType, "invoice_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "days_until_due": types.Int64Type, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "on_behalf_of": types.StringType, "transfer_data": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_percent": types.Float64Type, "destination": types.StringType}}}}); ok {
				if typedDefaultSettings, ok := nullDefaultSettings.(types.Object); ok {
					state.DefaultSettings = typedDefaultSettings
				}
			}
		}
	}
	{
		if rawValueEndBehavior, rawOk := plainValueAtPath(raw, "end_behavior"); rawOk {
			if valueEndBehavior, err := flattenPlainValue(rawValueEndBehavior, types.StringType, "end_behavior", "raw response"); err != nil {
				return err
			} else {
				if typedEndBehavior, ok := valueEndBehavior.(types.String); ok {
					state.EndBehavior = typedEndBehavior
				}
			}
		} else if !hasRaw {
			if responseValueEndBehavior, ok := plainFromResponseField(obj, "EndBehavior"); ok {
				if valueEndBehavior, err := flattenPlainValue(responseValueEndBehavior, types.StringType, "end_behavior", "response struct"); err != nil {
					return err
				} else {
					if typedEndBehavior, ok := valueEndBehavior.(types.String); ok {
						state.EndBehavior = typedEndBehavior
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
		if rawValuePhases, rawOk := plainValueAtPath(raw, "phases"); rawOk {
			if valuePhases, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawValuePhases, attrValueToPlain(state.Phases)), attrValueToPlain(state.Phases)), attrValueToPlain(state.Phases)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"add_invoice_items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"discountable": types.BoolType, "discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "metadata": types.MapType{ElemType: types.StringType}, "period": types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.ObjectType{AttrTypes: map[string]attr.Type{"timestamp": types.Int64Type, "type": types.StringType}}, "start": types.ObjectType{AttrTypes: map[string]attr.Type{"timestamp": types.Int64Type, "type": types.StringType}}}}, "price": types.StringType, "quantity": types.Int64Type, "tax_rates": types.ListType{ElemType: types.StringType}, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}}}}, "application_fee_percent": types.Float64Type, "automatic_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "billing_cycle_anchor": types.StringType, "billing_thresholds": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "reset_billing_cycle_anchor": types.BoolType}}, "collection_method": types.StringType, "currency": types.StringType, "default_payment_method": types.StringType, "default_tax_rates": types.ListType{ElemType: types.StringType}, "description": types.StringType, "discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "end_date": types.Int64Type, "invoice_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "days_until_due": types.Int64Type, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"billing_thresholds": types.ObjectType{AttrTypes: map[string]attr.Type{"usage_gte": types.Int64Type}}, "discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "metadata": types.MapType{ElemType: types.StringType}, "plan": types.StringType, "price": types.StringType, "quantity": types.Int64Type, "tax_rates": types.ListType{ElemType: types.StringType}, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}}}}, "metadata": types.MapType{ElemType: types.StringType}, "on_behalf_of": types.StringType, "proration_behavior": types.StringType, "start_date": types.Int64Type, "transfer_data": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_percent": types.Float64Type, "destination": types.StringType}}, "trial_end": types.Int64Type, "duration": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "trial": types.BoolType}}}, "phases", "raw response"); err != nil {
				return err
			} else {
				if typedPhases, ok := valuePhases.(types.List); ok {
					state.Phases = typedPhases
				}
			}
		} else if !hasRaw {
			if responseValuePhases, ok := plainFromResponseField(obj, "Phases"); ok {
				if valuePhases, err := flattenPlainValue(
					preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(responseValuePhases, attrValueToPlain(state.Phases)), attrValueToPlain(state.Phases)), attrValueToPlain(state.Phases)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"add_invoice_items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"discountable": types.BoolType, "discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "metadata": types.MapType{ElemType: types.StringType}, "period": types.ObjectType{AttrTypes: map[string]attr.Type{"end": types.ObjectType{AttrTypes: map[string]attr.Type{"timestamp": types.Int64Type, "type": types.StringType}}, "start": types.ObjectType{AttrTypes: map[string]attr.Type{"timestamp": types.Int64Type, "type": types.StringType}}}}, "price": types.StringType, "quantity": types.Int64Type, "tax_rates": types.ListType{ElemType: types.StringType}, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}}}}, "application_fee_percent": types.Float64Type, "automatic_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"disabled_reason": types.StringType, "enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "billing_cycle_anchor": types.StringType, "billing_thresholds": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_gte": types.Int64Type, "reset_billing_cycle_anchor": types.BoolType}}, "collection_method": types.StringType, "currency": types.StringType, "default_payment_method": types.StringType, "default_tax_rates": types.ListType{ElemType: types.StringType}, "description": types.StringType, "discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "end_date": types.Int64Type, "invoice_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "days_until_due": types.Int64Type, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"billing_thresholds": types.ObjectType{AttrTypes: map[string]attr.Type{"usage_gte": types.Int64Type}}, "discounts": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "metadata": types.MapType{ElemType: types.StringType}, "plan": types.StringType, "price": types.StringType, "quantity": types.Int64Type, "tax_rates": types.ListType{ElemType: types.StringType}, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}}}}, "metadata": types.MapType{ElemType: types.StringType}, "on_behalf_of": types.StringType, "proration_behavior": types.StringType, "start_date": types.Int64Type, "transfer_data": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_percent": types.Float64Type, "destination": types.StringType}}, "trial_end": types.Int64Type, "duration": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "trial": types.BoolType}}},
					"phases",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedPhases, ok := valuePhases.(types.List); ok {
						state.Phases = typedPhases
					}
				}
			}
		}
	}
	{
		if rawValueReleasedAt, rawOk := plainValueAtPath(raw, "released_at"); rawOk {
			if valueReleasedAt, err := flattenPlainValue(rawValueReleasedAt, types.Int64Type, "released_at", "raw response"); err != nil {
				return err
			} else {
				if typedReleasedAt, ok := valueReleasedAt.(types.Int64); ok {
					state.ReleasedAt = typedReleasedAt
				}
			}
		} else if !hasRaw {
			if responseValueReleasedAt, ok := plainFromResponseField(obj, "ReleasedAt"); ok {
				if valueReleasedAt, err := flattenPlainValue(responseValueReleasedAt, types.Int64Type, "released_at", "response struct"); err != nil {
					return err
				} else {
					if typedReleasedAt, ok := valueReleasedAt.(types.Int64); ok {
						state.ReleasedAt = typedReleasedAt
					}
				}
			}
		}
	}
	{
		if rawValueReleasedSubscription, rawOk := plainValueAtPath(raw, "released_subscription"); rawOk {
			if valueReleasedSubscription, err := flattenPlainValue(rawValueReleasedSubscription, types.StringType, "released_subscription", "raw response"); err != nil {
				return err
			} else {
				if typedReleasedSubscription, ok := valueReleasedSubscription.(types.String); ok {
					state.ReleasedSubscription = typedReleasedSubscription
				}
			}
		} else if !hasRaw {
			if responseValueReleasedSubscription, ok := plainFromResponseField(obj, "ReleasedSubscription"); ok {
				if valueReleasedSubscription, err := flattenPlainValue(responseValueReleasedSubscription, types.StringType, "released_subscription", "response struct"); err != nil {
					return err
				} else {
					if typedReleasedSubscription, ok := valueReleasedSubscription.(types.String); ok {
						state.ReleasedSubscription = typedReleasedSubscription
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
		if rawValueStartDate, rawOk := plainValueAtAnyPath(raw, [][]string{[]string{"start_date"}, []string{"phases", "0", "start_date"}}); rawOk {
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
	return nil
}
