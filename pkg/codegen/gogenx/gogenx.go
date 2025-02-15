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
