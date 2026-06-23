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

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &PaymentMethodDomainResource{}

var _ resource.ResourceWithConfigure = &PaymentMethodDomainResource{}

var _ resource.ResourceWithImportState = &PaymentMethodDomainResource{}

func NewPaymentMethodDomainResource() resource.Resource {
	return &PaymentMethodDomainResource{}
}

type PaymentMethodDomainResource struct {
	client *stripe.Client
}

type PaymentMethodDomainResourceModel struct {
	Object     types.String `tfsdk:"object"`
	AmazonPay  types.Object `tfsdk:"amazon_pay"`
	ApplePay   types.Object `tfsdk:"apple_pay"`
	Created    types.Int64  `tfsdk:"created"`
	DomainName types.String `tfsdk:"domain_name"`
	Enabled    types.Bool   `tfsdk:"enabled"`
	GooglePay  types.Object `tfsdk:"google_pay"`
	ID         types.String `tfsdk:"id"`
	Klarna     types.Object `tfsdk:"klarna"`
	Link       types.Object `tfsdk:"link"`
	Livemode   types.Bool   `tfsdk:"livemode"`
	Paypal     types.Object `tfsdk:"paypal"`
}

func (r *PaymentMethodDomainResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PaymentMethodDomainResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_payment_method_domain"
}

func (r *PaymentMethodDomainResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A payment method domain represents a web domain that you have registered with Stripe.\nStripe Elements use registered payment method domains to control where certain payment methods are shown.\n\nRelated guide: [Payment method domains](https://docs.stripe.com/payments/payment-methods/pmd-registration).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("payment_method_domain")},
			},
			"amazon_pay": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Indicates the status of a specific payment method on a payment method domain.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the payment method on the domain.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("active", "inactive")},
					},
					"status_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains additional details about the status of a payment method for a specific payment method domain.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"error_message": schema.StringAttribute{
								Computed:      true,
								Description:   "The error message associated with the status of the payment method on the domain.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"apple_pay": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Indicates the status of a specific payment method on a payment method domain.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the payment method on the domain.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("active", "inactive")},
					},
					"status_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains additional details about the status of a payment method for a specific payment method domain.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"error_message": schema.StringAttribute{
								Computed:      true,
								Description:   "The error message associated with the status of the payment method on the domain.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"domain_name": schema.StringAttribute{
				Required:      true,
				Description:   "The domain name that this payment method domain object represents.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"enabled": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether this payment method domain is enabled. If the domain is not enabled, payment methods that require a payment method domain will not appear in Elements.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"google_pay": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Indicates the status of a specific payment method on a payment method domain.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the payment method on the domain.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("active", "inactive")},
					},
					"status_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains additional details about the status of a payment method for a specific payment method domain.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"error_message": schema.StringAttribute{
								Computed:      true,
								Description:   "The error message associated with the status of the payment method on the domain.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"klarna": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Indicates the status of a specific payment method on a payment method domain.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the payment method on the domain.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("active", "inactive")},
					},
					"status_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains additional details about the status of a payment method for a specific payment method domain.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"error_message": schema.StringAttribute{
								Computed:      true,
								Description:   "The error message associated with the status of the payment method on the domain.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"link": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Indicates the status of a specific payment method on a payment method domain.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the payment method on the domain.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("active", "inactive")},
					},
					"status_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains additional details about the status of a payment method for a specific payment method domain.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"error_message": schema.StringAttribute{
								Computed:      true,
								Description:   "The error message associated with the status of the payment method on the domain.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"paypal": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Indicates the status of a specific payment method on a payment method domain.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The status of the payment method on the domain.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("active", "inactive")},
					},
					"status_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Contains additional details about the status of a payment method for a specific payment method domain.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"error_message": schema.StringAttribute{
								Computed:      true,
								Description:   "The error message associated with the status of the payment method on the domain.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
		},
	}
}

