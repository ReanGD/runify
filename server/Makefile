export RUNIFY_VERSION := "v0.1.0"
export GOLANGCILINT_VERSION := "v1.48.0"
export DIR_SCRIPTS := "../ci/make_scripts"

.PHONY: dep gocc test mod_tidy build release run help

dep: ## Get the dependencies
	@$(DIR_SCRIPTS)/install_server_deps.sh ${GOLANGCILINT_VERSION}

gocc: ## Generate the gocc parser
	@$(DIR_SCRIPTS)/gocc.sh

test: ## Run unittests
	@go test -v -count=1 `go list ./...`

mod_tidy: ## Removes unused dependencies from go.sum
	@go mod tidy -v

build: ## Build the binary file
	@$(DIR_SCRIPTS)/build_server.sh

release: build ## Make release for server
	@cp ./runify-server ../release/opt/runify/

run: build ## Build and run the binary file
	@./runify-server -c ../ci/cfg/dev.json

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
