#!/bin/bash
set -e
set -x

# run `go test` outside of goxc since there wasn't a clean way to ignore the vendor/ directory otherwise
# https://github.com/laher/goxc/issues/99
go test $($GOPATH/bin/glide nv)
$GOPATH/bin/goxc "$@"
