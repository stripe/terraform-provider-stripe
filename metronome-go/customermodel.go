//
// File generated from our OpenAPI spec
//

package metronome

type CustomerModel struct {
	APIResource
	CustomFields map[string]string `json:"custom_fields,omitempty"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest alias) that can be used in usage events
	ExternalID string `json:"external_id"`
	// the Metronome ID of the customer
	ID string `json:"id"`
	// aliases for this customer that can be used instead of the Metronome customer ID in usage events
	IngestAliases []string `json:"ingest_aliases"`
	Name          string   `json:"name"`
}
