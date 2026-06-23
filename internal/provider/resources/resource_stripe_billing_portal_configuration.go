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
	"reflect"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &BillingPortalConfigurationResource{}

var _ resource.ResourceWithConfigure = &BillingPortalConfigurationResource{}

var _ resource.ResourceWithImportState = &BillingPortalConfigurationResource{}

func NewBillingPortalConfigurationResource() resource.Resource {
	return &BillingPortalConfigurationResource{}
}

type BillingPortalConfigurationResource struct {
	client *stripe.Client
}

type BillingPortalConfigurationResourceModel struct {
	Object           types.String `tfsdk:"object"`
	Active           types.Bool   `tfsdk:"active"`
	Application      types.String `tfsdk:"application"`
	BusinessProfile  types.Object `tfsdk:"business_profile"`
	Created          types.Int64  `tfsdk:"created"`
	DefaultReturnURL types.String `tfsdk:"default_return_url"`
	Features         types.Object `tfsdk:"features"`
	ID               types.String `tfsdk:"id"`
	IsDefault        types.Bool   `tfsdk:"is_default"`
	Livemode         types.Bool   `tfsdk:"livemode"`
	LoginPage        types.Object `tfsdk:"login_page"`
	Metadata         types.Map    `tfsdk:"metadata"`
	Name             types.String `tfsdk:"name"`
}

func (r *BillingPortalConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *BillingPortalConfigurationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_billing_portal_configuration"
}

