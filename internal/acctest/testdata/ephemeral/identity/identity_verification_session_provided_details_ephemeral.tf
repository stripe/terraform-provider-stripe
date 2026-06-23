# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
ephemeral "stripe_identity_verification_session" "test" {
  type                = "document"
  client_reference_id = "acctest-identity-provided-details"

  metadata = {
    suite = "sdk-codegen"
    case  = "identity_verification_session_provided_details"
  }

  provided_details = {
    email = "sdk-codegen+identity@example.com"
    phone = "+15555550123"
  }
}
