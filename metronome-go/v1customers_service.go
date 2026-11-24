//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1CustomersService is used to invoke v1customers related APIs.
type v1CustomersService struct {
	B   Backend
	Key string
}

// Use this endpoint to archive a customer while preserving auditability. Archiving a customer will automatically archive all contracts as of the current date and void all corresponding invoices. Use this endpoint if a customer is onboarded by mistake.
//
// ### Usage guidelines:
// - Once a customer is archived, it cannot be unarchived.
// - Archived customers can still be viewed through the API or the UI for audit purposes.
// - Ingest aliases remain idempotent for archived customers. In order to reuse an ingest alias, first remove the ingest alias from the customer prior to archiving.
// - Any notifications associated with the customer will no longer be triggered.
func (c v1CustomersService) ArchiveCustomer(ctx context.Context, params *V1CustomersArchiveCustomerParams) (*V1ArchiveCustomerResponse, error) {
	if params == nil {
		params = &V1CustomersArchiveCustomerParams{}
	}
	params.Context = ctx
	v1archivecustomerresponse := &V1ArchiveCustomerResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/customers/archive", c.Key, params, v1archivecustomerresponse)
	return v1archivecustomerresponse, err
}

// Deprecate an existing billing configuration for a customer to handle churn or billing and collection preference changes. Archiving a billing configuration takes effect immediately. If there are active contracts using the configuration, Metronome will archive the configuration on the contract and immediately stop metering to downstream systems.
//
// ### Use this endpoint to:
// - Remove billing provider customer data and configurations when no longer needed
// - Clean up test or deprecated billing provider configurations
// - Free up uniqueness keys for reuse with new billing provider configurations
// - Disable threshold recharge configurations associated with archived billing providers
//
// ### Key response fields:
// A successful response returns:
// - `success`: Boolean indicating the operation completed successfully
// - `error`: Null on success, error message on failure
//
// ### Usage guidelines:
// - Archiving a contract configuration during a grace period will result in the invoice not being sent to the customer
// - Automatically disables both spend-based and credit-based threshold recharge configurations for contracts using the archived billing provider
// - You can archive multiple configurations for a single customer in a single request, but any validation failures for an individual configuration will prevent the entire operation from succeeding
func (c v1CustomersService) ArchiveCustomerBillingProviderConfigurations(ctx context.Context, params *V1CustomersArchiveCustomerBillingProviderConfigurationsParams) (*V1ArchiveCustomerBillingProviderConfigurationsResponse, error) {
	if params == nil {
		params = &V1CustomersArchiveCustomerBillingProviderConfigurationsParams{}
	}
	params.Context = ctx
	v1archivecustomerbillingproviderconfigurationsresponse := &V1ArchiveCustomerBillingProviderConfigurationsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/archiveCustomerBillingProviderConfigurations", c.Key, params, v1archivecustomerbillingproviderconfigurationsresponse)
	return v1archivecustomerbillingproviderconfigurationsresponse, err
}

// Archive existing revenue system configurations for a customer. Archiving a revenue system configuration takes effect immediately.
//
// ### Use this endpoint to:
// - Remove revenue system configurations when no longer needed
// - Clean up test or deprecated revenue system configurations
//
// ### Key response fields:
// A successful response returns:
// - `customer_id`: The customer ID the configurations belong to
// - `customer_revenue_system_configuration_ids`: The archived configuration IDs
func (c v1CustomersService) ArchiveCustomerRevenueSystemConfigurations(ctx context.Context, params *V1CustomersArchiveCustomerRevenueSystemConfigurationsParams) (*V1ArchiveCustomerRevenueSystemConfigurationsResponse, error) {
	if params == nil {
		params = &V1CustomersArchiveCustomerRevenueSystemConfigurationsParams{}
	}
	params.Context = ctx
	v1archivecustomerrevenuesystemconfigurationsresponse := &V1ArchiveCustomerRevenueSystemConfigurationsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/archiveCustomerRevenueSystemConfigurations", c.Key, params, v1archivecustomerrevenuesystemconfigurationsresponse)
	return v1archivecustomerrevenuesystemconfigurationsresponse, err
}

