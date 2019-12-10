#!/bin/sh
set -euf

SCRIPT_DIR="$(cd -- "$(dirname -- "$0")" && pwd)"
PROJ_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
COMPOSE_PROJECT_NAME="what-video"
DOMAIN="what-video.$USER.dev.lan"
IMAGE_TAG="${IMAGE_TAG:-}"
[ -n "$IMAGE_TAG" ] || IMAGE_TAG="$(git rev-parse HEAD)"

COMPOSE_FILE="docker-compose.yml"

export IMAGE_TAG COMPOSE_FILE COMPOSE_PROJECT_NAME DOMAIN

cd "$PROJ_DIR"

docker-compose up --remove-orphans --build mysql $@
