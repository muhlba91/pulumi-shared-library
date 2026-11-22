TESTPARALLELISM := 8

WORKING_DIR := $(shell pwd)

.PHONY: help check lint fix test coverage

.DEFAULT_GOAL := help

help: ## list the targets
	@grep -E '^[a-z-]+:.*## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*## "}; {printf "%-10s %s\n", $$1, $$2}'

check: lint test ## run the linters and the tests

lint: ## run golangci-lint and go vet
	golangci-lint run -c .golangci.yml
	go vet ./...

fix: ## format the code with golangci-lint
	golangci-lint fmt -c .golangci.yml

test: ## run the tests with coverage
	go test -v -tags=all -parallel ${TESTPARALLELISM} -timeout 2h -covermode atomic -coverprofile=covprofile github.com/muhlba91/pulumi-shared-library/pkg/...

coverage: ## open the HTML coverage report
	go tool cover -html=covprofile -o coverage.html
	open coverage.html
