BINARY_NAME = godemo
EXCUTE_PATH = $(shell pwd)
VERSION = $(shell git describe --tags)
BUILD_TIME = $(shell date -u "+%Y-%m-%d %H:%M:%S UTC")

define build
	@echo "Building $(BINARY_NAME)..."
	go build -ldflags="-X 'cmd.version=$(VERSION)' -X 'cmd.buildTime=$(BUILD_TIME)'" -o godemo main.go
	@echo "Build completed! Excutable path is $(EXCUTE_PATH)/$(BINARY_NAME)"
endef

wsl:
	go env -w CGO_ENABLED=0
	go env -w GOOS=linux
	go env -w GOARCH=amd64
	$(call build)

mac:
	go env -w CGO_ENABLED=0
	go env -w GOOS=darwin
	go env -w GOARCH=arm64
	$(call build)
