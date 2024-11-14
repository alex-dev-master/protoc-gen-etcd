package generator

import (
	"github.com/alex-dev-master/protoc-gen-etcd/entities"
	ipb "github.com/alex-dev-master/protoc-gen-etcd/pkg/proto"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"log/slog"
	"os"
)

func Run(cfg *entities.Config, plugin *protogen.Plugin) (err error) {
	logLevel := slog.LevelInfo
	if *cfg.LogLevelDebug {
		logLevel = slog.LevelDebug
	}

	textHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: false,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})
	slog.SetDefault(slog.New(textHandler))

	slog.Debug("Protogen plugin called with following files to be generated", "files", plugin.Request.FileToGenerate)

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		g := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".etcd.go", file.GoImportPath)
		GenerateHeader(g, file)

		for _, message := range file.Messages {
			etcdKeyParams := getEtcdKeyParamsParams(message)
			if etcdKeyParams == nil {
				slog.Debug("Пропущено сообщение %s, так как etcd_key_template не задано\n", message.GoIdent.GoName)
				continue
			}
			slog.Debug("etcdKeyParams", etcdKeyParams)

			var meta *EtcdMetadata
			if meta, err = GetEtcdMetadataFromMessage(message, etcdKeyParams); err != nil {
				slog.Debug("error GetEtcdMetadataFromMessage", etcdKeyParams)
				return err
			}

			if err = GenerateEtcdClient(g, meta); err != nil {
				slog.Debug("error GenerateEtcdClient", etcdKeyParams)
				return err
			}

			if err = GenerateEtcdMethodGet(g, meta); err != nil {
				slog.Debug("error GenerateEtcdMethodGet", etcdKeyParams)
				return err
			}
		}
	}
	return nil
}

// getEtcdKeyParamsParams получает структуру etcd_key_params из опции сообщения
func getEtcdKeyParamsParams(message *protogen.Message) *ipb.EtcdKeyParams {
	opts := message.Desc.Options().(*descriptorpb.MessageOptions)
	if proto.HasExtension(opts, ipb.E_EtcdKeyParams) {
		keyTemplate := proto.GetExtension(opts, ipb.E_EtcdKeyParams).(*ipb.EtcdKeyParams)
		return keyTemplate
	}
	return nil
}
