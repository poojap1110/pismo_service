#!/usr/bin/env bash


set -ex

# Golang/dep SSH workaround
mkdir ~/.ssh
echo "Host *\n\tStrictHostKeyChecking no" > ~/.ssh/config
chmod 400 ~/.ssh/config

# Installs golang/dep and runs it to pull dependencies
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
dep init
dep ensure -v
dep status

# Remove workaround
rm ~/.ssh/config