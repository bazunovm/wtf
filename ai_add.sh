#!/bin/bash
if [ -z "$1" ]; then
  echo "Usage: source ai_add.sh <API_KEY>"
  return 1 2>/dev/null || exit 1
fi

export WTF_AI_API_KEY="$1"
echo "WTF_AI_API_KEY set"
