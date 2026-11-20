//
// File generated from our OpenAPI spec
//

package metronome

type UsageFilterV2Model struct {
	APIResource
	// This will match contract ending_before value if usage filter is active until the end of the contract. It will be undefined if the contract is open-ended.
	EndingBefore string   `json:"ending_before,omitempty"`
	GroupKey     string   `json:"group_key"`
	GroupValues  []string `json:"group_values"`
	// This will match contract starting_at value if usage filter is active from the beginning of the contract.
	StartingAt string `json:"starting_at"`
}
