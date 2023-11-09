# Go gRPC Gateway Server Example

Start here - [grpc-ecosystem.github.io/grpc-gateway/](https://grpc-ecosystem.github.io/grpc-gateway/)

### Notes

- Initialise using go mod init

- Create the proto file in the `proto` folder

- `buf.yaml` & `buf.gen.yaml` needed for `buf` to generate the bindings

- Use `buf generate --path ./proto/helloworld/hello_world.proto` to generate the binding

- `*.gw.go` - gateway

- Create gRPC Server

- connection to the gRPC Server

- mux + handler

- HTTP Server - gRPC Gateway

## Setup

1. Run the gRPC Server and the gateway

```
go run main.go
```

```
2023/11/09 11:41:40 Serving gRPC on 0.0.0.0:8080
2023/11/09 11:41:40 Serving gRPC-Gateway on http://0.0.0.0:8090
```

2. Test using curl

```
curl -X POST -k http://localhost:8090/v1/example/echo -d '{"name": "Meow"}'
```

```
{"message":"Meow world"}
```

## Installations Needed

- gRPC stuff

```
$ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- buf tool

```
# Substitute BIN for your bin directory.
# Substitute VERSION for the current released version.
BIN="/usr/local/bin" && \
VERSION="1.27.2" && \
sudo curl -sSL \
"https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
-o "${BIN}/buf" && \
sudo chmod +x "${BIN}/buf"
```
