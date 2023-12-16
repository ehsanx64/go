# grpc-helloworld

These sources are the simplest usage of gRPC and is based on [grpc-go/examples/helloworld](https://github.com/grpc/grpc-go/tree/master/examples/helloworld).

## How to Run

First of all make sure the `protoc` compiler is installed and accessible. And use the provided Makefile as simply as:

```bash
# Compile the helloworld.proto file and generate Go code for it
make protoc

# Make sure dependencies are install
make tidy

# Run the server
make run-server

# Run the client (this will not pass the name parameter to the client so defaultName will be used)
make run-client

# Run the client (this will pass a name parameter to the client)
make run-client-gopher
```
