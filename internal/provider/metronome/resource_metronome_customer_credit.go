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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeCustomerCreditResource struct{ transport bindingruntime.Transport }

type MetronomeCustomerCreditResourceModel struct {
	AccessSchedule       types.Object `tfsdk:"access_schedule"`
	ApplicableProductIDs types.Set    `tfsdk:"applicable_product_ids"`
	CommitType           types.String `tfsdk:"commit_type"`
	CustomerID           types.String `tfsdk:"customer_id"`
	DeletionProtection   types.Bool   `tfsdk:"deletion_protection"`
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Priority             types.Number `tfsdk:"priority"`
	ProductID            types.String `tfsdk:"product_id"`
	UniquenessKey        types.String `tfsdk:"uniqueness_key"`
}

func expandMetronomeCustomerCreditResourceModel(ctx context.Context, model *MetronomeCustomerCreditResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"access_schedule": model.AccessSchedule.Type(ctx), "applicable_product_ids": model.ApplicableProductIDs.Type(ctx), "commit_type": model.CommitType.Type(ctx), "customer_id": model.CustomerID.Type(ctx), "deletion_protection": model.DeletionProtection.Type(ctx), "id": model.ID.Type(ctx), "name": model.Name.Type(ctx), "priority": model.Priority.Type(ctx), "product_id": model.ProductID.Type(ctx), "uniqueness_key": model.UniquenessKey.Type(ctx)}, map[string]attr.Value{"access_schedule": model.AccessSchedule, "applicable_product_ids": model.ApplicableProductIDs, "commit_type": model.CommitType, "customer_id": model.CustomerID, "deletion_protection": model.DeletionProtection, "id": model.ID, "name": model.Name, "priority": model.Priority, "product_id": model.ProductID, "uniqueness_key": model.UniquenessKey}))
}

func NewMetronomeCustomerCreditResource() resource.Resource {
	return &MetronomeCustomerCreditResource{}
}

func (r *MetronomeCustomerCreditResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_customer_credit"
}

func (r *MetronomeCustomerCreditResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A customer-scoped Metronome credit. Import as customer_id/credit_id. Destroy archives remaining balance.",
		Attributes: map[string]schema.Attribute{
			"access_schedule": schema.SingleNestedAttribute{
				Required:      true,
				Description:   "ScheduleDurationInput from the pinned OpenAPI specification.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"access_type": schema.StringAttribute{
						Required:    true,
						Description: "SPEND or QUANTITY.",
					},
					"credit_type_id": schema.StringAttribute{
						Required:    true,
						Description: "Pricing unit ID for this balance.",
					},
					"schedule_items": schema.ListNestedAttribute{
						Required: true,

						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"amount": schema.NumberAttribute{
									Required: true,
								},
								"ending_before": schema.StringAttribute{
									Required:    true,
									Description: "RFC 3339 exclusive end.",
								},
								"starting_at": schema.StringAttribute{
									Required:    true,
									Description: "RFC 3339 inclusive start.",
								},
							},
						},
					},
				},
			},
			"applicable_product_ids": schema.SetAttribute{
				Optional:      true,
				Description:   "Products this balance applies to. Omit to use the API default of all products.",
				PlanModifiers: []planmodifier.Set{setplanmodifier.RequiresReplace()},
				ElementType:   types.StringType,
			},
			"commit_type": schema.StringAttribute{
				Computed:      true,
				Description:   "Always CREDIT for a customer credit.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"customer_id": schema.StringAttribute{
				Required: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"deletion_protection": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set false and apply before archiving. Archiving removes remaining spendable balance.",
				Default:     booldefault.StaticBool(true),
			},
			"id": schema.StringAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"priority": schema.NumberAttribute{
				Required:      true,
				Description:   "Lower priority balances are spent first.",
				PlanModifiers: []planmodifier.Number{numberplanmodifier.RequiresReplace()},
			},
			"product_id": schema.StringAttribute{
				Required:      true,
				Description:   "Fixed product used for this balance.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"uniqueness_key": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Stable Metronome idempotency key; required for new balances. Imported balances without a key can leave it null.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *MetronomeCustomerCreditResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeCustomerCreditResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeCustomerCreditResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerCreditResourceModel(ctx, &valuesModel)
	if err := validateCustomerCredit(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeCustomerCreditResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}
	var valuesModel MetronomeCustomerCreditResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerCreditResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCustomerCreditResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCustomerCreditResourceModel(ctx, &priorModel)
	replacement, err := planCustomerBalance(ctx, values, prior)
	if err != nil {
		resp.Diagnostics.AddError("Unsafe Metronome change", err.Error())
		return
	}
	for _, attribute := range replacement {
		resp.RequiresReplace = append(resp.RequiresReplace, path.Root(attribute))
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, values.Object)...)
}

func (r *MetronomeCustomerCreditResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeCustomerCreditResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerCreditResourceModel(ctx, &valuesModel)
	if err := validateCustomerCredit(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
	if !values.Known("deletion_protection") {
		if err := values.SetJSON(ctx, "deletion_protection", bindingruntime.Literal("true")); err != nil {
			resp.Diagnostics.AddError("Invalid state default", err.Error())
			return
		}
	}
	if !values.Known("commit_type") {
		if err := values.SetJSON(ctx, "commit_type", bindingruntime.Literal("\"CREDIT\"")); err != nil {
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
		Operation: bindingruntime.Operation{ID: "createCustomerCredit-v1", Service: "metronome", Method: "POST", Path: "/v1/contracts/customerCredits/create", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "product_id", Constant: nil, Path: []string{"product_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "priority", Constant: nil, Path: []string{"priority"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "access_schedule", Constant: nil, Path: []string{"access_schedule"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "applicable_product_ids", Constant: nil, Path: []string{"applicable_product_ids"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "uniqueness_key", Constant: nil, Path: []string{"uniqueness_key"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{{Attribute: "id", Path: []string{"id"}, Required: true, PreserveMissing: false, PreserveConfigured: false, EmptyStringIsNull: false}},
		Root:     []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: balanceRequest, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
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

func (r *MetronomeCustomerCreditResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeCustomerCreditResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerCreditResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "listCustomerCredits-v1", Service: "metronome", Method: "POST", Path: "/v1/contracts/customerCredits/list", Read: true},
		Request: []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "id", Constant: nil, Path: []string{"credit_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "", Constant: json.RawMessage("false"), Path: []string{"include_archived"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "", Constant: json.RawMessage("25"), Path: []string{"limit"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{},
		Root:     []string{"data"}, Selector: map[string]string{"id": "id"},
		Paginate: true, PaginationLocation: "body", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readCustomerCredit, ForEach: nil, ForEachQuery: "",
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
	if !values.Known("commit_type") {
		if err := values.SetJSON(ctx, "commit_type", bindingruntime.Literal("\"CREDIT\"")); err != nil {
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

func (r *MetronomeCustomerCreditResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeCustomerCreditResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerCreditResourceModel(ctx, &valuesModel)
	var priorModel MetronomeCustomerCreditResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeCustomerCreditResourceModel(ctx, &priorModel)
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

func (r *MetronomeCustomerCreditResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomeCustomerCreditResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeCustomerCreditResourceModel(ctx, &valuesModel)
	if protection, ok := values.Get("deletion_protection").(types.Bool); !ok || protection.IsNull() || protection.IsUnknown() || protection.ValueBool() {
		resp.Diagnostics.AddError("Metronome deletion protected", "Set deletion_protection = false and apply before destroying this resource.")
		return
	}
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getCustomer-v1", Service: "metronome", Method: "GET", Path: "/v1/customers/{customer_id}", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "path", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{[]string{"archived_at"}}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "archiveCredit-v2", Service: "metronome", Method: "POST", Path: "/v2/contracts/credits/archive", Read: false},
			Request: []bindingruntime.Mapping{{Attribute: "customer_id", Constant: nil, Path: []string{"customer_id"}, Location: "body", Wrap: false, OmitUnknown: false},
				{Attribute: "id", Constant: nil, Path: []string{"credit_id"}, Location: "body", Wrap: false, OmitUnknown: false}},
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

func (r *MetronomeCustomerCreditResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	if !values.Known("commit_type") {
		if err := values.SetJSON(ctx, "commit_type", bindingruntime.Literal("\"CREDIT\"")); err != nil {
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

func (r *MetronomeCustomerCreditResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_customer_credit" {
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
