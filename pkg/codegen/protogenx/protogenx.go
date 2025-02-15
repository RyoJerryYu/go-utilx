package protogenx

import (
	"bytes"
	"fmt"
	"io"
)

type ProtoOption struct {
	PackageName       ProtoPackage
	GenFileImportPath string
	Syntax            string // default "proto3"
}

type ProtoFileBuf struct {
	buf          bytes.Buffer
	Opts         ProtoOption
	imports      []ProtoImport
	packageNames map[ProtoPackage]string
}

func NewProtoFileBuf(opts ProtoOption) *ProtoFileBuf {
	return &ProtoFileBuf{
		Opts:         opts,
		imports:      make([]ProtoImport, 0),
		packageNames: make(map[ProtoPackage]string),
	}
}

func (f *ProtoFileBuf) Apply(w io.Writer) error {
	content := f.buf.Bytes()
	imports := f.ImportStatement()
	syntax := f.Opts.Syntax
	if syntax == "" {
		syntax = "proto3"
	}
	res := fmt.Sprintf(`syntax = "%s";

package %s;

%s

%s

`, syntax, f.Opts.PackageName, imports, content)
	_, err := w.Write([]byte(res))
	return err
}

func (f *ProtoFileBuf) Write(p []byte) (n int, err error) {
	return f.buf.Write(p)
}

func (f *ProtoFileBuf) P(v ...any) {
	newV := f.pConv(v...)
	for _, x := range newV {
		fmt.Fprint(&f.buf, x)
	}
	fmt.Fprintln(&f.buf)
}

func (f *ProtoFileBuf) Pf(format string, v ...any) {
	newV := f.pConv(v...)
	fmt.Fprintf(&f.buf, format, newV...)
	fmt.Fprintln(&f.buf)
}

func (f *ProtoFileBuf) pConv(v ...any) []any {
	newV := make([]any, len(v))
	for i, x := range v {
		switch x := x.(type) {
		case ProtoIdent:
			newV[i] = f.QualifiedProtoIdent(x)
		case Comments:
			newV[i] = x.String()
		default:
			newV[i] = x
		}
	}
	return newV
}
