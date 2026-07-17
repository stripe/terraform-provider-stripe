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

var _ resource.Resource = &TerminalConfigurationResource{}

var _ resource.ResourceWithConfigure = &TerminalConfigurationResource{}

var _ resource.ResourceWithImportState = &TerminalConfigurationResource{}

func NewTerminalConfigurationResource() resource.Resource {
	return &TerminalConfigurationResource{}
}

type TerminalConfigurationResource struct {
	client *stripe.Client
}

type TerminalConfigurationResourceModel struct {
	Object           types.String `tfsdk:"object"`
	BBPOSWisePad3    types.Object `tfsdk:"bbpos_wisepad3"`
	BBPOSWisePOSE    types.Object `tfsdk:"bbpos_wisepos_e"`
	Cellular         types.Object `tfsdk:"cellular"`
	ID               types.String `tfsdk:"id"`
	IsAccountDefault types.Bool   `tfsdk:"is_account_default"`
	Livemode         types.Bool   `tfsdk:"livemode"`
	Name             types.String `tfsdk:"name"`
	Offline          types.Object `tfsdk:"offline"`
	RebootWindow     types.Object `tfsdk:"reboot_window"`
	StripeS700       types.Object `tfsdk:"stripe_s700"`
	StripeS710       types.Object `tfsdk:"stripe_s710"`
	Tipping          types.Object `tfsdk:"tipping"`
	VerifoneM425     types.Object `tfsdk:"verifone_m425"`
	VerifoneP400     types.Object `tfsdk:"verifone_p400"`
	VerifoneP630     types.Object `tfsdk:"verifone_p630"`
	VerifoneUx700    types.Object `tfsdk:"verifone_ux700"`
	VerifoneV660p    types.Object `tfsdk:"verifone_v660p"`
	Wifi             types.Object `tfsdk:"wifi"`
}

func (r *TerminalConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TerminalConfigurationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_terminal_configuration"
}

func (r *TerminalConfigurationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Configurations object represents how features should be configured for terminal readers.\nFor information about how to use it, see the [Terminal configurations documentation](https://docs.stripe.com/terminal/fleet/configurations-overview).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("terminal.configuration")},
			},
			"bbpos_wisepad3": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"splashscreen": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A File ID representing an image to display on the reader",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"bbpos_wisepos_e": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"splashscreen": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A File ID representing an image to display on the reader",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"cellular": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Required:    true,
						Description: "Whether a cellular-capable reader can connect to the internet over cellular.",
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"is_account_default": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether this Configuration is the default for your account",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "String indicating the name of the Configuration object, set by the user",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"offline": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Required:    true,
						Description: "Determines whether to allow transactions to be collected while reader is offline. Defaults to false.",
					},
				},
			},
			"reboot_window": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"end_hour": schema.Int64Attribute{
						Required:    true,
						Description: "Integer between 0 to 23 that represents the end hour of the reboot time window. The value must be different than the start_hour.",
					},
					"start_hour": schema.Int64Attribute{
						Required:    true,
						Description: "Integer between 0 to 23 that represents the start hour of the reboot time window.",
					},
				},
			},
			"stripe_s700": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"splashscreen": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A File ID representing an image to display on the reader",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"stripe_s710": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"splashscreen": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A File ID representing an image to display on the reader",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"tipping": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"aed": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"aud": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"cad": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"chf": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"czk": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"dkk": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"eur": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"gbp": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"gip": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"hkd": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"huf": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"jpy": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"mxn": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"myr": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"nok": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"nzd": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"pln": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"ron": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"sek": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"sgd": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"usd": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"fixed_amounts": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Fixed amounts displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"percentages": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Percentages displayed when collecting a tip",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.Int64Type,
							},
							"smart_tip_threshold": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Below this amount, fixed amounts will be displayed; above it, percentages will be displayed",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"verifone_m425": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"splashscreen": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A File ID representing an image to display on the reader",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"verifone_p400": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"splashscreen": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A File ID representing an image to display on the reader",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"verifone_p630": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"splashscreen": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A File ID representing an image to display on the reader",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"verifone_ux700": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"splashscreen": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A File ID representing an image to display on the reader",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"verifone_v660p": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"splashscreen": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A File ID representing an image to display on the reader",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"wifi": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"enterprise_eap_peap": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"ca_certificate_file": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "A File ID representing a PEM file containing the server certificate",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"password": schema.StringAttribute{
								Required:    true,
								Description: "Password for connecting to the WiFi network",
								Sensitive:   true,
							},
							"ssid": schema.StringAttribute{
								Required:    true,
								Description: "Name of the WiFi network",
							},
							"username": schema.StringAttribute{
								Required:    true,
								Description: "Username for connecting to the WiFi network",
							},
						},
					},
					"enterprise_eap_tls": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"ca_certificate_file": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "A File ID representing a PEM file containing the server certificate",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"client_certificate_file": schema.StringAttribute{
								Required:    true,
								Description: "A File ID representing a PEM file containing the client certificate",
							},
							"private_key_file": schema.StringAttribute{
								Required:    true,
								Description: "A File ID representing a PEM file containing the client RSA private key",
							},
							"private_key_file_password": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Password for the private key file",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Sensitive:     true,
							},
							"ssid": schema.StringAttribute{
								Required:    true,
								Description: "Name of the WiFi network",
							},
						},
					},
					"personal_psk": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"password": schema.StringAttribute{
								Required:    true,
								Description: "Password for connecting to the WiFi network",
								Sensitive:   true,
							},
							"ssid": schema.StringAttribute{
								Required:    true,
								Description: "Name of the WiFi network",
							},
						},
					},
					"type": schema.StringAttribute{
						Required:    true,
						Description: "Security type of the WiFi network. The hash with the corresponding name contains the credentials for this security type.",
						Validators:  []validator.String{stringvalidator.OneOf("enterprise_eap_peap", "enterprise_eap_tls", "personal_psk")},
					},
				},
			},
		},
	}
}

