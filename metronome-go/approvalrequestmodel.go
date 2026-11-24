//
// File generated from our OpenAPI spec
//

package metronome

type ApprovalRequestModelStatus string

// List of values that ApprovalRequestModelStatus can take
const (
	ApprovalRequestModelStatusPendingApproval    ApprovalRequestModelStatus = "pending_approval"
	ApprovalRequestModelStatusApproved           ApprovalRequestModelStatus = "approved"
	ApprovalRequestModelStatusDenied             ApprovalRequestModelStatus = "denied"
	ApprovalRequestModelStatusExecutionSucceeded ApprovalRequestModelStatus = "execution_succeeded"
	ApprovalRequestModelStatusExecutionError     ApprovalRequestModelStatus = "execution_error"
	ApprovalRequestModelStatusExpired            ApprovalRequestModelStatus = "expired"
)

type ApprovalRequestModelApprovalRequestPayload struct {
	Body        map[string]JSON `json:"body"`
	Description string          `json:"description"`
	Method      string          `json:"method"`
	Name        string          `json:"name"`
	Path        string          `json:"path"`
}
type ApprovalRequestModelResponse struct {
	ResponseBody   map[string]JSON `json:"response_body"`
	ResponseStatus Number          `json:"response_status"`
}
type ApprovalRequestModel struct {
	APIResource
	ApprovalRequestPayload *ApprovalRequestModelApprovalRequestPayload `json:"approval_request_payload"`
	ApprovalURL            string                                      `json:"approval_url"`
	CompletedAt            string                                      `json:"completed_at"`
	CreatedAt              string                                      `json:"created_at"`
	ExpiresAt              string                                      `json:"expires_at"`
	ID                     string                                      `json:"id"`
	Response               *ApprovalRequestModelResponse               `json:"response"`
	ReviewedAt             string                                      `json:"reviewed_at"`
	ReviewedBy             string                                      `json:"reviewed_by"`
	Status                 ApprovalRequestModelStatus                  `json:"status"`
	TokenName              string                                      `json:"token_name"`
}
