all: build

build: build-proto build-client build-generator build-server

build-client:
	@GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/nil-client cmd/client/main.go

.PHONY: build-docker
build-docker:
	@docker build -t nil-client -f infra/docker/Dockerfile.client .
	@docker build -t nil-server -f infra/docker/Dockerfile.server .

build-generator:
	@GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/nil-generator cmd/gh-generator/main.go

build-proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/proto/nil.proto

build-server:
	@GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/nil-server cmd/server/main.go

.PHONY: run-docker
run-docker: build-docker
	sudo docker-compose -f infra/docker/docker-compose.yaml up

.PHONY: run-terraform
run-terraform:
	terraform -chdir=infra/terraform init
	terraform -chdir=infra/terraform plan -out=plan.json
	terraform -chdir=infra/terraform apply ./plan.json
