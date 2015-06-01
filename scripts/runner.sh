#!/usr/bin/env bash

sudo -u vagrant -H bash -c "~/base.sh"
sudo -u vagrant -H bash -c "~/docker.sh"
sudo -u vagrant -H bash -c "~/golang.sh"
sudo -u vagrant -H bash -c "~/redb-database.sh"

