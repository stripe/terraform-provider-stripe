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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &PaymentMethodConfigurationResource{}

var _ resource.ResourceWithConfigure = &PaymentMethodConfigurationResource{}

var _ resource.ResourceWithImportState = &PaymentMethodConfigurationResource{}

func NewPaymentMethodConfigurationResource() resource.Resource {
	return &PaymentMethodConfigurationResource{}
}

type PaymentMethodConfigurationResource struct {
	client *stripe.Client
}

type PaymentMethodConfigurationResourceModel struct {
	Object           types.String `tfsdk:"object"`
	ACSSDebit        types.Object `tfsdk:"acss_debit"`
	Active           types.Bool   `tfsdk:"active"`
	Affirm           types.Object `tfsdk:"affirm"`
	AfterpayClearpay types.Object `tfsdk:"afterpay_clearpay"`
	Alipay           types.Object `tfsdk:"alipay"`
	Alma             types.Object `tfsdk:"alma"`
	AmazonPay        types.Object `tfsdk:"amazon_pay"`
	ApplePay         types.Object `tfsdk:"apple_pay"`
	Application      types.String `tfsdk:"application"`
	AUBECSDebit      types.Object `tfsdk:"au_becs_debit"`
	BACSDebit        types.Object `tfsdk:"bacs_debit"`
	Bancontact       types.Object `tfsdk:"bancontact"`
	Billie           types.Object `tfsdk:"billie"`
	Bizum            types.Object `tfsdk:"bizum"`
	BLIK             types.Object `tfsdk:"blik"`
	Boleto           types.Object `tfsdk:"boleto"`
	Card             types.Object `tfsdk:"card"`
	CartesBancaires  types.Object `tfsdk:"cartes_bancaires"`
	CashApp          types.Object `tfsdk:"cashapp"`
	Crypto           types.Object `tfsdk:"crypto"`
	CustomerBalance  types.Object `tfsdk:"customer_balance"`
	EPS              types.Object `tfsdk:"eps"`
	FPX              types.Object `tfsdk:"fpx"`
	Giropay          types.Object `tfsdk:"giropay"`
	GooglePay        types.Object `tfsdk:"google_pay"`
	Grabpay          types.Object `tfsdk:"grabpay"`
	ID               types.String `tfsdk:"id"`
	IDEAL            types.Object `tfsdk:"ideal"`
	IsDefault        types.Bool   `tfsdk:"is_default"`
	JCB              types.Object `tfsdk:"jcb"`
	KakaoPay         types.Object `tfsdk:"kakao_pay"`
	Klarna           types.Object `tfsdk:"klarna"`
	Konbini          types.Object `tfsdk:"konbini"`
	KrCard           types.Object `tfsdk:"kr_card"`
	Link             types.Object `tfsdk:"link"`
	Livemode         types.Bool   `tfsdk:"livemode"`
	MbWay            types.Object `tfsdk:"mb_way"`
	Mobilepay        types.Object `tfsdk:"mobilepay"`
	Multibanco       types.Object `tfsdk:"multibanco"`
	Name             types.String `tfsdk:"name"`
	NaverPay         types.Object `tfsdk:"naver_pay"`
	NzBankAccount    types.Object `tfsdk:"nz_bank_account"`
	OXXO             types.Object `tfsdk:"oxxo"`
	P24              types.Object `tfsdk:"p24"`
	Parent           types.String `tfsdk:"parent"`
	PayByBank        types.Object `tfsdk:"pay_by_bank"`
	Payco            types.Object `tfsdk:"payco"`
	PayNow           types.Object `tfsdk:"paynow"`
	Paypal           types.Object `tfsdk:"paypal"`
	Payto            types.Object `tfsdk:"payto"`
	Pix              types.Object `tfsdk:"pix"`
	PromptPay        types.Object `tfsdk:"promptpay"`
	RevolutPay       types.Object `tfsdk:"revolut_pay"`
	SamsungPay       types.Object `tfsdk:"samsung_pay"`
	Satispay         types.Object `tfsdk:"satispay"`
	Scalapay         types.Object `tfsdk:"scalapay"`
	SEPADebit        types.Object `tfsdk:"sepa_debit"`
	Sofort           types.Object `tfsdk:"sofort"`
	Sunbit           types.Object `tfsdk:"sunbit"`
	Swish            types.Object `tfsdk:"swish"`
	TWINT            types.Object `tfsdk:"twint"`
	Upi              types.Object `tfsdk:"upi"`
	USBankAccount    types.Object `tfsdk:"us_bank_account"`
	WeChatPay        types.Object `tfsdk:"wechat_pay"`
	Zip              types.Object `tfsdk:"zip"`
	ApplePayLater    types.Object `tfsdk:"apple_pay_later"`
}

func (r *PaymentMethodConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PaymentMethodConfigurationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_payment_method_configuration"
}

func (r *PaymentMethodConfigurationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "PaymentMethodConfigurations control which payment methods are displayed to your customers when you don't explicitly specify payment method types. You can have multiple configurations with different sets of payment methods for different scenarios.\n\nThere are two types of PaymentMethodConfigurations. Which is used depends on the [charge type](https://docs.stripe.com/connect/charges):\n\n**Direct** configurations apply to payments created on your account, including Connect destination charges, Connect separate charges and transfers, and payments not involving Connect.\n\n**Child** configurations apply to payments created on your connected accounts using direct charges, and charges with the on_behalf_of parameter.\n\nChild configurations have a `parent` that sets default values and controls which settings connected accounts may override. You can specify a parent ID at payment time, and Stripe will automatically resolve the connected account’s associated child configuration. Parent configurations are [managed in the dashboard](https://dashboard.stripe.com/settings/payment_methods/connected_accounts) and are not available in this API.\n\nRelated guides:\n- [Payment Method Configurations API](https://docs.stripe.com/connect/payment-method-configurations)\n- [Multiple configurations on dynamic payment methods](https://docs.stripe.com/payments/multiple-payment-method-configs)\n- [Multiple configurations for your Connect accounts](https://docs.stripe.com/connect/multiple-payment-method-configurations)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("payment_method_configuration")},
			},
			"acss_debit": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the configuration can be used for new payments.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"affirm": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"afterpay_clearpay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"alipay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"alma": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"amazon_pay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"apple_pay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "For child configs, the Connect application associated with the configuration.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"au_becs_debit": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"bacs_debit": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"bancontact": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"billie": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"bizum": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"blik": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"boleto": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"card": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
						},
					},
				},
			},
			"cartes_bancaires": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"cashapp": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"crypto": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"customer_balance": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"eps": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"fpx": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"giropay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"google_pay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"grabpay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ideal": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"is_default": schema.BoolAttribute{
				Computed:      true,
				Description:   "The default configuration is used whenever a payment method configuration is not specified.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"jcb": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"kakao_pay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"klarna": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"konbini": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"kr_card": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"link": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
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
			"mb_way": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"mobilepay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"multibanco": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The configuration's name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"naver_pay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"nz_bank_account": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"oxxo": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"p24": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"parent": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "For child configs, the configuration's parent configuration.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"pay_by_bank": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"payco": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"paynow": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"paypal": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"payto": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"pix": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"promptpay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"revolut_pay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"samsung_pay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"satispay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"scalapay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"sepa_debit": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"sofort": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"sunbit": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"swish": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"twint": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"upi": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"us_bank_account": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"wechat_pay": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"zip": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"available": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this payment method may be offered at checkout. True if `display_preference` is `on` and the payment method's capability is active.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"display_preference": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"overridable": schema.BoolAttribute{
								Computed:      true,
								Description:   "For child configs, whether or not the account's preference will be observed. If `false`, the parent configuration's default is used.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"preference": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The account's display preference.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("none", "off", "on")},
							},
							"value": schema.StringAttribute{
								Computed:      true,
								Description:   "The effective display preference value.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("off", "on")},
							},
						},
					},
				},
			},
			"apple_pay_later": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Apple Pay Later, a payment method for customers to buy now and pay later, gives your customers a way to split purchases into four installments across six weeks.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"display_preference": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "Whether or not the payment method should be displayed.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"preference": schema.StringAttribute{
								Optional:    true,
								Description: "The account's preference for whether or not to display this payment method.",
								WriteOnly:   true,
							},
						},
					},
				},
			},
		},
	}
}

func (r *PaymentMethodConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PaymentMethodConfigurationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PaymentMethodConfigurationResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"ApplePayLater"}, []string{"ApplePayLater", "display_preference"}, []string{"ApplePayLater", "display_preference", "preference"}})

	params, err := expandPaymentMethodConfigurationCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentMethodConfiguration create params", err.Error())
		return
	}

	obj, err := r.client.V1PaymentMethodConfigurations.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating PaymentMethodConfiguration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentMethodConfigurations.B, r.client.V1PaymentMethodConfigurations.Key, stripe.FormatURLPath("/v1/payment_method_configurations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentMethodConfiguration create raw response", err.Error())
		return
	}

	var createdState PaymentMethodConfigurationResourceModel
	if err := flattenPaymentMethodConfiguration(obj, &createdState); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentMethodConfiguration create response", err.Error())
		return
	}
	normalizeUnknownValues(&createdState)

	diffPlan := plan
	diffCreatedState := createdState

	postCreateParams, err := expandPaymentMethodConfigurationPostCreateUpdate(diffPlan, diffCreatedState)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentMethodConfiguration post-create update params", err.Error())
		return
	}

	if paramsHaveValues(postCreateParams) {
		obj, err = r.client.V1PaymentMethodConfigurations.Update(ctx, createdState.ID.ValueString(), postCreateParams)
		if err != nil {
			resp.Diagnostics.AddError("Error finalizing PaymentMethodConfiguration after create", err.Error())
			return
		}
		if err := ensureRawResponse(obj, r.client.V1PaymentMethodConfigurations.B, r.client.V1PaymentMethodConfigurations.Key, stripe.FormatURLPath("/v1/payment_method_configurations/%s", obj.ID), nil); err != nil {
			resp.Diagnostics.AddError("Error hydrating PaymentMethodConfiguration post-create update raw response", err.Error())
			return
		}
	}

	if err := flattenPaymentMethodConfiguration(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentMethodConfiguration create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"ApplePayLater"}, []string{"ApplePayLater", "display_preference"}, []string{"ApplePayLater", "display_preference", "preference"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentMethodConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState PaymentMethodConfigurationResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PaymentMethodConfigurationResourceModel
	state = priorState

	obj, err := r.client.V1PaymentMethodConfigurations.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading PaymentMethodConfiguration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentMethodConfigurations.B, r.client.V1PaymentMethodConfigurations.Key, stripe.FormatURLPath("/v1/payment_method_configurations/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentMethodConfiguration raw response", err.Error())
		return
	}

	if err := flattenPaymentMethodConfiguration(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentMethodConfiguration read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"ApplePayLater"}, []string{"ApplePayLater", "display_preference"}, []string{"ApplePayLater", "display_preference", "preference"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentMethodConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PaymentMethodConfigurationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PaymentMethodConfigurationResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"ApplePayLater"}, []string{"ApplePayLater", "display_preference"}, []string{"ApplePayLater", "display_preference", "preference"}})

	var state PaymentMethodConfigurationResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandPaymentMethodConfigurationUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentMethodConfiguration update params", err.Error())
		return
	}

	obj, err := r.client.V1PaymentMethodConfigurations.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating PaymentMethodConfiguration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentMethodConfigurations.B, r.client.V1PaymentMethodConfigurations.Key, stripe.FormatURLPath("/v1/payment_method_configurations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentMethodConfiguration update raw response", err.Error())
		return
	}

	if err := flattenPaymentMethodConfiguration(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentMethodConfiguration update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"ApplePayLater"}, []string{"ApplePayLater", "display_preference"}, []string{"ApplePayLater", "display_preference", "preference"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentMethodConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PaymentMethodConfigurationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.Active.IsNull() && !state.Active.IsUnknown() && !state.Active.ValueBool() {
		return
	}

	params := &stripe.PaymentMethodConfigurationUpdateParams{}
	activeField := reflect.ValueOf(params).Elem().FieldByName("Active")
	if activeField.IsValid() && activeField.CanSet() {
		if activeField.Kind() == reflect.Pointer && activeField.Type().Elem().Kind() == reflect.Bool {
			activeField.Set(reflect.ValueOf(stripe.Bool(false)))
		} else if activeField.Kind() == reflect.Bool {
			activeField.SetBool(false)
		}
	}

	_, err := r.client.V1PaymentMethodConfigurations.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating PaymentMethodConfiguration", err.Error())
		return
	}
}

