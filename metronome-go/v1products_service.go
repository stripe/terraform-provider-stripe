//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1ProductsService is used to invoke v1products related APIs.
type v1ProductsService struct {
	B   Backend
	Key string
}

// Archive a product. Any current rate cards associated with this product will continue to function as normal. However, it will no longer be available as an option for newly created rates. Once you archive a product, you can still retrieve it in the UI and API, but you cannot unarchive it.
func (c v1ProductsService) ArchiveProductListItem(ctx context.Context, params *V1ProductsArchiveProductListItemParams) (*V1ArchiveProductListItemResponse, error) {
	if params == nil {
		params = &V1ProductsArchiveProductListItemParams{}
	}
	params.Context = ctx
	v1archiveproductlistitemresponse := &V1ArchiveProductListItemResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/products/archive", c.Key, params, v1archiveproductlistitemresponse)
	return v1archiveproductlistitemresponse, err
}

// Create a new product object. Products in Metronome represent your company's individual product or service offerings. A Product can be thought of as the basic unit of a line item on the invoice. This is analogous to SKUs or items in an ERP system. Give the product a meaningful name as they will appear on customer invoices.
func (c v1ProductsService) CreateProduct(ctx context.Context, params *V1ProductsCreateProductParams) (*V1CreateProductResponse, error) {
	if params == nil {
		params = &V1ProductsCreateProductParams{}
	}
	params.Context = ctx
	v1createproductresponse := &V1CreateProductResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/products/create", c.Key, params, v1createproductresponse)
	return v1createproductresponse, err
}

// Retrieve a product by its ID, including all metadata and historical changes.
func (c v1ProductsService) GetProduct(ctx context.Context, params *V1ProductsGetProductParams) (*V1GetProductResponse, error) {
	if params == nil {
		params = &V1ProductsGetProductParams{}
	}
	params.Context = ctx
	v1getproductresponse := &V1GetProductResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/products/get", c.Key, params, v1getproductresponse)
	return v1getproductresponse, err
}

// Get a paginated list of all products in your organization with their complete configuration, version history, and metadata. By default excludes archived products unless explicitly requested via the `archive_filter` parameter.
func (c v1ProductsService) ListProducts(ctx context.Context, params *V1ProductsListProductsParams) (*V1ListProductsResponse, error) {
	if params == nil {
		params = &V1ProductsListProductsParams{}
	}
	params.Context = ctx
	v1listproductsresponse := &V1ListProductsResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/products/list", c.Key, params, v1listproductsresponse)
	return v1listproductsresponse, err
}

// Updates a product's configuration while maintaining billing continuity for active customers. Use this endpoint to modify product names, metrics, pricing rules, and composite settings without disrupting ongoing billing cycles. Changes are scheduled using the starting_at timestamp, which must be on an hour boundary—set future dates to schedule updates ahead of time, or past dates for retroactive changes. Returns the updated product ID upon success.
//
// ### Usage guidance:
// - Product type cannot be changed after creation. For incorrect product types, create a new product and archive the original instead.
func (c v1ProductsService) UpdateProduct(ctx context.Context, params *V1ProductsUpdateProductParams) (*V1UpdateProductResponse, error) {
	if params == nil {
		params = &V1ProductsUpdateProductParams{}
	}
	params.Context = ctx
	v1updateproductresponse := &V1UpdateProductResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/contract-pricing/products/update", c.Key, params, v1updateproductresponse)
	return v1updateproductresponse, err
}
