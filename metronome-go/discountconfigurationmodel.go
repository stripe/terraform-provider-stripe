//
// File generated from our OpenAPI spec
//

package metronome

// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
type DiscountConfigurationModelCap struct {
	// Accumulated spend ceiling above which the discount stops applying.
	Amount Number `json:"amount"`
	// Alias of the spend tracker this cap is measured against.
	SpendTrackerAlias string `json:"spend_tracker_alias"`
}
type DiscountConfigurationModel struct {
	APIResource
	// If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.
	Cap *DiscountConfigurationModelCap `json:"cap,omitempty"`
	// The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).
	PaymentFraction Number `json:"payment_fraction"`
}
