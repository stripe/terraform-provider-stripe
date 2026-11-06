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

var _ resource.Resource = &PlanResource{}

var _ resource.ResourceWithConfigure = &PlanResource{}

var _ resource.ResourceWithImportState = &PlanResource{}

func NewPlanResource() resource.Resource {
	return &PlanResource{}
}

type PlanResource struct {
	client *stripe.Client
}

type PlanResourceModel struct {
	Object          types.String  `tfsdk:"object"`
	Active          types.Bool    `tfsdk:"active"`
	Amount          types.Int64   `tfsdk:"amount"`
	AmountDecimal   types.Float64 `tfsdk:"amount_decimal"`
	BillingScheme   types.String  `tfsdk:"billing_scheme"`
	Created         types.Int64   `tfsdk:"created"`
	Currency        types.String  `tfsdk:"currency"`
	ID              types.String  `tfsdk:"id"`
	Interval        types.String  `tfsdk:"interval"`
	IntervalCount   types.Int64   `tfsdk:"interval_count"`
	Livemode        types.Bool    `tfsdk:"livemode"`
	Metadata        types.Map     `tfsdk:"metadata"`
	Meter           types.String  `tfsdk:"meter"`
	Nickname        types.String  `tfsdk:"nickname"`
	Product         types.String  `tfsdk:"product"`
	Tiers           types.List    `tfsdk:"tiers"`
	TiersMode       types.String  `tfsdk:"tiers_mode"`
	TransformUsage  types.Object  `tfsdk:"transform_usage"`
	TrialPeriodDays types.Int64   `tfsdk:"trial_period_days"`
	UsageType       types.String  `tfsdk:"usage_type"`
}

func (r *PlanResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PlanResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plan"
}

func (r *PlanResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "You can now model subscriptions more flexibly using the [Prices API](https://api.stripe.com#prices). It replaces the Plans API and is backwards compatible to simplify your migration.\n\nPlans define the base price, currency, and billing cycle for recurring purchases of products.\n[Products](https://api.stripe.com#products) help you track inventory or provisioning, and plans help you track pricing. Different physical goods or levels of service should be represented by products, and pricing options should be represented by plans. This approach lets you change prices without having to change your provisioning scheme.\n\nFor example, you might have a single \"gold\" product that has plans for $10/month, $100/year, €9/month, and €90/year.\n\nRelated guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription) and more about [products and prices](https://docs.stripe.com/products-prices/overview).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("plan")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the plan can be used for new purchases.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
			},
			"amount_decimal": schema.Float64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The unit amount in cents (or local equivalent) to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.",
				PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown(), float64planmodifier.RequiresReplace()},
			},
			"billing_scheme": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `amount`) will be charged per unit in `quantity` (for plans with `usage_type=licensed`), or per unit of total usage (for plans with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("per_unit", "tiered")},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Required:      true,
				Description:   "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"interval": schema.StringAttribute{
				Required:      true,
				Description:   "The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("day", "month", "week", "year")},
			},
			"interval_count": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
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
			"meter": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The meter tracking the usage of a metered price",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"nickname": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A brief description of the plan, hidden from customers.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"product": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The product whose pricing this plan determines.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"tiers": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown(), listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"flat_amount": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "Price for the entire tier.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"flat_amount_decimal": schema.Float64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.",
							PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown(), float64planmodifier.RequiresReplace()},
						},
						"unit_amount": schema.Int64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "Per unit price for units relevant to the tier.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown(), int64planmodifier.RequiresReplace()},
						},
						"unit_amount_decimal": schema.Float64Attribute{
							Optional:      true,
							Computed:      true,
							Description:   "Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.",
							PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown(), float64planmodifier.RequiresReplace()},
						},
						"up_to": schema.Int64Attribute{
							Required:      true,
							Description:   "Up to and including to this quantity will be contained in the tier.",
							PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
						},
					},
				},
			},
			"tiers_mode": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price. In `graduated` tiering, pricing can change as the quantity grows.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("graduated", "volume")},
			},
			"transform_usage": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Apply a transformation to the reported usage or set quantity before computing the amount billed. Cannot be combined with `tiers`.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"divide_by": schema.Int64Attribute{
						Required:      true,
						Description:   "Divide usage by this number.",
						PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
					},
					"round": schema.StringAttribute{
						Required:      true,
						Description:   "After division, either round the result `up` or `down`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
						Validators:    []validator.String{stringvalidator.OneOf("down", "up")},
					},
				},
			},
			"trial_period_days": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Default number of trial days when subscribing a customer to this plan using [`trial_from_plan=true`](https://docs.stripe.com/api#create_subscription-trial_from_plan).",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"usage_type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("licensed", "metered")},
			},
		},
	}
}

