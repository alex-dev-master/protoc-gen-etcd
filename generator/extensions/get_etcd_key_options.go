package extensions

import (
	"errors"
	ipb "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GetEtcdKeyOptions(method *protogen.Method) (*ipb.EtcdKeyOptions, error) {
	opts := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !proto.HasExtension(opts, ipb.E_EtcdKeyOptions) {
		return nil, errors.New("etcd_key_options extension missing")
	}
	return proto.GetExtension(opts, ipb.E_EtcdKeyOptions).(*ipb.EtcdKeyOptions), nil
}
