#!/usr/bin/env bash

## Versions: v1.0.0
## Variables:
## - $__ROOT   - root directory
## - $__SCRIPT - scripts directory
## Functions:
## - _debug "$format" "${args[@]}"
## - _error "$format" "${args[@]}"
## - _throw 1 "$format" "${args[@]}"
## - _cmd echo "test"

main() {
  local is_coverage="${COVER:-$CI}"
  local is_coverage_report="${REPORT}"
  local is_dryrun="$DRY"
  local is_debug="$DEBUG"
  local is_ci="$CI"

  local args=("-r")

  args+=("--timeout=5m")
  args+=("--randomize-all" "--randomize-suites")
  args+=("--keep-going" "--race" "--vet=")
  args+=(
    "--output-dir=reports"
    "--json-report=test-results.json"
    "--junit-report=test-results.xml"
  )

  if test -n "$is_dryrun"; then
    args+=("--dry-run")
  fi

  if test -n "$is_debug"; then
    args+=("-vv" "--trace")
  fi

  if test -n "$is_ci"; then
    args+=("-procs=1" "--compilers=1")
    args+=("--no-color")
  else
    args+=("-p")
  fi

  if test -n "$is_coverage"; then
    args+=(
      "--cover"
      "--coverprofile=coverage.out"
      "--covermode=atomic"
    )
  fi

  if ! _cmd go run github.com/onsi/ginkgo/v2/ginkgo \
    run "${args[@]}"; then
    return $?
  fi

  local report="$PWD/reports/coverage.out"
  if test -n "$is_coverage" &&
    test -n "$is_coverage_report" &&
    test -f "$report"; then
    _cmd go tool cover \
      "-html=$report" \
      "-o=reports/coverage.html"
  fi
}

main "$@"
