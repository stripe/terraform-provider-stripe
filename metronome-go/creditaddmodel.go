//
// File generated from our OpenAPI spec
//

package metronome

type CreditAddModelType string

// List of values that CreditAddModelType can take
const (
	CreditAddModelTypeCREDIT CreditAddModelType = "CREDIT"
)

type CreditAddModelRateType string

// List of values that CreditAddModelRateType can take
const (
	CreditAddModelRateTypeCOMMITRATE CreditAddModelRateType = "COMMIT_RATE"
	CreditAddModelRateTypeLISTRATE   CreditAddModelRateType = "LIST_RATE"
)

type CreditAddModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type CreditAddModel struct {
	APIResource
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        *ScheduleDurationModel `json:"access_schedule,omitempty"`
	ApplicableProductIDs  []string               `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string               `json:"applicable_product_tags,omitempty"`
	Description           string                 `json:"description,omitempty"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration *CommitHierarchyConfigurationModel `json:"hierarchy_configuration,omitempty"`
	ID                     string                             `json:"id"`
	Name                   string                             `json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits or commits are applicable, the one with the lower priority will apply first.
	Priority         Number                 `json:"priority,omitempty"`
	Product          *CreditAddModelProduct `json:"product"`
	RateType         CreditAddModelRateType `json:"rate_type,omitempty"`
	RolloverFraction Number                 `json:"rollover_fraction,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*CommitSpecifierInputModel `json:"specifiers,omitempty"`
	Type       CreditAddModelType           `json:"type"`
}
