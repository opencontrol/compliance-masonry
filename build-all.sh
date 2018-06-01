#!/bin/bash

GOPATH=$1
GO=$2
package=$3
if [[ $# -lt 3 ]]; then
  echo "usage: $0 <gopath> <gobinary> <package-name>"
  exit 1
fi

set -x

export GOPATH=${GOPATH}

package_split=(${package//\// })
package_name=${package_split[-1]}

platforms=("windows/amd64" "windows/386" "darwin/amd64" "darwin/386" "linux/386" "linux/amd64" "linux/arm" "linux/arm64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name="./build/${GOOS}_${GOARCH}/${package_name}"
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi  

    GOOS=$GOOS GOARCH=$GOARCH ${GO} build -o $output_name $package
done
