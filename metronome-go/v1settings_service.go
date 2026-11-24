//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1SettingsService is used to invoke v1settings related APIs.
type v1SettingsService struct {
	B   Backend
	Key string
}

// Archive a custom pricing unit. Once archived, it will no longer appear in pricing unit selectors by default.
func (c v1SettingsService) ArchiveCreditType(ctx context.Context, params *V1SettingsArchiveCreditTypeParams) (*V1ArchiveCreditTypeResponse, error) {
	if params == nil {
		params = &V1SettingsArchiveCreditTypeParams{}
	}
	params.Context = ctx
	v1archivecredittyperesponse := &V1ArchiveCreditTypeResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/credit-types/archive", c.Key, params, v1archivecredittyperesponse)
	return v1archivecredittyperesponse, err
}

// Create a custom pricing unit. Custom pricing units can be used to charge for usage in a non-fiat pricing unit, for example AI credits.
func (c v1SettingsService) CreateCreditType(ctx context.Context, params *V1SettingsCreateCreditTypeParams) (*V1CreateCreditTypeResponse, error) {
	if params == nil {
		params = &V1SettingsCreateCreditTypeParams{}
	}
	params.Context = ctx
	v1createcredittyperesponse := &V1CreateCreditTypeResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/credit-types/create", c.Key, params, v1createcredittyperesponse)
	return v1createcredittyperesponse, err
}

// Lists all configured billing providers and their delivery method configurations for your account. Returns provider details, delivery method IDs, and configuration settings needed for mapping individual customer contracts to billing integrations.
func (c v1SettingsService) ListConfiguredBillingProviders(ctx context.Context, params *V1SettingsListConfiguredBillingProvidersParams) (*V1ListConfiguredBillingProvidersResponse, error) {
	if params == nil {
		params = &V1SettingsListConfiguredBillingProvidersParams{}
	}
	params.Context = ctx
	v1listconfiguredbillingprovidersresponse := &V1ListConfiguredBillingProvidersResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/listConfiguredBillingProviders", c.Key, params, v1listconfiguredbillingprovidersresponse)
	return v1listconfiguredbillingprovidersresponse, err
}

// List all pricing units. All fiat currency types (for example, USD or GBP) will be included, as well as any custom pricing units that were configured. Custom pricing units can be used to charge for usage in a non-fiat pricing unit, for example AI credits.
//
// Note: The USD (cents) pricing unit is 2714e483-4ff1-48e4-9e25-ac732e8f24f2.
func (c v1SettingsService) ListCreditTypes(ctx context.Context, params *V1SettingsListCreditTypesParams) (*V1ListCreditTypesResponse, error) {
	if params == nil {
		params = &V1SettingsListCreditTypesParams{}
	}
	params.Context = ctx
	v1listcredittypesresponse := &V1ListCreditTypesResponse{}
	err := c.B.Call(
		http.MethodGet, "/v1/credit-types/list", c.Key, params, v1listcredittypesresponse)
	return v1listcredittypesresponse, err
}

// Set up account-level configuration for a billing provider. Once configured, individual contracts across customers can be mapped to this configuration using the returned delivery_method_id.
func (c v1SettingsService) SetUpBillingProvider(ctx context.Context, params *V1SettingsSetUpBillingProviderParams) (*V1SetUpBillingProviderResponse, error) {
	if params == nil {
		params = &V1SettingsSetUpBillingProviderParams{}
	}
	params.Context = ctx
	v1setupbillingproviderresponse := &V1SetUpBillingProviderResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/setUpBillingProvider", c.Key, params, v1setupbillingproviderresponse)
	return v1setupbillingproviderresponse, err
}

// Set the Anrok API token for some specified delivery_method_ids, which can be found in the `/listConfiguredBillingProviders` response. This maps the Anrok key to the appropriate billing entity. These API tokens are only used for Threshold Billing workflows today.
func (c v1SettingsService) UpsertAnrokAPIToken(ctx context.Context, params *V1SettingsUpsertAnrokAPITokenParams) (*V1UpsertAnrokAPITokenResponse, error) {
	if params == nil {
		params = &V1SettingsUpsertAnrokAPITokenParams{}
	}
	params.Context = ctx
	v1upsertanrokapitokenresponse := &V1UpsertAnrokAPITokenResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/upsertAnrokApiToken", c.Key, params, v1upsertanrokapitokenresponse)
	return v1upsertanrokapitokenresponse, err
}

// Set the Avalara credentials for some specified `delivery_method_ids`, which can be found in the `/listConfiguredBillingProviders` response. This maps the Avalara credentials to the appropriate billing entity. These credentials are only used for PLG Invoicing today.
func (c v1SettingsService) UpsertAvalaraCredentials(ctx context.Context, params *V1SettingsUpsertAvalaraCredentialsParams) (*V1UpsertAvalaraCredentialsResponse, error) {
	if params == nil {
		params = &V1SettingsUpsertAvalaraCredentialsParams{}
	}
	params.Context = ctx
	v1upsertavalaracredentialsresponse := &V1UpsertAvalaraCredentialsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/upsertAvalaraCredentials", c.Key, params, v1upsertavalaracredentialsresponse)
	return v1upsertavalaracredentialsresponse, err
}
