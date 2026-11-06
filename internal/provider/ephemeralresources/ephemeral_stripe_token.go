//
// File generated from our OpenAPI spec
//

package ephemeralresources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	ephemeralSchema "github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ ephemeral.EphemeralResource = &TokenEphemeralResource{}
var _ ephemeral.EphemeralResourceWithConfigure = &TokenEphemeralResource{}

func NewTokenEphemeralResource() ephemeral.EphemeralResource {
	return &TokenEphemeralResource{}
}

type TokenEphemeralResource struct {
	client *stripe.Client
}

type TokenResourceModel struct {
	Object      types.String `tfsdk:"object"`
	BankAccount types.Object `tfsdk:"bank_account"`
	Card        types.Object `tfsdk:"card"`
	ClientIP    types.String `tfsdk:"client_ip"`
	Created     types.Int64  `tfsdk:"created"`
	ID          types.String `tfsdk:"id"`
	Livemode    types.Bool   `tfsdk:"livemode"`
	Type        types.String `tfsdk:"type"`
	Used        types.Bool   `tfsdk:"used"`
	Account     types.Object `tfsdk:"account"`
	Customer    types.String `tfsdk:"customer"`
	CVCUpdate   types.Object `tfsdk:"cvc_update"`
	Person      types.Object `tfsdk:"person"`
	PII         types.Object `tfsdk:"pii"`
}

func (r *TokenEphemeralResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_token"
}

