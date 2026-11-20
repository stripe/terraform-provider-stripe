//
// File generated from our OpenAPI spec
//

package metronome

type CustomerBillingProviderArchivePayloadModel struct {
	APIResource
	// Array of billing provider configuration IDs to archive
	CustomerBillingProviderConfigurationIDs []string `json:"customer_billing_provider_configuration_ids"`
	// The customer ID the billing provider configurations belong to
	CustomerID string `json:"customer_id"`
}
