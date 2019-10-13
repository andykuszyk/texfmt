#!/bin/bash
archs=("amd64" "386")
oss=("darwin" "linux" "windows")

for arch in ${archs[*]}; do
    for os in ${oss[*]}; do
        if [[ "$os" == "windows" ]]; then
            ext=".exe"
        else
            ext=""
        fi
        echo "Building texfmt for os/arch: $os / $arch"
        GOOS=$os GOARCH=$arch go build -o texfmt$ext ./cmd/texfmt/main.go
        tar -cvf texfmt-$os-$arch.tar texfmt$ext
        rm texfmt$ext
    done
done
