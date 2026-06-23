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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
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

var _ resource.Resource = &ClimateOrderResource{}

var _ resource.ResourceWithConfigure = &ClimateOrderResource{}

var _ resource.ResourceWithImportState = &ClimateOrderResource{}

func NewClimateOrderResource() resource.Resource {
	return &ClimateOrderResource{}
}

type ClimateOrderResource struct {
	client *stripe.Client
}

type ClimateOrderResourceModel struct {
	Object               types.String  `tfsdk:"object"`
	AmountFees           types.Int64   `tfsdk:"amount_fees"`
	AmountSubtotal       types.Int64   `tfsdk:"amount_subtotal"`
	AmountTotal          types.Int64   `tfsdk:"amount_total"`
	Beneficiary          types.Object  `tfsdk:"beneficiary"`
	CanceledAt           types.Int64   `tfsdk:"canceled_at"`
	CancellationReason   types.String  `tfsdk:"cancellation_reason"`
	Certificate          types.String  `tfsdk:"certificate"`
	ConfirmedAt          types.Int64   `tfsdk:"confirmed_at"`
	Created              types.Int64   `tfsdk:"created"`
	Currency             types.String  `tfsdk:"currency"`
	DelayedAt            types.Int64   `tfsdk:"delayed_at"`
	DeliveredAt          types.Int64   `tfsdk:"delivered_at"`
	DeliveryDetails      types.List    `tfsdk:"delivery_details"`
	ExpectedDeliveryYear types.Int64   `tfsdk:"expected_delivery_year"`
	ID                   types.String  `tfsdk:"id"`
	Livemode             types.Bool    `tfsdk:"livemode"`
	Metadata             types.Map     `tfsdk:"metadata"`
	MetricTons           types.Float64 `tfsdk:"metric_tons"`
	Product              types.String  `tfsdk:"product"`
	ProductSubstitutedAt types.Int64   `tfsdk:"product_substituted_at"`
	Status               types.String  `tfsdk:"status"`
	Amount               types.Int64   `tfsdk:"amount"`
}

func (r *ClimateOrderResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ClimateOrderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_climate_order"
}

func (r *ClimateOrderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Orders represent your intent to purchase a particular Climate product. When you create an order, the\npayment is deducted from your merchant balance.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("climate.order")},
			},
			"amount_fees": schema.Int64Attribute{
				Computed:      true,
				Description:   "Total amount of [Frontier](https://frontierclimate.com/)'s service fees in the currency's smallest unit.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"amount_subtotal": schema.Int64Attribute{
				Computed:      true,
				Description:   "Total amount of the carbon removal in the currency's smallest unit.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"amount_total": schema.Int64Attribute{
				Computed:      true,
				Description:   "Total amount of the order including fees in the currency's smallest unit.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"beneficiary": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"public_name": schema.StringAttribute{
						Required:    true,
						Description: "Publicly displayable name for the end beneficiary of carbon removal.",
					},
				},
			},
			"canceled_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the order was canceled. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"cancellation_reason": schema.StringAttribute{
				Computed:      true,
				Description:   "Reason for the cancellation of this order.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("expired", "product_unavailable", "requested")},
			},
			"certificate": schema.StringAttribute{
				Computed:      true,
				Description:   "For delivered orders, a URL to a delivery certificate for the order.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"confirmed_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the order was confirmed. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase, representing the currency for this order.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"delayed_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the order's expected_delivery_year was delayed. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"delivered_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the order was delivered. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"delivery_details": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "Details about the delivery of carbon removal for this order.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"delivered_at": schema.Int64Attribute{
							Computed:      true,
							Description:   "Time at which the delivery occurred. Measured in seconds since the Unix epoch.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
						},
						"location": schema.SingleNestedAttribute{
							Computed:      true,
							Description:   "Specific location of this delivery.",
							PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
							Attributes: map[string]schema.Attribute{
								"city": schema.StringAttribute{
									Computed:      true,
									Description:   "The city where the supplier is located.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"country": schema.StringAttribute{
									Computed:      true,
									Description:   "Two-letter ISO code representing the country where the supplier is located.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"latitude": schema.Float64Attribute{
									Computed:      true,
									Description:   "The geographic latitude where the supplier is located.",
									PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
								},
								"longitude": schema.Float64Attribute{
									Computed:      true,
									Description:   "The geographic longitude where the supplier is located.",
									PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
								},
								"region": schema.StringAttribute{
									Computed:      true,
									Description:   "The state/county/province/region where the supplier is located.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
							},
						},
						"metric_tons": schema.StringAttribute{
							Computed:      true,
							Description:   "Quantity of carbon removal supplied by this delivery.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"registry_url": schema.StringAttribute{
							Computed:      true,
							Description:   "Once retired, a URL to the registry entry for the tons from this delivery.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
						"supplier": schema.StringAttribute{
							Computed:      true,
							Description:   "A supplier of carbon removal.",
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						},
					},
				},
			},
			"expected_delivery_year": schema.Int64Attribute{
				Computed:      true,
				Description:   "The year this order is expected to be delivered.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"metric_tons": schema.Float64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Quantity of carbon removal that is included in this order.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown(), float64planmodifier.RequiresReplace()},
			},
			"product": schema.StringAttribute{
				Required:      true,
				Description:   "Unique ID for the Climate `Product` this order is purchasing.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"product_substituted_at": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the order's product was substituted for a different product. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The current status of this order.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("awaiting_funds", "canceled", "confirmed", "delivered", "open")},
			},
			"amount": schema.Int64Attribute{
				Optional:      true,
				Description:   "Requested amount of carbon removal units. Either this or `metric_tons` must be specified.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
		},
	}
}

