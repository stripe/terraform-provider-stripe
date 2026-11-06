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

var _ resource.Resource = &FileLinkResource{}

var _ resource.ResourceWithConfigure = &FileLinkResource{}

var _ resource.ResourceWithImportState = &FileLinkResource{}

func NewFileLinkResource() resource.Resource {
	return &FileLinkResource{}
}

type FileLinkResource struct {
	client *stripe.Client
}

type FileLinkResourceModel struct {
	Object    types.String `tfsdk:"object"`
	Created   types.Int64  `tfsdk:"created"`
	Expired   types.Bool   `tfsdk:"expired"`
	ExpiresAt types.Int64  `tfsdk:"expires_at"`
	File      types.String `tfsdk:"file"`
	ID        types.String `tfsdk:"id"`
	Livemode  types.Bool   `tfsdk:"livemode"`
	Metadata  types.Map    `tfsdk:"metadata"`
	URL       types.String `tfsdk:"url"`
}

func (r *FileLinkResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *FileLinkResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file_link"
}

func (r *FileLinkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "To share the contents of a `File` object with non-Stripe users, you can\ncreate a `FileLink`. `FileLink`s contain a URL that you can use to\nretrieve the contents of the file without authentication.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("file_link")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"expired": schema.BoolAttribute{
				Computed:      true,
				Description:   "Returns if the link is already expired.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"expires_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Time that the link expires.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"file": schema.StringAttribute{
				Required:      true,
				Description:   "The file object this link points to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
			"url": schema.StringAttribute{
				Computed:      true,
				Description:   "The publicly accessible URL to download the file.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *FileLinkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FileLinkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandFileLinkCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building FileLink create params", err.Error())
		return
	}

	obj, err := r.client.V1FileLinks.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating FileLink", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1FileLinks.B, r.client.V1FileLinks.Key, stripe.FormatURLPath("/v1/file_links/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating FileLink create raw response", err.Error())
		return
	}

	if err := flattenFileLink(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening FileLink create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *FileLinkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState FileLinkResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state FileLinkResourceModel
	state = priorState

	obj, err := r.client.V1FileLinks.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading FileLink", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1FileLinks.B, r.client.V1FileLinks.Key, stripe.FormatURLPath("/v1/file_links/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating FileLink raw response", err.Error())
		return
	}

	if err := flattenFileLink(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening FileLink read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *FileLinkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FileLinkResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state FileLinkResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandFileLinkUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building FileLink update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building FileLink update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1FileLinks.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating FileLink", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1FileLinks.B, r.client.V1FileLinks.Key, stripe.FormatURLPath("/v1/file_links/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating FileLink update raw response", err.Error())
		return
	}

	if err := flattenFileLink(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening FileLink update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *FileLinkResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *FileLinkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandFileLinkCreate(plan FileLinkResourceModel) (*stripe.FileLinkCreateParams, error) {
	params := &stripe.FileLinkCreateParams{}

	if !plan.ExpiresAt.IsNull() && !plan.ExpiresAt.IsUnknown() {
		params.ExpiresAt = stripe.Int64(plan.ExpiresAt.ValueInt64())
	}
	if !plan.File.IsNull() && !plan.File.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FileID", "File", plan.File.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "file", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}

	return params, nil
}

func expandFileLinkUpdate(plan FileLinkResourceModel, state FileLinkResourceModel) (*stripe.FileLinkUpdateParams, error) {
	params := &stripe.FileLinkUpdateParams{}

	if !plan.ExpiresAt.Equal(state.ExpiresAt) && !plan.ExpiresAt.IsNull() && !plan.ExpiresAt.IsUnknown() {
		params.ExpiresAt = stripe.Int64(plan.ExpiresAt.ValueInt64())
	}
	if !plan.Metadata.Equal(state.Metadata) && !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			if !plan.Metadata.Equal(state.Metadata) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "metadata", params)
			}
		}
	}

	return params, nil
}

func flattenFileLink(obj *stripe.FileLink, state *FileLinkResourceModel) error {
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
		if rawValueExpired, rawOk := plainValueAtPath(raw, "expired"); rawOk {
			if valueExpired, err := flattenPlainValue(rawValueExpired, types.BoolType, "expired", "raw response"); err != nil {
				return err
			} else {
				if typedExpired, ok := valueExpired.(types.Bool); ok {
					state.Expired = typedExpired
				}
			}
		} else if !hasRaw {
			if responseValueExpired, ok := plainFromResponseField(obj, "Expired"); ok {
				if valueExpired, err := flattenPlainValue(responseValueExpired, types.BoolType, "expired", "response struct"); err != nil {
					return err
				} else {
					if typedExpired, ok := valueExpired.(types.Bool); ok {
						state.Expired = typedExpired
					}
				}
			}
		}
	}
	{
		if rawValueExpiresAt, rawOk := plainValueAtPath(raw, "expires_at"); rawOk {
			if valueExpiresAt, err := flattenPlainValue(rawValueExpiresAt, types.Int64Type, "expires_at", "raw response"); err != nil {
				return err
			} else {
				if typedExpiresAt, ok := valueExpiresAt.(types.Int64); ok {
					state.ExpiresAt = typedExpiresAt
				}
			}
		} else if !hasRaw {
			if responseValueExpiresAt, ok := plainFromResponseField(obj, "ExpiresAt"); ok {
				if valueExpiresAt, err := flattenPlainValue(responseValueExpiresAt, types.Int64Type, "expires_at", "response struct"); err != nil {
					return err
				} else {
					if typedExpiresAt, ok := valueExpiresAt.(types.Int64); ok {
						state.ExpiresAt = typedExpiresAt
					}
				}
			}
		}
	}
	{
		if state.File.IsNull() || state.File.IsUnknown() {
			if rawValueFile, rawOk := plainValueAtPath(raw, "file"); rawOk {
				if typedFile, ok := plainToStringIDValue(rawValueFile); ok {
					state.File = typedFile
				}
			} else if !hasRaw {
				if responseValueFile, ok := plainFromResponseField(obj, "File"); ok {
					if typedFile, ok := plainToStringIDValue(responseValueFile); ok {
						state.File = typedFile
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
		if rawValueURL, rawOk := plainValueAtPath(raw, "url"); rawOk {
			if valueURL, err := flattenPlainValue(rawValueURL, types.StringType, "url", "raw response"); err != nil {
				return err
			} else {
				if typedURL, ok := valueURL.(types.String); ok {
					state.URL = typedURL
				}
			}
		} else if !hasRaw {
			if responseValueURL, ok := plainFromResponseField(obj, "URL"); ok {
				if valueURL, err := flattenPlainValue(responseValueURL, types.StringType, "url", "response struct"); err != nil {
					return err
				} else {
					if typedURL, ok := valueURL.(types.String); ok {
						state.URL = typedURL
					}
				}
			}
		}
	}
	return nil
}
