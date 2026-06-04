default: run clean

build:
	@echo "[BUILD] ********************************"
	cd lib ; go build -buildmode=plugin
	go build -o app

run: build
	@echo "[RUN  ] ********************************"
	./app

clean:
	@echo "[CLEAN] ********************************"
	rm -f ./app
	find . -name "*.so" | xargs -I{} rm -f {}

file-tree:
	find . -print | sed -e "s;[^/]*/;|____;g;s;____|; |;g"