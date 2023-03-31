#!/usr/bin/env bash

__TMP="$PWD"
cd "$(dirname "$0")/.." || exit 1
__ROOT="$PWD"
__SCRIPT="$__ROOT/scripts"

__SEPARATOR="__"
__COMMAND_DIR="$__SCRIPT/commands"

__EXECUTE_FILE=""

main() {
  local index="$#"
  local args=("$@") _args=()

  # support directory as first parameters
  if test -d "$__ROOT/${args[0]}"; then
    # $PWD changed from $__ROOT to input directory
    cd "$__ROOT/${args[0]}" || exit 1
    # remove first argument
    args=("${args[@]:1}")
    # remove first index
    ((index--))
  fi

  __EXECUTE_FILE="${args[*]}"
  __EXECUTE_FILE="${__EXECUTE_FILE// /$__SEPARATOR}.sh"

  while true; do
    [ $index -le 0 ] &&
      break
    _args=("${args[@]:$index}")

    _debug "checking(%d) $%s" \
      "$index" "$__EXECUTE_FILE"
    __execute "${_args[@]}"

    __EXECUTE_FILE="${__EXECUTE_FILE%"$__SEPARATOR"*}.sh"
    ((index--))
  done

  __EXECUTE_FILE="_catch.sh"
  __execute "${args[@]}"
}

_cmd() {
  if test -n "$DRY"; then
    printf "[DRYRN] %s\n" "$*" >&2
    return 0
  fi
  if test -n "$CI"; then
    printf "[DEBUG] %s\n" "$*"
  fi

  "$@"
}

_debug() {
  test -z "$DEBUG" && return 0

  local format="$1"
  shift
  local args=("$@")

  # shellcheck disable=SC2059
  printf "[DEBUG] $format\n" "${args[@]}"
}

_error() {
  local format="$1"
  shift
  local args=("$@")

  # shellcheck disable=SC2059
  printf "[ERROR] $format\n" "${args[@]}" >&2
}

_throw() {
  local code="$1"
  shift
  _error "$@"
  exit "$code"
}

__execute() {
  local path="$__COMMAND_DIR/$__EXECUTE_FILE"
  if test -f "$path"; then
    _debug "executing %s '%s'" \
      "$path" "$*"

    # shellcheck disable=SC1090
    source "$path" "$@"
    exit "$?"
  fi
}

main "$@"

cd "$__TMP" || exit 1
unset __TMP __SCRIPT __ROOT
unset __SEPARATOR
unset __COMMAND_DIR
unset __EXECUTE_FILE
