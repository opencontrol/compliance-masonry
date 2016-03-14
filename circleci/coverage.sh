#!/bin/bash

cov_file=/home/ubuntu/coverage.out

# Get the list of packages.
pkgs=`go list ./...`

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

/home/ubuntu/.go_workspace/bin/goveralls -coverprofile=$cov_file -service=circle-ci -repotoken=$COVERALLS_TOKEN