func (r *PaymentMethodDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PaymentMethodDomainResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandPaymentMethodDomainCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentMethodDomain create params", err.Error())
		return
	}

	obj, err := r.client.V1PaymentMethodDomains.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating PaymentMethodDomain", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentMethodDomains.B, r.client.V1PaymentMethodDomains.Key, stripe.FormatURLPath("/v1/payment_method_domains/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentMethodDomain create raw response", err.Error())
		return
	}

	if err := flattenPaymentMethodDomain(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentMethodDomain create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentMethodDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState PaymentMethodDomainResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PaymentMethodDomainResourceModel
	state = priorState

	obj, err := r.client.V1PaymentMethodDomains.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading PaymentMethodDomain", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentMethodDomains.B, r.client.V1PaymentMethodDomains.Key, stripe.FormatURLPath("/v1/payment_method_domains/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentMethodDomain raw response", err.Error())
		return
	}

	if err := flattenPaymentMethodDomain(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentMethodDomain read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentMethodDomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PaymentMethodDomainResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PaymentMethodDomainResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandPaymentMethodDomainUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building PaymentMethodDomain update params", err.Error())
		return
	}

	obj, err := r.client.V1PaymentMethodDomains.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating PaymentMethodDomain", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1PaymentMethodDomains.B, r.client.V1PaymentMethodDomains.Key, stripe.FormatURLPath("/v1/payment_method_domains/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating PaymentMethodDomain update raw response", err.Error())
		return
	}

	if err := flattenPaymentMethodDomain(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening PaymentMethodDomain update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PaymentMethodDomainResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *PaymentMethodDomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandPaymentMethodDomainCreate(plan PaymentMethodDomainResourceModel) (*stripe.PaymentMethodDomainCreateParams, error) {
	params := &stripe.PaymentMethodDomainCreateParams{}

	if !plan.DomainName.IsNull() && !plan.DomainName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DomainName", "DomainName", plan.DomainName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "domain_name", params)
		}
	}
	if !plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
		params.Enabled = stripe.Bool(plan.Enabled.ValueBool())
	}

	return params, nil
}

func expandPaymentMethodDomainUpdate(plan PaymentMethodDomainResourceModel, state PaymentMethodDomainResourceModel) (*stripe.PaymentMethodDomainUpdateParams, error) {
	params := &stripe.PaymentMethodDomainUpdateParams{}

	if !plan.Enabled.Equal(state.Enabled) && !plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
		params.Enabled = stripe.Bool(plan.Enabled.ValueBool())
	}

	return params, nil
}

