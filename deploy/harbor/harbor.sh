#!/bin/bash

kubectl create namespace harbor
helm install --namespace harbor myhrbr harbor/harbor --set expose.type=clusterIP --set expose.tls.enabled=false

helm uninstall --namespace harbor myhrbr
kubectl delete namespace harbor


