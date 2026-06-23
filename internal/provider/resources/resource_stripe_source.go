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

var _ resource.Resource = &SourceResource{}

var _ resource.ResourceWithConfigure = &SourceResource{}

var _ resource.ResourceWithImportState = &SourceResource{}

func NewSourceResource() resource.Resource {
	return &SourceResource{}
}

type SourceResource struct {
	client *stripe.Client
}

type SourceResourceModel struct {
	Object              types.String `tfsdk:"object"`
	ACHCreditTransfer   types.Object `tfsdk:"ach_credit_transfer"`
	ACHDebit            types.Object `tfsdk:"ach_debit"`
	ACSSDebit           types.Object `tfsdk:"acss_debit"`
	Alipay              types.Object `tfsdk:"alipay"`
	AllowRedisplay      types.String `tfsdk:"allow_redisplay"`
	Amount              types.Int64  `tfsdk:"amount"`
	AUBECSDebit         types.Object `tfsdk:"au_becs_debit"`
	Bancontact          types.Object `tfsdk:"bancontact"`
	Card                types.Object `tfsdk:"card"`
	CardPresent         types.Object `tfsdk:"card_present"`
	ClientSecret        types.String `tfsdk:"client_secret"`
	CodeVerification    types.Object `tfsdk:"code_verification"`
	Created             types.Int64  `tfsdk:"created"`
	Currency            types.String `tfsdk:"currency"`
	Customer            types.String `tfsdk:"customer"`
	EPS                 types.Object `tfsdk:"eps"`
	Flow                types.String `tfsdk:"flow"`
	Giropay             types.Object `tfsdk:"giropay"`
	ID                  types.String `tfsdk:"id"`
	IDEAL               types.Object `tfsdk:"ideal"`
	Klarna              types.Object `tfsdk:"klarna"`
	Livemode            types.Bool   `tfsdk:"livemode"`
	Metadata            types.Map    `tfsdk:"metadata"`
	Multibanco          types.Object `tfsdk:"multibanco"`
	Owner               types.Object `tfsdk:"owner"`
	P24                 types.Object `tfsdk:"p24"`
	Receiver            types.Object `tfsdk:"receiver"`
	Redirect            types.Object `tfsdk:"redirect"`
	SEPACreditTransfer  types.Object `tfsdk:"sepa_credit_transfer"`
	SEPADebit           types.Object `tfsdk:"sepa_debit"`
	Sofort              types.Object `tfsdk:"sofort"`
	SourceOrder         types.Object `tfsdk:"source_order"`
	StatementDescriptor types.String `tfsdk:"statement_descriptor"`
	Status              types.String `tfsdk:"status"`
	ThreeDSecure        types.Object `tfsdk:"three_d_secure"`
	Type                types.String `tfsdk:"type"`
	Usage               types.String `tfsdk:"usage"`
	WeChat              types.Object `tfsdk:"wechat"`
	Mandate             types.Object `tfsdk:"mandate"`
	OriginalSource      types.String `tfsdk:"original_source"`
	Token               types.String `tfsdk:"token"`
}

func (r *SourceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source"
}

