#!/bin/bash

# set -e stops the execution of a script if a command or pipeline has an error
set -e

# The name of your Docker image
IMAGE_NAME="kajo-api"

# The tag for your Docker image
IMAGE_TAG="dev-latest"

# The path to your Kubernetes manifest
MANIFEST_PATH="./deployment/k8s/dev.yaml"

# Build the Docker image
docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .

# Uncomment the next line if you need to push the image to a Docker registry
# docker push ${IMAGE_NAME}:${IMAGE_TAG}

# Apply the Kubernetes manifest
kubectl apply -f ${MANIFEST_PATH}