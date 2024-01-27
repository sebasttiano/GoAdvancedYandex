package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	//Sun, 19 Sep 2021 15:42:00 MSK
	fmt.Println(now.Format(time.RFC1123))
}
