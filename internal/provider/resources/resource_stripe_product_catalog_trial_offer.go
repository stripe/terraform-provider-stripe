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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"reflect"

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &ProductCatalogTrialOfferResource{}

var _ resource.ResourceWithConfigure = &ProductCatalogTrialOfferResource{}

var _ resource.ResourceWithImportState = &ProductCatalogTrialOfferResource{}

func NewProductCatalogTrialOfferResource() resource.Resource {
	return &ProductCatalogTrialOfferResource{}
}

type ProductCatalogTrialOfferResource struct {
	client *stripe.Client
}

type ProductCatalogTrialOfferResourceModel struct {
	Object      types.String `tfsdk:"object"`
	Active      types.Bool   `tfsdk:"active"`
	Duration    types.Object `tfsdk:"duration"`
	EndBehavior types.Object `tfsdk:"end_behavior"`
	ID          types.String `tfsdk:"id"`
	Livemode    types.Bool   `tfsdk:"livemode"`
	Nickname    types.String `tfsdk:"nickname"`
	Price       types.String `tfsdk:"price"`
}

func (r *ProductCatalogTrialOfferResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ProductCatalogTrialOfferResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_product_catalog_trial_offer"
}

func (r *ProductCatalogTrialOfferResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Trial offers let you define free or paid introductory pricing for a subscription item.\nA TrialOffer specifies the price to charge during the trial, how many billing intervals\nthe trial lasts, and what price the subscription item transitions to when the trial ends.\nYou attach a TrialOffer to a subscription item\nusing `items[current_trial][trial_offer]` when creating or updating a subscription.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("product_catalog.trial_offer")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the trial offer is active. Set to false to archive the trial offer.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"duration": schema.SingleNestedAttribute{
				Required: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"relative": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"iterations": schema.Int64Attribute{
								Required:      true,
								Description:   "The number of iterations of the price's interval for this trial offer.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
						},
					},
					"type": schema.StringAttribute{
						Required:      true,
						Description:   "The type of trial offer duration.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("relative")},
					},
				},
			},
			"end_behavior": schema.SingleNestedAttribute{
				Required: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"transition": schema.SingleNestedAttribute{
						Required: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"price": schema.StringAttribute{
								Required:      true,
								Description:   "The new price to use at the end of the trial offer period.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
						},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "The type of behavior when the trial offer ends.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("transition")},
					},
				},
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
			"nickname": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A brief description of the trial offer, hidden from customers.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"price": schema.StringAttribute{
				Required:      true,
				Description:   "The price during the trial offer.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *ProductCatalogTrialOfferResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProductCatalogTrialOfferResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandProductCatalogTrialOfferCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building ProductCatalogTrialOffer create params", err.Error())
		return
	}

	obj, err := r.client.V1ProductCatalogTrialOffers.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating ProductCatalogTrialOffer", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ProductCatalogTrialOffers.B, r.client.V1ProductCatalogTrialOffers.Key, stripe.FormatURLPath("/v1/product_catalog/trial_offers/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ProductCatalogTrialOffer create raw response", err.Error())
		return
	}

	if err := flattenProductCatalogTrialOffer(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening ProductCatalogTrialOffer create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ProductCatalogTrialOfferResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState ProductCatalogTrialOfferResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ProductCatalogTrialOfferResourceModel
	state = priorState

	obj, err := r.client.V1ProductCatalogTrialOffers.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading ProductCatalogTrialOffer", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ProductCatalogTrialOffers.B, r.client.V1ProductCatalogTrialOffers.Key, stripe.FormatURLPath("/v1/product_catalog/trial_offers/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ProductCatalogTrialOffer raw response", err.Error())
		return
	}

	if err := flattenProductCatalogTrialOffer(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening ProductCatalogTrialOffer read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *ProductCatalogTrialOfferResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ProductCatalogTrialOfferResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ProductCatalogTrialOfferResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandProductCatalogTrialOfferUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building ProductCatalogTrialOffer update params", err.Error())
		return
	}

	obj, err := r.client.V1ProductCatalogTrialOffers.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating ProductCatalogTrialOffer", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ProductCatalogTrialOffers.B, r.client.V1ProductCatalogTrialOffers.Key, stripe.FormatURLPath("/v1/product_catalog/trial_offers/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ProductCatalogTrialOffer update raw response", err.Error())
		return
	}

	if err := flattenProductCatalogTrialOffer(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening ProductCatalogTrialOffer update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ProductCatalogTrialOfferResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ProductCatalogTrialOfferResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.Active.IsNull() && !state.Active.IsUnknown() && !state.Active.ValueBool() {
		return
	}

	params := &stripe.ProductCatalogTrialOfferUpdateParams{}
	activeField := reflect.ValueOf(params).Elem().FieldByName("Active")
	if activeField.IsValid() && activeField.CanSet() {
		if activeField.Kind() == reflect.Pointer && activeField.Type().Elem().Kind() == reflect.Bool {
			activeField.Set(reflect.ValueOf(stripe.Bool(false)))
		} else if activeField.Kind() == reflect.Bool {
			activeField.SetBool(false)
		}
	}

	_, err := r.client.V1ProductCatalogTrialOffers.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating ProductCatalogTrialOffer", err.Error())
		return
	}
}

func (r *ProductCatalogTrialOfferResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandProductCatalogTrialOfferCreate(plan ProductCatalogTrialOfferResourceModel) (*stripe.ProductCatalogTrialOfferCreateParams, error) {
	params := &stripe.ProductCatalogTrialOfferCreateParams{}

	if !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.Duration.IsNull() && !plan.Duration.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Duration", plan.Duration) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "duration", params)
		}
	}
	if !plan.EndBehavior.IsNull() && !plan.EndBehavior.IsUnknown() {
		if !assignAttrValueToNamedField(params, "EndBehavior", plan.EndBehavior) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "end_behavior", params)
		}
	}
	if !plan.Nickname.IsNull() && !plan.Nickname.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Nickname", "Nickname", plan.Nickname.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "nickname", params)
		}
	}
	if !plan.Price.IsNull() && !plan.Price.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PriceID", "Price", plan.Price.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "price", params)
		}
	}

	return params, nil
}

func expandProductCatalogTrialOfferUpdate(plan ProductCatalogTrialOfferResourceModel, state ProductCatalogTrialOfferResourceModel) (*stripe.ProductCatalogTrialOfferUpdateParams, error) {
	params := &stripe.ProductCatalogTrialOfferUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}

	return params, nil
}

func flattenProductCatalogTrialOffer(obj *stripe.ProductCatalogTrialOffer, state *ProductCatalogTrialOfferResourceModel) error {
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
		if rawValueActive, rawOk := plainValueAtPath(raw, "active"); rawOk {
			if valueActive, err := flattenPlainValue(rawValueActive, types.BoolType, "active", "raw response"); err != nil {
				return err
			} else {
				if typedActive, ok := valueActive.(types.Bool); ok {
					state.Active = typedActive
				}
			}
		} else if !hasRaw {
			if responseValueActive, ok := plainFromResponseField(obj, "Active"); ok {
				if valueActive, err := flattenPlainValue(responseValueActive, types.BoolType, "active", "response struct"); err != nil {
					return err
				} else {
					if typedActive, ok := valueActive.(types.Bool); ok {
						state.Active = typedActive
					}
				}
			}
		}
	}
	{
		assignedDuration := false
		hadRawDuration := false
		if rawValueDuration, rawOk := plainValueAtPath(raw, "duration"); rawOk {
			hadRawDuration = true
			if rawValueDuration != nil {
				sourceDuration := applyConfiguredKeyedListShapes(rawValueDuration, attrValueToPlain(state.Duration))
				if valueDuration, err := flattenPlainValue(sourceDuration, types.ObjectType{AttrTypes: map[string]attr.Type{"relative": types.ObjectType{AttrTypes: map[string]attr.Type{"iterations": types.Int64Type}}, "type": types.StringType}}, "duration", "raw response"); err != nil {
					return err
				} else {
					if typedDuration, ok := valueDuration.(types.Object); ok {
						state.Duration = typedDuration
						assignedDuration = true
					}
				}
			}
		}
		if !assignedDuration {
			if !hasRaw {
				if responseValueDuration, ok := plainFromResponseField(obj, "Duration"); ok {
					sourceDuration := applyConfiguredKeyedListShapes(responseValueDuration, attrValueToPlain(state.Duration))
					if valueDuration, err := flattenPlainValue(
						sourceDuration,
						types.ObjectType{AttrTypes: map[string]attr.Type{"relative": types.ObjectType{AttrTypes: map[string]attr.Type{"iterations": types.Int64Type}}, "type": types.StringType}},
						"duration",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedDuration, ok := valueDuration.(types.Object); ok {
							state.Duration = typedDuration
							assignedDuration = true
						}
					}
				}
			}
		}
		if !assignedDuration && hadRawDuration {
			if nullDuration, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"relative": types.ObjectType{AttrTypes: map[string]attr.Type{"iterations": types.Int64Type}}, "type": types.StringType}}); ok {
				if typedDuration, ok := nullDuration.(types.Object); ok {
					state.Duration = typedDuration
				}
			}
		}
	}
	{
		assignedEndBehavior := false
		hadRawEndBehavior := false
		if rawValueEndBehavior, rawOk := plainValueAtPath(raw, "end_behavior"); rawOk {
			hadRawEndBehavior = true
			if rawValueEndBehavior != nil {
				sourceEndBehavior := applyConfiguredKeyedListShapes(rawValueEndBehavior, attrValueToPlain(state.EndBehavior))
				if valueEndBehavior, err := flattenPlainValue(sourceEndBehavior, types.ObjectType{AttrTypes: map[string]attr.Type{"transition": types.ObjectType{AttrTypes: map[string]attr.Type{"price": types.StringType}}, "type": types.StringType}}, "end_behavior", "raw response"); err != nil {
					return err
				} else {
					if typedEndBehavior, ok := valueEndBehavior.(types.Object); ok {
						state.EndBehavior = typedEndBehavior
						assignedEndBehavior = true
					}
				}
			}
		}
		if !assignedEndBehavior {
			if !hasRaw {
				if responseValueEndBehavior, ok := plainFromResponseField(obj, "EndBehavior"); ok {
					sourceEndBehavior := applyConfiguredKeyedListShapes(responseValueEndBehavior, attrValueToPlain(state.EndBehavior))
					if valueEndBehavior, err := flattenPlainValue(
						sourceEndBehavior,
						types.ObjectType{AttrTypes: map[string]attr.Type{"transition": types.ObjectType{AttrTypes: map[string]attr.Type{"price": types.StringType}}, "type": types.StringType}},
						"end_behavior",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedEndBehavior, ok := valueEndBehavior.(types.Object); ok {
							state.EndBehavior = typedEndBehavior
							assignedEndBehavior = true
						}
					}
				}
			}
		}
		if !assignedEndBehavior && hadRawEndBehavior {
			if nullEndBehavior, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"transition": types.ObjectType{AttrTypes: map[string]attr.Type{"price": types.StringType}}, "type": types.StringType}}); ok {
				if typedEndBehavior, ok := nullEndBehavior.(types.Object); ok {
					state.EndBehavior = typedEndBehavior
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
		if rawValueNickname, rawOk := plainValueAtPath(raw, "nickname"); rawOk {
			if valueNickname, err := flattenPlainValue(rawValueNickname, types.StringType, "nickname", "raw response"); err != nil {
				return err
			} else {
				if typedNickname, ok := valueNickname.(types.String); ok {
					state.Nickname = typedNickname
				}
			}
		} else if !hasRaw {
			if responseValueNickname, ok := plainFromResponseField(obj, "Nickname"); ok {
				if valueNickname, err := flattenPlainValue(responseValueNickname, types.StringType, "nickname", "response struct"); err != nil {
					return err
				} else {
					if typedNickname, ok := valueNickname.(types.String); ok {
						state.Nickname = typedNickname
					}
				}
			}
		}
	}
	{
		if state.Price.IsNull() || state.Price.IsUnknown() {
			if rawValuePrice, rawOk := plainValueAtPath(raw, "price"); rawOk {
				if typedPrice, ok := plainToStringIDValue(rawValuePrice); ok {
					state.Price = typedPrice
				}
			} else if !hasRaw {
				if responseValuePrice, ok := plainFromResponseField(obj, "Price"); ok {
					if typedPrice, ok := plainToStringIDValue(responseValuePrice); ok {
						state.Price = typedPrice
					}
				}
			}
		}
	}
	return nil
}
