# Contributing to MakLogger

First off, thank you for considering contributing to MakLogger! 🎉

The following is a set of guidelines for contributing to MakLogger. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

## Code of Conduct

This project and everyone participating in it is governed by our commitment to creating a welcoming and inclusive environment. By participating, you are expected to uphold this standard.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible:

- **Use a clear and descriptive title**
- **Describe the exact steps which reproduce the problem**
- **Provide specific examples to demonstrate the steps**
- **Describe the behavior you observed after following the steps**
- **Explain which behavior you expected to see instead and why**
- **Include screenshots and animated GIFs if helpful**

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, please include:

- **Use a clear and descriptive title**
- **Provide a step-by-step description of the suggested enhancement**
- **Provide specific examples to demonstrate the steps**
- **Describe the current behavior and explain which behavior you expected to see instead**
- **Explain why this enhancement would be useful**

### Pull Requests

1. Fork the repo and create your branch from `main`
2. If you've added code that should be tested, add tests
3. If you've changed APIs, update the documentation
4. Ensure the test suite passes
5. Make sure your code lints
6. Issue that pull request!

## Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/makhkets/maklogger.git
   cd maklogger
   ```

2. **Install Go 1.21 or higher**
   - Download from [golang.org](https://golang.org/dl/)

3. **Run tests**
   ```bash
   go test -v ./...
   ```

4. **Run examples**
   ```bash
   cd examples/basic && go run main.go
   cd ../advanced && go run main.go
   cd ../no-colors && go run main.go
   ```

## Style Guide

### Go Code Style

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code
- Use `golint` and `go vet` to catch common mistakes
- Write meaningful commit messages

### Git Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line

### Documentation

- Keep README.md up to date
- Comment exported functions and types
- Provide examples for new features
- Update CHANGELOG.md for significant changes

## Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting PR
- Include both positive and negative test cases
- Test cross-platform compatibility when relevant

## Performance

- Profile performance-critical code changes
- Include benchmarks for performance improvements
- Consider memory allocation patterns

## Questions?

Feel free to open an issue with your question or reach out to the maintainers.

Thank you for contributing! 🚀