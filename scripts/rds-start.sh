#!/usr/bin/env bash

########################################################################################

NORM=0
BOLD=1
UNLN=4
RED=31
GREEN=32
YELLOW=33
BLUE=34
MAG=35
CYAN=36
GREY=37
DARK=90

CL_NORM="\e[0m"
CL_BOLD="\e[0;${BOLD};49m"
CL_UNLN="\e[0;${UNLN};49m"
CL_RED="\e[0;${RED};49m"
CL_GREEN="\e[0;${GREEN};49m"
CL_YELLOW="\e[0;${YELLOW};49m"
CL_BLUE="\e[0;${BLUE};49m"
CL_MAG="\e[0;${MAG};49m"
CL_CYAN="\e[0;${CYAN};49m"
CL_GREY="\e[0;${GREY};49m"
CL_DARK="\e[0;${DARK};49m"

################################################################################

main() {
  local dir="$1"
  local list="$2"

  if [[ $# -lt 2 || ! -d "$dir" || ! -f "$list" ]] ; then
    show "${CL_BOLD}Usage:${CL_NORM} ./rds-start.sh main-dir instance-list"
    exit 0
  fi

  checkMainDir "$dir"
  startInstances "$dir" "$list"
}

checkMainDir() {
  local dir="$1"

  if [[ ! -d "$1" ]] ; then
    error "There is no directory $1"
    exit 1
  fi

  if [[ ! -d "$1/pid" ]] ; then
    error "There is no directory $1/pid"
    exit 1
  fi
}

startInstances() {
  local dir="$1"
  local list="$2"

  local id

  while read -r id ; do

    if [[ -e "$dir/pid/$id.pid" ]] ; then
      continue
    fi

    cat ~/rds-su | rds start "$id"
    
    sleep 30

  done < <(awk 1 "$list")
}

show() {
  if [[ -n "$2" ]] ; then
    echo -e "\e[${2}m${1}\e[0m"
  else
    echo -e "$*"
  fi
}

error() {
  show "$@" $RED 1>&2
}

################################################################################

main "$@"
