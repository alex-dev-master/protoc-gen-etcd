package extension

import (
	"errors"
	ipb "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GetEtcdValueOptions(fieldOutput *protogen.Field) (*ipb.EtcdValueOptions, error) {
	fieldOutputOpts := fieldOutput.Desc.Options().(*descriptorpb.FieldOptions)
	if !proto.HasExtension(fieldOutputOpts, ipb.E_EtcdValueOptions) {
		return nil, errors.New("etcd_value_options extension missing")
	}
	return proto.GetExtension(fieldOutputOpts, ipb.E_EtcdValueOptions).(*ipb.EtcdValueOptions), nil
}
