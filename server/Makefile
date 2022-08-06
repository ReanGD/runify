export GOLANGCILINT_VERSION := "v1.48.0"

.PHONY: dep build help

dep: ## Get the dependencies
	@../ci/make_scripts/install_server_deps.sh ${GOLANGCILINT_VERSION}

build: ## Build the binary file
	@../ci/make_scripts/build_server.sh

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
