//
// File generated from our OpenAPI spec
//

package metronome

type UsageStatementScheduleTemplateModelFrequency string

// List of values that UsageStatementScheduleTemplateModelFrequency can take
const (
	UsageStatementScheduleTemplateModelFrequencyMONTHLY   UsageStatementScheduleTemplateModelFrequency = "MONTHLY"
	UsageStatementScheduleTemplateModelFrequencyQUARTERLY UsageStatementScheduleTemplateModelFrequency = "QUARTERLY"
	UsageStatementScheduleTemplateModelFrequencyANNUAL    UsageStatementScheduleTemplateModelFrequency = "ANNUAL"
	UsageStatementScheduleTemplateModelFrequencyWEEKLY    UsageStatementScheduleTemplateModelFrequency = "WEEKLY"
)

type UsageStatementScheduleTemplateModelDay string

// List of values that UsageStatementScheduleTemplateModelDay can take
const (
	UsageStatementScheduleTemplateModelDayFIRSTOFMONTH  UsageStatementScheduleTemplateModelDay = "FIRST_OF_MONTH"
	UsageStatementScheduleTemplateModelDayCONTRACTSTART UsageStatementScheduleTemplateModelDay = "CONTRACT_START"
)

type UsageStatementScheduleTemplateModel struct {
	APIResource
	Day       UsageStatementScheduleTemplateModelDay       `json:"day,omitempty"`
	Frequency UsageStatementScheduleTemplateModelFrequency `json:"frequency"`
}
