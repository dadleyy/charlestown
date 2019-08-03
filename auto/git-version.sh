#!/bin/bash

function gitversion() {
  local revparse=`git rev-parse HEAD 2> /dev/null`
  local tag=`git describe --abbrev=0 2> /dev/null`
  local short=${revparse:0:10}
  local modifier=""

  if [[ "" != "$(git status --porcelain 2> /dev/null)" ]]; then
    modifier="-develop"
  fi

  if [[ $tag != "" ]]; then
    echo "${short}+${tag}${modifier}"
    return
  fi

  echo "${short}${modifier}"
}

gitversion
