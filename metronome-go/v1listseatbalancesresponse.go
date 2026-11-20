//
// File generated from our OpenAPI spec
//

package metronome

type V1ListSeatBalancesResponsePagination struct {
	// Token to retrieve the next page of results. Null if no more pages available
	NextPage string `json:"next_page,omitempty"`
	// Number of seats available to fetch in the next page
	SeatsAvailableForNextPage Number `json:"seats_available_for_next_page"`
	// Number of seats included in this response
	SeatsIncluded Number `json:"seats_included"`
}
type V1ListSeatBalancesResponse struct {
	APIResource
	Data       []*SeatBalanceModel                   `json:"data"`
	Pagination *V1ListSeatBalancesResponsePagination `json:"pagination"`
}
