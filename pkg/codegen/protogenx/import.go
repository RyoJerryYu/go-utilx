package protogenx

import (
	"fmt"
	"strings"

	"github.com/RyoJerryYu/go-utilx/pkg/container/slicex"
)

// ProtoPackage represents the protobuf package name.
//
// google.protobuf.compiler
type ProtoPackage string

func (p ProtoPackage) Import(path string) ProtoImport {
	return ProtoImport{
		ProtoPackage: p,
		ImportPath:   path,
	}
}

func (p ProtoPackage) Ident(path string, name string) ProtoIdent {
	return ProtoIdent{
		ProtoImport: ProtoImport{
			ProtoPackage: p,
			ImportPath:   path,
		},
		Name: name,
	}
}

// ProtoImport represents a file imported by the protobuf file.
type ProtoImport struct {
	ProtoPackage ProtoPackage // google.protobuf.compiler
	ImportPath   string       // google/protobuf/compiler/plugin.proto
}

func (p ProtoImport) Ident(name string) ProtoIdent {
	return ProtoIdent{
		ProtoImport: p,
		Name:        name,
	}
}

// ProtoIdent represents a protobuf identifier.
type ProtoIdent struct {
	ProtoImport ProtoImport // which file the identifier is in
	Name        string      // the identifier name, e.g. "Message"
}

// for package schemax.schema.v1.test to import ident schemax.schema.v1.test.Message
// just use Message
//
// for package schemax.schema.v1.test to import ident schemax.schema.v1.plugin.Message
// use plugin.Message
//
// for package schemax.schema.v1.test to import ident schemax.table.v1.test.Message
// use table.v1.test.Message
//
// for package schemax.schema.v1.test to import ident google.protobuf.compiler.Message
// use google.protobuf.compiler.Message
//
// for package schemax.schema.v1.test to import ident schemax.schema.Message
// use Message
func (p *ProtoFileBuf) QualifiedProtoIdent(ident ProtoIdent) string {
	if ident.ProtoImport.ProtoPackage == p.Opts.PackageName &&
		ident.ProtoImport.ImportPath == p.Opts.GenFileImportPath {
		return ident.Name
	}

	if packageName, ok := p.packageNames[ident.ProtoImport.ProtoPackage]; ok {
		return p.formatIdent(packageName, ident.Name)
	}

	importPackageParts := strings.Split(string(ident.ProtoImport.ProtoPackage), ".")
	thisPackageParts := strings.Split(string(p.Opts.PackageName), ".")

	i := 0
	for i < len(importPackageParts) && i < len(thisPackageParts) {
		if importPackageParts[i] != thisPackageParts[i] {
			break
		}
		i++
	}

	importPackageParts = importPackageParts[i:]
	packageName := strings.Join(importPackageParts, ".")

	p.imports = append(p.imports, ident.ProtoImport)
	p.packageNames[ident.ProtoImport.ProtoPackage] = packageName

	return p.formatIdent(packageName, ident.Name)
}

func (g *ProtoFileBuf) formatIdent(packageName string, ident string) string {
	if packageName == "" {
		return ident
	}
	return fmt.Sprintf("%s.%s", packageName, ident)
}

func (g *ProtoFileBuf) ImportStatements() []string {
	return slicex.To(g.imports, func(i ProtoImport) string {
		return fmt.Sprintf("import \"%s\";", i.ImportPath)
	})
}

func (g *ProtoFileBuf) ImportStatement() string {
	return strings.Join(g.ImportStatements(), "\n")
}
