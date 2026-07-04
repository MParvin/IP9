#!/usr/bin/env sh
set -eu

ROOT="$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)"
DATA_DIR="${ROOT}/data"
BASE_URL="https://github.com/lionsoul2014/ip2region/raw/master/data"

mkdir -p "${DATA_DIR}"

for file in ip2region_v4.xdb ip2region_v6.xdb; do
  echo "Downloading ${file}..."
  curl -fsSL -o "${DATA_DIR}/${file}" "${BASE_URL}/${file}"
done

echo "ip2region databases saved to ${DATA_DIR}"
