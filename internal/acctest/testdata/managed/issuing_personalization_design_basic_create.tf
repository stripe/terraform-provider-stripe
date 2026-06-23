# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_issuing_personalization_design" "test" {
  physical_bundle = "{{ISSUING_PHYSICAL_BUNDLE}}"
  card_logo       = "{{ISSUING_CARD_LOGO}}"
  name            = "SDK Codegen Personalization Create"
  lookup_key      = "sdk-codegen-design-create-{{RAND}}"

  carrier_text = {
    header_title = "SDK Codegen"
    header_body  = "Create"
    footer_title = "Support"
    footer_body  = "support@example.com"
  }

  preferences = {
    is_default = false
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "issuing_personalization_design_basic"
    phase = "create"
  }
}
