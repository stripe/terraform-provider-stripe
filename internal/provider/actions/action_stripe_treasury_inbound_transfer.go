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

var _ action.Action = &TreasuryInboundTransferAction{}
var _ action.ActionWithConfigure = &TreasuryInboundTransferAction{}

func NewTreasuryInboundTransferAction() action.Action {
	return &TreasuryInboundTransferAction{}
}

type TreasuryInboundTransferAction struct {
	client *stripe.Client
}

type TreasuryInboundTransferResourceModel struct {
	Amount              types.Int64  `tfsdk:"amount"`
	Currency            types.String `tfsdk:"currency"`
	Description         types.String `tfsdk:"description"`
	FinancialAccount    types.String `tfsdk:"financial_account"`
	Metadata            types.Map    `tfsdk:"metadata"`
	OriginPaymentMethod types.String `tfsdk:"origin_payment_method"`
	StatementDescriptor types.String `tfsdk:"statement_descriptor"`
}

func (r *TreasuryInboundTransferAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_treasury_inbound_transfer"
}

func (r *TreasuryInboundTransferAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionSchema.Schema{
		Description: "Use [InboundTransfers](https://docs.stripe.com/docs/treasury/moving-money/financial-accounts/into/inbound-transfers) to add funds to your [FinancialAccount](https://api.stripe.com#financial_accounts) via a PaymentMethod that is owned by you. The funds will be transferred via an ACH debit.\n\nRelated guide: [Moving money with Treasury using InboundTransfer objects](https://docs.stripe.com/docs/treasury/moving-money/financial-accounts/into/inbound-transfers)",
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
			"financial_account": actionSchema.StringAttribute{
				Required:    true,
				Description: "The FinancialAccount that received the funds.",
			},
			"metadata": actionSchema.MapAttribute{
				Optional:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"origin_payment_method": actionSchema.StringAttribute{
				Required:    true,
				Description: "The origin payment method to be debited for an InboundTransfer.",
			},
			"statement_descriptor": actionSchema.StringAttribute{
				Optional:    true,
				Description: "Statement descriptor shown when funds are debited from the source. Not all payment networks support `statement_descriptor`.",
			},
		},
	}
}

func (r *TreasuryInboundTransferAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func expandTreasuryInboundTransferCreate(plan TreasuryInboundTransferResourceModel) (*stripe.TreasuryInboundTransferCreateParams, error) {
	params := &stripe.TreasuryInboundTransferCreateParams{}

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
	if !plan.OriginPaymentMethod.IsNull() && !plan.OriginPaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "OriginPaymentMethod", "OriginPaymentMethod", plan.OriginPaymentMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "origin_payment_method", params)
		}
	}
	if !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "statement_descriptor", params)
		}
	}

	return params, nil
}

func (r *TreasuryInboundTransferAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config TreasuryInboundTransferResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTreasuryInboundTransferCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building TreasuryInboundTransfer action params", err.Error())
		return
	}

	obj, err := r.client.V1TreasuryInboundTransfers.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking TreasuryInboundTransfer action", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Created stripe_treasury_inbound_transfer %s", obj.ID)})
	}
}
