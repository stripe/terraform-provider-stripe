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

var _ resource.Resource = &TreasuryFinancialAccountResource{}

var _ resource.ResourceWithConfigure = &TreasuryFinancialAccountResource{}

var _ resource.ResourceWithImportState = &TreasuryFinancialAccountResource{}

func NewTreasuryFinancialAccountResource() resource.Resource {
	return &TreasuryFinancialAccountResource{}
}

type TreasuryFinancialAccountResource struct {
	client *stripe.Client
}

type TreasuryFinancialAccountResourceModel struct {
	Object               types.String `tfsdk:"object"`
	ActiveFeatures       types.List   `tfsdk:"active_features"`
	Balance              types.Object `tfsdk:"balance"`
	Country              types.String `tfsdk:"country"`
	Created              types.Int64  `tfsdk:"created"`
	Features             types.Object `tfsdk:"features"`
	FinancialAddresses   types.List   `tfsdk:"financial_addresses"`
	ID                   types.String `tfsdk:"id"`
	IsDefault            types.Bool   `tfsdk:"is_default"`
	Livemode             types.Bool   `tfsdk:"livemode"`
	Metadata             types.Map    `tfsdk:"metadata"`
	Nickname             types.String `tfsdk:"nickname"`
	PendingFeatures      types.Set    `tfsdk:"pending_features"`
	PlatformRestrictions types.Object `tfsdk:"platform_restrictions"`
	RestrictedFeatures   types.List   `tfsdk:"restricted_features"`
	Status               types.String `tfsdk:"status"`
	StatusDetails        types.Object `tfsdk:"status_details"`
	SupportedCurrencies  types.List   `tfsdk:"supported_currencies"`
}

func (r *TreasuryFinancialAccountResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TreasuryFinancialAccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_treasury_financial_account"
}

