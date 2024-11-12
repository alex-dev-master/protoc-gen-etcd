//go:generate protoc -I=. -I=../internal/proto -I=./vendor -I=./proto --go_opt=paths=source_relative  --plugin=protoc-gen-etcd=../protoc-gen-etcd --etcd_out=. --etcd_opt=paths=source_relative --go_out=. proto/example.proto

package main
