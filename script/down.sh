#!/bin/sh
set -euf

SCRIPT_DIR="$(cd -- "$(dirname -- "$0")" && pwd)"
PROJ_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

export COMPOSE_PROJECT_NAME="what-video"

docker-compose -f "$PROJ_DIR/docker/docker-compose.yml" $@ down
