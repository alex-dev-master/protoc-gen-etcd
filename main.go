package main

import (
	"flag"
	"github.com/alex-dev-master/protoc-gen-etcd/entities"
	"github.com/alex-dev-master/protoc-gen-etcd/generator"
	"google.golang.org/protobuf/compiler/protogen"
)

var flags flag.FlagSet

func main() {
	cfg := &entities.Config{
		LogLevelDebug: flags.Bool("logLevelDebug", false, "enable debug log level"),
	}

	opts := protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(generator.NewGenerator(cfg).Run)
}