func (r *TokenEphemeralResource) Schema(_ context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = ephemeralSchema.Schema{
		Description: "Tokenization is the process Stripe uses to collect sensitive card or bank\naccount details, or personally identifiable information (PII), directly from\nyour customers in a secure manner. A token representing this information is\nreturned to your server to use. Use our\n[recommended payments integrations](https://docs.stripe.com/payments) to perform this process\non the client-side. This guarantees that no sensitive card data touches your server,\nand allows your integration to operate in a PCI-compliant way.\n\nIf you can't use client-side tokenization, you can also create tokens using\nthe API with either your publishable or secret API key. If\nyour integration uses this method, you're responsible for any PCI compliance\nthat it might require, and you must keep your secret API key safe. Unlike with\nclient-side tokenization, your customer's information isn't sent directly to\nStripe, so we can't determine how it's handled or stored.\n\nYou can't store or use tokens more than once. To store card or bank account\ninformation for later use, create [Customer](https://docs.stripe.com/api#customers)\nobjects or [External accounts](/api#external_accounts).\n[Radar](https://docs.stripe.com/radar), our integrated solution for automatic fraud protection,\nperforms best with integrations that use client-side tokenization.",
		Attributes: map[string]ephemeralSchema.Attribute{
			"object": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "String representing the object's type. Objects of the same type share the same value.",
				Validators:  []validator.String{stringvalidator.OneOf("token")},
			},
			"bank_account": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "These bank accounts are payment methods on `Customer` objects.\n\nOn the other hand [External Accounts](/api#external_accounts) are transfer\ndestinations on `Account` objects for connected accounts.\nThey can be bank accounts or debit cards as well, and are documented in the links above.\n\nRelated guide: [Bank debits and transfers](/payments/bank-debits-transfers)",
				Attributes: map[string]ephemeralSchema.Attribute{
					"account_holder_name": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The name of the person or business that owns the bank account. This field is required when attaching the bank account to a `Customer` object.",
					},
					"account_holder_type": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The type of entity that holds the account. It can be `company` or `individual`. This field is required when attaching the bank account to a `Customer` object.",
					},
					"account_number": ephemeralSchema.StringAttribute{
						Required:    true,
						Description: "The account number for the bank account, in string form. Must be a checking account.",
					},
					"account_type": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The bank account type. This can only be `checking` or `savings` in most countries. In Japan, this can only be `futsu` or `toza`.",
					},
					"country": ephemeralSchema.StringAttribute{
						Required:    true,
						Description: "The country in which the bank account is located.",
					},
					"currency": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The currency the bank account is in. This must be a country/currency pairing that [Stripe supports.](https://docs.stripe.com/payouts)",
					},
					"payment_method": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The ID of a Payment Method with a `type` of `us_bank_account`. The Payment Method's bank account information will be copied and returned as a Bank Account Token. This parameter is exclusive with respect to all other parameters in the `bank_account` hash. You must include the top-level `customer` parameter if the Payment Method is attached to a `Customer` object. If the Payment Method is not attached to a `Customer` object, it will be consumed and cannot be used again. You may not use Payment Methods which were created by a Setup Intent with `attach_to_self=true`.",
					},
					"routing_number": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The routing number, sort code, or other country-appropriate institution number for the bank account. For US bank accounts, this is required and should be the ACH routing number, not the wire routing number. If you are providing an IBAN for `account_number`, this field is not required.",
					},
				},
			},
			"card": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "You can store multiple cards on a customer in order to charge the customer\nlater. You can also store multiple debit cards on a recipient in order to\ntransfer to those cards later.\n\nRelated guide: [Card payments with Sources](https://docs.stripe.com/sources/cards)",
				Attributes: map[string]ephemeralSchema.Attribute{
					"number": ephemeralSchema.StringAttribute{
						Required:    true,
						Description: "The card number, as a string without any separators.",
					},
					"exp_month": ephemeralSchema.Int64Attribute{
						Required:    true,
						Description: "Two-digit number representing the card's expiration month.",
					},
					"exp_year": ephemeralSchema.Int64Attribute{
						Required:    true,
						Description: "Two- or four-digit number representing the card's expiration year.",
					},
					"cvc": ephemeralSchema.StringAttribute{
						Required:    true,
						Description: "Card security code.",
					},
				},
			},
			"client_ip": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "IP address of the client that generates the token.",
			},
			"created": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was created. Measured in seconds since the Unix epoch.",
			},
			"id": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier for the object.",
			},
			"livemode": ephemeralSchema.BoolAttribute{
				Computed:    true,
				Description: "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
			},
			"type": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "Type of the token: `account`, `bank_account`, `card`, or `pii`.",
			},
			"used": ephemeralSchema.BoolAttribute{
				Computed:    true,
				Description: "Determines if you have already used this token (you can only use tokens once).",
			},
			"account": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "Information for the account this token represents.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"business_type": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The business type.",
					},
					"company": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Information about the company or business.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"address": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The company's primary address.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"city": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "City, district, suburb, town, or village.",
									},
									"country": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
									},
									"line1": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Address line 1, such as the street, PO Box, or company name.",
									},
									"line2": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Address line 2, such as the apartment, suite, unit, or building.",
									},
									"postal_code": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "ZIP or postal code.",
									},
									"state": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
									},
								},
							},
							"address_kana": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The Kana variation of the company's primary address (Japan only).",
								Attributes: map[string]ephemeralSchema.Attribute{
									"city": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "City or ward.",
									},
									"country": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
									},
									"line1": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Block or building number.",
									},
									"line2": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Building details.",
									},
									"postal_code": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Postal code.",
									},
									"state": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Prefecture.",
									},
									"town": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Town or cho-me.",
									},
								},
							},
							"address_kanji": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The Kanji variation of the company's primary address (Japan only).",
								Attributes: map[string]ephemeralSchema.Attribute{
									"city": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "City or ward.",
									},
									"country": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
									},
									"line1": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Block or building number.",
									},
									"line2": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Building details.",
									},
									"postal_code": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Postal code.",
									},
									"state": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Prefecture.",
									},
									"town": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Town or cho-me.",
									},
								},
							},
							"directors_provided": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Whether the company's directors have been provided. Set this Boolean to `true` after creating all the company's directors with [the Persons API](/api/persons) for accounts with a `relationship.director` requirement. This value is not automatically set to `true` after creating directors, so it needs to be updated to indicate all directors have been provided.",
							},
							"directorship_declaration": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "This hash is used to attest that the directors information provided to Stripe is both current and correct.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"date": ephemeralSchema.Int64Attribute{
										Optional:    true,
										Description: "The Unix timestamp marking when the directorship declaration attestation was made.",
									},
									"ip": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The IP address from which the directorship declaration attestation was made.",
									},
									"user_agent": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The user agent of the browser from which the directorship declaration attestation was made.",
									},
								},
							},
							"executives_provided": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Whether the company's executives have been provided. Set this Boolean to `true` after creating all the company's executives with [the Persons API](/api/persons) for accounts with a `relationship.executive` requirement.",
							},
							"export_license_id": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The export license ID number of the company, also referred as Import Export Code (India only).",
							},
							"export_purpose_code": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The purpose code to use for export transactions (India only).",
							},
							"name": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The company's legal name.",
							},
							"name_kana": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The Kana variation of the company's legal name (Japan only).",
							},
							"name_kanji": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The Kanji variation of the company's legal name (Japan only).",
							},
							"owners_provided": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Whether the company's owners have been provided. Set this Boolean to `true` after creating all the company's owners with [the Persons API](/api/persons) for accounts with a `relationship.owner` requirement.",
							},
							"ownership_declaration": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "This hash is used to attest that the beneficial owner information provided to Stripe is both current and correct.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"date": ephemeralSchema.Int64Attribute{
										Optional:    true,
										Description: "The Unix timestamp marking when the beneficial owner attestation was made.",
									},
									"ip": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The IP address from which the beneficial owner attestation was made.",
									},
									"user_agent": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The user agent of the browser from which the beneficial owner attestation was made.",
									},
								},
							},
							"ownership_declaration_shown_and_signed": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Whether the user described by the data in the token has been shown the Ownership Declaration and indicated that it is correct.",
							},
							"ownership_exemption_reason": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "This value is used to determine if a business is exempt from providing ultimate beneficial owners. See [this support article](https://support.stripe.com/questions/exemption-from-providing-ownership-details) and [changelog](https://docs.stripe.com/changelog/acacia/2025-01-27/ownership-exemption-reason-accounts-api) for more details.",
							},
							"phone": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The company's phone number (used for verification).",
							},
							"registration_date": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "When the business was incorporated or registered.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"day": ephemeralSchema.Int64Attribute{
										Required:    true,
										Description: "The day of registration, between 1 and 31.",
									},
									"month": ephemeralSchema.Int64Attribute{
										Required:    true,
										Description: "The month of registration, between 1 and 12.",
									},
									"year": ephemeralSchema.Int64Attribute{
										Required:    true,
										Description: "The four-digit year of registration.",
									},
								},
							},
							"registration_number": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The identification number given to a company when it is registered or incorporated, if distinct from the identification number used for filing taxes. (Examples are the CIN for companies and LLP IN for partnerships in India, and the Company Registration Number in Hong Kong).",
							},
							"representative_declaration": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "This hash is used to attest that the representative is authorized to act as the representative of their legal entity.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"date": ephemeralSchema.Int64Attribute{
										Optional:    true,
										Description: "The Unix timestamp marking when the representative declaration attestation was made.",
									},
									"ip": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The IP address from which the representative declaration attestation was made.",
									},
									"user_agent": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The user agent of the browser from which the representative declaration attestation was made.",
									},
								},
							},
							"structure": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The category identifying the legal structure of the company or legal entity. See [Business structure](/connect/identity-verification#business-structure) for more details. Pass an empty string to unset this value.",
							},
							"tax_id": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The business ID number of the company, as appropriate for the company’s country. (Examples are an Employer ID Number in the U.S., a Business Number in Canada, or a Company Number in the UK.)",
							},
							"tax_id_registrar": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The jurisdiction in which the `tax_id` is registered (Germany-based companies only).",
							},
							"vat_id": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The VAT number of the company.",
							},
							"verification": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "Information on the verification state of the company.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"document": ephemeralSchema.SingleNestedAttribute{
										Optional:    true,
										Description: "A document verifying the business.",
										Attributes: map[string]ephemeralSchema.Attribute{
											"back": ephemeralSchema.StringAttribute{
												Optional:    true,
												Description: "The back of a document returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `additional_verification`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.",
											},
											"front": ephemeralSchema.StringAttribute{
												Optional:    true,
												Description: "The front of a document returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `additional_verification`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.",
											},
										},
									},
								},
							},
						},
					},
					"individual": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Information about the person represented by the account.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"address": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The individual's primary address.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"city": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "City, district, suburb, town, or village.",
									},
									"country": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
									},
									"line1": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Address line 1, such as the street, PO Box, or company name.",
									},
									"line2": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Address line 2, such as the apartment, suite, unit, or building.",
									},
									"postal_code": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "ZIP or postal code.",
									},
									"state": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
									},
								},
							},
							"address_kana": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The Kana variation of the individual's primary address (Japan only).",
								Attributes: map[string]ephemeralSchema.Attribute{
									"city": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "City or ward.",
									},
									"country": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
									},
									"line1": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Block or building number.",
									},
									"line2": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Building details.",
									},
									"postal_code": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Postal code.",
									},
									"state": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Prefecture.",
									},
									"town": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Town or cho-me.",
									},
								},
							},
							"address_kanji": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The Kanji variation of the individual's primary address (Japan only).",
								Attributes: map[string]ephemeralSchema.Attribute{
									"city": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "City or ward.",
									},
									"country": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
									},
									"line1": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Block or building number.",
									},
									"line2": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Building details.",
									},
									"postal_code": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Postal code.",
									},
									"state": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Prefecture.",
									},
									"town": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Town or cho-me.",
									},
								},
							},
							"dob": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The individual's date of birth.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"day": ephemeralSchema.Int64Attribute{
										Required:    true,
										Description: "The day of birth, between 1 and 31.",
									},
									"month": ephemeralSchema.Int64Attribute{
										Required:    true,
										Description: "The month of birth, between 1 and 12.",
									},
									"year": ephemeralSchema.Int64Attribute{
										Required:    true,
										Description: "The four-digit year of birth.",
									},
								},
							},
							"email": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The individual's email address.",
							},
							"first_name": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The individual's first name.",
							},
							"first_name_kana": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The Kana variation of the individual's first name (Japan only).",
							},
							"first_name_kanji": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The Kanji variation of the individual's first name (Japan only).",
							},
							"full_name_aliases": ephemeralSchema.ListAttribute{
								Optional:    true,
								Description: "A list of alternate names or aliases that the individual is known by.",
								ElementType: types.StringType,
							},
							"gender": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The individual's gender",
							},
							"id_number": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The government-issued ID number of the individual, as appropriate for the representative's country. (Examples are a Social Security Number in the U.S., or a Social Insurance Number in Canada). Instead of the number itself, you can also provide a [PII token created with Stripe.js](/js/tokens/create_token?type=pii).",
							},
							"id_number_secondary": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The government-issued secondary ID number of the individual, as appropriate for the representative's country, will be used for enhanced verification checks. In Thailand, this would be the laser code found on the back of an ID card. Instead of the number itself, you can also provide a [PII token created with Stripe.js](/js/tokens/create_token?type=pii).",
							},
							"last_name": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The individual's last name.",
							},
							"last_name_kana": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The Kana variation of the individual's last name (Japan only).",
							},
							"last_name_kanji": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The Kanji variation of the individual's last name (Japan only).",
							},
							"maiden_name": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The individual's maiden name.",
							},
							"metadata": ephemeralSchema.MapAttribute{
								Optional:    true,
								Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
								ElementType: types.StringType,
							},
							"phone": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The individual's phone number.",
							},
							"political_exposure": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Indicates if the person or any of their representatives, family members, or other closely related persons, declares that they hold or have held an important public job or function, in any jurisdiction.",
							},
							"registered_address": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The individual's registered address.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"city": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "City, district, suburb, town, or village.",
									},
									"country": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
									},
									"line1": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Address line 1, such as the street, PO Box, or company name.",
									},
									"line2": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Address line 2, such as the apartment, suite, unit, or building.",
									},
									"postal_code": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "ZIP or postal code.",
									},
									"state": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
									},
								},
							},
							"relationship": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "Describes the person’s relationship to the account.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"director": ephemeralSchema.BoolAttribute{
										Optional:    true,
										Description: "Whether the person is a director of the account's legal entity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.",
									},
									"executive": ephemeralSchema.BoolAttribute{
										Optional:    true,
										Description: "Whether the person has significant responsibility to control, manage, or direct the organization.",
									},
									"owner": ephemeralSchema.BoolAttribute{
										Optional:    true,
										Description: "Whether the person is an owner of the account’s legal entity.",
									},
									"percent_ownership": ephemeralSchema.Float64Attribute{
										Optional:    true,
										Description: "The percent owned by the person of the account's legal entity.",
									},
									"title": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The person's title (e.g., CEO, Support Engineer).",
									},
								},
							},
							"ssn_last_4": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The last four digits of the individual's Social Security Number (U.S. only).",
							},
							"verification": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The individual's verification document information.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"additional_document": ephemeralSchema.SingleNestedAttribute{
										Optional:    true,
										Description: "A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.",
										Attributes: map[string]ephemeralSchema.Attribute{
											"back": ephemeralSchema.StringAttribute{
												Optional:    true,
												Description: "The back of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.",
											},
											"front": ephemeralSchema.StringAttribute{
												Optional:    true,
												Description: "The front of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.",
											},
										},
									},
									"document": ephemeralSchema.SingleNestedAttribute{
										Optional:    true,
										Description: "An identifying document, either a passport or local ID card.",
										Attributes: map[string]ephemeralSchema.Attribute{
											"back": ephemeralSchema.StringAttribute{
												Optional:    true,
												Description: "The back of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.",
											},
											"front": ephemeralSchema.StringAttribute{
												Optional:    true,
												Description: "The front of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.",
											},
										},
									},
								},
							},
						},
					},
					"tos_shown_and_accepted": ephemeralSchema.BoolAttribute{
						Optional:    true,
						Description: "Whether the user described by the data in the token has been shown [the Stripe Connected Account Agreement](/connect/account-tokens#stripe-connected-account-agreement). When creating an account token to create a new Connect account, this value must be `true`.",
					},
				},
			},
			"customer": ephemeralSchema.StringAttribute{
				Optional:    true,
				Description: "Create a token for the customer, which is owned by the application's account. You can only use this with an [OAuth access token](https://docs.stripe.com/connect/standard-accounts) or [Stripe-Account header](https://docs.stripe.com/connect/authentication). Learn more about [cloning saved payment methods](https://docs.stripe.com/connect/cloning-saved-payment-methods).",
			},
			"cvc_update": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "The updated CVC value this token represents.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"cvc": ephemeralSchema.StringAttribute{
						Required:    true,
						Description: "The CVC value, in string form.",
					},
				},
			},
			"person": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "Information for the person this token represents.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"additional_tos_acceptances": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Details on the legal guardian's or authorizer's acceptance of the required Stripe agreements.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"account": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "Details on the legal guardian's acceptance of the main Stripe service agreement.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"date": ephemeralSchema.Int64Attribute{
										Optional:    true,
										Description: "The Unix timestamp marking when the account representative accepted the service agreement.",
									},
									"ip": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The IP address from which the account representative accepted the service agreement.",
									},
									"user_agent": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The user agent of the browser from which the account representative accepted the service agreement.",
									},
								},
							},
						},
					},
					"address": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "The person's address.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"city": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "City, district, suburb, town, or village.",
							},
							"country": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
							},
							"line1": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Address line 1, such as the street, PO Box, or company name.",
							},
							"line2": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Address line 2, such as the apartment, suite, unit, or building.",
							},
							"postal_code": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "ZIP or postal code.",
							},
							"state": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
							},
						},
					},
					"address_kana": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "The Kana variation of the person's address (Japan only).",
						Attributes: map[string]ephemeralSchema.Attribute{
							"city": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "City or ward.",
							},
							"country": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
							},
							"line1": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Block or building number.",
							},
							"line2": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Building details.",
							},
							"postal_code": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Postal code.",
							},
							"state": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Prefecture.",
							},
							"town": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Town or cho-me.",
							},
						},
					},
					"address_kanji": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "The Kanji variation of the person's address (Japan only).",
						Attributes: map[string]ephemeralSchema.Attribute{
							"city": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "City or ward.",
							},
							"country": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
							},
							"line1": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Block or building number.",
							},
							"line2": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Building details.",
							},
							"postal_code": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Postal code.",
							},
							"state": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Prefecture.",
							},
							"town": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Town or cho-me.",
							},
						},
					},
					"dob": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "The person's date of birth.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"day": ephemeralSchema.Int64Attribute{
								Required:    true,
								Description: "The day of birth, between 1 and 31.",
							},
							"month": ephemeralSchema.Int64Attribute{
								Required:    true,
								Description: "The month of birth, between 1 and 12.",
							},
							"year": ephemeralSchema.Int64Attribute{
								Required:    true,
								Description: "The four-digit year of birth.",
							},
						},
					},
					"documents": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Documents that may be submitted to satisfy various informational requests.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"company_authorization": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "One or more documents that demonstrate proof that this person is authorized to represent the company.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"files": ephemeralSchema.ListAttribute{
										Optional:    true,
										Description: "One or more document ids returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `account_requirement`.",
										ElementType: types.StringType,
									},
								},
							},
							"passport": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "One or more documents showing the person's passport page with photo and personal data.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"files": ephemeralSchema.ListAttribute{
										Optional:    true,
										Description: "One or more document ids returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `account_requirement`.",
										ElementType: types.StringType,
									},
								},
							},
							"visa": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "One or more documents showing the person's visa required for living in the country where they are residing.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"files": ephemeralSchema.ListAttribute{
										Optional:    true,
										Description: "One or more document ids returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `account_requirement`.",
										ElementType: types.StringType,
									},
								},
							},
						},
					},
					"email": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The person's email address.",
					},
					"first_name": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The person's first name.",
					},
					"first_name_kana": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The Kana variation of the person's first name (Japan only).",
					},
					"first_name_kanji": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The Kanji variation of the person's first name (Japan only).",
					},
					"full_name_aliases": ephemeralSchema.ListAttribute{
						Optional:    true,
						Description: "A list of alternate names or aliases that the person is known by.",
						ElementType: types.StringType,
					},
					"gender": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The person's gender (International regulations require either \"male\" or \"female\").",
					},
					"id_number": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The person's ID number, as appropriate for their country. For example, a social security number in the U.S., social insurance number in Canada, etc. Instead of the number itself, you can also provide a [PII token provided by Stripe.js](https://docs.stripe.com/js/tokens/create_token?type=pii).",
					},
					"id_number_secondary": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The person's secondary ID number, as appropriate for their country, will be used for enhanced verification checks. In Thailand, this would be the laser code found on the back of an ID card. Instead of the number itself, you can also provide a [PII token provided by Stripe.js](https://docs.stripe.com/js/tokens/create_token?type=pii).",
					},
					"last_name": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The person's last name.",
					},
					"last_name_kana": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The Kana variation of the person's last name (Japan only).",
					},
					"last_name_kanji": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The Kanji variation of the person's last name (Japan only).",
					},
					"maiden_name": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The person's maiden name.",
					},
					"metadata": ephemeralSchema.MapAttribute{
						Optional:    true,
						Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.",
						ElementType: types.StringType,
					},
					"nationality": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The country where the person is a national. Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)), or \"XX\" if unavailable.",
					},
					"phone": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The person's phone number.",
					},
					"political_exposure": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Indicates if the person or any of their representatives, family members, or other closely related persons, declares that they hold or have held an important public job or function, in any jurisdiction.",
					},
					"registered_address": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "The person's registered address.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"city": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "City, district, suburb, town, or village.",
							},
							"country": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
							},
							"line1": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Address line 1, such as the street, PO Box, or company name.",
							},
							"line2": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "Address line 2, such as the apartment, suite, unit, or building.",
							},
							"postal_code": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "ZIP or postal code.",
							},
							"state": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
							},
						},
					},
					"relationship": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "The relationship that this person has with the account's legal entity.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"authorizer": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Whether the person is the authorizer of the account's representative.",
							},
							"director": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Whether the person is a director of the account's legal entity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.",
							},
							"executive": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Whether the person has significant responsibility to control, manage, or direct the organization.",
							},
							"legal_guardian": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Whether the person is the legal guardian of the account's representative.",
							},
							"owner": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Whether the person is an owner of the account’s legal entity.",
							},
							"percent_ownership": ephemeralSchema.Float64Attribute{
								Optional:    true,
								Description: "The percent owned by the person of the account's legal entity.",
							},
							"representative": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Whether the person is authorized as the primary representative of the account. This is the person nominated by the business to provide information about themselves, and general information about the account. There can only be one representative at any given time. At the time the account is created, this person should be set to the person responsible for opening the account.",
							},
							"title": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The person's title (e.g., CEO, Support Engineer).",
							},
						},
					},
					"ssn_last_4": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The last four digits of the person's Social Security number (U.S. only).",
					},
					"us_cfpb_data": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "Demographic data related to the person.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"ethnicity_details": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The persons ethnicity details",
								Attributes: map[string]ephemeralSchema.Attribute{
									"ethnicity": ephemeralSchema.ListAttribute{
										Optional:    true,
										Description: "The persons ethnicity",
										ElementType: types.StringType,
									},
									"ethnicity_other": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Please specify your origin, when other is selected.",
									},
								},
							},
							"race_details": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "The persons race details",
								Attributes: map[string]ephemeralSchema.Attribute{
									"race": ephemeralSchema.ListAttribute{
										Optional:    true,
										Description: "The persons race.",
										ElementType: types.StringType,
									},
									"race_other": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "Please specify your race, when other is selected.",
									},
								},
							},
							"self_identified_gender": ephemeralSchema.StringAttribute{
								Optional:    true,
								Description: "The persons self-identified gender",
							},
						},
					},
					"verification": ephemeralSchema.SingleNestedAttribute{
						Optional:    true,
						Description: "The person's verification status.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"additional_document": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"back": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The back of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.",
									},
									"front": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The front of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.",
									},
								},
							},
							"document": ephemeralSchema.SingleNestedAttribute{
								Optional:    true,
								Description: "An identifying document, either a passport or local ID card.",
								Attributes: map[string]ephemeralSchema.Attribute{
									"back": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The back of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.",
									},
									"front": ephemeralSchema.StringAttribute{
										Optional:    true,
										Description: "The front of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`. The uploaded file needs to be a color image (smaller than 8,000px by 8,000px), in JPG, PNG, or PDF format, and less than 10 MB in size.",
									},
								},
							},
						},
					},
				},
			},
			"pii": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "The PII this token represents.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"id_number": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "The `id_number` for the PII, in string form.",
					},
				},
			},
		},
	}
}

