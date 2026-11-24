//
// File generated from our OpenAPI spec
//

package metronome

// Type of the threshold notification
type AlertModelType string

// List of values that AlertModelType can take
const (
	AlertModelTypeSpendThresholdReached                                AlertModelType = "spend_threshold_reached"
	AlertModelTypeMonthlyInvoiceTotalSpendThresholdReached             AlertModelType = "monthly_invoice_total_spend_threshold_reached"
	AlertModelTypeLowRemainingDaysForCommitSegmentReached              AlertModelType = "low_remaining_days_for_commit_segment_reached"
	AlertModelTypeLowRemainingCommitBalanceReached                     AlertModelType = "low_remaining_commit_balance_reached"
	AlertModelTypeLowRemainingCommitPercentageReached                  AlertModelType = "low_remaining_commit_percentage_reached"
	AlertModelTypeLowRemainingDaysForContractCreditSegmentReached      AlertModelType = "low_remaining_days_for_contract_credit_segment_reached"
	AlertModelTypeLowRemainingContractCreditBalanceReached             AlertModelType = "low_remaining_contract_credit_balance_reached"
	AlertModelTypeLowRemainingContractCreditPercentageReached          AlertModelType = "low_remaining_contract_credit_percentage_reached"
	AlertModelTypeLowRemainingContractCreditAndCommitBalanceReached    AlertModelType = "low_remaining_contract_credit_and_commit_balance_reached"
	AlertModelTypeLowRemainingContractCreditAndCommitPercentageReached AlertModelType = "low_remaining_contract_credit_and_commit_percentage_reached"
	AlertModelTypeLowRemainingSeatBalanceReached                       AlertModelType = "low_remaining_seat_balance_reached"
	AlertModelTypeInvoiceTotalReached                                  AlertModelType = "invoice_total_reached"
)

// Status of the threshold notification
type AlertModelStatus string

// List of values that AlertModelStatus can take
const (
	AlertModelStatusEnabled  AlertModelStatus = "enabled"
	AlertModelStatusArchived AlertModelStatus = "archived"
	AlertModelStatusDisabled AlertModelStatus = "disabled"
)

// Indicates the commit access type this notification is scoped to. Defaults to `SPEND` if not otherwise specified. Only present for `low_remaining_commit_balance_reached`, `low_remaining_commit_percentage_reached`, `low_remaining_contract_credit_and_commit_balance_reached`, `low_remaining_contract_credit_and_commit_percentage_reached`, `low_remaining_contract_credit_balance_reached`, `low_remaining_contract_credit_percentage_reached`, and `low_remaining_seat_balance_reached` notifications.
type AlertModelAccessType string

// List of values that AlertModelAccessType can take
const (
	AlertModelAccessTypeSPEND    AlertModelAccessType = "SPEND"
	AlertModelAccessTypeQUANTITY AlertModelAccessType = "QUANTITY"
)

// Only present for low_remaining_seat_balance_reached notifications. The seat group key or seat group key-value pair the alert is scoped to.
type AlertModelSeatFilter struct {
	// The seat group key (e.g., "seat_id", "user_id") that the alert is scoped to.
	SeatGroupKey string `json:"seat_group_key"`
	// The seat group value that the alert is scoped to.
	SeatGroupValue string `json:"seat_group_value,omitempty"`
}
type AlertModel struct {
	APIResource
	// Indicates the commit access type this notification is scoped to. Defaults to `SPEND` if not otherwise specified. Only present for `low_remaining_commit_balance_reached`, `low_remaining_commit_percentage_reached`, `low_remaining_contract_credit_and_commit_balance_reached`, `low_remaining_contract_credit_and_commit_percentage_reached`, `low_remaining_contract_credit_balance_reached`, `low_remaining_contract_credit_percentage_reached`, and `low_remaining_seat_balance_reached` notifications.
	AccessType AlertModelAccessType `json:"access_type,omitempty"`
	// Present for `low_remaining_contract_credit_and_commit_balance_reached` and `low_remaining_contract_credit_and_commit_percentage_reached` notifications. The filters that define the commits and credits used to calculate the remaining balance or percentage.
	AlertSpecifiers []*AlertSpecifierModel `json:"alert_specifiers,omitempty"`
	// An array of strings, representing a way to filter the credit grant this threshold notification applies to, by looking at the credit_grant_type field on the credit grant. This field is only defined for CreditPercentage and CreditBalance notifications
	CreditGrantTypeFilters []string         `json:"credit_grant_type_filters,omitempty"`
	CreditType             *CreditTypeModel `json:"credit_type,omitempty"`
	// A list of custom field filters for notification types that support advanced filtering
	CustomFieldFilters []*CustomFieldFilterTypeModel `json:"custom_field_filters,omitempty"`
	GroupKeyFilter     *GroupKeyFilterTypeModel      `json:"group_key_filter,omitempty"`
	// Only present for `spend_threshold_reached` notifications. Scope notification to a specific group key on individual line items.
	GroupValues []*GroupValueFilterTypeModel `json:"group_values,omitempty"`
	// the Metronome ID of the threshold notification
	ID                 string   `json:"id"`
	InvoiceTypesFilter []string `json:"invoice_types_filter,omitempty"`
	// Name of the threshold notification
	Name string `json:"name"`
	// Only present for low_remaining_seat_balance_reached notifications. The seat group key or seat group key-value pair the alert is scoped to.
	SeatFilter *AlertModelSeatFilter `json:"seat_filter,omitempty"`
	// Status of the threshold notification
	Status AlertModelStatus `json:"status"`
	// Threshold value of the notification policy
	Threshold Number `json:"threshold"`
	// Type of the threshold notification
	Type          AlertModelType `json:"type"`
	UniquenessKey string         `json:"uniqueness_key,omitempty"`
	// Timestamp for when the threshold notification's customer status was last updated
	UpdatedAt string `json:"updated_at"`
}
