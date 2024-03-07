#!/usr/bin/env bash

set -e

if [[ "$OUT" == "" ]]; then
  OUT="$PWD/result"
fi

case "$1" in
  test)
    cd src
    $GO test
    ;;
  docker)
    if [[ "$TAG" == "" ]]; then
      echo "Specify TAG env variable."
      exit 1
    fi
    # set image name to 673156464838.dkr.ecr.us-west-2.amazonaws.com/uptime-service-backend if IMAGE_NAME is not set
    IMAGE_NAME=${IMAGE_NAME:-673156464838.dkr.ecr.us-west-2.amazonaws.com/cassandra-updater}
    docker build -t "$IMAGE_NAME:$TAG" .
    ;;
  "")
    cd src
    $GO build -o "$OUT/bin/cassandra-updater"
    ;;
  *)
    echo "unknown command $1"
    exit 2
    ;;
esac