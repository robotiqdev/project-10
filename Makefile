.PHONY: all test cover bench build lint

all: test cover build

test:
	go test ./...

cover:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out -o coverage.html

bench:
	go test -bench=. -benchmem ./...

build:
	go build -o calc .

lint:
	go vet ./...
