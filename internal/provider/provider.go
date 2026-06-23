//
// File generated from our OpenAPI spec
//

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/terraform-provider-stripe/internal/provider/actions"
	"github.com/stripe/terraform-provider-stripe/internal/provider/ephemeralresources"
	"github.com/stripe/terraform-provider-stripe/internal/provider/resources"
)

var _ provider.Provider = &StripeProvider{}
var _ provider.ProviderWithEphemeralResources = &StripeProvider{}
var _ provider.ProviderWithActions = &StripeProvider{}

type StripeProvider struct {
	version string
}

type StripeProviderModel struct {
	APIKey        types.String `tfsdk:"api_key"`
	StripeAccount types.String `tfsdk:"stripe_account"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &StripeProvider{version: version}
	}
}

func (p *StripeProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "stripe"
	resp.Version = p.version
}

func (p *StripeProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage Stripe resources with Terraform.",
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Stripe API key. Can also be set via the STRIPE_API_KEY environment variable.",
			},
			"stripe_account": schema.StringAttribute{
				Optional:    true,
				Description: "Connected account context for Connect-scoped requests. Can also be set via the STRIPE_ACCOUNT environment variable.",
			},
		},
	}
}

func (p *StripeProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config StripeProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiKey := config.APIKey.ValueString()
	if apiKey == "" {
		apiKey = os.Getenv("STRIPE_API_KEY")
	}
	if apiKey == "" {
		resp.Diagnostics.AddError(
			"Missing API Key",
			"Set the api_key provider attribute or the STRIPE_API_KEY environment variable.",
		)
		return
	}

	stripeAccount := config.StripeAccount.ValueString()
	if stripeAccount == "" {
		stripeAccount = os.Getenv("STRIPE_ACCOUNT")
	}

	var client *stripe.Client
	if stripeAccount != "" {
		backends := stripe.NewBackendsWithConfig(&stripe.BackendConfig{
			StripeContext: stripe.String(stripeAccount),
		})
		client = stripe.NewClient(apiKey, stripe.WithBackends(backends))
	} else {
		client = stripe.NewClient(apiKey)
	}
	resp.DataSourceData = client
	resp.ResourceData = client
	resp.EphemeralResourceData = client
	resp.ActionData = client
}

func (p *StripeProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *StripeProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewPersonResource,
		resources.NewApplePayDomainResource,
		resources.NewBillingAlertResource,
		resources.NewBillingCreditGrantResource,
		resources.NewBillingMeterResource,
		resources.NewBillingPortalConfigurationResource,
		resources.NewChargeResource,
		resources.NewClimateOrderResource,
		resources.NewCouponResource,
		resources.NewCreditNoteResource,
		resources.NewCustomerResource,
		resources.NewCustomerBalanceTransactionResource,
		resources.NewEntitlementsFeatureResource,
		resources.NewFileResource,
		resources.NewFileLinkResource,
		resources.NewInvoiceResource,
		resources.NewInvoiceItemResource,
		resources.NewIssuingCardResource,
		resources.NewIssuingCardholderResource,
		resources.NewIssuingDisputeResource,
		resources.NewIssuingPersonalizationDesignResource,
		resources.NewPaymentIntentResource,
		resources.NewPaymentLinkResource,
		resources.NewPaymentMethodResource,
		resources.NewPaymentMethodConfigurationResource,
		resources.NewPaymentMethodDomainResource,
		resources.NewPlanResource,
		resources.NewPriceResource,
		resources.NewProductResource,
		resources.NewProductFeatureResource,
		resources.NewPromotionCodeResource,
		resources.NewQuoteResource,
		resources.NewRadarValueListResource,
		resources.NewRadarValueListItemResource,
		resources.NewSetupIntentResource,
		resources.NewShippingRateResource,
		resources.NewSourceResource,
		resources.NewSubscriptionResource,
		resources.NewSubscriptionItemResource,
		resources.NewSubscriptionScheduleResource,
		resources.NewTaxRegistrationResource,
		resources.NewTaxIDResource,
		resources.NewTaxRateResource,
		resources.NewTerminalConfigurationResource,
		resources.NewTerminalLocationResource,
		resources.NewTerminalReaderResource,
		resources.NewTreasuryFinancialAccountResource,
		resources.NewWebhookEndpointResource,
		resources.NewV2CoreEventDestinationResource,
	}
}

func (p *StripeProvider) EphemeralResources(_ context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{
		ephemeralresources.NewCheckoutSessionEphemeralResource,
		ephemeralresources.NewFinancialConnectionsSessionEphemeralResource,
		ephemeralresources.NewIdentityVerificationSessionEphemeralResource,
		ephemeralresources.NewReportingReportRunEphemeralResource,
		ephemeralresources.NewTaxCalculationEphemeralResource,
		ephemeralresources.NewTokenEphemeralResource,
	}
}

func (p *StripeProvider) Actions(_ context.Context) []func() action.Action {
	return []func() action.Action{
		actions.NewFeeRefundAction,
		actions.NewForwardingRequestAction,
		actions.NewPayoutAction,
		actions.NewRefundAction,
		actions.NewTopupAction,
		actions.NewTransferAction,
		actions.NewTransferReversalAction,
		actions.NewTreasuryCreditReversalAction,
		actions.NewTreasuryDebitReversalAction,
		actions.NewTreasuryInboundTransferAction,
		actions.NewTreasuryOutboundPaymentAction,
		actions.NewTreasuryOutboundTransferAction,
	}
}
