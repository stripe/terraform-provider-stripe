//
// File generated from our OpenAPI spec
//

package metronome

type CreditV2ModelType string

// List of values that CreditV2ModelType can take
const (
	CreditV2ModelTypeCREDIT CreditV2ModelType = "CREDIT"
)

type CreditV2ModelContract struct {
	ID string `json:"id"`
}
type CreditV2ModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type CreditV2ModelRolledOverFrom struct {
	ContractID string `json:"contract_id"`
	CreditID   string `json:"credit_id"`
}
type CreditV2Model struct {
	APIResource
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        *ScheduleDurationModel `json:"access_schedule,omitempty"`
	ApplicableContractIDs []string               `json:"applicable_contract_ids,omitempty"`
	ApplicableProductIDs  []string               `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string               `json:"applicable_product_tags,omitempty"`
	ArchivedAt            string                 `json:"archived_at,omitempty"`
	Balance               Number                 `json:"balance,omitempty"`
	Contract              *CreditV2ModelContract `json:"contract,omitempty"`
	// Timestamp of when the credit was created.
	// - Recurring credits: latter of credit service period date and parent credit start date
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
	Priority Number                `json:"priority,omitempty"`
	Product  *CreditV2ModelProduct `json:"product"`
	// The ID of the recurring credit that created this credit
	RecurringCreditID string                       `json:"recurring_credit_id,omitempty"`
	RolledOverFrom    *CreditV2ModelRolledOverFrom `json:"rolled_over_from,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *RecurringCommitSubscriptionConfigModel `json:"subscription_config,omitempty"`
	Type               CreditV2ModelType                       `json:"type"`
}
