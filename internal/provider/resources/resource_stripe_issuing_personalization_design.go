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

var _ resource.Resource = &IssuingPersonalizationDesignResource{}

var _ resource.ResourceWithConfigure = &IssuingPersonalizationDesignResource{}

var _ resource.ResourceWithImportState = &IssuingPersonalizationDesignResource{}

func NewIssuingPersonalizationDesignResource() resource.Resource {
	return &IssuingPersonalizationDesignResource{}
}

type IssuingPersonalizationDesignResource struct {
	client *stripe.Client
}

type IssuingPersonalizationDesignResourceModel struct {
	Object            types.String `tfsdk:"object"`
	CardLogo          types.String `tfsdk:"card_logo"`
	CarrierText       types.Object `tfsdk:"carrier_text"`
	Created           types.Int64  `tfsdk:"created"`
	ID                types.String `tfsdk:"id"`
	Livemode          types.Bool   `tfsdk:"livemode"`
	LookupKey         types.String `tfsdk:"lookup_key"`
	Metadata          types.Map    `tfsdk:"metadata"`
	Name              types.String `tfsdk:"name"`
	PhysicalBundle    types.String `tfsdk:"physical_bundle"`
	Preferences       types.Object `tfsdk:"preferences"`
	RejectionReasons  types.Object `tfsdk:"rejection_reasons"`
	Status            types.String `tfsdk:"status"`
	TransferLookupKey types.Bool   `tfsdk:"transfer_lookup_key"`
}

func (r *IssuingPersonalizationDesignResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IssuingPersonalizationDesignResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_issuing_personalization_design"
}

func (r *IssuingPersonalizationDesignResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Personalization Design is a logical grouping of a Physical Bundle, card logo, and carrier text that represents a product line.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("issuing.personalization_design")},
			},
			"card_logo": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The file for the card logo to use with physical bundles that support card logos. Must have a `purpose` value of `issuing_logo`. Image must be in PNG format with dimensions of 1000px by 200px. It must be a binary (black and white) image containing a black logo on a white background. We don't accept grayscale.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"carrier_text": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Hash containing carrier text, for use with physical bundles that support carrier text.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"footer_body": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The footer body text of the carrier letter.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"footer_title": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The footer title text of the carrier letter.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"header_body": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The header body text of the carrier letter.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"header_title": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The header title text of the carrier letter.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
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
			"lookup_key": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A lookup key used to retrieve personalization designs dynamically from a static string. This may be up to 200 characters.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
				Description:   "Friendly display name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"physical_bundle": schema.StringAttribute{
				Required:    true,
				Description: "The physical bundle object belonging to this personalization design.",
			},
			"preferences": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"is_default": schema.BoolAttribute{
						Required:    true,
						Description: "Whether we use this personalization design to create cards when one isn't specified. A connected account uses the Connect platform's default design if no personalization design is set as the default design.",
					},
					"is_platform_default": schema.BoolAttribute{
						Computed:      true,
						Description:   "Whether this personalization design is used to create cards when one is not specified and a default for this connected account does not exist.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"rejection_reasons": schema.SingleNestedAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"card_logo": schema.ListAttribute{
						Computed:      true,
						Description:   "The reason(s) the card logo was rejected.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"carrier_text": schema.ListAttribute{
						Computed:      true,
						Description:   "The reason(s) the carrier text was rejected.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
				},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Whether this personalization design can be used to create cards.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "inactive", "rejected", "review")},
			},
			"transfer_lookup_key": schema.BoolAttribute{
				Optional:    true,
				Description: "If set to true, will atomically remove the lookup key from the existing personalization design, and assign it to this personalization design.",
				WriteOnly:   true,
			},
		},
	}
}

func (r *IssuingPersonalizationDesignResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IssuingPersonalizationDesignResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config IssuingPersonalizationDesignResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"TransferLookupKey"}})

	params, err := expandIssuingPersonalizationDesignCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building IssuingPersonalizationDesign create params", err.Error())
		return
	}

	obj, err := r.client.V1IssuingPersonalizationDesigns.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating IssuingPersonalizationDesign", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingPersonalizationDesigns.B, r.client.V1IssuingPersonalizationDesigns.Key, stripe.FormatURLPath("/v1/issuing/personalization_designs/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingPersonalizationDesign create raw response", err.Error())
		return
	}

	if err := flattenIssuingPersonalizationDesign(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingPersonalizationDesign create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"TransferLookupKey"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingPersonalizationDesignResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState IssuingPersonalizationDesignResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state IssuingPersonalizationDesignResourceModel
	state = priorState

	obj, err := r.client.V1IssuingPersonalizationDesigns.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading IssuingPersonalizationDesign", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingPersonalizationDesigns.B, r.client.V1IssuingPersonalizationDesigns.Key, stripe.FormatURLPath("/v1/issuing/personalization_designs/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingPersonalizationDesign raw response", err.Error())
		return
	}

	if err := flattenIssuingPersonalizationDesign(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingPersonalizationDesign read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"TransferLookupKey"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingPersonalizationDesignResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IssuingPersonalizationDesignResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config IssuingPersonalizationDesignResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"TransferLookupKey"}})

	var state IssuingPersonalizationDesignResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandIssuingPersonalizationDesignUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building IssuingPersonalizationDesign update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building IssuingPersonalizationDesign update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1IssuingPersonalizationDesigns.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating IssuingPersonalizationDesign", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingPersonalizationDesigns.B, r.client.V1IssuingPersonalizationDesigns.Key, stripe.FormatURLPath("/v1/issuing/personalization_designs/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingPersonalizationDesign update raw response", err.Error())
		return
	}

	if err := flattenIssuingPersonalizationDesign(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingPersonalizationDesign update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"TransferLookupKey"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingPersonalizationDesignResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *IssuingPersonalizationDesignResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandIssuingPersonalizationDesignCreate(plan IssuingPersonalizationDesignResourceModel) (*stripe.IssuingPersonalizationDesignCreateParams, error) {
	params := &stripe.IssuingPersonalizationDesignCreateParams{}

	if !plan.CardLogo.IsNull() && !plan.CardLogo.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CardLogoID", "CardLogo", plan.CardLogo.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "card_logo", params)
		}
	}
	if !plan.CarrierText.IsNull() && !plan.CarrierText.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CarrierText", plan.CarrierText) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "carrier_text", params)
		}
	}
	if !plan.LookupKey.IsNull() && !plan.LookupKey.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LookupKey", "LookupKey", plan.LookupKey.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "lookup_key", params)
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
	if !plan.PhysicalBundle.IsNull() && !plan.PhysicalBundle.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PhysicalBundleID", "PhysicalBundle", plan.PhysicalBundle.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "physical_bundle", params)
		}
	}
	if !plan.Preferences.IsNull() && !plan.Preferences.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Preferences", plan.Preferences) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "preferences", params)
		}
	}
	if !plan.TransferLookupKey.IsNull() && !plan.TransferLookupKey.IsUnknown() {
		params.TransferLookupKey = stripe.Bool(plan.TransferLookupKey.ValueBool())
	}

	return params, nil
}

func expandIssuingPersonalizationDesignUpdate(plan IssuingPersonalizationDesignResourceModel, state IssuingPersonalizationDesignResourceModel) (*stripe.IssuingPersonalizationDesignUpdateParams, error) {
	params := &stripe.IssuingPersonalizationDesignUpdateParams{}

	if !plan.CardLogo.Equal(state.CardLogo) && !plan.CardLogo.IsNull() && !plan.CardLogo.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CardLogoID", "CardLogo", plan.CardLogo.ValueString()) {
			if !plan.CardLogo.Equal(state.CardLogo) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "card_logo", params)
			}
		}
	}
	if !plan.CarrierText.Equal(state.CarrierText) && !plan.CarrierText.IsNull() && !plan.CarrierText.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CarrierText", plan.CarrierText) {
			if !plan.CarrierText.Equal(state.CarrierText) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "carrier_text", params)
			}
		}
	}
	if !plan.LookupKey.Equal(state.LookupKey) && !plan.LookupKey.IsNull() && !plan.LookupKey.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LookupKey", "LookupKey", plan.LookupKey.ValueString()) {
			if !plan.LookupKey.Equal(state.LookupKey) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "lookup_key", params)
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
	if !plan.PhysicalBundle.Equal(state.PhysicalBundle) && !plan.PhysicalBundle.IsNull() && !plan.PhysicalBundle.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PhysicalBundleID", "PhysicalBundle", plan.PhysicalBundle.ValueString()) {
			if !plan.PhysicalBundle.Equal(state.PhysicalBundle) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "physical_bundle", params)
			}
		}
	}
	if !plan.Preferences.Equal(state.Preferences) && !plan.Preferences.IsNull() && !plan.Preferences.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Preferences", plan.Preferences) {
			if !plan.Preferences.Equal(state.Preferences) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "preferences", params)
			}
		}
	}
	if !plan.TransferLookupKey.Equal(state.TransferLookupKey) && !plan.TransferLookupKey.IsNull() && !plan.TransferLookupKey.IsUnknown() {
		params.TransferLookupKey = stripe.Bool(plan.TransferLookupKey.ValueBool())
	}

	return params, nil
}

