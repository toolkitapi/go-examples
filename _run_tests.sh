#!/usr/bin/env bash
# Test runner for all Go examples.
# Runs each .go file via `go run`, handles webhook chaining.
# Requires: go, jq
# Exit code: 0 = all enabled tests pass; 1 = one or more failures.
#
# Must be run from examples/go/ (where go.mod lives), or it will cd there.
# Usage: export TOOLKITAPI_KEY=tk_live_...; bash _run_tests.sh

set -uo pipefail

API_KEY="${TOOLKITAPI_KEY:-}"
if [ -z "$API_KEY" ]; then
  echo "Error: TOOLKITAPI_KEY is not set" >&2
  exit 1
fi

# Ensure we run from the directory containing go.mod
cd "$(dirname "${BASH_SOURCE[0]}")"

PASS=0
FAIL=0
SKIP=0
declare -a FAILURES=()

# ── Helpers ───────────────────────────────────────────────────────────────────

sep() {
  echo ""
  echo "============================================================"
  echo "  $1"
  echo "============================================================"
}

record_pass() { printf "  ✅ PASS  %-35s\n" "$2"; ((PASS++)); }
record_fail() { printf "  ❌ FAIL  %-35s %s\n" "$2" "${3:0:120}"; ((FAIL++)); FAILURES+=("$1/$2: ${3:-}"); }
record_skip() { printf "  ⚠️  SKIP  %-35s %s\n" "$2" "${3:-}"; ((SKIP++)); }

run_toolkit() {
  local toolkit="$1"; shift
  sep "${toolkit^^}"
  for f in "$@"; do
    local path="$toolkit/$f"
    if [ ! -f "$path" ]; then
      record_fail "$toolkit" "$f" "FILE MISSING"
      continue
    fi
    local out
    if out=$(timeout 60 go run "$path" 2>&1); then
      record_pass "$toolkit" "$f"
    else
      record_fail "$toolkit" "$f" "$(echo "$out" | head -1 | cut -c1-120)"
    fi
  done
}

# ── Devtools ──────────────────────────────────────────────────────────────────
run_toolkit devtools \
  generate_uuid.go json_validate.go yaml_validate.go \
  regex_run.go cron_parse.go diff_text.go \
  slugify.go math_eval.go

# ── DNS ───────────────────────────────────────────────────────────────────────
run_toolkit dns \
  lookup.go lookup_all.go lookup_bulk.go whois.go \
  available.go domain_age.go certificate.go \
  typosquat.go propagation.go health.go

# ── Email ─────────────────────────────────────────────────────────────────────
run_toolkit email \
  validate_email.go normalize.go role_check.go \
  catch_all.go provider.go security.go \
  validate_batch.go spam_score.go

# ── Auth ──────────────────────────────────────────────────────────────────────
run_toolkit auth \
  generate_password.go password_strength.go \
  hash_password.go generate_key.go base64_encode.go \
  encrypt.go jwt_generate.go jwt_verify.go \
  totp_generate.go verify_password.go

# ── Barcode ───────────────────────────────────────────────────────────────────
sep "BARCODE (plan unavailable — expected skip)"
for f in types.go qr_generate.go qr_bulk.go generate.go qr_decode.go decode.go; do
  record_skip "barcode" "$f" "API key does not have barcode access"
done

# ── Geo ───────────────────────────────────────────────────────────────────────
run_toolkit geo \
  ip_lookup.go ip_threat.go country_info.go \
  timezone_convert.go distance.go phone_validate.go

# ── Textanalysis ──────────────────────────────────────────────────────────────
run_toolkit textanalysis \
  detect_language.go summarize.go word_frequency.go \
  readability_score.go text_similarity.go \
  profanity_filter.go data_mask.go

# ── Scrape ────────────────────────────────────────────────────────────────────
run_toolkit scrape \
  scrape.go seo_audit.go seo_keyword_density.go \
  seo_mobile_friendly.go seo_pagespeed.go

