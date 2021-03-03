
default: run

prepare:
	@git describe --tags --abbrev=0 > version.txt
	@git rev-list -1 HEAD >> version.txt
	@git describe --tags >> version.txt

run: prepare
	@go run -race .
