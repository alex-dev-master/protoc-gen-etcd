package generator

import (
	"errors"
	"github.com/alex-dev-master/protoc-gen-etcd/generator/extensions"
	"github.com/alex-dev-master/protoc-gen-etcd/generator/metadata"
	ipb "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"log/slog"
)

func processMethodLayer(g *protogen.GeneratedFile, method *protogen.Method) (err error) {
	opts := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !proto.HasExtension(opts, ipb.E_EtcdKeyOptions) {
		return errors.New("etcd_key_options extension missing")
	}
	etcdKeyOptions := proto.GetExtension(opts, ipb.E_EtcdKeyOptions).(*ipb.EtcdKeyOptions)

	etcdKeyParamOptions := make(map[string]*metadata.FieldWithEtcdKeyParamOptions)
	for _, fieldInput := range method.Input.Fields {
		fieldInputOptsExt, errE := extensions.GetEtcdKeyParamOptions(fieldInput)
		if errE != nil {
			slog.Debug(errE.Error())
			continue
		}

		f := &metadata.FieldWithEtcdKeyParamOptions{
			GoName:            fieldInput.GoName,
			ProtoName:         fieldInput.Desc.JSONName(),
			Kind:              fieldInput.Desc.Kind(),
			Cardinality:       fieldInput.Desc.Cardinality(),
			FieldInputOptsExt: fieldInputOptsExt,
			GoTypeStr:         metadata.ProtoTypeToGoTypeField(fieldInput.Desc.Kind()),
		}

		key := f.ProtoName
		if f.FieldInputOptsExt.GetTargetName() != "" {
			key = f.FieldInputOptsExt.GetTargetName()
		}
		etcdKeyParamOptions[key] = f
	}

	etcdMethodMetadata := metadata.NewEtcdMethodMetadata(
		etcdKeyOptions,
		etcdKeyParamOptions,
		method.Output.GoIdent,
		method.Input.GoIdent)

	if err = GenerateEtcdMethodGet(g, etcdMethodMetadata); err != nil {
		return err
	}

	return nil
}
