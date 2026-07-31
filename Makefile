BINARY  := scalys
PREFIX  ?= $(HOME)/.local
GOFLAGS ?=

# Tool versions used by `make tools`.
GOLANGCI_VERSION := v2.12.2
GOFUMPT_VERSION  := v0.11.0

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help
	@grep -hE '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) \
		| awk 'BEGIN{FS=":.*?## "}{printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the binary into ./scalys
	go build $(GOFLAGS) -o $(BINARY) .

.PHONY: install
install: build ## Build and install to ~/.local/bin (override with PREFIX=)
	install -Dm755 $(BINARY) $(PREFIX)/bin/$(BINARY)
	@echo "installed -> $(PREFIX)/bin/$(BINARY)"

.PHONY: run
run: ## Run the TUI without installing
	go run .

.PHONY: fmt
fmt: ## Format all Go files (gofumpt + goimports)
	golangci-lint fmt

.PHONY: fmt-check
fmt-check: ## Fail if any file needs formatting
	@golangci-lint fmt --diff

.PHONY: lint
lint: ## Run the linters
	golangci-lint run

.PHONY: lint-fix
lint-fix: ## Run the linters and apply autofixes
	golangci-lint run --fix

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: test
test: ## Run tests with the race detector
	go test -race ./...

.PHONY: cover
cover: ## Run tests and open an HTML coverage report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "wrote coverage.html"

.PHONY: tidy
tidy: ## Tidy and verify go.mod / go.sum
	go mod tidy
	go mod verify

.PHONY: check
check: fmt-check lint test ## Everything CI would run
	@echo "all checks passed"

.PHONY: clean
clean: ## Remove build and coverage artifacts
	rm -f $(BINARY) coverage.out coverage.html

.PHONY: tools
tools: ## Install the pinned lint/format tools
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_VERSION)
	go install mvdan.cc/gofumpt@$(GOFUMPT_VERSION)
