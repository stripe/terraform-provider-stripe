//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeCreditTypeResource struct{ transport bindingruntime.Transport }

type MetronomeCreditTypeResourceModel struct {
	ID         types.String `tfsdk:"id"`
	IsCurrency types.Bool   `tfsdk:"is_currency"`
	Name       types.String `tfsdk:"name"`
}

func expandMetronomeCreditTypeResourceModel(ctx context.Context, model *MetronomeCreditTypeResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"id": model.ID.Type(ctx), "is_currency": model.IsCurrency.Type(ctx), "name": model.Name.Type(ctx)}, map[string]attr.Value{"id": model.ID, "is_currency": model.IsCurrency, "name": model.Name}))
}

func NewMetronomeCreditTypeResource() resource.Resource { return &MetronomeCreditTypeResource{} }

func (r *MetronomeCreditTypeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_credit_type"
}

func (r *MetronomeCreditTypeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A custom Metronome credit type. Import by ID. Destroy archives it; built-in currency credit types cannot be archived by this resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"is_currency": schema.BoolAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:      true,
				Description:   "Name of the custom pricing unit shown on invoices.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *MetronomeCreditTypeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeCreditTypeResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeCreditTypeResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCreditTypeResourceModel(ctx, &valuesModel)
	if err := validateCreditType(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeCreditTypeResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}
	var valuesModel MetronomeCreditTypeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCreditTypeResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCreditTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCreditTypeResourceModel(ctx, &priorModel)
	replacement, err := planCreditType(ctx, values, prior)
	if err != nil {
		resp.Diagnostics.AddError("Unsafe Metronome change", err.Error())
		return
	}
	for _, attribute := range replacement {
		resp.RequiresReplace = append(resp.RequiresReplace, path.Root(attribute))
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, values.Object)...)
}

func (r *MetronomeCreditTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeCreditTypeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCreditTypeResourceModel(ctx, &valuesModel)
	if err := validateCreditType(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
	if !values.Known("is_currency") {
		if err := values.SetJSON(ctx, "is_currency", bindingruntime.Literal("false")); err != nil {
			resp.Diagnostics.AddError("Invalid state default", err.Error())
			return
		}
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
		Operation: bindingruntime.Operation{ID: "createCreditType-v1", Service: "metronome", Method: "POST", Path: "/v1/credit-types/create", Read: false},
		Request:   []bindingruntime.Mapping{{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{{Attribute: "id", Path: []string{"id"}, Required: true, PreserveMissing: false, PreserveConfigured: false, EmptyStringIsNull: false}},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
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

func (r *MetronomeCreditTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeCreditTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCreditTypeResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "listCreditTypes-v1", Service: "metronome", Method: "GET", Path: "/v1/credit-types/list", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "", Constant: json.RawMessage("\"100\""), Path: []string{"limit"}, Location: "query", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{"id": "id"},
		Paginate: true, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readCreditType, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	}}, values, nil, nil); err != nil {
		if bindingruntime.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Could not refresh Metronome resource", err.Error())
		return
	}
	if !values.Known("is_currency") {
		if err := values.SetJSON(ctx, "is_currency", bindingruntime.Literal("false")); err != nil {
			resp.Diagnostics.AddError("Invalid state default", err.Error())
			return
		}
	}
	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeCreditTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeCreditTypeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCreditTypeResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCreditTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCreditTypeResourceModel(ctx, &priorModel)
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
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{}, values, prior, persist); err != nil {
		resp.Diagnostics.AddError("Could not update Metronome resource", err.Error())
		return
	}
	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeCreditTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomeCreditTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCreditTypeResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "listCreditTypes-v1", Service: "metronome", Method: "GET", Path: "/v1/credit-types/list", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "", Constant: json.RawMessage("\"100\""), Path: []string{"limit"}, Location: "query", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{"id": "id"},
		Paginate: true, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readCreditType, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "archiveCreditType-v1", Service: "metronome", Method: "POST", Path: "/v1/credit-types/archive", Read: false},
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

func (r *MetronomeCreditTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var definition resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &definition)
	values := bindingruntime.EmptyValues(definition.Schema.Type().(types.ObjectType))
	if err := bindingruntime.ImportIdentity(ctx, values, req.ID, []string{"id"}, "/"); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome import", err.Error())
		return
	}
	if !values.Known("is_currency") {
		if err := values.SetJSON(ctx, "is_currency", bindingruntime.Literal("false")); err != nil {
			resp.Diagnostics.AddError("Invalid state default", err.Error())
			return
		}
	}
	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeCreditTypeResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_credit_type" {
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
