//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeCustomerResource struct{ transport bindingruntime.Transport }

type MetronomeCustomerResourceModel struct {
	DeletionProtection types.Bool   `tfsdk:"deletion_protection"`
	ID                 types.String `tfsdk:"id"`
	IngestAliases      types.Set    `tfsdk:"ingest_aliases"`
	Name               types.String `tfsdk:"name"`
}

func expandMetronomeCustomerResourceModel(ctx context.Context, model *MetronomeCustomerResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"deletion_protection": model.DeletionProtection.Type(ctx), "id": model.ID.Type(ctx), "ingest_aliases": model.IngestAliases.Type(ctx), "name": model.Name.Type(ctx)}, map[string]attr.Value{"deletion_protection": model.DeletionProtection, "id": model.ID, "ingest_aliases": model.IngestAliases, "name": model.Name}))
}

func NewMetronomeCustomerResource() resource.Resource { return &MetronomeCustomerResource{} }

func (r *MetronomeCustomerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_customer"
}

func (r *MetronomeCustomerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Metronome customer and its usage ingest aliases. Customer archival can archive contracts and void invoices; deletion protection is enabled by default.",
		Attributes: map[string]schema.Attribute{
			"deletion_protection": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set false and apply before destroying. Archiving a customer can affect invoices and every contract, including unmanaged ones.",
				Default:     booldefault.StaticBool(true),
			},
			"id": schema.StringAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ingest_aliases": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Aliases used as customer_id in usage events. Setting this replaces the full alias set.",
				Default:     setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{})),
				ElementType: types.StringType,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *MetronomeCustomerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeCustomerResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeCustomerResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerResourceModel(ctx, &valuesModel)
	if err := validateCustomer(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeCustomerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeCustomerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerResourceModel(ctx, &valuesModel)
	if err := validateCustomer(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
	if !values.Known("deletion_protection") {
		if err := values.SetJSON(ctx, "deletion_protection", bindingruntime.Literal("true")); err != nil {
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
		Operation: bindingruntime.Operation{ID: "listCustomers-v1", Service: "metronome", Method: "GET", Path: "/v1/customers", Read: true},
		Request:   []bindingruntime.Mapping{},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: true, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: checkAliasOwnership, ForEach: customerAliases, ForEachQuery: "ingest_alias",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "createCustomer-v1", Service: "metronome", Method: "POST", Path: "/v1/customers", Read: false},
			Request: []bindingruntime.Mapping{{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: false},
				{Attribute: "ingest_aliases", Constant: nil, Path: []string{"ingest_aliases"}, Location: "body", Wrap: false, OmitUnknown: false}},
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

func (r *MetronomeCustomerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeCustomerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getCustomer-v1", Service: "metronome", Method: "GET", Path: "/v1/customers/{customer_id}", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"customer_id"}, Location: "path", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readCustomer, ForEach: nil, ForEachQuery: "",
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
	if !values.Known("deletion_protection") {
		if err := values.SetJSON(ctx, "deletion_protection", bindingruntime.Literal("true")); err != nil {
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

func (r *MetronomeCustomerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeCustomerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCustomerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCustomerResourceModel(ctx, &priorModel)
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
		Operation: bindingruntime.Operation{ID: "listCustomers-v1", Service: "metronome", Method: "GET", Path: "/v1/customers", Read: true},
		Request:   []bindingruntime.Mapping{},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: true, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: checkAliasOwnership, ForEach: customerAliases, ForEachQuery: "ingest_alias",
		WhenChanged: []string{"ingest_aliases"}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "setCustomerName-v1", Service: "metronome", Method: "POST", Path: "/v1/customers/{customer_id}/setName", Read: false},
			Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"customer_id"}, Location: "path", Wrap: false, OmitUnknown: false},
				{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response: []bindingruntime.ResponseField{},
			Root:     []string{"data"}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{"name"}, WhenMissing: "", Persist: true, PersistAttributes: []string{"name"},
			ArchivedPaths: [][]string{}, MissingStatuses: []int{},
			WaitForReady: false, IgnoreResponse: true, OnError: nil,
		},
		{
			Operation: bindingruntime.Operation{ID: "setIngestAliases-v1", Service: "metronome", Method: "POST", Path: "/v1/customers/{customer_id}/setIngestAliases", Read: false},
			Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"customer_id"}, Location: "path", Wrap: false, OmitUnknown: false},
				{Attribute: "ingest_aliases", Constant: nil, Path: []string{"ingest_aliases"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response: []bindingruntime.ResponseField{},
			Root:     []string{"data"}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{"ingest_aliases"}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
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

func (r *MetronomeCustomerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomeCustomerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerResourceModel(ctx, &valuesModel)
	if protection, ok := values.Get("deletion_protection").(types.Bool); !ok || protection.IsNull() || protection.IsUnknown() || protection.ValueBool() {
		resp.Diagnostics.AddError("Metronome deletion protected", "Set deletion_protection = false and apply before destroying this resource.")
		return
	}
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getCustomer-v1", Service: "metronome", Method: "GET", Path: "/v1/customers/{customer_id}", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"customer_id"}, Location: "path", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readCustomer, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{[]string{"archived_at"}}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "listContracts-v2", Service: "metronome", Method: "POST", Path: "/v2/contracts/list", Read: true},
			Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
				{Attribute: "", Constant: json.RawMessage("false"), Path: []string{"include_archived"}, Location: "body", Wrap: false, OmitUnknown: false},
				{Attribute: "", Constant: json.RawMessage("20"), Path: []string{"limit"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response: []bindingruntime.ResponseField{},
			Root:     []string{"data"}, Selector: map[string]string{},
			Paginate: true, PaginationLocation: "body", PaginationCursor: "cursor",
			RequestHook: nil, ResponseHook: checkCustomerContracts, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
			ArchivedPaths: [][]string{}, MissingStatuses: []int{-1},
			WaitForReady: false, IgnoreResponse: false, OnError: nil,
		},
		{
			Operation: bindingruntime.Operation{ID: "setIngestAliases-v1", Service: "metronome", Method: "POST", Path: "/v1/customers/{customer_id}/setIngestAliases", Read: false},
			Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"customer_id"}, Location: "path", Wrap: false, OmitUnknown: false},
				{Attribute: "", Constant: json.RawMessage("[]"), Path: []string{"ingest_aliases"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response: []bindingruntime.ResponseField{},
			Root:     []string{"data"}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
			ArchivedPaths: [][]string{}, MissingStatuses: []int{},
			WaitForReady: false, IgnoreResponse: true, OnError: nil,
		},
		{
			Operation: bindingruntime.Operation{ID: "archiveCustomer-v1", Service: "metronome", Method: "POST", Path: "/v1/customers/archive", Read: false},
			Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"id"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response:  []bindingruntime.ResponseField{},
			Root:      []string{"data"}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
			ArchivedPaths: [][]string{}, MissingStatuses: []int{},
			WaitForReady: false, IgnoreResponse: true, OnError: []bindingruntime.Step{{
				Operation: bindingruntime.Operation{ID: "getCustomer-v1", Service: "metronome", Method: "GET", Path: "/v1/customers/{customer_id}", Read: true},
				Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"customer_id"}, Location: "path", Wrap: false, OmitUnknown: false}},
				Response:  []bindingruntime.ResponseField{},
				Root:      []string{"data"}, Selector: map[string]string{},
				Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
				RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
				WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
				ArchivedPaths: [][]string{[]string{"archived_at"}}, MissingStatuses: []int{},
				WaitForReady: false, IgnoreResponse: false, OnError: nil,
			},
				{
					Operation: bindingruntime.Operation{ID: "setIngestAliases-v1", Service: "metronome", Method: "POST", Path: "/v1/customers/{customer_id}/setIngestAliases", Read: false},
					Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"customer_id"}, Location: "path", Wrap: false, OmitUnknown: false},
						{Attribute: "ingest_aliases", Constant: nil, Path: []string{"ingest_aliases"}, Location: "body", Wrap: false, OmitUnknown: false}},
					Response: []bindingruntime.ResponseField{},
					Root:     []string{"data"}, Selector: map[string]string{},
					Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
					RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
					WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
					ArchivedPaths: [][]string{}, MissingStatuses: []int{},
					WaitForReady: false, IgnoreResponse: true, OnError: nil,
				}},
		}}, values, nil, nil); err != nil && !bindingruntime.IsMissing(err) {
		resp.Diagnostics.AddError("Could not archive Metronome resource", err.Error())
		return
	}
}

func (r *MetronomeCustomerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var definition resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &definition)
	values := bindingruntime.EmptyValues(definition.Schema.Type().(types.ObjectType))
	if err := bindingruntime.ImportIdentity(ctx, values, req.ID, []string{"id"}, "/"); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome import", err.Error())
		return
	}
	if !values.Known("deletion_protection") {
		if err := values.SetJSON(ctx, "deletion_protection", bindingruntime.Literal("true")); err != nil {
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

func (r *MetronomeCustomerResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_customer" {
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
