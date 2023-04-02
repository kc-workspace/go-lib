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
  local cmd="golangci-lint"
  local args=()
  if command -v "$cmd" >/dev/null; then
    args+=("--fix")
    test -n "$DEBUG" && args+=("-v")

    if [ "$#" -gt 0 ]; then
      args+=("$@")
    fi

    # shellcheck disable=2046
    "$cmd" run "${args[@]}" $("$__SCRIPT/main.sh" modules name)
  else
    _throw 2 "cannot found commandline: %s" \
      "$cmd"
  fi
}

main "$@"
