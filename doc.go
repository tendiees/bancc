package main

//go:generate protoc --go_opt=paths=source_relative -I=../../proto --go_out=bancc ../../proto/bancc.proto
//go:generate protoc --go-grpc_opt=paths=source_relative -I=../../proto --go-grpc_out=bancc ../../proto/bancc.proto
