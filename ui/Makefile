.PHONY: build_debug build_profile build_release help

build_debug: ## Build debug version
	@../ci/make_scripts/build_ui.sh "debug"

build_profile: ## Build profile version
	@../ci/make_scripts/build_ui.sh "profile"

build_release: ## Build release version
	@../ci/make_scripts/build_ui.sh "release"

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
