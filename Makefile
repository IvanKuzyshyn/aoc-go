.PHONY: build lint static-checks

tools := golangci-lint staticcheck
$(foreach bin,$(tools),\
    $(if $(shell command -v $(bin) 2> /dev/null),,$(error Please install `$(bin)`)))

build:
	go build -o aoc .

lint: static-checks
	$(info Running linter...)
	golangci-lint run

static-checks:
	$(info Running static checks...)
	staticcheck ./...

test:
	$(info Running tests...)
	go test ./... --race -count=1;
