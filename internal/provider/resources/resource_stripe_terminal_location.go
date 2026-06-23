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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &TerminalLocationResource{}

var _ resource.ResourceWithConfigure = &TerminalLocationResource{}

var _ resource.ResourceWithImportState = &TerminalLocationResource{}

func NewTerminalLocationResource() resource.Resource {
	return &TerminalLocationResource{}
}

type TerminalLocationResource struct {
	client *stripe.Client
}

type TerminalLocationResourceModel struct {
	Object                 types.String `tfsdk:"object"`
	Address                types.Object `tfsdk:"address"`
	AddressKana            types.Object `tfsdk:"address_kana"`
	ConfigurationOverrides types.String `tfsdk:"configuration_overrides"`
	DisplayName            types.String `tfsdk:"display_name"`
	DisplayNameKana        types.String `tfsdk:"display_name_kana"`
	DisplayNameKanji       types.String `tfsdk:"display_name_kanji"`
	ID                     types.String `tfsdk:"id"`
	Livemode               types.Bool   `tfsdk:"livemode"`
	Metadata               types.Map    `tfsdk:"metadata"`
	Phone                  types.String `tfsdk:"phone"`
}

func (r *TerminalLocationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TerminalLocationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_terminal_location"
}

func (r *TerminalLocationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Location represents a grouping of readers.\n\nRelated guide: [Fleet management](https://docs.stripe.com/terminal/fleet/locations)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("terminal.location")},
			},
			"address": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"city": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "City, district, suburb, town, or village.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Required:    true,
						Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
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
			"address_kana": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"city": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "City/Ward.",
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
						Description:   "Block/Building number.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"line2": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Building details.",
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
						Description:   "Prefecture.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"town": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Town/cho-me.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"configuration_overrides": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ID of a configuration that will be used to customize all readers in this location.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"display_name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The display name of the location.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"display_name_kana": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The Kana variation of the display name of the location.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"display_name_kanji": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The Kanji variation of the display name of the location.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"phone": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The phone number of the location.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *TerminalLocationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TerminalLocationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTerminalLocationCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building TerminalLocation create params", err.Error())
		return
	}

	obj, err := r.client.V1TerminalLocations.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating TerminalLocation", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TerminalLocations.B, r.client.V1TerminalLocations.Key, stripe.FormatURLPath("/v1/terminal/locations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TerminalLocation create raw response", err.Error())
		return
	}

	if err := flattenTerminalLocation(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TerminalLocation create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TerminalLocationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState TerminalLocationResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TerminalLocationResourceModel
	state = priorState

	obj, err := r.client.V1TerminalLocations.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading TerminalLocation", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TerminalLocations.B, r.client.V1TerminalLocations.Key, stripe.FormatURLPath("/v1/terminal/locations/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TerminalLocation raw response", err.Error())
		return
	}

	if err := flattenTerminalLocation(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening TerminalLocation read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *TerminalLocationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TerminalLocationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TerminalLocationResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandTerminalLocationUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building TerminalLocation update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building TerminalLocation update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1TerminalLocations.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating TerminalLocation", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TerminalLocations.B, r.client.V1TerminalLocations.Key, stripe.FormatURLPath("/v1/terminal/locations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TerminalLocation update raw response", err.Error())
		return
	}

	if err := flattenTerminalLocation(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TerminalLocation update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TerminalLocationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TerminalLocationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1TerminalLocations.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting TerminalLocation", err.Error())
		return
	}
}

func (r *TerminalLocationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandTerminalLocationCreate(plan TerminalLocationResourceModel) (*stripe.TerminalLocationCreateParams, error) {
	params := &stripe.TerminalLocationCreateParams{}

	if !plan.Address.IsNull() && !plan.Address.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Address", plan.Address) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "address", params)
		}
	}
	if !plan.AddressKana.IsNull() && !plan.AddressKana.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AddressKana", plan.AddressKana) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "address_kana", params)
		}
	}
	if !plan.ConfigurationOverrides.IsNull() && !plan.ConfigurationOverrides.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ConfigurationOverrides", "ConfigurationOverrides", plan.ConfigurationOverrides.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "configuration_overrides", params)
		}
	}
	if !plan.DisplayName.IsNull() && !plan.DisplayName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayName", "DisplayName", plan.DisplayName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "display_name", params)
		}
	}
	if !plan.DisplayNameKana.IsNull() && !plan.DisplayNameKana.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayNameKana", "DisplayNameKana", plan.DisplayNameKana.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "display_name_kana", params)
		}
	}
	if !plan.DisplayNameKanji.IsNull() && !plan.DisplayNameKanji.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayNameKanji", "DisplayNameKanji", plan.DisplayNameKanji.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "display_name_kanji", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Phone.IsNull() && !plan.Phone.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Phone", "Phone", plan.Phone.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "phone", params)
		}
	}

	return params, nil
}

