//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"net/http"
)

// v1PackagesService is used to invoke v1packages related APIs.
type v1PackagesService struct {
	B   Backend
	Key string
}

// Archive a package. Archived packages cannot be used to create new contracts. However, existing contracts associated with the package will continue to function as normal. Once you archive a package, you can still retrieve it in the UI and API, but you cannot unarchive it.
func (c v1PackagesService) ArchivePackage(ctx context.Context, params *V1PackagesArchivePackageParams) (*V1ArchivePackageResponse, error) {
	if params == nil {
		params = &V1PackagesArchivePackageParams{}
	}
	params.Context = ctx
	v1archivepackageresponse := &V1ArchivePackageResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/packages/archive", c.Key, params, v1archivepackageresponse)
	return v1archivepackageresponse, err
}

// Gets the details for a specific package, including name, aliases, duration, and terms. Use this endpoint to understand a package's alias schedule, or display a specific package's details to end customers.
func (c v1PackagesService) GetPackage(ctx context.Context, params *V1PackagesGetPackageParams) (*V1GetPackageResponse, error) {
	if params == nil {
		params = &V1PackagesGetPackageParams{}
	}
	params.Context = ctx
	v1getpackageresponse := &V1GetPackageResponse{}
	err := c.B.Call(
		http.MethodPost, "/v1/packages/get", c.Key, params, v1getpackageresponse)
	return v1getpackageresponse, err
}
