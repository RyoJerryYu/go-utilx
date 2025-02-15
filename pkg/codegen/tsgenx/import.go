package tsgenx

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/RyoJerryYu/go-utilx/pkg/container/containerx"
	"github.com/RyoJerryYu/go-utilx/pkg/container/slicex"
)

type TSModule struct {
	ModuleName string
	Path       string // path relative to the generate root, or the absolute path
	Relative   bool   // whether the path is relative to the current file
}

func (m TSModule) Ident(name string) TSIdent {
	return TSIdent{
		TSModule: m,
		Name:     name,
	}
}

func (m TSModule) AsIdent() TSIdent {
	return TSIdent{
		TSModule: m,
	}
}

type TSIdent struct {
	TSModule
	Name string
}

func (m TSIdent) GetName() string {
	return m.Name
}

func (r *TSFileBuf) QualifiedTSIdent(ident TSIdent) string {
	if _, ok := r.ImportIdents[ident.Path]; !ok {
		r.ImportIdents[ident.Path] = []TSIdent{}
	}
	r.ImportIdents[ident.Path] = append(r.ImportIdents[ident.Path], ident)
	return ident.Name
}

func tsRelativeImportPath(thisPath string, modulePath string) (string, bool) {
	thisDir := filepath.Dir(thisPath)
	relativePath, err := filepath.Rel(thisDir, modulePath)
	if err != nil {
		return "", false
	}
	if !strings.Contains(relativePath, "/") && !strings.HasPrefix(relativePath, ".") {
		relativePath = "./" + relativePath
	}
	return strings.TrimSuffix(relativePath, ".ts"), true
}

func (g *TSFileBuf) thisModulePath() string {
	return g.Opts.GenFilePath
}

func (g *TSFileBuf) ImportSegments() string {
	thisModulePath := g.thisModulePath()
	var imports []string
	modulePaths := make([]string, 0, len(g.ImportIdents))
	for path := range g.ImportIdents {
		modulePaths = append(modulePaths, path)
	}
	// sort by module import path
	sort.Slice(modulePaths, func(i, j int) bool {
		return modulePaths[i] < modulePaths[j]
	})

	for _, modulePath := range modulePaths {
		idents := g.ImportIdents[modulePath]
		module := idents[0].TSModule
		importPath := module.Path
		ok := false
		if module.Relative {
			importPath, ok = tsRelativeImportPath(thisModulePath, module.Path)
			if !ok {
				continue
			}
		}
		imports = append(imports, g.importSegmentDirect(importPath, idents))
	}
	return strings.Join(imports, "\n")
}

func (g *TSFileBuf) importSegmentDirect(importPath string, idents []TSIdent) string {
	if len(idents) == 0 {
		return ""
	}
	identNames := containerx.ToNames(idents)
	identNames = slicex.Deduplicate(identNames)
	sort.Slice(identNames, func(i, j int) bool {
		return identNames[i] < identNames[j]
	})
	importPackageNum := 0
	for _, identName := range identNames {
		if identName == "" {
			importPackageNum += 1
			break
		}
	}

	importContents := []string{}
	if importPackageNum > 0 {
		importContents = append(importContents, idents[0].ModuleName)
	}
	if len(identNames) > importPackageNum {
		importContents = append(importContents, fmt.Sprintf("{ %s }", strings.Join(identNames[importPackageNum:], ", ")))
	}

	return fmt.Sprintf("import %s from '%s';", strings.Join(importContents, ", "), importPath)
}
