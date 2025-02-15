package gogenx

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type GoOption struct {
	GenFileSuffix     string
	GenFileImportPath GoImportPath
}

type GoFileBuf struct {
	buf              bytes.Buffer
	Opts             GoOption
	packageNames     map[GoImportPath]GoPackageName
	usedPackageNames map[GoPackageName]bool
}

func NewGoFileBuf(opts GoOption) *GoFileBuf {
	return &GoFileBuf{
		Opts:             opts,
		packageNames:     make(map[GoImportPath]GoPackageName),
		usedPackageNames: make(map[GoPackageName]bool),
	}
}

func (g *GoFileBuf) Apply(w io.Writer) error {
	if !strings.HasSuffix(g.Opts.GenFileSuffix, ".go") {
		_, err := io.Copy(w, &g.buf)
		return err
	}
	content := g.buf.Bytes()
	imports := g.ImportStatement()
	res := fmt.Sprintf("%s\n\n%s", imports, content)
	_, err := w.Write([]byte(res))
	return err
}

func (g *GoFileBuf) Write(p []byte) (n int, err error) {
	return g.buf.Write(p)
}

func (g *GoFileBuf) P(v ...any) {
	newV := g.pConv(v...)
	for _, x := range newV {
		fmt.Fprint(&g.buf, x)
	}
	fmt.Fprintln(&g.buf)
}

func (g *GoFileBuf) Pf(format string, v ...any) {
	newV := g.pConv(v...)
	fmt.Fprintf(&g.buf, format, newV...)
	fmt.Fprintln(&g.buf)
}

func (g *GoFileBuf) pConv(v ...any) []any {
	newV := make([]any, len(v))
	for i, x := range v {
		switch x := x.(type) {
		case GoIdent:
			newV[i] = g.QualifiedGoIdent(x)
		case Comments:
			newV[i] = x.String()
		default:
			newV[i] = x
		}
	}
	return newV
}
