#!/bin/bash

# Variables
DOCKER_USERNAME="nicholasliem"
IMAGE_TAG="latest"
SERVICES=("client-service" "grand-oak-service" "pine-valley-service")

echo "Logging into Docker Hub..."
docker login || { echo "Docker login failed"; exit 1; }

for SERVICE in "${SERVICES[@]}"; do
  echo "Building image for $SERVICE..."
  
  if [ -d "$SERVICE" ]; then
    docker build -t "$DOCKER_USERNAME/$SERVICE:$IMAGE_TAG" "$SERVICE" || { echo "Failed to build $SERVICE"; exit 1; }
  else
    echo "Directory $SERVICE does not exist. Skipping..."
    continue
  fi

  echo "Pushing image for $SERVICE..."
  docker push "$DOCKER_USERNAME/$SERVICE:$IMAGE_TAG" || { echo "Failed to push $SERVICE"; exit 1; }
done

echo "All images built and pushed successfully!"