func (r *TerminalConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TerminalConfigurationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTerminalConfigurationCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building TerminalConfiguration create params", err.Error())
		return
	}

	obj, err := r.client.V1TerminalConfigurations.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating TerminalConfiguration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TerminalConfigurations.B, r.client.V1TerminalConfigurations.Key, stripe.FormatURLPath("/v1/terminal/configurations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TerminalConfiguration create raw response", err.Error())
		return
	}

	if err := flattenTerminalConfiguration(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TerminalConfiguration create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TerminalConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState TerminalConfigurationResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TerminalConfigurationResourceModel
	state = priorState

	obj, err := r.client.V1TerminalConfigurations.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading TerminalConfiguration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TerminalConfigurations.B, r.client.V1TerminalConfigurations.Key, stripe.FormatURLPath("/v1/terminal/configurations/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TerminalConfiguration raw response", err.Error())
		return
	}

	if err := flattenTerminalConfiguration(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening TerminalConfiguration read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *TerminalConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TerminalConfigurationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state TerminalConfigurationResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandTerminalConfigurationUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building TerminalConfiguration update params", err.Error())
		return
	}

	obj, err := r.client.V1TerminalConfigurations.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating TerminalConfiguration", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1TerminalConfigurations.B, r.client.V1TerminalConfigurations.Key, stripe.FormatURLPath("/v1/terminal/configurations/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating TerminalConfiguration update raw response", err.Error())
		return
	}

	if err := flattenTerminalConfiguration(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening TerminalConfiguration update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *TerminalConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TerminalConfigurationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.V1TerminalConfigurations.Delete(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting TerminalConfiguration", err.Error())
		return
	}
}

func (r *TerminalConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandTerminalConfigurationCreate(plan TerminalConfigurationResourceModel) (*stripe.TerminalConfigurationCreateParams, error) {
	params := &stripe.TerminalConfigurationCreateParams{}

	if !plan.BBPOSWisePad3.IsNull() && !plan.BBPOSWisePad3.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BBPOSWisePad3", plan.BBPOSWisePad3) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "bbpos_wisepad3", params)
		}
	}
	if !plan.BBPOSWisePOSE.IsNull() && !plan.BBPOSWisePOSE.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BBPOSWisePOSE", plan.BBPOSWisePOSE) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "bbpos_wisepos_e", params)
		}
	}
	if !plan.Cellular.IsNull() && !plan.Cellular.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Cellular", plan.Cellular) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "cellular", params)
		}
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Name", "Name", plan.Name.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "name", params)
		}
	}
	if !plan.Offline.IsNull() && !plan.Offline.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Offline", plan.Offline) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "offline", params)
		}
	}
	if !plan.RebootWindow.IsNull() && !plan.RebootWindow.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RebootWindow", plan.RebootWindow) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "reboot_window", params)
		}
	}
	if !plan.StripeS700.IsNull() && !plan.StripeS700.IsUnknown() {
		if !assignAttrValueToNamedField(params, "StripeS700", plan.StripeS700) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "stripe_s700", params)
		}
	}
	if !plan.StripeS710.IsNull() && !plan.StripeS710.IsUnknown() {
		if !assignAttrValueToNamedField(params, "StripeS710", plan.StripeS710) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "stripe_s710", params)
		}
	}
	if !plan.Tipping.IsNull() && !plan.Tipping.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Tipping", plan.Tipping) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "tipping", params)
		}
	}
	if !plan.VerifoneM425.IsNull() && !plan.VerifoneM425.IsUnknown() {
		if !assignAttrValueToNamedField(params, "VerifoneM425", plan.VerifoneM425) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "verifone_m425", params)
		}
	}
	if !plan.VerifoneP400.IsNull() && !plan.VerifoneP400.IsUnknown() {
		if !assignAttrValueToNamedField(params, "VerifoneP400", plan.VerifoneP400) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "verifone_p400", params)
		}
	}
	if !plan.VerifoneP630.IsNull() && !plan.VerifoneP630.IsUnknown() {
		if !assignAttrValueToNamedField(params, "VerifoneP630", plan.VerifoneP630) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "verifone_p630", params)
		}
	}
	if !plan.VerifoneUx700.IsNull() && !plan.VerifoneUx700.IsUnknown() {
		if !assignAttrValueToNamedField(params, "VerifoneUx700", plan.VerifoneUx700) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "verifone_ux700", params)
		}
	}
	if !plan.VerifoneV660p.IsNull() && !plan.VerifoneV660p.IsUnknown() {
		if !assignAttrValueToNamedField(params, "VerifoneV660p", plan.VerifoneV660p) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "verifone_v660p", params)
		}
	}
	if !plan.Wifi.IsNull() && !plan.Wifi.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Wifi", plan.Wifi) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "wifi", params)
		}
	}

	return params, nil
}

