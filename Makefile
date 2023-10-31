.PHONY: test

BIN=go

test:
	${BIN} test -race -v ./...
watch-test:
	reflex -t 50ms -s -- sh -c '${BIN} test -race -v ./...'

coverage:
	${BIN} test -v -coverprofile=cover.out -covermode=atomic .
	${BIN} tool cover -html=cover.out -o cover.html

bench:
	${BIN} test -benchmem -count 3 -bench ./...
watch-bench:
	reflex -t 50ms -s -- sh -c '${BIN} test -benchmem -count 3 -bench ./...'

tools:
	${BIN} install github.com/cespare/reflex@latest
#	${BIN} install github.com/rakyll/gotest@latest
#	${BIN} install github.com/psampaz/go-mod-outdated@latest
#	${BIN} install github.com/jondot/goweight@latest
	${BIN} install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
#	${BIN} get -t -u golang.org/x/tools/cmd/cover
#	${BIN} install github.com/sonatype-nexus-community/nancy@latest
#	${BIN} mod tidy

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...
lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...

generate:
	go run cmd/generate_tuples.go -output tuples.go