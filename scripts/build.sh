#!/bin/bash

CGO_ENABLED=0

if [ "${1}" == "windows" ]
then
    go build -v -ldflags "-s -w" -o builds/webp-server-${1}-${2}.exe
fi


if [ "${1}" == "osx" ]
then
    go build -v -ldflags "-s -w" -o builds/webp-server-darwin-${2}
else
    go build -v -ldflags "-s -w" -o builds/webp-server-${1}-${2}
fi

if [ "${1}" == "linux" ]; then
    export CC=x86_64-linux-musl-gcc
    if [ $2 == "amd64" ]; then
        export CC=x86_64-linux-musl-gcc

    fi

    if [ $2 == "386" ]; then
        export CC=i486-linux-musl-gcc
    fi

    if [ $2 == "arm64" ]; then
        export CC=aarch64-linux-musl-gcc
    fi

    if [ $2 == "arm" ]; then
        export CC=arm-linux-musleabi-gcc
    fi

    go build -ldflags "-s -w"  builds/webp-server-${1}-${2}
fi

for file in builds/*
do
    sha256sum ${file} > ${file}.sha256
done

echo "build done!"
ls builds