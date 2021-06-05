.PHONY: build install

build:
	@go build ./cmd/botsend

install:
	@go install ./...