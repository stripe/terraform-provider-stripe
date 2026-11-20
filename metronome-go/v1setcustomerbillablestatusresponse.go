//
// File generated from our OpenAPI spec
//

package metronome

type V1SetCustomerBillableStatusResponseDataCurrentBillableStatus string

// List of values that V1SetCustomerBillableStatusResponseDataCurrentBillableStatus can take
const (
	V1SetCustomerBillableStatusResponseDataCurrentBillableStatusBillable   V1SetCustomerBillableStatusResponseDataCurrentBillableStatus = "billable"
	V1SetCustomerBillableStatusResponseDataCurrentBillableStatusUnbillable V1SetCustomerBillableStatusResponseDataCurrentBillableStatus = "unbillable"
)

type V1SetCustomerBillableStatusResponseData struct {
	CurrentBillableStatus V1SetCustomerBillableStatusResponseDataCurrentBillableStatus `json:"current_billable_status"`
	ID                    string                                                       `json:"id"`
}
type V1SetCustomerBillableStatusResponse struct {
	APIResource
	Data *V1SetCustomerBillableStatusResponseData `json:"data"`
}
