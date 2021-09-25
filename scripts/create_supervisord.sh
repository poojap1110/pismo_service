#!/usr/bin/env bash

# exit on error
set -e

if test "$#" -ne 1; then
    echo "Illegal number of parameters"
fi

cat <<EOF
[supervisord]
nodaemon=true

[program:${1}]
command=${1}
autostart=true
autorestart=true
startretries=10
user=root
stderr_logfile=/log/server/${1}.err.log
stdout_logfile=/log/server/${1}.out.log
EOF