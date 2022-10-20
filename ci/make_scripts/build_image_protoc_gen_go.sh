#!/bin/bash

SCRIPT_DIR=`cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd`
CI_DIR=`dirname $SCRIPT_DIR`
DOCKER_IMAGE_DIR=$CI_DIR/docker_images/protoc_gen_go

IMAGE_VERSION=${1}
GO_GRPC_PLUGIN_VERSION=${2}
GO_PROTOBUF_PLUGIN_VERSION=${3}
docker build -f $DOCKER_IMAGE_DIR/Dockerfile --build-arg GO_GRPC_PLUGIN_VERSION=${GO_GRPC_PLUGIN_VERSION} --build-arg GO_PROTOBUF_PLUGIN_VERSION=${GO_PROTOBUF_PLUGIN_VERSION} -t runify_protoc_gen_go:${IMAGE_VERSION} .
