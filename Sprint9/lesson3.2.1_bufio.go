package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	w.WriteString("Привет, ") // пишем в w
	fmt.Fprint(w, "мир!")     // другой метод записи в w

	fmt.Println(w.Available(), w.Buffered(), w.Size())
	w.Flush()
}
