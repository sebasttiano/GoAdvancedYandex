package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {
	src := `/* Тестовый пакет */
package main

// Double умножает значение на 2.
func Double(i int) int {
    return i*2
}

func main() {
   // умножаем в цикле
   for i := 1; i < 5; i++ {
      fmt.Println(Double(i))
   }
}`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}
	for _, gr := range f.Comments {
		for _, c := range gr.List {
			fmt.Printf("%v %v \n", fset.Position(c.Slash), c.Text)
		}
	}
	// допишите код
	// ...
}
