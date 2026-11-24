//
// File generated from our OpenAPI spec
//

package metronome

type V2ListSystemNotificationConfigsResponse struct {
	APIResource
	Cursor string                                         `json:"cursor,omitempty"`
	Data   []*LifecycleEventSystemNotificationConfigModel `json:"data"`
}
