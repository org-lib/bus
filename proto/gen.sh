#!/usr/bin/env bash

echo "Processing..."

GOPATH=${GOPATH:-$(go env GOPATH)}
GOBIN=${GOBIN:-$(go env GOBIN)}

if [[ $GOBIN == "" ]]; then
  GOBIN=${GOPATH}/bin
fi
go install -v google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 复制去 GOROOT

echo "Use protoc-gen-go and protoc-gen-go-grpc in $GOBIN."

protoc --go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
--plugin=protoc-gen-go=${GOBIN}/protoc-gen-go \
--plugin=protoc-gen-go-grpc=${GOBIN}/protoc-gen-go-grpc \
hello.proto
#protoc -I . --go_out=plugins=grpc:. ./hello.proto
if [ $? -eq 0 ]; then
  echo "Generated successfully."
fi
