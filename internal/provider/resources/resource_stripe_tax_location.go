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

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &TaxLocationResource{}

var _ resource.ResourceWithConfigure = &TaxLocationResource{}

var _ resource.ResourceWithImportState = &TaxLocationResource{}

func NewTaxLocationResource() resource.Resource {
	return &TaxLocationResource{}
}

type TaxLocationResource struct {
	client *stripe.Client
}

type TaxLocationResourceModel struct {
	Object      types.String `tfsdk:"object"`
	Address     types.Object `tfsdk:"address"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Livemode    types.Bool   `tfsdk:"livemode"`
	Type        types.String `tfsdk:"type"`
}

func (r *TaxLocationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TaxLocationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tax_location"
}

func (r *TaxLocationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Tax locations represent venues for services, tickets, or other product types.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("tax.location")},
			},
			"address": schema.SingleNestedAttribute{
				Required: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"city": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "City, district, suburb, town, or village.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"country": schema.StringAttribute{
						Required:      true,
						Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
			"description": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A descriptive text providing additional context about the tax location. This can include information about the venue, types of events held, services available, or any relevant details for better identification (for example, \"A spacious auditorium suitable for large concerts and events.\").",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
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
			"type": schema.StringAttribute{
				Required:      true,
				Description:   "The type of tax location to be defined. Currently the only option is `performance`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("performance")},
			},
		},
	}
}

func (r *TaxLocationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TaxLocationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTaxLocationCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building TaxLocation create params", err.Error())
		return
	}

	obj, err := r.client.V1TaxLocations.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating TaxLocation", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TaxLocations.B, r.client.V1TaxLocations.Key, stripe.FormatURLPath("/v1/tax/locations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TaxLocation create raw response", err.Error())
		return
	}

	if err := flattenTaxLocation(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxLocation create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TaxLocationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState TaxLocationResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TaxLocationResourceModel
	state = priorState

	obj, err := r.client.V1TaxLocations.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading TaxLocation", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TaxLocations.B, r.client.V1TaxLocations.Key, stripe.FormatURLPath("/v1/tax/locations/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TaxLocation raw response", err.Error())
		return
	}

	if err := flattenTaxLocation(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxLocation read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *TaxLocationResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "TaxLocation does not support updates")
}

func (r *TaxLocationResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *TaxLocationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandTaxLocationCreate(plan TaxLocationResourceModel) (*stripe.TaxLocationCreateParams, error) {
	params := &stripe.TaxLocationCreateParams{}

	if !plan.Address.IsNull() && !plan.Address.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Address", plan.Address) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "address", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}

	return params, nil
}

func flattenTaxLocation(obj *stripe.TaxLocation, state *TaxLocationResourceModel) error {
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
