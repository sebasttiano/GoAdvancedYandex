package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	b1 := make([]byte, 10)
	b2 := make([]byte, 4)
	b3 := make([]byte, 8)

	r := bytes.NewReader([]byte("Привет, мир!"))
	// читается столько байт, сколько соответствует размеру целевого слайса
	// один символ кириллицы занимает два байта
	r.Read(b1)               // в b1 запишется "Приве"
	r.ReadAt(b2, 4)          // в b2 запишется "ив"
	r.Seek(14, io.SeekStart) // встали на символе 'м'
	r.Read(b3)               // в b3 запишется "мир!"
	fmt.Printf("%d %s %s %s\n", r.Size(), b1, b2, b3)
}
