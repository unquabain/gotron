SHELL:=/bin/zsh
JSSRC=$(shell ls front-end/src/*.js front-end/src/**/*.js)
CSSSRC=$(shell ls front-end/src/*.css front-end/src/**/*.css)
GOSRC= \
	action/action.go \
	fs.go \
	gotron.go \
	state/state.go \
	state/fortune.go \
	store/store.go

bin/gotron: $(GOSRC) $(JSSRC) $(CSSSRC) assets/index.html
	go generate gotron/...
	go build -o bin/gotron 

run: bin/gotron
	bin/gotron

bin/miniserver: tools/miniserver/miniserver.go
	go build -o bin/miniserver gotron/tools/miniserver

miniserve: bin/miniserver assets/index.html
	bin/miniserver -root assets

clean:
	- rm -rf front-end/build/*
	- rm bin/*
	- rm assets/js/*.js
	- rm assets/css/*.css

.PHONY: run miniserve clean
