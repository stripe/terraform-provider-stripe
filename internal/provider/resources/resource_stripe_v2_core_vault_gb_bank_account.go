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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v87"
)

var _ resource.Resource = &V2CoreVaultGBBankAccountResource{}

var _ resource.ResourceWithConfigure = &V2CoreVaultGBBankAccountResource{}

var _ resource.ResourceWithImportState = &V2CoreVaultGBBankAccountResource{}

func NewV2CoreVaultGBBankAccountResource() resource.Resource {
	return &V2CoreVaultGBBankAccountResource{}
}

type V2CoreVaultGBBankAccountResource struct {
	client *stripe.Client
}

type V2CoreVaultGBBankAccountResourceModel struct {
	Object              types.String `tfsdk:"object"`
	Archived            types.Bool   `tfsdk:"archived"`
	BankAccountType     types.String `tfsdk:"bank_account_type"`
	BankName            types.String `tfsdk:"bank_name"`
	Created             types.String `tfsdk:"created"`
	ID                  types.String `tfsdk:"id"`
	Last4               types.String `tfsdk:"last4"`
	Livemode            types.Bool   `tfsdk:"livemode"`
	Restricted          types.Bool   `tfsdk:"restricted"`
	SortCode            types.String `tfsdk:"sort_code"`
	SupportedCurrencies types.List   `tfsdk:"supported_currencies"`
	AccountNumber       types.String `tfsdk:"account_number"`
	Currency            types.String `tfsdk:"currency"`
	IBAN                types.String `tfsdk:"iban"`
}

func (r *V2CoreVaultGBBankAccountResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *V2CoreVaultGBBankAccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_v2_core_vault_gb_bank_account"
}

func (r *V2CoreVaultGBBankAccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Use the GBBankAccounts API to create and manage GB bank account objects",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value of the object field.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("v2.core.vault.gb_bank_account")},
			},
			"archived": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether this bank account object was archived. Bank account objects can be archived through\nthe /archive API, and they will not be automatically archived by Stripe. Archived bank account objects\ncannot be used as outbound destinations and will not appear in the outbound destination list.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"bank_account_type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Closed Enum. The type of the bank account (checking or savings).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("checking", "savings")},
			},
			"bank_name": schema.StringAttribute{
				Computed:      true,
				Description:   "The name of the bank.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created": schema.StringAttribute{
				Computed:      true,
				Description:   "Creation time.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "The ID of the GBBankAccount object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"last4": schema.StringAttribute{
				Computed:      true,
				Description:   "The last 4 digits of the account number or IBAN.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"restricted": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether the GB Bank Account is currently unusable for money movement, despite potentially being correctly set up. Please reach out to Stripe Support for more information.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"sort_code": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The Sort Code of the bank account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"supported_currencies": schema.ListAttribute{
				Computed:      true,
				Description:   "The list of currencies supported by this bank account.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"account_number": schema.StringAttribute{
				Optional:      true,
				Description:   "The Account Number of the bank account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"currency": schema.StringAttribute{
				Required:      true,
				Description:   "The currency of the bank account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
			"iban": schema.StringAttribute{
				Optional:      true,
				Description:   "The IBAN of the bank account.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				WriteOnly:     true,
			},
		},
	}
}

func (r *V2CoreVaultGBBankAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan V2CoreVaultGBBankAccountResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config V2CoreVaultGBBankAccountResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"AccountNumber"}, []string{"Currency"}, []string{"IBAN"}})

	params, err := expandV2CoreVaultGBBankAccountCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building V2CoreVaultGBBankAccount create params", err.Error())
		return
	}

	obj, err := r.client.V2CoreVaultGBBankAccounts.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating V2CoreVaultGBBankAccount", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2CoreVaultGBBankAccounts.B, r.client.V2CoreVaultGBBankAccounts.Key, stripe.FormatURLPath("/v2/core/vault/gb_bank_accounts/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2CoreVaultGBBankAccount create raw response", err.Error())
		return
	}

	if err := flattenV2CoreVaultGBBankAccount(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening V2CoreVaultGBBankAccount create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"AccountNumber"}, []string{"Currency"}, []string{"IBAN"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *V2CoreVaultGBBankAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState V2CoreVaultGBBankAccountResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state V2CoreVaultGBBankAccountResourceModel
	state = priorState

	obj, err := r.client.V2CoreVaultGBBankAccounts.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading V2CoreVaultGBBankAccount", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V2CoreVaultGBBankAccounts.B, r.client.V2CoreVaultGBBankAccounts.Key, stripe.FormatURLPath("/v2/core/vault/gb_bank_accounts/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating V2CoreVaultGBBankAccount raw response", err.Error())
		return
	}

	if err := flattenV2CoreVaultGBBankAccount(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening V2CoreVaultGBBankAccount read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"AccountNumber"}, []string{"Currency"}, []string{"IBAN"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *V2CoreVaultGBBankAccountResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "V2CoreVaultGBBankAccount does not support updates")
}

func (r *V2CoreVaultGBBankAccountResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *V2CoreVaultGBBankAccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandV2CoreVaultGBBankAccountCreate(plan V2CoreVaultGBBankAccountResourceModel) (*stripe.V2CoreVaultGBBankAccountCreateParams, error) {
	params := &stripe.V2CoreVaultGBBankAccountCreateParams{}

	if !plan.BankAccountType.IsNull() && !plan.BankAccountType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "BankAccountType", "BankAccountType", plan.BankAccountType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "bank_account_type", params)
		}
	}
	if !plan.SortCode.IsNull() && !plan.SortCode.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SortCode", "SortCode", plan.SortCode.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "sort_code", params)
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
	if !plan.IBAN.IsNull() && !plan.IBAN.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "IBAN", "IBAN", plan.IBAN.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "iban", params)
		}
	}

	return params, nil
}

func flattenV2CoreVaultGBBankAccount(obj *stripe.V2CoreVaultGBBankAccount, state *V2CoreVaultGBBankAccountResourceModel) error {
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
		if rawValueSortCode, rawOk := plainValueAtPath(raw, "sort_code"); rawOk {
			if valueSortCode, err := flattenPlainValue(rawValueSortCode, types.StringType, "sort_code", "raw response"); err != nil {
				return err
			} else {
				if typedSortCode, ok := valueSortCode.(types.String); ok {
					state.SortCode = typedSortCode
				}
			}
		} else if !hasRaw {
			if responseValueSortCode, ok := plainFromResponseField(obj, "SortCode"); ok {
				if valueSortCode, err := flattenPlainValue(responseValueSortCode, types.StringType, "sort_code", "response struct"); err != nil {
					return err
				} else {
					if typedSortCode, ok := valueSortCode.(types.String); ok {
						state.SortCode = typedSortCode
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
	return nil
}
