//
// File generated from our OpenAPI spec
//

package metronome

// The billing provider set for this configuration.
type CustomerBillingProviderConfigurationOutputModelBillingProvider string

// List of values that CustomerBillingProviderConfigurationOutputModelBillingProvider can take
const (
	CustomerBillingProviderConfigurationOutputModelBillingProviderAwsMarketplace   CustomerBillingProviderConfigurationOutputModelBillingProvider = "aws_marketplace"
	CustomerBillingProviderConfigurationOutputModelBillingProviderStripe           CustomerBillingProviderConfigurationOutputModelBillingProvider = "stripe"
	CustomerBillingProviderConfigurationOutputModelBillingProviderNetsuite         CustomerBillingProviderConfigurationOutputModelBillingProvider = "netsuite"
	CustomerBillingProviderConfigurationOutputModelBillingProviderCustom           CustomerBillingProviderConfigurationOutputModelBillingProvider = "custom"
	CustomerBillingProviderConfigurationOutputModelBillingProviderAzureMarketplace CustomerBillingProviderConfigurationOutputModelBillingProvider = "azure_marketplace"
	CustomerBillingProviderConfigurationOutputModelBillingProviderQuickbooksOnline CustomerBillingProviderConfigurationOutputModelBillingProvider = "quickbooks_online"
	CustomerBillingProviderConfigurationOutputModelBillingProviderWorkday          CustomerBillingProviderConfigurationOutputModelBillingProvider = "workday"
	CustomerBillingProviderConfigurationOutputModelBillingProviderGcpMarketplace   CustomerBillingProviderConfigurationOutputModelBillingProvider = "gcp_marketplace"
	CustomerBillingProviderConfigurationOutputModelBillingProviderMetronome        CustomerBillingProviderConfigurationOutputModelBillingProvider = "metronome"
)

// The tax provider set for this configuration.
type CustomerBillingProviderConfigurationOutputModelTaxProvider string

// List of values that CustomerBillingProviderConfigurationOutputModelTaxProvider can take
const (
	CustomerBillingProviderConfigurationOutputModelTaxProviderAnrok   CustomerBillingProviderConfigurationOutputModelTaxProvider = "anrok"
	CustomerBillingProviderConfigurationOutputModelTaxProviderAvalara CustomerBillingProviderConfigurationOutputModelTaxProvider = "avalara"
	CustomerBillingProviderConfigurationOutputModelTaxProviderStripe  CustomerBillingProviderConfigurationOutputModelTaxProvider = "stripe"
)

type CustomerBillingProviderConfigurationOutputModel struct {
	APIResource
	// The billing provider set for this configuration.
	BillingProvider CustomerBillingProviderConfigurationOutputModelBillingProvider `json:"billing_provider,omitempty"`
	// Configuration for the billing provider. The structure of this object is specific to the billing provider and delivery method combination.
	Configuration map[string]JSON `json:"configuration,omitempty"`
	// ID of the customer this configuration is associated with.
	CustomerID string `json:"customer_id,omitempty"`
	// ID of the delivery method used for this customer configuration.
	DeliveryMethodID string `json:"delivery_method_id,omitempty"`
	// ID of the created configuration
	ID string `json:"id,omitempty"`
	// The tax provider set for this configuration.
	TaxProvider CustomerBillingProviderConfigurationOutputModelTaxProvider `json:"tax_provider,omitempty"`
}