func flattenPaymentMethodDomain(obj *stripe.PaymentMethodDomain, state *PaymentMethodDomainResourceModel) error {
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
		assignedAmazonPay := false
		hadRawAmazonPay := false
		if rawValueAmazonPay, rawOk := plainValueAtPath(raw, "amazon_pay"); rawOk {
			hadRawAmazonPay = true
			if rawValueAmazonPay != nil {
				sourceAmazonPay := applyConfiguredKeyedListShapes(rawValueAmazonPay, attrValueToPlain(state.AmazonPay))
				if valueAmazonPay, err := flattenPlainValue(sourceAmazonPay, types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}, "amazon_pay", "raw response"); err != nil {
					return err
				} else {
					if typedAmazonPay, ok := valueAmazonPay.(types.Object); ok {
						state.AmazonPay = typedAmazonPay
						assignedAmazonPay = true
					}
				}
			}
		}
		if !assignedAmazonPay {
			if !hasRaw {
				if responseValueAmazonPay, ok := plainFromResponseField(obj, "AmazonPay"); ok {
					sourceAmazonPay := applyConfiguredKeyedListShapes(responseValueAmazonPay, attrValueToPlain(state.AmazonPay))
					if valueAmazonPay, err := flattenPlainValue(
						sourceAmazonPay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}},
						"amazon_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAmazonPay, ok := valueAmazonPay.(types.Object); ok {
							state.AmazonPay = typedAmazonPay
							assignedAmazonPay = true
						}
					}
				}
			}
		}
		if !assignedAmazonPay && hadRawAmazonPay {
			if nullAmazonPay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}); ok {
				if typedAmazonPay, ok := nullAmazonPay.(types.Object); ok {
					state.AmazonPay = typedAmazonPay
				}
			}
		}
	}
	{
		assignedApplePay := false
		hadRawApplePay := false
		if rawValueApplePay, rawOk := plainValueAtPath(raw, "apple_pay"); rawOk {
			hadRawApplePay = true
			if rawValueApplePay != nil {
				sourceApplePay := applyConfiguredKeyedListShapes(rawValueApplePay, attrValueToPlain(state.ApplePay))
				if valueApplePay, err := flattenPlainValue(sourceApplePay, types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}, "apple_pay", "raw response"); err != nil {
					return err
				} else {
					if typedApplePay, ok := valueApplePay.(types.Object); ok {
						state.ApplePay = typedApplePay
						assignedApplePay = true
					}
				}
			}
		}
		if !assignedApplePay {
			if !hasRaw {
				if responseValueApplePay, ok := plainFromResponseField(obj, "ApplePay"); ok {
					sourceApplePay := applyConfiguredKeyedListShapes(responseValueApplePay, attrValueToPlain(state.ApplePay))
					if valueApplePay, err := flattenPlainValue(
						sourceApplePay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}},
						"apple_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedApplePay, ok := valueApplePay.(types.Object); ok {
							state.ApplePay = typedApplePay
							assignedApplePay = true
						}
					}
				}
			}
		}
		if !assignedApplePay && hadRawApplePay {
			if nullApplePay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}); ok {
				if typedApplePay, ok := nullApplePay.(types.Object); ok {
					state.ApplePay = typedApplePay
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
		if rawValueEnabled, rawOk := plainValueAtPath(raw, "enabled"); rawOk {
			if valueEnabled, err := flattenPlainValue(rawValueEnabled, types.BoolType, "enabled", "raw response"); err != nil {
				return err
			} else {
				if typedEnabled, ok := valueEnabled.(types.Bool); ok {
					state.Enabled = typedEnabled
				}
			}
		} else if !hasRaw {
			if responseValueEnabled, ok := plainFromResponseField(obj, "Enabled"); ok {
				if valueEnabled, err := flattenPlainValue(responseValueEnabled, types.BoolType, "enabled", "response struct"); err != nil {
					return err
				} else {
					if typedEnabled, ok := valueEnabled.(types.Bool); ok {
						state.Enabled = typedEnabled
					}
				}
			}
		}
	}
	{
		assignedGooglePay := false
		hadRawGooglePay := false
		if rawValueGooglePay, rawOk := plainValueAtPath(raw, "google_pay"); rawOk {
			hadRawGooglePay = true
			if rawValueGooglePay != nil {
				sourceGooglePay := applyConfiguredKeyedListShapes(rawValueGooglePay, attrValueToPlain(state.GooglePay))
				if valueGooglePay, err := flattenPlainValue(sourceGooglePay, types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}, "google_pay", "raw response"); err != nil {
					return err
				} else {
					if typedGooglePay, ok := valueGooglePay.(types.Object); ok {
						state.GooglePay = typedGooglePay
						assignedGooglePay = true
					}
				}
			}
		}
		if !assignedGooglePay {
			if !hasRaw {
				if responseValueGooglePay, ok := plainFromResponseField(obj, "GooglePay"); ok {
					sourceGooglePay := applyConfiguredKeyedListShapes(responseValueGooglePay, attrValueToPlain(state.GooglePay))
					if valueGooglePay, err := flattenPlainValue(
						sourceGooglePay,
						types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}},
						"google_pay",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedGooglePay, ok := valueGooglePay.(types.Object); ok {
							state.GooglePay = typedGooglePay
							assignedGooglePay = true
						}
					}
				}
			}
		}
		if !assignedGooglePay && hadRawGooglePay {
			if nullGooglePay, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}); ok {
				if typedGooglePay, ok := nullGooglePay.(types.Object); ok {
					state.GooglePay = typedGooglePay
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
		assignedKlarna := false
		hadRawKlarna := false
		if rawValueKlarna, rawOk := plainValueAtPath(raw, "klarna"); rawOk {
			hadRawKlarna = true
			if rawValueKlarna != nil {
				sourceKlarna := applyConfiguredKeyedListShapes(rawValueKlarna, attrValueToPlain(state.Klarna))
				if valueKlarna, err := flattenPlainValue(sourceKlarna, types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}, "klarna", "raw response"); err != nil {
					return err
				} else {
					if typedKlarna, ok := valueKlarna.(types.Object); ok {
						state.Klarna = typedKlarna
						assignedKlarna = true
					}
				}
			}
		}
		if !assignedKlarna {
			if !hasRaw {
				if responseValueKlarna, ok := plainFromResponseField(obj, "Klarna"); ok {
					sourceKlarna := applyConfiguredKeyedListShapes(responseValueKlarna, attrValueToPlain(state.Klarna))
					if valueKlarna, err := flattenPlainValue(
						sourceKlarna,
						types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}},
						"klarna",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedKlarna, ok := valueKlarna.(types.Object); ok {
							state.Klarna = typedKlarna
							assignedKlarna = true
						}
					}
				}
			}
		}
		if !assignedKlarna && hadRawKlarna {
			if nullKlarna, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}); ok {
				if typedKlarna, ok := nullKlarna.(types.Object); ok {
					state.Klarna = typedKlarna
				}
			}
		}
	}
	{
		assignedLink := false
		hadRawLink := false
		if rawValueLink, rawOk := plainValueAtPath(raw, "link"); rawOk {
			hadRawLink = true
			if rawValueLink != nil {
				sourceLink := applyConfiguredKeyedListShapes(rawValueLink, attrValueToPlain(state.Link))
				if valueLink, err := flattenPlainValue(sourceLink, types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}, "link", "raw response"); err != nil {
					return err
				} else {
					if typedLink, ok := valueLink.(types.Object); ok {
						state.Link = typedLink
						assignedLink = true
					}
				}
			}
		}
		if !assignedLink {
			if !hasRaw {
				if responseValueLink, ok := plainFromResponseField(obj, "Link"); ok {
					sourceLink := applyConfiguredKeyedListShapes(responseValueLink, attrValueToPlain(state.Link))
					if valueLink, err := flattenPlainValue(
						sourceLink,
						types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}},
						"link",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLink, ok := valueLink.(types.Object); ok {
							state.Link = typedLink
							assignedLink = true
						}
					}
				}
			}
		}
		if !assignedLink && hadRawLink {
			if nullLink, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}); ok {
				if typedLink, ok := nullLink.(types.Object); ok {
					state.Link = typedLink
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
		assignedPaypal := false
		hadRawPaypal := false
		if rawValuePaypal, rawOk := plainValueAtPath(raw, "paypal"); rawOk {
			hadRawPaypal = true
			if rawValuePaypal != nil {
				sourcePaypal := applyConfiguredKeyedListShapes(rawValuePaypal, attrValueToPlain(state.Paypal))
				if valuePaypal, err := flattenPlainValue(sourcePaypal, types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}, "paypal", "raw response"); err != nil {
					return err
				} else {
					if typedPaypal, ok := valuePaypal.(types.Object); ok {
						state.Paypal = typedPaypal
						assignedPaypal = true
					}
				}
			}
		}
		if !assignedPaypal {
			if !hasRaw {
				if responseValuePaypal, ok := plainFromResponseField(obj, "Paypal"); ok {
					sourcePaypal := applyConfiguredKeyedListShapes(responseValuePaypal, attrValueToPlain(state.Paypal))
					if valuePaypal, err := flattenPlainValue(
						sourcePaypal,
						types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}},
						"paypal",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedPaypal, ok := valuePaypal.(types.Object); ok {
							state.Paypal = typedPaypal
							assignedPaypal = true
						}
					}
				}
			}
		}
		if !assignedPaypal && hadRawPaypal {
			if nullPaypal, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "status_details": types.ObjectType{AttrTypes: map[string]attr.Type{"error_message": types.StringType}}}}); ok {
				if typedPaypal, ok := nullPaypal.(types.Object); ok {
					state.Paypal = typedPaypal
				}
			}
		}
	}
	return nil
}
