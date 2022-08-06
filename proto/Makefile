IMAGE_VERSION := "latest"
GO_GRPC_PLUGIN_VERSION := "v1.2"
GO_PROTOBUF_PLUGIN_VERSION := "v1.28.1"
DART_PROTOC_PLUGIN_VERSION := "20.0.1"

.PHONY: build_go_image proto help

build_go_image: ## Build image with protoc-gen-go
	@../ci/make_scripts/build_image_protoc_gen_go.sh ${IMAGE_VERSION} ${GO_GRPC_PLUGIN_VERSION} ${GO_PROTOBUF_PLUGIN_VERSION}

build_dart_image: ## Build image with protoc-gen-dart
	@../ci/make_scripts/build_image_protoc_gen_dart.sh ${IMAGE_VERSION} ${DART_PROTOC_PLUGIN_VERSION}

proto: build_go_image build_dart_image ## Build proto files
	@../ci/make_scripts/proto.sh ${IMAGE_VERSION}

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