func expandTerminalLocationUpdate(plan TerminalLocationResourceModel, state TerminalLocationResourceModel) (*stripe.TerminalLocationUpdateParams, error) {
	params := &stripe.TerminalLocationUpdateParams{}

	if !plan.Address.Equal(state.Address) && !plan.Address.IsNull() && !plan.Address.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Address", plan.Address) {
			if !plan.Address.Equal(state.Address) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "address", params)
			}
		}
	}
	if !plan.AddressKana.Equal(state.AddressKana) && !plan.AddressKana.IsNull() && !plan.AddressKana.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AddressKana", plan.AddressKana) {
			if !plan.AddressKana.Equal(state.AddressKana) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "address_kana", params)
			}
		}
	}
	if !plan.ConfigurationOverrides.Equal(state.ConfigurationOverrides) && !plan.ConfigurationOverrides.IsNull() && !plan.ConfigurationOverrides.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ConfigurationOverrides", "ConfigurationOverrides", plan.ConfigurationOverrides.ValueString()) {
			if !plan.ConfigurationOverrides.Equal(state.ConfigurationOverrides) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "configuration_overrides", params)
			}
		}
	}
	if !plan.DisplayName.Equal(state.DisplayName) && !plan.DisplayName.IsNull() && !plan.DisplayName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayName", "DisplayName", plan.DisplayName.ValueString()) {
			if !plan.DisplayName.Equal(state.DisplayName) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "display_name", params)
			}
		}
	}
	if !plan.DisplayNameKana.Equal(state.DisplayNameKana) && !plan.DisplayNameKana.IsNull() && !plan.DisplayNameKana.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayNameKana", "DisplayNameKana", plan.DisplayNameKana.ValueString()) {
			if !plan.DisplayNameKana.Equal(state.DisplayNameKana) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "display_name_kana", params)
			}
		}
	}
	if !plan.DisplayNameKanji.Equal(state.DisplayNameKanji) && !plan.DisplayNameKanji.IsNull() && !plan.DisplayNameKanji.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DisplayNameKanji", "DisplayNameKanji", plan.DisplayNameKanji.ValueString()) {
			if !plan.DisplayNameKanji.Equal(state.DisplayNameKanji) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "display_name_kanji", params)
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
	if !plan.Phone.Equal(state.Phone) && !plan.Phone.IsNull() && !plan.Phone.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Phone", "Phone", plan.Phone.ValueString()) {
			if !plan.Phone.Equal(state.Phone) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "phone", params)
			}
		}
	}

	return params, nil
}

