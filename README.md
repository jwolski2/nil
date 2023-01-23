# :zero: nil

A Go implementation of an extended ZKP protocol which adds 1-factor authentication

## :hatching_chick: Getting Started

Hello, and welcome to the repository! This exercise was completed over the
course of 3 or 4 days. I broke implementation down into multiple stages:

1. Set development environment: basic Go code, Makefile, Docker.
2. Implement the protocol: gRPC, crypto.
3. Write tests.
4. Documentation.

I'm fairly proficient in all parts but the crypto part. I used several resources
to accomplish what's there, and even still, I can't confidently stand behind or
reason deeply about the math. Though, I had fun putting everything together.
Thanks for your review and consideration.

To get started, refer to the guide that best suits your use case:

* [Reviewer's Guide](#student-reviewers-guide) - for those that want to learn about the repo, implementation and thought process
* [Developer's Guide](#construction_worker-developers-guide) - for those that want to build, run and test the implementation
* [User's Guide](#bust_in_silhouette-users-guide) - for those that want to use the client CLI to register and login

## :student: Reviewer's Guide

### What's in this repo?

In this repo, you'll find:

* Code for a ZKP client, server and value generator
* Dockerfiles for the client (CLI) and server
* A docker-compose file which launches the server and runs client commands
* Unit tests for the crypto package
* Functional tests for the register/login flow
* An end-to-end test of the client and server using Docker Compose
* Pre-computed values for p, g, h and q under [data](./data)

**:1234: About the code**

The code is organized in 2 main directories: [cmd](./cmd) and [pkg](./pkg).

The **cmd** dir contains the `main` functions for each of the client CLI, the
server and value generator.

The **pkg** dir contains the authentication protocol implementation and consists
of [client](./pkg/client), [server](./pkg/server), [crypto](./pkg/crypto) and
[proto](./pkg/proto) packages.

**:whale: About the Docker setup**

The Docker setup is comprised of a Dockerfile for the
[client](./infra/docker/Dockerfile.client) and the
[server](./infra/docker/Dockerfile.server) and
a [docker-compose.yaml](./infra/docker/docker-compose.yaml) file.

The Docker images are built on Alpine Linux and use Go 1.19.

The Docker Compose setup starts the server first, waits until the server is
reachable, starts the client and executes the register/login flow.

**:test_tube: About the tests**

There are 3 types of tests: unit, functional and end-to-end.

The [unit tests](./pkg/crypto/crypto_test.go) are written in Go and test that r1
and r2 values can be verified outside the context of the client/server
application.

The [functional tests](./pkg/client/client_test.go) are written in Go and test
the client/server interaction at the functional-level by ensuring the
correctness of the registration and login processes.

The [end-to-end test](./scripts/test-e2e) is written as a shell script which
wraps Docker Compose and ensures the correctness of the register/login flow by
running the client and server Docker images.

**:canned_food: About the data**

There's a [data](./data) dir which contains JSON files of pre-computed values for
p, g, h and q. The files are generated by the [generator](./cmd/generator).

I know that these values are meant to be 'public'. And I know that they are not
meant to be shared by multiple users / across multiple executions of the
protocol. But I was unable to implement a method of generating them efficiently
at run-time. The most practical thing I thought to do was generate them ahead of
time and have them loaded by the client/server at startup.

These data files double as test fixtures for the functional tests mentioned
above.

## :construction_worker: Developer's Guide

As a developer or tester, you may want to build and run tests for the
client/server applications. Here's how you can accomplish that.

First, ensure you've got the requisite software installed on your system by
running the Make target:

```
make setup
```

If successful, you can expect the output to look like:

```
👋 Welcome, nil users!

Checking for pre-requisite software...

✔️ Found docker
✔️ Found docker-compose
✔️ Found go
✔️ Found make
✔️ Found terraform
✔️ Found protoc plugin: protoc-gen-go
✔️ Found protoc plugin: protoc-gen-go-grpc

🚀 All is good. Refer to README.md for further instruction!
```

### Building the Client/Server

To build the client/server (and value generator), run the Make target:

```
make build
```

This target builds Linux/amd64 binaries and deposits them in `bin/`:

```
ls bin
nil-client  nil-generator  nil-server
```


### Running Tests

There are 3 types of tests that have been implemented: unit, functional and
end-to-end.

The unit and functional tests test the crypto and the registration and login
logic between client and server. To run them, run the Make target:

```
make test
```

The end-to-end tests test the registration and login logic using a Dockerized
client and server. To run them, run the Make target:

```
make test-e2e
```

If successful, the output should look like this:

```
Please wait. This may take several seconds...

✅ test_succeeds passed
✅ test_fails passed

2 passed, 0 failed, 2 total.
```

## :bust_in_silhouette: User's Guide

As a user, you may want to run the client and server applications. First,
familarize yourself with their usage. Then, choose between running
[on-host](#running-on-host) or [in a container](#running-in-docker).

__Server usage__

```
./bin/nil-server --help
NAME:
   nil-server - A Nil server

USAGE:
   nil-server [global options] command [command options] [arguments...]

COMMANDS:
   start    Start the server
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

__Client usage__

```
./bin/nil-client --help
NAME:
   nil-client - A CLI for the Nil server

USAGE:
   nil-client [global options] command [command options] [arguments...]

COMMANDS:
   login     Login with the Nil server
   register  Register with the Nil server
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --hostname value  Hostname for nil-server (default: "nil-server")
   --port value      Port for nil-server (default: 9999)
   --help, -h        show help (default: false)
```
Then, choose between running on-host or in Docker.

### Running on-host

If running the applications on-host, start the server first:

```
./bin/nil-server start
{"level":"info","time":"2023-01-21T19:15:50+01:00","message":"Server is listening on :9999"}
```

Then, register via the client by running:

```
./bin/nil-client --hostname localhost register wolski 8675309
User has been registered!
```

Then, complete the authentication process by logging in:

```
./bin/nil-client --hostname localhost login wolski 8675309
Login successful. Session ID is 1aacf8d5312b827f82f7d69f806fc1f4.
```

### Running in Docker

To run the same in Docker, run the Make target:

```
make run-docker
```

This target uses Docker Compose to build the client and server images. It then
starts the server container and runs the client container to completion after
executing a single register/login flow. After a slew of build output, you should
also see:

```
Starting docker_nil-server_1 ... done
Starting docker_nil-client_1 ... done
Attaching to docker_nil-server_1, docker_nil-client_1
nil-client_1  | User has been registered!
nil-client_1  | Login successful. Session ID is d896619fd6b35e7dd07a6155cd3b3d2.
docker_nil-client_1 exited with code 0
```

### Error Scenarios

Once you get the client and server running, you should know a bit about the
rules of the protocol:

* **If you register the same user multiple times**, an error is reported saying
the user already exists.
* **If you attempt to login with the same user multiple times**, an error is
reported saying there is already an active session.
* **If you login with the wrong password**, an error is reported saying the
password is incorrect.
