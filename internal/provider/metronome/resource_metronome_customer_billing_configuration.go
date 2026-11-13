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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeCustomerBillingConfigurationResource struct{ transport bindingruntime.Transport }

type MetronomeCustomerBillingConfigurationResourceModel struct {
	Configuration      types.Dynamic `tfsdk:"configuration"`
	CustomerID         types.String  `tfsdk:"customer_id"`
	DeletionProtection types.Bool    `tfsdk:"deletion_protection"`
	DeliveryMethod     types.String  `tfsdk:"delivery_method"`
	DeliveryMethodID   types.String  `tfsdk:"delivery_method_id"`
	ID                 types.String  `tfsdk:"id"`
	ProviderName       types.String  `tfsdk:"provider_name"`
	TaxProvider        types.String  `tfsdk:"tax_provider"`
}

func expandMetronomeCustomerBillingConfigurationResourceModel(ctx context.Context, model *MetronomeCustomerBillingConfigurationResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"configuration": model.Configuration.Type(ctx), "customer_id": model.CustomerID.Type(ctx), "deletion_protection": model.DeletionProtection.Type(ctx), "delivery_method": model.DeliveryMethod.Type(ctx), "delivery_method_id": model.DeliveryMethodID.Type(ctx), "id": model.ID.Type(ctx), "provider_name": model.ProviderName.Type(ctx), "tax_provider": model.TaxProvider.Type(ctx)}, map[string]attr.Value{"configuration": model.Configuration, "customer_id": model.CustomerID, "deletion_protection": model.DeletionProtection, "delivery_method": model.DeliveryMethod, "delivery_method_id": model.DeliveryMethodID, "id": model.ID, "provider_name": model.ProviderName, "tax_provider": model.TaxProvider}))
}

func NewMetronomeCustomerBillingConfigurationResource() resource.Resource {
	return &MetronomeCustomerBillingConfigurationResource{}
}

func (r *MetronomeCustomerBillingConfigurationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_customer_billing_configuration"
}

func (r *MetronomeCustomerBillingConfigurationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A customer-scoped billing provider configuration. Import as customer_id/configuration_id. Destroy archives only this configuration.",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.DynamicAttribute{
				Required:      true,
				Description:   "Provider-specific configuration object from CustomerBillingProviderConfigurationInput or CustomerRevenueSystemConfigurationInput. Server-populated keys may be omitted from HCL. Terraform still stores this value in state; use an encrypted, access-controlled backend.",
				PlanModifiers: []planmodifier.Dynamic{bindingruntime.RequiresReplaceSubset()},
				Sensitive:     true,
			},
			"customer_id": schema.StringAttribute{
				Required: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"deletion_protection": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set false and apply before archiving this configuration.",
				Default:     booldefault.StaticBool(true),
			},
			"delivery_method": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Invoice delivery method. Billing configuration creation needs this or delivery_method_id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"delivery_method_id": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Existing delivery method ID. The API returns this after creation.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"provider_name": schema.StringAttribute{
				Required:      true,
				Description:   "Maps to Metronome billing_provider or revenue provider. Terraform reserves the attribute name provider.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"tax_provider": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Billing tax provider, if set on create. Metronome GET omits this field, so imported configurations should omit it unless they are being replaced.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *MetronomeCustomerBillingConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeCustomerBillingConfigurationResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeCustomerBillingConfigurationResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerBillingConfigurationResourceModel(ctx, &valuesModel)
	if err := validateBillingConfiguration(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeCustomerBillingConfigurationResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}
	var valuesModel MetronomeCustomerBillingConfigurationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerBillingConfigurationResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCustomerBillingConfigurationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCustomerBillingConfigurationResourceModel(ctx, &priorModel)
	replacement, err := planConfiguration(ctx, values, prior)
	if err != nil {
		resp.Diagnostics.AddError("Unsafe Metronome change", err.Error())
		return
	}
	for _, attribute := range replacement {
		resp.RequiresReplace = append(resp.RequiresReplace, path.Root(attribute))
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, values.Object)...)
}

func (r *MetronomeCustomerBillingConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeCustomerBillingConfigurationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerBillingConfigurationResourceModel(ctx, &valuesModel)
	if err := validateBillingConfiguration(ctx, values); err != nil {
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
		Operation: bindingruntime.Operation{ID: "getCustomerBillingProviderConfigurations-v1", Service: "metronome", Method: "POST", Path: "/v1/getCustomerBillingProviderConfigurations", Read: true},
		Request: []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "", Constant: json.RawMessage("true"), Path: []string{"include_archived"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{},
		Root:     []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: checkBillingConfigurationAbsent, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "setCustomerBillingProviderConfigurations-v1", Service: "metronome", Method: "POST", Path: "/v1/setCustomerBillingProviderConfigurations", Read: false},
			Request: []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
				{Attribute: "provider_name", Constant: nil, Path: []string{"billing_provider"}, Location: "body", Wrap: false, OmitUnknown: false},
				{Attribute: "configuration", Constant: nil, Path: []string{"configuration"}, Location: "body", Wrap: false, OmitUnknown: false},
				{Attribute: "delivery_method_id", Constant: nil, Path: []string{"delivery_method_id"}, Location: "body", Wrap: false, OmitUnknown: true},
				{Attribute: "delivery_method", Constant: nil, Path: []string{"delivery_method"}, Location: "body", Wrap: false, OmitUnknown: true},
				{Attribute: "tax_provider", Constant: nil, Path: []string{"tax_provider"}, Location: "body", Wrap: false, OmitUnknown: true}},
			Response: []bindingruntime.ResponseField{},
			Root:     []string{"data"}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: wrapConfigurationRequest, ResponseHook: createdConfiguration, ForEach: nil, ForEachQuery: "",
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

func (r *MetronomeCustomerBillingConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeCustomerBillingConfigurationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerBillingConfigurationResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getCustomerBillingProviderConfigurations-v1", Service: "metronome", Method: "POST", Path: "/v1/getCustomerBillingProviderConfigurations", Read: true},
		Request: []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "", Constant: json.RawMessage("true"), Path: []string{"include_archived"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{},
		Root:     []string{"data"}, Selector: map[string]string{"id": "id"},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readBillingConfiguration, ForEach: nil, ForEachQuery: "",
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

func (r *MetronomeCustomerBillingConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeCustomerBillingConfigurationResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerBillingConfigurationResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCustomerBillingConfigurationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCustomerBillingConfigurationResourceModel(ctx, &priorModel)
	if err := values.Set("id", prior.Get("id")); err != nil {
		resp.Diagnostics.AddError("Invalid resource identity", err.Error())
		return
	}
	if err := updateConfiguration(ctx, values, prior); err != nil {
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

func (r *MetronomeCustomerBillingConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomeCustomerBillingConfigurationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerBillingConfigurationResourceModel(ctx, &valuesModel)
	if protection, ok := values.Get("deletion_protection").(types.Bool); !ok || protection.IsNull() || protection.IsUnknown() || protection.ValueBool() {
		resp.Diagnostics.AddError("Metronome deletion protected", "Set deletion_protection = false and apply before destroying this resource.")
		return
	}
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "archiveCustomerBillingProviderConfigurations-v1", Service: "metronome", Method: "POST", Path: "/v1/archiveCustomerBillingProviderConfigurations", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "id", Constant: nil, Path: []string{"customer_billing_provider_configuration_ids"}, Location: "body", Wrap: true, OmitUnknown: false}},
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

func (r *MetronomeCustomerBillingConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var definition resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &definition)
	values := bindingruntime.EmptyValues(definition.Schema.Type().(types.ObjectType))
	if err := bindingruntime.ImportIdentity(ctx, values, req.ID, []string{"customer_id", "id"}, "/"); err != nil {
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

func (r *MetronomeCustomerBillingConfigurationResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_customer_billing_configuration" {
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
