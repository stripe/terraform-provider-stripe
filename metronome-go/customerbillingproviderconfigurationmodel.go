//
// File generated from our OpenAPI spec
//

package metronome

// The billing provider set for this configuration.
type CustomerBillingProviderConfigurationModelBillingProvider string

// List of values that CustomerBillingProviderConfigurationModelBillingProvider can take
const (
	CustomerBillingProviderConfigurationModelBillingProviderAwsMarketplace   CustomerBillingProviderConfigurationModelBillingProvider = "aws_marketplace"
	CustomerBillingProviderConfigurationModelBillingProviderStripe           CustomerBillingProviderConfigurationModelBillingProvider = "stripe"
	CustomerBillingProviderConfigurationModelBillingProviderNetsuite         CustomerBillingProviderConfigurationModelBillingProvider = "netsuite"
	CustomerBillingProviderConfigurationModelBillingProviderCustom           CustomerBillingProviderConfigurationModelBillingProvider = "custom"
	CustomerBillingProviderConfigurationModelBillingProviderAzureMarketplace CustomerBillingProviderConfigurationModelBillingProvider = "azure_marketplace"
	CustomerBillingProviderConfigurationModelBillingProviderQuickbooksOnline CustomerBillingProviderConfigurationModelBillingProvider = "quickbooks_online"
	CustomerBillingProviderConfigurationModelBillingProviderWorkday          CustomerBillingProviderConfigurationModelBillingProvider = "workday"
	CustomerBillingProviderConfigurationModelBillingProviderGcpMarketplace   CustomerBillingProviderConfigurationModelBillingProvider = "gcp_marketplace"
	CustomerBillingProviderConfigurationModelBillingProviderMetronome        CustomerBillingProviderConfigurationModelBillingProvider = "metronome"
)

// The method to use for delivering invoices to this customer.
type CustomerBillingProviderConfigurationModelDeliveryMethod string

// List of values that CustomerBillingProviderConfigurationModelDeliveryMethod can take
const (
	CustomerBillingProviderConfigurationModelDeliveryMethodDirectToBillingProvider CustomerBillingProviderConfigurationModelDeliveryMethod = "direct_to_billing_provider"
	CustomerBillingProviderConfigurationModelDeliveryMethodAwsSqs                  CustomerBillingProviderConfigurationModelDeliveryMethod = "aws_sqs"
	CustomerBillingProviderConfigurationModelDeliveryMethodTackle                  CustomerBillingProviderConfigurationModelDeliveryMethod = "tackle"
	CustomerBillingProviderConfigurationModelDeliveryMethodAwsSns                  CustomerBillingProviderConfigurationModelDeliveryMethod = "aws_sns"
)

type CustomerBillingProviderConfigurationModel struct {
	APIResource
	ArchivedAt string `json:"archived_at"`
	// The billing provider set for this configuration.
	BillingProvider CustomerBillingProviderConfigurationModelBillingProvider `json:"billing_provider"`
	// Configuration for the billing provider. The structure of this object is specific to the billing provider.
	Configuration map[string]JSON `json:"configuration"`
	CustomerID    string          `json:"customer_id"`
	// The method to use for delivering invoices to this customer.
	DeliveryMethod CustomerBillingProviderConfigurationModelDeliveryMethod `json:"delivery_method"`
	// Configuration for the delivery method. The structure of this object is specific to the delivery method.
	DeliveryMethodConfiguration map[string]JSON `json:"delivery_method_configuration"`
	// ID of the delivery method to use for this customer.
	DeliveryMethodID string `json:"delivery_method_id"`
	// ID of this configuration; can be provided as the billing_provider_configuration_id when creating a contract.
	ID string `json:"id"`
	// Rules that stop matching invoices from being sent to the billing provider. Only supported for Stripe billing provider configurations. When omitted, every invoice is sent to the billing provider.
	UnbillableInvoicesConfiguration []*UnbillableInvoiceRuleModel `json:"unbillable_invoices_configuration,omitempty"`
}
