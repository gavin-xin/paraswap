.PHONY : tools mock
tools:
	# fill it

mod:
	go mod tidy

build: mod
	go build -o paraswap ./cmd/main.go

test:
	go test `go list ./...`

cover:
	go test -cover `go list ./...`
