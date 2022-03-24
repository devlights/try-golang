default:

run:
	@echo '--- go build with -ldflags ---'
	cd ./with_ldflags \
		&& go build -race -ldflags \
			" \
				-X main.version=$(shell git describe --tag --abbrev=0) \
			 	-X main.revision=$(shell git rev-list -1 HEAD) \
			 	-X main.build=$(shell git describe --tags) \
			"
	@cd ./with_ldflags && ./with_ldflags
	@cd ./with_ldflags && go clean
	@echo ''

	@echo '--- go run with -ldflags ---'
	cd ./with_ldflags \
		&& go run -race -ldflags \
			" \
				-X main.version=$(shell git describe --tag --abbrev=0) \
			 	-X main.revision=$(shell git rev-list -1 HEAD) \
			 	-X main.build=$(shell git describe --tags) \
			" .
	@echo ''
