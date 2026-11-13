//
// File generated from our OpenAPI spec
//

package combinedprovider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stripe/terraform-provider-stripe/internal/provider/actions"
	"github.com/stripe/terraform-provider-stripe/internal/provider/ephemeralresources"
	"github.com/stripe/terraform-provider-stripe/internal/provider/resources"
	"github.com/stripe/terraform-provider-stripe/internal/providerdata"
)

type stripePersonResource struct{ *resources.PersonResource }

func newStripePersonResource() resource.Resource {
	return &stripePersonResource{PersonResource: resources.NewPersonResource().(*resources.PersonResource)}
}
func (r *stripePersonResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.PersonResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.PersonResource.Configure(ctx, req, resp)
}

type stripeApplePayDomainResource struct {
	*resources.ApplePayDomainResource
}

func newStripeApplePayDomainResource() resource.Resource {
	return &stripeApplePayDomainResource{ApplePayDomainResource: resources.NewApplePayDomainResource().(*resources.ApplePayDomainResource)}
}
func (r *stripeApplePayDomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.ApplePayDomainResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.ApplePayDomainResource.Configure(ctx, req, resp)
}

type stripeFeeRefundAction struct{ *actions.FeeRefundAction }

func newStripeFeeRefundAction() action.Action {
	return &stripeFeeRefundAction{FeeRefundAction: actions.NewFeeRefundAction().(*actions.FeeRefundAction)}
}
func (r *stripeFeeRefundAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.FeeRefundAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.FeeRefundAction.Configure(ctx, req, resp)
}

type stripeBillingAlertResource struct {
	*resources.BillingAlertResource
}

func newStripeBillingAlertResource() resource.Resource {
	return &stripeBillingAlertResource{BillingAlertResource: resources.NewBillingAlertResource().(*resources.BillingAlertResource)}
}
func (r *stripeBillingAlertResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.BillingAlertResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.BillingAlertResource.Configure(ctx, req, resp)
}

type stripeBillingCreditGrantResource struct {
	*resources.BillingCreditGrantResource
}

func newStripeBillingCreditGrantResource() resource.Resource {
	return &stripeBillingCreditGrantResource{BillingCreditGrantResource: resources.NewBillingCreditGrantResource().(*resources.BillingCreditGrantResource)}
}
func (r *stripeBillingCreditGrantResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.BillingCreditGrantResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.BillingCreditGrantResource.Configure(ctx, req, resp)
}

type stripeBillingMeterResource struct {
	*resources.BillingMeterResource
}

func newStripeBillingMeterResource() resource.Resource {
	return &stripeBillingMeterResource{BillingMeterResource: resources.NewBillingMeterResource().(*resources.BillingMeterResource)}
}
func (r *stripeBillingMeterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.BillingMeterResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.BillingMeterResource.Configure(ctx, req, resp)
}

type stripeBillingPortalConfigurationResource struct {
	*resources.BillingPortalConfigurationResource
}

func newStripeBillingPortalConfigurationResource() resource.Resource {
	return &stripeBillingPortalConfigurationResource{BillingPortalConfigurationResource: resources.NewBillingPortalConfigurationResource().(*resources.BillingPortalConfigurationResource)}
}
func (r *stripeBillingPortalConfigurationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.BillingPortalConfigurationResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.BillingPortalConfigurationResource.Configure(ctx, req, resp)
}

type stripeChargeResource struct{ *resources.ChargeResource }

func newStripeChargeResource() resource.Resource {
	return &stripeChargeResource{ChargeResource: resources.NewChargeResource().(*resources.ChargeResource)}
}
func (r *stripeChargeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.ChargeResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.ChargeResource.Configure(ctx, req, resp)
}

type stripeCheckoutSessionEphemeralResource struct {
	*ephemeralresources.CheckoutSessionEphemeralResource
}

func newStripeCheckoutSessionEphemeralResource() ephemeral.EphemeralResource {
	return &stripeCheckoutSessionEphemeralResource{CheckoutSessionEphemeralResource: ephemeralresources.NewCheckoutSessionEphemeralResource().(*ephemeralresources.CheckoutSessionEphemeralResource)}
}
func (r *stripeCheckoutSessionEphemeralResource) Configure(ctx context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		r.CheckoutSessionEphemeralResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.CheckoutSessionEphemeralResource.Configure(ctx, req, resp)
}

