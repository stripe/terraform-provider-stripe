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

var _ resource.Resource = &ApplePayDomainResource{}

var _ resource.ResourceWithConfigure = &ApplePayDomainResource{}

var _ resource.ResourceWithImportState = &ApplePayDomainResource{}

func NewApplePayDomainResource() resource.Resource {
	return &ApplePayDomainResource{}
}

type ApplePayDomainResource struct {
	client *stripe.Client
}

type ApplePayDomainResourceModel struct {
	Object     types.String `tfsdk:"object"`
	Created    types.Int64  `tfsdk:"created"`
	DomainName types.String `tfsdk:"domain_name"`
	ID         types.String `tfsdk:"id"`
	Livemode   types.Bool   `tfsdk:"livemode"`
}

func (r *ApplePayDomainResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ApplePayDomainResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apple_pay_domain"
}

func (r *ApplePayDomainResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("apple_pay_domain")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"domain_name": schema.StringAttribute{
				Required: true,

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
		},
	}
}

func (r *ApplePayDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ApplePayDomainResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandApplePayDomainCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building ApplePayDomain create params", err.Error())
		return
	}

	obj, err := r.client.V1ApplePayDomains.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating ApplePayDomain", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ApplePayDomains.B, r.client.V1ApplePayDomains.Key, stripe.FormatURLPath("/v1/apple_pay/domains/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ApplePayDomain create raw response", err.Error())
		return
	}

	if err := flattenApplePayDomain(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening ApplePayDomain create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ApplePayDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState ApplePayDomainResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ApplePayDomainResourceModel
	state = priorState

	obj, err := r.client.V1ApplePayDomains.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading ApplePayDomain", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ApplePayDomains.B, r.client.V1ApplePayDomains.Key, stripe.FormatURLPath("/v1/apple_pay/domains/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ApplePayDomain raw response", err.Error())
		return
	}

	if err := flattenApplePayDomain(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening ApplePayDomain read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *ApplePayDomainResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "ApplePayDomain does not support updates")
}

func (r *ApplePayDomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ApplePayDomainResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1ApplePayDomains.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting ApplePayDomain", err.Error())
		return
	}
}

func (r *ApplePayDomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandApplePayDomainCreate(plan ApplePayDomainResourceModel) (*stripe.ApplePayDomainCreateParams, error) {
	params := &stripe.ApplePayDomainCreateParams{}

	if !plan.DomainName.IsNull() && !plan.DomainName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DomainName", "DomainName", plan.DomainName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "domain_name", params)
		}
	}

	return params, nil
}

func flattenApplePayDomain(obj *stripe.ApplePayDomain, state *ApplePayDomainResourceModel) error {
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
		if rawValueDomainName, rawOk := plainValueAtPath(raw, "domain_name"); rawOk {
			if valueDomainName, err := flattenPlainValue(rawValueDomainName, types.StringType, "domain_name", "raw response"); err != nil {
				return err
			} else {
				if typedDomainName, ok := valueDomainName.(types.String); ok {
					state.DomainName = typedDomainName
				}
			}
		} else if !hasRaw {
			if responseValueDomainName, ok := plainFromResponseField(obj, "DomainName"); ok {
				if valueDomainName, err := flattenPlainValue(responseValueDomainName, types.StringType, "domain_name", "response struct"); err != nil {
					return err
				} else {
					if typedDomainName, ok := valueDomainName.(types.String); ok {
						state.DomainName = typedDomainName
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
	return nil
}
