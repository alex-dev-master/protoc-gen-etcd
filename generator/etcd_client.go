package generator

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/alex-dev-master/protoc-gen-etcd/generator/metadata"
	"google.golang.org/protobuf/compiler/protogen"
	"log/slog"
	"text/template"
)

//go:embed template/etcd_client.tmpl
var etcdClientTmpl string

// GenerateEtcdClient генерирует код клиента etcd
func GenerateEtcdClient(
	g *protogen.GeneratedFile,
	meta *metadata.EtcdClientMetadata,
) (err error) {
	t, err := template.New("etcd-client").Parse(etcdClientTmpl)
	if err != nil {
		slog.Debug("error template.New", err)
		return err
	}

	genRes, err := GenerateTemplateFile(meta, t)
	if err != nil {
		slog.Debug("error GenerateTemplateFile", err)
		return err
	}

	_, err = g.Write(genRes)
	if err != nil {
		slog.Debug("error GenerateTemplateFile", err)
		return err
	}
	return nil
}

func GenerateTemplateFile(
	meta *metadata.EtcdClientMetadata,
	template *template.Template,
) ([]byte, error) {
	var buf bytes.Buffer
	if err := template.Execute(&buf, meta); err != nil {
		return nil, fmt.Errorf("failed to generate remplate: %w", err)
	}
	return buf.Bytes(), nil
}
