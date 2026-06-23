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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &RadarValueListResource{}

var _ resource.ResourceWithConfigure = &RadarValueListResource{}

var _ resource.ResourceWithImportState = &RadarValueListResource{}

func NewRadarValueListResource() resource.Resource {
	return &RadarValueListResource{}
}

type RadarValueListResource struct {
	client *stripe.Client
}

type RadarValueListResourceModel struct {
	Object    types.String `tfsdk:"object"`
	Alias     types.String `tfsdk:"alias"`
	Created   types.Int64  `tfsdk:"created"`
	CreatedBy types.String `tfsdk:"created_by"`
	ID        types.String `tfsdk:"id"`
	ItemType  types.String `tfsdk:"item_type"`
	Livemode  types.Bool   `tfsdk:"livemode"`
	Metadata  types.Map    `tfsdk:"metadata"`
	Name      types.String `tfsdk:"name"`
}

func (r *RadarValueListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RadarValueListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radar_value_list"
}

func (r *RadarValueListResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Value lists allow you to group values together which can then be referenced in rules.\n\nRelated guide: [Default Stripe lists](https://docs.stripe.com/radar/lists#managing-list-items)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("radar.value_list")},
			},
			"alias": schema.StringAttribute{
				Required:    true,
				Description: "The name of the value list for use in rules.",
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"created_by": schema.StringAttribute{
				Computed:      true,
				Description:   "The name or email address of the user who created this value list.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"item_type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The type of items in the value list. One of `card_fingerprint`, `card_bin`, `crypto_fingerprint`, `email`, `ip_address`, `country`, `string`, `case_sensitive_string`, `customer_id`, `account`, `sepa_debit_fingerprint`, or `us_bank_account_fingerprint`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("account", "card_bin", "card_fingerprint", "case_sensitive_string", "country", "crypto_fingerprint", "customer_id", "email", "ip_address", "sepa_debit_fingerprint", "string", "us_bank_account_fingerprint")},
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
				Required:    true,
				Description: "The name of the value list.",
			},
		},
	}
}

func (r *RadarValueListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RadarValueListResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandRadarValueListCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building RadarValueList create params", err.Error())
		return
	}

	obj, err := r.client.V1RadarValueLists.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating RadarValueList", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1RadarValueLists.B, r.client.V1RadarValueLists.Key, stripe.FormatURLPath("/v1/radar/value_lists/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating RadarValueList create raw response", err.Error())
		return
	}

	if err := flattenRadarValueList(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening RadarValueList create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RadarValueListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState RadarValueListResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state RadarValueListResourceModel
	state = priorState

	obj, err := r.client.V1RadarValueLists.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading RadarValueList", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1RadarValueLists.B, r.client.V1RadarValueLists.Key, stripe.FormatURLPath("/v1/radar/value_lists/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating RadarValueList raw response", err.Error())
		return
	}

	if err := flattenRadarValueList(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening RadarValueList read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *RadarValueListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RadarValueListResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state RadarValueListResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandRadarValueListUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building RadarValueList update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building RadarValueList update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1RadarValueLists.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating RadarValueList", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1RadarValueLists.B, r.client.V1RadarValueLists.Key, stripe.FormatURLPath("/v1/radar/value_lists/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating RadarValueList update raw response", err.Error())
		return
	}

	if err := flattenRadarValueList(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening RadarValueList update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *RadarValueListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RadarValueListResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1RadarValueLists.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting RadarValueList", err.Error())
		return
	}
}

func (r *RadarValueListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandRadarValueListCreate(plan RadarValueListResourceModel) (*stripe.RadarValueListCreateParams, error) {
	params := &stripe.RadarValueListCreateParams{}

	if !plan.Alias.IsNull() && !plan.Alias.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Alias", "Alias", plan.Alias.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "alias", params)
		}
	}
	if !plan.ItemType.IsNull() && !plan.ItemType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ItemType", "ItemType", plan.ItemType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "item_type", params)
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

	return params, nil
}

func expandRadarValueListUpdate(plan RadarValueListResourceModel, state RadarValueListResourceModel) (*stripe.RadarValueListUpdateParams, error) {
	params := &stripe.RadarValueListUpdateParams{}

	if !plan.Alias.Equal(state.Alias) && !plan.Alias.IsNull() && !plan.Alias.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Alias", "Alias", plan.Alias.ValueString()) {
			if !plan.Alias.Equal(state.Alias) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "alias", params)
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

	return params, nil
}

func flattenRadarValueList(obj *stripe.RadarValueList, state *RadarValueListResourceModel) error {
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
		if rawValueAlias, rawOk := plainValueAtPath(raw, "alias"); rawOk {
			if valueAlias, err := flattenPlainValue(rawValueAlias, types.StringType, "alias", "raw response"); err != nil {
				return err
			} else {
				if typedAlias, ok := valueAlias.(types.String); ok {
					state.Alias = typedAlias
				}
			}
		} else if !hasRaw {
			if responseValueAlias, ok := plainFromResponseField(obj, "Alias"); ok {
				if valueAlias, err := flattenPlainValue(responseValueAlias, types.StringType, "alias", "response struct"); err != nil {
					return err
				} else {
					if typedAlias, ok := valueAlias.(types.String); ok {
						state.Alias = typedAlias
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
		if rawValueItemType, rawOk := plainValueAtPath(raw, "item_type"); rawOk {
			if valueItemType, err := flattenPlainValue(rawValueItemType, types.StringType, "item_type", "raw response"); err != nil {
				return err
			} else {
				if typedItemType, ok := valueItemType.(types.String); ok {
					state.ItemType = typedItemType
				}
			}
		} else if !hasRaw {
			if responseValueItemType, ok := plainFromResponseField(obj, "ItemType"); ok {
				if valueItemType, err := flattenPlainValue(responseValueItemType, types.StringType, "item_type", "response struct"); err != nil {
					return err
				} else {
					if typedItemType, ok := valueItemType.(types.String); ok {
						state.ItemType = typedItemType
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
	return nil
}
