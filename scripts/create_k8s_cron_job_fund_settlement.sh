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
  name: ${1}-fund-settlement
  namespace: default
  labels:
    app: ${1}-fund-settlement
spec:
  schedule: 30 10 * * *
  concurrencyPolicy: Allow
  suspend: false
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: ${1}-fund-settlement
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
                  echo "Starting Fund Settlement Cron Call"
                  curl -X POST http://platform-utilities-sgprefund-service-nodeport:8080/api/v1/fund_settlement -H 'Content-Type: application/json' -d '{}'
                  echo "End Fund Settlement Cron call"
              imagePullPolicy: Always
          restartPolicy: OnFailure
          terminationGracePeriodSeconds: 30
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 1
EOF