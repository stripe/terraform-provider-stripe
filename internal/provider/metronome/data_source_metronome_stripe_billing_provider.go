//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomeStripeBillingProviderDataSource struct{ transport bindingruntime.Transport }

type MetronomeStripeBillingProviderDataSourceModel struct {
	ID                types.String `tfsdk:"id"`
	ExpectedAccountID types.String `tfsdk:"expected_account_id"`
}

func expandMetronomeStripeBillingProviderDataSourceModel(ctx context.Context, model *MetronomeStripeBillingProviderDataSourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"id": model.ID.Type(ctx), "expected_account_id": model.ExpectedAccountID.Type(ctx)}, map[string]attr.Value{"id": model.ID, "expected_account_id": model.ExpectedAccountID}))
}

func NewMetronomeStripeBillingProviderDataSource() datasource.DataSource {
	return &MetronomeStripeBillingProviderDataSource{}
}

func (r *MetronomeStripeBillingProviderDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_stripe_billing_provider"
}

func (r *MetronomeStripeBillingProviderDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Checks that exactly one direct Stripe billing integration exists in the Metronome account.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"expected_account_id": schema.StringAttribute{
				Optional:    true,
				Description: "Require this Stripe account ID in integration metadata.",
			},
		},
	}
}

func (r *MetronomeStripeBillingProviderDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MetronomeStripeBillingProviderDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var valuesModel MetronomeStripeBillingProviderDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomeStripeBillingProviderDataSourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "listConfiguredBillingProviders-v1", Service: "metronome", Method: "POST", Path: "/v1/listConfiguredBillingProviders", Read: true},
		Request:   []bindingruntime.Mapping{},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: true, PaginationLocation: "body", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readStripeBillingProvider, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	}}, values, nil, nil); err != nil {
		resp.Diagnostics.AddError("Could not read Metronome data", err.Error())
		return
	}
	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}
