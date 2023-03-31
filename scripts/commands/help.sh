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
  local path name
  echo "list of available commands:"
  for path in "$__COMMAND_DIR"/*; do
    name="$(basename "$path")"
    [[ "$name" =~ ^_ ]] && continue
    name="${name/\.sh/}"
    name="${name//__/ }"

    echo "- $name"
  done
}

main "$@"
