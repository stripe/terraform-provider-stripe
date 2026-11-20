//
// File generated from our OpenAPI spec
//

package metronome

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
type V1SecurityGetAuditLogsParams struct {
	Params `form:"*"`
	// RFC 3339 timestamp (exclusive). Cannot be used with 'next_page'.
	EndingBefore *string `form:"ending_before" json:"ending_before,omitempty"`
	// Max number of results that should be returned
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Optional parameter that can be used to filter which audit logs are returned. If you specify resource_id, you must also specify resource_type.
	ResourceID *string `form:"resource_id" json:"resource_id,omitempty"`
	// Optional parameter that can be used to filter which audit logs are returned. If you specify resource_type, you must also specify resource_id.
	ResourceType *string `form:"resource_type" json:"resource_type,omitempty"`
	// Sort order by timestamp, e.g. date_asc or date_desc. Defaults to date_asc.
	Sort *string `form:"sort" json:"sort,omitempty"`
	// RFC 3339 timestamp of the earliest audit log to return. Cannot be used with 'next_page'.
	StartingOn *string `form:"starting_on" json:"starting_on,omitempty"`
}

// Gets Metronome's service registry with associated IP addresses for security allowlisting and firewall configuration. Use this endpoint to maintain an up-to-date list of IPs that your systems should trust for Metronome communications. Returns service names and their current IP ranges, with new IPs typically appearing 30+ days before first use to ensure smooth allowlist updates.
type V1SecurityGetServicesParams struct {
	Params `form:"*"`
}
