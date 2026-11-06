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

var _ action.Action = &TransferReversalAction{}
var _ action.ActionWithConfigure = &TransferReversalAction{}

func NewTransferReversalAction() action.Action {
	return &TransferReversalAction{}
}

type TransferReversalAction struct {
	client *stripe.Client
}

type TransferReversalResourceModel struct {
	Amount               types.Int64  `tfsdk:"amount"`
	Metadata             types.Map    `tfsdk:"metadata"`
	Transfer             types.String `tfsdk:"transfer"`
	Description          types.String `tfsdk:"description"`
	RefundApplicationFee types.Bool   `tfsdk:"refund_application_fee"`
}

func (r *TransferReversalAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transfer_reversal"
}

func (r *TransferReversalAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionSchema.Schema{
		Description: "[Stripe Connect](https://docs.stripe.com/connect) platforms can reverse transfers made to a\nconnected account, either entirely or partially, and can also specify whether\nto refund any related application fees. Transfer reversals add to the\nplatform's balance and subtract from the destination account's balance.\n\nReversing a transfer that was made for a [destination\ncharge](/docs/connect/destination-charges) is allowed only up to the amount of\nthe charge. It is possible to reverse a\n[transfer_group](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options)\ntransfer only if the destination account has enough balance to cover the\nreversal.\n\nRelated guide: [Reverse transfers](https://docs.stripe.com/connect/separate-charges-and-transfers#reverse-transfers)",
		Attributes: map[string]actionSchema.Attribute{
			"amount": actionSchema.Int64Attribute{
				Optional:    true,
				Description: "Amount, in cents (or local equivalent).",
			},
			"metadata": actionSchema.MapAttribute{
				Optional:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"transfer": actionSchema.StringAttribute{
				Required:    true,
				Description: "ID of the transfer that was reversed.",
			},
			"description": actionSchema.StringAttribute{
				Optional:    true,
				Description: "An arbitrary string which you can attach to a reversal object. This will be unset if you POST an empty value.",
			},
			"refund_application_fee": actionSchema.BoolAttribute{
				Optional:    true,
				Description: "Boolean indicating whether the application fee should be refunded when reversing this transfer. If a full transfer reversal is given, the full application fee will be refunded. Otherwise, the application fee will be refunded with an amount proportional to the amount of the transfer reversed.",
			},
		},
	}
}

func (r *TransferReversalAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func expandTransferReversalCreate(plan TransferReversalResourceModel) (*stripe.TransferReversalCreateParams, error) {
	params := &stripe.TransferReversalCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.RefundApplicationFee.IsNull() && !plan.RefundApplicationFee.IsUnknown() {
		params.RefundApplicationFee = stripe.Bool(plan.RefundApplicationFee.ValueBool())
	}

	return params, nil
}

func (r *TransferReversalAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config TransferReversalResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTransferReversalCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building TransferReversal action params", err.Error())
		return
	}

	if !config.Transfer.IsNull() && !config.Transfer.IsUnknown() {
		if !assignStringToNamedField(params, "Transfer", "ID", config.Transfer.ValueString()) {
			resp.Diagnostics.AddError("Error building TransferReversal create path params", fmt.Sprintf("Failed to assign path parameter %q on %T", "transfer", params))
			return
		}
	}

	obj, err := r.client.V1TransferReversals.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking TransferReversal action", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Created stripe_transfer_reversal %s", obj.ID)})
	}
}
