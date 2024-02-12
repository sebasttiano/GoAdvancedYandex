package main

import "fmt"

func main() {
	chIn := make(chan int)
	chOut := make(chan int)
	quit := make(chan struct{})

	go func() {
		for i := 0; i < 15; i++ {
			chIn <- i
		}
		close(quit)
	}()
	go func() {
		var x int
		for x = range chIn {
			chOut <- x * 2
		}
	}()
	go func() {
		for x := range chOut {
			fmt.Printf("%d ", x)
		}
		quit <- struct{}{}
	}()
	<-quit
}
