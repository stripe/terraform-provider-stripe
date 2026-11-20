//
// File generated from our OpenAPI spec
//

package metronome

type SubscriptionProrationV2ModelInvoiceBehavior string

// List of values that SubscriptionProrationV2ModelInvoiceBehavior can take
const (
	SubscriptionProrationV2ModelInvoiceBehaviorBILLIMMEDIATELY          SubscriptionProrationV2ModelInvoiceBehavior = "BILL_IMMEDIATELY"
	SubscriptionProrationV2ModelInvoiceBehaviorBILLONNEXTCOLLECTIONDATE SubscriptionProrationV2ModelInvoiceBehavior = "BILL_ON_NEXT_COLLECTION_DATE"
	SubscriptionProrationV2ModelInvoiceBehaviorBillImmediately          SubscriptionProrationV2ModelInvoiceBehavior = "bill_immediately"
	SubscriptionProrationV2ModelInvoiceBehaviorBillOnNextCollectionDate SubscriptionProrationV2ModelInvoiceBehavior = "bill_on_next_collection_date"
)

type SubscriptionProrationV2Model struct {
	APIResource
	InvoiceBehavior SubscriptionProrationV2ModelInvoiceBehavior `json:"invoice_behavior"`
	IsProrated      bool                                        `json:"is_prorated"`
	Rounding        *ProrationRoundingConfigV2Model             `json:"rounding,omitempty"`
}
