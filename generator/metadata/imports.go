package metadata

import "google.golang.org/protobuf/compiler/protogen"

func CreateImportResolvers(out *protogen.GeneratedFile) map[string]*ImportResolver {
	return map[string]*ImportResolver{
		"context":  newImportResolver(out, "context"),
		"fmt":      newImportResolver(out, "encoding/json"),
		"clientv3": newImportResolver(out, "go.etcd.io/etcd/client/v3"),
		"time":     newImportResolver(out, "time"),
		"errors":   newImportResolver(out, "errors"),
	}
}
