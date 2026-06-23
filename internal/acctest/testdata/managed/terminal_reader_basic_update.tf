# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_terminal_reader" "test" {
  registration_code = "{{TERMINAL_READER_REGISTRATION_CODE}}"
  label             = "SDK Codegen Reader Basic Updated"
  location          = "{{TERMINAL_LOCATION}}"

  metadata = {
    suite = "sdk-codegen"
    case  = "terminal_reader_basic"
    phase = "update"
  }
}
