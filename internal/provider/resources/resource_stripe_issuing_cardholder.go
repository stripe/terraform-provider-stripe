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

var _ resource.Resource = &IssuingCardholderResource{}

var _ resource.ResourceWithConfigure = &IssuingCardholderResource{}

var _ resource.ResourceWithImportState = &IssuingCardholderResource{}

func NewIssuingCardholderResource() resource.Resource {
	return &IssuingCardholderResource{}
}

type IssuingCardholderResource struct {
	client *stripe.Client
}

type IssuingCardholderResourceModel struct {
	Object           types.String `tfsdk:"object"`
	Billing          types.Object `tfsdk:"billing"`
	Company          types.Object `tfsdk:"company"`
	Created          types.Int64  `tfsdk:"created"`
	Email            types.String `tfsdk:"email"`
	ID               types.String `tfsdk:"id"`
	Individual       types.Object `tfsdk:"individual"`
	Livemode         types.Bool   `tfsdk:"livemode"`
	Metadata         types.Map    `tfsdk:"metadata"`
	Name             types.String `tfsdk:"name"`
	PhoneNumber      types.String `tfsdk:"phone_number"`
	PreferredLocales types.List   `tfsdk:"preferred_locales"`
	Requirements     types.Object `tfsdk:"requirements"`
	SpendingControls types.Object `tfsdk:"spending_controls"`
	Status           types.String `tfsdk:"status"`
	Type             types.String `tfsdk:"type"`
}

func (r *IssuingCardholderResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IssuingCardholderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_issuing_cardholder"
}

