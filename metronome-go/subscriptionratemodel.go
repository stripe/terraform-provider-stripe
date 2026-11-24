//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionRateModelBillingFrequency string

// List of values that SubscriptionRateModelBillingFrequency can take
const (
	SubscriptionRateModelBillingFrequencyMONTHLY   SubscriptionRateModelBillingFrequency = "MONTHLY"
	SubscriptionRateModelBillingFrequencyQUARTERLY SubscriptionRateModelBillingFrequency = "QUARTERLY"
	SubscriptionRateModelBillingFrequencyANNUAL    SubscriptionRateModelBillingFrequency = "ANNUAL"
	SubscriptionRateModelBillingFrequencyWEEKLY    SubscriptionRateModelBillingFrequency = "WEEKLY"
	SubscriptionRateModelBillingFrequencyMonthly   SubscriptionRateModelBillingFrequency = "monthly"
	SubscriptionRateModelBillingFrequencyQuarterly SubscriptionRateModelBillingFrequency = "quarterly"
	SubscriptionRateModelBillingFrequencyAnnual    SubscriptionRateModelBillingFrequency = "annual"
	SubscriptionRateModelBillingFrequencyWeekly    SubscriptionRateModelBillingFrequency = "weekly"
)

type SubscriptionRateModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type SubscriptionRateModel struct {
	APIResource
	BillingFrequency SubscriptionRateModelBillingFrequency `json:"billing_frequency"`
	Product          *SubscriptionRateModelProduct         `json:"product"`
}
