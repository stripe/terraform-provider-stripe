//
// File generated from our OpenAPI spec
//

package metronome

// Commit ledger type
type SeatBalanceCommitLedgerModelType string

// List of values that SeatBalanceCommitLedgerModelType can take
const (
	SeatBalanceCommitLedgerModelTypePREPAIDCOMMITSEGMENTSTART              SeatBalanceCommitLedgerModelType = "PREPAID_COMMIT_SEGMENT_START"
	SeatBalanceCommitLedgerModelTypePREPAIDCOMMITAUTOMATEDINVOICEDEDUCTION SeatBalanceCommitLedgerModelType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	SeatBalanceCommitLedgerModelTypePREPAIDCOMMITROLLOVER                  SeatBalanceCommitLedgerModelType = "PREPAID_COMMIT_ROLLOVER"
	SeatBalanceCommitLedgerModelTypePREPAIDCOMMITEXPIRATION                SeatBalanceCommitLedgerModelType = "PREPAID_COMMIT_EXPIRATION"
	SeatBalanceCommitLedgerModelTypePREPAIDCOMMITCANCELED                  SeatBalanceCommitLedgerModelType = "PREPAID_COMMIT_CANCELED"
	SeatBalanceCommitLedgerModelTypePREPAIDCOMMITCREDITED                  SeatBalanceCommitLedgerModelType = "PREPAID_COMMIT_CREDITED"
	SeatBalanceCommitLedgerModelTypePREPAIDCOMMITMANUAL                    SeatBalanceCommitLedgerModelType = "PREPAID_COMMIT_MANUAL"
	SeatBalanceCommitLedgerModelTypePREPAIDCOMMITSEATBASEDADJUSTMENT       SeatBalanceCommitLedgerModelType = "PREPAID_COMMIT_SEAT_BASED_ADJUSTMENT"
)

type SeatBalanceCommitLedgerModel struct {
	APIResource
	// Amount of the ledger entry
	Amount Number `json:"amount"`
	// The datetime when the ledger is created
	Timestamp string `json:"timestamp"`
	// Commit ledger type
	Type SeatBalanceCommitLedgerModelType `json:"type"`
}
