//
// File generated from our OpenAPI spec
//

package metronome

type PrepaidBalanceThresholdConfigurationModel struct {
	APIResource
	Commit *PrepaidBalanceThresholdCommitModel `json:"commit"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID    string                      `json:"custom_credit_type_id,omitempty"`
	DiscountConfiguration *DiscountConfigurationModel `json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         bool                    `json:"is_enabled"`
	PaymentGateConfig *PaymentGateConfigModel `json:"payment_gate_config"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount Number `json:"recharge_to_amount"`
	// Specify the threshold amount for the contract. Each time the contract's prepaid balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount Number `json:"threshold_amount"`
	// Determines which balances are excluded from remaining balance calculation for threshold billing.
	ThresholdBalanceSpecifiers []*ThresholdBalanceSpecifierModel `json:"threshold_balance_specifiers,omitempty"`
}
