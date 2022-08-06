FROM golang:1.19.0-alpine3.16

ARG GRPC_PLUGIN_VERSION
ARG PROTOBUF_PLUGIN_VERSION
ENV GOPATH="/go"

RUN apk update \
    && apk add --no-cache protobuf make git \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@${GRPC_PLUGIN_VERSION} \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@${PROTOBUF_PLUGIN_VERSION}
