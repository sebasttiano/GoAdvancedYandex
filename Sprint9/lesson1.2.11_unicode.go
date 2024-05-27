package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	// считаем количество символов в строке
	str := "Привет, гофер!"
	// наивная попытка встроенной функцией len
	fmt.Println("Длина в байтах ", len(str))
	fmt.Println("Количество символов ", utf8.RuneCountInString(str))
	for _, c := range "\tIt's me! 7\n" {
		fmt.Printf("%q:", c)
		if unicode.IsControl(c) {
			fmt.Print(" control")
		}
		if unicode.IsDigit(c) {
			fmt.Print(" digit")
		}
		if unicode.IsLetter(c) {
			fmt.Print(" letter")
		}
		if unicode.IsLower(c) {
			fmt.Print(" lower")
		}
		if unicode.IsPunct(c) {
			fmt.Print(" punct")
		}
		if unicode.IsSpace(c) {
			fmt.Print(" space")
		}
		if unicode.IsUpper(c) {
			fmt.Print(" upper")
		}
		fmt.Print("\n")
	}
}
