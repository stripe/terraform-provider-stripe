//
// File generated from our OpenAPI spec
//

package metronome

// The status of the threshold notification. If the notification is archived, null will be returned.
type CustomerAlertModelCustomerStatus string

// List of values that CustomerAlertModelCustomerStatus can take
const (
	CustomerAlertModelCustomerStatusOk         CustomerAlertModelCustomerStatus = "ok"
	CustomerAlertModelCustomerStatusInAlarm    CustomerAlertModelCustomerStatus = "in_alarm"
	CustomerAlertModelCustomerStatusEvaluating CustomerAlertModelCustomerStatus = "evaluating"
)

type CustomerAlertModel struct {
	APIResource
	Alert *AlertModel `json:"alert"`
	// The status of the threshold notification. If the notification is archived, null will be returned.
	CustomerStatus CustomerAlertModelCustomerStatus `json:"customer_status"`
	// If present, indicates the reason the threshold notification was triggered.
	TriggeredBy string `json:"triggered_by,omitempty"`
}
