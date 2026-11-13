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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &V2MoneyManagementFinancialAddressResource{}

var _ resource.ResourceWithConfigure = &V2MoneyManagementFinancialAddressResource{}

var _ resource.ResourceWithImportState = &V2MoneyManagementFinancialAddressResource{}

func NewV2MoneyManagementFinancialAddressResource() resource.Resource {
	return &V2MoneyManagementFinancialAddressResource{}
}

type V2MoneyManagementFinancialAddressResource struct {
	client *stripe.Client
}

type V2MoneyManagementFinancialAddressResourceModel struct {
	Object             types.String `tfsdk:"object"`
	BankAccount        types.Object `tfsdk:"bank_account"`
	Created            types.String `tfsdk:"created"`
	FinancialAccount   types.String `tfsdk:"financial_account"`
	ID                 types.String `tfsdk:"id"`
	Livemode           types.Bool   `tfsdk:"livemode"`
	SettlementCurrency types.String `tfsdk:"settlement_currency"`
	Status             types.String `tfsdk:"status"`
	Type               types.String `tfsdk:"type"`
}

func (r *V2MoneyManagementFinancialAddressResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *V2MoneyManagementFinancialAddressResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v2_money_management_financial_address"
}

func (r *V2MoneyManagementFinancialAddressResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A FinancialAddress contains information needed to transfer money to a Financial Account. A Financial Account can have more than one Financial Address.",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value of the object field.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("v2.money_management.financial_address")},
			},
			"bank_account": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Bank account details for this FinancialAddress.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"aba": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "ABA bank account details (US).",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account_holder_name": schema.StringAttribute{
								Computed:      true,
								Description:   "The name of the account holder.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"account_number": schema.StringAttribute{
								Computed:      true,
								Description:   "The full account number. Empty while the bank account is pending allocation.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"address": schema.SingleNestedAttribute{
								Computed:      true,
								Description:   "The address of the account holder.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
								Attributes: map[string]schema.Attribute{
									"city": schema.StringAttribute{
										Computed:      true,
										Description:   "City.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"country": schema.StringAttribute{
										Computed:      true,
										Description:   "Country.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"line1": schema.StringAttribute{
										Computed:      true,
										Description:   "Address line 1.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"line2": schema.StringAttribute{
										Computed:      true,
										Description:   "Address line 2.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"postal_code": schema.StringAttribute{
										Computed:      true,
										Description:   "Postal code.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"state": schema.StringAttribute{
										Computed:      true,
										Description:   "State or province.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
									"town": schema.StringAttribute{
										Computed:      true,
										Description:   "Town or suburb.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									},
								},
							},
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "The name of the bank.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"routing_number": schema.StringAttribute{
								Computed:      true,
								Description:   "The ABA routing number.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"country": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The country of the bank account. Always included in the response, but null while the bank\naccount is pending and its country is not yet known.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"currency": schema.StringAttribute{
						Required:      true,
						Description:   "Open Enum. The currency of the bank account.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"iban": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "IBAN bank account details.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account_holder_name": schema.StringAttribute{
								Computed:      true,
								Description:   "The name of the account holder.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"bank_name": schema.StringAttribute{
								Computed:      true,
								Description:   "The name of the bank.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"iban": schema.StringAttribute{
								Computed:      true,
								Description:   "The full IBAN. Empty while the bank account is pending allocation.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"sort_code": schema.SingleNestedAttribute{
						Computed:      true,
						Description:   "Sort code bank account details (UK).",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"account_holder_name": schema.StringAttribute{
								Computed:      true,
								Description:   "The name of the account holder.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"account_number": schema.StringAttribute{
								Computed:      true,
								Description:   "The full account number. Empty while the bank account is pending allocation.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"sort_code": schema.StringAttribute{
								Computed:      true,
								Description:   "The sort code.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "Open Enum. The type of bank account details.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("aba", "iban", "sort_code")},
					},
				},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				Description:   "The creation timestamp of the FinancialAddress.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"financial_account": schema.StringAttribute{
				Required:      true,
				Description:   "The ID of the FinancialAccount this FinancialAddress corresponds to.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the FinancialAddress.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"settlement_currency": schema.StringAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Closed Enum. The status of the FinancialAddress.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "archived", "failed", "pending")},
			},
			"type": schema.StringAttribute{
				Required:      true,
				Description:   "Open Enum. The type of FinancialAddress.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("bank_account")},
			},
		},
	}
}

func (r *V2MoneyManagementFinancialAddressResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan V2MoneyManagementFinancialAddressResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandV2MoneyManagementFinancialAddressCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2MoneyManagementFinancialAddress create params", err.Error())
		return
	}

	obj, err := r.client.V2MoneyManagementFinancialAddresses.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating V2MoneyManagementFinancialAddress", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementFinancialAddresses.B, r.client.V2MoneyManagementFinancialAddresses.Key, stripe.FormatURLPath("/v2/money_management/financial_addresses/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementFinancialAddress create raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementFinancialAddress(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementFinancialAddress create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementFinancialAddressResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState V2MoneyManagementFinancialAddressResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2MoneyManagementFinancialAddressResourceModel
	state = priorState

	obj, err := r.client.V2MoneyManagementFinancialAddresses.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading V2MoneyManagementFinancialAddress", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2MoneyManagementFinancialAddresses.B, r.client.V2MoneyManagementFinancialAddresses.Key, stripe.FormatURLPath("/v2/money_management/financial_addresses/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2MoneyManagementFinancialAddress raw response", err.Error())
		return
	}

	if err := flattenV2MoneyManagementFinancialAddress(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening V2MoneyManagementFinancialAddress read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *V2MoneyManagementFinancialAddressResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "V2MoneyManagementFinancialAddress does not support updates")
}

func (r *V2MoneyManagementFinancialAddressResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *V2MoneyManagementFinancialAddressResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandV2MoneyManagementFinancialAddressCreate(plan V2MoneyManagementFinancialAddressResourceModel) (*stripe.V2MoneyManagementFinancialAddressCreateParams, error) {
	params := &stripe.V2MoneyManagementFinancialAddressCreateParams{}

	if !plan.BankAccount.IsNull() && !plan.BankAccount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BankAccount", plan.BankAccount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "bank_account", params)
		}
	}
	if !plan.FinancialAccount.IsNull() && !plan.FinancialAccount.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FinancialAccount", "FinancialAccount", plan.FinancialAccount.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "financial_account", params)
		}
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}

	return params, nil
}

