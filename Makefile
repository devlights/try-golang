GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOGENERATE=$(GOCMD) generate

PRJ_NAME=try-golang
GITHUB_USER=devlights
PKG_NAME=github.com/$(GITHUB_USER)/$(PRJ_NAME)
CMD_PKG=$(PKG_NAME)/cmd/trygolang

EXAMPLE=""

ifdef ComSpec
	SEP=\
	RM_CMD=del
	BIN_NAME=trygolang.exe
else
	SEP=/
	RM_CMD=rm -f
	BIN_NAME=trygolang
endif

all: clean build test

build:
	$(GOBUILD) -o $(BIN_NAME) -race $(CMD_PKG)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN) $(CMD_PKG)
	$(RM_CMD) .$(SEP)$(BIN_NAME)

run: clean generate
	$(GORUN) $(CMD_PKG) -onetime -example ${EXAMPLE}

generate:
	$(GOGENERATE) -x ./...

.PHONY: all build test clean run generate

