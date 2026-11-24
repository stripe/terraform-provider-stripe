//
// File generated from our OpenAPI spec
//

package metronome

// The revenue system provider (e.g. netsuite).
type CustomerRevenueSystemConfigurationV2ModelProvider string

// List of values that CustomerRevenueSystemConfigurationV2ModelProvider can take
const (
	CustomerRevenueSystemConfigurationV2ModelProviderNetsuite CustomerRevenueSystemConfigurationV2ModelProvider = "netsuite"
)

// The method to use for delivering data to the revenue system.
type CustomerRevenueSystemConfigurationV2ModelDeliveryMethod string

// List of values that CustomerRevenueSystemConfigurationV2ModelDeliveryMethod can take
const (
	CustomerRevenueSystemConfigurationV2ModelDeliveryMethodDirectToBillingProvider CustomerRevenueSystemConfigurationV2ModelDeliveryMethod = "direct_to_billing_provider"
	CustomerRevenueSystemConfigurationV2ModelDeliveryMethodAwsSqs                  CustomerRevenueSystemConfigurationV2ModelDeliveryMethod = "aws_sqs"
	CustomerRevenueSystemConfigurationV2ModelDeliveryMethodTackle                  CustomerRevenueSystemConfigurationV2ModelDeliveryMethod = "tackle"
	CustomerRevenueSystemConfigurationV2ModelDeliveryMethodAwsSns                  CustomerRevenueSystemConfigurationV2ModelDeliveryMethod = "aws_sns"
)

type CustomerRevenueSystemConfigurationV2Model struct {
	APIResource
	ArchivedAt string `json:"archived_at,omitempty"`
	// Configuration for the revenue system. The structure of this object is specific to the provider.
	Configuration map[string]JSON `json:"configuration"`
	CustomerID    string          `json:"customer_id"`
	// The method to use for delivering data to the revenue system.
	DeliveryMethod CustomerRevenueSystemConfigurationV2ModelDeliveryMethod `json:"delivery_method,omitempty"`
	// Configuration for the delivery method. The structure of this object is specific to the delivery method.
	DeliveryMethodConfiguration map[string]JSON `json:"delivery_method_configuration,omitempty"`
	// ID of the delivery method used for this customer configuration.
	DeliveryMethodID string `json:"delivery_method_id"`
	// ID of the revenue system configuration.
	ID string `json:"id"`
	// The revenue system provider (e.g. netsuite).
	Provider CustomerRevenueSystemConfigurationV2ModelProvider `json:"provider"`
}
