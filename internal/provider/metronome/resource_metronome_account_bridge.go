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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeAccountBridgeResource struct{ transport bindingruntime.Transport }

type MetronomeAccountBridgeResourceModel struct {
	ID               types.String `tfsdk:"id"`
	StripeAccountID  types.String `tfsdk:"stripe_account_id"`
	Mode             types.String `tfsdk:"mode"`
	ClientID         types.String `tfsdk:"client_id"`
	ReadinessTimeout types.String `tfsdk:"readiness_timeout"`
}

func expandMetronomeAccountBridgeResourceModel(ctx context.Context, model *MetronomeAccountBridgeResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"id": model.ID.Type(ctx), "stripe_account_id": model.StripeAccountID.Type(ctx), "mode": model.Mode.Type(ctx), "client_id": model.ClientID.Type(ctx), "readiness_timeout": model.ReadinessTimeout.Type(ctx)}, map[string]attr.Value{"id": model.ID, "stripe_account_id": model.StripeAccountID, "mode": model.Mode, "client_id": model.ClientID, "readiness_timeout": model.ReadinessTimeout}))
}

func NewMetronomeAccountBridgeResource() resource.Resource { return &MetronomeAccountBridgeResource{} }

func (r *MetronomeAccountBridgeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_account_bridge"
}

func (r *MetronomeAccountBridgeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Activates the Stripe account for Metronome and waits for shared-key access. Destroy forgets state without revoking access.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"stripe_account_id": schema.StringAttribute{
				Computed: true,
			},
			"mode": schema.StringAttribute{
				Computed: true,
			},
			"client_id": schema.StringAttribute{
				Computed: true,
			},
			"readiness_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Wait after activation for authorization caches to expire. Must be at least 5m30s.",
				Default:     stringdefault.StaticString("7m"),
			},
		},
	}
}

func (r *MetronomeAccountBridgeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeAccountBridgeResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeAccountBridgeResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeAccountBridgeResourceModel(ctx, &valuesModel)
	if err := validateBridge(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeAccountBridgeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeAccountBridgeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeAccountBridgeResourceModel(ctx, &valuesModel)
	if err := validateBridge(ctx, values); err != nil {
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
		Operation: bindingruntime.Operation{ID: "stripeAccountIdentity", Service: "stripe", Method: "GET", Path: "/v1/account", Read: true},
		Request:   []bindingruntime.Mapping{},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readBridgeIdentity, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{-1},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "stripeMetronomeActivate", Service: "stripe", Method: "POST", Path: "/v1/experimental-met-activate", Read: false},
			Request:   []bindingruntime.Mapping{},
			Response:  []bindingruntime.ResponseField{},
			Root:      []string{}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: activatedBridge, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: true, PersistAttributes: []string{},
			ArchivedPaths: [][]string{}, MissingStatuses: []int{-1},
			WaitForReady: false, IgnoreResponse: false, OnError: nil,
		},
		{
			Operation: bindingruntime.Operation{ID: "listCreditTypes-v1", Service: "metronome", Method: "GET", Path: "/v1/credit-types/list", Read: true},
			Request:   []bindingruntime.Mapping{{Attribute: "", Constant: json.RawMessage("\"1\""), Path: []string{"limit"}, Location: "query", Wrap: false, OmitUnknown: false}},
			Response:  []bindingruntime.ResponseField{},
			Root:      []string{}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: checkBridgeProbe, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
			ArchivedPaths: [][]string{}, MissingStatuses: []int{401, 403, 404},
			WaitForReady: true, IgnoreResponse: false, OnError: nil,
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

func (r *MetronomeAccountBridgeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeAccountBridgeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeAccountBridgeResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "stripeAccountIdentity", Service: "stripe", Method: "GET", Path: "/v1/account", Read: true},
		Request:   []bindingruntime.Mapping{},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readBridgeIdentity, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{-1},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "listCreditTypes-v1", Service: "metronome", Method: "GET", Path: "/v1/credit-types/list", Read: true},
			Request:   []bindingruntime.Mapping{{Attribute: "", Constant: json.RawMessage("\"1\""), Path: []string{"limit"}, Location: "query", Wrap: false, OmitUnknown: false}},
			Response:  []bindingruntime.ResponseField{},
			Root:      []string{}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: checkBridgeProbe, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
			ArchivedPaths: [][]string{}, MissingStatuses: []int{401, 403, 404},
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

func (r *MetronomeAccountBridgeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeAccountBridgeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeAccountBridgeResourceModel(ctx, &valuesModel)
	var priorModel MetronomeAccountBridgeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeAccountBridgeResourceModel(ctx, &priorModel)
	if err := values.Set("id", prior.Get("id")); err != nil {
		resp.Diagnostics.AddError("Invalid resource identity", err.Error())
		return
	}
	if err := updateBridge(ctx, values, prior); err != nil {
		resp.Diagnostics.AddError("Unsupported Metronome update", err.Error())
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

func (r *MetronomeAccountBridgeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *MetronomeAccountBridgeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var definition resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &definition)
	values := bindingruntime.EmptyValues(definition.Schema.Type().(types.ObjectType))
	if err := importBridge(ctx, values, req.ID); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome import", err.Error())
		return
	}
	if err := values.SetJSON(ctx, "readiness_timeout", bindingruntime.Literal("\"7m\"")); err != nil {
		resp.Diagnostics.AddError("Invalid import default", err.Error())
		return
	}
	if err := bindingruntime.ComposeIdentity(ctx, values, "id", []string{"stripe_account_id", "mode"}, "/"); err != nil {
		resp.Diagnostics.AddError("Invalid bridge identity", err.Error())
		return
	}

	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}
