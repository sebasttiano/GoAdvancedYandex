package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	src := `package main
    
func main() {
     ids := 77
     id := ids + 1
     fmt.Println("id равно:", id/2 )
}`

	// FileSet
	fset := token.NewFileSet()
	// парсинг
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	// обходим дерево разбора
	ast.Inspect(f, func(n ast.Node) bool {
		// если узел имеет тип *ast.Ident
		if v, ok := n.(*ast.Ident); ok {
			if v.Name == "id" {
				// изменяем узел
				v.Name = "Ident"
			}
		}
		return true
	})
	// выливаем изменённый AST
	// обратно в текст исходного кода
	// и выводим в консоль
	printer.Fprint(os.Stdout, fset, f)

	// допишите код
	// ...
}
