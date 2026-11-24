//
// File generated from our OpenAPI spec
//

package metronome

type V2EditContractResponseData struct {
	Edit *ContractEditModel `json:"edit,omitempty"`
	ID   string             `json:"id"`
}
type V2EditContractResponse struct {
	APIResource
	Data *V2EditContractResponseData `json:"data"`
}
