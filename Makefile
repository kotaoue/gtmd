.PHONY: build clean install

BINARY_NAME=gtmd

build:
	go build -o $(BINARY_NAME) .

run:
	go run .
