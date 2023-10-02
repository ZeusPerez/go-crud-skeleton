.PHONY: help check test fmt vet lint shellcheck fix-fmt compile tidy build acceptance acceptance-ruby

help: ## Show this help
	@echo "Help"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

##
### Code validation
check: ## Run all checks: test, fmt, vet, lint, tidy, shellcheck
	@bash scripts/check.sh

test: ## Run tests for all go packages
	@bash scripts/checks/test.sh

coverage: test ## Run tests and open coverage report
	@go tool cover -html=.test_coverage.out

fmt: ## Run goimports on all packages, printing files that don't match code-format if any
	@bash scripts/checks/fmt.sh

vet: ## Run vet on all packages (more info running `go doc cmd/vet`)
	@bash scripts/checks/vet.sh

tidy: ## Prefetch deps to ensure required versions are downloaded
	@bash scripts/checks/tidy.sh

lint: ## Run lint on the codebase, printing any style errors
	@bash scripts/checks/lint.sh

fix-fmt: ## Run goimports on all packages, fix files that don't match code-style
	@bash scripts/local/fix-fmt.sh

fix-tidy: ## Fix go.mod inconsistency
	@bash scripts/local/fix-tidy.sh

compile: ## Compile the binary
	@bash scripts/compile.sh

build: ## Build the image
	@bash scripts/build.sh

acceptance: ## Run acceptance tests for the built image
	@bash scripts/acceptance.sh

install-tools: ## Install external tools
	@bash scripts/install-tools.sh

install-deps: ## Prefetch deps to ensure required versions are downloaded
	@go mod tidy
	@go mod verify
	@go mod vendor

