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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeProductResource struct{ transport bindingruntime.Transport }

type MetronomeProductResourceModel struct {
	BillableMetricID    types.String `tfsdk:"billable_metric_id"`
	CompositeProductIDs types.List   `tfsdk:"composite_product_ids"`
	CompositeTags       types.List   `tfsdk:"composite_tags"`
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Type                types.String `tfsdk:"type"`
}

func expandMetronomeProductResourceModel(ctx context.Context, model *MetronomeProductResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"billable_metric_id": model.BillableMetricID.Type(ctx), "composite_product_ids": model.CompositeProductIDs.Type(ctx), "composite_tags": model.CompositeTags.Type(ctx), "id": model.ID.Type(ctx), "name": model.Name.Type(ctx), "type": model.Type.Type(ctx)}, map[string]attr.Value{"billable_metric_id": model.BillableMetricID, "composite_product_ids": model.CompositeProductIDs, "composite_tags": model.CompositeTags, "id": model.ID, "name": model.Name, "type": model.Type}))
}

func NewMetronomeProductResource() resource.Resource { return &MetronomeProductResource{} }

func (r *MetronomeProductResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_product"
}

func (r *MetronomeProductResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A USAGE, FIXED, SUBSCRIPTION, or COMPOSITE Metronome product. Destroy archives the product.",
		Attributes: map[string]schema.Attribute{
			"billable_metric_id": schema.StringAttribute{
				Optional:      true,
				Description:   "Required for USAGE; omit for other types.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"composite_product_ids": schema.ListAttribute{
				Optional:      true,
				Description:   "Required for COMPOSITE products.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				ElementType:   types.StringType,
			},
			"composite_tags": schema.ListAttribute{
				Optional:      true,
				Description:   "Required for COMPOSITE products.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				ElementType:   types.StringType,
			},
			"id": schema.StringAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required:      true,
				Description:   "USAGE, FIXED, SUBSCRIPTION, or COMPOSITE.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *MetronomeProductResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeProductResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeProductResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeProductResourceModel(ctx, &valuesModel)
	if err := validateProduct(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeProductResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeProductResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeProductResourceModel(ctx, &valuesModel)
	if err := validateProduct(ctx, values); err != nil {
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
		Operation: bindingruntime.Operation{ID: "createProduct-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/products/create", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "type", Constant: nil, Path: []string{"type"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "billable_metric_id", Constant: nil, Path: []string{"billable_metric_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "composite_product_ids", Constant: nil, Path: []string{"composite_product_ids"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "composite_tags", Constant: nil, Path: []string{"composite_tags"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{{Attribute: "id", Path: []string{"id"}, Required: true, PreserveMissing: false, PreserveConfigured: false, EmptyStringIsNull: false}},
		Root:     []string{"data"}, Selector: map[string]string{},
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

func (r *MetronomeProductResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeProductResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeProductResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getProduct-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/products/get", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"id"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readProduct, ForEach: nil, ForEachQuery: "",
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

func (r *MetronomeProductResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeProductResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeProductResourceModel(ctx, &valuesModel)
	var priorModel MetronomeProductResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeProductResourceModel(ctx, &priorModel)
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
		Operation: bindingruntime.Operation{ID: "updateProduct-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/products/update", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"product_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{},
		Root:     []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: true, OnError: nil,
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

func (r *MetronomeProductResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomeProductResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeProductResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "archiveProductListItem-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/products/archive", Read: false},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"product_id"}, Location: "body", Wrap: false, OmitUnknown: false}},
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

func (r *MetronomeProductResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *MetronomeProductResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_product" {
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
