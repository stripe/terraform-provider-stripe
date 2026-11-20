//
// File generated from our OpenAPI spec
//

package metronome

type V1ListPackagesResponse struct {
	APIResource
	Data     []*PackageModel `json:"data"`
	NextPage string          `json:"next_page"`
}
