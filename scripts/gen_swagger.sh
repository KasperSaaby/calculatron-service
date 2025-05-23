#!/bin/bash

# Exit on error
set -e

# Check if docker is available
command -v docker >/dev/null 2>&1 || { echo "Error: docker is required but not installed."; exit 1; }

# Check if swagger.yaml exists
if [ ! -f "../api/swagger.yaml" ]; then
    echo "Error: ../api/swagger.yaml not found"
    exit 1
fi

echo "Generating Swagger server code..."
docker run --rm -it \
    --user "$(id -u):$(id -g)" \
    -v "$HOME:$HOME" \
    -w "$PWD" \
    quay.io/goswagger/swagger generate server \
    --spec ../api/swagger.yaml \
    --name calculatron-service \
    --exclude-main \
    --target ../generated || {
    echo "Error: Failed to generate Swagger code"
    exit 1
}

echo "Successfully generated Swagger server code in ../generated"
