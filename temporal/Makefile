APP_NAME := watermill-servicebus
BUILD_DIR := build
MAIN_FILE := .
BINARY := $(BUILD_DIR)/$(APP_NAME)

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building the application..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BINARY) $(MAIN_FILE)

.PHONY: run
run:
	@echo "Running the application..."
	go run $(MAIN_FILE)

.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -v

.PHONY: clean
clean:
	@echo "Cleaning up build artifacts..."
	rm -rf $(BUILD_DIR)
	go clean -testcache
