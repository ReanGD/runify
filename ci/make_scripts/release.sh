#!/bin/bash

SCRIPT_DIR=`cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd`
CI_DIR=`dirname $SCRIPT_DIR`
PROJECT_DIR=`dirname $CI_DIR`
SERVER_DIR=$PROJECT_DIR/server
UI_DIR=$PROJECT_DIR/ui

cd "${PROJECT_DIR}"
echo "Prepare release dir..."
rm -rf ${PROJECT_DIR}/release
mkdir -p ${PROJECT_DIR}/release/opt/runify
cp -r ${CI_DIR}/release_files/* ${PROJECT_DIR}/release/

cd "${SERVER_DIR}"
echo "Building server..."
make -s release

cd "${UI_DIR}"
echo "Building UI..."
make -s release
