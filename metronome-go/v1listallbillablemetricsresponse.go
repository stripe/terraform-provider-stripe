//
// File generated from our OpenAPI spec
//

package metronome

type V1ListAllBillableMetricsResponse struct {
	APIResource
	Data     []*BillableMetricV1Model `json:"data"`
	NextPage string                   `json:"next_page"`
}
