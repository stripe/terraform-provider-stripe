//
// File generated from our OpenAPI spec
//

package metronome

// If set, the commit's rate type was updated to the specified value.
type CommitUpdateModelRateType string

// List of values that CommitUpdateModelRateType can take
const (
	CommitUpdateModelRateTypeCOMMITRATE CommitUpdateModelRateType = "COMMIT_RATE"
	CommitUpdateModelRateTypeLISTRATE   CommitUpdateModelRateType = "LIST_RATE"
)

type CommitUpdateModel struct {
	APIResource
	AccessSchedule *AccessScheduleUpdateModel `json:"access_schedule,omitempty"`
	// Which products the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductIDs []string `json:"applicable_product_ids,omitempty"`
	// Which tags the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	Description           string   `json:"description,omitempty"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration *CommitHierarchyConfigurationModel `json:"hierarchy_configuration,omitempty"`
	ID                     string                             `json:"id"`
	InvoiceSchedule        *InvoiceScheduleUpdateModel        `json:"invoice_schedule,omitempty"`
	Name                   string                             `json:"name,omitempty"`
	NetsuiteSalesOrderID   string                             `json:"netsuite_sales_order_id,omitempty"`
	// If multiple commits are applicable, the one with the lower priority will apply first.
	Priority  Number `json:"priority,omitempty"`
	ProductID string `json:"product_id,omitempty"`
	// If set, the commit's rate type was updated to the specified value.
	RateType         CommitUpdateModelRateType `json:"rate_type,omitempty"`
	RolloverFraction Number                    `json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`. Instead, to target usage by product or product tag, pass those values in the body of `specifiers`.
	Specifiers []*CommitSpecifierInputModel `json:"specifiers,omitempty"`
}
