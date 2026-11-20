//
// File generated from our OpenAPI spec
//

package metronome

type V1ListContractsOnPackageResponse struct {
	APIResource
	Data     []*ContractProjectionSummaryModel `json:"data"`
	NextPage string                            `json:"next_page,omitempty"`
}
