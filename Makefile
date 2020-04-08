.PHONY: build build-windows build-linux run test coverage fmt doc

ifndef VERBOSE
MAKEFLAGS+=--no-print-directory
endif

ifeq ($(UNAME),Darwin)
ECHO=echo
else
ECHO=echo -e
endif

# Package
PACKAGE_NAME=motronic_checksum
BUILD=$(shell git rev-list --count HEAD)
LDFLAGS=-ldflags '-w -s -v'

SRCS=./cmd/motronic/*.go

default: build

build:
	-@$(ECHO) "\n\033[0;35m%%% Building tools\033[0m"
	-@$(ECHO) "Building..."
	CGO_ENABLED=0 go build $(LDFLAGS) -v -o ./dist/$(PACKAGE_NAME)_macos $(SRCS)
	-@$(ECHO) "\n\033[1;32mDone!\033[0m\n"

build-windows:
	-@$(ECHO) "\n\033[0;35m%%% Building tools for Windows\033[0m"
	-@$(ECHO) "Building.."
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -v -o ./dist/$(PACKAGE_NAME)_win64.exe $(SRCS)
	-@$(ECHO) "\n\033[1;32mDone!\033[0m\n"

build-linux:
	-@$(ECHO) "\n\033[0;35m%%% Building tools for Linux\033[0m"
	-@$(ECHO) "Building.."
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build $(LDFLAGS) -v -o ./dist/$(PACKAGE_NAME)_linux64 $(SRCS)
	-@$(ECHO) "\n\033[1;32mDone!\033[0m\n"

run:
	go run ./cmd/motronic/main.go

test:
	-@$(ECHO) "\n\033[0;35m%%% Running tests\033[0m"
	go test -v ./...

coverage:
	-@$(ECHO) "\n\033[0;35m%%% Running test coverage\033[0m"
	go test -cover ./...

doc:
  godoc -http=:6060 -index

fmt:
	go fmt ./...