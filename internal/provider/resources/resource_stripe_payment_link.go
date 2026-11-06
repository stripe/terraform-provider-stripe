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
	"reflect"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &PaymentLinkResource{}

var _ resource.ResourceWithConfigure = &PaymentLinkResource{}

var _ resource.ResourceWithImportState = &PaymentLinkResource{}

func NewPaymentLinkResource() resource.Resource {
	return &PaymentLinkResource{}
}

type PaymentLinkResource struct {
	client *stripe.Client
}

type PaymentLinkResourceModel struct {
	Object                    types.String  `tfsdk:"object"`
	Active                    types.Bool    `tfsdk:"active"`
	AfterCompletion           types.Object  `tfsdk:"after_completion"`
	AllowPromotionCodes       types.Bool    `tfsdk:"allow_promotion_codes"`
	Application               types.String  `tfsdk:"application"`
	ApplicationFeeAmount      types.Int64   `tfsdk:"application_fee_amount"`
	ApplicationFeePercent     types.Float64 `tfsdk:"application_fee_percent"`
	AutomaticTax              types.Object  `tfsdk:"automatic_tax"`
	BillingAddressCollection  types.String  `tfsdk:"billing_address_collection"`
	ConsentCollection         types.Object  `tfsdk:"consent_collection"`
	Currency                  types.String  `tfsdk:"currency"`
	CustomFields              types.List    `tfsdk:"custom_fields"`
	CustomText                types.Object  `tfsdk:"custom_text"`
	CustomerCreation          types.String  `tfsdk:"customer_creation"`
	ID                        types.String  `tfsdk:"id"`
	InactiveMessage           types.String  `tfsdk:"inactive_message"`
	InvoiceCreation           types.Object  `tfsdk:"invoice_creation"`
	LineItems                 types.List    `tfsdk:"line_items"`
	Livemode                  types.Bool    `tfsdk:"livemode"`
	ManagedPayments           types.Object  `tfsdk:"managed_payments"`
	Metadata                  types.Map     `tfsdk:"metadata"`
	NameCollection            types.Object  `tfsdk:"name_collection"`
	OnBehalfOf                types.String  `tfsdk:"on_behalf_of"`
	OptionalItems             types.List    `tfsdk:"optional_items"`
	PaymentIntentData         types.Object  `tfsdk:"payment_intent_data"`
	PaymentMethodCollection   types.String  `tfsdk:"payment_method_collection"`
	PaymentMethodOptions      types.Object  `tfsdk:"payment_method_options"`
	PaymentMethodTypes        types.List    `tfsdk:"payment_method_types"`
	PhoneNumberCollection     types.Object  `tfsdk:"phone_number_collection"`
	Restrictions              types.Object  `tfsdk:"restrictions"`
	ShippingAddressCollection types.Object  `tfsdk:"shipping_address_collection"`
	ShippingOptions           types.List    `tfsdk:"shipping_options"`
	SubmitType                types.String  `tfsdk:"submit_type"`
	SubscriptionData          types.Object  `tfsdk:"subscription_data"`
	TaxIDCollection           types.Object  `tfsdk:"tax_id_collection"`
	TransferData              types.Object  `tfsdk:"transfer_data"`
	URL                       types.String  `tfsdk:"url"`
}

func (r *PaymentLinkResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PaymentLinkResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_payment_link"
}

