//
// File generated from our OpenAPI spec
//

package metronome

type V1GetSubscriptionSeatsHistoryResponse struct {
	APIResource
	Data []*SubscriptionSeatsScheduleEntryModel `json:"data"`
	// Cursor for the next page of results
	NextPage string `json:"next_page"`
}
