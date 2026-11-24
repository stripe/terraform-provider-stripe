//
// File generated from our OpenAPI spec
//

package metronome

// Use this endpoint to archive a customer while preserving auditability. Archiving a customer will automatically archive all contracts as of the current date and void all corresponding invoices. Use this endpoint if a customer is onboarded by mistake.
//
// ### Usage guidelines:
// - Once a customer is archived, it cannot be unarchived.
// - Archived customers can still be viewed through the API or the UI for audit purposes.
// - Ingest aliases remain idempotent for archived customers. In order to reuse an ingest alias, first remove the ingest alias from the customer prior to archiving.
// - Any notifications associated with the customer will no longer be triggered.
type V1CustomersArchiveCustomerParams struct {
	Params `form:"*"`
	ID     *string `form:"id" json:"id"`
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
type V1CustomersArchiveCustomerBillingProviderConfigurationsParams struct {
	Params `form:"*"`
	// Array of billing provider configuration IDs to archive
	CustomerBillingProviderConfigurationIDs []*string `form:"customer_billing_provider_configuration_ids" json:"customer_billing_provider_configuration_ids"`
	// The customer ID the billing provider configurations belong to
	CustomerID *string `form:"customer_id" json:"customer_id"`
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
type V1CustomersArchiveCustomerRevenueSystemConfigurationsParams struct {
	Params `form:"*"`
	// The customer ID the revenue system configurations belong to
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Array of revenue system configuration IDs to archive
	CustomerRevenueSystemConfigurationIDs []*string `form:"customer_revenue_system_configuration_ids" json:"customer_revenue_system_configuration_ids"`
}
type V1CustomersCreateCustomerCustomerBillingProviderConfigurationParams struct {
	// The billing provider set for this configuration.
	BillingProvider *string `form:"billing_provider" json:"billing_provider"`
	// Configuration for the billing provider. The structure of this object is specific to the billing provider and delivery provider combination. Defaults to an empty object, however, for most billing provider + delivery method combinations, it will not be a valid configuration.
	Configuration map[string]JSON `form:"configuration" json:"configuration,omitempty"`
	// The method to use for delivering invoices to this customer. If not provided, the `delivery_method_id` must be provided.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
	// ID of the delivery method to use for this customer. If not provided, the `delivery_method` must be provided.
	DeliveryMethodID *string `form:"delivery_method_id" json:"delivery_method_id,omitempty"`
	// Specifies which tax provider Metronome should use for tax calculation when billing through Stripe. This is only supported for Stripe billing provider configurations.
	TaxProvider *string `form:"tax_provider" json:"tax_provider,omitempty"`
}
type V1CustomersCreateCustomerCustomerRevenueSystemConfigurationParams struct {
	// Configuration for the revenue system provider. The structure of this object is specific to the revenue system provider. For NetSuite, this should contain `netsuite_customer_id`.
	Configuration map[string]JSON `form:"configuration" json:"configuration,omitempty"`
	// The method to use for delivering invoices to this customer. If not provided, the `delivery_method_id` must be provided.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
	// ID of the delivery method to use for this customer. If not provided, the `delivery_method` must be provided.
	DeliveryMethodID *string `form:"delivery_method_id" json:"delivery_method_id,omitempty"`
	// The revenue system provider set for this configuration.
	Provider *string `form:"provider" json:"provider"`
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
type V1CustomersCreateCustomerParams struct {
	Params                                `form:"*"`
	CustomerBillingProviderConfigurations []*V1CustomersCreateCustomerCustomerBillingProviderConfigurationParams `form:"customer_billing_provider_configurations" json:"customer_billing_provider_configurations,omitempty"`
	CustomerRevenueSystemConfigurations   []*V1CustomersCreateCustomerCustomerRevenueSystemConfigurationParams   `form:"customer_revenue_system_configurations" json:"customer_revenue_system_configurations,omitempty"`
	CustomFields                          map[string]string                                                      `form:"custom_fields" json:"custom_fields,omitempty"`
	// (deprecated, use ingest_aliases instead) an alias that can be used to refer to this customer in usage events
	ExternalID *string `form:"external_id" json:"external_id,omitempty"`
	// Aliases that can be used to refer to this customer in usage events
	IngestAliases []*string `form:"ingest_aliases" json:"ingest_aliases,omitempty"`
	// This will be truncated to 160 characters if the provided name is longer.
	Name *string `form:"name" json:"name"`
}

// Optional dashboard specific options
type V1CustomersEmbeddableDashboardDashboardOptionParams struct {
	// The option key name
	Key *string `form:"key" json:"key"`
	// The option value. For show_zero_usage_line_items: "true" or "false" (default "false"). For contract_id: a UUID filtering invoices to a specific contract. For invoice_type: "USAGE" or "SCHEDULED". For invoice_status_filter: "VOID", "FINALIZED", "DRAFT", "FINALIZED_AND_DRAFT", or "ALL". For hide_voided_invoices (deprecated): "true" or "false".
	Value *string `form:"value" json:"value"`
}

// Optional list of colors to override
type V1CustomersEmbeddableDashboardColorOverrideParams struct {
	// The color to override
	Name *string `form:"name" json:"name,omitempty"`
	// Hex value representation of the color
	Value *string `form:"value" json:"value,omitempty"`
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
type V1CustomersEmbeddableDashboardParams struct {
	Params `form:"*"`
	// Optional list of colors to override
	ColorOverrides []*V1CustomersEmbeddableDashboardColorOverrideParams `form:"color_overrides" json:"color_overrides,omitempty"`
	CustomerID     *string                                              `form:"customer_id" json:"customer_id"`
	// The type of dashboard to retrieve.
	Dashboard *string `form:"dashboard" json:"dashboard"`
	// Optional dashboard specific options
	DashboardOptions []*V1CustomersEmbeddableDashboardDashboardOptionParams `form:"dashboard_options" json:"dashboard_options,omitempty"`
}

// Get detailed information for a specific customer by their Metronome ID. Returns customer profile data including name, creation date, ingest aliases, configuration settings, and custom fields. Use this endpoint to fetch complete customer details for billing operations or account management.
//
// Note: If searching for a customer billing configuration, use the `/getCustomerBillingConfigurations` endpoint.
type V1CustomersGetCustomerParams struct {
	Params `form:"*"`
}

// Returns all billing configurations previously set for the customer. Use during the contract provisioning process to fetch the `billing_provider_configuration_id` needed to set the contract billing configuration.
type V1CustomersGetCustomerBillingProviderConfigurationsParams struct {
	Params          `form:"*"`
	CustomerID      *string `form:"customer_id" json:"customer_id"`
	IncludeArchived *bool   `form:"include_archived" json:"include_archived,omitempty"`
}

// Returns all revenue system configurations previously set for the customer.
type V1CustomersGetCustomerRevenueSystemConfigurationsParams struct {
	Params     `form:"*"`
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Whether to include archived configurations
	IncludeArchived *bool `form:"include_archived" json:"include_archived,omitempty"`
	// Filter configurations by revenue system provider
	Provider *string `form:"provider" json:"provider,omitempty"`
}

// Gets a paginated list of all customers in your Metronome account. Use this endpoint to browse your customer base, implement customer search functionality, or sync customer data with external systems. Returns customer details including IDs, names, and configuration settings. Supports filtering and pagination parameters for efficient data retrieval.
type V1CustomersListCustomersParams struct {
	Params `form:"*"`
	// Filter the customer list by customer_id.  Up to 100 ids can be provided.
	CustomerIDs []*string `form:"customer_ids" json:"customer_ids,omitempty"`
	// Filter the customer list by ingest_alias
	IngestAlias *string `form:"ingest_alias" json:"ingest_alias,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Filter the customer list to only return archived customers. By default, only active customers are returned.
	OnlyArchived *bool `form:"only_archived" json:"only_archived,omitempty"`
	// Filter the customer list by salesforce_account_id.  Up to 100 ids can be provided.
	SalesforceAccountIDs []*string `form:"salesforce_account_ids" json:"salesforce_account_ids,omitempty"`
}

// Rules that stop matching invoices from being sent to the billing provider. Only supported for Stripe billing provider configurations. When omitted, every invoice is sent to the billing provider.
type V1CustomersSetCustomerBillingProviderConfigurationsDataUnbillableInvoicesConfigurationParams struct {
	// Restricts the rule to invoices in this fiat currency. Omit for a catch-all rule that applies to every currency of the `invoice_type`. Required when `max_amount` is set.
	FiatCreditTypeID *string `form:"fiat_credit_type_id" json:"fiat_credit_type_id,omitempty"`
	// The type of invoice this rule applies to.
	InvoiceType *string `form:"invoice_type" json:"invoice_type"`
	// A positive decimal, in the units of `fiat_credit_type_id`. Only invoices whose total is at or below this amount are suppressed; a higher total is still sent to the billing provider. When omitted, every matching invoice is suppressed regardless of amount.
	MaxAmount *Number `form:"max_amount" json:"max_amount,omitempty"`
}
type V1CustomersSetCustomerBillingProviderConfigurationsDataParams struct {
	// The billing provider set for this configuration.
	BillingProvider *string `form:"billing_provider" json:"billing_provider"`
	// Configuration for the billing provider. The structure of this object is specific to the billing provider and delivery method combination. Defaults to an empty object, however, for most billing provider + delivery method combinations, it will not be a valid configuration.  For AWS marketplace configurations, the aws_is_subscription_product flag can be used to indicate a product with usage-based pricing.  More information can be found [here](https://docs.metronome.com/invoice-customers/solutions/marketplaces/invoice-aws/#provision-aws-marketplace-customers-in-metronome).
	Configuration map[string]JSON `form:"configuration" json:"configuration,omitempty"`
	CustomerID    *string         `form:"customer_id" json:"customer_id"`
	// The method to use for delivering invoices to this customer. If not provided, the `delivery_method_id` must be provided.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method,omitempty"`
	// ID of the delivery method to use for this customer. If not provided, the `delivery_method` must be provided.
	DeliveryMethodID *string `form:"delivery_method_id" json:"delivery_method_id,omitempty"`
	// Specifies which tax provider Metronome should use for tax calculation when billing through Stripe. This is only supported for Stripe billing provider configurations.
	TaxProvider *string `form:"tax_provider" json:"tax_provider,omitempty"`
	// Rules that stop matching invoices from being sent to the billing provider. Only supported for Stripe billing provider configurations. When omitted, every invoice is sent to the billing provider.
	UnbillableInvoicesConfiguration []*V1CustomersSetCustomerBillingProviderConfigurationsDataUnbillableInvoicesConfigurationParams `form:"unbillable_invoices_configuration" json:"unbillable_invoices_configuration,omitempty"`
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
type V1CustomersSetCustomerBillingProviderConfigurationsParams struct {
	Params `form:"*"`
	Data   []*V1CustomersSetCustomerBillingProviderConfigurationsDataParams `form:"data" json:"data"`
}

// Updates the display name for a customer record. Use this to correct customer names, update business names after rebranding, or maintain accurate customer information for invoicing and reporting. Returns the updated customer object with the new name applied immediately across all billing documents and interfaces.
type V1CustomersSetCustomerNameParams struct {
	Params `form:"*"`
	// The new name for the customer. This will be truncated to 160 characters if the provided name is longer.
	Name *string `form:"name" json:"name"`
}
type V1CustomersSetCustomerRevenueSystemConfigurationsDataParams struct {
	// Configuration for the revenue system. The structure of this object is specific to the provider.
	Configuration map[string]JSON `form:"configuration" json:"configuration"`
	CustomerID    *string         `form:"customer_id" json:"customer_id"`
	// Optional ID of the delivery method to use for this customer configuration, if multiple delivery methods are configured for the same provider.
	DeliveryMethodID *string `form:"delivery_method_id" json:"delivery_method_id,omitempty"`
	// The revenue system provider (e.g. netsuite).
	Provider *string `form:"provider" json:"provider"`
}

// Create a revenue system configuration for a customer. Once created, these configurations are available to associate to a contract and dictates which downstream system to use for revenue workflows. The configuration formats are distinct for each downstream provider.
//
// ### Use this endpoint to:
// - Add the initial configuration to an existing customer. Once created, the revenue system configuration can then be associated to the customer's contract.
// - Add a new configuration to an existing customer.
//
// ### Key response fields:
// Returns the inserted configuration objects, including the id of the revenue system configuration(s). These ids can be used to associate a revenue system to a contract.
type V1CustomersSetCustomerRevenueSystemConfigurationsParams struct {
	Params `form:"*"`
	Data   []*V1CustomersSetCustomerRevenueSystemConfigurationsDataParams `form:"data" json:"data"`
}

// Sets the ingest aliases for a customer. Use this endpoint to associate a Metronome customer with an internal ID for easier tracking between systems. Ingest aliases can be used in the `customer_id` field when sending usage events to Metronome.
//
// ### Usage guidelines:
// - This call is idempotent and fully replaces the set of ingest aliases for the given customer.
// - Switching an ingest alias from one customer to another will associate all corresponding usage to the new customer.
// - Use multiple ingest aliases to model child organizations within a single Metronome customer.
type V1CustomersSetIngestAliasesParams struct {
	Params        `form:"*"`
	IngestAliases []*string `form:"ingest_aliases" json:"ingest_aliases"`
}

// Update configuration settings for a specific customer, such as external system integrations (e.g., Salesforce account ID) and other customer-specific billing parameters. Use this endpoint to modify customer configurations without affecting core customer data like name or ingest aliases.
type V1CustomersUpdateCustomerConfigParams struct {
	Params `form:"*"`
	// The Salesforce account ID for the customer
	SalesforceAccountID *string `form:"salesforce_account_id" json:"salesforce_account_id,omitempty"`
}
