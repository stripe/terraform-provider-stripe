//
// File generated from our OpenAPI spec
//

package metronome

type CommitSpecifierInputModel struct {
	APIResource
	// If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.
	Exclude []*ExcludeSpecifierModel `json:"exclude,omitempty"`
	// If provided, the specifier will apply to product usage with these set of presentation group values.
	PresentationGroupValues map[string]string `json:"presentation_group_values,omitempty"`
	// If provided, the specifier will apply to product usage with these set of pricing group values.
	PricingGroupValues map[string]string `json:"pricing_group_values,omitempty"`
	// If provided, the specifier will only apply to the product with the specified ID.
	ProductID string `json:"product_id,omitempty"`
	// If provided, the specifier will only apply to products with all the specified tags.
	ProductTags []string `json:"product_tags,omitempty"`
}
