//
// File generated from our OpenAPI spec
//

package ephemeralresources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	ephemeralSchema "github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ ephemeral.EphemeralResource = &IdentityVerificationSessionEphemeralResource{}
var _ ephemeral.EphemeralResourceWithConfigure = &IdentityVerificationSessionEphemeralResource{}

func NewIdentityVerificationSessionEphemeralResource() ephemeral.EphemeralResource {
	return &IdentityVerificationSessionEphemeralResource{}
}

type IdentityVerificationSessionEphemeralResource struct {
	client *stripe.Client
}

type IdentityVerificationSessionResourceModel struct {
	Object                 types.String `tfsdk:"object"`
	ClientReferenceID      types.String `tfsdk:"client_reference_id"`
	ClientSecret           types.String `tfsdk:"client_secret"`
	Created                types.Int64  `tfsdk:"created"`
	ID                     types.String `tfsdk:"id"`
	LastError              types.Object `tfsdk:"last_error"`
	LastVerificationReport types.String `tfsdk:"last_verification_report"`
	Livemode               types.Bool   `tfsdk:"livemode"`
	Metadata               types.Map    `tfsdk:"metadata"`
	Options                types.Object `tfsdk:"options"`
	ProvidedDetails        types.Object `tfsdk:"provided_details"`
	Redaction              types.Object `tfsdk:"redaction"`
	RelatedCustomer        types.String `tfsdk:"related_customer"`
	RelatedCustomerAccount types.String `tfsdk:"related_customer_account"`
	RelatedPerson          types.Object `tfsdk:"related_person"`
	Status                 types.String `tfsdk:"status"`
	Type                   types.String `tfsdk:"type"`
	URL                    types.String `tfsdk:"url"`
	VerificationFlow       types.String `tfsdk:"verification_flow"`
	VerifiedOutputs        types.Object `tfsdk:"verified_outputs"`
	ReturnURL              types.String `tfsdk:"return_url"`
}

func (r *IdentityVerificationSessionEphemeralResource) Metadata(_ context.Context, req ephemeral.MetadataRequest, resp *ephemeral.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_identity_verification_session"
}

