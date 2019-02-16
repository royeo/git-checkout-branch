.PHONY: tool
tool:
	go get -u golang.org/x/lint/golint

.PHONY: lint
lint:
	go vet ./...
	golint ./...

.PHONY: mod-update
mod-update:
	GO111MODULE=on go get -u -m
	GO111MODULE=on go mod tidy

.PHONY: install
install:
	go install

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build

.PHONY: bin
bin:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -v -o bin/git-checkout-branch-Darwin-x86_64
	GOOS=linux  GOARCH=amd64 go build -ldflags "-s -w" -v -o bin/git-checkout-branch-Linux-x86_64
