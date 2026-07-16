#!/usr/bin/env bash
# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
set -euo pipefail

LEGACY_TAG="${LEGACY_TAG:-v0.2.2}"
RUN_REGEX='TestAccManaged.*LegacyUpgrade$'
PACKAGE='./internal/acctest/cases/managed/...'
AUDIT_ONLY="0"
IGNORE_MISSING_REGEX="${IGNORE_MISSING_REGEX:-^v2_billing_}"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --audit-only)
      AUDIT_ONLY="1"
      shift
      ;;
    --legacy-tag)
      LEGACY_TAG="$2"
      shift 2
      ;;
    --package)
      PACKAGE="$2"
      shift 2
      ;;
    --run)
      RUN_REGEX="$2"
      shift 2
      ;;
    --ignore-missing-regex)
      IGNORE_MISSING_REGEX="$2"
      shift 2
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

export TF_ACC=1

if [[ -z "${TF_ACC_TERRAFORM_PATH:-}" ]] && command -v terraform >/dev/null 2>&1; then
  export TF_ACC_TERRAFORM_PATH
  TF_ACC_TERRAFORM_PATH="$(command -v terraform)"
fi

list_legacy_resources() {
  git ls-tree -r --name-only "$LEGACY_TAG" internal/provider/resources |
    sed -n 's#^internal/provider/resources/resource_\(.*\)\.go$#\1#p' |
    sort -u
}

list_current_resources() {
  ls internal/provider/resources |
    sed -n 's#^resource_stripe_\(.*\)\.go$#\1#p' |
    sort -u
}

list_state_upgrader_resources() {
  rg -l 'ResourceWithUpgradeState' internal/provider/resources/resource_stripe_*.go |
    sed -E 's#^.*/resource_stripe_(.*)\.go$#\1#' |
    sort -u
}

coverage_registry_path="internal/acctest/runner/coverage_registry_generated.go"

list_legacy_tested_resources() {
  awk '
    /Surface:/ {
      surface = $2
      gsub(/[",]/, "", surface)
      sub(/^stripe_/, "", surface)
    }
    /LegacyUpgradeCaseNames:/ && $0 !~ /nil/ {
      print surface
    }
  ' "$coverage_registry_path" | sort -u
}

list_required_module_upgrade_resources() {
  awk '
    /Surface:/ {
      surface = $2
      gsub(/[",]/, "", surface)
      sub(/^stripe_/, "", surface)
    }
    /RequiresLegacyModuleUpgrade:/ && /true/ {
      print surface
    }
  ' "$coverage_registry_path" | sort -u
}

list_module_upgrade_tested_resources() {
  awk '
    /Surface:/ {
      surface = $2
      gsub(/[",]/, "", surface)
      sub(/^stripe_/, "", surface)
    }
    /LegacyModuleUpgradeCaseNames:/ && $0 !~ /nil/ {
      print surface
    }
  ' "$coverage_registry_path" | sort -u
}

print_set() {
  local title="$1"
  local values="$2"

  echo "$title"
  if [[ -z "$values" ]]; then
    echo "  (none)"
    return
  fi

  while IFS= read -r value; do
    [[ -n "$value" ]] || continue
    printf '  - %s\n' "$value"
  done <<<"$values"
}

filter_ignored_missing_resources() {
  local values="$1"

  if [[ -z "$IGNORE_MISSING_REGEX" ]]; then
    printf '%s\n' "$values"
    return
  fi

  printf '%s\n' "$values" | rg -v "$IGNORE_MISSING_REGEX" || true
}

echo "==> Legacy upgrade audit"
printf '  Legacy tag: %s\n' "$LEGACY_TAG"
if [[ -n "$IGNORE_MISSING_REGEX" ]]; then
  printf '  Ignoring missing resources matching: %s\n' "$IGNORE_MISSING_REGEX"
fi

legacy_resources="$(list_legacy_resources)"
current_resources="$(list_current_resources)"
state_upgrader_resources="$(list_state_upgrader_resources)"
legacy_tested_resources="$(list_legacy_tested_resources)"
required_module_upgrade_resources="$(list_required_module_upgrade_resources)"
module_upgrade_tested_resources="$(list_module_upgrade_tested_resources)"

common_legacy_resources="$(comm -12 <(printf '%s\n' "$legacy_resources") <(printf '%s\n' "$current_resources"))"
missing_current_resources="$(comm -23 <(printf '%s\n' "$legacy_resources") <(printf '%s\n' "$current_resources"))"
missing_current_resources="$(filter_ignored_missing_resources "$missing_current_resources")"
common_without_state_upgrader="$(comm -23 <(printf '%s\n' "$common_legacy_resources") <(printf '%s\n' "$state_upgrader_resources"))"
common_without_legacy_test="$(comm -23 <(printf '%s\n' "$common_legacy_resources") <(printf '%s\n' "$legacy_tested_resources"))"
required_module_upgrade_resources="$(comm -12 <(printf '%s\n' "$common_legacy_resources") <(printf '%s\n' "$required_module_upgrade_resources"))"
common_without_module_upgrade_test="$(comm -23 <(printf '%s\n' "$required_module_upgrade_resources") <(printf '%s\n' "$module_upgrade_tested_resources"))"

printf '  Legacy resources in %s: %s\n' "$LEGACY_TAG" "$(printf '%s\n' "$legacy_resources" | sed '/^$/d' | wc -l | tr -d ' ')"
printf '  Current resource files: %s\n' "$(printf '%s\n' "$current_resources" | sed '/^$/d' | wc -l | tr -d ' ')"
printf '  Shared legacy resources still present: %s\n' "$(printf '%s\n' "$common_legacy_resources" | sed '/^$/d' | wc -l | tr -d ' ')"
printf '  Shared legacy resources requiring module-upgrade coverage: %s\n' "$(printf '%s\n' "$required_module_upgrade_resources" | sed '/^$/d' | wc -l | tr -d ' ')"

compat_status=0

if [[ -n "$missing_current_resources" ]]; then
  compat_status=1
  print_set "Resources present in ${LEGACY_TAG} but missing now:" "$missing_current_resources"
fi

if [[ -n "$common_without_state_upgrader" ]]; then
  compat_status=1
  print_set "Shared legacy resources without a state upgrader:" "$common_without_state_upgrader"
fi

if [[ -n "$common_without_legacy_test" ]]; then
  compat_status=1
  print_set "Shared legacy resources without a legacy-upgrade acceptance case:" "$common_without_legacy_test"
fi

if [[ -n "$common_without_module_upgrade_test" ]]; then
  compat_status=1
  print_set "Shared legacy resources without a module-style legacy-upgrade acceptance case:" "$common_without_module_upgrade_test"
fi

if [[ "$compat_status" -eq 0 ]]; then
  echo "  Compatibility audit: PASS"
else
  echo "  Compatibility audit: FAIL"
fi

if [[ "$AUDIT_ONLY" == "1" ]]; then
  exit "$compat_status"
fi

echo
echo "==> Running legacy-upgrade acceptance tests"
printf '  Package: %s\n' "$PACKAGE"
printf '  Regex: %s\n' "$RUN_REGEX"
echo

test_status=0
if ! go test "$PACKAGE" -run "$RUN_REGEX" -count=1 -v; then
  test_status=1
fi

if [[ "$compat_status" -ne 0 || "$test_status" -ne 0 ]]; then
  exit 1
fi
