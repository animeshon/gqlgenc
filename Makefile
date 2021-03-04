MAKEFLAGS=--no-builtin-rules --no-builtin-variables --always-make

fmt:
	gofumports -local github.com/animeshon/gqlgenc -w .

lint:
	golangci-lint cache clean && golangci-lint run

test:
	go test -v ./...
