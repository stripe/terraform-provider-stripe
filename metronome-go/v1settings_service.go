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

// Create billing settings for Stripe.
//
// If you plan on using prepaid credit grants and recharges, your Stripe API key needs the following access:
//
// - Read & Write
//   - Invoices
//   - PaymentIntents
//
// - Read only:
//   - PaymentMethods
//   - Customers
func (c v1SettingsService) CreateStripeBillingSettings(ctx context.Context, params *V1SettingsCreateStripeBillingSettingsParams) (*V1CreateStripeBillingSettingsResponse, error) {
	if params == nil {
		params = &V1SettingsCreateStripeBillingSettingsParams{}
	}
	params.Context = ctx
	v1createstripebillingsettingsresponse := &V1CreateStripeBillingSettingsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/client/billing-config/stripe", c.Key, params, v1createstripebillingsettingsresponse)
	return v1createstripebillingsettingsresponse, err
}

// Delete all billing settings for Stripe
func (c v1SettingsService) DeleteStripeBillingSettings(ctx context.Context, params *V1SettingsDeleteStripeBillingSettingsParams) (*V1DeleteStripeBillingSettingsResponse, error) {
	if params == nil {
		params = &V1SettingsDeleteStripeBillingSettingsParams{}
	}
	params.Context = ctx
	v1deletestripebillingsettingsresponse := &V1DeleteStripeBillingSettingsResponse{}
	err := c.B.Call(
		http.MethodDelete, "/v1/client/billing-config/stripe", c.Key, params, v1deletestripebillingsettingsresponse)
	return v1deletestripebillingsettingsresponse, err
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

// Rotate the HMAC secret used to sign delta stream messages. The new secret will be used for all messages after the `starting_on` time. If a `starting_on` time is not provided, the new secret will be used in 10 minutes.
// If this endpoint is called while a rotation is already in progress, the second call will replace the existing rotation and the previous rotation will be cancelled.
func (c v1SettingsService) RotateDeltaStreamSecret(ctx context.Context, params *V1SettingsRotateDeltaStreamSecretParams) (*V1RotateDeltaStreamSecretResponse, error) {
	if params == nil {
		params = &V1SettingsRotateDeltaStreamSecretParams{}
	}
	params.Context = ctx
	v1rotatedeltastreamsecretresponse := &V1RotateDeltaStreamSecretResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/rotateDeltaStreamSecret", c.Key, params, v1rotatedeltastreamsecretresponse)
	return v1rotatedeltastreamsecretresponse, err
}

// (**Deprecated:** use [Create Stripe billing settings](https://docs.stripe.com/api#operation/createStripeBillingSettings) instead) Sets the api key for the billing provider. This is used for sending invoices to the billing provider. For example, if the billing provider is Stripe, this is the restricted key.
func (c v1SettingsService) SetBillingProviderAPIKey(ctx context.Context, id string, params *V1SettingsSetBillingProviderAPIKeyParams) (*V1SetBillingProviderAPIKeyResponse, error) {
	if params == nil {
		params = &V1SettingsSetBillingProviderAPIKeyParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/client/billing-config/%s/apiKey", id)
	v1setbillingproviderapikeyresponse := &V1SetBillingProviderAPIKeyResponse{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, v1setbillingproviderapikeyresponse)
	return v1setbillingproviderapikeyresponse, err
}

// Sets the webhook secret for the client.  This is used for validating webhook messages send to the external Metronome webhook endpoint.
func (c v1SettingsService) SetClientWebhookSecret(ctx context.Context, params *V1SettingsSetClientWebhookSecretParams) (*V1SetClientWebhookSecretResponse, error) {
	if params == nil {
		params = &V1SettingsSetClientWebhookSecretParams{}
	}
	params.Context = ctx
	v1setclientwebhooksecretresponse := &V1SetClientWebhookSecretResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/client/config/webhook_secret", c.Key, params, v1setclientwebhooksecretresponse)
	return v1setclientwebhooksecretresponse, err
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

// Apply partial updates to billing settings for Stripe. Individual settings can be deleted by specifying values as `null`.
func (c v1SettingsService) UpdateStripeBillingSettings(ctx context.Context, params *V1SettingsUpdateStripeBillingSettingsParams) (*V1UpdateStripeBillingSettingsResponse, error) {
	if params == nil {
		params = &V1SettingsUpdateStripeBillingSettingsParams{}
	}
	params.Context = ctx
	v1updatestripebillingsettingsresponse := &V1UpdateStripeBillingSettingsResponse{}
	err := c.B.Call(
		http.MethodPatch, "/v1/client/billing-config/stripe", c.Key, params, v1updatestripebillingsettingsresponse)
	return v1updatestripebillingsettingsresponse, err
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
