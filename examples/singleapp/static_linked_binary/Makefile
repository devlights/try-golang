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
	CGO_ENABLED=0 go build -o app.static

showinfo:
	@echo '[info] app.dynamic: '; file app.dynamic | tr , '\n' | grep -F 'linked'
	@echo '[info] app.static:  '; file app.static  | tr , '\n' | grep -F 'linked'

lddinfo:
	@echo '[ldd ] app.dynamic: '; ldd app.dynamic || :
	@echo '[ldd ] app.static:  '; ldd app.static || :

execdynamic:
	@echo '[exec] app.dynamic: '; sudo chroot . ./app.dynamic || :

execstatic:
	@echo '[exec] app.static:  '; sudo chroot . ./app.static || :
