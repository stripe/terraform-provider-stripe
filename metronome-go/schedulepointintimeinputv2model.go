//
// File generated from our OpenAPI spec
//

package metronome

type SchedulePointInTimeInputV2ModelRecurringScheduleFrequency string

// List of values that SchedulePointInTimeInputV2ModelRecurringScheduleFrequency can take
const (
	SchedulePointInTimeInputV2ModelRecurringScheduleFrequencyMONTHLY    SchedulePointInTimeInputV2ModelRecurringScheduleFrequency = "MONTHLY"
	SchedulePointInTimeInputV2ModelRecurringScheduleFrequencyMonthly    SchedulePointInTimeInputV2ModelRecurringScheduleFrequency = "monthly"
	SchedulePointInTimeInputV2ModelRecurringScheduleFrequencyQUARTERLY  SchedulePointInTimeInputV2ModelRecurringScheduleFrequency = "QUARTERLY"
	SchedulePointInTimeInputV2ModelRecurringScheduleFrequencyQuarterly  SchedulePointInTimeInputV2ModelRecurringScheduleFrequency = "quarterly"
	SchedulePointInTimeInputV2ModelRecurringScheduleFrequencySEMIANNUAL SchedulePointInTimeInputV2ModelRecurringScheduleFrequency = "SEMI_ANNUAL"
	SchedulePointInTimeInputV2ModelRecurringScheduleFrequencySemiAnnual SchedulePointInTimeInputV2ModelRecurringScheduleFrequency = "semi_annual"
	SchedulePointInTimeInputV2ModelRecurringScheduleFrequencyANNUAL     SchedulePointInTimeInputV2ModelRecurringScheduleFrequency = "ANNUAL"
	SchedulePointInTimeInputV2ModelRecurringScheduleFrequencyAnnual     SchedulePointInTimeInputV2ModelRecurringScheduleFrequency = "annual"
	SchedulePointInTimeInputV2ModelRecurringScheduleFrequencyWEEKLY     SchedulePointInTimeInputV2ModelRecurringScheduleFrequency = "WEEKLY"
	SchedulePointInTimeInputV2ModelRecurringScheduleFrequencyWeekly     SchedulePointInTimeInputV2ModelRecurringScheduleFrequency = "weekly"
)

type SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistribution string

// List of values that SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistribution can take
const (
	SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistributionDIVIDED        SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistribution = "DIVIDED"
	SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistributionDivided        SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistribution = "divided"
	SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistributionDIVIDEDROUNDED SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistributionDividedRounded SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistribution = "divided_rounded"
	SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistributionEACH           SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistribution = "EACH"
	SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistributionEach           SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistribution = "each"
)

// Either provide amount or provide both unit_price and quantity.
type SchedulePointInTimeInputV2ModelScheduleItem struct {
	// Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.
	Amount Number `json:"amount,omitempty"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.
	Quantity Number `json:"quantity,omitempty"`
	// timestamp of the scheduled event
	Timestamp string `json:"timestamp"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.
	UnitPrice Number `json:"unit_price,omitempty"`
}

// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
type SchedulePointInTimeInputV2ModelRecurringSchedule struct {
	// Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.
	Amount             Number                                                             `json:"amount,omitempty"`
	AmountDistribution SchedulePointInTimeInputV2ModelRecurringScheduleAmountDistribution `json:"amount_distribution"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore string                                                    `json:"ending_before"`
	Frequency    SchedulePointInTimeInputV2ModelRecurringScheduleFrequency `json:"frequency"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.
	Quantity Number `json:"quantity,omitempty"`
	// RFC 3339 timestamp (inclusive).
	StartingAt string `json:"starting_at"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.
	UnitPrice Number `json:"unit_price,omitempty"`
}

// Must provide either schedule_items or recurring_schedule.
type SchedulePointInTimeInputV2Model struct {
	APIResource
	// Defaults to USD (cents) if not passed.
	CreditTypeID string `json:"credit_type_id,omitempty"`
	// This field is only applicable to commit invoice schedules. If true, this schedule will not generate an invoice.
	DoNotInvoice bool `json:"do_not_invoice,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *SchedulePointInTimeInputV2ModelRecurringSchedule `json:"recurring_schedule,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []*SchedulePointInTimeInputV2ModelScheduleItem `json:"schedule_items,omitempty"`
}
