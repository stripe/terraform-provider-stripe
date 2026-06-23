# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_issuing_personalization_design" "test" {
  physical_bundle = "{{ISSUING_PHYSICAL_BUNDLE}}"
  card_logo       = "{{ISSUING_CARD_LOGO}}"
  name            = "SDK Codegen Design Default Create"
  lookup_key      = "sdk-codegen-design-default-create-{{RAND}}"

  carrier_text = {
    header_title = "SDK Default"
    header_body  = "Create"
    footer_title = "Desk"
    footer_body  = "create@example.com"
  }

  preferences = {
    is_default = true
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "issuing_personalization_design_default_toggle"
    phase = "create"
  }
}
