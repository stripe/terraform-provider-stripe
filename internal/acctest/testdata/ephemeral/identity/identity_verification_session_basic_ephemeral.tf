# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
ephemeral "stripe_identity_verification_session" "test" {
  type                = "document"
  client_reference_id = "acctest-identity-basic"

  metadata = {
    suite = "sdk-codegen"
    case  = "identity_verification_session_basic"
  }

  options = {
    document = {
      allowed_types        = ["passport"]
      require_live_capture = true
    }
  }
}
