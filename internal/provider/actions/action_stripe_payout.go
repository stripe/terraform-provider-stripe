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

var _ action.Action = &PayoutAction{}
var _ action.ActionWithConfigure = &PayoutAction{}

func NewPayoutAction() action.Action {
	return &PayoutAction{}
}

type PayoutAction struct {
	client *stripe.Client
}

type PayoutResourceModel struct {
	Amount              types.Int64  `tfsdk:"amount"`
	Currency            types.String `tfsdk:"currency"`
	Description         types.String `tfsdk:"description"`
	Destination         types.String `tfsdk:"destination"`
	Metadata            types.Map    `tfsdk:"metadata"`
	Method              types.String `tfsdk:"method"`
	PayoutMethod        types.String `tfsdk:"payout_method"`
	SourceType          types.String `tfsdk:"source_type"`
	StatementDescriptor types.String `tfsdk:"statement_descriptor"`
}

func (r *PayoutAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_payout"
}

func (r *PayoutAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionSchema.Schema{
		Description: "A `Payout` object is created when you receive funds from Stripe, or when you\ninitiate a payout to either a bank account or debit card of a [connected\nStripe account](/docs/connect/bank-debit-card-payouts). You can retrieve individual payouts,\nand list all payouts. Payouts are made on [varying\nschedules](/docs/connect/manage-payout-schedule), depending on your country and\nindustry.\n\nRelated guide: [Receiving payouts](https://docs.stripe.com/payouts)",
		Attributes: map[string]actionSchema.Attribute{
			"amount": actionSchema.Int64Attribute{
				Required:    true,
				Description: "The amount (in cents (or local equivalent)) that transfers to your bank account or debit card.",
			},
			"currency": actionSchema.StringAttribute{
				Required:    true,
				Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
			},
			"description": actionSchema.StringAttribute{
				Optional:    true,
				Description: "An arbitrary string attached to the object. Often useful for displaying to users.",
			},
			"destination": actionSchema.StringAttribute{
				Optional:    true,
				Description: "ID of the bank account or card the payout is sent to.",
			},
			"metadata": actionSchema.MapAttribute{
				Optional:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"method": actionSchema.StringAttribute{
				Optional:    true,
				Description: "The method used to send this payout, which can be `standard` or `instant`. `instant` is supported for payouts to debit cards and bank accounts in certain countries. Learn more about [bank support for Instant Payouts](https://stripe.com/docs/payouts/instant-payouts-banks).",
			},
			"payout_method": actionSchema.StringAttribute{
				Optional:    true,
				Description: "ID of the v2 FinancialAccount the funds are sent to.",
			},
			"source_type": actionSchema.StringAttribute{
				Optional:    true,
				Description: "The source balance this payout came from, which can be one of the following: `card`, `fpx`, or `bank_account`.",
			},
			"statement_descriptor": actionSchema.StringAttribute{
				Optional:    true,
				Description: "Extra information about a payout that displays on the user's bank statement.",
			},
		},
	}
}

func (r *PayoutAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func expandPayoutCreate(plan PayoutResourceModel) (*stripe.PayoutCreateParams, error) {
	params := &stripe.PayoutCreateParams{}

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
	if !plan.Destination.IsNull() && !plan.Destination.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DestinationID", "Destination", plan.Destination.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "destination", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Method.IsNull() && !plan.Method.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Method", "Method", plan.Method.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "method", params)
		}
	}
	if !plan.PayoutMethod.IsNull() && !plan.PayoutMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PayoutMethod", "PayoutMethod", plan.PayoutMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "payout_method", params)
		}
	}
	if !plan.SourceType.IsNull() && !plan.SourceType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SourceType", "SourceType", plan.SourceType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "source_type", params)
		}
	}
	if !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "statement_descriptor", params)
		}
	}

	return params, nil
}

func (r *PayoutAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config PayoutResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandPayoutCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building Payout action params", err.Error())
		return
	}

	obj, err := r.client.V1Payouts.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking Payout action", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Created stripe_payout %s", obj.ID)})
	}
}
