//
// File generated from our OpenAPI spec
//

package metronome

type ProductListItemModelType string

// List of values that ProductListItemModelType can take
const (
	ProductListItemModelTypeUSAGE        ProductListItemModelType = "USAGE"
	ProductListItemModelTypeSUBSCRIPTION ProductListItemModelType = "SUBSCRIPTION"
	ProductListItemModelTypeCOMPOSITE    ProductListItemModelType = "COMPOSITE"
	ProductListItemModelTypeFIXED        ProductListItemModelType = "FIXED"
	ProductListItemModelTypePROSERVICE   ProductListItemModelType = "PRO_SERVICE"
)

type ProductListItemModel struct {
	APIResource
	ArchivedAt   string                        `json:"archived_at,omitempty"`
	Current      *ProductListItemStateModel    `json:"current"`
	CustomFields map[string]string             `json:"custom_fields,omitempty"`
	ID           string                        `json:"id"`
	Initial      *ProductListItemStateModel    `json:"initial"`
	Type         ProductListItemModelType      `json:"type"`
	Updates      []*ProductListItemUpdateModel `json:"updates"`
}
