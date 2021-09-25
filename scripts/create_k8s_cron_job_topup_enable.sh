#!/usr/bin/env bash

# exit on error
set -eu

if test "$#" -ne 9; then
    echo "Illegal number of parameters"
fi

cat <<EOF
---
kind: CronJob
apiVersion: batch/v1beta1
metadata:
  name: ${1}-topup-enable
  namespace: default
  labels:
    app: ${1}-topup-enable
spec:
  schedule: 35 10 * * *
  concurrencyPolicy: Allow
  suspend: false
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: ${1}-topup-enable
              image: yauritux/busybox-curl
              envFrom:
              - secretRef:
                  name: ${1}-credentials
              args:
                - "/bin/sh"
                - "-c"
              command:
                - "/bin/sh"
                - "-ec"
                - |
                  echo "Starting Enable Cron Call"
                  curl -X POST http://platform-utilities-sgprefund-service-nodeport:8080/api/v1/topup -H 'Content-Type: application/json' -d '{"scope":"99", "topup_status":1}'
                  echo "End Enable Cron call"
              imagePullPolicy: Always
          restartPolicy: OnFailure
          terminationGracePeriodSeconds: 30
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 1
EOF