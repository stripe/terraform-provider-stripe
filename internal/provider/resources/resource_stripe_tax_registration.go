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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &TaxRegistrationResource{}

var _ resource.ResourceWithConfigure = &TaxRegistrationResource{}

var _ resource.ResourceWithImportState = &TaxRegistrationResource{}

func NewTaxRegistrationResource() resource.Resource {
	return &TaxRegistrationResource{}
}

type TaxRegistrationResource struct {
	client *stripe.Client
}

type TaxRegistrationResourceModel struct {
	Object         types.String `tfsdk:"object"`
	ActiveFrom     types.Int64  `tfsdk:"active_from"`
	Country        types.String `tfsdk:"country"`
	CountryOptions types.Object `tfsdk:"country_options"`
	Created        types.Int64  `tfsdk:"created"`
	ExpiresAt      types.Int64  `tfsdk:"expires_at"`
	ID             types.String `tfsdk:"id"`
	Livemode       types.Bool   `tfsdk:"livemode"`
	Status         types.String `tfsdk:"status"`
}

func (r *TaxRegistrationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TaxRegistrationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tax_registration"
}

func (r *TaxRegistrationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Tax `Registration` lets us know that your business is registered to collect tax on payments within a region, enabling you to [automatically collect tax](https://docs.stripe.com/tax).\n\nStripe doesn't register on your behalf with the relevant authorities when you create a Tax `Registration` object. For more information on how to register to collect tax, see [our guide](https://docs.stripe.com/tax/registering).\n\nRelated guide: [Using the Registrations API](https://docs.stripe.com/tax/registrations-api)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("tax.registration")},
			},
			"active_from": schema.Int64Attribute{
				Required:    true,
				Description: "Time at which the registration becomes active. Measured in seconds since the Unix epoch.",
			},
			"country": schema.StringAttribute{
				Required:      true,
				Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"country_options": schema.SingleNestedAttribute{
				Required: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"ae": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Place of supply scheme used in an Default standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
						},
					},
					"al": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"am": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"ao": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"at": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"au": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Place of supply scheme used in an Default standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
						},
					},
					"aw": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"az": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"ba": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"bb": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"bd": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"be": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"bf": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"bg": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"bh": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"bj": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"bs": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"by": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"ca": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"province_standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"province": schema.StringAttribute{
										Required:      true,
										Description:   "Two-letter CA province code ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in Canada.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("province_standard", "simplified", "standard")},
							},
						},
					},
					"cd": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"ch": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Place of supply scheme used in an Default standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
						},
					},
					"cl": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"cm": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"co": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"cr": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"cv": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"cy": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"cz": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"de": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"dk": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"ec": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"ee": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"eg": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"es": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"et": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"fi": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"fr": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"gb": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Place of supply scheme used in an Default standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
						},
					},
					"ge": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"gn": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"gr": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"hr": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"hu": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"id": schema.SingleNestedAttribute{
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"ie": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"in": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"is": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"it": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"jp": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Place of supply scheme used in an Default standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
						},
					},
					"ke": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"kg": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"kh": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"kr": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"kz": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"la": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"lk": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"lt": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"lu": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"lv": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"ma": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"md": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"me": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"mk": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"mr": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"mt": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"mx": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"my": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"ng": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"nl": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"no": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Place of supply scheme used in an Default standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
						},
					},
					"np": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"nz": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Place of supply scheme used in an Default standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
						},
					},
					"om": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"pe": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"ph": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"pl": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"pt": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"ro": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"rs": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"ru": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"sa": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"se": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"sg": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Computed:      true,
										Description:   "Place of supply scheme used in an Default standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
						},
					},
					"si": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"sk": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"standard": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Required:      true,
										Description:   "Place of supply scheme used in an EU standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
										Validators:    []validator.String{stringvalidator.OneOf("inbound_goods", "small_seller", "standard")},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in an EU country.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("ioss", "oss_non_union", "oss_union", "standard")},
							},
						},
					},
					"sn": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"sr": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"th": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"tj": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"tr": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"tw": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"tz": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"ua": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"ug": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"us": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"local_amusement_tax": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"jurisdiction": schema.StringAttribute{
										Required:      true,
										Description:   "A [FIPS code](https://www.census.gov/library/reference/code-lists/ansi.html) representing the local jurisdiction.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
							"local_lease_tax": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"jurisdiction": schema.StringAttribute{
										Required:      true,
										Description:   "A [FIPS code](https://www.census.gov/library/reference/code-lists/ansi.html) representing the local jurisdiction.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
							"state": schema.StringAttribute{
								Required:      true,
								Description:   "Two-letter US state code ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
							},
							"state_sales_tax": schema.SingleNestedAttribute{
								Optional: true,
								Computed: true,

								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"elections": schema.ListNestedAttribute{
										Required:      true,
										Description:   "Elections for the state sales tax registration.",
										PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"jurisdiction": schema.StringAttribute{
													Optional:      true,
													Computed:      true,
													Description:   "A [FIPS code](https://www.census.gov/library/reference/code-lists/ansi.html) representing the local jurisdiction.",
													PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
												},
												"type": schema.StringAttribute{
													Required:      true,
													Description:   "The type of the election for the state sales tax registration.",
													PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
													Validators:    []validator.String{stringvalidator.OneOf("local_use_tax", "simplified_sellers_use_tax", "single_local_use_tax")},
												},
											},
										},
									},
								},
							},
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in the US.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("local_amusement_tax", "local_lease_tax", "state_communications_tax", "state_retail_delivery_fee", "state_sales_tax")},
							},
						},
					},
					"uy": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"uz": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"vn": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"za": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
							},
						},
					},
					"zm": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("simplified")},
							},
						},
					},
					"zw": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Required:      true,
								Description:   "Type of registration in `country`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
								Validators:    []validator.String{stringvalidator.OneOf("standard")},
							},
							"standard": schema.SingleNestedAttribute{
								Optional:      true,
								Description:   "Options for the standard registration.",
								PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
								Attributes: map[string]schema.Attribute{
									"place_of_supply_scheme": schema.StringAttribute{
										Optional:      true,
										Description:   "Place of supply scheme used in an standard registration.",
										PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
									},
								},
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
			"expires_at": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "If set, the registration stops being active at this time. If not set, the registration will be active indefinitely. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
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
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "The status of the registration. This field is present for convenience and can be deduced from `active_from` and `expires_at`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("active", "expired", "scheduled")},
			},
		},
	}
}

