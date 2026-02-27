#!/bin/bash

REPO_URL="https://github.com/Ashking-tech/git-Visualise"
BINARY_NAME="git-visualize"

echo "Installing ${BINARY_NAME}..."

# Get latest release URL
DOWNLOAD_URL=$(curl -sSL "${REPO_URL}/releases/latest/download/${BINARY_NAME}")

# Download binary
curl -sSL "${DOWNLOAD_URL}" -o "/tmp/${BINARY_NAME}"

# Move to system path
sudo mv "/tmp/${BINARY_NAME}" "/usr/local/bin/${BINARY_NAME}"
sudo chmod +x "/usr/local/bin/${BINARY_NAME}"

echo "Installed! Run: git-visualize -email your@email.com"
