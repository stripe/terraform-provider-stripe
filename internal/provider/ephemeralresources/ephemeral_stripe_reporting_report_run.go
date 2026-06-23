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

var _ ephemeral.EphemeralResource = &ReportingReportRunEphemeralResource{}
var _ ephemeral.EphemeralResourceWithConfigure = &ReportingReportRunEphemeralResource{}

func NewReportingReportRunEphemeralResource() ephemeral.EphemeralResource {
	return &ReportingReportRunEphemeralResource{}
}

type ReportingReportRunEphemeralResource struct {
	client *stripe.Client
}

type ReportingReportRunResourceModel struct {
	Object      types.String `tfsdk:"object"`
	Created     types.Int64  `tfsdk:"created"`
	Error       types.String `tfsdk:"error"`
	ID          types.String `tfsdk:"id"`
	Livemode    types.Bool   `tfsdk:"livemode"`
	Parameters  types.Object `tfsdk:"parameters"`
	ReportType  types.String `tfsdk:"report_type"`
	Result      types.String `tfsdk:"result"`
	Status      types.String `tfsdk:"status"`
	SucceededAt types.Int64  `tfsdk:"succeeded_at"`
}

func (r *ReportingReportRunEphemeralResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_reporting_report_run"
}

func (r *ReportingReportRunEphemeralResource) Schema(_ context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = ephemeralSchema.Schema{
		Description: "The Report Run object represents an instance of a report type generated with\nspecific run parameters. Once the object is created, Stripe begins processing the report.\nWhen the report has finished running, it will give you a reference to a file\nwhere you can retrieve your results. For an overview, see\n[API Access to Reports](https://docs.stripe.com/reporting/statements/api).\n\nNote that certain report types can only be run based on your live-mode data (not test-mode\ndata), and will error when queried without a [live-mode API key](https://docs.stripe.com/keys#test-live-modes).",
		Attributes: map[string]ephemeralSchema.Attribute{
			"object": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "String representing the object's type. Objects of the same type share the same value.",
				Validators:  []validator.String{stringvalidator.OneOf("reporting.report_run")},
			},
			"created": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was created. Measured in seconds since the Unix epoch.",
			},
			"error": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "If something should go wrong during the run, a message about the failure (populated when\n `status=failed`).",
			},
			"id": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier for the object.",
			},
			"livemode": ephemeralSchema.BoolAttribute{
				Computed:    true,
				Description: "`true` if the report is run on live mode data and `false` if it is run on test mode data.",
			},
			"parameters": ephemeralSchema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"columns": ephemeralSchema.ListAttribute{
						Optional:    true,
						Computed:    true,
						Description: "The set of output columns requested for inclusion in the report run.",
						ElementType: types.StringType,
					},
					"connected_account": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Connected account ID by which to filter the report run.",
					},
					"currency": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Currency of objects to be included in the report run.",
					},
					"interval_end": ephemeralSchema.Int64Attribute{
						Optional:    true,
						Computed:    true,
						Description: "Ending timestamp of data to be included in the report run. Can be any UTC timestamp between 1 second after the user specified `interval_start` and 1 second before this report's last `data_available_end` value.",
					},
					"interval_start": ephemeralSchema.Int64Attribute{
						Optional:    true,
						Computed:    true,
						Description: "Starting timestamp of data to be included in the report run. Can be any UTC timestamp between 1 second after this report's `data_available_start` and 1 second before the user specified `interval_end` value.",
					},
					"payout": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Payout ID by which to filter the report run.",
					},
					"reporting_category": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Category of balance transactions to be included in the report run.",
					},
					"timezone": ephemeralSchema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Defaults to `Etc/UTC`. The output timezone for all timestamps in the report. A list of possible time zone values is maintained at the [IANA Time Zone Database](http://www.iana.org/time-zones). Has no effect on `interval_start` or `interval_end`.",
					},
				},
			},
			"report_type": ephemeralSchema.StringAttribute{
				Required:    true,
				Description: "The ID of the [report type](https://docs.stripe.com/reports/report-types) to run, such as `\"balance.summary.1\"`.",
			},
			"result": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The file object representing the result of the report run (populated when\n `status=succeeded`).",
			},
			"status": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "Status of this report run. This will be `pending` when the run is initially created.\n When the run finishes, this will be set to `succeeded` and the `result` field will be populated.\n Rarely, we may encounter an error, at which point this will be set to `failed` and the `error` field will be populated.",
			},
			"succeeded_at": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "Timestamp at which this run successfully finished (populated when\n `status=succeeded`). Measured in seconds since the Unix epoch.",
			},
		},
	}
}

func (r *ReportingReportRunEphemeralResource) Configure(_ context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
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

func expandReportingReportRunCreate(plan ReportingReportRunResourceModel) (*stripe.ReportingReportRunCreateParams, error) {
	params := &stripe.ReportingReportRunCreateParams{}

	if !plan.Parameters.IsNull() && !plan.Parameters.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Parameters", plan.Parameters) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "parameters", params)
		}
	}
	if !plan.ReportType.IsNull() && !plan.ReportType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReportType", "ReportType", plan.ReportType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "report_type", params)
		}
	}

	return params, nil
}

