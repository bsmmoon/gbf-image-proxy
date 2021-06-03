-include configs/.env

install:
	go mod tidy

clean:
	go clean -modcache

build:
	go build -o bin/

server:
	make build
	./bin/go-project

kill:
	kill -9 ${shell lsof -ti :$(PORT)}

test:
	go test ./...

