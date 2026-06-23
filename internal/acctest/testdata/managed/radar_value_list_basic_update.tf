# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_radar_value_list" "test" {
  alias     = "sdk_codegen_list_{{RAND}}"
  name      = "SDK Codegen Radar List Update"
  item_type = "string"

  metadata = {
    suite = "sdk-codegen"
    case  = "radar_value_list_basic"
    phase = "update"
  }
}