func expandTerminalConfigurationUpdate(plan TerminalConfigurationResourceModel, state TerminalConfigurationResourceModel) (*stripe.TerminalConfigurationUpdateParams, error) {
	params := &stripe.TerminalConfigurationUpdateParams{}

	if !plan.BBPOSWisePad3.Equal(state.BBPOSWisePad3) && !plan.BBPOSWisePad3.IsNull() && !plan.BBPOSWisePad3.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BBPOSWisePad3", plan.BBPOSWisePad3) {
			if !plan.BBPOSWisePad3.Equal(state.BBPOSWisePad3) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "bbpos_wisepad3", params)
			}
		}
	}
	if !plan.BBPOSWisePOSE.Equal(state.BBPOSWisePOSE) && !plan.BBPOSWisePOSE.IsNull() && !plan.BBPOSWisePOSE.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BBPOSWisePOSE", plan.BBPOSWisePOSE) {
			if !plan.BBPOSWisePOSE.Equal(state.BBPOSWisePOSE) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "bbpos_wisepos_e", params)
			}
		}
	}
	if !plan.Cellular.Equal(state.Cellular) && !plan.Cellular.IsNull() && !plan.Cellular.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Cellular", plan.Cellular) {
			if !plan.Cellular.Equal(state.Cellular) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "cellular", params)
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
	if !plan.Offline.Equal(state.Offline) && !plan.Offline.IsNull() && !plan.Offline.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Offline", plan.Offline) {
			if !plan.Offline.Equal(state.Offline) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "offline", params)
			}
		}
	}
	if !plan.RebootWindow.Equal(state.RebootWindow) && !plan.RebootWindow.IsNull() && !plan.RebootWindow.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RebootWindow", plan.RebootWindow) {
			if !plan.RebootWindow.Equal(state.RebootWindow) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "reboot_window", params)
			}
		}
	}
	if !plan.StripeS700.Equal(state.StripeS700) && !plan.StripeS700.IsNull() && !plan.StripeS700.IsUnknown() {
		if !assignAttrValueToNamedField(params, "StripeS700", plan.StripeS700) {
			if !plan.StripeS700.Equal(state.StripeS700) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "stripe_s700", params)
			}
		}
	}
	if !plan.StripeS710.Equal(state.StripeS710) && !plan.StripeS710.IsNull() && !plan.StripeS710.IsUnknown() {
		if !assignAttrValueToNamedField(params, "StripeS710", plan.StripeS710) {
			if !plan.StripeS710.Equal(state.StripeS710) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "stripe_s710", params)
			}
		}
	}
	if !plan.Tipping.Equal(state.Tipping) && !plan.Tipping.IsNull() && !plan.Tipping.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Tipping", plan.Tipping) {
			if !plan.Tipping.Equal(state.Tipping) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "tipping", params)
			}
		}
	}
	if !plan.VerifoneM425.Equal(state.VerifoneM425) && !plan.VerifoneM425.IsNull() && !plan.VerifoneM425.IsUnknown() {
		if !assignAttrValueToNamedField(params, "VerifoneM425", plan.VerifoneM425) {
			if !plan.VerifoneM425.Equal(state.VerifoneM425) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "verifone_m425", params)
			}
		}
	}
	if !plan.VerifoneP400.Equal(state.VerifoneP400) && !plan.VerifoneP400.IsNull() && !plan.VerifoneP400.IsUnknown() {
		if !assignAttrValueToNamedField(params, "VerifoneP400", plan.VerifoneP400) {
			if !plan.VerifoneP400.Equal(state.VerifoneP400) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "verifone_p400", params)
			}
		}
	}
	if !plan.VerifoneP630.Equal(state.VerifoneP630) && !plan.VerifoneP630.IsNull() && !plan.VerifoneP630.IsUnknown() {
		if !assignAttrValueToNamedField(params, "VerifoneP630", plan.VerifoneP630) {
			if !plan.VerifoneP630.Equal(state.VerifoneP630) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "verifone_p630", params)
			}
		}
	}
	if !plan.VerifoneUx700.Equal(state.VerifoneUx700) && !plan.VerifoneUx700.IsNull() && !plan.VerifoneUx700.IsUnknown() {
		if !assignAttrValueToNamedField(params, "VerifoneUx700", plan.VerifoneUx700) {
			if !plan.VerifoneUx700.Equal(state.VerifoneUx700) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "verifone_ux700", params)
			}
		}
	}
	if !plan.VerifoneV660p.Equal(state.VerifoneV660p) && !plan.VerifoneV660p.IsNull() && !plan.VerifoneV660p.IsUnknown() {
		if !assignAttrValueToNamedField(params, "VerifoneV660p", plan.VerifoneV660p) {
			if !plan.VerifoneV660p.Equal(state.VerifoneV660p) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "verifone_v660p", params)
			}
		}
	}
	if !plan.Wifi.Equal(state.Wifi) && !plan.Wifi.IsNull() && !plan.Wifi.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Wifi", plan.Wifi) {
			if !plan.Wifi.Equal(state.Wifi) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "wifi", params)
			}
		}
	}

	return params, nil
}

