//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type RecurringCommitTemplateModelAccessAmountAccessType string

// List of values that RecurringCommitTemplateModelAccessAmountAccessType can take
const (
	RecurringCommitTemplateModelAccessAmountAccessTypeSPEND    RecurringCommitTemplateModelAccessAmountAccessType = "SPEND"
	RecurringCommitTemplateModelAccessAmountAccessTypeQUANTITY RecurringCommitTemplateModelAccessAmountAccessType = "QUANTITY"
)

// Whether the created commits will use the commit rate or list rate
type RecurringCommitTemplateModelRateType string

// List of values that RecurringCommitTemplateModelRateType can take
const (
	RecurringCommitTemplateModelRateTypeCOMMITRATE RecurringCommitTemplateModelRateType = "COMMIT_RATE"
	RecurringCommitTemplateModelRateTypeLISTRATE   RecurringCommitTemplateModelRateType = "LIST_RATE"
)

type RecurringCommitTemplateModelCommitDurationUnit string

// List of values that RecurringCommitTemplateModelCommitDurationUnit can take
const (
	RecurringCommitTemplateModelCommitDurationUnitPERIODS RecurringCommitTemplateModelCommitDurationUnit = "PERIODS"
)

// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
type RecurringCommitTemplateModelRecurrenceFrequency string

// List of values that RecurringCommitTemplateModelRecurrenceFrequency can take
const (
	RecurringCommitTemplateModelRecurrenceFrequencyMONTHLY   RecurringCommitTemplateModelRecurrenceFrequency = "MONTHLY"
	RecurringCommitTemplateModelRecurrenceFrequencyQUARTERLY RecurringCommitTemplateModelRecurrenceFrequency = "QUARTERLY"
	RecurringCommitTemplateModelRecurrenceFrequencyANNUAL    RecurringCommitTemplateModelRecurrenceFrequency = "ANNUAL"
	RecurringCommitTemplateModelRecurrenceFrequencyWEEKLY    RecurringCommitTemplateModelRecurrenceFrequency = "WEEKLY"
	RecurringCommitTemplateModelRecurrenceFrequencyDAILY     RecurringCommitTemplateModelRecurrenceFrequency = "DAILY"
)

// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type RecurringCommitTemplateModelProration string

// List of values that RecurringCommitTemplateModelProration can take
const (
	RecurringCommitTemplateModelProrationNONE         RecurringCommitTemplateModelProration = "NONE"
	RecurringCommitTemplateModelProrationFIRST        RecurringCommitTemplateModelProration = "FIRST"
	RecurringCommitTemplateModelProrationLAST         RecurringCommitTemplateModelProration = "LAST"
	RecurringCommitTemplateModelProrationFIRSTANDLAST RecurringCommitTemplateModelProration = "FIRST_AND_LAST"
)

type RecurringCommitTemplateModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// The amount of commit to grant.
type RecurringCommitTemplateModelAccessAmount struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType RecurringCommitTemplateModelAccessAmountAccessType `json:"access_type,omitempty"`
	// This ID identifies the credit type for the access amount. Quantity-based recurring commits and credits return the null credit type UUID.
	CreditTypeID string `json:"credit_type_id"`
	Quantity     Number `json:"quantity,omitempty"`
	UnitPrice    Number `json:"unit_price"`
}

// The amount of time each of the created commits will be valid for
type RecurringCommitTemplateModelCommitDuration struct {
	Unit  RecurringCommitTemplateModelCommitDurationUnit `json:"unit,omitempty"`
	Value Number                                         `json:"value"`
}

// The amount the customer should be billed for the commit.
type RecurringCommitTemplateModelInvoiceAmount struct {
	CreditTypeID string `json:"credit_type_id"`
	Quantity     Number `json:"quantity"`
	UnitPrice    Number `json:"unit_price"`
}

// Rounding configuration for prorated recurring commit amounts.
type RecurringCommitTemplateModelProrationRounding struct {
	Access  *ProrationRoundingConfigModel `json:"access,omitempty"`
	Invoice *ProrationRoundingConfigModel `json:"invoice,omitempty"`
}
type RecurringCommitTemplateModel struct {
	APIResource
	// The amount of commit to grant.
	AccessAmount *RecurringCommitTemplateModelAccessAmount `json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	// The amount of time each of the created commits will be valid for
	CommitDuration *RecurringCommitTemplateModelCommitDuration `json:"commit_duration"`
	Description    string                                      `json:"description,omitempty"`
	// Offset relative to the recurring credit start that determines when the contract will stop creating recurring commits. optional
	Duration *RelativeDateModel `json:"duration,omitempty"`
	ID       string             `json:"id"`
	// The amount the customer should be billed for the commit.
	InvoiceAmount *RecurringCommitTemplateModelInvoiceAmount `json:"invoice_amount,omitempty"`
	Name          string                                     `json:"name,omitempty"`
	Priority      Number                                     `json:"priority"`
	Product       *RecurringCommitTemplateModelProduct       `json:"product"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration RecurringCommitTemplateModelProration `json:"proration,omitempty"`
	// Rounding configuration for prorated recurring commit amounts.
	ProrationRounding *RecurringCommitTemplateModelProrationRounding `json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType RecurringCommitTemplateModelRateType `json:"rate_type"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency RecurringCommitTemplateModelRecurrenceFrequency `json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction Number `json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	// Offset relative to the contract start date that determines the start time for the first commit
	StartingAtOffset *RelativeDateModel `json:"starting_at_offset"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *RecurringCommitSubscriptionConfigTemplateModel `json:"subscription_config,omitempty"`
}
