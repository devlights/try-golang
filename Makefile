GOCMD=go1.16
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
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
	RM_CMD=del
	BIN_NAME=.\trygolang.exe
	BIN_DIR=.\cmd\trygolang
else
	RM_CMD=rm -f
	BIN_NAME=./trygolang
	BIN_DIR=./cmd/trygolang
endif

.PHONY: all
all: clean build test

.PHONY: build
build:
	$(GOBUILD) -race -o $(BIN_NAME) $(CMD_PKG)

.PHONY: test
test:
	$(GOTEST) -race -v ./...

.PHONY: clean
clean:
	$(GOCLEAN) -i $(CMD_PKG)
	$(RM_CMD) $(BIN_NAME)

.PHONY: install
install:
	$(GOINSTALL) $(BIN_DIR)

.PHONY: run
run:
	$(GORUN) -race $(CMD_PKG) -onetime -example ${EXAMPLE}

.PHONY: generate
generate:
	$(GOGENERATE) -x ./...

.PHONY: docker-build
docker-build:
	sudo $(DOCKER_BUILD) -t try-golang .

.PHONY: docker-run
docker-run: docker-build
	sudo $(DOCKER_RUN) -it --rm --name try-golang try-golang

.PHONY: docker-sh
docker-sh: docker-build
	sudo $(DOCKER_RUN) -it --rm --name try-golang try-golang bash

.PHONY: docker
docker: docker-run

installgo116:
	GO111MODULE=off go get golang.org/dl/go1.16
	go1.16 download
