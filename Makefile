projectPath = $(shell pwd)
export GOPATH=$(PWD)

all: build

build: get-sys-deps get-deps
	go build -o $(projectPath)/bin/app $(projectPath)/src/app/app.go

clean:
	rm -rf $(projectPath)/bin

get-deps:
	go get -u -v azul3d.org/engine/gfx

get-sys-deps:
	sudo apt-get install build-essential git mesa-common-dev libx11-dev libx11-xcb-dev libxcb-icccm4-dev libxcb-image0-dev libxcb-randr0-dev libxcb-render-util0-dev libxcb-xkb-dev libfreetype6-dev libbz2-dev libxxf86vm-dev libgl1-mesa-dev libxrandr-dev libxcursor-dev libxi-dev