package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	r := bufio.NewReader(strings.NewReader("Пример для bufio.Reader"))
	word, err := r.ReadString(' ') // читаем слово вместе с пробелом
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word, r.Buffered(), r.Size())

	r.Discard(7) // пропускаем три символа на кириллице и пробел

	// получаем пять символов, не сдвигая позицию чтения
	peek, err := r.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 24)
	r.Read(buf) // читаем оставшиеся данные
	fmt.Println(string(peek), string(buf))
}
