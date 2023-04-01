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
  local cmd="cspell"
  local args=()
  if command -v "$cmd"; then
    args+=("$cmd")
  else
    args+=("npx" "$cmd")
  fi

  args+=(
    "lint"
    "--config"
    "$__ROOT/.github/linters/cspell.json"
    "$PWD"
  )

  _cmd "${args[@]}"
}

main "$@"
