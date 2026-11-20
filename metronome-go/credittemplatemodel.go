//
// File generated from our OpenAPI spec
//

package metronome

type CreditTemplateModelRateType string

// List of values that CreditTemplateModelRateType can take
const (
	CreditTemplateModelRateTypeCOMMITRATE CreditTemplateModelRateType = "COMMIT_RATE"
	CreditTemplateModelRateTypeLISTRATE   CreditTemplateModelRateType = "LIST_RATE"
)

type CreditTemplateModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type CreditTemplateModel struct {
	APIResource
	AccessSchedule        *RelativeScheduleDurationModel `json:"access_schedule,omitempty"`
	ApplicableProductIDs  []string                       `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string                       `json:"applicable_product_tags,omitempty"`
	CustomFields          map[string]string              `json:"custom_fields,omitempty"`
	Description           string                         `json:"description,omitempty"`
	ID                    string                         `json:"id"`
	Name                  string                         `json:"name,omitempty"`
	Priority              Number                         `json:"priority,omitempty"`
	Product               *CreditTemplateModelProduct    `json:"product"`
	RateType              CreditTemplateModelRateType    `json:"rate_type,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
}
