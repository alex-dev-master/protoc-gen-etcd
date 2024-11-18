package metadata

import "google.golang.org/protobuf/compiler/protogen"

type ImportResolver struct {
	f          *protogen.GeneratedFile
	importPath protogen.GoImportPath
}

func newImportResolver(gf *protogen.GeneratedFile, pkg string) *ImportResolver {
	return &ImportResolver{
		f:          gf,
		importPath: protogen.GoImportPath(pkg),
	}
}

func (ir *ImportResolver) Ident(s string) string {
	return ir.f.QualifiedGoIdent(ir.importPath.Ident(s))
}

func (ir *ImportResolver) Method(s string) string {
	return ir.f.QualifiedGoIdent(ir.importPath.Ident(s))
}
