package generator

import (
	_ "embed"
	"fmt"
	"github.com/alex-dev-master/protoc-gen-etcd/generator/metadata"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

// GenerateEtcdMethodGet генерирует код GET метода для ключа
func GenerateEtcdMethodGet(
	g *protogen.GeneratedFile,
	meta *metadata.EtcdMethodMetadata,
) (err error) {
	keyPath := meta.EtcdKeyOptions.GetKeyPath()
	params := make([]string, 0)
	if len(meta.EtcdKeyParamOptions) > 0 {
		if keyPath, params, err = meta.GetMethodRequestParams(); err != nil {
			return err
		}
	}

	g.P()
	g.P(`func (c *EtcdClient) Get(ctx `, contextPackage.Ident("Context"), `, rq *`, meta.InputRequest, `) (resp *`, meta.ValueOfKey, `, err error) {`)
	keyStr := `	key := ` + keyPath
	if len(params) > 0 {
		keyStr = fmt.Sprintf(`	key := fmt.Sprintf("%s", %s)`, keyPath, strings.Join(params, ", "))
	}
	g.P(keyStr)
	g.P(`	resp, err := c.client.Get(ctx, key)`)
	g.P(`	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return nil, fmt.Errorf("key not found")
	}

	for _, v := range resp.Kvs {`)
	genUnmarshalResponseStruct(g)

	g.P(`		if err != nil {
			return nil, err
		}
		return resp, nil
	}`)
	g.P(`return resp, err
}`)
	return nil
}

// genUnmarshalResponseStruct generates unmarshalling from []byte to struct for response
func genUnmarshalResponseStruct(g *protogen.GeneratedFile) {
	g.P("		err = ", jsonPackage.Ident("Unmarshal"), "(v.Value, resp)")
}
