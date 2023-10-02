#!/usr/bin/env bash
set -euo pipefail

# Check required commands are in place
command -v docker-compose >/dev/null 2>&1 || { echo 'Please install docker-compose or use image that has it'; exit 1; }

if [[ -z "${TESTED_IMAGE:-}" ]]; then
    export TESTED_IMAGE="devs-crud:local"
fi
export ACCEPTANCE_NETWORK="devs-crud-network"

COMPOSE="docker-compose -f docker-compose.yml -f acceptance/acceptance-tester.yml"

if [ $# -gt 0 ]; then
    # use the arguments as parameters for compose if passed
    $COMPOSE "$@"
else
    # else stop and run acceptance tests
    $COMPOSE down
    $COMPOSE up --build --abort-on-container-exit --exit-code-from acceptance-tester
fi
