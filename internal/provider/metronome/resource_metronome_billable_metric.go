//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeBillableMetricResource struct{ transport bindingruntime.Transport }

type MetronomeBillableMetricResourceModel struct {
	AggregationKey  types.String `tfsdk:"aggregation_key"`
	AggregationType types.String `tfsdk:"aggregation_type"`
	EventType       types.String `tfsdk:"event_type"`
	ID              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
}

func expandMetronomeBillableMetricResourceModel(ctx context.Context, model *MetronomeBillableMetricResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"aggregation_key": model.AggregationKey.Type(ctx), "aggregation_type": model.AggregationType.Type(ctx), "event_type": model.EventType.Type(ctx), "id": model.ID.Type(ctx), "name": model.Name.Type(ctx)}, map[string]attr.Value{"aggregation_key": model.AggregationKey, "aggregation_type": model.AggregationType, "event_type": model.EventType, "id": model.ID, "name": model.Name}))
}

func NewMetronomeBillableMetricResource() resource.Resource {
	return &MetronomeBillableMetricResource{}
}

func (r *MetronomeBillableMetricResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_billable_metric"
}

func (r *MetronomeBillableMetricResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "COUNT or SUM billable metric filtered to one event type. Destroy archives the metric.",
		Attributes: map[string]schema.Attribute{
			"aggregation_key": schema.StringAttribute{
				Optional:      true,
				Description:   "Event property to sum when aggregation_type is SUM. The provider requires this property to exist on matched events.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"aggregation_type": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "COUNT events or SUM a numeric event property.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Default:       stringdefault.StaticString("COUNT"),
			},
			"event_type": schema.StringAttribute{
				Required:      true,
				Description:   "Event type matched by this metric.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *MetronomeBillableMetricResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeBillableMetricResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeBillableMetricResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeBillableMetricResourceModel(ctx, &valuesModel)
	if err := validateMetric(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeBillableMetricResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeBillableMetricResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeBillableMetricResourceModel(ctx, &valuesModel)
	if err := validateMetric(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}

	persist := func(ctx context.Context, state *bindingruntime.Values) error {
		if err := state.ResolveUnknowns(); err != nil {
			return err
		}
		resp.Diagnostics.Append(resp.State.Set(ctx, state.Object)...)
		if resp.Diagnostics.HasError() {
			return errors.New("could not persist created identity")
		}
		return nil
	}
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "createBillableMetricV1-v1", Service: "metronome", Method: "POST", Path: "/v1/billable-metrics/create", Read: false},
		Request:   []bindingruntime.Mapping{},
		Response:  []bindingruntime.ResponseField{{Attribute: "id", Path: []string{"id"}, Required: true, PreserveMissing: false, PreserveConfigured: false, EmptyStringIsNull: false}},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: metricRequest, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: true, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	}}, values, nil, persist); err != nil {
		resp.Diagnostics.AddError("Could not create Metronome resource", err.Error())
		return
	}
	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeBillableMetricResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeBillableMetricResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeBillableMetricResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getBillableMetric-v1", Service: "metronome", Method: "GET", Path: "/v1/billable-metrics/{billable_metric_id}", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"billable_metric_id"}, Location: "path", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readMetric, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{[]string{"archived_at"}}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	}}, values, nil, nil); err != nil {
		if bindingruntime.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Could not refresh Metronome resource", err.Error())
		return
	}

	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeBillableMetricResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeBillableMetricResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeBillableMetricResourceModel(ctx, &valuesModel)
	var priorModel MetronomeBillableMetricResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeBillableMetricResourceModel(ctx, &priorModel)
	if err := values.Set("id", prior.Get("id")); err != nil {
		resp.Diagnostics.AddError("Invalid resource identity", err.Error())
		return
	}
	persist := func(ctx context.Context, state *bindingruntime.Values) error {
		resp.Diagnostics.Append(resp.State.Set(ctx, state.Object)...)
		if resp.Diagnostics.HasError() {
			return errors.New("could not persist partial update")
		}
		return nil
	}
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "updateBillableMetric-v1", Service: "metronome", Method: "PUT", Path: "/v1/billable-metrics/{billable_metric_id}", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"billable_metric_id"}, Location: "path", Wrap: false, OmitUnknown: false},
			{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{},
		Root:     []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: checkIdentity, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{"name"}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	}}, values, prior, persist); err != nil {
		resp.Diagnostics.AddError("Could not update Metronome resource", err.Error())
		return
	}
	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeBillableMetricResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomeBillableMetricResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeBillableMetricResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "archiveBillableMetric-v1", Service: "metronome", Method: "POST", Path: "/v1/billable-metrics/archive", Read: false},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"id"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: true, OnError: nil,
	}}, values, nil, nil); err != nil && !bindingruntime.IsMissing(err) {
		resp.Diagnostics.AddError("Could not archive Metronome resource", err.Error())
		return
	}
}

func (r *MetronomeBillableMetricResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var definition resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &definition)
	values := bindingruntime.EmptyValues(definition.Schema.Type().(types.ObjectType))
	if err := bindingruntime.ImportIdentity(ctx, values, req.ID, []string{"id"}, "/"); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome import", err.Error())
		return
	}

	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeBillableMetricResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_billable_metric" {
			return
		}
		switch req.SourceSchemaVersion {
		case 0:
		default:
			resp.Diagnostics.AddError("Unsupported Metronome source state", "The standalone schema version is not in the reviewed migration contract.")
			return
		}
		if req.SourceRawState == nil {
			resp.Diagnostics.AddError("Missing source state", "No legacy resource state was supplied.")
			return
		}
		var definition resource.SchemaResponse
		r.Schema(ctx, resource.SchemaRequest{}, &definition)
		raw, err := req.SourceRawState.Unmarshal(definition.Schema.Type().TerraformType(ctx))
		if err != nil {
			resp.Diagnostics.AddError("Invalid legacy Metronome state", err.Error())
			return
		}
		resp.TargetState = tfsdk.State{Schema: definition.Schema, Raw: raw}
		resp.TargetPrivate = req.SourcePrivate
	}}}
}
