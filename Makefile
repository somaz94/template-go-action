.PHONY: build test test-unit cover cover-html fmt vet clean help

BINARY := myaction

## Build

build: ## Build the binary
	go build -o $(BINARY) ./cmd/main.go

## Test

test: test-unit ## Run unit tests (alias)

test-unit: ## Run unit tests with coverage
	go test ./internal/... ./cmd/... -v -cover

## Coverage

cover: ## Generate coverage report
	go test ./internal/... ./cmd/... -coverprofile=coverage.out
	go tool cover -func=coverage.out

cover-html: cover ## Open coverage report in browser
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

## Quality

fmt: ## Format code
	gofmt -s -w .

vet: ## Run go vet
	go vet ./...

## Cleanup

clean: ## Remove build artifacts and coverage files
	rm -f $(BINARY) coverage.out coverage.html

## Help

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
