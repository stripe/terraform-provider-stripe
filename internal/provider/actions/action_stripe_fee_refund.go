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

var _ action.Action = &FeeRefundAction{}
var _ action.ActionWithConfigure = &FeeRefundAction{}

func NewFeeRefundAction() action.Action {
	return &FeeRefundAction{}
}

type FeeRefundAction struct {
	client *stripe.Client
}

type FeeRefundResourceModel struct {
	Amount   types.Int64  `tfsdk:"amount"`
	Fee      types.String `tfsdk:"fee"`
	Metadata types.Map    `tfsdk:"metadata"`
}

func (r *FeeRefundAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fee_refund"
}

func (r *FeeRefundAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionSchema.Schema{
		Description: "`Application Fee Refund` objects allow you to refund an application fee that\nhas previously been created but not yet refunded. Funds will be refunded to\nthe Stripe account from which the fee was originally collected.\n\nRelated guide: [Refunding application fees](https://docs.stripe.com/connect/destination-charges#refunding-app-fee)",
		Attributes: map[string]actionSchema.Attribute{
			"amount": actionSchema.Int64Attribute{
				Optional:    true,
				Description: "Amount, in cents (or local equivalent).",
			},
			"fee": actionSchema.StringAttribute{
				Required:    true,
				Description: "ID of the application fee that was refunded.",
			},
			"metadata": actionSchema.MapAttribute{
				Optional:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
		},
	}
}

func (r *FeeRefundAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func expandFeeRefundCreate(plan FeeRefundResourceModel) (*stripe.FeeRefundCreateParams, error) {
	params := &stripe.FeeRefundCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}

	return params, nil
}

func (r *FeeRefundAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config FeeRefundResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandFeeRefundCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building FeeRefund action params", err.Error())
		return
	}

	if !config.Fee.IsNull() && !config.Fee.IsUnknown() {
		if !assignStringToNamedField(params, "Fee", "ID", config.Fee.ValueString()) {
			resp.Diagnostics.AddError("Error building FeeRefund create path params", fmt.Sprintf("Failed to assign path parameter %q on %T", "fee", params))
			return
		}
	}

	obj, err := r.client.V1FeeRefunds.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking FeeRefund action", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Created stripe_fee_refund %s", obj.ID)})
	}
}
