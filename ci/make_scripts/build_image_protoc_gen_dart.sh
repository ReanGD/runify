#!/usr/bin/env sh

SCRIPT_DIR=`cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd`
CI_DIR=`dirname $SCRIPT_DIR`
DOCKER_IMAGE_DIR=$CI_DIR/docker_images/protoc_gen_dart

IMAGE_VERSION=${1}
DART_PROTOC_PLUGIN_VERSION=${2}
docker build -f $DOCKER_IMAGE_DIR/Dockerfile --build-arg DART_PROTOC_PLUGIN_VERSION=${DART_PROTOC_PLUGIN_VERSION} -t runify_protoc_gen_dart:${IMAGE_VERSION} .