func (r *IssuingCardholderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "An Issuing `Cardholder` object represents an individual or business entity who is [issued](https://docs.stripe.com/issuing) cards.\n\nRelated guide: [How to create a cardholder](https://docs.stripe.com/issuing/cards/virtual/issue-cards#create-cardholder)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("issuing.cardholder")},
			},
			"billing": schema.SingleNestedAttribute{
				Required: true,

				Attributes: map[string]schema.Attribute{
					"address": schema.SingleNestedAttribute{
						Required: true,

						Attributes: map[string]schema.Attribute{
							"city": schema.StringAttribute{
								Required:    true,
								Description: "City, district, suburb, town, or village.",
							},
							"country": schema.StringAttribute{
								Required:    true,
								Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
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
								Required:    true,
								Description: "ZIP or postal code.",
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
			"company": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Additional information about a `company` cardholder.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"tax_id_provided": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether the company's business ID number was provided.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"tax_id": schema.StringAttribute{
						Optional:    true,
						Description: "The entity's business ID number.",
					},
				},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"email": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The cardholder's email address.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"individual": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Additional information about an `individual` cardholder.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"card_issuing": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Information related to the card_issuing program for this cardholder.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"user_terms_acceptance": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Information about cardholder acceptance of Celtic [Authorized User Terms](https://stripe.com/docs/issuing/cards#accept-authorized-user-terms). Required for cards backed by a Celtic program.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"date": schema.Int64Attribute{
										Optional:      true,
										Computed:      true,
										Description:   "The Unix timestamp marking when the cardholder accepted the Authorized User Terms.",
										PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
									},
									"ip": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The IP address from which the cardholder accepted the Authorized User Terms.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"user_agent": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The user agent of the browser from which the cardholder accepted the Authorized User Terms.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
						},
					},
					"dob": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The date of birth of this cardholder.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"day": schema.Int64Attribute{
								Required:    true,
								Description: "The day of birth, between 1 and 31.",
							},
							"month": schema.Int64Attribute{
								Required:    true,
								Description: "The month of birth, between 1 and 12.",
							},
							"year": schema.Int64Attribute{
								Required:    true,
								Description: "The four-digit year of birth.",
							},
						},
					},
					"first_name": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The first name of this cardholder. Required before activating Cards. This field cannot contain any numbers, special characters (except periods, commas, hyphens, spaces and apostrophes) or non-latin letters.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"last_name": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The last name of this cardholder. Required before activating Cards. This field cannot contain any numbers, special characters (except periods, commas, hyphens, spaces and apostrophes) or non-latin letters.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"verification": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Government-issued ID document for this cardholder.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"document": schema.SingleNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "An identifying document, either a passport or local ID card.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"back": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The back of a document returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"front": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "The front of a document returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
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
				Required:      true,
				Description:   "The cardholder's name. This will be printed on cards issued to them.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"phone_number": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The cardholder's phone number. This is required for all cardholders who will be creating EU cards. See the [3D Secure documentation](https://docs.stripe.com/issuing/3d-secure#when-is-3d-secure-applied) for more details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"preferred_locales": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The cardholder’s preferred locales (languages), ordered by preference. Locales can be `da`, `de`, `en`, `es`, `fr`, `it`, `pl`, or `sv`.\n This changes the language of the [3D Secure flow](https://docs.stripe.com/issuing/3d-secure) and one-time password messages sent to the cardholder.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"requirements": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"past_due": schema.ListAttribute{
						Computed:      true,
						Description:   "Array of fields that need to be collected in order to verify and re-enable the cardholder.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
				},
			},
			"spending_controls": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Rules that control spending across this cardholder's cards. Refer to our [documentation](https://docs.stripe.com/issuing/controls/spending-controls) for more details.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"allowed_card_presences": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of card presence statuses from which authorizations will be allowed. Possible options are `present`, `not_present`. All other statuses will be blocked. Cannot be set with `blocked_card_presences`. Provide an empty value to unset this control.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"allowed_categories": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of strings containing [categories](https://docs.stripe.com/api#issuing_authorization_object-merchant_data-category) of authorizations to allow. All other categories will be blocked. Cannot be set with `blocked_categories`.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"allowed_merchant_countries": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of strings containing representing countries from which authorizations will be allowed. Authorizations from merchants in all other countries will be declined. Country codes should be ISO 3166 alpha-2 country codes (e.g. `US`). Cannot be set with `blocked_merchant_countries`. Provide an empty value to unset this control.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"blocked_card_presences": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of card presence statuses from which authorizations will be declined. Possible options are `present`, `not_present`. Cannot be set with `allowed_card_presences`. Provide an empty value to unset this control.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"blocked_categories": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of strings containing [categories](https://docs.stripe.com/api#issuing_authorization_object-merchant_data-category) of authorizations to decline. All other categories will be allowed. Cannot be set with `allowed_categories`.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"blocked_merchant_countries": schema.ListAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Array of strings containing representing countries from which authorizations will be declined. Country codes should be ISO 3166 alpha-2 country codes (e.g. `US`). Cannot be set with `allowed_merchant_countries`. Provide an empty value to unset this control.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"spending_limits": schema.ListNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Limit spending with amount-based rules that apply across this cardholder's cards.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"amount": schema.Int64Attribute{
									Required:    true,
									Description: "Maximum amount allowed to spend per interval. This amount is in the card's currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).",
								},
								"categories": schema.ListAttribute{
									Optional:      true,
									Computed:      true,
									Description:   "Array of strings containing [categories](https://docs.stripe.com/api#issuing_authorization_object-merchant_data-category) this limit applies to. Omitting this field will apply the limit to all categories.",
									PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
									ElementType:   types.StringType,
								},
								"interval": schema.StringAttribute{
									Required:    true,
									Description: "Interval (or event) to which the amount applies.",
									Validators:  []validator.String{stringvalidator.OneOf("all_time", "daily", "monthly", "per_authorization", "weekly", "yearly")},
								},
							},
						},
					},
					"spending_limits_currency": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Currency of the amounts within `spending_limits`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"status": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Specifies whether to permit authorizations on this cardholder's cards.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "blocked", "inactive")},
			},
			"type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "One of `individual` or `company`. See [Choose a cardholder type](https://docs.stripe.com/issuing/other/choose-cardholder) for more details.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("company", "individual")},
			},
		},
	}
}

