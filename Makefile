VERSION := 0.0.1

.PHONY: build
build:
	go build -ldflags "-X main.VERSION=$(VERSION)" -o go-deconstruct .

.PHONY: test
test:
	go test ./...
