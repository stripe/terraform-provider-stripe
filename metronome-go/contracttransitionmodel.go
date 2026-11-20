//
// File generated from our OpenAPI spec
//

package metronome

type ContractTransitionModelType string

// List of values that ContractTransitionModelType can take
const (
	ContractTransitionModelTypeRENEWAL ContractTransitionModelType = "RENEWAL"
)

type ContractTransitionModel struct {
	APIResource
	FromContractID string                      `json:"from_contract_id"`
	ToContractID   string                      `json:"to_contract_id"`
	Type           ContractTransitionModelType `json:"type"`
}
