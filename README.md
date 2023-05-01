# go-web

## Issues
1. bazel does not work for multiple proto files
1. `"@com_google_protobuf//:timestamp_proto"` cannot be found
1. if we used the WORKSPACE file that includes the rules for `protobuf` and `grpc`, it will fail with the error message
saying `go.bzl` is missing. This can be resolved by first not including the these rules, and run `./go-get.sh` to get 
a package, which will produce the `go.bzl` file, and then add back these rules.

## Setup
### Install protoc
[instructions](https://grpc.io/docs/protoc-installation/https://grpc.io/docs/protoc-installation/)
```bash
# linux
sudo apt install -y protobuf-compiler
protoc --version

# macos
brew install protobuf
protoc --version
```

### Install protobuf golang plugins
[instructions](https://grpc.io/docs/languages/go/quickstart/)
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

### Use protoc to generate go code
```bash
protoc --proto_path=proto proto/*.proto  --go_out=paths=source_relative:pb
```


