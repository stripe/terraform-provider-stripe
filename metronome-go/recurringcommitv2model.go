//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type RecurringCommitV2ModelAccessAmountAccessType string

// List of values that RecurringCommitV2ModelAccessAmountAccessType can take
const (
	RecurringCommitV2ModelAccessAmountAccessTypeSPEND    RecurringCommitV2ModelAccessAmountAccessType = "SPEND"
	RecurringCommitV2ModelAccessAmountAccessTypeQUANTITY RecurringCommitV2ModelAccessAmountAccessType = "QUANTITY"
)

// Whether the created commits will use the commit rate or list rate
type RecurringCommitV2ModelRateType string

// List of values that RecurringCommitV2ModelRateType can take
const (
	RecurringCommitV2ModelRateTypeCOMMITRATE RecurringCommitV2ModelRateType = "COMMIT_RATE"
	RecurringCommitV2ModelRateTypeCommitRate RecurringCommitV2ModelRateType = "commit_rate"
	RecurringCommitV2ModelRateTypeLISTRATE   RecurringCommitV2ModelRateType = "LIST_RATE"
	RecurringCommitV2ModelRateTypeListRate   RecurringCommitV2ModelRateType = "list_rate"
)

type RecurringCommitV2ModelCommitDurationUnit string

// List of values that RecurringCommitV2ModelCommitDurationUnit can take
const (
	RecurringCommitV2ModelCommitDurationUnitPeriods RecurringCommitV2ModelCommitDurationUnit = "periods"
	RecurringCommitV2ModelCommitDurationUnitPERIODS RecurringCommitV2ModelCommitDurationUnit = "PERIODS"
)

// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
type RecurringCommitV2ModelRecurrenceFrequency string

// List of values that RecurringCommitV2ModelRecurrenceFrequency can take
const (
	RecurringCommitV2ModelRecurrenceFrequencyMONTHLY   RecurringCommitV2ModelRecurrenceFrequency = "MONTHLY"
	RecurringCommitV2ModelRecurrenceFrequencyMonthly   RecurringCommitV2ModelRecurrenceFrequency = "monthly"
	RecurringCommitV2ModelRecurrenceFrequencyQUARTERLY RecurringCommitV2ModelRecurrenceFrequency = "QUARTERLY"
	RecurringCommitV2ModelRecurrenceFrequencyQuarterly RecurringCommitV2ModelRecurrenceFrequency = "quarterly"
	RecurringCommitV2ModelRecurrenceFrequencyANNUAL    RecurringCommitV2ModelRecurrenceFrequency = "ANNUAL"
	RecurringCommitV2ModelRecurrenceFrequencyAnnual    RecurringCommitV2ModelRecurrenceFrequency = "annual"
	RecurringCommitV2ModelRecurrenceFrequencyWEEKLY    RecurringCommitV2ModelRecurrenceFrequency = "WEEKLY"
	RecurringCommitV2ModelRecurrenceFrequencyWeekly    RecurringCommitV2ModelRecurrenceFrequency = "weekly"
	RecurringCommitV2ModelRecurrenceFrequencyDAILY     RecurringCommitV2ModelRecurrenceFrequency = "DAILY"
	RecurringCommitV2ModelRecurrenceFrequencyDaily     RecurringCommitV2ModelRecurrenceFrequency = "daily"
)

// Determines whether the first and last commit will be prorated. If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type RecurringCommitV2ModelProration string

// List of values that RecurringCommitV2ModelProration can take
const (
	RecurringCommitV2ModelProrationNONE         RecurringCommitV2ModelProration = "NONE"
	RecurringCommitV2ModelProrationNone         RecurringCommitV2ModelProration = "none"
	RecurringCommitV2ModelProrationFIRST        RecurringCommitV2ModelProration = "FIRST"
	RecurringCommitV2ModelProrationFirst        RecurringCommitV2ModelProration = "first"
	RecurringCommitV2ModelProrationLAST         RecurringCommitV2ModelProration = "LAST"
	RecurringCommitV2ModelProrationLast         RecurringCommitV2ModelProration = "last"
	RecurringCommitV2ModelProrationFIRSTANDLAST RecurringCommitV2ModelProration = "FIRST_AND_LAST"
	RecurringCommitV2ModelProrationFirstAndLast RecurringCommitV2ModelProration = "first_and_last"
)

type RecurringCommitV2ModelContract struct {
	ID string `json:"id"`
}
type RecurringCommitV2ModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// The amount of commit to grant.
type RecurringCommitV2ModelAccessAmount struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType RecurringCommitV2ModelAccessAmountAccessType `json:"access_type,omitempty"`
	// This ID identifies the credit type for the access amount. Quantity-based recurring commits and credits return the null credit type UUID.
	CreditTypeID string `json:"credit_type_id"`
	Quantity     Number `json:"quantity,omitempty"`
	UnitPrice    Number `json:"unit_price"`
}

// The amount of time the created commits will be valid for
type RecurringCommitV2ModelCommitDuration struct {
	Unit  RecurringCommitV2ModelCommitDurationUnit `json:"unit,omitempty"`
	Value Number                                   `json:"value"`
}

// The amount the customer should be billed for the commit. Not required.
type RecurringCommitV2ModelInvoiceAmount struct {
	CreditTypeID string `json:"credit_type_id"`
	Quantity     Number `json:"quantity"`
	UnitPrice    Number `json:"unit_price"`
}

// Rounding configuration for prorated recurring commit amounts.
type RecurringCommitV2ModelProrationRounding struct {
	Access  *ProrationRoundingConfigV2Model `json:"access,omitempty"`
	Invoice *ProrationRoundingConfigV2Model `json:"invoice,omitempty"`
}
type RecurringCommitV2Model struct {
	APIResource
	// The amount of commit to grant.
	AccessAmount *RecurringCommitV2ModelAccessAmount `json:"access_amount"`
	// The date this recurring commit's billing periods are anchored to.
	AnchorDate string `json:"anchor_date"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	// The amount of time the created commits will be valid for
	CommitDuration *RecurringCommitV2ModelCommitDuration `json:"commit_duration"`
	Contract       *RecurringCommitV2ModelContract       `json:"contract,omitempty"`
	// Will be passed down to the individual commits
	Description string `json:"description,omitempty"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore string `json:"ending_before,omitempty"`
	// Optional configuration for recurring credit hierarchy access control
	HierarchyConfiguration *CommitHierarchyConfigurationModel `json:"hierarchy_configuration,omitempty"`
	ID                     string                             `json:"id"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount *RecurringCommitV2ModelInvoiceAmount `json:"invoice_amount,omitempty"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name,omitempty"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	// Will be passed down to the individual commits
	Priority Number                         `json:"priority"`
	Product  *RecurringCommitV2ModelProduct `json:"product"`
	// Determines whether the first and last commit will be prorated. If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration RecurringCommitV2ModelProration `json:"proration,omitempty"`
	// Rounding configuration for prorated recurring commit amounts.
	ProrationRounding *RecurringCommitV2ModelProrationRounding `json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType RecurringCommitV2ModelRateType `json:"rate_type"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency RecurringCommitV2ModelRecurrenceFrequency `json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction Number `json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	// Determines the start time for the first commit
	StartingAt string `json:"starting_at"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *RecurringCommitSubscriptionConfigModel `json:"subscription_config,omitempty"`
}
