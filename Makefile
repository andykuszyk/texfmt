dep:
	dep ensure

install: build
	cp texfmt ~/go/bin

build: dep
	go build -o texfmt ./cmd/texfmt/main.go

test: build
	go test -v ./...

watch:
	find . | grep -e go$ | entr -c make test

package:
	./scripts/package.sh
