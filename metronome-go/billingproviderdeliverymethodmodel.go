//
// File generated from our OpenAPI spec
//

package metronome

// The billing provider set for this configuration.
type BillingProviderDeliveryMethodModelBillingProvider string

// List of values that BillingProviderDeliveryMethodModelBillingProvider can take
const (
	BillingProviderDeliveryMethodModelBillingProviderAwsMarketplace   BillingProviderDeliveryMethodModelBillingProvider = "aws_marketplace"
	BillingProviderDeliveryMethodModelBillingProviderStripe           BillingProviderDeliveryMethodModelBillingProvider = "stripe"
	BillingProviderDeliveryMethodModelBillingProviderNetsuite         BillingProviderDeliveryMethodModelBillingProvider = "netsuite"
	BillingProviderDeliveryMethodModelBillingProviderCustom           BillingProviderDeliveryMethodModelBillingProvider = "custom"
	BillingProviderDeliveryMethodModelBillingProviderAzureMarketplace BillingProviderDeliveryMethodModelBillingProvider = "azure_marketplace"
	BillingProviderDeliveryMethodModelBillingProviderQuickbooksOnline BillingProviderDeliveryMethodModelBillingProvider = "quickbooks_online"
	BillingProviderDeliveryMethodModelBillingProviderWorkday          BillingProviderDeliveryMethodModelBillingProvider = "workday"
	BillingProviderDeliveryMethodModelBillingProviderGcpMarketplace   BillingProviderDeliveryMethodModelBillingProvider = "gcp_marketplace"
	BillingProviderDeliveryMethodModelBillingProviderMetronome        BillingProviderDeliveryMethodModelBillingProvider = "metronome"
)

// The method to use for delivering invoices to this customer.
type BillingProviderDeliveryMethodModelDeliveryMethod string

// List of values that BillingProviderDeliveryMethodModelDeliveryMethod can take
const (
	BillingProviderDeliveryMethodModelDeliveryMethodDirectToBillingProvider BillingProviderDeliveryMethodModelDeliveryMethod = "direct_to_billing_provider"
	BillingProviderDeliveryMethodModelDeliveryMethodAwsSqs                  BillingProviderDeliveryMethodModelDeliveryMethod = "aws_sqs"
	BillingProviderDeliveryMethodModelDeliveryMethodTackle                  BillingProviderDeliveryMethodModelDeliveryMethod = "tackle"
	BillingProviderDeliveryMethodModelDeliveryMethodAwsSns                  BillingProviderDeliveryMethodModelDeliveryMethod = "aws_sns"
)

type BillingProviderDeliveryMethodModel struct {
	APIResource
	// The billing provider set for this configuration.
	BillingProvider BillingProviderDeliveryMethodModelBillingProvider `json:"billing_provider"`
	// The method to use for delivering invoices to this customer.
	DeliveryMethod BillingProviderDeliveryMethodModelDeliveryMethod `json:"delivery_method"`
	// Configuration for the delivery method. The structure of this object is specific to the delivery method. Some configuration may be omitted for security reasons.
	DeliveryMethodConfiguration map[string]JSON `json:"delivery_method_configuration"`
	// ID of the delivery method to use for this customer.
	DeliveryMethodID string `json:"delivery_method_id"`
}
