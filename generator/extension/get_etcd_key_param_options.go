package extension

import (
	"errors"
	ipb "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GetEtcdKeyParamOptions(fieldInput *protogen.Field) (*ipb.EtcdKeyParamOptions, error) {
	fieldInputOpts := fieldInput.Desc.Options().(*descriptorpb.FieldOptions)
	if !proto.HasExtension(fieldInputOpts, ipb.E_EtcdKeyParamOptions) {
		return nil, errors.New("etcd_key_param_options extension missing")
	}
	return proto.GetExtension(fieldInputOpts, ipb.E_EtcdKeyParamOptions).(*ipb.EtcdKeyParamOptions), nil
}
