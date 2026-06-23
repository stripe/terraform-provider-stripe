//
// File generated from our OpenAPI spec
//

package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &PersonResource{}

var _ resource.ResourceWithConfigure = &PersonResource{}

var _ resource.ResourceWithImportState = &PersonResource{}

func NewPersonResource() resource.Resource {
	return &PersonResource{}
}

type PersonResource struct {
	client *stripe.Client
}

type PersonResourceModel struct {
	Object                    types.String `tfsdk:"object"`
	Account                   types.String `tfsdk:"account"`
	AdditionalTOSAcceptances  types.Object `tfsdk:"additional_tos_acceptances"`
	Address                   types.Object `tfsdk:"address"`
	AddressKana               types.Object `tfsdk:"address_kana"`
	AddressKanji              types.Object `tfsdk:"address_kanji"`
	Created                   types.Int64  `tfsdk:"created"`
	DOB                       types.Object `tfsdk:"dob"`
	Email                     types.String `tfsdk:"email"`
	FirstName                 types.String `tfsdk:"first_name"`
	FirstNameKana             types.String `tfsdk:"first_name_kana"`
	FirstNameKanji            types.String `tfsdk:"first_name_kanji"`
	FullNameAliases           types.List   `tfsdk:"full_name_aliases"`
	FutureRequirements        types.Object `tfsdk:"future_requirements"`
	Gender                    types.String `tfsdk:"gender"`
	ID                        types.String `tfsdk:"id"`
	IDNumberProvided          types.Bool   `tfsdk:"id_number_provided"`
	IDNumberSecondaryProvided types.Bool   `tfsdk:"id_number_secondary_provided"`
	LastName                  types.String `tfsdk:"last_name"`
	LastNameKana              types.String `tfsdk:"last_name_kana"`
	LastNameKanji             types.String `tfsdk:"last_name_kanji"`
	MaidenName                types.String `tfsdk:"maiden_name"`
	Metadata                  types.Map    `tfsdk:"metadata"`
	Nationality               types.String `tfsdk:"nationality"`
	Phone                     types.String `tfsdk:"phone"`
	PoliticalExposure         types.String `tfsdk:"political_exposure"`
	RegisteredAddress         types.Object `tfsdk:"registered_address"`
	Relationship              types.Object `tfsdk:"relationship"`
	Requirements              types.Object `tfsdk:"requirements"`
	SSNLast4Provided          types.Bool   `tfsdk:"ssn_last_4_provided"`
	USCfpbData                types.Object `tfsdk:"us_cfpb_data"`
	Verification              types.Object `tfsdk:"verification"`
	Documents                 types.Object `tfsdk:"documents"`
	IDNumber                  types.String `tfsdk:"id_number"`
	IDNumberSecondary         types.String `tfsdk:"id_number_secondary"`
	PersonToken               types.String `tfsdk:"person_token"`
	SSNLast4                  types.String `tfsdk:"ssn_last_4"`
}

func (r *PersonResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stripe.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *stripe.Client, got: %T", req.ProviderData))
		return
	}

	r.client = client
}

func (r *PersonResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_person"
}

