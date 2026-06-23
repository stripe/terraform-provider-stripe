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

var _ action.Action = &TransferAction{}
var _ action.ActionWithConfigure = &TransferAction{}

func NewTransferAction() action.Action {
	return &TransferAction{}
}

type TransferAction struct {
	client *stripe.Client
}

type TransferResourceModel struct {
	Amount            types.Int64  `tfsdk:"amount"`
	Currency          types.String `tfsdk:"currency"`
	Description       types.String `tfsdk:"description"`
	Destination       types.String `tfsdk:"destination"`
	Metadata          types.Map    `tfsdk:"metadata"`
	SourceTransaction types.String `tfsdk:"source_transaction"`
	SourceType        types.String `tfsdk:"source_type"`
	TransferGroup     types.String `tfsdk:"transfer_group"`
}

func (r *TransferAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transfer"
}

func (r *TransferAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionSchema.Schema{
		Description: "A `Transfer` object is created when you move funds between Stripe accounts as\npart of Connect.\n\nBefore April 6, 2017, transfers also represented movement of funds from a\nStripe account to a card or bank account. This behavior has since been split\nout into a [Payout](https://api.stripe.com#payout_object) object, with corresponding payout endpoints. For more\ninformation, read about the\n[transfer/payout split](https://docs.stripe.com/transfer-payout-split).\n\nRelated guide: [Creating separate charges and transfers](https://docs.stripe.com/connect/separate-charges-and-transfers)",
		Attributes: map[string]actionSchema.Attribute{
			"amount": actionSchema.Int64Attribute{
				Optional:    true,
				Description: "Amount in cents (or local equivalent) to be transferred.",
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
				Required:    true,
				Description: "ID of the Stripe account the transfer was sent to.",
			},
			"metadata": actionSchema.MapAttribute{
				Optional:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"source_transaction": actionSchema.StringAttribute{
				Optional:    true,
				Description: "ID of the charge that was used to fund the transfer. If null, the transfer was funded from the available balance.",
			},
			"source_type": actionSchema.StringAttribute{
				Optional:    true,
				Description: "The source balance this transfer came from. One of `card`, `fpx`, or `bank_account`.",
			},
			"transfer_group": actionSchema.StringAttribute{
				Optional:    true,
				Description: "A string that identifies this transaction as part of a group. See the [Connect documentation](https://docs.stripe.com/connect/separate-charges-and-transfers#transfer-options) for details.",
			},
		},
	}
}

func (r *TransferAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
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

func expandTransferCreate(plan TransferResourceModel) (*stripe.TransferCreateParams, error) {
	params := &stripe.TransferCreateParams{}

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
	if !plan.SourceTransaction.IsNull() && !plan.SourceTransaction.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SourceTransactionID", "SourceTransaction", plan.SourceTransaction.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "source_transaction", params)
		}
	}
	if !plan.SourceType.IsNull() && !plan.SourceType.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SourceType", "SourceType", plan.SourceType.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "source_type", params)
		}
	}
	if !plan.TransferGroup.IsNull() && !plan.TransferGroup.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TransferGroup", "TransferGroup", plan.TransferGroup.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "transfer_group", params)
		}
	}

	return params, nil
}

func (r *TransferAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config TransferResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTransferCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building Transfer action params", err.Error())
		return
	}

	obj, err := r.client.V1Transfers.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking Transfer action", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Created stripe_transfer %s", obj.ID)})
	}
}
