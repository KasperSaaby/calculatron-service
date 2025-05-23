#!/bin/bash

# Exit on error
set -e

# Check if required commands are available
command -v docker >/dev/null 2>&1 || { echo "Error: docker is required but not installed."; exit 1; }
command -v gcloud >/dev/null 2>&1 || { echo "Error: gcloud is required but not installed."; exit 1; }

# Set your GCP project ID
PROJECT_ID=$(gcloud config get-value project)
if [ -z "$PROJECT_ID" ]; then
    echo "Error: Could not get GCP project ID. Make sure you're logged in and have a project set."
    exit 1
fi

# Choose a region (e.g., us-central1, europe-west1)
REGION="europe-west1" # Using current location's approximate region
SERVICE_NAME="calculatron-service"

echo "Building Docker image..."
# Build the Docker image locally
docker buildx build --platform linux/amd64 -t gcr.io/${PROJECT_ID}/${SERVICE_NAME} ../ || {
    echo "Error: Docker build failed"
    exit 1
}

echo "Pushing image to GCR..."
# Push the image to Google Container Registry (GCR) or Artifact Registry
# For new projects, Artifact Registry is generally preferred.
# If using Artifact Registry for the first time in your project/region:
# gcloud artifacts repositories create docker --repository-format=docker --location=${REGION}
docker push gcr.io/${PROJECT_ID}/${SERVICE_NAME} || {
    echo "Error: Docker push failed"
    exit 1
}

echo "Successfully built and pushed image to gcr.io/${PROJECT_ID}/${SERVICE_NAME}"
