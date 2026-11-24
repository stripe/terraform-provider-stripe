//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionModelCollectionSchedule string

// List of values that SubscriptionModelCollectionSchedule can take
const (
	SubscriptionModelCollectionScheduleADVANCE SubscriptionModelCollectionSchedule = "ADVANCE"
	SubscriptionModelCollectionScheduleARREARS SubscriptionModelCollectionSchedule = "ARREARS"
	SubscriptionModelCollectionScheduleAdvance SubscriptionModelCollectionSchedule = "advance"
	SubscriptionModelCollectionScheduleArrears SubscriptionModelCollectionSchedule = "arrears"
)

// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a **SEAT_BASED** subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
type SubscriptionModelQuantityManagementMode string

// List of values that SubscriptionModelQuantityManagementMode can take
const (
	SubscriptionModelQuantityManagementModeSEATBASED    SubscriptionModelQuantityManagementMode = "SEAT_BASED"
	SubscriptionModelQuantityManagementModeSeatBased    SubscriptionModelQuantityManagementMode = "seat_based"
	SubscriptionModelQuantityManagementModeQUANTITYONLY SubscriptionModelQuantityManagementMode = "QUANTITY_ONLY"
	SubscriptionModelQuantityManagementModeQuantityOnly SubscriptionModelQuantityManagementMode = "quantity_only"
)

type SubscriptionModelStatus string

// List of values that SubscriptionModelStatus can take
const (
	SubscriptionModelStatusNOTSTARTED SubscriptionModelStatus = "NOT_STARTED"
	SubscriptionModelStatusINCOMPLETE SubscriptionModelStatus = "INCOMPLETE"
	SubscriptionModelStatusACTIVE     SubscriptionModelStatus = "ACTIVE"
	SubscriptionModelStatusUNPAID     SubscriptionModelStatus = "UNPAID"
	SubscriptionModelStatusENDED      SubscriptionModelStatus = "ENDED"
)

type SubscriptionModel struct {
	APIResource
	BillingCycleConfig *SubscriptionBillingCycleConfigModel `json:"billing_cycle_config,omitempty"`
	BillingPeriods     *SubscriptionBillingPeriodsModel     `json:"billing_periods"`
	CollectionSchedule SubscriptionModelCollectionSchedule  `json:"collection_schedule"`
	CustomFields       map[string]string                    `json:"custom_fields,omitempty"`
	Description        string                               `json:"description,omitempty"`
	EndingBefore       string                               `json:"ending_before,omitempty"`
	FiatCreditTypeID   string                               `json:"fiat_credit_type_id,omitempty"`
	ID                 string                               `json:"id,omitempty"`
	Name               string                               `json:"name,omitempty"`
	PaidQuantity       Number                               `json:"paid_quantity,omitempty"`
	PaymentGateConfig  *SubscriptionPaymentGateConfigModel  `json:"payment_gate_config,omitempty"`
	// Custom fields from the subscription product referenced by `subscription_rate.product`. These are distinct from the subscription instance's `custom_fields`.
	ProductCustomFields map[string]string           `json:"product_custom_fields,omitempty"`
	Proration           *SubscriptionProrationModel `json:"proration"`
	// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a **SEAT_BASED** subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
	QuantityManagementMode SubscriptionModelQuantityManagementMode `json:"quantity_management_mode"`
	// List of quantity schedule items for the subscription. Only includes the current quantity and future quantity changes.
	QuantitySchedule []*SubscriptionQuantityScheduleModel `json:"quantity_schedule"`
	SeatConfig       *SubscriptionSeatConfigModel         `json:"seat_config,omitempty"`
	StartingAt       string                               `json:"starting_at"`
	Status           SubscriptionModelStatus              `json:"status,omitempty"`
	SubscriptionRate *SubscriptionRateModel               `json:"subscription_rate"`
}
