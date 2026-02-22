package analyzer

import (
	"go/ast"
	"go/token"

	"github.com/VeraVinogradova/loglinter/internal/config"
	"github.com/VeraVinogradova/loglinter/internal/logger"
	"github.com/VeraVinogradova/loglinter/internal/rules"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func New() (*analysis.Analyzer, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return &analysis.Analyzer{
		Name:     "loglinter",
		Doc:      "checks log messages",
		Run:      run(cfg),
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}, nil
}

func run(cfg *config.Config) func(*analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

		inspect.Preorder([]ast.Node{(*ast.CallExpr)(nil)}, func(n ast.Node) {
			call := n.(*ast.CallExpr)

			msg, pos, lit, ok := extract(pass, call)
			if !ok {
				return
			}

			check(pass, pos, msg, lit, call, cfg)
		})

		return nil, nil
	}
}

func extract(pass *analysis.Pass, call *ast.CallExpr) (string, token.Pos, *ast.BasicLit, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", 0, nil, false
	}

	if !logger.IsLogger(pass, sel) {
		return "", 0, nil, false
	}

	if len(call.Args) == 0 {
		return "", 0, nil, false
	}

	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return "", 0, nil, false
	}

	if len(lit.Value) < 2 {
		return "", 0, nil, false
	}

	return lit.Value[1 : len(lit.Value)-1], lit.Pos(), lit, true
}

func check(pass *analysis.Pass, pos token.Pos, msg string, lit *ast.BasicLit, call *ast.CallExpr, cfg *config.Config) {
	checks := []struct {
		name string
		fn   func(string, *config.Config) (bool, string)
	}{
		{"lowercase", rules.CheckLowercase},
		{"english", rules.CheckEnglish},
		{"specialchars", rules.CheckSpecialChars},
		{"sensitive", rules.CheckSensitive},
	}

	for _, c := range checks {
		if ok, fix := c.fn(msg, cfg); !ok {
			diag := analysis.Diagnostic{
				Pos:     pos,
				Message: c.name + " rule violation",
			}
			if fix != "" && cfg.Fix {
				diag.SuggestedFixes = []analysis.SuggestedFix{{
					Message: "fix " + c.name,
					TextEdits: []analysis.TextEdit{{
						Pos:     lit.Pos(),
						End:     lit.End(),
						NewText: []byte(`"` + fix + `"`),
					}},
				}}
			}
			pass.Report(diag)
			return
		}
	}
}
