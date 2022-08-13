export GOLANGCILINT_VERSION := "v1.48.0"

.PHONY: dep test mod_tidy build run help

dep: ## Get the dependencies
	@../ci/make_scripts/install_server_deps.sh ${GOLANGCILINT_VERSION}

test: ## Run unittests
	@go test -v -count=1 `go list ./...`

mod_tidy: ## Removes unused dependencies from go.sum
	@go mod tidy -v

build: ## Build the binary file
	@../ci/make_scripts/build_server.sh

run: build ## Build and run the binary file
	@./runify-server

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
