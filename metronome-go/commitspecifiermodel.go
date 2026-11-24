//
// File generated from our OpenAPI spec
//

package metronome

type CommitSpecifierModel struct {
	APIResource
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude                 []*ExcludeSpecifierModel `json:"exclude,omitempty"`
	PresentationGroupValues map[string]string        `json:"presentation_group_values,omitempty"`
	PricingGroupValues      map[string]string        `json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []string `json:"product_tags,omitempty"`
}
