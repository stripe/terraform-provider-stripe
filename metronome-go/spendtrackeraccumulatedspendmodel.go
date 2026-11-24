//
// File generated from our OpenAPI spec
//

package metronome

type SpendTrackerAccumulatedSpendModel struct {
	APIResource
	Amount             Number `json:"amount"`
	PeriodEndingBefore string `json:"period_ending_before"`
	PeriodStartingAt   string `json:"period_starting_at"`
}
