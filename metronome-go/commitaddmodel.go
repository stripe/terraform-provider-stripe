//
// File generated from our OpenAPI spec
//

package metronome

type CommitAddModelType string

// List of values that CommitAddModelType can take
const (
	CommitAddModelTypePREPAID  CommitAddModelType = "PREPAID"
	CommitAddModelTypePOSTPAID CommitAddModelType = "POSTPAID"
)

type CommitAddModelRateType string

// List of values that CommitAddModelRateType can take
const (
	CommitAddModelRateTypeCOMMITRATE CommitAddModelRateType = "COMMIT_RATE"
	CommitAddModelRateTypeLISTRATE   CommitAddModelRateType = "LIST_RATE"
)

type CommitAddModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type CommitAddModel struct {
	APIResource
	// The schedule that the customer will gain access to the credits purposed with this commit.
	AccessSchedule        *ScheduleDurationModel `json:"access_schedule,omitempty"`
	ApplicableProductIDs  []string               `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string               `json:"applicable_product_tags,omitempty"`
	Description           string                 `json:"description,omitempty"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration *CommitHierarchyConfigurationModel `json:"hierarchy_configuration,omitempty"`
	ID                     string                             `json:"id"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule *CommitAddInvoiceScheduleModel `json:"invoice_schedule,omitempty"`
	Name            string                         `json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits or commits are applicable, the one with the lower priority will apply first.
	Priority         Number                 `json:"priority,omitempty"`
	Product          *CommitAddModelProduct `json:"product"`
	RateType         CommitAddModelRateType `json:"rate_type,omitempty"`
	RolloverFraction Number                 `json:"rollover_fraction,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*CommitSpecifierInputModel `json:"specifiers,omitempty"`
	Type       CommitAddModelType           `json:"type"`
}