func (r *SourceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "`Source` objects allow you to accept a variety of payment methods. They\nrepresent a customer's payment instrument, and can be used with the Stripe API\njust like a `Card` object: once chargeable, they can be charged, or can be\nattached to customers.\n\nStripe doesn't recommend using the deprecated [Sources API](https://docs.stripe.com/api/sources).\nWe recommend that you adopt the [PaymentMethods API](https://docs.stripe.com/api/payment_methods).\nThis newer API provides access to our latest features and payment method types.\n\nRelated guides: [Sources API](https://docs.stripe.com/sources) and [Sources & Customers](https://docs.stripe.com/sources/customers).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("source")},
			},
			"ach_credit_transfer": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"account_number": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bank_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"fingerprint": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_type": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_routing_number": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"routing_number": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"swift_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"ach_debit": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bank_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"fingerprint": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"routing_number": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"type": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"acss_debit": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bank_address_city": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bank_address_line_1": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bank_address_line_2": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bank_address_postal_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bank_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"category": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"fingerprint": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"routing_number": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"alipay": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"data_string": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"native_url": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"statement_descriptor": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"allow_redisplay": schema.StringAttribute{
				Computed:      true,
				Description:   "This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow. The field defaults to “unspecified”.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("always", "limited", "unspecified")},
			},
			"amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount associated with the source. This is the amount for which the source will be chargeable once ready. Required for `single_use` sources.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"au_becs_debit": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bsb_number": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"fingerprint": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"bancontact": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bank_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bank_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bic": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"iban_last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"preferred_language": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"statement_descriptor": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"card": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"address_line1_check": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"address_zip_check": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"brand": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"cvc_check": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"description": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"dynamic_last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"exp_month": schema.Int64Attribute{
						Computed: true,

						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"exp_year": schema.Int64Attribute{
						Computed: true,

						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"fingerprint": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"funding": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"iin": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"issuer": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"three_d_secure": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"tokenization_method": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"card_present": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"application_cryptogram": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"application_preferred_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"authorization_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"authorization_response_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"brand": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"cvm_type": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"data_type": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"dedicated_file_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"description": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"emv_auth_data": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"evidence_customer_signature": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"evidence_transaction_certificate": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"exp_month": schema.Int64Attribute{
						Computed: true,

						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"exp_year": schema.Int64Attribute{
						Computed: true,

						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"fingerprint": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"funding": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"iin": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"issuer": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pos_device_id": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pos_entry_mode": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"read_method": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"reader": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"terminal_verification_results": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"transaction_status_information": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"client_secret": schema.StringAttribute{
				Computed:      true,
				Description:   "The client secret of the source. Used for client-side retrieval using a publishable key.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"code_verification": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"attempts_remaining": schema.Int64Attribute{
						Computed:      true,
						Description:   "The number of attempts remaining to authenticate the source object with a verification code.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the code verification, either `pending` (awaiting verification, `attempts_remaining` should be greater than 0), `succeeded` (successful verification) or `failed` (failed verification, cannot be verified anymore as `attempts_remaining` should be 0).",
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
				Description:   "Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) associated with the source. This is the currency for which the source will be chargeable once ready. Required for `single_use` sources.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ID of the customer to which this source is attached. This will not be present when the source has not been attached to a customer.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"eps": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"reference": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"statement_descriptor": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"flow": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The authentication `flow` of the source. `flow` is one of `redirect`, `receiver`, `code_verification`, `none`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"giropay": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bank_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bank_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bic": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"statement_descriptor": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ideal": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bank": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bic": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"iban_last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"statement_descriptor": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"klarna": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"background_image_url": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"client_token": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"first_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"last_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"locale": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"logo_url": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"page_title": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_later_asset_urls_descriptive": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_later_asset_urls_standard": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_later_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_later_redirect_url": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_now_asset_urls_descriptive": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_now_asset_urls_standard": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_now_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_now_redirect_url": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_over_time_asset_urls_descriptive": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_over_time_asset_urls_standard": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_over_time_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"pay_over_time_redirect_url": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"payment_method_categories": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"purchase_country": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"purchase_type": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"redirect_url": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"shipping_delay": schema.Int64Attribute{
						Computed: true,

						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"shipping_first_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"shipping_last_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"multibanco": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"entity": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"reference": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_city": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_country": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_line1": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_line2": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_postal_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_state": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_iban": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"owner": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Information about the owner of the payment instrument that may be used or required by particular source types.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"address": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Owner's address.",
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
					"email": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Owner's email address.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"name": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Owner's full name.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"phone": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Owner's phone number (including extension).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"verified_address": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Verified owner's address. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.",
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
					"verified_email": schema.StringAttribute{
						Computed:      true,
						Description:   "Verified owner's email address. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"verified_name": schema.StringAttribute{
						Computed:      true,
						Description:   "Verified owner's full name. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"verified_phone": schema.StringAttribute{
						Computed:      true,
						Description:   "Verified owner's phone number (including extension). Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"p24": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"reference": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"receiver": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Computed:      true,
						Description:   "The address of the receiver source. This is the value that should be communicated to the customer to send their funds to.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"amount_charged": schema.Int64Attribute{
						Computed:      true,
						Description:   "The total amount that was moved to your balance. This is almost always equal to the amount charged. In rare cases when customers deposit excess funds and we are unable to refund those, those funds get moved to your balance and show up in amount_charged as well. The amount charged is expressed in the source's currency.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"amount_received": schema.Int64Attribute{
						Computed:      true,
						Description:   "The total amount received by the receiver source. `amount_received = amount_returned + amount_charged` should be true for consumed sources unless customers deposit excess funds. The amount received is expressed in the source's currency.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"amount_returned": schema.Int64Attribute{
						Computed:      true,
						Description:   "The total amount that was returned to the customer. The amount returned is expressed in the source's currency.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"refund_attributes_method": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Type of refund attribute method, one of `email`, `manual`, or `none`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"refund_attributes_status": schema.StringAttribute{
						Computed:      true,
						Description:   "Type of refund attribute status, one of `missing`, `requested`, or `available`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"redirect": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"failure_reason": schema.StringAttribute{
						Computed:      true,
						Description:   "The failure reason for the redirect, either `user_abort` (the customer aborted or dropped out of the redirect flow), `declined` (the authentication failed or the transaction was declined), or `processing_error` (the redirect failed due to a technical error). Present only if the redirect status is `failed`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"return_url": schema.StringAttribute{
						Required:      true,
						Description:   "The URL you provide to redirect the customer to after they authenticated their payment.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the redirect, either `pending` (ready to be used by your customer to authenticate the transaction), `succeeded` (successful authentication, cannot be reused) or `not_required` (redirect should not be used) or `failed` (failed authentication, cannot be reused).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"url": schema.StringAttribute{
						Computed:      true,
						Description:   "The URL provided to you to redirect a customer to as part of a `redirect` authentication flow.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"sepa_credit_transfer": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bank_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bic": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"iban": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_city": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_country": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_line1": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_line2": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_postal_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_address_state": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_account_holder_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"refund_iban": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"sepa_debit": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bank_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"branch_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"fingerprint": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"mandate_reference": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"mandate_url": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"sofort": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"bank_code": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bank_name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"bic": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"iban_last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"preferred_language": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"statement_descriptor": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"source_order": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"amount": schema.Int64Attribute{
						Computed:      true,
						Description:   "A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount for the order.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"currency": schema.StringAttribute{
						Computed:      true,
						Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"email": schema.StringAttribute{
						Computed:      true,
						Description:   "The email address of the customer placing the order.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"items": schema.ListNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "List of items constituting the order.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"amount": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "The amount (price) for this order item.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"currency": schema.StringAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "This currency of this order item. Required when `amount` is present.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"description": schema.StringAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "Human-readable description for this order item.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"parent": schema.StringAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "The ID of the associated object for this line item. Expandable if not null (e.g., expandable to a SKU).",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"quantity": schema.Int64Attribute{
									Optional:      true,
									Computed:      true,
									Description:   "The quantity of this order item. When type is `sku`, this is the number of instances of the SKU to be ordered.",
									PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
								},
								"type": schema.StringAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "The type of this order item. Must be `sku`, `tax`, or `shipping`.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
							},
						},
					},
					"shipping": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

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
										Required:    true,
										Description: "Address line 1, such as the street, PO Box, or company name.",
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
								Optional:      true,
								Computed:      true,
								Description:   "The delivery service that shipped a physical product, such as Fedex, UPS, USPS, etc.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"name": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Recipient name.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"phone": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Recipient phone (including extension).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"tracking_number": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"statement_descriptor": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Extra information about a source. This will appear on your customer's statement every time you charge the source.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The status of the source, one of `canceled`, `chargeable`, `consumed`, `failed`, or `pending`. Only `chargeable` sources can be used to create a charge.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"three_d_secure": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"address_line1_check": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"address_zip_check": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"authenticated": schema.BoolAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"brand": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"card": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"customer": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"cvc_check": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"description": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"dynamic_last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"exp_month": schema.Int64Attribute{
						Computed: true,

						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"exp_year": schema.Int64Attribute{
						Computed: true,

						PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
					},
					"fingerprint": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"funding": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"iin": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"issuer": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"last4": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"name": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"three_d_secure": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"tokenization_method": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The `type` of the source. The `type` is a payment method, one of `ach_credit_transfer`, `ach_debit`, `alipay`, `bancontact`, `card`, `card_present`, `eps`, `giropay`, `ideal`, `multibanco`, `klarna`, `p24`, `sepa_debit`, `sofort`, `three_d_secure`, or `wechat`. An additional hash is included on the source with a name matching this value. It contains additional information specific to the [payment method](https://docs.stripe.com/sources) used.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("ach_credit_transfer", "ach_debit", "acss_debit", "alipay", "au_becs_debit", "bancontact", "card", "card_present", "eps", "giropay", "ideal", "klarna", "multibanco", "p24", "sepa_credit_transfer", "sepa_debit", "sofort", "three_d_secure", "wechat")},
			},
			"usage": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Either `reusable` or `single_use`. Whether this source should be reusable or not. Some source types may or may not be reusable by construction, while others may leave the option at creation. If an incompatible value is passed, an error will be returned.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"wechat": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"prepay_id": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"qr_code_url": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"statement_descriptor": schema.StringAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"mandate": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Information about a mandate possibility attached to a source object (generally for bank debits) as well as its acceptance status.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"acceptance": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "The parameters required to notify Stripe of a mandate acceptance or refusal by the customer.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"date": schema.Int64Attribute{
								Optional:    true,
								Description: "The Unix timestamp (in seconds) when the mandate was accepted or refused by the customer.",
								WriteOnly:   true,
							},
							"ip": schema.StringAttribute{
								Optional:    true,
								Description: "The IP address from which the mandate was accepted or refused by the customer.",
								WriteOnly:   true,
							},
							"offline": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "The parameters required to store a mandate accepted offline. Should only be set if `mandate[type]` is `offline`",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"contact_email": schema.StringAttribute{
										Required:    true,
										Description: "An email to contact you with if a copy of the mandate is requested, required if `type` is `offline`.",
										WriteOnly:   true,
									},
								},
							},
							"online": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "The parameters required to store a mandate accepted online. Should only be set if `mandate[type]` is `online`",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"date": schema.Int64Attribute{
										Optional:    true,
										Description: "The Unix timestamp (in seconds) when the mandate was accepted or refused by the customer.",
										WriteOnly:   true,
									},
									"ip": schema.StringAttribute{
										Optional:    true,
										Description: "The IP address from which the mandate was accepted or refused by the customer.",
										WriteOnly:   true,
									},
									"user_agent": schema.StringAttribute{
										Optional:    true,
										Description: "The user agent of the browser from which the mandate was accepted or refused by the customer.",
										WriteOnly:   true,
									},
								},
							},
							"status": schema.StringAttribute{
								Required:    true,
								Description: "The status of the mandate acceptance. Either `accepted` (the mandate was accepted) or `refused` (the mandate was refused).",
								WriteOnly:   true,
							},
							"type": schema.StringAttribute{
								Optional:    true,
								Description: "The type of acceptance information included with the mandate. Either `online` or `offline`",
								WriteOnly:   true,
							},
							"user_agent": schema.StringAttribute{
								Optional:    true,
								Description: "The user agent of the browser from which the mandate was accepted or refused by the customer.",
								WriteOnly:   true,
							},
						},
					},
					"amount": schema.Int64Attribute{
						Optional:    true,
						Description: "The amount specified by the mandate. (Leave null for a mandate covering all amounts)",
						WriteOnly:   true,
					},
					"currency": schema.StringAttribute{
						Optional:    true,
						Description: "The currency specified by the mandate. (Must match `currency` of the source)",
						WriteOnly:   true,
					},
					"interval": schema.StringAttribute{
						Optional:    true,
						Description: "The interval of debits permitted by the mandate. Either `one_time` (just permitting a single debit), `scheduled` (with debits on an agreed schedule or for clearly-defined events), or `variable`(for debits with any frequency)",
						WriteOnly:   true,
					},
					"notification_method": schema.StringAttribute{
						Optional:    true,
						Description: "The method Stripe should use to notify the customer of upcoming debit instructions and/or mandate confirmation as required by the underlying debit network. Either `email` (an email is sent directly to the customer), `manual` (a `source.mandate_notification` event is sent to your webhooks endpoint and you should handle the notification) or `none` (the underlying debit network does not require any notification).",
						WriteOnly:   true,
					},
				},
			},
			"original_source": schema.StringAttribute{
				Optional:      true,
				Description:   "The source to share.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"token": schema.StringAttribute{
				Optional:      true,
				Description:   "An optional token used to create the source. When passed, token properties will override source parameters.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
		},
	}
}

