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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeContractResource struct{ transport bindingruntime.Transport }

type MetronomeContractResourceModel struct {
	CustomerID            types.String `tfsdk:"customer_id"`
	EndingBefore          types.String `tfsdk:"ending_before"`
	ID                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	PackageID             types.String `tfsdk:"package_id"`
	RateCardID            types.String `tfsdk:"rate_card_id"`
	StartingAt            types.String `tfsdk:"starting_at"`
	UniquenessKey         types.String `tfsdk:"uniqueness_key"`
	VoidInvoicesOnDestroy types.Bool   `tfsdk:"void_invoices_on_destroy"`
}

func expandMetronomeContractResourceModel(ctx context.Context, model *MetronomeContractResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"customer_id": model.CustomerID.Type(ctx), "ending_before": model.EndingBefore.Type(ctx), "id": model.ID.Type(ctx), "name": model.Name.Type(ctx), "package_id": model.PackageID.Type(ctx), "rate_card_id": model.RateCardID.Type(ctx), "starting_at": model.StartingAt.Type(ctx), "uniqueness_key": model.UniquenessKey.Type(ctx), "void_invoices_on_destroy": model.VoidInvoicesOnDestroy.Type(ctx)}, map[string]attr.Value{"customer_id": model.CustomerID, "ending_before": model.EndingBefore, "id": model.ID, "name": model.Name, "package_id": model.PackageID, "rate_card_id": model.RateCardID, "starting_at": model.StartingAt, "uniqueness_key": model.UniquenessKey, "void_invoices_on_destroy": model.VoidInvoicesOnDestroy}))
}

func NewMetronomeContractResource() resource.Resource { return &MetronomeContractResource{} }

func (r *MetronomeContractResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_contract"
}

func (r *MetronomeContractResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A customer contract provisioned from a package or created directly against a rate card. Destroy archives the contract and can cancel draft invoices or expire balances. Import as customer_id/contract_id.",
		Attributes: map[string]schema.Attribute{
			"customer_id": schema.StringAttribute{
				Required: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"ending_before": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Exclusive contract end. A new instant uses editContract in place; equivalent timestamp formatting only updates Terraform state. Omitting the field adopts the remote value and does not clear an existing end date; use the explicit editContract action to clear one.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Name of a direct contract; cannot be set with package_id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"package_id": schema.StringAttribute{
				Optional:      true,
				Description:   "Package to provision from. Set exactly one of package_id or rate_card_id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"rate_card_id": schema.StringAttribute{
				Optional:      true,
				Description:   "Rate card for a direct contract. Set exactly one of package_id or rate_card_id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"starting_at": schema.StringAttribute{
				Required:      true,
				Description:   "RFC 3339 start, aligned to the hour. Equivalent timestamp formatting updates Terraform state without replacing the contract.",
				PlanModifiers: []planmodifier.String{bindingruntime.RequiresReplaceInstant()},
			},
			"uniqueness_key": schema.StringAttribute{
				Required:      true,
				Description:   "Unique 1–128 character key. Replacements require a new key even after archival.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"void_invoices_on_destroy": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether contract archival voids finalized invoices. False preserves finalized invoices, but archive still affects draft invoices and prepaid balances.",
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}

func (r *MetronomeContractResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeContractResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeContractResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeContractResourceModel(ctx, &valuesModel)
	if err := validateContract(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeContractResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}
	var valuesModel MetronomeContractResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeContractResourceModel(ctx, &valuesModel)
	var priorModel MetronomeContractResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeContractResourceModel(ctx, &priorModel)
	replacement, err := planContract(ctx, values, prior)
	if err != nil {
		resp.Diagnostics.AddError("Unsafe Metronome change", err.Error())
		return
	}
	for _, attribute := range replacement {
		resp.RequiresReplace = append(resp.RequiresReplace, path.Root(attribute))
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, values.Object)...)
}

func (r *MetronomeContractResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeContractResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeContractResourceModel(ctx, &valuesModel)
	if err := validateContract(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
	if !values.Known("void_invoices_on_destroy") {
		if err := values.SetJSON(ctx, "void_invoices_on_destroy", bindingruntime.Literal("false")); err != nil {
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
		Operation: bindingruntime.Operation{ID: "createContract-v1", Service: "metronome", Method: "POST", Path: "/v1/contracts/create", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "package_id", Constant: nil, Path: []string{"package_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "rate_card_id", Constant: nil, Path: []string{"rate_card_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: true},
			{Attribute: "starting_at", Constant: nil, Path: []string{"starting_at"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "uniqueness_key", Constant: nil, Path: []string{"uniqueness_key"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "ending_before", Constant: nil, Path: []string{"ending_before"}, Location: "body", Wrap: false, OmitUnknown: true}},
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

func (r *MetronomeContractResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeContractResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeContractResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getContract-v2", Service: "metronome", Method: "POST", Path: "/v2/contracts/get", Read: true},
		Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"contract_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{},
		Root:     []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readContract, ForEach: nil, ForEachQuery: "",
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
	if !values.Known("void_invoices_on_destroy") {
		if err := values.SetJSON(ctx, "void_invoices_on_destroy", bindingruntime.Literal("false")); err != nil {
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

func (r *MetronomeContractResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeContractResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeContractResourceModel(ctx, &valuesModel)
	var priorModel MetronomeContractResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeContractResourceModel(ctx, &priorModel)
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
		Operation: bindingruntime.Operation{ID: "editContract-v2", Service: "metronome", Method: "POST", Path: "/v2/contracts/edit", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"contract_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "ending_before", Constant: nil, Path: []string{"update_contract_end_date"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{},
		Root:     []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: contractEndRequest, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{"ending_before"}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
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

func (r *MetronomeContractResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomeContractResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeContractResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "archiveContract-v1", Service: "metronome", Method: "POST", Path: "/v1/contracts/archive", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"contract_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "void_invoices_on_destroy", Constant: nil, Path: []string{"void_invoices"}, Location: "body", Wrap: false, OmitUnknown: false}},
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

func (r *MetronomeContractResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var definition resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &definition)
	values := bindingruntime.EmptyValues(definition.Schema.Type().(types.ObjectType))
	if err := bindingruntime.ImportIdentity(ctx, values, req.ID, []string{"customer_id", "id"}, "/"); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome import", err.Error())
		return
	}
	if !values.Known("void_invoices_on_destroy") {
		if err := values.SetJSON(ctx, "void_invoices_on_destroy", bindingruntime.Literal("false")); err != nil {
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

func (r *MetronomeContractResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_contract" {
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