type stripeClimateOrderResource struct {
	*resources.ClimateOrderResource
}

func newStripeClimateOrderResource() resource.Resource {
	return &stripeClimateOrderResource{ClimateOrderResource: resources.NewClimateOrderResource().(*resources.ClimateOrderResource)}
}
func (r *stripeClimateOrderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.ClimateOrderResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.ClimateOrderResource.Configure(ctx, req, resp)
}

type stripeCouponResource struct{ *resources.CouponResource }

func newStripeCouponResource() resource.Resource {
	return &stripeCouponResource{CouponResource: resources.NewCouponResource().(*resources.CouponResource)}
}
func (r *stripeCouponResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.CouponResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.CouponResource.Configure(ctx, req, resp)
}

type stripeCreditNoteResource struct{ *resources.CreditNoteResource }

func newStripeCreditNoteResource() resource.Resource {
	return &stripeCreditNoteResource{CreditNoteResource: resources.NewCreditNoteResource().(*resources.CreditNoteResource)}
}
func (r *stripeCreditNoteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.CreditNoteResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.CreditNoteResource.Configure(ctx, req, resp)
}

type stripeCustomerResource struct{ *resources.CustomerResource }

func newStripeCustomerResource() resource.Resource {
	return &stripeCustomerResource{CustomerResource: resources.NewCustomerResource().(*resources.CustomerResource)}
}
func (r *stripeCustomerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.CustomerResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.CustomerResource.Configure(ctx, req, resp)
}

type stripeCustomerBalanceTransactionResource struct {
	*resources.CustomerBalanceTransactionResource
}

func newStripeCustomerBalanceTransactionResource() resource.Resource {
	return &stripeCustomerBalanceTransactionResource{CustomerBalanceTransactionResource: resources.NewCustomerBalanceTransactionResource().(*resources.CustomerBalanceTransactionResource)}
}
func (r *stripeCustomerBalanceTransactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.CustomerBalanceTransactionResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.CustomerBalanceTransactionResource.Configure(ctx, req, resp)
}

type stripeEntitlementsFeatureResource struct {
	*resources.EntitlementsFeatureResource
}

func newStripeEntitlementsFeatureResource() resource.Resource {
	return &stripeEntitlementsFeatureResource{EntitlementsFeatureResource: resources.NewEntitlementsFeatureResource().(*resources.EntitlementsFeatureResource)}
}
func (r *stripeEntitlementsFeatureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.EntitlementsFeatureResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.EntitlementsFeatureResource.Configure(ctx, req, resp)
}

type stripeFileResource struct{ *resources.FileResource }

func newStripeFileResource() resource.Resource {
	return &stripeFileResource{FileResource: resources.NewFileResource().(*resources.FileResource)}
}
func (r *stripeFileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.FileResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.FileResource.Configure(ctx, req, resp)
}

type stripeFileLinkResource struct{ *resources.FileLinkResource }

func newStripeFileLinkResource() resource.Resource {
	return &stripeFileLinkResource{FileLinkResource: resources.NewFileLinkResource().(*resources.FileLinkResource)}
}
func (r *stripeFileLinkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.FileLinkResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.FileLinkResource.Configure(ctx, req, resp)
}

type stripeFinancialConnectionsSessionEphemeralResource struct {
	*ephemeralresources.FinancialConnectionsSessionEphemeralResource
}

func newStripeFinancialConnectionsSessionEphemeralResource() ephemeral.EphemeralResource {
	return &stripeFinancialConnectionsSessionEphemeralResource{FinancialConnectionsSessionEphemeralResource: ephemeralresources.NewFinancialConnectionsSessionEphemeralResource().(*ephemeralresources.FinancialConnectionsSessionEphemeralResource)}
}
func (r *stripeFinancialConnectionsSessionEphemeralResource) Configure(ctx context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		r.FinancialConnectionsSessionEphemeralResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.FinancialConnectionsSessionEphemeralResource.Configure(ctx, req, resp)
}

type stripeForwardingRequestAction struct {
	*actions.ForwardingRequestAction
}

func newStripeForwardingRequestAction() action.Action {
	return &stripeForwardingRequestAction{ForwardingRequestAction: actions.NewForwardingRequestAction().(*actions.ForwardingRequestAction)}
}
func (r *stripeForwardingRequestAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.ForwardingRequestAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.ForwardingRequestAction.Configure(ctx, req, resp)
}

