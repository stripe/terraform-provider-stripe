//
// File generated from our OpenAPI spec
//

package metronome

type V1RotateDeltaStreamSecretResponse struct {
	APIResource
	// The new secret
	Secret string `json:"secret"`
	// When the new secret will start being used
	StartingOn string `json:"starting_on"`
}
