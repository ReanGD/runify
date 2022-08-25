#!/usr/bin/env sh

GOLANGCILINT_VERSION=${1}

go mod download

GOPATH_BIN=$(go env GOPATH | cut -d ':' -f1)/bin

go install github.com/boyter/scc@latest
go install github.com/sqs/goreturns@latest
go install github.com/vektra/mockery/v2/...@v2.11.0

(cd ..; wget -O - -q  https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH_BIN} ${GOLANGCILINT_VERSION})
