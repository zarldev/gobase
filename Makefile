
BINARY_NAME:=${BINARY_NAME}
DOCKER_IMAGE_NAME:=${DOCKER_IMAGE_NAME}
VERSION:=${VERSION}
BUILD_TAGS:=${BUILD_TAGS}
DIST_DIR:=${DIST_DIR}
ENVIRONMENT:=${ENVIRONMENT}

GO:=go
BUILD:=$(shell git rev-parse --short HEAD)
LD_FLAGS=-ldflags "-X main.Version=$(VERSION) -X main.Build=$(BUILD)"

.PHONY: all dev build build-linux-arm64 build-linux-amd64 build-windows-amd64 build-darwin-amd64 build-darwin-arm64 build-all generate test test-coverage-html lint clean release release-linux release-windows release-darwin docker-build dev-setup help

default: help

setup:
	@echo "Setting up development environment"
	@go install github.com/cosmtrek/air@latest
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

	@npm install tailwindcss@latest postcss@latest autoprefixer@latest
	@mv .env.example .env
	@echo "Development environment setup complete"
	
run: dev

dev: binary
	@echo "Starting development environment"
	@air
	@echo "Development environment started"

binary:
	@go mod tidy
	$(GO) build $(LD_FLAGS) -tags $(ENVIRONMENT),$(BUILD_TAGS) -o $(DIST_DIR)/$(BINARY_NAME)

build: binary test lint generate
	@echo "Build complete."
	@echo
	@echo "Summary:"
	@echo 
	@echo "Binary Name: $(BINARY_NAME)"
	@echo "Environment: $(ENVIRONMENT)"
	@echo "Version: $(VERSION)"
	@echo "Build: $(BUILD)"
	@echo "Build Tags: $(BUILD_TAGS)"
	@echo "Docker Image Name: $(DOCKER_IMAGE_NAME)"
	@echo "Binary Location: $(DIST_DIR)/$(BINARY_NAME)"
	@echo "Binary Size: $(shell du -h $(DIST_DIR)/$(BINARY_NAME) | awk '{print $$1}')"

generate:
	$(GO) generate ./...
	
test: 
	$(GO) test -v ./... -coverprofile=coverage.out -covermode=atomic -coverpkg=./... 

test-coverage-html:
	$(GO) tool cover -html=coverage.out

lint:
	@echo "Running linting"
	golangci-lint run ./...

clean:
	rm -f $(DIST_DIR)/$(BINARY_NAME)
	rm -f coverage.out
	rm -f $(DIST_DIR)/$(BINARY_NAME)-$(VERSION).zip

build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GO) build $(LD_FLAGS) -tags $(ENVIRONMENT),$(BUILD_TAGS) -o $(DIST_DIR)/$(BINARY_NAME)

build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GO) build $(LD_FLAGS) -tags $(ENVIRONMENT),$(BUILD_TAGS) -o $(DIST_DIR)/$(BINARY_NAME)

build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(GO) build $(LD_FLAGS) -tags $(ENVIRONMENT),$(BUILD_TAGS) -o $(DIST_DIR)/$(BINARY_NAME).exe

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GO) build $(LD_FLAGS) -tags $(ENVIRONMENT),$(BUILD_TAGS) -o $(DIST_DIR)/$(BINARY_NAME)

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GO) build $(LD_FLAGS) -tags $(ENVIRONMENT),$(BUILD_TAGS) -o $(DIST_DIR)/$(BINARY_NAME)

build-all: build-linux-amd64 build-linux-arm64 build-windows-amd64 build-darwin-amd64 build-darwin-arm64

release: release-linux release-windows release-darwin

release-linux: build-linux-amd64 build-linux-arm64
	zip -r $(DIST_DIR)/$(BINARY_NAME)-$(VERSION)-linux-amd64.zip $(DIST_DIR)/$(BINARY_NAME)
	zip -r $(DIST_DIR)/$(BINARY_NAME)-$(VERSION)-linux-arm64.zip $(DIST_DIR)/$(BINARY_NAME)

release-windows: build-windows-amd64
	zip -r $(DIST_DIR)/$(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(DIST_DIR)/$(BINARY_NAME).exe

release-darwin: build-darwin-amd64 build-darwin-arm64
	zip -r $(DIST_DIR)/$(BINARY_NAME)-$(VERSION)-darwin-amd64.zip $(DIST_DIR)/$(BINARY_NAME)
	zip -r $(DIST_DIR)/$(BINARY_NAME)-$(VERSION)-darwin-arm64.zip $(DIST_DIR)/$(BINARY_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

help:
	@echo "Makefile commands for $(BINARY_NAME)"
	@echo
	@echo "Usage:"
	@echo "    make setup               Setup development environment"
	@echo "    make dev                 Start development environment"
	@echo "    make build               Build the binary"
	@echo "    make generate            Generate code"
	@echo "    make test                Run tests"
	@echo "    make test-coverage-html  Generate test coverage report"
	@echo "    make lint                Run linting"
	@echo "    make clean               Remove previous build"
	@echo "    make build-linux-arm64   Build for linux arm64"
	@echo "    make build-linux-amd64   Build for linux amd64"
	@echo "    make build-windows-amd64 Build for windows amd64"
	@echo "    make build-darwin-amd64  Build for darwin amd64"
	@echo "    make build-darwin-arm64  Build for darwin arm64"
	@echo "    make build-all           Build for all platforms"
	@echo "    make release             Build and package for all platforms"
	@echo "    make release-linux       Build and package for linux"
	@echo "    make release-windows     Build and package for windows"
	@echo "    make release-darwin      Build and package for darwin"
	@echo "    make docker-build        Build docker image"
	@echo "    make help                Display this help message"
	@echo
	@echo "Environment variables:"
	@echo "    <VARIABLE>:          <CURRENT_VALUE>"
	@echo "    BINARY_NAME:         $(BINARY_NAME)"
	@echo "    DOCKER_IMAGE_NAME:   $(DOCKER_IMAGE_NAME)"
	@echo "    VERSION:             $(VERSION)"
	@echo "    DIST_DIR:            $(DIST_DIR)"
	@echo "    ENVIRONMENT:         $(ENVIRONMENT)"
	@echo "    BUILD_TAGS:          $(BUILD_TAGS)"
	@echo

		