func (r *PlanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PlanResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandPlanCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Plan create params", err.Error())
		return
	}

	obj, err := r.client.V1Plans.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Plan", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Plans.B, r.client.V1Plans.Key, stripe.FormatURLPath("/v1/plans/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Plan create raw response", err.Error())
		return
	}

	if err := flattenPlan(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Plan create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PlanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState PlanResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PlanResourceModel
	state = priorState

	obj, err := r.client.V1Plans.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Plan", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Plans.B, r.client.V1Plans.Key, stripe.FormatURLPath("/v1/plans/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Plan raw response", err.Error())
		return
	}

	if err := flattenPlan(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Plan read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *PlanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PlanResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PlanResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandPlanUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Plan update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Plan update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Plans.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Plan", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Plans.B, r.client.V1Plans.Key, stripe.FormatURLPath("/v1/plans/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating Plan update raw response", err.Error())
		return
	}

	if err := flattenPlan(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Plan update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PlanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PlanResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1Plans.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting Plan", err.Error())
		return
	}
}

func (r *PlanResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandPlanCreate(plan PlanResourceModel) (*stripe.PlanCreateParams, error) {
	params := &stripe.PlanCreateParams{}

	if !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.AmountDecimal.IsNull() && !plan.AmountDecimal.IsUnknown() {
		params.AmountDecimal = stripe.Float64(plan.AmountDecimal.ValueFloat64())
	}
	if !plan.BillingScheme.IsNull() && !plan.BillingScheme.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "BillingScheme", "BillingScheme", plan.BillingScheme.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_scheme", params)
		}
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.Interval.IsNull() && !plan.Interval.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Interval", "Interval", plan.Interval.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "interval", params)
		}
	}
	if !plan.IntervalCount.IsNull() && !plan.IntervalCount.IsUnknown() {
		params.IntervalCount = stripe.Int64(plan.IntervalCount.ValueInt64())
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Meter.IsNull() && !plan.Meter.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Meter", "Meter", plan.Meter.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "meter", params)
		}
	}
	if !plan.Nickname.IsNull() && !plan.Nickname.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Nickname", "Nickname", plan.Nickname.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "nickname", params)
		}
	}
	if !plan.Product.IsNull() && !plan.Product.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ProductID", "Product", plan.Product.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "product", params)
		}
	}
	if !plan.Tiers.IsNull() && !plan.Tiers.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Tiers", plan.Tiers) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tiers", params)
		}
	}
	if !plan.TiersMode.IsNull() && !plan.TiersMode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TiersMode", "TiersMode", plan.TiersMode.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tiers_mode", params)
		}
	}
	if !plan.TransformUsage.IsNull() && !plan.TransformUsage.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TransformUsage", plan.TransformUsage) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "transform_usage", params)
		}
	}
	if !plan.TrialPeriodDays.IsNull() && !plan.TrialPeriodDays.IsUnknown() {
		params.TrialPeriodDays = stripe.Int64(plan.TrialPeriodDays.ValueInt64())
	}
	if !plan.UsageType.IsNull() && !plan.UsageType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "UsageType", "UsageType", plan.UsageType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "usage_type", params)
		}
	}

	return params, nil
}

