#!/bin/bash
set -e

go test $($GOPATH/bin/glide nv)
$GOPATH/bin/goxc
