package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	generate(16)
}
func generate(number int) {
	// определяем слайс байт нужной длины
	b := make([]byte, number)
	_, err := rand.Read(b) // записываем байты в слайс b
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Println(base64.StdEncoding.EncodeToString(b))
}