type stripeIdentityVerificationSessionEphemeralResource struct {
	*ephemeralresources.IdentityVerificationSessionEphemeralResource
}

func newStripeIdentityVerificationSessionEphemeralResource() ephemeral.EphemeralResource {
	return &stripeIdentityVerificationSessionEphemeralResource{IdentityVerificationSessionEphemeralResource: ephemeralresources.NewIdentityVerificationSessionEphemeralResource().(*ephemeralresources.IdentityVerificationSessionEphemeralResource)}
}
func (r *stripeIdentityVerificationSessionEphemeralResource) Configure(ctx context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		r.IdentityVerificationSessionEphemeralResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.IdentityVerificationSessionEphemeralResource.Configure(ctx, req, resp)
}

type stripeInvoiceResource struct{ *resources.InvoiceResource }

func newStripeInvoiceResource() resource.Resource {
	return &stripeInvoiceResource{InvoiceResource: resources.NewInvoiceResource().(*resources.InvoiceResource)}
}
func (r *stripeInvoiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.InvoiceResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.InvoiceResource.Configure(ctx, req, resp)
}

type stripeInvoiceItemResource struct{ *resources.InvoiceItemResource }

func newStripeInvoiceItemResource() resource.Resource {
	return &stripeInvoiceItemResource{InvoiceItemResource: resources.NewInvoiceItemResource().(*resources.InvoiceItemResource)}
}
func (r *stripeInvoiceItemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.InvoiceItemResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.InvoiceItemResource.Configure(ctx, req, resp)
}

type stripeIssuingCardResource struct{ *resources.IssuingCardResource }

func newStripeIssuingCardResource() resource.Resource {
	return &stripeIssuingCardResource{IssuingCardResource: resources.NewIssuingCardResource().(*resources.IssuingCardResource)}
}
func (r *stripeIssuingCardResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.IssuingCardResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.IssuingCardResource.Configure(ctx, req, resp)
}

type stripeIssuingCardholderResource struct {
	*resources.IssuingCardholderResource
}

func newStripeIssuingCardholderResource() resource.Resource {
	return &stripeIssuingCardholderResource{IssuingCardholderResource: resources.NewIssuingCardholderResource().(*resources.IssuingCardholderResource)}
}
func (r *stripeIssuingCardholderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.IssuingCardholderResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.IssuingCardholderResource.Configure(ctx, req, resp)
}

type stripeIssuingDisputeResource struct {
	*resources.IssuingDisputeResource
}

func newStripeIssuingDisputeResource() resource.Resource {
	return &stripeIssuingDisputeResource{IssuingDisputeResource: resources.NewIssuingDisputeResource().(*resources.IssuingDisputeResource)}
}
func (r *stripeIssuingDisputeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.IssuingDisputeResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.IssuingDisputeResource.Configure(ctx, req, resp)
}

type stripeIssuingPersonalizationDesignResource struct {
	*resources.IssuingPersonalizationDesignResource
}

func newStripeIssuingPersonalizationDesignResource() resource.Resource {
	return &stripeIssuingPersonalizationDesignResource{IssuingPersonalizationDesignResource: resources.NewIssuingPersonalizationDesignResource().(*resources.IssuingPersonalizationDesignResource)}
}
func (r *stripeIssuingPersonalizationDesignResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.IssuingPersonalizationDesignResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.IssuingPersonalizationDesignResource.Configure(ctx, req, resp)
}

type stripePaymentIntentResource struct {
	*resources.PaymentIntentResource
}

func newStripePaymentIntentResource() resource.Resource {
	return &stripePaymentIntentResource{PaymentIntentResource: resources.NewPaymentIntentResource().(*resources.PaymentIntentResource)}
}
func (r *stripePaymentIntentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.PaymentIntentResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.PaymentIntentResource.Configure(ctx, req, resp)
}

type stripePaymentLinkResource struct{ *resources.PaymentLinkResource }

func newStripePaymentLinkResource() resource.Resource {
	return &stripePaymentLinkResource{PaymentLinkResource: resources.NewPaymentLinkResource().(*resources.PaymentLinkResource)}
}
func (r *stripePaymentLinkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.PaymentLinkResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.PaymentLinkResource.Configure(ctx, req, resp)
}

type stripePaymentMethodResource struct {
	*resources.PaymentMethodResource
}

