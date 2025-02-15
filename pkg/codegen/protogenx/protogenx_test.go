package protogenx

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtoFileBuf(t *testing.T) {
	f := NewProtoFileBuf(ProtoOption{
		PackageName:       "api.v1.test",
		GenFileImportPath: "api/v1/test.proto",
	})
	f.P("test")
	f.Pf("test %s", "test")

	f.P(ProtoIdent{ProtoImport: ProtoImport{ProtoPackage: "schemax.schema.v1.test", ImportPath: "schemax/schema/v1/test.proto"}, Name: "Message"})

	res := bytes.NewBufferString("")
	f.Apply(res)

	assert.Equal(t, `syntax = "proto3";

package api.v1.test;

import "schemax/schema/v1/test.proto";

test
test test
schemax.schema.v1.test.Message


`, res.String())
}
