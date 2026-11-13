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

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &V2MoneyManagementOutboundSetupIntentResource{}

var _ resource.ResourceWithConfigure = &V2MoneyManagementOutboundSetupIntentResource{}

var _ resource.ResourceWithImportState = &V2MoneyManagementOutboundSetupIntentResource{}

func NewV2MoneyManagementOutboundSetupIntentResource() resource.Resource {
	return &V2MoneyManagementOutboundSetupIntentResource{}
}

type V2MoneyManagementOutboundSetupIntentResource struct {
	client *stripe.Client
}

type V2MoneyManagementOutboundSetupIntentResourceModel struct {
	Object           types.String `tfsdk:"object"`
	Created          types.String `tfsdk:"created"`
	ID               types.String `tfsdk:"id"`
	Livemode         types.Bool   `tfsdk:"livemode"`
	PayoutMethod     types.String `tfsdk:"payout_method"`
	Status           types.String `tfsdk:"status"`
	UsageIntent      types.String `tfsdk:"usage_intent"`
	PayoutMethodData types.Object `tfsdk:"payout_method_data"`
}

func (r *V2MoneyManagementOutboundSetupIntentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *V2MoneyManagementOutboundSetupIntentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v2_money_management_outbound_setup_intent"
}

func (r *V2MoneyManagementOutboundSetupIntentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Use the OutboundSetupIntent API to create and setup usable payout methods.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value of the object field.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("v2.money_management.outbound_setup_intent")},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				Description:   "Created timestamp.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the outbound setup intent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"payout_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Use the PayoutMethods API to list and interact with PayoutMethod objects.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Closed Enum. Status of the outbound setup intent.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("canceled", "requires_action", "requires_payout_method", "succeeded")},
			},
			"usage_intent": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The intended money movement flow this payout method should be set up for, specified in params.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("payment", "transfer")},
			},
			"payout_method_data": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "If no payout_method provided, used to create the underlying credential that is set up for outbound money movement.\nIf a payout_method provided, used to update data on the credential linked to this setup intent.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"bank_account": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "The type specific details of the bank account payout method.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"account_number": schema.StringAttribute{
								Required:    true,
								Description: "The account number or IBAN of the bank account.",
								WriteOnly:   true,
							},
							"bank_account_type": schema.StringAttribute{
								Optional:    true,
								Description: "Closed Enum. The type of the bank account (checking or savings).",
								WriteOnly:   true,
							},
							"branch_number": schema.StringAttribute{
								Optional:    true,
								Description: "The branch number of the bank account, if present.",
								WriteOnly:   true,
							},
							"country": schema.StringAttribute{
								Required:    true,
								Description: "The country code of the bank account.",
								WriteOnly:   true,
							},
							"currency": schema.StringAttribute{
								Required:    true,
								Description: "The currency of the bank account.",
								WriteOnly:   true,
							},
							"routing_number": schema.StringAttribute{
								Optional:    true,
								Description: "The routing number of the bank account, if present.",
								WriteOnly:   true,
							},
							"swift_code": schema.StringAttribute{
								Optional:    true,
								Description: "The swift code of the bank account, if present.",
								WriteOnly:   true,
							},
						},
					},
					"type": schema.StringAttribute{
						Required:    true,
						Description: "Open Enum. The type of payout method to be created.",
						WriteOnly:   true,
					},
				},
			},
		},
	}
}

