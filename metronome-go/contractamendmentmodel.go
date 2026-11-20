//
// File generated from our OpenAPI spec
//

package metronome

type ContractAmendmentModel struct {
	APIResource
	Commits   []*CommitModel `json:"commits"`
	CreatedAt string         `json:"created_at"`
	CreatedBy string         `json:"created_by"`
	Credits   []*CreditModel `json:"credits,omitempty"`
	// This field's availability is dependent on your client's configuration.
	Discounts []*DiscountModel `json:"discounts,omitempty"`
	ID        string           `json:"id"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string           `json:"netsuite_sales_order_id,omitempty"`
	Overrides            []*OverrideModel `json:"overrides"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []*ProServiceModel `json:"professional_services,omitempty"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []*ResellerRoyaltyOrUpdateModel `json:"reseller_royalties,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                  `json:"salesforce_opportunity_id,omitempty"`
	ScheduledCharges        []*ScheduledChargeModel `json:"scheduled_charges"`
	StartingAt              string                  `json:"starting_at"`
}
