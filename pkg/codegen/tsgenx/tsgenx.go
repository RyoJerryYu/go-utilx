package tsgenx

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type TSOption struct {
	GenFileSuffix string
	GenFilePath   string
}

type TSFileBuf struct {
	buf          bytes.Buffer
	Opts         TSOption
	ImportIdents map[string][]TSIdent // map<module_path, TSIdent>
}

func NewTSFileBuf(opts TSOption) *TSFileBuf {
	return &TSFileBuf{
		Opts:         opts,
		ImportIdents: make(map[string][]TSIdent),
	}
}

func (g *TSFileBuf) Apply(w io.Writer) error {
	if !strings.HasSuffix(g.Opts.GenFileSuffix, ".ts") {
		_, err := io.Copy(w, &g.buf)
		return err
	}
	content := g.buf.Bytes()
	imports := g.ImportSegments()
	res := fmt.Sprintf("%s\n\n%s", imports, content)
	_, err := w.Write([]byte(res))
	return err
}

func (g *TSFileBuf) Write(p []byte) (n int, err error) {
	return g.buf.Write(p)
}

func (g *TSFileBuf) P(v ...any) {
	for _, x := range v {
		switch x := x.(type) {
		case TSIdent:
			fmt.Fprint(&g.buf, g.QualifiedTSIdent(x))
		default:
			fmt.Fprint(&g.buf, x)
		}
	}
	fmt.Fprintln(&g.buf)
}

// Pf is same as P, but with formatted string.
func (opt *TSFileBuf) Pf(format string, v ...any) {
	newV := make([]any, len(v))
	for i, x := range v {
		switch x := x.(type) {
		case TSIdent:
			newV[i] = opt.QualifiedTSIdent(x)
		default:
			newV[i] = x
		}
	}
	opt.P(fmt.Sprintf(format, newV...))
}
