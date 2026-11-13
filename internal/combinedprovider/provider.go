//
// File generated from our OpenAPI spec
//

package combinedprovider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	stripe "github.com/stripe/stripe-go/v87"
	metronomeclient "github.com/stripe/terraform-provider-stripe/internal/metronomeclient"
	original "github.com/stripe/terraform-provider-stripe/internal/provider"
	"github.com/stripe/terraform-provider-stripe/internal/provider/metronome"
	"github.com/stripe/terraform-provider-stripe/internal/providerdata"
	"os"
)

type Provider struct {
	stripe  provider.Provider
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider { return &Provider{stripe: original.New(version)(), version: version} }
}
func (p *Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	p.stripe.Metadata(ctx, req, resp)
}
func (p *Provider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	p.stripe.Schema(ctx, req, resp)
	attrs := make(map[string]schema.Attribute, len(resp.Schema.Attributes)+1)
	for key, value := range resp.Schema.Attributes {
		attrs[key] = value
	}
	attrs["metronome"] = schema.SingleNestedAttribute{Optional: true, Description: "Optional Metronome transport settings. Uses the same Stripe API key as Stripe.", Attributes: map[string]schema.Attribute{
		"base_url":     schema.StringAttribute{Optional: true, Description: "HTTPS API origin; defaults to METRONOME_BASE_URL, then https://api.metronome.com."},
		"ca_cert_file": schema.StringAttribute{Optional: true, Description: "PEM CA file; defaults to METRONOME_CA_CERT_FILE."},
	}}
	resp.Schema.Attributes = attrs
}
func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	data := &providerdata.Data{StripeStatus: providerdata.Status{State: "unconfigured"}, MetronomeStatus: providerdata.Status{State: "unconfigured"}}
	var key, account types.String
	var met types.Object
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("api_key"), &key)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("stripe_account"), &account)...)
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("metronome"), &met)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resolvedKey := key.ValueString()
	if resolvedKey == "" && !key.IsUnknown() {
		resolvedKey = os.Getenv("STRIPE_API_KEY")
	}
	data.StripeKey = resolvedKey
	data.StripeAccount = account.ValueString()
	if data.StripeAccount == "" && !account.IsUnknown() {
		data.StripeAccount = os.Getenv("STRIPE_ACCOUNT")
	}
	switch {
	case key.IsUnknown() || account.IsUnknown():
		data.StripeStatus = providerdata.Status{State: "unknown", Detail: "Stripe credentials or account context are not known until apply."}
	case resolvedKey == "":
		data.StripeStatus = providerdata.Status{State: "missing", Detail: "Set api_key or STRIPE_API_KEY."}
	default:
		var definition provider.SchemaResponse
		p.stripe.Schema(ctx, provider.SchemaRequest{}, &definition)
		var configured map[string]tftypes.Value
		if err := req.Config.Raw.As(&configured); err != nil {
			resp.Diagnostics.AddError("Invalid provider configuration", "Cannot project Stripe provider fields.")
			return
		}
		projected := map[string]tftypes.Value{}
		for key := range definition.Schema.Attributes {
			projected[key] = configured[key]
		}
		originalRequest := req
		originalRequest.Config = tfsdk.Config{Schema: definition.Schema, Raw: tftypes.NewValue(definition.Schema.Type().TerraformType(ctx), projected)}
		var originalResponse provider.ConfigureResponse
		p.stripe.Configure(ctx, originalRequest, &originalResponse)
		if originalResponse.Diagnostics.HasError() {
			data.StripeStatus = providerdata.Status{State: "invalid", Detail: "Stripe client could not be initialized."}
		} else {
			client, ok := originalResponse.ResourceData.(*stripe.Client)
			if !ok {
				data.StripeStatus = providerdata.Status{State: "invalid", Detail: "Stripe client configuration returned an unexpected type."}
			} else {
				data.Stripe = client
				data.StripeStatus = providerdata.Status{State: "ready"}
				data.ActivationBackend = providerdata.NewActivationBackend()
			}
		}
	}
	data.MetronomeStatus = data.StripeStatus
	if !key.IsUnknown() && resolvedKey != "" && !account.IsUnknown() {
		base, ca := "", ""
		unknown := met.IsUnknown()
		attrs := map[string]attr.Value{}
		if !met.IsNull() && !unknown {
			attrs = met.Attributes()
		}
		resolve := func(name, environment string) (string, bool) {
			value, ok := attrs[name].(types.String)
			if !ok || value.IsNull() {
				return os.Getenv(environment), false
			}
			if value.IsUnknown() {
				return "", true
			}
			return value.ValueString(), false
		}
		var baseUnknown, caUnknown bool
		base, baseUnknown = resolve("base_url", "METRONOME_BASE_URL")
		ca, caUnknown = resolve("ca_cert_file", "METRONOME_CA_CERT_FILE")
		if unknown || baseUnknown || caUnknown {
			data.MetronomeStatus = providerdata.Status{State: "unknown", Detail: "Metronome transport settings are not known until apply."}
		} else {
			client, err := metronomeclient.New(resolvedKey, base, ca, p.version)
			if err != nil {
				data.MetronomeStatus = providerdata.Status{State: "invalid", Detail: err.Error()}
			} else {
				data.Metronome = client
				data.MetronomeStatus = providerdata.Status{State: "ready"}
			}
		}
	}
	resp.ResourceData = data
	resp.DataSourceData = data
	resp.ActionData = data
	resp.EphemeralResourceData = data
}
func (p *Provider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{newStripePersonResource,
		newStripeApplePayDomainResource,
		newStripeAppsInstallResource,
		newStripeBillingAlertResource,
		newStripeBillingCreditGrantResource,
		newStripeBillingFeedbackOptionResource,
		newStripeBillingMeterResource,
		newStripeBillingPortalConfigurationResource,
		newStripeChargeResource,
		newStripeClimateOrderResource,
		newStripeCouponResource,
		newStripeCreditNoteResource,
		newStripeCustomerResource,
		newStripeCustomerBalanceTransactionResource,
		newStripeEntitlementsFeatureResource,
		newStripeFileResource,
		newStripeFileLinkResource,
		newStripeInvoiceResource,
		newStripeInvoiceItemResource,
		newStripeIssuingCardResource,
		newStripeIssuingCardholderResource,
		newStripeIssuingDisputeResource,
		newStripeIssuingPersonalizationDesignResource,
		newStripePaymentIntentResource,
		newStripePaymentLinkResource,
		newStripePaymentMethodResource,
		newStripePaymentMethodConfigurationResource,
		newStripePaymentMethodDomainResource,
		newStripePlanResource,
		newStripePriceResource,
		newStripeProductResource,
		newStripeProductFeatureResource,
		newStripeProductCatalogTrialOfferResource,
		newStripePromotionCodeResource,
		newStripeQuoteResource,
		newStripeRadarValueListResource,
		newStripeRadarValueListItemResource,
		newStripeSetupIntentResource,
		newStripeShippingRateResource,
		newStripeSourceResource,
		newStripeSubscriptionResource,
		newStripeSubscriptionItemResource,
		newStripeSubscriptionScheduleResource,
		newStripeTaxLocationResource,
		newStripeTaxRegistrationResource,
		newStripeTaxIDResource,
		newStripeTaxRateResource,
		newStripeTerminalConfigurationResource,
		newStripeTerminalLocationResource,
		newStripeTerminalReaderResource,
		newStripeThreeDSecureAuthenticationResource,
		newStripeTreasuryFinancialAccountResource,
		newStripeWebhookEndpointResource,
		newStripeV2CoreEventDestinationResource,
		newStripeV2CoreVaultGBBankAccountResource,
		newStripeV2CoreVaultUSBankAccountResource,
		newStripeV2MoneyManagementFinancialAddressResource,
		newStripeV2MoneyManagementInboundTransferResource,
		newStripeV2MoneyManagementOutboundPaymentResource,
		newStripeV2MoneyManagementOutboundSetupIntentResource,
		newStripeV2MoneyManagementPayoutIntentResource,
		newStripeV2MoneyManagementRecipientVerificationResource,
		metronome.NewMetronomeProductResource,
		metronome.NewMetronomeBillableMetricResource,
		metronome.NewMetronomeCreditTypeResource,
		metronome.NewMetronomeCustomFieldKeyResource,
		metronome.NewMetronomeOffsetNotificationResource,
		metronome.NewMetronomeCustomerResource,
		metronome.NewMetronomeContractResource,
		metronome.NewMetronomeCustomerAlertResource,
		metronome.NewMetronomeCustomerCreditResource,
		metronome.NewMetronomeCustomerCommitResource,
		metronome.NewMetronomeCustomerBillingConfigurationResource,
		metronome.NewMetronomeCustomerRevenueConfigurationResource,
		metronome.NewMetronomeRateCardResource,
		metronome.NewMetronomePackageResource,
		metronome.NewMetronomeAccountBridgeResource}
}
func (p *Provider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{metronome.NewMetronomeGetApprovalRequestV1DataSource,
		metronome.NewMetronomeGetAuditLogsV1DataSource,
		metronome.NewMetronomeGetBillableMetricV1DataSource,
		metronome.NewMetronomeGetContractV1DataSource,
		metronome.NewMetronomeGetContractV2DataSource,
		metronome.NewMetronomeGetContractEditHistoryV2DataSource,
		metronome.NewMetronomeGetContractNamedScheduleV1DataSource,
		metronome.NewMetronomeGetContractRateScheduleV1DataSource,
		metronome.NewMetronomeGetCustomerV1DataSource,
		metronome.NewMetronomeGetCustomerAlertV1DataSource,
		metronome.NewMetronomeGetCustomerNamedScheduleV1DataSource,
		metronome.NewMetronomeGetInvoiceV1DataSource,
		metronome.NewMetronomeGetInvoicePdfV1DataSource,
		metronome.NewMetronomeGetNetBalanceV1DataSource,
		metronome.NewMetronomeGetNotificationConfigV2DataSource,
		metronome.NewMetronomeGetPackageV1DataSource,
		metronome.NewMetronomeGetPagedUsageV1DataSource,
		metronome.NewMetronomeGetProductV1DataSource,
		metronome.NewMetronomeGetPurchasedSeatsV1DataSource,
		metronome.NewMetronomeGetRateCardV1DataSource,
		metronome.NewMetronomeGetRateCardNamedScheduleV1DataSource,
		metronome.NewMetronomeGetRateScheduleV1DataSource,
		metronome.NewMetronomeGetRatesV1DataSource,
		metronome.NewMetronomeGetSeatsUsageV1DataSource,
		metronome.NewMetronomeGetServicesV1DataSource,
		metronome.NewMetronomeGetSubscriptionQuantityHistoryV1DataSource,
		metronome.NewMetronomeGetSubscriptionSeatsHistoryV1DataSource,
		metronome.NewMetronomeGetSubscriptionStatusV1DataSource,
		metronome.NewMetronomeGetUsageBatchV1DataSource,
		metronome.NewMetronomeGetWebhookDeliveryV1DataSource,
		metronome.NewMetronomeListAllBillableMetricsV1DataSource,
		metronome.NewMetronomeListApprovalRequestsV1DataSource,
		metronome.NewMetronomeListBillableMetricsV1DataSource,
		metronome.NewMetronomeListBreakdownInvoicesV1DataSource,
		metronome.NewMetronomeListContractsV1DataSource,
		metronome.NewMetronomeListContractsV2DataSource,
		metronome.NewMetronomeListContractsNamedSchedulesV1DataSource,
		metronome.NewMetronomeListContractsOnPackageV1DataSource,
		metronome.NewMetronomeListCreditTypesV1DataSource,
		metronome.NewMetronomeListCustomFieldKeysV1DataSource,
		metronome.NewMetronomeListCustomerAlertsV1DataSource,
		metronome.NewMetronomeListCustomerBalancesV1DataSource,
		metronome.NewMetronomeListCustomerCommitsV1DataSource,
		metronome.NewMetronomeListCustomerCreditsV1DataSource,
		metronome.NewMetronomeListCustomersV1DataSource,
		metronome.NewMetronomeListHistoricalBalancesV1DataSource,
		metronome.NewMetronomeListInvoicesV1DataSource,
		metronome.NewMetronomeListOffsetNotificationConfigsV2DataSource,
		metronome.NewMetronomeListPackagesV1DataSource,
		metronome.NewMetronomeListProductsV1DataSource,
		metronome.NewMetronomeListRateCardsV1DataSource,
		metronome.NewMetronomeListSeatBalancesV1DataSource,
		metronome.NewMetronomeListSpendBreakdownInvoicesV1DataSource,
		metronome.NewMetronomeListSystemNotificationConfigsV2DataSource,
		metronome.NewMetronomeListWebhookDeliveriesV1DataSource,
		metronome.NewMetronomePreviewCustomerEventsV1DataSource,
		metronome.NewMetronomeSearchEventsV1DataSource,
		metronome.NewMetronomeBillableMetricsDataSource,
		metronome.NewMetronomeStripeBillingProviderDataSource}
}
func (p *Provider) Actions(_ context.Context) []func() action.Action {
	return []func() action.Action{newStripeFeeRefundAction,
		newStripeForwardingRequestAction,
		newStripePayoutAction,
		newStripeRefundAction,
		newStripeTopupAction,
		newStripeTransferAction,
		newStripeTransferReversalAction,
		newStripeTreasuryCreditReversalAction,
		newStripeTreasuryDebitReversalAction,
		newStripeTreasuryInboundTransferAction,
		newStripeTreasuryOutboundPaymentAction,
		newStripeTreasuryOutboundTransferAction,
		metronome.NewMetronomeAddCustomFieldKeyV1Action,
		metronome.NewMetronomeAddManualBalanceLedgerEntryV1Action,
		metronome.NewMetronomeAddRateV1Action,
		metronome.NewMetronomeAddRatesV1Action,
		metronome.NewMetronomeAmendContractV1Action,
		metronome.NewMetronomeApproveApprovalRequestV1Action,
		metronome.NewMetronomeArchiveAlertV1Action,
		metronome.NewMetronomeArchiveBillableMetricV1Action,
		metronome.NewMetronomeArchiveCommitV2Action,
		metronome.NewMetronomeArchiveContractV1Action,
		metronome.NewMetronomeArchiveCreditV2Action,
		metronome.NewMetronomeArchiveCreditTypeV1Action,
		metronome.NewMetronomeArchiveCustomerV1Action,
		metronome.NewMetronomeArchiveCustomerBillingProviderConfigurationsV1Action,
		metronome.NewMetronomeArchiveCustomerRevenueSystemConfigurationsV1Action,
		metronome.NewMetronomeArchiveNotificationConfigV2Action,
		metronome.NewMetronomeArchivePackageV1Action,
		metronome.NewMetronomeArchiveProductListItemV1Action,
		metronome.NewMetronomeArchiveRateCardV1Action,
		metronome.NewMetronomeChargeSeatsV1Action,
		metronome.NewMetronomeCreateAlertV1Action,
		metronome.NewMetronomeCreateBillableMetricV1Action,
		metronome.NewMetronomeCreateBillableMetricV1V1Action,
		metronome.NewMetronomeCreateContractV1Action,
		metronome.NewMetronomeCreateCreditTypeV1Action,
		metronome.NewMetronomeCreateCustomerV1Action,
		metronome.NewMetronomeCreateCustomerCommitV1Action,
		metronome.NewMetronomeCreateCustomerCreditV1Action,
		metronome.NewMetronomeCreateCustomerWithContractV1Action,
		metronome.NewMetronomeCreateHistoricalContractUsageInvoicesV1Action,
		metronome.NewMetronomeCreateNotificationConfigV2Action,
		metronome.NewMetronomeCreatePackageV1Action,
		metronome.NewMetronomeCreateProductV1Action,
		metronome.NewMetronomeCreateRateCardV1Action,
		metronome.NewMetronomeDeleteCustomFieldsV1Action,
		metronome.NewMetronomeDeleteStripeBillingSettingsV1Action,
		metronome.NewMetronomeDenyApprovalRequestV1Action,
		metronome.NewMetronomeDisableCommitTrueupV1Action,
		metronome.NewMetronomeDisableCustomFieldKeyV1Action,
		metronome.NewMetronomeEditCommitV2Action,
		metronome.NewMetronomeEditContractV2Action,
		metronome.NewMetronomeEditCreditV2Action,
		metronome.NewMetronomeEditNotificationConfigV2Action,
		metronome.NewMetronomeEmbeddableDashboardV1Action,
		metronome.NewMetronomeIngestV1Action,
		metronome.NewMetronomeIntegrationCloudwatchLoggerV1Action,
		metronome.NewMetronomeMoveRateCardProductsV1Action,
		metronome.NewMetronomeMoveToHistoryV2Action,
		metronome.NewMetronomeRegenerateInvoiceV1Action,
		metronome.NewMetronomeReleaseExternalPaymentGateThresholdCommitV1Action,
		metronome.NewMetronomeResetCustomerAlertsV1Action,
		metronome.NewMetronomeScheduleProServicesInvoiceV1Action,
		metronome.NewMetronomeSetCustomFieldsV1Action,
		metronome.NewMetronomeSetCustomerBillableStatusV1Action,
		metronome.NewMetronomeSetCustomerNameV1Action,
		metronome.NewMetronomeSetIngestAliasesV1Action,
		metronome.NewMetronomeSetRateCardProductsOrderV1Action,
		metronome.NewMetronomeSetUpBillingProviderV1Action,
		metronome.NewMetronomeSetUsageFilterV1Action,
		metronome.NewMetronomeSubmitApprovalRequestsV1Action,
		metronome.NewMetronomeUpdateActiveRechargeConfigV1Action,
		metronome.NewMetronomeUpdateBillableMetricV1Action,
		metronome.NewMetronomeUpdateCommitEndDateV1Action,
		metronome.NewMetronomeUpdateContractEndDateV1Action,
		metronome.NewMetronomeUpdateContractNamedScheduleV1Action,
		metronome.NewMetronomeUpdateCreditEndDateV1Action,
		metronome.NewMetronomeUpdateCustomerConfigV1Action,
		metronome.NewMetronomeUpdateCustomerNamedScheduleV1Action,
		metronome.NewMetronomeUpdateInvoiceIssueDateV1Action,
		metronome.NewMetronomeUpdateProductV1Action,
		metronome.NewMetronomeUpdateRateCardV1Action,
		metronome.NewMetronomeUpdateRateCardNamedScheduleV1Action,
		metronome.NewMetronomeVoidInvoiceV1Action}
}
func (p *Provider) EphemeralResources(_ context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{newStripeCheckoutSessionEphemeralResource,
		newStripeFinancialConnectionsSessionEphemeralResource,
		newStripeIdentityVerificationSessionEphemeralResource,
		newStripeReportingReportRunEphemeralResource,
		newStripeTaxCalculationEphemeralResource,
		newStripeTokenEphemeralResource}
}
