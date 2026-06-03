#!/usr/bin/env bash

set -e 

protoc \
  --go_out=src/ \
  --go-grpc_out=src/ \
  src/proto/raft.proto

mkdir -p dist

go build -o dist/raft ./src
