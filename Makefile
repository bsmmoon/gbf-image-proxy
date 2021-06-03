-include configs/.env

install:
	go mod tidy
	GOBIN=${GOBIN} go install test.com/go-project

run:
	${GOBIN}/go-project

clean:
	go clean -modcache

build:
	go build

server:
	make build
	./go-project

kill:
	kill -9 ${shell lsof -ti :$(PORT)}

test:
	go test ./...