func (r *PersonResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "This is an object representing a person associated with a Stripe account.\n\nA platform can only access a subset of data in a person for an account where [account.controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `stripe`, which includes Standard and Express accounts, after creating an Account Link or Account Session to start Connect onboarding.\n\nSee the [Standard onboarding](/connect/standard-accounts) or [Express onboarding](/connect/express-accounts) documentation for information about prefilling information and account onboarding steps. Learn more about [handling identity verification with the API](/connect/handling-api-verification#person-information).",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("person")},
			},
			"account": schema.StringAttribute{
				Required:      true,
				Description:   "The account the person is associated with.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"additional_tos_acceptances": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"account": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Details on the legal guardian's acceptance of the main Stripe service agreement.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"date": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "The Unix timestamp marking when the legal guardian accepted the service agreement.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"ip": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The IP address from which the legal guardian accepted the service agreement.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"user_agent": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The user agent of the browser from which the legal guardian accepted the service agreement.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"address": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"city": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "City, district, suburb, town, or village.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"line1": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Address line 1, such as the street, PO Box, or company name.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"line2": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Address line 2, such as the apartment, suite, unit, or building.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"postal_code": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "ZIP or postal code.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"state": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"address_kana": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The Kana variation of the person's address (Japan only).",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"city": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "City/Ward.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"line1": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Block/Building number.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"line2": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Building details.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"postal_code": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "ZIP or postal code.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"state": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Prefecture.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"town": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Town/cho-me.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"address_kanji": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The Kanji variation of the person's address (Japan only).",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"city": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "City/Ward.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"line1": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Block/Building number.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"line2": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Building details.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"postal_code": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "ZIP or postal code.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"state": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Prefecture.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"town": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Town/cho-me.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"dob": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"day": schema.Int64Attribute{
						Required:    true,
						Description: "The day of birth, between 1 and 31.",
					},
					"month": schema.Int64Attribute{
						Required:    true,
						Description: "The month of birth, between 1 and 12.",
					},
					"year": schema.Int64Attribute{
						Required:    true,
						Description: "The four-digit year of birth.",
					},
				},
			},
			"email": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The person's email address. Also available for accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `stripe`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"first_name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The person's first name. Also available for accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `stripe`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"first_name_kana": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The Kana variation of the person's first name (Japan only). Also available for accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `stripe`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"first_name_kanji": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The Kanji variation of the person's first name (Japan only). Also available for accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `stripe`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"full_name_aliases": schema.ListAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "A list of alternate names or aliases that the person is known by. Also available for accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `stripe`.",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"future_requirements": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Information about the [upcoming new requirements for this person](https://docs.stripe.com/connect/custom-accounts/future-requirements), including what information needs to be collected, and by when.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"alternatives": schema.ListNestedAttribute{
						Computed:      true,
						Description:   "Fields that are due and can be resolved by providing the corresponding alternative fields instead. Many alternatives can list the same `original_fields_due`, and any of these alternatives can serve as a pathway for attempting to resolve the fields again. Re-providing `original_fields_due` also serves as a pathway for attempting to resolve the fields again.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"alternative_fields_due": schema.ListAttribute{
									Computed:      true,
									Description:   "Fields that can be provided to resolve all fields in `original_fields_due`.",
									PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
									ElementType:   types.StringType,
								},
								"original_fields_due": schema.ListAttribute{
									Computed:      true,
									Description:   "Fields that are due and can be resolved by providing all fields in `alternative_fields_due`.",
									PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
									ElementType:   types.StringType,
								},
							},
						},
					},
					"currently_due": schema.ListAttribute{
						Computed:      true,
						Description:   "Fields that need to be resolved to keep the person's account enabled. If not resolved by the account's `future_requirements[current_deadline]`, these fields will transition to the main `requirements` hash, and may immediately become `past_due`, but the account may also be given a grace period depending on the account's enablement state prior to transition.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"errors": schema.ListNestedAttribute{
						Computed:      true,
						Description:   "Details about validation and verification failures for `due` requirements that must be resolved.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"code": schema.StringAttribute{
									Computed:      true,
									Description:   "The code for the type of error.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									Validators:    []validator.String{stringvalidator.OneOf("external_request", "information_missing", "invalid_address_city_state_postal_code", "invalid_address_highway_contract_box", "invalid_address_private_mailbox", "invalid_business_profile_name", "invalid_business_profile_name_denylisted", "invalid_company_name_denylisted", "invalid_dob_age_over_maximum", "invalid_dob_age_under_18", "invalid_dob_age_under_minimum", "invalid_product_description_length", "invalid_product_description_url_match", "invalid_representative_country", "invalid_signator", "invalid_statement_descriptor_business_mismatch", "invalid_statement_descriptor_denylisted", "invalid_statement_descriptor_length", "invalid_statement_descriptor_prefix_denylisted", "invalid_statement_descriptor_prefix_mismatch", "invalid_street_address", "invalid_tax_id", "invalid_tax_id_format", "invalid_tos_acceptance", "invalid_url_denylisted", "invalid_url_format", "invalid_url_length", "invalid_url_web_presence_detected", "invalid_url_website_business_information_mismatch", "invalid_url_website_empty", "invalid_url_website_inaccessible", "invalid_url_website_inaccessible_geoblocked", "invalid_url_website_inaccessible_password_protected", "invalid_url_website_incomplete", "invalid_url_website_incomplete_cancellation_policy", "invalid_url_website_incomplete_customer_service_details", "invalid_url_website_incomplete_legal_restrictions", "invalid_url_website_incomplete_refund_policy", "invalid_url_website_incomplete_return_policy", "invalid_url_website_incomplete_terms_and_conditions", "invalid_url_website_incomplete_under_construction", "invalid_url_website_other", "invalid_value_other", "unsupported_business_type", "verification_directors_mismatch", "verification_document_address_mismatch", "verification_document_address_missing", "verification_document_corrupt", "verification_document_country_not_supported", "verification_document_directors_mismatch", "verification_document_dob_mismatch", "verification_document_duplicate_type", "verification_document_expired", "verification_document_failed_copy", "verification_document_failed_greyscale", "verification_document_failed_other", "verification_document_failed_test_mode", "verification_document_fraudulent", "verification_document_id_number_mismatch", "verification_document_id_number_missing", "verification_document_incomplete", "verification_document_invalid", "verification_document_issue_or_expiry_date_missing", "verification_document_manipulated", "verification_document_missing_back", "verification_document_missing_front", "verification_document_name_mismatch", "verification_document_name_missing", "verification_document_nationality_mismatch", "verification_document_not_readable", "verification_document_not_signed", "verification_document_not_uploaded", "verification_document_photo_mismatch", "verification_document_too_large", "verification_document_type_not_supported", "verification_extraneous_directors", "verification_failed_address_match", "verification_failed_authorizer_authority", "verification_failed_business_iec_number", "verification_failed_document_match", "verification_failed_id_number_match", "verification_failed_keyed_identity", "verification_failed_keyed_match", "verification_failed_name_match", "verification_failed_other", "verification_failed_representative_authority", "verification_failed_residential_address", "verification_failed_tax_id_match", "verification_failed_tax_id_not_issued", "verification_legal_entity_structure_mismatch", "verification_missing_directors", "verification_missing_executives", "verification_missing_owners", "verification_rejected_ownership_exemption_reason", "verification_requires_additional_memorandum_of_associations", "verification_requires_additional_proof_of_registration", "verification_supportability")},
								},
								"reason": schema.StringAttribute{
									Computed:      true,
									Description:   "An informative message that indicates the error type and provides additional details about the error.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"requirement": schema.StringAttribute{
									Computed:      true,
									Description:   "The specific user onboarding requirement field (in the requirements hash) that needs to be resolved.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
							},
						},
					},
					"eventually_due": schema.ListAttribute{
						Computed:      true,
						Description:   "Fields you must collect when all thresholds are reached. As they become required, they appear in `currently_due` as well, and the account's `future_requirements[current_deadline]` becomes set.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"past_due": schema.ListAttribute{
						Computed:      true,
						Description:   "Fields that haven't been resolved by the account's `requirements.current_deadline`. These fields need to be resolved to enable the person's account. `future_requirements.past_due` is a subset of `requirements.past_due`.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"pending_verification": schema.ListAttribute{
						Computed:      true,
						Description:   "Fields that are being reviewed, or might become required depending on the results of a review. If the review fails, these fields can move to `eventually_due`, `currently_due`, `past_due` or `alternatives`. Fields might appear in `eventually_due`, `currently_due`, `past_due` or `alternatives` and in `pending_verification` if one verification fails but another is still pending.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
				},
			},
			"gender": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The person's gender.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"id_number_provided": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether the person's `id_number` was provided. True if either the full ID number was provided or if only the required part of the ID number was provided (ex. last four of an individual's SSN for the US indicated by `ssn_last_4_provided`).",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"id_number_secondary_provided": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether the person's `id_number_secondary` was provided.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"last_name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The person's last name. Also available for accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `stripe`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"last_name_kana": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The Kana variation of the person's last name (Japan only). Also available for accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `stripe`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"last_name_kanji": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The Kanji variation of the person's last name (Japan only). Also available for accounts where [controller.requirement_collection](/api/accounts/object#account_object-controller-requirement_collection) is `stripe`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"maiden_name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The person's maiden name.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"nationality": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The country where the person is a national.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"phone": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The person's phone number.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"political_exposure": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Indicates if the person or any of their representatives, family members, or other closely related persons, declares that they hold or have held an important public job or function, in any jurisdiction.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("existing", "none")},
			},
			"registered_address": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"city": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "City, district, suburb, town, or village.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"country": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"line1": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Address line 1, such as the street, PO Box, or company name.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"line2": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Address line 2, such as the apartment, suite, unit, or building.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"postal_code": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "ZIP or postal code.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"state": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"relationship": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"authorizer": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Whether the person is the authorizer of the account's representative.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"director": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Whether the person is a director of the account's legal entity. Directors are typically members of the governing board of the company, or responsible for ensuring the company meets its regulatory obligations.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"executive": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Whether the person has significant responsibility to control, manage, or direct the organization.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"legal_guardian": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Whether the person is the legal guardian of the account's representative.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"owner": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Whether the person is an owner of the account’s legal entity.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"percent_ownership": schema.Float64Attribute{
						Optional:      true,
						Computed:      true,
						Description:   "The percent owned by the person of the account's legal entity.",
						PlanModifiers: []planmodifier.Float64{float64planmodifier.UseStateForUnknown()},
					},
					"representative": schema.BoolAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "Whether the person is authorized as the primary representative of the account. This is the person nominated by the business to provide information about themselves, and general information about the account. There can only be one representative at any given time. At the time the account is created, this person should be set to the person responsible for opening the account.",
						PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
					},
					"title": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The person's title (e.g., CEO, Support Engineer).",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"requirements": schema.SingleNestedAttribute{
				Computed:      true,
				Description:   "Information about the requirements for this person, including what information needs to be collected, and by when.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"alternatives": schema.ListNestedAttribute{
						Computed:      true,
						Description:   "Fields that are due and can be resolved by providing the corresponding alternative fields instead. Many alternatives can list the same `original_fields_due`, and any of these alternatives can serve as a pathway for attempting to resolve the fields again. Re-providing `original_fields_due` also serves as a pathway for attempting to resolve the fields again.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"alternative_fields_due": schema.ListAttribute{
									Computed:      true,
									Description:   "Fields that can be provided to resolve all fields in `original_fields_due`.",
									PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
									ElementType:   types.StringType,
								},
								"original_fields_due": schema.ListAttribute{
									Computed:      true,
									Description:   "Fields that are due and can be resolved by providing all fields in `alternative_fields_due`.",
									PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
									ElementType:   types.StringType,
								},
							},
						},
					},
					"currently_due": schema.ListAttribute{
						Computed:      true,
						Description:   "Fields that need to be resolved to keep the person's account enabled. If not resolved by the account's `current_deadline`, these fields will appear in `past_due` as well, and the account is disabled.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"errors": schema.ListNestedAttribute{
						Computed:      true,
						Description:   "Details about validation and verification failures for `due` requirements that must be resolved.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"code": schema.StringAttribute{
									Computed:      true,
									Description:   "The code for the type of error.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
									Validators:    []validator.String{stringvalidator.OneOf("external_request", "information_missing", "invalid_address_city_state_postal_code", "invalid_address_highway_contract_box", "invalid_address_private_mailbox", "invalid_business_profile_name", "invalid_business_profile_name_denylisted", "invalid_company_name_denylisted", "invalid_dob_age_over_maximum", "invalid_dob_age_under_18", "invalid_dob_age_under_minimum", "invalid_product_description_length", "invalid_product_description_url_match", "invalid_representative_country", "invalid_signator", "invalid_statement_descriptor_business_mismatch", "invalid_statement_descriptor_denylisted", "invalid_statement_descriptor_length", "invalid_statement_descriptor_prefix_denylisted", "invalid_statement_descriptor_prefix_mismatch", "invalid_street_address", "invalid_tax_id", "invalid_tax_id_format", "invalid_tos_acceptance", "invalid_url_denylisted", "invalid_url_format", "invalid_url_length", "invalid_url_web_presence_detected", "invalid_url_website_business_information_mismatch", "invalid_url_website_empty", "invalid_url_website_inaccessible", "invalid_url_website_inaccessible_geoblocked", "invalid_url_website_inaccessible_password_protected", "invalid_url_website_incomplete", "invalid_url_website_incomplete_cancellation_policy", "invalid_url_website_incomplete_customer_service_details", "invalid_url_website_incomplete_legal_restrictions", "invalid_url_website_incomplete_refund_policy", "invalid_url_website_incomplete_return_policy", "invalid_url_website_incomplete_terms_and_conditions", "invalid_url_website_incomplete_under_construction", "invalid_url_website_other", "invalid_value_other", "unsupported_business_type", "verification_directors_mismatch", "verification_document_address_mismatch", "verification_document_address_missing", "verification_document_corrupt", "verification_document_country_not_supported", "verification_document_directors_mismatch", "verification_document_dob_mismatch", "verification_document_duplicate_type", "verification_document_expired", "verification_document_failed_copy", "verification_document_failed_greyscale", "verification_document_failed_other", "verification_document_failed_test_mode", "verification_document_fraudulent", "verification_document_id_number_mismatch", "verification_document_id_number_missing", "verification_document_incomplete", "verification_document_invalid", "verification_document_issue_or_expiry_date_missing", "verification_document_manipulated", "verification_document_missing_back", "verification_document_missing_front", "verification_document_name_mismatch", "verification_document_name_missing", "verification_document_nationality_mismatch", "verification_document_not_readable", "verification_document_not_signed", "verification_document_not_uploaded", "verification_document_photo_mismatch", "verification_document_too_large", "verification_document_type_not_supported", "verification_extraneous_directors", "verification_failed_address_match", "verification_failed_authorizer_authority", "verification_failed_business_iec_number", "verification_failed_document_match", "verification_failed_id_number_match", "verification_failed_keyed_identity", "verification_failed_keyed_match", "verification_failed_name_match", "verification_failed_other", "verification_failed_representative_authority", "verification_failed_residential_address", "verification_failed_tax_id_match", "verification_failed_tax_id_not_issued", "verification_legal_entity_structure_mismatch", "verification_missing_directors", "verification_missing_executives", "verification_missing_owners", "verification_rejected_ownership_exemption_reason", "verification_requires_additional_memorandum_of_associations", "verification_requires_additional_proof_of_registration", "verification_supportability")},
								},
								"reason": schema.StringAttribute{
									Computed:      true,
									Description:   "An informative message that indicates the error type and provides additional details about the error.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
								"requirement": schema.StringAttribute{
									Computed:      true,
									Description:   "The specific user onboarding requirement field (in the requirements hash) that needs to be resolved.",
									PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								},
							},
						},
					},
					"eventually_due": schema.ListAttribute{
						Computed:      true,
						Description:   "Fields you must collect when all thresholds are reached. As they become required, they appear in `currently_due` as well, and the account's `current_deadline` becomes set.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"past_due": schema.ListAttribute{
						Computed:      true,
						Description:   "Fields that haven't been resolved by `current_deadline`. These fields need to be resolved to enable the person's account.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
					"pending_verification": schema.ListAttribute{
						Computed:      true,
						Description:   "Fields that are being reviewed, or might become required depending on the results of a review. If the review fails, these fields can move to `eventually_due`, `currently_due`, `past_due` or `alternatives`. Fields might appear in `eventually_due`, `currently_due`, `past_due` or `alternatives` and in `pending_verification` if one verification fails but another is still pending.",
						PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
						ElementType:   types.StringType,
					},
				},
			},
			"ssn_last_4_provided": schema.BoolAttribute{
				Computed:      true,
				Description:   "Whether the last four digits of the person's Social Security number have been provided (U.S. only).",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"us_cfpb_data": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Demographic data related to the person.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"ethnicity_details": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The persons ethnicity details",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"ethnicity": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The persons ethnicity",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"ethnicity_other": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Please specify your origin, when other is selected.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"race_details": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The persons race details",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"race": schema.ListAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The persons race.",
								PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
								ElementType:   types.StringType,
							},
							"race_other": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Please specify your race, when other is selected.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"self_identified_gender": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The persons self-identified gender",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"verification": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"additional_document": schema.SingleNestedAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "A document showing address, either a passport, local ID card, or utility bill from a well-known utility company.",
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"back": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The back of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"details": schema.StringAttribute{
								Computed:      true,
								Description:   "A user-displayable string describing the verification state of this document. For example, if a document is uploaded and the picture is too fuzzy, this may say \"Identity document is too unclear to read\".",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"details_code": schema.StringAttribute{
								Computed:      true,
								Description:   "One of `document_corrupt`, `document_country_not_supported`, `document_expired`, `document_failed_copy`, `document_failed_other`, `document_failed_test_mode`, `document_fraudulent`, `document_failed_greyscale`, `document_incomplete`, `document_invalid`, `document_manipulated`, `document_missing_back`, `document_missing_front`, `document_not_readable`, `document_not_uploaded`, `document_photo_mismatch`, `document_too_large`, or `document_type_not_supported`. A machine-readable code specifying the verification state for this document.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"front": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The front of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"details": schema.StringAttribute{
						Computed:      true,
						Description:   "A user-displayable string describing the verification state for the person. For example, this may say \"Provided identity information could not be verified\".",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"details_code": schema.StringAttribute{
						Computed:      true,
						Description:   "One of `document_address_mismatch`, `document_dob_mismatch`, `document_duplicate_type`, `document_id_number_mismatch`, `document_name_mismatch`, `document_nationality_mismatch`, `failed_keyed_identity`, or `failed_other`. A machine-readable code specifying the verification state for the person.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"document": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"back": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The back of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"details": schema.StringAttribute{
								Computed:      true,
								Description:   "A user-displayable string describing the verification state of this document. For example, if a document is uploaded and the picture is too fuzzy, this may say \"Identity document is too unclear to read\".",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"details_code": schema.StringAttribute{
								Computed:      true,
								Description:   "One of `document_corrupt`, `document_country_not_supported`, `document_expired`, `document_failed_copy`, `document_failed_other`, `document_failed_test_mode`, `document_fraudulent`, `document_failed_greyscale`, `document_incomplete`, `document_invalid`, `document_manipulated`, `document_missing_back`, `document_missing_front`, `document_not_readable`, `document_not_uploaded`, `document_photo_mismatch`, `document_too_large`, or `document_type_not_supported`. A machine-readable code specifying the verification state for this document.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"front": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "The front of an ID returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `identity_document`.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"status": schema.StringAttribute{
						Computed:      true,
						Description:   "The state of verification for the person. Possible values are `unverified`, `pending`, or `verified`. Please refer [guide](https://docs.stripe.com/connect/handling-api-verification) to handle verification updates.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
				},
			},
			"documents": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Documents that may be submitted to satisfy various informational requests.",
				WriteOnly:   true,
				Attributes: map[string]schema.Attribute{
					"company_authorization": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "One or more documents that demonstrate proof that this person is authorized to represent the company.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"files": schema.ListAttribute{
								Optional:    true,
								Description: "One or more document ids returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `account_requirement`.",
								WriteOnly:   true,
								ElementType: types.StringType,
							},
						},
					},
					"passport": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "One or more documents showing the person's passport page with photo and personal data.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"files": schema.ListAttribute{
								Optional:    true,
								Description: "One or more document ids returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `account_requirement`.",
								WriteOnly:   true,
								ElementType: types.StringType,
							},
						},
					},
					"visa": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "One or more documents showing the person's visa required for living in the country where they are residing.",
						WriteOnly:   true,
						Attributes: map[string]schema.Attribute{
							"files": schema.ListAttribute{
								Optional:    true,
								Description: "One or more document ids returned by a [file upload](https://api.stripe.com#create_file) with a `purpose` value of `account_requirement`.",
								WriteOnly:   true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"id_number": schema.StringAttribute{
				Optional:    true,
				Description: "The person's ID number, as appropriate for their country. For example, a social security number in the U.S., social insurance number in Canada, etc. Instead of the number itself, you can also provide a [PII token provided by Stripe.js](https://docs.stripe.com/js/tokens/create_token?type=pii).",
				WriteOnly:   true,
			},
			"id_number_secondary": schema.StringAttribute{
				Optional:    true,
				Description: "The person's secondary ID number, as appropriate for their country, will be used for enhanced verification checks. In Thailand, this would be the laser code found on the back of an ID card. Instead of the number itself, you can also provide a [PII token provided by Stripe.js](https://docs.stripe.com/js/tokens/create_token?type=pii).",
				WriteOnly:   true,
			},
			"person_token": schema.StringAttribute{
				Optional:    true,
				Description: "A [person token](https://docs.stripe.com/connect/account-tokens), used to securely provide details to the person.",
				WriteOnly:   true,
			},
			"ssn_last_4": schema.StringAttribute{
				Optional:    true,
				Description: "The last four digits of the person's Social Security number (U.S. only).",
				WriteOnly:   true,
			},
		},
	}
}

