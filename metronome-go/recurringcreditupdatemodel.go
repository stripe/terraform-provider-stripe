//
// File generated from our OpenAPI spec
//

package metronome

type RecurringCreditUpdateModelRateType string

// List of values that RecurringCreditUpdateModelRateType can take
const (
	RecurringCreditUpdateModelRateTypeLISTRATE   RecurringCreditUpdateModelRateType = "LIST_RATE"
	RecurringCreditUpdateModelRateTypeListRate   RecurringCreditUpdateModelRateType = "list_rate"
	RecurringCreditUpdateModelRateTypeCOMMITRATE RecurringCreditUpdateModelRateType = "COMMIT_RATE"
	RecurringCreditUpdateModelRateTypeCommitRate RecurringCreditUpdateModelRateType = "commit_rate"
)

type RecurringCreditUpdateModelAccessAmount struct {
	Quantity  Number `json:"quantity,omitempty"`
	UnitPrice Number `json:"unit_price,omitempty"`
}

// Rounding configuration for prorated recurring credit amounts.
type RecurringCreditUpdateModelProrationRounding struct {
	Access *ProrationRoundingConfigV2Model `json:"access,omitempty"`
}
type RecurringCreditUpdateModel struct {
	APIResource
	AccessAmount *RecurringCreditUpdateModelAccessAmount `json:"access_amount,omitempty"`
	EndingBefore string                                  `json:"ending_before,omitempty"`
	ID           string                                  `json:"id"`
	// Rounding configuration for prorated recurring credit amounts.
	ProrationRounding *RecurringCreditUpdateModelProrationRounding `json:"proration_rounding,omitempty"`
	RateType          RecurringCreditUpdateModelRateType           `json:"rate_type,omitempty"`
}
