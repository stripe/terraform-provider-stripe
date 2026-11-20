//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1ApprovalRequestsService is used to invoke v1approvalrequests related APIs.
type v1ApprovalRequestsService struct {
	B   Backend
	Key string
}

// Approves a pending approval request, allowing the queued operation to proceed.
//
// Once approved, the original write operation will be executed automatically and the
// approval request status will transition to reflect the execution result.
func (c v1ApprovalRequestsService) ApproveApprovalRequest(ctx context.Context, id string, params *V1ApprovalRequestsApproveApprovalRequestParams) (*V1ApproveApprovalRequestResponse, error) {
	if params == nil {
		params = &V1ApprovalRequestsApproveApprovalRequestParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/approval_requests/%s/approve", id)
	v1approveapprovalrequestresponse := &V1ApproveApprovalRequestResponse{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, v1approveapprovalrequestresponse)
	return v1approveapprovalrequestresponse, err
}

// Denies a pending approval request, preventing the queued operation from executing.
//
// Once denied, the approval request is closed and the original write operation will
// not be performed.
func (c v1ApprovalRequestsService) DenyApprovalRequest(ctx context.Context, id string, params *V1ApprovalRequestsDenyApprovalRequestParams) (*V1DenyApprovalRequestResponse, error) {
	if params == nil {
		params = &V1ApprovalRequestsDenyApprovalRequestParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/approval_requests/%s/deny", id)
	v1denyapprovalrequestresponse := &V1DenyApprovalRequestResponse{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, v1denyapprovalrequestresponse)
	return v1denyapprovalrequestresponse, err
}

// Retrieves the details of a single approval request by its ID.
func (c v1ApprovalRequestsService) GetApprovalRequest(ctx context.Context, id string, params *V1ApprovalRequestsGetApprovalRequestParams) (*V1GetApprovalRequestResponse, error) {
	if params == nil {
		params = &V1ApprovalRequestsGetApprovalRequestParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/approval_requests/%s", id)
	v1getapprovalrequestresponse := &V1GetApprovalRequestResponse{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, v1getapprovalrequestresponse)
	return v1getapprovalrequestresponse, err
}

// Lists all approval requests submitted by agentic tokens for your account. Returns request details including status, token name, approval URL, and the original request payload. Use status and date filters to narrow results.
func (c v1ApprovalRequestsService) ListApprovalRequests(ctx context.Context, params *V1ApprovalRequestsListApprovalRequestsParams) (*V1ListApprovalRequestsResponse, error) {
	if params == nil {
		params = &V1ApprovalRequestsListApprovalRequestsParams{}
	}
	params.Context = ctx
	v1listapprovalrequestsresponse := &V1ListApprovalRequestsResponse{}
	err := c.B.Call(
		http.MethodGet, "/v1/approval_requests", c.Key, params, v1listapprovalrequestsresponse)
	return v1listapprovalrequestsresponse, err
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
func (c v1ApprovalRequestsService) SubmitApprovalRequests(ctx context.Context, params *V1ApprovalRequestsSubmitApprovalRequestsParams) (*V1SubmitApprovalRequestsResponse, error) {
	if params == nil {
		params = &V1ApprovalRequestsSubmitApprovalRequestsParams{}
	}
	params.Context = ctx
	v1submitapprovalrequestsresponse := &V1SubmitApprovalRequestsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/approval_requests", c.Key, params, v1submitapprovalrequestsresponse)
	return v1submitapprovalrequestsresponse, err
}