func flattenTerminalConfiguration(obj *stripe.TerminalConfiguration, state *TerminalConfigurationResourceModel) error {
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
		assignedBBPOSWisePad3 := false
		hadRawBBPOSWisePad3 := false
		if rawValueBBPOSWisePad3, rawOk := plainValueAtPath(raw, "bbpos_wisepad3"); rawOk {
			hadRawBBPOSWisePad3 = true
			if rawValueBBPOSWisePad3 != nil {
				sourceBBPOSWisePad3 := applyConfiguredKeyedListShapes(rawValueBBPOSWisePad3, attrValueToPlain(state.BBPOSWisePad3))
				if valueBBPOSWisePad3, err := flattenPlainValue(sourceBBPOSWisePad3, types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}, "bbpos_wisepad3", "raw response"); err != nil {
					return err
				} else {
					if typedBBPOSWisePad3, ok := valueBBPOSWisePad3.(types.Object); ok {
						state.BBPOSWisePad3 = typedBBPOSWisePad3
						assignedBBPOSWisePad3 = true
					}
				}
			}
		}
		if !assignedBBPOSWisePad3 {
			if !hasRaw {
				if responseValueBBPOSWisePad3, ok := plainFromResponseField(obj, "BBPOSWisePad3"); ok {
					sourceBBPOSWisePad3 := applyConfiguredKeyedListShapes(responseValueBBPOSWisePad3, attrValueToPlain(state.BBPOSWisePad3))
					if valueBBPOSWisePad3, err := flattenPlainValue(
						sourceBBPOSWisePad3,
						types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}},
						"bbpos_wisepad3",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBBPOSWisePad3, ok := valueBBPOSWisePad3.(types.Object); ok {
							state.BBPOSWisePad3 = typedBBPOSWisePad3
							assignedBBPOSWisePad3 = true
						}
					}
				}
			}
		}
		if !assignedBBPOSWisePad3 && hadRawBBPOSWisePad3 {
			if nullBBPOSWisePad3, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}); ok {
				if typedBBPOSWisePad3, ok := nullBBPOSWisePad3.(types.Object); ok {
					state.BBPOSWisePad3 = typedBBPOSWisePad3
				}
			}
		}
	}
	{
		assignedBBPOSWisePOSE := false
		hadRawBBPOSWisePOSE := false
		if rawValueBBPOSWisePOSE, rawOk := plainValueAtPath(raw, "bbpos_wisepos_e"); rawOk {
			hadRawBBPOSWisePOSE = true
			if rawValueBBPOSWisePOSE != nil {
				sourceBBPOSWisePOSE := applyConfiguredKeyedListShapes(rawValueBBPOSWisePOSE, attrValueToPlain(state.BBPOSWisePOSE))
				if valueBBPOSWisePOSE, err := flattenPlainValue(sourceBBPOSWisePOSE, types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}, "bbpos_wisepos_e", "raw response"); err != nil {
					return err
				} else {
					if typedBBPOSWisePOSE, ok := valueBBPOSWisePOSE.(types.Object); ok {
						state.BBPOSWisePOSE = typedBBPOSWisePOSE
						assignedBBPOSWisePOSE = true
					}
				}
			}
		}
		if !assignedBBPOSWisePOSE {
			if !hasRaw {
				if responseValueBBPOSWisePOSE, ok := plainFromResponseField(obj, "BBPOSWisePOSE"); ok {
					sourceBBPOSWisePOSE := applyConfiguredKeyedListShapes(responseValueBBPOSWisePOSE, attrValueToPlain(state.BBPOSWisePOSE))
					if valueBBPOSWisePOSE, err := flattenPlainValue(
						sourceBBPOSWisePOSE,
						types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}},
						"bbpos_wisepos_e",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedBBPOSWisePOSE, ok := valueBBPOSWisePOSE.(types.Object); ok {
							state.BBPOSWisePOSE = typedBBPOSWisePOSE
							assignedBBPOSWisePOSE = true
						}
					}
				}
			}
		}
		if !assignedBBPOSWisePOSE && hadRawBBPOSWisePOSE {
			if nullBBPOSWisePOSE, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}); ok {
				if typedBBPOSWisePOSE, ok := nullBBPOSWisePOSE.(types.Object); ok {
					state.BBPOSWisePOSE = typedBBPOSWisePOSE
				}
			}
		}
	}
	{
		assignedCellular := false
		hadRawCellular := false
		if rawValueCellular, rawOk := plainValueAtPath(raw, "cellular"); rawOk {
			hadRawCellular = true
			if rawValueCellular != nil {
				sourceCellular := applyConfiguredKeyedListShapes(rawValueCellular, attrValueToPlain(state.Cellular))
				if valueCellular, err := flattenPlainValue(sourceCellular, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "cellular", "raw response"); err != nil {
					return err
				} else {
					if typedCellular, ok := valueCellular.(types.Object); ok {
						state.Cellular = typedCellular
						assignedCellular = true
					}
				}
			}
		}
		if !assignedCellular {
			if !hasRaw {
				if responseValueCellular, ok := plainFromResponseField(obj, "Cellular"); ok {
					sourceCellular := applyConfiguredKeyedListShapes(responseValueCellular, attrValueToPlain(state.Cellular))
					if valueCellular, err := flattenPlainValue(
						sourceCellular,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}},
						"cellular",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedCellular, ok := valueCellular.(types.Object); ok {
							state.Cellular = typedCellular
							assignedCellular = true
						}
					}
				}
			}
		}
		if !assignedCellular && hadRawCellular {
			if nullCellular, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}); ok {
				if typedCellular, ok := nullCellular.(types.Object); ok {
					state.Cellular = typedCellular
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
		if rawValueIsAccountDefault, rawOk := plainValueAtPath(raw, "is_account_default"); rawOk {
			if valueIsAccountDefault, err := flattenPlainValue(rawValueIsAccountDefault, types.BoolType, "is_account_default", "raw response"); err != nil {
				return err
			} else {
				if typedIsAccountDefault, ok := valueIsAccountDefault.(types.Bool); ok {
					state.IsAccountDefault = typedIsAccountDefault
				}
			}
		} else if !hasRaw {
			if responseValueIsAccountDefault, ok := plainFromResponseField(obj, "IsAccountDefault"); ok {
				if valueIsAccountDefault, err := flattenPlainValue(responseValueIsAccountDefault, types.BoolType, "is_account_default", "response struct"); err != nil {
					return err
				} else {
					if typedIsAccountDefault, ok := valueIsAccountDefault.(types.Bool); ok {
						state.IsAccountDefault = typedIsAccountDefault
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
	{
		assignedOffline := false
		hadRawOffline := false
		if rawValueOffline, rawOk := plainValueAtPath(raw, "offline"); rawOk {
			hadRawOffline = true
			if rawValueOffline != nil {
				sourceOffline := applyConfiguredKeyedListShapes(rawValueOffline, attrValueToPlain(state.Offline))
				if valueOffline, err := flattenPlainValue(sourceOffline, types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}, "offline", "raw response"); err != nil {
					return err
				} else {
					if typedOffline, ok := valueOffline.(types.Object); ok {
						state.Offline = typedOffline
						assignedOffline = true
					}
				}
			}
		}
		if !assignedOffline {
			if !hasRaw {
				if responseValueOffline, ok := plainFromResponseField(obj, "Offline"); ok {
					sourceOffline := applyConfiguredKeyedListShapes(responseValueOffline, attrValueToPlain(state.Offline))
					if valueOffline, err := flattenPlainValue(
						sourceOffline,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}},
						"offline",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedOffline, ok := valueOffline.(types.Object); ok {
							state.Offline = typedOffline
							assignedOffline = true
						}
					}
				}
			}
		}
		if !assignedOffline && hadRawOffline {
			if nullOffline, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enabled": types.BoolType}}); ok {
				if typedOffline, ok := nullOffline.(types.Object); ok {
					state.Offline = typedOffline
				}
			}
		}
	}
	{
		assignedRebootWindow := false
		hadRawRebootWindow := false
		if rawValueRebootWindow, rawOk := plainValueAtPath(raw, "reboot_window"); rawOk {
			hadRawRebootWindow = true
			if rawValueRebootWindow != nil {
				sourceRebootWindow := applyConfiguredKeyedListShapes(rawValueRebootWindow, attrValueToPlain(state.RebootWindow))
				if valueRebootWindow, err := flattenPlainValue(sourceRebootWindow, types.ObjectType{AttrTypes: map[string]attr.Type{"end_hour": types.Int64Type, "start_hour": types.Int64Type}}, "reboot_window", "raw response"); err != nil {
					return err
				} else {
					if typedRebootWindow, ok := valueRebootWindow.(types.Object); ok {
						state.RebootWindow = typedRebootWindow
						assignedRebootWindow = true
					}
				}
			}
		}
		if !assignedRebootWindow {
			if !hasRaw {
				if responseValueRebootWindow, ok := plainFromResponseField(obj, "RebootWindow"); ok {
					sourceRebootWindow := applyConfiguredKeyedListShapes(responseValueRebootWindow, attrValueToPlain(state.RebootWindow))
					if valueRebootWindow, err := flattenPlainValue(
						sourceRebootWindow,
						types.ObjectType{AttrTypes: map[string]attr.Type{"end_hour": types.Int64Type, "start_hour": types.Int64Type}},
						"reboot_window",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRebootWindow, ok := valueRebootWindow.(types.Object); ok {
							state.RebootWindow = typedRebootWindow
							assignedRebootWindow = true
						}
					}
				}
			}
		}
		if !assignedRebootWindow && hadRawRebootWindow {
			if nullRebootWindow, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"end_hour": types.Int64Type, "start_hour": types.Int64Type}}); ok {
				if typedRebootWindow, ok := nullRebootWindow.(types.Object); ok {
					state.RebootWindow = typedRebootWindow
				}
			}
		}
	}
	{
		assignedStripeS700 := false
		hadRawStripeS700 := false
		if rawValueStripeS700, rawOk := plainValueAtPath(raw, "stripe_s700"); rawOk {
			hadRawStripeS700 = true
			if rawValueStripeS700 != nil {
				sourceStripeS700 := applyConfiguredKeyedListShapes(rawValueStripeS700, attrValueToPlain(state.StripeS700))
				if valueStripeS700, err := flattenPlainValue(sourceStripeS700, types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}, "stripe_s700", "raw response"); err != nil {
					return err
				} else {
					if typedStripeS700, ok := valueStripeS700.(types.Object); ok {
						state.StripeS700 = typedStripeS700
						assignedStripeS700 = true
					}
				}
			}
		}
		if !assignedStripeS700 {
			if !hasRaw {
				if responseValueStripeS700, ok := plainFromResponseField(obj, "StripeS700"); ok {
					sourceStripeS700 := applyConfiguredKeyedListShapes(responseValueStripeS700, attrValueToPlain(state.StripeS700))
					if valueStripeS700, err := flattenPlainValue(
						sourceStripeS700,
						types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}},
						"stripe_s700",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedStripeS700, ok := valueStripeS700.(types.Object); ok {
							state.StripeS700 = typedStripeS700
							assignedStripeS700 = true
						}
					}
				}
			}
		}
		if !assignedStripeS700 && hadRawStripeS700 {
			if nullStripeS700, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}); ok {
				if typedStripeS700, ok := nullStripeS700.(types.Object); ok {
					state.StripeS700 = typedStripeS700
				}
			}
		}
	}
	{
		assignedStripeS710 := false
		hadRawStripeS710 := false
		if rawValueStripeS710, rawOk := plainValueAtPath(raw, "stripe_s710"); rawOk {
			hadRawStripeS710 = true
			if rawValueStripeS710 != nil {
				sourceStripeS710 := applyConfiguredKeyedListShapes(rawValueStripeS710, attrValueToPlain(state.StripeS710))
				if valueStripeS710, err := flattenPlainValue(sourceStripeS710, types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}, "stripe_s710", "raw response"); err != nil {
					return err
				} else {
					if typedStripeS710, ok := valueStripeS710.(types.Object); ok {
						state.StripeS710 = typedStripeS710
						assignedStripeS710 = true
					}
				}
			}
		}
		if !assignedStripeS710 {
			if !hasRaw {
				if responseValueStripeS710, ok := plainFromResponseField(obj, "StripeS710"); ok {
					sourceStripeS710 := applyConfiguredKeyedListShapes(responseValueStripeS710, attrValueToPlain(state.StripeS710))
					if valueStripeS710, err := flattenPlainValue(
						sourceStripeS710,
						types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}},
						"stripe_s710",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedStripeS710, ok := valueStripeS710.(types.Object); ok {
							state.StripeS710 = typedStripeS710
							assignedStripeS710 = true
						}
					}
				}
			}
		}
		if !assignedStripeS710 && hadRawStripeS710 {
			if nullStripeS710, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}); ok {
				if typedStripeS710, ok := nullStripeS710.(types.Object); ok {
					state.StripeS710 = typedStripeS710
				}
			}
		}
	}
	{
		assignedTipping := false
		hadRawTipping := false
		if rawValueTipping, rawOk := plainValueAtPath(raw, "tipping"); rawOk {
			hadRawTipping = true
			if rawValueTipping != nil {
				sourceTipping := applyConfiguredKeyedListShapes(rawValueTipping, attrValueToPlain(state.Tipping))
				if valueTipping, err := flattenPlainValue(sourceTipping, types.ObjectType{AttrTypes: map[string]attr.Type{"aed": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "aud": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "cad": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "chf": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "czk": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "dkk": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "eur": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "gbp": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "gip": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "hkd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "huf": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "jpy": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "mxn": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "myr": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "nok": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "nzd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "pln": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "ron": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "sek": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "sgd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "usd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}}}, "tipping", "raw response"); err != nil {
					return err
				} else {
					if typedTipping, ok := valueTipping.(types.Object); ok {
						state.Tipping = typedTipping
						assignedTipping = true
					}
				}
			}
		}
		if !assignedTipping {
			if !hasRaw {
				if responseValueTipping, ok := plainFromResponseField(obj, "Tipping"); ok {
					sourceTipping := applyConfiguredKeyedListShapes(responseValueTipping, attrValueToPlain(state.Tipping))
					if valueTipping, err := flattenPlainValue(
						sourceTipping,
						types.ObjectType{AttrTypes: map[string]attr.Type{"aed": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "aud": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "cad": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "chf": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "czk": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "dkk": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "eur": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "gbp": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "gip": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "hkd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "huf": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "jpy": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "mxn": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "myr": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "nok": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "nzd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "pln": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "ron": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "sek": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "sgd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "usd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}}},
						"tipping",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTipping, ok := valueTipping.(types.Object); ok {
							state.Tipping = typedTipping
							assignedTipping = true
						}
					}
				}
			}
		}
		if !assignedTipping && hadRawTipping {
			if nullTipping, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"aed": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "aud": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "cad": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "chf": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "czk": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "dkk": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "eur": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "gbp": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "gip": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "hkd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "huf": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "jpy": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "mxn": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "myr": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "nok": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "nzd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "pln": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "ron": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "sek": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "sgd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}, "usd": types.ObjectType{AttrTypes: map[string]attr.Type{"fixed_amounts": types.ListType{ElemType: types.Int64Type}, "percentages": types.ListType{ElemType: types.Int64Type}, "smart_tip_threshold": types.Int64Type}}}}); ok {
				if typedTipping, ok := nullTipping.(types.Object); ok {
					state.Tipping = typedTipping
				}
			}
		}
	}
	{
		assignedVerifoneM425 := false
		hadRawVerifoneM425 := false
		if rawValueVerifoneM425, rawOk := plainValueAtPath(raw, "verifone_m425"); rawOk {
			hadRawVerifoneM425 = true
			if rawValueVerifoneM425 != nil {
				sourceVerifoneM425 := applyConfiguredKeyedListShapes(rawValueVerifoneM425, attrValueToPlain(state.VerifoneM425))
				if valueVerifoneM425, err := flattenPlainValue(sourceVerifoneM425, types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}, "verifone_m425", "raw response"); err != nil {
					return err
				} else {
					if typedVerifoneM425, ok := valueVerifoneM425.(types.Object); ok {
						state.VerifoneM425 = typedVerifoneM425
						assignedVerifoneM425 = true
					}
				}
			}
		}
		if !assignedVerifoneM425 {
			if !hasRaw {
				if responseValueVerifoneM425, ok := plainFromResponseField(obj, "VerifoneM425"); ok {
					sourceVerifoneM425 := applyConfiguredKeyedListShapes(responseValueVerifoneM425, attrValueToPlain(state.VerifoneM425))
					if valueVerifoneM425, err := flattenPlainValue(
						sourceVerifoneM425,
						types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}},
						"verifone_m425",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedVerifoneM425, ok := valueVerifoneM425.(types.Object); ok {
							state.VerifoneM425 = typedVerifoneM425
							assignedVerifoneM425 = true
						}
					}
				}
			}
		}
		if !assignedVerifoneM425 && hadRawVerifoneM425 {
			if nullVerifoneM425, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}); ok {
				if typedVerifoneM425, ok := nullVerifoneM425.(types.Object); ok {
					state.VerifoneM425 = typedVerifoneM425
				}
			}
		}
	}
	{
		assignedVerifoneP400 := false
		hadRawVerifoneP400 := false
		if rawValueVerifoneP400, rawOk := plainValueAtPath(raw, "verifone_p400"); rawOk {
			hadRawVerifoneP400 = true
			if rawValueVerifoneP400 != nil {
				sourceVerifoneP400 := applyConfiguredKeyedListShapes(rawValueVerifoneP400, attrValueToPlain(state.VerifoneP400))
				if valueVerifoneP400, err := flattenPlainValue(sourceVerifoneP400, types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}, "verifone_p400", "raw response"); err != nil {
					return err
				} else {
					if typedVerifoneP400, ok := valueVerifoneP400.(types.Object); ok {
						state.VerifoneP400 = typedVerifoneP400
						assignedVerifoneP400 = true
					}
				}
			}
		}
		if !assignedVerifoneP400 {
			if !hasRaw {
				if responseValueVerifoneP400, ok := plainFromResponseField(obj, "VerifoneP400"); ok {
					sourceVerifoneP400 := applyConfiguredKeyedListShapes(responseValueVerifoneP400, attrValueToPlain(state.VerifoneP400))
					if valueVerifoneP400, err := flattenPlainValue(
						sourceVerifoneP400,
						types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}},
						"verifone_p400",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedVerifoneP400, ok := valueVerifoneP400.(types.Object); ok {
							state.VerifoneP400 = typedVerifoneP400
							assignedVerifoneP400 = true
						}
					}
				}
			}
		}
		if !assignedVerifoneP400 && hadRawVerifoneP400 {
			if nullVerifoneP400, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}); ok {
				if typedVerifoneP400, ok := nullVerifoneP400.(types.Object); ok {
					state.VerifoneP400 = typedVerifoneP400
				}
			}
		}
	}
	{
		assignedVerifoneP630 := false
		hadRawVerifoneP630 := false
		if rawValueVerifoneP630, rawOk := plainValueAtPath(raw, "verifone_p630"); rawOk {
			hadRawVerifoneP630 = true
			if rawValueVerifoneP630 != nil {
				sourceVerifoneP630 := applyConfiguredKeyedListShapes(rawValueVerifoneP630, attrValueToPlain(state.VerifoneP630))
				if valueVerifoneP630, err := flattenPlainValue(sourceVerifoneP630, types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}, "verifone_p630", "raw response"); err != nil {
					return err
				} else {
					if typedVerifoneP630, ok := valueVerifoneP630.(types.Object); ok {
						state.VerifoneP630 = typedVerifoneP630
						assignedVerifoneP630 = true
					}
				}
			}
		}
		if !assignedVerifoneP630 {
			if !hasRaw {
				if responseValueVerifoneP630, ok := plainFromResponseField(obj, "VerifoneP630"); ok {
					sourceVerifoneP630 := applyConfiguredKeyedListShapes(responseValueVerifoneP630, attrValueToPlain(state.VerifoneP630))
					if valueVerifoneP630, err := flattenPlainValue(
						sourceVerifoneP630,
						types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}},
						"verifone_p630",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedVerifoneP630, ok := valueVerifoneP630.(types.Object); ok {
							state.VerifoneP630 = typedVerifoneP630
							assignedVerifoneP630 = true
						}
					}
				}
			}
		}
		if !assignedVerifoneP630 && hadRawVerifoneP630 {
			if nullVerifoneP630, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}); ok {
				if typedVerifoneP630, ok := nullVerifoneP630.(types.Object); ok {
					state.VerifoneP630 = typedVerifoneP630
				}
			}
		}
	}
	{
		assignedVerifoneUx700 := false
		hadRawVerifoneUx700 := false
		if rawValueVerifoneUx700, rawOk := plainValueAtPath(raw, "verifone_ux700"); rawOk {
			hadRawVerifoneUx700 = true
			if rawValueVerifoneUx700 != nil {
				sourceVerifoneUx700 := applyConfiguredKeyedListShapes(rawValueVerifoneUx700, attrValueToPlain(state.VerifoneUx700))
				if valueVerifoneUx700, err := flattenPlainValue(sourceVerifoneUx700, types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}, "verifone_ux700", "raw response"); err != nil {
					return err
				} else {
					if typedVerifoneUx700, ok := valueVerifoneUx700.(types.Object); ok {
						state.VerifoneUx700 = typedVerifoneUx700
						assignedVerifoneUx700 = true
					}
				}
			}
		}
		if !assignedVerifoneUx700 {
			if !hasRaw {
				if responseValueVerifoneUx700, ok := plainFromResponseField(obj, "VerifoneUx700"); ok {
					sourceVerifoneUx700 := applyConfiguredKeyedListShapes(responseValueVerifoneUx700, attrValueToPlain(state.VerifoneUx700))
					if valueVerifoneUx700, err := flattenPlainValue(
						sourceVerifoneUx700,
						types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}},
						"verifone_ux700",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedVerifoneUx700, ok := valueVerifoneUx700.(types.Object); ok {
							state.VerifoneUx700 = typedVerifoneUx700
							assignedVerifoneUx700 = true
						}
					}
				}
			}
		}
		if !assignedVerifoneUx700 && hadRawVerifoneUx700 {
			if nullVerifoneUx700, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}); ok {
				if typedVerifoneUx700, ok := nullVerifoneUx700.(types.Object); ok {
					state.VerifoneUx700 = typedVerifoneUx700
				}
			}
		}
	}
	{
		assignedVerifoneV660p := false
		hadRawVerifoneV660p := false
		if rawValueVerifoneV660p, rawOk := plainValueAtPath(raw, "verifone_v660p"); rawOk {
			hadRawVerifoneV660p = true
			if rawValueVerifoneV660p != nil {
				sourceVerifoneV660p := applyConfiguredKeyedListShapes(rawValueVerifoneV660p, attrValueToPlain(state.VerifoneV660p))
				if valueVerifoneV660p, err := flattenPlainValue(sourceVerifoneV660p, types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}, "verifone_v660p", "raw response"); err != nil {
					return err
				} else {
					if typedVerifoneV660p, ok := valueVerifoneV660p.(types.Object); ok {
						state.VerifoneV660p = typedVerifoneV660p
						assignedVerifoneV660p = true
					}
				}
			}
		}
		if !assignedVerifoneV660p {
			if !hasRaw {
				if responseValueVerifoneV660p, ok := plainFromResponseField(obj, "VerifoneV660p"); ok {
					sourceVerifoneV660p := applyConfiguredKeyedListShapes(responseValueVerifoneV660p, attrValueToPlain(state.VerifoneV660p))
					if valueVerifoneV660p, err := flattenPlainValue(
						sourceVerifoneV660p,
						types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}},
						"verifone_v660p",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedVerifoneV660p, ok := valueVerifoneV660p.(types.Object); ok {
							state.VerifoneV660p = typedVerifoneV660p
							assignedVerifoneV660p = true
						}
					}
				}
			}
		}
		if !assignedVerifoneV660p && hadRawVerifoneV660p {
			if nullVerifoneV660p, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"splashscreen": types.StringType}}); ok {
				if typedVerifoneV660p, ok := nullVerifoneV660p.(types.Object); ok {
					state.VerifoneV660p = typedVerifoneV660p
				}
			}
		}
	}
	{
		assignedWifi := false
		hadRawWifi := false
		if rawValueWifi, rawOk := plainValueAtPath(raw, "wifi"); rawOk {
			hadRawWifi = true
			if rawValueWifi != nil {
				sourceWifi := applyConfiguredKeyedListShapes(rawValueWifi, attrValueToPlain(state.Wifi))
				if valueWifi, err := flattenPlainValue(sourceWifi, types.ObjectType{AttrTypes: map[string]attr.Type{"enterprise_eap_peap": types.ObjectType{AttrTypes: map[string]attr.Type{"ca_certificate_file": types.StringType, "password": types.StringType, "ssid": types.StringType, "username": types.StringType}}, "enterprise_eap_tls": types.ObjectType{AttrTypes: map[string]attr.Type{"ca_certificate_file": types.StringType, "client_certificate_file": types.StringType, "private_key_file": types.StringType, "private_key_file_password": types.StringType, "ssid": types.StringType}}, "personal_psk": types.ObjectType{AttrTypes: map[string]attr.Type{"password": types.StringType, "ssid": types.StringType}}, "type": types.StringType}}, "wifi", "raw response"); err != nil {
					return err
				} else {
					if typedWifi, ok := valueWifi.(types.Object); ok {
						state.Wifi = typedWifi
						assignedWifi = true
					}
				}
			}
		}
		if !assignedWifi {
			if !hasRaw {
				if responseValueWifi, ok := plainFromResponseField(obj, "Wifi"); ok {
					sourceWifi := applyConfiguredKeyedListShapes(responseValueWifi, attrValueToPlain(state.Wifi))
					if valueWifi, err := flattenPlainValue(
						sourceWifi,
						types.ObjectType{AttrTypes: map[string]attr.Type{"enterprise_eap_peap": types.ObjectType{AttrTypes: map[string]attr.Type{"ca_certificate_file": types.StringType, "password": types.StringType, "ssid": types.StringType, "username": types.StringType}}, "enterprise_eap_tls": types.ObjectType{AttrTypes: map[string]attr.Type{"ca_certificate_file": types.StringType, "client_certificate_file": types.StringType, "private_key_file": types.StringType, "private_key_file_password": types.StringType, "ssid": types.StringType}}, "personal_psk": types.ObjectType{AttrTypes: map[string]attr.Type{"password": types.StringType, "ssid": types.StringType}}, "type": types.StringType}},
						"wifi",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedWifi, ok := valueWifi.(types.Object); ok {
							state.Wifi = typedWifi
							assignedWifi = true
						}
					}
				}
			}
		}
		if !assignedWifi && hadRawWifi {
			if nullWifi, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"enterprise_eap_peap": types.ObjectType{AttrTypes: map[string]attr.Type{"ca_certificate_file": types.StringType, "password": types.StringType, "ssid": types.StringType, "username": types.StringType}}, "enterprise_eap_tls": types.ObjectType{AttrTypes: map[string]attr.Type{"ca_certificate_file": types.StringType, "client_certificate_file": types.StringType, "private_key_file": types.StringType, "private_key_file_password": types.StringType, "ssid": types.StringType}}, "personal_psk": types.ObjectType{AttrTypes: map[string]attr.Type{"password": types.StringType, "ssid": types.StringType}}, "type": types.StringType}}); ok {
				if typedWifi, ok := nullWifi.(types.Object); ok {
					state.Wifi = typedWifi
				}
			}
		}
	}
	return nil
}
