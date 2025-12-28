#!/usr/bin/env bash

set -e

echo "🧹 Uninstalling WTF..."

# Remove binary
if [ -f /usr/local/bin/wtf ]; then
    sudo rm -f /usr/local/bin/wtf
    echo "✔ Removed /usr/local/bin/wtf"
else
    echo "ℹ /usr/local/bin/wtf not found"
fi

# Remove shared data
if [ -d /usr/local/share/wtf ]; then
    sudo rm -rf /usr/local/share/wtf
    echo "✔ Removed /usr/local/share/wtf"
else
    echo "ℹ /usr/local/share/wtf not found"
fi

# Remove config dir

if [ -d /etc/wtf ]; then
    sudo rm -rf /etc/wtf
    echo "✔ Removed /etc/wtf"
else
    echo "ℹ /etc/wtf not found"
fi

# Remove local project directory (relative)
if [ -d "../wtf" ]; then
    rm -rf ../wtf
    echo "✔ Removed local ./wtf directory"
else
    echo "ℹ Local ./wtf directory not found"
fi

echo "✅ WTF successfully removed"

