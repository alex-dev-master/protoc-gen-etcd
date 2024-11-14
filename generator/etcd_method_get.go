package generator

import (
	_ "embed"
	"google.golang.org/protobuf/compiler/protogen"
	"log/slog"
	"text/template"
)

//go:embed template/etcd_method_get.tmpl
var etcdMethodGetTmpl string

// GenerateEtcdMethodGet генерирует код GET метода для ключа
func GenerateEtcdMethodGet(
	g *protogen.GeneratedFile,
	meta *EtcdMetadata,
) (err error) {
	t, err := template.New("etcd-method-get").Parse(etcdMethodGetTmpl)
	if err != nil {
		slog.Debug("error template.New", err)
		return err
	}

	meta.getMethodRequestParams()

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