func (r *PaymentLinkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A payment link is a shareable URL that will take your customers to a hosted payment page. A payment link can be shared and used multiple times.\n\nWhen a customer opens a payment link it will open a new [checkout session](https://docs.stripe.com/api/checkout/sessions) to render the payment page. You can use [checkout session events](https://docs.stripe.com/api/events/types#event_types-checkout.session.completed) to track payments through payment links.\n\nRelated guide: [Payment Links API](https://docs.stripe.com/payment-links)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("payment_link")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the payment link's `url` is active. If `false`, customers visiting the URL will be shown a page saying that the link has been deactivated.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"after_completion": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"hosted_confirmation": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"custom_message": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The custom message that is displayed to the customer after the purchase is complete.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"redirect": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"url": schema.StringAttribute{
								Required:    true,
								Description: "The URL the customer will be redirected to after the purchase is complete.",
							},
						},
					},
					"type": schema.StringAttribute{
						Required:    true,
						Description: "The specified behavior after the purchase is complete.",
						Validators:  []validator.String{stringvalidator.OneOf("hosted_confirmation", "redirect")},
					},
				},
			},
			"allow_promotion_codes": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether user redeemable promotion codes are enabled.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the Connect application that created the Payment Link.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"application_fee_amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The amount of the application fee (if any) that will be requested to be applied to the payment and transferred to the application owner's Stripe account.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"application_fee_percent": schema.Float64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "This represents the percentage of the subscription invoice total that will be transferred to the application owner's Stripe account.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown(), float64planmodifier.RequiresReplace()},
			},
			"automatic_tax": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Required:    true,
						Description: "If `true`, tax will be calculated automatically using the customer's location.",
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
			"billing_address_collection": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Configuration for collecting the customer's billing address. Defaults to `auto`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("auto", "required")},
			},
			"consent_collection": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "When set, provides configuration to gather active consent from customers.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"payment_method_reuse_agreement": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Settings related to the payment method reuse text shown in the Checkout UI.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"position": schema.StringAttribute{
								Required:      true,
								Description:   "Determines the position and visibility of the payment method reuse agreement in the UI. When set to `auto`, Stripe's defaults will be used.\n\nWhen set to `hidden`, the payment method reuse agreement text will always be hidden in the UI.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("auto", "hidden")},
							},
						},
					},
					"promotions": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "If set to `auto`, enables the collection of customer consent for promotional communications.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("auto", "none")},
					},
					"terms_of_service": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "If set to `required`, it requires cutomers to accept the terms of service before being able to pay. If set to `none`, customers won't be shown a checkbox to accept the terms of service.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("none", "required")},
					},
				},
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
				Description:   "Collect additional information from your customer using custom fields. Up to 3 fields are supported. You can't set this parameter if `ui_mode` is `custom`.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"dropdown": schema.SingleNestedAttribute{
							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"default_value": schema.StringAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "The value that pre-fills on the payment page.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"options": schema.ListNestedAttribute{
									Required:    true,
									Description: "The options available for the customer to select. Up to 200 options allowed.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"label": schema.StringAttribute{
												Required:    true,
												Description: "The label for the option, displayed to the customer. Up to 100 characters.",
											},
											"value": schema.StringAttribute{
												Required:    true,
												Description: "The value for this option, not displayed to the customer, used by your integration to reconcile the option selected by the customer. Must be unique to this option, alphanumeric, and up to 100 characters.",
											},
										},
									},
								},
							},
						},
						"key": schema.StringAttribute{
							Required:    true,
							Description: "String of your choice that your integration can use to reconcile this field. Must be unique to this field, alphanumeric, and up to 200 characters.",
						},
						"label": schema.SingleNestedAttribute{
							Required: true,

							Attributes: map[string]schema.Attribute{
								"custom": schema.StringAttribute{
									Required:    true,
									Description: "Custom text for the label, displayed to the customer. Up to 50 characters.",
								},
								"type": schema.StringAttribute{
									Required:    true,
									Description: "The type of the label.",
									Validators:  []validator.String{stringvalidator.OneOf("custom")},
								},
							},
						},
						"numeric": schema.SingleNestedAttribute{
							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"default_value": schema.StringAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "The value that pre-fills the field on the payment page.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"maximum_length": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "The maximum character length constraint for the customer's input.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"minimum_length": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "The minimum character length requirement for the customer's input.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
							},
						},
						"optional": schema.BoolAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "Whether the customer is required to complete the field before completing the Checkout Session. Defaults to `false`.",
							PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
						},
						"text": schema.SingleNestedAttribute{
							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"default_value": schema.StringAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "The value that pre-fills the field on the payment page.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"maximum_length": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "The maximum character length constraint for the customer's input.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"minimum_length": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "The minimum character length requirement for the customer's input.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
							},
						},
						"type": schema.StringAttribute{
							Required:    true,
							Description: "The type of the field.",
							Validators:  []validator.String{stringvalidator.OneOf("dropdown", "numeric", "text")},
						},
					},
				},
			},
			"custom_text": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"after_submit": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Custom text that should be displayed after the payment confirmation button.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"message": schema.StringAttribute{
								Required:    true,
								Description: "Text can be up to 1200 characters in length.",
							},
						},
					},
					"shipping_address": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Custom text that should be displayed alongside shipping address collection.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"message": schema.StringAttribute{
								Required:    true,
								Description: "Text can be up to 1200 characters in length.",
							},
						},
					},
					"submit": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Custom text that should be displayed alongside the payment confirmation button.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"message": schema.StringAttribute{
								Required:    true,
								Description: "Text can be up to 1200 characters in length.",
							},
						},
					},
					"terms_of_service_acceptance": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Custom text that should be displayed in place of the default terms of service agreement text.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"message": schema.StringAttribute{
								Required:    true,
								Description: "Text can be up to 1200 characters in length.",
							},
						},
					},
				},
			},
			"customer_creation": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Configuration for Customer creation during checkout.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("always", "if_required")},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"inactive_message": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The custom message to be displayed to a customer when a payment link is no longer active.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"invoice_creation": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Configuration for creating invoice for payment mode payment links.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Required:    true,
						Description: "Enable creating an invoice on successful payment.",
					},
					"invoice_data": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Configuration for the invoice. Default invoice values will be used if unspecified.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account_tax_ids": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account tax IDs associated with the invoice.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"custom_fields": schema.ListNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "A list of up to 4 custom fields to be displayed on the invoice.",
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
							"description": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "An arbitrary string attached to the object. Often useful for displaying to users.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"footer": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Footer to be displayed on the invoice.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
							"metadata": schema.MapAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
								PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"rendering_options": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Options for invoice PDF rendering.",
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
										Description:   "ID of the invoice rendering template to be used for the generated invoice.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
				},
			},
			"line_items": schema.ListNestedAttribute{
				Required:    true,
				Description: "The line items representing what is being sold.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:      true,
							Description:   "Unique identifier for the object.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"adjustable_quantity": schema.SingleNestedAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "When set, provides configuration for this item’s quantity to be adjusted by the customer during checkout.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Required:    true,
									Description: "Set to true if the quantity can be adjusted to any non-negative Integer.",
								},
								"maximum": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "The maximum quantity the customer can purchase. By default this value is 99. You can specify a value up to 999999.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"minimum": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "The minimum quantity the customer can purchase. By default this value is 0. If there is only one item in the cart then that item's quantity cannot go down to 0.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
							},
						},
						"price": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The ID of the [Price](https://docs.stripe.com/api/prices) or [Plan](https://docs.stripe.com/api/plans) object. One of `price` or `price_data` is required.",
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
									Optional:      true,
									Computed:      true,
									Description:   "The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to. One of `product` or `product_data` is required.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"product_data": schema.SingleNestedAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "Data used to generate a new [Product](https://docs.stripe.com/api/products) object inline. One of `product` or `product_data` is required.",
									PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
									Attributes: map[string]schema.Attribute{
										"description": schema.StringAttribute{
											Optional:      true,
											Computed:      true,
											Description:   "The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.",
											PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										},
										"images": schema.ListAttribute{
											Optional:      true,
											Computed:      true,
											Description:   "A list of up to 8 URLs of images for this product, meant to be displayable to the customer.",
											PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
											ElementType:   types.StringType,
										},
										"metadata": schema.MapAttribute{
											Optional:      true,
											Computed:      true,
											Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
											PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
											ElementType:   types.StringType,
										},
										"name": schema.StringAttribute{
											Required:    true,
											Description: "The product's name, meant to be displayable to the customer.",
										},
										"tax_code": schema.StringAttribute{
											Optional:      true,
											Computed:      true,
											Description:   "A [tax code](https://docs.stripe.com/tax/tax-categories) ID.",
											PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										},
										"unit_label": schema.StringAttribute{
											Optional:      true,
											Computed:      true,
											Description:   "A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.",
											PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										},
									},
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
									Description:   "A non-negative integer in cents (or local equivalent) representing how much to charge. One of `unit_amount` or `unit_amount_decimal` is required.",
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
							Required:    true,
							Description: "The quantity of the line item being purchased.",
							WriteOnly:   true,
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
				Optional:      true,
				Computed:      true,
				Description:   "Settings for Managed Payments for this Payment Link and resulting [CheckoutSessions](/api/checkout/sessions/object), [PaymentIntents](/api/payment_intents/object), [Invoices](/api/invoices/object), and [Subscriptions](/api/subscriptions/object).",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Set to `true` to enable [Managed Payments](https://docs.stripe.com/payments/managed-payments), Stripe's merchant of record solution, for this session.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown(), boolplanmodifier.RequiresReplace()},
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
			"name_collection": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"business": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Required:    true,
								Description: "Indicates whether business name collection is enabled for the payment link.",
							},
							"optional": schema.BoolAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Whether the customer is required to complete the field before checking out. Defaults to `false`.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"individual": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Required:    true,
								Description: "Indicates whether individual name collection is enabled for the payment link.",
							},
							"optional": schema.BoolAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Whether the customer is required to complete the field before checking out. Defaults to `false`.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"on_behalf_of": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account on behalf of which to charge. See the [Connect documentation](https://support.stripe.com/questions/sending-invoices-on-behalf-of-connected-accounts) for details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"optional_items": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The optional items presented to the customer at checkout.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"adjustable_quantity": schema.SingleNestedAttribute{
							Optional: true,
							Computed: true,

							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Required:    true,
									Description: "Set to true if the quantity can be adjusted to any non-negative integer.",
								},
								"maximum": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "The maximum quantity of this item the customer can purchase. By default this value is 99.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"minimum": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "The minimum quantity of this item the customer must purchase, if they choose to purchase it. Because this item is optional, the customer will always be able to remove it from their order, even if the `minimum` configured here is greater than 0. By default this value is 0.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
							},
						},
						"price": schema.StringAttribute{
							Required: true,
						},
						"quantity": schema.Int64Attribute{
							Required: true,
						},
					},
				},
			},
			"payment_intent_data": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Indicates the parameters to be passed to PaymentIntent creation during checkout.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"capture_method": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Indicates when the funds will be captured from the customer's account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("automatic", "automatic_async", "manual")},
					},
					"description": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "An arbitrary string attached to the object. Often useful for displaying to users.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"metadata": schema.MapAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will set metadata on [Payment Intents](https://docs.stripe.com/api/payment_intents) generated from this payment link.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"setup_future_usage": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Indicates that you intend to make future payments with the payment method collected during checkout.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("off_session", "on_session")},
					},
					"statement_descriptor": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "For a non-card payment, information about the charge that appears on the customer's statement when this payment succeeds in creating a charge.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"statement_descriptor_suffix": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "For a card payment, information about the charge that appears on the customer's statement when this payment succeeds in creating a charge. Concatenated with the account's statement descriptor prefix to form the complete statement descriptor.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"transfer_group": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A string that identifies the resulting payment as part of a group. See the PaymentIntents [use case for connected accounts](https://docs.stripe.com/connect/separate-charges-and-transfers) for details.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"payment_method_collection": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Configuration for collecting a payment method during checkout. Defaults to `always`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("always", "if_required")},
			},
			"payment_method_options": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Payment-method-specific configuration.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"card": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Configuration for `card` payment methods.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"restrictions": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Restrictions to apply to the card payment method. For example, you can block specific card brands.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"brands_blocked": schema.ListAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The card brands to block. If a customer enters or selects a card belonging to a blocked brand, they can't complete the payment.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										ElementType:   types.StringType,
									},
								},
							},
						},
					},
				},
			},
			"payment_method_types": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The list of payment method types that customers can use. When `null`, Stripe will dynamically show relevant payment methods you've enabled in your [payment method settings](https://dashboard.stripe.com/settings/payment_methods).",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"phone_number_collection": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Required:    true,
						Description: "If `true`, a phone number will be collected during checkout.",
					},
				},
			},
			"restrictions": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Settings that restrict the usage of a payment link.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"completed_sessions": schema.SingleNestedAttribute{
						Required: true,

						Attributes: map[string]schema.Attribute{
							"count": schema.Int64Attribute{
								Computed:      true,
								Description:   "The current number of checkout sessions that have been completed on the payment link which count towards the `completed_sessions` restriction to be met.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"limit": schema.Int64Attribute{
								Required:    true,
								Description: "The maximum number of checkout sessions that can be completed for the `completed_sessions` restriction to be met.",
							},
						},
					},
				},
			},
			"shipping_address_collection": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Configuration for collecting the customer's shipping address.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"allowed_countries": schema.ListAttribute{
						Required:    true,
						Description: "An array of two-letter ISO country codes representing which countries Checkout should provide as options for shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SD, SY, UM, VI`.",
						ElementType: types.StringType,
					},
				},
			},
			"shipping_options": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The shipping rate options applied to the session.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"shipping_amount": schema.Int64Attribute{
							Computed:      true,
							Description:   "A non-negative integer in cents representing how much to charge.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"shipping_rate": schema.StringAttribute{
							Optional:      true,
							Computed:      true,
							Description:   "The ID of the Shipping Rate to use for this shipping option.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
						},
					},
				},
			},
			"submit_type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Indicates the type of transaction being performed which customizes relevant text on the page, such as the submit button.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("auto", "book", "donate", "pay", "subscribe")},
			},
			"subscription_data": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "When creating a subscription, the specified configuration data will be used. There must be at least one line item with a recurring price to use `subscription_data`.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"invoice_settings": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
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
					"metadata": schema.MapAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that will set metadata on [Subscriptions](https://docs.stripe.com/api/subscriptions) generated from this payment link.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"trial_period_days": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "Integer representing the number of trial period days before the customer is charged for the first time.",
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
								Description: "Defines how a subscription behaves when a free trial ends.",
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
				},
			},
			"tax_id_collection": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Required:    true,
						Description: "Indicates whether tax ID collection is enabled for the session.",
					},
					"required": schema.StringAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("if_supported", "never")},
					},
				},
			},
			"transfer_data": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"amount": schema.Int64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The amount in cents (or local equivalent) that will be transferred to the destination account. By default, the entire amount is transferred to the destination.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
					},
					"destination": schema.StringAttribute{
						Required:      true,
						Description:   "The connected account receiving the transfer.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"url": schema.StringAttribute{
				Computed:      true,
				Description:   "The public URL that can be shared with customers.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *PaymentLinkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PaymentLinkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PaymentLinkResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"LineItems", "*", "quantity"}})

	params, err := expandPaymentLinkCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentLink create params", err.Error())
		return
	}

	obj, err := r.client.V1PaymentLinks.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating PaymentLink", err.Error())
		return
	}

	rawReadParams := &stripe.PaymentLinkRetrieveParams{}
	rawReadParams.AddExpand("line_items")

	if err := ensureRawResponse(obj, r.client.V1PaymentLinks.B, r.client.V1PaymentLinks.Key, stripe.FormatURLPath("/v1/payment_links/%s", obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentLink create raw response", err.Error())
		return
	}

	var createdState PaymentLinkResourceModel
	if err := flattenPaymentLink(obj, &createdState); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentLink create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&createdState, &config, [][]string{[]string{"LineItems", "*", "adjustable_quantity"}, []string{"LineItems", "*", "adjustable_quantity", "enabled"}, []string{"LineItems", "*", "adjustable_quantity", "maximum"}, []string{"LineItems", "*", "adjustable_quantity", "minimum"}, []string{"LineItems", "*", "price"}, []string{"LineItems", "*", "price_data"}, []string{"LineItems", "*", "price_data", "currency"}, []string{"LineItems", "*", "price_data", "product"}, []string{"LineItems", "*", "price_data", "product_data"}, []string{"LineItems", "*", "price_data", "product_data", "description"}, []string{"LineItems", "*", "price_data", "product_data", "images"}, []string{"LineItems", "*", "price_data", "product_data", "metadata"}, []string{"LineItems", "*", "price_data", "product_data", "name"}, []string{"LineItems", "*", "price_data", "product_data", "tax_code"}, []string{"LineItems", "*", "price_data", "product_data", "unit_label"}, []string{"LineItems", "*", "price_data", "recurring"}, []string{"LineItems", "*", "price_data", "recurring", "interval"}, []string{"LineItems", "*", "price_data", "recurring", "interval_count"}, []string{"LineItems", "*", "price_data", "tax_behavior"}, []string{"LineItems", "*", "price_data", "unit_amount"}, []string{"LineItems", "*", "price_data", "unit_amount_decimal"}})
	normalizeUnknownValues(&createdState)

	diffPlan := plan
	diffCreatedState := createdState

	postCreateParams, err := expandPaymentLinkPostCreateUpdate(diffPlan, diffCreatedState)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentLink post-create update params", err.Error())
		return
	}

	if paramsHaveValues(postCreateParams) {
		if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
			!createdState.Metadata.IsNull() && !createdState.Metadata.IsUnknown() {
			if !assignMetadataDiffToNamedField(postCreateParams, "Metadata", plan.Metadata, createdState.Metadata) {
				resp.Diagnostics.AddError("Error building PaymentLink update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", postCreateParams))
				return
			}
		}
		obj, err = r.client.V1PaymentLinks.Update(ctx, createdState.ID.ValueString(), postCreateParams)
		if err != nil {
			resp.Diagnostics.AddError("Error finalizing PaymentLink after create", err.Error())
			return
		}
		rawReadParams := &stripe.PaymentLinkRetrieveParams{}
		rawReadParams.AddExpand("line_items")

		if err := ensureRawResponse(obj, r.client.V1PaymentLinks.B, r.client.V1PaymentLinks.Key, stripe.FormatURLPath("/v1/payment_links/%s", obj.ID), rawReadParams); err != nil {
			resp.Diagnostics.AddError("Error hydrating PaymentLink post-create update raw response", err.Error())
			return
		}
	}

	if err := flattenPaymentLink(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentLink create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"LineItems", "*", "adjustable_quantity"}, []string{"LineItems", "*", "adjustable_quantity", "enabled"}, []string{"LineItems", "*", "adjustable_quantity", "maximum"}, []string{"LineItems", "*", "adjustable_quantity", "minimum"}, []string{"LineItems", "*", "price"}, []string{"LineItems", "*", "price_data"}, []string{"LineItems", "*", "price_data", "currency"}, []string{"LineItems", "*", "price_data", "product"}, []string{"LineItems", "*", "price_data", "product_data"}, []string{"LineItems", "*", "price_data", "product_data", "description"}, []string{"LineItems", "*", "price_data", "product_data", "images"}, []string{"LineItems", "*", "price_data", "product_data", "metadata"}, []string{"LineItems", "*", "price_data", "product_data", "name"}, []string{"LineItems", "*", "price_data", "product_data", "tax_code"}, []string{"LineItems", "*", "price_data", "product_data", "unit_label"}, []string{"LineItems", "*", "price_data", "recurring"}, []string{"LineItems", "*", "price_data", "recurring", "interval"}, []string{"LineItems", "*", "price_data", "recurring", "interval_count"}, []string{"LineItems", "*", "price_data", "tax_behavior"}, []string{"LineItems", "*", "price_data", "unit_amount"}, []string{"LineItems", "*", "price_data", "unit_amount_decimal"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"LineItems", "*", "quantity"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentLinkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState PaymentLinkResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PaymentLinkResourceModel
	state = priorState

	params := &stripe.PaymentLinkRetrieveParams{}
	params.AddExpand("line_items")

	obj, err := r.client.V1PaymentLinks.Retrieve(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error reading PaymentLink", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentLinks.B, r.client.V1PaymentLinks.Key, stripe.FormatURLPath("/v1/payment_links/%s", state.ID.ValueString()), params); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentLink raw response", err.Error())
		return
	}

	if err := flattenPaymentLink(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentLink read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"LineItems", "*", "adjustable_quantity"}, []string{"LineItems", "*", "adjustable_quantity", "enabled"}, []string{"LineItems", "*", "adjustable_quantity", "maximum"}, []string{"LineItems", "*", "adjustable_quantity", "minimum"}, []string{"LineItems", "*", "price"}, []string{"LineItems", "*", "price_data"}, []string{"LineItems", "*", "price_data", "currency"}, []string{"LineItems", "*", "price_data", "product"}, []string{"LineItems", "*", "price_data", "product_data"}, []string{"LineItems", "*", "price_data", "product_data", "description"}, []string{"LineItems", "*", "price_data", "product_data", "images"}, []string{"LineItems", "*", "price_data", "product_data", "metadata"}, []string{"LineItems", "*", "price_data", "product_data", "name"}, []string{"LineItems", "*", "price_data", "product_data", "tax_code"}, []string{"LineItems", "*", "price_data", "product_data", "unit_label"}, []string{"LineItems", "*", "price_data", "recurring"}, []string{"LineItems", "*", "price_data", "recurring", "interval"}, []string{"LineItems", "*", "price_data", "recurring", "interval_count"}, []string{"LineItems", "*", "price_data", "tax_behavior"}, []string{"LineItems", "*", "price_data", "unit_amount"}, []string{"LineItems", "*", "price_data", "unit_amount_decimal"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"LineItems", "*", "quantity"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentLinkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PaymentLinkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PaymentLinkResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"LineItems", "*", "quantity"}})

	var state PaymentLinkResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandPaymentLinkUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentLink update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building PaymentLink update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1PaymentLinks.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating PaymentLink", err.Error())
		return
	}

	rawReadParams := &stripe.PaymentLinkRetrieveParams{}
	rawReadParams.AddExpand("line_items")

	if err := ensureRawResponse(obj, r.client.V1PaymentLinks.B, r.client.V1PaymentLinks.Key, stripe.FormatURLPath("/v1/payment_links/%s", obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentLink update raw response", err.Error())
		return
	}

	if err := flattenPaymentLink(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentLink update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"LineItems", "*", "adjustable_quantity"}, []string{"LineItems", "*", "adjustable_quantity", "enabled"}, []string{"LineItems", "*", "adjustable_quantity", "maximum"}, []string{"LineItems", "*", "adjustable_quantity", "minimum"}, []string{"LineItems", "*", "price"}, []string{"LineItems", "*", "price_data"}, []string{"LineItems", "*", "price_data", "currency"}, []string{"LineItems", "*", "price_data", "product"}, []string{"LineItems", "*", "price_data", "product_data"}, []string{"LineItems", "*", "price_data", "product_data", "description"}, []string{"LineItems", "*", "price_data", "product_data", "images"}, []string{"LineItems", "*", "price_data", "product_data", "metadata"}, []string{"LineItems", "*", "price_data", "product_data", "name"}, []string{"LineItems", "*", "price_data", "product_data", "tax_code"}, []string{"LineItems", "*", "price_data", "product_data", "unit_label"}, []string{"LineItems", "*", "price_data", "recurring"}, []string{"LineItems", "*", "price_data", "recurring", "interval"}, []string{"LineItems", "*", "price_data", "recurring", "interval_count"}, []string{"LineItems", "*", "price_data", "tax_behavior"}, []string{"LineItems", "*", "price_data", "unit_amount"}, []string{"LineItems", "*", "price_data", "unit_amount_decimal"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"LineItems", "*", "quantity"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentLinkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PaymentLinkResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.Active.IsNull() && !state.Active.IsUnknown() && !state.Active.ValueBool() {
		return
	}

	params := &stripe.PaymentLinkUpdateParams{}
	activeField := reflect.ValueOf(params).Elem().FieldByName("Active")
	if activeField.IsValid() && activeField.CanSet() {
		if activeField.Kind() == reflect.Pointer && activeField.Type().Elem().Kind() == reflect.Bool {
			activeField.Set(reflect.ValueOf(stripe.Bool(false)))
		} else if activeField.Kind() == reflect.Bool {
			activeField.SetBool(false)
		}
	}

	_, err := r.client.V1PaymentLinks.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating PaymentLink", err.Error())
		return
	}
}