// Create a new customer in Metronome and optionally the billing configuration (recommended) which dictates where invoices for the customer will be sent or where payment will be collected.
//
// ### Use this endpoint to:
// Execute your customer provisioning workflows for either PLG motions, where customers originate in your platform, or SLG motions, where customers originate in your sales system.
//
// ### Key response fields:
// This end-point returns the `customer_id` created by the request. This id can be used to fetch relevant billing configurations and create contracts.
//
// ### Example workflow:
// - Generally, Metronome recommends first creating the customer in the downstream payment / ERP system when payment method is collected and then creating the customer in Metronome using the response (i.e. `customer_id`) from the downstream system. If you do not create a billing configuration on customer creation, you can add it later.
// - Once a customer is created, you can then create a contract for the customer. In the contract creation process, you will need to add the customer billing configuration to the contract to ensure Metronome invoices the customer correctly. This is because a customer can have multiple configurations.
// - As part of the customer creation process, set the ingest alias for the customer which will ensure usage is accurately mapped to the customer. Ingest aliases can be added or changed after the creation process as well.
//
// ### Usage guidelines:
// For details on different billing configurations for different systems, review the `/setCustomerBillingConfiguration` end-point.
func (c v1CustomersService) CreateCustomer(ctx context.Context, params *V1CustomersCreateCustomerParams) (*V1CreateCustomerResponse, error) {
	if params == nil {
		params = &V1CustomersCreateCustomerParams{}
	}
	params.Context = ctx
	v1createcustomerresponse := &V1CreateCustomerResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/customers", c.Key, params, v1createcustomerresponse)
	return v1createcustomerresponse, err
}

// Generate secure, embeddable dashboard URLs that allow you to seamlessly integrate Metronome's billing visualizations directly into your application. This endpoint creates authenticated iframe-ready URLs for customer-specific dashboards, providing a white-labeled billing experience without building custom UI.
//
// ### Use this endpoint to:
// - Embed billing dashboards directly in your customer portal or admin interface
// - Provide self-service access to invoices, usage data, and credit balances
// - Build white-labeled billing experiences with minimal development effort
//
// ### Key response fields:
// - A secure, time-limited URL that can be embedded in an iframe
// - The URL includes authentication tokens and configuration parameters
// - URLs are customer-specific and respect your security settings
//
// ### Usage guidelines:
// - Dashboard types: Choose from `invoices`, `usage`, or `commits_and_credits`
// - Customization options:
//   - `dashboard_options`: Configure dashboard behavior. Supported for the invoices dashboard only. Available keys include: `show_zero_usage_line_items` ("true"/"false"), `contract_id` (UUID, filters invoices by contract), `invoice_type` ("USAGE" or "SCHEDULED", filters by invoice type), and `invoice_status_filter` ("VOID", "FINALIZED", "DRAFT", "FINALIZED_AND_DRAFT", or "ALL")
//   - `color_overrides`: Match your brand's color palette
//
// - Iframe implementation: Embed the returned URL directly in an iframe element
// - Responsive design: Dashboards automatically adapt to container dimensions
func (c v1CustomersService) EmbeddableDashboard(ctx context.Context, params *V1CustomersEmbeddableDashboardParams) (*V1EmbeddableDashboardResponse, error) {
	if params == nil {
		params = &V1CustomersEmbeddableDashboardParams{}
	}
	params.Context = ctx
	v1embeddabledashboardresponse := &V1EmbeddableDashboardResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/dashboards/getEmbeddableUrl", c.Key, params, v1embeddabledashboardresponse)
	return v1embeddabledashboardresponse, err
}

