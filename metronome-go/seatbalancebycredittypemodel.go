//
// File generated from our OpenAPI spec
//

package metronome

// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
type SeatBalanceByCreditTypeModelAccessType string

// List of values that SeatBalanceByCreditTypeModelAccessType can take
const (
	SeatBalanceByCreditTypeModelAccessTypeSPEND    SeatBalanceByCreditTypeModelAccessType = "SPEND"
	SeatBalanceByCreditTypeModelAccessTypeQUANTITY SeatBalanceByCreditTypeModelAccessType = "QUANTITY"
)

type SeatBalanceByCreditTypeModel struct {
	APIResource
	// Indicates how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.
	AccessType SeatBalanceByCreditTypeModelAccessType `json:"access_type,omitempty"`
	// The total balance across all commits and credits for this seat, of this credit type.
	Balance Number `json:"balance"`
	// This ID identifies the credit type for the balance. Quantity-based balances return the null credit type UUID.
	CreditTypeID string `json:"credit_type_id"`
	// The total initial balances of all commits and credits for this seat, of this credit type.
	StartingBalance Number `json:"starting_balance"`
}
