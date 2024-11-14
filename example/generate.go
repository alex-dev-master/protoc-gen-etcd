//go:generate protoc -I=. -I=./vendor -I=./proto --go_opt=paths=source_relative  --plugin=protoc-gen-etcd=../protoc-gen-etcd --etcd_out=. --etcd_opt=paths=source_relative,logLevelDebug=true --go_out=. proto/example.proto

package main

import (
	_ "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
)
