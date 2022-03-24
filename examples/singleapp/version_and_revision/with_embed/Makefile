
default: run

prepare:
	@git describe --tags --abbrev=0 > version.txt
	@git rev-list -1 HEAD > revision.txt
	@git describe --tags > build.txt

run: prepare
	@go run -race .
