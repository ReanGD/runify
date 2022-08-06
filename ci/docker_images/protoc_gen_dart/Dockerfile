FROM dart:2.17.6

ARG DART_PROTOC_PLUGIN_VERSION
ENV PATH="${PATH}:/root/.pub-cache/bin"

RUN DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get install --no-install-recommends --assume-yes protobuf-compiler \
    && dart pub global activate protoc_plugin ${DART_PROTOC_PLUGIN_VERSION} \
    && chmod -R 777 /root
