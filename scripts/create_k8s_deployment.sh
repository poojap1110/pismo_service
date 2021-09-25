#!/usr/bin/env bash

# exit on error
set -eu

if test "$#" -ne 9; then
    echo "Illegal number of parameters"
fi

cat <<EOF
---
kind: Deployment
apiVersion: extensions/v1beta1
metadata:
  name: ${1}
  namespace: default
  labels:
    app: ${1}
spec:
  replicas: ${9}
  selector:
    matchLabels:
      app: ${1}
  template:
    metadata:
      creationTimestamp: 
      labels:
        app: ${1}
    spec:
      volumes:
      - name: ${1}-log
        emptyDir: {}
      containers:
      - name: ${1}
        image: ${2}/${1}:${3}
        ports:
        - containerPort: ${4}
          protocol: TCP
        readinessProbe:
          tcpSocket:
            port: ${4}
          initialDelaySeconds: 5
          periodSeconds: 10
        livenessProbe:
          tcpSocket:
            port: ${4}
          initialDelaySeconds: 15
          periodSeconds: 20
        envFrom:
        - secretRef:
            name: ${1}-credentials
        resources:
          limits:
            cpu: ${5}
            memory: ${6}
          requests:
            cpu: ${7}
            memory: ${8}
        volumeMounts:
        - name: ${1}-log
          mountPath: "/log/server"
        imagePullPolicy: IfNotPresent
      - name: ${1}-logger
        image: busybox
        args:
        - "/bin/sh"
        - "-c"
        - tail -f /log/server/${1}*.log
        resources: {}
        volumeMounts:
        - name: ${1}-log
          mountPath: "/log/server"
        terminationMessagePath: "/dev/termination-log"
        terminationMessagePolicy: File
        imagePullPolicy: IfNotPresent
EOF