//
// File generated from our OpenAPI spec
//

package metronome

type SpendThresholdConfigurationV2Model struct {
	APIResource
	Commit                *SpendThresholdCommitModel  `json:"commit"`
	DiscountConfiguration *DiscountConfigurationModel `json:"discount_configuration,omitempty"`
	// When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.
	IsEnabled         bool                      `json:"is_enabled"`
	PaymentGateConfig *PaymentGateConfigV2Model `json:"payment_gate_config"`
	// Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.
	ThresholdAmount Number `json:"threshold_amount"`
}
