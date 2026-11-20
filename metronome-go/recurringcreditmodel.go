//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type RecurringCreditModelAccessAmountAccessType string

// List of values that RecurringCreditModelAccessAmountAccessType can take
const (
	RecurringCreditModelAccessAmountAccessTypeSPEND    RecurringCreditModelAccessAmountAccessType = "SPEND"
	RecurringCreditModelAccessAmountAccessTypeQUANTITY RecurringCreditModelAccessAmountAccessType = "QUANTITY"
)

// Whether the created commits will use the commit rate or list rate
type RecurringCreditModelRateType string

// List of values that RecurringCreditModelRateType can take
const (
	RecurringCreditModelRateTypeCOMMITRATE RecurringCreditModelRateType = "COMMIT_RATE"
	RecurringCreditModelRateTypeCommitRate RecurringCreditModelRateType = "commit_rate"
	RecurringCreditModelRateTypeLISTRATE   RecurringCreditModelRateType = "LIST_RATE"
	RecurringCreditModelRateTypeListRate   RecurringCreditModelRateType = "list_rate"
)

type RecurringCreditModelCommitDurationUnit string

// List of values that RecurringCreditModelCommitDurationUnit can take
const (
	RecurringCreditModelCommitDurationUnitPeriods RecurringCreditModelCommitDurationUnit = "periods"
	RecurringCreditModelCommitDurationUnitPERIODS RecurringCreditModelCommitDurationUnit = "PERIODS"
)

// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
type RecurringCreditModelRecurrenceFrequency string

// List of values that RecurringCreditModelRecurrenceFrequency can take
const (
	RecurringCreditModelRecurrenceFrequencyMONTHLY   RecurringCreditModelRecurrenceFrequency = "MONTHLY"
	RecurringCreditModelRecurrenceFrequencyMonthly   RecurringCreditModelRecurrenceFrequency = "monthly"
	RecurringCreditModelRecurrenceFrequencyQUARTERLY RecurringCreditModelRecurrenceFrequency = "QUARTERLY"
	RecurringCreditModelRecurrenceFrequencyQuarterly RecurringCreditModelRecurrenceFrequency = "quarterly"
	RecurringCreditModelRecurrenceFrequencyANNUAL    RecurringCreditModelRecurrenceFrequency = "ANNUAL"
	RecurringCreditModelRecurrenceFrequencyAnnual    RecurringCreditModelRecurrenceFrequency = "annual"
	RecurringCreditModelRecurrenceFrequencyWEEKLY    RecurringCreditModelRecurrenceFrequency = "WEEKLY"
	RecurringCreditModelRecurrenceFrequencyWeekly    RecurringCreditModelRecurrenceFrequency = "weekly"
	RecurringCreditModelRecurrenceFrequencyDAILY     RecurringCreditModelRecurrenceFrequency = "DAILY"
	RecurringCreditModelRecurrenceFrequencyDaily     RecurringCreditModelRecurrenceFrequency = "daily"
)

// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type RecurringCreditModelProration string

// List of values that RecurringCreditModelProration can take
const (
	RecurringCreditModelProrationNONE         RecurringCreditModelProration = "NONE"
	RecurringCreditModelProrationNone         RecurringCreditModelProration = "none"
	RecurringCreditModelProrationFIRST        RecurringCreditModelProration = "FIRST"
	RecurringCreditModelProrationFirst        RecurringCreditModelProration = "first"
	RecurringCreditModelProrationLAST         RecurringCreditModelProration = "LAST"
	RecurringCreditModelProrationLast         RecurringCreditModelProration = "last"
	RecurringCreditModelProrationFIRSTANDLAST RecurringCreditModelProration = "FIRST_AND_LAST"
	RecurringCreditModelProrationFirstAndLast RecurringCreditModelProration = "first_and_last"
)

type RecurringCreditModelContract struct {
	ID string `json:"id"`
}
type RecurringCreditModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// The amount of commit to grant.
type RecurringCreditModelAccessAmount struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType RecurringCreditModelAccessAmountAccessType `json:"access_type,omitempty"`
	// This ID identifies the credit type for the access amount. Quantity-based recurring commits and credits return the null credit type UUID.
	CreditTypeID string `json:"credit_type_id"`
	Quantity     Number `json:"quantity,omitempty"`
	UnitPrice    Number `json:"unit_price"`
}

// The amount of time the created commits will be valid for
type RecurringCreditModelCommitDuration struct {
	Unit  RecurringCreditModelCommitDurationUnit `json:"unit,omitempty"`
	Value Number                                 `json:"value"`
}

// Rounding configuration for prorated recurring credit amounts.
type RecurringCreditModelProrationRounding struct {
	Access *ProrationRoundingConfigModel `json:"access,omitempty"`
}
type RecurringCreditModel struct {
	APIResource
	// The amount of commit to grant.
	AccessAmount *RecurringCreditModelAccessAmount `json:"access_amount"`
	// The date this recurring commit's billing periods are anchored to.
	AnchorDate string `json:"anchor_date"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	// The amount of time the created commits will be valid for
	CommitDuration *RecurringCreditModelCommitDuration `json:"commit_duration"`
	Contract       *RecurringCreditModelContract       `json:"contract,omitempty"`
	// Will be passed down to the individual commits
	Description string `json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore string `json:"ending_before,omitempty"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration *CommitHierarchyConfigurationModel `json:"hierarchy_configuration,omitempty"`
	ID                     string                             `json:"id"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority Number                       `json:"priority"`
	Product  *RecurringCreditModelProduct `json:"product"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration RecurringCreditModelProration `json:"proration,omitempty"`
	// Rounding configuration for prorated recurring credit amounts.
	ProrationRounding *RecurringCreditModelProrationRounding `json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType RecurringCreditModelRateType `json:"rate_type"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency RecurringCreditModelRecurrenceFrequency `json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction Number `json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	// Determines the start time for the first commit
	StartingAt string `json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *RecurringCommitSubscriptionConfigModel `json:"subscription_config,omitempty"`
}
