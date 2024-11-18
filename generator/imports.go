package generator

import "google.golang.org/protobuf/compiler/protogen"

var (
	contextPackage = protogen.GoImportPath("context")
	jsonPackage    = protogen.GoImportPath("encoding/json")

	easyjsonPackage  = protogen.GoImportPath("github.com/mailru/easyjson")
	protojsonPackage = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
)
