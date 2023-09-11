.PHONY: all build clean test run

# Project variables
BINARY_NAME ?= rose-go-driver
IMAGE_NAME ?= quay.io/rose/rose-go-driver
MAIN_FILE ?= cmd/main.go
PORT ?= 8081

all: build

build:
	@echo "Building..."
	go build -o $(BINARY_NAME) $(MAIN_FILE)

clean:
	@echo "Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)

test:
	@echo "Running tests..."
	go test -v ./...

run: build
	@echo "Running..."
	./$(BINARY_NAME)

build-image:
	@echo "Building Docker image..."
	podman build -t $(IMAGE_NAME) .

run-image:
	@echo "Running container image ..."
	podman run --rm \
		--network host \
		-it $(IMAGE_NAME) \
		-port $(PORT)

deps:
	@echo "Fetching dependencies..."
	go mod tidy
	go mod download
