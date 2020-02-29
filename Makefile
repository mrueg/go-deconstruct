VERSION := $(shell cat VERSION)

.PHONY: build
build:
	go build -ldflags "-X main.VERSION=$(VERSION)" -o go-deconstruct .

.PHONY: test
test:
	go test ./...
