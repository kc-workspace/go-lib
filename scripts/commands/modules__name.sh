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

__is_github() {
  test -n "$CI" && test -n "$GITHUB_ACTIONS"
}

main() {
  if __is_github; then
    local modules=()
    while IFS= read -r line; do
      modules+=("$line")
    done < <(go list -f '{{.Dir}}' -m)

    printf "%s=%s\n" "names" "${modules[*]}"
  else
    go list -f '{{.Dir}}' -m
  fi
}

if __is_github; then
  main "$@" >>"$GITHUB_OUTPUT"
else
  main "$@"
fi
