//
// File generated from our OpenAPI spec
//

package metronome

type UsageStatementScheduleModelFrequency string

// List of values that UsageStatementScheduleModelFrequency can take
const (
	UsageStatementScheduleModelFrequencyMONTHLY   UsageStatementScheduleModelFrequency = "MONTHLY"
	UsageStatementScheduleModelFrequencyMonthly   UsageStatementScheduleModelFrequency = "monthly"
	UsageStatementScheduleModelFrequencyQUARTERLY UsageStatementScheduleModelFrequency = "QUARTERLY"
	UsageStatementScheduleModelFrequencyQuarterly UsageStatementScheduleModelFrequency = "quarterly"
	UsageStatementScheduleModelFrequencyANNUAL    UsageStatementScheduleModelFrequency = "ANNUAL"
	UsageStatementScheduleModelFrequencyAnnual    UsageStatementScheduleModelFrequency = "annual"
	UsageStatementScheduleModelFrequencyWEEKLY    UsageStatementScheduleModelFrequency = "WEEKLY"
	UsageStatementScheduleModelFrequencyWeekly    UsageStatementScheduleModelFrequency = "weekly"
)

type UsageStatementScheduleModel struct {
	APIResource
	// Contract usage statements follow a selected cadence based on this date.
	BillingAnchorDate string                               `json:"billing_anchor_date"`
	Frequency         UsageStatementScheduleModelFrequency `json:"frequency"`
}
