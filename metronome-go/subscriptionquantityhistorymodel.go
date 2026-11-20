//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionQuantityHistoryModelHistoryData struct {
	Quantity  Number `json:"quantity"`
	Total     Number `json:"total"`
	UnitPrice Number `json:"unit_price"`
}
type SubscriptionQuantityHistoryModelHistory struct {
	Data       []*SubscriptionQuantityHistoryModelHistoryData `json:"data"`
	StartingAt string                                         `json:"starting_at"`
}
type SubscriptionQuantityHistoryModel struct {
	APIResource
	FiatCreditTypeID string                                     `json:"fiat_credit_type_id,omitempty"`
	History          []*SubscriptionQuantityHistoryModelHistory `json:"history,omitempty"`
	SubscriptionID   string                                     `json:"subscription_id,omitempty"`
}
