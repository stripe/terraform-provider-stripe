//
// File generated from our OpenAPI spec
//

package metronome

type V1GetNetBalanceResponseData struct {
	// The combined net balance that the customer has access to use at this moment across all pertinent commits and credits.
	Balance Number `json:"balance"`
	// This ID identifies the credit type for the balance. The credit type can be fiat or a custom pricing unit. Quantity-based balances return the null credit type UUID.
	CreditTypeID string `json:"credit_type_id"`
}
type V1GetNetBalanceResponse struct {
	APIResource
	Data *V1GetNetBalanceResponseData `json:"data"`
}
