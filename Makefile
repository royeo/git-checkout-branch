.PHONY: install mod tool fmt lint test build bin

install:
	go install

mod:
	GO111MODULE=on go get -d
	GO111MODULE=on go mod tidy

tool:
	go get -u golang.org/x/lint/golint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.21.0

fmt:
	gofmt -l -s -w .

lint:
	golangci-lint run -c .golangci.yml

test:
	go test -cover ./...

build:
	go build

bin:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -v -o bin/git-checkout-branch-Darwin-x86_64
	GOOS=linux  GOARCH=amd64 go build -ldflags "-s -w" -v -o bin/git-checkout-branch-Linux-x86_64
