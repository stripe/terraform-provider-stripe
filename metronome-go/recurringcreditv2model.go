//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type RecurringCreditV2ModelAccessAmountAccessType string

// List of values that RecurringCreditV2ModelAccessAmountAccessType can take
const (
	RecurringCreditV2ModelAccessAmountAccessTypeSPEND    RecurringCreditV2ModelAccessAmountAccessType = "SPEND"
	RecurringCreditV2ModelAccessAmountAccessTypeQUANTITY RecurringCreditV2ModelAccessAmountAccessType = "QUANTITY"
)

// Whether the created commits will use the commit rate or list rate
type RecurringCreditV2ModelRateType string

// List of values that RecurringCreditV2ModelRateType can take
const (
	RecurringCreditV2ModelRateTypeCOMMITRATE RecurringCreditV2ModelRateType = "COMMIT_RATE"
	RecurringCreditV2ModelRateTypeCommitRate RecurringCreditV2ModelRateType = "commit_rate"
	RecurringCreditV2ModelRateTypeLISTRATE   RecurringCreditV2ModelRateType = "LIST_RATE"
	RecurringCreditV2ModelRateTypeListRate   RecurringCreditV2ModelRateType = "list_rate"
)

type RecurringCreditV2ModelCommitDurationUnit string

// List of values that RecurringCreditV2ModelCommitDurationUnit can take
const (
	RecurringCreditV2ModelCommitDurationUnitPeriods RecurringCreditV2ModelCommitDurationUnit = "periods"
	RecurringCreditV2ModelCommitDurationUnitPERIODS RecurringCreditV2ModelCommitDurationUnit = "PERIODS"
)

// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
type RecurringCreditV2ModelRecurrenceFrequency string

// List of values that RecurringCreditV2ModelRecurrenceFrequency can take
const (
	RecurringCreditV2ModelRecurrenceFrequencyMONTHLY   RecurringCreditV2ModelRecurrenceFrequency = "MONTHLY"
	RecurringCreditV2ModelRecurrenceFrequencyMonthly   RecurringCreditV2ModelRecurrenceFrequency = "monthly"
	RecurringCreditV2ModelRecurrenceFrequencyQUARTERLY RecurringCreditV2ModelRecurrenceFrequency = "QUARTERLY"
	RecurringCreditV2ModelRecurrenceFrequencyQuarterly RecurringCreditV2ModelRecurrenceFrequency = "quarterly"
	RecurringCreditV2ModelRecurrenceFrequencyANNUAL    RecurringCreditV2ModelRecurrenceFrequency = "ANNUAL"
	RecurringCreditV2ModelRecurrenceFrequencyAnnual    RecurringCreditV2ModelRecurrenceFrequency = "annual"
	RecurringCreditV2ModelRecurrenceFrequencyWEEKLY    RecurringCreditV2ModelRecurrenceFrequency = "WEEKLY"
	RecurringCreditV2ModelRecurrenceFrequencyWeekly    RecurringCreditV2ModelRecurrenceFrequency = "weekly"
	RecurringCreditV2ModelRecurrenceFrequencyDAILY     RecurringCreditV2ModelRecurrenceFrequency = "DAILY"
	RecurringCreditV2ModelRecurrenceFrequencyDaily     RecurringCreditV2ModelRecurrenceFrequency = "daily"
)

// Determines whether the first and last commit will be prorated. If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type RecurringCreditV2ModelProration string

// List of values that RecurringCreditV2ModelProration can take
const (
	RecurringCreditV2ModelProrationNONE         RecurringCreditV2ModelProration = "NONE"
	RecurringCreditV2ModelProrationNone         RecurringCreditV2ModelProration = "none"
	RecurringCreditV2ModelProrationFIRST        RecurringCreditV2ModelProration = "FIRST"
	RecurringCreditV2ModelProrationFirst        RecurringCreditV2ModelProration = "first"
	RecurringCreditV2ModelProrationLAST         RecurringCreditV2ModelProration = "LAST"
	RecurringCreditV2ModelProrationLast         RecurringCreditV2ModelProration = "last"
	RecurringCreditV2ModelProrationFIRSTANDLAST RecurringCreditV2ModelProration = "FIRST_AND_LAST"
	RecurringCreditV2ModelProrationFirstAndLast RecurringCreditV2ModelProration = "first_and_last"
)

type RecurringCreditV2ModelContract struct {
	ID string `json:"id"`
}
type RecurringCreditV2ModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// The amount of commit to grant.
type RecurringCreditV2ModelAccessAmount struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType RecurringCreditV2ModelAccessAmountAccessType `json:"access_type,omitempty"`
	// This ID identifies the credit type for the access amount. Quantity-based recurring commits and credits return the null credit type UUID.
	CreditTypeID string `json:"credit_type_id"`
	Quantity     Number `json:"quantity,omitempty"`
	UnitPrice    Number `json:"unit_price"`
}

// The amount of time the created commits will be valid for
type RecurringCreditV2ModelCommitDuration struct {
	Unit  RecurringCreditV2ModelCommitDurationUnit `json:"unit,omitempty"`
	Value Number                                   `json:"value"`
}

// Rounding configuration for prorated recurring credit amounts.
type RecurringCreditV2ModelProrationRounding struct {
	Access *ProrationRoundingConfigV2Model `json:"access,omitempty"`
}
type RecurringCreditV2Model struct {
	APIResource
	// The amount of commit to grant.
	AccessAmount *RecurringCreditV2ModelAccessAmount `json:"access_amount"`
	// The date this recurring commit's billing periods are anchored to.
	AnchorDate string `json:"anchor_date"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	// The amount of time the created commits will be valid for
	CommitDuration *RecurringCreditV2ModelCommitDuration `json:"commit_duration"`
	Contract       *RecurringCreditV2ModelContract       `json:"contract,omitempty"`
	// Will be passed down to the individual commits
	Description string `json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore string `json:"ending_before,omitempty"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration *CommitHierarchyConfigurationModel `json:"hierarchy_configuration,omitempty"`
	ID                     string                             `json:"id"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority Number                         `json:"priority"`
	Product  *RecurringCreditV2ModelProduct `json:"product"`
	// Determines whether the first and last commit will be prorated. If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration RecurringCreditV2ModelProration `json:"proration,omitempty"`
	// Rounding configuration for prorated recurring credit amounts.
	ProrationRounding *RecurringCreditV2ModelProrationRounding `json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType RecurringCreditV2ModelRateType `json:"rate_type"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency RecurringCreditV2ModelRecurrenceFrequency `json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction Number `json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	// Determines the start time for the first commit
	StartingAt string `json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *RecurringCommitSubscriptionConfigModel `json:"subscription_config,omitempty"`
}