func (r *IssuingCardholderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IssuingCardholderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config IssuingCardholderResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandIssuingCardholderCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building IssuingCardholder create params", err.Error())
		return
	}

	obj, err := r.client.V1IssuingCardholders.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating IssuingCardholder", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingCardholders.B, r.client.V1IssuingCardholders.Key, stripe.FormatURLPath("/v1/issuing/cardholders/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingCardholder create raw response", err.Error())
		return
	}

	if err := flattenIssuingCardholder(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingCardholder create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Company", "tax_id"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingCardholderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState IssuingCardholderResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state IssuingCardholderResourceModel
	state = priorState

	obj, err := r.client.V1IssuingCardholders.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading IssuingCardholder", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingCardholders.B, r.client.V1IssuingCardholders.Key, stripe.FormatURLPath("/v1/issuing/cardholders/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingCardholder raw response", err.Error())
		return
	}

	if err := flattenIssuingCardholder(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingCardholder read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"Company", "tax_id"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingCardholderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IssuingCardholderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config IssuingCardholderResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state IssuingCardholderResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandIssuingCardholderUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building IssuingCardholder update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building IssuingCardholder update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1IssuingCardholders.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating IssuingCardholder", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingCardholders.B, r.client.V1IssuingCardholders.Key, stripe.FormatURLPath("/v1/issuing/cardholders/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingCardholder update raw response", err.Error())
		return
	}

	if err := flattenIssuingCardholder(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingCardholder update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Company", "tax_id"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingCardholderResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *IssuingCardholderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandIssuingCardholderCreate(plan IssuingCardholderResourceModel) (*stripe.IssuingCardholderCreateParams, error) {
	params := &stripe.IssuingCardholderCreateParams{}

	if !plan.Billing.IsNull() && !plan.Billing.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Billing", plan.Billing) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing", params)
		}
	}
	if !plan.Company.IsNull() && !plan.Company.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Company", plan.Company) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "company", params)
		}
	}
	if !plan.Email.IsNull() && !plan.Email.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Email", "Email", plan.Email.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "email", params)
		}
	}
	if !plan.Individual.IsNull() && !plan.Individual.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Individual", plan.Individual) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "individual", params)
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
	if !plan.PhoneNumber.IsNull() && !plan.PhoneNumber.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PhoneNumber", "PhoneNumber", plan.PhoneNumber.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "phone_number", params)
		}
	}
	if !plan.PreferredLocales.IsNull() && !plan.PreferredLocales.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PreferredLocales", plan.PreferredLocales) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "preferred_locales", params)
		}
	}
	if !plan.SpendingControls.IsNull() && !plan.SpendingControls.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SpendingControls", plan.SpendingControls) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "spending_controls", params)
		}
	}
	if !plan.Status.IsNull() && !plan.Status.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Status", "Status", plan.Status.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "status", params)
		}
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}

	return params, nil
}

func expandIssuingCardholderUpdate(plan IssuingCardholderResourceModel, state IssuingCardholderResourceModel) (*stripe.IssuingCardholderUpdateParams, error) {
	params := &stripe.IssuingCardholderUpdateParams{}

	if !plan.Billing.Equal(state.Billing) && !plan.Billing.IsNull() && !plan.Billing.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Billing", plan.Billing) {
			if !plan.Billing.Equal(state.Billing) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "billing", params)
			}
		}
	}
	if !plan.Company.Equal(state.Company) && !plan.Company.IsNull() && !plan.Company.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Company", plan.Company) {
			if !plan.Company.Equal(state.Company) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "company", params)
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
	if !plan.Individual.Equal(state.Individual) && !plan.Individual.IsNull() && !plan.Individual.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Individual", plan.Individual) {
			if !plan.Individual.Equal(state.Individual) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "individual", params)
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
	if !plan.PhoneNumber.Equal(state.PhoneNumber) && !plan.PhoneNumber.IsNull() && !plan.PhoneNumber.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PhoneNumber", "PhoneNumber", plan.PhoneNumber.ValueString()) {
			if !plan.PhoneNumber.Equal(state.PhoneNumber) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "phone_number", params)
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
	if !plan.SpendingControls.Equal(state.SpendingControls) && !plan.SpendingControls.IsNull() && !plan.SpendingControls.IsUnknown() {
		if !assignAttrValueToNamedField(params, "SpendingControls", plan.SpendingControls) {
			if !plan.SpendingControls.Equal(state.SpendingControls) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "spending_controls", params)
			}
		}
	}
	if !plan.Status.Equal(state.Status) && !plan.Status.IsNull() && !plan.Status.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Status", "Status", plan.Status.ValueString()) {
			if !plan.Status.Equal(state.Status) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "status", params)
			}
		}
	}

	return params, nil
}