func (r *ClimateOrderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ClimateOrderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config ClimateOrderResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Amount"}})

	params, err := expandClimateOrderCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building ClimateOrder create params", err.Error())
		return
	}

	obj, err := r.client.V1ClimateOrders.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating ClimateOrder", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ClimateOrders.B, r.client.V1ClimateOrders.Key, stripe.FormatURLPath("/v1/climate/orders/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ClimateOrder create raw response", err.Error())
		return
	}

	if err := flattenClimateOrder(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening ClimateOrder create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Amount"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ClimateOrderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState ClimateOrderResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state ClimateOrderResourceModel
	state = priorState

	obj, err := r.client.V1ClimateOrders.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading ClimateOrder", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ClimateOrders.B, r.client.V1ClimateOrders.Key, stripe.FormatURLPath("/v1/climate/orders/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ClimateOrder raw response", err.Error())
		return
	}

	if err := flattenClimateOrder(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening ClimateOrder read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"Amount"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *ClimateOrderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ClimateOrderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config ClimateOrderResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Amount"}})

	var state ClimateOrderResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"Amount"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"Amount"}})

	params, err := expandClimateOrderUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building ClimateOrder update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building ClimateOrder update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1ClimateOrders.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating ClimateOrder", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1ClimateOrders.B, r.client.V1ClimateOrders.Key, stripe.FormatURLPath("/v1/climate/orders/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating ClimateOrder update raw response", err.Error())
		return
	}

	if err := flattenClimateOrder(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening ClimateOrder update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Amount"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *ClimateOrderResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *ClimateOrderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandClimateOrderCreate(plan ClimateOrderResourceModel) (*stripe.ClimateOrderCreateParams, error) {
	params := &stripe.ClimateOrderCreateParams{}

	if !plan.Beneficiary.IsNull() && !plan.Beneficiary.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Beneficiary", plan.Beneficiary) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "beneficiary", params)
		}
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.MetricTons.IsNull() && !plan.MetricTons.IsUnknown() {
		params.MetricTons = stripe.Float64(plan.MetricTons.ValueFloat64())
	}
	if !plan.Product.IsNull() && !plan.Product.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ProductID", "Product", plan.Product.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "product", params)
		}
	}
	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}

	return params, nil
}

func expandClimateOrderUpdate(plan ClimateOrderResourceModel, state ClimateOrderResourceModel) (*stripe.ClimateOrderUpdateParams, error) {
	params := &stripe.ClimateOrderUpdateParams{}

	if !plan.Beneficiary.Equal(state.Beneficiary) && !plan.Beneficiary.IsNull() && !plan.Beneficiary.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Beneficiary", plan.Beneficiary) {
			if !plan.Beneficiary.Equal(state.Beneficiary) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "beneficiary", params)
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

	return params, nil
}

