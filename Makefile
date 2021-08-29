-include configs/.env

install:
	go mod tidy

clean:
	go clean -modcache

build:
	go build -o bin/

start:
	make build
	PORT=${PORT} ./bin/gbf-image-proxy

kill:
	kill -9 ${shell lsof -ti :$(PORT)}

test:
	go test ./...

deploy:
	make build
	gcloud app deploy

then:
	echo ""