func (r *IdentityVerificationSessionEphemeralResource) Schema(_ context.Context, _ ephemeral.SchemaRequest, resp *ephemeral.SchemaResponse) {
	resp.Schema = ephemeralSchema.Schema{
		Description: "A VerificationSession guides you through the process of collecting and verifying the identities\nof your users. It contains details about the type of verification, such as what [verification\ncheck](/docs/identity/verification-checks) to perform. Only create one VerificationSession for\neach verification in your system.\n\nA VerificationSession transitions through [multiple\nstatuses](/docs/identity/how-sessions-work) throughout its lifetime as it progresses through\nthe verification flow. The VerificationSession contains the user's verified data after\nverification checks are complete.\n\nRelated guide: [The Verification Sessions API](https://docs.stripe.com/identity/verification-sessions)",
		Attributes: map[string]ephemeralSchema.Attribute{
			"object": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "String representing the object's type. Objects of the same type share the same value.",
				Validators:  []validator.String{stringvalidator.OneOf("identity.verification_session")},
			},
			"client_reference_id": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A string to reference this user. This can be a customer ID, a session ID, or similar, and can be used to reconcile this verification with your internal systems.",
			},
			"client_secret": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The short-lived client secret used by Stripe.js to [show a verification modal](https://docs.stripe.com/js/identity/modal) inside your app. This client secret expires after 24 hours and can only be used once. Don’t store it, log it, embed it in a URL, or expose it to anyone other than the user. Make sure that you have TLS enabled on any page that includes the client secret. Refer to our docs on [passing the client secret to the frontend](https://docs.stripe.com/identity/verification-sessions#client-secret) to learn more.",
				Sensitive:   true,
			},
			"created": ephemeralSchema.Int64Attribute{
				Computed:    true,
				Description: "Time at which the object was created. Measured in seconds since the Unix epoch.",
			},
			"id": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier for the object.",
			},
			"last_error": ephemeralSchema.SingleNestedAttribute{
				Computed:    true,
				Description: "If present, this property tells you the last error encountered when processing the verification.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"code": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "A short machine-readable string giving the reason for the verification or user-session failure.",
						Validators:  []validator.String{stringvalidator.OneOf("abandoned", "consent_declined", "country_not_supported", "device_not_supported", "document_expired", "document_type_not_supported", "document_unverified_other", "email_unverified_other", "email_verification_declined", "id_number_insufficient_document_data", "id_number_mismatch", "id_number_unverified_other", "phone_unverified_other", "phone_verification_declined", "selfie_document_missing_photo", "selfie_face_mismatch", "selfie_manipulated", "selfie_unverified_other", "under_supported_age")},
					},
					"reason": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "A message that explains the reason for verification or user-session failure.",
					},
				},
			},
			"last_verification_report": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "ID of the most recent VerificationReport. [Learn more about accessing detailed verification results.](https://docs.stripe.com/identity/verification-sessions#results)",
			},
			"livemode": ephemeralSchema.BoolAttribute{
				Computed:    true,
				Description: "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
			},
			"metadata": ephemeralSchema.MapAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				ElementType: types.StringType,
			},
			"options": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A set of options for the session’s verification checks.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"document": ephemeralSchema.SingleNestedAttribute{
						Optional: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"allowed_types": ephemeralSchema.ListAttribute{
								Optional:    true,
								Description: "Array of strings of allowed identity document types. If the provided identity document isn’t one of the allowed types, the verification check will fail with a document_type_not_allowed error code.",
								ElementType: types.StringType,
							},
							"require_id_number": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Collect an ID number and perform an [ID number check](https://docs.stripe.com/identity/verification-checks?type=id-number) with the document’s extracted name and date of birth.",
							},
							"require_live_capture": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Disable image uploads, identity document images have to be captured using the device’s camera.",
							},
							"require_matching_selfie": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Capture a face image and perform a [selfie check](https://docs.stripe.com/identity/verification-checks?type=selfie) comparing a photo ID and a picture of your user’s face. [Learn more](https://docs.stripe.com/identity/selfie).",
							},
						},
					},
					"email": ephemeralSchema.SingleNestedAttribute{
						Optional: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"require_verification": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Request one time password verification of `provided_details.email`.",
							},
						},
					},
					"matching": ephemeralSchema.SingleNestedAttribute{
						Computed: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"dob": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "Strictness of the DOB matching policy to apply.",
								Validators:  []validator.String{stringvalidator.OneOf("none", "similar")},
							},
							"name": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "Strictness of the name matching policy to apply.",
								Validators:  []validator.String{stringvalidator.OneOf("none", "similar")},
							},
						},
					},
					"phone": ephemeralSchema.SingleNestedAttribute{
						Optional: true,

						Attributes: map[string]ephemeralSchema.Attribute{
							"require_verification": ephemeralSchema.BoolAttribute{
								Optional:    true,
								Description: "Request one time password verification of `provided_details.phone`.",
							},
						},
					},
				},
			},
			"provided_details": ephemeralSchema.SingleNestedAttribute{
				Optional:    true,
				Description: "Details provided about the user being verified. These details may be shown to the user.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"email": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Email of user being verified",
					},
					"phone": ephemeralSchema.StringAttribute{
						Optional:    true,
						Description: "Phone number of user being verified",
					},
				},
			},
			"redaction": ephemeralSchema.SingleNestedAttribute{
				Computed:    true,
				Description: "Redaction status of this VerificationSession. If the VerificationSession is not redacted, this field will be null.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"status": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "Indicates whether this object and its related objects have been redacted or not.",
						Validators:  []validator.String{stringvalidator.OneOf("processing", "redacted")},
					},
				},
			},
			"related_customer": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Customer ID",
			},
			"related_customer_account": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The ID of the Account representing a customer.",
			},
			"related_person": ephemeralSchema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				Attributes: map[string]ephemeralSchema.Attribute{
					"account": ephemeralSchema.StringAttribute{
						Required:    true,
						Description: "Token referencing the associated Account of the related Person resource.",
					},
					"person": ephemeralSchema.StringAttribute{
						Required:    true,
						Description: "Token referencing the related Person resource.",
					},
				},
			},
			"status": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "Status of this VerificationSession. [Learn more about the lifecycle of sessions](https://docs.stripe.com/identity/how-sessions-work).",
				Validators:  []validator.String{stringvalidator.OneOf("canceled", "processing", "requires_input", "verified")},
			},
			"type": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of [verification check](https://docs.stripe.com/identity/verification-checks) to be performed.",
				Validators:  []validator.String{stringvalidator.OneOf("document", "id_number", "verification_flow")},
			},
			"url": ephemeralSchema.StringAttribute{
				Computed:    true,
				Description: "The short-lived URL that you use to redirect a user to Stripe to submit their identity information. This URL expires after 48 hours and can only be used once. Don’t store it, log it, send it in emails or expose it to anyone other than the user. Refer to our docs on [verifying identity documents](https://docs.stripe.com/identity/verify-identity-documents?platform=web&type=redirect) to learn how to redirect users to Stripe.",
			},
			"verification_flow": ephemeralSchema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The configuration token of a verification flow from the dashboard.",
			},
			"verified_outputs": ephemeralSchema.SingleNestedAttribute{
				Computed:    true,
				Description: "The user’s verified data.",
				Attributes: map[string]ephemeralSchema.Attribute{
					"address": ephemeralSchema.SingleNestedAttribute{
						Computed:    true,
						Description: "The user's verified address.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"city": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "City, district, suburb, town, or village.",
							},
							"country": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).",
							},
							"line1": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "Address line 1, such as the street, PO Box, or company name.",
							},
							"line2": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "Address line 2, such as the apartment, suite, unit, or building.",
							},
							"postal_code": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "ZIP or postal code.",
							},
							"state": ephemeralSchema.StringAttribute{
								Computed:    true,
								Description: "State, county, province, or region ([ISO 3166-2](https://en.wikipedia.org/wiki/ISO_3166-2)).",
							},
						},
					},
					"dob": ephemeralSchema.SingleNestedAttribute{
						Computed:    true,
						Description: "The user’s verified date of birth.",
						Attributes: map[string]ephemeralSchema.Attribute{
							"day": ephemeralSchema.Int64Attribute{
								Computed:    true,
								Description: "Numerical day between 1 and 31.",
							},
							"month": ephemeralSchema.Int64Attribute{
								Computed:    true,
								Description: "Numerical month between 1 and 12.",
							},
							"year": ephemeralSchema.Int64Attribute{
								Computed:    true,
								Description: "The four-digit year.",
							},
						},
					},
					"email": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The user's verified email address",
					},
					"first_name": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The user's verified first name.",
					},
					"id_number": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The user's verified id number.",
					},
					"id_number_type": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The user's verified id number type.",
						Validators:  []validator.String{stringvalidator.OneOf("br_cpf", "sg_nric", "us_ssn")},
					},
					"last_name": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The user's verified last name.",
					},
					"phone": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The user's verified phone number",
					},
					"sex": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The user's verified sex.",
						Validators:  []validator.String{stringvalidator.OneOf("[redacted]", "female", "male", "unknown")},
					},
					"unparsed_place_of_birth": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The user's verified place of birth as it appears in the document.",
					},
					"unparsed_sex": ephemeralSchema.StringAttribute{
						Computed:    true,
						Description: "The user's verified sex as it appears in the document.",
					},
				},
			},
			"return_url": ephemeralSchema.StringAttribute{
				Optional:    true,
				Description: "The URL that the user will be redirected to upon completing the verification flow.",
			},
		},
	}
}

