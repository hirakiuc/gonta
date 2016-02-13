.PHONY: build run lint clean vendor_get vendor_clean vet

GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

default: build

build: vet
	go build -v -o ./gonta

run: build
	./gonta

lint:
	golint ./main.go ./slack ./logger ./plugin

clean:
	rm -f ./gonta

vendor_get: vendor_clean
	gom install

vendor_clean:
	rm -rf ./_vendor
	mkdir _vendor

vet:
	go vet ./
