//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type SeatCommitBalanceModelAccessType string

// List of values that SeatCommitBalanceModelAccessType can take
const (
	SeatCommitBalanceModelAccessTypeSPEND    SeatCommitBalanceModelAccessType = "SPEND"
	SeatCommitBalanceModelAccessTypeQUANTITY SeatCommitBalanceModelAccessType = "QUANTITY"
)

type SeatCommitBalanceModel struct {
	APIResource
	// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType SeatCommitBalanceModelAccessType `json:"access_type,omitempty"`
	// The current balance for this commit for this specific seat
	Balance Number `json:"balance"`
	// The credit type for this commit. Quantity-based commits return the null credit type UUID.
	CreditTypeID string `json:"credit_type_id,omitempty"`
	// The datetime when the commit expires
	EndDate string `json:"end_date,omitempty"`
	// The commit or credit ID
	ID string `json:"id"`
	// Transaction history for this commit for this seat (only included if include_ledgers=true)
	LedgerEntries []*SeatBalanceCommitLedgerModel `json:"ledger_entries,omitempty"`
	// The datetime when the commit becomes active
	StartDate string `json:"start_date"`
}
