GOPATH ?= $(HOME)/go
APP_NAME ?= left-it-api
VERSION ?= 0.0.1

.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: watch
watch:
	$(GOPATH)/bin/air

.PHONY: build
build:
	go build -ldflags="-s -w -X 'main.Version=${VERSION}'" -o ./bin/$(APP_NAME) ./cmd
