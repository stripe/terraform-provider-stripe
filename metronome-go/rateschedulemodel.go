//
// File generated from our OpenAPI spec
//

package metronome

type RateScheduleModelBillingFrequency string

// List of values that RateScheduleModelBillingFrequency can take
const (
	RateScheduleModelBillingFrequencyMONTHLY   RateScheduleModelBillingFrequency = "MONTHLY"
	RateScheduleModelBillingFrequencyQUARTERLY RateScheduleModelBillingFrequency = "QUARTERLY"
	RateScheduleModelBillingFrequencyANNUAL    RateScheduleModelBillingFrequency = "ANNUAL"
	RateScheduleModelBillingFrequencyWEEKLY    RateScheduleModelBillingFrequency = "WEEKLY"
)

type RateScheduleModel struct {
	APIResource
	BillingFrequency    RateScheduleModelBillingFrequency `json:"billing_frequency,omitempty"`
	CommitRate          *CommitRateModel                  `json:"commit_rate,omitempty"`
	EndingBefore        string                            `json:"ending_before,omitempty"`
	Entitled            bool                              `json:"entitled"`
	PricingGroupValues  map[string]string                 `json:"pricing_group_values,omitempty"`
	ProductCustomFields map[string]string                 `json:"product_custom_fields"`
	ProductID           string                            `json:"product_id"`
	ProductName         string                            `json:"product_name"`
	ProductTags         []string                          `json:"product_tags"`
	Rate                *RateModel                        `json:"rate"`
	StartingAt          string                            `json:"starting_at"`
}
