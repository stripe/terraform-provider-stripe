//
// File generated from our OpenAPI spec
//

package metronome

// A unified API for updating prepaid balance threshold configs for clients who are migrating customers from Metronome 1.0 and 2.0. Customer can only have one active config across all Plans or Contracts.
type V1ThresholdBillingUpdateActiveRechargeConfigParams struct {
	Params     `form:"*"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
	Enabled    *bool   `form:"enabled" json:"enabled,omitempty"`
	// the threshold at which to recharge
	RechargeThreshold *Number `form:"recharge_threshold" json:"recharge_threshold,omitempty"`
	// the amount to recharge to
	RechargeToAmount *Number `form:"recharge_to_amount" json:"recharge_to_amount,omitempty"`
}
