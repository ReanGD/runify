FROM golang:1.19.0-alpine3.16

ARG GO_GRPC_PLUGIN_VERSION
ARG GO_PROTOBUF_PLUGIN_VERSION
ENV GOPATH="/go"

RUN apk update \
    && apk add --no-cache protobuf \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${GO_GRPC_PLUGIN_VERSION} \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@${GO_PROTOBUF_PLUGIN_VERSION}
