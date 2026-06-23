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
	"strings"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &TaxIDResource{}

var _ resource.ResourceWithConfigure = &TaxIDResource{}

var _ resource.ResourceWithImportState = &TaxIDResource{}

func NewTaxIDResource() resource.Resource {
	return &TaxIDResource{}
}

type TaxIDResource struct {
	client *stripe.Client
}

type TaxIDResourceModel struct {
	Object          types.String `tfsdk:"object"`
	Country         types.String `tfsdk:"country"`
	Created         types.Int64  `tfsdk:"created"`
	Customer        types.String `tfsdk:"customer"`
	CustomerAccount types.String `tfsdk:"customer_account"`
	ID              types.String `tfsdk:"id"`
	Livemode        types.Bool   `tfsdk:"livemode"`
	Owner           types.Object `tfsdk:"owner"`
	Type            types.String `tfsdk:"type"`
	Value           types.String `tfsdk:"value"`
	Verification    types.Object `tfsdk:"verification"`
}

func (r *TaxIDResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TaxIDResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tax_id"
}

func (r *TaxIDResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "You can add one or multiple tax IDs to a [customer](https://docs.stripe.com/api/customers) or account.\nCustomer and account tax IDs get displayed on related invoices and credit notes.\n\nRelated guides: [Customer tax identification numbers](https://docs.stripe.com/billing/taxes/tax-ids), [Account tax IDs](https://docs.stripe.com/invoicing/connect#account-tax-ids)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("tax_id")},
			},
			"country": schema.StringAttribute{
				Computed:      true,
				Description:   "Two-letter ISO code representing the country of the tax ID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"customer": schema.StringAttribute{
				Required:      true,
				Description:   "ID of the customer.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"customer_account": schema.StringAttribute{
				Computed:      true,
				Description:   "ID of the Account representing the customer.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"owner": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "The account or customer the tax ID belongs to.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"account": schema.StringAttribute{
						Computed:      true,
						Description:   "The account being referenced when `type` is `account`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"application": schema.StringAttribute{
						Computed:      true,
						Description:   "The Connect Application being referenced when `type` is `application`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"customer": schema.StringAttribute{
						Computed:      true,
						Description:   "The customer being referenced when `type` is `customer`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"customer_account": schema.StringAttribute{
						Computed:      true,
						Description:   "The Account representing the customer being referenced when `type` is `customer`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"type": schema.StringAttribute{
						Computed:      true,
						Description:   "Type of owner referenced.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("account", "application", "customer", "self")},
					},
				},
			},
			"type": schema.StringAttribute{
				Required:      true,
				Description:   "Type of the tax ID, one of `ad_nrt`, `ae_trn`, `al_tin`, `am_tin`, `ao_tin`, `ar_cuit`, `au_abn`, `au_arn`, `aw_tin`, `az_tin`, `ba_tin`, `bb_tin`, `bd_bin`, `bf_ifu`, `bg_uic`, `bh_vat`, `bj_ifu`, `bo_tin`, `br_cnpj`, `br_cpf`, `bs_tin`, `by_tin`, `ca_bn`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `ca_qst`, `cd_nif`, `ch_uid`, `ch_vat`, `cl_tin`, `cm_niu`, `cn_tin`, `co_nit`, `cr_tin`, `cv_nif`, `de_stn`, `do_rcn`, `ec_ruc`, `eg_tin`, `es_cif`, `et_tin`, `eu_oss_vat`, `eu_vat`, `fo_vat`, `gb_vat`, `ge_vat`, `gi_tin`, `gn_nif`, `hk_br`, `hr_oib`, `hu_tin`, `id_npwp`, `il_vat`, `in_gst`, `is_vat`, `it_cf`, `jp_cn`, `jp_rn`, `jp_trn`, `ke_pin`, `kg_tin`, `kh_tin`, `kr_brn`, `kz_bin`, `la_tin`, `li_uid`, `li_vat`, `lk_vat`, `ma_vat`, `md_vat`, `me_pib`, `mk_vat`, `mr_nif`, `mx_rfc`, `my_frp`, `my_itn`, `my_sst`, `ng_tin`, `no_vat`, `no_voec`, `np_pan`, `nz_gst`, `om_vat`, `pe_ruc`, `ph_tin`, `pl_nip`, `py_ruc`, `ro_tin`, `rs_pib`, `ru_inn`, `ru_kpp`, `sa_vat`, `sg_gst`, `sg_uen`, `si_tin`, `sn_ninea`, `sr_fin`, `sv_nit`, `th_vat`, `tj_tin`, `tr_tin`, `tw_vat`, `tz_vat`, `ua_vat`, `ug_tin`, `us_ein`, `uy_ruc`, `uz_tin`, `uz_vat`, `ve_rif`, `vn_tin`, `za_vat`, `zm_tin`, or `zw_tin`. Note that some legacy tax IDs have type `unknown`",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Validators:    []validator.String{stringvalidator.OneOf("ad_nrt", "ae_trn", "al_tin", "am_tin", "ao_tin", "ar_cuit", "au_abn", "au_arn", "aw_tin", "az_tin", "ba_tin", "bb_tin", "bd_bin", "bf_ifu", "bg_uic", "bh_vat", "bj_ifu", "bo_tin", "br_cnpj", "br_cpf", "bs_tin", "by_tin", "ca_bn", "ca_gst_hst", "ca_pst_bc", "ca_pst_mb", "ca_pst_sk", "ca_qst", "cd_nif", "ch_uid", "ch_vat", "cl_tin", "cm_niu", "cn_tin", "co_nit", "cr_tin", "cv_nif", "de_stn", "do_rcn", "ec_ruc", "eg_tin", "es_cif", "et_tin", "eu_oss_vat", "eu_vat", "fo_vat", "gb_vat", "ge_vat", "gi_tin", "gn_nif", "hk_br", "hr_oib", "hu_tin", "id_npwp", "il_vat", "in_gst", "is_vat", "it_cf", "jp_cn", "jp_rn", "jp_trn", "ke_pin", "kg_tin", "kh_tin", "kr_brn", "kz_bin", "la_tin", "li_uid", "li_vat", "lk_vat", "ma_vat", "md_vat", "me_pib", "mk_vat", "mr_nif", "mx_rfc", "my_frp", "my_itn", "my_sst", "ng_tin", "no_vat", "no_voec", "np_pan", "nz_gst", "om_vat", "pe_ruc", "ph_tin", "pl_nip", "py_ruc", "ro_tin", "rs_pib", "ru_inn", "ru_kpp", "sa_vat", "sg_gst", "sg_uen", "si_tin", "sn_ninea", "sr_fin", "sv_nit", "th_vat", "tj_tin", "tr_tin", "tw_vat", "tz_vat", "ua_vat", "ug_tin", "unknown", "us_ein", "uy_ruc", "uz_tin", "uz_vat", "ve_rif", "vn_tin", "za_vat", "zm_tin", "zw_tin")},
			},
			"value": schema.StringAttribute{
				Required:      true,
				Description:   "Value of the tax ID.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"verification": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Tax ID verification information.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "Verification status, one of `pending`, `verified`, `unverified`, or `unavailable`.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("pending", "unavailable", "unverified", "verified")},
					},
					"verified_address": schema.StringAttribute{
						Computed:      true,
						Description:   "Verified address.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"verified_name": schema.StringAttribute{
						Computed:      true,
						Description:   "Verified name.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
		},
	}
}

