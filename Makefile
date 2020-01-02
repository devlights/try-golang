GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

PRJ_NAME=try-golang
GITHUB_USER=devlights
PKG_NAME=github.com/$(GITHUB_USER)/$(PRJ_NAME)
BIN_NAME=trygolang
CMD_PKG=$(PKG_NAME)/cmd/$(BIN_NAME)

.PHONY: all
all: clean build test

.PHONY: build
build:
	$(GOBUILD) -o $(BIN_NAME) $(CMD_PKG)

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f ./$(BIN_NAME)

.PHONY: run
run: clean build
	./$(BIN_NAME) -onetime
	rm -f ./$(BIN_NAME)

