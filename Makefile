-include configs/.env

install:
	go mod tidy
	GOBIN=${GOBIN} go install test.com/go-project

clean:
	go clean -modcache

build:
	go build

run:
	${GOBIN}/go-project

test:
	go test ./...

