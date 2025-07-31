# Makefile for CleanURL

# Binary name
BINARY_NAME=cleanurl

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-X main.Version=$(shell git describe --tags --always --dirty)"

# Default target
all: test build

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) $(LDFLAGS)

# Build for multiple platforms
build-all: build-linux build-darwin build-windows

build-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-linux $(LDFLAGS)

build-darwin:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-darwin $(LDFLAGS)

build-windows:
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-windows.exe $(LDFLAGS)

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-linux
	rm -f $(BINARY_NAME)-darwin
	rm -f $(BINARY_NAME)-windows.exe

# Run tests
test:
	$(GOTEST) -v

# Run tests with coverage
test-coverage:
	$(GOTEST) -v -cover

# Run tests with coverage report
test-coverage-html:
	$(GOTEST) -v -coverprofile=coverage.out
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Install the binary globally (requires sudo on Unix systems)
install:
	cp $(BINARY_NAME) /usr/local/bin/

# Uninstall the binary
uninstall:
	rm -f /usr/local/bin/$(BINARY_NAME)

# Run the application (for testing)
run:
	./$(BINARY_NAME)

# Show help
help:
	@echo "Available targets:"
	@echo "  build              - Build the application"
	@echo "  build-all          - Build for Linux, macOS, and Windows"
	@echo "  build-linux        - Build for Linux"
	@echo "  build-darwin       - Build for macOS"
	@echo "  build-windows      - Build for Windows"
	@echo "  clean              - Clean build artifacts"
	@echo "  test               - Run tests"
	@echo "  test-coverage      - Run tests with coverage"
	@echo "  test-coverage-html - Run tests with HTML coverage report"
	@echo "  deps               - Install dependencies"
	@echo "  install            - Install binary globally"
	@echo "  uninstall          - Uninstall binary"
	@echo "  run                - Run the application"
	@echo "  help               - Show this help"

# Phony targets
.PHONY: all build build-all build-linux build-darwin build-windows clean test test-coverage test-coverage-html deps install uninstall run help 