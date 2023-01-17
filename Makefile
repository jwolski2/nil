all: build

build:
	GOARCH=arm64 GOOS=darwin go build -mod vendor -o bin/zkp-extended-arm64 main.go
	GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/zkp-extended main.go
