#!/usr/bin/env bash

# exit on error
set -eu

if test "$#" -ne 4; then
    echo "Illegal number of parameters"
fi

if [ "${4}" = "external-loadbalancer" ]
then

cat <<EOF
---
kind: Service
apiVersion: v1
metadata:
  name: ${1}-${4}
  namespace: default
  labels:
    app: ${1}
spec:
  ports:
  - protocol: TCP
    port: ${2}
    targetPort: ${3}
  selector:
    app: ${1}
  type: LoadBalancer
EOF

elif [ "${4}" = "internal-loadbalancer" ]
then

cat <<EOF
---
kind: Service
apiVersion: v1
metadata:
  name: ${1}-${4}
  namespace: default
  labels:
    app: ${1}
  annotations:
    service.beta.kubernetes.io/aws-load-balancer-internal: 0.0.0.0/0
spec:
  ports:
  - protocol: TCP
    port: ${2}
    targetPort: ${3}
  selector:
    app: ${1}
  type: LoadBalancer
EOF

elif [ "${4}" = "nodeport" ]
then

cat <<EOF
---
kind: Service
apiVersion: v1
metadata:
  name: ${1}-${4}
  namespace: default
  labels:
    app: ${1}
spec:
  ports:
  - protocol: TCP
    port: ${2}
    targetPort: ${3}
  selector:
    app: ${1}
  type: NodePort
EOF

else

cat <<EOF
---
kind: Service
apiVersion: v1
metadata:
  name: ${1}-clusterip
  namespace: default
  labels:
    app: ${1}
spec:
  ports:
  - protocol: TCP
    port: ${2}
    targetPort: ${3}
  selector:
    app: ${1}
EOF

fi