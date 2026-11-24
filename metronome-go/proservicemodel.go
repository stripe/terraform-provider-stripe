//
// File generated from our OpenAPI spec
//

package metronome

type ProServiceModel struct {
	APIResource
	CustomFields map[string]string `json:"custom_fields,omitempty"`
	Description  string            `json:"description,omitempty"`
	ID           string            `json:"id"`
	// Maximum amount for the term.
	MaxAmount Number `json:"max_amount"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id,omitempty"`
	ProductID            string `json:"product_id"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount.
	Quantity Number `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified.
	UnitPrice Number `json:"unit_price"`
}
