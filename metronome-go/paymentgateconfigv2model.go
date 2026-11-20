//
// File generated from our OpenAPI spec
//

package metronome

// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
type PaymentGateConfigV2ModelPaymentGateType string

// List of values that PaymentGateConfigV2ModelPaymentGateType can take
const (
	PaymentGateConfigV2ModelPaymentGateTypeNONE     PaymentGateConfigV2ModelPaymentGateType = "NONE"
	PaymentGateConfigV2ModelPaymentGateTypeSTRIPE   PaymentGateConfigV2ModelPaymentGateType = "STRIPE"
	PaymentGateConfigV2ModelPaymentGateTypeEXTERNAL PaymentGateConfigV2ModelPaymentGateType = "EXTERNAL"
)

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish  Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
type PaymentGateConfigV2ModelTaxType string

// List of values that PaymentGateConfigV2ModelTaxType can take
const (
	PaymentGateConfigV2ModelTaxTypeNONE          PaymentGateConfigV2ModelTaxType = "NONE"
	PaymentGateConfigV2ModelTaxTypeSTRIPE        PaymentGateConfigV2ModelTaxType = "STRIPE"
	PaymentGateConfigV2ModelTaxTypeANROK         PaymentGateConfigV2ModelTaxType = "ANROK"
	PaymentGateConfigV2ModelTaxTypePRECALCULATED PaymentGateConfigV2ModelTaxType = "PRECALCULATED"
)

// If left blank, will default to INVOICE
type PaymentGateConfigV2ModelStripeConfigPaymentType string

// List of values that PaymentGateConfigV2ModelStripeConfigPaymentType can take
const (
	PaymentGateConfigV2ModelStripeConfigPaymentTypeINVOICE       PaymentGateConfigV2ModelStripeConfigPaymentType = "INVOICE"
	PaymentGateConfigV2ModelStripeConfigPaymentTypePAYMENTINTENT PaymentGateConfigV2ModelStripeConfigPaymentType = "PAYMENT_INTENT"
)

// Only applicable if using STRIPE as your payment gateway type.
type PaymentGateConfigV2ModelStripeConfig struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType PaymentGateConfigV2ModelStripeConfigPaymentType `json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type PaymentGateConfigV2ModelPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount Number `json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName string `json:"tax_name,omitempty"`
}
type PaymentGateConfigV2Model struct {
	APIResource
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType PaymentGateConfigV2ModelPaymentGateType `json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *PaymentGateConfigV2ModelPrecalculatedTaxConfig `json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gateway type.
	StripeConfig *PaymentGateConfigV2ModelStripeConfig `json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish  Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType PaymentGateConfigV2ModelTaxType `json:"tax_type,omitempty"`
}
