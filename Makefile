# Makefile
export GO111MODULE=on

GO_CMD=go
GO_GET=$(GO_CMD) get -v
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_TEST_COVER=$(GO_CMD) test -cover
GO_INSTALL=$(GO_CMD) install -v

BIN=glox
DIR=.
MAIN=glox.go

SOURCE_PKG_DIR= .
SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

all: dependencies tests build-server

dependencies:
	@echo "==> Installing dependencies ...";
	@$(GO_GET) ./...

tests:
	@echo "==> Running tests ...";
	@$(GO_TEST_COVER) $(SOURCE_PKG_DIR)
	# @$(GO_TEST_COVER) $(SOURCE_PKG_DIR)/ishell

build:
	@echo "==> Building glox ...";
	@$(GO_BUILD) -o $(BIN) -ldflags "-w -s" $(DIR)/$(MAIN) || exit 1;
	@chmod 755 $(BIN)

run:
	./$(BIN)

format:
	for f in `find . -name '*.go'`; do go fmt $f; done
