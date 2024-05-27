package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := []byte("🌝🌖🌗🌘🌚🌒🌓🌔🌝")
	fmt.Println(bytes.IndexRune(b, '🌚'))
	// допишите код
}
