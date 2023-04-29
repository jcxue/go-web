 #!/bin/bash

set -e

GO=./bazel-go-web/external/go_sdk/bin/go
cd $(dirname "${BASH_SOURCE[0]}")

./bazel build @go_sdk//...
GO111MODULE=on $GO get -v $@
GO111MODULE=on ./bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=go.bzl%go_repositories -prune