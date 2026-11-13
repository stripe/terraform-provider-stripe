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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
	"time"
)

type MetronomeRateCardResource struct{ transport bindingruntime.Transport }

type MetronomeRateCardResourceModel struct {
	CreditTypeConversions types.List   `tfsdk:"credit_type_conversions"`
	CreditTypeID          types.String `tfsdk:"credit_type_id"`
	Description           types.String `tfsdk:"description"`
	ID                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	PriceCents            types.Int64  `tfsdk:"price_cents"`
	RateMode              types.String `tfsdk:"rate_mode"`
	Rates                 types.List   `tfsdk:"rates"`
	StartingAt            types.String `tfsdk:"starting_at"`
	UsageProductID        types.String `tfsdk:"usage_product_id"`
}

func expandMetronomeRateCardResourceModel(ctx context.Context, model *MetronomeRateCardResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"credit_type_conversions": model.CreditTypeConversions.Type(ctx), "credit_type_id": model.CreditTypeID.Type(ctx), "description": model.Description.Type(ctx), "id": model.ID.Type(ctx), "name": model.Name.Type(ctx), "price_cents": model.PriceCents.Type(ctx), "rate_mode": model.RateMode.Type(ctx), "rates": model.Rates.Type(ctx), "starting_at": model.StartingAt.Type(ctx), "usage_product_id": model.UsageProductID.Type(ctx)}, map[string]attr.Value{"credit_type_conversions": model.CreditTypeConversions, "credit_type_id": model.CreditTypeID, "description": model.Description, "id": model.ID, "name": model.Name, "price_cents": model.PriceCents, "rate_mode": model.RateMode, "rates": model.Rates, "starting_at": model.StartingAt, "usage_product_id": model.UsageProductID}))
}

func NewMetronomeRateCardResource() resource.Resource { return &MetronomeRateCardResource{} }

func (r *MetronomeRateCardResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_rate_card"
}

func (r *MetronomeRateCardResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A Metronome rate card with optional explicit rates or one legacy entitled FLAT rate. External add_rate actions are appropriate when rates are unmanaged or outside the managed product/start/frequency entries. Changes to managed rates can require card replacement. Destroy archives the card.",
		Attributes: map[string]schema.Attribute{
			"credit_type_conversions": schema.ListNestedAttribute{
				Optional:      true,
				Description:   "Custom pricing-unit conversions into fiat minor units. Changing this immutable conversion list replaces the card. Omitted conversions are unmanaged; explicit lists track only their listed credit types.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"custom_credit_type_id": schema.StringAttribute{
							Required: true,
						},
						"fiat_per_custom_credit": schema.NumberAttribute{
							Required:    true,
							Description: "Positive finite fiat minor units per custom credit, e.g. 2 USD cents per AI credit.",
						},
					},
				},
			},
			"credit_type_id": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Metronome credit type ID. Omit to use the account's default USD (cents) credit type.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"id": schema.StringAttribute{
				Computed: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"price_cents": schema.Int64Attribute{
				Optional:      true,
				Description:   "Legacy convenience: minor units per metric unit. Required with usage_product_id.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"rate_mode": schema.StringAttribute{
				Computed:      true,
				Description:   "Provider tracking mode: none, legacy, or list.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"rates": schema.ListNestedAttribute{
				Optional:      true,
				Description:   "Explicit rates to add when creating the card. Each is an addRate API call. Rates cannot be deleted individually; replacing the card archives it.",
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"billing_frequency": schema.StringAttribute{
							Optional:    true,
							Description: "Required for a FLAT rate on a SUBSCRIPTION product (e.g. MONTHLY).",
						},
						"credit_type_id": schema.StringAttribute{
							Optional:    true,
							Description: "Defaults to the rate card's credit type.",
						},
						"entitled": schema.BoolAttribute{
							Required: true,

							PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
						},
						"price": schema.NumberAttribute{
							Required:    true,
							Description: "Minor units for FLAT; fractional percentage for PERCENTAGE.",
						},
						"product_id": schema.StringAttribute{
							Required: true,
						},
						"rate_type": schema.StringAttribute{
							Required:    true,
							Description: "FLAT or PERCENTAGE. FLAT with billing_frequency prices a SUBSCRIPTION product.",
						},
						"starting_at": schema.StringAttribute{
							Required:    true,
							Description: "Inclusive ISO 8601 effective timestamp.",
						},
					},
				},
			},
			"starting_at": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "UTC hour-boundary effective timestamp. Omit to start at the beginning of the current UTC hour when created.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"usage_product_id": schema.StringAttribute{
				Optional:      true,
				Description:   "Legacy convenience: usage product for a single FLAT rate. Omit to create the card alone.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *MetronomeRateCardResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MetronomeRateCardResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomeRateCardResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeRateCardResourceModel(ctx, &valuesModel)
	if err := validateRateCard(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomeRateCardResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomeRateCardResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeRateCardResourceModel(ctx, &valuesModel)
	if err := validateRateCard(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
	if err := prepareRateCard(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome request", err.Error())
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
		Operation: bindingruntime.Operation{ID: "createRateCard-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/create", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "description", Constant: nil, Path: []string{"description"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "credit_type_id", Constant: nil, Path: []string{"fiat_credit_type_id"}, Location: "body", Wrap: false, OmitUnknown: true},
			{Attribute: "credit_type_conversions", Constant: nil, Path: []string{"credit_type_conversions"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{{Attribute: "id", Path: []string{"id"}, Required: true, PreserveMissing: false, PreserveConfigured: false, EmptyStringIsNull: false}},
		Root:     []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: true, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "getRateCard-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/get", Read: true},
			Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"id"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response:  []bindingruntime.ResponseField{},
			Root:      []string{"data"}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: readRateCard, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "credit_type_id", Persist: false, PersistAttributes: []string{},
			ArchivedPaths: [][]string{[]string{"archived_at"}}, MissingStatuses: []int{},
			WaitForReady: false, IgnoreResponse: false, OnError: nil,
		},
		{
			Operation: bindingruntime.Operation{ID: "addRate-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/addRate", Read: false},
			Request:   []bindingruntime.Mapping{},
			Response:  []bindingruntime.ResponseField{},
			Root:      []string{"data"}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: nil, ForEach: rateRequests, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
			ArchivedPaths: [][]string{}, MissingStatuses: []int{},
			WaitForReady: false, IgnoreResponse: true, OnError: nil,
		},
		{
			Operation: bindingruntime.Operation{ID: "getRateCard-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/get", Read: true},
			Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"id"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response:  []bindingruntime.ResponseField{},
			Root:      []string{"data"}, Selector: map[string]string{},
			Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: nil, ResponseHook: readRateCard, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
			ArchivedPaths: [][]string{[]string{"archived_at"}}, MissingStatuses: []int{},
			WaitForReady: false, IgnoreResponse: false, OnError: nil,
		}}, values, nil, persist); err != nil {
		if values.Known("id") {
			cleanupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second)
			defer cancel()
			cleanupErr := bindingruntime.Run(cleanupCtx, r.transport, []bindingruntime.Step{{
				Operation: bindingruntime.Operation{ID: "archiveRateCard-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/archive", Read: false},
				Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"id"}, Location: "body", Wrap: false, OmitUnknown: false}},
				Response:  []bindingruntime.ResponseField{},
				Root:      []string{"data"}, Selector: map[string]string{},
				Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
				RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
				WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
				ArchivedPaths: [][]string{}, MissingStatuses: []int{},
				WaitForReady: false, IgnoreResponse: true, OnError: nil,
			}}, values, nil, nil)
			if cleanupErr == nil || bindingruntime.IsMissing(cleanupErr) {
				resp.State.RemoveResource(ctx)
			} else {
				resp.Diagnostics.AddError("Cleanup failed", "The created identity remains in state for recovery: "+cleanupErr.Error())
			}
		}
		resp.Diagnostics.AddError("Could not create Metronome resource", err.Error())
		return
	}
	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeRateCardResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomeRateCardResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeRateCardResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getRateCard-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/get", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"id"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readRateCard, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{[]string{"archived_at"}}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "getRateSchedule-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/getRateSchedule", Read: true},
			Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"rate_card_id"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response:  []bindingruntime.ResponseField{},
			Root:      []string{"data"}, Selector: map[string]string{},
			Paginate: true, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: rateScheduleRequest, ResponseHook: readRateSchedule, ForEach: nil, ForEachQuery: "",
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

