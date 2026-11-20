//
// File generated from our OpenAPI spec
//

package metronome

type V1ListRateCardsResponse struct {
	APIResource
	Data     []*RateCardModel `json:"data"`
	NextPage string           `json:"next_page"`
}