func (r *BillingPortalConfigurationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A portal configuration describes the functionality and behavior you embed in a portal session. Related guide: [Configure the customer portal](/customer-management/configure-portal).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("billing_portal.configuration")},
			},
			"active": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether the configuration is active and can be used to create portal sessions.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"application": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the Connect Application that created the configuration.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"business_profile": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"headline": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The messaging shown to customers in the portal.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"privacy_policy_url": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A link to the business’s publicly available privacy policy.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"terms_of_service_url": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A link to the business’s publicly available terms of service.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"default_return_url": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The default URL to redirect customers to when they click on the portal's link to return to your website. This can be [overriden](https://docs.stripe.com/api/customer_portal/sessions/create#create_portal_session-return_url) when creating the session.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"features": schema.SingleNestedAttribute{
				Required: true,

				Attributes: map[string]schema.Attribute{
					"customer_update": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"allowed_updates": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The types of customer updates that are supported. When empty, customers are not updateable.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"enabled": schema.BoolAttribute{
								Required:    true,
								Description: "Whether the feature is enabled.",
							},
						},
					},
					"invoice_history": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Required:    true,
								Description: "Whether the feature is enabled.",
							},
						},
					},
					"payment_method_update": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Required:    true,
								Description: "Whether the feature is enabled.",
							},
							"payment_method_configuration": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The [Payment Method Configuration](/api/payment_method_configurations) to use for this portal session. When specified, customers will be able to update their payment method to one of the options specified by the payment method configuration. If not set, the default payment method configuration is used.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"subscription_cancel": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"cancellation_reason": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Required:    true,
										Description: "Whether the feature is enabled.",
									},
									"options": schema.ListAttribute{
										Required:    true,
										Description: "Which cancellation reasons will be given as options to the customer.",
										ElementType: types.StringType,
									},
								},
							},
							"enabled": schema.BoolAttribute{
								Required:    true,
								Description: "Whether the feature is enabled.",
							},
							"mode": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Whether to cancel subscriptions immediately or at the end of the billing period.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("at_period_end", "immediately")},
							},
							"proration_behavior": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Whether to create prorations when canceling subscriptions. Possible values are `none` and `create_prorations`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("always_invoice", "create_prorations", "none")},
							},
						},
					},
					"subscription_update": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"billing_cycle_anchor": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Determines the value to use for the billing cycle anchor on subscription updates. Valid values are `now` or `unchanged`, and the default value is `unchanged`. Setting the value to `now` resets the subscription's billing cycle anchor to the current time (in UTC). For more information, see the billing cycle [documentation](https://docs.stripe.com/billing/subscriptions/billing-cycle).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("now", "unchanged")},
							},
							"default_allowed_updates": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The types of subscription updates that are supported for items listed in the `products` attribute. When empty, subscriptions are not updateable.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"enabled": schema.BoolAttribute{
								Required:    true,
								Description: "Whether the feature is enabled.",
							},
							"products": schema.ListNestedAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The list of up to 10 products that support subscription updates.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"adjustable_quantity": schema.SingleNestedAttribute{
											Optional: true,
											Computed: true,

											PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
											Attributes: map[string]schema.Attribute{
												"enabled": schema.BoolAttribute{
													Required:    true,
													Description: "If true, the quantity can be adjusted to any non-negative integer.",
												},
												"maximum": schema.Int64Attribute{
													Optional:      true,
													Computed:      true,
													Description:   "The maximum quantity that can be set for the product.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
												},
												"minimum": schema.Int64Attribute{
													Optional:      true,
													Computed:      true,
													Description:   "The minimum quantity that can be set for the product.",
													PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
												},
											},
										},
										"prices": schema.ListAttribute{
											Required:    true,
											Description: "The list of price IDs which, when subscribed to, a subscription can be updated.",
											ElementType: types.StringType,
										},
										"product": schema.StringAttribute{
											Required:    true,
											Description: "The product ID.",
										},
									},
								},
							},
							"proration_behavior": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Determines how to handle prorations resulting from subscription updates. Valid values are `none`, `create_prorations`, and `always_invoice`. Defaults to a value of `none` if you don't set it during creation.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("always_invoice", "create_prorations", "none")},
							},
							"schedule_at_period_end": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"conditions": schema.ListNestedAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "List of conditions. When any condition is true, an update will be scheduled at the end of the current period.",
										PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													Required:    true,
													Description: "The type of condition.",
													Validators:  []validator.String{stringvalidator.OneOf("decreasing_item_amount", "shortening_interval")},
												},
											},
										},
									},
								},
							},
							"trial_update_behavior": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Determines how handle updates to trialing subscriptions. Valid values are `end_trial` and `continue_trial`. Defaults to a value of `end_trial` if you don't set it during creation.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("continue_trial", "end_trial")},
							},
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"is_default": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether the configuration is the default. If `true`, this configuration can be managed in the Dashboard and portal sessions will use this configuration unless it is overriden when creating the session.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"login_page": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Required:    true,
						Description: "If `true`, a shareable `url` will be generated that will take your customers to a hosted login page for the customer portal.\n\nIf `false`, the previously generated `url`, if any, will be deactivated.",
					},
					"url": schema.StringAttribute{
						Computed:      true,
						Description:   "A shareable URL to the hosted portal login page. Your customers will be able to log in with their [email](https://docs.stripe.com/api/customers/object#customer_object-email) and receive a link to their customer portal.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
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
				Description:   "The name of the configuration.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func (r *BillingPortalConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan BillingPortalConfigurationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandBillingPortalConfigurationCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building BillingPortalConfiguration create params", err.Error())
		return
	}

	obj, err := r.client.V1BillingPortalConfigurations.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating BillingPortalConfiguration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingPortalConfigurations.B, r.client.V1BillingPortalConfigurations.Key, stripe.FormatURLPath("/v1/billing_portal/configurations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingPortalConfiguration create raw response", err.Error())
		return
	}

	var createdState BillingPortalConfigurationResourceModel
	if err := flattenBillingPortalConfiguration(obj, &createdState); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingPortalConfiguration create response", err.Error())
		return
	}
	normalizeUnknownValues(&createdState)

	diffPlan := plan
	diffCreatedState := createdState

	postCreateParams, err := expandBillingPortalConfigurationPostCreateUpdate(diffPlan, diffCreatedState)
	if err != nil {
		resp.Diagnostics.AddError("Error building BillingPortalConfiguration post-create update params", err.Error())
		return
	}

	if paramsHaveValues(postCreateParams) {
		if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
			!createdState.Metadata.IsNull() && !createdState.Metadata.IsUnknown() {
			if !assignMetadataDiffToNamedField(postCreateParams, "Metadata", plan.Metadata, createdState.Metadata) {
				resp.Diagnostics.AddError("Error building BillingPortalConfiguration update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", postCreateParams))
				return
			}
		}
		obj, err = r.client.V1BillingPortalConfigurations.Update(ctx, createdState.ID.ValueString(), postCreateParams)
		if err != nil {
			resp.Diagnostics.AddError("Error finalizing BillingPortalConfiguration after create", err.Error())
			return
		}
		if err := ensureRawResponse(obj, r.client.V1BillingPortalConfigurations.B, r.client.V1BillingPortalConfigurations.Key, stripe.FormatURLPath("/v1/billing_portal/configurations/%s", obj.ID), nil); err != nil {
			resp.Diagnostics.AddError("Error hydrating BillingPortalConfiguration post-create update raw response", err.Error())
			return
		}
	}

	if err := flattenBillingPortalConfiguration(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingPortalConfiguration create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingPortalConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState BillingPortalConfigurationResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state BillingPortalConfigurationResourceModel
	state = priorState

	obj, err := r.client.V1BillingPortalConfigurations.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading BillingPortalConfiguration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingPortalConfigurations.B, r.client.V1BillingPortalConfigurations.Key, stripe.FormatURLPath("/v1/billing_portal/configurations/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingPortalConfiguration raw response", err.Error())
		return
	}

	if err := flattenBillingPortalConfiguration(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingPortalConfiguration read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingPortalConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan BillingPortalConfigurationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state BillingPortalConfigurationResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandBillingPortalConfigurationUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building BillingPortalConfiguration update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building BillingPortalConfiguration update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1BillingPortalConfigurations.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating BillingPortalConfiguration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1BillingPortalConfigurations.B, r.client.V1BillingPortalConfigurations.Key, stripe.FormatURLPath("/v1/billing_portal/configurations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating BillingPortalConfiguration update raw response", err.Error())
		return
	}

	if err := flattenBillingPortalConfiguration(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening BillingPortalConfiguration update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *BillingPortalConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state BillingPortalConfigurationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !state.Active.IsNull() && !state.Active.IsUnknown() && !state.Active.ValueBool() {
		return
	}

	params := &stripe.BillingPortalConfigurationUpdateParams{}
	activeField := reflect.ValueOf(params).Elem().FieldByName("Active")
	if activeField.IsValid() && activeField.CanSet() {
		if activeField.Kind() == reflect.Pointer && activeField.Type().Elem().Kind() == reflect.Bool {
			activeField.Set(reflect.ValueOf(stripe.Bool(false)))
		} else if activeField.Kind() == reflect.Bool {
			activeField.SetBool(false)
		}
	}

	_, err := r.client.V1BillingPortalConfigurations.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deactivating BillingPortalConfiguration", err.Error())
		return
	}
}

