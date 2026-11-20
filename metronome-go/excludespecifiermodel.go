//
// File generated from our OpenAPI spec
//

package metronome

type ExcludeSpecifierModel struct {
	APIResource
	// If provided, the specifier will not apply to products with all the specified tags.
	ProductTags []string `json:"product_tags,omitempty"`
}
