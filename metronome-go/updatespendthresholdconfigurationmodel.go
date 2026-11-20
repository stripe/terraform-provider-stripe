//
// File generated from our OpenAPI spec
//

package metronome

// Update the discount cap. Set to null to remove an existing cap.
type UpdateSpendThresholdConfigurationModelDiscountConfigurationCap struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount Number `json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias string `json:"spend_tracker_alias"`
}
type UpdateSpendThresholdConfigurationModelDiscountConfiguration struct {
	// Update the discount cap. Set to null to remove an existing cap.
	Cap *UpdateSpendThresholdConfigurationModelDiscountConfigurationCap `json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. Set to null to remove the discount fraction. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	//
	PaymentFraction Number `json:"payment_fraction,omitempty"`
}
type UpdateSpendThresholdConfigurationModel struct {
	APIResource
	Commit                *UpdateSpendThresholdCommitModel                             `json:"commit,omitempty"`
	DiscountConfiguration *UpdateSpendThresholdConfigurationModelDiscountConfiguration `json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         bool                      `json:"is_enabled,omitempty"`
	PaymentGateConfig *PaymentGateConfigV2Model `json:"payment_gate_config,omitempty"`
	// Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.
	ThresholdAmount Number `json:"threshold_amount,omitempty"`
}
