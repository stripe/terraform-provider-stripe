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

var _ resource.Resource = &BillingCreditGrantResource{}

var _ resource.ResourceWithConfigure = &BillingCreditGrantResource{}

var _ resource.ResourceWithImportState = &BillingCreditGrantResource{}

func NewBillingCreditGrantResource() resource.Resource {
	return &BillingCreditGrantResource{}
}

type BillingCreditGrantResource struct {
	client *stripe.Client
}

type BillingCreditGrantResourceModel struct {
	Object              types.String `tfsdk:"object"`
	Amount              types.Object `tfsdk:"amount"`
	ApplicabilityConfig types.Object `tfsdk:"applicability_config"`
	Category            types.String `tfsdk:"category"`
	Created             types.Int64  `tfsdk:"created"`
	Customer            types.String `tfsdk:"customer"`
	CustomerAccount     types.String `tfsdk:"customer_account"`
	EffectiveAt         types.Int64  `tfsdk:"effective_at"`
	ExpiresAt           types.Int64  `tfsdk:"expires_at"`
	ID                  types.String `tfsdk:"id"`
	Livemode            types.Bool   `tfsdk:"livemode"`
	Metadata            types.Map    `tfsdk:"metadata"`
	Name                types.String `tfsdk:"name"`
	Priority            types.Int64  `tfsdk:"priority"`
	TestClock           types.String `tfsdk:"test_clock"`
	Updated             types.Int64  `tfsdk:"updated"`
	VoidedAt            types.Int64  `tfsdk:"voided_at"`
}

func (r *BillingCreditGrantResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *BillingCreditGrantResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_billing_credit_grant"
}

func (r *BillingCreditGrantResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A credit grant is an API resource that documents the allocation of some billing credits to a customer.\n\nRelated guide: [Billing credits](https://docs.stripe.com/billing/subscriptions/usage-based/billing-credits)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("billing.credit_grant")},
			},
			"amount": schema.SingleNestedAttribute{
				Required: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"monetary": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The monetary amount.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"currency": schema.StringAttribute{
								Required:      true,
								Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
							"value": schema.Int64Attribute{
								Required:      true,
								Description:   "A positive integer representing the amount.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
							},
						},
					},
					"type": schema.StringAttribute{
						Required:      true,
						Description:   "The type of this amount. We currently only support `monetary` billing credits.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("monetary")},
					},
				},
			},
			"applicability_config": schema.SingleNestedAttribute{
				Required: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"scope": schema.SingleNestedAttribute{
						Required: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"price_type": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The price type that credit grants can apply to. We currently only support the `metered` price type. This refers to prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them. Cannot be used in combination with `prices`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("metered")},
							},
							"prices": schema.ListNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The prices that credit grants can apply to. We currently only support `metered` prices. This refers to prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them. Cannot be used in combination with `price_type`.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed:      true,
											Description:   "Unique identifier for the object.",
											PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
										},
									},
								},
							},
						},
					},
				},
			},
			"category": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The category of this credit grant. This is for tracking purposes and isn't displayed to the customer.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("paid", "promotional")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"customer": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the customer receiving the billing credits.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"customer_account": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "ID of the account representing the customer receiving the billing credits",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"effective_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The time when the billing credits become effective-when they're eligible for use.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"expires_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The time when the billing credits expire. If not present, the billing credits don't expire.",
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
				Description:   "A descriptive name shown in dashboard.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"priority": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The priority for applying this credit grant. The highest priority is 0 and the lowest is 100.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"test_clock": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the test clock this credit grant belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"updated": schema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was last updated. Measured in seconds since the Unix epoch.",
			},
			"voided_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "The time when this credit grant was voided. If not present, the credit grant hasn't been voided.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *BillingCreditGrantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan BillingCreditGrantResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandBillingCreditGrantCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building BillingCreditGrant create params", err.Error())
		return
	}

	obj, err := r.client.V1BillingCreditGrants.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating BillingCreditGrant", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingCreditGrants.B, r.client.V1BillingCreditGrants.Key, stripe.FormatURLPath("/v1/billing/credit_grants/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingCreditGrant create raw response", err.Error())
		return
	}

	if err := flattenBillingCreditGrant(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingCreditGrant create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingCreditGrantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState BillingCreditGrantResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state BillingCreditGrantResourceModel
	state = priorState

	obj, err := r.client.V1BillingCreditGrants.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading BillingCreditGrant", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingCreditGrants.B, r.client.V1BillingCreditGrants.Key, stripe.FormatURLPath("/v1/billing/credit_grants/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingCreditGrant raw response", err.Error())
		return
	}

	if err := flattenBillingCreditGrant(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingCreditGrant read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingCreditGrantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan BillingCreditGrantResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state BillingCreditGrantResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandBillingCreditGrantUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building BillingCreditGrant update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building BillingCreditGrant update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1BillingCreditGrants.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating BillingCreditGrant", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingCreditGrants.B, r.client.V1BillingCreditGrants.Key, stripe.FormatURLPath("/v1/billing/credit_grants/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingCreditGrant update raw response", err.Error())
		return
	}

	if err := flattenBillingCreditGrant(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingCreditGrant update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingCreditGrantResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *BillingCreditGrantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandBillingCreditGrantCreate(plan BillingCreditGrantResourceModel) (*stripe.BillingCreditGrantCreateParams, error) {
	params := &stripe.BillingCreditGrantCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Amount", plan.Amount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "amount", params)
		}
	}
	if !plan.ApplicabilityConfig.IsNull() && !plan.ApplicabilityConfig.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ApplicabilityConfig", plan.ApplicabilityConfig) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "applicability_config", params)
		}
	}
	if !plan.Category.IsNull() && !plan.Category.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Category", "Category", plan.Category.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "category", params)
		}
	}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerID", "Customer", plan.Customer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer", params)
		}
	}
	if !plan.CustomerAccount.IsNull() && !plan.CustomerAccount.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "CustomerAccount", "CustomerAccount", plan.CustomerAccount.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer_account", params)
		}
	}
	if !plan.EffectiveAt.IsNull() && !plan.EffectiveAt.IsUnknown() {
		params.EffectiveAt = stripe.Int64(plan.EffectiveAt.ValueInt64())
	}
	if !plan.ExpiresAt.IsNull() && !plan.ExpiresAt.IsUnknown() {
		params.ExpiresAt = stripe.Int64(plan.ExpiresAt.ValueInt64())
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
	if !plan.Priority.IsNull() && !plan.Priority.IsUnknown() {
		params.Priority = stripe.Int64(plan.Priority.ValueInt64())
	}

	return params, nil
}

