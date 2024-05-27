package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer

	b.Grow(32)
	b.Write([]byte("123"))
	b.WriteString("Hello")
	b.WriteByte(33)
	b.WriteRune('Я')
	fmt.Fprintf(&b, " %d + %d = %d", 2, 3, 5)
	// печатаем содержимое буфера
	fmt.Printf("%s", &b)

	// чтение из буфера
	var target bytes.Buffer
	tmp := make([]byte, 5)
	b.Read(tmp)          // читаем 5 байт в tmp
	b.WriteString(`end`) // дописываем в буфер
	target.ReadFrom(&b)  // читаем остальные байты из b в target
	fmt.Printf(" => %s:%s", tmp, &target)
}
