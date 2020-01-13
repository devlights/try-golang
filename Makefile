GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

PRJ_NAME=try-golang
GITHUB_USER=devlights
PKG_NAME=github.com/$(GITHUB_USER)/$(PRJ_NAME)
CMD_PKG=$(PKG_NAME)/cmd/trygolang

EXAMPLE=""

ifdef ComSpec
	SEP=\\
	RM_CMD=del
	BIN_NAME=trygolang.exe
else
	SEP=/
	RM_CMD=rm -f
	BIN_NAME=trygolang
endif

.PHONY: all
all: clean build test

.PHONY: build
build:
	$(GOBUILD) -o $(BIN_NAME) -race $(CMD_PKG)

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN) $(CMD_PKG)
	$(RM_CMD) .$(SEP)$(BIN_NAME)

.PHONY: run
run: clean
	$(GORUN) $(CMD_PKG) -onetime -example ${EXAMPLE}

