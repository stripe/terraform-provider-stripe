//
// File generated from our OpenAPI spec
//

package metronome

type BalanceCustomFieldFilterTypeModelEntity string

// List of values that BalanceCustomFieldFilterTypeModelEntity can take
const (
	BalanceCustomFieldFilterTypeModelEntityCommit                 BalanceCustomFieldFilterTypeModelEntity = "Commit"
	BalanceCustomFieldFilterTypeModelEntityContractCredit         BalanceCustomFieldFilterTypeModelEntity = "ContractCredit"
	BalanceCustomFieldFilterTypeModelEntityContractCreditOrCommit BalanceCustomFieldFilterTypeModelEntity = "ContractCreditOrCommit"
)

type BalanceCustomFieldFilterTypeModel struct {
	APIResource
	Entity BalanceCustomFieldFilterTypeModelEntity `json:"entity"`
	Key    string                                  `json:"key"`
	Value  string                                  `json:"value"`
}
