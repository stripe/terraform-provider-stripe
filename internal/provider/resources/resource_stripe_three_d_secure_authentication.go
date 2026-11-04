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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &ThreeDSecureAuthenticationResource{}

var _ resource.ResourceWithConfigure = &ThreeDSecureAuthenticationResource{}

var _ resource.ResourceWithImportState = &ThreeDSecureAuthenticationResource{}

func NewThreeDSecureAuthenticationResource() resource.Resource {
	return &ThreeDSecureAuthenticationResource{}
}

type ThreeDSecureAuthenticationResource struct {
	client *stripe.Client
}

type ThreeDSecureAuthenticationResourceModel struct {
	Object            types.String `tfsdk:"object"`
	AcquirerDetails   types.Object `tfsdk:"acquirer_details"`
	Amount            types.Int64  `tfsdk:"amount"`
	ChallengeURL      types.String `tfsdk:"challenge_url"`
	Channel           types.Object `tfsdk:"channel"`
	Created           types.Int64  `tfsdk:"created"`
	Currency          types.String `tfsdk:"currency"`
	DirectoryServer   types.String `tfsdk:"directory_server"`
	FingerprintingURL types.String `tfsdk:"fingerprinting_url"`
	FlowPreference    types.Object `tfsdk:"flow_preference"`
	FutureUsage       types.Object `tfsdk:"future_usage"`
	ID                types.String `tfsdk:"id"`
	Livemode          types.Bool   `tfsdk:"livemode"`
	MessageCategory   types.String `tfsdk:"message_category"`
	Metadata          types.Map    `tfsdk:"metadata"`
	Outcome           types.String `tfsdk:"outcome"`
	OutcomeDetails    types.Object `tfsdk:"outcome_details"`
	PaymentMethod     types.String `tfsdk:"payment_method"`
	Reason            types.String `tfsdk:"reason"`
	ShippingAddress   types.Object `tfsdk:"shipping_address"`
	Status            types.String `tfsdk:"status"`
	PaymentMethodData types.Object `tfsdk:"payment_method_data"`
	Submit            types.String `tfsdk:"submit"`
}

func (r *ThreeDSecureAuthenticationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ThreeDSecureAuthenticationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_three_d_secure_authentication"
}

func (r *ThreeDSecureAuthenticationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The Standalone 3DS API allows you to run EMV 3D Secure (3DS) authentication using Stripe while authorizing the payment with any PSP.\n\nRelated guide: [Standalone 3DS](/payments/3d-secure/standalone-3d-secure)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("three_d_secure.authentication")},
			},
			"acquirer_details": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Contains additional details about the acquirer for a 3DS Authentication.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"acquirer_bin": schema.StringAttribute{
						Required:      true,
						Description:   "The Acquirer BIN (specific to the directory_server).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"acquirer_country": schema.StringAttribute{
						Required:      true,
						Description:   "The two-letter country code of the acquirer ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"acquirer_merchant_id": schema.StringAttribute{
						Required:      true,
						Description:   "The Merchant ID (or Card Acceptor ID) that your acquirer assigned you (specific to the directory_server).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"mcc": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) as defined by each payment system or directory server.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"merchant_name": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The merchant name assigned by the acquirer or payment system. Same name used in the authorization message as defined in [ISO 8583](https://en.wikipedia.org/wiki/ISO_8583).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"requestor_id": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Requestor ID if you’re enrolled in the card network’s 3DS program. Otherwise, you can omit this field because Stripe assigns a Requestor ID with the card networks.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The amount for this 3DS Authentication.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"challenge_url": schema.StringAttribute{
				Computed:      true,
				Description:   "The URL for presenting a challenge to your cardholder, present if status is requires_challenge.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"channel": schema.SingleNestedAttribute{
				Required:      true,
				Description:   "Contains details on the channel used (browser, 3RI) for a standalone 3DS Authentication.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"browser": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Contains details on the browser for a standalone 3DS Authentication.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"accept_header": schema.StringAttribute{
								Required:      true,
								Description:   "The HTTP accept headers from the cardholder's browser.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
							"color_depth": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The color depth of the cardholder’s screen.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
							},
							"ip_address": schema.StringAttribute{
								Required:      true,
								Description:   "The IP address of the browser.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
							"java_enabled": schema.BoolAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The cardholder browser’s ability to execute Java.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown(), boolplanmodifier.RequiresReplace()},
							},
							"javascript_enabled": schema.BoolAttribute{
								Required:      true,
								Description:   "The cardholder browser’s ability to execute JavaScript.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
							},
							"language": schema.StringAttribute{
								Required:      true,
								Description:   "An IETF BCP 47 language tag representing the browser language.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
							"screen_height": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The total height of the cardholder’s screen in pixels.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
							},
							"screen_width": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The total width of the cardholder’s screen in pixels.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
							},
							"timezone_offset": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The time difference between UTC time and the local time of the cardholder’s browser, in minutes.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
							},
							"user_agent": schema.StringAttribute{
								Required:      true,
								Description:   "The browser user agent.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
							"device_id": schema.StringAttribute{
								Optional:      true,
								Description:   "Unique and immutable identifier linked to a device that is consistent across 3DS transactions for the specific user device. For example: hardware device ID or a platform-calculated device fingerprint.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
						},
					},
					"three_r_i": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Contains details for a 3RI standalone 3DS Authentication.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"previous_authentication": schema.StringAttribute{
								Required:      true,
								Description:   "ID of the previous initial authenticated 3DS Authentication object.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of the 3RI Authentication.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("delayed_shipment", "other_payment", "recurring", "split_shipment")},
							},
						},
					},
					"type": schema.StringAttribute{
						Required:      true,
						Description:   "Type of channel you would prefer to use for this 3DS Authentication. Only browser.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("browser", "three_r_i")},
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
			"directory_server": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The 3DS directory server with which this 3DS Authentication was processed.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("american_express", "cartes_bancaires", "discover", "mastercard", "visa")},
			},
			"fingerprinting_url": schema.StringAttribute{
				Computed:      true,
				Description:   "The URL for performing issuer fingerprinting, present if fingerprinting is supported for the given payment method.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"flow_preference": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Contains details of the flow preference used for a standalone 3DS Authentication.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"challenge": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of challenge flow you requested for this 3DS Authentication.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("mandated", "preferred")},
							},
						},
					},
					"data_share": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of data share flow you requested for this 3DS Authentication.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ds_specific", "emv_standard")},
							},
						},
					},
					"frictionless": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of frictionless flow you requested for this 3DS Authentication.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("low_risk", "none")},
							},
						},
					},
					"type": schema.StringAttribute{
						Required:      true,
						Description:   "Type of flow you requested for this 3DS Authentication.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("challenge", "data_share", "frictionless")},
					},
				},
			},
			"future_usage": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Contains information about the future authorisations related to this authentication",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"installment": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Details about installment payments",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"amount": schema.Int64Attribute{
								Required:      true,
								Description:   "A non-negative integer representing the amount in the [smallest currency unit](/currencies#zero-decimal).",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
							"expiry": schema.SingleNestedAttribute{
								Required:      true,
								Description:   "Information about recurring payment expiry",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"date": schema.StringAttribute{
										Optional: true,
										Computed: true,

										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
									},
									"type": schema.StringAttribute{
										Required: true,

										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("date", "never")},
									},
								},
							},
							"interval": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The unit of time for `interval_count`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("day")},
							},
							"interval_count": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The minimum number of time intervals between authorizations.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
							},
							"number": schema.Int64Attribute{
								Required:      true,
								Description:   "The maximum number of installments.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
						},
					},
					"recurring": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Details about recurring payments",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"amount": schema.Int64Attribute{
								Required:      true,
								Description:   "A non-negative integer representing the amount in the [smallest currency unit](/currencies#zero-decimal).",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
							"expiry": schema.SingleNestedAttribute{
								Required:      true,
								Description:   "Information about recurring payment expiry",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"date": schema.StringAttribute{
										Optional: true,
										Computed: true,

										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
									},
									"type": schema.StringAttribute{
										Required: true,

										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("date", "never")},
									},
								},
							},
							"interval": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The unit of time for `interval_count`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("day")},
							},
							"interval_count": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The minimum number of time intervals between authorizations.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
							},
						},
					},
					"type": schema.StringAttribute{
						Required:      true,
						Description:   "The type of future usage declared for this 3DS Authentication.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("card_on_file", "installment", "recurring")},
					},
				},
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
			"message_category": schema.StringAttribute{
				Required:      true,
				Description:   "Indicates whether this 3DS Authentication is being performed for a payment or non-payment use case.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("non_payment_authentication", "payment_authentication")},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown(), mapplanmodifier.RequiresReplace()},
				ElementType:   types.StringType,
			},
			"outcome": schema.StringAttribute{
				Computed:      true,
				Description:   "The outcome of this 3DS Authentication.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("abandoned", "attempt_acknowledged", "authenticated", "canceled", "denied", "informational", "internal_error", "not_supported", "not_triggered", "processing_error", "rejected")},
			},
			"outcome_details": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Contains details on the result for a standalone 3DS Authentication.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"acs_transaction_id": schema.StringAttribute{
						Computed:      true,
						Description:   "Universally unique transaction identifier assigned by the issuer to identify the transaction.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"ares": schema.StringAttribute{
						Computed:      true,
						Description:   "The Authentication Response Message (ARes) is the issuer's response to the AReq message.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"ares_trans_status": schema.StringAttribute{
						Computed:      true,
						Description:   "TransStatus field on the ARes",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("A", "C", "D", "I", "N", "R", "S", "U", "Y")},
					},
					"cryptogram": schema.StringAttribute{
						Computed:      true,
						Description:   "A 28-character Base64 string proving that 3DS was completed. Store this value securely, and don’t reuse it for multiple authorizations.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"ds_transaction_id": schema.StringAttribute{
						Computed:      true,
						Description:   "The 3DS2 Directory Server Transaction ID.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"eci": schema.StringAttribute{
						Computed:      true,
						Description:   "Electronic Commerce Indicator provided by the issuer to indicate the result of this 3DS Authentication.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"network_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains details specific to the individual network.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"cartes_bancaires": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "Contains details for Cartes Bancaires specific fields in the authentication outcomes.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"avalgo": schema.StringAttribute{
										Computed:      true,
										Description:   "The cryptogram calculation algorithm used by the card Issuer's ACS to calculate the Authentication cryptogram. Also known as cavvAlgorithm. ARes/RReq messageExtension: `CB-AVALGO`",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"cb_exemption": schema.StringAttribute{
										Computed:      true,
										Description:   "The exemption indicator returned from Cartes Bancaires in the ARes. This is a 3 byte bitmap (lowest significant byte first and most significant bit first) that has been Base64 encoded. String (4 characters). ARes message extension: `CB-EXEMPTION`",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"cb_score": schema.StringAttribute{
										Computed:      true,
										Description:   "The risk score returned from Cartes Bancaires in the ARes. Numeric value 0-99. ARes/RReq message extension: `CB-SCORE`",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
					"protocol_version": schema.StringAttribute{
						Computed:      true,
						Description:   "The 3DS protocol version used for this 3DS Authentication.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("2.1.0", "2.2.0", "2.3.1")},
					},
					"requestor_challenge_indicator": schema.StringAttribute{
						Computed:      true,
						Description:   "The indicator provided to the issuer by Stripe in the AReq that indicates whether a challenge is requested for this Authentication. This indicator should match the flow_preference you specified but may be overridden (for compliance reasons for example).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("01", "02", "03", "04", "05", "06")},
					},
					"rreq": schema.StringAttribute{
						Computed:      true,
						Description:   "The Results Request Message (RReq) communicates the results of the authentication or verification.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"rreq_trans_status": schema.StringAttribute{
						Computed:      true,
						Description:   "TransStatus field on the RReq",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("A", "C", "D", "I", "N", "R", "S", "U", "Y")},
					},
					"three_ds_server_transaction_id": schema.StringAttribute{
						Computed:      true,
						Description:   "Universally unique transaction identifier assigned by Stripe to identify the transaction.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"payment_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the payment method (a PaymentMethod object) to attach to this 3DS Authentication.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"reason": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The reason for invoking this 3DS Authentication.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("cardholder_authentication", "issuer_requested", "liability_shift", "processing_costs", "regulatory_compliance")},
			},
			"shipping_address": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Contains details about the shipping address for a 3DS Authentication.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"city": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "City, district, suburb, town, or village.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"country": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"line1": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Address line 1, such as the street, PO Box, or company name.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"line2": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Address line 2, such as the apartment, suite, unit, or building.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"postal_code": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "ZIP or postal code.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"state": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
				},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Status of this Authentication.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("canceled", "error", "failed", "requires_challenge", "requires_submission", "succeeded")},
			},
			"payment_method_data": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "Hash used to generate the PaymentMethod to be used for this Authentication. This is mutually exclusive with the `payment_method` parameter.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				WriteOnly:     true,
				Attributes: map[string]schema.Attribute{
					"billing_details": schema.SingleNestedAttribute{
						Optional:      true,
						Description:   "Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
						WriteOnly:     true,
						Attributes: map[string]schema.Attribute{
							"address": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Billing address.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								WriteOnly:     true,
								Attributes: map[string]schema.Attribute{
									"city": schema.StringAttribute{
										Optional:      true,
										Description:   "City, district, suburb, town, or village.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										WriteOnly:     true,
									},
									"country": schema.StringAttribute{
										Optional:      true,
										Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										WriteOnly:     true,
									},
									"line1": schema.StringAttribute{
										Optional:      true,
										Description:   "Address line 1, such as the street, PO Box, or company name.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										WriteOnly:     true,
									},
									"line2": schema.StringAttribute{
										Optional:      true,
										Description:   "Address line 2, such as the apartment, suite, unit, or building.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										WriteOnly:     true,
									},
									"postal_code": schema.StringAttribute{
										Optional:      true,
										Description:   "ZIP or postal code.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										WriteOnly:     true,
									},
									"state": schema.StringAttribute{
										Optional:      true,
										Description:   "Country subdivision code defined in ISO 3166-2.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										WriteOnly:     true,
									},
								},
							},
							"email": schema.StringAttribute{
								Optional:      true,
								Description:   "Email address.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								WriteOnly:     true,
							},
							"name": schema.StringAttribute{
								Optional:      true,
								Description:   "Full name.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								WriteOnly:     true,
							},
							"phone": schema.StringAttribute{
								Optional:      true,
								Description:   "Billing phone number (including extension).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								WriteOnly:     true,
							},
						},
					},
					"card": schema.SingleNestedAttribute{
						Required: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
						WriteOnly:     true,
						Attributes: map[string]schema.Attribute{
							"cvc": schema.StringAttribute{
								Optional: true,

								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								WriteOnly:     true,
							},
							"exp_month": schema.Int64Attribute{
								Optional: true,

								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
								WriteOnly:     true,
							},
							"exp_year": schema.Int64Attribute{
								Optional: true,

								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
								WriteOnly:     true,
							},
							"number": schema.StringAttribute{
								Optional: true,

								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								WriteOnly:     true,
							},
							"token": schema.StringAttribute{
								Optional: true,

								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								WriteOnly:     true,
							},
						},
					},
					"type": schema.StringAttribute{
						Required:      true,
						Description:   "The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
				},
			},
			"submit": schema.StringAttribute{
				Optional:      true,
				Description:   "Set to `always` to skip the fingerprinting step and submit this Authentication immediately or `if_fingerprinting_not_supported` to submit this Authentication only if fingerprinting is not available. This parameter defaults to `never`.\n\nRefer to the [Submit at creation section of the standalone 3DS guide](/payments/3d-secure/standalone-3d-secure#submit-at-creation) for more information.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
				Validators:    []validator.String{stringvalidator.OneOf("always", "if_fingerprinting_not_supported", "never")},
			},
		},
	}
}

func (r *ThreeDSecureAuthenticationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ThreeDSecureAuthenticationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config ThreeDSecureAuthenticationResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PaymentMethodData"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "card"}, []string{"PaymentMethodData", "card", "cvc"}, []string{"PaymentMethodData", "card", "exp_month"}, []string{"PaymentMethodData", "card", "exp_year"}, []string{"PaymentMethodData", "card", "number"}, []string{"PaymentMethodData", "card", "token"}, []string{"PaymentMethodData", "type"}, []string{"Submit"}})

	params, err := expandThreeDSecureAuthenticationCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building ThreeDSecureAuthentication create params", err.Error())
		return
	}

	obj, err := r.client.V1ThreeDSecureAuthentications.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating ThreeDSecureAuthentication", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ThreeDSecureAuthentications.B, r.client.V1ThreeDSecureAuthentications.Key, stripe.FormatURLPath("/v1/three_d_secure/authentications/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ThreeDSecureAuthentication create raw response", err.Error())
		return
	}

	if err := flattenThreeDSecureAuthentication(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening ThreeDSecureAuthentication create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Channel", "browser", "device_id"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"PaymentMethodData"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "card"}, []string{"PaymentMethodData", "card", "cvc"}, []string{"PaymentMethodData", "card", "exp_month"}, []string{"PaymentMethodData", "card", "exp_year"}, []string{"PaymentMethodData", "card", "number"}, []string{"PaymentMethodData", "card", "token"}, []string{"PaymentMethodData", "type"}, []string{"Submit"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ThreeDSecureAuthenticationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState ThreeDSecureAuthenticationResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ThreeDSecureAuthenticationResourceModel
	state = priorState

	obj, err := r.client.V1ThreeDSecureAuthentications.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading ThreeDSecureAuthentication", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ThreeDSecureAuthentications.B, r.client.V1ThreeDSecureAuthentications.Key, stripe.FormatURLPath("/v1/three_d_secure/authentications/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ThreeDSecureAuthentication raw response", err.Error())
		return
	}

	if err := flattenThreeDSecureAuthentication(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening ThreeDSecureAuthentication read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"Channel", "browser", "device_id"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"PaymentMethodData"}, []string{"PaymentMethodData", "billing_details"}, []string{"PaymentMethodData", "billing_details", "address"}, []string{"PaymentMethodData", "billing_details", "address", "city"}, []string{"PaymentMethodData", "billing_details", "address", "country"}, []string{"PaymentMethodData", "billing_details", "address", "line1"}, []string{"PaymentMethodData", "billing_details", "address", "line2"}, []string{"PaymentMethodData", "billing_details", "address", "postal_code"}, []string{"PaymentMethodData", "billing_details", "address", "state"}, []string{"PaymentMethodData", "billing_details", "email"}, []string{"PaymentMethodData", "billing_details", "name"}, []string{"PaymentMethodData", "billing_details", "phone"}, []string{"PaymentMethodData", "card"}, []string{"PaymentMethodData", "card", "cvc"}, []string{"PaymentMethodData", "card", "exp_month"}, []string{"PaymentMethodData", "card", "exp_year"}, []string{"PaymentMethodData", "card", "number"}, []string{"PaymentMethodData", "card", "token"}, []string{"PaymentMethodData", "type"}, []string{"Submit"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *ThreeDSecureAuthenticationResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "ThreeDSecureAuthentication does not support updates")
}

func (r *ThreeDSecureAuthenticationResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *ThreeDSecureAuthenticationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandThreeDSecureAuthenticationCreate(plan ThreeDSecureAuthenticationResourceModel) (*stripe.ThreeDSecureAuthenticationCreateParams, error) {
	params := &stripe.ThreeDSecureAuthenticationCreateParams{}

	if !plan.AcquirerDetails.IsNull() && !plan.AcquirerDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AcquirerDetails", plan.AcquirerDetails) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "acquirer_details", params)
		}
	}
	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Channel.IsNull() && !plan.Channel.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Channel", plan.Channel) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "channel", params)
		}
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.DirectoryServer.IsNull() && !plan.DirectoryServer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DirectoryServer", "DirectoryServer", plan.DirectoryServer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "directory_server", params)
		}
	}
	if !plan.FlowPreference.IsNull() && !plan.FlowPreference.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FlowPreference", plan.FlowPreference) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "flow_preference", params)
		}
	}
	if !plan.FutureUsage.IsNull() && !plan.FutureUsage.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FutureUsage", plan.FutureUsage) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "future_usage", params)
		}
	}
	if !plan.MessageCategory.IsNull() && !plan.MessageCategory.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "MessageCategory", "MessageCategory", plan.MessageCategory.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "message_category", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.PaymentMethod.IsNull() && !plan.PaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentMethodID", "PaymentMethod", plan.PaymentMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method", params)
		}
	}
	if !plan.Reason.IsNull() && !plan.Reason.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Reason", "Reason", plan.Reason.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "reason", params)
		}
	}
	if !plan.ShippingAddress.IsNull() && !plan.ShippingAddress.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingAddress", plan.ShippingAddress) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping_address", params)
		}
	}
	if !plan.PaymentMethodData.IsNull() && !plan.PaymentMethodData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PaymentMethodData", plan.PaymentMethodData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_method_data", params)
		}
	}
	if !plan.Submit.IsNull() && !plan.Submit.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Submit", "Submit", plan.Submit.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "submit", params)
		}
	}

	return params, nil
}

