GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

GITHUB_USER=devlights
PKG_NAME=try-golang
BIN_NAME=trygolang
CMD_PKG=github.com/$(GITHUB_USER)/$(PKG_NAME)/cmd/$(BIN_NAME)

.PHONY: all
all: clean build test

.PHONY: build
build:
	$(GOBUILD) -o $(BIN_NAME) -v $(CMD_PKG)

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(BIN_NAME)

.PHONY: run
run: clean build
	./$(BIN_NAME) -onetime

