//
// File generated from our OpenAPI spec
//

package metronome

// Scopes threshold notification evaluation to a specific presentation group key on individual line items. Only present for spend notifications.
type GroupKeyFilterTypeModel struct {
	APIResource
	Key   string `json:"key"`
	Value string `json:"value"`
}