# ── Media ─────────────────────────────────────────────────────────────────────
run_toolkit media \
  youtube_transcript.go youtube_video.go \
  youtube_channel.go youtube_search.go

# ── Image ─────────────────────────────────────────────────────────────────────
run_toolkit image \
  metadata.go colors.go resize.go \
  compress.go strip_exif.go remove_background.go

# ── PDF ───────────────────────────────────────────────────────────────────────
run_toolkit pdf \
  text_extract.go metadata.go split.go \
  compress.go merge.go watermark.go

# ── Convert ───────────────────────────────────────────────────────────────────
run_toolkit convert \
  list_formats.go data.go markup.go \
  json_to_typescript.go document.go spreadsheet.go

# ── Webhook (chained) ────────────────────────────────────────────────────────
sep "WEBHOOK (chained)"

# 1. create_bin
BIN_ID=""
if out=$(timeout 60 go run webhook/create_bin.go 2>&1); then
  record_pass "webhook" "create_bin.go"
  BIN_ID=$(echo "$out" | jq -r '.bin.bin_id // .bin_id // empty' 2>/dev/null || true)
else
  record_fail "webhook" "create_bin.go" "$(echo "$out" | head -1)"
fi

# 2. get_bin
if [ -n "${BIN_ID:-}" ]; then
  TMP="webhook/_tmp_get_bin.go"
  sed "s/your-bin-id-here/${BIN_ID}/g" "webhook/get_bin.go" > "$TMP"
  if out=$(timeout 60 go run "$TMP" 2>&1); then
    record_pass "webhook" "get_bin.go"
  else
    record_fail "webhook" "get_bin.go" "$(echo "$out" | head -1)"
  fi
  rm -f "$TMP"
else
  record_skip "webhook" "get_bin.go" "skipped (no bin_id from create_bin)"
fi

# 3. list_requests
if [ -n "${BIN_ID:-}" ]; then
  TMP="webhook/_tmp_list_requests.go"
  sed "s/your-bin-id-here/${BIN_ID}/g" "webhook/list_requests.go" > "$TMP"
  if out=$(timeout 60 go run "$TMP" 2>&1); then
    record_pass "webhook" "list_requests.go"
  else
    record_fail "webhook" "list_requests.go" "$(echo "$out" | head -1)"
  fi
  rm -f "$TMP"
else
  record_skip "webhook" "list_requests.go" "skipped (no bin_id from create_bin)"
fi

# 4. create_mock
MOCK_ID=""
if out=$(timeout 60 go run webhook/create_mock.go 2>&1); then
  record_pass "webhook" "create_mock.go"
  MOCK_ID=$(echo "$out" | jq -r '.mock.mock_id // .mock_id // empty' 2>/dev/null || true)
else
  record_fail "webhook" "create_mock.go" "$(echo "$out" | head -1)"
fi

# 5. hit_mock
if [ -n "${MOCK_ID:-}" ]; then
  TMP="webhook/_tmp_hit_mock.go"
  sed "s/your-mock-id-here/${MOCK_ID}/g" "webhook/hit_mock.go" > "$TMP"
  if out=$(timeout 60 go run "$TMP" 2>&1); then
    record_pass "webhook" "hit_mock.go"
  else
    record_fail "webhook" "hit_mock.go" "$(echo "$out" | head -1)"
  fi
  rm -f "$TMP"
else
  record_skip "webhook" "hit_mock.go" "skipped (no mock_id from create_mock)"
fi

# ── Summary ───────────────────────────────────────────────────────────────────
TOTAL=$((PASS + FAIL + SKIP))
echo ""
echo "============================================================"
echo "  RESULTS: ${PASS}/${TOTAL} passed, ${FAIL} failed, ${SKIP} skipped"
echo "============================================================"

if [ "${#FAILURES[@]}" -gt 0 ]; then
  echo ""
  echo "Failed:"
  for entry in "${FAILURES[@]}"; do
    echo "  $entry"
  done
  exit 1
fi

exit 0
