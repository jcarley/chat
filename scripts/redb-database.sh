#!/usr/bin/env bash

source /etc/lsb-release && echo "deb http://download.rethinkdb.com/apt $DISTRIB_CODENAME main" | sudo tee /etc/apt/sources.list.d/rethinkdb.list
wget -qO- http://download.rethinkdb.com/apt/pubkey.gpg | sudo apt-key add -
sudo apt-get update
sudo apt-get install rethinkdb -y

PWD=$(pwd)
sudo cp ${PWD}/files/rethinkdb/instance1.conf /etc/rethinkdb/instances.d/instance1.conf
sudo /etc/init.d/rethinkdb restart