func (r *PaymentMethodConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandPaymentMethodConfigurationCreate(plan PaymentMethodConfigurationResourceModel) (*stripe.PaymentMethodConfigurationCreateParams, error) {
	params := &stripe.PaymentMethodConfigurationCreateParams{}

	if !plan.ACSSDebit.IsNull() && !plan.ACSSDebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ACSSDebit", plan.ACSSDebit) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "acss_debit", params)
		}
	}
	if !plan.Affirm.IsNull() && !plan.Affirm.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Affirm", plan.Affirm) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "affirm", params)
		}
	}
	if !plan.AfterpayClearpay.IsNull() && !plan.AfterpayClearpay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AfterpayClearpay", plan.AfterpayClearpay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "afterpay_clearpay", params)
		}
	}
	if !plan.Alipay.IsNull() && !plan.Alipay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Alipay", plan.Alipay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "alipay", params)
		}
	}
	if !plan.Alma.IsNull() && !plan.Alma.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Alma", plan.Alma) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "alma", params)
		}
	}
	if !plan.AmazonPay.IsNull() && !plan.AmazonPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AmazonPay", plan.AmazonPay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "amazon_pay", params)
		}
	}
	if !plan.ApplePay.IsNull() && !plan.ApplePay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ApplePay", plan.ApplePay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "apple_pay", params)
		}
	}
	if !plan.AUBECSDebit.IsNull() && !plan.AUBECSDebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AUBECSDebit", plan.AUBECSDebit) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "au_becs_debit", params)
		}
	}
	if !plan.BACSDebit.IsNull() && !plan.BACSDebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BACSDebit", plan.BACSDebit) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "bacs_debit", params)
		}
	}
	if !plan.Bancontact.IsNull() && !plan.Bancontact.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Bancontact", plan.Bancontact) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "bancontact", params)
		}
	}
	if !plan.Billie.IsNull() && !plan.Billie.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Billie", plan.Billie) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billie", params)
		}
	}
	if !plan.Bizum.IsNull() && !plan.Bizum.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Bizum", plan.Bizum) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "bizum", params)
		}
	}
	if !plan.BLIK.IsNull() && !plan.BLIK.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BLIK", plan.BLIK) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "blik", params)
		}
	}
	if !plan.Boleto.IsNull() && !plan.Boleto.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Boleto", plan.Boleto) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "boleto", params)
		}
	}
	if !plan.Card.IsNull() && !plan.Card.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Card", plan.Card) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "card", params)
		}
	}
	if !plan.CartesBancaires.IsNull() && !plan.CartesBancaires.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CartesBancaires", plan.CartesBancaires) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "cartes_bancaires", params)
		}
	}
	if !plan.CashApp.IsNull() && !plan.CashApp.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CashApp", plan.CashApp) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "cashapp", params)
		}
	}
	if !plan.Crypto.IsNull() && !plan.Crypto.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Crypto", plan.Crypto) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "crypto", params)
		}
	}
	if !plan.CustomerBalance.IsNull() && !plan.CustomerBalance.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomerBalance", plan.CustomerBalance) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer_balance", params)
		}
	}
	if !plan.EPS.IsNull() && !plan.EPS.IsUnknown() {
		if !assignAttrValueToNamedField(params, "EPS", plan.EPS) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "eps", params)
		}
	}
	if !plan.FPX.IsNull() && !plan.FPX.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FPX", plan.FPX) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "fpx", params)
		}
	}
	if !plan.Giropay.IsNull() && !plan.Giropay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Giropay", plan.Giropay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "giropay", params)
		}
	}
	if !plan.GooglePay.IsNull() && !plan.GooglePay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "GooglePay", plan.GooglePay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "google_pay", params)
		}
	}
	if !plan.Grabpay.IsNull() && !plan.Grabpay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Grabpay", plan.Grabpay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "grabpay", params)
		}
	}
	if !plan.IDEAL.IsNull() && !plan.IDEAL.IsUnknown() {
		if !assignAttrValueToNamedField(params, "IDEAL", plan.IDEAL) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "ideal", params)
		}
	}
	if !plan.JCB.IsNull() && !plan.JCB.IsUnknown() {
		if !assignAttrValueToNamedField(params, "JCB", plan.JCB) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "jcb", params)
		}
	}
	if !plan.KakaoPay.IsNull() && !plan.KakaoPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "KakaoPay", plan.KakaoPay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "kakao_pay", params)
		}
	}
	if !plan.Klarna.IsNull() && !plan.Klarna.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Klarna", plan.Klarna) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "klarna", params)
		}
	}
	if !plan.Konbini.IsNull() && !plan.Konbini.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Konbini", plan.Konbini) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "konbini", params)
		}
	}
	if !plan.KrCard.IsNull() && !plan.KrCard.IsUnknown() {
		if !assignAttrValueToNamedField(params, "KrCard", plan.KrCard) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "kr_card", params)
		}
	}
	if !plan.Link.IsNull() && !plan.Link.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Link", plan.Link) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "link", params)
		}
	}
	if !plan.MbWay.IsNull() && !plan.MbWay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "MbWay", plan.MbWay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "mb_way", params)
		}
	}
	if !plan.Mobilepay.IsNull() && !plan.Mobilepay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Mobilepay", plan.Mobilepay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "mobilepay", params)
		}
	}
	if !plan.Multibanco.IsNull() && !plan.Multibanco.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Multibanco", plan.Multibanco) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "multibanco", params)
		}
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Name", "Name", plan.Name.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "name", params)
		}
	}
	if !plan.NaverPay.IsNull() && !plan.NaverPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "NaverPay", plan.NaverPay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "naver_pay", params)
		}
	}
	if !plan.NzBankAccount.IsNull() && !plan.NzBankAccount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "NzBankAccount", plan.NzBankAccount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "nz_bank_account", params)
		}
	}
	if !plan.OXXO.IsNull() && !plan.OXXO.IsUnknown() {
		if !assignAttrValueToNamedField(params, "OXXO", plan.OXXO) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "oxxo", params)
		}
	}
	if !plan.P24.IsNull() && !plan.P24.IsUnknown() {
		if !assignAttrValueToNamedField(params, "P24", plan.P24) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "p24", params)
		}
	}
	if !plan.Parent.IsNull() && !plan.Parent.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Parent", "Parent", plan.Parent.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "parent", params)
		}
	}
	if !plan.PayByBank.IsNull() && !plan.PayByBank.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PayByBank", plan.PayByBank) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "pay_by_bank", params)
		}
	}
	if !plan.Payco.IsNull() && !plan.Payco.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Payco", plan.Payco) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payco", params)
		}
	}
	if !plan.PayNow.IsNull() && !plan.PayNow.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PayNow", plan.PayNow) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "paynow", params)
		}
	}
	if !plan.Paypal.IsNull() && !plan.Paypal.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Paypal", plan.Paypal) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "paypal", params)
		}
	}
	if !plan.Payto.IsNull() && !plan.Payto.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Payto", plan.Payto) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payto", params)
		}
	}
	if !plan.Pix.IsNull() && !plan.Pix.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Pix", plan.Pix) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "pix", params)
		}
	}
	if !plan.PromptPay.IsNull() && !plan.PromptPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PromptPay", plan.PromptPay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "promptpay", params)
		}
	}
	if !plan.RevolutPay.IsNull() && !plan.RevolutPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RevolutPay", plan.RevolutPay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "revolut_pay", params)
		}
	}
	if !plan.SamsungPay.IsNull() && !plan.SamsungPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SamsungPay", plan.SamsungPay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "samsung_pay", params)
		}
	}
	if !plan.Satispay.IsNull() && !plan.Satispay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Satispay", plan.Satispay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "satispay", params)
		}
	}
	if !plan.Scalapay.IsNull() && !plan.Scalapay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Scalapay", plan.Scalapay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "scalapay", params)
		}
	}
	if !plan.SEPADebit.IsNull() && !plan.SEPADebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SEPADebit", plan.SEPADebit) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "sepa_debit", params)
		}
	}
	if !plan.Sofort.IsNull() && !plan.Sofort.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Sofort", plan.Sofort) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "sofort", params)
		}
	}
	if !plan.Sunbit.IsNull() && !plan.Sunbit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Sunbit", plan.Sunbit) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "sunbit", params)
		}
	}
	if !plan.Swish.IsNull() && !plan.Swish.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Swish", plan.Swish) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "swish", params)
		}
	}
	if !plan.TWINT.IsNull() && !plan.TWINT.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TWINT", plan.TWINT) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "twint", params)
		}
	}
	if !plan.Upi.IsNull() && !plan.Upi.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Upi", plan.Upi) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "upi", params)
		}
	}
	if !plan.USBankAccount.IsNull() && !plan.USBankAccount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "USBankAccount", plan.USBankAccount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "us_bank_account", params)
		}
	}
	if !plan.WeChatPay.IsNull() && !plan.WeChatPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "WeChatPay", plan.WeChatPay) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "wechat_pay", params)
		}
	}
	if !plan.Zip.IsNull() && !plan.Zip.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Zip", plan.Zip) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "zip", params)
		}
	}
	if !plan.ApplePayLater.IsNull() && !plan.ApplePayLater.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ApplePayLater", plan.ApplePayLater) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "apple_pay_later", params)
		}
	}

	return params, nil
}

func expandPaymentMethodConfigurationUpdate(plan PaymentMethodConfigurationResourceModel, state PaymentMethodConfigurationResourceModel) (*stripe.PaymentMethodConfigurationUpdateParams, error) {
	params := &stripe.PaymentMethodConfigurationUpdateParams{}

	if !plan.ACSSDebit.Equal(state.ACSSDebit) && !plan.ACSSDebit.IsNull() && !plan.ACSSDebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ACSSDebit", plan.ACSSDebit) {
			if !plan.ACSSDebit.Equal(state.ACSSDebit) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "acss_debit", params)
			}
		}
	}
	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.Affirm.Equal(state.Affirm) && !plan.Affirm.IsNull() && !plan.Affirm.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Affirm", plan.Affirm) {
			if !plan.Affirm.Equal(state.Affirm) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "affirm", params)
			}
		}
	}
	if !plan.AfterpayClearpay.Equal(state.AfterpayClearpay) && !plan.AfterpayClearpay.IsNull() && !plan.AfterpayClearpay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AfterpayClearpay", plan.AfterpayClearpay) {
			if !plan.AfterpayClearpay.Equal(state.AfterpayClearpay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "afterpay_clearpay", params)
			}
		}
	}
	if !plan.Alipay.Equal(state.Alipay) && !plan.Alipay.IsNull() && !plan.Alipay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Alipay", plan.Alipay) {
			if !plan.Alipay.Equal(state.Alipay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "alipay", params)
			}
		}
	}
	if !plan.Alma.Equal(state.Alma) && !plan.Alma.IsNull() && !plan.Alma.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Alma", plan.Alma) {
			if !plan.Alma.Equal(state.Alma) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "alma", params)
			}
		}
	}
	if !plan.AmazonPay.Equal(state.AmazonPay) && !plan.AmazonPay.IsNull() && !plan.AmazonPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AmazonPay", plan.AmazonPay) {
			if !plan.AmazonPay.Equal(state.AmazonPay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "amazon_pay", params)
			}
		}
	}
	if !plan.ApplePay.Equal(state.ApplePay) && !plan.ApplePay.IsNull() && !plan.ApplePay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ApplePay", plan.ApplePay) {
			if !plan.ApplePay.Equal(state.ApplePay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "apple_pay", params)
			}
		}
	}
	if !plan.AUBECSDebit.Equal(state.AUBECSDebit) && !plan.AUBECSDebit.IsNull() && !plan.AUBECSDebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AUBECSDebit", plan.AUBECSDebit) {
			if !plan.AUBECSDebit.Equal(state.AUBECSDebit) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "au_becs_debit", params)
			}
		}
	}
	if !plan.BACSDebit.Equal(state.BACSDebit) && !plan.BACSDebit.IsNull() && !plan.BACSDebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BACSDebit", plan.BACSDebit) {
			if !plan.BACSDebit.Equal(state.BACSDebit) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "bacs_debit", params)
			}
		}
	}
	if !plan.Bancontact.Equal(state.Bancontact) && !plan.Bancontact.IsNull() && !plan.Bancontact.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Bancontact", plan.Bancontact) {
			if !plan.Bancontact.Equal(state.Bancontact) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "bancontact", params)
			}
		}
	}
	if !plan.Billie.Equal(state.Billie) && !plan.Billie.IsNull() && !plan.Billie.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Billie", plan.Billie) {
			if !plan.Billie.Equal(state.Billie) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "billie", params)
			}
		}
	}
	if !plan.Bizum.Equal(state.Bizum) && !plan.Bizum.IsNull() && !plan.Bizum.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Bizum", plan.Bizum) {
			if !plan.Bizum.Equal(state.Bizum) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "bizum", params)
			}
		}
	}
	if !plan.BLIK.Equal(state.BLIK) && !plan.BLIK.IsNull() && !plan.BLIK.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BLIK", plan.BLIK) {
			if !plan.BLIK.Equal(state.BLIK) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "blik", params)
			}
		}
	}
	if !plan.Boleto.Equal(state.Boleto) && !plan.Boleto.IsNull() && !plan.Boleto.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Boleto", plan.Boleto) {
			if !plan.Boleto.Equal(state.Boleto) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "boleto", params)
			}
		}
	}
	if !plan.Card.Equal(state.Card) && !plan.Card.IsNull() && !plan.Card.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Card", plan.Card) {
			if !plan.Card.Equal(state.Card) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "card", params)
			}
		}
	}
	if !plan.CartesBancaires.Equal(state.CartesBancaires) && !plan.CartesBancaires.IsNull() && !plan.CartesBancaires.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CartesBancaires", plan.CartesBancaires) {
			if !plan.CartesBancaires.Equal(state.CartesBancaires) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "cartes_bancaires", params)
			}
		}
	}
	if !plan.CashApp.Equal(state.CashApp) && !plan.CashApp.IsNull() && !plan.CashApp.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CashApp", plan.CashApp) {
			if !plan.CashApp.Equal(state.CashApp) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "cashapp", params)
			}
		}
	}
	if !plan.Crypto.Equal(state.Crypto) && !plan.Crypto.IsNull() && !plan.Crypto.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Crypto", plan.Crypto) {
			if !plan.Crypto.Equal(state.Crypto) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "crypto", params)
			}
		}
	}
	if !plan.CustomerBalance.Equal(state.CustomerBalance) && !plan.CustomerBalance.IsNull() && !plan.CustomerBalance.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomerBalance", plan.CustomerBalance) {
			if !plan.CustomerBalance.Equal(state.CustomerBalance) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "customer_balance", params)
			}
		}
	}
	if !plan.EPS.Equal(state.EPS) && !plan.EPS.IsNull() && !plan.EPS.IsUnknown() {
		if !assignAttrValueToNamedField(params, "EPS", plan.EPS) {
			if !plan.EPS.Equal(state.EPS) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "eps", params)
			}
		}
	}
	if !plan.FPX.Equal(state.FPX) && !plan.FPX.IsNull() && !plan.FPX.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FPX", plan.FPX) {
			if !plan.FPX.Equal(state.FPX) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "fpx", params)
			}
		}
	}
	if !plan.Giropay.Equal(state.Giropay) && !plan.Giropay.IsNull() && !plan.Giropay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Giropay", plan.Giropay) {
			if !plan.Giropay.Equal(state.Giropay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "giropay", params)
			}
		}
	}
	if !plan.GooglePay.Equal(state.GooglePay) && !plan.GooglePay.IsNull() && !plan.GooglePay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "GooglePay", plan.GooglePay) {
			if !plan.GooglePay.Equal(state.GooglePay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "google_pay", params)
			}
		}
	}
	if !plan.Grabpay.Equal(state.Grabpay) && !plan.Grabpay.IsNull() && !plan.Grabpay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Grabpay", plan.Grabpay) {
			if !plan.Grabpay.Equal(state.Grabpay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "grabpay", params)
			}
		}
	}
	if !plan.IDEAL.Equal(state.IDEAL) && !plan.IDEAL.IsNull() && !plan.IDEAL.IsUnknown() {
		if !assignAttrValueToNamedField(params, "IDEAL", plan.IDEAL) {
			if !plan.IDEAL.Equal(state.IDEAL) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "ideal", params)
			}
		}
	}
	if !plan.JCB.Equal(state.JCB) && !plan.JCB.IsNull() && !plan.JCB.IsUnknown() {
		if !assignAttrValueToNamedField(params, "JCB", plan.JCB) {
			if !plan.JCB.Equal(state.JCB) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "jcb", params)
			}
		}
	}
	if !plan.KakaoPay.Equal(state.KakaoPay) && !plan.KakaoPay.IsNull() && !plan.KakaoPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "KakaoPay", plan.KakaoPay) {
			if !plan.KakaoPay.Equal(state.KakaoPay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "kakao_pay", params)
			}
		}
	}
	if !plan.Klarna.Equal(state.Klarna) && !plan.Klarna.IsNull() && !plan.Klarna.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Klarna", plan.Klarna) {
			if !plan.Klarna.Equal(state.Klarna) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "klarna", params)
			}
		}
	}
	if !plan.Konbini.Equal(state.Konbini) && !plan.Konbini.IsNull() && !plan.Konbini.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Konbini", plan.Konbini) {
			if !plan.Konbini.Equal(state.Konbini) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "konbini", params)
			}
		}
	}
	if !plan.KrCard.Equal(state.KrCard) && !plan.KrCard.IsNull() && !plan.KrCard.IsUnknown() {
		if !assignAttrValueToNamedField(params, "KrCard", plan.KrCard) {
			if !plan.KrCard.Equal(state.KrCard) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "kr_card", params)
			}
		}
	}
	if !plan.Link.Equal(state.Link) && !plan.Link.IsNull() && !plan.Link.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Link", plan.Link) {
			if !plan.Link.Equal(state.Link) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "link", params)
			}
		}
	}
	if !plan.MbWay.Equal(state.MbWay) && !plan.MbWay.IsNull() && !plan.MbWay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "MbWay", plan.MbWay) {
			if !plan.MbWay.Equal(state.MbWay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "mb_way", params)
			}
		}
	}
	if !plan.Mobilepay.Equal(state.Mobilepay) && !plan.Mobilepay.IsNull() && !plan.Mobilepay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Mobilepay", plan.Mobilepay) {
			if !plan.Mobilepay.Equal(state.Mobilepay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "mobilepay", params)
			}
		}
	}
	if !plan.Multibanco.Equal(state.Multibanco) && !plan.Multibanco.IsNull() && !plan.Multibanco.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Multibanco", plan.Multibanco) {
			if !plan.Multibanco.Equal(state.Multibanco) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "multibanco", params)
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
	if !plan.NaverPay.Equal(state.NaverPay) && !plan.NaverPay.IsNull() && !plan.NaverPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "NaverPay", plan.NaverPay) {
			if !plan.NaverPay.Equal(state.NaverPay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "naver_pay", params)
			}
		}
	}
	if !plan.NzBankAccount.Equal(state.NzBankAccount) && !plan.NzBankAccount.IsNull() && !plan.NzBankAccount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "NzBankAccount", plan.NzBankAccount) {
			if !plan.NzBankAccount.Equal(state.NzBankAccount) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "nz_bank_account", params)
			}
		}
	}
	if !plan.OXXO.Equal(state.OXXO) && !plan.OXXO.IsNull() && !plan.OXXO.IsUnknown() {
		if !assignAttrValueToNamedField(params, "OXXO", plan.OXXO) {
			if !plan.OXXO.Equal(state.OXXO) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "oxxo", params)
			}
		}
	}
	if !plan.P24.Equal(state.P24) && !plan.P24.IsNull() && !plan.P24.IsUnknown() {
		if !assignAttrValueToNamedField(params, "P24", plan.P24) {
			if !plan.P24.Equal(state.P24) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "p24", params)
			}
		}
	}
	if !plan.PayByBank.Equal(state.PayByBank) && !plan.PayByBank.IsNull() && !plan.PayByBank.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PayByBank", plan.PayByBank) {
			if !plan.PayByBank.Equal(state.PayByBank) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "pay_by_bank", params)
			}
		}
	}
	if !plan.Payco.Equal(state.Payco) && !plan.Payco.IsNull() && !plan.Payco.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Payco", plan.Payco) {
			if !plan.Payco.Equal(state.Payco) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payco", params)
			}
		}
	}
	if !plan.PayNow.Equal(state.PayNow) && !plan.PayNow.IsNull() && !plan.PayNow.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PayNow", plan.PayNow) {
			if !plan.PayNow.Equal(state.PayNow) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "paynow", params)
			}
		}
	}
	if !plan.Paypal.Equal(state.Paypal) && !plan.Paypal.IsNull() && !plan.Paypal.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Paypal", plan.Paypal) {
			if !plan.Paypal.Equal(state.Paypal) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "paypal", params)
			}
		}
	}
	if !plan.Payto.Equal(state.Payto) && !plan.Payto.IsNull() && !plan.Payto.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Payto", plan.Payto) {
			if !plan.Payto.Equal(state.Payto) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payto", params)
			}
		}
	}
	if !plan.Pix.Equal(state.Pix) && !plan.Pix.IsNull() && !plan.Pix.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Pix", plan.Pix) {
			if !plan.Pix.Equal(state.Pix) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "pix", params)
			}
		}
	}
	if !plan.PromptPay.Equal(state.PromptPay) && !plan.PromptPay.IsNull() && !plan.PromptPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PromptPay", plan.PromptPay) {
			if !plan.PromptPay.Equal(state.PromptPay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "promptpay", params)
			}
		}
	}
	if !plan.RevolutPay.Equal(state.RevolutPay) && !plan.RevolutPay.IsNull() && !plan.RevolutPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RevolutPay", plan.RevolutPay) {
			if !plan.RevolutPay.Equal(state.RevolutPay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "revolut_pay", params)
			}
		}
	}
	if !plan.SamsungPay.Equal(state.SamsungPay) && !plan.SamsungPay.IsNull() && !plan.SamsungPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SamsungPay", plan.SamsungPay) {
			if !plan.SamsungPay.Equal(state.SamsungPay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "samsung_pay", params)
			}
		}
	}
	if !plan.Satispay.Equal(state.Satispay) && !plan.Satispay.IsNull() && !plan.Satispay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Satispay", plan.Satispay) {
			if !plan.Satispay.Equal(state.Satispay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "satispay", params)
			}
		}
	}
	if !plan.Scalapay.Equal(state.Scalapay) && !plan.Scalapay.IsNull() && !plan.Scalapay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Scalapay", plan.Scalapay) {
			if !plan.Scalapay.Equal(state.Scalapay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "scalapay", params)
			}
		}
	}
	if !plan.SEPADebit.Equal(state.SEPADebit) && !plan.SEPADebit.IsNull() && !plan.SEPADebit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SEPADebit", plan.SEPADebit) {
			if !plan.SEPADebit.Equal(state.SEPADebit) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "sepa_debit", params)
			}
		}
	}
	if !plan.Sofort.Equal(state.Sofort) && !plan.Sofort.IsNull() && !plan.Sofort.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Sofort", plan.Sofort) {
			if !plan.Sofort.Equal(state.Sofort) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "sofort", params)
			}
		}
	}
	if !plan.Sunbit.Equal(state.Sunbit) && !plan.Sunbit.IsNull() && !plan.Sunbit.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Sunbit", plan.Sunbit) {
			if !plan.Sunbit.Equal(state.Sunbit) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "sunbit", params)
			}
		}
	}
	if !plan.Swish.Equal(state.Swish) && !plan.Swish.IsNull() && !plan.Swish.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Swish", plan.Swish) {
			if !plan.Swish.Equal(state.Swish) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "swish", params)
			}
		}
	}
	if !plan.TWINT.Equal(state.TWINT) && !plan.TWINT.IsNull() && !plan.TWINT.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TWINT", plan.TWINT) {
			if !plan.TWINT.Equal(state.TWINT) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "twint", params)
			}
		}
	}
	if !plan.Upi.Equal(state.Upi) && !plan.Upi.IsNull() && !plan.Upi.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Upi", plan.Upi) {
			if !plan.Upi.Equal(state.Upi) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "upi", params)
			}
		}
	}
	if !plan.USBankAccount.Equal(state.USBankAccount) && !plan.USBankAccount.IsNull() && !plan.USBankAccount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "USBankAccount", plan.USBankAccount) {
			if !plan.USBankAccount.Equal(state.USBankAccount) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "us_bank_account", params)
			}
		}
	}
	if !plan.WeChatPay.Equal(state.WeChatPay) && !plan.WeChatPay.IsNull() && !plan.WeChatPay.IsUnknown() {
		if !assignAttrValueToNamedField(params, "WeChatPay", plan.WeChatPay) {
			if !plan.WeChatPay.Equal(state.WeChatPay) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "wechat_pay", params)
			}
		}
	}
	if !plan.Zip.Equal(state.Zip) && !plan.Zip.IsNull() && !plan.Zip.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Zip", plan.Zip) {
			if !plan.Zip.Equal(state.Zip) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "zip", params)
			}
		}
	}
	if !plan.ApplePayLater.Equal(state.ApplePayLater) && !plan.ApplePayLater.IsNull() && !plan.ApplePayLater.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ApplePayLater", plan.ApplePayLater) {
			if !plan.ApplePayLater.Equal(state.ApplePayLater) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "apple_pay_later", params)
			}
		}
	}

	return params, nil
}