func (r *PaymentLinkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandPaymentLinkCreate(plan PaymentLinkResourceModel) (*stripe.PaymentLinkCreateParams, error) {
	params := &stripe.PaymentLinkCreateParams{}

	if !plan.AfterCompletion.IsNull() && !plan.AfterCompletion.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AfterCompletion", plan.AfterCompletion) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "after_completion", params)
		}
	}
	if !plan.AllowPromotionCodes.IsNull() && !plan.AllowPromotionCodes.IsUnknown() {
		params.AllowPromotionCodes = stripe.Bool(plan.AllowPromotionCodes.ValueBool())
	}
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
	if !plan.BillingAddressCollection.IsNull() && !plan.BillingAddressCollection.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "BillingAddressCollection", "BillingAddressCollection", plan.BillingAddressCollection.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_address_collection", params)
		}
	}
	if !plan.ConsentCollection.IsNull() && !plan.ConsentCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ConsentCollection", plan.ConsentCollection) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "consent_collection", params)
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
	if !plan.CustomText.IsNull() && !plan.CustomText.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomText", plan.CustomText) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "custom_text", params)
		}
	}
	if !plan.CustomerCreation.IsNull() && !plan.CustomerCreation.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerCreation", "CustomerCreation", plan.CustomerCreation.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer_creation", params)
		}
	}
	if !plan.InactiveMessage.IsNull() && !plan.InactiveMessage.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "InactiveMessage", "InactiveMessage", plan.InactiveMessage.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "inactive_message", params)
		}
	}
	if !plan.InvoiceCreation.IsNull() && !plan.InvoiceCreation.IsUnknown() {
		if !assignAttrValueToNamedField(params, "InvoiceCreation", plan.InvoiceCreation) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "invoice_creation", params)
		}
	}
	if !plan.LineItems.IsNull() && !plan.LineItems.IsUnknown() {
		if !assignAttrValueToNamedField(params, "LineItems", plan.LineItems) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "line_items", params)
		}
	}
	if !plan.ManagedPayments.IsNull() && !plan.ManagedPayments.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ManagedPayments", plan.ManagedPayments) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "managed_payments", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.NameCollection.IsNull() && !plan.NameCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "NameCollection", plan.NameCollection) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "name_collection", params)
		}
	}
	if !plan.OnBehalfOf.IsNull() && !plan.OnBehalfOf.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "OnBehalfOfID", "OnBehalfOf", plan.OnBehalfOf.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "on_behalf_of", params)
		}
	}
	if !plan.OptionalItems.IsNull() && !plan.OptionalItems.IsUnknown() {
		if !assignAttrValueToNamedField(params, "OptionalItems", plan.OptionalItems) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "optional_items", params)
		}
	}
	if !plan.PaymentIntentData.IsNull() && !plan.PaymentIntentData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentIntentData", plan.PaymentIntentData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_intent_data", params)
		}
	}
	if !plan.PaymentMethodCollection.IsNull() && !plan.PaymentMethodCollection.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethodCollection", "PaymentMethodCollection", plan.PaymentMethodCollection.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method_collection", params)
		}
	}
	if !plan.PaymentMethodOptions.IsNull() && !plan.PaymentMethodOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodOptions", plan.PaymentMethodOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method_options", params)
		}
	}
	if !plan.PaymentMethodTypes.IsNull() && !plan.PaymentMethodTypes.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodTypes", plan.PaymentMethodTypes) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method_types", params)
		}
	}
	if !plan.PhoneNumberCollection.IsNull() && !plan.PhoneNumberCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PhoneNumberCollection", plan.PhoneNumberCollection) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "phone_number_collection", params)
		}
	}
	if !plan.Restrictions.IsNull() && !plan.Restrictions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Restrictions", plan.Restrictions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "restrictions", params)
		}
	}
	if !plan.ShippingAddressCollection.IsNull() && !plan.ShippingAddressCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingAddressCollection", plan.ShippingAddressCollection) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping_address_collection", params)
		}
	}
	if !plan.ShippingOptions.IsNull() && !plan.ShippingOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingOptions", plan.ShippingOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping_options", params)
		}
	}
	if !plan.SubmitType.IsNull() && !plan.SubmitType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SubmitType", "SubmitType", plan.SubmitType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "submit_type", params)
		}
	}
	if !plan.SubscriptionData.IsNull() && !plan.SubscriptionData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SubscriptionData", plan.SubscriptionData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "subscription_data", params)
		}
	}
	if !plan.TaxIDCollection.IsNull() && !plan.TaxIDCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TaxIDCollection", plan.TaxIDCollection) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_id_collection", params)
		}
	}
	if !plan.TransferData.IsNull() && !plan.TransferData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TransferData", plan.TransferData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "transfer_data", params)
		}
	}

	return params, nil
}