func (r *SourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SourceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config SourceResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Mandate"}, []string{"Mandate", "acceptance"}, []string{"Mandate", "acceptance", "date"}, []string{"Mandate", "acceptance", "ip"}, []string{"Mandate", "acceptance", "offline"}, []string{"Mandate", "acceptance", "offline", "contact_email"}, []string{"Mandate", "acceptance", "online"}, []string{"Mandate", "acceptance", "online", "date"}, []string{"Mandate", "acceptance", "online", "ip"}, []string{"Mandate", "acceptance", "online", "user_agent"}, []string{"Mandate", "acceptance", "status"}, []string{"Mandate", "acceptance", "type"}, []string{"Mandate", "acceptance", "user_agent"}, []string{"Mandate", "amount"}, []string{"Mandate", "currency"}, []string{"Mandate", "interval"}, []string{"Mandate", "notification_method"}, []string{"OriginalSource"}, []string{"Token"}})

	params, err := expandSourceCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Source create params", err.Error())
		return
	}

	obj, err := r.client.V1Sources.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Source", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Sources.B, r.client.V1Sources.Key, stripe.FormatURLPath("/v1/sources/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Source create raw response", err.Error())
		return
	}

	if err := flattenSource(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Source create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Mandate"}, []string{"Mandate", "acceptance"}, []string{"Mandate", "acceptance", "date"}, []string{"Mandate", "acceptance", "ip"}, []string{"Mandate", "acceptance", "offline"}, []string{"Mandate", "acceptance", "offline", "contact_email"}, []string{"Mandate", "acceptance", "online"}, []string{"Mandate", "acceptance", "online", "date"}, []string{"Mandate", "acceptance", "online", "ip"}, []string{"Mandate", "acceptance", "online", "user_agent"}, []string{"Mandate", "acceptance", "status"}, []string{"Mandate", "acceptance", "type"}, []string{"Mandate", "acceptance", "user_agent"}, []string{"Mandate", "amount"}, []string{"Mandate", "currency"}, []string{"Mandate", "interval"}, []string{"Mandate", "notification_method"}, []string{"OriginalSource"}, []string{"Token"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState SourceResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state SourceResourceModel
	state = priorState

	obj, err := r.client.V1Sources.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Source", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Sources.B, r.client.V1Sources.Key, stripe.FormatURLPath("/v1/sources/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Source raw response", err.Error())
		return
	}

	if err := flattenSource(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Source read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"Mandate"}, []string{"Mandate", "acceptance"}, []string{"Mandate", "acceptance", "date"}, []string{"Mandate", "acceptance", "ip"}, []string{"Mandate", "acceptance", "offline"}, []string{"Mandate", "acceptance", "offline", "contact_email"}, []string{"Mandate", "acceptance", "online"}, []string{"Mandate", "acceptance", "online", "date"}, []string{"Mandate", "acceptance", "online", "ip"}, []string{"Mandate", "acceptance", "online", "user_agent"}, []string{"Mandate", "acceptance", "status"}, []string{"Mandate", "acceptance", "type"}, []string{"Mandate", "acceptance", "user_agent"}, []string{"Mandate", "amount"}, []string{"Mandate", "currency"}, []string{"Mandate", "interval"}, []string{"Mandate", "notification_method"}, []string{"OriginalSource"}, []string{"Token"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *SourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SourceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config SourceResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Mandate"}, []string{"Mandate", "acceptance"}, []string{"Mandate", "acceptance", "date"}, []string{"Mandate", "acceptance", "ip"}, []string{"Mandate", "acceptance", "offline"}, []string{"Mandate", "acceptance", "offline", "contact_email"}, []string{"Mandate", "acceptance", "online"}, []string{"Mandate", "acceptance", "online", "date"}, []string{"Mandate", "acceptance", "online", "ip"}, []string{"Mandate", "acceptance", "online", "user_agent"}, []string{"Mandate", "acceptance", "status"}, []string{"Mandate", "acceptance", "type"}, []string{"Mandate", "acceptance", "user_agent"}, []string{"Mandate", "amount"}, []string{"Mandate", "currency"}, []string{"Mandate", "interval"}, []string{"Mandate", "notification_method"}, []string{"OriginalSource"}, []string{"Token"}})

	var state SourceResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"OriginalSource"}, []string{"Token"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"OriginalSource"}, []string{"Token"}})

	params, err := expandSourceUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Source update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Source update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Sources.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Source", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Sources.B, r.client.V1Sources.Key, stripe.FormatURLPath("/v1/sources/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Source update raw response", err.Error())
		return
	}

	if err := flattenSource(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Source update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Mandate"}, []string{"Mandate", "acceptance"}, []string{"Mandate", "acceptance", "date"}, []string{"Mandate", "acceptance", "ip"}, []string{"Mandate", "acceptance", "offline"}, []string{"Mandate", "acceptance", "offline", "contact_email"}, []string{"Mandate", "acceptance", "online"}, []string{"Mandate", "acceptance", "online", "date"}, []string{"Mandate", "acceptance", "online", "ip"}, []string{"Mandate", "acceptance", "online", "user_agent"}, []string{"Mandate", "acceptance", "status"}, []string{"Mandate", "acceptance", "type"}, []string{"Mandate", "acceptance", "user_agent"}, []string{"Mandate", "amount"}, []string{"Mandate", "currency"}, []string{"Mandate", "interval"}, []string{"Mandate", "notification_method"}, []string{"OriginalSource"}, []string{"Token"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SourceResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *SourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandSourceCreate(plan SourceResourceModel) (*stripe.SourceCreateParams, error) {
	params := &stripe.SourceCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Customer", "Customer", plan.Customer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer", params)
		}
	}
	if !plan.Flow.IsNull() && !plan.Flow.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Flow", "Flow", plan.Flow.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "flow", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Owner.IsNull() && !plan.Owner.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Owner", plan.Owner) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "owner", params)
		}
	}
	if !plan.Receiver.IsNull() && !plan.Receiver.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Receiver", plan.Receiver) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "receiver", params)
		}
	}
	if !plan.Redirect.IsNull() && !plan.Redirect.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Redirect", plan.Redirect) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "redirect", params)
		}
	}
	if !plan.SourceOrder.IsNull() && !plan.SourceOrder.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SourceOrder", plan.SourceOrder) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "source_order", params)
		}
	}
	if !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "statement_descriptor", params)
		}
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}
	if !plan.Usage.IsNull() && !plan.Usage.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Usage", "Usage", plan.Usage.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "usage", params)
		}
	}
	if !plan.Mandate.IsNull() && !plan.Mandate.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Mandate", plan.Mandate) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "mandate", params)
		}
	}
	if !plan.OriginalSource.IsNull() && !plan.OriginalSource.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "OriginalSource", "OriginalSource", plan.OriginalSource.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "original_source", params)
		}
	}
	if !plan.Token.IsNull() && !plan.Token.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Token", "Token", plan.Token.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "token", params)
		}
	}

	return params, nil
}

func expandSourceUpdate(plan SourceResourceModel, state SourceResourceModel) (*stripe.SourceUpdateParams, error) {
	params := &stripe.SourceUpdateParams{}

	if !plan.Amount.Equal(state.Amount) && !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Metadata.Equal(state.Metadata) && !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			if !plan.Metadata.Equal(state.Metadata) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "metadata", params)
			}
		}
	}
	if !plan.Owner.Equal(state.Owner) && !plan.Owner.IsNull() && !plan.Owner.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Owner", plan.Owner) {
			if !plan.Owner.Equal(state.Owner) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "owner", params)
			}
		}
	}
	if !plan.SourceOrder.Equal(state.SourceOrder) && !plan.SourceOrder.IsNull() && !plan.SourceOrder.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SourceOrder", plan.SourceOrder) {
			if !plan.SourceOrder.Equal(state.SourceOrder) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "source_order", params)
			}
		}
	}
	if !plan.Mandate.Equal(state.Mandate) && !plan.Mandate.IsNull() && !plan.Mandate.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Mandate", plan.Mandate) {
			if !plan.Mandate.Equal(state.Mandate) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "mandate", params)
			}
		}
	}

	return params, nil
}

