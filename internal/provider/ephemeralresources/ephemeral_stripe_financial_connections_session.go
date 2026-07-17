//
// File generated from our OpenAPI spec
//

package ephemeralresources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	ephemeralSchema "github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ ephemeral.EphemeralResource = &FinancialConnectionsSessionEphemeralResource{}
var _ ephemeral.EphemeralResourceWithConfigure = &FinancialConnectionsSessionEphemeralResource{}

func NewFinancialConnectionsSessionEphemeralResource() ephemeral.EphemeralResource {
	return &FinancialConnectionsSessionEphemeralResource{}
}

type FinancialConnectionsSessionEphemeralResource struct {
	client *stripe.Client
}

type FinancialConnectionsSessionResourceModel struct {
	Object        types.String `tfsdk:"object"`
	AccountHolder types.Object `tfsdk:"account_holder"`
	ClientSecret  types.String `tfsdk:"client_secret"`
	Filters       types.Object `tfsdk:"filters"`
	ID            types.String `tfsdk:"id"`
	Livemode      types.Bool   `tfsdk:"livemode"`
	Permissions   types.List   `tfsdk:"permissions"`
	Prefetch      types.List   `tfsdk:"prefetch"`
	ReturnURL     types.String `tfsdk:"return_url"`
}

func (r *FinancialConnectionsSessionEphemeralResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_financial_connections_session"
}

func (r *FinancialConnectionsSessionEphemeralResource) Schema(_ context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = ephemeralSchema.Schema{
		Description: "A Financial Connections Session is the secure way to programmatically launch the client-side Stripe.js modal that lets your users link their accounts.",
		Attributes: map[string]ephemeralSchema.Attribute{
			"object": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "String representing the object's type. Objects of the same type share the same value.",
				Validators:  []validator.String{stringvalidator.OneOf("financial_connections.session")},
			},
			"account_holder": ephemeralSchema.SingleNestedAttribute{
				Required:    true,
				Description: "The account holder for whom accounts are collected in this session.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"account": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The ID of the Stripe account that this account belongs to. Only available when `account_holder.type` is `account`.",
					},
					"customer": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The ID for an Account representing a customer that this account belongs to. Only available when `account_holder.type` is `customer`.",
					},
					"customer_account": ephemeralSchema.StringAttribute{
						Optional: true,
						Computed: true,
					},
					"type": ephemeralSchema.StringAttribute{
						Required:    true,
						Description: "Type of account holder that this account belongs to.",
						Validators:  []validator.String{stringvalidator.OneOf("account", "customer")},
					},
				},
			},
			"client_secret": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "A value that will be passed to the client to launch the authentication flow.",
				Sensitive:   true,
			},
			"filters": ephemeralSchema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"account_subcategories": ephemeralSchema.ListAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Restricts the Session to subcategories of accounts that can be linked. Valid subcategories are: `checking`, `savings`, `mortgage`, `line_of_credit`, `credit_card`.",
						ElementType: types.StringType,
					},
					"countries": ephemeralSchema.ListAttribute{
						Optional:    true,
						Computed:    true,
						Description: "List of countries from which to filter accounts.",
						ElementType: types.StringType,
					},
				},
			},
			"id": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier for the object.",
			},
			"livemode": ephemeralSchema.BoolAttribute{
				Computed:    true,
				Description: "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
			},
			"permissions": ephemeralSchema.ListAttribute{
				Required:    true,
				Description: "Permissions requested for accounts collected during this session.",
				ElementType: types.StringType,
			},
			"prefetch": ephemeralSchema.ListAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Data features requested to be retrieved upon account creation.",
				ElementType: types.StringType,
			},
			"return_url": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.",
			},
		},
	}
}

func (r *FinancialConnectionsSessionEphemeralResource) Configure(_ context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stripe.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Ephemeral Resource Configure Type", fmt.Sprintf("Expected *stripe.Client, got: %T", req.ProviderData))
		return
	}

	r.client = client
}

