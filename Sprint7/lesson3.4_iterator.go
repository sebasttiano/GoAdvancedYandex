package main

import "fmt"

func newEven() func() int {
	n := 0
	// функциональный литерал замкнёт переменную n
	return func() int {
		n += 2
		return n
	}
}

func main() {
	next := newEven()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