func (r *TreasuryFinancialAccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Stripe Treasury provides users with a container for money called a FinancialAccount that is separate from their Payments balance.\nFinancialAccounts serve as the source and destination of Treasury’s money movement APIs.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("treasury.financial_account")},
			},
			"active_features": schema.ListAttribute{
				Computed:      true,
				Description:   "The array of paths to active Features in the Features hash.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"balance": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Balance information for the FinancialAccount",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"cash": schema.MapAttribute{
						Computed:      true,
						Description:   "Funds the user can spend right now.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						ElementType:   types.Int64Type,
					},
					"inbound_pending": schema.MapAttribute{
						Computed:      true,
						Description:   "Funds not spendable yet, but will become available at a later time.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						ElementType:   types.Int64Type,
					},
					"outbound_pending": schema.MapAttribute{
						Computed:      true,
						Description:   "Funds in the account, but not spendable because they are being held for pending outbound flows.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
						ElementType:   types.Int64Type,
					},
				},
			},
			"country": schema.StringAttribute{
				Computed:      true,
				Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"features": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Encodes whether a FinancialAccount has access to a particular Feature, with a `status` enum and associated `status_details`.\nStripe or the platform can control Features via the requested field.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"object": schema.StringAttribute{
						Optional:    true,
						Description: "String representing the object's type. Objects of the same type share the same value.",
						WriteOnly:   true,
						Validators:  []validator.String{stringvalidator.OneOf("treasury.financial_account_features")},
					},
					"card_issuing": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "Toggle settings for enabling/disabling a feature",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"requested": schema.BoolAttribute{
								Required:    true,
								Description: "Whether the FinancialAccount should have the Feature.",
								WriteOnly:   true,
							},
							"status": schema.StringAttribute{
								Optional:    true,
								Description: "Whether the Feature is operational.",
								WriteOnly:   true,
								Validators:  []validator.String{stringvalidator.OneOf("active", "pending", "restricted")},
							},
							"status_details": schema.ListNestedAttribute{
								Optional:    true,
								Description: "Additional details; includes at least one entry when the status is not `active`.",
								WriteOnly:   true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"code": schema.StringAttribute{
											Optional:    true,
											Description: "Represents the reason why the status is `pending` or `restricted`.",
											WriteOnly:   true,
											Validators:  []validator.String{stringvalidator.OneOf("activating", "capability_not_requested", "financial_account_closed", "rejected_other", "rejected_unsupported_business", "requirements_past_due", "requirements_pending_verification", "restricted_by_platform", "restricted_other")},
										},
										"resolution": schema.StringAttribute{
											Optional:    true,
											Description: "Represents what the user should do, if anything, to activate the Feature.",
											WriteOnly:   true,
											Validators:  []validator.String{stringvalidator.OneOf("contact_stripe", "provide_information", "remove_restriction")},
										},
										"restriction": schema.StringAttribute{
											Optional:    true,
											Description: "The `platform_restrictions` that are restricting this Feature.",
											WriteOnly:   true,
											Validators:  []validator.String{stringvalidator.OneOf("inbound_flows", "outbound_flows")},
										},
									},
								},
							},
						},
					},
					"deposit_insurance": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "Toggle settings for enabling/disabling a feature",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"requested": schema.BoolAttribute{
								Required:    true,
								Description: "Whether the FinancialAccount should have the Feature.",
								WriteOnly:   true,
							},
							"status": schema.StringAttribute{
								Optional:    true,
								Description: "Whether the Feature is operational.",
								WriteOnly:   true,
								Validators:  []validator.String{stringvalidator.OneOf("active", "pending", "restricted")},
							},
							"status_details": schema.ListNestedAttribute{
								Optional:    true,
								Description: "Additional details; includes at least one entry when the status is not `active`.",
								WriteOnly:   true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"code": schema.StringAttribute{
											Optional:    true,
											Description: "Represents the reason why the status is `pending` or `restricted`.",
											WriteOnly:   true,
											Validators:  []validator.String{stringvalidator.OneOf("activating", "capability_not_requested", "financial_account_closed", "rejected_other", "rejected_unsupported_business", "requirements_past_due", "requirements_pending_verification", "restricted_by_platform", "restricted_other")},
										},
										"resolution": schema.StringAttribute{
											Optional:    true,
											Description: "Represents what the user should do, if anything, to activate the Feature.",
											WriteOnly:   true,
											Validators:  []validator.String{stringvalidator.OneOf("contact_stripe", "provide_information", "remove_restriction")},
										},
										"restriction": schema.StringAttribute{
											Optional:    true,
											Description: "The `platform_restrictions` that are restricting this Feature.",
											WriteOnly:   true,
											Validators:  []validator.String{stringvalidator.OneOf("inbound_flows", "outbound_flows")},
										},
									},
								},
							},
						},
					},
					"financial_addresses": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "Settings related to Financial Addresses features on a Financial Account",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"aba": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Toggle settings for enabling/disabling the ABA address feature",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"requested": schema.BoolAttribute{
										Required:    true,
										Description: "Whether the FinancialAccount should have the Feature.",
										WriteOnly:   true,
									},
									"status": schema.StringAttribute{
										Optional:    true,
										Description: "Whether the Feature is operational.",
										WriteOnly:   true,
										Validators:  []validator.String{stringvalidator.OneOf("active", "pending", "restricted")},
									},
									"status_details": schema.ListNestedAttribute{
										Optional:    true,
										Description: "Additional details; includes at least one entry when the status is not `active`.",
										WriteOnly:   true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"code": schema.StringAttribute{
													Optional:    true,
													Description: "Represents the reason why the status is `pending` or `restricted`.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("activating", "capability_not_requested", "financial_account_closed", "rejected_other", "rejected_unsupported_business", "requirements_past_due", "requirements_pending_verification", "restricted_by_platform", "restricted_other")},
												},
												"resolution": schema.StringAttribute{
													Optional:    true,
													Description: "Represents what the user should do, if anything, to activate the Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("contact_stripe", "provide_information", "remove_restriction")},
												},
												"restriction": schema.StringAttribute{
													Optional:    true,
													Description: "The `platform_restrictions` that are restricting this Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("inbound_flows", "outbound_flows")},
												},
											},
										},
									},
								},
							},
						},
					},
					"inbound_transfers": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "InboundTransfers contains inbound transfers features for a FinancialAccount.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"ach": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Toggle settings for enabling/disabling an inbound ACH specific feature",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"requested": schema.BoolAttribute{
										Required:    true,
										Description: "Whether the FinancialAccount should have the Feature.",
										WriteOnly:   true,
									},
									"status": schema.StringAttribute{
										Optional:    true,
										Description: "Whether the Feature is operational.",
										WriteOnly:   true,
										Validators:  []validator.String{stringvalidator.OneOf("active", "pending", "restricted")},
									},
									"status_details": schema.ListNestedAttribute{
										Optional:    true,
										Description: "Additional details; includes at least one entry when the status is not `active`.",
										WriteOnly:   true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"code": schema.StringAttribute{
													Optional:    true,
													Description: "Represents the reason why the status is `pending` or `restricted`.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("activating", "capability_not_requested", "financial_account_closed", "rejected_other", "rejected_unsupported_business", "requirements_past_due", "requirements_pending_verification", "restricted_by_platform", "restricted_other")},
												},
												"resolution": schema.StringAttribute{
													Optional:    true,
													Description: "Represents what the user should do, if anything, to activate the Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("contact_stripe", "provide_information", "remove_restriction")},
												},
												"restriction": schema.StringAttribute{
													Optional:    true,
													Description: "The `platform_restrictions` that are restricting this Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("inbound_flows", "outbound_flows")},
												},
											},
										},
									},
								},
							},
						},
					},
					"intra_stripe_flows": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "Toggle settings for enabling/disabling a feature",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"requested": schema.BoolAttribute{
								Required:    true,
								Description: "Whether the FinancialAccount should have the Feature.",
								WriteOnly:   true,
							},
							"status": schema.StringAttribute{
								Optional:    true,
								Description: "Whether the Feature is operational.",
								WriteOnly:   true,
								Validators:  []validator.String{stringvalidator.OneOf("active", "pending", "restricted")},
							},
							"status_details": schema.ListNestedAttribute{
								Optional:    true,
								Description: "Additional details; includes at least one entry when the status is not `active`.",
								WriteOnly:   true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"code": schema.StringAttribute{
											Optional:    true,
											Description: "Represents the reason why the status is `pending` or `restricted`.",
											WriteOnly:   true,
											Validators:  []validator.String{stringvalidator.OneOf("activating", "capability_not_requested", "financial_account_closed", "rejected_other", "rejected_unsupported_business", "requirements_past_due", "requirements_pending_verification", "restricted_by_platform", "restricted_other")},
										},
										"resolution": schema.StringAttribute{
											Optional:    true,
											Description: "Represents what the user should do, if anything, to activate the Feature.",
											WriteOnly:   true,
											Validators:  []validator.String{stringvalidator.OneOf("contact_stripe", "provide_information", "remove_restriction")},
										},
										"restriction": schema.StringAttribute{
											Optional:    true,
											Description: "The `platform_restrictions` that are restricting this Feature.",
											WriteOnly:   true,
											Validators:  []validator.String{stringvalidator.OneOf("inbound_flows", "outbound_flows")},
										},
									},
								},
							},
						},
					},
					"outbound_payments": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "Settings related to Outbound Payments features on a Financial Account",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"ach": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Toggle settings for enabling/disabling an outbound ACH specific feature",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"requested": schema.BoolAttribute{
										Required:    true,
										Description: "Whether the FinancialAccount should have the Feature.",
										WriteOnly:   true,
									},
									"status": schema.StringAttribute{
										Optional:    true,
										Description: "Whether the Feature is operational.",
										WriteOnly:   true,
										Validators:  []validator.String{stringvalidator.OneOf("active", "pending", "restricted")},
									},
									"status_details": schema.ListNestedAttribute{
										Optional:    true,
										Description: "Additional details; includes at least one entry when the status is not `active`.",
										WriteOnly:   true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"code": schema.StringAttribute{
													Optional:    true,
													Description: "Represents the reason why the status is `pending` or `restricted`.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("activating", "capability_not_requested", "financial_account_closed", "rejected_other", "rejected_unsupported_business", "requirements_past_due", "requirements_pending_verification", "restricted_by_platform", "restricted_other")},
												},
												"resolution": schema.StringAttribute{
													Optional:    true,
													Description: "Represents what the user should do, if anything, to activate the Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("contact_stripe", "provide_information", "remove_restriction")},
												},
												"restriction": schema.StringAttribute{
													Optional:    true,
													Description: "The `platform_restrictions` that are restricting this Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("inbound_flows", "outbound_flows")},
												},
											},
										},
									},
								},
							},
							"us_domestic_wire": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Toggle settings for enabling/disabling a feature",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"requested": schema.BoolAttribute{
										Required:    true,
										Description: "Whether the FinancialAccount should have the Feature.",
										WriteOnly:   true,
									},
									"status": schema.StringAttribute{
										Optional:    true,
										Description: "Whether the Feature is operational.",
										WriteOnly:   true,
										Validators:  []validator.String{stringvalidator.OneOf("active", "pending", "restricted")},
									},
									"status_details": schema.ListNestedAttribute{
										Optional:    true,
										Description: "Additional details; includes at least one entry when the status is not `active`.",
										WriteOnly:   true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"code": schema.StringAttribute{
													Optional:    true,
													Description: "Represents the reason why the status is `pending` or `restricted`.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("activating", "capability_not_requested", "financial_account_closed", "rejected_other", "rejected_unsupported_business", "requirements_past_due", "requirements_pending_verification", "restricted_by_platform", "restricted_other")},
												},
												"resolution": schema.StringAttribute{
													Optional:    true,
													Description: "Represents what the user should do, if anything, to activate the Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("contact_stripe", "provide_information", "remove_restriction")},
												},
												"restriction": schema.StringAttribute{
													Optional:    true,
													Description: "The `platform_restrictions` that are restricting this Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("inbound_flows", "outbound_flows")},
												},
											},
										},
									},
								},
							},
						},
					},
					"outbound_transfers": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "OutboundTransfers contains outbound transfers features for a FinancialAccount.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"ach": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Toggle settings for enabling/disabling an outbound ACH specific feature",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"requested": schema.BoolAttribute{
										Required:    true,
										Description: "Whether the FinancialAccount should have the Feature.",
										WriteOnly:   true,
									},
									"status": schema.StringAttribute{
										Optional:    true,
										Description: "Whether the Feature is operational.",
										WriteOnly:   true,
										Validators:  []validator.String{stringvalidator.OneOf("active", "pending", "restricted")},
									},
									"status_details": schema.ListNestedAttribute{
										Optional:    true,
										Description: "Additional details; includes at least one entry when the status is not `active`.",
										WriteOnly:   true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"code": schema.StringAttribute{
													Optional:    true,
													Description: "Represents the reason why the status is `pending` or `restricted`.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("activating", "capability_not_requested", "financial_account_closed", "rejected_other", "rejected_unsupported_business", "requirements_past_due", "requirements_pending_verification", "restricted_by_platform", "restricted_other")},
												},
												"resolution": schema.StringAttribute{
													Optional:    true,
													Description: "Represents what the user should do, if anything, to activate the Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("contact_stripe", "provide_information", "remove_restriction")},
												},
												"restriction": schema.StringAttribute{
													Optional:    true,
													Description: "The `platform_restrictions` that are restricting this Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("inbound_flows", "outbound_flows")},
												},
											},
										},
									},
								},
							},
							"us_domestic_wire": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "Toggle settings for enabling/disabling a feature",
								WriteOnly:   true,
								Attributes: map[string]schema.Attribute{
									"requested": schema.BoolAttribute{
										Required:    true,
										Description: "Whether the FinancialAccount should have the Feature.",
										WriteOnly:   true,
									},
									"status": schema.StringAttribute{
										Optional:    true,
										Description: "Whether the Feature is operational.",
										WriteOnly:   true,
										Validators:  []validator.String{stringvalidator.OneOf("active", "pending", "restricted")},
									},
									"status_details": schema.ListNestedAttribute{
										Optional:    true,
										Description: "Additional details; includes at least one entry when the status is not `active`.",
										WriteOnly:   true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"code": schema.StringAttribute{
													Optional:    true,
													Description: "Represents the reason why the status is `pending` or `restricted`.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("activating", "capability_not_requested", "financial_account_closed", "rejected_other", "rejected_unsupported_business", "requirements_past_due", "requirements_pending_verification", "restricted_by_platform", "restricted_other")},
												},
												"resolution": schema.StringAttribute{
													Optional:    true,
													Description: "Represents what the user should do, if anything, to activate the Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("contact_stripe", "provide_information", "remove_restriction")},
												},
												"restriction": schema.StringAttribute{
													Optional:    true,
													Description: "The `platform_restrictions` that are restricting this Feature.",
													WriteOnly:   true,
													Validators:  []validator.String{stringvalidator.OneOf("inbound_flows", "outbound_flows")},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"financial_addresses": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "The set of credentials that resolve to a FinancialAccount.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"aba": schema.SingleNestedAttribute{
							Computed:      true,
							Description:   "ABA Records contain U.S. bank account details per the ABA format.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"account_holder_name": schema.StringAttribute{
									Computed:      true,
									Description:   "The name of the person or business that owns the bank account.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"account_number": schema.StringAttribute{
									Computed:      true,
									Description:   "The account number.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"account_number_last4": schema.StringAttribute{
									Computed:      true,
									Description:   "The last four characters of the account number.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"bank_name": schema.StringAttribute{
									Computed:      true,
									Description:   "Name of the bank.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"routing_number": schema.StringAttribute{
									Computed:      true,
									Description:   "Routing number for the account.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
							},
						},
						"supported_networks": schema.ListAttribute{
							Computed:      true,
							Description:   "The list of networks that the address supports",
							PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
							ElementType:   types.StringType,
						},
						"type": schema.StringAttribute{
							Computed:      true,
							Description:   "The type of financial address",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							Validators:    []validator.String{stringvalidator.OneOf("aba")},
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"is_default": schema.BoolAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
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
			"nickname": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The nickname for the FinancialAccount.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"pending_features": schema.SetAttribute{
				Computed:    true,
				Description: "The array of paths to pending Features in the Features hash.",
				ElementType: types.StringType,
			},
			"platform_restrictions": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The set of functionalities that the platform can restrict on the FinancialAccount.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"inbound_flows": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Restricts all inbound money movement.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("restricted", "unrestricted")},
					},
					"outbound_flows": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Restricts all outbound money movement.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("restricted", "unrestricted")},
					},
				},
			},
			"restricted_features": schema.ListAttribute{
				Computed:      true,
				Description:   "The array of paths to restricted Features in the Features hash.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Status of this FinancialAccount.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("closed", "open")},
			},
			"status_details": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"closed": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Details related to the closure of this FinancialAccount",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"reasons": schema.ListAttribute{
								Computed:      true,
								Description:   "The array that contains reasons for a FinancialAccount closure.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
						},
					},
				},
			},
			"supported_currencies": schema.ListAttribute{
				Required:      true,
				Description:   "The currencies the FinancialAccount can hold a balance in. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				ElementType:   types.StringType,
			},
		},
	}
}

