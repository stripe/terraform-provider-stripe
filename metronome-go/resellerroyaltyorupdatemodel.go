//
// File generated from our OpenAPI spec
//

package metronome

type ResellerRoyaltyOrUpdateModelResellerType string

// List of values that ResellerRoyaltyOrUpdateModelResellerType can take
const (
	ResellerRoyaltyOrUpdateModelResellerTypeAWS           ResellerRoyaltyOrUpdateModelResellerType = "AWS"
	ResellerRoyaltyOrUpdateModelResellerTypeAWSPROSERVICE ResellerRoyaltyOrUpdateModelResellerType = "AWS_PRO_SERVICE"
	ResellerRoyaltyOrUpdateModelResellerTypeGCP           ResellerRoyaltyOrUpdateModelResellerType = "GCP"
	ResellerRoyaltyOrUpdateModelResellerTypeGCPPROSERVICE ResellerRoyaltyOrUpdateModelResellerType = "GCP_PRO_SERVICE"
)

type ResellerRoyaltyOrUpdateModel struct {
	APIResource
	AwsAccountNumber      string                                   `json:"aws_account_number,omitempty"`
	AwsOfferID            string                                   `json:"aws_offer_id,omitempty"`
	AwsPayerReferenceID   string                                   `json:"aws_payer_reference_id,omitempty"`
	EndingBefore          string                                   `json:"ending_before,omitempty"`
	Fraction              Number                                   `json:"fraction,omitempty"`
	GcpAccountID          string                                   `json:"gcp_account_id,omitempty"`
	GcpOfferID            string                                   `json:"gcp_offer_id,omitempty"`
	NetsuiteResellerID    string                                   `json:"netsuite_reseller_id,omitempty"`
	ResellerContractValue Number                                   `json:"reseller_contract_value,omitempty"`
	ResellerType          ResellerRoyaltyOrUpdateModelResellerType `json:"reseller_type"`
	StartingAt            string                                   `json:"starting_at,omitempty"`
}
