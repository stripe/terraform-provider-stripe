//
// File generated from our OpenAPI spec
//

package metronome

type V1ListApprovalRequestsResponse struct {
	APIResource
	Data []*ApprovalRequestModel `json:"data"`
	// Cursor to use for fetching the next page of results.
	NextPage string `json:"next_page"`
}
