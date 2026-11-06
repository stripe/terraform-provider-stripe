#!/usr/bin/env bash
# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
set -euo pipefail

GROUP=""
SURFACE=""
CASE_NAME=""
CASE_NAMES=()
REQUIRE_ENV="0"
KIND=""

while [[ $# -gt 0 ]]; do
  case "$1" in
    --kind)
      KIND="$2"
      shift 2
      ;;
    --group)
      GROUP="$2"
      shift 2
      ;;
    --surface)
      SURFACE="$2"
      shift 2
      ;;
    --case)
      CASE_NAME="$2"
      CASE_NAMES+=("$2")
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

if [[ -f .env ]]; then
  set -a
  # shellcheck disable=SC1091
  source ./.env
  set +a
fi

export TF_ACC=1
export STRIPE_TF_ACCTEST_GROUP="$GROUP"
export STRIPE_TF_ACCTEST_SURFACE="$SURFACE"
export STRIPE_TF_ACCTEST_CASE="$CASE_NAME"
export STRIPE_TF_ACCTEST_REQUIRE_ENV="$REQUIRE_ENV"

if [[ -z "${TF_ACC_TERRAFORM_PATH:-}" ]] && command -v terraform >/dev/null 2>&1; then
  export TF_ACC_TERRAFORM_PATH
  TF_ACC_TERRAFORM_PATH="$(command -v terraform)"
fi

# Only an unfiltered managed/resource run should fail on skips. Selector-based
# runs intentionally skip unrelated cases in the same package set.
should_require_full_suite_coverage() {
  [[ "$REQUIRE_ENV" == "1" &&
    -z "$GROUP" &&
    -z "$SURFACE" &&
    -z "$CASE_NAME" &&
    ( -z "$KIND" || "$KIND" == "managed" || "$KIND" == "resource" ) ]]
}

should_seed_dispute_txn="0"
if [[ "$KIND" != "action" && ( -z "$GROUP" || "$GROUP" == "issuing" ) ]]; then
  should_seed_dispute_txn="1"
fi
if [[ "$SURFACE" == "stripe_issuing_dispute" || "$CASE_NAME" == "issuing_dispute_basic" ]]; then
  should_seed_dispute_txn="1"
fi

strict_dispute_target="0"
if [[ "$SURFACE" == "stripe_issuing_dispute" || "$CASE_NAME" == "issuing_dispute_basic" ]]; then
  strict_dispute_target="1"
fi

should_seed_credit_note_invoice="0"
if [[ "$KIND" != "action" && "$KIND" != "ephemeral" && ( -z "$GROUP" || "$GROUP" == "base" ) ]]; then
  should_seed_credit_note_invoice="1"
fi
if [[ "$SURFACE" == "stripe_credit_note" || "$CASE_NAME" == "credit_note_basic" ]]; then
  should_seed_credit_note_invoice="1"
fi

strict_credit_note_target="0"
if [[ "$SURFACE" == "stripe_credit_note" || "$CASE_NAME" == "credit_note_basic" ]]; then
  strict_credit_note_target="1"
fi

should_prepare_treasury_financial_account="0"
if [[ "$KIND" != "action" && "$KIND" != "ephemeral" && ( -z "$GROUP" || "$GROUP" == "treasury" ) ]]; then
  should_prepare_treasury_financial_account="1"
fi
if [[ "$SURFACE" == "stripe_treasury_financial_account" || "$CASE_NAME" == "treasury_financial_account_basic" ]]; then
  should_prepare_treasury_financial_account="1"
fi

strict_treasury_financial_account_target="0"
if [[ "$SURFACE" == "stripe_treasury_financial_account" || "$CASE_NAME" == "treasury_financial_account_basic" ]]; then
  strict_treasury_financial_account_target="1"
fi

env_status() {
  local name="$1"
  if [[ -n "${!name:-}" ]]; then
    printf '  [set] %s\n' "$name"
  else
    printf '  [missing] %s\n' "$name"
  fi
}

