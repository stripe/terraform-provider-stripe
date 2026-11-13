//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeListSystemNotificationConfigsV2DataSource struct{ transport bindingruntime.Transport }

type MetronomeListSystemNotificationConfigsV2DataSourceModel struct {
	ID             types.String  `tfsdk:"id"`
	Request        types.Dynamic `tfsdk:"request"`
	RequestJson    types.String  `tfsdk:"request_json"`
	PathParams     types.Map     `tfsdk:"path_params"`
	QueryParams    types.Map     `tfsdk:"query_params"`
	ResponseJson   types.String  `tfsdk:"response_json"`
	ResponseBase64 types.String  `tfsdk:"response_base64"`
}

func expandMetronomeListSystemNotificationConfigsV2DataSourceModel(ctx context.Context, model *MetronomeListSystemNotificationConfigsV2DataSourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"id": model.ID.Type(ctx), "request": model.Request.Type(ctx), "request_json": model.RequestJson.Type(ctx), "path_params": model.PathParams.Type(ctx), "query_params": model.QueryParams.Type(ctx), "response_json": model.ResponseJson.Type(ctx), "response_base64": model.ResponseBase64.Type(ctx)}, map[string]attr.Value{"id": model.ID, "request": model.Request, "request_json": model.RequestJson, "path_params": model.PathParams, "query_params": model.QueryParams, "response_json": model.ResponseJson, "response_base64": model.ResponseBase64}))
}

func NewMetronomeListSystemNotificationConfigsV2DataSource() datasource.DataSource {
	return &MetronomeListSystemNotificationConfigsV2DataSource{}
}

func (r *MetronomeListSystemNotificationConfigsV2DataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_list_system_notification_configs_v2"
}

func (r *MetronomeListSystemNotificationConfigsV2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Metronome OpenAPI operation listSystemNotificationConfigs-v2 (POST /v2/notifications/system/list). Request and response use the API's JSON field names.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"request": schema.DynamicAttribute{
				Optional:    true,
				Description: "HCL object using Metronome OpenAPI field names. Prefer this over request_json.",
			},
			"request_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON request body using the Metronome OpenAPI field names.",
				Sensitive:   true,
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
			"response_json": schema.StringAttribute{
				Computed:    true,
				Description: "Complete JSON response. Sensitive because some API responses contain customer data.",
				Sensitive:   true,
			},
			"response_base64": schema.StringAttribute{
				Computed:    true,
				Description: "Base64 response for binary operations, such as invoice PDF.",
				Sensitive:   true,
			},
		},
	}
}

func (r *MetronomeListSystemNotificationConfigsV2DataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MetronomeListSystemNotificationConfigsV2DataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var valuesModel MetronomeListSystemNotificationConfigsV2DataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeListSystemNotificationConfigsV2DataSourceModel(ctx, &valuesModel)
	if err := bindingruntime.ReadOperation(ctx, r.transport, bindingruntime.Operation{ID: "listSystemNotificationConfigs-v2", Service: "metronome", Method: "POST", Path: "/v2/notifications/system/list", Read: true}, values, false); err != nil {
		resp.Diagnostics.AddError("Metronome read failed", err.Error())
		return
	}
	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}
