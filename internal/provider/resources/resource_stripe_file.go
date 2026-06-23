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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &FileResource{}

var _ resource.ResourceWithConfigure = &FileResource{}

var _ resource.ResourceWithImportState = &FileResource{}

func NewFileResource() resource.Resource {
	return &FileResource{}
}

type FileResource struct {
	client *stripe.Client
}

type FileResourceModel struct {
	Object       types.String `tfsdk:"object"`
	Created      types.Int64  `tfsdk:"created"`
	ExpiresAt    types.Int64  `tfsdk:"expires_at"`
	Filename     types.String `tfsdk:"filename"`
	ID           types.String `tfsdk:"id"`
	Purpose      types.String `tfsdk:"purpose"`
	Size         types.Int64  `tfsdk:"size"`
	Title        types.String `tfsdk:"title"`
	Type         types.String `tfsdk:"type"`
	URL          types.String `tfsdk:"url"`
	File         types.String `tfsdk:"file"`
	FileLinkData types.Object `tfsdk:"file_link_data"`
}

func (r *FileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *FileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file"
}

func (r *FileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "This object represents files hosted on Stripe's servers. You can upload\nfiles with the [create file](https://api.stripe.com#create_file) request\n(for example, when uploading dispute evidence). Stripe also\ncreates files independently (for example, the results of a [Sigma scheduled\nquery](#scheduled_queries)).\n\nRelated guide: [File upload guide](https://docs.stripe.com/file-upload)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("file")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"expires_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "The file expires and isn't available at this time in epoch seconds.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"filename": schema.StringAttribute{
				Computed:      true,
				Description:   "The suitable name for saving the file to a filesystem.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"purpose": schema.StringAttribute{
				Required:      true,
				Description:   "The [purpose](https://docs.stripe.com/file-upload#uploading-a-file) of the uploaded file.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("account_requirement", "additional_verification", "business_icon", "business_logo", "customer_signature", "dispute_evidence", "document_provider_identity_document", "finance_report_run", "financial_account_statement", "identity_document", "identity_document_downloadable", "issuing_regulatory_reporting", "pci_document", "platform_terms_of_service", "selfie", "sigma_scheduled_query", "tax_document_user_upload", "terminal_android_apk", "terminal_reader_splashscreen", "terminal_wifi_certificate", "terminal_wifi_private_key")},
			},
			"size": schema.Int64Attribute{
				Computed:      true,
				Description:   "The size of the file object in bytes.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"title": schema.StringAttribute{
				Computed:      true,
				Description:   "A suitable title for the document.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"type": schema.StringAttribute{
				Computed:      true,
				Description:   "The returned file type (for example, `csv`, `pdf`, `jpg`, or `png`).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"url": schema.StringAttribute{
				Computed:      true,
				Description:   "Use your live secret API key to download the file from this URL.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"file": schema.StringAttribute{
				Required:      true,
				Description:   "A file to upload. Make sure that the specifications follow RFC 2388, which defines file transfers for the `multipart/form-data` protocol.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"file_link_data": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "Optional parameters that automatically create a [file link](https://api.stripe.com#file_links) for the newly created file.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				WriteOnly:     true,
				Attributes: map[string]schema.Attribute{
					"create": schema.BoolAttribute{
						Required:      true,
						Description:   "Set this to `true` to create a file link for the newly created file. Creating a link is only possible when the file's `purpose` is one of the following: `business_icon`, `business_logo`, `customer_signature`, `dispute_evidence`, `issuing_regulatory_reporting`, `pci_document`, `tax_document_user_upload`, `terminal_android_apk`, or `terminal_reader_splashscreen`.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
					"expires_at": schema.Int64Attribute{
						Optional:      true,
						Description:   "The link isn't available after this future timestamp.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						WriteOnly:     true,
					},
					"metadata": schema.MapAttribute{
						Optional:      true,
						Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
						PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
						WriteOnly:     true,
						ElementType:   types.StringType,
					},
				},
			},
		},
	}
}

