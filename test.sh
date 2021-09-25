#!/usr/bin/env bash

go get github.com/axw/gocov/gocov

DBUSER="user"
DBPASSWORD="password"

set -xe
source env.local

go test ./app/resource/api/*_test.go -v -failfast
#go test ./modules/helper/*_test.go -v -failfast

# Code Coverage + Convert to XML
echo "mode: set" > coverage.out | grep -v mode: | sort -r | \
awk '{if($1 != last) {print $0;last=$1}}' >> coverage.out
go tool cover -func=coverage.out
gocov convert coverage.out | gocov-xml > coverage.xml
