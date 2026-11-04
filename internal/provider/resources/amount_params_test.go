// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
package resources

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	stripe "github.com/stripe/stripe-go/v87"
)

func TestV2AmountParams(t *testing.T) {
	amountType := map[string]attr.Type{"currency": types.StringType, "value": types.Int64Type}
	for _, value := range []int64{0, 1234, 9007199254740993} {
		amount := types.ObjectValueMust(amountType, map[string]attr.Value{
			"currency": types.StringValue("usd"),
			"value":    types.Int64Value(value),
		})
		inbound, err := expandV2MoneyManagementInboundTransferCreate(V2MoneyManagementInboundTransferResourceModel{Amount: amount})
		if err != nil {
			t.Fatal(err)
		}
		outbound, err := expandV2MoneyManagementOutboundPaymentCreate(V2MoneyManagementOutboundPaymentResourceModel{Amount: amount})
		if err != nil {
			t.Fatal(err)
		}
		payout, err := expandV2MoneyManagementPayoutIntentCreate(V2MoneyManagementPayoutIntentResourceModel{Amount: amount})
		if err != nil {
			t.Fatal(err)
		}
		for name, params := range map[string]interface{}{"inbound": inbound, "outbound": outbound, "payout": payout} {
			encoded, err := json.Marshal(params)
			if err != nil {
				t.Fatal(err)
			}
			var decoded struct {
				Amount stripe.Amount `json:"amount"`
			}
			if err := json.Unmarshal(encoded, &decoded); err != nil {
				t.Fatal(err)
			}
			if decoded.Amount.Currency != "usd" || decoded.Amount.Value != value {
				t.Errorf("%s amount = %+v, want usd %d", name, decoded.Amount, value)
			}
		}
		// Imports start without an amount in state and must reconstruct both fields.
		state := V2MoneyManagementInboundTransferResourceModel{Amount: types.ObjectNull(amountType)}
		if err := flattenV2MoneyManagementInboundTransfer(&stripe.V2MoneyManagementInboundTransfer{Amount: *inbound.Amount}, &state); err != nil {
			t.Fatal(err)
		}
		if !state.Amount.Equal(amount) {
			t.Errorf("flattened amount = %v, want %v", state.Amount, amount)
		}
	}
	for _, amount := range []types.Object{types.ObjectNull(amountType), types.ObjectUnknown(amountType)} {
		params, err := expandV2MoneyManagementInboundTransferCreate(V2MoneyManagementInboundTransferResourceModel{Amount: amount})
		if err != nil {
			t.Fatal(err)
		}
		if params.Amount != nil {
			t.Errorf("absent amount sent: %+v", params.Amount)
		}
	}
}
