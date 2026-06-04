default: \
	dynamiclinked \
	staticlinked \
	showinfo \
	lddinfo \
	execdynamic \
	execstatic

dynamiclinked:
	go build -o app.dynamic

staticlinked:
	CGO_ENABLED=0 go build -o app.static1
	go build -tags osusergo,netgo -o app.static2

showinfo:
	@echo '[info] app.dynamic: '; file app.dynamic | tr , '\n' | grep -F 'linked'
	@echo '[info] app.static1:  '; file app.static1  | tr , '\n' | grep -F 'linked'
	@echo '[info] app.static2:  '; file app.static2  | tr , '\n' | grep -F 'linked'

lddinfo:
	@echo '[ldd ] app.dynamic: '; ldd app.dynamic || :
	@echo '[ldd ] app.static1:  '; ldd app.static1 || :
	@echo '[ldd ] app.static2:  '; ldd app.static2 || :

execdynamic:
	@echo '[exec] app.dynamic: '; sudo chroot . ./app.dynamic || :

execstatic:
	@echo '[exec] app.static1:  '; sudo chroot . ./app.static1 || :
	@echo '[exec] app.static2:  '; sudo chroot . ./app.static2 || :
