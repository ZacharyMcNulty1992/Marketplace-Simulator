projectPath = $(shell pwd)
export GOPATH=$(PWD)
APP_NAME = app

all: build

build: | .sys-deps .go-deps
	go build -d -ldflags "-s" -o $(projectPath)/bin/$(APP_NAME)

clean:
	rm -rf $(projectPath)/bin

.go-deps:
	go get -u github.com/g3n/engine/...
	touch .go-deps

.sys-deps:
	sudo apt-get install libopenal1 libopenal1 libgl1-mesa-dev xorg-dev
	touch .sys-deps

.PHONY clean:
	rm -rf pkg/*
	rm -rf .sys-deps
	rm -rf .go-deps
	rm -rf ./bin/app
