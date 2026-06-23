#!/usr/bin/env bash
# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
set -euo pipefail

REQUIRE_ENV="0"

while [[ $# -gt 0 ]]; do
  case "$1" in
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

if [[ -f .env ]]; then
  set -a
  # shellcheck disable=SC1091
  source ./.env
  set +a
fi

declare -a PHASE_NAMES=()
declare -a PHASE_RESULTS=()

run_phase() {
  local name="$1"
  shift

  PHASE_NAMES+=("$name")

  echo
  echo "==> $name"
  echo "Command: $*"
  echo

  if "$@"; then
    PHASE_RESULTS+=("PASS")
    echo
    echo "==> $name: PASS"
    return 0
  fi

  PHASE_RESULTS+=("FAIL")
  echo
  echo "==> $name: FAIL"
  return 1
}

unit_status=0
acceptance_status=0

ARGS=()
if [[ "$REQUIRE_ENV" == "1" ]]; then
  ARGS+=(--require-env)
fi

declare -a UNIT_PACKAGES=()
while IFS= read -r package; do
  if [[ "$package" != */internal/acctest/* ]]; then
    UNIT_PACKAGES+=("$package")
  fi
done < <(go list ./...)

if ! run_phase "Go Tests (non-acctest)" go test "${UNIT_PACKAGES[@]}" -count=1 -v; then
  unit_status=1
fi

if ! run_phase \
  "Terraform Acceptance Tests" \
  bash ./scripts/testacc-nightly.sh "${ARGS[@]}"; then
  acceptance_status=1
fi

echo
echo "==> Test Summary"
for i in "${!PHASE_NAMES[@]}"; do
  printf '  [%s] %s\n' "${PHASE_RESULTS[$i]}" "${PHASE_NAMES[$i]}"
done

if [[ "$unit_status" -ne 0 || "$acceptance_status" -ne 0 ]]; then
  exit 1
fi
