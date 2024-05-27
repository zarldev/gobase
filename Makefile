include .env

BINARY_NAME:=${BINARY_NAME}
DOCKER_IMAGE_NAME:=${DOCKER_IMAGE_NAME}
VERSION:=${VERSION}
GO:=go
BUILD:=$(shell git rev-parse --short HEAD)
DIST:=$(shell pwd)/dist

LD_FLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

.DEFAULT_GOAL := help

.PHONY: all build test test-coverage test-coverage-html lint clean release docker-build help

default: .DEFAULT_GOAL

all: help

dev:
	@echo "Starting development environment"
	@air -c .air.toml
	@echo "Development environment started"

build: test lint generate
	$(GO) build $(LD_FLAGS) -o $(DIST)/$(BINARY_NAME)
	@ls -lah  $(DIST)/$(BINARY_NAME) | awk '{print "Location:" $$9, "Size:" $$5}' | column -t
	@echo "Build complete"

build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GO) build $(LD_FLAGS) -o $(DIST)/$(BINARY_NAME)

build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GO) build $(LD_FLAGS) -o $(DIST)/$(BINARY_NAME)

build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(GO) build $(LD_FLAGS) -o $(DIST)/$(BINARY_NAME).exe

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GO) build $(LD_FLAGS) -o $(DIST)/$(BINARY_NAME)

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GO) build $(LD_FLAGS) -o $(DIST)/$(BINARY_NAME)

build-all: build-linux-amd64 build-linux-arm64 build-windows-amd64 build-darwin-amd64 build-darwin-arm64

generate:
	$(GO) generate ./...
	@templ generate
	@npx tailwindcss -i ui/static/css/styles.css -o ui/static/css/styles.css --minify
	
test: 
	$(GO) test -v ./... -coverprofile=coverage.out -covermode=atomic -coverpkg=./... 

test-coverage-html:
	$(GO) tool cover -html=coverage.out

lint:
	golangci-lint run

clean:
	rm -f $(DIST)/$(BINARY_NAME)
	rm -f coverage.out
	rm -f $(DIST)/$(BINARY_NAME)-$(VERSION).zip

release: release-linux release-windows release-darwin

release-linux: build-linux-amd64 build-linux-arm64
	zip -r $(DIST)/$(BINARY_NAME)-$(VERSION)-linux-amd64.zip $(BINARY_NAME)
	zip -r $(DIST)/$(BINARY_NAME)-$(VERSION)-linux-arm64.zip $(BINARY_NAME)

release-windows: build-windows-amd64
	zip -r $(DIST)/$(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(BINARY_NAME).exe

release-darwin: build-darwin-amd64 build-darwin-arm64
	zip -r $(DIST)/$(BINARY_NAME)-$(VERSION)-darwin-amd64.zip $(BINARY_NAME)
	zip -r $(DIST)/$(BINARY_NAME)-$(VERSION)-darwin-arm64.zip $(BINARY_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

dev-setup:
	@echo "Setting up development environment"
	@go install github.com/cosmtrek/air@latest
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install 
	@npm install tailwindcss@latest postcss@latest autoprefixer@latest
	@echo "Development environment setup complete"

help:
	@echo "Available targets:"
	@echo "  build                 Build the project"
	@echo "  test                  Run tests"
	@echo "  test-coverage         Run tests with coverage"
	@echo "  test-coverage-html    Generate HTML coverage report"
	@echo "  lint                  Run linter"
	@echo "  clean                 Clean up build artifacts"
	@echo "  release               Build and create a release zip"
	@echo "  docker-build          Build Docker image"
	@echo "  help                  Show this help message"