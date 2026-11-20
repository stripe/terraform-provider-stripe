//
// File generated from our OpenAPI spec
//

package metronome

type RecurringCommitUpdateModelRateType string

// List of values that RecurringCommitUpdateModelRateType can take
const (
	RecurringCommitUpdateModelRateTypeLISTRATE   RecurringCommitUpdateModelRateType = "LIST_RATE"
	RecurringCommitUpdateModelRateTypeListRate   RecurringCommitUpdateModelRateType = "list_rate"
	RecurringCommitUpdateModelRateTypeCOMMITRATE RecurringCommitUpdateModelRateType = "COMMIT_RATE"
	RecurringCommitUpdateModelRateTypeCommitRate RecurringCommitUpdateModelRateType = "commit_rate"
)

type RecurringCommitUpdateModelAccessAmount struct {
	Quantity  Number `json:"quantity,omitempty"`
	UnitPrice Number `json:"unit_price,omitempty"`
}
type RecurringCommitUpdateModelInvoiceAmount struct {
	Quantity  Number `json:"quantity,omitempty"`
	UnitPrice Number `json:"unit_price,omitempty"`
}

// Rounding configuration for prorated recurring commit amounts.
type RecurringCommitUpdateModelProrationRounding struct {
	Access  *ProrationRoundingConfigV2Model `json:"access,omitempty"`
	Invoice *ProrationRoundingConfigV2Model `json:"invoice,omitempty"`
}
type RecurringCommitUpdateModel struct {
	APIResource
	AccessAmount  *RecurringCommitUpdateModelAccessAmount  `json:"access_amount,omitempty"`
	EndingBefore  string                                   `json:"ending_before,omitempty"`
	ID            string                                   `json:"id"`
	InvoiceAmount *RecurringCommitUpdateModelInvoiceAmount `json:"invoice_amount,omitempty"`
	// Rounding configuration for prorated recurring commit amounts.
	ProrationRounding *RecurringCommitUpdateModelProrationRounding `json:"proration_rounding,omitempty"`
	RateType          RecurringCommitUpdateModelRateType           `json:"rate_type,omitempty"`
}
