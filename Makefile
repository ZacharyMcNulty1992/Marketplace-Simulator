projectPath = $(shell pwd)
export GOPATH=$(PWD)

all: build

build: | .sys-deps .go-deps
	go build -o $(projectPath)/bin/app $(projectPath)/src/app/app.go

clean:
	rm -rf $(projectPath)/bin

.go-deps:
	go get -u -v azul3d.org/engine/gfx 
	touch .go-deps

.sys-deps:
	sudo apt-get install build-essential git mesa-common-dev libx11-dev libx11-xcb-dev libxcb-icccm4-dev libxcb-image0-dev libxcb-randr0-dev libxcb-render-util0-dev libxcb-xkb-dev libfreetype6-dev libbz2-dev libxxf86vm-dev libgl1-mesa-dev libxrandr-dev libxcursor-dev libxi-dev
	touch .sys-deps