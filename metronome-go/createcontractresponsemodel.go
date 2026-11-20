//
// File generated from our OpenAPI spec
//

package metronome

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first.
type CreateContractResponseModelContractMultiplierOverridePrioritization string

// List of values that CreateContractResponseModelContractMultiplierOverridePrioritization can take
const (
	CreateContractResponseModelContractMultiplierOverridePrioritizationLOWESTMULTIPLIER CreateContractResponseModelContractMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	CreateContractResponseModelContractMultiplierOverridePrioritizationLowestMultiplier CreateContractResponseModelContractMultiplierOverridePrioritization = "lowest_multiplier"
	CreateContractResponseModelContractMultiplierOverridePrioritizationEXPLICIT         CreateContractResponseModelContractMultiplierOverridePrioritization = "EXPLICIT"
	CreateContractResponseModelContractMultiplierOverridePrioritizationExplicit         CreateContractResponseModelContractMultiplierOverridePrioritization = "explicit"
)

type CreateContractResponseModelContractScheduledChargesOnUsageInvoices string

// List of values that CreateContractResponseModelContractScheduledChargesOnUsageInvoices can take
const (
	CreateContractResponseModelContractScheduledChargesOnUsageInvoicesALL CreateContractResponseModelContractScheduledChargesOnUsageInvoices = "ALL"
)

type CreateContractResponseModelContractUsageFilter struct {
	EndingBefore string   `json:"ending_before,omitempty"`
	GroupKey     string   `json:"group_key"`
	GroupValues  []string `json:"group_values"`
	StartingAt   string   `json:"starting_at"`
}

// The created contract.
type CreateContractResponseModelContract struct {
	Commits                              []*CommitModel                             `json:"commits"`
	CreatedAt                            string                                     `json:"created_at"`
	CreatedBy                            string                                     `json:"created_by"`
	Credits                              []*CreditModel                             `json:"credits,omitempty"`
	CustomerBillingProviderConfiguration *CustomerBillingProviderConfigurationModel `json:"customer_billing_provider_configuration,omitempty"`
	CustomerID                           string                                     `json:"customer_id"`
	CustomFields                         map[string]string                          `json:"custom_fields,omitempty"`
	EndingBefore                         string                                     `json:"ending_before,omitempty"`
	HasMore                              *HasMoreModel                              `json:"has_more,omitempty"`
	HierarchyConfiguration               JSON                                       `json:"hierarchy_configuration,omitempty"`
	ID                                   string                                     `json:"id"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first.
	MultiplierOverridePrioritization CreateContractResponseModelContractMultiplierOverridePrioritization `json:"multiplier_override_prioritization,omitempty"`
	Name                             string                                                              `json:"name,omitempty"`
	NetPaymentTermsDays              Number                                                              `json:"net_payment_terms_days,omitempty"`
	Overrides                        []*OverrideModel                                                    `json:"overrides"`
	// ID of the package this contract was created from, if applicable.
	PackageID                            string                                                             `json:"package_id,omitempty"`
	PrepaidBalanceThresholdConfiguration *PrepaidBalanceThresholdConfigurationModel                         `json:"prepaid_balance_threshold_configuration,omitempty"`
	RateCardID                           string                                                             `json:"rate_card_id,omitempty"`
	RecurringCommits                     []*RecurringCommitModel                                            `json:"recurring_commits,omitempty"`
	RecurringCredits                     []*RecurringCreditModel                                            `json:"recurring_credits,omitempty"`
	ScheduledCharges                     []*ScheduledChargeModel                                            `json:"scheduled_charges"`
	ScheduledChargesOnUsageInvoices      CreateContractResponseModelContractScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices,omitempty"`
	SpendThresholdConfiguration          *SpendThresholdConfigurationModel                                  `json:"spend_threshold_configuration,omitempty"`
	StartingAt                           string                                                             `json:"starting_at"`
	Subscriptions                        []*SubscriptionModel                                               `json:"subscriptions,omitempty"`
	Transitions                          []*ContractTransitionModel                                         `json:"transitions"`
	// Optional uniqueness key to prevent duplicate contract creations.
	UniquenessKey          string                                            `json:"uniqueness_key,omitempty"`
	UsageFilter            []*CreateContractResponseModelContractUsageFilter `json:"usage_filter"`
	UsageStatementSchedule *UsageStatementScheduleModel                      `json:"usage_statement_schedule"`
}
type CreateContractResponseModel struct {
	APIResource
	// The created contract.
	Contract *CreateContractResponseModelContract `json:"contract,omitempty"`
	ID       string                               `json:"id"`
}