func (r *PersonResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PersonResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PersonResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Documents"}, []string{"Documents", "company_authorization"}, []string{"Documents", "company_authorization", "files"}, []string{"Documents", "passport"}, []string{"Documents", "passport", "files"}, []string{"Documents", "visa"}, []string{"Documents", "visa", "files"}, []string{"IDNumber"}, []string{"IDNumberSecondary"}, []string{"PersonToken"}, []string{"SSNLast4"}})

	params, err := expandPersonCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building Person create params", err.Error())
		return
	}

	if !plan.Account.IsNull() && !plan.Account.IsUnknown() {
		if !assignStringToNamedField(params, "Account", "ID", plan.Account.ValueString()) {
			resp.Diagnostics.AddError("Error building Person create path params", fmt.Sprintf("Failed to assign path parameter %q on %T", "account", params))
			return
		}
	}
	obj, err := r.client.V1Persons.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Person", err.Error())
		return
	}

	rawReadParams := &stripe.PersonRetrieveParams{}
	if !plan.Account.IsNull() && !plan.Account.IsUnknown() {
		if !assignStringToNamedField(rawReadParams, "Account", "ID", plan.Account.ValueString()) {
			resp.Diagnostics.AddError("Error building Person read params for raw hydration after create", fmt.Sprintf("Failed to assign path parameter %q on %T", "account", rawReadParams))
			return
		}
	}

	if err := ensureRawResponse(obj, r.client.V1Persons.B, r.client.V1Persons.Key, stripe.FormatURLPath("/v1/accounts/%s/persons/%s", plan.Account.ValueString(), obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating Person create raw response", err.Error())
		return
	}

	if err := flattenPerson(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Person create response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Documents"}, []string{"Documents", "company_authorization"}, []string{"Documents", "company_authorization", "files"}, []string{"Documents", "passport"}, []string{"Documents", "passport", "files"}, []string{"Documents", "visa"}, []string{"Documents", "visa", "files"}, []string{"IDNumber"}, []string{"IDNumberSecondary"}, []string{"PersonToken"}, []string{"SSNLast4"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PersonResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState PersonResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state PersonResourceModel
	state = priorState

	params := &stripe.PersonRetrieveParams{}
	if !state.Account.IsNull() && !state.Account.IsUnknown() {
		if !assignStringToNamedField(params, "Account", "ID", state.Account.ValueString()) {
			resp.Diagnostics.AddError("Error building Person read params", fmt.Sprintf("Failed to assign path parameter %q on %T", "account", params))
			return
		}
	}

	obj, err := r.client.V1Persons.Retrieve(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Person", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1Persons.B, r.client.V1Persons.Key, stripe.FormatURLPath("/v1/accounts/%s/persons/%s", state.Account.ValueString(), state.ID.ValueString()), params); err != nil {
		resp.Diagnostics.AddError("Error hydrating Person raw response", err.Error())
		return
	}

	if err := flattenPerson(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening Person read response", err.Error())
		return
	}
	clearWriteOnlyPaths(&state, [][]string{[]string{"Documents"}, []string{"Documents", "company_authorization"}, []string{"Documents", "company_authorization", "files"}, []string{"Documents", "passport"}, []string{"Documents", "passport", "files"}, []string{"Documents", "visa"}, []string{"Documents", "visa", "files"}, []string{"IDNumber"}, []string{"IDNumberSecondary"}, []string{"PersonToken"}, []string{"SSNLast4"}})
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *PersonResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PersonResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config PersonResourceModel
	diags = req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	hydrateWriteOnlyPaths(&plan, &config, [][]string{[]string{"Documents"}, []string{"Documents", "company_authorization"}, []string{"Documents", "company_authorization", "files"}, []string{"Documents", "passport"}, []string{"Documents", "passport", "files"}, []string{"Documents", "visa"}, []string{"Documents", "visa", "files"}, []string{"IDNumber"}, []string{"IDNumberSecondary"}, []string{"PersonToken"}, []string{"SSNLast4"}})

	var state PersonResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandPersonUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building Person update params", err.Error())
		return
	}

	if !state.Account.IsNull() && !state.Account.IsUnknown() {
		if !assignStringToNamedField(params, "Account", "ID", state.Account.ValueString()) {
			resp.Diagnostics.AddError("Error building Person update path params", fmt.Sprintf("Failed to assign path parameter %q on %T", "account", params))
			return
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building Person update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1Persons.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Person", err.Error())
		return
	}

	rawReadParams := &stripe.PersonRetrieveParams{}
	if !state.Account.IsNull() && !state.Account.IsUnknown() {
		if !assignStringToNamedField(rawReadParams, "Account", "ID", state.Account.ValueString()) {
			resp.Diagnostics.AddError("Error building Person read params for raw hydration after update", fmt.Sprintf("Failed to assign path parameter %q on %T", "account", rawReadParams))
			return
		}
	}

	if err := ensureRawResponse(obj, r.client.V1Persons.B, r.client.V1Persons.Key, stripe.FormatURLPath("/v1/accounts/%s/persons/%s", state.Account.ValueString(), obj.ID), rawReadParams); err != nil {
		resp.Diagnostics.AddError("Error hydrating Person update raw response", err.Error())
		return
	}

	if err := flattenPerson(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening Person update response", err.Error())
		return
	}
	clearWriteOnlyPaths(&plan, [][]string{[]string{"Documents"}, []string{"Documents", "company_authorization"}, []string{"Documents", "company_authorization", "files"}, []string{"Documents", "passport"}, []string{"Documents", "passport", "files"}, []string{"Documents", "visa"}, []string{"Documents", "visa", "files"}, []string{"IDNumber"}, []string{"IDNumberSecondary"}, []string{"PersonToken"}, []string{"SSNLast4"}})
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *PersonResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PersonResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params := &stripe.PersonDeleteParams{}
	if !state.Account.IsNull() && !state.Account.IsUnknown() {
		if !assignStringToNamedField(params, "Account", "ID", state.Account.ValueString()) {
			resp.Diagnostics.AddError("Error building Person delete params", fmt.Sprintf("Failed to assign path parameter %q on %T", "account", params))
			return
		}
	}

	_, err := r.client.V1Persons.Delete(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting Person", err.Error())
		return
	}
}

func (r *PersonResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Unexpected import identifier", fmt.Sprintf("Expected import identifier in the form \"account/id\", got %q", req.ID))
		return
	}

	diags := resp.State.SetAttribute(ctx, path.Root("account"), types.StringValue(parts[0]))
	resp.Diagnostics.Append(diags...)
	diags = resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(parts[1]))
	resp.Diagnostics.Append(diags...)
}

