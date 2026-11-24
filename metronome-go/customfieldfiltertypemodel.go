//
// File generated from our OpenAPI spec
//

package metronome

type CustomFieldFilterTypeModelEntity string

// List of values that CustomFieldFilterTypeModelEntity can take
const (
	CustomFieldFilterTypeModelEntityContract               CustomFieldFilterTypeModelEntity = "Contract"
	CustomFieldFilterTypeModelEntityCommit                 CustomFieldFilterTypeModelEntity = "Commit"
	CustomFieldFilterTypeModelEntityContractCredit         CustomFieldFilterTypeModelEntity = "ContractCredit"
	CustomFieldFilterTypeModelEntityContractCreditOrCommit CustomFieldFilterTypeModelEntity = "ContractCreditOrCommit"
)

type CustomFieldFilterTypeModel struct {
	APIResource
	Entity CustomFieldFilterTypeModelEntity `json:"entity"`
	Key    string                           `json:"key"`
	Value  string                           `json:"value"`
}
