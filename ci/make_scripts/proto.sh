#!/usr/bin/env sh

SCRIPT_DIR=`cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd`
CI_DIR=`dirname $SCRIPT_DIR`
PROJECT_DIR=`dirname $CI_DIR`
PROTO_DIR=$PROJECT_DIR/proto

IMAGE_VERSION=${1}

cd "${PROJECT_DIR}"
docker run --rm -it -v `pwd`:/project:rw -w /project -u $(id -u ${USER}):$(id -g ${USER}) --network none runify_protoc_gen_go:${IMAGE_VERSION} protoc --go_out=./server --go-grpc_out=./server ./proto/runify.proto
mkdir ./ui/lib/pb
docker run --rm -it -v `pwd`:/project:rw -w /project -u $(id -u ${USER}):$(id -g ${USER}) --network none runify_protoc_gen_dart:${IMAGE_VERSION} protoc --dart_out=grpc:./ui/lib/pb --proto_path=proto ./proto/runify.proto
