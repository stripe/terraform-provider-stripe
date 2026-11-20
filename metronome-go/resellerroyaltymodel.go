//
// File generated from our OpenAPI spec
//

package metronome

type ResellerRoyaltyModelResellerType string

// List of values that ResellerRoyaltyModelResellerType can take
const (
	ResellerRoyaltyModelResellerTypeAWS           ResellerRoyaltyModelResellerType = "AWS"
	ResellerRoyaltyModelResellerTypeAWSPROSERVICE ResellerRoyaltyModelResellerType = "AWS_PRO_SERVICE"
	ResellerRoyaltyModelResellerTypeGCP           ResellerRoyaltyModelResellerType = "GCP"
	ResellerRoyaltyModelResellerTypeGCPPROSERVICE ResellerRoyaltyModelResellerType = "GCP_PRO_SERVICE"
)

type ResellerRoyaltyModel struct {
	APIResource
	ApplicableProductIDs  []string                         `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string                         `json:"applicable_product_tags,omitempty"`
	AwsAccountNumber      string                           `json:"aws_account_number,omitempty"`
	AwsOfferID            string                           `json:"aws_offer_id,omitempty"`
	AwsPayerReferenceID   string                           `json:"aws_payer_reference_id,omitempty"`
	EndingBefore          string                           `json:"ending_before,omitempty"`
	Fraction              Number                           `json:"fraction"`
	GcpAccountID          string                           `json:"gcp_account_id,omitempty"`
	GcpOfferID            string                           `json:"gcp_offer_id,omitempty"`
	NetsuiteResellerID    string                           `json:"netsuite_reseller_id"`
	ResellerContractValue Number                           `json:"reseller_contract_value,omitempty"`
	ResellerType          ResellerRoyaltyModelResellerType `json:"reseller_type"`
	StartingAt            string                           `json:"starting_at"`
}