func (r *TreasuryFinancialAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TreasuryFinancialAccountResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config TreasuryFinancialAccountResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Features"}})

	params, err := expandTreasuryFinancialAccountCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building TreasuryFinancialAccount create params", err.Error())
		return
	}

	obj, err := r.client.V1TreasuryFinancialAccounts.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating TreasuryFinancialAccount", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TreasuryFinancialAccounts.B, r.client.V1TreasuryFinancialAccounts.Key, stripe.FormatURLPath("/v1/treasury/financial_accounts/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TreasuryFinancialAccount create raw response", err.Error())
		return
	}

	if err := flattenTreasuryFinancialAccount(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TreasuryFinancialAccount create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Features"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TreasuryFinancialAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState TreasuryFinancialAccountResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TreasuryFinancialAccountResourceModel
	state = priorState

	obj, err := r.client.V1TreasuryFinancialAccounts.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading TreasuryFinancialAccount", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TreasuryFinancialAccounts.B, r.client.V1TreasuryFinancialAccounts.Key, stripe.FormatURLPath("/v1/treasury/financial_accounts/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TreasuryFinancialAccount raw response", err.Error())
		return
	}

	if err := flattenTreasuryFinancialAccount(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening TreasuryFinancialAccount read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"Features"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *TreasuryFinancialAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TreasuryFinancialAccountResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config TreasuryFinancialAccountResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Features"}})

	var state TreasuryFinancialAccountResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandTreasuryFinancialAccountUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building TreasuryFinancialAccount update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building TreasuryFinancialAccount update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1TreasuryFinancialAccounts.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating TreasuryFinancialAccount", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TreasuryFinancialAccounts.B, r.client.V1TreasuryFinancialAccounts.Key, stripe.FormatURLPath("/v1/treasury/financial_accounts/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TreasuryFinancialAccount update raw response", err.Error())
		return
	}

	if err := flattenTreasuryFinancialAccount(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TreasuryFinancialAccount update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Features"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TreasuryFinancialAccountResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *TreasuryFinancialAccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandTreasuryFinancialAccountCreate(plan TreasuryFinancialAccountResourceModel) (*stripe.TreasuryFinancialAccountCreateParams, error) {
	params := &stripe.TreasuryFinancialAccountCreateParams{}

	if !plan.Features.IsNull() && !plan.Features.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Features", plan.Features) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "features", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Nickname.IsNull() && !plan.Nickname.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Nickname", "Nickname", plan.Nickname.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "nickname", params)
		}
	}
	if !plan.PlatformRestrictions.IsNull() && !plan.PlatformRestrictions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PlatformRestrictions", plan.PlatformRestrictions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "platform_restrictions", params)
		}
	}
	if !plan.SupportedCurrencies.IsNull() && !plan.SupportedCurrencies.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SupportedCurrencies", plan.SupportedCurrencies) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "supported_currencies", params)
		}
	}

	return params, nil
}

