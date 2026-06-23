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

var _ resource.Resource = &SubscriptionItemResource{}

var _ resource.ResourceWithConfigure = &SubscriptionItemResource{}

var _ resource.ResourceWithImportState = &SubscriptionItemResource{}

func NewSubscriptionItemResource() resource.Resource {
	return &SubscriptionItemResource{}
}

type SubscriptionItemResource struct {
	client *stripe.Client
}

type SubscriptionItemResourceModel struct {
	Object             types.String `tfsdk:"object"`
	BilledUntil        types.Int64  `tfsdk:"billed_until"`
	BillingThresholds  types.Object `tfsdk:"billing_thresholds"`
	Created            types.Int64  `tfsdk:"created"`
	CurrentPeriodEnd   types.Int64  `tfsdk:"current_period_end"`
	CurrentPeriodStart types.Int64  `tfsdk:"current_period_start"`
	Discounts          types.List   `tfsdk:"discounts"`
	ID                 types.String `tfsdk:"id"`
	Metadata           types.Map    `tfsdk:"metadata"`
	Plan               types.String `tfsdk:"plan"`
	Price              types.String `tfsdk:"price"`
	Quantity           types.Int64  `tfsdk:"quantity"`
	Subscription       types.String `tfsdk:"subscription"`
	TaxRates           types.List   `tfsdk:"tax_rates"`
	PaymentBehavior    types.String `tfsdk:"payment_behavior"`
	PriceData          types.Object `tfsdk:"price_data"`
	ProrationBehavior  types.String `tfsdk:"proration_behavior"`
	ProrationDate      types.Int64  `tfsdk:"proration_date"`
}

func (r *SubscriptionItemResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SubscriptionItemResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscription_item"
}

