# planet-express

This repo contains my solutions for [devops-se-homework](https://github.com/DivvyPayHQ/devops-se-homework).
I enjoyed working on this project. This project involved working with the following technologies:
* Golang
* GRPC
* GraphQL
* Docker
* Kubernetes
* Helm

## Run locally
* build the project ```make build```
* start the grpc server ```./ship/ship```. GRPC server will listen on port ```10000```.
* start the graphql server ```./headquarters/headquarters```. Graphql server will listen on port ```8080``` by default. To run on a different port ```./headquarters/headquarters --graphql_addr <desired-port>```.

## Run with docker-compose
```docker-compose up -d```. This will spin up two containers one for grpc server and other for graphql server.

## Run grpc server in a local kubernetes cluster
```helm install shipserver --kubeconfig <"path/to/kubeconfig/file"> ./shipserver_helm/```

-----

## Description (reproducing the description here solely for the sake of convenience.)

You are contracted to refactor Professor Farnsworth's legacy client/server code, and you decide to spice weasel (BAM!) things up with K8s, gRPC, and GraphQL.
The professor wants to run a gateway service at the Planet Express headquarters that can communicate with Planet Express ship via gRPC.
Furthermore, he wants a nice React UI to show data about the ship, deliveries, crew members, etc that will communicate with the gateway service via GraphQL.

In this architecture the ship is a backend service, the headquarters is a api gateway, and the frontend is a React SPA.

## Setup

When you are done shopping for lightspeed briefs you decide to setup the new project. You realize that some of it is already set up for you. The `./headquarters` directory contains the gateway/client code and the
`./ship` directory contains the server code.

1. Create a gitlab repo called `planet-express`
2. Setup Your Dev Environment
    - Install go (https://golang.org/doc/install)
    - Install protoc (https://grpc.io/docs/protoc-installation/)
    - Follow the getting started guide to install the protoc go plugin (https://grpc.io/docs/languages/go/quickstart/)
3. `go get google.golang.org/grpc`
4. `go get github.com/golang/protobuf`
5. Copy the files in this repo over to your new repo
6. Run `go mod init gitlab.com/<gitlab-username>/planet-express`
7. Your go mod should look something like the included `_go.mod` file (delete this file when done)
8. Fix the import statements in `./ship/main.go` and `./headquarters/main.go` (the lines commented like so `// pb "gitlab.com/<gitlab-username>/planet-express/ship/pkg/planetexpress"`)
9. Validate that you can build and genarate go code by running `make`
10. Validate you have the executables `./ship/ship` and `./headquarters/headquarters`
11. Run the server `./ship/ship` and the client `./headquarters/headquarters` in separate terminal windows and confirm the output.

## Generate Go Code

The protoc compiler with the go plugin will autogenerate client and server grpc stubs for you. Make sure this runs successfully with `make`

You should try changing the protobuf files in `./proto` and re-run `make`. The generated files are put in `./ship/pkg/planetexpress` and `./headquarters/pkg/planetexpress`

## Build

A makefile has been provided for you that will generate go code with the protoc compiler and build the executables. Run `make` and validate that you have executables at `./ship/ship` and `./headquarters/headquarters`

## Your Assignment

Time Estimate: 3 hours

The basic client and server code has been provided for you. Your job is to do the following:

- [x] 1. Implement `crew.proto` and `delivery.proto` and the functions listed in `planet_express_service.proto`
    - You should experiment with different protobuf types for the crew and delivery resources. Each resource should demonstrate an understanding of protobuf.
- [x] 2. Refactor `ship.proto` to use the `crew` and `delivery` resources.
- [x] 3. Implement the rpc functions (defined in `planet_express_service.proto`) in both the client (`./headquarters`) and server (`./ship`), following the `GetShip` example. You do not need to add a database. Experiment with using data contained in the gRPC request.
- [x] 4. Output data about the planet express ship, deliveries, and crew members to a json file called `planet_express.json`. Feel free to get creative with your data and protobuf resources.
- [x] 5. Add a dockerfile to `./ship` so that we can run the server with docker.

### Above and Beyond

Time Estimate: 4 hours

- [x] 1. Provide a basic helm template for the `./ship` server. This should contain a kubernetes deployment and service. We should be able to port-forward your service to call the `./ship` rpc functions.
- [x] 2. Add additional protobuf resources like `ShipEngine`. You can get creative with this.
- [x] 3. Add graphql support to the api gateway (headquarters). A good library for this is https://github.com/graph-gophers/graphql-go. Map a graphql endpoint/resource to a gRPC call, like for GetShip.
A GraphQL client should be able to query data from the headquarters which in turn resolves the query by getting data from the ship backend via gRPC.

### Bite my shiny metal a$$ (extra, extra)

Time Estimate: 4 hours

- [ ] 1. Add a simple React SPA that displays data about the ship using GraphQL to communicate with the gateway (headquarters service). You can put this in the `./dashboard` directory.
- [ ] 2. Add some go unit tests to the ship server.

** These would really impress us here at Planet Express! **
