package generator

import (
	"github.com/alex-dev-master/protoc-gen-etcd/generator/metadata"
	ipb "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"log/slog"
)

func (g *generator) processServiceLayer(
	genFile *protogen.GeneratedFile,
	etcdOpts *ipb.EtcdOptions,
) (err error) {
	etcdClientMetadata := metadata.NewEtcdClientMetadata(
		etcdOpts.GetServiceKeyPrefix(),
		g.imports,
	)
	if err = GenerateEtcdClient(genFile, etcdClientMetadata); err != nil {
		slog.Debug("error GenerateEtcdClient")
		return err
	}
	return nil
}
