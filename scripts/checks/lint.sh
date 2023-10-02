#!/bin/bash
# vim: ai:ts=8:sw=8:noet
# Check this lib
# Intended to be run from local machine or CI
set -eufo pipefail
export SHELLOPTS	# propagate set to children by default
IFS=$'\t\n'

# Check required commands are in place
command -v golangci-lint >/dev/null 2>&1 || { echo 'please install golangci-lint or use image that has it'; exit 1; }

golangci-lint run --disable-all \
  --deadline 20m0s \
  --skip-files .*.autogen.go.* \
  -e composites \
  -E goimports \
  -E golint \
  -E govet