func expandBillingCreditGrantUpdate(plan BillingCreditGrantResourceModel, state BillingCreditGrantResourceModel) (*stripe.BillingCreditGrantUpdateParams, error) {
	params := &stripe.BillingCreditGrantUpdateParams{}

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

func flattenBillingCreditGrant(obj *stripe.BillingCreditGrant, state *BillingCreditGrantResourceModel) error {
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
		assignedAmount := false
		hadRawAmount := false
		if rawValueAmount, rawOk := plainValueAtPath(raw, "amount"); rawOk {
			hadRawAmount = true
			if rawValueAmount != nil {
				sourceAmount := applyConfiguredKeyedListShapes(rawValueAmount, attrValueToPlain(state.Amount))
				if valueAmount, err := flattenPlainValue(sourceAmount, types.ObjectType{AttrTypes: map[string]attr.Type{"monetary": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "type": types.StringType}}, "amount", "raw response"); err != nil {
					return err
				} else {
					if typedAmount, ok := valueAmount.(types.Object); ok {
						state.Amount = typedAmount
						assignedAmount = true
					}
				}
			}
		}
		if !assignedAmount {
			if !hasRaw {
				if responseValueAmount, ok := plainFromResponseField(obj, "Amount"); ok {
					sourceAmount := applyConfiguredKeyedListShapes(responseValueAmount, attrValueToPlain(state.Amount))
					if valueAmount, err := flattenPlainValue(
						sourceAmount,
						types.ObjectType{AttrTypes: map[string]attr.Type{"monetary": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "type": types.StringType}},
						"amount",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAmount, ok := valueAmount.(types.Object); ok {
							state.Amount = typedAmount
							assignedAmount = true
						}
					}
				}
			}
		}
		if !assignedAmount && hadRawAmount {
			if nullAmount, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"monetary": types.ObjectType{AttrTypes: map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}}, "type": types.StringType}}); ok {
				if typedAmount, ok := nullAmount.(types.Object); ok {
					state.Amount = typedAmount
				}
			}
		}
	}
	{
		assignedApplicabilityConfig := false
		hadRawApplicabilityConfig := false
		if rawValueApplicabilityConfig, rawOk := plainValueAtPath(raw, "applicability_config"); rawOk {
			hadRawApplicabilityConfig = true
			if rawValueApplicabilityConfig != nil {
				sourceApplicabilityConfig := applyConfiguredKeyedListShapes(rawValueApplicabilityConfig, attrValueToPlain(state.ApplicabilityConfig))
				if valueApplicabilityConfig, err := flattenPlainValue(sourceApplicabilityConfig, types.ObjectType{AttrTypes: map[string]attr.Type{"scope": types.ObjectType{AttrTypes: map[string]attr.Type{"price_type": types.StringType, "prices": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType}}}}}}}, "applicability_config", "raw response"); err != nil {
					return err
				} else {
					if typedApplicabilityConfig, ok := valueApplicabilityConfig.(types.Object); ok {
						state.ApplicabilityConfig = typedApplicabilityConfig
						assignedApplicabilityConfig = true
					}
				}
			}
		}
		if !assignedApplicabilityConfig {
			if !hasRaw {
				if responseValueApplicabilityConfig, ok := plainFromResponseField(obj, "ApplicabilityConfig"); ok {
					sourceApplicabilityConfig := applyConfiguredKeyedListShapes(responseValueApplicabilityConfig, attrValueToPlain(state.ApplicabilityConfig))
					if valueApplicabilityConfig, err := flattenPlainValue(
						sourceApplicabilityConfig,
						types.ObjectType{AttrTypes: map[string]attr.Type{"scope": types.ObjectType{AttrTypes: map[string]attr.Type{"price_type": types.StringType, "prices": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType}}}}}}},
						"applicability_config",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedApplicabilityConfig, ok := valueApplicabilityConfig.(types.Object); ok {
							state.ApplicabilityConfig = typedApplicabilityConfig
							assignedApplicabilityConfig = true
						}
					}
				}
			}
		}
		if !assignedApplicabilityConfig && hadRawApplicabilityConfig {
			if nullApplicabilityConfig, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"scope": types.ObjectType{AttrTypes: map[string]attr.Type{"price_type": types.StringType, "prices": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType}}}}}}}); ok {
				if typedApplicabilityConfig, ok := nullApplicabilityConfig.(types.Object); ok {
					state.ApplicabilityConfig = typedApplicabilityConfig
				}
			}
		}
	}
	{
		if rawValueCategory, rawOk := plainValueAtPath(raw, "category"); rawOk {
			if valueCategory, err := flattenPlainValue(rawValueCategory, types.StringType, "category", "raw response"); err != nil {
				return err
			} else {
				if typedCategory, ok := valueCategory.(types.String); ok {
					state.Category = typedCategory
				}
			}
		} else if !hasRaw {
			if responseValueCategory, ok := plainFromResponseField(obj, "Category"); ok {
				if valueCategory, err := flattenPlainValue(responseValueCategory, types.StringType, "category", "response struct"); err != nil {
					return err
				} else {
					if typedCategory, ok := valueCategory.(types.String); ok {
						state.Category = typedCategory
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
		if state.Customer.IsNull() || state.Customer.IsUnknown() {
			if rawValueCustomer, rawOk := plainValueAtPath(raw, "customer"); rawOk {
				if typedCustomer, ok := plainToStringIDValue(rawValueCustomer); ok {
					state.Customer = typedCustomer
				}
			} else if !hasRaw {
				if responseValueCustomer, ok := plainFromResponseField(obj, "Customer"); ok {
					if typedCustomer, ok := plainToStringIDValue(responseValueCustomer); ok {
						state.Customer = typedCustomer
					}
				}
			}
		}
	}
	{
		if rawValueCustomerAccount, rawOk := plainValueAtPath(raw, "customer_account"); rawOk {
			if valueCustomerAccount, err := flattenPlainValue(rawValueCustomerAccount, types.StringType, "customer_account", "raw response"); err != nil {
				return err
			} else {
				if typedCustomerAccount, ok := valueCustomerAccount.(types.String); ok {
					state.CustomerAccount = typedCustomerAccount
				}
			}
		} else if !hasRaw {
			if responseValueCustomerAccount, ok := plainFromResponseField(obj, "CustomerAccount"); ok {
				if valueCustomerAccount, err := flattenPlainValue(responseValueCustomerAccount, types.StringType, "customer_account", "response struct"); err != nil {
					return err
				} else {
					if typedCustomerAccount, ok := valueCustomerAccount.(types.String); ok {
						state.CustomerAccount = typedCustomerAccount
					}
				}
			}
		}
	}
	{
		if rawValueEffectiveAt, rawOk := plainValueAtPath(raw, "effective_at"); rawOk {
			if valueEffectiveAt, err := flattenPlainValue(rawValueEffectiveAt, types.Int64Type, "effective_at", "raw response"); err != nil {
				return err
			} else {
				if typedEffectiveAt, ok := valueEffectiveAt.(types.Int64); ok {
					state.EffectiveAt = typedEffectiveAt
				}
			}
		} else if !hasRaw {
			if responseValueEffectiveAt, ok := plainFromResponseField(obj, "EffectiveAt"); ok {
				if valueEffectiveAt, err := flattenPlainValue(responseValueEffectiveAt, types.Int64Type, "effective_at", "response struct"); err != nil {
					return err
				} else {
					if typedEffectiveAt, ok := valueEffectiveAt.(types.Int64); ok {
						state.EffectiveAt = typedEffectiveAt
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
		if rawValuePriority, rawOk := plainValueAtPath(raw, "priority"); rawOk {
			if valuePriority, err := flattenPlainValue(rawValuePriority, types.Int64Type, "priority", "raw response"); err != nil {
				return err
			} else {
				if typedPriority, ok := valuePriority.(types.Int64); ok {
					state.Priority = typedPriority
				}
			}
		} else if !hasRaw {
			if responseValuePriority, ok := plainFromResponseField(obj, "Priority"); ok {
				if valuePriority, err := flattenPlainValue(responseValuePriority, types.Int64Type, "priority", "response struct"); err != nil {
					return err
				} else {
					if typedPriority, ok := valuePriority.(types.Int64); ok {
						state.Priority = typedPriority
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueTestClock, rawOk := plainValueAtPath(raw, "test_clock"); rawOk {
				if typedTestClock, ok := plainToStringIDValue(rawValueTestClock); ok {
					state.TestClock = typedTestClock
				}
			} else if !hasRaw {
				if responseValueTestClock, ok := plainFromResponseField(obj, "TestClock"); ok {
					if typedTestClock, ok := plainToStringIDValue(responseValueTestClock); ok {
						state.TestClock = typedTestClock
					}
				}
			}
		}
	}
	{
		if rawValueUpdated, rawOk := plainValueAtPath(raw, "updated"); rawOk {
			if valueUpdated, err := flattenPlainValue(rawValueUpdated, types.Int64Type, "updated", "raw response"); err != nil {
				return err
			} else {
				if typedUpdated, ok := valueUpdated.(types.Int64); ok {
					state.Updated = typedUpdated
				}
			}
		} else if !hasRaw {
			if responseValueUpdated, ok := plainFromResponseField(obj, "Updated"); ok {
				if valueUpdated, err := flattenPlainValue(responseValueUpdated, types.Int64Type, "updated", "response struct"); err != nil {
					return err
				} else {
					if typedUpdated, ok := valueUpdated.(types.Int64); ok {
						state.Updated = typedUpdated
					}
				}
			}
		}
	}
	{
		if rawValueVoidedAt, rawOk := plainValueAtPath(raw, "voided_at"); rawOk {
			if valueVoidedAt, err := flattenPlainValue(rawValueVoidedAt, types.Int64Type, "voided_at", "raw response"); err != nil {
				return err
			} else {
				if typedVoidedAt, ok := valueVoidedAt.(types.Int64); ok {
					state.VoidedAt = typedVoidedAt
				}
			}
		} else if !hasRaw {
			if responseValueVoidedAt, ok := plainFromResponseField(obj, "VoidedAt"); ok {
				if valueVoidedAt, err := flattenPlainValue(responseValueVoidedAt, types.Int64Type, "voided_at", "response struct"); err != nil {
					return err
				} else {
					if typedVoidedAt, ok := valueVoidedAt.(types.Int64); ok {
						state.VoidedAt = typedVoidedAt
					}
				}
			}
		}
	}
	return nil
}
