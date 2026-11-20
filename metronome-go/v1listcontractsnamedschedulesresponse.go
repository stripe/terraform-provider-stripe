//
// File generated from our OpenAPI spec
//

package metronome

type V1ListContractsNamedSchedulesResponse struct {
	APIResource
	Data []*CustomerContractNamedScheduleModel `json:"data"`
	// Cursor for the next page of results.
	NextPage string `json:"next_page"`
}
