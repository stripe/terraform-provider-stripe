//
// File generated from our OpenAPI spec
//

package metronome

type SeatsUsageAggregateModel struct {
	APIResource
	EndingBefore string `json:"ending_before"`
	StartingOn   string `json:"starting_on"`
	Value        Number `json:"value"`
}
