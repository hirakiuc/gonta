.PHONY: build run lint clean vendor_get vendor_clean vet

default: build

build: check
	go build

run: build
	go run main.go

check:
	go vet . ./internal/bot ./internal/slack ./internal/logger ./internal/plugin
	golint ./main.go ./internal/slack/*.go ./internal/logger/*.go ./internal/plugin/*.go

clean:
	go clean

vendor_get:
	glide install
