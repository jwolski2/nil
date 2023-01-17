all: build

build: build-darwin build-linux

build-darwin: build-proto
	@GOARCH=arm64 GOOS=darwin go build -mod vendor -o bin/zkp-client-arm64 cmd/client/main.go
	@GOARCH=arm64 GOOS=darwin go build -mod vendor -o bin/zkp-server-arm64 cmd/server/main.go

build-linux: build-proto build-linux-client build-linux-server

build-linux-client:
	@GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/zkp-client cmd/client/main.go

build-linux-server:
	@GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/zkp-server cmd/server/main.go

build-proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/proto/zkp_extended.proto

.PHONY: build-docker
build-docker:
	@docker build -t zkp-client -f infra/docker/Dockerfile.client .
	@docker build -t zkp-server -f infra/docker/Dockerfile.server .

.PHONY: run-docker
run-docker: build-docker
	sudo docker-compose -f infra/docker/docker-compose.yaml up

.PHONY: run-terraform
run-terraform:
	terraform -chdir=infra/terraform init
	terraform -chdir=infra/terraform plan -out=plan.json
	terraform -chdir=infra/terraform apply ./plan.json