print_env_summary() {
  echo "==> Acceptance environment"
  env_status STRIPE_API_KEY

  if [[ -z "$GROUP" || "$GROUP" == "connect" ]]; then
    env_status STRIPE_ACCOUNT
  fi
  if [[ -z "$GROUP" || "$GROUP" == "issuing" ]]; then
    env_status STRIPE_ISSUING_ACCOUNT
  fi
  if [[ -z "$GROUP" || "$GROUP" == "treasury" ]]; then
    env_status STRIPE_TREASURY_ACCOUNT
  fi

  if [[ -z "${STRIPE_FORWARDING_DESTINATION:-}" ]]; then
    echo "  [default] STRIPE_FORWARDING_DESTINATION=https://forwarding-api-demo.stripedemos.com/payments"
  else
    echo "  [set] STRIPE_FORWARDING_DESTINATION"
  fi

  if should_require_full_suite_coverage; then
    echo "  Full-suite fixture env:"
    env_status STRIPE_PAYMENT_METHOD_DOMAIN
    env_status STRIPE_APPLE_PAY_DOMAIN
    env_status STRIPE_TERMINAL_LOCATION
    env_status STRIPE_TERMINAL_READER_REGISTRATION_CODE
    env_status STRIPE_TERMINAL_READER_REGISTRATION_CODE_UPDATE
    env_status STRIPE_TERMINAL_READER_REGISTRATION_CODE_REGRESSION
    env_status STRIPE_ISSUING_FINANCIAL_ACCOUNT
    env_status STRIPE_ISSUING_PHYSICAL_BUNDLE
    env_status STRIPE_ISSUING_CARD_LOGO
  fi
}

print_test_summary() {
  local log_path="$1"
  local test_status="$2"
  local pass_count
  local fail_count
  local skip_count
  pass_count="$(awk '/^--- PASS: (TestAcc|TestActionFixture)/ {count++} END {print count + 0}' "$log_path")"
  fail_count="$(awk '/^--- FAIL: (TestAcc|TestActionFixture)/ {count++} END {print count + 0}' "$log_path")"
  skip_count="$(awk '/^--- SKIP: (TestAcc|TestActionFixture)/ {count++} END {print count + 0}' "$log_path")"

  echo
  echo "==> Terraform Acceptance Summary"
  printf '  PASS: %s\n' "$pass_count"
  printf '  FAIL: %s\n' "$fail_count"
  printf '  SKIP: %s\n' "$skip_count"
  printf '  go test exit code: %s\n' "$test_status"

  if [[ "$fail_count" -gt 0 ]]; then
    echo "  Failed tests:"
    awk '/^--- FAIL: (TestAcc|TestActionFixture)/ {print "    - " $3}' "$log_path" | sort -u
  fi

  if [[ "$REQUIRE_ENV" == "1" && "$skip_count" -gt 0 ]]; then
    echo "  Skipped tests:"
    awk '/^--- SKIP: (TestAcc|TestActionFixture)/ {print "    - " $3}' "$log_path" | sort -u
  fi

  if [[ "$test_status" -ne 0 || "$fail_count" -gt 0 || "$pass_count" -eq 0 ]]; then
    echo "TERRAFORM ACCEPTANCE: FAIL"
    return 1
  fi

  if should_require_full_suite_coverage && [[ "$skip_count" -gt 0 ]]; then
    echo "TERRAFORM ACCEPTANCE: FAIL"
    return 1
  fi

  echo "TERRAFORM ACCEPTANCE: PASS"
  return 0
}

