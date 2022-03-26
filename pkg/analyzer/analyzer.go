package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "gofuncdocprefix",
	Doc:      "Checks that function documentation starts with the function name.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.FuncDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		funcDecl := node.(*ast.FuncDecl)

		if strings.HasPrefix(funcDecl.Doc.Text()+" ", funcDecl.Name.Name) {
			return
		}

		firstLine := strings.Split(funcDecl.Doc.Text(), "\n")[0]
		pass.Reportf(node.Pos(), "comment '%s' should start with function name '%s'", firstLine, funcDecl.Name.Name)
	})

	return nil, nil
}
