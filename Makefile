all: build

build: build-darwin build-linux

build-darwin:
	GOARCH=arm64 GOOS=darwin go build -mod vendor -o bin/zkp-client-arm64 cmd/client/main.go
	GOARCH=arm64 GOOS=darwin go build -mod vendor -o bin/zkp-server-arm64 cmd/server/main.go

build-linux: build-linux-client build-linux-server

build-linux-client:
	GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/zkp-client cmd/client/main.go

build-linux-server:
	GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/zkp-server cmd/server/main.go

.PHONY: build-docker
build-docker:
	docker build -t zkp-client -f docker/Dockerfile.client .
	docker build -t zkp-server -f docker/Dockerfile.server .

.PHONY: run-docker
run-docker: build-docker
	sudo docker-compose -f docker/docker-compose.yaml up
