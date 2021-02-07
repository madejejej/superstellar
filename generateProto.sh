#!/bin/sh

protoc -I=protobuf --go_out=backend/pb protobuf/superstellar.proto 
node_modules/protobufjs/bin/pbjs protobuf/superstellar.proto > webroot/js/superstellar_proto.json
protoc --rust_out=rust-backend/src/ protobuf/superstellar.proto