func (r *TaxIDResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TaxIDResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTaxIDCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building TaxID create params", err.Error())
		return
	}

	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedField(params, "Customer", "ID", plan.Customer.ValueString()) {
			resp.Diagnostics.AddError("Error building TaxID create path params", fmt.Sprintf("Failed to assign path parameter %q on %T", "customer", params))
			return
		}
	}
	obj, err := r.client.V1TaxIDs.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating TaxID", err.Error())
		return
	}

	rawReadParams := &stripe.TaxIDRetrieveParams{}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedField(rawReadParams, "Customer", "ID", plan.Customer.ValueString()) {
			resp.Diagnostics.AddError("Error building TaxID read params for raw hydration after create", fmt.Sprintf("Failed to assign path parameter %q on %T", "customer", rawReadParams))
			return
		}
	}

	if err := ensureRawResponse(obj, r.client.V1TaxIDs.B, r.client.V1TaxIDs.Key, stripe.FormatURLPath("/v1/customers/%s/tax_ids/%s", plan.Customer.ValueString(), obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating TaxID create raw response", err.Error())
		return
	}

	if err := flattenTaxID(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxID create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TaxIDResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState TaxIDResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TaxIDResourceModel
	state = priorState

	params := &stripe.TaxIDRetrieveParams{}
	if !state.Customer.IsNull() && !state.Customer.IsUnknown() {
		if !assignStringToNamedField(params, "Customer", "ID", state.Customer.ValueString()) {
			resp.Diagnostics.AddError("Error building TaxID read params", fmt.Sprintf("Failed to assign path parameter %q on %T", "customer", params))
			return
		}
	}

	obj, err := r.client.V1TaxIDs.Retrieve(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error reading TaxID", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TaxIDs.B, r.client.V1TaxIDs.Key, stripe.FormatURLPath("/v1/customers/%s/tax_ids/%s", state.Customer.ValueString(), state.ID.ValueString()), params); err != nil {
		resp.Diagnostics.AddError("Error hydrating TaxID raw response", err.Error())
		return
	}

	if err := flattenTaxID(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxID read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *TaxIDResource) Update(_ context.Context, _ resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError("Update not supported", "TaxID does not support updates")
}

func (r *TaxIDResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TaxIDResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := &stripe.TaxIDDeleteParams{}
	if !state.Customer.IsNull() && !state.Customer.IsUnknown() {
		if !assignStringToNamedField(params, "Customer", "ID", state.Customer.ValueString()) {
			resp.Diagnostics.AddError("Error building TaxID delete params", fmt.Sprintf("Failed to assign path parameter %q on %T", "customer", params))
			return
		}
	}

	_, err := r.client.V1TaxIDs.Delete(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting TaxID", err.Error())
		return
	}
}

func (r *TaxIDResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Unexpected import identifier", fmt.Sprintf("Expected import identifier in the form \"customer/id\", got %q", req.ID))
		return
	}

	diags := resp.State.SetAttribute(ctx, path.Root("customer"), types.StringValue(parts[0]))
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(parts[1]))
	resp.Diagnostics.Append(diags...)
}

