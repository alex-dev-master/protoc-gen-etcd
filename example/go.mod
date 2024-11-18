module github.com/alex-dev-master/protoc-gen-etcd/example

go 1.23.0

require google.golang.org/protobuf v1.35.1 // indirect

require github.com/alex-dev-master/protoc-gen-etcd v0.0.0-20241118125410-c7c784c977ae

require (
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
)

//replace github.com/alex-dev-master/protoc-gen-etcd/proto => ../