func (r *IdentityVerificationSessionEphemeralResource) Configure(_ context.Context, req ephemeral.ConfigureRequest, resp *ephemeral.ConfigureResponse) {
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

func expandIdentityVerificationSessionCreate(plan IdentityVerificationSessionResourceModel) (*stripe.IdentityVerificationSessionCreateParams, error) {
	params := &stripe.IdentityVerificationSessionCreateParams{}

	if !plan.ClientReferenceID.IsNull() && !plan.ClientReferenceID.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ClientReferenceID", "ClientReferenceID", plan.ClientReferenceID.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "client_reference_id", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Options.IsNull() && !plan.Options.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Options", plan.Options) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "options", params)
		}
	}
	if !plan.ProvidedDetails.IsNull() && !plan.ProvidedDetails.IsUnknown() {
		if !assignAttrValueToNamedField(params, "ProvidedDetails", plan.ProvidedDetails) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "provided_details", params)
		}
	}
	if !plan.RelatedCustomer.IsNull() && !plan.RelatedCustomer.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "RelatedCustomer", "RelatedCustomer", plan.RelatedCustomer.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "related_customer", params)
		}
	}
	if !plan.RelatedCustomerAccount.IsNull() && !plan.RelatedCustomerAccount.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "RelatedCustomerAccount", "RelatedCustomerAccount", plan.RelatedCustomerAccount.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "related_customer_account", params)
		}
	}
	if !plan.RelatedPerson.IsNull() && !plan.RelatedPerson.IsUnknown() {
		if !assignAttrValueToNamedField(params, "RelatedPerson", plan.RelatedPerson) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "related_person", params)
		}
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "Type", "Type", plan.Type.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "type", params)
		}
	}
	if !plan.VerificationFlow.IsNull() && !plan.VerificationFlow.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "VerificationFlow", "VerificationFlow", plan.VerificationFlow.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "verification_flow", params)
		}
	}
	if !plan.ReturnURL.IsNull() && !plan.ReturnURL.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "ReturnURL", "ReturnURL", plan.ReturnURL.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "return_url", params)
		}
	}

	return params, nil
}

