package gogenx

import (
	"fmt"
	"go/token"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type GoImportPath string

type GoIdent struct {
	GoImportPath GoImportPath
	Name         string
}

func (m GoImportPath) Ident(name string) GoIdent {
	return GoIdent{
		GoImportPath: m,
		Name:         name,
	}
}

type GoPackageName string

// goSanitized converts a string to a valid Go identifier.
func goSanitized(s string) string {
	// Sanitize the input to the set of valid characters,
	// which must be '_' or be in the Unicode L or N categories.
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return '_'
	}, s)

	// Prepend '_' in the event of a Go keyword conflict or if
	// the identifier is invalid (does not start in the Unicode L category).
	r, _ := utf8.DecodeRuneInString(s)
	if token.Lookup(s).IsKeyword() || !unicode.IsLetter(r) {
		return "_" + s
	}
	return s
}

// cleanPackageName converts a string to a valid Go package name.
func cleanPackageName(name string) GoPackageName {
	return GoPackageName(goSanitized(name))
}

func (g *GoFileBuf) ImportStatementItems() []string {
	imports := make([]string, 0, len(g.packageNames))
	for importPath := range g.packageNames {
		imports = append(imports, fmt.Sprintf("%s \"%s\"", g.packageNames[importPath], importPath))
	}
	sort.Slice(imports, func(i, j int) bool {
		return imports[i] < imports[j]
	})
	return imports
}

func (g *GoFileBuf) ImportStatement() string {
	return fmt.Sprintf("import (\n%s\n)", strings.Join(g.ImportStatementItems(), "\n"))
}
