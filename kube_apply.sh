#!/bin/bash

YAML_FILES=("config/client-service.yaml" "config/grand-oak-service.yaml" "config/pine-valley-service.yaml" "config/istio-gateway.yaml")

for YAML in "${YAML_FILES[@]}"; do
  echo "Applying $YAML..."
  kubectl apply -f "$YAML" || { echo "Failed to apply $YAML"; exit 1; }
done

echo "All Kubernetes YAMLs applied successfully!"
