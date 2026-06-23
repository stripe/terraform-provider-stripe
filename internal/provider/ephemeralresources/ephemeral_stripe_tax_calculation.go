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

var _ ephemeral.EphemeralResource = &TaxCalculationEphemeralResource{}
var _ ephemeral.EphemeralResourceWithConfigure = &TaxCalculationEphemeralResource{}

func NewTaxCalculationEphemeralResource() ephemeral.EphemeralResource {
	return &TaxCalculationEphemeralResource{}
}

type TaxCalculationEphemeralResource struct {
	client *stripe.Client
}

type TaxCalculationResourceModel struct {
	Object             types.String `tfsdk:"object"`
	AmountTotal        types.Int64  `tfsdk:"amount_total"`
	Currency           types.String `tfsdk:"currency"`
	Customer           types.String `tfsdk:"customer"`
	CustomerDetails    types.Object `tfsdk:"customer_details"`
	ExpiresAt          types.Int64  `tfsdk:"expires_at"`
	ID                 types.String `tfsdk:"id"`
	LineItems          types.List   `tfsdk:"line_items"`
	Livemode           types.Bool   `tfsdk:"livemode"`
	ShipFromDetails    types.Object `tfsdk:"ship_from_details"`
	ShippingCost       types.Object `tfsdk:"shipping_cost"`
	TaxAmountExclusive types.Int64  `tfsdk:"tax_amount_exclusive"`
	TaxAmountInclusive types.Int64  `tfsdk:"tax_amount_inclusive"`
	TaxBreakdown       types.List   `tfsdk:"tax_breakdown"`
	TaxDate            types.Int64  `tfsdk:"tax_date"`
}

func (r *TaxCalculationEphemeralResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tax_calculation"
}

