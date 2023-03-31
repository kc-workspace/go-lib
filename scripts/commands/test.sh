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
  local args=("-r")

  args+=("--timeout=5m")
  args+=("--randomize-all" "--randomize-suites")
  args+=("--keep-going" "--race" "--vet=")
  args+=(
    "--output-dir=reports"
    "--json-report=test-results.json"
    "--junit-report=test-results.xml"
  )

  if test -n "$COVER"; then
    args+=(
      "--cover"
      "--coverprofile=coverage.out"
      "--covermode=atomic"
    )
  fi

  if test -n "$DRY"; then
    args+=("--dry-run")
  fi

  if test -n "$DEBUG"; then
    args+=("-vv" "--trace")
  fi

  if test -n "$CI"; then
    args+=("-procs=1" "--compilers=1")
    args+=("--no-color")
  else
    args+=("-p")
  fi

  _cmd go run github.com/onsi/ginkgo/v2/ginkgo \
    run "${args[@]}"
}

main "$@"
