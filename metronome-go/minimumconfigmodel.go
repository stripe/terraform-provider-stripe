//
// File generated from our OpenAPI spec
//

package metronome

// Only set for TIERED_PERCENTAGE or PERCENTAGE rate_type. Any commit-specific overrides will not apply if there is a minimum set on the rate/applied override.
type MinimumConfigModel struct {
	APIResource
	Minimum Number `json:"minimum"`
}
