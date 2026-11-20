//
// File generated from our OpenAPI spec
//

package metronome

type V1EmbeddableDashboardResponseData struct {
	URL string `json:"url,omitempty"`
}
type V1EmbeddableDashboardResponse struct {
	APIResource
	Data *V1EmbeddableDashboardResponseData `json:"data"`
}
