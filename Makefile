projectPath = $(shell pwd)
export GOBIN = $(projectPath)/bin/
APP_NAME = game
export GOPATH=$(PWD)

all: build

run: build
	$(projectPath)/bin/$(APP_NAME)

build: | .sys-deps .go-deps
	go build -o $(projectPath)/bin/$(APP_NAME) $(projectPath)/src/$(APP_NAME)/app/$(APP_NAME).go

.go-deps:
	go get -u -v azul3d.org/engine/gfx 
	#go get -u github.com/g3n/engine/...
	#$(projectPath)/bin/glide get azul3d.org/engine/gfx 
	touch .go-deps

.sys-deps:
	sudo apt-get install build-essential git mesa-common-dev libx11-dev libx11-xcb-dev libxcb-icccm4-dev libxcb-image0-dev libxcb-randr0-dev libxcb-render-util0-dev libxcb-xkb-dev libfreetype6-dev libbz2-dev libxxf86vm-dev libgl1-mesa-dev libxrandr-dev libxcursor-dev libxi-dev
	#sudo apt-get install libopenal1 libopenal1 libgl1-mesa-dev xorg-dev
	touch .sys-deps

.init-glide:
	$(projectPath)/bin/glide init

.get-glide:
	curl https://glide.sh/get | sh


.glide: | .get-glide .init-glide 
	touch .glide

clean:
	rm -rf pkg/*
	rm -rf src/github.com
	rm -rf src/azul3d.org
	rm -rf .sys-deps
	rm -rf .go-deps
	rm -rf ./bin/app

clean-bin:
	rm /bin/app