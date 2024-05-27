package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	myConvert := func(r rune) rune {
		switch {
		case unicode.IsUpper(r):
			return unicode.ToLower(r)
		case unicode.IsPunct(r):
			return -1
		}
		return r
	}
	fmt.Printf("%s", bytes.Map(myConvert, []byte("Доброе Утро, Страна!")))
}
