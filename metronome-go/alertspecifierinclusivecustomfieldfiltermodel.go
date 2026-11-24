//
// File generated from our OpenAPI spec
//

package metronome

type AlertSpecifierInclusiveCustomFieldFilterModelEntity string

// List of values that AlertSpecifierInclusiveCustomFieldFilterModelEntity can take
const (
	AlertSpecifierInclusiveCustomFieldFilterModelEntityContract               AlertSpecifierInclusiveCustomFieldFilterModelEntity = "Contract"
	AlertSpecifierInclusiveCustomFieldFilterModelEntityCommit                 AlertSpecifierInclusiveCustomFieldFilterModelEntity = "Commit"
	AlertSpecifierInclusiveCustomFieldFilterModelEntityContractCredit         AlertSpecifierInclusiveCustomFieldFilterModelEntity = "ContractCredit"
	AlertSpecifierInclusiveCustomFieldFilterModelEntityContractCreditOrCommit AlertSpecifierInclusiveCustomFieldFilterModelEntity = "ContractCreditOrCommit"
)

type AlertSpecifierInclusiveCustomFieldFilterModel struct {
	APIResource
	Entity AlertSpecifierInclusiveCustomFieldFilterModelEntity `json:"entity"`
	Key    string                                              `json:"key"`
	Value  string                                              `json:"value,omitempty"`
}
