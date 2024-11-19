package generator

import (
	"errors"
	"github.com/alex-dev-master/protoc-gen-etcd/entities"
	"github.com/alex-dev-master/protoc-gen-etcd/generator/extension"
	"github.com/alex-dev-master/protoc-gen-etcd/generator/metadata"
	"google.golang.org/protobuf/compiler/protogen"
	"log/slog"
	"os"
)

type (
	generator struct {
		cfg     *entities.Config
		imports map[string]*metadata.ImportResolver
	}
)

func NewGenerator(cfg *entities.Config) *generator {
	return &generator{
		cfg: cfg,
	}
}

func (g *generator) Run(plugin *protogen.Plugin) (err error) {
	g.initLogger()
	slog.Debug("Protogen plugin called with following files to be generated", "files", plugin.Request.FileToGenerate)

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		genFile := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+".etcd.go", file.GoImportPath)
		g.imports = metadata.CreateImportResolvers(genFile)

		GenerateHeader(genFile, file)

		for _, service := range file.Services {
			etcdOpts := extension.GetEtcdOptions(service)
			if etcdOpts == nil {
				return errors.New("error getEtcdOptions")
			}
			if err = g.processServiceLayer(genFile, etcdOpts); err != nil {
				slog.Debug("error processServiceLayer")
				return err
			}

			for _, method := range service.Methods {
				if err = g.processMethodLayer(genFile, method, etcdOpts.GetServiceKeyPrefix()); err != nil {
					slog.Debug("error processMethodLayer")
					return err
				}
			}
		}
	}
	return nil
}

func (g *generator) initLogger() {
	logLevel := slog.LevelInfo
	if *g.cfg.LogLevelDebug {
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
}
