#!/usr/bin/env bash
set -euo pipefail
if [ $# -ne 2 ]; then
  echo "usage: $0 <module_path> <plugin_name>" >&2
  exit 1
fi
MOD="$1"
PLG="$2"

# replace placeholders
git ls-files | while read -r f; do
  sed -i \
    -e "s|{{MODULE_PATH}}|$MOD|g" \
    -e "s|{{plugin}}|$PLG|g" \
    "$f"
done

# move folders
git mv "cmd/{{plugin}}" "cmd/$PLG"
git mv "pkg/{{plugin}}" "pkg/$PLG"
