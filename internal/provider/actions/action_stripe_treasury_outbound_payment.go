//
// File generated from our OpenAPI spec
//

package actions

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/action"
	actionSchema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ action.Action = &TreasuryOutboundPaymentAction{}
var _ action.ActionWithConfigure = &TreasuryOutboundPaymentAction{}

func NewTreasuryOutboundPaymentAction() action.Action {
	return &TreasuryOutboundPaymentAction{}
}

type TreasuryOutboundPaymentAction struct {
	client *stripe.Client
}

type TreasuryOutboundPaymentResourceModel struct {
	Amount                          types.Int64  `tfsdk:"amount"`
	Currency                        types.String `tfsdk:"currency"`
	Customer                        types.String `tfsdk:"customer"`
	Description                     types.String `tfsdk:"description"`
	DestinationPaymentMethod        types.String `tfsdk:"destination_payment_method"`
	EndUserDetails                  types.Object `tfsdk:"end_user_details"`
	FinancialAccount                types.String `tfsdk:"financial_account"`
	Metadata                        types.Map    `tfsdk:"metadata"`
	StatementDescriptor             types.String `tfsdk:"statement_descriptor"`
	DestinationPaymentMethodData    types.Object `tfsdk:"destination_payment_method_data"`
	DestinationPaymentMethodOptions types.Object `tfsdk:"destination_payment_method_options"`
}

func (r *TreasuryOutboundPaymentAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_treasury_outbound_payment"
}

func (r *TreasuryOutboundPaymentAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = actionSchema.Schema{
		Description: "Use [OutboundPayments](https://docs.stripe.com/docs/treasury/moving-money/financial-accounts/out-of/outbound-payments) to send funds to another party's external bank account or [FinancialAccount](https://api.stripe.com#financial_accounts). To send money to an account belonging to the same user, use an [OutboundTransfer](https://api.stripe.com#outbound_transfers).\n\nSimulate OutboundPayment state changes with the `/v1/test_helpers/treasury/outbound_payments` endpoints. These methods can only be called on test mode objects.\n\nRelated guide: [Moving money with Treasury using OutboundPayment objects](https://docs.stripe.com/docs/treasury/moving-money/financial-accounts/out-of/outbound-payments)",
		Attributes: map[string]actionSchema.Attribute{
			"amount": actionSchema.Int64Attribute{
				Required:    true,
				Description: "Amount (in cents) transferred.",
			},
			"currency": actionSchema.StringAttribute{
				Required:    true,
				Description: "Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).",
			},
			"customer": actionSchema.StringAttribute{
				Optional:    true,
				Description: "ID of the [customer](https://docs.stripe.com/api/customers) to whom an OutboundPayment is sent.",
			},
			"description": actionSchema.StringAttribute{
				Optional:    true,
				Description: "An arbitrary string attached to the object. Often useful for displaying to users.",
			},
			"destination_payment_method": actionSchema.StringAttribute{
				Optional:    true,
				Description: "The PaymentMethod via which an OutboundPayment is sent. This field can be empty if the OutboundPayment was created using `destination_payment_method_data`.",
			},
			"end_user_details": actionSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "Details about the end user.",
				Attributes: map[string]actionSchema.Attribute{
					"ip_address": actionSchema.StringAttribute{
						Optional:    true,
						Description: "IP address of the user initiating the OutboundPayment. Set if `present` is set to `true`. IP address collection is required for risk and compliance reasons. This will be used to help determine if the OutboundPayment is authorized or should be blocked.",
					},
					"present": actionSchema.BoolAttribute{
						Required:    true,
						Description: "`true` if the OutboundPayment creation request is being made on behalf of an end user by a platform. Otherwise, `false`.",
					},
				},
			},
			"financial_account": actionSchema.StringAttribute{
				Required:    true,
				Description: "The FinancialAccount that funds were pulled from.",
			},
			"metadata": actionSchema.MapAttribute{
				Optional:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"statement_descriptor": actionSchema.StringAttribute{
				Optional:    true,
				Description: "The description that appears on the receiving end for an OutboundPayment (for example, bank statement for external bank transfer).",
			},
			"destination_payment_method_data": actionSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "Hash used to generate the PaymentMethod to be used for this OutboundPayment. Exclusive with `destination_payment_method`.",
				Attributes: map[string]actionSchema.Attribute{
					"billing_details": actionSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.",
						Attributes: map[string]actionSchema.Attribute{
							"address": actionSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "Billing address.",
								Attributes: map[string]actionSchema.Attribute{
									"city": actionSchema.StringAttribute{
										Optional:    true,
										Description: "City, district, suburb, town, or village.",
									},
									"country": actionSchema.StringAttribute{
										Optional:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
									},
									"line1": actionSchema.StringAttribute{
										Optional:    true,
										Description: "Address line 1, such as the street, PO Box, or company name.",
									},
									"line2": actionSchema.StringAttribute{
										Optional:    true,
										Description: "Address line 2, such as the apartment, suite, unit, or building.",
									},
									"postal_code": actionSchema.StringAttribute{
										Optional:    true,
										Description: "ZIP or postal code.",
									},
									"state": actionSchema.StringAttribute{
										Optional:    true,
										Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
									},
								},
							},
							"email": actionSchema.StringAttribute{
								Optional:    true,
								Description: "Email address.",
							},
							"name": actionSchema.StringAttribute{
								Optional:    true,
								Description: "Full name.",
							},
							"phone": actionSchema.StringAttribute{
								Optional:    true,
								Description: "Billing phone number (including extension).",
							},
						},
					},
					"financial_account": actionSchema.StringAttribute{
						Optional:    true,
						Description: "Required if type is set to `financial_account`. The FinancialAccount ID to send funds to.",
					},
					"metadata": actionSchema.MapAttribute{
						Optional:    true,
						Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
						ElementType: types.StringType,
					},
					"type": actionSchema.StringAttribute{
						Required:    true,
						Description: "The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.",
					},
					"us_bank_account": actionSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Required hash if type is set to `us_bank_account`.",
						Attributes: map[string]actionSchema.Attribute{
							"account_holder_type": actionSchema.StringAttribute{
								Optional:    true,
								Description: "Account holder type: individual or company.",
							},
							"account_number": actionSchema.StringAttribute{
								Optional:    true,
								Description: "Account number of the bank account.",
							},
							"account_type": actionSchema.StringAttribute{
								Optional:    true,
								Description: "Account type: checkings or savings. Defaults to checking if omitted.",
							},
							"financial_connections_account": actionSchema.StringAttribute{
								Optional:    true,
								Description: "The ID of a Financial Connections Account to use as a payment method.",
							},
							"routing_number": actionSchema.StringAttribute{
								Optional:    true,
								Description: "Routing number of the bank account.",
							},
						},
					},
				},
			},
			"destination_payment_method_options": actionSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "Payment method-specific configuration for this OutboundPayment.",
				Attributes: map[string]actionSchema.Attribute{
					"us_bank_account": actionSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Optional fields for `us_bank_account`.",
						Attributes: map[string]actionSchema.Attribute{
							"network": actionSchema.StringAttribute{
								Optional:    true,
								Description: "Specifies the network rails to be used. If not set, will default to the PaymentMethod's preferred network. See the [docs](https://docs.stripe.com/treasury/money-movement/timelines) to learn more about money movement timelines for each network type.",
							},
						},
					},
				},
			},
		},
	}
}