func expandPlanUpdate(plan PlanResourceModel, state PlanResourceModel) (*stripe.PlanUpdateParams, error) {
	params := &stripe.PlanUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.Metadata.Equal(state.Metadata) && !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			if !plan.Metadata.Equal(state.Metadata) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "metadata", params)
			}
		}
	}
	if !plan.Nickname.Equal(state.Nickname) && !plan.Nickname.IsNull() && !plan.Nickname.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Nickname", "Nickname", plan.Nickname.ValueString()) {
			if !plan.Nickname.Equal(state.Nickname) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "nickname", params)
			}
		}
	}
	if !plan.Product.Equal(state.Product) && !plan.Product.IsNull() && !plan.Product.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ProductID", "Product", plan.Product.ValueString()) {
			if !plan.Product.Equal(state.Product) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "product", params)
			}
		}
	}
	if !plan.TrialPeriodDays.Equal(state.TrialPeriodDays) && !plan.TrialPeriodDays.IsNull() && !plan.TrialPeriodDays.IsUnknown() {
		params.TrialPeriodDays = stripe.Int64(plan.TrialPeriodDays.ValueInt64())
	}

	return params, nil
}

func flattenPlan(obj *stripe.Plan, state *PlanResourceModel) error {
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
		if rawValueAmount, rawOk := plainValueAtPath(raw, "amount"); rawOk {
			if valueAmount, err := flattenPlainValue(rawValueAmount, types.Int64Type, "amount", "raw response"); err != nil {
				return err
			} else {
				if typedAmount, ok := valueAmount.(types.Int64); ok {
					state.Amount = typedAmount
				}
			}
		} else if !hasRaw {
			if responseValueAmount, ok := plainFromResponseField(obj, "Amount"); ok {
				if valueAmount, err := flattenPlainValue(responseValueAmount, types.Int64Type, "amount", "response struct"); err != nil {
					return err
				} else {
					if typedAmount, ok := valueAmount.(types.Int64); ok {
						state.Amount = typedAmount
					}
				}
			}
		}
	}
	{
		if rawValueAmountDecimal, rawOk := plainValueAtPath(raw, "amount_decimal"); rawOk {
			if valueAmountDecimal, err := flattenPlainValue(rawValueAmountDecimal, types.Float64Type, "amount_decimal", "raw response"); err != nil {
				return err
			} else {
				if typedAmountDecimal, ok := valueAmountDecimal.(types.Float64); ok {
					state.AmountDecimal = typedAmountDecimal
				}
			}
		} else if !hasRaw {
			if responseValueAmountDecimal, ok := plainFromResponseField(obj, "AmountDecimal"); ok {
				if valueAmountDecimal, err := flattenPlainValue(responseValueAmountDecimal, types.Float64Type, "amount_decimal", "response struct"); err != nil {
					return err
				} else {
					if typedAmountDecimal, ok := valueAmountDecimal.(types.Float64); ok {
						state.AmountDecimal = typedAmountDecimal
					}
				}
			}
		}
	}
	{
		if rawValueBillingScheme, rawOk := plainValueAtPath(raw, "billing_scheme"); rawOk {
			if valueBillingScheme, err := flattenPlainValue(rawValueBillingScheme, types.StringType, "billing_scheme", "raw response"); err != nil {
				return err
			} else {
				if typedBillingScheme, ok := valueBillingScheme.(types.String); ok {
					state.BillingScheme = typedBillingScheme
				}
			}
		} else if !hasRaw {
			if responseValueBillingScheme, ok := plainFromResponseField(obj, "BillingScheme"); ok {
				if valueBillingScheme, err := flattenPlainValue(responseValueBillingScheme, types.StringType, "billing_scheme", "response struct"); err != nil {
					return err
				} else {
					if typedBillingScheme, ok := valueBillingScheme.(types.String); ok {
						state.BillingScheme = typedBillingScheme
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
		if rawValueInterval, rawOk := plainValueAtPath(raw, "interval"); rawOk {
			if valueInterval, err := flattenPlainValue(rawValueInterval, types.StringType, "interval", "raw response"); err != nil {
				return err
			} else {
				if typedInterval, ok := valueInterval.(types.String); ok {
					state.Interval = typedInterval
				}
			}
		} else if !hasRaw {
			if responseValueInterval, ok := plainFromResponseField(obj, "Interval"); ok {
				if valueInterval, err := flattenPlainValue(responseValueInterval, types.StringType, "interval", "response struct"); err != nil {
					return err
				} else {
					if typedInterval, ok := valueInterval.(types.String); ok {
						state.Interval = typedInterval
					}
				}
			}
		}
	}
	{
		if rawValueIntervalCount, rawOk := plainValueAtPath(raw, "interval_count"); rawOk {
			if valueIntervalCount, err := flattenPlainValue(rawValueIntervalCount, types.Int64Type, "interval_count", "raw response"); err != nil {
				return err
			} else {
				if typedIntervalCount, ok := valueIntervalCount.(types.Int64); ok {
					state.IntervalCount = typedIntervalCount
				}
			}
		} else if !hasRaw {
			if responseValueIntervalCount, ok := plainFromResponseField(obj, "IntervalCount"); ok {
				if valueIntervalCount, err := flattenPlainValue(responseValueIntervalCount, types.Int64Type, "interval_count", "response struct"); err != nil {
					return err
				} else {
					if typedIntervalCount, ok := valueIntervalCount.(types.Int64); ok {
						state.IntervalCount = typedIntervalCount
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
		if rawValueMeter, rawOk := plainValueAtPath(raw, "meter"); rawOk {
			if valueMeter, err := flattenPlainValue(rawValueMeter, types.StringType, "meter", "raw response"); err != nil {
				return err
			} else {
				if typedMeter, ok := valueMeter.(types.String); ok {
					state.Meter = typedMeter
				}
			}
		} else if !hasRaw {
			if responseValueMeter, ok := plainFromResponseField(obj, "Meter"); ok {
				if valueMeter, err := flattenPlainValue(responseValueMeter, types.StringType, "meter", "response struct"); err != nil {
					return err
				} else {
					if typedMeter, ok := valueMeter.(types.String); ok {
						state.Meter = typedMeter
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
		if true {
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
		if rawValueTiers, rawOk := plainValueAtPath(raw, "tiers"); rawOk {
			if valueTiers, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawValueTiers, attrValueToPlain(state.Tiers)), attrValueToPlain(state.Tiers)), attrValueToPlain(state.Tiers)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.Float64Type, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type, "up_to": types.Int64Type}}}, "tiers", "raw response"); err != nil {
				return err
			} else {
				if typedTiers, ok := valueTiers.(types.List); ok {
					state.Tiers = typedTiers
				}
			}
		} else if !hasRaw {
			if responseValueTiers, ok := plainFromResponseField(obj, "Tiers"); ok {
				if valueTiers, err := flattenPlainValue(
					preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(responseValueTiers, attrValueToPlain(state.Tiers)), attrValueToPlain(state.Tiers)), attrValueToPlain(state.Tiers)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"flat_amount": types.Int64Type, "flat_amount_decimal": types.Float64Type, "unit_amount": types.Int64Type, "unit_amount_decimal": types.Float64Type, "up_to": types.Int64Type}}},
					"tiers",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedTiers, ok := valueTiers.(types.List); ok {
						state.Tiers = typedTiers
					}
				}
			}
		}
	}
	{
		if rawValueTiersMode, rawOk := plainValueAtPath(raw, "tiers_mode"); rawOk {
			if valueTiersMode, err := flattenPlainValue(rawValueTiersMode, types.StringType, "tiers_mode", "raw response"); err != nil {
				return err
			} else {
				if typedTiersMode, ok := valueTiersMode.(types.String); ok {
					state.TiersMode = typedTiersMode
				}
			}
		} else if !hasRaw {
			if responseValueTiersMode, ok := plainFromResponseField(obj, "TiersMode"); ok {
				if valueTiersMode, err := flattenPlainValue(responseValueTiersMode, types.StringType, "tiers_mode", "response struct"); err != nil {
					return err
				} else {
					if typedTiersMode, ok := valueTiersMode.(types.String); ok {
						state.TiersMode = typedTiersMode
					}
				}
			}
		}
	}
	{
		assignedTransformUsage := false
		hadRawTransformUsage := false
		if rawValueTransformUsage, rawOk := plainValueAtPath(raw, "transform_usage"); rawOk {
			hadRawTransformUsage = true
			if rawValueTransformUsage != nil {
				sourceTransformUsage := applyConfiguredKeyedListShapes(rawValueTransformUsage, attrValueToPlain(state.TransformUsage))
				if valueTransformUsage, err := flattenPlainValue(sourceTransformUsage, types.ObjectType{AttrTypes: map[string]attr.Type{"divide_by": types.Int64Type, "round": types.StringType}}, "transform_usage", "raw response"); err != nil {
					return err
				} else {
					if typedTransformUsage, ok := valueTransformUsage.(types.Object); ok {
						state.TransformUsage = typedTransformUsage
						assignedTransformUsage = true
					}
				}
			}
		}
		if !assignedTransformUsage {
			if !hasRaw {
				if responseValueTransformUsage, ok := plainFromResponseField(obj, "TransformUsage"); ok {
					sourceTransformUsage := applyConfiguredKeyedListShapes(responseValueTransformUsage, attrValueToPlain(state.TransformUsage))
					if valueTransformUsage, err := flattenPlainValue(
						sourceTransformUsage,
						types.ObjectType{AttrTypes: map[string]attr.Type{"divide_by": types.Int64Type, "round": types.StringType}},
						"transform_usage",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTransformUsage, ok := valueTransformUsage.(types.Object); ok {
							state.TransformUsage = typedTransformUsage
							assignedTransformUsage = true
						}
					}
				}
			}
		}
		if !assignedTransformUsage && hadRawTransformUsage {
			if nullTransformUsage, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"divide_by": types.Int64Type, "round": types.StringType}}); ok {
				if typedTransformUsage, ok := nullTransformUsage.(types.Object); ok {
					state.TransformUsage = typedTransformUsage
				}
			}
		}
	}
	{
		if rawValueTrialPeriodDays, rawOk := plainValueAtPath(raw, "trial_period_days"); rawOk {
			if valueTrialPeriodDays, err := flattenPlainValue(rawValueTrialPeriodDays, types.Int64Type, "trial_period_days", "raw response"); err != nil {
				return err
			} else {
				if typedTrialPeriodDays, ok := valueTrialPeriodDays.(types.Int64); ok {
					state.TrialPeriodDays = typedTrialPeriodDays
				}
			}
		} else if !hasRaw {
			if responseValueTrialPeriodDays, ok := plainFromResponseField(obj, "TrialPeriodDays"); ok {
				if valueTrialPeriodDays, err := flattenPlainValue(responseValueTrialPeriodDays, types.Int64Type, "trial_period_days", "response struct"); err != nil {
					return err
				} else {
					if typedTrialPeriodDays, ok := valueTrialPeriodDays.(types.Int64); ok {
						state.TrialPeriodDays = typedTrialPeriodDays
					}
				}
			}
		}
	}
	{
		if rawValueUsageType, rawOk := plainValueAtPath(raw, "usage_type"); rawOk {
			if valueUsageType, err := flattenPlainValue(rawValueUsageType, types.StringType, "usage_type", "raw response"); err != nil {
				return err
			} else {
				if typedUsageType, ok := valueUsageType.(types.String); ok {
					state.UsageType = typedUsageType
				}
			}
		} else if !hasRaw {
			if responseValueUsageType, ok := plainFromResponseField(obj, "UsageType"); ok {
				if valueUsageType, err := flattenPlainValue(responseValueUsageType, types.StringType, "usage_type", "response struct"); err != nil {
					return err
				} else {
					if typedUsageType, ok := valueUsageType.(types.String); ok {
						state.UsageType = typedUsageType
					}
				}
			}
		}
	}
	return nil
}
