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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &V2CoreVaultUSBankAccountResource{}

var _ resource.ResourceWithConfigure = &V2CoreVaultUSBankAccountResource{}

var _ resource.ResourceWithImportState = &V2CoreVaultUSBankAccountResource{}

func NewV2CoreVaultUSBankAccountResource() resource.Resource {
	return &V2CoreVaultUSBankAccountResource{}
}

type V2CoreVaultUSBankAccountResource struct {
	client *stripe.Client
}

type V2CoreVaultUSBankAccountResourceModel struct {
	Object                      types.String `tfsdk:"object"`
	Archived                    types.Bool   `tfsdk:"archived"`
	BankAccountType             types.String `tfsdk:"bank_account_type"`
	BankName                    types.String `tfsdk:"bank_name"`
	Created                     types.String `tfsdk:"created"`
	FedwireRoutingNumber        types.String `tfsdk:"fedwire_routing_number"`
	FinancialConnectionsAccount types.String `tfsdk:"financial_connections_account"`
	ID                          types.String `tfsdk:"id"`
	Last4                       types.String `tfsdk:"last4"`
	Livemode                    types.Bool   `tfsdk:"livemode"`
	Restricted                  types.Bool   `tfsdk:"restricted"`
	RoutingNumber               types.String `tfsdk:"routing_number"`
	SupportedCurrencies         types.List   `tfsdk:"supported_currencies"`
	Verification                types.Object `tfsdk:"verification"`
	AccountNumber               types.String `tfsdk:"account_number"`
	Currency                    types.String `tfsdk:"currency"`
}

func (r *V2CoreVaultUSBankAccountResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *V2CoreVaultUSBankAccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v2_core_vault_us_bank_account"
}

func (r *V2CoreVaultUSBankAccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Use the USBankAccounts API to create and manage US bank accounts objects that you can use to receive funds. Note that these are not interchangeable with v1 Tokens.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value of the object field.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("v2.core.vault.us_bank_account")},
			},
			"archived": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether this USBankAccount object was archived.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"bank_account_type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Closed Enum. The type of bank account (checking or savings).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("checking", "savings")},
			},
			"bank_name": schema.StringAttribute{
				Computed:      true,
				Description:   "The name of the bank this bank account belongs to. This field is populated automatically by Stripe based on the routing number.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				Description:   "Creation time of the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"fedwire_routing_number": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The fedwire routing number of the bank account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"financial_connections_account": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the Financial Connections Account used to create the bank account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the USBankAccount object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"last4": schema.StringAttribute{
				Computed:      true,
				Description:   "The last 4 digits of the account number.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"restricted": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether the US Bank Account is currently unusable for money movement, despite potentially being correctly set up. Please reach out to Stripe Support for more information.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"routing_number": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The ACH routing number of the bank account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"supported_currencies": schema.ListAttribute{
				Computed:      true,
				Description:   "The list of currencies supported by this bank account.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"verification": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The bank account verification details.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"microdeposit_verification_details": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "The microdeposit verification details if the status is awaiting verification.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"expires": schema.StringAttribute{
								Computed:      true,
								Description:   "Time when microdeposits will expire and have to be re-sent.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"microdeposit_type": schema.StringAttribute{
								Computed:      true,
								Description:   "Microdeposit type can be amounts or descriptor_type.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("amounts", "descriptor_code")},
							},
							"sent": schema.StringAttribute{
								Computed:      true,
								Description:   "Time when microdeposits were sent.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The bank account verification status.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("awaiting_verification", "unverified", "verification_failed", "verified")},
					},
				},
			},
			"account_number": schema.StringAttribute{
				Required:      true,
				Description:   "The account number of the bank account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"currency": schema.StringAttribute{
				Required:      true,
				Description:   "The currency of the bank account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
		},
	}
}