func (r *TokenEphemeralResource) Configure(_ context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stripe.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Ephemeral Resource Configure Type", fmt.Sprintf("Expected *stripe.Client, got: %T", req.ProviderData))
		return
	}

	r.client = client
}

func expandTokenCreate(plan TokenResourceModel) (*stripe.TokenCreateParams, error) {
	params := &stripe.TokenCreateParams{}

	if !plan.BankAccount.IsNull() && !plan.BankAccount.IsUnknown() {
		if !assignAttrValueToNamedField(params, "BankAccount", plan.BankAccount) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "bank_account", params)
		}
	}
	if !plan.Card.IsNull() && !plan.Card.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Card", plan.Card) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "card", params)
		}
	}
	if !plan.Account.IsNull() && !plan.Account.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Account", plan.Account) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "account", params)
		}
	}
	if !plan.Customer.IsNull() && !plan.Customer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Customer", "Customer", plan.Customer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "customer", params)
		}
	}
	if !plan.CVCUpdate.IsNull() && !plan.CVCUpdate.IsUnknown() {
		if !assignAttrValueToNamedField(params, "CVCUpdate", plan.CVCUpdate) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "cvc_update", params)
		}
	}
	if !plan.Person.IsNull() && !plan.Person.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Person", plan.Person) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "person", params)
		}
	}
	if !plan.PII.IsNull() && !plan.PII.IsUnknown() {
		if !assignAttrValueToNamedField(params, "PII", plan.PII) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "pii", params)
		}
	}

	return params, nil
}

