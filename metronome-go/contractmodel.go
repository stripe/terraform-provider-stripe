//
// File generated from our OpenAPI spec
//

package metronome

type ContractModelScheduledChargesOnUsageInvoices string

// List of values that ContractModelScheduledChargesOnUsageInvoices can take
const (
	ContractModelScheduledChargesOnUsageInvoicesALL ContractModelScheduledChargesOnUsageInvoices = "ALL"
)

type ContractModelSpendTrackerResetFrequency string

// List of values that ContractModelSpendTrackerResetFrequency can take
const (
	ContractModelSpendTrackerResetFrequencyBILLINGPERIOD ContractModelSpendTrackerResetFrequency = "BILLING_PERIOD"
)

type ContractModelSpendTrackerApplicableSpendSpecifierSpendType string

// List of values that ContractModelSpendTrackerApplicableSpendSpecifierSpendType can take
const (
	ContractModelSpendTrackerApplicableSpendSpecifierSpendTypeCOMMITPURCHASE ContractModelSpendTrackerApplicableSpendSpecifierSpendType = "COMMIT_PURCHASE"
)

type ContractModelSpendTrackerApplicableSpendSpecifierSource string

// List of values that ContractModelSpendTrackerApplicableSpendSpecifierSource can take
const (
	ContractModelSpendTrackerApplicableSpendSpecifierSourceTHRESHOLDRECHARGE ContractModelSpendTrackerApplicableSpendSpecifierSource = "THRESHOLD_RECHARGE"
	ContractModelSpendTrackerApplicableSpendSpecifierSourceMANUAL            ContractModelSpendTrackerApplicableSpendSpecifierSource = "MANUAL"
)

type ContractModelSpendTrackerApplicableSpendSpecifierDiscounted string

// List of values that ContractModelSpendTrackerApplicableSpendSpecifierDiscounted can take
const (
	ContractModelSpendTrackerApplicableSpendSpecifierDiscountedANY              ContractModelSpendTrackerApplicableSpendSpecifierDiscounted = "ANY"
	ContractModelSpendTrackerApplicableSpendSpecifierDiscountedDISCOUNTEDONLY   ContractModelSpendTrackerApplicableSpendSpecifierDiscounted = "DISCOUNTED_ONLY"
	ContractModelSpendTrackerApplicableSpendSpecifierDiscountedUNDISCOUNTEDONLY ContractModelSpendTrackerApplicableSpendSpecifierDiscounted = "UNDISCOUNTED_ONLY"
)

type ContractModelSpendTrackerApplicableSpendSpecifier struct {
	Discounted ContractModelSpendTrackerApplicableSpendSpecifierDiscounted `json:"discounted,omitempty"`
	Sources    []ContractModelSpendTrackerApplicableSpendSpecifierSource   `json:"sources"`
	SpendType  ContractModelSpendTrackerApplicableSpendSpecifierSpendType  `json:"spend_type"`
}
type ContractModelSpendTrackerAccumulatedSpend struct {
	Amount             Number `json:"amount"`
	PeriodEndingBefore string `json:"period_ending_before"`
	PeriodStartingAt   string `json:"period_starting_at"`
}

// Spend trackers attached to this contract.
type ContractModelSpendTracker struct {
	AccumulatedSpend *ContractModelSpendTrackerAccumulatedSpend `json:"accumulated_spend,omitempty"`
	// Human-readable identifier, unique per contract.
	Alias                     string                                               `json:"alias"`
	ApplicableSpendSpecifiers []*ContractModelSpendTrackerApplicableSpendSpecifier `json:"applicable_spend_specifiers"`
	CreditTypeID              string                                               `json:"credit_type_id"`
	ResetFrequency            ContractModelSpendTrackerResetFrequency              `json:"reset_frequency"`
}
type ContractModel struct {
	APIResource
	Amendments []*ContractAmendmentModel `json:"amendments"`
	// RFC 3339 timestamp indicating when the contract was archived. If not returned, the contract is not archived.
	ArchivedAt                           string                                     `json:"archived_at,omitempty"`
	Current                              *ContractWithoutAmendmentsModel            `json:"current"`
	CustomerBillingProviderConfiguration *CustomerBillingProviderConfigurationModel `json:"customer_billing_provider_configuration,omitempty"`
	CustomerID                           string                                     `json:"customer_id"`
	CustomFields                         map[string]string                          `json:"custom_fields,omitempty"`
	ID                                   string                                     `json:"id"`
	Initial                              *ContractWithoutAmendmentsModel            `json:"initial"`
	// ID of the package this contract was created from, if applicable.
	PackageID                            string                                       `json:"package_id,omitempty"`
	PrepaidBalanceThresholdConfiguration *PrepaidBalanceThresholdConfigurationModel   `json:"prepaid_balance_threshold_configuration,omitempty"`
	ScheduledChargesOnUsageInvoices      ContractModelScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices,omitempty"`
	SpendThresholdConfiguration          *SpendThresholdConfigurationModel            `json:"spend_threshold_configuration,omitempty"`
	// Spend trackers attached to this contract.
	SpendTrackers []*ContractModelSpendTracker `json:"spend_trackers,omitempty"`
	Subscriptions []*SubscriptionModel         `json:"subscriptions,omitempty"`
	UniquenessKey string                       `json:"uniqueness_key,omitempty"`
}