func (r *V2MoneyManagementOutboundSetupIntentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan V2MoneyManagementOutboundSetupIntentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config V2MoneyManagementOutboundSetupIntentResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PayoutMethodData"}, []string{"PayoutMethodData", "bank_account"}, []string{"PayoutMethodData", "bank_account", "account_number"}, []string{"PayoutMethodData", "bank_account", "bank_account_type"}, []string{"PayoutMethodData", "bank_account", "branch_number"}, []string{"PayoutMethodData", "bank_account", "country"}, []string{"PayoutMethodData", "bank_account", "currency"}, []string{"PayoutMethodData", "bank_account", "routing_number"}, []string{"PayoutMethodData", "bank_account", "swift_code"}, []string{"PayoutMethodData", "type"}})

	params, err := expandV2MoneyManagementOutboundSetupIntentCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2MoneyManagementOutboundSetupIntent create params", err.Error())
		return
	}

	obj, err := r.client.V2MoneyManagementOutboundSetupIntents.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating V2MoneyManagementOutboundSetupIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementOutboundSetupIntents.B, r.client.V2MoneyManagementOutboundSetupIntents.Key, stripe.FormatURLPath("/v2/money_management/outbound_setup_intents/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementOutboundSetupIntent create raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementOutboundSetupIntent(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementOutboundSetupIntent create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"PayoutMethodData"}, []string{"PayoutMethodData", "bank_account"}, []string{"PayoutMethodData", "bank_account", "account_number"}, []string{"PayoutMethodData", "bank_account", "bank_account_type"}, []string{"PayoutMethodData", "bank_account", "branch_number"}, []string{"PayoutMethodData", "bank_account", "country"}, []string{"PayoutMethodData", "bank_account", "currency"}, []string{"PayoutMethodData", "bank_account", "routing_number"}, []string{"PayoutMethodData", "bank_account", "swift_code"}, []string{"PayoutMethodData", "type"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementOutboundSetupIntentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState V2MoneyManagementOutboundSetupIntentResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2MoneyManagementOutboundSetupIntentResourceModel
	state = priorState

	obj, err := r.client.V2MoneyManagementOutboundSetupIntents.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading V2MoneyManagementOutboundSetupIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementOutboundSetupIntents.B, r.client.V2MoneyManagementOutboundSetupIntents.Key, stripe.FormatURLPath("/v2/money_management/outbound_setup_intents/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementOutboundSetupIntent raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementOutboundSetupIntent(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementOutboundSetupIntent read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"PayoutMethodData"}, []string{"PayoutMethodData", "bank_account"}, []string{"PayoutMethodData", "bank_account", "account_number"}, []string{"PayoutMethodData", "bank_account", "bank_account_type"}, []string{"PayoutMethodData", "bank_account", "branch_number"}, []string{"PayoutMethodData", "bank_account", "country"}, []string{"PayoutMethodData", "bank_account", "currency"}, []string{"PayoutMethodData", "bank_account", "routing_number"}, []string{"PayoutMethodData", "bank_account", "swift_code"}, []string{"PayoutMethodData", "type"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementOutboundSetupIntentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan V2MoneyManagementOutboundSetupIntentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config V2MoneyManagementOutboundSetupIntentResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"PayoutMethodData"}, []string{"PayoutMethodData", "bank_account"}, []string{"PayoutMethodData", "bank_account", "account_number"}, []string{"PayoutMethodData", "bank_account", "bank_account_type"}, []string{"PayoutMethodData", "bank_account", "branch_number"}, []string{"PayoutMethodData", "bank_account", "country"}, []string{"PayoutMethodData", "bank_account", "currency"}, []string{"PayoutMethodData", "bank_account", "routing_number"}, []string{"PayoutMethodData", "bank_account", "swift_code"}, []string{"PayoutMethodData", "type"}})

	var state V2MoneyManagementOutboundSetupIntentResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandV2MoneyManagementOutboundSetupIntentUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2MoneyManagementOutboundSetupIntent update params", err.Error())
		return
	}

	obj, err := r.client.V2MoneyManagementOutboundSetupIntents.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating V2MoneyManagementOutboundSetupIntent", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementOutboundSetupIntents.B, r.client.V2MoneyManagementOutboundSetupIntents.Key, stripe.FormatURLPath("/v2/money_management/outbound_setup_intents/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementOutboundSetupIntent update raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementOutboundSetupIntent(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementOutboundSetupIntent update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"PayoutMethodData"}, []string{"PayoutMethodData", "bank_account"}, []string{"PayoutMethodData", "bank_account", "account_number"}, []string{"PayoutMethodData", "bank_account", "bank_account_type"}, []string{"PayoutMethodData", "bank_account", "branch_number"}, []string{"PayoutMethodData", "bank_account", "country"}, []string{"PayoutMethodData", "bank_account", "currency"}, []string{"PayoutMethodData", "bank_account", "routing_number"}, []string{"PayoutMethodData", "bank_account", "swift_code"}, []string{"PayoutMethodData", "type"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementOutboundSetupIntentResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *V2MoneyManagementOutboundSetupIntentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandV2MoneyManagementOutboundSetupIntentCreate(plan V2MoneyManagementOutboundSetupIntentResourceModel) (*stripe.V2MoneyManagementOutboundSetupIntentCreateParams, error) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentCreateParams{}

	if !plan.PayoutMethod.IsNull() && !plan.PayoutMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PayoutMethodID", "PayoutMethod", plan.PayoutMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payout_method", params)
		}
	}
	if !plan.UsageIntent.IsNull() && !plan.UsageIntent.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "UsageIntent", "UsageIntent", plan.UsageIntent.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "usage_intent", params)
		}
	}
	if !plan.PayoutMethodData.IsNull() && !plan.PayoutMethodData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PayoutMethodData", plan.PayoutMethodData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payout_method_data", params)
		}
	}

	return params, nil
}

