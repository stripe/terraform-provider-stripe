# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_radar_value_list" "test" {
  alias     = "sdk_codegen_email_list_update_{{RAND}}"
  name      = "SDK Codegen Radar Email List Update"
  item_type = "email"

  metadata = {
    suite = "sdk-codegen"
    case  = "radar_value_list_email_alias_update"
    phase = "update"
  }
}
