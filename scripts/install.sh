#!/usr/bin/env bash
set -euo pipefail

REPO="flarebyte/clingy-code-detective"
BINARY_NAME="clingy"
VERSION="${1:-latest}"

# Detect OS and ARCH
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  aarch64 | arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Resolve version
if [[ "$VERSION" == "latest" ]]; then
  VERSION=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep -oP '"tag_name":\s*"\K(.*)(?=")')
fi

FILENAME="clingy-${OS}-${ARCH}"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${FILENAME}"

echo "Downloading $FILENAME from $URL..."
curl -L -o "$BINARY_NAME" "$URL"
chmod +x "$BINARY_NAME"
sudo mv "$BINARY_NAME" /usr/local/bin

echo "✅ Installed $BINARY_NAME $VERSION to /usr/local/bin"
