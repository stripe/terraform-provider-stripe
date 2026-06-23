#!/usr/bin/env bash
# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
set -euo pipefail

REQUIRE_ENV="0"
KIND=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --kind)
      KIND="$2"
      shift 2
      ;;
    --require-env)
      REQUIRE_ENV="1"
      shift
      ;;
    *)
      echo "unknown argument: $1" >&2
      exit 1
      ;;
  esac
done

ARGS=()
if [[ -n "$KIND" ]]; then
  ARGS+=(--kind "$KIND")
fi
if [[ "$REQUIRE_ENV" == "1" ]]; then
  ARGS+=(--require-env)
fi

bash ./scripts/testacc.sh "${ARGS[@]}"
