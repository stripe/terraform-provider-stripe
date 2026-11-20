//
// File generated from our OpenAPI spec
//

package metronome

type ContractV2ModelScheduledChargesOnUsageInvoices string

// List of values that ContractV2ModelScheduledChargesOnUsageInvoices can take
const (
	ContractV2ModelScheduledChargesOnUsageInvoicesALL ContractV2ModelScheduledChargesOnUsageInvoices = "ALL"
)

type ContractV2ModelResellerRoyaltyResellerType string

// List of values that ContractV2ModelResellerRoyaltyResellerType can take
const (
	ContractV2ModelResellerRoyaltyResellerTypeAWS           ContractV2ModelResellerRoyaltyResellerType = "AWS"
	ContractV2ModelResellerRoyaltyResellerTypeAWSPROSERVICE ContractV2ModelResellerRoyaltyResellerType = "AWS_PRO_SERVICE"
	ContractV2ModelResellerRoyaltyResellerTypeGCP           ContractV2ModelResellerRoyaltyResellerType = "GCP"
	ContractV2ModelResellerRoyaltyResellerTypeGCPPROSERVICE ContractV2ModelResellerRoyaltyResellerType = "GCP_PRO_SERVICE"
)

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first.
type ContractV2ModelMultiplierOverridePrioritization string

// List of values that ContractV2ModelMultiplierOverridePrioritization can take
const (
	ContractV2ModelMultiplierOverridePrioritizationLOWESTMULTIPLIER ContractV2ModelMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	ContractV2ModelMultiplierOverridePrioritizationLowestMultiplier ContractV2ModelMultiplierOverridePrioritization = "lowest_multiplier"
	ContractV2ModelMultiplierOverridePrioritizationEXPLICIT         ContractV2ModelMultiplierOverridePrioritization = "EXPLICIT"
	ContractV2ModelMultiplierOverridePrioritizationExplicit         ContractV2ModelMultiplierOverridePrioritization = "explicit"
)

// This field's availability is dependent on your client's configuration.
type ContractV2ModelResellerRoyalty struct {
	ResellerType ContractV2ModelResellerRoyaltyResellerType `json:"reseller_type"`
	Segments     []*ResellerRoyaltyModel                    `json:"segments"`
}

// The schedule of billing provider configuration changes on the contract, ordered by effective_at ascending.
type ContractV2ModelBillingProviderConfigurationSchedule struct {
	BillingProviderConfiguration *CustomerBillingProviderConfigurationModel `json:"billing_provider_configuration"`
	// The date this billing provider configuration became or becomes active.
	EffectiveAt string `json:"effective_at"`
	// The date this billing provider configuration is superseded by the next entry. Null for the last entry in the schedule.
	EffectiveUntil string `json:"effective_until,omitempty"`
}

// The schedule of revenue system configuration changes on the contract, ordered by effective_at ascending.
type ContractV2ModelRevenueSystemConfigurationSchedule struct {
	// The date this revenue system configuration became or becomes active.
	EffectiveAt string `json:"effective_at"`
	// The date this revenue system configuration is superseded by the next entry. Null for the last entry in the schedule.
	EffectiveUntil             string                                     `json:"effective_until,omitempty"`
	RevenueSystemConfiguration *CustomerRevenueSystemConfigurationV2Model `json:"revenue_system_configuration"`
}
type ContractV2Model struct {
	APIResource
	ArchivedAt string `json:"archived_at,omitempty"`
	// The schedule of billing provider configuration changes on the contract, ordered by effective_at ascending.
	BillingProviderConfigurationSchedule []*ContractV2ModelBillingProviderConfigurationSchedule `json:"billing_provider_configuration_schedule,omitempty"`
	Commits                              []*CommitV2Model                                       `json:"commits"`
	CreatedAt                            string                                                 `json:"created_at"`
	CreatedBy                            string                                                 `json:"created_by"`
	Credits                              []*CreditV2Model                                       `json:"credits,omitempty"`
	CustomerBillingProviderConfiguration *CustomerBillingProviderConfigurationModel             `json:"customer_billing_provider_configuration,omitempty"`
	CustomerID                           string                                                 `json:"customer_id"`
	CustomFields                         map[string]string                                      `json:"custom_fields,omitempty"`
	// This field's availability is dependent on your client's configuration.
	Discounts              []*DiscountModel `json:"discounts,omitempty"`
	EndingBefore           string           `json:"ending_before,omitempty"`
	HasMore                *HasMoreModel    `json:"has_more,omitempty"`
	HierarchyConfiguration JSON             `json:"hierarchy_configuration,omitempty"`
	ID                     string           `json:"id"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first.
	MultiplierOverridePrioritization ContractV2ModelMultiplierOverridePrioritization `json:"multiplier_override_prioritization,omitempty"`
	Name                             string                                          `json:"name,omitempty"`
	NetPaymentTermsDays              Number                                          `json:"net_payment_terms_days,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string             `json:"netsuite_sales_order_id,omitempty"`
	Overrides            []*OverrideV2Model `json:"overrides"`
	// (BETA) ID of the package this contract was created from, if applicable.
	PackageID                            string                                       `json:"package_id,omitempty"`
	PrepaidBalanceThresholdConfiguration *PrepaidBalanceThresholdConfigurationV2Model `json:"prepaid_balance_threshold_configuration,omitempty"`
	// Priority of the contract.
	Priority Number `json:"priority,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []*ProServiceModel        `json:"professional_services,omitempty"`
	RateCardID           string                    `json:"rate_card_id,omitempty"`
	RecurringCommits     []*RecurringCommitV2Model `json:"recurring_commits,omitempty"`
	RecurringCredits     []*RecurringCreditV2Model `json:"recurring_credits,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []*ContractV2ModelResellerRoyalty `json:"reseller_royalties,omitempty"`
	// The schedule of revenue system configuration changes on the contract, ordered by effective_at ascending.
	RevenueSystemConfigurationSchedule []*ContractV2ModelRevenueSystemConfigurationSchedule `json:"revenue_system_configuration_schedule,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID         string                                         `json:"salesforce_opportunity_id,omitempty"`
	ScheduledCharges                []*ScheduledChargeModel                        `json:"scheduled_charges"`
	ScheduledChargesOnUsageInvoices ContractV2ModelScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices,omitempty"`
	SpendThresholdConfiguration     *SpendThresholdConfigurationV2Model            `json:"spend_threshold_configuration,omitempty"`
	// Spend trackers attached to this contract.
	SpendTrackers      []*SpendTrackerModel       `json:"spend_trackers,omitempty"`
	StartingAt         string                     `json:"starting_at"`
	Subscriptions      []*SubscriptionV2Model     `json:"subscriptions,omitempty"`
	TotalContractValue Number                     `json:"total_contract_value,omitempty"`
	Transitions        []*ContractTransitionModel `json:"transitions"`
	// Optional uniqueness key to prevent duplicate contract creations.
	UniquenessKey          string                       `json:"uniqueness_key,omitempty"`
	UsageFilter            []*UsageFilterV2Model        `json:"usage_filter"`
	UsageStatementSchedule *UsageStatementScheduleModel `json:"usage_statement_schedule"`
}
