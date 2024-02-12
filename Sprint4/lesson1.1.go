package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// NewSource возвращает новый псевдослучайный источник с заданным значением.
	sec1 := rand.New(rand.NewSource(10))
	sec2 := rand.New(rand.NewSource(10))

	for i := 0; i < 5; i++ {
		// генерация случайных значений из источников
		rnd1 := sec1.Int()
		rnd2 := sec2.Int()
		if rnd1 != rnd2 {
			fmt.Println("Сгенерированная случайным образом последовательность")
			break
		} else {
			fmt.Printf("rnd1: %d, rnd2: %d\n", rnd1, rnd2)
		}
	}
}
