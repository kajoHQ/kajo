#!/bin/bash

# set -e stops the execution of a script if a command or pipeline has an error
set -e

IMAGE_NAME="kajo-api"

IMAGE_TAG="dev-latest"

MANIFEST_PATH="./deployment/k8s/dev.yaml"

kubectl delete -f ${MANIFEST_PATH}