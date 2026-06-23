# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
ephemeral "stripe_identity_verification_session" "test" {
  type                = "document"
  client_reference_id = "acctest-identity-document-options"

  metadata = {
    suite = "sdk-codegen"
    case  = "identity_verification_session_document_options"
  }

  options = {
    document = {
      allowed_types           = ["driving_license", "passport"]
      require_id_number       = true
      require_matching_selfie = true
    }
  }
}
