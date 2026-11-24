//
// File generated from our OpenAPI spec
//

package metronome

type LifecycleEventOffsetPolicyModel struct {
	APIResource
	// ISO-8601 duration string indicating how much time before or after the base event this notification should be sent. Positive values indicate notifications after the event, negative values indicate notifications before the event. Examples: "P1D" (1 day after), "-PT2H" (2 hours before)
	//
	Offset string `json:"offset"`
	// The type of lifecycle event that this offset is based on.
	//
	Type string `json:"type"`
}
