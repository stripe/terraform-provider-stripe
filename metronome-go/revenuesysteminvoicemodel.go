//
// File generated from our OpenAPI spec
//

package metronome

type RevenueSystemInvoiceModel struct {
	APIResource
	// The error message from the revenue system, if available.
	ErrorMessage                    string `json:"error_message,omitempty"`
	RevenueSystemExternalEntityID   string `json:"revenue_system_external_entity_id,omitempty"`
	RevenueSystemExternalEntityType string `json:"revenue_system_external_entity_type"`
	RevenueSystemProvider           string `json:"revenue_system_provider"`
	SyncStatus                      string `json:"sync_status"`
}