func (r *TaxRegistrationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TaxRegistrationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config TaxRegistrationResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTaxRegistrationCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building TaxRegistration create params", err.Error())
		return
	}

	obj, err := r.client.V1TaxRegistrations.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating TaxRegistration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TaxRegistrations.B, r.client.V1TaxRegistrations.Key, stripe.FormatURLPath("/v1/tax/registrations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TaxRegistration create raw response", err.Error())
		return
	}

	if err := flattenTaxRegistration(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxRegistration create response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"CountryOptions", "al", "standard"}, []string{"CountryOptions", "al", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "ao", "standard"}, []string{"CountryOptions", "ao", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "aw", "standard"}, []string{"CountryOptions", "aw", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "ba", "standard"}, []string{"CountryOptions", "ba", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bb", "standard"}, []string{"CountryOptions", "bb", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bd", "standard"}, []string{"CountryOptions", "bd", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bf", "standard"}, []string{"CountryOptions", "bf", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bh", "standard"}, []string{"CountryOptions", "bh", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bs", "standard"}, []string{"CountryOptions", "bs", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "cd", "standard"}, []string{"CountryOptions", "cd", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "et", "standard"}, []string{"CountryOptions", "et", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "gn", "standard"}, []string{"CountryOptions", "gn", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "is", "standard"}, []string{"CountryOptions", "is", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "me", "standard"}, []string{"CountryOptions", "me", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "mk", "standard"}, []string{"CountryOptions", "mk", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "mr", "standard"}, []string{"CountryOptions", "mr", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "om", "standard"}, []string{"CountryOptions", "om", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "rs", "standard"}, []string{"CountryOptions", "rs", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "sr", "standard"}, []string{"CountryOptions", "sr", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "uy", "standard"}, []string{"CountryOptions", "uy", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "za", "standard"}, []string{"CountryOptions", "za", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "zw", "standard"}, []string{"CountryOptions", "zw", "standard", "place_of_supply_scheme"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TaxRegistrationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState TaxRegistrationResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TaxRegistrationResourceModel
	state = priorState

	obj, err := r.client.V1TaxRegistrations.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading TaxRegistration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TaxRegistrations.B, r.client.V1TaxRegistrations.Key, stripe.FormatURLPath("/v1/tax/registrations/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TaxRegistration raw response", err.Error())
		return
	}

	if err := flattenTaxRegistration(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxRegistration read response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&state, &priorState, [][]string{[]string{"CountryOptions", "al", "standard"}, []string{"CountryOptions", "al", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "ao", "standard"}, []string{"CountryOptions", "ao", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "aw", "standard"}, []string{"CountryOptions", "aw", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "ba", "standard"}, []string{"CountryOptions", "ba", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bb", "standard"}, []string{"CountryOptions", "bb", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bd", "standard"}, []string{"CountryOptions", "bd", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bf", "standard"}, []string{"CountryOptions", "bf", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bh", "standard"}, []string{"CountryOptions", "bh", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bs", "standard"}, []string{"CountryOptions", "bs", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "cd", "standard"}, []string{"CountryOptions", "cd", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "et", "standard"}, []string{"CountryOptions", "et", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "gn", "standard"}, []string{"CountryOptions", "gn", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "is", "standard"}, []string{"CountryOptions", "is", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "me", "standard"}, []string{"CountryOptions", "me", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "mk", "standard"}, []string{"CountryOptions", "mk", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "mr", "standard"}, []string{"CountryOptions", "mr", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "om", "standard"}, []string{"CountryOptions", "om", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "rs", "standard"}, []string{"CountryOptions", "rs", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "sr", "standard"}, []string{"CountryOptions", "sr", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "uy", "standard"}, []string{"CountryOptions", "uy", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "za", "standard"}, []string{"CountryOptions", "za", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "zw", "standard"}, []string{"CountryOptions", "zw", "standard", "place_of_supply_scheme"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *TaxRegistrationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TaxRegistrationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config TaxRegistrationResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TaxRegistrationResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandTaxRegistrationUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building TaxRegistration update params", err.Error())
		return
	}

	obj, err := r.client.V1TaxRegistrations.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating TaxRegistration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TaxRegistrations.B, r.client.V1TaxRegistrations.Key, stripe.FormatURLPath("/v1/tax/registrations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TaxRegistration update raw response", err.Error())
		return
	}

	if err := flattenTaxRegistration(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TaxRegistration update response", err.Error())
		return
	}
	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"CountryOptions", "al", "standard"}, []string{"CountryOptions", "al", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "ao", "standard"}, []string{"CountryOptions", "ao", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "aw", "standard"}, []string{"CountryOptions", "aw", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "ba", "standard"}, []string{"CountryOptions", "ba", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bb", "standard"}, []string{"CountryOptions", "bb", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bd", "standard"}, []string{"CountryOptions", "bd", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bf", "standard"}, []string{"CountryOptions", "bf", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bh", "standard"}, []string{"CountryOptions", "bh", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "bs", "standard"}, []string{"CountryOptions", "bs", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "cd", "standard"}, []string{"CountryOptions", "cd", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "et", "standard"}, []string{"CountryOptions", "et", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "gn", "standard"}, []string{"CountryOptions", "gn", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "is", "standard"}, []string{"CountryOptions", "is", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "me", "standard"}, []string{"CountryOptions", "me", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "mk", "standard"}, []string{"CountryOptions", "mk", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "mr", "standard"}, []string{"CountryOptions", "mr", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "om", "standard"}, []string{"CountryOptions", "om", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "rs", "standard"}, []string{"CountryOptions", "rs", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "sr", "standard"}, []string{"CountryOptions", "sr", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "uy", "standard"}, []string{"CountryOptions", "uy", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "za", "standard"}, []string{"CountryOptions", "za", "standard", "place_of_supply_scheme"}, []string{"CountryOptions", "zw", "standard"}, []string{"CountryOptions", "zw", "standard", "place_of_supply_scheme"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TaxRegistrationResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *TaxRegistrationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandTaxRegistrationCreate(plan TaxRegistrationResourceModel) (*stripe.TaxRegistrationCreateParams, error) {
	params := &stripe.TaxRegistrationCreateParams{}

	if !plan.ActiveFrom.IsNull() && !plan.ActiveFrom.IsUnknown() {
		params.ActiveFrom = stripe.Int64(plan.ActiveFrom.ValueInt64())
	}
	if !plan.Country.IsNull() && !plan.Country.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Country", "Country", plan.Country.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "country", params)
		}
	}
	if !plan.CountryOptions.IsNull() && !plan.CountryOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CountryOptions", plan.CountryOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "country_options", params)
		}
	}
	if !plan.ExpiresAt.IsNull() && !plan.ExpiresAt.IsUnknown() {
		params.ExpiresAt = stripe.Int64(plan.ExpiresAt.ValueInt64())
	}

	return params, nil
}

