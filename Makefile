.PHONY: lint mod-update install test

tool:
	go get -u golang.org/x/lint/golint

lint:
	go vet ./...
	golint ./...

mod-update:
	GO111MODULE=on go get -u -m
	GO111MODULE=on go mod tidy

install:
	go install

test:
	go test ./...
