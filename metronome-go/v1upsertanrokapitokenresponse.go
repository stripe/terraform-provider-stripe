//
// File generated from our OpenAPI spec
//

package metronome

type V1UpsertAnrokAPITokenResponseData struct {
	// Whether the update was successful.
	Success bool `json:"success"`
}
type V1UpsertAnrokAPITokenResponse struct {
	APIResource
	Data *V1UpsertAnrokAPITokenResponseData `json:"data"`
}