func (r *FileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FileResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config FileResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"File"}, []string{"FileLinkData"}, []string{"FileLinkData", "create"}, []string{"FileLinkData", "expires_at"}, []string{"FileLinkData", "metadata"}})

	params, err := expandFileCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building File create params", err.Error())
		return
	}

	obj, err := r.client.V1Files.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating File", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Files.B, r.client.V1Files.Key, stripe.FormatURLPath("/v1/files/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating File create raw response", err.Error())
		return
	}

	if err := flattenFile(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening File create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"File"}, []string{"FileLinkData"}, []string{"FileLinkData", "create"}, []string{"FileLinkData", "expires_at"}, []string{"FileLinkData", "metadata"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *FileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState FileResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state FileResourceModel
	state = priorState

	obj, err := r.client.V1Files.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading File", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Files.B, r.client.V1Files.Key, stripe.FormatURLPath("/v1/files/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating File raw response", err.Error())
		return
	}

	if err := flattenFile(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening File read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"File"}, []string{"FileLinkData"}, []string{"FileLinkData", "create"}, []string{"FileLinkData", "expires_at"}, []string{"FileLinkData", "metadata"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *FileResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "File does not support updates")
}

func (r *FileResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *FileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandFileCreate(plan FileResourceModel) (*stripe.FileCreateParams, error) {
	params := &stripe.FileCreateParams{}

	if !plan.Purpose.IsNull() && !plan.Purpose.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Purpose", "Purpose", plan.Purpose.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "purpose", params)
		}
	}
	if !plan.File.IsNull() && !plan.File.IsUnknown() {
		if err := assignFilePathToNamedField(params, "File", plan.File.ValueString()); err != nil {
			return nil, fmt.Errorf("failed to assign attribute %q on %T: %w", "file", params, err)
		}
	}
	if !plan.FileLinkData.IsNull() && !plan.FileLinkData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FileLinkData", plan.FileLinkData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "file_link_data", params)
		}
	}

	return params, nil
}

func flattenFile(obj *stripe.File, state *FileResourceModel) error {
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
		if rawValueFilename, rawOk := plainValueAtPath(raw, "filename"); rawOk {
			if valueFilename, err := flattenPlainValue(rawValueFilename, types.StringType, "filename", "raw response"); err != nil {
				return err
			} else {
				if typedFilename, ok := valueFilename.(types.String); ok {
					state.Filename = typedFilename
				}
			}
		} else if !hasRaw {
			if responseValueFilename, ok := plainFromResponseField(obj, "Filename"); ok {
				if valueFilename, err := flattenPlainValue(responseValueFilename, types.StringType, "filename", "response struct"); err != nil {
					return err
				} else {
					if typedFilename, ok := valueFilename.(types.String); ok {
						state.Filename = typedFilename
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
		if rawValuePurpose, rawOk := plainValueAtPath(raw, "purpose"); rawOk {
			if valuePurpose, err := flattenPlainValue(rawValuePurpose, types.StringType, "purpose", "raw response"); err != nil {
				return err
			} else {
				if typedPurpose, ok := valuePurpose.(types.String); ok {
					state.Purpose = typedPurpose
				}
			}
		} else if !hasRaw {
			if responseValuePurpose, ok := plainFromResponseField(obj, "Purpose"); ok {
				if valuePurpose, err := flattenPlainValue(responseValuePurpose, types.StringType, "purpose", "response struct"); err != nil {
					return err
				} else {
					if typedPurpose, ok := valuePurpose.(types.String); ok {
						state.Purpose = typedPurpose
					}
				}
			}
		}
	}
	{
		if rawValueSize, rawOk := plainValueAtPath(raw, "size"); rawOk {
			if valueSize, err := flattenPlainValue(rawValueSize, types.Int64Type, "size", "raw response"); err != nil {
				return err
			} else {
				if typedSize, ok := valueSize.(types.Int64); ok {
					state.Size = typedSize
				}
			}
		} else if !hasRaw {
			if responseValueSize, ok := plainFromResponseField(obj, "Size"); ok {
				if valueSize, err := flattenPlainValue(responseValueSize, types.Int64Type, "size", "response struct"); err != nil {
					return err
				} else {
					if typedSize, ok := valueSize.(types.Int64); ok {
						state.Size = typedSize
					}
				}
			}
		}
	}
	{
		if rawValueTitle, rawOk := plainValueAtPath(raw, "title"); rawOk {
			if valueTitle, err := flattenPlainValue(rawValueTitle, types.StringType, "title", "raw response"); err != nil {
				return err
			} else {
				if typedTitle, ok := valueTitle.(types.String); ok {
					state.Title = typedTitle
				}
			}
		} else if !hasRaw {
			if responseValueTitle, ok := plainFromResponseField(obj, "Title"); ok {
				if valueTitle, err := flattenPlainValue(responseValueTitle, types.StringType, "title", "response struct"); err != nil {
					return err
				} else {
					if typedTitle, ok := valueTitle.(types.String); ok {
						state.Title = typedTitle
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
