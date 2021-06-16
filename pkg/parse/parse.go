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
	Name string
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

	walker := &ASTWalker{fs}
	walker.Visit(astFile)
	return &ParseResult{}, nil
}

type ASTWalker struct {
	fset *token.FileSet
}

func (a *ASTWalker) Visit(node ast.Node) ast.Visitor {
	return a
}
