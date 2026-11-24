//
// File generated from our OpenAPI spec
//

package metronome

// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
type PaymentGateConfigModelPaymentGateType string

// List of values that PaymentGateConfigModelPaymentGateType can take
const (
	PaymentGateConfigModelPaymentGateTypeNONE     PaymentGateConfigModelPaymentGateType = "NONE"
	PaymentGateConfigModelPaymentGateTypeSTRIPE   PaymentGateConfigModelPaymentGateType = "STRIPE"
	PaymentGateConfigModelPaymentGateTypeEXTERNAL PaymentGateConfigModelPaymentGateType = "EXTERNAL"
)

// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
type PaymentGateConfigModelTaxType string

// List of values that PaymentGateConfigModelTaxType can take
const (
	PaymentGateConfigModelTaxTypeNONE          PaymentGateConfigModelTaxType = "NONE"
	PaymentGateConfigModelTaxTypeSTRIPE        PaymentGateConfigModelTaxType = "STRIPE"
	PaymentGateConfigModelTaxTypeANROK         PaymentGateConfigModelTaxType = "ANROK"
	PaymentGateConfigModelTaxTypePRECALCULATED PaymentGateConfigModelTaxType = "PRECALCULATED"
)

// If left blank, will default to INVOICE
type PaymentGateConfigModelStripeConfigPaymentType string

// List of values that PaymentGateConfigModelStripeConfigPaymentType can take
const (
	PaymentGateConfigModelStripeConfigPaymentTypeINVOICE       PaymentGateConfigModelStripeConfigPaymentType = "INVOICE"
	PaymentGateConfigModelStripeConfigPaymentTypePAYMENTINTENT PaymentGateConfigModelStripeConfigPaymentType = "PAYMENT_INTENT"
)

// Only applicable if using STRIPE as your payment gate type.
type PaymentGateConfigModelStripeConfig struct {
	// Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.
	InvoiceMetadata map[string]string `json:"invoice_metadata,omitempty"`
	// If left blank, will default to INVOICE
	PaymentType PaymentGateConfigModelStripeConfigPaymentType `json:"payment_type"`
}

// Only applicable if using PRECALCULATED as your tax type.
type PaymentGateConfigModelPrecalculatedTaxConfig struct {
	// Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule
	TaxAmount Number `json:"tax_amount"`
	// Name of the tax to be applied. This may be used in an invoice line item description.
	TaxName string `json:"tax_name,omitempty"`
}
type PaymentGateConfigModel struct {
	APIResource
	// Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.
	PaymentGateType PaymentGateConfigModelPaymentGateType `json:"payment_gate_type"`
	// Only applicable if using PRECALCULATED as your tax type.
	PrecalculatedTaxConfig *PaymentGateConfigModelPrecalculatedTaxConfig `json:"precalculated_tax_config,omitempty"`
	// Only applicable if using STRIPE as your payment gate type.
	StripeConfig *PaymentGateConfigModelStripeConfig `json:"stripe_config,omitempty"`
	// Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.
	TaxType PaymentGateConfigModelTaxType `json:"tax_type,omitempty"`
}