func flattenTerminalLocation(obj *stripe.TerminalLocation, state *TerminalLocationResourceModel) error {
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
				sourceAddress := applyConfiguredKeyedListShapes(rawValueAddress, attrValueToPlain(state.Address))
				if valueAddress, err := flattenPlainValue(sourceAddress, types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "address", "raw response"); err != nil {
					return err
				} else {
					if typedAddress, ok := valueAddress.(types.Object); ok {
						state.Address = typedAddress
						assignedAddress = true
					}
				}
			}
		}
		if !assignedAddress {
			if !hasRaw {
				if responseValueAddress, ok := plainFromResponseField(obj, "Address"); ok {
					sourceAddress := applyConfiguredKeyedListShapes(responseValueAddress, attrValueToPlain(state.Address))
					if valueAddress, err := flattenPlainValue(
						sourceAddress,
						types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}},
						"address",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAddress, ok := valueAddress.(types.Object); ok {
							state.Address = typedAddress
							assignedAddress = true
						}
					}
				}
			}
		}
		if !assignedAddress && hadRawAddress {
			if nullAddress, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}); ok {
				if typedAddress, ok := nullAddress.(types.Object); ok {
					state.Address = typedAddress
				}
			}
		}
	}
	{
		assignedAddressKana := false
		hadRawAddressKana := false
		if rawValueAddressKana, rawOk := plainValueAtPath(raw, "address_kana"); rawOk {
			hadRawAddressKana = true
			if rawValueAddressKana != nil {
				sourceAddressKana := applyConfiguredKeyedListShapes(rawValueAddressKana, attrValueToPlain(state.AddressKana))
				if valueAddressKana, err := flattenPlainValue(sourceAddressKana, types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}}, "address_kana", "raw response"); err != nil {
					return err
				} else {
					if typedAddressKana, ok := valueAddressKana.(types.Object); ok {
						state.AddressKana = typedAddressKana
						assignedAddressKana = true
					}
				}
			}
		}
		if !assignedAddressKana {
			if !hasRaw {
				if responseValueAddressKana, ok := plainFromResponseField(obj, "AddressKana"); ok {
					sourceAddressKana := applyConfiguredKeyedListShapes(responseValueAddressKana, attrValueToPlain(state.AddressKana))
					if valueAddressKana, err := flattenPlainValue(
						sourceAddressKana,
						types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}},
						"address_kana",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAddressKana, ok := valueAddressKana.(types.Object); ok {
							state.AddressKana = typedAddressKana
							assignedAddressKana = true
						}
					}
				}
			}
		}
		if !assignedAddressKana && hadRawAddressKana {
			if nullAddressKana, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}}); ok {
				if typedAddressKana, ok := nullAddressKana.(types.Object); ok {
					state.AddressKana = typedAddressKana
				}
			}
		}
	}
	{
		if rawValueConfigurationOverrides, rawOk := plainValueAtPath(raw, "configuration_overrides"); rawOk {
			if valueConfigurationOverrides, err := flattenPlainValue(rawValueConfigurationOverrides, types.StringType, "configuration_overrides", "raw response"); err != nil {
				return err
			} else {
				if typedConfigurationOverrides, ok := valueConfigurationOverrides.(types.String); ok {
					state.ConfigurationOverrides = typedConfigurationOverrides
				}
			}
		} else if !hasRaw {
			if responseValueConfigurationOverrides, ok := plainFromResponseField(obj, "ConfigurationOverrides"); ok {
				if valueConfigurationOverrides, err := flattenPlainValue(responseValueConfigurationOverrides, types.StringType, "configuration_overrides", "response struct"); err != nil {
					return err
				} else {
					if typedConfigurationOverrides, ok := valueConfigurationOverrides.(types.String); ok {
						state.ConfigurationOverrides = typedConfigurationOverrides
					}
				}
			}
		}
	}
	{
		if rawValueDisplayName, rawOk := plainValueAtPath(raw, "display_name"); rawOk {
			if valueDisplayName, err := flattenPlainValue(rawValueDisplayName, types.StringType, "display_name", "raw response"); err != nil {
				return err
			} else {
				if typedDisplayName, ok := valueDisplayName.(types.String); ok {
					state.DisplayName = typedDisplayName
				}
			}
		} else if !hasRaw {
			if responseValueDisplayName, ok := plainFromResponseField(obj, "DisplayName"); ok {
				if valueDisplayName, err := flattenPlainValue(responseValueDisplayName, types.StringType, "display_name", "response struct"); err != nil {
					return err
				} else {
					if typedDisplayName, ok := valueDisplayName.(types.String); ok {
						state.DisplayName = typedDisplayName
					}
				}
			}
		}
	}
	{
		if rawValueDisplayNameKana, rawOk := plainValueAtPath(raw, "display_name_kana"); rawOk {
			if valueDisplayNameKana, err := flattenPlainValue(rawValueDisplayNameKana, types.StringType, "display_name_kana", "raw response"); err != nil {
				return err
			} else {
				if typedDisplayNameKana, ok := valueDisplayNameKana.(types.String); ok {
					state.DisplayNameKana = typedDisplayNameKana
				}
			}
		} else if !hasRaw {
			if responseValueDisplayNameKana, ok := plainFromResponseField(obj, "DisplayNameKana"); ok {
				if valueDisplayNameKana, err := flattenPlainValue(responseValueDisplayNameKana, types.StringType, "display_name_kana", "response struct"); err != nil {
					return err
				} else {
					if typedDisplayNameKana, ok := valueDisplayNameKana.(types.String); ok {
						state.DisplayNameKana = typedDisplayNameKana
					}
				}
			}
		}
	}
	{
		if rawValueDisplayNameKanji, rawOk := plainValueAtPath(raw, "display_name_kanji"); rawOk {
			if valueDisplayNameKanji, err := flattenPlainValue(rawValueDisplayNameKanji, types.StringType, "display_name_kanji", "raw response"); err != nil {
				return err
			} else {
				if typedDisplayNameKanji, ok := valueDisplayNameKanji.(types.String); ok {
					state.DisplayNameKanji = typedDisplayNameKanji
				}
			}
		} else if !hasRaw {
			if responseValueDisplayNameKanji, ok := plainFromResponseField(obj, "DisplayNameKanji"); ok {
				if valueDisplayNameKanji, err := flattenPlainValue(responseValueDisplayNameKanji, types.StringType, "display_name_kanji", "response struct"); err != nil {
					return err
				} else {
					if typedDisplayNameKanji, ok := valueDisplayNameKanji.(types.String); ok {
						state.DisplayNameKanji = typedDisplayNameKanji
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
	return nil
}
