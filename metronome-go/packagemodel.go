//
// File generated from our OpenAPI spec
//

package metronome

type PackageModelScheduledChargesOnUsageInvoices string

// List of values that PackageModelScheduledChargesOnUsageInvoices can take
const (
	PackageModelScheduledChargesOnUsageInvoicesALL PackageModelScheduledChargesOnUsageInvoices = "ALL"
)

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first.
type PackageModelMultiplierOverridePrioritization string

// List of values that PackageModelMultiplierOverridePrioritization can take
const (
	PackageModelMultiplierOverridePrioritizationLOWESTMULTIPLIER PackageModelMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	PackageModelMultiplierOverridePrioritizationEXPLICIT         PackageModelMultiplierOverridePrioritization = "EXPLICIT"
)

type PackageModelBillingProvider string

// List of values that PackageModelBillingProvider can take
const (
	PackageModelBillingProviderAwsMarketplace   PackageModelBillingProvider = "aws_marketplace"
	PackageModelBillingProviderStripe           PackageModelBillingProvider = "stripe"
	PackageModelBillingProviderNetsuite         PackageModelBillingProvider = "netsuite"
	PackageModelBillingProviderCustom           PackageModelBillingProvider = "custom"
	PackageModelBillingProviderAzureMarketplace PackageModelBillingProvider = "azure_marketplace"
	PackageModelBillingProviderQuickbooksOnline PackageModelBillingProvider = "quickbooks_online"
	PackageModelBillingProviderWorkday          PackageModelBillingProvider = "workday"
	PackageModelBillingProviderGcpMarketplace   PackageModelBillingProvider = "gcp_marketplace"
	PackageModelBillingProviderMetronome        PackageModelBillingProvider = "metronome"
)

type PackageModelDeliveryMethod string

// List of values that PackageModelDeliveryMethod can take
const (
	PackageModelDeliveryMethodDirectToBillingProvider PackageModelDeliveryMethod = "direct_to_billing_provider"
	PackageModelDeliveryMethodAwsSqs                  PackageModelDeliveryMethod = "aws_sqs"
	PackageModelDeliveryMethodTackle                  PackageModelDeliveryMethod = "tackle"
	PackageModelDeliveryMethodAwsSns                  PackageModelDeliveryMethod = "aws_sns"
)

type PackageModel struct {
	APIResource
	Aliases         []*PackageAliasModel        `json:"aliases,omitempty"`
	ArchivedAt      string                      `json:"archived_at,omitempty"`
	BillingProvider PackageModelBillingProvider `json:"billing_provider,omitempty"`
	Commits         []*CommitTemplateModel      `json:"commits"`
	// The name to use for contracts created from this package.
	ContractName   string                     `json:"contract_name,omitempty"`
	CreatedAt      string                     `json:"created_at"`
	CreatedBy      string                     `json:"created_by"`
	Credits        []*CreditTemplateModel     `json:"credits,omitempty"`
	DeliveryMethod PackageModelDeliveryMethod `json:"delivery_method,omitempty"`
	Duration       *RelativeDateModel         `json:"duration,omitempty"`
	ID             string                     `json:"id"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first.
	MultiplierOverridePrioritization     PackageModelMultiplierOverridePrioritization `json:"multiplier_override_prioritization,omitempty"`
	Name                                 string                                       `json:"name,omitempty"`
	NetPaymentTermsDays                  Number                                       `json:"net_payment_terms_days,omitempty"`
	Overrides                            []*OverrideTemplateModel                     `json:"overrides"`
	PrepaidBalanceThresholdConfiguration *PrepaidBalanceThresholdConfigurationModel   `json:"prepaid_balance_threshold_configuration,omitempty"`
	RateCardID                           string                                       `json:"rate_card_id,omitempty"`
	RecurringCommits                     []*RecurringCommitTemplateModel              `json:"recurring_commits,omitempty"`
	RecurringCredits                     []*RecurringCreditTemplateModel              `json:"recurring_credits,omitempty"`
	ScheduledCharges                     []*ScheduledChargeTemplateModel              `json:"scheduled_charges"`
	ScheduledChargesOnUsageInvoices      PackageModelScheduledChargesOnUsageInvoices  `json:"scheduled_charges_on_usage_invoices,omitempty"`
	SpendThresholdConfiguration          *SpendThresholdConfigurationModel            `json:"spend_threshold_configuration,omitempty"`
	SpendTrackers                        []*SpendTrackerTemplateModel                 `json:"spend_trackers,omitempty"`
	Subscriptions                        []*SubscriptionTemplateModel                 `json:"subscriptions,omitempty"`
	UniquenessKey                        string                                       `json:"uniqueness_key,omitempty"`
	UsageStatementSchedule               *UsageStatementScheduleTemplateModel         `json:"usage_statement_schedule"`
}
