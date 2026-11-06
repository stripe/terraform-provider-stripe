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

var _ action.Action = &TreasuryOutboundTransferAction{}
var _ action.ActionWithConfigure = &TreasuryOutboundTransferAction{}

func NewTreasuryOutboundTransferAction() action.Action {
	return &TreasuryOutboundTransferAction{}
}

type TreasuryOutboundTransferAction struct {
	client *stripe.Client
}

type TreasuryOutboundTransferResourceModel struct {
	Amount                          types.Int64  `tfsdk:"amount"`
	Currency                        types.String `tfsdk:"currency"`
	Description                     types.String `tfsdk:"description"`
	DestinationPaymentMethod        types.String `tfsdk:"destination_payment_method"`
	FinancialAccount                types.String `tfsdk:"financial_account"`
	Metadata                        types.Map    `tfsdk:"metadata"`
	StatementDescriptor             types.String `tfsdk:"statement_descriptor"`
	DestinationPaymentMethodData    types.Object `tfsdk:"destination_payment_method_data"`
	DestinationPaymentMethodOptions types.Object `tfsdk:"destination_payment_method_options"`
}

func (r *TreasuryOutboundTransferAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_treasury_outbound_transfer"
}

func (r *TreasuryOutboundTransferAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionSchema.Schema{
		Description: "Use [OutboundTransfers](https://docs.stripe.com/docs/treasury/moving-money/financial-accounts/out-of/outbound-transfers) to transfer funds from a [FinancialAccount](https://api.stripe.com#financial_accounts) to a PaymentMethod belonging to the same entity. To send funds to a different party, use [OutboundPayments](https://api.stripe.com#outbound_payments) instead. You can send funds over ACH rails or through a domestic wire transfer to a user's own external bank account.\n\nSimulate OutboundTransfer state changes with the `/v1/test_helpers/treasury/outbound_transfers` endpoints. These methods can only be called on test mode objects.\n\nRelated guide: [Moving money with Treasury using OutboundTransfer objects](https://docs.stripe.com/docs/treasury/moving-money/financial-accounts/out-of/outbound-transfers)",
		Attributes: map[string]actionSchema.Attribute{
			"amount": actionSchema.Int64Attribute{
				Required:    true,
				Description: "Amount (in cents) transferred.",
			},
			"currency": actionSchema.StringAttribute{
				Required:    true,
				Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
			},
			"description": actionSchema.StringAttribute{
				Optional:    true,
				Description: "An arbitrary string attached to the object. Often useful for displaying to users.",
			},
			"destination_payment_method": actionSchema.StringAttribute{
				Optional:    true,
				Description: "The PaymentMethod used as the payment instrument for an OutboundTransfer.",
			},
			"financial_account": actionSchema.StringAttribute{
				Required:    true,
				Description: "The FinancialAccount that funds were pulled from.",
			},
			"metadata": actionSchema.MapAttribute{
				Optional:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"statement_descriptor": actionSchema.StringAttribute{
				Optional:    true,
				Description: "Information about the OutboundTransfer to be sent to the recipient account.",
			},
			"destination_payment_method_data": actionSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "Hash used to generate the PaymentMethod to be used for this OutboundTransfer. Exclusive with `destination_payment_method`.",
				Attributes: map[string]actionSchema.Attribute{
					"financial_account": actionSchema.StringAttribute{
						Optional:    true,
						Description: "Required if type is set to `financial_account`. The FinancialAccount ID to send funds to.",
					},
					"type": actionSchema.StringAttribute{
						Required:    true,
						Description: "The type of the destination.",
					},
				},
			},
			"destination_payment_method_options": actionSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "Hash describing payment method configuration details.",
				Attributes: map[string]actionSchema.Attribute{
					"us_bank_account": actionSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Optional fields for `us_bank_account`.",
						Attributes: map[string]actionSchema.Attribute{
							"network": actionSchema.StringAttribute{
								Optional:    true,
								Description: "Specifies the network rails to be used. If not set, will default to the PaymentMethod's preferred network. See the [docs](https://docs.stripe.com/treasury/money-movement/timelines) to learn more about money movement timelines for each network type.",
							},
						},
					},
				},
			},
		},
	}
}

func (r *TreasuryOutboundTransferAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func expandTreasuryOutboundTransferCreate(plan TreasuryOutboundTransferResourceModel) (*stripe.TreasuryOutboundTransferCreateParams, error) {
	params := &stripe.TreasuryOutboundTransferCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.DestinationPaymentMethod.IsNull() && !plan.DestinationPaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DestinationPaymentMethod", "DestinationPaymentMethod", plan.DestinationPaymentMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "destination_payment_method", params)
		}
	}
	if !plan.FinancialAccount.IsNull() && !plan.FinancialAccount.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FinancialAccount", "FinancialAccount", plan.FinancialAccount.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "financial_account", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "statement_descriptor", params)
		}
	}
	if !plan.DestinationPaymentMethodData.IsNull() && !plan.DestinationPaymentMethodData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DestinationPaymentMethodData", plan.DestinationPaymentMethodData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "destination_payment_method_data", params)
		}
	}
	if !plan.DestinationPaymentMethodOptions.IsNull() && !plan.DestinationPaymentMethodOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DestinationPaymentMethodOptions", plan.DestinationPaymentMethodOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "destination_payment_method_options", params)
		}
	}

	return params, nil
}

func (r *TreasuryOutboundTransferAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config TreasuryOutboundTransferResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTreasuryOutboundTransferCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building TreasuryOutboundTransfer action params", err.Error())
		return
	}

	obj, err := r.client.V1TreasuryOutboundTransfers.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking TreasuryOutboundTransfer action", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Created stripe_treasury_outbound_transfer %s", obj.ID)})
	}
}
