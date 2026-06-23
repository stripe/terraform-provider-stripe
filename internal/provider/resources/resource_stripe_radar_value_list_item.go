//
// File generated from our OpenAPI spec
//

package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &RadarValueListItemResource{}

var _ resource.ResourceWithConfigure = &RadarValueListItemResource{}

var _ resource.ResourceWithImportState = &RadarValueListItemResource{}

func NewRadarValueListItemResource() resource.Resource {
	return &RadarValueListItemResource{}
}

type RadarValueListItemResource struct {
	client *stripe.Client
}

type RadarValueListItemResourceModel struct {
	Object    types.String `tfsdk:"object"`
	Created   types.Int64  `tfsdk:"created"`
	CreatedBy types.String `tfsdk:"created_by"`
	ID        types.String `tfsdk:"id"`
	Livemode  types.Bool   `tfsdk:"livemode"`
	Value     types.String `tfsdk:"value"`
	ValueList types.String `tfsdk:"value_list"`
}

func (r *RadarValueListItemResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RadarValueListItemResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radar_value_list_item"
}

func (r *RadarValueListItemResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Value list items allow you to add specific values to a given Radar value list, which can then be used in rules.\n\nRelated guide: [Managing list items](https://docs.stripe.com/radar/lists#managing-list-items)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("radar.value_list_item")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"created_by": schema.StringAttribute{
				Computed:      true,
				Description:   "The name or email address of the user who added this item to the value list.",
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
			"value": schema.StringAttribute{
				Required:      true,
				Description:   "The value of the item.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"value_list": schema.StringAttribute{
				Required:      true,
				Description:   "The identifier of the value list this item belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *RadarValueListItemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RadarValueListItemResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandRadarValueListItemCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building RadarValueListItem create params", err.Error())
		return
	}

	obj, err := r.client.V1RadarValueListItems.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating RadarValueListItem", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1RadarValueListItems.B, r.client.V1RadarValueListItems.Key, stripe.FormatURLPath("/v1/radar/value_list_items/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating RadarValueListItem create raw response", err.Error())
		return
	}

	if err := flattenRadarValueListItem(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening RadarValueListItem create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RadarValueListItemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState RadarValueListItemResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state RadarValueListItemResourceModel
	state = priorState

	obj, err := r.client.V1RadarValueListItems.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading RadarValueListItem", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1RadarValueListItems.B, r.client.V1RadarValueListItems.Key, stripe.FormatURLPath("/v1/radar/value_list_items/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating RadarValueListItem raw response", err.Error())
		return
	}

	if err := flattenRadarValueListItem(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening RadarValueListItem read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *RadarValueListItemResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "RadarValueListItem does not support updates")
}

func (r *RadarValueListItemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RadarValueListItemResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1RadarValueListItems.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting RadarValueListItem", err.Error())
		return
	}
}

func (r *RadarValueListItemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandRadarValueListItemCreate(plan RadarValueListItemResourceModel) (*stripe.RadarValueListItemCreateParams, error) {
	params := &stripe.RadarValueListItemCreateParams{}

	if !plan.Value.IsNull() && !plan.Value.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Value", "Value", plan.Value.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "value", params)
		}
	}
	if !plan.ValueList.IsNull() && !plan.ValueList.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ValueList", "ValueList", plan.ValueList.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "value_list", params)
		}
	}

	return params, nil
}

func flattenRadarValueListItem(obj *stripe.RadarValueListItem, state *RadarValueListItemResourceModel) error {
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
		if rawValueCreatedBy, rawOk := plainValueAtPath(raw, "created_by"); rawOk {
			if valueCreatedBy, err := flattenPlainValue(rawValueCreatedBy, types.StringType, "created_by", "raw response"); err != nil {
				return err
			} else {
				if typedCreatedBy, ok := valueCreatedBy.(types.String); ok {
					state.CreatedBy = typedCreatedBy
				}
			}
		} else if !hasRaw {
			if responseValueCreatedBy, ok := plainFromResponseField(obj, "CreatedBy"); ok {
				if valueCreatedBy, err := flattenPlainValue(responseValueCreatedBy, types.StringType, "created_by", "response struct"); err != nil {
					return err
				} else {
					if typedCreatedBy, ok := valueCreatedBy.(types.String); ok {
						state.CreatedBy = typedCreatedBy
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
		if rawValueValue, rawOk := plainValueAtPath(raw, "value"); rawOk {
			if valueValue, err := flattenPlainValue(rawValueValue, types.StringType, "value", "raw response"); err != nil {
				return err
			} else {
				if typedValue, ok := valueValue.(types.String); ok {
					state.Value = typedValue
				}
			}
		} else if !hasRaw {
			if responseValueValue, ok := plainFromResponseField(obj, "Value"); ok {
				if valueValue, err := flattenPlainValue(responseValueValue, types.StringType, "value", "response struct"); err != nil {
					return err
				} else {
					if typedValue, ok := valueValue.(types.String); ok {
						state.Value = typedValue
					}
				}
			}
		}
	}
	{
		if rawValueValueList, rawOk := plainValueAtPath(raw, "value_list"); rawOk {
			if valueValueList, err := flattenPlainValue(rawValueValueList, types.StringType, "value_list", "raw response"); err != nil {
				return err
			} else {
				if typedValueList, ok := valueValueList.(types.String); ok {
					state.ValueList = typedValueList
				}
			}
		} else if !hasRaw {
			if responseValueValueList, ok := plainFromResponseField(obj, "ValueList"); ok {
				if valueValueList, err := flattenPlainValue(responseValueValueList, types.StringType, "value_list", "response struct"); err != nil {
					return err
				} else {
					if typedValueList, ok := valueValueList.(types.String); ok {
						state.ValueList = typedValueList
					}
				}
			}
		}
	}
	return nil
}
