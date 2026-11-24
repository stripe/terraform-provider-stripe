//
// File generated from our OpenAPI spec
//

package metronome

type CustomerAndContractModel struct {
	APIResource
	ContractID string `json:"contract_id"`
	// the Metronome ID of the customer
	CustomerID   string            `json:"customer_id"`
	CustomerName string            `json:"customer_name"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
	// aliases for this customer that can be used instead of the Metronome customer ID in usage events
	IngestAliases []string `json:"ingest_aliases"`
}
