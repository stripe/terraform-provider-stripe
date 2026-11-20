//
// File generated from our OpenAPI spec
//

package metronome

type CreditModelType string

// List of values that CreditModelType can take
const (
	CreditModelTypeCREDIT CreditModelType = "CREDIT"
)

type CreditModelSubscriptionConfigAllocation string

// List of values that CreditModelSubscriptionConfigAllocation can take
const (
	CreditModelSubscriptionConfigAllocationINDIVIDUAL CreditModelSubscriptionConfigAllocation = "INDIVIDUAL"
	CreditModelSubscriptionConfigAllocationPOOLED     CreditModelSubscriptionConfigAllocation = "POOLED"
)

// Controls when a customer can access this credit when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases the credit after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases the credit after the first payment.
type CreditModelSubscriptionConfigAccessPolicy string

// List of values that CreditModelSubscriptionConfigAccessPolicy can take
const (
	CreditModelSubscriptionConfigAccessPolicyBILLINGPERIODPAID            CreditModelSubscriptionConfigAccessPolicy = "BILLING_PERIOD_PAID"
	CreditModelSubscriptionConfigAccessPolicyINITIALBILLINGPERIODPAIDONLY CreditModelSubscriptionConfigAccessPolicy = "INITIAL_BILLING_PERIOD_PAID_ONLY"
)

type CreditModelRateType string

// List of values that CreditModelRateType can take
const (
	CreditModelRateTypeCOMMITRATE CreditModelRateType = "COMMIT_RATE"
	CreditModelRateTypeLISTRATE   CreditModelRateType = "LIST_RATE"
)

type CreditModelContract struct {
	ID string `json:"id"`
}
type CreditModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// The subscription configuration for this credit, if it was generated from a recurring credit with a subscription attached.
type CreditModelSubscriptionConfig struct {
	// Controls when a customer can access this credit when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases the credit after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases the credit after the first payment.
	AccessPolicy            CreditModelSubscriptionConfigAccessPolicy       `json:"access_policy,omitempty"`
	Allocation              CreditModelSubscriptionConfigAllocation         `json:"allocation,omitempty"`
	ApplySeatIncreaseConfig *ApplySeatIncreaseConfigForRecurringCommitModel `json:"apply_seat_increase_config,omitempty"`
	SubscriptionID          string                                          `json:"subscription_id,omitempty"`
}
type CreditModelRolledOverFrom struct {
	ContractID string `json:"contract_id"`
	CreditID   string `json:"credit_id"`
}
type CreditModel struct {
	APIResource
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        *ScheduleDurationModel `json:"access_schedule,omitempty"`
	ApplicableContractIDs []string               `json:"applicable_contract_ids,omitempty"`
	ApplicableProductIDs  []string               `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string               `json:"applicable_product_tags,omitempty"`
	Balance               Number                 `json:"balance,omitempty"`
	Contract              *CreditModelContract   `json:"contract,omitempty"`
	// Timestamp of when the credit was created.
	// - Recurring credit: latter of credit service period date and parent credit start date
	// - Rollover credit: when the new contract started
	//
	CreatedAt string `json:"created_at,omitempty"`
	// The actor who created this credit. Omitted for system-generated credits such as recurring credits.
	CreatedBy    string            `json:"created_by,omitempty"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
	Description  string            `json:"description,omitempty"`
	// Optional configuration for credit hierarchy access control
	HierarchyConfiguration *CommitHierarchyConfigurationModel `json:"hierarchy_configuration,omitempty"`
	ID                     string                             `json:"id"`
	// A list of ordered events that impact the balance of a credit. For example, an invoice deduction or an expiration.
	Ledger []JSON `json:"ledger,omitempty"`
	Name   string `json:"name,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	// If multiple credits or commits are applicable, the one with the lower priority will apply first.
	Priority Number              `json:"priority,omitempty"`
	Product  *CreditModelProduct `json:"product"`
	RateType CreditModelRateType `json:"rate_type,omitempty"`
	// The ID of the recurring credit that this credit was generated from, if applicable.
	RecurringCreditID string                     `json:"recurring_credit_id,omitempty"`
	RolledOverFrom    *CreditModelRolledOverFrom `json:"rolled_over_from,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	// The subscription configuration for this credit, if it was generated from a recurring credit with a subscription attached.
	SubscriptionConfig *CreditModelSubscriptionConfig `json:"subscription_config,omitempty"`
	Type               CreditModelType                `json:"type"`
	UniquenessKey      string                         `json:"uniqueness_key,omitempty"`
}