func newStripePaymentMethodResource() resource.Resource {
	return &stripePaymentMethodResource{PaymentMethodResource: resources.NewPaymentMethodResource().(*resources.PaymentMethodResource)}
}
func (r *stripePaymentMethodResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.PaymentMethodResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.PaymentMethodResource.Configure(ctx, req, resp)
}

type stripePaymentMethodConfigurationResource struct {
	*resources.PaymentMethodConfigurationResource
}

func newStripePaymentMethodConfigurationResource() resource.Resource {
	return &stripePaymentMethodConfigurationResource{PaymentMethodConfigurationResource: resources.NewPaymentMethodConfigurationResource().(*resources.PaymentMethodConfigurationResource)}
}
func (r *stripePaymentMethodConfigurationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.PaymentMethodConfigurationResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.PaymentMethodConfigurationResource.Configure(ctx, req, resp)
}

type stripePaymentMethodDomainResource struct {
	*resources.PaymentMethodDomainResource
}

func newStripePaymentMethodDomainResource() resource.Resource {
	return &stripePaymentMethodDomainResource{PaymentMethodDomainResource: resources.NewPaymentMethodDomainResource().(*resources.PaymentMethodDomainResource)}
}
func (r *stripePaymentMethodDomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.PaymentMethodDomainResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.PaymentMethodDomainResource.Configure(ctx, req, resp)
}

type stripePayoutAction struct{ *actions.PayoutAction }

func newStripePayoutAction() action.Action {
	return &stripePayoutAction{PayoutAction: actions.NewPayoutAction().(*actions.PayoutAction)}
}
func (r *stripePayoutAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.PayoutAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.PayoutAction.Configure(ctx, req, resp)
}

type stripePlanResource struct{ *resources.PlanResource }

func newStripePlanResource() resource.Resource {
	return &stripePlanResource{PlanResource: resources.NewPlanResource().(*resources.PlanResource)}
}
func (r *stripePlanResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.PlanResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.PlanResource.Configure(ctx, req, resp)
}

type stripePriceResource struct{ *resources.PriceResource }

func newStripePriceResource() resource.Resource {
	return &stripePriceResource{PriceResource: resources.NewPriceResource().(*resources.PriceResource)}
}
func (r *stripePriceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.PriceResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.PriceResource.Configure(ctx, req, resp)
}

type stripeProductResource struct{ *resources.ProductResource }

func newStripeProductResource() resource.Resource {
	return &stripeProductResource{ProductResource: resources.NewProductResource().(*resources.ProductResource)}
}
func (r *stripeProductResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.ProductResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.ProductResource.Configure(ctx, req, resp)
}

type stripeProductFeatureResource struct {
	*resources.ProductFeatureResource
}

func newStripeProductFeatureResource() resource.Resource {
	return &stripeProductFeatureResource{ProductFeatureResource: resources.NewProductFeatureResource().(*resources.ProductFeatureResource)}
}
func (r *stripeProductFeatureResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.ProductFeatureResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.ProductFeatureResource.Configure(ctx, req, resp)
}

type stripePromotionCodeResource struct {
	*resources.PromotionCodeResource
}

func newStripePromotionCodeResource() resource.Resource {
	return &stripePromotionCodeResource{PromotionCodeResource: resources.NewPromotionCodeResource().(*resources.PromotionCodeResource)}
}
func (r *stripePromotionCodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.PromotionCodeResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.PromotionCodeResource.Configure(ctx, req, resp)
}

type stripeQuoteResource struct{ *resources.QuoteResource }

func newStripeQuoteResource() resource.Resource {
	return &stripeQuoteResource{QuoteResource: resources.NewQuoteResource().(*resources.QuoteResource)}
}
func (r *stripeQuoteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.QuoteResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.QuoteResource.Configure(ctx, req, resp)
}

type stripeRadarValueListResource struct {
	*resources.RadarValueListResource
}

func newStripeRadarValueListResource() resource.Resource {
	return &stripeRadarValueListResource{RadarValueListResource: resources.NewRadarValueListResource().(*resources.RadarValueListResource)}
}
func (r *stripeRadarValueListResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.RadarValueListResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.RadarValueListResource.Configure(ctx, req, resp)
}

type stripeRadarValueListItemResource struct {
	*resources.RadarValueListItemResource
}

