package generator

import (
	"github.com/alex-dev-master/protoc-gen-etcd/generator/extensions"
	"github.com/alex-dev-master/protoc-gen-etcd/generator/metadata"
	ipb "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"log/slog"
)

func (g *generator) processMethodLayer(genFile *protogen.GeneratedFile, method *protogen.Method, keyPrefix string) (err error) {
	var etcdKeyOptions *ipb.EtcdKeyOptions
	if etcdKeyOptions, err = extensions.GetEtcdKeyOptions(method); err != nil {
		return err
	}

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

	var etcdMethodMetadata *metadata.EtcdMethodMetadata
	if etcdMethodMetadata, err = metadata.NewEtcdMethodMetadata(
		&metadata.NewEtcdMethodMetadataRequest{
			EtcdKeyOptions:      etcdKeyOptions,
			EtcdKeyParamOptions: etcdKeyParamOptions,
			ValueOfKey:          method.Output.GoIdent,
			InputRequest:        method.Input.GoIdent,
			KeyPrefix:           keyPrefix,
			MethodName:          method.GoName,
			Imports:             g.imports,
		},
	); err != nil {
		slog.Debug("error NewEtcdMethodMetadata")
		return err
	}

	if err = GenerateEtcdMethodGet(genFile, etcdMethodMetadata); err != nil {
		return err
	}

	return nil
}