seed_issuing_dispute_transaction() {
  local output_env="${1:-STRIPE_ISSUING_DISPUTE_TRANSACTION}"
  if [[ -z "${STRIPE_API_KEY:-}" || -z "${STRIPE_ISSUING_ACCOUNT:-}" ]]; then
    return 1
  fi

  local amount_cents="${STRIPE_ISSUING_DISPUTE_SEED_AMOUNT_CENTS:-123}"
  local buffer_cents="${STRIPE_ISSUING_DISPUTE_SEED_BUFFER_CENTS:-500}"

  api() {
    local method="$1"
    local path="$2"
    shift 2
    curl -sS -X "$method" "https://api.stripe.com${path}" \
      -u "${STRIPE_API_KEY}:" \
      -H "Stripe-Account: ${STRIPE_ISSUING_ACCOUNT}" \
      "$@"
  }

  create_seed_card() {
    local target_financial_account="$1"
    local now
    now="$(date +%s)"

    local cardholder_id
    cardholder_id="$(
      api POST "/v1/issuing/cardholders" \
        -d "type=individual" \
        -d "name=SDK Dispute Seeder" \
        -d "email=sdk-dispute-seeder@example.com" \
        -d "phone_number=+15555550199" \
        -d "status=active" \
        -d "billing[address][line1]=100 Main St" \
        -d "billing[address][city]=San Francisco" \
        -d "billing[address][postal_code]=94105" \
        -d "billing[address][country]=US" \
        -d "billing[address][state]=CA" \
        -d "individual[first_name]=Grace" \
        -d "individual[last_name]=Hopper" \
        -d "individual[dob][day]=9" \
        -d "individual[dob][month]=12" \
        -d "individual[dob][year]=1985" \
        -d "individual[card_issuing][user_terms_acceptance][ip]=127.0.0.1" \
        -d "individual[card_issuing][user_terms_acceptance][date]=${now}" \
        | jq -r '.id'
    )"
    if [[ -z "$cardholder_id" || "$cardholder_id" == "null" ]]; then
      return 1
    fi

    api POST "/v1/issuing/cards" \
      -d "cardholder=${cardholder_id}" \
      -d "financial_account=${target_financial_account}" \
      -d "currency=usd" \
      -d "type=virtual" \
      -d "status=active" \
      | jq -r '.id'
  }

  local card_id="${STRIPE_ISSUING_CARD:-}"
  local financial_account="${STRIPE_ISSUING_FINANCIAL_ACCOUNT:-}"

  if [[ -z "$card_id" || -z "$financial_account" ]]; then
    local active_card_row
    active_card_row="$(
      api GET "/v1/issuing/cards?limit=100" \
        | jq -r '
            .data[]
            | select(.status == "active" and .currency == "usd" and (.financial_account // "") != "")
            | [.id, .financial_account]
            | @tsv
          ' \
        | head -n 1
    )"

    if [[ -n "$active_card_row" ]]; then
      if [[ -z "$card_id" ]]; then
        card_id="$(printf '%s\n' "$active_card_row" | cut -f1)"
      fi
      if [[ -z "$financial_account" ]]; then
        financial_account="$(printf '%s\n' "$active_card_row" | cut -f2)"
      fi
    fi

    if [[ -z "$financial_account" ]]; then
      financial_account="$(
        api GET "/v1/issuing/cards?limit=100" \
          | jq -r '
              .data[]
              | select(.currency == "usd" and (.financial_account // "") != "")
              | .financial_account
            ' \
          | head -n 1
      )"
    fi

    if [[ -z "$card_id" ]]; then
      if [[ -z "$financial_account" || "$financial_account" == "null" ]]; then
        return 1
      fi
      card_id="$(create_seed_card "$financial_account")"
    fi
  fi

  if [[ -z "$card_id" || -z "$financial_account" ]]; then
    return 1
  fi

  local cash_balance
  cash_balance="$(
    api GET "/v1/treasury/financial_accounts/${financial_account}" \
      | jq -r '.balance.cash.usd'
  )"
  if [[ -z "$cash_balance" || "$cash_balance" == "null" ]]; then
    return 1
  fi

  local required_cash=$((amount_cents + buffer_cents))
  if (( cash_balance < required_cash )); then
    local topup_amount=$((required_cash - cash_balance))
    api POST "/v1/test_helpers/treasury/received_credits" \
      -d financial_account="${financial_account}" \
      -d amount="${topup_amount}" \
      -d currency="usd" \
      -d network="ach" >/dev/null
  fi

  local authorization_id
  authorization_id="$(
    api POST "/v1/test_helpers/issuing/authorizations" \
      -d card="${card_id}" \
      -d amount="${amount_cents}" \
      -d currency="usd" \
      -d "merchant_data[name]=SDK Dispute Seed" \
      -d "merchant_data[category]=taxicabs_limousines" \
      | jq -r '.id'
  )"
  if [[ -z "$authorization_id" || "$authorization_id" == "null" ]]; then
    return 1
  fi

  local transaction_id
  transaction_id="$(
    api POST "/v1/test_helpers/issuing/authorizations/${authorization_id}/capture" \
      | jq -r '.transactions[0].id'
  )"
  if [[ -z "$transaction_id" || "$transaction_id" == "null" ]]; then
    return 1
  fi

  printf -v "$output_env" '%s' "$transaction_id"
  export "$output_env"
  echo "auto-seeded ${output_env}=$transaction_id"
  return 0
}

