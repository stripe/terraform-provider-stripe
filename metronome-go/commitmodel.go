//
// File generated from our OpenAPI spec
//

package metronome

type CommitModelType string

// List of values that CommitModelType can take
const (
	CommitModelTypePREPAID  CommitModelType = "PREPAID"
	CommitModelTypePOSTPAID CommitModelType = "POSTPAID"
)

type CommitModelRateType string

// List of values that CommitModelRateType can take
const (
	CommitModelRateTypeCOMMITRATE CommitModelRateType = "COMMIT_RATE"
	CommitModelRateTypeLISTRATE   CommitModelRateType = "LIST_RATE"
)

type CommitModelSubscriptionConfigAllocation string

// List of values that CommitModelSubscriptionConfigAllocation can take
const (
	CommitModelSubscriptionConfigAllocationINDIVIDUAL CommitModelSubscriptionConfigAllocation = "INDIVIDUAL"
	CommitModelSubscriptionConfigAllocationPOOLED     CommitModelSubscriptionConfigAllocation = "POOLED"
)

// Controls when a customer can access this commit when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases the commit after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases the commit after the first payment.
type CommitModelSubscriptionConfigAccessPolicy string

// List of values that CommitModelSubscriptionConfigAccessPolicy can take
const (
	CommitModelSubscriptionConfigAccessPolicyBILLINGPERIODPAID            CommitModelSubscriptionConfigAccessPolicy = "BILLING_PERIOD_PAID"
	CommitModelSubscriptionConfigAccessPolicyINITIALBILLINGPERIODPAIDONLY CommitModelSubscriptionConfigAccessPolicy = "INITIAL_BILLING_PERIOD_PAID_ONLY"
)

type CommitModelContract struct {
	ID string `json:"id"`
}
type CommitModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// The contract that this commit will be billed on.
type CommitModelInvoiceContract struct {
	ID string `json:"id"`
}

// The subscription configuration for this commit, if it was generated from a recurring commit with a subscription attached.
type CommitModelSubscriptionConfig struct {
	// Controls when a customer can access this commit when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases the commit after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases the commit after the first payment.
	AccessPolicy            CommitModelSubscriptionConfigAccessPolicy       `json:"access_policy,omitempty"`
	Allocation              CommitModelSubscriptionConfigAllocation         `json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *ApplySeatIncreaseConfigForRecurringCommitModel `json:"apply_seat_increase_config,omitempty"`
	SubscriptionID          string                                          `json:"subscription_id,omitempty"`
}
type CommitModelRolledOverFrom struct {
	CommitID   string `json:"commit_id"`
	ContractID string `json:"contract_id"`
}
type CommitModel struct {
	APIResource
	// The schedule that the customer will gain access to the credits purposed with this commit.
	AccessSchedule *ScheduleDurationModel `json:"access_schedule,omitempty"`
	// (DEPRECATED) Use access_schedule + invoice_schedule instead.
	Amount                Number   `json:"amount,omitempty"`
	ApplicableContractIDs []string `json:"applicable_contract_ids,omitempty"`
	ApplicableProductIDs  []string `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	// RFC 3339 timestamp indicating when the commit was archived. If not provided, the commit is not archived.
	ArchivedAt string               `json:"archived_at,omitempty"`
	Balance    Number               `json:"balance,omitempty"`
	Contract   *CommitModelContract `json:"contract,omitempty"`
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
	InvoiceContract *CommitModelInvoiceContract `json:"invoice_contract,omitempty"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule *SchedulePointInTimeModel `json:"invoice_schedule,omitempty"`
	// A list of ordered events that impact the balance of a commit. For example, an invoice deduction or a rollover.
	Ledger []JSON `json:"ledger,omitempty"`
	Name   string `json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits or commits are applicable, the one with the lower priority will apply first.
	Priority Number              `json:"priority,omitempty"`
	Product  *CommitModelProduct `json:"product"`
	RateType CommitModelRateType `json:"rate_type,omitempty"`
	// The ID of the recurring commit that this commit was generated from, if applicable.
	RecurringCommitID string                     `json:"recurring_commit_id,omitempty"`
	RolledOverFrom    *CommitModelRolledOverFrom `json:"rolled_over_from,omitempty"`
	RolloverFraction  Number                     `json:"rollover_fraction,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	// Optional attributes controlling how this commit interacts with spend trackers.
	SpendTrackerAttributes *SpendTrackerAttributesModel `json:"spend_tracker_attributes,omitempty"`
	// The subscription configuration for this commit, if it was generated from a recurring commit with a subscription attached.
	SubscriptionConfig *CommitModelSubscriptionConfig `json:"subscription_config,omitempty"`
	Type               CommitModelType                `json:"type"`
	UniquenessKey      string                         `json:"uniqueness_key,omitempty"`
}
