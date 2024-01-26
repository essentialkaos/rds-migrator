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
  if [[ -z "$1" ]] ; then
    show "${CL_BOLD}Usage:${CL_NORM} ./rds-recreate.sh main-dir"
    exit 0
  fi

  checkMainDir "$1"
  recreate "$1"
}

checkMainDir() {
  local dir="$1"

  if [[ ! -d "$1" ]] ; then
    error "There is no directory $1"
    exit 1
  fi

  if [[ ! -d "$1/meta" ]] ; then
    error "There is no directory $1/meta"
    exit 1
  fi

  if [[ ! -d "$1/log" ]] ; then
    error "There is no directory $1/log"
    exit 1
  fi

  if [[ ! -d "$1/data" ]] ; then
    error "There is no directory $1/data"
    exit 1
  fi
}

recreate() {
  local dir="$1"
  local meta id

  show ""

  for meta in "$dir"/meta/* ; do
    id=$(basename "$meta")

    if [[ -d "$dir/data/$id" && -d "$dir/log/$id" ]] ; then
      continue
    fi

    mkdir -p "$dir/data/$id"
    mkdir -p "$dir/log/$id"

    touch "$dir/log/$id/redis.log"

    chown -R redis:redis "$dir/data/$id" "$dir/log/$id"

    show "${CL_GREEN}✔ ${CL_NORM} $id recreated"
  done

  show ""
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
