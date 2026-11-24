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
	original "github.com/stripe/terraform-provider-stripe/internal/provider"
	metronomeactions "github.com/stripe/terraform-provider-stripe/internal/provider/metronome/actions"
	metronomedatasources "github.com/stripe/terraform-provider-stripe/internal/provider/metronome/datasources"
	metronomeresources "github.com/stripe/terraform-provider-stripe/internal/provider/metronome/resources"
	"github.com/stripe/terraform-provider-stripe/internal/providerdata"
	metronomeclient "github.com/stripe/terraform-provider-stripe/metronome-go"
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
		metronomeresources.NewMetronomeProductResource,
		metronomeresources.NewMetronomeBillableMetricResource,
		metronomeresources.NewMetronomeCreditTypeResource,
		metronomeresources.NewMetronomeCustomFieldKeyResource,
		metronomeresources.NewMetronomeOffsetNotificationResource,
		metronomeresources.NewMetronomeCustomerResource,
		metronomeresources.NewMetronomeContractResource,
		metronomeresources.NewMetronomeCustomerAlertResource,
		metronomeresources.NewMetronomeCustomerCreditResource,
		metronomeresources.NewMetronomeCustomerCommitResource,
		metronomeresources.NewMetronomeCustomerBillingConfigurationResource,
		metronomeresources.NewMetronomeCustomerRevenueConfigurationResource,
		metronomeresources.NewMetronomeRateCardResource,
		metronomeresources.NewMetronomePackageResource}
}
func (p *Provider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{metronomedatasources.NewMetronomeGetApprovalRequestV1DataSource,
		metronomedatasources.NewMetronomeGetAuditLogsV1DataSource,
		metronomedatasources.NewMetronomeGetBillableMetricV1DataSource,
		metronomedatasources.NewMetronomeGetContractV1DataSource,
		metronomedatasources.NewMetronomeGetContractV2DataSource,
		metronomedatasources.NewMetronomeGetContractEditHistoryV2DataSource,
		metronomedatasources.NewMetronomeGetContractNamedScheduleV1DataSource,
		metronomedatasources.NewMetronomeGetContractRateScheduleV1DataSource,
		metronomedatasources.NewMetronomeGetCustomerV1DataSource,
		metronomedatasources.NewMetronomeGetCustomerAlertV1DataSource,
		metronomedatasources.NewMetronomeGetCustomerNamedScheduleV1DataSource,
		metronomedatasources.NewMetronomeGetInvoiceV1DataSource,
		metronomedatasources.NewMetronomeGetInvoicePdfV1DataSource,
		metronomedatasources.NewMetronomeGetNetBalanceV1DataSource,
		metronomedatasources.NewMetronomeGetNotificationConfigV2DataSource,
		metronomedatasources.NewMetronomeGetPackageV1DataSource,
		metronomedatasources.NewMetronomeGetPagedUsageV1DataSource,
		metronomedatasources.NewMetronomeGetProductV1DataSource,
		metronomedatasources.NewMetronomeGetRateCardV1DataSource,
		metronomedatasources.NewMetronomeGetRateCardNamedScheduleV1DataSource,
		metronomedatasources.NewMetronomeGetRateScheduleV1DataSource,
		metronomedatasources.NewMetronomeGetRatesV1DataSource,
		metronomedatasources.NewMetronomeGetServicesV1DataSource,
		metronomedatasources.NewMetronomeGetSubscriptionQuantityHistoryV1DataSource,
		metronomedatasources.NewMetronomeGetSubscriptionSeatsHistoryV1DataSource,
		metronomedatasources.NewMetronomeGetSubscriptionStatusV1DataSource,
		metronomedatasources.NewMetronomeGetUsageBatchV1DataSource,
		metronomedatasources.NewMetronomeGetWebhookDeliveryV1DataSource,
		metronomedatasources.NewMetronomeListAllBillableMetricsV1DataSource,
		metronomedatasources.NewMetronomeListApprovalRequestsV1DataSource,
		metronomedatasources.NewMetronomeListBillableMetricsV1DataSource,
		metronomedatasources.NewMetronomeListBreakdownInvoicesV1DataSource,
		metronomedatasources.NewMetronomeListContractsV1DataSource,
		metronomedatasources.NewMetronomeListContractsV2DataSource,
		metronomedatasources.NewMetronomeListContractsNamedSchedulesV1DataSource,
		metronomedatasources.NewMetronomeListContractsOnPackageV1DataSource,
		metronomedatasources.NewMetronomeListCreditTypesV1DataSource,
		metronomedatasources.NewMetronomeListCustomFieldKeysV1DataSource,
		metronomedatasources.NewMetronomeListCustomerAlertsV1DataSource,
		metronomedatasources.NewMetronomeListCustomerBalancesV1DataSource,
		metronomedatasources.NewMetronomeListCustomerCommitsV1DataSource,
		metronomedatasources.NewMetronomeListCustomerCreditsV1DataSource,
		metronomedatasources.NewMetronomeListCustomersV1DataSource,
		metronomedatasources.NewMetronomeListHistoricalBalancesV1DataSource,
		metronomedatasources.NewMetronomeListInvoicesV1DataSource,
		metronomedatasources.NewMetronomeListOffsetNotificationConfigsV2DataSource,
		metronomedatasources.NewMetronomeListPackagesV1DataSource,
		metronomedatasources.NewMetronomeListProductsV1DataSource,
		metronomedatasources.NewMetronomeListRateCardsV1DataSource,
		metronomedatasources.NewMetronomeListSeatBalancesV1DataSource,
		metronomedatasources.NewMetronomeListSpendBreakdownInvoicesV1DataSource,
		metronomedatasources.NewMetronomeListSystemNotificationConfigsV2DataSource,
		metronomedatasources.NewMetronomeListWebhookDeliveriesV1DataSource,
		metronomedatasources.NewMetronomePreviewCustomerEventsV1DataSource,
		metronomedatasources.NewMetronomeSearchEventsV1DataSource,
		metronomedatasources.NewMetronomeBillableMetricsDataSource,
		metronomedatasources.NewMetronomeStripeBillingProviderDataSource}
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
		metronomeactions.NewMetronomeAddCustomFieldKeyV1Action,
		metronomeactions.NewMetronomeAddManualBalanceLedgerEntryV1Action,
		metronomeactions.NewMetronomeAddRateV1Action,
		metronomeactions.NewMetronomeAddRatesV1Action,
		metronomeactions.NewMetronomeAmendContractV1Action,
		metronomeactions.NewMetronomeApproveApprovalRequestV1Action,
		metronomeactions.NewMetronomeArchiveAlertV1Action,
		metronomeactions.NewMetronomeArchiveBillableMetricV1Action,
		metronomeactions.NewMetronomeArchiveCommitV2Action,
		metronomeactions.NewMetronomeArchiveContractV1Action,
		metronomeactions.NewMetronomeArchiveCreditV2Action,
		metronomeactions.NewMetronomeArchiveCreditTypeV1Action,
		metronomeactions.NewMetronomeArchiveCustomerV1Action,
		metronomeactions.NewMetronomeArchiveCustomerBillingProviderConfigurationsV1Action,
		metronomeactions.NewMetronomeArchiveCustomerRevenueSystemConfigurationsV1Action,
		metronomeactions.NewMetronomeArchiveNotificationConfigV2Action,
		metronomeactions.NewMetronomeArchivePackageV1Action,
		metronomeactions.NewMetronomeArchiveProductListItemV1Action,
		metronomeactions.NewMetronomeArchiveRateCardV1Action,
		metronomeactions.NewMetronomeChargeSeatsV1Action,
		metronomeactions.NewMetronomeCreateAlertV1Action,
		metronomeactions.NewMetronomeCreateBillableMetricV1Action,
		metronomeactions.NewMetronomeCreateBillableMetricV1V1Action,
		metronomeactions.NewMetronomeCreateContractV1Action,
		metronomeactions.NewMetronomeCreateCreditTypeV1Action,
		metronomeactions.NewMetronomeCreateCustomerV1Action,
		metronomeactions.NewMetronomeCreateCustomerCommitV1Action,
		metronomeactions.NewMetronomeCreateCustomerCreditV1Action,
		metronomeactions.NewMetronomeCreateCustomerWithContractV1Action,
		metronomeactions.NewMetronomeCreateHistoricalContractUsageInvoicesV1Action,
		metronomeactions.NewMetronomeCreateNotificationConfigV2Action,
		metronomeactions.NewMetronomeCreatePackageV1Action,
		metronomeactions.NewMetronomeCreateProductV1Action,
		metronomeactions.NewMetronomeCreateRateCardV1Action,
		metronomeactions.NewMetronomeDeleteCustomFieldsV1Action,
		metronomeactions.NewMetronomeDenyApprovalRequestV1Action,
		metronomeactions.NewMetronomeDisableCommitTrueupV1Action,
		metronomeactions.NewMetronomeDisableCustomFieldKeyV1Action,
		metronomeactions.NewMetronomeEditCommitV2Action,
		metronomeactions.NewMetronomeEditContractV2Action,
		metronomeactions.NewMetronomeEditCreditV2Action,
		metronomeactions.NewMetronomeEditNotificationConfigV2Action,
		metronomeactions.NewMetronomeEmbeddableDashboardV1Action,
		metronomeactions.NewMetronomeIngestV1Action,
		metronomeactions.NewMetronomeMoveRateCardProductsV1Action,
		metronomeactions.NewMetronomeMoveToHistoryV2Action,
		metronomeactions.NewMetronomeRegenerateInvoiceV1Action,
		metronomeactions.NewMetronomeReleaseExternalPaymentGateThresholdCommitV1Action,
		metronomeactions.NewMetronomeResetCustomerAlertsV1Action,
		metronomeactions.NewMetronomeScheduleProServicesInvoiceV1Action,
		metronomeactions.NewMetronomeSetCustomFieldsV1Action,
		metronomeactions.NewMetronomeSetCustomerNameV1Action,
		metronomeactions.NewMetronomeSetIngestAliasesV1Action,
		metronomeactions.NewMetronomeSetRateCardProductsOrderV1Action,
		metronomeactions.NewMetronomeSetUpBillingProviderV1Action,
		metronomeactions.NewMetronomeSetUsageFilterV1Action,
		metronomeactions.NewMetronomeSubmitApprovalRequestsV1Action,
		metronomeactions.NewMetronomeUpdateBillableMetricV1Action,
		metronomeactions.NewMetronomeUpdateCommitEndDateV1Action,
		metronomeactions.NewMetronomeUpdateContractEndDateV1Action,
		metronomeactions.NewMetronomeUpdateContractNamedScheduleV1Action,
		metronomeactions.NewMetronomeUpdateCreditEndDateV1Action,
		metronomeactions.NewMetronomeUpdateCustomerConfigV1Action,
		metronomeactions.NewMetronomeUpdateCustomerNamedScheduleV1Action,
		metronomeactions.NewMetronomeUpdateInvoiceIssueDateV1Action,
		metronomeactions.NewMetronomeUpdateProductV1Action,
		metronomeactions.NewMetronomeUpdateRateCardV1Action,
		metronomeactions.NewMetronomeUpdateRateCardNamedScheduleV1Action,
		metronomeactions.NewMetronomeVoidInvoiceV1Action}
}
func (p *Provider) EphemeralResources(_ context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{newStripeCheckoutSessionEphemeralResource,
		newStripeFinancialConnectionsSessionEphemeralResource,
		newStripeIdentityVerificationSessionEphemeralResource,
		newStripeReportingReportRunEphemeralResource,
		newStripeTaxCalculationEphemeralResource,
		newStripeTokenEphemeralResource}
}
