//
// File generated from our OpenAPI spec
//

package metronome

// Update the discount cap. Set to null to remove an existing cap.
type UpdatePrepaidBalanceThresholdConfigurationModelDiscountConfigurationCap struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount Number `json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias string `json:"spend_tracker_alias"`
}
type UpdatePrepaidBalanceThresholdConfigurationModelDiscountConfiguration struct {
	// Update the discount cap. Set to null to remove an existing cap.
	Cap *UpdatePrepaidBalanceThresholdConfigurationModelDiscountConfigurationCap `json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. Set to null to remove the discount fraction. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	//
	PaymentFraction Number `json:"payment_fraction,omitempty"`
}
type UpdatePrepaidBalanceThresholdConfigurationModel struct {
	APIResource
	Commit *UpdatePrepaidBalanceThresholdCommitModel `json:"commit,omitempty"`
	// If provided, the threshold, recharge-to amount, and the resulting threshold commit amount will be in terms of this credit type instead of the fiat currency.
	CustomCreditTypeID    string                                                                `json:"custom_credit_type_id,omitempty"`
	DiscountConfiguration *UpdatePrepaidBalanceThresholdConfigurationModelDiscountConfiguration `json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         bool                      `json:"is_enabled,omitempty"`
	PaymentGateConfig *PaymentGateConfigV2Model `json:"payment_gate_config,omitempty"`
	// Specify the amount the balance should be recharged to.
	RechargeToAmount Number `json:"recharge_to_amount,omitempty"`
	// Specify the threshold amount for the contract. Each time the contract's balance lowers to this amount, a threshold charge will be initiated.
	ThresholdAmount Number `json:"threshold_amount,omitempty"`
	// Determines which balances are excluded from remaining balance calculation for threshold billing.
	ThresholdBalanceSpecifiers []*ThresholdBalanceSpecifierV2Model `json:"threshold_balance_specifiers,omitempty"`
}
