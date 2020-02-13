# Use `go list` to enumerate the full path of all go files
# See `go list -f '{{range .GoFiles}}{{$.Dir}}/{{.}}{{end}}' ./...`
BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}{{end}}' ./...)
ENTRY_POINT=./cli/go-boilerplate.go
BINARY_NAME=go-boilerplate

bin/${BINARY_NAME}: $(BUILD_FILES)
	@go build -o "$@" ${ENTRY_POINT}

# see `go tool dist list` for full list
dist:
	GOOS=darwin GOARCH=amd64 go build -o dist/${BINARY_NAME}-darwin-amd64 ${ENTRY_POINT}
	GOOS=linux GOARCH=amd64 go build -o dist/${BINARY_NAME}-linux-amd64 ${ENTRY_POINT}
	GOOS=windows GOARCH=amd64 go build -o dist/${BINARY_NAME}-windows-amd64 ${ENTRY_POINT}

run:
	go run ${ENTRY_POINT}

test:
	go test ./...

.PHONY:	test dist
