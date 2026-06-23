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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &ProductFeatureResource{}

var _ resource.ResourceWithConfigure = &ProductFeatureResource{}

var _ resource.ResourceWithImportState = &ProductFeatureResource{}

func NewProductFeatureResource() resource.Resource {
	return &ProductFeatureResource{}
}

type ProductFeatureResource struct {
	client *stripe.Client
}

type ProductFeatureResourceModel struct {
	Object             types.String `tfsdk:"object"`
	EntitlementFeature types.String `tfsdk:"entitlement_feature"`
	ID                 types.String `tfsdk:"id"`
	Livemode           types.Bool   `tfsdk:"livemode"`
	Product            types.String `tfsdk:"product"`
}

func (r *ProductFeatureResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ProductFeatureResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_product_feature"
}

func (r *ProductFeatureResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A product_feature represents an attachment between a feature and a product.\nWhen a product is purchased that has a feature attached, Stripe will create an entitlement to the feature for the purchasing customer.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("product_feature")},
			},
			"entitlement_feature": schema.StringAttribute{
				Required:      true,
				Description:   "A feature represents a monetizable ability or functionality in your system.\nFeatures can be assigned to products, and when those products are purchased, Stripe will create an entitlement to the feature for the purchasing customer.",
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
			"product": schema.StringAttribute{
				Required: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *ProductFeatureResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProductFeatureResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandProductFeatureCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building ProductFeature create params", err.Error())
		return
	}

	if !plan.Product.IsNull() && !plan.Product.IsUnknown() {
		if !assignStringToNamedField(params, "Product", "ID", plan.Product.ValueString()) {
			resp.Diagnostics.AddError("Error building ProductFeature create path params", fmt.Sprintf("Failed to assign path parameter %q on %T", "product", params))
			return
		}
	}
	obj, err := r.client.V1ProductFeatures.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating ProductFeature", err.Error())
		return
	}

	rawReadParams := &stripe.ProductFeatureRetrieveParams{}
	if !plan.Product.IsNull() && !plan.Product.IsUnknown() {
		if !assignStringToNamedField(rawReadParams, "Product", "ID", plan.Product.ValueString()) {
			resp.Diagnostics.AddError("Error building ProductFeature read params for raw hydration after create", fmt.Sprintf("Failed to assign path parameter %q on %T", "product", rawReadParams))
			return
		}
	}

	if err := ensureRawResponse(obj, r.client.V1ProductFeatures.B, r.client.V1ProductFeatures.Key, stripe.FormatURLPath("/v1/products/%s/features/%s", plan.Product.ValueString(), obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating ProductFeature create raw response", err.Error())
		return
	}

	if err := flattenProductFeature(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening ProductFeature create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ProductFeatureResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState ProductFeatureResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ProductFeatureResourceModel
	state = priorState

	params := &stripe.ProductFeatureRetrieveParams{}
	if !state.Product.IsNull() && !state.Product.IsUnknown() {
		if !assignStringToNamedField(params, "Product", "ID", state.Product.ValueString()) {
			resp.Diagnostics.AddError("Error building ProductFeature read params", fmt.Sprintf("Failed to assign path parameter %q on %T", "product", params))
			return
		}
	}

	obj, err := r.client.V1ProductFeatures.Retrieve(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error reading ProductFeature", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ProductFeatures.B, r.client.V1ProductFeatures.Key, stripe.FormatURLPath("/v1/products/%s/features/%s", state.Product.ValueString(), state.ID.ValueString()), params); err != nil {
		resp.Diagnostics.AddError("Error hydrating ProductFeature raw response", err.Error())
		return
	}

	if err := flattenProductFeature(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening ProductFeature read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *ProductFeatureResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "ProductFeature does not support updates")
}

func (r *ProductFeatureResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ProductFeatureResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := &stripe.ProductFeatureDeleteParams{}
	if !state.Product.IsNull() && !state.Product.IsUnknown() {
		if !assignStringToNamedField(params, "Product", "ID", state.Product.ValueString()) {
			resp.Diagnostics.AddError("Error building ProductFeature delete params", fmt.Sprintf("Failed to assign path parameter %q on %T", "product", params))
			return
		}
	}

	_, err := r.client.V1ProductFeatures.Delete(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting ProductFeature", err.Error())
		return
	}
}

func (r *ProductFeatureResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Unexpected import identifier", fmt.Sprintf("Expected import identifier in the form \"product/id\", got %q", req.ID))
		return
	}

	diags := resp.State.SetAttribute(ctx, path.Root("product"), types.StringValue(parts[0]))
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(parts[1]))
	resp.Diagnostics.Append(diags...)
}

func expandProductFeatureCreate(plan ProductFeatureResourceModel) (*stripe.ProductFeatureCreateParams, error) {
	params := &stripe.ProductFeatureCreateParams{}

	if !plan.EntitlementFeature.IsNull() && !plan.EntitlementFeature.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "EntitlementFeatureID", "EntitlementFeature", plan.EntitlementFeature.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "entitlement_feature", params)
		}
	}

	return params, nil
}

func flattenProductFeature(obj *stripe.ProductFeature, state *ProductFeatureResourceModel) error {
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
		if state.EntitlementFeature.IsNull() || state.EntitlementFeature.IsUnknown() {
			if rawValueEntitlementFeature, rawOk := plainValueAtPath(raw, "entitlement_feature"); rawOk {
				if typedEntitlementFeature, ok := plainToStringIDValue(rawValueEntitlementFeature); ok {
					state.EntitlementFeature = typedEntitlementFeature
				}
			} else if !hasRaw {
				if responseValueEntitlementFeature, ok := plainFromResponseField(obj, "EntitlementFeature"); ok {
					if typedEntitlementFeature, ok := plainToStringIDValue(responseValueEntitlementFeature); ok {
						state.EntitlementFeature = typedEntitlementFeature
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