func flattenIssuingPersonalizationDesign(obj *stripe.IssuingPersonalizationDesign, state *IssuingPersonalizationDesignResourceModel) error {
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
			if rawValueCardLogo, rawOk := plainValueAtPath(raw, "card_logo"); rawOk {
				if typedCardLogo, ok := plainToStringIDValue(rawValueCardLogo); ok {
					state.CardLogo = typedCardLogo
				}
			} else if !hasRaw {
				if responseValueCardLogo, ok := plainFromResponseField(obj, "CardLogo"); ok {
					if typedCardLogo, ok := plainToStringIDValue(responseValueCardLogo); ok {
						state.CardLogo = typedCardLogo
					}
				}
			}
		}
	}
	{
		assignedCarrierText := false
		hadRawCarrierText := false
		if rawValueCarrierText, rawOk := plainValueAtPath(raw, "carrier_text"); rawOk {
			hadRawCarrierText = true
			if rawValueCarrierText != nil {
				sourceCarrierText := applyConfiguredKeyedListShapes(rawValueCarrierText, attrValueToPlain(state.CarrierText))
				if valueCarrierText, err := flattenPlainValue(sourceCarrierText, types.ObjectType{AttrTypes: map[string]attr.Type{"footer_body": types.StringType, "footer_title": types.StringType, "header_body": types.StringType, "header_title": types.StringType}}, "carrier_text", "raw response"); err != nil {
					return err
				} else {
					if typedCarrierText, ok := valueCarrierText.(types.Object); ok {
						state.CarrierText = typedCarrierText
						assignedCarrierText = true
					}
				}
			}
		}
		if !assignedCarrierText {
			if !hasRaw {
				if responseValueCarrierText, ok := plainFromResponseField(obj, "CarrierText"); ok {
					sourceCarrierText := applyConfiguredKeyedListShapes(responseValueCarrierText, attrValueToPlain(state.CarrierText))
					if valueCarrierText, err := flattenPlainValue(
						sourceCarrierText,
						types.ObjectType{AttrTypes: map[string]attr.Type{"footer_body": types.StringType, "footer_title": types.StringType, "header_body": types.StringType, "header_title": types.StringType}},
						"carrier_text",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCarrierText, ok := valueCarrierText.(types.Object); ok {
							state.CarrierText = typedCarrierText
							assignedCarrierText = true
						}
					}
				}
			}
		}
		if !assignedCarrierText && hadRawCarrierText {
			if nullCarrierText, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"footer_body": types.StringType, "footer_title": types.StringType, "header_body": types.StringType, "header_title": types.StringType}}); ok {
				if typedCarrierText, ok := nullCarrierText.(types.Object); ok {
					state.CarrierText = typedCarrierText
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
		if rawValueLookupKey, rawOk := plainValueAtPath(raw, "lookup_key"); rawOk {
			if valueLookupKey, err := flattenPlainValue(rawValueLookupKey, types.StringType, "lookup_key", "raw response"); err != nil {
				return err
			} else {
				if typedLookupKey, ok := valueLookupKey.(types.String); ok {
					state.LookupKey = typedLookupKey
				}
			}
		} else if !hasRaw {
			if responseValueLookupKey, ok := plainFromResponseField(obj, "LookupKey"); ok {
				if valueLookupKey, err := flattenPlainValue(responseValueLookupKey, types.StringType, "lookup_key", "response struct"); err != nil {
					return err
				} else {
					if typedLookupKey, ok := valueLookupKey.(types.String); ok {
						state.LookupKey = typedLookupKey
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
		if true {
			if rawValuePhysicalBundle, rawOk := plainValueAtPath(raw, "physical_bundle"); rawOk {
				if typedPhysicalBundle, ok := plainToStringIDValue(rawValuePhysicalBundle); ok {
					state.PhysicalBundle = typedPhysicalBundle
				}
			} else if !hasRaw {
				if responseValuePhysicalBundle, ok := plainFromResponseField(obj, "PhysicalBundle"); ok {
					if typedPhysicalBundle, ok := plainToStringIDValue(responseValuePhysicalBundle); ok {
						state.PhysicalBundle = typedPhysicalBundle
					}
				}
			}
		}
	}
	{
		assignedPreferences := false
		hadRawPreferences := false
		if rawValuePreferences, rawOk := plainValueAtPath(raw, "preferences"); rawOk {
			hadRawPreferences = true
			if rawValuePreferences != nil {
				sourcePreferences := applyConfiguredKeyedListShapes(rawValuePreferences, attrValueToPlain(state.Preferences))
				if valuePreferences, err := flattenPlainValue(sourcePreferences, types.ObjectType{AttrTypes: map[string]attr.Type{"is_default": types.BoolType, "is_platform_default": types.BoolType}}, "preferences", "raw response"); err != nil {
					return err
				} else {
					if typedPreferences, ok := valuePreferences.(types.Object); ok {
						state.Preferences = typedPreferences
						assignedPreferences = true
					}
				}
			}
		}
		if !assignedPreferences {
			if !hasRaw {
				if responseValuePreferences, ok := plainFromResponseField(obj, "Preferences"); ok {
					sourcePreferences := applyConfiguredKeyedListShapes(responseValuePreferences, attrValueToPlain(state.Preferences))
					if valuePreferences, err := flattenPlainValue(
						sourcePreferences,
						types.ObjectType{AttrTypes: map[string]attr.Type{"is_default": types.BoolType, "is_platform_default": types.BoolType}},
						"preferences",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPreferences, ok := valuePreferences.(types.Object); ok {
							state.Preferences = typedPreferences
							assignedPreferences = true
						}
					}
				}
			}
		}
		if !assignedPreferences && hadRawPreferences {
			if nullPreferences, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"is_default": types.BoolType, "is_platform_default": types.BoolType}}); ok {
				if typedPreferences, ok := nullPreferences.(types.Object); ok {
					state.Preferences = typedPreferences
				}
			}
		}
	}
	{
		assignedRejectionReasons := false
		hadRawRejectionReasons := false
		if rawValueRejectionReasons, rawOk := plainValueAtPath(raw, "rejection_reasons"); rawOk {
			hadRawRejectionReasons = true
			if rawValueRejectionReasons != nil {
				sourceRejectionReasons := applyConfiguredKeyedListShapes(rawValueRejectionReasons, attrValueToPlain(state.RejectionReasons))
				if valueRejectionReasons, err := flattenPlainValue(sourceRejectionReasons, types.ObjectType{AttrTypes: map[string]attr.Type{"card_logo": types.ListType{ElemType: types.StringType}, "carrier_text": types.ListType{ElemType: types.StringType}}}, "rejection_reasons", "raw response"); err != nil {
					return err
				} else {
					if typedRejectionReasons, ok := valueRejectionReasons.(types.Object); ok {
						state.RejectionReasons = typedRejectionReasons
						assignedRejectionReasons = true
					}
				}
			}
		}
		if !assignedRejectionReasons {
			if !hasRaw {
				if responseValueRejectionReasons, ok := plainFromResponseField(obj, "RejectionReasons"); ok {
					sourceRejectionReasons := applyConfiguredKeyedListShapes(responseValueRejectionReasons, attrValueToPlain(state.RejectionReasons))
					if valueRejectionReasons, err := flattenPlainValue(
						sourceRejectionReasons,
						types.ObjectType{AttrTypes: map[string]attr.Type{"card_logo": types.ListType{ElemType: types.StringType}, "carrier_text": types.ListType{ElemType: types.StringType}}},
						"rejection_reasons",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRejectionReasons, ok := valueRejectionReasons.(types.Object); ok {
							state.RejectionReasons = typedRejectionReasons
							assignedRejectionReasons = true
						}
					}
				}
			}
		}
		if !assignedRejectionReasons && hadRawRejectionReasons {
			if nullRejectionReasons, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"card_logo": types.ListType{ElemType: types.StringType}, "carrier_text": types.ListType{ElemType: types.StringType}}}); ok {
				if typedRejectionReasons, ok := nullRejectionReasons.(types.Object); ok {
					state.RejectionReasons = typedRejectionReasons
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
