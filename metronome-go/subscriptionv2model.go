//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionV2ModelCollectionSchedule string

// List of values that SubscriptionV2ModelCollectionSchedule can take
const (
	SubscriptionV2ModelCollectionScheduleADVANCE SubscriptionV2ModelCollectionSchedule = "ADVANCE"
	SubscriptionV2ModelCollectionScheduleARREARS SubscriptionV2ModelCollectionSchedule = "ARREARS"
	SubscriptionV2ModelCollectionScheduleAdvance SubscriptionV2ModelCollectionSchedule = "advance"
	SubscriptionV2ModelCollectionScheduleArrears SubscriptionV2ModelCollectionSchedule = "arrears"
)

// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a **SEAT_BASED** subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
type SubscriptionV2ModelQuantityManagementMode string

// List of values that SubscriptionV2ModelQuantityManagementMode can take
const (
	SubscriptionV2ModelQuantityManagementModeSEATBASED    SubscriptionV2ModelQuantityManagementMode = "SEAT_BASED"
	SubscriptionV2ModelQuantityManagementModeSeatBased    SubscriptionV2ModelQuantityManagementMode = "seat_based"
	SubscriptionV2ModelQuantityManagementModeQUANTITYONLY SubscriptionV2ModelQuantityManagementMode = "QUANTITY_ONLY"
	SubscriptionV2ModelQuantityManagementModeQuantityOnly SubscriptionV2ModelQuantityManagementMode = "quantity_only"
)

type SubscriptionV2ModelStatus string

// List of values that SubscriptionV2ModelStatus can take
const (
	SubscriptionV2ModelStatusNOTSTARTED SubscriptionV2ModelStatus = "NOT_STARTED"
	SubscriptionV2ModelStatusINCOMPLETE SubscriptionV2ModelStatus = "INCOMPLETE"
	SubscriptionV2ModelStatusACTIVE     SubscriptionV2ModelStatus = "ACTIVE"
	SubscriptionV2ModelStatusUNPAID     SubscriptionV2ModelStatus = "UNPAID"
	SubscriptionV2ModelStatusENDED      SubscriptionV2ModelStatus = "ENDED"
)

type SubscriptionV2Model struct {
	APIResource
	BillingCycleConfig *SubscriptionBillingCycleConfigModel  `json:"billing_cycle_config,omitempty"`
	BillingPeriods     *SubscriptionBillingPeriodsModel      `json:"billing_periods"`
	CollectionSchedule SubscriptionV2ModelCollectionSchedule `json:"collection_schedule"`
	CustomFields       map[string]string                     `json:"custom_fields,omitempty"`
	Description        string                                `json:"description,omitempty"`
	EndingBefore       string                                `json:"ending_before,omitempty"`
	FiatCreditTypeID   string                                `json:"fiat_credit_type_id,omitempty"`
	ID                 string                                `json:"id,omitempty"`
	Name               string                                `json:"name,omitempty"`
	PaidQuantity       Number                                `json:"paid_quantity,omitempty"`
	PaymentGateConfig  *SubscriptionPaymentGateConfigModel   `json:"payment_gate_config,omitempty"`
	// Custom fields from the subscription product referenced by `subscription_rate.product`. These are distinct from the subscription instance's `custom_fields`.
	ProductCustomFields map[string]string             `json:"product_custom_fields,omitempty"`
	Proration           *SubscriptionProrationV2Model `json:"proration"`
	// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a **SEAT_BASED** subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
	QuantityManagementMode SubscriptionV2ModelQuantityManagementMode `json:"quantity_management_mode"`
	// List of quantity schedule items for the subscription. Only includes the current quantity and future quantity changes.
	QuantitySchedule []*SubscriptionQuantityScheduleModel `json:"quantity_schedule"`
	SeatConfig       *SubscriptionSeatConfigModel         `json:"seat_config,omitempty"`
	StartingAt       string                               `json:"starting_at"`
	Status           SubscriptionV2ModelStatus            `json:"status,omitempty"`
	SubscriptionRate *SubscriptionRateModel               `json:"subscription_rate"`
}
