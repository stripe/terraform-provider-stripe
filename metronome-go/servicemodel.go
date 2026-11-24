//
// File generated from our OpenAPI spec
//

package metronome

type ServiceModelUsage string

// List of values that ServiceModelUsage can take
const (
	ServiceModelUsageMakesConnectionsFrom ServiceModelUsage = "makes_connections_from"
	ServiceModelUsageAcceptsConnectionsAt ServiceModelUsage = "accepts_connections_at"
)

type ServiceModel struct {
	APIResource
	Ips   []string          `json:"ips"`
	Name  string            `json:"name"`
	Usage ServiceModelUsage `json:"usage"`
}
