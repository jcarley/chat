#!/usr/bin/env bash


sudo apt-get update
sudo apt-get install -y wget --no-install-recommends
wget -qO- https://get.docker.com/ | sh
sudo rm -rf /var/lib/apt/lists/*

sudo usermod -aG docker $(whoami)


