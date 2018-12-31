.PHONY: test
test: fmtcheck ## Run all tests
	go test -i $$(go list ./... | grep -v 'vendor') $(TESTARGS) -race -timeout=30s -parallel=4

.PHONY: fmtcheck
fmtcheck: ## Run gofmt on all .go files (dry run)
	@! gofmt -d $$(find . -name '*.go' | grep -v vendor) | grep '^'

.PHONY: fmt
fmt: ## Run gofmt on all .go files
	gofmt -w $$(find . -name '*.go' | grep -v vendor)

.PHONY: vet
vet: ## Run govet on all .go files
	go tool vet -all -v $$(find . -name '*.go' | grep -v vendor)

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
