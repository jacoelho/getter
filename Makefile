# disable default rules
.SUFFIXES:
MAKEFLAGS+=-r -R

.PHONY: generate
generate:
	go generate -v ./...

.PHONY: test
test:
	go test -race -v ./...

.PHONY: build
build:
	go build -v ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: install
install:
	go install ./...

.PHONY: ci-tidy
ci-tidy:
	go mod tidy
	git status --porcelain go.mod go.sum || { echo "Please run 'go mod tidy'."; exit 1; }
