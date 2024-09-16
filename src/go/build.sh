#!/bin/bash

if [ "$DUCKDB_PLATFORM" == "osx_amd64" ]; then
GOOS=darwin
GOARCH=amd64
elif [ "$DUCKDB_PLATFORM" == "osx_arm64" ]; then
GOOS=darwin
GOARCH=arm64
elif [ "$DUCKDB_PLATFORM" == "linux_amd64_gcc4" ]; then
GOOS=linux
GOARCH=amd64
elif [ "$DUCKDB_PLATFORM" == "linux_arm64" ]; then
CC=aarch64-linux-gnu-gcc
GOOS=linux
GOARCH=arm64
elif [ "$DUCKDB_PLATFORM" == "linux_amd64" ]; then
GOOS=linux
GOARCH=amd64
fi

CURRENT_DIR=$(pwd)
SRC_DIR="$1"
DUCKDB_LIB_PATH=$( cd "$2" && pwd -P )
(cd "$SRC_DIR" && go mod tidy && CC=$CC GOOS=$GOOS GOARCH=$GOARCH CGO_CFLAGS="-I${DUCKDB_LIB_PATH}/" CGO_LDFLAGS="-L${DUCKDB_LIB_PATH}" CGO_ENABLED=1 GOWORK=off go build -buildmode=c-archive -o "$CURRENT_DIR" .)