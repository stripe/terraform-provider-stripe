//
// File generated from our OpenAPI spec
//

package metronome

type V1GetSubscriptionStatusResponseDataStatus string

// List of values that V1GetSubscriptionStatusResponseDataStatus can take
const (
	V1GetSubscriptionStatusResponseDataStatusNOTSTARTED V1GetSubscriptionStatusResponseDataStatus = "NOT_STARTED"
	V1GetSubscriptionStatusResponseDataStatusINCOMPLETE V1GetSubscriptionStatusResponseDataStatus = "INCOMPLETE"
	V1GetSubscriptionStatusResponseDataStatusACTIVE     V1GetSubscriptionStatusResponseDataStatus = "ACTIVE"
	V1GetSubscriptionStatusResponseDataStatusUNPAID     V1GetSubscriptionStatusResponseDataStatus = "UNPAID"
	V1GetSubscriptionStatusResponseDataStatusENDED      V1GetSubscriptionStatusResponseDataStatus = "ENDED"
)

type V1GetSubscriptionStatusResponseDataOpenInvoiceReason string

// List of values that V1GetSubscriptionStatusResponseDataOpenInvoiceReason can take
const (
	V1GetSubscriptionStatusResponseDataOpenInvoiceReasonSubscriptionCycle          V1GetSubscriptionStatusResponseDataOpenInvoiceReason = "subscription_cycle"
	V1GetSubscriptionStatusResponseDataOpenInvoiceReasonSubscriptionQuantityUpdate V1GetSubscriptionStatusResponseDataOpenInvoiceReason = "subscription_quantity_update"
)

type V1GetSubscriptionStatusResponseDataLatestCycleInvoice struct {
	EndTimestamp    string                                        `json:"end_timestamp"`
	ExternalInvoice *PaymentGatedSubscriptionExternalInvoiceModel `json:"external_invoice"`
	ID              string                                        `json:"id"`
	Quantity        Number                                        `json:"quantity"`
	StartTimestamp  string                                        `json:"start_timestamp"`
}
type V1GetSubscriptionStatusResponseDataOpenInvoice struct {
	EndTimestamp    string                                               `json:"end_timestamp"`
	ExternalInvoice *PaymentGatedSubscriptionExternalInvoiceModel        `json:"external_invoice"`
	ID              string                                               `json:"id"`
	Quantity        Number                                               `json:"quantity"`
	Reason          V1GetSubscriptionStatusResponseDataOpenInvoiceReason `json:"reason"`
	StartTimestamp  string                                               `json:"start_timestamp"`
}
type V1GetSubscriptionStatusResponseData struct {
	ID                 string                                                 `json:"id"`
	LatestCycleInvoice *V1GetSubscriptionStatusResponseDataLatestCycleInvoice `json:"latest_cycle_invoice"`
	OpenInvoices       []*V1GetSubscriptionStatusResponseDataOpenInvoice      `json:"open_invoices"`
	PaidQuantity       Number                                                 `json:"paid_quantity"`
	Status             V1GetSubscriptionStatusResponseDataStatus              `json:"status"`
}
type V1GetSubscriptionStatusResponse struct {
	APIResource
	Data *V1GetSubscriptionStatusResponseData `json:"data"`
}
