//
// File generated from our OpenAPI spec
//

package metronome

type BaseUsageFilterModel struct {
	APIResource
	GroupKey    string   `json:"group_key"`
	GroupValues []string `json:"group_values"`
	StartingAt  string   `json:"starting_at,omitempty"`
}