func flattenThreeDSecureAuthentication(obj *stripe.ThreeDSecureAuthentication, state *ThreeDSecureAuthenticationResourceModel) error {
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
		assignedAcquirerDetails := false
		hadRawAcquirerDetails := false
		if rawValueAcquirerDetails, rawOk := plainValueAtPath(raw, "acquirer_details"); rawOk {
			hadRawAcquirerDetails = true
			if rawValueAcquirerDetails != nil {
				sourceAcquirerDetails := applyConfiguredKeyedListShapes(rawValueAcquirerDetails, attrValueToPlain(state.AcquirerDetails))
				if valueAcquirerDetails, err := flattenPlainValue(sourceAcquirerDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"acquirer_bin": types.StringType, "acquirer_country": types.StringType, "acquirer_merchant_id": types.StringType, "mcc": types.StringType, "merchant_name": types.StringType, "requestor_id": types.StringType}}, "acquirer_details", "raw response"); err != nil {
					return err
				} else {
					if typedAcquirerDetails, ok := valueAcquirerDetails.(types.Object); ok {
						state.AcquirerDetails = typedAcquirerDetails
						assignedAcquirerDetails = true
					}
				}
			}
		}
		if !assignedAcquirerDetails {
			if !hasRaw {
				if responseValueAcquirerDetails, ok := plainFromResponseField(obj, "AcquirerDetails"); ok {
					sourceAcquirerDetails := applyConfiguredKeyedListShapes(responseValueAcquirerDetails, attrValueToPlain(state.AcquirerDetails))
					if valueAcquirerDetails, err := flattenPlainValue(
						sourceAcquirerDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"acquirer_bin": types.StringType, "acquirer_country": types.StringType, "acquirer_merchant_id": types.StringType, "mcc": types.StringType, "merchant_name": types.StringType, "requestor_id": types.StringType}},
						"acquirer_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAcquirerDetails, ok := valueAcquirerDetails.(types.Object); ok {
							state.AcquirerDetails = typedAcquirerDetails
							assignedAcquirerDetails = true
						}
					}
				}
			}
		}
		if !assignedAcquirerDetails && hadRawAcquirerDetails {
			if nullAcquirerDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"acquirer_bin": types.StringType, "acquirer_country": types.StringType, "acquirer_merchant_id": types.StringType, "mcc": types.StringType, "merchant_name": types.StringType, "requestor_id": types.StringType}}); ok {
				if typedAcquirerDetails, ok := nullAcquirerDetails.(types.Object); ok {
					state.AcquirerDetails = typedAcquirerDetails
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
		if rawValueChallengeURL, rawOk := plainValueAtPath(raw, "challenge_url"); rawOk {
			if valueChallengeURL, err := flattenPlainValue(rawValueChallengeURL, types.StringType, "challenge_url", "raw response"); err != nil {
				return err
			} else {
				if typedChallengeURL, ok := valueChallengeURL.(types.String); ok {
					state.ChallengeURL = typedChallengeURL
				}
			}
		} else if !hasRaw {
			if responseValueChallengeURL, ok := plainFromResponseField(obj, "ChallengeURL"); ok {
				if valueChallengeURL, err := flattenPlainValue(responseValueChallengeURL, types.StringType, "challenge_url", "response struct"); err != nil {
					return err
				} else {
					if typedChallengeURL, ok := valueChallengeURL.(types.String); ok {
						state.ChallengeURL = typedChallengeURL
					}
				}
			}
		}
	}
	{
		assignedChannel := false
		hadRawChannel := false
		if rawValueChannel, rawOk := plainValueAtPath(raw, "channel"); rawOk {
			hadRawChannel = true
			if rawValueChannel != nil {
				sourceChannel := applyConfiguredKeyedListShapes(rawValueChannel, attrValueToPlain(state.Channel))
				if valueChannel, err := flattenPlainValue(sourceChannel, types.ObjectType{AttrTypes: map[string]attr.Type{"browser": types.ObjectType{AttrTypes: map[string]attr.Type{"accept_header": types.StringType, "color_depth": types.Int64Type, "ip_address": types.StringType, "java_enabled": types.BoolType, "javascript_enabled": types.BoolType, "language": types.StringType, "screen_height": types.Int64Type, "screen_width": types.Int64Type, "timezone_offset": types.Int64Type, "user_agent": types.StringType, "device_id": types.StringType}}, "three_r_i": types.ObjectType{AttrTypes: map[string]attr.Type{"previous_authentication": types.StringType, "type": types.StringType}}, "type": types.StringType}}, "channel", "raw response"); err != nil {
					return err
				} else {
					if typedChannel, ok := valueChannel.(types.Object); ok {
						state.Channel = typedChannel
						assignedChannel = true
					}
				}
			}
		}
		if !assignedChannel {
			if !hasRaw {
				if responseValueChannel, ok := plainFromResponseField(obj, "Channel"); ok {
					sourceChannel := applyConfiguredKeyedListShapes(responseValueChannel, attrValueToPlain(state.Channel))
					if valueChannel, err := flattenPlainValue(
						sourceChannel,
						types.ObjectType{AttrTypes: map[string]attr.Type{"browser": types.ObjectType{AttrTypes: map[string]attr.Type{"accept_header": types.StringType, "color_depth": types.Int64Type, "ip_address": types.StringType, "java_enabled": types.BoolType, "javascript_enabled": types.BoolType, "language": types.StringType, "screen_height": types.Int64Type, "screen_width": types.Int64Type, "timezone_offset": types.Int64Type, "user_agent": types.StringType, "device_id": types.StringType}}, "three_r_i": types.ObjectType{AttrTypes: map[string]attr.Type{"previous_authentication": types.StringType, "type": types.StringType}}, "type": types.StringType}},
						"channel",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedChannel, ok := valueChannel.(types.Object); ok {
							state.Channel = typedChannel
							assignedChannel = true
						}
					}
				}
			}
		}
		if !assignedChannel && hadRawChannel {
			if nullChannel, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"browser": types.ObjectType{AttrTypes: map[string]attr.Type{"accept_header": types.StringType, "color_depth": types.Int64Type, "ip_address": types.StringType, "java_enabled": types.BoolType, "javascript_enabled": types.BoolType, "language": types.StringType, "screen_height": types.Int64Type, "screen_width": types.Int64Type, "timezone_offset": types.Int64Type, "user_agent": types.StringType, "device_id": types.StringType}}, "three_r_i": types.ObjectType{AttrTypes: map[string]attr.Type{"previous_authentication": types.StringType, "type": types.StringType}}, "type": types.StringType}}); ok {
				if typedChannel, ok := nullChannel.(types.Object); ok {
					state.Channel = typedChannel
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
		if rawValueDirectoryServer, rawOk := plainValueAtPath(raw, "directory_server"); rawOk {
			if valueDirectoryServer, err := flattenPlainValue(rawValueDirectoryServer, types.StringType, "directory_server", "raw response"); err != nil {
				return err
			} else {
				if typedDirectoryServer, ok := valueDirectoryServer.(types.String); ok {
					state.DirectoryServer = typedDirectoryServer
				}
			}
		} else if !hasRaw {
			if responseValueDirectoryServer, ok := plainFromResponseField(obj, "DirectoryServer"); ok {
				if valueDirectoryServer, err := flattenPlainValue(responseValueDirectoryServer, types.StringType, "directory_server", "response struct"); err != nil {
					return err
				} else {
					if typedDirectoryServer, ok := valueDirectoryServer.(types.String); ok {
						state.DirectoryServer = typedDirectoryServer
					}
				}
			}
		}
	}
	{
		if rawValueFingerprintingURL, rawOk := plainValueAtPath(raw, "fingerprinting_url"); rawOk {
			if valueFingerprintingURL, err := flattenPlainValue(rawValueFingerprintingURL, types.StringType, "fingerprinting_url", "raw response"); err != nil {
				return err
			} else {
				if typedFingerprintingURL, ok := valueFingerprintingURL.(types.String); ok {
					state.FingerprintingURL = typedFingerprintingURL
				}
			}
		} else if !hasRaw {
			if responseValueFingerprintingURL, ok := plainFromResponseField(obj, "FingerprintingURL"); ok {
				if valueFingerprintingURL, err := flattenPlainValue(responseValueFingerprintingURL, types.StringType, "fingerprinting_url", "response struct"); err != nil {
					return err
				} else {
					if typedFingerprintingURL, ok := valueFingerprintingURL.(types.String); ok {
						state.FingerprintingURL = typedFingerprintingURL
					}
				}
			}
		}
	}
	{
		assignedFlowPreference := false
		hadRawFlowPreference := false
		if rawValueFlowPreference, rawOk := plainValueAtPath(raw, "flow_preference"); rawOk {
			hadRawFlowPreference = true
			if rawValueFlowPreference != nil {
				sourceFlowPreference := applyConfiguredKeyedListShapes(rawValueFlowPreference, attrValueToPlain(state.FlowPreference))
				if valueFlowPreference, err := flattenPlainValue(sourceFlowPreference, types.ObjectType{AttrTypes: map[string]attr.Type{"challenge": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "data_share": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "frictionless": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "type": types.StringType}}, "flow_preference", "raw response"); err != nil {
					return err
				} else {
					if typedFlowPreference, ok := valueFlowPreference.(types.Object); ok {
						state.FlowPreference = typedFlowPreference
						assignedFlowPreference = true
					}
				}
			}
		}
		if !assignedFlowPreference {
			if !hasRaw {
				if responseValueFlowPreference, ok := plainFromResponseField(obj, "FlowPreference"); ok {
					sourceFlowPreference := applyConfiguredKeyedListShapes(responseValueFlowPreference, attrValueToPlain(state.FlowPreference))
					if valueFlowPreference, err := flattenPlainValue(
						sourceFlowPreference,
						types.ObjectType{AttrTypes: map[string]attr.Type{"challenge": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "data_share": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "frictionless": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "type": types.StringType}},
						"flow_preference",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFlowPreference, ok := valueFlowPreference.(types.Object); ok {
							state.FlowPreference = typedFlowPreference
							assignedFlowPreference = true
						}
					}
				}
			}
		}
		if !assignedFlowPreference && hadRawFlowPreference {
			if nullFlowPreference, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"challenge": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "data_share": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "frictionless": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "type": types.StringType}}); ok {
				if typedFlowPreference, ok := nullFlowPreference.(types.Object); ok {
					state.FlowPreference = typedFlowPreference
				}
			}
		}
	}
	{
		assignedFutureUsage := false
		hadRawFutureUsage := false
		if rawValueFutureUsage, rawOk := plainValueAtPath(raw, "future_usage"); rawOk {
			hadRawFutureUsage = true
			if rawValueFutureUsage != nil {
				sourceFutureUsage := applyConfiguredKeyedListShapes(rawValueFutureUsage, attrValueToPlain(state.FutureUsage))
				if valueFutureUsage, err := flattenPlainValue(sourceFutureUsage, types.ObjectType{AttrTypes: map[string]attr.Type{"installment": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "expiry": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.StringType, "type": types.StringType}}, "interval": types.StringType, "interval_count": types.Int64Type, "number": types.Int64Type}}, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "expiry": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.StringType, "type": types.StringType}}, "interval": types.StringType, "interval_count": types.Int64Type}}, "type": types.StringType}}, "future_usage", "raw response"); err != nil {
					return err
				} else {
					if typedFutureUsage, ok := valueFutureUsage.(types.Object); ok {
						state.FutureUsage = typedFutureUsage
						assignedFutureUsage = true
					}
				}
			}
		}
		if !assignedFutureUsage {
			if !hasRaw {
				if responseValueFutureUsage, ok := plainFromResponseField(obj, "FutureUsage"); ok {
					sourceFutureUsage := applyConfiguredKeyedListShapes(responseValueFutureUsage, attrValueToPlain(state.FutureUsage))
					if valueFutureUsage, err := flattenPlainValue(
						sourceFutureUsage,
						types.ObjectType{AttrTypes: map[string]attr.Type{"installment": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "expiry": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.StringType, "type": types.StringType}}, "interval": types.StringType, "interval_count": types.Int64Type, "number": types.Int64Type}}, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "expiry": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.StringType, "type": types.StringType}}, "interval": types.StringType, "interval_count": types.Int64Type}}, "type": types.StringType}},
						"future_usage",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFutureUsage, ok := valueFutureUsage.(types.Object); ok {
							state.FutureUsage = typedFutureUsage
							assignedFutureUsage = true
						}
					}
				}
			}
		}
		if !assignedFutureUsage && hadRawFutureUsage {
			if nullFutureUsage, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"installment": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "expiry": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.StringType, "type": types.StringType}}, "interval": types.StringType, "interval_count": types.Int64Type, "number": types.Int64Type}}, "recurring": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "expiry": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.StringType, "type": types.StringType}}, "interval": types.StringType, "interval_count": types.Int64Type}}, "type": types.StringType}}); ok {
				if typedFutureUsage, ok := nullFutureUsage.(types.Object); ok {
					state.FutureUsage = typedFutureUsage
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
		if rawValueMessageCategory, rawOk := plainValueAtPath(raw, "message_category"); rawOk {
			if valueMessageCategory, err := flattenPlainValue(rawValueMessageCategory, types.StringType, "message_category", "raw response"); err != nil {
				return err
			} else {
				if typedMessageCategory, ok := valueMessageCategory.(types.String); ok {
					state.MessageCategory = typedMessageCategory
				}
			}
		} else if !hasRaw {
			if responseValueMessageCategory, ok := plainFromResponseField(obj, "MessageCategory"); ok {
				if valueMessageCategory, err := flattenPlainValue(responseValueMessageCategory, types.StringType, "message_category", "response struct"); err != nil {
					return err
				} else {
					if typedMessageCategory, ok := valueMessageCategory.(types.String); ok {
						state.MessageCategory = typedMessageCategory
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
		if rawValueOutcome, rawOk := plainValueAtPath(raw, "outcome"); rawOk {
			if valueOutcome, err := flattenPlainValue(rawValueOutcome, types.StringType, "outcome", "raw response"); err != nil {
				return err
			} else {
				if typedOutcome, ok := valueOutcome.(types.String); ok {
					state.Outcome = typedOutcome
				}
			}
		} else if !hasRaw {
			if responseValueOutcome, ok := plainFromResponseField(obj, "Outcome"); ok {
				if valueOutcome, err := flattenPlainValue(responseValueOutcome, types.StringType, "outcome", "response struct"); err != nil {
					return err
				} else {
					if typedOutcome, ok := valueOutcome.(types.String); ok {
						state.Outcome = typedOutcome
					}
				}
			}
		}
	}
	{
		assignedOutcomeDetails := false
		hadRawOutcomeDetails := false
		if rawValueOutcomeDetails, rawOk := plainValueAtPath(raw, "outcome_details"); rawOk {
			hadRawOutcomeDetails = true
			if rawValueOutcomeDetails != nil {
				sourceOutcomeDetails := applyConfiguredKeyedListShapes(rawValueOutcomeDetails, attrValueToPlain(state.OutcomeDetails))
				if valueOutcomeDetails, err := flattenPlainValue(sourceOutcomeDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"acs_transaction_id": types.StringType, "ares": types.StringType, "ares_trans_status": types.StringType, "cryptogram": types.StringType, "ds_transaction_id": types.StringType, "eci": types.StringType, "network_details": types.ObjectType{AttrTypes: map[string]attr.Type{"cartes_bancaires": types.ObjectType{AttrTypes: map[string]attr.Type{"avalgo": types.StringType, "cb_exemption": types.StringType, "cb_score": types.StringType}}}}, "protocol_version": types.StringType, "requestor_challenge_indicator": types.StringType, "rreq": types.StringType, "rreq_trans_status": types.StringType, "three_ds_server_transaction_id": types.StringType}}, "outcome_details", "raw response"); err != nil {
					return err
				} else {
					if typedOutcomeDetails, ok := valueOutcomeDetails.(types.Object); ok {
						state.OutcomeDetails = typedOutcomeDetails
						assignedOutcomeDetails = true
					}
				}
			}
		}
		if !assignedOutcomeDetails {
			if !hasRaw {
				if responseValueOutcomeDetails, ok := plainFromResponseField(obj, "OutcomeDetails"); ok {
					sourceOutcomeDetails := applyConfiguredKeyedListShapes(responseValueOutcomeDetails, attrValueToPlain(state.OutcomeDetails))
					if valueOutcomeDetails, err := flattenPlainValue(
						sourceOutcomeDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"acs_transaction_id": types.StringType, "ares": types.StringType, "ares_trans_status": types.StringType, "cryptogram": types.StringType, "ds_transaction_id": types.StringType, "eci": types.StringType, "network_details": types.ObjectType{AttrTypes: map[string]attr.Type{"cartes_bancaires": types.ObjectType{AttrTypes: map[string]attr.Type{"avalgo": types.StringType, "cb_exemption": types.StringType, "cb_score": types.StringType}}}}, "protocol_version": types.StringType, "requestor_challenge_indicator": types.StringType, "rreq": types.StringType, "rreq_trans_status": types.StringType, "three_ds_server_transaction_id": types.StringType}},
						"outcome_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedOutcomeDetails, ok := valueOutcomeDetails.(types.Object); ok {
							state.OutcomeDetails = typedOutcomeDetails
							assignedOutcomeDetails = true
						}
					}
				}
			}
		}
		if !assignedOutcomeDetails && hadRawOutcomeDetails {
			if nullOutcomeDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"acs_transaction_id": types.StringType, "ares": types.StringType, "ares_trans_status": types.StringType, "cryptogram": types.StringType, "ds_transaction_id": types.StringType, "eci": types.StringType, "network_details": types.ObjectType{AttrTypes: map[string]attr.Type{"cartes_bancaires": types.ObjectType{AttrTypes: map[string]attr.Type{"avalgo": types.StringType, "cb_exemption": types.StringType, "cb_score": types.StringType}}}}, "protocol_version": types.StringType, "requestor_challenge_indicator": types.StringType, "rreq": types.StringType, "rreq_trans_status": types.StringType, "three_ds_server_transaction_id": types.StringType}}); ok {
				if typedOutcomeDetails, ok := nullOutcomeDetails.(types.Object); ok {
					state.OutcomeDetails = typedOutcomeDetails
				}
			}
		}
	}
	{
		if state.PaymentMethod.IsNull() || state.PaymentMethod.IsUnknown() {
			if rawValuePaymentMethod, rawOk := plainValueAtPath(raw, "payment_method"); rawOk {
				if typedPaymentMethod, ok := plainToStringIDValue(rawValuePaymentMethod); ok {
					state.PaymentMethod = typedPaymentMethod
				}
			} else if !hasRaw {
				if responseValuePaymentMethod, ok := plainFromResponseField(obj, "PaymentMethod"); ok {
					if typedPaymentMethod, ok := plainToStringIDValue(responseValuePaymentMethod); ok {
						state.PaymentMethod = typedPaymentMethod
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
		assignedShippingAddress := false
		hadRawShippingAddress := false
		if rawValueShippingAddress, rawOk := plainValueAtPath(raw, "shipping_address"); rawOk {
			hadRawShippingAddress = true
			if rawValueShippingAddress != nil {
				sourceShippingAddress := applyConfiguredKeyedListShapes(rawValueShippingAddress, attrValueToPlain(state.ShippingAddress))
				if valueShippingAddress, err := flattenPlainValue(sourceShippingAddress, types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "shipping_address", "raw response"); err != nil {
					return err
				} else {
					if typedShippingAddress, ok := valueShippingAddress.(types.Object); ok {
						state.ShippingAddress = typedShippingAddress
						assignedShippingAddress = true
					}
				}
			}
		}
		if !assignedShippingAddress {
			if !hasRaw {
				if responseValueShippingAddress, ok := plainFromResponseField(obj, "ShippingAddress"); ok {
					sourceShippingAddress := applyConfiguredKeyedListShapes(responseValueShippingAddress, attrValueToPlain(state.ShippingAddress))
					if valueShippingAddress, err := flattenPlainValue(
						sourceShippingAddress,
						types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}},
						"shipping_address",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedShippingAddress, ok := valueShippingAddress.(types.Object); ok {
							state.ShippingAddress = typedShippingAddress
							assignedShippingAddress = true
						}
					}
				}
			}
		}
		if !assignedShippingAddress && hadRawShippingAddress {
			if nullShippingAddress, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}); ok {
				if typedShippingAddress, ok := nullShippingAddress.(types.Object); ok {
					state.ShippingAddress = typedShippingAddress
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
	return nil
}
