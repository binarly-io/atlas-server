#!/bin/bash

NAMESPACE="${NAMESPACE:-minio}"

source "minio-instance-params.sh"

watch "kubectl -n ${NAMESPACE} get MinIOInstance,all,pvc,pv"

