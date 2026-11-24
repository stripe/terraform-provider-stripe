//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type SeatCreditBalanceModelAccessType string

// List of values that SeatCreditBalanceModelAccessType can take
const (
	SeatCreditBalanceModelAccessTypeSPEND    SeatCreditBalanceModelAccessType = "SPEND"
	SeatCreditBalanceModelAccessTypeQUANTITY SeatCreditBalanceModelAccessType = "QUANTITY"
)

type SeatCreditBalanceModel struct {
	APIResource
	// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType SeatCreditBalanceModelAccessType `json:"access_type,omitempty"`
	// The current balance for this credit for this specific seat
	Balance Number `json:"balance"`
	// The credit type for this credit. Quantity-based credits return the null credit type UUID.
	CreditTypeID string `json:"credit_type_id,omitempty"`
	// The datetime when the credit expires
	EndDate string `json:"end_date,omitempty"`
	// The credit ID
	ID string `json:"id"`
	// Transaction history for this credit for this seat (only included if include_ledgers=true)
	LedgerEntries []*SeatBalanceCreditLedgerModel `json:"ledger_entries,omitempty"`
	// The datetime when the credit becomes active
	StartDate string `json:"start_date"`
}
