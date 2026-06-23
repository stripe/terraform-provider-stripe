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

var _ action.Action = &TreasuryCreditReversalAction{}
var _ action.ActionWithConfigure = &TreasuryCreditReversalAction{}

func NewTreasuryCreditReversalAction() action.Action {
	return &TreasuryCreditReversalAction{}
}

type TreasuryCreditReversalAction struct {
	client *stripe.Client
}

type TreasuryCreditReversalResourceModel struct {
	Metadata       types.Map    `tfsdk:"metadata"`
	ReceivedCredit types.String `tfsdk:"received_credit"`
}

func (r *TreasuryCreditReversalAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_treasury_credit_reversal"
}

func (r *TreasuryCreditReversalAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionSchema.Schema{
		Description: "You can reverse some [ReceivedCredits](https://api.stripe.com#received_credits) depending on their network and source flow. Reversing a ReceivedCredit leads to the creation of a new object known as a CreditReversal.",
		Attributes: map[string]actionSchema.Attribute{
			"metadata": actionSchema.MapAttribute{
				Optional:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"received_credit": actionSchema.StringAttribute{
				Required:    true,
				Description: "The ReceivedCredit being reversed.",
			},
		},
	}
}

func (r *TreasuryCreditReversalAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func expandTreasuryCreditReversalCreate(plan TreasuryCreditReversalResourceModel) (*stripe.TreasuryCreditReversalCreateParams, error) {
	params := &stripe.TreasuryCreditReversalCreateParams{}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.ReceivedCredit.IsNull() && !plan.ReceivedCredit.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReceivedCredit", "ReceivedCredit", plan.ReceivedCredit.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "received_credit", params)
		}
	}

	return params, nil
}

func (r *TreasuryCreditReversalAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config TreasuryCreditReversalResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTreasuryCreditReversalCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building TreasuryCreditReversal action params", err.Error())
		return
	}

	obj, err := r.client.V1TreasuryCreditReversals.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking TreasuryCreditReversal action", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Created stripe_treasury_credit_reversal %s", obj.ID)})
	}
}
