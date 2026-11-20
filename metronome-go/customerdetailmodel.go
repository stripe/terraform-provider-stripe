//
// File generated from our OpenAPI spec
//

package metronome

type CustomerDetailModelCurrentBillableStatusValue string

// List of values that CustomerDetailModelCurrentBillableStatusValue can take
const (
	CustomerDetailModelCurrentBillableStatusValueBillable   CustomerDetailModelCurrentBillableStatusValue = "billable"
	CustomerDetailModelCurrentBillableStatusValueUnbillable CustomerDetailModelCurrentBillableStatusValue = "unbillable"
)

// This field's availability is dependent on your client's configuration.
type CustomerDetailModelCurrentBillableStatus struct {
	EffectiveAt string                                        `json:"effective_at,omitempty"`
	Value       CustomerDetailModelCurrentBillableStatusValue `json:"value"`
}
type CustomerDetailModel struct {
	APIResource
	// RFC 3339 timestamp indicating when the customer was archived. Null if the customer is active.
	ArchivedAt string `json:"archived_at,omitempty"`
	// RFC 3339 timestamp indicating when the customer was created.
	CreatedAt string `json:"created_at"`
	// This field's availability is dependent on your client's configuration.
	CurrentBillableStatus *CustomerDetailModelCurrentBillableStatus `json:"current_billable_status,omitempty"`
	CustomerConfig        *CustomerConfigModel                      `json:"customer_config"`
	CustomFields          map[string]string                         `json:"custom_fields"`
	// (deprecated, use ingest_aliases instead) the first ID (Metronome or ingest alias) that can be used in usage events
	ExternalID string `json:"external_id"`
	// the Metronome ID of the customer
	ID string `json:"id"`
	// aliases for this customer that can be used instead of the Metronome customer ID in usage events
	IngestAliases []string `json:"ingest_aliases"`
	Name          string   `json:"name"`
	// RFC 3339 timestamp indicating when the customer was last updated.
	UpdatedAt string `json:"updated_at"`
}
