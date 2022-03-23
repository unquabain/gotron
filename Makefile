SHELL:=/bin/zsh
JSSRC=$(shell ls front-end/src/*.js front-end/src/**/*.js)
CSSSRC=$(shell ls front-end/src/*.css front-end/src/**/*.css)
GOSRC= \
	back-end/action/action.go \
	back-end/fs.go \
	back-end/gotron.go \
	back-end/state/state.go \
	back-end/state/fortune.go \
	back-end/store/store.go

bin/gotron: $(GOSRC) $(JSSRC) $(CSSSRC) back-end/assets/index.html
	go generate gotron/...
	go build -o bin/gotron gotron/back-end

run: bin/gotron
	bin/gotron

bin/miniserver: tools/miniserver/miniserver.go
	go build -o bin/miniserver gotron/tools/miniserver

miniserve: bin/miniserver back-end/assets/index.html
	bin/miniserver -root back-end/assets

clean:
	- rm -rf front-end/build/*
	- rm bin/*
	- rm back-end/assets/js/*.js
	- rm back-end/assets/css/*.css

.PHONY: run miniserve clean
