//
// File generated from our OpenAPI spec
//

package metronome

type V1GetPurchasedSeatsResponseData struct {
	ChargeID  string `json:"charge_id"`
	SeatCount int64  `json:"seat_count"`
}
type V1GetPurchasedSeatsResponse struct {
	APIResource
	Data []*V1GetPurchasedSeatsResponseData `json:"data"`
}
