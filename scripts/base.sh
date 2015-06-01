#!/usr/bin/env bash

fancy_echo() {
  printf "\n%b\n" "$1"
}

trap 'ret=$?; test $ret -ne 0 && printf "failed\n\n" >&2; exit $ret' EXIT

set -e

fancy_echo "Upgrade all packages ..."
  sudo apt-get update

fancy_echo "Updating system packages ..."
  if command -v aptitude >/dev/null; then
    fancy_echo "Using aptitude ..."
  else
    fancy_echo "Installing aptitude ..."
    sudo apt-get install -y aptitude
  fi

  sudo aptitude update

fancy_echo "Installing curl, for making web requests ..."
  sudo aptitude install -y curl

fancy_echo "Installing vim, for editing files on the server ..."
  sudo aptitude install -y vim

fancy_echo "Installing git, for source control management ..."
  sudo aptitude install -y git

fancy_echo "Installing Redis, a good key-value database ..."
  sudo aptitude install -y redis-server