func flattenClimateOrder(obj *stripe.ClimateOrder, state *ClimateOrderResourceModel) error {
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
		if rawValueAmountFees, rawOk := plainValueAtPath(raw, "amount_fees"); rawOk {
			if valueAmountFees, err := flattenPlainValue(rawValueAmountFees, types.Int64Type, "amount_fees", "raw response"); err != nil {
				return err
			} else {
				if typedAmountFees, ok := valueAmountFees.(types.Int64); ok {
					state.AmountFees = typedAmountFees
				}
			}
		} else if !hasRaw {
			if responseValueAmountFees, ok := plainFromResponseField(obj, "AmountFees"); ok {
				if valueAmountFees, err := flattenPlainValue(responseValueAmountFees, types.Int64Type, "amount_fees", "response struct"); err != nil {
					return err
				} else {
					if typedAmountFees, ok := valueAmountFees.(types.Int64); ok {
						state.AmountFees = typedAmountFees
					}
				}
			}
		}
	}
	{
		if rawValueAmountSubtotal, rawOk := plainValueAtPath(raw, "amount_subtotal"); rawOk {
			if valueAmountSubtotal, err := flattenPlainValue(rawValueAmountSubtotal, types.Int64Type, "amount_subtotal", "raw response"); err != nil {
				return err
			} else {
				if typedAmountSubtotal, ok := valueAmountSubtotal.(types.Int64); ok {
					state.AmountSubtotal = typedAmountSubtotal
				}
			}
		} else if !hasRaw {
			if responseValueAmountSubtotal, ok := plainFromResponseField(obj, "AmountSubtotal"); ok {
				if valueAmountSubtotal, err := flattenPlainValue(responseValueAmountSubtotal, types.Int64Type, "amount_subtotal", "response struct"); err != nil {
					return err
				} else {
					if typedAmountSubtotal, ok := valueAmountSubtotal.(types.Int64); ok {
						state.AmountSubtotal = typedAmountSubtotal
					}
				}
			}
		}
	}
	{
		if rawValueAmountTotal, rawOk := plainValueAtPath(raw, "amount_total"); rawOk {
			if valueAmountTotal, err := flattenPlainValue(rawValueAmountTotal, types.Int64Type, "amount_total", "raw response"); err != nil {
				return err
			} else {
				if typedAmountTotal, ok := valueAmountTotal.(types.Int64); ok {
					state.AmountTotal = typedAmountTotal
				}
			}
		} else if !hasRaw {
			if responseValueAmountTotal, ok := plainFromResponseField(obj, "AmountTotal"); ok {
				if valueAmountTotal, err := flattenPlainValue(responseValueAmountTotal, types.Int64Type, "amount_total", "response struct"); err != nil {
					return err
				} else {
					if typedAmountTotal, ok := valueAmountTotal.(types.Int64); ok {
						state.AmountTotal = typedAmountTotal
					}
				}
			}
		}
	}
	{
		assignedBeneficiary := false
		hadRawBeneficiary := false
		if rawValueBeneficiary, rawOk := plainValueAtPath(raw, "beneficiary"); rawOk {
			hadRawBeneficiary = true
			if rawValueBeneficiary != nil {
				sourceBeneficiary := applyConfiguredKeyedListShapes(rawValueBeneficiary, attrValueToPlain(state.Beneficiary))
				if valueBeneficiary, err := flattenPlainValue(sourceBeneficiary, types.ObjectType{AttrTypes: map[string]attr.Type{"public_name": types.StringType}}, "beneficiary", "raw response"); err != nil {
					return err
				} else {
					if typedBeneficiary, ok := valueBeneficiary.(types.Object); ok {
						state.Beneficiary = typedBeneficiary
						assignedBeneficiary = true
					}
				}
			}
		}
		if !assignedBeneficiary {
			if !hasRaw {
				if responseValueBeneficiary, ok := plainFromResponseField(obj, "Beneficiary"); ok {
					sourceBeneficiary := applyConfiguredKeyedListShapes(responseValueBeneficiary, attrValueToPlain(state.Beneficiary))
					if valueBeneficiary, err := flattenPlainValue(
						sourceBeneficiary,
						types.ObjectType{AttrTypes: map[string]attr.Type{"public_name": types.StringType}},
						"beneficiary",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBeneficiary, ok := valueBeneficiary.(types.Object); ok {
							state.Beneficiary = typedBeneficiary
							assignedBeneficiary = true
						}
					}
				}
			}
		}
		if !assignedBeneficiary && hadRawBeneficiary {
			if nullBeneficiary, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"public_name": types.StringType}}); ok {
				if typedBeneficiary, ok := nullBeneficiary.(types.Object); ok {
					state.Beneficiary = typedBeneficiary
				}
			}
		}
	}
	{
		if rawValueCanceledAt, rawOk := plainValueAtPath(raw, "canceled_at"); rawOk {
			if valueCanceledAt, err := flattenPlainValue(rawValueCanceledAt, types.Int64Type, "canceled_at", "raw response"); err != nil {
				return err
			} else {
				if typedCanceledAt, ok := valueCanceledAt.(types.Int64); ok {
					state.CanceledAt = typedCanceledAt
				}
			}
		} else if !hasRaw {
			if responseValueCanceledAt, ok := plainFromResponseField(obj, "CanceledAt"); ok {
				if valueCanceledAt, err := flattenPlainValue(responseValueCanceledAt, types.Int64Type, "canceled_at", "response struct"); err != nil {
					return err
				} else {
					if typedCanceledAt, ok := valueCanceledAt.(types.Int64); ok {
						state.CanceledAt = typedCanceledAt
					}
				}
			}
		}
	}
	{
		if rawValueCancellationReason, rawOk := plainValueAtPath(raw, "cancellation_reason"); rawOk {
			if valueCancellationReason, err := flattenPlainValue(rawValueCancellationReason, types.StringType, "cancellation_reason", "raw response"); err != nil {
				return err
			} else {
				if typedCancellationReason, ok := valueCancellationReason.(types.String); ok {
					state.CancellationReason = typedCancellationReason
				}
			}
		} else if !hasRaw {
			if responseValueCancellationReason, ok := plainFromResponseField(obj, "CancellationReason"); ok {
				if valueCancellationReason, err := flattenPlainValue(responseValueCancellationReason, types.StringType, "cancellation_reason", "response struct"); err != nil {
					return err
				} else {
					if typedCancellationReason, ok := valueCancellationReason.(types.String); ok {
						state.CancellationReason = typedCancellationReason
					}
				}
			}
		}
	}
	{
		if rawValueCertificate, rawOk := plainValueAtPath(raw, "certificate"); rawOk {
			if valueCertificate, err := flattenPlainValue(rawValueCertificate, types.StringType, "certificate", "raw response"); err != nil {
				return err
			} else {
				if typedCertificate, ok := valueCertificate.(types.String); ok {
					state.Certificate = typedCertificate
				}
			}
		} else if !hasRaw {
			if responseValueCertificate, ok := plainFromResponseField(obj, "Certificate"); ok {
				if valueCertificate, err := flattenPlainValue(responseValueCertificate, types.StringType, "certificate", "response struct"); err != nil {
					return err
				} else {
					if typedCertificate, ok := valueCertificate.(types.String); ok {
						state.Certificate = typedCertificate
					}
				}
			}
		}
	}
	{
		if rawValueConfirmedAt, rawOk := plainValueAtPath(raw, "confirmed_at"); rawOk {
			if valueConfirmedAt, err := flattenPlainValue(rawValueConfirmedAt, types.Int64Type, "confirmed_at", "raw response"); err != nil {
				return err
			} else {
				if typedConfirmedAt, ok := valueConfirmedAt.(types.Int64); ok {
					state.ConfirmedAt = typedConfirmedAt
				}
			}
		} else if !hasRaw {
			if responseValueConfirmedAt, ok := plainFromResponseField(obj, "ConfirmedAt"); ok {
				if valueConfirmedAt, err := flattenPlainValue(responseValueConfirmedAt, types.Int64Type, "confirmed_at", "response struct"); err != nil {
					return err
				} else {
					if typedConfirmedAt, ok := valueConfirmedAt.(types.Int64); ok {
						state.ConfirmedAt = typedConfirmedAt
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
		if rawValueCurrency, rawOk := plainValueAtPath(raw, "currency"); rawOk {
			if valueCurrency, err := flattenPlainValue(rawValueCurrency, types.StringType, "currency", "raw response"); err != nil {
				return err
			} else {
				if typedCurrency, ok := valueCurrency.(types.String); ok {
					state.Currency = typedCurrency
				}
			}
		} else if !hasRaw {
			if responseValueCurrency, ok := plainFromResponseField(obj, "Currency"); ok {
				if valueCurrency, err := flattenPlainValue(responseValueCurrency, types.StringType, "currency", "response struct"); err != nil {
					return err
				} else {
					if typedCurrency, ok := valueCurrency.(types.String); ok {
						state.Currency = typedCurrency
					}
				}
			}
		}
	}
	{
		if rawValueDelayedAt, rawOk := plainValueAtPath(raw, "delayed_at"); rawOk {
			if valueDelayedAt, err := flattenPlainValue(rawValueDelayedAt, types.Int64Type, "delayed_at", "raw response"); err != nil {
				return err
			} else {
				if typedDelayedAt, ok := valueDelayedAt.(types.Int64); ok {
					state.DelayedAt = typedDelayedAt
				}
			}
		} else if !hasRaw {
			if responseValueDelayedAt, ok := plainFromResponseField(obj, "DelayedAt"); ok {
				if valueDelayedAt, err := flattenPlainValue(responseValueDelayedAt, types.Int64Type, "delayed_at", "response struct"); err != nil {
					return err
				} else {
					if typedDelayedAt, ok := valueDelayedAt.(types.Int64); ok {
						state.DelayedAt = typedDelayedAt
					}
				}
			}
		}
	}
	{
		if rawValueDeliveredAt, rawOk := plainValueAtPath(raw, "delivered_at"); rawOk {
			if valueDeliveredAt, err := flattenPlainValue(rawValueDeliveredAt, types.Int64Type, "delivered_at", "raw response"); err != nil {
				return err
			} else {
				if typedDeliveredAt, ok := valueDeliveredAt.(types.Int64); ok {
					state.DeliveredAt = typedDeliveredAt
				}
			}
		} else if !hasRaw {
			if responseValueDeliveredAt, ok := plainFromResponseField(obj, "DeliveredAt"); ok {
				if valueDeliveredAt, err := flattenPlainValue(responseValueDeliveredAt, types.Int64Type, "delivered_at", "response struct"); err != nil {
					return err
				} else {
					if typedDeliveredAt, ok := valueDeliveredAt.(types.Int64); ok {
						state.DeliveredAt = typedDeliveredAt
					}
				}
			}
		}
	}
	{
		if rawValueDeliveryDetails, rawOk := plainValueAtPath(raw, "delivery_details"); rawOk {
			if valueDeliveryDetails, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueDeliveryDetails, attrValueToPlain(state.DeliveryDetails)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"delivered_at": types.Int64Type, "location": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "latitude": types.Float64Type, "longitude": types.Float64Type, "region": types.StringType}}, "metric_tons": types.StringType, "registry_url": types.StringType, "supplier": types.StringType}}}, "delivery_details", "raw response"); err != nil {
				return err
			} else {
				if typedDeliveryDetails, ok := valueDeliveryDetails.(types.List); ok {
					state.DeliveryDetails = typedDeliveryDetails
				}
			}
		} else if !hasRaw {
			if responseValueDeliveryDetails, ok := plainFromResponseField(obj, "DeliveryDetails"); ok {
				if valueDeliveryDetails, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueDeliveryDetails, attrValueToPlain(state.DeliveryDetails)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"delivered_at": types.Int64Type, "location": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "latitude": types.Float64Type, "longitude": types.Float64Type, "region": types.StringType}}, "metric_tons": types.StringType, "registry_url": types.StringType, "supplier": types.StringType}}},
					"delivery_details",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedDeliveryDetails, ok := valueDeliveryDetails.(types.List); ok {
						state.DeliveryDetails = typedDeliveryDetails
					}
				}
			}
		}
	}
	{
		if rawValueExpectedDeliveryYear, rawOk := plainValueAtPath(raw, "expected_delivery_year"); rawOk {
			if valueExpectedDeliveryYear, err := flattenPlainValue(rawValueExpectedDeliveryYear, types.Int64Type, "expected_delivery_year", "raw response"); err != nil {
				return err
			} else {
				if typedExpectedDeliveryYear, ok := valueExpectedDeliveryYear.(types.Int64); ok {
					state.ExpectedDeliveryYear = typedExpectedDeliveryYear
				}
			}
		} else if !hasRaw {
			if responseValueExpectedDeliveryYear, ok := plainFromResponseField(obj, "ExpectedDeliveryYear"); ok {
				if valueExpectedDeliveryYear, err := flattenPlainValue(responseValueExpectedDeliveryYear, types.Int64Type, "expected_delivery_year", "response struct"); err != nil {
					return err
				} else {
					if typedExpectedDeliveryYear, ok := valueExpectedDeliveryYear.(types.Int64); ok {
						state.ExpectedDeliveryYear = typedExpectedDeliveryYear
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
		if rawValueMetricTons, rawOk := plainValueAtPath(raw, "metric_tons"); rawOk {
			if valueMetricTons, err := flattenPlainValue(rawValueMetricTons, types.Float64Type, "metric_tons", "raw response"); err != nil {
				return err
			} else {
				if typedMetricTons, ok := valueMetricTons.(types.Float64); ok {
					state.MetricTons = typedMetricTons
				}
			}
		} else if !hasRaw {
			if responseValueMetricTons, ok := plainFromResponseField(obj, "MetricTons"); ok {
				if valueMetricTons, err := flattenPlainValue(responseValueMetricTons, types.Float64Type, "metric_tons", "response struct"); err != nil {
					return err
				} else {
					if typedMetricTons, ok := valueMetricTons.(types.Float64); ok {
						state.MetricTons = typedMetricTons
					}
				}
			}
		}
	}
	{
		if state.Product.IsNull() || state.Product.IsUnknown() {
			if rawValueProduct, rawOk := plainValueAtPath(raw, "product"); rawOk {
				if typedProduct, ok := plainToStringIDValue(rawValueProduct); ok {
					state.Product = typedProduct
				}
			} else if !hasRaw {
				if responseValueProduct, ok := plainFromResponseField(obj, "Product"); ok {
					if typedProduct, ok := plainToStringIDValue(responseValueProduct); ok {
						state.Product = typedProduct
					}
				}
			}
		}
	}
	{
		if rawValueProductSubstitutedAt, rawOk := plainValueAtPath(raw, "product_substituted_at"); rawOk {
			if valueProductSubstitutedAt, err := flattenPlainValue(rawValueProductSubstitutedAt, types.Int64Type, "product_substituted_at", "raw response"); err != nil {
				return err
			} else {
				if typedProductSubstitutedAt, ok := valueProductSubstitutedAt.(types.Int64); ok {
					state.ProductSubstitutedAt = typedProductSubstitutedAt
				}
			}
		} else if !hasRaw {
			if responseValueProductSubstitutedAt, ok := plainFromResponseField(obj, "ProductSubstitutedAt"); ok {
				if valueProductSubstitutedAt, err := flattenPlainValue(responseValueProductSubstitutedAt, types.Int64Type, "product_substituted_at", "response struct"); err != nil {
					return err
				} else {
					if typedProductSubstitutedAt, ok := valueProductSubstitutedAt.(types.Int64); ok {
						state.ProductSubstitutedAt = typedProductSubstitutedAt
					}
				}
			}
		}
	}
	{
		if rawValueStatus, rawOk := plainValueAtPath(raw, "status"); rawOk {
			if valueStatus, err := flattenPlainValue(rawValueStatus, types.StringType, "status", "raw response"); err != nil {
				return err
			} else {
				if typedStatus, ok := valueStatus.(types.String); ok {
					state.Status = typedStatus
				}
			}
		} else if !hasRaw {
			if responseValueStatus, ok := plainFromResponseField(obj, "Status"); ok {
				if valueStatus, err := flattenPlainValue(responseValueStatus, types.StringType, "status", "response struct"); err != nil {
					return err
				} else {
					if typedStatus, ok := valueStatus.(types.String); ok {
						state.Status = typedStatus
					}
				}
			}
		}
	}
	return nil
}