seed_credit_note_invoice() {
  if [[ -z "${STRIPE_API_KEY:-}" ]]; then
    return 1
  fi

  local credit_note_amount="${STRIPE_CREDIT_NOTE_TEST_AMOUNT_CENTS:-100}"
  local invoice_seed_amount="${STRIPE_CREDIT_NOTE_SEED_INVOICE_AMOUNT_CENTS:-700}"

  # Credit note acceptance runs in the base group, which targets the API-key owner
  # account. Seed in the same account by default to avoid cross-account invoice IDs.
  # Allow explicit override when needed for targeted debugging.
  local seed_account="${STRIPE_CREDIT_NOTE_SEED_ACCOUNT:-}"
  local -a base_headers=()
  if [[ -n "${seed_account}" ]]; then
    base_headers=(-H "Stripe-Account: ${seed_account}")
  fi

  api_base() {
    local method="$1"
    local path="$2"
    shift 2
    if [[ "${#base_headers[@]}" -gt 0 ]]; then
      curl -sS -X "$method" "https://api.stripe.com${path}" \
        -u "${STRIPE_API_KEY}:" \
        "${base_headers[@]}" \
        "$@"
    else
      curl -sS -X "$method" "https://api.stripe.com${path}" \
        -u "${STRIPE_API_KEY}:" \
        "$@"
    fi
  }

  local current_invoice="${STRIPE_CREDIT_NOTE_INVOICE:-}"
  if [[ -n "$current_invoice" ]]; then
    local remaining
    local current_invoice_response
    local current_finalized_at
    current_invoice_response="$(
      api_base GET "/v1/invoices/${current_invoice}"
    )"
    remaining="$(jq -r '.amount_remaining' <<<"$current_invoice_response")"
    current_finalized_at="$(jq -r '.status_transitions.finalized_at // empty' <<<"$current_invoice_response")"
    if [[ "$remaining" =~ ^-?[0-9]+$ ]] && (( remaining >= credit_note_amount )); then
      if [[ "$current_finalized_at" =~ ^[0-9]+$ ]]; then
        export STRIPE_CREDIT_NOTE_INVOICE_FINALIZED_AT="$current_finalized_at"
      fi
      return 0
    fi
  fi

  local customer_id
  customer_id="$(
    api_base POST "/v1/customers" \
      -d "name=SDK Credit Note Customer" \
      -d "email=sdk-credit-note@example.com" \
      | jq -r '.id'
  )"
  if [[ -z "$customer_id" || "$customer_id" == "null" ]]; then
    return 1
  fi

  api_base POST "/v1/invoiceitems" \
    -d "customer=${customer_id}" \
    -d "amount=${invoice_seed_amount}" \
    -d "currency=usd" \
    -d "description=sdk-codegen credit note seed item" >/dev/null

  local invoice_id
  invoice_id="$(
    api_base POST "/v1/invoices" \
      -d "customer=${customer_id}" \
      -d "auto_advance=false" \
      -d "collection_method=send_invoice" \
      -d "days_until_due=30" \
      -d "pending_invoice_items_behavior=include" \
      | jq -r '.id'
  )"
  if [[ -z "$invoice_id" || "$invoice_id" == "null" ]]; then
    return 1
  fi

  local finalized_id
  local finalized_at
  local finalized_invoice
  finalized_invoice="$(
    api_base POST "/v1/invoices/${invoice_id}/finalize"
  )"
  finalized_id="$(jq -r '.id' <<<"$finalized_invoice")"
  finalized_at="$(jq -r '.status_transitions.finalized_at // empty' <<<"$finalized_invoice")"
  if [[ -z "$finalized_id" || "$finalized_id" == "null" ]]; then
    return 1
  fi

  export STRIPE_CREDIT_NOTE_INVOICE="$finalized_id"
  if [[ "$finalized_at" =~ ^[0-9]+$ ]]; then
    export STRIPE_CREDIT_NOTE_INVOICE_FINALIZED_AT="$finalized_at"
  fi
  echo "auto-seeded STRIPE_CREDIT_NOTE_INVOICE=$finalized_id"
  return 0
}

prepare_treasury_financial_account_capacity() {
  if [[ -z "${STRIPE_API_KEY:-}" || -z "${STRIPE_TREASURY_ACCOUNT:-}" ]]; then
    return 1
  fi

  local protected_financial_account="${STRIPE_ISSUING_FINANCIAL_ACCOUNT:-}"
  if [[ -z "$protected_financial_account" ]]; then
    return 1
  fi

  api_treasury() {
    local method="$1"
    local path="$2"
    shift 2
    curl -sS -X "$method" "https://api.stripe.com${path}" \
      -u "${STRIPE_API_KEY}:" \
      -H "Stripe-Account: ${STRIPE_TREASURY_ACCOUNT}" \
      "$@"
  }

  local attempted_close_ids=""
  while true; do
    local open_count
    open_count="$(
      api_treasury GET "/v1/treasury/financial_accounts?limit=100" \
        | jq '[.data[] | select(.status == "open")] | length'
    )"

    if ! [[ "$open_count" =~ ^[0-9]+$ ]]; then
      return 1
    fi
    if (( open_count < 3 )); then
      return 0
    fi

    local close_candidate
    close_candidate="$(
      api_treasury GET "/v1/treasury/financial_accounts?limit=100" \
        | jq -r --arg protected "$protected_financial_account" --arg attempted "$attempted_close_ids" '
            .data
            | map(select(
                .status == "open"
                and .id != $protected
                and (.id as $id | ($attempted | split(",") | index($id) | not))
              ))
            | if length == 0 then "" else .[-1].id end
          '
    )"

    if [[ -z "$close_candidate" || "$close_candidate" == "null" ]]; then
      return 1
    fi

    local close_status
    close_status="$(
      # Stripe requires forwarding settings when closing an account with recent
      # activity so incoming/outgoing flows can drain into another open FA.
      api_treasury POST "/v1/treasury/financial_accounts/${close_candidate}/close" \
        -d "forwarding_settings[type]=financial_account" \
        -d "forwarding_settings[financial_account]=${protected_financial_account}" \
        | jq -r '.status // empty'
    )"
    if [[ "$close_status" != "closed" ]]; then
      attempted_close_ids="${attempted_close_ids},${close_candidate}"
      continue
    fi

    attempted_close_ids="${attempted_close_ids},${close_candidate}"
    echo "auto-closed treasury financial account ${close_candidate} to free capacity"
  done
}

