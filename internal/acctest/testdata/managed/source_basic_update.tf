# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_source" "test" {
  type  = "card"
  token = "tok_visa"

  owner = {
    name  = "SDK Codegen Source Update"
    email = "sdk-codegen+source-update@example.com"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "source_basic"
    phase = "update"
  }
}
