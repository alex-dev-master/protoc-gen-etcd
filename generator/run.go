package generator

import (
	"errors"
	"github.com/alex-dev-master/protoc-gen-etcd/entities"
	"github.com/alex-dev-master/protoc-gen-etcd/generator/extensions"
	"google.golang.org/protobuf/compiler/protogen"
	"log/slog"
	"os"
)

func Run(cfg *entities.Config, plugin *protogen.Plugin) (err error) {
	logLevel := slog.LevelInfo
	if *cfg.LogLevelDebug {
		logLevel = slog.LevelDebug
	}

	lTextHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: false,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})
	slog.SetDefault(slog.New(lTextHandler))

	slog.Debug("Protogen plugin called with following files to be generated", "files", plugin.Request.FileToGenerate)

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		g := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".etcd.go", file.GoImportPath)
		GenerateHeader(g, file)

		for _, service := range file.Services {
			etcdOpts := extensions.GetEtcdOptions(service)
			if etcdOpts == nil {
				return errors.New("error getEtcdOptions")
			}
			if err = processServiceLayer(g, etcdOpts); err != nil {
				slog.Debug("error processServiceLayer")
				return err
			}

			for _, method := range service.Methods {
				if err = processMethodLayer(g, method, etcdOpts.GetServiceKeyPrefix()); err != nil {
					slog.Debug("error processMethodLayer")
					return err
				}
			}
		}
	}
	return nil
}
