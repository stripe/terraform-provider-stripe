# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_entitlements_feature" "test" {
  lookup_key = "acctest-feature-metadata-{{RAND}}"
  name       = "SDK Feature Metadata {{RAND}}"

  metadata = {
    suite   = "sdk-codegen"
    case    = "entitlements_feature_metadata_update"
    phase   = "create"
    rollout = "alpha"
    owner   = "billing"
  }
}