func expandPaymentLinkUpdate(plan PaymentLinkResourceModel, state PaymentLinkResourceModel) (*stripe.PaymentLinkUpdateParams, error) {
	params := &stripe.PaymentLinkUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.AfterCompletion.Equal(state.AfterCompletion) && !plan.AfterCompletion.IsNull() && !plan.AfterCompletion.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AfterCompletion", plan.AfterCompletion) {
			if !plan.AfterCompletion.Equal(state.AfterCompletion) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "after_completion", params)
			}
		}
	}
	if !plan.AllowPromotionCodes.Equal(state.AllowPromotionCodes) && !plan.AllowPromotionCodes.IsNull() && !plan.AllowPromotionCodes.IsUnknown() {
		params.AllowPromotionCodes = stripe.Bool(plan.AllowPromotionCodes.ValueBool())
	}
	if !plan.AutomaticTax.Equal(state.AutomaticTax) && !plan.AutomaticTax.IsNull() && !plan.AutomaticTax.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AutomaticTax", plan.AutomaticTax) {
			if !plan.AutomaticTax.Equal(state.AutomaticTax) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "automatic_tax", params)
			}
		}
	}
	if !plan.BillingAddressCollection.Equal(state.BillingAddressCollection) && !plan.BillingAddressCollection.IsNull() && !plan.BillingAddressCollection.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "BillingAddressCollection", "BillingAddressCollection", plan.BillingAddressCollection.ValueString()) {
			if !plan.BillingAddressCollection.Equal(state.BillingAddressCollection) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "billing_address_collection", params)
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
	if !plan.CustomText.Equal(state.CustomText) && !plan.CustomText.IsNull() && !plan.CustomText.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomText", plan.CustomText) {
			if !plan.CustomText.Equal(state.CustomText) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "custom_text", params)
			}
		}
	}
	if !plan.CustomerCreation.Equal(state.CustomerCreation) && !plan.CustomerCreation.IsNull() && !plan.CustomerCreation.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerCreation", "CustomerCreation", plan.CustomerCreation.ValueString()) {
			if !plan.CustomerCreation.Equal(state.CustomerCreation) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "customer_creation", params)
			}
		}
	}
	if !plan.InactiveMessage.Equal(state.InactiveMessage) && !plan.InactiveMessage.IsNull() && !plan.InactiveMessage.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "InactiveMessage", "InactiveMessage", plan.InactiveMessage.ValueString()) {
			if !plan.InactiveMessage.Equal(state.InactiveMessage) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "inactive_message", params)
			}
		}
	}
	if !plan.InvoiceCreation.Equal(state.InvoiceCreation) && !plan.InvoiceCreation.IsNull() && !plan.InvoiceCreation.IsUnknown() {
		if !assignAttrValueToNamedField(params, "InvoiceCreation", plan.InvoiceCreation) {
			if !plan.InvoiceCreation.Equal(state.InvoiceCreation) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "invoice_creation", params)
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
	if !plan.NameCollection.Equal(state.NameCollection) && !plan.NameCollection.IsNull() && !plan.NameCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "NameCollection", plan.NameCollection) {
			if !plan.NameCollection.Equal(state.NameCollection) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "name_collection", params)
			}
		}
	}
	if !plan.OptionalItems.Equal(state.OptionalItems) && !plan.OptionalItems.IsNull() && !plan.OptionalItems.IsUnknown() {
		if !assignAttrValueToNamedField(params, "OptionalItems", plan.OptionalItems) {
			if !plan.OptionalItems.Equal(state.OptionalItems) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "optional_items", params)
			}
		}
	}
	if !plan.PaymentIntentData.Equal(state.PaymentIntentData) && !plan.PaymentIntentData.IsNull() && !plan.PaymentIntentData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentIntentData", plan.PaymentIntentData) {
			if !plan.PaymentIntentData.Equal(state.PaymentIntentData) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_intent_data", params)
			}
		}
	}
	if !plan.PaymentMethodCollection.Equal(state.PaymentMethodCollection) && !plan.PaymentMethodCollection.IsNull() && !plan.PaymentMethodCollection.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethodCollection", "PaymentMethodCollection", plan.PaymentMethodCollection.ValueString()) {
			if !plan.PaymentMethodCollection.Equal(state.PaymentMethodCollection) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_method_collection", params)
			}
		}
	}
	if !plan.PaymentMethodOptions.Equal(state.PaymentMethodOptions) && !plan.PaymentMethodOptions.IsNull() && !plan.PaymentMethodOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodOptions", plan.PaymentMethodOptions) {
			if !plan.PaymentMethodOptions.Equal(state.PaymentMethodOptions) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_method_options", params)
			}
		}
	}
	if !plan.PaymentMethodTypes.Equal(state.PaymentMethodTypes) && !plan.PaymentMethodTypes.IsNull() && !plan.PaymentMethodTypes.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodTypes", plan.PaymentMethodTypes) {
			if !plan.PaymentMethodTypes.Equal(state.PaymentMethodTypes) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_method_types", params)
			}
		}
	}
	if !plan.PhoneNumberCollection.Equal(state.PhoneNumberCollection) && !plan.PhoneNumberCollection.IsNull() && !plan.PhoneNumberCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PhoneNumberCollection", plan.PhoneNumberCollection) {
			if !plan.PhoneNumberCollection.Equal(state.PhoneNumberCollection) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "phone_number_collection", params)
			}
		}
	}
	if !plan.Restrictions.Equal(state.Restrictions) && !plan.Restrictions.IsNull() && !plan.Restrictions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Restrictions", plan.Restrictions) {
			if !plan.Restrictions.Equal(state.Restrictions) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "restrictions", params)
			}
		}
	}
	if !plan.ShippingAddressCollection.Equal(state.ShippingAddressCollection) && !plan.ShippingAddressCollection.IsNull() && !plan.ShippingAddressCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingAddressCollection", plan.ShippingAddressCollection) {
			if !plan.ShippingAddressCollection.Equal(state.ShippingAddressCollection) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "shipping_address_collection", params)
			}
		}
	}
	if !plan.SubmitType.Equal(state.SubmitType) && !plan.SubmitType.IsNull() && !plan.SubmitType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SubmitType", "SubmitType", plan.SubmitType.ValueString()) {
			if !plan.SubmitType.Equal(state.SubmitType) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "submit_type", params)
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
	if !plan.TaxIDCollection.Equal(state.TaxIDCollection) && !plan.TaxIDCollection.IsNull() && !plan.TaxIDCollection.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TaxIDCollection", plan.TaxIDCollection) {
			if !plan.TaxIDCollection.Equal(state.TaxIDCollection) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tax_id_collection", params)
			}
		}
	}

	return params, nil
}

