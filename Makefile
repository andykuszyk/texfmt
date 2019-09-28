build:
	go build ./cmd/texfmt/main.go

test:
	go test -v ./...

watch:
	find . | grep -e go$ | entr -c make test
