# Contributing to Safe Core SDK Golang

Thank you for your interest in contributing to the Safe Core SDK Golang! This document provides guidelines and information for contributors.

## Table of Contents

- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Submitting Changes](#submitting-changes)
- [Issue Reporting](#issue-reporting)

## Getting Started

### Prerequisites

- Go 1.21 or later
- Git
- Make (optional, for using Makefile commands)

### Development Setup

1. **Fork and clone the repository**
   ```bash
   git clone https://github.com/your-username/safe-core-sdk-golang.git
   cd safe-core-sdk-golang
   ```

2. **Install dependencies**
   ```bash
   make deps
   # or
   go mod download
   ```

3. **Run tests to verify setup**
   ```bash
   make test-unit
   ```

4. **Set up development tools**
   ```bash
   make dev-setup
   ```

## Project Structure

```
safe-core-sdk-golang/
├── types/              # Core type definitions
├── protocol/           # Safe Smart Account interaction
│   ├── managers/       # Account management (owners, modules, etc.)
│   ├── contracts/      # Contract interfaces
│   └── utils/          # Utility functions
├── api/                # Safe Transaction Service API client
├── examples/           # Example applications
│   ├── basic/          # Basic usage examples
│   └── advanced/       # Advanced usage examples
├── tests/              # Test files
│   ├── unit/           # Unit tests
│   └── integration/    # Integration tests
└── docs/               # Documentation (if any)
```

## Coding Standards

### Go Style Guide

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` for code formatting
- Use `golangci-lint` for additional linting
- Write clear, self-documenting code with appropriate comments

### Code Formatting

```bash
# Format code
make fmt

# Run linter
make lint

# Run vet
make vet
```

### Naming Conventions

- **Packages**: Short, lowercase, single words
- **Types**: PascalCase (e.g., `SafeTransaction`)
- **Functions/Methods**: PascalCase for exported, camelCase for unexported
- **Variables**: camelCase
- **Constants**: PascalCase or UPPER_CASE for package-level constants

### Documentation

- All exported functions, types, and constants must have comments
- Comments should start with the name of the item being documented
- Use complete sentences and proper grammar

Example:
```go
// SafeTransaction represents a Safe transaction with signatures.
type SafeTransaction struct {
    Data       SafeTransactionData      `json:"data"`
    Signatures map[string]SafeSignature `json:"signatures"`
}

// GetSignature returns the signature for a specific signer.
// Returns nil if no signature is found for the given signer.
func (st *SafeTransaction) GetSignature(signer string) *SafeSignature {
    // implementation
}
```

## Testing

### Unit Tests

- Write unit tests for all new functionality
- Aim for high test coverage
- Use table-driven tests where appropriate
- Mock external dependencies

```bash
# Run unit tests
make test-unit

# Run tests with coverage
make test-coverage
```

### Integration Tests

- Integration tests require real network connections
- Set environment variables for integration testing:
  ```bash
  export RUN_INTEGRATION_TESTS=true
  export RPC_URL="https://mainnet.infura.io/v3/your-project-id"
  export SAFE_ADDRESS="0x..."
  export SAFE_API_KEY="your-api-key"
  ```

```bash
# Run integration tests
make test-integration
```

### Test Structure

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name    string
        input   InputType
        want    ExpectedType
        wantErr bool
    }{
        {
            name:    "valid input",
            input:   validInput,
            want:    expectedOutput,
            wantErr: false,
        },
        // more test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := FunctionName(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("FunctionName() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("FunctionName() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Submitting Changes

### Pull Request Process

1. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make your changes**
   - Write code following the coding standards
   - Add tests for new functionality
   - Update documentation if needed

3. **Test your changes**
   ```bash
   make check  # Runs fmt, vet, lint, and unit tests
   ```

4. **Commit your changes**
   ```bash
   git add .
   git commit -m "feat: add new feature description"
   ```

5. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```

6. **Create a Pull Request**
   - Use a clear, descriptive title
   - Include a detailed description of changes
   - Reference any related issues
   - Ensure all checks pass

### Commit Message Format

Use conventional commit format:

```
type(scope): description

[optional body]

[optional footer]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:
```
feat(api): add support for Safe message signing
fix(protocol): handle edge case in transaction creation
docs: update README with new examples
test(types): add tests for SafeTransaction methods
```

## Issue Reporting

### Bug Reports

When reporting bugs, include:

1. **Description**: Clear description of the issue
2. **Steps to Reproduce**: Detailed steps to reproduce the bug
3. **Expected Behavior**: What you expected to happen
4. **Actual Behavior**: What actually happened
5. **Environment**: Go version, OS, SDK version
6. **Code Sample**: Minimal code that reproduces the issue

### Feature Requests

When requesting features, include:

1. **Use Case**: Why you need this feature
2. **Description**: Detailed description of the feature
3. **Examples**: Code examples of how it would be used
4. **Alternatives**: Alternative solutions you've considered

### Security Issues

For security-related issues, please do not open a public issue. Instead:
- Email security@safe.global
- Include detailed description and proof of concept
- Allow time for fix before public disclosure

## Development Workflow

### Local Development

1. **Start with tests**: Write tests first (TDD approach)
2. **Implement functionality**: Write the minimal code to pass tests
3. **Refactor**: Improve code quality while keeping tests green
4. **Document**: Add or update documentation

### Before Submitting

```bash
# Run all quality checks
make check

# Ensure integration tests pass (if applicable)
make test-integration

# Verify examples still work
make examples
```

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Safe Documentation](https://docs.safe.global/)
- [Ethereum Go Client](https://github.com/ethereum/go-ethereum)
- [Original TypeScript SDK](https://github.com/safe-global/safe-core-sdk)

## Questions?

If you have questions about contributing:

1. Check existing issues and discussions
2. Create a new issue with the "question" label
3. Join the Safe community channels

Thank you for contributing to Safe Core SDK Golang! 🚀