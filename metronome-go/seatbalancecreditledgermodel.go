//
// File generated from our OpenAPI spec
//

package metronome

// Credit ledger type
type SeatBalanceCreditLedgerModelType string

// List of values that SeatBalanceCreditLedgerModelType can take
const (
	SeatBalanceCreditLedgerModelTypeCREDITSEGMENTSTART              SeatBalanceCreditLedgerModelType = "CREDIT_SEGMENT_START"
	SeatBalanceCreditLedgerModelTypeCREDITAUTOMATEDINVOICEDEDUCTION SeatBalanceCreditLedgerModelType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	SeatBalanceCreditLedgerModelTypeCREDITEXPIRATION                SeatBalanceCreditLedgerModelType = "CREDIT_EXPIRATION"
	SeatBalanceCreditLedgerModelTypeCREDITCANCELED                  SeatBalanceCreditLedgerModelType = "CREDIT_CANCELED"
	SeatBalanceCreditLedgerModelTypeCREDITCREDITED                  SeatBalanceCreditLedgerModelType = "CREDIT_CREDITED"
	SeatBalanceCreditLedgerModelTypeCREDITMANUAL                    SeatBalanceCreditLedgerModelType = "CREDIT_MANUAL"
	SeatBalanceCreditLedgerModelTypeCREDITSEATBASEDADJUSTMENT       SeatBalanceCreditLedgerModelType = "CREDIT_SEAT_BASED_ADJUSTMENT"
	SeatBalanceCreditLedgerModelTypeCREDITROLLOVER                  SeatBalanceCreditLedgerModelType = "CREDIT_ROLLOVER"
)

type SeatBalanceCreditLedgerModel struct {
	APIResource
	// Amount of the ledger entry
	Amount Number `json:"amount"`
	// The datetime when the ledger is created
	Timestamp string `json:"timestamp"`
	// Credit ledger type
	Type SeatBalanceCreditLedgerModelType `json:"type"`
}
