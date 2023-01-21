all: build

build: build-proto build-client build-generator build-server

build-client:
	@GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/nil-client cmd/client/main.go

build-generator:
	@GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/nil-generator cmd/generator/main.go

build-proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/proto/nil.proto

build-server:
	@GOARCH=amd64 GOOS=linux go build -mod vendor -o bin/nil-server cmd/server/main.go

.PHONY: run-docker
run-docker:
	sudo docker-compose -f infra/docker/docker-compose.yaml up --build

.PHONY: run-terraform
run-terraform:
	terraform -chdir=infra/terraform init
	terraform -chdir=infra/terraform plan -out=plan.json
	terraform -chdir=infra/terraform apply ./plan.json

.PHONY: test
test:
	@go test -v ./...

.PHONY: test-e2e
test-e2e:
	./scripts/test-e2e-registration
