package protogenx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQualifiedProtoIdent(t *testing.T) {
	cases := []struct {
		ident ProtoIdent
		want  string
	}{
		{
			ident: ProtoIdent{ProtoImport: ProtoImport{ProtoPackage: "schemax.schema.v1.test", ImportPath: "schemax/schema/v1/test.proto"}, Name: "Message"},
			want:  "Message",
		},
		{
			ident: ProtoIdent{ProtoImport: ProtoImport{ProtoPackage: "schemax.schema.v1.test", ImportPath: "schemax/schema/v1/test2.proto"}, Name: "Message"},
			want:  "Message",
		},
		{
			ident: ProtoIdent{ProtoImport: ProtoImport{ProtoPackage: "schemax.schema.v1.plugin", ImportPath: "schemax/schema/v1/plugin.proto"}, Name: "Message"},
			want:  "plugin.Message",
		},
		{
			ident: ProtoIdent{ProtoImport: ProtoImport{ProtoPackage: "schemax.table.v1.test", ImportPath: "schemax/table/v1/test.proto"}, Name: "Message"},
			want:  "table.v1.test.Message",
		},
		{
			ident: ProtoIdent{ProtoImport: ProtoImport{ProtoPackage: "google.protobuf.compiler", ImportPath: "google/protobuf/compiler/plugin.proto"}, Name: "Message"},
			want:  "google.protobuf.compiler.Message",
		},
		{
			ident: ProtoIdent{ProtoImport: ProtoImport{ProtoPackage: "schemax", ImportPath: "schemax/test.proto"}, Name: "Message"},
			want:  "Message",
		},
	}

	for _, c := range cases {
		t.Run(c.ident.Name, func(t *testing.T) {
			f := NewProtoFileBuf(ProtoOption{
				PackageName:       "schemax.schema.v1.test",
				GenFileImportPath: "schemax/schema/v1/test.proto",
			})

			got := f.QualifiedProtoIdent(c.ident)
			require.Equal(t, c.want, got)
			if c.ident.ProtoImport.ImportPath != f.Opts.GenFileImportPath {
				assert.Contains(t, f.ImportStatement(), c.ident.ProtoImport.ImportPath)
			}
		})
	}
}
