# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_terminal_reader" "test" {
  registration_code = "{{TERMINAL_READER_REGISTRATION_CODE_REGRESSION}}"
  label             = "SDK Codegen Reader Regression"
  location          = "{{TERMINAL_LOCATION}}"

  metadata = {
    suite = "sdk-codegen"
    case  = "terminal_reader_registration_code_regression"
  }
}