func (r *SubscriptionItemResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Subscription items allow you to create customer subscriptions with more than\none plan, making it easy to represent complex billing relationships.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("subscription_item")},
			},
			"billed_until": schema.Int64Attribute{
				Computed:      true,
				Description:   "The time period the subscription item has been billed for.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"billing_thresholds": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"usage_gte": schema.Int64Attribute{
						Required:    true,
						Description: "Usage threshold that triggers the subscription to create an invoice",
					},
				},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"current_period_end": schema.Int64Attribute{
				Computed:      true,
				Description:   "The end time of this subscription item's current billing period.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"current_period_start": schema.Int64Attribute{
				Computed:      true,
				Description:   "The start time of this subscription item's current billing period.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"discounts": schema.ListNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The discounts applied to the subscription item. Subscription item discounts are applied before subscription discounts. Use `expand[]=discounts` to expand each discount.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"coupon": schema.StringAttribute{
							Optional:    true,
							Description: "ID of the coupon to create a new discount for.",
						},
						"discount": schema.StringAttribute{
							Optional:    true,
							Description: "ID of an existing discount on the object (or one of its ancestors) to reuse.",
						},
						"promotion_code": schema.StringAttribute{
							Optional:    true,
							Description: "ID of the promotion code to create a new discount for.",
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"plan": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "You can now model subscriptions more flexibly using the [Prices API](https://api.stripe.com#prices). It replaces the Plans API and is backwards compatible to simplify your migration.\n\nPlans define the base price, currency, and billing cycle for recurring purchases of products.\n[Products](https://api.stripe.com#products) help you track inventory or provisioning, and plans help you track pricing. Different physical goods or levels of service should be represented by products, and pricing options should be represented by plans. This approach lets you change prices without having to change your provisioning scheme.\n\nFor example, you might have a single \"gold\" product that has plans for $10/month, $100/year, €9/month, and €90/year.\n\nRelated guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription) and more about [products and prices](https://docs.stripe.com/products-prices/overview).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"price": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.\n[Products](https://api.stripe.com#products) help you track inventory or provisioning, and prices help you track payment terms. Different physical goods or levels of service should be represented by products, and pricing options should be represented by prices. This approach lets you change prices without having to change your provisioning scheme.\n\nFor example, you might have a single \"gold\" product that has prices for $10/month, $100/year, and €9 once.\n\nRelated guides: [Set up a subscription](https://docs.stripe.com/billing/subscriptions/set-up-subscription), [create an invoice](https://docs.stripe.com/billing/invoices/create), and more about [products and prices](https://docs.stripe.com/products-prices/overview).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"quantity": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "The [quantity](https://docs.stripe.com/subscriptions/quantities) of the plan to which the customer should be subscribed.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"subscription": schema.StringAttribute{
				Required:      true,
				Description:   "The `subscription` this `subscription_item` belongs to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"tax_rates": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The tax rates which apply to this `subscription_item`. When set, the `default_tax_rates` on the subscription do not apply to this `subscription_item`.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"payment_behavior": schema.StringAttribute{
				Optional:    true,
				Description: "Controls how Stripe handles payment when a subscription update requires payment and `collection_method=charge_automatically`.",
				WriteOnly:   true,
				Validators:  []validator.String{stringvalidator.OneOf("allow_incomplete", "default_incomplete", "error_if_incomplete", "pending_if_incomplete")},
			},
			"price_data": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Data used to generate a new [Price](https://docs.stripe.com/api/prices) object inline.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"currency": schema.StringAttribute{
						Required:    true,
						Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
						WriteOnly:   true,
					},
					"product": schema.StringAttribute{
						Required:    true,
						Description: "The ID of the [Product](https://docs.stripe.com/api/products) that this [Price](https://docs.stripe.com/api/prices) will belong to.",
						WriteOnly:   true,
					},
					"recurring": schema.SingleNestedAttribute{
						Required:    true,
						Description: "The recurring components of a price such as `interval` and `interval_count`.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"interval": schema.StringAttribute{
								Required:    true,
								Description: "Specifies billing frequency. Either `day`, `week`, `month` or `year`.",
								WriteOnly:   true,
							},
							"interval_count": schema.Int64Attribute{
								Optional:    true,
								Description: "The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).",
								WriteOnly:   true,
							},
						},
					},
					"tax_behavior": schema.StringAttribute{
						Optional:    true,
						Description: "Only required if a [default tax behavior](https://docs.stripe.com/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.",
						WriteOnly:   true,
					},
					"unit_amount": schema.Int64Attribute{
						Optional:    true,
						Description: "A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.",
						WriteOnly:   true,
					},
					"unit_amount_decimal": schema.Float64Attribute{
						Optional:    true,
						Description: "Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.",
						WriteOnly:   true,
					},
				},
			},
			"proration_behavior": schema.StringAttribute{
				Optional:    true,
				Description: "Determines how to handle [prorations](https://docs.stripe.com/billing/subscriptions/prorations) when the billing cycle changes (e.g., when switching plans, resetting `billing_cycle_anchor=now`, or starting a trial), or if an item's `quantity` changes. The default value is `create_prorations`.",
				WriteOnly:   true,
				Validators:  []validator.String{stringvalidator.OneOf("always_invoice", "create_prorations", "none")},
			},
			"proration_date": schema.Int64Attribute{
				Optional:    true,
				Description: "If set, the proration will be calculated as though the subscription was updated at the given time. This can be used to apply the same proration that was previewed with the [upcoming invoice](/api/invoices/create_preview) endpoint.",
				WriteOnly:   true,
			},
		},
	}
}

