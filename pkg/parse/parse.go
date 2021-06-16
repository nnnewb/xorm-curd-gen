package parse

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

type ParseResult struct {
	Models []Model
}

type Model struct {
	Name   string
	Fields []*ast.Field
}

func ParseFile(f string) (*ParseResult, error) {
	fs := token.NewFileSet()
	content, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	astFile, err := parser.ParseFile(fs, f, content, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return nil, err
	}

	models := make([]Model, 0)

	ast.Inspect(astFile, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			// We only care about const declarations.
			return true
		}

		for _, spec := range decl.Specs {
			tspec := spec.(*ast.TypeSpec)
			if stype, ok := tspec.Type.(*ast.StructType); ok {
				models = append(models, Model{
					Name:   tspec.Name.Name,
					Fields: stype.Fields.List,
				})
			}
		}

		return false
	})

	return &ParseResult{models}, nil
}
