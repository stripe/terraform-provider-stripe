//
// File generated from our OpenAPI spec
//

package metronome

// Determines what spend contributes to calculating this charge.
type ProductListItemStateModelCompositeScope string

// List of values that ProductListItemStateModelCompositeScope can take
const (
	ProductListItemStateModelCompositeScopeCUSTOMER ProductListItemStateModelCompositeScope = "CUSTOMER"
	ProductListItemStateModelCompositeScopeCONTRACT ProductListItemStateModelCompositeScope = "CONTRACT"
)

type ProductListItemStateModel struct {
	APIResource
	BillableMetricID    string   `json:"billable_metric_id,omitempty"`
	CompositeProductIDs []string `json:"composite_product_ids,omitempty"`
	// Determines what spend contributes to calculating this charge.
	CompositeScope   ProductListItemStateModelCompositeScope `json:"composite_scope,omitempty"`
	CompositeTags    []string                                `json:"composite_tags,omitempty"`
	CreatedAt        string                                  `json:"created_at"`
	CreatedBy        string                                  `json:"created_by"`
	ExcludeFreeUsage bool                                    `json:"exclude_free_usage,omitempty"`
	// Only for composite products. If true, allows a composite to incorporate spend from other composite products. Defaults to false
	IncludeCompositeSpend bool `json:"include_composite_spend,omitempty"`
	// This field's availability is dependent on your client's configuration.
	IsRefundable bool   `json:"is_refundable,omitempty"`
	Name         string `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID string `json:"netsuite_internal_item_id,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID string                   `json:"netsuite_overage_item_id,omitempty"`
	PresentationGroupKey  []string                 `json:"presentation_group_key,omitempty"`
	PricingGroupKey       []string                 `json:"pricing_group_key,omitempty"`
	QuantityConversion    *QuantityConversionModel `json:"quantity_conversion,omitempty"`
	QuantityRounding      *QuantityRoundingModel   `json:"quantity_rounding,omitempty"`
	StartingAt            string                   `json:"starting_at,omitempty"`
	Tags                  []string                 `json:"tags,omitempty"`
}
