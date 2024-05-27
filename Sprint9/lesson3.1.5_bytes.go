package main

import (
	"bytes"
	"fmt"
)

func main() {
	items := bytes.Split([]byte("т,е,с,т"), []byte(","))
	joined := bytes.Join(items, []byte("-"))
	fmt.Printf("%q %s %q\n", items, joined, bytes.SplitAfter(joined, []byte("-")))
}
