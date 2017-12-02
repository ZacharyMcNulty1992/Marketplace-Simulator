projectPath = $(shell pwd)
export GOPATH=$(PWD)

all: build

build:
	go build -o $(projectPath)/bin/app $(projectPath)/src/app/app.go

clean:
	rm -rf $(projectPath)/bin