func flattenReportingReportRun(obj *stripe.ReportingReportRun, state *ReportingReportRunResourceModel) error {
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
		if rawValueError, rawOk := plainValueAtPath(raw, "error"); rawOk {
			if valueError, err := flattenPlainValue(rawValueError, types.StringType, "error", "raw response"); err != nil {
				return err
			} else {
				if typedError, ok := valueError.(types.String); ok {
					state.Error = typedError
				}
			}
		} else if !hasRaw {
			if responseValueError, ok := plainFromResponseField(obj, "Error"); ok {
				if valueError, err := flattenPlainValue(responseValueError, types.StringType, "error", "response struct"); err != nil {
					return err
				} else {
					if typedError, ok := valueError.(types.String); ok {
						state.Error = typedError
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
		assignedParameters := false
		hadRawParameters := false
		if rawValueParameters, rawOk := plainValueAtPath(raw, "parameters"); rawOk {
			hadRawParameters = true
			if rawValueParameters != nil {
				sourceParameters := applyConfiguredKeyedListShapes(rawValueParameters, attrValueToPlain(state.Parameters))
				if valueParameters, err := flattenPlainValue(sourceParameters, types.ObjectType{AttrTypes: map[string]attr.Type{"columns": types.ListType{ElemType: types.StringType}, "connected_account": types.StringType, "currency": types.StringType, "interval_end": types.Int64Type, "interval_start": types.Int64Type, "payout": types.StringType, "reporting_category": types.StringType, "timezone": types.StringType}}, "parameters", "raw response"); err != nil {
					return err
				} else {
					if typedParameters, ok := valueParameters.(types.Object); ok {
						state.Parameters = typedParameters
						assignedParameters = true
					}
				}
			}
		}
		if !assignedParameters {
			if !hasRaw {
				if responseValueParameters, ok := plainFromResponseField(obj, "Parameters"); ok {
					sourceParameters := applyConfiguredKeyedListShapes(responseValueParameters, attrValueToPlain(state.Parameters))
					if valueParameters, err := flattenPlainValue(
						sourceParameters,
						types.ObjectType{AttrTypes: map[string]attr.Type{"columns": types.ListType{ElemType: types.StringType}, "connected_account": types.StringType, "currency": types.StringType, "interval_end": types.Int64Type, "interval_start": types.Int64Type, "payout": types.StringType, "reporting_category": types.StringType, "timezone": types.StringType}},
						"parameters",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedParameters, ok := valueParameters.(types.Object); ok {
							state.Parameters = typedParameters
							assignedParameters = true
						}
					}
				}
			}
		}
		if !assignedParameters && hadRawParameters {
			if nullParameters, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"columns": types.ListType{ElemType: types.StringType}, "connected_account": types.StringType, "currency": types.StringType, "interval_end": types.Int64Type, "interval_start": types.Int64Type, "payout": types.StringType, "reporting_category": types.StringType, "timezone": types.StringType}}); ok {
				if typedParameters, ok := nullParameters.(types.Object); ok {
					state.Parameters = typedParameters
				}
			}
		}
	}
	{
		if rawValueReportType, rawOk := plainValueAtPath(raw, "report_type"); rawOk {
			if valueReportType, err := flattenPlainValue(rawValueReportType, types.StringType, "report_type", "raw response"); err != nil {
				return err
			} else {
				if typedReportType, ok := valueReportType.(types.String); ok {
					state.ReportType = typedReportType
				}
			}
		} else if !hasRaw {
			if responseValueReportType, ok := plainFromResponseField(obj, "ReportType"); ok {
				if valueReportType, err := flattenPlainValue(responseValueReportType, types.StringType, "report_type", "response struct"); err != nil {
					return err
				} else {
					if typedReportType, ok := valueReportType.(types.String); ok {
						state.ReportType = typedReportType
					}
				}
			}
		}
	}
	{
		if true {
			if rawValueResult, rawOk := plainValueAtPath(raw, "result"); rawOk {
				if typedResult, ok := plainToStringIDValue(rawValueResult); ok {
					state.Result = typedResult
				}
			} else if !hasRaw {
				if responseValueResult, ok := plainFromResponseField(obj, "Result"); ok {
					if typedResult, ok := plainToStringIDValue(responseValueResult); ok {
						state.Result = typedResult
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
		if rawValueSucceededAt, rawOk := plainValueAtPath(raw, "succeeded_at"); rawOk {
			if valueSucceededAt, err := flattenPlainValue(rawValueSucceededAt, types.Int64Type, "succeeded_at", "raw response"); err != nil {
				return err
			} else {
				if typedSucceededAt, ok := valueSucceededAt.(types.Int64); ok {
					state.SucceededAt = typedSucceededAt
				}
			}
		} else if !hasRaw {
			if responseValueSucceededAt, ok := plainFromResponseField(obj, "SucceededAt"); ok {
				if valueSucceededAt, err := flattenPlainValue(responseValueSucceededAt, types.Int64Type, "succeeded_at", "response struct"); err != nil {
					return err
				} else {
					if typedSucceededAt, ok := valueSucceededAt.(types.Int64); ok {
						state.SucceededAt = typedSucceededAt
					}
				}
			}
		}
	}
	return nil
}

func (r *ReportingReportRunEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var config ReportingReportRunResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandReportingReportRunCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building ReportingReportRun ephemeral params", err.Error())
		return
	}

	obj, err := r.client.V1ReportingReportRuns.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error opening ReportingReportRun ephemeral resource", err.Error())
		return
	}

	result := config
	if err := flattenReportingReportRun(obj, &result); err != nil {
		resp.Diagnostics.AddError("Error flattening ReportingReportRun ephemeral response", err.Error())
		return
	}
	normalizeUnknownValues(&result)
	diags = resp.Result.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
}
