# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_entitlements_feature" "test" {
  lookup_key = "acctest-feature-{{RAND}}"
  name       = "SDK Feature {{RAND}}"

  metadata = {
    suite = "sdk-codegen"
    case  = "entitlements_feature_basic"
    phase = "update"
  }
}