func expandPaymentMethodConfigurationPostCreateUpdate(plan PaymentMethodConfigurationResourceModel, state PaymentMethodConfigurationResourceModel) (*stripe.PaymentMethodConfigurationUpdateParams, error) {
	params := &stripe.PaymentMethodConfigurationUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}

	return params, nil
}

func flattenPaymentMethodConfiguration(obj *stripe.PaymentMethodConfiguration, state *PaymentMethodConfigurationResourceModel) error {
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
		assignedACSSDebit := false
		hadRawACSSDebit := false
		if rawValueACSSDebit, rawOk := plainValueAtPath(raw, "acss_debit"); rawOk {
			hadRawACSSDebit = true
			if rawValueACSSDebit != nil {
				sourceACSSDebit := applyConfiguredKeyedListShapes(rawValueACSSDebit, attrValueToPlain(state.ACSSDebit))
				if valueACSSDebit, err := flattenPlainValue(sourceACSSDebit, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "acss_debit", "raw response"); err != nil {
					return err
				} else {
					if typedACSSDebit, ok := valueACSSDebit.(types.Object); ok {
						state.ACSSDebit = typedACSSDebit
						assignedACSSDebit = true
					}
				}
			}
		}
		if !assignedACSSDebit {
			if !hasRaw {
				if responseValueACSSDebit, ok := plainFromResponseField(obj, "ACSSDebit"); ok {
					sourceACSSDebit := applyConfiguredKeyedListShapes(responseValueACSSDebit, attrValueToPlain(state.ACSSDebit))
					if valueACSSDebit, err := flattenPlainValue(
						sourceACSSDebit,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"acss_debit",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedACSSDebit, ok := valueACSSDebit.(types.Object); ok {
							state.ACSSDebit = typedACSSDebit
							assignedACSSDebit = true
						}
					}
				}
			}
		}
		if !assignedACSSDebit && hadRawACSSDebit {
			if nullACSSDebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedACSSDebit, ok := nullACSSDebit.(types.Object); ok {
					state.ACSSDebit = typedACSSDebit
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
		assignedAffirm := false
		hadRawAffirm := false
		if rawValueAffirm, rawOk := plainValueAtPath(raw, "affirm"); rawOk {
			hadRawAffirm = true
			if rawValueAffirm != nil {
				sourceAffirm := applyConfiguredKeyedListShapes(rawValueAffirm, attrValueToPlain(state.Affirm))
				if valueAffirm, err := flattenPlainValue(sourceAffirm, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "affirm", "raw response"); err != nil {
					return err
				} else {
					if typedAffirm, ok := valueAffirm.(types.Object); ok {
						state.Affirm = typedAffirm
						assignedAffirm = true
					}
				}
			}
		}
		if !assignedAffirm {
			if !hasRaw {
				if responseValueAffirm, ok := plainFromResponseField(obj, "Affirm"); ok {
					sourceAffirm := applyConfiguredKeyedListShapes(responseValueAffirm, attrValueToPlain(state.Affirm))
					if valueAffirm, err := flattenPlainValue(
						sourceAffirm,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"affirm",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAffirm, ok := valueAffirm.(types.Object); ok {
							state.Affirm = typedAffirm
							assignedAffirm = true
						}
					}
				}
			}
		}
		if !assignedAffirm && hadRawAffirm {
			if nullAffirm, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedAffirm, ok := nullAffirm.(types.Object); ok {
					state.Affirm = typedAffirm
				}
			}
		}
	}
	{
		assignedAfterpayClearpay := false
		hadRawAfterpayClearpay := false
		if rawValueAfterpayClearpay, rawOk := plainValueAtPath(raw, "afterpay_clearpay"); rawOk {
			hadRawAfterpayClearpay = true
			if rawValueAfterpayClearpay != nil {
				sourceAfterpayClearpay := applyConfiguredKeyedListShapes(rawValueAfterpayClearpay, attrValueToPlain(state.AfterpayClearpay))
				if valueAfterpayClearpay, err := flattenPlainValue(sourceAfterpayClearpay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "afterpay_clearpay", "raw response"); err != nil {
					return err
				} else {
					if typedAfterpayClearpay, ok := valueAfterpayClearpay.(types.Object); ok {
						state.AfterpayClearpay = typedAfterpayClearpay
						assignedAfterpayClearpay = true
					}
				}
			}
		}
		if !assignedAfterpayClearpay {
			if !hasRaw {
				if responseValueAfterpayClearpay, ok := plainFromResponseField(obj, "AfterpayClearpay"); ok {
					sourceAfterpayClearpay := applyConfiguredKeyedListShapes(responseValueAfterpayClearpay, attrValueToPlain(state.AfterpayClearpay))
					if valueAfterpayClearpay, err := flattenPlainValue(
						sourceAfterpayClearpay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"afterpay_clearpay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAfterpayClearpay, ok := valueAfterpayClearpay.(types.Object); ok {
							state.AfterpayClearpay = typedAfterpayClearpay
							assignedAfterpayClearpay = true
						}
					}
				}
			}
		}
		if !assignedAfterpayClearpay && hadRawAfterpayClearpay {
			if nullAfterpayClearpay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedAfterpayClearpay, ok := nullAfterpayClearpay.(types.Object); ok {
					state.AfterpayClearpay = typedAfterpayClearpay
				}
			}
		}
	}
	{
		assignedAlipay := false
		hadRawAlipay := false
		if rawValueAlipay, rawOk := plainValueAtPath(raw, "alipay"); rawOk {
			hadRawAlipay = true
			if rawValueAlipay != nil {
				sourceAlipay := applyConfiguredKeyedListShapes(rawValueAlipay, attrValueToPlain(state.Alipay))
				if valueAlipay, err := flattenPlainValue(sourceAlipay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "alipay", "raw response"); err != nil {
					return err
				} else {
					if typedAlipay, ok := valueAlipay.(types.Object); ok {
						state.Alipay = typedAlipay
						assignedAlipay = true
					}
				}
			}
		}
		if !assignedAlipay {
			if !hasRaw {
				if responseValueAlipay, ok := plainFromResponseField(obj, "Alipay"); ok {
					sourceAlipay := applyConfiguredKeyedListShapes(responseValueAlipay, attrValueToPlain(state.Alipay))
					if valueAlipay, err := flattenPlainValue(
						sourceAlipay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"alipay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAlipay, ok := valueAlipay.(types.Object); ok {
							state.Alipay = typedAlipay
							assignedAlipay = true
						}
					}
				}
			}
		}
		if !assignedAlipay && hadRawAlipay {
			if nullAlipay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedAlipay, ok := nullAlipay.(types.Object); ok {
					state.Alipay = typedAlipay
				}
			}
		}
	}
	{
		assignedAlma := false
		hadRawAlma := false
		if rawValueAlma, rawOk := plainValueAtPath(raw, "alma"); rawOk {
			hadRawAlma = true
			if rawValueAlma != nil {
				sourceAlma := applyConfiguredKeyedListShapes(rawValueAlma, attrValueToPlain(state.Alma))
				if valueAlma, err := flattenPlainValue(sourceAlma, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "alma", "raw response"); err != nil {
					return err
				} else {
					if typedAlma, ok := valueAlma.(types.Object); ok {
						state.Alma = typedAlma
						assignedAlma = true
					}
				}
			}
		}
		if !assignedAlma {
			if !hasRaw {
				if responseValueAlma, ok := plainFromResponseField(obj, "Alma"); ok {
					sourceAlma := applyConfiguredKeyedListShapes(responseValueAlma, attrValueToPlain(state.Alma))
					if valueAlma, err := flattenPlainValue(
						sourceAlma,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"alma",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAlma, ok := valueAlma.(types.Object); ok {
							state.Alma = typedAlma
							assignedAlma = true
						}
					}
				}
			}
		}
		if !assignedAlma && hadRawAlma {
			if nullAlma, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedAlma, ok := nullAlma.(types.Object); ok {
					state.Alma = typedAlma
				}
			}
		}
	}
	{
		assignedAmazonPay := false
		hadRawAmazonPay := false
		if rawValueAmazonPay, rawOk := plainValueAtPath(raw, "amazon_pay"); rawOk {
			hadRawAmazonPay = true
			if rawValueAmazonPay != nil {
				sourceAmazonPay := applyConfiguredKeyedListShapes(rawValueAmazonPay, attrValueToPlain(state.AmazonPay))
				if valueAmazonPay, err := flattenPlainValue(sourceAmazonPay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "amazon_pay", "raw response"); err != nil {
					return err
				} else {
					if typedAmazonPay, ok := valueAmazonPay.(types.Object); ok {
						state.AmazonPay = typedAmazonPay
						assignedAmazonPay = true
					}
				}
			}
		}
		if !assignedAmazonPay {
			if !hasRaw {
				if responseValueAmazonPay, ok := plainFromResponseField(obj, "AmazonPay"); ok {
					sourceAmazonPay := applyConfiguredKeyedListShapes(responseValueAmazonPay, attrValueToPlain(state.AmazonPay))
					if valueAmazonPay, err := flattenPlainValue(
						sourceAmazonPay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"amazon_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAmazonPay, ok := valueAmazonPay.(types.Object); ok {
							state.AmazonPay = typedAmazonPay
							assignedAmazonPay = true
						}
					}
				}
			}
		}
		if !assignedAmazonPay && hadRawAmazonPay {
			if nullAmazonPay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedAmazonPay, ok := nullAmazonPay.(types.Object); ok {
					state.AmazonPay = typedAmazonPay
				}
			}
		}
	}
	{
		assignedApplePay := false
		hadRawApplePay := false
		if rawValueApplePay, rawOk := plainValueAtPath(raw, "apple_pay"); rawOk {
			hadRawApplePay = true
			if rawValueApplePay != nil {
				sourceApplePay := applyConfiguredKeyedListShapes(rawValueApplePay, attrValueToPlain(state.ApplePay))
				if valueApplePay, err := flattenPlainValue(sourceApplePay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "apple_pay", "raw response"); err != nil {
					return err
				} else {
					if typedApplePay, ok := valueApplePay.(types.Object); ok {
						state.ApplePay = typedApplePay
						assignedApplePay = true
					}
				}
			}
		}
		if !assignedApplePay {
			if !hasRaw {
				if responseValueApplePay, ok := plainFromResponseField(obj, "ApplePay"); ok {
					sourceApplePay := applyConfiguredKeyedListShapes(responseValueApplePay, attrValueToPlain(state.ApplePay))
					if valueApplePay, err := flattenPlainValue(
						sourceApplePay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"apple_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedApplePay, ok := valueApplePay.(types.Object); ok {
							state.ApplePay = typedApplePay
							assignedApplePay = true
						}
					}
				}
			}
		}
		if !assignedApplePay && hadRawApplePay {
			if nullApplePay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedApplePay, ok := nullApplePay.(types.Object); ok {
					state.ApplePay = typedApplePay
				}
			}
		}
	}
	{
		if rawValueApplication, rawOk := plainValueAtPath(raw, "application"); rawOk {
			if valueApplication, err := flattenPlainValue(rawValueApplication, types.StringType, "application", "raw response"); err != nil {
				return err
			} else {
				if typedApplication, ok := valueApplication.(types.String); ok {
					state.Application = typedApplication
				}
			}
		} else if !hasRaw {
			if responseValueApplication, ok := plainFromResponseField(obj, "Application"); ok {
				if valueApplication, err := flattenPlainValue(responseValueApplication, types.StringType, "application", "response struct"); err != nil {
					return err
				} else {
					if typedApplication, ok := valueApplication.(types.String); ok {
						state.Application = typedApplication
					}
				}
			}
		}
	}
	{
		assignedAUBECSDebit := false
		hadRawAUBECSDebit := false
		if rawValueAUBECSDebit, rawOk := plainValueAtPath(raw, "au_becs_debit"); rawOk {
			hadRawAUBECSDebit = true
			if rawValueAUBECSDebit != nil {
				sourceAUBECSDebit := applyConfiguredKeyedListShapes(rawValueAUBECSDebit, attrValueToPlain(state.AUBECSDebit))
				if valueAUBECSDebit, err := flattenPlainValue(sourceAUBECSDebit, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "au_becs_debit", "raw response"); err != nil {
					return err
				} else {
					if typedAUBECSDebit, ok := valueAUBECSDebit.(types.Object); ok {
						state.AUBECSDebit = typedAUBECSDebit
						assignedAUBECSDebit = true
					}
				}
			}
		}
		if !assignedAUBECSDebit {
			if !hasRaw {
				if responseValueAUBECSDebit, ok := plainFromResponseField(obj, "AUBECSDebit"); ok {
					sourceAUBECSDebit := applyConfiguredKeyedListShapes(responseValueAUBECSDebit, attrValueToPlain(state.AUBECSDebit))
					if valueAUBECSDebit, err := flattenPlainValue(
						sourceAUBECSDebit,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"au_becs_debit",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAUBECSDebit, ok := valueAUBECSDebit.(types.Object); ok {
							state.AUBECSDebit = typedAUBECSDebit
							assignedAUBECSDebit = true
						}
					}
				}
			}
		}
		if !assignedAUBECSDebit && hadRawAUBECSDebit {
			if nullAUBECSDebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedAUBECSDebit, ok := nullAUBECSDebit.(types.Object); ok {
					state.AUBECSDebit = typedAUBECSDebit
				}
			}
		}
	}
	{
		assignedBACSDebit := false
		hadRawBACSDebit := false
		if rawValueBACSDebit, rawOk := plainValueAtPath(raw, "bacs_debit"); rawOk {
			hadRawBACSDebit = true
			if rawValueBACSDebit != nil {
				sourceBACSDebit := applyConfiguredKeyedListShapes(rawValueBACSDebit, attrValueToPlain(state.BACSDebit))
				if valueBACSDebit, err := flattenPlainValue(sourceBACSDebit, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "bacs_debit", "raw response"); err != nil {
					return err
				} else {
					if typedBACSDebit, ok := valueBACSDebit.(types.Object); ok {
						state.BACSDebit = typedBACSDebit
						assignedBACSDebit = true
					}
				}
			}
		}
		if !assignedBACSDebit {
			if !hasRaw {
				if responseValueBACSDebit, ok := plainFromResponseField(obj, "BACSDebit"); ok {
					sourceBACSDebit := applyConfiguredKeyedListShapes(responseValueBACSDebit, attrValueToPlain(state.BACSDebit))
					if valueBACSDebit, err := flattenPlainValue(
						sourceBACSDebit,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"bacs_debit",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBACSDebit, ok := valueBACSDebit.(types.Object); ok {
							state.BACSDebit = typedBACSDebit
							assignedBACSDebit = true
						}
					}
				}
			}
		}
		if !assignedBACSDebit && hadRawBACSDebit {
			if nullBACSDebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedBACSDebit, ok := nullBACSDebit.(types.Object); ok {
					state.BACSDebit = typedBACSDebit
				}
			}
		}
	}
	{
		assignedBancontact := false
		hadRawBancontact := false
		if rawValueBancontact, rawOk := plainValueAtPath(raw, "bancontact"); rawOk {
			hadRawBancontact = true
			if rawValueBancontact != nil {
				sourceBancontact := applyConfiguredKeyedListShapes(rawValueBancontact, attrValueToPlain(state.Bancontact))
				if valueBancontact, err := flattenPlainValue(sourceBancontact, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "bancontact", "raw response"); err != nil {
					return err
				} else {
					if typedBancontact, ok := valueBancontact.(types.Object); ok {
						state.Bancontact = typedBancontact
						assignedBancontact = true
					}
				}
			}
		}
		if !assignedBancontact {
			if !hasRaw {
				if responseValueBancontact, ok := plainFromResponseField(obj, "Bancontact"); ok {
					sourceBancontact := applyConfiguredKeyedListShapes(responseValueBancontact, attrValueToPlain(state.Bancontact))
					if valueBancontact, err := flattenPlainValue(
						sourceBancontact,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"bancontact",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBancontact, ok := valueBancontact.(types.Object); ok {
							state.Bancontact = typedBancontact
							assignedBancontact = true
						}
					}
				}
			}
		}
		if !assignedBancontact && hadRawBancontact {
			if nullBancontact, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedBancontact, ok := nullBancontact.(types.Object); ok {
					state.Bancontact = typedBancontact
				}
			}
		}
	}
	{
		assignedBillie := false
		hadRawBillie := false
		if rawValueBillie, rawOk := plainValueAtPath(raw, "billie"); rawOk {
			hadRawBillie = true
			if rawValueBillie != nil {
				sourceBillie := applyConfiguredKeyedListShapes(rawValueBillie, attrValueToPlain(state.Billie))
				if valueBillie, err := flattenPlainValue(sourceBillie, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "billie", "raw response"); err != nil {
					return err
				} else {
					if typedBillie, ok := valueBillie.(types.Object); ok {
						state.Billie = typedBillie
						assignedBillie = true
					}
				}
			}
		}
		if !assignedBillie {
			if !hasRaw {
				if responseValueBillie, ok := plainFromResponseField(obj, "Billie"); ok {
					sourceBillie := applyConfiguredKeyedListShapes(responseValueBillie, attrValueToPlain(state.Billie))
					if valueBillie, err := flattenPlainValue(
						sourceBillie,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"billie",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBillie, ok := valueBillie.(types.Object); ok {
							state.Billie = typedBillie
							assignedBillie = true
						}
					}
				}
			}
		}
		if !assignedBillie && hadRawBillie {
			if nullBillie, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedBillie, ok := nullBillie.(types.Object); ok {
					state.Billie = typedBillie
				}
			}
		}
	}
	{
		assignedBizum := false
		hadRawBizum := false
		if rawValueBizum, rawOk := plainValueAtPath(raw, "bizum"); rawOk {
			hadRawBizum = true
			if rawValueBizum != nil {
				sourceBizum := applyConfiguredKeyedListShapes(rawValueBizum, attrValueToPlain(state.Bizum))
				if valueBizum, err := flattenPlainValue(sourceBizum, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "bizum", "raw response"); err != nil {
					return err
				} else {
					if typedBizum, ok := valueBizum.(types.Object); ok {
						state.Bizum = typedBizum
						assignedBizum = true
					}
				}
			}
		}
		if !assignedBizum {
			if !hasRaw {
				if responseValueBizum, ok := plainFromResponseField(obj, "Bizum"); ok {
					sourceBizum := applyConfiguredKeyedListShapes(responseValueBizum, attrValueToPlain(state.Bizum))
					if valueBizum, err := flattenPlainValue(
						sourceBizum,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"bizum",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBizum, ok := valueBizum.(types.Object); ok {
							state.Bizum = typedBizum
							assignedBizum = true
						}
					}
				}
			}
		}
		if !assignedBizum && hadRawBizum {
			if nullBizum, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedBizum, ok := nullBizum.(types.Object); ok {
					state.Bizum = typedBizum
				}
			}
		}
	}
	{
		assignedBLIK := false
		hadRawBLIK := false
		if rawValueBLIK, rawOk := plainValueAtPath(raw, "blik"); rawOk {
			hadRawBLIK = true
			if rawValueBLIK != nil {
				sourceBLIK := applyConfiguredKeyedListShapes(rawValueBLIK, attrValueToPlain(state.BLIK))
				if valueBLIK, err := flattenPlainValue(sourceBLIK, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "blik", "raw response"); err != nil {
					return err
				} else {
					if typedBLIK, ok := valueBLIK.(types.Object); ok {
						state.BLIK = typedBLIK
						assignedBLIK = true
					}
				}
			}
		}
		if !assignedBLIK {
			if !hasRaw {
				if responseValueBLIK, ok := plainFromResponseField(obj, "BLIK"); ok {
					sourceBLIK := applyConfiguredKeyedListShapes(responseValueBLIK, attrValueToPlain(state.BLIK))
					if valueBLIK, err := flattenPlainValue(
						sourceBLIK,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"blik",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBLIK, ok := valueBLIK.(types.Object); ok {
							state.BLIK = typedBLIK
							assignedBLIK = true
						}
					}
				}
			}
		}
		if !assignedBLIK && hadRawBLIK {
			if nullBLIK, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedBLIK, ok := nullBLIK.(types.Object); ok {
					state.BLIK = typedBLIK
				}
			}
		}
	}
	{
		assignedBoleto := false
		hadRawBoleto := false
		if rawValueBoleto, rawOk := plainValueAtPath(raw, "boleto"); rawOk {
			hadRawBoleto = true
			if rawValueBoleto != nil {
				sourceBoleto := applyConfiguredKeyedListShapes(rawValueBoleto, attrValueToPlain(state.Boleto))
				if valueBoleto, err := flattenPlainValue(sourceBoleto, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "boleto", "raw response"); err != nil {
					return err
				} else {
					if typedBoleto, ok := valueBoleto.(types.Object); ok {
						state.Boleto = typedBoleto
						assignedBoleto = true
					}
				}
			}
		}
		if !assignedBoleto {
			if !hasRaw {
				if responseValueBoleto, ok := plainFromResponseField(obj, "Boleto"); ok {
					sourceBoleto := applyConfiguredKeyedListShapes(responseValueBoleto, attrValueToPlain(state.Boleto))
					if valueBoleto, err := flattenPlainValue(
						sourceBoleto,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"boleto",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBoleto, ok := valueBoleto.(types.Object); ok {
							state.Boleto = typedBoleto
							assignedBoleto = true
						}
					}
				}
			}
		}
		if !assignedBoleto && hadRawBoleto {
			if nullBoleto, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedBoleto, ok := nullBoleto.(types.Object); ok {
					state.Boleto = typedBoleto
				}
			}
		}
	}
	{
		assignedCard := false
		hadRawCard := false
		if rawValueCard, rawOk := plainValueAtPath(raw, "card"); rawOk {
			hadRawCard = true
			if rawValueCard != nil {
				sourceCard := applyConfiguredKeyedListShapes(rawValueCard, attrValueToPlain(state.Card))
				if valueCard, err := flattenPlainValue(sourceCard, types.ObjectType{AttrTypes: map[string]attr.Type{"display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType}}}}, "card", "raw response"); err != nil {
					return err
				} else {
					if typedCard, ok := valueCard.(types.Object); ok {
						state.Card = typedCard
						assignedCard = true
					}
				}
			}
		}
		if !assignedCard {
			if !hasRaw {
				if responseValueCard, ok := plainFromResponseField(obj, "Card"); ok {
					sourceCard := applyConfiguredKeyedListShapes(responseValueCard, attrValueToPlain(state.Card))
					if valueCard, err := flattenPlainValue(
						sourceCard,
						types.ObjectType{AttrTypes: map[string]attr.Type{"display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType}}}},
						"card",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCard, ok := valueCard.(types.Object); ok {
							state.Card = typedCard
							assignedCard = true
						}
					}
				}
			}
		}
		if !assignedCard && hadRawCard {
			if nullCard, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType}}}}); ok {
				if typedCard, ok := nullCard.(types.Object); ok {
					state.Card = typedCard
				}
			}
		}
	}
	{
		assignedCartesBancaires := false
		hadRawCartesBancaires := false
		if rawValueCartesBancaires, rawOk := plainValueAtPath(raw, "cartes_bancaires"); rawOk {
			hadRawCartesBancaires = true
			if rawValueCartesBancaires != nil {
				sourceCartesBancaires := applyConfiguredKeyedListShapes(rawValueCartesBancaires, attrValueToPlain(state.CartesBancaires))
				if valueCartesBancaires, err := flattenPlainValue(sourceCartesBancaires, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "cartes_bancaires", "raw response"); err != nil {
					return err
				} else {
					if typedCartesBancaires, ok := valueCartesBancaires.(types.Object); ok {
						state.CartesBancaires = typedCartesBancaires
						assignedCartesBancaires = true
					}
				}
			}
		}
		if !assignedCartesBancaires {
			if !hasRaw {
				if responseValueCartesBancaires, ok := plainFromResponseField(obj, "CartesBancaires"); ok {
					sourceCartesBancaires := applyConfiguredKeyedListShapes(responseValueCartesBancaires, attrValueToPlain(state.CartesBancaires))
					if valueCartesBancaires, err := flattenPlainValue(
						sourceCartesBancaires,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"cartes_bancaires",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCartesBancaires, ok := valueCartesBancaires.(types.Object); ok {
							state.CartesBancaires = typedCartesBancaires
							assignedCartesBancaires = true
						}
					}
				}
			}
		}
		if !assignedCartesBancaires && hadRawCartesBancaires {
			if nullCartesBancaires, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedCartesBancaires, ok := nullCartesBancaires.(types.Object); ok {
					state.CartesBancaires = typedCartesBancaires
				}
			}
		}
	}
	{
		assignedCashApp := false
		hadRawCashApp := false
		if rawValueCashApp, rawOk := plainValueAtPath(raw, "cashapp"); rawOk {
			hadRawCashApp = true
			if rawValueCashApp != nil {
				sourceCashApp := applyConfiguredKeyedListShapes(rawValueCashApp, attrValueToPlain(state.CashApp))
				if valueCashApp, err := flattenPlainValue(sourceCashApp, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "cashapp", "raw response"); err != nil {
					return err
				} else {
					if typedCashApp, ok := valueCashApp.(types.Object); ok {
						state.CashApp = typedCashApp
						assignedCashApp = true
					}
				}
			}
		}
		if !assignedCashApp {
			if !hasRaw {
				if responseValueCashApp, ok := plainFromResponseField(obj, "CashApp"); ok {
					sourceCashApp := applyConfiguredKeyedListShapes(responseValueCashApp, attrValueToPlain(state.CashApp))
					if valueCashApp, err := flattenPlainValue(
						sourceCashApp,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"cashapp",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCashApp, ok := valueCashApp.(types.Object); ok {
							state.CashApp = typedCashApp
							assignedCashApp = true
						}
					}
				}
			}
		}
		if !assignedCashApp && hadRawCashApp {
			if nullCashApp, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedCashApp, ok := nullCashApp.(types.Object); ok {
					state.CashApp = typedCashApp
				}
			}
		}
	}
	{
		assignedCrypto := false
		hadRawCrypto := false
		if rawValueCrypto, rawOk := plainValueAtPath(raw, "crypto"); rawOk {
			hadRawCrypto = true
			if rawValueCrypto != nil {
				sourceCrypto := applyConfiguredKeyedListShapes(rawValueCrypto, attrValueToPlain(state.Crypto))
				if valueCrypto, err := flattenPlainValue(sourceCrypto, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "crypto", "raw response"); err != nil {
					return err
				} else {
					if typedCrypto, ok := valueCrypto.(types.Object); ok {
						state.Crypto = typedCrypto
						assignedCrypto = true
					}
				}
			}
		}
		if !assignedCrypto {
			if !hasRaw {
				if responseValueCrypto, ok := plainFromResponseField(obj, "Crypto"); ok {
					sourceCrypto := applyConfiguredKeyedListShapes(responseValueCrypto, attrValueToPlain(state.Crypto))
					if valueCrypto, err := flattenPlainValue(
						sourceCrypto,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"crypto",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCrypto, ok := valueCrypto.(types.Object); ok {
							state.Crypto = typedCrypto
							assignedCrypto = true
						}
					}
				}
			}
		}
		if !assignedCrypto && hadRawCrypto {
			if nullCrypto, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedCrypto, ok := nullCrypto.(types.Object); ok {
					state.Crypto = typedCrypto
				}
			}
		}
	}
	{
		assignedCustomerBalance := false
		hadRawCustomerBalance := false
		if rawValueCustomerBalance, rawOk := plainValueAtPath(raw, "customer_balance"); rawOk {
			hadRawCustomerBalance = true
			if rawValueCustomerBalance != nil {
				sourceCustomerBalance := applyConfiguredKeyedListShapes(rawValueCustomerBalance, attrValueToPlain(state.CustomerBalance))
				if valueCustomerBalance, err := flattenPlainValue(sourceCustomerBalance, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "customer_balance", "raw response"); err != nil {
					return err
				} else {
					if typedCustomerBalance, ok := valueCustomerBalance.(types.Object); ok {
						state.CustomerBalance = typedCustomerBalance
						assignedCustomerBalance = true
					}
				}
			}
		}
		if !assignedCustomerBalance {
			if !hasRaw {
				if responseValueCustomerBalance, ok := plainFromResponseField(obj, "CustomerBalance"); ok {
					sourceCustomerBalance := applyConfiguredKeyedListShapes(responseValueCustomerBalance, attrValueToPlain(state.CustomerBalance))
					if valueCustomerBalance, err := flattenPlainValue(
						sourceCustomerBalance,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"customer_balance",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCustomerBalance, ok := valueCustomerBalance.(types.Object); ok {
							state.CustomerBalance = typedCustomerBalance
							assignedCustomerBalance = true
						}
					}
				}
			}
		}
		if !assignedCustomerBalance && hadRawCustomerBalance {
			if nullCustomerBalance, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedCustomerBalance, ok := nullCustomerBalance.(types.Object); ok {
					state.CustomerBalance = typedCustomerBalance
				}
			}
		}
	}
	{
		assignedEPS := false
		hadRawEPS := false
		if rawValueEPS, rawOk := plainValueAtPath(raw, "eps"); rawOk {
			hadRawEPS = true
			if rawValueEPS != nil {
				sourceEPS := applyConfiguredKeyedListShapes(rawValueEPS, attrValueToPlain(state.EPS))
				if valueEPS, err := flattenPlainValue(sourceEPS, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "eps", "raw response"); err != nil {
					return err
				} else {
					if typedEPS, ok := valueEPS.(types.Object); ok {
						state.EPS = typedEPS
						assignedEPS = true
					}
				}
			}
		}
		if !assignedEPS {
			if !hasRaw {
				if responseValueEPS, ok := plainFromResponseField(obj, "EPS"); ok {
					sourceEPS := applyConfiguredKeyedListShapes(responseValueEPS, attrValueToPlain(state.EPS))
					if valueEPS, err := flattenPlainValue(
						sourceEPS,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"eps",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedEPS, ok := valueEPS.(types.Object); ok {
							state.EPS = typedEPS
							assignedEPS = true
						}
					}
				}
			}
		}
		if !assignedEPS && hadRawEPS {
			if nullEPS, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedEPS, ok := nullEPS.(types.Object); ok {
					state.EPS = typedEPS
				}
			}
		}
	}
	{
		assignedFPX := false
		hadRawFPX := false
		if rawValueFPX, rawOk := plainValueAtPath(raw, "fpx"); rawOk {
			hadRawFPX = true
			if rawValueFPX != nil {
				sourceFPX := applyConfiguredKeyedListShapes(rawValueFPX, attrValueToPlain(state.FPX))
				if valueFPX, err := flattenPlainValue(sourceFPX, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "fpx", "raw response"); err != nil {
					return err
				} else {
					if typedFPX, ok := valueFPX.(types.Object); ok {
						state.FPX = typedFPX
						assignedFPX = true
					}
				}
			}
		}
		if !assignedFPX {
			if !hasRaw {
				if responseValueFPX, ok := plainFromResponseField(obj, "FPX"); ok {
					sourceFPX := applyConfiguredKeyedListShapes(responseValueFPX, attrValueToPlain(state.FPX))
					if valueFPX, err := flattenPlainValue(
						sourceFPX,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"fpx",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFPX, ok := valueFPX.(types.Object); ok {
							state.FPX = typedFPX
							assignedFPX = true
						}
					}
				}
			}
		}
		if !assignedFPX && hadRawFPX {
			if nullFPX, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedFPX, ok := nullFPX.(types.Object); ok {
					state.FPX = typedFPX
				}
			}
		}
	}
	{
		assignedGiropay := false
		hadRawGiropay := false
		if rawValueGiropay, rawOk := plainValueAtPath(raw, "giropay"); rawOk {
			hadRawGiropay = true
			if rawValueGiropay != nil {
				sourceGiropay := applyConfiguredKeyedListShapes(rawValueGiropay, attrValueToPlain(state.Giropay))
				if valueGiropay, err := flattenPlainValue(sourceGiropay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "giropay", "raw response"); err != nil {
					return err
				} else {
					if typedGiropay, ok := valueGiropay.(types.Object); ok {
						state.Giropay = typedGiropay
						assignedGiropay = true
					}
				}
			}
		}
		if !assignedGiropay {
			if !hasRaw {
				if responseValueGiropay, ok := plainFromResponseField(obj, "Giropay"); ok {
					sourceGiropay := applyConfiguredKeyedListShapes(responseValueGiropay, attrValueToPlain(state.Giropay))
					if valueGiropay, err := flattenPlainValue(
						sourceGiropay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"giropay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedGiropay, ok := valueGiropay.(types.Object); ok {
							state.Giropay = typedGiropay
							assignedGiropay = true
						}
					}
				}
			}
		}
		if !assignedGiropay && hadRawGiropay {
			if nullGiropay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedGiropay, ok := nullGiropay.(types.Object); ok {
					state.Giropay = typedGiropay
				}
			}
		}
	}
	{
		assignedGooglePay := false
		hadRawGooglePay := false
		if rawValueGooglePay, rawOk := plainValueAtPath(raw, "google_pay"); rawOk {
			hadRawGooglePay = true
			if rawValueGooglePay != nil {
				sourceGooglePay := applyConfiguredKeyedListShapes(rawValueGooglePay, attrValueToPlain(state.GooglePay))
				if valueGooglePay, err := flattenPlainValue(sourceGooglePay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "google_pay", "raw response"); err != nil {
					return err
				} else {
					if typedGooglePay, ok := valueGooglePay.(types.Object); ok {
						state.GooglePay = typedGooglePay
						assignedGooglePay = true
					}
				}
			}
		}
		if !assignedGooglePay {
			if !hasRaw {
				if responseValueGooglePay, ok := plainFromResponseField(obj, "GooglePay"); ok {
					sourceGooglePay := applyConfiguredKeyedListShapes(responseValueGooglePay, attrValueToPlain(state.GooglePay))
					if valueGooglePay, err := flattenPlainValue(
						sourceGooglePay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"google_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedGooglePay, ok := valueGooglePay.(types.Object); ok {
							state.GooglePay = typedGooglePay
							assignedGooglePay = true
						}
					}
				}
			}
		}
		if !assignedGooglePay && hadRawGooglePay {
			if nullGooglePay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedGooglePay, ok := nullGooglePay.(types.Object); ok {
					state.GooglePay = typedGooglePay
				}
			}
		}
	}
	{
		assignedGrabpay := false
		hadRawGrabpay := false
		if rawValueGrabpay, rawOk := plainValueAtPath(raw, "grabpay"); rawOk {
			hadRawGrabpay = true
			if rawValueGrabpay != nil {
				sourceGrabpay := applyConfiguredKeyedListShapes(rawValueGrabpay, attrValueToPlain(state.Grabpay))
				if valueGrabpay, err := flattenPlainValue(sourceGrabpay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "grabpay", "raw response"); err != nil {
					return err
				} else {
					if typedGrabpay, ok := valueGrabpay.(types.Object); ok {
						state.Grabpay = typedGrabpay
						assignedGrabpay = true
					}
				}
			}
		}
		if !assignedGrabpay {
			if !hasRaw {
				if responseValueGrabpay, ok := plainFromResponseField(obj, "Grabpay"); ok {
					sourceGrabpay := applyConfiguredKeyedListShapes(responseValueGrabpay, attrValueToPlain(state.Grabpay))
					if valueGrabpay, err := flattenPlainValue(
						sourceGrabpay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"grabpay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedGrabpay, ok := valueGrabpay.(types.Object); ok {
							state.Grabpay = typedGrabpay
							assignedGrabpay = true
						}
					}
				}
			}
		}
		if !assignedGrabpay && hadRawGrabpay {
			if nullGrabpay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedGrabpay, ok := nullGrabpay.(types.Object); ok {
					state.Grabpay = typedGrabpay
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
		assignedIDEAL := false
		hadRawIDEAL := false
		if rawValueIDEAL, rawOk := plainValueAtPath(raw, "ideal"); rawOk {
			hadRawIDEAL = true
			if rawValueIDEAL != nil {
				sourceIDEAL := applyConfiguredKeyedListShapes(rawValueIDEAL, attrValueToPlain(state.IDEAL))
				if valueIDEAL, err := flattenPlainValue(sourceIDEAL, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "ideal", "raw response"); err != nil {
					return err
				} else {
					if typedIDEAL, ok := valueIDEAL.(types.Object); ok {
						state.IDEAL = typedIDEAL
						assignedIDEAL = true
					}
				}
			}
		}
		if !assignedIDEAL {
			if !hasRaw {
				if responseValueIDEAL, ok := plainFromResponseField(obj, "IDEAL"); ok {
					sourceIDEAL := applyConfiguredKeyedListShapes(responseValueIDEAL, attrValueToPlain(state.IDEAL))
					if valueIDEAL, err := flattenPlainValue(
						sourceIDEAL,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"ideal",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedIDEAL, ok := valueIDEAL.(types.Object); ok {
							state.IDEAL = typedIDEAL
							assignedIDEAL = true
						}
					}
				}
			}
		}
		if !assignedIDEAL && hadRawIDEAL {
			if nullIDEAL, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedIDEAL, ok := nullIDEAL.(types.Object); ok {
					state.IDEAL = typedIDEAL
				}
			}
		}
	}
	{
		if rawValueIsDefault, rawOk := plainValueAtPath(raw, "is_default"); rawOk {
			if valueIsDefault, err := flattenPlainValue(rawValueIsDefault, types.BoolType, "is_default", "raw response"); err != nil {
				return err
			} else {
				if typedIsDefault, ok := valueIsDefault.(types.Bool); ok {
					state.IsDefault = typedIsDefault
				}
			}
		} else if !hasRaw {
			if responseValueIsDefault, ok := plainFromResponseField(obj, "IsDefault"); ok {
				if valueIsDefault, err := flattenPlainValue(responseValueIsDefault, types.BoolType, "is_default", "response struct"); err != nil {
					return err
				} else {
					if typedIsDefault, ok := valueIsDefault.(types.Bool); ok {
						state.IsDefault = typedIsDefault
					}
				}
			}
		}
	}
	{
		assignedJCB := false
		hadRawJCB := false
		if rawValueJCB, rawOk := plainValueAtPath(raw, "jcb"); rawOk {
			hadRawJCB = true
			if rawValueJCB != nil {
				sourceJCB := applyConfiguredKeyedListShapes(rawValueJCB, attrValueToPlain(state.JCB))
				if valueJCB, err := flattenPlainValue(sourceJCB, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "jcb", "raw response"); err != nil {
					return err
				} else {
					if typedJCB, ok := valueJCB.(types.Object); ok {
						state.JCB = typedJCB
						assignedJCB = true
					}
				}
			}
		}
		if !assignedJCB {
			if !hasRaw {
				if responseValueJCB, ok := plainFromResponseField(obj, "JCB"); ok {
					sourceJCB := applyConfiguredKeyedListShapes(responseValueJCB, attrValueToPlain(state.JCB))
					if valueJCB, err := flattenPlainValue(
						sourceJCB,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"jcb",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedJCB, ok := valueJCB.(types.Object); ok {
							state.JCB = typedJCB
							assignedJCB = true
						}
					}
				}
			}
		}
		if !assignedJCB && hadRawJCB {
			if nullJCB, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedJCB, ok := nullJCB.(types.Object); ok {
					state.JCB = typedJCB
				}
			}
		}
	}
	{
		assignedKakaoPay := false
		hadRawKakaoPay := false
		if rawValueKakaoPay, rawOk := plainValueAtPath(raw, "kakao_pay"); rawOk {
			hadRawKakaoPay = true
			if rawValueKakaoPay != nil {
				sourceKakaoPay := applyConfiguredKeyedListShapes(rawValueKakaoPay, attrValueToPlain(state.KakaoPay))
				if valueKakaoPay, err := flattenPlainValue(sourceKakaoPay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "kakao_pay", "raw response"); err != nil {
					return err
				} else {
					if typedKakaoPay, ok := valueKakaoPay.(types.Object); ok {
						state.KakaoPay = typedKakaoPay
						assignedKakaoPay = true
					}
				}
			}
		}
		if !assignedKakaoPay {
			if !hasRaw {
				if responseValueKakaoPay, ok := plainFromResponseField(obj, "KakaoPay"); ok {
					sourceKakaoPay := applyConfiguredKeyedListShapes(responseValueKakaoPay, attrValueToPlain(state.KakaoPay))
					if valueKakaoPay, err := flattenPlainValue(
						sourceKakaoPay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"kakao_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedKakaoPay, ok := valueKakaoPay.(types.Object); ok {
							state.KakaoPay = typedKakaoPay
							assignedKakaoPay = true
						}
					}
				}
			}
		}
		if !assignedKakaoPay && hadRawKakaoPay {
			if nullKakaoPay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedKakaoPay, ok := nullKakaoPay.(types.Object); ok {
					state.KakaoPay = typedKakaoPay
				}
			}
		}
	}
	{
		assignedKlarna := false
		hadRawKlarna := false
		if rawValueKlarna, rawOk := plainValueAtPath(raw, "klarna"); rawOk {
			hadRawKlarna = true
			if rawValueKlarna != nil {
				sourceKlarna := applyConfiguredKeyedListShapes(rawValueKlarna, attrValueToPlain(state.Klarna))
				if valueKlarna, err := flattenPlainValue(sourceKlarna, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "klarna", "raw response"); err != nil {
					return err
				} else {
					if typedKlarna, ok := valueKlarna.(types.Object); ok {
						state.Klarna = typedKlarna
						assignedKlarna = true
					}
				}
			}
		}
		if !assignedKlarna {
			if !hasRaw {
				if responseValueKlarna, ok := plainFromResponseField(obj, "Klarna"); ok {
					sourceKlarna := applyConfiguredKeyedListShapes(responseValueKlarna, attrValueToPlain(state.Klarna))
					if valueKlarna, err := flattenPlainValue(
						sourceKlarna,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"klarna",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedKlarna, ok := valueKlarna.(types.Object); ok {
							state.Klarna = typedKlarna
							assignedKlarna = true
						}
					}
				}
			}
		}
		if !assignedKlarna && hadRawKlarna {
			if nullKlarna, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedKlarna, ok := nullKlarna.(types.Object); ok {
					state.Klarna = typedKlarna
				}
			}
		}
	}
	{
		assignedKonbini := false
		hadRawKonbini := false
		if rawValueKonbini, rawOk := plainValueAtPath(raw, "konbini"); rawOk {
			hadRawKonbini = true
			if rawValueKonbini != nil {
				sourceKonbini := applyConfiguredKeyedListShapes(rawValueKonbini, attrValueToPlain(state.Konbini))
				if valueKonbini, err := flattenPlainValue(sourceKonbini, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "konbini", "raw response"); err != nil {
					return err
				} else {
					if typedKonbini, ok := valueKonbini.(types.Object); ok {
						state.Konbini = typedKonbini
						assignedKonbini = true
					}
				}
			}
		}
		if !assignedKonbini {
			if !hasRaw {
				if responseValueKonbini, ok := plainFromResponseField(obj, "Konbini"); ok {
					sourceKonbini := applyConfiguredKeyedListShapes(responseValueKonbini, attrValueToPlain(state.Konbini))
					if valueKonbini, err := flattenPlainValue(
						sourceKonbini,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"konbini",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedKonbini, ok := valueKonbini.(types.Object); ok {
							state.Konbini = typedKonbini
							assignedKonbini = true
						}
					}
				}
			}
		}
		if !assignedKonbini && hadRawKonbini {
			if nullKonbini, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedKonbini, ok := nullKonbini.(types.Object); ok {
					state.Konbini = typedKonbini
				}
			}
		}
	}
	{
		assignedKrCard := false
		hadRawKrCard := false
		if rawValueKrCard, rawOk := plainValueAtPath(raw, "kr_card"); rawOk {
			hadRawKrCard = true
			if rawValueKrCard != nil {
				sourceKrCard := applyConfiguredKeyedListShapes(rawValueKrCard, attrValueToPlain(state.KrCard))
				if valueKrCard, err := flattenPlainValue(sourceKrCard, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "kr_card", "raw response"); err != nil {
					return err
				} else {
					if typedKrCard, ok := valueKrCard.(types.Object); ok {
						state.KrCard = typedKrCard
						assignedKrCard = true
					}
				}
			}
		}
		if !assignedKrCard {
			if !hasRaw {
				if responseValueKrCard, ok := plainFromResponseField(obj, "KrCard"); ok {
					sourceKrCard := applyConfiguredKeyedListShapes(responseValueKrCard, attrValueToPlain(state.KrCard))
					if valueKrCard, err := flattenPlainValue(
						sourceKrCard,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"kr_card",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedKrCard, ok := valueKrCard.(types.Object); ok {
							state.KrCard = typedKrCard
							assignedKrCard = true
						}
					}
				}
			}
		}
		if !assignedKrCard && hadRawKrCard {
			if nullKrCard, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedKrCard, ok := nullKrCard.(types.Object); ok {
					state.KrCard = typedKrCard
				}
			}
		}
	}
	{
		assignedLink := false
		hadRawLink := false
		if rawValueLink, rawOk := plainValueAtPath(raw, "link"); rawOk {
			hadRawLink = true
			if rawValueLink != nil {
				sourceLink := applyConfiguredKeyedListShapes(rawValueLink, attrValueToPlain(state.Link))
				if valueLink, err := flattenPlainValue(sourceLink, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "link", "raw response"); err != nil {
					return err
				} else {
					if typedLink, ok := valueLink.(types.Object); ok {
						state.Link = typedLink
						assignedLink = true
					}
				}
			}
		}
		if !assignedLink {
			if !hasRaw {
				if responseValueLink, ok := plainFromResponseField(obj, "Link"); ok {
					sourceLink := applyConfiguredKeyedListShapes(responseValueLink, attrValueToPlain(state.Link))
					if valueLink, err := flattenPlainValue(
						sourceLink,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"link",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLink, ok := valueLink.(types.Object); ok {
							state.Link = typedLink
							assignedLink = true
						}
					}
				}
			}
		}
		if !assignedLink && hadRawLink {
			if nullLink, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedLink, ok := nullLink.(types.Object); ok {
					state.Link = typedLink
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
		assignedMbWay := false
		hadRawMbWay := false
		if rawValueMbWay, rawOk := plainValueAtPath(raw, "mb_way"); rawOk {
			hadRawMbWay = true
			if rawValueMbWay != nil {
				sourceMbWay := applyConfiguredKeyedListShapes(rawValueMbWay, attrValueToPlain(state.MbWay))
				if valueMbWay, err := flattenPlainValue(sourceMbWay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "mb_way", "raw response"); err != nil {
					return err
				} else {
					if typedMbWay, ok := valueMbWay.(types.Object); ok {
						state.MbWay = typedMbWay
						assignedMbWay = true
					}
				}
			}
		}
		if !assignedMbWay {
			if !hasRaw {
				if responseValueMbWay, ok := plainFromResponseField(obj, "MbWay"); ok {
					sourceMbWay := applyConfiguredKeyedListShapes(responseValueMbWay, attrValueToPlain(state.MbWay))
					if valueMbWay, err := flattenPlainValue(
						sourceMbWay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"mb_way",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedMbWay, ok := valueMbWay.(types.Object); ok {
							state.MbWay = typedMbWay
							assignedMbWay = true
						}
					}
				}
			}
		}
		if !assignedMbWay && hadRawMbWay {
			if nullMbWay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedMbWay, ok := nullMbWay.(types.Object); ok {
					state.MbWay = typedMbWay
				}
			}
		}
	}
	{
		assignedMobilepay := false
		hadRawMobilepay := false
		if rawValueMobilepay, rawOk := plainValueAtPath(raw, "mobilepay"); rawOk {
			hadRawMobilepay = true
			if rawValueMobilepay != nil {
				sourceMobilepay := applyConfiguredKeyedListShapes(rawValueMobilepay, attrValueToPlain(state.Mobilepay))
				if valueMobilepay, err := flattenPlainValue(sourceMobilepay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "mobilepay", "raw response"); err != nil {
					return err
				} else {
					if typedMobilepay, ok := valueMobilepay.(types.Object); ok {
						state.Mobilepay = typedMobilepay
						assignedMobilepay = true
					}
				}
			}
		}
		if !assignedMobilepay {
			if !hasRaw {
				if responseValueMobilepay, ok := plainFromResponseField(obj, "Mobilepay"); ok {
					sourceMobilepay := applyConfiguredKeyedListShapes(responseValueMobilepay, attrValueToPlain(state.Mobilepay))
					if valueMobilepay, err := flattenPlainValue(
						sourceMobilepay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"mobilepay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedMobilepay, ok := valueMobilepay.(types.Object); ok {
							state.Mobilepay = typedMobilepay
							assignedMobilepay = true
						}
					}
				}
			}
		}
		if !assignedMobilepay && hadRawMobilepay {
			if nullMobilepay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedMobilepay, ok := nullMobilepay.(types.Object); ok {
					state.Mobilepay = typedMobilepay
				}
			}
		}
	}
	{
		assignedMultibanco := false
		hadRawMultibanco := false
		if rawValueMultibanco, rawOk := plainValueAtPath(raw, "multibanco"); rawOk {
			hadRawMultibanco = true
			if rawValueMultibanco != nil {
				sourceMultibanco := applyConfiguredKeyedListShapes(rawValueMultibanco, attrValueToPlain(state.Multibanco))
				if valueMultibanco, err := flattenPlainValue(sourceMultibanco, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "multibanco", "raw response"); err != nil {
					return err
				} else {
					if typedMultibanco, ok := valueMultibanco.(types.Object); ok {
						state.Multibanco = typedMultibanco
						assignedMultibanco = true
					}
				}
			}
		}
		if !assignedMultibanco {
			if !hasRaw {
				if responseValueMultibanco, ok := plainFromResponseField(obj, "Multibanco"); ok {
					sourceMultibanco := applyConfiguredKeyedListShapes(responseValueMultibanco, attrValueToPlain(state.Multibanco))
					if valueMultibanco, err := flattenPlainValue(
						sourceMultibanco,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"multibanco",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedMultibanco, ok := valueMultibanco.(types.Object); ok {
							state.Multibanco = typedMultibanco
							assignedMultibanco = true
						}
					}
				}
			}
		}
		if !assignedMultibanco && hadRawMultibanco {
			if nullMultibanco, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedMultibanco, ok := nullMultibanco.(types.Object); ok {
					state.Multibanco = typedMultibanco
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
		assignedNaverPay := false
		hadRawNaverPay := false
		if rawValueNaverPay, rawOk := plainValueAtPath(raw, "naver_pay"); rawOk {
			hadRawNaverPay = true
			if rawValueNaverPay != nil {
				sourceNaverPay := applyConfiguredKeyedListShapes(rawValueNaverPay, attrValueToPlain(state.NaverPay))
				if valueNaverPay, err := flattenPlainValue(sourceNaverPay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "naver_pay", "raw response"); err != nil {
					return err
				} else {
					if typedNaverPay, ok := valueNaverPay.(types.Object); ok {
						state.NaverPay = typedNaverPay
						assignedNaverPay = true
					}
				}
			}
		}
		if !assignedNaverPay {
			if !hasRaw {
				if responseValueNaverPay, ok := plainFromResponseField(obj, "NaverPay"); ok {
					sourceNaverPay := applyConfiguredKeyedListShapes(responseValueNaverPay, attrValueToPlain(state.NaverPay))
					if valueNaverPay, err := flattenPlainValue(
						sourceNaverPay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"naver_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedNaverPay, ok := valueNaverPay.(types.Object); ok {
							state.NaverPay = typedNaverPay
							assignedNaverPay = true
						}
					}
				}
			}
		}
		if !assignedNaverPay && hadRawNaverPay {
			if nullNaverPay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedNaverPay, ok := nullNaverPay.(types.Object); ok {
					state.NaverPay = typedNaverPay
				}
			}
		}
	}
	{
		assignedNzBankAccount := false
		hadRawNzBankAccount := false
		if rawValueNzBankAccount, rawOk := plainValueAtPath(raw, "nz_bank_account"); rawOk {
			hadRawNzBankAccount = true
			if rawValueNzBankAccount != nil {
				sourceNzBankAccount := applyConfiguredKeyedListShapes(rawValueNzBankAccount, attrValueToPlain(state.NzBankAccount))
				if valueNzBankAccount, err := flattenPlainValue(sourceNzBankAccount, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "nz_bank_account", "raw response"); err != nil {
					return err
				} else {
					if typedNzBankAccount, ok := valueNzBankAccount.(types.Object); ok {
						state.NzBankAccount = typedNzBankAccount
						assignedNzBankAccount = true
					}
				}
			}
		}
		if !assignedNzBankAccount {
			if !hasRaw {
				if responseValueNzBankAccount, ok := plainFromResponseField(obj, "NzBankAccount"); ok {
					sourceNzBankAccount := applyConfiguredKeyedListShapes(responseValueNzBankAccount, attrValueToPlain(state.NzBankAccount))
					if valueNzBankAccount, err := flattenPlainValue(
						sourceNzBankAccount,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"nz_bank_account",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedNzBankAccount, ok := valueNzBankAccount.(types.Object); ok {
							state.NzBankAccount = typedNzBankAccount
							assignedNzBankAccount = true
						}
					}
				}
			}
		}
		if !assignedNzBankAccount && hadRawNzBankAccount {
			if nullNzBankAccount, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedNzBankAccount, ok := nullNzBankAccount.(types.Object); ok {
					state.NzBankAccount = typedNzBankAccount
				}
			}
		}
	}
	{
		assignedOXXO := false
		hadRawOXXO := false
		if rawValueOXXO, rawOk := plainValueAtPath(raw, "oxxo"); rawOk {
			hadRawOXXO = true
			if rawValueOXXO != nil {
				sourceOXXO := applyConfiguredKeyedListShapes(rawValueOXXO, attrValueToPlain(state.OXXO))
				if valueOXXO, err := flattenPlainValue(sourceOXXO, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "oxxo", "raw response"); err != nil {
					return err
				} else {
					if typedOXXO, ok := valueOXXO.(types.Object); ok {
						state.OXXO = typedOXXO
						assignedOXXO = true
					}
				}
			}
		}
		if !assignedOXXO {
			if !hasRaw {
				if responseValueOXXO, ok := plainFromResponseField(obj, "OXXO"); ok {
					sourceOXXO := applyConfiguredKeyedListShapes(responseValueOXXO, attrValueToPlain(state.OXXO))
					if valueOXXO, err := flattenPlainValue(
						sourceOXXO,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"oxxo",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedOXXO, ok := valueOXXO.(types.Object); ok {
							state.OXXO = typedOXXO
							assignedOXXO = true
						}
					}
				}
			}
		}
		if !assignedOXXO && hadRawOXXO {
			if nullOXXO, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedOXXO, ok := nullOXXO.(types.Object); ok {
					state.OXXO = typedOXXO
				}
			}
		}
	}
	{
		assignedP24 := false
		hadRawP24 := false
		if rawValueP24, rawOk := plainValueAtPath(raw, "p24"); rawOk {
			hadRawP24 = true
			if rawValueP24 != nil {
				sourceP24 := applyConfiguredKeyedListShapes(rawValueP24, attrValueToPlain(state.P24))
				if valueP24, err := flattenPlainValue(sourceP24, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "p24", "raw response"); err != nil {
					return err
				} else {
					if typedP24, ok := valueP24.(types.Object); ok {
						state.P24 = typedP24
						assignedP24 = true
					}
				}
			}
		}
		if !assignedP24 {
			if !hasRaw {
				if responseValueP24, ok := plainFromResponseField(obj, "P24"); ok {
					sourceP24 := applyConfiguredKeyedListShapes(responseValueP24, attrValueToPlain(state.P24))
					if valueP24, err := flattenPlainValue(
						sourceP24,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"p24",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedP24, ok := valueP24.(types.Object); ok {
							state.P24 = typedP24
							assignedP24 = true
						}
					}
				}
			}
		}
		if !assignedP24 && hadRawP24 {
			if nullP24, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedP24, ok := nullP24.(types.Object); ok {
					state.P24 = typedP24
				}
			}
		}
	}
	{
		if rawValueParent, rawOk := plainValueAtPath(raw, "parent"); rawOk {
			if valueParent, err := flattenPlainValue(rawValueParent, types.StringType, "parent", "raw response"); err != nil {
				return err
			} else {
				if typedParent, ok := valueParent.(types.String); ok {
					state.Parent = typedParent
				}
			}
		} else if !hasRaw {
			if responseValueParent, ok := plainFromResponseField(obj, "Parent"); ok {
				if valueParent, err := flattenPlainValue(responseValueParent, types.StringType, "parent", "response struct"); err != nil {
					return err
				} else {
					if typedParent, ok := valueParent.(types.String); ok {
						state.Parent = typedParent
					}
				}
			}
		}
	}
	{
		assignedPayByBank := false
		hadRawPayByBank := false
		if rawValuePayByBank, rawOk := plainValueAtPath(raw, "pay_by_bank"); rawOk {
			hadRawPayByBank = true
			if rawValuePayByBank != nil {
				sourcePayByBank := applyConfiguredKeyedListShapes(rawValuePayByBank, attrValueToPlain(state.PayByBank))
				if valuePayByBank, err := flattenPlainValue(sourcePayByBank, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "pay_by_bank", "raw response"); err != nil {
					return err
				} else {
					if typedPayByBank, ok := valuePayByBank.(types.Object); ok {
						state.PayByBank = typedPayByBank
						assignedPayByBank = true
					}
				}
			}
		}
		if !assignedPayByBank {
			if !hasRaw {
				if responseValuePayByBank, ok := plainFromResponseField(obj, "PayByBank"); ok {
					sourcePayByBank := applyConfiguredKeyedListShapes(responseValuePayByBank, attrValueToPlain(state.PayByBank))
					if valuePayByBank, err := flattenPlainValue(
						sourcePayByBank,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"pay_by_bank",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPayByBank, ok := valuePayByBank.(types.Object); ok {
							state.PayByBank = typedPayByBank
							assignedPayByBank = true
						}
					}
				}
			}
		}
		if !assignedPayByBank && hadRawPayByBank {
			if nullPayByBank, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedPayByBank, ok := nullPayByBank.(types.Object); ok {
					state.PayByBank = typedPayByBank
				}
			}
		}
	}
	{
		assignedPayco := false
		hadRawPayco := false
		if rawValuePayco, rawOk := plainValueAtPath(raw, "payco"); rawOk {
			hadRawPayco = true
			if rawValuePayco != nil {
				sourcePayco := applyConfiguredKeyedListShapes(rawValuePayco, attrValueToPlain(state.Payco))
				if valuePayco, err := flattenPlainValue(sourcePayco, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "payco", "raw response"); err != nil {
					return err
				} else {
					if typedPayco, ok := valuePayco.(types.Object); ok {
						state.Payco = typedPayco
						assignedPayco = true
					}
				}
			}
		}
		if !assignedPayco {
			if !hasRaw {
				if responseValuePayco, ok := plainFromResponseField(obj, "Payco"); ok {
					sourcePayco := applyConfiguredKeyedListShapes(responseValuePayco, attrValueToPlain(state.Payco))
					if valuePayco, err := flattenPlainValue(
						sourcePayco,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"payco",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPayco, ok := valuePayco.(types.Object); ok {
							state.Payco = typedPayco
							assignedPayco = true
						}
					}
				}
			}
		}
		if !assignedPayco && hadRawPayco {
			if nullPayco, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedPayco, ok := nullPayco.(types.Object); ok {
					state.Payco = typedPayco
				}
			}
		}
	}
	{
		assignedPayNow := false
		hadRawPayNow := false
		if rawValuePayNow, rawOk := plainValueAtPath(raw, "paynow"); rawOk {
			hadRawPayNow = true
			if rawValuePayNow != nil {
				sourcePayNow := applyConfiguredKeyedListShapes(rawValuePayNow, attrValueToPlain(state.PayNow))
				if valuePayNow, err := flattenPlainValue(sourcePayNow, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "paynow", "raw response"); err != nil {
					return err
				} else {
					if typedPayNow, ok := valuePayNow.(types.Object); ok {
						state.PayNow = typedPayNow
						assignedPayNow = true
					}
				}
			}
		}
		if !assignedPayNow {
			if !hasRaw {
				if responseValuePayNow, ok := plainFromResponseField(obj, "PayNow"); ok {
					sourcePayNow := applyConfiguredKeyedListShapes(responseValuePayNow, attrValueToPlain(state.PayNow))
					if valuePayNow, err := flattenPlainValue(
						sourcePayNow,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"paynow",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPayNow, ok := valuePayNow.(types.Object); ok {
							state.PayNow = typedPayNow
							assignedPayNow = true
						}
					}
				}
			}
		}
		if !assignedPayNow && hadRawPayNow {
			if nullPayNow, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedPayNow, ok := nullPayNow.(types.Object); ok {
					state.PayNow = typedPayNow
				}
			}
		}
	}
	{
		assignedPaypal := false
		hadRawPaypal := false
		if rawValuePaypal, rawOk := plainValueAtPath(raw, "paypal"); rawOk {
			hadRawPaypal = true
			if rawValuePaypal != nil {
				sourcePaypal := applyConfiguredKeyedListShapes(rawValuePaypal, attrValueToPlain(state.Paypal))
				if valuePaypal, err := flattenPlainValue(sourcePaypal, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "paypal", "raw response"); err != nil {
					return err
				} else {
					if typedPaypal, ok := valuePaypal.(types.Object); ok {
						state.Paypal = typedPaypal
						assignedPaypal = true
					}
				}
			}
		}
		if !assignedPaypal {
			if !hasRaw {
				if responseValuePaypal, ok := plainFromResponseField(obj, "Paypal"); ok {
					sourcePaypal := applyConfiguredKeyedListShapes(responseValuePaypal, attrValueToPlain(state.Paypal))
					if valuePaypal, err := flattenPlainValue(
						sourcePaypal,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"paypal",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPaypal, ok := valuePaypal.(types.Object); ok {
							state.Paypal = typedPaypal
							assignedPaypal = true
						}
					}
				}
			}
		}
		if !assignedPaypal && hadRawPaypal {
			if nullPaypal, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedPaypal, ok := nullPaypal.(types.Object); ok {
					state.Paypal = typedPaypal
				}
			}
		}
	}
	{
		assignedPayto := false
		hadRawPayto := false
		if rawValuePayto, rawOk := plainValueAtPath(raw, "payto"); rawOk {
			hadRawPayto = true
			if rawValuePayto != nil {
				sourcePayto := applyConfiguredKeyedListShapes(rawValuePayto, attrValueToPlain(state.Payto))
				if valuePayto, err := flattenPlainValue(sourcePayto, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "payto", "raw response"); err != nil {
					return err
				} else {
					if typedPayto, ok := valuePayto.(types.Object); ok {
						state.Payto = typedPayto
						assignedPayto = true
					}
				}
			}
		}
		if !assignedPayto {
			if !hasRaw {
				if responseValuePayto, ok := plainFromResponseField(obj, "Payto"); ok {
					sourcePayto := applyConfiguredKeyedListShapes(responseValuePayto, attrValueToPlain(state.Payto))
					if valuePayto, err := flattenPlainValue(
						sourcePayto,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"payto",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPayto, ok := valuePayto.(types.Object); ok {
							state.Payto = typedPayto
							assignedPayto = true
						}
					}
				}
			}
		}
		if !assignedPayto && hadRawPayto {
			if nullPayto, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedPayto, ok := nullPayto.(types.Object); ok {
					state.Payto = typedPayto
				}
			}
		}
	}
	{
		assignedPix := false
		hadRawPix := false
		if rawValuePix, rawOk := plainValueAtPath(raw, "pix"); rawOk {
			hadRawPix = true
			if rawValuePix != nil {
				sourcePix := applyConfiguredKeyedListShapes(rawValuePix, attrValueToPlain(state.Pix))
				if valuePix, err := flattenPlainValue(sourcePix, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "pix", "raw response"); err != nil {
					return err
				} else {
					if typedPix, ok := valuePix.(types.Object); ok {
						state.Pix = typedPix
						assignedPix = true
					}
				}
			}
		}
		if !assignedPix {
			if !hasRaw {
				if responseValuePix, ok := plainFromResponseField(obj, "Pix"); ok {
					sourcePix := applyConfiguredKeyedListShapes(responseValuePix, attrValueToPlain(state.Pix))
					if valuePix, err := flattenPlainValue(
						sourcePix,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"pix",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPix, ok := valuePix.(types.Object); ok {
							state.Pix = typedPix
							assignedPix = true
						}
					}
				}
			}
		}
		if !assignedPix && hadRawPix {
			if nullPix, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedPix, ok := nullPix.(types.Object); ok {
					state.Pix = typedPix
				}
			}
		}
	}
	{
		assignedPromptPay := false
		hadRawPromptPay := false
		if rawValuePromptPay, rawOk := plainValueAtPath(raw, "promptpay"); rawOk {
			hadRawPromptPay = true
			if rawValuePromptPay != nil {
				sourcePromptPay := applyConfiguredKeyedListShapes(rawValuePromptPay, attrValueToPlain(state.PromptPay))
				if valuePromptPay, err := flattenPlainValue(sourcePromptPay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "promptpay", "raw response"); err != nil {
					return err
				} else {
					if typedPromptPay, ok := valuePromptPay.(types.Object); ok {
						state.PromptPay = typedPromptPay
						assignedPromptPay = true
					}
				}
			}
		}
		if !assignedPromptPay {
			if !hasRaw {
				if responseValuePromptPay, ok := plainFromResponseField(obj, "PromptPay"); ok {
					sourcePromptPay := applyConfiguredKeyedListShapes(responseValuePromptPay, attrValueToPlain(state.PromptPay))
					if valuePromptPay, err := flattenPlainValue(
						sourcePromptPay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"promptpay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPromptPay, ok := valuePromptPay.(types.Object); ok {
							state.PromptPay = typedPromptPay
							assignedPromptPay = true
						}
					}
				}
			}
		}
		if !assignedPromptPay && hadRawPromptPay {
			if nullPromptPay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedPromptPay, ok := nullPromptPay.(types.Object); ok {
					state.PromptPay = typedPromptPay
				}
			}
		}
	}
	{
		assignedRevolutPay := false
		hadRawRevolutPay := false
		if rawValueRevolutPay, rawOk := plainValueAtPath(raw, "revolut_pay"); rawOk {
			hadRawRevolutPay = true
			if rawValueRevolutPay != nil {
				sourceRevolutPay := applyConfiguredKeyedListShapes(rawValueRevolutPay, attrValueToPlain(state.RevolutPay))
				if valueRevolutPay, err := flattenPlainValue(sourceRevolutPay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "revolut_pay", "raw response"); err != nil {
					return err
				} else {
					if typedRevolutPay, ok := valueRevolutPay.(types.Object); ok {
						state.RevolutPay = typedRevolutPay
						assignedRevolutPay = true
					}
				}
			}
		}
		if !assignedRevolutPay {
			if !hasRaw {
				if responseValueRevolutPay, ok := plainFromResponseField(obj, "RevolutPay"); ok {
					sourceRevolutPay := applyConfiguredKeyedListShapes(responseValueRevolutPay, attrValueToPlain(state.RevolutPay))
					if valueRevolutPay, err := flattenPlainValue(
						sourceRevolutPay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"revolut_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRevolutPay, ok := valueRevolutPay.(types.Object); ok {
							state.RevolutPay = typedRevolutPay
							assignedRevolutPay = true
						}
					}
				}
			}
		}
		if !assignedRevolutPay && hadRawRevolutPay {
			if nullRevolutPay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedRevolutPay, ok := nullRevolutPay.(types.Object); ok {
					state.RevolutPay = typedRevolutPay
				}
			}
		}
	}
	{
		assignedSamsungPay := false
		hadRawSamsungPay := false
		if rawValueSamsungPay, rawOk := plainValueAtPath(raw, "samsung_pay"); rawOk {
			hadRawSamsungPay = true
			if rawValueSamsungPay != nil {
				sourceSamsungPay := applyConfiguredKeyedListShapes(rawValueSamsungPay, attrValueToPlain(state.SamsungPay))
				if valueSamsungPay, err := flattenPlainValue(sourceSamsungPay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "samsung_pay", "raw response"); err != nil {
					return err
				} else {
					if typedSamsungPay, ok := valueSamsungPay.(types.Object); ok {
						state.SamsungPay = typedSamsungPay
						assignedSamsungPay = true
					}
				}
			}
		}
		if !assignedSamsungPay {
			if !hasRaw {
				if responseValueSamsungPay, ok := plainFromResponseField(obj, "SamsungPay"); ok {
					sourceSamsungPay := applyConfiguredKeyedListShapes(responseValueSamsungPay, attrValueToPlain(state.SamsungPay))
					if valueSamsungPay, err := flattenPlainValue(
						sourceSamsungPay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"samsung_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSamsungPay, ok := valueSamsungPay.(types.Object); ok {
							state.SamsungPay = typedSamsungPay
							assignedSamsungPay = true
						}
					}
				}
			}
		}
		if !assignedSamsungPay && hadRawSamsungPay {
			if nullSamsungPay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedSamsungPay, ok := nullSamsungPay.(types.Object); ok {
					state.SamsungPay = typedSamsungPay
				}
			}
		}
	}
	{
		assignedSatispay := false
		hadRawSatispay := false
		if rawValueSatispay, rawOk := plainValueAtPath(raw, "satispay"); rawOk {
			hadRawSatispay = true
			if rawValueSatispay != nil {
				sourceSatispay := applyConfiguredKeyedListShapes(rawValueSatispay, attrValueToPlain(state.Satispay))
				if valueSatispay, err := flattenPlainValue(sourceSatispay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "satispay", "raw response"); err != nil {
					return err
				} else {
					if typedSatispay, ok := valueSatispay.(types.Object); ok {
						state.Satispay = typedSatispay
						assignedSatispay = true
					}
				}
			}
		}
		if !assignedSatispay {
			if !hasRaw {
				if responseValueSatispay, ok := plainFromResponseField(obj, "Satispay"); ok {
					sourceSatispay := applyConfiguredKeyedListShapes(responseValueSatispay, attrValueToPlain(state.Satispay))
					if valueSatispay, err := flattenPlainValue(
						sourceSatispay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"satispay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSatispay, ok := valueSatispay.(types.Object); ok {
							state.Satispay = typedSatispay
							assignedSatispay = true
						}
					}
				}
			}
		}
		if !assignedSatispay && hadRawSatispay {
			if nullSatispay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedSatispay, ok := nullSatispay.(types.Object); ok {
					state.Satispay = typedSatispay
				}
			}
		}
	}
	{
		assignedScalapay := false
		hadRawScalapay := false
		if rawValueScalapay, rawOk := plainValueAtPath(raw, "scalapay"); rawOk {
			hadRawScalapay = true
			if rawValueScalapay != nil {
				sourceScalapay := applyConfiguredKeyedListShapes(rawValueScalapay, attrValueToPlain(state.Scalapay))
				if valueScalapay, err := flattenPlainValue(sourceScalapay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "scalapay", "raw response"); err != nil {
					return err
				} else {
					if typedScalapay, ok := valueScalapay.(types.Object); ok {
						state.Scalapay = typedScalapay
						assignedScalapay = true
					}
				}
			}
		}
		if !assignedScalapay {
			if !hasRaw {
				if responseValueScalapay, ok := plainFromResponseField(obj, "Scalapay"); ok {
					sourceScalapay := applyConfiguredKeyedListShapes(responseValueScalapay, attrValueToPlain(state.Scalapay))
					if valueScalapay, err := flattenPlainValue(
						sourceScalapay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"scalapay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedScalapay, ok := valueScalapay.(types.Object); ok {
							state.Scalapay = typedScalapay
							assignedScalapay = true
						}
					}
				}
			}
		}
		if !assignedScalapay && hadRawScalapay {
			if nullScalapay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedScalapay, ok := nullScalapay.(types.Object); ok {
					state.Scalapay = typedScalapay
				}
			}
		}
	}
	{
		assignedSEPADebit := false
		hadRawSEPADebit := false
		if rawValueSEPADebit, rawOk := plainValueAtPath(raw, "sepa_debit"); rawOk {
			hadRawSEPADebit = true
			if rawValueSEPADebit != nil {
				sourceSEPADebit := applyConfiguredKeyedListShapes(rawValueSEPADebit, attrValueToPlain(state.SEPADebit))
				if valueSEPADebit, err := flattenPlainValue(sourceSEPADebit, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "sepa_debit", "raw response"); err != nil {
					return err
				} else {
					if typedSEPADebit, ok := valueSEPADebit.(types.Object); ok {
						state.SEPADebit = typedSEPADebit
						assignedSEPADebit = true
					}
				}
			}
		}
		if !assignedSEPADebit {
			if !hasRaw {
				if responseValueSEPADebit, ok := plainFromResponseField(obj, "SEPADebit"); ok {
					sourceSEPADebit := applyConfiguredKeyedListShapes(responseValueSEPADebit, attrValueToPlain(state.SEPADebit))
					if valueSEPADebit, err := flattenPlainValue(
						sourceSEPADebit,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"sepa_debit",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSEPADebit, ok := valueSEPADebit.(types.Object); ok {
							state.SEPADebit = typedSEPADebit
							assignedSEPADebit = true
						}
					}
				}
			}
		}
		if !assignedSEPADebit && hadRawSEPADebit {
			if nullSEPADebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedSEPADebit, ok := nullSEPADebit.(types.Object); ok {
					state.SEPADebit = typedSEPADebit
				}
			}
		}
	}
	{
		assignedSofort := false
		hadRawSofort := false
		if rawValueSofort, rawOk := plainValueAtPath(raw, "sofort"); rawOk {
			hadRawSofort = true
			if rawValueSofort != nil {
				sourceSofort := applyConfiguredKeyedListShapes(rawValueSofort, attrValueToPlain(state.Sofort))
				if valueSofort, err := flattenPlainValue(sourceSofort, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "sofort", "raw response"); err != nil {
					return err
				} else {
					if typedSofort, ok := valueSofort.(types.Object); ok {
						state.Sofort = typedSofort
						assignedSofort = true
					}
				}
			}
		}
		if !assignedSofort {
			if !hasRaw {
				if responseValueSofort, ok := plainFromResponseField(obj, "Sofort"); ok {
					sourceSofort := applyConfiguredKeyedListShapes(responseValueSofort, attrValueToPlain(state.Sofort))
					if valueSofort, err := flattenPlainValue(
						sourceSofort,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"sofort",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSofort, ok := valueSofort.(types.Object); ok {
							state.Sofort = typedSofort
							assignedSofort = true
						}
					}
				}
			}
		}
		if !assignedSofort && hadRawSofort {
			if nullSofort, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedSofort, ok := nullSofort.(types.Object); ok {
					state.Sofort = typedSofort
				}
			}
		}
	}
	{
		assignedSunbit := false
		hadRawSunbit := false
		if rawValueSunbit, rawOk := plainValueAtPath(raw, "sunbit"); rawOk {
			hadRawSunbit = true
			if rawValueSunbit != nil {
				sourceSunbit := applyConfiguredKeyedListShapes(rawValueSunbit, attrValueToPlain(state.Sunbit))
				if valueSunbit, err := flattenPlainValue(sourceSunbit, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "sunbit", "raw response"); err != nil {
					return err
				} else {
					if typedSunbit, ok := valueSunbit.(types.Object); ok {
						state.Sunbit = typedSunbit
						assignedSunbit = true
					}
				}
			}
		}
		if !assignedSunbit {
			if !hasRaw {
				if responseValueSunbit, ok := plainFromResponseField(obj, "Sunbit"); ok {
					sourceSunbit := applyConfiguredKeyedListShapes(responseValueSunbit, attrValueToPlain(state.Sunbit))
					if valueSunbit, err := flattenPlainValue(
						sourceSunbit,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"sunbit",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSunbit, ok := valueSunbit.(types.Object); ok {
							state.Sunbit = typedSunbit
							assignedSunbit = true
						}
					}
				}
			}
		}
		if !assignedSunbit && hadRawSunbit {
			if nullSunbit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedSunbit, ok := nullSunbit.(types.Object); ok {
					state.Sunbit = typedSunbit
				}
			}
		}
	}
	{
		assignedSwish := false
		hadRawSwish := false
		if rawValueSwish, rawOk := plainValueAtPath(raw, "swish"); rawOk {
			hadRawSwish = true
			if rawValueSwish != nil {
				sourceSwish := applyConfiguredKeyedListShapes(rawValueSwish, attrValueToPlain(state.Swish))
				if valueSwish, err := flattenPlainValue(sourceSwish, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "swish", "raw response"); err != nil {
					return err
				} else {
					if typedSwish, ok := valueSwish.(types.Object); ok {
						state.Swish = typedSwish
						assignedSwish = true
					}
				}
			}
		}
		if !assignedSwish {
			if !hasRaw {
				if responseValueSwish, ok := plainFromResponseField(obj, "Swish"); ok {
					sourceSwish := applyConfiguredKeyedListShapes(responseValueSwish, attrValueToPlain(state.Swish))
					if valueSwish, err := flattenPlainValue(
						sourceSwish,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"swish",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSwish, ok := valueSwish.(types.Object); ok {
							state.Swish = typedSwish
							assignedSwish = true
						}
					}
				}
			}
		}
		if !assignedSwish && hadRawSwish {
			if nullSwish, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedSwish, ok := nullSwish.(types.Object); ok {
					state.Swish = typedSwish
				}
			}
		}
	}
	{
		assignedTWINT := false
		hadRawTWINT := false
		if rawValueTWINT, rawOk := plainValueAtPath(raw, "twint"); rawOk {
			hadRawTWINT = true
			if rawValueTWINT != nil {
				sourceTWINT := applyConfiguredKeyedListShapes(rawValueTWINT, attrValueToPlain(state.TWINT))
				if valueTWINT, err := flattenPlainValue(sourceTWINT, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "twint", "raw response"); err != nil {
					return err
				} else {
					if typedTWINT, ok := valueTWINT.(types.Object); ok {
						state.TWINT = typedTWINT
						assignedTWINT = true
					}
				}
			}
		}
		if !assignedTWINT {
			if !hasRaw {
				if responseValueTWINT, ok := plainFromResponseField(obj, "TWINT"); ok {
					sourceTWINT := applyConfiguredKeyedListShapes(responseValueTWINT, attrValueToPlain(state.TWINT))
					if valueTWINT, err := flattenPlainValue(
						sourceTWINT,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"twint",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTWINT, ok := valueTWINT.(types.Object); ok {
							state.TWINT = typedTWINT
							assignedTWINT = true
						}
					}
				}
			}
		}
		if !assignedTWINT && hadRawTWINT {
			if nullTWINT, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedTWINT, ok := nullTWINT.(types.Object); ok {
					state.TWINT = typedTWINT
				}
			}
		}
	}
	{
		assignedUpi := false
		hadRawUpi := false
		if rawValueUpi, rawOk := plainValueAtPath(raw, "upi"); rawOk {
			hadRawUpi = true
			if rawValueUpi != nil {
				sourceUpi := applyConfiguredKeyedListShapes(rawValueUpi, attrValueToPlain(state.Upi))
				if valueUpi, err := flattenPlainValue(sourceUpi, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "upi", "raw response"); err != nil {
					return err
				} else {
					if typedUpi, ok := valueUpi.(types.Object); ok {
						state.Upi = typedUpi
						assignedUpi = true
					}
				}
			}
		}
		if !assignedUpi {
			if !hasRaw {
				if responseValueUpi, ok := plainFromResponseField(obj, "Upi"); ok {
					sourceUpi := applyConfiguredKeyedListShapes(responseValueUpi, attrValueToPlain(state.Upi))
					if valueUpi, err := flattenPlainValue(
						sourceUpi,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"upi",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedUpi, ok := valueUpi.(types.Object); ok {
							state.Upi = typedUpi
							assignedUpi = true
						}
					}
				}
			}
		}
		if !assignedUpi && hadRawUpi {
			if nullUpi, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedUpi, ok := nullUpi.(types.Object); ok {
					state.Upi = typedUpi
				}
			}
		}
	}
	{
		assignedUSBankAccount := false
		hadRawUSBankAccount := false
		if rawValueUSBankAccount, rawOk := plainValueAtPath(raw, "us_bank_account"); rawOk {
			hadRawUSBankAccount = true
			if rawValueUSBankAccount != nil {
				sourceUSBankAccount := applyConfiguredKeyedListShapes(rawValueUSBankAccount, attrValueToPlain(state.USBankAccount))
				if valueUSBankAccount, err := flattenPlainValue(sourceUSBankAccount, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "us_bank_account", "raw response"); err != nil {
					return err
				} else {
					if typedUSBankAccount, ok := valueUSBankAccount.(types.Object); ok {
						state.USBankAccount = typedUSBankAccount
						assignedUSBankAccount = true
					}
				}
			}
		}
		if !assignedUSBankAccount {
			if !hasRaw {
				if responseValueUSBankAccount, ok := plainFromResponseField(obj, "USBankAccount"); ok {
					sourceUSBankAccount := applyConfiguredKeyedListShapes(responseValueUSBankAccount, attrValueToPlain(state.USBankAccount))
					if valueUSBankAccount, err := flattenPlainValue(
						sourceUSBankAccount,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"us_bank_account",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedUSBankAccount, ok := valueUSBankAccount.(types.Object); ok {
							state.USBankAccount = typedUSBankAccount
							assignedUSBankAccount = true
						}
					}
				}
			}
		}
		if !assignedUSBankAccount && hadRawUSBankAccount {
			if nullUSBankAccount, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedUSBankAccount, ok := nullUSBankAccount.(types.Object); ok {
					state.USBankAccount = typedUSBankAccount
				}
			}
		}
	}
	{
		assignedWeChatPay := false
		hadRawWeChatPay := false
		if rawValueWeChatPay, rawOk := plainValueAtPath(raw, "wechat_pay"); rawOk {
			hadRawWeChatPay = true
			if rawValueWeChatPay != nil {
				sourceWeChatPay := applyConfiguredKeyedListShapes(rawValueWeChatPay, attrValueToPlain(state.WeChatPay))
				if valueWeChatPay, err := flattenPlainValue(sourceWeChatPay, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "wechat_pay", "raw response"); err != nil {
					return err
				} else {
					if typedWeChatPay, ok := valueWeChatPay.(types.Object); ok {
						state.WeChatPay = typedWeChatPay
						assignedWeChatPay = true
					}
				}
			}
		}
		if !assignedWeChatPay {
			if !hasRaw {
				if responseValueWeChatPay, ok := plainFromResponseField(obj, "WeChatPay"); ok {
					sourceWeChatPay := applyConfiguredKeyedListShapes(responseValueWeChatPay, attrValueToPlain(state.WeChatPay))
					if valueWeChatPay, err := flattenPlainValue(
						sourceWeChatPay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"wechat_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedWeChatPay, ok := valueWeChatPay.(types.Object); ok {
							state.WeChatPay = typedWeChatPay
							assignedWeChatPay = true
						}
					}
				}
			}
		}
		if !assignedWeChatPay && hadRawWeChatPay {
			if nullWeChatPay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedWeChatPay, ok := nullWeChatPay.(types.Object); ok {
					state.WeChatPay = typedWeChatPay
				}
			}
		}
	}
	{
		assignedZip := false
		hadRawZip := false
		if rawValueZip, rawOk := plainValueAtPath(raw, "zip"); rawOk {
			hadRawZip = true
			if rawValueZip != nil {
				sourceZip := applyConfiguredKeyedListShapes(rawValueZip, attrValueToPlain(state.Zip))
				if valueZip, err := flattenPlainValue(sourceZip, types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}, "zip", "raw response"); err != nil {
					return err
				} else {
					if typedZip, ok := valueZip.(types.Object); ok {
						state.Zip = typedZip
						assignedZip = true
					}
				}
			}
		}
		if !assignedZip {
			if !hasRaw {
				if responseValueZip, ok := plainFromResponseField(obj, "Zip"); ok {
					sourceZip := applyConfiguredKeyedListShapes(responseValueZip, attrValueToPlain(state.Zip))
					if valueZip, err := flattenPlainValue(
						sourceZip,
						types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}},
						"zip",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedZip, ok := valueZip.(types.Object); ok {
							state.Zip = typedZip
							assignedZip = true
						}
					}
				}
			}
		}
		if !assignedZip && hadRawZip {
			if nullZip, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"available": types.BoolType, "display_preference": types.ObjectType{AttrTypes: map[string]attr.Type{"overridable": types.BoolType, "preference": types.StringType, "value": types.StringType}}}}); ok {
				if typedZip, ok := nullZip.(types.Object); ok {
					state.Zip = typedZip
				}
			}
		}
	}
	return nil
}