print_env_summary

if [[ "$should_seed_dispute_txn" == "1" ]]; then
  if ! seed_issuing_dispute_transaction STRIPE_ISSUING_DISPUTE_TRANSACTION ||
    ! seed_issuing_dispute_transaction STRIPE_ISSUING_DISPUTE_OTHER_TRANSACTION; then
    if [[ "$strict_dispute_target" == "1" || "$REQUIRE_ENV" == "1" ]]; then
      echo "failed to auto-seed an issuing dispute transaction" >&2
      exit 1
    fi
    echo "warning: unable to auto-seed issuing dispute transaction; continuing" >&2
  fi
fi

if [[ "$should_seed_credit_note_invoice" == "1" ]]; then
  if ! seed_credit_note_invoice; then
    if [[ "$strict_credit_note_target" == "1" || "$REQUIRE_ENV" == "1" ]]; then
      echo "failed to auto-seed a credit note invoice" >&2
      exit 1
    fi
    echo "warning: unable to auto-seed credit note invoice; continuing" >&2
  fi
fi

if [[ "$should_prepare_treasury_financial_account" == "1" ]]; then
  if ! prepare_treasury_financial_account_capacity; then
    if [[ "$strict_treasury_financial_account_target" == "1" || "$REQUIRE_ENV" == "1" ]]; then
      echo "failed to prepare treasury financial account capacity" >&2
      exit 1
    fi
    echo "warning: unable to prepare treasury financial account capacity; continuing" >&2
  fi
fi

case "$KIND" in
  "")
    packages=(./internal/acctest/...)
    ;;
  action)
    packages=(./internal/acctest/cases/action/...)
    ;;
  ephemeral)
    packages=(./internal/acctest/cases/ephemeral/...)
    ;;
  managed|resource)
    packages=(./internal/acctest/cases/managed/...)
    ;;
  *)
    echo "unsupported kind: $KIND" >&2
    exit 1
    ;;
esac

case_name_to_go_test_regex() {
  local case_name="$1"
  local kind="$2"
  local prefix=""
  case "$kind" in
    action)
      prefix="TestAccAction"
      ;;
    ephemeral)
      prefix="TestAccEphemeral"
      ;;
    managed|resource|*)
      prefix="TestAccManaged"
      ;;
  esac

  local camel=""
  local part
  IFS='_' read -r -a parts <<< "$case_name"
  for part in "${parts[@]}"; do
    camel+="${part^}"
  done
  printf '%s%s$' "$prefix" "$camel"
}

join_case_regex() {
  local kind="$1"
  shift
  local -a regexes=()
  local case_name
  for case_name in "$@"; do
    regexes+=("$(case_name_to_go_test_regex "$case_name" "$kind")")
  done
  local IFS='|'
  printf '%s' "${regexes[*]}"
}

log_path="$(mktemp "${TMPDIR:-/tmp}/terraform-acceptance-log.XXXXXX")"
echo "==> Running Terraform acceptance tests"
echo "    log: ${log_path}"

set +e
run_regex='TestAcc|TestActionFixture'
if [[ "${#CASE_NAMES[@]}" -gt 0 ]]; then
  run_regex="$(join_case_regex "$KIND" "${CASE_NAMES[@]}")"
fi
go test "${packages[@]}" -run "$run_regex" -count=1 -v 2>&1 | tee "$log_path"
test_status="${PIPESTATUS[0]}"
set -e

print_test_summary "$log_path" "$test_status"
