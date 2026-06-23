# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_file" "test" {
  purpose = "dispute_evidence"
  file    = "{{FILE_UPLOAD_FIXTURE_PATH}}"
}
