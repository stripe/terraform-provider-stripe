//
// File generated from our OpenAPI spec
//

package metronome

type CustomerRevenueSystemArchivePayloadModel struct {
	APIResource
	// The customer ID the revenue system configurations belong to
	CustomerID string `json:"customer_id"`
	// Array of revenue system configuration IDs to archive
	CustomerRevenueSystemConfigurationIDs []string `json:"customer_revenue_system_configuration_ids"`
}
