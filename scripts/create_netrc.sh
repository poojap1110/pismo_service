#!/usr/bin/env bash

# exit on error
set -e

if test "$#" -ne 2; then
    echo "Illegal number of parameters"
fi

cat <<EOF
machine bitbucket.org
        login ${1}
        password ${2}
EOF
