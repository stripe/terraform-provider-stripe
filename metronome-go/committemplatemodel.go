//
// File generated from our OpenAPI spec
//

package metronome

type CommitTemplateModelType string

// List of values that CommitTemplateModelType can take
const (
	CommitTemplateModelTypePREPAID  CommitTemplateModelType = "PREPAID"
	CommitTemplateModelTypePOSTPAID CommitTemplateModelType = "POSTPAID"
)

type CommitTemplateModelRateType string

// List of values that CommitTemplateModelRateType can take
const (
	CommitTemplateModelRateTypeCOMMITRATE CommitTemplateModelRateType = "COMMIT_RATE"
	CommitTemplateModelRateTypeLISTRATE   CommitTemplateModelRateType = "LIST_RATE"
)

type CommitTemplateModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type CommitTemplateModelInvoiceScheduleScheduleItem struct {
	DateOffset *RelativeDateModel `json:"date_offset"`
	ID         string             `json:"id"`
	Quantity   Number             `json:"quantity"`
	UnitPrice  Number             `json:"unit_price"`
}

// The schedule that the customer will be invoiced for this commit.
type CommitTemplateModelInvoiceSchedule struct {
	CreditType *CreditTypeModel `json:"credit_type"`
	// If true, this schedule will not generate an invoice.
	DoNotInvoice  bool                                              `json:"do_not_invoice"`
	ScheduleItems []*CommitTemplateModelInvoiceScheduleScheduleItem `json:"schedule_items"`
}
type CommitTemplateModel struct {
	APIResource
	// The schedule that the customer will gain access to the credits purposed with this commit.
	AccessSchedule        *RelativeScheduleDurationModel `json:"access_schedule,omitempty"`
	ApplicableProductIDs  []string                       `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string                       `json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string              `json:"custom_fields,omitempty"`
	Description           string                         `json:"description,omitempty"`
	ID                    string                         `json:"id"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule *CommitTemplateModelInvoiceSchedule `json:"invoice_schedule,omitempty"`
	Name            string                              `json:"name,omitempty"`
	// If multiple credits or commits are applicable, the one with the lower priority will apply first.
	Priority         Number                      `json:"priority,omitempty"`
	Product          *CommitTemplateModelProduct `json:"product"`
	RateType         CommitTemplateModelRateType `json:"rate_type,omitempty"`
	RolloverFraction Number                      `json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	Type       CommitTemplateModelType `json:"type"`
}