func (r *V2CoreVaultUSBankAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan V2CoreVaultUSBankAccountResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config V2CoreVaultUSBankAccountResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"AccountNumber"}, []string{"Currency"}})

	params, err := expandV2CoreVaultUSBankAccountCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2CoreVaultUSBankAccount create params", err.Error())
		return
	}

	obj, err := r.client.V2CoreVaultUSBankAccounts.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating V2CoreVaultUSBankAccount", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2CoreVaultUSBankAccounts.B, r.client.V2CoreVaultUSBankAccounts.Key, stripe.FormatURLPath("/v2/core/vault/us_bank_accounts/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2CoreVaultUSBankAccount create raw response", err.Error())
		return
	}

	if err := flattenV2CoreVaultUSBankAccount(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2CoreVaultUSBankAccount create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"AccountNumber"}, []string{"Currency"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2CoreVaultUSBankAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState V2CoreVaultUSBankAccountResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2CoreVaultUSBankAccountResourceModel
	state = priorState

	obj, err := r.client.V2CoreVaultUSBankAccounts.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading V2CoreVaultUSBankAccount", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2CoreVaultUSBankAccounts.B, r.client.V2CoreVaultUSBankAccounts.Key, stripe.FormatURLPath("/v2/core/vault/us_bank_accounts/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2CoreVaultUSBankAccount raw response", err.Error())
		return
	}

	if err := flattenV2CoreVaultUSBankAccount(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening V2CoreVaultUSBankAccount read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"AccountNumber"}, []string{"Currency"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *V2CoreVaultUSBankAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan V2CoreVaultUSBankAccountResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config V2CoreVaultUSBankAccountResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"AccountNumber"}, []string{"Currency"}})

	var state V2CoreVaultUSBankAccountResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state
	clearWriteOnlyPaths(&diffPlan, [][]string{[]string{"AccountNumber"}, []string{"Currency"}})
	clearWriteOnlyPaths(&diffState, [][]string{[]string{"AccountNumber"}, []string{"Currency"}})

	params, err := expandV2CoreVaultUSBankAccountUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2CoreVaultUSBankAccount update params", err.Error())
		return
	}

	obj, err := r.client.V2CoreVaultUSBankAccounts.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating V2CoreVaultUSBankAccount", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2CoreVaultUSBankAccounts.B, r.client.V2CoreVaultUSBankAccounts.Key, stripe.FormatURLPath("/v2/core/vault/us_bank_accounts/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2CoreVaultUSBankAccount update raw response", err.Error())
		return
	}

	if err := flattenV2CoreVaultUSBankAccount(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2CoreVaultUSBankAccount update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"AccountNumber"}, []string{"Currency"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2CoreVaultUSBankAccountResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *V2CoreVaultUSBankAccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandV2CoreVaultUSBankAccountCreate(plan V2CoreVaultUSBankAccountResourceModel) (*stripe.V2CoreVaultUSBankAccountCreateParams, error) {
	params := &stripe.V2CoreVaultUSBankAccountCreateParams{}

	if !plan.BankAccountType.IsNull() && !plan.BankAccountType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "BankAccountType", "BankAccountType", plan.BankAccountType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "bank_account_type", params)
		}
	}
	if !plan.FedwireRoutingNumber.IsNull() && !plan.FedwireRoutingNumber.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FedwireRoutingNumber", "FedwireRoutingNumber", plan.FedwireRoutingNumber.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "fedwire_routing_number", params)
		}
	}
	if !plan.RoutingNumber.IsNull() && !plan.RoutingNumber.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "RoutingNumber", "RoutingNumber", plan.RoutingNumber.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "routing_number", params)
		}
	}
	if !plan.AccountNumber.IsNull() && !plan.AccountNumber.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "AccountNumber", "AccountNumber", plan.AccountNumber.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "account_number", params)
		}
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}

	return params, nil
}

func expandV2CoreVaultUSBankAccountUpdate(plan V2CoreVaultUSBankAccountResourceModel, state V2CoreVaultUSBankAccountResourceModel) (*stripe.V2CoreVaultUSBankAccountUpdateParams, error) {
	params := &stripe.V2CoreVaultUSBankAccountUpdateParams{}

	if !plan.FedwireRoutingNumber.Equal(state.FedwireRoutingNumber) && !plan.FedwireRoutingNumber.IsNull() && !plan.FedwireRoutingNumber.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FedwireRoutingNumber", "FedwireRoutingNumber", plan.FedwireRoutingNumber.ValueString()) {
			if !plan.FedwireRoutingNumber.Equal(state.FedwireRoutingNumber) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "fedwire_routing_number", params)
			}
		}
	}
	if !plan.RoutingNumber.Equal(state.RoutingNumber) && !plan.RoutingNumber.IsNull() && !plan.RoutingNumber.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "RoutingNumber", "RoutingNumber", plan.RoutingNumber.ValueString()) {
			if !plan.RoutingNumber.Equal(state.RoutingNumber) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "routing_number", params)
			}
		}
	}

	return params, nil
}

