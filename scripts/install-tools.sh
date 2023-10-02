#!/bin/bash

# Install external tools
# Intended to be run from local machine or CI

set -eufo pipefail
IFS=$'\t\n'

# Check required commands are in place
command -v go > /dev/null 2>&1 || { echo 'please install go or use image that has it'; exit 1; }

GO111MODULE=off go get -u -v \
		golang.org/x/tools/cmd/goimports \
		golang.org/x/lint/golint \
		honnef.co/go/tools/cmd/staticcheck \
		github.com/kisielk/errcheck \
		github.com/kyoh86/richgo \
		github.com/vektra/mockery/.../
