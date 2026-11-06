//
// File generated from our OpenAPI spec
//

package actions

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/action"
	actionSchema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ action.Action = &TreasuryDebitReversalAction{}
var _ action.ActionWithConfigure = &TreasuryDebitReversalAction{}

func NewTreasuryDebitReversalAction() action.Action {
	return &TreasuryDebitReversalAction{}
}

type TreasuryDebitReversalAction struct {
	client *stripe.Client
}

type TreasuryDebitReversalResourceModel struct {
	Metadata      types.Map    `tfsdk:"metadata"`
	ReceivedDebit types.String `tfsdk:"received_debit"`
}

func (r *TreasuryDebitReversalAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_treasury_debit_reversal"
}

func (r *TreasuryDebitReversalAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionSchema.Schema{
		Description: "You can reverse some [ReceivedDebits](https://api.stripe.com#received_debits) depending on their network and source flow. Reversing a ReceivedDebit leads to the creation of a new object known as a DebitReversal.",
		Attributes: map[string]actionSchema.Attribute{
			"metadata": actionSchema.MapAttribute{
				Optional:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"received_debit": actionSchema.StringAttribute{
				Required:    true,
				Description: "The ReceivedDebit being reversed.",
			},
		},
	}
}

func (r *TreasuryDebitReversalAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stripe.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Action Configure Type", fmt.Sprintf("Expected *stripe.Client, got: %T", req.ProviderData))
		return
	}

	r.client = client
}

func expandTreasuryDebitReversalCreate(plan TreasuryDebitReversalResourceModel) (*stripe.TreasuryDebitReversalCreateParams, error) {
	params := &stripe.TreasuryDebitReversalCreateParams{}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.ReceivedDebit.IsNull() && !plan.ReceivedDebit.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReceivedDebit", "ReceivedDebit", plan.ReceivedDebit.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "received_debit", params)
		}
	}

	return params, nil
}

func (r *TreasuryDebitReversalAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config TreasuryDebitReversalResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTreasuryDebitReversalCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building TreasuryDebitReversal action params", err.Error())
		return
	}

	obj, err := r.client.V1TreasuryDebitReversals.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking TreasuryDebitReversal action", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Created stripe_treasury_debit_reversal %s", obj.ID)})
	}
}
