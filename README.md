# :zero: zkp-extended

An implementation of an extended ZKP Protocol which adds 1-factor authentication

## :hatching_chick: Getting Started

To get started with this project, please install the following pre-requisites:

* Go, version 1.19
* Docker
* Docker Compose
* Make
* protoc ([Docs](https://grpc.io/docs/protoc-installation/))

To check whether you meet all pre-requisites, you can run the
[setup](./scripts/setup) script:

```
./scripts/setup
```

Now, you can follow the guide that best suits your use case:

* Building the project: [locally](#building-the-go-binary) or from [Docker](#building-the-docker-image)
* [Running the application](#running-the-application-from-docker)
* [Provisioning the application on AWS](#provisioning-the-application-on-aws)

## Building the Go Binary

To build this project, run the `make build` target:

```
make build
```

This target builds 4 binaries, 2 for each of the supported platforms:
Linux/amd64 and MacOS/arm64. The binaries are written to `bin/`.

## Building the Docker Image

To build Docker images for this project, run the `make build-docker` target.

```
make build-docker
```

This target builds 2 images tagged `zkp-client:latest` and `zkp-server:latest`
for the client and server applications, respectively.

## Running the Application from Docker

To run the application, run the `make run-docker` target:

```
make run-docker
```

This target uses `docker-compose` to fire up the client and server application,
each in their own container.

## Provisioning the Application on AWS

To provision the application and supporting infrastructure on AWS, run the `make
run-terraform` target:

```
make run-terraform
```

This target uses Terraform to launch 2 EC2 instances in the same VPC/subnet. The
instance userdata is configured to install the full zkp-extended toolchain
including Golang and Docker. After building the client and server images, on
their respective instances, the Terraform ensures the client `login` command
exits `0`.

:bulb: _The Terraform stores its state locally. All caveats apply._
