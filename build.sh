#!/bin/sh
export GOPROXY=https://goproxy.io
export CGO_ENABLED=0
echo "building..."
go build -v
echo "done"