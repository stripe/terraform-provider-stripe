# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
set shell := ["bash", "-eu", "-o", "pipefail", "-c"]

format:
    gofmt -w $(find . -name '*.go' -not -path './.terraform/*')

testacc *args:
    bash ./scripts/testacc.sh {{ args }}

testacc-nightly *args:
    bash ./scripts/testacc-nightly.sh {{ args }}

testacc-upgrade *args:
    bash ./scripts/testacc-upgrade.sh {{ args }}

terraform-acceptance *args:
    bash ./scripts/testacc.sh --require-env {{ args }}

testall *args:
    bash ./scripts/testall.sh {{ args }}
