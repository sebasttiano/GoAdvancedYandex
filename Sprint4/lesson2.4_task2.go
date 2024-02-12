package main

import (
	"fmt"
)

func main() {
	ch := generator2("Hello")
	for msg := range ch {
		fmt.Println(msg)
	}
}

// Тут ваш генератор
func generator2(msg string) chan string {
	inputCh := make(chan string)

	go func() {
		defer close(inputCh)
		for i := 0; i < 5; i++ {
			inputCh <- fmt.Sprintf("%s %d", msg, i)
		}
	}()

	return inputCh
}
