//
// File generated from our OpenAPI spec
//

package metronome

type V1GetAuditLogsResponse struct {
	APIResource
	Data []*AuditLogModel `json:"data"`
	// The next_page parameter is always returned to support ongoing log retrieval. It enables continuous querying, even when some requests return no new data. Save the next_page token from each response and use it for future requests to ensure no logs are missed. This setup is ideal for regular updates via automated processes, like cron jobs, to fetch logs continuously as they become available. When you receive an empty data array, it indicates a temporary absence of new logs, but subsequent requests might return new data.
	NextPage string `json:"next_page"`
}
