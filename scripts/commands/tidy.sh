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
  while IFS= read -r module; do
    cd "$module" || exit 2
    go mod tidy
  done < <("$__SCRIPT/main.sh" modules name)
}

main "$@"