func expandV2MoneyManagementOutboundSetupIntentUpdate(plan V2MoneyManagementOutboundSetupIntentResourceModel, state V2MoneyManagementOutboundSetupIntentResourceModel) (*stripe.V2MoneyManagementOutboundSetupIntentUpdateParams, error) {
	params := &stripe.V2MoneyManagementOutboundSetupIntentUpdateParams{}

	if !plan.PayoutMethod.Equal(state.PayoutMethod) && !plan.PayoutMethod.IsNull() && !plan.PayoutMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PayoutMethodID", "PayoutMethod", plan.PayoutMethod.ValueString()) {
			if !plan.PayoutMethod.Equal(state.PayoutMethod) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payout_method", params)
			}
		}
	}
	if !plan.PayoutMethodData.Equal(state.PayoutMethodData) && !plan.PayoutMethodData.IsNull() && !plan.PayoutMethodData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PayoutMethodData", plan.PayoutMethodData) {
			if !plan.PayoutMethodData.Equal(state.PayoutMethodData) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "payout_method_data", params)
			}
		}
	}

	return params, nil
}

func flattenV2MoneyManagementOutboundSetupIntent(obj *stripe.V2MoneyManagementOutboundSetupIntent, state *V2MoneyManagementOutboundSetupIntentResourceModel) error {
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
			if valueCreated, err := flattenPlainValue(rawValueCreated, types.StringType, "created", "raw response"); err != nil {
				return err
			} else {
				if typedCreated, ok := valueCreated.(types.String); ok {
					state.Created = typedCreated
				}
			}
		} else if !hasRaw {
			if responseValueCreated, ok := plainFromResponseField(obj, "Created"); ok {
				if valueCreated, err := flattenPlainValue(responseValueCreated, types.StringType, "created", "response struct"); err != nil {
					return err
				} else {
					if typedCreated, ok := valueCreated.(types.String); ok {
						state.Created = typedCreated
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
		if true {
			if rawValuePayoutMethod, rawOk := plainValueAtPath(raw, "payout_method"); rawOk {
				if typedPayoutMethod, ok := plainToStringIDValue(rawValuePayoutMethod); ok {
					state.PayoutMethod = typedPayoutMethod
				}
			} else if !hasRaw {
				if responseValuePayoutMethod, ok := plainFromResponseField(obj, "PayoutMethod"); ok {
					if typedPayoutMethod, ok := plainToStringIDValue(responseValuePayoutMethod); ok {
						state.PayoutMethod = typedPayoutMethod
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
	{
		if rawValueUsageIntent, rawOk := plainValueAtPath(raw, "usage_intent"); rawOk {
			if valueUsageIntent, err := flattenPlainValue(rawValueUsageIntent, types.StringType, "usage_intent", "raw response"); err != nil {
				return err
			} else {
				if typedUsageIntent, ok := valueUsageIntent.(types.String); ok {
					state.UsageIntent = typedUsageIntent
				}
			}
		} else if !hasRaw {
			if responseValueUsageIntent, ok := plainFromResponseField(obj, "UsageIntent"); ok {
				if valueUsageIntent, err := flattenPlainValue(responseValueUsageIntent, types.StringType, "usage_intent", "response struct"); err != nil {
					return err
				} else {
					if typedUsageIntent, ok := valueUsageIntent.(types.String); ok {
						state.UsageIntent = typedUsageIntent
					}
				}
			}
		}
	}
	return nil
}