func expandTreasuryFinancialAccountUpdate(plan TreasuryFinancialAccountResourceModel, state TreasuryFinancialAccountResourceModel) (*stripe.TreasuryFinancialAccountUpdateParams, error) {
	params := &stripe.TreasuryFinancialAccountUpdateParams{}

	if !plan.Features.Equal(state.Features) && !plan.Features.IsNull() && !plan.Features.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Features", plan.Features) {
			if !plan.Features.Equal(state.Features) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "features", params)
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
	if !plan.Nickname.Equal(state.Nickname) && !plan.Nickname.IsNull() && !plan.Nickname.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Nickname", "Nickname", plan.Nickname.ValueString()) {
			if !plan.Nickname.Equal(state.Nickname) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "nickname", params)
			}
		}
	}
	if !plan.PlatformRestrictions.Equal(state.PlatformRestrictions) && !plan.PlatformRestrictions.IsNull() && !plan.PlatformRestrictions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PlatformRestrictions", plan.PlatformRestrictions) {
			if !plan.PlatformRestrictions.Equal(state.PlatformRestrictions) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "platform_restrictions", params)
			}
		}
	}

	return params, nil
}

func flattenTreasuryFinancialAccount(obj *stripe.TreasuryFinancialAccount, state *TreasuryFinancialAccountResourceModel) error {
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
		if rawValueActiveFeatures, rawOk := plainValueAtPath(raw, "active_features"); rawOk {
			if valueActiveFeatures, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueActiveFeatures, attrValueToPlain(state.ActiveFeatures)), types.ListType{ElemType: types.StringType}, "active_features", "raw response"); err != nil {
				return err
			} else {
				if typedActiveFeatures, ok := valueActiveFeatures.(types.List); ok {
					state.ActiveFeatures = typedActiveFeatures
				}
			}
		} else if !hasRaw {
			if responseValueActiveFeatures, ok := plainFromResponseField(obj, "ActiveFeatures"); ok {
				if valueActiveFeatures, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueActiveFeatures, attrValueToPlain(state.ActiveFeatures)),
					types.ListType{ElemType: types.StringType},
					"active_features",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedActiveFeatures, ok := valueActiveFeatures.(types.List); ok {
						state.ActiveFeatures = typedActiveFeatures
					}
				}
			}
		}
	}
	{
		assignedBalance := false
		hadRawBalance := false
		if rawValueBalance, rawOk := plainValueAtPath(raw, "balance"); rawOk {
			hadRawBalance = true
			if rawValueBalance != nil {
				sourceBalance := applyConfiguredKeyedListShapes(rawValueBalance, attrValueToPlain(state.Balance))
				if valueBalance, err := flattenPlainValue(sourceBalance, types.ObjectType{AttrTypes: map[string]attr.Type{"cash": types.MapType{ElemType: types.Int64Type}, "inbound_pending": types.MapType{ElemType: types.Int64Type}, "outbound_pending": types.MapType{ElemType: types.Int64Type}}}, "balance", "raw response"); err != nil {
					return err
				} else {
					if typedBalance, ok := valueBalance.(types.Object); ok {
						state.Balance = typedBalance
						assignedBalance = true
					}
				}
			}
		}
		if !assignedBalance {
			if !hasRaw {
				if responseValueBalance, ok := plainFromResponseField(obj, "Balance"); ok {
					sourceBalance := applyConfiguredKeyedListShapes(responseValueBalance, attrValueToPlain(state.Balance))
					if valueBalance, err := flattenPlainValue(
						sourceBalance,
						types.ObjectType{AttrTypes: map[string]attr.Type{"cash": types.MapType{ElemType: types.Int64Type}, "inbound_pending": types.MapType{ElemType: types.Int64Type}, "outbound_pending": types.MapType{ElemType: types.Int64Type}}},
						"balance",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBalance, ok := valueBalance.(types.Object); ok {
							state.Balance = typedBalance
							assignedBalance = true
						}
					}
				}
			}
		}
		if !assignedBalance && hadRawBalance {
			if nullBalance, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"cash": types.MapType{ElemType: types.Int64Type}, "inbound_pending": types.MapType{ElemType: types.Int64Type}, "outbound_pending": types.MapType{ElemType: types.Int64Type}}}); ok {
				if typedBalance, ok := nullBalance.(types.Object); ok {
					state.Balance = typedBalance
				}
			}
		}
	}
	{
		if rawValueCountry, rawOk := plainValueAtPath(raw, "country"); rawOk {
			if valueCountry, err := flattenPlainValue(rawValueCountry, types.StringType, "country", "raw response"); err != nil {
				return err
			} else {
				if typedCountry, ok := valueCountry.(types.String); ok {
					state.Country = typedCountry
				}
			}
		} else if !hasRaw {
			if responseValueCountry, ok := plainFromResponseField(obj, "Country"); ok {
				if valueCountry, err := flattenPlainValue(responseValueCountry, types.StringType, "country", "response struct"); err != nil {
					return err
				} else {
					if typedCountry, ok := valueCountry.(types.String); ok {
						state.Country = typedCountry
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
		assignedFeatures := false
		hadRawFeatures := false
		if rawValueFeatures, rawOk := plainValueAtPath(raw, "features"); rawOk {
			hadRawFeatures = true
			if rawValueFeatures != nil {
				sourceFeatures := applyConfiguredKeyedListShapes(rawValueFeatures, attrValueToPlain(state.Features))
				if valueFeatures, err := flattenPlainValue(sourceFeatures, types.ObjectType{AttrTypes: map[string]attr.Type{"object": types.StringType, "card_issuing": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "deposit_insurance": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "financial_addresses": types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}, "inbound_transfers": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}, "intra_stripe_flows": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "outbound_payments": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "us_domestic_wire": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}, "outbound_transfers": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "us_domestic_wire": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}}}, "features", "raw response"); err != nil {
					return err
				} else {
					if typedFeatures, ok := valueFeatures.(types.Object); ok {
						state.Features = typedFeatures
						assignedFeatures = true
					}
				}
			}
		}
		if !assignedFeatures {
			if !hasRaw {
				if responseValueFeatures, ok := plainFromResponseField(obj, "Features"); ok {
					sourceFeatures := applyConfiguredKeyedListShapes(responseValueFeatures, attrValueToPlain(state.Features))
					if valueFeatures, err := flattenPlainValue(
						sourceFeatures,
						types.ObjectType{AttrTypes: map[string]attr.Type{"object": types.StringType, "card_issuing": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "deposit_insurance": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "financial_addresses": types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}, "inbound_transfers": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}, "intra_stripe_flows": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "outbound_payments": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "us_domestic_wire": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}, "outbound_transfers": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "us_domestic_wire": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}}},
						"features",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFeatures, ok := valueFeatures.(types.Object); ok {
							state.Features = typedFeatures
							assignedFeatures = true
						}
					}
				}
			}
		}
		if !assignedFeatures && hadRawFeatures {
			if nullFeatures, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"object": types.StringType, "card_issuing": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "deposit_insurance": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "financial_addresses": types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}, "inbound_transfers": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}, "intra_stripe_flows": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "outbound_payments": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "us_domestic_wire": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}, "outbound_transfers": types.ObjectType{AttrTypes: map[string]attr.Type{"ach": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}, "us_domestic_wire": types.ObjectType{AttrTypes: map[string]attr.Type{"requested": types.BoolType, "status": types.StringType, "status_details": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "resolution": types.StringType, "restriction": types.StringType}}}}}}}}}); ok {
				if typedFeatures, ok := nullFeatures.(types.Object); ok {
					state.Features = typedFeatures
				}
			}
		}
	}
	{
		if rawValueFinancialAddresses, rawOk := plainValueAtPath(raw, "financial_addresses"); rawOk {
			if valueFinancialAddresses, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueFinancialAddresses, attrValueToPlain(state.FinancialAddresses)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "account_number": types.StringType, "account_number_last4": types.StringType, "bank_name": types.StringType, "routing_number": types.StringType}}, "supported_networks": types.ListType{ElemType: types.StringType}, "type": types.StringType}}}, "financial_addresses", "raw response"); err != nil {
				return err
			} else {
				if typedFinancialAddresses, ok := valueFinancialAddresses.(types.List); ok {
					state.FinancialAddresses = typedFinancialAddresses
				}
			}
		} else if !hasRaw {
			if responseValueFinancialAddresses, ok := plainFromResponseField(obj, "FinancialAddresses"); ok {
				if valueFinancialAddresses, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueFinancialAddresses, attrValueToPlain(state.FinancialAddresses)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "account_number": types.StringType, "account_number_last4": types.StringType, "bank_name": types.StringType, "routing_number": types.StringType}}, "supported_networks": types.ListType{ElemType: types.StringType}, "type": types.StringType}}},
					"financial_addresses",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedFinancialAddresses, ok := valueFinancialAddresses.(types.List); ok {
						state.FinancialAddresses = typedFinancialAddresses
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
		if rawValueNickname, rawOk := plainValueAtPath(raw, "nickname"); rawOk {
			if valueNickname, err := flattenPlainValue(rawValueNickname, types.StringType, "nickname", "raw response"); err != nil {
				return err
			} else {
				if typedNickname, ok := valueNickname.(types.String); ok {
					state.Nickname = typedNickname
				}
			}
		} else if !hasRaw {
			if responseValueNickname, ok := plainFromResponseField(obj, "Nickname"); ok {
				if valueNickname, err := flattenPlainValue(responseValueNickname, types.StringType, "nickname", "response struct"); err != nil {
					return err
				} else {
					if typedNickname, ok := valueNickname.(types.String); ok {
						state.Nickname = typedNickname
					}
				}
			}
		}
	}
	{
		if rawValuePendingFeatures, rawOk := plainValueAtPath(raw, "pending_features"); rawOk {
			if valuePendingFeatures, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValuePendingFeatures, attrValueToPlain(state.PendingFeatures)), types.SetType{ElemType: types.StringType}, "pending_features", "raw response"); err != nil {
				return err
			} else {
				if typedPendingFeatures, ok := valuePendingFeatures.(types.Set); ok {
					state.PendingFeatures = typedPendingFeatures
				}
			}
		} else if !hasRaw {
			if responseValuePendingFeatures, ok := plainFromResponseField(obj, "PendingFeatures"); ok {
				if valuePendingFeatures, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValuePendingFeatures, attrValueToPlain(state.PendingFeatures)),
					types.SetType{ElemType: types.StringType},
					"pending_features",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedPendingFeatures, ok := valuePendingFeatures.(types.Set); ok {
						state.PendingFeatures = typedPendingFeatures
					}
				}
			}
		}
	}
	{
		assignedPlatformRestrictions := false
		hadRawPlatformRestrictions := false
		if rawValuePlatformRestrictions, rawOk := plainValueAtPath(raw, "platform_restrictions"); rawOk {
			hadRawPlatformRestrictions = true
			if rawValuePlatformRestrictions != nil {
				sourcePlatformRestrictions := applyConfiguredKeyedListShapes(rawValuePlatformRestrictions, attrValueToPlain(state.PlatformRestrictions))
				if valuePlatformRestrictions, err := flattenPlainValue(sourcePlatformRestrictions, types.ObjectType{AttrTypes: map[string]attr.Type{"inbound_flows": types.StringType, "outbound_flows": types.StringType}}, "platform_restrictions", "raw response"); err != nil {
					return err
				} else {
					if typedPlatformRestrictions, ok := valuePlatformRestrictions.(types.Object); ok {
						state.PlatformRestrictions = typedPlatformRestrictions
						assignedPlatformRestrictions = true
					}
				}
			}
		}
		if !assignedPlatformRestrictions {
			if !hasRaw {
				if responseValuePlatformRestrictions, ok := plainFromResponseField(obj, "PlatformRestrictions"); ok {
					sourcePlatformRestrictions := applyConfiguredKeyedListShapes(responseValuePlatformRestrictions, attrValueToPlain(state.PlatformRestrictions))
					if valuePlatformRestrictions, err := flattenPlainValue(
						sourcePlatformRestrictions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"inbound_flows": types.StringType, "outbound_flows": types.StringType}},
						"platform_restrictions",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPlatformRestrictions, ok := valuePlatformRestrictions.(types.Object); ok {
							state.PlatformRestrictions = typedPlatformRestrictions
							assignedPlatformRestrictions = true
						}
					}
				}
			}
		}
		if !assignedPlatformRestrictions && hadRawPlatformRestrictions {
			if nullPlatformRestrictions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"inbound_flows": types.StringType, "outbound_flows": types.StringType}}); ok {
				if typedPlatformRestrictions, ok := nullPlatformRestrictions.(types.Object); ok {
					state.PlatformRestrictions = typedPlatformRestrictions
				}
			}
		}
	}
	{
		if rawValueRestrictedFeatures, rawOk := plainValueAtPath(raw, "restricted_features"); rawOk {
			if valueRestrictedFeatures, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueRestrictedFeatures, attrValueToPlain(state.RestrictedFeatures)), types.ListType{ElemType: types.StringType}, "restricted_features", "raw response"); err != nil {
				return err
			} else {
				if typedRestrictedFeatures, ok := valueRestrictedFeatures.(types.List); ok {
					state.RestrictedFeatures = typedRestrictedFeatures
				}
			}
		} else if !hasRaw {
			if responseValueRestrictedFeatures, ok := plainFromResponseField(obj, "RestrictedFeatures"); ok {
				if valueRestrictedFeatures, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueRestrictedFeatures, attrValueToPlain(state.RestrictedFeatures)),
					types.ListType{ElemType: types.StringType},
					"restricted_features",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedRestrictedFeatures, ok := valueRestrictedFeatures.(types.List); ok {
						state.RestrictedFeatures = typedRestrictedFeatures
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
		assignedStatusDetails := false
		hadRawStatusDetails := false
		if rawValueStatusDetails, rawOk := plainValueAtPath(raw, "status_details"); rawOk {
			hadRawStatusDetails = true
			if rawValueStatusDetails != nil {
				sourceStatusDetails := applyConfiguredKeyedListShapes(rawValueStatusDetails, attrValueToPlain(state.StatusDetails))
				if valueStatusDetails, err := flattenPlainValue(sourceStatusDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"closed": types.ObjectType{AttrTypes: map[string]attr.Type{"reasons": types.ListType{ElemType: types.StringType}}}}}, "status_details", "raw response"); err != nil {
					return err
				} else {
					if typedStatusDetails, ok := valueStatusDetails.(types.Object); ok {
						state.StatusDetails = typedStatusDetails
						assignedStatusDetails = true
					}
				}
			}
		}
		if !assignedStatusDetails {
			if !hasRaw {
				if responseValueStatusDetails, ok := plainFromResponseField(obj, "StatusDetails"); ok {
					sourceStatusDetails := applyConfiguredKeyedListShapes(responseValueStatusDetails, attrValueToPlain(state.StatusDetails))
					if valueStatusDetails, err := flattenPlainValue(
						sourceStatusDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"closed": types.ObjectType{AttrTypes: map[string]attr.Type{"reasons": types.ListType{ElemType: types.StringType}}}}},
						"status_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedStatusDetails, ok := valueStatusDetails.(types.Object); ok {
							state.StatusDetails = typedStatusDetails
							assignedStatusDetails = true
						}
					}
				}
			}
		}
		if !assignedStatusDetails && hadRawStatusDetails {
			if nullStatusDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"closed": types.ObjectType{AttrTypes: map[string]attr.Type{"reasons": types.ListType{ElemType: types.StringType}}}}}); ok {
				if typedStatusDetails, ok := nullStatusDetails.(types.Object); ok {
					state.StatusDetails = typedStatusDetails
				}
			}
		}
	}
	{
		if rawValueSupportedCurrencies, rawOk := plainValueAtPath(raw, "supported_currencies"); rawOk {
			if valueSupportedCurrencies, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueSupportedCurrencies, attrValueToPlain(state.SupportedCurrencies)), types.ListType{ElemType: types.StringType}, "supported_currencies", "raw response"); err != nil {
				return err
			} else {
				if typedSupportedCurrencies, ok := valueSupportedCurrencies.(types.List); ok {
					state.SupportedCurrencies = typedSupportedCurrencies
				}
			}
		} else if !hasRaw {
			if responseValueSupportedCurrencies, ok := plainFromResponseField(obj, "SupportedCurrencies"); ok {
				if valueSupportedCurrencies, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueSupportedCurrencies, attrValueToPlain(state.SupportedCurrencies)),
					types.ListType{ElemType: types.StringType},
					"supported_currencies",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedSupportedCurrencies, ok := valueSupportedCurrencies.(types.List); ok {
						state.SupportedCurrencies = typedSupportedCurrencies
					}
				}
			}
		}
	}
	return nil
}
