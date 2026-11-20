//
// File generated from our OpenAPI spec
//

package metronome

// Approves a pending approval request, allowing the queued operation to proceed.
//
// Once approved, the original write operation will be executed automatically and the
// approval request status will transition to reflect the execution result.
type V1ApprovalRequestsApproveApprovalRequestParams struct {
	Params `form:"*"`
}

// Denies a pending approval request, preventing the queued operation from executing.
//
// Once denied, the approval request is closed and the original write operation will
// not be performed.
type V1ApprovalRequestsDenyApprovalRequestParams struct {
	Params `form:"*"`
}

// Retrieves the details of a single approval request by its ID.
type V1ApprovalRequestsGetApprovalRequestParams struct {
	Params `form:"*"`
}

// Lists all approval requests submitted by agentic tokens for your account. Returns request details including status, token name, approval URL, and the original request payload. Use status and date filters to narrow results.
type V1ApprovalRequestsListApprovalRequestsParams struct {
	Params `form:"*"`
	// Filter to requests created on or after this timestamp (inclusive).
	CreatedAfter *string `form:"created_after" json:"created_after,omitempty"`
	// Filter to requests created before this timestamp (exclusive).
	CreatedBefore *string `form:"created_before" json:"created_before,omitempty"`
	// Number of items to return per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Cursor for pagination. Use the value from the previous response's next_page field.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
	// Filter by approval request status.
	Status *string `form:"status" json:"status,omitempty"`
}

// Queues a write operation for human review when an agentic token is in use.
//
// When an agent attempts a write operation (e.g. `POST /v1/contracts/create`) using an
// agentic token, the server returns a `403 AgenticMutationBlocked` error with a pointer to
// this endpoint. The agent should then call this endpoint with the original request details
// so a human can review and approve the action.
//
// On success, a pending approval request is created and the response includes an
// `approval_url` linking to the Approval Dashboard where a human can review, approve, or
// reject the request. Once approved, the original operation is executed automatically and
// the dashboard reflects the completed status.
//
// Agents should surface the `approval_url` to the human user (e.g. via chat message or UI link)
// so they can navigate directly to the approval dashboard to review the request.
type V1ApprovalRequestsSubmitApprovalRequestsParams struct {
	Params `form:"*"`
	// The full request body of the operation being requested.
	Body        map[string]JSON `form:"body" json:"body"`
	Description *string         `form:"description" json:"description"`
	// The HTTP method of the operation being requested.
	Method *string `form:"method" json:"method"`
	Name   *string `form:"name" json:"name"`
	// The API path of the operation being requested (e.g. /v1/contracts/create).
	Path *string `form:"path" json:"path"`
}