func flattenIssuingCardholder(obj *stripe.IssuingCardholder, state *IssuingCardholderResourceModel) error {
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
		assignedBilling := false
		hadRawBilling := false
		if rawValueBilling, rawOk := plainValueAtPath(raw, "billing"); rawOk {
			hadRawBilling = true
			if rawValueBilling != nil {
				sourceBilling := applyConfiguredKeyedListShapes(rawValueBilling, attrValueToPlain(state.Billing))
				if valueBilling, err := flattenPlainValue(sourceBilling, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}, "billing", "raw response"); err != nil {
					return err
				} else {
					if typedBilling, ok := valueBilling.(types.Object); ok {
						state.Billing = typedBilling
						assignedBilling = true
					}
				}
			}
		}
		if !assignedBilling {
			if !hasRaw {
				if responseValueBilling, ok := plainFromResponseField(obj, "Billing"); ok {
					sourceBilling := applyConfiguredKeyedListShapes(responseValueBilling, attrValueToPlain(state.Billing))
					if valueBilling, err := flattenPlainValue(
						sourceBilling,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}},
						"billing",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBilling, ok := valueBilling.(types.Object); ok {
							state.Billing = typedBilling
							assignedBilling = true
						}
					}
				}
			}
		}
		if !assignedBilling && hadRawBilling {
			if nullBilling, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}); ok {
				if typedBilling, ok := nullBilling.(types.Object); ok {
					state.Billing = typedBilling
				}
			}
		}
	}
	{
		assignedCompany := false
		hadRawCompany := false
		if rawValueCompany, rawOk := plainValueAtPath(raw, "company"); rawOk {
			hadRawCompany = true
			if rawValueCompany != nil {
				sourceCompany := applyConfiguredKeyedListShapes(rawValueCompany, attrValueToPlain(state.Company))
				if valueCompany, err := flattenPlainValue(sourceCompany, types.ObjectType{AttrTypes: map[string]attr.Type{"tax_id_provided": types.BoolType, "tax_id": types.StringType}}, "company", "raw response"); err != nil {
					return err
				} else {
					if typedCompany, ok := valueCompany.(types.Object); ok {
						state.Company = typedCompany
						assignedCompany = true
					}
				}
			}
		}
		if !assignedCompany {
			if !hasRaw {
				if responseValueCompany, ok := plainFromResponseField(obj, "Company"); ok {
					sourceCompany := applyConfiguredKeyedListShapes(responseValueCompany, attrValueToPlain(state.Company))
					if valueCompany, err := flattenPlainValue(
						sourceCompany,
						types.ObjectType{AttrTypes: map[string]attr.Type{"tax_id_provided": types.BoolType, "tax_id": types.StringType}},
						"company",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCompany, ok := valueCompany.(types.Object); ok {
							state.Company = typedCompany
							assignedCompany = true
						}
					}
				}
			}
		}
		if !assignedCompany && hadRawCompany {
			if nullCompany, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"tax_id_provided": types.BoolType, "tax_id": types.StringType}}); ok {
				if typedCompany, ok := nullCompany.(types.Object); ok {
					state.Company = typedCompany
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
		assignedIndividual := false
		hadRawIndividual := false
		if rawValueIndividual, rawOk := plainValueAtPath(raw, "individual"); rawOk {
			hadRawIndividual = true
			if rawValueIndividual != nil {
				sourceIndividual := applyConfiguredKeyedListShapes(rawValueIndividual, attrValueToPlain(state.Individual))
				if valueIndividual, err := flattenPlainValue(sourceIndividual, types.ObjectType{AttrTypes: map[string]attr.Type{"card_issuing": types.ObjectType{AttrTypes: map[string]attr.Type{"user_terms_acceptance": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.Int64Type, "ip": types.StringType, "user_agent": types.StringType}}}}, "dob": types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}, "first_name": types.StringType, "last_name": types.StringType, "verification": types.ObjectType{AttrTypes: map[string]attr.Type{"document": types.ObjectType{AttrTypes: map[string]attr.Type{"back": types.StringType, "front": types.StringType}}}}}}, "individual", "raw response"); err != nil {
					return err
				} else {
					if typedIndividual, ok := valueIndividual.(types.Object); ok {
						state.Individual = typedIndividual
						assignedIndividual = true
					}
				}
			}
		}
		if !assignedIndividual {
			if !hasRaw {
				if responseValueIndividual, ok := plainFromResponseField(obj, "Individual"); ok {
					sourceIndividual := applyConfiguredKeyedListShapes(responseValueIndividual, attrValueToPlain(state.Individual))
					if valueIndividual, err := flattenPlainValue(
						sourceIndividual,
						types.ObjectType{AttrTypes: map[string]attr.Type{"card_issuing": types.ObjectType{AttrTypes: map[string]attr.Type{"user_terms_acceptance": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.Int64Type, "ip": types.StringType, "user_agent": types.StringType}}}}, "dob": types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}, "first_name": types.StringType, "last_name": types.StringType, "verification": types.ObjectType{AttrTypes: map[string]attr.Type{"document": types.ObjectType{AttrTypes: map[string]attr.Type{"back": types.StringType, "front": types.StringType}}}}}},
						"individual",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedIndividual, ok := valueIndividual.(types.Object); ok {
							state.Individual = typedIndividual
							assignedIndividual = true
						}
					}
				}
			}
		}
		if !assignedIndividual && hadRawIndividual {
			if nullIndividual, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"card_issuing": types.ObjectType{AttrTypes: map[string]attr.Type{"user_terms_acceptance": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.Int64Type, "ip": types.StringType, "user_agent": types.StringType}}}}, "dob": types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}, "first_name": types.StringType, "last_name": types.StringType, "verification": types.ObjectType{AttrTypes: map[string]attr.Type{"document": types.ObjectType{AttrTypes: map[string]attr.Type{"back": types.StringType, "front": types.StringType}}}}}}); ok {
				if typedIndividual, ok := nullIndividual.(types.Object); ok {
					state.Individual = typedIndividual
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
		if rawValuePhoneNumber, rawOk := plainValueAtPath(raw, "phone_number"); rawOk {
			if valuePhoneNumber, err := flattenPlainValue(rawValuePhoneNumber, types.StringType, "phone_number", "raw response"); err != nil {
				return err
			} else {
				if typedPhoneNumber, ok := valuePhoneNumber.(types.String); ok {
					state.PhoneNumber = typedPhoneNumber
				}
			}
		} else if !hasRaw {
			if responseValuePhoneNumber, ok := plainFromResponseField(obj, "PhoneNumber"); ok {
				if valuePhoneNumber, err := flattenPlainValue(responseValuePhoneNumber, types.StringType, "phone_number", "response struct"); err != nil {
					return err
				} else {
					if typedPhoneNumber, ok := valuePhoneNumber.(types.String); ok {
						state.PhoneNumber = typedPhoneNumber
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
		assignedRequirements := false
		hadRawRequirements := false
		if rawValueRequirements, rawOk := plainValueAtPath(raw, "requirements"); rawOk {
			hadRawRequirements = true
			if rawValueRequirements != nil {
				sourceRequirements := applyConfiguredKeyedListShapes(rawValueRequirements, attrValueToPlain(state.Requirements))
				if valueRequirements, err := flattenPlainValue(sourceRequirements, types.ObjectType{AttrTypes: map[string]attr.Type{"past_due": types.ListType{ElemType: types.StringType}}}, "requirements", "raw response"); err != nil {
					return err
				} else {
					if typedRequirements, ok := valueRequirements.(types.Object); ok {
						state.Requirements = typedRequirements
						assignedRequirements = true
					}
				}
			}
		}
		if !assignedRequirements {
			if !hasRaw {
				if responseValueRequirements, ok := plainFromResponseField(obj, "Requirements"); ok {
					sourceRequirements := applyConfiguredKeyedListShapes(responseValueRequirements, attrValueToPlain(state.Requirements))
					if valueRequirements, err := flattenPlainValue(
						sourceRequirements,
						types.ObjectType{AttrTypes: map[string]attr.Type{"past_due": types.ListType{ElemType: types.StringType}}},
						"requirements",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRequirements, ok := valueRequirements.(types.Object); ok {
							state.Requirements = typedRequirements
							assignedRequirements = true
						}
					}
				}
			}
		}
		if !assignedRequirements && hadRawRequirements {
			if nullRequirements, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"past_due": types.ListType{ElemType: types.StringType}}}); ok {
				if typedRequirements, ok := nullRequirements.(types.Object); ok {
					state.Requirements = typedRequirements
				}
			}
		}
	}
	{
		assignedSpendingControls := false
		hadRawSpendingControls := false
		if rawValueSpendingControls, rawOk := plainValueAtPath(raw, "spending_controls"); rawOk {
			hadRawSpendingControls = true
			if rawValueSpendingControls != nil {
				sourceSpendingControls := applyConfiguredKeyedListShapes(rawValueSpendingControls, attrValueToPlain(state.SpendingControls))
				if valueSpendingControls, err := flattenPlainValue(sourceSpendingControls, types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_card_presences": types.ListType{ElemType: types.StringType}, "allowed_categories": types.ListType{ElemType: types.StringType}, "allowed_merchant_countries": types.ListType{ElemType: types.StringType}, "blocked_card_presences": types.ListType{ElemType: types.StringType}, "blocked_categories": types.ListType{ElemType: types.StringType}, "blocked_merchant_countries": types.ListType{ElemType: types.StringType}, "spending_limits": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "categories": types.ListType{ElemType: types.StringType}, "interval": types.StringType}}}, "spending_limits_currency": types.StringType}}, "spending_controls", "raw response"); err != nil {
					return err
				} else {
					if typedSpendingControls, ok := valueSpendingControls.(types.Object); ok {
						state.SpendingControls = typedSpendingControls
						assignedSpendingControls = true
					}
				}
			}
		}
		if !assignedSpendingControls {
			if !hasRaw {
				if responseValueSpendingControls, ok := plainFromResponseField(obj, "SpendingControls"); ok {
					sourceSpendingControls := applyConfiguredKeyedListShapes(responseValueSpendingControls, attrValueToPlain(state.SpendingControls))
					if valueSpendingControls, err := flattenPlainValue(
						sourceSpendingControls,
						types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_card_presences": types.ListType{ElemType: types.StringType}, "allowed_categories": types.ListType{ElemType: types.StringType}, "allowed_merchant_countries": types.ListType{ElemType: types.StringType}, "blocked_card_presences": types.ListType{ElemType: types.StringType}, "blocked_categories": types.ListType{ElemType: types.StringType}, "blocked_merchant_countries": types.ListType{ElemType: types.StringType}, "spending_limits": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "categories": types.ListType{ElemType: types.StringType}, "interval": types.StringType}}}, "spending_limits_currency": types.StringType}},
						"spending_controls",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedSpendingControls, ok := valueSpendingControls.(types.Object); ok {
							state.SpendingControls = typedSpendingControls
							assignedSpendingControls = true
						}
					}
				}
			}
		}
		if !assignedSpendingControls && hadRawSpendingControls {
			if nullSpendingControls, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_card_presences": types.ListType{ElemType: types.StringType}, "allowed_categories": types.ListType{ElemType: types.StringType}, "allowed_merchant_countries": types.ListType{ElemType: types.StringType}, "blocked_card_presences": types.ListType{ElemType: types.StringType}, "blocked_categories": types.ListType{ElemType: types.StringType}, "blocked_merchant_countries": types.ListType{ElemType: types.StringType}, "spending_limits": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "categories": types.ListType{ElemType: types.StringType}, "interval": types.StringType}}}, "spending_limits_currency": types.StringType}}); ok {
				if typedSpendingControls, ok := nullSpendingControls.(types.Object); ok {
					state.SpendingControls = typedSpendingControls
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
	return nil
}
