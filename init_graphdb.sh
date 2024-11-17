#!/bin/bash
set -x

GRAPHDB_URL="http://graphdb:7200"
REPOSITORY_NAME="ponegraph-music"
DATA_FILE="/data/final-rdf.ttl"
REPO_CONFIG_FILE="/data/repo-config.ttl"

repository_exists() {
  curl -sf "${GRAPHDB_URL}/rest/repositories/${REPOSITORY_NAME}" > /dev/null 2>&1
}

delete_repository() {
  echo "Deleting repository: ${REPOSITORY_NAME}"
  curl -X DELETE \
    "${GRAPHDB_URL}/rest/repositories/${REPOSITORY_NAME}"
  echo "Repository ${REPOSITORY_NAME} deleted successfully."
}

create_repository() {
  echo "Creating repository: ${REPOSITORY_NAME}"
  curl -X POST \
    "${GRAPHDB_URL}/rest/repositories" \
    -H "Content-Type: multipart/form-data" \
    -F "config=@${REPO_CONFIG_FILE}"
  echo "Repository ${REPOSITORY_NAME} created successfully."
}

import_data() {
  echo "Importing data from ${DATA_FILE} into repository: ${REPOSITORY_NAME}"
  curl -X POST \
    "${GRAPHDB_URL}/repositories/${REPOSITORY_NAME}/statements" \
    -H "Content-Type: application/x-turtle" \
    --data-binary "@${DATA_FILE}"
  echo "Data imported successfully."
}

# main
if repository_exists; then
  echo "Repository ${REPOSITORY_NAME} already exists."
  delete_repository
fi

if [ -f "${REPO_CONFIG_FILE}" ]; then
  create_repository
  if [ -f "${DATA_FILE}" ]; then
    import_data
  else
    echo "Data file ${DATA_FILE} not found. Skipping import."
  fi
else
  echo "Repository configuration file ${REPO_CONFIG_FILE} not found. Cannot create repository."
  exit 1
fi
