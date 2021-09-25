#!/usr/bin/env bash
# exit on error

set -eu

cd ${APP_DIR}

#run the migratoin script

go run migration/migration.go migrate