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
  go build "github.com/kc-workspace/go-lib/..."
}

main "$@"