func flattenSource(obj *stripe.Source, state *SourceResourceModel) error {
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
		assignedACHCreditTransfer := false
		hadRawACHCreditTransfer := false
		if rawValueACHCreditTransfer, rawOk := plainValueAtPath(raw, "ach_credit_transfer"); rawOk {
			hadRawACHCreditTransfer = true
			if rawValueACHCreditTransfer != nil {
				sourceACHCreditTransfer := applyConfiguredKeyedListShapes(rawValueACHCreditTransfer, attrValueToPlain(state.ACHCreditTransfer))
				if valueACHCreditTransfer, err := flattenPlainValue(sourceACHCreditTransfer, types.ObjectType{AttrTypes: map[string]attr.Type{"account_number": types.StringType, "bank_name": types.StringType, "fingerprint": types.StringType, "refund_account_holder_name": types.StringType, "refund_account_holder_type": types.StringType, "refund_routing_number": types.StringType, "routing_number": types.StringType, "swift_code": types.StringType}}, "ach_credit_transfer", "raw response"); err != nil {
					return err
				} else {
					if typedACHCreditTransfer, ok := valueACHCreditTransfer.(types.Object); ok {
						state.ACHCreditTransfer = typedACHCreditTransfer
						assignedACHCreditTransfer = true
					}
				}
			}
		}
		if !assignedACHCreditTransfer {
			if !hasRaw {
				if responseValueACHCreditTransfer, ok := plainFromResponseField(obj, "ACHCreditTransfer"); ok {
					sourceACHCreditTransfer := applyConfiguredKeyedListShapes(responseValueACHCreditTransfer, attrValueToPlain(state.ACHCreditTransfer))
					if valueACHCreditTransfer, err := flattenPlainValue(
						sourceACHCreditTransfer,
						types.ObjectType{AttrTypes: map[string]attr.Type{"account_number": types.StringType, "bank_name": types.StringType, "fingerprint": types.StringType, "refund_account_holder_name": types.StringType, "refund_account_holder_type": types.StringType, "refund_routing_number": types.StringType, "routing_number": types.StringType, "swift_code": types.StringType}},
						"ach_credit_transfer",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedACHCreditTransfer, ok := valueACHCreditTransfer.(types.Object); ok {
							state.ACHCreditTransfer = typedACHCreditTransfer
							assignedACHCreditTransfer = true
						}
					}
				}
			}
		}
		if !assignedACHCreditTransfer && hadRawACHCreditTransfer {
			if nullACHCreditTransfer, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account_number": types.StringType, "bank_name": types.StringType, "fingerprint": types.StringType, "refund_account_holder_name": types.StringType, "refund_account_holder_type": types.StringType, "refund_routing_number": types.StringType, "routing_number": types.StringType, "swift_code": types.StringType}}); ok {
				if typedACHCreditTransfer, ok := nullACHCreditTransfer.(types.Object); ok {
					state.ACHCreditTransfer = typedACHCreditTransfer
				}
			}
		}
	}
	{
		assignedACHDebit := false
		hadRawACHDebit := false
		if rawValueACHDebit, rawOk := plainValueAtPath(raw, "ach_debit"); rawOk {
			hadRawACHDebit = true
			if rawValueACHDebit != nil {
				sourceACHDebit := applyConfiguredKeyedListShapes(rawValueACHDebit, attrValueToPlain(state.ACHDebit))
				if valueACHDebit, err := flattenPlainValue(sourceACHDebit, types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "routing_number": types.StringType, "type": types.StringType}}, "ach_debit", "raw response"); err != nil {
					return err
				} else {
					if typedACHDebit, ok := valueACHDebit.(types.Object); ok {
						state.ACHDebit = typedACHDebit
						assignedACHDebit = true
					}
				}
			}
		}
		if !assignedACHDebit {
			if !hasRaw {
				if responseValueACHDebit, ok := plainFromResponseField(obj, "ACHDebit"); ok {
					sourceACHDebit := applyConfiguredKeyedListShapes(responseValueACHDebit, attrValueToPlain(state.ACHDebit))
					if valueACHDebit, err := flattenPlainValue(
						sourceACHDebit,
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "routing_number": types.StringType, "type": types.StringType}},
						"ach_debit",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedACHDebit, ok := valueACHDebit.(types.Object); ok {
							state.ACHDebit = typedACHDebit
							assignedACHDebit = true
						}
					}
				}
			}
		}
		if !assignedACHDebit && hadRawACHDebit {
			if nullACHDebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "routing_number": types.StringType, "type": types.StringType}}); ok {
				if typedACHDebit, ok := nullACHDebit.(types.Object); ok {
					state.ACHDebit = typedACHDebit
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
				if valueACSSDebit, err := flattenPlainValue(sourceACSSDebit, types.ObjectType{AttrTypes: map[string]attr.Type{"bank_address_city": types.StringType, "bank_address_line_1": types.StringType, "bank_address_line_2": types.StringType, "bank_address_postal_code": types.StringType, "bank_name": types.StringType, "category": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "routing_number": types.StringType}}, "acss_debit", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank_address_city": types.StringType, "bank_address_line_1": types.StringType, "bank_address_line_2": types.StringType, "bank_address_postal_code": types.StringType, "bank_name": types.StringType, "category": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "routing_number": types.StringType}},
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
			if nullACSSDebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank_address_city": types.StringType, "bank_address_line_1": types.StringType, "bank_address_line_2": types.StringType, "bank_address_postal_code": types.StringType, "bank_name": types.StringType, "category": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "routing_number": types.StringType}}); ok {
				if typedACSSDebit, ok := nullACSSDebit.(types.Object); ok {
					state.ACSSDebit = typedACSSDebit
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
				if valueAlipay, err := flattenPlainValue(sourceAlipay, types.ObjectType{AttrTypes: map[string]attr.Type{"data_string": types.StringType, "native_url": types.StringType, "statement_descriptor": types.StringType}}, "alipay", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"data_string": types.StringType, "native_url": types.StringType, "statement_descriptor": types.StringType}},
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
			if nullAlipay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"data_string": types.StringType, "native_url": types.StringType, "statement_descriptor": types.StringType}}); ok {
				if typedAlipay, ok := nullAlipay.(types.Object); ok {
					state.Alipay = typedAlipay
				}
			}
		}
	}
	{
		if rawValueAllowRedisplay, rawOk := plainValueAtPath(raw, "allow_redisplay"); rawOk {
			if valueAllowRedisplay, err := flattenPlainValue(rawValueAllowRedisplay, types.StringType, "allow_redisplay", "raw response"); err != nil {
				return err
			} else {
				if typedAllowRedisplay, ok := valueAllowRedisplay.(types.String); ok {
					state.AllowRedisplay = typedAllowRedisplay
				}
			}
		} else if !hasRaw {
			if responseValueAllowRedisplay, ok := plainFromResponseField(obj, "AllowRedisplay"); ok {
				if valueAllowRedisplay, err := flattenPlainValue(responseValueAllowRedisplay, types.StringType, "allow_redisplay", "response struct"); err != nil {
					return err
				} else {
					if typedAllowRedisplay, ok := valueAllowRedisplay.(types.String); ok {
						state.AllowRedisplay = typedAllowRedisplay
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
		assignedAUBECSDebit := false
		hadRawAUBECSDebit := false
		if rawValueAUBECSDebit, rawOk := plainValueAtPath(raw, "au_becs_debit"); rawOk {
			hadRawAUBECSDebit = true
			if rawValueAUBECSDebit != nil {
				sourceAUBECSDebit := applyConfiguredKeyedListShapes(rawValueAUBECSDebit, attrValueToPlain(state.AUBECSDebit))
				if valueAUBECSDebit, err := flattenPlainValue(sourceAUBECSDebit, types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "fingerprint": types.StringType, "last4": types.StringType}}, "au_becs_debit", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "fingerprint": types.StringType, "last4": types.StringType}},
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
			if nullAUBECSDebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bsb_number": types.StringType, "fingerprint": types.StringType, "last4": types.StringType}}); ok {
				if typedAUBECSDebit, ok := nullAUBECSDebit.(types.Object); ok {
					state.AUBECSDebit = typedAUBECSDebit
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
				if valueBancontact, err := flattenPlainValue(sourceBancontact, types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "statement_descriptor": types.StringType}}, "bancontact", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "statement_descriptor": types.StringType}},
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
			if nullBancontact, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "statement_descriptor": types.StringType}}); ok {
				if typedBancontact, ok := nullBancontact.(types.Object); ok {
					state.Bancontact = typedBancontact
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
				if valueCard, err := flattenPlainValue(sourceCard, types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_zip_check": types.StringType, "brand": types.StringType, "country": types.StringType, "cvc_check": types.StringType, "description": types.StringType, "dynamic_last4": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "name": types.StringType, "three_d_secure": types.StringType, "tokenization_method": types.StringType}}, "card", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_zip_check": types.StringType, "brand": types.StringType, "country": types.StringType, "cvc_check": types.StringType, "description": types.StringType, "dynamic_last4": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "name": types.StringType, "three_d_secure": types.StringType, "tokenization_method": types.StringType}},
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
			if nullCard, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_zip_check": types.StringType, "brand": types.StringType, "country": types.StringType, "cvc_check": types.StringType, "description": types.StringType, "dynamic_last4": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "name": types.StringType, "three_d_secure": types.StringType, "tokenization_method": types.StringType}}); ok {
				if typedCard, ok := nullCard.(types.Object); ok {
					state.Card = typedCard
				}
			}
		}
	}
	{
		assignedCardPresent := false
		hadRawCardPresent := false
		if rawValueCardPresent, rawOk := plainValueAtPath(raw, "card_present"); rawOk {
			hadRawCardPresent = true
			if rawValueCardPresent != nil {
				sourceCardPresent := applyConfiguredKeyedListShapes(rawValueCardPresent, attrValueToPlain(state.CardPresent))
				if valueCardPresent, err := flattenPlainValue(sourceCardPresent, types.ObjectType{AttrTypes: map[string]attr.Type{"application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "brand": types.StringType, "country": types.StringType, "cvm_type": types.StringType, "data_type": types.StringType, "dedicated_file_name": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "evidence_customer_signature": types.StringType, "evidence_transaction_certificate": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "pos_device_id": types.StringType, "pos_entry_mode": types.StringType, "read_method": types.StringType, "reader": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}, "card_present", "raw response"); err != nil {
					return err
				} else {
					if typedCardPresent, ok := valueCardPresent.(types.Object); ok {
						state.CardPresent = typedCardPresent
						assignedCardPresent = true
					}
				}
			}
		}
		if !assignedCardPresent {
			if !hasRaw {
				if responseValueCardPresent, ok := plainFromResponseField(obj, "CardPresent"); ok {
					sourceCardPresent := applyConfiguredKeyedListShapes(responseValueCardPresent, attrValueToPlain(state.CardPresent))
					if valueCardPresent, err := flattenPlainValue(
						sourceCardPresent,
						types.ObjectType{AttrTypes: map[string]attr.Type{"application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "brand": types.StringType, "country": types.StringType, "cvm_type": types.StringType, "data_type": types.StringType, "dedicated_file_name": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "evidence_customer_signature": types.StringType, "evidence_transaction_certificate": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "pos_device_id": types.StringType, "pos_entry_mode": types.StringType, "read_method": types.StringType, "reader": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}},
						"card_present",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCardPresent, ok := valueCardPresent.(types.Object); ok {
							state.CardPresent = typedCardPresent
							assignedCardPresent = true
						}
					}
				}
			}
		}
		if !assignedCardPresent && hadRawCardPresent {
			if nullCardPresent, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"application_cryptogram": types.StringType, "application_preferred_name": types.StringType, "authorization_code": types.StringType, "authorization_response_code": types.StringType, "brand": types.StringType, "country": types.StringType, "cvm_type": types.StringType, "data_type": types.StringType, "dedicated_file_name": types.StringType, "description": types.StringType, "emv_auth_data": types.StringType, "evidence_customer_signature": types.StringType, "evidence_transaction_certificate": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "pos_device_id": types.StringType, "pos_entry_mode": types.StringType, "read_method": types.StringType, "reader": types.StringType, "terminal_verification_results": types.StringType, "transaction_status_information": types.StringType}}); ok {
				if typedCardPresent, ok := nullCardPresent.(types.Object); ok {
					state.CardPresent = typedCardPresent
				}
			}
		}
	}
	{
		if rawValueClientSecret, rawOk := plainValueAtPath(raw, "client_secret"); rawOk {
			if valueClientSecret, err := flattenPlainValue(rawValueClientSecret, types.StringType, "client_secret", "raw response"); err != nil {
				return err
			} else {
				if typedClientSecret, ok := valueClientSecret.(types.String); ok {
					state.ClientSecret = typedClientSecret
				}
			}
		} else if !hasRaw {
			if responseValueClientSecret, ok := plainFromResponseField(obj, "ClientSecret"); ok {
				if valueClientSecret, err := flattenPlainValue(responseValueClientSecret, types.StringType, "client_secret", "response struct"); err != nil {
					return err
				} else {
					if typedClientSecret, ok := valueClientSecret.(types.String); ok {
						state.ClientSecret = typedClientSecret
					}
				}
			}
		}
	}
	{
		assignedCodeVerification := false
		hadRawCodeVerification := false
		if rawValueCodeVerification, rawOk := plainValueAtPath(raw, "code_verification"); rawOk {
			hadRawCodeVerification = true
			if rawValueCodeVerification != nil {
				sourceCodeVerification := applyConfiguredKeyedListShapes(rawValueCodeVerification, attrValueToPlain(state.CodeVerification))
				if valueCodeVerification, err := flattenPlainValue(sourceCodeVerification, types.ObjectType{AttrTypes: map[string]attr.Type{"attempts_remaining": types.Int64Type, "status": types.StringType}}, "code_verification", "raw response"); err != nil {
					return err
				} else {
					if typedCodeVerification, ok := valueCodeVerification.(types.Object); ok {
						state.CodeVerification = typedCodeVerification
						assignedCodeVerification = true
					}
				}
			}
		}
		if !assignedCodeVerification {
			if !hasRaw {
				if responseValueCodeVerification, ok := plainFromResponseField(obj, "CodeVerification"); ok {
					sourceCodeVerification := applyConfiguredKeyedListShapes(responseValueCodeVerification, attrValueToPlain(state.CodeVerification))
					if valueCodeVerification, err := flattenPlainValue(
						sourceCodeVerification,
						types.ObjectType{AttrTypes: map[string]attr.Type{"attempts_remaining": types.Int64Type, "status": types.StringType}},
						"code_verification",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCodeVerification, ok := valueCodeVerification.(types.Object); ok {
							state.CodeVerification = typedCodeVerification
							assignedCodeVerification = true
						}
					}
				}
			}
		}
		if !assignedCodeVerification && hadRawCodeVerification {
			if nullCodeVerification, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"attempts_remaining": types.Int64Type, "status": types.StringType}}); ok {
				if typedCodeVerification, ok := nullCodeVerification.(types.Object); ok {
					state.CodeVerification = typedCodeVerification
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
		if rawValueCustomer, rawOk := plainValueAtPath(raw, "customer"); rawOk {
			if valueCustomer, err := flattenPlainValue(rawValueCustomer, types.StringType, "customer", "raw response"); err != nil {
				return err
			} else {
				if typedCustomer, ok := valueCustomer.(types.String); ok {
					state.Customer = typedCustomer
				}
			}
		} else if !hasRaw {
			if responseValueCustomer, ok := plainFromResponseField(obj, "Customer"); ok {
				if valueCustomer, err := flattenPlainValue(responseValueCustomer, types.StringType, "customer", "response struct"); err != nil {
					return err
				} else {
					if typedCustomer, ok := valueCustomer.(types.String); ok {
						state.Customer = typedCustomer
					}
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
				if valueEPS, err := flattenPlainValue(sourceEPS, types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType, "statement_descriptor": types.StringType}}, "eps", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType, "statement_descriptor": types.StringType}},
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
			if nullEPS, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType, "statement_descriptor": types.StringType}}); ok {
				if typedEPS, ok := nullEPS.(types.Object); ok {
					state.EPS = typedEPS
				}
			}
		}
	}
	{
		if rawValueFlow, rawOk := plainValueAtPath(raw, "flow"); rawOk {
			if valueFlow, err := flattenPlainValue(rawValueFlow, types.StringType, "flow", "raw response"); err != nil {
				return err
			} else {
				if typedFlow, ok := valueFlow.(types.String); ok {
					state.Flow = typedFlow
				}
			}
		} else if !hasRaw {
			if responseValueFlow, ok := plainFromResponseField(obj, "Flow"); ok {
				if valueFlow, err := flattenPlainValue(responseValueFlow, types.StringType, "flow", "response struct"); err != nil {
					return err
				} else {
					if typedFlow, ok := valueFlow.(types.String); ok {
						state.Flow = typedFlow
					}
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
				if valueGiropay, err := flattenPlainValue(sourceGiropay, types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "statement_descriptor": types.StringType}}, "giropay", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "statement_descriptor": types.StringType}},
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
			if nullGiropay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "statement_descriptor": types.StringType}}); ok {
				if typedGiropay, ok := nullGiropay.(types.Object); ok {
					state.Giropay = typedGiropay
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
				if valueIDEAL, err := flattenPlainValue(sourceIDEAL, types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "bic": types.StringType, "iban_last4": types.StringType, "statement_descriptor": types.StringType}}, "ideal", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "bic": types.StringType, "iban_last4": types.StringType, "statement_descriptor": types.StringType}},
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
			if nullIDEAL, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank": types.StringType, "bic": types.StringType, "iban_last4": types.StringType, "statement_descriptor": types.StringType}}); ok {
				if typedIDEAL, ok := nullIDEAL.(types.Object); ok {
					state.IDEAL = typedIDEAL
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
				if valueKlarna, err := flattenPlainValue(sourceKlarna, types.ObjectType{AttrTypes: map[string]attr.Type{"background_image_url": types.StringType, "client_token": types.StringType, "first_name": types.StringType, "last_name": types.StringType, "locale": types.StringType, "logo_url": types.StringType, "page_title": types.StringType, "pay_later_asset_urls_descriptive": types.StringType, "pay_later_asset_urls_standard": types.StringType, "pay_later_name": types.StringType, "pay_later_redirect_url": types.StringType, "pay_now_asset_urls_descriptive": types.StringType, "pay_now_asset_urls_standard": types.StringType, "pay_now_name": types.StringType, "pay_now_redirect_url": types.StringType, "pay_over_time_asset_urls_descriptive": types.StringType, "pay_over_time_asset_urls_standard": types.StringType, "pay_over_time_name": types.StringType, "pay_over_time_redirect_url": types.StringType, "payment_method_categories": types.StringType, "purchase_country": types.StringType, "purchase_type": types.StringType, "redirect_url": types.StringType, "shipping_delay": types.Int64Type, "shipping_first_name": types.StringType, "shipping_last_name": types.StringType}}, "klarna", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"background_image_url": types.StringType, "client_token": types.StringType, "first_name": types.StringType, "last_name": types.StringType, "locale": types.StringType, "logo_url": types.StringType, "page_title": types.StringType, "pay_later_asset_urls_descriptive": types.StringType, "pay_later_asset_urls_standard": types.StringType, "pay_later_name": types.StringType, "pay_later_redirect_url": types.StringType, "pay_now_asset_urls_descriptive": types.StringType, "pay_now_asset_urls_standard": types.StringType, "pay_now_name": types.StringType, "pay_now_redirect_url": types.StringType, "pay_over_time_asset_urls_descriptive": types.StringType, "pay_over_time_asset_urls_standard": types.StringType, "pay_over_time_name": types.StringType, "pay_over_time_redirect_url": types.StringType, "payment_method_categories": types.StringType, "purchase_country": types.StringType, "purchase_type": types.StringType, "redirect_url": types.StringType, "shipping_delay": types.Int64Type, "shipping_first_name": types.StringType, "shipping_last_name": types.StringType}},
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
			if nullKlarna, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"background_image_url": types.StringType, "client_token": types.StringType, "first_name": types.StringType, "last_name": types.StringType, "locale": types.StringType, "logo_url": types.StringType, "page_title": types.StringType, "pay_later_asset_urls_descriptive": types.StringType, "pay_later_asset_urls_standard": types.StringType, "pay_later_name": types.StringType, "pay_later_redirect_url": types.StringType, "pay_now_asset_urls_descriptive": types.StringType, "pay_now_asset_urls_standard": types.StringType, "pay_now_name": types.StringType, "pay_now_redirect_url": types.StringType, "pay_over_time_asset_urls_descriptive": types.StringType, "pay_over_time_asset_urls_standard": types.StringType, "pay_over_time_name": types.StringType, "pay_over_time_redirect_url": types.StringType, "payment_method_categories": types.StringType, "purchase_country": types.StringType, "purchase_type": types.StringType, "redirect_url": types.StringType, "shipping_delay": types.Int64Type, "shipping_first_name": types.StringType, "shipping_last_name": types.StringType}}); ok {
				if typedKlarna, ok := nullKlarna.(types.Object); ok {
					state.Klarna = typedKlarna
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
		assignedMultibanco := false
		hadRawMultibanco := false
		if rawValueMultibanco, rawOk := plainValueAtPath(raw, "multibanco"); rawOk {
			hadRawMultibanco = true
			if rawValueMultibanco != nil {
				sourceMultibanco := applyConfiguredKeyedListShapes(rawValueMultibanco, attrValueToPlain(state.Multibanco))
				if valueMultibanco, err := flattenPlainValue(sourceMultibanco, types.ObjectType{AttrTypes: map[string]attr.Type{"entity": types.StringType, "reference": types.StringType, "refund_account_holder_address_city": types.StringType, "refund_account_holder_address_country": types.StringType, "refund_account_holder_address_line1": types.StringType, "refund_account_holder_address_line2": types.StringType, "refund_account_holder_address_postal_code": types.StringType, "refund_account_holder_address_state": types.StringType, "refund_account_holder_name": types.StringType, "refund_iban": types.StringType}}, "multibanco", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"entity": types.StringType, "reference": types.StringType, "refund_account_holder_address_city": types.StringType, "refund_account_holder_address_country": types.StringType, "refund_account_holder_address_line1": types.StringType, "refund_account_holder_address_line2": types.StringType, "refund_account_holder_address_postal_code": types.StringType, "refund_account_holder_address_state": types.StringType, "refund_account_holder_name": types.StringType, "refund_iban": types.StringType}},
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
			if nullMultibanco, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"entity": types.StringType, "reference": types.StringType, "refund_account_holder_address_city": types.StringType, "refund_account_holder_address_country": types.StringType, "refund_account_holder_address_line1": types.StringType, "refund_account_holder_address_line2": types.StringType, "refund_account_holder_address_postal_code": types.StringType, "refund_account_holder_address_state": types.StringType, "refund_account_holder_name": types.StringType, "refund_iban": types.StringType}}); ok {
				if typedMultibanco, ok := nullMultibanco.(types.Object); ok {
					state.Multibanco = typedMultibanco
				}
			}
		}
	}
	{
		assignedOwner := false
		hadRawOwner := false
		if rawValueOwner, rawOk := plainValueAtPath(raw, "owner"); rawOk {
			hadRawOwner = true
			if rawValueOwner != nil {
				sourceOwner := applyConfiguredKeyedListShapes(rawValueOwner, attrValueToPlain(state.Owner))
				if valueOwner, err := flattenPlainValue(sourceOwner, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "phone": types.StringType, "verified_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "verified_email": types.StringType, "verified_name": types.StringType, "verified_phone": types.StringType}}, "owner", "raw response"); err != nil {
					return err
				} else {
					if typedOwner, ok := valueOwner.(types.Object); ok {
						state.Owner = typedOwner
						assignedOwner = true
					}
				}
			}
		}
		if !assignedOwner {
			if !hasRaw {
				if responseValueOwner, ok := plainFromResponseField(obj, "Owner"); ok {
					sourceOwner := applyConfiguredKeyedListShapes(responseValueOwner, attrValueToPlain(state.Owner))
					if valueOwner, err := flattenPlainValue(
						sourceOwner,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "phone": types.StringType, "verified_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "verified_email": types.StringType, "verified_name": types.StringType, "verified_phone": types.StringType}},
						"owner",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedOwner, ok := valueOwner.(types.Object); ok {
							state.Owner = typedOwner
							assignedOwner = true
						}
					}
				}
			}
		}
		if !assignedOwner && hadRawOwner {
			if nullOwner, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "email": types.StringType, "name": types.StringType, "phone": types.StringType, "verified_address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "verified_email": types.StringType, "verified_name": types.StringType, "verified_phone": types.StringType}}); ok {
				if typedOwner, ok := nullOwner.(types.Object); ok {
					state.Owner = typedOwner
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
				if valueP24, err := flattenPlainValue(sourceP24, types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType}}, "p24", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType}},
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
			if nullP24, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"reference": types.StringType}}); ok {
				if typedP24, ok := nullP24.(types.Object); ok {
					state.P24 = typedP24
				}
			}
		}
	}
	{
		assignedReceiver := false
		hadRawReceiver := false
		if rawValueReceiver, rawOk := plainValueAtPath(raw, "receiver"); rawOk {
			hadRawReceiver = true
			if rawValueReceiver != nil {
				sourceReceiver := applyConfiguredKeyedListShapes(rawValueReceiver, attrValueToPlain(state.Receiver))
				if valueReceiver, err := flattenPlainValue(sourceReceiver, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.StringType, "amount_charged": types.Int64Type, "amount_received": types.Int64Type, "amount_returned": types.Int64Type, "refund_attributes_method": types.StringType, "refund_attributes_status": types.StringType}}, "receiver", "raw response"); err != nil {
					return err
				} else {
					if typedReceiver, ok := valueReceiver.(types.Object); ok {
						state.Receiver = typedReceiver
						assignedReceiver = true
					}
				}
			}
		}
		if !assignedReceiver {
			if !hasRaw {
				if responseValueReceiver, ok := plainFromResponseField(obj, "Receiver"); ok {
					sourceReceiver := applyConfiguredKeyedListShapes(responseValueReceiver, attrValueToPlain(state.Receiver))
					if valueReceiver, err := flattenPlainValue(
						sourceReceiver,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.StringType, "amount_charged": types.Int64Type, "amount_received": types.Int64Type, "amount_returned": types.Int64Type, "refund_attributes_method": types.StringType, "refund_attributes_status": types.StringType}},
						"receiver",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedReceiver, ok := valueReceiver.(types.Object); ok {
							state.Receiver = typedReceiver
							assignedReceiver = true
						}
					}
				}
			}
		}
		if !assignedReceiver && hadRawReceiver {
			if nullReceiver, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.StringType, "amount_charged": types.Int64Type, "amount_received": types.Int64Type, "amount_returned": types.Int64Type, "refund_attributes_method": types.StringType, "refund_attributes_status": types.StringType}}); ok {
				if typedReceiver, ok := nullReceiver.(types.Object); ok {
					state.Receiver = typedReceiver
				}
			}
		}
	}
	{
		assignedRedirect := false
		hadRawRedirect := false
		if rawValueRedirect, rawOk := plainValueAtPath(raw, "redirect"); rawOk {
			hadRawRedirect = true
			if rawValueRedirect != nil {
				sourceRedirect := applyConfiguredKeyedListShapes(rawValueRedirect, attrValueToPlain(state.Redirect))
				if valueRedirect, err := flattenPlainValue(sourceRedirect, types.ObjectType{AttrTypes: map[string]attr.Type{"failure_reason": types.StringType, "return_url": types.StringType, "status": types.StringType, "url": types.StringType}}, "redirect", "raw response"); err != nil {
					return err
				} else {
					if typedRedirect, ok := valueRedirect.(types.Object); ok {
						state.Redirect = typedRedirect
						assignedRedirect = true
					}
				}
			}
		}
		if !assignedRedirect {
			if !hasRaw {
				if responseValueRedirect, ok := plainFromResponseField(obj, "Redirect"); ok {
					sourceRedirect := applyConfiguredKeyedListShapes(responseValueRedirect, attrValueToPlain(state.Redirect))
					if valueRedirect, err := flattenPlainValue(
						sourceRedirect,
						types.ObjectType{AttrTypes: map[string]attr.Type{"failure_reason": types.StringType, "return_url": types.StringType, "status": types.StringType, "url": types.StringType}},
						"redirect",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRedirect, ok := valueRedirect.(types.Object); ok {
							state.Redirect = typedRedirect
							assignedRedirect = true
						}
					}
				}
			}
		}
		if !assignedRedirect && hadRawRedirect {
			if nullRedirect, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"failure_reason": types.StringType, "return_url": types.StringType, "status": types.StringType, "url": types.StringType}}); ok {
				if typedRedirect, ok := nullRedirect.(types.Object); ok {
					state.Redirect = typedRedirect
				}
			}
		}
	}
	{
		assignedSEPACreditTransfer := false
		hadRawSEPACreditTransfer := false
		if rawValueSEPACreditTransfer, rawOk := plainValueAtPath(raw, "sepa_credit_transfer"); rawOk {
			hadRawSEPACreditTransfer = true
			if rawValueSEPACreditTransfer != nil {
				sourceSEPACreditTransfer := applyConfiguredKeyedListShapes(rawValueSEPACreditTransfer, attrValueToPlain(state.SEPACreditTransfer))
				if valueSEPACreditTransfer, err := flattenPlainValue(sourceSEPACreditTransfer, types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "bic": types.StringType, "iban": types.StringType, "refund_account_holder_address_city": types.StringType, "refund_account_holder_address_country": types.StringType, "refund_account_holder_address_line1": types.StringType, "refund_account_holder_address_line2": types.StringType, "refund_account_holder_address_postal_code": types.StringType, "refund_account_holder_address_state": types.StringType, "refund_account_holder_name": types.StringType, "refund_iban": types.StringType}}, "sepa_credit_transfer", "raw response"); err != nil {
					return err
				} else {
					if typedSEPACreditTransfer, ok := valueSEPACreditTransfer.(types.Object); ok {
						state.SEPACreditTransfer = typedSEPACreditTransfer
						assignedSEPACreditTransfer = true
					}
				}
			}
		}
		if !assignedSEPACreditTransfer {
			if !hasRaw {
				if responseValueSEPACreditTransfer, ok := plainFromResponseField(obj, "SEPACreditTransfer"); ok {
					sourceSEPACreditTransfer := applyConfiguredKeyedListShapes(responseValueSEPACreditTransfer, attrValueToPlain(state.SEPACreditTransfer))
					if valueSEPACreditTransfer, err := flattenPlainValue(
						sourceSEPACreditTransfer,
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "bic": types.StringType, "iban": types.StringType, "refund_account_holder_address_city": types.StringType, "refund_account_holder_address_country": types.StringType, "refund_account_holder_address_line1": types.StringType, "refund_account_holder_address_line2": types.StringType, "refund_account_holder_address_postal_code": types.StringType, "refund_account_holder_address_state": types.StringType, "refund_account_holder_name": types.StringType, "refund_iban": types.StringType}},
						"sepa_credit_transfer",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSEPACreditTransfer, ok := valueSEPACreditTransfer.(types.Object); ok {
							state.SEPACreditTransfer = typedSEPACreditTransfer
							assignedSEPACreditTransfer = true
						}
					}
				}
			}
		}
		if !assignedSEPACreditTransfer && hadRawSEPACreditTransfer {
			if nullSEPACreditTransfer, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank_name": types.StringType, "bic": types.StringType, "iban": types.StringType, "refund_account_holder_address_city": types.StringType, "refund_account_holder_address_country": types.StringType, "refund_account_holder_address_line1": types.StringType, "refund_account_holder_address_line2": types.StringType, "refund_account_holder_address_postal_code": types.StringType, "refund_account_holder_address_state": types.StringType, "refund_account_holder_name": types.StringType, "refund_iban": types.StringType}}); ok {
				if typedSEPACreditTransfer, ok := nullSEPACreditTransfer.(types.Object); ok {
					state.SEPACreditTransfer = typedSEPACreditTransfer
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
				if valueSEPADebit, err := flattenPlainValue(sourceSEPADebit, types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "branch_code": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate_reference": types.StringType, "mandate_url": types.StringType}}, "sepa_debit", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "branch_code": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate_reference": types.StringType, "mandate_url": types.StringType}},
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
			if nullSEPADebit, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "branch_code": types.StringType, "country": types.StringType, "fingerprint": types.StringType, "last4": types.StringType, "mandate_reference": types.StringType, "mandate_url": types.StringType}}); ok {
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
				if valueSofort, err := flattenPlainValue(sourceSofort, types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "country": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "statement_descriptor": types.StringType}}, "sofort", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "country": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "statement_descriptor": types.StringType}},
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
			if nullSofort, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"bank_code": types.StringType, "bank_name": types.StringType, "bic": types.StringType, "country": types.StringType, "iban_last4": types.StringType, "preferred_language": types.StringType, "statement_descriptor": types.StringType}}); ok {
				if typedSofort, ok := nullSofort.(types.Object); ok {
					state.Sofort = typedSofort
				}
			}
		}
	}
	{
		assignedSourceOrder := false
		hadRawSourceOrder := false
		if rawValueSourceOrder, rawOk := plainValueAtPath(raw, "source_order"); rawOk {
			hadRawSourceOrder = true
			if rawValueSourceOrder != nil {
				sourceSourceOrder := applyConfiguredKeyedListShapes(rawValueSourceOrder, attrValueToPlain(state.SourceOrder))
				if valueSourceOrder, err := flattenPlainValue(sourceSourceOrder, types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "email": types.StringType, "items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "description": types.StringType, "parent": types.StringType, "quantity": types.Int64Type, "type": types.StringType}}}, "shipping": types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}}}, "source_order", "raw response"); err != nil {
					return err
				} else {
					if typedSourceOrder, ok := valueSourceOrder.(types.Object); ok {
						state.SourceOrder = typedSourceOrder
						assignedSourceOrder = true
					}
				}
			}
		}
		if !assignedSourceOrder {
			if !hasRaw {
				if responseValueSourceOrder, ok := plainFromResponseField(obj, "SourceOrder"); ok {
					sourceSourceOrder := applyConfiguredKeyedListShapes(responseValueSourceOrder, attrValueToPlain(state.SourceOrder))
					if valueSourceOrder, err := flattenPlainValue(
						sourceSourceOrder,
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "email": types.StringType, "items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "description": types.StringType, "parent": types.StringType, "quantity": types.Int64Type, "type": types.StringType}}}, "shipping": types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}}},
						"source_order",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSourceOrder, ok := valueSourceOrder.(types.Object); ok {
							state.SourceOrder = typedSourceOrder
							assignedSourceOrder = true
						}
					}
				}
			}
		}
		if !assignedSourceOrder && hadRawSourceOrder {
			if nullSourceOrder, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "email": types.StringType, "items": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType, "description": types.StringType, "parent": types.StringType, "quantity": types.Int64Type, "type": types.StringType}}}, "shipping": types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "carrier": types.StringType, "name": types.StringType, "phone": types.StringType, "tracking_number": types.StringType}}}}); ok {
				if typedSourceOrder, ok := nullSourceOrder.(types.Object); ok {
					state.SourceOrder = typedSourceOrder
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
		assignedThreeDSecure := false
		hadRawThreeDSecure := false
		if rawValueThreeDSecure, rawOk := plainValueAtPath(raw, "three_d_secure"); rawOk {
			hadRawThreeDSecure = true
			if rawValueThreeDSecure != nil {
				sourceThreeDSecure := applyConfiguredKeyedListShapes(rawValueThreeDSecure, attrValueToPlain(state.ThreeDSecure))
				if valueThreeDSecure, err := flattenPlainValue(sourceThreeDSecure, types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_zip_check": types.StringType, "authenticated": types.BoolType, "brand": types.StringType, "card": types.StringType, "country": types.StringType, "customer": types.StringType, "cvc_check": types.StringType, "description": types.StringType, "dynamic_last4": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "name": types.StringType, "three_d_secure": types.StringType, "tokenization_method": types.StringType}}, "three_d_secure", "raw response"); err != nil {
					return err
				} else {
					if typedThreeDSecure, ok := valueThreeDSecure.(types.Object); ok {
						state.ThreeDSecure = typedThreeDSecure
						assignedThreeDSecure = true
					}
				}
			}
		}
		if !assignedThreeDSecure {
			if !hasRaw {
				if responseValueThreeDSecure, ok := plainFromResponseField(obj, "ThreeDSecure"); ok {
					sourceThreeDSecure := applyConfiguredKeyedListShapes(responseValueThreeDSecure, attrValueToPlain(state.ThreeDSecure))
					if valueThreeDSecure, err := flattenPlainValue(
						sourceThreeDSecure,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_zip_check": types.StringType, "authenticated": types.BoolType, "brand": types.StringType, "card": types.StringType, "country": types.StringType, "customer": types.StringType, "cvc_check": types.StringType, "description": types.StringType, "dynamic_last4": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "name": types.StringType, "three_d_secure": types.StringType, "tokenization_method": types.StringType}},
						"three_d_secure",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedThreeDSecure, ok := valueThreeDSecure.(types.Object); ok {
							state.ThreeDSecure = typedThreeDSecure
							assignedThreeDSecure = true
						}
					}
				}
			}
		}
		if !assignedThreeDSecure && hadRawThreeDSecure {
			if nullThreeDSecure, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address_line1_check": types.StringType, "address_zip_check": types.StringType, "authenticated": types.BoolType, "brand": types.StringType, "card": types.StringType, "country": types.StringType, "customer": types.StringType, "cvc_check": types.StringType, "description": types.StringType, "dynamic_last4": types.StringType, "exp_month": types.Int64Type, "exp_year": types.Int64Type, "fingerprint": types.StringType, "funding": types.StringType, "iin": types.StringType, "issuer": types.StringType, "last4": types.StringType, "name": types.StringType, "three_d_secure": types.StringType, "tokenization_method": types.StringType}}); ok {
				if typedThreeDSecure, ok := nullThreeDSecure.(types.Object); ok {
					state.ThreeDSecure = typedThreeDSecure
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
		if rawValueUsage, rawOk := plainValueAtPath(raw, "usage"); rawOk {
			if valueUsage, err := flattenPlainValue(rawValueUsage, types.StringType, "usage", "raw response"); err != nil {
				return err
			} else {
				if typedUsage, ok := valueUsage.(types.String); ok {
					state.Usage = typedUsage
				}
			}
		} else if !hasRaw {
			if responseValueUsage, ok := plainFromResponseField(obj, "Usage"); ok {
				if valueUsage, err := flattenPlainValue(responseValueUsage, types.StringType, "usage", "response struct"); err != nil {
					return err
				} else {
					if typedUsage, ok := valueUsage.(types.String); ok {
						state.Usage = typedUsage
					}
				}
			}
		}
	}
	{
		assignedWeChat := false
		hadRawWeChat := false
		if rawValueWeChat, rawOk := plainValueAtPath(raw, "wechat"); rawOk {
			hadRawWeChat = true
			if rawValueWeChat != nil {
				sourceWeChat := applyConfiguredKeyedListShapes(rawValueWeChat, attrValueToPlain(state.WeChat))
				if valueWeChat, err := flattenPlainValue(sourceWeChat, types.ObjectType{AttrTypes: map[string]attr.Type{"prepay_id": types.StringType, "qr_code_url": types.StringType, "statement_descriptor": types.StringType}}, "wechat", "raw response"); err != nil {
					return err
				} else {
					if typedWeChat, ok := valueWeChat.(types.Object); ok {
						state.WeChat = typedWeChat
						assignedWeChat = true
					}
				}
			}
		}
		if !assignedWeChat {
			if !hasRaw {
				if responseValueWeChat, ok := plainFromResponseField(obj, "WeChat"); ok {
					sourceWeChat := applyConfiguredKeyedListShapes(responseValueWeChat, attrValueToPlain(state.WeChat))
					if valueWeChat, err := flattenPlainValue(
						sourceWeChat,
						types.ObjectType{AttrTypes: map[string]attr.Type{"prepay_id": types.StringType, "qr_code_url": types.StringType, "statement_descriptor": types.StringType}},
						"wechat",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedWeChat, ok := valueWeChat.(types.Object); ok {
							state.WeChat = typedWeChat
							assignedWeChat = true
						}
					}
				}
			}
		}
		if !assignedWeChat && hadRawWeChat {
			if nullWeChat, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"prepay_id": types.StringType, "qr_code_url": types.StringType, "statement_descriptor": types.StringType}}); ok {
				if typedWeChat, ok := nullWeChat.(types.Object); ok {
					state.WeChat = typedWeChat
				}
			}
		}
	}
	return nil
}
