package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	for i := 1; i <= 10; i++ {
		ticker := time.NewTicker(2 * time.Second)
		tickTime := <-ticker.C
		fmt.Println(int(tickTime.Sub(start).Seconds()))
	}
}
