//
// File generated from our OpenAPI spec
//

package metronome

type SpendTrackerAttributesModel struct {
	APIResource
	// If true, this commit is included in spend trackers with discounted set to DISCOUNTED_ONLY
	CountsAsDiscounted bool `json:"counts_as_discounted"`
}
