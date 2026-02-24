#!/bin/sh
set -e

case "$(uname -s)" in
  Darwin) OS="darwin" ;;
  Linux)  OS="linux" ;;
  *) echo "Unsupported OS"; exit 1 ;;
esac

case "$(uname -m)" in
  x86_64) ARCH="amd64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Unsupported arch"; exit 1 ;;
esac

# resolve latest release version from GitHub
TAG=$(curl -fsSL https://api.github.com/repos/iamtraction/sage/releases/latest \
  | grep '"tag_name":' | cut -d '"' -f4)
VERSION="${TAG#v}"
ARTIFACT="sage_${VERSION}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/iamtraction/sage/releases/download/${TAG}/${ARTIFACT}"

INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"

# download and extract binary
echo "Installing sage ${VERSION} to ${INSTALL_DIR}"
curl -fsSL "$URL" | tar -xz -C /tmp
mv /tmp/sage "${INSTALL_DIR}/sage"
chmod +x "${INSTALL_DIR}/sage"

echo "Installed."
echo "Ensure ${INSTALL_DIR} is in your PATH."