func (r *SubscriptionItemResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SubscriptionItemResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config SubscriptionItemResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PaymentBehavior"}, []string{"PriceData"}, []string{"PriceData", "currency"}, []string{"PriceData", "product"}, []string{"PriceData", "recurring"}, []string{"PriceData", "recurring", "interval"}, []string{"PriceData", "recurring", "interval_count"}, []string{"PriceData", "tax_behavior"}, []string{"PriceData", "unit_amount"}, []string{"PriceData", "unit_amount_decimal"}, []string{"ProrationBehavior"}, []string{"ProrationDate"}})

	params, err := expandSubscriptionItemCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building SubscriptionItem create params", err.Error())
		return
	}

	obj, err := r.client.V1SubscriptionItems.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating SubscriptionItem", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1SubscriptionItems.B, r.client.V1SubscriptionItems.Key, stripe.FormatURLPath("/v1/subscription_items/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating SubscriptionItem create raw response", err.Error())
		return
	}

	if err := flattenSubscriptionItem(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening SubscriptionItem create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"PaymentBehavior"}, []string{"PriceData"}, []string{"PriceData", "currency"}, []string{"PriceData", "product"}, []string{"PriceData", "recurring"}, []string{"PriceData", "recurring", "interval"}, []string{"PriceData", "recurring", "interval_count"}, []string{"PriceData", "tax_behavior"}, []string{"PriceData", "unit_amount"}, []string{"PriceData", "unit_amount_decimal"}, []string{"ProrationBehavior"}, []string{"ProrationDate"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SubscriptionItemResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState SubscriptionItemResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state SubscriptionItemResourceModel
	state = priorState

	obj, err := r.client.V1SubscriptionItems.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading SubscriptionItem", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1SubscriptionItems.B, r.client.V1SubscriptionItems.Key, stripe.FormatURLPath("/v1/subscription_items/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating SubscriptionItem raw response", err.Error())
		return
	}

	if err := flattenSubscriptionItem(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening SubscriptionItem read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}})
	clearWriteOnlyPaths(&state, [][]string{[]string{"PaymentBehavior"}, []string{"PriceData"}, []string{"PriceData", "currency"}, []string{"PriceData", "product"}, []string{"PriceData", "recurring"}, []string{"PriceData", "recurring", "interval"}, []string{"PriceData", "recurring", "interval_count"}, []string{"PriceData", "tax_behavior"}, []string{"PriceData", "unit_amount"}, []string{"PriceData", "unit_amount_decimal"}, []string{"ProrationBehavior"}, []string{"ProrationDate"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *SubscriptionItemResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SubscriptionItemResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config SubscriptionItemResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PaymentBehavior"}, []string{"PriceData"}, []string{"PriceData", "currency"}, []string{"PriceData", "product"}, []string{"PriceData", "recurring"}, []string{"PriceData", "recurring", "interval"}, []string{"PriceData", "recurring", "interval_count"}, []string{"PriceData", "tax_behavior"}, []string{"PriceData", "unit_amount"}, []string{"PriceData", "unit_amount_decimal"}, []string{"ProrationBehavior"}, []string{"ProrationDate"}})

	var state SubscriptionItemResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandSubscriptionItemUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building SubscriptionItem update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building SubscriptionItem update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1SubscriptionItems.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating SubscriptionItem", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1SubscriptionItems.B, r.client.V1SubscriptionItems.Key, stripe.FormatURLPath("/v1/subscription_items/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating SubscriptionItem update raw response", err.Error())
		return
	}

	if err := flattenSubscriptionItem(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening SubscriptionItem update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Discounts", "*", "coupon"}, []string{"Discounts", "*", "discount"}, []string{"Discounts", "*", "promotion_code"}})
	clearWriteOnlyPaths(&plan, [][]string{[]string{"PaymentBehavior"}, []string{"PriceData"}, []string{"PriceData", "currency"}, []string{"PriceData", "product"}, []string{"PriceData", "recurring"}, []string{"PriceData", "recurring", "interval"}, []string{"PriceData", "recurring", "interval_count"}, []string{"PriceData", "tax_behavior"}, []string{"PriceData", "unit_amount"}, []string{"PriceData", "unit_amount_decimal"}, []string{"ProrationBehavior"}, []string{"ProrationDate"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *SubscriptionItemResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SubscriptionItemResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1SubscriptionItems.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting SubscriptionItem", err.Error())
		return
	}
}

func (r *SubscriptionItemResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandSubscriptionItemCreate(plan SubscriptionItemResourceModel) (*stripe.SubscriptionItemCreateParams, error) {
	params := &stripe.SubscriptionItemCreateParams{}

	if !plan.BillingThresholds.IsNull() && !plan.BillingThresholds.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingThresholds", plan.BillingThresholds) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "billing_thresholds", params)
		}
	}
	if !plan.Discounts.IsNull() && !plan.Discounts.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Discounts", plan.Discounts) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "discounts", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Plan.IsNull() && !plan.Plan.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PlanID", "Plan", plan.Plan.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "plan", params)
		}
	}
	if !plan.Price.IsNull() && !plan.Price.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PriceID", "Price", plan.Price.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "price", params)
		}
	}
	if !plan.Quantity.IsNull() && !plan.Quantity.IsUnknown() {
		params.Quantity = stripe.Int64(plan.Quantity.ValueInt64())
	}
	if !plan.Subscription.IsNull() && !plan.Subscription.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Subscription", "Subscription", plan.Subscription.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "subscription", params)
		}
	}
	if !plan.TaxRates.IsNull() && !plan.TaxRates.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TaxRates", plan.TaxRates) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tax_rates", params)
		}
	}
	if !plan.PaymentBehavior.IsNull() && !plan.PaymentBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentBehavior", "PaymentBehavior", plan.PaymentBehavior.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payment_behavior", params)
		}
	}
	if !plan.PriceData.IsNull() && !plan.PriceData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PriceData", plan.PriceData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "price_data", params)
		}
	}
	if !plan.ProrationBehavior.IsNull() && !plan.ProrationBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ProrationBehavior", "ProrationBehavior", plan.ProrationBehavior.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "proration_behavior", params)
		}
	}
	if !plan.ProrationDate.IsNull() && !plan.ProrationDate.IsUnknown() {
		params.ProrationDate = stripe.Int64(plan.ProrationDate.ValueInt64())
	}

	return params, nil
}