func expandFinancialConnectionsSessionCreate(plan FinancialConnectionsSessionResourceModel) (*stripe.FinancialConnectionsSessionCreateParams, error) {
	params := &stripe.FinancialConnectionsSessionCreateParams{}

	if !plan.AccountHolder.IsNull() && !plan.AccountHolder.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AccountHolder", plan.AccountHolder) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "account_holder", params)
		}
	}
	if !plan.Filters.IsNull() && !plan.Filters.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Filters", plan.Filters) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "filters", params)
		}
	}
	if !plan.Permissions.IsNull() && !plan.Permissions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Permissions", plan.Permissions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "permissions", params)
		}
	}
	if !plan.Prefetch.IsNull() && !plan.Prefetch.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Prefetch", plan.Prefetch) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "prefetch", params)
		}
	}
	if !plan.ReturnURL.IsNull() && !plan.ReturnURL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReturnURL", "ReturnURL", plan.ReturnURL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "return_url", params)
		}
	}

	return params, nil
}

func flattenFinancialConnectionsSession(obj *stripe.FinancialConnectionsSession, state *FinancialConnectionsSessionResourceModel) error {
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
		assignedAccountHolder := false
		hadRawAccountHolder := false
		if rawValueAccountHolder, rawOk := plainValueAtPath(raw, "account_holder"); rawOk {
			hadRawAccountHolder = true
			if rawValueAccountHolder != nil {
				sourceAccountHolder := applyConfiguredKeyedListShapes(rawValueAccountHolder, attrValueToPlain(state.AccountHolder))
				if valueAccountHolder, err := flattenPlainValue(sourceAccountHolder, types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "customer": types.StringType, "customer_account": types.StringType, "type": types.StringType}}, "account_holder", "raw response"); err != nil {
					return err
				} else {
					if typedAccountHolder, ok := valueAccountHolder.(types.Object); ok {
						state.AccountHolder = typedAccountHolder
						assignedAccountHolder = true
					}
				}
			}
		}
		if !assignedAccountHolder {
			if !hasRaw {
				if responseValueAccountHolder, ok := plainFromResponseField(obj, "AccountHolder"); ok {
					sourceAccountHolder := applyConfiguredKeyedListShapes(responseValueAccountHolder, attrValueToPlain(state.AccountHolder))
					if valueAccountHolder, err := flattenPlainValue(
						sourceAccountHolder,
						types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "customer": types.StringType, "customer_account": types.StringType, "type": types.StringType}},
						"account_holder",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAccountHolder, ok := valueAccountHolder.(types.Object); ok {
							state.AccountHolder = typedAccountHolder
							assignedAccountHolder = true
						}
					}
				}
			}
		}
		if !assignedAccountHolder && hadRawAccountHolder {
			if nullAccountHolder, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "customer": types.StringType, "customer_account": types.StringType, "type": types.StringType}}); ok {
				if typedAccountHolder, ok := nullAccountHolder.(types.Object); ok {
					state.AccountHolder = typedAccountHolder
				}
			}
		}
	}
	{
		if rawValueClientSecret, rawOk := plainValueAtPath(raw, "client_secret"); rawOk {
			if valueClientSecret, err := flattenPlainValue(rawValueClientSecret, types.StringType, "client_secret", "raw response"); err != nil {
				return err
			} else {
				if typedClientSecret, ok := valueClientSecret.(types.String); ok {
					state.ClientSecret = typedClientSecret
				}
			}
		} else if !hasRaw {
			if responseValueClientSecret, ok := plainFromResponseField(obj, "ClientSecret"); ok {
				if valueClientSecret, err := flattenPlainValue(responseValueClientSecret, types.StringType, "client_secret", "response struct"); err != nil {
					return err
				} else {
					if typedClientSecret, ok := valueClientSecret.(types.String); ok {
						state.ClientSecret = typedClientSecret
					}
				}
			}
		}
	}
	{
		assignedFilters := false
		hadRawFilters := false
		if rawValueFilters, rawOk := plainValueAtPath(raw, "filters"); rawOk {
			hadRawFilters = true
			if rawValueFilters != nil {
				sourceFilters := applyConfiguredKeyedListShapes(rawValueFilters, attrValueToPlain(state.Filters))
				if valueFilters, err := flattenPlainValue(sourceFilters, types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}, "countries": types.ListType{ElemType: types.StringType}}}, "filters", "raw response"); err != nil {
					return err
				} else {
					if typedFilters, ok := valueFilters.(types.Object); ok {
						state.Filters = typedFilters
						assignedFilters = true
					}
				}
			}
		}
		if !assignedFilters {
			if !hasRaw {
				if responseValueFilters, ok := plainFromResponseField(obj, "Filters"); ok {
					sourceFilters := applyConfiguredKeyedListShapes(responseValueFilters, attrValueToPlain(state.Filters))
					if valueFilters, err := flattenPlainValue(
						sourceFilters,
						types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}, "countries": types.ListType{ElemType: types.StringType}}},
						"filters",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFilters, ok := valueFilters.(types.Object); ok {
							state.Filters = typedFilters
							assignedFilters = true
						}
					}
				}
			}
		}
		if !assignedFilters && hadRawFilters {
			if nullFilters, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account_subcategories": types.ListType{ElemType: types.StringType}, "countries": types.ListType{ElemType: types.StringType}}}); ok {
				if typedFilters, ok := nullFilters.(types.Object); ok {
					state.Filters = typedFilters
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
		if rawValuePermissions, rawOk := plainValueAtPath(raw, "permissions"); rawOk {
			if valuePermissions, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValuePermissions, attrValueToPlain(state.Permissions)), types.ListType{ElemType: types.StringType}, "permissions", "raw response"); err != nil {
				return err
			} else {
				if typedPermissions, ok := valuePermissions.(types.List); ok {
					state.Permissions = typedPermissions
				}
			}
		} else if !hasRaw {
			if responseValuePermissions, ok := plainFromResponseField(obj, "Permissions"); ok {
				if valuePermissions, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValuePermissions, attrValueToPlain(state.Permissions)),
					types.ListType{ElemType: types.StringType},
					"permissions",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedPermissions, ok := valuePermissions.(types.List); ok {
						state.Permissions = typedPermissions
					}
				}
			}
		}
	}
	{
		if rawValuePrefetch, rawOk := plainValueAtPath(raw, "prefetch"); rawOk {
			if valuePrefetch, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValuePrefetch, attrValueToPlain(state.Prefetch)), types.ListType{ElemType: types.StringType}, "prefetch", "raw response"); err != nil {
				return err
			} else {
				if typedPrefetch, ok := valuePrefetch.(types.List); ok {
					state.Prefetch = typedPrefetch
				}
			}
		} else if !hasRaw {
			if responseValuePrefetch, ok := plainFromResponseField(obj, "Prefetch"); ok {
				if valuePrefetch, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValuePrefetch, attrValueToPlain(state.Prefetch)),
					types.ListType{ElemType: types.StringType},
					"prefetch",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedPrefetch, ok := valuePrefetch.(types.List); ok {
						state.Prefetch = typedPrefetch
					}
				}
			}
		}
	}
	{
		if rawValueReturnURL, rawOk := plainValueAtPath(raw, "return_url"); rawOk {
			if valueReturnURL, err := flattenPlainValue(rawValueReturnURL, types.StringType, "return_url", "raw response"); err != nil {
				return err
			} else {
				if typedReturnURL, ok := valueReturnURL.(types.String); ok {
					state.ReturnURL = typedReturnURL
				}
			}
		} else if !hasRaw {
			if responseValueReturnURL, ok := plainFromResponseField(obj, "ReturnURL"); ok {
				if valueReturnURL, err := flattenPlainValue(responseValueReturnURL, types.StringType, "return_url", "response struct"); err != nil {
					return err
				} else {
					if typedReturnURL, ok := valueReturnURL.(types.String); ok {
						state.ReturnURL = typedReturnURL
					}
				}
			}
		}
	}
	return nil
}

func (r *FinancialConnectionsSessionEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var config FinancialConnectionsSessionResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandFinancialConnectionsSessionCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building FinancialConnectionsSession ephemeral params", err.Error())
		return
	}

	obj, err := r.client.V1FinancialConnectionsSessions.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error opening FinancialConnectionsSession ephemeral resource", err.Error())
		return
	}

	result := config
	if err := flattenFinancialConnectionsSession(obj, &result); err != nil {
		resp.Diagnostics.AddError("Error flattening FinancialConnectionsSession ephemeral response", err.Error())
		return
	}
	normalizeUnknownValues(&result)
	diags = resp.Result.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
}
