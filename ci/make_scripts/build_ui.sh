#!/usr/bin/env sh
MODE=--${1}

SCRIPT_DIR=`cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd`
CI_DIR=`dirname $SCRIPT_DIR`
PROJECT_DIR=`dirname $CI_DIR`
UI_DIR=$PROJECT_DIR/server

cd "${SERVER_DIR}"
flutter build linux ${MODE}
