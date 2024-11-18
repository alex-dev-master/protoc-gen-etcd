package extensions

import (
	ipb "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func GetEtcdOptions(s *protogen.Service) *ipb.EtcdOptions {
	sOpt := s.Desc.Options().(*descriptorpb.ServiceOptions)
	if proto.HasExtension(sOpt, ipb.E_EtcdOptions) {
		return proto.GetExtension(sOpt, ipb.E_EtcdOptions).(*ipb.EtcdOptions)
	}
	return nil
}