func expandTaxIDCreate(plan TaxIDResourceModel) (*stripe.TaxIDCreateParams, error) {
	params := &stripe.TaxIDCreateParams{}

	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}
	if !plan.Value.IsNull() && !plan.Value.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Value", "Value", plan.Value.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "value", params)
		}
	}

	return params, nil
}

func flattenTaxID(obj *stripe.TaxID, state *TaxIDResourceModel) error {
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
		if rawValueCountry, rawOk := plainValueAtPath(raw, "country"); rawOk {
			if valueCountry, err := flattenPlainValue(rawValueCountry, types.StringType, "country", "raw response"); err != nil {
				return err
			} else {
				if typedCountry, ok := valueCountry.(types.String); ok {
					state.Country = typedCountry
				}
			}
		} else if !hasRaw {
			if responseValueCountry, ok := plainFromResponseField(obj, "Country"); ok {
				if valueCountry, err := flattenPlainValue(responseValueCountry, types.StringType, "country", "response struct"); err != nil {
					return err
				} else {
					if typedCountry, ok := valueCountry.(types.String); ok {
						state.Country = typedCountry
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
		if state.Customer.IsNull() || state.Customer.IsUnknown() {
			if rawValueCustomer, rawOk := plainValueAtPath(raw, "customer"); rawOk {
				if typedCustomer, ok := plainToStringIDValue(rawValueCustomer); ok {
					state.Customer = typedCustomer
				}
			} else if !hasRaw {
				if responseValueCustomer, ok := plainFromResponseField(obj, "Customer"); ok {
					if typedCustomer, ok := plainToStringIDValue(responseValueCustomer); ok {
						state.Customer = typedCustomer
					}
				}
			}
		}
	}
	{
		if rawValueCustomerAccount, rawOk := plainValueAtPath(raw, "customer_account"); rawOk {
			if valueCustomerAccount, err := flattenPlainValue(rawValueCustomerAccount, types.StringType, "customer_account", "raw response"); err != nil {
				return err
			} else {
				if typedCustomerAccount, ok := valueCustomerAccount.(types.String); ok {
					state.CustomerAccount = typedCustomerAccount
				}
			}
		} else if !hasRaw {
			if responseValueCustomerAccount, ok := plainFromResponseField(obj, "CustomerAccount"); ok {
				if valueCustomerAccount, err := flattenPlainValue(responseValueCustomerAccount, types.StringType, "customer_account", "response struct"); err != nil {
					return err
				} else {
					if typedCustomerAccount, ok := valueCustomerAccount.(types.String); ok {
						state.CustomerAccount = typedCustomerAccount
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
		assignedOwner := false
		hadRawOwner := false
		if rawValueOwner, rawOk := plainValueAtPath(raw, "owner"); rawOk {
			hadRawOwner = true
			if rawValueOwner != nil {
				sourceOwner := applyConfiguredKeyedListShapes(rawValueOwner, attrValueToPlain(state.Owner))
				if valueOwner, err := flattenPlainValue(sourceOwner, types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "application": types.StringType, "customer": types.StringType, "customer_account": types.StringType, "type": types.StringType}}, "owner", "raw response"); err != nil {
					return err
				} else {
					if typedOwner, ok := valueOwner.(types.Object); ok {
						state.Owner = typedOwner
						assignedOwner = true
					}
				}
			}
		}
		if !assignedOwner {
			if !hasRaw {
				if responseValueOwner, ok := plainFromResponseField(obj, "Owner"); ok {
					sourceOwner := applyConfiguredKeyedListShapes(responseValueOwner, attrValueToPlain(state.Owner))
					if valueOwner, err := flattenPlainValue(
						sourceOwner,
						types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "application": types.StringType, "customer": types.StringType, "customer_account": types.StringType, "type": types.StringType}},
						"owner",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedOwner, ok := valueOwner.(types.Object); ok {
							state.Owner = typedOwner
							assignedOwner = true
						}
					}
				}
			}
		}
		if !assignedOwner && hadRawOwner {
			if nullOwner, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "application": types.StringType, "customer": types.StringType, "customer_account": types.StringType, "type": types.StringType}}); ok {
				if typedOwner, ok := nullOwner.(types.Object); ok {
					state.Owner = typedOwner
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
	{
		if rawValueValue, rawOk := plainValueAtPath(raw, "value"); rawOk {
			if valueValue, err := flattenPlainValue(rawValueValue, types.StringType, "value", "raw response"); err != nil {
				return err
			} else {
				if typedValue, ok := valueValue.(types.String); ok {
					state.Value = typedValue
				}
			}
		} else if !hasRaw {
			if responseValueValue, ok := plainFromResponseField(obj, "Value"); ok {
				if valueValue, err := flattenPlainValue(responseValueValue, types.StringType, "value", "response struct"); err != nil {
					return err
				} else {
					if typedValue, ok := valueValue.(types.String); ok {
						state.Value = typedValue
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
				if valueVerification, err := flattenPlainValue(sourceVerification, types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "verified_address": types.StringType, "verified_name": types.StringType}}, "verification", "raw response"); err != nil {
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
						types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "verified_address": types.StringType, "verified_name": types.StringType}},
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
			if nullVerification, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType, "verified_address": types.StringType, "verified_name": types.StringType}}); ok {
				if typedVerification, ok := nullVerification.(types.Object); ok {
					state.Verification = typedVerification
				}
			}
		}
	}
	return nil
}
