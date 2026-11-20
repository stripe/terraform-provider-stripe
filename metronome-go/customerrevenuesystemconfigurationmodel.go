//
// File generated from our OpenAPI spec
//

package metronome

// The revenue system provider (e.g. netsuite).
type CustomerRevenueSystemConfigurationModelProvider string

// List of values that CustomerRevenueSystemConfigurationModelProvider can take
const (
	CustomerRevenueSystemConfigurationModelProviderNetsuite CustomerRevenueSystemConfigurationModelProvider = "netsuite"
)

// The method to use for delivering data to the revenue system.
type CustomerRevenueSystemConfigurationModelDeliveryMethod string

// List of values that CustomerRevenueSystemConfigurationModelDeliveryMethod can take
const (
	CustomerRevenueSystemConfigurationModelDeliveryMethodDirectToBillingProvider CustomerRevenueSystemConfigurationModelDeliveryMethod = "direct_to_billing_provider"
	CustomerRevenueSystemConfigurationModelDeliveryMethodAwsSqs                  CustomerRevenueSystemConfigurationModelDeliveryMethod = "aws_sqs"
	CustomerRevenueSystemConfigurationModelDeliveryMethodTackle                  CustomerRevenueSystemConfigurationModelDeliveryMethod = "tackle"
	CustomerRevenueSystemConfigurationModelDeliveryMethodAwsSns                  CustomerRevenueSystemConfigurationModelDeliveryMethod = "aws_sns"
)

type CustomerRevenueSystemConfigurationModel struct {
	APIResource
	ArchivedAt string `json:"archived_at,omitempty"`
	// Configuration for the revenue system. The structure of this object is specific to the provider.
	Configuration map[string]JSON `json:"configuration"`
	CustomerID    string          `json:"customer_id"`
	// The method to use for delivering data to the revenue system.
	DeliveryMethod CustomerRevenueSystemConfigurationModelDeliveryMethod `json:"delivery_method,omitempty"`
	// Configuration for the delivery method. The structure of this object is specific to the delivery method.
	DeliveryMethodConfiguration map[string]JSON `json:"delivery_method_configuration,omitempty"`
	// ID of the delivery method used for this customer configuration.
	DeliveryMethodID string `json:"delivery_method_id"`
	// ID of this configuration; can be provided when associating to a contract.
	ID string `json:"id"`
	// The revenue system provider (e.g. netsuite).
	Provider CustomerRevenueSystemConfigurationModelProvider `json:"provider"`
}