// Get detailed information for a specific customer by their Metronome ID. Returns customer profile data including name, creation date, ingest aliases, configuration settings, and custom fields. Use this endpoint to fetch complete customer details for billing operations or account management.
//
// Note: If searching for a customer billing configuration, use the `/getCustomerBillingConfigurations` endpoint.
func (c v1CustomersService) GetCustomer(ctx context.Context, id string, params *V1CustomersGetCustomerParams) (*V1GetCustomerResponse, error) {
	if params == nil {
		params = &V1CustomersGetCustomerParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s", id)
	v1getcustomerresponse := &V1GetCustomerResponse{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, v1getcustomerresponse)
	return v1getcustomerresponse, err
}

// Returns all billing configurations previously set for the customer. Use during the contract provisioning process to fetch the `billing_provider_configuration_id` needed to set the contract billing configuration.
func (c v1CustomersService) GetCustomerBillingProviderConfigurations(ctx context.Context, params *V1CustomersGetCustomerBillingProviderConfigurationsParams) (*V1GetCustomerBillingProviderConfigurationsResponse, error) {
	if params == nil {
		params = &V1CustomersGetCustomerBillingProviderConfigurationsParams{}
	}
	params.Context = ctx
	v1getcustomerbillingproviderconfigurationsresponse := &V1GetCustomerBillingProviderConfigurationsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/getCustomerBillingProviderConfigurations", c.Key, params, v1getcustomerbillingproviderconfigurationsresponse)
	return v1getcustomerbillingproviderconfigurationsresponse, err
}

// Returns all revenue system configurations previously set for the customer.
func (c v1CustomersService) GetCustomerRevenueSystemConfigurations(ctx context.Context, params *V1CustomersGetCustomerRevenueSystemConfigurationsParams) (*V1GetCustomerRevenueSystemConfigurationsResponse, error) {
	if params == nil {
		params = &V1CustomersGetCustomerRevenueSystemConfigurationsParams{}
	}
	params.Context = ctx
	v1getcustomerrevenuesystemconfigurationsresponse := &V1GetCustomerRevenueSystemConfigurationsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/getCustomerRevenueSystemConfigurations", c.Key, params, v1getcustomerrevenuesystemconfigurationsresponse)
	return v1getcustomerrevenuesystemconfigurationsresponse, err
}

// Gets a paginated list of all customers in your Metronome account. Use this endpoint to browse your customer base, implement customer search functionality, or sync customer data with external systems. Returns customer details including IDs, names, and configuration settings. Supports filtering and pagination parameters for efficient data retrieval.
func (c v1CustomersService) ListCustomers(ctx context.Context, params *V1CustomersListCustomersParams) (*V1ListCustomersResponse, error) {
	if params == nil {
		params = &V1CustomersListCustomersParams{}
	}
	params.Context = ctx
	v1listcustomersresponse := &V1ListCustomersResponse{}
	err := c.B.Call(
		http.MethodGet, "/v1/customers", c.Key, params, v1listcustomersresponse)
	return v1listcustomersresponse, err
}

// Create a billing configuration for a customer. Once created, these configurations are available to associate to a contract and dictates which downstream system to collect payment in or send the invoice to. You can create multiple configurations per customer. The configuration formats are distinct for each downstream provider.
//
// ### Use this endpoint to:
// - Add the initial configuration to an existing customer. Once created, the billing configuration can then be associated to the customer's contract.
// - Add a new configuration to an existing customer. This might be used as part of an upgrade or downgrade workflow where the customer was previously billed through system A (e.g. Stripe) but will now be billed through system B (e.g. AWS). Once created, the new configuration can then be associated to the customer's contract.
// - Multiple configurations can be added per destination. For example, you can create two Stripe billing configurations for a Metronome customer that each have a distinct `collection_method`.
//
// ### Delivery method options:
// - `direct_to_billing_provider`: Use when Metronome should send invoices directly to the billing provider's API (e.g., Stripe, NetSuite). This is the most common method for automated billing workflows.
// - `tackle`: Use specifically for AWS Marketplace transactions that require Tackle's co-selling platform for partner attribution and commission tracking.
// - `aws_sqs`: Use when you want invoice data delivered to an AWS SQS queue for custom processing before sending to your billing system.
// - `aws_sns`: Use when you want invoice notifications published to an AWS SNS topic for event-driven billing workflows.
//
// ### Key response fields:
// The id for the customer billing configuration. This id can be used to associate the billing configuration to a contract.
//
// ### Usage guidelines:
// Must use the `delivery_method_id` if you have multiple Stripe accounts connected to Metronome.
func (c v1CustomersService) SetCustomerBillingProviderConfigurations(ctx context.Context, params *V1CustomersSetCustomerBillingProviderConfigurationsParams) (*V1SetCustomerBillingProviderConfigurationsResponse, error) {
	if params == nil {
		params = &V1CustomersSetCustomerBillingProviderConfigurationsParams{}
	}
	params.Context = ctx
	v1setcustomerbillingproviderconfigurationsresponse := &V1SetCustomerBillingProviderConfigurationsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/setCustomerBillingProviderConfigurations", c.Key, params, v1setcustomerbillingproviderconfigurationsresponse)
	return v1setcustomerbillingproviderconfigurationsresponse, err
}

// Updates the display name for a customer record. Use this to correct customer names, update business names after rebranding, or maintain accurate customer information for invoicing and reporting. Returns the updated customer object with the new name applied immediately across all billing documents and interfaces.
func (c v1CustomersService) SetCustomerName(ctx context.Context, id string, params *V1CustomersSetCustomerNameParams) (*V1SetCustomerNameResponse, error) {
	if params == nil {
		params = &V1CustomersSetCustomerNameParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/setName", id)
	v1setcustomernameresponse := &V1SetCustomerNameResponse{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, v1setcustomernameresponse)
	return v1setcustomernameresponse, err
}

// Create a revenue system configuration for a customer. Once created, these configurations are available to associate to a contract and dictates which downstream system to use for revenue workflows. The configuration formats are distinct for each downstream provider.
//
// ### Use this endpoint to:
// - Add the initial configuration to an existing customer. Once created, the revenue system configuration can then be associated to the customer's contract.
// - Add a new configuration to an existing customer.
//
// ### Key response fields:
// Returns the inserted configuration objects, including the id of the revenue system configuration(s). These ids can be used to associate a revenue system to a contract.
func (c v1CustomersService) SetCustomerRevenueSystemConfigurations(ctx context.Context, params *V1CustomersSetCustomerRevenueSystemConfigurationsParams) (*V1SetCustomerRevenueSystemConfigurationsResponse, error) {
	if params == nil {
		params = &V1CustomersSetCustomerRevenueSystemConfigurationsParams{}
	}
	params.Context = ctx
	v1setcustomerrevenuesystemconfigurationsresponse := &V1SetCustomerRevenueSystemConfigurationsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/setCustomerRevenueSystemConfigurations", c.Key, params, v1setcustomerrevenuesystemconfigurationsresponse)
	return v1setcustomerrevenuesystemconfigurationsresponse, err
}

// Sets the ingest aliases for a customer. Use this endpoint to associate a Metronome customer with an internal ID for easier tracking between systems. Ingest aliases can be used in the `customer_id` field when sending usage events to Metronome.
//
// ### Usage guidelines:
// - This call is idempotent and fully replaces the set of ingest aliases for the given customer.
// - Switching an ingest alias from one customer to another will associate all corresponding usage to the new customer.
// - Use multiple ingest aliases to model child organizations within a single Metronome customer.
func (c v1CustomersService) SetIngestAliases(ctx context.Context, id string, params *V1CustomersSetIngestAliasesParams) (*V1SetIngestAliasesResponse, error) {
	if params == nil {
		params = &V1CustomersSetIngestAliasesParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/setIngestAliases", id)
	v1setingestaliasesresponse := &V1SetIngestAliasesResponse{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, v1setingestaliasesresponse)
	return v1setingestaliasesresponse, err
}

// Update configuration settings for a specific customer, such as external system integrations (e.g., Salesforce account ID) and other customer-specific billing parameters. Use this endpoint to modify customer configurations without affecting core customer data like name or ingest aliases.
func (c v1CustomersService) UpdateCustomerConfig(ctx context.Context, id string, params *V1CustomersUpdateCustomerConfigParams) (*V1UpdateCustomerConfigResponse, error) {
	if params == nil {
		params = &V1CustomersUpdateCustomerConfigParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/customers/%s/updateConfig", id)
	v1updatecustomerconfigresponse := &V1UpdateCustomerConfigResponse{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, v1updatecustomerconfigresponse)
	return v1updatecustomerconfigresponse, err
}
