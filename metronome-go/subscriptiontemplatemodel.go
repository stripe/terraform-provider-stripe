//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionTemplateModelCollectionSchedule string

// List of values that SubscriptionTemplateModelCollectionSchedule can take
const (
	SubscriptionTemplateModelCollectionScheduleADVANCE SubscriptionTemplateModelCollectionSchedule = "ADVANCE"
	SubscriptionTemplateModelCollectionScheduleARREARS SubscriptionTemplateModelCollectionSchedule = "ARREARS"
	SubscriptionTemplateModelCollectionScheduleAdvance SubscriptionTemplateModelCollectionSchedule = "advance"
	SubscriptionTemplateModelCollectionScheduleArrears SubscriptionTemplateModelCollectionSchedule = "arrears"
)

// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a SEAT_BASED subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
type SubscriptionTemplateModelQuantityManagementMode string

// List of values that SubscriptionTemplateModelQuantityManagementMode can take
const (
	SubscriptionTemplateModelQuantityManagementModeSEATBASED    SubscriptionTemplateModelQuantityManagementMode = "SEAT_BASED"
	SubscriptionTemplateModelQuantityManagementModeSeatBased    SubscriptionTemplateModelQuantityManagementMode = "seat_based"
	SubscriptionTemplateModelQuantityManagementModeQUANTITYONLY SubscriptionTemplateModelQuantityManagementMode = "QUANTITY_ONLY"
	SubscriptionTemplateModelQuantityManagementModeQuantityOnly SubscriptionTemplateModelQuantityManagementMode = "quantity_only"
)

type SubscriptionTemplateModel struct {
	APIResource
	BillingCycleConfig *SubscriptionBillingCycleConfigTemplateModel `json:"billing_cycle_config,omitempty"`
	CollectionSchedule SubscriptionTemplateModelCollectionSchedule  `json:"collection_schedule"`
	CustomFields       map[string]string                            `json:"custom_fields,omitempty"`
	Description        string                                       `json:"description,omitempty"`
	Duration           *RelativeDateModel                           `json:"duration,omitempty"`
	FiatCreditTypeID   string                                       `json:"fiat_credit_type_id,omitempty"`
	ID                 string                                       `json:"id,omitempty"`
	InitialQuantity    Number                                       `json:"initial_quantity,omitempty"`
	Name               string                                       `json:"name,omitempty"`
	PaymentGateConfig  *SubscriptionPaymentGateConfigModel          `json:"payment_gate_config,omitempty"`
	Proration          *SubscriptionProrationModel                  `json:"proration"`
	// Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a SEAT_BASED subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.
	QuantityManagementMode SubscriptionTemplateModelQuantityManagementMode `json:"quantity_management_mode,omitempty"`
	SeatConfig             *SubscriptionSeatConfigModel                    `json:"seat_config,omitempty"`
	StartingAtOffset       *RelativeDateModel                              `json:"starting_at_offset,omitempty"`
	// Human-readable, per-package-unique alias for this subscription template, if one was set at package creation. Can be used to reference the subscription in contract customizations.
	SubscriptionAlias string                 `json:"subscription_alias,omitempty"`
	SubscriptionRate  *SubscriptionRateModel `json:"subscription_rate"`
}
