#!/bin/bash

set -e

cov_file=/home/ubuntu/coverage.txt

# Get the list of packages.
pkgs=`go list $(/home/ubuntu/.go_workspace/bin/glide novendor)`

echo "mode: count" > $cov_file
for pkg in $pkgs
do
	go test -v -covermode=count $pkg -coverprofile=tmp.cov
	if [ -f tmp.cov ]
	then
		cat tmp.cov | tail -n +2 >> $cov_file
		rm tmp.cov
	fi
done

go tool cover -func $cov_file

mv $cov_file . && bash <(curl -s https://codecov.io/bash)
