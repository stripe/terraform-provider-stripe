//
// File generated from our OpenAPI spec
//

package metronome

type SpendTrackerTemplateModelResetFrequency string

// List of values that SpendTrackerTemplateModelResetFrequency can take
const (
	SpendTrackerTemplateModelResetFrequencyBILLINGPERIOD SpendTrackerTemplateModelResetFrequency = "BILLING_PERIOD"
)

type SpendTrackerTemplateModel struct {
	APIResource
	// Human-readable identifier, unique per contract.
	Alias                     string                                       `json:"alias"`
	ApplicableSpendSpecifiers []*SpendTrackerApplicableSpendSpecifierModel `json:"applicable_spend_specifiers"`
	CreditTypeID              string                                       `json:"credit_type_id"`
	ResetFrequency            SpendTrackerTemplateModelResetFrequency      `json:"reset_frequency"`
}
