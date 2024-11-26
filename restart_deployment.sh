#!/bin/bash

DEPLOYMENTS=("client-service" "grand-oak-service" "pine-valley-service")

restart_deployments() {
  for DEPLOYMENT in "${DEPLOYMENTS[@]}"; do
    echo "Restarting deployment: $DEPLOYMENT..."
    kubectl rollout restart deployment/$DEPLOYMENT || {
      echo "Failed to restart deployment: $DEPLOYMENT"; exit 1;
    }
    echo "Deployment $DEPLOYMENT restarted successfully!"
  done
}

restart_deployments
