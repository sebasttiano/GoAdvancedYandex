package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.Replace([]byte("тест тест"), []byte("т"), []byte("Т"), 3)
	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", bytes.ReplaceAll(b, []byte("ес"), []byte("ЕС")))
}
