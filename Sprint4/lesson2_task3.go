package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[int]int)
	var mx sync.Mutex
	for i := 0; i < 100; i++ {
		go func(v int) {
			mx.Lock()
			m[v] = 1
			mx.Unlock()
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(len(m))
}
