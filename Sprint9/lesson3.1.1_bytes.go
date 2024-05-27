package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("привет")
	b := []byte("Привет")
	fmt.Println(bytes.Compare(a, b), bytes.Equal(a, b), bytes.EqualFold(a, b))
}
