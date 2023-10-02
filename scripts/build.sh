#!/bin/bash
# vim: ai:ts=8:sw=8:noet
# Build the image
# Intended to be run from CI or local
set -eufo pipefail
export SHELLOPTS	# propagate set to children by default
IFS=$'\t\n'

# Check required commands are in place
command -v docker >/dev/null 2>&1 || { echo 'Please install docker or use image that has it'; exit 1; }

if [[ -z "${TAG:-}" ]]; then
    TAG="latest"
fi
if [[ -z "${NAME:-}" ]]; then
    NAME="devs-crud"
fi
if [[ -z "${IMAGE:-}" ]]; then
    IMAGE="$NAME:local"
fi

echo "Building $NAME as '$IMAGE'"
docker build "$PWD" -t "$IMAGE" --build-arg BUILD_TAG="$TAG"
echo "Done"
