//
// File generated from our OpenAPI spec
//

package metronome

// Archive a custom pricing unit. Once archived, it will no longer appear in pricing unit selectors by default.
type V1SettingsArchiveCreditTypeParams struct {
	Params `form:"*"`
	ID     *string `form:"id" json:"id"`
}

// Create a custom pricing unit. Custom pricing units can be used to charge for usage in a non-fiat pricing unit, for example AI credits.
type V1SettingsCreateCreditTypeParams struct {
	Params `form:"*"`
	// The name of the custom pricing unit. This will appear on invoices.
	Name *string `form:"name" json:"name"`
}

// Lists all configured billing providers and their delivery method configurations for your account. Returns provider details, delivery method IDs, and configuration settings needed for mapping individual customer contracts to billing integrations.
type V1SettingsListConfiguredBillingProvidersParams struct {
	Params `form:"*"`
	// The cursor to the next page of results
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
}

// List all pricing units. All fiat currency types (for example, USD or GBP) will be included, as well as any custom pricing units that were configured. Custom pricing units can be used to charge for usage in a non-fiat pricing unit, for example AI credits.
//
// Note: The USD (cents) pricing unit is 2714e483-4ff1-48e4-9e25-ac732e8f24f2.
type V1SettingsListCreditTypesParams struct {
	Params `form:"*"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
}

// Set up account-level configuration for a billing provider. Once configured, individual contracts across customers can be mapped to this configuration using the returned delivery_method_id.
type V1SettingsSetUpBillingProviderParams struct {
	Params `form:"*"`
	// The billing provider set for this configuration.
	BillingProvider *string `form:"billing_provider" json:"billing_provider"`
	// Account-level configuration for the billing provider. The structure of this object is specific to the billing provider and delivery provider combination. See examples below.
	Configuration map[string]JSON `form:"configuration" json:"configuration"`
	// The method to use for delivering invoices for this configuration.
	DeliveryMethod *string `form:"delivery_method" json:"delivery_method"`
}

// Set the Anrok API token for some specified delivery_method_ids, which can be found in the `/listConfiguredBillingProviders` response. This maps the Anrok key to the appropriate billing entity. These API tokens are only used for Threshold Billing workflows today.
type V1SettingsUpsertAnrokAPITokenParams struct {
	Params `form:"*"`
	// The Anrok API token that is added to the configuration.
	AnrokAPIToken *string `form:"anrok_api_token" json:"anrok_api_token"`
	// The delivery method IDs of the billing provider configurations to update, can be found in the response of the `/listConfiguredBillingProviders` endpoint.
	DeliveryMethodIDs []*string `form:"delivery_method_ids" json:"delivery_method_ids"`
}

// Set the Avalara credentials for some specified `delivery_method_ids`, which can be found in the `/listConfiguredBillingProviders` response. This maps the Avalara credentials to the appropriate billing entity. These credentials are only used for PLG Invoicing today.
type V1SettingsUpsertAvalaraCredentialsParams struct {
	Params `form:"*"`
	// The Avalara environment to use (SANDBOX or PRODUCTION).
	AvalaraEnvironment *string `form:"avalara_environment" json:"avalara_environment"`
	// The password for the Avalara account.
	AvalaraPassword *string `form:"avalara_password" json:"avalara_password"`
	// The username for the Avalara account.
	AvalaraUsername *string `form:"avalara_username" json:"avalara_username"`
	// Commit transactions if you want Metronome tax calculations used for reporting and tax filings.
	CommitTransactions *bool `form:"commit_transactions" json:"commit_transactions,omitempty"`
	// The delivery method IDs of the billing provider configurations to update, can be found in the response of the `/listConfiguredBillingProviders` endpoint.
	DeliveryMethodIDs []*string `form:"delivery_method_ids" json:"delivery_method_ids"`
}
