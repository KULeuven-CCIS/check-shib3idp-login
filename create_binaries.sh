#!/bin/bash -e
set -xv
mkdir -p binaries
bin="binaries/check-shib3idp-login"
platforms=("windows/amd64" "windows/386" "darwin/amd64" "darwin/386" "linux/amd64" "linux/386")

function build {
    GOOS=$1
    GOARCH=$2
    output="${bin}-${GOOS}-${GOARCH}"
    if [ $GOOS = "windows" ]; then
        output+='.exe'
    fi
    GOOS=$GOOS GOARCH=$GOARCH go build -o $output
    sha512sum $output > $output.sha512
}

for i in ${platforms[@]}; do
    platform_split=(${i//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    build $GOOS $GOARCH
done
