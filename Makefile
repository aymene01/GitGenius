BINARY_NAME=ggenius

.PHONY: build run internal-release

build:
	@echo "Building..."
	@mkdir -p bin
	@go build -o bin/$(BINARY_NAME) -v
	@echo "Build complete, binary located at bin/$(BINARY_NAME)"

run:
	@echo "Running..."
	@go run main.go

internal-release:
	@echo "Building for internal release..."
	@mkdir -p ~/bin
	@go build -o ~/bin/$(BINARY_NAME) -v
	@echo "Done! Binary for internal release is located at ~/bin/$(BINARY_NAME)"