func flattenV2CoreVaultUSBankAccount(obj *stripe.V2CoreVaultUSBankAccount, state *V2CoreVaultUSBankAccountResourceModel) error {
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
		if rawValueArchived, rawOk := plainValueAtPath(raw, "archived"); rawOk {
			if valueArchived, err := flattenPlainValue(rawValueArchived, types.BoolType, "archived", "raw response"); err != nil {
				return err
			} else {
				if typedArchived, ok := valueArchived.(types.Bool); ok {
					state.Archived = typedArchived
				}
			}
		} else if !hasRaw {
			if responseValueArchived, ok := plainFromResponseField(obj, "Archived"); ok {
				if valueArchived, err := flattenPlainValue(responseValueArchived, types.BoolType, "archived", "response struct"); err != nil {
					return err
				} else {
					if typedArchived, ok := valueArchived.(types.Bool); ok {
						state.Archived = typedArchived
					}
				}
			}
		}
	}
	{
		if rawValueBankAccountType, rawOk := plainValueAtPath(raw, "bank_account_type"); rawOk {
			if valueBankAccountType, err := flattenPlainValue(rawValueBankAccountType, types.StringType, "bank_account_type", "raw response"); err != nil {
				return err
			} else {
				if typedBankAccountType, ok := valueBankAccountType.(types.String); ok {
					state.BankAccountType = typedBankAccountType
				}
			}
		} else if !hasRaw {
			if responseValueBankAccountType, ok := plainFromResponseField(obj, "BankAccountType"); ok {
				if valueBankAccountType, err := flattenPlainValue(responseValueBankAccountType, types.StringType, "bank_account_type", "response struct"); err != nil {
					return err
				} else {
					if typedBankAccountType, ok := valueBankAccountType.(types.String); ok {
						state.BankAccountType = typedBankAccountType
					}
				}
			}
		}
	}
	{
		if rawValueBankName, rawOk := plainValueAtPath(raw, "bank_name"); rawOk {
			if valueBankName, err := flattenPlainValue(rawValueBankName, types.StringType, "bank_name", "raw response"); err != nil {
				return err
			} else {
				if typedBankName, ok := valueBankName.(types.String); ok {
					state.BankName = typedBankName
				}
			}
		} else if !hasRaw {
			if responseValueBankName, ok := plainFromResponseField(obj, "BankName"); ok {
				if valueBankName, err := flattenPlainValue(responseValueBankName, types.StringType, "bank_name", "response struct"); err != nil {
					return err
				} else {
					if typedBankName, ok := valueBankName.(types.String); ok {
						state.BankName = typedBankName
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
		if rawValueFedwireRoutingNumber, rawOk := plainValueAtPath(raw, "fedwire_routing_number"); rawOk {
			if valueFedwireRoutingNumber, err := flattenPlainValue(rawValueFedwireRoutingNumber, types.StringType, "fedwire_routing_number", "raw response"); err != nil {
				return err
			} else {
				if typedFedwireRoutingNumber, ok := valueFedwireRoutingNumber.(types.String); ok {
					state.FedwireRoutingNumber = typedFedwireRoutingNumber
				}
			}
		} else if !hasRaw {
			if responseValueFedwireRoutingNumber, ok := plainFromResponseField(obj, "FedwireRoutingNumber"); ok {
				if valueFedwireRoutingNumber, err := flattenPlainValue(responseValueFedwireRoutingNumber, types.StringType, "fedwire_routing_number", "response struct"); err != nil {
					return err
				} else {
					if typedFedwireRoutingNumber, ok := valueFedwireRoutingNumber.(types.String); ok {
						state.FedwireRoutingNumber = typedFedwireRoutingNumber
					}
				}
			}
		}
	}
	{
		if rawValueFinancialConnectionsAccount, rawOk := plainValueAtPath(raw, "financial_connections_account"); rawOk {
			if valueFinancialConnectionsAccount, err := flattenPlainValue(rawValueFinancialConnectionsAccount, types.StringType, "financial_connections_account", "raw response"); err != nil {
				return err
			} else {
				if typedFinancialConnectionsAccount, ok := valueFinancialConnectionsAccount.(types.String); ok {
					state.FinancialConnectionsAccount = typedFinancialConnectionsAccount
				}
			}
		} else if !hasRaw {
			if responseValueFinancialConnectionsAccount, ok := plainFromResponseField(obj, "FinancialConnectionsAccount"); ok {
				if valueFinancialConnectionsAccount, err := flattenPlainValue(responseValueFinancialConnectionsAccount, types.StringType, "financial_connections_account", "response struct"); err != nil {
					return err
				} else {
					if typedFinancialConnectionsAccount, ok := valueFinancialConnectionsAccount.(types.String); ok {
						state.FinancialConnectionsAccount = typedFinancialConnectionsAccount
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
		if rawValueLast4, rawOk := plainValueAtPath(raw, "last4"); rawOk {
			if valueLast4, err := flattenPlainValue(rawValueLast4, types.StringType, "last4", "raw response"); err != nil {
				return err
			} else {
				if typedLast4, ok := valueLast4.(types.String); ok {
					state.Last4 = typedLast4
				}
			}
		} else if !hasRaw {
			if responseValueLast4, ok := plainFromResponseField(obj, "Last4"); ok {
				if valueLast4, err := flattenPlainValue(responseValueLast4, types.StringType, "last4", "response struct"); err != nil {
					return err
				} else {
					if typedLast4, ok := valueLast4.(types.String); ok {
						state.Last4 = typedLast4
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
		if rawValueRestricted, rawOk := plainValueAtPath(raw, "restricted"); rawOk {
			if valueRestricted, err := flattenPlainValue(rawValueRestricted, types.BoolType, "restricted", "raw response"); err != nil {
				return err
			} else {
				if typedRestricted, ok := valueRestricted.(types.Bool); ok {
					state.Restricted = typedRestricted
				}
			}
		} else if !hasRaw {
			if responseValueRestricted, ok := plainFromResponseField(obj, "Restricted"); ok {
				if valueRestricted, err := flattenPlainValue(responseValueRestricted, types.BoolType, "restricted", "response struct"); err != nil {
					return err
				} else {
					if typedRestricted, ok := valueRestricted.(types.Bool); ok {
						state.Restricted = typedRestricted
					}
				}
			}
		}
	}
	{
		if rawValueRoutingNumber, rawOk := plainValueAtPath(raw, "routing_number"); rawOk {
			if valueRoutingNumber, err := flattenPlainValue(rawValueRoutingNumber, types.StringType, "routing_number", "raw response"); err != nil {
				return err
			} else {
				if typedRoutingNumber, ok := valueRoutingNumber.(types.String); ok {
					state.RoutingNumber = typedRoutingNumber
				}
			}
		} else if !hasRaw {
			if responseValueRoutingNumber, ok := plainFromResponseField(obj, "RoutingNumber"); ok {
				if valueRoutingNumber, err := flattenPlainValue(responseValueRoutingNumber, types.StringType, "routing_number", "response struct"); err != nil {
					return err
				} else {
					if typedRoutingNumber, ok := valueRoutingNumber.(types.String); ok {
						state.RoutingNumber = typedRoutingNumber
					}
				}
			}
		}
	}
	{
		if rawValueSupportedCurrencies, rawOk := plainValueAtPath(raw, "supported_currencies"); rawOk {
			if valueSupportedCurrencies, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueSupportedCurrencies, attrValueToPlain(state.SupportedCurrencies)), types.ListType{ElemType: types.StringType}, "supported_currencies", "raw response"); err != nil {
				return err
			} else {
				if typedSupportedCurrencies, ok := valueSupportedCurrencies.(types.List); ok {
					state.SupportedCurrencies = typedSupportedCurrencies
				}
			}
		} else if !hasRaw {
			if responseValueSupportedCurrencies, ok := plainFromResponseField(obj, "SupportedCurrencies"); ok {
				if valueSupportedCurrencies, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueSupportedCurrencies, attrValueToPlain(state.SupportedCurrencies)),
					types.ListType{ElemType: types.StringType},
					"supported_currencies",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedSupportedCurrencies, ok := valueSupportedCurrencies.(types.List); ok {
						state.SupportedCurrencies = typedSupportedCurrencies
					}
				}
			}
		}
	}
	{
		assignedVerification := false
		hadRawVerification := false
		if rawValueVerification, rawOk := plainValueAtPath(raw, "verification"); rawOk {
			hadRawVerification = true
			if rawValueVerification != nil {
				sourceVerification := applyConfiguredKeyedListShapes(rawValueVerification, attrValueToPlain(state.Verification))
				if valueVerification, err := flattenPlainValue(sourceVerification, types.ObjectType{AttrTypes: map[string]attr.Type{"microdeposit_verification_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires": types.StringType, "microdeposit_type": types.StringType, "sent": types.StringType}}, "status": types.StringType}}, "verification", "raw response"); err != nil {
					return err
				} else {
					if typedVerification, ok := valueVerification.(types.Object); ok {
						state.Verification = typedVerification
						assignedVerification = true
					}
				}
			}
		}
		if !assignedVerification {
			if !hasRaw {
				if responseValueVerification, ok := plainFromResponseField(obj, "Verification"); ok {
					sourceVerification := applyConfiguredKeyedListShapes(responseValueVerification, attrValueToPlain(state.Verification))
					if valueVerification, err := flattenPlainValue(
						sourceVerification,
						types.ObjectType{AttrTypes: map[string]attr.Type{"microdeposit_verification_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires": types.StringType, "microdeposit_type": types.StringType, "sent": types.StringType}}, "status": types.StringType}},
						"verification",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedVerification, ok := valueVerification.(types.Object); ok {
							state.Verification = typedVerification
							assignedVerification = true
						}
					}
				}
			}
		}
		if !assignedVerification && hadRawVerification {
			if nullVerification, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"microdeposit_verification_details": types.ObjectType{AttrTypes: map[string]attr.Type{"expires": types.StringType, "microdeposit_type": types.StringType, "sent": types.StringType}}, "status": types.StringType}}); ok {
				if typedVerification, ok := nullVerification.(types.Object); ok {
					state.Verification = typedVerification
				}
			}
		}
	}
	return nil
}