func expandSubscriptionItemUpdate(plan SubscriptionItemResourceModel, state SubscriptionItemResourceModel) (*stripe.SubscriptionItemUpdateParams, error) {
	params := &stripe.SubscriptionItemUpdateParams{}

	if !plan.BillingThresholds.Equal(state.BillingThresholds) && !plan.BillingThresholds.IsNull() && !plan.BillingThresholds.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BillingThresholds", plan.BillingThresholds) {
			if !plan.BillingThresholds.Equal(state.BillingThresholds) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "billing_thresholds", params)
			}
		}
	}
	if !plan.Discounts.Equal(state.Discounts) && !plan.Discounts.IsNull() && !plan.Discounts.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Discounts", plan.Discounts) {
			if !plan.Discounts.Equal(state.Discounts) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "discounts", params)
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
	if !plan.Plan.Equal(state.Plan) && !plan.Plan.IsNull() && !plan.Plan.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PlanID", "Plan", plan.Plan.ValueString()) {
			if !plan.Plan.Equal(state.Plan) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "plan", params)
			}
		}
	}
	if !plan.Price.Equal(state.Price) && !plan.Price.IsNull() && !plan.Price.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PriceID", "Price", plan.Price.ValueString()) {
			if !plan.Price.Equal(state.Price) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "price", params)
			}
		}
	}
	if !plan.Quantity.Equal(state.Quantity) && !plan.Quantity.IsNull() && !plan.Quantity.IsUnknown() {
		params.Quantity = stripe.Int64(plan.Quantity.ValueInt64())
	}
	if !plan.TaxRates.Equal(state.TaxRates) && !plan.TaxRates.IsNull() && !plan.TaxRates.IsUnknown() {
		if !assignAttrValueToNamedField(params, "TaxRates", plan.TaxRates) {
			if !plan.TaxRates.Equal(state.TaxRates) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tax_rates", params)
			}
		}
	}
	if !plan.PaymentBehavior.Equal(state.PaymentBehavior) && !plan.PaymentBehavior.IsNull() && !plan.PaymentBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PaymentBehavior", "PaymentBehavior", plan.PaymentBehavior.ValueString()) {
			if !plan.PaymentBehavior.Equal(state.PaymentBehavior) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payment_behavior", params)
			}
		}
	}
	if !plan.PriceData.Equal(state.PriceData) && !plan.PriceData.IsNull() && !plan.PriceData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PriceData", plan.PriceData) {
			if !plan.PriceData.Equal(state.PriceData) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "price_data", params)
			}
		}
	}
	if !plan.ProrationBehavior.Equal(state.ProrationBehavior) && !plan.ProrationBehavior.IsNull() && !plan.ProrationBehavior.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ProrationBehavior", "ProrationBehavior", plan.ProrationBehavior.ValueString()) {
			if !plan.ProrationBehavior.Equal(state.ProrationBehavior) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "proration_behavior", params)
			}
		}
	}
	if !plan.ProrationDate.Equal(state.ProrationDate) && !plan.ProrationDate.IsNull() && !plan.ProrationDate.IsUnknown() {
		params.ProrationDate = stripe.Int64(plan.ProrationDate.ValueInt64())
	}

	return params, nil
}

