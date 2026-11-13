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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeCustomerAlertResource struct{ transport bindingruntime.Transport }

type MetronomeCustomerAlertResourceModel struct {
	AlertType          types.String `tfsdk:"alert_type"`
	BillableMetricID   types.String `tfsdk:"billable_metric_id"`
	CreditTypeID       types.String `tfsdk:"credit_type_id"`
	CustomerID         types.String `tfsdk:"customer_id"`
	DeletionProtection types.Bool   `tfsdk:"deletion_protection"`
	EvaluateOnCreate   types.Bool   `tfsdk:"evaluate_on_create"`
	ID                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	Status             types.String `tfsdk:"status"`
	Threshold          types.Number `tfsdk:"threshold"`
	UniquenessKey      types.String `tfsdk:"uniqueness_key"`
}

func expandMetronomeCustomerAlertResourceModel(ctx context.Context, model *MetronomeCustomerAlertResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"alert_type": model.AlertType.Type(ctx), "billable_metric_id": model.BillableMetricID.Type(ctx), "credit_type_id": model.CreditTypeID.Type(ctx), "customer_id": model.CustomerID.Type(ctx), "deletion_protection": model.DeletionProtection.Type(ctx), "evaluate_on_create": model.EvaluateOnCreate.Type(ctx), "id": model.ID.Type(ctx), "name": model.Name.Type(ctx), "status": model.Status.Type(ctx), "threshold": model.Threshold.Type(ctx), "uniqueness_key": model.UniquenessKey.Type(ctx)}, map[string]attr.Value{"alert_type": model.AlertType, "billable_metric_id": model.BillableMetricID, "credit_type_id": model.CreditTypeID, "customer_id": model.CustomerID, "deletion_protection": model.DeletionProtection, "evaluate_on_create": model.EvaluateOnCreate, "id": model.ID, "name": model.Name, "status": model.Status, "threshold": model.Threshold, "uniqueness_key": model.UniquenessKey}))
}

func NewMetronomeCustomerAlertResource() resource.Resource { return &MetronomeCustomerAlertResource{} }

func (r *MetronomeCustomerAlertResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_customer_alert"
}

func (r *MetronomeCustomerAlertResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A customer-scoped Metronome alert created by this resource. Import is unavailable because the alert GET response does not identify whether an alert is customer-scoped or account-wide. Destroy archives it. Advanced filters remain available through the generated create-alert action.",
		Attributes: map[string]schema.Attribute{
			"alert_type": schema.StringAttribute{
				Required:      true,
				Description:   "CreateCustomerAlertPayload.alert_type from pinned OpenAPI.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"billable_metric_id": schema.StringAttribute{
				Optional:      true,
				Description:   "For usage threshold alerts. The alert GET response omits this field, so Terraform retains the configured value. Changes replace the alert.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"credit_type_id": schema.StringAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"customer_id": schema.StringAttribute{
				Required: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"deletion_protection": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set false and apply before archiving this alert.",
				Default:     booldefault.StaticBool(true),
			},
			"evaluate_on_create": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Whether to evaluate the alert immediately against existing usage. Defaults to false.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
				Default:       booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"status": schema.StringAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"threshold": schema.NumberAttribute{
				Required: true,

				PlanModifiers: []planmodifier.Number{numberplanmodifier.RequiresReplace()},
			},
			"uniqueness_key": schema.StringAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *MetronomeCustomerAlertResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeCustomerAlertResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeCustomerAlertResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerAlertResourceModel(ctx, &valuesModel)
	if err := validateAlert(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeCustomerAlertResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}
	var valuesModel MetronomeCustomerAlertResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerAlertResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCustomerAlertResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCustomerAlertResourceModel(ctx, &priorModel)
	replacement, err := planAlert(ctx, values, prior)
	if err != nil {
		resp.Diagnostics.AddError("Unsafe Metronome change", err.Error())
		return
	}
	for _, attribute := range replacement {
		resp.RequiresReplace = append(resp.RequiresReplace, path.Root(attribute))
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, values.Object)...)
}

func (r *MetronomeCustomerAlertResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeCustomerAlertResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerAlertResourceModel(ctx, &valuesModel)
	if err := validateAlert(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
	if !values.Known("status") {
		if err := values.SetJSON(ctx, "status", bindingruntime.Literal("\"enabled\"")); err != nil {
			resp.Diagnostics.AddError("Invalid state default", err.Error())
			return
		}
	}
	if !values.Known("deletion_protection") {
		if err := values.SetJSON(ctx, "deletion_protection", bindingruntime.Literal("true")); err != nil {
			resp.Diagnostics.AddError("Invalid state default", err.Error())
			return
		}
	}
	if !values.Known("evaluate_on_create") {
		if err := values.SetJSON(ctx, "evaluate_on_create", bindingruntime.Literal("false")); err != nil {
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
		Operation: bindingruntime.Operation{ID: "createAlert-v1", Service: "metronome", Method: "POST", Path: "/v1/alerts/create", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "alert_type", Constant: nil, Path: []string{"alert_type"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "threshold", Constant: nil, Path: []string{"threshold"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "credit_type_id", Constant: nil, Path: []string{"credit_type_id"}, Location: "body", Wrap: false, OmitUnknown: true},
			{Attribute: "billable_metric_id", Constant: nil, Path: []string{"billable_metric_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "uniqueness_key", Constant: nil, Path: []string{"uniqueness_key"}, Location: "body", Wrap: false, OmitUnknown: true},
			{Attribute: "evaluate_on_create", Constant: nil, Path: []string{"evaluate_on_create"}, Location: "body", Wrap: false, OmitUnknown: false}},
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

func (r *MetronomeCustomerAlertResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeCustomerAlertResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerAlertResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getCustomerAlert-v1", Service: "metronome", Method: "POST", Path: "/v1/customer-alerts/get", Read: true},
		Request: []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "id", Constant: nil, Path: []string{"alert_id"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{},
		Root:     []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readAlert, ForEach: nil, ForEachQuery: "",
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
	if !values.Known("status") {
		if err := values.SetJSON(ctx, "status", bindingruntime.Literal("\"enabled\"")); err != nil {
			resp.Diagnostics.AddError("Invalid state default", err.Error())
			return
		}
	}
	if !values.Known("deletion_protection") {
		if err := values.SetJSON(ctx, "deletion_protection", bindingruntime.Literal("true")); err != nil {
			resp.Diagnostics.AddError("Invalid state default", err.Error())
			return
		}
	}
	if !values.Known("evaluate_on_create") {
		if err := values.SetJSON(ctx, "evaluate_on_create", bindingruntime.Literal("false")); err != nil {
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

func (r *MetronomeCustomerAlertResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeCustomerAlertResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerAlertResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCustomerAlertResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCustomerAlertResourceModel(ctx, &priorModel)
	if err := values.Set("id", prior.Get("id")); err != nil {
		resp.Diagnostics.AddError("Invalid resource identity", err.Error())
		return
	}
	if err := updateAlert(ctx, values, prior); err != nil {
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

func (r *MetronomeCustomerAlertResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomeCustomerAlertResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerAlertResourceModel(ctx, &valuesModel)
	if protection, ok := values.Get("deletion_protection").(types.Bool); !ok || protection.IsNull() || protection.IsUnknown() || protection.ValueBool() {
		resp.Diagnostics.AddError("Metronome deletion protected", "Set deletion_protection = false and apply before destroying this resource.")
		return
	}
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "archiveAlert-v1", Service: "metronome", Method: "POST", Path: "/v1/alerts/archive", Read: false},
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

func (r *MetronomeCustomerAlertResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_customer_alert" {
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
