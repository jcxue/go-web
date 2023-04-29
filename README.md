# go-web

## Issues
1. bazel does not work for multiple proto files
1. `"@com_google_protobuf//:timestamp_proto"` cannot be found
1. if we used the WORKSPACE file that includes the rules for `protobuf` and `grpc`, it will fail with the error message
saying `go.bzl` is missing. This can be resolved by first not including the these rules, and run `./go-get.sh` to get 
a package, which will produce the `go.bzl` file, and then add back these rules.


