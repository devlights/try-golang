EXAMPLE=""

all: clean build test

prepare: prepare-release
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/go-task/task/v3/cmd/task@latest

prepare-release:
	go mod download

build: 
	go build -race

build-static: 
	CGO_ENABLED=0 go build -a -tags osusergo,netgo -ldflags '-extldflags "-static"'

build-release:
	CGO_ENABLED=0 go build -a -tags osusergo,netgo -ldflags '-s -w -extldflags "-static"' -trimpath

test: 
	go test -race -coverprofile /tmp/try-golang-cover $(shell go list ./... | grep -v /examples/ | grep -v /cmd)

clean: 
	go clean

install: 
	go install

vet:
	go vet ./...
	staticcheck ./...

run:
	go run -race main.go -onetime -example ${EXAMPLE}

generate: 
	go generate -x ./...

docker-build:
	DOCKER_BUILDKIT=1 docker image build -t try-golang -f Dockerfile ${PWD}

docker-run: docker-build
	docker container run -it --rm --name try-golang try-golang

docker-sh: docker-build
	docker container run -it --rm --name try-golang try-golang bash

docker: docker-run