func flattenSubscriptionItem(obj *stripe.SubscriptionItem, state *SubscriptionItemResourceModel) error {
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
		if rawValueBilledUntil, rawOk := plainValueAtPath(raw, "billed_until"); rawOk {
			if valueBilledUntil, err := flattenPlainValue(rawValueBilledUntil, types.Int64Type, "billed_until", "raw response"); err != nil {
				return err
			} else {
				if typedBilledUntil, ok := valueBilledUntil.(types.Int64); ok {
					state.BilledUntil = typedBilledUntil
				}
			}
		} else if !hasRaw {
			if responseValueBilledUntil, ok := plainFromResponseField(obj, "BilledUntil"); ok {
				if valueBilledUntil, err := flattenPlainValue(responseValueBilledUntil, types.Int64Type, "billed_until", "response struct"); err != nil {
					return err
				} else {
					if typedBilledUntil, ok := valueBilledUntil.(types.Int64); ok {
						state.BilledUntil = typedBilledUntil
					}
				}
			}
		}
	}
	{
		assignedBillingThresholds := false
		hadRawBillingThresholds := false
		if rawValueBillingThresholds, rawOk := plainValueAtPath(raw, "billing_thresholds"); rawOk {
			hadRawBillingThresholds = true
			if rawValueBillingThresholds != nil {
				sourceBillingThresholds := applyConfiguredKeyedListShapes(rawValueBillingThresholds, attrValueToPlain(state.BillingThresholds))
				if valueBillingThresholds, err := flattenPlainValue(sourceBillingThresholds, types.ObjectType{AttrTypes: map[string]attr.Type{"usage_gte": types.Int64Type}}, "billing_thresholds", "raw response"); err != nil {
					return err
				} else {
					if typedBillingThresholds, ok := valueBillingThresholds.(types.Object); ok {
						state.BillingThresholds = typedBillingThresholds
						assignedBillingThresholds = true
					}
				}
			}
		}
		if !assignedBillingThresholds {
			if !hasRaw {
				if responseValueBillingThresholds, ok := plainFromResponseField(obj, "BillingThresholds"); ok {
					sourceBillingThresholds := applyConfiguredKeyedListShapes(responseValueBillingThresholds, attrValueToPlain(state.BillingThresholds))
					if valueBillingThresholds, err := flattenPlainValue(
						sourceBillingThresholds,
						types.ObjectType{AttrTypes: map[string]attr.Type{"usage_gte": types.Int64Type}},
						"billing_thresholds",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBillingThresholds, ok := valueBillingThresholds.(types.Object); ok {
							state.BillingThresholds = typedBillingThresholds
							assignedBillingThresholds = true
						}
					}
				}
			}
		}
		if !assignedBillingThresholds && hadRawBillingThresholds {
			if nullBillingThresholds, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"usage_gte": types.Int64Type}}); ok {
				if typedBillingThresholds, ok := nullBillingThresholds.(types.Object); ok {
					state.BillingThresholds = typedBillingThresholds
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
		if rawValueCurrentPeriodEnd, rawOk := plainValueAtPath(raw, "current_period_end"); rawOk {
			if valueCurrentPeriodEnd, err := flattenPlainValue(rawValueCurrentPeriodEnd, types.Int64Type, "current_period_end", "raw response"); err != nil {
				return err
			} else {
				if typedCurrentPeriodEnd, ok := valueCurrentPeriodEnd.(types.Int64); ok {
					state.CurrentPeriodEnd = typedCurrentPeriodEnd
				}
			}
		} else if !hasRaw {
			if responseValueCurrentPeriodEnd, ok := plainFromResponseField(obj, "CurrentPeriodEnd"); ok {
				if valueCurrentPeriodEnd, err := flattenPlainValue(responseValueCurrentPeriodEnd, types.Int64Type, "current_period_end", "response struct"); err != nil {
					return err
				} else {
					if typedCurrentPeriodEnd, ok := valueCurrentPeriodEnd.(types.Int64); ok {
						state.CurrentPeriodEnd = typedCurrentPeriodEnd
					}
				}
			}
		}
	}
	{
		if rawValueCurrentPeriodStart, rawOk := plainValueAtPath(raw, "current_period_start"); rawOk {
			if valueCurrentPeriodStart, err := flattenPlainValue(rawValueCurrentPeriodStart, types.Int64Type, "current_period_start", "raw response"); err != nil {
				return err
			} else {
				if typedCurrentPeriodStart, ok := valueCurrentPeriodStart.(types.Int64); ok {
					state.CurrentPeriodStart = typedCurrentPeriodStart
				}
			}
		} else if !hasRaw {
			if responseValueCurrentPeriodStart, ok := plainFromResponseField(obj, "CurrentPeriodStart"); ok {
				if valueCurrentPeriodStart, err := flattenPlainValue(responseValueCurrentPeriodStart, types.Int64Type, "current_period_start", "response struct"); err != nil {
					return err
				} else {
					if typedCurrentPeriodStart, ok := valueCurrentPeriodStart.(types.Int64); ok {
						state.CurrentPeriodStart = typedCurrentPeriodStart
					}
				}
			}
		}
	}
	{
		if rawValueDiscounts, rawOk := plainValueAtPath(raw, "discounts"); rawOk {
			if valueDiscounts, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueDiscounts, attrValueToPlain(state.Discounts)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}}, "discounts", "raw response"); err != nil {
				return err
			} else {
				if typedDiscounts, ok := valueDiscounts.(types.List); ok {
					state.Discounts = typedDiscounts
				}
			}
		} else if !hasRaw {
			if responseValueDiscounts, ok := plainFromResponseField(obj, "Discounts"); ok {
				if valueDiscounts, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueDiscounts, attrValueToPlain(state.Discounts)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"coupon": types.StringType, "discount": types.StringType, "promotion_code": types.StringType}}},
					"discounts",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedDiscounts, ok := valueDiscounts.(types.List); ok {
						state.Discounts = typedDiscounts
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
		if true {
			if rawValuePlan, rawOk := plainValueAtPath(raw, "plan"); rawOk {
				if typedPlan, ok := plainToStringIDValue(rawValuePlan); ok {
					state.Plan = typedPlan
				}
			} else if !hasRaw {
				if responseValuePlan, ok := plainFromResponseField(obj, "Plan"); ok {
					if typedPlan, ok := plainToStringIDValue(responseValuePlan); ok {
						state.Plan = typedPlan
					}
				}
			}
		}
	}
	{
		if true {
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
	{
		if rawValueQuantity, rawOk := plainValueAtPath(raw, "quantity"); rawOk {
			if valueQuantity, err := flattenPlainValue(rawValueQuantity, types.Int64Type, "quantity", "raw response"); err != nil {
				return err
			} else {
				if typedQuantity, ok := valueQuantity.(types.Int64); ok {
					state.Quantity = typedQuantity
				}
			}
		} else if !hasRaw {
			if responseValueQuantity, ok := plainFromResponseField(obj, "Quantity"); ok {
				if valueQuantity, err := flattenPlainValue(responseValueQuantity, types.Int64Type, "quantity", "response struct"); err != nil {
					return err
				} else {
					if typedQuantity, ok := valueQuantity.(types.Int64); ok {
						state.Quantity = typedQuantity
					}
				}
			}
		}
	}
	{
		if rawValueSubscription, rawOk := plainValueAtPath(raw, "subscription"); rawOk {
			if valueSubscription, err := flattenPlainValue(rawValueSubscription, types.StringType, "subscription", "raw response"); err != nil {
				return err
			} else {
				if typedSubscription, ok := valueSubscription.(types.String); ok {
					state.Subscription = typedSubscription
				}
			}
		} else if !hasRaw {
			if responseValueSubscription, ok := plainFromResponseField(obj, "Subscription"); ok {
				if valueSubscription, err := flattenPlainValue(responseValueSubscription, types.StringType, "subscription", "response struct"); err != nil {
					return err
				} else {
					if typedSubscription, ok := valueSubscription.(types.String); ok {
						state.Subscription = typedSubscription
					}
				}
			}
		}
	}
	{
		if rawValueTaxRates, rawOk := plainValueAtPath(raw, "tax_rates"); rawOk {
			if valueTaxRates, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueTaxRates, attrValueToPlain(state.TaxRates)), types.ListType{ElemType: types.StringType}, "tax_rates", "raw response"); err != nil {
				return err
			} else {
				if typedTaxRates, ok := valueTaxRates.(types.List); ok {
					state.TaxRates = typedTaxRates
				}
			}
		} else if !hasRaw {
			if responseValueTaxRates, ok := plainFromResponseField(obj, "TaxRates"); ok {
				if valueTaxRates, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueTaxRates, attrValueToPlain(state.TaxRates)),
					types.ListType{ElemType: types.StringType},
					"tax_rates",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedTaxRates, ok := valueTaxRates.(types.List); ok {
						state.TaxRates = typedTaxRates
					}
				}
			}
		}
	}
	return nil
}
