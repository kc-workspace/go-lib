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
  local name="$1"
  test -z "$name" && _throw 2 "name argument is required"

  _cmd go run github.com/onsi/ginkgo/v2/ginkgo \
    generate "$name"
}

main "$@"
