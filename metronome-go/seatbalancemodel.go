//
// File generated from our OpenAPI spec
//

package metronome

type SeatBalanceModel struct {
	APIResource
	Balances []*SeatBalanceByCreditTypeModel `json:"balances"`
	// Array of commits applicable to this seat with their balances
	Commits []*SeatCommitBalanceModel `json:"commits,omitempty"`
	// Array of credits applicable to this seat with their balances
	Credits []*SeatCreditBalanceModel `json:"credits,omitempty"`
	// The unique identifier for the seat
	SeatID string `json:"seat_id"`
}