func expandPersonCreate(plan PersonResourceModel) (*stripe.PersonCreateParams, error) {
	params := &stripe.PersonCreateParams{}

	if !plan.AdditionalTOSAcceptances.IsNull() && !plan.AdditionalTOSAcceptances.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AdditionalTOSAcceptances", plan.AdditionalTOSAcceptances) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "additional_tos_acceptances", params)
		}
	}
	if !plan.Address.IsNull() && !plan.Address.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Address", plan.Address) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "address", params)
		}
	}
	if !plan.AddressKana.IsNull() && !plan.AddressKana.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AddressKana", plan.AddressKana) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "address_kana", params)
		}
	}
	if !plan.AddressKanji.IsNull() && !plan.AddressKanji.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AddressKanji", plan.AddressKanji) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "address_kanji", params)
		}
	}
	if !plan.DOB.IsNull() && !plan.DOB.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DOB", plan.DOB) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "dob", params)
		}
	}
	if !plan.Email.IsNull() && !plan.Email.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Email", "Email", plan.Email.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "email", params)
		}
	}
	if !plan.FirstName.IsNull() && !plan.FirstName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FirstName", "FirstName", plan.FirstName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "first_name", params)
		}
	}
	if !plan.FirstNameKana.IsNull() && !plan.FirstNameKana.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FirstNameKana", "FirstNameKana", plan.FirstNameKana.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "first_name_kana", params)
		}
	}
	if !plan.FirstNameKanji.IsNull() && !plan.FirstNameKanji.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FirstNameKanji", "FirstNameKanji", plan.FirstNameKanji.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "first_name_kanji", params)
		}
	}
	if !plan.FullNameAliases.IsNull() && !plan.FullNameAliases.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FullNameAliases", plan.FullNameAliases) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "full_name_aliases", params)
		}
	}
	if !plan.Gender.IsNull() && !plan.Gender.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Gender", "Gender", plan.Gender.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "gender", params)
		}
	}
	if !plan.LastName.IsNull() && !plan.LastName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LastName", "LastName", plan.LastName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "last_name", params)
		}
	}
	if !plan.LastNameKana.IsNull() && !plan.LastNameKana.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LastNameKana", "LastNameKana", plan.LastNameKana.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "last_name_kana", params)
		}
	}
	if !plan.LastNameKanji.IsNull() && !plan.LastNameKanji.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LastNameKanji", "LastNameKanji", plan.LastNameKanji.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "last_name_kanji", params)
		}
	}
	if !plan.MaidenName.IsNull() && !plan.MaidenName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "MaidenName", "MaidenName", plan.MaidenName.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "maiden_name", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Nationality.IsNull() && !plan.Nationality.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Nationality", "Nationality", plan.Nationality.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "nationality", params)
		}
	}
	if !plan.Phone.IsNull() && !plan.Phone.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Phone", "Phone", plan.Phone.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "phone", params)
		}
	}
	if !plan.PoliticalExposure.IsNull() && !plan.PoliticalExposure.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PoliticalExposure", "PoliticalExposure", plan.PoliticalExposure.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "political_exposure", params)
		}
	}
	if !plan.RegisteredAddress.IsNull() && !plan.RegisteredAddress.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RegisteredAddress", plan.RegisteredAddress) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "registered_address", params)
		}
	}
	if !plan.Relationship.IsNull() && !plan.Relationship.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Relationship", plan.Relationship) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "relationship", params)
		}
	}
	if !plan.USCfpbData.IsNull() && !plan.USCfpbData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "USCfpbData", plan.USCfpbData) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "us_cfpb_data", params)
		}
	}
	if !plan.Verification.IsNull() && !plan.Verification.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Verification", plan.Verification) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "verification", params)
		}
	}
	if !plan.Documents.IsNull() && !plan.Documents.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Documents", plan.Documents) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "documents", params)
		}
	}
	if !plan.IDNumber.IsNull() && !plan.IDNumber.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "IDNumber", "IDNumber", plan.IDNumber.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "id_number", params)
		}
	}
	if !plan.IDNumberSecondary.IsNull() && !plan.IDNumberSecondary.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "IDNumberSecondary", "IDNumberSecondary", plan.IDNumberSecondary.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "id_number_secondary", params)
		}
	}
	if !plan.PersonToken.IsNull() && !plan.PersonToken.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PersonToken", "PersonToken", plan.PersonToken.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "person_token", params)
		}
	}
	if !plan.SSNLast4.IsNull() && !plan.SSNLast4.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SSNLast4", "SSNLast4", plan.SSNLast4.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "ssn_last_4", params)
		}
	}

	return params, nil
}

func expandPersonUpdate(plan PersonResourceModel, state PersonResourceModel) (*stripe.PersonUpdateParams, error) {
	params := &stripe.PersonUpdateParams{}

	if !plan.AdditionalTOSAcceptances.Equal(state.AdditionalTOSAcceptances) && !plan.AdditionalTOSAcceptances.IsNull() && !plan.AdditionalTOSAcceptances.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AdditionalTOSAcceptances", plan.AdditionalTOSAcceptances) {
			if !plan.AdditionalTOSAcceptances.Equal(state.AdditionalTOSAcceptances) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "additional_tos_acceptances", params)
			}
		}
	}
	if !plan.Address.Equal(state.Address) && !plan.Address.IsNull() && !plan.Address.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Address", plan.Address) {
			if !plan.Address.Equal(state.Address) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "address", params)
			}
		}
	}
	if !plan.AddressKana.Equal(state.AddressKana) && !plan.AddressKana.IsNull() && !plan.AddressKana.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AddressKana", plan.AddressKana) {
			if !plan.AddressKana.Equal(state.AddressKana) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "address_kana", params)
			}
		}
	}
	if !plan.AddressKanji.Equal(state.AddressKanji) && !plan.AddressKanji.IsNull() && !plan.AddressKanji.IsUnknown() {
		if !assignAttrValueToNamedField(params, "AddressKanji", plan.AddressKanji) {
			if !plan.AddressKanji.Equal(state.AddressKanji) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "address_kanji", params)
			}
		}
	}
	if !plan.DOB.Equal(state.DOB) && !plan.DOB.IsNull() && !plan.DOB.IsUnknown() {
		if !assignAttrValueToNamedField(params, "DOB", plan.DOB) {
			if !plan.DOB.Equal(state.DOB) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "dob", params)
			}
		}
	}
	if !plan.Email.Equal(state.Email) && !plan.Email.IsNull() && !plan.Email.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Email", "Email", plan.Email.ValueString()) {
			if !plan.Email.Equal(state.Email) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "email", params)
			}
		}
	}
	if !plan.FirstName.Equal(state.FirstName) && !plan.FirstName.IsNull() && !plan.FirstName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FirstName", "FirstName", plan.FirstName.ValueString()) {
			if !plan.FirstName.Equal(state.FirstName) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "first_name", params)
			}
		}
	}
	if !plan.FirstNameKana.Equal(state.FirstNameKana) && !plan.FirstNameKana.IsNull() && !plan.FirstNameKana.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FirstNameKana", "FirstNameKana", plan.FirstNameKana.ValueString()) {
			if !plan.FirstNameKana.Equal(state.FirstNameKana) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "first_name_kana", params)
			}
		}
	}
	if !plan.FirstNameKanji.Equal(state.FirstNameKanji) && !plan.FirstNameKanji.IsNull() && !plan.FirstNameKanji.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "FirstNameKanji", "FirstNameKanji", plan.FirstNameKanji.ValueString()) {
			if !plan.FirstNameKanji.Equal(state.FirstNameKanji) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "first_name_kanji", params)
			}
		}
	}
	if !plan.FullNameAliases.Equal(state.FullNameAliases) && !plan.FullNameAliases.IsNull() && !plan.FullNameAliases.IsUnknown() {
		if !assignAttrValueToNamedField(params, "FullNameAliases", plan.FullNameAliases) {
			if !plan.FullNameAliases.Equal(state.FullNameAliases) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "full_name_aliases", params)
			}
		}
	}
	if !plan.Gender.Equal(state.Gender) && !plan.Gender.IsNull() && !plan.Gender.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Gender", "Gender", plan.Gender.ValueString()) {
			if !plan.Gender.Equal(state.Gender) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "gender", params)
			}
		}
	}
	if !plan.LastName.Equal(state.LastName) && !plan.LastName.IsNull() && !plan.LastName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LastName", "LastName", plan.LastName.ValueString()) {
			if !plan.LastName.Equal(state.LastName) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "last_name", params)
			}
		}
	}
	if !plan.LastNameKana.Equal(state.LastNameKana) && !plan.LastNameKana.IsNull() && !plan.LastNameKana.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LastNameKana", "LastNameKana", plan.LastNameKana.ValueString()) {
			if !plan.LastNameKana.Equal(state.LastNameKana) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "last_name_kana", params)
			}
		}
	}
	if !plan.LastNameKanji.Equal(state.LastNameKanji) && !plan.LastNameKanji.IsNull() && !plan.LastNameKanji.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "LastNameKanji", "LastNameKanji", plan.LastNameKanji.ValueString()) {
			if !plan.LastNameKanji.Equal(state.LastNameKanji) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "last_name_kanji", params)
			}
		}
	}
	if !plan.MaidenName.Equal(state.MaidenName) && !plan.MaidenName.IsNull() && !plan.MaidenName.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "MaidenName", "MaidenName", plan.MaidenName.ValueString()) {
			if !plan.MaidenName.Equal(state.MaidenName) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "maiden_name", params)
			}
		}
	}
	if !plan.Metadata.Equal(state.Metadata) && !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			if !plan.Metadata.Equal(state.Metadata) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "metadata", params)
			}
		}
	}
	if !plan.Nationality.Equal(state.Nationality) && !plan.Nationality.IsNull() && !plan.Nationality.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Nationality", "Nationality", plan.Nationality.ValueString()) {
			if !plan.Nationality.Equal(state.Nationality) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "nationality", params)
			}
		}
	}
	if !plan.Phone.Equal(state.Phone) && !plan.Phone.IsNull() && !plan.Phone.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Phone", "Phone", plan.Phone.ValueString()) {
			if !plan.Phone.Equal(state.Phone) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "phone", params)
			}
		}
	}
	if !plan.PoliticalExposure.Equal(state.PoliticalExposure) && !plan.PoliticalExposure.IsNull() && !plan.PoliticalExposure.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PoliticalExposure", "PoliticalExposure", plan.PoliticalExposure.ValueString()) {
			if !plan.PoliticalExposure.Equal(state.PoliticalExposure) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "political_exposure", params)
			}
		}
	}
	if !plan.RegisteredAddress.Equal(state.RegisteredAddress) && !plan.RegisteredAddress.IsNull() && !plan.RegisteredAddress.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RegisteredAddress", plan.RegisteredAddress) {
			if !plan.RegisteredAddress.Equal(state.RegisteredAddress) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "registered_address", params)
			}
		}
	}
	if !plan.Relationship.Equal(state.Relationship) && !plan.Relationship.IsNull() && !plan.Relationship.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Relationship", plan.Relationship) {
			if !plan.Relationship.Equal(state.Relationship) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "relationship", params)
			}
		}
	}
	if !plan.USCfpbData.Equal(state.USCfpbData) && !plan.USCfpbData.IsNull() && !plan.USCfpbData.IsUnknown() {
		if !assignAttrValueToNamedField(params, "USCfpbData", plan.USCfpbData) {
			if !plan.USCfpbData.Equal(state.USCfpbData) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "us_cfpb_data", params)
			}
		}
	}
	if !plan.Verification.Equal(state.Verification) && !plan.Verification.IsNull() && !plan.Verification.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Verification", plan.Verification) {
			if !plan.Verification.Equal(state.Verification) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "verification", params)
			}
		}
	}
	if !plan.Documents.Equal(state.Documents) && !plan.Documents.IsNull() && !plan.Documents.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Documents", plan.Documents) {
			if !plan.Documents.Equal(state.Documents) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "documents", params)
			}
		}
	}
	if !plan.IDNumber.Equal(state.IDNumber) && !plan.IDNumber.IsNull() && !plan.IDNumber.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "IDNumber", "IDNumber", plan.IDNumber.ValueString()) {
			if !plan.IDNumber.Equal(state.IDNumber) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "id_number", params)
			}
		}
	}
	if !plan.IDNumberSecondary.Equal(state.IDNumberSecondary) && !plan.IDNumberSecondary.IsNull() && !plan.IDNumberSecondary.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "IDNumberSecondary", "IDNumberSecondary", plan.IDNumberSecondary.ValueString()) {
			if !plan.IDNumberSecondary.Equal(state.IDNumberSecondary) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "id_number_secondary", params)
			}
		}
	}
	if !plan.PersonToken.Equal(state.PersonToken) && !plan.PersonToken.IsNull() && !plan.PersonToken.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "PersonToken", "PersonToken", plan.PersonToken.ValueString()) {
			if !plan.PersonToken.Equal(state.PersonToken) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "person_token", params)
			}
		}
	}
	if !plan.SSNLast4.Equal(state.SSNLast4) && !plan.SSNLast4.IsNull() && !plan.SSNLast4.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "SSNLast4", "SSNLast4", plan.SSNLast4.ValueString()) {
			if !plan.SSNLast4.Equal(state.SSNLast4) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "ssn_last_4", params)
			}
		}
	}

	return params, nil
}