func (r *BillingPortalConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandBillingPortalConfigurationCreate(plan BillingPortalConfigurationResourceModel) (*stripe.BillingPortalConfigurationCreateParams, error) {
	params := &stripe.BillingPortalConfigurationCreateParams{}

	if !plan.BusinessProfile.IsNull() && !plan.BusinessProfile.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BusinessProfile", plan.BusinessProfile) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "business_profile", params)
		}
	}
	if !plan.DefaultReturnURL.IsNull() && !plan.DefaultReturnURL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DefaultReturnURL", "DefaultReturnURL", plan.DefaultReturnURL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "default_return_url", params)
		}
	}
	if !plan.Features.IsNull() && !plan.Features.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Features", plan.Features) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "features", params)
		}
	}
	if !plan.LoginPage.IsNull() && !plan.LoginPage.IsUnknown() {
		if !assignAttrValueToNamedField(params, "LoginPage", plan.LoginPage) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "login_page", params)
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

func expandBillingPortalConfigurationUpdate(plan BillingPortalConfigurationResourceModel, state BillingPortalConfigurationResourceModel) (*stripe.BillingPortalConfigurationUpdateParams, error) {
	params := &stripe.BillingPortalConfigurationUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}
	if !plan.BusinessProfile.Equal(state.BusinessProfile) && !plan.BusinessProfile.IsNull() && !plan.BusinessProfile.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BusinessProfile", plan.BusinessProfile) {
			if !plan.BusinessProfile.Equal(state.BusinessProfile) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "business_profile", params)
			}
		}
	}
	if !plan.DefaultReturnURL.Equal(state.DefaultReturnURL) && !plan.DefaultReturnURL.IsNull() && !plan.DefaultReturnURL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DefaultReturnURL", "DefaultReturnURL", plan.DefaultReturnURL.ValueString()) {
			if !plan.DefaultReturnURL.Equal(state.DefaultReturnURL) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "default_return_url", params)
			}
		}
	}
	if !plan.Features.Equal(state.Features) && !plan.Features.IsNull() && !plan.Features.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Features", plan.Features) {
			if !plan.Features.Equal(state.Features) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "features", params)
			}
		}
	}
	if !plan.LoginPage.Equal(state.LoginPage) && !plan.LoginPage.IsNull() && !plan.LoginPage.IsUnknown() {
		if !assignAttrValueToNamedField(params, "LoginPage", plan.LoginPage) {
			if !plan.LoginPage.Equal(state.LoginPage) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "login_page", params)
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

func expandBillingPortalConfigurationPostCreateUpdate(plan BillingPortalConfigurationResourceModel, state BillingPortalConfigurationResourceModel) (*stripe.BillingPortalConfigurationUpdateParams, error) {
	params := &stripe.BillingPortalConfigurationUpdateParams{}

	if !plan.Active.Equal(state.Active) && !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		params.Active = stripe.Bool(plan.Active.ValueBool())
	}

	return params, nil
}

