package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriterSize(os.Stdout, 32)
	w.Write([]byte("Ещё один пример "))
	w.WriteString("по работе с bufio")
	fmt.Fprintln(w, w.Available(), w.Buffered(), w.Size())
	w.Flush()
}
