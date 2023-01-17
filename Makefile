all: build

build:
	GOARCH=arm64 GOOS=darwin go build -mod vendor -o bin/zkp-client-arm64 cmd/client/main.go
	GOARCH=arm64 GOOS=darwin go build -mod vendor -o bin/zkp-server-arm64 cmd/server/main.go
	GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/zkp-client cmd/client/main.go
	GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/zkp-server cmd/server/main.go
