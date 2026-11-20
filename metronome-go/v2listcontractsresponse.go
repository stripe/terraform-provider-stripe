//
// File generated from our OpenAPI spec
//

package metronome

type V2ListContractsResponse struct {
	APIResource
	// Pass this value as `cursor` in a subsequent request to fetch the next page of contracts. Null if there are no more contracts.
	Cursor string             `json:"cursor,omitempty"`
	Data   []*ContractV2Model `json:"data"`
}
