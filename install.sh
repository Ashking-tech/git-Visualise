#!/bin/bash

BINARY_NAME="git-visualize"
REPO_URL="https://github.com/Ashking-tech/git-Visualise"

echo "Installing ${BINARY_NAME}..."

# Direct download URL from release
curl -sSL "${REPO_URL}/releases/download/v1.0.0/${BINARY_NAME}" -o "/tmp/${BINARY_NAME}"

# Move to system path
sudo mv "/tmp/${BINARY_NAME}" "/usr/local/bin/${BINARY_NAME}"
sudo chmod +x "/usr/local/bin/${BINARY_NAME}"

echo "Installed! Run: git-visualize -email your@email.com"
