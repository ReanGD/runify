export DIR_SCRIPTS := "ci/make_scripts"

.PHONY: cloc release help

cloc: ## Count line of code
	@scc --not-match="pb" --not-match="gocc" --not-match="shortcut_ids.go" --not-match="shortcut_key_map.go" --not-match="x11_keysymdef.go" .

release: ## Make release
	@$(DIR_SCRIPTS)/release.sh

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
