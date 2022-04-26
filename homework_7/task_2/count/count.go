package count

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func CountFunc(n string, nameFunc string) (int, error) {
	count := 0
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, n, nil, 0)
	if err != nil {
		return 0, err
	}

	// Данная конструкция не работает и учитывает только 1 запуск hello()
	// К сожалению не смог придумать как учитывать запуск функции hello() в горутинах
	for _, f := range node.Decls {
		fn, ok := f.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if nameFunc == (fn.Name.Name) {
			count++
		}
	}
	return count, nil
}
