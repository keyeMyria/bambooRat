#!/bin/bash
if [ $# -ne 1 ]; then
	echo "dir name null"
	exit 0
fi
set -e
mkdir -p $GOPATH/src/github.com/$1
cd $GOPATH/src/github.com/$1
touch main.go

dep init
# 下载protobuff
dep ensure -add github.com/golang/protobuf/protoc-gen-go
# 下载micro组件
dep ensure -add github.com/micro/protoc-gen-micro
tree -L 2

set +e


