package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("Привет, мир!")

	fmt.Println(bytes.Contains(b, []byte("мир")), bytes.HasPrefix(b, []byte("При")))
	fmt.Println(bytes.IndexRune(b, 'м'), bytes.LastIndexByte(b, '!'), bytes.IndexAny(b, "abc"))
}
