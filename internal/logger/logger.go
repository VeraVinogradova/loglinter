package logger

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var pkgs = map[string]bool{
	"log/slog":        true,
	"go.uber.org/zap": true,
}

var funcs = map[string]bool{
	"Info": true, "Debug": true, "Warn": true, "Error": true,
}

func IsLogger(pass *analysis.Pass, sel *ast.SelectorExpr) bool {
	if !funcs[sel.Sel.Name] {
		return false
	}

	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}

	obj := pass.TypesInfo.ObjectOf(ident)
	if obj == nil || obj.Pkg() == nil {
		return false
	}

	return pkgs[obj.Pkg().Path()]
}