func expandTaxRegistrationUpdate(plan TaxRegistrationResourceModel, state TaxRegistrationResourceModel) (*stripe.TaxRegistrationUpdateParams, error) {
	params := &stripe.TaxRegistrationUpdateParams{}

	if !plan.ActiveFrom.Equal(state.ActiveFrom) && !plan.ActiveFrom.IsNull() && !plan.ActiveFrom.IsUnknown() {
		params.ActiveFrom = stripe.Int64(plan.ActiveFrom.ValueInt64())
	}
	if !plan.ExpiresAt.Equal(state.ExpiresAt) && !plan.ExpiresAt.IsNull() && !plan.ExpiresAt.IsUnknown() {
		params.ExpiresAt = stripe.Int64(plan.ExpiresAt.ValueInt64())
	}

	return params, nil
}

func flattenTaxRegistration(obj *stripe.TaxRegistration, state *TaxRegistrationResourceModel) error {
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
		if rawValueActiveFrom, rawOk := plainValueAtPath(raw, "active_from"); rawOk {
			if valueActiveFrom, err := flattenPlainValue(rawValueActiveFrom, types.Int64Type, "active_from", "raw response"); err != nil {
				return err
			} else {
				if typedActiveFrom, ok := valueActiveFrom.(types.Int64); ok {
					state.ActiveFrom = typedActiveFrom
				}
			}
		} else if !hasRaw {
			if responseValueActiveFrom, ok := plainFromResponseField(obj, "ActiveFrom"); ok {
				if valueActiveFrom, err := flattenPlainValue(responseValueActiveFrom, types.Int64Type, "active_from", "response struct"); err != nil {
					return err
				} else {
					if typedActiveFrom, ok := valueActiveFrom.(types.Int64); ok {
						state.ActiveFrom = typedActiveFrom
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
		assignedCountryOptions := false
		hadRawCountryOptions := false
		if rawValueCountryOptions, rawOk := plainValueAtPath(raw, "country_options"); rawOk {
			hadRawCountryOptions = true
			if rawValueCountryOptions != nil {
				sourceCountryOptions := applyConfiguredKeyedListShapes(rawValueCountryOptions, attrValueToPlain(state.CountryOptions))
				if valueCountryOptions, err := flattenPlainValue(sourceCountryOptions, types.ObjectType{AttrTypes: map[string]attr.Type{"ae": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "al": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "am": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ao": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "at": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "au": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "aw": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "az": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ba": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bb": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bd": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "be": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "bf": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bg": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "bh": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bj": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "bs": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "by": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ca": types.ObjectType{AttrTypes: map[string]attr.Type{"province_standard": types.ObjectType{AttrTypes: map[string]attr.Type{"province": types.StringType}}, "type": types.StringType}}, "cd": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "ch": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "cl": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cm": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "co": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cv": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cy": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "cz": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "de": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "dk": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ec": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ee": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "eg": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "es": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "et": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "fi": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "fr": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "gb": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ge": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "gn": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "gr": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "hr": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "hu": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "id": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ie": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "in": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "is": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "it": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "jp": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ke": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kg": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kh": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kz": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "la": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "lk": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "lt": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "lu": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "lv": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ma": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "md": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "me": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "mk": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "mr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "mt": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "mx": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "my": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ng": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "nl": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "no": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "np": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "nz": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "om": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "pe": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ph": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "pl": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "pt": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ro": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "rs": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "ru": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "sa": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "se": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "sg": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "si": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "sk": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "sn": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "sr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "th": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tj": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tw": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tz": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ua": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ug": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "us": types.ObjectType{AttrTypes: map[string]attr.Type{"local_amusement_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"jurisdiction": types.StringType}}, "local_lease_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"jurisdiction": types.StringType}}, "state": types.StringType, "state_sales_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"elections": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"jurisdiction": types.StringType, "type": types.StringType}}}}}, "type": types.StringType}}, "uy": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "uz": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "vn": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "za": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "zm": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "zw": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}}}, "country_options", "raw response"); err != nil {
					return err
				} else {
					if typedCountryOptions, ok := valueCountryOptions.(types.Object); ok {
						state.CountryOptions = typedCountryOptions
						assignedCountryOptions = true
					}
				}
			}
		}
		if !assignedCountryOptions {
			if !hasRaw {
				if responseValueCountryOptions, ok := plainFromResponseField(obj, "CountryOptions"); ok {
					sourceCountryOptions := applyConfiguredKeyedListShapes(responseValueCountryOptions, attrValueToPlain(state.CountryOptions))
					if valueCountryOptions, err := flattenPlainValue(
						sourceCountryOptions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"ae": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "al": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "am": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ao": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "at": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "au": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "aw": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "az": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ba": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bb": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bd": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "be": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "bf": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bg": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "bh": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bj": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "bs": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "by": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ca": types.ObjectType{AttrTypes: map[string]attr.Type{"province_standard": types.ObjectType{AttrTypes: map[string]attr.Type{"province": types.StringType}}, "type": types.StringType}}, "cd": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "ch": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "cl": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cm": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "co": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cv": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cy": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "cz": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "de": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "dk": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ec": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ee": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "eg": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "es": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "et": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "fi": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "fr": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "gb": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ge": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "gn": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "gr": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "hr": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "hu": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "id": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ie": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "in": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "is": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "it": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "jp": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ke": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kg": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kh": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kz": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "la": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "lk": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "lt": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "lu": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "lv": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ma": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "md": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "me": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "mk": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "mr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "mt": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "mx": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "my": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ng": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "nl": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "no": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "np": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "nz": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "om": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "pe": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ph": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "pl": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "pt": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ro": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "rs": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "ru": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "sa": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "se": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "sg": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "si": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "sk": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "sn": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "sr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "th": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tj": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tw": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tz": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ua": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ug": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "us": types.ObjectType{AttrTypes: map[string]attr.Type{"local_amusement_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"jurisdiction": types.StringType}}, "local_lease_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"jurisdiction": types.StringType}}, "state": types.StringType, "state_sales_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"elections": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"jurisdiction": types.StringType, "type": types.StringType}}}}}, "type": types.StringType}}, "uy": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "uz": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "vn": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "za": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "zm": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "zw": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}}},
						"country_options",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCountryOptions, ok := valueCountryOptions.(types.Object); ok {
							state.CountryOptions = typedCountryOptions
							assignedCountryOptions = true
						}
					}
				}
			}
		}
		if !assignedCountryOptions && hadRawCountryOptions {
			if nullCountryOptions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"ae": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "al": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "am": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ao": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "at": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "au": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "aw": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "az": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ba": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bb": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bd": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "be": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "bf": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bg": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "bh": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "bj": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "bs": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "by": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ca": types.ObjectType{AttrTypes: map[string]attr.Type{"province_standard": types.ObjectType{AttrTypes: map[string]attr.Type{"province": types.StringType}}, "type": types.StringType}}, "cd": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "ch": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "cl": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cm": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "co": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cv": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "cy": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "cz": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "de": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "dk": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ec": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ee": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "eg": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "es": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "et": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "fi": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "fr": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "gb": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ge": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "gn": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "gr": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "hr": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "hu": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "id": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ie": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "in": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "is": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "it": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "jp": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ke": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kg": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kh": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "kz": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "la": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "lk": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "lt": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "lu": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "lv": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ma": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "md": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "me": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "mk": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "mr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "mt": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "mx": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "my": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ng": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "nl": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "no": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "np": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "nz": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "om": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "pe": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ph": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "pl": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "pt": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "ro": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "rs": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "ru": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "sa": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "se": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "sg": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "si": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "sk": types.ObjectType{AttrTypes: map[string]attr.Type{"standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}, "type": types.StringType}}, "sn": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "sr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "th": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tj": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tr": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tw": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "tz": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ua": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "ug": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "us": types.ObjectType{AttrTypes: map[string]attr.Type{"local_amusement_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"jurisdiction": types.StringType}}, "local_lease_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"jurisdiction": types.StringType}}, "state": types.StringType, "state_sales_tax": types.ObjectType{AttrTypes: map[string]attr.Type{"elections": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"jurisdiction": types.StringType, "type": types.StringType}}}}}, "type": types.StringType}}, "uy": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "uz": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "vn": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "za": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}, "zm": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType}}, "zw": types.ObjectType{AttrTypes: map[string]attr.Type{"type": types.StringType, "standard": types.ObjectType{AttrTypes: map[string]attr.Type{"place_of_supply_scheme": types.StringType}}}}}}); ok {
				if typedCountryOptions, ok := nullCountryOptions.(types.Object); ok {
					state.CountryOptions = typedCountryOptions
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
	return nil
}
