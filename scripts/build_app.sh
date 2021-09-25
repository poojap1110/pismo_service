#!/usr/bin/env bash
# exit on error

set -eu

cd ${APP_DIR}

echo ">>>>>>>> Setting up netrc..."
sh scripts/create_netrc.sh ${NETRC_LOGIN} ${NETRC_PASSWORD} > /root/.netrc

echo ">>>>>>>> Installing Go packages..."
sh packages.sh

echo ">>>>>>>> Building the Go binary..."
go build -o $GOPATH/bin/${APP_NAME} .
mv $GOPATH/bin/${APP_NAME} .