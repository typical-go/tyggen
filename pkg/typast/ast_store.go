package typast

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
)

// ASTStore responsible to store filename, declaration and annotation
type ASTStore struct {
	Filenames []string
	Decls     []*Decl
	Annots    []*Annotation
}

// CreateASTStore to walk through the filenames and store declaration and annotations
func CreateASTStore(filenames ...string) (a *ASTStore) {
	var (
		decls  []*Decl
		annots []*Annotation
		err    error
	)

	if decls, err = parseFiles(filenames); err != nil {
		panic(err.Error())
	}

	for _, decl := range decls {
		annots = append(annots, parseAnnots(decl)...)
	}

	return &ASTStore{
		Filenames: filenames,
		Decls:     decls,
		Annots:    annots,
	}
}

func parseFiles(filenames []string) (decls []*Decl, err error) {
	var (
		f *ast.File
	)
	fset := token.NewFileSet() // positions are relative to fset
	for _, filename := range filenames {
		if f, err = parser.ParseFile(fset, filename, nil, parser.ParseComments); err != nil {
			return
		}
		for _, d := range f.Decls {
			if decl := parseDecl(filename, f, d); decl != nil {
				decls = append(decls, decl)
			}
		}
	}
	return
}

func parseDecl(path string, f *ast.File, decl ast.Decl) *Decl {
	switch decl.(type) {
	case *ast.FuncDecl:
		funcDecl := decl.(*ast.FuncDecl)
		return &Decl{
			Type:       Function,
			SourceName: funcDecl.Name.Name,
			SourceObj:  funcDecl,
			Path:       path,
			File:       f,
			Doc:        funcDecl.Doc,
		}
	case *ast.GenDecl:
		genDecl := decl.(*ast.GenDecl)
		for _, spec := range genDecl.Specs {
			switch spec.(type) {
			case *ast.TypeSpec:
				typeSpec := spec.(*ast.TypeSpec)
				declType := Generic
				switch typeSpec.Type.(type) {
				case *ast.InterfaceType:
					declType = Interface
				case *ast.StructType:
					declType = Struct
				}
				return &Decl{
					Type:       declType,
					SourceName: typeSpec.Name.Name,
					SourceObj:  typeSpec,
					Path:       path,
					File:       f,
					Doc:        genDecl.Doc,
				}
			}
		}
	}
	return nil
}

func parseAnnots(decl *Decl) (annotations []*Annotation) {
	if decl.Doc == nil {
		return
	}

	r, _ := regexp.Compile("\\[(.*?)\\]")
	for _, s := range r.FindAllString(decl.Doc.Text(), -1) {
		if a := CreateAnnotation(decl, s); a != nil {
			annotations = append(annotations, a)
		}
	}
	return
}
