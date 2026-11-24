//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1SecurityService is used to invoke v1security related APIs.
type v1SecurityService struct {
	B   Backend
	Key string
}

// Get a comprehensive audit trail of all operations performed in your Metronome account, whether initiated through the API, web interface, or automated processes. This endpoint provides detailed logs of who did what and when, enabling compliance reporting, security monitoring, and operational troubleshooting across all interaction channels.
//
// ### Use this endpoint to:
// - Monitor all account activity for security and compliance purposes
// - Track configuration changes regardless of source (API, UI, or system)
// - Investigate issues by reviewing historical operations
//
// ### Key response fields:
// An array of AuditLog objects containing:
// - id: Unique identifier for the audit log entry
// - timestamp: When the action occurred (RFC 3339 format)
// - actor: Information about who performed the action
// - request: Details including request ID, IP address, and user agent
// - `resource_type`: The type of resource affected (e.g., customer, contract, invoice)
// - `resource_id`: The specific resource identifier
// - `action`: The operation performed
// - `next_page`: Cursor for continuous log retrieval
//
// ### Usage guidelines:
// - Continuous retrieval: The next_page token enables uninterrupted log streaming—save it between requests to ensure no logs are missed
// - Empty responses: An empty data array means no new logs yet; continue polling with the same next_page token
// - Date filtering:
//   - `starting_on`: Earliest logs to return (inclusive)
//   - `ending_before`: Latest logs to return (exclusive)
//   - Cannot be used with `next_page`
//
// - Resource filtering: Must specify both `resource_type` and `resource_id` together
// - Sort order: Default is `date_asc`; use `date_desc` for newest first
func (c v1SecurityService) GetAuditLogs(ctx context.Context, params *V1SecurityGetAuditLogsParams) (*V1GetAuditLogsResponse, error) {
	if params == nil {
		params = &V1SecurityGetAuditLogsParams{}
	}
	params.Context = ctx
	v1getauditlogsresponse := &V1GetAuditLogsResponse{}
	err := c.B.Call(
		http.MethodGet, "/v1/auditLogs", c.Key, params, v1getauditlogsresponse)
	return v1getauditlogsresponse, err
}

// Gets Metronome's service registry with associated IP addresses for security allowlisting and firewall configuration. Use this endpoint to maintain an up-to-date list of IPs that your systems should trust for Metronome communications. Returns service names and their current IP ranges, with new IPs typically appearing 30+ days before first use to ensure smooth allowlist updates.
func (c v1SecurityService) GetServices(ctx context.Context, params *V1SecurityGetServicesParams) (*V1GetServicesResponse, error) {
	if params == nil {
		params = &V1SecurityGetServicesParams{}
	}
	params.Context = ctx
	v1getservicesresponse := &V1GetServicesResponse{}
	err := c.B.Call(
		http.MethodGet, "/v1/services", c.Key, params, v1getservicesresponse)
	return v1getservicesresponse, err
}
