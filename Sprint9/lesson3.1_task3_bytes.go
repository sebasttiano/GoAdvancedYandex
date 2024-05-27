package main

import (
	"bytes"
	"fmt"
)

func miniYandexTranslate(r rune) rune {
	switch r {
	case '😅':
		return 'a'
	case '😎':
		return 'e'
	case '😂':
		return 'g'
	case '🤓':
		return 'j'
	case '🥶':
		return 'h'
	case '🥰':
		return 'o'
	case '😶':
		return 'p'
	case '🐱':
		return 'r'
	default:
		return -1
	}
}

func main() {
	b := []byte("😂🥰😶🥶😎🐱🌚")
	fmt.Printf("%s\n", bytes.Map(miniYandexTranslate, b))
}