func newStripeRadarValueListItemResource() resource.Resource {
	return &stripeRadarValueListItemResource{RadarValueListItemResource: resources.NewRadarValueListItemResource().(*resources.RadarValueListItemResource)}
}
func (r *stripeRadarValueListItemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.RadarValueListItemResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.RadarValueListItemResource.Configure(ctx, req, resp)
}

type stripeRefundAction struct{ *actions.RefundAction }

func newStripeRefundAction() action.Action {
	return &stripeRefundAction{RefundAction: actions.NewRefundAction().(*actions.RefundAction)}
}
func (r *stripeRefundAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.RefundAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.RefundAction.Configure(ctx, req, resp)
}

type stripeReportingReportRunEphemeralResource struct {
	*ephemeralresources.ReportingReportRunEphemeralResource
}

func newStripeReportingReportRunEphemeralResource() ephemeral.EphemeralResource {
	return &stripeReportingReportRunEphemeralResource{ReportingReportRunEphemeralResource: ephemeralresources.NewReportingReportRunEphemeralResource().(*ephemeralresources.ReportingReportRunEphemeralResource)}
}
func (r *stripeReportingReportRunEphemeralResource) Configure(ctx context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		r.ReportingReportRunEphemeralResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.ReportingReportRunEphemeralResource.Configure(ctx, req, resp)
}

type stripeSetupIntentResource struct{ *resources.SetupIntentResource }

func newStripeSetupIntentResource() resource.Resource {
	return &stripeSetupIntentResource{SetupIntentResource: resources.NewSetupIntentResource().(*resources.SetupIntentResource)}
}
func (r *stripeSetupIntentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.SetupIntentResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.SetupIntentResource.Configure(ctx, req, resp)
}

type stripeShippingRateResource struct {
	*resources.ShippingRateResource
}

func newStripeShippingRateResource() resource.Resource {
	return &stripeShippingRateResource{ShippingRateResource: resources.NewShippingRateResource().(*resources.ShippingRateResource)}
}
func (r *stripeShippingRateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.ShippingRateResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.ShippingRateResource.Configure(ctx, req, resp)
}

type stripeSourceResource struct{ *resources.SourceResource }

func newStripeSourceResource() resource.Resource {
	return &stripeSourceResource{SourceResource: resources.NewSourceResource().(*resources.SourceResource)}
}
func (r *stripeSourceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.SourceResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.SourceResource.Configure(ctx, req, resp)
}

type stripeSubscriptionResource struct {
	*resources.SubscriptionResource
}

func newStripeSubscriptionResource() resource.Resource {
	return &stripeSubscriptionResource{SubscriptionResource: resources.NewSubscriptionResource().(*resources.SubscriptionResource)}
}
func (r *stripeSubscriptionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.SubscriptionResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.SubscriptionResource.Configure(ctx, req, resp)
}

type stripeSubscriptionItemResource struct {
	*resources.SubscriptionItemResource
}

func newStripeSubscriptionItemResource() resource.Resource {
	return &stripeSubscriptionItemResource{SubscriptionItemResource: resources.NewSubscriptionItemResource().(*resources.SubscriptionItemResource)}
}
func (r *stripeSubscriptionItemResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.SubscriptionItemResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.SubscriptionItemResource.Configure(ctx, req, resp)
}

type stripeSubscriptionScheduleResource struct {
	*resources.SubscriptionScheduleResource
}

func newStripeSubscriptionScheduleResource() resource.Resource {
	return &stripeSubscriptionScheduleResource{SubscriptionScheduleResource: resources.NewSubscriptionScheduleResource().(*resources.SubscriptionScheduleResource)}
}
func (r *stripeSubscriptionScheduleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.SubscriptionScheduleResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.SubscriptionScheduleResource.Configure(ctx, req, resp)
}

type stripeTaxCalculationEphemeralResource struct {
	*ephemeralresources.TaxCalculationEphemeralResource
}

func newStripeTaxCalculationEphemeralResource() ephemeral.EphemeralResource {
	return &stripeTaxCalculationEphemeralResource{TaxCalculationEphemeralResource: ephemeralresources.NewTaxCalculationEphemeralResource().(*ephemeralresources.TaxCalculationEphemeralResource)}
}
func (r *stripeTaxCalculationEphemeralResource) Configure(ctx context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TaxCalculationEphemeralResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TaxCalculationEphemeralResource.Configure(ctx, req, resp)
}

