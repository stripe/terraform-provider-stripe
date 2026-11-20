//
// File generated from our OpenAPI spec
//

package metronome

type ThresholdBalanceSpecifierModel struct {
	APIResource
	// If any of the exclude specifier is met, the balance is not considered when evaluating threshold billing
	Exclude []*ExcludeThresholdBalanceSpecifierModel `json:"exclude"`
}