func expandPaymentLinkPostCreateUpdate(plan PaymentLinkResourceModel, state PaymentLinkResourceModel) (*stripe.PaymentLinkUpdateParams, error) {
	params := &stripe.PaymentLinkUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}

	return params, nil
}

func flattenPaymentLink(obj *stripe.PaymentLink, state *PaymentLinkResourceModel) error {
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
		assignedAfterCompletion := false
		hadRawAfterCompletion := false
		if rawValueAfterCompletion, rawOk := plainValueAtPath(raw, "after_completion"); rawOk {
			hadRawAfterCompletion = true
			if rawValueAfterCompletion != nil {
				sourceAfterCompletion := applyConfiguredKeyedListShapes(rawValueAfterCompletion, attrValueToPlain(state.AfterCompletion))
				if valueAfterCompletion, err := flattenPlainValue(sourceAfterCompletion, types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_confirmation": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_message": types.StringType}}, "redirect": types.ObjectType{AttrTypes: map[string]attr.Type{"url": types.StringType}}, "type": types.StringType}}, "after_completion", "raw response"); err != nil {
					return err
				} else {
					if typedAfterCompletion, ok := valueAfterCompletion.(types.Object); ok {
						state.AfterCompletion = typedAfterCompletion
						assignedAfterCompletion = true
					}
				}
			}
		}
		if !assignedAfterCompletion {
			if !hasRaw {
				if responseValueAfterCompletion, ok := plainFromResponseField(obj, "AfterCompletion"); ok {
					sourceAfterCompletion := applyConfiguredKeyedListShapes(responseValueAfterCompletion, attrValueToPlain(state.AfterCompletion))
					if valueAfterCompletion, err := flattenPlainValue(
						sourceAfterCompletion,
						types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_confirmation": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_message": types.StringType}}, "redirect": types.ObjectType{AttrTypes: map[string]attr.Type{"url": types.StringType}}, "type": types.StringType}},
						"after_completion",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAfterCompletion, ok := valueAfterCompletion.(types.Object); ok {
							state.AfterCompletion = typedAfterCompletion
							assignedAfterCompletion = true
						}
					}
				}
			}
		}
		if !assignedAfterCompletion && hadRawAfterCompletion {
			if nullAfterCompletion, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"hosted_confirmation": types.ObjectType{AttrTypes: map[string]attr.Type{"custom_message": types.StringType}}, "redirect": types.ObjectType{AttrTypes: map[string]attr.Type{"url": types.StringType}}, "type": types.StringType}}); ok {
				if typedAfterCompletion, ok := nullAfterCompletion.(types.Object); ok {
					state.AfterCompletion = typedAfterCompletion
				}
			}
		}
	}
	{
		if rawValueAllowPromotionCodes, rawOk := plainValueAtPath(raw, "allow_promotion_codes"); rawOk {
			if valueAllowPromotionCodes, err := flattenPlainValue(rawValueAllowPromotionCodes, types.BoolType, "allow_promotion_codes", "raw response"); err != nil {
				return err
			} else {
				if typedAllowPromotionCodes, ok := valueAllowPromotionCodes.(types.Bool); ok {
					state.AllowPromotionCodes = typedAllowPromotionCodes
				}
			}
		} else if !hasRaw {
			if responseValueAllowPromotionCodes, ok := plainFromResponseField(obj, "AllowPromotionCodes"); ok {
				if valueAllowPromotionCodes, err := flattenPlainValue(responseValueAllowPromotionCodes, types.BoolType, "allow_promotion_codes", "response struct"); err != nil {
					return err
				} else {
					if typedAllowPromotionCodes, ok := valueAllowPromotionCodes.(types.Bool); ok {
						state.AllowPromotionCodes = typedAllowPromotionCodes
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
				if valueAutomaticTax, err := flattenPlainValue(sourceAutomaticTax, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "automatic_tax", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}},
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
			if nullAutomaticTax, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "liability": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}); ok {
				if typedAutomaticTax, ok := nullAutomaticTax.(types.Object); ok {
					state.AutomaticTax = typedAutomaticTax
				}
			}
		}
	}
	{
		if rawValueBillingAddressCollection, rawOk := plainValueAtPath(raw, "billing_address_collection"); rawOk {
			if valueBillingAddressCollection, err := flattenPlainValue(rawValueBillingAddressCollection, types.StringType, "billing_address_collection", "raw response"); err != nil {
				return err
			} else {
				if typedBillingAddressCollection, ok := valueBillingAddressCollection.(types.String); ok {
					state.BillingAddressCollection = typedBillingAddressCollection
				}
			}
		} else if !hasRaw {
			if responseValueBillingAddressCollection, ok := plainFromResponseField(obj, "BillingAddressCollection"); ok {
				if valueBillingAddressCollection, err := flattenPlainValue(responseValueBillingAddressCollection, types.StringType, "billing_address_collection", "response struct"); err != nil {
					return err
				} else {
					if typedBillingAddressCollection, ok := valueBillingAddressCollection.(types.String); ok {
						state.BillingAddressCollection = typedBillingAddressCollection
					}
				}
			}
		}
	}
	{
		assignedConsentCollection := false
		hadRawConsentCollection := false
		if rawValueConsentCollection, rawOk := plainValueAtPath(raw, "consent_collection"); rawOk {
			hadRawConsentCollection = true
			if rawValueConsentCollection != nil {
				sourceConsentCollection := applyConfiguredKeyedListShapes(rawValueConsentCollection, attrValueToPlain(state.ConsentCollection))
				if valueConsentCollection, err := flattenPlainValue(sourceConsentCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"payment_method_reuse_agreement": types.ObjectType{AttrTypes: map[string]attr.Type{"position": types.StringType}}, "promotions": types.StringType, "terms_of_service": types.StringType}}, "consent_collection", "raw response"); err != nil {
					return err
				} else {
					if typedConsentCollection, ok := valueConsentCollection.(types.Object); ok {
						state.ConsentCollection = typedConsentCollection
						assignedConsentCollection = true
					}
				}
			}
		}
		if !assignedConsentCollection {
			if !hasRaw {
				if responseValueConsentCollection, ok := plainFromResponseField(obj, "ConsentCollection"); ok {
					sourceConsentCollection := applyConfiguredKeyedListShapes(responseValueConsentCollection, attrValueToPlain(state.ConsentCollection))
					if valueConsentCollection, err := flattenPlainValue(
						sourceConsentCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"payment_method_reuse_agreement": types.ObjectType{AttrTypes: map[string]attr.Type{"position": types.StringType}}, "promotions": types.StringType, "terms_of_service": types.StringType}},
						"consent_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedConsentCollection, ok := valueConsentCollection.(types.Object); ok {
							state.ConsentCollection = typedConsentCollection
							assignedConsentCollection = true
						}
					}
				}
			}
		}
		if !assignedConsentCollection && hadRawConsentCollection {
			if nullConsentCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"payment_method_reuse_agreement": types.ObjectType{AttrTypes: map[string]attr.Type{"position": types.StringType}}, "promotions": types.StringType, "terms_of_service": types.StringType}}); ok {
				if typedConsentCollection, ok := nullConsentCollection.(types.Object); ok {
					state.ConsentCollection = typedConsentCollection
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
			if valueCustomFields, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueCustomFields, attrValueToPlain(state.CustomFields)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"dropdown": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"label": types.StringType, "value": types.StringType}}}}}, "key": types.StringType, "label": types.ObjectType{AttrTypes: map[string]attr.Type{"custom": types.StringType, "type": types.StringType}}, "numeric": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "maximum_length": types.Int64Type, "minimum_length": types.Int64Type}}, "optional": types.BoolType, "text": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "maximum_length": types.Int64Type, "minimum_length": types.Int64Type}}, "type": types.StringType}}}, "custom_fields", "raw response"); err != nil {
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
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"dropdown": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "options": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"label": types.StringType, "value": types.StringType}}}}}, "key": types.StringType, "label": types.ObjectType{AttrTypes: map[string]attr.Type{"custom": types.StringType, "type": types.StringType}}, "numeric": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "maximum_length": types.Int64Type, "minimum_length": types.Int64Type}}, "optional": types.BoolType, "text": types.ObjectType{AttrTypes: map[string]attr.Type{"default_value": types.StringType, "maximum_length": types.Int64Type, "minimum_length": types.Int64Type}}, "type": types.StringType}}},
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
		assignedCustomText := false
		hadRawCustomText := false
		if rawValueCustomText, rawOk := plainValueAtPath(raw, "custom_text"); rawOk {
			hadRawCustomText = true
			if rawValueCustomText != nil {
				sourceCustomText := applyConfiguredKeyedListShapes(rawValueCustomText, attrValueToPlain(state.CustomText))
				if valueCustomText, err := flattenPlainValue(sourceCustomText, types.ObjectType{AttrTypes: map[string]attr.Type{"after_submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "terms_of_service_acceptance": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}}}, "custom_text", "raw response"); err != nil {
					return err
				} else {
					if typedCustomText, ok := valueCustomText.(types.Object); ok {
						state.CustomText = typedCustomText
						assignedCustomText = true
					}
				}
			}
		}
		if !assignedCustomText {
			if !hasRaw {
				if responseValueCustomText, ok := plainFromResponseField(obj, "CustomText"); ok {
					sourceCustomText := applyConfiguredKeyedListShapes(responseValueCustomText, attrValueToPlain(state.CustomText))
					if valueCustomText, err := flattenPlainValue(
						sourceCustomText,
						types.ObjectType{AttrTypes: map[string]attr.Type{"after_submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "terms_of_service_acceptance": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}}},
						"custom_text",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCustomText, ok := valueCustomText.(types.Object); ok {
							state.CustomText = typedCustomText
							assignedCustomText = true
						}
					}
				}
			}
		}
		if !assignedCustomText && hadRawCustomText {
			if nullCustomText, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"after_submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "shipping_address": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "submit": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}, "terms_of_service_acceptance": types.ObjectType{AttrTypes: map[string]attr.Type{"message": types.StringType}}}}); ok {
				if typedCustomText, ok := nullCustomText.(types.Object); ok {
					state.CustomText = typedCustomText
				}
			}
		}
	}
	{
		if rawValueCustomerCreation, rawOk := plainValueAtPath(raw, "customer_creation"); rawOk {
			if valueCustomerCreation, err := flattenPlainValue(rawValueCustomerCreation, types.StringType, "customer_creation", "raw response"); err != nil {
				return err
			} else {
				if typedCustomerCreation, ok := valueCustomerCreation.(types.String); ok {
					state.CustomerCreation = typedCustomerCreation
				}
			}
		} else if !hasRaw {
			if responseValueCustomerCreation, ok := plainFromResponseField(obj, "CustomerCreation"); ok {
				if valueCustomerCreation, err := flattenPlainValue(responseValueCustomerCreation, types.StringType, "customer_creation", "response struct"); err != nil {
					return err
				} else {
					if typedCustomerCreation, ok := valueCustomerCreation.(types.String); ok {
						state.CustomerCreation = typedCustomerCreation
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
		if rawValueInactiveMessage, rawOk := plainValueAtPath(raw, "inactive_message"); rawOk {
			if valueInactiveMessage, err := flattenPlainValue(rawValueInactiveMessage, types.StringType, "inactive_message", "raw response"); err != nil {
				return err
			} else {
				if typedInactiveMessage, ok := valueInactiveMessage.(types.String); ok {
					state.InactiveMessage = typedInactiveMessage
				}
			}
		} else if !hasRaw {
			if responseValueInactiveMessage, ok := plainFromResponseField(obj, "InactiveMessage"); ok {
				if valueInactiveMessage, err := flattenPlainValue(responseValueInactiveMessage, types.StringType, "inactive_message", "response struct"); err != nil {
					return err
				} else {
					if typedInactiveMessage, ok := valueInactiveMessage.(types.String); ok {
						state.InactiveMessage = typedInactiveMessage
					}
				}
			}
		}
	}
	{
		assignedInvoiceCreation := false
		hadRawInvoiceCreation := false
		if rawValueInvoiceCreation, rawOk := plainValueAtPath(raw, "invoice_creation"); rawOk {
			hadRawInvoiceCreation = true
			if rawValueInvoiceCreation != nil {
				sourceInvoiceCreation := applyConfiguredKeyedListShapes(rawValueInvoiceCreation, attrValueToPlain(state.InvoiceCreation))
				if valueInvoiceCreation, err := flattenPlainValue(sourceInvoiceCreation, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "invoice_data": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "custom_fields": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "description": types.StringType, "footer": types.StringType, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "metadata": types.MapType{ElemType: types.StringType}, "rendering_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}}}}, "invoice_creation", "raw response"); err != nil {
					return err
				} else {
					if typedInvoiceCreation, ok := valueInvoiceCreation.(types.Object); ok {
						state.InvoiceCreation = typedInvoiceCreation
						assignedInvoiceCreation = true
					}
				}
			}
		}
		if !assignedInvoiceCreation {
			if !hasRaw {
				if responseValueInvoiceCreation, ok := plainFromResponseField(obj, "InvoiceCreation"); ok {
					sourceInvoiceCreation := applyConfiguredKeyedListShapes(responseValueInvoiceCreation, attrValueToPlain(state.InvoiceCreation))
					if valueInvoiceCreation, err := flattenPlainValue(
						sourceInvoiceCreation,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "invoice_data": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "custom_fields": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "description": types.StringType, "footer": types.StringType, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "metadata": types.MapType{ElemType: types.StringType}, "rendering_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}}}},
						"invoice_creation",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedInvoiceCreation, ok := valueInvoiceCreation.(types.Object); ok {
							state.InvoiceCreation = typedInvoiceCreation
							assignedInvoiceCreation = true
						}
					}
				}
			}
		}
		if !assignedInvoiceCreation && hadRawInvoiceCreation {
			if nullInvoiceCreation, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "invoice_data": types.ObjectType{AttrTypes: map[string]attr.Type{"account_tax_ids": types.ListType{ElemType: types.StringType}, "custom_fields": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType, "value": types.StringType}}}, "description": types.StringType, "footer": types.StringType, "issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}, "metadata": types.MapType{ElemType: types.StringType}, "rendering_options": types.ObjectType{AttrTypes: map[string]attr.Type{"amount_tax_display": types.StringType, "template": types.StringType}}}}}}); ok {
				if typedInvoiceCreation, ok := nullInvoiceCreation.(types.Object); ok {
					state.InvoiceCreation = typedInvoiceCreation
				}
			}
		}
	}
	{
		if rawValueLineItems, rawOk := plainValueAtPath(raw, "line_items"); rawOk {
			rawPlainLineItems := extractListObjectData(rawValueLineItems)
			if valueLineItems, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawPlainLineItems, attrValueToPlain(state.LineItems)), attrValueToPlain(state.LineItems)), attrValueToPlain(state.LineItems)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "price": types.StringType, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "product_data": types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "images": types.ListType{ElemType: types.StringType}, "metadata": types.MapType{ElemType: types.StringType}, "name": types.StringType, "tax_code": types.StringType, "unit_label": types.StringType}}, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}, "quantity": types.Int64Type}}}, "line_items", "raw response"); err != nil {
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
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "price": types.StringType, "price_data": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "product": types.StringType, "product_data": types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "images": types.ListType{ElemType: types.StringType}, "metadata": types.MapType{ElemType: types.StringType}, "name": types.StringType, "tax_code": types.StringType, "unit_label": types.StringType}}, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"interval": types.StringType, "interval_count": types.Int64Type}}, "tax_behavior": types.StringType, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type}}, "quantity": types.Int64Type}}},
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
		assignedNameCollection := false
		hadRawNameCollection := false
		if rawValueNameCollection, rawOk := plainValueAtPath(raw, "name_collection"); rawOk {
			hadRawNameCollection = true
			if rawValueNameCollection != nil {
				sourceNameCollection := applyConfiguredKeyedListShapes(rawValueNameCollection, attrValueToPlain(state.NameCollection))
				if valueNameCollection, err := flattenPlainValue(sourceNameCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"business": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}, "individual": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}}}, "name_collection", "raw response"); err != nil {
					return err
				} else {
					if typedNameCollection, ok := valueNameCollection.(types.Object); ok {
						state.NameCollection = typedNameCollection
						assignedNameCollection = true
					}
				}
			}
		}
		if !assignedNameCollection {
			if !hasRaw {
				if responseValueNameCollection, ok := plainFromResponseField(obj, "NameCollection"); ok {
					sourceNameCollection := applyConfiguredKeyedListShapes(responseValueNameCollection, attrValueToPlain(state.NameCollection))
					if valueNameCollection, err := flattenPlainValue(
						sourceNameCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"business": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}, "individual": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}}},
						"name_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedNameCollection, ok := valueNameCollection.(types.Object); ok {
							state.NameCollection = typedNameCollection
							assignedNameCollection = true
						}
					}
				}
			}
		}
		if !assignedNameCollection && hadRawNameCollection {
			if nullNameCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"business": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}, "individual": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "optional": types.BoolType}}}}); ok {
				if typedNameCollection, ok := nullNameCollection.(types.Object); ok {
					state.NameCollection = typedNameCollection
				}
			}
		}
	}
	{
		if state.OnBehalfOf.IsNull() || state.OnBehalfOf.IsUnknown() {
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
		if rawValueOptionalItems, rawOk := plainValueAtPath(raw, "optional_items"); rawOk {
			if valueOptionalItems, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueOptionalItems, attrValueToPlain(state.OptionalItems)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "price": types.StringType, "quantity": types.Int64Type}}}, "optional_items", "raw response"); err != nil {
				return err
			} else {
				if typedOptionalItems, ok := valueOptionalItems.(types.List); ok {
					state.OptionalItems = typedOptionalItems
				}
			}
		} else if !hasRaw {
			if responseValueOptionalItems, ok := plainFromResponseField(obj, "OptionalItems"); ok {
				if valueOptionalItems, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueOptionalItems, attrValueToPlain(state.OptionalItems)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "price": types.StringType, "quantity": types.Int64Type}}},
					"optional_items",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedOptionalItems, ok := valueOptionalItems.(types.List); ok {
						state.OptionalItems = typedOptionalItems
					}
				}
			}
		}
	}
	{
		assignedPaymentIntentData := false
		hadRawPaymentIntentData := false
		if rawValuePaymentIntentData, rawOk := plainValueAtPath(raw, "payment_intent_data"); rawOk {
			hadRawPaymentIntentData = true
			if rawValuePaymentIntentData != nil {
				sourcePaymentIntentData := applyConfiguredKeyedListShapes(rawValuePaymentIntentData, attrValueToPlain(state.PaymentIntentData))
				if valuePaymentIntentData, err := flattenPlainValue(sourcePaymentIntentData, types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "description": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "setup_future_usage": types.StringType, "statement_descriptor": types.StringType, "statement_descriptor_suffix": types.StringType, "transfer_group": types.StringType}}, "payment_intent_data", "raw response"); err != nil {
					return err
				} else {
					if typedPaymentIntentData, ok := valuePaymentIntentData.(types.Object); ok {
						state.PaymentIntentData = typedPaymentIntentData
						assignedPaymentIntentData = true
					}
				}
			}
		}
		if !assignedPaymentIntentData {
			if !hasRaw {
				if responseValuePaymentIntentData, ok := plainFromResponseField(obj, "PaymentIntentData"); ok {
					sourcePaymentIntentData := applyConfiguredKeyedListShapes(responseValuePaymentIntentData, attrValueToPlain(state.PaymentIntentData))
					if valuePaymentIntentData, err := flattenPlainValue(
						sourcePaymentIntentData,
						types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "description": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "setup_future_usage": types.StringType, "statement_descriptor": types.StringType, "statement_descriptor_suffix": types.StringType, "transfer_group": types.StringType}},
						"payment_intent_data",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPaymentIntentData, ok := valuePaymentIntentData.(types.Object); ok {
							state.PaymentIntentData = typedPaymentIntentData
							assignedPaymentIntentData = true
						}
					}
				}
			}
		}
		if !assignedPaymentIntentData && hadRawPaymentIntentData {
			if nullPaymentIntentData, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"capture_method": types.StringType, "description": types.StringType, "metadata": types.MapType{ElemType: types.StringType}, "setup_future_usage": types.StringType, "statement_descriptor": types.StringType, "statement_descriptor_suffix": types.StringType, "transfer_group": types.StringType}}); ok {
				if typedPaymentIntentData, ok := nullPaymentIntentData.(types.Object); ok {
					state.PaymentIntentData = typedPaymentIntentData
				}
			}
		}
	}
	{
		if rawValuePaymentMethodCollection, rawOk := plainValueAtPath(raw, "payment_method_collection"); rawOk {
			if valuePaymentMethodCollection, err := flattenPlainValue(rawValuePaymentMethodCollection, types.StringType, "payment_method_collection", "raw response"); err != nil {
				return err
			} else {
				if typedPaymentMethodCollection, ok := valuePaymentMethodCollection.(types.String); ok {
					state.PaymentMethodCollection = typedPaymentMethodCollection
				}
			}
		} else if !hasRaw {
			if responseValuePaymentMethodCollection, ok := plainFromResponseField(obj, "PaymentMethodCollection"); ok {
				if valuePaymentMethodCollection, err := flattenPlainValue(responseValuePaymentMethodCollection, types.StringType, "payment_method_collection", "response struct"); err != nil {
					return err
				} else {
					if typedPaymentMethodCollection, ok := valuePaymentMethodCollection.(types.String); ok {
						state.PaymentMethodCollection = typedPaymentMethodCollection
					}
				}
			}
		}
	}
	{
		assignedPaymentMethodOptions := false
		hadRawPaymentMethodOptions := false
		if rawValuePaymentMethodOptions, rawOk := plainValueAtPath(raw, "payment_method_options"); rawOk {
			hadRawPaymentMethodOptions = true
			if rawValuePaymentMethodOptions != nil {
				sourcePaymentMethodOptions := applyConfiguredKeyedListShapes(rawValuePaymentMethodOptions, attrValueToPlain(state.PaymentMethodOptions))
				if valuePaymentMethodOptions, err := flattenPlainValue(sourcePaymentMethodOptions, types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"restrictions": types.ObjectType{AttrTypes: map[string]attr.Type{"brands_blocked": types.ListType{ElemType: types.StringType}}}}}}}, "payment_method_options", "raw response"); err != nil {
					return err
				} else {
					if typedPaymentMethodOptions, ok := valuePaymentMethodOptions.(types.Object); ok {
						state.PaymentMethodOptions = typedPaymentMethodOptions
						assignedPaymentMethodOptions = true
					}
				}
			}
		}
		if !assignedPaymentMethodOptions {
			if !hasRaw {
				if responseValuePaymentMethodOptions, ok := plainFromResponseField(obj, "PaymentMethodOptions"); ok {
					sourcePaymentMethodOptions := applyConfiguredKeyedListShapes(responseValuePaymentMethodOptions, attrValueToPlain(state.PaymentMethodOptions))
					if valuePaymentMethodOptions, err := flattenPlainValue(
						sourcePaymentMethodOptions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"restrictions": types.ObjectType{AttrTypes: map[string]attr.Type{"brands_blocked": types.ListType{ElemType: types.StringType}}}}}}},
						"payment_method_options",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPaymentMethodOptions, ok := valuePaymentMethodOptions.(types.Object); ok {
							state.PaymentMethodOptions = typedPaymentMethodOptions
							assignedPaymentMethodOptions = true
						}
					}
				}
			}
		}
		if !assignedPaymentMethodOptions && hadRawPaymentMethodOptions {
			if nullPaymentMethodOptions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"card": types.ObjectType{AttrTypes: map[string]attr.Type{"restrictions": types.ObjectType{AttrTypes: map[string]attr.Type{"brands_blocked": types.ListType{ElemType: types.StringType}}}}}}}); ok {
				if typedPaymentMethodOptions, ok := nullPaymentMethodOptions.(types.Object); ok {
					state.PaymentMethodOptions = typedPaymentMethodOptions
				}
			}
		}
	}
	{
		if rawValuePaymentMethodTypes, rawOk := plainValueAtPath(raw, "payment_method_types"); rawOk {
			if valuePaymentMethodTypes, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValuePaymentMethodTypes, attrValueToPlain(state.PaymentMethodTypes)), types.ListType{ElemType: types.StringType}, "payment_method_types", "raw response"); err != nil {
				return err
			} else {
				if typedPaymentMethodTypes, ok := valuePaymentMethodTypes.(types.List); ok {
					state.PaymentMethodTypes = typedPaymentMethodTypes
				}
			}
		} else if !hasRaw {
			if responseValuePaymentMethodTypes, ok := plainFromResponseField(obj, "PaymentMethodTypes"); ok {
				if valuePaymentMethodTypes, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValuePaymentMethodTypes, attrValueToPlain(state.PaymentMethodTypes)),
					types.ListType{ElemType: types.StringType},
					"payment_method_types",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedPaymentMethodTypes, ok := valuePaymentMethodTypes.(types.List); ok {
						state.PaymentMethodTypes = typedPaymentMethodTypes
					}
				}
			}
		}
	}
	{
		assignedPhoneNumberCollection := false
		hadRawPhoneNumberCollection := false
		if rawValuePhoneNumberCollection, rawOk := plainValueAtPath(raw, "phone_number_collection"); rawOk {
			hadRawPhoneNumberCollection = true
			if rawValuePhoneNumberCollection != nil {
				sourcePhoneNumberCollection := applyConfiguredKeyedListShapes(rawValuePhoneNumberCollection, attrValueToPlain(state.PhoneNumberCollection))
				if valuePhoneNumberCollection, err := flattenPlainValue(sourcePhoneNumberCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "phone_number_collection", "raw response"); err != nil {
					return err
				} else {
					if typedPhoneNumberCollection, ok := valuePhoneNumberCollection.(types.Object); ok {
						state.PhoneNumberCollection = typedPhoneNumberCollection
						assignedPhoneNumberCollection = true
					}
				}
			}
		}
		if !assignedPhoneNumberCollection {
			if !hasRaw {
				if responseValuePhoneNumberCollection, ok := plainFromResponseField(obj, "PhoneNumberCollection"); ok {
					sourcePhoneNumberCollection := applyConfiguredKeyedListShapes(responseValuePhoneNumberCollection, attrValueToPlain(state.PhoneNumberCollection))
					if valuePhoneNumberCollection, err := flattenPlainValue(
						sourcePhoneNumberCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}},
						"phone_number_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPhoneNumberCollection, ok := valuePhoneNumberCollection.(types.Object); ok {
							state.PhoneNumberCollection = typedPhoneNumberCollection
							assignedPhoneNumberCollection = true
						}
					}
				}
			}
		}
		if !assignedPhoneNumberCollection && hadRawPhoneNumberCollection {
			if nullPhoneNumberCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}); ok {
				if typedPhoneNumberCollection, ok := nullPhoneNumberCollection.(types.Object); ok {
					state.PhoneNumberCollection = typedPhoneNumberCollection
				}
			}
		}
	}
	{
		assignedRestrictions := false
		hadRawRestrictions := false
		if rawValueRestrictions, rawOk := plainValueAtPath(raw, "restrictions"); rawOk {
			hadRawRestrictions = true
			if rawValueRestrictions != nil {
				sourceRestrictions := applyConfiguredKeyedListShapes(rawValueRestrictions, attrValueToPlain(state.Restrictions))
				if valueRestrictions, err := flattenPlainValue(sourceRestrictions, types.ObjectType{AttrTypes: map[string]attr.Type{"completed_sessions": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "limit": types.Int64Type}}}}, "restrictions", "raw response"); err != nil {
					return err
				} else {
					if typedRestrictions, ok := valueRestrictions.(types.Object); ok {
						state.Restrictions = typedRestrictions
						assignedRestrictions = true
					}
				}
			}
		}
		if !assignedRestrictions {
			if !hasRaw {
				if responseValueRestrictions, ok := plainFromResponseField(obj, "Restrictions"); ok {
					sourceRestrictions := applyConfiguredKeyedListShapes(responseValueRestrictions, attrValueToPlain(state.Restrictions))
					if valueRestrictions, err := flattenPlainValue(
						sourceRestrictions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"completed_sessions": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "limit": types.Int64Type}}}},
						"restrictions",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRestrictions, ok := valueRestrictions.(types.Object); ok {
							state.Restrictions = typedRestrictions
							assignedRestrictions = true
						}
					}
				}
			}
		}
		if !assignedRestrictions && hadRawRestrictions {
			if nullRestrictions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"completed_sessions": types.ObjectType{AttrTypes: map[string]attr.Type{"count": types.Int64Type, "limit": types.Int64Type}}}}); ok {
				if typedRestrictions, ok := nullRestrictions.(types.Object); ok {
					state.Restrictions = typedRestrictions
				}
			}
		}
	}
	{
		assignedShippingAddressCollection := false
		hadRawShippingAddressCollection := false
		if rawValueShippingAddressCollection, rawOk := plainValueAtPath(raw, "shipping_address_collection"); rawOk {
			hadRawShippingAddressCollection = true
			if rawValueShippingAddressCollection != nil {
				sourceShippingAddressCollection := applyConfiguredKeyedListShapes(rawValueShippingAddressCollection, attrValueToPlain(state.ShippingAddressCollection))
				if valueShippingAddressCollection, err := flattenPlainValue(sourceShippingAddressCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_countries": types.ListType{ElemType: types.StringType}}}, "shipping_address_collection", "raw response"); err != nil {
					return err
				} else {
					if typedShippingAddressCollection, ok := valueShippingAddressCollection.(types.Object); ok {
						state.ShippingAddressCollection = typedShippingAddressCollection
						assignedShippingAddressCollection = true
					}
				}
			}
		}
		if !assignedShippingAddressCollection {
			if !hasRaw {
				if responseValueShippingAddressCollection, ok := plainFromResponseField(obj, "ShippingAddressCollection"); ok {
					sourceShippingAddressCollection := applyConfiguredKeyedListShapes(responseValueShippingAddressCollection, attrValueToPlain(state.ShippingAddressCollection))
					if valueShippingAddressCollection, err := flattenPlainValue(
						sourceShippingAddressCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_countries": types.ListType{ElemType: types.StringType}}},
						"shipping_address_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedShippingAddressCollection, ok := valueShippingAddressCollection.(types.Object); ok {
							state.ShippingAddressCollection = typedShippingAddressCollection
							assignedShippingAddressCollection = true
						}
					}
				}
			}
		}
		if !assignedShippingAddressCollection && hadRawShippingAddressCollection {
			if nullShippingAddressCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_countries": types.ListType{ElemType: types.StringType}}}); ok {
				if typedShippingAddressCollection, ok := nullShippingAddressCollection.(types.Object); ok {
					state.ShippingAddressCollection = typedShippingAddressCollection
				}
			}
		}
	}
	{
		if rawValueShippingOptions, rawOk := plainValueAtPath(raw, "shipping_options"); rawOk {
			if valueShippingOptions, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueShippingOptions, attrValueToPlain(state.ShippingOptions)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"shipping_amount": types.Int64Type, "shipping_rate": types.StringType}}}, "shipping_options", "raw response"); err != nil {
				return err
			} else {
				if typedShippingOptions, ok := valueShippingOptions.(types.List); ok {
					state.ShippingOptions = typedShippingOptions
				}
			}
		} else if !hasRaw {
			if responseValueShippingOptions, ok := plainFromResponseField(obj, "ShippingOptions"); ok {
				if valueShippingOptions, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueShippingOptions, attrValueToPlain(state.ShippingOptions)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"shipping_amount": types.Int64Type, "shipping_rate": types.StringType}}},
					"shipping_options",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedShippingOptions, ok := valueShippingOptions.(types.List); ok {
						state.ShippingOptions = typedShippingOptions
					}
				}
			}
		}
	}
	{
		if rawValueSubmitType, rawOk := plainValueAtPath(raw, "submit_type"); rawOk {
			if valueSubmitType, err := flattenPlainValue(rawValueSubmitType, types.StringType, "submit_type", "raw response"); err != nil {
				return err
			} else {
				if typedSubmitType, ok := valueSubmitType.(types.String); ok {
					state.SubmitType = typedSubmitType
				}
			}
		} else if !hasRaw {
			if responseValueSubmitType, ok := plainFromResponseField(obj, "SubmitType"); ok {
				if valueSubmitType, err := flattenPlainValue(responseValueSubmitType, types.StringType, "submit_type", "response struct"); err != nil {
					return err
				} else {
					if typedSubmitType, ok := valueSubmitType.(types.String); ok {
						state.SubmitType = typedSubmitType
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
				if valueSubscriptionData, err := flattenPlainValue(sourceSubscriptionData, types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "invoice_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "metadata": types.MapType{ElemType: types.StringType}, "trial_period_days": types.Int64Type, "trial_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"end_behavior": types.ObjectType{AttrTypes: map[string]attr.Type{"missing_payment_method": types.StringType}}}}}}, "subscription_data", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "invoice_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "metadata": types.MapType{ElemType: types.StringType}, "trial_period_days": types.Int64Type, "trial_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"end_behavior": types.ObjectType{AttrTypes: map[string]attr.Type{"missing_payment_method": types.StringType}}}}}},
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
			if nullSubscriptionData, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"description": types.StringType, "invoice_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"issuer": types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "type": types.StringType}}}}, "metadata": types.MapType{ElemType: types.StringType}, "trial_period_days": types.Int64Type, "trial_settings": types.ObjectType{AttrTypes: map[string]attr.Type{"end_behavior": types.ObjectType{AttrTypes: map[string]attr.Type{"missing_payment_method": types.StringType}}}}}}); ok {
				if typedSubscriptionData, ok := nullSubscriptionData.(types.Object); ok {
					state.SubscriptionData = typedSubscriptionData
				}
			}
		}
	}
	{
		assignedTaxIDCollection := false
		hadRawTaxIDCollection := false
		if rawValueTaxIDCollection, rawOk := plainValueAtPath(raw, "tax_id_collection"); rawOk {
			hadRawTaxIDCollection = true
			if rawValueTaxIDCollection != nil {
				sourceTaxIDCollection := applyConfiguredKeyedListShapes(rawValueTaxIDCollection, attrValueToPlain(state.TaxIDCollection))
				if valueTaxIDCollection, err := flattenPlainValue(sourceTaxIDCollection, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "required": types.StringType}}, "tax_id_collection", "raw response"); err != nil {
					return err
				} else {
					if typedTaxIDCollection, ok := valueTaxIDCollection.(types.Object); ok {
						state.TaxIDCollection = typedTaxIDCollection
						assignedTaxIDCollection = true
					}
				}
			}
		}
		if !assignedTaxIDCollection {
			if !hasRaw {
				if responseValueTaxIDCollection, ok := plainFromResponseField(obj, "TaxIDCollection"); ok {
					sourceTaxIDCollection := applyConfiguredKeyedListShapes(responseValueTaxIDCollection, attrValueToPlain(state.TaxIDCollection))
					if valueTaxIDCollection, err := flattenPlainValue(
						sourceTaxIDCollection,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "required": types.StringType}},
						"tax_id_collection",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTaxIDCollection, ok := valueTaxIDCollection.(types.Object); ok {
							state.TaxIDCollection = typedTaxIDCollection
							assignedTaxIDCollection = true
						}
					}
				}
			}
		}
		if !assignedTaxIDCollection && hadRawTaxIDCollection {
			if nullTaxIDCollection, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "required": types.StringType}}); ok {
				if typedTaxIDCollection, ok := nullTaxIDCollection.(types.Object); ok {
					state.TaxIDCollection = typedTaxIDCollection
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
				if valueTransferData, err := flattenPlainValue(sourceTransferData, types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "destination": types.StringType}}, "transfer_data", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "destination": types.StringType}},
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
			if nullTransferData, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "destination": types.StringType}}); ok {
				if typedTransferData, ok := nullTransferData.(types.Object); ok {
					state.TransferData = typedTransferData
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
	return nil
}
