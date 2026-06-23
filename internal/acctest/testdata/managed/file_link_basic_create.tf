# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_file" "file" {
  purpose = "dispute_evidence"
  file    = "{{FILE_UPLOAD_FIXTURE_PATH}}"
}

resource "stripe_file_link" "test" {
  file       = stripe_file.file.id
  expires_at = 1893456000

  metadata = {
    suite = "sdk-codegen"
    case  = "file_link_basic"
    phase = "create"
  }
}