func flattenToken(obj *stripe.Token, state *TokenResourceModel) error {
	raw, hasRaw := rawResponseToPlain(obj)
	{
		if rawValueObject, rawOk := plainValueAtPath(raw, "object"); rawOk {
			if valueObject, err := flattenPlainValue(rawValueObject, types.StringType, "object", "raw response"); err != nil {
				return err
			} else {
				if typedObject, ok := valueObject.(types.String); ok {
					state.Object = typedObject
				}
			}
		} else if !hasRaw {
			if responseValueObject, ok := plainFromResponseField(obj, "Object"); ok {
				if valueObject, err := flattenPlainValue(responseValueObject, types.StringType, "object", "response struct"); err != nil {
					return err
				} else {
					if typedObject, ok := valueObject.(types.String); ok {
						state.Object = typedObject
					}
				}
			}
		}
	}
	{
		if rawValueClientIP, rawOk := plainValueAtPath(raw, "client_ip"); rawOk {
			if valueClientIP, err := flattenPlainValue(rawValueClientIP, types.StringType, "client_ip", "raw response"); err != nil {
				return err
			} else {
				if typedClientIP, ok := valueClientIP.(types.String); ok {
					state.ClientIP = typedClientIP
				}
			}
		} else if !hasRaw {
			if responseValueClientIP, ok := plainFromResponseField(obj, "ClientIP"); ok {
				if valueClientIP, err := flattenPlainValue(responseValueClientIP, types.StringType, "client_ip", "response struct"); err != nil {
					return err
				} else {
					if typedClientIP, ok := valueClientIP.(types.String); ok {
						state.ClientIP = typedClientIP
					}
				}
			}
		}
	}
	{
		if rawValueCreated, rawOk := plainValueAtPath(raw, "created"); rawOk {
			if valueCreated, err := flattenPlainValue(rawValueCreated, types.Int64Type, "created", "raw response"); err != nil {
				return err
			} else {
				if typedCreated, ok := valueCreated.(types.Int64); ok {
					state.Created = typedCreated
				}
			}
		} else if !hasRaw {
			if responseValueCreated, ok := plainFromResponseField(obj, "Created"); ok {
				if valueCreated, err := flattenPlainValue(responseValueCreated, types.Int64Type, "created", "response struct"); err != nil {
					return err
				} else {
					if typedCreated, ok := valueCreated.(types.Int64); ok {
						state.Created = typedCreated
					}
				}
			}
		}
	}
	{
		if rawValueID, rawOk := plainValueAtPath(raw, "id"); rawOk {
			if valueID, err := flattenPlainValue(rawValueID, types.StringType, "id", "raw response"); err != nil {
				return err
			} else {
				if typedID, ok := valueID.(types.String); ok {
					state.ID = typedID
				}
			}
		} else if !hasRaw {
			if responseValueID, ok := plainFromResponseField(obj, "ID"); ok {
				if valueID, err := flattenPlainValue(responseValueID, types.StringType, "id", "response struct"); err != nil {
					return err
				} else {
					if typedID, ok := valueID.(types.String); ok {
						state.ID = typedID
					}
				}
			}
		}
	}
	{
		if rawValueLivemode, rawOk := plainValueAtPath(raw, "livemode"); rawOk {
			if valueLivemode, err := flattenPlainValue(rawValueLivemode, types.BoolType, "livemode", "raw response"); err != nil {
				return err
			} else {
				if typedLivemode, ok := valueLivemode.(types.Bool); ok {
					state.Livemode = typedLivemode
				}
			}
		} else if !hasRaw {
			if responseValueLivemode, ok := plainFromResponseField(obj, "Livemode"); ok {
				if valueLivemode, err := flattenPlainValue(responseValueLivemode, types.BoolType, "livemode", "response struct"); err != nil {
					return err
				} else {
					if typedLivemode, ok := valueLivemode.(types.Bool); ok {
						state.Livemode = typedLivemode
					}
				}
			}
		}
	}
	{
		if rawValueType, rawOk := plainValueAtPath(raw, "type"); rawOk {
			if valueType, err := flattenPlainValue(rawValueType, types.StringType, "type", "raw response"); err != nil {
				return err
			} else {
				if typedType, ok := valueType.(types.String); ok {
					state.Type = typedType
				}
			}
		} else if !hasRaw {
			if responseValueType, ok := plainFromResponseField(obj, "Type"); ok {
				if valueType, err := flattenPlainValue(responseValueType, types.StringType, "type", "response struct"); err != nil {
					return err
				} else {
					if typedType, ok := valueType.(types.String); ok {
						state.Type = typedType
					}
				}
			}
		}
	}
	{
		if rawValueUsed, rawOk := plainValueAtPath(raw, "used"); rawOk {
			if valueUsed, err := flattenPlainValue(rawValueUsed, types.BoolType, "used", "raw response"); err != nil {
				return err
			} else {
				if typedUsed, ok := valueUsed.(types.Bool); ok {
					state.Used = typedUsed
				}
			}
		} else if !hasRaw {
			if responseValueUsed, ok := plainFromResponseField(obj, "Used"); ok {
				if valueUsed, err := flattenPlainValue(responseValueUsed, types.BoolType, "used", "response struct"); err != nil {
					return err
				} else {
					if typedUsed, ok := valueUsed.(types.Bool); ok {
						state.Used = typedUsed
					}
				}
			}
		}
	}
	return nil
}

func (r *TokenEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var config TokenResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandTokenCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building Token ephemeral params", err.Error())
		return
	}

	obj, err := r.client.V1Tokens.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error opening Token ephemeral resource", err.Error())
		return
	}

	result := config
	if err := flattenToken(obj, &result); err != nil {
		resp.Diagnostics.AddError("Error flattening Token ephemeral response", err.Error())
		return
	}
	normalizeUnknownValues(&result)
	diags = resp.Result.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
}
