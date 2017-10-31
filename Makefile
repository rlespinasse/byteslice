.DEFAULT_GOAL := test

test: ## Start tests
	@go test -coverprofile=/tmp/byteslice-cover .

bench: ## Start benchmark
	@go test -bench=. .

help: ## Print this message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: bench help test
