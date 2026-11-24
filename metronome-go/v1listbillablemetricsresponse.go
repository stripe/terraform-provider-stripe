//
// File generated from our OpenAPI spec
//

package metronome

type V1ListBillableMetricsResponse struct {
	APIResource
	Data     []*BillableMetricWithDeprecatedFieldsModel `json:"data"`
	NextPage string                                     `json:"next_page"`
}
