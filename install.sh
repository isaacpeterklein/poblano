#!/bin/sh
set -e

REPO="isaacpeterklein/poblano"
BIN="poblano"
INSTALL_DIR="/usr/local/bin"

# Detect OS
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
  Darwin)
    case "$ARCH" in
      arm64) SUFFIX="macos-arm64" ;;
      x86_64) SUFFIX="macos-amd64" ;;
      *) echo "Unsupported architecture: $ARCH" && exit 1 ;;
    esac
    ;;
  Linux)
    case "$ARCH" in
      aarch64|arm64) SUFFIX="linux-arm64" ;;
      x86_64) SUFFIX="linux-amd64" ;;
      *) echo "Unsupported architecture: $ARCH" && exit 1 ;;
    esac
    ;;
  *)
    echo "Unsupported OS: $OS"
    echo "Windows users: download from https://github.com/$REPO/releases"
    exit 1
    ;;
esac

# Get latest release version
VERSION="$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | sed -E 's/.*"([^"]+)".*/\1/')"

if [ -z "$VERSION" ]; then
  echo "Could not find latest release. Check https://github.com/$REPO/releases"
  exit 1
fi

URL="https://github.com/$REPO/releases/download/$VERSION/poblano-$SUFFIX"

echo "Installing Poblano $VERSION ($SUFFIX)..."
curl -fsSL "$URL" -o "/tmp/poblano"
chmod +x "/tmp/poblano"

# Move to install dir (may need sudo)
if [ -w "$INSTALL_DIR" ]; then
  mv "/tmp/poblano" "$INSTALL_DIR/$BIN"
else
  echo "Needs permission to write to $INSTALL_DIR, trying sudo..."
  sudo mv "/tmp/poblano" "$INSTALL_DIR/$BIN"
fi

echo ""
echo "Poblano installed! Try:"
echo "  poblano new mysite"
echo "  poblano build"
echo "  poblano serve"
