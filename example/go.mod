module github.com/alex-dev-master/protoc-gen-etcd/example

go 1.23.0

require google.golang.org/protobuf v1.35.1 // indirect

require github.com/alex-dev-master/protoc-gen-etcd v0.0.0-20241114144525-8628fda2361d

require (
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
)

//replace github.com/alex-dev-master/protoc-gen-etcd/proto => ../
