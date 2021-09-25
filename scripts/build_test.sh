#!/usr/bin/env bash
# exit on error

set -eu

cd ${APP_DIR}


go test -failfast ./ -v --cover

go test -failfast ./app/ -v --cover

go test -failfast ./app/resource/api/*_test.go -v --cover
