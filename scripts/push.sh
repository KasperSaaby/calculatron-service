# Set your GCP project ID
PROJECT_ID=$(gcloud config get-value project)
# Choose a region (e.g., us-central1, europe-west1)
REGION="europe-west1" # Using current location's approximate region
SERVICE_NAME="calculatron-service"

# Build the Docker image locally
docker buildx build --platform linux/amd64 -t gcr.io/${PROJECT_ID}/${SERVICE_NAME} ../

# Push the image to Google Container Registry (GCR) or Artifact Registry
# For new projects, Artifact Registry is generally preferred.
# If using Artifact Registry for the first time in your project/region:
# gcloud artifacts repositories create docker --repository-format=docker --location=${REGION}
docker push gcr.io/${PROJECT_ID}/${SERVICE_NAME}