func (r *TaxCalculationEphemeralResource) Schema(_ context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = ephemeralSchema.Schema{
		Description: "A Tax Calculation allows you to calculate the tax to collect from your customer.\n\nRelated guide: [Calculate tax in your custom payment flow](https://docs.stripe.com/tax/custom)",
		Attributes: map[string]ephemeralSchema.Attribute{
			"object": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "String representing the object's type. Objects of the same type share the same value.",
				Validators:  []validator.String{stringvalidator.OneOf("tax.calculation")},
			},
			"amount_total": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "Total amount after taxes in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).",
			},
			"currency": ephemeralSchema.StringAttribute{
				Required:    true,
				Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
			},
			"customer": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The ID of an existing [Customer](https://docs.stripe.com/api/customers/object) used for the resource.",
			},
			"customer_details": ephemeralSchema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"address": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The customer's postal address (for example, home or business location).",
						Attributes: map[string]ephemeralSchema.Attribute{
							"city": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "City, district, suburb, town, or village.",
							},
							"country": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
							},
							"line1": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Address line 1, such as the street, PO Box, or company name.",
							},
							"line2": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Address line 2, such as the apartment, suite, unit, or building.",
							},
							"postal_code": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "ZIP or postal code.",
							},
							"state": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "State/province as an [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) subdivision code, without country prefix, such as \"NY\" or \"TX\".",
							},
						},
					},
					"address_source": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The type of customer address provided.",
						Validators:  []validator.String{stringvalidator.OneOf("billing", "shipping")},
					},
					"ip_address": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The customer's IP address (IPv4 or IPv6).",
					},
					"tax_ids": ephemeralSchema.ListNestedAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The customer's tax IDs (for example, EU VAT numbers).",
						NestedObject: ephemeralSchema.NestedAttributeObject{
							Attributes: map[string]ephemeralSchema.Attribute{
								"type": ephemeralSchema.StringAttribute{
									Required:    true,
									Description: "The type of the tax ID, one of `ad_nrt`, `ar_cuit`, `eu_vat`, `bo_tin`, `br_cnpj`, `br_cpf`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `eu_oss_vat`, `hr_oib`, `pe_ruc`, `ro_tin`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, `vn_tin`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `no_voec`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `pl_nip`, `it_cf`, `fo_vat`, `gi_tin`, `py_ruc`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `li_vat`, `lk_vat`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, `al_tin`, `bh_vat`, `kz_bin`, `ng_tin`, `om_vat`, `de_stn`, `ch_uid`, `tz_vat`, `uz_vat`, `uz_tin`, `md_vat`, `ma_vat`, `by_tin`, `ao_tin`, `bs_tin`, `bb_tin`, `cd_nif`, `mr_nif`, `me_pib`, `zw_tin`, `ba_tin`, `gn_nif`, `mk_vat`, `sr_fin`, `sn_ninea`, `am_tin`, `np_pan`, `tj_tin`, `ug_tin`, `zm_tin`, `kh_tin`, `aw_tin`, `az_tin`, `bd_bin`, `bj_ifu`, `et_tin`, `kg_tin`, `la_tin`, `cm_niu`, `cv_nif`, `bf_ifu`, or `unknown`",
									Validators:  []validator.String{stringvalidator.OneOf("ad_nrt", "ae_trn", "al_tin", "am_tin", "ao_tin", "ar_cuit", "au_abn", "au_arn", "aw_tin", "az_tin", "ba_tin", "bb_tin", "bd_bin", "bf_ifu", "bg_uic", "bh_vat", "bj_ifu", "bo_tin", "br_cnpj", "br_cpf", "bs_tin", "by_tin", "ca_bn", "ca_gst_hst", "ca_pst_bc", "ca_pst_mb", "ca_pst_sk", "ca_qst", "cd_nif", "ch_uid", "ch_vat", "cl_tin", "cm_niu", "cn_tin", "co_nit", "cr_tin", "cv_nif", "de_stn", "do_rcn", "ec_ruc", "eg_tin", "es_cif", "et_tin", "eu_oss_vat", "eu_vat", "fo_vat", "gb_vat", "ge_vat", "gi_tin", "gn_nif", "hk_br", "hr_oib", "hu_tin", "id_npwp", "il_vat", "in_gst", "is_vat", "it_cf", "jp_cn", "jp_rn", "jp_trn", "ke_pin", "kg_tin", "kh_tin", "kr_brn", "kz_bin", "la_tin", "li_uid", "li_vat", "lk_vat", "ma_vat", "md_vat", "me_pib", "mk_vat", "mr_nif", "mx_rfc", "my_frp", "my_itn", "my_sst", "ng_tin", "no_vat", "no_voec", "np_pan", "nz_gst", "om_vat", "pe_ruc", "ph_tin", "pl_nip", "py_ruc", "ro_tin", "rs_pib", "ru_inn", "ru_kpp", "sa_vat", "sg_gst", "sg_uen", "si_tin", "sn_ninea", "sr_fin", "sv_nit", "th_vat", "tj_tin", "tr_tin", "tw_vat", "tz_vat", "ua_vat", "ug_tin", "unknown", "us_ein", "uy_ruc", "uz_tin", "uz_vat", "ve_rif", "vn_tin", "za_vat", "zm_tin", "zw_tin")},
								},
								"value": ephemeralSchema.StringAttribute{
									Required:    true,
									Description: "The value of the tax ID.",
								},
							},
						},
					},
					"taxability_override": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The taxability override used for taxation.",
						Validators:  []validator.String{stringvalidator.OneOf("customer_exempt", "none", "reverse_charge")},
					},
				},
			},
			"expires_at": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "Timestamp of date at which the tax calculation will expire.",
			},
			"id": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier for the calculation.",
			},
			"line_items": ephemeralSchema.ListNestedAttribute{
				Required:    true,
				Description: "The list of items the customer is purchasing.",
				NestedObject: ephemeralSchema.NestedAttributeObject{
					Attributes: map[string]ephemeralSchema.Attribute{
						"amount": ephemeralSchema.Int64Attribute{
							Required:    true,
							Description: "A positive integer representing the line item's total price in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).\nIf `tax_behavior=inclusive`, then this amount includes taxes. Otherwise, taxes are calculated on top of this amount.",
						},
						"metadata": ephemeralSchema.MapAttribute{
							Optional:    true,
							Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
							ElementType: types.StringType,
						},
						"product": ephemeralSchema.StringAttribute{
							Optional:    true,
							Description: "If provided, the product's `tax_code` will be used as the line item's `tax_code`.",
						},
						"quantity": ephemeralSchema.Int64Attribute{
							Optional:    true,
							Description: "The number of units of the item being purchased. Used to calculate the per-unit price from the total `amount` for the line. For example, if `amount=100` and `quantity=4`, the calculated unit price is 25.",
						},
						"reference": ephemeralSchema.StringAttribute{
							Optional:    true,
							Description: "A custom identifier for this line item, which must be unique across the line items in the calculation. The reference helps identify each line item in exported [tax reports](https://docs.stripe.com/tax/reports).",
						},
						"tax_behavior": ephemeralSchema.StringAttribute{
							Optional:    true,
							Description: "Specifies whether the `amount` includes taxes. Defaults to `exclusive`.",
						},
						"tax_code": ephemeralSchema.StringAttribute{
							Optional:    true,
							Description: "A [tax code](https://docs.stripe.com/tax/tax-categories) ID to use for this line item. If not provided, we will use the tax code from the provided `product` param. If neither `tax_code` nor `product` is provided, we will use the default tax code from your Tax Settings.",
						},
					},
				},
			},
			"livemode": ephemeralSchema.BoolAttribute{
				Computed:    true,
				Description: "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
			},
			"ship_from_details": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The details of the ship from location, such as the address.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"address": ephemeralSchema.SingleNestedAttribute{
						Required: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"city": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "City, district, suburb, town, or village.",
							},
							"country": ephemeralSchema.StringAttribute{
								Required:    true,
								Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
							},
							"line1": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Address line 1, such as the street, PO Box, or company name.",
							},
							"line2": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "Address line 2, such as the apartment, suite, unit, or building.",
							},
							"postal_code": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "ZIP or postal code.",
							},
							"state": ephemeralSchema.StringAttribute{
								Optional:    true,
								Computed:    true,
								Description: "State/province as an [ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2) subdivision code, without country prefix, such as \"NY\" or \"TX\".",
							},
						},
					},
				},
			},
			"shipping_cost": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The shipping cost details for the calculation.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"amount": ephemeralSchema.Int64Attribute{
						Optional:    true,
						Computed:    true,
						Description: "The shipping amount in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units). If `tax_behavior=inclusive`, then this amount includes taxes. Otherwise, taxes were calculated on top of this amount.",
					},
					"amount_tax": ephemeralSchema.Int64Attribute{
						Computed:    true,
						Description: "The amount of tax calculated for shipping, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).",
					},
					"shipping_rate": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The ID of an existing [ShippingRate](https://docs.stripe.com/api/shipping_rates/object).",
					},
					"tax_behavior": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Specifies whether the `amount` includes taxes. If `tax_behavior=inclusive`, then the amount includes taxes.",
						Validators:  []validator.String{stringvalidator.OneOf("exclusive", "inclusive")},
					},
					"tax_breakdown": ephemeralSchema.ListNestedAttribute{
						Computed:    true,
						Description: "Detailed account of taxes relevant to shipping cost.",
						NestedObject: ephemeralSchema.NestedAttributeObject{
							Attributes: map[string]ephemeralSchema.Attribute{
								"amount": ephemeralSchema.Int64Attribute{
									Computed:    true,
									Description: "The amount of tax, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).",
								},
								"jurisdiction": ephemeralSchema.SingleNestedAttribute{
									Computed: true,

									Attributes: map[string]ephemeralSchema.Attribute{
										"country": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
										},
										"display_name": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "A human-readable name for the jurisdiction imposing the tax.",
										},
										"level": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "Indicates the level of the jurisdiction imposing the tax.",
											Validators:  []validator.String{stringvalidator.OneOf("city", "country", "county", "district", "state")},
										},
										"state": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "[ISO 3166-2 subdivision code](https://en.wikipedia.org/wiki/ISO_3166-2), without country prefix. For example, \"NY\" for New York, United States.",
										},
									},
								},
								"sourcing": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "Indicates whether the jurisdiction was determined by the origin (merchant's address) or destination (customer's address).",
									Validators:  []validator.String{stringvalidator.OneOf("destination", "origin")},
								},
								"tax_rate_details": ephemeralSchema.SingleNestedAttribute{
									Computed:    true,
									Description: "Details regarding the rate for this tax. This field will be `null` when the tax is not imposed, for example if the product is exempt from tax.",
									Attributes: map[string]ephemeralSchema.Attribute{
										"display_name": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "A localized display name for tax type, intended to be human-readable. For example, \"Local Sales and Use Tax\", \"Value-added tax (VAT)\", or \"Umsatzsteuer (USt.)\".",
										},
										"percentage_decimal": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "The tax rate percentage as a string. For example, 8.5% is represented as \"8.5\".",
										},
										"tax_type": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "The tax type, such as `vat` or `sales_tax`.",
											Validators:  []validator.String{stringvalidator.OneOf("amusement_tax", "communications_tax", "gst", "hst", "igst", "jct", "lease_tax", "pst", "qst", "retail_delivery_fee", "rst", "sales_tax", "service_tax", "vat")},
										},
									},
								},
								"taxability_reason": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.",
									Validators:  []validator.String{stringvalidator.OneOf("customer_exempt", "not_collecting", "not_subject_to_tax", "not_supported", "portion_product_exempt", "portion_reduced_rated", "portion_standard_rated", "product_exempt", "product_exempt_holiday", "proportionally_rated", "reduced_rated", "reverse_charge", "standard_rated", "taxable_basis_reduced", "zero_rated")},
								},
								"taxable_amount": ephemeralSchema.Int64Attribute{
									Computed:    true,
									Description: "The amount on which tax is calculated, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).",
								},
							},
						},
					},
					"tax_code": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The [tax code](https://docs.stripe.com/tax/tax-categories) ID used for shipping.",
					},
				},
			},
			"tax_amount_exclusive": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "The amount of tax to be collected on top of the line item prices.",
			},
			"tax_amount_inclusive": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "The amount of tax already included in the line item prices.",
			},
			"tax_breakdown": ephemeralSchema.ListNestedAttribute{
				Computed:    true,
				Description: "Breakdown of individual tax amounts that add up to the total.",
				NestedObject: ephemeralSchema.NestedAttributeObject{
					Attributes: map[string]ephemeralSchema.Attribute{
						"amount": ephemeralSchema.Int64Attribute{
							Computed:    true,
							Description: "The amount of tax, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).",
						},
						"inclusive": ephemeralSchema.BoolAttribute{
							Computed:    true,
							Description: "Specifies whether the tax amount is included in the line item amount.",
						},
						"tax_rate_details": ephemeralSchema.SingleNestedAttribute{
							Computed: true,

							Attributes: map[string]ephemeralSchema.Attribute{
								"country": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
								},
								"flat_amount": ephemeralSchema.SingleNestedAttribute{
									Computed:    true,
									Description: "The amount of the tax rate when the `rate_type` is `flat_amount`. Tax rates with `rate_type` `percentage` can vary based on the transaction, resulting in this field being `null`. This field exposes the amount and currency of the flat tax rate.",
									Attributes: map[string]ephemeralSchema.Attribute{
										"amount": ephemeralSchema.Int64Attribute{
											Computed:    true,
											Description: "Amount of the tax when the `rate_type` is `flat_amount`. This positive integer represents how much to charge in the smallest currency unit (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).",
										},
										"currency": ephemeralSchema.StringAttribute{
											Computed:    true,
											Description: "Three-letter ISO currency code, in lowercase.",
										},
									},
								},
								"percentage_decimal": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "The tax rate percentage as a string. For example, 8.5% is represented as `\"8.5\"`.",
								},
								"rate_type": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "Indicates the type of tax rate applied to the taxable amount. This value can be `null` when no tax applies to the location. This field is only present for TaxRates created by Stripe Tax.",
									Validators:  []validator.String{stringvalidator.OneOf("flat_amount", "percentage")},
								},
								"state": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
								},
								"tax_type": ephemeralSchema.StringAttribute{
									Computed:    true,
									Description: "The tax type, such as `vat` or `sales_tax`.",
									Validators:  []validator.String{stringvalidator.OneOf("amusement_tax", "communications_tax", "gst", "hst", "igst", "jct", "lease_tax", "pst", "qst", "retail_delivery_fee", "rst", "sales_tax", "service_tax", "vat")},
								},
							},
						},
						"taxability_reason": ephemeralSchema.StringAttribute{
							Computed:    true,
							Description: "The reasoning behind this tax, for example, if the product is tax exempt. We might extend the possible values for this field to support new tax rules.",
							Validators:  []validator.String{stringvalidator.OneOf("customer_exempt", "not_collecting", "not_subject_to_tax", "not_supported", "portion_product_exempt", "portion_reduced_rated", "portion_standard_rated", "product_exempt", "product_exempt_holiday", "proportionally_rated", "reduced_rated", "reverse_charge", "standard_rated", "taxable_basis_reduced", "zero_rated")},
						},
						"taxable_amount": ephemeralSchema.Int64Attribute{
							Computed:    true,
							Description: "The amount on which tax is calculated, in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).",
						},
					},
				},
			},
			"tax_date": ephemeralSchema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The calculation uses the tax rules and rates that are in effect at this timestamp. You can use a date up to 31 days in the past or up to 31 days in the future. If you use a future date, Stripe doesn't guarantee that the expected tax rules and rate being used match the actual rules and rate that will be in effect on that date. We deploy tax changes before their effective date, but not within a fixed window.",
			},
		},
	}
}

