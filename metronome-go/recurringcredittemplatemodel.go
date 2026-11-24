//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type RecurringCreditTemplateModelAccessAmountAccessType string

// List of values that RecurringCreditTemplateModelAccessAmountAccessType can take
const (
	RecurringCreditTemplateModelAccessAmountAccessTypeSPEND    RecurringCreditTemplateModelAccessAmountAccessType = "SPEND"
	RecurringCreditTemplateModelAccessAmountAccessTypeQUANTITY RecurringCreditTemplateModelAccessAmountAccessType = "QUANTITY"
)

// Whether the created commits will use the commit rate or list rate
type RecurringCreditTemplateModelRateType string

// List of values that RecurringCreditTemplateModelRateType can take
const (
	RecurringCreditTemplateModelRateTypeCOMMITRATE RecurringCreditTemplateModelRateType = "COMMIT_RATE"
	RecurringCreditTemplateModelRateTypeLISTRATE   RecurringCreditTemplateModelRateType = "LIST_RATE"
)

type RecurringCreditTemplateModelCommitDurationUnit string

// List of values that RecurringCreditTemplateModelCommitDurationUnit can take
const (
	RecurringCreditTemplateModelCommitDurationUnitPERIODS RecurringCreditTemplateModelCommitDurationUnit = "PERIODS"
)

// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
type RecurringCreditTemplateModelRecurrenceFrequency string

// List of values that RecurringCreditTemplateModelRecurrenceFrequency can take
const (
	RecurringCreditTemplateModelRecurrenceFrequencyMONTHLY   RecurringCreditTemplateModelRecurrenceFrequency = "MONTHLY"
	RecurringCreditTemplateModelRecurrenceFrequencyQUARTERLY RecurringCreditTemplateModelRecurrenceFrequency = "QUARTERLY"
	RecurringCreditTemplateModelRecurrenceFrequencyANNUAL    RecurringCreditTemplateModelRecurrenceFrequency = "ANNUAL"
	RecurringCreditTemplateModelRecurrenceFrequencyWEEKLY    RecurringCreditTemplateModelRecurrenceFrequency = "WEEKLY"
	RecurringCreditTemplateModelRecurrenceFrequencyDAILY     RecurringCreditTemplateModelRecurrenceFrequency = "DAILY"
)

// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type RecurringCreditTemplateModelProration string

// List of values that RecurringCreditTemplateModelProration can take
const (
	RecurringCreditTemplateModelProrationNONE         RecurringCreditTemplateModelProration = "NONE"
	RecurringCreditTemplateModelProrationFIRST        RecurringCreditTemplateModelProration = "FIRST"
	RecurringCreditTemplateModelProrationLAST         RecurringCreditTemplateModelProration = "LAST"
	RecurringCreditTemplateModelProrationFIRSTANDLAST RecurringCreditTemplateModelProration = "FIRST_AND_LAST"
)

type RecurringCreditTemplateModelProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// The amount of commit to grant.
type RecurringCreditTemplateModelAccessAmount struct {
	// Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType RecurringCreditTemplateModelAccessAmountAccessType `json:"access_type,omitempty"`
	// This ID identifies the credit type for the access amount. Quantity-based recurring commits and credits return the null credit type UUID.
	CreditTypeID string `json:"credit_type_id"`
	Quantity     Number `json:"quantity,omitempty"`
	UnitPrice    Number `json:"unit_price"`
}

// The amount of time each of the created commits will be valid for
type RecurringCreditTemplateModelCommitDuration struct {
	Unit  RecurringCreditTemplateModelCommitDurationUnit `json:"unit,omitempty"`
	Value Number                                         `json:"value"`
}

// Rounding configuration for prorated recurring credit amounts.
type RecurringCreditTemplateModelProrationRounding struct {
	Access *ProrationRoundingConfigModel `json:"access,omitempty"`
}
type RecurringCreditTemplateModel struct {
	APIResource
	// The amount of commit to grant.
	AccessAmount *RecurringCreditTemplateModelAccessAmount `json:"access_amount"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids,omitempty"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	// The amount of time each of the created commits will be valid for
	CommitDuration *RecurringCreditTemplateModelCommitDuration `json:"commit_duration"`
	Description    string                                      `json:"description,omitempty"`
	// Offset relative to the recurring credit start that determines when the contract will stop creating recurring commits. optional
	Duration *RelativeDateModel                   `json:"duration,omitempty"`
	ID       string                               `json:"id"`
	Name     string                               `json:"name,omitempty"`
	Priority Number                               `json:"priority"`
	Product  *RecurringCreditTemplateModelProduct `json:"product"`
	// Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration RecurringCreditTemplateModelProration `json:"proration,omitempty"`
	// Rounding configuration for prorated recurring credit amounts.
	ProrationRounding *RecurringCreditTemplateModelProrationRounding `json:"proration_rounding,omitempty"`
	// Whether the created commits will use the commit rate or list rate
	RateType RecurringCreditTemplateModelRateType `json:"rate_type"`
	// The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions
	RecurrenceFrequency RecurringCreditTemplateModelRecurrenceFrequency `json:"recurrence_frequency,omitempty"`
	// Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.
	RolloverFraction Number `json:"rollover_fraction,omitempty"`
	// List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown.
	Specifiers []*CommitSpecifierModel `json:"specifiers,omitempty"`
	// Offset relative to the contract start date that determines the start time for the first commit
	StartingAtOffset *RelativeDateModel `json:"starting_at_offset"`
	// Attach a subscription to the recurring commit/credit.
	SubscriptionConfig *RecurringCommitSubscriptionConfigTemplateModel `json:"subscription_config,omitempty"`
}
