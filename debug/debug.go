package main

import (
	"github.com/alex-dev-master/protoc-gen-etcd/entities"
	"github.com/alex-dev-master/protoc-gen-etcd/generator"
	"io"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	opts := protogen.Options{}
	file, err := os.Open("./debug/code_generator_request.pb.bin")
	if err != nil {
		return err
	}
	in, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err = proto.Unmarshal(in, req); err != nil {
		return err
	}
	gen, err := opts.New(req)
	if err != nil {
		return err
	}
	cfg := &entities.Config{
		LogLevelDebug: ptr(true),
	}
	if err = generator.Run(cfg, gen); err != nil {
		return err
	}
	resp := gen.Response()
	out, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	if _, err = os.Stdout.Write(out); err != nil {
		return err
	}
	return nil
}

// ptr is a helper that returns a pointer to v.
func ptr[T any](v T) *T {
	return &v
}
