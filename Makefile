.PHONY: all build deps generate help test version

CHECK_FILES ?= $$(go list ./... | grep -v /vendor/)

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: deps generate build test ## Run all steps

build: ## Build all
	go build ./...

deps: ## Download dependencies.
	go mod tidy

generate: ## Generate theme files
	cd tools && go run generate-theme.go

test: ## Run tests
	go test -v $(CHECK_FILES)

version: ## Run the goversion tool. Usage: make version bump="patch"
	@# Check that the 'bump' variable is provided.
	@if [ -z "$(bump)" ]; then \
		echo "Error: You must provide a bump argument, e.g. bump=\"patch\"" >&2; \
		exit 1; \
	fi
	@echo "Running goversion with bump: $(bump)"
	go tool github.com/bcomnes/goversion/v2 -bump-file  extension.toml $(bump)
