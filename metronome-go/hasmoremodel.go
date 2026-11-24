//
// File generated from our OpenAPI spec
//

package metronome

// Indicates whether there are more items than the limit for this endpoint. Use the respective list endpoints to get the full lists.
type HasMoreModel struct {
	APIResource
	// Whether there are more commits on this contract than the limit for this endpoint. Use the /contracts/customerCommits/list endpoint to get the full list of commits.
	Commits bool `json:"commits"`
	// Whether there are more credits on this contract than the limit for this endpoint. Use the /contracts/customerCredits/list endpoint to get the full list of credits.
	Credits bool `json:"credits"`
}
