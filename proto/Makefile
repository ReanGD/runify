IMAGE_VERSION := "latest"
GRPC_PLUGIN_VERSION := "v1.2"
PROTOBUF_PLUGIN_VERSION := "v1.28.1"

.PHONY: build_go_image proto help

build_go_image: ## Build image with protoc-gen-go
	@../ci/make_scripts/build_image_protoc_gen_go.sh ${IMAGE_VERSION} ${GRPC_PLUGIN_VERSION} ${PROTOBUF_PLUGIN_VERSION}

proto: build_go_image ## Build proto files
	@../ci/make_scripts/proto.sh ${IMAGE_VERSION}

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
