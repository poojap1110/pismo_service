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
  name: ${1}-topup-disable
  namespace: default
  labels:
    app: ${1}-topup-disable
spec:
  schedule: 25 10 * * *
  concurrencyPolicy: Allow
  suspend: false
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: ${1}-topup-disable
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
                  echo "Starting Disable Cron Call"
                  curl -X POST http://platform-utilities-sgprefund-service-nodeport:8080/api/v1/topup -H 'Content-Type: application/json' -d '{"scope":"97", "topup_status":0}'
                  echo "End Disable Cron call"
              imagePullPolicy: Always
          restartPolicy: OnFailure
          terminationGracePeriodSeconds: 30
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 1
EOF