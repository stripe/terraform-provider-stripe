//
// File generated from our OpenAPI spec
//

package metronome

// Per-commit details on why specific commits were not moved to historical
type MoveToHistoryResponseModelSkippedReason struct {
	CommitID string `json:"commit_id"`
	// Why the commit was skipped
	Reason string `json:"reason"`
}
type MoveToHistoryResponseModel struct {
	APIResource
	// IDs of the commits that were moved to historical
	HistoricalCommitIDs []string `json:"historical_commit_ids"`
	// Number of commits successfully moved to historical in this request
	HistoricalCount Number `json:"historical_count"`
	// Number of commits that were skipped due to ineligibility
	SkippedCount Number `json:"skipped_count"`
	// Per-commit details on why specific commits were not moved to historical
	SkippedReasons []*MoveToHistoryResponseModelSkippedReason `json:"skipped_reasons"`
}
