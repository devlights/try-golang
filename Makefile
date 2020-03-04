GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOGENERATE=$(GOCMD) generate
DOCKER=docker
DOCKER_BUILD=$(DOCKER) build
DOCKER_RUN=$(DOCKER) run

PRJ_NAME=try-golang
GITHUB_USER=devlights
PKG_NAME=github.com/$(GITHUB_USER)/$(PRJ_NAME)
CMD_PKG=$(PKG_NAME)/cmd/trygolang

EXAMPLE=""

ifdef ComSpec
	SEP="\\"
	RM_CMD=del
	BIN_NAME=trygolang.exe
else
	SEP=/
	RM_CMD=rm -f
	BIN_NAME=trygolang
endif

all: clean build test

build:
	$(GOBUILD) -o $(BIN_NAME) $(CMD_PKG)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN) $(CMD_PKG)
	$(RM_CMD) .$(SEP)$(BIN_NAME)

run: clean
	$(GORUN) $(CMD_PKG) -onetime -example ${EXAMPLE}

generate:
	$(GOGENERATE) -x ./...

docker-build:
	$(DOCKER_BUILD) -t try-golang .

docker-run: docker-build
	$(DOCKER_RUN) -it --rm --name try-golang try-golang

docker-sh: docker-build
	$(DOCKER_RUN) -it --rm --name try-golang try-golang bash

docker: docker-run

.PHONY: all build test clean run generate docker-build docker-run docker-sh docker

