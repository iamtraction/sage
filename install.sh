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
VERSION=$(curl -fsSL https://api.github.com/repos/iamtraction/sage/releases/latest \
  | grep '"tag_name":' | cut -d '"' -f4)
ARTIFACT="git-sage_${VERSION}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/iamtraction/sage/releases/download/${VERSION}/${ARTIFACT}"

INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"

# download and extract binary
echo "Installing git-sage ${VERSION} to ${INSTALL_DIR}"
curl -fsSL "$URL" | tar -xz -C /tmp
mv /tmp/git-sage "${INSTALL_DIR}/git-sage"
chmod +x "${INSTALL_DIR}/git-sage"

echo "Installed."
echo "Ensure ${INSTALL_DIR} is in your PATH."
