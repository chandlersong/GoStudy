#!/bin/bash

set -o errexit
BASEDIR="$1"
echo $BASEDIR

sam local start-api \
  --host 0.0.0.0 \
  --docker-volume-basedir "${BASEDIR}"