func (r *MetronomeRateCardResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomeRateCardResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeRateCardResourceModel(ctx, &valuesModel)
	var priorModel MetronomeRateCardResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomeRateCardResourceModel(ctx, &priorModel)
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
		Operation: bindingruntime.Operation{ID: "updateRateCard-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/update", Read: false},
		Request: []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"rate_card_id"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "name", Constant: nil, Path: []string{"name"}, Location: "body", Wrap: false, OmitUnknown: false},
			{Attribute: "description", Constant: nil, Path: []string{"description"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response: []bindingruntime.ResponseField{},
		Root:     []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: rateCardUpdateRequest, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
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

func (r *MetronomeRateCardResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomeRateCardResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeRateCardResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "archiveRateCard-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/archive", Read: false},
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

func (r *MetronomeRateCardResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var definition resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &definition)
	values := bindingruntime.EmptyValues(definition.Schema.Type().(types.ObjectType))
	if err := importRateCard(ctx, values, req.ID); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome import", err.Error())
		return
	}
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getRateCard-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/get", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"id"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: verifyRateCardImport, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{[]string{"archived_at"}}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	},
		{
			Operation: bindingruntime.Operation{ID: "getRateSchedule-v1", Service: "metronome", Method: "POST", Path: "/v1/contract-pricing/rate-cards/getRateSchedule", Read: true},
			Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"rate_card_id"}, Location: "body", Wrap: false, OmitUnknown: false}},
			Response:  []bindingruntime.ResponseField{},
			Root:      []string{"data"}, Selector: map[string]string{},
			Paginate: true, PaginationLocation: "query", PaginationCursor: "next_page",
			RequestHook: rateScheduleRequest, ResponseHook: verifyRateImport, ForEach: nil, ForEachQuery: "",
			WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
			ArchivedPaths: [][]string{}, MissingStatuses: []int{},
			WaitForReady: false, IgnoreResponse: false, OnError: nil,
		}}, values, nil, nil); err != nil {
		resp.Diagnostics.AddError("Could not verify import", err.Error())
		return
	}

	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomeRateCardResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_rate_card" {
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
