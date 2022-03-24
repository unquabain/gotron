SHELL:=/bin/zsh
JSSRC=$(shell ls ui/src/*.js ui/src/**/*.js)
CSSSRC=$(shell ls ui/src/*.css ui/src/**/*.css)
GOSRC= \
	action/action.go \
	index.go \
	app.go \
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
	- rm -rf ui/build/*
	- rm bin/*
	- rm assets/js/*.js
	- rm assets/css/*.css

.PHONY: run miniserve clean
