#!/usr/bin/env bash

# install dependencies
sudo apt-get update
sudo apt-get install -y gcc libc6-dev make --no-install-recommends
sudo rm -rf /var/lib/apt/lists/*

# download and extract
export GOLANG_VERSION="1.4.2"
sudo curl -sSL https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz | sudo tar -v -C /usr/local -xz

# compile
cd /usr/local/go/src && sudo ./make.bash --no-clean 2>&1

# su - vagrant
# echo 'vagrant' | sudo -E -S bash '{{.Path}}'

run_as_user=vagrant
run_as_user_home="/home/${run_as_user}"

# put the go tools on the path
if ! grep -qs "GOLANG_VERSION" $run_as_user_home/.bashrc; then
  printf 'export GOLANG_VERSION="1.4.2"\n' >> $run_as_user_home/.bashrc
  printf 'export PATH="/usr/local/go/bin:$PATH"\n' >> $run_as_user_home/.bashrc
fi

# this is the project directory
# mkdir -p $run_as_user_home/go/src $run_as_user_home/go/bin && chmod -R 777 $run_as_user_home/go

# export the GOPATH for projects
if ! grep -qs "GOPATH" $run_as_user_home/.bashrc; then
  printf 'export GOPATH="$HOME/go"\n' >> $run_as_user_home/.bashrc
  printf 'export PATH="$HOME/go/bin:$PATH"\n' >> $run_as_user_home/.bashrc
fi
