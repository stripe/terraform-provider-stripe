//
// File generated from our OpenAPI spec
//

package metronome

// The type of invoice this rule applies to.
type UnbillableInvoiceRuleModelInvoiceType string

// List of values that UnbillableInvoiceRuleModelInvoiceType can take
const (
	UnbillableInvoiceRuleModelInvoiceTypeUSAGE     UnbillableInvoiceRuleModelInvoiceType = "USAGE"
	UnbillableInvoiceRuleModelInvoiceTypeSCHEDULED UnbillableInvoiceRuleModelInvoiceType = "SCHEDULED"
	UnbillableInvoiceRuleModelInvoiceTypeUsage     UnbillableInvoiceRuleModelInvoiceType = "usage"
	UnbillableInvoiceRuleModelInvoiceTypeScheduled UnbillableInvoiceRuleModelInvoiceType = "scheduled"
)

// An individual rule that, when evaluated to true, indicates that any invoices for this billing provider will not be sent to its associated destination for the associated contract. Rules only apply to the specified `invoice_type` (or all invoices if omitted) and `fiat_credit_type_id` (or all invoices if omitted). Rule precedence is evaluated from more specific to less specific. This method will fail with a 400 if multiple rules with the same specificity are included.
type UnbillableInvoiceRuleModel struct {
	APIResource
	// Restricts the rule to invoices in this fiat currency. Omit for a catch-all rule that applies to every currency of the `invoice_type`. Required when `max_amount` is set.
	FiatCreditTypeID string `json:"fiat_credit_type_id,omitempty"`
	// The type of invoice this rule applies to.
	InvoiceType UnbillableInvoiceRuleModelInvoiceType `json:"invoice_type"`
	// A positive decimal, in the units of `fiat_credit_type_id`. Only invoices whose total is at or below this amount are suppressed; a higher total is still sent to the billing provider. When omitted, every matching invoice is suppressed regardless of amount.
	MaxAmount Number `json:"max_amount,omitempty"`
}
