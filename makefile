APP_NAME = mawXscanner
BUILD_DIR = build

all: build

build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) ./...
	@echo "Build complete! Binary at $(BUILD_DIR)/$(APP_NAME)"

run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleanup complete!"

build-win:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME).exe ./...
	@echo "Windows build complete! Binary at $(BUILD_DIR)/$(APP_NAME).exe"

build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME) ./...
	@echo "Linux build complete! Binary at $(BUILD_DIR)/$(APP_NAME)"
