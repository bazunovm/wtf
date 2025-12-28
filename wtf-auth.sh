#!/usr/bin/env bash

KEY="$1"
CONFIG_DIR="/etc/wtf"
CONFIG_FILE="$CONFIG_DIR/config"

if [ -z "$KEY" ]; then
  echo "Usage: sudo wtf-auth <API_KEY>"
  exit 1
fi

if [ "$EUID" -ne 0 ]; then
  echo "Please run as root (use sudo)"
  exit 1
fi

mkdir -p "$CONFIG_DIR"

cat > "$CONFIG_FILE" <<EOF
WTF_AI_API_KEY=$KEY
EOF

chmod 600 "$CONFIG_FILE"

echo "✅ API key saved to $CONFIG_FILE"