type stripeTaxRegistrationResource struct {
	*resources.TaxRegistrationResource
}

func newStripeTaxRegistrationResource() resource.Resource {
	return &stripeTaxRegistrationResource{TaxRegistrationResource: resources.NewTaxRegistrationResource().(*resources.TaxRegistrationResource)}
}
func (r *stripeTaxRegistrationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TaxRegistrationResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TaxRegistrationResource.Configure(ctx, req, resp)
}

type stripeTaxIDResource struct{ *resources.TaxIDResource }

func newStripeTaxIDResource() resource.Resource {
	return &stripeTaxIDResource{TaxIDResource: resources.NewTaxIDResource().(*resources.TaxIDResource)}
}
func (r *stripeTaxIDResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TaxIDResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TaxIDResource.Configure(ctx, req, resp)
}

type stripeTaxRateResource struct{ *resources.TaxRateResource }

func newStripeTaxRateResource() resource.Resource {
	return &stripeTaxRateResource{TaxRateResource: resources.NewTaxRateResource().(*resources.TaxRateResource)}
}
func (r *stripeTaxRateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TaxRateResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TaxRateResource.Configure(ctx, req, resp)
}

type stripeTerminalConfigurationResource struct {
	*resources.TerminalConfigurationResource
}

func newStripeTerminalConfigurationResource() resource.Resource {
	return &stripeTerminalConfigurationResource{TerminalConfigurationResource: resources.NewTerminalConfigurationResource().(*resources.TerminalConfigurationResource)}
}
func (r *stripeTerminalConfigurationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TerminalConfigurationResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TerminalConfigurationResource.Configure(ctx, req, resp)
}

type stripeTerminalLocationResource struct {
	*resources.TerminalLocationResource
}

func newStripeTerminalLocationResource() resource.Resource {
	return &stripeTerminalLocationResource{TerminalLocationResource: resources.NewTerminalLocationResource().(*resources.TerminalLocationResource)}
}
func (r *stripeTerminalLocationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TerminalLocationResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TerminalLocationResource.Configure(ctx, req, resp)
}

type stripeTerminalReaderResource struct {
	*resources.TerminalReaderResource
}

func newStripeTerminalReaderResource() resource.Resource {
	return &stripeTerminalReaderResource{TerminalReaderResource: resources.NewTerminalReaderResource().(*resources.TerminalReaderResource)}
}
func (r *stripeTerminalReaderResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TerminalReaderResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TerminalReaderResource.Configure(ctx, req, resp)
}

type stripeTokenEphemeralResource struct {
	*ephemeralresources.TokenEphemeralResource
}

func newStripeTokenEphemeralResource() ephemeral.EphemeralResource {
	return &stripeTokenEphemeralResource{TokenEphemeralResource: ephemeralresources.NewTokenEphemeralResource().(*ephemeralresources.TokenEphemeralResource)}
}
func (r *stripeTokenEphemeralResource) Configure(ctx context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TokenEphemeralResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TokenEphemeralResource.Configure(ctx, req, resp)
}

type stripeTopupAction struct{ *actions.TopupAction }

func newStripeTopupAction() action.Action {
	return &stripeTopupAction{TopupAction: actions.NewTopupAction().(*actions.TopupAction)}
}
func (r *stripeTopupAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TopupAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TopupAction.Configure(ctx, req, resp)
}

type stripeTransferAction struct{ *actions.TransferAction }

func newStripeTransferAction() action.Action {
	return &stripeTransferAction{TransferAction: actions.NewTransferAction().(*actions.TransferAction)}
}
func (r *stripeTransferAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TransferAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TransferAction.Configure(ctx, req, resp)
}

type stripeTransferReversalAction struct {
	*actions.TransferReversalAction
}

func newStripeTransferReversalAction() action.Action {
	return &stripeTransferReversalAction{TransferReversalAction: actions.NewTransferReversalAction().(*actions.TransferReversalAction)}
}
func (r *stripeTransferReversalAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TransferReversalAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TransferReversalAction.Configure(ctx, req, resp)
}

type stripeTreasuryCreditReversalAction struct {
	*actions.TreasuryCreditReversalAction
}

