.PHONY: all init build
all: init clean build

init:
	go mod tidy

clean:
	rm -f build/*

build:
	go build -o build/main cmd/main/main.go

