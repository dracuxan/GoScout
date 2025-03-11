APP_NAME := goscout
BUILD_DIR := bin
MAIN_FILE := main.go

all: deps build  

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME)

run: build
	./$(BUILD_DIR)/$(APP_NAME) $(ARGS)

test:
	@go test -v ./...

cover:
	@go test -cover ./...

clean:
	@rm -rf $(BUILD_DIR)

format:
	@go fmt ./...

deps:
	@go mod tidy

help: 
	@echo "Usage: make [target] [ARGS=\"flags\"]"
	@echo "Available targets:"
	@echo "  build      Build the application"
	@echo "  run        Run the application (supports ARGS=\"--flag value\")"
	@echo "  test       Run all tests"
	@echo "  cover      Run coverage tests"
	@echo "  clean      Clean up build artifacts"
	@echo "  format     Format the code"
	@echo "  deps       Install dependencies"
	@echo "Examples:"
	@echo "  make run ARGS=\"-p <directory_path> -q <search_query>\"   # Run with flags"
	@echo "  make build																						# Build only"

