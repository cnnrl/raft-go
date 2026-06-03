#!/usr/bin/env bash

protoc \
  --go_out=src/ \
  --go-grpc_out=src/ \
  src/proto/raft.proto
