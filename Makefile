GOCMD=go
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

RM_CMD=rm -f
BIN_NAME=trygolang
BIN_DIR=bin

.PHONY: all
all: clean build test

.PHONY: prepare
prepare: \
	_go_get \

_go_get:
	$(GOCMD) mod download
	go install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: build
build: 
	$(GOBUILD) -o bin/$(BIN_NAME) $(CMD_PKG)

.PHONY: buildstatic
buildstatic: 
	CGO_ENABLED=0 $(GOBUILD) -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' -o bin/$(BIN_NAME)_staticlink $(CMD_PKG)

.PHONY: test
test: 
	$(GOTEST) -coverprofile /tmp/try-golang-cover $(shell go list ./... | grep -v /examples/ | grep -v /cmd/)

.PHONY: clean
clean: 
	$(GOCLEAN) -i $(CMD_PKG)
	$(RM_CMD) $(BIN_DIR)/$(BIN_NAME)

.PHONY: install
install: 
	$(GOINSTALL) $(BIN_DIR)

.PHONY: vet
vet:
	go vet ./...
	staticcheck ./...

.PHONY: run
run: vet
	$(GORUN) $(CMD_PKG) -onetime -example ${EXAMPLE}

.PHONY: generate
generate: 
	$(GOGENERATE) -x ./...

.PHONY: docker-build
docker-build:
	sudo $(DOCKER_BUILD) -t try-golang -f Dockerfile ${PWD}

.PHONY: docker-run
docker-run: docker-build
	sudo $(DOCKER_RUN) -it --rm --name try-golang try-golang

.PHONY: docker-sh
docker-sh: docker-build
	sudo $(DOCKER_RUN) -it --rm --name try-golang try-golang bash

.PHONY: docker
docker: docker-run

ldflags_example:
	@echo '--- go build with -ldflags ---'
	cd ./cmd/version_and_revision/with_ldflags \
		&& $(GOBUILD) -race -ldflags \
			" \
				-X main.version=$(shell git describe --tag --abbrev=0) \
			 	-X main.revision=$(shell git rev-list -1 HEAD) \
			 	-X main.build=$(shell git describe --tags) \
			"
	@cd ./cmd/version_and_revision/with_ldflags && ./with_ldflags
	@cd ./cmd/version_and_revision/with_ldflags && go clean
	@echo ''

	@echo '--- go run with -ldflags ---'
	cd ./cmd/version_and_revision/with_ldflags \
		&& $(GORUN) -race -ldflags \
			" \
				-X main.version=$(shell git describe --tag --abbrev=0) \
			 	-X main.revision=$(shell git rev-list -1 HEAD) \
			 	-X main.build=$(shell git describe --tags) \
			" .
	@echo ''
