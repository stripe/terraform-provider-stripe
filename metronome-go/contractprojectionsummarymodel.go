//
// File generated from our OpenAPI spec
//

package metronome

type ContractProjectionSummaryModel struct {
	APIResource
	ArchivedAt   string `json:"archived_at,omitempty"`
	ContractID   string `json:"contract_id"`
	CustomerID   string `json:"customer_id"`
	EndingBefore string `json:"ending_before,omitempty"`
	StartingAt   string `json:"starting_at"`
}
