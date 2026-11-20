//
// File generated from our OpenAPI spec
//

package metronome

// Archive a package. Archived packages cannot be used to create new contracts. However, existing contracts associated with the package will continue to function as normal. Once you archive a package, you can still retrieve it in the UI and API, but you cannot unarchive it.
type V1PackagesArchivePackageParams struct {
	Params `form:"*"`
	// ID of the package to archive
	PackageID *string `form:"package_id" json:"package_id"`
}

// Gets the details for a specific package, including name, aliases, duration, and terms. Use this endpoint to understand a package's alias schedule, or display a specific package's details to end customers.
type V1PackagesGetPackageParams struct {
	Params    `form:"*"`
	PackageID *string `form:"package_id" json:"package_id"`
}
