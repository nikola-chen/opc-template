.PHONY: build design generate run heal clean

BINARY_NAME=bin/opc

build:
	go build -o $(BINARY_NAME) ./cmd/opc

init:
	mkdir -p runtime/logs .opc design frontend backend infra docs

design:
	go run ./cmd/opc/main.go design

generate:
	go run ./cmd/opc/main.go generate

run:
	go run ./cmd/opc/main.go run

heal:
	go run ./cmd/opc/main.go heal

clean:
	rm -rf bin/ runtime/logs/
