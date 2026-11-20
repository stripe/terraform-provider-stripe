//
// File generated from our OpenAPI spec
//

package metronome

type ContractEditModelArchiveCommit struct {
	ID string `json:"id"`
}
type ContractEditModelArchiveCredit struct {
	ID string `json:"id"`
}
type ContractEditModelArchiveScheduledCharge struct {
	ID string `json:"id"`
}
type ContractEditModelRemoveOverride struct {
	ID string `json:"id"`
}
type ContractEditModel struct {
	APIResource
	AddCommits                              []*CommitAddModel                            `json:"add_commits,omitempty"`
	AddCredits                              []*CreditAddModel                            `json:"add_credits,omitempty"`
	AddDiscounts                            []*DiscountModel                             `json:"add_discounts,omitempty"`
	AddOverrides                            []*OverrideV2Model                           `json:"add_overrides,omitempty"`
	AddPrepaidBalanceThresholdConfiguration *PrepaidBalanceThresholdConfigurationV2Model `json:"add_prepaid_balance_threshold_configuration,omitempty"`
	AddProServices                          []*ProServiceModel                           `json:"add_pro_services,omitempty"`
	AddRecurringCommits                     []*RecurringCommitV2Model                    `json:"add_recurring_commits,omitempty"`
	AddRecurringCredits                     []*RecurringCreditV2Model                    `json:"add_recurring_credits,omitempty"`
	AddResellerRoyalties                    []*ResellerRoyaltyOrUpdateV2Model            `json:"add_reseller_royalties,omitempty"`
	AddScheduledCharges                     []*ScheduledChargeAddModel                   `json:"add_scheduled_charges,omitempty"`
	AddSpendThresholdConfiguration          *SpendThresholdConfigurationV2Model          `json:"add_spend_threshold_configuration,omitempty"`
	AddSubscriptions                        []*SubscriptionV2Model                       `json:"add_subscriptions,omitempty"`
	AddUsageFilters                         []*UsageFilterV2Model                        `json:"add_usage_filters,omitempty"`
	ArchiveCommits                          []*ContractEditModelArchiveCommit            `json:"archive_commits,omitempty"`
	ArchiveCredits                          []*ContractEditModelArchiveCredit            `json:"archive_credits,omitempty"`
	ArchiveScheduledCharges                 []*ContractEditModelArchiveScheduledCharge   `json:"archive_scheduled_charges,omitempty"`
	ID                                      string                                       `json:"id"`
	RemoveOverrides                         []*ContractEditModelRemoveOverride           `json:"remove_overrides,omitempty"`
	Timestamp                               string                                       `json:"timestamp,omitempty"`
	UniquenessKey                           string                                       `json:"uniqueness_key,omitempty"`
	UpdateCommits                           []*CommitUpdateModel                         `json:"update_commits,omitempty"`
	UpdateContractEndDate                   string                                       `json:"update_contract_end_date,omitempty"`
	// Value to update the contract name to. If not provided, the contract name will remain unchanged.
	UpdateContractName                         string                                           `json:"update_contract_name,omitempty"`
	UpdateCredits                              []*CreditUpdateModel                             `json:"update_credits,omitempty"`
	UpdateDiscounts                            []*DiscountUpdateModel                           `json:"update_discounts,omitempty"`
	UpdatePrepaidBalanceThresholdConfiguration *UpdatePrepaidBalanceThresholdConfigurationModel `json:"update_prepaid_balance_threshold_configuration,omitempty"`
	UpdateRecurringCommits                     []*RecurringCommitUpdateModel                    `json:"update_recurring_commits,omitempty"`
	UpdateRecurringCredits                     []*RecurringCreditUpdateModel                    `json:"update_recurring_credits,omitempty"`
	UpdateRefundInvoices                       []*RefundInvoiceUpdateModel                      `json:"update_refund_invoices,omitempty"`
	UpdateScheduledCharges                     []*ScheduledChargeUpdateModel                    `json:"update_scheduled_charges,omitempty"`
	UpdateSpendThresholdConfiguration          *UpdateSpendThresholdConfigurationModel          `json:"update_spend_threshold_configuration,omitempty"`
	// Optional list of subscriptions to update.
	UpdateSubscriptions []*SubscriptionsUpdateModel `json:"update_subscriptions,omitempty"`
}
