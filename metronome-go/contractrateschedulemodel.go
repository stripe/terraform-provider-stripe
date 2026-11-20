//
// File generated from our OpenAPI spec
//

package metronome

type ContractRateScheduleModelBillingFrequency string

// List of values that ContractRateScheduleModelBillingFrequency can take
const (
	ContractRateScheduleModelBillingFrequencyMONTHLY   ContractRateScheduleModelBillingFrequency = "MONTHLY"
	ContractRateScheduleModelBillingFrequencyQUARTERLY ContractRateScheduleModelBillingFrequency = "QUARTERLY"
	ContractRateScheduleModelBillingFrequencyANNUAL    ContractRateScheduleModelBillingFrequency = "ANNUAL"
	ContractRateScheduleModelBillingFrequencyWEEKLY    ContractRateScheduleModelBillingFrequency = "WEEKLY"
)

type ContractRateScheduleModel struct {
	APIResource
	BillingFrequency    ContractRateScheduleModelBillingFrequency `json:"billing_frequency,omitempty"`
	CommitRate          *CommitRateModel                          `json:"commit_rate,omitempty"`
	EndingBefore        string                                    `json:"ending_before,omitempty"`
	Entitled            bool                                      `json:"entitled"`
	ListRate            *RateModel                                `json:"list_rate"`
	OverrideRate        *RateModel                                `json:"override_rate,omitempty"`
	PricingGroupValues  map[string]string                         `json:"pricing_group_values,omitempty"`
	ProductCustomFields map[string]string                         `json:"product_custom_fields"`
	ProductID           string                                    `json:"product_id"`
	ProductName         string                                    `json:"product_name"`
	ProductTags         []string                                  `json:"product_tags"`
	RateCardID          string                                    `json:"rate_card_id"`
	StartingAt          string                                    `json:"starting_at"`
}
