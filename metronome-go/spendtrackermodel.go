//
// File generated from our OpenAPI spec
//

package metronome

type SpendTrackerModelResetFrequency string

// List of values that SpendTrackerModelResetFrequency can take
const (
	SpendTrackerModelResetFrequencyBILLINGPERIOD SpendTrackerModelResetFrequency = "BILLING_PERIOD"
)

type SpendTrackerModel struct {
	APIResource
	AccumulatedSpend *SpendTrackerAccumulatedSpendModel `json:"accumulated_spend,omitempty"`
	// Human-readable identifier, unique per contract.
	Alias                     string                                       `json:"alias"`
	ApplicableSpendSpecifiers []*SpendTrackerApplicableSpendSpecifierModel `json:"applicable_spend_specifiers"`
	CreditTypeID              string                                       `json:"credit_type_id"`
	ResetFrequency            SpendTrackerModelResetFrequency              `json:"reset_frequency"`
}
