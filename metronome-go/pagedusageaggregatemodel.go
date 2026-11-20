//
// File generated from our OpenAPI spec
//

package metronome

type PagedUsageAggregateModel struct {
	APIResource
	EndingBefore string `json:"ending_before"`
	// Map of group key to their value for this usage aggregate.
	// For simple group keys, this should be a single key e.g. `{"region": "US-East"}`
	// For compound group keys, this should contain the values of each group key that forms the compound e.g. `{"region": "US-East", "team": "engineering"}`
	//
	Group map[string]string `json:"group,omitempty"`
	// Use `group` instead. The group key for single-key grouping.
	GroupKey string `json:"group_key"`
	// Use `group` instead. The group value for single-key grouping.
	GroupValue string `json:"group_value"`
	StartingOn string `json:"starting_on"`
	Value      Number `json:"value"`
}
