#!/usr/bin/env sh

SCRIPT_DIR=`cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd`
CI_DIR=`dirname $SCRIPT_DIR`
PROJECT_DIR=`dirname $CI_DIR`
SERVER_DIR=$PROJECT_DIR/server

cd "${SERVER_DIR}"
go build -o runify-server
