export RUNIFY_VERSION := "v0.3.0"
export GOLANGCILINT_VERSION := "v1.48.0"
export DIR_SCRIPTS := "../ci/make_scripts"

.PHONY: dep_fmt dep gocc test mod_tidy fmt_check fmt build release clear_dev_log run help

dep_fmt: ## Get the dependencies for the fmt target
	@go install mvdan.cc/gofumpt@latest

dep: dep_fmt ## Get the dependencies
	@$(DIR_SCRIPTS)/install_server_deps.sh ${GOLANGCILINT_VERSION}

gocc: ## Generate the gocc parser
	@$(DIR_SCRIPTS)/gocc.sh

test: ## Run unittests
	@go test -v -count=1 `go list ./...`

mod_tidy: ## Removes unused dependencies from go.sum
	@go mod tidy -v

fmt_check: ## Check if the code is formatted
	@$(DIR_SCRIPTS)/fmt_check_server.sh

fmt: ## Format source code
	@gofumpt -l -w .

build: ## Build the binary file
	@$(DIR_SCRIPTS)/build_server.sh

release: build ## Make release for server
	@cp ./runify-server ../release/opt/runify/

clear_dev_log: ## Clear the dev log file
	@echo > ./runify.log

run: build clear_dev_log ## Build and run the binary file
	@./runify-server -c ../ci/cfg/dev.json

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
