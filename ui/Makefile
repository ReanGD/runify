export DIR_SCRIPTS := "../ci/make_scripts"

.PHONY: build_debug build_profile build_release release help

build_debug: ## Build debug version
	@$(DIR_SCRIPTS)/build_ui.sh "debug"

build_profile: ## Build profile version
	@$(DIR_SCRIPTS)/build_ui.sh "profile"

build_release: ## Build release version
	@$(DIR_SCRIPTS)/build_ui.sh "release"

release: build_release ## Make release for UI
	@cp -r ./build/linux/x64/release/bundle/* ../release/

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
