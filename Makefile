
projectPath = $(shell pwd)

all: build

build:
	go build -o $(projectPath)/bin/app $(projectPath)/app/app.go

clean:
	rm -rf $(projectPath)/bin
