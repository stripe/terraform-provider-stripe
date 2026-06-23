# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_radar_value_list" "list" {
  alias     = "sdk_codegen_item_list_{{RAND}}"
  name      = "SDK Codegen Radar Item List"
  item_type = "email"
}

resource "stripe_radar_value_list_item" "test" {
  value      = "sdk-codegen@example.com"
  value_list = stripe_radar_value_list.list.id
}