func flattenV2MoneyManagementFinancialAddress(obj *stripe.V2MoneyManagementFinancialAddress, state *V2MoneyManagementFinancialAddressResourceModel) error {
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
		assignedBankAccount := false
		hadRawBankAccount := false
		if rawValueBankAccount, rawOk := plainValueAtPath(raw, "bank_account"); rawOk {
			hadRawBankAccount = true
			if rawValueBankAccount != nil {
				sourceBankAccount := applyConfiguredKeyedListShapes(rawValueBankAccount, attrValueToPlain(state.BankAccount))
				if valueBankAccount, err := flattenPlainValue(sourceBankAccount, types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "account_number": types.StringType, "address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}}, "bank_name": types.StringType, "routing_number": types.StringType}}, "country": types.StringType, "currency": types.StringType, "iban": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "bank_name": types.StringType, "iban": types.StringType}}, "sort_code": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "account_number": types.StringType, "sort_code": types.StringType}}, "type": types.StringType}}, "bank_account", "raw response"); err != nil {
					return err
				} else {
					if typedBankAccount, ok := valueBankAccount.(types.Object); ok {
						state.BankAccount = typedBankAccount
						assignedBankAccount = true
					}
				}
			}
		}
		if !assignedBankAccount {
			if !hasRaw {
				if responseValueBankAccount, ok := plainFromResponseField(obj, "BankAccount"); ok {
					sourceBankAccount := applyConfiguredKeyedListShapes(responseValueBankAccount, attrValueToPlain(state.BankAccount))
					if valueBankAccount, err := flattenPlainValue(
						sourceBankAccount,
						types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "account_number": types.StringType, "address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}}, "bank_name": types.StringType, "routing_number": types.StringType}}, "country": types.StringType, "currency": types.StringType, "iban": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "bank_name": types.StringType, "iban": types.StringType}}, "sort_code": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "account_number": types.StringType, "sort_code": types.StringType}}, "type": types.StringType}},
						"bank_account",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBankAccount, ok := valueBankAccount.(types.Object); ok {
							state.BankAccount = typedBankAccount
							assignedBankAccount = true
						}
					}
				}
			}
		}
		if !assignedBankAccount && hadRawBankAccount {
			if nullBankAccount, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"aba": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "account_number": types.StringType, "address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}}, "bank_name": types.StringType, "routing_number": types.StringType}}, "country": types.StringType, "currency": types.StringType, "iban": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "bank_name": types.StringType, "iban": types.StringType}}, "sort_code": types.ObjectType{AttrTypes: map[string]attr.Type{"account_holder_name": types.StringType, "account_number": types.StringType, "sort_code": types.StringType}}, "type": types.StringType}}); ok {
				if typedBankAccount, ok := nullBankAccount.(types.Object); ok {
					state.BankAccount = typedBankAccount
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
		if rawValueFinancialAccount, rawOk := plainValueAtPath(raw, "financial_account"); rawOk {
			if valueFinancialAccount, err := flattenPlainValue(rawValueFinancialAccount, types.StringType, "financial_account", "raw response"); err != nil {
				return err
			} else {
				if typedFinancialAccount, ok := valueFinancialAccount.(types.String); ok {
					state.FinancialAccount = typedFinancialAccount
				}
			}
		} else if !hasRaw {
			if responseValueFinancialAccount, ok := plainFromResponseField(obj, "FinancialAccount"); ok {
				if valueFinancialAccount, err := flattenPlainValue(responseValueFinancialAccount, types.StringType, "financial_account", "response struct"); err != nil {
					return err
				} else {
					if typedFinancialAccount, ok := valueFinancialAccount.(types.String); ok {
						state.FinancialAccount = typedFinancialAccount
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
		if rawValueSettlementCurrency, rawOk := plainValueAtPath(raw, "settlement_currency"); rawOk {
			if valueSettlementCurrency, err := flattenPlainValue(rawValueSettlementCurrency, types.StringType, "settlement_currency", "raw response"); err != nil {
				return err
			} else {
				if typedSettlementCurrency, ok := valueSettlementCurrency.(types.String); ok {
					state.SettlementCurrency = typedSettlementCurrency
				}
			}
		} else if !hasRaw {
			if responseValueSettlementCurrency, ok := plainFromResponseField(obj, "SettlementCurrency"); ok {
				if valueSettlementCurrency, err := flattenPlainValue(responseValueSettlementCurrency, types.StringType, "settlement_currency", "response struct"); err != nil {
					return err
				} else {
					if typedSettlementCurrency, ok := valueSettlementCurrency.(types.String); ok {
						state.SettlementCurrency = typedSettlementCurrency
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
		if rawValueType, rawOk := plainValueAtPath(raw, "type"); rawOk {
			if valueType, err := flattenPlainValue(rawValueType, types.StringType, "type", "raw response"); err != nil {
				return err
			} else {
				if typedType, ok := valueType.(types.String); ok {
					state.Type = typedType
				}
			}
		} else if !hasRaw {
			if responseValueType, ok := plainFromResponseField(obj, "Type"); ok {
				if valueType, err := flattenPlainValue(responseValueType, types.StringType, "type", "response struct"); err != nil {
					return err
				} else {
					if typedType, ok := valueType.(types.String); ok {
						state.Type = typedType
					}
				}
			}
		}
	}
	return nil
}
