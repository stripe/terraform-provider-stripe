//
// File generated from our OpenAPI spec
//

package metronome

type ContractWithoutAmendmentsModelScheduledChargesOnUsageInvoices string

// List of values that ContractWithoutAmendmentsModelScheduledChargesOnUsageInvoices can take
const (
	ContractWithoutAmendmentsModelScheduledChargesOnUsageInvoicesALL ContractWithoutAmendmentsModelScheduledChargesOnUsageInvoices = "ALL"
)

type ContractWithoutAmendmentsModelSpendTrackerResetFrequency string

// List of values that ContractWithoutAmendmentsModelSpendTrackerResetFrequency can take
const (
	ContractWithoutAmendmentsModelSpendTrackerResetFrequencyBILLINGPERIOD ContractWithoutAmendmentsModelSpendTrackerResetFrequency = "BILLING_PERIOD"
)

type ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSpendType string

// List of values that ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSpendType can take
const (
	ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSpendTypeCOMMITPURCHASE ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSpendType = "COMMIT_PURCHASE"
)

type ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSource string

// List of values that ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSource can take
const (
	ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSourceTHRESHOLDRECHARGE ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSource = "THRESHOLD_RECHARGE"
	ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSourceMANUAL            ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSource = "MANUAL"
)

type ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierDiscounted string

// List of values that ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierDiscounted can take
const (
	ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierDiscountedANY              ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierDiscounted = "ANY"
	ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierDiscountedDISCOUNTEDONLY   ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierDiscounted = "DISCOUNTED_ONLY"
	ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierDiscountedUNDISCOUNTEDONLY ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierDiscounted = "UNDISCOUNTED_ONLY"
)

type ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifier struct {
	Discounted ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierDiscounted `json:"discounted,omitempty"`
	Sources    []ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSource   `json:"sources"`
	SpendType  ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifierSpendType  `json:"spend_type"`
}
type ContractWithoutAmendmentsModelSpendTrackerAccumulatedSpend struct {
	Amount             Number `json:"amount"`
	PeriodEndingBefore string `json:"period_ending_before"`
	PeriodStartingAt   string `json:"period_starting_at"`
}

// Spend trackers attached to this contract.
type ContractWithoutAmendmentsModelSpendTracker struct {
	AccumulatedSpend *ContractWithoutAmendmentsModelSpendTrackerAccumulatedSpend `json:"accumulated_spend,omitempty"`
	// Human-readable identifier, unique per contract.
	Alias                     string                                                                `json:"alias"`
	ApplicableSpendSpecifiers []*ContractWithoutAmendmentsModelSpendTrackerApplicableSpendSpecifier `json:"applicable_spend_specifiers"`
	CreditTypeID              string                                                                `json:"credit_type_id"`
	ResetFrequency            ContractWithoutAmendmentsModelSpendTrackerResetFrequency              `json:"reset_frequency"`
}
type ContractWithoutAmendmentsModel struct {
	APIResource
	Commits   []*CommitModel `json:"commits"`
	CreatedAt string         `json:"created_at"`
	CreatedBy string         `json:"created_by"`
	Credits   []*CreditModel `json:"credits,omitempty"`
	// This field's availability is dependent on your client's
	Discounts              []*DiscountModel `json:"discounts,omitempty"`
	EndingBefore           string           `json:"ending_before,omitempty"`
	HierarchyConfiguration JSON             `json:"hierarchy_configuration,omitempty"`
	Name                   string           `json:"name,omitempty"`
	NetPaymentTermsDays    Number           `json:"net_payment_terms_days,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID                 string                                     `json:"netsuite_sales_order_id,omitempty"`
	Overrides                            []*OverrideModel                           `json:"overrides"`
	PrepaidBalanceThresholdConfiguration *PrepaidBalanceThresholdConfigurationModel `json:"prepaid_balance_threshold_configuration,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []*ProServiceModel      `json:"professional_services,omitempty"`
	RateCardID           string                  `json:"rate_card_id,omitempty"`
	RecurringCommits     []*RecurringCommitModel `json:"recurring_commits,omitempty"`
	RecurringCredits     []*RecurringCreditModel `json:"recurring_credits,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []*ResellerRoyaltyModel `json:"reseller_royalties,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID         string                                                        `json:"salesforce_opportunity_id,omitempty"`
	ScheduledCharges                []*ScheduledChargeModel                                       `json:"scheduled_charges"`
	ScheduledChargesOnUsageInvoices ContractWithoutAmendmentsModelScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices,omitempty"`
	SpendThresholdConfiguration     *SpendThresholdConfigurationModel                             `json:"spend_threshold_configuration,omitempty"`
	// Spend trackers attached to this contract.
	SpendTrackers []*ContractWithoutAmendmentsModelSpendTracker `json:"spend_trackers,omitempty"`
	StartingAt    string                                        `json:"starting_at"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue     Number                       `json:"total_contract_value,omitempty"`
	Transitions            []*ContractTransitionModel   `json:"transitions"`
	UsageFilter            *UsageFilterModel            `json:"usage_filter,omitempty"`
	UsageStatementSchedule *UsageStatementScheduleModel `json:"usage_statement_schedule"`
}