func flattenBillingPortalConfiguration(obj *stripe.BillingPortalConfiguration, state *BillingPortalConfigurationResourceModel) error {
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
		if true {
			if rawValueApplication, rawOk := plainValueAtPath(raw, "application"); rawOk {
				if typedApplication, ok := plainToStringIDValue(rawValueApplication); ok {
					state.Application = typedApplication
				}
			} else if !hasRaw {
				if responseValueApplication, ok := plainFromResponseField(obj, "Application"); ok {
					if typedApplication, ok := plainToStringIDValue(responseValueApplication); ok {
						state.Application = typedApplication
					}
				}
			}
		}
	}
	{
		assignedBusinessProfile := false
		hadRawBusinessProfile := false
		if rawValueBusinessProfile, rawOk := plainValueAtPath(raw, "business_profile"); rawOk {
			hadRawBusinessProfile = true
			if rawValueBusinessProfile != nil {
				sourceBusinessProfile := applyConfiguredKeyedListShapes(rawValueBusinessProfile, attrValueToPlain(state.BusinessProfile))
				if valueBusinessProfile, err := flattenPlainValue(sourceBusinessProfile, types.ObjectType{AttrTypes: map[string]attr.Type{"headline": types.StringType, "privacy_policy_url": types.StringType, "terms_of_service_url": types.StringType}}, "business_profile", "raw response"); err != nil {
					return err
				} else {
					if typedBusinessProfile, ok := valueBusinessProfile.(types.Object); ok {
						state.BusinessProfile = typedBusinessProfile
						assignedBusinessProfile = true
					}
				}
			}
		}
		if !assignedBusinessProfile {
			if !hasRaw {
				if responseValueBusinessProfile, ok := plainFromResponseField(obj, "BusinessProfile"); ok {
					sourceBusinessProfile := applyConfiguredKeyedListShapes(responseValueBusinessProfile, attrValueToPlain(state.BusinessProfile))
					if valueBusinessProfile, err := flattenPlainValue(
						sourceBusinessProfile,
						types.ObjectType{AttrTypes: map[string]attr.Type{"headline": types.StringType, "privacy_policy_url": types.StringType, "terms_of_service_url": types.StringType}},
						"business_profile",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBusinessProfile, ok := valueBusinessProfile.(types.Object); ok {
							state.BusinessProfile = typedBusinessProfile
							assignedBusinessProfile = true
						}
					}
				}
			}
		}
		if !assignedBusinessProfile && hadRawBusinessProfile {
			if nullBusinessProfile, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"headline": types.StringType, "privacy_policy_url": types.StringType, "terms_of_service_url": types.StringType}}); ok {
				if typedBusinessProfile, ok := nullBusinessProfile.(types.Object); ok {
					state.BusinessProfile = typedBusinessProfile
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
		if rawValueDefaultReturnURL, rawOk := plainValueAtPath(raw, "default_return_url"); rawOk {
			if valueDefaultReturnURL, err := flattenPlainValue(rawValueDefaultReturnURL, types.StringType, "default_return_url", "raw response"); err != nil {
				return err
			} else {
				if typedDefaultReturnURL, ok := valueDefaultReturnURL.(types.String); ok {
					state.DefaultReturnURL = typedDefaultReturnURL
				}
			}
		} else if !hasRaw {
			if responseValueDefaultReturnURL, ok := plainFromResponseField(obj, "DefaultReturnURL"); ok {
				if valueDefaultReturnURL, err := flattenPlainValue(responseValueDefaultReturnURL, types.StringType, "default_return_url", "response struct"); err != nil {
					return err
				} else {
					if typedDefaultReturnURL, ok := valueDefaultReturnURL.(types.String); ok {
						state.DefaultReturnURL = typedDefaultReturnURL
					}
				}
			}
		}
	}
	{
		assignedFeatures := false
		hadRawFeatures := false
		if rawValueFeatures, rawOk := plainValueAtPath(raw, "features"); rawOk {
			hadRawFeatures = true
			if rawValueFeatures != nil {
				sourceFeatures := applyConfiguredKeyedListShapes(rawValueFeatures, attrValueToPlain(state.Features))
				if valueFeatures, err := flattenPlainValue(sourceFeatures, types.ObjectType{AttrTypes: map[string]attr.Type{"customer_update": types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_updates": types.ListType{ElemType: types.StringType}, "enabled": types.BoolType}}, "invoice_history": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "payment_method_update": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "payment_method_configuration": types.StringType}}, "subscription_cancel": types.ObjectType{AttrTypes: map[string]attr.Type{"cancellation_reason": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "options": types.ListType{ElemType: types.StringType}}}, "enabled": types.BoolType, "mode": types.StringType, "proration_behavior": types.StringType}}, "subscription_update": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_cycle_anchor": types.StringType, "default_allowed_updates": types.ListType{ElemType: types.StringType}, "enabled": types.BoolType, "products": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "prices": types.ListType{ElemType: types.StringType}, "product": types.StringType}}}, "proration_behavior": types.StringType, "schedule_at_period_end": types.ObjectType{AttrTypes: map[string]attr.Type{"conditions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}}, "trial_update_behavior": types.StringType}}}}, "features", "raw response"); err != nil {
					return err
				} else {
					if typedFeatures, ok := valueFeatures.(types.Object); ok {
						state.Features = typedFeatures
						assignedFeatures = true
					}
				}
			}
		}
		if !assignedFeatures {
			if !hasRaw {
				if responseValueFeatures, ok := plainFromResponseField(obj, "Features"); ok {
					sourceFeatures := applyConfiguredKeyedListShapes(responseValueFeatures, attrValueToPlain(state.Features))
					if valueFeatures, err := flattenPlainValue(
						sourceFeatures,
						types.ObjectType{AttrTypes: map[string]attr.Type{"customer_update": types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_updates": types.ListType{ElemType: types.StringType}, "enabled": types.BoolType}}, "invoice_history": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "payment_method_update": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "payment_method_configuration": types.StringType}}, "subscription_cancel": types.ObjectType{AttrTypes: map[string]attr.Type{"cancellation_reason": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "options": types.ListType{ElemType: types.StringType}}}, "enabled": types.BoolType, "mode": types.StringType, "proration_behavior": types.StringType}}, "subscription_update": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_cycle_anchor": types.StringType, "default_allowed_updates": types.ListType{ElemType: types.StringType}, "enabled": types.BoolType, "products": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "prices": types.ListType{ElemType: types.StringType}, "product": types.StringType}}}, "proration_behavior": types.StringType, "schedule_at_period_end": types.ObjectType{AttrTypes: map[string]attr.Type{"conditions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}}, "trial_update_behavior": types.StringType}}}},
						"features",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFeatures, ok := valueFeatures.(types.Object); ok {
							state.Features = typedFeatures
							assignedFeatures = true
						}
					}
				}
			}
		}
		if !assignedFeatures && hadRawFeatures {
			if nullFeatures, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"customer_update": types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_updates": types.ListType{ElemType: types.StringType}, "enabled": types.BoolType}}, "invoice_history": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "payment_method_update": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "payment_method_configuration": types.StringType}}, "subscription_cancel": types.ObjectType{AttrTypes: map[string]attr.Type{"cancellation_reason": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "options": types.ListType{ElemType: types.StringType}}}, "enabled": types.BoolType, "mode": types.StringType, "proration_behavior": types.StringType}}, "subscription_update": types.ObjectType{AttrTypes: map[string]attr.Type{"billing_cycle_anchor": types.StringType, "default_allowed_updates": types.ListType{ElemType: types.StringType}, "enabled": types.BoolType, "products": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"adjustable_quantity": types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "maximum": types.Int64Type, "minimum": types.Int64Type}}, "prices": types.ListType{ElemType: types.StringType}, "product": types.StringType}}}, "proration_behavior": types.StringType, "schedule_at_period_end": types.ObjectType{AttrTypes: map[string]attr.Type{"conditions": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}}}}, "trial_update_behavior": types.StringType}}}}); ok {
				if typedFeatures, ok := nullFeatures.(types.Object); ok {
					state.Features = typedFeatures
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
		if rawValueIsDefault, rawOk := plainValueAtPath(raw, "is_default"); rawOk {
			if valueIsDefault, err := flattenPlainValue(rawValueIsDefault, types.BoolType, "is_default", "raw response"); err != nil {
				return err
			} else {
				if typedIsDefault, ok := valueIsDefault.(types.Bool); ok {
					state.IsDefault = typedIsDefault
				}
			}
		} else if !hasRaw {
			if responseValueIsDefault, ok := plainFromResponseField(obj, "IsDefault"); ok {
				if valueIsDefault, err := flattenPlainValue(responseValueIsDefault, types.BoolType, "is_default", "response struct"); err != nil {
					return err
				} else {
					if typedIsDefault, ok := valueIsDefault.(types.Bool); ok {
						state.IsDefault = typedIsDefault
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
		assignedLoginPage := false
		if rawValueLoginPage, rawOk := plainValueAtPath(raw, "login_page"); rawOk {
			if rawValueLoginPage != nil {
				sourceLoginPage := mergeMissingPlainLeaves(applyConfiguredKeyedListShapes(rawValueLoginPage, attrValueToPlain(state.LoginPage)), attrValueToPlain(state.LoginPage))
				if valueLoginPage, err := flattenPlainValue(sourceLoginPage, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "url": types.StringType}}, "login_page", "raw response"); err != nil {
					return err
				} else {
					if typedLoginPage, ok := valueLoginPage.(types.Object); ok {
						state.LoginPage = typedLoginPage
						assignedLoginPage = true
					}
				}
			}
		}
		if !assignedLoginPage {
			if !hasRaw {
				if responseValueLoginPage, ok := plainFromResponseField(obj, "LoginPage"); ok {
					sourceLoginPage := mergeMissingPlainLeaves(applyConfiguredKeyedListShapes(responseValueLoginPage, attrValueToPlain(state.LoginPage)), attrValueToPlain(state.LoginPage))
					if valueLoginPage, err := flattenPlainValue(
						sourceLoginPage,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType, "url": types.StringType}},
						"login_page",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLoginPage, ok := valueLoginPage.(types.Object); ok {
							state.LoginPage = typedLoginPage
							assignedLoginPage = true
						}
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
