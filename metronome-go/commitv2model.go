//
// File generated from our OpenAPI spec
//

package metronome

type CommitV2ModelType string

// List of values that CommitV2ModelType can take
const (
	CommitV2ModelTypePREPAID  CommitV2ModelType = "PREPAID"
	CommitV2ModelTypePOSTPAID CommitV2ModelType = "POSTPAID"
)

type CommitV2ModelRateType string

// List of values that CommitV2ModelRateType can take
const (
	CommitV2ModelRateTypeCOMMITRATE CommitV2ModelRateType = "COMMIT_RATE"
	CommitV2ModelRateTypeLISTRATE   CommitV2ModelRateType = "LIST_RATE"
)

type CommitV2ModelContract struct {
	ID string `json:"id"`
}
type CommitV2ModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// The contract that this commit will be billed on.
type CommitV2ModelInvoiceContract struct {
	ID string `json:"id"`
}
type CommitV2ModelRolledOverFrom struct {
	CommitID   string `json:"commit_id"`
	ContractID string `json:"contract_id"`
}
type CommitV2Model struct {
	APIResource
	// The schedule that the customer will gain access to the credits purposed with this commit.
	AccessSchedule        *ScheduleDurationModel `json:"access_schedule,omitempty"`
	ApplicableContractIDs []string               `json:"applicable_contract_ids,omitempty"`
	ApplicableProductIDs  []string               `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string               `json:"applicable_product_tags,omitempty"`
	ArchivedAt            string                 `json:"archived_at,omitempty"`
	Balance               Number                 `json:"balance,omitempty"`
	Contract              *CommitV2ModelContract `json:"contract,omitempty"`
	// The ratio of the amount paid for the commit to the amount of credit granted.
	CostBasis Number `json:"cost_basis,omitempty"`
	// Timestamp of when the commit was created.
	// - Recurring commits: latter of commit service period date and parent commit start date
	// - Rollover commits: when the new contract started
	//
	CreatedAt string `json:"created_at"`
	// The actor who created this commit. Omitted for system-generated commits such as recurring commits, rollover commits, and threshold commits.
	CreatedBy    string            `json:"created_by,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
	Description  string            `json:"description,omitempty"`
	// Optional configuration for commit hierarchy access control
	HierarchyConfiguration *CommitHierarchyConfigurationModel `json:"hierarchy_configuration,omitempty"`
	ID                     string                             `json:"id"`
	// The contract that this commit will be billed on.
	InvoiceContract *CommitV2ModelInvoiceContract `json:"invoice_contract,omitempty"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule *SchedulePointInTimeModel `json:"invoice_schedule,omitempty"`
	// A list of ordered events that impact the balance of a commit. For example, an invoice deduction or a rollover.
	Ledger []JSON `json:"ledger,omitempty"`
	Name   string `json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits or commits are applicable, the one with the lower priority will apply first.
	Priority Number                `json:"priority,omitempty"`
	Product  *CommitV2ModelProduct `json:"product"`
	RateType CommitV2ModelRateType `json:"rate_type,omitempty"`
	// The ID of the recurring commit that created this commit
	RecurringCommitID string                       `json:"recurring_commit_id,omitempty"`
	RolledOverFrom    *CommitV2ModelRolledOverFrom `json:"rolled_over_from,omitempty"`
	RolloverFraction  Number                       `json:"rollover_fraction,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	// Optional attributes controlling how this commit interacts with spend trackers.
	SpendTrackerAttributes *SpendTrackerAttributesModel `json:"spend_tracker_attributes,omitempty"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *RecurringCommitSubscriptionConfigModel `json:"subscription_config,omitempty"`
	Type               CommitV2ModelType                       `json:"type"`
}