func flattenIdentityVerificationSession(obj *stripe.IdentityVerificationSession, state *IdentityVerificationSessionResourceModel) error {
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
		if rawValueClientReferenceID, rawOk := plainValueAtPath(raw, "client_reference_id"); rawOk {
			if valueClientReferenceID, err := flattenPlainValue(rawValueClientReferenceID, types.StringType, "client_reference_id", "raw response"); err != nil {
				return err
			} else {
				if typedClientReferenceID, ok := valueClientReferenceID.(types.String); ok {
					state.ClientReferenceID = typedClientReferenceID
				}
			}
		} else if !hasRaw {
			if responseValueClientReferenceID, ok := plainFromResponseField(obj, "ClientReferenceID"); ok {
				if valueClientReferenceID, err := flattenPlainValue(responseValueClientReferenceID, types.StringType, "client_reference_id", "response struct"); err != nil {
					return err
				} else {
					if typedClientReferenceID, ok := valueClientReferenceID.(types.String); ok {
						state.ClientReferenceID = typedClientReferenceID
					}
				}
			}
		}
	}
	{
		if rawValueClientSecret, rawOk := plainValueAtPath(raw, "client_secret"); rawOk {
			if valueClientSecret, err := flattenPlainValue(rawValueClientSecret, types.StringType, "client_secret", "raw response"); err != nil {
				return err
			} else {
				if typedClientSecret, ok := valueClientSecret.(types.String); ok {
					state.ClientSecret = typedClientSecret
				}
			}
		} else if !hasRaw {
			if responseValueClientSecret, ok := plainFromResponseField(obj, "ClientSecret"); ok {
				if valueClientSecret, err := flattenPlainValue(responseValueClientSecret, types.StringType, "client_secret", "response struct"); err != nil {
					return err
				} else {
					if typedClientSecret, ok := valueClientSecret.(types.String); ok {
						state.ClientSecret = typedClientSecret
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
		assignedLastError := false
		hadRawLastError := false
		if rawValueLastError, rawOk := plainValueAtPath(raw, "last_error"); rawOk {
			hadRawLastError = true
			if rawValueLastError != nil {
				sourceLastError := applyConfiguredKeyedListShapes(rawValueLastError, attrValueToPlain(state.LastError))
				if valueLastError, err := flattenPlainValue(sourceLastError, types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "reason": types.StringType}}, "last_error", "raw response"); err != nil {
					return err
				} else {
					if typedLastError, ok := valueLastError.(types.Object); ok {
						state.LastError = typedLastError
						assignedLastError = true
					}
				}
			}
		}
		if !assignedLastError {
			if !hasRaw {
				if responseValueLastError, ok := plainFromResponseField(obj, "LastError"); ok {
					sourceLastError := applyConfiguredKeyedListShapes(responseValueLastError, attrValueToPlain(state.LastError))
					if valueLastError, err := flattenPlainValue(
						sourceLastError,
						types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "reason": types.StringType}},
						"last_error",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedLastError, ok := valueLastError.(types.Object); ok {
							state.LastError = typedLastError
							assignedLastError = true
						}
					}
				}
			}
		}
		if !assignedLastError && hadRawLastError {
			if nullLastError, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"code": types.StringType, "reason": types.StringType}}); ok {
				if typedLastError, ok := nullLastError.(types.Object); ok {
					state.LastError = typedLastError
				}
			}
		}
	}
	{
		if true {
			if rawValueLastVerificationReport, rawOk := plainValueAtPath(raw, "last_verification_report"); rawOk {
				if typedLastVerificationReport, ok := plainToStringIDValue(rawValueLastVerificationReport); ok {
					state.LastVerificationReport = typedLastVerificationReport
				}
			} else if !hasRaw {
				if responseValueLastVerificationReport, ok := plainFromResponseField(obj, "LastVerificationReport"); ok {
					if typedLastVerificationReport, ok := plainToStringIDValue(responseValueLastVerificationReport); ok {
						state.LastVerificationReport = typedLastVerificationReport
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
		assignedOptions := false
		hadRawOptions := false
		if rawValueOptions, rawOk := plainValueAtPath(raw, "options"); rawOk {
			hadRawOptions = true
			if rawValueOptions != nil {
				sourceOptions := applyConfiguredKeyedListShapes(rawValueOptions, attrValueToPlain(state.Options))
				if valueOptions, err := flattenPlainValue(sourceOptions, types.ObjectType{AttrTypes: map[string]attr.Type{"document": types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_types": types.ListType{ElemType: types.StringType}, "require_id_number": types.BoolType, "require_live_capture": types.BoolType, "require_matching_selfie": types.BoolType}}, "email": types.ObjectType{AttrTypes: map[string]attr.Type{"require_verification": types.BoolType}}, "matching": types.ObjectType{AttrTypes: map[string]attr.Type{"dob": types.StringType, "name": types.StringType}}, "phone": types.ObjectType{AttrTypes: map[string]attr.Type{"require_verification": types.BoolType}}}}, "options", "raw response"); err != nil {
					return err
				} else {
					if typedOptions, ok := valueOptions.(types.Object); ok {
						state.Options = typedOptions
						assignedOptions = true
					}
				}
			}
		}
		if !assignedOptions {
			if !hasRaw {
				if responseValueOptions, ok := plainFromResponseField(obj, "Options"); ok {
					sourceOptions := applyConfiguredKeyedListShapes(responseValueOptions, attrValueToPlain(state.Options))
					if valueOptions, err := flattenPlainValue(
						sourceOptions,
						types.ObjectType{AttrTypes: map[string]attr.Type{"document": types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_types": types.ListType{ElemType: types.StringType}, "require_id_number": types.BoolType, "require_live_capture": types.BoolType, "require_matching_selfie": types.BoolType}}, "email": types.ObjectType{AttrTypes: map[string]attr.Type{"require_verification": types.BoolType}}, "matching": types.ObjectType{AttrTypes: map[string]attr.Type{"dob": types.StringType, "name": types.StringType}}, "phone": types.ObjectType{AttrTypes: map[string]attr.Type{"require_verification": types.BoolType}}}},
						"options",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedOptions, ok := valueOptions.(types.Object); ok {
							state.Options = typedOptions
							assignedOptions = true
						}
					}
				}
			}
		}
		if !assignedOptions && hadRawOptions {
			if nullOptions, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"document": types.ObjectType{AttrTypes: map[string]attr.Type{"allowed_types": types.ListType{ElemType: types.StringType}, "require_id_number": types.BoolType, "require_live_capture": types.BoolType, "require_matching_selfie": types.BoolType}}, "email": types.ObjectType{AttrTypes: map[string]attr.Type{"require_verification": types.BoolType}}, "matching": types.ObjectType{AttrTypes: map[string]attr.Type{"dob": types.StringType, "name": types.StringType}}, "phone": types.ObjectType{AttrTypes: map[string]attr.Type{"require_verification": types.BoolType}}}}); ok {
				if typedOptions, ok := nullOptions.(types.Object); ok {
					state.Options = typedOptions
				}
			}
		}
	}
	{
		assignedProvidedDetails := false
		hadRawProvidedDetails := false
		if rawValueProvidedDetails, rawOk := plainValueAtPath(raw, "provided_details"); rawOk {
			hadRawProvidedDetails = true
			if rawValueProvidedDetails != nil {
				sourceProvidedDetails := applyConfiguredKeyedListShapes(rawValueProvidedDetails, attrValueToPlain(state.ProvidedDetails))
				if valueProvidedDetails, err := flattenPlainValue(sourceProvidedDetails, types.ObjectType{AttrTypes: map[string]attr.Type{"email": types.StringType, "phone": types.StringType}}, "provided_details", "raw response"); err != nil {
					return err
				} else {
					if typedProvidedDetails, ok := valueProvidedDetails.(types.Object); ok {
						state.ProvidedDetails = typedProvidedDetails
						assignedProvidedDetails = true
					}
				}
			}
		}
		if !assignedProvidedDetails {
			if !hasRaw {
				if responseValueProvidedDetails, ok := plainFromResponseField(obj, "ProvidedDetails"); ok {
					sourceProvidedDetails := applyConfiguredKeyedListShapes(responseValueProvidedDetails, attrValueToPlain(state.ProvidedDetails))
					if valueProvidedDetails, err := flattenPlainValue(
						sourceProvidedDetails,
						types.ObjectType{AttrTypes: map[string]attr.Type{"email": types.StringType, "phone": types.StringType}},
						"provided_details",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedProvidedDetails, ok := valueProvidedDetails.(types.Object); ok {
							state.ProvidedDetails = typedProvidedDetails
							assignedProvidedDetails = true
						}
					}
				}
			}
		}
		if !assignedProvidedDetails && hadRawProvidedDetails {
			if nullProvidedDetails, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"email": types.StringType, "phone": types.StringType}}); ok {
				if typedProvidedDetails, ok := nullProvidedDetails.(types.Object); ok {
					state.ProvidedDetails = typedProvidedDetails
				}
			}
		}
	}
	{
		assignedRedaction := false
		hadRawRedaction := false
		if rawValueRedaction, rawOk := plainValueAtPath(raw, "redaction"); rawOk {
			hadRawRedaction = true
			if rawValueRedaction != nil {
				sourceRedaction := applyConfiguredKeyedListShapes(rawValueRedaction, attrValueToPlain(state.Redaction))
				if valueRedaction, err := flattenPlainValue(sourceRedaction, types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}, "redaction", "raw response"); err != nil {
					return err
				} else {
					if typedRedaction, ok := valueRedaction.(types.Object); ok {
						state.Redaction = typedRedaction
						assignedRedaction = true
					}
				}
			}
		}
		if !assignedRedaction {
			if !hasRaw {
				if responseValueRedaction, ok := plainFromResponseField(obj, "Redaction"); ok {
					sourceRedaction := applyConfiguredKeyedListShapes(responseValueRedaction, attrValueToPlain(state.Redaction))
					if valueRedaction, err := flattenPlainValue(
						sourceRedaction,
						types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}},
						"redaction",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRedaction, ok := valueRedaction.(types.Object); ok {
							state.Redaction = typedRedaction
							assignedRedaction = true
						}
					}
				}
			}
		}
		if !assignedRedaction && hadRawRedaction {
			if nullRedaction, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"status": types.StringType}}); ok {
				if typedRedaction, ok := nullRedaction.(types.Object); ok {
					state.Redaction = typedRedaction
				}
			}
		}
	}
	{
		if rawValueRelatedCustomer, rawOk := plainValueAtPath(raw, "related_customer"); rawOk {
			if valueRelatedCustomer, err := flattenPlainValue(rawValueRelatedCustomer, types.StringType, "related_customer", "raw response"); err != nil {
				return err
			} else {
				if typedRelatedCustomer, ok := valueRelatedCustomer.(types.String); ok {
					state.RelatedCustomer = typedRelatedCustomer
				}
			}
		} else if !hasRaw {
			if responseValueRelatedCustomer, ok := plainFromResponseField(obj, "RelatedCustomer"); ok {
				if valueRelatedCustomer, err := flattenPlainValue(responseValueRelatedCustomer, types.StringType, "related_customer", "response struct"); err != nil {
					return err
				} else {
					if typedRelatedCustomer, ok := valueRelatedCustomer.(types.String); ok {
						state.RelatedCustomer = typedRelatedCustomer
					}
				}
			}
		}
	}
	{
		if rawValueRelatedCustomerAccount, rawOk := plainValueAtPath(raw, "related_customer_account"); rawOk {
			if valueRelatedCustomerAccount, err := flattenPlainValue(rawValueRelatedCustomerAccount, types.StringType, "related_customer_account", "raw response"); err != nil {
				return err
			} else {
				if typedRelatedCustomerAccount, ok := valueRelatedCustomerAccount.(types.String); ok {
					state.RelatedCustomerAccount = typedRelatedCustomerAccount
				}
			}
		} else if !hasRaw {
			if responseValueRelatedCustomerAccount, ok := plainFromResponseField(obj, "RelatedCustomerAccount"); ok {
				if valueRelatedCustomerAccount, err := flattenPlainValue(responseValueRelatedCustomerAccount, types.StringType, "related_customer_account", "response struct"); err != nil {
					return err
				} else {
					if typedRelatedCustomerAccount, ok := valueRelatedCustomerAccount.(types.String); ok {
						state.RelatedCustomerAccount = typedRelatedCustomerAccount
					}
				}
			}
		}
	}
	{
		assignedRelatedPerson := false
		hadRawRelatedPerson := false
		if rawValueRelatedPerson, rawOk := plainValueAtPath(raw, "related_person"); rawOk {
			hadRawRelatedPerson = true
			if rawValueRelatedPerson != nil {
				sourceRelatedPerson := applyConfiguredKeyedListShapes(rawValueRelatedPerson, attrValueToPlain(state.RelatedPerson))
				if valueRelatedPerson, err := flattenPlainValue(sourceRelatedPerson, types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "person": types.StringType}}, "related_person", "raw response"); err != nil {
					return err
				} else {
					if typedRelatedPerson, ok := valueRelatedPerson.(types.Object); ok {
						state.RelatedPerson = typedRelatedPerson
						assignedRelatedPerson = true
					}
				}
			}
		}
		if !assignedRelatedPerson {
			if !hasRaw {
				if responseValueRelatedPerson, ok := plainFromResponseField(obj, "RelatedPerson"); ok {
					sourceRelatedPerson := applyConfiguredKeyedListShapes(responseValueRelatedPerson, attrValueToPlain(state.RelatedPerson))
					if valueRelatedPerson, err := flattenPlainValue(
						sourceRelatedPerson,
						types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "person": types.StringType}},
						"related_person",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedRelatedPerson, ok := valueRelatedPerson.(types.Object); ok {
							state.RelatedPerson = typedRelatedPerson
							assignedRelatedPerson = true
						}
					}
				}
			}
		}
		if !assignedRelatedPerson && hadRawRelatedPerson {
			if nullRelatedPerson, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"account": types.StringType, "person": types.StringType}}); ok {
				if typedRelatedPerson, ok := nullRelatedPerson.(types.Object); ok {
					state.RelatedPerson = typedRelatedPerson
				}
			}
		}
	}
	{
		if rawValueStatus, rawOk := plainValueAtPath(raw, "status"); rawOk {
			if valueStatus, err := flattenPlainValue(rawValueStatus, types.StringType, "status", "raw response"); err != nil {
				return err
			} else {
				if typedStatus, ok := valueStatus.(types.String); ok {
					state.Status = typedStatus
				}
			}
		} else if !hasRaw {
			if responseValueStatus, ok := plainFromResponseField(obj, "Status"); ok {
				if valueStatus, err := flattenPlainValue(responseValueStatus, types.StringType, "status", "response struct"); err != nil {
					return err
				} else {
					if typedStatus, ok := valueStatus.(types.String); ok {
						state.Status = typedStatus
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
		if rawValueURL, rawOk := plainValueAtPath(raw, "url"); rawOk {
			if valueURL, err := flattenPlainValue(rawValueURL, types.StringType, "url", "raw response"); err != nil {
				return err
			} else {
				if typedURL, ok := valueURL.(types.String); ok {
					state.URL = typedURL
				}
			}
		} else if !hasRaw {
			if responseValueURL, ok := plainFromResponseField(obj, "URL"); ok {
				if valueURL, err := flattenPlainValue(responseValueURL, types.StringType, "url", "response struct"); err != nil {
					return err
				} else {
					if typedURL, ok := valueURL.(types.String); ok {
						state.URL = typedURL
					}
				}
			}
		}
	}
	{
		if rawValueVerificationFlow, rawOk := plainValueAtPath(raw, "verification_flow"); rawOk {
			if valueVerificationFlow, err := flattenPlainValue(rawValueVerificationFlow, types.StringType, "verification_flow", "raw response"); err != nil {
				return err
			} else {
				if typedVerificationFlow, ok := valueVerificationFlow.(types.String); ok {
					state.VerificationFlow = typedVerificationFlow
				}
			}
		} else if !hasRaw {
			if responseValueVerificationFlow, ok := plainFromResponseField(obj, "VerificationFlow"); ok {
				if valueVerificationFlow, err := flattenPlainValue(responseValueVerificationFlow, types.StringType, "verification_flow", "response struct"); err != nil {
					return err
				} else {
					if typedVerificationFlow, ok := valueVerificationFlow.(types.String); ok {
						state.VerificationFlow = typedVerificationFlow
					}
				}
			}
		}
	}
	{
		assignedVerifiedOutputs := false
		hadRawVerifiedOutputs := false
		if rawValueVerifiedOutputs, rawOk := plainValueAtPath(raw, "verified_outputs"); rawOk {
			hadRawVerifiedOutputs = true
			if rawValueVerifiedOutputs != nil {
				sourceVerifiedOutputs := applyConfiguredKeyedListShapes(rawValueVerifiedOutputs, attrValueToPlain(state.VerifiedOutputs))
				if valueVerifiedOutputs, err := flattenPlainValue(sourceVerifiedOutputs, types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "dob": types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}, "email": types.StringType, "first_name": types.StringType, "id_number": types.StringType, "id_number_type": types.StringType, "last_name": types.StringType, "phone": types.StringType, "sex": types.StringType, "unparsed_place_of_birth": types.StringType, "unparsed_sex": types.StringType}}, "verified_outputs", "raw response"); err != nil {
					return err
				} else {
					if typedVerifiedOutputs, ok := valueVerifiedOutputs.(types.Object); ok {
						state.VerifiedOutputs = typedVerifiedOutputs
						assignedVerifiedOutputs = true
					}
				}
			}
		}
		if !assignedVerifiedOutputs {
			if !hasRaw {
				if responseValueVerifiedOutputs, ok := plainFromResponseField(obj, "VerifiedOutputs"); ok {
					sourceVerifiedOutputs := applyConfiguredKeyedListShapes(responseValueVerifiedOutputs, attrValueToPlain(state.VerifiedOutputs))
					if valueVerifiedOutputs, err := flattenPlainValue(
						sourceVerifiedOutputs,
						types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "dob": types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}, "email": types.StringType, "first_name": types.StringType, "id_number": types.StringType, "id_number_type": types.StringType, "last_name": types.StringType, "phone": types.StringType, "sex": types.StringType, "unparsed_place_of_birth": types.StringType, "unparsed_sex": types.StringType}},
						"verified_outputs",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedVerifiedOutputs, ok := valueVerifiedOutputs.(types.Object); ok {
							state.VerifiedOutputs = typedVerifiedOutputs
							assignedVerifiedOutputs = true
						}
					}
				}
			}
		}
		if !assignedVerifiedOutputs && hadRawVerifiedOutputs {
			if nullVerifiedOutputs, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"address": types.ObjectType{AttrTypes: map[string]attr.Type{"city": types.StringType, "country": types.StringType, "line1": types.StringType, "line2": types.StringType, "postal_code": types.StringType, "state": types.StringType}}, "dob": types.ObjectType{AttrTypes: map[string]attr.Type{"day": types.Int64Type, "month": types.Int64Type, "year": types.Int64Type}}, "email": types.StringType, "first_name": types.StringType, "id_number": types.StringType, "id_number_type": types.StringType, "last_name": types.StringType, "phone": types.StringType, "sex": types.StringType, "unparsed_place_of_birth": types.StringType, "unparsed_sex": types.StringType}}); ok {
				if typedVerifiedOutputs, ok := nullVerifiedOutputs.(types.Object); ok {
					state.VerifiedOutputs = typedVerifiedOutputs
				}
			}
		}
	}
	return nil
}

func (r *IdentityVerificationSessionEphemeralResource) Open(ctx context.Context, req ephemeral.OpenRequest, resp *ephemeral.OpenResponse) {
	var config IdentityVerificationSessionResourceModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandIdentityVerificationSessionCreate(config)
	if err != nil {
		resp.Diagnostics.AddError("Error building IdentityVerificationSession ephemeral params", err.Error())
		return
	}

	obj, err := r.client.V1IdentityVerificationSessions.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error opening IdentityVerificationSession ephemeral resource", err.Error())
		return
	}

	result := config
	if err := flattenIdentityVerificationSession(obj, &result); err != nil {
		resp.Diagnostics.AddError("Error flattening IdentityVerificationSession ephemeral response", err.Error())
		return
	}
	normalizeUnknownValues(&result)
	diags = resp.Result.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
}
