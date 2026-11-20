//
// File generated from our OpenAPI spec
//

package metronome

type V1SetUpBillingProviderResponseData struct {
	DeliveryMethodID string `json:"delivery_method_id"`
}
type V1SetUpBillingProviderResponse struct {
	APIResource
	Data *V1SetUpBillingProviderResponseData `json:"data"`
}
