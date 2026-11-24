//
// File generated from our OpenAPI spec
//

package metronome

type V2ListOffsetNotificationConfigsResponse struct {
	APIResource
	Cursor string                                         `json:"cursor,omitempty"`
	Data   []*LifecycleEventOffsetNotificationConfigModel `json:"data"`
}
