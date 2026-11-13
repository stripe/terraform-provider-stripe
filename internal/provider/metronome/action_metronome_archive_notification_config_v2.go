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

type MetronomeArchiveNotificationConfigV2Action struct{ transport bindingruntime.Transport }

type MetronomeArchiveNotificationConfigV2ActionModel struct {
	Request     types.Dynamic `tfsdk:"request"`
	RequestJson types.String  `tfsdk:"request_json"`
	PathParams  types.Map     `tfsdk:"path_params"`
	QueryParams types.Map     `tfsdk:"query_params"`
}

func expandMetronomeArchiveNotificationConfigV2ActionModel(ctx context.Context, model *MetronomeArchiveNotificationConfigV2ActionModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"request": model.Request.Type(ctx), "request_json": model.RequestJson.Type(ctx), "path_params": model.PathParams.Type(ctx), "query_params": model.QueryParams.Type(ctx)}, map[string]attr.Value{"request": model.Request, "request_json": model.RequestJson, "path_params": model.PathParams, "query_params": model.QueryParams}))
}

func NewMetronomeArchiveNotificationConfigV2Action() action.Action {
	return &MetronomeArchiveNotificationConfigV2Action{}
}

func (r *MetronomeArchiveNotificationConfigV2Action) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_archive_notification_config_v2"
}

func (r *MetronomeArchiveNotificationConfigV2Action) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Invoke Metronome OpenAPI operation archiveNotificationConfig-v2 (POST /v2/notifications/archive). This action runs only when Terraform invokes it; it does not manage a durable object.",
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

func (r *MetronomeArchiveNotificationConfigV2Action) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func (r *MetronomeArchiveNotificationConfigV2Action) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var valuesModel MetronomeArchiveNotificationConfigV2ActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeArchiveNotificationConfigV2ActionModel(ctx, &valuesModel)
	if _, _, err := bindingruntime.InvokeOperation(ctx, r.transport, bindingruntime.Operation{ID: "archiveNotificationConfig-v2", Service: "metronome", Method: "POST", Path: "/v2/notifications/archive", Read: false}, values); err != nil {
		resp.Diagnostics.AddError("Metronome action failed", err.Error())
		return
	}
	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: "Metronome operation archiveNotificationConfig-v2 succeeded"})
	}
}
