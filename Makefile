# Go application name
APP_NAME=fetch-demo

# Output directories
BUILD_DIR=./build
WINDOWS_BUILD=$(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe
MAC_BUILD=$(BUILD_DIR)/$(APP_NAME)-darwin-amd64
LINUX_BUILD=$(BUILD_DIR)/$(APP_NAME)-linux-amd64

# Default target
all: build-linux build-mac build-windows

# Build for Windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(WINDOWS_BUILD)

# Build for macOS
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(MAC_BUILD)

# Build for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(LINUX_BUILD)

# Clean the build directory
clean:
	rm -rf $(BUILD_DIR)

# Remove any previously built binaries
distclean: clean
	rm -rf $(WINDOWS_BUILD) $(MAC_BUILD) $(LINUX_BUILD)

# Create the build directory
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

.PHONY: all build-windows build-mac build-linux clean distclean
