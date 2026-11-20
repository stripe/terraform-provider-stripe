//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type RecurringCommitModelAccessAmountAccessType string

// List of values that RecurringCommitModelAccessAmountAccessType can take
const (
	RecurringCommitModelAccessAmountAccessTypeSPEND    RecurringCommitModelAccessAmountAccessType = "SPEND"
	RecurringCommitModelAccessAmountAccessTypeQUANTITY RecurringCommitModelAccessAmountAccessType = "QUANTITY"
)

// Whether the created commits will use the commit rate or list rate
type RecurringCommitModelRateType string

// List of values that RecurringCommitModelRateType can take
const (
	RecurringCommitModelRateTypeCOMMITRATE RecurringCommitModelRateType = "COMMIT_RATE"
	RecurringCommitModelRateTypeCommitRate RecurringCommitModelRateType = "commit_rate"
	RecurringCommitModelRateTypeLISTRATE   RecurringCommitModelRateType = "LIST_RATE"
	RecurringCommitModelRateTypeListRate   RecurringCommitModelRateType = "list_rate"
)

type RecurringCommitModelCommitDurationUnit string

// List of values that RecurringCommitModelCommitDurationUnit can take
const (
	RecurringCommitModelCommitDurationUnitPeriods RecurringCommitModelCommitDurationUnit = "periods"
	RecurringCommitModelCommitDurationUnitPERIODS RecurringCommitModelCommitDurationUnit = "PERIODS"
)

// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
type RecurringCommitModelRecurrenceFrequency string

// List of values that RecurringCommitModelRecurrenceFrequency can take
const (
	RecurringCommitModelRecurrenceFrequencyMONTHLY   RecurringCommitModelRecurrenceFrequency = "MONTHLY"
	RecurringCommitModelRecurrenceFrequencyMonthly   RecurringCommitModelRecurrenceFrequency = "monthly"
	RecurringCommitModelRecurrenceFrequencyQUARTERLY RecurringCommitModelRecurrenceFrequency = "QUARTERLY"
	RecurringCommitModelRecurrenceFrequencyQuarterly RecurringCommitModelRecurrenceFrequency = "quarterly"
	RecurringCommitModelRecurrenceFrequencyANNUAL    RecurringCommitModelRecurrenceFrequency = "ANNUAL"
	RecurringCommitModelRecurrenceFrequencyAnnual    RecurringCommitModelRecurrenceFrequency = "annual"
	RecurringCommitModelRecurrenceFrequencyWEEKLY    RecurringCommitModelRecurrenceFrequency = "WEEKLY"
	RecurringCommitModelRecurrenceFrequencyWeekly    RecurringCommitModelRecurrenceFrequency = "weekly"
	RecurringCommitModelRecurrenceFrequencyDAILY     RecurringCommitModelRecurrenceFrequency = "DAILY"
	RecurringCommitModelRecurrenceFrequencyDaily     RecurringCommitModelRecurrenceFrequency = "daily"
)

// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type RecurringCommitModelProration string

// List of values that RecurringCommitModelProration can take
const (
	RecurringCommitModelProrationNONE         RecurringCommitModelProration = "NONE"
	RecurringCommitModelProrationNone         RecurringCommitModelProration = "none"
	RecurringCommitModelProrationFIRST        RecurringCommitModelProration = "FIRST"
	RecurringCommitModelProrationFirst        RecurringCommitModelProration = "first"
	RecurringCommitModelProrationLAST         RecurringCommitModelProration = "LAST"
	RecurringCommitModelProrationLast         RecurringCommitModelProration = "last"
	RecurringCommitModelProrationFIRSTANDLAST RecurringCommitModelProration = "FIRST_AND_LAST"
	RecurringCommitModelProrationFirstAndLast RecurringCommitModelProration = "first_and_last"
)

type RecurringCommitModelContract struct {
	ID string `json:"id"`
}
type RecurringCommitModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// The amount of commit to grant.
type RecurringCommitModelAccessAmount struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType RecurringCommitModelAccessAmountAccessType `json:"access_type,omitempty"`
	// This ID identifies the credit type for the access amount. Quantity-based recurring commits and credits return the null credit type UUID.
	CreditTypeID string `json:"credit_type_id"`
	Quantity     Number `json:"quantity,omitempty"`
	UnitPrice    Number `json:"unit_price"`
}

// The amount of time the created commits will be valid for
type RecurringCommitModelCommitDuration struct {
	Unit  RecurringCommitModelCommitDurationUnit `json:"unit,omitempty"`
	Value Number                                 `json:"value"`
}

// The amount the customer should be billed for the commit. Not required.
type RecurringCommitModelInvoiceAmount struct {
	CreditTypeID string `json:"credit_type_id"`
	Quantity     Number `json:"quantity"`
	UnitPrice    Number `json:"unit_price"`
}

// Rounding configuration for prorated recurring commit amounts.
type RecurringCommitModelProrationRounding struct {
	Access  *ProrationRoundingConfigModel `json:"access,omitempty"`
	Invoice *ProrationRoundingConfigModel `json:"invoice,omitempty"`
}
type RecurringCommitModel struct {
	APIResource
	// The amount of commit to grant.
	AccessAmount *RecurringCommitModelAccessAmount `json:"access_amount"`
	// The date this recurring commit's billing periods are anchored to.
	AnchorDate string `json:"anchor_date"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	// The amount of time the created commits will be valid for
	CommitDuration *RecurringCommitModelCommitDuration `json:"commit_duration"`
	Contract       *RecurringCommitModelContract       `json:"contract,omitempty"`
	// Will be passed down to the individual commits
	Description string `json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore string `json:"ending_before,omitempty"`
	// Optional configuration for recurring commit/credit hierarchy access control
	HierarchyConfiguration *CommitHierarchyConfigurationModel `json:"hierarchy_configuration,omitempty"`
	ID                     string                             `json:"id"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount *RecurringCommitModelInvoiceAmount `json:"invoice_amount,omitempty"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority Number                       `json:"priority"`
	Product  *RecurringCommitModelProduct `json:"product"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration RecurringCommitModelProration `json:"proration,omitempty"`
	// Rounding configuration for prorated recurring commit amounts.
	ProrationRounding *RecurringCommitModelProrationRounding `json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType RecurringCommitModelRateType `json:"rate_type"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency RecurringCommitModelRecurrenceFrequency `json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction Number `json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	// Determines the start time for the first commit
	StartingAt string `json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *RecurringCommitSubscriptionConfigModel `json:"subscription_config,omitempty"`
}