func (r *TreasuryOutboundPaymentAction) Configure(_ context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stripe.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Action Configure Type", fmt.Sprintf("Expected *stripe.Client, got: %T", req.ProviderData))
		return
	}

	r.client = client
}

func expandTreasuryOutboundPaymentCreate(plan TreasuryOutboundPaymentResourceModel) (*stripe.TreasuryOutboundPaymentCreateParams, error) {
	params := &stripe.TreasuryOutboundPaymentCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Currency.IsNull() && !plan.Currency.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Currency", "Currency", plan.Currency.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "currency", params)
		}
	}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Customer", "Customer", plan.Customer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer", params)
		}
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Description", "Description", plan.Description.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "description", params)
		}
	}
	if !plan.DestinationPaymentMethod.IsNull() && !plan.DestinationPaymentMethod.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "DestinationPaymentMethod", "DestinationPaymentMethod", plan.DestinationPaymentMethod.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "destination_payment_method", params)
		}
	}
	if !plan.EndUserDetails.IsNull() && !plan.EndUserDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "EndUserDetails", plan.EndUserDetails) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "end_user_details", params)
		}
	}
	if !plan.FinancialAccount.IsNull() && !plan.FinancialAccount.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FinancialAccount", "FinancialAccount", plan.FinancialAccount.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "financial_account", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.StatementDescriptor.IsNull() && !plan.StatementDescriptor.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "StatementDescriptor", "StatementDescriptor", plan.StatementDescriptor.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "statement_descriptor", params)
		}
	}
	if !plan.DestinationPaymentMethodData.IsNull() && !plan.DestinationPaymentMethodData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DestinationPaymentMethodData", plan.DestinationPaymentMethodData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "destination_payment_method_data", params)
		}
	}
	if !plan.DestinationPaymentMethodOptions.IsNull() && !plan.DestinationPaymentMethodOptions.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DestinationPaymentMethodOptions", plan.DestinationPaymentMethodOptions) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "destination_payment_method_options", params)
		}
	}

	return params, nil
}

func (r *TreasuryOutboundPaymentAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config TreasuryOutboundPaymentResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTreasuryOutboundPaymentCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building TreasuryOutboundPayment action params", err.Error())
		return
	}

	obj, err := r.client.V1TreasuryOutboundPayments.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error invoking TreasuryOutboundPayment action", err.Error())
		return
	}

	if resp.SendProgress != nil {
		resp.SendProgress(action.InvokeProgressEvent{Message: fmt.Sprintf("Created stripe_treasury_outbound_payment %s", obj.ID)})
	}
}
