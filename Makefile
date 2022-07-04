GO ?= go
GOLINT ?= golint
GOSEC ?= gosec
VERSION ?=$(shell git describe --tags --always)
PACKAGES = $(shell go list -f {{.Dir}} ./... | grep -v /vendor/)
DATE = $(shell date -R)

.PHONY: help
help: ## Show this help.
	@echo "Targets:"
	@grep -E '^[a-zA-Z\/_-]*:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\t%-20s: %s\n", $$1, $$2}'

.PHONY: src-fmt
src-fmt: ## format source code.
	gofmt -s -w ${PACKAGES}

.PHONY: test
test: ## Run all the tests.
	$(GO) test -race ./... -v -cover --coverprofile=coverage.out

.PHONY: cover
cover: ## Show test coverage.
	$(GO) tool cover -html=coverage.out

.PHONY: gosec
gosec: ## Run gosec static code security checker.
	$(GOSEC) ./...

.PHONY: install-tools
install-tools: ## install dependencies if needed.
	$(GO) install github.com/securego/gosec/v2/cmd/gosec@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin latest
