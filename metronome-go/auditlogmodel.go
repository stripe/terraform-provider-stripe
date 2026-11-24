//
// File generated from our OpenAPI spec
//

package metronome

type AuditLogModelStatus string

// List of values that AuditLogModelStatus can take
const (
	AuditLogModelStatusSuccess AuditLogModelStatus = "success"
	AuditLogModelStatusFailure AuditLogModelStatus = "failure"
	AuditLogModelStatusPending AuditLogModelStatus = "pending"
)

type AuditLogModelRequest struct {
	ID        string `json:"id"`
	IP        string `json:"ip,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
}
type AuditLogModel struct {
	APIResource
	Action       string                `json:"action,omitempty"`
	Actor        *ActorModel           `json:"actor,omitempty"`
	Description  string                `json:"description,omitempty"`
	ID           string                `json:"id"`
	Request      *AuditLogModelRequest `json:"request"`
	ResourceID   string                `json:"resource_id,omitempty"`
	ResourceType string                `json:"resource_type,omitempty"`
	Status       AuditLogModelStatus   `json:"status,omitempty"`
	Timestamp    string                `json:"timestamp"`
}