func flattenPerson(obj *stripe.Person, state *PersonResourceModel) error {
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
		if rawValueAccount, rawOk := plainValueAtPath(raw, "account"); rawOk {
			if valueAccount, err := flattenPlainValue(rawValueAccount, types.StringType, "account", "raw response"); err != nil {
				return err
			} else {
				if typedAccount, ok := valueAccount.(types.String); ok {
					state.Account = typedAccount
				}
			}
		} else if !hasRaw {
			if responseValueAccount, ok := plainFromResponseField(obj, "Account"); ok {
				if valueAccount, err := flattenPlainValue(responseValueAccount, types.StringType, "account", "response struct"); err != nil {
					return err
				} else {
					if typedAccount, ok := valueAccount.(types.String); ok {
						state.Account = typedAccount
					}
				}
			}
		}
	}
	{
		assignedAdditionalTOSAcceptances := false
		hadRawAdditionalTOSAcceptances := false
		if rawValueAdditionalTOSAcceptances, rawOk := plainValueAtPath(raw, "additional_tos_acceptances"); rawOk {
			hadRawAdditionalTOSAcceptances = true
			if rawValueAdditionalTOSAcceptances != nil {
				sourceAdditionalTOSAcceptances := applyConfiguredKeyedListShapes(rawValueAdditionalTOSAcceptances, attrValueToPlain(state.AdditionalTOSAcceptances))
				if valueAdditionalTOSAcceptances, err := flattenPlainValue(sourceAdditionalTOSAcceptances, types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.Int64Type, "ip": types.StringType, "user_agent": types.StringType}}}}, "additional_tos_acceptances", "raw response"); err != nil {
					return err
				} else {
					if typedAdditionalTOSAcceptances, ok := valueAdditionalTOSAcceptances.(types.Object); ok {
						state.AdditionalTOSAcceptances = typedAdditionalTOSAcceptances
						assignedAdditionalTOSAcceptances = true
					}
				}
			}
		}
		if !assignedAdditionalTOSAcceptances {
			if !hasRaw {
				if responseValueAdditionalTOSAcceptances, ok := plainFromResponseField(obj, "AdditionalTOSAcceptances"); ok {
					sourceAdditionalTOSAcceptances := applyConfiguredKeyedListShapes(responseValueAdditionalTOSAcceptances, attrValueToPlain(state.AdditionalTOSAcceptances))
					if valueAdditionalTOSAcceptances, err := flattenPlainValue(
						sourceAdditionalTOSAcceptances,
						types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.Int64Type, "ip": types.StringType, "user_agent": types.StringType}}}},
						"additional_tos_acceptances",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAdditionalTOSAcceptances, ok := valueAdditionalTOSAcceptances.(types.Object); ok {
							state.AdditionalTOSAcceptances = typedAdditionalTOSAcceptances
							assignedAdditionalTOSAcceptances = true
						}
					}
				}
			}
		}
		if !assignedAdditionalTOSAcceptances && hadRawAdditionalTOSAcceptances {
			if nullAdditionalTOSAcceptances, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.ObjectType{AttrTypes: map[string]attr.Type{"date": types.Int64Type, "ip": types.StringType, "user_agent": types.StringType}}}}); ok {
				if typedAdditionalTOSAcceptances, ok := nullAdditionalTOSAcceptances.(types.Object); ok {
					state.AdditionalTOSAcceptances = typedAdditionalTOSAcceptances
				}
			}
		}
	}
	{
		assignedAddress := false
		hadRawAddress := false
		if rawValueAddress, rawOk := plainValueAtPath(raw, "address"); rawOk {
			hadRawAddress = true
			if rawValueAddress != nil {
				sourceAddress := applyConfiguredKeyedListShapes(rawValueAddress, attrValueToPlain(state.Address))
				if valueAddress, err := flattenPlainValue(sourceAddress, types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "address", "raw response"); err != nil {
					return err
				} else {
					if typedAddress, ok := valueAddress.(types.Object); ok {
						state.Address = typedAddress
						assignedAddress = true
					}
				}
			}
		}
		if !assignedAddress {
			if !hasRaw {
				if responseValueAddress, ok := plainFromResponseField(obj, "Address"); ok {
					sourceAddress := applyConfiguredKeyedListShapes(responseValueAddress, attrValueToPlain(state.Address))
					if valueAddress, err := flattenPlainValue(
						sourceAddress,
						types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}},
						"address",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAddress, ok := valueAddress.(types.Object); ok {
							state.Address = typedAddress
							assignedAddress = true
						}
					}
				}
			}
		}
		if !assignedAddress && hadRawAddress {
			if nullAddress, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}); ok {
				if typedAddress, ok := nullAddress.(types.Object); ok {
					state.Address = typedAddress
				}
			}
		}
	}
	{
		assignedAddressKana := false
		hadRawAddressKana := false
		if rawValueAddressKana, rawOk := plainValueAtPath(raw, "address_kana"); rawOk {
			hadRawAddressKana = true
			if rawValueAddressKana != nil {
				sourceAddressKana := applyConfiguredKeyedListShapes(rawValueAddressKana, attrValueToPlain(state.AddressKana))
				if valueAddressKana, err := flattenPlainValue(sourceAddressKana, types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}}, "address_kana", "raw response"); err != nil {
					return err
				} else {
					if typedAddressKana, ok := valueAddressKana.(types.Object); ok {
						state.AddressKana = typedAddressKana
						assignedAddressKana = true
					}
				}
			}
		}
		if !assignedAddressKana {
			if !hasRaw {
				if responseValueAddressKana, ok := plainFromResponseField(obj, "AddressKana"); ok {
					sourceAddressKana := applyConfiguredKeyedListShapes(responseValueAddressKana, attrValueToPlain(state.AddressKana))
					if valueAddressKana, err := flattenPlainValue(
						sourceAddressKana,
						types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}},
						"address_kana",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAddressKana, ok := valueAddressKana.(types.Object); ok {
							state.AddressKana = typedAddressKana
							assignedAddressKana = true
						}
					}
				}
			}
		}
		if !assignedAddressKana && hadRawAddressKana {
			if nullAddressKana, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}}); ok {
				if typedAddressKana, ok := nullAddressKana.(types.Object); ok {
					state.AddressKana = typedAddressKana
				}
			}
		}
	}
	{
		assignedAddressKanji := false
		hadRawAddressKanji := false
		if rawValueAddressKanji, rawOk := plainValueAtPath(raw, "address_kanji"); rawOk {
			hadRawAddressKanji = true
			if rawValueAddressKanji != nil {
				sourceAddressKanji := applyConfiguredKeyedListShapes(rawValueAddressKanji, attrValueToPlain(state.AddressKanji))
				if valueAddressKanji, err := flattenPlainValue(sourceAddressKanji, types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}}, "address_kanji", "raw response"); err != nil {
					return err
				} else {
					if typedAddressKanji, ok := valueAddressKanji.(types.Object); ok {
						state.AddressKanji = typedAddressKanji
						assignedAddressKanji = true
					}
				}
			}
		}
		if !assignedAddressKanji {
			if !hasRaw {
				if responseValueAddressKanji, ok := plainFromResponseField(obj, "AddressKanji"); ok {
					sourceAddressKanji := applyConfiguredKeyedListShapes(responseValueAddressKanji, attrValueToPlain(state.AddressKanji))
					if valueAddressKanji, err := flattenPlainValue(
						sourceAddressKanji,
						types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}},
						"address_kanji",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedAddressKanji, ok := valueAddressKanji.(types.Object); ok {
							state.AddressKanji = typedAddressKanji
							assignedAddressKanji = true
						}
					}
				}
			}
		}
		if !assignedAddressKanji && hadRawAddressKanji {
			if nullAddressKanji, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType, "town": types.StringType}}); ok {
				if typedAddressKanji, ok := nullAddressKanji.(types.Object); ok {
					state.AddressKanji = typedAddressKanji
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
		assignedDOB := false
		hadRawDOB := false
		if rawValueDOB, rawOk := plainValueAtPath(raw, "dob"); rawOk {
			hadRawDOB = true
			if rawValueDOB != nil {
				sourceDOB := applyConfiguredKeyedListShapes(rawValueDOB, attrValueToPlain(state.DOB))
				if valueDOB, err := flattenPlainValue(sourceDOB, types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}, "dob", "raw response"); err != nil {
					return err
				} else {
					if typedDOB, ok := valueDOB.(types.Object); ok {
						state.DOB = typedDOB
						assignedDOB = true
					}
				}
			}
		}
		if !assignedDOB {
			if !hasRaw {
				if responseValueDOB, ok := plainFromResponseField(obj, "DOB"); ok {
					sourceDOB := applyConfiguredKeyedListShapes(responseValueDOB, attrValueToPlain(state.DOB))
					if valueDOB, err := flattenPlainValue(
						sourceDOB,
						types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}},
						"dob",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedDOB, ok := valueDOB.(types.Object); ok {
							state.DOB = typedDOB
							assignedDOB = true
						}
					}
				}
			}
		}
		if !assignedDOB && hadRawDOB {
			if nullDOB, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}); ok {
				if typedDOB, ok := nullDOB.(types.Object); ok {
					state.DOB = typedDOB
				}
			}
		}
	}
	{
		if rawValueEmail, rawOk := plainValueAtPath(raw, "email"); rawOk {
			if valueEmail, err := flattenPlainValue(rawValueEmail, types.StringType, "email", "raw response"); err != nil {
				return err
			} else {
				if typedEmail, ok := valueEmail.(types.String); ok {
					state.Email = typedEmail
				}
			}
		} else if !hasRaw {
			if responseValueEmail, ok := plainFromResponseField(obj, "Email"); ok {
				if valueEmail, err := flattenPlainValue(responseValueEmail, types.StringType, "email", "response struct"); err != nil {
					return err
				} else {
					if typedEmail, ok := valueEmail.(types.String); ok {
						state.Email = typedEmail
					}
				}
			}
		}
	}
	{
		if rawValueFirstName, rawOk := plainValueAtPath(raw, "first_name"); rawOk {
			if valueFirstName, err := flattenPlainValue(rawValueFirstName, types.StringType, "first_name", "raw response"); err != nil {
				return err
			} else {
				if typedFirstName, ok := valueFirstName.(types.String); ok {
					state.FirstName = typedFirstName
				}
			}
		} else if !hasRaw {
			if responseValueFirstName, ok := plainFromResponseField(obj, "FirstName"); ok {
				if valueFirstName, err := flattenPlainValue(responseValueFirstName, types.StringType, "first_name", "response struct"); err != nil {
					return err
				} else {
					if typedFirstName, ok := valueFirstName.(types.String); ok {
						state.FirstName = typedFirstName
					}
				}
			}
		}
	}
	{
		if rawValueFirstNameKana, rawOk := plainValueAtPath(raw, "first_name_kana"); rawOk {
			if valueFirstNameKana, err := flattenPlainValue(rawValueFirstNameKana, types.StringType, "first_name_kana", "raw response"); err != nil {
				return err
			} else {
				if typedFirstNameKana, ok := valueFirstNameKana.(types.String); ok {
					state.FirstNameKana = typedFirstNameKana
				}
			}
		} else if !hasRaw {
			if responseValueFirstNameKana, ok := plainFromResponseField(obj, "FirstNameKana"); ok {
				if valueFirstNameKana, err := flattenPlainValue(responseValueFirstNameKana, types.StringType, "first_name_kana", "response struct"); err != nil {
					return err
				} else {
					if typedFirstNameKana, ok := valueFirstNameKana.(types.String); ok {
						state.FirstNameKana = typedFirstNameKana
					}
				}
			}
		}
	}
	{
		if rawValueFirstNameKanji, rawOk := plainValueAtPath(raw, "first_name_kanji"); rawOk {
			if valueFirstNameKanji, err := flattenPlainValue(rawValueFirstNameKanji, types.StringType, "first_name_kanji", "raw response"); err != nil {
				return err
			} else {
				if typedFirstNameKanji, ok := valueFirstNameKanji.(types.String); ok {
					state.FirstNameKanji = typedFirstNameKanji
				}
			}
		} else if !hasRaw {
			if responseValueFirstNameKanji, ok := plainFromResponseField(obj, "FirstNameKanji"); ok {
				if valueFirstNameKanji, err := flattenPlainValue(responseValueFirstNameKanji, types.StringType, "first_name_kanji", "response struct"); err != nil {
					return err
				} else {
					if typedFirstNameKanji, ok := valueFirstNameKanji.(types.String); ok {
						state.FirstNameKanji = typedFirstNameKanji
					}
				}
			}
		}
	}
	{
		if rawValueFullNameAliases, rawOk := plainValueAtPath(raw, "full_name_aliases"); rawOk {
			if valueFullNameAliases, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueFullNameAliases, attrValueToPlain(state.FullNameAliases)), types.ListType{ElemType: types.StringType}, "full_name_aliases", "raw response"); err != nil {
				return err
			} else {
				if typedFullNameAliases, ok := valueFullNameAliases.(types.List); ok {
					state.FullNameAliases = typedFullNameAliases
				}
			}
		} else if !hasRaw {
			if responseValueFullNameAliases, ok := plainFromResponseField(obj, "FullNameAliases"); ok {
				if valueFullNameAliases, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueFullNameAliases, attrValueToPlain(state.FullNameAliases)),
					types.ListType{ElemType: types.StringType},
					"full_name_aliases",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedFullNameAliases, ok := valueFullNameAliases.(types.List); ok {
						state.FullNameAliases = typedFullNameAliases
					}
				}
			}
		}
	}
	{
		assignedFutureRequirements := false
		hadRawFutureRequirements := false
		if rawValueFutureRequirements, rawOk := plainValueAtPath(raw, "future_requirements"); rawOk {
			hadRawFutureRequirements = true
			if rawValueFutureRequirements != nil {
				sourceFutureRequirements := applyConfiguredKeyedListShapes(rawValueFutureRequirements, attrValueToPlain(state.FutureRequirements))
				if valueFutureRequirements, err := flattenPlainValue(sourceFutureRequirements, types.ObjectType{AttrTypes: map[string]attr.Type{"alternatives": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"alternative_fields_due": types.ListType{ElemType: types.StringType}, "original_fields_due": types.ListType{ElemType: types.StringType}}}}, "currently_due": types.ListType{ElemType: types.StringType}, "errors": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "reason": types.StringType, "requirement": types.StringType}}}, "eventually_due": types.ListType{ElemType: types.StringType}, "past_due": types.ListType{ElemType: types.StringType}, "pending_verification": types.ListType{ElemType: types.StringType}}}, "future_requirements", "raw response"); err != nil {
					return err
				} else {
					if typedFutureRequirements, ok := valueFutureRequirements.(types.Object); ok {
						state.FutureRequirements = typedFutureRequirements
						assignedFutureRequirements = true
					}
				}
			}
		}
		if !assignedFutureRequirements {
			if !hasRaw {
				if responseValueFutureRequirements, ok := plainFromResponseField(obj, "FutureRequirements"); ok {
					sourceFutureRequirements := applyConfiguredKeyedListShapes(responseValueFutureRequirements, attrValueToPlain(state.FutureRequirements))
					if valueFutureRequirements, err := flattenPlainValue(
						sourceFutureRequirements,
						types.ObjectType{AttrTypes: map[string]attr.Type{"alternatives": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"alternative_fields_due": types.ListType{ElemType: types.StringType}, "original_fields_due": types.ListType{ElemType: types.StringType}}}}, "currently_due": types.ListType{ElemType: types.StringType}, "errors": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "reason": types.StringType, "requirement": types.StringType}}}, "eventually_due": types.ListType{ElemType: types.StringType}, "past_due": types.ListType{ElemType: types.StringType}, "pending_verification": types.ListType{ElemType: types.StringType}}},
						"future_requirements",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedFutureRequirements, ok := valueFutureRequirements.(types.Object); ok {
							state.FutureRequirements = typedFutureRequirements
							assignedFutureRequirements = true
						}
					}
				}
			}
		}
		if !assignedFutureRequirements && hadRawFutureRequirements {
			if nullFutureRequirements, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"alternatives": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"alternative_fields_due": types.ListType{ElemType: types.StringType}, "original_fields_due": types.ListType{ElemType: types.StringType}}}}, "currently_due": types.ListType{ElemType: types.StringType}, "errors": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "reason": types.StringType, "requirement": types.StringType}}}, "eventually_due": types.ListType{ElemType: types.StringType}, "past_due": types.ListType{ElemType: types.StringType}, "pending_verification": types.ListType{ElemType: types.StringType}}}); ok {
				if typedFutureRequirements, ok := nullFutureRequirements.(types.Object); ok {
					state.FutureRequirements = typedFutureRequirements
				}
			}
		}
	}
	{
		if rawValueGender, rawOk := plainValueAtPath(raw, "gender"); rawOk {
			if valueGender, err := flattenPlainValue(rawValueGender, types.StringType, "gender", "raw response"); err != nil {
				return err
			} else {
				if typedGender, ok := valueGender.(types.String); ok {
					state.Gender = typedGender
				}
			}
		} else if !hasRaw {
			if responseValueGender, ok := plainFromResponseField(obj, "Gender"); ok {
				if valueGender, err := flattenPlainValue(responseValueGender, types.StringType, "gender", "response struct"); err != nil {
					return err
				} else {
					if typedGender, ok := valueGender.(types.String); ok {
						state.Gender = typedGender
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
		if rawValueIDNumberProvided, rawOk := plainValueAtPath(raw, "id_number_provided"); rawOk {
			if valueIDNumberProvided, err := flattenPlainValue(rawValueIDNumberProvided, types.BoolType, "id_number_provided", "raw response"); err != nil {
				return err
			} else {
				if typedIDNumberProvided, ok := valueIDNumberProvided.(types.Bool); ok {
					state.IDNumberProvided = typedIDNumberProvided
				}
			}
		} else if !hasRaw {
			if responseValueIDNumberProvided, ok := plainFromResponseField(obj, "IDNumberProvided"); ok {
				if valueIDNumberProvided, err := flattenPlainValue(responseValueIDNumberProvided, types.BoolType, "id_number_provided", "response struct"); err != nil {
					return err
				} else {
					if typedIDNumberProvided, ok := valueIDNumberProvided.(types.Bool); ok {
						state.IDNumberProvided = typedIDNumberProvided
					}
				}
			}
		}
	}
	{
		if rawValueIDNumberSecondaryProvided, rawOk := plainValueAtPath(raw, "id_number_secondary_provided"); rawOk {
			if valueIDNumberSecondaryProvided, err := flattenPlainValue(rawValueIDNumberSecondaryProvided, types.BoolType, "id_number_secondary_provided", "raw response"); err != nil {
				return err
			} else {
				if typedIDNumberSecondaryProvided, ok := valueIDNumberSecondaryProvided.(types.Bool); ok {
					state.IDNumberSecondaryProvided = typedIDNumberSecondaryProvided
				}
			}
		} else if !hasRaw {
			if responseValueIDNumberSecondaryProvided, ok := plainFromResponseField(obj, "IDNumberSecondaryProvided"); ok {
				if valueIDNumberSecondaryProvided, err := flattenPlainValue(responseValueIDNumberSecondaryProvided, types.BoolType, "id_number_secondary_provided", "response struct"); err != nil {
					return err
				} else {
					if typedIDNumberSecondaryProvided, ok := valueIDNumberSecondaryProvided.(types.Bool); ok {
						state.IDNumberSecondaryProvided = typedIDNumberSecondaryProvided
					}
				}
			}
		}
	}
	{
		if rawValueLastName, rawOk := plainValueAtPath(raw, "last_name"); rawOk {
			if valueLastName, err := flattenPlainValue(rawValueLastName, types.StringType, "last_name", "raw response"); err != nil {
				return err
			} else {
				if typedLastName, ok := valueLastName.(types.String); ok {
					state.LastName = typedLastName
				}
			}
		} else if !hasRaw {
			if responseValueLastName, ok := plainFromResponseField(obj, "LastName"); ok {
				if valueLastName, err := flattenPlainValue(responseValueLastName, types.StringType, "last_name", "response struct"); err != nil {
					return err
				} else {
					if typedLastName, ok := valueLastName.(types.String); ok {
						state.LastName = typedLastName
					}
				}
			}
		}
	}
	{
		if rawValueLastNameKana, rawOk := plainValueAtPath(raw, "last_name_kana"); rawOk {
			if valueLastNameKana, err := flattenPlainValue(rawValueLastNameKana, types.StringType, "last_name_kana", "raw response"); err != nil {
				return err
			} else {
				if typedLastNameKana, ok := valueLastNameKana.(types.String); ok {
					state.LastNameKana = typedLastNameKana
				}
			}
		} else if !hasRaw {
			if responseValueLastNameKana, ok := plainFromResponseField(obj, "LastNameKana"); ok {
				if valueLastNameKana, err := flattenPlainValue(responseValueLastNameKana, types.StringType, "last_name_kana", "response struct"); err != nil {
					return err
				} else {
					if typedLastNameKana, ok := valueLastNameKana.(types.String); ok {
						state.LastNameKana = typedLastNameKana
					}
				}
			}
		}
	}
	{
		if rawValueLastNameKanji, rawOk := plainValueAtPath(raw, "last_name_kanji"); rawOk {
			if valueLastNameKanji, err := flattenPlainValue(rawValueLastNameKanji, types.StringType, "last_name_kanji", "raw response"); err != nil {
				return err
			} else {
				if typedLastNameKanji, ok := valueLastNameKanji.(types.String); ok {
					state.LastNameKanji = typedLastNameKanji
				}
			}
		} else if !hasRaw {
			if responseValueLastNameKanji, ok := plainFromResponseField(obj, "LastNameKanji"); ok {
				if valueLastNameKanji, err := flattenPlainValue(responseValueLastNameKanji, types.StringType, "last_name_kanji", "response struct"); err != nil {
					return err
				} else {
					if typedLastNameKanji, ok := valueLastNameKanji.(types.String); ok {
						state.LastNameKanji = typedLastNameKanji
					}
				}
			}
		}
	}
	{
		if rawValueMaidenName, rawOk := plainValueAtPath(raw, "maiden_name"); rawOk {
			if valueMaidenName, err := flattenPlainValue(rawValueMaidenName, types.StringType, "maiden_name", "raw response"); err != nil {
				return err
			} else {
				if typedMaidenName, ok := valueMaidenName.(types.String); ok {
					state.MaidenName = typedMaidenName
				}
			}
		} else if !hasRaw {
			if responseValueMaidenName, ok := plainFromResponseField(obj, "MaidenName"); ok {
				if valueMaidenName, err := flattenPlainValue(responseValueMaidenName, types.StringType, "maiden_name", "response struct"); err != nil {
					return err
				} else {
					if typedMaidenName, ok := valueMaidenName.(types.String); ok {
						state.MaidenName = typedMaidenName
					}
				}
			}
		}
	}
	{
		if rawValueMetadata, rawOk := plainValueAtPath(raw, "metadata"); rawOk {
			if valueMetadata, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueMetadata, attrValueToPlain(state.Metadata)), types.MapType{ElemType: types.StringType}, "metadata", "raw response"); err != nil {
				return err
			} else {
				if typedMetadata, ok := valueMetadata.(types.Map); ok {
					state.Metadata = typedMetadata
				}
			}
		} else if !hasRaw {
			if responseValueMetadata, ok := plainFromResponseField(obj, "Metadata"); ok {
				if valueMetadata, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueMetadata, attrValueToPlain(state.Metadata)),
					types.MapType{ElemType: types.StringType},
					"metadata",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedMetadata, ok := valueMetadata.(types.Map); ok {
						state.Metadata = typedMetadata
					}
				}
			}
		}
	}
	{
		if rawValueNationality, rawOk := plainValueAtPath(raw, "nationality"); rawOk {
			if valueNationality, err := flattenPlainValue(rawValueNationality, types.StringType, "nationality", "raw response"); err != nil {
				return err
			} else {
				if typedNationality, ok := valueNationality.(types.String); ok {
					state.Nationality = typedNationality
				}
			}
		} else if !hasRaw {
			if responseValueNationality, ok := plainFromResponseField(obj, "Nationality"); ok {
				if valueNationality, err := flattenPlainValue(responseValueNationality, types.StringType, "nationality", "response struct"); err != nil {
					return err
				} else {
					if typedNationality, ok := valueNationality.(types.String); ok {
						state.Nationality = typedNationality
					}
				}
			}
		}
	}
	{
		if rawValuePhone, rawOk := plainValueAtPath(raw, "phone"); rawOk {
			if valuePhone, err := flattenPlainValue(rawValuePhone, types.StringType, "phone", "raw response"); err != nil {
				return err
			} else {
				if typedPhone, ok := valuePhone.(types.String); ok {
					state.Phone = typedPhone
				}
			}
		} else if !hasRaw {
			if responseValuePhone, ok := plainFromResponseField(obj, "Phone"); ok {
				if valuePhone, err := flattenPlainValue(responseValuePhone, types.StringType, "phone", "response struct"); err != nil {
					return err
				} else {
					if typedPhone, ok := valuePhone.(types.String); ok {
						state.Phone = typedPhone
					}
				}
			}
		}
	}
	{
		if rawValuePoliticalExposure, rawOk := plainValueAtPath(raw, "political_exposure"); rawOk {
			if valuePoliticalExposure, err := flattenPlainValue(rawValuePoliticalExposure, types.StringType, "political_exposure", "raw response"); err != nil {
				return err
			} else {
				if typedPoliticalExposure, ok := valuePoliticalExposure.(types.String); ok {
					state.PoliticalExposure = typedPoliticalExposure
				}
			}
		} else if !hasRaw {
			if responseValuePoliticalExposure, ok := plainFromResponseField(obj, "PoliticalExposure"); ok {
				if valuePoliticalExposure, err := flattenPlainValue(responseValuePoliticalExposure, types.StringType, "political_exposure", "response struct"); err != nil {
					return err
				} else {
					if typedPoliticalExposure, ok := valuePoliticalExposure.(types.String); ok {
						state.PoliticalExposure = typedPoliticalExposure
					}
				}
			}
		}
	}
	{
		assignedRegisteredAddress := false
		hadRawRegisteredAddress := false
		if rawValueRegisteredAddress, rawOk := plainValueAtPath(raw, "registered_address"); rawOk {
			hadRawRegisteredAddress = true
			if rawValueRegisteredAddress != nil {
				sourceRegisteredAddress := applyConfiguredKeyedListShapes(rawValueRegisteredAddress, attrValueToPlain(state.RegisteredAddress))
				if valueRegisteredAddress, err := flattenPlainValue(sourceRegisteredAddress, types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "registered_address", "raw response"); err != nil {
					return err
				} else {
					if typedRegisteredAddress, ok := valueRegisteredAddress.(types.Object); ok {
						state.RegisteredAddress = typedRegisteredAddress
						assignedRegisteredAddress = true
					}
				}
			}
		}
		if !assignedRegisteredAddress {
			if !hasRaw {
				if responseValueRegisteredAddress, ok := plainFromResponseField(obj, "RegisteredAddress"); ok {
					sourceRegisteredAddress := applyConfiguredKeyedListShapes(responseValueRegisteredAddress, attrValueToPlain(state.RegisteredAddress))
					if valueRegisteredAddress, err := flattenPlainValue(
						sourceRegisteredAddress,
						types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}},
						"registered_address",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRegisteredAddress, ok := valueRegisteredAddress.(types.Object); ok {
							state.RegisteredAddress = typedRegisteredAddress
							assignedRegisteredAddress = true
						}
					}
				}
			}
		}
		if !assignedRegisteredAddress && hadRawRegisteredAddress {
			if nullRegisteredAddress, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}); ok {
				if typedRegisteredAddress, ok := nullRegisteredAddress.(types.Object); ok {
					state.RegisteredAddress = typedRegisteredAddress
				}
			}
		}
	}
	{
		assignedRelationship := false
		hadRawRelationship := false
		if rawValueRelationship, rawOk := plainValueAtPath(raw, "relationship"); rawOk {
			hadRawRelationship = true
			if rawValueRelationship != nil {
				sourceRelationship := applyConfiguredKeyedListShapes(rawValueRelationship, attrValueToPlain(state.Relationship))
				if valueRelationship, err := flattenPlainValue(sourceRelationship, types.ObjectType{AttrTypes: map[string]attr.Type{"authorizer": types.BoolType, "director": types.BoolType, "executive": types.BoolType, "legal_guardian": types.BoolType, "owner": types.BoolType, "percent_ownership": types.Float64Type, "representative": types.BoolType, "title": types.StringType}}, "relationship", "raw response"); err != nil {
					return err
				} else {
					if typedRelationship, ok := valueRelationship.(types.Object); ok {
						state.Relationship = typedRelationship
						assignedRelationship = true
					}
				}
			}
		}
		if !assignedRelationship {
			if !hasRaw {
				if responseValueRelationship, ok := plainFromResponseField(obj, "Relationship"); ok {
					sourceRelationship := applyConfiguredKeyedListShapes(responseValueRelationship, attrValueToPlain(state.Relationship))
					if valueRelationship, err := flattenPlainValue(
						sourceRelationship,
						types.ObjectType{AttrTypes: map[string]attr.Type{"authorizer": types.BoolType, "director": types.BoolType, "executive": types.BoolType, "legal_guardian": types.BoolType, "owner": types.BoolType, "percent_ownership": types.Float64Type, "representative": types.BoolType, "title": types.StringType}},
						"relationship",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRelationship, ok := valueRelationship.(types.Object); ok {
							state.Relationship = typedRelationship
							assignedRelationship = true
						}
					}
				}
			}
		}
		if !assignedRelationship && hadRawRelationship {
			if nullRelationship, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"authorizer": types.BoolType, "director": types.BoolType, "executive": types.BoolType, "legal_guardian": types.BoolType, "owner": types.BoolType, "percent_ownership": types.Float64Type, "representative": types.BoolType, "title": types.StringType}}); ok {
				if typedRelationship, ok := nullRelationship.(types.Object); ok {
					state.Relationship = typedRelationship
				}
			}
		}
	}
	{
		assignedRequirements := false
		hadRawRequirements := false
		if rawValueRequirements, rawOk := plainValueAtPath(raw, "requirements"); rawOk {
			hadRawRequirements = true
			if rawValueRequirements != nil {
				sourceRequirements := applyConfiguredKeyedListShapes(rawValueRequirements, attrValueToPlain(state.Requirements))
				if valueRequirements, err := flattenPlainValue(sourceRequirements, types.ObjectType{AttrTypes: map[string]attr.Type{"alternatives": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"alternative_fields_due": types.ListType{ElemType: types.StringType}, "original_fields_due": types.ListType{ElemType: types.StringType}}}}, "currently_due": types.ListType{ElemType: types.StringType}, "errors": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "reason": types.StringType, "requirement": types.StringType}}}, "eventually_due": types.ListType{ElemType: types.StringType}, "past_due": types.ListType{ElemType: types.StringType}, "pending_verification": types.ListType{ElemType: types.StringType}}}, "requirements", "raw response"); err != nil {
					return err
				} else {
					if typedRequirements, ok := valueRequirements.(types.Object); ok {
						state.Requirements = typedRequirements
						assignedRequirements = true
					}
				}
			}
		}
		if !assignedRequirements {
			if !hasRaw {
				if responseValueRequirements, ok := plainFromResponseField(obj, "Requirements"); ok {
					sourceRequirements := applyConfiguredKeyedListShapes(responseValueRequirements, attrValueToPlain(state.Requirements))
					if valueRequirements, err := flattenPlainValue(
						sourceRequirements,
						types.ObjectType{AttrTypes: map[string]attr.Type{"alternatives": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"alternative_fields_due": types.ListType{ElemType: types.StringType}, "original_fields_due": types.ListType{ElemType: types.StringType}}}}, "currently_due": types.ListType{ElemType: types.StringType}, "errors": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "reason": types.StringType, "requirement": types.StringType}}}, "eventually_due": types.ListType{ElemType: types.StringType}, "past_due": types.ListType{ElemType: types.StringType}, "pending_verification": types.ListType{ElemType: types.StringType}}},
						"requirements",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRequirements, ok := valueRequirements.(types.Object); ok {
							state.Requirements = typedRequirements
							assignedRequirements = true
						}
					}
				}
			}
		}
		if !assignedRequirements && hadRawRequirements {
			if nullRequirements, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"alternatives": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"alternative_fields_due": types.ListType{ElemType: types.StringType}, "original_fields_due": types.ListType{ElemType: types.StringType}}}}, "currently_due": types.ListType{ElemType: types.StringType}, "errors": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "reason": types.StringType, "requirement": types.StringType}}}, "eventually_due": types.ListType{ElemType: types.StringType}, "past_due": types.ListType{ElemType: types.StringType}, "pending_verification": types.ListType{ElemType: types.StringType}}}); ok {
				if typedRequirements, ok := nullRequirements.(types.Object); ok {
					state.Requirements = typedRequirements
				}
			}
		}
	}
	{
		if rawValueSSNLast4Provided, rawOk := plainValueAtPath(raw, "ssn_last_4_provided"); rawOk {
			if valueSSNLast4Provided, err := flattenPlainValue(rawValueSSNLast4Provided, types.BoolType, "ssn_last_4_provided", "raw response"); err != nil {
				return err
			} else {
				if typedSSNLast4Provided, ok := valueSSNLast4Provided.(types.Bool); ok {
					state.SSNLast4Provided = typedSSNLast4Provided
				}
			}
		} else if !hasRaw {
			if responseValueSSNLast4Provided, ok := plainFromResponseField(obj, "SSNLast4Provided"); ok {
				if valueSSNLast4Provided, err := flattenPlainValue(responseValueSSNLast4Provided, types.BoolType, "ssn_last_4_provided", "response struct"); err != nil {
					return err
				} else {
					if typedSSNLast4Provided, ok := valueSSNLast4Provided.(types.Bool); ok {
						state.SSNLast4Provided = typedSSNLast4Provided
					}
				}
			}
		}
	}
	{
		assignedUSCfpbData := false
		hadRawUSCfpbData := false
		if rawValueUSCfpbData, rawOk := plainValueAtPath(raw, "us_cfpb_data"); rawOk {
			hadRawUSCfpbData = true
			if rawValueUSCfpbData != nil {
				sourceUSCfpbData := applyConfiguredKeyedListShapes(rawValueUSCfpbData, attrValueToPlain(state.USCfpbData))
				if valueUSCfpbData, err := flattenPlainValue(sourceUSCfpbData, types.ObjectType{AttrTypes: map[string]attr.Type{"ethnicity_details": types.ObjectType{AttrTypes: map[string]attr.Type{"ethnicity": types.ListType{ElemType: types.StringType}, "ethnicity_other": types.StringType}}, "race_details": types.ObjectType{AttrTypes: map[string]attr.Type{"race": types.ListType{ElemType: types.StringType}, "race_other": types.StringType}}, "self_identified_gender": types.StringType}}, "us_cfpb_data", "raw response"); err != nil {
					return err
				} else {
					if typedUSCfpbData, ok := valueUSCfpbData.(types.Object); ok {
						state.USCfpbData = typedUSCfpbData
						assignedUSCfpbData = true
					}
				}
			}
		}
		if !assignedUSCfpbData {
			if !hasRaw {
				if responseValueUSCfpbData, ok := plainFromResponseField(obj, "USCfpbData"); ok {
					sourceUSCfpbData := applyConfiguredKeyedListShapes(responseValueUSCfpbData, attrValueToPlain(state.USCfpbData))
					if valueUSCfpbData, err := flattenPlainValue(
						sourceUSCfpbData,
						types.ObjectType{AttrTypes: map[string]attr.Type{"ethnicity_details": types.ObjectType{AttrTypes: map[string]attr.Type{"ethnicity": types.ListType{ElemType: types.StringType}, "ethnicity_other": types.StringType}}, "race_details": types.ObjectType{AttrTypes: map[string]attr.Type{"race": types.ListType{ElemType: types.StringType}, "race_other": types.StringType}}, "self_identified_gender": types.StringType}},
						"us_cfpb_data",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedUSCfpbData, ok := valueUSCfpbData.(types.Object); ok {
							state.USCfpbData = typedUSCfpbData
							assignedUSCfpbData = true
						}
					}
				}
			}
		}
		if !assignedUSCfpbData && hadRawUSCfpbData {
			if nullUSCfpbData, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"ethnicity_details": types.ObjectType{AttrTypes: map[string]attr.Type{"ethnicity": types.ListType{ElemType: types.StringType}, "ethnicity_other": types.StringType}}, "race_details": types.ObjectType{AttrTypes: map[string]attr.Type{"race": types.ListType{ElemType: types.StringType}, "race_other": types.StringType}}, "self_identified_gender": types.StringType}}); ok {
				if typedUSCfpbData, ok := nullUSCfpbData.(types.Object); ok {
					state.USCfpbData = typedUSCfpbData
				}
			}
		}
	}
	{
		assignedVerification := false
		hadRawVerification := false
		if rawValueVerification, rawOk := plainValueAtPath(raw, "verification"); rawOk {
			hadRawVerification = true
			if rawValueVerification != nil {
				sourceVerification := applyConfiguredKeyedListShapes(rawValueVerification, attrValueToPlain(state.Verification))
				if valueVerification, err := flattenPlainValue(sourceVerification, types.ObjectType{AttrTypes: map[string]attr.Type{"additional_document": types.ObjectType{AttrTypes: map[string]attr.Type{"back": types.StringType, "details": types.StringType, "details_code": types.StringType, "front": types.StringType}}, "details": types.StringType, "details_code": types.StringType, "document": types.ObjectType{AttrTypes: map[string]attr.Type{"back": types.StringType, "details": types.StringType, "details_code": types.StringType, "front": types.StringType}}, "status": types.StringType}}, "verification", "raw response"); err != nil {
					return err
				} else {
					if typedVerification, ok := valueVerification.(types.Object); ok {
						state.Verification = typedVerification
						assignedVerification = true
					}
				}
			}
		}
		if !assignedVerification {
			if !hasRaw {
				if responseValueVerification, ok := plainFromResponseField(obj, "Verification"); ok {
					sourceVerification := applyConfiguredKeyedListShapes(responseValueVerification, attrValueToPlain(state.Verification))
					if valueVerification, err := flattenPlainValue(
						sourceVerification,
						types.ObjectType{AttrTypes: map[string]attr.Type{"additional_document": types.ObjectType{AttrTypes: map[string]attr.Type{"back": types.StringType, "details": types.StringType, "details_code": types.StringType, "front": types.StringType}}, "details": types.StringType, "details_code": types.StringType, "document": types.ObjectType{AttrTypes: map[string]attr.Type{"back": types.StringType, "details": types.StringType, "details_code": types.StringType, "front": types.StringType}}, "status": types.StringType}},
						"verification",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedVerification, ok := valueVerification.(types.Object); ok {
							state.Verification = typedVerification
							assignedVerification = true
						}
					}
				}
			}
		}
		if !assignedVerification && hadRawVerification {
			if nullVerification, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"additional_document": types.ObjectType{AttrTypes: map[string]attr.Type{"back": types.StringType, "details": types.StringType, "details_code": types.StringType, "front": types.StringType}}, "details": types.StringType, "details_code": types.StringType, "document": types.ObjectType{AttrTypes: map[string]attr.Type{"back": types.StringType, "details": types.StringType, "details_code": types.StringType, "front": types.StringType}}, "status": types.StringType}}); ok {
				if typedVerification, ok := nullVerification.(types.Object); ok {
					state.Verification = typedVerification
				}
			}
		}
	}
	return nil
}
