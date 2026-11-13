//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeCreateHistoricalContractUsageInvoicesV1Action struct{ transport bindingruntime.Transport }

type MetronomeCreateHistoricalContractUsageInvoicesV1ActionModel struct {
	Request     types.Dynamic `tfsdk:"request"`
	RequestJson types.String  `tfsdk:"request_json"`
	PathParams  types.Map     `tfsdk:"path_params"`
	QueryParams types.Map     `tfsdk:"query_params"`
}

func expandMetronomeCreateHistoricalContractUsageInvoicesV1ActionModel(ctx context.Context, model *MetronomeCreateHistoricalContractUsageInvoicesV1ActionModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"request": model.Request.Type(ctx), "request_json": model.RequestJson.Type(ctx), "path_params": model.PathParams.Type(ctx), "query_params": model.QueryParams.Type(ctx)}, map[string]attr.Value{"request": model.Request, "request_json": model.RequestJson, "path_params": model.PathParams, "query_params": model.QueryParams}))
}

func NewMetronomeCreateHistoricalContractUsageInvoicesV1Action() action.Action {
	return &MetronomeCreateHistoricalContractUsageInvoicesV1Action{}
}

func (r *MetronomeCreateHistoricalContractUsageInvoicesV1Action) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_create_historical_contract_usage_invoices_v1"
}

func (r *MetronomeCreateHistoricalContractUsageInvoicesV1Action) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Invoke Metronome OpenAPI operation createHistoricalContractUsageInvoices-v1 (POST /v1/contracts/createHistoricalInvoices). This action runs only when Terraform invokes it; it does not manage a durable object.",
		Attributes: map[string]schema.Attribute{
			"request": schema.DynamicAttribute{
				Optional:    true,
				Description: "HCL object using Metronome OpenAPI field names. Prefer this over request_json.",
			},
			"request_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON request body using the Metronome OpenAPI field names.",
			},
			"path_params": schema.MapAttribute{
				Optional:    true,
				Description: "Values for named path parameters, such as customer_id.",
				ElementType: types.StringType,
			},
			"query_params": schema.MapAttribute{
				Optional:    true,
				Description: "URL query parameters.",
				ElementType: types.StringType,
			},
		},
	}
}

func (r *MetronomeCreateHistoricalContractUsageInvoicesV1Action) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	data, ok := req.ProviderData.(bindingruntime.Transport)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected the combined provider service clients.")
		return
	}
	r.transport = data
}

func (r *MetronomeCreateHistoricalContractUsageInvoicesV1Action) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var valuesModel MetronomeCreateHistoricalContractUsageInvoicesV1ActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCreateHistoricalContractUsageInvoicesV1ActionModel(ctx, &valuesModel)
	if _, _, err := bindingruntime.InvokeOperation(ctx, r.transport, bindingruntime.Operation{ID: "createHistoricalContractUsageInvoices-v1", Service: "metronome", Method: "POST", Path: "/v1/contracts/createHistoricalInvoices", Read: false}, values); err != nil {
		resp.Diagnostics.AddError("Metronome action failed", err.Error())
		return
	}
	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: "Metronome operation createHistoricalContractUsageInvoices-v1 succeeded"})
	}
}