func newStripeTreasuryCreditReversalAction() action.Action {
	return &stripeTreasuryCreditReversalAction{TreasuryCreditReversalAction: actions.NewTreasuryCreditReversalAction().(*actions.TreasuryCreditReversalAction)}
}
func (r *stripeTreasuryCreditReversalAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TreasuryCreditReversalAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TreasuryCreditReversalAction.Configure(ctx, req, resp)
}

type stripeTreasuryDebitReversalAction struct {
	*actions.TreasuryDebitReversalAction
}

func newStripeTreasuryDebitReversalAction() action.Action {
	return &stripeTreasuryDebitReversalAction{TreasuryDebitReversalAction: actions.NewTreasuryDebitReversalAction().(*actions.TreasuryDebitReversalAction)}
}
func (r *stripeTreasuryDebitReversalAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TreasuryDebitReversalAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TreasuryDebitReversalAction.Configure(ctx, req, resp)
}

type stripeTreasuryFinancialAccountResource struct {
	*resources.TreasuryFinancialAccountResource
}

func newStripeTreasuryFinancialAccountResource() resource.Resource {
	return &stripeTreasuryFinancialAccountResource{TreasuryFinancialAccountResource: resources.NewTreasuryFinancialAccountResource().(*resources.TreasuryFinancialAccountResource)}
}
func (r *stripeTreasuryFinancialAccountResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TreasuryFinancialAccountResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TreasuryFinancialAccountResource.Configure(ctx, req, resp)
}

type stripeTreasuryInboundTransferAction struct {
	*actions.TreasuryInboundTransferAction
}

func newStripeTreasuryInboundTransferAction() action.Action {
	return &stripeTreasuryInboundTransferAction{TreasuryInboundTransferAction: actions.NewTreasuryInboundTransferAction().(*actions.TreasuryInboundTransferAction)}
}
func (r *stripeTreasuryInboundTransferAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TreasuryInboundTransferAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TreasuryInboundTransferAction.Configure(ctx, req, resp)
}

type stripeTreasuryOutboundPaymentAction struct {
	*actions.TreasuryOutboundPaymentAction
}

func newStripeTreasuryOutboundPaymentAction() action.Action {
	return &stripeTreasuryOutboundPaymentAction{TreasuryOutboundPaymentAction: actions.NewTreasuryOutboundPaymentAction().(*actions.TreasuryOutboundPaymentAction)}
}
func (r *stripeTreasuryOutboundPaymentAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TreasuryOutboundPaymentAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TreasuryOutboundPaymentAction.Configure(ctx, req, resp)
}

type stripeTreasuryOutboundTransferAction struct {
	*actions.TreasuryOutboundTransferAction
}

func newStripeTreasuryOutboundTransferAction() action.Action {
	return &stripeTreasuryOutboundTransferAction{TreasuryOutboundTransferAction: actions.NewTreasuryOutboundTransferAction().(*actions.TreasuryOutboundTransferAction)}
}
func (r *stripeTreasuryOutboundTransferAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		r.TreasuryOutboundTransferAction.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.TreasuryOutboundTransferAction.Configure(ctx, req, resp)
}

type stripeWebhookEndpointResource struct {
	*resources.WebhookEndpointResource
}

func newStripeWebhookEndpointResource() resource.Resource {
	return &stripeWebhookEndpointResource{WebhookEndpointResource: resources.NewWebhookEndpointResource().(*resources.WebhookEndpointResource)}
}
func (r *stripeWebhookEndpointResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.WebhookEndpointResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.WebhookEndpointResource.Configure(ctx, req, resp)
}

type stripeV2CoreEventDestinationResource struct {
	*resources.V2CoreEventDestinationResource
}

func newStripeV2CoreEventDestinationResource() resource.Resource {
	return &stripeV2CoreEventDestinationResource{V2CoreEventDestinationResource: resources.NewV2CoreEventDestinationResource().(*resources.V2CoreEventDestinationResource)}
}
func (r *stripeV2CoreEventDestinationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		r.V2CoreEventDestinationResource.Configure(ctx, req, resp)
		return
	}
	data, ok := req.ProviderData.(*providerdata.Data)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected combined Stripe service clients.")
		return
	}
	if err := data.StripeStatus.Err("Stripe"); err != nil {
		resp.Diagnostics.AddError("Stripe client unavailable", err.Error())
		return
	}
	req.ProviderData = data.Stripe
	r.V2CoreEventDestinationResource.Configure(ctx, req, resp)
}
