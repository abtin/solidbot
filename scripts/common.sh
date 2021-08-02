#!/bin/sh

debug=1

fail() {
  echo $1
  exit 1
}

bye() {
  echo $1
  exit
}

log() {
  if ( debug eq 1 )
  then
    echo $1
  fi
}