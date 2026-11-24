//
// File generated from our OpenAPI spec
//

package metronome

type V1ListProductsResponse struct {
	APIResource
	Data     []*ProductListItemModel `json:"data"`
	NextPage string                  `json:"next_page"`
}