func (r *TaxCalculationEphemeralResource) Configure(_ context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
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

func expandTaxCalculationCreate(plan TaxCalculationResourceModel) (*stripe.TaxCalculationCreateParams, error) {
	params := &stripe.TaxCalculationCreateParams{}

	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Customer", "Customer", plan.Customer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer", params)
		}
	}
	if !plan.CustomerDetails.IsNull() && !plan.CustomerDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CustomerDetails", plan.CustomerDetails) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer_details", params)
		}
	}
	if !plan.LineItems.IsNull() && !plan.LineItems.IsUnknown() {
		if !assignAttrValueToNamedField(params, "LineItems", plan.LineItems) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "line_items", params)
		}
	}
	if !plan.ShipFromDetails.IsNull() && !plan.ShipFromDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShipFromDetails", plan.ShipFromDetails) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "ship_from_details", params)
		}
	}
	if !plan.ShippingCost.IsNull() && !plan.ShippingCost.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ShippingCost", plan.ShippingCost) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "shipping_cost", params)
		}
	}
	if !plan.TaxDate.IsNull() && !plan.TaxDate.IsUnknown() {
		params.TaxDate = stripe.Int64(plan.TaxDate.ValueInt64())
	}

	return params, nil
}

func flattenTaxCalculation(obj *stripe.TaxCalculation, state *TaxCalculationResourceModel) error {
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
		if rawValueAmountTotal, rawOk := plainValueAtPath(raw, "amount_total"); rawOk {
			if valueAmountTotal, err := flattenPlainValue(rawValueAmountTotal, types.Int64Type, "amount_total", "raw response"); err != nil {
				return err
			} else {
				if typedAmountTotal, ok := valueAmountTotal.(types.Int64); ok {
					state.AmountTotal = typedAmountTotal
				}
			}
		} else if !hasRaw {
			if responseValueAmountTotal, ok := plainFromResponseField(obj, "AmountTotal"); ok {
				if valueAmountTotal, err := flattenPlainValue(responseValueAmountTotal, types.Int64Type, "amount_total", "response struct"); err != nil {
					return err
				} else {
					if typedAmountTotal, ok := valueAmountTotal.(types.Int64); ok {
						state.AmountTotal = typedAmountTotal
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
		if rawValueCustomer, rawOk := plainValueAtPath(raw, "customer"); rawOk {
			if valueCustomer, err := flattenPlainValue(rawValueCustomer, types.StringType, "customer", "raw response"); err != nil {
				return err
			} else {
				if typedCustomer, ok := valueCustomer.(types.String); ok {
					state.Customer = typedCustomer
				}
			}
		} else if !hasRaw {
			if responseValueCustomer, ok := plainFromResponseField(obj, "Customer"); ok {
				if valueCustomer, err := flattenPlainValue(responseValueCustomer, types.StringType, "customer", "response struct"); err != nil {
					return err
				} else {
					if typedCustomer, ok := valueCustomer.(types.String); ok {
						state.Customer = typedCustomer
					}
				}
			}
		}
	}
	{
		assignedCustomerDetails := false
		hadRawCustomerDetails := false
		if rawValueCustomerDetails, rawOk := plainValueAtPath(raw, "customer_details"); rawOk {
			hadRawCustomerDetails = true
			if rawValueCustomerDetails != nil {
				sourceCustomerDetails := applyConfiguredKeyedListShapes(rawValueCustomerDetails, attrValueToPlain(state.CustomerDetails))
				if valueCustomerDetails, err := flattenPlainValue(sourceCustomerDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "address_source": types.StringType, "ip_address": types.StringType, "tax_ids": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}}, "taxability_override": types.StringType}}, "customer_details", "raw response"); err != nil {
					return err
				} else {
					if typedCustomerDetails, ok := valueCustomerDetails.(types.Object); ok {
						state.CustomerDetails = typedCustomerDetails
						assignedCustomerDetails = true
					}
				}
			}
		}
		if !assignedCustomerDetails {
			if !hasRaw {
				if responseValueCustomerDetails, ok := plainFromResponseField(obj, "CustomerDetails"); ok {
					sourceCustomerDetails := applyConfiguredKeyedListShapes(responseValueCustomerDetails, attrValueToPlain(state.CustomerDetails))
					if valueCustomerDetails, err := flattenPlainValue(
						sourceCustomerDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "address_source": types.StringType, "ip_address": types.StringType, "tax_ids": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}}, "taxability_override": types.StringType}},
						"customer_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCustomerDetails, ok := valueCustomerDetails.(types.Object); ok {
							state.CustomerDetails = typedCustomerDetails
							assignedCustomerDetails = true
						}
					}
				}
			}
		}
		if !assignedCustomerDetails && hadRawCustomerDetails {
			if nullCustomerDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "address_source": types.StringType, "ip_address": types.StringType, "tax_ids": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "value": types.StringType}}}, "taxability_override": types.StringType}}); ok {
				if typedCustomerDetails, ok := nullCustomerDetails.(types.Object); ok {
					state.CustomerDetails = typedCustomerDetails
				}
			}
		}
	}
	{
		if rawValueExpiresAt, rawOk := plainValueAtPath(raw, "expires_at"); rawOk {
			if valueExpiresAt, err := flattenPlainValue(rawValueExpiresAt, types.Int64Type, "expires_at", "raw response"); err != nil {
				return err
			} else {
				if typedExpiresAt, ok := valueExpiresAt.(types.Int64); ok {
					state.ExpiresAt = typedExpiresAt
				}
			}
		} else if !hasRaw {
			if responseValueExpiresAt, ok := plainFromResponseField(obj, "ExpiresAt"); ok {
				if valueExpiresAt, err := flattenPlainValue(responseValueExpiresAt, types.Int64Type, "expires_at", "response struct"); err != nil {
					return err
				} else {
					if typedExpiresAt, ok := valueExpiresAt.(types.Int64); ok {
						state.ExpiresAt = typedExpiresAt
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
		if rawValueLineItems, rawOk := plainValueAtPath(raw, "line_items"); rawOk {
			rawPlainLineItems := extractListObjectData(rawValueLineItems)
			if valueLineItems, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawPlainLineItems, attrValueToPlain(state.LineItems)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "metadata": types.MapType{ElemType: types.StringType}, "product": types.StringType, "quantity": types.Int64Type, "reference": types.StringType, "tax_behavior": types.StringType, "tax_code": types.StringType}}}, "line_items", "raw response"); err != nil {
				return err
			} else {
				if typedLineItems, ok := valueLineItems.(types.List); ok {
					state.LineItems = typedLineItems
				}
			}
		} else if !hasRaw {
			if responseValueLineItems, ok := plainFromResponseField(obj, "LineItems"); ok {
				fallbackPlainLineItems := extractListObjectData(responseValueLineItems)
				if valueLineItems, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(fallbackPlainLineItems, attrValueToPlain(state.LineItems)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "metadata": types.MapType{ElemType: types.StringType}, "product": types.StringType, "quantity": types.Int64Type, "reference": types.StringType, "tax_behavior": types.StringType, "tax_code": types.StringType}}},
					"line_items",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedLineItems, ok := valueLineItems.(types.List); ok {
						state.LineItems = typedLineItems
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
		assignedShipFromDetails := false
		hadRawShipFromDetails := false
		if rawValueShipFromDetails, rawOk := plainValueAtPath(raw, "ship_from_details"); rawOk {
			hadRawShipFromDetails = true
			if rawValueShipFromDetails != nil {
				sourceShipFromDetails := applyConfiguredKeyedListShapes(rawValueShipFromDetails, attrValueToPlain(state.ShipFromDetails))
				if valueShipFromDetails, err := flattenPlainValue(sourceShipFromDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}, "ship_from_details", "raw response"); err != nil {
					return err
				} else {
					if typedShipFromDetails, ok := valueShipFromDetails.(types.Object); ok {
						state.ShipFromDetails = typedShipFromDetails
						assignedShipFromDetails = true
					}
				}
			}
		}
		if !assignedShipFromDetails {
			if !hasRaw {
				if responseValueShipFromDetails, ok := plainFromResponseField(obj, "ShipFromDetails"); ok {
					sourceShipFromDetails := applyConfiguredKeyedListShapes(responseValueShipFromDetails, attrValueToPlain(state.ShipFromDetails))
					if valueShipFromDetails, err := flattenPlainValue(
						sourceShipFromDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}},
						"ship_from_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedShipFromDetails, ok := valueShipFromDetails.(types.Object); ok {
							state.ShipFromDetails = typedShipFromDetails
							assignedShipFromDetails = true
						}
					}
				}
			}
		}
		if !assignedShipFromDetails && hadRawShipFromDetails {
			if nullShipFromDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}}}); ok {
				if typedShipFromDetails, ok := nullShipFromDetails.(types.Object); ok {
					state.ShipFromDetails = typedShipFromDetails
				}
			}
		}
	}
	{
		assignedShippingCost := false
		hadRawShippingCost := false
		if rawValueShippingCost, rawOk := plainValueAtPath(raw, "shipping_cost"); rawOk {
			hadRawShippingCost = true
			if rawValueShippingCost != nil {
				sourceShippingCost := preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawValueShippingCost, attrValueToPlain(state.ShippingCost)), attrValueToPlain(state.ShippingCost)), attrValueToPlain(state.ShippingCost))
				if valueShippingCost, err := flattenPlainValue(sourceShippingCost, types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_tax": types.Int64Type, "shipping_rate": types.StringType, "tax_behavior": types.StringType, "tax_breakdown": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "jurisdiction": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "display_name": types.StringType, "level": types.StringType, "state": types.StringType}}, "sourcing": types.StringType, "tax_rate_details": types.ObjectType{AttrTypes: map[string]attr.Type{"display_name": types.StringType, "percentage_decimal": types.StringType, "tax_type": types.StringType}}, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}, "tax_code": types.StringType}}, "shipping_cost", "raw response"); err != nil {
					return err
				} else {
					if typedShippingCost, ok := valueShippingCost.(types.Object); ok {
						state.ShippingCost = typedShippingCost
						assignedShippingCost = true
					}
				}
			}
		}
		if !assignedShippingCost {
			if !hasRaw {
				if responseValueShippingCost, ok := plainFromResponseField(obj, "ShippingCost"); ok {
					sourceShippingCost := preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(responseValueShippingCost, attrValueToPlain(state.ShippingCost)), attrValueToPlain(state.ShippingCost)), attrValueToPlain(state.ShippingCost))
					if valueShippingCost, err := flattenPlainValue(
						sourceShippingCost,
						types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_tax": types.Int64Type, "shipping_rate": types.StringType, "tax_behavior": types.StringType, "tax_breakdown": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "jurisdiction": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "display_name": types.StringType, "level": types.StringType, "state": types.StringType}}, "sourcing": types.StringType, "tax_rate_details": types.ObjectType{AttrTypes: map[string]attr.Type{"display_name": types.StringType, "percentage_decimal": types.StringType, "tax_type": types.StringType}}, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}, "tax_code": types.StringType}},
						"shipping_cost",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedShippingCost, ok := valueShippingCost.(types.Object); ok {
							state.ShippingCost = typedShippingCost
							assignedShippingCost = true
						}
					}
				}
			}
		}
		if !assignedShippingCost && hadRawShippingCost {
			if nullShippingCost, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "amount_tax": types.Int64Type, "shipping_rate": types.StringType, "tax_behavior": types.StringType, "tax_breakdown": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "jurisdiction": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "display_name": types.StringType, "level": types.StringType, "state": types.StringType}}, "sourcing": types.StringType, "tax_rate_details": types.ObjectType{AttrTypes: map[string]attr.Type{"display_name": types.StringType, "percentage_decimal": types.StringType, "tax_type": types.StringType}}, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}, "tax_code": types.StringType}}); ok {
				if typedShippingCost, ok := nullShippingCost.(types.Object); ok {
					state.ShippingCost = typedShippingCost
				}
			}
		}
	}
	{
		if rawValueTaxAmountExclusive, rawOk := plainValueAtPath(raw, "tax_amount_exclusive"); rawOk {
			if valueTaxAmountExclusive, err := flattenPlainValue(rawValueTaxAmountExclusive, types.Int64Type, "tax_amount_exclusive", "raw response"); err != nil {
				return err
			} else {
				if typedTaxAmountExclusive, ok := valueTaxAmountExclusive.(types.Int64); ok {
					state.TaxAmountExclusive = typedTaxAmountExclusive
				}
			}
		} else if !hasRaw {
			if responseValueTaxAmountExclusive, ok := plainFromResponseField(obj, "TaxAmountExclusive"); ok {
				if valueTaxAmountExclusive, err := flattenPlainValue(responseValueTaxAmountExclusive, types.Int64Type, "tax_amount_exclusive", "response struct"); err != nil {
					return err
				} else {
					if typedTaxAmountExclusive, ok := valueTaxAmountExclusive.(types.Int64); ok {
						state.TaxAmountExclusive = typedTaxAmountExclusive
					}
				}
			}
		}
	}
	{
		if rawValueTaxAmountInclusive, rawOk := plainValueAtPath(raw, "tax_amount_inclusive"); rawOk {
			if valueTaxAmountInclusive, err := flattenPlainValue(rawValueTaxAmountInclusive, types.Int64Type, "tax_amount_inclusive", "raw response"); err != nil {
				return err
			} else {
				if typedTaxAmountInclusive, ok := valueTaxAmountInclusive.(types.Int64); ok {
					state.TaxAmountInclusive = typedTaxAmountInclusive
				}
			}
		} else if !hasRaw {
			if responseValueTaxAmountInclusive, ok := plainFromResponseField(obj, "TaxAmountInclusive"); ok {
				if valueTaxAmountInclusive, err := flattenPlainValue(responseValueTaxAmountInclusive, types.Int64Type, "tax_amount_inclusive", "response struct"); err != nil {
					return err
				} else {
					if typedTaxAmountInclusive, ok := valueTaxAmountInclusive.(types.Int64); ok {
						state.TaxAmountInclusive = typedTaxAmountInclusive
					}
				}
			}
		}
	}
	{
		if rawValueTaxBreakdown, rawOk := plainValueAtPath(raw, "tax_breakdown"); rawOk {
			if valueTaxBreakdown, err := flattenPlainValue(preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(rawValueTaxBreakdown, attrValueToPlain(state.TaxBreakdown)), attrValueToPlain(state.TaxBreakdown)), attrValueToPlain(state.TaxBreakdown)), types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "inclusive": types.BoolType, "tax_rate_details": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "flat_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType}}, "percentage_decimal": types.StringType, "rate_type": types.StringType, "state": types.StringType, "tax_type": types.StringType}}, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}}, "tax_breakdown", "raw response"); err != nil {
				return err
			} else {
				if typedTaxBreakdown, ok := valueTaxBreakdown.(types.List); ok {
					state.TaxBreakdown = typedTaxBreakdown
				}
			}
		} else if !hasRaw {
			if responseValueTaxBreakdown, ok := plainFromResponseField(obj, "TaxBreakdown"); ok {
				if valueTaxBreakdown, err := flattenPlainValue(
					preserveEquivalentDecimalStringLeaves(suppressUnconfiguredDecimalMirrorLeaves(applyConfiguredKeyedListShapes(responseValueTaxBreakdown, attrValueToPlain(state.TaxBreakdown)), attrValueToPlain(state.TaxBreakdown)), attrValueToPlain(state.TaxBreakdown)),
					types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "inclusive": types.BoolType, "tax_rate_details": types.ObjectType{AttrTypes: map[string]attr.Type{"country": types.StringType, "flat_amount": types.ObjectType{AttrTypes: map[string]attr.Type{"amount": types.Int64Type, "currency": types.StringType}}, "percentage_decimal": types.StringType, "rate_type": types.StringType, "state": types.StringType, "tax_type": types.StringType}}, "taxability_reason": types.StringType, "taxable_amount": types.Int64Type}}},
					"tax_breakdown",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedTaxBreakdown, ok := valueTaxBreakdown.(types.List); ok {
						state.TaxBreakdown = typedTaxBreakdown
					}
				}
			}
		}
	}
	{
		if rawValueTaxDate, rawOk := plainValueAtPath(raw, "tax_date"); rawOk {
			if valueTaxDate, err := flattenPlainValue(rawValueTaxDate, types.Int64Type, "tax_date", "raw response"); err != nil {
				return err
			} else {
				if typedTaxDate, ok := valueTaxDate.(types.Int64); ok {
					state.TaxDate = typedTaxDate
				}
			}
		} else if !hasRaw {
			if responseValueTaxDate, ok := plainFromResponseField(obj, "TaxDate"); ok {
				if valueTaxDate, err := flattenPlainValue(responseValueTaxDate, types.Int64Type, "tax_date", "response struct"); err != nil {
					return err
				} else {
					if typedTaxDate, ok := valueTaxDate.(types.Int64); ok {
						state.TaxDate = typedTaxDate
					}
				}
			}
		}
	}
	return nil
}

func (r *TaxCalculationEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var config TaxCalculationResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTaxCalculationCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building TaxCalculation ephemeral params", err.Error())
		return
	}

	obj, err := r.client.V1TaxCalculations.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error opening TaxCalculation ephemeral resource", err.Error())
		return
	}

	result := config
	if err := flattenTaxCalculation(obj, &result); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxCalculation ephemeral response", err.Error())
		return
	}
	normalizeUnknownValues(&result)
	diags = resp.Result.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
}
