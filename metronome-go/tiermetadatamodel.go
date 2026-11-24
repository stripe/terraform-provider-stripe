//
// File generated from our OpenAPI spec
//

package metronome

type TierMetadataModel struct {
	APIResource
	Level      Number `json:"level"`
	Size       string `json:"size,omitempty"`
	StartingAt string `json:"starting_at"`
}
