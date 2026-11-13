//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeCustomFieldKeyResource struct{ transport bindingruntime.Transport }

type MetronomeCustomFieldKeyResourceModel struct {
	EnforceUniqueness types.Bool   `tfsdk:"enforce_uniqueness"`
	Entity            types.String `tfsdk:"entity"`
	ID                types.String `tfsdk:"id"`
	Key               types.String `tfsdk:"key"`
}

func expandMetronomeCustomFieldKeyResourceModel(ctx context.Context, model *MetronomeCustomFieldKeyResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"enforce_uniqueness": model.EnforceUniqueness.Type(ctx), "entity": model.Entity.Type(ctx), "id": model.ID.Type(ctx), "key": model.Key.Type(ctx)}, map[string]attr.Value{"enforce_uniqueness": model.EnforceUniqueness, "entity": model.Entity, "id": model.ID, "key": model.Key}))
}

func NewMetronomeCustomFieldKeyResource() resource.Resource {
	return &MetronomeCustomFieldKeyResource{}
}

func (r *MetronomeCustomFieldKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_custom_field_key"
}

func (r *MetronomeCustomFieldKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Metronome custom field key. Import as entity/key. Destroy disables the key; existing values are not deleted.",
		Attributes: map[string]schema.Attribute{
			"enforce_uniqueness": schema.BoolAttribute{
				Required: true,

				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"entity": schema.StringAttribute{
				Required:      true,
				Description:   "Metronome ManagedEntity value from the pinned OpenAPI enum.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"key": schema.StringAttribute{
				Required: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *MetronomeCustomFieldKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeCustomFieldKeyResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeCustomFieldKeyResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomFieldKeyResourceModel(ctx, &valuesModel)
	if err := validateCustomFieldKey(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeCustomFieldKeyResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}
	var valuesModel MetronomeCustomFieldKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomFieldKeyResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCustomFieldKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCustomFieldKeyResourceModel(ctx, &priorModel)
	replacement, err := planCustomFieldKey(ctx, values, prior)
	if err != nil {
		resp.Diagnostics.AddError("Unsafe Metronome change", err.Error())
		return
	}
	for _, attribute := range replacement {
		resp.RequiresReplace = append(resp.RequiresReplace, path.Root(attribute))
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, values.Object)...)
}

func (r *MetronomeCustomFieldKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeCustomFieldKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomFieldKeyResourceModel(ctx, &valuesModel)
	if err := validateCustomFieldKey(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}

	if err := bindingruntime.ComposeIdentity(ctx, values, "id", []string{"entity", "key"}, "/"); err != nil {
		resp.Diagnostics.AddError("Invalid resource identity", err.Error())
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
		Operation: bindingruntime.Operation{ID: "listCustomFieldKeys-v1", Service: "metronome", Method: "POST", Path: "/v1/customFields/listKeys", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "entity", Constant: nil, Path: []string{"entities"}, Location: "body", Wrap: true, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: true, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: checkCustomKeyAbsent, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "addCustomFieldKey-v1", Service: "metronome", Method: "POST", Path: "/v1/customFields/addKey", Read: false},
			Request: []bindingruntime.Mapping{{Attribute: "entity", Constant: nil, Path: []string{"entity"}, Location: "body", Wrap: false, OmitUnknown: false},
				{Attribute: "key", Constant: nil, Path: []string{"key"}, Location: "body", Wrap: false, OmitUnknown: false},
				{Attribute: "enforce_uniqueness", Constant: nil, Path: []string{"enforce_uniqueness"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response: []bindingruntime.ResponseField{},
			Root:     []string{"data"}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: true, PersistAttributes: []string{},
			ArchivedPaths: [][]string{}, MissingStatuses: []int{},
			WaitForReady: false, IgnoreResponse: true, OnError: nil,
		},
		{
			Operation: bindingruntime.Operation{ID: "listCustomFieldKeys-v1", Service: "metronome", Method: "POST", Path: "/v1/customFields/listKeys", Read: true},
			Request:   []bindingruntime.Mapping{{Attribute: "entity", Constant: nil, Path: []string{"entities"}, Location: "body", Wrap: true, OmitUnknown: false}},
			Response:  []bindingruntime.ResponseField{},
			Root:      []string{"data"}, Selector: map[string]string{"entity": "entity", "key": "key"},
			Paginate: true, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: checkCreatedCustomKey, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
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

func (r *MetronomeCustomFieldKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeCustomFieldKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomFieldKeyResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "listCustomFieldKeys-v1", Service: "metronome", Method: "POST", Path: "/v1/customFields/listKeys", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "entity", Constant: nil, Path: []string{"entities"}, Location: "body", Wrap: true, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{{Attribute: "entity", Path: []string{"entity"}, Required: false, PreserveMissing: false, PreserveConfigured: false, EmptyStringIsNull: false},
			{Attribute: "key", Path: []string{"key"}, Required: false, PreserveMissing: false, PreserveConfigured: false, EmptyStringIsNull: false},
			{Attribute: "enforce_uniqueness", Path: []string{"enforce_uniqueness"}, Required: false, PreserveMissing: false, PreserveConfigured: false, EmptyStringIsNull: false}},
		Root: []string{"data"}, Selector: map[string]string{"entity": "entity", "key": "key"},
		Paginate: true, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
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

	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeCustomFieldKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeCustomFieldKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomFieldKeyResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCustomFieldKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCustomFieldKeyResourceModel(ctx, &priorModel)
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

func (r *MetronomeCustomFieldKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomeCustomFieldKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomFieldKeyResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "listCustomFieldKeys-v1", Service: "metronome", Method: "POST", Path: "/v1/customFields/listKeys", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "entity", Constant: nil, Path: []string{"entities"}, Location: "body", Wrap: true, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{"entity": "entity", "key": "key"},
		Paginate: true, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "disableCustomFieldKey-v1", Service: "metronome", Method: "POST", Path: "/v1/customFields/removeKey", Read: false},
			Request: []bindingruntime.Mapping{{Attribute: "entity", Constant: nil, Path: []string{"entity"}, Location: "body", Wrap: false, OmitUnknown: false},
				{Attribute: "key", Constant: nil, Path: []string{"key"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response: []bindingruntime.ResponseField{},
			Root:     []string{"data"}, Selector: map[string]string{},
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

func (r *MetronomeCustomFieldKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var definition resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &definition)
	values := bindingruntime.EmptyValues(definition.Schema.Type().(types.ObjectType))
	if err := bindingruntime.ImportIdentity(ctx, values, req.ID, []string{"entity", "key"}, "/"); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome import", err.Error())
		return
	}
	if err := bindingruntime.ComposeIdentity(ctx, values, "id", []string{"entity", "key"}, "/"); err != nil {
		resp.Diagnostics.AddError("Invalid identity", err.Error())
		return
	}

	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeCustomFieldKeyResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_custom_field_key" {
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
