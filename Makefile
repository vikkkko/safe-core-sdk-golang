# Safe Core SDK Golang Makefile

.PHONY: help test test-unit test-integration build clean examples lint fmt vet deps

# Default target
help:
	@echo "Safe Core SDK Golang"
	@echo "Available targets:"
	@echo "  help              Show this help message"
	@echo "  deps              Download dependencies"
	@echo "  build             Build the project"
	@echo "  test              Run all tests"
	@echo "  test-unit         Run unit tests only"
	@echo "  test-integration  Run integration tests (requires env vars)"
	@echo "  lint              Run linter"
	@echo "  fmt               Format code"
	@echo "  vet               Run go vet"
	@echo "  examples          Run example applications"
	@echo "  clean             Clean build artifacts"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Build the project
build: deps
	@echo "Building project..."
	go build ./...

# Run all tests
test: test-unit test-integration

# Run unit tests
test-unit:
	@echo "Running unit tests..."
	go test -v ./tests/unit/...

# Run integration tests (requires environment variables)
test-integration:
	@echo "Running integration tests..."
	@echo "Note: Set RUN_INTEGRATION_TESTS=true and required env vars to run these tests"
	go test -v ./tests/integration/...

# Run unit tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./tests/unit/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Run golangci-lint (install first with: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2)
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install it with:"; \
		echo "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b \$$(go env GOPATH)/bin v1.54.2"; \
	fi

# Run examples
examples: build
	@echo "Running basic example..."
	@echo "Note: Update RPC_URL and other config in examples before running"
	@echo "Uncomment the following line to run:"
	# go run ./examples/basic/
	@echo ""
	@echo "Running advanced example..."
	@echo "Uncomment the following line to run:"
	# go run ./examples/advanced/

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	go clean ./...
	rm -f coverage.out coverage.html

# Run all quality checks
check: fmt vet lint test-unit
	@echo "All quality checks passed!"

# Development setup
dev-setup: deps
	@echo "Setting up development environment..."
	@echo "Installing development tools..."
	go install golang.org/x/tools/cmd/goimports@latest
	@echo "Development setup complete!"

# Create a new release (example)
release:
	@echo "Creating release..."
	@echo "Make sure to:"
	@echo "1. Update version in go.mod"
	@echo "2. Update CHANGELOG.md"
	@echo "3. Create and push git tag"
	@echo "4. Create GitHub release"

# Docker build (if needed)
docker-build:
	@echo "Building Docker image..."
	docker build -t safe-core-sdk-golang .

# Show project status
status:
	@echo "Project Status:"
	@echo "==============="
	go version
	@echo ""
	@echo "Dependencies:"
	go list -m all
	@echo ""
	@echo "Test coverage:"
	@if [ -f coverage.out ]; then \
		go tool cover -func=coverage.out | tail -1; \
	else \
		echo "No coverage data. Run 'make test-coverage' first."; \
	fi