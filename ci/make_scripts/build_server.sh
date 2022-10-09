#!/usr/bin/env sh

SCRIPT_DIR=`cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd`
CI_DIR=`dirname $SCRIPT_DIR`
PROJECT_DIR=`dirname $CI_DIR`
SERVER_DIR=$PROJECT_DIR/server

cd "${SERVER_DIR}"

if [ -z ${GITHUB_RUN_ID+x} ]
then
    BUILD_ID="undefined"
else
    BUILD_ID=$GITHUB_RUN_ID
fi

if [ -z ${GITHUB_SHA+x} ]
then
    BUILD_COMMIT="undefined (undefined)"
else
    BUILD_COMMIT="$GITHUB_REF ($GITHUB_SHA)"
fi

if [ -z ${GITHUB_ACTOR+x} ]
then
    BUILD_USER="undefined"
else
    BUILD_USER=$GITHUB_ACTOR
fi

BUILD_VERSION="$RUNIFY_VERSION"
BUILD_DATE_TIME=$(date "+%Y-%m-%d %H:%M:%S")

RUNIFY_BUILD_AGRS="-X 'main.version=$BUILD_VERSION' -X 'main.buildID=$BUILD_ID' -X 'main.buildCommit=$BUILD_COMMIT' -X 'main.buildUser=$BUILD_USER' -X 'main.buildDateTime=$BUILD_DATE_TIME'"

go build -ldflags "$RUNIFY_BUILD_AGRS" -o runify-server
